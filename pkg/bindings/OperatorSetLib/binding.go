// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package OperatorSetLib

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

// OperatorSetLibMetaData contains all meta data concerning the OperatorSetLib contract.
var OperatorSetLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60556032600b8282823980515f1a607314602657634e487b7160e01b5f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040525f5ffdfea2646970667358221220c6f39a4279e4b76d3e2615353cdbc6c6c997db9a5b4e2de6bdd01f1036a3e64a64736f6c634300081b0033",
}

// OperatorSetLibABI is the input ABI used to generate the binding from.
// Deprecated: Use OperatorSetLibMetaData.ABI instead.
var OperatorSetLibABI = OperatorSetLibMetaData.ABI

// OperatorSetLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OperatorSetLibMetaData.Bin instead.
var OperatorSetLibBin = OperatorSetLibMetaData.Bin

// DeployOperatorSetLib deploys a new Ethereum contract, binding an instance of OperatorSetLib to it.
func DeployOperatorSetLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OperatorSetLib, error) {
	parsed, err := OperatorSetLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OperatorSetLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OperatorSetLib{OperatorSetLibCaller: OperatorSetLibCaller{contract: contract}, OperatorSetLibTransactor: OperatorSetLibTransactor{contract: contract}, OperatorSetLibFilterer: OperatorSetLibFilterer{contract: contract}}, nil
}

// OperatorSetLib is an auto generated Go binding around an Ethereum contract.
type OperatorSetLib struct {
	OperatorSetLibCaller     // Read-only binding to the contract
	OperatorSetLibTransactor // Write-only binding to the contract
	OperatorSetLibFilterer   // Log filterer for contract events
}

// OperatorSetLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type OperatorSetLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OperatorSetLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OperatorSetLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OperatorSetLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OperatorSetLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OperatorSetLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OperatorSetLibSession struct {
	Contract     *OperatorSetLib   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OperatorSetLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OperatorSetLibCallerSession struct {
	Contract *OperatorSetLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// OperatorSetLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OperatorSetLibTransactorSession struct {
	Contract     *OperatorSetLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// OperatorSetLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type OperatorSetLibRaw struct {
	Contract *OperatorSetLib // Generic contract binding to access the raw methods on
}

// OperatorSetLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OperatorSetLibCallerRaw struct {
	Contract *OperatorSetLibCaller // Generic read-only contract binding to access the raw methods on
}

// OperatorSetLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OperatorSetLibTransactorRaw struct {
	Contract *OperatorSetLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOperatorSetLib creates a new instance of OperatorSetLib, bound to a specific deployed contract.
func NewOperatorSetLib(address common.Address, backend bind.ContractBackend) (*OperatorSetLib, error) {
	contract, err := bindOperatorSetLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OperatorSetLib{OperatorSetLibCaller: OperatorSetLibCaller{contract: contract}, OperatorSetLibTransactor: OperatorSetLibTransactor{contract: contract}, OperatorSetLibFilterer: OperatorSetLibFilterer{contract: contract}}, nil
}

// NewOperatorSetLibCaller creates a new read-only instance of OperatorSetLib, bound to a specific deployed contract.
func NewOperatorSetLibCaller(address common.Address, caller bind.ContractCaller) (*OperatorSetLibCaller, error) {
	contract, err := bindOperatorSetLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OperatorSetLibCaller{contract: contract}, nil
}

// NewOperatorSetLibTransactor creates a new write-only instance of OperatorSetLib, bound to a specific deployed contract.
func NewOperatorSetLibTransactor(address common.Address, transactor bind.ContractTransactor) (*OperatorSetLibTransactor, error) {
	contract, err := bindOperatorSetLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OperatorSetLibTransactor{contract: contract}, nil
}

// NewOperatorSetLibFilterer creates a new log filterer instance of OperatorSetLib, bound to a specific deployed contract.
func NewOperatorSetLibFilterer(address common.Address, filterer bind.ContractFilterer) (*OperatorSetLibFilterer, error) {
	contract, err := bindOperatorSetLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OperatorSetLibFilterer{contract: contract}, nil
}

// bindOperatorSetLib binds a generic wrapper to an already deployed contract.
func bindOperatorSetLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OperatorSetLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OperatorSetLib *OperatorSetLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OperatorSetLib.Contract.OperatorSetLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OperatorSetLib *OperatorSetLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OperatorSetLib.Contract.OperatorSetLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OperatorSetLib *OperatorSetLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OperatorSetLib.Contract.OperatorSetLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OperatorSetLib *OperatorSetLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OperatorSetLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OperatorSetLib *OperatorSetLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OperatorSetLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OperatorSetLib *OperatorSetLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OperatorSetLib.Contract.contract.Transact(opts, method, params...)
}
