// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IStrategyFactory

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

// IStrategyFactoryMetaData contains all meta data concerning the IStrategyFactory contract.
var IStrategyFactoryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"deployNewStrategy\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"outputs\":[{\"name\":\"newStrategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployedStrategies\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeStrategiesFromWhitelist\",\"inputs\":[{\"name\":\"strategiesToRemoveFromWhitelist\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"strategyBeacon\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBeacon\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"whitelistStrategies\",\"inputs\":[{\"name\":\"strategiesToWhitelist\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"StrategyBeaconModified\",\"inputs\":[{\"name\":\"previousBeacon\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIBeacon\"},{\"name\":\"newBeacon\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIBeacon\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategySetForToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIERC20\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenBlacklisted\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIERC20\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyBlacklisted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BlacklistedToken\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategyAlreadyExists\",\"inputs\":[]}]",
}

// IStrategyFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use IStrategyFactoryMetaData.ABI instead.
var IStrategyFactoryABI = IStrategyFactoryMetaData.ABI

// IStrategyFactory is an auto generated Go binding around an Ethereum contract.
type IStrategyFactory struct {
	IStrategyFactoryCaller     // Read-only binding to the contract
	IStrategyFactoryTransactor // Write-only binding to the contract
	IStrategyFactoryFilterer   // Log filterer for contract events
}

// IStrategyFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IStrategyFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStrategyFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IStrategyFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStrategyFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IStrategyFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStrategyFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IStrategyFactorySession struct {
	Contract     *IStrategyFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IStrategyFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IStrategyFactoryCallerSession struct {
	Contract *IStrategyFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IStrategyFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IStrategyFactoryTransactorSession struct {
	Contract     *IStrategyFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IStrategyFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IStrategyFactoryRaw struct {
	Contract *IStrategyFactory // Generic contract binding to access the raw methods on
}

// IStrategyFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IStrategyFactoryCallerRaw struct {
	Contract *IStrategyFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// IStrategyFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IStrategyFactoryTransactorRaw struct {
	Contract *IStrategyFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIStrategyFactory creates a new instance of IStrategyFactory, bound to a specific deployed contract.
func NewIStrategyFactory(address common.Address, backend bind.ContractBackend) (*IStrategyFactory, error) {
	contract, err := bindIStrategyFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IStrategyFactory{IStrategyFactoryCaller: IStrategyFactoryCaller{contract: contract}, IStrategyFactoryTransactor: IStrategyFactoryTransactor{contract: contract}, IStrategyFactoryFilterer: IStrategyFactoryFilterer{contract: contract}}, nil
}

// NewIStrategyFactoryCaller creates a new read-only instance of IStrategyFactory, bound to a specific deployed contract.
func NewIStrategyFactoryCaller(address common.Address, caller bind.ContractCaller) (*IStrategyFactoryCaller, error) {
	contract, err := bindIStrategyFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IStrategyFactoryCaller{contract: contract}, nil
}

// NewIStrategyFactoryTransactor creates a new write-only instance of IStrategyFactory, bound to a specific deployed contract.
func NewIStrategyFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*IStrategyFactoryTransactor, error) {
	contract, err := bindIStrategyFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IStrategyFactoryTransactor{contract: contract}, nil
}

// NewIStrategyFactoryFilterer creates a new log filterer instance of IStrategyFactory, bound to a specific deployed contract.
func NewIStrategyFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*IStrategyFactoryFilterer, error) {
	contract, err := bindIStrategyFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IStrategyFactoryFilterer{contract: contract}, nil
}

// bindIStrategyFactory binds a generic wrapper to an already deployed contract.
func bindIStrategyFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IStrategyFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStrategyFactory *IStrategyFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStrategyFactory.Contract.IStrategyFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStrategyFactory *IStrategyFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStrategyFactory.Contract.IStrategyFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStrategyFactory *IStrategyFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStrategyFactory.Contract.IStrategyFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStrategyFactory *IStrategyFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStrategyFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStrategyFactory *IStrategyFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStrategyFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStrategyFactory *IStrategyFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStrategyFactory.Contract.contract.Transact(opts, method, params...)
}

// DeployedStrategies is a free data retrieval call binding the contract method 0x581dfd65.
//
// Solidity: function deployedStrategies(address token) view returns(address)
func (_IStrategyFactory *IStrategyFactoryCaller) DeployedStrategies(opts *bind.CallOpts, token common.Address) (common.Address, error) {
	var out []interface{}
	err := _IStrategyFactory.contract.Call(opts, &out, "deployedStrategies", token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DeployedStrategies is a free data retrieval call binding the contract method 0x581dfd65.
//
// Solidity: function deployedStrategies(address token) view returns(address)
func (_IStrategyFactory *IStrategyFactorySession) DeployedStrategies(token common.Address) (common.Address, error) {
	return _IStrategyFactory.Contract.DeployedStrategies(&_IStrategyFactory.CallOpts, token)
}

// DeployedStrategies is a free data retrieval call binding the contract method 0x581dfd65.
//
// Solidity: function deployedStrategies(address token) view returns(address)
func (_IStrategyFactory *IStrategyFactoryCallerSession) DeployedStrategies(token common.Address) (common.Address, error) {
	return _IStrategyFactory.Contract.DeployedStrategies(&_IStrategyFactory.CallOpts, token)
}

// StrategyBeacon is a free data retrieval call binding the contract method 0xf0062d9a.
//
// Solidity: function strategyBeacon() view returns(address)
func (_IStrategyFactory *IStrategyFactoryCaller) StrategyBeacon(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IStrategyFactory.contract.Call(opts, &out, "strategyBeacon")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StrategyBeacon is a free data retrieval call binding the contract method 0xf0062d9a.
//
// Solidity: function strategyBeacon() view returns(address)
func (_IStrategyFactory *IStrategyFactorySession) StrategyBeacon() (common.Address, error) {
	return _IStrategyFactory.Contract.StrategyBeacon(&_IStrategyFactory.CallOpts)
}

// StrategyBeacon is a free data retrieval call binding the contract method 0xf0062d9a.
//
// Solidity: function strategyBeacon() view returns(address)
func (_IStrategyFactory *IStrategyFactoryCallerSession) StrategyBeacon() (common.Address, error) {
	return _IStrategyFactory.Contract.StrategyBeacon(&_IStrategyFactory.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_IStrategyFactory *IStrategyFactoryCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IStrategyFactory.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_IStrategyFactory *IStrategyFactorySession) Version() (string, error) {
	return _IStrategyFactory.Contract.Version(&_IStrategyFactory.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_IStrategyFactory *IStrategyFactoryCallerSession) Version() (string, error) {
	return _IStrategyFactory.Contract.Version(&_IStrategyFactory.CallOpts)
}

// DeployNewStrategy is a paid mutator transaction binding the contract method 0x6b9b6229.
//
// Solidity: function deployNewStrategy(address token) returns(address newStrategy)
func (_IStrategyFactory *IStrategyFactoryTransactor) DeployNewStrategy(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _IStrategyFactory.contract.Transact(opts, "deployNewStrategy", token)
}

// DeployNewStrategy is a paid mutator transaction binding the contract method 0x6b9b6229.
//
// Solidity: function deployNewStrategy(address token) returns(address newStrategy)
func (_IStrategyFactory *IStrategyFactorySession) DeployNewStrategy(token common.Address) (*types.Transaction, error) {
	return _IStrategyFactory.Contract.DeployNewStrategy(&_IStrategyFactory.TransactOpts, token)
}

// DeployNewStrategy is a paid mutator transaction binding the contract method 0x6b9b6229.
//
// Solidity: function deployNewStrategy(address token) returns(address newStrategy)
func (_IStrategyFactory *IStrategyFactoryTransactorSession) DeployNewStrategy(token common.Address) (*types.Transaction, error) {
	return _IStrategyFactory.Contract.DeployNewStrategy(&_IStrategyFactory.TransactOpts, token)
}

// RemoveStrategiesFromWhitelist is a paid mutator transaction binding the contract method 0xfe38b32d.
//
// Solidity: function removeStrategiesFromWhitelist(address[] strategiesToRemoveFromWhitelist) returns()
func (_IStrategyFactory *IStrategyFactoryTransactor) RemoveStrategiesFromWhitelist(opts *bind.TransactOpts, strategiesToRemoveFromWhitelist []common.Address) (*types.Transaction, error) {
	return _IStrategyFactory.contract.Transact(opts, "removeStrategiesFromWhitelist", strategiesToRemoveFromWhitelist)
}

// RemoveStrategiesFromWhitelist is a paid mutator transaction binding the contract method 0xfe38b32d.
//
// Solidity: function removeStrategiesFromWhitelist(address[] strategiesToRemoveFromWhitelist) returns()
func (_IStrategyFactory *IStrategyFactorySession) RemoveStrategiesFromWhitelist(strategiesToRemoveFromWhitelist []common.Address) (*types.Transaction, error) {
	return _IStrategyFactory.Contract.RemoveStrategiesFromWhitelist(&_IStrategyFactory.TransactOpts, strategiesToRemoveFromWhitelist)
}

// RemoveStrategiesFromWhitelist is a paid mutator transaction binding the contract method 0xfe38b32d.
//
// Solidity: function removeStrategiesFromWhitelist(address[] strategiesToRemoveFromWhitelist) returns()
func (_IStrategyFactory *IStrategyFactoryTransactorSession) RemoveStrategiesFromWhitelist(strategiesToRemoveFromWhitelist []common.Address) (*types.Transaction, error) {
	return _IStrategyFactory.Contract.RemoveStrategiesFromWhitelist(&_IStrategyFactory.TransactOpts, strategiesToRemoveFromWhitelist)
}

// WhitelistStrategies is a paid mutator transaction binding the contract method 0xb768ebc9.
//
// Solidity: function whitelistStrategies(address[] strategiesToWhitelist) returns()
func (_IStrategyFactory *IStrategyFactoryTransactor) WhitelistStrategies(opts *bind.TransactOpts, strategiesToWhitelist []common.Address) (*types.Transaction, error) {
	return _IStrategyFactory.contract.Transact(opts, "whitelistStrategies", strategiesToWhitelist)
}

// WhitelistStrategies is a paid mutator transaction binding the contract method 0xb768ebc9.
//
// Solidity: function whitelistStrategies(address[] strategiesToWhitelist) returns()
func (_IStrategyFactory *IStrategyFactorySession) WhitelistStrategies(strategiesToWhitelist []common.Address) (*types.Transaction, error) {
	return _IStrategyFactory.Contract.WhitelistStrategies(&_IStrategyFactory.TransactOpts, strategiesToWhitelist)
}

// WhitelistStrategies is a paid mutator transaction binding the contract method 0xb768ebc9.
//
// Solidity: function whitelistStrategies(address[] strategiesToWhitelist) returns()
func (_IStrategyFactory *IStrategyFactoryTransactorSession) WhitelistStrategies(strategiesToWhitelist []common.Address) (*types.Transaction, error) {
	return _IStrategyFactory.Contract.WhitelistStrategies(&_IStrategyFactory.TransactOpts, strategiesToWhitelist)
}

// IStrategyFactoryStrategyBeaconModifiedIterator is returned from FilterStrategyBeaconModified and is used to iterate over the raw logs and unpacked data for StrategyBeaconModified events raised by the IStrategyFactory contract.
type IStrategyFactoryStrategyBeaconModifiedIterator struct {
	Event *IStrategyFactoryStrategyBeaconModified // Event containing the contract specifics and raw log

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
func (it *IStrategyFactoryStrategyBeaconModifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStrategyFactoryStrategyBeaconModified)
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
		it.Event = new(IStrategyFactoryStrategyBeaconModified)
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
func (it *IStrategyFactoryStrategyBeaconModifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStrategyFactoryStrategyBeaconModifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStrategyFactoryStrategyBeaconModified represents a StrategyBeaconModified event raised by the IStrategyFactory contract.
type IStrategyFactoryStrategyBeaconModified struct {
	PreviousBeacon common.Address
	NewBeacon      common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterStrategyBeaconModified is a free log retrieval operation binding the contract event 0xe21755962a7d7e100b59b9c3e4d4b54085b146313719955efb6a7a25c5c7feee.
//
// Solidity: event StrategyBeaconModified(address previousBeacon, address newBeacon)
func (_IStrategyFactory *IStrategyFactoryFilterer) FilterStrategyBeaconModified(opts *bind.FilterOpts) (*IStrategyFactoryStrategyBeaconModifiedIterator, error) {

	logs, sub, err := _IStrategyFactory.contract.FilterLogs(opts, "StrategyBeaconModified")
	if err != nil {
		return nil, err
	}
	return &IStrategyFactoryStrategyBeaconModifiedIterator{contract: _IStrategyFactory.contract, event: "StrategyBeaconModified", logs: logs, sub: sub}, nil
}

// WatchStrategyBeaconModified is a free log subscription operation binding the contract event 0xe21755962a7d7e100b59b9c3e4d4b54085b146313719955efb6a7a25c5c7feee.
//
// Solidity: event StrategyBeaconModified(address previousBeacon, address newBeacon)
func (_IStrategyFactory *IStrategyFactoryFilterer) WatchStrategyBeaconModified(opts *bind.WatchOpts, sink chan<- *IStrategyFactoryStrategyBeaconModified) (event.Subscription, error) {

	logs, sub, err := _IStrategyFactory.contract.WatchLogs(opts, "StrategyBeaconModified")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStrategyFactoryStrategyBeaconModified)
				if err := _IStrategyFactory.contract.UnpackLog(event, "StrategyBeaconModified", log); err != nil {
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
func (_IStrategyFactory *IStrategyFactoryFilterer) ParseStrategyBeaconModified(log types.Log) (*IStrategyFactoryStrategyBeaconModified, error) {
	event := new(IStrategyFactoryStrategyBeaconModified)
	if err := _IStrategyFactory.contract.UnpackLog(event, "StrategyBeaconModified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStrategyFactoryStrategySetForTokenIterator is returned from FilterStrategySetForToken and is used to iterate over the raw logs and unpacked data for StrategySetForToken events raised by the IStrategyFactory contract.
type IStrategyFactoryStrategySetForTokenIterator struct {
	Event *IStrategyFactoryStrategySetForToken // Event containing the contract specifics and raw log

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
func (it *IStrategyFactoryStrategySetForTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStrategyFactoryStrategySetForToken)
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
		it.Event = new(IStrategyFactoryStrategySetForToken)
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
func (it *IStrategyFactoryStrategySetForTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStrategyFactoryStrategySetForTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStrategyFactoryStrategySetForToken represents a StrategySetForToken event raised by the IStrategyFactory contract.
type IStrategyFactoryStrategySetForToken struct {
	Token    common.Address
	Strategy common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStrategySetForToken is a free log retrieval operation binding the contract event 0x6852a55230ef089d785bce7ffbf757985de34026df90a87d7b4a6e56f95d251f.
//
// Solidity: event StrategySetForToken(address token, address strategy)
func (_IStrategyFactory *IStrategyFactoryFilterer) FilterStrategySetForToken(opts *bind.FilterOpts) (*IStrategyFactoryStrategySetForTokenIterator, error) {

	logs, sub, err := _IStrategyFactory.contract.FilterLogs(opts, "StrategySetForToken")
	if err != nil {
		return nil, err
	}
	return &IStrategyFactoryStrategySetForTokenIterator{contract: _IStrategyFactory.contract, event: "StrategySetForToken", logs: logs, sub: sub}, nil
}

// WatchStrategySetForToken is a free log subscription operation binding the contract event 0x6852a55230ef089d785bce7ffbf757985de34026df90a87d7b4a6e56f95d251f.
//
// Solidity: event StrategySetForToken(address token, address strategy)
func (_IStrategyFactory *IStrategyFactoryFilterer) WatchStrategySetForToken(opts *bind.WatchOpts, sink chan<- *IStrategyFactoryStrategySetForToken) (event.Subscription, error) {

	logs, sub, err := _IStrategyFactory.contract.WatchLogs(opts, "StrategySetForToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStrategyFactoryStrategySetForToken)
				if err := _IStrategyFactory.contract.UnpackLog(event, "StrategySetForToken", log); err != nil {
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
func (_IStrategyFactory *IStrategyFactoryFilterer) ParseStrategySetForToken(log types.Log) (*IStrategyFactoryStrategySetForToken, error) {
	event := new(IStrategyFactoryStrategySetForToken)
	if err := _IStrategyFactory.contract.UnpackLog(event, "StrategySetForToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStrategyFactoryTokenBlacklistedIterator is returned from FilterTokenBlacklisted and is used to iterate over the raw logs and unpacked data for TokenBlacklisted events raised by the IStrategyFactory contract.
type IStrategyFactoryTokenBlacklistedIterator struct {
	Event *IStrategyFactoryTokenBlacklisted // Event containing the contract specifics and raw log

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
func (it *IStrategyFactoryTokenBlacklistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStrategyFactoryTokenBlacklisted)
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
		it.Event = new(IStrategyFactoryTokenBlacklisted)
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
func (it *IStrategyFactoryTokenBlacklistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStrategyFactoryTokenBlacklistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStrategyFactoryTokenBlacklisted represents a TokenBlacklisted event raised by the IStrategyFactory contract.
type IStrategyFactoryTokenBlacklisted struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenBlacklisted is a free log retrieval operation binding the contract event 0x75519c51f39873ec0e27dd3bbc09549e4865a113f505393fb9eab5898f6418b3.
//
// Solidity: event TokenBlacklisted(address token)
func (_IStrategyFactory *IStrategyFactoryFilterer) FilterTokenBlacklisted(opts *bind.FilterOpts) (*IStrategyFactoryTokenBlacklistedIterator, error) {

	logs, sub, err := _IStrategyFactory.contract.FilterLogs(opts, "TokenBlacklisted")
	if err != nil {
		return nil, err
	}
	return &IStrategyFactoryTokenBlacklistedIterator{contract: _IStrategyFactory.contract, event: "TokenBlacklisted", logs: logs, sub: sub}, nil
}

// WatchTokenBlacklisted is a free log subscription operation binding the contract event 0x75519c51f39873ec0e27dd3bbc09549e4865a113f505393fb9eab5898f6418b3.
//
// Solidity: event TokenBlacklisted(address token)
func (_IStrategyFactory *IStrategyFactoryFilterer) WatchTokenBlacklisted(opts *bind.WatchOpts, sink chan<- *IStrategyFactoryTokenBlacklisted) (event.Subscription, error) {

	logs, sub, err := _IStrategyFactory.contract.WatchLogs(opts, "TokenBlacklisted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStrategyFactoryTokenBlacklisted)
				if err := _IStrategyFactory.contract.UnpackLog(event, "TokenBlacklisted", log); err != nil {
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
func (_IStrategyFactory *IStrategyFactoryFilterer) ParseTokenBlacklisted(log types.Log) (*IStrategyFactoryTokenBlacklisted, error) {
	event := new(IStrategyFactoryTokenBlacklisted)
	if err := _IStrategyFactory.contract.UnpackLog(event, "TokenBlacklisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
