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
	Bin: "0x608060405234801561000f575f5ffd5b506040516105c83803806105c883398101604081905261002e916101c2565b5f5b825181101561006b5761006383828151811061004e5761004e61029e565b6020026020010151600161007c60201b60201c565b600101610030565b5061007581610103565b50506102b2565b6001600160a01b0382166100a3576040516339b190bb60e11b815260040160405180910390fd5b6001600160a01b0382165f8181526020818152604091829020805460ff19168515159081179091558251938452908301527f65d3a1fd4c13f05cba164f80d03ce90fb4b5e21946bfc3ab7dbd434c2d0b9152910160405180910390a15050565b6001600160a01b03811661012a576040516339b190bb60e11b815260040160405180910390fd5b600154604080516001600160a01b03928316815291831660208301527f06b4167a2528887a1e97a366eefe8549bfbf1ea3e6ac81cb2564a934d20e8892910160405180910390a1600180546001600160a01b0319166001600160a01b0392909216919091179055565b634e487b7160e01b5f52604160045260245ffd5b80516001600160a01b03811681146101bd575f5ffd5b919050565b5f5f604083850312156101d3575f5ffd5b82516001600160401b038111156101e8575f5ffd5b8301601f810185136101f8575f5ffd5b80516001600160401b0381111561021157610211610193565b604051600582901b90603f8201601f191681016001600160401b038111828210171561023f5761023f610193565b60405291825260208184018101929081018884111561025c575f5ffd5b6020850194505b8385101561028257610274856101a7565b815260209485019401610263565b50945061029592505050602084016101a7565b90509250929050565b634e487b7160e01b5f52603260045260245ffd5b610309806102bf5f395ff3fe608060405234801561000f575f5ffd5b506004361061004a575f3560e01c806346fbf68e1461004e5780638568520614610085578063ce5484281461009a578063eab66d7a146100ad575b5f5ffd5b61007061005c36600461027a565b5f6020819052908152604090205460ff1681565b60405190151581526020015b60405180910390f35b61009861009336600461029a565b6100d8565b005b6100986100a836600461027a565b610111565b6001546100c0906001600160a01b031681565b6040516001600160a01b03909116815260200161007c565b6001546001600160a01b031633146101035760405163794821ff60e01b815260040160405180910390fd5b61010d8282610148565b5050565b6001546001600160a01b0316331461013c5760405163794821ff60e01b815260040160405180910390fd5b610145816101cf565b50565b6001600160a01b03821661016f576040516339b190bb60e11b815260040160405180910390fd5b6001600160a01b0382165f8181526020818152604091829020805460ff19168515159081179091558251938452908301527f65d3a1fd4c13f05cba164f80d03ce90fb4b5e21946bfc3ab7dbd434c2d0b9152910160405180910390a15050565b6001600160a01b0381166101f6576040516339b190bb60e11b815260040160405180910390fd5b600154604080516001600160a01b03928316815291831660208301527f06b4167a2528887a1e97a366eefe8549bfbf1ea3e6ac81cb2564a934d20e8892910160405180910390a1600180546001600160a01b0319166001600160a01b0392909216919091179055565b80356001600160a01b0381168114610275575f5ffd5b919050565b5f6020828403121561028a575f5ffd5b6102938261025f565b9392505050565b5f5f604083850312156102ab575f5ffd5b6102b48361025f565b9150602083013580151581146102c8575f5ffd5b80915050925092905056fea26469706673582212202917159df0783c456a04923c0900b80ef8974eef653605675c7597eb49e08fe864736f6c634300081b0033",
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
