// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Merkle

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

// MerkleMetaData contains all meta data concerning the Merkle contract.
var MerkleMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220de4dce1b90fa553d4573afc57ce3ea0458378cac0669dc216b6be8ae471c99ea64736f6c634300080c0033",
}

// MerkleABI is the input ABI used to generate the binding from.
// Deprecated: Use MerkleMetaData.ABI instead.
var MerkleABI = MerkleMetaData.ABI

// MerkleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MerkleMetaData.Bin instead.
var MerkleBin = MerkleMetaData.Bin

// DeployMerkle deploys a new Ethereum contract, binding an instance of Merkle to it.
func DeployMerkle(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Merkle, error) {
	parsed, err := MerkleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MerkleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Merkle{MerkleCaller: MerkleCaller{contract: contract}, MerkleTransactor: MerkleTransactor{contract: contract}, MerkleFilterer: MerkleFilterer{contract: contract}}, nil
}

// Merkle is an auto generated Go binding around an Ethereum contract.
type Merkle struct {
	MerkleCaller     // Read-only binding to the contract
	MerkleTransactor // Write-only binding to the contract
	MerkleFilterer   // Log filterer for contract events
}

// MerkleCaller is an auto generated read-only Go binding around an Ethereum contract.
type MerkleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MerkleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MerkleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MerkleSession struct {
	Contract     *Merkle           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MerkleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MerkleCallerSession struct {
	Contract *MerkleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MerkleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MerkleTransactorSession struct {
	Contract     *MerkleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MerkleRaw is an auto generated low-level Go binding around an Ethereum contract.
type MerkleRaw struct {
	Contract *Merkle // Generic contract binding to access the raw methods on
}

// MerkleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MerkleCallerRaw struct {
	Contract *MerkleCaller // Generic read-only contract binding to access the raw methods on
}

// MerkleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MerkleTransactorRaw struct {
	Contract *MerkleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMerkle creates a new instance of Merkle, bound to a specific deployed contract.
func NewMerkle(address common.Address, backend bind.ContractBackend) (*Merkle, error) {
	contract, err := bindMerkle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Merkle{MerkleCaller: MerkleCaller{contract: contract}, MerkleTransactor: MerkleTransactor{contract: contract}, MerkleFilterer: MerkleFilterer{contract: contract}}, nil
}

// NewMerkleCaller creates a new read-only instance of Merkle, bound to a specific deployed contract.
func NewMerkleCaller(address common.Address, caller bind.ContractCaller) (*MerkleCaller, error) {
	contract, err := bindMerkle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleCaller{contract: contract}, nil
}

// NewMerkleTransactor creates a new write-only instance of Merkle, bound to a specific deployed contract.
func NewMerkleTransactor(address common.Address, transactor bind.ContractTransactor) (*MerkleTransactor, error) {
	contract, err := bindMerkle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleTransactor{contract: contract}, nil
}

// NewMerkleFilterer creates a new log filterer instance of Merkle, bound to a specific deployed contract.
func NewMerkleFilterer(address common.Address, filterer bind.ContractFilterer) (*MerkleFilterer, error) {
	contract, err := bindMerkle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MerkleFilterer{contract: contract}, nil
}

// bindMerkle binds a generic wrapper to an already deployed contract.
func bindMerkle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MerkleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Merkle *MerkleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Merkle.Contract.MerkleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Merkle *MerkleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Merkle.Contract.MerkleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Merkle *MerkleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Merkle.Contract.MerkleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Merkle *MerkleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Merkle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Merkle *MerkleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Merkle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Merkle *MerkleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Merkle.Contract.contract.Transact(opts, method, params...)
}
