// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package SignatureUtils

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

// SignatureUtilsMetaData contains all meta data concerning the SignatureUtils contract.
var SignatureUtilsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"domainSeparator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignatureExpired\",\"inputs\":[]}]",
}

// SignatureUtilsABI is the input ABI used to generate the binding from.
// Deprecated: Use SignatureUtilsMetaData.ABI instead.
var SignatureUtilsABI = SignatureUtilsMetaData.ABI

// SignatureUtils is an auto generated Go binding around an Ethereum contract.
type SignatureUtils struct {
	SignatureUtilsCaller     // Read-only binding to the contract
	SignatureUtilsTransactor // Write-only binding to the contract
	SignatureUtilsFilterer   // Log filterer for contract events
}

// SignatureUtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type SignatureUtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignatureUtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SignatureUtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignatureUtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SignatureUtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignatureUtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SignatureUtilsSession struct {
	Contract     *SignatureUtils   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SignatureUtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SignatureUtilsCallerSession struct {
	Contract *SignatureUtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// SignatureUtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SignatureUtilsTransactorSession struct {
	Contract     *SignatureUtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// SignatureUtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type SignatureUtilsRaw struct {
	Contract *SignatureUtils // Generic contract binding to access the raw methods on
}

// SignatureUtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SignatureUtilsCallerRaw struct {
	Contract *SignatureUtilsCaller // Generic read-only contract binding to access the raw methods on
}

// SignatureUtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SignatureUtilsTransactorRaw struct {
	Contract *SignatureUtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSignatureUtils creates a new instance of SignatureUtils, bound to a specific deployed contract.
func NewSignatureUtils(address common.Address, backend bind.ContractBackend) (*SignatureUtils, error) {
	contract, err := bindSignatureUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SignatureUtils{SignatureUtilsCaller: SignatureUtilsCaller{contract: contract}, SignatureUtilsTransactor: SignatureUtilsTransactor{contract: contract}, SignatureUtilsFilterer: SignatureUtilsFilterer{contract: contract}}, nil
}

// NewSignatureUtilsCaller creates a new read-only instance of SignatureUtils, bound to a specific deployed contract.
func NewSignatureUtilsCaller(address common.Address, caller bind.ContractCaller) (*SignatureUtilsCaller, error) {
	contract, err := bindSignatureUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SignatureUtilsCaller{contract: contract}, nil
}

// NewSignatureUtilsTransactor creates a new write-only instance of SignatureUtils, bound to a specific deployed contract.
func NewSignatureUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*SignatureUtilsTransactor, error) {
	contract, err := bindSignatureUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SignatureUtilsTransactor{contract: contract}, nil
}

// NewSignatureUtilsFilterer creates a new log filterer instance of SignatureUtils, bound to a specific deployed contract.
func NewSignatureUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*SignatureUtilsFilterer, error) {
	contract, err := bindSignatureUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SignatureUtilsFilterer{contract: contract}, nil
}

// bindSignatureUtils binds a generic wrapper to an already deployed contract.
func bindSignatureUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SignatureUtilsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SignatureUtils *SignatureUtilsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SignatureUtils.Contract.SignatureUtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SignatureUtils *SignatureUtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SignatureUtils.Contract.SignatureUtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SignatureUtils *SignatureUtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SignatureUtils.Contract.SignatureUtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SignatureUtils *SignatureUtilsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SignatureUtils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SignatureUtils *SignatureUtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SignatureUtils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SignatureUtils *SignatureUtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SignatureUtils.Contract.contract.Transact(opts, method, params...)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_SignatureUtils *SignatureUtilsCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SignatureUtils.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_SignatureUtils *SignatureUtilsSession) DomainSeparator() ([32]byte, error) {
	return _SignatureUtils.Contract.DomainSeparator(&_SignatureUtils.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_SignatureUtils *SignatureUtilsCallerSession) DomainSeparator() ([32]byte, error) {
	return _SignatureUtils.Contract.DomainSeparator(&_SignatureUtils.CallOpts)
}
