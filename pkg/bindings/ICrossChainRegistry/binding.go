// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ICrossChainRegistry

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

// OperatorSet is an auto generated low-level Go binding around an user-defined struct.
type OperatorSet struct {
	Avs common.Address
	Id  uint32
}

// ICrossChainRegistryMetaData contains all meta data concerning the ICrossChainRegistry contract.
var ICrossChainRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"addChainIDToWhitelist\",\"inputs\":[{\"name\":\"chainID\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addTransportDestination\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"chainID\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getActiveGenerationReservations\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"contractIOperatorTableCalculator[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getOperatorTableCalculator\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIOperatorTableCalculator\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTransportDestinations\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeChainIDFromWhitelist\",\"inputs\":[{\"name\":\"chainID\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeGenerationReservation\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeTransportDestination\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"chainID\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requestGenerationReservation\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operatorTableCalculator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOperatorTableCalculator\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"calculator\",\"type\":\"address\",\"internalType\":\"contractIOperatorTableCalculator\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ChainIDAddedToWhitelist\",\"inputs\":[{\"name\":\"chainID\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainIDRemovedFromWhitelist\",\"inputs\":[{\"name\":\"chainID\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GenerationReservationMade\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operatorTableCalculator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIOperatorTableCalculator\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GenerationReservationRemoved\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operatorTableCalculator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIOperatorTableCalculator\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TransportDestinationAdded\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"chainID\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TransportDestinationRemoved\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"chainID\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidChainId\",\"inputs\":[]}]",
}

// ICrossChainRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ICrossChainRegistryMetaData.ABI instead.
var ICrossChainRegistryABI = ICrossChainRegistryMetaData.ABI

// ICrossChainRegistry is an auto generated Go binding around an Ethereum contract.
type ICrossChainRegistry struct {
	ICrossChainRegistryCaller     // Read-only binding to the contract
	ICrossChainRegistryTransactor // Write-only binding to the contract
	ICrossChainRegistryFilterer   // Log filterer for contract events
}

// ICrossChainRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ICrossChainRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICrossChainRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ICrossChainRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICrossChainRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ICrossChainRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICrossChainRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ICrossChainRegistrySession struct {
	Contract     *ICrossChainRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ICrossChainRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ICrossChainRegistryCallerSession struct {
	Contract *ICrossChainRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// ICrossChainRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ICrossChainRegistryTransactorSession struct {
	Contract     *ICrossChainRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// ICrossChainRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ICrossChainRegistryRaw struct {
	Contract *ICrossChainRegistry // Generic contract binding to access the raw methods on
}

// ICrossChainRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ICrossChainRegistryCallerRaw struct {
	Contract *ICrossChainRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ICrossChainRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ICrossChainRegistryTransactorRaw struct {
	Contract *ICrossChainRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewICrossChainRegistry creates a new instance of ICrossChainRegistry, bound to a specific deployed contract.
func NewICrossChainRegistry(address common.Address, backend bind.ContractBackend) (*ICrossChainRegistry, error) {
	contract, err := bindICrossChainRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ICrossChainRegistry{ICrossChainRegistryCaller: ICrossChainRegistryCaller{contract: contract}, ICrossChainRegistryTransactor: ICrossChainRegistryTransactor{contract: contract}, ICrossChainRegistryFilterer: ICrossChainRegistryFilterer{contract: contract}}, nil
}

// NewICrossChainRegistryCaller creates a new read-only instance of ICrossChainRegistry, bound to a specific deployed contract.
func NewICrossChainRegistryCaller(address common.Address, caller bind.ContractCaller) (*ICrossChainRegistryCaller, error) {
	contract, err := bindICrossChainRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ICrossChainRegistryCaller{contract: contract}, nil
}

// NewICrossChainRegistryTransactor creates a new write-only instance of ICrossChainRegistry, bound to a specific deployed contract.
func NewICrossChainRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ICrossChainRegistryTransactor, error) {
	contract, err := bindICrossChainRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ICrossChainRegistryTransactor{contract: contract}, nil
}

// NewICrossChainRegistryFilterer creates a new log filterer instance of ICrossChainRegistry, bound to a specific deployed contract.
func NewICrossChainRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ICrossChainRegistryFilterer, error) {
	contract, err := bindICrossChainRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ICrossChainRegistryFilterer{contract: contract}, nil
}

// bindICrossChainRegistry binds a generic wrapper to an already deployed contract.
func bindICrossChainRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ICrossChainRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICrossChainRegistry *ICrossChainRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICrossChainRegistry.Contract.ICrossChainRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICrossChainRegistry *ICrossChainRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICrossChainRegistry.Contract.ICrossChainRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICrossChainRegistry *ICrossChainRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICrossChainRegistry.Contract.ICrossChainRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICrossChainRegistry *ICrossChainRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICrossChainRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICrossChainRegistry *ICrossChainRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICrossChainRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICrossChainRegistry *ICrossChainRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICrossChainRegistry.Contract.contract.Transact(opts, method, params...)
}

// GetSupportedChains is a free data retrieval call binding the contract method 0xc4bffe2b.
//
// Solidity: function getSupportedChains() view returns(uint32[])
func (_ICrossChainRegistry *ICrossChainRegistryCaller) GetSupportedChains(opts *bind.CallOpts) ([]uint32, error) {
	var out []interface{}
	err := _ICrossChainRegistry.contract.Call(opts, &out, "getSupportedChains")

	if err != nil {
		return *new([]uint32), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint32)).(*[]uint32)

	return out0, err

}

// GetSupportedChains is a free data retrieval call binding the contract method 0xc4bffe2b.
//
// Solidity: function getSupportedChains() view returns(uint32[])
func (_ICrossChainRegistry *ICrossChainRegistrySession) GetSupportedChains() ([]uint32, error) {
	return _ICrossChainRegistry.Contract.GetSupportedChains(&_ICrossChainRegistry.CallOpts)
}

// GetSupportedChains is a free data retrieval call binding the contract method 0xc4bffe2b.
//
// Solidity: function getSupportedChains() view returns(uint32[])
func (_ICrossChainRegistry *ICrossChainRegistryCallerSession) GetSupportedChains() ([]uint32, error) {
	return _ICrossChainRegistry.Contract.GetSupportedChains(&_ICrossChainRegistry.CallOpts)
}

// GetTransportDestinations is a free data retrieval call binding the contract method 0x3c75fddf.
//
// Solidity: function getTransportDestinations((address,uint32) operatorSet) view returns(uint32[])
func (_ICrossChainRegistry *ICrossChainRegistryCaller) GetTransportDestinations(opts *bind.CallOpts, operatorSet OperatorSet) ([]uint32, error) {
	var out []interface{}
	err := _ICrossChainRegistry.contract.Call(opts, &out, "getTransportDestinations", operatorSet)

	if err != nil {
		return *new([]uint32), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint32)).(*[]uint32)

	return out0, err

}

// GetTransportDestinations is a free data retrieval call binding the contract method 0x3c75fddf.
//
// Solidity: function getTransportDestinations((address,uint32) operatorSet) view returns(uint32[])
func (_ICrossChainRegistry *ICrossChainRegistrySession) GetTransportDestinations(operatorSet OperatorSet) ([]uint32, error) {
	return _ICrossChainRegistry.Contract.GetTransportDestinations(&_ICrossChainRegistry.CallOpts, operatorSet)
}

// GetTransportDestinations is a free data retrieval call binding the contract method 0x3c75fddf.
//
// Solidity: function getTransportDestinations((address,uint32) operatorSet) view returns(uint32[])
func (_ICrossChainRegistry *ICrossChainRegistryCallerSession) GetTransportDestinations(operatorSet OperatorSet) ([]uint32, error) {
	return _ICrossChainRegistry.Contract.GetTransportDestinations(&_ICrossChainRegistry.CallOpts, operatorSet)
}

// AddChainIDToWhitelist is a paid mutator transaction binding the contract method 0x16df80b5.
//
// Solidity: function addChainIDToWhitelist(uint32 chainID) returns()
func (_ICrossChainRegistry *ICrossChainRegistryTransactor) AddChainIDToWhitelist(opts *bind.TransactOpts, chainID uint32) (*types.Transaction, error) {
	return _ICrossChainRegistry.contract.Transact(opts, "addChainIDToWhitelist", chainID)
}

// AddChainIDToWhitelist is a paid mutator transaction binding the contract method 0x16df80b5.
//
// Solidity: function addChainIDToWhitelist(uint32 chainID) returns()
func (_ICrossChainRegistry *ICrossChainRegistrySession) AddChainIDToWhitelist(chainID uint32) (*types.Transaction, error) {
	return _ICrossChainRegistry.Contract.AddChainIDToWhitelist(&_ICrossChainRegistry.TransactOpts, chainID)
}

// AddChainIDToWhitelist is a paid mutator transaction binding the contract method 0x16df80b5.
//
// Solidity: function addChainIDToWhitelist(uint32 chainID) returns()
func (_ICrossChainRegistry *ICrossChainRegistryTransactorSession) AddChainIDToWhitelist(chainID uint32) (*types.Transaction, error) {
	return _ICrossChainRegistry.Contract.AddChainIDToWhitelist(&_ICrossChainRegistry.TransactOpts, chainID)
}

// AddTransportDestination is a paid mutator transaction binding the contract method 0x814e731b.
//
// Solidity: function addTransportDestination((address,uint32) operatorSet, uint32 chainID) returns()
func (_ICrossChainRegistry *ICrossChainRegistryTransactor) AddTransportDestination(opts *bind.TransactOpts, operatorSet OperatorSet, chainID uint32) (*types.Transaction, error) {
	return _ICrossChainRegistry.contract.Transact(opts, "addTransportDestination", operatorSet, chainID)
}

// AddTransportDestination is a paid mutator transaction binding the contract method 0x814e731b.
//
// Solidity: function addTransportDestination((address,uint32) operatorSet, uint32 chainID) returns()
func (_ICrossChainRegistry *ICrossChainRegistrySession) AddTransportDestination(operatorSet OperatorSet, chainID uint32) (*types.Transaction, error) {
	return _ICrossChainRegistry.Contract.AddTransportDestination(&_ICrossChainRegistry.TransactOpts, operatorSet, chainID)
}

// AddTransportDestination is a paid mutator transaction binding the contract method 0x814e731b.
//
// Solidity: function addTransportDestination((address,uint32) operatorSet, uint32 chainID) returns()
func (_ICrossChainRegistry *ICrossChainRegistryTransactorSession) AddTransportDestination(operatorSet OperatorSet, chainID uint32) (*types.Transaction, error) {
	return _ICrossChainRegistry.Contract.AddTransportDestination(&_ICrossChainRegistry.TransactOpts, operatorSet, chainID)
}

// GetActiveGenerationReservations is a paid mutator transaction binding the contract method 0xd09b978b.
//
// Solidity: function getActiveGenerationReservations() returns((address,uint32)[], address[])
func (_ICrossChainRegistry *ICrossChainRegistryTransactor) GetActiveGenerationReservations(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICrossChainRegistry.contract.Transact(opts, "getActiveGenerationReservations")
}

// GetActiveGenerationReservations is a paid mutator transaction binding the contract method 0xd09b978b.
//
// Solidity: function getActiveGenerationReservations() returns((address,uint32)[], address[])
func (_ICrossChainRegistry *ICrossChainRegistrySession) GetActiveGenerationReservations() (*types.Transaction, error) {
	return _ICrossChainRegistry.Contract.GetActiveGenerationReservations(&_ICrossChainRegistry.TransactOpts)
}

// GetActiveGenerationReservations is a paid mutator transaction binding the contract method 0xd09b978b.
//
// Solidity: function getActiveGenerationReservations() returns((address,uint32)[], address[])
func (_ICrossChainRegistry *ICrossChainRegistryTransactorSession) GetActiveGenerationReservations() (*types.Transaction, error) {
	return _ICrossChainRegistry.Contract.GetActiveGenerationReservations(&_ICrossChainRegistry.TransactOpts)
}

// GetOperatorTableCalculator is a paid mutator transaction binding the contract method 0x75e4b539.
//
// Solidity: function getOperatorTableCalculator((address,uint32) operatorSet) returns(address)
func (_ICrossChainRegistry *ICrossChainRegistryTransactor) GetOperatorTableCalculator(opts *bind.TransactOpts, operatorSet OperatorSet) (*types.Transaction, error) {
	return _ICrossChainRegistry.contract.Transact(opts, "getOperatorTableCalculator", operatorSet)
}

// GetOperatorTableCalculator is a paid mutator transaction binding the contract method 0x75e4b539.
//
// Solidity: function getOperatorTableCalculator((address,uint32) operatorSet) returns(address)
func (_ICrossChainRegistry *ICrossChainRegistrySession) GetOperatorTableCalculator(operatorSet OperatorSet) (*types.Transaction, error) {
	return _ICrossChainRegistry.Contract.GetOperatorTableCalculator(&_ICrossChainRegistry.TransactOpts, operatorSet)
}

// GetOperatorTableCalculator is a paid mutator transaction binding the contract method 0x75e4b539.
//
// Solidity: function getOperatorTableCalculator((address,uint32) operatorSet) returns(address)
func (_ICrossChainRegistry *ICrossChainRegistryTransactorSession) GetOperatorTableCalculator(operatorSet OperatorSet) (*types.Transaction, error) {
	return _ICrossChainRegistry.Contract.GetOperatorTableCalculator(&_ICrossChainRegistry.TransactOpts, operatorSet)
}

// RemoveChainIDFromWhitelist is a paid mutator transaction binding the contract method 0x2ab0b20f.
//
// Solidity: function removeChainIDFromWhitelist(uint32 chainID) returns()
func (_ICrossChainRegistry *ICrossChainRegistryTransactor) RemoveChainIDFromWhitelist(opts *bind.TransactOpts, chainID uint32) (*types.Transaction, error) {
	return _ICrossChainRegistry.contract.Transact(opts, "removeChainIDFromWhitelist", chainID)
}

// RemoveChainIDFromWhitelist is a paid mutator transaction binding the contract method 0x2ab0b20f.
//
// Solidity: function removeChainIDFromWhitelist(uint32 chainID) returns()
func (_ICrossChainRegistry *ICrossChainRegistrySession) RemoveChainIDFromWhitelist(chainID uint32) (*types.Transaction, error) {
	return _ICrossChainRegistry.Contract.RemoveChainIDFromWhitelist(&_ICrossChainRegistry.TransactOpts, chainID)
}

// RemoveChainIDFromWhitelist is a paid mutator transaction binding the contract method 0x2ab0b20f.
//
// Solidity: function removeChainIDFromWhitelist(uint32 chainID) returns()
func (_ICrossChainRegistry *ICrossChainRegistryTransactorSession) RemoveChainIDFromWhitelist(chainID uint32) (*types.Transaction, error) {
	return _ICrossChainRegistry.Contract.RemoveChainIDFromWhitelist(&_ICrossChainRegistry.TransactOpts, chainID)
}

// RemoveGenerationReservation is a paid mutator transaction binding the contract method 0x6c55a37f.
//
// Solidity: function removeGenerationReservation((address,uint32) operatorSet) returns()
func (_ICrossChainRegistry *ICrossChainRegistryTransactor) RemoveGenerationReservation(opts *bind.TransactOpts, operatorSet OperatorSet) (*types.Transaction, error) {
	return _ICrossChainRegistry.contract.Transact(opts, "removeGenerationReservation", operatorSet)
}

// RemoveGenerationReservation is a paid mutator transaction binding the contract method 0x6c55a37f.
//
// Solidity: function removeGenerationReservation((address,uint32) operatorSet) returns()
func (_ICrossChainRegistry *ICrossChainRegistrySession) RemoveGenerationReservation(operatorSet OperatorSet) (*types.Transaction, error) {
	return _ICrossChainRegistry.Contract.RemoveGenerationReservation(&_ICrossChainRegistry.TransactOpts, operatorSet)
}

// RemoveGenerationReservation is a paid mutator transaction binding the contract method 0x6c55a37f.
//
// Solidity: function removeGenerationReservation((address,uint32) operatorSet) returns()
func (_ICrossChainRegistry *ICrossChainRegistryTransactorSession) RemoveGenerationReservation(operatorSet OperatorSet) (*types.Transaction, error) {
	return _ICrossChainRegistry.Contract.RemoveGenerationReservation(&_ICrossChainRegistry.TransactOpts, operatorSet)
}

// RemoveTransportDestination is a paid mutator transaction binding the contract method 0x2132488f.
//
// Solidity: function removeTransportDestination((address,uint32) operatorSet, uint32 chainID) returns()
func (_ICrossChainRegistry *ICrossChainRegistryTransactor) RemoveTransportDestination(opts *bind.TransactOpts, operatorSet OperatorSet, chainID uint32) (*types.Transaction, error) {
	return _ICrossChainRegistry.contract.Transact(opts, "removeTransportDestination", operatorSet, chainID)
}

// RemoveTransportDestination is a paid mutator transaction binding the contract method 0x2132488f.
//
// Solidity: function removeTransportDestination((address,uint32) operatorSet, uint32 chainID) returns()
func (_ICrossChainRegistry *ICrossChainRegistrySession) RemoveTransportDestination(operatorSet OperatorSet, chainID uint32) (*types.Transaction, error) {
	return _ICrossChainRegistry.Contract.RemoveTransportDestination(&_ICrossChainRegistry.TransactOpts, operatorSet, chainID)
}

// RemoveTransportDestination is a paid mutator transaction binding the contract method 0x2132488f.
//
// Solidity: function removeTransportDestination((address,uint32) operatorSet, uint32 chainID) returns()
func (_ICrossChainRegistry *ICrossChainRegistryTransactorSession) RemoveTransportDestination(operatorSet OperatorSet, chainID uint32) (*types.Transaction, error) {
	return _ICrossChainRegistry.Contract.RemoveTransportDestination(&_ICrossChainRegistry.TransactOpts, operatorSet, chainID)
}

// RequestGenerationReservation is a paid mutator transaction binding the contract method 0x4d9d1e48.
//
// Solidity: function requestGenerationReservation((address,uint32) operatorSet, address operatorTableCalculator) returns()
func (_ICrossChainRegistry *ICrossChainRegistryTransactor) RequestGenerationReservation(opts *bind.TransactOpts, operatorSet OperatorSet, operatorTableCalculator common.Address) (*types.Transaction, error) {
	return _ICrossChainRegistry.contract.Transact(opts, "requestGenerationReservation", operatorSet, operatorTableCalculator)
}

// RequestGenerationReservation is a paid mutator transaction binding the contract method 0x4d9d1e48.
//
// Solidity: function requestGenerationReservation((address,uint32) operatorSet, address operatorTableCalculator) returns()
func (_ICrossChainRegistry *ICrossChainRegistrySession) RequestGenerationReservation(operatorSet OperatorSet, operatorTableCalculator common.Address) (*types.Transaction, error) {
	return _ICrossChainRegistry.Contract.RequestGenerationReservation(&_ICrossChainRegistry.TransactOpts, operatorSet, operatorTableCalculator)
}

// RequestGenerationReservation is a paid mutator transaction binding the contract method 0x4d9d1e48.
//
// Solidity: function requestGenerationReservation((address,uint32) operatorSet, address operatorTableCalculator) returns()
func (_ICrossChainRegistry *ICrossChainRegistryTransactorSession) RequestGenerationReservation(operatorSet OperatorSet, operatorTableCalculator common.Address) (*types.Transaction, error) {
	return _ICrossChainRegistry.Contract.RequestGenerationReservation(&_ICrossChainRegistry.TransactOpts, operatorSet, operatorTableCalculator)
}

// SetOperatorTableCalculator is a paid mutator transaction binding the contract method 0x1ca9142a.
//
// Solidity: function setOperatorTableCalculator((address,uint32) operatorSet, address calculator) returns()
func (_ICrossChainRegistry *ICrossChainRegistryTransactor) SetOperatorTableCalculator(opts *bind.TransactOpts, operatorSet OperatorSet, calculator common.Address) (*types.Transaction, error) {
	return _ICrossChainRegistry.contract.Transact(opts, "setOperatorTableCalculator", operatorSet, calculator)
}

// SetOperatorTableCalculator is a paid mutator transaction binding the contract method 0x1ca9142a.
//
// Solidity: function setOperatorTableCalculator((address,uint32) operatorSet, address calculator) returns()
func (_ICrossChainRegistry *ICrossChainRegistrySession) SetOperatorTableCalculator(operatorSet OperatorSet, calculator common.Address) (*types.Transaction, error) {
	return _ICrossChainRegistry.Contract.SetOperatorTableCalculator(&_ICrossChainRegistry.TransactOpts, operatorSet, calculator)
}

// SetOperatorTableCalculator is a paid mutator transaction binding the contract method 0x1ca9142a.
//
// Solidity: function setOperatorTableCalculator((address,uint32) operatorSet, address calculator) returns()
func (_ICrossChainRegistry *ICrossChainRegistryTransactorSession) SetOperatorTableCalculator(operatorSet OperatorSet, calculator common.Address) (*types.Transaction, error) {
	return _ICrossChainRegistry.Contract.SetOperatorTableCalculator(&_ICrossChainRegistry.TransactOpts, operatorSet, calculator)
}

// ICrossChainRegistryChainIDAddedToWhitelistIterator is returned from FilterChainIDAddedToWhitelist and is used to iterate over the raw logs and unpacked data for ChainIDAddedToWhitelist events raised by the ICrossChainRegistry contract.
type ICrossChainRegistryChainIDAddedToWhitelistIterator struct {
	Event *ICrossChainRegistryChainIDAddedToWhitelist // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ICrossChainRegistryChainIDAddedToWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICrossChainRegistryChainIDAddedToWhitelist)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ICrossChainRegistryChainIDAddedToWhitelist)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ICrossChainRegistryChainIDAddedToWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICrossChainRegistryChainIDAddedToWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICrossChainRegistryChainIDAddedToWhitelist represents a ChainIDAddedToWhitelist event raised by the ICrossChainRegistry contract.
type ICrossChainRegistryChainIDAddedToWhitelist struct {
	ChainID uint32
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterChainIDAddedToWhitelist is a free log retrieval operation binding the contract event 0x554da9a23c6204a2a98605c0fc1f29bf2988ff3f3bbfe2ac1e3c11ee1866e5b9.
//
// Solidity: event ChainIDAddedToWhitelist(uint32 chainID)
func (_ICrossChainRegistry *ICrossChainRegistryFilterer) FilterChainIDAddedToWhitelist(opts *bind.FilterOpts) (*ICrossChainRegistryChainIDAddedToWhitelistIterator, error) {

	logs, sub, err := _ICrossChainRegistry.contract.FilterLogs(opts, "ChainIDAddedToWhitelist")
	if err != nil {
		return nil, err
	}
	return &ICrossChainRegistryChainIDAddedToWhitelistIterator{contract: _ICrossChainRegistry.contract, event: "ChainIDAddedToWhitelist", logs: logs, sub: sub}, nil
}

// WatchChainIDAddedToWhitelist is a free log subscription operation binding the contract event 0x554da9a23c6204a2a98605c0fc1f29bf2988ff3f3bbfe2ac1e3c11ee1866e5b9.
//
// Solidity: event ChainIDAddedToWhitelist(uint32 chainID)
func (_ICrossChainRegistry *ICrossChainRegistryFilterer) WatchChainIDAddedToWhitelist(opts *bind.WatchOpts, sink chan<- *ICrossChainRegistryChainIDAddedToWhitelist) (event.Subscription, error) {

	logs, sub, err := _ICrossChainRegistry.contract.WatchLogs(opts, "ChainIDAddedToWhitelist")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICrossChainRegistryChainIDAddedToWhitelist)
				if err := _ICrossChainRegistry.contract.UnpackLog(event, "ChainIDAddedToWhitelist", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseChainIDAddedToWhitelist is a log parse operation binding the contract event 0x554da9a23c6204a2a98605c0fc1f29bf2988ff3f3bbfe2ac1e3c11ee1866e5b9.
//
// Solidity: event ChainIDAddedToWhitelist(uint32 chainID)
func (_ICrossChainRegistry *ICrossChainRegistryFilterer) ParseChainIDAddedToWhitelist(log types.Log) (*ICrossChainRegistryChainIDAddedToWhitelist, error) {
	event := new(ICrossChainRegistryChainIDAddedToWhitelist)
	if err := _ICrossChainRegistry.contract.UnpackLog(event, "ChainIDAddedToWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICrossChainRegistryChainIDRemovedFromWhitelistIterator is returned from FilterChainIDRemovedFromWhitelist and is used to iterate over the raw logs and unpacked data for ChainIDRemovedFromWhitelist events raised by the ICrossChainRegistry contract.
type ICrossChainRegistryChainIDRemovedFromWhitelistIterator struct {
	Event *ICrossChainRegistryChainIDRemovedFromWhitelist // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ICrossChainRegistryChainIDRemovedFromWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICrossChainRegistryChainIDRemovedFromWhitelist)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ICrossChainRegistryChainIDRemovedFromWhitelist)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ICrossChainRegistryChainIDRemovedFromWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICrossChainRegistryChainIDRemovedFromWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICrossChainRegistryChainIDRemovedFromWhitelist represents a ChainIDRemovedFromWhitelist event raised by the ICrossChainRegistry contract.
type ICrossChainRegistryChainIDRemovedFromWhitelist struct {
	ChainID uint32
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterChainIDRemovedFromWhitelist is a free log retrieval operation binding the contract event 0x2069bbe11cd6c0fc8dd1e409735f3a6309718e89b813b4d648ede2399ae37f40.
//
// Solidity: event ChainIDRemovedFromWhitelist(uint32 chainID)
func (_ICrossChainRegistry *ICrossChainRegistryFilterer) FilterChainIDRemovedFromWhitelist(opts *bind.FilterOpts) (*ICrossChainRegistryChainIDRemovedFromWhitelistIterator, error) {

	logs, sub, err := _ICrossChainRegistry.contract.FilterLogs(opts, "ChainIDRemovedFromWhitelist")
	if err != nil {
		return nil, err
	}
	return &ICrossChainRegistryChainIDRemovedFromWhitelistIterator{contract: _ICrossChainRegistry.contract, event: "ChainIDRemovedFromWhitelist", logs: logs, sub: sub}, nil
}

// WatchChainIDRemovedFromWhitelist is a free log subscription operation binding the contract event 0x2069bbe11cd6c0fc8dd1e409735f3a6309718e89b813b4d648ede2399ae37f40.
//
// Solidity: event ChainIDRemovedFromWhitelist(uint32 chainID)
func (_ICrossChainRegistry *ICrossChainRegistryFilterer) WatchChainIDRemovedFromWhitelist(opts *bind.WatchOpts, sink chan<- *ICrossChainRegistryChainIDRemovedFromWhitelist) (event.Subscription, error) {

	logs, sub, err := _ICrossChainRegistry.contract.WatchLogs(opts, "ChainIDRemovedFromWhitelist")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICrossChainRegistryChainIDRemovedFromWhitelist)
				if err := _ICrossChainRegistry.contract.UnpackLog(event, "ChainIDRemovedFromWhitelist", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseChainIDRemovedFromWhitelist is a log parse operation binding the contract event 0x2069bbe11cd6c0fc8dd1e409735f3a6309718e89b813b4d648ede2399ae37f40.
//
// Solidity: event ChainIDRemovedFromWhitelist(uint32 chainID)
func (_ICrossChainRegistry *ICrossChainRegistryFilterer) ParseChainIDRemovedFromWhitelist(log types.Log) (*ICrossChainRegistryChainIDRemovedFromWhitelist, error) {
	event := new(ICrossChainRegistryChainIDRemovedFromWhitelist)
	if err := _ICrossChainRegistry.contract.UnpackLog(event, "ChainIDRemovedFromWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICrossChainRegistryGenerationReservationMadeIterator is returned from FilterGenerationReservationMade and is used to iterate over the raw logs and unpacked data for GenerationReservationMade events raised by the ICrossChainRegistry contract.
type ICrossChainRegistryGenerationReservationMadeIterator struct {
	Event *ICrossChainRegistryGenerationReservationMade // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ICrossChainRegistryGenerationReservationMadeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICrossChainRegistryGenerationReservationMade)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ICrossChainRegistryGenerationReservationMade)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ICrossChainRegistryGenerationReservationMadeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICrossChainRegistryGenerationReservationMadeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICrossChainRegistryGenerationReservationMade represents a GenerationReservationMade event raised by the ICrossChainRegistry contract.
type ICrossChainRegistryGenerationReservationMade struct {
	OperatorSet             OperatorSet
	OperatorTableCalculator common.Address
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterGenerationReservationMade is a free log retrieval operation binding the contract event 0x14772150bf340e929cfaf20fe5e530cb3f57c1c78a7e2fe90caca8452b5d868b.
//
// Solidity: event GenerationReservationMade((address,uint32) operatorSet, address operatorTableCalculator)
func (_ICrossChainRegistry *ICrossChainRegistryFilterer) FilterGenerationReservationMade(opts *bind.FilterOpts) (*ICrossChainRegistryGenerationReservationMadeIterator, error) {

	logs, sub, err := _ICrossChainRegistry.contract.FilterLogs(opts, "GenerationReservationMade")
	if err != nil {
		return nil, err
	}
	return &ICrossChainRegistryGenerationReservationMadeIterator{contract: _ICrossChainRegistry.contract, event: "GenerationReservationMade", logs: logs, sub: sub}, nil
}

// WatchGenerationReservationMade is a free log subscription operation binding the contract event 0x14772150bf340e929cfaf20fe5e530cb3f57c1c78a7e2fe90caca8452b5d868b.
//
// Solidity: event GenerationReservationMade((address,uint32) operatorSet, address operatorTableCalculator)
func (_ICrossChainRegistry *ICrossChainRegistryFilterer) WatchGenerationReservationMade(opts *bind.WatchOpts, sink chan<- *ICrossChainRegistryGenerationReservationMade) (event.Subscription, error) {

	logs, sub, err := _ICrossChainRegistry.contract.WatchLogs(opts, "GenerationReservationMade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICrossChainRegistryGenerationReservationMade)
				if err := _ICrossChainRegistry.contract.UnpackLog(event, "GenerationReservationMade", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGenerationReservationMade is a log parse operation binding the contract event 0x14772150bf340e929cfaf20fe5e530cb3f57c1c78a7e2fe90caca8452b5d868b.
//
// Solidity: event GenerationReservationMade((address,uint32) operatorSet, address operatorTableCalculator)
func (_ICrossChainRegistry *ICrossChainRegistryFilterer) ParseGenerationReservationMade(log types.Log) (*ICrossChainRegistryGenerationReservationMade, error) {
	event := new(ICrossChainRegistryGenerationReservationMade)
	if err := _ICrossChainRegistry.contract.UnpackLog(event, "GenerationReservationMade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICrossChainRegistryGenerationReservationRemovedIterator is returned from FilterGenerationReservationRemoved and is used to iterate over the raw logs and unpacked data for GenerationReservationRemoved events raised by the ICrossChainRegistry contract.
type ICrossChainRegistryGenerationReservationRemovedIterator struct {
	Event *ICrossChainRegistryGenerationReservationRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ICrossChainRegistryGenerationReservationRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICrossChainRegistryGenerationReservationRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ICrossChainRegistryGenerationReservationRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ICrossChainRegistryGenerationReservationRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICrossChainRegistryGenerationReservationRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICrossChainRegistryGenerationReservationRemoved represents a GenerationReservationRemoved event raised by the ICrossChainRegistry contract.
type ICrossChainRegistryGenerationReservationRemoved struct {
	OperatorSet             OperatorSet
	OperatorTableCalculator common.Address
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterGenerationReservationRemoved is a free log retrieval operation binding the contract event 0x14c80ad57c80435a6a5df1a9e493d2bf1cee67dd762db709dc5363315fa224dd.
//
// Solidity: event GenerationReservationRemoved((address,uint32) operatorSet, address operatorTableCalculator)
func (_ICrossChainRegistry *ICrossChainRegistryFilterer) FilterGenerationReservationRemoved(opts *bind.FilterOpts) (*ICrossChainRegistryGenerationReservationRemovedIterator, error) {

	logs, sub, err := _ICrossChainRegistry.contract.FilterLogs(opts, "GenerationReservationRemoved")
	if err != nil {
		return nil, err
	}
	return &ICrossChainRegistryGenerationReservationRemovedIterator{contract: _ICrossChainRegistry.contract, event: "GenerationReservationRemoved", logs: logs, sub: sub}, nil
}

// WatchGenerationReservationRemoved is a free log subscription operation binding the contract event 0x14c80ad57c80435a6a5df1a9e493d2bf1cee67dd762db709dc5363315fa224dd.
//
// Solidity: event GenerationReservationRemoved((address,uint32) operatorSet, address operatorTableCalculator)
func (_ICrossChainRegistry *ICrossChainRegistryFilterer) WatchGenerationReservationRemoved(opts *bind.WatchOpts, sink chan<- *ICrossChainRegistryGenerationReservationRemoved) (event.Subscription, error) {

	logs, sub, err := _ICrossChainRegistry.contract.WatchLogs(opts, "GenerationReservationRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICrossChainRegistryGenerationReservationRemoved)
				if err := _ICrossChainRegistry.contract.UnpackLog(event, "GenerationReservationRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGenerationReservationRemoved is a log parse operation binding the contract event 0x14c80ad57c80435a6a5df1a9e493d2bf1cee67dd762db709dc5363315fa224dd.
//
// Solidity: event GenerationReservationRemoved((address,uint32) operatorSet, address operatorTableCalculator)
func (_ICrossChainRegistry *ICrossChainRegistryFilterer) ParseGenerationReservationRemoved(log types.Log) (*ICrossChainRegistryGenerationReservationRemoved, error) {
	event := new(ICrossChainRegistryGenerationReservationRemoved)
	if err := _ICrossChainRegistry.contract.UnpackLog(event, "GenerationReservationRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICrossChainRegistryTransportDestinationAddedIterator is returned from FilterTransportDestinationAdded and is used to iterate over the raw logs and unpacked data for TransportDestinationAdded events raised by the ICrossChainRegistry contract.
type ICrossChainRegistryTransportDestinationAddedIterator struct {
	Event *ICrossChainRegistryTransportDestinationAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ICrossChainRegistryTransportDestinationAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICrossChainRegistryTransportDestinationAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ICrossChainRegistryTransportDestinationAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ICrossChainRegistryTransportDestinationAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICrossChainRegistryTransportDestinationAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICrossChainRegistryTransportDestinationAdded represents a TransportDestinationAdded event raised by the ICrossChainRegistry contract.
type ICrossChainRegistryTransportDestinationAdded struct {
	OperatorSet OperatorSet
	ChainID     uint32
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTransportDestinationAdded is a free log retrieval operation binding the contract event 0x8b4440933d21e00b2cb9054306b21777f3da3f19c8fd32d61bb225c02d2b9877.
//
// Solidity: event TransportDestinationAdded((address,uint32) operatorSet, uint32 chainID)
func (_ICrossChainRegistry *ICrossChainRegistryFilterer) FilterTransportDestinationAdded(opts *bind.FilterOpts) (*ICrossChainRegistryTransportDestinationAddedIterator, error) {

	logs, sub, err := _ICrossChainRegistry.contract.FilterLogs(opts, "TransportDestinationAdded")
	if err != nil {
		return nil, err
	}
	return &ICrossChainRegistryTransportDestinationAddedIterator{contract: _ICrossChainRegistry.contract, event: "TransportDestinationAdded", logs: logs, sub: sub}, nil
}

// WatchTransportDestinationAdded is a free log subscription operation binding the contract event 0x8b4440933d21e00b2cb9054306b21777f3da3f19c8fd32d61bb225c02d2b9877.
//
// Solidity: event TransportDestinationAdded((address,uint32) operatorSet, uint32 chainID)
func (_ICrossChainRegistry *ICrossChainRegistryFilterer) WatchTransportDestinationAdded(opts *bind.WatchOpts, sink chan<- *ICrossChainRegistryTransportDestinationAdded) (event.Subscription, error) {

	logs, sub, err := _ICrossChainRegistry.contract.WatchLogs(opts, "TransportDestinationAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICrossChainRegistryTransportDestinationAdded)
				if err := _ICrossChainRegistry.contract.UnpackLog(event, "TransportDestinationAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransportDestinationAdded is a log parse operation binding the contract event 0x8b4440933d21e00b2cb9054306b21777f3da3f19c8fd32d61bb225c02d2b9877.
//
// Solidity: event TransportDestinationAdded((address,uint32) operatorSet, uint32 chainID)
func (_ICrossChainRegistry *ICrossChainRegistryFilterer) ParseTransportDestinationAdded(log types.Log) (*ICrossChainRegistryTransportDestinationAdded, error) {
	event := new(ICrossChainRegistryTransportDestinationAdded)
	if err := _ICrossChainRegistry.contract.UnpackLog(event, "TransportDestinationAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICrossChainRegistryTransportDestinationRemovedIterator is returned from FilterTransportDestinationRemoved and is used to iterate over the raw logs and unpacked data for TransportDestinationRemoved events raised by the ICrossChainRegistry contract.
type ICrossChainRegistryTransportDestinationRemovedIterator struct {
	Event *ICrossChainRegistryTransportDestinationRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ICrossChainRegistryTransportDestinationRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICrossChainRegistryTransportDestinationRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ICrossChainRegistryTransportDestinationRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ICrossChainRegistryTransportDestinationRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICrossChainRegistryTransportDestinationRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICrossChainRegistryTransportDestinationRemoved represents a TransportDestinationRemoved event raised by the ICrossChainRegistry contract.
type ICrossChainRegistryTransportDestinationRemoved struct {
	OperatorSet OperatorSet
	ChainID     uint32
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTransportDestinationRemoved is a free log retrieval operation binding the contract event 0x6492ce451b3fd0b7d79c6695d244676a8ee605aa63dab2ec2382796695d9461c.
//
// Solidity: event TransportDestinationRemoved((address,uint32) operatorSet, uint32 chainID)
func (_ICrossChainRegistry *ICrossChainRegistryFilterer) FilterTransportDestinationRemoved(opts *bind.FilterOpts) (*ICrossChainRegistryTransportDestinationRemovedIterator, error) {

	logs, sub, err := _ICrossChainRegistry.contract.FilterLogs(opts, "TransportDestinationRemoved")
	if err != nil {
		return nil, err
	}
	return &ICrossChainRegistryTransportDestinationRemovedIterator{contract: _ICrossChainRegistry.contract, event: "TransportDestinationRemoved", logs: logs, sub: sub}, nil
}

// WatchTransportDestinationRemoved is a free log subscription operation binding the contract event 0x6492ce451b3fd0b7d79c6695d244676a8ee605aa63dab2ec2382796695d9461c.
//
// Solidity: event TransportDestinationRemoved((address,uint32) operatorSet, uint32 chainID)
func (_ICrossChainRegistry *ICrossChainRegistryFilterer) WatchTransportDestinationRemoved(opts *bind.WatchOpts, sink chan<- *ICrossChainRegistryTransportDestinationRemoved) (event.Subscription, error) {

	logs, sub, err := _ICrossChainRegistry.contract.WatchLogs(opts, "TransportDestinationRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICrossChainRegistryTransportDestinationRemoved)
				if err := _ICrossChainRegistry.contract.UnpackLog(event, "TransportDestinationRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransportDestinationRemoved is a log parse operation binding the contract event 0x6492ce451b3fd0b7d79c6695d244676a8ee605aa63dab2ec2382796695d9461c.
//
// Solidity: event TransportDestinationRemoved((address,uint32) operatorSet, uint32 chainID)
func (_ICrossChainRegistry *ICrossChainRegistryFilterer) ParseTransportDestinationRemoved(log types.Log) (*ICrossChainRegistryTransportDestinationRemoved, error) {
	event := new(ICrossChainRegistryTransportDestinationRemoved)
	if err := _ICrossChainRegistry.contract.UnpackLog(event, "TransportDestinationRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
