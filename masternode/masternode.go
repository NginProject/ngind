// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package masternode

import (
	"github.com/NginProject/ngind/accounts/abi"
	"github.com/NginProject/ngind/accounts/abi/bind"
	"github.com/NginProject/ngind/common"
	"github.com/NginProject/ngind/core/types"
	"math/big"
	"strings"
)

// MNABI is the input ABI used to generate the binding from.
const MNABI = `[{"constant":true,"inputs":[{"name":"","type":"address"}],"name":"nodeList","outputs":[{"name":"balance","type":"uint256"},{"name":"isActive","type":"bool"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[],"name":"withdraw","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"account","type":"address"},{"name":"amount","type":"uint256"},{"name":"h","type":"uint256"}],"name":"checkDeposit","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[],"name":"updateStatus","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"","type":"uint256"}],"name":"itList","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"h","type":"uint256"}],"name":"circulatingSupply","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"pure","type":"function"},{"constant":false,"inputs":[{"name":"amount","type":"uint256"}],"name":"deposit","outputs":[],"payable":true,"stateMutability":"payable","type":"function"},{"payable":true,"stateMutability":"payable","type":"fallback"},{"anonymous":false,"inputs":[{"indexed":false,"name":"_account","type":"address"}],"name":"Active","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"_account","type":"address"}],"name":"Deactive","type":"event"}]`

// MNBin is the compiled bytecode used for deploying new contracts.
const MNBin = `0x608060405234801561001057600080fd5b50610681806100206000396000f3fe60806040526004361061007c577c0100000000000000000000000000000000000000000000000000000000600035046337befda481146100865780633ccfd60b146100d25780634a7e60f0146100e75780635f60f4851461013a578063649862ee1461014f57806392ff6aea14610195578063b6b55f25146101d1575b6100846101ee565b005b34801561009257600080fd5b506100b9600480360360208110156100a957600080fd5b5035600160a060020a03166102f8565b6040805192835290151560208301528051918290030190f35b3480156100de57600080fd5b50610084610314565b3480156100f357600080fd5b506101266004803603606081101561010a57600080fd5b50600160a060020a0381351690602081013590604001356103e8565b604080519115158252519081900360200190f35b34801561014657600080fd5b506100846101ee565b34801561015b57600080fd5b506101796004803603602081101561017257600080fd5b5035610429565b60408051600160a060020a039092168252519081900360200190f35b3480156101a157600080fd5b506101bf600480360360208110156101b857600080fd5b5035610451565b60408051918252519081900360200190f35b610084600480360360208110156101e757600080fd5b50356104de565b60005b6000548110156102f55761024d60008281548110151561020d57fe5b60009182526020822001548154600160a060020a0390911691908490811061023157fe5b600091825260209091200154600160a060020a031631436103e8565b156102a1576001806000808481548110151561026557fe5b600091825260208083209190910154600160a060020a031683528201929092526040019020600101805460ff19169115159190911790556102ed565b60006001600080848154811015156102b557fe5b600091825260208083209190910154600160a060020a031683528201929092526040019020600101805460ff19169115159190911790555b6001016101f1565b50565b6001602081905260009182526040909120805491015460ff1682565b3360008181526001602052604080822054905190929183156108fc02918491818181858888f19350505050158015610350573d6000803e3d6000fd5b5060005b6000548110156103c757600080543391908390811061036f57fe5b600091825260209091200154600160a060020a031614156103bf57600080548290811061039857fe5b6000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff191690555b600101610354565b5050336000908152600160208190526040822091825501805460ff19169055565b60008060326103f684610451565b8115156103ff57fe5b600160a060020a038716600090815260016020526040902054919004908501119150509392505050565b600080548290811061043757fe5b600091825260209091200154600160a060020a0316905081565b6000620186a0820481805b8281116104d657828114156104a75760008160fa0a8260f90a81151561047e57fe5b04678ac7230489e80000029050620186a084028681151561049b57fe5b060291909101906104ce565b60008160fa0a8260f90a8115156104ba57fe5b0469d3c21bcecceda1000000029290920191505b60010161045c565b509392505050565b34811461054c57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f4572726f723a2077726f6e67206465706f73697420616d6f756e740000000000604482015290519081900360640190fd5b600160005b60005481101561059857600080543391908390811061056c57fe5b600091825260209091200154600160a060020a031614156105905760009150610598565b600101610551565b5080156105ec57600080546001810182559080527f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e56301805473ffffffffffffffffffffffffffffffffffffffff1916331790555b33600081815260016020526040902080548401905561060c9083436103e8565b15610635573360009081526001602081905260409091208101805460ff19169091179055610651565b33600090815260016020819052604090912001805460ff191690555b505056fea165627a7a7230582023f0eb3faadace18ed4bea7081844d9b0f83568ce3b5ad4dfa391f61f37e2b7c0029`

// DeployMN deploys a new Ethereum contract, binding an instance of MN to it.
func DeployMN(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MN, error) {
	parsed, err := abi.JSON(strings.NewReader(MNABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MNBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MN{MNCaller: MNCaller{contract: contract}, MNTransactor: MNTransactor{contract: contract}}, nil
}

// MN is an auto generated Go binding around an Ethereum contract.
type MN struct {
	MNCaller     // Read-only binding to the contract
	MNTransactor // Write-only binding to the contract
}

// MNCaller is an auto generated read-only Go binding around an Ethereum contract.
type MNCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MNTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MNTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MNSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MNSession struct {
	Contract     *MN               // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MNCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MNCallerSession struct {
	Contract *MNCaller     // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MNTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MNTransactorSession struct {
	Contract     *MNTransactor     // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MNRaw is an auto generated low-level Go binding around an Ethereum contract.
type MNRaw struct {
	Contract *MN // Generic contract binding to access the raw methods on
}

// MNCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MNCallerRaw struct {
	Contract *MNCaller // Generic read-only contract binding to access the raw methods on
}

// MNTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MNTransactorRaw struct {
	Contract *MNTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMN creates a new instance of MN, bound to a specific deployed contract.
func NewMN(address common.Address, backend bind.ContractBackend) (*MN, error) {
	contract, err := bindMN(address, backend.(bind.ContractCaller), backend.(bind.ContractTransactor))
	if err != nil {
		return nil, err
	}
	return &MN{MNCaller: MNCaller{contract: contract}, MNTransactor: MNTransactor{contract: contract}}, nil
}

// NewMNCaller creates a new read-only instance of MN, bound to a specific deployed contract.
func NewMNCaller(address common.Address, caller bind.ContractCaller) (*MNCaller, error) {
	contract, err := bindMN(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &MNCaller{contract: contract}, nil
}

// NewMNTransactor creates a new write-only instance of MN, bound to a specific deployed contract.
func NewMNTransactor(address common.Address, transactor bind.ContractTransactor) (*MNTransactor, error) {
	contract, err := bindMN(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &MNTransactor{contract: contract}, nil
}

// bindMN binds a generic wrapper to an already deployed contract.
func bindMN(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MNABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MN *MNRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MN.Contract.MNCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MN *MNRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MN.Contract.MNTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MN *MNRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MN.Contract.MNTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MN *MNCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MN.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MN *MNTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MN.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MN *MNTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MN.Contract.contract.Transact(opts, method, params...)
}

// CheckDeposit is a free data retrieval call binding the contract method 0x4a7e60f0.
//
// Solidity: function checkDeposit(account address, amount uint256, h uint256) constant returns(bool)
func (_MN *MNCaller) CheckDeposit(opts *bind.CallOpts, account common.Address, amount *big.Int, h *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MN.contract.Call(opts, out, "checkDeposit", account, amount, h)
	return *ret0, err
}

// CheckDeposit is a free data retrieval call binding the contract method 0x4a7e60f0.
//
// Solidity: function checkDeposit(account address, amount uint256, h uint256) constant returns(bool)
func (_MN *MNSession) CheckDeposit(account common.Address, amount *big.Int, h *big.Int) (bool, error) {
	return _MN.Contract.CheckDeposit(&_MN.CallOpts, account, amount, h)
}

// CheckDeposit is a free data retrieval call binding the contract method 0x4a7e60f0.
//
// Solidity: function checkDeposit(account address, amount uint256, h uint256) constant returns(bool)
func (_MN *MNCallerSession) CheckDeposit(account common.Address, amount *big.Int, h *big.Int) (bool, error) {
	return _MN.Contract.CheckDeposit(&_MN.CallOpts, account, amount, h)
}

// CirculatingSupply is a free data retrieval call binding the contract method 0x92ff6aea.
//
// Solidity: function circulatingSupply(h uint256) constant returns(uint256)
func (_MN *MNCaller) CirculatingSupply(opts *bind.CallOpts, h *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MN.contract.Call(opts, out, "circulatingSupply", h)
	return *ret0, err
}

// CirculatingSupply is a free data retrieval call binding the contract method 0x92ff6aea.
//
// Solidity: function circulatingSupply(h uint256) constant returns(uint256)
func (_MN *MNSession) CirculatingSupply(h *big.Int) (*big.Int, error) {
	return _MN.Contract.CirculatingSupply(&_MN.CallOpts, h)
}

// CirculatingSupply is a free data retrieval call binding the contract method 0x92ff6aea.
//
// Solidity: function circulatingSupply(h uint256) constant returns(uint256)
func (_MN *MNCallerSession) CirculatingSupply(h *big.Int) (*big.Int, error) {
	return _MN.Contract.CirculatingSupply(&_MN.CallOpts, h)
}

// ItList is a free data retrieval call binding the contract method 0x649862ee.
//
// Solidity: function itList( uint256) constant returns(address)
func (_MN *MNCaller) ItList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MN.contract.Call(opts, out, "itList", arg0)
	return *ret0, err
}

// ItList is a free data retrieval call binding the contract method 0x649862ee.
//
// Solidity: function itList( uint256) constant returns(address)
func (_MN *MNSession) ItList(arg0 *big.Int) (common.Address, error) {
	return _MN.Contract.ItList(&_MN.CallOpts, arg0)
}

// ItList is a free data retrieval call binding the contract method 0x649862ee.
//
// Solidity: function itList( uint256) constant returns(address)
func (_MN *MNCallerSession) ItList(arg0 *big.Int) (common.Address, error) {
	return _MN.Contract.ItList(&_MN.CallOpts, arg0)
}

// NodeList is a free data retrieval call binding the contract method 0x37befda4.
//
// Solidity: function nodeList( address) constant returns(balance uint256, isActive bool)
func (_MN *MNCaller) NodeList(opts *bind.CallOpts, arg0 common.Address) (struct {
	Balance  *big.Int
	IsActive bool
}, error) {
	ret := new(struct {
		Balance  *big.Int
		IsActive bool
	})
	out := ret
	err := _MN.contract.Call(opts, out, "nodeList", arg0)
	return *ret, err
}

// NodeList is a free data retrieval call binding the contract method 0x37befda4.
//
// Solidity: function nodeList( address) constant returns(balance uint256, isActive bool)
func (_MN *MNSession) NodeList(arg0 common.Address) (struct {
	Balance  *big.Int
	IsActive bool
}, error) {
	return _MN.Contract.NodeList(&_MN.CallOpts, arg0)
}

// NodeList is a free data retrieval call binding the contract method 0x37befda4.
//
// Solidity: function nodeList( address) constant returns(balance uint256, isActive bool)
func (_MN *MNCallerSession) NodeList(arg0 common.Address) (struct {
	Balance  *big.Int
	IsActive bool
}, error) {
	return _MN.Contract.NodeList(&_MN.CallOpts, arg0)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(amount uint256) returns()
func (_MN *MNTransactor) Deposit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _MN.contract.Transact(opts, "deposit", amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(amount uint256) returns()
func (_MN *MNSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _MN.Contract.Deposit(&_MN.TransactOpts, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(amount uint256) returns()
func (_MN *MNTransactorSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _MN.Contract.Deposit(&_MN.TransactOpts, amount)
}

// UpdateStatus is a paid mutator transaction binding the contract method 0x5f60f485.
//
// Solidity: function updateStatus() returns()
func (_MN *MNTransactor) UpdateStatus(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MN.contract.Transact(opts, "updateStatus")
}

// UpdateStatus is a paid mutator transaction binding the contract method 0x5f60f485.
//
// Solidity: function updateStatus() returns()
func (_MN *MNSession) UpdateStatus() (*types.Transaction, error) {
	return _MN.Contract.UpdateStatus(&_MN.TransactOpts)
}

// UpdateStatus is a paid mutator transaction binding the contract method 0x5f60f485.
//
// Solidity: function updateStatus() returns()
func (_MN *MNTransactorSession) UpdateStatus() (*types.Transaction, error) {
	return _MN.Contract.UpdateStatus(&_MN.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_MN *MNTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MN.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_MN *MNSession) Withdraw() (*types.Transaction, error) {
	return _MN.Contract.Withdraw(&_MN.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_MN *MNTransactorSession) Withdraw() (*types.Transaction, error) {
	return _MN.Contract.Withdraw(&_MN.TransactOpts)
}
