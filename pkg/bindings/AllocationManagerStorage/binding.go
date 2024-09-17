// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package AllocationManagerStorage

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

// AllocationManagerStorageMetaData contains all meta data concerning the AllocationManagerStorage contract.
var AllocationManagerStorageMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"ALLOCATION_DELAY_CONFIGURATION_DELAY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEALLOCATION_DELAY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DOMAIN_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAGNITUDE_ADJUSTMENT_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allocationDelay\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"isSet\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"delay\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"avsDirectory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAVSDirectory\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateMagnitudeAllocationDigestHash\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allocations\",\"type\":\"tuple[]\",\"internalType\":\"structIAllocationManager.MagnitudeAllocation[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"expectedTotalMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"operatorSets\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"magnitudes\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}]},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cancelSalt\",\"inputs\":[{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"domainSeparator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocatableMagnitude\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"numToComplete\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentSlashableMagnitudes\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"\",\"type\":\"uint64[][]\",\"internalType\":\"uint64[][]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPendingAllocations\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"operatorSets\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPendingDeallocations\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"operatorSets\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSlashableMagnitudes\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"timestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"\",\"type\":\"uint64[][]\",\"internalType\":\"uint64[][]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSlashablePPM\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"timestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"linear\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint24[]\",\"internalType\":\"uint24[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalAndAllocatedMagnitudes\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalMagnitude\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalMagnitudeAtTimestamp\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"timestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalMagnitudes\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalMagnitudesAtTimestamp\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"timestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperatorSlashable\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"modifyAllocations\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allocations\",\"type\":\"tuple[]\",\"internalType\":\"structIAllocationManager.MagnitudeAllocation[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"expectedTotalMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"operatorSets\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"magnitudes\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}]},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"operatorMagnitudeInfo\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"nextPendingFreeMagnitudeIndex\",\"type\":\"uint192\",\"internalType\":\"uint192\"},{\"name\":\"freeMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorSaltIsSpent\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setAllocationDelay\",\"inputs\":[{\"name\":\"delay\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slashOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"bipsToSlash\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateFreeMagnitude\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"numToComplete\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllocationDelaySet\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"delay\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MagnitudeAllocated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"magnitudeToAllocate\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"effectTimestamp\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MagnitudeDeallocationCompleted\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"freeMagnitudeAdded\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MagnitudeQueueDeallocated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"magnitudeToDeallocate\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"completableTimestamp\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSetCreated\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSlashed\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"bipsToSlash\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InputArrayLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientAllocatableMagnitude\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidExpectedTotalMagnitude\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperatorSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorNotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorSetsNotInAscendingOrder\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PendingAllocationOrDeallocation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SaltSpent\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignatureExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UninitializedAllocationDelay\",\"inputs\":[]}]",
}

// AllocationManagerStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use AllocationManagerStorageMetaData.ABI instead.
var AllocationManagerStorageABI = AllocationManagerStorageMetaData.ABI

// AllocationManagerStorage is an auto generated Go binding around an Ethereum contract.
type AllocationManagerStorage struct {
	AllocationManagerStorageCaller     // Read-only binding to the contract
	AllocationManagerStorageTransactor // Write-only binding to the contract
	AllocationManagerStorageFilterer   // Log filterer for contract events
}

// AllocationManagerStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type AllocationManagerStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AllocationManagerStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AllocationManagerStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AllocationManagerStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AllocationManagerStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AllocationManagerStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AllocationManagerStorageSession struct {
	Contract     *AllocationManagerStorage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// AllocationManagerStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AllocationManagerStorageCallerSession struct {
	Contract *AllocationManagerStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// AllocationManagerStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AllocationManagerStorageTransactorSession struct {
	Contract     *AllocationManagerStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// AllocationManagerStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type AllocationManagerStorageRaw struct {
	Contract *AllocationManagerStorage // Generic contract binding to access the raw methods on
}

// AllocationManagerStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AllocationManagerStorageCallerRaw struct {
	Contract *AllocationManagerStorageCaller // Generic read-only contract binding to access the raw methods on
}

// AllocationManagerStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AllocationManagerStorageTransactorRaw struct {
	Contract *AllocationManagerStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAllocationManagerStorage creates a new instance of AllocationManagerStorage, bound to a specific deployed contract.
func NewAllocationManagerStorage(address common.Address, backend bind.ContractBackend) (*AllocationManagerStorage, error) {
	contract, err := bindAllocationManagerStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AllocationManagerStorage{AllocationManagerStorageCaller: AllocationManagerStorageCaller{contract: contract}, AllocationManagerStorageTransactor: AllocationManagerStorageTransactor{contract: contract}, AllocationManagerStorageFilterer: AllocationManagerStorageFilterer{contract: contract}}, nil
}

// NewAllocationManagerStorageCaller creates a new read-only instance of AllocationManagerStorage, bound to a specific deployed contract.
func NewAllocationManagerStorageCaller(address common.Address, caller bind.ContractCaller) (*AllocationManagerStorageCaller, error) {
	contract, err := bindAllocationManagerStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AllocationManagerStorageCaller{contract: contract}, nil
}

// NewAllocationManagerStorageTransactor creates a new write-only instance of AllocationManagerStorage, bound to a specific deployed contract.
func NewAllocationManagerStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*AllocationManagerStorageTransactor, error) {
	contract, err := bindAllocationManagerStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AllocationManagerStorageTransactor{contract: contract}, nil
}

// NewAllocationManagerStorageFilterer creates a new log filterer instance of AllocationManagerStorage, bound to a specific deployed contract.
func NewAllocationManagerStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*AllocationManagerStorageFilterer, error) {
	contract, err := bindAllocationManagerStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AllocationManagerStorageFilterer{contract: contract}, nil
}

// bindAllocationManagerStorage binds a generic wrapper to an already deployed contract.
func bindAllocationManagerStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AllocationManagerStorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AllocationManagerStorage *AllocationManagerStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AllocationManagerStorage.Contract.AllocationManagerStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AllocationManagerStorage *AllocationManagerStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AllocationManagerStorage.Contract.AllocationManagerStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AllocationManagerStorage *AllocationManagerStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AllocationManagerStorage.Contract.AllocationManagerStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AllocationManagerStorage *AllocationManagerStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AllocationManagerStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AllocationManagerStorage *AllocationManagerStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AllocationManagerStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AllocationManagerStorage *AllocationManagerStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AllocationManagerStorage.Contract.contract.Transact(opts, method, params...)
}

// ALLOCATIONDELAYCONFIGURATIONDELAY is a free data retrieval call binding the contract method 0x757c4ff2.
//
// Solidity: function ALLOCATION_DELAY_CONFIGURATION_DELAY() view returns(uint32)
func (_AllocationManagerStorage *AllocationManagerStorageCaller) ALLOCATIONDELAYCONFIGURATIONDELAY(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _AllocationManagerStorage.contract.Call(opts, &out, "ALLOCATION_DELAY_CONFIGURATION_DELAY")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ALLOCATIONDELAYCONFIGURATIONDELAY is a free data retrieval call binding the contract method 0x757c4ff2.
//
// Solidity: function ALLOCATION_DELAY_CONFIGURATION_DELAY() view returns(uint32)
func (_AllocationManagerStorage *AllocationManagerStorageSession) ALLOCATIONDELAYCONFIGURATIONDELAY() (uint32, error) {
	return _AllocationManagerStorage.Contract.ALLOCATIONDELAYCONFIGURATIONDELAY(&_AllocationManagerStorage.CallOpts)
}

// ALLOCATIONDELAYCONFIGURATIONDELAY is a free data retrieval call binding the contract method 0x757c4ff2.
//
// Solidity: function ALLOCATION_DELAY_CONFIGURATION_DELAY() view returns(uint32)
func (_AllocationManagerStorage *AllocationManagerStorageCallerSession) ALLOCATIONDELAYCONFIGURATIONDELAY() (uint32, error) {
	return _AllocationManagerStorage.Contract.ALLOCATIONDELAYCONFIGURATIONDELAY(&_AllocationManagerStorage.CallOpts)
}

// DEALLOCATIONDELAY is a free data retrieval call binding the contract method 0x2981eb77.
//
// Solidity: function DEALLOCATION_DELAY() view returns(uint32)
func (_AllocationManagerStorage *AllocationManagerStorageCaller) DEALLOCATIONDELAY(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _AllocationManagerStorage.contract.Call(opts, &out, "DEALLOCATION_DELAY")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// DEALLOCATIONDELAY is a free data retrieval call binding the contract method 0x2981eb77.
//
// Solidity: function DEALLOCATION_DELAY() view returns(uint32)
func (_AllocationManagerStorage *AllocationManagerStorageSession) DEALLOCATIONDELAY() (uint32, error) {
	return _AllocationManagerStorage.Contract.DEALLOCATIONDELAY(&_AllocationManagerStorage.CallOpts)
}

// DEALLOCATIONDELAY is a free data retrieval call binding the contract method 0x2981eb77.
//
// Solidity: function DEALLOCATION_DELAY() view returns(uint32)
func (_AllocationManagerStorage *AllocationManagerStorageCallerSession) DEALLOCATIONDELAY() (uint32, error) {
	return _AllocationManagerStorage.Contract.DEALLOCATIONDELAY(&_AllocationManagerStorage.CallOpts)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_AllocationManagerStorage *AllocationManagerStorageCaller) DOMAINTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AllocationManagerStorage.contract.Call(opts, &out, "DOMAIN_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_AllocationManagerStorage *AllocationManagerStorageSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _AllocationManagerStorage.Contract.DOMAINTYPEHASH(&_AllocationManagerStorage.CallOpts)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_AllocationManagerStorage *AllocationManagerStorageCallerSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _AllocationManagerStorage.Contract.DOMAINTYPEHASH(&_AllocationManagerStorage.CallOpts)
}

// MAGNITUDEADJUSTMENTTYPEHASH is a free data retrieval call binding the contract method 0x7b205de3.
//
// Solidity: function MAGNITUDE_ADJUSTMENT_TYPEHASH() view returns(bytes32)
func (_AllocationManagerStorage *AllocationManagerStorageCaller) MAGNITUDEADJUSTMENTTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AllocationManagerStorage.contract.Call(opts, &out, "MAGNITUDE_ADJUSTMENT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MAGNITUDEADJUSTMENTTYPEHASH is a free data retrieval call binding the contract method 0x7b205de3.
//
// Solidity: function MAGNITUDE_ADJUSTMENT_TYPEHASH() view returns(bytes32)
func (_AllocationManagerStorage *AllocationManagerStorageSession) MAGNITUDEADJUSTMENTTYPEHASH() ([32]byte, error) {
	return _AllocationManagerStorage.Contract.MAGNITUDEADJUSTMENTTYPEHASH(&_AllocationManagerStorage.CallOpts)
}

// MAGNITUDEADJUSTMENTTYPEHASH is a free data retrieval call binding the contract method 0x7b205de3.
//
// Solidity: function MAGNITUDE_ADJUSTMENT_TYPEHASH() view returns(bytes32)
func (_AllocationManagerStorage *AllocationManagerStorageCallerSession) MAGNITUDEADJUSTMENTTYPEHASH() ([32]byte, error) {
	return _AllocationManagerStorage.Contract.MAGNITUDEADJUSTMENTTYPEHASH(&_AllocationManagerStorage.CallOpts)
}

// AllocationDelay is a free data retrieval call binding the contract method 0xbe4e1fd3.
//
// Solidity: function allocationDelay(address operator) view returns(bool isSet, uint32 delay)
func (_AllocationManagerStorage *AllocationManagerStorageCaller) AllocationDelay(opts *bind.CallOpts, operator common.Address) (struct {
	IsSet bool
	Delay uint32
}, error) {
	var out []interface{}
	err := _AllocationManagerStorage.contract.Call(opts, &out, "allocationDelay", operator)

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
func (_AllocationManagerStorage *AllocationManagerStorageSession) AllocationDelay(operator common.Address) (struct {
	IsSet bool
	Delay uint32
}, error) {
	return _AllocationManagerStorage.Contract.AllocationDelay(&_AllocationManagerStorage.CallOpts, operator)
}

// AllocationDelay is a free data retrieval call binding the contract method 0xbe4e1fd3.
//
// Solidity: function allocationDelay(address operator) view returns(bool isSet, uint32 delay)
func (_AllocationManagerStorage *AllocationManagerStorageCallerSession) AllocationDelay(operator common.Address) (struct {
	IsSet bool
	Delay uint32
}, error) {
	return _AllocationManagerStorage.Contract.AllocationDelay(&_AllocationManagerStorage.CallOpts, operator)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_AllocationManagerStorage *AllocationManagerStorageCaller) AvsDirectory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AllocationManagerStorage.contract.Call(opts, &out, "avsDirectory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_AllocationManagerStorage *AllocationManagerStorageSession) AvsDirectory() (common.Address, error) {
	return _AllocationManagerStorage.Contract.AvsDirectory(&_AllocationManagerStorage.CallOpts)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_AllocationManagerStorage *AllocationManagerStorageCallerSession) AvsDirectory() (common.Address, error) {
	return _AllocationManagerStorage.Contract.AvsDirectory(&_AllocationManagerStorage.CallOpts)
}

// CalculateMagnitudeAllocationDigestHash is a free data retrieval call binding the contract method 0x686b686e.
//
// Solidity: function calculateMagnitudeAllocationDigestHash(address operator, (address,uint64,(address,uint32)[],uint64[])[] allocations, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_AllocationManagerStorage *AllocationManagerStorageCaller) CalculateMagnitudeAllocationDigestHash(opts *bind.CallOpts, operator common.Address, allocations []IAllocationManagerMagnitudeAllocation, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _AllocationManagerStorage.contract.Call(opts, &out, "calculateMagnitudeAllocationDigestHash", operator, allocations, salt, expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateMagnitudeAllocationDigestHash is a free data retrieval call binding the contract method 0x686b686e.
//
// Solidity: function calculateMagnitudeAllocationDigestHash(address operator, (address,uint64,(address,uint32)[],uint64[])[] allocations, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_AllocationManagerStorage *AllocationManagerStorageSession) CalculateMagnitudeAllocationDigestHash(operator common.Address, allocations []IAllocationManagerMagnitudeAllocation, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _AllocationManagerStorage.Contract.CalculateMagnitudeAllocationDigestHash(&_AllocationManagerStorage.CallOpts, operator, allocations, salt, expiry)
}

// CalculateMagnitudeAllocationDigestHash is a free data retrieval call binding the contract method 0x686b686e.
//
// Solidity: function calculateMagnitudeAllocationDigestHash(address operator, (address,uint64,(address,uint32)[],uint64[])[] allocations, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_AllocationManagerStorage *AllocationManagerStorageCallerSession) CalculateMagnitudeAllocationDigestHash(operator common.Address, allocations []IAllocationManagerMagnitudeAllocation, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _AllocationManagerStorage.Contract.CalculateMagnitudeAllocationDigestHash(&_AllocationManagerStorage.CallOpts, operator, allocations, salt, expiry)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_AllocationManagerStorage *AllocationManagerStorageCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AllocationManagerStorage.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_AllocationManagerStorage *AllocationManagerStorageSession) Delegation() (common.Address, error) {
	return _AllocationManagerStorage.Contract.Delegation(&_AllocationManagerStorage.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_AllocationManagerStorage *AllocationManagerStorageCallerSession) Delegation() (common.Address, error) {
	return _AllocationManagerStorage.Contract.Delegation(&_AllocationManagerStorage.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_AllocationManagerStorage *AllocationManagerStorageCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AllocationManagerStorage.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_AllocationManagerStorage *AllocationManagerStorageSession) DomainSeparator() ([32]byte, error) {
	return _AllocationManagerStorage.Contract.DomainSeparator(&_AllocationManagerStorage.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_AllocationManagerStorage *AllocationManagerStorageCallerSession) DomainSeparator() ([32]byte, error) {
	return _AllocationManagerStorage.Contract.DomainSeparator(&_AllocationManagerStorage.CallOpts)
}

// GetAllocatableMagnitude is a free data retrieval call binding the contract method 0x3f78612f.
//
// Solidity: function getAllocatableMagnitude(address operator, address strategy, uint16 numToComplete) view returns(uint64)
func (_AllocationManagerStorage *AllocationManagerStorageCaller) GetAllocatableMagnitude(opts *bind.CallOpts, operator common.Address, strategy common.Address, numToComplete uint16) (uint64, error) {
	var out []interface{}
	err := _AllocationManagerStorage.contract.Call(opts, &out, "getAllocatableMagnitude", operator, strategy, numToComplete)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetAllocatableMagnitude is a free data retrieval call binding the contract method 0x3f78612f.
//
// Solidity: function getAllocatableMagnitude(address operator, address strategy, uint16 numToComplete) view returns(uint64)
func (_AllocationManagerStorage *AllocationManagerStorageSession) GetAllocatableMagnitude(operator common.Address, strategy common.Address, numToComplete uint16) (uint64, error) {
	return _AllocationManagerStorage.Contract.GetAllocatableMagnitude(&_AllocationManagerStorage.CallOpts, operator, strategy, numToComplete)
}

// GetAllocatableMagnitude is a free data retrieval call binding the contract method 0x3f78612f.
//
// Solidity: function getAllocatableMagnitude(address operator, address strategy, uint16 numToComplete) view returns(uint64)
func (_AllocationManagerStorage *AllocationManagerStorageCallerSession) GetAllocatableMagnitude(operator common.Address, strategy common.Address, numToComplete uint16) (uint64, error) {
	return _AllocationManagerStorage.Contract.GetAllocatableMagnitude(&_AllocationManagerStorage.CallOpts, operator, strategy, numToComplete)
}

// GetCurrentSlashableMagnitudes is a free data retrieval call binding the contract method 0x1d95b524.
//
// Solidity: function getCurrentSlashableMagnitudes(address operator, address[] strategies) view returns((address,uint32)[], uint64[][])
func (_AllocationManagerStorage *AllocationManagerStorageCaller) GetCurrentSlashableMagnitudes(opts *bind.CallOpts, operator common.Address, strategies []common.Address) ([]OperatorSet, [][]uint64, error) {
	var out []interface{}
	err := _AllocationManagerStorage.contract.Call(opts, &out, "getCurrentSlashableMagnitudes", operator, strategies)

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
func (_AllocationManagerStorage *AllocationManagerStorageSession) GetCurrentSlashableMagnitudes(operator common.Address, strategies []common.Address) ([]OperatorSet, [][]uint64, error) {
	return _AllocationManagerStorage.Contract.GetCurrentSlashableMagnitudes(&_AllocationManagerStorage.CallOpts, operator, strategies)
}

// GetCurrentSlashableMagnitudes is a free data retrieval call binding the contract method 0x1d95b524.
//
// Solidity: function getCurrentSlashableMagnitudes(address operator, address[] strategies) view returns((address,uint32)[], uint64[][])
func (_AllocationManagerStorage *AllocationManagerStorageCallerSession) GetCurrentSlashableMagnitudes(operator common.Address, strategies []common.Address) ([]OperatorSet, [][]uint64, error) {
	return _AllocationManagerStorage.Contract.GetCurrentSlashableMagnitudes(&_AllocationManagerStorage.CallOpts, operator, strategies)
}

// GetPendingAllocations is a free data retrieval call binding the contract method 0x67d973ef.
//
// Solidity: function getPendingAllocations(address operator, address strategy, (address,uint32)[] operatorSets) view returns(uint64[], uint32[])
func (_AllocationManagerStorage *AllocationManagerStorageCaller) GetPendingAllocations(opts *bind.CallOpts, operator common.Address, strategy common.Address, operatorSets []OperatorSet) ([]uint64, []uint32, error) {
	var out []interface{}
	err := _AllocationManagerStorage.contract.Call(opts, &out, "getPendingAllocations", operator, strategy, operatorSets)

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
func (_AllocationManagerStorage *AllocationManagerStorageSession) GetPendingAllocations(operator common.Address, strategy common.Address, operatorSets []OperatorSet) ([]uint64, []uint32, error) {
	return _AllocationManagerStorage.Contract.GetPendingAllocations(&_AllocationManagerStorage.CallOpts, operator, strategy, operatorSets)
}

// GetPendingAllocations is a free data retrieval call binding the contract method 0x67d973ef.
//
// Solidity: function getPendingAllocations(address operator, address strategy, (address,uint32)[] operatorSets) view returns(uint64[], uint32[])
func (_AllocationManagerStorage *AllocationManagerStorageCallerSession) GetPendingAllocations(operator common.Address, strategy common.Address, operatorSets []OperatorSet) ([]uint64, []uint32, error) {
	return _AllocationManagerStorage.Contract.GetPendingAllocations(&_AllocationManagerStorage.CallOpts, operator, strategy, operatorSets)
}

// GetPendingDeallocations is a free data retrieval call binding the contract method 0x07053717.
//
// Solidity: function getPendingDeallocations(address operator, address strategy, (address,uint32)[] operatorSets) view returns(uint64[], uint32[])
func (_AllocationManagerStorage *AllocationManagerStorageCaller) GetPendingDeallocations(opts *bind.CallOpts, operator common.Address, strategy common.Address, operatorSets []OperatorSet) ([]uint64, []uint32, error) {
	var out []interface{}
	err := _AllocationManagerStorage.contract.Call(opts, &out, "getPendingDeallocations", operator, strategy, operatorSets)

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
func (_AllocationManagerStorage *AllocationManagerStorageSession) GetPendingDeallocations(operator common.Address, strategy common.Address, operatorSets []OperatorSet) ([]uint64, []uint32, error) {
	return _AllocationManagerStorage.Contract.GetPendingDeallocations(&_AllocationManagerStorage.CallOpts, operator, strategy, operatorSets)
}

// GetPendingDeallocations is a free data retrieval call binding the contract method 0x07053717.
//
// Solidity: function getPendingDeallocations(address operator, address strategy, (address,uint32)[] operatorSets) view returns(uint64[], uint32[])
func (_AllocationManagerStorage *AllocationManagerStorageCallerSession) GetPendingDeallocations(operator common.Address, strategy common.Address, operatorSets []OperatorSet) ([]uint64, []uint32, error) {
	return _AllocationManagerStorage.Contract.GetPendingDeallocations(&_AllocationManagerStorage.CallOpts, operator, strategy, operatorSets)
}

// GetSlashableMagnitudes is a free data retrieval call binding the contract method 0x43b0592d.
//
// Solidity: function getSlashableMagnitudes(address operator, address[] strategies, uint32 timestamp) view returns((address,uint32)[], uint64[][])
func (_AllocationManagerStorage *AllocationManagerStorageCaller) GetSlashableMagnitudes(opts *bind.CallOpts, operator common.Address, strategies []common.Address, timestamp uint32) ([]OperatorSet, [][]uint64, error) {
	var out []interface{}
	err := _AllocationManagerStorage.contract.Call(opts, &out, "getSlashableMagnitudes", operator, strategies, timestamp)

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
func (_AllocationManagerStorage *AllocationManagerStorageSession) GetSlashableMagnitudes(operator common.Address, strategies []common.Address, timestamp uint32) ([]OperatorSet, [][]uint64, error) {
	return _AllocationManagerStorage.Contract.GetSlashableMagnitudes(&_AllocationManagerStorage.CallOpts, operator, strategies, timestamp)
}

// GetSlashableMagnitudes is a free data retrieval call binding the contract method 0x43b0592d.
//
// Solidity: function getSlashableMagnitudes(address operator, address[] strategies, uint32 timestamp) view returns((address,uint32)[], uint64[][])
func (_AllocationManagerStorage *AllocationManagerStorageCallerSession) GetSlashableMagnitudes(operator common.Address, strategies []common.Address, timestamp uint32) ([]OperatorSet, [][]uint64, error) {
	return _AllocationManagerStorage.Contract.GetSlashableMagnitudes(&_AllocationManagerStorage.CallOpts, operator, strategies, timestamp)
}

// GetSlashablePPM is a free data retrieval call binding the contract method 0x1c699d65.
//
// Solidity: function getSlashablePPM(address operator, (address,uint32) operatorSet, address[] strategies, uint32 timestamp, bool linear) view returns(uint24[])
func (_AllocationManagerStorage *AllocationManagerStorageCaller) GetSlashablePPM(opts *bind.CallOpts, operator common.Address, operatorSet OperatorSet, strategies []common.Address, timestamp uint32, linear bool) ([]*big.Int, error) {
	var out []interface{}
	err := _AllocationManagerStorage.contract.Call(opts, &out, "getSlashablePPM", operator, operatorSet, strategies, timestamp, linear)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetSlashablePPM is a free data retrieval call binding the contract method 0x1c699d65.
//
// Solidity: function getSlashablePPM(address operator, (address,uint32) operatorSet, address[] strategies, uint32 timestamp, bool linear) view returns(uint24[])
func (_AllocationManagerStorage *AllocationManagerStorageSession) GetSlashablePPM(operator common.Address, operatorSet OperatorSet, strategies []common.Address, timestamp uint32, linear bool) ([]*big.Int, error) {
	return _AllocationManagerStorage.Contract.GetSlashablePPM(&_AllocationManagerStorage.CallOpts, operator, operatorSet, strategies, timestamp, linear)
}

// GetSlashablePPM is a free data retrieval call binding the contract method 0x1c699d65.
//
// Solidity: function getSlashablePPM(address operator, (address,uint32) operatorSet, address[] strategies, uint32 timestamp, bool linear) view returns(uint24[])
func (_AllocationManagerStorage *AllocationManagerStorageCallerSession) GetSlashablePPM(operator common.Address, operatorSet OperatorSet, strategies []common.Address, timestamp uint32, linear bool) ([]*big.Int, error) {
	return _AllocationManagerStorage.Contract.GetSlashablePPM(&_AllocationManagerStorage.CallOpts, operator, operatorSet, strategies, timestamp, linear)
}

// GetTotalAndAllocatedMagnitudes is a free data retrieval call binding the contract method 0xc2348318.
//
// Solidity: function getTotalAndAllocatedMagnitudes(address operator, (address,uint32) operatorSet, address[] strategies) view returns(uint64[], uint64[])
func (_AllocationManagerStorage *AllocationManagerStorageCaller) GetTotalAndAllocatedMagnitudes(opts *bind.CallOpts, operator common.Address, operatorSet OperatorSet, strategies []common.Address) ([]uint64, []uint64, error) {
	var out []interface{}
	err := _AllocationManagerStorage.contract.Call(opts, &out, "getTotalAndAllocatedMagnitudes", operator, operatorSet, strategies)

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
func (_AllocationManagerStorage *AllocationManagerStorageSession) GetTotalAndAllocatedMagnitudes(operator common.Address, operatorSet OperatorSet, strategies []common.Address) ([]uint64, []uint64, error) {
	return _AllocationManagerStorage.Contract.GetTotalAndAllocatedMagnitudes(&_AllocationManagerStorage.CallOpts, operator, operatorSet, strategies)
}

// GetTotalAndAllocatedMagnitudes is a free data retrieval call binding the contract method 0xc2348318.
//
// Solidity: function getTotalAndAllocatedMagnitudes(address operator, (address,uint32) operatorSet, address[] strategies) view returns(uint64[], uint64[])
func (_AllocationManagerStorage *AllocationManagerStorageCallerSession) GetTotalAndAllocatedMagnitudes(operator common.Address, operatorSet OperatorSet, strategies []common.Address) ([]uint64, []uint64, error) {
	return _AllocationManagerStorage.Contract.GetTotalAndAllocatedMagnitudes(&_AllocationManagerStorage.CallOpts, operator, operatorSet, strategies)
}

// GetTotalMagnitude is a free data retrieval call binding the contract method 0xb47265e2.
//
// Solidity: function getTotalMagnitude(address operator, address strategy) view returns(uint64)
func (_AllocationManagerStorage *AllocationManagerStorageCaller) GetTotalMagnitude(opts *bind.CallOpts, operator common.Address, strategy common.Address) (uint64, error) {
	var out []interface{}
	err := _AllocationManagerStorage.contract.Call(opts, &out, "getTotalMagnitude", operator, strategy)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetTotalMagnitude is a free data retrieval call binding the contract method 0xb47265e2.
//
// Solidity: function getTotalMagnitude(address operator, address strategy) view returns(uint64)
func (_AllocationManagerStorage *AllocationManagerStorageSession) GetTotalMagnitude(operator common.Address, strategy common.Address) (uint64, error) {
	return _AllocationManagerStorage.Contract.GetTotalMagnitude(&_AllocationManagerStorage.CallOpts, operator, strategy)
}

// GetTotalMagnitude is a free data retrieval call binding the contract method 0xb47265e2.
//
// Solidity: function getTotalMagnitude(address operator, address strategy) view returns(uint64)
func (_AllocationManagerStorage *AllocationManagerStorageCallerSession) GetTotalMagnitude(operator common.Address, strategy common.Address) (uint64, error) {
	return _AllocationManagerStorage.Contract.GetTotalMagnitude(&_AllocationManagerStorage.CallOpts, operator, strategy)
}

// GetTotalMagnitudeAtTimestamp is a free data retrieval call binding the contract method 0xdabe97c7.
//
// Solidity: function getTotalMagnitudeAtTimestamp(address operator, address strategy, uint32 timestamp) view returns(uint64)
func (_AllocationManagerStorage *AllocationManagerStorageCaller) GetTotalMagnitudeAtTimestamp(opts *bind.CallOpts, operator common.Address, strategy common.Address, timestamp uint32) (uint64, error) {
	var out []interface{}
	err := _AllocationManagerStorage.contract.Call(opts, &out, "getTotalMagnitudeAtTimestamp", operator, strategy, timestamp)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetTotalMagnitudeAtTimestamp is a free data retrieval call binding the contract method 0xdabe97c7.
//
// Solidity: function getTotalMagnitudeAtTimestamp(address operator, address strategy, uint32 timestamp) view returns(uint64)
func (_AllocationManagerStorage *AllocationManagerStorageSession) GetTotalMagnitudeAtTimestamp(operator common.Address, strategy common.Address, timestamp uint32) (uint64, error) {
	return _AllocationManagerStorage.Contract.GetTotalMagnitudeAtTimestamp(&_AllocationManagerStorage.CallOpts, operator, strategy, timestamp)
}

// GetTotalMagnitudeAtTimestamp is a free data retrieval call binding the contract method 0xdabe97c7.
//
// Solidity: function getTotalMagnitudeAtTimestamp(address operator, address strategy, uint32 timestamp) view returns(uint64)
func (_AllocationManagerStorage *AllocationManagerStorageCallerSession) GetTotalMagnitudeAtTimestamp(operator common.Address, strategy common.Address, timestamp uint32) (uint64, error) {
	return _AllocationManagerStorage.Contract.GetTotalMagnitudeAtTimestamp(&_AllocationManagerStorage.CallOpts, operator, strategy, timestamp)
}

// GetTotalMagnitudes is a free data retrieval call binding the contract method 0x39a9a3ed.
//
// Solidity: function getTotalMagnitudes(address operator, address[] strategies) view returns(uint64[])
func (_AllocationManagerStorage *AllocationManagerStorageCaller) GetTotalMagnitudes(opts *bind.CallOpts, operator common.Address, strategies []common.Address) ([]uint64, error) {
	var out []interface{}
	err := _AllocationManagerStorage.contract.Call(opts, &out, "getTotalMagnitudes", operator, strategies)

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

// GetTotalMagnitudes is a free data retrieval call binding the contract method 0x39a9a3ed.
//
// Solidity: function getTotalMagnitudes(address operator, address[] strategies) view returns(uint64[])
func (_AllocationManagerStorage *AllocationManagerStorageSession) GetTotalMagnitudes(operator common.Address, strategies []common.Address) ([]uint64, error) {
	return _AllocationManagerStorage.Contract.GetTotalMagnitudes(&_AllocationManagerStorage.CallOpts, operator, strategies)
}

// GetTotalMagnitudes is a free data retrieval call binding the contract method 0x39a9a3ed.
//
// Solidity: function getTotalMagnitudes(address operator, address[] strategies) view returns(uint64[])
func (_AllocationManagerStorage *AllocationManagerStorageCallerSession) GetTotalMagnitudes(operator common.Address, strategies []common.Address) ([]uint64, error) {
	return _AllocationManagerStorage.Contract.GetTotalMagnitudes(&_AllocationManagerStorage.CallOpts, operator, strategies)
}

// GetTotalMagnitudesAtTimestamp is a free data retrieval call binding the contract method 0x858d0b47.
//
// Solidity: function getTotalMagnitudesAtTimestamp(address operator, address[] strategies, uint32 timestamp) view returns(uint64[])
func (_AllocationManagerStorage *AllocationManagerStorageCaller) GetTotalMagnitudesAtTimestamp(opts *bind.CallOpts, operator common.Address, strategies []common.Address, timestamp uint32) ([]uint64, error) {
	var out []interface{}
	err := _AllocationManagerStorage.contract.Call(opts, &out, "getTotalMagnitudesAtTimestamp", operator, strategies, timestamp)

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

// GetTotalMagnitudesAtTimestamp is a free data retrieval call binding the contract method 0x858d0b47.
//
// Solidity: function getTotalMagnitudesAtTimestamp(address operator, address[] strategies, uint32 timestamp) view returns(uint64[])
func (_AllocationManagerStorage *AllocationManagerStorageSession) GetTotalMagnitudesAtTimestamp(operator common.Address, strategies []common.Address, timestamp uint32) ([]uint64, error) {
	return _AllocationManagerStorage.Contract.GetTotalMagnitudesAtTimestamp(&_AllocationManagerStorage.CallOpts, operator, strategies, timestamp)
}

// GetTotalMagnitudesAtTimestamp is a free data retrieval call binding the contract method 0x858d0b47.
//
// Solidity: function getTotalMagnitudesAtTimestamp(address operator, address[] strategies, uint32 timestamp) view returns(uint64[])
func (_AllocationManagerStorage *AllocationManagerStorageCallerSession) GetTotalMagnitudesAtTimestamp(operator common.Address, strategies []common.Address, timestamp uint32) ([]uint64, error) {
	return _AllocationManagerStorage.Contract.GetTotalMagnitudesAtTimestamp(&_AllocationManagerStorage.CallOpts, operator, strategies, timestamp)
}

// IsOperatorSlashable is a free data retrieval call binding the contract method 0x1352c3e6.
//
// Solidity: function isOperatorSlashable(address operator, (address,uint32) operatorSet) view returns(bool)
func (_AllocationManagerStorage *AllocationManagerStorageCaller) IsOperatorSlashable(opts *bind.CallOpts, operator common.Address, operatorSet OperatorSet) (bool, error) {
	var out []interface{}
	err := _AllocationManagerStorage.contract.Call(opts, &out, "isOperatorSlashable", operator, operatorSet)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperatorSlashable is a free data retrieval call binding the contract method 0x1352c3e6.
//
// Solidity: function isOperatorSlashable(address operator, (address,uint32) operatorSet) view returns(bool)
func (_AllocationManagerStorage *AllocationManagerStorageSession) IsOperatorSlashable(operator common.Address, operatorSet OperatorSet) (bool, error) {
	return _AllocationManagerStorage.Contract.IsOperatorSlashable(&_AllocationManagerStorage.CallOpts, operator, operatorSet)
}

// IsOperatorSlashable is a free data retrieval call binding the contract method 0x1352c3e6.
//
// Solidity: function isOperatorSlashable(address operator, (address,uint32) operatorSet) view returns(bool)
func (_AllocationManagerStorage *AllocationManagerStorageCallerSession) IsOperatorSlashable(operator common.Address, operatorSet OperatorSet) (bool, error) {
	return _AllocationManagerStorage.Contract.IsOperatorSlashable(&_AllocationManagerStorage.CallOpts, operator, operatorSet)
}

// OperatorMagnitudeInfo is a free data retrieval call binding the contract method 0xcb13e56b.
//
// Solidity: function operatorMagnitudeInfo(address , address ) view returns(uint192 nextPendingFreeMagnitudeIndex, uint64 freeMagnitude)
func (_AllocationManagerStorage *AllocationManagerStorageCaller) OperatorMagnitudeInfo(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (struct {
	NextPendingFreeMagnitudeIndex *big.Int
	FreeMagnitude                 uint64
}, error) {
	var out []interface{}
	err := _AllocationManagerStorage.contract.Call(opts, &out, "operatorMagnitudeInfo", arg0, arg1)

	outstruct := new(struct {
		NextPendingFreeMagnitudeIndex *big.Int
		FreeMagnitude                 uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NextPendingFreeMagnitudeIndex = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.FreeMagnitude = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// OperatorMagnitudeInfo is a free data retrieval call binding the contract method 0xcb13e56b.
//
// Solidity: function operatorMagnitudeInfo(address , address ) view returns(uint192 nextPendingFreeMagnitudeIndex, uint64 freeMagnitude)
func (_AllocationManagerStorage *AllocationManagerStorageSession) OperatorMagnitudeInfo(arg0 common.Address, arg1 common.Address) (struct {
	NextPendingFreeMagnitudeIndex *big.Int
	FreeMagnitude                 uint64
}, error) {
	return _AllocationManagerStorage.Contract.OperatorMagnitudeInfo(&_AllocationManagerStorage.CallOpts, arg0, arg1)
}

// OperatorMagnitudeInfo is a free data retrieval call binding the contract method 0xcb13e56b.
//
// Solidity: function operatorMagnitudeInfo(address , address ) view returns(uint192 nextPendingFreeMagnitudeIndex, uint64 freeMagnitude)
func (_AllocationManagerStorage *AllocationManagerStorageCallerSession) OperatorMagnitudeInfo(arg0 common.Address, arg1 common.Address) (struct {
	NextPendingFreeMagnitudeIndex *big.Int
	FreeMagnitude                 uint64
}, error) {
	return _AllocationManagerStorage.Contract.OperatorMagnitudeInfo(&_AllocationManagerStorage.CallOpts, arg0, arg1)
}

// OperatorSaltIsSpent is a free data retrieval call binding the contract method 0x374823b5.
//
// Solidity: function operatorSaltIsSpent(address , bytes32 ) view returns(bool)
func (_AllocationManagerStorage *AllocationManagerStorageCaller) OperatorSaltIsSpent(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (bool, error) {
	var out []interface{}
	err := _AllocationManagerStorage.contract.Call(opts, &out, "operatorSaltIsSpent", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OperatorSaltIsSpent is a free data retrieval call binding the contract method 0x374823b5.
//
// Solidity: function operatorSaltIsSpent(address , bytes32 ) view returns(bool)
func (_AllocationManagerStorage *AllocationManagerStorageSession) OperatorSaltIsSpent(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _AllocationManagerStorage.Contract.OperatorSaltIsSpent(&_AllocationManagerStorage.CallOpts, arg0, arg1)
}

// OperatorSaltIsSpent is a free data retrieval call binding the contract method 0x374823b5.
//
// Solidity: function operatorSaltIsSpent(address , bytes32 ) view returns(bool)
func (_AllocationManagerStorage *AllocationManagerStorageCallerSession) OperatorSaltIsSpent(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _AllocationManagerStorage.Contract.OperatorSaltIsSpent(&_AllocationManagerStorage.CallOpts, arg0, arg1)
}

// CancelSalt is a paid mutator transaction binding the contract method 0xec76f442.
//
// Solidity: function cancelSalt(bytes32 salt) returns()
func (_AllocationManagerStorage *AllocationManagerStorageTransactor) CancelSalt(opts *bind.TransactOpts, salt [32]byte) (*types.Transaction, error) {
	return _AllocationManagerStorage.contract.Transact(opts, "cancelSalt", salt)
}

// CancelSalt is a paid mutator transaction binding the contract method 0xec76f442.
//
// Solidity: function cancelSalt(bytes32 salt) returns()
func (_AllocationManagerStorage *AllocationManagerStorageSession) CancelSalt(salt [32]byte) (*types.Transaction, error) {
	return _AllocationManagerStorage.Contract.CancelSalt(&_AllocationManagerStorage.TransactOpts, salt)
}

// CancelSalt is a paid mutator transaction binding the contract method 0xec76f442.
//
// Solidity: function cancelSalt(bytes32 salt) returns()
func (_AllocationManagerStorage *AllocationManagerStorageTransactorSession) CancelSalt(salt [32]byte) (*types.Transaction, error) {
	return _AllocationManagerStorage.Contract.CancelSalt(&_AllocationManagerStorage.TransactOpts, salt)
}

// ModifyAllocations is a paid mutator transaction binding the contract method 0xf8e91d16.
//
// Solidity: function modifyAllocations(address operator, (address,uint64,(address,uint32)[],uint64[])[] allocations, (bytes,bytes32,uint256) operatorSignature) returns()
func (_AllocationManagerStorage *AllocationManagerStorageTransactor) ModifyAllocations(opts *bind.TransactOpts, operator common.Address, allocations []IAllocationManagerMagnitudeAllocation, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _AllocationManagerStorage.contract.Transact(opts, "modifyAllocations", operator, allocations, operatorSignature)
}

// ModifyAllocations is a paid mutator transaction binding the contract method 0xf8e91d16.
//
// Solidity: function modifyAllocations(address operator, (address,uint64,(address,uint32)[],uint64[])[] allocations, (bytes,bytes32,uint256) operatorSignature) returns()
func (_AllocationManagerStorage *AllocationManagerStorageSession) ModifyAllocations(operator common.Address, allocations []IAllocationManagerMagnitudeAllocation, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _AllocationManagerStorage.Contract.ModifyAllocations(&_AllocationManagerStorage.TransactOpts, operator, allocations, operatorSignature)
}

// ModifyAllocations is a paid mutator transaction binding the contract method 0xf8e91d16.
//
// Solidity: function modifyAllocations(address operator, (address,uint64,(address,uint32)[],uint64[])[] allocations, (bytes,bytes32,uint256) operatorSignature) returns()
func (_AllocationManagerStorage *AllocationManagerStorageTransactorSession) ModifyAllocations(operator common.Address, allocations []IAllocationManagerMagnitudeAllocation, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _AllocationManagerStorage.Contract.ModifyAllocations(&_AllocationManagerStorage.TransactOpts, operator, allocations, operatorSignature)
}

// SetAllocationDelay is a paid mutator transaction binding the contract method 0x5c489bb5.
//
// Solidity: function setAllocationDelay(uint32 delay) returns()
func (_AllocationManagerStorage *AllocationManagerStorageTransactor) SetAllocationDelay(opts *bind.TransactOpts, delay uint32) (*types.Transaction, error) {
	return _AllocationManagerStorage.contract.Transact(opts, "setAllocationDelay", delay)
}

// SetAllocationDelay is a paid mutator transaction binding the contract method 0x5c489bb5.
//
// Solidity: function setAllocationDelay(uint32 delay) returns()
func (_AllocationManagerStorage *AllocationManagerStorageSession) SetAllocationDelay(delay uint32) (*types.Transaction, error) {
	return _AllocationManagerStorage.Contract.SetAllocationDelay(&_AllocationManagerStorage.TransactOpts, delay)
}

// SetAllocationDelay is a paid mutator transaction binding the contract method 0x5c489bb5.
//
// Solidity: function setAllocationDelay(uint32 delay) returns()
func (_AllocationManagerStorage *AllocationManagerStorageTransactorSession) SetAllocationDelay(delay uint32) (*types.Transaction, error) {
	return _AllocationManagerStorage.Contract.SetAllocationDelay(&_AllocationManagerStorage.TransactOpts, delay)
}

// SlashOperator is a paid mutator transaction binding the contract method 0xbd74a06c.
//
// Solidity: function slashOperator(address operator, uint32 operatorSetId, address[] strategies, uint16 bipsToSlash) returns()
func (_AllocationManagerStorage *AllocationManagerStorageTransactor) SlashOperator(opts *bind.TransactOpts, operator common.Address, operatorSetId uint32, strategies []common.Address, bipsToSlash uint16) (*types.Transaction, error) {
	return _AllocationManagerStorage.contract.Transact(opts, "slashOperator", operator, operatorSetId, strategies, bipsToSlash)
}

// SlashOperator is a paid mutator transaction binding the contract method 0xbd74a06c.
//
// Solidity: function slashOperator(address operator, uint32 operatorSetId, address[] strategies, uint16 bipsToSlash) returns()
func (_AllocationManagerStorage *AllocationManagerStorageSession) SlashOperator(operator common.Address, operatorSetId uint32, strategies []common.Address, bipsToSlash uint16) (*types.Transaction, error) {
	return _AllocationManagerStorage.Contract.SlashOperator(&_AllocationManagerStorage.TransactOpts, operator, operatorSetId, strategies, bipsToSlash)
}

// SlashOperator is a paid mutator transaction binding the contract method 0xbd74a06c.
//
// Solidity: function slashOperator(address operator, uint32 operatorSetId, address[] strategies, uint16 bipsToSlash) returns()
func (_AllocationManagerStorage *AllocationManagerStorageTransactorSession) SlashOperator(operator common.Address, operatorSetId uint32, strategies []common.Address, bipsToSlash uint16) (*types.Transaction, error) {
	return _AllocationManagerStorage.Contract.SlashOperator(&_AllocationManagerStorage.TransactOpts, operator, operatorSetId, strategies, bipsToSlash)
}

// UpdateFreeMagnitude is a paid mutator transaction binding the contract method 0xce770388.
//
// Solidity: function updateFreeMagnitude(address operator, address[] strategies, uint16[] numToComplete) returns()
func (_AllocationManagerStorage *AllocationManagerStorageTransactor) UpdateFreeMagnitude(opts *bind.TransactOpts, operator common.Address, strategies []common.Address, numToComplete []uint16) (*types.Transaction, error) {
	return _AllocationManagerStorage.contract.Transact(opts, "updateFreeMagnitude", operator, strategies, numToComplete)
}

// UpdateFreeMagnitude is a paid mutator transaction binding the contract method 0xce770388.
//
// Solidity: function updateFreeMagnitude(address operator, address[] strategies, uint16[] numToComplete) returns()
func (_AllocationManagerStorage *AllocationManagerStorageSession) UpdateFreeMagnitude(operator common.Address, strategies []common.Address, numToComplete []uint16) (*types.Transaction, error) {
	return _AllocationManagerStorage.Contract.UpdateFreeMagnitude(&_AllocationManagerStorage.TransactOpts, operator, strategies, numToComplete)
}

// UpdateFreeMagnitude is a paid mutator transaction binding the contract method 0xce770388.
//
// Solidity: function updateFreeMagnitude(address operator, address[] strategies, uint16[] numToComplete) returns()
func (_AllocationManagerStorage *AllocationManagerStorageTransactorSession) UpdateFreeMagnitude(operator common.Address, strategies []common.Address, numToComplete []uint16) (*types.Transaction, error) {
	return _AllocationManagerStorage.Contract.UpdateFreeMagnitude(&_AllocationManagerStorage.TransactOpts, operator, strategies, numToComplete)
}

// AllocationManagerStorageAllocationDelaySetIterator is returned from FilterAllocationDelaySet and is used to iterate over the raw logs and unpacked data for AllocationDelaySet events raised by the AllocationManagerStorage contract.
type AllocationManagerStorageAllocationDelaySetIterator struct {
	Event *AllocationManagerStorageAllocationDelaySet // Event containing the contract specifics and raw log

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
func (it *AllocationManagerStorageAllocationDelaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AllocationManagerStorageAllocationDelaySet)
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
		it.Event = new(AllocationManagerStorageAllocationDelaySet)
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
func (it *AllocationManagerStorageAllocationDelaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AllocationManagerStorageAllocationDelaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AllocationManagerStorageAllocationDelaySet represents a AllocationDelaySet event raised by the AllocationManagerStorage contract.
type AllocationManagerStorageAllocationDelaySet struct {
	Operator common.Address
	Delay    uint32
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAllocationDelaySet is a free log retrieval operation binding the contract event 0xc17479bf29bcb9669d4dd3580ba716c0424d52c939d248d49b07efc02a32952d.
//
// Solidity: event AllocationDelaySet(address operator, uint32 delay)
func (_AllocationManagerStorage *AllocationManagerStorageFilterer) FilterAllocationDelaySet(opts *bind.FilterOpts) (*AllocationManagerStorageAllocationDelaySetIterator, error) {

	logs, sub, err := _AllocationManagerStorage.contract.FilterLogs(opts, "AllocationDelaySet")
	if err != nil {
		return nil, err
	}
	return &AllocationManagerStorageAllocationDelaySetIterator{contract: _AllocationManagerStorage.contract, event: "AllocationDelaySet", logs: logs, sub: sub}, nil
}

// WatchAllocationDelaySet is a free log subscription operation binding the contract event 0xc17479bf29bcb9669d4dd3580ba716c0424d52c939d248d49b07efc02a32952d.
//
// Solidity: event AllocationDelaySet(address operator, uint32 delay)
func (_AllocationManagerStorage *AllocationManagerStorageFilterer) WatchAllocationDelaySet(opts *bind.WatchOpts, sink chan<- *AllocationManagerStorageAllocationDelaySet) (event.Subscription, error) {

	logs, sub, err := _AllocationManagerStorage.contract.WatchLogs(opts, "AllocationDelaySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AllocationManagerStorageAllocationDelaySet)
				if err := _AllocationManagerStorage.contract.UnpackLog(event, "AllocationDelaySet", log); err != nil {
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
func (_AllocationManagerStorage *AllocationManagerStorageFilterer) ParseAllocationDelaySet(log types.Log) (*AllocationManagerStorageAllocationDelaySet, error) {
	event := new(AllocationManagerStorageAllocationDelaySet)
	if err := _AllocationManagerStorage.contract.UnpackLog(event, "AllocationDelaySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AllocationManagerStorageMagnitudeAllocatedIterator is returned from FilterMagnitudeAllocated and is used to iterate over the raw logs and unpacked data for MagnitudeAllocated events raised by the AllocationManagerStorage contract.
type AllocationManagerStorageMagnitudeAllocatedIterator struct {
	Event *AllocationManagerStorageMagnitudeAllocated // Event containing the contract specifics and raw log

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
func (it *AllocationManagerStorageMagnitudeAllocatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AllocationManagerStorageMagnitudeAllocated)
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
		it.Event = new(AllocationManagerStorageMagnitudeAllocated)
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
func (it *AllocationManagerStorageMagnitudeAllocatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AllocationManagerStorageMagnitudeAllocatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AllocationManagerStorageMagnitudeAllocated represents a MagnitudeAllocated event raised by the AllocationManagerStorage contract.
type AllocationManagerStorageMagnitudeAllocated struct {
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
func (_AllocationManagerStorage *AllocationManagerStorageFilterer) FilterMagnitudeAllocated(opts *bind.FilterOpts) (*AllocationManagerStorageMagnitudeAllocatedIterator, error) {

	logs, sub, err := _AllocationManagerStorage.contract.FilterLogs(opts, "MagnitudeAllocated")
	if err != nil {
		return nil, err
	}
	return &AllocationManagerStorageMagnitudeAllocatedIterator{contract: _AllocationManagerStorage.contract, event: "MagnitudeAllocated", logs: logs, sub: sub}, nil
}

// WatchMagnitudeAllocated is a free log subscription operation binding the contract event 0x6d7d0079582cb2c5e70d4135b37f36711415ee6c260778b716bd65e026eb4f1a.
//
// Solidity: event MagnitudeAllocated(address operator, address strategy, (address,uint32) operatorSet, uint64 magnitudeToAllocate, uint32 effectTimestamp)
func (_AllocationManagerStorage *AllocationManagerStorageFilterer) WatchMagnitudeAllocated(opts *bind.WatchOpts, sink chan<- *AllocationManagerStorageMagnitudeAllocated) (event.Subscription, error) {

	logs, sub, err := _AllocationManagerStorage.contract.WatchLogs(opts, "MagnitudeAllocated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AllocationManagerStorageMagnitudeAllocated)
				if err := _AllocationManagerStorage.contract.UnpackLog(event, "MagnitudeAllocated", log); err != nil {
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
func (_AllocationManagerStorage *AllocationManagerStorageFilterer) ParseMagnitudeAllocated(log types.Log) (*AllocationManagerStorageMagnitudeAllocated, error) {
	event := new(AllocationManagerStorageMagnitudeAllocated)
	if err := _AllocationManagerStorage.contract.UnpackLog(event, "MagnitudeAllocated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AllocationManagerStorageMagnitudeDeallocationCompletedIterator is returned from FilterMagnitudeDeallocationCompleted and is used to iterate over the raw logs and unpacked data for MagnitudeDeallocationCompleted events raised by the AllocationManagerStorage contract.
type AllocationManagerStorageMagnitudeDeallocationCompletedIterator struct {
	Event *AllocationManagerStorageMagnitudeDeallocationCompleted // Event containing the contract specifics and raw log

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
func (it *AllocationManagerStorageMagnitudeDeallocationCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AllocationManagerStorageMagnitudeDeallocationCompleted)
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
		it.Event = new(AllocationManagerStorageMagnitudeDeallocationCompleted)
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
func (it *AllocationManagerStorageMagnitudeDeallocationCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AllocationManagerStorageMagnitudeDeallocationCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AllocationManagerStorageMagnitudeDeallocationCompleted represents a MagnitudeDeallocationCompleted event raised by the AllocationManagerStorage contract.
type AllocationManagerStorageMagnitudeDeallocationCompleted struct {
	Operator           common.Address
	Strategy           common.Address
	OperatorSet        OperatorSet
	FreeMagnitudeAdded uint64
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMagnitudeDeallocationCompleted is a free log retrieval operation binding the contract event 0x1e5c8e13c62c31d6252ac205e592477d643c7e95831d5b46d99a3c60c2fad8db.
//
// Solidity: event MagnitudeDeallocationCompleted(address operator, address strategy, (address,uint32) operatorSet, uint64 freeMagnitudeAdded)
func (_AllocationManagerStorage *AllocationManagerStorageFilterer) FilterMagnitudeDeallocationCompleted(opts *bind.FilterOpts) (*AllocationManagerStorageMagnitudeDeallocationCompletedIterator, error) {

	logs, sub, err := _AllocationManagerStorage.contract.FilterLogs(opts, "MagnitudeDeallocationCompleted")
	if err != nil {
		return nil, err
	}
	return &AllocationManagerStorageMagnitudeDeallocationCompletedIterator{contract: _AllocationManagerStorage.contract, event: "MagnitudeDeallocationCompleted", logs: logs, sub: sub}, nil
}

// WatchMagnitudeDeallocationCompleted is a free log subscription operation binding the contract event 0x1e5c8e13c62c31d6252ac205e592477d643c7e95831d5b46d99a3c60c2fad8db.
//
// Solidity: event MagnitudeDeallocationCompleted(address operator, address strategy, (address,uint32) operatorSet, uint64 freeMagnitudeAdded)
func (_AllocationManagerStorage *AllocationManagerStorageFilterer) WatchMagnitudeDeallocationCompleted(opts *bind.WatchOpts, sink chan<- *AllocationManagerStorageMagnitudeDeallocationCompleted) (event.Subscription, error) {

	logs, sub, err := _AllocationManagerStorage.contract.WatchLogs(opts, "MagnitudeDeallocationCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AllocationManagerStorageMagnitudeDeallocationCompleted)
				if err := _AllocationManagerStorage.contract.UnpackLog(event, "MagnitudeDeallocationCompleted", log); err != nil {
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
func (_AllocationManagerStorage *AllocationManagerStorageFilterer) ParseMagnitudeDeallocationCompleted(log types.Log) (*AllocationManagerStorageMagnitudeDeallocationCompleted, error) {
	event := new(AllocationManagerStorageMagnitudeDeallocationCompleted)
	if err := _AllocationManagerStorage.contract.UnpackLog(event, "MagnitudeDeallocationCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AllocationManagerStorageMagnitudeQueueDeallocatedIterator is returned from FilterMagnitudeQueueDeallocated and is used to iterate over the raw logs and unpacked data for MagnitudeQueueDeallocated events raised by the AllocationManagerStorage contract.
type AllocationManagerStorageMagnitudeQueueDeallocatedIterator struct {
	Event *AllocationManagerStorageMagnitudeQueueDeallocated // Event containing the contract specifics and raw log

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
func (it *AllocationManagerStorageMagnitudeQueueDeallocatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AllocationManagerStorageMagnitudeQueueDeallocated)
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
		it.Event = new(AllocationManagerStorageMagnitudeQueueDeallocated)
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
func (it *AllocationManagerStorageMagnitudeQueueDeallocatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AllocationManagerStorageMagnitudeQueueDeallocatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AllocationManagerStorageMagnitudeQueueDeallocated represents a MagnitudeQueueDeallocated event raised by the AllocationManagerStorage contract.
type AllocationManagerStorageMagnitudeQueueDeallocated struct {
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
func (_AllocationManagerStorage *AllocationManagerStorageFilterer) FilterMagnitudeQueueDeallocated(opts *bind.FilterOpts) (*AllocationManagerStorageMagnitudeQueueDeallocatedIterator, error) {

	logs, sub, err := _AllocationManagerStorage.contract.FilterLogs(opts, "MagnitudeQueueDeallocated")
	if err != nil {
		return nil, err
	}
	return &AllocationManagerStorageMagnitudeQueueDeallocatedIterator{contract: _AllocationManagerStorage.contract, event: "MagnitudeQueueDeallocated", logs: logs, sub: sub}, nil
}

// WatchMagnitudeQueueDeallocated is a free log subscription operation binding the contract event 0x2e68db1fe51107d7e451ae268d1631796989ab9d7925054e9b247854cb5be950.
//
// Solidity: event MagnitudeQueueDeallocated(address operator, address strategy, (address,uint32) operatorSet, uint64 magnitudeToDeallocate, uint32 completableTimestamp)
func (_AllocationManagerStorage *AllocationManagerStorageFilterer) WatchMagnitudeQueueDeallocated(opts *bind.WatchOpts, sink chan<- *AllocationManagerStorageMagnitudeQueueDeallocated) (event.Subscription, error) {

	logs, sub, err := _AllocationManagerStorage.contract.WatchLogs(opts, "MagnitudeQueueDeallocated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AllocationManagerStorageMagnitudeQueueDeallocated)
				if err := _AllocationManagerStorage.contract.UnpackLog(event, "MagnitudeQueueDeallocated", log); err != nil {
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
func (_AllocationManagerStorage *AllocationManagerStorageFilterer) ParseMagnitudeQueueDeallocated(log types.Log) (*AllocationManagerStorageMagnitudeQueueDeallocated, error) {
	event := new(AllocationManagerStorageMagnitudeQueueDeallocated)
	if err := _AllocationManagerStorage.contract.UnpackLog(event, "MagnitudeQueueDeallocated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AllocationManagerStorageOperatorSetCreatedIterator is returned from FilterOperatorSetCreated and is used to iterate over the raw logs and unpacked data for OperatorSetCreated events raised by the AllocationManagerStorage contract.
type AllocationManagerStorageOperatorSetCreatedIterator struct {
	Event *AllocationManagerStorageOperatorSetCreated // Event containing the contract specifics and raw log

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
func (it *AllocationManagerStorageOperatorSetCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AllocationManagerStorageOperatorSetCreated)
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
		it.Event = new(AllocationManagerStorageOperatorSetCreated)
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
func (it *AllocationManagerStorageOperatorSetCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AllocationManagerStorageOperatorSetCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AllocationManagerStorageOperatorSetCreated represents a OperatorSetCreated event raised by the AllocationManagerStorage contract.
type AllocationManagerStorageOperatorSetCreated struct {
	OperatorSet OperatorSet
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorSetCreated is a free log retrieval operation binding the contract event 0x31629285ead2335ae0933f86ed2ae63321f7af77b4e6eaabc42c057880977e6c.
//
// Solidity: event OperatorSetCreated((address,uint32) operatorSet)
func (_AllocationManagerStorage *AllocationManagerStorageFilterer) FilterOperatorSetCreated(opts *bind.FilterOpts) (*AllocationManagerStorageOperatorSetCreatedIterator, error) {

	logs, sub, err := _AllocationManagerStorage.contract.FilterLogs(opts, "OperatorSetCreated")
	if err != nil {
		return nil, err
	}
	return &AllocationManagerStorageOperatorSetCreatedIterator{contract: _AllocationManagerStorage.contract, event: "OperatorSetCreated", logs: logs, sub: sub}, nil
}

// WatchOperatorSetCreated is a free log subscription operation binding the contract event 0x31629285ead2335ae0933f86ed2ae63321f7af77b4e6eaabc42c057880977e6c.
//
// Solidity: event OperatorSetCreated((address,uint32) operatorSet)
func (_AllocationManagerStorage *AllocationManagerStorageFilterer) WatchOperatorSetCreated(opts *bind.WatchOpts, sink chan<- *AllocationManagerStorageOperatorSetCreated) (event.Subscription, error) {

	logs, sub, err := _AllocationManagerStorage.contract.WatchLogs(opts, "OperatorSetCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AllocationManagerStorageOperatorSetCreated)
				if err := _AllocationManagerStorage.contract.UnpackLog(event, "OperatorSetCreated", log); err != nil {
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
func (_AllocationManagerStorage *AllocationManagerStorageFilterer) ParseOperatorSetCreated(log types.Log) (*AllocationManagerStorageOperatorSetCreated, error) {
	event := new(AllocationManagerStorageOperatorSetCreated)
	if err := _AllocationManagerStorage.contract.UnpackLog(event, "OperatorSetCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AllocationManagerStorageOperatorSlashedIterator is returned from FilterOperatorSlashed and is used to iterate over the raw logs and unpacked data for OperatorSlashed events raised by the AllocationManagerStorage contract.
type AllocationManagerStorageOperatorSlashedIterator struct {
	Event *AllocationManagerStorageOperatorSlashed // Event containing the contract specifics and raw log

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
func (it *AllocationManagerStorageOperatorSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AllocationManagerStorageOperatorSlashed)
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
		it.Event = new(AllocationManagerStorageOperatorSlashed)
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
func (it *AllocationManagerStorageOperatorSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AllocationManagerStorageOperatorSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AllocationManagerStorageOperatorSlashed represents a OperatorSlashed event raised by the AllocationManagerStorage contract.
type AllocationManagerStorageOperatorSlashed struct {
	Operator      common.Address
	OperatorSetId uint32
	Strategy      common.Address
	BipsToSlash   uint16
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOperatorSlashed is a free log retrieval operation binding the contract event 0xe672839d3c371691acdb52de9fefc94b3dbf407dc0920ef566c7c059ad575b1c.
//
// Solidity: event OperatorSlashed(address operator, uint32 operatorSetId, address strategy, uint16 bipsToSlash)
func (_AllocationManagerStorage *AllocationManagerStorageFilterer) FilterOperatorSlashed(opts *bind.FilterOpts) (*AllocationManagerStorageOperatorSlashedIterator, error) {

	logs, sub, err := _AllocationManagerStorage.contract.FilterLogs(opts, "OperatorSlashed")
	if err != nil {
		return nil, err
	}
	return &AllocationManagerStorageOperatorSlashedIterator{contract: _AllocationManagerStorage.contract, event: "OperatorSlashed", logs: logs, sub: sub}, nil
}

// WatchOperatorSlashed is a free log subscription operation binding the contract event 0xe672839d3c371691acdb52de9fefc94b3dbf407dc0920ef566c7c059ad575b1c.
//
// Solidity: event OperatorSlashed(address operator, uint32 operatorSetId, address strategy, uint16 bipsToSlash)
func (_AllocationManagerStorage *AllocationManagerStorageFilterer) WatchOperatorSlashed(opts *bind.WatchOpts, sink chan<- *AllocationManagerStorageOperatorSlashed) (event.Subscription, error) {

	logs, sub, err := _AllocationManagerStorage.contract.WatchLogs(opts, "OperatorSlashed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AllocationManagerStorageOperatorSlashed)
				if err := _AllocationManagerStorage.contract.UnpackLog(event, "OperatorSlashed", log); err != nil {
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
func (_AllocationManagerStorage *AllocationManagerStorageFilterer) ParseOperatorSlashed(log types.Log) (*AllocationManagerStorageOperatorSlashed, error) {
	event := new(AllocationManagerStorageOperatorSlashed)
	if err := _AllocationManagerStorage.contract.UnpackLog(event, "OperatorSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
