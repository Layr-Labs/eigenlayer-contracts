// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package DurationVaultStrategyStorage

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

// DurationVaultStrategyStorageMetaData contains all meta data concerning the DurationVaultStrategyStorage contract.
var DurationVaultStrategyStorageMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"advanceToWithdrawals\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allocationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allocationsActive\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"arbitrator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"beforeAddShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"beforeRemoveShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositsOpen\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"duration\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"explanation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isLocked\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isMatured\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lock\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockedAt\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"markMatured\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"maturedAt\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxPerDeposit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxTotalDeposits\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"metadataURI\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorIntegrationConfigured\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorSetInfo\",\"inputs\":[],\"outputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorSetRegistered\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"rewardsCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRewardsCoordinator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setRewardsClaimer\",\"inputs\":[{\"name\":\"claimer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"shares\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sharesToUnderlying\",\"inputs\":[{\"name\":\"amountShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sharesToUnderlyingView\",\"inputs\":[{\"name\":\"amountShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakeCap\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"state\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIDurationVaultStrategyTypes.VaultState\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalShares\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"underlyingToShares\",\"inputs\":[{\"name\":\"amountUnderlying\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"underlyingToSharesView\",\"inputs\":[{\"name\":\"amountUnderlying\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"underlyingToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unlockAt\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unlockTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateDelegationApprover\",\"inputs\":[{\"name\":\"newDelegationApprover\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateMetadataURI\",\"inputs\":[{\"name\":\"newMetadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOperatorMetadataURI\",\"inputs\":[{\"name\":\"newOperatorMetadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateTVLLimits\",\"inputs\":[{\"name\":\"newMaxPerDeposit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"newStakeCap\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"userUnderlying\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"userUnderlyingView\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"vaultAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amountShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawalsOpen\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"DeallocateAttempted\",\"inputs\":[{\"name\":\"success\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DeregisterAttempted\",\"inputs\":[{\"name\":\"success\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExchangeRateEmitted\",\"inputs\":[{\"name\":\"rate\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MaxPerDepositUpdated\",\"inputs\":[{\"name\":\"previousValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MaxTotalDepositsUpdated\",\"inputs\":[{\"name\":\"previousValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MetadataURIUpdated\",\"inputs\":[{\"name\":\"newMetadataURI\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyTokenSet\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIERC20\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VaultAdvancedToWithdrawals\",\"inputs\":[{\"name\":\"arbitrator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"maturedAt\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VaultInitialized\",\"inputs\":[{\"name\":\"vaultAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"arbitrator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"underlyingToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"contractIERC20\"},{\"name\":\"duration\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"maxPerDeposit\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"stakeCap\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"metadataURI\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VaultLocked\",\"inputs\":[{\"name\":\"lockedAt\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"unlockAt\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VaultMatured\",\"inputs\":[{\"name\":\"maturedAt\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BalanceExceedsMaxTotalDeposits\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DepositExceedsMaxPerDeposit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DepositsLocked\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DurationAlreadyElapsed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DurationNotElapsed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidArbitrator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDuration\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidVaultAdmin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MaxPerDepositExceedsMax\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeDelegatedToVaultOperator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NewSharesZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyArbitrator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyStrategyManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnderlyingToken\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyVaultAdmin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorIntegrationInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PendingAllocation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategyNotSupportedByOperatorSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TotalSharesExceedsMax\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnderlyingTokenBlacklisted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"VaultAlreadyLocked\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"VaultNotLocked\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawalAmountExceedsTotalDeposits\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawalsLockedDuringAllocations\",\"inputs\":[]}]",
}

// DurationVaultStrategyStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use DurationVaultStrategyStorageMetaData.ABI instead.
var DurationVaultStrategyStorageABI = DurationVaultStrategyStorageMetaData.ABI

// DurationVaultStrategyStorage is an auto generated Go binding around an Ethereum contract.
type DurationVaultStrategyStorage struct {
	DurationVaultStrategyStorageCaller     // Read-only binding to the contract
	DurationVaultStrategyStorageTransactor // Write-only binding to the contract
	DurationVaultStrategyStorageFilterer   // Log filterer for contract events
}

// DurationVaultStrategyStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type DurationVaultStrategyStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DurationVaultStrategyStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DurationVaultStrategyStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DurationVaultStrategyStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DurationVaultStrategyStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DurationVaultStrategyStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DurationVaultStrategyStorageSession struct {
	Contract     *DurationVaultStrategyStorage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// DurationVaultStrategyStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DurationVaultStrategyStorageCallerSession struct {
	Contract *DurationVaultStrategyStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// DurationVaultStrategyStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DurationVaultStrategyStorageTransactorSession struct {
	Contract     *DurationVaultStrategyStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// DurationVaultStrategyStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type DurationVaultStrategyStorageRaw struct {
	Contract *DurationVaultStrategyStorage // Generic contract binding to access the raw methods on
}

// DurationVaultStrategyStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DurationVaultStrategyStorageCallerRaw struct {
	Contract *DurationVaultStrategyStorageCaller // Generic read-only contract binding to access the raw methods on
}

// DurationVaultStrategyStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DurationVaultStrategyStorageTransactorRaw struct {
	Contract *DurationVaultStrategyStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDurationVaultStrategyStorage creates a new instance of DurationVaultStrategyStorage, bound to a specific deployed contract.
func NewDurationVaultStrategyStorage(address common.Address, backend bind.ContractBackend) (*DurationVaultStrategyStorage, error) {
	contract, err := bindDurationVaultStrategyStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DurationVaultStrategyStorage{DurationVaultStrategyStorageCaller: DurationVaultStrategyStorageCaller{contract: contract}, DurationVaultStrategyStorageTransactor: DurationVaultStrategyStorageTransactor{contract: contract}, DurationVaultStrategyStorageFilterer: DurationVaultStrategyStorageFilterer{contract: contract}}, nil
}

// NewDurationVaultStrategyStorageCaller creates a new read-only instance of DurationVaultStrategyStorage, bound to a specific deployed contract.
func NewDurationVaultStrategyStorageCaller(address common.Address, caller bind.ContractCaller) (*DurationVaultStrategyStorageCaller, error) {
	contract, err := bindDurationVaultStrategyStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DurationVaultStrategyStorageCaller{contract: contract}, nil
}

// NewDurationVaultStrategyStorageTransactor creates a new write-only instance of DurationVaultStrategyStorage, bound to a specific deployed contract.
func NewDurationVaultStrategyStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*DurationVaultStrategyStorageTransactor, error) {
	contract, err := bindDurationVaultStrategyStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DurationVaultStrategyStorageTransactor{contract: contract}, nil
}

// NewDurationVaultStrategyStorageFilterer creates a new log filterer instance of DurationVaultStrategyStorage, bound to a specific deployed contract.
func NewDurationVaultStrategyStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*DurationVaultStrategyStorageFilterer, error) {
	contract, err := bindDurationVaultStrategyStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DurationVaultStrategyStorageFilterer{contract: contract}, nil
}

// bindDurationVaultStrategyStorage binds a generic wrapper to an already deployed contract.
func bindDurationVaultStrategyStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DurationVaultStrategyStorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DurationVaultStrategyStorage.Contract.DurationVaultStrategyStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DurationVaultStrategyStorage.Contract.DurationVaultStrategyStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DurationVaultStrategyStorage.Contract.DurationVaultStrategyStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DurationVaultStrategyStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DurationVaultStrategyStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DurationVaultStrategyStorage.Contract.contract.Transact(opts, method, params...)
}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageCaller) AllocationManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DurationVaultStrategyStorage.contract.Call(opts, &out, "allocationManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageSession) AllocationManager() (common.Address, error) {
	return _DurationVaultStrategyStorage.Contract.AllocationManager(&_DurationVaultStrategyStorage.CallOpts)
}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageCallerSession) AllocationManager() (common.Address, error) {
	return _DurationVaultStrategyStorage.Contract.AllocationManager(&_DurationVaultStrategyStorage.CallOpts)
}

// AllocationsActive is a free data retrieval call binding the contract method 0xfb4d86b4.
//
// Solidity: function allocationsActive() view returns(bool)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageCaller) AllocationsActive(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DurationVaultStrategyStorage.contract.Call(opts, &out, "allocationsActive")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllocationsActive is a free data retrieval call binding the contract method 0xfb4d86b4.
//
// Solidity: function allocationsActive() view returns(bool)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageSession) AllocationsActive() (bool, error) {
	return _DurationVaultStrategyStorage.Contract.AllocationsActive(&_DurationVaultStrategyStorage.CallOpts)
}

// AllocationsActive is a free data retrieval call binding the contract method 0xfb4d86b4.
//
// Solidity: function allocationsActive() view returns(bool)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageCallerSession) AllocationsActive() (bool, error) {
	return _DurationVaultStrategyStorage.Contract.AllocationsActive(&_DurationVaultStrategyStorage.CallOpts)
}

// Arbitrator is a free data retrieval call binding the contract method 0x6cc6cde1.
//
// Solidity: function arbitrator() view returns(address)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageCaller) Arbitrator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DurationVaultStrategyStorage.contract.Call(opts, &out, "arbitrator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Arbitrator is a free data retrieval call binding the contract method 0x6cc6cde1.
//
// Solidity: function arbitrator() view returns(address)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageSession) Arbitrator() (common.Address, error) {
	return _DurationVaultStrategyStorage.Contract.Arbitrator(&_DurationVaultStrategyStorage.CallOpts)
}

// Arbitrator is a free data retrieval call binding the contract method 0x6cc6cde1.
//
// Solidity: function arbitrator() view returns(address)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageCallerSession) Arbitrator() (common.Address, error) {
	return _DurationVaultStrategyStorage.Contract.Arbitrator(&_DurationVaultStrategyStorage.CallOpts)
}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageCaller) DelegationManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DurationVaultStrategyStorage.contract.Call(opts, &out, "delegationManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageSession) DelegationManager() (common.Address, error) {
	return _DurationVaultStrategyStorage.Contract.DelegationManager(&_DurationVaultStrategyStorage.CallOpts)
}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageCallerSession) DelegationManager() (common.Address, error) {
	return _DurationVaultStrategyStorage.Contract.DelegationManager(&_DurationVaultStrategyStorage.CallOpts)
}

// DepositsOpen is a free data retrieval call binding the contract method 0x549c4627.
//
// Solidity: function depositsOpen() view returns(bool)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageCaller) DepositsOpen(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DurationVaultStrategyStorage.contract.Call(opts, &out, "depositsOpen")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DepositsOpen is a free data retrieval call binding the contract method 0x549c4627.
//
// Solidity: function depositsOpen() view returns(bool)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageSession) DepositsOpen() (bool, error) {
	return _DurationVaultStrategyStorage.Contract.DepositsOpen(&_DurationVaultStrategyStorage.CallOpts)
}

// DepositsOpen is a free data retrieval call binding the contract method 0x549c4627.
//
// Solidity: function depositsOpen() view returns(bool)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageCallerSession) DepositsOpen() (bool, error) {
	return _DurationVaultStrategyStorage.Contract.DepositsOpen(&_DurationVaultStrategyStorage.CallOpts)
}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint32)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageCaller) Duration(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _DurationVaultStrategyStorage.contract.Call(opts, &out, "duration")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint32)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageSession) Duration() (uint32, error) {
	return _DurationVaultStrategyStorage.Contract.Duration(&_DurationVaultStrategyStorage.CallOpts)
}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint32)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageCallerSession) Duration() (uint32, error) {
	return _DurationVaultStrategyStorage.Contract.Duration(&_DurationVaultStrategyStorage.CallOpts)
}

// Explanation is a free data retrieval call binding the contract method 0xab5921e1.
//
// Solidity: function explanation() view returns(string)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageCaller) Explanation(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DurationVaultStrategyStorage.contract.Call(opts, &out, "explanation")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Explanation is a free data retrieval call binding the contract method 0xab5921e1.
//
// Solidity: function explanation() view returns(string)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageSession) Explanation() (string, error) {
	return _DurationVaultStrategyStorage.Contract.Explanation(&_DurationVaultStrategyStorage.CallOpts)
}

// Explanation is a free data retrieval call binding the contract method 0xab5921e1.
//
// Solidity: function explanation() view returns(string)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageCallerSession) Explanation() (string, error) {
	return _DurationVaultStrategyStorage.Contract.Explanation(&_DurationVaultStrategyStorage.CallOpts)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() view returns(bool)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageCaller) IsLocked(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DurationVaultStrategyStorage.contract.Call(opts, &out, "isLocked")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() view returns(bool)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageSession) IsLocked() (bool, error) {
	return _DurationVaultStrategyStorage.Contract.IsLocked(&_DurationVaultStrategyStorage.CallOpts)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() view returns(bool)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageCallerSession) IsLocked() (bool, error) {
	return _DurationVaultStrategyStorage.Contract.IsLocked(&_DurationVaultStrategyStorage.CallOpts)
}

// IsMatured is a free data retrieval call binding the contract method 0x7f2b6a0d.
//
// Solidity: function isMatured() view returns(bool)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageCaller) IsMatured(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DurationVaultStrategyStorage.contract.Call(opts, &out, "isMatured")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMatured is a free data retrieval call binding the contract method 0x7f2b6a0d.
//
// Solidity: function isMatured() view returns(bool)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageSession) IsMatured() (bool, error) {
	return _DurationVaultStrategyStorage.Contract.IsMatured(&_DurationVaultStrategyStorage.CallOpts)
}

// IsMatured is a free data retrieval call binding the contract method 0x7f2b6a0d.
//
// Solidity: function isMatured() view returns(bool)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageCallerSession) IsMatured() (bool, error) {
	return _DurationVaultStrategyStorage.Contract.IsMatured(&_DurationVaultStrategyStorage.CallOpts)
}

// LockedAt is a free data retrieval call binding the contract method 0xb2163482.
//
// Solidity: function lockedAt() view returns(uint32)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageCaller) LockedAt(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _DurationVaultStrategyStorage.contract.Call(opts, &out, "lockedAt")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LockedAt is a free data retrieval call binding the contract method 0xb2163482.
//
// Solidity: function lockedAt() view returns(uint32)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageSession) LockedAt() (uint32, error) {
	return _DurationVaultStrategyStorage.Contract.LockedAt(&_DurationVaultStrategyStorage.CallOpts)
}

// LockedAt is a free data retrieval call binding the contract method 0xb2163482.
//
// Solidity: function lockedAt() view returns(uint32)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageCallerSession) LockedAt() (uint32, error) {
	return _DurationVaultStrategyStorage.Contract.LockedAt(&_DurationVaultStrategyStorage.CallOpts)
}

// MaturedAt is a free data retrieval call binding the contract method 0x1f1cc9a3.
//
// Solidity: function maturedAt() view returns(uint32)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageCaller) MaturedAt(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _DurationVaultStrategyStorage.contract.Call(opts, &out, "maturedAt")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MaturedAt is a free data retrieval call binding the contract method 0x1f1cc9a3.
//
// Solidity: function maturedAt() view returns(uint32)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageSession) MaturedAt() (uint32, error) {
	return _DurationVaultStrategyStorage.Contract.MaturedAt(&_DurationVaultStrategyStorage.CallOpts)
}

// MaturedAt is a free data retrieval call binding the contract method 0x1f1cc9a3.
//
// Solidity: function maturedAt() view returns(uint32)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageCallerSession) MaturedAt() (uint32, error) {
	return _DurationVaultStrategyStorage.Contract.MaturedAt(&_DurationVaultStrategyStorage.CallOpts)
}

// MaxPerDeposit is a free data retrieval call binding the contract method 0x43fe08b0.
//
// Solidity: function maxPerDeposit() view returns(uint256)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageCaller) MaxPerDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DurationVaultStrategyStorage.contract.Call(opts, &out, "maxPerDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxPerDeposit is a free data retrieval call binding the contract method 0x43fe08b0.
//
// Solidity: function maxPerDeposit() view returns(uint256)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageSession) MaxPerDeposit() (*big.Int, error) {
	return _DurationVaultStrategyStorage.Contract.MaxPerDeposit(&_DurationVaultStrategyStorage.CallOpts)
}

// MaxPerDeposit is a free data retrieval call binding the contract method 0x43fe08b0.
//
// Solidity: function maxPerDeposit() view returns(uint256)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageCallerSession) MaxPerDeposit() (*big.Int, error) {
	return _DurationVaultStrategyStorage.Contract.MaxPerDeposit(&_DurationVaultStrategyStorage.CallOpts)
}

// MaxTotalDeposits is a free data retrieval call binding the contract method 0x61b01b5d.
//
// Solidity: function maxTotalDeposits() view returns(uint256)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageCaller) MaxTotalDeposits(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DurationVaultStrategyStorage.contract.Call(opts, &out, "maxTotalDeposits")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxTotalDeposits is a free data retrieval call binding the contract method 0x61b01b5d.
//
// Solidity: function maxTotalDeposits() view returns(uint256)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageSession) MaxTotalDeposits() (*big.Int, error) {
	return _DurationVaultStrategyStorage.Contract.MaxTotalDeposits(&_DurationVaultStrategyStorage.CallOpts)
}

// MaxTotalDeposits is a free data retrieval call binding the contract method 0x61b01b5d.
//
// Solidity: function maxTotalDeposits() view returns(uint256)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageCallerSession) MaxTotalDeposits() (*big.Int, error) {
	return _DurationVaultStrategyStorage.Contract.MaxTotalDeposits(&_DurationVaultStrategyStorage.CallOpts)
}

// MetadataURI is a free data retrieval call binding the contract method 0x03ee438c.
//
// Solidity: function metadataURI() view returns(string)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageCaller) MetadataURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DurationVaultStrategyStorage.contract.Call(opts, &out, "metadataURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// MetadataURI is a free data retrieval call binding the contract method 0x03ee438c.
//
// Solidity: function metadataURI() view returns(string)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageSession) MetadataURI() (string, error) {
	return _DurationVaultStrategyStorage.Contract.MetadataURI(&_DurationVaultStrategyStorage.CallOpts)
}

// MetadataURI is a free data retrieval call binding the contract method 0x03ee438c.
//
// Solidity: function metadataURI() view returns(string)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageCallerSession) MetadataURI() (string, error) {
	return _DurationVaultStrategyStorage.Contract.MetadataURI(&_DurationVaultStrategyStorage.CallOpts)
}

// OperatorIntegrationConfigured is a free data retrieval call binding the contract method 0x5438a8c7.
//
// Solidity: function operatorIntegrationConfigured() view returns(bool)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageCaller) OperatorIntegrationConfigured(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DurationVaultStrategyStorage.contract.Call(opts, &out, "operatorIntegrationConfigured")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OperatorIntegrationConfigured is a free data retrieval call binding the contract method 0x5438a8c7.
//
// Solidity: function operatorIntegrationConfigured() view returns(bool)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageSession) OperatorIntegrationConfigured() (bool, error) {
	return _DurationVaultStrategyStorage.Contract.OperatorIntegrationConfigured(&_DurationVaultStrategyStorage.CallOpts)
}

// OperatorIntegrationConfigured is a free data retrieval call binding the contract method 0x5438a8c7.
//
// Solidity: function operatorIntegrationConfigured() view returns(bool)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageCallerSession) OperatorIntegrationConfigured() (bool, error) {
	return _DurationVaultStrategyStorage.Contract.OperatorIntegrationConfigured(&_DurationVaultStrategyStorage.CallOpts)
}

// OperatorSetInfo is a free data retrieval call binding the contract method 0xd4deae81.
//
// Solidity: function operatorSetInfo() view returns(address avs, uint32 operatorSetId)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageCaller) OperatorSetInfo(opts *bind.CallOpts) (struct {
	Avs           common.Address
	OperatorSetId uint32
}, error) {
	var out []interface{}
	err := _DurationVaultStrategyStorage.contract.Call(opts, &out, "operatorSetInfo")

	outstruct := new(struct {
		Avs           common.Address
		OperatorSetId uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Avs = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.OperatorSetId = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

// OperatorSetInfo is a free data retrieval call binding the contract method 0xd4deae81.
//
// Solidity: function operatorSetInfo() view returns(address avs, uint32 operatorSetId)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageSession) OperatorSetInfo() (struct {
	Avs           common.Address
	OperatorSetId uint32
}, error) {
	return _DurationVaultStrategyStorage.Contract.OperatorSetInfo(&_DurationVaultStrategyStorage.CallOpts)
}

// OperatorSetInfo is a free data retrieval call binding the contract method 0xd4deae81.
//
// Solidity: function operatorSetInfo() view returns(address avs, uint32 operatorSetId)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageCallerSession) OperatorSetInfo() (struct {
	Avs           common.Address
	OperatorSetId uint32
}, error) {
	return _DurationVaultStrategyStorage.Contract.OperatorSetInfo(&_DurationVaultStrategyStorage.CallOpts)
}

// OperatorSetRegistered is a free data retrieval call binding the contract method 0x59d915ff.
//
// Solidity: function operatorSetRegistered() view returns(bool)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageCaller) OperatorSetRegistered(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DurationVaultStrategyStorage.contract.Call(opts, &out, "operatorSetRegistered")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OperatorSetRegistered is a free data retrieval call binding the contract method 0x59d915ff.
//
// Solidity: function operatorSetRegistered() view returns(bool)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageSession) OperatorSetRegistered() (bool, error) {
	return _DurationVaultStrategyStorage.Contract.OperatorSetRegistered(&_DurationVaultStrategyStorage.CallOpts)
}

// OperatorSetRegistered is a free data retrieval call binding the contract method 0x59d915ff.
//
// Solidity: function operatorSetRegistered() view returns(bool)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageCallerSession) OperatorSetRegistered() (bool, error) {
	return _DurationVaultStrategyStorage.Contract.OperatorSetRegistered(&_DurationVaultStrategyStorage.CallOpts)
}

// RewardsCoordinator is a free data retrieval call binding the contract method 0x8a2fc4e3.
//
// Solidity: function rewardsCoordinator() view returns(address)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageCaller) RewardsCoordinator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DurationVaultStrategyStorage.contract.Call(opts, &out, "rewardsCoordinator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardsCoordinator is a free data retrieval call binding the contract method 0x8a2fc4e3.
//
// Solidity: function rewardsCoordinator() view returns(address)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageSession) RewardsCoordinator() (common.Address, error) {
	return _DurationVaultStrategyStorage.Contract.RewardsCoordinator(&_DurationVaultStrategyStorage.CallOpts)
}

// RewardsCoordinator is a free data retrieval call binding the contract method 0x8a2fc4e3.
//
// Solidity: function rewardsCoordinator() view returns(address)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageCallerSession) RewardsCoordinator() (common.Address, error) {
	return _DurationVaultStrategyStorage.Contract.RewardsCoordinator(&_DurationVaultStrategyStorage.CallOpts)
}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares(address user) view returns(uint256)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageCaller) Shares(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DurationVaultStrategyStorage.contract.Call(opts, &out, "shares", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares(address user) view returns(uint256)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageSession) Shares(user common.Address) (*big.Int, error) {
	return _DurationVaultStrategyStorage.Contract.Shares(&_DurationVaultStrategyStorage.CallOpts, user)
}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares(address user) view returns(uint256)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageCallerSession) Shares(user common.Address) (*big.Int, error) {
	return _DurationVaultStrategyStorage.Contract.Shares(&_DurationVaultStrategyStorage.CallOpts, user)
}

// SharesToUnderlyingView is a free data retrieval call binding the contract method 0x7a8b2637.
//
// Solidity: function sharesToUnderlyingView(uint256 amountShares) view returns(uint256)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageCaller) SharesToUnderlyingView(opts *bind.CallOpts, amountShares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DurationVaultStrategyStorage.contract.Call(opts, &out, "sharesToUnderlyingView", amountShares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SharesToUnderlyingView is a free data retrieval call binding the contract method 0x7a8b2637.
//
// Solidity: function sharesToUnderlyingView(uint256 amountShares) view returns(uint256)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageSession) SharesToUnderlyingView(amountShares *big.Int) (*big.Int, error) {
	return _DurationVaultStrategyStorage.Contract.SharesToUnderlyingView(&_DurationVaultStrategyStorage.CallOpts, amountShares)
}

// SharesToUnderlyingView is a free data retrieval call binding the contract method 0x7a8b2637.
//
// Solidity: function sharesToUnderlyingView(uint256 amountShares) view returns(uint256)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageCallerSession) SharesToUnderlyingView(amountShares *big.Int) (*big.Int, error) {
	return _DurationVaultStrategyStorage.Contract.SharesToUnderlyingView(&_DurationVaultStrategyStorage.CallOpts, amountShares)
}

// StakeCap is a free data retrieval call binding the contract method 0xba28fd2e.
//
// Solidity: function stakeCap() view returns(uint256)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageCaller) StakeCap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DurationVaultStrategyStorage.contract.Call(opts, &out, "stakeCap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeCap is a free data retrieval call binding the contract method 0xba28fd2e.
//
// Solidity: function stakeCap() view returns(uint256)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageSession) StakeCap() (*big.Int, error) {
	return _DurationVaultStrategyStorage.Contract.StakeCap(&_DurationVaultStrategyStorage.CallOpts)
}

// StakeCap is a free data retrieval call binding the contract method 0xba28fd2e.
//
// Solidity: function stakeCap() view returns(uint256)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageCallerSession) StakeCap() (*big.Int, error) {
	return _DurationVaultStrategyStorage.Contract.StakeCap(&_DurationVaultStrategyStorage.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageCaller) State(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _DurationVaultStrategyStorage.contract.Call(opts, &out, "state")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageSession) State() (uint8, error) {
	return _DurationVaultStrategyStorage.Contract.State(&_DurationVaultStrategyStorage.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageCallerSession) State() (uint8, error) {
	return _DurationVaultStrategyStorage.Contract.State(&_DurationVaultStrategyStorage.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageCaller) TotalShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DurationVaultStrategyStorage.contract.Call(opts, &out, "totalShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageSession) TotalShares() (*big.Int, error) {
	return _DurationVaultStrategyStorage.Contract.TotalShares(&_DurationVaultStrategyStorage.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageCallerSession) TotalShares() (*big.Int, error) {
	return _DurationVaultStrategyStorage.Contract.TotalShares(&_DurationVaultStrategyStorage.CallOpts)
}

// UnderlyingToSharesView is a free data retrieval call binding the contract method 0xe3dae51c.
//
// Solidity: function underlyingToSharesView(uint256 amountUnderlying) view returns(uint256)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageCaller) UnderlyingToSharesView(opts *bind.CallOpts, amountUnderlying *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DurationVaultStrategyStorage.contract.Call(opts, &out, "underlyingToSharesView", amountUnderlying)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnderlyingToSharesView is a free data retrieval call binding the contract method 0xe3dae51c.
//
// Solidity: function underlyingToSharesView(uint256 amountUnderlying) view returns(uint256)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageSession) UnderlyingToSharesView(amountUnderlying *big.Int) (*big.Int, error) {
	return _DurationVaultStrategyStorage.Contract.UnderlyingToSharesView(&_DurationVaultStrategyStorage.CallOpts, amountUnderlying)
}

// UnderlyingToSharesView is a free data retrieval call binding the contract method 0xe3dae51c.
//
// Solidity: function underlyingToSharesView(uint256 amountUnderlying) view returns(uint256)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageCallerSession) UnderlyingToSharesView(amountUnderlying *big.Int) (*big.Int, error) {
	return _DurationVaultStrategyStorage.Contract.UnderlyingToSharesView(&_DurationVaultStrategyStorage.CallOpts, amountUnderlying)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageCaller) UnderlyingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DurationVaultStrategyStorage.contract.Call(opts, &out, "underlyingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageSession) UnderlyingToken() (common.Address, error) {
	return _DurationVaultStrategyStorage.Contract.UnderlyingToken(&_DurationVaultStrategyStorage.CallOpts)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageCallerSession) UnderlyingToken() (common.Address, error) {
	return _DurationVaultStrategyStorage.Contract.UnderlyingToken(&_DurationVaultStrategyStorage.CallOpts)
}

// UnlockAt is a free data retrieval call binding the contract method 0xaa5dec6f.
//
// Solidity: function unlockAt() view returns(uint32)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageCaller) UnlockAt(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _DurationVaultStrategyStorage.contract.Call(opts, &out, "unlockAt")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// UnlockAt is a free data retrieval call binding the contract method 0xaa5dec6f.
//
// Solidity: function unlockAt() view returns(uint32)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageSession) UnlockAt() (uint32, error) {
	return _DurationVaultStrategyStorage.Contract.UnlockAt(&_DurationVaultStrategyStorage.CallOpts)
}

// UnlockAt is a free data retrieval call binding the contract method 0xaa5dec6f.
//
// Solidity: function unlockAt() view returns(uint32)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageCallerSession) UnlockAt() (uint32, error) {
	return _DurationVaultStrategyStorage.Contract.UnlockAt(&_DurationVaultStrategyStorage.CallOpts)
}

// UnlockTimestamp is a free data retrieval call binding the contract method 0xaa082a9d.
//
// Solidity: function unlockTimestamp() view returns(uint32)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageCaller) UnlockTimestamp(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _DurationVaultStrategyStorage.contract.Call(opts, &out, "unlockTimestamp")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// UnlockTimestamp is a free data retrieval call binding the contract method 0xaa082a9d.
//
// Solidity: function unlockTimestamp() view returns(uint32)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageSession) UnlockTimestamp() (uint32, error) {
	return _DurationVaultStrategyStorage.Contract.UnlockTimestamp(&_DurationVaultStrategyStorage.CallOpts)
}

// UnlockTimestamp is a free data retrieval call binding the contract method 0xaa082a9d.
//
// Solidity: function unlockTimestamp() view returns(uint32)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageCallerSession) UnlockTimestamp() (uint32, error) {
	return _DurationVaultStrategyStorage.Contract.UnlockTimestamp(&_DurationVaultStrategyStorage.CallOpts)
}

// UserUnderlyingView is a free data retrieval call binding the contract method 0x553ca5f8.
//
// Solidity: function userUnderlyingView(address user) view returns(uint256)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageCaller) UserUnderlyingView(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DurationVaultStrategyStorage.contract.Call(opts, &out, "userUnderlyingView", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserUnderlyingView is a free data retrieval call binding the contract method 0x553ca5f8.
//
// Solidity: function userUnderlyingView(address user) view returns(uint256)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageSession) UserUnderlyingView(user common.Address) (*big.Int, error) {
	return _DurationVaultStrategyStorage.Contract.UserUnderlyingView(&_DurationVaultStrategyStorage.CallOpts, user)
}

// UserUnderlyingView is a free data retrieval call binding the contract method 0x553ca5f8.
//
// Solidity: function userUnderlyingView(address user) view returns(uint256)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageCallerSession) UserUnderlyingView(user common.Address) (*big.Int, error) {
	return _DurationVaultStrategyStorage.Contract.UserUnderlyingView(&_DurationVaultStrategyStorage.CallOpts, user)
}

// VaultAdmin is a free data retrieval call binding the contract method 0xe7f6f225.
//
// Solidity: function vaultAdmin() view returns(address)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageCaller) VaultAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DurationVaultStrategyStorage.contract.Call(opts, &out, "vaultAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VaultAdmin is a free data retrieval call binding the contract method 0xe7f6f225.
//
// Solidity: function vaultAdmin() view returns(address)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageSession) VaultAdmin() (common.Address, error) {
	return _DurationVaultStrategyStorage.Contract.VaultAdmin(&_DurationVaultStrategyStorage.CallOpts)
}

// VaultAdmin is a free data retrieval call binding the contract method 0xe7f6f225.
//
// Solidity: function vaultAdmin() view returns(address)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageCallerSession) VaultAdmin() (common.Address, error) {
	return _DurationVaultStrategyStorage.Contract.VaultAdmin(&_DurationVaultStrategyStorage.CallOpts)
}

// WithdrawalsOpen is a free data retrieval call binding the contract method 0x94aad677.
//
// Solidity: function withdrawalsOpen() view returns(bool)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageCaller) WithdrawalsOpen(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DurationVaultStrategyStorage.contract.Call(opts, &out, "withdrawalsOpen")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WithdrawalsOpen is a free data retrieval call binding the contract method 0x94aad677.
//
// Solidity: function withdrawalsOpen() view returns(bool)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageSession) WithdrawalsOpen() (bool, error) {
	return _DurationVaultStrategyStorage.Contract.WithdrawalsOpen(&_DurationVaultStrategyStorage.CallOpts)
}

// WithdrawalsOpen is a free data retrieval call binding the contract method 0x94aad677.
//
// Solidity: function withdrawalsOpen() view returns(bool)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageCallerSession) WithdrawalsOpen() (bool, error) {
	return _DurationVaultStrategyStorage.Contract.WithdrawalsOpen(&_DurationVaultStrategyStorage.CallOpts)
}

// AdvanceToWithdrawals is a paid mutator transaction binding the contract method 0x6325f655.
//
// Solidity: function advanceToWithdrawals() returns()
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageTransactor) AdvanceToWithdrawals(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DurationVaultStrategyStorage.contract.Transact(opts, "advanceToWithdrawals")
}

// AdvanceToWithdrawals is a paid mutator transaction binding the contract method 0x6325f655.
//
// Solidity: function advanceToWithdrawals() returns()
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageSession) AdvanceToWithdrawals() (*types.Transaction, error) {
	return _DurationVaultStrategyStorage.Contract.AdvanceToWithdrawals(&_DurationVaultStrategyStorage.TransactOpts)
}

// AdvanceToWithdrawals is a paid mutator transaction binding the contract method 0x6325f655.
//
// Solidity: function advanceToWithdrawals() returns()
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageTransactorSession) AdvanceToWithdrawals() (*types.Transaction, error) {
	return _DurationVaultStrategyStorage.Contract.AdvanceToWithdrawals(&_DurationVaultStrategyStorage.TransactOpts)
}

// BeforeAddShares is a paid mutator transaction binding the contract method 0x73e3c280.
//
// Solidity: function beforeAddShares(address staker, uint256 shares) returns()
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageTransactor) BeforeAddShares(opts *bind.TransactOpts, staker common.Address, shares *big.Int) (*types.Transaction, error) {
	return _DurationVaultStrategyStorage.contract.Transact(opts, "beforeAddShares", staker, shares)
}

// BeforeAddShares is a paid mutator transaction binding the contract method 0x73e3c280.
//
// Solidity: function beforeAddShares(address staker, uint256 shares) returns()
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageSession) BeforeAddShares(staker common.Address, shares *big.Int) (*types.Transaction, error) {
	return _DurationVaultStrategyStorage.Contract.BeforeAddShares(&_DurationVaultStrategyStorage.TransactOpts, staker, shares)
}

// BeforeAddShares is a paid mutator transaction binding the contract method 0x73e3c280.
//
// Solidity: function beforeAddShares(address staker, uint256 shares) returns()
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageTransactorSession) BeforeAddShares(staker common.Address, shares *big.Int) (*types.Transaction, error) {
	return _DurationVaultStrategyStorage.Contract.BeforeAddShares(&_DurationVaultStrategyStorage.TransactOpts, staker, shares)
}

// BeforeRemoveShares is a paid mutator transaction binding the contract method 0x03e3e6eb.
//
// Solidity: function beforeRemoveShares(address staker, uint256 shares) returns()
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageTransactor) BeforeRemoveShares(opts *bind.TransactOpts, staker common.Address, shares *big.Int) (*types.Transaction, error) {
	return _DurationVaultStrategyStorage.contract.Transact(opts, "beforeRemoveShares", staker, shares)
}

// BeforeRemoveShares is a paid mutator transaction binding the contract method 0x03e3e6eb.
//
// Solidity: function beforeRemoveShares(address staker, uint256 shares) returns()
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageSession) BeforeRemoveShares(staker common.Address, shares *big.Int) (*types.Transaction, error) {
	return _DurationVaultStrategyStorage.Contract.BeforeRemoveShares(&_DurationVaultStrategyStorage.TransactOpts, staker, shares)
}

// BeforeRemoveShares is a paid mutator transaction binding the contract method 0x03e3e6eb.
//
// Solidity: function beforeRemoveShares(address staker, uint256 shares) returns()
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageTransactorSession) BeforeRemoveShares(staker common.Address, shares *big.Int) (*types.Transaction, error) {
	return _DurationVaultStrategyStorage.Contract.BeforeRemoveShares(&_DurationVaultStrategyStorage.TransactOpts, staker, shares)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns(uint256)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageTransactor) Deposit(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DurationVaultStrategyStorage.contract.Transact(opts, "deposit", token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns(uint256)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageSession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DurationVaultStrategyStorage.Contract.Deposit(&_DurationVaultStrategyStorage.TransactOpts, token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns(uint256)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageTransactorSession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DurationVaultStrategyStorage.Contract.Deposit(&_DurationVaultStrategyStorage.TransactOpts, token, amount)
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageTransactor) Lock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DurationVaultStrategyStorage.contract.Transact(opts, "lock")
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageSession) Lock() (*types.Transaction, error) {
	return _DurationVaultStrategyStorage.Contract.Lock(&_DurationVaultStrategyStorage.TransactOpts)
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageTransactorSession) Lock() (*types.Transaction, error) {
	return _DurationVaultStrategyStorage.Contract.Lock(&_DurationVaultStrategyStorage.TransactOpts)
}

// MarkMatured is a paid mutator transaction binding the contract method 0x6d8690a9.
//
// Solidity: function markMatured() returns()
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageTransactor) MarkMatured(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DurationVaultStrategyStorage.contract.Transact(opts, "markMatured")
}

// MarkMatured is a paid mutator transaction binding the contract method 0x6d8690a9.
//
// Solidity: function markMatured() returns()
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageSession) MarkMatured() (*types.Transaction, error) {
	return _DurationVaultStrategyStorage.Contract.MarkMatured(&_DurationVaultStrategyStorage.TransactOpts)
}

// MarkMatured is a paid mutator transaction binding the contract method 0x6d8690a9.
//
// Solidity: function markMatured() returns()
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageTransactorSession) MarkMatured() (*types.Transaction, error) {
	return _DurationVaultStrategyStorage.Contract.MarkMatured(&_DurationVaultStrategyStorage.TransactOpts)
}

// SetRewardsClaimer is a paid mutator transaction binding the contract method 0xb501d660.
//
// Solidity: function setRewardsClaimer(address claimer) returns()
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageTransactor) SetRewardsClaimer(opts *bind.TransactOpts, claimer common.Address) (*types.Transaction, error) {
	return _DurationVaultStrategyStorage.contract.Transact(opts, "setRewardsClaimer", claimer)
}

// SetRewardsClaimer is a paid mutator transaction binding the contract method 0xb501d660.
//
// Solidity: function setRewardsClaimer(address claimer) returns()
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageSession) SetRewardsClaimer(claimer common.Address) (*types.Transaction, error) {
	return _DurationVaultStrategyStorage.Contract.SetRewardsClaimer(&_DurationVaultStrategyStorage.TransactOpts, claimer)
}

// SetRewardsClaimer is a paid mutator transaction binding the contract method 0xb501d660.
//
// Solidity: function setRewardsClaimer(address claimer) returns()
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageTransactorSession) SetRewardsClaimer(claimer common.Address) (*types.Transaction, error) {
	return _DurationVaultStrategyStorage.Contract.SetRewardsClaimer(&_DurationVaultStrategyStorage.TransactOpts, claimer)
}

// SharesToUnderlying is a paid mutator transaction binding the contract method 0xf3e73875.
//
// Solidity: function sharesToUnderlying(uint256 amountShares) returns(uint256)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageTransactor) SharesToUnderlying(opts *bind.TransactOpts, amountShares *big.Int) (*types.Transaction, error) {
	return _DurationVaultStrategyStorage.contract.Transact(opts, "sharesToUnderlying", amountShares)
}

// SharesToUnderlying is a paid mutator transaction binding the contract method 0xf3e73875.
//
// Solidity: function sharesToUnderlying(uint256 amountShares) returns(uint256)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageSession) SharesToUnderlying(amountShares *big.Int) (*types.Transaction, error) {
	return _DurationVaultStrategyStorage.Contract.SharesToUnderlying(&_DurationVaultStrategyStorage.TransactOpts, amountShares)
}

// SharesToUnderlying is a paid mutator transaction binding the contract method 0xf3e73875.
//
// Solidity: function sharesToUnderlying(uint256 amountShares) returns(uint256)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageTransactorSession) SharesToUnderlying(amountShares *big.Int) (*types.Transaction, error) {
	return _DurationVaultStrategyStorage.Contract.SharesToUnderlying(&_DurationVaultStrategyStorage.TransactOpts, amountShares)
}

// UnderlyingToShares is a paid mutator transaction binding the contract method 0x8c871019.
//
// Solidity: function underlyingToShares(uint256 amountUnderlying) returns(uint256)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageTransactor) UnderlyingToShares(opts *bind.TransactOpts, amountUnderlying *big.Int) (*types.Transaction, error) {
	return _DurationVaultStrategyStorage.contract.Transact(opts, "underlyingToShares", amountUnderlying)
}

// UnderlyingToShares is a paid mutator transaction binding the contract method 0x8c871019.
//
// Solidity: function underlyingToShares(uint256 amountUnderlying) returns(uint256)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageSession) UnderlyingToShares(amountUnderlying *big.Int) (*types.Transaction, error) {
	return _DurationVaultStrategyStorage.Contract.UnderlyingToShares(&_DurationVaultStrategyStorage.TransactOpts, amountUnderlying)
}

// UnderlyingToShares is a paid mutator transaction binding the contract method 0x8c871019.
//
// Solidity: function underlyingToShares(uint256 amountUnderlying) returns(uint256)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageTransactorSession) UnderlyingToShares(amountUnderlying *big.Int) (*types.Transaction, error) {
	return _DurationVaultStrategyStorage.Contract.UnderlyingToShares(&_DurationVaultStrategyStorage.TransactOpts, amountUnderlying)
}

// UpdateDelegationApprover is a paid mutator transaction binding the contract method 0xb4e20f13.
//
// Solidity: function updateDelegationApprover(address newDelegationApprover) returns()
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageTransactor) UpdateDelegationApprover(opts *bind.TransactOpts, newDelegationApprover common.Address) (*types.Transaction, error) {
	return _DurationVaultStrategyStorage.contract.Transact(opts, "updateDelegationApprover", newDelegationApprover)
}

// UpdateDelegationApprover is a paid mutator transaction binding the contract method 0xb4e20f13.
//
// Solidity: function updateDelegationApprover(address newDelegationApprover) returns()
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageSession) UpdateDelegationApprover(newDelegationApprover common.Address) (*types.Transaction, error) {
	return _DurationVaultStrategyStorage.Contract.UpdateDelegationApprover(&_DurationVaultStrategyStorage.TransactOpts, newDelegationApprover)
}

// UpdateDelegationApprover is a paid mutator transaction binding the contract method 0xb4e20f13.
//
// Solidity: function updateDelegationApprover(address newDelegationApprover) returns()
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageTransactorSession) UpdateDelegationApprover(newDelegationApprover common.Address) (*types.Transaction, error) {
	return _DurationVaultStrategyStorage.Contract.UpdateDelegationApprover(&_DurationVaultStrategyStorage.TransactOpts, newDelegationApprover)
}

// UpdateMetadataURI is a paid mutator transaction binding the contract method 0x53fd3e81.
//
// Solidity: function updateMetadataURI(string newMetadataURI) returns()
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageTransactor) UpdateMetadataURI(opts *bind.TransactOpts, newMetadataURI string) (*types.Transaction, error) {
	return _DurationVaultStrategyStorage.contract.Transact(opts, "updateMetadataURI", newMetadataURI)
}

// UpdateMetadataURI is a paid mutator transaction binding the contract method 0x53fd3e81.
//
// Solidity: function updateMetadataURI(string newMetadataURI) returns()
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageSession) UpdateMetadataURI(newMetadataURI string) (*types.Transaction, error) {
	return _DurationVaultStrategyStorage.Contract.UpdateMetadataURI(&_DurationVaultStrategyStorage.TransactOpts, newMetadataURI)
}

// UpdateMetadataURI is a paid mutator transaction binding the contract method 0x53fd3e81.
//
// Solidity: function updateMetadataURI(string newMetadataURI) returns()
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageTransactorSession) UpdateMetadataURI(newMetadataURI string) (*types.Transaction, error) {
	return _DurationVaultStrategyStorage.Contract.UpdateMetadataURI(&_DurationVaultStrategyStorage.TransactOpts, newMetadataURI)
}

// UpdateOperatorMetadataURI is a paid mutator transaction binding the contract method 0x99be81c8.
//
// Solidity: function updateOperatorMetadataURI(string newOperatorMetadataURI) returns()
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageTransactor) UpdateOperatorMetadataURI(opts *bind.TransactOpts, newOperatorMetadataURI string) (*types.Transaction, error) {
	return _DurationVaultStrategyStorage.contract.Transact(opts, "updateOperatorMetadataURI", newOperatorMetadataURI)
}

// UpdateOperatorMetadataURI is a paid mutator transaction binding the contract method 0x99be81c8.
//
// Solidity: function updateOperatorMetadataURI(string newOperatorMetadataURI) returns()
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageSession) UpdateOperatorMetadataURI(newOperatorMetadataURI string) (*types.Transaction, error) {
	return _DurationVaultStrategyStorage.Contract.UpdateOperatorMetadataURI(&_DurationVaultStrategyStorage.TransactOpts, newOperatorMetadataURI)
}

// UpdateOperatorMetadataURI is a paid mutator transaction binding the contract method 0x99be81c8.
//
// Solidity: function updateOperatorMetadataURI(string newOperatorMetadataURI) returns()
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageTransactorSession) UpdateOperatorMetadataURI(newOperatorMetadataURI string) (*types.Transaction, error) {
	return _DurationVaultStrategyStorage.Contract.UpdateOperatorMetadataURI(&_DurationVaultStrategyStorage.TransactOpts, newOperatorMetadataURI)
}

// UpdateTVLLimits is a paid mutator transaction binding the contract method 0xaf6eb2be.
//
// Solidity: function updateTVLLimits(uint256 newMaxPerDeposit, uint256 newStakeCap) returns()
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageTransactor) UpdateTVLLimits(opts *bind.TransactOpts, newMaxPerDeposit *big.Int, newStakeCap *big.Int) (*types.Transaction, error) {
	return _DurationVaultStrategyStorage.contract.Transact(opts, "updateTVLLimits", newMaxPerDeposit, newStakeCap)
}

// UpdateTVLLimits is a paid mutator transaction binding the contract method 0xaf6eb2be.
//
// Solidity: function updateTVLLimits(uint256 newMaxPerDeposit, uint256 newStakeCap) returns()
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageSession) UpdateTVLLimits(newMaxPerDeposit *big.Int, newStakeCap *big.Int) (*types.Transaction, error) {
	return _DurationVaultStrategyStorage.Contract.UpdateTVLLimits(&_DurationVaultStrategyStorage.TransactOpts, newMaxPerDeposit, newStakeCap)
}

// UpdateTVLLimits is a paid mutator transaction binding the contract method 0xaf6eb2be.
//
// Solidity: function updateTVLLimits(uint256 newMaxPerDeposit, uint256 newStakeCap) returns()
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageTransactorSession) UpdateTVLLimits(newMaxPerDeposit *big.Int, newStakeCap *big.Int) (*types.Transaction, error) {
	return _DurationVaultStrategyStorage.Contract.UpdateTVLLimits(&_DurationVaultStrategyStorage.TransactOpts, newMaxPerDeposit, newStakeCap)
}

// UserUnderlying is a paid mutator transaction binding the contract method 0x8f6a6240.
//
// Solidity: function userUnderlying(address user) returns(uint256)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageTransactor) UserUnderlying(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _DurationVaultStrategyStorage.contract.Transact(opts, "userUnderlying", user)
}

// UserUnderlying is a paid mutator transaction binding the contract method 0x8f6a6240.
//
// Solidity: function userUnderlying(address user) returns(uint256)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageSession) UserUnderlying(user common.Address) (*types.Transaction, error) {
	return _DurationVaultStrategyStorage.Contract.UserUnderlying(&_DurationVaultStrategyStorage.TransactOpts, user)
}

// UserUnderlying is a paid mutator transaction binding the contract method 0x8f6a6240.
//
// Solidity: function userUnderlying(address user) returns(uint256)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageTransactorSession) UserUnderlying(user common.Address) (*types.Transaction, error) {
	return _DurationVaultStrategyStorage.Contract.UserUnderlying(&_DurationVaultStrategyStorage.TransactOpts, user)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address recipient, address token, uint256 amountShares) returns(uint256)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageTransactor) Withdraw(opts *bind.TransactOpts, recipient common.Address, token common.Address, amountShares *big.Int) (*types.Transaction, error) {
	return _DurationVaultStrategyStorage.contract.Transact(opts, "withdraw", recipient, token, amountShares)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address recipient, address token, uint256 amountShares) returns(uint256)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageSession) Withdraw(recipient common.Address, token common.Address, amountShares *big.Int) (*types.Transaction, error) {
	return _DurationVaultStrategyStorage.Contract.Withdraw(&_DurationVaultStrategyStorage.TransactOpts, recipient, token, amountShares)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address recipient, address token, uint256 amountShares) returns(uint256)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageTransactorSession) Withdraw(recipient common.Address, token common.Address, amountShares *big.Int) (*types.Transaction, error) {
	return _DurationVaultStrategyStorage.Contract.Withdraw(&_DurationVaultStrategyStorage.TransactOpts, recipient, token, amountShares)
}

// DurationVaultStrategyStorageDeallocateAttemptedIterator is returned from FilterDeallocateAttempted and is used to iterate over the raw logs and unpacked data for DeallocateAttempted events raised by the DurationVaultStrategyStorage contract.
type DurationVaultStrategyStorageDeallocateAttemptedIterator struct {
	Event *DurationVaultStrategyStorageDeallocateAttempted // Event containing the contract specifics and raw log

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
func (it *DurationVaultStrategyStorageDeallocateAttemptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DurationVaultStrategyStorageDeallocateAttempted)
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
		it.Event = new(DurationVaultStrategyStorageDeallocateAttempted)
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
func (it *DurationVaultStrategyStorageDeallocateAttemptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DurationVaultStrategyStorageDeallocateAttemptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DurationVaultStrategyStorageDeallocateAttempted represents a DeallocateAttempted event raised by the DurationVaultStrategyStorage contract.
type DurationVaultStrategyStorageDeallocateAttempted struct {
	Success bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeallocateAttempted is a free log retrieval operation binding the contract event 0x72f957da7daaea6b52e4ff7820cb464206fd51e9f502f3027f45b5017caf4c8b.
//
// Solidity: event DeallocateAttempted(bool success)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageFilterer) FilterDeallocateAttempted(opts *bind.FilterOpts) (*DurationVaultStrategyStorageDeallocateAttemptedIterator, error) {

	logs, sub, err := _DurationVaultStrategyStorage.contract.FilterLogs(opts, "DeallocateAttempted")
	if err != nil {
		return nil, err
	}
	return &DurationVaultStrategyStorageDeallocateAttemptedIterator{contract: _DurationVaultStrategyStorage.contract, event: "DeallocateAttempted", logs: logs, sub: sub}, nil
}

// WatchDeallocateAttempted is a free log subscription operation binding the contract event 0x72f957da7daaea6b52e4ff7820cb464206fd51e9f502f3027f45b5017caf4c8b.
//
// Solidity: event DeallocateAttempted(bool success)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageFilterer) WatchDeallocateAttempted(opts *bind.WatchOpts, sink chan<- *DurationVaultStrategyStorageDeallocateAttempted) (event.Subscription, error) {

	logs, sub, err := _DurationVaultStrategyStorage.contract.WatchLogs(opts, "DeallocateAttempted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DurationVaultStrategyStorageDeallocateAttempted)
				if err := _DurationVaultStrategyStorage.contract.UnpackLog(event, "DeallocateAttempted", log); err != nil {
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

// ParseDeallocateAttempted is a log parse operation binding the contract event 0x72f957da7daaea6b52e4ff7820cb464206fd51e9f502f3027f45b5017caf4c8b.
//
// Solidity: event DeallocateAttempted(bool success)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageFilterer) ParseDeallocateAttempted(log types.Log) (*DurationVaultStrategyStorageDeallocateAttempted, error) {
	event := new(DurationVaultStrategyStorageDeallocateAttempted)
	if err := _DurationVaultStrategyStorage.contract.UnpackLog(event, "DeallocateAttempted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DurationVaultStrategyStorageDeregisterAttemptedIterator is returned from FilterDeregisterAttempted and is used to iterate over the raw logs and unpacked data for DeregisterAttempted events raised by the DurationVaultStrategyStorage contract.
type DurationVaultStrategyStorageDeregisterAttemptedIterator struct {
	Event *DurationVaultStrategyStorageDeregisterAttempted // Event containing the contract specifics and raw log

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
func (it *DurationVaultStrategyStorageDeregisterAttemptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DurationVaultStrategyStorageDeregisterAttempted)
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
		it.Event = new(DurationVaultStrategyStorageDeregisterAttempted)
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
func (it *DurationVaultStrategyStorageDeregisterAttemptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DurationVaultStrategyStorageDeregisterAttemptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DurationVaultStrategyStorageDeregisterAttempted represents a DeregisterAttempted event raised by the DurationVaultStrategyStorage contract.
type DurationVaultStrategyStorageDeregisterAttempted struct {
	Success bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeregisterAttempted is a free log retrieval operation binding the contract event 0xd0791dbc9180cb64588d7eb7658a1022dcf734b8825eb7eec68bd9516872d168.
//
// Solidity: event DeregisterAttempted(bool success)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageFilterer) FilterDeregisterAttempted(opts *bind.FilterOpts) (*DurationVaultStrategyStorageDeregisterAttemptedIterator, error) {

	logs, sub, err := _DurationVaultStrategyStorage.contract.FilterLogs(opts, "DeregisterAttempted")
	if err != nil {
		return nil, err
	}
	return &DurationVaultStrategyStorageDeregisterAttemptedIterator{contract: _DurationVaultStrategyStorage.contract, event: "DeregisterAttempted", logs: logs, sub: sub}, nil
}

// WatchDeregisterAttempted is a free log subscription operation binding the contract event 0xd0791dbc9180cb64588d7eb7658a1022dcf734b8825eb7eec68bd9516872d168.
//
// Solidity: event DeregisterAttempted(bool success)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageFilterer) WatchDeregisterAttempted(opts *bind.WatchOpts, sink chan<- *DurationVaultStrategyStorageDeregisterAttempted) (event.Subscription, error) {

	logs, sub, err := _DurationVaultStrategyStorage.contract.WatchLogs(opts, "DeregisterAttempted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DurationVaultStrategyStorageDeregisterAttempted)
				if err := _DurationVaultStrategyStorage.contract.UnpackLog(event, "DeregisterAttempted", log); err != nil {
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

// ParseDeregisterAttempted is a log parse operation binding the contract event 0xd0791dbc9180cb64588d7eb7658a1022dcf734b8825eb7eec68bd9516872d168.
//
// Solidity: event DeregisterAttempted(bool success)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageFilterer) ParseDeregisterAttempted(log types.Log) (*DurationVaultStrategyStorageDeregisterAttempted, error) {
	event := new(DurationVaultStrategyStorageDeregisterAttempted)
	if err := _DurationVaultStrategyStorage.contract.UnpackLog(event, "DeregisterAttempted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DurationVaultStrategyStorageExchangeRateEmittedIterator is returned from FilterExchangeRateEmitted and is used to iterate over the raw logs and unpacked data for ExchangeRateEmitted events raised by the DurationVaultStrategyStorage contract.
type DurationVaultStrategyStorageExchangeRateEmittedIterator struct {
	Event *DurationVaultStrategyStorageExchangeRateEmitted // Event containing the contract specifics and raw log

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
func (it *DurationVaultStrategyStorageExchangeRateEmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DurationVaultStrategyStorageExchangeRateEmitted)
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
		it.Event = new(DurationVaultStrategyStorageExchangeRateEmitted)
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
func (it *DurationVaultStrategyStorageExchangeRateEmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DurationVaultStrategyStorageExchangeRateEmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DurationVaultStrategyStorageExchangeRateEmitted represents a ExchangeRateEmitted event raised by the DurationVaultStrategyStorage contract.
type DurationVaultStrategyStorageExchangeRateEmitted struct {
	Rate *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterExchangeRateEmitted is a free log retrieval operation binding the contract event 0xd2494f3479e5da49d386657c292c610b5b01df313d07c62eb0cfa49924a31be8.
//
// Solidity: event ExchangeRateEmitted(uint256 rate)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageFilterer) FilterExchangeRateEmitted(opts *bind.FilterOpts) (*DurationVaultStrategyStorageExchangeRateEmittedIterator, error) {

	logs, sub, err := _DurationVaultStrategyStorage.contract.FilterLogs(opts, "ExchangeRateEmitted")
	if err != nil {
		return nil, err
	}
	return &DurationVaultStrategyStorageExchangeRateEmittedIterator{contract: _DurationVaultStrategyStorage.contract, event: "ExchangeRateEmitted", logs: logs, sub: sub}, nil
}

// WatchExchangeRateEmitted is a free log subscription operation binding the contract event 0xd2494f3479e5da49d386657c292c610b5b01df313d07c62eb0cfa49924a31be8.
//
// Solidity: event ExchangeRateEmitted(uint256 rate)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageFilterer) WatchExchangeRateEmitted(opts *bind.WatchOpts, sink chan<- *DurationVaultStrategyStorageExchangeRateEmitted) (event.Subscription, error) {

	logs, sub, err := _DurationVaultStrategyStorage.contract.WatchLogs(opts, "ExchangeRateEmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DurationVaultStrategyStorageExchangeRateEmitted)
				if err := _DurationVaultStrategyStorage.contract.UnpackLog(event, "ExchangeRateEmitted", log); err != nil {
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

// ParseExchangeRateEmitted is a log parse operation binding the contract event 0xd2494f3479e5da49d386657c292c610b5b01df313d07c62eb0cfa49924a31be8.
//
// Solidity: event ExchangeRateEmitted(uint256 rate)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageFilterer) ParseExchangeRateEmitted(log types.Log) (*DurationVaultStrategyStorageExchangeRateEmitted, error) {
	event := new(DurationVaultStrategyStorageExchangeRateEmitted)
	if err := _DurationVaultStrategyStorage.contract.UnpackLog(event, "ExchangeRateEmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DurationVaultStrategyStorageMaxPerDepositUpdatedIterator is returned from FilterMaxPerDepositUpdated and is used to iterate over the raw logs and unpacked data for MaxPerDepositUpdated events raised by the DurationVaultStrategyStorage contract.
type DurationVaultStrategyStorageMaxPerDepositUpdatedIterator struct {
	Event *DurationVaultStrategyStorageMaxPerDepositUpdated // Event containing the contract specifics and raw log

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
func (it *DurationVaultStrategyStorageMaxPerDepositUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DurationVaultStrategyStorageMaxPerDepositUpdated)
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
		it.Event = new(DurationVaultStrategyStorageMaxPerDepositUpdated)
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
func (it *DurationVaultStrategyStorageMaxPerDepositUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DurationVaultStrategyStorageMaxPerDepositUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DurationVaultStrategyStorageMaxPerDepositUpdated represents a MaxPerDepositUpdated event raised by the DurationVaultStrategyStorage contract.
type DurationVaultStrategyStorageMaxPerDepositUpdated struct {
	PreviousValue *big.Int
	NewValue      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMaxPerDepositUpdated is a free log retrieval operation binding the contract event 0xf97ed4e083acac67830025ecbc756d8fe847cdbdca4cee3fe1e128e98b54ecb5.
//
// Solidity: event MaxPerDepositUpdated(uint256 previousValue, uint256 newValue)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageFilterer) FilterMaxPerDepositUpdated(opts *bind.FilterOpts) (*DurationVaultStrategyStorageMaxPerDepositUpdatedIterator, error) {

	logs, sub, err := _DurationVaultStrategyStorage.contract.FilterLogs(opts, "MaxPerDepositUpdated")
	if err != nil {
		return nil, err
	}
	return &DurationVaultStrategyStorageMaxPerDepositUpdatedIterator{contract: _DurationVaultStrategyStorage.contract, event: "MaxPerDepositUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxPerDepositUpdated is a free log subscription operation binding the contract event 0xf97ed4e083acac67830025ecbc756d8fe847cdbdca4cee3fe1e128e98b54ecb5.
//
// Solidity: event MaxPerDepositUpdated(uint256 previousValue, uint256 newValue)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageFilterer) WatchMaxPerDepositUpdated(opts *bind.WatchOpts, sink chan<- *DurationVaultStrategyStorageMaxPerDepositUpdated) (event.Subscription, error) {

	logs, sub, err := _DurationVaultStrategyStorage.contract.WatchLogs(opts, "MaxPerDepositUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DurationVaultStrategyStorageMaxPerDepositUpdated)
				if err := _DurationVaultStrategyStorage.contract.UnpackLog(event, "MaxPerDepositUpdated", log); err != nil {
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

// ParseMaxPerDepositUpdated is a log parse operation binding the contract event 0xf97ed4e083acac67830025ecbc756d8fe847cdbdca4cee3fe1e128e98b54ecb5.
//
// Solidity: event MaxPerDepositUpdated(uint256 previousValue, uint256 newValue)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageFilterer) ParseMaxPerDepositUpdated(log types.Log) (*DurationVaultStrategyStorageMaxPerDepositUpdated, error) {
	event := new(DurationVaultStrategyStorageMaxPerDepositUpdated)
	if err := _DurationVaultStrategyStorage.contract.UnpackLog(event, "MaxPerDepositUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DurationVaultStrategyStorageMaxTotalDepositsUpdatedIterator is returned from FilterMaxTotalDepositsUpdated and is used to iterate over the raw logs and unpacked data for MaxTotalDepositsUpdated events raised by the DurationVaultStrategyStorage contract.
type DurationVaultStrategyStorageMaxTotalDepositsUpdatedIterator struct {
	Event *DurationVaultStrategyStorageMaxTotalDepositsUpdated // Event containing the contract specifics and raw log

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
func (it *DurationVaultStrategyStorageMaxTotalDepositsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DurationVaultStrategyStorageMaxTotalDepositsUpdated)
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
		it.Event = new(DurationVaultStrategyStorageMaxTotalDepositsUpdated)
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
func (it *DurationVaultStrategyStorageMaxTotalDepositsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DurationVaultStrategyStorageMaxTotalDepositsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DurationVaultStrategyStorageMaxTotalDepositsUpdated represents a MaxTotalDepositsUpdated event raised by the DurationVaultStrategyStorage contract.
type DurationVaultStrategyStorageMaxTotalDepositsUpdated struct {
	PreviousValue *big.Int
	NewValue      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMaxTotalDepositsUpdated is a free log retrieval operation binding the contract event 0x6ab181e0440bfbf4bacdf2e99674735ce6638005490688c5f994f5399353e452.
//
// Solidity: event MaxTotalDepositsUpdated(uint256 previousValue, uint256 newValue)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageFilterer) FilterMaxTotalDepositsUpdated(opts *bind.FilterOpts) (*DurationVaultStrategyStorageMaxTotalDepositsUpdatedIterator, error) {

	logs, sub, err := _DurationVaultStrategyStorage.contract.FilterLogs(opts, "MaxTotalDepositsUpdated")
	if err != nil {
		return nil, err
	}
	return &DurationVaultStrategyStorageMaxTotalDepositsUpdatedIterator{contract: _DurationVaultStrategyStorage.contract, event: "MaxTotalDepositsUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxTotalDepositsUpdated is a free log subscription operation binding the contract event 0x6ab181e0440bfbf4bacdf2e99674735ce6638005490688c5f994f5399353e452.
//
// Solidity: event MaxTotalDepositsUpdated(uint256 previousValue, uint256 newValue)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageFilterer) WatchMaxTotalDepositsUpdated(opts *bind.WatchOpts, sink chan<- *DurationVaultStrategyStorageMaxTotalDepositsUpdated) (event.Subscription, error) {

	logs, sub, err := _DurationVaultStrategyStorage.contract.WatchLogs(opts, "MaxTotalDepositsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DurationVaultStrategyStorageMaxTotalDepositsUpdated)
				if err := _DurationVaultStrategyStorage.contract.UnpackLog(event, "MaxTotalDepositsUpdated", log); err != nil {
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

// ParseMaxTotalDepositsUpdated is a log parse operation binding the contract event 0x6ab181e0440bfbf4bacdf2e99674735ce6638005490688c5f994f5399353e452.
//
// Solidity: event MaxTotalDepositsUpdated(uint256 previousValue, uint256 newValue)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageFilterer) ParseMaxTotalDepositsUpdated(log types.Log) (*DurationVaultStrategyStorageMaxTotalDepositsUpdated, error) {
	event := new(DurationVaultStrategyStorageMaxTotalDepositsUpdated)
	if err := _DurationVaultStrategyStorage.contract.UnpackLog(event, "MaxTotalDepositsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DurationVaultStrategyStorageMetadataURIUpdatedIterator is returned from FilterMetadataURIUpdated and is used to iterate over the raw logs and unpacked data for MetadataURIUpdated events raised by the DurationVaultStrategyStorage contract.
type DurationVaultStrategyStorageMetadataURIUpdatedIterator struct {
	Event *DurationVaultStrategyStorageMetadataURIUpdated // Event containing the contract specifics and raw log

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
func (it *DurationVaultStrategyStorageMetadataURIUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DurationVaultStrategyStorageMetadataURIUpdated)
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
		it.Event = new(DurationVaultStrategyStorageMetadataURIUpdated)
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
func (it *DurationVaultStrategyStorageMetadataURIUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DurationVaultStrategyStorageMetadataURIUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DurationVaultStrategyStorageMetadataURIUpdated represents a MetadataURIUpdated event raised by the DurationVaultStrategyStorage contract.
type DurationVaultStrategyStorageMetadataURIUpdated struct {
	NewMetadataURI string
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterMetadataURIUpdated is a free log retrieval operation binding the contract event 0xefafb90526da1636e1335eac0151301742fb755d986954c613b90e891778ba39.
//
// Solidity: event MetadataURIUpdated(string newMetadataURI)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageFilterer) FilterMetadataURIUpdated(opts *bind.FilterOpts) (*DurationVaultStrategyStorageMetadataURIUpdatedIterator, error) {

	logs, sub, err := _DurationVaultStrategyStorage.contract.FilterLogs(opts, "MetadataURIUpdated")
	if err != nil {
		return nil, err
	}
	return &DurationVaultStrategyStorageMetadataURIUpdatedIterator{contract: _DurationVaultStrategyStorage.contract, event: "MetadataURIUpdated", logs: logs, sub: sub}, nil
}

// WatchMetadataURIUpdated is a free log subscription operation binding the contract event 0xefafb90526da1636e1335eac0151301742fb755d986954c613b90e891778ba39.
//
// Solidity: event MetadataURIUpdated(string newMetadataURI)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageFilterer) WatchMetadataURIUpdated(opts *bind.WatchOpts, sink chan<- *DurationVaultStrategyStorageMetadataURIUpdated) (event.Subscription, error) {

	logs, sub, err := _DurationVaultStrategyStorage.contract.WatchLogs(opts, "MetadataURIUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DurationVaultStrategyStorageMetadataURIUpdated)
				if err := _DurationVaultStrategyStorage.contract.UnpackLog(event, "MetadataURIUpdated", log); err != nil {
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

// ParseMetadataURIUpdated is a log parse operation binding the contract event 0xefafb90526da1636e1335eac0151301742fb755d986954c613b90e891778ba39.
//
// Solidity: event MetadataURIUpdated(string newMetadataURI)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageFilterer) ParseMetadataURIUpdated(log types.Log) (*DurationVaultStrategyStorageMetadataURIUpdated, error) {
	event := new(DurationVaultStrategyStorageMetadataURIUpdated)
	if err := _DurationVaultStrategyStorage.contract.UnpackLog(event, "MetadataURIUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DurationVaultStrategyStorageStrategyTokenSetIterator is returned from FilterStrategyTokenSet and is used to iterate over the raw logs and unpacked data for StrategyTokenSet events raised by the DurationVaultStrategyStorage contract.
type DurationVaultStrategyStorageStrategyTokenSetIterator struct {
	Event *DurationVaultStrategyStorageStrategyTokenSet // Event containing the contract specifics and raw log

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
func (it *DurationVaultStrategyStorageStrategyTokenSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DurationVaultStrategyStorageStrategyTokenSet)
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
		it.Event = new(DurationVaultStrategyStorageStrategyTokenSet)
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
func (it *DurationVaultStrategyStorageStrategyTokenSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DurationVaultStrategyStorageStrategyTokenSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DurationVaultStrategyStorageStrategyTokenSet represents a StrategyTokenSet event raised by the DurationVaultStrategyStorage contract.
type DurationVaultStrategyStorageStrategyTokenSet struct {
	Token    common.Address
	Decimals uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStrategyTokenSet is a free log retrieval operation binding the contract event 0x1c540707b00eb5427b6b774fc799d756516a54aee108b64b327acc55af557507.
//
// Solidity: event StrategyTokenSet(address token, uint8 decimals)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageFilterer) FilterStrategyTokenSet(opts *bind.FilterOpts) (*DurationVaultStrategyStorageStrategyTokenSetIterator, error) {

	logs, sub, err := _DurationVaultStrategyStorage.contract.FilterLogs(opts, "StrategyTokenSet")
	if err != nil {
		return nil, err
	}
	return &DurationVaultStrategyStorageStrategyTokenSetIterator{contract: _DurationVaultStrategyStorage.contract, event: "StrategyTokenSet", logs: logs, sub: sub}, nil
}

// WatchStrategyTokenSet is a free log subscription operation binding the contract event 0x1c540707b00eb5427b6b774fc799d756516a54aee108b64b327acc55af557507.
//
// Solidity: event StrategyTokenSet(address token, uint8 decimals)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageFilterer) WatchStrategyTokenSet(opts *bind.WatchOpts, sink chan<- *DurationVaultStrategyStorageStrategyTokenSet) (event.Subscription, error) {

	logs, sub, err := _DurationVaultStrategyStorage.contract.WatchLogs(opts, "StrategyTokenSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DurationVaultStrategyStorageStrategyTokenSet)
				if err := _DurationVaultStrategyStorage.contract.UnpackLog(event, "StrategyTokenSet", log); err != nil {
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

// ParseStrategyTokenSet is a log parse operation binding the contract event 0x1c540707b00eb5427b6b774fc799d756516a54aee108b64b327acc55af557507.
//
// Solidity: event StrategyTokenSet(address token, uint8 decimals)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageFilterer) ParseStrategyTokenSet(log types.Log) (*DurationVaultStrategyStorageStrategyTokenSet, error) {
	event := new(DurationVaultStrategyStorageStrategyTokenSet)
	if err := _DurationVaultStrategyStorage.contract.UnpackLog(event, "StrategyTokenSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DurationVaultStrategyStorageVaultAdvancedToWithdrawalsIterator is returned from FilterVaultAdvancedToWithdrawals and is used to iterate over the raw logs and unpacked data for VaultAdvancedToWithdrawals events raised by the DurationVaultStrategyStorage contract.
type DurationVaultStrategyStorageVaultAdvancedToWithdrawalsIterator struct {
	Event *DurationVaultStrategyStorageVaultAdvancedToWithdrawals // Event containing the contract specifics and raw log

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
func (it *DurationVaultStrategyStorageVaultAdvancedToWithdrawalsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DurationVaultStrategyStorageVaultAdvancedToWithdrawals)
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
		it.Event = new(DurationVaultStrategyStorageVaultAdvancedToWithdrawals)
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
func (it *DurationVaultStrategyStorageVaultAdvancedToWithdrawalsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DurationVaultStrategyStorageVaultAdvancedToWithdrawalsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DurationVaultStrategyStorageVaultAdvancedToWithdrawals represents a VaultAdvancedToWithdrawals event raised by the DurationVaultStrategyStorage contract.
type DurationVaultStrategyStorageVaultAdvancedToWithdrawals struct {
	Arbitrator common.Address
	MaturedAt  uint32
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVaultAdvancedToWithdrawals is a free log retrieval operation binding the contract event 0x96c49d03ef64591194500229a104cd087b2d45c68234c96444c3a2a6abb0bb97.
//
// Solidity: event VaultAdvancedToWithdrawals(address indexed arbitrator, uint32 maturedAt)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageFilterer) FilterVaultAdvancedToWithdrawals(opts *bind.FilterOpts, arbitrator []common.Address) (*DurationVaultStrategyStorageVaultAdvancedToWithdrawalsIterator, error) {

	var arbitratorRule []interface{}
	for _, arbitratorItem := range arbitrator {
		arbitratorRule = append(arbitratorRule, arbitratorItem)
	}

	logs, sub, err := _DurationVaultStrategyStorage.contract.FilterLogs(opts, "VaultAdvancedToWithdrawals", arbitratorRule)
	if err != nil {
		return nil, err
	}
	return &DurationVaultStrategyStorageVaultAdvancedToWithdrawalsIterator{contract: _DurationVaultStrategyStorage.contract, event: "VaultAdvancedToWithdrawals", logs: logs, sub: sub}, nil
}

// WatchVaultAdvancedToWithdrawals is a free log subscription operation binding the contract event 0x96c49d03ef64591194500229a104cd087b2d45c68234c96444c3a2a6abb0bb97.
//
// Solidity: event VaultAdvancedToWithdrawals(address indexed arbitrator, uint32 maturedAt)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageFilterer) WatchVaultAdvancedToWithdrawals(opts *bind.WatchOpts, sink chan<- *DurationVaultStrategyStorageVaultAdvancedToWithdrawals, arbitrator []common.Address) (event.Subscription, error) {

	var arbitratorRule []interface{}
	for _, arbitratorItem := range arbitrator {
		arbitratorRule = append(arbitratorRule, arbitratorItem)
	}

	logs, sub, err := _DurationVaultStrategyStorage.contract.WatchLogs(opts, "VaultAdvancedToWithdrawals", arbitratorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DurationVaultStrategyStorageVaultAdvancedToWithdrawals)
				if err := _DurationVaultStrategyStorage.contract.UnpackLog(event, "VaultAdvancedToWithdrawals", log); err != nil {
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

// ParseVaultAdvancedToWithdrawals is a log parse operation binding the contract event 0x96c49d03ef64591194500229a104cd087b2d45c68234c96444c3a2a6abb0bb97.
//
// Solidity: event VaultAdvancedToWithdrawals(address indexed arbitrator, uint32 maturedAt)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageFilterer) ParseVaultAdvancedToWithdrawals(log types.Log) (*DurationVaultStrategyStorageVaultAdvancedToWithdrawals, error) {
	event := new(DurationVaultStrategyStorageVaultAdvancedToWithdrawals)
	if err := _DurationVaultStrategyStorage.contract.UnpackLog(event, "VaultAdvancedToWithdrawals", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DurationVaultStrategyStorageVaultInitializedIterator is returned from FilterVaultInitialized and is used to iterate over the raw logs and unpacked data for VaultInitialized events raised by the DurationVaultStrategyStorage contract.
type DurationVaultStrategyStorageVaultInitializedIterator struct {
	Event *DurationVaultStrategyStorageVaultInitialized // Event containing the contract specifics and raw log

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
func (it *DurationVaultStrategyStorageVaultInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DurationVaultStrategyStorageVaultInitialized)
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
		it.Event = new(DurationVaultStrategyStorageVaultInitialized)
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
func (it *DurationVaultStrategyStorageVaultInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DurationVaultStrategyStorageVaultInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DurationVaultStrategyStorageVaultInitialized represents a VaultInitialized event raised by the DurationVaultStrategyStorage contract.
type DurationVaultStrategyStorageVaultInitialized struct {
	VaultAdmin      common.Address
	Arbitrator      common.Address
	UnderlyingToken common.Address
	Duration        uint32
	MaxPerDeposit   *big.Int
	StakeCap        *big.Int
	MetadataURI     string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterVaultInitialized is a free log retrieval operation binding the contract event 0xbdbff63632f473bb2a7c6a4aafbc096b71fbda12e22c6b51643bfd64f13d2b9e.
//
// Solidity: event VaultInitialized(address indexed vaultAdmin, address indexed arbitrator, address indexed underlyingToken, uint32 duration, uint256 maxPerDeposit, uint256 stakeCap, string metadataURI)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageFilterer) FilterVaultInitialized(opts *bind.FilterOpts, vaultAdmin []common.Address, arbitrator []common.Address, underlyingToken []common.Address) (*DurationVaultStrategyStorageVaultInitializedIterator, error) {

	var vaultAdminRule []interface{}
	for _, vaultAdminItem := range vaultAdmin {
		vaultAdminRule = append(vaultAdminRule, vaultAdminItem)
	}
	var arbitratorRule []interface{}
	for _, arbitratorItem := range arbitrator {
		arbitratorRule = append(arbitratorRule, arbitratorItem)
	}
	var underlyingTokenRule []interface{}
	for _, underlyingTokenItem := range underlyingToken {
		underlyingTokenRule = append(underlyingTokenRule, underlyingTokenItem)
	}

	logs, sub, err := _DurationVaultStrategyStorage.contract.FilterLogs(opts, "VaultInitialized", vaultAdminRule, arbitratorRule, underlyingTokenRule)
	if err != nil {
		return nil, err
	}
	return &DurationVaultStrategyStorageVaultInitializedIterator{contract: _DurationVaultStrategyStorage.contract, event: "VaultInitialized", logs: logs, sub: sub}, nil
}

// WatchVaultInitialized is a free log subscription operation binding the contract event 0xbdbff63632f473bb2a7c6a4aafbc096b71fbda12e22c6b51643bfd64f13d2b9e.
//
// Solidity: event VaultInitialized(address indexed vaultAdmin, address indexed arbitrator, address indexed underlyingToken, uint32 duration, uint256 maxPerDeposit, uint256 stakeCap, string metadataURI)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageFilterer) WatchVaultInitialized(opts *bind.WatchOpts, sink chan<- *DurationVaultStrategyStorageVaultInitialized, vaultAdmin []common.Address, arbitrator []common.Address, underlyingToken []common.Address) (event.Subscription, error) {

	var vaultAdminRule []interface{}
	for _, vaultAdminItem := range vaultAdmin {
		vaultAdminRule = append(vaultAdminRule, vaultAdminItem)
	}
	var arbitratorRule []interface{}
	for _, arbitratorItem := range arbitrator {
		arbitratorRule = append(arbitratorRule, arbitratorItem)
	}
	var underlyingTokenRule []interface{}
	for _, underlyingTokenItem := range underlyingToken {
		underlyingTokenRule = append(underlyingTokenRule, underlyingTokenItem)
	}

	logs, sub, err := _DurationVaultStrategyStorage.contract.WatchLogs(opts, "VaultInitialized", vaultAdminRule, arbitratorRule, underlyingTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DurationVaultStrategyStorageVaultInitialized)
				if err := _DurationVaultStrategyStorage.contract.UnpackLog(event, "VaultInitialized", log); err != nil {
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

// ParseVaultInitialized is a log parse operation binding the contract event 0xbdbff63632f473bb2a7c6a4aafbc096b71fbda12e22c6b51643bfd64f13d2b9e.
//
// Solidity: event VaultInitialized(address indexed vaultAdmin, address indexed arbitrator, address indexed underlyingToken, uint32 duration, uint256 maxPerDeposit, uint256 stakeCap, string metadataURI)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageFilterer) ParseVaultInitialized(log types.Log) (*DurationVaultStrategyStorageVaultInitialized, error) {
	event := new(DurationVaultStrategyStorageVaultInitialized)
	if err := _DurationVaultStrategyStorage.contract.UnpackLog(event, "VaultInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DurationVaultStrategyStorageVaultLockedIterator is returned from FilterVaultLocked and is used to iterate over the raw logs and unpacked data for VaultLocked events raised by the DurationVaultStrategyStorage contract.
type DurationVaultStrategyStorageVaultLockedIterator struct {
	Event *DurationVaultStrategyStorageVaultLocked // Event containing the contract specifics and raw log

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
func (it *DurationVaultStrategyStorageVaultLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DurationVaultStrategyStorageVaultLocked)
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
		it.Event = new(DurationVaultStrategyStorageVaultLocked)
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
func (it *DurationVaultStrategyStorageVaultLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DurationVaultStrategyStorageVaultLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DurationVaultStrategyStorageVaultLocked represents a VaultLocked event raised by the DurationVaultStrategyStorage contract.
type DurationVaultStrategyStorageVaultLocked struct {
	LockedAt uint32
	UnlockAt uint32
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterVaultLocked is a free log retrieval operation binding the contract event 0x42cd6d7338516695d9c9ff8969dbdcf89ce22e3f2f76fda2fc11e973fe4860e4.
//
// Solidity: event VaultLocked(uint32 lockedAt, uint32 unlockAt)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageFilterer) FilterVaultLocked(opts *bind.FilterOpts) (*DurationVaultStrategyStorageVaultLockedIterator, error) {

	logs, sub, err := _DurationVaultStrategyStorage.contract.FilterLogs(opts, "VaultLocked")
	if err != nil {
		return nil, err
	}
	return &DurationVaultStrategyStorageVaultLockedIterator{contract: _DurationVaultStrategyStorage.contract, event: "VaultLocked", logs: logs, sub: sub}, nil
}

// WatchVaultLocked is a free log subscription operation binding the contract event 0x42cd6d7338516695d9c9ff8969dbdcf89ce22e3f2f76fda2fc11e973fe4860e4.
//
// Solidity: event VaultLocked(uint32 lockedAt, uint32 unlockAt)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageFilterer) WatchVaultLocked(opts *bind.WatchOpts, sink chan<- *DurationVaultStrategyStorageVaultLocked) (event.Subscription, error) {

	logs, sub, err := _DurationVaultStrategyStorage.contract.WatchLogs(opts, "VaultLocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DurationVaultStrategyStorageVaultLocked)
				if err := _DurationVaultStrategyStorage.contract.UnpackLog(event, "VaultLocked", log); err != nil {
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

// ParseVaultLocked is a log parse operation binding the contract event 0x42cd6d7338516695d9c9ff8969dbdcf89ce22e3f2f76fda2fc11e973fe4860e4.
//
// Solidity: event VaultLocked(uint32 lockedAt, uint32 unlockAt)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageFilterer) ParseVaultLocked(log types.Log) (*DurationVaultStrategyStorageVaultLocked, error) {
	event := new(DurationVaultStrategyStorageVaultLocked)
	if err := _DurationVaultStrategyStorage.contract.UnpackLog(event, "VaultLocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DurationVaultStrategyStorageVaultMaturedIterator is returned from FilterVaultMatured and is used to iterate over the raw logs and unpacked data for VaultMatured events raised by the DurationVaultStrategyStorage contract.
type DurationVaultStrategyStorageVaultMaturedIterator struct {
	Event *DurationVaultStrategyStorageVaultMatured // Event containing the contract specifics and raw log

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
func (it *DurationVaultStrategyStorageVaultMaturedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DurationVaultStrategyStorageVaultMatured)
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
		it.Event = new(DurationVaultStrategyStorageVaultMatured)
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
func (it *DurationVaultStrategyStorageVaultMaturedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DurationVaultStrategyStorageVaultMaturedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DurationVaultStrategyStorageVaultMatured represents a VaultMatured event raised by the DurationVaultStrategyStorage contract.
type DurationVaultStrategyStorageVaultMatured struct {
	MaturedAt uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterVaultMatured is a free log retrieval operation binding the contract event 0xff979382d3040b1602e0a02f0f2a454b2250aa36e891d2da0ceb95d70d11a8f2.
//
// Solidity: event VaultMatured(uint32 maturedAt)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageFilterer) FilterVaultMatured(opts *bind.FilterOpts) (*DurationVaultStrategyStorageVaultMaturedIterator, error) {

	logs, sub, err := _DurationVaultStrategyStorage.contract.FilterLogs(opts, "VaultMatured")
	if err != nil {
		return nil, err
	}
	return &DurationVaultStrategyStorageVaultMaturedIterator{contract: _DurationVaultStrategyStorage.contract, event: "VaultMatured", logs: logs, sub: sub}, nil
}

// WatchVaultMatured is a free log subscription operation binding the contract event 0xff979382d3040b1602e0a02f0f2a454b2250aa36e891d2da0ceb95d70d11a8f2.
//
// Solidity: event VaultMatured(uint32 maturedAt)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageFilterer) WatchVaultMatured(opts *bind.WatchOpts, sink chan<- *DurationVaultStrategyStorageVaultMatured) (event.Subscription, error) {

	logs, sub, err := _DurationVaultStrategyStorage.contract.WatchLogs(opts, "VaultMatured")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DurationVaultStrategyStorageVaultMatured)
				if err := _DurationVaultStrategyStorage.contract.UnpackLog(event, "VaultMatured", log); err != nil {
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

// ParseVaultMatured is a log parse operation binding the contract event 0xff979382d3040b1602e0a02f0f2a454b2250aa36e891d2da0ceb95d70d11a8f2.
//
// Solidity: event VaultMatured(uint32 maturedAt)
func (_DurationVaultStrategyStorage *DurationVaultStrategyStorageFilterer) ParseVaultMatured(log types.Log) (*DurationVaultStrategyStorageVaultMatured, error) {
	event := new(DurationVaultStrategyStorageVaultMatured)
	if err := _DurationVaultStrategyStorage.contract.UnpackLog(event, "VaultMatured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
