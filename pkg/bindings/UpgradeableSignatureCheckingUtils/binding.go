// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package UpgradeableSignatureCheckingUtils

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

// UpgradeableSignatureCheckingUtilsMetaData contains all meta data concerning the UpgradeableSignatureCheckingUtils contract.
var UpgradeableSignatureCheckingUtilsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"DOMAIN_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"domainSeparator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false}]",
}

// UpgradeableSignatureCheckingUtilsABI is the input ABI used to generate the binding from.
// Deprecated: Use UpgradeableSignatureCheckingUtilsMetaData.ABI instead.
var UpgradeableSignatureCheckingUtilsABI = UpgradeableSignatureCheckingUtilsMetaData.ABI

// UpgradeableSignatureCheckingUtils is an auto generated Go binding around an Ethereum contract.
type UpgradeableSignatureCheckingUtils struct {
	UpgradeableSignatureCheckingUtilsCaller     // Read-only binding to the contract
	UpgradeableSignatureCheckingUtilsTransactor // Write-only binding to the contract
	UpgradeableSignatureCheckingUtilsFilterer   // Log filterer for contract events
}

// UpgradeableSignatureCheckingUtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type UpgradeableSignatureCheckingUtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradeableSignatureCheckingUtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UpgradeableSignatureCheckingUtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradeableSignatureCheckingUtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UpgradeableSignatureCheckingUtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradeableSignatureCheckingUtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UpgradeableSignatureCheckingUtilsSession struct {
	Contract     *UpgradeableSignatureCheckingUtils // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                      // Call options to use throughout this session
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// UpgradeableSignatureCheckingUtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UpgradeableSignatureCheckingUtilsCallerSession struct {
	Contract *UpgradeableSignatureCheckingUtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                            // Call options to use throughout this session
}

// UpgradeableSignatureCheckingUtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UpgradeableSignatureCheckingUtilsTransactorSession struct {
	Contract     *UpgradeableSignatureCheckingUtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                            // Transaction auth options to use throughout this session
}

// UpgradeableSignatureCheckingUtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type UpgradeableSignatureCheckingUtilsRaw struct {
	Contract *UpgradeableSignatureCheckingUtils // Generic contract binding to access the raw methods on
}

// UpgradeableSignatureCheckingUtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UpgradeableSignatureCheckingUtilsCallerRaw struct {
	Contract *UpgradeableSignatureCheckingUtilsCaller // Generic read-only contract binding to access the raw methods on
}

// UpgradeableSignatureCheckingUtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UpgradeableSignatureCheckingUtilsTransactorRaw struct {
	Contract *UpgradeableSignatureCheckingUtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUpgradeableSignatureCheckingUtils creates a new instance of UpgradeableSignatureCheckingUtils, bound to a specific deployed contract.
func NewUpgradeableSignatureCheckingUtils(address common.Address, backend bind.ContractBackend) (*UpgradeableSignatureCheckingUtils, error) {
	contract, err := bindUpgradeableSignatureCheckingUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UpgradeableSignatureCheckingUtils{UpgradeableSignatureCheckingUtilsCaller: UpgradeableSignatureCheckingUtilsCaller{contract: contract}, UpgradeableSignatureCheckingUtilsTransactor: UpgradeableSignatureCheckingUtilsTransactor{contract: contract}, UpgradeableSignatureCheckingUtilsFilterer: UpgradeableSignatureCheckingUtilsFilterer{contract: contract}}, nil
}

// NewUpgradeableSignatureCheckingUtilsCaller creates a new read-only instance of UpgradeableSignatureCheckingUtils, bound to a specific deployed contract.
func NewUpgradeableSignatureCheckingUtilsCaller(address common.Address, caller bind.ContractCaller) (*UpgradeableSignatureCheckingUtilsCaller, error) {
	contract, err := bindUpgradeableSignatureCheckingUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UpgradeableSignatureCheckingUtilsCaller{contract: contract}, nil
}

// NewUpgradeableSignatureCheckingUtilsTransactor creates a new write-only instance of UpgradeableSignatureCheckingUtils, bound to a specific deployed contract.
func NewUpgradeableSignatureCheckingUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*UpgradeableSignatureCheckingUtilsTransactor, error) {
	contract, err := bindUpgradeableSignatureCheckingUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UpgradeableSignatureCheckingUtilsTransactor{contract: contract}, nil
}

// NewUpgradeableSignatureCheckingUtilsFilterer creates a new log filterer instance of UpgradeableSignatureCheckingUtils, bound to a specific deployed contract.
func NewUpgradeableSignatureCheckingUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*UpgradeableSignatureCheckingUtilsFilterer, error) {
	contract, err := bindUpgradeableSignatureCheckingUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UpgradeableSignatureCheckingUtilsFilterer{contract: contract}, nil
}

// bindUpgradeableSignatureCheckingUtils binds a generic wrapper to an already deployed contract.
func bindUpgradeableSignatureCheckingUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UpgradeableSignatureCheckingUtilsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UpgradeableSignatureCheckingUtils *UpgradeableSignatureCheckingUtilsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UpgradeableSignatureCheckingUtils.Contract.UpgradeableSignatureCheckingUtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UpgradeableSignatureCheckingUtils *UpgradeableSignatureCheckingUtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UpgradeableSignatureCheckingUtils.Contract.UpgradeableSignatureCheckingUtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UpgradeableSignatureCheckingUtils *UpgradeableSignatureCheckingUtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UpgradeableSignatureCheckingUtils.Contract.UpgradeableSignatureCheckingUtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UpgradeableSignatureCheckingUtils *UpgradeableSignatureCheckingUtilsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UpgradeableSignatureCheckingUtils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UpgradeableSignatureCheckingUtils *UpgradeableSignatureCheckingUtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UpgradeableSignatureCheckingUtils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UpgradeableSignatureCheckingUtils *UpgradeableSignatureCheckingUtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UpgradeableSignatureCheckingUtils.Contract.contract.Transact(opts, method, params...)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_UpgradeableSignatureCheckingUtils *UpgradeableSignatureCheckingUtilsCaller) DOMAINTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _UpgradeableSignatureCheckingUtils.contract.Call(opts, &out, "DOMAIN_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_UpgradeableSignatureCheckingUtils *UpgradeableSignatureCheckingUtilsSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _UpgradeableSignatureCheckingUtils.Contract.DOMAINTYPEHASH(&_UpgradeableSignatureCheckingUtils.CallOpts)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_UpgradeableSignatureCheckingUtils *UpgradeableSignatureCheckingUtilsCallerSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _UpgradeableSignatureCheckingUtils.Contract.DOMAINTYPEHASH(&_UpgradeableSignatureCheckingUtils.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_UpgradeableSignatureCheckingUtils *UpgradeableSignatureCheckingUtilsCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _UpgradeableSignatureCheckingUtils.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_UpgradeableSignatureCheckingUtils *UpgradeableSignatureCheckingUtilsSession) DomainSeparator() ([32]byte, error) {
	return _UpgradeableSignatureCheckingUtils.Contract.DomainSeparator(&_UpgradeableSignatureCheckingUtils.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_UpgradeableSignatureCheckingUtils *UpgradeableSignatureCheckingUtilsCallerSession) DomainSeparator() ([32]byte, error) {
	return _UpgradeableSignatureCheckingUtils.Contract.DomainSeparator(&_UpgradeableSignatureCheckingUtils.CallOpts)
}

// UpgradeableSignatureCheckingUtilsInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the UpgradeableSignatureCheckingUtils contract.
type UpgradeableSignatureCheckingUtilsInitializedIterator struct {
	Event *UpgradeableSignatureCheckingUtilsInitialized // Event containing the contract specifics and raw log

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
func (it *UpgradeableSignatureCheckingUtilsInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeableSignatureCheckingUtilsInitialized)
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
		it.Event = new(UpgradeableSignatureCheckingUtilsInitialized)
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
func (it *UpgradeableSignatureCheckingUtilsInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeableSignatureCheckingUtilsInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeableSignatureCheckingUtilsInitialized represents a Initialized event raised by the UpgradeableSignatureCheckingUtils contract.
type UpgradeableSignatureCheckingUtilsInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_UpgradeableSignatureCheckingUtils *UpgradeableSignatureCheckingUtilsFilterer) FilterInitialized(opts *bind.FilterOpts) (*UpgradeableSignatureCheckingUtilsInitializedIterator, error) {

	logs, sub, err := _UpgradeableSignatureCheckingUtils.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &UpgradeableSignatureCheckingUtilsInitializedIterator{contract: _UpgradeableSignatureCheckingUtils.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_UpgradeableSignatureCheckingUtils *UpgradeableSignatureCheckingUtilsFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *UpgradeableSignatureCheckingUtilsInitialized) (event.Subscription, error) {

	logs, sub, err := _UpgradeableSignatureCheckingUtils.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeableSignatureCheckingUtilsInitialized)
				if err := _UpgradeableSignatureCheckingUtils.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_UpgradeableSignatureCheckingUtils *UpgradeableSignatureCheckingUtilsFilterer) ParseInitialized(log types.Log) (*UpgradeableSignatureCheckingUtilsInitialized, error) {
	event := new(UpgradeableSignatureCheckingUtilsInitialized)
	if err := _UpgradeableSignatureCheckingUtils.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
