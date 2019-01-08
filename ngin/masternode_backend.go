package ngin

import (
	"fmt"
	"github.com/NginProject/ngind/logger"
	"github.com/NginProject/ngind/logger/glog"
	"github.com/NginProject/ngind/masternode/contract"
	"github.com/NginProject/ngind/ngindb"
	"math/big"
	"sync"

	"time"

	"github.com/NginProject/ngind/common"

	"github.com/NginProject/ngind/core"
	"github.com/NginProject/ngind/masternode"
	"github.com/NginProject/ngind/ngin/downloader"
	"github.com/NginProject/ngind/p2p"
	"github.com/NginProject/ngind/p2p/discover"
)

var (
	statsReportInterval = 10 * time.Second // Time interval to report vote pool stats
)

type MasternodeManager struct {
	beats map[common.Hash]time.Time // Last heartbeat from each known vote

	db     *ngindb.Database
	active *masternode.ActiveMasternode
	mu     sync.Mutex
	// channels for fetcher, syncer, txsyncLoop
	newPeerCh chan *peer

	IsMasternode uint32
	srvr         *p2p.Server
	blockchain   *core.BlockChain

	currentCycle uint64        // Current vote of the block chain
	Lifetime     time.Duration // Maximum amount of time vote are queued

	txPool *core.TxPool

	downloader *downloader.Downloader
}

func NewMasternodeManager(db *ngindb.MemDatabase, blockchain *core.BlockChain,contract *contract.MN, txPool *core.TxPool) *MasternodeManager {

	// Create the masternode manager with its initial settings
	manager := &MasternodeManager{
		db:         db,
		blockchain: blockchain,
		beats:      make(map[common.Hash]time.Time),
		Lifetime:   13 * time.Second,
		contract:   contract,
		txPool:     txPool,
	}
	return manager
}

func (self *MasternodeManager) Clear() {
	self.mu.Lock()
	defer self.mu.Unlock()

}

func (self *MasternodeManager) Start(srvr *p2p.Server, peers *peerSet, downloader *downloader.Downloader) {
	self.srvr = srvr
	self.downloader = downloader
	glog.D(logger.Detail).Infoln("Masternode Manager Start.")
	self.active = masternode.NewActiveMasternode(srvr)
	go self.masternodeLoop()
}

func (self *MasternodeManager) Stop() {

}

func (mm *MasternodeManager) masternodeLoop() {
	// MN Should DO

	ping := time.NewTimer(masternode.MASTERNODE_PING_INTERVAL)
	defer ping.Stop()
	ntp := time.NewTimer(time.Second)
	defer ntp.Stop()

	report := time.NewTicker(statsReportInterval)
	defer report.Stop()

	for {
		select {
		case <-ntp.C:
			ntp.Reset(10 * time.Minute)
			go discover.CheckClockDrift()
		case <-ping.C:
			ping.Reset(masternode.MASTERNODE_PING_INTERVAL)
			if mm.active.State() != masternode.ACTIVE_MASTERNODE_STARTED {
				break
			}
			if mm.downloader.Synchronising() {
				break
			}

			address := mm.active.NodeAccount
			stateDB, _ := mm.blockchain.State()

			// requirement
			currentBlockNum := mm.blockchain.CurrentBlock().Number()
			eraLen := big.NewInt(100000)
			era := core.GetBlockEra(currentBlockNum, eraLen)
			totalRequirement := new(big.Int).Mul(big.NewInt(1e+18), new(big.Int).Mul(era, eraLen.Mul(eraLen, big.NewInt(10))))
			requirement := totalRequirement.Div(totalRequirement, big.NewInt(50))
			if stateDB.GetBalance(address).Cmp(requirement) < 0 {
				fmt.Println("Failed to Check Masternode balance: ", address.Hex())
				break
			}
			fmt.Println("Send ping message ...")
		}
	}
}

func (mm *MasternodeManager) updateActiveMasternode(isMasternode bool) {
	var state int
	if isMasternode {
		state = masternode.ACTIVE_MASTERNODE_STARTED
	} else {
		state = masternode.ACTIVE_MASTERNODE_NOT_CAPABLE
	}
	mm.active.SetState(state)
}
