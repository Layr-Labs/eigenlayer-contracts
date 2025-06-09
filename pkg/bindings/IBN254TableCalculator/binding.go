// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IBN254TableCalculator

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

// BN254G1Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G1Point struct {
	X *big.Int
	Y *big.Int
}

// IBN254TableCalculatorTypesBN254OperatorInfo is an auto generated low-level Go binding around an user-defined struct.
type IBN254TableCalculatorTypesBN254OperatorInfo struct {
	Pubkey  BN254G1Point
	Weights []*big.Int
}

// IBN254TableCalculatorTypesBN254OperatorSetInfo is an auto generated low-level Go binding around an user-defined struct.
type IBN254TableCalculatorTypesBN254OperatorSetInfo struct {
	OperatorInfoTreeRoot [32]byte
	NumOperators         *big.Int
	AggregatePubkey      BN254G1Point
	TotalWeights         []*big.Int
}

// OperatorSet is an auto generated low-level Go binding around an user-defined struct.
type OperatorSet struct {
	Avs common.Address
	Id  uint32
}

// IBN254TableCalculatorMetaData contains all meta data concerning the IBN254TableCalculator contract.
var IBN254TableCalculatorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"calculateOperatorTable\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"operatorSetInfo\",\"type\":\"tuple\",\"internalType\":\"structIBN254TableCalculatorTypes.BN254OperatorSetInfo\",\"components\":[{\"name\":\"operatorInfoTreeRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"numOperators\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"aggregatePubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"totalWeights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateOperatorTableBytes\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"operatorTableBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorInfos\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"operatorInfos\",\"type\":\"tuple[]\",\"internalType\":\"structIBN254TableCalculatorTypes.BN254OperatorInfo[]\",\"components\":[{\"name\":\"pubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorWeight\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorWeights\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"weights\",\"type\":\"uint256[][]\",\"internalType\":\"uint256[][]\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"LookaheadBlocksSet\",\"inputs\":[{\"name\":\"lookaheadBlocks\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"LookaheadBlocksTooHigh\",\"inputs\":[]}]",
}

// IBN254TableCalculatorABI is the input ABI used to generate the binding from.
// Deprecated: Use IBN254TableCalculatorMetaData.ABI instead.
var IBN254TableCalculatorABI = IBN254TableCalculatorMetaData.ABI

// IBN254TableCalculator is an auto generated Go binding around an Ethereum contract.
type IBN254TableCalculator struct {
	IBN254TableCalculatorCaller     // Read-only binding to the contract
	IBN254TableCalculatorTransactor // Write-only binding to the contract
	IBN254TableCalculatorFilterer   // Log filterer for contract events
}

// IBN254TableCalculatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBN254TableCalculatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBN254TableCalculatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBN254TableCalculatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBN254TableCalculatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBN254TableCalculatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBN254TableCalculatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBN254TableCalculatorSession struct {
	Contract     *IBN254TableCalculator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IBN254TableCalculatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBN254TableCalculatorCallerSession struct {
	Contract *IBN254TableCalculatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// IBN254TableCalculatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBN254TableCalculatorTransactorSession struct {
	Contract     *IBN254TableCalculatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// IBN254TableCalculatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBN254TableCalculatorRaw struct {
	Contract *IBN254TableCalculator // Generic contract binding to access the raw methods on
}

// IBN254TableCalculatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBN254TableCalculatorCallerRaw struct {
	Contract *IBN254TableCalculatorCaller // Generic read-only contract binding to access the raw methods on
}

// IBN254TableCalculatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBN254TableCalculatorTransactorRaw struct {
	Contract *IBN254TableCalculatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBN254TableCalculator creates a new instance of IBN254TableCalculator, bound to a specific deployed contract.
func NewIBN254TableCalculator(address common.Address, backend bind.ContractBackend) (*IBN254TableCalculator, error) {
	contract, err := bindIBN254TableCalculator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBN254TableCalculator{IBN254TableCalculatorCaller: IBN254TableCalculatorCaller{contract: contract}, IBN254TableCalculatorTransactor: IBN254TableCalculatorTransactor{contract: contract}, IBN254TableCalculatorFilterer: IBN254TableCalculatorFilterer{contract: contract}}, nil
}

// NewIBN254TableCalculatorCaller creates a new read-only instance of IBN254TableCalculator, bound to a specific deployed contract.
func NewIBN254TableCalculatorCaller(address common.Address, caller bind.ContractCaller) (*IBN254TableCalculatorCaller, error) {
	contract, err := bindIBN254TableCalculator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBN254TableCalculatorCaller{contract: contract}, nil
}

// NewIBN254TableCalculatorTransactor creates a new write-only instance of IBN254TableCalculator, bound to a specific deployed contract.
func NewIBN254TableCalculatorTransactor(address common.Address, transactor bind.ContractTransactor) (*IBN254TableCalculatorTransactor, error) {
	contract, err := bindIBN254TableCalculator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBN254TableCalculatorTransactor{contract: contract}, nil
}

// NewIBN254TableCalculatorFilterer creates a new log filterer instance of IBN254TableCalculator, bound to a specific deployed contract.
func NewIBN254TableCalculatorFilterer(address common.Address, filterer bind.ContractFilterer) (*IBN254TableCalculatorFilterer, error) {
	contract, err := bindIBN254TableCalculator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBN254TableCalculatorFilterer{contract: contract}, nil
}

// bindIBN254TableCalculator binds a generic wrapper to an already deployed contract.
func bindIBN254TableCalculator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IBN254TableCalculatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBN254TableCalculator *IBN254TableCalculatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBN254TableCalculator.Contract.IBN254TableCalculatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBN254TableCalculator *IBN254TableCalculatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBN254TableCalculator.Contract.IBN254TableCalculatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBN254TableCalculator *IBN254TableCalculatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBN254TableCalculator.Contract.IBN254TableCalculatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBN254TableCalculator *IBN254TableCalculatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBN254TableCalculator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBN254TableCalculator *IBN254TableCalculatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBN254TableCalculator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBN254TableCalculator *IBN254TableCalculatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBN254TableCalculator.Contract.contract.Transact(opts, method, params...)
}

// CalculateOperatorTable is a free data retrieval call binding the contract method 0x124c87e0.
//
// Solidity: function calculateOperatorTable((address,uint32) operatorSet) view returns((bytes32,uint256,(uint256,uint256),uint256[]) operatorSetInfo)
func (_IBN254TableCalculator *IBN254TableCalculatorCaller) CalculateOperatorTable(opts *bind.CallOpts, operatorSet OperatorSet) (IBN254TableCalculatorTypesBN254OperatorSetInfo, error) {
	var out []interface{}
	err := _IBN254TableCalculator.contract.Call(opts, &out, "calculateOperatorTable", operatorSet)

	if err != nil {
		return *new(IBN254TableCalculatorTypesBN254OperatorSetInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IBN254TableCalculatorTypesBN254OperatorSetInfo)).(*IBN254TableCalculatorTypesBN254OperatorSetInfo)

	return out0, err

}

// CalculateOperatorTable is a free data retrieval call binding the contract method 0x124c87e0.
//
// Solidity: function calculateOperatorTable((address,uint32) operatorSet) view returns((bytes32,uint256,(uint256,uint256),uint256[]) operatorSetInfo)
func (_IBN254TableCalculator *IBN254TableCalculatorSession) CalculateOperatorTable(operatorSet OperatorSet) (IBN254TableCalculatorTypesBN254OperatorSetInfo, error) {
	return _IBN254TableCalculator.Contract.CalculateOperatorTable(&_IBN254TableCalculator.CallOpts, operatorSet)
}

// CalculateOperatorTable is a free data retrieval call binding the contract method 0x124c87e0.
//
// Solidity: function calculateOperatorTable((address,uint32) operatorSet) view returns((bytes32,uint256,(uint256,uint256),uint256[]) operatorSetInfo)
func (_IBN254TableCalculator *IBN254TableCalculatorCallerSession) CalculateOperatorTable(operatorSet OperatorSet) (IBN254TableCalculatorTypesBN254OperatorSetInfo, error) {
	return _IBN254TableCalculator.Contract.CalculateOperatorTable(&_IBN254TableCalculator.CallOpts, operatorSet)
}

// CalculateOperatorTableBytes is a free data retrieval call binding the contract method 0x41ee6d0e.
//
// Solidity: function calculateOperatorTableBytes((address,uint32) operatorSet) view returns(bytes operatorTableBytes)
func (_IBN254TableCalculator *IBN254TableCalculatorCaller) CalculateOperatorTableBytes(opts *bind.CallOpts, operatorSet OperatorSet) ([]byte, error) {
	var out []interface{}
	err := _IBN254TableCalculator.contract.Call(opts, &out, "calculateOperatorTableBytes", operatorSet)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// CalculateOperatorTableBytes is a free data retrieval call binding the contract method 0x41ee6d0e.
//
// Solidity: function calculateOperatorTableBytes((address,uint32) operatorSet) view returns(bytes operatorTableBytes)
func (_IBN254TableCalculator *IBN254TableCalculatorSession) CalculateOperatorTableBytes(operatorSet OperatorSet) ([]byte, error) {
	return _IBN254TableCalculator.Contract.CalculateOperatorTableBytes(&_IBN254TableCalculator.CallOpts, operatorSet)
}

// CalculateOperatorTableBytes is a free data retrieval call binding the contract method 0x41ee6d0e.
//
// Solidity: function calculateOperatorTableBytes((address,uint32) operatorSet) view returns(bytes operatorTableBytes)
func (_IBN254TableCalculator *IBN254TableCalculatorCallerSession) CalculateOperatorTableBytes(operatorSet OperatorSet) ([]byte, error) {
	return _IBN254TableCalculator.Contract.CalculateOperatorTableBytes(&_IBN254TableCalculator.CallOpts, operatorSet)
}

// GetOperatorInfos is a free data retrieval call binding the contract method 0xcf2d90ef.
//
// Solidity: function getOperatorInfos((address,uint32) operatorSet) view returns(((uint256,uint256),uint256[])[] operatorInfos)
func (_IBN254TableCalculator *IBN254TableCalculatorCaller) GetOperatorInfos(opts *bind.CallOpts, operatorSet OperatorSet) ([]IBN254TableCalculatorTypesBN254OperatorInfo, error) {
	var out []interface{}
	err := _IBN254TableCalculator.contract.Call(opts, &out, "getOperatorInfos", operatorSet)

	if err != nil {
		return *new([]IBN254TableCalculatorTypesBN254OperatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IBN254TableCalculatorTypesBN254OperatorInfo)).(*[]IBN254TableCalculatorTypesBN254OperatorInfo)

	return out0, err

}

// GetOperatorInfos is a free data retrieval call binding the contract method 0xcf2d90ef.
//
// Solidity: function getOperatorInfos((address,uint32) operatorSet) view returns(((uint256,uint256),uint256[])[] operatorInfos)
func (_IBN254TableCalculator *IBN254TableCalculatorSession) GetOperatorInfos(operatorSet OperatorSet) ([]IBN254TableCalculatorTypesBN254OperatorInfo, error) {
	return _IBN254TableCalculator.Contract.GetOperatorInfos(&_IBN254TableCalculator.CallOpts, operatorSet)
}

// GetOperatorInfos is a free data retrieval call binding the contract method 0xcf2d90ef.
//
// Solidity: function getOperatorInfos((address,uint32) operatorSet) view returns(((uint256,uint256),uint256[])[] operatorInfos)
func (_IBN254TableCalculator *IBN254TableCalculatorCallerSession) GetOperatorInfos(operatorSet OperatorSet) ([]IBN254TableCalculatorTypesBN254OperatorInfo, error) {
	return _IBN254TableCalculator.Contract.GetOperatorInfos(&_IBN254TableCalculator.CallOpts, operatorSet)
}

// GetOperatorWeight is a free data retrieval call binding the contract method 0x1088794a.
//
// Solidity: function getOperatorWeight((address,uint32) operatorSet, address operator) view returns(uint256 weight)
func (_IBN254TableCalculator *IBN254TableCalculatorCaller) GetOperatorWeight(opts *bind.CallOpts, operatorSet OperatorSet, operator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IBN254TableCalculator.contract.Call(opts, &out, "getOperatorWeight", operatorSet, operator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOperatorWeight is a free data retrieval call binding the contract method 0x1088794a.
//
// Solidity: function getOperatorWeight((address,uint32) operatorSet, address operator) view returns(uint256 weight)
func (_IBN254TableCalculator *IBN254TableCalculatorSession) GetOperatorWeight(operatorSet OperatorSet, operator common.Address) (*big.Int, error) {
	return _IBN254TableCalculator.Contract.GetOperatorWeight(&_IBN254TableCalculator.CallOpts, operatorSet, operator)
}

// GetOperatorWeight is a free data retrieval call binding the contract method 0x1088794a.
//
// Solidity: function getOperatorWeight((address,uint32) operatorSet, address operator) view returns(uint256 weight)
func (_IBN254TableCalculator *IBN254TableCalculatorCallerSession) GetOperatorWeight(operatorSet OperatorSet, operator common.Address) (*big.Int, error) {
	return _IBN254TableCalculator.Contract.GetOperatorWeight(&_IBN254TableCalculator.CallOpts, operatorSet, operator)
}

// GetOperatorWeights is a free data retrieval call binding the contract method 0x71ca71d9.
//
// Solidity: function getOperatorWeights((address,uint32) operatorSet) view returns(address[] operators, uint256[][] weights)
func (_IBN254TableCalculator *IBN254TableCalculatorCaller) GetOperatorWeights(opts *bind.CallOpts, operatorSet OperatorSet) (struct {
	Operators []common.Address
	Weights   [][]*big.Int
}, error) {
	var out []interface{}
	err := _IBN254TableCalculator.contract.Call(opts, &out, "getOperatorWeights", operatorSet)

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
func (_IBN254TableCalculator *IBN254TableCalculatorSession) GetOperatorWeights(operatorSet OperatorSet) (struct {
	Operators []common.Address
	Weights   [][]*big.Int
}, error) {
	return _IBN254TableCalculator.Contract.GetOperatorWeights(&_IBN254TableCalculator.CallOpts, operatorSet)
}

// GetOperatorWeights is a free data retrieval call binding the contract method 0x71ca71d9.
//
// Solidity: function getOperatorWeights((address,uint32) operatorSet) view returns(address[] operators, uint256[][] weights)
func (_IBN254TableCalculator *IBN254TableCalculatorCallerSession) GetOperatorWeights(operatorSet OperatorSet) (struct {
	Operators []common.Address
	Weights   [][]*big.Int
}, error) {
	return _IBN254TableCalculator.Contract.GetOperatorWeights(&_IBN254TableCalculator.CallOpts, operatorSet)
}

// IBN254TableCalculatorLookaheadBlocksSetIterator is returned from FilterLookaheadBlocksSet and is used to iterate over the raw logs and unpacked data for LookaheadBlocksSet events raised by the IBN254TableCalculator contract.
type IBN254TableCalculatorLookaheadBlocksSetIterator struct {
	Event *IBN254TableCalculatorLookaheadBlocksSet // Event containing the contract specifics and raw log

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
func (it *IBN254TableCalculatorLookaheadBlocksSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBN254TableCalculatorLookaheadBlocksSet)
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
		it.Event = new(IBN254TableCalculatorLookaheadBlocksSet)
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
func (it *IBN254TableCalculatorLookaheadBlocksSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBN254TableCalculatorLookaheadBlocksSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBN254TableCalculatorLookaheadBlocksSet represents a LookaheadBlocksSet event raised by the IBN254TableCalculator contract.
type IBN254TableCalculatorLookaheadBlocksSet struct {
	LookaheadBlocks *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLookaheadBlocksSet is a free log retrieval operation binding the contract event 0xa41e64dd47db91b61b43ccea8b57d75abfa496f23efc708c22753c4bc9d68842.
//
// Solidity: event LookaheadBlocksSet(uint256 lookaheadBlocks)
func (_IBN254TableCalculator *IBN254TableCalculatorFilterer) FilterLookaheadBlocksSet(opts *bind.FilterOpts) (*IBN254TableCalculatorLookaheadBlocksSetIterator, error) {

	logs, sub, err := _IBN254TableCalculator.contract.FilterLogs(opts, "LookaheadBlocksSet")
	if err != nil {
		return nil, err
	}
	return &IBN254TableCalculatorLookaheadBlocksSetIterator{contract: _IBN254TableCalculator.contract, event: "LookaheadBlocksSet", logs: logs, sub: sub}, nil
}

// WatchLookaheadBlocksSet is a free log subscription operation binding the contract event 0xa41e64dd47db91b61b43ccea8b57d75abfa496f23efc708c22753c4bc9d68842.
//
// Solidity: event LookaheadBlocksSet(uint256 lookaheadBlocks)
func (_IBN254TableCalculator *IBN254TableCalculatorFilterer) WatchLookaheadBlocksSet(opts *bind.WatchOpts, sink chan<- *IBN254TableCalculatorLookaheadBlocksSet) (event.Subscription, error) {

	logs, sub, err := _IBN254TableCalculator.contract.WatchLogs(opts, "LookaheadBlocksSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBN254TableCalculatorLookaheadBlocksSet)
				if err := _IBN254TableCalculator.contract.UnpackLog(event, "LookaheadBlocksSet", log); err != nil {
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
func (_IBN254TableCalculator *IBN254TableCalculatorFilterer) ParseLookaheadBlocksSet(log types.Log) (*IBN254TableCalculatorLookaheadBlocksSet, error) {
	event := new(IBN254TableCalculatorLookaheadBlocksSet)
	if err := _IBN254TableCalculator.contract.UnpackLog(event, "LookaheadBlocksSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
