// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package AVSDirectoryStorage

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

// IAVSDirectoryMagnitudeAllocation is an auto generated low-level Go binding around an user-defined struct.
type IAVSDirectoryMagnitudeAllocation struct {
	Strategy               common.Address
	ExpectedTotalMagnitude uint64
	OperatorSets           []IAVSDirectoryOperatorSet
	Magnitudes             []uint64
}

// IAVSDirectoryOperatorSet is an auto generated low-level Go binding around an user-defined struct.
type IAVSDirectoryOperatorSet struct {
	Avs           common.Address
	OperatorSetId uint32
}

// ISignatureUtilsSignatureWithSaltAndExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithSaltAndExpiry struct {
	Signature []byte
	Salt      [32]byte
	Expiry    *big.Int
}

// AVSDirectoryStorageMetaData contains all meta data concerning the AVSDirectoryStorage contract.
var AVSDirectoryStorageMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"DOMAIN_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"INITIAL_TOTAL_MAGNITUDE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAGNITUDE_ADJUSTMENT_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"OPERATOR_AVS_REGISTRATION_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"OPERATOR_SET_FORCE_DEREGISTRATION_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"OPERATOR_SET_REGISTRATION_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allocationDelay\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"isSet\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"allocationDelay\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"avsOperatorStatus\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIAVSDirectory.OperatorAVSRegistrationStatus\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"becomeOperatorSetAVS\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"calculateMagnitudeAllocationDigestHash\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allocations\",\"type\":\"tuple[]\",\"internalType\":\"structIAVSDirectory.MagnitudeAllocation[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"expectedTotalMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"operatorSets\",\"type\":\"tuple[]\",\"internalType\":\"structIAVSDirectory.OperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"magnitudes\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}]},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateOperatorAVSRegistrationDigestHash\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateOperatorSetForceDeregistrationTypehash\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateOperatorSetRegistrationDigestHash\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cancelSalt\",\"inputs\":[{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createOperatorSets\",\"inputs\":[{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregisterOperatorFromOperatorSets\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"domainSeparator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"forceDeregisterFromOperatorSets\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"freeMagnitude\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocatableMagnitude\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"numToComplete\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocationDelay\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSlashablePPM\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structIAVSDirectory.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"timestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"linear\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint24[]\",\"internalType\":\"uint24[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initializeAllocationDelay\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delay\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isMember\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structIAVSDirectory.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperatorSet\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperatorSetAVS\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperatorSlashable\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structIAVSDirectory.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"migrateOperatorsToOperatorSets\",\"inputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"modifyAllocations\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allocations\",\"type\":\"tuple[]\",\"internalType\":\"structIAVSDirectory.MagnitudeAllocation[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"expectedTotalMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"operatorSets\",\"type\":\"tuple[]\",\"internalType\":\"structIAVSDirectory.OperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"magnitudes\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}]},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"operatorSaltIsSpent\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorSetMemberCount\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorSetStatus\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"registered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"lastDeregisteredTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerOperatorToOperatorSets\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slashOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"bipsToSlash\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAVSMetadataURI\",\"inputs\":[{\"name\":\"metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateFreeMagnitude\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"numToComplete\",\"type\":\"uint8[]\",\"internalType\":\"uint8[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AVSMetadataURIUpdated\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"metadataURI\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AVSMigratedToOperatorSets\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MagnitudeAllocated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIAVSDirectory.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"magnitudeToAllocate\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"effectTimestamp\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MagnitudeDeallocationCompleted\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIAVSDirectory.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"freeMagnitudeAdded\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MagnitudeQueueDeallocated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIAVSDirectory.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"magnitudeToDeallocate\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"completableTimestamp\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorAVSRegistrationStatusUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"status\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumIAVSDirectory.OperatorAVSRegistrationStatus\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorAddedToOperatorSet\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIAVSDirectory.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorMigratedToOperatorSets\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"indexed\":false,\"internalType\":\"uint32[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRemovedFromOperatorSet\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIAVSDirectory.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSetCreated\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIAVSDirectory.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSlashed\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"bipsToSlash\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false}]",
}

// AVSDirectoryStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use AVSDirectoryStorageMetaData.ABI instead.
var AVSDirectoryStorageABI = AVSDirectoryStorageMetaData.ABI

// AVSDirectoryStorage is an auto generated Go binding around an Ethereum contract.
type AVSDirectoryStorage struct {
	AVSDirectoryStorageCaller     // Read-only binding to the contract
	AVSDirectoryStorageTransactor // Write-only binding to the contract
	AVSDirectoryStorageFilterer   // Log filterer for contract events
}

// AVSDirectoryStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type AVSDirectoryStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AVSDirectoryStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AVSDirectoryStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AVSDirectoryStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AVSDirectoryStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AVSDirectoryStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AVSDirectoryStorageSession struct {
	Contract     *AVSDirectoryStorage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// AVSDirectoryStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AVSDirectoryStorageCallerSession struct {
	Contract *AVSDirectoryStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// AVSDirectoryStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AVSDirectoryStorageTransactorSession struct {
	Contract     *AVSDirectoryStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// AVSDirectoryStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type AVSDirectoryStorageRaw struct {
	Contract *AVSDirectoryStorage // Generic contract binding to access the raw methods on
}

// AVSDirectoryStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AVSDirectoryStorageCallerRaw struct {
	Contract *AVSDirectoryStorageCaller // Generic read-only contract binding to access the raw methods on
}

// AVSDirectoryStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AVSDirectoryStorageTransactorRaw struct {
	Contract *AVSDirectoryStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAVSDirectoryStorage creates a new instance of AVSDirectoryStorage, bound to a specific deployed contract.
func NewAVSDirectoryStorage(address common.Address, backend bind.ContractBackend) (*AVSDirectoryStorage, error) {
	contract, err := bindAVSDirectoryStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AVSDirectoryStorage{AVSDirectoryStorageCaller: AVSDirectoryStorageCaller{contract: contract}, AVSDirectoryStorageTransactor: AVSDirectoryStorageTransactor{contract: contract}, AVSDirectoryStorageFilterer: AVSDirectoryStorageFilterer{contract: contract}}, nil
}

// NewAVSDirectoryStorageCaller creates a new read-only instance of AVSDirectoryStorage, bound to a specific deployed contract.
func NewAVSDirectoryStorageCaller(address common.Address, caller bind.ContractCaller) (*AVSDirectoryStorageCaller, error) {
	contract, err := bindAVSDirectoryStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AVSDirectoryStorageCaller{contract: contract}, nil
}

// NewAVSDirectoryStorageTransactor creates a new write-only instance of AVSDirectoryStorage, bound to a specific deployed contract.
func NewAVSDirectoryStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*AVSDirectoryStorageTransactor, error) {
	contract, err := bindAVSDirectoryStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AVSDirectoryStorageTransactor{contract: contract}, nil
}

// NewAVSDirectoryStorageFilterer creates a new log filterer instance of AVSDirectoryStorage, bound to a specific deployed contract.
func NewAVSDirectoryStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*AVSDirectoryStorageFilterer, error) {
	contract, err := bindAVSDirectoryStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AVSDirectoryStorageFilterer{contract: contract}, nil
}

// bindAVSDirectoryStorage binds a generic wrapper to an already deployed contract.
func bindAVSDirectoryStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AVSDirectoryStorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AVSDirectoryStorage *AVSDirectoryStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AVSDirectoryStorage.Contract.AVSDirectoryStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AVSDirectoryStorage *AVSDirectoryStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AVSDirectoryStorage.Contract.AVSDirectoryStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AVSDirectoryStorage *AVSDirectoryStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AVSDirectoryStorage.Contract.AVSDirectoryStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AVSDirectoryStorage *AVSDirectoryStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AVSDirectoryStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AVSDirectoryStorage *AVSDirectoryStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AVSDirectoryStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AVSDirectoryStorage *AVSDirectoryStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AVSDirectoryStorage.Contract.contract.Transact(opts, method, params...)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_AVSDirectoryStorage *AVSDirectoryStorageCaller) DOMAINTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AVSDirectoryStorage.contract.Call(opts, &out, "DOMAIN_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_AVSDirectoryStorage *AVSDirectoryStorageSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _AVSDirectoryStorage.Contract.DOMAINTYPEHASH(&_AVSDirectoryStorage.CallOpts)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_AVSDirectoryStorage *AVSDirectoryStorageCallerSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _AVSDirectoryStorage.Contract.DOMAINTYPEHASH(&_AVSDirectoryStorage.CallOpts)
}

// INITIALTOTALMAGNITUDE is a free data retrieval call binding the contract method 0x9a543ca4.
//
// Solidity: function INITIAL_TOTAL_MAGNITUDE() view returns(uint64)
func (_AVSDirectoryStorage *AVSDirectoryStorageCaller) INITIALTOTALMAGNITUDE(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _AVSDirectoryStorage.contract.Call(opts, &out, "INITIAL_TOTAL_MAGNITUDE")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// INITIALTOTALMAGNITUDE is a free data retrieval call binding the contract method 0x9a543ca4.
//
// Solidity: function INITIAL_TOTAL_MAGNITUDE() view returns(uint64)
func (_AVSDirectoryStorage *AVSDirectoryStorageSession) INITIALTOTALMAGNITUDE() (uint64, error) {
	return _AVSDirectoryStorage.Contract.INITIALTOTALMAGNITUDE(&_AVSDirectoryStorage.CallOpts)
}

// INITIALTOTALMAGNITUDE is a free data retrieval call binding the contract method 0x9a543ca4.
//
// Solidity: function INITIAL_TOTAL_MAGNITUDE() view returns(uint64)
func (_AVSDirectoryStorage *AVSDirectoryStorageCallerSession) INITIALTOTALMAGNITUDE() (uint64, error) {
	return _AVSDirectoryStorage.Contract.INITIALTOTALMAGNITUDE(&_AVSDirectoryStorage.CallOpts)
}

// MAGNITUDEADJUSTMENTTYPEHASH is a free data retrieval call binding the contract method 0x7b205de3.
//
// Solidity: function MAGNITUDE_ADJUSTMENT_TYPEHASH() view returns(bytes32)
func (_AVSDirectoryStorage *AVSDirectoryStorageCaller) MAGNITUDEADJUSTMENTTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AVSDirectoryStorage.contract.Call(opts, &out, "MAGNITUDE_ADJUSTMENT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MAGNITUDEADJUSTMENTTYPEHASH is a free data retrieval call binding the contract method 0x7b205de3.
//
// Solidity: function MAGNITUDE_ADJUSTMENT_TYPEHASH() view returns(bytes32)
func (_AVSDirectoryStorage *AVSDirectoryStorageSession) MAGNITUDEADJUSTMENTTYPEHASH() ([32]byte, error) {
	return _AVSDirectoryStorage.Contract.MAGNITUDEADJUSTMENTTYPEHASH(&_AVSDirectoryStorage.CallOpts)
}

// MAGNITUDEADJUSTMENTTYPEHASH is a free data retrieval call binding the contract method 0x7b205de3.
//
// Solidity: function MAGNITUDE_ADJUSTMENT_TYPEHASH() view returns(bytes32)
func (_AVSDirectoryStorage *AVSDirectoryStorageCallerSession) MAGNITUDEADJUSTMENTTYPEHASH() ([32]byte, error) {
	return _AVSDirectoryStorage.Contract.MAGNITUDEADJUSTMENTTYPEHASH(&_AVSDirectoryStorage.CallOpts)
}

// OPERATORAVSREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0xd79aceab.
//
// Solidity: function OPERATOR_AVS_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_AVSDirectoryStorage *AVSDirectoryStorageCaller) OPERATORAVSREGISTRATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AVSDirectoryStorage.contract.Call(opts, &out, "OPERATOR_AVS_REGISTRATION_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORAVSREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0xd79aceab.
//
// Solidity: function OPERATOR_AVS_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_AVSDirectoryStorage *AVSDirectoryStorageSession) OPERATORAVSREGISTRATIONTYPEHASH() ([32]byte, error) {
	return _AVSDirectoryStorage.Contract.OPERATORAVSREGISTRATIONTYPEHASH(&_AVSDirectoryStorage.CallOpts)
}

// OPERATORAVSREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0xd79aceab.
//
// Solidity: function OPERATOR_AVS_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_AVSDirectoryStorage *AVSDirectoryStorageCallerSession) OPERATORAVSREGISTRATIONTYPEHASH() ([32]byte, error) {
	return _AVSDirectoryStorage.Contract.OPERATORAVSREGISTRATIONTYPEHASH(&_AVSDirectoryStorage.CallOpts)
}

// OPERATORSETFORCEDEREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0xdce974b9.
//
// Solidity: function OPERATOR_SET_FORCE_DEREGISTRATION_TYPEHASH() view returns(bytes32)
func (_AVSDirectoryStorage *AVSDirectoryStorageCaller) OPERATORSETFORCEDEREGISTRATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AVSDirectoryStorage.contract.Call(opts, &out, "OPERATOR_SET_FORCE_DEREGISTRATION_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORSETFORCEDEREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0xdce974b9.
//
// Solidity: function OPERATOR_SET_FORCE_DEREGISTRATION_TYPEHASH() view returns(bytes32)
func (_AVSDirectoryStorage *AVSDirectoryStorageSession) OPERATORSETFORCEDEREGISTRATIONTYPEHASH() ([32]byte, error) {
	return _AVSDirectoryStorage.Contract.OPERATORSETFORCEDEREGISTRATIONTYPEHASH(&_AVSDirectoryStorage.CallOpts)
}

// OPERATORSETFORCEDEREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0xdce974b9.
//
// Solidity: function OPERATOR_SET_FORCE_DEREGISTRATION_TYPEHASH() view returns(bytes32)
func (_AVSDirectoryStorage *AVSDirectoryStorageCallerSession) OPERATORSETFORCEDEREGISTRATIONTYPEHASH() ([32]byte, error) {
	return _AVSDirectoryStorage.Contract.OPERATORSETFORCEDEREGISTRATIONTYPEHASH(&_AVSDirectoryStorage.CallOpts)
}

// OPERATORSETREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0xc825fe68.
//
// Solidity: function OPERATOR_SET_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_AVSDirectoryStorage *AVSDirectoryStorageCaller) OPERATORSETREGISTRATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AVSDirectoryStorage.contract.Call(opts, &out, "OPERATOR_SET_REGISTRATION_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORSETREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0xc825fe68.
//
// Solidity: function OPERATOR_SET_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_AVSDirectoryStorage *AVSDirectoryStorageSession) OPERATORSETREGISTRATIONTYPEHASH() ([32]byte, error) {
	return _AVSDirectoryStorage.Contract.OPERATORSETREGISTRATIONTYPEHASH(&_AVSDirectoryStorage.CallOpts)
}

// OPERATORSETREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0xc825fe68.
//
// Solidity: function OPERATOR_SET_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_AVSDirectoryStorage *AVSDirectoryStorageCallerSession) OPERATORSETREGISTRATIONTYPEHASH() ([32]byte, error) {
	return _AVSDirectoryStorage.Contract.OPERATORSETREGISTRATIONTYPEHASH(&_AVSDirectoryStorage.CallOpts)
}

// AllocationDelay is a free data retrieval call binding the contract method 0xbe4e1fd3.
//
// Solidity: function allocationDelay(address ) view returns(bool isSet, uint32 allocationDelay)
func (_AVSDirectoryStorage *AVSDirectoryStorageCaller) AllocationDelay(opts *bind.CallOpts, arg0 common.Address) (struct {
	IsSet           bool
	AllocationDelay uint32
}, error) {
	var out []interface{}
	err := _AVSDirectoryStorage.contract.Call(opts, &out, "allocationDelay", arg0)

	outstruct := new(struct {
		IsSet           bool
		AllocationDelay uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsSet = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.AllocationDelay = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

// AllocationDelay is a free data retrieval call binding the contract method 0xbe4e1fd3.
//
// Solidity: function allocationDelay(address ) view returns(bool isSet, uint32 allocationDelay)
func (_AVSDirectoryStorage *AVSDirectoryStorageSession) AllocationDelay(arg0 common.Address) (struct {
	IsSet           bool
	AllocationDelay uint32
}, error) {
	return _AVSDirectoryStorage.Contract.AllocationDelay(&_AVSDirectoryStorage.CallOpts, arg0)
}

// AllocationDelay is a free data retrieval call binding the contract method 0xbe4e1fd3.
//
// Solidity: function allocationDelay(address ) view returns(bool isSet, uint32 allocationDelay)
func (_AVSDirectoryStorage *AVSDirectoryStorageCallerSession) AllocationDelay(arg0 common.Address) (struct {
	IsSet           bool
	AllocationDelay uint32
}, error) {
	return _AVSDirectoryStorage.Contract.AllocationDelay(&_AVSDirectoryStorage.CallOpts, arg0)
}

// AvsOperatorStatus is a free data retrieval call binding the contract method 0x49075da3.
//
// Solidity: function avsOperatorStatus(address , address ) view returns(uint8)
func (_AVSDirectoryStorage *AVSDirectoryStorageCaller) AvsOperatorStatus(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (uint8, error) {
	var out []interface{}
	err := _AVSDirectoryStorage.contract.Call(opts, &out, "avsOperatorStatus", arg0, arg1)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// AvsOperatorStatus is a free data retrieval call binding the contract method 0x49075da3.
//
// Solidity: function avsOperatorStatus(address , address ) view returns(uint8)
func (_AVSDirectoryStorage *AVSDirectoryStorageSession) AvsOperatorStatus(arg0 common.Address, arg1 common.Address) (uint8, error) {
	return _AVSDirectoryStorage.Contract.AvsOperatorStatus(&_AVSDirectoryStorage.CallOpts, arg0, arg1)
}

// AvsOperatorStatus is a free data retrieval call binding the contract method 0x49075da3.
//
// Solidity: function avsOperatorStatus(address , address ) view returns(uint8)
func (_AVSDirectoryStorage *AVSDirectoryStorageCallerSession) AvsOperatorStatus(arg0 common.Address, arg1 common.Address) (uint8, error) {
	return _AVSDirectoryStorage.Contract.AvsOperatorStatus(&_AVSDirectoryStorage.CallOpts, arg0, arg1)
}

// CalculateMagnitudeAllocationDigestHash is a free data retrieval call binding the contract method 0x686b686e.
//
// Solidity: function calculateMagnitudeAllocationDigestHash(address operator, (address,uint64,(address,uint32)[],uint64[])[] allocations, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_AVSDirectoryStorage *AVSDirectoryStorageCaller) CalculateMagnitudeAllocationDigestHash(opts *bind.CallOpts, operator common.Address, allocations []IAVSDirectoryMagnitudeAllocation, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _AVSDirectoryStorage.contract.Call(opts, &out, "calculateMagnitudeAllocationDigestHash", operator, allocations, salt, expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateMagnitudeAllocationDigestHash is a free data retrieval call binding the contract method 0x686b686e.
//
// Solidity: function calculateMagnitudeAllocationDigestHash(address operator, (address,uint64,(address,uint32)[],uint64[])[] allocations, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_AVSDirectoryStorage *AVSDirectoryStorageSession) CalculateMagnitudeAllocationDigestHash(operator common.Address, allocations []IAVSDirectoryMagnitudeAllocation, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _AVSDirectoryStorage.Contract.CalculateMagnitudeAllocationDigestHash(&_AVSDirectoryStorage.CallOpts, operator, allocations, salt, expiry)
}

// CalculateMagnitudeAllocationDigestHash is a free data retrieval call binding the contract method 0x686b686e.
//
// Solidity: function calculateMagnitudeAllocationDigestHash(address operator, (address,uint64,(address,uint32)[],uint64[])[] allocations, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_AVSDirectoryStorage *AVSDirectoryStorageCallerSession) CalculateMagnitudeAllocationDigestHash(operator common.Address, allocations []IAVSDirectoryMagnitudeAllocation, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _AVSDirectoryStorage.Contract.CalculateMagnitudeAllocationDigestHash(&_AVSDirectoryStorage.CallOpts, operator, allocations, salt, expiry)
}

// CalculateOperatorAVSRegistrationDigestHash is a free data retrieval call binding the contract method 0xa1060c88.
//
// Solidity: function calculateOperatorAVSRegistrationDigestHash(address operator, address avs, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_AVSDirectoryStorage *AVSDirectoryStorageCaller) CalculateOperatorAVSRegistrationDigestHash(opts *bind.CallOpts, operator common.Address, avs common.Address, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _AVSDirectoryStorage.contract.Call(opts, &out, "calculateOperatorAVSRegistrationDigestHash", operator, avs, salt, expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateOperatorAVSRegistrationDigestHash is a free data retrieval call binding the contract method 0xa1060c88.
//
// Solidity: function calculateOperatorAVSRegistrationDigestHash(address operator, address avs, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_AVSDirectoryStorage *AVSDirectoryStorageSession) CalculateOperatorAVSRegistrationDigestHash(operator common.Address, avs common.Address, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _AVSDirectoryStorage.Contract.CalculateOperatorAVSRegistrationDigestHash(&_AVSDirectoryStorage.CallOpts, operator, avs, salt, expiry)
}

// CalculateOperatorAVSRegistrationDigestHash is a free data retrieval call binding the contract method 0xa1060c88.
//
// Solidity: function calculateOperatorAVSRegistrationDigestHash(address operator, address avs, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_AVSDirectoryStorage *AVSDirectoryStorageCallerSession) CalculateOperatorAVSRegistrationDigestHash(operator common.Address, avs common.Address, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _AVSDirectoryStorage.Contract.CalculateOperatorAVSRegistrationDigestHash(&_AVSDirectoryStorage.CallOpts, operator, avs, salt, expiry)
}

// CalculateOperatorSetForceDeregistrationTypehash is a free data retrieval call binding the contract method 0xb2841d48.
//
// Solidity: function calculateOperatorSetForceDeregistrationTypehash(address avs, uint32[] operatorSetIds, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_AVSDirectoryStorage *AVSDirectoryStorageCaller) CalculateOperatorSetForceDeregistrationTypehash(opts *bind.CallOpts, avs common.Address, operatorSetIds []uint32, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _AVSDirectoryStorage.contract.Call(opts, &out, "calculateOperatorSetForceDeregistrationTypehash", avs, operatorSetIds, salt, expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateOperatorSetForceDeregistrationTypehash is a free data retrieval call binding the contract method 0xb2841d48.
//
// Solidity: function calculateOperatorSetForceDeregistrationTypehash(address avs, uint32[] operatorSetIds, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_AVSDirectoryStorage *AVSDirectoryStorageSession) CalculateOperatorSetForceDeregistrationTypehash(avs common.Address, operatorSetIds []uint32, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _AVSDirectoryStorage.Contract.CalculateOperatorSetForceDeregistrationTypehash(&_AVSDirectoryStorage.CallOpts, avs, operatorSetIds, salt, expiry)
}

// CalculateOperatorSetForceDeregistrationTypehash is a free data retrieval call binding the contract method 0xb2841d48.
//
// Solidity: function calculateOperatorSetForceDeregistrationTypehash(address avs, uint32[] operatorSetIds, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_AVSDirectoryStorage *AVSDirectoryStorageCallerSession) CalculateOperatorSetForceDeregistrationTypehash(avs common.Address, operatorSetIds []uint32, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _AVSDirectoryStorage.Contract.CalculateOperatorSetForceDeregistrationTypehash(&_AVSDirectoryStorage.CallOpts, avs, operatorSetIds, salt, expiry)
}

// CalculateOperatorSetRegistrationDigestHash is a free data retrieval call binding the contract method 0x955e6696.
//
// Solidity: function calculateOperatorSetRegistrationDigestHash(address avs, uint32[] operatorSetIds, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_AVSDirectoryStorage *AVSDirectoryStorageCaller) CalculateOperatorSetRegistrationDigestHash(opts *bind.CallOpts, avs common.Address, operatorSetIds []uint32, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _AVSDirectoryStorage.contract.Call(opts, &out, "calculateOperatorSetRegistrationDigestHash", avs, operatorSetIds, salt, expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateOperatorSetRegistrationDigestHash is a free data retrieval call binding the contract method 0x955e6696.
//
// Solidity: function calculateOperatorSetRegistrationDigestHash(address avs, uint32[] operatorSetIds, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_AVSDirectoryStorage *AVSDirectoryStorageSession) CalculateOperatorSetRegistrationDigestHash(avs common.Address, operatorSetIds []uint32, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _AVSDirectoryStorage.Contract.CalculateOperatorSetRegistrationDigestHash(&_AVSDirectoryStorage.CallOpts, avs, operatorSetIds, salt, expiry)
}

// CalculateOperatorSetRegistrationDigestHash is a free data retrieval call binding the contract method 0x955e6696.
//
// Solidity: function calculateOperatorSetRegistrationDigestHash(address avs, uint32[] operatorSetIds, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_AVSDirectoryStorage *AVSDirectoryStorageCallerSession) CalculateOperatorSetRegistrationDigestHash(avs common.Address, operatorSetIds []uint32, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _AVSDirectoryStorage.Contract.CalculateOperatorSetRegistrationDigestHash(&_AVSDirectoryStorage.CallOpts, avs, operatorSetIds, salt, expiry)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_AVSDirectoryStorage *AVSDirectoryStorageCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AVSDirectoryStorage.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_AVSDirectoryStorage *AVSDirectoryStorageSession) Delegation() (common.Address, error) {
	return _AVSDirectoryStorage.Contract.Delegation(&_AVSDirectoryStorage.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_AVSDirectoryStorage *AVSDirectoryStorageCallerSession) Delegation() (common.Address, error) {
	return _AVSDirectoryStorage.Contract.Delegation(&_AVSDirectoryStorage.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_AVSDirectoryStorage *AVSDirectoryStorageCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> 6b2ab7ed (build: bindings)
	err := _AVSDirectoryStorage.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_AVSDirectoryStorage *AVSDirectoryStorageSession) DomainSeparator() ([32]byte, error) {
	return _AVSDirectoryStorage.Contract.DomainSeparator(&_AVSDirectoryStorage.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_AVSDirectoryStorage *AVSDirectoryStorageCallerSession) DomainSeparator() ([32]byte, error) {
	return _AVSDirectoryStorage.Contract.DomainSeparator(&_AVSDirectoryStorage.CallOpts)
}

<<<<<<< HEAD
// GetNumOperatorsInOperatorSet is a free data retrieval call binding the contract method 0x1023aa35.
//
// Solidity: function getNumOperatorsInOperatorSet((address,uint32) operatorSet) view returns(uint256)
func (_AVSDirectoryStorage *AVSDirectoryStorageCaller) GetNumOperatorsInOperatorSet(opts *bind.CallOpts, operatorSet IAVSDirectoryOperatorSet) (*big.Int, error) {
	var out []interface{}
	err := _AVSDirectoryStorage.contract.Call(opts, &out, "getNumOperatorsInOperatorSet", operatorSet)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
=======
// FreeMagnitude is a free data retrieval call binding the contract method 0x5fd6abfd.
//
// Solidity: function freeMagnitude(address , address ) view returns(uint64)
func (_AVSDirectoryStorage *AVSDirectoryStorageCaller) FreeMagnitude(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (uint64, error) {
	var out []interface{}
	err := _AVSDirectoryStorage.contract.Call(opts, &out, "freeMagnitude", arg0, arg1)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
>>>>>>> 6b2ab7ed (build: bindings)

	return out0, err

}

<<<<<<< HEAD
// GetNumOperatorsInOperatorSet is a free data retrieval call binding the contract method 0x1023aa35.
//
// Solidity: function getNumOperatorsInOperatorSet((address,uint32) operatorSet) view returns(uint256)
func (_AVSDirectoryStorage *AVSDirectoryStorageSession) GetNumOperatorsInOperatorSet(operatorSet IAVSDirectoryOperatorSet) (*big.Int, error) {
	return _AVSDirectoryStorage.Contract.GetNumOperatorsInOperatorSet(&_AVSDirectoryStorage.CallOpts, operatorSet)
}

// GetNumOperatorsInOperatorSet is a free data retrieval call binding the contract method 0x1023aa35.
//
// Solidity: function getNumOperatorsInOperatorSet((address,uint32) operatorSet) view returns(uint256)
func (_AVSDirectoryStorage *AVSDirectoryStorageCallerSession) GetNumOperatorsInOperatorSet(operatorSet IAVSDirectoryOperatorSet) (*big.Int, error) {
	return _AVSDirectoryStorage.Contract.GetNumOperatorsInOperatorSet(&_AVSDirectoryStorage.CallOpts, operatorSet)
}

// GetOperatorSetsOfOperator is a free data retrieval call binding the contract method 0x16ae76cb.
//
// Solidity: function getOperatorSetsOfOperator(address operator, uint256 start, uint256 length) view returns((address,uint32)[] operatorSets)
func (_AVSDirectoryStorage *AVSDirectoryStorageCaller) GetOperatorSetsOfOperator(opts *bind.CallOpts, operator common.Address, start *big.Int, length *big.Int) ([]IAVSDirectoryOperatorSet, error) {
	var out []interface{}
	err := _AVSDirectoryStorage.contract.Call(opts, &out, "getOperatorSetsOfOperator", operator, start, length)

	if err != nil {
		return *new([]IAVSDirectoryOperatorSet), err
	}

	out0 := *abi.ConvertType(out[0], new([]IAVSDirectoryOperatorSet)).(*[]IAVSDirectoryOperatorSet)
=======
// FreeMagnitude is a free data retrieval call binding the contract method 0x5fd6abfd.
//
// Solidity: function freeMagnitude(address , address ) view returns(uint64)
func (_AVSDirectoryStorage *AVSDirectoryStorageSession) FreeMagnitude(arg0 common.Address, arg1 common.Address) (uint64, error) {
	return _AVSDirectoryStorage.Contract.FreeMagnitude(&_AVSDirectoryStorage.CallOpts, arg0, arg1)
}

// FreeMagnitude is a free data retrieval call binding the contract method 0x5fd6abfd.
//
// Solidity: function freeMagnitude(address , address ) view returns(uint64)
func (_AVSDirectoryStorage *AVSDirectoryStorageCallerSession) FreeMagnitude(arg0 common.Address, arg1 common.Address) (uint64, error) {
	return _AVSDirectoryStorage.Contract.FreeMagnitude(&_AVSDirectoryStorage.CallOpts, arg0, arg1)
}

// GetAllocatableMagnitude is a free data retrieval call binding the contract method 0xdd63ac60.
//
// Solidity: function getAllocatableMagnitude(address operator, address strategy, uint8 numToComplete) view returns(uint64)
func (_AVSDirectoryStorage *AVSDirectoryStorageCaller) GetAllocatableMagnitude(opts *bind.CallOpts, operator common.Address, strategy common.Address, numToComplete uint8) (uint64, error) {
	var out []interface{}
	err := _AVSDirectoryStorage.contract.Call(opts, &out, "getAllocatableMagnitude", operator, strategy, numToComplete)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetAllocatableMagnitude is a free data retrieval call binding the contract method 0xdd63ac60.
//
// Solidity: function getAllocatableMagnitude(address operator, address strategy, uint8 numToComplete) view returns(uint64)
func (_AVSDirectoryStorage *AVSDirectoryStorageSession) GetAllocatableMagnitude(operator common.Address, strategy common.Address, numToComplete uint8) (uint64, error) {
	return _AVSDirectoryStorage.Contract.GetAllocatableMagnitude(&_AVSDirectoryStorage.CallOpts, operator, strategy, numToComplete)
}

// GetAllocatableMagnitude is a free data retrieval call binding the contract method 0xdd63ac60.
//
// Solidity: function getAllocatableMagnitude(address operator, address strategy, uint8 numToComplete) view returns(uint64)
func (_AVSDirectoryStorage *AVSDirectoryStorageCallerSession) GetAllocatableMagnitude(operator common.Address, strategy common.Address, numToComplete uint8) (uint64, error) {
	return _AVSDirectoryStorage.Contract.GetAllocatableMagnitude(&_AVSDirectoryStorage.CallOpts, operator, strategy, numToComplete)
}

// GetAllocationDelay is a free data retrieval call binding the contract method 0xb9fbaed1.
//
// Solidity: function getAllocationDelay(address operator) view returns(uint32)
func (_AVSDirectoryStorage *AVSDirectoryStorageCaller) GetAllocationDelay(opts *bind.CallOpts, operator common.Address) (uint32, error) {
	var out []interface{}
	err := _AVSDirectoryStorage.contract.Call(opts, &out, "getAllocationDelay", operator)

	if err != nil {
		return *new(uint32), err
	}

<<<<<<< HEAD
	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)
>>>>>>> 6b2ab7ed (build: bindings)
=======
	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)
>>>>>>> 49569901 (feat: set allocations (#691))

	return out0, err

}

<<<<<<< HEAD
<<<<<<< HEAD
// GetOperatorSetsOfOperator is a free data retrieval call binding the contract method 0x16ae76cb.
//
// Solidity: function getOperatorSetsOfOperator(address operator, uint256 start, uint256 length) view returns((address,uint32)[] operatorSets)
func (_AVSDirectoryStorage *AVSDirectoryStorageSession) GetOperatorSetsOfOperator(operator common.Address, start *big.Int, length *big.Int) ([]IAVSDirectoryOperatorSet, error) {
	return _AVSDirectoryStorage.Contract.GetOperatorSetsOfOperator(&_AVSDirectoryStorage.CallOpts, operator, start, length)
}

// GetOperatorSetsOfOperator is a free data retrieval call binding the contract method 0x16ae76cb.
//
// Solidity: function getOperatorSetsOfOperator(address operator, uint256 start, uint256 length) view returns((address,uint32)[] operatorSets)
func (_AVSDirectoryStorage *AVSDirectoryStorageCallerSession) GetOperatorSetsOfOperator(operator common.Address, start *big.Int, length *big.Int) ([]IAVSDirectoryOperatorSet, error) {
	return _AVSDirectoryStorage.Contract.GetOperatorSetsOfOperator(&_AVSDirectoryStorage.CallOpts, operator, start, length)
}

// GetOperatorsInOperatorSet is a free data retrieval call binding the contract method 0x7357723b.
//
// Solidity: function getOperatorsInOperatorSet((address,uint32) operatorSet, uint256 start, uint256 length) view returns(address[] operators)
func (_AVSDirectoryStorage *AVSDirectoryStorageCaller) GetOperatorsInOperatorSet(opts *bind.CallOpts, operatorSet IAVSDirectoryOperatorSet, start *big.Int, length *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _AVSDirectoryStorage.contract.Call(opts, &out, "getOperatorsInOperatorSet", operatorSet, start, length)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOperatorsInOperatorSet is a free data retrieval call binding the contract method 0x7357723b.
//
// Solidity: function getOperatorsInOperatorSet((address,uint32) operatorSet, uint256 start, uint256 length) view returns(address[] operators)
func (_AVSDirectoryStorage *AVSDirectoryStorageSession) GetOperatorsInOperatorSet(operatorSet IAVSDirectoryOperatorSet, start *big.Int, length *big.Int) ([]common.Address, error) {
	return _AVSDirectoryStorage.Contract.GetOperatorsInOperatorSet(&_AVSDirectoryStorage.CallOpts, operatorSet, start, length)
}

// GetOperatorsInOperatorSet is a free data retrieval call binding the contract method 0x7357723b.
//
// Solidity: function getOperatorsInOperatorSet((address,uint32) operatorSet, uint256 start, uint256 length) view returns(address[] operators)
func (_AVSDirectoryStorage *AVSDirectoryStorageCallerSession) GetOperatorsInOperatorSet(operatorSet IAVSDirectoryOperatorSet, start *big.Int, length *big.Int) ([]common.Address, error) {
	return _AVSDirectoryStorage.Contract.GetOperatorsInOperatorSet(&_AVSDirectoryStorage.CallOpts, operatorSet, start, length)
}

// InTotalOperatorSets is a free data retrieval call binding the contract method 0xcbdf0e42.
//
// Solidity: function inTotalOperatorSets(address operator) view returns(uint256)
func (_AVSDirectoryStorage *AVSDirectoryStorageCaller) InTotalOperatorSets(opts *bind.CallOpts, operator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AVSDirectoryStorage.contract.Call(opts, &out, "inTotalOperatorSets", operator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InTotalOperatorSets is a free data retrieval call binding the contract method 0xcbdf0e42.
//
// Solidity: function inTotalOperatorSets(address operator) view returns(uint256)
func (_AVSDirectoryStorage *AVSDirectoryStorageSession) InTotalOperatorSets(operator common.Address) (*big.Int, error) {
	return _AVSDirectoryStorage.Contract.InTotalOperatorSets(&_AVSDirectoryStorage.CallOpts, operator)
}

// InTotalOperatorSets is a free data retrieval call binding the contract method 0xcbdf0e42.
//
// Solidity: function inTotalOperatorSets(address operator) view returns(uint256)
func (_AVSDirectoryStorage *AVSDirectoryStorageCallerSession) InTotalOperatorSets(operator common.Address) (*big.Int, error) {
	return _AVSDirectoryStorage.Contract.InTotalOperatorSets(&_AVSDirectoryStorage.CallOpts, operator)
=======
// GetSlashableBips is a free data retrieval call binding the contract method 0x33429a6a.
=======
// GetAllocationDelay is a free data retrieval call binding the contract method 0xb9fbaed1.
>>>>>>> 49569901 (feat: set allocations (#691))
//
// Solidity: function getAllocationDelay(address operator) view returns(uint32)
func (_AVSDirectoryStorage *AVSDirectoryStorageSession) GetAllocationDelay(operator common.Address) (uint32, error) {
	return _AVSDirectoryStorage.Contract.GetAllocationDelay(&_AVSDirectoryStorage.CallOpts, operator)
}

// GetAllocationDelay is a free data retrieval call binding the contract method 0xb9fbaed1.
//
<<<<<<< HEAD
// Solidity: function getSlashableBips(address operator, (address,uint32) operatorSet, address strategy, uint32 timestamp) view returns(uint16)
func (_AVSDirectoryStorage *AVSDirectoryStorageCallerSession) GetSlashableBips(operator common.Address, operatorSet IAVSDirectoryOperatorSet, strategy common.Address, timestamp uint32) (uint16, error) {
	return _AVSDirectoryStorage.Contract.GetSlashableBips(&_AVSDirectoryStorage.CallOpts, operator, operatorSet, strategy, timestamp)
>>>>>>> 6b2ab7ed (build: bindings)
=======
// Solidity: function getAllocationDelay(address operator) view returns(uint32)
func (_AVSDirectoryStorage *AVSDirectoryStorageCallerSession) GetAllocationDelay(operator common.Address) (uint32, error) {
	return _AVSDirectoryStorage.Contract.GetAllocationDelay(&_AVSDirectoryStorage.CallOpts, operator)
}

// GetSlashablePPM is a free data retrieval call binding the contract method 0x1c699d65.
//
// Solidity: function getSlashablePPM(address operator, (address,uint32) operatorSet, address[] strategies, uint32 timestamp, bool linear) view returns(uint24[])
func (_AVSDirectoryStorage *AVSDirectoryStorageCaller) GetSlashablePPM(opts *bind.CallOpts, operator common.Address, operatorSet IAVSDirectoryOperatorSet, strategies []common.Address, timestamp uint32, linear bool) ([]*big.Int, error) {
	var out []interface{}
	err := _AVSDirectoryStorage.contract.Call(opts, &out, "getSlashablePPM", operator, operatorSet, strategies, timestamp, linear)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetSlashablePPM is a free data retrieval call binding the contract method 0x1c699d65.
//
// Solidity: function getSlashablePPM(address operator, (address,uint32) operatorSet, address[] strategies, uint32 timestamp, bool linear) view returns(uint24[])
func (_AVSDirectoryStorage *AVSDirectoryStorageSession) GetSlashablePPM(operator common.Address, operatorSet IAVSDirectoryOperatorSet, strategies []common.Address, timestamp uint32, linear bool) ([]*big.Int, error) {
	return _AVSDirectoryStorage.Contract.GetSlashablePPM(&_AVSDirectoryStorage.CallOpts, operator, operatorSet, strategies, timestamp, linear)
}

// GetSlashablePPM is a free data retrieval call binding the contract method 0x1c699d65.
//
// Solidity: function getSlashablePPM(address operator, (address,uint32) operatorSet, address[] strategies, uint32 timestamp, bool linear) view returns(uint24[])
func (_AVSDirectoryStorage *AVSDirectoryStorageCallerSession) GetSlashablePPM(operator common.Address, operatorSet IAVSDirectoryOperatorSet, strategies []common.Address, timestamp uint32, linear bool) ([]*big.Int, error) {
	return _AVSDirectoryStorage.Contract.GetSlashablePPM(&_AVSDirectoryStorage.CallOpts, operator, operatorSet, strategies, timestamp, linear)
>>>>>>> 49569901 (feat: set allocations (#691))
}

// IsMember is a free data retrieval call binding the contract method 0xda2ff05d.
//
// Solidity: function isMember(address operator, (address,uint32) operatorSet) view returns(bool)
func (_AVSDirectoryStorage *AVSDirectoryStorageCaller) IsMember(opts *bind.CallOpts, operator common.Address, operatorSet IAVSDirectoryOperatorSet) (bool, error) {
	var out []interface{}
	err := _AVSDirectoryStorage.contract.Call(opts, &out, "isMember", operator, operatorSet)
<<<<<<< HEAD
=======
	err := _AVSDirectoryStorage.contract.Call(opts, &out, "isMember", arg0, arg1, arg2)
>>>>>>> ce94473c (feat: operator commission bips (#627))
=======
>>>>>>> 6b2ab7ed (build: bindings)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMember is a free data retrieval call binding the contract method 0xda2ff05d.
//
// Solidity: function isMember(address operator, (address,uint32) operatorSet) view returns(bool)
func (_AVSDirectoryStorage *AVSDirectoryStorageSession) IsMember(operator common.Address, operatorSet IAVSDirectoryOperatorSet) (bool, error) {
	return _AVSDirectoryStorage.Contract.IsMember(&_AVSDirectoryStorage.CallOpts, operator, operatorSet)
}

// IsMember is a free data retrieval call binding the contract method 0xda2ff05d.
//
// Solidity: function isMember(address operator, (address,uint32) operatorSet) view returns(bool)
func (_AVSDirectoryStorage *AVSDirectoryStorageCallerSession) IsMember(operator common.Address, operatorSet IAVSDirectoryOperatorSet) (bool, error) {
	return _AVSDirectoryStorage.Contract.IsMember(&_AVSDirectoryStorage.CallOpts, operator, operatorSet)
}

// IsOperatorSet is a free data retrieval call binding the contract method 0x84d76f7b.
//
// Solidity: function isOperatorSet(address , uint32 ) view returns(bool)
func (_AVSDirectoryStorage *AVSDirectoryStorageCaller) IsOperatorSet(opts *bind.CallOpts, arg0 common.Address, arg1 uint32) (bool, error) {
	var out []interface{}
	err := _AVSDirectoryStorage.contract.Call(opts, &out, "isOperatorSet", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperatorSet is a free data retrieval call binding the contract method 0x84d76f7b.
//
// Solidity: function isOperatorSet(address , uint32 ) view returns(bool)
func (_AVSDirectoryStorage *AVSDirectoryStorageSession) IsOperatorSet(arg0 common.Address, arg1 uint32) (bool, error) {
	return _AVSDirectoryStorage.Contract.IsOperatorSet(&_AVSDirectoryStorage.CallOpts, arg0, arg1)
}

// IsOperatorSet is a free data retrieval call binding the contract method 0x84d76f7b.
//
// Solidity: function isOperatorSet(address , uint32 ) view returns(bool)
func (_AVSDirectoryStorage *AVSDirectoryStorageCallerSession) IsOperatorSet(arg0 common.Address, arg1 uint32) (bool, error) {
	return _AVSDirectoryStorage.Contract.IsOperatorSet(&_AVSDirectoryStorage.CallOpts, arg0, arg1)
}

// IsOperatorSetAVS is a free data retrieval call binding the contract method 0x7673e93a.
//
// Solidity: function isOperatorSetAVS(address ) view returns(bool)
func (_AVSDirectoryStorage *AVSDirectoryStorageCaller) IsOperatorSetAVS(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _AVSDirectoryStorage.contract.Call(opts, &out, "isOperatorSetAVS", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperatorSetAVS is a free data retrieval call binding the contract method 0x7673e93a.
//
// Solidity: function isOperatorSetAVS(address ) view returns(bool)
func (_AVSDirectoryStorage *AVSDirectoryStorageSession) IsOperatorSetAVS(arg0 common.Address) (bool, error) {
	return _AVSDirectoryStorage.Contract.IsOperatorSetAVS(&_AVSDirectoryStorage.CallOpts, arg0)
}

// IsOperatorSetAVS is a free data retrieval call binding the contract method 0x7673e93a.
//
// Solidity: function isOperatorSetAVS(address ) view returns(bool)
func (_AVSDirectoryStorage *AVSDirectoryStorageCallerSession) IsOperatorSetAVS(arg0 common.Address) (bool, error) {
	return _AVSDirectoryStorage.Contract.IsOperatorSetAVS(&_AVSDirectoryStorage.CallOpts, arg0)
}

// IsOperatorSlashable is a free data retrieval call binding the contract method 0x1352c3e6.
//
// Solidity: function isOperatorSlashable(address operator, (address,uint32) operatorSet) view returns(bool)
func (_AVSDirectoryStorage *AVSDirectoryStorageCaller) IsOperatorSlashable(opts *bind.CallOpts, operator common.Address, operatorSet IAVSDirectoryOperatorSet) (bool, error) {
	var out []interface{}
	err := _AVSDirectoryStorage.contract.Call(opts, &out, "isOperatorSlashable", operator, operatorSet)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperatorSlashable is a free data retrieval call binding the contract method 0x1352c3e6.
//
// Solidity: function isOperatorSlashable(address operator, (address,uint32) operatorSet) view returns(bool)
func (_AVSDirectoryStorage *AVSDirectoryStorageSession) IsOperatorSlashable(operator common.Address, operatorSet IAVSDirectoryOperatorSet) (bool, error) {
	return _AVSDirectoryStorage.Contract.IsOperatorSlashable(&_AVSDirectoryStorage.CallOpts, operator, operatorSet)
}

// IsOperatorSlashable is a free data retrieval call binding the contract method 0x1352c3e6.
//
// Solidity: function isOperatorSlashable(address operator, (address,uint32) operatorSet) view returns(bool)
func (_AVSDirectoryStorage *AVSDirectoryStorageCallerSession) IsOperatorSlashable(operator common.Address, operatorSet IAVSDirectoryOperatorSet) (bool, error) {
	return _AVSDirectoryStorage.Contract.IsOperatorSlashable(&_AVSDirectoryStorage.CallOpts, operator, operatorSet)
}

// OperatorSaltIsSpent is a free data retrieval call binding the contract method 0x374823b5.
//
// Solidity: function operatorSaltIsSpent(address , bytes32 ) view returns(bool)
func (_AVSDirectoryStorage *AVSDirectoryStorageCaller) OperatorSaltIsSpent(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (bool, error) {
	var out []interface{}
	err := _AVSDirectoryStorage.contract.Call(opts, &out, "operatorSaltIsSpent", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OperatorSaltIsSpent is a free data retrieval call binding the contract method 0x374823b5.
//
// Solidity: function operatorSaltIsSpent(address , bytes32 ) view returns(bool)
func (_AVSDirectoryStorage *AVSDirectoryStorageSession) OperatorSaltIsSpent(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _AVSDirectoryStorage.Contract.OperatorSaltIsSpent(&_AVSDirectoryStorage.CallOpts, arg0, arg1)
}

// OperatorSaltIsSpent is a free data retrieval call binding the contract method 0x374823b5.
//
// Solidity: function operatorSaltIsSpent(address , bytes32 ) view returns(bool)
func (_AVSDirectoryStorage *AVSDirectoryStorageCallerSession) OperatorSaltIsSpent(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _AVSDirectoryStorage.Contract.OperatorSaltIsSpent(&_AVSDirectoryStorage.CallOpts, arg0, arg1)
}

<<<<<<< HEAD
<<<<<<< HEAD
// OperatorSetMemberAtIndex is a free data retrieval call binding the contract method 0x411d415b.
//
// Solidity: function operatorSetMemberAtIndex((address,uint32) operatorSet, uint256 index) view returns(address)
func (_AVSDirectoryStorage *AVSDirectoryStorageCaller) OperatorSetMemberAtIndex(opts *bind.CallOpts, operatorSet IAVSDirectoryOperatorSet, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AVSDirectoryStorage.contract.Call(opts, &out, "operatorSetMemberAtIndex", operatorSet, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OperatorSetMemberAtIndex is a free data retrieval call binding the contract method 0x411d415b.
//
// Solidity: function operatorSetMemberAtIndex((address,uint32) operatorSet, uint256 index) view returns(address)
func (_AVSDirectoryStorage *AVSDirectoryStorageSession) OperatorSetMemberAtIndex(operatorSet IAVSDirectoryOperatorSet, index *big.Int) (common.Address, error) {
	return _AVSDirectoryStorage.Contract.OperatorSetMemberAtIndex(&_AVSDirectoryStorage.CallOpts, operatorSet, index)
}

// OperatorSetMemberAtIndex is a free data retrieval call binding the contract method 0x411d415b.
//
// Solidity: function operatorSetMemberAtIndex((address,uint32) operatorSet, uint256 index) view returns(address)
func (_AVSDirectoryStorage *AVSDirectoryStorageCallerSession) OperatorSetMemberAtIndex(operatorSet IAVSDirectoryOperatorSet, index *big.Int) (common.Address, error) {
	return _AVSDirectoryStorage.Contract.OperatorSetMemberAtIndex(&_AVSDirectoryStorage.CallOpts, operatorSet, index)
}

// OperatorSetsMemberOfAtIndex is a free data retrieval call binding the contract method 0xb5a768ca.
//
// Solidity: function operatorSetsMemberOfAtIndex(address operator, uint256 index) view returns((address,uint32))
func (_AVSDirectoryStorage *AVSDirectoryStorageCaller) OperatorSetsMemberOfAtIndex(opts *bind.CallOpts, operator common.Address, index *big.Int) (IAVSDirectoryOperatorSet, error) {
	var out []interface{}
	err := _AVSDirectoryStorage.contract.Call(opts, &out, "operatorSetsMemberOfAtIndex", operator, index)

	if err != nil {
		return *new(IAVSDirectoryOperatorSet), err
	}

	out0 := *abi.ConvertType(out[0], new(IAVSDirectoryOperatorSet)).(*IAVSDirectoryOperatorSet)

	return out0, err

}

// OperatorSetsMemberOfAtIndex is a free data retrieval call binding the contract method 0xb5a768ca.
//
// Solidity: function operatorSetsMemberOfAtIndex(address operator, uint256 index) view returns((address,uint32))
func (_AVSDirectoryStorage *AVSDirectoryStorageSession) OperatorSetsMemberOfAtIndex(operator common.Address, index *big.Int) (IAVSDirectoryOperatorSet, error) {
	return _AVSDirectoryStorage.Contract.OperatorSetsMemberOfAtIndex(&_AVSDirectoryStorage.CallOpts, operator, index)
}

// OperatorSetsMemberOfAtIndex is a free data retrieval call binding the contract method 0xb5a768ca.
//
// Solidity: function operatorSetsMemberOfAtIndex(address operator, uint256 index) view returns((address,uint32))
func (_AVSDirectoryStorage *AVSDirectoryStorageCallerSession) OperatorSetsMemberOfAtIndex(operator common.Address, index *big.Int) (IAVSDirectoryOperatorSet, error) {
	return _AVSDirectoryStorage.Contract.OperatorSetsMemberOfAtIndex(&_AVSDirectoryStorage.CallOpts, operator, index)
}

// BecomeOperatorSetAVS is a paid mutator transaction binding the contract method 0xaec205c5.
//
// Solidity: function becomeOperatorSetAVS() returns()
func (_AVSDirectoryStorage *AVSDirectoryStorageTransactor) BecomeOperatorSetAVS(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AVSDirectoryStorage.contract.Transact(opts, "becomeOperatorSetAVS")
}

// BecomeOperatorSetAVS is a paid mutator transaction binding the contract method 0xaec205c5.
//
// Solidity: function becomeOperatorSetAVS() returns()
func (_AVSDirectoryStorage *AVSDirectoryStorageSession) BecomeOperatorSetAVS() (*types.Transaction, error) {
	return _AVSDirectoryStorage.Contract.BecomeOperatorSetAVS(&_AVSDirectoryStorage.TransactOpts)
}

// BecomeOperatorSetAVS is a paid mutator transaction binding the contract method 0xaec205c5.
//
// Solidity: function becomeOperatorSetAVS() returns()
func (_AVSDirectoryStorage *AVSDirectoryStorageTransactorSession) BecomeOperatorSetAVS() (*types.Transaction, error) {
	return _AVSDirectoryStorage.Contract.BecomeOperatorSetAVS(&_AVSDirectoryStorage.TransactOpts)
}

// CancelSalt is a paid mutator transaction binding the contract method 0xec76f442.
//
// Solidity: function cancelSalt(bytes32 salt) returns()
func (_AVSDirectoryStorage *AVSDirectoryStorageTransactor) CancelSalt(opts *bind.TransactOpts, salt [32]byte) (*types.Transaction, error) {
	return _AVSDirectoryStorage.contract.Transact(opts, "cancelSalt", salt)
}

// CancelSalt is a paid mutator transaction binding the contract method 0xec76f442.
//
// Solidity: function cancelSalt(bytes32 salt) returns()
func (_AVSDirectoryStorage *AVSDirectoryStorageSession) CancelSalt(salt [32]byte) (*types.Transaction, error) {
	return _AVSDirectoryStorage.Contract.CancelSalt(&_AVSDirectoryStorage.TransactOpts, salt)
}

// CancelSalt is a paid mutator transaction binding the contract method 0xec76f442.
//
// Solidity: function cancelSalt(bytes32 salt) returns()
func (_AVSDirectoryStorage *AVSDirectoryStorageTransactorSession) CancelSalt(salt [32]byte) (*types.Transaction, error) {
	return _AVSDirectoryStorage.Contract.CancelSalt(&_AVSDirectoryStorage.TransactOpts, salt)
}

// CreateOperatorSets is a paid mutator transaction binding the contract method 0xafe02ed5.
//
// Solidity: function createOperatorSets(uint32[] operatorSetIds) returns()
func (_AVSDirectoryStorage *AVSDirectoryStorageTransactor) CreateOperatorSets(opts *bind.TransactOpts, operatorSetIds []uint32) (*types.Transaction, error) {
	return _AVSDirectoryStorage.contract.Transact(opts, "createOperatorSets", operatorSetIds)
}

// CreateOperatorSets is a paid mutator transaction binding the contract method 0xafe02ed5.
//
// Solidity: function createOperatorSets(uint32[] operatorSetIds) returns()
func (_AVSDirectoryStorage *AVSDirectoryStorageSession) CreateOperatorSets(operatorSetIds []uint32) (*types.Transaction, error) {
	return _AVSDirectoryStorage.Contract.CreateOperatorSets(&_AVSDirectoryStorage.TransactOpts, operatorSetIds)
}

// CreateOperatorSets is a paid mutator transaction binding the contract method 0xafe02ed5.
//
// Solidity: function createOperatorSets(uint32[] operatorSetIds) returns()
func (_AVSDirectoryStorage *AVSDirectoryStorageTransactorSession) CreateOperatorSets(operatorSetIds []uint32) (*types.Transaction, error) {
	return _AVSDirectoryStorage.Contract.CreateOperatorSets(&_AVSDirectoryStorage.TransactOpts, operatorSetIds)
}

=======
>>>>>>> ce94473c (feat: operator commission bips (#627))
// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
=======
// OperatorSetMemberCount is a free data retrieval call binding the contract method 0xdae226b6.
>>>>>>> 6b2ab7ed (build: bindings)
//
// Solidity: function operatorSetMemberCount(address , uint32 ) view returns(uint256)
func (_AVSDirectoryStorage *AVSDirectoryStorageCaller) OperatorSetMemberCount(opts *bind.CallOpts, arg0 common.Address, arg1 uint32) (*big.Int, error) {
	var out []interface{}
	err := _AVSDirectoryStorage.contract.Call(opts, &out, "operatorSetMemberCount", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatorSetMemberCount is a free data retrieval call binding the contract method 0xdae226b6.
//
// Solidity: function operatorSetMemberCount(address , uint32 ) view returns(uint256)
func (_AVSDirectoryStorage *AVSDirectoryStorageSession) OperatorSetMemberCount(arg0 common.Address, arg1 uint32) (*big.Int, error) {
	return _AVSDirectoryStorage.Contract.OperatorSetMemberCount(&_AVSDirectoryStorage.CallOpts, arg0, arg1)
}

// OperatorSetMemberCount is a free data retrieval call binding the contract method 0xdae226b6.
//
// Solidity: function operatorSetMemberCount(address , uint32 ) view returns(uint256)
func (_AVSDirectoryStorage *AVSDirectoryStorageCallerSession) OperatorSetMemberCount(arg0 common.Address, arg1 uint32) (*big.Int, error) {
	return _AVSDirectoryStorage.Contract.OperatorSetMemberCount(&_AVSDirectoryStorage.CallOpts, arg0, arg1)
}

// OperatorSetStatus is a free data retrieval call binding the contract method 0x1e68134e.
//
// Solidity: function operatorSetStatus(address , address , uint32 ) view returns(bool registered, uint32 lastDeregisteredTimestamp)
func (_AVSDirectoryStorage *AVSDirectoryStorageCaller) OperatorSetStatus(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 uint32) (struct {
	Registered                bool
	LastDeregisteredTimestamp uint32
}, error) {
	var out []interface{}
	err := _AVSDirectoryStorage.contract.Call(opts, &out, "operatorSetStatus", arg0, arg1, arg2)

	outstruct := new(struct {
		Registered                bool
		LastDeregisteredTimestamp uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Registered = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.LastDeregisteredTimestamp = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

// OperatorSetStatus is a free data retrieval call binding the contract method 0x1e68134e.
//
// Solidity: function operatorSetStatus(address , address , uint32 ) view returns(bool registered, uint32 lastDeregisteredTimestamp)
func (_AVSDirectoryStorage *AVSDirectoryStorageSession) OperatorSetStatus(arg0 common.Address, arg1 common.Address, arg2 uint32) (struct {
	Registered                bool
	LastDeregisteredTimestamp uint32
}, error) {
	return _AVSDirectoryStorage.Contract.OperatorSetStatus(&_AVSDirectoryStorage.CallOpts, arg0, arg1, arg2)
}

// OperatorSetStatus is a free data retrieval call binding the contract method 0x1e68134e.
//
// Solidity: function operatorSetStatus(address , address , uint32 ) view returns(bool registered, uint32 lastDeregisteredTimestamp)
func (_AVSDirectoryStorage *AVSDirectoryStorageCallerSession) OperatorSetStatus(arg0 common.Address, arg1 common.Address, arg2 uint32) (struct {
	Registered                bool
	LastDeregisteredTimestamp uint32
}, error) {
	return _AVSDirectoryStorage.Contract.OperatorSetStatus(&_AVSDirectoryStorage.CallOpts, arg0, arg1, arg2)
}

// BecomeOperatorSetAVS is a paid mutator transaction binding the contract method 0xaec205c5.
//
// Solidity: function becomeOperatorSetAVS() returns()
func (_AVSDirectoryStorage *AVSDirectoryStorageTransactor) BecomeOperatorSetAVS(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AVSDirectoryStorage.contract.Transact(opts, "becomeOperatorSetAVS")
}

// BecomeOperatorSetAVS is a paid mutator transaction binding the contract method 0xaec205c5.
//
// Solidity: function becomeOperatorSetAVS() returns()
func (_AVSDirectoryStorage *AVSDirectoryStorageSession) BecomeOperatorSetAVS() (*types.Transaction, error) {
	return _AVSDirectoryStorage.Contract.BecomeOperatorSetAVS(&_AVSDirectoryStorage.TransactOpts)
}

// BecomeOperatorSetAVS is a paid mutator transaction binding the contract method 0xaec205c5.
//
// Solidity: function becomeOperatorSetAVS() returns()
func (_AVSDirectoryStorage *AVSDirectoryStorageTransactorSession) BecomeOperatorSetAVS() (*types.Transaction, error) {
	return _AVSDirectoryStorage.Contract.BecomeOperatorSetAVS(&_AVSDirectoryStorage.TransactOpts)
}

// CancelSalt is a paid mutator transaction binding the contract method 0xec76f442.
//
// Solidity: function cancelSalt(bytes32 salt) returns()
func (_AVSDirectoryStorage *AVSDirectoryStorageTransactor) CancelSalt(opts *bind.TransactOpts, salt [32]byte) (*types.Transaction, error) {
	return _AVSDirectoryStorage.contract.Transact(opts, "cancelSalt", salt)
}

// CancelSalt is a paid mutator transaction binding the contract method 0xec76f442.
//
// Solidity: function cancelSalt(bytes32 salt) returns()
func (_AVSDirectoryStorage *AVSDirectoryStorageSession) CancelSalt(salt [32]byte) (*types.Transaction, error) {
	return _AVSDirectoryStorage.Contract.CancelSalt(&_AVSDirectoryStorage.TransactOpts, salt)
}

// CancelSalt is a paid mutator transaction binding the contract method 0xec76f442.
//
// Solidity: function cancelSalt(bytes32 salt) returns()
func (_AVSDirectoryStorage *AVSDirectoryStorageTransactorSession) CancelSalt(salt [32]byte) (*types.Transaction, error) {
	return _AVSDirectoryStorage.Contract.CancelSalt(&_AVSDirectoryStorage.TransactOpts, salt)
}

// CreateOperatorSets is a paid mutator transaction binding the contract method 0xafe02ed5.
//
// Solidity: function createOperatorSets(uint32[] operatorSetIds) returns()
func (_AVSDirectoryStorage *AVSDirectoryStorageTransactor) CreateOperatorSets(opts *bind.TransactOpts, operatorSetIds []uint32) (*types.Transaction, error) {
	return _AVSDirectoryStorage.contract.Transact(opts, "createOperatorSets", operatorSetIds)
}

// CreateOperatorSets is a paid mutator transaction binding the contract method 0xafe02ed5.
//
// Solidity: function createOperatorSets(uint32[] operatorSetIds) returns()
func (_AVSDirectoryStorage *AVSDirectoryStorageSession) CreateOperatorSets(operatorSetIds []uint32) (*types.Transaction, error) {
	return _AVSDirectoryStorage.Contract.CreateOperatorSets(&_AVSDirectoryStorage.TransactOpts, operatorSetIds)
}

// CreateOperatorSets is a paid mutator transaction binding the contract method 0xafe02ed5.
//
// Solidity: function createOperatorSets(uint32[] operatorSetIds) returns()
func (_AVSDirectoryStorage *AVSDirectoryStorageTransactorSession) CreateOperatorSets(operatorSetIds []uint32) (*types.Transaction, error) {
	return _AVSDirectoryStorage.Contract.CreateOperatorSets(&_AVSDirectoryStorage.TransactOpts, operatorSetIds)
}

// DeregisterOperatorFromOperatorSets is a paid mutator transaction binding the contract method 0xc1a8e2c5.
//
// Solidity: function deregisterOperatorFromOperatorSets(address operator, uint32[] operatorSetIds) returns()
func (_AVSDirectoryStorage *AVSDirectoryStorageTransactor) DeregisterOperatorFromOperatorSets(opts *bind.TransactOpts, operator common.Address, operatorSetIds []uint32) (*types.Transaction, error) {
	return _AVSDirectoryStorage.contract.Transact(opts, "deregisterOperatorFromOperatorSets", operator, operatorSetIds)
}

// DeregisterOperatorFromOperatorSets is a paid mutator transaction binding the contract method 0xc1a8e2c5.
//
// Solidity: function deregisterOperatorFromOperatorSets(address operator, uint32[] operatorSetIds) returns()
func (_AVSDirectoryStorage *AVSDirectoryStorageSession) DeregisterOperatorFromOperatorSets(operator common.Address, operatorSetIds []uint32) (*types.Transaction, error) {
	return _AVSDirectoryStorage.Contract.DeregisterOperatorFromOperatorSets(&_AVSDirectoryStorage.TransactOpts, operator, operatorSetIds)
}

// DeregisterOperatorFromOperatorSets is a paid mutator transaction binding the contract method 0xc1a8e2c5.
//
// Solidity: function deregisterOperatorFromOperatorSets(address operator, uint32[] operatorSetIds) returns()
func (_AVSDirectoryStorage *AVSDirectoryStorageTransactorSession) DeregisterOperatorFromOperatorSets(operator common.Address, operatorSetIds []uint32) (*types.Transaction, error) {
	return _AVSDirectoryStorage.Contract.DeregisterOperatorFromOperatorSets(&_AVSDirectoryStorage.TransactOpts, operator, operatorSetIds)
}

// ForceDeregisterFromOperatorSets is a paid mutator transaction binding the contract method 0x3fee332d.
//
// Solidity: function forceDeregisterFromOperatorSets(address operator, address avs, uint32[] operatorSetIds, (bytes,bytes32,uint256) operatorSignature) returns()
func (_AVSDirectoryStorage *AVSDirectoryStorageTransactor) ForceDeregisterFromOperatorSets(opts *bind.TransactOpts, operator common.Address, avs common.Address, operatorSetIds []uint32, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _AVSDirectoryStorage.contract.Transact(opts, "forceDeregisterFromOperatorSets", operator, avs, operatorSetIds, operatorSignature)
}

// ForceDeregisterFromOperatorSets is a paid mutator transaction binding the contract method 0x3fee332d.
//
// Solidity: function forceDeregisterFromOperatorSets(address operator, address avs, uint32[] operatorSetIds, (bytes,bytes32,uint256) operatorSignature) returns()
func (_AVSDirectoryStorage *AVSDirectoryStorageSession) ForceDeregisterFromOperatorSets(operator common.Address, avs common.Address, operatorSetIds []uint32, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _AVSDirectoryStorage.Contract.ForceDeregisterFromOperatorSets(&_AVSDirectoryStorage.TransactOpts, operator, avs, operatorSetIds, operatorSignature)
}

// ForceDeregisterFromOperatorSets is a paid mutator transaction binding the contract method 0x3fee332d.
//
// Solidity: function forceDeregisterFromOperatorSets(address operator, address avs, uint32[] operatorSetIds, (bytes,bytes32,uint256) operatorSignature) returns()
func (_AVSDirectoryStorage *AVSDirectoryStorageTransactorSession) ForceDeregisterFromOperatorSets(operator common.Address, avs common.Address, operatorSetIds []uint32, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _AVSDirectoryStorage.Contract.ForceDeregisterFromOperatorSets(&_AVSDirectoryStorage.TransactOpts, operator, avs, operatorSetIds, operatorSignature)
}

// InitializeAllocationDelay is a paid mutator transaction binding the contract method 0x80fe6da4.
//
// Solidity: function initializeAllocationDelay(address operator, uint32 delay) returns()
func (_AVSDirectoryStorage *AVSDirectoryStorageTransactor) InitializeAllocationDelay(opts *bind.TransactOpts, operator common.Address, delay uint32) (*types.Transaction, error) {
	return _AVSDirectoryStorage.contract.Transact(opts, "initializeAllocationDelay", operator, delay)
}

// InitializeAllocationDelay is a paid mutator transaction binding the contract method 0x80fe6da4.
//
// Solidity: function initializeAllocationDelay(address operator, uint32 delay) returns()
func (_AVSDirectoryStorage *AVSDirectoryStorageSession) InitializeAllocationDelay(operator common.Address, delay uint32) (*types.Transaction, error) {
	return _AVSDirectoryStorage.Contract.InitializeAllocationDelay(&_AVSDirectoryStorage.TransactOpts, operator, delay)
}

// InitializeAllocationDelay is a paid mutator transaction binding the contract method 0x80fe6da4.
//
// Solidity: function initializeAllocationDelay(address operator, uint32 delay) returns()
func (_AVSDirectoryStorage *AVSDirectoryStorageTransactorSession) InitializeAllocationDelay(operator common.Address, delay uint32) (*types.Transaction, error) {
	return _AVSDirectoryStorage.Contract.InitializeAllocationDelay(&_AVSDirectoryStorage.TransactOpts, operator, delay)
}

// MigrateOperatorsToOperatorSets is a paid mutator transaction binding the contract method 0xef2dfa8d.
//
// Solidity: function migrateOperatorsToOperatorSets(address[] operators, uint32[][] operatorSetIds) returns()
func (_AVSDirectoryStorage *AVSDirectoryStorageTransactor) MigrateOperatorsToOperatorSets(opts *bind.TransactOpts, operators []common.Address, operatorSetIds [][]uint32) (*types.Transaction, error) {
	return _AVSDirectoryStorage.contract.Transact(opts, "migrateOperatorsToOperatorSets", operators, operatorSetIds)
}

// MigrateOperatorsToOperatorSets is a paid mutator transaction binding the contract method 0xef2dfa8d.
//
// Solidity: function migrateOperatorsToOperatorSets(address[] operators, uint32[][] operatorSetIds) returns()
func (_AVSDirectoryStorage *AVSDirectoryStorageSession) MigrateOperatorsToOperatorSets(operators []common.Address, operatorSetIds [][]uint32) (*types.Transaction, error) {
	return _AVSDirectoryStorage.Contract.MigrateOperatorsToOperatorSets(&_AVSDirectoryStorage.TransactOpts, operators, operatorSetIds)
}

// MigrateOperatorsToOperatorSets is a paid mutator transaction binding the contract method 0xef2dfa8d.
//
// Solidity: function migrateOperatorsToOperatorSets(address[] operators, uint32[][] operatorSetIds) returns()
func (_AVSDirectoryStorage *AVSDirectoryStorageTransactorSession) MigrateOperatorsToOperatorSets(operators []common.Address, operatorSetIds [][]uint32) (*types.Transaction, error) {
	return _AVSDirectoryStorage.Contract.MigrateOperatorsToOperatorSets(&_AVSDirectoryStorage.TransactOpts, operators, operatorSetIds)
}

// ModifyAllocations is a paid mutator transaction binding the contract method 0xf8e91d16.
//
// Solidity: function modifyAllocations(address operator, (address,uint64,(address,uint32)[],uint64[])[] allocations, (bytes,bytes32,uint256) operatorSignature) returns()
func (_AVSDirectoryStorage *AVSDirectoryStorageTransactor) ModifyAllocations(opts *bind.TransactOpts, operator common.Address, allocations []IAVSDirectoryMagnitudeAllocation, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _AVSDirectoryStorage.contract.Transact(opts, "modifyAllocations", operator, allocations, operatorSignature)
}

// ModifyAllocations is a paid mutator transaction binding the contract method 0xf8e91d16.
//
// Solidity: function modifyAllocations(address operator, (address,uint64,(address,uint32)[],uint64[])[] allocations, (bytes,bytes32,uint256) operatorSignature) returns()
func (_AVSDirectoryStorage *AVSDirectoryStorageSession) ModifyAllocations(operator common.Address, allocations []IAVSDirectoryMagnitudeAllocation, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _AVSDirectoryStorage.Contract.ModifyAllocations(&_AVSDirectoryStorage.TransactOpts, operator, allocations, operatorSignature)
}

// ModifyAllocations is a paid mutator transaction binding the contract method 0xf8e91d16.
//
// Solidity: function modifyAllocations(address operator, (address,uint64,(address,uint32)[],uint64[])[] allocations, (bytes,bytes32,uint256) operatorSignature) returns()
func (_AVSDirectoryStorage *AVSDirectoryStorageTransactorSession) ModifyAllocations(operator common.Address, allocations []IAVSDirectoryMagnitudeAllocation, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _AVSDirectoryStorage.Contract.ModifyAllocations(&_AVSDirectoryStorage.TransactOpts, operator, allocations, operatorSignature)
}

// RegisterOperatorToOperatorSets is a paid mutator transaction binding the contract method 0x1e2199e2.
//
// Solidity: function registerOperatorToOperatorSets(address operator, uint32[] operatorSetIds, (bytes,bytes32,uint256) operatorSignature) returns()
func (_AVSDirectoryStorage *AVSDirectoryStorageTransactor) RegisterOperatorToOperatorSets(opts *bind.TransactOpts, operator common.Address, operatorSetIds []uint32, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _AVSDirectoryStorage.contract.Transact(opts, "registerOperatorToOperatorSets", operator, operatorSetIds, operatorSignature)
}

// RegisterOperatorToOperatorSets is a paid mutator transaction binding the contract method 0x1e2199e2.
//
// Solidity: function registerOperatorToOperatorSets(address operator, uint32[] operatorSetIds, (bytes,bytes32,uint256) operatorSignature) returns()
func (_AVSDirectoryStorage *AVSDirectoryStorageSession) RegisterOperatorToOperatorSets(operator common.Address, operatorSetIds []uint32, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _AVSDirectoryStorage.Contract.RegisterOperatorToOperatorSets(&_AVSDirectoryStorage.TransactOpts, operator, operatorSetIds, operatorSignature)
}

// RegisterOperatorToOperatorSets is a paid mutator transaction binding the contract method 0x1e2199e2.
//
// Solidity: function registerOperatorToOperatorSets(address operator, uint32[] operatorSetIds, (bytes,bytes32,uint256) operatorSignature) returns()
func (_AVSDirectoryStorage *AVSDirectoryStorageTransactorSession) RegisterOperatorToOperatorSets(operator common.Address, operatorSetIds []uint32, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _AVSDirectoryStorage.Contract.RegisterOperatorToOperatorSets(&_AVSDirectoryStorage.TransactOpts, operator, operatorSetIds, operatorSignature)
}

// SlashOperator is a paid mutator transaction binding the contract method 0xbd74a06c.
//
// Solidity: function slashOperator(address operator, uint32 operatorSetId, address[] strategies, uint16 bipsToSlash) returns()
func (_AVSDirectoryStorage *AVSDirectoryStorageTransactor) SlashOperator(opts *bind.TransactOpts, operator common.Address, operatorSetId uint32, strategies []common.Address, bipsToSlash uint16) (*types.Transaction, error) {
	return _AVSDirectoryStorage.contract.Transact(opts, "slashOperator", operator, operatorSetId, strategies, bipsToSlash)
}

// SlashOperator is a paid mutator transaction binding the contract method 0xbd74a06c.
//
// Solidity: function slashOperator(address operator, uint32 operatorSetId, address[] strategies, uint16 bipsToSlash) returns()
func (_AVSDirectoryStorage *AVSDirectoryStorageSession) SlashOperator(operator common.Address, operatorSetId uint32, strategies []common.Address, bipsToSlash uint16) (*types.Transaction, error) {
	return _AVSDirectoryStorage.Contract.SlashOperator(&_AVSDirectoryStorage.TransactOpts, operator, operatorSetId, strategies, bipsToSlash)
}

// SlashOperator is a paid mutator transaction binding the contract method 0xbd74a06c.
//
// Solidity: function slashOperator(address operator, uint32 operatorSetId, address[] strategies, uint16 bipsToSlash) returns()
func (_AVSDirectoryStorage *AVSDirectoryStorageTransactorSession) SlashOperator(operator common.Address, operatorSetId uint32, strategies []common.Address, bipsToSlash uint16) (*types.Transaction, error) {
	return _AVSDirectoryStorage.Contract.SlashOperator(&_AVSDirectoryStorage.TransactOpts, operator, operatorSetId, strategies, bipsToSlash)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string metadataURI) returns()
func (_AVSDirectoryStorage *AVSDirectoryStorageTransactor) UpdateAVSMetadataURI(opts *bind.TransactOpts, metadataURI string) (*types.Transaction, error) {
	return _AVSDirectoryStorage.contract.Transact(opts, "updateAVSMetadataURI", metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string metadataURI) returns()
func (_AVSDirectoryStorage *AVSDirectoryStorageSession) UpdateAVSMetadataURI(metadataURI string) (*types.Transaction, error) {
	return _AVSDirectoryStorage.Contract.UpdateAVSMetadataURI(&_AVSDirectoryStorage.TransactOpts, metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string metadataURI) returns()
func (_AVSDirectoryStorage *AVSDirectoryStorageTransactorSession) UpdateAVSMetadataURI(metadataURI string) (*types.Transaction, error) {
	return _AVSDirectoryStorage.Contract.UpdateAVSMetadataURI(&_AVSDirectoryStorage.TransactOpts, metadataURI)
}

// UpdateFreeMagnitude is a paid mutator transaction binding the contract method 0x84b78fc2.
//
// Solidity: function updateFreeMagnitude(address operator, address[] strategies, uint8[] numToComplete) returns()
func (_AVSDirectoryStorage *AVSDirectoryStorageTransactor) UpdateFreeMagnitude(opts *bind.TransactOpts, operator common.Address, strategies []common.Address, numToComplete []uint8) (*types.Transaction, error) {
	return _AVSDirectoryStorage.contract.Transact(opts, "updateFreeMagnitude", operator, strategies, numToComplete)
}

// UpdateFreeMagnitude is a paid mutator transaction binding the contract method 0x84b78fc2.
//
// Solidity: function updateFreeMagnitude(address operator, address[] strategies, uint8[] numToComplete) returns()
func (_AVSDirectoryStorage *AVSDirectoryStorageSession) UpdateFreeMagnitude(operator common.Address, strategies []common.Address, numToComplete []uint8) (*types.Transaction, error) {
	return _AVSDirectoryStorage.Contract.UpdateFreeMagnitude(&_AVSDirectoryStorage.TransactOpts, operator, strategies, numToComplete)
}

// UpdateFreeMagnitude is a paid mutator transaction binding the contract method 0x84b78fc2.
//
// Solidity: function updateFreeMagnitude(address operator, address[] strategies, uint8[] numToComplete) returns()
func (_AVSDirectoryStorage *AVSDirectoryStorageTransactorSession) UpdateFreeMagnitude(operator common.Address, strategies []common.Address, numToComplete []uint8) (*types.Transaction, error) {
	return _AVSDirectoryStorage.Contract.UpdateFreeMagnitude(&_AVSDirectoryStorage.TransactOpts, operator, strategies, numToComplete)
}

// AVSDirectoryStorageAVSMetadataURIUpdatedIterator is returned from FilterAVSMetadataURIUpdated and is used to iterate over the raw logs and unpacked data for AVSMetadataURIUpdated events raised by the AVSDirectoryStorage contract.
type AVSDirectoryStorageAVSMetadataURIUpdatedIterator struct {
	Event *AVSDirectoryStorageAVSMetadataURIUpdated // Event containing the contract specifics and raw log

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
func (it *AVSDirectoryStorageAVSMetadataURIUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSDirectoryStorageAVSMetadataURIUpdated)
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
		it.Event = new(AVSDirectoryStorageAVSMetadataURIUpdated)
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
func (it *AVSDirectoryStorageAVSMetadataURIUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSDirectoryStorageAVSMetadataURIUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSDirectoryStorageAVSMetadataURIUpdated represents a AVSMetadataURIUpdated event raised by the AVSDirectoryStorage contract.
type AVSDirectoryStorageAVSMetadataURIUpdated struct {
	Avs         common.Address
	MetadataURI string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAVSMetadataURIUpdated is a free log retrieval operation binding the contract event 0xa89c1dc243d8908a96dd84944bcc97d6bc6ac00dd78e20621576be6a3c943713.
//
// Solidity: event AVSMetadataURIUpdated(address indexed avs, string metadataURI)
func (_AVSDirectoryStorage *AVSDirectoryStorageFilterer) FilterAVSMetadataURIUpdated(opts *bind.FilterOpts, avs []common.Address) (*AVSDirectoryStorageAVSMetadataURIUpdatedIterator, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _AVSDirectoryStorage.contract.FilterLogs(opts, "AVSMetadataURIUpdated", avsRule)
	if err != nil {
		return nil, err
	}
	return &AVSDirectoryStorageAVSMetadataURIUpdatedIterator{contract: _AVSDirectoryStorage.contract, event: "AVSMetadataURIUpdated", logs: logs, sub: sub}, nil
}

// WatchAVSMetadataURIUpdated is a free log subscription operation binding the contract event 0xa89c1dc243d8908a96dd84944bcc97d6bc6ac00dd78e20621576be6a3c943713.
//
// Solidity: event AVSMetadataURIUpdated(address indexed avs, string metadataURI)
func (_AVSDirectoryStorage *AVSDirectoryStorageFilterer) WatchAVSMetadataURIUpdated(opts *bind.WatchOpts, sink chan<- *AVSDirectoryStorageAVSMetadataURIUpdated, avs []common.Address) (event.Subscription, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _AVSDirectoryStorage.contract.WatchLogs(opts, "AVSMetadataURIUpdated", avsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSDirectoryStorageAVSMetadataURIUpdated)
				if err := _AVSDirectoryStorage.contract.UnpackLog(event, "AVSMetadataURIUpdated", log); err != nil {
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

// ParseAVSMetadataURIUpdated is a log parse operation binding the contract event 0xa89c1dc243d8908a96dd84944bcc97d6bc6ac00dd78e20621576be6a3c943713.
//
// Solidity: event AVSMetadataURIUpdated(address indexed avs, string metadataURI)
func (_AVSDirectoryStorage *AVSDirectoryStorageFilterer) ParseAVSMetadataURIUpdated(log types.Log) (*AVSDirectoryStorageAVSMetadataURIUpdated, error) {
	event := new(AVSDirectoryStorageAVSMetadataURIUpdated)
	if err := _AVSDirectoryStorage.contract.UnpackLog(event, "AVSMetadataURIUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSDirectoryStorageAVSMigratedToOperatorSetsIterator is returned from FilterAVSMigratedToOperatorSets and is used to iterate over the raw logs and unpacked data for AVSMigratedToOperatorSets events raised by the AVSDirectoryStorage contract.
type AVSDirectoryStorageAVSMigratedToOperatorSetsIterator struct {
	Event *AVSDirectoryStorageAVSMigratedToOperatorSets // Event containing the contract specifics and raw log

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
func (it *AVSDirectoryStorageAVSMigratedToOperatorSetsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSDirectoryStorageAVSMigratedToOperatorSets)
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
		it.Event = new(AVSDirectoryStorageAVSMigratedToOperatorSets)
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
func (it *AVSDirectoryStorageAVSMigratedToOperatorSetsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSDirectoryStorageAVSMigratedToOperatorSetsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSDirectoryStorageAVSMigratedToOperatorSets represents a AVSMigratedToOperatorSets event raised by the AVSDirectoryStorage contract.
type AVSDirectoryStorageAVSMigratedToOperatorSets struct {
	Avs common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAVSMigratedToOperatorSets is a free log retrieval operation binding the contract event 0x702b0c1f6cb1cf511aaa81f72bc05a215bb3497632d72c690c822b044ab494bf.
//
// Solidity: event AVSMigratedToOperatorSets(address indexed avs)
func (_AVSDirectoryStorage *AVSDirectoryStorageFilterer) FilterAVSMigratedToOperatorSets(opts *bind.FilterOpts, avs []common.Address) (*AVSDirectoryStorageAVSMigratedToOperatorSetsIterator, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _AVSDirectoryStorage.contract.FilterLogs(opts, "AVSMigratedToOperatorSets", avsRule)
	if err != nil {
		return nil, err
	}
	return &AVSDirectoryStorageAVSMigratedToOperatorSetsIterator{contract: _AVSDirectoryStorage.contract, event: "AVSMigratedToOperatorSets", logs: logs, sub: sub}, nil
}

// WatchAVSMigratedToOperatorSets is a free log subscription operation binding the contract event 0x702b0c1f6cb1cf511aaa81f72bc05a215bb3497632d72c690c822b044ab494bf.
//
// Solidity: event AVSMigratedToOperatorSets(address indexed avs)
func (_AVSDirectoryStorage *AVSDirectoryStorageFilterer) WatchAVSMigratedToOperatorSets(opts *bind.WatchOpts, sink chan<- *AVSDirectoryStorageAVSMigratedToOperatorSets, avs []common.Address) (event.Subscription, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _AVSDirectoryStorage.contract.WatchLogs(opts, "AVSMigratedToOperatorSets", avsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSDirectoryStorageAVSMigratedToOperatorSets)
				if err := _AVSDirectoryStorage.contract.UnpackLog(event, "AVSMigratedToOperatorSets", log); err != nil {
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

// ParseAVSMigratedToOperatorSets is a log parse operation binding the contract event 0x702b0c1f6cb1cf511aaa81f72bc05a215bb3497632d72c690c822b044ab494bf.
//
// Solidity: event AVSMigratedToOperatorSets(address indexed avs)
func (_AVSDirectoryStorage *AVSDirectoryStorageFilterer) ParseAVSMigratedToOperatorSets(log types.Log) (*AVSDirectoryStorageAVSMigratedToOperatorSets, error) {
	event := new(AVSDirectoryStorageAVSMigratedToOperatorSets)
	if err := _AVSDirectoryStorage.contract.UnpackLog(event, "AVSMigratedToOperatorSets", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSDirectoryStorageMagnitudeAllocatedIterator is returned from FilterMagnitudeAllocated and is used to iterate over the raw logs and unpacked data for MagnitudeAllocated events raised by the AVSDirectoryStorage contract.
type AVSDirectoryStorageMagnitudeAllocatedIterator struct {
	Event *AVSDirectoryStorageMagnitudeAllocated // Event containing the contract specifics and raw log

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
func (it *AVSDirectoryStorageMagnitudeAllocatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSDirectoryStorageMagnitudeAllocated)
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
		it.Event = new(AVSDirectoryStorageMagnitudeAllocated)
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
func (it *AVSDirectoryStorageMagnitudeAllocatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSDirectoryStorageMagnitudeAllocatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSDirectoryStorageMagnitudeAllocated represents a MagnitudeAllocated event raised by the AVSDirectoryStorage contract.
type AVSDirectoryStorageMagnitudeAllocated struct {
	Operator            common.Address
	Strategy            common.Address
	OperatorSet         IAVSDirectoryOperatorSet
	MagnitudeToAllocate uint64
	EffectTimestamp     uint32
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterMagnitudeAllocated is a free log retrieval operation binding the contract event 0x6d7d0079582cb2c5e70d4135b37f36711415ee6c260778b716bd65e026eb4f1a.
//
// Solidity: event MagnitudeAllocated(address operator, address strategy, (address,uint32) operatorSet, uint64 magnitudeToAllocate, uint32 effectTimestamp)
func (_AVSDirectoryStorage *AVSDirectoryStorageFilterer) FilterMagnitudeAllocated(opts *bind.FilterOpts) (*AVSDirectoryStorageMagnitudeAllocatedIterator, error) {

	logs, sub, err := _AVSDirectoryStorage.contract.FilterLogs(opts, "MagnitudeAllocated")
	if err != nil {
		return nil, err
	}
	return &AVSDirectoryStorageMagnitudeAllocatedIterator{contract: _AVSDirectoryStorage.contract, event: "MagnitudeAllocated", logs: logs, sub: sub}, nil
}

// WatchMagnitudeAllocated is a free log subscription operation binding the contract event 0x6d7d0079582cb2c5e70d4135b37f36711415ee6c260778b716bd65e026eb4f1a.
//
// Solidity: event MagnitudeAllocated(address operator, address strategy, (address,uint32) operatorSet, uint64 magnitudeToAllocate, uint32 effectTimestamp)
func (_AVSDirectoryStorage *AVSDirectoryStorageFilterer) WatchMagnitudeAllocated(opts *bind.WatchOpts, sink chan<- *AVSDirectoryStorageMagnitudeAllocated) (event.Subscription, error) {

	logs, sub, err := _AVSDirectoryStorage.contract.WatchLogs(opts, "MagnitudeAllocated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSDirectoryStorageMagnitudeAllocated)
				if err := _AVSDirectoryStorage.contract.UnpackLog(event, "MagnitudeAllocated", log); err != nil {
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
func (_AVSDirectoryStorage *AVSDirectoryStorageFilterer) ParseMagnitudeAllocated(log types.Log) (*AVSDirectoryStorageMagnitudeAllocated, error) {
	event := new(AVSDirectoryStorageMagnitudeAllocated)
	if err := _AVSDirectoryStorage.contract.UnpackLog(event, "MagnitudeAllocated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSDirectoryStorageMagnitudeDeallocationCompletedIterator is returned from FilterMagnitudeDeallocationCompleted and is used to iterate over the raw logs and unpacked data for MagnitudeDeallocationCompleted events raised by the AVSDirectoryStorage contract.
type AVSDirectoryStorageMagnitudeDeallocationCompletedIterator struct {
	Event *AVSDirectoryStorageMagnitudeDeallocationCompleted // Event containing the contract specifics and raw log

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
func (it *AVSDirectoryStorageMagnitudeDeallocationCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSDirectoryStorageMagnitudeDeallocationCompleted)
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
		it.Event = new(AVSDirectoryStorageMagnitudeDeallocationCompleted)
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
func (it *AVSDirectoryStorageMagnitudeDeallocationCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSDirectoryStorageMagnitudeDeallocationCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSDirectoryStorageMagnitudeDeallocationCompleted represents a MagnitudeDeallocationCompleted event raised by the AVSDirectoryStorage contract.
type AVSDirectoryStorageMagnitudeDeallocationCompleted struct {
	Operator           common.Address
	Strategy           common.Address
	OperatorSet        IAVSDirectoryOperatorSet
	FreeMagnitudeAdded uint64
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMagnitudeDeallocationCompleted is a free log retrieval operation binding the contract event 0x1e5c8e13c62c31d6252ac205e592477d643c7e95831d5b46d99a3c60c2fad8db.
//
// Solidity: event MagnitudeDeallocationCompleted(address operator, address strategy, (address,uint32) operatorSet, uint64 freeMagnitudeAdded)
func (_AVSDirectoryStorage *AVSDirectoryStorageFilterer) FilterMagnitudeDeallocationCompleted(opts *bind.FilterOpts) (*AVSDirectoryStorageMagnitudeDeallocationCompletedIterator, error) {

	logs, sub, err := _AVSDirectoryStorage.contract.FilterLogs(opts, "MagnitudeDeallocationCompleted")
	if err != nil {
		return nil, err
	}
	return &AVSDirectoryStorageMagnitudeDeallocationCompletedIterator{contract: _AVSDirectoryStorage.contract, event: "MagnitudeDeallocationCompleted", logs: logs, sub: sub}, nil
}

// WatchMagnitudeDeallocationCompleted is a free log subscription operation binding the contract event 0x1e5c8e13c62c31d6252ac205e592477d643c7e95831d5b46d99a3c60c2fad8db.
//
// Solidity: event MagnitudeDeallocationCompleted(address operator, address strategy, (address,uint32) operatorSet, uint64 freeMagnitudeAdded)
func (_AVSDirectoryStorage *AVSDirectoryStorageFilterer) WatchMagnitudeDeallocationCompleted(opts *bind.WatchOpts, sink chan<- *AVSDirectoryStorageMagnitudeDeallocationCompleted) (event.Subscription, error) {

	logs, sub, err := _AVSDirectoryStorage.contract.WatchLogs(opts, "MagnitudeDeallocationCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSDirectoryStorageMagnitudeDeallocationCompleted)
				if err := _AVSDirectoryStorage.contract.UnpackLog(event, "MagnitudeDeallocationCompleted", log); err != nil {
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
func (_AVSDirectoryStorage *AVSDirectoryStorageFilterer) ParseMagnitudeDeallocationCompleted(log types.Log) (*AVSDirectoryStorageMagnitudeDeallocationCompleted, error) {
	event := new(AVSDirectoryStorageMagnitudeDeallocationCompleted)
	if err := _AVSDirectoryStorage.contract.UnpackLog(event, "MagnitudeDeallocationCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSDirectoryStorageMagnitudeQueueDeallocatedIterator is returned from FilterMagnitudeQueueDeallocated and is used to iterate over the raw logs and unpacked data for MagnitudeQueueDeallocated events raised by the AVSDirectoryStorage contract.
type AVSDirectoryStorageMagnitudeQueueDeallocatedIterator struct {
	Event *AVSDirectoryStorageMagnitudeQueueDeallocated // Event containing the contract specifics and raw log

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
func (it *AVSDirectoryStorageMagnitudeQueueDeallocatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSDirectoryStorageMagnitudeQueueDeallocated)
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
		it.Event = new(AVSDirectoryStorageMagnitudeQueueDeallocated)
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
func (it *AVSDirectoryStorageMagnitudeQueueDeallocatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSDirectoryStorageMagnitudeQueueDeallocatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSDirectoryStorageMagnitudeQueueDeallocated represents a MagnitudeQueueDeallocated event raised by the AVSDirectoryStorage contract.
type AVSDirectoryStorageMagnitudeQueueDeallocated struct {
	Operator              common.Address
	Strategy              common.Address
	OperatorSet           IAVSDirectoryOperatorSet
	MagnitudeToDeallocate uint64
	CompletableTimestamp  uint32
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterMagnitudeQueueDeallocated is a free log retrieval operation binding the contract event 0x2e68db1fe51107d7e451ae268d1631796989ab9d7925054e9b247854cb5be950.
//
// Solidity: event MagnitudeQueueDeallocated(address operator, address strategy, (address,uint32) operatorSet, uint64 magnitudeToDeallocate, uint32 completableTimestamp)
func (_AVSDirectoryStorage *AVSDirectoryStorageFilterer) FilterMagnitudeQueueDeallocated(opts *bind.FilterOpts) (*AVSDirectoryStorageMagnitudeQueueDeallocatedIterator, error) {

	logs, sub, err := _AVSDirectoryStorage.contract.FilterLogs(opts, "MagnitudeQueueDeallocated")
	if err != nil {
		return nil, err
	}
	return &AVSDirectoryStorageMagnitudeQueueDeallocatedIterator{contract: _AVSDirectoryStorage.contract, event: "MagnitudeQueueDeallocated", logs: logs, sub: sub}, nil
}

// WatchMagnitudeQueueDeallocated is a free log subscription operation binding the contract event 0x2e68db1fe51107d7e451ae268d1631796989ab9d7925054e9b247854cb5be950.
//
// Solidity: event MagnitudeQueueDeallocated(address operator, address strategy, (address,uint32) operatorSet, uint64 magnitudeToDeallocate, uint32 completableTimestamp)
func (_AVSDirectoryStorage *AVSDirectoryStorageFilterer) WatchMagnitudeQueueDeallocated(opts *bind.WatchOpts, sink chan<- *AVSDirectoryStorageMagnitudeQueueDeallocated) (event.Subscription, error) {

	logs, sub, err := _AVSDirectoryStorage.contract.WatchLogs(opts, "MagnitudeQueueDeallocated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSDirectoryStorageMagnitudeQueueDeallocated)
				if err := _AVSDirectoryStorage.contract.UnpackLog(event, "MagnitudeQueueDeallocated", log); err != nil {
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
func (_AVSDirectoryStorage *AVSDirectoryStorageFilterer) ParseMagnitudeQueueDeallocated(log types.Log) (*AVSDirectoryStorageMagnitudeQueueDeallocated, error) {
	event := new(AVSDirectoryStorageMagnitudeQueueDeallocated)
	if err := _AVSDirectoryStorage.contract.UnpackLog(event, "MagnitudeQueueDeallocated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSDirectoryStorageOperatorAVSRegistrationStatusUpdatedIterator is returned from FilterOperatorAVSRegistrationStatusUpdated and is used to iterate over the raw logs and unpacked data for OperatorAVSRegistrationStatusUpdated events raised by the AVSDirectoryStorage contract.
type AVSDirectoryStorageOperatorAVSRegistrationStatusUpdatedIterator struct {
	Event *AVSDirectoryStorageOperatorAVSRegistrationStatusUpdated // Event containing the contract specifics and raw log

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
func (it *AVSDirectoryStorageOperatorAVSRegistrationStatusUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSDirectoryStorageOperatorAVSRegistrationStatusUpdated)
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
		it.Event = new(AVSDirectoryStorageOperatorAVSRegistrationStatusUpdated)
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
func (it *AVSDirectoryStorageOperatorAVSRegistrationStatusUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSDirectoryStorageOperatorAVSRegistrationStatusUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSDirectoryStorageOperatorAVSRegistrationStatusUpdated represents a OperatorAVSRegistrationStatusUpdated event raised by the AVSDirectoryStorage contract.
type AVSDirectoryStorageOperatorAVSRegistrationStatusUpdated struct {
	Operator common.Address
	Avs      common.Address
	Status   uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorAVSRegistrationStatusUpdated is a free log retrieval operation binding the contract event 0xf0952b1c65271d819d39983d2abb044b9cace59bcc4d4dd389f586ebdcb15b41.
//
// Solidity: event OperatorAVSRegistrationStatusUpdated(address indexed operator, address indexed avs, uint8 status)
func (_AVSDirectoryStorage *AVSDirectoryStorageFilterer) FilterOperatorAVSRegistrationStatusUpdated(opts *bind.FilterOpts, operator []common.Address, avs []common.Address) (*AVSDirectoryStorageOperatorAVSRegistrationStatusUpdatedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _AVSDirectoryStorage.contract.FilterLogs(opts, "OperatorAVSRegistrationStatusUpdated", operatorRule, avsRule)
	if err != nil {
		return nil, err
	}
	return &AVSDirectoryStorageOperatorAVSRegistrationStatusUpdatedIterator{contract: _AVSDirectoryStorage.contract, event: "OperatorAVSRegistrationStatusUpdated", logs: logs, sub: sub}, nil
}

// WatchOperatorAVSRegistrationStatusUpdated is a free log subscription operation binding the contract event 0xf0952b1c65271d819d39983d2abb044b9cace59bcc4d4dd389f586ebdcb15b41.
//
// Solidity: event OperatorAVSRegistrationStatusUpdated(address indexed operator, address indexed avs, uint8 status)
func (_AVSDirectoryStorage *AVSDirectoryStorageFilterer) WatchOperatorAVSRegistrationStatusUpdated(opts *bind.WatchOpts, sink chan<- *AVSDirectoryStorageOperatorAVSRegistrationStatusUpdated, operator []common.Address, avs []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _AVSDirectoryStorage.contract.WatchLogs(opts, "OperatorAVSRegistrationStatusUpdated", operatorRule, avsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSDirectoryStorageOperatorAVSRegistrationStatusUpdated)
				if err := _AVSDirectoryStorage.contract.UnpackLog(event, "OperatorAVSRegistrationStatusUpdated", log); err != nil {
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

// ParseOperatorAVSRegistrationStatusUpdated is a log parse operation binding the contract event 0xf0952b1c65271d819d39983d2abb044b9cace59bcc4d4dd389f586ebdcb15b41.
//
// Solidity: event OperatorAVSRegistrationStatusUpdated(address indexed operator, address indexed avs, uint8 status)
func (_AVSDirectoryStorage *AVSDirectoryStorageFilterer) ParseOperatorAVSRegistrationStatusUpdated(log types.Log) (*AVSDirectoryStorageOperatorAVSRegistrationStatusUpdated, error) {
	event := new(AVSDirectoryStorageOperatorAVSRegistrationStatusUpdated)
	if err := _AVSDirectoryStorage.contract.UnpackLog(event, "OperatorAVSRegistrationStatusUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSDirectoryStorageOperatorAddedToOperatorSetIterator is returned from FilterOperatorAddedToOperatorSet and is used to iterate over the raw logs and unpacked data for OperatorAddedToOperatorSet events raised by the AVSDirectoryStorage contract.
type AVSDirectoryStorageOperatorAddedToOperatorSetIterator struct {
	Event *AVSDirectoryStorageOperatorAddedToOperatorSet // Event containing the contract specifics and raw log

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
func (it *AVSDirectoryStorageOperatorAddedToOperatorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSDirectoryStorageOperatorAddedToOperatorSet)
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
		it.Event = new(AVSDirectoryStorageOperatorAddedToOperatorSet)
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
func (it *AVSDirectoryStorageOperatorAddedToOperatorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSDirectoryStorageOperatorAddedToOperatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSDirectoryStorageOperatorAddedToOperatorSet represents a OperatorAddedToOperatorSet event raised by the AVSDirectoryStorage contract.
type AVSDirectoryStorageOperatorAddedToOperatorSet struct {
	Operator    common.Address
	OperatorSet IAVSDirectoryOperatorSet
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorAddedToOperatorSet is a free log retrieval operation binding the contract event 0x43232edf9071753d2321e5fa7e018363ee248e5f2142e6c08edd3265bfb4895e.
//
// Solidity: event OperatorAddedToOperatorSet(address indexed operator, (address,uint32) operatorSet)
func (_AVSDirectoryStorage *AVSDirectoryStorageFilterer) FilterOperatorAddedToOperatorSet(opts *bind.FilterOpts, operator []common.Address) (*AVSDirectoryStorageOperatorAddedToOperatorSetIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _AVSDirectoryStorage.contract.FilterLogs(opts, "OperatorAddedToOperatorSet", operatorRule)
	if err != nil {
		return nil, err
	}
	return &AVSDirectoryStorageOperatorAddedToOperatorSetIterator{contract: _AVSDirectoryStorage.contract, event: "OperatorAddedToOperatorSet", logs: logs, sub: sub}, nil
}

// WatchOperatorAddedToOperatorSet is a free log subscription operation binding the contract event 0x43232edf9071753d2321e5fa7e018363ee248e5f2142e6c08edd3265bfb4895e.
//
// Solidity: event OperatorAddedToOperatorSet(address indexed operator, (address,uint32) operatorSet)
func (_AVSDirectoryStorage *AVSDirectoryStorageFilterer) WatchOperatorAddedToOperatorSet(opts *bind.WatchOpts, sink chan<- *AVSDirectoryStorageOperatorAddedToOperatorSet, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _AVSDirectoryStorage.contract.WatchLogs(opts, "OperatorAddedToOperatorSet", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSDirectoryStorageOperatorAddedToOperatorSet)
				if err := _AVSDirectoryStorage.contract.UnpackLog(event, "OperatorAddedToOperatorSet", log); err != nil {
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

// ParseOperatorAddedToOperatorSet is a log parse operation binding the contract event 0x43232edf9071753d2321e5fa7e018363ee248e5f2142e6c08edd3265bfb4895e.
//
// Solidity: event OperatorAddedToOperatorSet(address indexed operator, (address,uint32) operatorSet)
func (_AVSDirectoryStorage *AVSDirectoryStorageFilterer) ParseOperatorAddedToOperatorSet(log types.Log) (*AVSDirectoryStorageOperatorAddedToOperatorSet, error) {
	event := new(AVSDirectoryStorageOperatorAddedToOperatorSet)
	if err := _AVSDirectoryStorage.contract.UnpackLog(event, "OperatorAddedToOperatorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSDirectoryStorageOperatorMigratedToOperatorSetsIterator is returned from FilterOperatorMigratedToOperatorSets and is used to iterate over the raw logs and unpacked data for OperatorMigratedToOperatorSets events raised by the AVSDirectoryStorage contract.
type AVSDirectoryStorageOperatorMigratedToOperatorSetsIterator struct {
	Event *AVSDirectoryStorageOperatorMigratedToOperatorSets // Event containing the contract specifics and raw log

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
func (it *AVSDirectoryStorageOperatorMigratedToOperatorSetsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSDirectoryStorageOperatorMigratedToOperatorSets)
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
		it.Event = new(AVSDirectoryStorageOperatorMigratedToOperatorSets)
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
func (it *AVSDirectoryStorageOperatorMigratedToOperatorSetsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSDirectoryStorageOperatorMigratedToOperatorSetsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSDirectoryStorageOperatorMigratedToOperatorSets represents a OperatorMigratedToOperatorSets event raised by the AVSDirectoryStorage contract.
type AVSDirectoryStorageOperatorMigratedToOperatorSets struct {
	Operator       common.Address
	Avs            common.Address
	OperatorSetIds []uint32
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOperatorMigratedToOperatorSets is a free log retrieval operation binding the contract event 0x54f33cfdd1ca703d795986b986fd47d742eab1904ecd2a5fdb8d6595e5904a01.
//
// Solidity: event OperatorMigratedToOperatorSets(address indexed operator, address indexed avs, uint32[] operatorSetIds)
func (_AVSDirectoryStorage *AVSDirectoryStorageFilterer) FilterOperatorMigratedToOperatorSets(opts *bind.FilterOpts, operator []common.Address, avs []common.Address) (*AVSDirectoryStorageOperatorMigratedToOperatorSetsIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _AVSDirectoryStorage.contract.FilterLogs(opts, "OperatorMigratedToOperatorSets", operatorRule, avsRule)
	if err != nil {
		return nil, err
	}
	return &AVSDirectoryStorageOperatorMigratedToOperatorSetsIterator{contract: _AVSDirectoryStorage.contract, event: "OperatorMigratedToOperatorSets", logs: logs, sub: sub}, nil
}

// WatchOperatorMigratedToOperatorSets is a free log subscription operation binding the contract event 0x54f33cfdd1ca703d795986b986fd47d742eab1904ecd2a5fdb8d6595e5904a01.
//
// Solidity: event OperatorMigratedToOperatorSets(address indexed operator, address indexed avs, uint32[] operatorSetIds)
func (_AVSDirectoryStorage *AVSDirectoryStorageFilterer) WatchOperatorMigratedToOperatorSets(opts *bind.WatchOpts, sink chan<- *AVSDirectoryStorageOperatorMigratedToOperatorSets, operator []common.Address, avs []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _AVSDirectoryStorage.contract.WatchLogs(opts, "OperatorMigratedToOperatorSets", operatorRule, avsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSDirectoryStorageOperatorMigratedToOperatorSets)
				if err := _AVSDirectoryStorage.contract.UnpackLog(event, "OperatorMigratedToOperatorSets", log); err != nil {
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

// ParseOperatorMigratedToOperatorSets is a log parse operation binding the contract event 0x54f33cfdd1ca703d795986b986fd47d742eab1904ecd2a5fdb8d6595e5904a01.
//
// Solidity: event OperatorMigratedToOperatorSets(address indexed operator, address indexed avs, uint32[] operatorSetIds)
func (_AVSDirectoryStorage *AVSDirectoryStorageFilterer) ParseOperatorMigratedToOperatorSets(log types.Log) (*AVSDirectoryStorageOperatorMigratedToOperatorSets, error) {
	event := new(AVSDirectoryStorageOperatorMigratedToOperatorSets)
	if err := _AVSDirectoryStorage.contract.UnpackLog(event, "OperatorMigratedToOperatorSets", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSDirectoryStorageOperatorRemovedFromOperatorSetIterator is returned from FilterOperatorRemovedFromOperatorSet and is used to iterate over the raw logs and unpacked data for OperatorRemovedFromOperatorSet events raised by the AVSDirectoryStorage contract.
type AVSDirectoryStorageOperatorRemovedFromOperatorSetIterator struct {
	Event *AVSDirectoryStorageOperatorRemovedFromOperatorSet // Event containing the contract specifics and raw log

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
func (it *AVSDirectoryStorageOperatorRemovedFromOperatorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSDirectoryStorageOperatorRemovedFromOperatorSet)
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
		it.Event = new(AVSDirectoryStorageOperatorRemovedFromOperatorSet)
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
func (it *AVSDirectoryStorageOperatorRemovedFromOperatorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSDirectoryStorageOperatorRemovedFromOperatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSDirectoryStorageOperatorRemovedFromOperatorSet represents a OperatorRemovedFromOperatorSet event raised by the AVSDirectoryStorage contract.
type AVSDirectoryStorageOperatorRemovedFromOperatorSet struct {
	Operator    common.Address
	OperatorSet IAVSDirectoryOperatorSet
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorRemovedFromOperatorSet is a free log retrieval operation binding the contract event 0xad34c3070be1dffbcaa499d000ba2b8d9848aefcac3059df245dd95c4ece14fe.
//
// Solidity: event OperatorRemovedFromOperatorSet(address indexed operator, (address,uint32) operatorSet)
func (_AVSDirectoryStorage *AVSDirectoryStorageFilterer) FilterOperatorRemovedFromOperatorSet(opts *bind.FilterOpts, operator []common.Address) (*AVSDirectoryStorageOperatorRemovedFromOperatorSetIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _AVSDirectoryStorage.contract.FilterLogs(opts, "OperatorRemovedFromOperatorSet", operatorRule)
	if err != nil {
		return nil, err
	}
	return &AVSDirectoryStorageOperatorRemovedFromOperatorSetIterator{contract: _AVSDirectoryStorage.contract, event: "OperatorRemovedFromOperatorSet", logs: logs, sub: sub}, nil
}

// WatchOperatorRemovedFromOperatorSet is a free log subscription operation binding the contract event 0xad34c3070be1dffbcaa499d000ba2b8d9848aefcac3059df245dd95c4ece14fe.
//
// Solidity: event OperatorRemovedFromOperatorSet(address indexed operator, (address,uint32) operatorSet)
func (_AVSDirectoryStorage *AVSDirectoryStorageFilterer) WatchOperatorRemovedFromOperatorSet(opts *bind.WatchOpts, sink chan<- *AVSDirectoryStorageOperatorRemovedFromOperatorSet, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _AVSDirectoryStorage.contract.WatchLogs(opts, "OperatorRemovedFromOperatorSet", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSDirectoryStorageOperatorRemovedFromOperatorSet)
				if err := _AVSDirectoryStorage.contract.UnpackLog(event, "OperatorRemovedFromOperatorSet", log); err != nil {
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

// ParseOperatorRemovedFromOperatorSet is a log parse operation binding the contract event 0xad34c3070be1dffbcaa499d000ba2b8d9848aefcac3059df245dd95c4ece14fe.
//
// Solidity: event OperatorRemovedFromOperatorSet(address indexed operator, (address,uint32) operatorSet)
func (_AVSDirectoryStorage *AVSDirectoryStorageFilterer) ParseOperatorRemovedFromOperatorSet(log types.Log) (*AVSDirectoryStorageOperatorRemovedFromOperatorSet, error) {
	event := new(AVSDirectoryStorageOperatorRemovedFromOperatorSet)
	if err := _AVSDirectoryStorage.contract.UnpackLog(event, "OperatorRemovedFromOperatorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSDirectoryStorageOperatorSetCreatedIterator is returned from FilterOperatorSetCreated and is used to iterate over the raw logs and unpacked data for OperatorSetCreated events raised by the AVSDirectoryStorage contract.
type AVSDirectoryStorageOperatorSetCreatedIterator struct {
	Event *AVSDirectoryStorageOperatorSetCreated // Event containing the contract specifics and raw log

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
func (it *AVSDirectoryStorageOperatorSetCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSDirectoryStorageOperatorSetCreated)
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
		it.Event = new(AVSDirectoryStorageOperatorSetCreated)
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
func (it *AVSDirectoryStorageOperatorSetCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSDirectoryStorageOperatorSetCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSDirectoryStorageOperatorSetCreated represents a OperatorSetCreated event raised by the AVSDirectoryStorage contract.
type AVSDirectoryStorageOperatorSetCreated struct {
	OperatorSet IAVSDirectoryOperatorSet
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorSetCreated is a free log retrieval operation binding the contract event 0x31629285ead2335ae0933f86ed2ae63321f7af77b4e6eaabc42c057880977e6c.
//
// Solidity: event OperatorSetCreated((address,uint32) operatorSet)
func (_AVSDirectoryStorage *AVSDirectoryStorageFilterer) FilterOperatorSetCreated(opts *bind.FilterOpts) (*AVSDirectoryStorageOperatorSetCreatedIterator, error) {

	logs, sub, err := _AVSDirectoryStorage.contract.FilterLogs(opts, "OperatorSetCreated")
	if err != nil {
		return nil, err
	}
	return &AVSDirectoryStorageOperatorSetCreatedIterator{contract: _AVSDirectoryStorage.contract, event: "OperatorSetCreated", logs: logs, sub: sub}, nil
}

// WatchOperatorSetCreated is a free log subscription operation binding the contract event 0x31629285ead2335ae0933f86ed2ae63321f7af77b4e6eaabc42c057880977e6c.
//
// Solidity: event OperatorSetCreated((address,uint32) operatorSet)
func (_AVSDirectoryStorage *AVSDirectoryStorageFilterer) WatchOperatorSetCreated(opts *bind.WatchOpts, sink chan<- *AVSDirectoryStorageOperatorSetCreated) (event.Subscription, error) {

	logs, sub, err := _AVSDirectoryStorage.contract.WatchLogs(opts, "OperatorSetCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSDirectoryStorageOperatorSetCreated)
				if err := _AVSDirectoryStorage.contract.UnpackLog(event, "OperatorSetCreated", log); err != nil {
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
func (_AVSDirectoryStorage *AVSDirectoryStorageFilterer) ParseOperatorSetCreated(log types.Log) (*AVSDirectoryStorageOperatorSetCreated, error) {
	event := new(AVSDirectoryStorageOperatorSetCreated)
	if err := _AVSDirectoryStorage.contract.UnpackLog(event, "OperatorSetCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSDirectoryStorageOperatorSlashedIterator is returned from FilterOperatorSlashed and is used to iterate over the raw logs and unpacked data for OperatorSlashed events raised by the AVSDirectoryStorage contract.
type AVSDirectoryStorageOperatorSlashedIterator struct {
	Event *AVSDirectoryStorageOperatorSlashed // Event containing the contract specifics and raw log

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
func (it *AVSDirectoryStorageOperatorSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSDirectoryStorageOperatorSlashed)
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
		it.Event = new(AVSDirectoryStorageOperatorSlashed)
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
func (it *AVSDirectoryStorageOperatorSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSDirectoryStorageOperatorSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSDirectoryStorageOperatorSlashed represents a OperatorSlashed event raised by the AVSDirectoryStorage contract.
type AVSDirectoryStorageOperatorSlashed struct {
	Operator      common.Address
	OperatorSetId uint32
	Strategy      common.Address
	BipsToSlash   uint16
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOperatorSlashed is a free log retrieval operation binding the contract event 0xe672839d3c371691acdb52de9fefc94b3dbf407dc0920ef566c7c059ad575b1c.
//
// Solidity: event OperatorSlashed(address operator, uint32 operatorSetId, address strategy, uint16 bipsToSlash)
func (_AVSDirectoryStorage *AVSDirectoryStorageFilterer) FilterOperatorSlashed(opts *bind.FilterOpts) (*AVSDirectoryStorageOperatorSlashedIterator, error) {

	logs, sub, err := _AVSDirectoryStorage.contract.FilterLogs(opts, "OperatorSlashed")
	if err != nil {
		return nil, err
	}
	return &AVSDirectoryStorageOperatorSlashedIterator{contract: _AVSDirectoryStorage.contract, event: "OperatorSlashed", logs: logs, sub: sub}, nil
}

// WatchOperatorSlashed is a free log subscription operation binding the contract event 0xe672839d3c371691acdb52de9fefc94b3dbf407dc0920ef566c7c059ad575b1c.
//
// Solidity: event OperatorSlashed(address operator, uint32 operatorSetId, address strategy, uint16 bipsToSlash)
func (_AVSDirectoryStorage *AVSDirectoryStorageFilterer) WatchOperatorSlashed(opts *bind.WatchOpts, sink chan<- *AVSDirectoryStorageOperatorSlashed) (event.Subscription, error) {

	logs, sub, err := _AVSDirectoryStorage.contract.WatchLogs(opts, "OperatorSlashed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSDirectoryStorageOperatorSlashed)
				if err := _AVSDirectoryStorage.contract.UnpackLog(event, "OperatorSlashed", log); err != nil {
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
func (_AVSDirectoryStorage *AVSDirectoryStorageFilterer) ParseOperatorSlashed(log types.Log) (*AVSDirectoryStorageOperatorSlashed, error) {
	event := new(AVSDirectoryStorageOperatorSlashed)
	if err := _AVSDirectoryStorage.contract.UnpackLog(event, "OperatorSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
