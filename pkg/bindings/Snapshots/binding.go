// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Snapshots

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

// SnapshotsMetaData contains all meta data concerning the Snapshots contract.
var SnapshotsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"error\",\"name\":\"InvalidSnapshotOrdering\",\"inputs\":[]}]",
	Bin: "0x60556032600b8282823980515f1a607314602657634e487b7160e01b5f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040525f5ffdfea264697066735822122039ae8d80d6b79cfe3f3fa1acda95d5467ae7b0e1b54d0fccbf6fba10d9bd1f4864736f6c634300081b0033",
}

// SnapshotsABI is the input ABI used to generate the binding from.
// Deprecated: Use SnapshotsMetaData.ABI instead.
var SnapshotsABI = SnapshotsMetaData.ABI

// SnapshotsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SnapshotsMetaData.Bin instead.
var SnapshotsBin = SnapshotsMetaData.Bin

// DeploySnapshots deploys a new Ethereum contract, binding an instance of Snapshots to it.
func DeploySnapshots(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Snapshots, error) {
	parsed, err := SnapshotsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SnapshotsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Snapshots{SnapshotsCaller: SnapshotsCaller{contract: contract}, SnapshotsTransactor: SnapshotsTransactor{contract: contract}, SnapshotsFilterer: SnapshotsFilterer{contract: contract}}, nil
}

// Snapshots is an auto generated Go binding around an Ethereum contract.
type Snapshots struct {
	SnapshotsCaller     // Read-only binding to the contract
	SnapshotsTransactor // Write-only binding to the contract
	SnapshotsFilterer   // Log filterer for contract events
}

// SnapshotsCaller is an auto generated read-only Go binding around an Ethereum contract.
type SnapshotsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SnapshotsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SnapshotsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SnapshotsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SnapshotsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SnapshotsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SnapshotsSession struct {
	Contract     *Snapshots        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SnapshotsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SnapshotsCallerSession struct {
	Contract *SnapshotsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SnapshotsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SnapshotsTransactorSession struct {
	Contract     *SnapshotsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SnapshotsRaw is an auto generated low-level Go binding around an Ethereum contract.
type SnapshotsRaw struct {
	Contract *Snapshots // Generic contract binding to access the raw methods on
}

// SnapshotsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SnapshotsCallerRaw struct {
	Contract *SnapshotsCaller // Generic read-only contract binding to access the raw methods on
}

// SnapshotsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SnapshotsTransactorRaw struct {
	Contract *SnapshotsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSnapshots creates a new instance of Snapshots, bound to a specific deployed contract.
func NewSnapshots(address common.Address, backend bind.ContractBackend) (*Snapshots, error) {
	contract, err := bindSnapshots(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Snapshots{SnapshotsCaller: SnapshotsCaller{contract: contract}, SnapshotsTransactor: SnapshotsTransactor{contract: contract}, SnapshotsFilterer: SnapshotsFilterer{contract: contract}}, nil
}

// NewSnapshotsCaller creates a new read-only instance of Snapshots, bound to a specific deployed contract.
func NewSnapshotsCaller(address common.Address, caller bind.ContractCaller) (*SnapshotsCaller, error) {
	contract, err := bindSnapshots(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SnapshotsCaller{contract: contract}, nil
}

// NewSnapshotsTransactor creates a new write-only instance of Snapshots, bound to a specific deployed contract.
func NewSnapshotsTransactor(address common.Address, transactor bind.ContractTransactor) (*SnapshotsTransactor, error) {
	contract, err := bindSnapshots(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SnapshotsTransactor{contract: contract}, nil
}

// NewSnapshotsFilterer creates a new log filterer instance of Snapshots, bound to a specific deployed contract.
func NewSnapshotsFilterer(address common.Address, filterer bind.ContractFilterer) (*SnapshotsFilterer, error) {
	contract, err := bindSnapshots(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SnapshotsFilterer{contract: contract}, nil
}

// bindSnapshots binds a generic wrapper to an already deployed contract.
func bindSnapshots(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SnapshotsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Snapshots *SnapshotsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Snapshots.Contract.SnapshotsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Snapshots *SnapshotsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Snapshots.Contract.SnapshotsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Snapshots *SnapshotsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Snapshots.Contract.SnapshotsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Snapshots *SnapshotsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Snapshots.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Snapshots *SnapshotsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Snapshots.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Snapshots *SnapshotsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Snapshots.Contract.contract.Transact(opts, method, params...)
}
