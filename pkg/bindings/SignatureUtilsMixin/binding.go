// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package SignatureUtilsMixin

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

// SignatureUtilsMixinMetaData contains all meta data concerning the SignatureUtilsMixin contract.
var SignatureUtilsMixinMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"domainSeparator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignatureExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]}]",
}

// SignatureUtilsMixinABI is the input ABI used to generate the binding from.
// Deprecated: Use SignatureUtilsMixinMetaData.ABI instead.
var SignatureUtilsMixinABI = SignatureUtilsMixinMetaData.ABI

// SignatureUtilsMixin is an auto generated Go binding around an Ethereum contract.
type SignatureUtilsMixin struct {
	SignatureUtilsMixinCaller     // Read-only binding to the contract
	SignatureUtilsMixinTransactor // Write-only binding to the contract
	SignatureUtilsMixinFilterer   // Log filterer for contract events
}

// SignatureUtilsMixinCaller is an auto generated read-only Go binding around an Ethereum contract.
type SignatureUtilsMixinCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignatureUtilsMixinTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SignatureUtilsMixinTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignatureUtilsMixinFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SignatureUtilsMixinFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignatureUtilsMixinSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SignatureUtilsMixinSession struct {
	Contract     *SignatureUtilsMixin // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SignatureUtilsMixinCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SignatureUtilsMixinCallerSession struct {
	Contract *SignatureUtilsMixinCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// SignatureUtilsMixinTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SignatureUtilsMixinTransactorSession struct {
	Contract     *SignatureUtilsMixinTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// SignatureUtilsMixinRaw is an auto generated low-level Go binding around an Ethereum contract.
type SignatureUtilsMixinRaw struct {
	Contract *SignatureUtilsMixin // Generic contract binding to access the raw methods on
}

// SignatureUtilsMixinCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SignatureUtilsMixinCallerRaw struct {
	Contract *SignatureUtilsMixinCaller // Generic read-only contract binding to access the raw methods on
}

// SignatureUtilsMixinTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SignatureUtilsMixinTransactorRaw struct {
	Contract *SignatureUtilsMixinTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSignatureUtilsMixin creates a new instance of SignatureUtilsMixin, bound to a specific deployed contract.
func NewSignatureUtilsMixin(address common.Address, backend bind.ContractBackend) (*SignatureUtilsMixin, error) {
	contract, err := bindSignatureUtilsMixin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SignatureUtilsMixin{SignatureUtilsMixinCaller: SignatureUtilsMixinCaller{contract: contract}, SignatureUtilsMixinTransactor: SignatureUtilsMixinTransactor{contract: contract}, SignatureUtilsMixinFilterer: SignatureUtilsMixinFilterer{contract: contract}}, nil
}

// NewSignatureUtilsMixinCaller creates a new read-only instance of SignatureUtilsMixin, bound to a specific deployed contract.
func NewSignatureUtilsMixinCaller(address common.Address, caller bind.ContractCaller) (*SignatureUtilsMixinCaller, error) {
	contract, err := bindSignatureUtilsMixin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SignatureUtilsMixinCaller{contract: contract}, nil
}

// NewSignatureUtilsMixinTransactor creates a new write-only instance of SignatureUtilsMixin, bound to a specific deployed contract.
func NewSignatureUtilsMixinTransactor(address common.Address, transactor bind.ContractTransactor) (*SignatureUtilsMixinTransactor, error) {
	contract, err := bindSignatureUtilsMixin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SignatureUtilsMixinTransactor{contract: contract}, nil
}

// NewSignatureUtilsMixinFilterer creates a new log filterer instance of SignatureUtilsMixin, bound to a specific deployed contract.
func NewSignatureUtilsMixinFilterer(address common.Address, filterer bind.ContractFilterer) (*SignatureUtilsMixinFilterer, error) {
	contract, err := bindSignatureUtilsMixin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SignatureUtilsMixinFilterer{contract: contract}, nil
}

// bindSignatureUtilsMixin binds a generic wrapper to an already deployed contract.
func bindSignatureUtilsMixin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SignatureUtilsMixinMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SignatureUtilsMixin *SignatureUtilsMixinRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SignatureUtilsMixin.Contract.SignatureUtilsMixinCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SignatureUtilsMixin *SignatureUtilsMixinRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SignatureUtilsMixin.Contract.SignatureUtilsMixinTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SignatureUtilsMixin *SignatureUtilsMixinRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SignatureUtilsMixin.Contract.SignatureUtilsMixinTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SignatureUtilsMixin *SignatureUtilsMixinCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SignatureUtilsMixin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SignatureUtilsMixin *SignatureUtilsMixinTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SignatureUtilsMixin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SignatureUtilsMixin *SignatureUtilsMixinTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SignatureUtilsMixin.Contract.contract.Transact(opts, method, params...)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_SignatureUtilsMixin *SignatureUtilsMixinCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SignatureUtilsMixin.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_SignatureUtilsMixin *SignatureUtilsMixinSession) DomainSeparator() ([32]byte, error) {
	return _SignatureUtilsMixin.Contract.DomainSeparator(&_SignatureUtilsMixin.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_SignatureUtilsMixin *SignatureUtilsMixinCallerSession) DomainSeparator() ([32]byte, error) {
	return _SignatureUtilsMixin.Contract.DomainSeparator(&_SignatureUtilsMixin.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_SignatureUtilsMixin *SignatureUtilsMixinCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SignatureUtilsMixin.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_SignatureUtilsMixin *SignatureUtilsMixinSession) Version() (string, error) {
	return _SignatureUtilsMixin.Contract.Version(&_SignatureUtilsMixin.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_SignatureUtilsMixin *SignatureUtilsMixinCallerSession) Version() (string, error) {
	return _SignatureUtilsMixin.Contract.Version(&_SignatureUtilsMixin.CallOpts)
}
