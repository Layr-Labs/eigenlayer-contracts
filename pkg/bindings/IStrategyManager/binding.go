// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IStrategyManager

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

// IStrategyManagerMetaData contains all meta data concerning the IStrategyManager contract.
var IStrategyManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"addShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addStrategiesToDepositWhitelist\",\"inputs\":[{\"name\":\"strategiesToWhitelist\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burnShares\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"calculateStrategyDepositDigestHash\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"depositIntoStrategy\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"depositShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositIntoStrategyWithSignature\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"depositShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getBurnableShares\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDeposits\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStakerStrategyList\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStrategiesWithBurnableShares\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"increaseBurnableShares\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"addedSharesToBurn\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialStrategyWhitelister\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeDepositShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"depositSharesToRemove\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeStrategiesFromDepositWhitelist\",\"inputs\":[{\"name\":\"strategiesToRemoveFromWhitelist\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStrategyWhitelister\",\"inputs\":[{\"name\":\"newStrategyWhitelister\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakerDepositShares\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakerStrategyListLength\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"strategyIsWhitelistedForDeposit\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"strategyWhitelister\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawSharesAsTokens\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"BurnableSharesDecreased\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BurnableSharesIncreased\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposit\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyAddedToDepositWhitelist\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyRemovedFromDepositWhitelist\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyWhitelisterChanged\",\"inputs\":[{\"name\":\"previousAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"MaxStrategiesExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyDelegationManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyStrategyWhitelister\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SharesAmountTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SharesAmountZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StakerAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategyNotFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategyNotWhitelisted\",\"inputs\":[]}]",
}

// IStrategyManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use IStrategyManagerMetaData.ABI instead.
var IStrategyManagerABI = IStrategyManagerMetaData.ABI

// IStrategyManager is an auto generated Go binding around an Ethereum contract.
type IStrategyManager struct {
	IStrategyManagerCaller     // Read-only binding to the contract
	IStrategyManagerTransactor // Write-only binding to the contract
	IStrategyManagerFilterer   // Log filterer for contract events
}

// IStrategyManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IStrategyManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStrategyManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IStrategyManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStrategyManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IStrategyManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStrategyManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IStrategyManagerSession struct {
	Contract     *IStrategyManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IStrategyManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IStrategyManagerCallerSession struct {
	Contract *IStrategyManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IStrategyManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IStrategyManagerTransactorSession struct {
	Contract     *IStrategyManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IStrategyManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IStrategyManagerRaw struct {
	Contract *IStrategyManager // Generic contract binding to access the raw methods on
}

// IStrategyManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IStrategyManagerCallerRaw struct {
	Contract *IStrategyManagerCaller // Generic read-only contract binding to access the raw methods on
}

// IStrategyManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IStrategyManagerTransactorRaw struct {
	Contract *IStrategyManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIStrategyManager creates a new instance of IStrategyManager, bound to a specific deployed contract.
func NewIStrategyManager(address common.Address, backend bind.ContractBackend) (*IStrategyManager, error) {
	contract, err := bindIStrategyManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IStrategyManager{IStrategyManagerCaller: IStrategyManagerCaller{contract: contract}, IStrategyManagerTransactor: IStrategyManagerTransactor{contract: contract}, IStrategyManagerFilterer: IStrategyManagerFilterer{contract: contract}}, nil
}

// NewIStrategyManagerCaller creates a new read-only instance of IStrategyManager, bound to a specific deployed contract.
func NewIStrategyManagerCaller(address common.Address, caller bind.ContractCaller) (*IStrategyManagerCaller, error) {
	contract, err := bindIStrategyManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IStrategyManagerCaller{contract: contract}, nil
}

// NewIStrategyManagerTransactor creates a new write-only instance of IStrategyManager, bound to a specific deployed contract.
func NewIStrategyManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IStrategyManagerTransactor, error) {
	contract, err := bindIStrategyManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IStrategyManagerTransactor{contract: contract}, nil
}

// NewIStrategyManagerFilterer creates a new log filterer instance of IStrategyManager, bound to a specific deployed contract.
func NewIStrategyManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IStrategyManagerFilterer, error) {
	contract, err := bindIStrategyManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IStrategyManagerFilterer{contract: contract}, nil
}

// bindIStrategyManager binds a generic wrapper to an already deployed contract.
func bindIStrategyManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IStrategyManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStrategyManager *IStrategyManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStrategyManager.Contract.IStrategyManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStrategyManager *IStrategyManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStrategyManager.Contract.IStrategyManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStrategyManager *IStrategyManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStrategyManager.Contract.IStrategyManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStrategyManager *IStrategyManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStrategyManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStrategyManager *IStrategyManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStrategyManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStrategyManager *IStrategyManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStrategyManager.Contract.contract.Transact(opts, method, params...)
}

// CalculateStrategyDepositDigestHash is a free data retrieval call binding the contract method 0x9ac01d61.
//
// Solidity: function calculateStrategyDepositDigestHash(address staker, address strategy, address token, uint256 amount, uint256 nonce, uint256 expiry) view returns(bytes32)
func (_IStrategyManager *IStrategyManagerCaller) CalculateStrategyDepositDigestHash(opts *bind.CallOpts, staker common.Address, strategy common.Address, token common.Address, amount *big.Int, nonce *big.Int, expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _IStrategyManager.contract.Call(opts, &out, "calculateStrategyDepositDigestHash", staker, strategy, token, amount, nonce, expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateStrategyDepositDigestHash is a free data retrieval call binding the contract method 0x9ac01d61.
//
// Solidity: function calculateStrategyDepositDigestHash(address staker, address strategy, address token, uint256 amount, uint256 nonce, uint256 expiry) view returns(bytes32)
func (_IStrategyManager *IStrategyManagerSession) CalculateStrategyDepositDigestHash(staker common.Address, strategy common.Address, token common.Address, amount *big.Int, nonce *big.Int, expiry *big.Int) ([32]byte, error) {
	return _IStrategyManager.Contract.CalculateStrategyDepositDigestHash(&_IStrategyManager.CallOpts, staker, strategy, token, amount, nonce, expiry)
}

// CalculateStrategyDepositDigestHash is a free data retrieval call binding the contract method 0x9ac01d61.
//
// Solidity: function calculateStrategyDepositDigestHash(address staker, address strategy, address token, uint256 amount, uint256 nonce, uint256 expiry) view returns(bytes32)
func (_IStrategyManager *IStrategyManagerCallerSession) CalculateStrategyDepositDigestHash(staker common.Address, strategy common.Address, token common.Address, amount *big.Int, nonce *big.Int, expiry *big.Int) ([32]byte, error) {
	return _IStrategyManager.Contract.CalculateStrategyDepositDigestHash(&_IStrategyManager.CallOpts, staker, strategy, token, amount, nonce, expiry)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_IStrategyManager *IStrategyManagerCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IStrategyManager.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_IStrategyManager *IStrategyManagerSession) Delegation() (common.Address, error) {
	return _IStrategyManager.Contract.Delegation(&_IStrategyManager.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_IStrategyManager *IStrategyManagerCallerSession) Delegation() (common.Address, error) {
	return _IStrategyManager.Contract.Delegation(&_IStrategyManager.CallOpts)
}

// GetBurnableShares is a free data retrieval call binding the contract method 0xfd980423.
//
// Solidity: function getBurnableShares(address strategy) view returns(uint256)
func (_IStrategyManager *IStrategyManagerCaller) GetBurnableShares(opts *bind.CallOpts, strategy common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IStrategyManager.contract.Call(opts, &out, "getBurnableShares", strategy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBurnableShares is a free data retrieval call binding the contract method 0xfd980423.
//
// Solidity: function getBurnableShares(address strategy) view returns(uint256)
func (_IStrategyManager *IStrategyManagerSession) GetBurnableShares(strategy common.Address) (*big.Int, error) {
	return _IStrategyManager.Contract.GetBurnableShares(&_IStrategyManager.CallOpts, strategy)
}

// GetBurnableShares is a free data retrieval call binding the contract method 0xfd980423.
//
// Solidity: function getBurnableShares(address strategy) view returns(uint256)
func (_IStrategyManager *IStrategyManagerCallerSession) GetBurnableShares(strategy common.Address) (*big.Int, error) {
	return _IStrategyManager.Contract.GetBurnableShares(&_IStrategyManager.CallOpts, strategy)
}

// GetDeposits is a free data retrieval call binding the contract method 0x94f649dd.
//
// Solidity: function getDeposits(address staker) view returns(address[], uint256[])
func (_IStrategyManager *IStrategyManagerCaller) GetDeposits(opts *bind.CallOpts, staker common.Address) ([]common.Address, []*big.Int, error) {
	var out []interface{}
	err := _IStrategyManager.contract.Call(opts, &out, "getDeposits", staker)

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
func (_IStrategyManager *IStrategyManagerSession) GetDeposits(staker common.Address) ([]common.Address, []*big.Int, error) {
	return _IStrategyManager.Contract.GetDeposits(&_IStrategyManager.CallOpts, staker)
}

// GetDeposits is a free data retrieval call binding the contract method 0x94f649dd.
//
// Solidity: function getDeposits(address staker) view returns(address[], uint256[])
func (_IStrategyManager *IStrategyManagerCallerSession) GetDeposits(staker common.Address) ([]common.Address, []*big.Int, error) {
	return _IStrategyManager.Contract.GetDeposits(&_IStrategyManager.CallOpts, staker)
}

// GetStakerStrategyList is a free data retrieval call binding the contract method 0xde44acb6.
//
// Solidity: function getStakerStrategyList(address staker) view returns(address[])
func (_IStrategyManager *IStrategyManagerCaller) GetStakerStrategyList(opts *bind.CallOpts, staker common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _IStrategyManager.contract.Call(opts, &out, "getStakerStrategyList", staker)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetStakerStrategyList is a free data retrieval call binding the contract method 0xde44acb6.
//
// Solidity: function getStakerStrategyList(address staker) view returns(address[])
func (_IStrategyManager *IStrategyManagerSession) GetStakerStrategyList(staker common.Address) ([]common.Address, error) {
	return _IStrategyManager.Contract.GetStakerStrategyList(&_IStrategyManager.CallOpts, staker)
}

// GetStakerStrategyList is a free data retrieval call binding the contract method 0xde44acb6.
//
// Solidity: function getStakerStrategyList(address staker) view returns(address[])
func (_IStrategyManager *IStrategyManagerCallerSession) GetStakerStrategyList(staker common.Address) ([]common.Address, error) {
	return _IStrategyManager.Contract.GetStakerStrategyList(&_IStrategyManager.CallOpts, staker)
}

// GetStrategiesWithBurnableShares is a free data retrieval call binding the contract method 0x36a8c500.
//
// Solidity: function getStrategiesWithBurnableShares() view returns(address[], uint256[])
func (_IStrategyManager *IStrategyManagerCaller) GetStrategiesWithBurnableShares(opts *bind.CallOpts) ([]common.Address, []*big.Int, error) {
	var out []interface{}
	err := _IStrategyManager.contract.Call(opts, &out, "getStrategiesWithBurnableShares")

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
func (_IStrategyManager *IStrategyManagerSession) GetStrategiesWithBurnableShares() ([]common.Address, []*big.Int, error) {
	return _IStrategyManager.Contract.GetStrategiesWithBurnableShares(&_IStrategyManager.CallOpts)
}

// GetStrategiesWithBurnableShares is a free data retrieval call binding the contract method 0x36a8c500.
//
// Solidity: function getStrategiesWithBurnableShares() view returns(address[], uint256[])
func (_IStrategyManager *IStrategyManagerCallerSession) GetStrategiesWithBurnableShares() ([]common.Address, []*big.Int, error) {
	return _IStrategyManager.Contract.GetStrategiesWithBurnableShares(&_IStrategyManager.CallOpts)
}

// StakerDepositShares is a free data retrieval call binding the contract method 0xfe243a17.
//
// Solidity: function stakerDepositShares(address user, address strategy) view returns(uint256 shares)
func (_IStrategyManager *IStrategyManagerCaller) StakerDepositShares(opts *bind.CallOpts, user common.Address, strategy common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IStrategyManager.contract.Call(opts, &out, "stakerDepositShares", user, strategy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerDepositShares is a free data retrieval call binding the contract method 0xfe243a17.
//
// Solidity: function stakerDepositShares(address user, address strategy) view returns(uint256 shares)
func (_IStrategyManager *IStrategyManagerSession) StakerDepositShares(user common.Address, strategy common.Address) (*big.Int, error) {
	return _IStrategyManager.Contract.StakerDepositShares(&_IStrategyManager.CallOpts, user, strategy)
}

// StakerDepositShares is a free data retrieval call binding the contract method 0xfe243a17.
//
// Solidity: function stakerDepositShares(address user, address strategy) view returns(uint256 shares)
func (_IStrategyManager *IStrategyManagerCallerSession) StakerDepositShares(user common.Address, strategy common.Address) (*big.Int, error) {
	return _IStrategyManager.Contract.StakerDepositShares(&_IStrategyManager.CallOpts, user, strategy)
}

// StakerStrategyListLength is a free data retrieval call binding the contract method 0x8b8aac3c.
//
// Solidity: function stakerStrategyListLength(address staker) view returns(uint256)
func (_IStrategyManager *IStrategyManagerCaller) StakerStrategyListLength(opts *bind.CallOpts, staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IStrategyManager.contract.Call(opts, &out, "stakerStrategyListLength", staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerStrategyListLength is a free data retrieval call binding the contract method 0x8b8aac3c.
//
// Solidity: function stakerStrategyListLength(address staker) view returns(uint256)
func (_IStrategyManager *IStrategyManagerSession) StakerStrategyListLength(staker common.Address) (*big.Int, error) {
	return _IStrategyManager.Contract.StakerStrategyListLength(&_IStrategyManager.CallOpts, staker)
}

// StakerStrategyListLength is a free data retrieval call binding the contract method 0x8b8aac3c.
//
// Solidity: function stakerStrategyListLength(address staker) view returns(uint256)
func (_IStrategyManager *IStrategyManagerCallerSession) StakerStrategyListLength(staker common.Address) (*big.Int, error) {
	return _IStrategyManager.Contract.StakerStrategyListLength(&_IStrategyManager.CallOpts, staker)
}

// StrategyIsWhitelistedForDeposit is a free data retrieval call binding the contract method 0x663c1de4.
//
// Solidity: function strategyIsWhitelistedForDeposit(address strategy) view returns(bool)
func (_IStrategyManager *IStrategyManagerCaller) StrategyIsWhitelistedForDeposit(opts *bind.CallOpts, strategy common.Address) (bool, error) {
	var out []interface{}
	err := _IStrategyManager.contract.Call(opts, &out, "strategyIsWhitelistedForDeposit", strategy)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StrategyIsWhitelistedForDeposit is a free data retrieval call binding the contract method 0x663c1de4.
//
// Solidity: function strategyIsWhitelistedForDeposit(address strategy) view returns(bool)
func (_IStrategyManager *IStrategyManagerSession) StrategyIsWhitelistedForDeposit(strategy common.Address) (bool, error) {
	return _IStrategyManager.Contract.StrategyIsWhitelistedForDeposit(&_IStrategyManager.CallOpts, strategy)
}

// StrategyIsWhitelistedForDeposit is a free data retrieval call binding the contract method 0x663c1de4.
//
// Solidity: function strategyIsWhitelistedForDeposit(address strategy) view returns(bool)
func (_IStrategyManager *IStrategyManagerCallerSession) StrategyIsWhitelistedForDeposit(strategy common.Address) (bool, error) {
	return _IStrategyManager.Contract.StrategyIsWhitelistedForDeposit(&_IStrategyManager.CallOpts, strategy)
}

// StrategyWhitelister is a free data retrieval call binding the contract method 0x967fc0d2.
//
// Solidity: function strategyWhitelister() view returns(address)
func (_IStrategyManager *IStrategyManagerCaller) StrategyWhitelister(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IStrategyManager.contract.Call(opts, &out, "strategyWhitelister")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StrategyWhitelister is a free data retrieval call binding the contract method 0x967fc0d2.
//
// Solidity: function strategyWhitelister() view returns(address)
func (_IStrategyManager *IStrategyManagerSession) StrategyWhitelister() (common.Address, error) {
	return _IStrategyManager.Contract.StrategyWhitelister(&_IStrategyManager.CallOpts)
}

// StrategyWhitelister is a free data retrieval call binding the contract method 0x967fc0d2.
//
// Solidity: function strategyWhitelister() view returns(address)
func (_IStrategyManager *IStrategyManagerCallerSession) StrategyWhitelister() (common.Address, error) {
	return _IStrategyManager.Contract.StrategyWhitelister(&_IStrategyManager.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_IStrategyManager *IStrategyManagerCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IStrategyManager.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_IStrategyManager *IStrategyManagerSession) Version() (string, error) {
	return _IStrategyManager.Contract.Version(&_IStrategyManager.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_IStrategyManager *IStrategyManagerCallerSession) Version() (string, error) {
	return _IStrategyManager.Contract.Version(&_IStrategyManager.CallOpts)
}

// AddShares is a paid mutator transaction binding the contract method 0x50ff7225.
//
// Solidity: function addShares(address staker, address strategy, uint256 shares) returns(uint256, uint256)
func (_IStrategyManager *IStrategyManagerTransactor) AddShares(opts *bind.TransactOpts, staker common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error) {
	return _IStrategyManager.contract.Transact(opts, "addShares", staker, strategy, shares)
}

// AddShares is a paid mutator transaction binding the contract method 0x50ff7225.
//
// Solidity: function addShares(address staker, address strategy, uint256 shares) returns(uint256, uint256)
func (_IStrategyManager *IStrategyManagerSession) AddShares(staker common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error) {
	return _IStrategyManager.Contract.AddShares(&_IStrategyManager.TransactOpts, staker, strategy, shares)
}

// AddShares is a paid mutator transaction binding the contract method 0x50ff7225.
//
// Solidity: function addShares(address staker, address strategy, uint256 shares) returns(uint256, uint256)
func (_IStrategyManager *IStrategyManagerTransactorSession) AddShares(staker common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error) {
	return _IStrategyManager.Contract.AddShares(&_IStrategyManager.TransactOpts, staker, strategy, shares)
}

// AddStrategiesToDepositWhitelist is a paid mutator transaction binding the contract method 0x5de08ff2.
//
// Solidity: function addStrategiesToDepositWhitelist(address[] strategiesToWhitelist) returns()
func (_IStrategyManager *IStrategyManagerTransactor) AddStrategiesToDepositWhitelist(opts *bind.TransactOpts, strategiesToWhitelist []common.Address) (*types.Transaction, error) {
	return _IStrategyManager.contract.Transact(opts, "addStrategiesToDepositWhitelist", strategiesToWhitelist)
}

// AddStrategiesToDepositWhitelist is a paid mutator transaction binding the contract method 0x5de08ff2.
//
// Solidity: function addStrategiesToDepositWhitelist(address[] strategiesToWhitelist) returns()
func (_IStrategyManager *IStrategyManagerSession) AddStrategiesToDepositWhitelist(strategiesToWhitelist []common.Address) (*types.Transaction, error) {
	return _IStrategyManager.Contract.AddStrategiesToDepositWhitelist(&_IStrategyManager.TransactOpts, strategiesToWhitelist)
}

// AddStrategiesToDepositWhitelist is a paid mutator transaction binding the contract method 0x5de08ff2.
//
// Solidity: function addStrategiesToDepositWhitelist(address[] strategiesToWhitelist) returns()
func (_IStrategyManager *IStrategyManagerTransactorSession) AddStrategiesToDepositWhitelist(strategiesToWhitelist []common.Address) (*types.Transaction, error) {
	return _IStrategyManager.Contract.AddStrategiesToDepositWhitelist(&_IStrategyManager.TransactOpts, strategiesToWhitelist)
}

// BurnShares is a paid mutator transaction binding the contract method 0x4b6d5d6e.
//
// Solidity: function burnShares(address strategy) returns()
func (_IStrategyManager *IStrategyManagerTransactor) BurnShares(opts *bind.TransactOpts, strategy common.Address) (*types.Transaction, error) {
	return _IStrategyManager.contract.Transact(opts, "burnShares", strategy)
}

// BurnShares is a paid mutator transaction binding the contract method 0x4b6d5d6e.
//
// Solidity: function burnShares(address strategy) returns()
func (_IStrategyManager *IStrategyManagerSession) BurnShares(strategy common.Address) (*types.Transaction, error) {
	return _IStrategyManager.Contract.BurnShares(&_IStrategyManager.TransactOpts, strategy)
}

// BurnShares is a paid mutator transaction binding the contract method 0x4b6d5d6e.
//
// Solidity: function burnShares(address strategy) returns()
func (_IStrategyManager *IStrategyManagerTransactorSession) BurnShares(strategy common.Address) (*types.Transaction, error) {
	return _IStrategyManager.Contract.BurnShares(&_IStrategyManager.TransactOpts, strategy)
}

// DepositIntoStrategy is a paid mutator transaction binding the contract method 0xe7a050aa.
//
// Solidity: function depositIntoStrategy(address strategy, address token, uint256 amount) returns(uint256 depositShares)
func (_IStrategyManager *IStrategyManagerTransactor) DepositIntoStrategy(opts *bind.TransactOpts, strategy common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IStrategyManager.contract.Transact(opts, "depositIntoStrategy", strategy, token, amount)
}

// DepositIntoStrategy is a paid mutator transaction binding the contract method 0xe7a050aa.
//
// Solidity: function depositIntoStrategy(address strategy, address token, uint256 amount) returns(uint256 depositShares)
func (_IStrategyManager *IStrategyManagerSession) DepositIntoStrategy(strategy common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IStrategyManager.Contract.DepositIntoStrategy(&_IStrategyManager.TransactOpts, strategy, token, amount)
}

// DepositIntoStrategy is a paid mutator transaction binding the contract method 0xe7a050aa.
//
// Solidity: function depositIntoStrategy(address strategy, address token, uint256 amount) returns(uint256 depositShares)
func (_IStrategyManager *IStrategyManagerTransactorSession) DepositIntoStrategy(strategy common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IStrategyManager.Contract.DepositIntoStrategy(&_IStrategyManager.TransactOpts, strategy, token, amount)
}

// DepositIntoStrategyWithSignature is a paid mutator transaction binding the contract method 0x32e89ace.
//
// Solidity: function depositIntoStrategyWithSignature(address strategy, address token, uint256 amount, address staker, uint256 expiry, bytes signature) returns(uint256 depositShares)
func (_IStrategyManager *IStrategyManagerTransactor) DepositIntoStrategyWithSignature(opts *bind.TransactOpts, strategy common.Address, token common.Address, amount *big.Int, staker common.Address, expiry *big.Int, signature []byte) (*types.Transaction, error) {
	return _IStrategyManager.contract.Transact(opts, "depositIntoStrategyWithSignature", strategy, token, amount, staker, expiry, signature)
}

// DepositIntoStrategyWithSignature is a paid mutator transaction binding the contract method 0x32e89ace.
//
// Solidity: function depositIntoStrategyWithSignature(address strategy, address token, uint256 amount, address staker, uint256 expiry, bytes signature) returns(uint256 depositShares)
func (_IStrategyManager *IStrategyManagerSession) DepositIntoStrategyWithSignature(strategy common.Address, token common.Address, amount *big.Int, staker common.Address, expiry *big.Int, signature []byte) (*types.Transaction, error) {
	return _IStrategyManager.Contract.DepositIntoStrategyWithSignature(&_IStrategyManager.TransactOpts, strategy, token, amount, staker, expiry, signature)
}

// DepositIntoStrategyWithSignature is a paid mutator transaction binding the contract method 0x32e89ace.
//
// Solidity: function depositIntoStrategyWithSignature(address strategy, address token, uint256 amount, address staker, uint256 expiry, bytes signature) returns(uint256 depositShares)
func (_IStrategyManager *IStrategyManagerTransactorSession) DepositIntoStrategyWithSignature(strategy common.Address, token common.Address, amount *big.Int, staker common.Address, expiry *big.Int, signature []byte) (*types.Transaction, error) {
	return _IStrategyManager.Contract.DepositIntoStrategyWithSignature(&_IStrategyManager.TransactOpts, strategy, token, amount, staker, expiry, signature)
}

// IncreaseBurnableShares is a paid mutator transaction binding the contract method 0xdebe1eab.
//
// Solidity: function increaseBurnableShares(address strategy, uint256 addedSharesToBurn) returns()
func (_IStrategyManager *IStrategyManagerTransactor) IncreaseBurnableShares(opts *bind.TransactOpts, strategy common.Address, addedSharesToBurn *big.Int) (*types.Transaction, error) {
	return _IStrategyManager.contract.Transact(opts, "increaseBurnableShares", strategy, addedSharesToBurn)
}

// IncreaseBurnableShares is a paid mutator transaction binding the contract method 0xdebe1eab.
//
// Solidity: function increaseBurnableShares(address strategy, uint256 addedSharesToBurn) returns()
func (_IStrategyManager *IStrategyManagerSession) IncreaseBurnableShares(strategy common.Address, addedSharesToBurn *big.Int) (*types.Transaction, error) {
	return _IStrategyManager.Contract.IncreaseBurnableShares(&_IStrategyManager.TransactOpts, strategy, addedSharesToBurn)
}

// IncreaseBurnableShares is a paid mutator transaction binding the contract method 0xdebe1eab.
//
// Solidity: function increaseBurnableShares(address strategy, uint256 addedSharesToBurn) returns()
func (_IStrategyManager *IStrategyManagerTransactorSession) IncreaseBurnableShares(strategy common.Address, addedSharesToBurn *big.Int) (*types.Transaction, error) {
	return _IStrategyManager.Contract.IncreaseBurnableShares(&_IStrategyManager.TransactOpts, strategy, addedSharesToBurn)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address initialOwner, address initialStrategyWhitelister, uint256 initialPausedStatus) returns()
func (_IStrategyManager *IStrategyManagerTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, initialStrategyWhitelister common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _IStrategyManager.contract.Transact(opts, "initialize", initialOwner, initialStrategyWhitelister, initialPausedStatus)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address initialOwner, address initialStrategyWhitelister, uint256 initialPausedStatus) returns()
func (_IStrategyManager *IStrategyManagerSession) Initialize(initialOwner common.Address, initialStrategyWhitelister common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _IStrategyManager.Contract.Initialize(&_IStrategyManager.TransactOpts, initialOwner, initialStrategyWhitelister, initialPausedStatus)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address initialOwner, address initialStrategyWhitelister, uint256 initialPausedStatus) returns()
func (_IStrategyManager *IStrategyManagerTransactorSession) Initialize(initialOwner common.Address, initialStrategyWhitelister common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _IStrategyManager.Contract.Initialize(&_IStrategyManager.TransactOpts, initialOwner, initialStrategyWhitelister, initialPausedStatus)
}

// RemoveDepositShares is a paid mutator transaction binding the contract method 0x724af423.
//
// Solidity: function removeDepositShares(address staker, address strategy, uint256 depositSharesToRemove) returns(uint256)
func (_IStrategyManager *IStrategyManagerTransactor) RemoveDepositShares(opts *bind.TransactOpts, staker common.Address, strategy common.Address, depositSharesToRemove *big.Int) (*types.Transaction, error) {
	return _IStrategyManager.contract.Transact(opts, "removeDepositShares", staker, strategy, depositSharesToRemove)
}

// RemoveDepositShares is a paid mutator transaction binding the contract method 0x724af423.
//
// Solidity: function removeDepositShares(address staker, address strategy, uint256 depositSharesToRemove) returns(uint256)
func (_IStrategyManager *IStrategyManagerSession) RemoveDepositShares(staker common.Address, strategy common.Address, depositSharesToRemove *big.Int) (*types.Transaction, error) {
	return _IStrategyManager.Contract.RemoveDepositShares(&_IStrategyManager.TransactOpts, staker, strategy, depositSharesToRemove)
}

// RemoveDepositShares is a paid mutator transaction binding the contract method 0x724af423.
//
// Solidity: function removeDepositShares(address staker, address strategy, uint256 depositSharesToRemove) returns(uint256)
func (_IStrategyManager *IStrategyManagerTransactorSession) RemoveDepositShares(staker common.Address, strategy common.Address, depositSharesToRemove *big.Int) (*types.Transaction, error) {
	return _IStrategyManager.Contract.RemoveDepositShares(&_IStrategyManager.TransactOpts, staker, strategy, depositSharesToRemove)
}

// RemoveStrategiesFromDepositWhitelist is a paid mutator transaction binding the contract method 0xb5d8b5b8.
//
// Solidity: function removeStrategiesFromDepositWhitelist(address[] strategiesToRemoveFromWhitelist) returns()
func (_IStrategyManager *IStrategyManagerTransactor) RemoveStrategiesFromDepositWhitelist(opts *bind.TransactOpts, strategiesToRemoveFromWhitelist []common.Address) (*types.Transaction, error) {
	return _IStrategyManager.contract.Transact(opts, "removeStrategiesFromDepositWhitelist", strategiesToRemoveFromWhitelist)
}

// RemoveStrategiesFromDepositWhitelist is a paid mutator transaction binding the contract method 0xb5d8b5b8.
//
// Solidity: function removeStrategiesFromDepositWhitelist(address[] strategiesToRemoveFromWhitelist) returns()
func (_IStrategyManager *IStrategyManagerSession) RemoveStrategiesFromDepositWhitelist(strategiesToRemoveFromWhitelist []common.Address) (*types.Transaction, error) {
	return _IStrategyManager.Contract.RemoveStrategiesFromDepositWhitelist(&_IStrategyManager.TransactOpts, strategiesToRemoveFromWhitelist)
}

// RemoveStrategiesFromDepositWhitelist is a paid mutator transaction binding the contract method 0xb5d8b5b8.
//
// Solidity: function removeStrategiesFromDepositWhitelist(address[] strategiesToRemoveFromWhitelist) returns()
func (_IStrategyManager *IStrategyManagerTransactorSession) RemoveStrategiesFromDepositWhitelist(strategiesToRemoveFromWhitelist []common.Address) (*types.Transaction, error) {
	return _IStrategyManager.Contract.RemoveStrategiesFromDepositWhitelist(&_IStrategyManager.TransactOpts, strategiesToRemoveFromWhitelist)
}

// SetStrategyWhitelister is a paid mutator transaction binding the contract method 0xc6656702.
//
// Solidity: function setStrategyWhitelister(address newStrategyWhitelister) returns()
func (_IStrategyManager *IStrategyManagerTransactor) SetStrategyWhitelister(opts *bind.TransactOpts, newStrategyWhitelister common.Address) (*types.Transaction, error) {
	return _IStrategyManager.contract.Transact(opts, "setStrategyWhitelister", newStrategyWhitelister)
}

// SetStrategyWhitelister is a paid mutator transaction binding the contract method 0xc6656702.
//
// Solidity: function setStrategyWhitelister(address newStrategyWhitelister) returns()
func (_IStrategyManager *IStrategyManagerSession) SetStrategyWhitelister(newStrategyWhitelister common.Address) (*types.Transaction, error) {
	return _IStrategyManager.Contract.SetStrategyWhitelister(&_IStrategyManager.TransactOpts, newStrategyWhitelister)
}

// SetStrategyWhitelister is a paid mutator transaction binding the contract method 0xc6656702.
//
// Solidity: function setStrategyWhitelister(address newStrategyWhitelister) returns()
func (_IStrategyManager *IStrategyManagerTransactorSession) SetStrategyWhitelister(newStrategyWhitelister common.Address) (*types.Transaction, error) {
	return _IStrategyManager.Contract.SetStrategyWhitelister(&_IStrategyManager.TransactOpts, newStrategyWhitelister)
}

// WithdrawSharesAsTokens is a paid mutator transaction binding the contract method 0x2eae418c.
//
// Solidity: function withdrawSharesAsTokens(address staker, address strategy, address token, uint256 shares) returns()
func (_IStrategyManager *IStrategyManagerTransactor) WithdrawSharesAsTokens(opts *bind.TransactOpts, staker common.Address, strategy common.Address, token common.Address, shares *big.Int) (*types.Transaction, error) {
	return _IStrategyManager.contract.Transact(opts, "withdrawSharesAsTokens", staker, strategy, token, shares)
}

// WithdrawSharesAsTokens is a paid mutator transaction binding the contract method 0x2eae418c.
//
// Solidity: function withdrawSharesAsTokens(address staker, address strategy, address token, uint256 shares) returns()
func (_IStrategyManager *IStrategyManagerSession) WithdrawSharesAsTokens(staker common.Address, strategy common.Address, token common.Address, shares *big.Int) (*types.Transaction, error) {
	return _IStrategyManager.Contract.WithdrawSharesAsTokens(&_IStrategyManager.TransactOpts, staker, strategy, token, shares)
}

// WithdrawSharesAsTokens is a paid mutator transaction binding the contract method 0x2eae418c.
//
// Solidity: function withdrawSharesAsTokens(address staker, address strategy, address token, uint256 shares) returns()
func (_IStrategyManager *IStrategyManagerTransactorSession) WithdrawSharesAsTokens(staker common.Address, strategy common.Address, token common.Address, shares *big.Int) (*types.Transaction, error) {
	return _IStrategyManager.Contract.WithdrawSharesAsTokens(&_IStrategyManager.TransactOpts, staker, strategy, token, shares)
}

// IStrategyManagerBurnableSharesDecreasedIterator is returned from FilterBurnableSharesDecreased and is used to iterate over the raw logs and unpacked data for BurnableSharesDecreased events raised by the IStrategyManager contract.
type IStrategyManagerBurnableSharesDecreasedIterator struct {
	Event *IStrategyManagerBurnableSharesDecreased // Event containing the contract specifics and raw log

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
func (it *IStrategyManagerBurnableSharesDecreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStrategyManagerBurnableSharesDecreased)
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
		it.Event = new(IStrategyManagerBurnableSharesDecreased)
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
func (it *IStrategyManagerBurnableSharesDecreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStrategyManagerBurnableSharesDecreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStrategyManagerBurnableSharesDecreased represents a BurnableSharesDecreased event raised by the IStrategyManager contract.
type IStrategyManagerBurnableSharesDecreased struct {
	Strategy common.Address
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBurnableSharesDecreased is a free log retrieval operation binding the contract event 0xd9d082c3ec4f3a3ffa55c324939a06407f5fbcb87d5e0ce3b9508c92c84ed839.
//
// Solidity: event BurnableSharesDecreased(address strategy, uint256 shares)
func (_IStrategyManager *IStrategyManagerFilterer) FilterBurnableSharesDecreased(opts *bind.FilterOpts) (*IStrategyManagerBurnableSharesDecreasedIterator, error) {

	logs, sub, err := _IStrategyManager.contract.FilterLogs(opts, "BurnableSharesDecreased")
	if err != nil {
		return nil, err
	}
	return &IStrategyManagerBurnableSharesDecreasedIterator{contract: _IStrategyManager.contract, event: "BurnableSharesDecreased", logs: logs, sub: sub}, nil
}

// WatchBurnableSharesDecreased is a free log subscription operation binding the contract event 0xd9d082c3ec4f3a3ffa55c324939a06407f5fbcb87d5e0ce3b9508c92c84ed839.
//
// Solidity: event BurnableSharesDecreased(address strategy, uint256 shares)
func (_IStrategyManager *IStrategyManagerFilterer) WatchBurnableSharesDecreased(opts *bind.WatchOpts, sink chan<- *IStrategyManagerBurnableSharesDecreased) (event.Subscription, error) {

	logs, sub, err := _IStrategyManager.contract.WatchLogs(opts, "BurnableSharesDecreased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStrategyManagerBurnableSharesDecreased)
				if err := _IStrategyManager.contract.UnpackLog(event, "BurnableSharesDecreased", log); err != nil {
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
func (_IStrategyManager *IStrategyManagerFilterer) ParseBurnableSharesDecreased(log types.Log) (*IStrategyManagerBurnableSharesDecreased, error) {
	event := new(IStrategyManagerBurnableSharesDecreased)
	if err := _IStrategyManager.contract.UnpackLog(event, "BurnableSharesDecreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStrategyManagerBurnableSharesIncreasedIterator is returned from FilterBurnableSharesIncreased and is used to iterate over the raw logs and unpacked data for BurnableSharesIncreased events raised by the IStrategyManager contract.
type IStrategyManagerBurnableSharesIncreasedIterator struct {
	Event *IStrategyManagerBurnableSharesIncreased // Event containing the contract specifics and raw log

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
func (it *IStrategyManagerBurnableSharesIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStrategyManagerBurnableSharesIncreased)
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
		it.Event = new(IStrategyManagerBurnableSharesIncreased)
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
func (it *IStrategyManagerBurnableSharesIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStrategyManagerBurnableSharesIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStrategyManagerBurnableSharesIncreased represents a BurnableSharesIncreased event raised by the IStrategyManager contract.
type IStrategyManagerBurnableSharesIncreased struct {
	Strategy common.Address
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBurnableSharesIncreased is a free log retrieval operation binding the contract event 0xca3e02a4ab7ad3c47a8e36e5a624c30170791726ab720f1babfef21046d953ff.
//
// Solidity: event BurnableSharesIncreased(address strategy, uint256 shares)
func (_IStrategyManager *IStrategyManagerFilterer) FilterBurnableSharesIncreased(opts *bind.FilterOpts) (*IStrategyManagerBurnableSharesIncreasedIterator, error) {

	logs, sub, err := _IStrategyManager.contract.FilterLogs(opts, "BurnableSharesIncreased")
	if err != nil {
		return nil, err
	}
	return &IStrategyManagerBurnableSharesIncreasedIterator{contract: _IStrategyManager.contract, event: "BurnableSharesIncreased", logs: logs, sub: sub}, nil
}

// WatchBurnableSharesIncreased is a free log subscription operation binding the contract event 0xca3e02a4ab7ad3c47a8e36e5a624c30170791726ab720f1babfef21046d953ff.
//
// Solidity: event BurnableSharesIncreased(address strategy, uint256 shares)
func (_IStrategyManager *IStrategyManagerFilterer) WatchBurnableSharesIncreased(opts *bind.WatchOpts, sink chan<- *IStrategyManagerBurnableSharesIncreased) (event.Subscription, error) {

	logs, sub, err := _IStrategyManager.contract.WatchLogs(opts, "BurnableSharesIncreased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStrategyManagerBurnableSharesIncreased)
				if err := _IStrategyManager.contract.UnpackLog(event, "BurnableSharesIncreased", log); err != nil {
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
func (_IStrategyManager *IStrategyManagerFilterer) ParseBurnableSharesIncreased(log types.Log) (*IStrategyManagerBurnableSharesIncreased, error) {
	event := new(IStrategyManagerBurnableSharesIncreased)
	if err := _IStrategyManager.contract.UnpackLog(event, "BurnableSharesIncreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStrategyManagerDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the IStrategyManager contract.
type IStrategyManagerDepositIterator struct {
	Event *IStrategyManagerDeposit // Event containing the contract specifics and raw log

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
func (it *IStrategyManagerDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStrategyManagerDeposit)
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
		it.Event = new(IStrategyManagerDeposit)
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
func (it *IStrategyManagerDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStrategyManagerDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStrategyManagerDeposit represents a Deposit event raised by the IStrategyManager contract.
type IStrategyManagerDeposit struct {
	Staker   common.Address
	Strategy common.Address
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: event Deposit(address staker, address strategy, uint256 shares)
func (_IStrategyManager *IStrategyManagerFilterer) FilterDeposit(opts *bind.FilterOpts) (*IStrategyManagerDepositIterator, error) {

	logs, sub, err := _IStrategyManager.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &IStrategyManagerDepositIterator{contract: _IStrategyManager.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: event Deposit(address staker, address strategy, uint256 shares)
func (_IStrategyManager *IStrategyManagerFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *IStrategyManagerDeposit) (event.Subscription, error) {

	logs, sub, err := _IStrategyManager.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStrategyManagerDeposit)
				if err := _IStrategyManager.contract.UnpackLog(event, "Deposit", log); err != nil {
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
func (_IStrategyManager *IStrategyManagerFilterer) ParseDeposit(log types.Log) (*IStrategyManagerDeposit, error) {
	event := new(IStrategyManagerDeposit)
	if err := _IStrategyManager.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStrategyManagerStrategyAddedToDepositWhitelistIterator is returned from FilterStrategyAddedToDepositWhitelist and is used to iterate over the raw logs and unpacked data for StrategyAddedToDepositWhitelist events raised by the IStrategyManager contract.
type IStrategyManagerStrategyAddedToDepositWhitelistIterator struct {
	Event *IStrategyManagerStrategyAddedToDepositWhitelist // Event containing the contract specifics and raw log

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
func (it *IStrategyManagerStrategyAddedToDepositWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStrategyManagerStrategyAddedToDepositWhitelist)
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
		it.Event = new(IStrategyManagerStrategyAddedToDepositWhitelist)
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
func (it *IStrategyManagerStrategyAddedToDepositWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStrategyManagerStrategyAddedToDepositWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStrategyManagerStrategyAddedToDepositWhitelist represents a StrategyAddedToDepositWhitelist event raised by the IStrategyManager contract.
type IStrategyManagerStrategyAddedToDepositWhitelist struct {
	Strategy common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStrategyAddedToDepositWhitelist is a free log retrieval operation binding the contract event 0x0c35b17d91c96eb2751cd456e1252f42a386e524ef9ff26ecc9950859fdc04fe.
//
// Solidity: event StrategyAddedToDepositWhitelist(address strategy)
func (_IStrategyManager *IStrategyManagerFilterer) FilterStrategyAddedToDepositWhitelist(opts *bind.FilterOpts) (*IStrategyManagerStrategyAddedToDepositWhitelistIterator, error) {

	logs, sub, err := _IStrategyManager.contract.FilterLogs(opts, "StrategyAddedToDepositWhitelist")
	if err != nil {
		return nil, err
	}
	return &IStrategyManagerStrategyAddedToDepositWhitelistIterator{contract: _IStrategyManager.contract, event: "StrategyAddedToDepositWhitelist", logs: logs, sub: sub}, nil
}

// WatchStrategyAddedToDepositWhitelist is a free log subscription operation binding the contract event 0x0c35b17d91c96eb2751cd456e1252f42a386e524ef9ff26ecc9950859fdc04fe.
//
// Solidity: event StrategyAddedToDepositWhitelist(address strategy)
func (_IStrategyManager *IStrategyManagerFilterer) WatchStrategyAddedToDepositWhitelist(opts *bind.WatchOpts, sink chan<- *IStrategyManagerStrategyAddedToDepositWhitelist) (event.Subscription, error) {

	logs, sub, err := _IStrategyManager.contract.WatchLogs(opts, "StrategyAddedToDepositWhitelist")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStrategyManagerStrategyAddedToDepositWhitelist)
				if err := _IStrategyManager.contract.UnpackLog(event, "StrategyAddedToDepositWhitelist", log); err != nil {
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
func (_IStrategyManager *IStrategyManagerFilterer) ParseStrategyAddedToDepositWhitelist(log types.Log) (*IStrategyManagerStrategyAddedToDepositWhitelist, error) {
	event := new(IStrategyManagerStrategyAddedToDepositWhitelist)
	if err := _IStrategyManager.contract.UnpackLog(event, "StrategyAddedToDepositWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStrategyManagerStrategyRemovedFromDepositWhitelistIterator is returned from FilterStrategyRemovedFromDepositWhitelist and is used to iterate over the raw logs and unpacked data for StrategyRemovedFromDepositWhitelist events raised by the IStrategyManager contract.
type IStrategyManagerStrategyRemovedFromDepositWhitelistIterator struct {
	Event *IStrategyManagerStrategyRemovedFromDepositWhitelist // Event containing the contract specifics and raw log

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
func (it *IStrategyManagerStrategyRemovedFromDepositWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStrategyManagerStrategyRemovedFromDepositWhitelist)
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
		it.Event = new(IStrategyManagerStrategyRemovedFromDepositWhitelist)
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
func (it *IStrategyManagerStrategyRemovedFromDepositWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStrategyManagerStrategyRemovedFromDepositWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStrategyManagerStrategyRemovedFromDepositWhitelist represents a StrategyRemovedFromDepositWhitelist event raised by the IStrategyManager contract.
type IStrategyManagerStrategyRemovedFromDepositWhitelist struct {
	Strategy common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStrategyRemovedFromDepositWhitelist is a free log retrieval operation binding the contract event 0x4074413b4b443e4e58019f2855a8765113358c7c72e39509c6af45fc0f5ba030.
//
// Solidity: event StrategyRemovedFromDepositWhitelist(address strategy)
func (_IStrategyManager *IStrategyManagerFilterer) FilterStrategyRemovedFromDepositWhitelist(opts *bind.FilterOpts) (*IStrategyManagerStrategyRemovedFromDepositWhitelistIterator, error) {

	logs, sub, err := _IStrategyManager.contract.FilterLogs(opts, "StrategyRemovedFromDepositWhitelist")
	if err != nil {
		return nil, err
	}
	return &IStrategyManagerStrategyRemovedFromDepositWhitelistIterator{contract: _IStrategyManager.contract, event: "StrategyRemovedFromDepositWhitelist", logs: logs, sub: sub}, nil
}

// WatchStrategyRemovedFromDepositWhitelist is a free log subscription operation binding the contract event 0x4074413b4b443e4e58019f2855a8765113358c7c72e39509c6af45fc0f5ba030.
//
// Solidity: event StrategyRemovedFromDepositWhitelist(address strategy)
func (_IStrategyManager *IStrategyManagerFilterer) WatchStrategyRemovedFromDepositWhitelist(opts *bind.WatchOpts, sink chan<- *IStrategyManagerStrategyRemovedFromDepositWhitelist) (event.Subscription, error) {

	logs, sub, err := _IStrategyManager.contract.WatchLogs(opts, "StrategyRemovedFromDepositWhitelist")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStrategyManagerStrategyRemovedFromDepositWhitelist)
				if err := _IStrategyManager.contract.UnpackLog(event, "StrategyRemovedFromDepositWhitelist", log); err != nil {
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
func (_IStrategyManager *IStrategyManagerFilterer) ParseStrategyRemovedFromDepositWhitelist(log types.Log) (*IStrategyManagerStrategyRemovedFromDepositWhitelist, error) {
	event := new(IStrategyManagerStrategyRemovedFromDepositWhitelist)
	if err := _IStrategyManager.contract.UnpackLog(event, "StrategyRemovedFromDepositWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStrategyManagerStrategyWhitelisterChangedIterator is returned from FilterStrategyWhitelisterChanged and is used to iterate over the raw logs and unpacked data for StrategyWhitelisterChanged events raised by the IStrategyManager contract.
type IStrategyManagerStrategyWhitelisterChangedIterator struct {
	Event *IStrategyManagerStrategyWhitelisterChanged // Event containing the contract specifics and raw log

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
func (it *IStrategyManagerStrategyWhitelisterChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStrategyManagerStrategyWhitelisterChanged)
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
		it.Event = new(IStrategyManagerStrategyWhitelisterChanged)
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
func (it *IStrategyManagerStrategyWhitelisterChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStrategyManagerStrategyWhitelisterChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStrategyManagerStrategyWhitelisterChanged represents a StrategyWhitelisterChanged event raised by the IStrategyManager contract.
type IStrategyManagerStrategyWhitelisterChanged struct {
	PreviousAddress common.Address
	NewAddress      common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterStrategyWhitelisterChanged is a free log retrieval operation binding the contract event 0x4264275e593955ff9d6146a51a4525f6ddace2e81db9391abcc9d1ca48047d29.
//
// Solidity: event StrategyWhitelisterChanged(address previousAddress, address newAddress)
func (_IStrategyManager *IStrategyManagerFilterer) FilterStrategyWhitelisterChanged(opts *bind.FilterOpts) (*IStrategyManagerStrategyWhitelisterChangedIterator, error) {

	logs, sub, err := _IStrategyManager.contract.FilterLogs(opts, "StrategyWhitelisterChanged")
	if err != nil {
		return nil, err
	}
	return &IStrategyManagerStrategyWhitelisterChangedIterator{contract: _IStrategyManager.contract, event: "StrategyWhitelisterChanged", logs: logs, sub: sub}, nil
}

// WatchStrategyWhitelisterChanged is a free log subscription operation binding the contract event 0x4264275e593955ff9d6146a51a4525f6ddace2e81db9391abcc9d1ca48047d29.
//
// Solidity: event StrategyWhitelisterChanged(address previousAddress, address newAddress)
func (_IStrategyManager *IStrategyManagerFilterer) WatchStrategyWhitelisterChanged(opts *bind.WatchOpts, sink chan<- *IStrategyManagerStrategyWhitelisterChanged) (event.Subscription, error) {

	logs, sub, err := _IStrategyManager.contract.WatchLogs(opts, "StrategyWhitelisterChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStrategyManagerStrategyWhitelisterChanged)
				if err := _IStrategyManager.contract.UnpackLog(event, "StrategyWhitelisterChanged", log); err != nil {
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
func (_IStrategyManager *IStrategyManagerFilterer) ParseStrategyWhitelisterChanged(log types.Log) (*IStrategyManagerStrategyWhitelisterChanged, error) {
	event := new(IStrategyManagerStrategyWhitelisterChanged)
	if err := _IStrategyManager.contract.UnpackLog(event, "StrategyWhitelisterChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
