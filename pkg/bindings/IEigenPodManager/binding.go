// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IEigenPodManager

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

// IEigenPodManagerMetaData contains all meta data concerning the IEigenPodManager contract.
var IEigenPodManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"addShares\",\"inputs\":[{\"name\":\"podOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"beaconChainETHStrategy\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"beaconChainOracle\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBeaconChainOracle\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createPod\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"denebForkTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eigenPodBeacon\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBeacon\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ethPOS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIETHPOSDeposit\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBlockRootAtTimestamp\",\"inputs\":[{\"name\":\"timestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPod\",\"inputs\":[{\"name\":\"podOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEigenPod\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasPod\",\"inputs\":[{\"name\":\"podOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"numPods\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownerToPod\",\"inputs\":[{\"name\":\"podOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEigenPod\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"podOwnerShares\",\"inputs\":[{\"name\":\"podOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"recordBeaconChainETHBalanceUpdate\",\"inputs\":[{\"name\":\"podOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sharesDelta\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeShares\",\"inputs\":[{\"name\":\"podOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDenebForkTimestamp\",\"inputs\":[{\"name\":\"newDenebForkTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slasher\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractISlasher\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stake\",\"inputs\":[{\"name\":\"pubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"depositDataRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"strategyManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateBeaconChainOracle\",\"inputs\":[{\"name\":\"newBeaconChainOracle\",\"type\":\"address\",\"internalType\":\"contractIBeaconChainOracle\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawSharesAsTokens\",\"inputs\":[{\"name\":\"podOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destination\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"BeaconChainETHDeposited\",\"inputs\":[{\"name\":\"podOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BeaconChainETHWithdrawalCompleted\",\"inputs\":[{\"name\":\"podOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint96\",\"indexed\":false,\"internalType\":\"uint96\"},{\"name\":\"delegatedAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"withdrawalRoot\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BeaconOracleUpdated\",\"inputs\":[{\"name\":\"newOracleAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DenebForkTimestampUpdated\",\"inputs\":[{\"name\":\"newValue\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PodDeployed\",\"inputs\":[{\"name\":\"eigenPod\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"podOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PodSharesUpdated\",\"inputs\":[{\"name\":\"podOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sharesDelta\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// IEigenPodManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use IEigenPodManagerMetaData.ABI instead.
var IEigenPodManagerABI = IEigenPodManagerMetaData.ABI

// IEigenPodManager is an auto generated Go binding around an Ethereum contract.
type IEigenPodManager struct {
	IEigenPodManagerCaller     // Read-only binding to the contract
	IEigenPodManagerTransactor // Write-only binding to the contract
	IEigenPodManagerFilterer   // Log filterer for contract events
}

// IEigenPodManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IEigenPodManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEigenPodManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IEigenPodManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEigenPodManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IEigenPodManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEigenPodManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IEigenPodManagerSession struct {
	Contract     *IEigenPodManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IEigenPodManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IEigenPodManagerCallerSession struct {
	Contract *IEigenPodManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IEigenPodManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IEigenPodManagerTransactorSession struct {
	Contract     *IEigenPodManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IEigenPodManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IEigenPodManagerRaw struct {
	Contract *IEigenPodManager // Generic contract binding to access the raw methods on
}

// IEigenPodManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IEigenPodManagerCallerRaw struct {
	Contract *IEigenPodManagerCaller // Generic read-only contract binding to access the raw methods on
}

// IEigenPodManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IEigenPodManagerTransactorRaw struct {
	Contract *IEigenPodManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIEigenPodManager creates a new instance of IEigenPodManager, bound to a specific deployed contract.
func NewIEigenPodManager(address common.Address, backend bind.ContractBackend) (*IEigenPodManager, error) {
	contract, err := bindIEigenPodManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IEigenPodManager{IEigenPodManagerCaller: IEigenPodManagerCaller{contract: contract}, IEigenPodManagerTransactor: IEigenPodManagerTransactor{contract: contract}, IEigenPodManagerFilterer: IEigenPodManagerFilterer{contract: contract}}, nil
}

// NewIEigenPodManagerCaller creates a new read-only instance of IEigenPodManager, bound to a specific deployed contract.
func NewIEigenPodManagerCaller(address common.Address, caller bind.ContractCaller) (*IEigenPodManagerCaller, error) {
	contract, err := bindIEigenPodManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IEigenPodManagerCaller{contract: contract}, nil
}

// NewIEigenPodManagerTransactor creates a new write-only instance of IEigenPodManager, bound to a specific deployed contract.
func NewIEigenPodManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IEigenPodManagerTransactor, error) {
	contract, err := bindIEigenPodManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IEigenPodManagerTransactor{contract: contract}, nil
}

// NewIEigenPodManagerFilterer creates a new log filterer instance of IEigenPodManager, bound to a specific deployed contract.
func NewIEigenPodManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IEigenPodManagerFilterer, error) {
	contract, err := bindIEigenPodManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IEigenPodManagerFilterer{contract: contract}, nil
}

// bindIEigenPodManager binds a generic wrapper to an already deployed contract.
func bindIEigenPodManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IEigenPodManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEigenPodManager *IEigenPodManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IEigenPodManager.Contract.IEigenPodManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEigenPodManager *IEigenPodManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEigenPodManager.Contract.IEigenPodManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEigenPodManager *IEigenPodManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEigenPodManager.Contract.IEigenPodManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEigenPodManager *IEigenPodManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IEigenPodManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEigenPodManager *IEigenPodManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEigenPodManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEigenPodManager *IEigenPodManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEigenPodManager.Contract.contract.Transact(opts, method, params...)
}

// BeaconChainETHStrategy is a free data retrieval call binding the contract method 0x9104c319.
//
// Solidity: function beaconChainETHStrategy() view returns(address)
func (_IEigenPodManager *IEigenPodManagerCaller) BeaconChainETHStrategy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IEigenPodManager.contract.Call(opts, &out, "beaconChainETHStrategy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BeaconChainETHStrategy is a free data retrieval call binding the contract method 0x9104c319.
//
// Solidity: function beaconChainETHStrategy() view returns(address)
func (_IEigenPodManager *IEigenPodManagerSession) BeaconChainETHStrategy() (common.Address, error) {
	return _IEigenPodManager.Contract.BeaconChainETHStrategy(&_IEigenPodManager.CallOpts)
}

// BeaconChainETHStrategy is a free data retrieval call binding the contract method 0x9104c319.
//
// Solidity: function beaconChainETHStrategy() view returns(address)
func (_IEigenPodManager *IEigenPodManagerCallerSession) BeaconChainETHStrategy() (common.Address, error) {
	return _IEigenPodManager.Contract.BeaconChainETHStrategy(&_IEigenPodManager.CallOpts)
}

// BeaconChainOracle is a free data retrieval call binding the contract method 0xc052bd61.
//
// Solidity: function beaconChainOracle() view returns(address)
func (_IEigenPodManager *IEigenPodManagerCaller) BeaconChainOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IEigenPodManager.contract.Call(opts, &out, "beaconChainOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BeaconChainOracle is a free data retrieval call binding the contract method 0xc052bd61.
//
// Solidity: function beaconChainOracle() view returns(address)
func (_IEigenPodManager *IEigenPodManagerSession) BeaconChainOracle() (common.Address, error) {
	return _IEigenPodManager.Contract.BeaconChainOracle(&_IEigenPodManager.CallOpts)
}

// BeaconChainOracle is a free data retrieval call binding the contract method 0xc052bd61.
//
// Solidity: function beaconChainOracle() view returns(address)
func (_IEigenPodManager *IEigenPodManagerCallerSession) BeaconChainOracle() (common.Address, error) {
	return _IEigenPodManager.Contract.BeaconChainOracle(&_IEigenPodManager.CallOpts)
}

// DenebForkTimestamp is a free data retrieval call binding the contract method 0x44e71c80.
//
// Solidity: function denebForkTimestamp() view returns(uint64)
func (_IEigenPodManager *IEigenPodManagerCaller) DenebForkTimestamp(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _IEigenPodManager.contract.Call(opts, &out, "denebForkTimestamp")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// DenebForkTimestamp is a free data retrieval call binding the contract method 0x44e71c80.
//
// Solidity: function denebForkTimestamp() view returns(uint64)
func (_IEigenPodManager *IEigenPodManagerSession) DenebForkTimestamp() (uint64, error) {
	return _IEigenPodManager.Contract.DenebForkTimestamp(&_IEigenPodManager.CallOpts)
}

// DenebForkTimestamp is a free data retrieval call binding the contract method 0x44e71c80.
//
// Solidity: function denebForkTimestamp() view returns(uint64)
func (_IEigenPodManager *IEigenPodManagerCallerSession) DenebForkTimestamp() (uint64, error) {
	return _IEigenPodManager.Contract.DenebForkTimestamp(&_IEigenPodManager.CallOpts)
}

// EigenPodBeacon is a free data retrieval call binding the contract method 0x292b7b2b.
//
// Solidity: function eigenPodBeacon() view returns(address)
func (_IEigenPodManager *IEigenPodManagerCaller) EigenPodBeacon(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IEigenPodManager.contract.Call(opts, &out, "eigenPodBeacon")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EigenPodBeacon is a free data retrieval call binding the contract method 0x292b7b2b.
//
// Solidity: function eigenPodBeacon() view returns(address)
func (_IEigenPodManager *IEigenPodManagerSession) EigenPodBeacon() (common.Address, error) {
	return _IEigenPodManager.Contract.EigenPodBeacon(&_IEigenPodManager.CallOpts)
}

// EigenPodBeacon is a free data retrieval call binding the contract method 0x292b7b2b.
//
// Solidity: function eigenPodBeacon() view returns(address)
func (_IEigenPodManager *IEigenPodManagerCallerSession) EigenPodBeacon() (common.Address, error) {
	return _IEigenPodManager.Contract.EigenPodBeacon(&_IEigenPodManager.CallOpts)
}

// EthPOS is a free data retrieval call binding the contract method 0x74cdd798.
//
// Solidity: function ethPOS() view returns(address)
func (_IEigenPodManager *IEigenPodManagerCaller) EthPOS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IEigenPodManager.contract.Call(opts, &out, "ethPOS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EthPOS is a free data retrieval call binding the contract method 0x74cdd798.
//
// Solidity: function ethPOS() view returns(address)
func (_IEigenPodManager *IEigenPodManagerSession) EthPOS() (common.Address, error) {
	return _IEigenPodManager.Contract.EthPOS(&_IEigenPodManager.CallOpts)
}

// EthPOS is a free data retrieval call binding the contract method 0x74cdd798.
//
// Solidity: function ethPOS() view returns(address)
func (_IEigenPodManager *IEigenPodManagerCallerSession) EthPOS() (common.Address, error) {
	return _IEigenPodManager.Contract.EthPOS(&_IEigenPodManager.CallOpts)
}

// GetBlockRootAtTimestamp is a free data retrieval call binding the contract method 0xd1c64cc9.
//
// Solidity: function getBlockRootAtTimestamp(uint64 timestamp) view returns(bytes32)
func (_IEigenPodManager *IEigenPodManagerCaller) GetBlockRootAtTimestamp(opts *bind.CallOpts, timestamp uint64) ([32]byte, error) {
	var out []interface{}
	err := _IEigenPodManager.contract.Call(opts, &out, "getBlockRootAtTimestamp", timestamp)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBlockRootAtTimestamp is a free data retrieval call binding the contract method 0xd1c64cc9.
//
// Solidity: function getBlockRootAtTimestamp(uint64 timestamp) view returns(bytes32)
func (_IEigenPodManager *IEigenPodManagerSession) GetBlockRootAtTimestamp(timestamp uint64) ([32]byte, error) {
	return _IEigenPodManager.Contract.GetBlockRootAtTimestamp(&_IEigenPodManager.CallOpts, timestamp)
}

// GetBlockRootAtTimestamp is a free data retrieval call binding the contract method 0xd1c64cc9.
//
// Solidity: function getBlockRootAtTimestamp(uint64 timestamp) view returns(bytes32)
func (_IEigenPodManager *IEigenPodManagerCallerSession) GetBlockRootAtTimestamp(timestamp uint64) ([32]byte, error) {
	return _IEigenPodManager.Contract.GetBlockRootAtTimestamp(&_IEigenPodManager.CallOpts, timestamp)
}

// GetPod is a free data retrieval call binding the contract method 0xa38406a3.
//
// Solidity: function getPod(address podOwner) view returns(address)
func (_IEigenPodManager *IEigenPodManagerCaller) GetPod(opts *bind.CallOpts, podOwner common.Address) (common.Address, error) {
	var out []interface{}
	err := _IEigenPodManager.contract.Call(opts, &out, "getPod", podOwner)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPod is a free data retrieval call binding the contract method 0xa38406a3.
//
// Solidity: function getPod(address podOwner) view returns(address)
func (_IEigenPodManager *IEigenPodManagerSession) GetPod(podOwner common.Address) (common.Address, error) {
	return _IEigenPodManager.Contract.GetPod(&_IEigenPodManager.CallOpts, podOwner)
}

// GetPod is a free data retrieval call binding the contract method 0xa38406a3.
//
// Solidity: function getPod(address podOwner) view returns(address)
func (_IEigenPodManager *IEigenPodManagerCallerSession) GetPod(podOwner common.Address) (common.Address, error) {
	return _IEigenPodManager.Contract.GetPod(&_IEigenPodManager.CallOpts, podOwner)
}

// HasPod is a free data retrieval call binding the contract method 0xf6848d24.
//
// Solidity: function hasPod(address podOwner) view returns(bool)
func (_IEigenPodManager *IEigenPodManagerCaller) HasPod(opts *bind.CallOpts, podOwner common.Address) (bool, error) {
	var out []interface{}
	err := _IEigenPodManager.contract.Call(opts, &out, "hasPod", podOwner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasPod is a free data retrieval call binding the contract method 0xf6848d24.
//
// Solidity: function hasPod(address podOwner) view returns(bool)
func (_IEigenPodManager *IEigenPodManagerSession) HasPod(podOwner common.Address) (bool, error) {
	return _IEigenPodManager.Contract.HasPod(&_IEigenPodManager.CallOpts, podOwner)
}

// HasPod is a free data retrieval call binding the contract method 0xf6848d24.
//
// Solidity: function hasPod(address podOwner) view returns(bool)
func (_IEigenPodManager *IEigenPodManagerCallerSession) HasPod(podOwner common.Address) (bool, error) {
	return _IEigenPodManager.Contract.HasPod(&_IEigenPodManager.CallOpts, podOwner)
}

// NumPods is a free data retrieval call binding the contract method 0xa6a509be.
//
// Solidity: function numPods() view returns(uint256)
func (_IEigenPodManager *IEigenPodManagerCaller) NumPods(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IEigenPodManager.contract.Call(opts, &out, "numPods")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumPods is a free data retrieval call binding the contract method 0xa6a509be.
//
// Solidity: function numPods() view returns(uint256)
func (_IEigenPodManager *IEigenPodManagerSession) NumPods() (*big.Int, error) {
	return _IEigenPodManager.Contract.NumPods(&_IEigenPodManager.CallOpts)
}

// NumPods is a free data retrieval call binding the contract method 0xa6a509be.
//
// Solidity: function numPods() view returns(uint256)
func (_IEigenPodManager *IEigenPodManagerCallerSession) NumPods() (*big.Int, error) {
	return _IEigenPodManager.Contract.NumPods(&_IEigenPodManager.CallOpts)
}

// OwnerToPod is a free data retrieval call binding the contract method 0x9ba06275.
//
// Solidity: function ownerToPod(address podOwner) view returns(address)
func (_IEigenPodManager *IEigenPodManagerCaller) OwnerToPod(opts *bind.CallOpts, podOwner common.Address) (common.Address, error) {
	var out []interface{}
	err := _IEigenPodManager.contract.Call(opts, &out, "ownerToPod", podOwner)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerToPod is a free data retrieval call binding the contract method 0x9ba06275.
//
// Solidity: function ownerToPod(address podOwner) view returns(address)
func (_IEigenPodManager *IEigenPodManagerSession) OwnerToPod(podOwner common.Address) (common.Address, error) {
	return _IEigenPodManager.Contract.OwnerToPod(&_IEigenPodManager.CallOpts, podOwner)
}

// OwnerToPod is a free data retrieval call binding the contract method 0x9ba06275.
//
// Solidity: function ownerToPod(address podOwner) view returns(address)
func (_IEigenPodManager *IEigenPodManagerCallerSession) OwnerToPod(podOwner common.Address) (common.Address, error) {
	return _IEigenPodManager.Contract.OwnerToPod(&_IEigenPodManager.CallOpts, podOwner)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_IEigenPodManager *IEigenPodManagerCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _IEigenPodManager.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_IEigenPodManager *IEigenPodManagerSession) Paused(index uint8) (bool, error) {
	return _IEigenPodManager.Contract.Paused(&_IEigenPodManager.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_IEigenPodManager *IEigenPodManagerCallerSession) Paused(index uint8) (bool, error) {
	return _IEigenPodManager.Contract.Paused(&_IEigenPodManager.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_IEigenPodManager *IEigenPodManagerCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IEigenPodManager.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_IEigenPodManager *IEigenPodManagerSession) Paused0() (*big.Int, error) {
	return _IEigenPodManager.Contract.Paused0(&_IEigenPodManager.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_IEigenPodManager *IEigenPodManagerCallerSession) Paused0() (*big.Int, error) {
	return _IEigenPodManager.Contract.Paused0(&_IEigenPodManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_IEigenPodManager *IEigenPodManagerCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IEigenPodManager.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_IEigenPodManager *IEigenPodManagerSession) PauserRegistry() (common.Address, error) {
	return _IEigenPodManager.Contract.PauserRegistry(&_IEigenPodManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_IEigenPodManager *IEigenPodManagerCallerSession) PauserRegistry() (common.Address, error) {
	return _IEigenPodManager.Contract.PauserRegistry(&_IEigenPodManager.CallOpts)
}

// PodOwnerShares is a free data retrieval call binding the contract method 0x60f4062b.
//
// Solidity: function podOwnerShares(address podOwner) view returns(int256)
func (_IEigenPodManager *IEigenPodManagerCaller) PodOwnerShares(opts *bind.CallOpts, podOwner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IEigenPodManager.contract.Call(opts, &out, "podOwnerShares", podOwner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PodOwnerShares is a free data retrieval call binding the contract method 0x60f4062b.
//
// Solidity: function podOwnerShares(address podOwner) view returns(int256)
func (_IEigenPodManager *IEigenPodManagerSession) PodOwnerShares(podOwner common.Address) (*big.Int, error) {
	return _IEigenPodManager.Contract.PodOwnerShares(&_IEigenPodManager.CallOpts, podOwner)
}

// PodOwnerShares is a free data retrieval call binding the contract method 0x60f4062b.
//
// Solidity: function podOwnerShares(address podOwner) view returns(int256)
func (_IEigenPodManager *IEigenPodManagerCallerSession) PodOwnerShares(podOwner common.Address) (*big.Int, error) {
	return _IEigenPodManager.Contract.PodOwnerShares(&_IEigenPodManager.CallOpts, podOwner)
}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_IEigenPodManager *IEigenPodManagerCaller) Slasher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IEigenPodManager.contract.Call(opts, &out, "slasher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_IEigenPodManager *IEigenPodManagerSession) Slasher() (common.Address, error) {
	return _IEigenPodManager.Contract.Slasher(&_IEigenPodManager.CallOpts)
}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_IEigenPodManager *IEigenPodManagerCallerSession) Slasher() (common.Address, error) {
	return _IEigenPodManager.Contract.Slasher(&_IEigenPodManager.CallOpts)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_IEigenPodManager *IEigenPodManagerCaller) StrategyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IEigenPodManager.contract.Call(opts, &out, "strategyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_IEigenPodManager *IEigenPodManagerSession) StrategyManager() (common.Address, error) {
	return _IEigenPodManager.Contract.StrategyManager(&_IEigenPodManager.CallOpts)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_IEigenPodManager *IEigenPodManagerCallerSession) StrategyManager() (common.Address, error) {
	return _IEigenPodManager.Contract.StrategyManager(&_IEigenPodManager.CallOpts)
}

// AddShares is a paid mutator transaction binding the contract method 0x0e81073c.
//
// Solidity: function addShares(address podOwner, uint256 shares) returns(uint256)
func (_IEigenPodManager *IEigenPodManagerTransactor) AddShares(opts *bind.TransactOpts, podOwner common.Address, shares *big.Int) (*types.Transaction, error) {
	return _IEigenPodManager.contract.Transact(opts, "addShares", podOwner, shares)
}

// AddShares is a paid mutator transaction binding the contract method 0x0e81073c.
//
// Solidity: function addShares(address podOwner, uint256 shares) returns(uint256)
func (_IEigenPodManager *IEigenPodManagerSession) AddShares(podOwner common.Address, shares *big.Int) (*types.Transaction, error) {
	return _IEigenPodManager.Contract.AddShares(&_IEigenPodManager.TransactOpts, podOwner, shares)
}

// AddShares is a paid mutator transaction binding the contract method 0x0e81073c.
//
// Solidity: function addShares(address podOwner, uint256 shares) returns(uint256)
func (_IEigenPodManager *IEigenPodManagerTransactorSession) AddShares(podOwner common.Address, shares *big.Int) (*types.Transaction, error) {
	return _IEigenPodManager.Contract.AddShares(&_IEigenPodManager.TransactOpts, podOwner, shares)
}

// CreatePod is a paid mutator transaction binding the contract method 0x84d81062.
//
// Solidity: function createPod() returns(address)
func (_IEigenPodManager *IEigenPodManagerTransactor) CreatePod(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEigenPodManager.contract.Transact(opts, "createPod")
}

// CreatePod is a paid mutator transaction binding the contract method 0x84d81062.
//
// Solidity: function createPod() returns(address)
func (_IEigenPodManager *IEigenPodManagerSession) CreatePod() (*types.Transaction, error) {
	return _IEigenPodManager.Contract.CreatePod(&_IEigenPodManager.TransactOpts)
}

// CreatePod is a paid mutator transaction binding the contract method 0x84d81062.
//
// Solidity: function createPod() returns(address)
func (_IEigenPodManager *IEigenPodManagerTransactorSession) CreatePod() (*types.Transaction, error) {
	return _IEigenPodManager.Contract.CreatePod(&_IEigenPodManager.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_IEigenPodManager *IEigenPodManagerTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _IEigenPodManager.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_IEigenPodManager *IEigenPodManagerSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _IEigenPodManager.Contract.Pause(&_IEigenPodManager.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_IEigenPodManager *IEigenPodManagerTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _IEigenPodManager.Contract.Pause(&_IEigenPodManager.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_IEigenPodManager *IEigenPodManagerTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEigenPodManager.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_IEigenPodManager *IEigenPodManagerSession) PauseAll() (*types.Transaction, error) {
	return _IEigenPodManager.Contract.PauseAll(&_IEigenPodManager.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_IEigenPodManager *IEigenPodManagerTransactorSession) PauseAll() (*types.Transaction, error) {
	return _IEigenPodManager.Contract.PauseAll(&_IEigenPodManager.TransactOpts)
}

// RecordBeaconChainETHBalanceUpdate is a paid mutator transaction binding the contract method 0xc2c51c40.
//
// Solidity: function recordBeaconChainETHBalanceUpdate(address podOwner, int256 sharesDelta) returns()
func (_IEigenPodManager *IEigenPodManagerTransactor) RecordBeaconChainETHBalanceUpdate(opts *bind.TransactOpts, podOwner common.Address, sharesDelta *big.Int) (*types.Transaction, error) {
	return _IEigenPodManager.contract.Transact(opts, "recordBeaconChainETHBalanceUpdate", podOwner, sharesDelta)
}

// RecordBeaconChainETHBalanceUpdate is a paid mutator transaction binding the contract method 0xc2c51c40.
//
// Solidity: function recordBeaconChainETHBalanceUpdate(address podOwner, int256 sharesDelta) returns()
func (_IEigenPodManager *IEigenPodManagerSession) RecordBeaconChainETHBalanceUpdate(podOwner common.Address, sharesDelta *big.Int) (*types.Transaction, error) {
	return _IEigenPodManager.Contract.RecordBeaconChainETHBalanceUpdate(&_IEigenPodManager.TransactOpts, podOwner, sharesDelta)
}

// RecordBeaconChainETHBalanceUpdate is a paid mutator transaction binding the contract method 0xc2c51c40.
//
// Solidity: function recordBeaconChainETHBalanceUpdate(address podOwner, int256 sharesDelta) returns()
func (_IEigenPodManager *IEigenPodManagerTransactorSession) RecordBeaconChainETHBalanceUpdate(podOwner common.Address, sharesDelta *big.Int) (*types.Transaction, error) {
	return _IEigenPodManager.Contract.RecordBeaconChainETHBalanceUpdate(&_IEigenPodManager.TransactOpts, podOwner, sharesDelta)
}

// RemoveShares is a paid mutator transaction binding the contract method 0xbeffbb89.
//
// Solidity: function removeShares(address podOwner, uint256 shares) returns()
func (_IEigenPodManager *IEigenPodManagerTransactor) RemoveShares(opts *bind.TransactOpts, podOwner common.Address, shares *big.Int) (*types.Transaction, error) {
	return _IEigenPodManager.contract.Transact(opts, "removeShares", podOwner, shares)
}

// RemoveShares is a paid mutator transaction binding the contract method 0xbeffbb89.
//
// Solidity: function removeShares(address podOwner, uint256 shares) returns()
func (_IEigenPodManager *IEigenPodManagerSession) RemoveShares(podOwner common.Address, shares *big.Int) (*types.Transaction, error) {
	return _IEigenPodManager.Contract.RemoveShares(&_IEigenPodManager.TransactOpts, podOwner, shares)
}

// RemoveShares is a paid mutator transaction binding the contract method 0xbeffbb89.
//
// Solidity: function removeShares(address podOwner, uint256 shares) returns()
func (_IEigenPodManager *IEigenPodManagerTransactorSession) RemoveShares(podOwner common.Address, shares *big.Int) (*types.Transaction, error) {
	return _IEigenPodManager.Contract.RemoveShares(&_IEigenPodManager.TransactOpts, podOwner, shares)
}

// SetDenebForkTimestamp is a paid mutator transaction binding the contract method 0x463db038.
//
// Solidity: function setDenebForkTimestamp(uint64 newDenebForkTimestamp) returns()
func (_IEigenPodManager *IEigenPodManagerTransactor) SetDenebForkTimestamp(opts *bind.TransactOpts, newDenebForkTimestamp uint64) (*types.Transaction, error) {
	return _IEigenPodManager.contract.Transact(opts, "setDenebForkTimestamp", newDenebForkTimestamp)
}

// SetDenebForkTimestamp is a paid mutator transaction binding the contract method 0x463db038.
//
// Solidity: function setDenebForkTimestamp(uint64 newDenebForkTimestamp) returns()
func (_IEigenPodManager *IEigenPodManagerSession) SetDenebForkTimestamp(newDenebForkTimestamp uint64) (*types.Transaction, error) {
	return _IEigenPodManager.Contract.SetDenebForkTimestamp(&_IEigenPodManager.TransactOpts, newDenebForkTimestamp)
}

// SetDenebForkTimestamp is a paid mutator transaction binding the contract method 0x463db038.
//
// Solidity: function setDenebForkTimestamp(uint64 newDenebForkTimestamp) returns()
func (_IEigenPodManager *IEigenPodManagerTransactorSession) SetDenebForkTimestamp(newDenebForkTimestamp uint64) (*types.Transaction, error) {
	return _IEigenPodManager.Contract.SetDenebForkTimestamp(&_IEigenPodManager.TransactOpts, newDenebForkTimestamp)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_IEigenPodManager *IEigenPodManagerTransactor) SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error) {
	return _IEigenPodManager.contract.Transact(opts, "setPauserRegistry", newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_IEigenPodManager *IEigenPodManagerSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _IEigenPodManager.Contract.SetPauserRegistry(&_IEigenPodManager.TransactOpts, newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_IEigenPodManager *IEigenPodManagerTransactorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _IEigenPodManager.Contract.SetPauserRegistry(&_IEigenPodManager.TransactOpts, newPauserRegistry)
}

// Stake is a paid mutator transaction binding the contract method 0x9b4e4634.
//
// Solidity: function stake(bytes pubkey, bytes signature, bytes32 depositDataRoot) payable returns()
func (_IEigenPodManager *IEigenPodManagerTransactor) Stake(opts *bind.TransactOpts, pubkey []byte, signature []byte, depositDataRoot [32]byte) (*types.Transaction, error) {
	return _IEigenPodManager.contract.Transact(opts, "stake", pubkey, signature, depositDataRoot)
}

// Stake is a paid mutator transaction binding the contract method 0x9b4e4634.
//
// Solidity: function stake(bytes pubkey, bytes signature, bytes32 depositDataRoot) payable returns()
func (_IEigenPodManager *IEigenPodManagerSession) Stake(pubkey []byte, signature []byte, depositDataRoot [32]byte) (*types.Transaction, error) {
	return _IEigenPodManager.Contract.Stake(&_IEigenPodManager.TransactOpts, pubkey, signature, depositDataRoot)
}

// Stake is a paid mutator transaction binding the contract method 0x9b4e4634.
//
// Solidity: function stake(bytes pubkey, bytes signature, bytes32 depositDataRoot) payable returns()
func (_IEigenPodManager *IEigenPodManagerTransactorSession) Stake(pubkey []byte, signature []byte, depositDataRoot [32]byte) (*types.Transaction, error) {
	return _IEigenPodManager.Contract.Stake(&_IEigenPodManager.TransactOpts, pubkey, signature, depositDataRoot)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_IEigenPodManager *IEigenPodManagerTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _IEigenPodManager.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_IEigenPodManager *IEigenPodManagerSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _IEigenPodManager.Contract.Unpause(&_IEigenPodManager.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_IEigenPodManager *IEigenPodManagerTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _IEigenPodManager.Contract.Unpause(&_IEigenPodManager.TransactOpts, newPausedStatus)
}

// UpdateBeaconChainOracle is a paid mutator transaction binding the contract method 0xc1de3aef.
//
// Solidity: function updateBeaconChainOracle(address newBeaconChainOracle) returns()
func (_IEigenPodManager *IEigenPodManagerTransactor) UpdateBeaconChainOracle(opts *bind.TransactOpts, newBeaconChainOracle common.Address) (*types.Transaction, error) {
	return _IEigenPodManager.contract.Transact(opts, "updateBeaconChainOracle", newBeaconChainOracle)
}

// UpdateBeaconChainOracle is a paid mutator transaction binding the contract method 0xc1de3aef.
//
// Solidity: function updateBeaconChainOracle(address newBeaconChainOracle) returns()
func (_IEigenPodManager *IEigenPodManagerSession) UpdateBeaconChainOracle(newBeaconChainOracle common.Address) (*types.Transaction, error) {
	return _IEigenPodManager.Contract.UpdateBeaconChainOracle(&_IEigenPodManager.TransactOpts, newBeaconChainOracle)
}

// UpdateBeaconChainOracle is a paid mutator transaction binding the contract method 0xc1de3aef.
//
// Solidity: function updateBeaconChainOracle(address newBeaconChainOracle) returns()
func (_IEigenPodManager *IEigenPodManagerTransactorSession) UpdateBeaconChainOracle(newBeaconChainOracle common.Address) (*types.Transaction, error) {
	return _IEigenPodManager.Contract.UpdateBeaconChainOracle(&_IEigenPodManager.TransactOpts, newBeaconChainOracle)
}

// WithdrawSharesAsTokens is a paid mutator transaction binding the contract method 0x387b1300.
//
// Solidity: function withdrawSharesAsTokens(address podOwner, address destination, uint256 shares) returns()
func (_IEigenPodManager *IEigenPodManagerTransactor) WithdrawSharesAsTokens(opts *bind.TransactOpts, podOwner common.Address, destination common.Address, shares *big.Int) (*types.Transaction, error) {
	return _IEigenPodManager.contract.Transact(opts, "withdrawSharesAsTokens", podOwner, destination, shares)
}

// WithdrawSharesAsTokens is a paid mutator transaction binding the contract method 0x387b1300.
//
// Solidity: function withdrawSharesAsTokens(address podOwner, address destination, uint256 shares) returns()
func (_IEigenPodManager *IEigenPodManagerSession) WithdrawSharesAsTokens(podOwner common.Address, destination common.Address, shares *big.Int) (*types.Transaction, error) {
	return _IEigenPodManager.Contract.WithdrawSharesAsTokens(&_IEigenPodManager.TransactOpts, podOwner, destination, shares)
}

// WithdrawSharesAsTokens is a paid mutator transaction binding the contract method 0x387b1300.
//
// Solidity: function withdrawSharesAsTokens(address podOwner, address destination, uint256 shares) returns()
func (_IEigenPodManager *IEigenPodManagerTransactorSession) WithdrawSharesAsTokens(podOwner common.Address, destination common.Address, shares *big.Int) (*types.Transaction, error) {
	return _IEigenPodManager.Contract.WithdrawSharesAsTokens(&_IEigenPodManager.TransactOpts, podOwner, destination, shares)
}

// IEigenPodManagerBeaconChainETHDepositedIterator is returned from FilterBeaconChainETHDeposited and is used to iterate over the raw logs and unpacked data for BeaconChainETHDeposited events raised by the IEigenPodManager contract.
type IEigenPodManagerBeaconChainETHDepositedIterator struct {
	Event *IEigenPodManagerBeaconChainETHDeposited // Event containing the contract specifics and raw log

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
func (it *IEigenPodManagerBeaconChainETHDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEigenPodManagerBeaconChainETHDeposited)
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
		it.Event = new(IEigenPodManagerBeaconChainETHDeposited)
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
func (it *IEigenPodManagerBeaconChainETHDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEigenPodManagerBeaconChainETHDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEigenPodManagerBeaconChainETHDeposited represents a BeaconChainETHDeposited event raised by the IEigenPodManager contract.
type IEigenPodManagerBeaconChainETHDeposited struct {
	PodOwner common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBeaconChainETHDeposited is a free log retrieval operation binding the contract event 0x35a85cabc603f48abb2b71d9fbd8adea7c449d7f0be900ae7a2986ea369c3d0d.
//
// Solidity: event BeaconChainETHDeposited(address indexed podOwner, uint256 amount)
func (_IEigenPodManager *IEigenPodManagerFilterer) FilterBeaconChainETHDeposited(opts *bind.FilterOpts, podOwner []common.Address) (*IEigenPodManagerBeaconChainETHDepositedIterator, error) {

	var podOwnerRule []interface{}
	for _, podOwnerItem := range podOwner {
		podOwnerRule = append(podOwnerRule, podOwnerItem)
	}

	logs, sub, err := _IEigenPodManager.contract.FilterLogs(opts, "BeaconChainETHDeposited", podOwnerRule)
	if err != nil {
		return nil, err
	}
	return &IEigenPodManagerBeaconChainETHDepositedIterator{contract: _IEigenPodManager.contract, event: "BeaconChainETHDeposited", logs: logs, sub: sub}, nil
}

// WatchBeaconChainETHDeposited is a free log subscription operation binding the contract event 0x35a85cabc603f48abb2b71d9fbd8adea7c449d7f0be900ae7a2986ea369c3d0d.
//
// Solidity: event BeaconChainETHDeposited(address indexed podOwner, uint256 amount)
func (_IEigenPodManager *IEigenPodManagerFilterer) WatchBeaconChainETHDeposited(opts *bind.WatchOpts, sink chan<- *IEigenPodManagerBeaconChainETHDeposited, podOwner []common.Address) (event.Subscription, error) {

	var podOwnerRule []interface{}
	for _, podOwnerItem := range podOwner {
		podOwnerRule = append(podOwnerRule, podOwnerItem)
	}

	logs, sub, err := _IEigenPodManager.contract.WatchLogs(opts, "BeaconChainETHDeposited", podOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEigenPodManagerBeaconChainETHDeposited)
				if err := _IEigenPodManager.contract.UnpackLog(event, "BeaconChainETHDeposited", log); err != nil {
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
func (_IEigenPodManager *IEigenPodManagerFilterer) ParseBeaconChainETHDeposited(log types.Log) (*IEigenPodManagerBeaconChainETHDeposited, error) {
	event := new(IEigenPodManagerBeaconChainETHDeposited)
	if err := _IEigenPodManager.contract.UnpackLog(event, "BeaconChainETHDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IEigenPodManagerBeaconChainETHWithdrawalCompletedIterator is returned from FilterBeaconChainETHWithdrawalCompleted and is used to iterate over the raw logs and unpacked data for BeaconChainETHWithdrawalCompleted events raised by the IEigenPodManager contract.
type IEigenPodManagerBeaconChainETHWithdrawalCompletedIterator struct {
	Event *IEigenPodManagerBeaconChainETHWithdrawalCompleted // Event containing the contract specifics and raw log

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
func (it *IEigenPodManagerBeaconChainETHWithdrawalCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEigenPodManagerBeaconChainETHWithdrawalCompleted)
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
		it.Event = new(IEigenPodManagerBeaconChainETHWithdrawalCompleted)
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
func (it *IEigenPodManagerBeaconChainETHWithdrawalCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEigenPodManagerBeaconChainETHWithdrawalCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEigenPodManagerBeaconChainETHWithdrawalCompleted represents a BeaconChainETHWithdrawalCompleted event raised by the IEigenPodManager contract.
type IEigenPodManagerBeaconChainETHWithdrawalCompleted struct {
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
func (_IEigenPodManager *IEigenPodManagerFilterer) FilterBeaconChainETHWithdrawalCompleted(opts *bind.FilterOpts, podOwner []common.Address) (*IEigenPodManagerBeaconChainETHWithdrawalCompletedIterator, error) {

	var podOwnerRule []interface{}
	for _, podOwnerItem := range podOwner {
		podOwnerRule = append(podOwnerRule, podOwnerItem)
	}

	logs, sub, err := _IEigenPodManager.contract.FilterLogs(opts, "BeaconChainETHWithdrawalCompleted", podOwnerRule)
	if err != nil {
		return nil, err
	}
	return &IEigenPodManagerBeaconChainETHWithdrawalCompletedIterator{contract: _IEigenPodManager.contract, event: "BeaconChainETHWithdrawalCompleted", logs: logs, sub: sub}, nil
}

// WatchBeaconChainETHWithdrawalCompleted is a free log subscription operation binding the contract event 0xa6bab1d55a361fcea2eee2bc9491e4f01e6cf333df03c9c4f2c144466429f7d6.
//
// Solidity: event BeaconChainETHWithdrawalCompleted(address indexed podOwner, uint256 shares, uint96 nonce, address delegatedAddress, address withdrawer, bytes32 withdrawalRoot)
func (_IEigenPodManager *IEigenPodManagerFilterer) WatchBeaconChainETHWithdrawalCompleted(opts *bind.WatchOpts, sink chan<- *IEigenPodManagerBeaconChainETHWithdrawalCompleted, podOwner []common.Address) (event.Subscription, error) {

	var podOwnerRule []interface{}
	for _, podOwnerItem := range podOwner {
		podOwnerRule = append(podOwnerRule, podOwnerItem)
	}

	logs, sub, err := _IEigenPodManager.contract.WatchLogs(opts, "BeaconChainETHWithdrawalCompleted", podOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEigenPodManagerBeaconChainETHWithdrawalCompleted)
				if err := _IEigenPodManager.contract.UnpackLog(event, "BeaconChainETHWithdrawalCompleted", log); err != nil {
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
func (_IEigenPodManager *IEigenPodManagerFilterer) ParseBeaconChainETHWithdrawalCompleted(log types.Log) (*IEigenPodManagerBeaconChainETHWithdrawalCompleted, error) {
	event := new(IEigenPodManagerBeaconChainETHWithdrawalCompleted)
	if err := _IEigenPodManager.contract.UnpackLog(event, "BeaconChainETHWithdrawalCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IEigenPodManagerBeaconOracleUpdatedIterator is returned from FilterBeaconOracleUpdated and is used to iterate over the raw logs and unpacked data for BeaconOracleUpdated events raised by the IEigenPodManager contract.
type IEigenPodManagerBeaconOracleUpdatedIterator struct {
	Event *IEigenPodManagerBeaconOracleUpdated // Event containing the contract specifics and raw log

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
func (it *IEigenPodManagerBeaconOracleUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEigenPodManagerBeaconOracleUpdated)
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
		it.Event = new(IEigenPodManagerBeaconOracleUpdated)
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
func (it *IEigenPodManagerBeaconOracleUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEigenPodManagerBeaconOracleUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEigenPodManagerBeaconOracleUpdated represents a BeaconOracleUpdated event raised by the IEigenPodManager contract.
type IEigenPodManagerBeaconOracleUpdated struct {
	NewOracleAddress common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterBeaconOracleUpdated is a free log retrieval operation binding the contract event 0x08f0470754946ccfbb446ff7fd2d6ae6af1bbdae19f85794c0cc5ed5e8ceb4f6.
//
// Solidity: event BeaconOracleUpdated(address indexed newOracleAddress)
func (_IEigenPodManager *IEigenPodManagerFilterer) FilterBeaconOracleUpdated(opts *bind.FilterOpts, newOracleAddress []common.Address) (*IEigenPodManagerBeaconOracleUpdatedIterator, error) {

	var newOracleAddressRule []interface{}
	for _, newOracleAddressItem := range newOracleAddress {
		newOracleAddressRule = append(newOracleAddressRule, newOracleAddressItem)
	}

	logs, sub, err := _IEigenPodManager.contract.FilterLogs(opts, "BeaconOracleUpdated", newOracleAddressRule)
	if err != nil {
		return nil, err
	}
	return &IEigenPodManagerBeaconOracleUpdatedIterator{contract: _IEigenPodManager.contract, event: "BeaconOracleUpdated", logs: logs, sub: sub}, nil
}

// WatchBeaconOracleUpdated is a free log subscription operation binding the contract event 0x08f0470754946ccfbb446ff7fd2d6ae6af1bbdae19f85794c0cc5ed5e8ceb4f6.
//
// Solidity: event BeaconOracleUpdated(address indexed newOracleAddress)
func (_IEigenPodManager *IEigenPodManagerFilterer) WatchBeaconOracleUpdated(opts *bind.WatchOpts, sink chan<- *IEigenPodManagerBeaconOracleUpdated, newOracleAddress []common.Address) (event.Subscription, error) {

	var newOracleAddressRule []interface{}
	for _, newOracleAddressItem := range newOracleAddress {
		newOracleAddressRule = append(newOracleAddressRule, newOracleAddressItem)
	}

	logs, sub, err := _IEigenPodManager.contract.WatchLogs(opts, "BeaconOracleUpdated", newOracleAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEigenPodManagerBeaconOracleUpdated)
				if err := _IEigenPodManager.contract.UnpackLog(event, "BeaconOracleUpdated", log); err != nil {
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

// ParseBeaconOracleUpdated is a log parse operation binding the contract event 0x08f0470754946ccfbb446ff7fd2d6ae6af1bbdae19f85794c0cc5ed5e8ceb4f6.
//
// Solidity: event BeaconOracleUpdated(address indexed newOracleAddress)
func (_IEigenPodManager *IEigenPodManagerFilterer) ParseBeaconOracleUpdated(log types.Log) (*IEigenPodManagerBeaconOracleUpdated, error) {
	event := new(IEigenPodManagerBeaconOracleUpdated)
	if err := _IEigenPodManager.contract.UnpackLog(event, "BeaconOracleUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IEigenPodManagerDenebForkTimestampUpdatedIterator is returned from FilterDenebForkTimestampUpdated and is used to iterate over the raw logs and unpacked data for DenebForkTimestampUpdated events raised by the IEigenPodManager contract.
type IEigenPodManagerDenebForkTimestampUpdatedIterator struct {
	Event *IEigenPodManagerDenebForkTimestampUpdated // Event containing the contract specifics and raw log

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
func (it *IEigenPodManagerDenebForkTimestampUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEigenPodManagerDenebForkTimestampUpdated)
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
		it.Event = new(IEigenPodManagerDenebForkTimestampUpdated)
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
func (it *IEigenPodManagerDenebForkTimestampUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEigenPodManagerDenebForkTimestampUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEigenPodManagerDenebForkTimestampUpdated represents a DenebForkTimestampUpdated event raised by the IEigenPodManager contract.
type IEigenPodManagerDenebForkTimestampUpdated struct {
	NewValue uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDenebForkTimestampUpdated is a free log retrieval operation binding the contract event 0x19200b6fdad58f91b2f496b0c444fc4be3eff74a7e24b07770e04a7137bfd9db.
//
// Solidity: event DenebForkTimestampUpdated(uint64 newValue)
func (_IEigenPodManager *IEigenPodManagerFilterer) FilterDenebForkTimestampUpdated(opts *bind.FilterOpts) (*IEigenPodManagerDenebForkTimestampUpdatedIterator, error) {

	logs, sub, err := _IEigenPodManager.contract.FilterLogs(opts, "DenebForkTimestampUpdated")
	if err != nil {
		return nil, err
	}
	return &IEigenPodManagerDenebForkTimestampUpdatedIterator{contract: _IEigenPodManager.contract, event: "DenebForkTimestampUpdated", logs: logs, sub: sub}, nil
}

// WatchDenebForkTimestampUpdated is a free log subscription operation binding the contract event 0x19200b6fdad58f91b2f496b0c444fc4be3eff74a7e24b07770e04a7137bfd9db.
//
// Solidity: event DenebForkTimestampUpdated(uint64 newValue)
func (_IEigenPodManager *IEigenPodManagerFilterer) WatchDenebForkTimestampUpdated(opts *bind.WatchOpts, sink chan<- *IEigenPodManagerDenebForkTimestampUpdated) (event.Subscription, error) {

	logs, sub, err := _IEigenPodManager.contract.WatchLogs(opts, "DenebForkTimestampUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEigenPodManagerDenebForkTimestampUpdated)
				if err := _IEigenPodManager.contract.UnpackLog(event, "DenebForkTimestampUpdated", log); err != nil {
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

// ParseDenebForkTimestampUpdated is a log parse operation binding the contract event 0x19200b6fdad58f91b2f496b0c444fc4be3eff74a7e24b07770e04a7137bfd9db.
//
// Solidity: event DenebForkTimestampUpdated(uint64 newValue)
func (_IEigenPodManager *IEigenPodManagerFilterer) ParseDenebForkTimestampUpdated(log types.Log) (*IEigenPodManagerDenebForkTimestampUpdated, error) {
	event := new(IEigenPodManagerDenebForkTimestampUpdated)
	if err := _IEigenPodManager.contract.UnpackLog(event, "DenebForkTimestampUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IEigenPodManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the IEigenPodManager contract.
type IEigenPodManagerPausedIterator struct {
	Event *IEigenPodManagerPaused // Event containing the contract specifics and raw log

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
func (it *IEigenPodManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEigenPodManagerPaused)
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
		it.Event = new(IEigenPodManagerPaused)
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
func (it *IEigenPodManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEigenPodManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEigenPodManagerPaused represents a Paused event raised by the IEigenPodManager contract.
type IEigenPodManagerPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_IEigenPodManager *IEigenPodManagerFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*IEigenPodManagerPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IEigenPodManager.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &IEigenPodManagerPausedIterator{contract: _IEigenPodManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_IEigenPodManager *IEigenPodManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *IEigenPodManagerPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IEigenPodManager.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEigenPodManagerPaused)
				if err := _IEigenPodManager.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_IEigenPodManager *IEigenPodManagerFilterer) ParsePaused(log types.Log) (*IEigenPodManagerPaused, error) {
	event := new(IEigenPodManagerPaused)
	if err := _IEigenPodManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IEigenPodManagerPauserRegistrySetIterator is returned from FilterPauserRegistrySet and is used to iterate over the raw logs and unpacked data for PauserRegistrySet events raised by the IEigenPodManager contract.
type IEigenPodManagerPauserRegistrySetIterator struct {
	Event *IEigenPodManagerPauserRegistrySet // Event containing the contract specifics and raw log

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
func (it *IEigenPodManagerPauserRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEigenPodManagerPauserRegistrySet)
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
		it.Event = new(IEigenPodManagerPauserRegistrySet)
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
func (it *IEigenPodManagerPauserRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEigenPodManagerPauserRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEigenPodManagerPauserRegistrySet represents a PauserRegistrySet event raised by the IEigenPodManager contract.
type IEigenPodManagerPauserRegistrySet struct {
	PauserRegistry    common.Address
	NewPauserRegistry common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPauserRegistrySet is a free log retrieval operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_IEigenPodManager *IEigenPodManagerFilterer) FilterPauserRegistrySet(opts *bind.FilterOpts) (*IEigenPodManagerPauserRegistrySetIterator, error) {

	logs, sub, err := _IEigenPodManager.contract.FilterLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return &IEigenPodManagerPauserRegistrySetIterator{contract: _IEigenPodManager.contract, event: "PauserRegistrySet", logs: logs, sub: sub}, nil
}

// WatchPauserRegistrySet is a free log subscription operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_IEigenPodManager *IEigenPodManagerFilterer) WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *IEigenPodManagerPauserRegistrySet) (event.Subscription, error) {

	logs, sub, err := _IEigenPodManager.contract.WatchLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEigenPodManagerPauserRegistrySet)
				if err := _IEigenPodManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
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
func (_IEigenPodManager *IEigenPodManagerFilterer) ParsePauserRegistrySet(log types.Log) (*IEigenPodManagerPauserRegistrySet, error) {
	event := new(IEigenPodManagerPauserRegistrySet)
	if err := _IEigenPodManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IEigenPodManagerPodDeployedIterator is returned from FilterPodDeployed and is used to iterate over the raw logs and unpacked data for PodDeployed events raised by the IEigenPodManager contract.
type IEigenPodManagerPodDeployedIterator struct {
	Event *IEigenPodManagerPodDeployed // Event containing the contract specifics and raw log

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
func (it *IEigenPodManagerPodDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEigenPodManagerPodDeployed)
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
		it.Event = new(IEigenPodManagerPodDeployed)
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
func (it *IEigenPodManagerPodDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEigenPodManagerPodDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEigenPodManagerPodDeployed represents a PodDeployed event raised by the IEigenPodManager contract.
type IEigenPodManagerPodDeployed struct {
	EigenPod common.Address
	PodOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPodDeployed is a free log retrieval operation binding the contract event 0x21c99d0db02213c32fff5b05cf0a718ab5f858802b91498f80d82270289d856a.
//
// Solidity: event PodDeployed(address indexed eigenPod, address indexed podOwner)
func (_IEigenPodManager *IEigenPodManagerFilterer) FilterPodDeployed(opts *bind.FilterOpts, eigenPod []common.Address, podOwner []common.Address) (*IEigenPodManagerPodDeployedIterator, error) {

	var eigenPodRule []interface{}
	for _, eigenPodItem := range eigenPod {
		eigenPodRule = append(eigenPodRule, eigenPodItem)
	}
	var podOwnerRule []interface{}
	for _, podOwnerItem := range podOwner {
		podOwnerRule = append(podOwnerRule, podOwnerItem)
	}

	logs, sub, err := _IEigenPodManager.contract.FilterLogs(opts, "PodDeployed", eigenPodRule, podOwnerRule)
	if err != nil {
		return nil, err
	}
	return &IEigenPodManagerPodDeployedIterator{contract: _IEigenPodManager.contract, event: "PodDeployed", logs: logs, sub: sub}, nil
}

// WatchPodDeployed is a free log subscription operation binding the contract event 0x21c99d0db02213c32fff5b05cf0a718ab5f858802b91498f80d82270289d856a.
//
// Solidity: event PodDeployed(address indexed eigenPod, address indexed podOwner)
func (_IEigenPodManager *IEigenPodManagerFilterer) WatchPodDeployed(opts *bind.WatchOpts, sink chan<- *IEigenPodManagerPodDeployed, eigenPod []common.Address, podOwner []common.Address) (event.Subscription, error) {

	var eigenPodRule []interface{}
	for _, eigenPodItem := range eigenPod {
		eigenPodRule = append(eigenPodRule, eigenPodItem)
	}
	var podOwnerRule []interface{}
	for _, podOwnerItem := range podOwner {
		podOwnerRule = append(podOwnerRule, podOwnerItem)
	}

	logs, sub, err := _IEigenPodManager.contract.WatchLogs(opts, "PodDeployed", eigenPodRule, podOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEigenPodManagerPodDeployed)
				if err := _IEigenPodManager.contract.UnpackLog(event, "PodDeployed", log); err != nil {
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
func (_IEigenPodManager *IEigenPodManagerFilterer) ParsePodDeployed(log types.Log) (*IEigenPodManagerPodDeployed, error) {
	event := new(IEigenPodManagerPodDeployed)
	if err := _IEigenPodManager.contract.UnpackLog(event, "PodDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IEigenPodManagerPodSharesUpdatedIterator is returned from FilterPodSharesUpdated and is used to iterate over the raw logs and unpacked data for PodSharesUpdated events raised by the IEigenPodManager contract.
type IEigenPodManagerPodSharesUpdatedIterator struct {
	Event *IEigenPodManagerPodSharesUpdated // Event containing the contract specifics and raw log

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
func (it *IEigenPodManagerPodSharesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEigenPodManagerPodSharesUpdated)
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
		it.Event = new(IEigenPodManagerPodSharesUpdated)
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
func (it *IEigenPodManagerPodSharesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEigenPodManagerPodSharesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEigenPodManagerPodSharesUpdated represents a PodSharesUpdated event raised by the IEigenPodManager contract.
type IEigenPodManagerPodSharesUpdated struct {
	PodOwner    common.Address
	SharesDelta *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPodSharesUpdated is a free log retrieval operation binding the contract event 0x4e2b791dedccd9fb30141b088cabf5c14a8912b52f59375c95c010700b8c6193.
//
// Solidity: event PodSharesUpdated(address indexed podOwner, int256 sharesDelta)
func (_IEigenPodManager *IEigenPodManagerFilterer) FilterPodSharesUpdated(opts *bind.FilterOpts, podOwner []common.Address) (*IEigenPodManagerPodSharesUpdatedIterator, error) {

	var podOwnerRule []interface{}
	for _, podOwnerItem := range podOwner {
		podOwnerRule = append(podOwnerRule, podOwnerItem)
	}

	logs, sub, err := _IEigenPodManager.contract.FilterLogs(opts, "PodSharesUpdated", podOwnerRule)
	if err != nil {
		return nil, err
	}
	return &IEigenPodManagerPodSharesUpdatedIterator{contract: _IEigenPodManager.contract, event: "PodSharesUpdated", logs: logs, sub: sub}, nil
}

// WatchPodSharesUpdated is a free log subscription operation binding the contract event 0x4e2b791dedccd9fb30141b088cabf5c14a8912b52f59375c95c010700b8c6193.
//
// Solidity: event PodSharesUpdated(address indexed podOwner, int256 sharesDelta)
func (_IEigenPodManager *IEigenPodManagerFilterer) WatchPodSharesUpdated(opts *bind.WatchOpts, sink chan<- *IEigenPodManagerPodSharesUpdated, podOwner []common.Address) (event.Subscription, error) {

	var podOwnerRule []interface{}
	for _, podOwnerItem := range podOwner {
		podOwnerRule = append(podOwnerRule, podOwnerItem)
	}

	logs, sub, err := _IEigenPodManager.contract.WatchLogs(opts, "PodSharesUpdated", podOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEigenPodManagerPodSharesUpdated)
				if err := _IEigenPodManager.contract.UnpackLog(event, "PodSharesUpdated", log); err != nil {
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
func (_IEigenPodManager *IEigenPodManagerFilterer) ParsePodSharesUpdated(log types.Log) (*IEigenPodManagerPodSharesUpdated, error) {
	event := new(IEigenPodManagerPodSharesUpdated)
	if err := _IEigenPodManager.contract.UnpackLog(event, "PodSharesUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IEigenPodManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the IEigenPodManager contract.
type IEigenPodManagerUnpausedIterator struct {
	Event *IEigenPodManagerUnpaused // Event containing the contract specifics and raw log

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
func (it *IEigenPodManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEigenPodManagerUnpaused)
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
		it.Event = new(IEigenPodManagerUnpaused)
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
func (it *IEigenPodManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEigenPodManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEigenPodManagerUnpaused represents a Unpaused event raised by the IEigenPodManager contract.
type IEigenPodManagerUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_IEigenPodManager *IEigenPodManagerFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*IEigenPodManagerUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IEigenPodManager.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &IEigenPodManagerUnpausedIterator{contract: _IEigenPodManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_IEigenPodManager *IEigenPodManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *IEigenPodManagerUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IEigenPodManager.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEigenPodManagerUnpaused)
				if err := _IEigenPodManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_IEigenPodManager *IEigenPodManagerFilterer) ParseUnpaused(log types.Log) (*IEigenPodManagerUnpaused, error) {
	event := new(IEigenPodManagerUnpaused)
	if err := _IEigenPodManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
