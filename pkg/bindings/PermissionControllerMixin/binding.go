// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package PermissionControllerMixin

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

// PermissionControllerMixinMetaData contains all meta data concerning the PermissionControllerMixin contract.
var PermissionControllerMixinMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"permissionController\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"}],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"InvalidPermissions\",\"inputs\":[]}]",
}

// PermissionControllerMixinABI is the input ABI used to generate the binding from.
// Deprecated: Use PermissionControllerMixinMetaData.ABI instead.
var PermissionControllerMixinABI = PermissionControllerMixinMetaData.ABI

// PermissionControllerMixin is an auto generated Go binding around an Ethereum contract.
type PermissionControllerMixin struct {
	PermissionControllerMixinCaller     // Read-only binding to the contract
	PermissionControllerMixinTransactor // Write-only binding to the contract
	PermissionControllerMixinFilterer   // Log filterer for contract events
}

// PermissionControllerMixinCaller is an auto generated read-only Go binding around an Ethereum contract.
type PermissionControllerMixinCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermissionControllerMixinTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PermissionControllerMixinTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermissionControllerMixinFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PermissionControllerMixinFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermissionControllerMixinSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PermissionControllerMixinSession struct {
	Contract     *PermissionControllerMixin // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// PermissionControllerMixinCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PermissionControllerMixinCallerSession struct {
	Contract *PermissionControllerMixinCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// PermissionControllerMixinTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PermissionControllerMixinTransactorSession struct {
	Contract     *PermissionControllerMixinTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// PermissionControllerMixinRaw is an auto generated low-level Go binding around an Ethereum contract.
type PermissionControllerMixinRaw struct {
	Contract *PermissionControllerMixin // Generic contract binding to access the raw methods on
}

// PermissionControllerMixinCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PermissionControllerMixinCallerRaw struct {
	Contract *PermissionControllerMixinCaller // Generic read-only contract binding to access the raw methods on
}

// PermissionControllerMixinTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PermissionControllerMixinTransactorRaw struct {
	Contract *PermissionControllerMixinTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPermissionControllerMixin creates a new instance of PermissionControllerMixin, bound to a specific deployed contract.
func NewPermissionControllerMixin(address common.Address, backend bind.ContractBackend) (*PermissionControllerMixin, error) {
	contract, err := bindPermissionControllerMixin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PermissionControllerMixin{PermissionControllerMixinCaller: PermissionControllerMixinCaller{contract: contract}, PermissionControllerMixinTransactor: PermissionControllerMixinTransactor{contract: contract}, PermissionControllerMixinFilterer: PermissionControllerMixinFilterer{contract: contract}}, nil
}

// NewPermissionControllerMixinCaller creates a new read-only instance of PermissionControllerMixin, bound to a specific deployed contract.
func NewPermissionControllerMixinCaller(address common.Address, caller bind.ContractCaller) (*PermissionControllerMixinCaller, error) {
	contract, err := bindPermissionControllerMixin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PermissionControllerMixinCaller{contract: contract}, nil
}

// NewPermissionControllerMixinTransactor creates a new write-only instance of PermissionControllerMixin, bound to a specific deployed contract.
func NewPermissionControllerMixinTransactor(address common.Address, transactor bind.ContractTransactor) (*PermissionControllerMixinTransactor, error) {
	contract, err := bindPermissionControllerMixin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PermissionControllerMixinTransactor{contract: contract}, nil
}

// NewPermissionControllerMixinFilterer creates a new log filterer instance of PermissionControllerMixin, bound to a specific deployed contract.
func NewPermissionControllerMixinFilterer(address common.Address, filterer bind.ContractFilterer) (*PermissionControllerMixinFilterer, error) {
	contract, err := bindPermissionControllerMixin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PermissionControllerMixinFilterer{contract: contract}, nil
}

// bindPermissionControllerMixin binds a generic wrapper to an already deployed contract.
func bindPermissionControllerMixin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PermissionControllerMixinMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PermissionControllerMixin *PermissionControllerMixinRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PermissionControllerMixin.Contract.PermissionControllerMixinCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PermissionControllerMixin *PermissionControllerMixinRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PermissionControllerMixin.Contract.PermissionControllerMixinTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PermissionControllerMixin *PermissionControllerMixinRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PermissionControllerMixin.Contract.PermissionControllerMixinTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PermissionControllerMixin *PermissionControllerMixinCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PermissionControllerMixin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PermissionControllerMixin *PermissionControllerMixinTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PermissionControllerMixin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PermissionControllerMixin *PermissionControllerMixinTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PermissionControllerMixin.Contract.contract.Transact(opts, method, params...)
}

// PermissionController is a free data retrieval call binding the contract method 0x4657e26a.
//
// Solidity: function permissionController() view returns(address)
func (_PermissionControllerMixin *PermissionControllerMixinCaller) PermissionController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PermissionControllerMixin.contract.Call(opts, &out, "permissionController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PermissionController is a free data retrieval call binding the contract method 0x4657e26a.
//
// Solidity: function permissionController() view returns(address)
func (_PermissionControllerMixin *PermissionControllerMixinSession) PermissionController() (common.Address, error) {
	return _PermissionControllerMixin.Contract.PermissionController(&_PermissionControllerMixin.CallOpts)
}

// PermissionController is a free data retrieval call binding the contract method 0x4657e26a.
//
// Solidity: function permissionController() view returns(address)
func (_PermissionControllerMixin *PermissionControllerMixinCallerSession) PermissionController() (common.Address, error) {
	return _PermissionControllerMixin.Contract.PermissionController(&_PermissionControllerMixin.CallOpts)
}
