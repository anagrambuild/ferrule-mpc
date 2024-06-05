package committee

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"math/big"
	"os"
	"path"
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
)

type comitteeState struct {
	Genesis        []byte                     `json:"genesis,omitempty"`
	SecretShare    *keygen.LocalPartySaveData `json:"secretShare,omitempty"`
	EpochStart     *Epoch                     `json:"epochStart,omitempty"`
	CurrentParties []*tss.PartyID             `json:"currentParties,omitempty"`
}

type keygenUpdate struct {
	From    *tss.PartyID
	To      *tss.PartyID
	Message *schemas.KeygenMessage
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
	currentCommittee     util.SyncedMap[string, *tss.PartyID]
	nodeLabels           util.SyncedMap[NodeLabel, util.Set[string]]
	keyGenMsgQueue       util.SyncedMap[uint32, *lane.Queue[*keygenUpdate]]
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

	var cs comitteeState
	if cfg.CurrentSecretShare != nil && len(cfg.CurrentSecretShare) > 0 {
		err := json.Unmarshal(cfg.CurrentSecretShare, &cs)
		if err == nil {
			log.Info().Msg("Loaded current secret share")
		} else {
			log.Error().Err(err).Msg("Failed to load current secret share")
		}
	}
	c := &Committee{
		cluster:           cluster,
		currentState:      &cs,
		selfId:            tss.NewPartyID(cfg.NodeName, cfg.NodeName, util.PubkeyToBigInt(&cfg.Identity.PublicKey)),
		parties:           util.NewSyncedMap[string, *tss.PartyID](),
		epochSource:       epochSource,
		nodeLabels:        util.NewSyncedMap[NodeLabel, util.Set[string]](),
		forkThreshold:     max(2, cfg.ForkThreshold),
		stateSaveLocation: cfg.CommitteeStateFile,
		keyGenMsgQueue:    util.NewSyncedMap[uint32, *lane.Queue[*keygenUpdate]](),
		shutdown:          make(chan struct{}),
		ethConnection:     cfg.OpClient,
	}
	for _, label := range LabelTypes {
		c.nodeLabels.Set(label, util.NewSet[string]())
	}
	return c
}

func (c *Committee) SharedPubKey() *ecdsa.PublicKey {
	if c.currentState == nil || c.currentState.SecretShare == nil {
		return nil
	}
	return c.currentState.SecretShare.ECDSAPub.ToECDSAPubKey()
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
	selfAddr := ethcrypto.PubkeyToAddress(*c.SharedPubKey())
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
	ctx context.Context,
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
			log.Error().Err(err).Msg("Error starting keygen local party")
			errCh <- err
		}
	}()

	go func() {
		roundPartySignoff := map[uint32]map[string]bool{}
		currentRound := uint32(1)
		processing := false
		for {
			select {
			case <-ctx.Done():
			case <-c.shutdown:
			case <-errCh:
				return
			case <-time.After(time.Second * 5):
				{
					if currentRound == 4 {
						break
					}
					cq := c.keyGenMsgQueue.Get(currentRound)
					if cq == nil || processing {
						break
					}

					if cq.Size() > 0 {
						log.Debug().Msg("Starting keygen message processing")

						kmsg, ok := cq.Dequeue()
						log.Debug().Uint32("round", currentRound).Uint32("message round", kmsg.Message.Round).Msg("Processing keygen message")
						if roundPartySignoff[currentRound] == nil {
							roundPartySignoff[currentRound] = make(map[string]bool)
						}
						roundPartySignoff[currentRound][kmsg.From.Id] = true
						if ok {
							processing = true
							ok, err := localParty.UpdateFromBytes(kmsg.Message.Payload, kmsg.From, kmsg.Message.IsBroadcast)
							if err != nil {
								log.Error().Err(err).Msg("Error updating keygen")
							}

							if !ok {
								log.Error().Msg("Keygen update not okay")
							}
							log.Debug().Msg("Updated keygen")
							processing = false
						}
					}
					if cq.Size() == 0 {
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
						if moveToNextRound {
							currentRound += 1
						}
					}

				}
				break
			}
		}
	}()

	for {
		select {
		case <-ctx.Done():
			{
				c.EndKeygen(errCh, outCh, endCh)
				return
			}
		case <-c.shutdown:
			{
				c.EndKeygen(errCh, outCh, endCh)
				return
			}
		case err := <-errCh:
			{
				log.Error().Err(err).Msg("Error in keygen")
				c.EndKeygen(errCh, outCh, endCh)
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
					toParties := msg.GetTo()
					if toParties == nil {
						toParties = c.keygenCurrentParties
					}
					for _, to := range toParties {
						wireMessage := schemas.ClusterMessage{
							ParsedContent: &schemas.Content{
								Content: &schemas.Content_Keygen{
									Keygen: &schemas.KeygenMessage{
										IsBroadcast: msg.IsBroadcast(),
										SessionID:   sessionId,
										Payload:     bytes,
										Round:       util.GetRound(msg.Type()),
									},
								},
							},
						}
						if toNode := c.cluster.LookupNode(to.Id); toNode != nil {
							log.Debug().Msg("Sending keygen message")
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
				}
			}
			break
		case data := <-endCh:
			{

				log.Info().Msg("Keygen completed")
				c.currentState = &comitteeState{
					CurrentParties: c.keygenCurrentParties,
					SecretShare:    data,
				}
				//save current state to disk
				data, err := json.Marshal(c.currentState)
				if err != nil {
					log.Error().Err(err).Msg("Error marshalling save data")
					errCh <- err
				} else {
					err := os.MkdirAll(path.Dir(c.stateSaveLocation), 0755)
					if err != nil {

						log.Error().Err(err).Msg("Error creating directory for state save")
					}

					err = os.WriteFile(c.stateSaveLocation, data, 0644)
					if err != nil {
						log.Error().Err(err).Msg("Error saving current state")
						errCh <- err
					}
				}
				c.EndKeygen(errCh, outCh, endCh)
				return
			}
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

func (c *Committee) loadParties() {
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
func (c *Committee) saveCommittee() {
	newMap := make(map[string]*tss.PartyID)
	nodes := c.parties.Items()
	for _, node := range nodes {
		newMap[node.Id] = node
	}
	c.currentCommittee.Replace(newMap)
}

func (c *Committee) shouldProposeKeygen() bool {
	//add in epoch expiry
	if c.keyGenInProgress == false && c.parties.Len() >= c.forkThreshold && (c.currentState == nil || c.currentState.SecretShare == nil) {
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
				c.loadParties()
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
						if c.currentState == nil || c.currentState.SecretShare == nil || c.epochSource.Expired(c.currentState.EpochStart) {
							if proposingNodes := c.nodeLabels.Get(WantsKeygen); proposingNodes != nil && proposingNodes.Len() >= c.forkThreshold {
								parties := []*tss.PartyID{}
								c.saveCommittee()
								for _, node := range proposingNodes.Items() {
									if party := c.currentCommittee.Get(node); party != nil {
										parties = append(parties, party)
									}
								}
								ctxKg, _ := context.WithDeadline(context.Background(), time.Now().Add(5*time.Minute))
								go c.StartKeygen(
									ctxKg,
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
					if c.keyGenInProgress {
						msg := message.GetParsedContent().GetKeygen()
						log.Debug().Uint32("round", msg.Round).Msg("processing message")
						if msg == nil {
							log.Error().Msg("Invalid keygen message")
						}
						if pa := c.currentCommittee.Get(message.From); pa != nil {
							cq := c.keyGenMsgQueue.Get(msg.Round)
							if cq == nil {
								q := lane.NewQueue[*keygenUpdate]()
								c.keyGenMsgQueue.Set(msg.Round, q)
								cq = c.keyGenMsgQueue.Get(msg.Round)
							}
							cq.Enqueue(&keygenUpdate{
								From:    pa,
								To:      c.selfId,
								Message: msg,
							})
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
