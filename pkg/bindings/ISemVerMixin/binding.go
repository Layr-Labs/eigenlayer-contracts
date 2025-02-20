// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ISemVerMixin

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

// ISemVerMixinMetaData contains all meta data concerning the ISemVerMixin contract.
var ISemVerMixinMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"}]",
}

// ISemVerMixinABI is the input ABI used to generate the binding from.
// Deprecated: Use ISemVerMixinMetaData.ABI instead.
var ISemVerMixinABI = ISemVerMixinMetaData.ABI

// ISemVerMixin is an auto generated Go binding around an Ethereum contract.
type ISemVerMixin struct {
	ISemVerMixinCaller     // Read-only binding to the contract
	ISemVerMixinTransactor // Write-only binding to the contract
	ISemVerMixinFilterer   // Log filterer for contract events
}

// ISemVerMixinCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISemVerMixinCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISemVerMixinTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISemVerMixinTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISemVerMixinFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISemVerMixinFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISemVerMixinSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISemVerMixinSession struct {
	Contract     *ISemVerMixin     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISemVerMixinCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISemVerMixinCallerSession struct {
	Contract *ISemVerMixinCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ISemVerMixinTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISemVerMixinTransactorSession struct {
	Contract     *ISemVerMixinTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ISemVerMixinRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISemVerMixinRaw struct {
	Contract *ISemVerMixin // Generic contract binding to access the raw methods on
}

// ISemVerMixinCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISemVerMixinCallerRaw struct {
	Contract *ISemVerMixinCaller // Generic read-only contract binding to access the raw methods on
}

// ISemVerMixinTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISemVerMixinTransactorRaw struct {
	Contract *ISemVerMixinTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISemVerMixin creates a new instance of ISemVerMixin, bound to a specific deployed contract.
func NewISemVerMixin(address common.Address, backend bind.ContractBackend) (*ISemVerMixin, error) {
	contract, err := bindISemVerMixin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISemVerMixin{ISemVerMixinCaller: ISemVerMixinCaller{contract: contract}, ISemVerMixinTransactor: ISemVerMixinTransactor{contract: contract}, ISemVerMixinFilterer: ISemVerMixinFilterer{contract: contract}}, nil
}

// NewISemVerMixinCaller creates a new read-only instance of ISemVerMixin, bound to a specific deployed contract.
func NewISemVerMixinCaller(address common.Address, caller bind.ContractCaller) (*ISemVerMixinCaller, error) {
	contract, err := bindISemVerMixin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISemVerMixinCaller{contract: contract}, nil
}

// NewISemVerMixinTransactor creates a new write-only instance of ISemVerMixin, bound to a specific deployed contract.
func NewISemVerMixinTransactor(address common.Address, transactor bind.ContractTransactor) (*ISemVerMixinTransactor, error) {
	contract, err := bindISemVerMixin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISemVerMixinTransactor{contract: contract}, nil
}

// NewISemVerMixinFilterer creates a new log filterer instance of ISemVerMixin, bound to a specific deployed contract.
func NewISemVerMixinFilterer(address common.Address, filterer bind.ContractFilterer) (*ISemVerMixinFilterer, error) {
	contract, err := bindISemVerMixin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISemVerMixinFilterer{contract: contract}, nil
}

// bindISemVerMixin binds a generic wrapper to an already deployed contract.
func bindISemVerMixin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ISemVerMixinMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISemVerMixin *ISemVerMixinRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISemVerMixin.Contract.ISemVerMixinCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISemVerMixin *ISemVerMixinRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISemVerMixin.Contract.ISemVerMixinTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISemVerMixin *ISemVerMixinRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISemVerMixin.Contract.ISemVerMixinTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISemVerMixin *ISemVerMixinCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISemVerMixin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISemVerMixin *ISemVerMixinTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISemVerMixin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISemVerMixin *ISemVerMixinTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISemVerMixin.Contract.contract.Transact(opts, method, params...)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ISemVerMixin *ISemVerMixinCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ISemVerMixin.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ISemVerMixin *ISemVerMixinSession) Version() (string, error) {
	return _ISemVerMixin.Contract.Version(&_ISemVerMixin.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ISemVerMixin *ISemVerMixinCallerSession) Version() (string, error) {
	return _ISemVerMixin.Contract.Version(&_ISemVerMixin.CallOpts)
}
