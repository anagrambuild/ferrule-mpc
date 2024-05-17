package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"
	"regexp"
	"strconv"
	"sync"
	"time"

	"github.com/anagrambuild/ferrule/schemas"
	"github.com/bnb-chain/tss-lib/v2/ecdsa/keygen"
	"github.com/bnb-chain/tss-lib/v2/tss"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/hashicorp/memberlist"
	"github.com/kubeshark/goque"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"google.golang.org/protobuf/proto"

	"github.com/oleiade/lane/v2"
	"github.com/vechain/go-ecvrf"
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

	//setup self state from local storage and config

	//setup channels

	//setup cluster

	//setup cluster watcher
}



/*






type ConfigRaw struct {
	IdentityKpPath string   `mapstructure:"IDENTITY_KEY_PATH"`
	SecretShareDir string   `mapstructure:"SECRET_SHARE_DIR"`
	Peers          []string `mapstructure:"PEERS"`
	LogLevel       string   `mapstructure:"LOG_LEVEL"`
	DataDir        string   `mapstructure:"DATA_DIR"`
}

type ConfigLoaded struct {
	Identity          *ecdsa.PrivateKey
	Address           string
	PeerEncryptionKey []byte
	Peers             []string
	Logger            *zap.Logger
	DataDir           string
}

var roundRegex = regexp.MustCompile(`.+Round(\d{1}).+`)

func GetRound(messageType string) int32 {
	match := roundRegex.FindStringSubmatch(messageType)
	if len(match) < 2 {
		return -1
	}
	round, err := strconv.Atoi(match[1])
	if err != nil {
		return -1
	}
	return int32(round)
}

func loadIdentity(path string) *ecdsa.PrivateKey {
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



func Retry(attempts int, sleep time.Duration, fn func() error) error {
	var err error
	for i := 0; i < attempts; i++ {
		err = fn()
		if err == nil {
			return nil
		}
		time.Sleep(sleep)
		fmt.Printf("Attempt %d failed: %s\n", i, err)
		fmt.Printf("Retrying in %s\n", sleep)
		fmt.Println()
	}
	return err
}

func FormFerrule(
	config ConfigLoaded,
	delegate memberlist.Delegate,
) (*memberlist.Memberlist, error) {
	c := memberlist.DefaultLocalConfig()
	c.Name = config.Address
	fmt.Println(config.Address)
	c.Delegate = delegate
	m, err := memberlist.Create(c)
	if err != nil {
		panic("Failed to create memberlist: " + err.Error())
	}
	if len(config.Peers) > 0 {
		var peersConnected int
		rerr := Retry(
			10,
			time.Second,
			func() error {
				pc, err := m.Join(config.Peers)
				peersConnected = pc
				return err
			})

		if rerr != nil {
			return nil, err
		}
		config.Logger.Info("Peers connected", zap.Any("peers", peersConnected))
	}
	return m, nil
}

func MembersWorker(
	config *ConfigLoaded,
	m *memberlist.Memberlist,

) error {
	if config.Logger.Level() == zap.DebugLevel {
		for range time.Tick(time.Second * 10) {
			config.Logger.Info("Members", zap.Any("members", m.Members()))
		}
	}
	return nil
}

const WOKER_POOL_SIZE = 10

func ClusterActivityWorker(
	logger *zap.Logger,
	config *ConfigLoaded,
	m *memberlist.Memberlist,
	queue *QueueManager,
	msgCh chan []byte,
) error {
	for {
		select {
		case <-time.After(time.Millisecond * 10):
			{
				if queue.NumMessages() > 0 {
					msg, err := queue.GetMessage()
					if err != nil {
						logger.Error("Failed to get message from queue", zap.Error(err))
						continue
					}
					fmt.Println("got message")
					msgCh <- msg.Value
				}
			}
		}
	}
}

type KeygenUpdate struct {
	Round       int32
	From        string
	IsBroadcast bool
	Payload     []byte
}

//TODO come up with other network level decisions on when to propose keygen

func ProposeKeygen(
	config *ConfigLoaded,
	ctx context.Context,
	cl *Cluster,
	updateCh <-chan []byte,
	outCh chan<- tss.Message,
	endCh chan<- *keygen.LocalPartySaveData,
) error {
	errch := make(chan error, 1)
	var lp tss.Party
	self := cl.GetSelf()
	for cl.Memberlist.NumMembers() < 2 {
		select {
		case <-time.After(time.Second * 1):
			{
				config.Logger.Info("Waiting for enough members to propose keygen", zap.Any("members", cl.Memberlist.NumMembers()))
			}
		}
	}
	config.Logger.Info("Proposing keygen with enough members", zap.Any("members", cl.Memberlist.NumMembers()))
	preParams, _ := keygen.GeneratePreParams(1 * time.Minute)
	curve := tss.S256()
	parties := cl.Parties()
	fmt.Println(parties)
	pctx := tss.NewPeerContext(parties)

	parameters := tss.NewParameters(curve, pctx, self, len(parties), 1)
	party := keygen.NewLocalParty(parameters, outCh, endCh, *preParams)
	lp = party

	go func() {
		err := lp.Start()
		if err != nil {
			errch <- err
		}
	}()
	messageQueues := map[int32]*lane.Queue[KeygenUpdate]{}
	roundPartySignoff := map[int32]map[string]bool{}
	currentRound := int32(1)
	processing := false

	go func() {
		for {
			select {
			case <-time.After(time.Second):
				{
					if currentRound == 4 {
						break
					}
					fmt.Println("msgqueus", messageQueues)
					fmt.Println("round", currentRound)
					if messageQueues[currentRound] == nil || processing {
						continue
					}

					if messageQueues[currentRound].Size() > 0 {
						fmt.Println("waiting", lp.WaitingFor())
						fmt.Println("running", lp.Running())
						i, ok := messageQueues[currentRound].Dequeue()
						if roundPartySignoff[currentRound] == nil {
							roundPartySignoff[currentRound] = make(map[string]bool)
						}
						roundPartySignoff[currentRound][i.From] = true
						if ok {
							if c := cl.LookupParty(i.From); c != nil {
								processing = true
								ok, err := lp.UpdateFromBytes(i.Payload, c, i.IsBroadcast)
								if err != nil {
									fmt.Println("failed to update keygen", err)
								}
								if !ok {
									fmt.Println("keygen update not okay")
								}
								fmt.Println("updated keygen in round", currentRound)
								processing = false
							}
						}
					}
					if messageQueues[currentRound].Size() == 0 {
						moveToNextRound := true

						for _, party := range parties {
							if party.Id == self.Id {
								continue
							}
							if k, ok := roundPartySignoff[currentRound][party.Id]; !k || !ok {
								moveToNextRound = false
								break
							}
						}
						fmt.Println("signoff", roundPartySignoff)
						if moveToNextRound {
							currentRound += 1
						}
					}
				}
			}
		}
	}()

	for lp != nil {
		select {
		case update := <-updateCh:
			{
				fmt.Println("update")
				msg := schemas.KeygenMessage{}
				proto.Unmarshal(update, &msg)
				hashedPayload := ethcrypto.Keccak256(
					[]byte(msg.From),
					[]byte(msg.To),
					msg.Payload,
				)
				//todo replay attacks here, so we may need a clock or something
				pubkey, err := ethcrypto.SigToPub(hashedPayload, msg.Signature)
				if err != nil || pubkey == nil || ethcrypto.PubkeyToAddress(*pubkey).String() != msg.From {
					fmt.Println("invalid signature")
					return fmt.Errorf("Invalid signature")
				}
				if c := cl.LookupParty(msg.From); c != nil {
					fmt.Println("lookup", msg.Round)
					if messageQueues[msg.Round] == nil {
						messageQueues[msg.Round] = lane.NewQueue[KeygenUpdate]()
					}
					messageQueues[msg.Round].Enqueue(KeygenUpdate{
						Round:       msg.Round,
						From:        msg.From,
						IsBroadcast: msg.IsBroadcast,
						Payload:     msg.Payload,
					})
				} else {
					fmt.Println("not in committee")
				}
			}
		case err := <-errch:
			{
				fmt.Println("error", err)
				return err
			}
		}
	}
	return nil
}



func main() {
	ctx := context.Background()
	wg, _ := errgroup.WithContext(ctx)
	config := InitConfig()
	queue, err := NewQueueManager(config.DataDir)

	if err != nil {
		panic(err)
	}
	events := NewDelegate(
		&config.Identity.PublicKey,
		config.Logger,
		queue,
	)

	m, err := FormFerrule(*config, events)
	if err != nil {
		panic(err)
	}
	msgCh := make(chan []byte)
	wg.Go(func() error {
		kgDone := make(chan *keygen.LocalPartySaveData)
		loggingCh := make(chan tss.Message)
		selfParty := tss.NewPartyID(
			m.LocalNode().Name,
			m.LocalNode().Name,
			PubkeyToBigInt(&config.Identity.PublicKey),
		)
		cluster := NewCluster(selfParty, m)
		updateCh := make(chan []byte)
		go ProposeKeygen(
			config,
			context.Background(),
			cluster,
			updateCh,
			loggingCh,
			kgDone,
		)
		self := cluster.GetSelf()
		for {
			select {
			case m := <-msgCh:
				{

					updateCh <- m
				}
			case d := <-kgDone:
				{
					config.Logger.Error("Keygen done", zap.Any("data", d))
					config.Logger.Sync()
				}
			case d := <-loggingCh:
				{
					//todo track committe
					to := d.GetTo()
					if d.IsBroadcast() {
						to = cluster.Parties()
					}
					byts, r, err := d.WireBytes()
					fmt.Println("routing", r)
					if err != nil {
						config.Logger.Error("Failed to get wire bytes", zap.Error(err))
						continue
					}
					for _, party := range to {
						if party.Id == cluster.GetSelf().Id {
							continue
						}
						if c := cluster.LookupNode(party.Id); c != nil {
							payloadHash := ethcrypto.Keccak256(
								[]byte(self.Id),
								[]byte(party.Id),
								byts,
							)
							signature, err := ethcrypto.Sign(payloadHash, config.Identity)
							if err != nil {
								config.Logger.Error("Failed to sign payload", zap.Error(err))
								continue
							}
							msg := schemas.KeygenMessage{
								Round:       GetRound(d.Type()),
								From:        self.Id,
								To:          party.Id,
								Payload:     byts,
								Signature:   signature,
								IsBroadcast: d.IsBroadcast(),
							}
							finalBytes, err := proto.Marshal(&msg)
							if err != nil {
								config.Logger.Error("Failed to marshal message", zap.Error(err))
								continue
							}
							err = m.SendReliable(c, finalBytes)
							if err != nil {
								config.Logger.Error("Failed to send message", zap.Error(err))
								continue
							}
							fmt.Println("Sent ", d.Type())
						}
					}
				}
			}
		}

	})

	wg.Go(func() error {
		return MembersWorker(config, m)
	})
	wg.Go(func() error {
		return ClusterActivityWorker(config.Logger, config, m, queue, msgCh)
	})
	wg.Wait()
	config.Logger.Info("exiting")
}

type DQueueDelegate struct {
	publicKey    *ecdsa.PublicKey
	queueManager *QueueManager
	logger       *zap.Logger
}

func NewDelegate(
	pub *ecdsa.PublicKey,
	logger *zap.Logger,
	queueManager *QueueManager,
) *DQueueDelegate {
	return &DQueueDelegate{
		publicKey:    pub,
		queueManager: queueManager,
		logger:       logger,
	}
}

func (d *DQueueDelegate) NodeMeta(limit int) []byte {
	// Metadata about this node, return nil or a small byte slice
	return ethcrypto.FromECDSAPub(d.publicKey)
}

func (d *DQueueDelegate) NotifyMsg(b []byte) {
	// Enqueue received message into the queue
	if err := d.queueManager.RecvMessage(b); err != nil {
		d.logger.Info("NotifyMsg", zap.Error(err))
	}

}

func (d *DQueueDelegate) GetBroadcasts(overhead, limit int) [][]byte {
	// Return messages to be broadcasted to other nodes, nil for now
	return nil
}

func (d *DQueueDelegate) LocalState(join bool) []byte {

	// Return local state data, return nil for simplicity
	return nil
}

func (d *DQueueDelegate) MergeRemoteState(buf []byte, join bool) {
	// Merge received state from another node, no action for now
	d.logger.Info("MergeRemoteState", zap.Any("buf", buf), zap.Bool("join", join))
}

const WorkerPoolSize = 5

var MSG_PREFIX = []byte("msg")
var BROADCAST_PREFIX = []byte("bct")

// QueueManager manages queue operations
type QueueManager struct {
	mainQ *goque.PrefixQueue
}

// NewQueueManager initializes a new QueueManager
func NewQueueManager(path string) (*QueueManager, error) {
	q, err := goque.OpenPrefixQueue(path)
	if err != nil {
		return nil, err
	}
	return &QueueManager{
		mainQ: q,
	}, nil
}

// EnqueueMessage adds a message to the queue
func (qm *QueueManager) RecvMessage(data []byte) error {
	_, err := qm.mainQ.Enqueue(MSG_PREFIX, data)
	return err
}

func (qm *QueueManager) NumMessages() uint64 {
	return qm.mainQ.Length()
}

// func (qm *QueueManager) RecvBroadcast(data []byte) error {
// 	_, err := qm.mainQ.Enqueue(BROADCAST_PREFIX, data)
// 	return err
// }

func (qm *QueueManager) GetMessage() (*goque.Item, error) {
	return qm.mainQ.Dequeue(MSG_PREFIX)
}

// func (qm *QueueManager) GetBroadcast() ([]byte, error) {
// 	return qm.mainQ.Dequeue(BROADCAST_PREFIX)
// }

*/