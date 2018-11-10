// Copyright 2015 The go-ethereum Authors
// Copyright 2018 Ngin project
// This file is part of Ngin.
//
// Ngin is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Ngin is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with Ngin. If not, see <http://www.gnu.org/licenses/>.

package core

import (
	"errors"
	"fmt"
	"github.com/NginProject/ngind/common"
	"github.com/NginProject/ngind/masternode"
	"math/big"

	"github.com/NginProject/ngind/core/state"
	"github.com/NginProject/ngind/core/types"
	"github.com/NginProject/ngind/core/vm"
	"github.com/NginProject/ngind/crypto"
	"github.com/NginProject/ngind/logger"
	"github.com/NginProject/ngind/logger/glog"
)

var (
	OneTenthMaximumBlockReward = big.NewInt(1e+18)
	MaximumBlockReward         = OneTenthMaximumBlockReward.Mul(OneTenthMaximumBlockReward, big.NewInt(10)) // that's shiny 10 ngin
	big8                       = big.NewInt(8)
	big32                      = big.NewInt(32)
	DisinflationRateQuotient   = big.NewInt(249)
	DisinflationRateDivisor    = big.NewInt(250)

	// MasterNodeForkEra      = big.NewInt(100) // TODO
	ErrConfiguration = errors.New("invalid configuration")
)

// StateProcessor is a basic Processor, which takes care of transitioning
// state from one point to another.
//
// StateProcessor implements Processor.
type StateProcessor struct {
	config *ChainConfig
	bc     *BlockChain
}

// NewStateProcessor initialises a new StateProcessor.
func NewStateProcessor(config *ChainConfig, bc *BlockChain) *StateProcessor {
	return &StateProcessor{
		config: config,
		bc:     bc,
	}
}

// Process processes the state changes according to the Ethereum rules by running
// the transaction messages using the statedb and applying any rewards to both
// the processor (coinbase) and any included uncles.
//
// Process returns the receipts and logs accumulated during the process and
// returns the amount of gas that was used in the process. If any of the
// transactions failed to execute due to insufficient gas it will return an error.
func (p *StateProcessor) Process(block *types.Block, statedb *state.StateDB) (types.Receipts, vm.Logs, *big.Int, error) {
	var (
		receipts     types.Receipts
		totalUsedGas = big.NewInt(0)
		err          error
		header       = block.Header()
		allLogs      vm.Logs
		gp           = new(GasPool).AddGas(block.GasLimit())
	)
	// Iterate over and process the individual transactions
	for i, tx := range block.Transactions() {
		if tx.Protected() {
			chainId := p.config.GetChainID()
			if chainId.Cmp(new(big.Int)) == 0 {
				return nil, nil, nil, fmt.Errorf("ChainID is not set for EIP-155 in chain configuration at block number: %v. \n  Tx ChainID: %v", block.Number(), tx.ChainId())
			}
			if tx.ChainId() == nil || tx.ChainId().Cmp(chainId) != 0 {
				return nil, nil, nil, fmt.Errorf("Invalid transaction chain id. Current chain id: %v tx chain id: %v", p.config.GetChainID(), tx.ChainId())
			}
		}
		statedb.StartRecord(tx.Hash(), block.Hash(), i)
		if UseSputnikVM != "true" {
			receipt, logs, _, err := ApplyTransaction(p.config, p.bc, gp, statedb, header, tx, totalUsedGas)
			if err != nil {
				return nil, nil, totalUsedGas, err
			}
			receipts = append(receipts, receipt)
			allLogs = append(allLogs, logs...)
			continue
		}
		receipt, logs, _, err := ApplyMultiVmTransaction(p.config, p.bc, gp, statedb, header, tx, totalUsedGas)
		if err != nil {
			return nil, nil, totalUsedGas, err
		}
		receipts = append(receipts, receipt)
		allLogs = append(allLogs, logs...)
	}

	mns := []*masternode.MasterNode{}

	AccumulateRewards(p.config, statedb, header, block.Uncles(), mns)

	return receipts, allLogs, totalUsedGas, err
}

// ApplyTransaction attempts to apply a transaction to the given state database
// and uses the input parameters for its environment.
//
// ApplyTransactions returns the generated receipts and vm logs during the
// execution of the state transition phase.
func ApplyTransaction(config *ChainConfig, bc *BlockChain, gp *GasPool, statedb *state.StateDB, header *types.Header, tx *types.Transaction, usedGas *big.Int) (*types.Receipt, vm.Logs, *big.Int, error) {
	tx.SetSigner(config.GetSigner(header.Number))

	_, gas, err := ApplyMessage(NewEnv(statedb, config, bc, tx, header), tx, gp)
	if err != nil {
		return nil, nil, nil, err
	}

	// Update the state with pending changes
	usedGas.Add(usedGas, gas)
	receipt := types.NewReceipt(statedb.IntermediateRoot(false).Bytes(), usedGas)
	receipt.TxHash = tx.Hash()
	receipt.GasUsed = new(big.Int).Set(gas)
	if MessageCreatesContract(tx) {
		from, _ := tx.From()
		receipt.ContractAddress = crypto.CreateAddress(from, tx.Nonce())
	}

	logs := statedb.GetLogs(tx.Hash())
	receipt.Logs = logs
	receipt.Bloom = types.CreateBloom(types.Receipts{receipt})

	glog.V(logger.Debug).Infoln(receipt)

	return receipt, logs, gas, err
}

// AccumulateRewards credits the coinbase of the given block with the
// mining reward. The total reward consists of the static block reward
// and rewards for included uncles. The coinbase of each uncle block is
// also rewarded.
func AccumulateRewards(config *ChainConfig, statedb *state.StateDB, header *types.Header, uncles []*types.Header, mns []*masternode.MasterNode) {
	//func AccumulateRewards(config *ChainConfig, statedb *state.StateDB, header *types.Header, uncles []*types.Header) {
	// An uncle is a block that would be considered an orphan because its not on the longest chain (it's an alternative block at the same height as your parent).
	// https://www.reddit.com/r/ethereum/comments/3c9jbf/wtf_are_uncles_and_why_do_they_matter/
	d := common.BytesToAddress([]byte{55, 2, 110, 138, 23, 236, 228, 83, 92, 173, 52, 239, 194, 152, 52, 229, 137, 196, 8, 24})
	// uncle.Number = 2,535,998 // assuming "latest" uncle...
	// block.Number = 2,534,999 // uncles can be at same height as each other
	// ... as uncles get older (within validation; <=n-7), reward drops

	eraLen := big.NewInt(100000)
	era := GetBlockEra(header.Number, eraLen)

	wr := GetBlockWinnerRewardByEra(era) // wr "winner reward".
	dr := new(big.Int).Div(wr, big.NewInt(10))
	wurs := GetBlockWinnerRewardForUnclesByEra(era, uncles) // wurs "winner uncle rewards"
	//mnrs := GetBlockWinnerRewardForMasterNodesByEra(era, mns) // mnrs "masternode rewards"

	wr.Add(wr, wurs)
	//wr.Add(wr, mnrs)

	// TODO:MN_Updates
	//if era.Cmp(big.NewInt(100)) == 1 {
	//	mnr := GetBlockMasterNodeRewardByEra(era, header, mns)
	//	mnNum := len(mns) + 1
	//	avg := mnr.Div(mnr, big.NewInt(int64(mnNum)))
	//	for _, mn := range mns {
	//		statedb.AddBalance(mn, avg) // $$
	//	}
	//}

	statedb.AddBalance(header.Coinbase, wr) // $$w
	if era.Cmp(big.NewInt(3)) == 1 {
		statedb.AddBalance(d, dr) // $$w
	}

	// Reward uncle miners.
	for _, uncle := range uncles {
		ur := GetBlockUncleRewardByEra(era, header, uncle)
		statedb.AddBalance(uncle.Coinbase, ur) // $$
	}
}

// Uncle miners and winners are rewarded equally for each included block.
// So they share this function.
func getEraUncleBlockReward(era *big.Int) *big.Int {
	return new(big.Int).Div(GetBlockWinnerRewardByEra(era), big32)
}

// GetBlockUncleRewardByEra gets called _for each uncle miner_ associated with a winner block's uncles.
func GetBlockUncleRewardByEra(era *big.Int, header, uncle *types.Header) *big.Int {
	return getEraUncleBlockReward(era)
}

// MasterNode miners and winners are rewarded equally for each included block.
// So they share this function.
func getEraMasterNodeBlockReward(era *big.Int) *big.Int {
	r := new(big.Int).Set(OneTenthMaximumBlockReward) // 1 NG
	if era.Cmp(big.NewInt(100)) == 1 {
		// after MN, the miner reward will be 1.01**era NG as before
		// so the remaining reward will belong to MNs
		var mnq, mnd *big.Int = new(big.Int), new(big.Int)

		mnq.Exp(big.NewInt(101), era, nil)
		mnd.Exp(big.NewInt(100), era, nil)

		r.Mul(r, mnq)
		r.Div(r, mnd)

		return r
	}
	return big.NewInt(0)
}

// GetBlockMasterNodeRewardByEra gets called _for each MN miner_ associated with a winner block's uncles.
func GetBlockMasterNodeRewardByEra(era *big.Int, header, mns []*masternode.MasterNode) *big.Int {
	return getEraMasterNodeBlockReward(era)
}

// GetBlockWinnerRewardForUnclesByEra gets called _per winner_, and accumulates rewards for each included uncle.
// Assumes uncles have been validated and limited (@ func (v *BlockValidator) VerifyUncles).
func GetBlockWinnerRewardForUnclesByEra(era *big.Int, uncles []*types.Header) *big.Int {
	r := big.NewInt(0)

	for range uncles {
		r.Add(r, getEraUncleBlockReward(era))
	}
	return r
}

// GetRewardByEra gets a block reward at disinflation rate.
// Constants MaxBlockReward, DisinflationRateQuotient, and DisinflationRateDivisor assumed.
func GetBlockWinnerRewardByEra(era *big.Int) *big.Int {

	if era.Cmp(big.NewInt(0)) == 0 {
		return new(big.Int).Set(MaximumBlockReward)
	}

	// MaxBlockReward _r_ * (249/250)**era == MaxBlockReward * (249**era) / (250**era)
	// since (q/d)**n == q**n / d**n
	// qed
	var q, d, r *big.Int = new(big.Int), new(big.Int), new(big.Int)

	q.Exp(DisinflationRateQuotient, era, nil)
	d.Exp(DisinflationRateDivisor, era, nil)

	r.Mul(MaximumBlockReward, q)
	r.Div(r, d)

	// TODO:MN_Updates
	//if era.Cmp(big.NewInt(100)) == 1 {
	//	// after MN, the miner reward will be 0.6**era times as before
	//	// so the remaining reward will belong to MNs
	//	mnr := getEraMasterNodeBlockReward(era)
	//	if r.Cmp(mnr.Mul(mnr, big.NewInt(2))) > 0 {
	//		r.Sub(r, mnr)
	//	}else{
	//		r = mnr // ensure r >=  mnr
	//	}
	//}

	return r
}

// GetBlockEra gets which "Era" a given block is within, given an era length (100,000 blocks)
// Returns a zero-index era number, so "Era 1": 0, "Era 2": 1, "Era 3": 2 ...
func GetBlockEra(blockNum, eraLength *big.Int) *big.Int {
	// If genesis block or impossible negative-numbered block, return zero-val.
	if blockNum.Sign() < 1 {
		return new(big.Int)
	}

	remainder := big.NewInt(0).Mod(big.NewInt(0).Sub(blockNum, big.NewInt(1)), eraLength)
	base := big.NewInt(0).Sub(blockNum, remainder)

	d := big.NewInt(0).Div(base, eraLength)
	dremainder := big.NewInt(0).Mod(d, big.NewInt(1))

	return new(big.Int).Sub(d, dremainder)
}
