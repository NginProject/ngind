// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package masternode

import (
	"math/big"
	"strings"

	"github.com/NginProject/ngind/accounts/abi"
	"github.com/NginProject/ngind/accounts/abi/bind"
	"github.com/NginProject/ngind/common"
	"github.com/NginProject/ngind/core/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
)

// MNABI is the input ABI used to generate the binding from.
const MNABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"nodeList\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"},{\"name\":\"isActive\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"addrNum\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"checkDeposit\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"addrList\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"h\",\"type\":\"uint256\"}],\"name\":\"circulatingSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"id2Node\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]"

// MNBin is the compiled bytecode used for deploying new contracts.
const MNBin = `0x608060405234801561001057600080fd5b506000600155610578806100256000396000f3fe608060405260043610610098576000357c010000000000000000000000000000000000000000000000000000000090048063804025641161006b578063804025641461016f57806392ff6aea146101b5578063cddb0c04146101df578063d0e30db01461023357610098565b806337befda41461009a5780633ccfd60b146100e65780636249fa96146100fb5780636fd3627614610122575b005b3480156100a657600080fd5b506100cd600480360360208110156100bd57600080fd5b5035600160a060020a031661023b565b6040805192835290151560208301528051918290030190f35b3480156100f257600080fd5b50610098610257565b34801561010757600080fd5b506101106102a8565b60408051918252519081900360200190f35b34801561012e57600080fd5b5061015b6004803603604081101561014557600080fd5b50600160a060020a0381351690602001356102ae565b604080519115158252519081900360200190f35b34801561017b57600080fd5b506101996004803603602081101561019257600080fd5b50356102e8565b60408051600160a060020a039092168252519081900360200190f35b3480156101c157600080fd5b50610110600480360360208110156101d857600080fd5b5035610310565b3480156101eb57600080fd5b506102096004803603602081101561020257600080fd5b5035610325565b60408051600160a060020a0390941684526020840192909252151582820152519081900360600190f35b610098610436565b6002602052600090815260409020805460019091015460ff1682565b3360008181526002602052604080822080548382556001909101805460ff19169055905190929183156108fc02918491818181858888f193505050501580156102a4573d6000803e3d6000fd5b5050565b60015481565b6000438111156102bd57600080fd5b50600160a060020a0391909116600090815260026020526040902054670de0b6b3a764000091011190565b60008054829081106102f657fe5b600091825260209091200154600160a060020a0316905081565b620186a0900469d3c21bcecceda10000000290565b600080600061038e60008581548110151561033c57fe5b60009182526020822001548154600160a060020a039091169160029181908990811061036457fe5b6000918252602080832090910154600160a060020a031683528201929092526040019020546102ae565b50600080548590811061039d57fe5b60009182526020822001548154600160a060020a03909116916002918190889081106103c557fe5b6000918252602080832090910154600160a060020a03168352820192909252604001812054815490916002918190899081106103fd57fe5b6000918252602080832090910154600160a060020a03168352820192909252604001902060010154919450925060ff1690509193909250565b33600090815260026020526040902054348101101561045457600080fd5b33600090815260026020526040902054670de0b6b3a7640000349091011161047b57600080fd5b3360009081526002602052604090205415156104e55760008054600181810183559180527f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e56301805473ffffffffffffffffffffffffffffffffffffffff1916331790558054810190555b336000818152600260205260409020805434019081905561050691906102ae565b1561052f573360009081526002602052604090206001908101805460ff1916909117905561054a565b336000908152600260205260409020600101805460ff191690555b56fea165627a7a72305820adfaca5133c591546fdce64d85c6c7e50dad1643bdce3970b338ee59495c210e0029`

// DeployMN deploys a new Ngin contract, binding an instance of MN to it.
func DeployMN(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MN, error) {
	parsed, err := abi.JSON(strings.NewReader(MNABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MNBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MN{MNCaller: MNCaller{contract: contract}, MNTransactor: MNTransactor{contract: contract}, MNFilterer: MNFilterer{contract: contract}}, nil
}

// MN is an auto generated Go binding around a Ngin contract.
type MN struct {
	MNCaller     // Read-only binding to the contract
	MNTransactor // Write-only binding to the contract
	MNFilterer   // Log filterer for contract events
}

// MNCaller is an auto generated read-only Go binding around a Ngin contract.
type MNCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MNTransactor is an auto generated write-only Go binding around a Ngin contract.
type MNTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MNFilterer is an auto generated log filtering Go binding around a Ngin contract events.
type MNFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MNSession is an auto generated Go binding around a Ngin contract,
// with pre-set call and transact options.
type MNSession struct {
	Contract     *MN               // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MNCallerSession is an auto generated read-only Go binding around a Ngin contract,
// with pre-set call options.
type MNCallerSession struct {
	Contract *MNCaller     // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MNTransactorSession is an auto generated write-only Go binding around a Ngin contract,
// with pre-set transact options.
type MNTransactorSession struct {
	Contract     *MNTransactor     // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MNRaw is an auto generated low-level Go binding around a Ngin contract.
type MNRaw struct {
	Contract *MN // Generic contract binding to access the raw methods on
}

// MNCallerRaw is an auto generated low-level read-only Go binding around a Ngin contract.
type MNCallerRaw struct {
	Contract *MNCaller // Generic read-only contract binding to access the raw methods on
}

// MNTransactorRaw is an auto generated low-level write-only Go binding around a Ngin contract.
type MNTransactorRaw struct {
	Contract *MNTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMN creates a new instance of MN, bound to a specific deployed contract.
func NewMN(address common.Address, backend bind.ContractBackend) (*MN, error) {
	contract, err := bindMN(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MN{MNCaller: MNCaller{contract: contract}, MNTransactor: MNTransactor{contract: contract}, MNFilterer: MNFilterer{contract: contract}}, nil
}

// NewMNCaller creates a new read-only instance of MN, bound to a specific deployed contract.
func NewMNCaller(address common.Address, caller bind.ContractCaller) (*MNCaller, error) {
	contract, err := bindMN(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MNCaller{contract: contract}, nil
}

// NewMNTransactor creates a new write-only instance of MN, bound to a specific deployed contract.
func NewMNTransactor(address common.Address, transactor bind.ContractTransactor) (*MNTransactor, error) {
	contract, err := bindMN(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MNTransactor{contract: contract}, nil
}

// NewMNFilterer creates a new log filterer instance of MN, bound to a specific deployed contract.
func NewMNFilterer(address common.Address, filterer bind.ContractFilterer) (*MNFilterer, error) {
	contract, err := bindMN(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MNFilterer{contract: contract}, nil
}

// bindMN binds a generic wrapper to an already deployed contract.
func bindMN(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MNABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// AddrList is a free data retrieval call binding the contract method 0x80402564.
//
// Solidity: function addrList( uint256) constant returns(address)
func (_MN *MNCaller) AddrList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MN.contract.Call(opts, out, "addrList", arg0)
	return *ret0, err
}

// AddrList is a free data retrieval call binding the contract method 0x80402564.
//
// Solidity: function addrList( uint256) constant returns(address)
func (_MN *MNSession) AddrList(arg0 *big.Int) (common.Address, error) {
	return _MN.Contract.AddrList(&_MN.CallOpts, arg0)
}

// AddrList is a free data retrieval call binding the contract method 0x80402564.
//
// Solidity: function addrList( uint256) constant returns(address)
func (_MN *MNCallerSession) AddrList(arg0 *big.Int) (common.Address, error) {
	return _MN.Contract.AddrList(&_MN.CallOpts, arg0)
}

// AddrNum is a free data retrieval call binding the contract method 0x6249fa96.
//
// Solidity: function addrNum() constant returns(uint256)
func (_MN *MNCaller) AddrNum(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MN.contract.Call(opts, out, "addrNum")
	return *ret0, err
}

// AddrNum is a free data retrieval call binding the contract method 0x6249fa96.
//
// Solidity: function addrNum() constant returns(uint256)
func (_MN *MNSession) AddrNum() (*big.Int, error) {
	return _MN.Contract.AddrNum(&_MN.CallOpts)
}

// AddrNum is a free data retrieval call binding the contract method 0x6249fa96.
//
// Solidity: function addrNum() constant returns(uint256)
func (_MN *MNCallerSession) AddrNum() (*big.Int, error) {
	return _MN.Contract.AddrNum(&_MN.CallOpts)
}

// CheckDeposit is a free data retrieval call binding the contract method 0x6fd36276.
//
// Solidity: function checkDeposit(addr address, amount uint256) constant returns(bool)
func (_MN *MNCaller) CheckDeposit(opts *bind.CallOpts, addr common.Address, amount *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MN.contract.Call(opts, out, "checkDeposit", addr, amount)
	return *ret0, err
}

// CheckDeposit is a free data retrieval call binding the contract method 0x6fd36276.
//
// Solidity: function checkDeposit(addr address, amount uint256) constant returns(bool)
func (_MN *MNSession) CheckDeposit(addr common.Address, amount *big.Int) (bool, error) {
	return _MN.Contract.CheckDeposit(&_MN.CallOpts, addr, amount)
}

// CheckDeposit is a free data retrieval call binding the contract method 0x6fd36276.
//
// Solidity: function checkDeposit(addr address, amount uint256) constant returns(bool)
func (_MN *MNCallerSession) CheckDeposit(addr common.Address, amount *big.Int) (bool, error) {
	return _MN.Contract.CheckDeposit(&_MN.CallOpts, addr, amount)
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

// Id2Node is a free data retrieval call binding the contract method 0xcddb0c04.
//
// Solidity: function id2Node(id uint256) constant returns(address, uint256, bool)
func (_MN *MNCaller) Id2Node(opts *bind.CallOpts, id *big.Int) (common.Address, *big.Int, bool, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(*big.Int)
		ret2 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _MN.contract.Call(opts, out, "id2Node", id)
	return *ret0, *ret1, *ret2, err
}

// Id2Node is a free data retrieval call binding the contract method 0xcddb0c04.
//
// Solidity: function id2Node(id uint256) constant returns(address, uint256, bool)
func (_MN *MNSession) Id2Node(id *big.Int) (common.Address, *big.Int, bool, error) {
	return _MN.Contract.Id2Node(&_MN.CallOpts, id)
}

// Id2Node is a free data retrieval call binding the contract method 0xcddb0c04.
//
// Solidity: function id2Node(id uint256) constant returns(address, uint256, bool)
func (_MN *MNCallerSession) Id2Node(id *big.Int) (common.Address, *big.Int, bool, error) {
	return _MN.Contract.Id2Node(&_MN.CallOpts, id)
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

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_MN *MNTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MN.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_MN *MNSession) Deposit() (*types.Transaction, error) {
	return _MN.Contract.Deposit(&_MN.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_MN *MNTransactorSession) Deposit() (*types.Transaction, error) {
	return _MN.Contract.Deposit(&_MN.TransactOpts)
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
