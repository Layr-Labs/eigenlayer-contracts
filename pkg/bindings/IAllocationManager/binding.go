// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IAllocationManager

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

// IAllocationManagerMagnitudeAllocation is an auto generated low-level Go binding around an user-defined struct.
type IAllocationManagerMagnitudeAllocation struct {
	Strategy               common.Address
	ExpectedTotalMagnitude uint64
	OperatorSets           []OperatorSet
	Magnitudes             []uint64
}

// ISignatureUtilsSignatureWithSaltAndExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithSaltAndExpiry struct {
	Signature []byte
	Salt      [32]byte
	Expiry    *big.Int
}

// OperatorSet is an auto generated low-level Go binding around an user-defined struct.
type OperatorSet struct {
	Avs           common.Address
	OperatorSetId uint32
}

// IAllocationManagerMetaData contains all meta data concerning the IAllocationManager contract.
var IAllocationManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"allocationDelay\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"isSet\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"delay\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateMagnitudeAllocationDigestHash\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allocations\",\"type\":\"tuple[]\",\"internalType\":\"structIAllocationManager.MagnitudeAllocation[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"expectedTotalMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"operatorSets\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"magnitudes\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}]},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cancelSalt\",\"inputs\":[{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"domainSeparator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocatableMagnitude\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"numToComplete\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentSlashableMagnitudes\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"\",\"type\":\"uint64[][]\",\"internalType\":\"uint64[][]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPendingAllocations\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"operatorSets\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPendingDeallocations\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"operatorSets\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSlashableMagnitudes\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"timestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"\",\"type\":\"uint64[][]\",\"internalType\":\"uint64[][]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSlashablePPM\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"timestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"linear\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint24[]\",\"internalType\":\"uint24[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalAndAllocatedMagnitudes\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalMagnitude\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalMagnitudeAtTimestamp\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"timestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalMagnitudes\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalMagnitudesAtTimestamp\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"timestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperatorSlashable\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"modifyAllocations\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allocations\",\"type\":\"tuple[]\",\"internalType\":\"structIAllocationManager.MagnitudeAllocation[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"expectedTotalMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"operatorSets\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"magnitudes\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}]},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAllocationDelay\",\"inputs\":[{\"name\":\"delay\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slashOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"bipsToSlash\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateFreeMagnitude\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"numToComplete\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllocationDelaySet\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"delay\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MagnitudeAllocated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"magnitudeToAllocate\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"effectTimestamp\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MagnitudeDeallocationCompleted\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"freeMagnitudeAdded\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MagnitudeQueueDeallocated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"magnitudeToDeallocate\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"completableTimestamp\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSetCreated\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSlashed\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"bipsToSlash\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InputArrayLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientAllocatableMagnitude\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidExpectedTotalMagnitude\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperatorSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorNotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorSetsNotInAscendingOrder\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PendingAllocationOrDeallocation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SaltSpent\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignatureExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UninitializedAllocationDelay\",\"inputs\":[]}]",
}

// IAllocationManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use IAllocationManagerMetaData.ABI instead.
var IAllocationManagerABI = IAllocationManagerMetaData.ABI

// IAllocationManager is an auto generated Go binding around an Ethereum contract.
type IAllocationManager struct {
	IAllocationManagerCaller     // Read-only binding to the contract
	IAllocationManagerTransactor // Write-only binding to the contract
	IAllocationManagerFilterer   // Log filterer for contract events
}

// IAllocationManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAllocationManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAllocationManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAllocationManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAllocationManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAllocationManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAllocationManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAllocationManagerSession struct {
	Contract     *IAllocationManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IAllocationManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAllocationManagerCallerSession struct {
	Contract *IAllocationManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// IAllocationManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAllocationManagerTransactorSession struct {
	Contract     *IAllocationManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// IAllocationManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAllocationManagerRaw struct {
	Contract *IAllocationManager // Generic contract binding to access the raw methods on
}

// IAllocationManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAllocationManagerCallerRaw struct {
	Contract *IAllocationManagerCaller // Generic read-only contract binding to access the raw methods on
}

// IAllocationManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAllocationManagerTransactorRaw struct {
	Contract *IAllocationManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAllocationManager creates a new instance of IAllocationManager, bound to a specific deployed contract.
func NewIAllocationManager(address common.Address, backend bind.ContractBackend) (*IAllocationManager, error) {
	contract, err := bindIAllocationManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAllocationManager{IAllocationManagerCaller: IAllocationManagerCaller{contract: contract}, IAllocationManagerTransactor: IAllocationManagerTransactor{contract: contract}, IAllocationManagerFilterer: IAllocationManagerFilterer{contract: contract}}, nil
}

// NewIAllocationManagerCaller creates a new read-only instance of IAllocationManager, bound to a specific deployed contract.
func NewIAllocationManagerCaller(address common.Address, caller bind.ContractCaller) (*IAllocationManagerCaller, error) {
	contract, err := bindIAllocationManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAllocationManagerCaller{contract: contract}, nil
}

// NewIAllocationManagerTransactor creates a new write-only instance of IAllocationManager, bound to a specific deployed contract.
func NewIAllocationManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IAllocationManagerTransactor, error) {
	contract, err := bindIAllocationManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAllocationManagerTransactor{contract: contract}, nil
}

// NewIAllocationManagerFilterer creates a new log filterer instance of IAllocationManager, bound to a specific deployed contract.
func NewIAllocationManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IAllocationManagerFilterer, error) {
	contract, err := bindIAllocationManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAllocationManagerFilterer{contract: contract}, nil
}

// bindIAllocationManager binds a generic wrapper to an already deployed contract.
func bindIAllocationManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IAllocationManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAllocationManager *IAllocationManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAllocationManager.Contract.IAllocationManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAllocationManager *IAllocationManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAllocationManager.Contract.IAllocationManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAllocationManager *IAllocationManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAllocationManager.Contract.IAllocationManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAllocationManager *IAllocationManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAllocationManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAllocationManager *IAllocationManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAllocationManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAllocationManager *IAllocationManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAllocationManager.Contract.contract.Transact(opts, method, params...)
}

// AllocationDelay is a free data retrieval call binding the contract method 0xbe4e1fd3.
//
// Solidity: function allocationDelay(address operator) view returns(bool isSet, uint32 delay)
func (_IAllocationManager *IAllocationManagerCaller) AllocationDelay(opts *bind.CallOpts, operator common.Address) (struct {
	IsSet bool
	Delay uint32
}, error) {
	var out []interface{}
	err := _IAllocationManager.contract.Call(opts, &out, "allocationDelay", operator)

	outstruct := new(struct {
		IsSet bool
		Delay uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsSet = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Delay = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

// AllocationDelay is a free data retrieval call binding the contract method 0xbe4e1fd3.
//
// Solidity: function allocationDelay(address operator) view returns(bool isSet, uint32 delay)
func (_IAllocationManager *IAllocationManagerSession) AllocationDelay(operator common.Address) (struct {
	IsSet bool
	Delay uint32
}, error) {
	return _IAllocationManager.Contract.AllocationDelay(&_IAllocationManager.CallOpts, operator)
}

// AllocationDelay is a free data retrieval call binding the contract method 0xbe4e1fd3.
//
// Solidity: function allocationDelay(address operator) view returns(bool isSet, uint32 delay)
func (_IAllocationManager *IAllocationManagerCallerSession) AllocationDelay(operator common.Address) (struct {
	IsSet bool
	Delay uint32
}, error) {
	return _IAllocationManager.Contract.AllocationDelay(&_IAllocationManager.CallOpts, operator)
}

// CalculateMagnitudeAllocationDigestHash is a free data retrieval call binding the contract method 0x686b686e.
//
// Solidity: function calculateMagnitudeAllocationDigestHash(address operator, (address,uint64,(address,uint32)[],uint64[])[] allocations, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_IAllocationManager *IAllocationManagerCaller) CalculateMagnitudeAllocationDigestHash(opts *bind.CallOpts, operator common.Address, allocations []IAllocationManagerMagnitudeAllocation, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _IAllocationManager.contract.Call(opts, &out, "calculateMagnitudeAllocationDigestHash", operator, allocations, salt, expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateMagnitudeAllocationDigestHash is a free data retrieval call binding the contract method 0x686b686e.
//
// Solidity: function calculateMagnitudeAllocationDigestHash(address operator, (address,uint64,(address,uint32)[],uint64[])[] allocations, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_IAllocationManager *IAllocationManagerSession) CalculateMagnitudeAllocationDigestHash(operator common.Address, allocations []IAllocationManagerMagnitudeAllocation, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _IAllocationManager.Contract.CalculateMagnitudeAllocationDigestHash(&_IAllocationManager.CallOpts, operator, allocations, salt, expiry)
}

// CalculateMagnitudeAllocationDigestHash is a free data retrieval call binding the contract method 0x686b686e.
//
// Solidity: function calculateMagnitudeAllocationDigestHash(address operator, (address,uint64,(address,uint32)[],uint64[])[] allocations, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_IAllocationManager *IAllocationManagerCallerSession) CalculateMagnitudeAllocationDigestHash(operator common.Address, allocations []IAllocationManagerMagnitudeAllocation, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _IAllocationManager.Contract.CalculateMagnitudeAllocationDigestHash(&_IAllocationManager.CallOpts, operator, allocations, salt, expiry)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_IAllocationManager *IAllocationManagerCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IAllocationManager.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_IAllocationManager *IAllocationManagerSession) DomainSeparator() ([32]byte, error) {
	return _IAllocationManager.Contract.DomainSeparator(&_IAllocationManager.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_IAllocationManager *IAllocationManagerCallerSession) DomainSeparator() ([32]byte, error) {
	return _IAllocationManager.Contract.DomainSeparator(&_IAllocationManager.CallOpts)
}

// GetAllocatableMagnitude is a free data retrieval call binding the contract method 0x3f78612f.
//
// Solidity: function getAllocatableMagnitude(address operator, address strategy, uint16 numToComplete) view returns(uint64)
func (_IAllocationManager *IAllocationManagerCaller) GetAllocatableMagnitude(opts *bind.CallOpts, operator common.Address, strategy common.Address, numToComplete uint16) (uint64, error) {
	var out []interface{}
	err := _IAllocationManager.contract.Call(opts, &out, "getAllocatableMagnitude", operator, strategy, numToComplete)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetAllocatableMagnitude is a free data retrieval call binding the contract method 0x3f78612f.
//
// Solidity: function getAllocatableMagnitude(address operator, address strategy, uint16 numToComplete) view returns(uint64)
func (_IAllocationManager *IAllocationManagerSession) GetAllocatableMagnitude(operator common.Address, strategy common.Address, numToComplete uint16) (uint64, error) {
	return _IAllocationManager.Contract.GetAllocatableMagnitude(&_IAllocationManager.CallOpts, operator, strategy, numToComplete)
}

// GetAllocatableMagnitude is a free data retrieval call binding the contract method 0x3f78612f.
//
// Solidity: function getAllocatableMagnitude(address operator, address strategy, uint16 numToComplete) view returns(uint64)
func (_IAllocationManager *IAllocationManagerCallerSession) GetAllocatableMagnitude(operator common.Address, strategy common.Address, numToComplete uint16) (uint64, error) {
	return _IAllocationManager.Contract.GetAllocatableMagnitude(&_IAllocationManager.CallOpts, operator, strategy, numToComplete)
}

// GetCurrentSlashableMagnitudes is a free data retrieval call binding the contract method 0x1d95b524.
//
// Solidity: function getCurrentSlashableMagnitudes(address operator, address[] strategies) view returns((address,uint32)[], uint64[][])
func (_IAllocationManager *IAllocationManagerCaller) GetCurrentSlashableMagnitudes(opts *bind.CallOpts, operator common.Address, strategies []common.Address) ([]OperatorSet, [][]uint64, error) {
	var out []interface{}
	err := _IAllocationManager.contract.Call(opts, &out, "getCurrentSlashableMagnitudes", operator, strategies)

	if err != nil {
		return *new([]OperatorSet), *new([][]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]OperatorSet)).(*[]OperatorSet)
	out1 := *abi.ConvertType(out[1], new([][]uint64)).(*[][]uint64)

	return out0, out1, err

}

// GetCurrentSlashableMagnitudes is a free data retrieval call binding the contract method 0x1d95b524.
//
// Solidity: function getCurrentSlashableMagnitudes(address operator, address[] strategies) view returns((address,uint32)[], uint64[][])
func (_IAllocationManager *IAllocationManagerSession) GetCurrentSlashableMagnitudes(operator common.Address, strategies []common.Address) ([]OperatorSet, [][]uint64, error) {
	return _IAllocationManager.Contract.GetCurrentSlashableMagnitudes(&_IAllocationManager.CallOpts, operator, strategies)
}

// GetCurrentSlashableMagnitudes is a free data retrieval call binding the contract method 0x1d95b524.
//
// Solidity: function getCurrentSlashableMagnitudes(address operator, address[] strategies) view returns((address,uint32)[], uint64[][])
func (_IAllocationManager *IAllocationManagerCallerSession) GetCurrentSlashableMagnitudes(operator common.Address, strategies []common.Address) ([]OperatorSet, [][]uint64, error) {
	return _IAllocationManager.Contract.GetCurrentSlashableMagnitudes(&_IAllocationManager.CallOpts, operator, strategies)
}

// GetPendingAllocations is a free data retrieval call binding the contract method 0x67d973ef.
//
// Solidity: function getPendingAllocations(address operator, address strategy, (address,uint32)[] operatorSets) view returns(uint64[], uint32[])
func (_IAllocationManager *IAllocationManagerCaller) GetPendingAllocations(opts *bind.CallOpts, operator common.Address, strategy common.Address, operatorSets []OperatorSet) ([]uint64, []uint32, error) {
	var out []interface{}
	err := _IAllocationManager.contract.Call(opts, &out, "getPendingAllocations", operator, strategy, operatorSets)

	if err != nil {
		return *new([]uint64), *new([]uint32), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)
	out1 := *abi.ConvertType(out[1], new([]uint32)).(*[]uint32)

	return out0, out1, err

}

// GetPendingAllocations is a free data retrieval call binding the contract method 0x67d973ef.
//
// Solidity: function getPendingAllocations(address operator, address strategy, (address,uint32)[] operatorSets) view returns(uint64[], uint32[])
func (_IAllocationManager *IAllocationManagerSession) GetPendingAllocations(operator common.Address, strategy common.Address, operatorSets []OperatorSet) ([]uint64, []uint32, error) {
	return _IAllocationManager.Contract.GetPendingAllocations(&_IAllocationManager.CallOpts, operator, strategy, operatorSets)
}

// GetPendingAllocations is a free data retrieval call binding the contract method 0x67d973ef.
//
// Solidity: function getPendingAllocations(address operator, address strategy, (address,uint32)[] operatorSets) view returns(uint64[], uint32[])
func (_IAllocationManager *IAllocationManagerCallerSession) GetPendingAllocations(operator common.Address, strategy common.Address, operatorSets []OperatorSet) ([]uint64, []uint32, error) {
	return _IAllocationManager.Contract.GetPendingAllocations(&_IAllocationManager.CallOpts, operator, strategy, operatorSets)
}

// GetPendingDeallocations is a free data retrieval call binding the contract method 0x07053717.
//
// Solidity: function getPendingDeallocations(address operator, address strategy, (address,uint32)[] operatorSets) view returns(uint64[], uint32[])
func (_IAllocationManager *IAllocationManagerCaller) GetPendingDeallocations(opts *bind.CallOpts, operator common.Address, strategy common.Address, operatorSets []OperatorSet) ([]uint64, []uint32, error) {
	var out []interface{}
	err := _IAllocationManager.contract.Call(opts, &out, "getPendingDeallocations", operator, strategy, operatorSets)

	if err != nil {
		return *new([]uint64), *new([]uint32), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)
	out1 := *abi.ConvertType(out[1], new([]uint32)).(*[]uint32)

	return out0, out1, err

}

// GetPendingDeallocations is a free data retrieval call binding the contract method 0x07053717.
//
// Solidity: function getPendingDeallocations(address operator, address strategy, (address,uint32)[] operatorSets) view returns(uint64[], uint32[])
func (_IAllocationManager *IAllocationManagerSession) GetPendingDeallocations(operator common.Address, strategy common.Address, operatorSets []OperatorSet) ([]uint64, []uint32, error) {
	return _IAllocationManager.Contract.GetPendingDeallocations(&_IAllocationManager.CallOpts, operator, strategy, operatorSets)
}

// GetPendingDeallocations is a free data retrieval call binding the contract method 0x07053717.
//
// Solidity: function getPendingDeallocations(address operator, address strategy, (address,uint32)[] operatorSets) view returns(uint64[], uint32[])
func (_IAllocationManager *IAllocationManagerCallerSession) GetPendingDeallocations(operator common.Address, strategy common.Address, operatorSets []OperatorSet) ([]uint64, []uint32, error) {
	return _IAllocationManager.Contract.GetPendingDeallocations(&_IAllocationManager.CallOpts, operator, strategy, operatorSets)
}

// GetSlashableMagnitudes is a free data retrieval call binding the contract method 0x43b0592d.
//
// Solidity: function getSlashableMagnitudes(address operator, address[] strategies, uint32 timestamp) view returns((address,uint32)[], uint64[][])
func (_IAllocationManager *IAllocationManagerCaller) GetSlashableMagnitudes(opts *bind.CallOpts, operator common.Address, strategies []common.Address, timestamp uint32) ([]OperatorSet, [][]uint64, error) {
	var out []interface{}
	err := _IAllocationManager.contract.Call(opts, &out, "getSlashableMagnitudes", operator, strategies, timestamp)

	if err != nil {
		return *new([]OperatorSet), *new([][]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]OperatorSet)).(*[]OperatorSet)
	out1 := *abi.ConvertType(out[1], new([][]uint64)).(*[][]uint64)

	return out0, out1, err

}

// GetSlashableMagnitudes is a free data retrieval call binding the contract method 0x43b0592d.
//
// Solidity: function getSlashableMagnitudes(address operator, address[] strategies, uint32 timestamp) view returns((address,uint32)[], uint64[][])
func (_IAllocationManager *IAllocationManagerSession) GetSlashableMagnitudes(operator common.Address, strategies []common.Address, timestamp uint32) ([]OperatorSet, [][]uint64, error) {
	return _IAllocationManager.Contract.GetSlashableMagnitudes(&_IAllocationManager.CallOpts, operator, strategies, timestamp)
}

// GetSlashableMagnitudes is a free data retrieval call binding the contract method 0x43b0592d.
//
// Solidity: function getSlashableMagnitudes(address operator, address[] strategies, uint32 timestamp) view returns((address,uint32)[], uint64[][])
func (_IAllocationManager *IAllocationManagerCallerSession) GetSlashableMagnitudes(operator common.Address, strategies []common.Address, timestamp uint32) ([]OperatorSet, [][]uint64, error) {
	return _IAllocationManager.Contract.GetSlashableMagnitudes(&_IAllocationManager.CallOpts, operator, strategies, timestamp)
}

// GetSlashablePPM is a free data retrieval call binding the contract method 0x1c699d65.
//
// Solidity: function getSlashablePPM(address operator, (address,uint32) operatorSet, address[] strategies, uint32 timestamp, bool linear) view returns(uint24[])
func (_IAllocationManager *IAllocationManagerCaller) GetSlashablePPM(opts *bind.CallOpts, operator common.Address, operatorSet OperatorSet, strategies []common.Address, timestamp uint32, linear bool) ([]*big.Int, error) {
	var out []interface{}
	err := _IAllocationManager.contract.Call(opts, &out, "getSlashablePPM", operator, operatorSet, strategies, timestamp, linear)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetSlashablePPM is a free data retrieval call binding the contract method 0x1c699d65.
//
// Solidity: function getSlashablePPM(address operator, (address,uint32) operatorSet, address[] strategies, uint32 timestamp, bool linear) view returns(uint24[])
func (_IAllocationManager *IAllocationManagerSession) GetSlashablePPM(operator common.Address, operatorSet OperatorSet, strategies []common.Address, timestamp uint32, linear bool) ([]*big.Int, error) {
	return _IAllocationManager.Contract.GetSlashablePPM(&_IAllocationManager.CallOpts, operator, operatorSet, strategies, timestamp, linear)
}

// GetSlashablePPM is a free data retrieval call binding the contract method 0x1c699d65.
//
// Solidity: function getSlashablePPM(address operator, (address,uint32) operatorSet, address[] strategies, uint32 timestamp, bool linear) view returns(uint24[])
func (_IAllocationManager *IAllocationManagerCallerSession) GetSlashablePPM(operator common.Address, operatorSet OperatorSet, strategies []common.Address, timestamp uint32, linear bool) ([]*big.Int, error) {
	return _IAllocationManager.Contract.GetSlashablePPM(&_IAllocationManager.CallOpts, operator, operatorSet, strategies, timestamp, linear)
}

// GetTotalAndAllocatedMagnitudes is a free data retrieval call binding the contract method 0xc2348318.
//
// Solidity: function getTotalAndAllocatedMagnitudes(address operator, (address,uint32) operatorSet, address[] strategies) view returns(uint64[], uint64[])
func (_IAllocationManager *IAllocationManagerCaller) GetTotalAndAllocatedMagnitudes(opts *bind.CallOpts, operator common.Address, operatorSet OperatorSet, strategies []common.Address) ([]uint64, []uint64, error) {
	var out []interface{}
	err := _IAllocationManager.contract.Call(opts, &out, "getTotalAndAllocatedMagnitudes", operator, operatorSet, strategies)

	if err != nil {
		return *new([]uint64), *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)
	out1 := *abi.ConvertType(out[1], new([]uint64)).(*[]uint64)

	return out0, out1, err

}

// GetTotalAndAllocatedMagnitudes is a free data retrieval call binding the contract method 0xc2348318.
//
// Solidity: function getTotalAndAllocatedMagnitudes(address operator, (address,uint32) operatorSet, address[] strategies) view returns(uint64[], uint64[])
func (_IAllocationManager *IAllocationManagerSession) GetTotalAndAllocatedMagnitudes(operator common.Address, operatorSet OperatorSet, strategies []common.Address) ([]uint64, []uint64, error) {
	return _IAllocationManager.Contract.GetTotalAndAllocatedMagnitudes(&_IAllocationManager.CallOpts, operator, operatorSet, strategies)
}

// GetTotalAndAllocatedMagnitudes is a free data retrieval call binding the contract method 0xc2348318.
//
// Solidity: function getTotalAndAllocatedMagnitudes(address operator, (address,uint32) operatorSet, address[] strategies) view returns(uint64[], uint64[])
func (_IAllocationManager *IAllocationManagerCallerSession) GetTotalAndAllocatedMagnitudes(operator common.Address, operatorSet OperatorSet, strategies []common.Address) ([]uint64, []uint64, error) {
	return _IAllocationManager.Contract.GetTotalAndAllocatedMagnitudes(&_IAllocationManager.CallOpts, operator, operatorSet, strategies)
}

// GetTotalMagnitude is a free data retrieval call binding the contract method 0xb47265e2.
//
// Solidity: function getTotalMagnitude(address operator, address strategy) view returns(uint64)
func (_IAllocationManager *IAllocationManagerCaller) GetTotalMagnitude(opts *bind.CallOpts, operator common.Address, strategy common.Address) (uint64, error) {
	var out []interface{}
	err := _IAllocationManager.contract.Call(opts, &out, "getTotalMagnitude", operator, strategy)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetTotalMagnitude is a free data retrieval call binding the contract method 0xb47265e2.
//
// Solidity: function getTotalMagnitude(address operator, address strategy) view returns(uint64)
func (_IAllocationManager *IAllocationManagerSession) GetTotalMagnitude(operator common.Address, strategy common.Address) (uint64, error) {
	return _IAllocationManager.Contract.GetTotalMagnitude(&_IAllocationManager.CallOpts, operator, strategy)
}

// GetTotalMagnitude is a free data retrieval call binding the contract method 0xb47265e2.
//
// Solidity: function getTotalMagnitude(address operator, address strategy) view returns(uint64)
func (_IAllocationManager *IAllocationManagerCallerSession) GetTotalMagnitude(operator common.Address, strategy common.Address) (uint64, error) {
	return _IAllocationManager.Contract.GetTotalMagnitude(&_IAllocationManager.CallOpts, operator, strategy)
}

// GetTotalMagnitudeAtTimestamp is a free data retrieval call binding the contract method 0xdabe97c7.
//
// Solidity: function getTotalMagnitudeAtTimestamp(address operator, address strategy, uint32 timestamp) view returns(uint64)
func (_IAllocationManager *IAllocationManagerCaller) GetTotalMagnitudeAtTimestamp(opts *bind.CallOpts, operator common.Address, strategy common.Address, timestamp uint32) (uint64, error) {
	var out []interface{}
	err := _IAllocationManager.contract.Call(opts, &out, "getTotalMagnitudeAtTimestamp", operator, strategy, timestamp)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetTotalMagnitudeAtTimestamp is a free data retrieval call binding the contract method 0xdabe97c7.
//
// Solidity: function getTotalMagnitudeAtTimestamp(address operator, address strategy, uint32 timestamp) view returns(uint64)
func (_IAllocationManager *IAllocationManagerSession) GetTotalMagnitudeAtTimestamp(operator common.Address, strategy common.Address, timestamp uint32) (uint64, error) {
	return _IAllocationManager.Contract.GetTotalMagnitudeAtTimestamp(&_IAllocationManager.CallOpts, operator, strategy, timestamp)
}

// GetTotalMagnitudeAtTimestamp is a free data retrieval call binding the contract method 0xdabe97c7.
//
// Solidity: function getTotalMagnitudeAtTimestamp(address operator, address strategy, uint32 timestamp) view returns(uint64)
func (_IAllocationManager *IAllocationManagerCallerSession) GetTotalMagnitudeAtTimestamp(operator common.Address, strategy common.Address, timestamp uint32) (uint64, error) {
	return _IAllocationManager.Contract.GetTotalMagnitudeAtTimestamp(&_IAllocationManager.CallOpts, operator, strategy, timestamp)
}

// GetTotalMagnitudes is a free data retrieval call binding the contract method 0x39a9a3ed.
//
// Solidity: function getTotalMagnitudes(address operator, address[] strategies) view returns(uint64[])
func (_IAllocationManager *IAllocationManagerCaller) GetTotalMagnitudes(opts *bind.CallOpts, operator common.Address, strategies []common.Address) ([]uint64, error) {
	var out []interface{}
	err := _IAllocationManager.contract.Call(opts, &out, "getTotalMagnitudes", operator, strategies)

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

// GetTotalMagnitudes is a free data retrieval call binding the contract method 0x39a9a3ed.
//
// Solidity: function getTotalMagnitudes(address operator, address[] strategies) view returns(uint64[])
func (_IAllocationManager *IAllocationManagerSession) GetTotalMagnitudes(operator common.Address, strategies []common.Address) ([]uint64, error) {
	return _IAllocationManager.Contract.GetTotalMagnitudes(&_IAllocationManager.CallOpts, operator, strategies)
}

// GetTotalMagnitudes is a free data retrieval call binding the contract method 0x39a9a3ed.
//
// Solidity: function getTotalMagnitudes(address operator, address[] strategies) view returns(uint64[])
func (_IAllocationManager *IAllocationManagerCallerSession) GetTotalMagnitudes(operator common.Address, strategies []common.Address) ([]uint64, error) {
	return _IAllocationManager.Contract.GetTotalMagnitudes(&_IAllocationManager.CallOpts, operator, strategies)
}

// GetTotalMagnitudesAtTimestamp is a free data retrieval call binding the contract method 0x858d0b47.
//
// Solidity: function getTotalMagnitudesAtTimestamp(address operator, address[] strategies, uint32 timestamp) view returns(uint64[])
func (_IAllocationManager *IAllocationManagerCaller) GetTotalMagnitudesAtTimestamp(opts *bind.CallOpts, operator common.Address, strategies []common.Address, timestamp uint32) ([]uint64, error) {
	var out []interface{}
	err := _IAllocationManager.contract.Call(opts, &out, "getTotalMagnitudesAtTimestamp", operator, strategies, timestamp)

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

// GetTotalMagnitudesAtTimestamp is a free data retrieval call binding the contract method 0x858d0b47.
//
// Solidity: function getTotalMagnitudesAtTimestamp(address operator, address[] strategies, uint32 timestamp) view returns(uint64[])
func (_IAllocationManager *IAllocationManagerSession) GetTotalMagnitudesAtTimestamp(operator common.Address, strategies []common.Address, timestamp uint32) ([]uint64, error) {
	return _IAllocationManager.Contract.GetTotalMagnitudesAtTimestamp(&_IAllocationManager.CallOpts, operator, strategies, timestamp)
}

// GetTotalMagnitudesAtTimestamp is a free data retrieval call binding the contract method 0x858d0b47.
//
// Solidity: function getTotalMagnitudesAtTimestamp(address operator, address[] strategies, uint32 timestamp) view returns(uint64[])
func (_IAllocationManager *IAllocationManagerCallerSession) GetTotalMagnitudesAtTimestamp(operator common.Address, strategies []common.Address, timestamp uint32) ([]uint64, error) {
	return _IAllocationManager.Contract.GetTotalMagnitudesAtTimestamp(&_IAllocationManager.CallOpts, operator, strategies, timestamp)
}

// IsOperatorSlashable is a free data retrieval call binding the contract method 0x1352c3e6.
//
// Solidity: function isOperatorSlashable(address operator, (address,uint32) operatorSet) view returns(bool)
func (_IAllocationManager *IAllocationManagerCaller) IsOperatorSlashable(opts *bind.CallOpts, operator common.Address, operatorSet OperatorSet) (bool, error) {
	var out []interface{}
	err := _IAllocationManager.contract.Call(opts, &out, "isOperatorSlashable", operator, operatorSet)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperatorSlashable is a free data retrieval call binding the contract method 0x1352c3e6.
//
// Solidity: function isOperatorSlashable(address operator, (address,uint32) operatorSet) view returns(bool)
func (_IAllocationManager *IAllocationManagerSession) IsOperatorSlashable(operator common.Address, operatorSet OperatorSet) (bool, error) {
	return _IAllocationManager.Contract.IsOperatorSlashable(&_IAllocationManager.CallOpts, operator, operatorSet)
}

// IsOperatorSlashable is a free data retrieval call binding the contract method 0x1352c3e6.
//
// Solidity: function isOperatorSlashable(address operator, (address,uint32) operatorSet) view returns(bool)
func (_IAllocationManager *IAllocationManagerCallerSession) IsOperatorSlashable(operator common.Address, operatorSet OperatorSet) (bool, error) {
	return _IAllocationManager.Contract.IsOperatorSlashable(&_IAllocationManager.CallOpts, operator, operatorSet)
}

// CancelSalt is a paid mutator transaction binding the contract method 0xec76f442.
//
// Solidity: function cancelSalt(bytes32 salt) returns()
func (_IAllocationManager *IAllocationManagerTransactor) CancelSalt(opts *bind.TransactOpts, salt [32]byte) (*types.Transaction, error) {
	return _IAllocationManager.contract.Transact(opts, "cancelSalt", salt)
}

// CancelSalt is a paid mutator transaction binding the contract method 0xec76f442.
//
// Solidity: function cancelSalt(bytes32 salt) returns()
func (_IAllocationManager *IAllocationManagerSession) CancelSalt(salt [32]byte) (*types.Transaction, error) {
	return _IAllocationManager.Contract.CancelSalt(&_IAllocationManager.TransactOpts, salt)
}

// CancelSalt is a paid mutator transaction binding the contract method 0xec76f442.
//
// Solidity: function cancelSalt(bytes32 salt) returns()
func (_IAllocationManager *IAllocationManagerTransactorSession) CancelSalt(salt [32]byte) (*types.Transaction, error) {
	return _IAllocationManager.Contract.CancelSalt(&_IAllocationManager.TransactOpts, salt)
}

// ModifyAllocations is a paid mutator transaction binding the contract method 0xf8e91d16.
//
// Solidity: function modifyAllocations(address operator, (address,uint64,(address,uint32)[],uint64[])[] allocations, (bytes,bytes32,uint256) operatorSignature) returns()
func (_IAllocationManager *IAllocationManagerTransactor) ModifyAllocations(opts *bind.TransactOpts, operator common.Address, allocations []IAllocationManagerMagnitudeAllocation, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _IAllocationManager.contract.Transact(opts, "modifyAllocations", operator, allocations, operatorSignature)
}

// ModifyAllocations is a paid mutator transaction binding the contract method 0xf8e91d16.
//
// Solidity: function modifyAllocations(address operator, (address,uint64,(address,uint32)[],uint64[])[] allocations, (bytes,bytes32,uint256) operatorSignature) returns()
func (_IAllocationManager *IAllocationManagerSession) ModifyAllocations(operator common.Address, allocations []IAllocationManagerMagnitudeAllocation, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _IAllocationManager.Contract.ModifyAllocations(&_IAllocationManager.TransactOpts, operator, allocations, operatorSignature)
}

// ModifyAllocations is a paid mutator transaction binding the contract method 0xf8e91d16.
//
// Solidity: function modifyAllocations(address operator, (address,uint64,(address,uint32)[],uint64[])[] allocations, (bytes,bytes32,uint256) operatorSignature) returns()
func (_IAllocationManager *IAllocationManagerTransactorSession) ModifyAllocations(operator common.Address, allocations []IAllocationManagerMagnitudeAllocation, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _IAllocationManager.Contract.ModifyAllocations(&_IAllocationManager.TransactOpts, operator, allocations, operatorSignature)
}

// SetAllocationDelay is a paid mutator transaction binding the contract method 0x5c489bb5.
//
// Solidity: function setAllocationDelay(uint32 delay) returns()
func (_IAllocationManager *IAllocationManagerTransactor) SetAllocationDelay(opts *bind.TransactOpts, delay uint32) (*types.Transaction, error) {
	return _IAllocationManager.contract.Transact(opts, "setAllocationDelay", delay)
}

// SetAllocationDelay is a paid mutator transaction binding the contract method 0x5c489bb5.
//
// Solidity: function setAllocationDelay(uint32 delay) returns()
func (_IAllocationManager *IAllocationManagerSession) SetAllocationDelay(delay uint32) (*types.Transaction, error) {
	return _IAllocationManager.Contract.SetAllocationDelay(&_IAllocationManager.TransactOpts, delay)
}

// SetAllocationDelay is a paid mutator transaction binding the contract method 0x5c489bb5.
//
// Solidity: function setAllocationDelay(uint32 delay) returns()
func (_IAllocationManager *IAllocationManagerTransactorSession) SetAllocationDelay(delay uint32) (*types.Transaction, error) {
	return _IAllocationManager.Contract.SetAllocationDelay(&_IAllocationManager.TransactOpts, delay)
}

// SlashOperator is a paid mutator transaction binding the contract method 0xbd74a06c.
//
// Solidity: function slashOperator(address operator, uint32 operatorSetId, address[] strategies, uint16 bipsToSlash) returns()
func (_IAllocationManager *IAllocationManagerTransactor) SlashOperator(opts *bind.TransactOpts, operator common.Address, operatorSetId uint32, strategies []common.Address, bipsToSlash uint16) (*types.Transaction, error) {
	return _IAllocationManager.contract.Transact(opts, "slashOperator", operator, operatorSetId, strategies, bipsToSlash)
}

// SlashOperator is a paid mutator transaction binding the contract method 0xbd74a06c.
//
// Solidity: function slashOperator(address operator, uint32 operatorSetId, address[] strategies, uint16 bipsToSlash) returns()
func (_IAllocationManager *IAllocationManagerSession) SlashOperator(operator common.Address, operatorSetId uint32, strategies []common.Address, bipsToSlash uint16) (*types.Transaction, error) {
	return _IAllocationManager.Contract.SlashOperator(&_IAllocationManager.TransactOpts, operator, operatorSetId, strategies, bipsToSlash)
}

// SlashOperator is a paid mutator transaction binding the contract method 0xbd74a06c.
//
// Solidity: function slashOperator(address operator, uint32 operatorSetId, address[] strategies, uint16 bipsToSlash) returns()
func (_IAllocationManager *IAllocationManagerTransactorSession) SlashOperator(operator common.Address, operatorSetId uint32, strategies []common.Address, bipsToSlash uint16) (*types.Transaction, error) {
	return _IAllocationManager.Contract.SlashOperator(&_IAllocationManager.TransactOpts, operator, operatorSetId, strategies, bipsToSlash)
}

// UpdateFreeMagnitude is a paid mutator transaction binding the contract method 0xce770388.
//
// Solidity: function updateFreeMagnitude(address operator, address[] strategies, uint16[] numToComplete) returns()
func (_IAllocationManager *IAllocationManagerTransactor) UpdateFreeMagnitude(opts *bind.TransactOpts, operator common.Address, strategies []common.Address, numToComplete []uint16) (*types.Transaction, error) {
	return _IAllocationManager.contract.Transact(opts, "updateFreeMagnitude", operator, strategies, numToComplete)
}

// UpdateFreeMagnitude is a paid mutator transaction binding the contract method 0xce770388.
//
// Solidity: function updateFreeMagnitude(address operator, address[] strategies, uint16[] numToComplete) returns()
func (_IAllocationManager *IAllocationManagerSession) UpdateFreeMagnitude(operator common.Address, strategies []common.Address, numToComplete []uint16) (*types.Transaction, error) {
	return _IAllocationManager.Contract.UpdateFreeMagnitude(&_IAllocationManager.TransactOpts, operator, strategies, numToComplete)
}

// UpdateFreeMagnitude is a paid mutator transaction binding the contract method 0xce770388.
//
// Solidity: function updateFreeMagnitude(address operator, address[] strategies, uint16[] numToComplete) returns()
func (_IAllocationManager *IAllocationManagerTransactorSession) UpdateFreeMagnitude(operator common.Address, strategies []common.Address, numToComplete []uint16) (*types.Transaction, error) {
	return _IAllocationManager.Contract.UpdateFreeMagnitude(&_IAllocationManager.TransactOpts, operator, strategies, numToComplete)
}

// IAllocationManagerAllocationDelaySetIterator is returned from FilterAllocationDelaySet and is used to iterate over the raw logs and unpacked data for AllocationDelaySet events raised by the IAllocationManager contract.
type IAllocationManagerAllocationDelaySetIterator struct {
	Event *IAllocationManagerAllocationDelaySet // Event containing the contract specifics and raw log

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
func (it *IAllocationManagerAllocationDelaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAllocationManagerAllocationDelaySet)
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
		it.Event = new(IAllocationManagerAllocationDelaySet)
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
func (it *IAllocationManagerAllocationDelaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAllocationManagerAllocationDelaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAllocationManagerAllocationDelaySet represents a AllocationDelaySet event raised by the IAllocationManager contract.
type IAllocationManagerAllocationDelaySet struct {
	Operator common.Address
	Delay    uint32
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAllocationDelaySet is a free log retrieval operation binding the contract event 0xc17479bf29bcb9669d4dd3580ba716c0424d52c939d248d49b07efc02a32952d.
//
// Solidity: event AllocationDelaySet(address operator, uint32 delay)
func (_IAllocationManager *IAllocationManagerFilterer) FilterAllocationDelaySet(opts *bind.FilterOpts) (*IAllocationManagerAllocationDelaySetIterator, error) {

	logs, sub, err := _IAllocationManager.contract.FilterLogs(opts, "AllocationDelaySet")
	if err != nil {
		return nil, err
	}
	return &IAllocationManagerAllocationDelaySetIterator{contract: _IAllocationManager.contract, event: "AllocationDelaySet", logs: logs, sub: sub}, nil
}

// WatchAllocationDelaySet is a free log subscription operation binding the contract event 0xc17479bf29bcb9669d4dd3580ba716c0424d52c939d248d49b07efc02a32952d.
//
// Solidity: event AllocationDelaySet(address operator, uint32 delay)
func (_IAllocationManager *IAllocationManagerFilterer) WatchAllocationDelaySet(opts *bind.WatchOpts, sink chan<- *IAllocationManagerAllocationDelaySet) (event.Subscription, error) {

	logs, sub, err := _IAllocationManager.contract.WatchLogs(opts, "AllocationDelaySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAllocationManagerAllocationDelaySet)
				if err := _IAllocationManager.contract.UnpackLog(event, "AllocationDelaySet", log); err != nil {
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

// ParseAllocationDelaySet is a log parse operation binding the contract event 0xc17479bf29bcb9669d4dd3580ba716c0424d52c939d248d49b07efc02a32952d.
//
// Solidity: event AllocationDelaySet(address operator, uint32 delay)
func (_IAllocationManager *IAllocationManagerFilterer) ParseAllocationDelaySet(log types.Log) (*IAllocationManagerAllocationDelaySet, error) {
	event := new(IAllocationManagerAllocationDelaySet)
	if err := _IAllocationManager.contract.UnpackLog(event, "AllocationDelaySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAllocationManagerMagnitudeAllocatedIterator is returned from FilterMagnitudeAllocated and is used to iterate over the raw logs and unpacked data for MagnitudeAllocated events raised by the IAllocationManager contract.
type IAllocationManagerMagnitudeAllocatedIterator struct {
	Event *IAllocationManagerMagnitudeAllocated // Event containing the contract specifics and raw log

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
func (it *IAllocationManagerMagnitudeAllocatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAllocationManagerMagnitudeAllocated)
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
		it.Event = new(IAllocationManagerMagnitudeAllocated)
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
func (it *IAllocationManagerMagnitudeAllocatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAllocationManagerMagnitudeAllocatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAllocationManagerMagnitudeAllocated represents a MagnitudeAllocated event raised by the IAllocationManager contract.
type IAllocationManagerMagnitudeAllocated struct {
	Operator            common.Address
	Strategy            common.Address
	OperatorSet         OperatorSet
	MagnitudeToAllocate uint64
	EffectTimestamp     uint32
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterMagnitudeAllocated is a free log retrieval operation binding the contract event 0x6d7d0079582cb2c5e70d4135b37f36711415ee6c260778b716bd65e026eb4f1a.
//
// Solidity: event MagnitudeAllocated(address operator, address strategy, (address,uint32) operatorSet, uint64 magnitudeToAllocate, uint32 effectTimestamp)
func (_IAllocationManager *IAllocationManagerFilterer) FilterMagnitudeAllocated(opts *bind.FilterOpts) (*IAllocationManagerMagnitudeAllocatedIterator, error) {

	logs, sub, err := _IAllocationManager.contract.FilterLogs(opts, "MagnitudeAllocated")
	if err != nil {
		return nil, err
	}
	return &IAllocationManagerMagnitudeAllocatedIterator{contract: _IAllocationManager.contract, event: "MagnitudeAllocated", logs: logs, sub: sub}, nil
}

// WatchMagnitudeAllocated is a free log subscription operation binding the contract event 0x6d7d0079582cb2c5e70d4135b37f36711415ee6c260778b716bd65e026eb4f1a.
//
// Solidity: event MagnitudeAllocated(address operator, address strategy, (address,uint32) operatorSet, uint64 magnitudeToAllocate, uint32 effectTimestamp)
func (_IAllocationManager *IAllocationManagerFilterer) WatchMagnitudeAllocated(opts *bind.WatchOpts, sink chan<- *IAllocationManagerMagnitudeAllocated) (event.Subscription, error) {

	logs, sub, err := _IAllocationManager.contract.WatchLogs(opts, "MagnitudeAllocated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAllocationManagerMagnitudeAllocated)
				if err := _IAllocationManager.contract.UnpackLog(event, "MagnitudeAllocated", log); err != nil {
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

// ParseMagnitudeAllocated is a log parse operation binding the contract event 0x6d7d0079582cb2c5e70d4135b37f36711415ee6c260778b716bd65e026eb4f1a.
//
// Solidity: event MagnitudeAllocated(address operator, address strategy, (address,uint32) operatorSet, uint64 magnitudeToAllocate, uint32 effectTimestamp)
func (_IAllocationManager *IAllocationManagerFilterer) ParseMagnitudeAllocated(log types.Log) (*IAllocationManagerMagnitudeAllocated, error) {
	event := new(IAllocationManagerMagnitudeAllocated)
	if err := _IAllocationManager.contract.UnpackLog(event, "MagnitudeAllocated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAllocationManagerMagnitudeDeallocationCompletedIterator is returned from FilterMagnitudeDeallocationCompleted and is used to iterate over the raw logs and unpacked data for MagnitudeDeallocationCompleted events raised by the IAllocationManager contract.
type IAllocationManagerMagnitudeDeallocationCompletedIterator struct {
	Event *IAllocationManagerMagnitudeDeallocationCompleted // Event containing the contract specifics and raw log

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
func (it *IAllocationManagerMagnitudeDeallocationCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAllocationManagerMagnitudeDeallocationCompleted)
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
		it.Event = new(IAllocationManagerMagnitudeDeallocationCompleted)
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
func (it *IAllocationManagerMagnitudeDeallocationCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAllocationManagerMagnitudeDeallocationCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAllocationManagerMagnitudeDeallocationCompleted represents a MagnitudeDeallocationCompleted event raised by the IAllocationManager contract.
type IAllocationManagerMagnitudeDeallocationCompleted struct {
	Operator           common.Address
	Strategy           common.Address
	OperatorSet        OperatorSet
	FreeMagnitudeAdded uint64
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMagnitudeDeallocationCompleted is a free log retrieval operation binding the contract event 0x1e5c8e13c62c31d6252ac205e592477d643c7e95831d5b46d99a3c60c2fad8db.
//
// Solidity: event MagnitudeDeallocationCompleted(address operator, address strategy, (address,uint32) operatorSet, uint64 freeMagnitudeAdded)
func (_IAllocationManager *IAllocationManagerFilterer) FilterMagnitudeDeallocationCompleted(opts *bind.FilterOpts) (*IAllocationManagerMagnitudeDeallocationCompletedIterator, error) {

	logs, sub, err := _IAllocationManager.contract.FilterLogs(opts, "MagnitudeDeallocationCompleted")
	if err != nil {
		return nil, err
	}
	return &IAllocationManagerMagnitudeDeallocationCompletedIterator{contract: _IAllocationManager.contract, event: "MagnitudeDeallocationCompleted", logs: logs, sub: sub}, nil
}

// WatchMagnitudeDeallocationCompleted is a free log subscription operation binding the contract event 0x1e5c8e13c62c31d6252ac205e592477d643c7e95831d5b46d99a3c60c2fad8db.
//
// Solidity: event MagnitudeDeallocationCompleted(address operator, address strategy, (address,uint32) operatorSet, uint64 freeMagnitudeAdded)
func (_IAllocationManager *IAllocationManagerFilterer) WatchMagnitudeDeallocationCompleted(opts *bind.WatchOpts, sink chan<- *IAllocationManagerMagnitudeDeallocationCompleted) (event.Subscription, error) {

	logs, sub, err := _IAllocationManager.contract.WatchLogs(opts, "MagnitudeDeallocationCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAllocationManagerMagnitudeDeallocationCompleted)
				if err := _IAllocationManager.contract.UnpackLog(event, "MagnitudeDeallocationCompleted", log); err != nil {
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

// ParseMagnitudeDeallocationCompleted is a log parse operation binding the contract event 0x1e5c8e13c62c31d6252ac205e592477d643c7e95831d5b46d99a3c60c2fad8db.
//
// Solidity: event MagnitudeDeallocationCompleted(address operator, address strategy, (address,uint32) operatorSet, uint64 freeMagnitudeAdded)
func (_IAllocationManager *IAllocationManagerFilterer) ParseMagnitudeDeallocationCompleted(log types.Log) (*IAllocationManagerMagnitudeDeallocationCompleted, error) {
	event := new(IAllocationManagerMagnitudeDeallocationCompleted)
	if err := _IAllocationManager.contract.UnpackLog(event, "MagnitudeDeallocationCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAllocationManagerMagnitudeQueueDeallocatedIterator is returned from FilterMagnitudeQueueDeallocated and is used to iterate over the raw logs and unpacked data for MagnitudeQueueDeallocated events raised by the IAllocationManager contract.
type IAllocationManagerMagnitudeQueueDeallocatedIterator struct {
	Event *IAllocationManagerMagnitudeQueueDeallocated // Event containing the contract specifics and raw log

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
func (it *IAllocationManagerMagnitudeQueueDeallocatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAllocationManagerMagnitudeQueueDeallocated)
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
		it.Event = new(IAllocationManagerMagnitudeQueueDeallocated)
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
func (it *IAllocationManagerMagnitudeQueueDeallocatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAllocationManagerMagnitudeQueueDeallocatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAllocationManagerMagnitudeQueueDeallocated represents a MagnitudeQueueDeallocated event raised by the IAllocationManager contract.
type IAllocationManagerMagnitudeQueueDeallocated struct {
	Operator              common.Address
	Strategy              common.Address
	OperatorSet           OperatorSet
	MagnitudeToDeallocate uint64
	CompletableTimestamp  uint32
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterMagnitudeQueueDeallocated is a free log retrieval operation binding the contract event 0x2e68db1fe51107d7e451ae268d1631796989ab9d7925054e9b247854cb5be950.
//
// Solidity: event MagnitudeQueueDeallocated(address operator, address strategy, (address,uint32) operatorSet, uint64 magnitudeToDeallocate, uint32 completableTimestamp)
func (_IAllocationManager *IAllocationManagerFilterer) FilterMagnitudeQueueDeallocated(opts *bind.FilterOpts) (*IAllocationManagerMagnitudeQueueDeallocatedIterator, error) {

	logs, sub, err := _IAllocationManager.contract.FilterLogs(opts, "MagnitudeQueueDeallocated")
	if err != nil {
		return nil, err
	}
	return &IAllocationManagerMagnitudeQueueDeallocatedIterator{contract: _IAllocationManager.contract, event: "MagnitudeQueueDeallocated", logs: logs, sub: sub}, nil
}

// WatchMagnitudeQueueDeallocated is a free log subscription operation binding the contract event 0x2e68db1fe51107d7e451ae268d1631796989ab9d7925054e9b247854cb5be950.
//
// Solidity: event MagnitudeQueueDeallocated(address operator, address strategy, (address,uint32) operatorSet, uint64 magnitudeToDeallocate, uint32 completableTimestamp)
func (_IAllocationManager *IAllocationManagerFilterer) WatchMagnitudeQueueDeallocated(opts *bind.WatchOpts, sink chan<- *IAllocationManagerMagnitudeQueueDeallocated) (event.Subscription, error) {

	logs, sub, err := _IAllocationManager.contract.WatchLogs(opts, "MagnitudeQueueDeallocated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAllocationManagerMagnitudeQueueDeallocated)
				if err := _IAllocationManager.contract.UnpackLog(event, "MagnitudeQueueDeallocated", log); err != nil {
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

// ParseMagnitudeQueueDeallocated is a log parse operation binding the contract event 0x2e68db1fe51107d7e451ae268d1631796989ab9d7925054e9b247854cb5be950.
//
// Solidity: event MagnitudeQueueDeallocated(address operator, address strategy, (address,uint32) operatorSet, uint64 magnitudeToDeallocate, uint32 completableTimestamp)
func (_IAllocationManager *IAllocationManagerFilterer) ParseMagnitudeQueueDeallocated(log types.Log) (*IAllocationManagerMagnitudeQueueDeallocated, error) {
	event := new(IAllocationManagerMagnitudeQueueDeallocated)
	if err := _IAllocationManager.contract.UnpackLog(event, "MagnitudeQueueDeallocated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAllocationManagerOperatorSetCreatedIterator is returned from FilterOperatorSetCreated and is used to iterate over the raw logs and unpacked data for OperatorSetCreated events raised by the IAllocationManager contract.
type IAllocationManagerOperatorSetCreatedIterator struct {
	Event *IAllocationManagerOperatorSetCreated // Event containing the contract specifics and raw log

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
func (it *IAllocationManagerOperatorSetCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAllocationManagerOperatorSetCreated)
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
		it.Event = new(IAllocationManagerOperatorSetCreated)
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
func (it *IAllocationManagerOperatorSetCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAllocationManagerOperatorSetCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAllocationManagerOperatorSetCreated represents a OperatorSetCreated event raised by the IAllocationManager contract.
type IAllocationManagerOperatorSetCreated struct {
	OperatorSet OperatorSet
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorSetCreated is a free log retrieval operation binding the contract event 0x31629285ead2335ae0933f86ed2ae63321f7af77b4e6eaabc42c057880977e6c.
//
// Solidity: event OperatorSetCreated((address,uint32) operatorSet)
func (_IAllocationManager *IAllocationManagerFilterer) FilterOperatorSetCreated(opts *bind.FilterOpts) (*IAllocationManagerOperatorSetCreatedIterator, error) {

	logs, sub, err := _IAllocationManager.contract.FilterLogs(opts, "OperatorSetCreated")
	if err != nil {
		return nil, err
	}
	return &IAllocationManagerOperatorSetCreatedIterator{contract: _IAllocationManager.contract, event: "OperatorSetCreated", logs: logs, sub: sub}, nil
}

// WatchOperatorSetCreated is a free log subscription operation binding the contract event 0x31629285ead2335ae0933f86ed2ae63321f7af77b4e6eaabc42c057880977e6c.
//
// Solidity: event OperatorSetCreated((address,uint32) operatorSet)
func (_IAllocationManager *IAllocationManagerFilterer) WatchOperatorSetCreated(opts *bind.WatchOpts, sink chan<- *IAllocationManagerOperatorSetCreated) (event.Subscription, error) {

	logs, sub, err := _IAllocationManager.contract.WatchLogs(opts, "OperatorSetCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAllocationManagerOperatorSetCreated)
				if err := _IAllocationManager.contract.UnpackLog(event, "OperatorSetCreated", log); err != nil {
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

// ParseOperatorSetCreated is a log parse operation binding the contract event 0x31629285ead2335ae0933f86ed2ae63321f7af77b4e6eaabc42c057880977e6c.
//
// Solidity: event OperatorSetCreated((address,uint32) operatorSet)
func (_IAllocationManager *IAllocationManagerFilterer) ParseOperatorSetCreated(log types.Log) (*IAllocationManagerOperatorSetCreated, error) {
	event := new(IAllocationManagerOperatorSetCreated)
	if err := _IAllocationManager.contract.UnpackLog(event, "OperatorSetCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAllocationManagerOperatorSlashedIterator is returned from FilterOperatorSlashed and is used to iterate over the raw logs and unpacked data for OperatorSlashed events raised by the IAllocationManager contract.
type IAllocationManagerOperatorSlashedIterator struct {
	Event *IAllocationManagerOperatorSlashed // Event containing the contract specifics and raw log

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
func (it *IAllocationManagerOperatorSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAllocationManagerOperatorSlashed)
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
		it.Event = new(IAllocationManagerOperatorSlashed)
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
func (it *IAllocationManagerOperatorSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAllocationManagerOperatorSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAllocationManagerOperatorSlashed represents a OperatorSlashed event raised by the IAllocationManager contract.
type IAllocationManagerOperatorSlashed struct {
	Operator      common.Address
	OperatorSetId uint32
	Strategy      common.Address
	BipsToSlash   uint16
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOperatorSlashed is a free log retrieval operation binding the contract event 0xe672839d3c371691acdb52de9fefc94b3dbf407dc0920ef566c7c059ad575b1c.
//
// Solidity: event OperatorSlashed(address operator, uint32 operatorSetId, address strategy, uint16 bipsToSlash)
func (_IAllocationManager *IAllocationManagerFilterer) FilterOperatorSlashed(opts *bind.FilterOpts) (*IAllocationManagerOperatorSlashedIterator, error) {

	logs, sub, err := _IAllocationManager.contract.FilterLogs(opts, "OperatorSlashed")
	if err != nil {
		return nil, err
	}
	return &IAllocationManagerOperatorSlashedIterator{contract: _IAllocationManager.contract, event: "OperatorSlashed", logs: logs, sub: sub}, nil
}

// WatchOperatorSlashed is a free log subscription operation binding the contract event 0xe672839d3c371691acdb52de9fefc94b3dbf407dc0920ef566c7c059ad575b1c.
//
// Solidity: event OperatorSlashed(address operator, uint32 operatorSetId, address strategy, uint16 bipsToSlash)
func (_IAllocationManager *IAllocationManagerFilterer) WatchOperatorSlashed(opts *bind.WatchOpts, sink chan<- *IAllocationManagerOperatorSlashed) (event.Subscription, error) {

	logs, sub, err := _IAllocationManager.contract.WatchLogs(opts, "OperatorSlashed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAllocationManagerOperatorSlashed)
				if err := _IAllocationManager.contract.UnpackLog(event, "OperatorSlashed", log); err != nil {
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

// ParseOperatorSlashed is a log parse operation binding the contract event 0xe672839d3c371691acdb52de9fefc94b3dbf407dc0920ef566c7c059ad575b1c.
//
// Solidity: event OperatorSlashed(address operator, uint32 operatorSetId, address strategy, uint16 bipsToSlash)
func (_IAllocationManager *IAllocationManagerFilterer) ParseOperatorSlashed(log types.Log) (*IAllocationManagerOperatorSlashed, error) {
	event := new(IAllocationManagerOperatorSlashed)
	if err := _IAllocationManager.contract.UnpackLog(event, "OperatorSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
