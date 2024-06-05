package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"math/big"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/anagrambuild/ferrule/config"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/viper"
	"golang.org/x/crypto/hkdf"
)

var roundRegex = regexp.MustCompile(`.+Round(\d{1}).+`)

func GetRound(messageType string) uint32 {
	match := roundRegex.FindStringSubmatch(messageType)
	if len(match) < 2 {
		return 0
	}
	round, err := strconv.Atoi(match[1])
	if err != nil {
		return 0
	}
	log.Printf("Round %d", round)
	return uint32(round)
}

func LoadIdentity(path string) *ecdsa.PrivateKey {
	key, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	pkey, err := ethcrypto.HexToECDSA(string(key))
	if err != nil {
		panic(err)
	}
	return pkey
}

func LoadSecretShare(loc string) []byte {
	file, err := os.ReadFile(loc)
	if err != nil {
		fmt.Println("No secret share found, starting new committee.")
		return nil
	}
	return file
}

type ConfigRaw struct {
	IdentityKpPath string   `mapstructure:"IDENTITY_KEY_PATH"`
	SecretShareDir string   `mapstructure:"SECRET_SHARE_DIR"`
	Peers          []string `mapstructure:"PEERS"`
	LogLevel       string   `mapstructure:"LOG_LEVEL"`
	DataDir        string   `mapstructure:"DATA_DIR"`
	OpUrl          string   `mapstructure:"OP_URL"`
}

func InitConfig() *config.FerruleConfig {
	viper.SetEnvPrefix("FR")
	var cfg ConfigRaw
	for _, env := range os.Environ() {
		if strings.HasPrefix(env, "FR_") {
			split := strings.Split(env, "=")
			e := viper.BindEnv(strings.Replace(split[0], "FR_", "", 1))
			if e != nil {
				fmt.Println(e)
			}
		}
	}
	viper.AutomaticEnv()
	viper.Unmarshal(&cfg)
	client, err := ethclient.Dial(cfg.OpUrl)
	if err != nil {
		panic(err)
	}
	id := LoadIdentity(cfg.IdentityKpPath)
	loc := path.Join(cfg.SecretShareDir, "current.json")
	//todo peer encryption
	return &config.FerruleConfig{
		Identity:           id,
		NodeName:           PubKeyToNodeName(&id.PublicKey),
		Peers:              cfg.Peers,
		PeerEncryptionKey:  []byte{},
		DataDir:            cfg.DataDir,
		CurrentSecretShare: LoadSecretShare(loc),
		CommitteeStateFile: loc,
		ForkThreshold:      3,
		OpClient:           client,
	}
}

func NodeNameToPubKey(name string) (*ecdsa.PublicKey, error) {
	decoded, err := hex.DecodeString(name)
	if err != nil {
		return nil, err
	}
	return ethcrypto.DecompressPubkey(decoded)
}

func PubKeyToNodeName(key *ecdsa.PublicKey) string {
	compr := ethcrypto.CompressPubkey(key)
	compressedKey := hex.EncodeToString(compr)
	return compressedKey
}

func PrepareStorage(config config.FerruleConfig) {
	if _, err := os.Stat(config.DataDir); os.IsNotExist(err) {
		err := os.MkdirAll(config.DataDir, 0755)
		if err != nil {
			panic(err)
		}
	}
}

func PubkeyToBigInt(pubkey *ecdsa.PublicKey) *big.Int {
	return new(big.Int).SetBytes(secp256k1.S256().Marshal(pubkey.X, pubkey.Y))
}

func DeriveSharedSecret(priv *ecdsa.PrivateKey, pub *ecdsa.PublicKey) ([]byte, error) {
	ecdhPriv, err := priv.ECDH()
	if err != nil {
		return nil, err
	}
	ecdhPub, err := pub.ECDH()
	if err != nil {
		return nil, err
	}
	return ecdhPriv.ECDH(ecdhPub)
}

func DeriveSymKey(secret []byte, salt []byte) ([]byte, error) {
	derivedKey := make([]byte, 32)
	hk := hkdf.New(sha256.New, secret, salt, nil)
	_, err := io.ReadFull(hk, derivedKey)
	if err != nil {
		return nil, fmt.Errorf("failed to generate key: %w", err)
	}
	return derivedKey, nil
}

type SealedMessage struct {
	Nonce  []byte
	Cipher []byte
}

func EncryptMessage(
	key []byte,
	message []byte) (*SealedMessage, error) {
	nonce := make([]byte, 12)
	_, err := io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, fmt.Errorf("failed to generate nonce: %w", err)
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("failed to create cipher: %w", err)
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("failed to create GCM: %w", err)
	}
	return &SealedMessage{
		Nonce:  nonce,
		Cipher: aesgcm.Seal(nil, nonce, message, nil),
	}, nil
}

func DecryptMessage(
	key []byte,
	sealed *SealedMessage) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("failed to create cipher: %w", err)
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("failed to create GCM: %w", err)
	}
	return aesgcm.Open(nil, sealed.Nonce, sealed.Cipher, nil)
}

func NewSyncedMap[K comparable, V any]() SyncedMap[K, V] {
	return SyncedMap[K, V]{m: make(map[K]V)}
}

type SyncedMap[K comparable, V any] struct {
	mu sync.RWMutex
	m  map[K]V
}

func (m *SyncedMap[K, V]) Get(key K) V {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.m[key]
}

func (m *SyncedMap[K, V]) Set(key K, value V) {
	m.mu.Lock()
	m.m[key] = value
	m.mu.Unlock()
}

func (m *SyncedMap[K, V]) Delete(key K) {
	m.mu.Lock()
	delete(m.m, key)
	m.mu.Unlock()
}

func (m *SyncedMap[K, V]) Replace(new map[K]V) {
	m.mu.Lock()
	m.m = new
	m.mu.Unlock()
}
func (m *SyncedMap[K, V]) Len() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return len(m.m)
}

func (m *SyncedMap[K, V]) Items() []V {
	m.mu.RLock()
	defer m.mu.RUnlock()
	items := make([]V, 0, len(m.m))
	for _, item := range m.m {
		items = append(items, item)
	}
	return items
}

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](items ...T) Set[T] {
	set := make(Set[T])
	set.Add(items...)
	return set
}

func (s Set[T]) Add(items ...T) {
	for _, item := range items {
		s[item] = struct{}{}
	}
}

func (s Set[T]) Remove(items ...T) {
	for _, item := range items {
		delete(s, item)
	}
}

func (s Set[T]) Contains(item T) bool {
	_, exists := s[item]
	return exists
}

func (s Set[T]) Len() int {
	return len(s)
}

func (s Set[T]) Items() []T {
	items := make([]T, 0, len(s))
	for item := range s {
		items = append(items, item)
	}
	return items
}

func BackoffRetry(
	fn func() error,
	maxRetries int,
	dur time.Duration,
) error {
	var err error
	for i := 0; i < maxRetries; i++ {
		err = fn()
		if err == nil {
			return nil
		}
		if i == maxRetries-1 {
			return err
		}
		time.Sleep(dur)
	}
	return err
}
