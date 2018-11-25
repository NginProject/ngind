package masternode

import (
	"net"
	"sync"
	"crypto/ecdsa"
	"errors"

	"github.com/NginProject/ngind/p2p"
	"github.com/NginProject/ngind/p2p/discover"
	"github.com/NginProject/ngind/crypto"
	"github.com/NginProject/ngind/common"

)

const (
	ACTIVE_MASTERNODE_INITIAL         = 0 // initial state
	ACTIVE_MASTERNODE_SYNCING         = 2
	ACTIVE_MASTERNODE_NOT_CAPABLE     = 3
	ACTIVE_MASTERNODE_STARTED         = 4
)

// ErrUnknownMasternode is returned for any requested operation for which no backend
// provides the specified masternode.
var ErrUnknownMasternode = errors.New("unknown masternode")

//Responsible for activating the Masternode and pinging the network
type ActiveMasternode struct {
	ID          string
	NodeID      discover.NodeID
	NodeAccount common.Address
	PrivateKey  *ecdsa.PrivateKey
	activeState int
	Addr        net.TCPAddr

	mu sync.RWMutex
}

func NewActiveMasternode(srvr *p2p.Server) *ActiveMasternode {
	nodeId := srvr.Self().ID
	id := GetMasternodeID(nodeId)
	am := &ActiveMasternode{
		ID:          id,
		NodeID:      nodeId,
		activeState: ACTIVE_MASTERNODE_INITIAL,
		PrivateKey:  srvr.Config.PrivateKey,
		NodeAccount: crypto.PubkeyToAddress(srvr.Config.PrivateKey.PublicKey),
	}
	return am
}

func (am *ActiveMasternode) State() int {
	return am.activeState
}

func (am *ActiveMasternode) SetState(state int) {
	am.activeState = state
}

// SignHash calculates a ECDSA signature for the given hash. The produced
// signature is in the [R || S || V] format where V is 0 or 1.
func (a *ActiveMasternode) SignHash(id string, hash []byte) ([]byte, error) {
	// Look up the key to sign with and abort if it cannot be found
	a.mu.RLock()
	defer a.mu.RUnlock()

	if id != a.ID{
		return nil, ErrUnknownMasternode
	}
	// Sign the hash using plain ECDSA operations
	return crypto.Sign(hash, a.PrivateKey)
}