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
	ABI: "[{\"type\":\"function\",\"name\":\"addShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"beaconChainETHStrategy\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"beaconChainSlashingFactor\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burnableETHShares\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createPod\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eigenPodBeacon\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBeacon\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ethPOS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIETHPOSDeposit\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPod\",\"inputs\":[{\"name\":\"podOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEigenPod\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasPod\",\"inputs\":[{\"name\":\"podOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"increaseBurnableShares\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"addedSharesToBurn\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"numPods\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownerToPod\",\"inputs\":[{\"name\":\"podOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEigenPod\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pectraForkTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"podOwnerDepositShares\",\"inputs\":[{\"name\":\"podOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"recordBeaconChainETHBalanceUpdate\",\"inputs\":[{\"name\":\"podOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"prevRestakedBalanceWei\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balanceDeltaWei\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeDepositShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"depositSharesToRemove\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPectraForkTimestamp\",\"inputs\":[{\"name\":\"timestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setProofTimestampSetter\",\"inputs\":[{\"name\":\"newProofTimestampSetter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stake\",\"inputs\":[{\"name\":\"pubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"depositDataRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"stakerDepositShares\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"depositShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawSharesAsTokens\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"BeaconChainETHDeposited\",\"inputs\":[{\"name\":\"podOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BeaconChainETHWithdrawalCompleted\",\"inputs\":[{\"name\":\"podOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint96\",\"indexed\":false,\"internalType\":\"uint96\"},{\"name\":\"delegatedAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"withdrawalRoot\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BeaconChainSlashingFactorDecreased\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"prevBeaconChainSlashingFactor\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"newBeaconChainSlashingFactor\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BurnableETHSharesIncreased\",\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewTotalShares\",\"inputs\":[{\"name\":\"podOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newTotalShares\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PectraForkTimestampSet\",\"inputs\":[{\"name\":\"newPectraForkTimestamp\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PodDeployed\",\"inputs\":[{\"name\":\"eigenPod\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"podOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PodSharesUpdated\",\"inputs\":[{\"name\":\"podOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sharesDelta\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProofTimestampSetterSet\",\"inputs\":[{\"name\":\"newProofTimestampSetter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CurrentlyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EigenPodAlreadyExists\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNewPausedStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidStrategy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LegacyWithdrawalsNotCompleted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyDelegationManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyEigenPod\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyPauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyProofTimestampSetter\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnpauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SharesNegative\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SharesNotMultipleOfGwei\",\"inputs\":[]}]",
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

// BeaconChainSlashingFactor is a free data retrieval call binding the contract method 0xa3d75e09.
//
// Solidity: function beaconChainSlashingFactor(address staker) view returns(uint64)
func (_IEigenPodManager *IEigenPodManagerCaller) BeaconChainSlashingFactor(opts *bind.CallOpts, staker common.Address) (uint64, error) {
	var out []interface{}
	err := _IEigenPodManager.contract.Call(opts, &out, "beaconChainSlashingFactor", staker)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// BeaconChainSlashingFactor is a free data retrieval call binding the contract method 0xa3d75e09.
//
// Solidity: function beaconChainSlashingFactor(address staker) view returns(uint64)
func (_IEigenPodManager *IEigenPodManagerSession) BeaconChainSlashingFactor(staker common.Address) (uint64, error) {
	return _IEigenPodManager.Contract.BeaconChainSlashingFactor(&_IEigenPodManager.CallOpts, staker)
}

// BeaconChainSlashingFactor is a free data retrieval call binding the contract method 0xa3d75e09.
//
// Solidity: function beaconChainSlashingFactor(address staker) view returns(uint64)
func (_IEigenPodManager *IEigenPodManagerCallerSession) BeaconChainSlashingFactor(staker common.Address) (uint64, error) {
	return _IEigenPodManager.Contract.BeaconChainSlashingFactor(&_IEigenPodManager.CallOpts, staker)
}

// BurnableETHShares is a free data retrieval call binding the contract method 0xf5d4fed3.
//
// Solidity: function burnableETHShares() view returns(uint256)
func (_IEigenPodManager *IEigenPodManagerCaller) BurnableETHShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IEigenPodManager.contract.Call(opts, &out, "burnableETHShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BurnableETHShares is a free data retrieval call binding the contract method 0xf5d4fed3.
//
// Solidity: function burnableETHShares() view returns(uint256)
func (_IEigenPodManager *IEigenPodManagerSession) BurnableETHShares() (*big.Int, error) {
	return _IEigenPodManager.Contract.BurnableETHShares(&_IEigenPodManager.CallOpts)
}

// BurnableETHShares is a free data retrieval call binding the contract method 0xf5d4fed3.
//
// Solidity: function burnableETHShares() view returns(uint256)
func (_IEigenPodManager *IEigenPodManagerCallerSession) BurnableETHShares() (*big.Int, error) {
	return _IEigenPodManager.Contract.BurnableETHShares(&_IEigenPodManager.CallOpts)
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

// PectraForkTimestamp is a free data retrieval call binding the contract method 0x2704351a.
//
// Solidity: function pectraForkTimestamp() view returns(uint64)
func (_IEigenPodManager *IEigenPodManagerCaller) PectraForkTimestamp(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _IEigenPodManager.contract.Call(opts, &out, "pectraForkTimestamp")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// PectraForkTimestamp is a free data retrieval call binding the contract method 0x2704351a.
//
// Solidity: function pectraForkTimestamp() view returns(uint64)
func (_IEigenPodManager *IEigenPodManagerSession) PectraForkTimestamp() (uint64, error) {
	return _IEigenPodManager.Contract.PectraForkTimestamp(&_IEigenPodManager.CallOpts)
}

// PectraForkTimestamp is a free data retrieval call binding the contract method 0x2704351a.
//
// Solidity: function pectraForkTimestamp() view returns(uint64)
func (_IEigenPodManager *IEigenPodManagerCallerSession) PectraForkTimestamp() (uint64, error) {
	return _IEigenPodManager.Contract.PectraForkTimestamp(&_IEigenPodManager.CallOpts)
}

// PodOwnerDepositShares is a free data retrieval call binding the contract method 0xd48e8894.
//
// Solidity: function podOwnerDepositShares(address podOwner) view returns(int256)
func (_IEigenPodManager *IEigenPodManagerCaller) PodOwnerDepositShares(opts *bind.CallOpts, podOwner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IEigenPodManager.contract.Call(opts, &out, "podOwnerDepositShares", podOwner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PodOwnerDepositShares is a free data retrieval call binding the contract method 0xd48e8894.
//
// Solidity: function podOwnerDepositShares(address podOwner) view returns(int256)
func (_IEigenPodManager *IEigenPodManagerSession) PodOwnerDepositShares(podOwner common.Address) (*big.Int, error) {
	return _IEigenPodManager.Contract.PodOwnerDepositShares(&_IEigenPodManager.CallOpts, podOwner)
}

// PodOwnerDepositShares is a free data retrieval call binding the contract method 0xd48e8894.
//
// Solidity: function podOwnerDepositShares(address podOwner) view returns(int256)
func (_IEigenPodManager *IEigenPodManagerCallerSession) PodOwnerDepositShares(podOwner common.Address) (*big.Int, error) {
	return _IEigenPodManager.Contract.PodOwnerDepositShares(&_IEigenPodManager.CallOpts, podOwner)
}

// StakerDepositShares is a free data retrieval call binding the contract method 0xfe243a17.
//
// Solidity: function stakerDepositShares(address user, address strategy) view returns(uint256 depositShares)
func (_IEigenPodManager *IEigenPodManagerCaller) StakerDepositShares(opts *bind.CallOpts, user common.Address, strategy common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IEigenPodManager.contract.Call(opts, &out, "stakerDepositShares", user, strategy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerDepositShares is a free data retrieval call binding the contract method 0xfe243a17.
//
// Solidity: function stakerDepositShares(address user, address strategy) view returns(uint256 depositShares)
func (_IEigenPodManager *IEigenPodManagerSession) StakerDepositShares(user common.Address, strategy common.Address) (*big.Int, error) {
	return _IEigenPodManager.Contract.StakerDepositShares(&_IEigenPodManager.CallOpts, user, strategy)
}

// StakerDepositShares is a free data retrieval call binding the contract method 0xfe243a17.
//
// Solidity: function stakerDepositShares(address user, address strategy) view returns(uint256 depositShares)
func (_IEigenPodManager *IEigenPodManagerCallerSession) StakerDepositShares(user common.Address, strategy common.Address) (*big.Int, error) {
	return _IEigenPodManager.Contract.StakerDepositShares(&_IEigenPodManager.CallOpts, user, strategy)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_IEigenPodManager *IEigenPodManagerCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IEigenPodManager.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_IEigenPodManager *IEigenPodManagerSession) Version() (string, error) {
	return _IEigenPodManager.Contract.Version(&_IEigenPodManager.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_IEigenPodManager *IEigenPodManagerCallerSession) Version() (string, error) {
	return _IEigenPodManager.Contract.Version(&_IEigenPodManager.CallOpts)
}

// AddShares is a paid mutator transaction binding the contract method 0x50ff7225.
//
// Solidity: function addShares(address staker, address strategy, uint256 shares) returns(uint256, uint256)
func (_IEigenPodManager *IEigenPodManagerTransactor) AddShares(opts *bind.TransactOpts, staker common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error) {
	return _IEigenPodManager.contract.Transact(opts, "addShares", staker, strategy, shares)
}

// AddShares is a paid mutator transaction binding the contract method 0x50ff7225.
//
// Solidity: function addShares(address staker, address strategy, uint256 shares) returns(uint256, uint256)
func (_IEigenPodManager *IEigenPodManagerSession) AddShares(staker common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error) {
	return _IEigenPodManager.Contract.AddShares(&_IEigenPodManager.TransactOpts, staker, strategy, shares)
}

// AddShares is a paid mutator transaction binding the contract method 0x50ff7225.
//
// Solidity: function addShares(address staker, address strategy, uint256 shares) returns(uint256, uint256)
func (_IEigenPodManager *IEigenPodManagerTransactorSession) AddShares(staker common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error) {
	return _IEigenPodManager.Contract.AddShares(&_IEigenPodManager.TransactOpts, staker, strategy, shares)
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

// IncreaseBurnableShares is a paid mutator transaction binding the contract method 0xdebe1eab.
//
// Solidity: function increaseBurnableShares(address strategy, uint256 addedSharesToBurn) returns()
func (_IEigenPodManager *IEigenPodManagerTransactor) IncreaseBurnableShares(opts *bind.TransactOpts, strategy common.Address, addedSharesToBurn *big.Int) (*types.Transaction, error) {
	return _IEigenPodManager.contract.Transact(opts, "increaseBurnableShares", strategy, addedSharesToBurn)
}

// IncreaseBurnableShares is a paid mutator transaction binding the contract method 0xdebe1eab.
//
// Solidity: function increaseBurnableShares(address strategy, uint256 addedSharesToBurn) returns()
func (_IEigenPodManager *IEigenPodManagerSession) IncreaseBurnableShares(strategy common.Address, addedSharesToBurn *big.Int) (*types.Transaction, error) {
	return _IEigenPodManager.Contract.IncreaseBurnableShares(&_IEigenPodManager.TransactOpts, strategy, addedSharesToBurn)
}

// IncreaseBurnableShares is a paid mutator transaction binding the contract method 0xdebe1eab.
//
// Solidity: function increaseBurnableShares(address strategy, uint256 addedSharesToBurn) returns()
func (_IEigenPodManager *IEigenPodManagerTransactorSession) IncreaseBurnableShares(strategy common.Address, addedSharesToBurn *big.Int) (*types.Transaction, error) {
	return _IEigenPodManager.Contract.IncreaseBurnableShares(&_IEigenPodManager.TransactOpts, strategy, addedSharesToBurn)
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

// RecordBeaconChainETHBalanceUpdate is a paid mutator transaction binding the contract method 0xa1ca780b.
//
// Solidity: function recordBeaconChainETHBalanceUpdate(address podOwner, uint256 prevRestakedBalanceWei, int256 balanceDeltaWei) returns()
func (_IEigenPodManager *IEigenPodManagerTransactor) RecordBeaconChainETHBalanceUpdate(opts *bind.TransactOpts, podOwner common.Address, prevRestakedBalanceWei *big.Int, balanceDeltaWei *big.Int) (*types.Transaction, error) {
	return _IEigenPodManager.contract.Transact(opts, "recordBeaconChainETHBalanceUpdate", podOwner, prevRestakedBalanceWei, balanceDeltaWei)
}

// RecordBeaconChainETHBalanceUpdate is a paid mutator transaction binding the contract method 0xa1ca780b.
//
// Solidity: function recordBeaconChainETHBalanceUpdate(address podOwner, uint256 prevRestakedBalanceWei, int256 balanceDeltaWei) returns()
func (_IEigenPodManager *IEigenPodManagerSession) RecordBeaconChainETHBalanceUpdate(podOwner common.Address, prevRestakedBalanceWei *big.Int, balanceDeltaWei *big.Int) (*types.Transaction, error) {
	return _IEigenPodManager.Contract.RecordBeaconChainETHBalanceUpdate(&_IEigenPodManager.TransactOpts, podOwner, prevRestakedBalanceWei, balanceDeltaWei)
}

// RecordBeaconChainETHBalanceUpdate is a paid mutator transaction binding the contract method 0xa1ca780b.
//
// Solidity: function recordBeaconChainETHBalanceUpdate(address podOwner, uint256 prevRestakedBalanceWei, int256 balanceDeltaWei) returns()
func (_IEigenPodManager *IEigenPodManagerTransactorSession) RecordBeaconChainETHBalanceUpdate(podOwner common.Address, prevRestakedBalanceWei *big.Int, balanceDeltaWei *big.Int) (*types.Transaction, error) {
	return _IEigenPodManager.Contract.RecordBeaconChainETHBalanceUpdate(&_IEigenPodManager.TransactOpts, podOwner, prevRestakedBalanceWei, balanceDeltaWei)
}

// RemoveDepositShares is a paid mutator transaction binding the contract method 0x724af423.
//
// Solidity: function removeDepositShares(address staker, address strategy, uint256 depositSharesToRemove) returns(uint256)
func (_IEigenPodManager *IEigenPodManagerTransactor) RemoveDepositShares(opts *bind.TransactOpts, staker common.Address, strategy common.Address, depositSharesToRemove *big.Int) (*types.Transaction, error) {
	return _IEigenPodManager.contract.Transact(opts, "removeDepositShares", staker, strategy, depositSharesToRemove)
}

// RemoveDepositShares is a paid mutator transaction binding the contract method 0x724af423.
//
// Solidity: function removeDepositShares(address staker, address strategy, uint256 depositSharesToRemove) returns(uint256)
func (_IEigenPodManager *IEigenPodManagerSession) RemoveDepositShares(staker common.Address, strategy common.Address, depositSharesToRemove *big.Int) (*types.Transaction, error) {
	return _IEigenPodManager.Contract.RemoveDepositShares(&_IEigenPodManager.TransactOpts, staker, strategy, depositSharesToRemove)
}

// RemoveDepositShares is a paid mutator transaction binding the contract method 0x724af423.
//
// Solidity: function removeDepositShares(address staker, address strategy, uint256 depositSharesToRemove) returns(uint256)
func (_IEigenPodManager *IEigenPodManagerTransactorSession) RemoveDepositShares(staker common.Address, strategy common.Address, depositSharesToRemove *big.Int) (*types.Transaction, error) {
	return _IEigenPodManager.Contract.RemoveDepositShares(&_IEigenPodManager.TransactOpts, staker, strategy, depositSharesToRemove)
}

// SetPectraForkTimestamp is a paid mutator transaction binding the contract method 0x5a26fbf4.
//
// Solidity: function setPectraForkTimestamp(uint64 timestamp) returns()
func (_IEigenPodManager *IEigenPodManagerTransactor) SetPectraForkTimestamp(opts *bind.TransactOpts, timestamp uint64) (*types.Transaction, error) {
	return _IEigenPodManager.contract.Transact(opts, "setPectraForkTimestamp", timestamp)
}

// SetPectraForkTimestamp is a paid mutator transaction binding the contract method 0x5a26fbf4.
//
// Solidity: function setPectraForkTimestamp(uint64 timestamp) returns()
func (_IEigenPodManager *IEigenPodManagerSession) SetPectraForkTimestamp(timestamp uint64) (*types.Transaction, error) {
	return _IEigenPodManager.Contract.SetPectraForkTimestamp(&_IEigenPodManager.TransactOpts, timestamp)
}

// SetPectraForkTimestamp is a paid mutator transaction binding the contract method 0x5a26fbf4.
//
// Solidity: function setPectraForkTimestamp(uint64 timestamp) returns()
func (_IEigenPodManager *IEigenPodManagerTransactorSession) SetPectraForkTimestamp(timestamp uint64) (*types.Transaction, error) {
	return _IEigenPodManager.Contract.SetPectraForkTimestamp(&_IEigenPodManager.TransactOpts, timestamp)
}

// SetProofTimestampSetter is a paid mutator transaction binding the contract method 0x0d1e9de1.
//
// Solidity: function setProofTimestampSetter(address newProofTimestampSetter) returns()
func (_IEigenPodManager *IEigenPodManagerTransactor) SetProofTimestampSetter(opts *bind.TransactOpts, newProofTimestampSetter common.Address) (*types.Transaction, error) {
	return _IEigenPodManager.contract.Transact(opts, "setProofTimestampSetter", newProofTimestampSetter)
}

// SetProofTimestampSetter is a paid mutator transaction binding the contract method 0x0d1e9de1.
//
// Solidity: function setProofTimestampSetter(address newProofTimestampSetter) returns()
func (_IEigenPodManager *IEigenPodManagerSession) SetProofTimestampSetter(newProofTimestampSetter common.Address) (*types.Transaction, error) {
	return _IEigenPodManager.Contract.SetProofTimestampSetter(&_IEigenPodManager.TransactOpts, newProofTimestampSetter)
}

// SetProofTimestampSetter is a paid mutator transaction binding the contract method 0x0d1e9de1.
//
// Solidity: function setProofTimestampSetter(address newProofTimestampSetter) returns()
func (_IEigenPodManager *IEigenPodManagerTransactorSession) SetProofTimestampSetter(newProofTimestampSetter common.Address) (*types.Transaction, error) {
	return _IEigenPodManager.Contract.SetProofTimestampSetter(&_IEigenPodManager.TransactOpts, newProofTimestampSetter)
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

// WithdrawSharesAsTokens is a paid mutator transaction binding the contract method 0x2eae418c.
//
// Solidity: function withdrawSharesAsTokens(address staker, address strategy, address token, uint256 shares) returns()
func (_IEigenPodManager *IEigenPodManagerTransactor) WithdrawSharesAsTokens(opts *bind.TransactOpts, staker common.Address, strategy common.Address, token common.Address, shares *big.Int) (*types.Transaction, error) {
	return _IEigenPodManager.contract.Transact(opts, "withdrawSharesAsTokens", staker, strategy, token, shares)
}

// WithdrawSharesAsTokens is a paid mutator transaction binding the contract method 0x2eae418c.
//
// Solidity: function withdrawSharesAsTokens(address staker, address strategy, address token, uint256 shares) returns()
func (_IEigenPodManager *IEigenPodManagerSession) WithdrawSharesAsTokens(staker common.Address, strategy common.Address, token common.Address, shares *big.Int) (*types.Transaction, error) {
	return _IEigenPodManager.Contract.WithdrawSharesAsTokens(&_IEigenPodManager.TransactOpts, staker, strategy, token, shares)
}

// WithdrawSharesAsTokens is a paid mutator transaction binding the contract method 0x2eae418c.
//
// Solidity: function withdrawSharesAsTokens(address staker, address strategy, address token, uint256 shares) returns()
func (_IEigenPodManager *IEigenPodManagerTransactorSession) WithdrawSharesAsTokens(staker common.Address, strategy common.Address, token common.Address, shares *big.Int) (*types.Transaction, error) {
	return _IEigenPodManager.Contract.WithdrawSharesAsTokens(&_IEigenPodManager.TransactOpts, staker, strategy, token, shares)
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

// IEigenPodManagerBeaconChainSlashingFactorDecreasedIterator is returned from FilterBeaconChainSlashingFactorDecreased and is used to iterate over the raw logs and unpacked data for BeaconChainSlashingFactorDecreased events raised by the IEigenPodManager contract.
type IEigenPodManagerBeaconChainSlashingFactorDecreasedIterator struct {
	Event *IEigenPodManagerBeaconChainSlashingFactorDecreased // Event containing the contract specifics and raw log

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
func (it *IEigenPodManagerBeaconChainSlashingFactorDecreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEigenPodManagerBeaconChainSlashingFactorDecreased)
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
		it.Event = new(IEigenPodManagerBeaconChainSlashingFactorDecreased)
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
func (it *IEigenPodManagerBeaconChainSlashingFactorDecreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEigenPodManagerBeaconChainSlashingFactorDecreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEigenPodManagerBeaconChainSlashingFactorDecreased represents a BeaconChainSlashingFactorDecreased event raised by the IEigenPodManager contract.
type IEigenPodManagerBeaconChainSlashingFactorDecreased struct {
	Staker                        common.Address
	PrevBeaconChainSlashingFactor uint64
	NewBeaconChainSlashingFactor  uint64
	Raw                           types.Log // Blockchain specific contextual infos
}

// FilterBeaconChainSlashingFactorDecreased is a free log retrieval operation binding the contract event 0xb160ab8589bf47dc04ea11b50d46678d21590cea2ed3e454e7bd3e41510f98cf.
//
// Solidity: event BeaconChainSlashingFactorDecreased(address staker, uint64 prevBeaconChainSlashingFactor, uint64 newBeaconChainSlashingFactor)
func (_IEigenPodManager *IEigenPodManagerFilterer) FilterBeaconChainSlashingFactorDecreased(opts *bind.FilterOpts) (*IEigenPodManagerBeaconChainSlashingFactorDecreasedIterator, error) {

	logs, sub, err := _IEigenPodManager.contract.FilterLogs(opts, "BeaconChainSlashingFactorDecreased")
	if err != nil {
		return nil, err
	}
	return &IEigenPodManagerBeaconChainSlashingFactorDecreasedIterator{contract: _IEigenPodManager.contract, event: "BeaconChainSlashingFactorDecreased", logs: logs, sub: sub}, nil
}

// WatchBeaconChainSlashingFactorDecreased is a free log subscription operation binding the contract event 0xb160ab8589bf47dc04ea11b50d46678d21590cea2ed3e454e7bd3e41510f98cf.
//
// Solidity: event BeaconChainSlashingFactorDecreased(address staker, uint64 prevBeaconChainSlashingFactor, uint64 newBeaconChainSlashingFactor)
func (_IEigenPodManager *IEigenPodManagerFilterer) WatchBeaconChainSlashingFactorDecreased(opts *bind.WatchOpts, sink chan<- *IEigenPodManagerBeaconChainSlashingFactorDecreased) (event.Subscription, error) {

	logs, sub, err := _IEigenPodManager.contract.WatchLogs(opts, "BeaconChainSlashingFactorDecreased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEigenPodManagerBeaconChainSlashingFactorDecreased)
				if err := _IEigenPodManager.contract.UnpackLog(event, "BeaconChainSlashingFactorDecreased", log); err != nil {
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

// ParseBeaconChainSlashingFactorDecreased is a log parse operation binding the contract event 0xb160ab8589bf47dc04ea11b50d46678d21590cea2ed3e454e7bd3e41510f98cf.
//
// Solidity: event BeaconChainSlashingFactorDecreased(address staker, uint64 prevBeaconChainSlashingFactor, uint64 newBeaconChainSlashingFactor)
func (_IEigenPodManager *IEigenPodManagerFilterer) ParseBeaconChainSlashingFactorDecreased(log types.Log) (*IEigenPodManagerBeaconChainSlashingFactorDecreased, error) {
	event := new(IEigenPodManagerBeaconChainSlashingFactorDecreased)
	if err := _IEigenPodManager.contract.UnpackLog(event, "BeaconChainSlashingFactorDecreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IEigenPodManagerBurnableETHSharesIncreasedIterator is returned from FilterBurnableETHSharesIncreased and is used to iterate over the raw logs and unpacked data for BurnableETHSharesIncreased events raised by the IEigenPodManager contract.
type IEigenPodManagerBurnableETHSharesIncreasedIterator struct {
	Event *IEigenPodManagerBurnableETHSharesIncreased // Event containing the contract specifics and raw log

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
func (it *IEigenPodManagerBurnableETHSharesIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEigenPodManagerBurnableETHSharesIncreased)
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
		it.Event = new(IEigenPodManagerBurnableETHSharesIncreased)
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
func (it *IEigenPodManagerBurnableETHSharesIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEigenPodManagerBurnableETHSharesIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEigenPodManagerBurnableETHSharesIncreased represents a BurnableETHSharesIncreased event raised by the IEigenPodManager contract.
type IEigenPodManagerBurnableETHSharesIncreased struct {
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurnableETHSharesIncreased is a free log retrieval operation binding the contract event 0x1ed04b7fd262c0d9e50fa02957f32a81a151f03baaa367faeedc7521b001c4a4.
//
// Solidity: event BurnableETHSharesIncreased(uint256 shares)
func (_IEigenPodManager *IEigenPodManagerFilterer) FilterBurnableETHSharesIncreased(opts *bind.FilterOpts) (*IEigenPodManagerBurnableETHSharesIncreasedIterator, error) {

	logs, sub, err := _IEigenPodManager.contract.FilterLogs(opts, "BurnableETHSharesIncreased")
	if err != nil {
		return nil, err
	}
	return &IEigenPodManagerBurnableETHSharesIncreasedIterator{contract: _IEigenPodManager.contract, event: "BurnableETHSharesIncreased", logs: logs, sub: sub}, nil
}

// WatchBurnableETHSharesIncreased is a free log subscription operation binding the contract event 0x1ed04b7fd262c0d9e50fa02957f32a81a151f03baaa367faeedc7521b001c4a4.
//
// Solidity: event BurnableETHSharesIncreased(uint256 shares)
func (_IEigenPodManager *IEigenPodManagerFilterer) WatchBurnableETHSharesIncreased(opts *bind.WatchOpts, sink chan<- *IEigenPodManagerBurnableETHSharesIncreased) (event.Subscription, error) {

	logs, sub, err := _IEigenPodManager.contract.WatchLogs(opts, "BurnableETHSharesIncreased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEigenPodManagerBurnableETHSharesIncreased)
				if err := _IEigenPodManager.contract.UnpackLog(event, "BurnableETHSharesIncreased", log); err != nil {
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

// ParseBurnableETHSharesIncreased is a log parse operation binding the contract event 0x1ed04b7fd262c0d9e50fa02957f32a81a151f03baaa367faeedc7521b001c4a4.
//
// Solidity: event BurnableETHSharesIncreased(uint256 shares)
func (_IEigenPodManager *IEigenPodManagerFilterer) ParseBurnableETHSharesIncreased(log types.Log) (*IEigenPodManagerBurnableETHSharesIncreased, error) {
	event := new(IEigenPodManagerBurnableETHSharesIncreased)
	if err := _IEigenPodManager.contract.UnpackLog(event, "BurnableETHSharesIncreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IEigenPodManagerNewTotalSharesIterator is returned from FilterNewTotalShares and is used to iterate over the raw logs and unpacked data for NewTotalShares events raised by the IEigenPodManager contract.
type IEigenPodManagerNewTotalSharesIterator struct {
	Event *IEigenPodManagerNewTotalShares // Event containing the contract specifics and raw log

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
func (it *IEigenPodManagerNewTotalSharesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEigenPodManagerNewTotalShares)
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
		it.Event = new(IEigenPodManagerNewTotalShares)
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
func (it *IEigenPodManagerNewTotalSharesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEigenPodManagerNewTotalSharesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEigenPodManagerNewTotalShares represents a NewTotalShares event raised by the IEigenPodManager contract.
type IEigenPodManagerNewTotalShares struct {
	PodOwner       common.Address
	NewTotalShares *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewTotalShares is a free log retrieval operation binding the contract event 0xd4def76d6d2bed6f14d5cd9af73cc2913d618d00edde42432e81c09bfe077098.
//
// Solidity: event NewTotalShares(address indexed podOwner, int256 newTotalShares)
func (_IEigenPodManager *IEigenPodManagerFilterer) FilterNewTotalShares(opts *bind.FilterOpts, podOwner []common.Address) (*IEigenPodManagerNewTotalSharesIterator, error) {

	var podOwnerRule []interface{}
	for _, podOwnerItem := range podOwner {
		podOwnerRule = append(podOwnerRule, podOwnerItem)
	}

	logs, sub, err := _IEigenPodManager.contract.FilterLogs(opts, "NewTotalShares", podOwnerRule)
	if err != nil {
		return nil, err
	}
	return &IEigenPodManagerNewTotalSharesIterator{contract: _IEigenPodManager.contract, event: "NewTotalShares", logs: logs, sub: sub}, nil
}

// WatchNewTotalShares is a free log subscription operation binding the contract event 0xd4def76d6d2bed6f14d5cd9af73cc2913d618d00edde42432e81c09bfe077098.
//
// Solidity: event NewTotalShares(address indexed podOwner, int256 newTotalShares)
func (_IEigenPodManager *IEigenPodManagerFilterer) WatchNewTotalShares(opts *bind.WatchOpts, sink chan<- *IEigenPodManagerNewTotalShares, podOwner []common.Address) (event.Subscription, error) {

	var podOwnerRule []interface{}
	for _, podOwnerItem := range podOwner {
		podOwnerRule = append(podOwnerRule, podOwnerItem)
	}

	logs, sub, err := _IEigenPodManager.contract.WatchLogs(opts, "NewTotalShares", podOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEigenPodManagerNewTotalShares)
				if err := _IEigenPodManager.contract.UnpackLog(event, "NewTotalShares", log); err != nil {
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
func (_IEigenPodManager *IEigenPodManagerFilterer) ParseNewTotalShares(log types.Log) (*IEigenPodManagerNewTotalShares, error) {
	event := new(IEigenPodManagerNewTotalShares)
	if err := _IEigenPodManager.contract.UnpackLog(event, "NewTotalShares", log); err != nil {
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

// IEigenPodManagerPectraForkTimestampSetIterator is returned from FilterPectraForkTimestampSet and is used to iterate over the raw logs and unpacked data for PectraForkTimestampSet events raised by the IEigenPodManager contract.
type IEigenPodManagerPectraForkTimestampSetIterator struct {
	Event *IEigenPodManagerPectraForkTimestampSet // Event containing the contract specifics and raw log

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
func (it *IEigenPodManagerPectraForkTimestampSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEigenPodManagerPectraForkTimestampSet)
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
		it.Event = new(IEigenPodManagerPectraForkTimestampSet)
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
func (it *IEigenPodManagerPectraForkTimestampSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEigenPodManagerPectraForkTimestampSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEigenPodManagerPectraForkTimestampSet represents a PectraForkTimestampSet event raised by the IEigenPodManager contract.
type IEigenPodManagerPectraForkTimestampSet struct {
	NewPectraForkTimestamp uint64
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterPectraForkTimestampSet is a free log retrieval operation binding the contract event 0x1bc8f042a52db3a437620dea4548f2031fb2a16dd8d3b0b854295528dd2cdd33.
//
// Solidity: event PectraForkTimestampSet(uint64 newPectraForkTimestamp)
func (_IEigenPodManager *IEigenPodManagerFilterer) FilterPectraForkTimestampSet(opts *bind.FilterOpts) (*IEigenPodManagerPectraForkTimestampSetIterator, error) {

	logs, sub, err := _IEigenPodManager.contract.FilterLogs(opts, "PectraForkTimestampSet")
	if err != nil {
		return nil, err
	}
	return &IEigenPodManagerPectraForkTimestampSetIterator{contract: _IEigenPodManager.contract, event: "PectraForkTimestampSet", logs: logs, sub: sub}, nil
}

// WatchPectraForkTimestampSet is a free log subscription operation binding the contract event 0x1bc8f042a52db3a437620dea4548f2031fb2a16dd8d3b0b854295528dd2cdd33.
//
// Solidity: event PectraForkTimestampSet(uint64 newPectraForkTimestamp)
func (_IEigenPodManager *IEigenPodManagerFilterer) WatchPectraForkTimestampSet(opts *bind.WatchOpts, sink chan<- *IEigenPodManagerPectraForkTimestampSet) (event.Subscription, error) {

	logs, sub, err := _IEigenPodManager.contract.WatchLogs(opts, "PectraForkTimestampSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEigenPodManagerPectraForkTimestampSet)
				if err := _IEigenPodManager.contract.UnpackLog(event, "PectraForkTimestampSet", log); err != nil {
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

// ParsePectraForkTimestampSet is a log parse operation binding the contract event 0x1bc8f042a52db3a437620dea4548f2031fb2a16dd8d3b0b854295528dd2cdd33.
//
// Solidity: event PectraForkTimestampSet(uint64 newPectraForkTimestamp)
func (_IEigenPodManager *IEigenPodManagerFilterer) ParsePectraForkTimestampSet(log types.Log) (*IEigenPodManagerPectraForkTimestampSet, error) {
	event := new(IEigenPodManagerPectraForkTimestampSet)
	if err := _IEigenPodManager.contract.UnpackLog(event, "PectraForkTimestampSet", log); err != nil {
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

// IEigenPodManagerProofTimestampSetterSetIterator is returned from FilterProofTimestampSetterSet and is used to iterate over the raw logs and unpacked data for ProofTimestampSetterSet events raised by the IEigenPodManager contract.
type IEigenPodManagerProofTimestampSetterSetIterator struct {
	Event *IEigenPodManagerProofTimestampSetterSet // Event containing the contract specifics and raw log

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
func (it *IEigenPodManagerProofTimestampSetterSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEigenPodManagerProofTimestampSetterSet)
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
		it.Event = new(IEigenPodManagerProofTimestampSetterSet)
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
func (it *IEigenPodManagerProofTimestampSetterSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEigenPodManagerProofTimestampSetterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEigenPodManagerProofTimestampSetterSet represents a ProofTimestampSetterSet event raised by the IEigenPodManager contract.
type IEigenPodManagerProofTimestampSetterSet struct {
	NewProofTimestampSetter common.Address
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterProofTimestampSetterSet is a free log retrieval operation binding the contract event 0x7025c71a9fe60d709e71b377dc5f7c72c3e1d8539f8022574254e736ceca01e5.
//
// Solidity: event ProofTimestampSetterSet(address newProofTimestampSetter)
func (_IEigenPodManager *IEigenPodManagerFilterer) FilterProofTimestampSetterSet(opts *bind.FilterOpts) (*IEigenPodManagerProofTimestampSetterSetIterator, error) {

	logs, sub, err := _IEigenPodManager.contract.FilterLogs(opts, "ProofTimestampSetterSet")
	if err != nil {
		return nil, err
	}
	return &IEigenPodManagerProofTimestampSetterSetIterator{contract: _IEigenPodManager.contract, event: "ProofTimestampSetterSet", logs: logs, sub: sub}, nil
}

// WatchProofTimestampSetterSet is a free log subscription operation binding the contract event 0x7025c71a9fe60d709e71b377dc5f7c72c3e1d8539f8022574254e736ceca01e5.
//
// Solidity: event ProofTimestampSetterSet(address newProofTimestampSetter)
func (_IEigenPodManager *IEigenPodManagerFilterer) WatchProofTimestampSetterSet(opts *bind.WatchOpts, sink chan<- *IEigenPodManagerProofTimestampSetterSet) (event.Subscription, error) {

	logs, sub, err := _IEigenPodManager.contract.WatchLogs(opts, "ProofTimestampSetterSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEigenPodManagerProofTimestampSetterSet)
				if err := _IEigenPodManager.contract.UnpackLog(event, "ProofTimestampSetterSet", log); err != nil {
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

// ParseProofTimestampSetterSet is a log parse operation binding the contract event 0x7025c71a9fe60d709e71b377dc5f7c72c3e1d8539f8022574254e736ceca01e5.
//
// Solidity: event ProofTimestampSetterSet(address newProofTimestampSetter)
func (_IEigenPodManager *IEigenPodManagerFilterer) ParseProofTimestampSetterSet(log types.Log) (*IEigenPodManagerProofTimestampSetterSet, error) {
	event := new(IEigenPodManagerProofTimestampSetterSet)
	if err := _IEigenPodManager.contract.UnpackLog(event, "ProofTimestampSetterSet", log); err != nil {
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
