// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package SemVerMixin

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

// SemVerMixinMetaData contains all meta data concerning the SemVerMixin contract.
var SemVerMixinMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]}]",
}

// SemVerMixinABI is the input ABI used to generate the binding from.
// Deprecated: Use SemVerMixinMetaData.ABI instead.
var SemVerMixinABI = SemVerMixinMetaData.ABI

// SemVerMixin is an auto generated Go binding around an Ethereum contract.
type SemVerMixin struct {
	SemVerMixinCaller     // Read-only binding to the contract
	SemVerMixinTransactor // Write-only binding to the contract
	SemVerMixinFilterer   // Log filterer for contract events
}

// SemVerMixinCaller is an auto generated read-only Go binding around an Ethereum contract.
type SemVerMixinCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SemVerMixinTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SemVerMixinTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SemVerMixinFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SemVerMixinFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SemVerMixinSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SemVerMixinSession struct {
	Contract     *SemVerMixin      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SemVerMixinCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SemVerMixinCallerSession struct {
	Contract *SemVerMixinCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// SemVerMixinTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SemVerMixinTransactorSession struct {
	Contract     *SemVerMixinTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SemVerMixinRaw is an auto generated low-level Go binding around an Ethereum contract.
type SemVerMixinRaw struct {
	Contract *SemVerMixin // Generic contract binding to access the raw methods on
}

// SemVerMixinCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SemVerMixinCallerRaw struct {
	Contract *SemVerMixinCaller // Generic read-only contract binding to access the raw methods on
}

// SemVerMixinTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SemVerMixinTransactorRaw struct {
	Contract *SemVerMixinTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSemVerMixin creates a new instance of SemVerMixin, bound to a specific deployed contract.
func NewSemVerMixin(address common.Address, backend bind.ContractBackend) (*SemVerMixin, error) {
	contract, err := bindSemVerMixin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SemVerMixin{SemVerMixinCaller: SemVerMixinCaller{contract: contract}, SemVerMixinTransactor: SemVerMixinTransactor{contract: contract}, SemVerMixinFilterer: SemVerMixinFilterer{contract: contract}}, nil
}

// NewSemVerMixinCaller creates a new read-only instance of SemVerMixin, bound to a specific deployed contract.
func NewSemVerMixinCaller(address common.Address, caller bind.ContractCaller) (*SemVerMixinCaller, error) {
	contract, err := bindSemVerMixin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SemVerMixinCaller{contract: contract}, nil
}

// NewSemVerMixinTransactor creates a new write-only instance of SemVerMixin, bound to a specific deployed contract.
func NewSemVerMixinTransactor(address common.Address, transactor bind.ContractTransactor) (*SemVerMixinTransactor, error) {
	contract, err := bindSemVerMixin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SemVerMixinTransactor{contract: contract}, nil
}

// NewSemVerMixinFilterer creates a new log filterer instance of SemVerMixin, bound to a specific deployed contract.
func NewSemVerMixinFilterer(address common.Address, filterer bind.ContractFilterer) (*SemVerMixinFilterer, error) {
	contract, err := bindSemVerMixin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SemVerMixinFilterer{contract: contract}, nil
}

// bindSemVerMixin binds a generic wrapper to an already deployed contract.
func bindSemVerMixin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SemVerMixinMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SemVerMixin *SemVerMixinRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SemVerMixin.Contract.SemVerMixinCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SemVerMixin *SemVerMixinRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SemVerMixin.Contract.SemVerMixinTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SemVerMixin *SemVerMixinRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SemVerMixin.Contract.SemVerMixinTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SemVerMixin *SemVerMixinCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SemVerMixin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SemVerMixin *SemVerMixinTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SemVerMixin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SemVerMixin *SemVerMixinTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SemVerMixin.Contract.contract.Transact(opts, method, params...)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_SemVerMixin *SemVerMixinCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SemVerMixin.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_SemVerMixin *SemVerMixinSession) Version() (string, error) {
	return _SemVerMixin.Contract.Version(&_SemVerMixin.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_SemVerMixin *SemVerMixinCallerSession) Version() (string, error) {
	return _SemVerMixin.Contract.Version(&_SemVerMixin.CallOpts)
}
