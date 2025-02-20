// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package StrategyManagerStorage

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

// StrategyManagerStorageMetaData contains all meta data concerning the StrategyManagerStorage contract.
var StrategyManagerStorageMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"DEFAULT_BURN_ADDRESS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEPOSIT_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addStrategiesToDepositWhitelist\",\"inputs\":[{\"name\":\"strategiesToWhitelist\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burnShares\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"calculateStrategyDepositDigestHash\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"depositIntoStrategy\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"depositShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositIntoStrategyWithSignature\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"depositShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getBurnableShares\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDeposits\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStakerStrategyList\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStrategiesWithBurnableShares\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"increaseBurnableShares\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"addedSharesToBurn\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialStrategyWhitelister\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"nonces\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeDepositShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"depositSharesToRemove\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeStrategiesFromDepositWhitelist\",\"inputs\":[{\"name\":\"strategiesToRemoveFromWhitelist\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStrategyWhitelister\",\"inputs\":[{\"name\":\"newStrategyWhitelister\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakerDepositShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakerStrategyList\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"strategies\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakerStrategyListLength\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"strategyIsWhitelistedForDeposit\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"whitelisted\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"strategyWhitelister\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawSharesAsTokens\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"BurnableSharesDecreased\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BurnableSharesIncreased\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposit\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyAddedToDepositWhitelist\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyRemovedFromDepositWhitelist\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyWhitelisterChanged\",\"inputs\":[{\"name\":\"previousAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"MaxStrategiesExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyDelegationManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyStrategyWhitelister\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SharesAmountTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SharesAmountZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StakerAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategyNotFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategyNotWhitelisted\",\"inputs\":[]}]",
}

// StrategyManagerStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use StrategyManagerStorageMetaData.ABI instead.
var StrategyManagerStorageABI = StrategyManagerStorageMetaData.ABI

// StrategyManagerStorage is an auto generated Go binding around an Ethereum contract.
type StrategyManagerStorage struct {
	StrategyManagerStorageCaller     // Read-only binding to the contract
	StrategyManagerStorageTransactor // Write-only binding to the contract
	StrategyManagerStorageFilterer   // Log filterer for contract events
}

// StrategyManagerStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type StrategyManagerStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrategyManagerStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StrategyManagerStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrategyManagerStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StrategyManagerStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrategyManagerStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StrategyManagerStorageSession struct {
	Contract     *StrategyManagerStorage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// StrategyManagerStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StrategyManagerStorageCallerSession struct {
	Contract *StrategyManagerStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// StrategyManagerStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StrategyManagerStorageTransactorSession struct {
	Contract     *StrategyManagerStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// StrategyManagerStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type StrategyManagerStorageRaw struct {
	Contract *StrategyManagerStorage // Generic contract binding to access the raw methods on
}

// StrategyManagerStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StrategyManagerStorageCallerRaw struct {
	Contract *StrategyManagerStorageCaller // Generic read-only contract binding to access the raw methods on
}

// StrategyManagerStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StrategyManagerStorageTransactorRaw struct {
	Contract *StrategyManagerStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStrategyManagerStorage creates a new instance of StrategyManagerStorage, bound to a specific deployed contract.
func NewStrategyManagerStorage(address common.Address, backend bind.ContractBackend) (*StrategyManagerStorage, error) {
	contract, err := bindStrategyManagerStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StrategyManagerStorage{StrategyManagerStorageCaller: StrategyManagerStorageCaller{contract: contract}, StrategyManagerStorageTransactor: StrategyManagerStorageTransactor{contract: contract}, StrategyManagerStorageFilterer: StrategyManagerStorageFilterer{contract: contract}}, nil
}

// NewStrategyManagerStorageCaller creates a new read-only instance of StrategyManagerStorage, bound to a specific deployed contract.
func NewStrategyManagerStorageCaller(address common.Address, caller bind.ContractCaller) (*StrategyManagerStorageCaller, error) {
	contract, err := bindStrategyManagerStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StrategyManagerStorageCaller{contract: contract}, nil
}

// NewStrategyManagerStorageTransactor creates a new write-only instance of StrategyManagerStorage, bound to a specific deployed contract.
func NewStrategyManagerStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*StrategyManagerStorageTransactor, error) {
	contract, err := bindStrategyManagerStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StrategyManagerStorageTransactor{contract: contract}, nil
}

// NewStrategyManagerStorageFilterer creates a new log filterer instance of StrategyManagerStorage, bound to a specific deployed contract.
func NewStrategyManagerStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*StrategyManagerStorageFilterer, error) {
	contract, err := bindStrategyManagerStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StrategyManagerStorageFilterer{contract: contract}, nil
}

// bindStrategyManagerStorage binds a generic wrapper to an already deployed contract.
func bindStrategyManagerStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StrategyManagerStorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StrategyManagerStorage *StrategyManagerStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StrategyManagerStorage.Contract.StrategyManagerStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StrategyManagerStorage *StrategyManagerStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrategyManagerStorage.Contract.StrategyManagerStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StrategyManagerStorage *StrategyManagerStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StrategyManagerStorage.Contract.StrategyManagerStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StrategyManagerStorage *StrategyManagerStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StrategyManagerStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StrategyManagerStorage *StrategyManagerStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrategyManagerStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StrategyManagerStorage *StrategyManagerStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StrategyManagerStorage.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTBURNADDRESS is a free data retrieval call binding the contract method 0xf3b4a000.
//
// Solidity: function DEFAULT_BURN_ADDRESS() view returns(address)
func (_StrategyManagerStorage *StrategyManagerStorageCaller) DEFAULTBURNADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyManagerStorage.contract.Call(opts, &out, "DEFAULT_BURN_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DEFAULTBURNADDRESS is a free data retrieval call binding the contract method 0xf3b4a000.
//
// Solidity: function DEFAULT_BURN_ADDRESS() view returns(address)
func (_StrategyManagerStorage *StrategyManagerStorageSession) DEFAULTBURNADDRESS() (common.Address, error) {
	return _StrategyManagerStorage.Contract.DEFAULTBURNADDRESS(&_StrategyManagerStorage.CallOpts)
}

// DEFAULTBURNADDRESS is a free data retrieval call binding the contract method 0xf3b4a000.
//
// Solidity: function DEFAULT_BURN_ADDRESS() view returns(address)
func (_StrategyManagerStorage *StrategyManagerStorageCallerSession) DEFAULTBURNADDRESS() (common.Address, error) {
	return _StrategyManagerStorage.Contract.DEFAULTBURNADDRESS(&_StrategyManagerStorage.CallOpts)
}

// DEPOSITTYPEHASH is a free data retrieval call binding the contract method 0x48825e94.
//
// Solidity: function DEPOSIT_TYPEHASH() view returns(bytes32)
func (_StrategyManagerStorage *StrategyManagerStorageCaller) DEPOSITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StrategyManagerStorage.contract.Call(opts, &out, "DEPOSIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEPOSITTYPEHASH is a free data retrieval call binding the contract method 0x48825e94.
//
// Solidity: function DEPOSIT_TYPEHASH() view returns(bytes32)
func (_StrategyManagerStorage *StrategyManagerStorageSession) DEPOSITTYPEHASH() ([32]byte, error) {
	return _StrategyManagerStorage.Contract.DEPOSITTYPEHASH(&_StrategyManagerStorage.CallOpts)
}

// DEPOSITTYPEHASH is a free data retrieval call binding the contract method 0x48825e94.
//
// Solidity: function DEPOSIT_TYPEHASH() view returns(bytes32)
func (_StrategyManagerStorage *StrategyManagerStorageCallerSession) DEPOSITTYPEHASH() ([32]byte, error) {
	return _StrategyManagerStorage.Contract.DEPOSITTYPEHASH(&_StrategyManagerStorage.CallOpts)
}

// CalculateStrategyDepositDigestHash is a free data retrieval call binding the contract method 0x9ac01d61.
//
// Solidity: function calculateStrategyDepositDigestHash(address staker, address strategy, address token, uint256 amount, uint256 nonce, uint256 expiry) view returns(bytes32)
func (_StrategyManagerStorage *StrategyManagerStorageCaller) CalculateStrategyDepositDigestHash(opts *bind.CallOpts, staker common.Address, strategy common.Address, token common.Address, amount *big.Int, nonce *big.Int, expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _StrategyManagerStorage.contract.Call(opts, &out, "calculateStrategyDepositDigestHash", staker, strategy, token, amount, nonce, expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateStrategyDepositDigestHash is a free data retrieval call binding the contract method 0x9ac01d61.
//
// Solidity: function calculateStrategyDepositDigestHash(address staker, address strategy, address token, uint256 amount, uint256 nonce, uint256 expiry) view returns(bytes32)
func (_StrategyManagerStorage *StrategyManagerStorageSession) CalculateStrategyDepositDigestHash(staker common.Address, strategy common.Address, token common.Address, amount *big.Int, nonce *big.Int, expiry *big.Int) ([32]byte, error) {
	return _StrategyManagerStorage.Contract.CalculateStrategyDepositDigestHash(&_StrategyManagerStorage.CallOpts, staker, strategy, token, amount, nonce, expiry)
}

// CalculateStrategyDepositDigestHash is a free data retrieval call binding the contract method 0x9ac01d61.
//
// Solidity: function calculateStrategyDepositDigestHash(address staker, address strategy, address token, uint256 amount, uint256 nonce, uint256 expiry) view returns(bytes32)
func (_StrategyManagerStorage *StrategyManagerStorageCallerSession) CalculateStrategyDepositDigestHash(staker common.Address, strategy common.Address, token common.Address, amount *big.Int, nonce *big.Int, expiry *big.Int) ([32]byte, error) {
	return _StrategyManagerStorage.Contract.CalculateStrategyDepositDigestHash(&_StrategyManagerStorage.CallOpts, staker, strategy, token, amount, nonce, expiry)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_StrategyManagerStorage *StrategyManagerStorageCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyManagerStorage.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_StrategyManagerStorage *StrategyManagerStorageSession) Delegation() (common.Address, error) {
	return _StrategyManagerStorage.Contract.Delegation(&_StrategyManagerStorage.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_StrategyManagerStorage *StrategyManagerStorageCallerSession) Delegation() (common.Address, error) {
	return _StrategyManagerStorage.Contract.Delegation(&_StrategyManagerStorage.CallOpts)
}

// GetBurnableShares is a free data retrieval call binding the contract method 0xfd980423.
//
// Solidity: function getBurnableShares(address strategy) view returns(uint256)
func (_StrategyManagerStorage *StrategyManagerStorageCaller) GetBurnableShares(opts *bind.CallOpts, strategy common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StrategyManagerStorage.contract.Call(opts, &out, "getBurnableShares", strategy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBurnableShares is a free data retrieval call binding the contract method 0xfd980423.
//
// Solidity: function getBurnableShares(address strategy) view returns(uint256)
func (_StrategyManagerStorage *StrategyManagerStorageSession) GetBurnableShares(strategy common.Address) (*big.Int, error) {
	return _StrategyManagerStorage.Contract.GetBurnableShares(&_StrategyManagerStorage.CallOpts, strategy)
}

// GetBurnableShares is a free data retrieval call binding the contract method 0xfd980423.
//
// Solidity: function getBurnableShares(address strategy) view returns(uint256)
func (_StrategyManagerStorage *StrategyManagerStorageCallerSession) GetBurnableShares(strategy common.Address) (*big.Int, error) {
	return _StrategyManagerStorage.Contract.GetBurnableShares(&_StrategyManagerStorage.CallOpts, strategy)
}

// GetDeposits is a free data retrieval call binding the contract method 0x94f649dd.
//
// Solidity: function getDeposits(address staker) view returns(address[], uint256[])
func (_StrategyManagerStorage *StrategyManagerStorageCaller) GetDeposits(opts *bind.CallOpts, staker common.Address) ([]common.Address, []*big.Int, error) {
	var out []interface{}
	err := _StrategyManagerStorage.contract.Call(opts, &out, "getDeposits", staker)

	if err != nil {
		return *new([]common.Address), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// GetDeposits is a free data retrieval call binding the contract method 0x94f649dd.
//
// Solidity: function getDeposits(address staker) view returns(address[], uint256[])
func (_StrategyManagerStorage *StrategyManagerStorageSession) GetDeposits(staker common.Address) ([]common.Address, []*big.Int, error) {
	return _StrategyManagerStorage.Contract.GetDeposits(&_StrategyManagerStorage.CallOpts, staker)
}

// GetDeposits is a free data retrieval call binding the contract method 0x94f649dd.
//
// Solidity: function getDeposits(address staker) view returns(address[], uint256[])
func (_StrategyManagerStorage *StrategyManagerStorageCallerSession) GetDeposits(staker common.Address) ([]common.Address, []*big.Int, error) {
	return _StrategyManagerStorage.Contract.GetDeposits(&_StrategyManagerStorage.CallOpts, staker)
}

// GetStakerStrategyList is a free data retrieval call binding the contract method 0xde44acb6.
//
// Solidity: function getStakerStrategyList(address staker) view returns(address[])
func (_StrategyManagerStorage *StrategyManagerStorageCaller) GetStakerStrategyList(opts *bind.CallOpts, staker common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _StrategyManagerStorage.contract.Call(opts, &out, "getStakerStrategyList", staker)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetStakerStrategyList is a free data retrieval call binding the contract method 0xde44acb6.
//
// Solidity: function getStakerStrategyList(address staker) view returns(address[])
func (_StrategyManagerStorage *StrategyManagerStorageSession) GetStakerStrategyList(staker common.Address) ([]common.Address, error) {
	return _StrategyManagerStorage.Contract.GetStakerStrategyList(&_StrategyManagerStorage.CallOpts, staker)
}

// GetStakerStrategyList is a free data retrieval call binding the contract method 0xde44acb6.
//
// Solidity: function getStakerStrategyList(address staker) view returns(address[])
func (_StrategyManagerStorage *StrategyManagerStorageCallerSession) GetStakerStrategyList(staker common.Address) ([]common.Address, error) {
	return _StrategyManagerStorage.Contract.GetStakerStrategyList(&_StrategyManagerStorage.CallOpts, staker)
}

// GetStrategiesWithBurnableShares is a free data retrieval call binding the contract method 0x36a8c500.
//
// Solidity: function getStrategiesWithBurnableShares() view returns(address[], uint256[])
func (_StrategyManagerStorage *StrategyManagerStorageCaller) GetStrategiesWithBurnableShares(opts *bind.CallOpts) ([]common.Address, []*big.Int, error) {
	var out []interface{}
	err := _StrategyManagerStorage.contract.Call(opts, &out, "getStrategiesWithBurnableShares")

	if err != nil {
		return *new([]common.Address), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// GetStrategiesWithBurnableShares is a free data retrieval call binding the contract method 0x36a8c500.
//
// Solidity: function getStrategiesWithBurnableShares() view returns(address[], uint256[])
func (_StrategyManagerStorage *StrategyManagerStorageSession) GetStrategiesWithBurnableShares() ([]common.Address, []*big.Int, error) {
	return _StrategyManagerStorage.Contract.GetStrategiesWithBurnableShares(&_StrategyManagerStorage.CallOpts)
}

// GetStrategiesWithBurnableShares is a free data retrieval call binding the contract method 0x36a8c500.
//
// Solidity: function getStrategiesWithBurnableShares() view returns(address[], uint256[])
func (_StrategyManagerStorage *StrategyManagerStorageCallerSession) GetStrategiesWithBurnableShares() ([]common.Address, []*big.Int, error) {
	return _StrategyManagerStorage.Contract.GetStrategiesWithBurnableShares(&_StrategyManagerStorage.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address signer) view returns(uint256 nonce)
func (_StrategyManagerStorage *StrategyManagerStorageCaller) Nonces(opts *bind.CallOpts, signer common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StrategyManagerStorage.contract.Call(opts, &out, "nonces", signer)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address signer) view returns(uint256 nonce)
func (_StrategyManagerStorage *StrategyManagerStorageSession) Nonces(signer common.Address) (*big.Int, error) {
	return _StrategyManagerStorage.Contract.Nonces(&_StrategyManagerStorage.CallOpts, signer)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address signer) view returns(uint256 nonce)
func (_StrategyManagerStorage *StrategyManagerStorageCallerSession) Nonces(signer common.Address) (*big.Int, error) {
	return _StrategyManagerStorage.Contract.Nonces(&_StrategyManagerStorage.CallOpts, signer)
}

// StakerDepositShares is a free data retrieval call binding the contract method 0xfe243a17.
//
// Solidity: function stakerDepositShares(address staker, address strategy) view returns(uint256 shares)
func (_StrategyManagerStorage *StrategyManagerStorageCaller) StakerDepositShares(opts *bind.CallOpts, staker common.Address, strategy common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StrategyManagerStorage.contract.Call(opts, &out, "stakerDepositShares", staker, strategy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerDepositShares is a free data retrieval call binding the contract method 0xfe243a17.
//
// Solidity: function stakerDepositShares(address staker, address strategy) view returns(uint256 shares)
func (_StrategyManagerStorage *StrategyManagerStorageSession) StakerDepositShares(staker common.Address, strategy common.Address) (*big.Int, error) {
	return _StrategyManagerStorage.Contract.StakerDepositShares(&_StrategyManagerStorage.CallOpts, staker, strategy)
}

// StakerDepositShares is a free data retrieval call binding the contract method 0xfe243a17.
//
// Solidity: function stakerDepositShares(address staker, address strategy) view returns(uint256 shares)
func (_StrategyManagerStorage *StrategyManagerStorageCallerSession) StakerDepositShares(staker common.Address, strategy common.Address) (*big.Int, error) {
	return _StrategyManagerStorage.Contract.StakerDepositShares(&_StrategyManagerStorage.CallOpts, staker, strategy)
}

// StakerStrategyList is a free data retrieval call binding the contract method 0xcbc2bd62.
//
// Solidity: function stakerStrategyList(address staker, uint256 ) view returns(address strategies)
func (_StrategyManagerStorage *StrategyManagerStorageCaller) StakerStrategyList(opts *bind.CallOpts, staker common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _StrategyManagerStorage.contract.Call(opts, &out, "stakerStrategyList", staker, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakerStrategyList is a free data retrieval call binding the contract method 0xcbc2bd62.
//
// Solidity: function stakerStrategyList(address staker, uint256 ) view returns(address strategies)
func (_StrategyManagerStorage *StrategyManagerStorageSession) StakerStrategyList(staker common.Address, arg1 *big.Int) (common.Address, error) {
	return _StrategyManagerStorage.Contract.StakerStrategyList(&_StrategyManagerStorage.CallOpts, staker, arg1)
}

// StakerStrategyList is a free data retrieval call binding the contract method 0xcbc2bd62.
//
// Solidity: function stakerStrategyList(address staker, uint256 ) view returns(address strategies)
func (_StrategyManagerStorage *StrategyManagerStorageCallerSession) StakerStrategyList(staker common.Address, arg1 *big.Int) (common.Address, error) {
	return _StrategyManagerStorage.Contract.StakerStrategyList(&_StrategyManagerStorage.CallOpts, staker, arg1)
}

// StakerStrategyListLength is a free data retrieval call binding the contract method 0x8b8aac3c.
//
// Solidity: function stakerStrategyListLength(address staker) view returns(uint256)
func (_StrategyManagerStorage *StrategyManagerStorageCaller) StakerStrategyListLength(opts *bind.CallOpts, staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StrategyManagerStorage.contract.Call(opts, &out, "stakerStrategyListLength", staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerStrategyListLength is a free data retrieval call binding the contract method 0x8b8aac3c.
//
// Solidity: function stakerStrategyListLength(address staker) view returns(uint256)
func (_StrategyManagerStorage *StrategyManagerStorageSession) StakerStrategyListLength(staker common.Address) (*big.Int, error) {
	return _StrategyManagerStorage.Contract.StakerStrategyListLength(&_StrategyManagerStorage.CallOpts, staker)
}

// StakerStrategyListLength is a free data retrieval call binding the contract method 0x8b8aac3c.
//
// Solidity: function stakerStrategyListLength(address staker) view returns(uint256)
func (_StrategyManagerStorage *StrategyManagerStorageCallerSession) StakerStrategyListLength(staker common.Address) (*big.Int, error) {
	return _StrategyManagerStorage.Contract.StakerStrategyListLength(&_StrategyManagerStorage.CallOpts, staker)
}

// StrategyIsWhitelistedForDeposit is a free data retrieval call binding the contract method 0x663c1de4.
//
// Solidity: function strategyIsWhitelistedForDeposit(address strategy) view returns(bool whitelisted)
func (_StrategyManagerStorage *StrategyManagerStorageCaller) StrategyIsWhitelistedForDeposit(opts *bind.CallOpts, strategy common.Address) (bool, error) {
	var out []interface{}
	err := _StrategyManagerStorage.contract.Call(opts, &out, "strategyIsWhitelistedForDeposit", strategy)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StrategyIsWhitelistedForDeposit is a free data retrieval call binding the contract method 0x663c1de4.
//
// Solidity: function strategyIsWhitelistedForDeposit(address strategy) view returns(bool whitelisted)
func (_StrategyManagerStorage *StrategyManagerStorageSession) StrategyIsWhitelistedForDeposit(strategy common.Address) (bool, error) {
	return _StrategyManagerStorage.Contract.StrategyIsWhitelistedForDeposit(&_StrategyManagerStorage.CallOpts, strategy)
}

// StrategyIsWhitelistedForDeposit is a free data retrieval call binding the contract method 0x663c1de4.
//
// Solidity: function strategyIsWhitelistedForDeposit(address strategy) view returns(bool whitelisted)
func (_StrategyManagerStorage *StrategyManagerStorageCallerSession) StrategyIsWhitelistedForDeposit(strategy common.Address) (bool, error) {
	return _StrategyManagerStorage.Contract.StrategyIsWhitelistedForDeposit(&_StrategyManagerStorage.CallOpts, strategy)
}

// StrategyWhitelister is a free data retrieval call binding the contract method 0x967fc0d2.
//
// Solidity: function strategyWhitelister() view returns(address)
func (_StrategyManagerStorage *StrategyManagerStorageCaller) StrategyWhitelister(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyManagerStorage.contract.Call(opts, &out, "strategyWhitelister")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StrategyWhitelister is a free data retrieval call binding the contract method 0x967fc0d2.
//
// Solidity: function strategyWhitelister() view returns(address)
func (_StrategyManagerStorage *StrategyManagerStorageSession) StrategyWhitelister() (common.Address, error) {
	return _StrategyManagerStorage.Contract.StrategyWhitelister(&_StrategyManagerStorage.CallOpts)
}

// StrategyWhitelister is a free data retrieval call binding the contract method 0x967fc0d2.
//
// Solidity: function strategyWhitelister() view returns(address)
func (_StrategyManagerStorage *StrategyManagerStorageCallerSession) StrategyWhitelister() (common.Address, error) {
	return _StrategyManagerStorage.Contract.StrategyWhitelister(&_StrategyManagerStorage.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_StrategyManagerStorage *StrategyManagerStorageCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StrategyManagerStorage.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_StrategyManagerStorage *StrategyManagerStorageSession) Version() (string, error) {
	return _StrategyManagerStorage.Contract.Version(&_StrategyManagerStorage.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_StrategyManagerStorage *StrategyManagerStorageCallerSession) Version() (string, error) {
	return _StrategyManagerStorage.Contract.Version(&_StrategyManagerStorage.CallOpts)
}

// AddShares is a paid mutator transaction binding the contract method 0x50ff7225.
//
// Solidity: function addShares(address staker, address strategy, uint256 shares) returns(uint256, uint256)
func (_StrategyManagerStorage *StrategyManagerStorageTransactor) AddShares(opts *bind.TransactOpts, staker common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error) {
	return _StrategyManagerStorage.contract.Transact(opts, "addShares", staker, strategy, shares)
}

// AddShares is a paid mutator transaction binding the contract method 0x50ff7225.
//
// Solidity: function addShares(address staker, address strategy, uint256 shares) returns(uint256, uint256)
func (_StrategyManagerStorage *StrategyManagerStorageSession) AddShares(staker common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error) {
	return _StrategyManagerStorage.Contract.AddShares(&_StrategyManagerStorage.TransactOpts, staker, strategy, shares)
}

// AddShares is a paid mutator transaction binding the contract method 0x50ff7225.
//
// Solidity: function addShares(address staker, address strategy, uint256 shares) returns(uint256, uint256)
func (_StrategyManagerStorage *StrategyManagerStorageTransactorSession) AddShares(staker common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error) {
	return _StrategyManagerStorage.Contract.AddShares(&_StrategyManagerStorage.TransactOpts, staker, strategy, shares)
}

// AddStrategiesToDepositWhitelist is a paid mutator transaction binding the contract method 0x5de08ff2.
//
// Solidity: function addStrategiesToDepositWhitelist(address[] strategiesToWhitelist) returns()
func (_StrategyManagerStorage *StrategyManagerStorageTransactor) AddStrategiesToDepositWhitelist(opts *bind.TransactOpts, strategiesToWhitelist []common.Address) (*types.Transaction, error) {
	return _StrategyManagerStorage.contract.Transact(opts, "addStrategiesToDepositWhitelist", strategiesToWhitelist)
}

// AddStrategiesToDepositWhitelist is a paid mutator transaction binding the contract method 0x5de08ff2.
//
// Solidity: function addStrategiesToDepositWhitelist(address[] strategiesToWhitelist) returns()
func (_StrategyManagerStorage *StrategyManagerStorageSession) AddStrategiesToDepositWhitelist(strategiesToWhitelist []common.Address) (*types.Transaction, error) {
	return _StrategyManagerStorage.Contract.AddStrategiesToDepositWhitelist(&_StrategyManagerStorage.TransactOpts, strategiesToWhitelist)
}

// AddStrategiesToDepositWhitelist is a paid mutator transaction binding the contract method 0x5de08ff2.
//
// Solidity: function addStrategiesToDepositWhitelist(address[] strategiesToWhitelist) returns()
func (_StrategyManagerStorage *StrategyManagerStorageTransactorSession) AddStrategiesToDepositWhitelist(strategiesToWhitelist []common.Address) (*types.Transaction, error) {
	return _StrategyManagerStorage.Contract.AddStrategiesToDepositWhitelist(&_StrategyManagerStorage.TransactOpts, strategiesToWhitelist)
}

// BurnShares is a paid mutator transaction binding the contract method 0x4b6d5d6e.
//
// Solidity: function burnShares(address strategy) returns()
func (_StrategyManagerStorage *StrategyManagerStorageTransactor) BurnShares(opts *bind.TransactOpts, strategy common.Address) (*types.Transaction, error) {
	return _StrategyManagerStorage.contract.Transact(opts, "burnShares", strategy)
}

// BurnShares is a paid mutator transaction binding the contract method 0x4b6d5d6e.
//
// Solidity: function burnShares(address strategy) returns()
func (_StrategyManagerStorage *StrategyManagerStorageSession) BurnShares(strategy common.Address) (*types.Transaction, error) {
	return _StrategyManagerStorage.Contract.BurnShares(&_StrategyManagerStorage.TransactOpts, strategy)
}

// BurnShares is a paid mutator transaction binding the contract method 0x4b6d5d6e.
//
// Solidity: function burnShares(address strategy) returns()
func (_StrategyManagerStorage *StrategyManagerStorageTransactorSession) BurnShares(strategy common.Address) (*types.Transaction, error) {
	return _StrategyManagerStorage.Contract.BurnShares(&_StrategyManagerStorage.TransactOpts, strategy)
}

// DepositIntoStrategy is a paid mutator transaction binding the contract method 0xe7a050aa.
//
// Solidity: function depositIntoStrategy(address strategy, address token, uint256 amount) returns(uint256 depositShares)
func (_StrategyManagerStorage *StrategyManagerStorageTransactor) DepositIntoStrategy(opts *bind.TransactOpts, strategy common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StrategyManagerStorage.contract.Transact(opts, "depositIntoStrategy", strategy, token, amount)
}

// DepositIntoStrategy is a paid mutator transaction binding the contract method 0xe7a050aa.
//
// Solidity: function depositIntoStrategy(address strategy, address token, uint256 amount) returns(uint256 depositShares)
func (_StrategyManagerStorage *StrategyManagerStorageSession) DepositIntoStrategy(strategy common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StrategyManagerStorage.Contract.DepositIntoStrategy(&_StrategyManagerStorage.TransactOpts, strategy, token, amount)
}

// DepositIntoStrategy is a paid mutator transaction binding the contract method 0xe7a050aa.
//
// Solidity: function depositIntoStrategy(address strategy, address token, uint256 amount) returns(uint256 depositShares)
func (_StrategyManagerStorage *StrategyManagerStorageTransactorSession) DepositIntoStrategy(strategy common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StrategyManagerStorage.Contract.DepositIntoStrategy(&_StrategyManagerStorage.TransactOpts, strategy, token, amount)
}

// DepositIntoStrategyWithSignature is a paid mutator transaction binding the contract method 0x32e89ace.
//
// Solidity: function depositIntoStrategyWithSignature(address strategy, address token, uint256 amount, address staker, uint256 expiry, bytes signature) returns(uint256 depositShares)
func (_StrategyManagerStorage *StrategyManagerStorageTransactor) DepositIntoStrategyWithSignature(opts *bind.TransactOpts, strategy common.Address, token common.Address, amount *big.Int, staker common.Address, expiry *big.Int, signature []byte) (*types.Transaction, error) {
	return _StrategyManagerStorage.contract.Transact(opts, "depositIntoStrategyWithSignature", strategy, token, amount, staker, expiry, signature)
}

// DepositIntoStrategyWithSignature is a paid mutator transaction binding the contract method 0x32e89ace.
//
// Solidity: function depositIntoStrategyWithSignature(address strategy, address token, uint256 amount, address staker, uint256 expiry, bytes signature) returns(uint256 depositShares)
func (_StrategyManagerStorage *StrategyManagerStorageSession) DepositIntoStrategyWithSignature(strategy common.Address, token common.Address, amount *big.Int, staker common.Address, expiry *big.Int, signature []byte) (*types.Transaction, error) {
	return _StrategyManagerStorage.Contract.DepositIntoStrategyWithSignature(&_StrategyManagerStorage.TransactOpts, strategy, token, amount, staker, expiry, signature)
}

// DepositIntoStrategyWithSignature is a paid mutator transaction binding the contract method 0x32e89ace.
//
// Solidity: function depositIntoStrategyWithSignature(address strategy, address token, uint256 amount, address staker, uint256 expiry, bytes signature) returns(uint256 depositShares)
func (_StrategyManagerStorage *StrategyManagerStorageTransactorSession) DepositIntoStrategyWithSignature(strategy common.Address, token common.Address, amount *big.Int, staker common.Address, expiry *big.Int, signature []byte) (*types.Transaction, error) {
	return _StrategyManagerStorage.Contract.DepositIntoStrategyWithSignature(&_StrategyManagerStorage.TransactOpts, strategy, token, amount, staker, expiry, signature)
}

// IncreaseBurnableShares is a paid mutator transaction binding the contract method 0xdebe1eab.
//
// Solidity: function increaseBurnableShares(address strategy, uint256 addedSharesToBurn) returns()
func (_StrategyManagerStorage *StrategyManagerStorageTransactor) IncreaseBurnableShares(opts *bind.TransactOpts, strategy common.Address, addedSharesToBurn *big.Int) (*types.Transaction, error) {
	return _StrategyManagerStorage.contract.Transact(opts, "increaseBurnableShares", strategy, addedSharesToBurn)
}

// IncreaseBurnableShares is a paid mutator transaction binding the contract method 0xdebe1eab.
//
// Solidity: function increaseBurnableShares(address strategy, uint256 addedSharesToBurn) returns()
func (_StrategyManagerStorage *StrategyManagerStorageSession) IncreaseBurnableShares(strategy common.Address, addedSharesToBurn *big.Int) (*types.Transaction, error) {
	return _StrategyManagerStorage.Contract.IncreaseBurnableShares(&_StrategyManagerStorage.TransactOpts, strategy, addedSharesToBurn)
}

// IncreaseBurnableShares is a paid mutator transaction binding the contract method 0xdebe1eab.
//
// Solidity: function increaseBurnableShares(address strategy, uint256 addedSharesToBurn) returns()
func (_StrategyManagerStorage *StrategyManagerStorageTransactorSession) IncreaseBurnableShares(strategy common.Address, addedSharesToBurn *big.Int) (*types.Transaction, error) {
	return _StrategyManagerStorage.Contract.IncreaseBurnableShares(&_StrategyManagerStorage.TransactOpts, strategy, addedSharesToBurn)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address initialOwner, address initialStrategyWhitelister, uint256 initialPausedStatus) returns()
func (_StrategyManagerStorage *StrategyManagerStorageTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, initialStrategyWhitelister common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _StrategyManagerStorage.contract.Transact(opts, "initialize", initialOwner, initialStrategyWhitelister, initialPausedStatus)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address initialOwner, address initialStrategyWhitelister, uint256 initialPausedStatus) returns()
func (_StrategyManagerStorage *StrategyManagerStorageSession) Initialize(initialOwner common.Address, initialStrategyWhitelister common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _StrategyManagerStorage.Contract.Initialize(&_StrategyManagerStorage.TransactOpts, initialOwner, initialStrategyWhitelister, initialPausedStatus)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address initialOwner, address initialStrategyWhitelister, uint256 initialPausedStatus) returns()
func (_StrategyManagerStorage *StrategyManagerStorageTransactorSession) Initialize(initialOwner common.Address, initialStrategyWhitelister common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _StrategyManagerStorage.Contract.Initialize(&_StrategyManagerStorage.TransactOpts, initialOwner, initialStrategyWhitelister, initialPausedStatus)
}

// RemoveDepositShares is a paid mutator transaction binding the contract method 0x724af423.
//
// Solidity: function removeDepositShares(address staker, address strategy, uint256 depositSharesToRemove) returns(uint256)
func (_StrategyManagerStorage *StrategyManagerStorageTransactor) RemoveDepositShares(opts *bind.TransactOpts, staker common.Address, strategy common.Address, depositSharesToRemove *big.Int) (*types.Transaction, error) {
	return _StrategyManagerStorage.contract.Transact(opts, "removeDepositShares", staker, strategy, depositSharesToRemove)
}

// RemoveDepositShares is a paid mutator transaction binding the contract method 0x724af423.
//
// Solidity: function removeDepositShares(address staker, address strategy, uint256 depositSharesToRemove) returns(uint256)
func (_StrategyManagerStorage *StrategyManagerStorageSession) RemoveDepositShares(staker common.Address, strategy common.Address, depositSharesToRemove *big.Int) (*types.Transaction, error) {
	return _StrategyManagerStorage.Contract.RemoveDepositShares(&_StrategyManagerStorage.TransactOpts, staker, strategy, depositSharesToRemove)
}

// RemoveDepositShares is a paid mutator transaction binding the contract method 0x724af423.
//
// Solidity: function removeDepositShares(address staker, address strategy, uint256 depositSharesToRemove) returns(uint256)
func (_StrategyManagerStorage *StrategyManagerStorageTransactorSession) RemoveDepositShares(staker common.Address, strategy common.Address, depositSharesToRemove *big.Int) (*types.Transaction, error) {
	return _StrategyManagerStorage.Contract.RemoveDepositShares(&_StrategyManagerStorage.TransactOpts, staker, strategy, depositSharesToRemove)
}

// RemoveStrategiesFromDepositWhitelist is a paid mutator transaction binding the contract method 0xb5d8b5b8.
//
// Solidity: function removeStrategiesFromDepositWhitelist(address[] strategiesToRemoveFromWhitelist) returns()
func (_StrategyManagerStorage *StrategyManagerStorageTransactor) RemoveStrategiesFromDepositWhitelist(opts *bind.TransactOpts, strategiesToRemoveFromWhitelist []common.Address) (*types.Transaction, error) {
	return _StrategyManagerStorage.contract.Transact(opts, "removeStrategiesFromDepositWhitelist", strategiesToRemoveFromWhitelist)
}

// RemoveStrategiesFromDepositWhitelist is a paid mutator transaction binding the contract method 0xb5d8b5b8.
//
// Solidity: function removeStrategiesFromDepositWhitelist(address[] strategiesToRemoveFromWhitelist) returns()
func (_StrategyManagerStorage *StrategyManagerStorageSession) RemoveStrategiesFromDepositWhitelist(strategiesToRemoveFromWhitelist []common.Address) (*types.Transaction, error) {
	return _StrategyManagerStorage.Contract.RemoveStrategiesFromDepositWhitelist(&_StrategyManagerStorage.TransactOpts, strategiesToRemoveFromWhitelist)
}

// RemoveStrategiesFromDepositWhitelist is a paid mutator transaction binding the contract method 0xb5d8b5b8.
//
// Solidity: function removeStrategiesFromDepositWhitelist(address[] strategiesToRemoveFromWhitelist) returns()
func (_StrategyManagerStorage *StrategyManagerStorageTransactorSession) RemoveStrategiesFromDepositWhitelist(strategiesToRemoveFromWhitelist []common.Address) (*types.Transaction, error) {
	return _StrategyManagerStorage.Contract.RemoveStrategiesFromDepositWhitelist(&_StrategyManagerStorage.TransactOpts, strategiesToRemoveFromWhitelist)
}

// SetStrategyWhitelister is a paid mutator transaction binding the contract method 0xc6656702.
//
// Solidity: function setStrategyWhitelister(address newStrategyWhitelister) returns()
func (_StrategyManagerStorage *StrategyManagerStorageTransactor) SetStrategyWhitelister(opts *bind.TransactOpts, newStrategyWhitelister common.Address) (*types.Transaction, error) {
	return _StrategyManagerStorage.contract.Transact(opts, "setStrategyWhitelister", newStrategyWhitelister)
}

// SetStrategyWhitelister is a paid mutator transaction binding the contract method 0xc6656702.
//
// Solidity: function setStrategyWhitelister(address newStrategyWhitelister) returns()
func (_StrategyManagerStorage *StrategyManagerStorageSession) SetStrategyWhitelister(newStrategyWhitelister common.Address) (*types.Transaction, error) {
	return _StrategyManagerStorage.Contract.SetStrategyWhitelister(&_StrategyManagerStorage.TransactOpts, newStrategyWhitelister)
}

// SetStrategyWhitelister is a paid mutator transaction binding the contract method 0xc6656702.
//
// Solidity: function setStrategyWhitelister(address newStrategyWhitelister) returns()
func (_StrategyManagerStorage *StrategyManagerStorageTransactorSession) SetStrategyWhitelister(newStrategyWhitelister common.Address) (*types.Transaction, error) {
	return _StrategyManagerStorage.Contract.SetStrategyWhitelister(&_StrategyManagerStorage.TransactOpts, newStrategyWhitelister)
}

// WithdrawSharesAsTokens is a paid mutator transaction binding the contract method 0x2eae418c.
//
// Solidity: function withdrawSharesAsTokens(address staker, address strategy, address token, uint256 shares) returns()
func (_StrategyManagerStorage *StrategyManagerStorageTransactor) WithdrawSharesAsTokens(opts *bind.TransactOpts, staker common.Address, strategy common.Address, token common.Address, shares *big.Int) (*types.Transaction, error) {
	return _StrategyManagerStorage.contract.Transact(opts, "withdrawSharesAsTokens", staker, strategy, token, shares)
}

// WithdrawSharesAsTokens is a paid mutator transaction binding the contract method 0x2eae418c.
//
// Solidity: function withdrawSharesAsTokens(address staker, address strategy, address token, uint256 shares) returns()
func (_StrategyManagerStorage *StrategyManagerStorageSession) WithdrawSharesAsTokens(staker common.Address, strategy common.Address, token common.Address, shares *big.Int) (*types.Transaction, error) {
	return _StrategyManagerStorage.Contract.WithdrawSharesAsTokens(&_StrategyManagerStorage.TransactOpts, staker, strategy, token, shares)
}

// WithdrawSharesAsTokens is a paid mutator transaction binding the contract method 0x2eae418c.
//
// Solidity: function withdrawSharesAsTokens(address staker, address strategy, address token, uint256 shares) returns()
func (_StrategyManagerStorage *StrategyManagerStorageTransactorSession) WithdrawSharesAsTokens(staker common.Address, strategy common.Address, token common.Address, shares *big.Int) (*types.Transaction, error) {
	return _StrategyManagerStorage.Contract.WithdrawSharesAsTokens(&_StrategyManagerStorage.TransactOpts, staker, strategy, token, shares)
}

// StrategyManagerStorageBurnableSharesDecreasedIterator is returned from FilterBurnableSharesDecreased and is used to iterate over the raw logs and unpacked data for BurnableSharesDecreased events raised by the StrategyManagerStorage contract.
type StrategyManagerStorageBurnableSharesDecreasedIterator struct {
	Event *StrategyManagerStorageBurnableSharesDecreased // Event containing the contract specifics and raw log

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
func (it *StrategyManagerStorageBurnableSharesDecreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyManagerStorageBurnableSharesDecreased)
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
		it.Event = new(StrategyManagerStorageBurnableSharesDecreased)
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
func (it *StrategyManagerStorageBurnableSharesDecreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyManagerStorageBurnableSharesDecreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyManagerStorageBurnableSharesDecreased represents a BurnableSharesDecreased event raised by the StrategyManagerStorage contract.
type StrategyManagerStorageBurnableSharesDecreased struct {
	Strategy common.Address
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBurnableSharesDecreased is a free log retrieval operation binding the contract event 0xd9d082c3ec4f3a3ffa55c324939a06407f5fbcb87d5e0ce3b9508c92c84ed839.
//
// Solidity: event BurnableSharesDecreased(address strategy, uint256 shares)
func (_StrategyManagerStorage *StrategyManagerStorageFilterer) FilterBurnableSharesDecreased(opts *bind.FilterOpts) (*StrategyManagerStorageBurnableSharesDecreasedIterator, error) {

	logs, sub, err := _StrategyManagerStorage.contract.FilterLogs(opts, "BurnableSharesDecreased")
	if err != nil {
		return nil, err
	}
	return &StrategyManagerStorageBurnableSharesDecreasedIterator{contract: _StrategyManagerStorage.contract, event: "BurnableSharesDecreased", logs: logs, sub: sub}, nil
}

// WatchBurnableSharesDecreased is a free log subscription operation binding the contract event 0xd9d082c3ec4f3a3ffa55c324939a06407f5fbcb87d5e0ce3b9508c92c84ed839.
//
// Solidity: event BurnableSharesDecreased(address strategy, uint256 shares)
func (_StrategyManagerStorage *StrategyManagerStorageFilterer) WatchBurnableSharesDecreased(opts *bind.WatchOpts, sink chan<- *StrategyManagerStorageBurnableSharesDecreased) (event.Subscription, error) {

	logs, sub, err := _StrategyManagerStorage.contract.WatchLogs(opts, "BurnableSharesDecreased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyManagerStorageBurnableSharesDecreased)
				if err := _StrategyManagerStorage.contract.UnpackLog(event, "BurnableSharesDecreased", log); err != nil {
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

// ParseBurnableSharesDecreased is a log parse operation binding the contract event 0xd9d082c3ec4f3a3ffa55c324939a06407f5fbcb87d5e0ce3b9508c92c84ed839.
//
// Solidity: event BurnableSharesDecreased(address strategy, uint256 shares)
func (_StrategyManagerStorage *StrategyManagerStorageFilterer) ParseBurnableSharesDecreased(log types.Log) (*StrategyManagerStorageBurnableSharesDecreased, error) {
	event := new(StrategyManagerStorageBurnableSharesDecreased)
	if err := _StrategyManagerStorage.contract.UnpackLog(event, "BurnableSharesDecreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyManagerStorageBurnableSharesIncreasedIterator is returned from FilterBurnableSharesIncreased and is used to iterate over the raw logs and unpacked data for BurnableSharesIncreased events raised by the StrategyManagerStorage contract.
type StrategyManagerStorageBurnableSharesIncreasedIterator struct {
	Event *StrategyManagerStorageBurnableSharesIncreased // Event containing the contract specifics and raw log

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
func (it *StrategyManagerStorageBurnableSharesIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyManagerStorageBurnableSharesIncreased)
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
		it.Event = new(StrategyManagerStorageBurnableSharesIncreased)
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
func (it *StrategyManagerStorageBurnableSharesIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyManagerStorageBurnableSharesIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyManagerStorageBurnableSharesIncreased represents a BurnableSharesIncreased event raised by the StrategyManagerStorage contract.
type StrategyManagerStorageBurnableSharesIncreased struct {
	Strategy common.Address
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBurnableSharesIncreased is a free log retrieval operation binding the contract event 0xca3e02a4ab7ad3c47a8e36e5a624c30170791726ab720f1babfef21046d953ff.
//
// Solidity: event BurnableSharesIncreased(address strategy, uint256 shares)
func (_StrategyManagerStorage *StrategyManagerStorageFilterer) FilterBurnableSharesIncreased(opts *bind.FilterOpts) (*StrategyManagerStorageBurnableSharesIncreasedIterator, error) {

	logs, sub, err := _StrategyManagerStorage.contract.FilterLogs(opts, "BurnableSharesIncreased")
	if err != nil {
		return nil, err
	}
	return &StrategyManagerStorageBurnableSharesIncreasedIterator{contract: _StrategyManagerStorage.contract, event: "BurnableSharesIncreased", logs: logs, sub: sub}, nil
}

// WatchBurnableSharesIncreased is a free log subscription operation binding the contract event 0xca3e02a4ab7ad3c47a8e36e5a624c30170791726ab720f1babfef21046d953ff.
//
// Solidity: event BurnableSharesIncreased(address strategy, uint256 shares)
func (_StrategyManagerStorage *StrategyManagerStorageFilterer) WatchBurnableSharesIncreased(opts *bind.WatchOpts, sink chan<- *StrategyManagerStorageBurnableSharesIncreased) (event.Subscription, error) {

	logs, sub, err := _StrategyManagerStorage.contract.WatchLogs(opts, "BurnableSharesIncreased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyManagerStorageBurnableSharesIncreased)
				if err := _StrategyManagerStorage.contract.UnpackLog(event, "BurnableSharesIncreased", log); err != nil {
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

// ParseBurnableSharesIncreased is a log parse operation binding the contract event 0xca3e02a4ab7ad3c47a8e36e5a624c30170791726ab720f1babfef21046d953ff.
//
// Solidity: event BurnableSharesIncreased(address strategy, uint256 shares)
func (_StrategyManagerStorage *StrategyManagerStorageFilterer) ParseBurnableSharesIncreased(log types.Log) (*StrategyManagerStorageBurnableSharesIncreased, error) {
	event := new(StrategyManagerStorageBurnableSharesIncreased)
	if err := _StrategyManagerStorage.contract.UnpackLog(event, "BurnableSharesIncreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyManagerStorageDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the StrategyManagerStorage contract.
type StrategyManagerStorageDepositIterator struct {
	Event *StrategyManagerStorageDeposit // Event containing the contract specifics and raw log

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
func (it *StrategyManagerStorageDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyManagerStorageDeposit)
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
		it.Event = new(StrategyManagerStorageDeposit)
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
func (it *StrategyManagerStorageDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyManagerStorageDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyManagerStorageDeposit represents a Deposit event raised by the StrategyManagerStorage contract.
type StrategyManagerStorageDeposit struct {
	Staker   common.Address
	Strategy common.Address
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: event Deposit(address staker, address strategy, uint256 shares)
func (_StrategyManagerStorage *StrategyManagerStorageFilterer) FilterDeposit(opts *bind.FilterOpts) (*StrategyManagerStorageDepositIterator, error) {

	logs, sub, err := _StrategyManagerStorage.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &StrategyManagerStorageDepositIterator{contract: _StrategyManagerStorage.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: event Deposit(address staker, address strategy, uint256 shares)
func (_StrategyManagerStorage *StrategyManagerStorageFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *StrategyManagerStorageDeposit) (event.Subscription, error) {

	logs, sub, err := _StrategyManagerStorage.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyManagerStorageDeposit)
				if err := _StrategyManagerStorage.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: event Deposit(address staker, address strategy, uint256 shares)
func (_StrategyManagerStorage *StrategyManagerStorageFilterer) ParseDeposit(log types.Log) (*StrategyManagerStorageDeposit, error) {
	event := new(StrategyManagerStorageDeposit)
	if err := _StrategyManagerStorage.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyManagerStorageStrategyAddedToDepositWhitelistIterator is returned from FilterStrategyAddedToDepositWhitelist and is used to iterate over the raw logs and unpacked data for StrategyAddedToDepositWhitelist events raised by the StrategyManagerStorage contract.
type StrategyManagerStorageStrategyAddedToDepositWhitelistIterator struct {
	Event *StrategyManagerStorageStrategyAddedToDepositWhitelist // Event containing the contract specifics and raw log

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
func (it *StrategyManagerStorageStrategyAddedToDepositWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyManagerStorageStrategyAddedToDepositWhitelist)
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
		it.Event = new(StrategyManagerStorageStrategyAddedToDepositWhitelist)
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
func (it *StrategyManagerStorageStrategyAddedToDepositWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyManagerStorageStrategyAddedToDepositWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyManagerStorageStrategyAddedToDepositWhitelist represents a StrategyAddedToDepositWhitelist event raised by the StrategyManagerStorage contract.
type StrategyManagerStorageStrategyAddedToDepositWhitelist struct {
	Strategy common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStrategyAddedToDepositWhitelist is a free log retrieval operation binding the contract event 0x0c35b17d91c96eb2751cd456e1252f42a386e524ef9ff26ecc9950859fdc04fe.
//
// Solidity: event StrategyAddedToDepositWhitelist(address strategy)
func (_StrategyManagerStorage *StrategyManagerStorageFilterer) FilterStrategyAddedToDepositWhitelist(opts *bind.FilterOpts) (*StrategyManagerStorageStrategyAddedToDepositWhitelistIterator, error) {

	logs, sub, err := _StrategyManagerStorage.contract.FilterLogs(opts, "StrategyAddedToDepositWhitelist")
	if err != nil {
		return nil, err
	}
	return &StrategyManagerStorageStrategyAddedToDepositWhitelistIterator{contract: _StrategyManagerStorage.contract, event: "StrategyAddedToDepositWhitelist", logs: logs, sub: sub}, nil
}

// WatchStrategyAddedToDepositWhitelist is a free log subscription operation binding the contract event 0x0c35b17d91c96eb2751cd456e1252f42a386e524ef9ff26ecc9950859fdc04fe.
//
// Solidity: event StrategyAddedToDepositWhitelist(address strategy)
func (_StrategyManagerStorage *StrategyManagerStorageFilterer) WatchStrategyAddedToDepositWhitelist(opts *bind.WatchOpts, sink chan<- *StrategyManagerStorageStrategyAddedToDepositWhitelist) (event.Subscription, error) {

	logs, sub, err := _StrategyManagerStorage.contract.WatchLogs(opts, "StrategyAddedToDepositWhitelist")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyManagerStorageStrategyAddedToDepositWhitelist)
				if err := _StrategyManagerStorage.contract.UnpackLog(event, "StrategyAddedToDepositWhitelist", log); err != nil {
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

// ParseStrategyAddedToDepositWhitelist is a log parse operation binding the contract event 0x0c35b17d91c96eb2751cd456e1252f42a386e524ef9ff26ecc9950859fdc04fe.
//
// Solidity: event StrategyAddedToDepositWhitelist(address strategy)
func (_StrategyManagerStorage *StrategyManagerStorageFilterer) ParseStrategyAddedToDepositWhitelist(log types.Log) (*StrategyManagerStorageStrategyAddedToDepositWhitelist, error) {
	event := new(StrategyManagerStorageStrategyAddedToDepositWhitelist)
	if err := _StrategyManagerStorage.contract.UnpackLog(event, "StrategyAddedToDepositWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyManagerStorageStrategyRemovedFromDepositWhitelistIterator is returned from FilterStrategyRemovedFromDepositWhitelist and is used to iterate over the raw logs and unpacked data for StrategyRemovedFromDepositWhitelist events raised by the StrategyManagerStorage contract.
type StrategyManagerStorageStrategyRemovedFromDepositWhitelistIterator struct {
	Event *StrategyManagerStorageStrategyRemovedFromDepositWhitelist // Event containing the contract specifics and raw log

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
func (it *StrategyManagerStorageStrategyRemovedFromDepositWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyManagerStorageStrategyRemovedFromDepositWhitelist)
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
		it.Event = new(StrategyManagerStorageStrategyRemovedFromDepositWhitelist)
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
func (it *StrategyManagerStorageStrategyRemovedFromDepositWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyManagerStorageStrategyRemovedFromDepositWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyManagerStorageStrategyRemovedFromDepositWhitelist represents a StrategyRemovedFromDepositWhitelist event raised by the StrategyManagerStorage contract.
type StrategyManagerStorageStrategyRemovedFromDepositWhitelist struct {
	Strategy common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStrategyRemovedFromDepositWhitelist is a free log retrieval operation binding the contract event 0x4074413b4b443e4e58019f2855a8765113358c7c72e39509c6af45fc0f5ba030.
//
// Solidity: event StrategyRemovedFromDepositWhitelist(address strategy)
func (_StrategyManagerStorage *StrategyManagerStorageFilterer) FilterStrategyRemovedFromDepositWhitelist(opts *bind.FilterOpts) (*StrategyManagerStorageStrategyRemovedFromDepositWhitelistIterator, error) {

	logs, sub, err := _StrategyManagerStorage.contract.FilterLogs(opts, "StrategyRemovedFromDepositWhitelist")
	if err != nil {
		return nil, err
	}
	return &StrategyManagerStorageStrategyRemovedFromDepositWhitelistIterator{contract: _StrategyManagerStorage.contract, event: "StrategyRemovedFromDepositWhitelist", logs: logs, sub: sub}, nil
}

// WatchStrategyRemovedFromDepositWhitelist is a free log subscription operation binding the contract event 0x4074413b4b443e4e58019f2855a8765113358c7c72e39509c6af45fc0f5ba030.
//
// Solidity: event StrategyRemovedFromDepositWhitelist(address strategy)
func (_StrategyManagerStorage *StrategyManagerStorageFilterer) WatchStrategyRemovedFromDepositWhitelist(opts *bind.WatchOpts, sink chan<- *StrategyManagerStorageStrategyRemovedFromDepositWhitelist) (event.Subscription, error) {

	logs, sub, err := _StrategyManagerStorage.contract.WatchLogs(opts, "StrategyRemovedFromDepositWhitelist")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyManagerStorageStrategyRemovedFromDepositWhitelist)
				if err := _StrategyManagerStorage.contract.UnpackLog(event, "StrategyRemovedFromDepositWhitelist", log); err != nil {
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

// ParseStrategyRemovedFromDepositWhitelist is a log parse operation binding the contract event 0x4074413b4b443e4e58019f2855a8765113358c7c72e39509c6af45fc0f5ba030.
//
// Solidity: event StrategyRemovedFromDepositWhitelist(address strategy)
func (_StrategyManagerStorage *StrategyManagerStorageFilterer) ParseStrategyRemovedFromDepositWhitelist(log types.Log) (*StrategyManagerStorageStrategyRemovedFromDepositWhitelist, error) {
	event := new(StrategyManagerStorageStrategyRemovedFromDepositWhitelist)
	if err := _StrategyManagerStorage.contract.UnpackLog(event, "StrategyRemovedFromDepositWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyManagerStorageStrategyWhitelisterChangedIterator is returned from FilterStrategyWhitelisterChanged and is used to iterate over the raw logs and unpacked data for StrategyWhitelisterChanged events raised by the StrategyManagerStorage contract.
type StrategyManagerStorageStrategyWhitelisterChangedIterator struct {
	Event *StrategyManagerStorageStrategyWhitelisterChanged // Event containing the contract specifics and raw log

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
func (it *StrategyManagerStorageStrategyWhitelisterChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyManagerStorageStrategyWhitelisterChanged)
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
		it.Event = new(StrategyManagerStorageStrategyWhitelisterChanged)
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
func (it *StrategyManagerStorageStrategyWhitelisterChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyManagerStorageStrategyWhitelisterChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyManagerStorageStrategyWhitelisterChanged represents a StrategyWhitelisterChanged event raised by the StrategyManagerStorage contract.
type StrategyManagerStorageStrategyWhitelisterChanged struct {
	PreviousAddress common.Address
	NewAddress      common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterStrategyWhitelisterChanged is a free log retrieval operation binding the contract event 0x4264275e593955ff9d6146a51a4525f6ddace2e81db9391abcc9d1ca48047d29.
//
// Solidity: event StrategyWhitelisterChanged(address previousAddress, address newAddress)
func (_StrategyManagerStorage *StrategyManagerStorageFilterer) FilterStrategyWhitelisterChanged(opts *bind.FilterOpts) (*StrategyManagerStorageStrategyWhitelisterChangedIterator, error) {

	logs, sub, err := _StrategyManagerStorage.contract.FilterLogs(opts, "StrategyWhitelisterChanged")
	if err != nil {
		return nil, err
	}
	return &StrategyManagerStorageStrategyWhitelisterChangedIterator{contract: _StrategyManagerStorage.contract, event: "StrategyWhitelisterChanged", logs: logs, sub: sub}, nil
}

// WatchStrategyWhitelisterChanged is a free log subscription operation binding the contract event 0x4264275e593955ff9d6146a51a4525f6ddace2e81db9391abcc9d1ca48047d29.
//
// Solidity: event StrategyWhitelisterChanged(address previousAddress, address newAddress)
func (_StrategyManagerStorage *StrategyManagerStorageFilterer) WatchStrategyWhitelisterChanged(opts *bind.WatchOpts, sink chan<- *StrategyManagerStorageStrategyWhitelisterChanged) (event.Subscription, error) {

	logs, sub, err := _StrategyManagerStorage.contract.WatchLogs(opts, "StrategyWhitelisterChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyManagerStorageStrategyWhitelisterChanged)
				if err := _StrategyManagerStorage.contract.UnpackLog(event, "StrategyWhitelisterChanged", log); err != nil {
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

// ParseStrategyWhitelisterChanged is a log parse operation binding the contract event 0x4264275e593955ff9d6146a51a4525f6ddace2e81db9391abcc9d1ca48047d29.
//
// Solidity: event StrategyWhitelisterChanged(address previousAddress, address newAddress)
func (_StrategyManagerStorage *StrategyManagerStorageFilterer) ParseStrategyWhitelisterChanged(log types.Log) (*StrategyManagerStorageStrategyWhitelisterChanged, error) {
	event := new(StrategyManagerStorageStrategyWhitelisterChanged)
	if err := _StrategyManagerStorage.contract.UnpackLog(event, "StrategyWhitelisterChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
