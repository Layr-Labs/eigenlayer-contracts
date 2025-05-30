// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Deprecated_OwnableUpgradeable

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

// DeprecatedOwnableUpgradeableMetaData contains all meta data concerning the DeprecatedOwnableUpgradeable contract.
var DeprecatedOwnableUpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false}]",
}

// DeprecatedOwnableUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use DeprecatedOwnableUpgradeableMetaData.ABI instead.
var DeprecatedOwnableUpgradeableABI = DeprecatedOwnableUpgradeableMetaData.ABI

// DeprecatedOwnableUpgradeable is an auto generated Go binding around an Ethereum contract.
type DeprecatedOwnableUpgradeable struct {
	DeprecatedOwnableUpgradeableCaller     // Read-only binding to the contract
	DeprecatedOwnableUpgradeableTransactor // Write-only binding to the contract
	DeprecatedOwnableUpgradeableFilterer   // Log filterer for contract events
}

// DeprecatedOwnableUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type DeprecatedOwnableUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeprecatedOwnableUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DeprecatedOwnableUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeprecatedOwnableUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DeprecatedOwnableUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeprecatedOwnableUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DeprecatedOwnableUpgradeableSession struct {
	Contract     *DeprecatedOwnableUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// DeprecatedOwnableUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DeprecatedOwnableUpgradeableCallerSession struct {
	Contract *DeprecatedOwnableUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// DeprecatedOwnableUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DeprecatedOwnableUpgradeableTransactorSession struct {
	Contract     *DeprecatedOwnableUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// DeprecatedOwnableUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type DeprecatedOwnableUpgradeableRaw struct {
	Contract *DeprecatedOwnableUpgradeable // Generic contract binding to access the raw methods on
}

// DeprecatedOwnableUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DeprecatedOwnableUpgradeableCallerRaw struct {
	Contract *DeprecatedOwnableUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// DeprecatedOwnableUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DeprecatedOwnableUpgradeableTransactorRaw struct {
	Contract *DeprecatedOwnableUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDeprecatedOwnableUpgradeable creates a new instance of DeprecatedOwnableUpgradeable, bound to a specific deployed contract.
func NewDeprecatedOwnableUpgradeable(address common.Address, backend bind.ContractBackend) (*DeprecatedOwnableUpgradeable, error) {
	contract, err := bindDeprecatedOwnableUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DeprecatedOwnableUpgradeable{DeprecatedOwnableUpgradeableCaller: DeprecatedOwnableUpgradeableCaller{contract: contract}, DeprecatedOwnableUpgradeableTransactor: DeprecatedOwnableUpgradeableTransactor{contract: contract}, DeprecatedOwnableUpgradeableFilterer: DeprecatedOwnableUpgradeableFilterer{contract: contract}}, nil
}

// NewDeprecatedOwnableUpgradeableCaller creates a new read-only instance of DeprecatedOwnableUpgradeable, bound to a specific deployed contract.
func NewDeprecatedOwnableUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*DeprecatedOwnableUpgradeableCaller, error) {
	contract, err := bindDeprecatedOwnableUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DeprecatedOwnableUpgradeableCaller{contract: contract}, nil
}

// NewDeprecatedOwnableUpgradeableTransactor creates a new write-only instance of DeprecatedOwnableUpgradeable, bound to a specific deployed contract.
func NewDeprecatedOwnableUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*DeprecatedOwnableUpgradeableTransactor, error) {
	contract, err := bindDeprecatedOwnableUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DeprecatedOwnableUpgradeableTransactor{contract: contract}, nil
}

// NewDeprecatedOwnableUpgradeableFilterer creates a new log filterer instance of DeprecatedOwnableUpgradeable, bound to a specific deployed contract.
func NewDeprecatedOwnableUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*DeprecatedOwnableUpgradeableFilterer, error) {
	contract, err := bindDeprecatedOwnableUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DeprecatedOwnableUpgradeableFilterer{contract: contract}, nil
}

// bindDeprecatedOwnableUpgradeable binds a generic wrapper to an already deployed contract.
func bindDeprecatedOwnableUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DeprecatedOwnableUpgradeableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DeprecatedOwnableUpgradeable *DeprecatedOwnableUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DeprecatedOwnableUpgradeable.Contract.DeprecatedOwnableUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DeprecatedOwnableUpgradeable *DeprecatedOwnableUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DeprecatedOwnableUpgradeable.Contract.DeprecatedOwnableUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DeprecatedOwnableUpgradeable *DeprecatedOwnableUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DeprecatedOwnableUpgradeable.Contract.DeprecatedOwnableUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DeprecatedOwnableUpgradeable *DeprecatedOwnableUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DeprecatedOwnableUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DeprecatedOwnableUpgradeable *DeprecatedOwnableUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DeprecatedOwnableUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DeprecatedOwnableUpgradeable *DeprecatedOwnableUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DeprecatedOwnableUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// DeprecatedOwnableUpgradeableInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the DeprecatedOwnableUpgradeable contract.
type DeprecatedOwnableUpgradeableInitializedIterator struct {
	Event *DeprecatedOwnableUpgradeableInitialized // Event containing the contract specifics and raw log

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
func (it *DeprecatedOwnableUpgradeableInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeprecatedOwnableUpgradeableInitialized)
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
		it.Event = new(DeprecatedOwnableUpgradeableInitialized)
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
func (it *DeprecatedOwnableUpgradeableInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeprecatedOwnableUpgradeableInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeprecatedOwnableUpgradeableInitialized represents a Initialized event raised by the DeprecatedOwnableUpgradeable contract.
type DeprecatedOwnableUpgradeableInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_DeprecatedOwnableUpgradeable *DeprecatedOwnableUpgradeableFilterer) FilterInitialized(opts *bind.FilterOpts) (*DeprecatedOwnableUpgradeableInitializedIterator, error) {

	logs, sub, err := _DeprecatedOwnableUpgradeable.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &DeprecatedOwnableUpgradeableInitializedIterator{contract: _DeprecatedOwnableUpgradeable.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_DeprecatedOwnableUpgradeable *DeprecatedOwnableUpgradeableFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *DeprecatedOwnableUpgradeableInitialized) (event.Subscription, error) {

	logs, sub, err := _DeprecatedOwnableUpgradeable.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeprecatedOwnableUpgradeableInitialized)
				if err := _DeprecatedOwnableUpgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_DeprecatedOwnableUpgradeable *DeprecatedOwnableUpgradeableFilterer) ParseInitialized(log types.Log) (*DeprecatedOwnableUpgradeableInitialized, error) {
	event := new(DeprecatedOwnableUpgradeableInitialized)
	if err := _DeprecatedOwnableUpgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
