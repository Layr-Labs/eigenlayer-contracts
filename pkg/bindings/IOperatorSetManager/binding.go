// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IOperatorSetManager

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

// IOperatorSetManagerOperatorSet is an auto generated low-level Go binding around an user-defined struct.
type IOperatorSetManagerOperatorSet struct {
	Avs common.Address
	Id  [4]byte
}

// IOperatorSetManagerSlashingMagnitudeParameters is an auto generated low-level Go binding around an user-defined struct.
type IOperatorSetManagerSlashingMagnitudeParameters struct {
	Strategy            common.Address
	TotalMagnitude      uint64
	OperatorSets        []IOperatorSetManagerOperatorSet
	SlashableMagnitudes []uint64
}

// ISignatureUtilsSignatureWithExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithExpiry struct {
	Signature []byte
	Expiry    *big.Int
}

// IOperatorSetManagerMetaData contains all meta data concerning the IOperatorSetManager contract.
var IOperatorSetManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"getSlashableBips\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structIOperatorSetManager.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"epoch\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"slashableBips\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateSlashingParameters\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"slashingMagnitudeParameters\",\"type\":\"tuple[]\",\"internalType\":\"structIOperatorSetManager.SlashingMagnitudeParameters[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"totalMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"operatorSets\",\"type\":\"tuple[]\",\"internalType\":\"structIOperatorSetManager.OperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"name\":\"slashableMagnitudes\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}]},{\"name\":\"allocatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"effectEpoch\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"SlashableMagnitudeUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIOperatorSetManager.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"name\":\"slashableMagnitude\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"effectEpoch\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TotalMagnitudeUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"totalMagnitude\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"effectEpoch\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false}]",
}

// IOperatorSetManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use IOperatorSetManagerMetaData.ABI instead.
var IOperatorSetManagerABI = IOperatorSetManagerMetaData.ABI

// IOperatorSetManager is an auto generated Go binding around an Ethereum contract.
type IOperatorSetManager struct {
	IOperatorSetManagerCaller     // Read-only binding to the contract
	IOperatorSetManagerTransactor // Write-only binding to the contract
	IOperatorSetManagerFilterer   // Log filterer for contract events
}

// IOperatorSetManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IOperatorSetManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOperatorSetManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IOperatorSetManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOperatorSetManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IOperatorSetManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOperatorSetManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IOperatorSetManagerSession struct {
	Contract     *IOperatorSetManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IOperatorSetManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IOperatorSetManagerCallerSession struct {
	Contract *IOperatorSetManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// IOperatorSetManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IOperatorSetManagerTransactorSession struct {
	Contract     *IOperatorSetManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IOperatorSetManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IOperatorSetManagerRaw struct {
	Contract *IOperatorSetManager // Generic contract binding to access the raw methods on
}

// IOperatorSetManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IOperatorSetManagerCallerRaw struct {
	Contract *IOperatorSetManagerCaller // Generic read-only contract binding to access the raw methods on
}

// IOperatorSetManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IOperatorSetManagerTransactorRaw struct {
	Contract *IOperatorSetManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIOperatorSetManager creates a new instance of IOperatorSetManager, bound to a specific deployed contract.
func NewIOperatorSetManager(address common.Address, backend bind.ContractBackend) (*IOperatorSetManager, error) {
	contract, err := bindIOperatorSetManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IOperatorSetManager{IOperatorSetManagerCaller: IOperatorSetManagerCaller{contract: contract}, IOperatorSetManagerTransactor: IOperatorSetManagerTransactor{contract: contract}, IOperatorSetManagerFilterer: IOperatorSetManagerFilterer{contract: contract}}, nil
}

// NewIOperatorSetManagerCaller creates a new read-only instance of IOperatorSetManager, bound to a specific deployed contract.
func NewIOperatorSetManagerCaller(address common.Address, caller bind.ContractCaller) (*IOperatorSetManagerCaller, error) {
	contract, err := bindIOperatorSetManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IOperatorSetManagerCaller{contract: contract}, nil
}

// NewIOperatorSetManagerTransactor creates a new write-only instance of IOperatorSetManager, bound to a specific deployed contract.
func NewIOperatorSetManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IOperatorSetManagerTransactor, error) {
	contract, err := bindIOperatorSetManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IOperatorSetManagerTransactor{contract: contract}, nil
}

// NewIOperatorSetManagerFilterer creates a new log filterer instance of IOperatorSetManager, bound to a specific deployed contract.
func NewIOperatorSetManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IOperatorSetManagerFilterer, error) {
	contract, err := bindIOperatorSetManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IOperatorSetManagerFilterer{contract: contract}, nil
}

// bindIOperatorSetManager binds a generic wrapper to an already deployed contract.
func bindIOperatorSetManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IOperatorSetManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOperatorSetManager *IOperatorSetManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOperatorSetManager.Contract.IOperatorSetManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOperatorSetManager *IOperatorSetManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOperatorSetManager.Contract.IOperatorSetManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOperatorSetManager *IOperatorSetManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOperatorSetManager.Contract.IOperatorSetManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOperatorSetManager *IOperatorSetManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOperatorSetManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOperatorSetManager *IOperatorSetManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOperatorSetManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOperatorSetManager *IOperatorSetManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOperatorSetManager.Contract.contract.Transact(opts, method, params...)
}

// GetSlashableBips is a free data retrieval call binding the contract method 0x3f76c6c7.
//
// Solidity: function getSlashableBips(address operator, (address,bytes4) operatorSet, address strategy, uint32 epoch) view returns(uint16 slashableBips)
func (_IOperatorSetManager *IOperatorSetManagerCaller) GetSlashableBips(opts *bind.CallOpts, operator common.Address, operatorSet IOperatorSetManagerOperatorSet, strategy common.Address, epoch uint32) (uint16, error) {
	var out []interface{}
	err := _IOperatorSetManager.contract.Call(opts, &out, "getSlashableBips", operator, operatorSet, strategy, epoch)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetSlashableBips is a free data retrieval call binding the contract method 0x3f76c6c7.
//
// Solidity: function getSlashableBips(address operator, (address,bytes4) operatorSet, address strategy, uint32 epoch) view returns(uint16 slashableBips)
func (_IOperatorSetManager *IOperatorSetManagerSession) GetSlashableBips(operator common.Address, operatorSet IOperatorSetManagerOperatorSet, strategy common.Address, epoch uint32) (uint16, error) {
	return _IOperatorSetManager.Contract.GetSlashableBips(&_IOperatorSetManager.CallOpts, operator, operatorSet, strategy, epoch)
}

// GetSlashableBips is a free data retrieval call binding the contract method 0x3f76c6c7.
//
// Solidity: function getSlashableBips(address operator, (address,bytes4) operatorSet, address strategy, uint32 epoch) view returns(uint16 slashableBips)
func (_IOperatorSetManager *IOperatorSetManagerCallerSession) GetSlashableBips(operator common.Address, operatorSet IOperatorSetManagerOperatorSet, strategy common.Address, epoch uint32) (uint16, error) {
	return _IOperatorSetManager.Contract.GetSlashableBips(&_IOperatorSetManager.CallOpts, operator, operatorSet, strategy, epoch)
}

// UpdateSlashingParameters is a paid mutator transaction binding the contract method 0x516b9883.
//
// Solidity: function updateSlashingParameters(address operator, (address,uint64,(address,bytes4)[],uint64[])[] slashingMagnitudeParameters, (bytes,uint256) allocatorSignature) returns(uint32 effectEpoch)
func (_IOperatorSetManager *IOperatorSetManagerTransactor) UpdateSlashingParameters(opts *bind.TransactOpts, operator common.Address, slashingMagnitudeParameters []IOperatorSetManagerSlashingMagnitudeParameters, allocatorSignature ISignatureUtilsSignatureWithExpiry) (*types.Transaction, error) {
	return _IOperatorSetManager.contract.Transact(opts, "updateSlashingParameters", operator, slashingMagnitudeParameters, allocatorSignature)
}

// UpdateSlashingParameters is a paid mutator transaction binding the contract method 0x516b9883.
//
// Solidity: function updateSlashingParameters(address operator, (address,uint64,(address,bytes4)[],uint64[])[] slashingMagnitudeParameters, (bytes,uint256) allocatorSignature) returns(uint32 effectEpoch)
func (_IOperatorSetManager *IOperatorSetManagerSession) UpdateSlashingParameters(operator common.Address, slashingMagnitudeParameters []IOperatorSetManagerSlashingMagnitudeParameters, allocatorSignature ISignatureUtilsSignatureWithExpiry) (*types.Transaction, error) {
	return _IOperatorSetManager.Contract.UpdateSlashingParameters(&_IOperatorSetManager.TransactOpts, operator, slashingMagnitudeParameters, allocatorSignature)
}

// UpdateSlashingParameters is a paid mutator transaction binding the contract method 0x516b9883.
//
// Solidity: function updateSlashingParameters(address operator, (address,uint64,(address,bytes4)[],uint64[])[] slashingMagnitudeParameters, (bytes,uint256) allocatorSignature) returns(uint32 effectEpoch)
func (_IOperatorSetManager *IOperatorSetManagerTransactorSession) UpdateSlashingParameters(operator common.Address, slashingMagnitudeParameters []IOperatorSetManagerSlashingMagnitudeParameters, allocatorSignature ISignatureUtilsSignatureWithExpiry) (*types.Transaction, error) {
	return _IOperatorSetManager.Contract.UpdateSlashingParameters(&_IOperatorSetManager.TransactOpts, operator, slashingMagnitudeParameters, allocatorSignature)
}

// IOperatorSetManagerSlashableMagnitudeUpdatedIterator is returned from FilterSlashableMagnitudeUpdated and is used to iterate over the raw logs and unpacked data for SlashableMagnitudeUpdated events raised by the IOperatorSetManager contract.
type IOperatorSetManagerSlashableMagnitudeUpdatedIterator struct {
	Event *IOperatorSetManagerSlashableMagnitudeUpdated // Event containing the contract specifics and raw log

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
func (it *IOperatorSetManagerSlashableMagnitudeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOperatorSetManagerSlashableMagnitudeUpdated)
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
		it.Event = new(IOperatorSetManagerSlashableMagnitudeUpdated)
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
func (it *IOperatorSetManagerSlashableMagnitudeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOperatorSetManagerSlashableMagnitudeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOperatorSetManagerSlashableMagnitudeUpdated represents a SlashableMagnitudeUpdated event raised by the IOperatorSetManager contract.
type IOperatorSetManagerSlashableMagnitudeUpdated struct {
	Operator           common.Address
	Strategy           common.Address
	OperatorSet        IOperatorSetManagerOperatorSet
	SlashableMagnitude uint64
	EffectEpoch        uint32
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterSlashableMagnitudeUpdated is a free log retrieval operation binding the contract event 0xa051327ef1123f482ec636fa78d997e135019b2bcfa0fab32a90462654299506.
//
// Solidity: event SlashableMagnitudeUpdated(address operator, address strategy, (address,bytes4) operatorSet, uint64 slashableMagnitude, uint32 effectEpoch)
func (_IOperatorSetManager *IOperatorSetManagerFilterer) FilterSlashableMagnitudeUpdated(opts *bind.FilterOpts) (*IOperatorSetManagerSlashableMagnitudeUpdatedIterator, error) {

	logs, sub, err := _IOperatorSetManager.contract.FilterLogs(opts, "SlashableMagnitudeUpdated")
	if err != nil {
		return nil, err
	}
	return &IOperatorSetManagerSlashableMagnitudeUpdatedIterator{contract: _IOperatorSetManager.contract, event: "SlashableMagnitudeUpdated", logs: logs, sub: sub}, nil
}

// WatchSlashableMagnitudeUpdated is a free log subscription operation binding the contract event 0xa051327ef1123f482ec636fa78d997e135019b2bcfa0fab32a90462654299506.
//
// Solidity: event SlashableMagnitudeUpdated(address operator, address strategy, (address,bytes4) operatorSet, uint64 slashableMagnitude, uint32 effectEpoch)
func (_IOperatorSetManager *IOperatorSetManagerFilterer) WatchSlashableMagnitudeUpdated(opts *bind.WatchOpts, sink chan<- *IOperatorSetManagerSlashableMagnitudeUpdated) (event.Subscription, error) {

	logs, sub, err := _IOperatorSetManager.contract.WatchLogs(opts, "SlashableMagnitudeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOperatorSetManagerSlashableMagnitudeUpdated)
				if err := _IOperatorSetManager.contract.UnpackLog(event, "SlashableMagnitudeUpdated", log); err != nil {
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

// ParseSlashableMagnitudeUpdated is a log parse operation binding the contract event 0xa051327ef1123f482ec636fa78d997e135019b2bcfa0fab32a90462654299506.
//
// Solidity: event SlashableMagnitudeUpdated(address operator, address strategy, (address,bytes4) operatorSet, uint64 slashableMagnitude, uint32 effectEpoch)
func (_IOperatorSetManager *IOperatorSetManagerFilterer) ParseSlashableMagnitudeUpdated(log types.Log) (*IOperatorSetManagerSlashableMagnitudeUpdated, error) {
	event := new(IOperatorSetManagerSlashableMagnitudeUpdated)
	if err := _IOperatorSetManager.contract.UnpackLog(event, "SlashableMagnitudeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOperatorSetManagerTotalMagnitudeUpdatedIterator is returned from FilterTotalMagnitudeUpdated and is used to iterate over the raw logs and unpacked data for TotalMagnitudeUpdated events raised by the IOperatorSetManager contract.
type IOperatorSetManagerTotalMagnitudeUpdatedIterator struct {
	Event *IOperatorSetManagerTotalMagnitudeUpdated // Event containing the contract specifics and raw log

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
func (it *IOperatorSetManagerTotalMagnitudeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOperatorSetManagerTotalMagnitudeUpdated)
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
		it.Event = new(IOperatorSetManagerTotalMagnitudeUpdated)
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
func (it *IOperatorSetManagerTotalMagnitudeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOperatorSetManagerTotalMagnitudeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOperatorSetManagerTotalMagnitudeUpdated represents a TotalMagnitudeUpdated event raised by the IOperatorSetManager contract.
type IOperatorSetManagerTotalMagnitudeUpdated struct {
	Operator       common.Address
	Strategy       common.Address
	TotalMagnitude uint64
	EffectEpoch    uint32
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTotalMagnitudeUpdated is a free log retrieval operation binding the contract event 0x802e045376358152b85ba2107735ff6f465df424b0fcbcb4690c83951d73ebd6.
//
// Solidity: event TotalMagnitudeUpdated(address operator, address strategy, uint64 totalMagnitude, uint32 effectEpoch)
func (_IOperatorSetManager *IOperatorSetManagerFilterer) FilterTotalMagnitudeUpdated(opts *bind.FilterOpts) (*IOperatorSetManagerTotalMagnitudeUpdatedIterator, error) {

	logs, sub, err := _IOperatorSetManager.contract.FilterLogs(opts, "TotalMagnitudeUpdated")
	if err != nil {
		return nil, err
	}
	return &IOperatorSetManagerTotalMagnitudeUpdatedIterator{contract: _IOperatorSetManager.contract, event: "TotalMagnitudeUpdated", logs: logs, sub: sub}, nil
}

// WatchTotalMagnitudeUpdated is a free log subscription operation binding the contract event 0x802e045376358152b85ba2107735ff6f465df424b0fcbcb4690c83951d73ebd6.
//
// Solidity: event TotalMagnitudeUpdated(address operator, address strategy, uint64 totalMagnitude, uint32 effectEpoch)
func (_IOperatorSetManager *IOperatorSetManagerFilterer) WatchTotalMagnitudeUpdated(opts *bind.WatchOpts, sink chan<- *IOperatorSetManagerTotalMagnitudeUpdated) (event.Subscription, error) {

	logs, sub, err := _IOperatorSetManager.contract.WatchLogs(opts, "TotalMagnitudeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOperatorSetManagerTotalMagnitudeUpdated)
				if err := _IOperatorSetManager.contract.UnpackLog(event, "TotalMagnitudeUpdated", log); err != nil {
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

// ParseTotalMagnitudeUpdated is a log parse operation binding the contract event 0x802e045376358152b85ba2107735ff6f465df424b0fcbcb4690c83951d73ebd6.
//
// Solidity: event TotalMagnitudeUpdated(address operator, address strategy, uint64 totalMagnitude, uint32 effectEpoch)
func (_IOperatorSetManager *IOperatorSetManagerFilterer) ParseTotalMagnitudeUpdated(log types.Log) (*IOperatorSetManagerTotalMagnitudeUpdated, error) {
	event := new(IOperatorSetManagerTotalMagnitudeUpdated)
	if err := _IOperatorSetManager.contract.UnpackLog(event, "TotalMagnitudeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
