// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IBeaconChainOracle

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

// IBeaconChainOracleMetaData contains all meta data concerning the IBeaconChainOracle contract.
var IBeaconChainOracleMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"timestampToBlockRoot\",\"inputs\":[{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"}]",
}

// IBeaconChainOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use IBeaconChainOracleMetaData.ABI instead.
var IBeaconChainOracleABI = IBeaconChainOracleMetaData.ABI

// IBeaconChainOracle is an auto generated Go binding around an Ethereum contract.
type IBeaconChainOracle struct {
	IBeaconChainOracleCaller     // Read-only binding to the contract
	IBeaconChainOracleTransactor // Write-only binding to the contract
	IBeaconChainOracleFilterer   // Log filterer for contract events
}

// IBeaconChainOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBeaconChainOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBeaconChainOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBeaconChainOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBeaconChainOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBeaconChainOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBeaconChainOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBeaconChainOracleSession struct {
	Contract     *IBeaconChainOracle // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IBeaconChainOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBeaconChainOracleCallerSession struct {
	Contract *IBeaconChainOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// IBeaconChainOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBeaconChainOracleTransactorSession struct {
	Contract     *IBeaconChainOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// IBeaconChainOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBeaconChainOracleRaw struct {
	Contract *IBeaconChainOracle // Generic contract binding to access the raw methods on
}

// IBeaconChainOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBeaconChainOracleCallerRaw struct {
	Contract *IBeaconChainOracleCaller // Generic read-only contract binding to access the raw methods on
}

// IBeaconChainOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBeaconChainOracleTransactorRaw struct {
	Contract *IBeaconChainOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBeaconChainOracle creates a new instance of IBeaconChainOracle, bound to a specific deployed contract.
func NewIBeaconChainOracle(address common.Address, backend bind.ContractBackend) (*IBeaconChainOracle, error) {
	contract, err := bindIBeaconChainOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBeaconChainOracle{IBeaconChainOracleCaller: IBeaconChainOracleCaller{contract: contract}, IBeaconChainOracleTransactor: IBeaconChainOracleTransactor{contract: contract}, IBeaconChainOracleFilterer: IBeaconChainOracleFilterer{contract: contract}}, nil
}

// NewIBeaconChainOracleCaller creates a new read-only instance of IBeaconChainOracle, bound to a specific deployed contract.
func NewIBeaconChainOracleCaller(address common.Address, caller bind.ContractCaller) (*IBeaconChainOracleCaller, error) {
	contract, err := bindIBeaconChainOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBeaconChainOracleCaller{contract: contract}, nil
}

// NewIBeaconChainOracleTransactor creates a new write-only instance of IBeaconChainOracle, bound to a specific deployed contract.
func NewIBeaconChainOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*IBeaconChainOracleTransactor, error) {
	contract, err := bindIBeaconChainOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBeaconChainOracleTransactor{contract: contract}, nil
}

// NewIBeaconChainOracleFilterer creates a new log filterer instance of IBeaconChainOracle, bound to a specific deployed contract.
func NewIBeaconChainOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*IBeaconChainOracleFilterer, error) {
	contract, err := bindIBeaconChainOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBeaconChainOracleFilterer{contract: contract}, nil
}

// bindIBeaconChainOracle binds a generic wrapper to an already deployed contract.
func bindIBeaconChainOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IBeaconChainOracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBeaconChainOracle *IBeaconChainOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBeaconChainOracle.Contract.IBeaconChainOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBeaconChainOracle *IBeaconChainOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBeaconChainOracle.Contract.IBeaconChainOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBeaconChainOracle *IBeaconChainOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBeaconChainOracle.Contract.IBeaconChainOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBeaconChainOracle *IBeaconChainOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBeaconChainOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBeaconChainOracle *IBeaconChainOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBeaconChainOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBeaconChainOracle *IBeaconChainOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBeaconChainOracle.Contract.contract.Transact(opts, method, params...)
}

// TimestampToBlockRoot is a free data retrieval call binding the contract method 0x643599f2.
//
// Solidity: function timestampToBlockRoot(uint256 timestamp) view returns(bytes32)
func (_IBeaconChainOracle *IBeaconChainOracleCaller) TimestampToBlockRoot(opts *bind.CallOpts, timestamp *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _IBeaconChainOracle.contract.Call(opts, &out, "timestampToBlockRoot", timestamp)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TimestampToBlockRoot is a free data retrieval call binding the contract method 0x643599f2.
//
// Solidity: function timestampToBlockRoot(uint256 timestamp) view returns(bytes32)
func (_IBeaconChainOracle *IBeaconChainOracleSession) TimestampToBlockRoot(timestamp *big.Int) ([32]byte, error) {
	return _IBeaconChainOracle.Contract.TimestampToBlockRoot(&_IBeaconChainOracle.CallOpts, timestamp)
}

// TimestampToBlockRoot is a free data retrieval call binding the contract method 0x643599f2.
//
// Solidity: function timestampToBlockRoot(uint256 timestamp) view returns(bytes32)
func (_IBeaconChainOracle *IBeaconChainOracleCallerSession) TimestampToBlockRoot(timestamp *big.Int) ([32]byte, error) {
	return _IBeaconChainOracle.Contract.TimestampToBlockRoot(&_IBeaconChainOracle.CallOpts, timestamp)
}
