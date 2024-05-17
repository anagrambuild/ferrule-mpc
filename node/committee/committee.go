package comittee

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"time"

	"github.com/anagrambuild/ferrule/cluster"
	"github.com/anagrambuild/ferrule/config"
	"github.com/anagrambuild/ferrule/schemas"
	"github.com/anagrambuild/ferrule/util"
	"github.com/bnb-chain/tss-lib/v2/ecdsa/keygen"
	"github.com/bnb-chain/tss-lib/v2/tss"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/hashicorp/memberlist"
	"github.com/oleiade/lane/v2"
	"go.uber.org/zap"
)

type comitteeState struct {
	Genesis     []byte                     `json:"genesis"`
	SecretShare *keygen.LocalPartySaveData `json:"secretShare"`
	EpochStart  *Epoch                     `json:"epochStart"`
	CurrentParties []*tss.PartyID           `json:"currentParties"`
}

type NodeLabel string

const (
	WantsKeygen NodeLabel = "WantsKeygen"
)

var LabelTypes []NodeLabel = []NodeLabel{
	WantsKeygen,
}

type Committee struct {
	cluster              *cluster.Cluster
	currentState         *comitteeState
	keygenSessionId      string
	keygenCurrentParties []*tss.PartyID
	keyGenInProgress     bool
	selfId               *tss.PartyID
	parties              *util.SyncedMap[string, *tss.PartyID]
	nodeLabels           *util.SyncedMap[NodeLabel, util.Set[string]]
	logger               *zap.Logger
	keyGenMsgQueue       map[uint32]*lane.Queue[*schemas.KeygenMessage]
	epochSource          EpochSource
	forkThreshold        int
	stateSaveLocation    string
	shutdown             chan struct{}
}

func NewCommittee(
	cfg *config.FerruleConfig,
	cluster *cluster.Cluster,
	epochSource EpochSource,
) *Committee {
	logger, _ := zap.NewProduction()
	var cs *comitteeState
	if cfg.CurrentSecretShare != nil && len(cfg.CurrentSecretShare) > 0 {
		err := json.Unmarshal(cfg.CurrentSecretShare, cs)
		if err == nil {
			logger.Info("Loaded current secret share")
		} else {
			logger.Error("Error loading current secret share", zap.Error(err))
		}
	}

	parties := util.NewSyncedMap[string, *tss.PartyID]()
	nodeLabels := util.NewSyncedMap[NodeLabel, util.Set[string]]()
	for _, label := range LabelTypes {
		nodeLabels.Set(label, util.NewSet[string]())
	}

	return &Committee{
		cluster:           cluster,
		currentState:      cs,
		logger:            logger,
		selfId:            tss.NewPartyID(cfg.Address, cfg.Address, util.PubkeyToBigInt(&cfg.Identity.PublicKey)),
		parties:           &parties,
		epochSource:       epochSource,
		nodeLabels:        &nodeLabels,
		forkThreshold:     max(3, cfg.ForkThreshold),
		stateSaveLocation: cfg.CommitteeStateFile,
		keyGenMsgQueue:    make(map[uint32]*lane.Queue[*schemas.KeygenMessage]),
		shutdown:          make(chan struct{}),
	}
}

func (c *Committee) Shutdown() {
	close(c.shutdown)
}

func (c *Committee) StartKeygen(
	sessionId string,
	parties []*tss.PartyID,
) {
	c.keyGenInProgress = true
	c.keygenSessionId = sessionId
	c.keygenCurrentParties = parties
	c.logger.Info("Starting keygen")
	preParams, _ := keygen.GeneratePreParams(1 * time.Minute)
	curve := tss.S256()
	sortedParties := tss.SortPartyIDs(parties)
	fmt.Println(sortedParties)
	pctx := tss.NewPeerContext(sortedParties)
	outCh := make(chan tss.Message)
	endCh := make(chan *keygen.LocalPartySaveData)
	errCh := make(chan error)
	parameters := tss.NewParameters(curve, pctx, c.selfId, len(sortedParties), c.forkThreshold)
	localParty := keygen.NewLocalParty(parameters, outCh, endCh, *preParams)

	go func() {
		err := localParty.Start()
		if err != nil {
			errCh <- err
		}
	}()

	roundPartySignoff := map[uint32]map[string]bool{}
	currentRound := uint32(1)
	processing := false

	for {
		select {
		case <-c.shutdown:
			{
				close(outCh)
				close(endCh)
				close(errCh)
				return
			}
		case err := <-errCh:
			{
				c.logger.Error("Error starting keygen", zap.Error(err))
				close(outCh)
				close(endCh)
				close(errCh)
				return
			}
		case msg := <-outCh:
			{
				c.logger.Info("Sending message", zap.String("msg", msg.String()))

			}
			break
		case <-time.After(time.Millisecond * 100):
			{
				if currentRound == 4 {
					break
				}

				if c.keyGenMsgQueue[currentRound] == nil || processing {
					continue
				}

				if c.keyGenMsgQueue[currentRound].Size() > 0 {
					kmsg, ok := c.keyGenMsgQueue[currentRound].Dequeue()
					pid := string(kmsg.From)
					if roundPartySignoff[currentRound] == nil {
						roundPartySignoff[currentRound] = make(map[string]bool)
					}
					roundPartySignoff[currentRound][pid] = true
					if ok {
						if c := c.parties.Get(pid); c != nil {
							processing = true
							ok, err := localParty.UpdateFromBytes(kmsg.Payload, c, kmsg.IsBroadcast)
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
				if c.keyGenMsgQueue[currentRound].Size() == 0 {
					moveToNextRound := true

					for _, party := range parties {
						if party.Id == c.selfId.Id {
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
			break
		case data := <-endCh:
			{
				c.logger.Info("Keygen complete")
				c.currentState.SecretShare = data
				//save current state to disk
				data, err := json.Marshal(c.currentState)
				if err != nil {
					c.logger.Error("Error saving current state", zap.Error(err))
					errCh <- err
				} else {
					err := os.WriteFile(c.stateSaveLocation, data, 0644)
					if err != nil {
						c.logger.Error("Error saving current state", zap.Error(err))
					}
				}

			}
			break
		}
	}
}

func nodeToPartyId(node *memberlist.Node) *tss.PartyID {
	pk, err := ethcrypto.UnmarshalPubkey([]byte(node.Name))
	if err != nil {
		return nil
	}
	return tss.NewPartyID(node.Name, node.Name, util.PubkeyToBigInt(pk))
}

func (c *Committee) syncParties() {
	nodes := c.cluster.GetNodes()
	newMap := make(map[string]*tss.PartyID)
	for _, node := range nodes {
		if node.Name != c.cluster.SelfNode().Name {
			newMap[node.Name] = nodeToPartyId(node)
		}
	}
	c.parties.Replace(newMap)
}

func (c *Committee) shouldProposeKeygen() bool {
	if c.currentState == nil ||
		c.currentState.EpochStart == nil ||
		c.epochSource.Expired(c.currentState.EpochStart) ||
		c.currentState.SecretShare == nil {
		return true
	}
	return false
}

func sessionId(parties []*tss.PartyID) string {
	var combine string
	spid := tss.SortPartyIDs(parties)
	for _, party := range spid {
		combine += party.Id
	}
	sha := ethcrypto.Keccak256([]byte(combine))
	return string(sha)
}

func stateHandler(c *Committee) {
	for {
		select {
		case <-time.After(5 * time.Second):
			{
				c.syncParties()
				if c.shouldProposeKeygen() {
					c.logger.Info("Proposing keygen")
					c.nodeLabels.Get(WantsKeygen).Add(c.cluster.SelfNode().Name)
					c.cluster.Broadcast(&schemas.ClusterMessage{
						ParsedContent: &schemas.Content{
							Content: &schemas.Content_KeygenStart{
								KeygenStart: &schemas.KeygenStart{
									SessionID: c.keygenSessionId,
								},
							},
						},
					})
				}
			}
		}
	}
}

func (c *Committee) Start() {
	c.cluster.AddEventListener(
		&cluster.ClusterEventListener{
			Id: "comittee-event",
			Listener: func(event memberlist.NodeEvent) {
				switch event.Event {
				case memberlist.NodeJoin:
					{
						c.logger.Info("Node joined", zap.String("node", event.Node.Name))
						c.parties.Set(event.Node.Name, nodeToPartyId(event.Node))
					}
				case memberlist.NodeLeave:
					{
						c.logger.Info("Node left", zap.String("node", event.Node.Name))
						c.parties.Delete(event.Node.Name)
					}
				}
			},
		})

	c.cluster.AddListener(&cluster.ClusterMessageListener{
		Id: "comittee",
		Predicate: func(message *schemas.ClusterMessage) bool {
			types := []string{
				"KeygenMessage",
				"KeygenStart",
			}
			return slices.Contains(types, string(message.ParsedContent.ProtoReflect().Descriptor().Name()))
		},
		Listener: func(message *schemas.ClusterMessage) error {
			switch message.ParsedContent.ProtoReflect().Descriptor().Name() {
			case "KeygenStart":
				{
					c.nodeLabels.Get(WantsKeygen).Add(string(message.From))
					if c.currentState == nil || c.epochSource.Expired(c.currentState.EpochStart) {
						if proposingNodes := c.nodeLabels.Get(WantsKeygen); proposingNodes != nil && proposingNodes.Len() >= c.cluster.NumNodes()/c.forkThreshold {

							parties := []*tss.PartyID{}
							for _,node := range c.nodeLabels.Get(WantsKeygen).Items() {
								if party := c.parties.Get(node); party != nil {
									parties = append(parties, party)
								}
							}

							go c.StartKeygen(
								sessionId(parties),
								parties,
							)
						}
					} else {
						c.logger.Info("Keygen proposal recived, rejected but stored for later")
						c.nodeLabels.Get(WantsKeygen).Add(string(message.From))
					}
				}
				break
			case "KeygenMessage":
				{
					msg := message.GetParsedContent().GetKeygen()
					if msg == nil {
						c.logger.Error("Invalid keygen message")
						break
					}
					c.logger.Info("Keygen message recived")
					if pa := c.parties.Get(string(message.From)); pa != nil {
						c.logger.Info("Keygen message recived from committee member", zap.String("from", string(message.From)), zap.String("round", message.String()))
						if c.keyGenMsgQueue[msg.Round] == nil {
							q := lane.NewQueue[*schemas.KeygenMessage]()
							c.keyGenMsgQueue[msg.Round] = q
						}
						c.keyGenMsgQueue[msg.Round].Enqueue(msg)
					} else {
						fmt.Println("not in committee")
					}
				}
				break
			}
			return nil
		},
	})
	// Start the committee
}
