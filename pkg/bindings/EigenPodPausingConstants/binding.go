// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package EigenPodPausingConstants

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

// EigenPodPausingConstantsMetaData contains all meta data concerning the EigenPodPausingConstants contract.
var EigenPodPausingConstantsMetaData = &bind.MetaData{
	ABI: "[]",
}

// EigenPodPausingConstantsABI is the input ABI used to generate the binding from.
// Deprecated: Use EigenPodPausingConstantsMetaData.ABI instead.
var EigenPodPausingConstantsABI = EigenPodPausingConstantsMetaData.ABI

// EigenPodPausingConstants is an auto generated Go binding around an Ethereum contract.
type EigenPodPausingConstants struct {
	EigenPodPausingConstantsCaller     // Read-only binding to the contract
	EigenPodPausingConstantsTransactor // Write-only binding to the contract
	EigenPodPausingConstantsFilterer   // Log filterer for contract events
}

// EigenPodPausingConstantsCaller is an auto generated read-only Go binding around an Ethereum contract.
type EigenPodPausingConstantsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EigenPodPausingConstantsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EigenPodPausingConstantsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EigenPodPausingConstantsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EigenPodPausingConstantsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EigenPodPausingConstantsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EigenPodPausingConstantsSession struct {
	Contract     *EigenPodPausingConstants // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// EigenPodPausingConstantsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EigenPodPausingConstantsCallerSession struct {
	Contract *EigenPodPausingConstantsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// EigenPodPausingConstantsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EigenPodPausingConstantsTransactorSession struct {
	Contract     *EigenPodPausingConstantsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// EigenPodPausingConstantsRaw is an auto generated low-level Go binding around an Ethereum contract.
type EigenPodPausingConstantsRaw struct {
	Contract *EigenPodPausingConstants // Generic contract binding to access the raw methods on
}

// EigenPodPausingConstantsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EigenPodPausingConstantsCallerRaw struct {
	Contract *EigenPodPausingConstantsCaller // Generic read-only contract binding to access the raw methods on
}

// EigenPodPausingConstantsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EigenPodPausingConstantsTransactorRaw struct {
	Contract *EigenPodPausingConstantsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEigenPodPausingConstants creates a new instance of EigenPodPausingConstants, bound to a specific deployed contract.
func NewEigenPodPausingConstants(address common.Address, backend bind.ContractBackend) (*EigenPodPausingConstants, error) {
	contract, err := bindEigenPodPausingConstants(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EigenPodPausingConstants{EigenPodPausingConstantsCaller: EigenPodPausingConstantsCaller{contract: contract}, EigenPodPausingConstantsTransactor: EigenPodPausingConstantsTransactor{contract: contract}, EigenPodPausingConstantsFilterer: EigenPodPausingConstantsFilterer{contract: contract}}, nil
}

// NewEigenPodPausingConstantsCaller creates a new read-only instance of EigenPodPausingConstants, bound to a specific deployed contract.
func NewEigenPodPausingConstantsCaller(address common.Address, caller bind.ContractCaller) (*EigenPodPausingConstantsCaller, error) {
	contract, err := bindEigenPodPausingConstants(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EigenPodPausingConstantsCaller{contract: contract}, nil
}

// NewEigenPodPausingConstantsTransactor creates a new write-only instance of EigenPodPausingConstants, bound to a specific deployed contract.
func NewEigenPodPausingConstantsTransactor(address common.Address, transactor bind.ContractTransactor) (*EigenPodPausingConstantsTransactor, error) {
	contract, err := bindEigenPodPausingConstants(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EigenPodPausingConstantsTransactor{contract: contract}, nil
}

// NewEigenPodPausingConstantsFilterer creates a new log filterer instance of EigenPodPausingConstants, bound to a specific deployed contract.
func NewEigenPodPausingConstantsFilterer(address common.Address, filterer bind.ContractFilterer) (*EigenPodPausingConstantsFilterer, error) {
	contract, err := bindEigenPodPausingConstants(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EigenPodPausingConstantsFilterer{contract: contract}, nil
}

// bindEigenPodPausingConstants binds a generic wrapper to an already deployed contract.
func bindEigenPodPausingConstants(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EigenPodPausingConstantsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EigenPodPausingConstants *EigenPodPausingConstantsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EigenPodPausingConstants.Contract.EigenPodPausingConstantsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EigenPodPausingConstants *EigenPodPausingConstantsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EigenPodPausingConstants.Contract.EigenPodPausingConstantsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EigenPodPausingConstants *EigenPodPausingConstantsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EigenPodPausingConstants.Contract.EigenPodPausingConstantsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EigenPodPausingConstants *EigenPodPausingConstantsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EigenPodPausingConstants.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EigenPodPausingConstants *EigenPodPausingConstantsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EigenPodPausingConstants.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EigenPodPausingConstants *EigenPodPausingConstantsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EigenPodPausingConstants.Contract.contract.Transact(opts, method, params...)
}
