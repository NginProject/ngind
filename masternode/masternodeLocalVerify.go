package masternode

import (
	"github.com/NginProject/ngind/common"
	"github.com/NginProject/ngind/core/vm"
	"math/big"
)

// fulfillRequirement check the node is a MasterNode or Not
// first, it will check the ngin balance > currentEra * 100000 * 10 / 50
// then, check system hardware
// require run with --nat:extip:77.12.33.4
// require run with --masternode
func selfFulfillRequirements(env vm.Environment) bool {
	if env.Db().GetBalance(env.Coinbase()).Cmp(getCurrentMasterNodeBalanceThrehold(env)) > 0 {
		return true
	}
	return false
}

func checkFulfillRequirements(addr common.Address, env vm.Environment) bool {
	if env.Db().GetBalance(addr).Cmp(getCurrentMasterNodeBalanceThrehold(env)) > 0 {
		return true
	}
	return false
}

func getCurrentMasterNodeBalanceThrehold(env vm.Environment) *big.Int {
	blkNum := env.BlockNumber()
	return blkNum.Div(blkNum, big.NewInt(5)) // 1/50 cir supply
}
