package main

import (
	"os"
	"os/signal"
	"time"

	"github.com/anagrambuild/ferrule/cluster"
	"github.com/anagrambuild/ferrule/committee"
	"github.com/anagrambuild/ferrule/util"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/rs/zerolog/log"
)

/*
this system works of epochs where new members who
joined or left in the epoch will be accounted for in the keygen ceremony
if no new members joined or left in the epoch, the keygen will not be done.

If there is a need for a new keygen, the members will propose a new key,
while holding on to the old key shares. After the new key is decided on then.
The proposer will get a signature from all the parties to run the transfer method on the id registry contract.
If the fid does not exist this will be the first keygen, if it does exist, the keygen will form a transfer.

All operations before the settlement of the transfer will be done with the old key shares.
After the transfer is settled, the old key shares will be discarded and the new key shares will be used.

*/

func main() {
	// setup configuration
	cfg := util.InitConfig()
	cluster := cluster.NewCluster(cfg)
	epochSource := committee.NewOpBlockEpochSource("")

	err := cluster.Start()
	if err != nil {
		panic(err)
	}
	comittee := committee.NewCommittee(cfg, cluster, epochSource)
	comittee.Start()
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Kill)

	for {
		select {
		case <-time.After(10 * time.Second):
			{
				p := comittee.SharedPubKey()
				if p == nil {
					continue
				}
				log.Info().Interface("pubkey", ethcrypto.PubkeyToAddress(*p)).Msg("Current public key")
				break
			}
		case <-sigChan:
			comittee.Shutdown()
			cluster.Shutdown()
			os.Exit(0)
		}
	}
}
