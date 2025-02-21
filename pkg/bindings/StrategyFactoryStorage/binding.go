// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package StrategyFactoryStorage

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

// StrategyFactoryStorageMetaData contains all meta data concerning the StrategyFactoryStorage contract.
var StrategyFactoryStorageMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"deployNewStrategy\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"outputs\":[{\"name\":\"newStrategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployedStrategies\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isBlacklisted\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeStrategiesFromWhitelist\",\"inputs\":[{\"name\":\"strategiesToRemoveFromWhitelist\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"strategyBeacon\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBeacon\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"whitelistStrategies\",\"inputs\":[{\"name\":\"strategiesToWhitelist\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"StrategyBeaconModified\",\"inputs\":[{\"name\":\"previousBeacon\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIBeacon\"},{\"name\":\"newBeacon\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIBeacon\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategySetForToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIERC20\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenBlacklisted\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIERC20\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyBlacklisted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BlacklistedToken\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategyAlreadyExists\",\"inputs\":[]}]",
}

// StrategyFactoryStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use StrategyFactoryStorageMetaData.ABI instead.
var StrategyFactoryStorageABI = StrategyFactoryStorageMetaData.ABI

// StrategyFactoryStorage is an auto generated Go binding around an Ethereum contract.
type StrategyFactoryStorage struct {
	StrategyFactoryStorageCaller     // Read-only binding to the contract
	StrategyFactoryStorageTransactor // Write-only binding to the contract
	StrategyFactoryStorageFilterer   // Log filterer for contract events
}

// StrategyFactoryStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type StrategyFactoryStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrategyFactoryStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StrategyFactoryStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrategyFactoryStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StrategyFactoryStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrategyFactoryStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StrategyFactoryStorageSession struct {
	Contract     *StrategyFactoryStorage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// StrategyFactoryStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StrategyFactoryStorageCallerSession struct {
	Contract *StrategyFactoryStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// StrategyFactoryStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StrategyFactoryStorageTransactorSession struct {
	Contract     *StrategyFactoryStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// StrategyFactoryStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type StrategyFactoryStorageRaw struct {
	Contract *StrategyFactoryStorage // Generic contract binding to access the raw methods on
}

// StrategyFactoryStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StrategyFactoryStorageCallerRaw struct {
	Contract *StrategyFactoryStorageCaller // Generic read-only contract binding to access the raw methods on
}

// StrategyFactoryStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StrategyFactoryStorageTransactorRaw struct {
	Contract *StrategyFactoryStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStrategyFactoryStorage creates a new instance of StrategyFactoryStorage, bound to a specific deployed contract.
func NewStrategyFactoryStorage(address common.Address, backend bind.ContractBackend) (*StrategyFactoryStorage, error) {
	contract, err := bindStrategyFactoryStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StrategyFactoryStorage{StrategyFactoryStorageCaller: StrategyFactoryStorageCaller{contract: contract}, StrategyFactoryStorageTransactor: StrategyFactoryStorageTransactor{contract: contract}, StrategyFactoryStorageFilterer: StrategyFactoryStorageFilterer{contract: contract}}, nil
}

// NewStrategyFactoryStorageCaller creates a new read-only instance of StrategyFactoryStorage, bound to a specific deployed contract.
func NewStrategyFactoryStorageCaller(address common.Address, caller bind.ContractCaller) (*StrategyFactoryStorageCaller, error) {
	contract, err := bindStrategyFactoryStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StrategyFactoryStorageCaller{contract: contract}, nil
}

// NewStrategyFactoryStorageTransactor creates a new write-only instance of StrategyFactoryStorage, bound to a specific deployed contract.
func NewStrategyFactoryStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*StrategyFactoryStorageTransactor, error) {
	contract, err := bindStrategyFactoryStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StrategyFactoryStorageTransactor{contract: contract}, nil
}

// NewStrategyFactoryStorageFilterer creates a new log filterer instance of StrategyFactoryStorage, bound to a specific deployed contract.
func NewStrategyFactoryStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*StrategyFactoryStorageFilterer, error) {
	contract, err := bindStrategyFactoryStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StrategyFactoryStorageFilterer{contract: contract}, nil
}

// bindStrategyFactoryStorage binds a generic wrapper to an already deployed contract.
func bindStrategyFactoryStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StrategyFactoryStorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StrategyFactoryStorage *StrategyFactoryStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StrategyFactoryStorage.Contract.StrategyFactoryStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StrategyFactoryStorage *StrategyFactoryStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrategyFactoryStorage.Contract.StrategyFactoryStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StrategyFactoryStorage *StrategyFactoryStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StrategyFactoryStorage.Contract.StrategyFactoryStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StrategyFactoryStorage *StrategyFactoryStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StrategyFactoryStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StrategyFactoryStorage *StrategyFactoryStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrategyFactoryStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StrategyFactoryStorage *StrategyFactoryStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StrategyFactoryStorage.Contract.contract.Transact(opts, method, params...)
}

// DeployedStrategies is a free data retrieval call binding the contract method 0x581dfd65.
//
// Solidity: function deployedStrategies(address ) view returns(address)
func (_StrategyFactoryStorage *StrategyFactoryStorageCaller) DeployedStrategies(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _StrategyFactoryStorage.contract.Call(opts, &out, "deployedStrategies", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DeployedStrategies is a free data retrieval call binding the contract method 0x581dfd65.
//
// Solidity: function deployedStrategies(address ) view returns(address)
func (_StrategyFactoryStorage *StrategyFactoryStorageSession) DeployedStrategies(arg0 common.Address) (common.Address, error) {
	return _StrategyFactoryStorage.Contract.DeployedStrategies(&_StrategyFactoryStorage.CallOpts, arg0)
}

// DeployedStrategies is a free data retrieval call binding the contract method 0x581dfd65.
//
// Solidity: function deployedStrategies(address ) view returns(address)
func (_StrategyFactoryStorage *StrategyFactoryStorageCallerSession) DeployedStrategies(arg0 common.Address) (common.Address, error) {
	return _StrategyFactoryStorage.Contract.DeployedStrategies(&_StrategyFactoryStorage.CallOpts, arg0)
}

// IsBlacklisted is a free data retrieval call binding the contract method 0xfe575a87.
//
// Solidity: function isBlacklisted(address ) view returns(bool)
func (_StrategyFactoryStorage *StrategyFactoryStorageCaller) IsBlacklisted(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _StrategyFactoryStorage.contract.Call(opts, &out, "isBlacklisted", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBlacklisted is a free data retrieval call binding the contract method 0xfe575a87.
//
// Solidity: function isBlacklisted(address ) view returns(bool)
func (_StrategyFactoryStorage *StrategyFactoryStorageSession) IsBlacklisted(arg0 common.Address) (bool, error) {
	return _StrategyFactoryStorage.Contract.IsBlacklisted(&_StrategyFactoryStorage.CallOpts, arg0)
}

// IsBlacklisted is a free data retrieval call binding the contract method 0xfe575a87.
//
// Solidity: function isBlacklisted(address ) view returns(bool)
func (_StrategyFactoryStorage *StrategyFactoryStorageCallerSession) IsBlacklisted(arg0 common.Address) (bool, error) {
	return _StrategyFactoryStorage.Contract.IsBlacklisted(&_StrategyFactoryStorage.CallOpts, arg0)
}

// StrategyBeacon is a free data retrieval call binding the contract method 0xf0062d9a.
//
// Solidity: function strategyBeacon() view returns(address)
func (_StrategyFactoryStorage *StrategyFactoryStorageCaller) StrategyBeacon(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyFactoryStorage.contract.Call(opts, &out, "strategyBeacon")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StrategyBeacon is a free data retrieval call binding the contract method 0xf0062d9a.
//
// Solidity: function strategyBeacon() view returns(address)
func (_StrategyFactoryStorage *StrategyFactoryStorageSession) StrategyBeacon() (common.Address, error) {
	return _StrategyFactoryStorage.Contract.StrategyBeacon(&_StrategyFactoryStorage.CallOpts)
}

// StrategyBeacon is a free data retrieval call binding the contract method 0xf0062d9a.
//
// Solidity: function strategyBeacon() view returns(address)
func (_StrategyFactoryStorage *StrategyFactoryStorageCallerSession) StrategyBeacon() (common.Address, error) {
	return _StrategyFactoryStorage.Contract.StrategyBeacon(&_StrategyFactoryStorage.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_StrategyFactoryStorage *StrategyFactoryStorageCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StrategyFactoryStorage.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_StrategyFactoryStorage *StrategyFactoryStorageSession) Version() (string, error) {
	return _StrategyFactoryStorage.Contract.Version(&_StrategyFactoryStorage.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_StrategyFactoryStorage *StrategyFactoryStorageCallerSession) Version() (string, error) {
	return _StrategyFactoryStorage.Contract.Version(&_StrategyFactoryStorage.CallOpts)
}

// DeployNewStrategy is a paid mutator transaction binding the contract method 0x6b9b6229.
//
// Solidity: function deployNewStrategy(address token) returns(address newStrategy)
func (_StrategyFactoryStorage *StrategyFactoryStorageTransactor) DeployNewStrategy(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _StrategyFactoryStorage.contract.Transact(opts, "deployNewStrategy", token)
}

// DeployNewStrategy is a paid mutator transaction binding the contract method 0x6b9b6229.
//
// Solidity: function deployNewStrategy(address token) returns(address newStrategy)
func (_StrategyFactoryStorage *StrategyFactoryStorageSession) DeployNewStrategy(token common.Address) (*types.Transaction, error) {
	return _StrategyFactoryStorage.Contract.DeployNewStrategy(&_StrategyFactoryStorage.TransactOpts, token)
}

// DeployNewStrategy is a paid mutator transaction binding the contract method 0x6b9b6229.
//
// Solidity: function deployNewStrategy(address token) returns(address newStrategy)
func (_StrategyFactoryStorage *StrategyFactoryStorageTransactorSession) DeployNewStrategy(token common.Address) (*types.Transaction, error) {
	return _StrategyFactoryStorage.Contract.DeployNewStrategy(&_StrategyFactoryStorage.TransactOpts, token)
}

// RemoveStrategiesFromWhitelist is a paid mutator transaction binding the contract method 0xfe38b32d.
//
// Solidity: function removeStrategiesFromWhitelist(address[] strategiesToRemoveFromWhitelist) returns()
func (_StrategyFactoryStorage *StrategyFactoryStorageTransactor) RemoveStrategiesFromWhitelist(opts *bind.TransactOpts, strategiesToRemoveFromWhitelist []common.Address) (*types.Transaction, error) {
	return _StrategyFactoryStorage.contract.Transact(opts, "removeStrategiesFromWhitelist", strategiesToRemoveFromWhitelist)
}

// RemoveStrategiesFromWhitelist is a paid mutator transaction binding the contract method 0xfe38b32d.
//
// Solidity: function removeStrategiesFromWhitelist(address[] strategiesToRemoveFromWhitelist) returns()
func (_StrategyFactoryStorage *StrategyFactoryStorageSession) RemoveStrategiesFromWhitelist(strategiesToRemoveFromWhitelist []common.Address) (*types.Transaction, error) {
	return _StrategyFactoryStorage.Contract.RemoveStrategiesFromWhitelist(&_StrategyFactoryStorage.TransactOpts, strategiesToRemoveFromWhitelist)
}

// RemoveStrategiesFromWhitelist is a paid mutator transaction binding the contract method 0xfe38b32d.
//
// Solidity: function removeStrategiesFromWhitelist(address[] strategiesToRemoveFromWhitelist) returns()
func (_StrategyFactoryStorage *StrategyFactoryStorageTransactorSession) RemoveStrategiesFromWhitelist(strategiesToRemoveFromWhitelist []common.Address) (*types.Transaction, error) {
	return _StrategyFactoryStorage.Contract.RemoveStrategiesFromWhitelist(&_StrategyFactoryStorage.TransactOpts, strategiesToRemoveFromWhitelist)
}

// WhitelistStrategies is a paid mutator transaction binding the contract method 0xb768ebc9.
//
// Solidity: function whitelistStrategies(address[] strategiesToWhitelist) returns()
func (_StrategyFactoryStorage *StrategyFactoryStorageTransactor) WhitelistStrategies(opts *bind.TransactOpts, strategiesToWhitelist []common.Address) (*types.Transaction, error) {
	return _StrategyFactoryStorage.contract.Transact(opts, "whitelistStrategies", strategiesToWhitelist)
}

// WhitelistStrategies is a paid mutator transaction binding the contract method 0xb768ebc9.
//
// Solidity: function whitelistStrategies(address[] strategiesToWhitelist) returns()
func (_StrategyFactoryStorage *StrategyFactoryStorageSession) WhitelistStrategies(strategiesToWhitelist []common.Address) (*types.Transaction, error) {
	return _StrategyFactoryStorage.Contract.WhitelistStrategies(&_StrategyFactoryStorage.TransactOpts, strategiesToWhitelist)
}

// WhitelistStrategies is a paid mutator transaction binding the contract method 0xb768ebc9.
//
// Solidity: function whitelistStrategies(address[] strategiesToWhitelist) returns()
func (_StrategyFactoryStorage *StrategyFactoryStorageTransactorSession) WhitelistStrategies(strategiesToWhitelist []common.Address) (*types.Transaction, error) {
	return _StrategyFactoryStorage.Contract.WhitelistStrategies(&_StrategyFactoryStorage.TransactOpts, strategiesToWhitelist)
}

// StrategyFactoryStorageStrategyBeaconModifiedIterator is returned from FilterStrategyBeaconModified and is used to iterate over the raw logs and unpacked data for StrategyBeaconModified events raised by the StrategyFactoryStorage contract.
type StrategyFactoryStorageStrategyBeaconModifiedIterator struct {
	Event *StrategyFactoryStorageStrategyBeaconModified // Event containing the contract specifics and raw log

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
func (it *StrategyFactoryStorageStrategyBeaconModifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyFactoryStorageStrategyBeaconModified)
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
		it.Event = new(StrategyFactoryStorageStrategyBeaconModified)
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
func (it *StrategyFactoryStorageStrategyBeaconModifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyFactoryStorageStrategyBeaconModifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyFactoryStorageStrategyBeaconModified represents a StrategyBeaconModified event raised by the StrategyFactoryStorage contract.
type StrategyFactoryStorageStrategyBeaconModified struct {
	PreviousBeacon common.Address
	NewBeacon      common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterStrategyBeaconModified is a free log retrieval operation binding the contract event 0xe21755962a7d7e100b59b9c3e4d4b54085b146313719955efb6a7a25c5c7feee.
//
// Solidity: event StrategyBeaconModified(address previousBeacon, address newBeacon)
func (_StrategyFactoryStorage *StrategyFactoryStorageFilterer) FilterStrategyBeaconModified(opts *bind.FilterOpts) (*StrategyFactoryStorageStrategyBeaconModifiedIterator, error) {

	logs, sub, err := _StrategyFactoryStorage.contract.FilterLogs(opts, "StrategyBeaconModified")
	if err != nil {
		return nil, err
	}
	return &StrategyFactoryStorageStrategyBeaconModifiedIterator{contract: _StrategyFactoryStorage.contract, event: "StrategyBeaconModified", logs: logs, sub: sub}, nil
}

// WatchStrategyBeaconModified is a free log subscription operation binding the contract event 0xe21755962a7d7e100b59b9c3e4d4b54085b146313719955efb6a7a25c5c7feee.
//
// Solidity: event StrategyBeaconModified(address previousBeacon, address newBeacon)
func (_StrategyFactoryStorage *StrategyFactoryStorageFilterer) WatchStrategyBeaconModified(opts *bind.WatchOpts, sink chan<- *StrategyFactoryStorageStrategyBeaconModified) (event.Subscription, error) {

	logs, sub, err := _StrategyFactoryStorage.contract.WatchLogs(opts, "StrategyBeaconModified")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyFactoryStorageStrategyBeaconModified)
				if err := _StrategyFactoryStorage.contract.UnpackLog(event, "StrategyBeaconModified", log); err != nil {
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

// ParseStrategyBeaconModified is a log parse operation binding the contract event 0xe21755962a7d7e100b59b9c3e4d4b54085b146313719955efb6a7a25c5c7feee.
//
// Solidity: event StrategyBeaconModified(address previousBeacon, address newBeacon)
func (_StrategyFactoryStorage *StrategyFactoryStorageFilterer) ParseStrategyBeaconModified(log types.Log) (*StrategyFactoryStorageStrategyBeaconModified, error) {
	event := new(StrategyFactoryStorageStrategyBeaconModified)
	if err := _StrategyFactoryStorage.contract.UnpackLog(event, "StrategyBeaconModified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyFactoryStorageStrategySetForTokenIterator is returned from FilterStrategySetForToken and is used to iterate over the raw logs and unpacked data for StrategySetForToken events raised by the StrategyFactoryStorage contract.
type StrategyFactoryStorageStrategySetForTokenIterator struct {
	Event *StrategyFactoryStorageStrategySetForToken // Event containing the contract specifics and raw log

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
func (it *StrategyFactoryStorageStrategySetForTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyFactoryStorageStrategySetForToken)
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
		it.Event = new(StrategyFactoryStorageStrategySetForToken)
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
func (it *StrategyFactoryStorageStrategySetForTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyFactoryStorageStrategySetForTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyFactoryStorageStrategySetForToken represents a StrategySetForToken event raised by the StrategyFactoryStorage contract.
type StrategyFactoryStorageStrategySetForToken struct {
	Token    common.Address
	Strategy common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStrategySetForToken is a free log retrieval operation binding the contract event 0x6852a55230ef089d785bce7ffbf757985de34026df90a87d7b4a6e56f95d251f.
//
// Solidity: event StrategySetForToken(address token, address strategy)
func (_StrategyFactoryStorage *StrategyFactoryStorageFilterer) FilterStrategySetForToken(opts *bind.FilterOpts) (*StrategyFactoryStorageStrategySetForTokenIterator, error) {

	logs, sub, err := _StrategyFactoryStorage.contract.FilterLogs(opts, "StrategySetForToken")
	if err != nil {
		return nil, err
	}
	return &StrategyFactoryStorageStrategySetForTokenIterator{contract: _StrategyFactoryStorage.contract, event: "StrategySetForToken", logs: logs, sub: sub}, nil
}

// WatchStrategySetForToken is a free log subscription operation binding the contract event 0x6852a55230ef089d785bce7ffbf757985de34026df90a87d7b4a6e56f95d251f.
//
// Solidity: event StrategySetForToken(address token, address strategy)
func (_StrategyFactoryStorage *StrategyFactoryStorageFilterer) WatchStrategySetForToken(opts *bind.WatchOpts, sink chan<- *StrategyFactoryStorageStrategySetForToken) (event.Subscription, error) {

	logs, sub, err := _StrategyFactoryStorage.contract.WatchLogs(opts, "StrategySetForToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyFactoryStorageStrategySetForToken)
				if err := _StrategyFactoryStorage.contract.UnpackLog(event, "StrategySetForToken", log); err != nil {
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

// ParseStrategySetForToken is a log parse operation binding the contract event 0x6852a55230ef089d785bce7ffbf757985de34026df90a87d7b4a6e56f95d251f.
//
// Solidity: event StrategySetForToken(address token, address strategy)
func (_StrategyFactoryStorage *StrategyFactoryStorageFilterer) ParseStrategySetForToken(log types.Log) (*StrategyFactoryStorageStrategySetForToken, error) {
	event := new(StrategyFactoryStorageStrategySetForToken)
	if err := _StrategyFactoryStorage.contract.UnpackLog(event, "StrategySetForToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyFactoryStorageTokenBlacklistedIterator is returned from FilterTokenBlacklisted and is used to iterate over the raw logs and unpacked data for TokenBlacklisted events raised by the StrategyFactoryStorage contract.
type StrategyFactoryStorageTokenBlacklistedIterator struct {
	Event *StrategyFactoryStorageTokenBlacklisted // Event containing the contract specifics and raw log

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
func (it *StrategyFactoryStorageTokenBlacklistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyFactoryStorageTokenBlacklisted)
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
		it.Event = new(StrategyFactoryStorageTokenBlacklisted)
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
func (it *StrategyFactoryStorageTokenBlacklistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyFactoryStorageTokenBlacklistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyFactoryStorageTokenBlacklisted represents a TokenBlacklisted event raised by the StrategyFactoryStorage contract.
type StrategyFactoryStorageTokenBlacklisted struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenBlacklisted is a free log retrieval operation binding the contract event 0x75519c51f39873ec0e27dd3bbc09549e4865a113f505393fb9eab5898f6418b3.
//
// Solidity: event TokenBlacklisted(address token)
func (_StrategyFactoryStorage *StrategyFactoryStorageFilterer) FilterTokenBlacklisted(opts *bind.FilterOpts) (*StrategyFactoryStorageTokenBlacklistedIterator, error) {

	logs, sub, err := _StrategyFactoryStorage.contract.FilterLogs(opts, "TokenBlacklisted")
	if err != nil {
		return nil, err
	}
	return &StrategyFactoryStorageTokenBlacklistedIterator{contract: _StrategyFactoryStorage.contract, event: "TokenBlacklisted", logs: logs, sub: sub}, nil
}

// WatchTokenBlacklisted is a free log subscription operation binding the contract event 0x75519c51f39873ec0e27dd3bbc09549e4865a113f505393fb9eab5898f6418b3.
//
// Solidity: event TokenBlacklisted(address token)
func (_StrategyFactoryStorage *StrategyFactoryStorageFilterer) WatchTokenBlacklisted(opts *bind.WatchOpts, sink chan<- *StrategyFactoryStorageTokenBlacklisted) (event.Subscription, error) {

	logs, sub, err := _StrategyFactoryStorage.contract.WatchLogs(opts, "TokenBlacklisted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyFactoryStorageTokenBlacklisted)
				if err := _StrategyFactoryStorage.contract.UnpackLog(event, "TokenBlacklisted", log); err != nil {
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

// ParseTokenBlacklisted is a log parse operation binding the contract event 0x75519c51f39873ec0e27dd3bbc09549e4865a113f505393fb9eab5898f6418b3.
//
// Solidity: event TokenBlacklisted(address token)
func (_StrategyFactoryStorage *StrategyFactoryStorageFilterer) ParseTokenBlacklisted(log types.Log) (*StrategyFactoryStorageTokenBlacklisted, error) {
	event := new(StrategyFactoryStorageTokenBlacklisted)
	if err := _StrategyFactoryStorage.contract.UnpackLog(event, "TokenBlacklisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
