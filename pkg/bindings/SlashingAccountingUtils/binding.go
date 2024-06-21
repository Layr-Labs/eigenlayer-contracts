// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package SlashingAccountingUtils

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

// SlashingAccountingUtilsMetaData contains all meta data concerning the SlashingAccountingUtils contract.
var SlashingAccountingUtilsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122012cf4f8df15f613dfd59f63ccdc768eef75419a497be7791ae4891d4f9fd126564736f6c634300080c0033",
}

// SlashingAccountingUtilsABI is the input ABI used to generate the binding from.
// Deprecated: Use SlashingAccountingUtilsMetaData.ABI instead.
var SlashingAccountingUtilsABI = SlashingAccountingUtilsMetaData.ABI

// SlashingAccountingUtilsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SlashingAccountingUtilsMetaData.Bin instead.
var SlashingAccountingUtilsBin = SlashingAccountingUtilsMetaData.Bin

// DeploySlashingAccountingUtils deploys a new Ethereum contract, binding an instance of SlashingAccountingUtils to it.
func DeploySlashingAccountingUtils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SlashingAccountingUtils, error) {
	parsed, err := SlashingAccountingUtilsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SlashingAccountingUtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SlashingAccountingUtils{SlashingAccountingUtilsCaller: SlashingAccountingUtilsCaller{contract: contract}, SlashingAccountingUtilsTransactor: SlashingAccountingUtilsTransactor{contract: contract}, SlashingAccountingUtilsFilterer: SlashingAccountingUtilsFilterer{contract: contract}}, nil
}

// SlashingAccountingUtils is an auto generated Go binding around an Ethereum contract.
type SlashingAccountingUtils struct {
	SlashingAccountingUtilsCaller     // Read-only binding to the contract
	SlashingAccountingUtilsTransactor // Write-only binding to the contract
	SlashingAccountingUtilsFilterer   // Log filterer for contract events
}

// SlashingAccountingUtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type SlashingAccountingUtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SlashingAccountingUtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SlashingAccountingUtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SlashingAccountingUtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SlashingAccountingUtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SlashingAccountingUtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SlashingAccountingUtilsSession struct {
	Contract     *SlashingAccountingUtils // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SlashingAccountingUtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SlashingAccountingUtilsCallerSession struct {
	Contract *SlashingAccountingUtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// SlashingAccountingUtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SlashingAccountingUtilsTransactorSession struct {
	Contract     *SlashingAccountingUtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// SlashingAccountingUtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type SlashingAccountingUtilsRaw struct {
	Contract *SlashingAccountingUtils // Generic contract binding to access the raw methods on
}

// SlashingAccountingUtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SlashingAccountingUtilsCallerRaw struct {
	Contract *SlashingAccountingUtilsCaller // Generic read-only contract binding to access the raw methods on
}

// SlashingAccountingUtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SlashingAccountingUtilsTransactorRaw struct {
	Contract *SlashingAccountingUtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSlashingAccountingUtils creates a new instance of SlashingAccountingUtils, bound to a specific deployed contract.
func NewSlashingAccountingUtils(address common.Address, backend bind.ContractBackend) (*SlashingAccountingUtils, error) {
	contract, err := bindSlashingAccountingUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SlashingAccountingUtils{SlashingAccountingUtilsCaller: SlashingAccountingUtilsCaller{contract: contract}, SlashingAccountingUtilsTransactor: SlashingAccountingUtilsTransactor{contract: contract}, SlashingAccountingUtilsFilterer: SlashingAccountingUtilsFilterer{contract: contract}}, nil
}

// NewSlashingAccountingUtilsCaller creates a new read-only instance of SlashingAccountingUtils, bound to a specific deployed contract.
func NewSlashingAccountingUtilsCaller(address common.Address, caller bind.ContractCaller) (*SlashingAccountingUtilsCaller, error) {
	contract, err := bindSlashingAccountingUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SlashingAccountingUtilsCaller{contract: contract}, nil
}

// NewSlashingAccountingUtilsTransactor creates a new write-only instance of SlashingAccountingUtils, bound to a specific deployed contract.
func NewSlashingAccountingUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*SlashingAccountingUtilsTransactor, error) {
	contract, err := bindSlashingAccountingUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SlashingAccountingUtilsTransactor{contract: contract}, nil
}

// NewSlashingAccountingUtilsFilterer creates a new log filterer instance of SlashingAccountingUtils, bound to a specific deployed contract.
func NewSlashingAccountingUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*SlashingAccountingUtilsFilterer, error) {
	contract, err := bindSlashingAccountingUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SlashingAccountingUtilsFilterer{contract: contract}, nil
}

// bindSlashingAccountingUtils binds a generic wrapper to an already deployed contract.
func bindSlashingAccountingUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SlashingAccountingUtilsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SlashingAccountingUtils *SlashingAccountingUtilsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SlashingAccountingUtils.Contract.SlashingAccountingUtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SlashingAccountingUtils *SlashingAccountingUtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SlashingAccountingUtils.Contract.SlashingAccountingUtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SlashingAccountingUtils *SlashingAccountingUtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SlashingAccountingUtils.Contract.SlashingAccountingUtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SlashingAccountingUtils *SlashingAccountingUtilsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SlashingAccountingUtils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SlashingAccountingUtils *SlashingAccountingUtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SlashingAccountingUtils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SlashingAccountingUtils *SlashingAccountingUtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SlashingAccountingUtils.Contract.contract.Transact(opts, method, params...)
}
