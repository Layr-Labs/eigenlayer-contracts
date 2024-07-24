// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package EigenPodManagerStorage

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

// EigenPodManagerStorageMetaData contains all meta data concerning the EigenPodManagerStorage contract.
var EigenPodManagerStorageMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"addShares\",\"inputs\":[{\"name\":\"podOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"beaconChainETHStrategy\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createPod\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eigenPodBeacon\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBeacon\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ethPOS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIETHPOSDeposit\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPod\",\"inputs\":[{\"name\":\"podOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEigenPod\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasPod\",\"inputs\":[{\"name\":\"podOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"numPods\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownerToPod\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEigenPod\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"podOwnerShares\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"recordBeaconChainETHBalanceUpdate\",\"inputs\":[{\"name\":\"podOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sharesDelta\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeShares\",\"inputs\":[{\"name\":\"podOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slasher\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractISlasher\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stake\",\"inputs\":[{\"name\":\"pubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"depositDataRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"strategyManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawSharesAsTokens\",\"inputs\":[{\"name\":\"podOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destination\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"BeaconChainETHDeposited\",\"inputs\":[{\"name\":\"podOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BeaconChainETHWithdrawalCompleted\",\"inputs\":[{\"name\":\"podOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint96\",\"indexed\":false,\"internalType\":\"uint96\"},{\"name\":\"delegatedAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"withdrawalRoot\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewTotalShares\",\"inputs\":[{\"name\":\"podOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newTotalShares\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PodDeployed\",\"inputs\":[{\"name\":\"eigenPod\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"podOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PodSharesUpdated\",\"inputs\":[{\"name\":\"podOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sharesDelta\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// EigenPodManagerStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use EigenPodManagerStorageMetaData.ABI instead.
var EigenPodManagerStorageABI = EigenPodManagerStorageMetaData.ABI

// EigenPodManagerStorage is an auto generated Go binding around an Ethereum contract.
type EigenPodManagerStorage struct {
	EigenPodManagerStorageCaller     // Read-only binding to the contract
	EigenPodManagerStorageTransactor // Write-only binding to the contract
	EigenPodManagerStorageFilterer   // Log filterer for contract events
}

// EigenPodManagerStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type EigenPodManagerStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EigenPodManagerStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EigenPodManagerStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EigenPodManagerStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EigenPodManagerStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EigenPodManagerStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EigenPodManagerStorageSession struct {
	Contract     *EigenPodManagerStorage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// EigenPodManagerStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EigenPodManagerStorageCallerSession struct {
	Contract *EigenPodManagerStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// EigenPodManagerStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EigenPodManagerStorageTransactorSession struct {
	Contract     *EigenPodManagerStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// EigenPodManagerStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type EigenPodManagerStorageRaw struct {
	Contract *EigenPodManagerStorage // Generic contract binding to access the raw methods on
}

// EigenPodManagerStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EigenPodManagerStorageCallerRaw struct {
	Contract *EigenPodManagerStorageCaller // Generic read-only contract binding to access the raw methods on
}

// EigenPodManagerStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EigenPodManagerStorageTransactorRaw struct {
	Contract *EigenPodManagerStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEigenPodManagerStorage creates a new instance of EigenPodManagerStorage, bound to a specific deployed contract.
func NewEigenPodManagerStorage(address common.Address, backend bind.ContractBackend) (*EigenPodManagerStorage, error) {
	contract, err := bindEigenPodManagerStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EigenPodManagerStorage{EigenPodManagerStorageCaller: EigenPodManagerStorageCaller{contract: contract}, EigenPodManagerStorageTransactor: EigenPodManagerStorageTransactor{contract: contract}, EigenPodManagerStorageFilterer: EigenPodManagerStorageFilterer{contract: contract}}, nil
}

// NewEigenPodManagerStorageCaller creates a new read-only instance of EigenPodManagerStorage, bound to a specific deployed contract.
func NewEigenPodManagerStorageCaller(address common.Address, caller bind.ContractCaller) (*EigenPodManagerStorageCaller, error) {
	contract, err := bindEigenPodManagerStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EigenPodManagerStorageCaller{contract: contract}, nil
}

// NewEigenPodManagerStorageTransactor creates a new write-only instance of EigenPodManagerStorage, bound to a specific deployed contract.
func NewEigenPodManagerStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*EigenPodManagerStorageTransactor, error) {
	contract, err := bindEigenPodManagerStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EigenPodManagerStorageTransactor{contract: contract}, nil
}

// NewEigenPodManagerStorageFilterer creates a new log filterer instance of EigenPodManagerStorage, bound to a specific deployed contract.
func NewEigenPodManagerStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*EigenPodManagerStorageFilterer, error) {
	contract, err := bindEigenPodManagerStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EigenPodManagerStorageFilterer{contract: contract}, nil
}

// bindEigenPodManagerStorage binds a generic wrapper to an already deployed contract.
func bindEigenPodManagerStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EigenPodManagerStorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EigenPodManagerStorage *EigenPodManagerStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EigenPodManagerStorage.Contract.EigenPodManagerStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EigenPodManagerStorage *EigenPodManagerStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EigenPodManagerStorage.Contract.EigenPodManagerStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EigenPodManagerStorage *EigenPodManagerStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EigenPodManagerStorage.Contract.EigenPodManagerStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EigenPodManagerStorage *EigenPodManagerStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EigenPodManagerStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EigenPodManagerStorage *EigenPodManagerStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EigenPodManagerStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EigenPodManagerStorage *EigenPodManagerStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EigenPodManagerStorage.Contract.contract.Transact(opts, method, params...)
}

// BeaconChainETHStrategy is a free data retrieval call binding the contract method 0x9104c319.
//
// Solidity: function beaconChainETHStrategy() view returns(address)
func (_EigenPodManagerStorage *EigenPodManagerStorageCaller) BeaconChainETHStrategy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EigenPodManagerStorage.contract.Call(opts, &out, "beaconChainETHStrategy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BeaconChainETHStrategy is a free data retrieval call binding the contract method 0x9104c319.
//
// Solidity: function beaconChainETHStrategy() view returns(address)
func (_EigenPodManagerStorage *EigenPodManagerStorageSession) BeaconChainETHStrategy() (common.Address, error) {
	return _EigenPodManagerStorage.Contract.BeaconChainETHStrategy(&_EigenPodManagerStorage.CallOpts)
}

// BeaconChainETHStrategy is a free data retrieval call binding the contract method 0x9104c319.
//
// Solidity: function beaconChainETHStrategy() view returns(address)
func (_EigenPodManagerStorage *EigenPodManagerStorageCallerSession) BeaconChainETHStrategy() (common.Address, error) {
	return _EigenPodManagerStorage.Contract.BeaconChainETHStrategy(&_EigenPodManagerStorage.CallOpts)
}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_EigenPodManagerStorage *EigenPodManagerStorageCaller) DelegationManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EigenPodManagerStorage.contract.Call(opts, &out, "delegationManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_EigenPodManagerStorage *EigenPodManagerStorageSession) DelegationManager() (common.Address, error) {
	return _EigenPodManagerStorage.Contract.DelegationManager(&_EigenPodManagerStorage.CallOpts)
}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_EigenPodManagerStorage *EigenPodManagerStorageCallerSession) DelegationManager() (common.Address, error) {
	return _EigenPodManagerStorage.Contract.DelegationManager(&_EigenPodManagerStorage.CallOpts)
}

// EigenPodBeacon is a free data retrieval call binding the contract method 0x292b7b2b.
//
// Solidity: function eigenPodBeacon() view returns(address)
func (_EigenPodManagerStorage *EigenPodManagerStorageCaller) EigenPodBeacon(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EigenPodManagerStorage.contract.Call(opts, &out, "eigenPodBeacon")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EigenPodBeacon is a free data retrieval call binding the contract method 0x292b7b2b.
//
// Solidity: function eigenPodBeacon() view returns(address)
func (_EigenPodManagerStorage *EigenPodManagerStorageSession) EigenPodBeacon() (common.Address, error) {
	return _EigenPodManagerStorage.Contract.EigenPodBeacon(&_EigenPodManagerStorage.CallOpts)
}

// EigenPodBeacon is a free data retrieval call binding the contract method 0x292b7b2b.
//
// Solidity: function eigenPodBeacon() view returns(address)
func (_EigenPodManagerStorage *EigenPodManagerStorageCallerSession) EigenPodBeacon() (common.Address, error) {
	return _EigenPodManagerStorage.Contract.EigenPodBeacon(&_EigenPodManagerStorage.CallOpts)
}

// EthPOS is a free data retrieval call binding the contract method 0x74cdd798.
//
// Solidity: function ethPOS() view returns(address)
func (_EigenPodManagerStorage *EigenPodManagerStorageCaller) EthPOS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EigenPodManagerStorage.contract.Call(opts, &out, "ethPOS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EthPOS is a free data retrieval call binding the contract method 0x74cdd798.
//
// Solidity: function ethPOS() view returns(address)
func (_EigenPodManagerStorage *EigenPodManagerStorageSession) EthPOS() (common.Address, error) {
	return _EigenPodManagerStorage.Contract.EthPOS(&_EigenPodManagerStorage.CallOpts)
}

// EthPOS is a free data retrieval call binding the contract method 0x74cdd798.
//
// Solidity: function ethPOS() view returns(address)
func (_EigenPodManagerStorage *EigenPodManagerStorageCallerSession) EthPOS() (common.Address, error) {
	return _EigenPodManagerStorage.Contract.EthPOS(&_EigenPodManagerStorage.CallOpts)
}

// GetPod is a free data retrieval call binding the contract method 0xa38406a3.
//
// Solidity: function getPod(address podOwner) view returns(address)
func (_EigenPodManagerStorage *EigenPodManagerStorageCaller) GetPod(opts *bind.CallOpts, podOwner common.Address) (common.Address, error) {
	var out []interface{}
	err := _EigenPodManagerStorage.contract.Call(opts, &out, "getPod", podOwner)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPod is a free data retrieval call binding the contract method 0xa38406a3.
//
// Solidity: function getPod(address podOwner) view returns(address)
func (_EigenPodManagerStorage *EigenPodManagerStorageSession) GetPod(podOwner common.Address) (common.Address, error) {
	return _EigenPodManagerStorage.Contract.GetPod(&_EigenPodManagerStorage.CallOpts, podOwner)
}

// GetPod is a free data retrieval call binding the contract method 0xa38406a3.
//
// Solidity: function getPod(address podOwner) view returns(address)
func (_EigenPodManagerStorage *EigenPodManagerStorageCallerSession) GetPod(podOwner common.Address) (common.Address, error) {
	return _EigenPodManagerStorage.Contract.GetPod(&_EigenPodManagerStorage.CallOpts, podOwner)
}

// HasPod is a free data retrieval call binding the contract method 0xf6848d24.
//
// Solidity: function hasPod(address podOwner) view returns(bool)
func (_EigenPodManagerStorage *EigenPodManagerStorageCaller) HasPod(opts *bind.CallOpts, podOwner common.Address) (bool, error) {
	var out []interface{}
	err := _EigenPodManagerStorage.contract.Call(opts, &out, "hasPod", podOwner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasPod is a free data retrieval call binding the contract method 0xf6848d24.
//
// Solidity: function hasPod(address podOwner) view returns(bool)
func (_EigenPodManagerStorage *EigenPodManagerStorageSession) HasPod(podOwner common.Address) (bool, error) {
	return _EigenPodManagerStorage.Contract.HasPod(&_EigenPodManagerStorage.CallOpts, podOwner)
}

// HasPod is a free data retrieval call binding the contract method 0xf6848d24.
//
// Solidity: function hasPod(address podOwner) view returns(bool)
func (_EigenPodManagerStorage *EigenPodManagerStorageCallerSession) HasPod(podOwner common.Address) (bool, error) {
	return _EigenPodManagerStorage.Contract.HasPod(&_EigenPodManagerStorage.CallOpts, podOwner)
}

// NumPods is a free data retrieval call binding the contract method 0xa6a509be.
//
// Solidity: function numPods() view returns(uint256)
func (_EigenPodManagerStorage *EigenPodManagerStorageCaller) NumPods(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EigenPodManagerStorage.contract.Call(opts, &out, "numPods")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumPods is a free data retrieval call binding the contract method 0xa6a509be.
//
// Solidity: function numPods() view returns(uint256)
func (_EigenPodManagerStorage *EigenPodManagerStorageSession) NumPods() (*big.Int, error) {
	return _EigenPodManagerStorage.Contract.NumPods(&_EigenPodManagerStorage.CallOpts)
}

// NumPods is a free data retrieval call binding the contract method 0xa6a509be.
//
// Solidity: function numPods() view returns(uint256)
func (_EigenPodManagerStorage *EigenPodManagerStorageCallerSession) NumPods() (*big.Int, error) {
	return _EigenPodManagerStorage.Contract.NumPods(&_EigenPodManagerStorage.CallOpts)
}

// OwnerToPod is a free data retrieval call binding the contract method 0x9ba06275.
//
// Solidity: function ownerToPod(address ) view returns(address)
func (_EigenPodManagerStorage *EigenPodManagerStorageCaller) OwnerToPod(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _EigenPodManagerStorage.contract.Call(opts, &out, "ownerToPod", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerToPod is a free data retrieval call binding the contract method 0x9ba06275.
//
// Solidity: function ownerToPod(address ) view returns(address)
func (_EigenPodManagerStorage *EigenPodManagerStorageSession) OwnerToPod(arg0 common.Address) (common.Address, error) {
	return _EigenPodManagerStorage.Contract.OwnerToPod(&_EigenPodManagerStorage.CallOpts, arg0)
}

// OwnerToPod is a free data retrieval call binding the contract method 0x9ba06275.
//
// Solidity: function ownerToPod(address ) view returns(address)
func (_EigenPodManagerStorage *EigenPodManagerStorageCallerSession) OwnerToPod(arg0 common.Address) (common.Address, error) {
	return _EigenPodManagerStorage.Contract.OwnerToPod(&_EigenPodManagerStorage.CallOpts, arg0)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_EigenPodManagerStorage *EigenPodManagerStorageCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _EigenPodManagerStorage.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_EigenPodManagerStorage *EigenPodManagerStorageSession) Paused(index uint8) (bool, error) {
	return _EigenPodManagerStorage.Contract.Paused(&_EigenPodManagerStorage.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_EigenPodManagerStorage *EigenPodManagerStorageCallerSession) Paused(index uint8) (bool, error) {
	return _EigenPodManagerStorage.Contract.Paused(&_EigenPodManagerStorage.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_EigenPodManagerStorage *EigenPodManagerStorageCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EigenPodManagerStorage.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_EigenPodManagerStorage *EigenPodManagerStorageSession) Paused0() (*big.Int, error) {
	return _EigenPodManagerStorage.Contract.Paused0(&_EigenPodManagerStorage.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_EigenPodManagerStorage *EigenPodManagerStorageCallerSession) Paused0() (*big.Int, error) {
	return _EigenPodManagerStorage.Contract.Paused0(&_EigenPodManagerStorage.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_EigenPodManagerStorage *EigenPodManagerStorageCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EigenPodManagerStorage.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_EigenPodManagerStorage *EigenPodManagerStorageSession) PauserRegistry() (common.Address, error) {
	return _EigenPodManagerStorage.Contract.PauserRegistry(&_EigenPodManagerStorage.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_EigenPodManagerStorage *EigenPodManagerStorageCallerSession) PauserRegistry() (common.Address, error) {
	return _EigenPodManagerStorage.Contract.PauserRegistry(&_EigenPodManagerStorage.CallOpts)
}

// PodOwnerShares is a free data retrieval call binding the contract method 0x60f4062b.
//
// Solidity: function podOwnerShares(address ) view returns(int256)
func (_EigenPodManagerStorage *EigenPodManagerStorageCaller) PodOwnerShares(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EigenPodManagerStorage.contract.Call(opts, &out, "podOwnerShares", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PodOwnerShares is a free data retrieval call binding the contract method 0x60f4062b.
//
// Solidity: function podOwnerShares(address ) view returns(int256)
func (_EigenPodManagerStorage *EigenPodManagerStorageSession) PodOwnerShares(arg0 common.Address) (*big.Int, error) {
	return _EigenPodManagerStorage.Contract.PodOwnerShares(&_EigenPodManagerStorage.CallOpts, arg0)
}

// PodOwnerShares is a free data retrieval call binding the contract method 0x60f4062b.
//
// Solidity: function podOwnerShares(address ) view returns(int256)
func (_EigenPodManagerStorage *EigenPodManagerStorageCallerSession) PodOwnerShares(arg0 common.Address) (*big.Int, error) {
	return _EigenPodManagerStorage.Contract.PodOwnerShares(&_EigenPodManagerStorage.CallOpts, arg0)
}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_EigenPodManagerStorage *EigenPodManagerStorageCaller) Slasher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EigenPodManagerStorage.contract.Call(opts, &out, "slasher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_EigenPodManagerStorage *EigenPodManagerStorageSession) Slasher() (common.Address, error) {
	return _EigenPodManagerStorage.Contract.Slasher(&_EigenPodManagerStorage.CallOpts)
}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_EigenPodManagerStorage *EigenPodManagerStorageCallerSession) Slasher() (common.Address, error) {
	return _EigenPodManagerStorage.Contract.Slasher(&_EigenPodManagerStorage.CallOpts)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_EigenPodManagerStorage *EigenPodManagerStorageCaller) StrategyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EigenPodManagerStorage.contract.Call(opts, &out, "strategyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_EigenPodManagerStorage *EigenPodManagerStorageSession) StrategyManager() (common.Address, error) {
	return _EigenPodManagerStorage.Contract.StrategyManager(&_EigenPodManagerStorage.CallOpts)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_EigenPodManagerStorage *EigenPodManagerStorageCallerSession) StrategyManager() (common.Address, error) {
	return _EigenPodManagerStorage.Contract.StrategyManager(&_EigenPodManagerStorage.CallOpts)
}

// AddShares is a paid mutator transaction binding the contract method 0x0e81073c.
//
// Solidity: function addShares(address podOwner, uint256 shares) returns(uint256)
func (_EigenPodManagerStorage *EigenPodManagerStorageTransactor) AddShares(opts *bind.TransactOpts, podOwner common.Address, shares *big.Int) (*types.Transaction, error) {
	return _EigenPodManagerStorage.contract.Transact(opts, "addShares", podOwner, shares)
}

// AddShares is a paid mutator transaction binding the contract method 0x0e81073c.
//
// Solidity: function addShares(address podOwner, uint256 shares) returns(uint256)
func (_EigenPodManagerStorage *EigenPodManagerStorageSession) AddShares(podOwner common.Address, shares *big.Int) (*types.Transaction, error) {
	return _EigenPodManagerStorage.Contract.AddShares(&_EigenPodManagerStorage.TransactOpts, podOwner, shares)
}

// AddShares is a paid mutator transaction binding the contract method 0x0e81073c.
//
// Solidity: function addShares(address podOwner, uint256 shares) returns(uint256)
func (_EigenPodManagerStorage *EigenPodManagerStorageTransactorSession) AddShares(podOwner common.Address, shares *big.Int) (*types.Transaction, error) {
	return _EigenPodManagerStorage.Contract.AddShares(&_EigenPodManagerStorage.TransactOpts, podOwner, shares)
}

// CreatePod is a paid mutator transaction binding the contract method 0x84d81062.
//
// Solidity: function createPod() returns(address)
func (_EigenPodManagerStorage *EigenPodManagerStorageTransactor) CreatePod(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EigenPodManagerStorage.contract.Transact(opts, "createPod")
}

// CreatePod is a paid mutator transaction binding the contract method 0x84d81062.
//
// Solidity: function createPod() returns(address)
func (_EigenPodManagerStorage *EigenPodManagerStorageSession) CreatePod() (*types.Transaction, error) {
	return _EigenPodManagerStorage.Contract.CreatePod(&_EigenPodManagerStorage.TransactOpts)
}

// CreatePod is a paid mutator transaction binding the contract method 0x84d81062.
//
// Solidity: function createPod() returns(address)
func (_EigenPodManagerStorage *EigenPodManagerStorageTransactorSession) CreatePod() (*types.Transaction, error) {
	return _EigenPodManagerStorage.Contract.CreatePod(&_EigenPodManagerStorage.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_EigenPodManagerStorage *EigenPodManagerStorageTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _EigenPodManagerStorage.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_EigenPodManagerStorage *EigenPodManagerStorageSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _EigenPodManagerStorage.Contract.Pause(&_EigenPodManagerStorage.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_EigenPodManagerStorage *EigenPodManagerStorageTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _EigenPodManagerStorage.Contract.Pause(&_EigenPodManagerStorage.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_EigenPodManagerStorage *EigenPodManagerStorageTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EigenPodManagerStorage.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_EigenPodManagerStorage *EigenPodManagerStorageSession) PauseAll() (*types.Transaction, error) {
	return _EigenPodManagerStorage.Contract.PauseAll(&_EigenPodManagerStorage.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_EigenPodManagerStorage *EigenPodManagerStorageTransactorSession) PauseAll() (*types.Transaction, error) {
	return _EigenPodManagerStorage.Contract.PauseAll(&_EigenPodManagerStorage.TransactOpts)
}

// RecordBeaconChainETHBalanceUpdate is a paid mutator transaction binding the contract method 0xc2c51c40.
//
// Solidity: function recordBeaconChainETHBalanceUpdate(address podOwner, int256 sharesDelta) returns()
func (_EigenPodManagerStorage *EigenPodManagerStorageTransactor) RecordBeaconChainETHBalanceUpdate(opts *bind.TransactOpts, podOwner common.Address, sharesDelta *big.Int) (*types.Transaction, error) {
	return _EigenPodManagerStorage.contract.Transact(opts, "recordBeaconChainETHBalanceUpdate", podOwner, sharesDelta)
}

// RecordBeaconChainETHBalanceUpdate is a paid mutator transaction binding the contract method 0xc2c51c40.
//
// Solidity: function recordBeaconChainETHBalanceUpdate(address podOwner, int256 sharesDelta) returns()
func (_EigenPodManagerStorage *EigenPodManagerStorageSession) RecordBeaconChainETHBalanceUpdate(podOwner common.Address, sharesDelta *big.Int) (*types.Transaction, error) {
	return _EigenPodManagerStorage.Contract.RecordBeaconChainETHBalanceUpdate(&_EigenPodManagerStorage.TransactOpts, podOwner, sharesDelta)
}

// RecordBeaconChainETHBalanceUpdate is a paid mutator transaction binding the contract method 0xc2c51c40.
//
// Solidity: function recordBeaconChainETHBalanceUpdate(address podOwner, int256 sharesDelta) returns()
func (_EigenPodManagerStorage *EigenPodManagerStorageTransactorSession) RecordBeaconChainETHBalanceUpdate(podOwner common.Address, sharesDelta *big.Int) (*types.Transaction, error) {
	return _EigenPodManagerStorage.Contract.RecordBeaconChainETHBalanceUpdate(&_EigenPodManagerStorage.TransactOpts, podOwner, sharesDelta)
}

// RemoveShares is a paid mutator transaction binding the contract method 0xbeffbb89.
//
// Solidity: function removeShares(address podOwner, uint256 shares) returns()
func (_EigenPodManagerStorage *EigenPodManagerStorageTransactor) RemoveShares(opts *bind.TransactOpts, podOwner common.Address, shares *big.Int) (*types.Transaction, error) {
	return _EigenPodManagerStorage.contract.Transact(opts, "removeShares", podOwner, shares)
}

// RemoveShares is a paid mutator transaction binding the contract method 0xbeffbb89.
//
// Solidity: function removeShares(address podOwner, uint256 shares) returns()
func (_EigenPodManagerStorage *EigenPodManagerStorageSession) RemoveShares(podOwner common.Address, shares *big.Int) (*types.Transaction, error) {
	return _EigenPodManagerStorage.Contract.RemoveShares(&_EigenPodManagerStorage.TransactOpts, podOwner, shares)
}

// RemoveShares is a paid mutator transaction binding the contract method 0xbeffbb89.
//
// Solidity: function removeShares(address podOwner, uint256 shares) returns()
func (_EigenPodManagerStorage *EigenPodManagerStorageTransactorSession) RemoveShares(podOwner common.Address, shares *big.Int) (*types.Transaction, error) {
	return _EigenPodManagerStorage.Contract.RemoveShares(&_EigenPodManagerStorage.TransactOpts, podOwner, shares)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_EigenPodManagerStorage *EigenPodManagerStorageTransactor) SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error) {
	return _EigenPodManagerStorage.contract.Transact(opts, "setPauserRegistry", newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_EigenPodManagerStorage *EigenPodManagerStorageSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _EigenPodManagerStorage.Contract.SetPauserRegistry(&_EigenPodManagerStorage.TransactOpts, newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_EigenPodManagerStorage *EigenPodManagerStorageTransactorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _EigenPodManagerStorage.Contract.SetPauserRegistry(&_EigenPodManagerStorage.TransactOpts, newPauserRegistry)
}

// Stake is a paid mutator transaction binding the contract method 0x9b4e4634.
//
// Solidity: function stake(bytes pubkey, bytes signature, bytes32 depositDataRoot) payable returns()
func (_EigenPodManagerStorage *EigenPodManagerStorageTransactor) Stake(opts *bind.TransactOpts, pubkey []byte, signature []byte, depositDataRoot [32]byte) (*types.Transaction, error) {
	return _EigenPodManagerStorage.contract.Transact(opts, "stake", pubkey, signature, depositDataRoot)
}

// Stake is a paid mutator transaction binding the contract method 0x9b4e4634.
//
// Solidity: function stake(bytes pubkey, bytes signature, bytes32 depositDataRoot) payable returns()
func (_EigenPodManagerStorage *EigenPodManagerStorageSession) Stake(pubkey []byte, signature []byte, depositDataRoot [32]byte) (*types.Transaction, error) {
	return _EigenPodManagerStorage.Contract.Stake(&_EigenPodManagerStorage.TransactOpts, pubkey, signature, depositDataRoot)
}

// Stake is a paid mutator transaction binding the contract method 0x9b4e4634.
//
// Solidity: function stake(bytes pubkey, bytes signature, bytes32 depositDataRoot) payable returns()
func (_EigenPodManagerStorage *EigenPodManagerStorageTransactorSession) Stake(pubkey []byte, signature []byte, depositDataRoot [32]byte) (*types.Transaction, error) {
	return _EigenPodManagerStorage.Contract.Stake(&_EigenPodManagerStorage.TransactOpts, pubkey, signature, depositDataRoot)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_EigenPodManagerStorage *EigenPodManagerStorageTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _EigenPodManagerStorage.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_EigenPodManagerStorage *EigenPodManagerStorageSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _EigenPodManagerStorage.Contract.Unpause(&_EigenPodManagerStorage.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_EigenPodManagerStorage *EigenPodManagerStorageTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _EigenPodManagerStorage.Contract.Unpause(&_EigenPodManagerStorage.TransactOpts, newPausedStatus)
}

// WithdrawSharesAsTokens is a paid mutator transaction binding the contract method 0x387b1300.
//
// Solidity: function withdrawSharesAsTokens(address podOwner, address destination, uint256 shares) returns()
func (_EigenPodManagerStorage *EigenPodManagerStorageTransactor) WithdrawSharesAsTokens(opts *bind.TransactOpts, podOwner common.Address, destination common.Address, shares *big.Int) (*types.Transaction, error) {
	return _EigenPodManagerStorage.contract.Transact(opts, "withdrawSharesAsTokens", podOwner, destination, shares)
}

// WithdrawSharesAsTokens is a paid mutator transaction binding the contract method 0x387b1300.
//
// Solidity: function withdrawSharesAsTokens(address podOwner, address destination, uint256 shares) returns()
func (_EigenPodManagerStorage *EigenPodManagerStorageSession) WithdrawSharesAsTokens(podOwner common.Address, destination common.Address, shares *big.Int) (*types.Transaction, error) {
	return _EigenPodManagerStorage.Contract.WithdrawSharesAsTokens(&_EigenPodManagerStorage.TransactOpts, podOwner, destination, shares)
}

// WithdrawSharesAsTokens is a paid mutator transaction binding the contract method 0x387b1300.
//
// Solidity: function withdrawSharesAsTokens(address podOwner, address destination, uint256 shares) returns()
func (_EigenPodManagerStorage *EigenPodManagerStorageTransactorSession) WithdrawSharesAsTokens(podOwner common.Address, destination common.Address, shares *big.Int) (*types.Transaction, error) {
	return _EigenPodManagerStorage.Contract.WithdrawSharesAsTokens(&_EigenPodManagerStorage.TransactOpts, podOwner, destination, shares)
}

// EigenPodManagerStorageBeaconChainETHDepositedIterator is returned from FilterBeaconChainETHDeposited and is used to iterate over the raw logs and unpacked data for BeaconChainETHDeposited events raised by the EigenPodManagerStorage contract.
type EigenPodManagerStorageBeaconChainETHDepositedIterator struct {
	Event *EigenPodManagerStorageBeaconChainETHDeposited // Event containing the contract specifics and raw log

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
func (it *EigenPodManagerStorageBeaconChainETHDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenPodManagerStorageBeaconChainETHDeposited)
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
		it.Event = new(EigenPodManagerStorageBeaconChainETHDeposited)
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
func (it *EigenPodManagerStorageBeaconChainETHDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenPodManagerStorageBeaconChainETHDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenPodManagerStorageBeaconChainETHDeposited represents a BeaconChainETHDeposited event raised by the EigenPodManagerStorage contract.
type EigenPodManagerStorageBeaconChainETHDeposited struct {
	PodOwner common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBeaconChainETHDeposited is a free log retrieval operation binding the contract event 0x35a85cabc603f48abb2b71d9fbd8adea7c449d7f0be900ae7a2986ea369c3d0d.
//
// Solidity: event BeaconChainETHDeposited(address indexed podOwner, uint256 amount)
func (_EigenPodManagerStorage *EigenPodManagerStorageFilterer) FilterBeaconChainETHDeposited(opts *bind.FilterOpts, podOwner []common.Address) (*EigenPodManagerStorageBeaconChainETHDepositedIterator, error) {

	var podOwnerRule []interface{}
	for _, podOwnerItem := range podOwner {
		podOwnerRule = append(podOwnerRule, podOwnerItem)
	}

	logs, sub, err := _EigenPodManagerStorage.contract.FilterLogs(opts, "BeaconChainETHDeposited", podOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EigenPodManagerStorageBeaconChainETHDepositedIterator{contract: _EigenPodManagerStorage.contract, event: "BeaconChainETHDeposited", logs: logs, sub: sub}, nil
}

// WatchBeaconChainETHDeposited is a free log subscription operation binding the contract event 0x35a85cabc603f48abb2b71d9fbd8adea7c449d7f0be900ae7a2986ea369c3d0d.
//
// Solidity: event BeaconChainETHDeposited(address indexed podOwner, uint256 amount)
func (_EigenPodManagerStorage *EigenPodManagerStorageFilterer) WatchBeaconChainETHDeposited(opts *bind.WatchOpts, sink chan<- *EigenPodManagerStorageBeaconChainETHDeposited, podOwner []common.Address) (event.Subscription, error) {

	var podOwnerRule []interface{}
	for _, podOwnerItem := range podOwner {
		podOwnerRule = append(podOwnerRule, podOwnerItem)
	}

	logs, sub, err := _EigenPodManagerStorage.contract.WatchLogs(opts, "BeaconChainETHDeposited", podOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenPodManagerStorageBeaconChainETHDeposited)
				if err := _EigenPodManagerStorage.contract.UnpackLog(event, "BeaconChainETHDeposited", log); err != nil {
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

// ParseBeaconChainETHDeposited is a log parse operation binding the contract event 0x35a85cabc603f48abb2b71d9fbd8adea7c449d7f0be900ae7a2986ea369c3d0d.
//
// Solidity: event BeaconChainETHDeposited(address indexed podOwner, uint256 amount)
func (_EigenPodManagerStorage *EigenPodManagerStorageFilterer) ParseBeaconChainETHDeposited(log types.Log) (*EigenPodManagerStorageBeaconChainETHDeposited, error) {
	event := new(EigenPodManagerStorageBeaconChainETHDeposited)
	if err := _EigenPodManagerStorage.contract.UnpackLog(event, "BeaconChainETHDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenPodManagerStorageBeaconChainETHWithdrawalCompletedIterator is returned from FilterBeaconChainETHWithdrawalCompleted and is used to iterate over the raw logs and unpacked data for BeaconChainETHWithdrawalCompleted events raised by the EigenPodManagerStorage contract.
type EigenPodManagerStorageBeaconChainETHWithdrawalCompletedIterator struct {
	Event *EigenPodManagerStorageBeaconChainETHWithdrawalCompleted // Event containing the contract specifics and raw log

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
func (it *EigenPodManagerStorageBeaconChainETHWithdrawalCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenPodManagerStorageBeaconChainETHWithdrawalCompleted)
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
		it.Event = new(EigenPodManagerStorageBeaconChainETHWithdrawalCompleted)
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
func (it *EigenPodManagerStorageBeaconChainETHWithdrawalCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenPodManagerStorageBeaconChainETHWithdrawalCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenPodManagerStorageBeaconChainETHWithdrawalCompleted represents a BeaconChainETHWithdrawalCompleted event raised by the EigenPodManagerStorage contract.
type EigenPodManagerStorageBeaconChainETHWithdrawalCompleted struct {
	PodOwner         common.Address
	Shares           *big.Int
	Nonce            *big.Int
	DelegatedAddress common.Address
	Withdrawer       common.Address
	WithdrawalRoot   [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterBeaconChainETHWithdrawalCompleted is a free log retrieval operation binding the contract event 0xa6bab1d55a361fcea2eee2bc9491e4f01e6cf333df03c9c4f2c144466429f7d6.
//
// Solidity: event BeaconChainETHWithdrawalCompleted(address indexed podOwner, uint256 shares, uint96 nonce, address delegatedAddress, address withdrawer, bytes32 withdrawalRoot)
func (_EigenPodManagerStorage *EigenPodManagerStorageFilterer) FilterBeaconChainETHWithdrawalCompleted(opts *bind.FilterOpts, podOwner []common.Address) (*EigenPodManagerStorageBeaconChainETHWithdrawalCompletedIterator, error) {

	var podOwnerRule []interface{}
	for _, podOwnerItem := range podOwner {
		podOwnerRule = append(podOwnerRule, podOwnerItem)
	}

	logs, sub, err := _EigenPodManagerStorage.contract.FilterLogs(opts, "BeaconChainETHWithdrawalCompleted", podOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EigenPodManagerStorageBeaconChainETHWithdrawalCompletedIterator{contract: _EigenPodManagerStorage.contract, event: "BeaconChainETHWithdrawalCompleted", logs: logs, sub: sub}, nil
}

// WatchBeaconChainETHWithdrawalCompleted is a free log subscription operation binding the contract event 0xa6bab1d55a361fcea2eee2bc9491e4f01e6cf333df03c9c4f2c144466429f7d6.
//
// Solidity: event BeaconChainETHWithdrawalCompleted(address indexed podOwner, uint256 shares, uint96 nonce, address delegatedAddress, address withdrawer, bytes32 withdrawalRoot)
func (_EigenPodManagerStorage *EigenPodManagerStorageFilterer) WatchBeaconChainETHWithdrawalCompleted(opts *bind.WatchOpts, sink chan<- *EigenPodManagerStorageBeaconChainETHWithdrawalCompleted, podOwner []common.Address) (event.Subscription, error) {

	var podOwnerRule []interface{}
	for _, podOwnerItem := range podOwner {
		podOwnerRule = append(podOwnerRule, podOwnerItem)
	}

	logs, sub, err := _EigenPodManagerStorage.contract.WatchLogs(opts, "BeaconChainETHWithdrawalCompleted", podOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenPodManagerStorageBeaconChainETHWithdrawalCompleted)
				if err := _EigenPodManagerStorage.contract.UnpackLog(event, "BeaconChainETHWithdrawalCompleted", log); err != nil {
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

// ParseBeaconChainETHWithdrawalCompleted is a log parse operation binding the contract event 0xa6bab1d55a361fcea2eee2bc9491e4f01e6cf333df03c9c4f2c144466429f7d6.
//
// Solidity: event BeaconChainETHWithdrawalCompleted(address indexed podOwner, uint256 shares, uint96 nonce, address delegatedAddress, address withdrawer, bytes32 withdrawalRoot)
func (_EigenPodManagerStorage *EigenPodManagerStorageFilterer) ParseBeaconChainETHWithdrawalCompleted(log types.Log) (*EigenPodManagerStorageBeaconChainETHWithdrawalCompleted, error) {
	event := new(EigenPodManagerStorageBeaconChainETHWithdrawalCompleted)
	if err := _EigenPodManagerStorage.contract.UnpackLog(event, "BeaconChainETHWithdrawalCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenPodManagerStorageNewTotalSharesIterator is returned from FilterNewTotalShares and is used to iterate over the raw logs and unpacked data for NewTotalShares events raised by the EigenPodManagerStorage contract.
type EigenPodManagerStorageNewTotalSharesIterator struct {
	Event *EigenPodManagerStorageNewTotalShares // Event containing the contract specifics and raw log

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
func (it *EigenPodManagerStorageNewTotalSharesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenPodManagerStorageNewTotalShares)
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
		it.Event = new(EigenPodManagerStorageNewTotalShares)
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
func (it *EigenPodManagerStorageNewTotalSharesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenPodManagerStorageNewTotalSharesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenPodManagerStorageNewTotalShares represents a NewTotalShares event raised by the EigenPodManagerStorage contract.
type EigenPodManagerStorageNewTotalShares struct {
	PodOwner       common.Address
	NewTotalShares *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewTotalShares is a free log retrieval operation binding the contract event 0xd4def76d6d2bed6f14d5cd9af73cc2913d618d00edde42432e81c09bfe077098.
//
// Solidity: event NewTotalShares(address indexed podOwner, int256 newTotalShares)
func (_EigenPodManagerStorage *EigenPodManagerStorageFilterer) FilterNewTotalShares(opts *bind.FilterOpts, podOwner []common.Address) (*EigenPodManagerStorageNewTotalSharesIterator, error) {

	var podOwnerRule []interface{}
	for _, podOwnerItem := range podOwner {
		podOwnerRule = append(podOwnerRule, podOwnerItem)
	}

	logs, sub, err := _EigenPodManagerStorage.contract.FilterLogs(opts, "NewTotalShares", podOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EigenPodManagerStorageNewTotalSharesIterator{contract: _EigenPodManagerStorage.contract, event: "NewTotalShares", logs: logs, sub: sub}, nil
}

// WatchNewTotalShares is a free log subscription operation binding the contract event 0xd4def76d6d2bed6f14d5cd9af73cc2913d618d00edde42432e81c09bfe077098.
//
// Solidity: event NewTotalShares(address indexed podOwner, int256 newTotalShares)
func (_EigenPodManagerStorage *EigenPodManagerStorageFilterer) WatchNewTotalShares(opts *bind.WatchOpts, sink chan<- *EigenPodManagerStorageNewTotalShares, podOwner []common.Address) (event.Subscription, error) {

	var podOwnerRule []interface{}
	for _, podOwnerItem := range podOwner {
		podOwnerRule = append(podOwnerRule, podOwnerItem)
	}

	logs, sub, err := _EigenPodManagerStorage.contract.WatchLogs(opts, "NewTotalShares", podOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenPodManagerStorageNewTotalShares)
				if err := _EigenPodManagerStorage.contract.UnpackLog(event, "NewTotalShares", log); err != nil {
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

// ParseNewTotalShares is a log parse operation binding the contract event 0xd4def76d6d2bed6f14d5cd9af73cc2913d618d00edde42432e81c09bfe077098.
//
// Solidity: event NewTotalShares(address indexed podOwner, int256 newTotalShares)
func (_EigenPodManagerStorage *EigenPodManagerStorageFilterer) ParseNewTotalShares(log types.Log) (*EigenPodManagerStorageNewTotalShares, error) {
	event := new(EigenPodManagerStorageNewTotalShares)
	if err := _EigenPodManagerStorage.contract.UnpackLog(event, "NewTotalShares", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenPodManagerStoragePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the EigenPodManagerStorage contract.
type EigenPodManagerStoragePausedIterator struct {
	Event *EigenPodManagerStoragePaused // Event containing the contract specifics and raw log

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
func (it *EigenPodManagerStoragePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenPodManagerStoragePaused)
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
		it.Event = new(EigenPodManagerStoragePaused)
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
func (it *EigenPodManagerStoragePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenPodManagerStoragePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenPodManagerStoragePaused represents a Paused event raised by the EigenPodManagerStorage contract.
type EigenPodManagerStoragePaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_EigenPodManagerStorage *EigenPodManagerStorageFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*EigenPodManagerStoragePausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EigenPodManagerStorage.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &EigenPodManagerStoragePausedIterator{contract: _EigenPodManagerStorage.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_EigenPodManagerStorage *EigenPodManagerStorageFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *EigenPodManagerStoragePaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EigenPodManagerStorage.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenPodManagerStoragePaused)
				if err := _EigenPodManagerStorage.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_EigenPodManagerStorage *EigenPodManagerStorageFilterer) ParsePaused(log types.Log) (*EigenPodManagerStoragePaused, error) {
	event := new(EigenPodManagerStoragePaused)
	if err := _EigenPodManagerStorage.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenPodManagerStoragePauserRegistrySetIterator is returned from FilterPauserRegistrySet and is used to iterate over the raw logs and unpacked data for PauserRegistrySet events raised by the EigenPodManagerStorage contract.
type EigenPodManagerStoragePauserRegistrySetIterator struct {
	Event *EigenPodManagerStoragePauserRegistrySet // Event containing the contract specifics and raw log

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
func (it *EigenPodManagerStoragePauserRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenPodManagerStoragePauserRegistrySet)
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
		it.Event = new(EigenPodManagerStoragePauserRegistrySet)
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
func (it *EigenPodManagerStoragePauserRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenPodManagerStoragePauserRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenPodManagerStoragePauserRegistrySet represents a PauserRegistrySet event raised by the EigenPodManagerStorage contract.
type EigenPodManagerStoragePauserRegistrySet struct {
	PauserRegistry    common.Address
	NewPauserRegistry common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPauserRegistrySet is a free log retrieval operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_EigenPodManagerStorage *EigenPodManagerStorageFilterer) FilterPauserRegistrySet(opts *bind.FilterOpts) (*EigenPodManagerStoragePauserRegistrySetIterator, error) {

	logs, sub, err := _EigenPodManagerStorage.contract.FilterLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return &EigenPodManagerStoragePauserRegistrySetIterator{contract: _EigenPodManagerStorage.contract, event: "PauserRegistrySet", logs: logs, sub: sub}, nil
}

// WatchPauserRegistrySet is a free log subscription operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_EigenPodManagerStorage *EigenPodManagerStorageFilterer) WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *EigenPodManagerStoragePauserRegistrySet) (event.Subscription, error) {

	logs, sub, err := _EigenPodManagerStorage.contract.WatchLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenPodManagerStoragePauserRegistrySet)
				if err := _EigenPodManagerStorage.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
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
func (_EigenPodManagerStorage *EigenPodManagerStorageFilterer) ParsePauserRegistrySet(log types.Log) (*EigenPodManagerStoragePauserRegistrySet, error) {
	event := new(EigenPodManagerStoragePauserRegistrySet)
	if err := _EigenPodManagerStorage.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenPodManagerStoragePodDeployedIterator is returned from FilterPodDeployed and is used to iterate over the raw logs and unpacked data for PodDeployed events raised by the EigenPodManagerStorage contract.
type EigenPodManagerStoragePodDeployedIterator struct {
	Event *EigenPodManagerStoragePodDeployed // Event containing the contract specifics and raw log

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
func (it *EigenPodManagerStoragePodDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenPodManagerStoragePodDeployed)
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
		it.Event = new(EigenPodManagerStoragePodDeployed)
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
func (it *EigenPodManagerStoragePodDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenPodManagerStoragePodDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenPodManagerStoragePodDeployed represents a PodDeployed event raised by the EigenPodManagerStorage contract.
type EigenPodManagerStoragePodDeployed struct {
	EigenPod common.Address
	PodOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPodDeployed is a free log retrieval operation binding the contract event 0x21c99d0db02213c32fff5b05cf0a718ab5f858802b91498f80d82270289d856a.
//
// Solidity: event PodDeployed(address indexed eigenPod, address indexed podOwner)
func (_EigenPodManagerStorage *EigenPodManagerStorageFilterer) FilterPodDeployed(opts *bind.FilterOpts, eigenPod []common.Address, podOwner []common.Address) (*EigenPodManagerStoragePodDeployedIterator, error) {

	var eigenPodRule []interface{}
	for _, eigenPodItem := range eigenPod {
		eigenPodRule = append(eigenPodRule, eigenPodItem)
	}
	var podOwnerRule []interface{}
	for _, podOwnerItem := range podOwner {
		podOwnerRule = append(podOwnerRule, podOwnerItem)
	}

	logs, sub, err := _EigenPodManagerStorage.contract.FilterLogs(opts, "PodDeployed", eigenPodRule, podOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EigenPodManagerStoragePodDeployedIterator{contract: _EigenPodManagerStorage.contract, event: "PodDeployed", logs: logs, sub: sub}, nil
}

// WatchPodDeployed is a free log subscription operation binding the contract event 0x21c99d0db02213c32fff5b05cf0a718ab5f858802b91498f80d82270289d856a.
//
// Solidity: event PodDeployed(address indexed eigenPod, address indexed podOwner)
func (_EigenPodManagerStorage *EigenPodManagerStorageFilterer) WatchPodDeployed(opts *bind.WatchOpts, sink chan<- *EigenPodManagerStoragePodDeployed, eigenPod []common.Address, podOwner []common.Address) (event.Subscription, error) {

	var eigenPodRule []interface{}
	for _, eigenPodItem := range eigenPod {
		eigenPodRule = append(eigenPodRule, eigenPodItem)
	}
	var podOwnerRule []interface{}
	for _, podOwnerItem := range podOwner {
		podOwnerRule = append(podOwnerRule, podOwnerItem)
	}

	logs, sub, err := _EigenPodManagerStorage.contract.WatchLogs(opts, "PodDeployed", eigenPodRule, podOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenPodManagerStoragePodDeployed)
				if err := _EigenPodManagerStorage.contract.UnpackLog(event, "PodDeployed", log); err != nil {
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

// ParsePodDeployed is a log parse operation binding the contract event 0x21c99d0db02213c32fff5b05cf0a718ab5f858802b91498f80d82270289d856a.
//
// Solidity: event PodDeployed(address indexed eigenPod, address indexed podOwner)
func (_EigenPodManagerStorage *EigenPodManagerStorageFilterer) ParsePodDeployed(log types.Log) (*EigenPodManagerStoragePodDeployed, error) {
	event := new(EigenPodManagerStoragePodDeployed)
	if err := _EigenPodManagerStorage.contract.UnpackLog(event, "PodDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenPodManagerStoragePodSharesUpdatedIterator is returned from FilterPodSharesUpdated and is used to iterate over the raw logs and unpacked data for PodSharesUpdated events raised by the EigenPodManagerStorage contract.
type EigenPodManagerStoragePodSharesUpdatedIterator struct {
	Event *EigenPodManagerStoragePodSharesUpdated // Event containing the contract specifics and raw log

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
func (it *EigenPodManagerStoragePodSharesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenPodManagerStoragePodSharesUpdated)
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
		it.Event = new(EigenPodManagerStoragePodSharesUpdated)
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
func (it *EigenPodManagerStoragePodSharesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenPodManagerStoragePodSharesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenPodManagerStoragePodSharesUpdated represents a PodSharesUpdated event raised by the EigenPodManagerStorage contract.
type EigenPodManagerStoragePodSharesUpdated struct {
	PodOwner    common.Address
	SharesDelta *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPodSharesUpdated is a free log retrieval operation binding the contract event 0x4e2b791dedccd9fb30141b088cabf5c14a8912b52f59375c95c010700b8c6193.
//
// Solidity: event PodSharesUpdated(address indexed podOwner, int256 sharesDelta)
func (_EigenPodManagerStorage *EigenPodManagerStorageFilterer) FilterPodSharesUpdated(opts *bind.FilterOpts, podOwner []common.Address) (*EigenPodManagerStoragePodSharesUpdatedIterator, error) {

	var podOwnerRule []interface{}
	for _, podOwnerItem := range podOwner {
		podOwnerRule = append(podOwnerRule, podOwnerItem)
	}

	logs, sub, err := _EigenPodManagerStorage.contract.FilterLogs(opts, "PodSharesUpdated", podOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EigenPodManagerStoragePodSharesUpdatedIterator{contract: _EigenPodManagerStorage.contract, event: "PodSharesUpdated", logs: logs, sub: sub}, nil
}

// WatchPodSharesUpdated is a free log subscription operation binding the contract event 0x4e2b791dedccd9fb30141b088cabf5c14a8912b52f59375c95c010700b8c6193.
//
// Solidity: event PodSharesUpdated(address indexed podOwner, int256 sharesDelta)
func (_EigenPodManagerStorage *EigenPodManagerStorageFilterer) WatchPodSharesUpdated(opts *bind.WatchOpts, sink chan<- *EigenPodManagerStoragePodSharesUpdated, podOwner []common.Address) (event.Subscription, error) {

	var podOwnerRule []interface{}
	for _, podOwnerItem := range podOwner {
		podOwnerRule = append(podOwnerRule, podOwnerItem)
	}

	logs, sub, err := _EigenPodManagerStorage.contract.WatchLogs(opts, "PodSharesUpdated", podOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenPodManagerStoragePodSharesUpdated)
				if err := _EigenPodManagerStorage.contract.UnpackLog(event, "PodSharesUpdated", log); err != nil {
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

// ParsePodSharesUpdated is a log parse operation binding the contract event 0x4e2b791dedccd9fb30141b088cabf5c14a8912b52f59375c95c010700b8c6193.
//
// Solidity: event PodSharesUpdated(address indexed podOwner, int256 sharesDelta)
func (_EigenPodManagerStorage *EigenPodManagerStorageFilterer) ParsePodSharesUpdated(log types.Log) (*EigenPodManagerStoragePodSharesUpdated, error) {
	event := new(EigenPodManagerStoragePodSharesUpdated)
	if err := _EigenPodManagerStorage.contract.UnpackLog(event, "PodSharesUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenPodManagerStorageUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the EigenPodManagerStorage contract.
type EigenPodManagerStorageUnpausedIterator struct {
	Event *EigenPodManagerStorageUnpaused // Event containing the contract specifics and raw log

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
func (it *EigenPodManagerStorageUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenPodManagerStorageUnpaused)
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
		it.Event = new(EigenPodManagerStorageUnpaused)
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
func (it *EigenPodManagerStorageUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenPodManagerStorageUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenPodManagerStorageUnpaused represents a Unpaused event raised by the EigenPodManagerStorage contract.
type EigenPodManagerStorageUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_EigenPodManagerStorage *EigenPodManagerStorageFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*EigenPodManagerStorageUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EigenPodManagerStorage.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &EigenPodManagerStorageUnpausedIterator{contract: _EigenPodManagerStorage.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_EigenPodManagerStorage *EigenPodManagerStorageFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EigenPodManagerStorageUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EigenPodManagerStorage.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenPodManagerStorageUnpaused)
				if err := _EigenPodManagerStorage.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_EigenPodManagerStorage *EigenPodManagerStorageFilterer) ParseUnpaused(log types.Log) (*EigenPodManagerStorageUnpaused, error) {
	event := new(EigenPodManagerStorageUnpaused)
	if err := _EigenPodManagerStorage.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
