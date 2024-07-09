// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package SlasherStorage

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

// IOperatorSetManagerOperatorSet is an auto generated low-level Go binding around an user-defined struct.
type IOperatorSetManagerOperatorSet struct {
	Avs common.Address
	Id  [4]byte
}

// SlasherStorageMetaData contains all meta data concerning the SlasherStorage contract.
var SlasherStorageMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSlashedRate\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structIOperatorSetManager.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"name\":\"epoch\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorSetManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIOperatorSetManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"shareScalingFactor\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"shareScalingFactorAtEpoch\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"epoch\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"slashOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"bipsToSlash\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slashingEpochHistory\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"slashingUpdates\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"slashingRate\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"scalingFactor\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"strategyManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSlashed\",\"inputs\":[{\"name\":\"epoch\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIOperatorSetManager.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"name\":\"bipsToSlash\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"slashingRate\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// SlasherStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use SlasherStorageMetaData.ABI instead.
var SlasherStorageABI = SlasherStorageMetaData.ABI

// SlasherStorage is an auto generated Go binding around an Ethereum contract.
type SlasherStorage struct {
	SlasherStorageCaller     // Read-only binding to the contract
	SlasherStorageTransactor // Write-only binding to the contract
	SlasherStorageFilterer   // Log filterer for contract events
}

// SlasherStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type SlasherStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SlasherStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SlasherStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SlasherStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SlasherStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SlasherStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SlasherStorageSession struct {
	Contract     *SlasherStorage   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SlasherStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SlasherStorageCallerSession struct {
	Contract *SlasherStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// SlasherStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SlasherStorageTransactorSession struct {
	Contract     *SlasherStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// SlasherStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type SlasherStorageRaw struct {
	Contract *SlasherStorage // Generic contract binding to access the raw methods on
}

// SlasherStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SlasherStorageCallerRaw struct {
	Contract *SlasherStorageCaller // Generic read-only contract binding to access the raw methods on
}

// SlasherStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SlasherStorageTransactorRaw struct {
	Contract *SlasherStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSlasherStorage creates a new instance of SlasherStorage, bound to a specific deployed contract.
func NewSlasherStorage(address common.Address, backend bind.ContractBackend) (*SlasherStorage, error) {
	contract, err := bindSlasherStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SlasherStorage{SlasherStorageCaller: SlasherStorageCaller{contract: contract}, SlasherStorageTransactor: SlasherStorageTransactor{contract: contract}, SlasherStorageFilterer: SlasherStorageFilterer{contract: contract}}, nil
}

// NewSlasherStorageCaller creates a new read-only instance of SlasherStorage, bound to a specific deployed contract.
func NewSlasherStorageCaller(address common.Address, caller bind.ContractCaller) (*SlasherStorageCaller, error) {
	contract, err := bindSlasherStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SlasherStorageCaller{contract: contract}, nil
}

// NewSlasherStorageTransactor creates a new write-only instance of SlasherStorage, bound to a specific deployed contract.
func NewSlasherStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*SlasherStorageTransactor, error) {
	contract, err := bindSlasherStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SlasherStorageTransactor{contract: contract}, nil
}

// NewSlasherStorageFilterer creates a new log filterer instance of SlasherStorage, bound to a specific deployed contract.
func NewSlasherStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*SlasherStorageFilterer, error) {
	contract, err := bindSlasherStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SlasherStorageFilterer{contract: contract}, nil
}

// bindSlasherStorage binds a generic wrapper to an already deployed contract.
func bindSlasherStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SlasherStorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SlasherStorage *SlasherStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SlasherStorage.Contract.SlasherStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SlasherStorage *SlasherStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SlasherStorage.Contract.SlasherStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SlasherStorage *SlasherStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SlasherStorage.Contract.SlasherStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SlasherStorage *SlasherStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SlasherStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SlasherStorage *SlasherStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SlasherStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SlasherStorage *SlasherStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SlasherStorage.Contract.contract.Transact(opts, method, params...)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_SlasherStorage *SlasherStorageCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SlasherStorage.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_SlasherStorage *SlasherStorageSession) Delegation() (common.Address, error) {
	return _SlasherStorage.Contract.Delegation(&_SlasherStorage.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_SlasherStorage *SlasherStorageCallerSession) Delegation() (common.Address, error) {
	return _SlasherStorage.Contract.Delegation(&_SlasherStorage.CallOpts)
}

// GetSlashedRate is a free data retrieval call binding the contract method 0xb59d8fdb.
//
// Solidity: function getSlashedRate(address operator, address strategy, (address,bytes4) operatorSet, uint32 epoch) view returns(uint64)
func (_SlasherStorage *SlasherStorageCaller) GetSlashedRate(opts *bind.CallOpts, operator common.Address, strategy common.Address, operatorSet IOperatorSetManagerOperatorSet, epoch uint32) (uint64, error) {
	var out []interface{}
	err := _SlasherStorage.contract.Call(opts, &out, "getSlashedRate", operator, strategy, operatorSet, epoch)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetSlashedRate is a free data retrieval call binding the contract method 0xb59d8fdb.
//
// Solidity: function getSlashedRate(address operator, address strategy, (address,bytes4) operatorSet, uint32 epoch) view returns(uint64)
func (_SlasherStorage *SlasherStorageSession) GetSlashedRate(operator common.Address, strategy common.Address, operatorSet IOperatorSetManagerOperatorSet, epoch uint32) (uint64, error) {
	return _SlasherStorage.Contract.GetSlashedRate(&_SlasherStorage.CallOpts, operator, strategy, operatorSet, epoch)
}

// GetSlashedRate is a free data retrieval call binding the contract method 0xb59d8fdb.
//
// Solidity: function getSlashedRate(address operator, address strategy, (address,bytes4) operatorSet, uint32 epoch) view returns(uint64)
func (_SlasherStorage *SlasherStorageCallerSession) GetSlashedRate(operator common.Address, strategy common.Address, operatorSet IOperatorSetManagerOperatorSet, epoch uint32) (uint64, error) {
	return _SlasherStorage.Contract.GetSlashedRate(&_SlasherStorage.CallOpts, operator, strategy, operatorSet, epoch)
}

// OperatorSetManager is a free data retrieval call binding the contract method 0xc78d4bcd.
//
// Solidity: function operatorSetManager() view returns(address)
func (_SlasherStorage *SlasherStorageCaller) OperatorSetManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SlasherStorage.contract.Call(opts, &out, "operatorSetManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OperatorSetManager is a free data retrieval call binding the contract method 0xc78d4bcd.
//
// Solidity: function operatorSetManager() view returns(address)
func (_SlasherStorage *SlasherStorageSession) OperatorSetManager() (common.Address, error) {
	return _SlasherStorage.Contract.OperatorSetManager(&_SlasherStorage.CallOpts)
}

// OperatorSetManager is a free data retrieval call binding the contract method 0xc78d4bcd.
//
// Solidity: function operatorSetManager() view returns(address)
func (_SlasherStorage *SlasherStorageCallerSession) OperatorSetManager() (common.Address, error) {
	return _SlasherStorage.Contract.OperatorSetManager(&_SlasherStorage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SlasherStorage *SlasherStorageCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SlasherStorage.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SlasherStorage *SlasherStorageSession) Owner() (common.Address, error) {
	return _SlasherStorage.Contract.Owner(&_SlasherStorage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SlasherStorage *SlasherStorageCallerSession) Owner() (common.Address, error) {
	return _SlasherStorage.Contract.Owner(&_SlasherStorage.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_SlasherStorage *SlasherStorageCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _SlasherStorage.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_SlasherStorage *SlasherStorageSession) Paused(index uint8) (bool, error) {
	return _SlasherStorage.Contract.Paused(&_SlasherStorage.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_SlasherStorage *SlasherStorageCallerSession) Paused(index uint8) (bool, error) {
	return _SlasherStorage.Contract.Paused(&_SlasherStorage.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_SlasherStorage *SlasherStorageCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SlasherStorage.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_SlasherStorage *SlasherStorageSession) Paused0() (*big.Int, error) {
	return _SlasherStorage.Contract.Paused0(&_SlasherStorage.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_SlasherStorage *SlasherStorageCallerSession) Paused0() (*big.Int, error) {
	return _SlasherStorage.Contract.Paused0(&_SlasherStorage.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_SlasherStorage *SlasherStorageCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SlasherStorage.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_SlasherStorage *SlasherStorageSession) PauserRegistry() (common.Address, error) {
	return _SlasherStorage.Contract.PauserRegistry(&_SlasherStorage.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_SlasherStorage *SlasherStorageCallerSession) PauserRegistry() (common.Address, error) {
	return _SlasherStorage.Contract.PauserRegistry(&_SlasherStorage.CallOpts)
}

// ShareScalingFactor is a free data retrieval call binding the contract method 0x334f00d6.
//
// Solidity: function shareScalingFactor(address operator, address strategy) view returns(uint64)
func (_SlasherStorage *SlasherStorageCaller) ShareScalingFactor(opts *bind.CallOpts, operator common.Address, strategy common.Address) (uint64, error) {
	var out []interface{}
	err := _SlasherStorage.contract.Call(opts, &out, "shareScalingFactor", operator, strategy)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ShareScalingFactor is a free data retrieval call binding the contract method 0x334f00d6.
//
// Solidity: function shareScalingFactor(address operator, address strategy) view returns(uint64)
func (_SlasherStorage *SlasherStorageSession) ShareScalingFactor(operator common.Address, strategy common.Address) (uint64, error) {
	return _SlasherStorage.Contract.ShareScalingFactor(&_SlasherStorage.CallOpts, operator, strategy)
}

// ShareScalingFactor is a free data retrieval call binding the contract method 0x334f00d6.
//
// Solidity: function shareScalingFactor(address operator, address strategy) view returns(uint64)
func (_SlasherStorage *SlasherStorageCallerSession) ShareScalingFactor(operator common.Address, strategy common.Address) (uint64, error) {
	return _SlasherStorage.Contract.ShareScalingFactor(&_SlasherStorage.CallOpts, operator, strategy)
}

// ShareScalingFactorAtEpoch is a free data retrieval call binding the contract method 0xe49a1e84.
//
// Solidity: function shareScalingFactorAtEpoch(address operator, address strategy, uint32 epoch) view returns(uint64)
func (_SlasherStorage *SlasherStorageCaller) ShareScalingFactorAtEpoch(opts *bind.CallOpts, operator common.Address, strategy common.Address, epoch uint32) (uint64, error) {
	var out []interface{}
	err := _SlasherStorage.contract.Call(opts, &out, "shareScalingFactorAtEpoch", operator, strategy, epoch)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ShareScalingFactorAtEpoch is a free data retrieval call binding the contract method 0xe49a1e84.
//
// Solidity: function shareScalingFactorAtEpoch(address operator, address strategy, uint32 epoch) view returns(uint64)
func (_SlasherStorage *SlasherStorageSession) ShareScalingFactorAtEpoch(operator common.Address, strategy common.Address, epoch uint32) (uint64, error) {
	return _SlasherStorage.Contract.ShareScalingFactorAtEpoch(&_SlasherStorage.CallOpts, operator, strategy, epoch)
}

// ShareScalingFactorAtEpoch is a free data retrieval call binding the contract method 0xe49a1e84.
//
// Solidity: function shareScalingFactorAtEpoch(address operator, address strategy, uint32 epoch) view returns(uint64)
func (_SlasherStorage *SlasherStorageCallerSession) ShareScalingFactorAtEpoch(operator common.Address, strategy common.Address, epoch uint32) (uint64, error) {
	return _SlasherStorage.Contract.ShareScalingFactorAtEpoch(&_SlasherStorage.CallOpts, operator, strategy, epoch)
}

// SlashingEpochHistory is a free data retrieval call binding the contract method 0x4279a7e6.
//
// Solidity: function slashingEpochHistory(address , address , uint256 ) view returns(uint32)
func (_SlasherStorage *SlasherStorageCaller) SlashingEpochHistory(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) (uint32, error) {
	var out []interface{}
	err := _SlasherStorage.contract.Call(opts, &out, "slashingEpochHistory", arg0, arg1, arg2)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// SlashingEpochHistory is a free data retrieval call binding the contract method 0x4279a7e6.
//
// Solidity: function slashingEpochHistory(address , address , uint256 ) view returns(uint32)
func (_SlasherStorage *SlasherStorageSession) SlashingEpochHistory(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (uint32, error) {
	return _SlasherStorage.Contract.SlashingEpochHistory(&_SlasherStorage.CallOpts, arg0, arg1, arg2)
}

// SlashingEpochHistory is a free data retrieval call binding the contract method 0x4279a7e6.
//
// Solidity: function slashingEpochHistory(address , address , uint256 ) view returns(uint32)
func (_SlasherStorage *SlasherStorageCallerSession) SlashingEpochHistory(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (uint32, error) {
	return _SlasherStorage.Contract.SlashingEpochHistory(&_SlasherStorage.CallOpts, arg0, arg1, arg2)
}

// SlashingUpdates is a free data retrieval call binding the contract method 0x0b1b781e.
//
// Solidity: function slashingUpdates(address , address , uint32 ) view returns(uint64 slashingRate, uint64 scalingFactor)
func (_SlasherStorage *SlasherStorageCaller) SlashingUpdates(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 uint32) (struct {
	SlashingRate  uint64
	ScalingFactor uint64
}, error) {
	var out []interface{}
	err := _SlasherStorage.contract.Call(opts, &out, "slashingUpdates", arg0, arg1, arg2)

	outstruct := new(struct {
		SlashingRate  uint64
		ScalingFactor uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SlashingRate = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.ScalingFactor = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// SlashingUpdates is a free data retrieval call binding the contract method 0x0b1b781e.
//
// Solidity: function slashingUpdates(address , address , uint32 ) view returns(uint64 slashingRate, uint64 scalingFactor)
func (_SlasherStorage *SlasherStorageSession) SlashingUpdates(arg0 common.Address, arg1 common.Address, arg2 uint32) (struct {
	SlashingRate  uint64
	ScalingFactor uint64
}, error) {
	return _SlasherStorage.Contract.SlashingUpdates(&_SlasherStorage.CallOpts, arg0, arg1, arg2)
}

// SlashingUpdates is a free data retrieval call binding the contract method 0x0b1b781e.
//
// Solidity: function slashingUpdates(address , address , uint32 ) view returns(uint64 slashingRate, uint64 scalingFactor)
func (_SlasherStorage *SlasherStorageCallerSession) SlashingUpdates(arg0 common.Address, arg1 common.Address, arg2 uint32) (struct {
	SlashingRate  uint64
	ScalingFactor uint64
}, error) {
	return _SlasherStorage.Contract.SlashingUpdates(&_SlasherStorage.CallOpts, arg0, arg1, arg2)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_SlasherStorage *SlasherStorageCaller) StrategyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SlasherStorage.contract.Call(opts, &out, "strategyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_SlasherStorage *SlasherStorageSession) StrategyManager() (common.Address, error) {
	return _SlasherStorage.Contract.StrategyManager(&_SlasherStorage.CallOpts)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_SlasherStorage *SlasherStorageCallerSession) StrategyManager() (common.Address, error) {
	return _SlasherStorage.Contract.StrategyManager(&_SlasherStorage.CallOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_SlasherStorage *SlasherStorageTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _SlasherStorage.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_SlasherStorage *SlasherStorageSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _SlasherStorage.Contract.Pause(&_SlasherStorage.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_SlasherStorage *SlasherStorageTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _SlasherStorage.Contract.Pause(&_SlasherStorage.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_SlasherStorage *SlasherStorageTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SlasherStorage.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_SlasherStorage *SlasherStorageSession) PauseAll() (*types.Transaction, error) {
	return _SlasherStorage.Contract.PauseAll(&_SlasherStorage.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_SlasherStorage *SlasherStorageTransactorSession) PauseAll() (*types.Transaction, error) {
	return _SlasherStorage.Contract.PauseAll(&_SlasherStorage.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SlasherStorage *SlasherStorageTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SlasherStorage.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SlasherStorage *SlasherStorageSession) RenounceOwnership() (*types.Transaction, error) {
	return _SlasherStorage.Contract.RenounceOwnership(&_SlasherStorage.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SlasherStorage *SlasherStorageTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SlasherStorage.Contract.RenounceOwnership(&_SlasherStorage.TransactOpts)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_SlasherStorage *SlasherStorageTransactor) SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error) {
	return _SlasherStorage.contract.Transact(opts, "setPauserRegistry", newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_SlasherStorage *SlasherStorageSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _SlasherStorage.Contract.SetPauserRegistry(&_SlasherStorage.TransactOpts, newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_SlasherStorage *SlasherStorageTransactorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _SlasherStorage.Contract.SetPauserRegistry(&_SlasherStorage.TransactOpts, newPauserRegistry)
}

// SlashOperator is a paid mutator transaction binding the contract method 0x4a1def9a.
//
// Solidity: function slashOperator(address operator, bytes4 operatorSetId, address[] strategies, uint32 bipsToSlash) returns()
func (_SlasherStorage *SlasherStorageTransactor) SlashOperator(opts *bind.TransactOpts, operator common.Address, operatorSetId [4]byte, strategies []common.Address, bipsToSlash uint32) (*types.Transaction, error) {
	return _SlasherStorage.contract.Transact(opts, "slashOperator", operator, operatorSetId, strategies, bipsToSlash)
}

// SlashOperator is a paid mutator transaction binding the contract method 0x4a1def9a.
//
// Solidity: function slashOperator(address operator, bytes4 operatorSetId, address[] strategies, uint32 bipsToSlash) returns()
func (_SlasherStorage *SlasherStorageSession) SlashOperator(operator common.Address, operatorSetId [4]byte, strategies []common.Address, bipsToSlash uint32) (*types.Transaction, error) {
	return _SlasherStorage.Contract.SlashOperator(&_SlasherStorage.TransactOpts, operator, operatorSetId, strategies, bipsToSlash)
}

// SlashOperator is a paid mutator transaction binding the contract method 0x4a1def9a.
//
// Solidity: function slashOperator(address operator, bytes4 operatorSetId, address[] strategies, uint32 bipsToSlash) returns()
func (_SlasherStorage *SlasherStorageTransactorSession) SlashOperator(operator common.Address, operatorSetId [4]byte, strategies []common.Address, bipsToSlash uint32) (*types.Transaction, error) {
	return _SlasherStorage.Contract.SlashOperator(&_SlasherStorage.TransactOpts, operator, operatorSetId, strategies, bipsToSlash)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SlasherStorage *SlasherStorageTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SlasherStorage.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SlasherStorage *SlasherStorageSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SlasherStorage.Contract.TransferOwnership(&_SlasherStorage.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SlasherStorage *SlasherStorageTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SlasherStorage.Contract.TransferOwnership(&_SlasherStorage.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_SlasherStorage *SlasherStorageTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _SlasherStorage.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_SlasherStorage *SlasherStorageSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _SlasherStorage.Contract.Unpause(&_SlasherStorage.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_SlasherStorage *SlasherStorageTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _SlasherStorage.Contract.Unpause(&_SlasherStorage.TransactOpts, newPausedStatus)
}

// SlasherStorageInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SlasherStorage contract.
type SlasherStorageInitializedIterator struct {
	Event *SlasherStorageInitialized // Event containing the contract specifics and raw log

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
func (it *SlasherStorageInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlasherStorageInitialized)
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
		it.Event = new(SlasherStorageInitialized)
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
func (it *SlasherStorageInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlasherStorageInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlasherStorageInitialized represents a Initialized event raised by the SlasherStorage contract.
type SlasherStorageInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SlasherStorage *SlasherStorageFilterer) FilterInitialized(opts *bind.FilterOpts) (*SlasherStorageInitializedIterator, error) {

	logs, sub, err := _SlasherStorage.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SlasherStorageInitializedIterator{contract: _SlasherStorage.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SlasherStorage *SlasherStorageFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SlasherStorageInitialized) (event.Subscription, error) {

	logs, sub, err := _SlasherStorage.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlasherStorageInitialized)
				if err := _SlasherStorage.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SlasherStorage *SlasherStorageFilterer) ParseInitialized(log types.Log) (*SlasherStorageInitialized, error) {
	event := new(SlasherStorageInitialized)
	if err := _SlasherStorage.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlasherStorageOperatorSlashedIterator is returned from FilterOperatorSlashed and is used to iterate over the raw logs and unpacked data for OperatorSlashed events raised by the SlasherStorage contract.
type SlasherStorageOperatorSlashedIterator struct {
	Event *SlasherStorageOperatorSlashed // Event containing the contract specifics and raw log

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
func (it *SlasherStorageOperatorSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlasherStorageOperatorSlashed)
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
		it.Event = new(SlasherStorageOperatorSlashed)
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
func (it *SlasherStorageOperatorSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlasherStorageOperatorSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlasherStorageOperatorSlashed represents a OperatorSlashed event raised by the SlasherStorage contract.
type SlasherStorageOperatorSlashed struct {
	Epoch        uint32
	Operator     common.Address
	Strategy     common.Address
	OperatorSet  IOperatorSetManagerOperatorSet
	BipsToSlash  uint32
	SlashingRate uint64
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOperatorSlashed is a free log retrieval operation binding the contract event 0x471fe23f2a18902ad4f5859f431c6cc59256d682c861ee3405719f2faa09f937.
//
// Solidity: event OperatorSlashed(uint32 epoch, address operator, address strategy, (address,bytes4) operatorSet, uint32 bipsToSlash, uint64 slashingRate)
func (_SlasherStorage *SlasherStorageFilterer) FilterOperatorSlashed(opts *bind.FilterOpts) (*SlasherStorageOperatorSlashedIterator, error) {

	logs, sub, err := _SlasherStorage.contract.FilterLogs(opts, "OperatorSlashed")
	if err != nil {
		return nil, err
	}
	return &SlasherStorageOperatorSlashedIterator{contract: _SlasherStorage.contract, event: "OperatorSlashed", logs: logs, sub: sub}, nil
}

// WatchOperatorSlashed is a free log subscription operation binding the contract event 0x471fe23f2a18902ad4f5859f431c6cc59256d682c861ee3405719f2faa09f937.
//
// Solidity: event OperatorSlashed(uint32 epoch, address operator, address strategy, (address,bytes4) operatorSet, uint32 bipsToSlash, uint64 slashingRate)
func (_SlasherStorage *SlasherStorageFilterer) WatchOperatorSlashed(opts *bind.WatchOpts, sink chan<- *SlasherStorageOperatorSlashed) (event.Subscription, error) {

	logs, sub, err := _SlasherStorage.contract.WatchLogs(opts, "OperatorSlashed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlasherStorageOperatorSlashed)
				if err := _SlasherStorage.contract.UnpackLog(event, "OperatorSlashed", log); err != nil {
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

// ParseOperatorSlashed is a log parse operation binding the contract event 0x471fe23f2a18902ad4f5859f431c6cc59256d682c861ee3405719f2faa09f937.
//
// Solidity: event OperatorSlashed(uint32 epoch, address operator, address strategy, (address,bytes4) operatorSet, uint32 bipsToSlash, uint64 slashingRate)
func (_SlasherStorage *SlasherStorageFilterer) ParseOperatorSlashed(log types.Log) (*SlasherStorageOperatorSlashed, error) {
	event := new(SlasherStorageOperatorSlashed)
	if err := _SlasherStorage.contract.UnpackLog(event, "OperatorSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlasherStorageOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SlasherStorage contract.
type SlasherStorageOwnershipTransferredIterator struct {
	Event *SlasherStorageOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SlasherStorageOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlasherStorageOwnershipTransferred)
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
		it.Event = new(SlasherStorageOwnershipTransferred)
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
func (it *SlasherStorageOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlasherStorageOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlasherStorageOwnershipTransferred represents a OwnershipTransferred event raised by the SlasherStorage contract.
type SlasherStorageOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SlasherStorage *SlasherStorageFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SlasherStorageOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SlasherStorage.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SlasherStorageOwnershipTransferredIterator{contract: _SlasherStorage.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SlasherStorage *SlasherStorageFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SlasherStorageOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SlasherStorage.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlasherStorageOwnershipTransferred)
				if err := _SlasherStorage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SlasherStorage *SlasherStorageFilterer) ParseOwnershipTransferred(log types.Log) (*SlasherStorageOwnershipTransferred, error) {
	event := new(SlasherStorageOwnershipTransferred)
	if err := _SlasherStorage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlasherStoragePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the SlasherStorage contract.
type SlasherStoragePausedIterator struct {
	Event *SlasherStoragePaused // Event containing the contract specifics and raw log

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
func (it *SlasherStoragePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlasherStoragePaused)
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
		it.Event = new(SlasherStoragePaused)
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
func (it *SlasherStoragePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlasherStoragePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlasherStoragePaused represents a Paused event raised by the SlasherStorage contract.
type SlasherStoragePaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_SlasherStorage *SlasherStorageFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*SlasherStoragePausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SlasherStorage.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &SlasherStoragePausedIterator{contract: _SlasherStorage.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_SlasherStorage *SlasherStorageFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *SlasherStoragePaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SlasherStorage.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlasherStoragePaused)
				if err := _SlasherStorage.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_SlasherStorage *SlasherStorageFilterer) ParsePaused(log types.Log) (*SlasherStoragePaused, error) {
	event := new(SlasherStoragePaused)
	if err := _SlasherStorage.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlasherStoragePauserRegistrySetIterator is returned from FilterPauserRegistrySet and is used to iterate over the raw logs and unpacked data for PauserRegistrySet events raised by the SlasherStorage contract.
type SlasherStoragePauserRegistrySetIterator struct {
	Event *SlasherStoragePauserRegistrySet // Event containing the contract specifics and raw log

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
func (it *SlasherStoragePauserRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlasherStoragePauserRegistrySet)
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
		it.Event = new(SlasherStoragePauserRegistrySet)
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
func (it *SlasherStoragePauserRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlasherStoragePauserRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlasherStoragePauserRegistrySet represents a PauserRegistrySet event raised by the SlasherStorage contract.
type SlasherStoragePauserRegistrySet struct {
	PauserRegistry    common.Address
	NewPauserRegistry common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPauserRegistrySet is a free log retrieval operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_SlasherStorage *SlasherStorageFilterer) FilterPauserRegistrySet(opts *bind.FilterOpts) (*SlasherStoragePauserRegistrySetIterator, error) {

	logs, sub, err := _SlasherStorage.contract.FilterLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return &SlasherStoragePauserRegistrySetIterator{contract: _SlasherStorage.contract, event: "PauserRegistrySet", logs: logs, sub: sub}, nil
}

// WatchPauserRegistrySet is a free log subscription operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_SlasherStorage *SlasherStorageFilterer) WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *SlasherStoragePauserRegistrySet) (event.Subscription, error) {

	logs, sub, err := _SlasherStorage.contract.WatchLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlasherStoragePauserRegistrySet)
				if err := _SlasherStorage.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
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

// ParsePauserRegistrySet is a log parse operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_SlasherStorage *SlasherStorageFilterer) ParsePauserRegistrySet(log types.Log) (*SlasherStoragePauserRegistrySet, error) {
	event := new(SlasherStoragePauserRegistrySet)
	if err := _SlasherStorage.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlasherStorageUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the SlasherStorage contract.
type SlasherStorageUnpausedIterator struct {
	Event *SlasherStorageUnpaused // Event containing the contract specifics and raw log

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
func (it *SlasherStorageUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlasherStorageUnpaused)
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
		it.Event = new(SlasherStorageUnpaused)
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
func (it *SlasherStorageUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlasherStorageUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlasherStorageUnpaused represents a Unpaused event raised by the SlasherStorage contract.
type SlasherStorageUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_SlasherStorage *SlasherStorageFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*SlasherStorageUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SlasherStorage.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &SlasherStorageUnpausedIterator{contract: _SlasherStorage.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_SlasherStorage *SlasherStorageFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *SlasherStorageUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SlasherStorage.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlasherStorageUnpaused)
				if err := _SlasherStorage.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_SlasherStorage *SlasherStorageFilterer) ParseUnpaused(log types.Log) (*SlasherStorageUnpaused, error) {
	event := new(SlasherStorageUnpaused)
	if err := _SlasherStorage.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
