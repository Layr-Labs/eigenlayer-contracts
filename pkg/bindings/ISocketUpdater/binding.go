// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ISocketUpdater

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

// ISocketUpdaterMetaData contains all meta data concerning the ISocketUpdater contract.
var ISocketUpdaterMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"updateSocket\",\"inputs\":[{\"name\":\"socket\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"OperatorSocketUpdate\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"socket\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false}]",
}

// ISocketUpdaterABI is the input ABI used to generate the binding from.
// Deprecated: Use ISocketUpdaterMetaData.ABI instead.
var ISocketUpdaterABI = ISocketUpdaterMetaData.ABI

// ISocketUpdater is an auto generated Go binding around an Ethereum contract.
type ISocketUpdater struct {
	ISocketUpdaterCaller     // Read-only binding to the contract
	ISocketUpdaterTransactor // Write-only binding to the contract
	ISocketUpdaterFilterer   // Log filterer for contract events
}

// ISocketUpdaterCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISocketUpdaterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISocketUpdaterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISocketUpdaterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISocketUpdaterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISocketUpdaterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISocketUpdaterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISocketUpdaterSession struct {
	Contract     *ISocketUpdater   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISocketUpdaterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISocketUpdaterCallerSession struct {
	Contract *ISocketUpdaterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ISocketUpdaterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISocketUpdaterTransactorSession struct {
	Contract     *ISocketUpdaterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ISocketUpdaterRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISocketUpdaterRaw struct {
	Contract *ISocketUpdater // Generic contract binding to access the raw methods on
}

// ISocketUpdaterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISocketUpdaterCallerRaw struct {
	Contract *ISocketUpdaterCaller // Generic read-only contract binding to access the raw methods on
}

// ISocketUpdaterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISocketUpdaterTransactorRaw struct {
	Contract *ISocketUpdaterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISocketUpdater creates a new instance of ISocketUpdater, bound to a specific deployed contract.
func NewISocketUpdater(address common.Address, backend bind.ContractBackend) (*ISocketUpdater, error) {
	contract, err := bindISocketUpdater(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISocketUpdater{ISocketUpdaterCaller: ISocketUpdaterCaller{contract: contract}, ISocketUpdaterTransactor: ISocketUpdaterTransactor{contract: contract}, ISocketUpdaterFilterer: ISocketUpdaterFilterer{contract: contract}}, nil
}

// NewISocketUpdaterCaller creates a new read-only instance of ISocketUpdater, bound to a specific deployed contract.
func NewISocketUpdaterCaller(address common.Address, caller bind.ContractCaller) (*ISocketUpdaterCaller, error) {
	contract, err := bindISocketUpdater(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISocketUpdaterCaller{contract: contract}, nil
}

// NewISocketUpdaterTransactor creates a new write-only instance of ISocketUpdater, bound to a specific deployed contract.
func NewISocketUpdaterTransactor(address common.Address, transactor bind.ContractTransactor) (*ISocketUpdaterTransactor, error) {
	contract, err := bindISocketUpdater(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISocketUpdaterTransactor{contract: contract}, nil
}

// NewISocketUpdaterFilterer creates a new log filterer instance of ISocketUpdater, bound to a specific deployed contract.
func NewISocketUpdaterFilterer(address common.Address, filterer bind.ContractFilterer) (*ISocketUpdaterFilterer, error) {
	contract, err := bindISocketUpdater(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISocketUpdaterFilterer{contract: contract}, nil
}

// bindISocketUpdater binds a generic wrapper to an already deployed contract.
func bindISocketUpdater(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ISocketUpdaterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISocketUpdater *ISocketUpdaterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISocketUpdater.Contract.ISocketUpdaterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISocketUpdater *ISocketUpdaterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISocketUpdater.Contract.ISocketUpdaterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISocketUpdater *ISocketUpdaterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISocketUpdater.Contract.ISocketUpdaterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISocketUpdater *ISocketUpdaterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISocketUpdater.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISocketUpdater *ISocketUpdaterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISocketUpdater.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISocketUpdater *ISocketUpdaterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISocketUpdater.Contract.contract.Transact(opts, method, params...)
}

// UpdateSocket is a paid mutator transaction binding the contract method 0x0cf4b767.
//
// Solidity: function updateSocket(string socket) returns()
func (_ISocketUpdater *ISocketUpdaterTransactor) UpdateSocket(opts *bind.TransactOpts, socket string) (*types.Transaction, error) {
	return _ISocketUpdater.contract.Transact(opts, "updateSocket", socket)
}

// UpdateSocket is a paid mutator transaction binding the contract method 0x0cf4b767.
//
// Solidity: function updateSocket(string socket) returns()
func (_ISocketUpdater *ISocketUpdaterSession) UpdateSocket(socket string) (*types.Transaction, error) {
	return _ISocketUpdater.Contract.UpdateSocket(&_ISocketUpdater.TransactOpts, socket)
}

// UpdateSocket is a paid mutator transaction binding the contract method 0x0cf4b767.
//
// Solidity: function updateSocket(string socket) returns()
func (_ISocketUpdater *ISocketUpdaterTransactorSession) UpdateSocket(socket string) (*types.Transaction, error) {
	return _ISocketUpdater.Contract.UpdateSocket(&_ISocketUpdater.TransactOpts, socket)
}

// ISocketUpdaterOperatorSocketUpdateIterator is returned from FilterOperatorSocketUpdate and is used to iterate over the raw logs and unpacked data for OperatorSocketUpdate events raised by the ISocketUpdater contract.
type ISocketUpdaterOperatorSocketUpdateIterator struct {
	Event *ISocketUpdaterOperatorSocketUpdate // Event containing the contract specifics and raw log

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
func (it *ISocketUpdaterOperatorSocketUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ISocketUpdaterOperatorSocketUpdate)
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
		it.Event = new(ISocketUpdaterOperatorSocketUpdate)
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
func (it *ISocketUpdaterOperatorSocketUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ISocketUpdaterOperatorSocketUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ISocketUpdaterOperatorSocketUpdate represents a OperatorSocketUpdate event raised by the ISocketUpdater contract.
type ISocketUpdaterOperatorSocketUpdate struct {
	OperatorId [32]byte
	Socket     string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOperatorSocketUpdate is a free log retrieval operation binding the contract event 0xec2963ab21c1e50e1e582aa542af2e4bf7bf38e6e1403c27b42e1c5d6e621eaa.
//
// Solidity: event OperatorSocketUpdate(bytes32 indexed operatorId, string socket)
func (_ISocketUpdater *ISocketUpdaterFilterer) FilterOperatorSocketUpdate(opts *bind.FilterOpts, operatorId [][32]byte) (*ISocketUpdaterOperatorSocketUpdateIterator, error) {

	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _ISocketUpdater.contract.FilterLogs(opts, "OperatorSocketUpdate", operatorIdRule)
	if err != nil {
		return nil, err
	}
	return &ISocketUpdaterOperatorSocketUpdateIterator{contract: _ISocketUpdater.contract, event: "OperatorSocketUpdate", logs: logs, sub: sub}, nil
}

// WatchOperatorSocketUpdate is a free log subscription operation binding the contract event 0xec2963ab21c1e50e1e582aa542af2e4bf7bf38e6e1403c27b42e1c5d6e621eaa.
//
// Solidity: event OperatorSocketUpdate(bytes32 indexed operatorId, string socket)
func (_ISocketUpdater *ISocketUpdaterFilterer) WatchOperatorSocketUpdate(opts *bind.WatchOpts, sink chan<- *ISocketUpdaterOperatorSocketUpdate, operatorId [][32]byte) (event.Subscription, error) {

	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _ISocketUpdater.contract.WatchLogs(opts, "OperatorSocketUpdate", operatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ISocketUpdaterOperatorSocketUpdate)
				if err := _ISocketUpdater.contract.UnpackLog(event, "OperatorSocketUpdate", log); err != nil {
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

// ParseOperatorSocketUpdate is a log parse operation binding the contract event 0xec2963ab21c1e50e1e582aa542af2e4bf7bf38e6e1403c27b42e1c5d6e621eaa.
//
// Solidity: event OperatorSocketUpdate(bytes32 indexed operatorId, string socket)
func (_ISocketUpdater *ISocketUpdaterFilterer) ParseOperatorSocketUpdate(log types.Log) (*ISocketUpdaterOperatorSocketUpdate, error) {
	event := new(ISocketUpdaterOperatorSocketUpdate)
	if err := _ISocketUpdater.contract.UnpackLog(event, "OperatorSocketUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
