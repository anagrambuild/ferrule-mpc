package cluster

import (
	"crypto/ecdh"
	"crypto/ecdsa"
	"errors"

	"github.com/anagrambuild/ferrule/config"
	"github.com/anagrambuild/ferrule/schemas"
	"github.com/hashicorp/memberlist"
	"go.uber.org/zap"
)

type Self struct {
	Identity *ecdsa.PrivateKey
	ecdhKey  *ecdh.PrivateKey
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
	logger              *zap.Logger
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
	cfg config.FerruleConfig,
) *Cluster {
	logger, _ := zap.NewProduction()
	messages := make(chan *schemas.ClusterMessage)
	events := make(chan memberlist.NodeEvent)
	shutdown := make(chan struct{})
	ecdh, err := cfg.Identity.ECDH()
	if err != nil {
		logger.Error("Failed to derive ecdh key", zap.Error(err))
		panic(err)
	}
	cluster := &Cluster{
		self: &Self{
			Identity: cfg.Identity,
			Address:  cfg.Address,
			ecdhKey:  ecdh,
		},
		peers:    cfg.Peers,
		shutdown: shutdown,
		logger:   logger,
		operator: NewClusterOperator(
			cfg.Address,
			messages,
			events,
		),
		messages: messages,
		events:   events,
		decoderFn: &ClusterMessageListener{
			Id: "decoder",
			Predicate: func(message *schemas.ClusterMessage) bool {
				return true
			},
			Listener: func(message *schemas.ClusterMessage) error {
				if message.Encrypted {
					err := message.Decrypt(ecdh)
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
	err := cl.operator.Start()
	if err != nil {
		return err
	}
	go cl.recvMessages()
	go cl.recvEvents()
	_, err = cl.operator.Memberlist.Join(cl.peers)
	if err != nil {
		return err
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
	if message.From == nil || len(message.From) == 0 {
		return errors.New("missing from field in message")
	}
	if message.Encrypted && (message.Salt == nil || message.Nonce == nil || len(message.Salt) == 0 || len(message.Nonce) == 0) {
		return errors.New("missing salt or nonce for encrypted message")
	} else if !message.Encrypted && (message.Signature == nil || len(message.Signature) == 0) {
		return errors.New("missing signature for message")
	}
	return nil
}

func (cl *Cluster) Broadcast(message *schemas.ClusterMessage) error {
	if err := validateMessage(message); err != nil {
		return err
	}
	message.EncodeContent()
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
	if err := validateMessage(message); err != nil {
		return err
	}
	message.EncodeContent()
	bytes, err := schemas.EncodeClusterMessage(message)
	if err != nil {
		return err
	}
	return cl.operator.Memberlist.SendReliable(node, bytes)
}

func (cl *Cluster) Shutdown() {
	cl.shutdown <- struct{}{}
}

func (cl *Cluster) SelfNode() *memberlist.Node {
	return cl.operator.Memberlist.LocalNode()
}

func (cl *Cluster) LookupNode(name string) *memberlist.Node {
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

func (cl *Cluster) recvEvents() {
	listeners := []ClusterEventListener{}
	for {
		select {
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
		case <-cl.shutdown:
			return
		case event := <-cl.events:
			for _, listener := range listeners {
				go listener.Listener(event)
			}
		}
	}
}

func (cl *Cluster) recvMessages() {
	listeners := []ClusterMessageListener{
		*cl.decoderFn,
	}
	for {
		select {
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
		case <-cl.shutdown:
			return
		case message := <-cl.messages:
			cl.logger.Debug("Received message", zap.Any("message", &message))
			for _, listener := range listeners {
				if listener.Predicate(message) {
					listener.Listener(message)
				}
			}
		}
	}
}

func (cl *Cluster) Head() bool {
	return len(cl.peers) == 0
}
