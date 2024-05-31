package cluster

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"time"

	"encoding/hex"

	"github.com/anagrambuild/ferrule/config"
	"github.com/anagrambuild/ferrule/schemas"
	"github.com/anagrambuild/ferrule/util"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	"github.com/hashicorp/memberlist"
	"github.com/rs/zerolog/log"
)

type Self struct {
	Identity *ecdsa.PrivateKey
	encKey   *ecies.PrivateKey
	Address  string
	Node     *memberlist.Node
}

type ClusterMessageListener struct {
	Id        string
	Predicate func(message *schemas.ClusterMessage) bool
	Listener  func(message *schemas.ClusterMessage) error
}
type ListenerOpType int

var AddListenerOp ListenerOpType = 0
var RemoveListenerOp ListenerOpType = 1

type ListenerOp struct {
	Op       ListenerOpType
	Listener *ClusterMessageListener
}

type ClusterEventListener struct {
	Id       string
	Listener func(event memberlist.NodeEvent)
}

type EventListenerOp struct {
	Op       ListenerOpType
	Listener *ClusterEventListener
}

type Cluster struct {
	self                *Self
	peers               []string
	operator            *ClusterOperator
	shutdown            chan struct{}
	messages            chan *schemas.ClusterMessage
	events              chan memberlist.NodeEvent
	listenerOpChan      chan ListenerOp
	eventListenerOpChan chan EventListenerOp
	decoderFn           *ClusterMessageListener
}

type ClusterMsgBroadcast struct {
	msg []byte
}

func (b ClusterMsgBroadcast) Message() []byte {
	return b.msg
}

func (b ClusterMsgBroadcast) Invalidates(a memberlist.Broadcast) bool {
	return false
}

func (cl ClusterMsgBroadcast) Finished() {
}

func NewCluster(
	cfg *config.FerruleConfig,
) *Cluster {
	messages := make(chan *schemas.ClusterMessage)
	events := make(chan memberlist.NodeEvent)
	shutdown := make(chan struct{})
	eciesKey := ecies.ImportECDSA(cfg.Identity)
	cluster := &Cluster{
		self: &Self{
			Identity: cfg.Identity,
			Address:  cfg.NodeName,
			encKey:   eciesKey,
		},
		peers:    cfg.Peers,
		shutdown: shutdown,
		operator: NewClusterOperator(
			cfg.NodeName,
			messages,
			events,
		),
		eventListenerOpChan: make(chan EventListenerOp),
		listenerOpChan:      make(chan ListenerOp),
		messages:            messages,
		events:              events,
		decoderFn: &ClusterMessageListener{
			Id: "decoder",
			Predicate: func(message *schemas.ClusterMessage) bool {
				return true
			},
			Listener: func(message *schemas.ClusterMessage) error {
				if message.Encrypted {
					err := message.Decrypt(eciesKey)
					if err != nil {
						return err
					}

				} else {
					if !message.VerifySelf() {
						return errors.New("failed to verify message")
					}
				}
				return message.DecodeContent()
			},
		},
	}
	return cluster
}

func (cl *Cluster) Start() error {
	log.Info().Msg("Starting cluster")

	go cl.recvMessages(cl.shutdown)
	go cl.recvEvents(cl.shutdown)
	err := cl.operator.Start()
	if err != nil {
		return err
	}
	log.Info().Msg("Started cluster")
	if !cl.Head() {
		err := util.BackoffRetry(
			func() error {
				_, err = cl.operator.Memberlist.Join(cl.peers)
				return err
			},
			5,
			time.Duration(5)*time.Second,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func validateMessage(
	message *schemas.ClusterMessage,
) error {
	if message == nil {
		return errors.New("message is nil")
	}
	if message.Id == nil || len(message.Id) == 0 {
		return errors.New("missing id field in message")
	}
	if message.From == "" {
		return errors.New("missing from field in message")
	}
	if message.Encrypted && (message.Salt == nil || message.Nonce == nil || len(message.Salt) == 0 || len(message.Nonce) == 0) {
		return errors.New("missing salt or nonce for encrypted message")
	} else if !message.Encrypted && (message.Signature == nil || len(message.Signature) == 0) {
		return errors.New("missing signature for message")
	}
	if message.Content == nil && message.ParsedContent != nil {
		return errors.New("message content must be encoded")
	}
	if message.Content == nil && message.ParsedContent == nil {
		return errors.New("message content is nil")
	}
	return nil
}

func (cl *Cluster) Broadcast(message *schemas.ClusterMessage) error {
	err := message.EncodeContent()
	if err != nil {
		return err
	}

	if message.Encrypted {
		message.Encrypted = false
	}
	err = message.Sign(cl.self.Identity)
	if err != nil {
		return err
	}

	if err := validateMessage(message); err != nil {
		fmt.Println("error validating message", err)
		return err
	}
	bts, err := schemas.EncodeClusterMessage(message)
	if err != nil {
		return err
	}
	cl.operator.Broadcast(ClusterMsgBroadcast{
		msg: bts,
	})
	return nil
}

func (cl *Cluster) SendToNode(node *memberlist.Node, message *schemas.ClusterMessage) error {
	err := message.EncodeContent()
	if err != nil {
		return err
	}
	if message.Encrypted {
		message.Encrypted = false
	}
	err = message.Sign(cl.self.Identity)
	if err := validateMessage(message); err != nil {
		return err
	}
	bytes, err := schemas.EncodeClusterMessage(message)
	if err != nil {
		return err
	}
	return cl.operator.Memberlist.SendReliable(node, bytes)
}

func (cl *Cluster) SendToNodeEncrypted(node *memberlist.Node, message *schemas.ClusterMessage) error {
	err := message.EncodeContent()
	if err != nil {
		return err
	}
	eciesPubKey, err := getEciesPubKey(node)
	if err != nil {
		return err
	}
	err = message.Encrypt(
		cl.self.Identity,
		cl.self.encKey, eciesPubKey)
	if err != nil {
		return err
	}
	if err := validateMessage(message); err != nil {
		return err
	}
	bytes, err := schemas.EncodeClusterMessage(message)
	if err != nil {
		return err
	}
	return cl.operator.Memberlist.SendReliable(node, bytes)
}

func (cl *Cluster) Shutdown() {
	close(cl.shutdown)
}

func (cl *Cluster) SelfNode() *memberlist.Node {
	return cl.operator.Memberlist.LocalNode()
}

func (cl *Cluster) LookupNode(name string) *memberlist.Node {
	log.Debug().Str("name", name).Msg("Looking up node")
	log.Debug().Str("nodes", fmt.Sprintf("%+v", cl.operator.nodeMap.Items())).Msg("Nodes")
	return cl.operator.GetNode(name)
}

func (cl *Cluster) GetNodes() []*memberlist.Node {
	return cl.operator.Memberlist.Members()
}
func (cl *Cluster) NumNodes() int {
	return cl.operator.Memberlist.NumMembers()
}

func (cl *Cluster) AddEventListener(listener *ClusterEventListener) {
	cl.eventListenerOpChan <- EventListenerOp{
		Op:       AddListenerOp,
		Listener: listener,
	}
}

func (cl *Cluster) RemoveEventListener(listener *ClusterEventListener) {
	cl.eventListenerOpChan <- EventListenerOp{
		Op:       RemoveListenerOp,
		Listener: listener,
	}
}

func (cl *Cluster) AddListener(listener *ClusterMessageListener) {
	cl.listenerOpChan <- ListenerOp{
		Op:       AddListenerOp,
		Listener: listener,
	}
}

func (cl *Cluster) RemoveListener(listener *ClusterMessageListener) {
	cl.listenerOpChan <- ListenerOp{
		Op:       RemoveListenerOp,
		Listener: listener,
	}
}

func (cl *Cluster) recvEvents(
	shutdown chan struct{},
) {
	listeners := []ClusterEventListener{}
	for {
		select {
		case <-shutdown:
			{
				return
			}
		case levent := <-cl.eventListenerOpChan:
			{
				switch levent.Op {
				case AddListenerOp:
					{
						listeners = append(listeners, *levent.Listener)
						continue
					}
				case RemoveListenerOp:
					{
						for i, listener := range listeners {
							if listener.Id == levent.Listener.Id {
								listeners = append(listeners[:i], listeners[i+1:]...)
								break
							}
						}
						continue
					}
				}
			}
		case event := <-cl.events:
			for _, listener := range listeners {
				go listener.Listener(event)
			}
		}
	}
}

func getEciesPubKey(
	node *memberlist.Node,
) (*ecies.PublicKey, error) {
	pubkey := node.Name
	hex, err := hex.DecodeString(pubkey)
	if err != nil {
		return nil, err
	}
	pub, err := ethcrypto.UnmarshalPubkey(hex)
	if err != nil {
		return nil, err
	}
	ePub := ecies.ImportECDSAPublic(pub)
	return ePub, nil
}

func (cl *Cluster) recvMessages(
	shutdown chan struct{},
) {
	listeners := []ClusterMessageListener{
		*cl.decoderFn,
	}
	for {
		select {
		case <-shutdown:
			{
				return
			}
		case levent := <-cl.listenerOpChan:
			{
				switch levent.Op {
				case AddListenerOp:
					{
						listeners = append(listeners, *levent.Listener)
						continue
					}
				case RemoveListenerOp:
					{
						for i, listener := range listeners {
							if listener.Id == levent.Listener.Id {
								listeners = append(listeners[:i], listeners[i+1:]...)
								break
							}
						}
						continue
					}
				}
			}
		case message := <-cl.messages:
			for _, listener := range listeners {
				if listener.Predicate(message) {
					err := listener.Listener(message)
					if err != nil {
						log.Error().Err(err).Msg("Error processing message")
					}
				}
			}
		}
	}
}

func (cl *Cluster) Head() bool {
	return len(cl.peers) == 0
}
