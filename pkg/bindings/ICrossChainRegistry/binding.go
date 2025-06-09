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

// ICrossChainRegistryMetaData contains all meta data concerning the ICrossChainRegistry contract.
var ICrossChainRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"addChainIDsToWhitelist\",\"inputs\":[{\"name\":\"chainIDs\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"operatorTableUpdaters\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addTransportDestinations\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"chainIDs\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"calculateOperatorTableBytes\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createGenerationReservation\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operatorTableCalculator\",\"type\":\"address\",\"internalType\":\"contractIOperatorTableCalculator\"},{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structICrossChainRegistryTypes.OperatorSetConfig\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxStalenessPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"chainIDs\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getActiveGenerationReservations\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getActiveTransportReservations\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"\",\"type\":\"uint256[][]\",\"internalType\":\"uint256[][]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorSetConfig\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structICrossChainRegistryTypes.OperatorSetConfig\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxStalenessPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorTableCalculator\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIOperatorTableCalculator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTransportDestinations\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeChainIDsFromWhitelist\",\"inputs\":[{\"name\":\"chainIDs\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeGenerationReservation\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeTransportDestinations\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"chainIDs\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOperatorSetConfig\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structICrossChainRegistryTypes.OperatorSetConfig\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxStalenessPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOperatorTableCalculator\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operatorTableCalculator\",\"type\":\"address\",\"internalType\":\"contractIOperatorTableCalculator\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ChainIDAddedToWhitelist\",\"inputs\":[{\"name\":\"chainID\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"operatorTableUpdater\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainIDRemovedFromWhitelist\",\"inputs\":[{\"name\":\"chainID\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GenerationReservationCreated\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GenerationReservationRemoved\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSetConfigSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structICrossChainRegistryTypes.OperatorSetConfig\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxStalenessPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorTableCalculatorSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operatorTableCalculator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIOperatorTableCalculator\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TransportDestinationAdded\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"chainID\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TransportDestinationRemoved\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"chainID\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ArrayLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainIDAlreadyWhitelisted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainIDNotWhitelisted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EmptyChainIDsArray\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"GenerationReservationAlreadyExists\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"GenerationReservationDoesNotExist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidChainId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperatorSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperatorTableCalculator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NeedToDelete\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RequireAtLeastOneTransportDestination\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StalenessPeriodZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TransportDestinationAlreadyAdded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TransportDestinationNotFound\",\"inputs\":[]}]",
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

// CalculateOperatorTableBytes is a free data retrieval call binding the contract method 0x41ee6d0e.
//
// Solidity: function calculateOperatorTableBytes((address,uint32) operatorSet) view returns(bytes)
func (_ICrossChainRegistry *ICrossChainRegistryCaller) CalculateOperatorTableBytes(opts *bind.CallOpts, operatorSet OperatorSet) ([]byte, error) {
	var out []interface{}
	err := _ICrossChainRegistry.contract.Call(opts, &out, "calculateOperatorTableBytes", operatorSet)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// CalculateOperatorTableBytes is a free data retrieval call binding the contract method 0x41ee6d0e.
//
// Solidity: function calculateOperatorTableBytes((address,uint32) operatorSet) view returns(bytes)
func (_ICrossChainRegistry *ICrossChainRegistrySession) CalculateOperatorTableBytes(operatorSet OperatorSet) ([]byte, error) {
	return _ICrossChainRegistry.Contract.CalculateOperatorTableBytes(&_ICrossChainRegistry.CallOpts, operatorSet)
}

// CalculateOperatorTableBytes is a free data retrieval call binding the contract method 0x41ee6d0e.
//
// Solidity: function calculateOperatorTableBytes((address,uint32) operatorSet) view returns(bytes)
func (_ICrossChainRegistry *ICrossChainRegistryCallerSession) CalculateOperatorTableBytes(operatorSet OperatorSet) ([]byte, error) {
	return _ICrossChainRegistry.Contract.CalculateOperatorTableBytes(&_ICrossChainRegistry.CallOpts, operatorSet)
}

// GetActiveGenerationReservations is a free data retrieval call binding the contract method 0xd09b978b.
//
// Solidity: function getActiveGenerationReservations() view returns((address,uint32)[])
func (_ICrossChainRegistry *ICrossChainRegistryCaller) GetActiveGenerationReservations(opts *bind.CallOpts) ([]OperatorSet, error) {
	var out []interface{}
	err := _ICrossChainRegistry.contract.Call(opts, &out, "getActiveGenerationReservations")

	if err != nil {
		return *new([]OperatorSet), err
	}

	out0 := *abi.ConvertType(out[0], new([]OperatorSet)).(*[]OperatorSet)

	return out0, err

}

// GetActiveGenerationReservations is a free data retrieval call binding the contract method 0xd09b978b.
//
// Solidity: function getActiveGenerationReservations() view returns((address,uint32)[])
func (_ICrossChainRegistry *ICrossChainRegistrySession) GetActiveGenerationReservations() ([]OperatorSet, error) {
	return _ICrossChainRegistry.Contract.GetActiveGenerationReservations(&_ICrossChainRegistry.CallOpts)
}

// GetActiveGenerationReservations is a free data retrieval call binding the contract method 0xd09b978b.
//
// Solidity: function getActiveGenerationReservations() view returns((address,uint32)[])
func (_ICrossChainRegistry *ICrossChainRegistryCallerSession) GetActiveGenerationReservations() ([]OperatorSet, error) {
	return _ICrossChainRegistry.Contract.GetActiveGenerationReservations(&_ICrossChainRegistry.CallOpts)
}

// GetActiveTransportReservations is a free data retrieval call binding the contract method 0xbfda3b3d.
//
// Solidity: function getActiveTransportReservations() view returns((address,uint32)[], uint256[][])
func (_ICrossChainRegistry *ICrossChainRegistryCaller) GetActiveTransportReservations(opts *bind.CallOpts) ([]OperatorSet, [][]*big.Int, error) {
	var out []interface{}
	err := _ICrossChainRegistry.contract.Call(opts, &out, "getActiveTransportReservations")

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
func (_ICrossChainRegistry *ICrossChainRegistrySession) GetActiveTransportReservations() ([]OperatorSet, [][]*big.Int, error) {
	return _ICrossChainRegistry.Contract.GetActiveTransportReservations(&_ICrossChainRegistry.CallOpts)
}

// GetActiveTransportReservations is a free data retrieval call binding the contract method 0xbfda3b3d.
//
// Solidity: function getActiveTransportReservations() view returns((address,uint32)[], uint256[][])
func (_ICrossChainRegistry *ICrossChainRegistryCallerSession) GetActiveTransportReservations() ([]OperatorSet, [][]*big.Int, error) {
	return _ICrossChainRegistry.Contract.GetActiveTransportReservations(&_ICrossChainRegistry.CallOpts)
}

// GetOperatorSetConfig is a free data retrieval call binding the contract method 0x21fa7fdc.
//
// Solidity: function getOperatorSetConfig((address,uint32) operatorSet) view returns((address,uint32))
func (_ICrossChainRegistry *ICrossChainRegistryCaller) GetOperatorSetConfig(opts *bind.CallOpts, operatorSet OperatorSet) (ICrossChainRegistryTypesOperatorSetConfig, error) {
	var out []interface{}
	err := _ICrossChainRegistry.contract.Call(opts, &out, "getOperatorSetConfig", operatorSet)

	if err != nil {
		return *new(ICrossChainRegistryTypesOperatorSetConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(ICrossChainRegistryTypesOperatorSetConfig)).(*ICrossChainRegistryTypesOperatorSetConfig)

	return out0, err

}

// GetOperatorSetConfig is a free data retrieval call binding the contract method 0x21fa7fdc.
//
// Solidity: function getOperatorSetConfig((address,uint32) operatorSet) view returns((address,uint32))
func (_ICrossChainRegistry *ICrossChainRegistrySession) GetOperatorSetConfig(operatorSet OperatorSet) (ICrossChainRegistryTypesOperatorSetConfig, error) {
	return _ICrossChainRegistry.Contract.GetOperatorSetConfig(&_ICrossChainRegistry.CallOpts, operatorSet)
}

// GetOperatorSetConfig is a free data retrieval call binding the contract method 0x21fa7fdc.
//
// Solidity: function getOperatorSetConfig((address,uint32) operatorSet) view returns((address,uint32))
func (_ICrossChainRegistry *ICrossChainRegistryCallerSession) GetOperatorSetConfig(operatorSet OperatorSet) (ICrossChainRegistryTypesOperatorSetConfig, error) {
	return _ICrossChainRegistry.Contract.GetOperatorSetConfig(&_ICrossChainRegistry.CallOpts, operatorSet)
}

// GetOperatorTableCalculator is a free data retrieval call binding the contract method 0x75e4b539.
//
// Solidity: function getOperatorTableCalculator((address,uint32) operatorSet) view returns(address)
func (_ICrossChainRegistry *ICrossChainRegistryCaller) GetOperatorTableCalculator(opts *bind.CallOpts, operatorSet OperatorSet) (common.Address, error) {
	var out []interface{}
	err := _ICrossChainRegistry.contract.Call(opts, &out, "getOperatorTableCalculator", operatorSet)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOperatorTableCalculator is a free data retrieval call binding the contract method 0x75e4b539.
//
// Solidity: function getOperatorTableCalculator((address,uint32) operatorSet) view returns(address)
func (_ICrossChainRegistry *ICrossChainRegistrySession) GetOperatorTableCalculator(operatorSet OperatorSet) (common.Address, error) {
	return _ICrossChainRegistry.Contract.GetOperatorTableCalculator(&_ICrossChainRegistry.CallOpts, operatorSet)
}

// GetOperatorTableCalculator is a free data retrieval call binding the contract method 0x75e4b539.
//
// Solidity: function getOperatorTableCalculator((address,uint32) operatorSet) view returns(address)
func (_ICrossChainRegistry *ICrossChainRegistryCallerSession) GetOperatorTableCalculator(operatorSet OperatorSet) (common.Address, error) {
	return _ICrossChainRegistry.Contract.GetOperatorTableCalculator(&_ICrossChainRegistry.CallOpts, operatorSet)
}

// GetSupportedChains is a free data retrieval call binding the contract method 0xc4bffe2b.
//
// Solidity: function getSupportedChains() view returns(uint256[], address[])
func (_ICrossChainRegistry *ICrossChainRegistryCaller) GetSupportedChains(opts *bind.CallOpts) ([]*big.Int, []common.Address, error) {
	var out []interface{}
	err := _ICrossChainRegistry.contract.Call(opts, &out, "getSupportedChains")

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
func (_ICrossChainRegistry *ICrossChainRegistrySession) GetSupportedChains() ([]*big.Int, []common.Address, error) {
	return _ICrossChainRegistry.Contract.GetSupportedChains(&_ICrossChainRegistry.CallOpts)
}

// GetSupportedChains is a free data retrieval call binding the contract method 0xc4bffe2b.
//
// Solidity: function getSupportedChains() view returns(uint256[], address[])
func (_ICrossChainRegistry *ICrossChainRegistryCallerSession) GetSupportedChains() ([]*big.Int, []common.Address, error) {
	return _ICrossChainRegistry.Contract.GetSupportedChains(&_ICrossChainRegistry.CallOpts)
}

// GetTransportDestinations is a free data retrieval call binding the contract method 0x3c75fddf.
//
// Solidity: function getTransportDestinations((address,uint32) operatorSet) view returns(uint256[])
func (_ICrossChainRegistry *ICrossChainRegistryCaller) GetTransportDestinations(opts *bind.CallOpts, operatorSet OperatorSet) ([]*big.Int, error) {
	var out []interface{}
	err := _ICrossChainRegistry.contract.Call(opts, &out, "getTransportDestinations", operatorSet)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetTransportDestinations is a free data retrieval call binding the contract method 0x3c75fddf.
//
// Solidity: function getTransportDestinations((address,uint32) operatorSet) view returns(uint256[])
func (_ICrossChainRegistry *ICrossChainRegistrySession) GetTransportDestinations(operatorSet OperatorSet) ([]*big.Int, error) {
	return _ICrossChainRegistry.Contract.GetTransportDestinations(&_ICrossChainRegistry.CallOpts, operatorSet)
}

// GetTransportDestinations is a free data retrieval call binding the contract method 0x3c75fddf.
//
// Solidity: function getTransportDestinations((address,uint32) operatorSet) view returns(uint256[])
func (_ICrossChainRegistry *ICrossChainRegistryCallerSession) GetTransportDestinations(operatorSet OperatorSet) ([]*big.Int, error) {
	return _ICrossChainRegistry.Contract.GetTransportDestinations(&_ICrossChainRegistry.CallOpts, operatorSet)
}

// AddChainIDsToWhitelist is a paid mutator transaction binding the contract method 0x04e98be3.
//
// Solidity: function addChainIDsToWhitelist(uint256[] chainIDs, address[] operatorTableUpdaters) returns()
func (_ICrossChainRegistry *ICrossChainRegistryTransactor) AddChainIDsToWhitelist(opts *bind.TransactOpts, chainIDs []*big.Int, operatorTableUpdaters []common.Address) (*types.Transaction, error) {
	return _ICrossChainRegistry.contract.Transact(opts, "addChainIDsToWhitelist", chainIDs, operatorTableUpdaters)
}

// AddChainIDsToWhitelist is a paid mutator transaction binding the contract method 0x04e98be3.
//
// Solidity: function addChainIDsToWhitelist(uint256[] chainIDs, address[] operatorTableUpdaters) returns()
func (_ICrossChainRegistry *ICrossChainRegistrySession) AddChainIDsToWhitelist(chainIDs []*big.Int, operatorTableUpdaters []common.Address) (*types.Transaction, error) {
	return _ICrossChainRegistry.Contract.AddChainIDsToWhitelist(&_ICrossChainRegistry.TransactOpts, chainIDs, operatorTableUpdaters)
}

// AddChainIDsToWhitelist is a paid mutator transaction binding the contract method 0x04e98be3.
//
// Solidity: function addChainIDsToWhitelist(uint256[] chainIDs, address[] operatorTableUpdaters) returns()
func (_ICrossChainRegistry *ICrossChainRegistryTransactorSession) AddChainIDsToWhitelist(chainIDs []*big.Int, operatorTableUpdaters []common.Address) (*types.Transaction, error) {
	return _ICrossChainRegistry.Contract.AddChainIDsToWhitelist(&_ICrossChainRegistry.TransactOpts, chainIDs, operatorTableUpdaters)
}

// AddTransportDestinations is a paid mutator transaction binding the contract method 0x49be7d6f.
//
// Solidity: function addTransportDestinations((address,uint32) operatorSet, uint256[] chainIDs) returns()
func (_ICrossChainRegistry *ICrossChainRegistryTransactor) AddTransportDestinations(opts *bind.TransactOpts, operatorSet OperatorSet, chainIDs []*big.Int) (*types.Transaction, error) {
	return _ICrossChainRegistry.contract.Transact(opts, "addTransportDestinations", operatorSet, chainIDs)
}

// AddTransportDestinations is a paid mutator transaction binding the contract method 0x49be7d6f.
//
// Solidity: function addTransportDestinations((address,uint32) operatorSet, uint256[] chainIDs) returns()
func (_ICrossChainRegistry *ICrossChainRegistrySession) AddTransportDestinations(operatorSet OperatorSet, chainIDs []*big.Int) (*types.Transaction, error) {
	return _ICrossChainRegistry.Contract.AddTransportDestinations(&_ICrossChainRegistry.TransactOpts, operatorSet, chainIDs)
}

// AddTransportDestinations is a paid mutator transaction binding the contract method 0x49be7d6f.
//
// Solidity: function addTransportDestinations((address,uint32) operatorSet, uint256[] chainIDs) returns()
func (_ICrossChainRegistry *ICrossChainRegistryTransactorSession) AddTransportDestinations(operatorSet OperatorSet, chainIDs []*big.Int) (*types.Transaction, error) {
	return _ICrossChainRegistry.Contract.AddTransportDestinations(&_ICrossChainRegistry.TransactOpts, operatorSet, chainIDs)
}

// CreateGenerationReservation is a paid mutator transaction binding the contract method 0xfe596dee.
//
// Solidity: function createGenerationReservation((address,uint32) operatorSet, address operatorTableCalculator, (address,uint32) config, uint256[] chainIDs) returns()
func (_ICrossChainRegistry *ICrossChainRegistryTransactor) CreateGenerationReservation(opts *bind.TransactOpts, operatorSet OperatorSet, operatorTableCalculator common.Address, config ICrossChainRegistryTypesOperatorSetConfig, chainIDs []*big.Int) (*types.Transaction, error) {
	return _ICrossChainRegistry.contract.Transact(opts, "createGenerationReservation", operatorSet, operatorTableCalculator, config, chainIDs)
}

// CreateGenerationReservation is a paid mutator transaction binding the contract method 0xfe596dee.
//
// Solidity: function createGenerationReservation((address,uint32) operatorSet, address operatorTableCalculator, (address,uint32) config, uint256[] chainIDs) returns()
func (_ICrossChainRegistry *ICrossChainRegistrySession) CreateGenerationReservation(operatorSet OperatorSet, operatorTableCalculator common.Address, config ICrossChainRegistryTypesOperatorSetConfig, chainIDs []*big.Int) (*types.Transaction, error) {
	return _ICrossChainRegistry.Contract.CreateGenerationReservation(&_ICrossChainRegistry.TransactOpts, operatorSet, operatorTableCalculator, config, chainIDs)
}

// CreateGenerationReservation is a paid mutator transaction binding the contract method 0xfe596dee.
//
// Solidity: function createGenerationReservation((address,uint32) operatorSet, address operatorTableCalculator, (address,uint32) config, uint256[] chainIDs) returns()
func (_ICrossChainRegistry *ICrossChainRegistryTransactorSession) CreateGenerationReservation(operatorSet OperatorSet, operatorTableCalculator common.Address, config ICrossChainRegistryTypesOperatorSetConfig, chainIDs []*big.Int) (*types.Transaction, error) {
	return _ICrossChainRegistry.Contract.CreateGenerationReservation(&_ICrossChainRegistry.TransactOpts, operatorSet, operatorTableCalculator, config, chainIDs)
}

// RemoveChainIDsFromWhitelist is a paid mutator transaction binding the contract method 0xdfbd9dfd.
//
// Solidity: function removeChainIDsFromWhitelist(uint256[] chainIDs) returns()
func (_ICrossChainRegistry *ICrossChainRegistryTransactor) RemoveChainIDsFromWhitelist(opts *bind.TransactOpts, chainIDs []*big.Int) (*types.Transaction, error) {
	return _ICrossChainRegistry.contract.Transact(opts, "removeChainIDsFromWhitelist", chainIDs)
}

// RemoveChainIDsFromWhitelist is a paid mutator transaction binding the contract method 0xdfbd9dfd.
//
// Solidity: function removeChainIDsFromWhitelist(uint256[] chainIDs) returns()
func (_ICrossChainRegistry *ICrossChainRegistrySession) RemoveChainIDsFromWhitelist(chainIDs []*big.Int) (*types.Transaction, error) {
	return _ICrossChainRegistry.Contract.RemoveChainIDsFromWhitelist(&_ICrossChainRegistry.TransactOpts, chainIDs)
}

// RemoveChainIDsFromWhitelist is a paid mutator transaction binding the contract method 0xdfbd9dfd.
//
// Solidity: function removeChainIDsFromWhitelist(uint256[] chainIDs) returns()
func (_ICrossChainRegistry *ICrossChainRegistryTransactorSession) RemoveChainIDsFromWhitelist(chainIDs []*big.Int) (*types.Transaction, error) {
	return _ICrossChainRegistry.Contract.RemoveChainIDsFromWhitelist(&_ICrossChainRegistry.TransactOpts, chainIDs)
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

// RemoveTransportDestinations is a paid mutator transaction binding the contract method 0xf3e9f5d4.
//
// Solidity: function removeTransportDestinations((address,uint32) operatorSet, uint256[] chainIDs) returns()
func (_ICrossChainRegistry *ICrossChainRegistryTransactor) RemoveTransportDestinations(opts *bind.TransactOpts, operatorSet OperatorSet, chainIDs []*big.Int) (*types.Transaction, error) {
	return _ICrossChainRegistry.contract.Transact(opts, "removeTransportDestinations", operatorSet, chainIDs)
}

// RemoveTransportDestinations is a paid mutator transaction binding the contract method 0xf3e9f5d4.
//
// Solidity: function removeTransportDestinations((address,uint32) operatorSet, uint256[] chainIDs) returns()
func (_ICrossChainRegistry *ICrossChainRegistrySession) RemoveTransportDestinations(operatorSet OperatorSet, chainIDs []*big.Int) (*types.Transaction, error) {
	return _ICrossChainRegistry.Contract.RemoveTransportDestinations(&_ICrossChainRegistry.TransactOpts, operatorSet, chainIDs)
}

// RemoveTransportDestinations is a paid mutator transaction binding the contract method 0xf3e9f5d4.
//
// Solidity: function removeTransportDestinations((address,uint32) operatorSet, uint256[] chainIDs) returns()
func (_ICrossChainRegistry *ICrossChainRegistryTransactorSession) RemoveTransportDestinations(operatorSet OperatorSet, chainIDs []*big.Int) (*types.Transaction, error) {
	return _ICrossChainRegistry.Contract.RemoveTransportDestinations(&_ICrossChainRegistry.TransactOpts, operatorSet, chainIDs)
}

// SetOperatorSetConfig is a paid mutator transaction binding the contract method 0x277e1e62.
//
// Solidity: function setOperatorSetConfig((address,uint32) operatorSet, (address,uint32) config) returns()
func (_ICrossChainRegistry *ICrossChainRegistryTransactor) SetOperatorSetConfig(opts *bind.TransactOpts, operatorSet OperatorSet, config ICrossChainRegistryTypesOperatorSetConfig) (*types.Transaction, error) {
	return _ICrossChainRegistry.contract.Transact(opts, "setOperatorSetConfig", operatorSet, config)
}

// SetOperatorSetConfig is a paid mutator transaction binding the contract method 0x277e1e62.
//
// Solidity: function setOperatorSetConfig((address,uint32) operatorSet, (address,uint32) config) returns()
func (_ICrossChainRegistry *ICrossChainRegistrySession) SetOperatorSetConfig(operatorSet OperatorSet, config ICrossChainRegistryTypesOperatorSetConfig) (*types.Transaction, error) {
	return _ICrossChainRegistry.Contract.SetOperatorSetConfig(&_ICrossChainRegistry.TransactOpts, operatorSet, config)
}

// SetOperatorSetConfig is a paid mutator transaction binding the contract method 0x277e1e62.
//
// Solidity: function setOperatorSetConfig((address,uint32) operatorSet, (address,uint32) config) returns()
func (_ICrossChainRegistry *ICrossChainRegistryTransactorSession) SetOperatorSetConfig(operatorSet OperatorSet, config ICrossChainRegistryTypesOperatorSetConfig) (*types.Transaction, error) {
	return _ICrossChainRegistry.Contract.SetOperatorSetConfig(&_ICrossChainRegistry.TransactOpts, operatorSet, config)
}

// SetOperatorTableCalculator is a paid mutator transaction binding the contract method 0x1ca9142a.
//
// Solidity: function setOperatorTableCalculator((address,uint32) operatorSet, address operatorTableCalculator) returns()
func (_ICrossChainRegistry *ICrossChainRegistryTransactor) SetOperatorTableCalculator(opts *bind.TransactOpts, operatorSet OperatorSet, operatorTableCalculator common.Address) (*types.Transaction, error) {
	return _ICrossChainRegistry.contract.Transact(opts, "setOperatorTableCalculator", operatorSet, operatorTableCalculator)
}

// SetOperatorTableCalculator is a paid mutator transaction binding the contract method 0x1ca9142a.
//
// Solidity: function setOperatorTableCalculator((address,uint32) operatorSet, address operatorTableCalculator) returns()
func (_ICrossChainRegistry *ICrossChainRegistrySession) SetOperatorTableCalculator(operatorSet OperatorSet, operatorTableCalculator common.Address) (*types.Transaction, error) {
	return _ICrossChainRegistry.Contract.SetOperatorTableCalculator(&_ICrossChainRegistry.TransactOpts, operatorSet, operatorTableCalculator)
}

// SetOperatorTableCalculator is a paid mutator transaction binding the contract method 0x1ca9142a.
//
// Solidity: function setOperatorTableCalculator((address,uint32) operatorSet, address operatorTableCalculator) returns()
func (_ICrossChainRegistry *ICrossChainRegistryTransactorSession) SetOperatorTableCalculator(operatorSet OperatorSet, operatorTableCalculator common.Address) (*types.Transaction, error) {
	return _ICrossChainRegistry.Contract.SetOperatorTableCalculator(&_ICrossChainRegistry.TransactOpts, operatorSet, operatorTableCalculator)
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
	ChainID              *big.Int
	OperatorTableUpdater common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterChainIDAddedToWhitelist is a free log retrieval operation binding the contract event 0x7a0a76d85b582b17996dd7371a407aa7a79b870db8539247fba315c7b6beff62.
//
// Solidity: event ChainIDAddedToWhitelist(uint256 chainID, address operatorTableUpdater)
func (_ICrossChainRegistry *ICrossChainRegistryFilterer) FilterChainIDAddedToWhitelist(opts *bind.FilterOpts) (*ICrossChainRegistryChainIDAddedToWhitelistIterator, error) {

	logs, sub, err := _ICrossChainRegistry.contract.FilterLogs(opts, "ChainIDAddedToWhitelist")
	if err != nil {
		return nil, err
	}
	return &ICrossChainRegistryChainIDAddedToWhitelistIterator{contract: _ICrossChainRegistry.contract, event: "ChainIDAddedToWhitelist", logs: logs, sub: sub}, nil
}

// WatchChainIDAddedToWhitelist is a free log subscription operation binding the contract event 0x7a0a76d85b582b17996dd7371a407aa7a79b870db8539247fba315c7b6beff62.
//
// Solidity: event ChainIDAddedToWhitelist(uint256 chainID, address operatorTableUpdater)
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

// ParseChainIDAddedToWhitelist is a log parse operation binding the contract event 0x7a0a76d85b582b17996dd7371a407aa7a79b870db8539247fba315c7b6beff62.
//
// Solidity: event ChainIDAddedToWhitelist(uint256 chainID, address operatorTableUpdater)
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
	ChainID *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterChainIDRemovedFromWhitelist is a free log retrieval operation binding the contract event 0x6824d36084ecf2cd819b137cb5d837cc6e73afce1e0e348c9fdecaa81d0341e5.
//
// Solidity: event ChainIDRemovedFromWhitelist(uint256 chainID)
func (_ICrossChainRegistry *ICrossChainRegistryFilterer) FilterChainIDRemovedFromWhitelist(opts *bind.FilterOpts) (*ICrossChainRegistryChainIDRemovedFromWhitelistIterator, error) {

	logs, sub, err := _ICrossChainRegistry.contract.FilterLogs(opts, "ChainIDRemovedFromWhitelist")
	if err != nil {
		return nil, err
	}
	return &ICrossChainRegistryChainIDRemovedFromWhitelistIterator{contract: _ICrossChainRegistry.contract, event: "ChainIDRemovedFromWhitelist", logs: logs, sub: sub}, nil
}

// WatchChainIDRemovedFromWhitelist is a free log subscription operation binding the contract event 0x6824d36084ecf2cd819b137cb5d837cc6e73afce1e0e348c9fdecaa81d0341e5.
//
// Solidity: event ChainIDRemovedFromWhitelist(uint256 chainID)
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

// ParseChainIDRemovedFromWhitelist is a log parse operation binding the contract event 0x6824d36084ecf2cd819b137cb5d837cc6e73afce1e0e348c9fdecaa81d0341e5.
//
// Solidity: event ChainIDRemovedFromWhitelist(uint256 chainID)
func (_ICrossChainRegistry *ICrossChainRegistryFilterer) ParseChainIDRemovedFromWhitelist(log types.Log) (*ICrossChainRegistryChainIDRemovedFromWhitelist, error) {
	event := new(ICrossChainRegistryChainIDRemovedFromWhitelist)
	if err := _ICrossChainRegistry.contract.UnpackLog(event, "ChainIDRemovedFromWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICrossChainRegistryGenerationReservationCreatedIterator is returned from FilterGenerationReservationCreated and is used to iterate over the raw logs and unpacked data for GenerationReservationCreated events raised by the ICrossChainRegistry contract.
type ICrossChainRegistryGenerationReservationCreatedIterator struct {
	Event *ICrossChainRegistryGenerationReservationCreated // Event containing the contract specifics and raw log

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
func (it *ICrossChainRegistryGenerationReservationCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICrossChainRegistryGenerationReservationCreated)
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
		it.Event = new(ICrossChainRegistryGenerationReservationCreated)
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
func (it *ICrossChainRegistryGenerationReservationCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICrossChainRegistryGenerationReservationCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICrossChainRegistryGenerationReservationCreated represents a GenerationReservationCreated event raised by the ICrossChainRegistry contract.
type ICrossChainRegistryGenerationReservationCreated struct {
	OperatorSet OperatorSet
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterGenerationReservationCreated is a free log retrieval operation binding the contract event 0x4fb6efec7dd60036ce3a7af8d5c48425019daa0fb61eb471a966a7ac2c6fa6a6.
//
// Solidity: event GenerationReservationCreated((address,uint32) operatorSet)
func (_ICrossChainRegistry *ICrossChainRegistryFilterer) FilterGenerationReservationCreated(opts *bind.FilterOpts) (*ICrossChainRegistryGenerationReservationCreatedIterator, error) {

	logs, sub, err := _ICrossChainRegistry.contract.FilterLogs(opts, "GenerationReservationCreated")
	if err != nil {
		return nil, err
	}
	return &ICrossChainRegistryGenerationReservationCreatedIterator{contract: _ICrossChainRegistry.contract, event: "GenerationReservationCreated", logs: logs, sub: sub}, nil
}

// WatchGenerationReservationCreated is a free log subscription operation binding the contract event 0x4fb6efec7dd60036ce3a7af8d5c48425019daa0fb61eb471a966a7ac2c6fa6a6.
//
// Solidity: event GenerationReservationCreated((address,uint32) operatorSet)
func (_ICrossChainRegistry *ICrossChainRegistryFilterer) WatchGenerationReservationCreated(opts *bind.WatchOpts, sink chan<- *ICrossChainRegistryGenerationReservationCreated) (event.Subscription, error) {

	logs, sub, err := _ICrossChainRegistry.contract.WatchLogs(opts, "GenerationReservationCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICrossChainRegistryGenerationReservationCreated)
				if err := _ICrossChainRegistry.contract.UnpackLog(event, "GenerationReservationCreated", log); err != nil {
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
func (_ICrossChainRegistry *ICrossChainRegistryFilterer) ParseGenerationReservationCreated(log types.Log) (*ICrossChainRegistryGenerationReservationCreated, error) {
	event := new(ICrossChainRegistryGenerationReservationCreated)
	if err := _ICrossChainRegistry.contract.UnpackLog(event, "GenerationReservationCreated", log); err != nil {
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
	OperatorSet OperatorSet
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterGenerationReservationRemoved is a free log retrieval operation binding the contract event 0x4ffdfdd59e9e1e3c301608788f78dd458e61cb8c045ca92b62a7b484c80824fb.
//
// Solidity: event GenerationReservationRemoved((address,uint32) operatorSet)
func (_ICrossChainRegistry *ICrossChainRegistryFilterer) FilterGenerationReservationRemoved(opts *bind.FilterOpts) (*ICrossChainRegistryGenerationReservationRemovedIterator, error) {

	logs, sub, err := _ICrossChainRegistry.contract.FilterLogs(opts, "GenerationReservationRemoved")
	if err != nil {
		return nil, err
	}
	return &ICrossChainRegistryGenerationReservationRemovedIterator{contract: _ICrossChainRegistry.contract, event: "GenerationReservationRemoved", logs: logs, sub: sub}, nil
}

// WatchGenerationReservationRemoved is a free log subscription operation binding the contract event 0x4ffdfdd59e9e1e3c301608788f78dd458e61cb8c045ca92b62a7b484c80824fb.
//
// Solidity: event GenerationReservationRemoved((address,uint32) operatorSet)
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

// ParseGenerationReservationRemoved is a log parse operation binding the contract event 0x4ffdfdd59e9e1e3c301608788f78dd458e61cb8c045ca92b62a7b484c80824fb.
//
// Solidity: event GenerationReservationRemoved((address,uint32) operatorSet)
func (_ICrossChainRegistry *ICrossChainRegistryFilterer) ParseGenerationReservationRemoved(log types.Log) (*ICrossChainRegistryGenerationReservationRemoved, error) {
	event := new(ICrossChainRegistryGenerationReservationRemoved)
	if err := _ICrossChainRegistry.contract.UnpackLog(event, "GenerationReservationRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICrossChainRegistryOperatorSetConfigSetIterator is returned from FilterOperatorSetConfigSet and is used to iterate over the raw logs and unpacked data for OperatorSetConfigSet events raised by the ICrossChainRegistry contract.
type ICrossChainRegistryOperatorSetConfigSetIterator struct {
	Event *ICrossChainRegistryOperatorSetConfigSet // Event containing the contract specifics and raw log

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
func (it *ICrossChainRegistryOperatorSetConfigSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICrossChainRegistryOperatorSetConfigSet)
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
		it.Event = new(ICrossChainRegistryOperatorSetConfigSet)
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
func (it *ICrossChainRegistryOperatorSetConfigSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICrossChainRegistryOperatorSetConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICrossChainRegistryOperatorSetConfigSet represents a OperatorSetConfigSet event raised by the ICrossChainRegistry contract.
type ICrossChainRegistryOperatorSetConfigSet struct {
	OperatorSet OperatorSet
	Config      ICrossChainRegistryTypesOperatorSetConfig
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorSetConfigSet is a free log retrieval operation binding the contract event 0x3147846ee526009000671c20380b856a633345691300f82585f90034715cf0e2.
//
// Solidity: event OperatorSetConfigSet((address,uint32) operatorSet, (address,uint32) config)
func (_ICrossChainRegistry *ICrossChainRegistryFilterer) FilterOperatorSetConfigSet(opts *bind.FilterOpts) (*ICrossChainRegistryOperatorSetConfigSetIterator, error) {

	logs, sub, err := _ICrossChainRegistry.contract.FilterLogs(opts, "OperatorSetConfigSet")
	if err != nil {
		return nil, err
	}
	return &ICrossChainRegistryOperatorSetConfigSetIterator{contract: _ICrossChainRegistry.contract, event: "OperatorSetConfigSet", logs: logs, sub: sub}, nil
}

// WatchOperatorSetConfigSet is a free log subscription operation binding the contract event 0x3147846ee526009000671c20380b856a633345691300f82585f90034715cf0e2.
//
// Solidity: event OperatorSetConfigSet((address,uint32) operatorSet, (address,uint32) config)
func (_ICrossChainRegistry *ICrossChainRegistryFilterer) WatchOperatorSetConfigSet(opts *bind.WatchOpts, sink chan<- *ICrossChainRegistryOperatorSetConfigSet) (event.Subscription, error) {

	logs, sub, err := _ICrossChainRegistry.contract.WatchLogs(opts, "OperatorSetConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICrossChainRegistryOperatorSetConfigSet)
				if err := _ICrossChainRegistry.contract.UnpackLog(event, "OperatorSetConfigSet", log); err != nil {
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
func (_ICrossChainRegistry *ICrossChainRegistryFilterer) ParseOperatorSetConfigSet(log types.Log) (*ICrossChainRegistryOperatorSetConfigSet, error) {
	event := new(ICrossChainRegistryOperatorSetConfigSet)
	if err := _ICrossChainRegistry.contract.UnpackLog(event, "OperatorSetConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICrossChainRegistryOperatorTableCalculatorSetIterator is returned from FilterOperatorTableCalculatorSet and is used to iterate over the raw logs and unpacked data for OperatorTableCalculatorSet events raised by the ICrossChainRegistry contract.
type ICrossChainRegistryOperatorTableCalculatorSetIterator struct {
	Event *ICrossChainRegistryOperatorTableCalculatorSet // Event containing the contract specifics and raw log

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
func (it *ICrossChainRegistryOperatorTableCalculatorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICrossChainRegistryOperatorTableCalculatorSet)
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
		it.Event = new(ICrossChainRegistryOperatorTableCalculatorSet)
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
func (it *ICrossChainRegistryOperatorTableCalculatorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICrossChainRegistryOperatorTableCalculatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICrossChainRegistryOperatorTableCalculatorSet represents a OperatorTableCalculatorSet event raised by the ICrossChainRegistry contract.
type ICrossChainRegistryOperatorTableCalculatorSet struct {
	OperatorSet             OperatorSet
	OperatorTableCalculator common.Address
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterOperatorTableCalculatorSet is a free log retrieval operation binding the contract event 0x7f7ccafd92d20fdb39dee184a0dce002a9da420ed0def461f2a027abc9b3f6df.
//
// Solidity: event OperatorTableCalculatorSet((address,uint32) operatorSet, address operatorTableCalculator)
func (_ICrossChainRegistry *ICrossChainRegistryFilterer) FilterOperatorTableCalculatorSet(opts *bind.FilterOpts) (*ICrossChainRegistryOperatorTableCalculatorSetIterator, error) {

	logs, sub, err := _ICrossChainRegistry.contract.FilterLogs(opts, "OperatorTableCalculatorSet")
	if err != nil {
		return nil, err
	}
	return &ICrossChainRegistryOperatorTableCalculatorSetIterator{contract: _ICrossChainRegistry.contract, event: "OperatorTableCalculatorSet", logs: logs, sub: sub}, nil
}

// WatchOperatorTableCalculatorSet is a free log subscription operation binding the contract event 0x7f7ccafd92d20fdb39dee184a0dce002a9da420ed0def461f2a027abc9b3f6df.
//
// Solidity: event OperatorTableCalculatorSet((address,uint32) operatorSet, address operatorTableCalculator)
func (_ICrossChainRegistry *ICrossChainRegistryFilterer) WatchOperatorTableCalculatorSet(opts *bind.WatchOpts, sink chan<- *ICrossChainRegistryOperatorTableCalculatorSet) (event.Subscription, error) {

	logs, sub, err := _ICrossChainRegistry.contract.WatchLogs(opts, "OperatorTableCalculatorSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICrossChainRegistryOperatorTableCalculatorSet)
				if err := _ICrossChainRegistry.contract.UnpackLog(event, "OperatorTableCalculatorSet", log); err != nil {
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
func (_ICrossChainRegistry *ICrossChainRegistryFilterer) ParseOperatorTableCalculatorSet(log types.Log) (*ICrossChainRegistryOperatorTableCalculatorSet, error) {
	event := new(ICrossChainRegistryOperatorTableCalculatorSet)
	if err := _ICrossChainRegistry.contract.UnpackLog(event, "OperatorTableCalculatorSet", log); err != nil {
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
	ChainID     *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTransportDestinationAdded is a free log retrieval operation binding the contract event 0x8b3a5eb206599a7bd7dcffa354a014ae777667c6890b23d046ff6867cd968393.
//
// Solidity: event TransportDestinationAdded((address,uint32) operatorSet, uint256 chainID)
func (_ICrossChainRegistry *ICrossChainRegistryFilterer) FilterTransportDestinationAdded(opts *bind.FilterOpts) (*ICrossChainRegistryTransportDestinationAddedIterator, error) {

	logs, sub, err := _ICrossChainRegistry.contract.FilterLogs(opts, "TransportDestinationAdded")
	if err != nil {
		return nil, err
	}
	return &ICrossChainRegistryTransportDestinationAddedIterator{contract: _ICrossChainRegistry.contract, event: "TransportDestinationAdded", logs: logs, sub: sub}, nil
}

// WatchTransportDestinationAdded is a free log subscription operation binding the contract event 0x8b3a5eb206599a7bd7dcffa354a014ae777667c6890b23d046ff6867cd968393.
//
// Solidity: event TransportDestinationAdded((address,uint32) operatorSet, uint256 chainID)
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

// ParseTransportDestinationAdded is a log parse operation binding the contract event 0x8b3a5eb206599a7bd7dcffa354a014ae777667c6890b23d046ff6867cd968393.
//
// Solidity: event TransportDestinationAdded((address,uint32) operatorSet, uint256 chainID)
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
	ChainID     *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTransportDestinationRemoved is a free log retrieval operation binding the contract event 0x9fa5794dfaeae2ede46d5bdba347083580889221c3338813ca6c8d4b681ae8ee.
//
// Solidity: event TransportDestinationRemoved((address,uint32) operatorSet, uint256 chainID)
func (_ICrossChainRegistry *ICrossChainRegistryFilterer) FilterTransportDestinationRemoved(opts *bind.FilterOpts) (*ICrossChainRegistryTransportDestinationRemovedIterator, error) {

	logs, sub, err := _ICrossChainRegistry.contract.FilterLogs(opts, "TransportDestinationRemoved")
	if err != nil {
		return nil, err
	}
	return &ICrossChainRegistryTransportDestinationRemovedIterator{contract: _ICrossChainRegistry.contract, event: "TransportDestinationRemoved", logs: logs, sub: sub}, nil
}

// WatchTransportDestinationRemoved is a free log subscription operation binding the contract event 0x9fa5794dfaeae2ede46d5bdba347083580889221c3338813ca6c8d4b681ae8ee.
//
// Solidity: event TransportDestinationRemoved((address,uint32) operatorSet, uint256 chainID)
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

// ParseTransportDestinationRemoved is a log parse operation binding the contract event 0x9fa5794dfaeae2ede46d5bdba347083580889221c3338813ca6c8d4b681ae8ee.
//
// Solidity: event TransportDestinationRemoved((address,uint32) operatorSet, uint256 chainID)
func (_ICrossChainRegistry *ICrossChainRegistryFilterer) ParseTransportDestinationRemoved(log types.Log) (*ICrossChainRegistryTransportDestinationRemoved, error) {
	event := new(ICrossChainRegistryTransportDestinationRemoved)
	if err := _ICrossChainRegistry.contract.UnpackLog(event, "TransportDestinationRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
