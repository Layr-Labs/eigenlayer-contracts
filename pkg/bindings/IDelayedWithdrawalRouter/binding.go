// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IDelayedWithdrawalRouter

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

// IDelayedWithdrawalRouterDelayedWithdrawal is an auto generated low-level Go binding around an user-defined struct.
type IDelayedWithdrawalRouterDelayedWithdrawal struct {
	Amount       *big.Int
	BlockCreated uint32
}

// IDelayedWithdrawalRouterUserDelayedWithdrawals is an auto generated low-level Go binding around an user-defined struct.
type IDelayedWithdrawalRouterUserDelayedWithdrawals struct {
	DelayedWithdrawalsCompleted *big.Int
	DelayedWithdrawals          []IDelayedWithdrawalRouterDelayedWithdrawal
}

// IDelayedWithdrawalRouterMetaData contains all meta data concerning the IDelayedWithdrawalRouter contract.
var IDelayedWithdrawalRouterMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"canClaimDelayedWithdrawal\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimDelayedWithdrawals\",\"inputs\":[{\"name\":\"maxNumberOfWithdrawalsToClaim\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimDelayedWithdrawals\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxNumberOfWithdrawalsToClaim\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createDelayedWithdrawal\",\"inputs\":[{\"name\":\"podOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"getClaimableUserDelayedWithdrawals\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIDelayedWithdrawalRouter.DelayedWithdrawal[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint224\",\"internalType\":\"uint224\"},{\"name\":\"blockCreated\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUserDelayedWithdrawals\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIDelayedWithdrawalRouter.DelayedWithdrawal[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint224\",\"internalType\":\"uint224\"},{\"name\":\"blockCreated\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setWithdrawalDelayBlocks\",\"inputs\":[{\"name\":\"newValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"userDelayedWithdrawalByIndex\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIDelayedWithdrawalRouter.DelayedWithdrawal\",\"components\":[{\"name\":\"amount\",\"type\":\"uint224\",\"internalType\":\"uint224\"},{\"name\":\"blockCreated\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"userWithdrawals\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIDelayedWithdrawalRouter.UserDelayedWithdrawals\",\"components\":[{\"name\":\"delayedWithdrawalsCompleted\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delayedWithdrawals\",\"type\":\"tuple[]\",\"internalType\":\"structIDelayedWithdrawalRouter.DelayedWithdrawal[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint224\",\"internalType\":\"uint224\"},{\"name\":\"blockCreated\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"userWithdrawalsLength\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawalDelayBlocks\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"DelayedWithdrawalCreated\",\"inputs\":[{\"name\":\"podOwner\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DelayedWithdrawalsClaimed\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amountClaimed\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"delayedWithdrawalsCompleted\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawalDelayBlocksSet\",\"inputs\":[{\"name\":\"previousValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// IDelayedWithdrawalRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use IDelayedWithdrawalRouterMetaData.ABI instead.
var IDelayedWithdrawalRouterABI = IDelayedWithdrawalRouterMetaData.ABI

// IDelayedWithdrawalRouter is an auto generated Go binding around an Ethereum contract.
type IDelayedWithdrawalRouter struct {
	IDelayedWithdrawalRouterCaller     // Read-only binding to the contract
	IDelayedWithdrawalRouterTransactor // Write-only binding to the contract
	IDelayedWithdrawalRouterFilterer   // Log filterer for contract events
}

// IDelayedWithdrawalRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IDelayedWithdrawalRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDelayedWithdrawalRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IDelayedWithdrawalRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDelayedWithdrawalRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IDelayedWithdrawalRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDelayedWithdrawalRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IDelayedWithdrawalRouterSession struct {
	Contract     *IDelayedWithdrawalRouter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IDelayedWithdrawalRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IDelayedWithdrawalRouterCallerSession struct {
	Contract *IDelayedWithdrawalRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// IDelayedWithdrawalRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IDelayedWithdrawalRouterTransactorSession struct {
	Contract     *IDelayedWithdrawalRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// IDelayedWithdrawalRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IDelayedWithdrawalRouterRaw struct {
	Contract *IDelayedWithdrawalRouter // Generic contract binding to access the raw methods on
}

// IDelayedWithdrawalRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IDelayedWithdrawalRouterCallerRaw struct {
	Contract *IDelayedWithdrawalRouterCaller // Generic read-only contract binding to access the raw methods on
}

// IDelayedWithdrawalRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IDelayedWithdrawalRouterTransactorRaw struct {
	Contract *IDelayedWithdrawalRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIDelayedWithdrawalRouter creates a new instance of IDelayedWithdrawalRouter, bound to a specific deployed contract.
func NewIDelayedWithdrawalRouter(address common.Address, backend bind.ContractBackend) (*IDelayedWithdrawalRouter, error) {
	contract, err := bindIDelayedWithdrawalRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IDelayedWithdrawalRouter{IDelayedWithdrawalRouterCaller: IDelayedWithdrawalRouterCaller{contract: contract}, IDelayedWithdrawalRouterTransactor: IDelayedWithdrawalRouterTransactor{contract: contract}, IDelayedWithdrawalRouterFilterer: IDelayedWithdrawalRouterFilterer{contract: contract}}, nil
}

// NewIDelayedWithdrawalRouterCaller creates a new read-only instance of IDelayedWithdrawalRouter, bound to a specific deployed contract.
func NewIDelayedWithdrawalRouterCaller(address common.Address, caller bind.ContractCaller) (*IDelayedWithdrawalRouterCaller, error) {
	contract, err := bindIDelayedWithdrawalRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IDelayedWithdrawalRouterCaller{contract: contract}, nil
}

// NewIDelayedWithdrawalRouterTransactor creates a new write-only instance of IDelayedWithdrawalRouter, bound to a specific deployed contract.
func NewIDelayedWithdrawalRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*IDelayedWithdrawalRouterTransactor, error) {
	contract, err := bindIDelayedWithdrawalRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IDelayedWithdrawalRouterTransactor{contract: contract}, nil
}

// NewIDelayedWithdrawalRouterFilterer creates a new log filterer instance of IDelayedWithdrawalRouter, bound to a specific deployed contract.
func NewIDelayedWithdrawalRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*IDelayedWithdrawalRouterFilterer, error) {
	contract, err := bindIDelayedWithdrawalRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IDelayedWithdrawalRouterFilterer{contract: contract}, nil
}

// bindIDelayedWithdrawalRouter binds a generic wrapper to an already deployed contract.
func bindIDelayedWithdrawalRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IDelayedWithdrawalRouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDelayedWithdrawalRouter *IDelayedWithdrawalRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDelayedWithdrawalRouter.Contract.IDelayedWithdrawalRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDelayedWithdrawalRouter *IDelayedWithdrawalRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDelayedWithdrawalRouter.Contract.IDelayedWithdrawalRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDelayedWithdrawalRouter *IDelayedWithdrawalRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDelayedWithdrawalRouter.Contract.IDelayedWithdrawalRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDelayedWithdrawalRouter *IDelayedWithdrawalRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDelayedWithdrawalRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDelayedWithdrawalRouter *IDelayedWithdrawalRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDelayedWithdrawalRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDelayedWithdrawalRouter *IDelayedWithdrawalRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDelayedWithdrawalRouter.Contract.contract.Transact(opts, method, params...)
}

// CanClaimDelayedWithdrawal is a free data retrieval call binding the contract method 0x75608896.
//
// Solidity: function canClaimDelayedWithdrawal(address user, uint256 index) view returns(bool)
func (_IDelayedWithdrawalRouter *IDelayedWithdrawalRouterCaller) CanClaimDelayedWithdrawal(opts *bind.CallOpts, user common.Address, index *big.Int) (bool, error) {
	var out []interface{}
	err := _IDelayedWithdrawalRouter.contract.Call(opts, &out, "canClaimDelayedWithdrawal", user, index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanClaimDelayedWithdrawal is a free data retrieval call binding the contract method 0x75608896.
//
// Solidity: function canClaimDelayedWithdrawal(address user, uint256 index) view returns(bool)
func (_IDelayedWithdrawalRouter *IDelayedWithdrawalRouterSession) CanClaimDelayedWithdrawal(user common.Address, index *big.Int) (bool, error) {
	return _IDelayedWithdrawalRouter.Contract.CanClaimDelayedWithdrawal(&_IDelayedWithdrawalRouter.CallOpts, user, index)
}

// CanClaimDelayedWithdrawal is a free data retrieval call binding the contract method 0x75608896.
//
// Solidity: function canClaimDelayedWithdrawal(address user, uint256 index) view returns(bool)
func (_IDelayedWithdrawalRouter *IDelayedWithdrawalRouterCallerSession) CanClaimDelayedWithdrawal(user common.Address, index *big.Int) (bool, error) {
	return _IDelayedWithdrawalRouter.Contract.CanClaimDelayedWithdrawal(&_IDelayedWithdrawalRouter.CallOpts, user, index)
}

// GetClaimableUserDelayedWithdrawals is a free data retrieval call binding the contract method 0x1f39d87f.
//
// Solidity: function getClaimableUserDelayedWithdrawals(address user) view returns((uint224,uint32)[])
func (_IDelayedWithdrawalRouter *IDelayedWithdrawalRouterCaller) GetClaimableUserDelayedWithdrawals(opts *bind.CallOpts, user common.Address) ([]IDelayedWithdrawalRouterDelayedWithdrawal, error) {
	var out []interface{}
	err := _IDelayedWithdrawalRouter.contract.Call(opts, &out, "getClaimableUserDelayedWithdrawals", user)

	if err != nil {
		return *new([]IDelayedWithdrawalRouterDelayedWithdrawal), err
	}

	out0 := *abi.ConvertType(out[0], new([]IDelayedWithdrawalRouterDelayedWithdrawal)).(*[]IDelayedWithdrawalRouterDelayedWithdrawal)

	return out0, err

}

// GetClaimableUserDelayedWithdrawals is a free data retrieval call binding the contract method 0x1f39d87f.
//
// Solidity: function getClaimableUserDelayedWithdrawals(address user) view returns((uint224,uint32)[])
func (_IDelayedWithdrawalRouter *IDelayedWithdrawalRouterSession) GetClaimableUserDelayedWithdrawals(user common.Address) ([]IDelayedWithdrawalRouterDelayedWithdrawal, error) {
	return _IDelayedWithdrawalRouter.Contract.GetClaimableUserDelayedWithdrawals(&_IDelayedWithdrawalRouter.CallOpts, user)
}

// GetClaimableUserDelayedWithdrawals is a free data retrieval call binding the contract method 0x1f39d87f.
//
// Solidity: function getClaimableUserDelayedWithdrawals(address user) view returns((uint224,uint32)[])
func (_IDelayedWithdrawalRouter *IDelayedWithdrawalRouterCallerSession) GetClaimableUserDelayedWithdrawals(user common.Address) ([]IDelayedWithdrawalRouterDelayedWithdrawal, error) {
	return _IDelayedWithdrawalRouter.Contract.GetClaimableUserDelayedWithdrawals(&_IDelayedWithdrawalRouter.CallOpts, user)
}

// GetUserDelayedWithdrawals is a free data retrieval call binding the contract method 0x3e1de008.
//
// Solidity: function getUserDelayedWithdrawals(address user) view returns((uint224,uint32)[])
func (_IDelayedWithdrawalRouter *IDelayedWithdrawalRouterCaller) GetUserDelayedWithdrawals(opts *bind.CallOpts, user common.Address) ([]IDelayedWithdrawalRouterDelayedWithdrawal, error) {
	var out []interface{}
	err := _IDelayedWithdrawalRouter.contract.Call(opts, &out, "getUserDelayedWithdrawals", user)

	if err != nil {
		return *new([]IDelayedWithdrawalRouterDelayedWithdrawal), err
	}

	out0 := *abi.ConvertType(out[0], new([]IDelayedWithdrawalRouterDelayedWithdrawal)).(*[]IDelayedWithdrawalRouterDelayedWithdrawal)

	return out0, err

}

// GetUserDelayedWithdrawals is a free data retrieval call binding the contract method 0x3e1de008.
//
// Solidity: function getUserDelayedWithdrawals(address user) view returns((uint224,uint32)[])
func (_IDelayedWithdrawalRouter *IDelayedWithdrawalRouterSession) GetUserDelayedWithdrawals(user common.Address) ([]IDelayedWithdrawalRouterDelayedWithdrawal, error) {
	return _IDelayedWithdrawalRouter.Contract.GetUserDelayedWithdrawals(&_IDelayedWithdrawalRouter.CallOpts, user)
}

// GetUserDelayedWithdrawals is a free data retrieval call binding the contract method 0x3e1de008.
//
// Solidity: function getUserDelayedWithdrawals(address user) view returns((uint224,uint32)[])
func (_IDelayedWithdrawalRouter *IDelayedWithdrawalRouterCallerSession) GetUserDelayedWithdrawals(user common.Address) ([]IDelayedWithdrawalRouterDelayedWithdrawal, error) {
	return _IDelayedWithdrawalRouter.Contract.GetUserDelayedWithdrawals(&_IDelayedWithdrawalRouter.CallOpts, user)
}

// UserDelayedWithdrawalByIndex is a free data retrieval call binding the contract method 0x85594e58.
//
// Solidity: function userDelayedWithdrawalByIndex(address user, uint256 index) view returns((uint224,uint32))
func (_IDelayedWithdrawalRouter *IDelayedWithdrawalRouterCaller) UserDelayedWithdrawalByIndex(opts *bind.CallOpts, user common.Address, index *big.Int) (IDelayedWithdrawalRouterDelayedWithdrawal, error) {
	var out []interface{}
	err := _IDelayedWithdrawalRouter.contract.Call(opts, &out, "userDelayedWithdrawalByIndex", user, index)

	if err != nil {
		return *new(IDelayedWithdrawalRouterDelayedWithdrawal), err
	}

	out0 := *abi.ConvertType(out[0], new(IDelayedWithdrawalRouterDelayedWithdrawal)).(*IDelayedWithdrawalRouterDelayedWithdrawal)

	return out0, err

}

// UserDelayedWithdrawalByIndex is a free data retrieval call binding the contract method 0x85594e58.
//
// Solidity: function userDelayedWithdrawalByIndex(address user, uint256 index) view returns((uint224,uint32))
func (_IDelayedWithdrawalRouter *IDelayedWithdrawalRouterSession) UserDelayedWithdrawalByIndex(user common.Address, index *big.Int) (IDelayedWithdrawalRouterDelayedWithdrawal, error) {
	return _IDelayedWithdrawalRouter.Contract.UserDelayedWithdrawalByIndex(&_IDelayedWithdrawalRouter.CallOpts, user, index)
}

// UserDelayedWithdrawalByIndex is a free data retrieval call binding the contract method 0x85594e58.
//
// Solidity: function userDelayedWithdrawalByIndex(address user, uint256 index) view returns((uint224,uint32))
func (_IDelayedWithdrawalRouter *IDelayedWithdrawalRouterCallerSession) UserDelayedWithdrawalByIndex(user common.Address, index *big.Int) (IDelayedWithdrawalRouterDelayedWithdrawal, error) {
	return _IDelayedWithdrawalRouter.Contract.UserDelayedWithdrawalByIndex(&_IDelayedWithdrawalRouter.CallOpts, user, index)
}

// UserWithdrawals is a free data retrieval call binding the contract method 0xecb7cb1b.
//
// Solidity: function userWithdrawals(address user) view returns((uint256,(uint224,uint32)[]))
func (_IDelayedWithdrawalRouter *IDelayedWithdrawalRouterCaller) UserWithdrawals(opts *bind.CallOpts, user common.Address) (IDelayedWithdrawalRouterUserDelayedWithdrawals, error) {
	var out []interface{}
	err := _IDelayedWithdrawalRouter.contract.Call(opts, &out, "userWithdrawals", user)

	if err != nil {
		return *new(IDelayedWithdrawalRouterUserDelayedWithdrawals), err
	}

	out0 := *abi.ConvertType(out[0], new(IDelayedWithdrawalRouterUserDelayedWithdrawals)).(*IDelayedWithdrawalRouterUserDelayedWithdrawals)

	return out0, err

}

// UserWithdrawals is a free data retrieval call binding the contract method 0xecb7cb1b.
//
// Solidity: function userWithdrawals(address user) view returns((uint256,(uint224,uint32)[]))
func (_IDelayedWithdrawalRouter *IDelayedWithdrawalRouterSession) UserWithdrawals(user common.Address) (IDelayedWithdrawalRouterUserDelayedWithdrawals, error) {
	return _IDelayedWithdrawalRouter.Contract.UserWithdrawals(&_IDelayedWithdrawalRouter.CallOpts, user)
}

// UserWithdrawals is a free data retrieval call binding the contract method 0xecb7cb1b.
//
// Solidity: function userWithdrawals(address user) view returns((uint256,(uint224,uint32)[]))
func (_IDelayedWithdrawalRouter *IDelayedWithdrawalRouterCallerSession) UserWithdrawals(user common.Address) (IDelayedWithdrawalRouterUserDelayedWithdrawals, error) {
	return _IDelayedWithdrawalRouter.Contract.UserWithdrawals(&_IDelayedWithdrawalRouter.CallOpts, user)
}

// UserWithdrawalsLength is a free data retrieval call binding the contract method 0xe4f4f887.
//
// Solidity: function userWithdrawalsLength(address user) view returns(uint256)
func (_IDelayedWithdrawalRouter *IDelayedWithdrawalRouterCaller) UserWithdrawalsLength(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IDelayedWithdrawalRouter.contract.Call(opts, &out, "userWithdrawalsLength", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserWithdrawalsLength is a free data retrieval call binding the contract method 0xe4f4f887.
//
// Solidity: function userWithdrawalsLength(address user) view returns(uint256)
func (_IDelayedWithdrawalRouter *IDelayedWithdrawalRouterSession) UserWithdrawalsLength(user common.Address) (*big.Int, error) {
	return _IDelayedWithdrawalRouter.Contract.UserWithdrawalsLength(&_IDelayedWithdrawalRouter.CallOpts, user)
}

// UserWithdrawalsLength is a free data retrieval call binding the contract method 0xe4f4f887.
//
// Solidity: function userWithdrawalsLength(address user) view returns(uint256)
func (_IDelayedWithdrawalRouter *IDelayedWithdrawalRouterCallerSession) UserWithdrawalsLength(user common.Address) (*big.Int, error) {
	return _IDelayedWithdrawalRouter.Contract.UserWithdrawalsLength(&_IDelayedWithdrawalRouter.CallOpts, user)
}

// WithdrawalDelayBlocks is a free data retrieval call binding the contract method 0x50f73e7c.
//
// Solidity: function withdrawalDelayBlocks() view returns(uint256)
func (_IDelayedWithdrawalRouter *IDelayedWithdrawalRouterCaller) WithdrawalDelayBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IDelayedWithdrawalRouter.contract.Call(opts, &out, "withdrawalDelayBlocks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawalDelayBlocks is a free data retrieval call binding the contract method 0x50f73e7c.
//
// Solidity: function withdrawalDelayBlocks() view returns(uint256)
func (_IDelayedWithdrawalRouter *IDelayedWithdrawalRouterSession) WithdrawalDelayBlocks() (*big.Int, error) {
	return _IDelayedWithdrawalRouter.Contract.WithdrawalDelayBlocks(&_IDelayedWithdrawalRouter.CallOpts)
}

// WithdrawalDelayBlocks is a free data retrieval call binding the contract method 0x50f73e7c.
//
// Solidity: function withdrawalDelayBlocks() view returns(uint256)
func (_IDelayedWithdrawalRouter *IDelayedWithdrawalRouterCallerSession) WithdrawalDelayBlocks() (*big.Int, error) {
	return _IDelayedWithdrawalRouter.Contract.WithdrawalDelayBlocks(&_IDelayedWithdrawalRouter.CallOpts)
}

// ClaimDelayedWithdrawals is a paid mutator transaction binding the contract method 0xd44e1b76.
//
// Solidity: function claimDelayedWithdrawals(uint256 maxNumberOfWithdrawalsToClaim) returns()
func (_IDelayedWithdrawalRouter *IDelayedWithdrawalRouterTransactor) ClaimDelayedWithdrawals(opts *bind.TransactOpts, maxNumberOfWithdrawalsToClaim *big.Int) (*types.Transaction, error) {
	return _IDelayedWithdrawalRouter.contract.Transact(opts, "claimDelayedWithdrawals", maxNumberOfWithdrawalsToClaim)
}

// ClaimDelayedWithdrawals is a paid mutator transaction binding the contract method 0xd44e1b76.
//
// Solidity: function claimDelayedWithdrawals(uint256 maxNumberOfWithdrawalsToClaim) returns()
func (_IDelayedWithdrawalRouter *IDelayedWithdrawalRouterSession) ClaimDelayedWithdrawals(maxNumberOfWithdrawalsToClaim *big.Int) (*types.Transaction, error) {
	return _IDelayedWithdrawalRouter.Contract.ClaimDelayedWithdrawals(&_IDelayedWithdrawalRouter.TransactOpts, maxNumberOfWithdrawalsToClaim)
}

// ClaimDelayedWithdrawals is a paid mutator transaction binding the contract method 0xd44e1b76.
//
// Solidity: function claimDelayedWithdrawals(uint256 maxNumberOfWithdrawalsToClaim) returns()
func (_IDelayedWithdrawalRouter *IDelayedWithdrawalRouterTransactorSession) ClaimDelayedWithdrawals(maxNumberOfWithdrawalsToClaim *big.Int) (*types.Transaction, error) {
	return _IDelayedWithdrawalRouter.Contract.ClaimDelayedWithdrawals(&_IDelayedWithdrawalRouter.TransactOpts, maxNumberOfWithdrawalsToClaim)
}

// ClaimDelayedWithdrawals0 is a paid mutator transaction binding the contract method 0xe5db06c0.
//
// Solidity: function claimDelayedWithdrawals(address recipient, uint256 maxNumberOfWithdrawalsToClaim) returns()
func (_IDelayedWithdrawalRouter *IDelayedWithdrawalRouterTransactor) ClaimDelayedWithdrawals0(opts *bind.TransactOpts, recipient common.Address, maxNumberOfWithdrawalsToClaim *big.Int) (*types.Transaction, error) {
	return _IDelayedWithdrawalRouter.contract.Transact(opts, "claimDelayedWithdrawals0", recipient, maxNumberOfWithdrawalsToClaim)
}

// ClaimDelayedWithdrawals0 is a paid mutator transaction binding the contract method 0xe5db06c0.
//
// Solidity: function claimDelayedWithdrawals(address recipient, uint256 maxNumberOfWithdrawalsToClaim) returns()
func (_IDelayedWithdrawalRouter *IDelayedWithdrawalRouterSession) ClaimDelayedWithdrawals0(recipient common.Address, maxNumberOfWithdrawalsToClaim *big.Int) (*types.Transaction, error) {
	return _IDelayedWithdrawalRouter.Contract.ClaimDelayedWithdrawals0(&_IDelayedWithdrawalRouter.TransactOpts, recipient, maxNumberOfWithdrawalsToClaim)
}

// ClaimDelayedWithdrawals0 is a paid mutator transaction binding the contract method 0xe5db06c0.
//
// Solidity: function claimDelayedWithdrawals(address recipient, uint256 maxNumberOfWithdrawalsToClaim) returns()
func (_IDelayedWithdrawalRouter *IDelayedWithdrawalRouterTransactorSession) ClaimDelayedWithdrawals0(recipient common.Address, maxNumberOfWithdrawalsToClaim *big.Int) (*types.Transaction, error) {
	return _IDelayedWithdrawalRouter.Contract.ClaimDelayedWithdrawals0(&_IDelayedWithdrawalRouter.TransactOpts, recipient, maxNumberOfWithdrawalsToClaim)
}

// CreateDelayedWithdrawal is a paid mutator transaction binding the contract method 0xc0db354c.
//
// Solidity: function createDelayedWithdrawal(address podOwner, address recipient) payable returns()
func (_IDelayedWithdrawalRouter *IDelayedWithdrawalRouterTransactor) CreateDelayedWithdrawal(opts *bind.TransactOpts, podOwner common.Address, recipient common.Address) (*types.Transaction, error) {
	return _IDelayedWithdrawalRouter.contract.Transact(opts, "createDelayedWithdrawal", podOwner, recipient)
}

// CreateDelayedWithdrawal is a paid mutator transaction binding the contract method 0xc0db354c.
//
// Solidity: function createDelayedWithdrawal(address podOwner, address recipient) payable returns()
func (_IDelayedWithdrawalRouter *IDelayedWithdrawalRouterSession) CreateDelayedWithdrawal(podOwner common.Address, recipient common.Address) (*types.Transaction, error) {
	return _IDelayedWithdrawalRouter.Contract.CreateDelayedWithdrawal(&_IDelayedWithdrawalRouter.TransactOpts, podOwner, recipient)
}

// CreateDelayedWithdrawal is a paid mutator transaction binding the contract method 0xc0db354c.
//
// Solidity: function createDelayedWithdrawal(address podOwner, address recipient) payable returns()
func (_IDelayedWithdrawalRouter *IDelayedWithdrawalRouterTransactorSession) CreateDelayedWithdrawal(podOwner common.Address, recipient common.Address) (*types.Transaction, error) {
	return _IDelayedWithdrawalRouter.Contract.CreateDelayedWithdrawal(&_IDelayedWithdrawalRouter.TransactOpts, podOwner, recipient)
}

// SetWithdrawalDelayBlocks is a paid mutator transaction binding the contract method 0x4d50f9a4.
//
// Solidity: function setWithdrawalDelayBlocks(uint256 newValue) returns()
func (_IDelayedWithdrawalRouter *IDelayedWithdrawalRouterTransactor) SetWithdrawalDelayBlocks(opts *bind.TransactOpts, newValue *big.Int) (*types.Transaction, error) {
	return _IDelayedWithdrawalRouter.contract.Transact(opts, "setWithdrawalDelayBlocks", newValue)
}

// SetWithdrawalDelayBlocks is a paid mutator transaction binding the contract method 0x4d50f9a4.
//
// Solidity: function setWithdrawalDelayBlocks(uint256 newValue) returns()
func (_IDelayedWithdrawalRouter *IDelayedWithdrawalRouterSession) SetWithdrawalDelayBlocks(newValue *big.Int) (*types.Transaction, error) {
	return _IDelayedWithdrawalRouter.Contract.SetWithdrawalDelayBlocks(&_IDelayedWithdrawalRouter.TransactOpts, newValue)
}

// SetWithdrawalDelayBlocks is a paid mutator transaction binding the contract method 0x4d50f9a4.
//
// Solidity: function setWithdrawalDelayBlocks(uint256 newValue) returns()
func (_IDelayedWithdrawalRouter *IDelayedWithdrawalRouterTransactorSession) SetWithdrawalDelayBlocks(newValue *big.Int) (*types.Transaction, error) {
	return _IDelayedWithdrawalRouter.Contract.SetWithdrawalDelayBlocks(&_IDelayedWithdrawalRouter.TransactOpts, newValue)
}

// IDelayedWithdrawalRouterDelayedWithdrawalCreatedIterator is returned from FilterDelayedWithdrawalCreated and is used to iterate over the raw logs and unpacked data for DelayedWithdrawalCreated events raised by the IDelayedWithdrawalRouter contract.
type IDelayedWithdrawalRouterDelayedWithdrawalCreatedIterator struct {
	Event *IDelayedWithdrawalRouterDelayedWithdrawalCreated // Event containing the contract specifics and raw log

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
func (it *IDelayedWithdrawalRouterDelayedWithdrawalCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDelayedWithdrawalRouterDelayedWithdrawalCreated)
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
		it.Event = new(IDelayedWithdrawalRouterDelayedWithdrawalCreated)
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
func (it *IDelayedWithdrawalRouterDelayedWithdrawalCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDelayedWithdrawalRouterDelayedWithdrawalCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDelayedWithdrawalRouterDelayedWithdrawalCreated represents a DelayedWithdrawalCreated event raised by the IDelayedWithdrawalRouter contract.
type IDelayedWithdrawalRouterDelayedWithdrawalCreated struct {
	PodOwner  common.Address
	Recipient common.Address
	Amount    *big.Int
	Index     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDelayedWithdrawalCreated is a free log retrieval operation binding the contract event 0xb8f1b14c7caf74150801dcc9bc18d575cbeaf5b421943497e409df92c92e0f59.
//
// Solidity: event DelayedWithdrawalCreated(address podOwner, address recipient, uint256 amount, uint256 index)
func (_IDelayedWithdrawalRouter *IDelayedWithdrawalRouterFilterer) FilterDelayedWithdrawalCreated(opts *bind.FilterOpts) (*IDelayedWithdrawalRouterDelayedWithdrawalCreatedIterator, error) {

	logs, sub, err := _IDelayedWithdrawalRouter.contract.FilterLogs(opts, "DelayedWithdrawalCreated")
	if err != nil {
		return nil, err
	}
	return &IDelayedWithdrawalRouterDelayedWithdrawalCreatedIterator{contract: _IDelayedWithdrawalRouter.contract, event: "DelayedWithdrawalCreated", logs: logs, sub: sub}, nil
}

// WatchDelayedWithdrawalCreated is a free log subscription operation binding the contract event 0xb8f1b14c7caf74150801dcc9bc18d575cbeaf5b421943497e409df92c92e0f59.
//
// Solidity: event DelayedWithdrawalCreated(address podOwner, address recipient, uint256 amount, uint256 index)
func (_IDelayedWithdrawalRouter *IDelayedWithdrawalRouterFilterer) WatchDelayedWithdrawalCreated(opts *bind.WatchOpts, sink chan<- *IDelayedWithdrawalRouterDelayedWithdrawalCreated) (event.Subscription, error) {

	logs, sub, err := _IDelayedWithdrawalRouter.contract.WatchLogs(opts, "DelayedWithdrawalCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDelayedWithdrawalRouterDelayedWithdrawalCreated)
				if err := _IDelayedWithdrawalRouter.contract.UnpackLog(event, "DelayedWithdrawalCreated", log); err != nil {
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

// ParseDelayedWithdrawalCreated is a log parse operation binding the contract event 0xb8f1b14c7caf74150801dcc9bc18d575cbeaf5b421943497e409df92c92e0f59.
//
// Solidity: event DelayedWithdrawalCreated(address podOwner, address recipient, uint256 amount, uint256 index)
func (_IDelayedWithdrawalRouter *IDelayedWithdrawalRouterFilterer) ParseDelayedWithdrawalCreated(log types.Log) (*IDelayedWithdrawalRouterDelayedWithdrawalCreated, error) {
	event := new(IDelayedWithdrawalRouterDelayedWithdrawalCreated)
	if err := _IDelayedWithdrawalRouter.contract.UnpackLog(event, "DelayedWithdrawalCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDelayedWithdrawalRouterDelayedWithdrawalsClaimedIterator is returned from FilterDelayedWithdrawalsClaimed and is used to iterate over the raw logs and unpacked data for DelayedWithdrawalsClaimed events raised by the IDelayedWithdrawalRouter contract.
type IDelayedWithdrawalRouterDelayedWithdrawalsClaimedIterator struct {
	Event *IDelayedWithdrawalRouterDelayedWithdrawalsClaimed // Event containing the contract specifics and raw log

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
func (it *IDelayedWithdrawalRouterDelayedWithdrawalsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDelayedWithdrawalRouterDelayedWithdrawalsClaimed)
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
		it.Event = new(IDelayedWithdrawalRouterDelayedWithdrawalsClaimed)
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
func (it *IDelayedWithdrawalRouterDelayedWithdrawalsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDelayedWithdrawalRouterDelayedWithdrawalsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDelayedWithdrawalRouterDelayedWithdrawalsClaimed represents a DelayedWithdrawalsClaimed event raised by the IDelayedWithdrawalRouter contract.
type IDelayedWithdrawalRouterDelayedWithdrawalsClaimed struct {
	Recipient                   common.Address
	AmountClaimed               *big.Int
	DelayedWithdrawalsCompleted *big.Int
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterDelayedWithdrawalsClaimed is a free log retrieval operation binding the contract event 0x6b7151500bd0b5cc211bcc47b3029831b769004df4549e8e1c9a69da05bb0943.
//
// Solidity: event DelayedWithdrawalsClaimed(address recipient, uint256 amountClaimed, uint256 delayedWithdrawalsCompleted)
func (_IDelayedWithdrawalRouter *IDelayedWithdrawalRouterFilterer) FilterDelayedWithdrawalsClaimed(opts *bind.FilterOpts) (*IDelayedWithdrawalRouterDelayedWithdrawalsClaimedIterator, error) {

	logs, sub, err := _IDelayedWithdrawalRouter.contract.FilterLogs(opts, "DelayedWithdrawalsClaimed")
	if err != nil {
		return nil, err
	}
	return &IDelayedWithdrawalRouterDelayedWithdrawalsClaimedIterator{contract: _IDelayedWithdrawalRouter.contract, event: "DelayedWithdrawalsClaimed", logs: logs, sub: sub}, nil
}

// WatchDelayedWithdrawalsClaimed is a free log subscription operation binding the contract event 0x6b7151500bd0b5cc211bcc47b3029831b769004df4549e8e1c9a69da05bb0943.
//
// Solidity: event DelayedWithdrawalsClaimed(address recipient, uint256 amountClaimed, uint256 delayedWithdrawalsCompleted)
func (_IDelayedWithdrawalRouter *IDelayedWithdrawalRouterFilterer) WatchDelayedWithdrawalsClaimed(opts *bind.WatchOpts, sink chan<- *IDelayedWithdrawalRouterDelayedWithdrawalsClaimed) (event.Subscription, error) {

	logs, sub, err := _IDelayedWithdrawalRouter.contract.WatchLogs(opts, "DelayedWithdrawalsClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDelayedWithdrawalRouterDelayedWithdrawalsClaimed)
				if err := _IDelayedWithdrawalRouter.contract.UnpackLog(event, "DelayedWithdrawalsClaimed", log); err != nil {
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

// ParseDelayedWithdrawalsClaimed is a log parse operation binding the contract event 0x6b7151500bd0b5cc211bcc47b3029831b769004df4549e8e1c9a69da05bb0943.
//
// Solidity: event DelayedWithdrawalsClaimed(address recipient, uint256 amountClaimed, uint256 delayedWithdrawalsCompleted)
func (_IDelayedWithdrawalRouter *IDelayedWithdrawalRouterFilterer) ParseDelayedWithdrawalsClaimed(log types.Log) (*IDelayedWithdrawalRouterDelayedWithdrawalsClaimed, error) {
	event := new(IDelayedWithdrawalRouterDelayedWithdrawalsClaimed)
	if err := _IDelayedWithdrawalRouter.contract.UnpackLog(event, "DelayedWithdrawalsClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDelayedWithdrawalRouterWithdrawalDelayBlocksSetIterator is returned from FilterWithdrawalDelayBlocksSet and is used to iterate over the raw logs and unpacked data for WithdrawalDelayBlocksSet events raised by the IDelayedWithdrawalRouter contract.
type IDelayedWithdrawalRouterWithdrawalDelayBlocksSetIterator struct {
	Event *IDelayedWithdrawalRouterWithdrawalDelayBlocksSet // Event containing the contract specifics and raw log

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
func (it *IDelayedWithdrawalRouterWithdrawalDelayBlocksSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDelayedWithdrawalRouterWithdrawalDelayBlocksSet)
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
		it.Event = new(IDelayedWithdrawalRouterWithdrawalDelayBlocksSet)
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
func (it *IDelayedWithdrawalRouterWithdrawalDelayBlocksSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDelayedWithdrawalRouterWithdrawalDelayBlocksSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDelayedWithdrawalRouterWithdrawalDelayBlocksSet represents a WithdrawalDelayBlocksSet event raised by the IDelayedWithdrawalRouter contract.
type IDelayedWithdrawalRouterWithdrawalDelayBlocksSet struct {
	PreviousValue *big.Int
	NewValue      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalDelayBlocksSet is a free log retrieval operation binding the contract event 0x4ffb00400574147429ee377a5633386321e66d45d8b14676014b5fa393e61e9e.
//
// Solidity: event WithdrawalDelayBlocksSet(uint256 previousValue, uint256 newValue)
func (_IDelayedWithdrawalRouter *IDelayedWithdrawalRouterFilterer) FilterWithdrawalDelayBlocksSet(opts *bind.FilterOpts) (*IDelayedWithdrawalRouterWithdrawalDelayBlocksSetIterator, error) {

	logs, sub, err := _IDelayedWithdrawalRouter.contract.FilterLogs(opts, "WithdrawalDelayBlocksSet")
	if err != nil {
		return nil, err
	}
	return &IDelayedWithdrawalRouterWithdrawalDelayBlocksSetIterator{contract: _IDelayedWithdrawalRouter.contract, event: "WithdrawalDelayBlocksSet", logs: logs, sub: sub}, nil
}

// WatchWithdrawalDelayBlocksSet is a free log subscription operation binding the contract event 0x4ffb00400574147429ee377a5633386321e66d45d8b14676014b5fa393e61e9e.
//
// Solidity: event WithdrawalDelayBlocksSet(uint256 previousValue, uint256 newValue)
func (_IDelayedWithdrawalRouter *IDelayedWithdrawalRouterFilterer) WatchWithdrawalDelayBlocksSet(opts *bind.WatchOpts, sink chan<- *IDelayedWithdrawalRouterWithdrawalDelayBlocksSet) (event.Subscription, error) {

	logs, sub, err := _IDelayedWithdrawalRouter.contract.WatchLogs(opts, "WithdrawalDelayBlocksSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDelayedWithdrawalRouterWithdrawalDelayBlocksSet)
				if err := _IDelayedWithdrawalRouter.contract.UnpackLog(event, "WithdrawalDelayBlocksSet", log); err != nil {
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

// ParseWithdrawalDelayBlocksSet is a log parse operation binding the contract event 0x4ffb00400574147429ee377a5633386321e66d45d8b14676014b5fa393e61e9e.
//
// Solidity: event WithdrawalDelayBlocksSet(uint256 previousValue, uint256 newValue)
func (_IDelayedWithdrawalRouter *IDelayedWithdrawalRouterFilterer) ParseWithdrawalDelayBlocksSet(log types.Log) (*IDelayedWithdrawalRouterWithdrawalDelayBlocksSet, error) {
	event := new(IDelayedWithdrawalRouterWithdrawalDelayBlocksSet)
	if err := _IDelayedWithdrawalRouter.contract.UnpackLog(event, "WithdrawalDelayBlocksSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
