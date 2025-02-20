// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IShareManager

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

// IShareManagerMetaData contains all meta data concerning the IShareManager contract.
var IShareManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"addShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"increaseBurnableShares\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"addedSharesToBurn\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeDepositShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"depositSharesToRemove\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakerDepositShares\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"depositShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawSharesAsTokens\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
}

// IShareManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use IShareManagerMetaData.ABI instead.
var IShareManagerABI = IShareManagerMetaData.ABI

// IShareManager is an auto generated Go binding around an Ethereum contract.
type IShareManager struct {
	IShareManagerCaller     // Read-only binding to the contract
	IShareManagerTransactor // Write-only binding to the contract
	IShareManagerFilterer   // Log filterer for contract events
}

// IShareManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IShareManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IShareManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IShareManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IShareManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IShareManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IShareManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IShareManagerSession struct {
	Contract     *IShareManager    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IShareManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IShareManagerCallerSession struct {
	Contract *IShareManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IShareManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IShareManagerTransactorSession struct {
	Contract     *IShareManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IShareManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IShareManagerRaw struct {
	Contract *IShareManager // Generic contract binding to access the raw methods on
}

// IShareManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IShareManagerCallerRaw struct {
	Contract *IShareManagerCaller // Generic read-only contract binding to access the raw methods on
}

// IShareManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IShareManagerTransactorRaw struct {
	Contract *IShareManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIShareManager creates a new instance of IShareManager, bound to a specific deployed contract.
func NewIShareManager(address common.Address, backend bind.ContractBackend) (*IShareManager, error) {
	contract, err := bindIShareManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IShareManager{IShareManagerCaller: IShareManagerCaller{contract: contract}, IShareManagerTransactor: IShareManagerTransactor{contract: contract}, IShareManagerFilterer: IShareManagerFilterer{contract: contract}}, nil
}

// NewIShareManagerCaller creates a new read-only instance of IShareManager, bound to a specific deployed contract.
func NewIShareManagerCaller(address common.Address, caller bind.ContractCaller) (*IShareManagerCaller, error) {
	contract, err := bindIShareManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IShareManagerCaller{contract: contract}, nil
}

// NewIShareManagerTransactor creates a new write-only instance of IShareManager, bound to a specific deployed contract.
func NewIShareManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IShareManagerTransactor, error) {
	contract, err := bindIShareManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IShareManagerTransactor{contract: contract}, nil
}

// NewIShareManagerFilterer creates a new log filterer instance of IShareManager, bound to a specific deployed contract.
func NewIShareManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IShareManagerFilterer, error) {
	contract, err := bindIShareManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IShareManagerFilterer{contract: contract}, nil
}

// bindIShareManager binds a generic wrapper to an already deployed contract.
func bindIShareManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IShareManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IShareManager *IShareManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IShareManager.Contract.IShareManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IShareManager *IShareManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IShareManager.Contract.IShareManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IShareManager *IShareManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IShareManager.Contract.IShareManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IShareManager *IShareManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IShareManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IShareManager *IShareManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IShareManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IShareManager *IShareManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IShareManager.Contract.contract.Transact(opts, method, params...)
}

// StakerDepositShares is a free data retrieval call binding the contract method 0xfe243a17.
//
// Solidity: function stakerDepositShares(address user, address strategy) view returns(uint256 depositShares)
func (_IShareManager *IShareManagerCaller) StakerDepositShares(opts *bind.CallOpts, user common.Address, strategy common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IShareManager.contract.Call(opts, &out, "stakerDepositShares", user, strategy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerDepositShares is a free data retrieval call binding the contract method 0xfe243a17.
//
// Solidity: function stakerDepositShares(address user, address strategy) view returns(uint256 depositShares)
func (_IShareManager *IShareManagerSession) StakerDepositShares(user common.Address, strategy common.Address) (*big.Int, error) {
	return _IShareManager.Contract.StakerDepositShares(&_IShareManager.CallOpts, user, strategy)
}

// StakerDepositShares is a free data retrieval call binding the contract method 0xfe243a17.
//
// Solidity: function stakerDepositShares(address user, address strategy) view returns(uint256 depositShares)
func (_IShareManager *IShareManagerCallerSession) StakerDepositShares(user common.Address, strategy common.Address) (*big.Int, error) {
	return _IShareManager.Contract.StakerDepositShares(&_IShareManager.CallOpts, user, strategy)
}

// AddShares is a paid mutator transaction binding the contract method 0x50ff7225.
//
// Solidity: function addShares(address staker, address strategy, uint256 shares) returns(uint256, uint256)
func (_IShareManager *IShareManagerTransactor) AddShares(opts *bind.TransactOpts, staker common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error) {
	return _IShareManager.contract.Transact(opts, "addShares", staker, strategy, shares)
}

// AddShares is a paid mutator transaction binding the contract method 0x50ff7225.
//
// Solidity: function addShares(address staker, address strategy, uint256 shares) returns(uint256, uint256)
func (_IShareManager *IShareManagerSession) AddShares(staker common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error) {
	return _IShareManager.Contract.AddShares(&_IShareManager.TransactOpts, staker, strategy, shares)
}

// AddShares is a paid mutator transaction binding the contract method 0x50ff7225.
//
// Solidity: function addShares(address staker, address strategy, uint256 shares) returns(uint256, uint256)
func (_IShareManager *IShareManagerTransactorSession) AddShares(staker common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error) {
	return _IShareManager.Contract.AddShares(&_IShareManager.TransactOpts, staker, strategy, shares)
}

// IncreaseBurnableShares is a paid mutator transaction binding the contract method 0xdebe1eab.
//
// Solidity: function increaseBurnableShares(address strategy, uint256 addedSharesToBurn) returns()
func (_IShareManager *IShareManagerTransactor) IncreaseBurnableShares(opts *bind.TransactOpts, strategy common.Address, addedSharesToBurn *big.Int) (*types.Transaction, error) {
	return _IShareManager.contract.Transact(opts, "increaseBurnableShares", strategy, addedSharesToBurn)
}

// IncreaseBurnableShares is a paid mutator transaction binding the contract method 0xdebe1eab.
//
// Solidity: function increaseBurnableShares(address strategy, uint256 addedSharesToBurn) returns()
func (_IShareManager *IShareManagerSession) IncreaseBurnableShares(strategy common.Address, addedSharesToBurn *big.Int) (*types.Transaction, error) {
	return _IShareManager.Contract.IncreaseBurnableShares(&_IShareManager.TransactOpts, strategy, addedSharesToBurn)
}

// IncreaseBurnableShares is a paid mutator transaction binding the contract method 0xdebe1eab.
//
// Solidity: function increaseBurnableShares(address strategy, uint256 addedSharesToBurn) returns()
func (_IShareManager *IShareManagerTransactorSession) IncreaseBurnableShares(strategy common.Address, addedSharesToBurn *big.Int) (*types.Transaction, error) {
	return _IShareManager.Contract.IncreaseBurnableShares(&_IShareManager.TransactOpts, strategy, addedSharesToBurn)
}

// RemoveDepositShares is a paid mutator transaction binding the contract method 0x724af423.
//
// Solidity: function removeDepositShares(address staker, address strategy, uint256 depositSharesToRemove) returns(uint256)
func (_IShareManager *IShareManagerTransactor) RemoveDepositShares(opts *bind.TransactOpts, staker common.Address, strategy common.Address, depositSharesToRemove *big.Int) (*types.Transaction, error) {
	return _IShareManager.contract.Transact(opts, "removeDepositShares", staker, strategy, depositSharesToRemove)
}

// RemoveDepositShares is a paid mutator transaction binding the contract method 0x724af423.
//
// Solidity: function removeDepositShares(address staker, address strategy, uint256 depositSharesToRemove) returns(uint256)
func (_IShareManager *IShareManagerSession) RemoveDepositShares(staker common.Address, strategy common.Address, depositSharesToRemove *big.Int) (*types.Transaction, error) {
	return _IShareManager.Contract.RemoveDepositShares(&_IShareManager.TransactOpts, staker, strategy, depositSharesToRemove)
}

// RemoveDepositShares is a paid mutator transaction binding the contract method 0x724af423.
//
// Solidity: function removeDepositShares(address staker, address strategy, uint256 depositSharesToRemove) returns(uint256)
func (_IShareManager *IShareManagerTransactorSession) RemoveDepositShares(staker common.Address, strategy common.Address, depositSharesToRemove *big.Int) (*types.Transaction, error) {
	return _IShareManager.Contract.RemoveDepositShares(&_IShareManager.TransactOpts, staker, strategy, depositSharesToRemove)
}

// WithdrawSharesAsTokens is a paid mutator transaction binding the contract method 0x2eae418c.
//
// Solidity: function withdrawSharesAsTokens(address staker, address strategy, address token, uint256 shares) returns()
func (_IShareManager *IShareManagerTransactor) WithdrawSharesAsTokens(opts *bind.TransactOpts, staker common.Address, strategy common.Address, token common.Address, shares *big.Int) (*types.Transaction, error) {
	return _IShareManager.contract.Transact(opts, "withdrawSharesAsTokens", staker, strategy, token, shares)
}

// WithdrawSharesAsTokens is a paid mutator transaction binding the contract method 0x2eae418c.
//
// Solidity: function withdrawSharesAsTokens(address staker, address strategy, address token, uint256 shares) returns()
func (_IShareManager *IShareManagerSession) WithdrawSharesAsTokens(staker common.Address, strategy common.Address, token common.Address, shares *big.Int) (*types.Transaction, error) {
	return _IShareManager.Contract.WithdrawSharesAsTokens(&_IShareManager.TransactOpts, staker, strategy, token, shares)
}

// WithdrawSharesAsTokens is a paid mutator transaction binding the contract method 0x2eae418c.
//
// Solidity: function withdrawSharesAsTokens(address staker, address strategy, address token, uint256 shares) returns()
func (_IShareManager *IShareManagerTransactorSession) WithdrawSharesAsTokens(staker common.Address, strategy common.Address, token common.Address, shares *big.Int) (*types.Transaction, error) {
	return _IShareManager.Contract.WithdrawSharesAsTokens(&_IShareManager.TransactOpts, staker, strategy, token, shares)
}
