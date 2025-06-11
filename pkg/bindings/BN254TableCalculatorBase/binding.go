// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package BN254TableCalculatorBase

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

// BN254TableCalculatorBaseMetaData contains all meta data concerning the BN254TableCalculatorBase contract.
var BN254TableCalculatorBaseMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"calculateOperatorTable\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"operatorSetInfo\",\"type\":\"tuple\",\"internalType\":\"structIBN254TableCalculatorTypes.BN254OperatorSetInfo\",\"components\":[{\"name\":\"operatorInfoTreeRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"numOperators\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"aggregatePubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"totalWeights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateOperatorTableBytes\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"operatorTableBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorInfos\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIBN254TableCalculatorTypes.BN254OperatorInfo[]\",\"components\":[{\"name\":\"pubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorWeight\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorWeights\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"weights\",\"type\":\"uint256[][]\",\"internalType\":\"uint256[][]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"keyRegistrar\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIKeyRegistrar\"}],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"ECAddFailed\",\"inputs\":[]}]",
}

// BN254TableCalculatorBaseABI is the input ABI used to generate the binding from.
// Deprecated: Use BN254TableCalculatorBaseMetaData.ABI instead.
var BN254TableCalculatorBaseABI = BN254TableCalculatorBaseMetaData.ABI

// BN254TableCalculatorBase is an auto generated Go binding around an Ethereum contract.
type BN254TableCalculatorBase struct {
	BN254TableCalculatorBaseCaller     // Read-only binding to the contract
	BN254TableCalculatorBaseTransactor // Write-only binding to the contract
	BN254TableCalculatorBaseFilterer   // Log filterer for contract events
}

// BN254TableCalculatorBaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type BN254TableCalculatorBaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BN254TableCalculatorBaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BN254TableCalculatorBaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BN254TableCalculatorBaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BN254TableCalculatorBaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BN254TableCalculatorBaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BN254TableCalculatorBaseSession struct {
	Contract     *BN254TableCalculatorBase // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// BN254TableCalculatorBaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BN254TableCalculatorBaseCallerSession struct {
	Contract *BN254TableCalculatorBaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// BN254TableCalculatorBaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BN254TableCalculatorBaseTransactorSession struct {
	Contract     *BN254TableCalculatorBaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// BN254TableCalculatorBaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type BN254TableCalculatorBaseRaw struct {
	Contract *BN254TableCalculatorBase // Generic contract binding to access the raw methods on
}

// BN254TableCalculatorBaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BN254TableCalculatorBaseCallerRaw struct {
	Contract *BN254TableCalculatorBaseCaller // Generic read-only contract binding to access the raw methods on
}

// BN254TableCalculatorBaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BN254TableCalculatorBaseTransactorRaw struct {
	Contract *BN254TableCalculatorBaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBN254TableCalculatorBase creates a new instance of BN254TableCalculatorBase, bound to a specific deployed contract.
func NewBN254TableCalculatorBase(address common.Address, backend bind.ContractBackend) (*BN254TableCalculatorBase, error) {
	contract, err := bindBN254TableCalculatorBase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BN254TableCalculatorBase{BN254TableCalculatorBaseCaller: BN254TableCalculatorBaseCaller{contract: contract}, BN254TableCalculatorBaseTransactor: BN254TableCalculatorBaseTransactor{contract: contract}, BN254TableCalculatorBaseFilterer: BN254TableCalculatorBaseFilterer{contract: contract}}, nil
}

// NewBN254TableCalculatorBaseCaller creates a new read-only instance of BN254TableCalculatorBase, bound to a specific deployed contract.
func NewBN254TableCalculatorBaseCaller(address common.Address, caller bind.ContractCaller) (*BN254TableCalculatorBaseCaller, error) {
	contract, err := bindBN254TableCalculatorBase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BN254TableCalculatorBaseCaller{contract: contract}, nil
}

// NewBN254TableCalculatorBaseTransactor creates a new write-only instance of BN254TableCalculatorBase, bound to a specific deployed contract.
func NewBN254TableCalculatorBaseTransactor(address common.Address, transactor bind.ContractTransactor) (*BN254TableCalculatorBaseTransactor, error) {
	contract, err := bindBN254TableCalculatorBase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BN254TableCalculatorBaseTransactor{contract: contract}, nil
}

// NewBN254TableCalculatorBaseFilterer creates a new log filterer instance of BN254TableCalculatorBase, bound to a specific deployed contract.
func NewBN254TableCalculatorBaseFilterer(address common.Address, filterer bind.ContractFilterer) (*BN254TableCalculatorBaseFilterer, error) {
	contract, err := bindBN254TableCalculatorBase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BN254TableCalculatorBaseFilterer{contract: contract}, nil
}

// bindBN254TableCalculatorBase binds a generic wrapper to an already deployed contract.
func bindBN254TableCalculatorBase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BN254TableCalculatorBaseMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BN254TableCalculatorBase *BN254TableCalculatorBaseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BN254TableCalculatorBase.Contract.BN254TableCalculatorBaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BN254TableCalculatorBase *BN254TableCalculatorBaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BN254TableCalculatorBase.Contract.BN254TableCalculatorBaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BN254TableCalculatorBase *BN254TableCalculatorBaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BN254TableCalculatorBase.Contract.BN254TableCalculatorBaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BN254TableCalculatorBase *BN254TableCalculatorBaseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BN254TableCalculatorBase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BN254TableCalculatorBase *BN254TableCalculatorBaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BN254TableCalculatorBase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BN254TableCalculatorBase *BN254TableCalculatorBaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BN254TableCalculatorBase.Contract.contract.Transact(opts, method, params...)
}

// CalculateOperatorTable is a free data retrieval call binding the contract method 0x124c87e0.
//
// Solidity: function calculateOperatorTable((address,uint32) operatorSet) view returns((bytes32,uint256,(uint256,uint256),uint256[]) operatorSetInfo)
func (_BN254TableCalculatorBase *BN254TableCalculatorBaseCaller) CalculateOperatorTable(opts *bind.CallOpts, operatorSet OperatorSet) (IBN254TableCalculatorTypesBN254OperatorSetInfo, error) {
	var out []interface{}
	err := _BN254TableCalculatorBase.contract.Call(opts, &out, "calculateOperatorTable", operatorSet)

	if err != nil {
		return *new(IBN254TableCalculatorTypesBN254OperatorSetInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IBN254TableCalculatorTypesBN254OperatorSetInfo)).(*IBN254TableCalculatorTypesBN254OperatorSetInfo)

	return out0, err

}

// CalculateOperatorTable is a free data retrieval call binding the contract method 0x124c87e0.
//
// Solidity: function calculateOperatorTable((address,uint32) operatorSet) view returns((bytes32,uint256,(uint256,uint256),uint256[]) operatorSetInfo)
func (_BN254TableCalculatorBase *BN254TableCalculatorBaseSession) CalculateOperatorTable(operatorSet OperatorSet) (IBN254TableCalculatorTypesBN254OperatorSetInfo, error) {
	return _BN254TableCalculatorBase.Contract.CalculateOperatorTable(&_BN254TableCalculatorBase.CallOpts, operatorSet)
}

// CalculateOperatorTable is a free data retrieval call binding the contract method 0x124c87e0.
//
// Solidity: function calculateOperatorTable((address,uint32) operatorSet) view returns((bytes32,uint256,(uint256,uint256),uint256[]) operatorSetInfo)
func (_BN254TableCalculatorBase *BN254TableCalculatorBaseCallerSession) CalculateOperatorTable(operatorSet OperatorSet) (IBN254TableCalculatorTypesBN254OperatorSetInfo, error) {
	return _BN254TableCalculatorBase.Contract.CalculateOperatorTable(&_BN254TableCalculatorBase.CallOpts, operatorSet)
}

// CalculateOperatorTableBytes is a free data retrieval call binding the contract method 0x41ee6d0e.
//
// Solidity: function calculateOperatorTableBytes((address,uint32) operatorSet) view returns(bytes operatorTableBytes)
func (_BN254TableCalculatorBase *BN254TableCalculatorBaseCaller) CalculateOperatorTableBytes(opts *bind.CallOpts, operatorSet OperatorSet) ([]byte, error) {
	var out []interface{}
	err := _BN254TableCalculatorBase.contract.Call(opts, &out, "calculateOperatorTableBytes", operatorSet)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// CalculateOperatorTableBytes is a free data retrieval call binding the contract method 0x41ee6d0e.
//
// Solidity: function calculateOperatorTableBytes((address,uint32) operatorSet) view returns(bytes operatorTableBytes)
func (_BN254TableCalculatorBase *BN254TableCalculatorBaseSession) CalculateOperatorTableBytes(operatorSet OperatorSet) ([]byte, error) {
	return _BN254TableCalculatorBase.Contract.CalculateOperatorTableBytes(&_BN254TableCalculatorBase.CallOpts, operatorSet)
}

// CalculateOperatorTableBytes is a free data retrieval call binding the contract method 0x41ee6d0e.
//
// Solidity: function calculateOperatorTableBytes((address,uint32) operatorSet) view returns(bytes operatorTableBytes)
func (_BN254TableCalculatorBase *BN254TableCalculatorBaseCallerSession) CalculateOperatorTableBytes(operatorSet OperatorSet) ([]byte, error) {
	return _BN254TableCalculatorBase.Contract.CalculateOperatorTableBytes(&_BN254TableCalculatorBase.CallOpts, operatorSet)
}

// GetOperatorInfos is a free data retrieval call binding the contract method 0xcf2d90ef.
//
// Solidity: function getOperatorInfos((address,uint32) operatorSet) view returns(((uint256,uint256),uint256[])[])
func (_BN254TableCalculatorBase *BN254TableCalculatorBaseCaller) GetOperatorInfos(opts *bind.CallOpts, operatorSet OperatorSet) ([]IBN254TableCalculatorTypesBN254OperatorInfo, error) {
	var out []interface{}
	err := _BN254TableCalculatorBase.contract.Call(opts, &out, "getOperatorInfos", operatorSet)

	if err != nil {
		return *new([]IBN254TableCalculatorTypesBN254OperatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IBN254TableCalculatorTypesBN254OperatorInfo)).(*[]IBN254TableCalculatorTypesBN254OperatorInfo)

	return out0, err

}

// GetOperatorInfos is a free data retrieval call binding the contract method 0xcf2d90ef.
//
// Solidity: function getOperatorInfos((address,uint32) operatorSet) view returns(((uint256,uint256),uint256[])[])
func (_BN254TableCalculatorBase *BN254TableCalculatorBaseSession) GetOperatorInfos(operatorSet OperatorSet) ([]IBN254TableCalculatorTypesBN254OperatorInfo, error) {
	return _BN254TableCalculatorBase.Contract.GetOperatorInfos(&_BN254TableCalculatorBase.CallOpts, operatorSet)
}

// GetOperatorInfos is a free data retrieval call binding the contract method 0xcf2d90ef.
//
// Solidity: function getOperatorInfos((address,uint32) operatorSet) view returns(((uint256,uint256),uint256[])[])
func (_BN254TableCalculatorBase *BN254TableCalculatorBaseCallerSession) GetOperatorInfos(operatorSet OperatorSet) ([]IBN254TableCalculatorTypesBN254OperatorInfo, error) {
	return _BN254TableCalculatorBase.Contract.GetOperatorInfos(&_BN254TableCalculatorBase.CallOpts, operatorSet)
}

// GetOperatorWeight is a free data retrieval call binding the contract method 0x1088794a.
//
// Solidity: function getOperatorWeight((address,uint32) operatorSet, address operator) view returns(uint256 weight)
func (_BN254TableCalculatorBase *BN254TableCalculatorBaseCaller) GetOperatorWeight(opts *bind.CallOpts, operatorSet OperatorSet, operator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BN254TableCalculatorBase.contract.Call(opts, &out, "getOperatorWeight", operatorSet, operator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOperatorWeight is a free data retrieval call binding the contract method 0x1088794a.
//
// Solidity: function getOperatorWeight((address,uint32) operatorSet, address operator) view returns(uint256 weight)
func (_BN254TableCalculatorBase *BN254TableCalculatorBaseSession) GetOperatorWeight(operatorSet OperatorSet, operator common.Address) (*big.Int, error) {
	return _BN254TableCalculatorBase.Contract.GetOperatorWeight(&_BN254TableCalculatorBase.CallOpts, operatorSet, operator)
}

// GetOperatorWeight is a free data retrieval call binding the contract method 0x1088794a.
//
// Solidity: function getOperatorWeight((address,uint32) operatorSet, address operator) view returns(uint256 weight)
func (_BN254TableCalculatorBase *BN254TableCalculatorBaseCallerSession) GetOperatorWeight(operatorSet OperatorSet, operator common.Address) (*big.Int, error) {
	return _BN254TableCalculatorBase.Contract.GetOperatorWeight(&_BN254TableCalculatorBase.CallOpts, operatorSet, operator)
}

// GetOperatorWeights is a free data retrieval call binding the contract method 0x71ca71d9.
//
// Solidity: function getOperatorWeights((address,uint32) operatorSet) view returns(address[] operators, uint256[][] weights)
func (_BN254TableCalculatorBase *BN254TableCalculatorBaseCaller) GetOperatorWeights(opts *bind.CallOpts, operatorSet OperatorSet) (struct {
	Operators []common.Address
	Weights   [][]*big.Int
}, error) {
	var out []interface{}
	err := _BN254TableCalculatorBase.contract.Call(opts, &out, "getOperatorWeights", operatorSet)

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
func (_BN254TableCalculatorBase *BN254TableCalculatorBaseSession) GetOperatorWeights(operatorSet OperatorSet) (struct {
	Operators []common.Address
	Weights   [][]*big.Int
}, error) {
	return _BN254TableCalculatorBase.Contract.GetOperatorWeights(&_BN254TableCalculatorBase.CallOpts, operatorSet)
}

// GetOperatorWeights is a free data retrieval call binding the contract method 0x71ca71d9.
//
// Solidity: function getOperatorWeights((address,uint32) operatorSet) view returns(address[] operators, uint256[][] weights)
func (_BN254TableCalculatorBase *BN254TableCalculatorBaseCallerSession) GetOperatorWeights(operatorSet OperatorSet) (struct {
	Operators []common.Address
	Weights   [][]*big.Int
}, error) {
	return _BN254TableCalculatorBase.Contract.GetOperatorWeights(&_BN254TableCalculatorBase.CallOpts, operatorSet)
}

// KeyRegistrar is a free data retrieval call binding the contract method 0x3ec45c7e.
//
// Solidity: function keyRegistrar() view returns(address)
func (_BN254TableCalculatorBase *BN254TableCalculatorBaseCaller) KeyRegistrar(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BN254TableCalculatorBase.contract.Call(opts, &out, "keyRegistrar")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// KeyRegistrar is a free data retrieval call binding the contract method 0x3ec45c7e.
//
// Solidity: function keyRegistrar() view returns(address)
func (_BN254TableCalculatorBase *BN254TableCalculatorBaseSession) KeyRegistrar() (common.Address, error) {
	return _BN254TableCalculatorBase.Contract.KeyRegistrar(&_BN254TableCalculatorBase.CallOpts)
}

// KeyRegistrar is a free data retrieval call binding the contract method 0x3ec45c7e.
//
// Solidity: function keyRegistrar() view returns(address)
func (_BN254TableCalculatorBase *BN254TableCalculatorBaseCallerSession) KeyRegistrar() (common.Address, error) {
	return _BN254TableCalculatorBase.Contract.KeyRegistrar(&_BN254TableCalculatorBase.CallOpts)
}
