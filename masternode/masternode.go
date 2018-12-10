package masternode

import (
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/NginProject/ngind/common"
	"github.com/NginProject/ngind/p2p/discover"
)

const (
	MasternodeInit = iota
)

const (
	MASTERNODE_PING_INTERVAL  = 1200 * time.Second
)

var (
	errClosed            = errors.New("masternode set is closed")
	errAlreadyRegistered = errors.New("masternode is already registered")
	errNotRegistered     = errors.New("masternode is not registered")
)

type Masternode struct {
	ID          string
	NodeID      discover.NodeID
	Account     common.Address
	OriginBlock *big.Int
	State       int

	BlockOnlineAcc *big.Int
	BlockLastPing  *big.Int
}

func New(nodeId discover.NodeID, account common.Address, block, blockOnlineAcc, blockLastPing *big.Int) *Masternode {

	id := GetMasternodeID(nodeId)
	return &Masternode{
		ID:             id,
		NodeID:         nodeId,
		Account:        account,
		OriginBlock:    block,
		State:          MasternodeInit,
		BlockOnlineAcc: blockOnlineAcc,
		BlockLastPing:  blockLastPing,
	}
}

func (n *Masternode) String() string {
	return fmt.Sprintf("Node: %s\n", n.NodeID.String())
}

func GetMasternodeID(ID discover.NodeID) string {
	return fmt.Sprintf("%x", ID[:8])
}
