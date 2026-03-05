// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IDurationVaultStrategy

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

// IDurationVaultStrategyMetaData contains all meta data concerning the IDurationVaultStrategy contract.
var IDurationVaultStrategyMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"advanceToWithdrawals\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allocationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allocationsActive\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"arbitrator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"beforeAddShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"beforeRemoveShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositsOpen\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"duration\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"explanation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isLocked\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isMatured\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lock\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockedAt\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"markMatured\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"maxPerDeposit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxTotalDeposits\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"metadataURI\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorIntegrationConfigured\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorSetInfo\",\"inputs\":[],\"outputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorSetRegistered\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"rewardsCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRewardsCoordinator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setRewardsClaimer\",\"inputs\":[{\"name\":\"claimer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"shares\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sharesToUnderlying\",\"inputs\":[{\"name\":\"amountShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sharesToUnderlyingView\",\"inputs\":[{\"name\":\"amountShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakeCap\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"state\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIDurationVaultStrategyTypes.VaultState\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalShares\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"underlyingToShares\",\"inputs\":[{\"name\":\"amountUnderlying\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"underlyingToSharesView\",\"inputs\":[{\"name\":\"amountUnderlying\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"underlyingToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unlockTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateDelegationApprover\",\"inputs\":[{\"name\":\"newDelegationApprover\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateMetadataURI\",\"inputs\":[{\"name\":\"newMetadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOperatorMetadataURI\",\"inputs\":[{\"name\":\"newOperatorMetadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateTVLLimits\",\"inputs\":[{\"name\":\"newMaxPerDeposit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"newStakeCap\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"userUnderlying\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"userUnderlyingView\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"vaultAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amountShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawalsOpen\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"DeallocateAttempted\",\"inputs\":[{\"name\":\"success\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DeregisterAttempted\",\"inputs\":[{\"name\":\"success\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExchangeRateEmitted\",\"inputs\":[{\"name\":\"rate\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MaxPerDepositUpdated\",\"inputs\":[{\"name\":\"previousValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MaxTotalDepositsUpdated\",\"inputs\":[{\"name\":\"previousValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MetadataURIUpdated\",\"inputs\":[{\"name\":\"newMetadataURI\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyTokenSet\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIERC20\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VaultAdvancedToWithdrawals\",\"inputs\":[{\"name\":\"arbitrator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"maturedAt\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VaultInitialized\",\"inputs\":[{\"name\":\"vaultAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"arbitrator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"underlyingToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"contractIERC20\"},{\"name\":\"duration\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"maxPerDeposit\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"stakeCap\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"metadataURI\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VaultLocked\",\"inputs\":[{\"name\":\"lockedAt\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"unlockAt\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VaultMatured\",\"inputs\":[{\"name\":\"maturedAt\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BalanceExceedsMaxTotalDeposits\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DepositExceedsMaxPerDeposit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DepositsLocked\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DurationAlreadyElapsed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DurationNotElapsed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidArbitrator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDuration\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidVaultAdmin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MaxPerDepositExceedsMax\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeDelegatedToVaultOperator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NewSharesZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyArbitrator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyStrategyManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnderlyingToken\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyVaultAdmin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorIntegrationInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PendingAllocation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategyNotSupportedByOperatorSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TotalSharesExceedsMax\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnderlyingTokenBlacklisted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"VaultAlreadyLocked\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"VaultNotLocked\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawalAmountExceedsTotalDeposits\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawalsLockedDuringAllocations\",\"inputs\":[]}]",
}

// IDurationVaultStrategyABI is the input ABI used to generate the binding from.
// Deprecated: Use IDurationVaultStrategyMetaData.ABI instead.
var IDurationVaultStrategyABI = IDurationVaultStrategyMetaData.ABI

// IDurationVaultStrategy is an auto generated Go binding around an Ethereum contract.
type IDurationVaultStrategy struct {
	IDurationVaultStrategyCaller     // Read-only binding to the contract
	IDurationVaultStrategyTransactor // Write-only binding to the contract
	IDurationVaultStrategyFilterer   // Log filterer for contract events
}

// IDurationVaultStrategyCaller is an auto generated read-only Go binding around an Ethereum contract.
type IDurationVaultStrategyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDurationVaultStrategyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IDurationVaultStrategyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDurationVaultStrategyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IDurationVaultStrategyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDurationVaultStrategySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IDurationVaultStrategySession struct {
	Contract     *IDurationVaultStrategy // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IDurationVaultStrategyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IDurationVaultStrategyCallerSession struct {
	Contract *IDurationVaultStrategyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// IDurationVaultStrategyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IDurationVaultStrategyTransactorSession struct {
	Contract     *IDurationVaultStrategyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// IDurationVaultStrategyRaw is an auto generated low-level Go binding around an Ethereum contract.
type IDurationVaultStrategyRaw struct {
	Contract *IDurationVaultStrategy // Generic contract binding to access the raw methods on
}

// IDurationVaultStrategyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IDurationVaultStrategyCallerRaw struct {
	Contract *IDurationVaultStrategyCaller // Generic read-only contract binding to access the raw methods on
}

// IDurationVaultStrategyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IDurationVaultStrategyTransactorRaw struct {
	Contract *IDurationVaultStrategyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIDurationVaultStrategy creates a new instance of IDurationVaultStrategy, bound to a specific deployed contract.
func NewIDurationVaultStrategy(address common.Address, backend bind.ContractBackend) (*IDurationVaultStrategy, error) {
	contract, err := bindIDurationVaultStrategy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IDurationVaultStrategy{IDurationVaultStrategyCaller: IDurationVaultStrategyCaller{contract: contract}, IDurationVaultStrategyTransactor: IDurationVaultStrategyTransactor{contract: contract}, IDurationVaultStrategyFilterer: IDurationVaultStrategyFilterer{contract: contract}}, nil
}

// NewIDurationVaultStrategyCaller creates a new read-only instance of IDurationVaultStrategy, bound to a specific deployed contract.
func NewIDurationVaultStrategyCaller(address common.Address, caller bind.ContractCaller) (*IDurationVaultStrategyCaller, error) {
	contract, err := bindIDurationVaultStrategy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IDurationVaultStrategyCaller{contract: contract}, nil
}

// NewIDurationVaultStrategyTransactor creates a new write-only instance of IDurationVaultStrategy, bound to a specific deployed contract.
func NewIDurationVaultStrategyTransactor(address common.Address, transactor bind.ContractTransactor) (*IDurationVaultStrategyTransactor, error) {
	contract, err := bindIDurationVaultStrategy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IDurationVaultStrategyTransactor{contract: contract}, nil
}

// NewIDurationVaultStrategyFilterer creates a new log filterer instance of IDurationVaultStrategy, bound to a specific deployed contract.
func NewIDurationVaultStrategyFilterer(address common.Address, filterer bind.ContractFilterer) (*IDurationVaultStrategyFilterer, error) {
	contract, err := bindIDurationVaultStrategy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IDurationVaultStrategyFilterer{contract: contract}, nil
}

// bindIDurationVaultStrategy binds a generic wrapper to an already deployed contract.
func bindIDurationVaultStrategy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IDurationVaultStrategyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDurationVaultStrategy *IDurationVaultStrategyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDurationVaultStrategy.Contract.IDurationVaultStrategyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDurationVaultStrategy *IDurationVaultStrategyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDurationVaultStrategy.Contract.IDurationVaultStrategyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDurationVaultStrategy *IDurationVaultStrategyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDurationVaultStrategy.Contract.IDurationVaultStrategyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDurationVaultStrategy *IDurationVaultStrategyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDurationVaultStrategy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDurationVaultStrategy *IDurationVaultStrategyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDurationVaultStrategy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDurationVaultStrategy *IDurationVaultStrategyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDurationVaultStrategy.Contract.contract.Transact(opts, method, params...)
}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_IDurationVaultStrategy *IDurationVaultStrategyCaller) AllocationManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IDurationVaultStrategy.contract.Call(opts, &out, "allocationManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_IDurationVaultStrategy *IDurationVaultStrategySession) AllocationManager() (common.Address, error) {
	return _IDurationVaultStrategy.Contract.AllocationManager(&_IDurationVaultStrategy.CallOpts)
}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_IDurationVaultStrategy *IDurationVaultStrategyCallerSession) AllocationManager() (common.Address, error) {
	return _IDurationVaultStrategy.Contract.AllocationManager(&_IDurationVaultStrategy.CallOpts)
}

// AllocationsActive is a free data retrieval call binding the contract method 0xfb4d86b4.
//
// Solidity: function allocationsActive() view returns(bool)
func (_IDurationVaultStrategy *IDurationVaultStrategyCaller) AllocationsActive(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IDurationVaultStrategy.contract.Call(opts, &out, "allocationsActive")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllocationsActive is a free data retrieval call binding the contract method 0xfb4d86b4.
//
// Solidity: function allocationsActive() view returns(bool)
func (_IDurationVaultStrategy *IDurationVaultStrategySession) AllocationsActive() (bool, error) {
	return _IDurationVaultStrategy.Contract.AllocationsActive(&_IDurationVaultStrategy.CallOpts)
}

// AllocationsActive is a free data retrieval call binding the contract method 0xfb4d86b4.
//
// Solidity: function allocationsActive() view returns(bool)
func (_IDurationVaultStrategy *IDurationVaultStrategyCallerSession) AllocationsActive() (bool, error) {
	return _IDurationVaultStrategy.Contract.AllocationsActive(&_IDurationVaultStrategy.CallOpts)
}

// Arbitrator is a free data retrieval call binding the contract method 0x6cc6cde1.
//
// Solidity: function arbitrator() view returns(address)
func (_IDurationVaultStrategy *IDurationVaultStrategyCaller) Arbitrator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IDurationVaultStrategy.contract.Call(opts, &out, "arbitrator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Arbitrator is a free data retrieval call binding the contract method 0x6cc6cde1.
//
// Solidity: function arbitrator() view returns(address)
func (_IDurationVaultStrategy *IDurationVaultStrategySession) Arbitrator() (common.Address, error) {
	return _IDurationVaultStrategy.Contract.Arbitrator(&_IDurationVaultStrategy.CallOpts)
}

// Arbitrator is a free data retrieval call binding the contract method 0x6cc6cde1.
//
// Solidity: function arbitrator() view returns(address)
func (_IDurationVaultStrategy *IDurationVaultStrategyCallerSession) Arbitrator() (common.Address, error) {
	return _IDurationVaultStrategy.Contract.Arbitrator(&_IDurationVaultStrategy.CallOpts)
}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_IDurationVaultStrategy *IDurationVaultStrategyCaller) DelegationManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IDurationVaultStrategy.contract.Call(opts, &out, "delegationManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_IDurationVaultStrategy *IDurationVaultStrategySession) DelegationManager() (common.Address, error) {
	return _IDurationVaultStrategy.Contract.DelegationManager(&_IDurationVaultStrategy.CallOpts)
}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_IDurationVaultStrategy *IDurationVaultStrategyCallerSession) DelegationManager() (common.Address, error) {
	return _IDurationVaultStrategy.Contract.DelegationManager(&_IDurationVaultStrategy.CallOpts)
}

// DepositsOpen is a free data retrieval call binding the contract method 0x549c4627.
//
// Solidity: function depositsOpen() view returns(bool)
func (_IDurationVaultStrategy *IDurationVaultStrategyCaller) DepositsOpen(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IDurationVaultStrategy.contract.Call(opts, &out, "depositsOpen")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DepositsOpen is a free data retrieval call binding the contract method 0x549c4627.
//
// Solidity: function depositsOpen() view returns(bool)
func (_IDurationVaultStrategy *IDurationVaultStrategySession) DepositsOpen() (bool, error) {
	return _IDurationVaultStrategy.Contract.DepositsOpen(&_IDurationVaultStrategy.CallOpts)
}

// DepositsOpen is a free data retrieval call binding the contract method 0x549c4627.
//
// Solidity: function depositsOpen() view returns(bool)
func (_IDurationVaultStrategy *IDurationVaultStrategyCallerSession) DepositsOpen() (bool, error) {
	return _IDurationVaultStrategy.Contract.DepositsOpen(&_IDurationVaultStrategy.CallOpts)
}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint32)
func (_IDurationVaultStrategy *IDurationVaultStrategyCaller) Duration(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _IDurationVaultStrategy.contract.Call(opts, &out, "duration")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint32)
func (_IDurationVaultStrategy *IDurationVaultStrategySession) Duration() (uint32, error) {
	return _IDurationVaultStrategy.Contract.Duration(&_IDurationVaultStrategy.CallOpts)
}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint32)
func (_IDurationVaultStrategy *IDurationVaultStrategyCallerSession) Duration() (uint32, error) {
	return _IDurationVaultStrategy.Contract.Duration(&_IDurationVaultStrategy.CallOpts)
}

// Explanation is a free data retrieval call binding the contract method 0xab5921e1.
//
// Solidity: function explanation() view returns(string)
func (_IDurationVaultStrategy *IDurationVaultStrategyCaller) Explanation(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IDurationVaultStrategy.contract.Call(opts, &out, "explanation")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Explanation is a free data retrieval call binding the contract method 0xab5921e1.
//
// Solidity: function explanation() view returns(string)
func (_IDurationVaultStrategy *IDurationVaultStrategySession) Explanation() (string, error) {
	return _IDurationVaultStrategy.Contract.Explanation(&_IDurationVaultStrategy.CallOpts)
}

// Explanation is a free data retrieval call binding the contract method 0xab5921e1.
//
// Solidity: function explanation() view returns(string)
func (_IDurationVaultStrategy *IDurationVaultStrategyCallerSession) Explanation() (string, error) {
	return _IDurationVaultStrategy.Contract.Explanation(&_IDurationVaultStrategy.CallOpts)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() view returns(bool)
func (_IDurationVaultStrategy *IDurationVaultStrategyCaller) IsLocked(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IDurationVaultStrategy.contract.Call(opts, &out, "isLocked")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() view returns(bool)
func (_IDurationVaultStrategy *IDurationVaultStrategySession) IsLocked() (bool, error) {
	return _IDurationVaultStrategy.Contract.IsLocked(&_IDurationVaultStrategy.CallOpts)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() view returns(bool)
func (_IDurationVaultStrategy *IDurationVaultStrategyCallerSession) IsLocked() (bool, error) {
	return _IDurationVaultStrategy.Contract.IsLocked(&_IDurationVaultStrategy.CallOpts)
}

// IsMatured is a free data retrieval call binding the contract method 0x7f2b6a0d.
//
// Solidity: function isMatured() view returns(bool)
func (_IDurationVaultStrategy *IDurationVaultStrategyCaller) IsMatured(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IDurationVaultStrategy.contract.Call(opts, &out, "isMatured")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMatured is a free data retrieval call binding the contract method 0x7f2b6a0d.
//
// Solidity: function isMatured() view returns(bool)
func (_IDurationVaultStrategy *IDurationVaultStrategySession) IsMatured() (bool, error) {
	return _IDurationVaultStrategy.Contract.IsMatured(&_IDurationVaultStrategy.CallOpts)
}

// IsMatured is a free data retrieval call binding the contract method 0x7f2b6a0d.
//
// Solidity: function isMatured() view returns(bool)
func (_IDurationVaultStrategy *IDurationVaultStrategyCallerSession) IsMatured() (bool, error) {
	return _IDurationVaultStrategy.Contract.IsMatured(&_IDurationVaultStrategy.CallOpts)
}

// LockedAt is a free data retrieval call binding the contract method 0xb2163482.
//
// Solidity: function lockedAt() view returns(uint32)
func (_IDurationVaultStrategy *IDurationVaultStrategyCaller) LockedAt(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _IDurationVaultStrategy.contract.Call(opts, &out, "lockedAt")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LockedAt is a free data retrieval call binding the contract method 0xb2163482.
//
// Solidity: function lockedAt() view returns(uint32)
func (_IDurationVaultStrategy *IDurationVaultStrategySession) LockedAt() (uint32, error) {
	return _IDurationVaultStrategy.Contract.LockedAt(&_IDurationVaultStrategy.CallOpts)
}

// LockedAt is a free data retrieval call binding the contract method 0xb2163482.
//
// Solidity: function lockedAt() view returns(uint32)
func (_IDurationVaultStrategy *IDurationVaultStrategyCallerSession) LockedAt() (uint32, error) {
	return _IDurationVaultStrategy.Contract.LockedAt(&_IDurationVaultStrategy.CallOpts)
}

// MaxPerDeposit is a free data retrieval call binding the contract method 0x43fe08b0.
//
// Solidity: function maxPerDeposit() view returns(uint256)
func (_IDurationVaultStrategy *IDurationVaultStrategyCaller) MaxPerDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IDurationVaultStrategy.contract.Call(opts, &out, "maxPerDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxPerDeposit is a free data retrieval call binding the contract method 0x43fe08b0.
//
// Solidity: function maxPerDeposit() view returns(uint256)
func (_IDurationVaultStrategy *IDurationVaultStrategySession) MaxPerDeposit() (*big.Int, error) {
	return _IDurationVaultStrategy.Contract.MaxPerDeposit(&_IDurationVaultStrategy.CallOpts)
}

// MaxPerDeposit is a free data retrieval call binding the contract method 0x43fe08b0.
//
// Solidity: function maxPerDeposit() view returns(uint256)
func (_IDurationVaultStrategy *IDurationVaultStrategyCallerSession) MaxPerDeposit() (*big.Int, error) {
	return _IDurationVaultStrategy.Contract.MaxPerDeposit(&_IDurationVaultStrategy.CallOpts)
}

// MaxTotalDeposits is a free data retrieval call binding the contract method 0x61b01b5d.
//
// Solidity: function maxTotalDeposits() view returns(uint256)
func (_IDurationVaultStrategy *IDurationVaultStrategyCaller) MaxTotalDeposits(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IDurationVaultStrategy.contract.Call(opts, &out, "maxTotalDeposits")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxTotalDeposits is a free data retrieval call binding the contract method 0x61b01b5d.
//
// Solidity: function maxTotalDeposits() view returns(uint256)
func (_IDurationVaultStrategy *IDurationVaultStrategySession) MaxTotalDeposits() (*big.Int, error) {
	return _IDurationVaultStrategy.Contract.MaxTotalDeposits(&_IDurationVaultStrategy.CallOpts)
}

// MaxTotalDeposits is a free data retrieval call binding the contract method 0x61b01b5d.
//
// Solidity: function maxTotalDeposits() view returns(uint256)
func (_IDurationVaultStrategy *IDurationVaultStrategyCallerSession) MaxTotalDeposits() (*big.Int, error) {
	return _IDurationVaultStrategy.Contract.MaxTotalDeposits(&_IDurationVaultStrategy.CallOpts)
}

// MetadataURI is a free data retrieval call binding the contract method 0x03ee438c.
//
// Solidity: function metadataURI() view returns(string)
func (_IDurationVaultStrategy *IDurationVaultStrategyCaller) MetadataURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IDurationVaultStrategy.contract.Call(opts, &out, "metadataURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// MetadataURI is a free data retrieval call binding the contract method 0x03ee438c.
//
// Solidity: function metadataURI() view returns(string)
func (_IDurationVaultStrategy *IDurationVaultStrategySession) MetadataURI() (string, error) {
	return _IDurationVaultStrategy.Contract.MetadataURI(&_IDurationVaultStrategy.CallOpts)
}

// MetadataURI is a free data retrieval call binding the contract method 0x03ee438c.
//
// Solidity: function metadataURI() view returns(string)
func (_IDurationVaultStrategy *IDurationVaultStrategyCallerSession) MetadataURI() (string, error) {
	return _IDurationVaultStrategy.Contract.MetadataURI(&_IDurationVaultStrategy.CallOpts)
}

// OperatorIntegrationConfigured is a free data retrieval call binding the contract method 0x5438a8c7.
//
// Solidity: function operatorIntegrationConfigured() view returns(bool)
func (_IDurationVaultStrategy *IDurationVaultStrategyCaller) OperatorIntegrationConfigured(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IDurationVaultStrategy.contract.Call(opts, &out, "operatorIntegrationConfigured")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OperatorIntegrationConfigured is a free data retrieval call binding the contract method 0x5438a8c7.
//
// Solidity: function operatorIntegrationConfigured() view returns(bool)
func (_IDurationVaultStrategy *IDurationVaultStrategySession) OperatorIntegrationConfigured() (bool, error) {
	return _IDurationVaultStrategy.Contract.OperatorIntegrationConfigured(&_IDurationVaultStrategy.CallOpts)
}

// OperatorIntegrationConfigured is a free data retrieval call binding the contract method 0x5438a8c7.
//
// Solidity: function operatorIntegrationConfigured() view returns(bool)
func (_IDurationVaultStrategy *IDurationVaultStrategyCallerSession) OperatorIntegrationConfigured() (bool, error) {
	return _IDurationVaultStrategy.Contract.OperatorIntegrationConfigured(&_IDurationVaultStrategy.CallOpts)
}

// OperatorSetInfo is a free data retrieval call binding the contract method 0xd4deae81.
//
// Solidity: function operatorSetInfo() view returns(address avs, uint32 operatorSetId)
func (_IDurationVaultStrategy *IDurationVaultStrategyCaller) OperatorSetInfo(opts *bind.CallOpts) (struct {
	Avs           common.Address
	OperatorSetId uint32
}, error) {
	var out []interface{}
	err := _IDurationVaultStrategy.contract.Call(opts, &out, "operatorSetInfo")

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
func (_IDurationVaultStrategy *IDurationVaultStrategySession) OperatorSetInfo() (struct {
	Avs           common.Address
	OperatorSetId uint32
}, error) {
	return _IDurationVaultStrategy.Contract.OperatorSetInfo(&_IDurationVaultStrategy.CallOpts)
}

// OperatorSetInfo is a free data retrieval call binding the contract method 0xd4deae81.
//
// Solidity: function operatorSetInfo() view returns(address avs, uint32 operatorSetId)
func (_IDurationVaultStrategy *IDurationVaultStrategyCallerSession) OperatorSetInfo() (struct {
	Avs           common.Address
	OperatorSetId uint32
}, error) {
	return _IDurationVaultStrategy.Contract.OperatorSetInfo(&_IDurationVaultStrategy.CallOpts)
}

// OperatorSetRegistered is a free data retrieval call binding the contract method 0x59d915ff.
//
// Solidity: function operatorSetRegistered() view returns(bool)
func (_IDurationVaultStrategy *IDurationVaultStrategyCaller) OperatorSetRegistered(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IDurationVaultStrategy.contract.Call(opts, &out, "operatorSetRegistered")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OperatorSetRegistered is a free data retrieval call binding the contract method 0x59d915ff.
//
// Solidity: function operatorSetRegistered() view returns(bool)
func (_IDurationVaultStrategy *IDurationVaultStrategySession) OperatorSetRegistered() (bool, error) {
	return _IDurationVaultStrategy.Contract.OperatorSetRegistered(&_IDurationVaultStrategy.CallOpts)
}

// OperatorSetRegistered is a free data retrieval call binding the contract method 0x59d915ff.
//
// Solidity: function operatorSetRegistered() view returns(bool)
func (_IDurationVaultStrategy *IDurationVaultStrategyCallerSession) OperatorSetRegistered() (bool, error) {
	return _IDurationVaultStrategy.Contract.OperatorSetRegistered(&_IDurationVaultStrategy.CallOpts)
}

// RewardsCoordinator is a free data retrieval call binding the contract method 0x8a2fc4e3.
//
// Solidity: function rewardsCoordinator() view returns(address)
func (_IDurationVaultStrategy *IDurationVaultStrategyCaller) RewardsCoordinator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IDurationVaultStrategy.contract.Call(opts, &out, "rewardsCoordinator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardsCoordinator is a free data retrieval call binding the contract method 0x8a2fc4e3.
//
// Solidity: function rewardsCoordinator() view returns(address)
func (_IDurationVaultStrategy *IDurationVaultStrategySession) RewardsCoordinator() (common.Address, error) {
	return _IDurationVaultStrategy.Contract.RewardsCoordinator(&_IDurationVaultStrategy.CallOpts)
}

// RewardsCoordinator is a free data retrieval call binding the contract method 0x8a2fc4e3.
//
// Solidity: function rewardsCoordinator() view returns(address)
func (_IDurationVaultStrategy *IDurationVaultStrategyCallerSession) RewardsCoordinator() (common.Address, error) {
	return _IDurationVaultStrategy.Contract.RewardsCoordinator(&_IDurationVaultStrategy.CallOpts)
}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares(address user) view returns(uint256)
func (_IDurationVaultStrategy *IDurationVaultStrategyCaller) Shares(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IDurationVaultStrategy.contract.Call(opts, &out, "shares", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares(address user) view returns(uint256)
func (_IDurationVaultStrategy *IDurationVaultStrategySession) Shares(user common.Address) (*big.Int, error) {
	return _IDurationVaultStrategy.Contract.Shares(&_IDurationVaultStrategy.CallOpts, user)
}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares(address user) view returns(uint256)
func (_IDurationVaultStrategy *IDurationVaultStrategyCallerSession) Shares(user common.Address) (*big.Int, error) {
	return _IDurationVaultStrategy.Contract.Shares(&_IDurationVaultStrategy.CallOpts, user)
}

// SharesToUnderlyingView is a free data retrieval call binding the contract method 0x7a8b2637.
//
// Solidity: function sharesToUnderlyingView(uint256 amountShares) view returns(uint256)
func (_IDurationVaultStrategy *IDurationVaultStrategyCaller) SharesToUnderlyingView(opts *bind.CallOpts, amountShares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IDurationVaultStrategy.contract.Call(opts, &out, "sharesToUnderlyingView", amountShares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SharesToUnderlyingView is a free data retrieval call binding the contract method 0x7a8b2637.
//
// Solidity: function sharesToUnderlyingView(uint256 amountShares) view returns(uint256)
func (_IDurationVaultStrategy *IDurationVaultStrategySession) SharesToUnderlyingView(amountShares *big.Int) (*big.Int, error) {
	return _IDurationVaultStrategy.Contract.SharesToUnderlyingView(&_IDurationVaultStrategy.CallOpts, amountShares)
}

// SharesToUnderlyingView is a free data retrieval call binding the contract method 0x7a8b2637.
//
// Solidity: function sharesToUnderlyingView(uint256 amountShares) view returns(uint256)
func (_IDurationVaultStrategy *IDurationVaultStrategyCallerSession) SharesToUnderlyingView(amountShares *big.Int) (*big.Int, error) {
	return _IDurationVaultStrategy.Contract.SharesToUnderlyingView(&_IDurationVaultStrategy.CallOpts, amountShares)
}

// StakeCap is a free data retrieval call binding the contract method 0xba28fd2e.
//
// Solidity: function stakeCap() view returns(uint256)
func (_IDurationVaultStrategy *IDurationVaultStrategyCaller) StakeCap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IDurationVaultStrategy.contract.Call(opts, &out, "stakeCap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeCap is a free data retrieval call binding the contract method 0xba28fd2e.
//
// Solidity: function stakeCap() view returns(uint256)
func (_IDurationVaultStrategy *IDurationVaultStrategySession) StakeCap() (*big.Int, error) {
	return _IDurationVaultStrategy.Contract.StakeCap(&_IDurationVaultStrategy.CallOpts)
}

// StakeCap is a free data retrieval call binding the contract method 0xba28fd2e.
//
// Solidity: function stakeCap() view returns(uint256)
func (_IDurationVaultStrategy *IDurationVaultStrategyCallerSession) StakeCap() (*big.Int, error) {
	return _IDurationVaultStrategy.Contract.StakeCap(&_IDurationVaultStrategy.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_IDurationVaultStrategy *IDurationVaultStrategyCaller) State(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IDurationVaultStrategy.contract.Call(opts, &out, "state")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_IDurationVaultStrategy *IDurationVaultStrategySession) State() (uint8, error) {
	return _IDurationVaultStrategy.Contract.State(&_IDurationVaultStrategy.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_IDurationVaultStrategy *IDurationVaultStrategyCallerSession) State() (uint8, error) {
	return _IDurationVaultStrategy.Contract.State(&_IDurationVaultStrategy.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_IDurationVaultStrategy *IDurationVaultStrategyCaller) TotalShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IDurationVaultStrategy.contract.Call(opts, &out, "totalShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_IDurationVaultStrategy *IDurationVaultStrategySession) TotalShares() (*big.Int, error) {
	return _IDurationVaultStrategy.Contract.TotalShares(&_IDurationVaultStrategy.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_IDurationVaultStrategy *IDurationVaultStrategyCallerSession) TotalShares() (*big.Int, error) {
	return _IDurationVaultStrategy.Contract.TotalShares(&_IDurationVaultStrategy.CallOpts)
}

// UnderlyingToSharesView is a free data retrieval call binding the contract method 0xe3dae51c.
//
// Solidity: function underlyingToSharesView(uint256 amountUnderlying) view returns(uint256)
func (_IDurationVaultStrategy *IDurationVaultStrategyCaller) UnderlyingToSharesView(opts *bind.CallOpts, amountUnderlying *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IDurationVaultStrategy.contract.Call(opts, &out, "underlyingToSharesView", amountUnderlying)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnderlyingToSharesView is a free data retrieval call binding the contract method 0xe3dae51c.
//
// Solidity: function underlyingToSharesView(uint256 amountUnderlying) view returns(uint256)
func (_IDurationVaultStrategy *IDurationVaultStrategySession) UnderlyingToSharesView(amountUnderlying *big.Int) (*big.Int, error) {
	return _IDurationVaultStrategy.Contract.UnderlyingToSharesView(&_IDurationVaultStrategy.CallOpts, amountUnderlying)
}

// UnderlyingToSharesView is a free data retrieval call binding the contract method 0xe3dae51c.
//
// Solidity: function underlyingToSharesView(uint256 amountUnderlying) view returns(uint256)
func (_IDurationVaultStrategy *IDurationVaultStrategyCallerSession) UnderlyingToSharesView(amountUnderlying *big.Int) (*big.Int, error) {
	return _IDurationVaultStrategy.Contract.UnderlyingToSharesView(&_IDurationVaultStrategy.CallOpts, amountUnderlying)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_IDurationVaultStrategy *IDurationVaultStrategyCaller) UnderlyingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IDurationVaultStrategy.contract.Call(opts, &out, "underlyingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_IDurationVaultStrategy *IDurationVaultStrategySession) UnderlyingToken() (common.Address, error) {
	return _IDurationVaultStrategy.Contract.UnderlyingToken(&_IDurationVaultStrategy.CallOpts)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_IDurationVaultStrategy *IDurationVaultStrategyCallerSession) UnderlyingToken() (common.Address, error) {
	return _IDurationVaultStrategy.Contract.UnderlyingToken(&_IDurationVaultStrategy.CallOpts)
}

// UnlockTimestamp is a free data retrieval call binding the contract method 0xaa082a9d.
//
// Solidity: function unlockTimestamp() view returns(uint32)
func (_IDurationVaultStrategy *IDurationVaultStrategyCaller) UnlockTimestamp(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _IDurationVaultStrategy.contract.Call(opts, &out, "unlockTimestamp")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// UnlockTimestamp is a free data retrieval call binding the contract method 0xaa082a9d.
//
// Solidity: function unlockTimestamp() view returns(uint32)
func (_IDurationVaultStrategy *IDurationVaultStrategySession) UnlockTimestamp() (uint32, error) {
	return _IDurationVaultStrategy.Contract.UnlockTimestamp(&_IDurationVaultStrategy.CallOpts)
}

// UnlockTimestamp is a free data retrieval call binding the contract method 0xaa082a9d.
//
// Solidity: function unlockTimestamp() view returns(uint32)
func (_IDurationVaultStrategy *IDurationVaultStrategyCallerSession) UnlockTimestamp() (uint32, error) {
	return _IDurationVaultStrategy.Contract.UnlockTimestamp(&_IDurationVaultStrategy.CallOpts)
}

// UserUnderlyingView is a free data retrieval call binding the contract method 0x553ca5f8.
//
// Solidity: function userUnderlyingView(address user) view returns(uint256)
func (_IDurationVaultStrategy *IDurationVaultStrategyCaller) UserUnderlyingView(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IDurationVaultStrategy.contract.Call(opts, &out, "userUnderlyingView", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserUnderlyingView is a free data retrieval call binding the contract method 0x553ca5f8.
//
// Solidity: function userUnderlyingView(address user) view returns(uint256)
func (_IDurationVaultStrategy *IDurationVaultStrategySession) UserUnderlyingView(user common.Address) (*big.Int, error) {
	return _IDurationVaultStrategy.Contract.UserUnderlyingView(&_IDurationVaultStrategy.CallOpts, user)
}

// UserUnderlyingView is a free data retrieval call binding the contract method 0x553ca5f8.
//
// Solidity: function userUnderlyingView(address user) view returns(uint256)
func (_IDurationVaultStrategy *IDurationVaultStrategyCallerSession) UserUnderlyingView(user common.Address) (*big.Int, error) {
	return _IDurationVaultStrategy.Contract.UserUnderlyingView(&_IDurationVaultStrategy.CallOpts, user)
}

// VaultAdmin is a free data retrieval call binding the contract method 0xe7f6f225.
//
// Solidity: function vaultAdmin() view returns(address)
func (_IDurationVaultStrategy *IDurationVaultStrategyCaller) VaultAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IDurationVaultStrategy.contract.Call(opts, &out, "vaultAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VaultAdmin is a free data retrieval call binding the contract method 0xe7f6f225.
//
// Solidity: function vaultAdmin() view returns(address)
func (_IDurationVaultStrategy *IDurationVaultStrategySession) VaultAdmin() (common.Address, error) {
	return _IDurationVaultStrategy.Contract.VaultAdmin(&_IDurationVaultStrategy.CallOpts)
}

// VaultAdmin is a free data retrieval call binding the contract method 0xe7f6f225.
//
// Solidity: function vaultAdmin() view returns(address)
func (_IDurationVaultStrategy *IDurationVaultStrategyCallerSession) VaultAdmin() (common.Address, error) {
	return _IDurationVaultStrategy.Contract.VaultAdmin(&_IDurationVaultStrategy.CallOpts)
}

// WithdrawalsOpen is a free data retrieval call binding the contract method 0x94aad677.
//
// Solidity: function withdrawalsOpen() view returns(bool)
func (_IDurationVaultStrategy *IDurationVaultStrategyCaller) WithdrawalsOpen(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IDurationVaultStrategy.contract.Call(opts, &out, "withdrawalsOpen")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WithdrawalsOpen is a free data retrieval call binding the contract method 0x94aad677.
//
// Solidity: function withdrawalsOpen() view returns(bool)
func (_IDurationVaultStrategy *IDurationVaultStrategySession) WithdrawalsOpen() (bool, error) {
	return _IDurationVaultStrategy.Contract.WithdrawalsOpen(&_IDurationVaultStrategy.CallOpts)
}

// WithdrawalsOpen is a free data retrieval call binding the contract method 0x94aad677.
//
// Solidity: function withdrawalsOpen() view returns(bool)
func (_IDurationVaultStrategy *IDurationVaultStrategyCallerSession) WithdrawalsOpen() (bool, error) {
	return _IDurationVaultStrategy.Contract.WithdrawalsOpen(&_IDurationVaultStrategy.CallOpts)
}

// AdvanceToWithdrawals is a paid mutator transaction binding the contract method 0x6325f655.
//
// Solidity: function advanceToWithdrawals() returns()
func (_IDurationVaultStrategy *IDurationVaultStrategyTransactor) AdvanceToWithdrawals(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDurationVaultStrategy.contract.Transact(opts, "advanceToWithdrawals")
}

// AdvanceToWithdrawals is a paid mutator transaction binding the contract method 0x6325f655.
//
// Solidity: function advanceToWithdrawals() returns()
func (_IDurationVaultStrategy *IDurationVaultStrategySession) AdvanceToWithdrawals() (*types.Transaction, error) {
	return _IDurationVaultStrategy.Contract.AdvanceToWithdrawals(&_IDurationVaultStrategy.TransactOpts)
}

// AdvanceToWithdrawals is a paid mutator transaction binding the contract method 0x6325f655.
//
// Solidity: function advanceToWithdrawals() returns()
func (_IDurationVaultStrategy *IDurationVaultStrategyTransactorSession) AdvanceToWithdrawals() (*types.Transaction, error) {
	return _IDurationVaultStrategy.Contract.AdvanceToWithdrawals(&_IDurationVaultStrategy.TransactOpts)
}

// BeforeAddShares is a paid mutator transaction binding the contract method 0x73e3c280.
//
// Solidity: function beforeAddShares(address staker, uint256 shares) returns()
func (_IDurationVaultStrategy *IDurationVaultStrategyTransactor) BeforeAddShares(opts *bind.TransactOpts, staker common.Address, shares *big.Int) (*types.Transaction, error) {
	return _IDurationVaultStrategy.contract.Transact(opts, "beforeAddShares", staker, shares)
}

// BeforeAddShares is a paid mutator transaction binding the contract method 0x73e3c280.
//
// Solidity: function beforeAddShares(address staker, uint256 shares) returns()
func (_IDurationVaultStrategy *IDurationVaultStrategySession) BeforeAddShares(staker common.Address, shares *big.Int) (*types.Transaction, error) {
	return _IDurationVaultStrategy.Contract.BeforeAddShares(&_IDurationVaultStrategy.TransactOpts, staker, shares)
}

// BeforeAddShares is a paid mutator transaction binding the contract method 0x73e3c280.
//
// Solidity: function beforeAddShares(address staker, uint256 shares) returns()
func (_IDurationVaultStrategy *IDurationVaultStrategyTransactorSession) BeforeAddShares(staker common.Address, shares *big.Int) (*types.Transaction, error) {
	return _IDurationVaultStrategy.Contract.BeforeAddShares(&_IDurationVaultStrategy.TransactOpts, staker, shares)
}

// BeforeRemoveShares is a paid mutator transaction binding the contract method 0x03e3e6eb.
//
// Solidity: function beforeRemoveShares(address staker, uint256 shares) returns()
func (_IDurationVaultStrategy *IDurationVaultStrategyTransactor) BeforeRemoveShares(opts *bind.TransactOpts, staker common.Address, shares *big.Int) (*types.Transaction, error) {
	return _IDurationVaultStrategy.contract.Transact(opts, "beforeRemoveShares", staker, shares)
}

// BeforeRemoveShares is a paid mutator transaction binding the contract method 0x03e3e6eb.
//
// Solidity: function beforeRemoveShares(address staker, uint256 shares) returns()
func (_IDurationVaultStrategy *IDurationVaultStrategySession) BeforeRemoveShares(staker common.Address, shares *big.Int) (*types.Transaction, error) {
	return _IDurationVaultStrategy.Contract.BeforeRemoveShares(&_IDurationVaultStrategy.TransactOpts, staker, shares)
}

// BeforeRemoveShares is a paid mutator transaction binding the contract method 0x03e3e6eb.
//
// Solidity: function beforeRemoveShares(address staker, uint256 shares) returns()
func (_IDurationVaultStrategy *IDurationVaultStrategyTransactorSession) BeforeRemoveShares(staker common.Address, shares *big.Int) (*types.Transaction, error) {
	return _IDurationVaultStrategy.Contract.BeforeRemoveShares(&_IDurationVaultStrategy.TransactOpts, staker, shares)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns(uint256)
func (_IDurationVaultStrategy *IDurationVaultStrategyTransactor) Deposit(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IDurationVaultStrategy.contract.Transact(opts, "deposit", token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns(uint256)
func (_IDurationVaultStrategy *IDurationVaultStrategySession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IDurationVaultStrategy.Contract.Deposit(&_IDurationVaultStrategy.TransactOpts, token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns(uint256)
func (_IDurationVaultStrategy *IDurationVaultStrategyTransactorSession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IDurationVaultStrategy.Contract.Deposit(&_IDurationVaultStrategy.TransactOpts, token, amount)
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_IDurationVaultStrategy *IDurationVaultStrategyTransactor) Lock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDurationVaultStrategy.contract.Transact(opts, "lock")
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_IDurationVaultStrategy *IDurationVaultStrategySession) Lock() (*types.Transaction, error) {
	return _IDurationVaultStrategy.Contract.Lock(&_IDurationVaultStrategy.TransactOpts)
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_IDurationVaultStrategy *IDurationVaultStrategyTransactorSession) Lock() (*types.Transaction, error) {
	return _IDurationVaultStrategy.Contract.Lock(&_IDurationVaultStrategy.TransactOpts)
}

// MarkMatured is a paid mutator transaction binding the contract method 0x6d8690a9.
//
// Solidity: function markMatured() returns()
func (_IDurationVaultStrategy *IDurationVaultStrategyTransactor) MarkMatured(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDurationVaultStrategy.contract.Transact(opts, "markMatured")
}

// MarkMatured is a paid mutator transaction binding the contract method 0x6d8690a9.
//
// Solidity: function markMatured() returns()
func (_IDurationVaultStrategy *IDurationVaultStrategySession) MarkMatured() (*types.Transaction, error) {
	return _IDurationVaultStrategy.Contract.MarkMatured(&_IDurationVaultStrategy.TransactOpts)
}

// MarkMatured is a paid mutator transaction binding the contract method 0x6d8690a9.
//
// Solidity: function markMatured() returns()
func (_IDurationVaultStrategy *IDurationVaultStrategyTransactorSession) MarkMatured() (*types.Transaction, error) {
	return _IDurationVaultStrategy.Contract.MarkMatured(&_IDurationVaultStrategy.TransactOpts)
}

// SetRewardsClaimer is a paid mutator transaction binding the contract method 0xb501d660.
//
// Solidity: function setRewardsClaimer(address claimer) returns()
func (_IDurationVaultStrategy *IDurationVaultStrategyTransactor) SetRewardsClaimer(opts *bind.TransactOpts, claimer common.Address) (*types.Transaction, error) {
	return _IDurationVaultStrategy.contract.Transact(opts, "setRewardsClaimer", claimer)
}

// SetRewardsClaimer is a paid mutator transaction binding the contract method 0xb501d660.
//
// Solidity: function setRewardsClaimer(address claimer) returns()
func (_IDurationVaultStrategy *IDurationVaultStrategySession) SetRewardsClaimer(claimer common.Address) (*types.Transaction, error) {
	return _IDurationVaultStrategy.Contract.SetRewardsClaimer(&_IDurationVaultStrategy.TransactOpts, claimer)
}

// SetRewardsClaimer is a paid mutator transaction binding the contract method 0xb501d660.
//
// Solidity: function setRewardsClaimer(address claimer) returns()
func (_IDurationVaultStrategy *IDurationVaultStrategyTransactorSession) SetRewardsClaimer(claimer common.Address) (*types.Transaction, error) {
	return _IDurationVaultStrategy.Contract.SetRewardsClaimer(&_IDurationVaultStrategy.TransactOpts, claimer)
}

// SharesToUnderlying is a paid mutator transaction binding the contract method 0xf3e73875.
//
// Solidity: function sharesToUnderlying(uint256 amountShares) returns(uint256)
func (_IDurationVaultStrategy *IDurationVaultStrategyTransactor) SharesToUnderlying(opts *bind.TransactOpts, amountShares *big.Int) (*types.Transaction, error) {
	return _IDurationVaultStrategy.contract.Transact(opts, "sharesToUnderlying", amountShares)
}

// SharesToUnderlying is a paid mutator transaction binding the contract method 0xf3e73875.
//
// Solidity: function sharesToUnderlying(uint256 amountShares) returns(uint256)
func (_IDurationVaultStrategy *IDurationVaultStrategySession) SharesToUnderlying(amountShares *big.Int) (*types.Transaction, error) {
	return _IDurationVaultStrategy.Contract.SharesToUnderlying(&_IDurationVaultStrategy.TransactOpts, amountShares)
}

// SharesToUnderlying is a paid mutator transaction binding the contract method 0xf3e73875.
//
// Solidity: function sharesToUnderlying(uint256 amountShares) returns(uint256)
func (_IDurationVaultStrategy *IDurationVaultStrategyTransactorSession) SharesToUnderlying(amountShares *big.Int) (*types.Transaction, error) {
	return _IDurationVaultStrategy.Contract.SharesToUnderlying(&_IDurationVaultStrategy.TransactOpts, amountShares)
}

// UnderlyingToShares is a paid mutator transaction binding the contract method 0x8c871019.
//
// Solidity: function underlyingToShares(uint256 amountUnderlying) returns(uint256)
func (_IDurationVaultStrategy *IDurationVaultStrategyTransactor) UnderlyingToShares(opts *bind.TransactOpts, amountUnderlying *big.Int) (*types.Transaction, error) {
	return _IDurationVaultStrategy.contract.Transact(opts, "underlyingToShares", amountUnderlying)
}

// UnderlyingToShares is a paid mutator transaction binding the contract method 0x8c871019.
//
// Solidity: function underlyingToShares(uint256 amountUnderlying) returns(uint256)
func (_IDurationVaultStrategy *IDurationVaultStrategySession) UnderlyingToShares(amountUnderlying *big.Int) (*types.Transaction, error) {
	return _IDurationVaultStrategy.Contract.UnderlyingToShares(&_IDurationVaultStrategy.TransactOpts, amountUnderlying)
}

// UnderlyingToShares is a paid mutator transaction binding the contract method 0x8c871019.
//
// Solidity: function underlyingToShares(uint256 amountUnderlying) returns(uint256)
func (_IDurationVaultStrategy *IDurationVaultStrategyTransactorSession) UnderlyingToShares(amountUnderlying *big.Int) (*types.Transaction, error) {
	return _IDurationVaultStrategy.Contract.UnderlyingToShares(&_IDurationVaultStrategy.TransactOpts, amountUnderlying)
}

// UpdateDelegationApprover is a paid mutator transaction binding the contract method 0xb4e20f13.
//
// Solidity: function updateDelegationApprover(address newDelegationApprover) returns()
func (_IDurationVaultStrategy *IDurationVaultStrategyTransactor) UpdateDelegationApprover(opts *bind.TransactOpts, newDelegationApprover common.Address) (*types.Transaction, error) {
	return _IDurationVaultStrategy.contract.Transact(opts, "updateDelegationApprover", newDelegationApprover)
}

// UpdateDelegationApprover is a paid mutator transaction binding the contract method 0xb4e20f13.
//
// Solidity: function updateDelegationApprover(address newDelegationApprover) returns()
func (_IDurationVaultStrategy *IDurationVaultStrategySession) UpdateDelegationApprover(newDelegationApprover common.Address) (*types.Transaction, error) {
	return _IDurationVaultStrategy.Contract.UpdateDelegationApprover(&_IDurationVaultStrategy.TransactOpts, newDelegationApprover)
}

// UpdateDelegationApprover is a paid mutator transaction binding the contract method 0xb4e20f13.
//
// Solidity: function updateDelegationApprover(address newDelegationApprover) returns()
func (_IDurationVaultStrategy *IDurationVaultStrategyTransactorSession) UpdateDelegationApprover(newDelegationApprover common.Address) (*types.Transaction, error) {
	return _IDurationVaultStrategy.Contract.UpdateDelegationApprover(&_IDurationVaultStrategy.TransactOpts, newDelegationApprover)
}

// UpdateMetadataURI is a paid mutator transaction binding the contract method 0x53fd3e81.
//
// Solidity: function updateMetadataURI(string newMetadataURI) returns()
func (_IDurationVaultStrategy *IDurationVaultStrategyTransactor) UpdateMetadataURI(opts *bind.TransactOpts, newMetadataURI string) (*types.Transaction, error) {
	return _IDurationVaultStrategy.contract.Transact(opts, "updateMetadataURI", newMetadataURI)
}

// UpdateMetadataURI is a paid mutator transaction binding the contract method 0x53fd3e81.
//
// Solidity: function updateMetadataURI(string newMetadataURI) returns()
func (_IDurationVaultStrategy *IDurationVaultStrategySession) UpdateMetadataURI(newMetadataURI string) (*types.Transaction, error) {
	return _IDurationVaultStrategy.Contract.UpdateMetadataURI(&_IDurationVaultStrategy.TransactOpts, newMetadataURI)
}

// UpdateMetadataURI is a paid mutator transaction binding the contract method 0x53fd3e81.
//
// Solidity: function updateMetadataURI(string newMetadataURI) returns()
func (_IDurationVaultStrategy *IDurationVaultStrategyTransactorSession) UpdateMetadataURI(newMetadataURI string) (*types.Transaction, error) {
	return _IDurationVaultStrategy.Contract.UpdateMetadataURI(&_IDurationVaultStrategy.TransactOpts, newMetadataURI)
}

// UpdateOperatorMetadataURI is a paid mutator transaction binding the contract method 0x99be81c8.
//
// Solidity: function updateOperatorMetadataURI(string newOperatorMetadataURI) returns()
func (_IDurationVaultStrategy *IDurationVaultStrategyTransactor) UpdateOperatorMetadataURI(opts *bind.TransactOpts, newOperatorMetadataURI string) (*types.Transaction, error) {
	return _IDurationVaultStrategy.contract.Transact(opts, "updateOperatorMetadataURI", newOperatorMetadataURI)
}

// UpdateOperatorMetadataURI is a paid mutator transaction binding the contract method 0x99be81c8.
//
// Solidity: function updateOperatorMetadataURI(string newOperatorMetadataURI) returns()
func (_IDurationVaultStrategy *IDurationVaultStrategySession) UpdateOperatorMetadataURI(newOperatorMetadataURI string) (*types.Transaction, error) {
	return _IDurationVaultStrategy.Contract.UpdateOperatorMetadataURI(&_IDurationVaultStrategy.TransactOpts, newOperatorMetadataURI)
}

// UpdateOperatorMetadataURI is a paid mutator transaction binding the contract method 0x99be81c8.
//
// Solidity: function updateOperatorMetadataURI(string newOperatorMetadataURI) returns()
func (_IDurationVaultStrategy *IDurationVaultStrategyTransactorSession) UpdateOperatorMetadataURI(newOperatorMetadataURI string) (*types.Transaction, error) {
	return _IDurationVaultStrategy.Contract.UpdateOperatorMetadataURI(&_IDurationVaultStrategy.TransactOpts, newOperatorMetadataURI)
}

// UpdateTVLLimits is a paid mutator transaction binding the contract method 0xaf6eb2be.
//
// Solidity: function updateTVLLimits(uint256 newMaxPerDeposit, uint256 newStakeCap) returns()
func (_IDurationVaultStrategy *IDurationVaultStrategyTransactor) UpdateTVLLimits(opts *bind.TransactOpts, newMaxPerDeposit *big.Int, newStakeCap *big.Int) (*types.Transaction, error) {
	return _IDurationVaultStrategy.contract.Transact(opts, "updateTVLLimits", newMaxPerDeposit, newStakeCap)
}

// UpdateTVLLimits is a paid mutator transaction binding the contract method 0xaf6eb2be.
//
// Solidity: function updateTVLLimits(uint256 newMaxPerDeposit, uint256 newStakeCap) returns()
func (_IDurationVaultStrategy *IDurationVaultStrategySession) UpdateTVLLimits(newMaxPerDeposit *big.Int, newStakeCap *big.Int) (*types.Transaction, error) {
	return _IDurationVaultStrategy.Contract.UpdateTVLLimits(&_IDurationVaultStrategy.TransactOpts, newMaxPerDeposit, newStakeCap)
}

// UpdateTVLLimits is a paid mutator transaction binding the contract method 0xaf6eb2be.
//
// Solidity: function updateTVLLimits(uint256 newMaxPerDeposit, uint256 newStakeCap) returns()
func (_IDurationVaultStrategy *IDurationVaultStrategyTransactorSession) UpdateTVLLimits(newMaxPerDeposit *big.Int, newStakeCap *big.Int) (*types.Transaction, error) {
	return _IDurationVaultStrategy.Contract.UpdateTVLLimits(&_IDurationVaultStrategy.TransactOpts, newMaxPerDeposit, newStakeCap)
}

// UserUnderlying is a paid mutator transaction binding the contract method 0x8f6a6240.
//
// Solidity: function userUnderlying(address user) returns(uint256)
func (_IDurationVaultStrategy *IDurationVaultStrategyTransactor) UserUnderlying(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _IDurationVaultStrategy.contract.Transact(opts, "userUnderlying", user)
}

// UserUnderlying is a paid mutator transaction binding the contract method 0x8f6a6240.
//
// Solidity: function userUnderlying(address user) returns(uint256)
func (_IDurationVaultStrategy *IDurationVaultStrategySession) UserUnderlying(user common.Address) (*types.Transaction, error) {
	return _IDurationVaultStrategy.Contract.UserUnderlying(&_IDurationVaultStrategy.TransactOpts, user)
}

// UserUnderlying is a paid mutator transaction binding the contract method 0x8f6a6240.
//
// Solidity: function userUnderlying(address user) returns(uint256)
func (_IDurationVaultStrategy *IDurationVaultStrategyTransactorSession) UserUnderlying(user common.Address) (*types.Transaction, error) {
	return _IDurationVaultStrategy.Contract.UserUnderlying(&_IDurationVaultStrategy.TransactOpts, user)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address recipient, address token, uint256 amountShares) returns(uint256)
func (_IDurationVaultStrategy *IDurationVaultStrategyTransactor) Withdraw(opts *bind.TransactOpts, recipient common.Address, token common.Address, amountShares *big.Int) (*types.Transaction, error) {
	return _IDurationVaultStrategy.contract.Transact(opts, "withdraw", recipient, token, amountShares)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address recipient, address token, uint256 amountShares) returns(uint256)
func (_IDurationVaultStrategy *IDurationVaultStrategySession) Withdraw(recipient common.Address, token common.Address, amountShares *big.Int) (*types.Transaction, error) {
	return _IDurationVaultStrategy.Contract.Withdraw(&_IDurationVaultStrategy.TransactOpts, recipient, token, amountShares)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address recipient, address token, uint256 amountShares) returns(uint256)
func (_IDurationVaultStrategy *IDurationVaultStrategyTransactorSession) Withdraw(recipient common.Address, token common.Address, amountShares *big.Int) (*types.Transaction, error) {
	return _IDurationVaultStrategy.Contract.Withdraw(&_IDurationVaultStrategy.TransactOpts, recipient, token, amountShares)
}

// IDurationVaultStrategyDeallocateAttemptedIterator is returned from FilterDeallocateAttempted and is used to iterate over the raw logs and unpacked data for DeallocateAttempted events raised by the IDurationVaultStrategy contract.
type IDurationVaultStrategyDeallocateAttemptedIterator struct {
	Event *IDurationVaultStrategyDeallocateAttempted // Event containing the contract specifics and raw log

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
func (it *IDurationVaultStrategyDeallocateAttemptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDurationVaultStrategyDeallocateAttempted)
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
		it.Event = new(IDurationVaultStrategyDeallocateAttempted)
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
func (it *IDurationVaultStrategyDeallocateAttemptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDurationVaultStrategyDeallocateAttemptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDurationVaultStrategyDeallocateAttempted represents a DeallocateAttempted event raised by the IDurationVaultStrategy contract.
type IDurationVaultStrategyDeallocateAttempted struct {
	Success bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeallocateAttempted is a free log retrieval operation binding the contract event 0x72f957da7daaea6b52e4ff7820cb464206fd51e9f502f3027f45b5017caf4c8b.
//
// Solidity: event DeallocateAttempted(bool success)
func (_IDurationVaultStrategy *IDurationVaultStrategyFilterer) FilterDeallocateAttempted(opts *bind.FilterOpts) (*IDurationVaultStrategyDeallocateAttemptedIterator, error) {

	logs, sub, err := _IDurationVaultStrategy.contract.FilterLogs(opts, "DeallocateAttempted")
	if err != nil {
		return nil, err
	}
	return &IDurationVaultStrategyDeallocateAttemptedIterator{contract: _IDurationVaultStrategy.contract, event: "DeallocateAttempted", logs: logs, sub: sub}, nil
}

// WatchDeallocateAttempted is a free log subscription operation binding the contract event 0x72f957da7daaea6b52e4ff7820cb464206fd51e9f502f3027f45b5017caf4c8b.
//
// Solidity: event DeallocateAttempted(bool success)
func (_IDurationVaultStrategy *IDurationVaultStrategyFilterer) WatchDeallocateAttempted(opts *bind.WatchOpts, sink chan<- *IDurationVaultStrategyDeallocateAttempted) (event.Subscription, error) {

	logs, sub, err := _IDurationVaultStrategy.contract.WatchLogs(opts, "DeallocateAttempted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDurationVaultStrategyDeallocateAttempted)
				if err := _IDurationVaultStrategy.contract.UnpackLog(event, "DeallocateAttempted", log); err != nil {
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
func (_IDurationVaultStrategy *IDurationVaultStrategyFilterer) ParseDeallocateAttempted(log types.Log) (*IDurationVaultStrategyDeallocateAttempted, error) {
	event := new(IDurationVaultStrategyDeallocateAttempted)
	if err := _IDurationVaultStrategy.contract.UnpackLog(event, "DeallocateAttempted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDurationVaultStrategyDeregisterAttemptedIterator is returned from FilterDeregisterAttempted and is used to iterate over the raw logs and unpacked data for DeregisterAttempted events raised by the IDurationVaultStrategy contract.
type IDurationVaultStrategyDeregisterAttemptedIterator struct {
	Event *IDurationVaultStrategyDeregisterAttempted // Event containing the contract specifics and raw log

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
func (it *IDurationVaultStrategyDeregisterAttemptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDurationVaultStrategyDeregisterAttempted)
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
		it.Event = new(IDurationVaultStrategyDeregisterAttempted)
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
func (it *IDurationVaultStrategyDeregisterAttemptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDurationVaultStrategyDeregisterAttemptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDurationVaultStrategyDeregisterAttempted represents a DeregisterAttempted event raised by the IDurationVaultStrategy contract.
type IDurationVaultStrategyDeregisterAttempted struct {
	Success bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeregisterAttempted is a free log retrieval operation binding the contract event 0xd0791dbc9180cb64588d7eb7658a1022dcf734b8825eb7eec68bd9516872d168.
//
// Solidity: event DeregisterAttempted(bool success)
func (_IDurationVaultStrategy *IDurationVaultStrategyFilterer) FilterDeregisterAttempted(opts *bind.FilterOpts) (*IDurationVaultStrategyDeregisterAttemptedIterator, error) {

	logs, sub, err := _IDurationVaultStrategy.contract.FilterLogs(opts, "DeregisterAttempted")
	if err != nil {
		return nil, err
	}
	return &IDurationVaultStrategyDeregisterAttemptedIterator{contract: _IDurationVaultStrategy.contract, event: "DeregisterAttempted", logs: logs, sub: sub}, nil
}

// WatchDeregisterAttempted is a free log subscription operation binding the contract event 0xd0791dbc9180cb64588d7eb7658a1022dcf734b8825eb7eec68bd9516872d168.
//
// Solidity: event DeregisterAttempted(bool success)
func (_IDurationVaultStrategy *IDurationVaultStrategyFilterer) WatchDeregisterAttempted(opts *bind.WatchOpts, sink chan<- *IDurationVaultStrategyDeregisterAttempted) (event.Subscription, error) {

	logs, sub, err := _IDurationVaultStrategy.contract.WatchLogs(opts, "DeregisterAttempted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDurationVaultStrategyDeregisterAttempted)
				if err := _IDurationVaultStrategy.contract.UnpackLog(event, "DeregisterAttempted", log); err != nil {
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
func (_IDurationVaultStrategy *IDurationVaultStrategyFilterer) ParseDeregisterAttempted(log types.Log) (*IDurationVaultStrategyDeregisterAttempted, error) {
	event := new(IDurationVaultStrategyDeregisterAttempted)
	if err := _IDurationVaultStrategy.contract.UnpackLog(event, "DeregisterAttempted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDurationVaultStrategyExchangeRateEmittedIterator is returned from FilterExchangeRateEmitted and is used to iterate over the raw logs and unpacked data for ExchangeRateEmitted events raised by the IDurationVaultStrategy contract.
type IDurationVaultStrategyExchangeRateEmittedIterator struct {
	Event *IDurationVaultStrategyExchangeRateEmitted // Event containing the contract specifics and raw log

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
func (it *IDurationVaultStrategyExchangeRateEmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDurationVaultStrategyExchangeRateEmitted)
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
		it.Event = new(IDurationVaultStrategyExchangeRateEmitted)
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
func (it *IDurationVaultStrategyExchangeRateEmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDurationVaultStrategyExchangeRateEmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDurationVaultStrategyExchangeRateEmitted represents a ExchangeRateEmitted event raised by the IDurationVaultStrategy contract.
type IDurationVaultStrategyExchangeRateEmitted struct {
	Rate *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterExchangeRateEmitted is a free log retrieval operation binding the contract event 0xd2494f3479e5da49d386657c292c610b5b01df313d07c62eb0cfa49924a31be8.
//
// Solidity: event ExchangeRateEmitted(uint256 rate)
func (_IDurationVaultStrategy *IDurationVaultStrategyFilterer) FilterExchangeRateEmitted(opts *bind.FilterOpts) (*IDurationVaultStrategyExchangeRateEmittedIterator, error) {

	logs, sub, err := _IDurationVaultStrategy.contract.FilterLogs(opts, "ExchangeRateEmitted")
	if err != nil {
		return nil, err
	}
	return &IDurationVaultStrategyExchangeRateEmittedIterator{contract: _IDurationVaultStrategy.contract, event: "ExchangeRateEmitted", logs: logs, sub: sub}, nil
}

// WatchExchangeRateEmitted is a free log subscription operation binding the contract event 0xd2494f3479e5da49d386657c292c610b5b01df313d07c62eb0cfa49924a31be8.
//
// Solidity: event ExchangeRateEmitted(uint256 rate)
func (_IDurationVaultStrategy *IDurationVaultStrategyFilterer) WatchExchangeRateEmitted(opts *bind.WatchOpts, sink chan<- *IDurationVaultStrategyExchangeRateEmitted) (event.Subscription, error) {

	logs, sub, err := _IDurationVaultStrategy.contract.WatchLogs(opts, "ExchangeRateEmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDurationVaultStrategyExchangeRateEmitted)
				if err := _IDurationVaultStrategy.contract.UnpackLog(event, "ExchangeRateEmitted", log); err != nil {
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
func (_IDurationVaultStrategy *IDurationVaultStrategyFilterer) ParseExchangeRateEmitted(log types.Log) (*IDurationVaultStrategyExchangeRateEmitted, error) {
	event := new(IDurationVaultStrategyExchangeRateEmitted)
	if err := _IDurationVaultStrategy.contract.UnpackLog(event, "ExchangeRateEmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDurationVaultStrategyMaxPerDepositUpdatedIterator is returned from FilterMaxPerDepositUpdated and is used to iterate over the raw logs and unpacked data for MaxPerDepositUpdated events raised by the IDurationVaultStrategy contract.
type IDurationVaultStrategyMaxPerDepositUpdatedIterator struct {
	Event *IDurationVaultStrategyMaxPerDepositUpdated // Event containing the contract specifics and raw log

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
func (it *IDurationVaultStrategyMaxPerDepositUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDurationVaultStrategyMaxPerDepositUpdated)
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
		it.Event = new(IDurationVaultStrategyMaxPerDepositUpdated)
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
func (it *IDurationVaultStrategyMaxPerDepositUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDurationVaultStrategyMaxPerDepositUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDurationVaultStrategyMaxPerDepositUpdated represents a MaxPerDepositUpdated event raised by the IDurationVaultStrategy contract.
type IDurationVaultStrategyMaxPerDepositUpdated struct {
	PreviousValue *big.Int
	NewValue      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMaxPerDepositUpdated is a free log retrieval operation binding the contract event 0xf97ed4e083acac67830025ecbc756d8fe847cdbdca4cee3fe1e128e98b54ecb5.
//
// Solidity: event MaxPerDepositUpdated(uint256 previousValue, uint256 newValue)
func (_IDurationVaultStrategy *IDurationVaultStrategyFilterer) FilterMaxPerDepositUpdated(opts *bind.FilterOpts) (*IDurationVaultStrategyMaxPerDepositUpdatedIterator, error) {

	logs, sub, err := _IDurationVaultStrategy.contract.FilterLogs(opts, "MaxPerDepositUpdated")
	if err != nil {
		return nil, err
	}
	return &IDurationVaultStrategyMaxPerDepositUpdatedIterator{contract: _IDurationVaultStrategy.contract, event: "MaxPerDepositUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxPerDepositUpdated is a free log subscription operation binding the contract event 0xf97ed4e083acac67830025ecbc756d8fe847cdbdca4cee3fe1e128e98b54ecb5.
//
// Solidity: event MaxPerDepositUpdated(uint256 previousValue, uint256 newValue)
func (_IDurationVaultStrategy *IDurationVaultStrategyFilterer) WatchMaxPerDepositUpdated(opts *bind.WatchOpts, sink chan<- *IDurationVaultStrategyMaxPerDepositUpdated) (event.Subscription, error) {

	logs, sub, err := _IDurationVaultStrategy.contract.WatchLogs(opts, "MaxPerDepositUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDurationVaultStrategyMaxPerDepositUpdated)
				if err := _IDurationVaultStrategy.contract.UnpackLog(event, "MaxPerDepositUpdated", log); err != nil {
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
func (_IDurationVaultStrategy *IDurationVaultStrategyFilterer) ParseMaxPerDepositUpdated(log types.Log) (*IDurationVaultStrategyMaxPerDepositUpdated, error) {
	event := new(IDurationVaultStrategyMaxPerDepositUpdated)
	if err := _IDurationVaultStrategy.contract.UnpackLog(event, "MaxPerDepositUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDurationVaultStrategyMaxTotalDepositsUpdatedIterator is returned from FilterMaxTotalDepositsUpdated and is used to iterate over the raw logs and unpacked data for MaxTotalDepositsUpdated events raised by the IDurationVaultStrategy contract.
type IDurationVaultStrategyMaxTotalDepositsUpdatedIterator struct {
	Event *IDurationVaultStrategyMaxTotalDepositsUpdated // Event containing the contract specifics and raw log

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
func (it *IDurationVaultStrategyMaxTotalDepositsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDurationVaultStrategyMaxTotalDepositsUpdated)
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
		it.Event = new(IDurationVaultStrategyMaxTotalDepositsUpdated)
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
func (it *IDurationVaultStrategyMaxTotalDepositsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDurationVaultStrategyMaxTotalDepositsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDurationVaultStrategyMaxTotalDepositsUpdated represents a MaxTotalDepositsUpdated event raised by the IDurationVaultStrategy contract.
type IDurationVaultStrategyMaxTotalDepositsUpdated struct {
	PreviousValue *big.Int
	NewValue      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMaxTotalDepositsUpdated is a free log retrieval operation binding the contract event 0x6ab181e0440bfbf4bacdf2e99674735ce6638005490688c5f994f5399353e452.
//
// Solidity: event MaxTotalDepositsUpdated(uint256 previousValue, uint256 newValue)
func (_IDurationVaultStrategy *IDurationVaultStrategyFilterer) FilterMaxTotalDepositsUpdated(opts *bind.FilterOpts) (*IDurationVaultStrategyMaxTotalDepositsUpdatedIterator, error) {

	logs, sub, err := _IDurationVaultStrategy.contract.FilterLogs(opts, "MaxTotalDepositsUpdated")
	if err != nil {
		return nil, err
	}
	return &IDurationVaultStrategyMaxTotalDepositsUpdatedIterator{contract: _IDurationVaultStrategy.contract, event: "MaxTotalDepositsUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxTotalDepositsUpdated is a free log subscription operation binding the contract event 0x6ab181e0440bfbf4bacdf2e99674735ce6638005490688c5f994f5399353e452.
//
// Solidity: event MaxTotalDepositsUpdated(uint256 previousValue, uint256 newValue)
func (_IDurationVaultStrategy *IDurationVaultStrategyFilterer) WatchMaxTotalDepositsUpdated(opts *bind.WatchOpts, sink chan<- *IDurationVaultStrategyMaxTotalDepositsUpdated) (event.Subscription, error) {

	logs, sub, err := _IDurationVaultStrategy.contract.WatchLogs(opts, "MaxTotalDepositsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDurationVaultStrategyMaxTotalDepositsUpdated)
				if err := _IDurationVaultStrategy.contract.UnpackLog(event, "MaxTotalDepositsUpdated", log); err != nil {
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
func (_IDurationVaultStrategy *IDurationVaultStrategyFilterer) ParseMaxTotalDepositsUpdated(log types.Log) (*IDurationVaultStrategyMaxTotalDepositsUpdated, error) {
	event := new(IDurationVaultStrategyMaxTotalDepositsUpdated)
	if err := _IDurationVaultStrategy.contract.UnpackLog(event, "MaxTotalDepositsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDurationVaultStrategyMetadataURIUpdatedIterator is returned from FilterMetadataURIUpdated and is used to iterate over the raw logs and unpacked data for MetadataURIUpdated events raised by the IDurationVaultStrategy contract.
type IDurationVaultStrategyMetadataURIUpdatedIterator struct {
	Event *IDurationVaultStrategyMetadataURIUpdated // Event containing the contract specifics and raw log

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
func (it *IDurationVaultStrategyMetadataURIUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDurationVaultStrategyMetadataURIUpdated)
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
		it.Event = new(IDurationVaultStrategyMetadataURIUpdated)
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
func (it *IDurationVaultStrategyMetadataURIUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDurationVaultStrategyMetadataURIUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDurationVaultStrategyMetadataURIUpdated represents a MetadataURIUpdated event raised by the IDurationVaultStrategy contract.
type IDurationVaultStrategyMetadataURIUpdated struct {
	NewMetadataURI string
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterMetadataURIUpdated is a free log retrieval operation binding the contract event 0xefafb90526da1636e1335eac0151301742fb755d986954c613b90e891778ba39.
//
// Solidity: event MetadataURIUpdated(string newMetadataURI)
func (_IDurationVaultStrategy *IDurationVaultStrategyFilterer) FilterMetadataURIUpdated(opts *bind.FilterOpts) (*IDurationVaultStrategyMetadataURIUpdatedIterator, error) {

	logs, sub, err := _IDurationVaultStrategy.contract.FilterLogs(opts, "MetadataURIUpdated")
	if err != nil {
		return nil, err
	}
	return &IDurationVaultStrategyMetadataURIUpdatedIterator{contract: _IDurationVaultStrategy.contract, event: "MetadataURIUpdated", logs: logs, sub: sub}, nil
}

// WatchMetadataURIUpdated is a free log subscription operation binding the contract event 0xefafb90526da1636e1335eac0151301742fb755d986954c613b90e891778ba39.
//
// Solidity: event MetadataURIUpdated(string newMetadataURI)
func (_IDurationVaultStrategy *IDurationVaultStrategyFilterer) WatchMetadataURIUpdated(opts *bind.WatchOpts, sink chan<- *IDurationVaultStrategyMetadataURIUpdated) (event.Subscription, error) {

	logs, sub, err := _IDurationVaultStrategy.contract.WatchLogs(opts, "MetadataURIUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDurationVaultStrategyMetadataURIUpdated)
				if err := _IDurationVaultStrategy.contract.UnpackLog(event, "MetadataURIUpdated", log); err != nil {
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
func (_IDurationVaultStrategy *IDurationVaultStrategyFilterer) ParseMetadataURIUpdated(log types.Log) (*IDurationVaultStrategyMetadataURIUpdated, error) {
	event := new(IDurationVaultStrategyMetadataURIUpdated)
	if err := _IDurationVaultStrategy.contract.UnpackLog(event, "MetadataURIUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDurationVaultStrategyStrategyTokenSetIterator is returned from FilterStrategyTokenSet and is used to iterate over the raw logs and unpacked data for StrategyTokenSet events raised by the IDurationVaultStrategy contract.
type IDurationVaultStrategyStrategyTokenSetIterator struct {
	Event *IDurationVaultStrategyStrategyTokenSet // Event containing the contract specifics and raw log

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
func (it *IDurationVaultStrategyStrategyTokenSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDurationVaultStrategyStrategyTokenSet)
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
		it.Event = new(IDurationVaultStrategyStrategyTokenSet)
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
func (it *IDurationVaultStrategyStrategyTokenSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDurationVaultStrategyStrategyTokenSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDurationVaultStrategyStrategyTokenSet represents a StrategyTokenSet event raised by the IDurationVaultStrategy contract.
type IDurationVaultStrategyStrategyTokenSet struct {
	Token    common.Address
	Decimals uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStrategyTokenSet is a free log retrieval operation binding the contract event 0x1c540707b00eb5427b6b774fc799d756516a54aee108b64b327acc55af557507.
//
// Solidity: event StrategyTokenSet(address token, uint8 decimals)
func (_IDurationVaultStrategy *IDurationVaultStrategyFilterer) FilterStrategyTokenSet(opts *bind.FilterOpts) (*IDurationVaultStrategyStrategyTokenSetIterator, error) {

	logs, sub, err := _IDurationVaultStrategy.contract.FilterLogs(opts, "StrategyTokenSet")
	if err != nil {
		return nil, err
	}
	return &IDurationVaultStrategyStrategyTokenSetIterator{contract: _IDurationVaultStrategy.contract, event: "StrategyTokenSet", logs: logs, sub: sub}, nil
}

// WatchStrategyTokenSet is a free log subscription operation binding the contract event 0x1c540707b00eb5427b6b774fc799d756516a54aee108b64b327acc55af557507.
//
// Solidity: event StrategyTokenSet(address token, uint8 decimals)
func (_IDurationVaultStrategy *IDurationVaultStrategyFilterer) WatchStrategyTokenSet(opts *bind.WatchOpts, sink chan<- *IDurationVaultStrategyStrategyTokenSet) (event.Subscription, error) {

	logs, sub, err := _IDurationVaultStrategy.contract.WatchLogs(opts, "StrategyTokenSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDurationVaultStrategyStrategyTokenSet)
				if err := _IDurationVaultStrategy.contract.UnpackLog(event, "StrategyTokenSet", log); err != nil {
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
func (_IDurationVaultStrategy *IDurationVaultStrategyFilterer) ParseStrategyTokenSet(log types.Log) (*IDurationVaultStrategyStrategyTokenSet, error) {
	event := new(IDurationVaultStrategyStrategyTokenSet)
	if err := _IDurationVaultStrategy.contract.UnpackLog(event, "StrategyTokenSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDurationVaultStrategyVaultAdvancedToWithdrawalsIterator is returned from FilterVaultAdvancedToWithdrawals and is used to iterate over the raw logs and unpacked data for VaultAdvancedToWithdrawals events raised by the IDurationVaultStrategy contract.
type IDurationVaultStrategyVaultAdvancedToWithdrawalsIterator struct {
	Event *IDurationVaultStrategyVaultAdvancedToWithdrawals // Event containing the contract specifics and raw log

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
func (it *IDurationVaultStrategyVaultAdvancedToWithdrawalsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDurationVaultStrategyVaultAdvancedToWithdrawals)
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
		it.Event = new(IDurationVaultStrategyVaultAdvancedToWithdrawals)
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
func (it *IDurationVaultStrategyVaultAdvancedToWithdrawalsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDurationVaultStrategyVaultAdvancedToWithdrawalsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDurationVaultStrategyVaultAdvancedToWithdrawals represents a VaultAdvancedToWithdrawals event raised by the IDurationVaultStrategy contract.
type IDurationVaultStrategyVaultAdvancedToWithdrawals struct {
	Arbitrator common.Address
	MaturedAt  uint32
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVaultAdvancedToWithdrawals is a free log retrieval operation binding the contract event 0x96c49d03ef64591194500229a104cd087b2d45c68234c96444c3a2a6abb0bb97.
//
// Solidity: event VaultAdvancedToWithdrawals(address indexed arbitrator, uint32 maturedAt)
func (_IDurationVaultStrategy *IDurationVaultStrategyFilterer) FilterVaultAdvancedToWithdrawals(opts *bind.FilterOpts, arbitrator []common.Address) (*IDurationVaultStrategyVaultAdvancedToWithdrawalsIterator, error) {

	var arbitratorRule []interface{}
	for _, arbitratorItem := range arbitrator {
		arbitratorRule = append(arbitratorRule, arbitratorItem)
	}

	logs, sub, err := _IDurationVaultStrategy.contract.FilterLogs(opts, "VaultAdvancedToWithdrawals", arbitratorRule)
	if err != nil {
		return nil, err
	}
	return &IDurationVaultStrategyVaultAdvancedToWithdrawalsIterator{contract: _IDurationVaultStrategy.contract, event: "VaultAdvancedToWithdrawals", logs: logs, sub: sub}, nil
}

// WatchVaultAdvancedToWithdrawals is a free log subscription operation binding the contract event 0x96c49d03ef64591194500229a104cd087b2d45c68234c96444c3a2a6abb0bb97.
//
// Solidity: event VaultAdvancedToWithdrawals(address indexed arbitrator, uint32 maturedAt)
func (_IDurationVaultStrategy *IDurationVaultStrategyFilterer) WatchVaultAdvancedToWithdrawals(opts *bind.WatchOpts, sink chan<- *IDurationVaultStrategyVaultAdvancedToWithdrawals, arbitrator []common.Address) (event.Subscription, error) {

	var arbitratorRule []interface{}
	for _, arbitratorItem := range arbitrator {
		arbitratorRule = append(arbitratorRule, arbitratorItem)
	}

	logs, sub, err := _IDurationVaultStrategy.contract.WatchLogs(opts, "VaultAdvancedToWithdrawals", arbitratorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDurationVaultStrategyVaultAdvancedToWithdrawals)
				if err := _IDurationVaultStrategy.contract.UnpackLog(event, "VaultAdvancedToWithdrawals", log); err != nil {
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
func (_IDurationVaultStrategy *IDurationVaultStrategyFilterer) ParseVaultAdvancedToWithdrawals(log types.Log) (*IDurationVaultStrategyVaultAdvancedToWithdrawals, error) {
	event := new(IDurationVaultStrategyVaultAdvancedToWithdrawals)
	if err := _IDurationVaultStrategy.contract.UnpackLog(event, "VaultAdvancedToWithdrawals", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDurationVaultStrategyVaultInitializedIterator is returned from FilterVaultInitialized and is used to iterate over the raw logs and unpacked data for VaultInitialized events raised by the IDurationVaultStrategy contract.
type IDurationVaultStrategyVaultInitializedIterator struct {
	Event *IDurationVaultStrategyVaultInitialized // Event containing the contract specifics and raw log

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
func (it *IDurationVaultStrategyVaultInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDurationVaultStrategyVaultInitialized)
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
		it.Event = new(IDurationVaultStrategyVaultInitialized)
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
func (it *IDurationVaultStrategyVaultInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDurationVaultStrategyVaultInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDurationVaultStrategyVaultInitialized represents a VaultInitialized event raised by the IDurationVaultStrategy contract.
type IDurationVaultStrategyVaultInitialized struct {
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
func (_IDurationVaultStrategy *IDurationVaultStrategyFilterer) FilterVaultInitialized(opts *bind.FilterOpts, vaultAdmin []common.Address, arbitrator []common.Address, underlyingToken []common.Address) (*IDurationVaultStrategyVaultInitializedIterator, error) {

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

	logs, sub, err := _IDurationVaultStrategy.contract.FilterLogs(opts, "VaultInitialized", vaultAdminRule, arbitratorRule, underlyingTokenRule)
	if err != nil {
		return nil, err
	}
	return &IDurationVaultStrategyVaultInitializedIterator{contract: _IDurationVaultStrategy.contract, event: "VaultInitialized", logs: logs, sub: sub}, nil
}

// WatchVaultInitialized is a free log subscription operation binding the contract event 0xbdbff63632f473bb2a7c6a4aafbc096b71fbda12e22c6b51643bfd64f13d2b9e.
//
// Solidity: event VaultInitialized(address indexed vaultAdmin, address indexed arbitrator, address indexed underlyingToken, uint32 duration, uint256 maxPerDeposit, uint256 stakeCap, string metadataURI)
func (_IDurationVaultStrategy *IDurationVaultStrategyFilterer) WatchVaultInitialized(opts *bind.WatchOpts, sink chan<- *IDurationVaultStrategyVaultInitialized, vaultAdmin []common.Address, arbitrator []common.Address, underlyingToken []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _IDurationVaultStrategy.contract.WatchLogs(opts, "VaultInitialized", vaultAdminRule, arbitratorRule, underlyingTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDurationVaultStrategyVaultInitialized)
				if err := _IDurationVaultStrategy.contract.UnpackLog(event, "VaultInitialized", log); err != nil {
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
func (_IDurationVaultStrategy *IDurationVaultStrategyFilterer) ParseVaultInitialized(log types.Log) (*IDurationVaultStrategyVaultInitialized, error) {
	event := new(IDurationVaultStrategyVaultInitialized)
	if err := _IDurationVaultStrategy.contract.UnpackLog(event, "VaultInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDurationVaultStrategyVaultLockedIterator is returned from FilterVaultLocked and is used to iterate over the raw logs and unpacked data for VaultLocked events raised by the IDurationVaultStrategy contract.
type IDurationVaultStrategyVaultLockedIterator struct {
	Event *IDurationVaultStrategyVaultLocked // Event containing the contract specifics and raw log

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
func (it *IDurationVaultStrategyVaultLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDurationVaultStrategyVaultLocked)
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
		it.Event = new(IDurationVaultStrategyVaultLocked)
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
func (it *IDurationVaultStrategyVaultLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDurationVaultStrategyVaultLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDurationVaultStrategyVaultLocked represents a VaultLocked event raised by the IDurationVaultStrategy contract.
type IDurationVaultStrategyVaultLocked struct {
	LockedAt uint32
	UnlockAt uint32
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterVaultLocked is a free log retrieval operation binding the contract event 0x42cd6d7338516695d9c9ff8969dbdcf89ce22e3f2f76fda2fc11e973fe4860e4.
//
// Solidity: event VaultLocked(uint32 lockedAt, uint32 unlockAt)
func (_IDurationVaultStrategy *IDurationVaultStrategyFilterer) FilterVaultLocked(opts *bind.FilterOpts) (*IDurationVaultStrategyVaultLockedIterator, error) {

	logs, sub, err := _IDurationVaultStrategy.contract.FilterLogs(opts, "VaultLocked")
	if err != nil {
		return nil, err
	}
	return &IDurationVaultStrategyVaultLockedIterator{contract: _IDurationVaultStrategy.contract, event: "VaultLocked", logs: logs, sub: sub}, nil
}

// WatchVaultLocked is a free log subscription operation binding the contract event 0x42cd6d7338516695d9c9ff8969dbdcf89ce22e3f2f76fda2fc11e973fe4860e4.
//
// Solidity: event VaultLocked(uint32 lockedAt, uint32 unlockAt)
func (_IDurationVaultStrategy *IDurationVaultStrategyFilterer) WatchVaultLocked(opts *bind.WatchOpts, sink chan<- *IDurationVaultStrategyVaultLocked) (event.Subscription, error) {

	logs, sub, err := _IDurationVaultStrategy.contract.WatchLogs(opts, "VaultLocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDurationVaultStrategyVaultLocked)
				if err := _IDurationVaultStrategy.contract.UnpackLog(event, "VaultLocked", log); err != nil {
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
func (_IDurationVaultStrategy *IDurationVaultStrategyFilterer) ParseVaultLocked(log types.Log) (*IDurationVaultStrategyVaultLocked, error) {
	event := new(IDurationVaultStrategyVaultLocked)
	if err := _IDurationVaultStrategy.contract.UnpackLog(event, "VaultLocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDurationVaultStrategyVaultMaturedIterator is returned from FilterVaultMatured and is used to iterate over the raw logs and unpacked data for VaultMatured events raised by the IDurationVaultStrategy contract.
type IDurationVaultStrategyVaultMaturedIterator struct {
	Event *IDurationVaultStrategyVaultMatured // Event containing the contract specifics and raw log

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
func (it *IDurationVaultStrategyVaultMaturedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDurationVaultStrategyVaultMatured)
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
		it.Event = new(IDurationVaultStrategyVaultMatured)
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
func (it *IDurationVaultStrategyVaultMaturedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDurationVaultStrategyVaultMaturedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDurationVaultStrategyVaultMatured represents a VaultMatured event raised by the IDurationVaultStrategy contract.
type IDurationVaultStrategyVaultMatured struct {
	MaturedAt uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterVaultMatured is a free log retrieval operation binding the contract event 0xff979382d3040b1602e0a02f0f2a454b2250aa36e891d2da0ceb95d70d11a8f2.
//
// Solidity: event VaultMatured(uint32 maturedAt)
func (_IDurationVaultStrategy *IDurationVaultStrategyFilterer) FilterVaultMatured(opts *bind.FilterOpts) (*IDurationVaultStrategyVaultMaturedIterator, error) {

	logs, sub, err := _IDurationVaultStrategy.contract.FilterLogs(opts, "VaultMatured")
	if err != nil {
		return nil, err
	}
	return &IDurationVaultStrategyVaultMaturedIterator{contract: _IDurationVaultStrategy.contract, event: "VaultMatured", logs: logs, sub: sub}, nil
}

// WatchVaultMatured is a free log subscription operation binding the contract event 0xff979382d3040b1602e0a02f0f2a454b2250aa36e891d2da0ceb95d70d11a8f2.
//
// Solidity: event VaultMatured(uint32 maturedAt)
func (_IDurationVaultStrategy *IDurationVaultStrategyFilterer) WatchVaultMatured(opts *bind.WatchOpts, sink chan<- *IDurationVaultStrategyVaultMatured) (event.Subscription, error) {

	logs, sub, err := _IDurationVaultStrategy.contract.WatchLogs(opts, "VaultMatured")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDurationVaultStrategyVaultMatured)
				if err := _IDurationVaultStrategy.contract.UnpackLog(event, "VaultMatured", log); err != nil {
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
func (_IDurationVaultStrategy *IDurationVaultStrategyFilterer) ParseVaultMatured(log types.Log) (*IDurationVaultStrategyVaultMatured, error) {
	event := new(IDurationVaultStrategyVaultMatured)
	if err := _IDurationVaultStrategy.contract.UnpackLog(event, "VaultMatured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
