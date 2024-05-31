package cluster

import (
	"sync"

	"github.com/anagrambuild/ferrule/schemas"
	"github.com/anagrambuild/ferrule/util"
	"github.com/hashicorp/memberlist"
	"go.uber.org/zap"
)

type state struct {
	data   []byte
	rwLock sync.RWMutex
}
type ClusterOperator struct {
	memberlist.Delegate
	memberlist.EventDelegate
	mlistConfig  *memberlist.Config
	Memberlist   *memberlist.Memberlist
	tlq          *memberlist.TransmitLimitedQueue
	shutdown     chan struct{}
	memberEvents chan<- memberlist.NodeEvent
	newMessages  chan<- *schemas.ClusterMessage
	nodeMeta     state
	localState   state
	logger       *zap.Logger
	nodeMap      util.SyncedMap[string, *memberlist.Node]
}

func NewClusterOperator(
	selfName string,
	newMessages chan<- *schemas.ClusterMessage,
	memberEvents chan<- memberlist.NodeEvent,
) *ClusterOperator {
	dlg := &ClusterOperator{
		newMessages:  newMessages,
		memberEvents: memberEvents,
		nodeMeta: state{
			data:   []byte{},
			rwLock: sync.RWMutex{},
		},
		localState: state{
			data:   []byte{},
			rwLock: sync.RWMutex{},
		},
		nodeMap: util.NewSyncedMap[string, *memberlist.Node](),
	}
	config := memberlist.DefaultWANConfig()
	config.Delegate = dlg
	config.Name = selfName
	config.Events = dlg
	dlg.mlistConfig = config
	return dlg
}

func (d *ClusterOperator) Broadcast(cast ClusterMsgBroadcast) {
	d.tlq.QueueBroadcast(cast)
}

func (d *ClusterOperator) Started() bool {
	return d.Memberlist != nil && d.tlq != nil
}

func (d *ClusterOperator) Shutdown() {
	d.shutdown <- struct{}{}
	d.Memberlist.Shutdown()
}

func (d *ClusterOperator) NotifyJoin(node *memberlist.Node) {
	d.nodeMap.Set(node.Name, node)
	d.memberEvents <- memberlist.NodeEvent{
		Node:  node,
		Event: memberlist.NodeJoin,
	}
}

func (d *ClusterOperator) NotifyLeave(node *memberlist.Node) {
	d.nodeMap.Delete(node.Name)
	d.memberEvents <- memberlist.NodeEvent{
		Node:  node,
		Event: memberlist.NodeLeave,
	}
}

func (d *ClusterOperator) NotifyUpdate(node *memberlist.Node) {
	d.nodeMap.Set(node.Name, node)
	d.memberEvents <- memberlist.NodeEvent{
		Node:  node,
		Event: memberlist.NodeUpdate,
	}
}

func (d *ClusterOperator) GetNode(name string) *memberlist.Node {
	return d.nodeMap.Get(name)
}

func (d *ClusterOperator) Start() error {
	mlist, err := memberlist.Create(d.mlistConfig)
	if err != nil {
		return err
	}
	d.Memberlist = mlist
	d.tlq = &memberlist.TransmitLimitedQueue{
		NumNodes: func() int {
			return d.Memberlist.NumMembers()
		},
		RetransmitMult: 2,
	}
	return nil
}

func (d *ClusterOperator) SetNodeMeta(meta []byte) {
	d.nodeMeta.rwLock.Lock()
	defer d.nodeMeta.rwLock.Unlock()
	d.nodeMeta.data = meta
}

func (d *ClusterOperator) SetLocalState(state []byte) {
	d.localState.rwLock.Lock()
	defer d.localState.rwLock.Unlock()
	d.localState.data = state
}

func (d *ClusterOperator) NodeMeta(limit int) []byte {
	d.nodeMeta.rwLock.RLock()
	defer d.nodeMeta.rwLock.RUnlock()
	return d.nodeMeta.data
}

func (d *ClusterOperator) NotifyMsg(msg []byte) {
	clusterMsg, err := schemas.DecodeClusterMessage(msg)
	if err != nil {
		d.logger.Error("Failed to decode cluster message", zap.Error(err))
		return
	}
	d.newMessages <- clusterMsg
}

func (d *ClusterOperator) GetBroadcasts(overhead, limit int) [][]byte {
	return d.tlq.GetBroadcasts(overhead, limit)
}

func (d *ClusterOperator) LocalState(join bool) []byte {
	// disabled for now
	// d.localState.rwLock.RLock()
	// defer d.localState.rwLock.RUnlock()
	// return d.localState.data
	return []byte{}
}

func (d *ClusterOperator) MergeRemoteState(buf []byte, join bool) {
	// disabled for now
	// d.localState.rwLock.Lock()
	// defer d.localState.rwLock.Unlock()
	// // this is probbaly wrong/ need to use a crdt data structure
	// d.localState.data = buf
}
