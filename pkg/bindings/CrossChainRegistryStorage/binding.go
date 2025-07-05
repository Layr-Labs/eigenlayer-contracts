// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package CrossChainRegistryStorage

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

// ICrossChainRegistryTypesOperatorSetConfig is an auto generated low-level Go binding around an user-defined struct.
type ICrossChainRegistryTypesOperatorSetConfig struct {
	Owner              common.Address
	MaxStalenessPeriod uint32
}

// OperatorSet is an auto generated low-level Go binding around an user-defined struct.
type OperatorSet struct {
	Avs common.Address
	Id  uint32
}

// CrossChainRegistryStorageMetaData contains all meta data concerning the CrossChainRegistryStorage contract.
var CrossChainRegistryStorageMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"addChainIDsToWhitelist\",\"inputs\":[{\"name\":\"chainIDs\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"operatorTableUpdaters\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addTransportDestinations\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"chainIDs\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allocationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateOperatorTableBytes\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createGenerationReservation\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operatorTableCalculator\",\"type\":\"address\",\"internalType\":\"contractIOperatorTableCalculator\"},{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structICrossChainRegistryTypes.OperatorSetConfig\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxStalenessPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"chainIDs\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getActiveGenerationReservations\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getActiveTransportReservations\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"\",\"type\":\"uint256[][]\",\"internalType\":\"uint256[][]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorSetConfig\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structICrossChainRegistryTypes.OperatorSetConfig\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxStalenessPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorTableCalculator\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIOperatorTableCalculator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTransportDestinations\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"keyRegistrar\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIKeyRegistrar\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeChainIDsFromWhitelist\",\"inputs\":[{\"name\":\"chainIDs\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeGenerationReservation\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeTransportDestinations\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"chainIDs\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOperatorSetConfig\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structICrossChainRegistryTypes.OperatorSetConfig\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxStalenessPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOperatorTableCalculator\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operatorTableCalculator\",\"type\":\"address\",\"internalType\":\"contractIOperatorTableCalculator\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ChainIDAddedToWhitelist\",\"inputs\":[{\"name\":\"chainID\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"operatorTableUpdater\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainIDRemovedFromWhitelist\",\"inputs\":[{\"name\":\"chainID\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GenerationReservationCreated\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GenerationReservationRemoved\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSetConfigRemoved\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSetConfigSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structICrossChainRegistryTypes.OperatorSetConfig\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxStalenessPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorTableCalculatorRemoved\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorTableCalculatorSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operatorTableCalculator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIOperatorTableCalculator\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TransportDestinationChainAdded\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"chainID\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TransportDestinationChainRemoved\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"chainID\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TransportDestinationsRemoved\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ArrayLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainIDAlreadyWhitelisted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainIDNotWhitelisted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EmptyChainIDsArray\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"GenerationReservationAlreadyExists\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"GenerationReservationDoesNotExist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidChainId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperatorSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RequireAtLeastOneTransportDestination\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TransportDestinationAlreadyAdded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TransportDestinationNotFound\",\"inputs\":[]}]",
}

// CrossChainRegistryStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use CrossChainRegistryStorageMetaData.ABI instead.
var CrossChainRegistryStorageABI = CrossChainRegistryStorageMetaData.ABI

// CrossChainRegistryStorage is an auto generated Go binding around an Ethereum contract.
type CrossChainRegistryStorage struct {
	CrossChainRegistryStorageCaller     // Read-only binding to the contract
	CrossChainRegistryStorageTransactor // Write-only binding to the contract
	CrossChainRegistryStorageFilterer   // Log filterer for contract events
}

// CrossChainRegistryStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type CrossChainRegistryStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossChainRegistryStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CrossChainRegistryStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossChainRegistryStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrossChainRegistryStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossChainRegistryStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrossChainRegistryStorageSession struct {
	Contract     *CrossChainRegistryStorage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// CrossChainRegistryStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrossChainRegistryStorageCallerSession struct {
	Contract *CrossChainRegistryStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// CrossChainRegistryStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrossChainRegistryStorageTransactorSession struct {
	Contract     *CrossChainRegistryStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// CrossChainRegistryStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type CrossChainRegistryStorageRaw struct {
	Contract *CrossChainRegistryStorage // Generic contract binding to access the raw methods on
}

// CrossChainRegistryStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrossChainRegistryStorageCallerRaw struct {
	Contract *CrossChainRegistryStorageCaller // Generic read-only contract binding to access the raw methods on
}

// CrossChainRegistryStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrossChainRegistryStorageTransactorRaw struct {
	Contract *CrossChainRegistryStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCrossChainRegistryStorage creates a new instance of CrossChainRegistryStorage, bound to a specific deployed contract.
func NewCrossChainRegistryStorage(address common.Address, backend bind.ContractBackend) (*CrossChainRegistryStorage, error) {
	contract, err := bindCrossChainRegistryStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CrossChainRegistryStorage{CrossChainRegistryStorageCaller: CrossChainRegistryStorageCaller{contract: contract}, CrossChainRegistryStorageTransactor: CrossChainRegistryStorageTransactor{contract: contract}, CrossChainRegistryStorageFilterer: CrossChainRegistryStorageFilterer{contract: contract}}, nil
}

// NewCrossChainRegistryStorageCaller creates a new read-only instance of CrossChainRegistryStorage, bound to a specific deployed contract.
func NewCrossChainRegistryStorageCaller(address common.Address, caller bind.ContractCaller) (*CrossChainRegistryStorageCaller, error) {
	contract, err := bindCrossChainRegistryStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrossChainRegistryStorageCaller{contract: contract}, nil
}

// NewCrossChainRegistryStorageTransactor creates a new write-only instance of CrossChainRegistryStorage, bound to a specific deployed contract.
func NewCrossChainRegistryStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*CrossChainRegistryStorageTransactor, error) {
	contract, err := bindCrossChainRegistryStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrossChainRegistryStorageTransactor{contract: contract}, nil
}

// NewCrossChainRegistryStorageFilterer creates a new log filterer instance of CrossChainRegistryStorage, bound to a specific deployed contract.
func NewCrossChainRegistryStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*CrossChainRegistryStorageFilterer, error) {
	contract, err := bindCrossChainRegistryStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrossChainRegistryStorageFilterer{contract: contract}, nil
}

// bindCrossChainRegistryStorage binds a generic wrapper to an already deployed contract.
func bindCrossChainRegistryStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CrossChainRegistryStorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossChainRegistryStorage *CrossChainRegistryStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossChainRegistryStorage.Contract.CrossChainRegistryStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossChainRegistryStorage *CrossChainRegistryStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossChainRegistryStorage.Contract.CrossChainRegistryStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossChainRegistryStorage *CrossChainRegistryStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossChainRegistryStorage.Contract.CrossChainRegistryStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossChainRegistryStorage *CrossChainRegistryStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossChainRegistryStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossChainRegistryStorage *CrossChainRegistryStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossChainRegistryStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossChainRegistryStorage *CrossChainRegistryStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossChainRegistryStorage.Contract.contract.Transact(opts, method, params...)
}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_CrossChainRegistryStorage *CrossChainRegistryStorageCaller) AllocationManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossChainRegistryStorage.contract.Call(opts, &out, "allocationManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_CrossChainRegistryStorage *CrossChainRegistryStorageSession) AllocationManager() (common.Address, error) {
	return _CrossChainRegistryStorage.Contract.AllocationManager(&_CrossChainRegistryStorage.CallOpts)
}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_CrossChainRegistryStorage *CrossChainRegistryStorageCallerSession) AllocationManager() (common.Address, error) {
	return _CrossChainRegistryStorage.Contract.AllocationManager(&_CrossChainRegistryStorage.CallOpts)
}

// CalculateOperatorTableBytes is a free data retrieval call binding the contract method 0x41ee6d0e.
//
// Solidity: function calculateOperatorTableBytes((address,uint32) operatorSet) view returns(bytes)
func (_CrossChainRegistryStorage *CrossChainRegistryStorageCaller) CalculateOperatorTableBytes(opts *bind.CallOpts, operatorSet OperatorSet) ([]byte, error) {
	var out []interface{}
	err := _CrossChainRegistryStorage.contract.Call(opts, &out, "calculateOperatorTableBytes", operatorSet)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// CalculateOperatorTableBytes is a free data retrieval call binding the contract method 0x41ee6d0e.
//
// Solidity: function calculateOperatorTableBytes((address,uint32) operatorSet) view returns(bytes)
func (_CrossChainRegistryStorage *CrossChainRegistryStorageSession) CalculateOperatorTableBytes(operatorSet OperatorSet) ([]byte, error) {
	return _CrossChainRegistryStorage.Contract.CalculateOperatorTableBytes(&_CrossChainRegistryStorage.CallOpts, operatorSet)
}

// CalculateOperatorTableBytes is a free data retrieval call binding the contract method 0x41ee6d0e.
//
// Solidity: function calculateOperatorTableBytes((address,uint32) operatorSet) view returns(bytes)
func (_CrossChainRegistryStorage *CrossChainRegistryStorageCallerSession) CalculateOperatorTableBytes(operatorSet OperatorSet) ([]byte, error) {
	return _CrossChainRegistryStorage.Contract.CalculateOperatorTableBytes(&_CrossChainRegistryStorage.CallOpts, operatorSet)
}

// GetActiveGenerationReservations is a free data retrieval call binding the contract method 0xd09b978b.
//
// Solidity: function getActiveGenerationReservations() view returns((address,uint32)[])
func (_CrossChainRegistryStorage *CrossChainRegistryStorageCaller) GetActiveGenerationReservations(opts *bind.CallOpts) ([]OperatorSet, error) {
	var out []interface{}
	err := _CrossChainRegistryStorage.contract.Call(opts, &out, "getActiveGenerationReservations")

	if err != nil {
		return *new([]OperatorSet), err
	}

	out0 := *abi.ConvertType(out[0], new([]OperatorSet)).(*[]OperatorSet)

	return out0, err

}

// GetActiveGenerationReservations is a free data retrieval call binding the contract method 0xd09b978b.
//
// Solidity: function getActiveGenerationReservations() view returns((address,uint32)[])
func (_CrossChainRegistryStorage *CrossChainRegistryStorageSession) GetActiveGenerationReservations() ([]OperatorSet, error) {
	return _CrossChainRegistryStorage.Contract.GetActiveGenerationReservations(&_CrossChainRegistryStorage.CallOpts)
}

// GetActiveGenerationReservations is a free data retrieval call binding the contract method 0xd09b978b.
//
// Solidity: function getActiveGenerationReservations() view returns((address,uint32)[])
func (_CrossChainRegistryStorage *CrossChainRegistryStorageCallerSession) GetActiveGenerationReservations() ([]OperatorSet, error) {
	return _CrossChainRegistryStorage.Contract.GetActiveGenerationReservations(&_CrossChainRegistryStorage.CallOpts)
}

// GetActiveTransportReservations is a free data retrieval call binding the contract method 0xbfda3b3d.
//
// Solidity: function getActiveTransportReservations() view returns((address,uint32)[], uint256[][])
func (_CrossChainRegistryStorage *CrossChainRegistryStorageCaller) GetActiveTransportReservations(opts *bind.CallOpts) ([]OperatorSet, [][]*big.Int, error) {
	var out []interface{}
	err := _CrossChainRegistryStorage.contract.Call(opts, &out, "getActiveTransportReservations")

	if err != nil {
		return *new([]OperatorSet), *new([][]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]OperatorSet)).(*[]OperatorSet)
	out1 := *abi.ConvertType(out[1], new([][]*big.Int)).(*[][]*big.Int)

	return out0, out1, err

}

// GetActiveTransportReservations is a free data retrieval call binding the contract method 0xbfda3b3d.
//
// Solidity: function getActiveTransportReservations() view returns((address,uint32)[], uint256[][])
func (_CrossChainRegistryStorage *CrossChainRegistryStorageSession) GetActiveTransportReservations() ([]OperatorSet, [][]*big.Int, error) {
	return _CrossChainRegistryStorage.Contract.GetActiveTransportReservations(&_CrossChainRegistryStorage.CallOpts)
}

// GetActiveTransportReservations is a free data retrieval call binding the contract method 0xbfda3b3d.
//
// Solidity: function getActiveTransportReservations() view returns((address,uint32)[], uint256[][])
func (_CrossChainRegistryStorage *CrossChainRegistryStorageCallerSession) GetActiveTransportReservations() ([]OperatorSet, [][]*big.Int, error) {
	return _CrossChainRegistryStorage.Contract.GetActiveTransportReservations(&_CrossChainRegistryStorage.CallOpts)
}

// GetOperatorSetConfig is a free data retrieval call binding the contract method 0x21fa7fdc.
//
// Solidity: function getOperatorSetConfig((address,uint32) operatorSet) view returns((address,uint32))
func (_CrossChainRegistryStorage *CrossChainRegistryStorageCaller) GetOperatorSetConfig(opts *bind.CallOpts, operatorSet OperatorSet) (ICrossChainRegistryTypesOperatorSetConfig, error) {
	var out []interface{}
	err := _CrossChainRegistryStorage.contract.Call(opts, &out, "getOperatorSetConfig", operatorSet)

	if err != nil {
		return *new(ICrossChainRegistryTypesOperatorSetConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(ICrossChainRegistryTypesOperatorSetConfig)).(*ICrossChainRegistryTypesOperatorSetConfig)

	return out0, err

}

// GetOperatorSetConfig is a free data retrieval call binding the contract method 0x21fa7fdc.
//
// Solidity: function getOperatorSetConfig((address,uint32) operatorSet) view returns((address,uint32))
func (_CrossChainRegistryStorage *CrossChainRegistryStorageSession) GetOperatorSetConfig(operatorSet OperatorSet) (ICrossChainRegistryTypesOperatorSetConfig, error) {
	return _CrossChainRegistryStorage.Contract.GetOperatorSetConfig(&_CrossChainRegistryStorage.CallOpts, operatorSet)
}

// GetOperatorSetConfig is a free data retrieval call binding the contract method 0x21fa7fdc.
//
// Solidity: function getOperatorSetConfig((address,uint32) operatorSet) view returns((address,uint32))
func (_CrossChainRegistryStorage *CrossChainRegistryStorageCallerSession) GetOperatorSetConfig(operatorSet OperatorSet) (ICrossChainRegistryTypesOperatorSetConfig, error) {
	return _CrossChainRegistryStorage.Contract.GetOperatorSetConfig(&_CrossChainRegistryStorage.CallOpts, operatorSet)
}

// GetOperatorTableCalculator is a free data retrieval call binding the contract method 0x75e4b539.
//
// Solidity: function getOperatorTableCalculator((address,uint32) operatorSet) view returns(address)
func (_CrossChainRegistryStorage *CrossChainRegistryStorageCaller) GetOperatorTableCalculator(opts *bind.CallOpts, operatorSet OperatorSet) (common.Address, error) {
	var out []interface{}
	err := _CrossChainRegistryStorage.contract.Call(opts, &out, "getOperatorTableCalculator", operatorSet)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOperatorTableCalculator is a free data retrieval call binding the contract method 0x75e4b539.
//
// Solidity: function getOperatorTableCalculator((address,uint32) operatorSet) view returns(address)
func (_CrossChainRegistryStorage *CrossChainRegistryStorageSession) GetOperatorTableCalculator(operatorSet OperatorSet) (common.Address, error) {
	return _CrossChainRegistryStorage.Contract.GetOperatorTableCalculator(&_CrossChainRegistryStorage.CallOpts, operatorSet)
}

// GetOperatorTableCalculator is a free data retrieval call binding the contract method 0x75e4b539.
//
// Solidity: function getOperatorTableCalculator((address,uint32) operatorSet) view returns(address)
func (_CrossChainRegistryStorage *CrossChainRegistryStorageCallerSession) GetOperatorTableCalculator(operatorSet OperatorSet) (common.Address, error) {
	return _CrossChainRegistryStorage.Contract.GetOperatorTableCalculator(&_CrossChainRegistryStorage.CallOpts, operatorSet)
}

// GetSupportedChains is a free data retrieval call binding the contract method 0xc4bffe2b.
//
// Solidity: function getSupportedChains() view returns(uint256[], address[])
func (_CrossChainRegistryStorage *CrossChainRegistryStorageCaller) GetSupportedChains(opts *bind.CallOpts) ([]*big.Int, []common.Address, error) {
	var out []interface{}
	err := _CrossChainRegistryStorage.contract.Call(opts, &out, "getSupportedChains")

	if err != nil {
		return *new([]*big.Int), *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	out1 := *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)

	return out0, out1, err

}

// GetSupportedChains is a free data retrieval call binding the contract method 0xc4bffe2b.
//
// Solidity: function getSupportedChains() view returns(uint256[], address[])
func (_CrossChainRegistryStorage *CrossChainRegistryStorageSession) GetSupportedChains() ([]*big.Int, []common.Address, error) {
	return _CrossChainRegistryStorage.Contract.GetSupportedChains(&_CrossChainRegistryStorage.CallOpts)
}

// GetSupportedChains is a free data retrieval call binding the contract method 0xc4bffe2b.
//
// Solidity: function getSupportedChains() view returns(uint256[], address[])
func (_CrossChainRegistryStorage *CrossChainRegistryStorageCallerSession) GetSupportedChains() ([]*big.Int, []common.Address, error) {
	return _CrossChainRegistryStorage.Contract.GetSupportedChains(&_CrossChainRegistryStorage.CallOpts)
}

// GetTransportDestinations is a free data retrieval call binding the contract method 0x3c75fddf.
//
// Solidity: function getTransportDestinations((address,uint32) operatorSet) view returns(uint256[])
func (_CrossChainRegistryStorage *CrossChainRegistryStorageCaller) GetTransportDestinations(opts *bind.CallOpts, operatorSet OperatorSet) ([]*big.Int, error) {
	var out []interface{}
	err := _CrossChainRegistryStorage.contract.Call(opts, &out, "getTransportDestinations", operatorSet)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetTransportDestinations is a free data retrieval call binding the contract method 0x3c75fddf.
//
// Solidity: function getTransportDestinations((address,uint32) operatorSet) view returns(uint256[])
func (_CrossChainRegistryStorage *CrossChainRegistryStorageSession) GetTransportDestinations(operatorSet OperatorSet) ([]*big.Int, error) {
	return _CrossChainRegistryStorage.Contract.GetTransportDestinations(&_CrossChainRegistryStorage.CallOpts, operatorSet)
}

// GetTransportDestinations is a free data retrieval call binding the contract method 0x3c75fddf.
//
// Solidity: function getTransportDestinations((address,uint32) operatorSet) view returns(uint256[])
func (_CrossChainRegistryStorage *CrossChainRegistryStorageCallerSession) GetTransportDestinations(operatorSet OperatorSet) ([]*big.Int, error) {
	return _CrossChainRegistryStorage.Contract.GetTransportDestinations(&_CrossChainRegistryStorage.CallOpts, operatorSet)
}

// KeyRegistrar is a free data retrieval call binding the contract method 0x3ec45c7e.
//
// Solidity: function keyRegistrar() view returns(address)
func (_CrossChainRegistryStorage *CrossChainRegistryStorageCaller) KeyRegistrar(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossChainRegistryStorage.contract.Call(opts, &out, "keyRegistrar")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// KeyRegistrar is a free data retrieval call binding the contract method 0x3ec45c7e.
//
// Solidity: function keyRegistrar() view returns(address)
func (_CrossChainRegistryStorage *CrossChainRegistryStorageSession) KeyRegistrar() (common.Address, error) {
	return _CrossChainRegistryStorage.Contract.KeyRegistrar(&_CrossChainRegistryStorage.CallOpts)
}

// KeyRegistrar is a free data retrieval call binding the contract method 0x3ec45c7e.
//
// Solidity: function keyRegistrar() view returns(address)
func (_CrossChainRegistryStorage *CrossChainRegistryStorageCallerSession) KeyRegistrar() (common.Address, error) {
	return _CrossChainRegistryStorage.Contract.KeyRegistrar(&_CrossChainRegistryStorage.CallOpts)
}

// AddChainIDsToWhitelist is a paid mutator transaction binding the contract method 0x04e98be3.
//
// Solidity: function addChainIDsToWhitelist(uint256[] chainIDs, address[] operatorTableUpdaters) returns()
func (_CrossChainRegistryStorage *CrossChainRegistryStorageTransactor) AddChainIDsToWhitelist(opts *bind.TransactOpts, chainIDs []*big.Int, operatorTableUpdaters []common.Address) (*types.Transaction, error) {
	return _CrossChainRegistryStorage.contract.Transact(opts, "addChainIDsToWhitelist", chainIDs, operatorTableUpdaters)
}

// AddChainIDsToWhitelist is a paid mutator transaction binding the contract method 0x04e98be3.
//
// Solidity: function addChainIDsToWhitelist(uint256[] chainIDs, address[] operatorTableUpdaters) returns()
func (_CrossChainRegistryStorage *CrossChainRegistryStorageSession) AddChainIDsToWhitelist(chainIDs []*big.Int, operatorTableUpdaters []common.Address) (*types.Transaction, error) {
	return _CrossChainRegistryStorage.Contract.AddChainIDsToWhitelist(&_CrossChainRegistryStorage.TransactOpts, chainIDs, operatorTableUpdaters)
}

// AddChainIDsToWhitelist is a paid mutator transaction binding the contract method 0x04e98be3.
//
// Solidity: function addChainIDsToWhitelist(uint256[] chainIDs, address[] operatorTableUpdaters) returns()
func (_CrossChainRegistryStorage *CrossChainRegistryStorageTransactorSession) AddChainIDsToWhitelist(chainIDs []*big.Int, operatorTableUpdaters []common.Address) (*types.Transaction, error) {
	return _CrossChainRegistryStorage.Contract.AddChainIDsToWhitelist(&_CrossChainRegistryStorage.TransactOpts, chainIDs, operatorTableUpdaters)
}

// AddTransportDestinations is a paid mutator transaction binding the contract method 0x49be7d6f.
//
// Solidity: function addTransportDestinations((address,uint32) operatorSet, uint256[] chainIDs) returns()
func (_CrossChainRegistryStorage *CrossChainRegistryStorageTransactor) AddTransportDestinations(opts *bind.TransactOpts, operatorSet OperatorSet, chainIDs []*big.Int) (*types.Transaction, error) {
	return _CrossChainRegistryStorage.contract.Transact(opts, "addTransportDestinations", operatorSet, chainIDs)
}

// AddTransportDestinations is a paid mutator transaction binding the contract method 0x49be7d6f.
//
// Solidity: function addTransportDestinations((address,uint32) operatorSet, uint256[] chainIDs) returns()
func (_CrossChainRegistryStorage *CrossChainRegistryStorageSession) AddTransportDestinations(operatorSet OperatorSet, chainIDs []*big.Int) (*types.Transaction, error) {
	return _CrossChainRegistryStorage.Contract.AddTransportDestinations(&_CrossChainRegistryStorage.TransactOpts, operatorSet, chainIDs)
}

// AddTransportDestinations is a paid mutator transaction binding the contract method 0x49be7d6f.
//
// Solidity: function addTransportDestinations((address,uint32) operatorSet, uint256[] chainIDs) returns()
func (_CrossChainRegistryStorage *CrossChainRegistryStorageTransactorSession) AddTransportDestinations(operatorSet OperatorSet, chainIDs []*big.Int) (*types.Transaction, error) {
	return _CrossChainRegistryStorage.Contract.AddTransportDestinations(&_CrossChainRegistryStorage.TransactOpts, operatorSet, chainIDs)
}

// CreateGenerationReservation is a paid mutator transaction binding the contract method 0xfe596dee.
//
// Solidity: function createGenerationReservation((address,uint32) operatorSet, address operatorTableCalculator, (address,uint32) config, uint256[] chainIDs) returns()
func (_CrossChainRegistryStorage *CrossChainRegistryStorageTransactor) CreateGenerationReservation(opts *bind.TransactOpts, operatorSet OperatorSet, operatorTableCalculator common.Address, config ICrossChainRegistryTypesOperatorSetConfig, chainIDs []*big.Int) (*types.Transaction, error) {
	return _CrossChainRegistryStorage.contract.Transact(opts, "createGenerationReservation", operatorSet, operatorTableCalculator, config, chainIDs)
}

// CreateGenerationReservation is a paid mutator transaction binding the contract method 0xfe596dee.
//
// Solidity: function createGenerationReservation((address,uint32) operatorSet, address operatorTableCalculator, (address,uint32) config, uint256[] chainIDs) returns()
func (_CrossChainRegistryStorage *CrossChainRegistryStorageSession) CreateGenerationReservation(operatorSet OperatorSet, operatorTableCalculator common.Address, config ICrossChainRegistryTypesOperatorSetConfig, chainIDs []*big.Int) (*types.Transaction, error) {
	return _CrossChainRegistryStorage.Contract.CreateGenerationReservation(&_CrossChainRegistryStorage.TransactOpts, operatorSet, operatorTableCalculator, config, chainIDs)
}

// CreateGenerationReservation is a paid mutator transaction binding the contract method 0xfe596dee.
//
// Solidity: function createGenerationReservation((address,uint32) operatorSet, address operatorTableCalculator, (address,uint32) config, uint256[] chainIDs) returns()
func (_CrossChainRegistryStorage *CrossChainRegistryStorageTransactorSession) CreateGenerationReservation(operatorSet OperatorSet, operatorTableCalculator common.Address, config ICrossChainRegistryTypesOperatorSetConfig, chainIDs []*big.Int) (*types.Transaction, error) {
	return _CrossChainRegistryStorage.Contract.CreateGenerationReservation(&_CrossChainRegistryStorage.TransactOpts, operatorSet, operatorTableCalculator, config, chainIDs)
}

// RemoveChainIDsFromWhitelist is a paid mutator transaction binding the contract method 0xdfbd9dfd.
//
// Solidity: function removeChainIDsFromWhitelist(uint256[] chainIDs) returns()
func (_CrossChainRegistryStorage *CrossChainRegistryStorageTransactor) RemoveChainIDsFromWhitelist(opts *bind.TransactOpts, chainIDs []*big.Int) (*types.Transaction, error) {
	return _CrossChainRegistryStorage.contract.Transact(opts, "removeChainIDsFromWhitelist", chainIDs)
}

// RemoveChainIDsFromWhitelist is a paid mutator transaction binding the contract method 0xdfbd9dfd.
//
// Solidity: function removeChainIDsFromWhitelist(uint256[] chainIDs) returns()
func (_CrossChainRegistryStorage *CrossChainRegistryStorageSession) RemoveChainIDsFromWhitelist(chainIDs []*big.Int) (*types.Transaction, error) {
	return _CrossChainRegistryStorage.Contract.RemoveChainIDsFromWhitelist(&_CrossChainRegistryStorage.TransactOpts, chainIDs)
}

// RemoveChainIDsFromWhitelist is a paid mutator transaction binding the contract method 0xdfbd9dfd.
//
// Solidity: function removeChainIDsFromWhitelist(uint256[] chainIDs) returns()
func (_CrossChainRegistryStorage *CrossChainRegistryStorageTransactorSession) RemoveChainIDsFromWhitelist(chainIDs []*big.Int) (*types.Transaction, error) {
	return _CrossChainRegistryStorage.Contract.RemoveChainIDsFromWhitelist(&_CrossChainRegistryStorage.TransactOpts, chainIDs)
}

// RemoveGenerationReservation is a paid mutator transaction binding the contract method 0x6c55a37f.
//
// Solidity: function removeGenerationReservation((address,uint32) operatorSet) returns()
func (_CrossChainRegistryStorage *CrossChainRegistryStorageTransactor) RemoveGenerationReservation(opts *bind.TransactOpts, operatorSet OperatorSet) (*types.Transaction, error) {
	return _CrossChainRegistryStorage.contract.Transact(opts, "removeGenerationReservation", operatorSet)
}

// RemoveGenerationReservation is a paid mutator transaction binding the contract method 0x6c55a37f.
//
// Solidity: function removeGenerationReservation((address,uint32) operatorSet) returns()
func (_CrossChainRegistryStorage *CrossChainRegistryStorageSession) RemoveGenerationReservation(operatorSet OperatorSet) (*types.Transaction, error) {
	return _CrossChainRegistryStorage.Contract.RemoveGenerationReservation(&_CrossChainRegistryStorage.TransactOpts, operatorSet)
}

// RemoveGenerationReservation is a paid mutator transaction binding the contract method 0x6c55a37f.
//
// Solidity: function removeGenerationReservation((address,uint32) operatorSet) returns()
func (_CrossChainRegistryStorage *CrossChainRegistryStorageTransactorSession) RemoveGenerationReservation(operatorSet OperatorSet) (*types.Transaction, error) {
	return _CrossChainRegistryStorage.Contract.RemoveGenerationReservation(&_CrossChainRegistryStorage.TransactOpts, operatorSet)
}

// RemoveTransportDestinations is a paid mutator transaction binding the contract method 0xf3e9f5d4.
//
// Solidity: function removeTransportDestinations((address,uint32) operatorSet, uint256[] chainIDs) returns()
func (_CrossChainRegistryStorage *CrossChainRegistryStorageTransactor) RemoveTransportDestinations(opts *bind.TransactOpts, operatorSet OperatorSet, chainIDs []*big.Int) (*types.Transaction, error) {
	return _CrossChainRegistryStorage.contract.Transact(opts, "removeTransportDestinations", operatorSet, chainIDs)
}

// RemoveTransportDestinations is a paid mutator transaction binding the contract method 0xf3e9f5d4.
//
// Solidity: function removeTransportDestinations((address,uint32) operatorSet, uint256[] chainIDs) returns()
func (_CrossChainRegistryStorage *CrossChainRegistryStorageSession) RemoveTransportDestinations(operatorSet OperatorSet, chainIDs []*big.Int) (*types.Transaction, error) {
	return _CrossChainRegistryStorage.Contract.RemoveTransportDestinations(&_CrossChainRegistryStorage.TransactOpts, operatorSet, chainIDs)
}

// RemoveTransportDestinations is a paid mutator transaction binding the contract method 0xf3e9f5d4.
//
// Solidity: function removeTransportDestinations((address,uint32) operatorSet, uint256[] chainIDs) returns()
func (_CrossChainRegistryStorage *CrossChainRegistryStorageTransactorSession) RemoveTransportDestinations(operatorSet OperatorSet, chainIDs []*big.Int) (*types.Transaction, error) {
	return _CrossChainRegistryStorage.Contract.RemoveTransportDestinations(&_CrossChainRegistryStorage.TransactOpts, operatorSet, chainIDs)
}

// SetOperatorSetConfig is a paid mutator transaction binding the contract method 0x277e1e62.
//
// Solidity: function setOperatorSetConfig((address,uint32) operatorSet, (address,uint32) config) returns()
func (_CrossChainRegistryStorage *CrossChainRegistryStorageTransactor) SetOperatorSetConfig(opts *bind.TransactOpts, operatorSet OperatorSet, config ICrossChainRegistryTypesOperatorSetConfig) (*types.Transaction, error) {
	return _CrossChainRegistryStorage.contract.Transact(opts, "setOperatorSetConfig", operatorSet, config)
}

// SetOperatorSetConfig is a paid mutator transaction binding the contract method 0x277e1e62.
//
// Solidity: function setOperatorSetConfig((address,uint32) operatorSet, (address,uint32) config) returns()
func (_CrossChainRegistryStorage *CrossChainRegistryStorageSession) SetOperatorSetConfig(operatorSet OperatorSet, config ICrossChainRegistryTypesOperatorSetConfig) (*types.Transaction, error) {
	return _CrossChainRegistryStorage.Contract.SetOperatorSetConfig(&_CrossChainRegistryStorage.TransactOpts, operatorSet, config)
}

// SetOperatorSetConfig is a paid mutator transaction binding the contract method 0x277e1e62.
//
// Solidity: function setOperatorSetConfig((address,uint32) operatorSet, (address,uint32) config) returns()
func (_CrossChainRegistryStorage *CrossChainRegistryStorageTransactorSession) SetOperatorSetConfig(operatorSet OperatorSet, config ICrossChainRegistryTypesOperatorSetConfig) (*types.Transaction, error) {
	return _CrossChainRegistryStorage.Contract.SetOperatorSetConfig(&_CrossChainRegistryStorage.TransactOpts, operatorSet, config)
}

// SetOperatorTableCalculator is a paid mutator transaction binding the contract method 0x1ca9142a.
//
// Solidity: function setOperatorTableCalculator((address,uint32) operatorSet, address operatorTableCalculator) returns()
func (_CrossChainRegistryStorage *CrossChainRegistryStorageTransactor) SetOperatorTableCalculator(opts *bind.TransactOpts, operatorSet OperatorSet, operatorTableCalculator common.Address) (*types.Transaction, error) {
	return _CrossChainRegistryStorage.contract.Transact(opts, "setOperatorTableCalculator", operatorSet, operatorTableCalculator)
}

// SetOperatorTableCalculator is a paid mutator transaction binding the contract method 0x1ca9142a.
//
// Solidity: function setOperatorTableCalculator((address,uint32) operatorSet, address operatorTableCalculator) returns()
func (_CrossChainRegistryStorage *CrossChainRegistryStorageSession) SetOperatorTableCalculator(operatorSet OperatorSet, operatorTableCalculator common.Address) (*types.Transaction, error) {
	return _CrossChainRegistryStorage.Contract.SetOperatorTableCalculator(&_CrossChainRegistryStorage.TransactOpts, operatorSet, operatorTableCalculator)
}

// SetOperatorTableCalculator is a paid mutator transaction binding the contract method 0x1ca9142a.
//
// Solidity: function setOperatorTableCalculator((address,uint32) operatorSet, address operatorTableCalculator) returns()
func (_CrossChainRegistryStorage *CrossChainRegistryStorageTransactorSession) SetOperatorTableCalculator(operatorSet OperatorSet, operatorTableCalculator common.Address) (*types.Transaction, error) {
	return _CrossChainRegistryStorage.Contract.SetOperatorTableCalculator(&_CrossChainRegistryStorage.TransactOpts, operatorSet, operatorTableCalculator)
}

// CrossChainRegistryStorageChainIDAddedToWhitelistIterator is returned from FilterChainIDAddedToWhitelist and is used to iterate over the raw logs and unpacked data for ChainIDAddedToWhitelist events raised by the CrossChainRegistryStorage contract.
type CrossChainRegistryStorageChainIDAddedToWhitelistIterator struct {
	Event *CrossChainRegistryStorageChainIDAddedToWhitelist // Event containing the contract specifics and raw log

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
func (it *CrossChainRegistryStorageChainIDAddedToWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainRegistryStorageChainIDAddedToWhitelist)
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
		it.Event = new(CrossChainRegistryStorageChainIDAddedToWhitelist)
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
func (it *CrossChainRegistryStorageChainIDAddedToWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossChainRegistryStorageChainIDAddedToWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossChainRegistryStorageChainIDAddedToWhitelist represents a ChainIDAddedToWhitelist event raised by the CrossChainRegistryStorage contract.
type CrossChainRegistryStorageChainIDAddedToWhitelist struct {
	ChainID              *big.Int
	OperatorTableUpdater common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterChainIDAddedToWhitelist is a free log retrieval operation binding the contract event 0x7a0a76d85b582b17996dd7371a407aa7a79b870db8539247fba315c7b6beff62.
//
// Solidity: event ChainIDAddedToWhitelist(uint256 chainID, address operatorTableUpdater)
func (_CrossChainRegistryStorage *CrossChainRegistryStorageFilterer) FilterChainIDAddedToWhitelist(opts *bind.FilterOpts) (*CrossChainRegistryStorageChainIDAddedToWhitelistIterator, error) {

	logs, sub, err := _CrossChainRegistryStorage.contract.FilterLogs(opts, "ChainIDAddedToWhitelist")
	if err != nil {
		return nil, err
	}
	return &CrossChainRegistryStorageChainIDAddedToWhitelistIterator{contract: _CrossChainRegistryStorage.contract, event: "ChainIDAddedToWhitelist", logs: logs, sub: sub}, nil
}

// WatchChainIDAddedToWhitelist is a free log subscription operation binding the contract event 0x7a0a76d85b582b17996dd7371a407aa7a79b870db8539247fba315c7b6beff62.
//
// Solidity: event ChainIDAddedToWhitelist(uint256 chainID, address operatorTableUpdater)
func (_CrossChainRegistryStorage *CrossChainRegistryStorageFilterer) WatchChainIDAddedToWhitelist(opts *bind.WatchOpts, sink chan<- *CrossChainRegistryStorageChainIDAddedToWhitelist) (event.Subscription, error) {

	logs, sub, err := _CrossChainRegistryStorage.contract.WatchLogs(opts, "ChainIDAddedToWhitelist")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossChainRegistryStorageChainIDAddedToWhitelist)
				if err := _CrossChainRegistryStorage.contract.UnpackLog(event, "ChainIDAddedToWhitelist", log); err != nil {
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

// ParseChainIDAddedToWhitelist is a log parse operation binding the contract event 0x7a0a76d85b582b17996dd7371a407aa7a79b870db8539247fba315c7b6beff62.
//
// Solidity: event ChainIDAddedToWhitelist(uint256 chainID, address operatorTableUpdater)
func (_CrossChainRegistryStorage *CrossChainRegistryStorageFilterer) ParseChainIDAddedToWhitelist(log types.Log) (*CrossChainRegistryStorageChainIDAddedToWhitelist, error) {
	event := new(CrossChainRegistryStorageChainIDAddedToWhitelist)
	if err := _CrossChainRegistryStorage.contract.UnpackLog(event, "ChainIDAddedToWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossChainRegistryStorageChainIDRemovedFromWhitelistIterator is returned from FilterChainIDRemovedFromWhitelist and is used to iterate over the raw logs and unpacked data for ChainIDRemovedFromWhitelist events raised by the CrossChainRegistryStorage contract.
type CrossChainRegistryStorageChainIDRemovedFromWhitelistIterator struct {
	Event *CrossChainRegistryStorageChainIDRemovedFromWhitelist // Event containing the contract specifics and raw log

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
func (it *CrossChainRegistryStorageChainIDRemovedFromWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainRegistryStorageChainIDRemovedFromWhitelist)
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
		it.Event = new(CrossChainRegistryStorageChainIDRemovedFromWhitelist)
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
func (it *CrossChainRegistryStorageChainIDRemovedFromWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossChainRegistryStorageChainIDRemovedFromWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossChainRegistryStorageChainIDRemovedFromWhitelist represents a ChainIDRemovedFromWhitelist event raised by the CrossChainRegistryStorage contract.
type CrossChainRegistryStorageChainIDRemovedFromWhitelist struct {
	ChainID *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterChainIDRemovedFromWhitelist is a free log retrieval operation binding the contract event 0x6824d36084ecf2cd819b137cb5d837cc6e73afce1e0e348c9fdecaa81d0341e5.
//
// Solidity: event ChainIDRemovedFromWhitelist(uint256 chainID)
func (_CrossChainRegistryStorage *CrossChainRegistryStorageFilterer) FilterChainIDRemovedFromWhitelist(opts *bind.FilterOpts) (*CrossChainRegistryStorageChainIDRemovedFromWhitelistIterator, error) {

	logs, sub, err := _CrossChainRegistryStorage.contract.FilterLogs(opts, "ChainIDRemovedFromWhitelist")
	if err != nil {
		return nil, err
	}
	return &CrossChainRegistryStorageChainIDRemovedFromWhitelistIterator{contract: _CrossChainRegistryStorage.contract, event: "ChainIDRemovedFromWhitelist", logs: logs, sub: sub}, nil
}

// WatchChainIDRemovedFromWhitelist is a free log subscription operation binding the contract event 0x6824d36084ecf2cd819b137cb5d837cc6e73afce1e0e348c9fdecaa81d0341e5.
//
// Solidity: event ChainIDRemovedFromWhitelist(uint256 chainID)
func (_CrossChainRegistryStorage *CrossChainRegistryStorageFilterer) WatchChainIDRemovedFromWhitelist(opts *bind.WatchOpts, sink chan<- *CrossChainRegistryStorageChainIDRemovedFromWhitelist) (event.Subscription, error) {

	logs, sub, err := _CrossChainRegistryStorage.contract.WatchLogs(opts, "ChainIDRemovedFromWhitelist")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossChainRegistryStorageChainIDRemovedFromWhitelist)
				if err := _CrossChainRegistryStorage.contract.UnpackLog(event, "ChainIDRemovedFromWhitelist", log); err != nil {
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

// ParseChainIDRemovedFromWhitelist is a log parse operation binding the contract event 0x6824d36084ecf2cd819b137cb5d837cc6e73afce1e0e348c9fdecaa81d0341e5.
//
// Solidity: event ChainIDRemovedFromWhitelist(uint256 chainID)
func (_CrossChainRegistryStorage *CrossChainRegistryStorageFilterer) ParseChainIDRemovedFromWhitelist(log types.Log) (*CrossChainRegistryStorageChainIDRemovedFromWhitelist, error) {
	event := new(CrossChainRegistryStorageChainIDRemovedFromWhitelist)
	if err := _CrossChainRegistryStorage.contract.UnpackLog(event, "ChainIDRemovedFromWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossChainRegistryStorageGenerationReservationCreatedIterator is returned from FilterGenerationReservationCreated and is used to iterate over the raw logs and unpacked data for GenerationReservationCreated events raised by the CrossChainRegistryStorage contract.
type CrossChainRegistryStorageGenerationReservationCreatedIterator struct {
	Event *CrossChainRegistryStorageGenerationReservationCreated // Event containing the contract specifics and raw log

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
func (it *CrossChainRegistryStorageGenerationReservationCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainRegistryStorageGenerationReservationCreated)
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
		it.Event = new(CrossChainRegistryStorageGenerationReservationCreated)
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
func (it *CrossChainRegistryStorageGenerationReservationCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossChainRegistryStorageGenerationReservationCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossChainRegistryStorageGenerationReservationCreated represents a GenerationReservationCreated event raised by the CrossChainRegistryStorage contract.
type CrossChainRegistryStorageGenerationReservationCreated struct {
	OperatorSet OperatorSet
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterGenerationReservationCreated is a free log retrieval operation binding the contract event 0x4fb6efec7dd60036ce3a7af8d5c48425019daa0fb61eb471a966a7ac2c6fa6a6.
//
// Solidity: event GenerationReservationCreated((address,uint32) operatorSet)
func (_CrossChainRegistryStorage *CrossChainRegistryStorageFilterer) FilterGenerationReservationCreated(opts *bind.FilterOpts) (*CrossChainRegistryStorageGenerationReservationCreatedIterator, error) {

	logs, sub, err := _CrossChainRegistryStorage.contract.FilterLogs(opts, "GenerationReservationCreated")
	if err != nil {
		return nil, err
	}
	return &CrossChainRegistryStorageGenerationReservationCreatedIterator{contract: _CrossChainRegistryStorage.contract, event: "GenerationReservationCreated", logs: logs, sub: sub}, nil
}

// WatchGenerationReservationCreated is a free log subscription operation binding the contract event 0x4fb6efec7dd60036ce3a7af8d5c48425019daa0fb61eb471a966a7ac2c6fa6a6.
//
// Solidity: event GenerationReservationCreated((address,uint32) operatorSet)
func (_CrossChainRegistryStorage *CrossChainRegistryStorageFilterer) WatchGenerationReservationCreated(opts *bind.WatchOpts, sink chan<- *CrossChainRegistryStorageGenerationReservationCreated) (event.Subscription, error) {

	logs, sub, err := _CrossChainRegistryStorage.contract.WatchLogs(opts, "GenerationReservationCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossChainRegistryStorageGenerationReservationCreated)
				if err := _CrossChainRegistryStorage.contract.UnpackLog(event, "GenerationReservationCreated", log); err != nil {
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

// ParseGenerationReservationCreated is a log parse operation binding the contract event 0x4fb6efec7dd60036ce3a7af8d5c48425019daa0fb61eb471a966a7ac2c6fa6a6.
//
// Solidity: event GenerationReservationCreated((address,uint32) operatorSet)
func (_CrossChainRegistryStorage *CrossChainRegistryStorageFilterer) ParseGenerationReservationCreated(log types.Log) (*CrossChainRegistryStorageGenerationReservationCreated, error) {
	event := new(CrossChainRegistryStorageGenerationReservationCreated)
	if err := _CrossChainRegistryStorage.contract.UnpackLog(event, "GenerationReservationCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossChainRegistryStorageGenerationReservationRemovedIterator is returned from FilterGenerationReservationRemoved and is used to iterate over the raw logs and unpacked data for GenerationReservationRemoved events raised by the CrossChainRegistryStorage contract.
type CrossChainRegistryStorageGenerationReservationRemovedIterator struct {
	Event *CrossChainRegistryStorageGenerationReservationRemoved // Event containing the contract specifics and raw log

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
func (it *CrossChainRegistryStorageGenerationReservationRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainRegistryStorageGenerationReservationRemoved)
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
		it.Event = new(CrossChainRegistryStorageGenerationReservationRemoved)
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
func (it *CrossChainRegistryStorageGenerationReservationRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossChainRegistryStorageGenerationReservationRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossChainRegistryStorageGenerationReservationRemoved represents a GenerationReservationRemoved event raised by the CrossChainRegistryStorage contract.
type CrossChainRegistryStorageGenerationReservationRemoved struct {
	OperatorSet OperatorSet
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterGenerationReservationRemoved is a free log retrieval operation binding the contract event 0x4ffdfdd59e9e1e3c301608788f78dd458e61cb8c045ca92b62a7b484c80824fb.
//
// Solidity: event GenerationReservationRemoved((address,uint32) operatorSet)
func (_CrossChainRegistryStorage *CrossChainRegistryStorageFilterer) FilterGenerationReservationRemoved(opts *bind.FilterOpts) (*CrossChainRegistryStorageGenerationReservationRemovedIterator, error) {

	logs, sub, err := _CrossChainRegistryStorage.contract.FilterLogs(opts, "GenerationReservationRemoved")
	if err != nil {
		return nil, err
	}
	return &CrossChainRegistryStorageGenerationReservationRemovedIterator{contract: _CrossChainRegistryStorage.contract, event: "GenerationReservationRemoved", logs: logs, sub: sub}, nil
}

// WatchGenerationReservationRemoved is a free log subscription operation binding the contract event 0x4ffdfdd59e9e1e3c301608788f78dd458e61cb8c045ca92b62a7b484c80824fb.
//
// Solidity: event GenerationReservationRemoved((address,uint32) operatorSet)
func (_CrossChainRegistryStorage *CrossChainRegistryStorageFilterer) WatchGenerationReservationRemoved(opts *bind.WatchOpts, sink chan<- *CrossChainRegistryStorageGenerationReservationRemoved) (event.Subscription, error) {

	logs, sub, err := _CrossChainRegistryStorage.contract.WatchLogs(opts, "GenerationReservationRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossChainRegistryStorageGenerationReservationRemoved)
				if err := _CrossChainRegistryStorage.contract.UnpackLog(event, "GenerationReservationRemoved", log); err != nil {
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

// ParseGenerationReservationRemoved is a log parse operation binding the contract event 0x4ffdfdd59e9e1e3c301608788f78dd458e61cb8c045ca92b62a7b484c80824fb.
//
// Solidity: event GenerationReservationRemoved((address,uint32) operatorSet)
func (_CrossChainRegistryStorage *CrossChainRegistryStorageFilterer) ParseGenerationReservationRemoved(log types.Log) (*CrossChainRegistryStorageGenerationReservationRemoved, error) {
	event := new(CrossChainRegistryStorageGenerationReservationRemoved)
	if err := _CrossChainRegistryStorage.contract.UnpackLog(event, "GenerationReservationRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossChainRegistryStorageOperatorSetConfigRemovedIterator is returned from FilterOperatorSetConfigRemoved and is used to iterate over the raw logs and unpacked data for OperatorSetConfigRemoved events raised by the CrossChainRegistryStorage contract.
type CrossChainRegistryStorageOperatorSetConfigRemovedIterator struct {
	Event *CrossChainRegistryStorageOperatorSetConfigRemoved // Event containing the contract specifics and raw log

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
func (it *CrossChainRegistryStorageOperatorSetConfigRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainRegistryStorageOperatorSetConfigRemoved)
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
		it.Event = new(CrossChainRegistryStorageOperatorSetConfigRemoved)
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
func (it *CrossChainRegistryStorageOperatorSetConfigRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossChainRegistryStorageOperatorSetConfigRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossChainRegistryStorageOperatorSetConfigRemoved represents a OperatorSetConfigRemoved event raised by the CrossChainRegistryStorage contract.
type CrossChainRegistryStorageOperatorSetConfigRemoved struct {
	OperatorSet OperatorSet
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorSetConfigRemoved is a free log retrieval operation binding the contract event 0x210a1118a869246162804e2a7f21ef808ebd93f4be7ed512014fe29a7a8be02e.
//
// Solidity: event OperatorSetConfigRemoved((address,uint32) operatorSet)
func (_CrossChainRegistryStorage *CrossChainRegistryStorageFilterer) FilterOperatorSetConfigRemoved(opts *bind.FilterOpts) (*CrossChainRegistryStorageOperatorSetConfigRemovedIterator, error) {

	logs, sub, err := _CrossChainRegistryStorage.contract.FilterLogs(opts, "OperatorSetConfigRemoved")
	if err != nil {
		return nil, err
	}
	return &CrossChainRegistryStorageOperatorSetConfigRemovedIterator{contract: _CrossChainRegistryStorage.contract, event: "OperatorSetConfigRemoved", logs: logs, sub: sub}, nil
}

// WatchOperatorSetConfigRemoved is a free log subscription operation binding the contract event 0x210a1118a869246162804e2a7f21ef808ebd93f4be7ed512014fe29a7a8be02e.
//
// Solidity: event OperatorSetConfigRemoved((address,uint32) operatorSet)
func (_CrossChainRegistryStorage *CrossChainRegistryStorageFilterer) WatchOperatorSetConfigRemoved(opts *bind.WatchOpts, sink chan<- *CrossChainRegistryStorageOperatorSetConfigRemoved) (event.Subscription, error) {

	logs, sub, err := _CrossChainRegistryStorage.contract.WatchLogs(opts, "OperatorSetConfigRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossChainRegistryStorageOperatorSetConfigRemoved)
				if err := _CrossChainRegistryStorage.contract.UnpackLog(event, "OperatorSetConfigRemoved", log); err != nil {
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

// ParseOperatorSetConfigRemoved is a log parse operation binding the contract event 0x210a1118a869246162804e2a7f21ef808ebd93f4be7ed512014fe29a7a8be02e.
//
// Solidity: event OperatorSetConfigRemoved((address,uint32) operatorSet)
func (_CrossChainRegistryStorage *CrossChainRegistryStorageFilterer) ParseOperatorSetConfigRemoved(log types.Log) (*CrossChainRegistryStorageOperatorSetConfigRemoved, error) {
	event := new(CrossChainRegistryStorageOperatorSetConfigRemoved)
	if err := _CrossChainRegistryStorage.contract.UnpackLog(event, "OperatorSetConfigRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossChainRegistryStorageOperatorSetConfigSetIterator is returned from FilterOperatorSetConfigSet and is used to iterate over the raw logs and unpacked data for OperatorSetConfigSet events raised by the CrossChainRegistryStorage contract.
type CrossChainRegistryStorageOperatorSetConfigSetIterator struct {
	Event *CrossChainRegistryStorageOperatorSetConfigSet // Event containing the contract specifics and raw log

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
func (it *CrossChainRegistryStorageOperatorSetConfigSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainRegistryStorageOperatorSetConfigSet)
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
		it.Event = new(CrossChainRegistryStorageOperatorSetConfigSet)
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
func (it *CrossChainRegistryStorageOperatorSetConfigSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossChainRegistryStorageOperatorSetConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossChainRegistryStorageOperatorSetConfigSet represents a OperatorSetConfigSet event raised by the CrossChainRegistryStorage contract.
type CrossChainRegistryStorageOperatorSetConfigSet struct {
	OperatorSet OperatorSet
	Config      ICrossChainRegistryTypesOperatorSetConfig
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorSetConfigSet is a free log retrieval operation binding the contract event 0x3147846ee526009000671c20380b856a633345691300f82585f90034715cf0e2.
//
// Solidity: event OperatorSetConfigSet((address,uint32) operatorSet, (address,uint32) config)
func (_CrossChainRegistryStorage *CrossChainRegistryStorageFilterer) FilterOperatorSetConfigSet(opts *bind.FilterOpts) (*CrossChainRegistryStorageOperatorSetConfigSetIterator, error) {

	logs, sub, err := _CrossChainRegistryStorage.contract.FilterLogs(opts, "OperatorSetConfigSet")
	if err != nil {
		return nil, err
	}
	return &CrossChainRegistryStorageOperatorSetConfigSetIterator{contract: _CrossChainRegistryStorage.contract, event: "OperatorSetConfigSet", logs: logs, sub: sub}, nil
}

// WatchOperatorSetConfigSet is a free log subscription operation binding the contract event 0x3147846ee526009000671c20380b856a633345691300f82585f90034715cf0e2.
//
// Solidity: event OperatorSetConfigSet((address,uint32) operatorSet, (address,uint32) config)
func (_CrossChainRegistryStorage *CrossChainRegistryStorageFilterer) WatchOperatorSetConfigSet(opts *bind.WatchOpts, sink chan<- *CrossChainRegistryStorageOperatorSetConfigSet) (event.Subscription, error) {

	logs, sub, err := _CrossChainRegistryStorage.contract.WatchLogs(opts, "OperatorSetConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossChainRegistryStorageOperatorSetConfigSet)
				if err := _CrossChainRegistryStorage.contract.UnpackLog(event, "OperatorSetConfigSet", log); err != nil {
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

// ParseOperatorSetConfigSet is a log parse operation binding the contract event 0x3147846ee526009000671c20380b856a633345691300f82585f90034715cf0e2.
//
// Solidity: event OperatorSetConfigSet((address,uint32) operatorSet, (address,uint32) config)
func (_CrossChainRegistryStorage *CrossChainRegistryStorageFilterer) ParseOperatorSetConfigSet(log types.Log) (*CrossChainRegistryStorageOperatorSetConfigSet, error) {
	event := new(CrossChainRegistryStorageOperatorSetConfigSet)
	if err := _CrossChainRegistryStorage.contract.UnpackLog(event, "OperatorSetConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossChainRegistryStorageOperatorTableCalculatorRemovedIterator is returned from FilterOperatorTableCalculatorRemoved and is used to iterate over the raw logs and unpacked data for OperatorTableCalculatorRemoved events raised by the CrossChainRegistryStorage contract.
type CrossChainRegistryStorageOperatorTableCalculatorRemovedIterator struct {
	Event *CrossChainRegistryStorageOperatorTableCalculatorRemoved // Event containing the contract specifics and raw log

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
func (it *CrossChainRegistryStorageOperatorTableCalculatorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainRegistryStorageOperatorTableCalculatorRemoved)
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
		it.Event = new(CrossChainRegistryStorageOperatorTableCalculatorRemoved)
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
func (it *CrossChainRegistryStorageOperatorTableCalculatorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossChainRegistryStorageOperatorTableCalculatorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossChainRegistryStorageOperatorTableCalculatorRemoved represents a OperatorTableCalculatorRemoved event raised by the CrossChainRegistryStorage contract.
type CrossChainRegistryStorageOperatorTableCalculatorRemoved struct {
	OperatorSet OperatorSet
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorTableCalculatorRemoved is a free log retrieval operation binding the contract event 0xd7811913efd5d98fc7ea0d1fdd022b3d31987815360842d05b1d1cf55578d16a.
//
// Solidity: event OperatorTableCalculatorRemoved((address,uint32) operatorSet)
func (_CrossChainRegistryStorage *CrossChainRegistryStorageFilterer) FilterOperatorTableCalculatorRemoved(opts *bind.FilterOpts) (*CrossChainRegistryStorageOperatorTableCalculatorRemovedIterator, error) {

	logs, sub, err := _CrossChainRegistryStorage.contract.FilterLogs(opts, "OperatorTableCalculatorRemoved")
	if err != nil {
		return nil, err
	}
	return &CrossChainRegistryStorageOperatorTableCalculatorRemovedIterator{contract: _CrossChainRegistryStorage.contract, event: "OperatorTableCalculatorRemoved", logs: logs, sub: sub}, nil
}

// WatchOperatorTableCalculatorRemoved is a free log subscription operation binding the contract event 0xd7811913efd5d98fc7ea0d1fdd022b3d31987815360842d05b1d1cf55578d16a.
//
// Solidity: event OperatorTableCalculatorRemoved((address,uint32) operatorSet)
func (_CrossChainRegistryStorage *CrossChainRegistryStorageFilterer) WatchOperatorTableCalculatorRemoved(opts *bind.WatchOpts, sink chan<- *CrossChainRegistryStorageOperatorTableCalculatorRemoved) (event.Subscription, error) {

	logs, sub, err := _CrossChainRegistryStorage.contract.WatchLogs(opts, "OperatorTableCalculatorRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossChainRegistryStorageOperatorTableCalculatorRemoved)
				if err := _CrossChainRegistryStorage.contract.UnpackLog(event, "OperatorTableCalculatorRemoved", log); err != nil {
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

// ParseOperatorTableCalculatorRemoved is a log parse operation binding the contract event 0xd7811913efd5d98fc7ea0d1fdd022b3d31987815360842d05b1d1cf55578d16a.
//
// Solidity: event OperatorTableCalculatorRemoved((address,uint32) operatorSet)
func (_CrossChainRegistryStorage *CrossChainRegistryStorageFilterer) ParseOperatorTableCalculatorRemoved(log types.Log) (*CrossChainRegistryStorageOperatorTableCalculatorRemoved, error) {
	event := new(CrossChainRegistryStorageOperatorTableCalculatorRemoved)
	if err := _CrossChainRegistryStorage.contract.UnpackLog(event, "OperatorTableCalculatorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossChainRegistryStorageOperatorTableCalculatorSetIterator is returned from FilterOperatorTableCalculatorSet and is used to iterate over the raw logs and unpacked data for OperatorTableCalculatorSet events raised by the CrossChainRegistryStorage contract.
type CrossChainRegistryStorageOperatorTableCalculatorSetIterator struct {
	Event *CrossChainRegistryStorageOperatorTableCalculatorSet // Event containing the contract specifics and raw log

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
func (it *CrossChainRegistryStorageOperatorTableCalculatorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainRegistryStorageOperatorTableCalculatorSet)
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
		it.Event = new(CrossChainRegistryStorageOperatorTableCalculatorSet)
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
func (it *CrossChainRegistryStorageOperatorTableCalculatorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossChainRegistryStorageOperatorTableCalculatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossChainRegistryStorageOperatorTableCalculatorSet represents a OperatorTableCalculatorSet event raised by the CrossChainRegistryStorage contract.
type CrossChainRegistryStorageOperatorTableCalculatorSet struct {
	OperatorSet             OperatorSet
	OperatorTableCalculator common.Address
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterOperatorTableCalculatorSet is a free log retrieval operation binding the contract event 0x7f7ccafd92d20fdb39dee184a0dce002a9da420ed0def461f2a027abc9b3f6df.
//
// Solidity: event OperatorTableCalculatorSet((address,uint32) operatorSet, address operatorTableCalculator)
func (_CrossChainRegistryStorage *CrossChainRegistryStorageFilterer) FilterOperatorTableCalculatorSet(opts *bind.FilterOpts) (*CrossChainRegistryStorageOperatorTableCalculatorSetIterator, error) {

	logs, sub, err := _CrossChainRegistryStorage.contract.FilterLogs(opts, "OperatorTableCalculatorSet")
	if err != nil {
		return nil, err
	}
	return &CrossChainRegistryStorageOperatorTableCalculatorSetIterator{contract: _CrossChainRegistryStorage.contract, event: "OperatorTableCalculatorSet", logs: logs, sub: sub}, nil
}

// WatchOperatorTableCalculatorSet is a free log subscription operation binding the contract event 0x7f7ccafd92d20fdb39dee184a0dce002a9da420ed0def461f2a027abc9b3f6df.
//
// Solidity: event OperatorTableCalculatorSet((address,uint32) operatorSet, address operatorTableCalculator)
func (_CrossChainRegistryStorage *CrossChainRegistryStorageFilterer) WatchOperatorTableCalculatorSet(opts *bind.WatchOpts, sink chan<- *CrossChainRegistryStorageOperatorTableCalculatorSet) (event.Subscription, error) {

	logs, sub, err := _CrossChainRegistryStorage.contract.WatchLogs(opts, "OperatorTableCalculatorSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossChainRegistryStorageOperatorTableCalculatorSet)
				if err := _CrossChainRegistryStorage.contract.UnpackLog(event, "OperatorTableCalculatorSet", log); err != nil {
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

// ParseOperatorTableCalculatorSet is a log parse operation binding the contract event 0x7f7ccafd92d20fdb39dee184a0dce002a9da420ed0def461f2a027abc9b3f6df.
//
// Solidity: event OperatorTableCalculatorSet((address,uint32) operatorSet, address operatorTableCalculator)
func (_CrossChainRegistryStorage *CrossChainRegistryStorageFilterer) ParseOperatorTableCalculatorSet(log types.Log) (*CrossChainRegistryStorageOperatorTableCalculatorSet, error) {
	event := new(CrossChainRegistryStorageOperatorTableCalculatorSet)
	if err := _CrossChainRegistryStorage.contract.UnpackLog(event, "OperatorTableCalculatorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossChainRegistryStorageTransportDestinationChainAddedIterator is returned from FilterTransportDestinationChainAdded and is used to iterate over the raw logs and unpacked data for TransportDestinationChainAdded events raised by the CrossChainRegistryStorage contract.
type CrossChainRegistryStorageTransportDestinationChainAddedIterator struct {
	Event *CrossChainRegistryStorageTransportDestinationChainAdded // Event containing the contract specifics and raw log

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
func (it *CrossChainRegistryStorageTransportDestinationChainAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainRegistryStorageTransportDestinationChainAdded)
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
		it.Event = new(CrossChainRegistryStorageTransportDestinationChainAdded)
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
func (it *CrossChainRegistryStorageTransportDestinationChainAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossChainRegistryStorageTransportDestinationChainAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossChainRegistryStorageTransportDestinationChainAdded represents a TransportDestinationChainAdded event raised by the CrossChainRegistryStorage contract.
type CrossChainRegistryStorageTransportDestinationChainAdded struct {
	OperatorSet OperatorSet
	ChainID     *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTransportDestinationChainAdded is a free log retrieval operation binding the contract event 0x57a1fcb3d9cd447695c46f20944ba562d9547989dcddea0afb119115060c7f0b.
//
// Solidity: event TransportDestinationChainAdded((address,uint32) operatorSet, uint256 chainID)
func (_CrossChainRegistryStorage *CrossChainRegistryStorageFilterer) FilterTransportDestinationChainAdded(opts *bind.FilterOpts) (*CrossChainRegistryStorageTransportDestinationChainAddedIterator, error) {

	logs, sub, err := _CrossChainRegistryStorage.contract.FilterLogs(opts, "TransportDestinationChainAdded")
	if err != nil {
		return nil, err
	}
	return &CrossChainRegistryStorageTransportDestinationChainAddedIterator{contract: _CrossChainRegistryStorage.contract, event: "TransportDestinationChainAdded", logs: logs, sub: sub}, nil
}

// WatchTransportDestinationChainAdded is a free log subscription operation binding the contract event 0x57a1fcb3d9cd447695c46f20944ba562d9547989dcddea0afb119115060c7f0b.
//
// Solidity: event TransportDestinationChainAdded((address,uint32) operatorSet, uint256 chainID)
func (_CrossChainRegistryStorage *CrossChainRegistryStorageFilterer) WatchTransportDestinationChainAdded(opts *bind.WatchOpts, sink chan<- *CrossChainRegistryStorageTransportDestinationChainAdded) (event.Subscription, error) {

	logs, sub, err := _CrossChainRegistryStorage.contract.WatchLogs(opts, "TransportDestinationChainAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossChainRegistryStorageTransportDestinationChainAdded)
				if err := _CrossChainRegistryStorage.contract.UnpackLog(event, "TransportDestinationChainAdded", log); err != nil {
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

// ParseTransportDestinationChainAdded is a log parse operation binding the contract event 0x57a1fcb3d9cd447695c46f20944ba562d9547989dcddea0afb119115060c7f0b.
//
// Solidity: event TransportDestinationChainAdded((address,uint32) operatorSet, uint256 chainID)
func (_CrossChainRegistryStorage *CrossChainRegistryStorageFilterer) ParseTransportDestinationChainAdded(log types.Log) (*CrossChainRegistryStorageTransportDestinationChainAdded, error) {
	event := new(CrossChainRegistryStorageTransportDestinationChainAdded)
	if err := _CrossChainRegistryStorage.contract.UnpackLog(event, "TransportDestinationChainAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossChainRegistryStorageTransportDestinationChainRemovedIterator is returned from FilterTransportDestinationChainRemoved and is used to iterate over the raw logs and unpacked data for TransportDestinationChainRemoved events raised by the CrossChainRegistryStorage contract.
type CrossChainRegistryStorageTransportDestinationChainRemovedIterator struct {
	Event *CrossChainRegistryStorageTransportDestinationChainRemoved // Event containing the contract specifics and raw log

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
func (it *CrossChainRegistryStorageTransportDestinationChainRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainRegistryStorageTransportDestinationChainRemoved)
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
		it.Event = new(CrossChainRegistryStorageTransportDestinationChainRemoved)
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
func (it *CrossChainRegistryStorageTransportDestinationChainRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossChainRegistryStorageTransportDestinationChainRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossChainRegistryStorageTransportDestinationChainRemoved represents a TransportDestinationChainRemoved event raised by the CrossChainRegistryStorage contract.
type CrossChainRegistryStorageTransportDestinationChainRemoved struct {
	OperatorSet OperatorSet
	ChainID     *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTransportDestinationChainRemoved is a free log retrieval operation binding the contract event 0x499955d838e6f0ca31e83adf81d191cfe6cd8fe252bf826c75c9a80ba077e25e.
//
// Solidity: event TransportDestinationChainRemoved((address,uint32) operatorSet, uint256 chainID)
func (_CrossChainRegistryStorage *CrossChainRegistryStorageFilterer) FilterTransportDestinationChainRemoved(opts *bind.FilterOpts) (*CrossChainRegistryStorageTransportDestinationChainRemovedIterator, error) {

	logs, sub, err := _CrossChainRegistryStorage.contract.FilterLogs(opts, "TransportDestinationChainRemoved")
	if err != nil {
		return nil, err
	}
	return &CrossChainRegistryStorageTransportDestinationChainRemovedIterator{contract: _CrossChainRegistryStorage.contract, event: "TransportDestinationChainRemoved", logs: logs, sub: sub}, nil
}

// WatchTransportDestinationChainRemoved is a free log subscription operation binding the contract event 0x499955d838e6f0ca31e83adf81d191cfe6cd8fe252bf826c75c9a80ba077e25e.
//
// Solidity: event TransportDestinationChainRemoved((address,uint32) operatorSet, uint256 chainID)
func (_CrossChainRegistryStorage *CrossChainRegistryStorageFilterer) WatchTransportDestinationChainRemoved(opts *bind.WatchOpts, sink chan<- *CrossChainRegistryStorageTransportDestinationChainRemoved) (event.Subscription, error) {

	logs, sub, err := _CrossChainRegistryStorage.contract.WatchLogs(opts, "TransportDestinationChainRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossChainRegistryStorageTransportDestinationChainRemoved)
				if err := _CrossChainRegistryStorage.contract.UnpackLog(event, "TransportDestinationChainRemoved", log); err != nil {
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

// ParseTransportDestinationChainRemoved is a log parse operation binding the contract event 0x499955d838e6f0ca31e83adf81d191cfe6cd8fe252bf826c75c9a80ba077e25e.
//
// Solidity: event TransportDestinationChainRemoved((address,uint32) operatorSet, uint256 chainID)
func (_CrossChainRegistryStorage *CrossChainRegistryStorageFilterer) ParseTransportDestinationChainRemoved(log types.Log) (*CrossChainRegistryStorageTransportDestinationChainRemoved, error) {
	event := new(CrossChainRegistryStorageTransportDestinationChainRemoved)
	if err := _CrossChainRegistryStorage.contract.UnpackLog(event, "TransportDestinationChainRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossChainRegistryStorageTransportDestinationsRemovedIterator is returned from FilterTransportDestinationsRemoved and is used to iterate over the raw logs and unpacked data for TransportDestinationsRemoved events raised by the CrossChainRegistryStorage contract.
type CrossChainRegistryStorageTransportDestinationsRemovedIterator struct {
	Event *CrossChainRegistryStorageTransportDestinationsRemoved // Event containing the contract specifics and raw log

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
func (it *CrossChainRegistryStorageTransportDestinationsRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainRegistryStorageTransportDestinationsRemoved)
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
		it.Event = new(CrossChainRegistryStorageTransportDestinationsRemoved)
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
func (it *CrossChainRegistryStorageTransportDestinationsRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossChainRegistryStorageTransportDestinationsRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossChainRegistryStorageTransportDestinationsRemoved represents a TransportDestinationsRemoved event raised by the CrossChainRegistryStorage contract.
type CrossChainRegistryStorageTransportDestinationsRemoved struct {
	OperatorSet OperatorSet
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTransportDestinationsRemoved is a free log retrieval operation binding the contract event 0xaf209f19ac00e8ccb4539e96d4141cdc96fea479d258d99910307c7365e68759.
//
// Solidity: event TransportDestinationsRemoved((address,uint32) operatorSet)
func (_CrossChainRegistryStorage *CrossChainRegistryStorageFilterer) FilterTransportDestinationsRemoved(opts *bind.FilterOpts) (*CrossChainRegistryStorageTransportDestinationsRemovedIterator, error) {

	logs, sub, err := _CrossChainRegistryStorage.contract.FilterLogs(opts, "TransportDestinationsRemoved")
	if err != nil {
		return nil, err
	}
	return &CrossChainRegistryStorageTransportDestinationsRemovedIterator{contract: _CrossChainRegistryStorage.contract, event: "TransportDestinationsRemoved", logs: logs, sub: sub}, nil
}

// WatchTransportDestinationsRemoved is a free log subscription operation binding the contract event 0xaf209f19ac00e8ccb4539e96d4141cdc96fea479d258d99910307c7365e68759.
//
// Solidity: event TransportDestinationsRemoved((address,uint32) operatorSet)
func (_CrossChainRegistryStorage *CrossChainRegistryStorageFilterer) WatchTransportDestinationsRemoved(opts *bind.WatchOpts, sink chan<- *CrossChainRegistryStorageTransportDestinationsRemoved) (event.Subscription, error) {

	logs, sub, err := _CrossChainRegistryStorage.contract.WatchLogs(opts, "TransportDestinationsRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossChainRegistryStorageTransportDestinationsRemoved)
				if err := _CrossChainRegistryStorage.contract.UnpackLog(event, "TransportDestinationsRemoved", log); err != nil {
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

// ParseTransportDestinationsRemoved is a log parse operation binding the contract event 0xaf209f19ac00e8ccb4539e96d4141cdc96fea479d258d99910307c7365e68759.
//
// Solidity: event TransportDestinationsRemoved((address,uint32) operatorSet)
func (_CrossChainRegistryStorage *CrossChainRegistryStorageFilterer) ParseTransportDestinationsRemoved(log types.Log) (*CrossChainRegistryStorageTransportDestinationsRemoved, error) {
	event := new(CrossChainRegistryStorageTransportDestinationsRemoved)
	if err := _CrossChainRegistryStorage.contract.UnpackLog(event, "TransportDestinationsRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
