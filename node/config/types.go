package config

import (
	"crypto/ecdsa"
)

type FerruleConfig struct {
	Identity           *ecdsa.PrivateKey
	Address            string
	PeerEncryptionKey  []byte
	Peers              []string
	DataDir            string
	CurrentSecretShare []byte
	CommitteeStateFile string
	ForkThreshold      int
}
