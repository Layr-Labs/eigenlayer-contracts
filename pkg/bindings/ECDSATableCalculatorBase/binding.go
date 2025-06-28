// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ECDSATableCalculatorBase

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

// IECDSATableCalculatorTypesECDSAOperatorInfo is an auto generated low-level Go binding around an user-defined struct.
type IECDSATableCalculatorTypesECDSAOperatorInfo struct {
	Pubkey  common.Address
	Weights []*big.Int
}

// OperatorSet is an auto generated low-level Go binding around an user-defined struct.
type OperatorSet struct {
	Avs common.Address
	Id  uint32
}

// ECDSATableCalculatorBaseMetaData contains all meta data concerning the ECDSATableCalculatorBase contract.
var ECDSATableCalculatorBaseMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"calculateOperatorTable\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"operatorInfos\",\"type\":\"tuple[]\",\"internalType\":\"structIECDSATableCalculatorTypes.ECDSAOperatorInfo[]\",\"components\":[{\"name\":\"pubkey\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"weights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateOperatorTableBytes\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"operatorTableBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorWeight\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorWeights\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"weights\",\"type\":\"uint256[][]\",\"internalType\":\"uint256[][]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"keyRegistrar\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIKeyRegistrar\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"LookaheadBlocksSet\",\"inputs\":[{\"name\":\"lookaheadBlocks\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"LookaheadBlocksTooHigh\",\"inputs\":[]}]",
}

// ECDSATableCalculatorBaseABI is the input ABI used to generate the binding from.
// Deprecated: Use ECDSATableCalculatorBaseMetaData.ABI instead.
var ECDSATableCalculatorBaseABI = ECDSATableCalculatorBaseMetaData.ABI

// ECDSATableCalculatorBase is an auto generated Go binding around an Ethereum contract.
type ECDSATableCalculatorBase struct {
	ECDSATableCalculatorBaseCaller     // Read-only binding to the contract
	ECDSATableCalculatorBaseTransactor // Write-only binding to the contract
	ECDSATableCalculatorBaseFilterer   // Log filterer for contract events
}

// ECDSATableCalculatorBaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type ECDSATableCalculatorBaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSATableCalculatorBaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ECDSATableCalculatorBaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSATableCalculatorBaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ECDSATableCalculatorBaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSATableCalculatorBaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ECDSATableCalculatorBaseSession struct {
	Contract     *ECDSATableCalculatorBase // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ECDSATableCalculatorBaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ECDSATableCalculatorBaseCallerSession struct {
	Contract *ECDSATableCalculatorBaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// ECDSATableCalculatorBaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ECDSATableCalculatorBaseTransactorSession struct {
	Contract     *ECDSATableCalculatorBaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// ECDSATableCalculatorBaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type ECDSATableCalculatorBaseRaw struct {
	Contract *ECDSATableCalculatorBase // Generic contract binding to access the raw methods on
}

// ECDSATableCalculatorBaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ECDSATableCalculatorBaseCallerRaw struct {
	Contract *ECDSATableCalculatorBaseCaller // Generic read-only contract binding to access the raw methods on
}

// ECDSATableCalculatorBaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ECDSATableCalculatorBaseTransactorRaw struct {
	Contract *ECDSATableCalculatorBaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewECDSATableCalculatorBase creates a new instance of ECDSATableCalculatorBase, bound to a specific deployed contract.
func NewECDSATableCalculatorBase(address common.Address, backend bind.ContractBackend) (*ECDSATableCalculatorBase, error) {
	contract, err := bindECDSATableCalculatorBase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ECDSATableCalculatorBase{ECDSATableCalculatorBaseCaller: ECDSATableCalculatorBaseCaller{contract: contract}, ECDSATableCalculatorBaseTransactor: ECDSATableCalculatorBaseTransactor{contract: contract}, ECDSATableCalculatorBaseFilterer: ECDSATableCalculatorBaseFilterer{contract: contract}}, nil
}

// NewECDSATableCalculatorBaseCaller creates a new read-only instance of ECDSATableCalculatorBase, bound to a specific deployed contract.
func NewECDSATableCalculatorBaseCaller(address common.Address, caller bind.ContractCaller) (*ECDSATableCalculatorBaseCaller, error) {
	contract, err := bindECDSATableCalculatorBase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ECDSATableCalculatorBaseCaller{contract: contract}, nil
}

// NewECDSATableCalculatorBaseTransactor creates a new write-only instance of ECDSATableCalculatorBase, bound to a specific deployed contract.
func NewECDSATableCalculatorBaseTransactor(address common.Address, transactor bind.ContractTransactor) (*ECDSATableCalculatorBaseTransactor, error) {
	contract, err := bindECDSATableCalculatorBase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ECDSATableCalculatorBaseTransactor{contract: contract}, nil
}

// NewECDSATableCalculatorBaseFilterer creates a new log filterer instance of ECDSATableCalculatorBase, bound to a specific deployed contract.
func NewECDSATableCalculatorBaseFilterer(address common.Address, filterer bind.ContractFilterer) (*ECDSATableCalculatorBaseFilterer, error) {
	contract, err := bindECDSATableCalculatorBase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ECDSATableCalculatorBaseFilterer{contract: contract}, nil
}

// bindECDSATableCalculatorBase binds a generic wrapper to an already deployed contract.
func bindECDSATableCalculatorBase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ECDSATableCalculatorBaseMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECDSATableCalculatorBase *ECDSATableCalculatorBaseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ECDSATableCalculatorBase.Contract.ECDSATableCalculatorBaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECDSATableCalculatorBase *ECDSATableCalculatorBaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECDSATableCalculatorBase.Contract.ECDSATableCalculatorBaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECDSATableCalculatorBase *ECDSATableCalculatorBaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECDSATableCalculatorBase.Contract.ECDSATableCalculatorBaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECDSATableCalculatorBase *ECDSATableCalculatorBaseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ECDSATableCalculatorBase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECDSATableCalculatorBase *ECDSATableCalculatorBaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECDSATableCalculatorBase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECDSATableCalculatorBase *ECDSATableCalculatorBaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECDSATableCalculatorBase.Contract.contract.Transact(opts, method, params...)
}

// CalculateOperatorTable is a free data retrieval call binding the contract method 0x124c87e0.
//
// Solidity: function calculateOperatorTable((address,uint32) operatorSet) view returns((address,uint256[])[] operatorInfos)
func (_ECDSATableCalculatorBase *ECDSATableCalculatorBaseCaller) CalculateOperatorTable(opts *bind.CallOpts, operatorSet OperatorSet) ([]IECDSATableCalculatorTypesECDSAOperatorInfo, error) {
	var out []interface{}
	err := _ECDSATableCalculatorBase.contract.Call(opts, &out, "calculateOperatorTable", operatorSet)

	if err != nil {
		return *new([]IECDSATableCalculatorTypesECDSAOperatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IECDSATableCalculatorTypesECDSAOperatorInfo)).(*[]IECDSATableCalculatorTypesECDSAOperatorInfo)

	return out0, err

}

// CalculateOperatorTable is a free data retrieval call binding the contract method 0x124c87e0.
//
// Solidity: function calculateOperatorTable((address,uint32) operatorSet) view returns((address,uint256[])[] operatorInfos)
func (_ECDSATableCalculatorBase *ECDSATableCalculatorBaseSession) CalculateOperatorTable(operatorSet OperatorSet) ([]IECDSATableCalculatorTypesECDSAOperatorInfo, error) {
	return _ECDSATableCalculatorBase.Contract.CalculateOperatorTable(&_ECDSATableCalculatorBase.CallOpts, operatorSet)
}

// CalculateOperatorTable is a free data retrieval call binding the contract method 0x124c87e0.
//
// Solidity: function calculateOperatorTable((address,uint32) operatorSet) view returns((address,uint256[])[] operatorInfos)
func (_ECDSATableCalculatorBase *ECDSATableCalculatorBaseCallerSession) CalculateOperatorTable(operatorSet OperatorSet) ([]IECDSATableCalculatorTypesECDSAOperatorInfo, error) {
	return _ECDSATableCalculatorBase.Contract.CalculateOperatorTable(&_ECDSATableCalculatorBase.CallOpts, operatorSet)
}

// CalculateOperatorTableBytes is a free data retrieval call binding the contract method 0x41ee6d0e.
//
// Solidity: function calculateOperatorTableBytes((address,uint32) operatorSet) view returns(bytes operatorTableBytes)
func (_ECDSATableCalculatorBase *ECDSATableCalculatorBaseCaller) CalculateOperatorTableBytes(opts *bind.CallOpts, operatorSet OperatorSet) ([]byte, error) {
	var out []interface{}
	err := _ECDSATableCalculatorBase.contract.Call(opts, &out, "calculateOperatorTableBytes", operatorSet)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// CalculateOperatorTableBytes is a free data retrieval call binding the contract method 0x41ee6d0e.
//
// Solidity: function calculateOperatorTableBytes((address,uint32) operatorSet) view returns(bytes operatorTableBytes)
func (_ECDSATableCalculatorBase *ECDSATableCalculatorBaseSession) CalculateOperatorTableBytes(operatorSet OperatorSet) ([]byte, error) {
	return _ECDSATableCalculatorBase.Contract.CalculateOperatorTableBytes(&_ECDSATableCalculatorBase.CallOpts, operatorSet)
}

// CalculateOperatorTableBytes is a free data retrieval call binding the contract method 0x41ee6d0e.
//
// Solidity: function calculateOperatorTableBytes((address,uint32) operatorSet) view returns(bytes operatorTableBytes)
func (_ECDSATableCalculatorBase *ECDSATableCalculatorBaseCallerSession) CalculateOperatorTableBytes(operatorSet OperatorSet) ([]byte, error) {
	return _ECDSATableCalculatorBase.Contract.CalculateOperatorTableBytes(&_ECDSATableCalculatorBase.CallOpts, operatorSet)
}

// GetOperatorWeight is a free data retrieval call binding the contract method 0x1088794a.
//
// Solidity: function getOperatorWeight((address,uint32) operatorSet, address operator) view returns(uint256 weight)
func (_ECDSATableCalculatorBase *ECDSATableCalculatorBaseCaller) GetOperatorWeight(opts *bind.CallOpts, operatorSet OperatorSet, operator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ECDSATableCalculatorBase.contract.Call(opts, &out, "getOperatorWeight", operatorSet, operator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOperatorWeight is a free data retrieval call binding the contract method 0x1088794a.
//
// Solidity: function getOperatorWeight((address,uint32) operatorSet, address operator) view returns(uint256 weight)
func (_ECDSATableCalculatorBase *ECDSATableCalculatorBaseSession) GetOperatorWeight(operatorSet OperatorSet, operator common.Address) (*big.Int, error) {
	return _ECDSATableCalculatorBase.Contract.GetOperatorWeight(&_ECDSATableCalculatorBase.CallOpts, operatorSet, operator)
}

// GetOperatorWeight is a free data retrieval call binding the contract method 0x1088794a.
//
// Solidity: function getOperatorWeight((address,uint32) operatorSet, address operator) view returns(uint256 weight)
func (_ECDSATableCalculatorBase *ECDSATableCalculatorBaseCallerSession) GetOperatorWeight(operatorSet OperatorSet, operator common.Address) (*big.Int, error) {
	return _ECDSATableCalculatorBase.Contract.GetOperatorWeight(&_ECDSATableCalculatorBase.CallOpts, operatorSet, operator)
}

// GetOperatorWeights is a free data retrieval call binding the contract method 0x71ca71d9.
//
// Solidity: function getOperatorWeights((address,uint32) operatorSet) view returns(address[] operators, uint256[][] weights)
func (_ECDSATableCalculatorBase *ECDSATableCalculatorBaseCaller) GetOperatorWeights(opts *bind.CallOpts, operatorSet OperatorSet) (struct {
	Operators []common.Address
	Weights   [][]*big.Int
}, error) {
	var out []interface{}
	err := _ECDSATableCalculatorBase.contract.Call(opts, &out, "getOperatorWeights", operatorSet)

	outstruct := new(struct {
		Operators []common.Address
		Weights   [][]*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Operators = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Weights = *abi.ConvertType(out[1], new([][]*big.Int)).(*[][]*big.Int)

	return *outstruct, err

}

// GetOperatorWeights is a free data retrieval call binding the contract method 0x71ca71d9.
//
// Solidity: function getOperatorWeights((address,uint32) operatorSet) view returns(address[] operators, uint256[][] weights)
func (_ECDSATableCalculatorBase *ECDSATableCalculatorBaseSession) GetOperatorWeights(operatorSet OperatorSet) (struct {
	Operators []common.Address
	Weights   [][]*big.Int
}, error) {
	return _ECDSATableCalculatorBase.Contract.GetOperatorWeights(&_ECDSATableCalculatorBase.CallOpts, operatorSet)
}

// GetOperatorWeights is a free data retrieval call binding the contract method 0x71ca71d9.
//
// Solidity: function getOperatorWeights((address,uint32) operatorSet) view returns(address[] operators, uint256[][] weights)
func (_ECDSATableCalculatorBase *ECDSATableCalculatorBaseCallerSession) GetOperatorWeights(operatorSet OperatorSet) (struct {
	Operators []common.Address
	Weights   [][]*big.Int
}, error) {
	return _ECDSATableCalculatorBase.Contract.GetOperatorWeights(&_ECDSATableCalculatorBase.CallOpts, operatorSet)
}

// KeyRegistrar is a free data retrieval call binding the contract method 0x3ec45c7e.
//
// Solidity: function keyRegistrar() view returns(address)
func (_ECDSATableCalculatorBase *ECDSATableCalculatorBaseCaller) KeyRegistrar(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ECDSATableCalculatorBase.contract.Call(opts, &out, "keyRegistrar")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// KeyRegistrar is a free data retrieval call binding the contract method 0x3ec45c7e.
//
// Solidity: function keyRegistrar() view returns(address)
func (_ECDSATableCalculatorBase *ECDSATableCalculatorBaseSession) KeyRegistrar() (common.Address, error) {
	return _ECDSATableCalculatorBase.Contract.KeyRegistrar(&_ECDSATableCalculatorBase.CallOpts)
}

// KeyRegistrar is a free data retrieval call binding the contract method 0x3ec45c7e.
//
// Solidity: function keyRegistrar() view returns(address)
func (_ECDSATableCalculatorBase *ECDSATableCalculatorBaseCallerSession) KeyRegistrar() (common.Address, error) {
	return _ECDSATableCalculatorBase.Contract.KeyRegistrar(&_ECDSATableCalculatorBase.CallOpts)
}

// ECDSATableCalculatorBaseLookaheadBlocksSetIterator is returned from FilterLookaheadBlocksSet and is used to iterate over the raw logs and unpacked data for LookaheadBlocksSet events raised by the ECDSATableCalculatorBase contract.
type ECDSATableCalculatorBaseLookaheadBlocksSetIterator struct {
	Event *ECDSATableCalculatorBaseLookaheadBlocksSet // Event containing the contract specifics and raw log

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
func (it *ECDSATableCalculatorBaseLookaheadBlocksSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ECDSATableCalculatorBaseLookaheadBlocksSet)
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
		it.Event = new(ECDSATableCalculatorBaseLookaheadBlocksSet)
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
func (it *ECDSATableCalculatorBaseLookaheadBlocksSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ECDSATableCalculatorBaseLookaheadBlocksSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ECDSATableCalculatorBaseLookaheadBlocksSet represents a LookaheadBlocksSet event raised by the ECDSATableCalculatorBase contract.
type ECDSATableCalculatorBaseLookaheadBlocksSet struct {
	LookaheadBlocks *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLookaheadBlocksSet is a free log retrieval operation binding the contract event 0xa41e64dd47db91b61b43ccea8b57d75abfa496f23efc708c22753c4bc9d68842.
//
// Solidity: event LookaheadBlocksSet(uint256 lookaheadBlocks)
func (_ECDSATableCalculatorBase *ECDSATableCalculatorBaseFilterer) FilterLookaheadBlocksSet(opts *bind.FilterOpts) (*ECDSATableCalculatorBaseLookaheadBlocksSetIterator, error) {

	logs, sub, err := _ECDSATableCalculatorBase.contract.FilterLogs(opts, "LookaheadBlocksSet")
	if err != nil {
		return nil, err
	}
	return &ECDSATableCalculatorBaseLookaheadBlocksSetIterator{contract: _ECDSATableCalculatorBase.contract, event: "LookaheadBlocksSet", logs: logs, sub: sub}, nil
}

// WatchLookaheadBlocksSet is a free log subscription operation binding the contract event 0xa41e64dd47db91b61b43ccea8b57d75abfa496f23efc708c22753c4bc9d68842.
//
// Solidity: event LookaheadBlocksSet(uint256 lookaheadBlocks)
func (_ECDSATableCalculatorBase *ECDSATableCalculatorBaseFilterer) WatchLookaheadBlocksSet(opts *bind.WatchOpts, sink chan<- *ECDSATableCalculatorBaseLookaheadBlocksSet) (event.Subscription, error) {

	logs, sub, err := _ECDSATableCalculatorBase.contract.WatchLogs(opts, "LookaheadBlocksSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ECDSATableCalculatorBaseLookaheadBlocksSet)
				if err := _ECDSATableCalculatorBase.contract.UnpackLog(event, "LookaheadBlocksSet", log); err != nil {
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

// ParseLookaheadBlocksSet is a log parse operation binding the contract event 0xa41e64dd47db91b61b43ccea8b57d75abfa496f23efc708c22753c4bc9d68842.
//
// Solidity: event LookaheadBlocksSet(uint256 lookaheadBlocks)
func (_ECDSATableCalculatorBase *ECDSATableCalculatorBaseFilterer) ParseLookaheadBlocksSet(log types.Log) (*ECDSATableCalculatorBaseLookaheadBlocksSet, error) {
	event := new(ECDSATableCalculatorBaseLookaheadBlocksSet)
	if err := _ECDSATableCalculatorBase.contract.UnpackLog(event, "LookaheadBlocksSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
