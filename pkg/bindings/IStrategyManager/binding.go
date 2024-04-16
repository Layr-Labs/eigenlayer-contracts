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
	ABI: "[{\"type\":\"function\",\"name\":\"addShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addStrategiesToDepositWhitelist\",\"inputs\":[{\"name\":\"strategiesToWhitelist\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"thirdPartyTransfersForbiddenValues\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"depositIntoStrategy\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositIntoStrategyWithSignature\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eigenPodManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEigenPodManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDeposits\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeStrategiesFromDepositWhitelist\",\"inputs\":[{\"name\":\"strategiesToRemoveFromWhitelist\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setThirdPartyTransfersForbidden\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slasher\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractISlasher\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakerStrategyListLength\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakerStrategyShares\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"strategyIsWhitelistedForDeposit\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"strategyWhitelister\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"thirdPartyTransfersForbidden\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawSharesAsTokens\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Deposit\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIERC20\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyAddedToDepositWhitelist\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyRemovedFromDepositWhitelist\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyWhitelisterChanged\",\"inputs\":[{\"name\":\"previousAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedThirdPartyTransfersForbidden\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"value\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false}]",
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

// EigenPodManager is a free data retrieval call binding the contract method 0x4665bcda.
//
// Solidity: function eigenPodManager() view returns(address)
func (_IStrategyManager *IStrategyManagerCaller) EigenPodManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IStrategyManager.contract.Call(opts, &out, "eigenPodManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EigenPodManager is a free data retrieval call binding the contract method 0x4665bcda.
//
// Solidity: function eigenPodManager() view returns(address)
func (_IStrategyManager *IStrategyManagerSession) EigenPodManager() (common.Address, error) {
	return _IStrategyManager.Contract.EigenPodManager(&_IStrategyManager.CallOpts)
}

// EigenPodManager is a free data retrieval call binding the contract method 0x4665bcda.
//
// Solidity: function eigenPodManager() view returns(address)
func (_IStrategyManager *IStrategyManagerCallerSession) EigenPodManager() (common.Address, error) {
	return _IStrategyManager.Contract.EigenPodManager(&_IStrategyManager.CallOpts)
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

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_IStrategyManager *IStrategyManagerCaller) Slasher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IStrategyManager.contract.Call(opts, &out, "slasher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_IStrategyManager *IStrategyManagerSession) Slasher() (common.Address, error) {
	return _IStrategyManager.Contract.Slasher(&_IStrategyManager.CallOpts)
}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_IStrategyManager *IStrategyManagerCallerSession) Slasher() (common.Address, error) {
	return _IStrategyManager.Contract.Slasher(&_IStrategyManager.CallOpts)
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

// StakerStrategyShares is a free data retrieval call binding the contract method 0x7a7e0d92.
//
// Solidity: function stakerStrategyShares(address user, address strategy) view returns(uint256 shares)
func (_IStrategyManager *IStrategyManagerCaller) StakerStrategyShares(opts *bind.CallOpts, user common.Address, strategy common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IStrategyManager.contract.Call(opts, &out, "stakerStrategyShares", user, strategy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerStrategyShares is a free data retrieval call binding the contract method 0x7a7e0d92.
//
// Solidity: function stakerStrategyShares(address user, address strategy) view returns(uint256 shares)
func (_IStrategyManager *IStrategyManagerSession) StakerStrategyShares(user common.Address, strategy common.Address) (*big.Int, error) {
	return _IStrategyManager.Contract.StakerStrategyShares(&_IStrategyManager.CallOpts, user, strategy)
}

// StakerStrategyShares is a free data retrieval call binding the contract method 0x7a7e0d92.
//
// Solidity: function stakerStrategyShares(address user, address strategy) view returns(uint256 shares)
func (_IStrategyManager *IStrategyManagerCallerSession) StakerStrategyShares(user common.Address, strategy common.Address) (*big.Int, error) {
	return _IStrategyManager.Contract.StakerStrategyShares(&_IStrategyManager.CallOpts, user, strategy)
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

// ThirdPartyTransfersForbidden is a free data retrieval call binding the contract method 0x9b4da03d.
//
// Solidity: function thirdPartyTransfersForbidden(address strategy) view returns(bool)
func (_IStrategyManager *IStrategyManagerCaller) ThirdPartyTransfersForbidden(opts *bind.CallOpts, strategy common.Address) (bool, error) {
	var out []interface{}
	err := _IStrategyManager.contract.Call(opts, &out, "thirdPartyTransfersForbidden", strategy)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ThirdPartyTransfersForbidden is a free data retrieval call binding the contract method 0x9b4da03d.
//
// Solidity: function thirdPartyTransfersForbidden(address strategy) view returns(bool)
func (_IStrategyManager *IStrategyManagerSession) ThirdPartyTransfersForbidden(strategy common.Address) (bool, error) {
	return _IStrategyManager.Contract.ThirdPartyTransfersForbidden(&_IStrategyManager.CallOpts, strategy)
}

// ThirdPartyTransfersForbidden is a free data retrieval call binding the contract method 0x9b4da03d.
//
// Solidity: function thirdPartyTransfersForbidden(address strategy) view returns(bool)
func (_IStrategyManager *IStrategyManagerCallerSession) ThirdPartyTransfersForbidden(strategy common.Address) (bool, error) {
	return _IStrategyManager.Contract.ThirdPartyTransfersForbidden(&_IStrategyManager.CallOpts, strategy)
}

// AddShares is a paid mutator transaction binding the contract method 0xc4623ea1.
//
// Solidity: function addShares(address staker, address token, address strategy, uint256 shares) returns()
func (_IStrategyManager *IStrategyManagerTransactor) AddShares(opts *bind.TransactOpts, staker common.Address, token common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error) {
	return _IStrategyManager.contract.Transact(opts, "addShares", staker, token, strategy, shares)
}

// AddShares is a paid mutator transaction binding the contract method 0xc4623ea1.
//
// Solidity: function addShares(address staker, address token, address strategy, uint256 shares) returns()
func (_IStrategyManager *IStrategyManagerSession) AddShares(staker common.Address, token common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error) {
	return _IStrategyManager.Contract.AddShares(&_IStrategyManager.TransactOpts, staker, token, strategy, shares)
}

// AddShares is a paid mutator transaction binding the contract method 0xc4623ea1.
//
// Solidity: function addShares(address staker, address token, address strategy, uint256 shares) returns()
func (_IStrategyManager *IStrategyManagerTransactorSession) AddShares(staker common.Address, token common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error) {
	return _IStrategyManager.Contract.AddShares(&_IStrategyManager.TransactOpts, staker, token, strategy, shares)
}

// AddStrategiesToDepositWhitelist is a paid mutator transaction binding the contract method 0xdf5b3547.
//
// Solidity: function addStrategiesToDepositWhitelist(address[] strategiesToWhitelist, bool[] thirdPartyTransfersForbiddenValues) returns()
func (_IStrategyManager *IStrategyManagerTransactor) AddStrategiesToDepositWhitelist(opts *bind.TransactOpts, strategiesToWhitelist []common.Address, thirdPartyTransfersForbiddenValues []bool) (*types.Transaction, error) {
	return _IStrategyManager.contract.Transact(opts, "addStrategiesToDepositWhitelist", strategiesToWhitelist, thirdPartyTransfersForbiddenValues)
}

// AddStrategiesToDepositWhitelist is a paid mutator transaction binding the contract method 0xdf5b3547.
//
// Solidity: function addStrategiesToDepositWhitelist(address[] strategiesToWhitelist, bool[] thirdPartyTransfersForbiddenValues) returns()
func (_IStrategyManager *IStrategyManagerSession) AddStrategiesToDepositWhitelist(strategiesToWhitelist []common.Address, thirdPartyTransfersForbiddenValues []bool) (*types.Transaction, error) {
	return _IStrategyManager.Contract.AddStrategiesToDepositWhitelist(&_IStrategyManager.TransactOpts, strategiesToWhitelist, thirdPartyTransfersForbiddenValues)
}

// AddStrategiesToDepositWhitelist is a paid mutator transaction binding the contract method 0xdf5b3547.
//
// Solidity: function addStrategiesToDepositWhitelist(address[] strategiesToWhitelist, bool[] thirdPartyTransfersForbiddenValues) returns()
func (_IStrategyManager *IStrategyManagerTransactorSession) AddStrategiesToDepositWhitelist(strategiesToWhitelist []common.Address, thirdPartyTransfersForbiddenValues []bool) (*types.Transaction, error) {
	return _IStrategyManager.Contract.AddStrategiesToDepositWhitelist(&_IStrategyManager.TransactOpts, strategiesToWhitelist, thirdPartyTransfersForbiddenValues)
}

// DepositIntoStrategy is a paid mutator transaction binding the contract method 0xe7a050aa.
//
// Solidity: function depositIntoStrategy(address strategy, address token, uint256 amount) returns(uint256 shares)
func (_IStrategyManager *IStrategyManagerTransactor) DepositIntoStrategy(opts *bind.TransactOpts, strategy common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IStrategyManager.contract.Transact(opts, "depositIntoStrategy", strategy, token, amount)
}

// DepositIntoStrategy is a paid mutator transaction binding the contract method 0xe7a050aa.
//
// Solidity: function depositIntoStrategy(address strategy, address token, uint256 amount) returns(uint256 shares)
func (_IStrategyManager *IStrategyManagerSession) DepositIntoStrategy(strategy common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IStrategyManager.Contract.DepositIntoStrategy(&_IStrategyManager.TransactOpts, strategy, token, amount)
}

// DepositIntoStrategy is a paid mutator transaction binding the contract method 0xe7a050aa.
//
// Solidity: function depositIntoStrategy(address strategy, address token, uint256 amount) returns(uint256 shares)
func (_IStrategyManager *IStrategyManagerTransactorSession) DepositIntoStrategy(strategy common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IStrategyManager.Contract.DepositIntoStrategy(&_IStrategyManager.TransactOpts, strategy, token, amount)
}

// DepositIntoStrategyWithSignature is a paid mutator transaction binding the contract method 0x32e89ace.
//
// Solidity: function depositIntoStrategyWithSignature(address strategy, address token, uint256 amount, address staker, uint256 expiry, bytes signature) returns(uint256 shares)
func (_IStrategyManager *IStrategyManagerTransactor) DepositIntoStrategyWithSignature(opts *bind.TransactOpts, strategy common.Address, token common.Address, amount *big.Int, staker common.Address, expiry *big.Int, signature []byte) (*types.Transaction, error) {
	return _IStrategyManager.contract.Transact(opts, "depositIntoStrategyWithSignature", strategy, token, amount, staker, expiry, signature)
}

// DepositIntoStrategyWithSignature is a paid mutator transaction binding the contract method 0x32e89ace.
//
// Solidity: function depositIntoStrategyWithSignature(address strategy, address token, uint256 amount, address staker, uint256 expiry, bytes signature) returns(uint256 shares)
func (_IStrategyManager *IStrategyManagerSession) DepositIntoStrategyWithSignature(strategy common.Address, token common.Address, amount *big.Int, staker common.Address, expiry *big.Int, signature []byte) (*types.Transaction, error) {
	return _IStrategyManager.Contract.DepositIntoStrategyWithSignature(&_IStrategyManager.TransactOpts, strategy, token, amount, staker, expiry, signature)
}

// DepositIntoStrategyWithSignature is a paid mutator transaction binding the contract method 0x32e89ace.
//
// Solidity: function depositIntoStrategyWithSignature(address strategy, address token, uint256 amount, address staker, uint256 expiry, bytes signature) returns(uint256 shares)
func (_IStrategyManager *IStrategyManagerTransactorSession) DepositIntoStrategyWithSignature(strategy common.Address, token common.Address, amount *big.Int, staker common.Address, expiry *big.Int, signature []byte) (*types.Transaction, error) {
	return _IStrategyManager.Contract.DepositIntoStrategyWithSignature(&_IStrategyManager.TransactOpts, strategy, token, amount, staker, expiry, signature)
}

// RemoveShares is a paid mutator transaction binding the contract method 0x8c80d4e5.
//
// Solidity: function removeShares(address staker, address strategy, uint256 shares) returns()
func (_IStrategyManager *IStrategyManagerTransactor) RemoveShares(opts *bind.TransactOpts, staker common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error) {
	return _IStrategyManager.contract.Transact(opts, "removeShares", staker, strategy, shares)
}

// RemoveShares is a paid mutator transaction binding the contract method 0x8c80d4e5.
//
// Solidity: function removeShares(address staker, address strategy, uint256 shares) returns()
func (_IStrategyManager *IStrategyManagerSession) RemoveShares(staker common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error) {
	return _IStrategyManager.Contract.RemoveShares(&_IStrategyManager.TransactOpts, staker, strategy, shares)
}

// RemoveShares is a paid mutator transaction binding the contract method 0x8c80d4e5.
//
// Solidity: function removeShares(address staker, address strategy, uint256 shares) returns()
func (_IStrategyManager *IStrategyManagerTransactorSession) RemoveShares(staker common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error) {
	return _IStrategyManager.Contract.RemoveShares(&_IStrategyManager.TransactOpts, staker, strategy, shares)
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

// SetThirdPartyTransfersForbidden is a paid mutator transaction binding the contract method 0x4e5a4263.
//
// Solidity: function setThirdPartyTransfersForbidden(address strategy, bool value) returns()
func (_IStrategyManager *IStrategyManagerTransactor) SetThirdPartyTransfersForbidden(opts *bind.TransactOpts, strategy common.Address, value bool) (*types.Transaction, error) {
	return _IStrategyManager.contract.Transact(opts, "setThirdPartyTransfersForbidden", strategy, value)
}

// SetThirdPartyTransfersForbidden is a paid mutator transaction binding the contract method 0x4e5a4263.
//
// Solidity: function setThirdPartyTransfersForbidden(address strategy, bool value) returns()
func (_IStrategyManager *IStrategyManagerSession) SetThirdPartyTransfersForbidden(strategy common.Address, value bool) (*types.Transaction, error) {
	return _IStrategyManager.Contract.SetThirdPartyTransfersForbidden(&_IStrategyManager.TransactOpts, strategy, value)
}

// SetThirdPartyTransfersForbidden is a paid mutator transaction binding the contract method 0x4e5a4263.
//
// Solidity: function setThirdPartyTransfersForbidden(address strategy, bool value) returns()
func (_IStrategyManager *IStrategyManagerTransactorSession) SetThirdPartyTransfersForbidden(strategy common.Address, value bool) (*types.Transaction, error) {
	return _IStrategyManager.Contract.SetThirdPartyTransfersForbidden(&_IStrategyManager.TransactOpts, strategy, value)
}

// WithdrawSharesAsTokens is a paid mutator transaction binding the contract method 0xc608c7f3.
//
// Solidity: function withdrawSharesAsTokens(address recipient, address strategy, uint256 shares, address token) returns()
func (_IStrategyManager *IStrategyManagerTransactor) WithdrawSharesAsTokens(opts *bind.TransactOpts, recipient common.Address, strategy common.Address, shares *big.Int, token common.Address) (*types.Transaction, error) {
	return _IStrategyManager.contract.Transact(opts, "withdrawSharesAsTokens", recipient, strategy, shares, token)
}

// WithdrawSharesAsTokens is a paid mutator transaction binding the contract method 0xc608c7f3.
//
// Solidity: function withdrawSharesAsTokens(address recipient, address strategy, uint256 shares, address token) returns()
func (_IStrategyManager *IStrategyManagerSession) WithdrawSharesAsTokens(recipient common.Address, strategy common.Address, shares *big.Int, token common.Address) (*types.Transaction, error) {
	return _IStrategyManager.Contract.WithdrawSharesAsTokens(&_IStrategyManager.TransactOpts, recipient, strategy, shares, token)
}

// WithdrawSharesAsTokens is a paid mutator transaction binding the contract method 0xc608c7f3.
//
// Solidity: function withdrawSharesAsTokens(address recipient, address strategy, uint256 shares, address token) returns()
func (_IStrategyManager *IStrategyManagerTransactorSession) WithdrawSharesAsTokens(recipient common.Address, strategy common.Address, shares *big.Int, token common.Address) (*types.Transaction, error) {
	return _IStrategyManager.Contract.WithdrawSharesAsTokens(&_IStrategyManager.TransactOpts, recipient, strategy, shares, token)
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
	Token    common.Address
	Strategy common.Address
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x7cfff908a4b583f36430b25d75964c458d8ede8a99bd61be750e97ee1b2f3a96.
//
// Solidity: event Deposit(address staker, address token, address strategy, uint256 shares)
func (_IStrategyManager *IStrategyManagerFilterer) FilterDeposit(opts *bind.FilterOpts) (*IStrategyManagerDepositIterator, error) {

	logs, sub, err := _IStrategyManager.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &IStrategyManagerDepositIterator{contract: _IStrategyManager.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x7cfff908a4b583f36430b25d75964c458d8ede8a99bd61be750e97ee1b2f3a96.
//
// Solidity: event Deposit(address staker, address token, address strategy, uint256 shares)
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

// ParseDeposit is a log parse operation binding the contract event 0x7cfff908a4b583f36430b25d75964c458d8ede8a99bd61be750e97ee1b2f3a96.
//
// Solidity: event Deposit(address staker, address token, address strategy, uint256 shares)
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

// IStrategyManagerUpdatedThirdPartyTransfersForbiddenIterator is returned from FilterUpdatedThirdPartyTransfersForbidden and is used to iterate over the raw logs and unpacked data for UpdatedThirdPartyTransfersForbidden events raised by the IStrategyManager contract.
type IStrategyManagerUpdatedThirdPartyTransfersForbiddenIterator struct {
	Event *IStrategyManagerUpdatedThirdPartyTransfersForbidden // Event containing the contract specifics and raw log

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
func (it *IStrategyManagerUpdatedThirdPartyTransfersForbiddenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStrategyManagerUpdatedThirdPartyTransfersForbidden)
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
		it.Event = new(IStrategyManagerUpdatedThirdPartyTransfersForbidden)
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
func (it *IStrategyManagerUpdatedThirdPartyTransfersForbiddenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStrategyManagerUpdatedThirdPartyTransfersForbiddenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStrategyManagerUpdatedThirdPartyTransfersForbidden represents a UpdatedThirdPartyTransfersForbidden event raised by the IStrategyManager contract.
type IStrategyManagerUpdatedThirdPartyTransfersForbidden struct {
	Strategy common.Address
	Value    bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdatedThirdPartyTransfersForbidden is a free log retrieval operation binding the contract event 0x77d930df4937793473a95024d87a98fd2ccb9e92d3c2463b3dacd65d3e6a5786.
//
// Solidity: event UpdatedThirdPartyTransfersForbidden(address strategy, bool value)
func (_IStrategyManager *IStrategyManagerFilterer) FilterUpdatedThirdPartyTransfersForbidden(opts *bind.FilterOpts) (*IStrategyManagerUpdatedThirdPartyTransfersForbiddenIterator, error) {

	logs, sub, err := _IStrategyManager.contract.FilterLogs(opts, "UpdatedThirdPartyTransfersForbidden")
	if err != nil {
		return nil, err
	}
	return &IStrategyManagerUpdatedThirdPartyTransfersForbiddenIterator{contract: _IStrategyManager.contract, event: "UpdatedThirdPartyTransfersForbidden", logs: logs, sub: sub}, nil
}

// WatchUpdatedThirdPartyTransfersForbidden is a free log subscription operation binding the contract event 0x77d930df4937793473a95024d87a98fd2ccb9e92d3c2463b3dacd65d3e6a5786.
//
// Solidity: event UpdatedThirdPartyTransfersForbidden(address strategy, bool value)
func (_IStrategyManager *IStrategyManagerFilterer) WatchUpdatedThirdPartyTransfersForbidden(opts *bind.WatchOpts, sink chan<- *IStrategyManagerUpdatedThirdPartyTransfersForbidden) (event.Subscription, error) {

	logs, sub, err := _IStrategyManager.contract.WatchLogs(opts, "UpdatedThirdPartyTransfersForbidden")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStrategyManagerUpdatedThirdPartyTransfersForbidden)
				if err := _IStrategyManager.contract.UnpackLog(event, "UpdatedThirdPartyTransfersForbidden", log); err != nil {
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

// ParseUpdatedThirdPartyTransfersForbidden is a log parse operation binding the contract event 0x77d930df4937793473a95024d87a98fd2ccb9e92d3c2463b3dacd65d3e6a5786.
//
// Solidity: event UpdatedThirdPartyTransfersForbidden(address strategy, bool value)
func (_IStrategyManager *IStrategyManagerFilterer) ParseUpdatedThirdPartyTransfersForbidden(log types.Log) (*IStrategyManagerUpdatedThirdPartyTransfersForbidden, error) {
	event := new(IStrategyManagerUpdatedThirdPartyTransfersForbidden)
	if err := _IStrategyManager.contract.UnpackLog(event, "UpdatedThirdPartyTransfersForbidden", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
