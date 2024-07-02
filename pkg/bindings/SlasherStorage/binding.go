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
	ABI: "[{\"type\":\"function\",\"name\":\"canWithdraw\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"epoch\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"executeSlashing\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"epoch\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getPendingSlashingRate\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structIOperatorSetManager.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"name\":\"epoch\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequestedSlashingRate\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structIOperatorSetManager.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"name\":\"epoch\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalPendingSlashingRate\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"epoch\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWithdrawabilityAndScalingFactorAtEpoch\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"epoch\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"increaseRequestedBipsToSlash\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetID\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"bipsToIncrease\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"operatorSetManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIOperatorSetManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingShareScalingFactor\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"reduceRequestedBipsToSlash\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetID\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"epoch\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"bipsToReduce\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requestedSlashedBips\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"shareScalingFactor\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"shareScalingFactorAtEpoch\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"epoch\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"slashedEpochHistory\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"slashingRequestIds\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"lastCreatedSlashingRequestId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"lastExecutedSlashingRequestId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"slashingRequests\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"slashingRate\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"scalingFactor\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"strategyManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RequestedBipsToSlashModified\",\"inputs\":[{\"name\":\"epoch\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIOperatorSetManager.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"name\":\"strategies\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"contractIStrategy[]\"},{\"name\":\"bipsToModify\",\"type\":\"int32\",\"indexed\":false,\"internalType\":\"int32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SlashingExecuted\",\"inputs\":[{\"name\":\"epoch\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"slashingRate\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
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

// CanWithdraw is a free data retrieval call binding the contract method 0x79c415ec.
//
// Solidity: function canWithdraw(address operator, address strategy, uint32 epoch) view returns(bool)
func (_SlasherStorage *SlasherStorageCaller) CanWithdraw(opts *bind.CallOpts, operator common.Address, strategy common.Address, epoch uint32) (bool, error) {
	var out []interface{}
	err := _SlasherStorage.contract.Call(opts, &out, "canWithdraw", operator, strategy, epoch)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanWithdraw is a free data retrieval call binding the contract method 0x79c415ec.
//
// Solidity: function canWithdraw(address operator, address strategy, uint32 epoch) view returns(bool)
func (_SlasherStorage *SlasherStorageSession) CanWithdraw(operator common.Address, strategy common.Address, epoch uint32) (bool, error) {
	return _SlasherStorage.Contract.CanWithdraw(&_SlasherStorage.CallOpts, operator, strategy, epoch)
}

// CanWithdraw is a free data retrieval call binding the contract method 0x79c415ec.
//
// Solidity: function canWithdraw(address operator, address strategy, uint32 epoch) view returns(bool)
func (_SlasherStorage *SlasherStorageCallerSession) CanWithdraw(operator common.Address, strategy common.Address, epoch uint32) (bool, error) {
	return _SlasherStorage.Contract.CanWithdraw(&_SlasherStorage.CallOpts, operator, strategy, epoch)
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

// GetPendingSlashingRate is a free data retrieval call binding the contract method 0x4d54dc3c.
//
// Solidity: function getPendingSlashingRate(address operator, address strategy, (address,bytes4) operatorSet, uint32 epoch) view returns(uint32)
func (_SlasherStorage *SlasherStorageCaller) GetPendingSlashingRate(opts *bind.CallOpts, operator common.Address, strategy common.Address, operatorSet IOperatorSetManagerOperatorSet, epoch uint32) (uint32, error) {
	var out []interface{}
	err := _SlasherStorage.contract.Call(opts, &out, "getPendingSlashingRate", operator, strategy, operatorSet, epoch)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetPendingSlashingRate is a free data retrieval call binding the contract method 0x4d54dc3c.
//
// Solidity: function getPendingSlashingRate(address operator, address strategy, (address,bytes4) operatorSet, uint32 epoch) view returns(uint32)
func (_SlasherStorage *SlasherStorageSession) GetPendingSlashingRate(operator common.Address, strategy common.Address, operatorSet IOperatorSetManagerOperatorSet, epoch uint32) (uint32, error) {
	return _SlasherStorage.Contract.GetPendingSlashingRate(&_SlasherStorage.CallOpts, operator, strategy, operatorSet, epoch)
}

// GetPendingSlashingRate is a free data retrieval call binding the contract method 0x4d54dc3c.
//
// Solidity: function getPendingSlashingRate(address operator, address strategy, (address,bytes4) operatorSet, uint32 epoch) view returns(uint32)
func (_SlasherStorage *SlasherStorageCallerSession) GetPendingSlashingRate(operator common.Address, strategy common.Address, operatorSet IOperatorSetManagerOperatorSet, epoch uint32) (uint32, error) {
	return _SlasherStorage.Contract.GetPendingSlashingRate(&_SlasherStorage.CallOpts, operator, strategy, operatorSet, epoch)
}

// GetRequestedSlashingRate is a free data retrieval call binding the contract method 0x2421a64c.
//
// Solidity: function getRequestedSlashingRate(address operator, address strategy, (address,bytes4) operatorSet, uint32 epoch) view returns(uint32)
func (_SlasherStorage *SlasherStorageCaller) GetRequestedSlashingRate(opts *bind.CallOpts, operator common.Address, strategy common.Address, operatorSet IOperatorSetManagerOperatorSet, epoch uint32) (uint32, error) {
	var out []interface{}
	err := _SlasherStorage.contract.Call(opts, &out, "getRequestedSlashingRate", operator, strategy, operatorSet, epoch)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetRequestedSlashingRate is a free data retrieval call binding the contract method 0x2421a64c.
//
// Solidity: function getRequestedSlashingRate(address operator, address strategy, (address,bytes4) operatorSet, uint32 epoch) view returns(uint32)
func (_SlasherStorage *SlasherStorageSession) GetRequestedSlashingRate(operator common.Address, strategy common.Address, operatorSet IOperatorSetManagerOperatorSet, epoch uint32) (uint32, error) {
	return _SlasherStorage.Contract.GetRequestedSlashingRate(&_SlasherStorage.CallOpts, operator, strategy, operatorSet, epoch)
}

// GetRequestedSlashingRate is a free data retrieval call binding the contract method 0x2421a64c.
//
// Solidity: function getRequestedSlashingRate(address operator, address strategy, (address,bytes4) operatorSet, uint32 epoch) view returns(uint32)
func (_SlasherStorage *SlasherStorageCallerSession) GetRequestedSlashingRate(operator common.Address, strategy common.Address, operatorSet IOperatorSetManagerOperatorSet, epoch uint32) (uint32, error) {
	return _SlasherStorage.Contract.GetRequestedSlashingRate(&_SlasherStorage.CallOpts, operator, strategy, operatorSet, epoch)
}

// GetTotalPendingSlashingRate is a free data retrieval call binding the contract method 0x90e7cde1.
//
// Solidity: function getTotalPendingSlashingRate(address operator, address strategy, uint32 epoch) view returns(uint32)
func (_SlasherStorage *SlasherStorageCaller) GetTotalPendingSlashingRate(opts *bind.CallOpts, operator common.Address, strategy common.Address, epoch uint32) (uint32, error) {
	var out []interface{}
	err := _SlasherStorage.contract.Call(opts, &out, "getTotalPendingSlashingRate", operator, strategy, epoch)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetTotalPendingSlashingRate is a free data retrieval call binding the contract method 0x90e7cde1.
//
// Solidity: function getTotalPendingSlashingRate(address operator, address strategy, uint32 epoch) view returns(uint32)
func (_SlasherStorage *SlasherStorageSession) GetTotalPendingSlashingRate(operator common.Address, strategy common.Address, epoch uint32) (uint32, error) {
	return _SlasherStorage.Contract.GetTotalPendingSlashingRate(&_SlasherStorage.CallOpts, operator, strategy, epoch)
}

// GetTotalPendingSlashingRate is a free data retrieval call binding the contract method 0x90e7cde1.
//
// Solidity: function getTotalPendingSlashingRate(address operator, address strategy, uint32 epoch) view returns(uint32)
func (_SlasherStorage *SlasherStorageCallerSession) GetTotalPendingSlashingRate(operator common.Address, strategy common.Address, epoch uint32) (uint32, error) {
	return _SlasherStorage.Contract.GetTotalPendingSlashingRate(&_SlasherStorage.CallOpts, operator, strategy, epoch)
}

// GetWithdrawabilityAndScalingFactorAtEpoch is a free data retrieval call binding the contract method 0x3be2073b.
//
// Solidity: function getWithdrawabilityAndScalingFactorAtEpoch(address operator, address strategy, uint32 epoch) view returns(bool, uint64)
func (_SlasherStorage *SlasherStorageCaller) GetWithdrawabilityAndScalingFactorAtEpoch(opts *bind.CallOpts, operator common.Address, strategy common.Address, epoch uint32) (bool, uint64, error) {
	var out []interface{}
	err := _SlasherStorage.contract.Call(opts, &out, "getWithdrawabilityAndScalingFactorAtEpoch", operator, strategy, epoch)

	if err != nil {
		return *new(bool), *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return out0, out1, err

}

// GetWithdrawabilityAndScalingFactorAtEpoch is a free data retrieval call binding the contract method 0x3be2073b.
//
// Solidity: function getWithdrawabilityAndScalingFactorAtEpoch(address operator, address strategy, uint32 epoch) view returns(bool, uint64)
func (_SlasherStorage *SlasherStorageSession) GetWithdrawabilityAndScalingFactorAtEpoch(operator common.Address, strategy common.Address, epoch uint32) (bool, uint64, error) {
	return _SlasherStorage.Contract.GetWithdrawabilityAndScalingFactorAtEpoch(&_SlasherStorage.CallOpts, operator, strategy, epoch)
}

// GetWithdrawabilityAndScalingFactorAtEpoch is a free data retrieval call binding the contract method 0x3be2073b.
//
// Solidity: function getWithdrawabilityAndScalingFactorAtEpoch(address operator, address strategy, uint32 epoch) view returns(bool, uint64)
func (_SlasherStorage *SlasherStorageCallerSession) GetWithdrawabilityAndScalingFactorAtEpoch(operator common.Address, strategy common.Address, epoch uint32) (bool, uint64, error) {
	return _SlasherStorage.Contract.GetWithdrawabilityAndScalingFactorAtEpoch(&_SlasherStorage.CallOpts, operator, strategy, epoch)
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

// PendingShareScalingFactor is a free data retrieval call binding the contract method 0x3dd9e7c5.
//
// Solidity: function pendingShareScalingFactor(address operator, address strategy) view returns(uint64)
func (_SlasherStorage *SlasherStorageCaller) PendingShareScalingFactor(opts *bind.CallOpts, operator common.Address, strategy common.Address) (uint64, error) {
	var out []interface{}
	err := _SlasherStorage.contract.Call(opts, &out, "pendingShareScalingFactor", operator, strategy)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// PendingShareScalingFactor is a free data retrieval call binding the contract method 0x3dd9e7c5.
//
// Solidity: function pendingShareScalingFactor(address operator, address strategy) view returns(uint64)
func (_SlasherStorage *SlasherStorageSession) PendingShareScalingFactor(operator common.Address, strategy common.Address) (uint64, error) {
	return _SlasherStorage.Contract.PendingShareScalingFactor(&_SlasherStorage.CallOpts, operator, strategy)
}

// PendingShareScalingFactor is a free data retrieval call binding the contract method 0x3dd9e7c5.
//
// Solidity: function pendingShareScalingFactor(address operator, address strategy) view returns(uint64)
func (_SlasherStorage *SlasherStorageCallerSession) PendingShareScalingFactor(operator common.Address, strategy common.Address) (uint64, error) {
	return _SlasherStorage.Contract.PendingShareScalingFactor(&_SlasherStorage.CallOpts, operator, strategy)
}

// RequestedSlashedBips is a free data retrieval call binding the contract method 0xec65b53d.
//
// Solidity: function requestedSlashedBips(address , address , uint32 , bytes32 ) view returns(uint32)
func (_SlasherStorage *SlasherStorageCaller) RequestedSlashedBips(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 uint32, arg3 [32]byte) (uint32, error) {
	var out []interface{}
	err := _SlasherStorage.contract.Call(opts, &out, "requestedSlashedBips", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// RequestedSlashedBips is a free data retrieval call binding the contract method 0xec65b53d.
//
// Solidity: function requestedSlashedBips(address , address , uint32 , bytes32 ) view returns(uint32)
func (_SlasherStorage *SlasherStorageSession) RequestedSlashedBips(arg0 common.Address, arg1 common.Address, arg2 uint32, arg3 [32]byte) (uint32, error) {
	return _SlasherStorage.Contract.RequestedSlashedBips(&_SlasherStorage.CallOpts, arg0, arg1, arg2, arg3)
}

// RequestedSlashedBips is a free data retrieval call binding the contract method 0xec65b53d.
//
// Solidity: function requestedSlashedBips(address , address , uint32 , bytes32 ) view returns(uint32)
func (_SlasherStorage *SlasherStorageCallerSession) RequestedSlashedBips(arg0 common.Address, arg1 common.Address, arg2 uint32, arg3 [32]byte) (uint32, error) {
	return _SlasherStorage.Contract.RequestedSlashedBips(&_SlasherStorage.CallOpts, arg0, arg1, arg2, arg3)
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

// SlashedEpochHistory is a free data retrieval call binding the contract method 0x6c0d75d0.
//
// Solidity: function slashedEpochHistory(address , address , uint256 ) view returns(uint32)
func (_SlasherStorage *SlasherStorageCaller) SlashedEpochHistory(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) (uint32, error) {
	var out []interface{}
	err := _SlasherStorage.contract.Call(opts, &out, "slashedEpochHistory", arg0, arg1, arg2)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// SlashedEpochHistory is a free data retrieval call binding the contract method 0x6c0d75d0.
//
// Solidity: function slashedEpochHistory(address , address , uint256 ) view returns(uint32)
func (_SlasherStorage *SlasherStorageSession) SlashedEpochHistory(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (uint32, error) {
	return _SlasherStorage.Contract.SlashedEpochHistory(&_SlasherStorage.CallOpts, arg0, arg1, arg2)
}

// SlashedEpochHistory is a free data retrieval call binding the contract method 0x6c0d75d0.
//
// Solidity: function slashedEpochHistory(address , address , uint256 ) view returns(uint32)
func (_SlasherStorage *SlasherStorageCallerSession) SlashedEpochHistory(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (uint32, error) {
	return _SlasherStorage.Contract.SlashedEpochHistory(&_SlasherStorage.CallOpts, arg0, arg1, arg2)
}

// SlashingRequestIds is a free data retrieval call binding the contract method 0x7ef639a6.
//
// Solidity: function slashingRequestIds(address , address ) view returns(uint32 lastCreatedSlashingRequestId, uint32 lastExecutedSlashingRequestId)
func (_SlasherStorage *SlasherStorageCaller) SlashingRequestIds(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (struct {
	LastCreatedSlashingRequestId  uint32
	LastExecutedSlashingRequestId uint32
}, error) {
	var out []interface{}
	err := _SlasherStorage.contract.Call(opts, &out, "slashingRequestIds", arg0, arg1)

	outstruct := new(struct {
		LastCreatedSlashingRequestId  uint32
		LastExecutedSlashingRequestId uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LastCreatedSlashingRequestId = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.LastExecutedSlashingRequestId = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

// SlashingRequestIds is a free data retrieval call binding the contract method 0x7ef639a6.
//
// Solidity: function slashingRequestIds(address , address ) view returns(uint32 lastCreatedSlashingRequestId, uint32 lastExecutedSlashingRequestId)
func (_SlasherStorage *SlasherStorageSession) SlashingRequestIds(arg0 common.Address, arg1 common.Address) (struct {
	LastCreatedSlashingRequestId  uint32
	LastExecutedSlashingRequestId uint32
}, error) {
	return _SlasherStorage.Contract.SlashingRequestIds(&_SlasherStorage.CallOpts, arg0, arg1)
}

// SlashingRequestIds is a free data retrieval call binding the contract method 0x7ef639a6.
//
// Solidity: function slashingRequestIds(address , address ) view returns(uint32 lastCreatedSlashingRequestId, uint32 lastExecutedSlashingRequestId)
func (_SlasherStorage *SlasherStorageCallerSession) SlashingRequestIds(arg0 common.Address, arg1 common.Address) (struct {
	LastCreatedSlashingRequestId  uint32
	LastExecutedSlashingRequestId uint32
}, error) {
	return _SlasherStorage.Contract.SlashingRequestIds(&_SlasherStorage.CallOpts, arg0, arg1)
}

// SlashingRequests is a free data retrieval call binding the contract method 0x3f2201bb.
//
// Solidity: function slashingRequests(address , address , uint32 ) view returns(uint32 id, uint64 slashingRate, uint64 scalingFactor)
func (_SlasherStorage *SlasherStorageCaller) SlashingRequests(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 uint32) (struct {
	Id            uint32
	SlashingRate  uint64
	ScalingFactor uint64
}, error) {
	var out []interface{}
	err := _SlasherStorage.contract.Call(opts, &out, "slashingRequests", arg0, arg1, arg2)

	outstruct := new(struct {
		Id            uint32
		SlashingRate  uint64
		ScalingFactor uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.SlashingRate = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.ScalingFactor = *abi.ConvertType(out[2], new(uint64)).(*uint64)

	return *outstruct, err

}

// SlashingRequests is a free data retrieval call binding the contract method 0x3f2201bb.
//
// Solidity: function slashingRequests(address , address , uint32 ) view returns(uint32 id, uint64 slashingRate, uint64 scalingFactor)
func (_SlasherStorage *SlasherStorageSession) SlashingRequests(arg0 common.Address, arg1 common.Address, arg2 uint32) (struct {
	Id            uint32
	SlashingRate  uint64
	ScalingFactor uint64
}, error) {
	return _SlasherStorage.Contract.SlashingRequests(&_SlasherStorage.CallOpts, arg0, arg1, arg2)
}

// SlashingRequests is a free data retrieval call binding the contract method 0x3f2201bb.
//
// Solidity: function slashingRequests(address , address , uint32 ) view returns(uint32 id, uint64 slashingRate, uint64 scalingFactor)
func (_SlasherStorage *SlasherStorageCallerSession) SlashingRequests(arg0 common.Address, arg1 common.Address, arg2 uint32) (struct {
	Id            uint32
	SlashingRate  uint64
	ScalingFactor uint64
}, error) {
	return _SlasherStorage.Contract.SlashingRequests(&_SlasherStorage.CallOpts, arg0, arg1, arg2)
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

// ExecuteSlashing is a paid mutator transaction binding the contract method 0x4dcaafb8.
//
// Solidity: function executeSlashing(address operator, address[] strategies, uint32 epoch) returns()
func (_SlasherStorage *SlasherStorageTransactor) ExecuteSlashing(opts *bind.TransactOpts, operator common.Address, strategies []common.Address, epoch uint32) (*types.Transaction, error) {
	return _SlasherStorage.contract.Transact(opts, "executeSlashing", operator, strategies, epoch)
}

// ExecuteSlashing is a paid mutator transaction binding the contract method 0x4dcaafb8.
//
// Solidity: function executeSlashing(address operator, address[] strategies, uint32 epoch) returns()
func (_SlasherStorage *SlasherStorageSession) ExecuteSlashing(operator common.Address, strategies []common.Address, epoch uint32) (*types.Transaction, error) {
	return _SlasherStorage.Contract.ExecuteSlashing(&_SlasherStorage.TransactOpts, operator, strategies, epoch)
}

// ExecuteSlashing is a paid mutator transaction binding the contract method 0x4dcaafb8.
//
// Solidity: function executeSlashing(address operator, address[] strategies, uint32 epoch) returns()
func (_SlasherStorage *SlasherStorageTransactorSession) ExecuteSlashing(operator common.Address, strategies []common.Address, epoch uint32) (*types.Transaction, error) {
	return _SlasherStorage.Contract.ExecuteSlashing(&_SlasherStorage.TransactOpts, operator, strategies, epoch)
}

// IncreaseRequestedBipsToSlash is a paid mutator transaction binding the contract method 0x287a96da.
//
// Solidity: function increaseRequestedBipsToSlash(address operator, bytes4 operatorSetID, address[] strategies, uint32 bipsToIncrease) returns()
func (_SlasherStorage *SlasherStorageTransactor) IncreaseRequestedBipsToSlash(opts *bind.TransactOpts, operator common.Address, operatorSetID [4]byte, strategies []common.Address, bipsToIncrease uint32) (*types.Transaction, error) {
	return _SlasherStorage.contract.Transact(opts, "increaseRequestedBipsToSlash", operator, operatorSetID, strategies, bipsToIncrease)
}

// IncreaseRequestedBipsToSlash is a paid mutator transaction binding the contract method 0x287a96da.
//
// Solidity: function increaseRequestedBipsToSlash(address operator, bytes4 operatorSetID, address[] strategies, uint32 bipsToIncrease) returns()
func (_SlasherStorage *SlasherStorageSession) IncreaseRequestedBipsToSlash(operator common.Address, operatorSetID [4]byte, strategies []common.Address, bipsToIncrease uint32) (*types.Transaction, error) {
	return _SlasherStorage.Contract.IncreaseRequestedBipsToSlash(&_SlasherStorage.TransactOpts, operator, operatorSetID, strategies, bipsToIncrease)
}

// IncreaseRequestedBipsToSlash is a paid mutator transaction binding the contract method 0x287a96da.
//
// Solidity: function increaseRequestedBipsToSlash(address operator, bytes4 operatorSetID, address[] strategies, uint32 bipsToIncrease) returns()
func (_SlasherStorage *SlasherStorageTransactorSession) IncreaseRequestedBipsToSlash(operator common.Address, operatorSetID [4]byte, strategies []common.Address, bipsToIncrease uint32) (*types.Transaction, error) {
	return _SlasherStorage.Contract.IncreaseRequestedBipsToSlash(&_SlasherStorage.TransactOpts, operator, operatorSetID, strategies, bipsToIncrease)
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

// ReduceRequestedBipsToSlash is a paid mutator transaction binding the contract method 0x9d086ecb.
//
// Solidity: function reduceRequestedBipsToSlash(address operator, bytes4 operatorSetID, address[] strategies, uint32 epoch, uint32 bipsToReduce) returns()
func (_SlasherStorage *SlasherStorageTransactor) ReduceRequestedBipsToSlash(opts *bind.TransactOpts, operator common.Address, operatorSetID [4]byte, strategies []common.Address, epoch uint32, bipsToReduce uint32) (*types.Transaction, error) {
	return _SlasherStorage.contract.Transact(opts, "reduceRequestedBipsToSlash", operator, operatorSetID, strategies, epoch, bipsToReduce)
}

// ReduceRequestedBipsToSlash is a paid mutator transaction binding the contract method 0x9d086ecb.
//
// Solidity: function reduceRequestedBipsToSlash(address operator, bytes4 operatorSetID, address[] strategies, uint32 epoch, uint32 bipsToReduce) returns()
func (_SlasherStorage *SlasherStorageSession) ReduceRequestedBipsToSlash(operator common.Address, operatorSetID [4]byte, strategies []common.Address, epoch uint32, bipsToReduce uint32) (*types.Transaction, error) {
	return _SlasherStorage.Contract.ReduceRequestedBipsToSlash(&_SlasherStorage.TransactOpts, operator, operatorSetID, strategies, epoch, bipsToReduce)
}

// ReduceRequestedBipsToSlash is a paid mutator transaction binding the contract method 0x9d086ecb.
//
// Solidity: function reduceRequestedBipsToSlash(address operator, bytes4 operatorSetID, address[] strategies, uint32 epoch, uint32 bipsToReduce) returns()
func (_SlasherStorage *SlasherStorageTransactorSession) ReduceRequestedBipsToSlash(operator common.Address, operatorSetID [4]byte, strategies []common.Address, epoch uint32, bipsToReduce uint32) (*types.Transaction, error) {
	return _SlasherStorage.Contract.ReduceRequestedBipsToSlash(&_SlasherStorage.TransactOpts, operator, operatorSetID, strategies, epoch, bipsToReduce)
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

// SlasherStorageRequestedBipsToSlashModifiedIterator is returned from FilterRequestedBipsToSlashModified and is used to iterate over the raw logs and unpacked data for RequestedBipsToSlashModified events raised by the SlasherStorage contract.
type SlasherStorageRequestedBipsToSlashModifiedIterator struct {
	Event *SlasherStorageRequestedBipsToSlashModified // Event containing the contract specifics and raw log

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
func (it *SlasherStorageRequestedBipsToSlashModifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlasherStorageRequestedBipsToSlashModified)
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
		it.Event = new(SlasherStorageRequestedBipsToSlashModified)
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
func (it *SlasherStorageRequestedBipsToSlashModifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlasherStorageRequestedBipsToSlashModifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlasherStorageRequestedBipsToSlashModified represents a RequestedBipsToSlashModified event raised by the SlasherStorage contract.
type SlasherStorageRequestedBipsToSlashModified struct {
	Epoch        uint32
	Operator     common.Address
	OperatorSet  IOperatorSetManagerOperatorSet
	Strategies   []common.Address
	BipsToModify int32
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRequestedBipsToSlashModified is a free log retrieval operation binding the contract event 0x51b15dc60a707d9c43660fdd6af7cf86060e2778638d04ef462faa56241ea6bf.
//
// Solidity: event RequestedBipsToSlashModified(uint32 epoch, address operator, (address,bytes4) operatorSet, address[] strategies, int32 bipsToModify)
func (_SlasherStorage *SlasherStorageFilterer) FilterRequestedBipsToSlashModified(opts *bind.FilterOpts) (*SlasherStorageRequestedBipsToSlashModifiedIterator, error) {

	logs, sub, err := _SlasherStorage.contract.FilterLogs(opts, "RequestedBipsToSlashModified")
	if err != nil {
		return nil, err
	}
	return &SlasherStorageRequestedBipsToSlashModifiedIterator{contract: _SlasherStorage.contract, event: "RequestedBipsToSlashModified", logs: logs, sub: sub}, nil
}

// WatchRequestedBipsToSlashModified is a free log subscription operation binding the contract event 0x51b15dc60a707d9c43660fdd6af7cf86060e2778638d04ef462faa56241ea6bf.
//
// Solidity: event RequestedBipsToSlashModified(uint32 epoch, address operator, (address,bytes4) operatorSet, address[] strategies, int32 bipsToModify)
func (_SlasherStorage *SlasherStorageFilterer) WatchRequestedBipsToSlashModified(opts *bind.WatchOpts, sink chan<- *SlasherStorageRequestedBipsToSlashModified) (event.Subscription, error) {

	logs, sub, err := _SlasherStorage.contract.WatchLogs(opts, "RequestedBipsToSlashModified")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlasherStorageRequestedBipsToSlashModified)
				if err := _SlasherStorage.contract.UnpackLog(event, "RequestedBipsToSlashModified", log); err != nil {
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

// ParseRequestedBipsToSlashModified is a log parse operation binding the contract event 0x51b15dc60a707d9c43660fdd6af7cf86060e2778638d04ef462faa56241ea6bf.
//
// Solidity: event RequestedBipsToSlashModified(uint32 epoch, address operator, (address,bytes4) operatorSet, address[] strategies, int32 bipsToModify)
func (_SlasherStorage *SlasherStorageFilterer) ParseRequestedBipsToSlashModified(log types.Log) (*SlasherStorageRequestedBipsToSlashModified, error) {
	event := new(SlasherStorageRequestedBipsToSlashModified)
	if err := _SlasherStorage.contract.UnpackLog(event, "RequestedBipsToSlashModified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlasherStorageSlashingExecutedIterator is returned from FilterSlashingExecuted and is used to iterate over the raw logs and unpacked data for SlashingExecuted events raised by the SlasherStorage contract.
type SlasherStorageSlashingExecutedIterator struct {
	Event *SlasherStorageSlashingExecuted // Event containing the contract specifics and raw log

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
func (it *SlasherStorageSlashingExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlasherStorageSlashingExecuted)
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
		it.Event = new(SlasherStorageSlashingExecuted)
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
func (it *SlasherStorageSlashingExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlasherStorageSlashingExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlasherStorageSlashingExecuted represents a SlashingExecuted event raised by the SlasherStorage contract.
type SlasherStorageSlashingExecuted struct {
	Epoch        uint32
	Operator     common.Address
	Strategy     common.Address
	SlashingRate uint64
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSlashingExecuted is a free log retrieval operation binding the contract event 0x2f679597a08f229c142b2f79a954c91a30bbda82795ef8dee2775b84db969924.
//
// Solidity: event SlashingExecuted(uint32 epoch, address operator, address strategy, uint64 slashingRate)
func (_SlasherStorage *SlasherStorageFilterer) FilterSlashingExecuted(opts *bind.FilterOpts) (*SlasherStorageSlashingExecutedIterator, error) {

	logs, sub, err := _SlasherStorage.contract.FilterLogs(opts, "SlashingExecuted")
	if err != nil {
		return nil, err
	}
	return &SlasherStorageSlashingExecutedIterator{contract: _SlasherStorage.contract, event: "SlashingExecuted", logs: logs, sub: sub}, nil
}

// WatchSlashingExecuted is a free log subscription operation binding the contract event 0x2f679597a08f229c142b2f79a954c91a30bbda82795ef8dee2775b84db969924.
//
// Solidity: event SlashingExecuted(uint32 epoch, address operator, address strategy, uint64 slashingRate)
func (_SlasherStorage *SlasherStorageFilterer) WatchSlashingExecuted(opts *bind.WatchOpts, sink chan<- *SlasherStorageSlashingExecuted) (event.Subscription, error) {

	logs, sub, err := _SlasherStorage.contract.WatchLogs(opts, "SlashingExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlasherStorageSlashingExecuted)
				if err := _SlasherStorage.contract.UnpackLog(event, "SlashingExecuted", log); err != nil {
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

// ParseSlashingExecuted is a log parse operation binding the contract event 0x2f679597a08f229c142b2f79a954c91a30bbda82795ef8dee2775b84db969924.
//
// Solidity: event SlashingExecuted(uint32 epoch, address operator, address strategy, uint64 slashingRate)
func (_SlasherStorage *SlasherStorageFilterer) ParseSlashingExecuted(log types.Log) (*SlasherStorageSlashingExecuted, error) {
	event := new(SlasherStorageSlashingExecuted)
	if err := _SlasherStorage.contract.UnpackLog(event, "SlashingExecuted", log); err != nil {
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
