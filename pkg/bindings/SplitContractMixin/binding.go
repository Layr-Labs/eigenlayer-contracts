// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package SplitContractMixin

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

// SplitContractMixinMetaData contains all meta data concerning the SplitContractMixin contract.
var SplitContractMixinMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"viewImplementation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"}]",
}

// SplitContractMixinABI is the input ABI used to generate the binding from.
// Deprecated: Use SplitContractMixinMetaData.ABI instead.
var SplitContractMixinABI = SplitContractMixinMetaData.ABI

// SplitContractMixin is an auto generated Go binding around an Ethereum contract.
type SplitContractMixin struct {
	SplitContractMixinCaller     // Read-only binding to the contract
	SplitContractMixinTransactor // Write-only binding to the contract
	SplitContractMixinFilterer   // Log filterer for contract events
}

// SplitContractMixinCaller is an auto generated read-only Go binding around an Ethereum contract.
type SplitContractMixinCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SplitContractMixinTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SplitContractMixinTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SplitContractMixinFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SplitContractMixinFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SplitContractMixinSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SplitContractMixinSession struct {
	Contract     *SplitContractMixin // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SplitContractMixinCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SplitContractMixinCallerSession struct {
	Contract *SplitContractMixinCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// SplitContractMixinTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SplitContractMixinTransactorSession struct {
	Contract     *SplitContractMixinTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// SplitContractMixinRaw is an auto generated low-level Go binding around an Ethereum contract.
type SplitContractMixinRaw struct {
	Contract *SplitContractMixin // Generic contract binding to access the raw methods on
}

// SplitContractMixinCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SplitContractMixinCallerRaw struct {
	Contract *SplitContractMixinCaller // Generic read-only contract binding to access the raw methods on
}

// SplitContractMixinTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SplitContractMixinTransactorRaw struct {
	Contract *SplitContractMixinTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSplitContractMixin creates a new instance of SplitContractMixin, bound to a specific deployed contract.
func NewSplitContractMixin(address common.Address, backend bind.ContractBackend) (*SplitContractMixin, error) {
	contract, err := bindSplitContractMixin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SplitContractMixin{SplitContractMixinCaller: SplitContractMixinCaller{contract: contract}, SplitContractMixinTransactor: SplitContractMixinTransactor{contract: contract}, SplitContractMixinFilterer: SplitContractMixinFilterer{contract: contract}}, nil
}

// NewSplitContractMixinCaller creates a new read-only instance of SplitContractMixin, bound to a specific deployed contract.
func NewSplitContractMixinCaller(address common.Address, caller bind.ContractCaller) (*SplitContractMixinCaller, error) {
	contract, err := bindSplitContractMixin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SplitContractMixinCaller{contract: contract}, nil
}

// NewSplitContractMixinTransactor creates a new write-only instance of SplitContractMixin, bound to a specific deployed contract.
func NewSplitContractMixinTransactor(address common.Address, transactor bind.ContractTransactor) (*SplitContractMixinTransactor, error) {
	contract, err := bindSplitContractMixin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SplitContractMixinTransactor{contract: contract}, nil
}

// NewSplitContractMixinFilterer creates a new log filterer instance of SplitContractMixin, bound to a specific deployed contract.
func NewSplitContractMixinFilterer(address common.Address, filterer bind.ContractFilterer) (*SplitContractMixinFilterer, error) {
	contract, err := bindSplitContractMixin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SplitContractMixinFilterer{contract: contract}, nil
}

// bindSplitContractMixin binds a generic wrapper to an already deployed contract.
func bindSplitContractMixin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SplitContractMixinMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SplitContractMixin *SplitContractMixinRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SplitContractMixin.Contract.SplitContractMixinCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SplitContractMixin *SplitContractMixinRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SplitContractMixin.Contract.SplitContractMixinTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SplitContractMixin *SplitContractMixinRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SplitContractMixin.Contract.SplitContractMixinTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SplitContractMixin *SplitContractMixinCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SplitContractMixin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SplitContractMixin *SplitContractMixinTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SplitContractMixin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SplitContractMixin *SplitContractMixinTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SplitContractMixin.Contract.contract.Transact(opts, method, params...)
}

// ViewImplementation is a free data retrieval call binding the contract method 0x0b156bb6.
//
// Solidity: function viewImplementation() view returns(address)
func (_SplitContractMixin *SplitContractMixinCaller) ViewImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SplitContractMixin.contract.Call(opts, &out, "viewImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ViewImplementation is a free data retrieval call binding the contract method 0x0b156bb6.
//
// Solidity: function viewImplementation() view returns(address)
func (_SplitContractMixin *SplitContractMixinSession) ViewImplementation() (common.Address, error) {
	return _SplitContractMixin.Contract.ViewImplementation(&_SplitContractMixin.CallOpts)
}

// ViewImplementation is a free data retrieval call binding the contract method 0x0b156bb6.
//
// Solidity: function viewImplementation() view returns(address)
func (_SplitContractMixin *SplitContractMixinCallerSession) ViewImplementation() (common.Address, error) {
	return _SplitContractMixin.Contract.ViewImplementation(&_SplitContractMixin.CallOpts)
}
