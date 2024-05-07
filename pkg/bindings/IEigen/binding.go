// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IEigen

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

// IEigenMetaData contains all meta data concerning the IEigen contract.
var IEigenMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"CLOCK_MODE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"clock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"disableTransferRestrictions\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAllowedFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isAllowedFrom\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAllowedTo\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isAllowedTo\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unwrap\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"wrap\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// IEigenABI is the input ABI used to generate the binding from.
// Deprecated: Use IEigenMetaData.ABI instead.
var IEigenABI = IEigenMetaData.ABI

// IEigen is an auto generated Go binding around an Ethereum contract.
type IEigen struct {
	IEigenCaller     // Read-only binding to the contract
	IEigenTransactor // Write-only binding to the contract
	IEigenFilterer   // Log filterer for contract events
}

// IEigenCaller is an auto generated read-only Go binding around an Ethereum contract.
type IEigenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEigenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IEigenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEigenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IEigenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEigenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IEigenSession struct {
	Contract     *IEigen           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IEigenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IEigenCallerSession struct {
	Contract *IEigenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IEigenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IEigenTransactorSession struct {
	Contract     *IEigenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IEigenRaw is an auto generated low-level Go binding around an Ethereum contract.
type IEigenRaw struct {
	Contract *IEigen // Generic contract binding to access the raw methods on
}

// IEigenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IEigenCallerRaw struct {
	Contract *IEigenCaller // Generic read-only contract binding to access the raw methods on
}

// IEigenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IEigenTransactorRaw struct {
	Contract *IEigenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIEigen creates a new instance of IEigen, bound to a specific deployed contract.
func NewIEigen(address common.Address, backend bind.ContractBackend) (*IEigen, error) {
	contract, err := bindIEigen(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IEigen{IEigenCaller: IEigenCaller{contract: contract}, IEigenTransactor: IEigenTransactor{contract: contract}, IEigenFilterer: IEigenFilterer{contract: contract}}, nil
}

// NewIEigenCaller creates a new read-only instance of IEigen, bound to a specific deployed contract.
func NewIEigenCaller(address common.Address, caller bind.ContractCaller) (*IEigenCaller, error) {
	contract, err := bindIEigen(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IEigenCaller{contract: contract}, nil
}

// NewIEigenTransactor creates a new write-only instance of IEigen, bound to a specific deployed contract.
func NewIEigenTransactor(address common.Address, transactor bind.ContractTransactor) (*IEigenTransactor, error) {
	contract, err := bindIEigen(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IEigenTransactor{contract: contract}, nil
}

// NewIEigenFilterer creates a new log filterer instance of IEigen, bound to a specific deployed contract.
func NewIEigenFilterer(address common.Address, filterer bind.ContractFilterer) (*IEigenFilterer, error) {
	contract, err := bindIEigen(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IEigenFilterer{contract: contract}, nil
}

// bindIEigen binds a generic wrapper to an already deployed contract.
func bindIEigen(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IEigenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEigen *IEigenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IEigen.Contract.IEigenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEigen *IEigenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEigen.Contract.IEigenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEigen *IEigenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEigen.Contract.IEigenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEigen *IEigenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IEigen.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEigen *IEigenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEigen.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEigen *IEigenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEigen.Contract.contract.Transact(opts, method, params...)
}

// CLOCKMODE is a free data retrieval call binding the contract method 0x4bf5d7e9.
//
// Solidity: function CLOCK_MODE() pure returns(string)
func (_IEigen *IEigenCaller) CLOCKMODE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IEigen.contract.Call(opts, &out, "CLOCK_MODE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CLOCKMODE is a free data retrieval call binding the contract method 0x4bf5d7e9.
//
// Solidity: function CLOCK_MODE() pure returns(string)
func (_IEigen *IEigenSession) CLOCKMODE() (string, error) {
	return _IEigen.Contract.CLOCKMODE(&_IEigen.CallOpts)
}

// CLOCKMODE is a free data retrieval call binding the contract method 0x4bf5d7e9.
//
// Solidity: function CLOCK_MODE() pure returns(string)
func (_IEigen *IEigenCallerSession) CLOCKMODE() (string, error) {
	return _IEigen.Contract.CLOCKMODE(&_IEigen.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IEigen *IEigenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IEigen.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IEigen *IEigenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IEigen.Contract.Allowance(&_IEigen.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IEigen *IEigenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IEigen.Contract.Allowance(&_IEigen.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IEigen *IEigenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IEigen.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IEigen *IEigenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IEigen.Contract.BalanceOf(&_IEigen.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IEigen *IEigenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IEigen.Contract.BalanceOf(&_IEigen.CallOpts, account)
}

// Clock is a free data retrieval call binding the contract method 0x91ddadf4.
//
// Solidity: function clock() view returns(uint48)
func (_IEigen *IEigenCaller) Clock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IEigen.contract.Call(opts, &out, "clock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Clock is a free data retrieval call binding the contract method 0x91ddadf4.
//
// Solidity: function clock() view returns(uint48)
func (_IEigen *IEigenSession) Clock() (*big.Int, error) {
	return _IEigen.Contract.Clock(&_IEigen.CallOpts)
}

// Clock is a free data retrieval call binding the contract method 0x91ddadf4.
//
// Solidity: function clock() view returns(uint48)
func (_IEigen *IEigenCallerSession) Clock() (*big.Int, error) {
	return _IEigen.Contract.Clock(&_IEigen.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IEigen *IEigenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IEigen.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IEigen *IEigenSession) TotalSupply() (*big.Int, error) {
	return _IEigen.Contract.TotalSupply(&_IEigen.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IEigen *IEigenCallerSession) TotalSupply() (*big.Int, error) {
	return _IEigen.Contract.TotalSupply(&_IEigen.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IEigen *IEigenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IEigen.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IEigen *IEigenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IEigen.Contract.Approve(&_IEigen.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IEigen *IEigenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IEigen.Contract.Approve(&_IEigen.TransactOpts, spender, amount)
}

// DisableTransferRestrictions is a paid mutator transaction binding the contract method 0xeb415f45.
//
// Solidity: function disableTransferRestrictions() returns()
func (_IEigen *IEigenTransactor) DisableTransferRestrictions(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEigen.contract.Transact(opts, "disableTransferRestrictions")
}

// DisableTransferRestrictions is a paid mutator transaction binding the contract method 0xeb415f45.
//
// Solidity: function disableTransferRestrictions() returns()
func (_IEigen *IEigenSession) DisableTransferRestrictions() (*types.Transaction, error) {
	return _IEigen.Contract.DisableTransferRestrictions(&_IEigen.TransactOpts)
}

// DisableTransferRestrictions is a paid mutator transaction binding the contract method 0xeb415f45.
//
// Solidity: function disableTransferRestrictions() returns()
func (_IEigen *IEigenTransactorSession) DisableTransferRestrictions() (*types.Transaction, error) {
	return _IEigen.Contract.DisableTransferRestrictions(&_IEigen.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() returns()
func (_IEigen *IEigenTransactor) Mint(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEigen.contract.Transact(opts, "mint")
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() returns()
func (_IEigen *IEigenSession) Mint() (*types.Transaction, error) {
	return _IEigen.Contract.Mint(&_IEigen.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() returns()
func (_IEigen *IEigenTransactorSession) Mint() (*types.Transaction, error) {
	return _IEigen.Contract.Mint(&_IEigen.TransactOpts)
}

// SetAllowedFrom is a paid mutator transaction binding the contract method 0x1ffacdef.
//
// Solidity: function setAllowedFrom(address from, bool isAllowedFrom) returns()
func (_IEigen *IEigenTransactor) SetAllowedFrom(opts *bind.TransactOpts, from common.Address, isAllowedFrom bool) (*types.Transaction, error) {
	return _IEigen.contract.Transact(opts, "setAllowedFrom", from, isAllowedFrom)
}

// SetAllowedFrom is a paid mutator transaction binding the contract method 0x1ffacdef.
//
// Solidity: function setAllowedFrom(address from, bool isAllowedFrom) returns()
func (_IEigen *IEigenSession) SetAllowedFrom(from common.Address, isAllowedFrom bool) (*types.Transaction, error) {
	return _IEigen.Contract.SetAllowedFrom(&_IEigen.TransactOpts, from, isAllowedFrom)
}

// SetAllowedFrom is a paid mutator transaction binding the contract method 0x1ffacdef.
//
// Solidity: function setAllowedFrom(address from, bool isAllowedFrom) returns()
func (_IEigen *IEigenTransactorSession) SetAllowedFrom(from common.Address, isAllowedFrom bool) (*types.Transaction, error) {
	return _IEigen.Contract.SetAllowedFrom(&_IEigen.TransactOpts, from, isAllowedFrom)
}

// SetAllowedTo is a paid mutator transaction binding the contract method 0xb8c25594.
//
// Solidity: function setAllowedTo(address to, bool isAllowedTo) returns()
func (_IEigen *IEigenTransactor) SetAllowedTo(opts *bind.TransactOpts, to common.Address, isAllowedTo bool) (*types.Transaction, error) {
	return _IEigen.contract.Transact(opts, "setAllowedTo", to, isAllowedTo)
}

// SetAllowedTo is a paid mutator transaction binding the contract method 0xb8c25594.
//
// Solidity: function setAllowedTo(address to, bool isAllowedTo) returns()
func (_IEigen *IEigenSession) SetAllowedTo(to common.Address, isAllowedTo bool) (*types.Transaction, error) {
	return _IEigen.Contract.SetAllowedTo(&_IEigen.TransactOpts, to, isAllowedTo)
}

// SetAllowedTo is a paid mutator transaction binding the contract method 0xb8c25594.
//
// Solidity: function setAllowedTo(address to, bool isAllowedTo) returns()
func (_IEigen *IEigenTransactorSession) SetAllowedTo(to common.Address, isAllowedTo bool) (*types.Transaction, error) {
	return _IEigen.Contract.SetAllowedTo(&_IEigen.TransactOpts, to, isAllowedTo)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_IEigen *IEigenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IEigen.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_IEigen *IEigenSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IEigen.Contract.Transfer(&_IEigen.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_IEigen *IEigenTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IEigen.Contract.Transfer(&_IEigen.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_IEigen *IEigenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IEigen.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_IEigen *IEigenSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IEigen.Contract.TransferFrom(&_IEigen.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_IEigen *IEigenTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IEigen.Contract.TransferFrom(&_IEigen.TransactOpts, from, to, amount)
}

// Unwrap is a paid mutator transaction binding the contract method 0xde0e9a3e.
//
// Solidity: function unwrap(uint256 amount) returns()
func (_IEigen *IEigenTransactor) Unwrap(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _IEigen.contract.Transact(opts, "unwrap", amount)
}

// Unwrap is a paid mutator transaction binding the contract method 0xde0e9a3e.
//
// Solidity: function unwrap(uint256 amount) returns()
func (_IEigen *IEigenSession) Unwrap(amount *big.Int) (*types.Transaction, error) {
	return _IEigen.Contract.Unwrap(&_IEigen.TransactOpts, amount)
}

// Unwrap is a paid mutator transaction binding the contract method 0xde0e9a3e.
//
// Solidity: function unwrap(uint256 amount) returns()
func (_IEigen *IEigenTransactorSession) Unwrap(amount *big.Int) (*types.Transaction, error) {
	return _IEigen.Contract.Unwrap(&_IEigen.TransactOpts, amount)
}

// Wrap is a paid mutator transaction binding the contract method 0xea598cb0.
//
// Solidity: function wrap(uint256 amount) returns()
func (_IEigen *IEigenTransactor) Wrap(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _IEigen.contract.Transact(opts, "wrap", amount)
}

// Wrap is a paid mutator transaction binding the contract method 0xea598cb0.
//
// Solidity: function wrap(uint256 amount) returns()
func (_IEigen *IEigenSession) Wrap(amount *big.Int) (*types.Transaction, error) {
	return _IEigen.Contract.Wrap(&_IEigen.TransactOpts, amount)
}

// Wrap is a paid mutator transaction binding the contract method 0xea598cb0.
//
// Solidity: function wrap(uint256 amount) returns()
func (_IEigen *IEigenTransactorSession) Wrap(amount *big.Int) (*types.Transaction, error) {
	return _IEigen.Contract.Wrap(&_IEigen.TransactOpts, amount)
}

// IEigenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IEigen contract.
type IEigenApprovalIterator struct {
	Event *IEigenApproval // Event containing the contract specifics and raw log

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
func (it *IEigenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEigenApproval)
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
		it.Event = new(IEigenApproval)
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
func (it *IEigenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEigenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEigenApproval represents a Approval event raised by the IEigen contract.
type IEigenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IEigen *IEigenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IEigenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IEigen.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IEigenApprovalIterator{contract: _IEigen.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IEigen *IEigenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IEigenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IEigen.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEigenApproval)
				if err := _IEigen.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_IEigen *IEigenFilterer) ParseApproval(log types.Log) (*IEigenApproval, error) {
	event := new(IEigenApproval)
	if err := _IEigen.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IEigenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IEigen contract.
type IEigenTransferIterator struct {
	Event *IEigenTransfer // Event containing the contract specifics and raw log

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
func (it *IEigenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEigenTransfer)
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
		it.Event = new(IEigenTransfer)
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
func (it *IEigenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEigenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEigenTransfer represents a Transfer event raised by the IEigen contract.
type IEigenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IEigen *IEigenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IEigenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IEigen.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IEigenTransferIterator{contract: _IEigen.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IEigen *IEigenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IEigenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IEigen.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEigenTransfer)
				if err := _IEigen.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_IEigen *IEigenFilterer) ParseTransfer(log types.Log) (*IEigenTransfer, error) {
	event := new(IEigenTransfer)
	if err := _IEigen.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
