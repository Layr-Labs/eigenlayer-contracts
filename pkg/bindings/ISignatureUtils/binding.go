// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ISignatureUtils

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

// ISignatureUtilsMetaData contains all meta data concerning the ISignatureUtils contract.
var ISignatureUtilsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignatureExpired\",\"inputs\":[]}]",
}

// ISignatureUtilsABI is the input ABI used to generate the binding from.
// Deprecated: Use ISignatureUtilsMetaData.ABI instead.
var ISignatureUtilsABI = ISignatureUtilsMetaData.ABI

// ISignatureUtils is an auto generated Go binding around an Ethereum contract.
type ISignatureUtils struct {
	ISignatureUtilsCaller     // Read-only binding to the contract
	ISignatureUtilsTransactor // Write-only binding to the contract
	ISignatureUtilsFilterer   // Log filterer for contract events
}

// ISignatureUtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISignatureUtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISignatureUtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISignatureUtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISignatureUtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISignatureUtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISignatureUtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISignatureUtilsSession struct {
	Contract     *ISignatureUtils  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISignatureUtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISignatureUtilsCallerSession struct {
	Contract *ISignatureUtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ISignatureUtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISignatureUtilsTransactorSession struct {
	Contract     *ISignatureUtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ISignatureUtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISignatureUtilsRaw struct {
	Contract *ISignatureUtils // Generic contract binding to access the raw methods on
}

// ISignatureUtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISignatureUtilsCallerRaw struct {
	Contract *ISignatureUtilsCaller // Generic read-only contract binding to access the raw methods on
}

// ISignatureUtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISignatureUtilsTransactorRaw struct {
	Contract *ISignatureUtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISignatureUtils creates a new instance of ISignatureUtils, bound to a specific deployed contract.
func NewISignatureUtils(address common.Address, backend bind.ContractBackend) (*ISignatureUtils, error) {
	contract, err := bindISignatureUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISignatureUtils{ISignatureUtilsCaller: ISignatureUtilsCaller{contract: contract}, ISignatureUtilsTransactor: ISignatureUtilsTransactor{contract: contract}, ISignatureUtilsFilterer: ISignatureUtilsFilterer{contract: contract}}, nil
}

// NewISignatureUtilsCaller creates a new read-only instance of ISignatureUtils, bound to a specific deployed contract.
func NewISignatureUtilsCaller(address common.Address, caller bind.ContractCaller) (*ISignatureUtilsCaller, error) {
	contract, err := bindISignatureUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISignatureUtilsCaller{contract: contract}, nil
}

// NewISignatureUtilsTransactor creates a new write-only instance of ISignatureUtils, bound to a specific deployed contract.
func NewISignatureUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*ISignatureUtilsTransactor, error) {
	contract, err := bindISignatureUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISignatureUtilsTransactor{contract: contract}, nil
}

// NewISignatureUtilsFilterer creates a new log filterer instance of ISignatureUtils, bound to a specific deployed contract.
func NewISignatureUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*ISignatureUtilsFilterer, error) {
	contract, err := bindISignatureUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISignatureUtilsFilterer{contract: contract}, nil
}

// bindISignatureUtils binds a generic wrapper to an already deployed contract.
func bindISignatureUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ISignatureUtilsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISignatureUtils *ISignatureUtilsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISignatureUtils.Contract.ISignatureUtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISignatureUtils *ISignatureUtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISignatureUtils.Contract.ISignatureUtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISignatureUtils *ISignatureUtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISignatureUtils.Contract.ISignatureUtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISignatureUtils *ISignatureUtilsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISignatureUtils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISignatureUtils *ISignatureUtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISignatureUtils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISignatureUtils *ISignatureUtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISignatureUtils.Contract.contract.Transact(opts, method, params...)
}
