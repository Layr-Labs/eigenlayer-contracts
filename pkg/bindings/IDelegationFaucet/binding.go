// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IDelegationFaucet

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

// ISignatureUtilsSignatureWithExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithExpiry struct {
	Signature []byte
	Expiry    *big.Int
}

// IDelegationFaucetMetaData contains all meta data concerning the IDelegationFaucet contract.
var IDelegationFaucetMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"callAddress\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"completeQueuedWithdrawal\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"queuedWithdrawal\",\"type\":\"tuple\",\"internalType\":\"structIDelegationManager.Withdrawal\",\"components\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegatedTo\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"shares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]},{\"name\":\"tokens\",\"type\":\"address[]\",\"internalType\":\"contractIERC20[]\"},{\"name\":\"middlewareTimesIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"receiveAsTokens\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositIntoStrategy\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getStaker\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mintDepositAndDelegate\",\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"approverSignatureAndExpiry\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"approverSalt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_depositAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"queueWithdrawal\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"queuedWithdrawalParams\",\"type\":\"tuple[]\",\"internalType\":\"structIDelegationManager.QueuedWithdrawalParams[]\",\"components\":[{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"shares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"}]",
}

// IDelegationFaucetABI is the input ABI used to generate the binding from.
// Deprecated: Use IDelegationFaucetMetaData.ABI instead.
var IDelegationFaucetABI = IDelegationFaucetMetaData.ABI

// IDelegationFaucet is an auto generated Go binding around an Ethereum contract.
type IDelegationFaucet struct {
	IDelegationFaucetCaller     // Read-only binding to the contract
	IDelegationFaucetTransactor // Write-only binding to the contract
	IDelegationFaucetFilterer   // Log filterer for contract events
}

// IDelegationFaucetCaller is an auto generated read-only Go binding around an Ethereum contract.
type IDelegationFaucetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDelegationFaucetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IDelegationFaucetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDelegationFaucetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IDelegationFaucetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDelegationFaucetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IDelegationFaucetSession struct {
	Contract     *IDelegationFaucet // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IDelegationFaucetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IDelegationFaucetCallerSession struct {
	Contract *IDelegationFaucetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IDelegationFaucetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IDelegationFaucetTransactorSession struct {
	Contract     *IDelegationFaucetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IDelegationFaucetRaw is an auto generated low-level Go binding around an Ethereum contract.
type IDelegationFaucetRaw struct {
	Contract *IDelegationFaucet // Generic contract binding to access the raw methods on
}

// IDelegationFaucetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IDelegationFaucetCallerRaw struct {
	Contract *IDelegationFaucetCaller // Generic read-only contract binding to access the raw methods on
}

// IDelegationFaucetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IDelegationFaucetTransactorRaw struct {
	Contract *IDelegationFaucetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIDelegationFaucet creates a new instance of IDelegationFaucet, bound to a specific deployed contract.
func NewIDelegationFaucet(address common.Address, backend bind.ContractBackend) (*IDelegationFaucet, error) {
	contract, err := bindIDelegationFaucet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IDelegationFaucet{IDelegationFaucetCaller: IDelegationFaucetCaller{contract: contract}, IDelegationFaucetTransactor: IDelegationFaucetTransactor{contract: contract}, IDelegationFaucetFilterer: IDelegationFaucetFilterer{contract: contract}}, nil
}

// NewIDelegationFaucetCaller creates a new read-only instance of IDelegationFaucet, bound to a specific deployed contract.
func NewIDelegationFaucetCaller(address common.Address, caller bind.ContractCaller) (*IDelegationFaucetCaller, error) {
	contract, err := bindIDelegationFaucet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IDelegationFaucetCaller{contract: contract}, nil
}

// NewIDelegationFaucetTransactor creates a new write-only instance of IDelegationFaucet, bound to a specific deployed contract.
func NewIDelegationFaucetTransactor(address common.Address, transactor bind.ContractTransactor) (*IDelegationFaucetTransactor, error) {
	contract, err := bindIDelegationFaucet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IDelegationFaucetTransactor{contract: contract}, nil
}

// NewIDelegationFaucetFilterer creates a new log filterer instance of IDelegationFaucet, bound to a specific deployed contract.
func NewIDelegationFaucetFilterer(address common.Address, filterer bind.ContractFilterer) (*IDelegationFaucetFilterer, error) {
	contract, err := bindIDelegationFaucet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IDelegationFaucetFilterer{contract: contract}, nil
}

// bindIDelegationFaucet binds a generic wrapper to an already deployed contract.
func bindIDelegationFaucet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IDelegationFaucetMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDelegationFaucet *IDelegationFaucetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDelegationFaucet.Contract.IDelegationFaucetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDelegationFaucet *IDelegationFaucetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDelegationFaucet.Contract.IDelegationFaucetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDelegationFaucet *IDelegationFaucetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDelegationFaucet.Contract.IDelegationFaucetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDelegationFaucet *IDelegationFaucetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDelegationFaucet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDelegationFaucet *IDelegationFaucetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDelegationFaucet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDelegationFaucet *IDelegationFaucetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDelegationFaucet.Contract.contract.Transact(opts, method, params...)
}

// CallAddress is a paid mutator transaction binding the contract method 0xf6e8f39d.
//
// Solidity: function callAddress(address to, bytes data) payable returns(bytes)
func (_IDelegationFaucet *IDelegationFaucetTransactor) CallAddress(opts *bind.TransactOpts, to common.Address, data []byte) (*types.Transaction, error) {
	return _IDelegationFaucet.contract.Transact(opts, "callAddress", to, data)
}

// CallAddress is a paid mutator transaction binding the contract method 0xf6e8f39d.
//
// Solidity: function callAddress(address to, bytes data) payable returns(bytes)
func (_IDelegationFaucet *IDelegationFaucetSession) CallAddress(to common.Address, data []byte) (*types.Transaction, error) {
	return _IDelegationFaucet.Contract.CallAddress(&_IDelegationFaucet.TransactOpts, to, data)
}

// CallAddress is a paid mutator transaction binding the contract method 0xf6e8f39d.
//
// Solidity: function callAddress(address to, bytes data) payable returns(bytes)
func (_IDelegationFaucet *IDelegationFaucetTransactorSession) CallAddress(to common.Address, data []byte) (*types.Transaction, error) {
	return _IDelegationFaucet.Contract.CallAddress(&_IDelegationFaucet.TransactOpts, to, data)
}

// CompleteQueuedWithdrawal is a paid mutator transaction binding the contract method 0x063dcd50.
//
// Solidity: function completeQueuedWithdrawal(address staker, (address,address,address,uint256,uint32,address[],uint256[]) queuedWithdrawal, address[] tokens, uint256 middlewareTimesIndex, bool receiveAsTokens) returns(bytes)
func (_IDelegationFaucet *IDelegationFaucetTransactor) CompleteQueuedWithdrawal(opts *bind.TransactOpts, staker common.Address, queuedWithdrawal IDelegationManagerWithdrawal, tokens []common.Address, middlewareTimesIndex *big.Int, receiveAsTokens bool) (*types.Transaction, error) {
	return _IDelegationFaucet.contract.Transact(opts, "completeQueuedWithdrawal", staker, queuedWithdrawal, tokens, middlewareTimesIndex, receiveAsTokens)
}

// CompleteQueuedWithdrawal is a paid mutator transaction binding the contract method 0x063dcd50.
//
// Solidity: function completeQueuedWithdrawal(address staker, (address,address,address,uint256,uint32,address[],uint256[]) queuedWithdrawal, address[] tokens, uint256 middlewareTimesIndex, bool receiveAsTokens) returns(bytes)
func (_IDelegationFaucet *IDelegationFaucetSession) CompleteQueuedWithdrawal(staker common.Address, queuedWithdrawal IDelegationManagerWithdrawal, tokens []common.Address, middlewareTimesIndex *big.Int, receiveAsTokens bool) (*types.Transaction, error) {
	return _IDelegationFaucet.Contract.CompleteQueuedWithdrawal(&_IDelegationFaucet.TransactOpts, staker, queuedWithdrawal, tokens, middlewareTimesIndex, receiveAsTokens)
}

// CompleteQueuedWithdrawal is a paid mutator transaction binding the contract method 0x063dcd50.
//
// Solidity: function completeQueuedWithdrawal(address staker, (address,address,address,uint256,uint32,address[],uint256[]) queuedWithdrawal, address[] tokens, uint256 middlewareTimesIndex, bool receiveAsTokens) returns(bytes)
func (_IDelegationFaucet *IDelegationFaucetTransactorSession) CompleteQueuedWithdrawal(staker common.Address, queuedWithdrawal IDelegationManagerWithdrawal, tokens []common.Address, middlewareTimesIndex *big.Int, receiveAsTokens bool) (*types.Transaction, error) {
	return _IDelegationFaucet.Contract.CompleteQueuedWithdrawal(&_IDelegationFaucet.TransactOpts, staker, queuedWithdrawal, tokens, middlewareTimesIndex, receiveAsTokens)
}

// DepositIntoStrategy is a paid mutator transaction binding the contract method 0xa49ca158.
//
// Solidity: function depositIntoStrategy(address staker, address strategy, address token, uint256 amount) returns(bytes)
func (_IDelegationFaucet *IDelegationFaucetTransactor) DepositIntoStrategy(opts *bind.TransactOpts, staker common.Address, strategy common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IDelegationFaucet.contract.Transact(opts, "depositIntoStrategy", staker, strategy, token, amount)
}

// DepositIntoStrategy is a paid mutator transaction binding the contract method 0xa49ca158.
//
// Solidity: function depositIntoStrategy(address staker, address strategy, address token, uint256 amount) returns(bytes)
func (_IDelegationFaucet *IDelegationFaucetSession) DepositIntoStrategy(staker common.Address, strategy common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IDelegationFaucet.Contract.DepositIntoStrategy(&_IDelegationFaucet.TransactOpts, staker, strategy, token, amount)
}

// DepositIntoStrategy is a paid mutator transaction binding the contract method 0xa49ca158.
//
// Solidity: function depositIntoStrategy(address staker, address strategy, address token, uint256 amount) returns(bytes)
func (_IDelegationFaucet *IDelegationFaucetTransactorSession) DepositIntoStrategy(staker common.Address, strategy common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IDelegationFaucet.Contract.DepositIntoStrategy(&_IDelegationFaucet.TransactOpts, staker, strategy, token, amount)
}

// GetStaker is a paid mutator transaction binding the contract method 0xa23c44b1.
//
// Solidity: function getStaker(address operator) returns(address)
func (_IDelegationFaucet *IDelegationFaucetTransactor) GetStaker(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _IDelegationFaucet.contract.Transact(opts, "getStaker", operator)
}

// GetStaker is a paid mutator transaction binding the contract method 0xa23c44b1.
//
// Solidity: function getStaker(address operator) returns(address)
func (_IDelegationFaucet *IDelegationFaucetSession) GetStaker(operator common.Address) (*types.Transaction, error) {
	return _IDelegationFaucet.Contract.GetStaker(&_IDelegationFaucet.TransactOpts, operator)
}

// GetStaker is a paid mutator transaction binding the contract method 0xa23c44b1.
//
// Solidity: function getStaker(address operator) returns(address)
func (_IDelegationFaucet *IDelegationFaucetTransactorSession) GetStaker(operator common.Address) (*types.Transaction, error) {
	return _IDelegationFaucet.Contract.GetStaker(&_IDelegationFaucet.TransactOpts, operator)
}

// MintDepositAndDelegate is a paid mutator transaction binding the contract method 0x34489506.
//
// Solidity: function mintDepositAndDelegate(address _operator, (bytes,uint256) approverSignatureAndExpiry, bytes32 approverSalt, uint256 _depositAmount) returns()
func (_IDelegationFaucet *IDelegationFaucetTransactor) MintDepositAndDelegate(opts *bind.TransactOpts, _operator common.Address, approverSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSalt [32]byte, _depositAmount *big.Int) (*types.Transaction, error) {
	return _IDelegationFaucet.contract.Transact(opts, "mintDepositAndDelegate", _operator, approverSignatureAndExpiry, approverSalt, _depositAmount)
}

// MintDepositAndDelegate is a paid mutator transaction binding the contract method 0x34489506.
//
// Solidity: function mintDepositAndDelegate(address _operator, (bytes,uint256) approverSignatureAndExpiry, bytes32 approverSalt, uint256 _depositAmount) returns()
func (_IDelegationFaucet *IDelegationFaucetSession) MintDepositAndDelegate(_operator common.Address, approverSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSalt [32]byte, _depositAmount *big.Int) (*types.Transaction, error) {
	return _IDelegationFaucet.Contract.MintDepositAndDelegate(&_IDelegationFaucet.TransactOpts, _operator, approverSignatureAndExpiry, approverSalt, _depositAmount)
}

// MintDepositAndDelegate is a paid mutator transaction binding the contract method 0x34489506.
//
// Solidity: function mintDepositAndDelegate(address _operator, (bytes,uint256) approverSignatureAndExpiry, bytes32 approverSalt, uint256 _depositAmount) returns()
func (_IDelegationFaucet *IDelegationFaucetTransactorSession) MintDepositAndDelegate(_operator common.Address, approverSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSalt [32]byte, _depositAmount *big.Int) (*types.Transaction, error) {
	return _IDelegationFaucet.Contract.MintDepositAndDelegate(&_IDelegationFaucet.TransactOpts, _operator, approverSignatureAndExpiry, approverSalt, _depositAmount)
}

// QueueWithdrawal is a paid mutator transaction binding the contract method 0xa76a9d2d.
//
// Solidity: function queueWithdrawal(address staker, (address[],uint256[],address)[] queuedWithdrawalParams) returns(bytes)
func (_IDelegationFaucet *IDelegationFaucetTransactor) QueueWithdrawal(opts *bind.TransactOpts, staker common.Address, queuedWithdrawalParams []IDelegationManagerQueuedWithdrawalParams) (*types.Transaction, error) {
	return _IDelegationFaucet.contract.Transact(opts, "queueWithdrawal", staker, queuedWithdrawalParams)
}

// QueueWithdrawal is a paid mutator transaction binding the contract method 0xa76a9d2d.
//
// Solidity: function queueWithdrawal(address staker, (address[],uint256[],address)[] queuedWithdrawalParams) returns(bytes)
func (_IDelegationFaucet *IDelegationFaucetSession) QueueWithdrawal(staker common.Address, queuedWithdrawalParams []IDelegationManagerQueuedWithdrawalParams) (*types.Transaction, error) {
	return _IDelegationFaucet.Contract.QueueWithdrawal(&_IDelegationFaucet.TransactOpts, staker, queuedWithdrawalParams)
}

// QueueWithdrawal is a paid mutator transaction binding the contract method 0xa76a9d2d.
//
// Solidity: function queueWithdrawal(address staker, (address[],uint256[],address)[] queuedWithdrawalParams) returns(bytes)
func (_IDelegationFaucet *IDelegationFaucetTransactorSession) QueueWithdrawal(staker common.Address, queuedWithdrawalParams []IDelegationManagerQueuedWithdrawalParams) (*types.Transaction, error) {
	return _IDelegationFaucet.Contract.QueueWithdrawal(&_IDelegationFaucet.TransactOpts, staker, queuedWithdrawalParams)
}

// Transfer is a paid mutator transaction binding the contract method 0xf18d03cc.
//
// Solidity: function transfer(address staker, address token, address to, uint256 amount) returns(bytes)
func (_IDelegationFaucet *IDelegationFaucetTransactor) Transfer(opts *bind.TransactOpts, staker common.Address, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IDelegationFaucet.contract.Transact(opts, "transfer", staker, token, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xf18d03cc.
//
// Solidity: function transfer(address staker, address token, address to, uint256 amount) returns(bytes)
func (_IDelegationFaucet *IDelegationFaucetSession) Transfer(staker common.Address, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IDelegationFaucet.Contract.Transfer(&_IDelegationFaucet.TransactOpts, staker, token, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xf18d03cc.
//
// Solidity: function transfer(address staker, address token, address to, uint256 amount) returns(bytes)
func (_IDelegationFaucet *IDelegationFaucetTransactorSession) Transfer(staker common.Address, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IDelegationFaucet.Contract.Transfer(&_IDelegationFaucet.TransactOpts, staker, token, to, amount)
}
