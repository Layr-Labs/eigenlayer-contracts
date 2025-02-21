// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ISignatureUtilsMixin

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

// ISignatureUtilsMixinMetaData contains all meta data concerning the ISignatureUtilsMixin contract.
var ISignatureUtilsMixinMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"domainSeparator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignatureExpired\",\"inputs\":[]}]",
}

// ISignatureUtilsMixinABI is the input ABI used to generate the binding from.
// Deprecated: Use ISignatureUtilsMixinMetaData.ABI instead.
var ISignatureUtilsMixinABI = ISignatureUtilsMixinMetaData.ABI

// ISignatureUtilsMixin is an auto generated Go binding around an Ethereum contract.
type ISignatureUtilsMixin struct {
	ISignatureUtilsMixinCaller     // Read-only binding to the contract
	ISignatureUtilsMixinTransactor // Write-only binding to the contract
	ISignatureUtilsMixinFilterer   // Log filterer for contract events
}

// ISignatureUtilsMixinCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISignatureUtilsMixinCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISignatureUtilsMixinTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISignatureUtilsMixinTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISignatureUtilsMixinFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISignatureUtilsMixinFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISignatureUtilsMixinSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISignatureUtilsMixinSession struct {
	Contract     *ISignatureUtilsMixin // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ISignatureUtilsMixinCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISignatureUtilsMixinCallerSession struct {
	Contract *ISignatureUtilsMixinCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// ISignatureUtilsMixinTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISignatureUtilsMixinTransactorSession struct {
	Contract     *ISignatureUtilsMixinTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// ISignatureUtilsMixinRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISignatureUtilsMixinRaw struct {
	Contract *ISignatureUtilsMixin // Generic contract binding to access the raw methods on
}

// ISignatureUtilsMixinCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISignatureUtilsMixinCallerRaw struct {
	Contract *ISignatureUtilsMixinCaller // Generic read-only contract binding to access the raw methods on
}

// ISignatureUtilsMixinTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISignatureUtilsMixinTransactorRaw struct {
	Contract *ISignatureUtilsMixinTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISignatureUtilsMixin creates a new instance of ISignatureUtilsMixin, bound to a specific deployed contract.
func NewISignatureUtilsMixin(address common.Address, backend bind.ContractBackend) (*ISignatureUtilsMixin, error) {
	contract, err := bindISignatureUtilsMixin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISignatureUtilsMixin{ISignatureUtilsMixinCaller: ISignatureUtilsMixinCaller{contract: contract}, ISignatureUtilsMixinTransactor: ISignatureUtilsMixinTransactor{contract: contract}, ISignatureUtilsMixinFilterer: ISignatureUtilsMixinFilterer{contract: contract}}, nil
}

// NewISignatureUtilsMixinCaller creates a new read-only instance of ISignatureUtilsMixin, bound to a specific deployed contract.
func NewISignatureUtilsMixinCaller(address common.Address, caller bind.ContractCaller) (*ISignatureUtilsMixinCaller, error) {
	contract, err := bindISignatureUtilsMixin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISignatureUtilsMixinCaller{contract: contract}, nil
}

// NewISignatureUtilsMixinTransactor creates a new write-only instance of ISignatureUtilsMixin, bound to a specific deployed contract.
func NewISignatureUtilsMixinTransactor(address common.Address, transactor bind.ContractTransactor) (*ISignatureUtilsMixinTransactor, error) {
	contract, err := bindISignatureUtilsMixin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISignatureUtilsMixinTransactor{contract: contract}, nil
}

// NewISignatureUtilsMixinFilterer creates a new log filterer instance of ISignatureUtilsMixin, bound to a specific deployed contract.
func NewISignatureUtilsMixinFilterer(address common.Address, filterer bind.ContractFilterer) (*ISignatureUtilsMixinFilterer, error) {
	contract, err := bindISignatureUtilsMixin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISignatureUtilsMixinFilterer{contract: contract}, nil
}

// bindISignatureUtilsMixin binds a generic wrapper to an already deployed contract.
func bindISignatureUtilsMixin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ISignatureUtilsMixinMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISignatureUtilsMixin *ISignatureUtilsMixinRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISignatureUtilsMixin.Contract.ISignatureUtilsMixinCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISignatureUtilsMixin *ISignatureUtilsMixinRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISignatureUtilsMixin.Contract.ISignatureUtilsMixinTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISignatureUtilsMixin *ISignatureUtilsMixinRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISignatureUtilsMixin.Contract.ISignatureUtilsMixinTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISignatureUtilsMixin *ISignatureUtilsMixinCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISignatureUtilsMixin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISignatureUtilsMixin *ISignatureUtilsMixinTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISignatureUtilsMixin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISignatureUtilsMixin *ISignatureUtilsMixinTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISignatureUtilsMixin.Contract.contract.Transact(opts, method, params...)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_ISignatureUtilsMixin *ISignatureUtilsMixinCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ISignatureUtilsMixin.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_ISignatureUtilsMixin *ISignatureUtilsMixinSession) DomainSeparator() ([32]byte, error) {
	return _ISignatureUtilsMixin.Contract.DomainSeparator(&_ISignatureUtilsMixin.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_ISignatureUtilsMixin *ISignatureUtilsMixinCallerSession) DomainSeparator() ([32]byte, error) {
	return _ISignatureUtilsMixin.Contract.DomainSeparator(&_ISignatureUtilsMixin.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ISignatureUtilsMixin *ISignatureUtilsMixinCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ISignatureUtilsMixin.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ISignatureUtilsMixin *ISignatureUtilsMixinSession) Version() (string, error) {
	return _ISignatureUtilsMixin.Contract.Version(&_ISignatureUtilsMixin.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ISignatureUtilsMixin *ISignatureUtilsMixinCallerSession) Version() (string, error) {
	return _ISignatureUtilsMixin.Contract.Version(&_ISignatureUtilsMixin.CallOpts)
}
