// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IAVSRegistrar

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

// IAVSRegistrarMetaData contains all meta data concerning the IAVSRegistrar contract.
var IAVSRegistrarMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"deregisterOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsAVS\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"}]",
}

// IAVSRegistrarABI is the input ABI used to generate the binding from.
// Deprecated: Use IAVSRegistrarMetaData.ABI instead.
var IAVSRegistrarABI = IAVSRegistrarMetaData.ABI

// IAVSRegistrar is an auto generated Go binding around an Ethereum contract.
type IAVSRegistrar struct {
	IAVSRegistrarCaller     // Read-only binding to the contract
	IAVSRegistrarTransactor // Write-only binding to the contract
	IAVSRegistrarFilterer   // Log filterer for contract events
}

// IAVSRegistrarCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAVSRegistrarCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAVSRegistrarTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAVSRegistrarTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAVSRegistrarFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAVSRegistrarFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAVSRegistrarSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAVSRegistrarSession struct {
	Contract     *IAVSRegistrar    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAVSRegistrarCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAVSRegistrarCallerSession struct {
	Contract *IAVSRegistrarCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IAVSRegistrarTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAVSRegistrarTransactorSession struct {
	Contract     *IAVSRegistrarTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IAVSRegistrarRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAVSRegistrarRaw struct {
	Contract *IAVSRegistrar // Generic contract binding to access the raw methods on
}

// IAVSRegistrarCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAVSRegistrarCallerRaw struct {
	Contract *IAVSRegistrarCaller // Generic read-only contract binding to access the raw methods on
}

// IAVSRegistrarTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAVSRegistrarTransactorRaw struct {
	Contract *IAVSRegistrarTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAVSRegistrar creates a new instance of IAVSRegistrar, bound to a specific deployed contract.
func NewIAVSRegistrar(address common.Address, backend bind.ContractBackend) (*IAVSRegistrar, error) {
	contract, err := bindIAVSRegistrar(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAVSRegistrar{IAVSRegistrarCaller: IAVSRegistrarCaller{contract: contract}, IAVSRegistrarTransactor: IAVSRegistrarTransactor{contract: contract}, IAVSRegistrarFilterer: IAVSRegistrarFilterer{contract: contract}}, nil
}

// NewIAVSRegistrarCaller creates a new read-only instance of IAVSRegistrar, bound to a specific deployed contract.
func NewIAVSRegistrarCaller(address common.Address, caller bind.ContractCaller) (*IAVSRegistrarCaller, error) {
	contract, err := bindIAVSRegistrar(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAVSRegistrarCaller{contract: contract}, nil
}

// NewIAVSRegistrarTransactor creates a new write-only instance of IAVSRegistrar, bound to a specific deployed contract.
func NewIAVSRegistrarTransactor(address common.Address, transactor bind.ContractTransactor) (*IAVSRegistrarTransactor, error) {
	contract, err := bindIAVSRegistrar(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAVSRegistrarTransactor{contract: contract}, nil
}

// NewIAVSRegistrarFilterer creates a new log filterer instance of IAVSRegistrar, bound to a specific deployed contract.
func NewIAVSRegistrarFilterer(address common.Address, filterer bind.ContractFilterer) (*IAVSRegistrarFilterer, error) {
	contract, err := bindIAVSRegistrar(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAVSRegistrarFilterer{contract: contract}, nil
}

// bindIAVSRegistrar binds a generic wrapper to an already deployed contract.
func bindIAVSRegistrar(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IAVSRegistrarMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAVSRegistrar *IAVSRegistrarRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAVSRegistrar.Contract.IAVSRegistrarCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAVSRegistrar *IAVSRegistrarRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAVSRegistrar.Contract.IAVSRegistrarTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAVSRegistrar *IAVSRegistrarRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAVSRegistrar.Contract.IAVSRegistrarTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAVSRegistrar *IAVSRegistrarCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAVSRegistrar.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAVSRegistrar *IAVSRegistrarTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAVSRegistrar.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAVSRegistrar *IAVSRegistrarTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAVSRegistrar.Contract.contract.Transact(opts, method, params...)
}

// SupportsAVS is a free data retrieval call binding the contract method 0xb5265787.
//
// Solidity: function supportsAVS(address avs) view returns(bool)
func (_IAVSRegistrar *IAVSRegistrarCaller) SupportsAVS(opts *bind.CallOpts, avs common.Address) (bool, error) {
	var out []interface{}
	err := _IAVSRegistrar.contract.Call(opts, &out, "supportsAVS", avs)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsAVS is a free data retrieval call binding the contract method 0xb5265787.
//
// Solidity: function supportsAVS(address avs) view returns(bool)
func (_IAVSRegistrar *IAVSRegistrarSession) SupportsAVS(avs common.Address) (bool, error) {
	return _IAVSRegistrar.Contract.SupportsAVS(&_IAVSRegistrar.CallOpts, avs)
}

// SupportsAVS is a free data retrieval call binding the contract method 0xb5265787.
//
// Solidity: function supportsAVS(address avs) view returns(bool)
func (_IAVSRegistrar *IAVSRegistrarCallerSession) SupportsAVS(avs common.Address) (bool, error) {
	return _IAVSRegistrar.Contract.SupportsAVS(&_IAVSRegistrar.CallOpts, avs)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0x303ca956.
//
// Solidity: function deregisterOperator(address operator, address avs, uint32[] operatorSetIds) returns()
func (_IAVSRegistrar *IAVSRegistrarTransactor) DeregisterOperator(opts *bind.TransactOpts, operator common.Address, avs common.Address, operatorSetIds []uint32) (*types.Transaction, error) {
	return _IAVSRegistrar.contract.Transact(opts, "deregisterOperator", operator, avs, operatorSetIds)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0x303ca956.
//
// Solidity: function deregisterOperator(address operator, address avs, uint32[] operatorSetIds) returns()
func (_IAVSRegistrar *IAVSRegistrarSession) DeregisterOperator(operator common.Address, avs common.Address, operatorSetIds []uint32) (*types.Transaction, error) {
	return _IAVSRegistrar.Contract.DeregisterOperator(&_IAVSRegistrar.TransactOpts, operator, avs, operatorSetIds)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0x303ca956.
//
// Solidity: function deregisterOperator(address operator, address avs, uint32[] operatorSetIds) returns()
func (_IAVSRegistrar *IAVSRegistrarTransactorSession) DeregisterOperator(operator common.Address, avs common.Address, operatorSetIds []uint32) (*types.Transaction, error) {
	return _IAVSRegistrar.Contract.DeregisterOperator(&_IAVSRegistrar.TransactOpts, operator, avs, operatorSetIds)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xc63fd502.
//
// Solidity: function registerOperator(address operator, address avs, uint32[] operatorSetIds, bytes data) returns()
func (_IAVSRegistrar *IAVSRegistrarTransactor) RegisterOperator(opts *bind.TransactOpts, operator common.Address, avs common.Address, operatorSetIds []uint32, data []byte) (*types.Transaction, error) {
	return _IAVSRegistrar.contract.Transact(opts, "registerOperator", operator, avs, operatorSetIds, data)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xc63fd502.
//
// Solidity: function registerOperator(address operator, address avs, uint32[] operatorSetIds, bytes data) returns()
func (_IAVSRegistrar *IAVSRegistrarSession) RegisterOperator(operator common.Address, avs common.Address, operatorSetIds []uint32, data []byte) (*types.Transaction, error) {
	return _IAVSRegistrar.Contract.RegisterOperator(&_IAVSRegistrar.TransactOpts, operator, avs, operatorSetIds, data)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xc63fd502.
//
// Solidity: function registerOperator(address operator, address avs, uint32[] operatorSetIds, bytes data) returns()
func (_IAVSRegistrar *IAVSRegistrarTransactorSession) RegisterOperator(operator common.Address, avs common.Address, operatorSetIds []uint32, data []byte) (*types.Transaction, error) {
	return _IAVSRegistrar.Contract.RegisterOperator(&_IAVSRegistrar.TransactOpts, operator, avs, operatorSetIds, data)
}
