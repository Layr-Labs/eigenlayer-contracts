// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Checkpoints

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

// CheckpointsMetaData contains all meta data concerning the Checkpoints contract.
var CheckpointsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122033cd598fba8494d1a6958a8af518adce5917043f9d96f1543c5e5350d9ce39b364736f6c634300080c0033",
}

// CheckpointsABI is the input ABI used to generate the binding from.
// Deprecated: Use CheckpointsMetaData.ABI instead.
var CheckpointsABI = CheckpointsMetaData.ABI

// CheckpointsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CheckpointsMetaData.Bin instead.
var CheckpointsBin = CheckpointsMetaData.Bin

// DeployCheckpoints deploys a new Ethereum contract, binding an instance of Checkpoints to it.
func DeployCheckpoints(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Checkpoints, error) {
	parsed, err := CheckpointsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CheckpointsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Checkpoints{CheckpointsCaller: CheckpointsCaller{contract: contract}, CheckpointsTransactor: CheckpointsTransactor{contract: contract}, CheckpointsFilterer: CheckpointsFilterer{contract: contract}}, nil
}

// Checkpoints is an auto generated Go binding around an Ethereum contract.
type Checkpoints struct {
	CheckpointsCaller     // Read-only binding to the contract
	CheckpointsTransactor // Write-only binding to the contract
	CheckpointsFilterer   // Log filterer for contract events
}

// CheckpointsCaller is an auto generated read-only Go binding around an Ethereum contract.
type CheckpointsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CheckpointsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CheckpointsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CheckpointsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CheckpointsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CheckpointsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CheckpointsSession struct {
	Contract     *Checkpoints      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CheckpointsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CheckpointsCallerSession struct {
	Contract *CheckpointsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// CheckpointsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CheckpointsTransactorSession struct {
	Contract     *CheckpointsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// CheckpointsRaw is an auto generated low-level Go binding around an Ethereum contract.
type CheckpointsRaw struct {
	Contract *Checkpoints // Generic contract binding to access the raw methods on
}

// CheckpointsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CheckpointsCallerRaw struct {
	Contract *CheckpointsCaller // Generic read-only contract binding to access the raw methods on
}

// CheckpointsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CheckpointsTransactorRaw struct {
	Contract *CheckpointsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCheckpoints creates a new instance of Checkpoints, bound to a specific deployed contract.
func NewCheckpoints(address common.Address, backend bind.ContractBackend) (*Checkpoints, error) {
	contract, err := bindCheckpoints(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Checkpoints{CheckpointsCaller: CheckpointsCaller{contract: contract}, CheckpointsTransactor: CheckpointsTransactor{contract: contract}, CheckpointsFilterer: CheckpointsFilterer{contract: contract}}, nil
}

// NewCheckpointsCaller creates a new read-only instance of Checkpoints, bound to a specific deployed contract.
func NewCheckpointsCaller(address common.Address, caller bind.ContractCaller) (*CheckpointsCaller, error) {
	contract, err := bindCheckpoints(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CheckpointsCaller{contract: contract}, nil
}

// NewCheckpointsTransactor creates a new write-only instance of Checkpoints, bound to a specific deployed contract.
func NewCheckpointsTransactor(address common.Address, transactor bind.ContractTransactor) (*CheckpointsTransactor, error) {
	contract, err := bindCheckpoints(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CheckpointsTransactor{contract: contract}, nil
}

// NewCheckpointsFilterer creates a new log filterer instance of Checkpoints, bound to a specific deployed contract.
func NewCheckpointsFilterer(address common.Address, filterer bind.ContractFilterer) (*CheckpointsFilterer, error) {
	contract, err := bindCheckpoints(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CheckpointsFilterer{contract: contract}, nil
}

// bindCheckpoints binds a generic wrapper to an already deployed contract.
func bindCheckpoints(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CheckpointsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Checkpoints *CheckpointsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Checkpoints.Contract.CheckpointsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Checkpoints *CheckpointsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Checkpoints.Contract.CheckpointsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Checkpoints *CheckpointsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Checkpoints.Contract.CheckpointsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Checkpoints *CheckpointsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Checkpoints.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Checkpoints *CheckpointsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Checkpoints.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Checkpoints *CheckpointsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Checkpoints.Contract.contract.Transact(opts, method, params...)
}
