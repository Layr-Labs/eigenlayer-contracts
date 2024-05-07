// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package StructuredLinkedList

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

// StructuredLinkedListMetaData contains all meta data concerning the StructuredLinkedList contract.
var StructuredLinkedListMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212200b0f3aa07997f25ddb2b0dd17a1b6104304d811d2cec7f2d2dfcae3e3d618f6964736f6c634300080c0033",
}

// StructuredLinkedListABI is the input ABI used to generate the binding from.
// Deprecated: Use StructuredLinkedListMetaData.ABI instead.
var StructuredLinkedListABI = StructuredLinkedListMetaData.ABI

// StructuredLinkedListBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StructuredLinkedListMetaData.Bin instead.
var StructuredLinkedListBin = StructuredLinkedListMetaData.Bin

// DeployStructuredLinkedList deploys a new Ethereum contract, binding an instance of StructuredLinkedList to it.
func DeployStructuredLinkedList(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StructuredLinkedList, error) {
	parsed, err := StructuredLinkedListMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StructuredLinkedListBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StructuredLinkedList{StructuredLinkedListCaller: StructuredLinkedListCaller{contract: contract}, StructuredLinkedListTransactor: StructuredLinkedListTransactor{contract: contract}, StructuredLinkedListFilterer: StructuredLinkedListFilterer{contract: contract}}, nil
}

// StructuredLinkedList is an auto generated Go binding around an Ethereum contract.
type StructuredLinkedList struct {
	StructuredLinkedListCaller     // Read-only binding to the contract
	StructuredLinkedListTransactor // Write-only binding to the contract
	StructuredLinkedListFilterer   // Log filterer for contract events
}

// StructuredLinkedListCaller is an auto generated read-only Go binding around an Ethereum contract.
type StructuredLinkedListCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StructuredLinkedListTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StructuredLinkedListTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StructuredLinkedListFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StructuredLinkedListFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StructuredLinkedListSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StructuredLinkedListSession struct {
	Contract     *StructuredLinkedList // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// StructuredLinkedListCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StructuredLinkedListCallerSession struct {
	Contract *StructuredLinkedListCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// StructuredLinkedListTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StructuredLinkedListTransactorSession struct {
	Contract     *StructuredLinkedListTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// StructuredLinkedListRaw is an auto generated low-level Go binding around an Ethereum contract.
type StructuredLinkedListRaw struct {
	Contract *StructuredLinkedList // Generic contract binding to access the raw methods on
}

// StructuredLinkedListCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StructuredLinkedListCallerRaw struct {
	Contract *StructuredLinkedListCaller // Generic read-only contract binding to access the raw methods on
}

// StructuredLinkedListTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StructuredLinkedListTransactorRaw struct {
	Contract *StructuredLinkedListTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStructuredLinkedList creates a new instance of StructuredLinkedList, bound to a specific deployed contract.
func NewStructuredLinkedList(address common.Address, backend bind.ContractBackend) (*StructuredLinkedList, error) {
	contract, err := bindStructuredLinkedList(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StructuredLinkedList{StructuredLinkedListCaller: StructuredLinkedListCaller{contract: contract}, StructuredLinkedListTransactor: StructuredLinkedListTransactor{contract: contract}, StructuredLinkedListFilterer: StructuredLinkedListFilterer{contract: contract}}, nil
}

// NewStructuredLinkedListCaller creates a new read-only instance of StructuredLinkedList, bound to a specific deployed contract.
func NewStructuredLinkedListCaller(address common.Address, caller bind.ContractCaller) (*StructuredLinkedListCaller, error) {
	contract, err := bindStructuredLinkedList(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StructuredLinkedListCaller{contract: contract}, nil
}

// NewStructuredLinkedListTransactor creates a new write-only instance of StructuredLinkedList, bound to a specific deployed contract.
func NewStructuredLinkedListTransactor(address common.Address, transactor bind.ContractTransactor) (*StructuredLinkedListTransactor, error) {
	contract, err := bindStructuredLinkedList(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StructuredLinkedListTransactor{contract: contract}, nil
}

// NewStructuredLinkedListFilterer creates a new log filterer instance of StructuredLinkedList, bound to a specific deployed contract.
func NewStructuredLinkedListFilterer(address common.Address, filterer bind.ContractFilterer) (*StructuredLinkedListFilterer, error) {
	contract, err := bindStructuredLinkedList(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StructuredLinkedListFilterer{contract: contract}, nil
}

// bindStructuredLinkedList binds a generic wrapper to an already deployed contract.
func bindStructuredLinkedList(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StructuredLinkedListMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StructuredLinkedList *StructuredLinkedListRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StructuredLinkedList.Contract.StructuredLinkedListCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StructuredLinkedList *StructuredLinkedListRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StructuredLinkedList.Contract.StructuredLinkedListTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StructuredLinkedList *StructuredLinkedListRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StructuredLinkedList.Contract.StructuredLinkedListTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StructuredLinkedList *StructuredLinkedListCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StructuredLinkedList.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StructuredLinkedList *StructuredLinkedListTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StructuredLinkedList.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StructuredLinkedList *StructuredLinkedListTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StructuredLinkedList.Contract.contract.Transact(opts, method, params...)
}
