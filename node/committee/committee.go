package committee

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/anagrambuild/ferrule/cluster"
	"github.com/anagrambuild/ferrule/config"
	"github.com/anagrambuild/ferrule/contracts"
	"github.com/anagrambuild/ferrule/schemas"
	"github.com/anagrambuild/ferrule/util"
	"github.com/bnb-chain/tss-lib/v2/ecdsa/keygen"
	"github.com/bnb-chain/tss-lib/v2/tss"
	"github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hashicorp/memberlist"
	"github.com/oleiade/lane/v2"
	"github.com/rs/zerolog/log"
	"go.uber.org/zap"
)

type comitteeState struct {
	Genesis        []byte                     `json:"genesis"`
	SecretShare    *keygen.LocalPartySaveData `json:"secretShare"`
	EpochStart     *Epoch                     `json:"epochStart"`
	CurrentParties []*tss.PartyID             `json:"currentParties"`
	SharedPubKey   *ecdsa.PublicKey           `json:"sharedPubKey"`
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
	parties              util.SyncedMap[string, *tss.PartyID]
	nodeLabels           util.SyncedMap[NodeLabel, util.Set[string]]
	keyGenMsgQueue       map[uint32]*lane.Queue[*schemas.KeygenMessage]
	epochSource          EpochSource
	forkThreshold        int
	stateSaveLocation    string
	shutdown             chan struct{}
	ethConnection        *ethclient.Client
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
	c := &Committee{
		cluster:           cluster,
		currentState:      cs,
		selfId:            tss.NewPartyID(cfg.NodeName, cfg.NodeName, util.PubkeyToBigInt(&cfg.Identity.PublicKey)),
		parties:           util.NewSyncedMap[string, *tss.PartyID](),
		epochSource:       epochSource,
		nodeLabels:        util.NewSyncedMap[NodeLabel, util.Set[string]](),
		forkThreshold:     max(2, cfg.ForkThreshold),
		stateSaveLocation: cfg.CommitteeStateFile,
		keyGenMsgQueue:    make(map[uint32]*lane.Queue[*schemas.KeygenMessage]),
		shutdown:          make(chan struct{}),
		ethConnection:     cfg.OpClient,
	}
	for _, label := range LabelTypes {
		c.nodeLabels.Set(label, util.NewSet[string]())
	}
	return c
}

func (c *Committee) Shutdown() {
	close(c.shutdown)
}

func (c *Committee) ReadFid() (*big.Int, error) {
	addr := common.HexToAddress(contracts.ID_REGISTRY_ADDRESS)
	caller, err := contracts.NewIdRegistryCaller(addr, c.ethConnection)
	if err != nil {
		return nil, err
	}
	selfAddr := ethcrypto.PubkeyToAddress(*c.currentState.SharedPubKey)
	fid, err := caller.IdOf(nil, selfAddr)
	if err != nil {
		return nil, err
	}
	return fid, nil
}

func (c *Committee) EndKeygen(
	errch chan error,
	outch chan tss.Message,
	endch chan *keygen.LocalPartySaveData,
) {
	c.keyGenInProgress = false
	c.keygenSessionId = ""
	c.keygenCurrentParties = nil
	close(errch)
	close(outch)
	close(endch)
}

func (c *Committee) StartKeygen(
	sessionId string,
	parties []*tss.PartyID,
) {
	c.keyGenInProgress = true
	c.keygenSessionId = sessionId
	c.keygenCurrentParties = parties
	log.Debug().Msg("Starting keygen")
	preParams, err := keygen.GeneratePreParams(1 * time.Minute)
	if err != nil {
		log.Error().Err(err).Msg("Error generating pre-params")
		c.keyGenInProgress = false
		c.keygenSessionId = ""
		c.keygenCurrentParties = nil
		return
	}
	curve := tss.S256()
	sortedParties := tss.SortPartyIDs(parties)
	pctx := tss.NewPeerContext(sortedParties)
	outCh := make(chan tss.Message)
	endCh := make(chan *keygen.LocalPartySaveData)
	errCh := make(chan error)
	t := max(2, len(sortedParties)/3)
	parameters := tss.NewParameters(curve, pctx, c.selfId, len(sortedParties), t)
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
				log.Error().Err(err).Msg("Error starting keygen")
				close(outCh)
				close(endCh)
				close(errCh)
				return
			}
		case msg := <-outCh:
			{
				if msg != nil {
					bytes, _, err := msg.WireBytes()
					if err != nil {
						log.Error().Err(err).Msg("Error getting wire bytes")
						break
					}

					if msg.IsBroadcast() {
						log.Debug().Any("parties", c.keygenCurrentParties).Msg("Sending keygen broadcast message")
						//to big for gossip
						for _, to := range c.keygenCurrentParties {
							wireMessage := schemas.ClusterMessage{
								ParsedContent: &schemas.Content{
									Content: &schemas.Content_Keygen{
										Keygen: &schemas.KeygenMessage{
											From:        []byte(c.selfId.Id),
											IsBroadcast: false,
											SessionID:   sessionId,
											Payload:     bytes,
											Round:       currentRound,
										},
									},
								},
							}
							if toNode := c.cluster.LookupNode(to.Id); toNode != nil {
								log.Debug().Msg("Sending keygen broadcast message")
								err := c.cluster.SendToNode(toNode, &wireMessage)
								if err != nil {

									log.Error().Err(err).Msg("Error sending keygen message")
								}
							} else {
								log.Error().Msg("Node not found, or has left since keygen started")
								//end keygen
								if sortedParties.Len()-1 < t {
									c.EndKeygen(errCh, outCh, endCh)
									return
								}
							}
						}
					} else {
						log.Debug().Any("parties", msg.GetTo()).Msg("Sending keygen message")
						for _, to := range msg.GetTo() {
							wireMessage := schemas.ClusterMessage{
								ParsedContent: &schemas.Content{
									Content: &schemas.Content_Keygen{
										Keygen: &schemas.KeygenMessage{
											From:        []byte(c.selfId.Id),
											IsBroadcast: false,
											SessionID:   sessionId,
											Payload:     bytes,
											Round:       currentRound,
										},
									},
								},
							}
							if toNode := c.cluster.LookupNode(to.Id); toNode != nil {
								log.Debug().Msg("Sending keygen message")
								err := c.cluster.SendToNodeEncrypted(toNode, &wireMessage)
								if err != nil {
									log.Error().Err(err).Msg("Error sending keygen message")
								}
							} else {
								log.Error().Msg("Node not found, or has left since keygen started")
								//end keygen
								if sortedParties.Len()-1 < t {
									c.EndKeygen(errCh, outCh, endCh)
									return
								}
							}

						}
					}

				}
			}
			break
		case <-time.After(time.Second * 10):
			{
				log.Debug().Any("parties ", localParty.WaitingFor()).Msg("Keygen waiting")

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
								log.Error().Err(err).Msg("Error updating keygen")
							}
							if !ok {
								log.Error().Msg("Keygen update not okay")
							}
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
				log.Info().Msg("Keygen completed")
				c.currentState.SecretShare = data
				c.keyGenInProgress = false
				c.keygenSessionId = ""
				c.keygenCurrentParties = nil
				//save current state to disk
				data, err := json.Marshal(c.currentState)
				if err != nil {
					log.Error().Err(err).Msg("Error marshalling save data")
					errCh <- err
				} else {
					err := os.WriteFile(c.stateSaveLocation, data, 0644)
					if err != nil {
						log.Error().Err(err).Msg("Error saving current state")
						errCh <- err
					}
				}

			}
			break
		}
	}
}

func nodeToPartyId(node *memberlist.Node) *tss.PartyID {
	pk, err := util.NodeNameToPubKey(node.Name)
	if err != nil {
		log.Error().Err(err).Msg("Error getting pubkey from node")
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
	newMap[c.cluster.SelfNode().Name] = c.selfId
	c.parties.Replace(newMap)
}

func (c *Committee) shouldProposeKeygen() bool {
	//add in epoch expiry
	if c.keyGenInProgress == false && c.parties.Len() >= c.forkThreshold && c.currentState == nil {
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
	sha := ethcrypto.Keccak256Hash([]byte(combine))
	return sha.String()
}

func stateHandler(c *Committee) {
	for {
		select {
		case <-c.shutdown:
			{
				return
			}
		case <-time.After(5 * time.Second):
			{
				c.syncParties()
				if c.shouldProposeKeygen() {
					log.Debug().Msg("Proposing keygen")
					c.nodeLabels.Get(WantsKeygen).Add(c.cluster.SelfNode().Name)
					err := c.cluster.Broadcast(
						&schemas.ClusterMessage{
							ParsedContent: &schemas.Content{
								Content: &schemas.Content_KeygenStart{
									KeygenStart: &schemas.KeygenStart{
										SessionID: c.keygenSessionId,
									},
								},
							},
						})
					if err != nil {
						log.Error().Err(err).Msg("Error broadcasting keygen proposal")
					} else {
						log.Info().Msg("Keygen proposal broadcasted")
					}
				}
			}
		}
	}
}

func (c *Committee) Start() {
	log.Info().Msg("Starting committee")
	c.cluster.AddEventListener(
		&cluster.ClusterEventListener{
			Id: "comittee-event",
			Listener: func(event memberlist.NodeEvent) {
				switch event.Event {
				case memberlist.NodeJoin:
					{
						c.parties.Set(event.Node.Name, nodeToPartyId(event.Node))
					}
				case memberlist.NodeLeave:
					{
						c.parties.Delete(event.Node.Name)
					}
				}
			},
		})
	c.cluster.AddListener(&cluster.ClusterMessageListener{
		Id: "comittee",
		Predicate: func(message *schemas.ClusterMessage) bool {
			if message.ParsedContent == nil || message.ParsedContent.Content == nil {
				log.Error().Msg("Invalid message, nil content")
				return false
			}
			switch message.ParsedContent.Content.(type) {
			case *schemas.Content_KeygenStart:
				return true
			case *schemas.Content_Keygen:
				return true
			}
			return false
		},
		Listener: func(message *schemas.ClusterMessage) error {
			switch message.ParsedContent.Content.(type) {
			case *schemas.Content_KeygenStart:
				{
					if !c.keyGenInProgress {
						log.Info().Msg("Keygen proposal recived")
						c.nodeLabels.Get(WantsKeygen).Add(message.From)
						// move to same logic as should propose keygen
						if c.currentState == nil || c.epochSource.Expired(c.currentState.EpochStart) {
							if proposingNodes := c.nodeLabels.Get(WantsKeygen); proposingNodes != nil && proposingNodes.Len() >= c.forkThreshold {
								log.Info().Any("nodes", proposingNodes.Items()).Msg("Keygen proposal accepted")
								parties := []*tss.PartyID{}
								for _, node := range proposingNodes.Items() {
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
							log.Debug().Msg("Keygen proposal rejected, but saved")
							c.nodeLabels.Get(WantsKeygen).Add(string(message.From))
						}
					}
				}
				break
			case *schemas.Content_Keygen:
				{
					log.Info().Msg("Keygen message recived")
					if c.keyGenInProgress {
						log.Debug().Str("from", string(message.From)).Msg("Keygen message recived")
						msg := message.GetParsedContent().GetKeygen()
						if msg == nil {
							log.Error().Msg("Invalid keygen message")
							break
						}
						if pa := c.parties.Get(string(message.From)); pa != nil {
							log.Info().Str("from", string(message.From)).Msg("Keygen message recived")

							if c.keyGenMsgQueue[msg.Round] == nil {
								q := lane.NewQueue[*schemas.KeygenMessage]()
								c.keyGenMsgQueue[msg.Round] = q
							}
							c.keyGenMsgQueue[msg.Round].Enqueue(msg)
						} else {
							log.Error().Str("from", string(message.From)).Msg("Proposed keygen from non-comittee party")
						}
					}
				}
				break
			}
			return nil
		},
	})
	// Start the committee
	go stateHandler(c)
}
