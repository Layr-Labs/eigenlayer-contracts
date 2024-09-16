// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IPauserRegistry

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

// IPauserRegistryMetaData contains all meta data concerning the IPauserRegistry contract.
var IPauserRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"isPauser\",\"inputs\":[{\"name\":\"pauser\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpauser\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"PauserStatusChanged\",\"inputs\":[{\"name\":\"pauser\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"canPause\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UnpauserChanged\",\"inputs\":[{\"name\":\"previousUnpauser\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newUnpauser\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InputAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnpauser\",\"inputs\":[]}]",
}

// IPauserRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use IPauserRegistryMetaData.ABI instead.
var IPauserRegistryABI = IPauserRegistryMetaData.ABI

// IPauserRegistry is an auto generated Go binding around an Ethereum contract.
type IPauserRegistry struct {
	IPauserRegistryCaller     // Read-only binding to the contract
	IPauserRegistryTransactor // Write-only binding to the contract
	IPauserRegistryFilterer   // Log filterer for contract events
}

// IPauserRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPauserRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPauserRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPauserRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPauserRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPauserRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPauserRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPauserRegistrySession struct {
	Contract     *IPauserRegistry  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPauserRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPauserRegistryCallerSession struct {
	Contract *IPauserRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IPauserRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPauserRegistryTransactorSession struct {
	Contract     *IPauserRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IPauserRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPauserRegistryRaw struct {
	Contract *IPauserRegistry // Generic contract binding to access the raw methods on
}

// IPauserRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPauserRegistryCallerRaw struct {
	Contract *IPauserRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// IPauserRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPauserRegistryTransactorRaw struct {
	Contract *IPauserRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPauserRegistry creates a new instance of IPauserRegistry, bound to a specific deployed contract.
func NewIPauserRegistry(address common.Address, backend bind.ContractBackend) (*IPauserRegistry, error) {
	contract, err := bindIPauserRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPauserRegistry{IPauserRegistryCaller: IPauserRegistryCaller{contract: contract}, IPauserRegistryTransactor: IPauserRegistryTransactor{contract: contract}, IPauserRegistryFilterer: IPauserRegistryFilterer{contract: contract}}, nil
}

// NewIPauserRegistryCaller creates a new read-only instance of IPauserRegistry, bound to a specific deployed contract.
func NewIPauserRegistryCaller(address common.Address, caller bind.ContractCaller) (*IPauserRegistryCaller, error) {
	contract, err := bindIPauserRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPauserRegistryCaller{contract: contract}, nil
}

// NewIPauserRegistryTransactor creates a new write-only instance of IPauserRegistry, bound to a specific deployed contract.
func NewIPauserRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*IPauserRegistryTransactor, error) {
	contract, err := bindIPauserRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPauserRegistryTransactor{contract: contract}, nil
}

// NewIPauserRegistryFilterer creates a new log filterer instance of IPauserRegistry, bound to a specific deployed contract.
func NewIPauserRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*IPauserRegistryFilterer, error) {
	contract, err := bindIPauserRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPauserRegistryFilterer{contract: contract}, nil
}

// bindIPauserRegistry binds a generic wrapper to an already deployed contract.
func bindIPauserRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IPauserRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPauserRegistry *IPauserRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPauserRegistry.Contract.IPauserRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPauserRegistry *IPauserRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPauserRegistry.Contract.IPauserRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPauserRegistry *IPauserRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPauserRegistry.Contract.IPauserRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPauserRegistry *IPauserRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPauserRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPauserRegistry *IPauserRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPauserRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPauserRegistry *IPauserRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPauserRegistry.Contract.contract.Transact(opts, method, params...)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address pauser) view returns(bool)
func (_IPauserRegistry *IPauserRegistryCaller) IsPauser(opts *bind.CallOpts, pauser common.Address) (bool, error) {
	var out []interface{}
	err := _IPauserRegistry.contract.Call(opts, &out, "isPauser", pauser)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address pauser) view returns(bool)
func (_IPauserRegistry *IPauserRegistrySession) IsPauser(pauser common.Address) (bool, error) {
	return _IPauserRegistry.Contract.IsPauser(&_IPauserRegistry.CallOpts, pauser)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address pauser) view returns(bool)
func (_IPauserRegistry *IPauserRegistryCallerSession) IsPauser(pauser common.Address) (bool, error) {
	return _IPauserRegistry.Contract.IsPauser(&_IPauserRegistry.CallOpts, pauser)
}

// Unpauser is a free data retrieval call binding the contract method 0xeab66d7a.
//
// Solidity: function unpauser() view returns(address)
func (_IPauserRegistry *IPauserRegistryCaller) Unpauser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IPauserRegistry.contract.Call(opts, &out, "unpauser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Unpauser is a free data retrieval call binding the contract method 0xeab66d7a.
//
// Solidity: function unpauser() view returns(address)
func (_IPauserRegistry *IPauserRegistrySession) Unpauser() (common.Address, error) {
	return _IPauserRegistry.Contract.Unpauser(&_IPauserRegistry.CallOpts)
}

// Unpauser is a free data retrieval call binding the contract method 0xeab66d7a.
//
// Solidity: function unpauser() view returns(address)
func (_IPauserRegistry *IPauserRegistryCallerSession) Unpauser() (common.Address, error) {
	return _IPauserRegistry.Contract.Unpauser(&_IPauserRegistry.CallOpts)
}

// IPauserRegistryPauserStatusChangedIterator is returned from FilterPauserStatusChanged and is used to iterate over the raw logs and unpacked data for PauserStatusChanged events raised by the IPauserRegistry contract.
type IPauserRegistryPauserStatusChangedIterator struct {
	Event *IPauserRegistryPauserStatusChanged // Event containing the contract specifics and raw log

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
func (it *IPauserRegistryPauserStatusChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPauserRegistryPauserStatusChanged)
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
		it.Event = new(IPauserRegistryPauserStatusChanged)
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
func (it *IPauserRegistryPauserStatusChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPauserRegistryPauserStatusChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPauserRegistryPauserStatusChanged represents a PauserStatusChanged event raised by the IPauserRegistry contract.
type IPauserRegistryPauserStatusChanged struct {
	Pauser   common.Address
	CanPause bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPauserStatusChanged is a free log retrieval operation binding the contract event 0x65d3a1fd4c13f05cba164f80d03ce90fb4b5e21946bfc3ab7dbd434c2d0b9152.
//
// Solidity: event PauserStatusChanged(address pauser, bool canPause)
func (_IPauserRegistry *IPauserRegistryFilterer) FilterPauserStatusChanged(opts *bind.FilterOpts) (*IPauserRegistryPauserStatusChangedIterator, error) {

	logs, sub, err := _IPauserRegistry.contract.FilterLogs(opts, "PauserStatusChanged")
	if err != nil {
		return nil, err
	}
	return &IPauserRegistryPauserStatusChangedIterator{contract: _IPauserRegistry.contract, event: "PauserStatusChanged", logs: logs, sub: sub}, nil
}

// WatchPauserStatusChanged is a free log subscription operation binding the contract event 0x65d3a1fd4c13f05cba164f80d03ce90fb4b5e21946bfc3ab7dbd434c2d0b9152.
//
// Solidity: event PauserStatusChanged(address pauser, bool canPause)
func (_IPauserRegistry *IPauserRegistryFilterer) WatchPauserStatusChanged(opts *bind.WatchOpts, sink chan<- *IPauserRegistryPauserStatusChanged) (event.Subscription, error) {

	logs, sub, err := _IPauserRegistry.contract.WatchLogs(opts, "PauserStatusChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPauserRegistryPauserStatusChanged)
				if err := _IPauserRegistry.contract.UnpackLog(event, "PauserStatusChanged", log); err != nil {
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

// ParsePauserStatusChanged is a log parse operation binding the contract event 0x65d3a1fd4c13f05cba164f80d03ce90fb4b5e21946bfc3ab7dbd434c2d0b9152.
//
// Solidity: event PauserStatusChanged(address pauser, bool canPause)
func (_IPauserRegistry *IPauserRegistryFilterer) ParsePauserStatusChanged(log types.Log) (*IPauserRegistryPauserStatusChanged, error) {
	event := new(IPauserRegistryPauserStatusChanged)
	if err := _IPauserRegistry.contract.UnpackLog(event, "PauserStatusChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPauserRegistryUnpauserChangedIterator is returned from FilterUnpauserChanged and is used to iterate over the raw logs and unpacked data for UnpauserChanged events raised by the IPauserRegistry contract.
type IPauserRegistryUnpauserChangedIterator struct {
	Event *IPauserRegistryUnpauserChanged // Event containing the contract specifics and raw log

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
func (it *IPauserRegistryUnpauserChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPauserRegistryUnpauserChanged)
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
		it.Event = new(IPauserRegistryUnpauserChanged)
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
func (it *IPauserRegistryUnpauserChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPauserRegistryUnpauserChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPauserRegistryUnpauserChanged represents a UnpauserChanged event raised by the IPauserRegistry contract.
type IPauserRegistryUnpauserChanged struct {
	PreviousUnpauser common.Address
	NewUnpauser      common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterUnpauserChanged is a free log retrieval operation binding the contract event 0x06b4167a2528887a1e97a366eefe8549bfbf1ea3e6ac81cb2564a934d20e8892.
//
// Solidity: event UnpauserChanged(address previousUnpauser, address newUnpauser)
func (_IPauserRegistry *IPauserRegistryFilterer) FilterUnpauserChanged(opts *bind.FilterOpts) (*IPauserRegistryUnpauserChangedIterator, error) {

	logs, sub, err := _IPauserRegistry.contract.FilterLogs(opts, "UnpauserChanged")
	if err != nil {
		return nil, err
	}
	return &IPauserRegistryUnpauserChangedIterator{contract: _IPauserRegistry.contract, event: "UnpauserChanged", logs: logs, sub: sub}, nil
}

// WatchUnpauserChanged is a free log subscription operation binding the contract event 0x06b4167a2528887a1e97a366eefe8549bfbf1ea3e6ac81cb2564a934d20e8892.
//
// Solidity: event UnpauserChanged(address previousUnpauser, address newUnpauser)
func (_IPauserRegistry *IPauserRegistryFilterer) WatchUnpauserChanged(opts *bind.WatchOpts, sink chan<- *IPauserRegistryUnpauserChanged) (event.Subscription, error) {

	logs, sub, err := _IPauserRegistry.contract.WatchLogs(opts, "UnpauserChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPauserRegistryUnpauserChanged)
				if err := _IPauserRegistry.contract.UnpackLog(event, "UnpauserChanged", log); err != nil {
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

// ParseUnpauserChanged is a log parse operation binding the contract event 0x06b4167a2528887a1e97a366eefe8549bfbf1ea3e6ac81cb2564a934d20e8892.
//
// Solidity: event UnpauserChanged(address previousUnpauser, address newUnpauser)
func (_IPauserRegistry *IPauserRegistryFilterer) ParseUnpauserChanged(log types.Log) (*IPauserRegistryUnpauserChanged, error) {
	event := new(IPauserRegistryUnpauserChanged)
	if err := _IPauserRegistry.contract.UnpackLog(event, "UnpauserChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
