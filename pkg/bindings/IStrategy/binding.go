// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IStrategy

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

// IStrategyMetaData contains all meta data concerning the IStrategy contract.
var IStrategyMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"explanation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"shares\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sharesToUnderlying\",\"inputs\":[{\"name\":\"amountShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sharesToUnderlyingView\",\"inputs\":[{\"name\":\"amountShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalShares\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"underlyingToShares\",\"inputs\":[{\"name\":\"amountUnderlying\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"underlyingToSharesView\",\"inputs\":[{\"name\":\"amountUnderlying\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"underlyingToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"userUnderlying\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"userUnderlyingView\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amountShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
}

// IStrategyABI is the input ABI used to generate the binding from.
// Deprecated: Use IStrategyMetaData.ABI instead.
var IStrategyABI = IStrategyMetaData.ABI

// IStrategy is an auto generated Go binding around an Ethereum contract.
type IStrategy struct {
	IStrategyCaller     // Read-only binding to the contract
	IStrategyTransactor // Write-only binding to the contract
	IStrategyFilterer   // Log filterer for contract events
}

// IStrategyCaller is an auto generated read-only Go binding around an Ethereum contract.
type IStrategyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStrategyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IStrategyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStrategyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IStrategyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStrategySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IStrategySession struct {
	Contract     *IStrategy        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IStrategyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IStrategyCallerSession struct {
	Contract *IStrategyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IStrategyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IStrategyTransactorSession struct {
	Contract     *IStrategyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IStrategyRaw is an auto generated low-level Go binding around an Ethereum contract.
type IStrategyRaw struct {
	Contract *IStrategy // Generic contract binding to access the raw methods on
}

// IStrategyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IStrategyCallerRaw struct {
	Contract *IStrategyCaller // Generic read-only contract binding to access the raw methods on
}

// IStrategyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IStrategyTransactorRaw struct {
	Contract *IStrategyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIStrategy creates a new instance of IStrategy, bound to a specific deployed contract.
func NewIStrategy(address common.Address, backend bind.ContractBackend) (*IStrategy, error) {
	contract, err := bindIStrategy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IStrategy{IStrategyCaller: IStrategyCaller{contract: contract}, IStrategyTransactor: IStrategyTransactor{contract: contract}, IStrategyFilterer: IStrategyFilterer{contract: contract}}, nil
}

// NewIStrategyCaller creates a new read-only instance of IStrategy, bound to a specific deployed contract.
func NewIStrategyCaller(address common.Address, caller bind.ContractCaller) (*IStrategyCaller, error) {
	contract, err := bindIStrategy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IStrategyCaller{contract: contract}, nil
}

// NewIStrategyTransactor creates a new write-only instance of IStrategy, bound to a specific deployed contract.
func NewIStrategyTransactor(address common.Address, transactor bind.ContractTransactor) (*IStrategyTransactor, error) {
	contract, err := bindIStrategy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IStrategyTransactor{contract: contract}, nil
}

// NewIStrategyFilterer creates a new log filterer instance of IStrategy, bound to a specific deployed contract.
func NewIStrategyFilterer(address common.Address, filterer bind.ContractFilterer) (*IStrategyFilterer, error) {
	contract, err := bindIStrategy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IStrategyFilterer{contract: contract}, nil
}

// bindIStrategy binds a generic wrapper to an already deployed contract.
func bindIStrategy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IStrategyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStrategy *IStrategyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStrategy.Contract.IStrategyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStrategy *IStrategyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStrategy.Contract.IStrategyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStrategy *IStrategyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStrategy.Contract.IStrategyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStrategy *IStrategyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStrategy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStrategy *IStrategyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStrategy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStrategy *IStrategyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStrategy.Contract.contract.Transact(opts, method, params...)
}

// Explanation is a free data retrieval call binding the contract method 0xab5921e1.
//
// Solidity: function explanation() view returns(string)
func (_IStrategy *IStrategyCaller) Explanation(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IStrategy.contract.Call(opts, &out, "explanation")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Explanation is a free data retrieval call binding the contract method 0xab5921e1.
//
// Solidity: function explanation() view returns(string)
func (_IStrategy *IStrategySession) Explanation() (string, error) {
	return _IStrategy.Contract.Explanation(&_IStrategy.CallOpts)
}

// Explanation is a free data retrieval call binding the contract method 0xab5921e1.
//
// Solidity: function explanation() view returns(string)
func (_IStrategy *IStrategyCallerSession) Explanation() (string, error) {
	return _IStrategy.Contract.Explanation(&_IStrategy.CallOpts)
}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares(address user) view returns(uint256)
func (_IStrategy *IStrategyCaller) Shares(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IStrategy.contract.Call(opts, &out, "shares", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares(address user) view returns(uint256)
func (_IStrategy *IStrategySession) Shares(user common.Address) (*big.Int, error) {
	return _IStrategy.Contract.Shares(&_IStrategy.CallOpts, user)
}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares(address user) view returns(uint256)
func (_IStrategy *IStrategyCallerSession) Shares(user common.Address) (*big.Int, error) {
	return _IStrategy.Contract.Shares(&_IStrategy.CallOpts, user)
}

// SharesToUnderlyingView is a free data retrieval call binding the contract method 0x7a8b2637.
//
// Solidity: function sharesToUnderlyingView(uint256 amountShares) view returns(uint256)
func (_IStrategy *IStrategyCaller) SharesToUnderlyingView(opts *bind.CallOpts, amountShares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IStrategy.contract.Call(opts, &out, "sharesToUnderlyingView", amountShares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SharesToUnderlyingView is a free data retrieval call binding the contract method 0x7a8b2637.
//
// Solidity: function sharesToUnderlyingView(uint256 amountShares) view returns(uint256)
func (_IStrategy *IStrategySession) SharesToUnderlyingView(amountShares *big.Int) (*big.Int, error) {
	return _IStrategy.Contract.SharesToUnderlyingView(&_IStrategy.CallOpts, amountShares)
}

// SharesToUnderlyingView is a free data retrieval call binding the contract method 0x7a8b2637.
//
// Solidity: function sharesToUnderlyingView(uint256 amountShares) view returns(uint256)
func (_IStrategy *IStrategyCallerSession) SharesToUnderlyingView(amountShares *big.Int) (*big.Int, error) {
	return _IStrategy.Contract.SharesToUnderlyingView(&_IStrategy.CallOpts, amountShares)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_IStrategy *IStrategyCaller) TotalShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IStrategy.contract.Call(opts, &out, "totalShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_IStrategy *IStrategySession) TotalShares() (*big.Int, error) {
	return _IStrategy.Contract.TotalShares(&_IStrategy.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_IStrategy *IStrategyCallerSession) TotalShares() (*big.Int, error) {
	return _IStrategy.Contract.TotalShares(&_IStrategy.CallOpts)
}

// UnderlyingToSharesView is a free data retrieval call binding the contract method 0xe3dae51c.
//
// Solidity: function underlyingToSharesView(uint256 amountUnderlying) view returns(uint256)
func (_IStrategy *IStrategyCaller) UnderlyingToSharesView(opts *bind.CallOpts, amountUnderlying *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IStrategy.contract.Call(opts, &out, "underlyingToSharesView", amountUnderlying)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnderlyingToSharesView is a free data retrieval call binding the contract method 0xe3dae51c.
//
// Solidity: function underlyingToSharesView(uint256 amountUnderlying) view returns(uint256)
func (_IStrategy *IStrategySession) UnderlyingToSharesView(amountUnderlying *big.Int) (*big.Int, error) {
	return _IStrategy.Contract.UnderlyingToSharesView(&_IStrategy.CallOpts, amountUnderlying)
}

// UnderlyingToSharesView is a free data retrieval call binding the contract method 0xe3dae51c.
//
// Solidity: function underlyingToSharesView(uint256 amountUnderlying) view returns(uint256)
func (_IStrategy *IStrategyCallerSession) UnderlyingToSharesView(amountUnderlying *big.Int) (*big.Int, error) {
	return _IStrategy.Contract.UnderlyingToSharesView(&_IStrategy.CallOpts, amountUnderlying)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_IStrategy *IStrategyCaller) UnderlyingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IStrategy.contract.Call(opts, &out, "underlyingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_IStrategy *IStrategySession) UnderlyingToken() (common.Address, error) {
	return _IStrategy.Contract.UnderlyingToken(&_IStrategy.CallOpts)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_IStrategy *IStrategyCallerSession) UnderlyingToken() (common.Address, error) {
	return _IStrategy.Contract.UnderlyingToken(&_IStrategy.CallOpts)
}

// UserUnderlyingView is a free data retrieval call binding the contract method 0x553ca5f8.
//
// Solidity: function userUnderlyingView(address user) view returns(uint256)
func (_IStrategy *IStrategyCaller) UserUnderlyingView(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IStrategy.contract.Call(opts, &out, "userUnderlyingView", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserUnderlyingView is a free data retrieval call binding the contract method 0x553ca5f8.
//
// Solidity: function userUnderlyingView(address user) view returns(uint256)
func (_IStrategy *IStrategySession) UserUnderlyingView(user common.Address) (*big.Int, error) {
	return _IStrategy.Contract.UserUnderlyingView(&_IStrategy.CallOpts, user)
}

// UserUnderlyingView is a free data retrieval call binding the contract method 0x553ca5f8.
//
// Solidity: function userUnderlyingView(address user) view returns(uint256)
func (_IStrategy *IStrategyCallerSession) UserUnderlyingView(user common.Address) (*big.Int, error) {
	return _IStrategy.Contract.UserUnderlyingView(&_IStrategy.CallOpts, user)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns(uint256)
func (_IStrategy *IStrategyTransactor) Deposit(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IStrategy.contract.Transact(opts, "deposit", token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns(uint256)
func (_IStrategy *IStrategySession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IStrategy.Contract.Deposit(&_IStrategy.TransactOpts, token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns(uint256)
func (_IStrategy *IStrategyTransactorSession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IStrategy.Contract.Deposit(&_IStrategy.TransactOpts, token, amount)
}

// SharesToUnderlying is a paid mutator transaction binding the contract method 0xf3e73875.
//
// Solidity: function sharesToUnderlying(uint256 amountShares) returns(uint256)
func (_IStrategy *IStrategyTransactor) SharesToUnderlying(opts *bind.TransactOpts, amountShares *big.Int) (*types.Transaction, error) {
	return _IStrategy.contract.Transact(opts, "sharesToUnderlying", amountShares)
}

// SharesToUnderlying is a paid mutator transaction binding the contract method 0xf3e73875.
//
// Solidity: function sharesToUnderlying(uint256 amountShares) returns(uint256)
func (_IStrategy *IStrategySession) SharesToUnderlying(amountShares *big.Int) (*types.Transaction, error) {
	return _IStrategy.Contract.SharesToUnderlying(&_IStrategy.TransactOpts, amountShares)
}

// SharesToUnderlying is a paid mutator transaction binding the contract method 0xf3e73875.
//
// Solidity: function sharesToUnderlying(uint256 amountShares) returns(uint256)
func (_IStrategy *IStrategyTransactorSession) SharesToUnderlying(amountShares *big.Int) (*types.Transaction, error) {
	return _IStrategy.Contract.SharesToUnderlying(&_IStrategy.TransactOpts, amountShares)
}

// UnderlyingToShares is a paid mutator transaction binding the contract method 0x8c871019.
//
// Solidity: function underlyingToShares(uint256 amountUnderlying) returns(uint256)
func (_IStrategy *IStrategyTransactor) UnderlyingToShares(opts *bind.TransactOpts, amountUnderlying *big.Int) (*types.Transaction, error) {
	return _IStrategy.contract.Transact(opts, "underlyingToShares", amountUnderlying)
}

// UnderlyingToShares is a paid mutator transaction binding the contract method 0x8c871019.
//
// Solidity: function underlyingToShares(uint256 amountUnderlying) returns(uint256)
func (_IStrategy *IStrategySession) UnderlyingToShares(amountUnderlying *big.Int) (*types.Transaction, error) {
	return _IStrategy.Contract.UnderlyingToShares(&_IStrategy.TransactOpts, amountUnderlying)
}

// UnderlyingToShares is a paid mutator transaction binding the contract method 0x8c871019.
//
// Solidity: function underlyingToShares(uint256 amountUnderlying) returns(uint256)
func (_IStrategy *IStrategyTransactorSession) UnderlyingToShares(amountUnderlying *big.Int) (*types.Transaction, error) {
	return _IStrategy.Contract.UnderlyingToShares(&_IStrategy.TransactOpts, amountUnderlying)
}

// UserUnderlying is a paid mutator transaction binding the contract method 0x8f6a6240.
//
// Solidity: function userUnderlying(address user) returns(uint256)
func (_IStrategy *IStrategyTransactor) UserUnderlying(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _IStrategy.contract.Transact(opts, "userUnderlying", user)
}

// UserUnderlying is a paid mutator transaction binding the contract method 0x8f6a6240.
//
// Solidity: function userUnderlying(address user) returns(uint256)
func (_IStrategy *IStrategySession) UserUnderlying(user common.Address) (*types.Transaction, error) {
	return _IStrategy.Contract.UserUnderlying(&_IStrategy.TransactOpts, user)
}

// UserUnderlying is a paid mutator transaction binding the contract method 0x8f6a6240.
//
// Solidity: function userUnderlying(address user) returns(uint256)
func (_IStrategy *IStrategyTransactorSession) UserUnderlying(user common.Address) (*types.Transaction, error) {
	return _IStrategy.Contract.UserUnderlying(&_IStrategy.TransactOpts, user)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address recipient, address token, uint256 amountShares) returns()
func (_IStrategy *IStrategyTransactor) Withdraw(opts *bind.TransactOpts, recipient common.Address, token common.Address, amountShares *big.Int) (*types.Transaction, error) {
	return _IStrategy.contract.Transact(opts, "withdraw", recipient, token, amountShares)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address recipient, address token, uint256 amountShares) returns()
func (_IStrategy *IStrategySession) Withdraw(recipient common.Address, token common.Address, amountShares *big.Int) (*types.Transaction, error) {
	return _IStrategy.Contract.Withdraw(&_IStrategy.TransactOpts, recipient, token, amountShares)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address recipient, address token, uint256 amountShares) returns()
func (_IStrategy *IStrategyTransactorSession) Withdraw(recipient common.Address, token common.Address, amountShares *big.Int) (*types.Transaction, error) {
	return _IStrategy.Contract.Withdraw(&_IStrategy.TransactOpts, recipient, token, amountShares)
}
