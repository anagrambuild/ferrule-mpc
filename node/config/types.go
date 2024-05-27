package config

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/ethclient"
)

type FerruleConfig struct {
	Identity           *ecdsa.PrivateKey
	NodeName           string
	PeerEncryptionKey  []byte
	Peers              []string
	DataDir            string
	CurrentSecretShare []byte
	CommitteeStateFile string
	ForkThreshold      int
	OpClient           *ethclient.Client
}
