package schemas

import (
	"crypto/ecdh"
	"crypto/ecdsa"

	"github.com/anagrambuild/ferrule/util"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
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
	return nil
}

func (m *ClusterMessage) GetFromKey() (*ecdsa.PublicKey, error) {
	return ethcrypto.UnmarshalPubkey(m.From)
}

func (m *ClusterMessage) GetFromEcdh() (*ecdh.PublicKey, error) {
	key, err := m.GetFromKey()
	if err != nil {
		return nil, err
	}
	return key.ECDH()
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
	if err != nil {
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
	key *ecdh.PrivateKey,
) error {
	pubEc, err := m.GetFromEcdh()
	if err != nil {
		return err
	}
	secret, err := key.ECDH(pubEc)
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
