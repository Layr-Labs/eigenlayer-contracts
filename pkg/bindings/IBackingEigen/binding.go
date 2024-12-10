// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IBackingEigen

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// IBackingEigenMetaData contains all meta data concerning the IBackingEigen contract.
var IBackingEigenMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"CLOCK_MODE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"EIGEN\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"clock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"disableTransferRestrictions\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAllowedFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isAllowedFrom\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAllowedTo\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isAllowedTo\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setIsMinter\",\"inputs\":[{\"name\":\"minterAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"newStatus\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferRestrictionsDisabledAfter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// IBackingEigenABI is the input ABI used to generate the binding from.
// Deprecated: Use IBackingEigenMetaData.ABI instead.
var IBackingEigenABI = IBackingEigenMetaData.ABI

// IBackingEigen is an auto generated Go binding around an Ethereum contract.
type IBackingEigen struct {
	IBackingEigenCaller     // Read-only binding to the contract
	IBackingEigenTransactor // Write-only binding to the contract
	IBackingEigenFilterer   // Log filterer for contract events
}

// IBackingEigenCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBackingEigenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBackingEigenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBackingEigenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBackingEigenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBackingEigenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBackingEigenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBackingEigenSession struct {
	Contract     *IBackingEigen    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBackingEigenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBackingEigenCallerSession struct {
	Contract *IBackingEigenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IBackingEigenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBackingEigenTransactorSession struct {
	Contract     *IBackingEigenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IBackingEigenRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBackingEigenRaw struct {
	Contract *IBackingEigen // Generic contract binding to access the raw methods on
}

// IBackingEigenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBackingEigenCallerRaw struct {
	Contract *IBackingEigenCaller // Generic read-only contract binding to access the raw methods on
}

// IBackingEigenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBackingEigenTransactorRaw struct {
	Contract *IBackingEigenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBackingEigen creates a new instance of IBackingEigen, bound to a specific deployed contract.
func NewIBackingEigen(address common.Address, backend bind.ContractBackend) (*IBackingEigen, error) {
	contract, err := bindIBackingEigen(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBackingEigen{IBackingEigenCaller: IBackingEigenCaller{contract: contract}, IBackingEigenTransactor: IBackingEigenTransactor{contract: contract}, IBackingEigenFilterer: IBackingEigenFilterer{contract: contract}}, nil
}

// NewIBackingEigenCaller creates a new read-only instance of IBackingEigen, bound to a specific deployed contract.
func NewIBackingEigenCaller(address common.Address, caller bind.ContractCaller) (*IBackingEigenCaller, error) {
	contract, err := bindIBackingEigen(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBackingEigenCaller{contract: contract}, nil
}

// NewIBackingEigenTransactor creates a new write-only instance of IBackingEigen, bound to a specific deployed contract.
func NewIBackingEigenTransactor(address common.Address, transactor bind.ContractTransactor) (*IBackingEigenTransactor, error) {
	contract, err := bindIBackingEigen(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBackingEigenTransactor{contract: contract}, nil
}

// NewIBackingEigenFilterer creates a new log filterer instance of IBackingEigen, bound to a specific deployed contract.
func NewIBackingEigenFilterer(address common.Address, filterer bind.ContractFilterer) (*IBackingEigenFilterer, error) {
	contract, err := bindIBackingEigen(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBackingEigenFilterer{contract: contract}, nil
}

// bindIBackingEigen binds a generic wrapper to an already deployed contract.
func bindIBackingEigen(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IBackingEigenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBackingEigen *IBackingEigenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBackingEigen.Contract.IBackingEigenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBackingEigen *IBackingEigenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBackingEigen.Contract.IBackingEigenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBackingEigen *IBackingEigenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBackingEigen.Contract.IBackingEigenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBackingEigen *IBackingEigenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBackingEigen.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBackingEigen *IBackingEigenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBackingEigen.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBackingEigen *IBackingEigenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBackingEigen.Contract.contract.Transact(opts, method, params...)
}

// CLOCKMODE is a free data retrieval call binding the contract method 0x4bf5d7e9.
//
// Solidity: function CLOCK_MODE() pure returns(string)
func (_IBackingEigen *IBackingEigenCaller) CLOCKMODE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IBackingEigen.contract.Call(opts, &out, "CLOCK_MODE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CLOCKMODE is a free data retrieval call binding the contract method 0x4bf5d7e9.
//
// Solidity: function CLOCK_MODE() pure returns(string)
func (_IBackingEigen *IBackingEigenSession) CLOCKMODE() (string, error) {
	return _IBackingEigen.Contract.CLOCKMODE(&_IBackingEigen.CallOpts)
}

// CLOCKMODE is a free data retrieval call binding the contract method 0x4bf5d7e9.
//
// Solidity: function CLOCK_MODE() pure returns(string)
func (_IBackingEigen *IBackingEigenCallerSession) CLOCKMODE() (string, error) {
	return _IBackingEigen.Contract.CLOCKMODE(&_IBackingEigen.CallOpts)
}

// EIGEN is a free data retrieval call binding the contract method 0xfdc371ce.
//
// Solidity: function EIGEN() view returns(address)
func (_IBackingEigen *IBackingEigenCaller) EIGEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IBackingEigen.contract.Call(opts, &out, "EIGEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EIGEN is a free data retrieval call binding the contract method 0xfdc371ce.
//
// Solidity: function EIGEN() view returns(address)
func (_IBackingEigen *IBackingEigenSession) EIGEN() (common.Address, error) {
	return _IBackingEigen.Contract.EIGEN(&_IBackingEigen.CallOpts)
}

// EIGEN is a free data retrieval call binding the contract method 0xfdc371ce.
//
// Solidity: function EIGEN() view returns(address)
func (_IBackingEigen *IBackingEigenCallerSession) EIGEN() (common.Address, error) {
	return _IBackingEigen.Contract.EIGEN(&_IBackingEigen.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IBackingEigen *IBackingEigenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IBackingEigen.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IBackingEigen *IBackingEigenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IBackingEigen.Contract.Allowance(&_IBackingEigen.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IBackingEigen *IBackingEigenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IBackingEigen.Contract.Allowance(&_IBackingEigen.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IBackingEigen *IBackingEigenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IBackingEigen.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IBackingEigen *IBackingEigenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IBackingEigen.Contract.BalanceOf(&_IBackingEigen.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IBackingEigen *IBackingEigenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IBackingEigen.Contract.BalanceOf(&_IBackingEigen.CallOpts, account)
}

// Clock is a free data retrieval call binding the contract method 0x91ddadf4.
//
// Solidity: function clock() view returns(uint48)
func (_IBackingEigen *IBackingEigenCaller) Clock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IBackingEigen.contract.Call(opts, &out, "clock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Clock is a free data retrieval call binding the contract method 0x91ddadf4.
//
// Solidity: function clock() view returns(uint48)
func (_IBackingEigen *IBackingEigenSession) Clock() (*big.Int, error) {
	return _IBackingEigen.Contract.Clock(&_IBackingEigen.CallOpts)
}

// Clock is a free data retrieval call binding the contract method 0x91ddadf4.
//
// Solidity: function clock() view returns(uint48)
func (_IBackingEigen *IBackingEigenCallerSession) Clock() (*big.Int, error) {
	return _IBackingEigen.Contract.Clock(&_IBackingEigen.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IBackingEigen *IBackingEigenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IBackingEigen.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IBackingEigen *IBackingEigenSession) TotalSupply() (*big.Int, error) {
	return _IBackingEigen.Contract.TotalSupply(&_IBackingEigen.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IBackingEigen *IBackingEigenCallerSession) TotalSupply() (*big.Int, error) {
	return _IBackingEigen.Contract.TotalSupply(&_IBackingEigen.CallOpts)
}

// TransferRestrictionsDisabledAfter is a free data retrieval call binding the contract method 0x9aec4bae.
//
// Solidity: function transferRestrictionsDisabledAfter() view returns(uint256)
func (_IBackingEigen *IBackingEigenCaller) TransferRestrictionsDisabledAfter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IBackingEigen.contract.Call(opts, &out, "transferRestrictionsDisabledAfter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TransferRestrictionsDisabledAfter is a free data retrieval call binding the contract method 0x9aec4bae.
//
// Solidity: function transferRestrictionsDisabledAfter() view returns(uint256)
func (_IBackingEigen *IBackingEigenSession) TransferRestrictionsDisabledAfter() (*big.Int, error) {
	return _IBackingEigen.Contract.TransferRestrictionsDisabledAfter(&_IBackingEigen.CallOpts)
}

// TransferRestrictionsDisabledAfter is a free data retrieval call binding the contract method 0x9aec4bae.
//
// Solidity: function transferRestrictionsDisabledAfter() view returns(uint256)
func (_IBackingEigen *IBackingEigenCallerSession) TransferRestrictionsDisabledAfter() (*big.Int, error) {
	return _IBackingEigen.Contract.TransferRestrictionsDisabledAfter(&_IBackingEigen.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IBackingEigen *IBackingEigenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IBackingEigen.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IBackingEigen *IBackingEigenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IBackingEigen.Contract.Approve(&_IBackingEigen.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IBackingEigen *IBackingEigenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IBackingEigen.Contract.Approve(&_IBackingEigen.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_IBackingEigen *IBackingEigenTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _IBackingEigen.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_IBackingEigen *IBackingEigenSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _IBackingEigen.Contract.Burn(&_IBackingEigen.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_IBackingEigen *IBackingEigenTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _IBackingEigen.Contract.Burn(&_IBackingEigen.TransactOpts, amount)
}

// DisableTransferRestrictions is a paid mutator transaction binding the contract method 0xeb415f45.
//
// Solidity: function disableTransferRestrictions() returns()
func (_IBackingEigen *IBackingEigenTransactor) DisableTransferRestrictions(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBackingEigen.contract.Transact(opts, "disableTransferRestrictions")
}

// DisableTransferRestrictions is a paid mutator transaction binding the contract method 0xeb415f45.
//
// Solidity: function disableTransferRestrictions() returns()
func (_IBackingEigen *IBackingEigenSession) DisableTransferRestrictions() (*types.Transaction, error) {
	return _IBackingEigen.Contract.DisableTransferRestrictions(&_IBackingEigen.TransactOpts)
}

// DisableTransferRestrictions is a paid mutator transaction binding the contract method 0xeb415f45.
//
// Solidity: function disableTransferRestrictions() returns()
func (_IBackingEigen *IBackingEigenTransactorSession) DisableTransferRestrictions() (*types.Transaction, error) {
	return _IBackingEigen.Contract.DisableTransferRestrictions(&_IBackingEigen.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_IBackingEigen *IBackingEigenTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address) (*types.Transaction, error) {
	return _IBackingEigen.contract.Transact(opts, "initialize", initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_IBackingEigen *IBackingEigenSession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _IBackingEigen.Contract.Initialize(&_IBackingEigen.TransactOpts, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_IBackingEigen *IBackingEigenTransactorSession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _IBackingEigen.Contract.Initialize(&_IBackingEigen.TransactOpts, initialOwner)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_IBackingEigen *IBackingEigenTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IBackingEigen.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_IBackingEigen *IBackingEigenSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IBackingEigen.Contract.Mint(&_IBackingEigen.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_IBackingEigen *IBackingEigenTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IBackingEigen.Contract.Mint(&_IBackingEigen.TransactOpts, to, amount)
}

// SetAllowedFrom is a paid mutator transaction binding the contract method 0x1ffacdef.
//
// Solidity: function setAllowedFrom(address from, bool isAllowedFrom) returns()
func (_IBackingEigen *IBackingEigenTransactor) SetAllowedFrom(opts *bind.TransactOpts, from common.Address, isAllowedFrom bool) (*types.Transaction, error) {
	return _IBackingEigen.contract.Transact(opts, "setAllowedFrom", from, isAllowedFrom)
}

// SetAllowedFrom is a paid mutator transaction binding the contract method 0x1ffacdef.
//
// Solidity: function setAllowedFrom(address from, bool isAllowedFrom) returns()
func (_IBackingEigen *IBackingEigenSession) SetAllowedFrom(from common.Address, isAllowedFrom bool) (*types.Transaction, error) {
	return _IBackingEigen.Contract.SetAllowedFrom(&_IBackingEigen.TransactOpts, from, isAllowedFrom)
}

// SetAllowedFrom is a paid mutator transaction binding the contract method 0x1ffacdef.
//
// Solidity: function setAllowedFrom(address from, bool isAllowedFrom) returns()
func (_IBackingEigen *IBackingEigenTransactorSession) SetAllowedFrom(from common.Address, isAllowedFrom bool) (*types.Transaction, error) {
	return _IBackingEigen.Contract.SetAllowedFrom(&_IBackingEigen.TransactOpts, from, isAllowedFrom)
}

// SetAllowedTo is a paid mutator transaction binding the contract method 0xb8c25594.
//
// Solidity: function setAllowedTo(address to, bool isAllowedTo) returns()
func (_IBackingEigen *IBackingEigenTransactor) SetAllowedTo(opts *bind.TransactOpts, to common.Address, isAllowedTo bool) (*types.Transaction, error) {
	return _IBackingEigen.contract.Transact(opts, "setAllowedTo", to, isAllowedTo)
}

// SetAllowedTo is a paid mutator transaction binding the contract method 0xb8c25594.
//
// Solidity: function setAllowedTo(address to, bool isAllowedTo) returns()
func (_IBackingEigen *IBackingEigenSession) SetAllowedTo(to common.Address, isAllowedTo bool) (*types.Transaction, error) {
	return _IBackingEigen.Contract.SetAllowedTo(&_IBackingEigen.TransactOpts, to, isAllowedTo)
}

// SetAllowedTo is a paid mutator transaction binding the contract method 0xb8c25594.
//
// Solidity: function setAllowedTo(address to, bool isAllowedTo) returns()
func (_IBackingEigen *IBackingEigenTransactorSession) SetAllowedTo(to common.Address, isAllowedTo bool) (*types.Transaction, error) {
	return _IBackingEigen.Contract.SetAllowedTo(&_IBackingEigen.TransactOpts, to, isAllowedTo)
}

// SetIsMinter is a paid mutator transaction binding the contract method 0x66eb399f.
//
// Solidity: function setIsMinter(address minterAddress, bool newStatus) returns()
func (_IBackingEigen *IBackingEigenTransactor) SetIsMinter(opts *bind.TransactOpts, minterAddress common.Address, newStatus bool) (*types.Transaction, error) {
	return _IBackingEigen.contract.Transact(opts, "setIsMinter", minterAddress, newStatus)
}

// SetIsMinter is a paid mutator transaction binding the contract method 0x66eb399f.
//
// Solidity: function setIsMinter(address minterAddress, bool newStatus) returns()
func (_IBackingEigen *IBackingEigenSession) SetIsMinter(minterAddress common.Address, newStatus bool) (*types.Transaction, error) {
	return _IBackingEigen.Contract.SetIsMinter(&_IBackingEigen.TransactOpts, minterAddress, newStatus)
}

// SetIsMinter is a paid mutator transaction binding the contract method 0x66eb399f.
//
// Solidity: function setIsMinter(address minterAddress, bool newStatus) returns()
func (_IBackingEigen *IBackingEigenTransactorSession) SetIsMinter(minterAddress common.Address, newStatus bool) (*types.Transaction, error) {
	return _IBackingEigen.Contract.SetIsMinter(&_IBackingEigen.TransactOpts, minterAddress, newStatus)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_IBackingEigen *IBackingEigenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IBackingEigen.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_IBackingEigen *IBackingEigenSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IBackingEigen.Contract.Transfer(&_IBackingEigen.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_IBackingEigen *IBackingEigenTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IBackingEigen.Contract.Transfer(&_IBackingEigen.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_IBackingEigen *IBackingEigenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IBackingEigen.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_IBackingEigen *IBackingEigenSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IBackingEigen.Contract.TransferFrom(&_IBackingEigen.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_IBackingEigen *IBackingEigenTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IBackingEigen.Contract.TransferFrom(&_IBackingEigen.TransactOpts, from, to, amount)
}

// IBackingEigenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IBackingEigen contract.
type IBackingEigenApprovalIterator struct {
	Event *IBackingEigenApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IBackingEigenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBackingEigenApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IBackingEigenApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IBackingEigenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBackingEigenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBackingEigenApproval represents a Approval event raised by the IBackingEigen contract.
type IBackingEigenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IBackingEigen *IBackingEigenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IBackingEigenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IBackingEigen.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IBackingEigenApprovalIterator{contract: _IBackingEigen.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IBackingEigen *IBackingEigenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IBackingEigenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IBackingEigen.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBackingEigenApproval)
				if err := _IBackingEigen.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IBackingEigen *IBackingEigenFilterer) ParseApproval(log types.Log) (*IBackingEigenApproval, error) {
	event := new(IBackingEigenApproval)
	if err := _IBackingEigen.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBackingEigenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IBackingEigen contract.
type IBackingEigenTransferIterator struct {
	Event *IBackingEigenTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IBackingEigenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBackingEigenTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IBackingEigenTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IBackingEigenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBackingEigenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBackingEigenTransfer represents a Transfer event raised by the IBackingEigen contract.
type IBackingEigenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IBackingEigen *IBackingEigenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IBackingEigenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IBackingEigen.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IBackingEigenTransferIterator{contract: _IBackingEigen.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IBackingEigen *IBackingEigenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IBackingEigenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IBackingEigen.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBackingEigenTransfer)
				if err := _IBackingEigen.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IBackingEigen *IBackingEigenFilterer) ParseTransfer(log types.Log) (*IBackingEigenTransfer, error) {
	event := new(IBackingEigenTransfer)
	if err := _IBackingEigen.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
