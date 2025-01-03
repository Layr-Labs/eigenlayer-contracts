// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package PauserRegistry

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

// PauserRegistryMetaData contains all meta data concerning the PauserRegistry contract.
var PauserRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_pausers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_unpauser\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isPauser\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setIsPauser\",\"inputs\":[{\"name\":\"newPauser\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"canPause\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setUnpauser\",\"inputs\":[{\"name\":\"newUnpauser\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpauser\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"PauserStatusChanged\",\"inputs\":[{\"name\":\"pauser\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"canPause\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UnpauserChanged\",\"inputs\":[{\"name\":\"previousUnpauser\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newUnpauser\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InputAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnpauser\",\"inputs\":[]}]",
	Bin: "0x608060405234801561001057600080fd5b506040516105e23803806105e283398101604081905261002f916101c8565b60005b825181101561006d57610065838281518110610050576100506102a9565b6020026020010151600161007e60201b60201c565b600101610032565b5061007781610106565b50506102bf565b6001600160a01b0382166100a5576040516339b190bb60e11b815260040160405180910390fd5b6001600160a01b03821660008181526020818152604091829020805460ff19168515159081179091558251938452908301527f65d3a1fd4c13f05cba164f80d03ce90fb4b5e21946bfc3ab7dbd434c2d0b9152910160405180910390a15050565b6001600160a01b03811661012d576040516339b190bb60e11b815260040160405180910390fd5b600154604080516001600160a01b03928316815291831660208301527f06b4167a2528887a1e97a366eefe8549bfbf1ea3e6ac81cb2564a934d20e8892910160405180910390a1600180546001600160a01b0319166001600160a01b0392909216919091179055565b634e487b7160e01b600052604160045260246000fd5b80516001600160a01b03811681146101c357600080fd5b919050565b600080604083850312156101db57600080fd5b82516001600160401b038111156101f157600080fd5b8301601f8101851361020257600080fd5b80516001600160401b0381111561021b5761021b610196565b604051600582901b90603f8201601f191681016001600160401b038111828210171561024957610249610196565b60405291825260208184018101929081018884111561026757600080fd5b6020850194505b8385101561028d5761027f856101ac565b81526020948501940161026e565b5094506102a092505050602084016101ac565b90509250929050565b634e487b7160e01b600052603260045260246000fd5b610314806102ce6000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806346fbf68e146100515780638568520614610089578063ce5484281461009e578063eab66d7a146100b1575b600080fd5b61007461005f366004610280565b60006020819052908152604090205460ff1681565b60405190151581526020015b60405180910390f35b61009c6100973660046102a2565b6100dc565b005b61009c6100ac366004610280565b610115565b6001546100c4906001600160a01b031681565b6040516001600160a01b039091168152602001610080565b6001546001600160a01b031633146101075760405163794821ff60e01b815260040160405180910390fd5b610111828261014c565b5050565b6001546001600160a01b031633146101405760405163794821ff60e01b815260040160405180910390fd5b610149816101d4565b50565b6001600160a01b038216610173576040516339b190bb60e11b815260040160405180910390fd5b6001600160a01b03821660008181526020818152604091829020805460ff19168515159081179091558251938452908301527f65d3a1fd4c13f05cba164f80d03ce90fb4b5e21946bfc3ab7dbd434c2d0b9152910160405180910390a15050565b6001600160a01b0381166101fb576040516339b190bb60e11b815260040160405180910390fd5b600154604080516001600160a01b03928316815291831660208301527f06b4167a2528887a1e97a366eefe8549bfbf1ea3e6ac81cb2564a934d20e8892910160405180910390a1600180546001600160a01b0319166001600160a01b0392909216919091179055565b80356001600160a01b038116811461027b57600080fd5b919050565b60006020828403121561029257600080fd5b61029b82610264565b9392505050565b600080604083850312156102b557600080fd5b6102be83610264565b9150602083013580151581146102d357600080fd5b80915050925092905056fea2646970667358221220e165396a512ff2b6e2f77852623a0f3ee43d1eca10d557ff6863e2fc122b475c64736f6c634300081b0033",
}

// PauserRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use PauserRegistryMetaData.ABI instead.
var PauserRegistryABI = PauserRegistryMetaData.ABI

// PauserRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PauserRegistryMetaData.Bin instead.
var PauserRegistryBin = PauserRegistryMetaData.Bin

// DeployPauserRegistry deploys a new Ethereum contract, binding an instance of PauserRegistry to it.
func DeployPauserRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, _pausers []common.Address, _unpauser common.Address) (common.Address, *types.Transaction, *PauserRegistry, error) {
	parsed, err := PauserRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PauserRegistryBin), backend, _pausers, _unpauser)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PauserRegistry{PauserRegistryCaller: PauserRegistryCaller{contract: contract}, PauserRegistryTransactor: PauserRegistryTransactor{contract: contract}, PauserRegistryFilterer: PauserRegistryFilterer{contract: contract}}, nil
}

// PauserRegistry is an auto generated Go binding around an Ethereum contract.
type PauserRegistry struct {
	PauserRegistryCaller     // Read-only binding to the contract
	PauserRegistryTransactor // Write-only binding to the contract
	PauserRegistryFilterer   // Log filterer for contract events
}

// PauserRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type PauserRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PauserRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PauserRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PauserRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PauserRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PauserRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PauserRegistrySession struct {
	Contract     *PauserRegistry   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PauserRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PauserRegistryCallerSession struct {
	Contract *PauserRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// PauserRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PauserRegistryTransactorSession struct {
	Contract     *PauserRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// PauserRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type PauserRegistryRaw struct {
	Contract *PauserRegistry // Generic contract binding to access the raw methods on
}

// PauserRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PauserRegistryCallerRaw struct {
	Contract *PauserRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// PauserRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PauserRegistryTransactorRaw struct {
	Contract *PauserRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPauserRegistry creates a new instance of PauserRegistry, bound to a specific deployed contract.
func NewPauserRegistry(address common.Address, backend bind.ContractBackend) (*PauserRegistry, error) {
	contract, err := bindPauserRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PauserRegistry{PauserRegistryCaller: PauserRegistryCaller{contract: contract}, PauserRegistryTransactor: PauserRegistryTransactor{contract: contract}, PauserRegistryFilterer: PauserRegistryFilterer{contract: contract}}, nil
}

// NewPauserRegistryCaller creates a new read-only instance of PauserRegistry, bound to a specific deployed contract.
func NewPauserRegistryCaller(address common.Address, caller bind.ContractCaller) (*PauserRegistryCaller, error) {
	contract, err := bindPauserRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PauserRegistryCaller{contract: contract}, nil
}

// NewPauserRegistryTransactor creates a new write-only instance of PauserRegistry, bound to a specific deployed contract.
func NewPauserRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*PauserRegistryTransactor, error) {
	contract, err := bindPauserRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PauserRegistryTransactor{contract: contract}, nil
}

// NewPauserRegistryFilterer creates a new log filterer instance of PauserRegistry, bound to a specific deployed contract.
func NewPauserRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*PauserRegistryFilterer, error) {
	contract, err := bindPauserRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PauserRegistryFilterer{contract: contract}, nil
}

// bindPauserRegistry binds a generic wrapper to an already deployed contract.
func bindPauserRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PauserRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PauserRegistry *PauserRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PauserRegistry.Contract.PauserRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PauserRegistry *PauserRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PauserRegistry.Contract.PauserRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PauserRegistry *PauserRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PauserRegistry.Contract.PauserRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PauserRegistry *PauserRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PauserRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PauserRegistry *PauserRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PauserRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PauserRegistry *PauserRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PauserRegistry.Contract.contract.Transact(opts, method, params...)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address ) view returns(bool)
func (_PauserRegistry *PauserRegistryCaller) IsPauser(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _PauserRegistry.contract.Call(opts, &out, "isPauser", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address ) view returns(bool)
func (_PauserRegistry *PauserRegistrySession) IsPauser(arg0 common.Address) (bool, error) {
	return _PauserRegistry.Contract.IsPauser(&_PauserRegistry.CallOpts, arg0)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address ) view returns(bool)
func (_PauserRegistry *PauserRegistryCallerSession) IsPauser(arg0 common.Address) (bool, error) {
	return _PauserRegistry.Contract.IsPauser(&_PauserRegistry.CallOpts, arg0)
}

// Unpauser is a free data retrieval call binding the contract method 0xeab66d7a.
//
// Solidity: function unpauser() view returns(address)
func (_PauserRegistry *PauserRegistryCaller) Unpauser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PauserRegistry.contract.Call(opts, &out, "unpauser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Unpauser is a free data retrieval call binding the contract method 0xeab66d7a.
//
// Solidity: function unpauser() view returns(address)
func (_PauserRegistry *PauserRegistrySession) Unpauser() (common.Address, error) {
	return _PauserRegistry.Contract.Unpauser(&_PauserRegistry.CallOpts)
}

// Unpauser is a free data retrieval call binding the contract method 0xeab66d7a.
//
// Solidity: function unpauser() view returns(address)
func (_PauserRegistry *PauserRegistryCallerSession) Unpauser() (common.Address, error) {
	return _PauserRegistry.Contract.Unpauser(&_PauserRegistry.CallOpts)
}

// SetIsPauser is a paid mutator transaction binding the contract method 0x85685206.
//
// Solidity: function setIsPauser(address newPauser, bool canPause) returns()
func (_PauserRegistry *PauserRegistryTransactor) SetIsPauser(opts *bind.TransactOpts, newPauser common.Address, canPause bool) (*types.Transaction, error) {
	return _PauserRegistry.contract.Transact(opts, "setIsPauser", newPauser, canPause)
}

// SetIsPauser is a paid mutator transaction binding the contract method 0x85685206.
//
// Solidity: function setIsPauser(address newPauser, bool canPause) returns()
func (_PauserRegistry *PauserRegistrySession) SetIsPauser(newPauser common.Address, canPause bool) (*types.Transaction, error) {
	return _PauserRegistry.Contract.SetIsPauser(&_PauserRegistry.TransactOpts, newPauser, canPause)
}

// SetIsPauser is a paid mutator transaction binding the contract method 0x85685206.
//
// Solidity: function setIsPauser(address newPauser, bool canPause) returns()
func (_PauserRegistry *PauserRegistryTransactorSession) SetIsPauser(newPauser common.Address, canPause bool) (*types.Transaction, error) {
	return _PauserRegistry.Contract.SetIsPauser(&_PauserRegistry.TransactOpts, newPauser, canPause)
}

// SetUnpauser is a paid mutator transaction binding the contract method 0xce548428.
//
// Solidity: function setUnpauser(address newUnpauser) returns()
func (_PauserRegistry *PauserRegistryTransactor) SetUnpauser(opts *bind.TransactOpts, newUnpauser common.Address) (*types.Transaction, error) {
	return _PauserRegistry.contract.Transact(opts, "setUnpauser", newUnpauser)
}

// SetUnpauser is a paid mutator transaction binding the contract method 0xce548428.
//
// Solidity: function setUnpauser(address newUnpauser) returns()
func (_PauserRegistry *PauserRegistrySession) SetUnpauser(newUnpauser common.Address) (*types.Transaction, error) {
	return _PauserRegistry.Contract.SetUnpauser(&_PauserRegistry.TransactOpts, newUnpauser)
}

// SetUnpauser is a paid mutator transaction binding the contract method 0xce548428.
//
// Solidity: function setUnpauser(address newUnpauser) returns()
func (_PauserRegistry *PauserRegistryTransactorSession) SetUnpauser(newUnpauser common.Address) (*types.Transaction, error) {
	return _PauserRegistry.Contract.SetUnpauser(&_PauserRegistry.TransactOpts, newUnpauser)
}

// PauserRegistryPauserStatusChangedIterator is returned from FilterPauserStatusChanged and is used to iterate over the raw logs and unpacked data for PauserStatusChanged events raised by the PauserRegistry contract.
type PauserRegistryPauserStatusChangedIterator struct {
	Event *PauserRegistryPauserStatusChanged // Event containing the contract specifics and raw log

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
func (it *PauserRegistryPauserStatusChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PauserRegistryPauserStatusChanged)
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
		it.Event = new(PauserRegistryPauserStatusChanged)
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
func (it *PauserRegistryPauserStatusChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PauserRegistryPauserStatusChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PauserRegistryPauserStatusChanged represents a PauserStatusChanged event raised by the PauserRegistry contract.
type PauserRegistryPauserStatusChanged struct {
	Pauser   common.Address
	CanPause bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPauserStatusChanged is a free log retrieval operation binding the contract event 0x65d3a1fd4c13f05cba164f80d03ce90fb4b5e21946bfc3ab7dbd434c2d0b9152.
//
// Solidity: event PauserStatusChanged(address pauser, bool canPause)
func (_PauserRegistry *PauserRegistryFilterer) FilterPauserStatusChanged(opts *bind.FilterOpts) (*PauserRegistryPauserStatusChangedIterator, error) {

	logs, sub, err := _PauserRegistry.contract.FilterLogs(opts, "PauserStatusChanged")
	if err != nil {
		return nil, err
	}
	return &PauserRegistryPauserStatusChangedIterator{contract: _PauserRegistry.contract, event: "PauserStatusChanged", logs: logs, sub: sub}, nil
}

// WatchPauserStatusChanged is a free log subscription operation binding the contract event 0x65d3a1fd4c13f05cba164f80d03ce90fb4b5e21946bfc3ab7dbd434c2d0b9152.
//
// Solidity: event PauserStatusChanged(address pauser, bool canPause)
func (_PauserRegistry *PauserRegistryFilterer) WatchPauserStatusChanged(opts *bind.WatchOpts, sink chan<- *PauserRegistryPauserStatusChanged) (event.Subscription, error) {

	logs, sub, err := _PauserRegistry.contract.WatchLogs(opts, "PauserStatusChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PauserRegistryPauserStatusChanged)
				if err := _PauserRegistry.contract.UnpackLog(event, "PauserStatusChanged", log); err != nil {
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
func (_PauserRegistry *PauserRegistryFilterer) ParsePauserStatusChanged(log types.Log) (*PauserRegistryPauserStatusChanged, error) {
	event := new(PauserRegistryPauserStatusChanged)
	if err := _PauserRegistry.contract.UnpackLog(event, "PauserStatusChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PauserRegistryUnpauserChangedIterator is returned from FilterUnpauserChanged and is used to iterate over the raw logs and unpacked data for UnpauserChanged events raised by the PauserRegistry contract.
type PauserRegistryUnpauserChangedIterator struct {
	Event *PauserRegistryUnpauserChanged // Event containing the contract specifics and raw log

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
func (it *PauserRegistryUnpauserChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PauserRegistryUnpauserChanged)
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
		it.Event = new(PauserRegistryUnpauserChanged)
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
func (it *PauserRegistryUnpauserChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PauserRegistryUnpauserChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PauserRegistryUnpauserChanged represents a UnpauserChanged event raised by the PauserRegistry contract.
type PauserRegistryUnpauserChanged struct {
	PreviousUnpauser common.Address
	NewUnpauser      common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterUnpauserChanged is a free log retrieval operation binding the contract event 0x06b4167a2528887a1e97a366eefe8549bfbf1ea3e6ac81cb2564a934d20e8892.
//
// Solidity: event UnpauserChanged(address previousUnpauser, address newUnpauser)
func (_PauserRegistry *PauserRegistryFilterer) FilterUnpauserChanged(opts *bind.FilterOpts) (*PauserRegistryUnpauserChangedIterator, error) {

	logs, sub, err := _PauserRegistry.contract.FilterLogs(opts, "UnpauserChanged")
	if err != nil {
		return nil, err
	}
	return &PauserRegistryUnpauserChangedIterator{contract: _PauserRegistry.contract, event: "UnpauserChanged", logs: logs, sub: sub}, nil
}

// WatchUnpauserChanged is a free log subscription operation binding the contract event 0x06b4167a2528887a1e97a366eefe8549bfbf1ea3e6ac81cb2564a934d20e8892.
//
// Solidity: event UnpauserChanged(address previousUnpauser, address newUnpauser)
func (_PauserRegistry *PauserRegistryFilterer) WatchUnpauserChanged(opts *bind.WatchOpts, sink chan<- *PauserRegistryUnpauserChanged) (event.Subscription, error) {

	logs, sub, err := _PauserRegistry.contract.WatchLogs(opts, "UnpauserChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PauserRegistryUnpauserChanged)
				if err := _PauserRegistry.contract.UnpackLog(event, "UnpauserChanged", log); err != nil {
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
func (_PauserRegistry *PauserRegistryFilterer) ParseUnpauserChanged(log types.Log) (*PauserRegistryUnpauserChanged, error) {
	event := new(PauserRegistryUnpauserChanged)
	if err := _PauserRegistry.contract.UnpackLog(event, "UnpauserChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
