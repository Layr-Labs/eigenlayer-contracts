// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package EIP1271SignatureUtils

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

// EIP1271SignatureUtilsMetaData contains all meta data concerning the EIP1271SignatureUtils contract.
var EIP1271SignatureUtilsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212209832fdd04d36d0bbe78b7b9457ca6759978d3e81a33602cb69f518c735ee384664736f6c634300080c0033",
}

// EIP1271SignatureUtilsABI is the input ABI used to generate the binding from.
// Deprecated: Use EIP1271SignatureUtilsMetaData.ABI instead.
var EIP1271SignatureUtilsABI = EIP1271SignatureUtilsMetaData.ABI

// EIP1271SignatureUtilsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EIP1271SignatureUtilsMetaData.Bin instead.
var EIP1271SignatureUtilsBin = EIP1271SignatureUtilsMetaData.Bin

// DeployEIP1271SignatureUtils deploys a new Ethereum contract, binding an instance of EIP1271SignatureUtils to it.
func DeployEIP1271SignatureUtils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EIP1271SignatureUtils, error) {
	parsed, err := EIP1271SignatureUtilsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EIP1271SignatureUtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EIP1271SignatureUtils{EIP1271SignatureUtilsCaller: EIP1271SignatureUtilsCaller{contract: contract}, EIP1271SignatureUtilsTransactor: EIP1271SignatureUtilsTransactor{contract: contract}, EIP1271SignatureUtilsFilterer: EIP1271SignatureUtilsFilterer{contract: contract}}, nil
}

// EIP1271SignatureUtils is an auto generated Go binding around an Ethereum contract.
type EIP1271SignatureUtils struct {
	EIP1271SignatureUtilsCaller     // Read-only binding to the contract
	EIP1271SignatureUtilsTransactor // Write-only binding to the contract
	EIP1271SignatureUtilsFilterer   // Log filterer for contract events
}

// EIP1271SignatureUtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type EIP1271SignatureUtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EIP1271SignatureUtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EIP1271SignatureUtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EIP1271SignatureUtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EIP1271SignatureUtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EIP1271SignatureUtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EIP1271SignatureUtilsSession struct {
	Contract     *EIP1271SignatureUtils // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// EIP1271SignatureUtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EIP1271SignatureUtilsCallerSession struct {
	Contract *EIP1271SignatureUtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// EIP1271SignatureUtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EIP1271SignatureUtilsTransactorSession struct {
	Contract     *EIP1271SignatureUtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// EIP1271SignatureUtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type EIP1271SignatureUtilsRaw struct {
	Contract *EIP1271SignatureUtils // Generic contract binding to access the raw methods on
}

// EIP1271SignatureUtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EIP1271SignatureUtilsCallerRaw struct {
	Contract *EIP1271SignatureUtilsCaller // Generic read-only contract binding to access the raw methods on
}

// EIP1271SignatureUtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EIP1271SignatureUtilsTransactorRaw struct {
	Contract *EIP1271SignatureUtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEIP1271SignatureUtils creates a new instance of EIP1271SignatureUtils, bound to a specific deployed contract.
func NewEIP1271SignatureUtils(address common.Address, backend bind.ContractBackend) (*EIP1271SignatureUtils, error) {
	contract, err := bindEIP1271SignatureUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EIP1271SignatureUtils{EIP1271SignatureUtilsCaller: EIP1271SignatureUtilsCaller{contract: contract}, EIP1271SignatureUtilsTransactor: EIP1271SignatureUtilsTransactor{contract: contract}, EIP1271SignatureUtilsFilterer: EIP1271SignatureUtilsFilterer{contract: contract}}, nil
}

// NewEIP1271SignatureUtilsCaller creates a new read-only instance of EIP1271SignatureUtils, bound to a specific deployed contract.
func NewEIP1271SignatureUtilsCaller(address common.Address, caller bind.ContractCaller) (*EIP1271SignatureUtilsCaller, error) {
	contract, err := bindEIP1271SignatureUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EIP1271SignatureUtilsCaller{contract: contract}, nil
}

// NewEIP1271SignatureUtilsTransactor creates a new write-only instance of EIP1271SignatureUtils, bound to a specific deployed contract.
func NewEIP1271SignatureUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*EIP1271SignatureUtilsTransactor, error) {
	contract, err := bindEIP1271SignatureUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EIP1271SignatureUtilsTransactor{contract: contract}, nil
}

// NewEIP1271SignatureUtilsFilterer creates a new log filterer instance of EIP1271SignatureUtils, bound to a specific deployed contract.
func NewEIP1271SignatureUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*EIP1271SignatureUtilsFilterer, error) {
	contract, err := bindEIP1271SignatureUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EIP1271SignatureUtilsFilterer{contract: contract}, nil
}

// bindEIP1271SignatureUtils binds a generic wrapper to an already deployed contract.
func bindEIP1271SignatureUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EIP1271SignatureUtilsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EIP1271SignatureUtils *EIP1271SignatureUtilsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EIP1271SignatureUtils.Contract.EIP1271SignatureUtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EIP1271SignatureUtils *EIP1271SignatureUtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EIP1271SignatureUtils.Contract.EIP1271SignatureUtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EIP1271SignatureUtils *EIP1271SignatureUtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EIP1271SignatureUtils.Contract.EIP1271SignatureUtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EIP1271SignatureUtils *EIP1271SignatureUtilsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EIP1271SignatureUtils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EIP1271SignatureUtils *EIP1271SignatureUtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EIP1271SignatureUtils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EIP1271SignatureUtils *EIP1271SignatureUtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EIP1271SignatureUtils.Contract.contract.Transact(opts, method, params...)
}
