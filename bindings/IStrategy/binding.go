// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractIStrategy

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

// ContractIStrategyMetaData contains all meta data concerning the ContractIStrategy contract.
var ContractIStrategyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"explanation\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountShares\",\"type\":\"uint256\"}],\"name\":\"sharesToUnderlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountShares\",\"type\":\"uint256\"}],\"name\":\"sharesToUnderlyingView\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountUnderlying\",\"type\":\"uint256\"}],\"name\":\"underlyingToShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountUnderlying\",\"type\":\"uint256\"}],\"name\":\"underlyingToSharesView\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlyingToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"userUnderlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"userUnderlyingView\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountShares\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ContractIStrategyABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractIStrategyMetaData.ABI instead.
var ContractIStrategyABI = ContractIStrategyMetaData.ABI

// ContractIStrategy is an auto generated Go binding around an Ethereum contract.
type ContractIStrategy struct {
	ContractIStrategyCaller     // Read-only binding to the contract
	ContractIStrategyTransactor // Write-only binding to the contract
	ContractIStrategyFilterer   // Log filterer for contract events
}

// ContractIStrategyCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractIStrategyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIStrategyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractIStrategyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIStrategyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractIStrategyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIStrategySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractIStrategySession struct {
	Contract     *ContractIStrategy // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ContractIStrategyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractIStrategyCallerSession struct {
	Contract *ContractIStrategyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ContractIStrategyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractIStrategyTransactorSession struct {
	Contract     *ContractIStrategyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ContractIStrategyRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractIStrategyRaw struct {
	Contract *ContractIStrategy // Generic contract binding to access the raw methods on
}

// ContractIStrategyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractIStrategyCallerRaw struct {
	Contract *ContractIStrategyCaller // Generic read-only contract binding to access the raw methods on
}

// ContractIStrategyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractIStrategyTransactorRaw struct {
	Contract *ContractIStrategyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractIStrategy creates a new instance of ContractIStrategy, bound to a specific deployed contract.
func NewContractIStrategy(address common.Address, backend bind.ContractBackend) (*ContractIStrategy, error) {
	contract, err := bindContractIStrategy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractIStrategy{ContractIStrategyCaller: ContractIStrategyCaller{contract: contract}, ContractIStrategyTransactor: ContractIStrategyTransactor{contract: contract}, ContractIStrategyFilterer: ContractIStrategyFilterer{contract: contract}}, nil
}

// NewContractIStrategyCaller creates a new read-only instance of ContractIStrategy, bound to a specific deployed contract.
func NewContractIStrategyCaller(address common.Address, caller bind.ContractCaller) (*ContractIStrategyCaller, error) {
	contract, err := bindContractIStrategy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractIStrategyCaller{contract: contract}, nil
}

// NewContractIStrategyTransactor creates a new write-only instance of ContractIStrategy, bound to a specific deployed contract.
func NewContractIStrategyTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractIStrategyTransactor, error) {
	contract, err := bindContractIStrategy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractIStrategyTransactor{contract: contract}, nil
}

// NewContractIStrategyFilterer creates a new log filterer instance of ContractIStrategy, bound to a specific deployed contract.
func NewContractIStrategyFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractIStrategyFilterer, error) {
	contract, err := bindContractIStrategy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractIStrategyFilterer{contract: contract}, nil
}

// bindContractIStrategy binds a generic wrapper to an already deployed contract.
func bindContractIStrategy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractIStrategyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractIStrategy *ContractIStrategyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractIStrategy.Contract.ContractIStrategyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractIStrategy *ContractIStrategyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIStrategy.Contract.ContractIStrategyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractIStrategy *ContractIStrategyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractIStrategy.Contract.ContractIStrategyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractIStrategy *ContractIStrategyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractIStrategy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractIStrategy *ContractIStrategyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIStrategy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractIStrategy *ContractIStrategyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractIStrategy.Contract.contract.Transact(opts, method, params...)
}

// Explanation is a free data retrieval call binding the contract method 0xab5921e1.
//
// Solidity: function explanation() view returns(string)
func (_ContractIStrategy *ContractIStrategyCaller) Explanation(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ContractIStrategy.contract.Call(opts, &out, "explanation")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Explanation is a free data retrieval call binding the contract method 0xab5921e1.
//
// Solidity: function explanation() view returns(string)
func (_ContractIStrategy *ContractIStrategySession) Explanation() (string, error) {
	return _ContractIStrategy.Contract.Explanation(&_ContractIStrategy.CallOpts)
}

// Explanation is a free data retrieval call binding the contract method 0xab5921e1.
//
// Solidity: function explanation() view returns(string)
func (_ContractIStrategy *ContractIStrategyCallerSession) Explanation() (string, error) {
	return _ContractIStrategy.Contract.Explanation(&_ContractIStrategy.CallOpts)
}

// SharesToUnderlyingView is a free data retrieval call binding the contract method 0x7a8b2637.
//
// Solidity: function sharesToUnderlyingView(uint256 amountShares) view returns(uint256)
func (_ContractIStrategy *ContractIStrategyCaller) SharesToUnderlyingView(opts *bind.CallOpts, amountShares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ContractIStrategy.contract.Call(opts, &out, "sharesToUnderlyingView", amountShares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SharesToUnderlyingView is a free data retrieval call binding the contract method 0x7a8b2637.
//
// Solidity: function sharesToUnderlyingView(uint256 amountShares) view returns(uint256)
func (_ContractIStrategy *ContractIStrategySession) SharesToUnderlyingView(amountShares *big.Int) (*big.Int, error) {
	return _ContractIStrategy.Contract.SharesToUnderlyingView(&_ContractIStrategy.CallOpts, amountShares)
}

// SharesToUnderlyingView is a free data retrieval call binding the contract method 0x7a8b2637.
//
// Solidity: function sharesToUnderlyingView(uint256 amountShares) view returns(uint256)
func (_ContractIStrategy *ContractIStrategyCallerSession) SharesToUnderlyingView(amountShares *big.Int) (*big.Int, error) {
	return _ContractIStrategy.Contract.SharesToUnderlyingView(&_ContractIStrategy.CallOpts, amountShares)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_ContractIStrategy *ContractIStrategyCaller) TotalShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractIStrategy.contract.Call(opts, &out, "totalShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_ContractIStrategy *ContractIStrategySession) TotalShares() (*big.Int, error) {
	return _ContractIStrategy.Contract.TotalShares(&_ContractIStrategy.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_ContractIStrategy *ContractIStrategyCallerSession) TotalShares() (*big.Int, error) {
	return _ContractIStrategy.Contract.TotalShares(&_ContractIStrategy.CallOpts)
}

// UnderlyingToSharesView is a free data retrieval call binding the contract method 0xe3dae51c.
//
// Solidity: function underlyingToSharesView(uint256 amountUnderlying) view returns(uint256)
func (_ContractIStrategy *ContractIStrategyCaller) UnderlyingToSharesView(opts *bind.CallOpts, amountUnderlying *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ContractIStrategy.contract.Call(opts, &out, "underlyingToSharesView", amountUnderlying)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnderlyingToSharesView is a free data retrieval call binding the contract method 0xe3dae51c.
//
// Solidity: function underlyingToSharesView(uint256 amountUnderlying) view returns(uint256)
func (_ContractIStrategy *ContractIStrategySession) UnderlyingToSharesView(amountUnderlying *big.Int) (*big.Int, error) {
	return _ContractIStrategy.Contract.UnderlyingToSharesView(&_ContractIStrategy.CallOpts, amountUnderlying)
}

// UnderlyingToSharesView is a free data retrieval call binding the contract method 0xe3dae51c.
//
// Solidity: function underlyingToSharesView(uint256 amountUnderlying) view returns(uint256)
func (_ContractIStrategy *ContractIStrategyCallerSession) UnderlyingToSharesView(amountUnderlying *big.Int) (*big.Int, error) {
	return _ContractIStrategy.Contract.UnderlyingToSharesView(&_ContractIStrategy.CallOpts, amountUnderlying)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_ContractIStrategy *ContractIStrategyCaller) UnderlyingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIStrategy.contract.Call(opts, &out, "underlyingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_ContractIStrategy *ContractIStrategySession) UnderlyingToken() (common.Address, error) {
	return _ContractIStrategy.Contract.UnderlyingToken(&_ContractIStrategy.CallOpts)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_ContractIStrategy *ContractIStrategyCallerSession) UnderlyingToken() (common.Address, error) {
	return _ContractIStrategy.Contract.UnderlyingToken(&_ContractIStrategy.CallOpts)
}

// UserUnderlyingView is a free data retrieval call binding the contract method 0x553ca5f8.
//
// Solidity: function userUnderlyingView(address user) view returns(uint256)
func (_ContractIStrategy *ContractIStrategyCaller) UserUnderlyingView(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractIStrategy.contract.Call(opts, &out, "userUnderlyingView", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserUnderlyingView is a free data retrieval call binding the contract method 0x553ca5f8.
//
// Solidity: function userUnderlyingView(address user) view returns(uint256)
func (_ContractIStrategy *ContractIStrategySession) UserUnderlyingView(user common.Address) (*big.Int, error) {
	return _ContractIStrategy.Contract.UserUnderlyingView(&_ContractIStrategy.CallOpts, user)
}

// UserUnderlyingView is a free data retrieval call binding the contract method 0x553ca5f8.
//
// Solidity: function userUnderlyingView(address user) view returns(uint256)
func (_ContractIStrategy *ContractIStrategyCallerSession) UserUnderlyingView(user common.Address) (*big.Int, error) {
	return _ContractIStrategy.Contract.UserUnderlyingView(&_ContractIStrategy.CallOpts, user)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns(uint256)
func (_ContractIStrategy *ContractIStrategyTransactor) Deposit(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractIStrategy.contract.Transact(opts, "deposit", token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns(uint256)
func (_ContractIStrategy *ContractIStrategySession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractIStrategy.Contract.Deposit(&_ContractIStrategy.TransactOpts, token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns(uint256)
func (_ContractIStrategy *ContractIStrategyTransactorSession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractIStrategy.Contract.Deposit(&_ContractIStrategy.TransactOpts, token, amount)
}

// SharesToUnderlying is a paid mutator transaction binding the contract method 0xf3e73875.
//
// Solidity: function sharesToUnderlying(uint256 amountShares) returns(uint256)
func (_ContractIStrategy *ContractIStrategyTransactor) SharesToUnderlying(opts *bind.TransactOpts, amountShares *big.Int) (*types.Transaction, error) {
	return _ContractIStrategy.contract.Transact(opts, "sharesToUnderlying", amountShares)
}

// SharesToUnderlying is a paid mutator transaction binding the contract method 0xf3e73875.
//
// Solidity: function sharesToUnderlying(uint256 amountShares) returns(uint256)
func (_ContractIStrategy *ContractIStrategySession) SharesToUnderlying(amountShares *big.Int) (*types.Transaction, error) {
	return _ContractIStrategy.Contract.SharesToUnderlying(&_ContractIStrategy.TransactOpts, amountShares)
}

// SharesToUnderlying is a paid mutator transaction binding the contract method 0xf3e73875.
//
// Solidity: function sharesToUnderlying(uint256 amountShares) returns(uint256)
func (_ContractIStrategy *ContractIStrategyTransactorSession) SharesToUnderlying(amountShares *big.Int) (*types.Transaction, error) {
	return _ContractIStrategy.Contract.SharesToUnderlying(&_ContractIStrategy.TransactOpts, amountShares)
}

// UnderlyingToShares is a paid mutator transaction binding the contract method 0x8c871019.
//
// Solidity: function underlyingToShares(uint256 amountUnderlying) returns(uint256)
func (_ContractIStrategy *ContractIStrategyTransactor) UnderlyingToShares(opts *bind.TransactOpts, amountUnderlying *big.Int) (*types.Transaction, error) {
	return _ContractIStrategy.contract.Transact(opts, "underlyingToShares", amountUnderlying)
}

// UnderlyingToShares is a paid mutator transaction binding the contract method 0x8c871019.
//
// Solidity: function underlyingToShares(uint256 amountUnderlying) returns(uint256)
func (_ContractIStrategy *ContractIStrategySession) UnderlyingToShares(amountUnderlying *big.Int) (*types.Transaction, error) {
	return _ContractIStrategy.Contract.UnderlyingToShares(&_ContractIStrategy.TransactOpts, amountUnderlying)
}

// UnderlyingToShares is a paid mutator transaction binding the contract method 0x8c871019.
//
// Solidity: function underlyingToShares(uint256 amountUnderlying) returns(uint256)
func (_ContractIStrategy *ContractIStrategyTransactorSession) UnderlyingToShares(amountUnderlying *big.Int) (*types.Transaction, error) {
	return _ContractIStrategy.Contract.UnderlyingToShares(&_ContractIStrategy.TransactOpts, amountUnderlying)
}

// UserUnderlying is a paid mutator transaction binding the contract method 0x8f6a6240.
//
// Solidity: function userUnderlying(address user) returns(uint256)
func (_ContractIStrategy *ContractIStrategyTransactor) UserUnderlying(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _ContractIStrategy.contract.Transact(opts, "userUnderlying", user)
}

// UserUnderlying is a paid mutator transaction binding the contract method 0x8f6a6240.
//
// Solidity: function userUnderlying(address user) returns(uint256)
func (_ContractIStrategy *ContractIStrategySession) UserUnderlying(user common.Address) (*types.Transaction, error) {
	return _ContractIStrategy.Contract.UserUnderlying(&_ContractIStrategy.TransactOpts, user)
}

// UserUnderlying is a paid mutator transaction binding the contract method 0x8f6a6240.
//
// Solidity: function userUnderlying(address user) returns(uint256)
func (_ContractIStrategy *ContractIStrategyTransactorSession) UserUnderlying(user common.Address) (*types.Transaction, error) {
	return _ContractIStrategy.Contract.UserUnderlying(&_ContractIStrategy.TransactOpts, user)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address depositor, address token, uint256 amountShares) returns()
func (_ContractIStrategy *ContractIStrategyTransactor) Withdraw(opts *bind.TransactOpts, depositor common.Address, token common.Address, amountShares *big.Int) (*types.Transaction, error) {
	return _ContractIStrategy.contract.Transact(opts, "withdraw", depositor, token, amountShares)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address depositor, address token, uint256 amountShares) returns()
func (_ContractIStrategy *ContractIStrategySession) Withdraw(depositor common.Address, token common.Address, amountShares *big.Int) (*types.Transaction, error) {
	return _ContractIStrategy.Contract.Withdraw(&_ContractIStrategy.TransactOpts, depositor, token, amountShares)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address depositor, address token, uint256 amountShares) returns()
func (_ContractIStrategy *ContractIStrategyTransactorSession) Withdraw(depositor common.Address, token common.Address, amountShares *big.Int) (*types.Transaction, error) {
	return _ContractIStrategy.Contract.Withdraw(&_ContractIStrategy.TransactOpts, depositor, token, amountShares)
}
