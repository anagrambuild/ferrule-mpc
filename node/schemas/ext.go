package schemas

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/anagrambuild/ferrule/util"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
)

func DecodeClusterMessage(data []byte) (*ClusterMessage, error) {
	m := &ClusterMessage{}
	err := proto.Unmarshal(data, m)
	return m, err
}

func EncodeClusterMessage(m *ClusterMessage) ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ClusterMessage) EncodeContent() error {
	m.Id = []byte(uuid.New().String())
	content, err := proto.Marshal(m.ParsedContent)
	if err != nil {
		return err
	}
	m.Content = content
	m.ParsedContent = nil
	return nil
}

func (m *ClusterMessage) Sign(
	key *ecdsa.PrivateKey,
) error {
	hash := ethcrypto.Keccak256(m.GetContent())
	sig, err := ethcrypto.Sign(hash, key)
	if err != nil {
		return err
	}
	m.Signature = sig
	m.From = util.PubKeyToNodeName(&key.PublicKey)
	return nil
}

func (m *ClusterMessage) GetFromKey() (*ecdsa.PublicKey, error) {
	return util.NodeNameToPubKey(m.From)
}

func (m *ClusterMessage) GetFromEcies() (*ecies.PublicKey, error) {
	key, err := m.GetFromKey()
	if err != nil {
		return nil, err
	}
	return ecies.ImportECDSAPublic(key), nil
}

func (m *ClusterMessage) GetFromAddress() (string, error) {
	nodeKey, err := m.GetFromKey()
	if err != nil {
		return "", err
	}
	return ethcrypto.PubkeyToAddress(*nodeKey).Hex(), nil
}

func (m *ClusterMessage) Verify(
	key *ecdsa.PublicKey,
) bool {
	hash := ethcrypto.Keccak256(m.GetContent())
	pub, err := ethcrypto.SigToPub(hash, m.Signature)
	if err != nil {
		return false
	}
	return pub.Equal(key)
}

func (m *ClusterMessage) VerifySelf() bool {
	key, err := m.GetFromKey()
	fmt.Println(key)
	if err != nil {
		fmt.Println("error getting key", err)
		return false
	}
	return m.Verify(key)
}

func (m *ClusterMessage) DecodeContent() error {
	parsedContent := Content{}
	err := proto.Unmarshal(m.Content, &parsedContent)
	if err != nil {
		return err
	}
	m.ParsedContent = &parsedContent
	m.Content = nil
	return nil
}

func (m *ClusterMessage) Decrypt(
	key *ecies.PrivateKey,
) error {
	pubEc, err := m.GetFromEcies()
	if err != nil {
		return err
	}
	params := key.Params
	secret, err := key.GenerateShared(pubEc, params.KeyLen, params.KeyLen)
	if err != nil {
		return err
	}
	dk, err := util.DeriveSymKey(secret, m.Salt)
	if err != nil {
		return err
	}
	seal := util.SealedMessage{
		Nonce:  m.Nonce,
		Cipher: m.GetContent(),
	}
	content, err := util.DecryptMessage(dk, &seal)
	if err != nil {
		return err
	}
	m.Content = content
	m.DecodeContent()
	return nil
}

func (m *ClusterMessage) Encrypt(
	ecdsaKey *ecdsa.PrivateKey,
	key *ecies.PrivateKey,
	to *ecies.PublicKey,
) error {
	params := key.Params
	secret, err := key.GenerateShared(to, params.KeyLen, params.KeyLen)
	if err != nil {
		return err
	}
	dk, err := util.DeriveSymKey(secret, m.Salt)
	if err != nil {
		return err
	}
	seal, err := util.EncryptMessage(dk, m.GetContent())
	if err != nil {
		return err
	}
	m.Nonce = seal.Nonce
	m.Content = seal.Cipher
	m.From = util.PubKeyToNodeName(&ecdsaKey.PublicKey)
	return nil
}
