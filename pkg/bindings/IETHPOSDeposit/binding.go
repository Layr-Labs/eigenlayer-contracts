// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IETHPOSDeposit

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

// IETHPOSDepositMetaData contains all meta data concerning the IETHPOSDeposit contract.
var IETHPOSDepositMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"pubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"withdrawal_credentials\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"deposit_data_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"get_deposit_count\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"get_deposit_root\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"DepositEvent\",\"inputs\":[{\"name\":\"pubkey\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"withdrawal_credentials\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"index\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false}]",
}

// IETHPOSDepositABI is the input ABI used to generate the binding from.
// Deprecated: Use IETHPOSDepositMetaData.ABI instead.
var IETHPOSDepositABI = IETHPOSDepositMetaData.ABI

// IETHPOSDeposit is an auto generated Go binding around an Ethereum contract.
type IETHPOSDeposit struct {
	IETHPOSDepositCaller     // Read-only binding to the contract
	IETHPOSDepositTransactor // Write-only binding to the contract
	IETHPOSDepositFilterer   // Log filterer for contract events
}

// IETHPOSDepositCaller is an auto generated read-only Go binding around an Ethereum contract.
type IETHPOSDepositCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IETHPOSDepositTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IETHPOSDepositTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IETHPOSDepositFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IETHPOSDepositFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IETHPOSDepositSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IETHPOSDepositSession struct {
	Contract     *IETHPOSDeposit   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IETHPOSDepositCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IETHPOSDepositCallerSession struct {
	Contract *IETHPOSDepositCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IETHPOSDepositTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IETHPOSDepositTransactorSession struct {
	Contract     *IETHPOSDepositTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IETHPOSDepositRaw is an auto generated low-level Go binding around an Ethereum contract.
type IETHPOSDepositRaw struct {
	Contract *IETHPOSDeposit // Generic contract binding to access the raw methods on
}

// IETHPOSDepositCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IETHPOSDepositCallerRaw struct {
	Contract *IETHPOSDepositCaller // Generic read-only contract binding to access the raw methods on
}

// IETHPOSDepositTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IETHPOSDepositTransactorRaw struct {
	Contract *IETHPOSDepositTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIETHPOSDeposit creates a new instance of IETHPOSDeposit, bound to a specific deployed contract.
func NewIETHPOSDeposit(address common.Address, backend bind.ContractBackend) (*IETHPOSDeposit, error) {
	contract, err := bindIETHPOSDeposit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IETHPOSDeposit{IETHPOSDepositCaller: IETHPOSDepositCaller{contract: contract}, IETHPOSDepositTransactor: IETHPOSDepositTransactor{contract: contract}, IETHPOSDepositFilterer: IETHPOSDepositFilterer{contract: contract}}, nil
}

// NewIETHPOSDepositCaller creates a new read-only instance of IETHPOSDeposit, bound to a specific deployed contract.
func NewIETHPOSDepositCaller(address common.Address, caller bind.ContractCaller) (*IETHPOSDepositCaller, error) {
	contract, err := bindIETHPOSDeposit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IETHPOSDepositCaller{contract: contract}, nil
}

// NewIETHPOSDepositTransactor creates a new write-only instance of IETHPOSDeposit, bound to a specific deployed contract.
func NewIETHPOSDepositTransactor(address common.Address, transactor bind.ContractTransactor) (*IETHPOSDepositTransactor, error) {
	contract, err := bindIETHPOSDeposit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IETHPOSDepositTransactor{contract: contract}, nil
}

// NewIETHPOSDepositFilterer creates a new log filterer instance of IETHPOSDeposit, bound to a specific deployed contract.
func NewIETHPOSDepositFilterer(address common.Address, filterer bind.ContractFilterer) (*IETHPOSDepositFilterer, error) {
	contract, err := bindIETHPOSDeposit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IETHPOSDepositFilterer{contract: contract}, nil
}

// bindIETHPOSDeposit binds a generic wrapper to an already deployed contract.
func bindIETHPOSDeposit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IETHPOSDepositMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IETHPOSDeposit *IETHPOSDepositRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IETHPOSDeposit.Contract.IETHPOSDepositCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IETHPOSDeposit *IETHPOSDepositRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IETHPOSDeposit.Contract.IETHPOSDepositTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IETHPOSDeposit *IETHPOSDepositRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IETHPOSDeposit.Contract.IETHPOSDepositTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IETHPOSDeposit *IETHPOSDepositCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IETHPOSDeposit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IETHPOSDeposit *IETHPOSDepositTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IETHPOSDeposit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IETHPOSDeposit *IETHPOSDepositTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IETHPOSDeposit.Contract.contract.Transact(opts, method, params...)
}

// GetDepositCount is a free data retrieval call binding the contract method 0x621fd130.
//
// Solidity: function get_deposit_count() view returns(bytes)
func (_IETHPOSDeposit *IETHPOSDepositCaller) GetDepositCount(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _IETHPOSDeposit.contract.Call(opts, &out, "get_deposit_count")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetDepositCount is a free data retrieval call binding the contract method 0x621fd130.
//
// Solidity: function get_deposit_count() view returns(bytes)
func (_IETHPOSDeposit *IETHPOSDepositSession) GetDepositCount() ([]byte, error) {
	return _IETHPOSDeposit.Contract.GetDepositCount(&_IETHPOSDeposit.CallOpts)
}

// GetDepositCount is a free data retrieval call binding the contract method 0x621fd130.
//
// Solidity: function get_deposit_count() view returns(bytes)
func (_IETHPOSDeposit *IETHPOSDepositCallerSession) GetDepositCount() ([]byte, error) {
	return _IETHPOSDeposit.Contract.GetDepositCount(&_IETHPOSDeposit.CallOpts)
}

// GetDepositRoot is a free data retrieval call binding the contract method 0xc5f2892f.
//
// Solidity: function get_deposit_root() view returns(bytes32)
func (_IETHPOSDeposit *IETHPOSDepositCaller) GetDepositRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IETHPOSDeposit.contract.Call(opts, &out, "get_deposit_root")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDepositRoot is a free data retrieval call binding the contract method 0xc5f2892f.
//
// Solidity: function get_deposit_root() view returns(bytes32)
func (_IETHPOSDeposit *IETHPOSDepositSession) GetDepositRoot() ([32]byte, error) {
	return _IETHPOSDeposit.Contract.GetDepositRoot(&_IETHPOSDeposit.CallOpts)
}

// GetDepositRoot is a free data retrieval call binding the contract method 0xc5f2892f.
//
// Solidity: function get_deposit_root() view returns(bytes32)
func (_IETHPOSDeposit *IETHPOSDepositCallerSession) GetDepositRoot() ([32]byte, error) {
	return _IETHPOSDeposit.Contract.GetDepositRoot(&_IETHPOSDeposit.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x22895118.
//
// Solidity: function deposit(bytes pubkey, bytes withdrawal_credentials, bytes signature, bytes32 deposit_data_root) payable returns()
func (_IETHPOSDeposit *IETHPOSDepositTransactor) Deposit(opts *bind.TransactOpts, pubkey []byte, withdrawal_credentials []byte, signature []byte, deposit_data_root [32]byte) (*types.Transaction, error) {
	return _IETHPOSDeposit.contract.Transact(opts, "deposit", pubkey, withdrawal_credentials, signature, deposit_data_root)
}

// Deposit is a paid mutator transaction binding the contract method 0x22895118.
//
// Solidity: function deposit(bytes pubkey, bytes withdrawal_credentials, bytes signature, bytes32 deposit_data_root) payable returns()
func (_IETHPOSDeposit *IETHPOSDepositSession) Deposit(pubkey []byte, withdrawal_credentials []byte, signature []byte, deposit_data_root [32]byte) (*types.Transaction, error) {
	return _IETHPOSDeposit.Contract.Deposit(&_IETHPOSDeposit.TransactOpts, pubkey, withdrawal_credentials, signature, deposit_data_root)
}

// Deposit is a paid mutator transaction binding the contract method 0x22895118.
//
// Solidity: function deposit(bytes pubkey, bytes withdrawal_credentials, bytes signature, bytes32 deposit_data_root) payable returns()
func (_IETHPOSDeposit *IETHPOSDepositTransactorSession) Deposit(pubkey []byte, withdrawal_credentials []byte, signature []byte, deposit_data_root [32]byte) (*types.Transaction, error) {
	return _IETHPOSDeposit.Contract.Deposit(&_IETHPOSDeposit.TransactOpts, pubkey, withdrawal_credentials, signature, deposit_data_root)
}

// IETHPOSDepositDepositEventIterator is returned from FilterDepositEvent and is used to iterate over the raw logs and unpacked data for DepositEvent events raised by the IETHPOSDeposit contract.
type IETHPOSDepositDepositEventIterator struct {
	Event *IETHPOSDepositDepositEvent // Event containing the contract specifics and raw log

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
func (it *IETHPOSDepositDepositEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IETHPOSDepositDepositEvent)
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
		it.Event = new(IETHPOSDepositDepositEvent)
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
func (it *IETHPOSDepositDepositEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IETHPOSDepositDepositEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IETHPOSDepositDepositEvent represents a DepositEvent event raised by the IETHPOSDeposit contract.
type IETHPOSDepositDepositEvent struct {
	Pubkey                []byte
	WithdrawalCredentials []byte
	Amount                []byte
	Signature             []byte
	Index                 []byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterDepositEvent is a free log retrieval operation binding the contract event 0x649bbc62d0e31342afea4e5cd82d4049e7e1ee912fc0889aa790803be39038c5.
//
// Solidity: event DepositEvent(bytes pubkey, bytes withdrawal_credentials, bytes amount, bytes signature, bytes index)
func (_IETHPOSDeposit *IETHPOSDepositFilterer) FilterDepositEvent(opts *bind.FilterOpts) (*IETHPOSDepositDepositEventIterator, error) {

	logs, sub, err := _IETHPOSDeposit.contract.FilterLogs(opts, "DepositEvent")
	if err != nil {
		return nil, err
	}
	return &IETHPOSDepositDepositEventIterator{contract: _IETHPOSDeposit.contract, event: "DepositEvent", logs: logs, sub: sub}, nil
}

// WatchDepositEvent is a free log subscription operation binding the contract event 0x649bbc62d0e31342afea4e5cd82d4049e7e1ee912fc0889aa790803be39038c5.
//
// Solidity: event DepositEvent(bytes pubkey, bytes withdrawal_credentials, bytes amount, bytes signature, bytes index)
func (_IETHPOSDeposit *IETHPOSDepositFilterer) WatchDepositEvent(opts *bind.WatchOpts, sink chan<- *IETHPOSDepositDepositEvent) (event.Subscription, error) {

	logs, sub, err := _IETHPOSDeposit.contract.WatchLogs(opts, "DepositEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IETHPOSDepositDepositEvent)
				if err := _IETHPOSDeposit.contract.UnpackLog(event, "DepositEvent", log); err != nil {
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

// ParseDepositEvent is a log parse operation binding the contract event 0x649bbc62d0e31342afea4e5cd82d4049e7e1ee912fc0889aa790803be39038c5.
//
// Solidity: event DepositEvent(bytes pubkey, bytes withdrawal_credentials, bytes amount, bytes signature, bytes index)
func (_IETHPOSDeposit *IETHPOSDepositFilterer) ParseDepositEvent(log types.Log) (*IETHPOSDepositDepositEvent, error) {
	event := new(IETHPOSDepositDepositEvent)
	if err := _IETHPOSDeposit.contract.UnpackLog(event, "DepositEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
