// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IWhitelister

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

// IDelegationManagerQueuedWithdrawalParams is an auto generated low-level Go binding around an user-defined struct.
type IDelegationManagerQueuedWithdrawalParams struct {
	Strategies []common.Address
	Shares     []*big.Int
	Withdrawer common.Address
}

// IDelegationManagerWithdrawal is an auto generated low-level Go binding around an user-defined struct.
type IDelegationManagerWithdrawal struct {
	Staker      common.Address
	DelegatedTo common.Address
	Withdrawer  common.Address
	Nonce       *big.Int
	StartBlock  uint32
	Strategies  []common.Address
	Shares      []*big.Int
}

// IWhitelisterMetaData contains all meta data concerning the IWhitelister contract.
var IWhitelisterMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"callAddress\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"completeQueuedWithdrawal\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"queuedWithdrawal\",\"type\":\"tuple\",\"internalType\":\"structIDelegationManager.Withdrawal\",\"components\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegatedTo\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"shares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]},{\"name\":\"tokens\",\"type\":\"address[]\",\"internalType\":\"contractIERC20[]\"},{\"name\":\"middlewareTimesIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"receiveAsTokens\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositIntoStrategy\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getStaker\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"queueWithdrawal\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"queuedWithdrawalParams\",\"type\":\"tuple[]\",\"internalType\":\"structIDelegationManager.QueuedWithdrawalParams[]\",\"components\":[{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"shares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"whitelist\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
}

// IWhitelisterABI is the input ABI used to generate the binding from.
// Deprecated: Use IWhitelisterMetaData.ABI instead.
var IWhitelisterABI = IWhitelisterMetaData.ABI

// IWhitelister is an auto generated Go binding around an Ethereum contract.
type IWhitelister struct {
	IWhitelisterCaller     // Read-only binding to the contract
	IWhitelisterTransactor // Write-only binding to the contract
	IWhitelisterFilterer   // Log filterer for contract events
}

// IWhitelisterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IWhitelisterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWhitelisterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IWhitelisterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWhitelisterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IWhitelisterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWhitelisterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IWhitelisterSession struct {
	Contract     *IWhitelister     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IWhitelisterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IWhitelisterCallerSession struct {
	Contract *IWhitelisterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IWhitelisterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IWhitelisterTransactorSession struct {
	Contract     *IWhitelisterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IWhitelisterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IWhitelisterRaw struct {
	Contract *IWhitelister // Generic contract binding to access the raw methods on
}

// IWhitelisterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IWhitelisterCallerRaw struct {
	Contract *IWhitelisterCaller // Generic read-only contract binding to access the raw methods on
}

// IWhitelisterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IWhitelisterTransactorRaw struct {
	Contract *IWhitelisterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIWhitelister creates a new instance of IWhitelister, bound to a specific deployed contract.
func NewIWhitelister(address common.Address, backend bind.ContractBackend) (*IWhitelister, error) {
	contract, err := bindIWhitelister(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IWhitelister{IWhitelisterCaller: IWhitelisterCaller{contract: contract}, IWhitelisterTransactor: IWhitelisterTransactor{contract: contract}, IWhitelisterFilterer: IWhitelisterFilterer{contract: contract}}, nil
}

// NewIWhitelisterCaller creates a new read-only instance of IWhitelister, bound to a specific deployed contract.
func NewIWhitelisterCaller(address common.Address, caller bind.ContractCaller) (*IWhitelisterCaller, error) {
	contract, err := bindIWhitelister(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IWhitelisterCaller{contract: contract}, nil
}

// NewIWhitelisterTransactor creates a new write-only instance of IWhitelister, bound to a specific deployed contract.
func NewIWhitelisterTransactor(address common.Address, transactor bind.ContractTransactor) (*IWhitelisterTransactor, error) {
	contract, err := bindIWhitelister(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IWhitelisterTransactor{contract: contract}, nil
}

// NewIWhitelisterFilterer creates a new log filterer instance of IWhitelister, bound to a specific deployed contract.
func NewIWhitelisterFilterer(address common.Address, filterer bind.ContractFilterer) (*IWhitelisterFilterer, error) {
	contract, err := bindIWhitelister(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IWhitelisterFilterer{contract: contract}, nil
}

// bindIWhitelister binds a generic wrapper to an already deployed contract.
func bindIWhitelister(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IWhitelisterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWhitelister *IWhitelisterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWhitelister.Contract.IWhitelisterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWhitelister *IWhitelisterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWhitelister.Contract.IWhitelisterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWhitelister *IWhitelisterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWhitelister.Contract.IWhitelisterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWhitelister *IWhitelisterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWhitelister.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWhitelister *IWhitelisterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWhitelister.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWhitelister *IWhitelisterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWhitelister.Contract.contract.Transact(opts, method, params...)
}

// CallAddress is a paid mutator transaction binding the contract method 0xf6e8f39d.
//
// Solidity: function callAddress(address to, bytes data) payable returns(bytes)
func (_IWhitelister *IWhitelisterTransactor) CallAddress(opts *bind.TransactOpts, to common.Address, data []byte) (*types.Transaction, error) {
	return _IWhitelister.contract.Transact(opts, "callAddress", to, data)
}

// CallAddress is a paid mutator transaction binding the contract method 0xf6e8f39d.
//
// Solidity: function callAddress(address to, bytes data) payable returns(bytes)
func (_IWhitelister *IWhitelisterSession) CallAddress(to common.Address, data []byte) (*types.Transaction, error) {
	return _IWhitelister.Contract.CallAddress(&_IWhitelister.TransactOpts, to, data)
}

// CallAddress is a paid mutator transaction binding the contract method 0xf6e8f39d.
//
// Solidity: function callAddress(address to, bytes data) payable returns(bytes)
func (_IWhitelister *IWhitelisterTransactorSession) CallAddress(to common.Address, data []byte) (*types.Transaction, error) {
	return _IWhitelister.Contract.CallAddress(&_IWhitelister.TransactOpts, to, data)
}

// CompleteQueuedWithdrawal is a paid mutator transaction binding the contract method 0x063dcd50.
//
// Solidity: function completeQueuedWithdrawal(address staker, (address,address,address,uint256,uint32,address[],uint256[]) queuedWithdrawal, address[] tokens, uint256 middlewareTimesIndex, bool receiveAsTokens) returns(bytes)
func (_IWhitelister *IWhitelisterTransactor) CompleteQueuedWithdrawal(opts *bind.TransactOpts, staker common.Address, queuedWithdrawal IDelegationManagerWithdrawal, tokens []common.Address, middlewareTimesIndex *big.Int, receiveAsTokens bool) (*types.Transaction, error) {
	return _IWhitelister.contract.Transact(opts, "completeQueuedWithdrawal", staker, queuedWithdrawal, tokens, middlewareTimesIndex, receiveAsTokens)
}

// CompleteQueuedWithdrawal is a paid mutator transaction binding the contract method 0x063dcd50.
//
// Solidity: function completeQueuedWithdrawal(address staker, (address,address,address,uint256,uint32,address[],uint256[]) queuedWithdrawal, address[] tokens, uint256 middlewareTimesIndex, bool receiveAsTokens) returns(bytes)
func (_IWhitelister *IWhitelisterSession) CompleteQueuedWithdrawal(staker common.Address, queuedWithdrawal IDelegationManagerWithdrawal, tokens []common.Address, middlewareTimesIndex *big.Int, receiveAsTokens bool) (*types.Transaction, error) {
	return _IWhitelister.Contract.CompleteQueuedWithdrawal(&_IWhitelister.TransactOpts, staker, queuedWithdrawal, tokens, middlewareTimesIndex, receiveAsTokens)
}

// CompleteQueuedWithdrawal is a paid mutator transaction binding the contract method 0x063dcd50.
//
// Solidity: function completeQueuedWithdrawal(address staker, (address,address,address,uint256,uint32,address[],uint256[]) queuedWithdrawal, address[] tokens, uint256 middlewareTimesIndex, bool receiveAsTokens) returns(bytes)
func (_IWhitelister *IWhitelisterTransactorSession) CompleteQueuedWithdrawal(staker common.Address, queuedWithdrawal IDelegationManagerWithdrawal, tokens []common.Address, middlewareTimesIndex *big.Int, receiveAsTokens bool) (*types.Transaction, error) {
	return _IWhitelister.Contract.CompleteQueuedWithdrawal(&_IWhitelister.TransactOpts, staker, queuedWithdrawal, tokens, middlewareTimesIndex, receiveAsTokens)
}

// DepositIntoStrategy is a paid mutator transaction binding the contract method 0xa49ca158.
//
// Solidity: function depositIntoStrategy(address staker, address strategy, address token, uint256 amount) returns(bytes)
func (_IWhitelister *IWhitelisterTransactor) DepositIntoStrategy(opts *bind.TransactOpts, staker common.Address, strategy common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWhitelister.contract.Transact(opts, "depositIntoStrategy", staker, strategy, token, amount)
}

// DepositIntoStrategy is a paid mutator transaction binding the contract method 0xa49ca158.
//
// Solidity: function depositIntoStrategy(address staker, address strategy, address token, uint256 amount) returns(bytes)
func (_IWhitelister *IWhitelisterSession) DepositIntoStrategy(staker common.Address, strategy common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWhitelister.Contract.DepositIntoStrategy(&_IWhitelister.TransactOpts, staker, strategy, token, amount)
}

// DepositIntoStrategy is a paid mutator transaction binding the contract method 0xa49ca158.
//
// Solidity: function depositIntoStrategy(address staker, address strategy, address token, uint256 amount) returns(bytes)
func (_IWhitelister *IWhitelisterTransactorSession) DepositIntoStrategy(staker common.Address, strategy common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWhitelister.Contract.DepositIntoStrategy(&_IWhitelister.TransactOpts, staker, strategy, token, amount)
}

// GetStaker is a paid mutator transaction binding the contract method 0xa23c44b1.
//
// Solidity: function getStaker(address operator) returns(address)
func (_IWhitelister *IWhitelisterTransactor) GetStaker(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _IWhitelister.contract.Transact(opts, "getStaker", operator)
}

// GetStaker is a paid mutator transaction binding the contract method 0xa23c44b1.
//
// Solidity: function getStaker(address operator) returns(address)
func (_IWhitelister *IWhitelisterSession) GetStaker(operator common.Address) (*types.Transaction, error) {
	return _IWhitelister.Contract.GetStaker(&_IWhitelister.TransactOpts, operator)
}

// GetStaker is a paid mutator transaction binding the contract method 0xa23c44b1.
//
// Solidity: function getStaker(address operator) returns(address)
func (_IWhitelister *IWhitelisterTransactorSession) GetStaker(operator common.Address) (*types.Transaction, error) {
	return _IWhitelister.Contract.GetStaker(&_IWhitelister.TransactOpts, operator)
}

// QueueWithdrawal is a paid mutator transaction binding the contract method 0xa76a9d2d.
//
// Solidity: function queueWithdrawal(address staker, (address[],uint256[],address)[] queuedWithdrawalParams) returns(bytes)
func (_IWhitelister *IWhitelisterTransactor) QueueWithdrawal(opts *bind.TransactOpts, staker common.Address, queuedWithdrawalParams []IDelegationManagerQueuedWithdrawalParams) (*types.Transaction, error) {
	return _IWhitelister.contract.Transact(opts, "queueWithdrawal", staker, queuedWithdrawalParams)
}

// QueueWithdrawal is a paid mutator transaction binding the contract method 0xa76a9d2d.
//
// Solidity: function queueWithdrawal(address staker, (address[],uint256[],address)[] queuedWithdrawalParams) returns(bytes)
func (_IWhitelister *IWhitelisterSession) QueueWithdrawal(staker common.Address, queuedWithdrawalParams []IDelegationManagerQueuedWithdrawalParams) (*types.Transaction, error) {
	return _IWhitelister.Contract.QueueWithdrawal(&_IWhitelister.TransactOpts, staker, queuedWithdrawalParams)
}

// QueueWithdrawal is a paid mutator transaction binding the contract method 0xa76a9d2d.
//
// Solidity: function queueWithdrawal(address staker, (address[],uint256[],address)[] queuedWithdrawalParams) returns(bytes)
func (_IWhitelister *IWhitelisterTransactorSession) QueueWithdrawal(staker common.Address, queuedWithdrawalParams []IDelegationManagerQueuedWithdrawalParams) (*types.Transaction, error) {
	return _IWhitelister.Contract.QueueWithdrawal(&_IWhitelister.TransactOpts, staker, queuedWithdrawalParams)
}

// Transfer is a paid mutator transaction binding the contract method 0xf18d03cc.
//
// Solidity: function transfer(address staker, address token, address to, uint256 amount) returns(bytes)
func (_IWhitelister *IWhitelisterTransactor) Transfer(opts *bind.TransactOpts, staker common.Address, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWhitelister.contract.Transact(opts, "transfer", staker, token, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xf18d03cc.
//
// Solidity: function transfer(address staker, address token, address to, uint256 amount) returns(bytes)
func (_IWhitelister *IWhitelisterSession) Transfer(staker common.Address, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWhitelister.Contract.Transfer(&_IWhitelister.TransactOpts, staker, token, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xf18d03cc.
//
// Solidity: function transfer(address staker, address token, address to, uint256 amount) returns(bytes)
func (_IWhitelister *IWhitelisterTransactorSession) Transfer(staker common.Address, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWhitelister.Contract.Transfer(&_IWhitelister.TransactOpts, staker, token, to, amount)
}

// Whitelist is a paid mutator transaction binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address operator) returns()
func (_IWhitelister *IWhitelisterTransactor) Whitelist(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _IWhitelister.contract.Transact(opts, "whitelist", operator)
}

// Whitelist is a paid mutator transaction binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address operator) returns()
func (_IWhitelister *IWhitelisterSession) Whitelist(operator common.Address) (*types.Transaction, error) {
	return _IWhitelister.Contract.Whitelist(&_IWhitelister.TransactOpts, operator)
}

// Whitelist is a paid mutator transaction binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address operator) returns()
func (_IWhitelister *IWhitelisterTransactorSession) Whitelist(operator common.Address) (*types.Transaction, error) {
	return _IWhitelister.Contract.Whitelist(&_IWhitelister.TransactOpts, operator)
}
