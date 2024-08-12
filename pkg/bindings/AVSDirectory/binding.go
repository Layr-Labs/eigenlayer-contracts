// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package AVSDirectory

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

// IAVSDirectoryMagnitudeAdjustment is an auto generated low-level Go binding around an user-defined struct.
type IAVSDirectoryMagnitudeAdjustment struct {
	Strategy       common.Address
	OperatorSets   []IAVSDirectoryOperatorSet
	MagnitudeDiffs []uint64
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

// AVSDirectoryMetaData contains all meta data concerning the AVSDirectory contract.
var AVSDirectoryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_delegation\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ALLOCATION_DELAY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DOMAIN_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"OPERATOR_AVS_REGISTRATION_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"OPERATOR_SET_FORCE_DEREGISTRATION_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"OPERATOR_SET_REGISTRATION_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allocate\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allocations\",\"type\":\"tuple[]\",\"internalType\":\"structIAVSDirectory.MagnitudeAdjustment[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"operatorSets\",\"type\":\"tuple[]\",\"internalType\":\"structIAVSDirectory.OperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"magnitudeDiffs\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}]},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"avsOperatorStatus\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIAVSDirectory.OperatorAVSRegistrationStatus\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"becomeOperatorSetAVS\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"calculateOperatorAVSRegistrationDigestHash\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateOperatorSetForceDeregistrationTypehash\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateOperatorSetRegistrationDigestHash\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cancelSalt\",\"inputs\":[{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeDeallocations\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"operatorSets\",\"type\":\"tuple[][]\",\"internalType\":\"structIAVSDirectory.OperatorSet[][]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createOperatorSets\",\"inputs\":[{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregisterOperatorFromOperatorSets\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"domainSeparator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"forceDeregisterFromOperatorSets\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"freeMagnitude\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSlashableBips\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structIAVSDirectory.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"timestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"inTotalOperatorSets\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"initialPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isMember\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structIAVSDirectory.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperatorSet\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperatorSetAVS\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperatorSlashable\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structIAVSDirectory.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"migrateOperatorsToOperatorSets\",\"inputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"operatorSaltIsSpent\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorSetMemberCount\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorSetStatus\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"registered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"lastDeregisteredTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorSetsMemberOf\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"operatorSets\",\"type\":\"tuple[]\",\"internalType\":\"structIAVSDirectory.OperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorSetsMemberOf\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIAVSDirectory.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"queueDeallocate\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"deallocations\",\"type\":\"tuple[]\",\"internalType\":\"structIAVSDirectory.MagnitudeAdjustment[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"operatorSets\",\"type\":\"tuple[]\",\"internalType\":\"structIAVSDirectory.OperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"magnitudeDiffs\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}]},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerOperatorToOperatorSets\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slashOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"bipsToSlash\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAVSMetadataURI\",\"inputs\":[{\"name\":\"metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AVSMetadataURIUpdated\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"metadataURI\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AVSMigratedToOperatorSets\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MagnitudeAllocated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIAVSDirectory.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"magnitudeToAllocate\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"effectTimestamp\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MagnitudeDeallocationCompleted\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIAVSDirectory.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"freeMagnitudeAdded\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MagnitudeQueueDeallocated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIAVSDirectory.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"magnitudeToDeallocate\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"completableTimestamp\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorAVSRegistrationStatusUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"status\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumIAVSDirectory.OperatorAVSRegistrationStatus\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorAddedToOperatorSet\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIAVSDirectory.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorMigratedToOperatorSets\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"indexed\":false,\"internalType\":\"uint32[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRemovedFromOperatorSet\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIAVSDirectory.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSetCreated\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIAVSDirectory.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSlashed\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"bipsToSlash\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x60c06040523480156200001157600080fd5b506040516200597b3803806200597b833981016040819052620000349162000118565b6001600160a01b0381166080526200004b62000056565b504660a0526200014a565b600054610100900460ff1615620000c35760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff908116101562000116576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6000602082840312156200012b57600080fd5b81516001600160a01b03811681146200014357600080fd5b9392505050565b60805160a05161580462000177600039600061275a0152600081816107410152610d9101526158046000f3fe608060405234801561001057600080fd5b506004361061028a5760003560e01c8063886f11951161015c578063c825fe68116100ce578063df5cf72311610087578063df5cf7231461073c578063ec76f44214610763578063ef2dfa8d14610797578063f2fde38b146107aa578063f698da25146107bd578063fabc1cbc146107c557600080fd5b8063c825fe6814610676578063cbdf0e421461069d578063d79aceab146106b0578063da2ff05d146106d7578063dae226b6146106ea578063dce974b91461071557600080fd5b8063a98fb35511610120578063a98fb3551461060f578063aec205c514610622578063afe02ed51461062a578063b2841d481461063d578063bd74a06c14610650578063c1a8e2c51461066357600080fd5b8063886f11951461058d5780638da5cb5b146105b85780638de54944146105c9578063955e6696146105e9578063a1060c88146105fc57600080fd5b806349075da3116102005780635fd6abfd116101b95780635fd6abfd146104b65780636fbb031c146105025780637019670814610521578063715018a6146105345780637673e93a1461053c57806384d76f7b1461055f57600080fd5b806349075da3146104225780634ef8c5d41461045d578063595c6a67146104705780635a922613146104785780635ac86ab71461048b5780635c975abb146104ae57600080fd5b80631e2199e2116102525780631e2199e2146103165780631e68134e1461032957806320606b701461038657806333429a6a146103bb578063374823b5146103e15780633fee332d1461040f57600080fd5b80630d5387c51461028f57806310d67a2f146102b85780631352c3e6146102cd578063136439dd146102f05780631794bb3c14610303575b600080fd5b6102a261029d3660046148fc565b6107d8565b6040516102af9190614931565b60405180910390f35b6102cb6102c6366004614997565b61091d565b005b6102e06102db366004614a3b565b6109d9565b60405190151581526020016102af565b6102cb6102fe366004614ac0565b610a68565b6102cb610311366004614ad9565b610ba7565b6102cb610324366004614c10565b610cd1565b61036a610337366004614c8b565b609e60209081526000938452604080852082529284528284209052825290205460ff811690610100900463ffffffff1682565b60408051921515835263ffffffff9091166020830152016102af565b6103ad7f8cad95687ba82c2ce50e74f7b754645e5117c3a5bec8151c0726d5857980a86681565b6040519081526020016102af565b6103ce6103c9366004614cd2565b610ff1565b60405161ffff90911681526020016102af565b6102e06103ef366004614d36565b609960209081526000928352604080842090915290825290205460ff1681565b6102cb61041d366004614d62565b6110f9565b610450610430366004614df0565b609860209081526000928352604080842090915290825290205460ff1681565b6040516102af9190614e3f565b6102cb61046b366004614e67565b611343565b6102cb61139d565b6102cb610486366004614ee7565b611464565b6102e0610499366004614f69565b606654600160ff9092169190911b9081161490565b6066546103ad565b6104ea6104c4366004614df0565b60a06020908152600092835260408084209091529082529020546001600160401b031681565b6040516001600160401b0390911681526020016102af565b61050c621baf8081565b60405163ffffffff90911681526020016102af565b6102cb61052f366004614e67565b6115d2565b6102cb61162c565b6102e061054a366004614997565b609a6020526000908152604090205460ff1681565b6102e061056d366004614f8c565b609b60209081526000928352604080842090915290825290205460ff1681565b6065546105a0906001600160a01b031681565b6040516001600160a01b0390911681526020016102af565b6033546001600160a01b03166105a0565b6105dc6105d7366004614d36565b611640565b6040516102af9190614fc1565b6103ad6105f7366004614fe7565b611681565b6103ad61060a36600461504d565b6116e6565b6102cb61061d366004615093565b611750565b6102cb611797565b6102cb610638366004615104565b61185f565b6103ad61064b366004614fe7565b611a2c565b6102cb61065e366004615145565b611a6c565b6102cb6106713660046151c6565b611f53565b6103ad7f809c5ac049c45b7a7f050a20f00c16cf63797efbf8b1eb8d749fdfa39ff8f92981565b6103ad6106ab366004614997565b611f88565b6103ad7fda2c89bafdd34776a2b8bb9c83c82f419e20cc8c67207f70edd58249b92661bd81565b6102e06106e5366004614a3b565b611fa9565b6103ad6106f8366004614f8c565b609c60209081526000928352604080842090915290825290205481565b6103ad7f4ee65f64218c67b68da66fd0db16560040a6b973290b9e71912d661ee53fe49581565b6105a07f000000000000000000000000000000000000000000000000000000000000000081565b6102cb610771366004614ac0565b33600090815260996020908152604080832093835292905220805460ff19166001179055565b6102cb6107a536600461521a565b611fd5565b6102cb6107b8366004614997565b61237e565b6103ad6123f4565b6102cb6107d3366004614ac0565b612403565b6001600160a01b0383166000908152609d602052604081206060919084906107ff9061255f565b610809919061529b565b905080831115610817578092505b826001600160401b0381111561082f5761082f6149b4565b60405190808252806020026020018201604052801561087457816020015b604080518082019091526000808252602082015281526020019060019003908161084d5790505b50915060005b83811015610914576108e66108b061089283886152b2565b6001600160a01b0389166000908152609d6020526040902090612569565b60408051808201909152600080825260208201525060408051808201909152606082901c815263ffffffff909116602082015290565b8382815181106108f8576108f86152ca565b60200260200101819052508061090d906152e0565b905061087a565b50509392505050565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610970573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061099491906152fb565b6001600160a01b0316336001600160a01b0316146109cd5760405162461bcd60e51b81526004016109c490615318565b60405180910390fd5b6109d681612575565b50565b80516001600160a01b039081166000908152609e6020908152604080832093861683529281528282208185015163ffffffff908116845290825283832084518086019095525460ff811615158086526101009091049091169184019190915290919080610a5e575042621baf808260200151610a559190615362565b63ffffffff1610155b9150505b92915050565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015610ab0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ad4919061538a565b610af05760405162461bcd60e51b81526004016109c4906153ac565b60665481811614610b695760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c697479000000000000000060648201526084016109c4565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b600054610100900460ff1615808015610bc75750600054600160ff909116105b80610be15750303b158015610be1575060005460ff166001145b610c445760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016109c4565b6000805460ff191660011790558015610c67576000805461ff0019166101001790555b610c71838361266c565b610c79612756565b609755610c858461281f565b8015610ccb576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050565b60665460019060029081161415610cfa5760405162461bcd60e51b81526004016109c4906153f4565b4282604001511015610d725760405162461bcd60e51b815260206004820152604760248201526000805160206157af83398151915260448201527f70657261746f72536574733a206f70657261746f72207369676e617475726520606482015266195e1c1a5c995960ca1b608482015260a4016109c4565b6040516336b87bd760e11b81526001600160a01b0386811660048301527f00000000000000000000000000000000000000000000000000000000000000001690636d70f7ae90602401602060405180830381865afa158015610dd8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dfc919061538a565b610e7b5760405162461bcd60e51b815260206004820152605660248201526000805160206157af83398151915260448201527f70657261746f72536574733a206f70657261746f72206e6f7420726567697374606482015275195c9959081d1bc8115a59d95b93185e595c881e595d60521b608482015260a4016109c4565b336000908152609a602052604090205460ff16610f025760405162461bcd60e51b815260206004820152604b60248201526000805160206157af83398151915260448201527f70657261746f72536574733a20415653206973206e6f7420616e206f7065726160648201526a746f72207365742041565360a81b608482015260a4016109c4565b6001600160a01b038516600090815260996020908152604080832085830151845290915290205460ff1615610f8d5760405162461bcd60e51b815260206004820152603f60248201526000805160206157af83398151915260448201527f70657261746f72536574733a2073616c7420616c7265616479207370656e740060648201526084016109c4565b610fad85610fa633878787602001518860400151611681565b8451612871565b6001600160a01b03851660009081526099602090815260408083208583015184529091529020805460ff19166001179055610fea85338686612a2b565b5050505050565b6001600160a01b038085166000908152609f60209081526040808320938616835292905290812081906110249084612dfa565b6001600160a01b03808816600090815260a1602090815260408083209389168352928152918120929350916110c1918691908490611064908b018b614997565b6001600160a01b03166001600160a01b031681526020019081526020016000206000896020016020810190611099919061542b565b63ffffffff1663ffffffff168152602001908152602001600020612dfa90919063ffffffff16565b9050816001600160401b0316612710826001600160401b03166110e49190615446565b6110ee919061547b565b979650505050505050565b606654600190600290811614156111225760405162461bcd60e51b81526004016109c4906153f4565b8151516111ba57336001600160a01b038716146111b55760405162461bcd60e51b815260206004820152604560248201527f4156534469726563746f72792e666f7263654465726567697374657246726f6d60448201527f4f70657261746f72536574733a2063616c6c6572206d757374206265206f70656064820152643930ba37b960d91b608482015260a4016109c4565b61132f565b42826040015110156112455760405162461bcd60e51b815260206004820152604860248201527f4156534469726563746f72792e666f7263654465726567697374657246726f6d60448201527f4f70657261746f72536574733a206f70657261746f72207369676e617475726560648201526708195e1c1a5c995960c21b608482015260a4016109c4565b6001600160a01b038616600090815260996020908152604080832085830151845290915290205460ff16156112e4576040805162461bcd60e51b81526020600482015260248101919091527f4156534469726563746f72792e666f7263654465726567697374657246726f6d60448201527f4f70657261746f72536574733a2073616c7420616c7265616479207370656e7460648201526084016109c4565b6112fd86610fa687878787602001518860400151611a2c565b6001600160a01b03861660009081526099602090815260408083208583015184529091529020805460ff191660011790555b61133b85878686612e52565b505050505050565b6000611352621baf8042615362565b905060005b8381101561133b5761138d86868684818110611375576113756152ca565b9050602002810190611387919061549d565b84613034565b611396816152e0565b9050611357565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa1580156113e5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611409919061538a565b6114255760405162461bcd60e51b81526004016109c4906153ac565b600019606681905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b60005b8381101561133b576000805b848484818110611485576114856152ca565b905060200281019061149791906154bd565b90508110156115265761150a888888868181106114b6576114b66152ca565b90506020020160208101906114cb9190614997565b8787878181106114dd576114dd6152ca565b90506020028101906114ef91906154bd565b858181106114ff576114ff6152ca565b905060400201613518565b6115149083615506565b915061151f816152e0565b9050611473565b506001600160a01b038716600090815260a0602052604081208291888886818110611553576115536152ca565b90506020020160208101906115689190614997565b6001600160a01b0316815260208101919091526040016000908120805490919061159c9084906001600160401b0316615506565b92506101000a8154816001600160401b0302191690836001600160401b0316021790555050806115cb906152e0565b9050611467565b60006115e1621baf8042615362565b905060005b8381101561133b5761161c86868684818110611604576116046152ca565b9050602002810190611616919061549d565b846137a5565b611625816152e0565b90506115e6565b611634613cdf565b61163e600061281f565b565b60408051808201909152600080825260208201526001600160a01b0383166000908152609d6020526040902061167a906108b09084612569565b9392505050565b60006116dc7f809c5ac049c45b7a7f050a20f00c16cf63797efbf8b1eb8d749fdfa39ff8f92987878787876040516020016116c19695949392919061556c565b60405160208183030381529060405280519060200120613d39565b9695505050505050565b604080517fda2c89bafdd34776a2b8bb9c83c82f419e20cc8c67207f70edd58249b92661bd60208201526001600160a01b038087169282019290925290841660608201526080810183905260a081018290526000906117479060c0016116c1565b95945050505050565b336001600160a01b03167fa89c1dc243d8908a96dd84944bcc97d6bc6ac00dd78e20621576be6a3c943713838360405161178b9291906155ac565b60405180910390a25050565b336000908152609a602052604090205460ff161561181d5760405162461bcd60e51b815260206004820152603e60248201527f4156534469726563746f72792e6265636f6d654f70657261746f72536574415660448201527f533a20616c726561647920616e206f70657261746f722073657420415653000060648201526084016109c4565b336000818152609a6020526040808220805460ff19166001179055517f702b0c1f6cb1cf511aaa81f72bc05a215bb3497632d72c690c822b044ab494bf9190a2565b60005b81811015611a2757336000908152609b602052604081209084848481811061188c5761188c6152ca565b90506020020160208101906118a1919061542b565b63ffffffff16815260208101919091526040016000205460ff161561192e5760405162461bcd60e51b815260206004820152603b60248201527f4156534469726563746f72792e6372656174654f70657261746f725365743a2060448201527f6f70657261746f722073657420616c726561647920657869737473000000000060648201526084016109c4565b336000908152609b60205260408120600191858585818110611952576119526152ca565b9050602002016020810190611967919061542b565b63ffffffff1663ffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055507f31629285ead2335ae0933f86ed2ae63321f7af77b4e6eaabc42c057880977e6c6040518060400160405280336001600160a01b031681526020018585858181106119e5576119e56152ca565b90506020020160208101906119fa919061542b565b63ffffffff169052604051611a0f9190614fc1565b60405180910390a1611a20816152e0565b9050611862565b505050565b60006116dc7f4ee65f64218c67b68da66fd0db16560040a6b973290b9e71912d661ee53fe49587878787876040516020016116c19695949392919061556c565b611a99856040518060400160405280336001600160a01b031681526020018763ffffffff168152506109d9565b611b165760405162461bcd60e51b815260206004820152604260248201527f4156534469726563746f72792e736c6173684f70657261746f723a206f70657260448201527f61746f72206e6f7420736c61736861626c6520666f72206f70657261746f7253606482015261195d60f21b608482015260a4016109c4565b60005b8281101561133b576001600160a01b038616600090815260a1602052604081208190611ba490429083898988818110611b5457611b546152ca565b9050602002016020810190611b699190614997565b6001600160a01b0316815260208082019290925260409081016000908120338252835281812063ffffffff808e1683529352209190613d8016565b9050612710611bc06001600160401b03831661ffff8716615446565b611bca919061547b565b6001600160a01b038916600090815260a160205260408120919350611c5a9142916001600160401b03861691908a8a89818110611c0957611c096152ca565b9050602002016020810190611c1e9190614997565b6001600160a01b0316815260208082019290925260409081016000908120338252835281812063ffffffff808f168352935220929190613e2516565b506001600160a01b038716600090815260a260205260408120819081888887818110611c8857611c886152ca565b9050602002016020810190611c9d9190614997565b6001600160a01b0316815260208082019290925260409081016000908120338252835281812063ffffffff8c1682529092529020549050805b8015611e3a576001600160a01b038a16600090815260a260205260408120818a8a89818110611d0757611d076152ca565b9050602002016020810190611d1c9190614997565b6001600160a01b0316815260208082019290925260409081016000908120338252835281812063ffffffff8e1682529092529020611d5b60018461529b565b81548110611d6b57611d6b6152ca565b60009182526020909120018054909150600160401b900463ffffffff16611d95621baf8042615362565b63ffffffff161115611e2357805460009061271090611dc1906001600160401b031661ffff8b16615446565b611dcb919061547b565b825490915081908390600090611deb9084906001600160401b03166155db565b92506101000a8154816001600160401b0302191690836001600160401b031602179055508085611e1b9190615506565b945050611e29565b50611e3a565b50611e3381615603565b9050611cd6565b5050611f3e42826001600160401b0316846001600160401b0316611eca609f60008e6001600160a01b03166001600160a01b0316815260200190815260200160002060008c8c8b818110611e9057611e906152ca565b9050602002016020810190611ea59190614997565b6001600160a01b03166001600160a01b03168152602001908152602001600020613ec1565b611ed4919061561a565b611ede919061561a565b6001600160a01b038b166000908152609f60205260408120908a8a89818110611f0957611f096152ca565b9050602002016020810190611f1e9190614997565b6001600160a01b0316815260208101919091526040016000209190613efa565b5050505080611f4c906152e0565b9050611b19565b60665460019060029081161415611f7c5760405162461bcd60e51b81526004016109c4906153f4565b610ccb33858585612e52565b6001600160a01b0381166000908152609d60205260408120610a629061255f565b600061167a611fb783613f15565b6001600160a01b0385166000908152609d6020526040902090613f7a565b60665460019060029081161415611ffe5760405162461bcd60e51b81526004016109c4906153f4565b336000908152609a602052604090205460ff166120975760405162461bcd60e51b815260206004820152604b60248201527f4156534469726563746f72792e6d6967726174654f70657261746f7273546f4f60448201527f70657261746f72536574733a20415653206973206e6f7420616e206f7065726160648201526a746f72207365742041565360a81b608482015260a4016109c4565b60005b8481101561133b576001336000908152609860205260408120908888858181106120c6576120c66152ca565b90506020020160208101906120db9190614997565b6001600160a01b0316815260208101919091526040016000205460ff16600181111561210957612109614e29565b146121b55760405162461bcd60e51b815260206004820152606a60248201527f4156534469726563746f72792e6d6967726174654f70657261746f7273546f4f60448201527f70657261746f72536574733a206f70657261746f7220616c7265616479206d6960648201527f677261746564206f72206e6f742061206c656761637920726567697374657265608482015269321037b832b930ba37b960b11b60a482015260c4016109c4565b6122098686838181106121ca576121ca6152ca565b90506020020160208101906121df9190614997565b338686858181106121f2576121f26152ca565b9050602002810190612204919061563a565b612a2b565b3360009081526098602052604081208188888581811061222b5761222b6152ca565b90506020020160208101906122409190614997565b6001600160a01b031681526020810191909152604001600020805460ff19166001838181111561227257612272614e29565b02179055503386868381811061228a5761228a6152ca565b905060200201602081019061229f9190614997565b6001600160a01b03167ff0952b1c65271d819d39983d2abb044b9cace59bcc4d4dd389f586ebdcb15b4160006040516122d89190614e3f565b60405180910390a3338686838181106122f3576122f36152ca565b90506020020160208101906123089190614997565b6001600160a01b03167f54f33cfdd1ca703d795986b986fd47d742eab1904ecd2a5fdb8d6595e5904a01868685818110612344576123446152ca565b9050602002810190612356919061563a565b604051612364929190615683565b60405180910390a380612376816152e0565b91505061209a565b612386613cdf565b6001600160a01b0381166123eb5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016109c4565b6109d68161281f565b60006123fe612756565b905090565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612456573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061247a91906152fb565b6001600160a01b0316336001600160a01b0316146124aa5760405162461bcd60e51b81526004016109c490615318565b6066541981196066541916146125285760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c697479000000000000000060648201526084016109c4565b606681905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c90602001610b9c565b6000610a62825490565b600061167a8383613f92565b6001600160a01b0381166126035760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a4016109c4565b606554604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1606580546001600160a01b0319166001600160a01b0392909216919091179055565b6065546001600160a01b031615801561268d57506001600160a01b03821615155b61270f5760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a4016109c4565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a261275282612575565b5050565b60007f0000000000000000000000000000000000000000000000000000000000000000461415612787575060975490565b50604080518082018252600a81526922b4b3b2b72630bcb2b960b11b60209182015281517f8cad95687ba82c2ce50e74f7b754645e5117c3a5bec8151c0726d5857980a866818301527f71b625cfad44bac63b13dba07f2e1d6084ee04b6f8752101ece6126d584ee6ea81840152466060820152306080808301919091528351808303909101815260a0909101909252815191012090565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6001600160a01b0383163b1561298b57604051630b135d3f60e11b808252906001600160a01b03851690631626ba7e906128b1908690869060040161569f565b602060405180830381865afa1580156128ce573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906128f291906156fc565b6001600160e01b03191614611a275760405162461bcd60e51b815260206004820152605360248201527f454950313237315369676e61747572655574696c732e636865636b5369676e6160448201527f747572655f454950313237313a2045524331323731207369676e6174757265206064820152721d995c9a599a58d85d1a5bdb8819985a5b1959606a1b608482015260a4016109c4565b826001600160a01b031661299f8383613fbc565b6001600160a01b031614611a275760405162461bcd60e51b815260206004820152604760248201527f454950313237315369676e61747572655574696c732e636865636b5369676e6160448201527f747572655f454950313237313a207369676e6174757265206e6f742066726f6d6064820152661039b4b3b732b960c91b608482015260a4016109c4565b60005b81811015610fea5760006040518060400160405280866001600160a01b03168152602001858585818110612a6457612a646152ca565b9050602002016020810190612a79919061542b565b63ffffffff1690526001600160a01b0386166000908152609b60205260408120919250858585818110612aae57612aae6152ca565b9050602002016020810190612ac3919061542b565b63ffffffff16815260208101919091526040016000205460ff16612b5a5760405162461bcd60e51b815260206004820152604260248201527f4156534469726563746f72792e5f72656769737465724f70657261746f72546f60448201527f4f70657261746f72536574733a20696e76616c6964206f70657261746f722073606482015261195d60f21b608482015260a4016109c4565b612b648682611fa9565b15612bfd5760405162461bcd60e51b815260206004820152605960248201527f4156534469726563746f72792e5f72656769737465724f70657261746f72546f60448201527f4f70657261746f72536574733a206f70657261746f7220616c7265616479207260648201527f65676973746572656420746f206f70657261746f722073657400000000000000608482015260a4016109c4565b6001600160a01b0385166000908152609c6020526040812090858585818110612c2857612c286152ca565b9050602002016020810190612c3d919061542b565b63ffffffff1663ffffffff16815260200190815260200160002060008154612c64906152e0565b90915550612c92612c7482613f15565b6001600160a01b0388166000908152609d6020526040902090613fe0565b506001600160a01b038086166000908152609e60209081526040808320938a16835292905290812081868686818110612ccd57612ccd6152ca565b9050602002016020810190612ce2919061542b565b63ffffffff1681526020810191909152604001600020805490915060ff1615612d995760405162461bcd60e51b815260206004820152605a60248201527f4156534469726563746f72792e5f72656769737465724f70657261746f72546f60448201527f4f70657261746f72536574733a206f70657261746f7220616c7265616479207260648201527f65676973746572656420666f72206f70657261746f7220736574000000000000608482015260a4016109c4565b805460ff191660011781556040516001600160a01b038816907f43232edf9071753d2321e5fa7e018363ee248e5f2142e6c08edd3265bfb4895e90612ddf908590614fc1565b60405180910390a2505080612df3906152e0565b9050612a2e565b815460009081612e0c85858385613fec565b90508015612e4757612e3185612e2360018461529b565b600091825260209091200190565b54600160201b90046001600160e01b0316611747565b600095945050505050565b60005b81811015610fea5760006040518060400160405280876001600160a01b03168152602001858585818110612e8b57612e8b6152ca565b9050602002016020810190612ea0919061542b565b63ffffffff1690529050612eb48582611fa9565b612f4c5760405162461bcd60e51b815260206004820152605960248201527f4156534469726563746f72792e5f646572656769737465724f70657261746f7260448201527f46726f6d4f70657261746f725365743a206f70657261746f72206e6f7420726560648201527f676973746572656420666f72206f70657261746f722073657400000000000000608482015260a4016109c4565b6001600160a01b0386166000908152609c6020526040812090858585818110612f7757612f776152ca565b9050602002016020810190612f8c919061542b565b63ffffffff1663ffffffff16815260200190815260200160002060008154612fb390615603565b90915550612fe1612fc382613f15565b6001600160a01b0387166000908152609d6020526040902090614042565b50846001600160a01b03167fad34c3070be1dffbcaa499d000ba2b8d9848aefcac3059df245dd95c4ece14fe8260405161301b9190614fc1565b60405180910390a25061302d816152e0565b9050612e55565b60006130436020840184614997565b9050613052604084018461563a565b905061306160208501856154bd565b9050146130ee5760405162461bcd60e51b815260206004820152604f60248201527f4156534469726563746f72792e71756575654465616c6c6f636174696f6e3a2060448201527f6f70657261746f725365747320616e64206d61676e697475646544696666732060648201526e0d8cadccee8d040dad2e6dac2e8c6d608b1b608482015260a4016109c4565b3660006130fe60208601866154bd565b9150915060005b8181101561350f576001600160a01b03808816600090815260a16020908152604080832093881683529290529081206131d79042908387878781811061314d5761314d6152ca565b6131639260206040909202019081019150614997565b6001600160a01b03166001600160a01b031681526020019081526020016000206000878787818110613197576131976152ca565b90506040020160200160208101906131af919061542b565b63ffffffff1663ffffffff168152602001908152602001600020613d8090919063ffffffff16565b90506001600160401b0381166131f0604089018961563a565b84818110613200576132006152ca565b90506020020160208101906132159190615726565b6001600160401b031611156132a85760405162461bcd60e51b815260206004820152604d60248201527f4156534469726563746f72792e71756575654465616c6c6f636174696f6e3a2060448201527f63616e6e6f74206465616c6c6f63617465206d6f7265207468616e207768617460648201526c081a5cc8185b1b1bd8d85d1959609a1b608482015260a4016109c4565b6133b0426132b960408a018a61563a565b858181106132c9576132c96152ca565b90506020020160208101906132de9190615726565b6001600160a01b03808c16600090815260a160209081526040808320938c1683529290529081206001600160401b03929092169190888888818110613325576133256152ca565b61333b9260206040909202019081019150614997565b6001600160a01b03166001600160a01b03168152602001908152602001600020600088888881811061336f5761336f6152ca565b9050604002016020016020810190613387919061542b565b63ffffffff1663ffffffff168152602001908152602001600020613e259092919063ffffffff16565b6001600160a01b03808916600090815260a2602090815260408083209389168352929052908120908585858181106133ea576133ea6152ca565b6134009260206040909202019081019150614997565b6001600160a01b03166001600160a01b031681526020019081526020016000206000858585818110613434576134346152ca565b905060400201602001602081019061344c919061542b565b63ffffffff1663ffffffff1681526020019081526020016000206040518060400160405280898060400190613481919061563a565b86818110613491576134916152ca565b90506020020160208101906134a69190615726565b6001600160401b03908116825263ffffffff808b166020938401528454600181018655600095865294839020845195018054949093015116600160401b026bffffffffffffffffffffffff1990931693169290921717905550613508816152e0565b9050613105565b50505050505050565b6001600160a01b03808416600090815260a26020908152604080832093861683529281529181209091829190829061355290860186614997565b6001600160a01b03166001600160a01b031681526020019081526020016000206000846020016020810190613587919061542b565b63ffffffff1663ffffffff168152602001908152602001600020805480602002602001604051908101604052809291908181526020016000905b8282101561360d57600084815260209081902060408051808201909152908401546001600160401b0381168252600160401b900463ffffffff16818301528252600190920191016135c1565b5050506001600160a01b03808816600090815260a360209081526040808320938a168352928152918120939450929150829061364b90870187614997565b6001600160a01b03166001600160a01b031681526020019081526020016000206000856020016020810190613680919061542b565b63ffffffff1663ffffffff1681526020019081526020016000205490505b8151811015613718574263ffffffff168282815181106136c0576136c06152ca565b60200260200101516020015163ffffffff1611156136dd57613718565b8181815181106136ef576136ef6152ca565b602002602001015160000151836137069190615506565b9250613711816152e0565b905061369e565b6001600160a01b03808716600090815260a3602090815260408083209389168352928152918120839290919061375090880188614997565b6001600160a01b03166001600160a01b031681526020019081526020016000206000866020016020810190613785919061542b565b63ffffffff16815260208101919091526040016000205550509392505050565b60006137b46020840184614997565b90503660006137c660208601866154bd565b6001600160a01b03808916600090815260a0602090815260408083209389168352929052908120549294509092506001600160401b03909116905b82811015613c8f57613816604088018861563a565b82818110613826576138266152ca565b905060200201602081019061383b9190615726565b6001600160401b0316826001600160401b031610156138db5760405162461bcd60e51b815260206004820152605060248201527f4156534469726563746f72792e7175657565416c6c6f636174696f6e733a206960448201527f6e73756666696369656e7420617661696c61626c652066726565206d61676e6960648201526f7475646520746f20616c6c6f6361746560801b608482015260a4016109c4565b6001600160a01b03808916600090815260a160209081526040808320938916835292905290812081906139a79042908389898881811061391d5761391d6152ca565b6139339260206040909202019081019150614997565b6001600160a01b03166001600160a01b031681526020019081526020016000206000898988818110613967576139676152ca565b905060400201602001602081019061397f919061542b565b63ffffffff1663ffffffff16815260200190815260200160002061404e90919063ffffffff16565b91509150816001600160e01b031660001415806139c357508015155b15613b2d576001600160a01b03808b16600090815260a160209081526040808320938b168352929052908120600191613a899190898988818110613a0957613a096152ca565b613a1f9260206040909202019081019150614997565b6001600160a01b03166001600160a01b031681526020019081526020016000206000898988818110613a5357613a536152ca565b9050604002016020016020810190613a6b919061542b565b63ffffffff1663ffffffff1681526020019081526020016000205490565b613a93919061529b565b8114613b2d5760405162461bcd60e51b815260206004820152605a60248201527f4156534469726563746f72792e7175657565416c6c6f636174696f6e733a206f60448201527f6e6c79206f6e652070656e64696e6720616c6c6f636174696f6e20616c6c6f7760648201527f656420666f72206f702c206f705365742c207374726174656779000000000000608482015260a4016109c4565b613c3b88613b3e60408c018c61563a565b86818110613b4e57613b4e6152ca565b9050602002016020810190613b639190615726565b613b76906001600160401b03168561574f565b6001600160a01b03808e16600090815260a160209081526040808320938e168352929052908120908a8a89818110613bb057613bb06152ca565b613bc69260206040909202019081019150614997565b6001600160a01b03166001600160a01b0316815260200190815260200160002060008a8a89818110613bfa57613bfa6152ca565b9050604002016020016020810190613c12919061542b565b63ffffffff1663ffffffff168152602001908152602001600020613efa9092919063ffffffff16565b50613c4b905060408a018a61563a565b84818110613c5b57613c5b6152ca565b9050602002016020810190613c709190615726565b613c7a90856155db565b9350505080613c88906152e0565b9050613801565b506001600160a01b03968716600090815260a0602090815260408083209690991682529490945295909220805467ffffffffffffffff19166001600160401b039096169590951790945550505050565b6033546001600160a01b0316331461163e5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016109c4565b6000613d43612756565b60405161190160f01b6020820152602281019190915260428101839052606201604051602081830303815290604052805190602001209050919050565b815460009081816005811115613ddd576000613d9b84614107565b613da5908561529b565b60008881526020902090915081015463ffffffff9081169087161015613dcd57809150613ddb565b613dd88160016152b2565b92505b505b6000613deb87878585613fec565b90508015613e1857613e0287612e2360018461529b565b54600160201b90046001600160e01b03166110ee565b6000979650505050505050565b600080613e32858561404e565b91509150816001600160e01b03166000148015613e4d575080155b15613e5757506000195b84545b8082101561133b576000868152602090208201805485908290600490613e91908490600160201b90046001600160e01b031661561a565b92506101000a8154816001600160e01b0302191690836001600160e01b0316021790555082600101925050613e5a565b80546000908015613ef157613edb83612e2360018461529b565b54600160201b90046001600160e01b031661167a565b60009392505050565b600080613f088585856141ec565b915091505b935093915050565b60008160000151826020015163ffffffff16604051602001613f6292919060609290921b6bffffffffffffffffffffffff1916825260a01b6001600160a01b031916601482015260200190565b604051602081830303815290604052610a6290615771565b6000818152600183016020526040812054151561167a565b6000826000018281548110613fa957613fa96152ca565b9060005260206000200154905092915050565b6000806000613fcb8585614392565b91509150613fd8816143ff565b509392505050565b600061167a83836145ba565b60005b81831015613fd85760006140038484614609565b60008781526020902090915063ffffffff86169082015463ffffffff16111561402e5780925061403c565b6140398160016152b2565b93505b50613fef565b600061167a8383614624565b81546000908190818160058111156140ad57600061406b84614107565b614075908561529b565b60008981526020902090915081015463ffffffff908116908816101561409d578091506140ab565b6140a88160016152b2565b92505b505b60006140bb88888585613fec565b905080156140f3576140d288612e2360018461529b565b54600160201b90046001600160e01b03166140ee60018361529b565b6140f7565b6000805b95509550505050505b9250929050565b60008161411657506000919050565b6000600161412384614717565b901c6001901b9050600181848161413c5761413c615465565b048201901c9050600181848161415457614154615465565b048201901c9050600181848161416c5761416c615465565b048201901c9050600181848161418457614184615465565b048201901c9050600181848161419c5761419c615465565b048201901c905060018184816141b4576141b4615465565b048201901c905060018184816141cc576141cc615465565b048201901c905061167a818285816141e6576141e6615465565b046147ab565b82546000908190801561433957600061420a87612e2360018561529b565b60408051808201909152905463ffffffff808216808452600160201b9092046001600160e01b03166020840152919250908716101561428b5760405162461bcd60e51b815260206004820152601b60248201527f436865636b706f696e743a2064656372656173696e67206b657973000000000060448201526064016109c4565b8563ffffffff16816000015163ffffffff1614156142da57846142b388612e2360018661529b565b80546001600160e01b0392909216600160201b0263ffffffff909216919091179055614329565b6040805180820190915263ffffffff80881682526001600160e01b0380881660208085019182528b54600181018d5560008d81529190912094519151909216600160201b029216919091179101555b602001519250839150613f0d9050565b50506040805180820190915263ffffffff80851682526001600160e01b0380851660208085019182528854600181018a5560008a815291822095519251909316600160201b029190931617920191909155905081613f0d565b6000808251604114156143c95760208301516040840151606085015160001a6143bd878285856147c1565b94509450505050614100565b8251604014156143f357602083015160408401516143e88683836148ae565b935093505050614100565b50600090506002614100565b600081600481111561441357614413614e29565b141561441c5750565b600181600481111561443057614430614e29565b141561447e5760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e6174757265000000000000000060448201526064016109c4565b600281600481111561449257614492614e29565b14156144e05760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e6774680060448201526064016109c4565b60038160048111156144f4576144f4614e29565b141561454d5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b60648201526084016109c4565b600481600481111561456157614561614e29565b14156109d65760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b60648201526084016109c4565b600081815260018301602052604081205461460157508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610a62565b506000610a62565b6000614618600284841861547b565b61167a908484166152b2565b6000818152600183016020526040812054801561470d57600061464860018361529b565b855490915060009061465c9060019061529b565b90508181146146c157600086600001828154811061467c5761467c6152ca565b906000526020600020015490508087600001848154811061469f5761469f6152ca565b6000918252602080832090910192909255918252600188019052604090208390555b85548690806146d2576146d2615798565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050610a62565b6000915050610a62565b600080608083901c1561472c57608092831c92015b604083901c1561473e57604092831c92015b602083901c1561475057602092831c92015b601083901c1561476257601092831c92015b600883901c1561477457600892831c92015b600483901c1561478657600492831c92015b600283901c1561479857600292831c92015b600183901c15610a625760010192915050565b60008183106147ba578161167a565b5090919050565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08311156147f857506000905060036148a5565b8460ff16601b1415801561481057508460ff16601c14155b1561482157506000905060046148a5565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015614875573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b03811661489e576000600192509250506148a5565b9150600090505b94509492505050565b6000806001600160ff1b038316816148cb60ff86901c601b6152b2565b90506148d9878288856147c1565b935093505050935093915050565b6001600160a01b03811681146109d657600080fd5b60008060006060848603121561491157600080fd5b833561491c816148e7565b95602085013595506040909401359392505050565b602080825282518282018190526000919060409081850190868401855b8281101561498a5761497a84835180516001600160a01b0316825260209081015163ffffffff16910152565b928401929085019060010161494e565b5091979650505050505050565b6000602082840312156149a957600080fd5b813561167a816148e7565b634e487b7160e01b600052604160045260246000fd5b604051606081016001600160401b03811182821017156149ec576149ec6149b4565b60405290565b604051601f8201601f191681016001600160401b0381118282101715614a1a57614a1a6149b4565b604052919050565b803563ffffffff81168114614a3657600080fd5b919050565b6000808284036060811215614a4f57600080fd5b8335614a5a816148e7565b92506040601f1982011215614a6e57600080fd5b50604051604081018181106001600160401b0382111715614a9157614a916149b4565b6040526020840135614aa2816148e7565b8152614ab060408501614a22565b6020820152809150509250929050565b600060208284031215614ad257600080fd5b5035919050565b600080600060608486031215614aee57600080fd5b8335614af9816148e7565b92506020840135614b09816148e7565b929592945050506040919091013590565b60008083601f840112614b2c57600080fd5b5081356001600160401b03811115614b4357600080fd5b6020830191508360208260051b850101111561410057600080fd5b600060608284031215614b7057600080fd5b614b786149ca565b905081356001600160401b0380821115614b9157600080fd5b818401915084601f830112614ba557600080fd5b8135602082821115614bb957614bb96149b4565b614bcb601f8301601f191682016149f2565b92508183528681838601011115614be157600080fd5b818185018285013760008183850101528285528086013581860152505050506040820135604082015292915050565b60008060008060608587031215614c2657600080fd5b8435614c31816148e7565b935060208501356001600160401b0380821115614c4d57600080fd5b614c5988838901614b1a565b90955093506040870135915080821115614c7257600080fd5b50614c7f87828801614b5e565b91505092959194509250565b600080600060608486031215614ca057600080fd5b8335614cab816148e7565b92506020840135614cbb816148e7565b9150614cc960408501614a22565b90509250925092565b60008060008084860360a0811215614ce957600080fd5b8535614cf4816148e7565b94506040601f1982011215614d0857600080fd5b506020850192506060850135614d1d816148e7565b9150614d2b60808601614a22565b905092959194509250565b60008060408385031215614d4957600080fd5b8235614d54816148e7565b946020939093013593505050565b600080600080600060808688031215614d7a57600080fd5b8535614d85816148e7565b94506020860135614d95816148e7565b935060408601356001600160401b0380821115614db157600080fd5b614dbd89838a01614b1a565b90955093506060880135915080821115614dd657600080fd5b50614de388828901614b5e565b9150509295509295909350565b60008060408385031215614e0357600080fd5b8235614e0e816148e7565b91506020830135614e1e816148e7565b809150509250929050565b634e487b7160e01b600052602160045260246000fd5b6020810160028310614e6157634e487b7160e01b600052602160045260246000fd5b91905290565b60008060008060608587031215614e7d57600080fd5b8435614e88816148e7565b935060208501356001600160401b0380821115614ea457600080fd5b614eb088838901614b1a565b90955093506040870135915080821115614ec957600080fd5b50850160608188031215614edc57600080fd5b939692955090935050565b600080600080600060608688031215614eff57600080fd5b8535614f0a816148e7565b945060208601356001600160401b0380821115614f2657600080fd5b614f3289838a01614b1a565b90965094506040880135915080821115614f4b57600080fd5b50614f5888828901614b1a565b969995985093965092949392505050565b600060208284031215614f7b57600080fd5b813560ff8116811461167a57600080fd5b60008060408385031215614f9f57600080fd5b8235614faa816148e7565b9150614fb860208401614a22565b90509250929050565b81516001600160a01b0316815260208083015163ffffffff169082015260408101610a62565b600080600080600060808688031215614fff57600080fd5b853561500a816148e7565b945060208601356001600160401b0381111561502557600080fd5b61503188828901614b1a565b9699909850959660408101359660609091013595509350505050565b6000806000806080858703121561506357600080fd5b843561506e816148e7565b9350602085013561507e816148e7565b93969395505050506040820135916060013590565b600080602083850312156150a657600080fd5b82356001600160401b03808211156150bd57600080fd5b818501915085601f8301126150d157600080fd5b8135818111156150e057600080fd5b8660208285010111156150f257600080fd5b60209290920196919550909350505050565b6000806020838503121561511757600080fd5b82356001600160401b0381111561512d57600080fd5b61513985828601614b1a565b90969095509350505050565b60008060008060006080868803121561515d57600080fd5b8535615168816148e7565b945061517660208701614a22565b935060408601356001600160401b0381111561519157600080fd5b61519d88828901614b1a565b909450925050606086013561ffff811681146151b857600080fd5b809150509295509295909350565b6000806000604084860312156151db57600080fd5b83356151e6816148e7565b925060208401356001600160401b0381111561520157600080fd5b61520d86828701614b1a565b9497909650939450505050565b6000806000806040858703121561523057600080fd5b84356001600160401b038082111561524757600080fd5b61525388838901614b1a565b9096509450602087013591508082111561526c57600080fd5b5061527987828801614b1a565b95989497509550505050565b634e487b7160e01b600052601160045260246000fd5b6000828210156152ad576152ad615285565b500390565b600082198211156152c5576152c5615285565b500190565b634e487b7160e01b600052603260045260246000fd5b60006000198214156152f4576152f4615285565b5060010190565b60006020828403121561530d57600080fd5b815161167a816148e7565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b600063ffffffff80831681851680830382111561538157615381615285565b01949350505050565b60006020828403121561539c57600080fd5b8151801515811461167a57600080fd5b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b60208082526019908201527f5061757361626c653a20696e6465782069732070617573656400000000000000604082015260600190565b60006020828403121561543d57600080fd5b61167a82614a22565b600081600019048311821515161561546057615460615285565b500290565b634e487b7160e01b600052601260045260246000fd5b60008261549857634e487b7160e01b600052601260045260246000fd5b500490565b60008235605e198336030181126154b357600080fd5b9190910192915050565b6000808335601e198436030181126154d457600080fd5b8301803591506001600160401b038211156154ee57600080fd5b6020019150600681901b360382131561410057600080fd5b60006001600160401b0380831681851680830382111561538157615381615285565b8183526000602080850194508260005b858110156155615763ffffffff61554e83614a22565b1687529582019590820190600101615538565b509495945050505050565b8681526001600160a01b038616602082015260a0604082018190526000906155979083018688615528565b60608301949094525060800152949350505050565b60208152816020820152818360408301376000818301604090810191909152601f909201601f19160101919050565b60006001600160401b03838116908316818110156155fb576155fb615285565b039392505050565b60008161561257615612615285565b506000190190565b60006001600160e01b03838116908316818110156155fb576155fb615285565b6000808335601e1984360301811261565157600080fd5b8301803591506001600160401b0382111561566b57600080fd5b6020019150600581901b360382131561410057600080fd5b602081526000615697602083018486615528565b949350505050565b82815260006020604081840152835180604085015260005b818110156156d3578581018301518582016060015282016156b7565b818111156156e5576000606083870101525b50601f01601f191692909201606001949350505050565b60006020828403121561570e57600080fd5b81516001600160e01b03198116811461167a57600080fd5b60006020828403121561573857600080fd5b81356001600160401b038116811461167a57600080fd5b60006001600160e01b0382811684821680830382111561538157615381615285565b80516020808301519190811015615792576000198160200360031b1b821691505b50919050565b634e487b7160e01b600052603160045260246000fdfe4156534469726563746f72792e72656769737465724f70657261746f72546f4fa2646970667358221220513df02979800c403aae5cc96dc9610ab9fc2f9f49a7cd59e5a5b557da30581f64736f6c634300080c0033",
}

// AVSDirectoryABI is the input ABI used to generate the binding from.
// Deprecated: Use AVSDirectoryMetaData.ABI instead.
var AVSDirectoryABI = AVSDirectoryMetaData.ABI

// AVSDirectoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AVSDirectoryMetaData.Bin instead.
var AVSDirectoryBin = AVSDirectoryMetaData.Bin

// DeployAVSDirectory deploys a new Ethereum contract, binding an instance of AVSDirectory to it.
func DeployAVSDirectory(auth *bind.TransactOpts, backend bind.ContractBackend, _delegation common.Address) (common.Address, *types.Transaction, *AVSDirectory, error) {
	parsed, err := AVSDirectoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AVSDirectoryBin), backend, _delegation)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AVSDirectory{AVSDirectoryCaller: AVSDirectoryCaller{contract: contract}, AVSDirectoryTransactor: AVSDirectoryTransactor{contract: contract}, AVSDirectoryFilterer: AVSDirectoryFilterer{contract: contract}}, nil
}

// AVSDirectory is an auto generated Go binding around an Ethereum contract.
type AVSDirectory struct {
	AVSDirectoryCaller     // Read-only binding to the contract
	AVSDirectoryTransactor // Write-only binding to the contract
	AVSDirectoryFilterer   // Log filterer for contract events
}

// AVSDirectoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type AVSDirectoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AVSDirectoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AVSDirectoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AVSDirectoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AVSDirectoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AVSDirectorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AVSDirectorySession struct {
	Contract     *AVSDirectory     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AVSDirectoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AVSDirectoryCallerSession struct {
	Contract *AVSDirectoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// AVSDirectoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AVSDirectoryTransactorSession struct {
	Contract     *AVSDirectoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// AVSDirectoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type AVSDirectoryRaw struct {
	Contract *AVSDirectory // Generic contract binding to access the raw methods on
}

// AVSDirectoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AVSDirectoryCallerRaw struct {
	Contract *AVSDirectoryCaller // Generic read-only contract binding to access the raw methods on
}

// AVSDirectoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AVSDirectoryTransactorRaw struct {
	Contract *AVSDirectoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAVSDirectory creates a new instance of AVSDirectory, bound to a specific deployed contract.
func NewAVSDirectory(address common.Address, backend bind.ContractBackend) (*AVSDirectory, error) {
	contract, err := bindAVSDirectory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AVSDirectory{AVSDirectoryCaller: AVSDirectoryCaller{contract: contract}, AVSDirectoryTransactor: AVSDirectoryTransactor{contract: contract}, AVSDirectoryFilterer: AVSDirectoryFilterer{contract: contract}}, nil
}

// NewAVSDirectoryCaller creates a new read-only instance of AVSDirectory, bound to a specific deployed contract.
func NewAVSDirectoryCaller(address common.Address, caller bind.ContractCaller) (*AVSDirectoryCaller, error) {
	contract, err := bindAVSDirectory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AVSDirectoryCaller{contract: contract}, nil
}

// NewAVSDirectoryTransactor creates a new write-only instance of AVSDirectory, bound to a specific deployed contract.
func NewAVSDirectoryTransactor(address common.Address, transactor bind.ContractTransactor) (*AVSDirectoryTransactor, error) {
	contract, err := bindAVSDirectory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AVSDirectoryTransactor{contract: contract}, nil
}

// NewAVSDirectoryFilterer creates a new log filterer instance of AVSDirectory, bound to a specific deployed contract.
func NewAVSDirectoryFilterer(address common.Address, filterer bind.ContractFilterer) (*AVSDirectoryFilterer, error) {
	contract, err := bindAVSDirectory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AVSDirectoryFilterer{contract: contract}, nil
}

// bindAVSDirectory binds a generic wrapper to an already deployed contract.
func bindAVSDirectory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AVSDirectoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AVSDirectory *AVSDirectoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AVSDirectory.Contract.AVSDirectoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AVSDirectory *AVSDirectoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AVSDirectory.Contract.AVSDirectoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AVSDirectory *AVSDirectoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AVSDirectory.Contract.AVSDirectoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AVSDirectory *AVSDirectoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AVSDirectory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AVSDirectory *AVSDirectoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AVSDirectory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AVSDirectory *AVSDirectoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AVSDirectory.Contract.contract.Transact(opts, method, params...)
}

// ALLOCATIONDELAY is a free data retrieval call binding the contract method 0x6fbb031c.
//
// Solidity: function ALLOCATION_DELAY() view returns(uint32)
func (_AVSDirectory *AVSDirectoryCaller) ALLOCATIONDELAY(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _AVSDirectory.contract.Call(opts, &out, "ALLOCATION_DELAY")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ALLOCATIONDELAY is a free data retrieval call binding the contract method 0x6fbb031c.
//
// Solidity: function ALLOCATION_DELAY() view returns(uint32)
func (_AVSDirectory *AVSDirectorySession) ALLOCATIONDELAY() (uint32, error) {
	return _AVSDirectory.Contract.ALLOCATIONDELAY(&_AVSDirectory.CallOpts)
}

// ALLOCATIONDELAY is a free data retrieval call binding the contract method 0x6fbb031c.
//
// Solidity: function ALLOCATION_DELAY() view returns(uint32)
func (_AVSDirectory *AVSDirectoryCallerSession) ALLOCATIONDELAY() (uint32, error) {
	return _AVSDirectory.Contract.ALLOCATIONDELAY(&_AVSDirectory.CallOpts)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_AVSDirectory *AVSDirectoryCaller) DOMAINTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AVSDirectory.contract.Call(opts, &out, "DOMAIN_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_AVSDirectory *AVSDirectorySession) DOMAINTYPEHASH() ([32]byte, error) {
	return _AVSDirectory.Contract.DOMAINTYPEHASH(&_AVSDirectory.CallOpts)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_AVSDirectory *AVSDirectoryCallerSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _AVSDirectory.Contract.DOMAINTYPEHASH(&_AVSDirectory.CallOpts)
}

// OPERATORAVSREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0xd79aceab.
//
// Solidity: function OPERATOR_AVS_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_AVSDirectory *AVSDirectoryCaller) OPERATORAVSREGISTRATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AVSDirectory.contract.Call(opts, &out, "OPERATOR_AVS_REGISTRATION_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORAVSREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0xd79aceab.
//
// Solidity: function OPERATOR_AVS_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_AVSDirectory *AVSDirectorySession) OPERATORAVSREGISTRATIONTYPEHASH() ([32]byte, error) {
	return _AVSDirectory.Contract.OPERATORAVSREGISTRATIONTYPEHASH(&_AVSDirectory.CallOpts)
}

// OPERATORAVSREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0xd79aceab.
//
// Solidity: function OPERATOR_AVS_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_AVSDirectory *AVSDirectoryCallerSession) OPERATORAVSREGISTRATIONTYPEHASH() ([32]byte, error) {
	return _AVSDirectory.Contract.OPERATORAVSREGISTRATIONTYPEHASH(&_AVSDirectory.CallOpts)
}

// OPERATORSETFORCEDEREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0xdce974b9.
//
// Solidity: function OPERATOR_SET_FORCE_DEREGISTRATION_TYPEHASH() view returns(bytes32)
func (_AVSDirectory *AVSDirectoryCaller) OPERATORSETFORCEDEREGISTRATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AVSDirectory.contract.Call(opts, &out, "OPERATOR_SET_FORCE_DEREGISTRATION_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORSETFORCEDEREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0xdce974b9.
//
// Solidity: function OPERATOR_SET_FORCE_DEREGISTRATION_TYPEHASH() view returns(bytes32)
func (_AVSDirectory *AVSDirectorySession) OPERATORSETFORCEDEREGISTRATIONTYPEHASH() ([32]byte, error) {
	return _AVSDirectory.Contract.OPERATORSETFORCEDEREGISTRATIONTYPEHASH(&_AVSDirectory.CallOpts)
}

// OPERATORSETFORCEDEREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0xdce974b9.
//
// Solidity: function OPERATOR_SET_FORCE_DEREGISTRATION_TYPEHASH() view returns(bytes32)
func (_AVSDirectory *AVSDirectoryCallerSession) OPERATORSETFORCEDEREGISTRATIONTYPEHASH() ([32]byte, error) {
	return _AVSDirectory.Contract.OPERATORSETFORCEDEREGISTRATIONTYPEHASH(&_AVSDirectory.CallOpts)
}

// OPERATORSETREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0xc825fe68.
//
// Solidity: function OPERATOR_SET_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_AVSDirectory *AVSDirectoryCaller) OPERATORSETREGISTRATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AVSDirectory.contract.Call(opts, &out, "OPERATOR_SET_REGISTRATION_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORSETREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0xc825fe68.
//
// Solidity: function OPERATOR_SET_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_AVSDirectory *AVSDirectorySession) OPERATORSETREGISTRATIONTYPEHASH() ([32]byte, error) {
	return _AVSDirectory.Contract.OPERATORSETREGISTRATIONTYPEHASH(&_AVSDirectory.CallOpts)
}

// OPERATORSETREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0xc825fe68.
//
// Solidity: function OPERATOR_SET_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_AVSDirectory *AVSDirectoryCallerSession) OPERATORSETREGISTRATIONTYPEHASH() ([32]byte, error) {
	return _AVSDirectory.Contract.OPERATORSETREGISTRATIONTYPEHASH(&_AVSDirectory.CallOpts)
}

// AvsOperatorStatus is a free data retrieval call binding the contract method 0x49075da3.
//
// Solidity: function avsOperatorStatus(address , address ) view returns(uint8)
func (_AVSDirectory *AVSDirectoryCaller) AvsOperatorStatus(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (uint8, error) {
	var out []interface{}
	err := _AVSDirectory.contract.Call(opts, &out, "avsOperatorStatus", arg0, arg1)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// AvsOperatorStatus is a free data retrieval call binding the contract method 0x49075da3.
//
// Solidity: function avsOperatorStatus(address , address ) view returns(uint8)
func (_AVSDirectory *AVSDirectorySession) AvsOperatorStatus(arg0 common.Address, arg1 common.Address) (uint8, error) {
	return _AVSDirectory.Contract.AvsOperatorStatus(&_AVSDirectory.CallOpts, arg0, arg1)
}

// AvsOperatorStatus is a free data retrieval call binding the contract method 0x49075da3.
//
// Solidity: function avsOperatorStatus(address , address ) view returns(uint8)
func (_AVSDirectory *AVSDirectoryCallerSession) AvsOperatorStatus(arg0 common.Address, arg1 common.Address) (uint8, error) {
	return _AVSDirectory.Contract.AvsOperatorStatus(&_AVSDirectory.CallOpts, arg0, arg1)
}

// CalculateOperatorAVSRegistrationDigestHash is a free data retrieval call binding the contract method 0xa1060c88.
//
// Solidity: function calculateOperatorAVSRegistrationDigestHash(address operator, address avs, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_AVSDirectory *AVSDirectoryCaller) CalculateOperatorAVSRegistrationDigestHash(opts *bind.CallOpts, operator common.Address, avs common.Address, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _AVSDirectory.contract.Call(opts, &out, "calculateOperatorAVSRegistrationDigestHash", operator, avs, salt, expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateOperatorAVSRegistrationDigestHash is a free data retrieval call binding the contract method 0xa1060c88.
//
// Solidity: function calculateOperatorAVSRegistrationDigestHash(address operator, address avs, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_AVSDirectory *AVSDirectorySession) CalculateOperatorAVSRegistrationDigestHash(operator common.Address, avs common.Address, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _AVSDirectory.Contract.CalculateOperatorAVSRegistrationDigestHash(&_AVSDirectory.CallOpts, operator, avs, salt, expiry)
}

// CalculateOperatorAVSRegistrationDigestHash is a free data retrieval call binding the contract method 0xa1060c88.
//
// Solidity: function calculateOperatorAVSRegistrationDigestHash(address operator, address avs, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_AVSDirectory *AVSDirectoryCallerSession) CalculateOperatorAVSRegistrationDigestHash(operator common.Address, avs common.Address, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _AVSDirectory.Contract.CalculateOperatorAVSRegistrationDigestHash(&_AVSDirectory.CallOpts, operator, avs, salt, expiry)
}

// CalculateOperatorSetForceDeregistrationTypehash is a free data retrieval call binding the contract method 0xb2841d48.
//
// Solidity: function calculateOperatorSetForceDeregistrationTypehash(address avs, uint32[] operatorSetIds, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_AVSDirectory *AVSDirectoryCaller) CalculateOperatorSetForceDeregistrationTypehash(opts *bind.CallOpts, avs common.Address, operatorSetIds []uint32, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _AVSDirectory.contract.Call(opts, &out, "calculateOperatorSetForceDeregistrationTypehash", avs, operatorSetIds, salt, expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateOperatorSetForceDeregistrationTypehash is a free data retrieval call binding the contract method 0xb2841d48.
//
// Solidity: function calculateOperatorSetForceDeregistrationTypehash(address avs, uint32[] operatorSetIds, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_AVSDirectory *AVSDirectorySession) CalculateOperatorSetForceDeregistrationTypehash(avs common.Address, operatorSetIds []uint32, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _AVSDirectory.Contract.CalculateOperatorSetForceDeregistrationTypehash(&_AVSDirectory.CallOpts, avs, operatorSetIds, salt, expiry)
}

// CalculateOperatorSetForceDeregistrationTypehash is a free data retrieval call binding the contract method 0xb2841d48.
//
// Solidity: function calculateOperatorSetForceDeregistrationTypehash(address avs, uint32[] operatorSetIds, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_AVSDirectory *AVSDirectoryCallerSession) CalculateOperatorSetForceDeregistrationTypehash(avs common.Address, operatorSetIds []uint32, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _AVSDirectory.Contract.CalculateOperatorSetForceDeregistrationTypehash(&_AVSDirectory.CallOpts, avs, operatorSetIds, salt, expiry)
}

// CalculateOperatorSetRegistrationDigestHash is a free data retrieval call binding the contract method 0x955e6696.
//
// Solidity: function calculateOperatorSetRegistrationDigestHash(address avs, uint32[] operatorSetIds, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_AVSDirectory *AVSDirectoryCaller) CalculateOperatorSetRegistrationDigestHash(opts *bind.CallOpts, avs common.Address, operatorSetIds []uint32, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _AVSDirectory.contract.Call(opts, &out, "calculateOperatorSetRegistrationDigestHash", avs, operatorSetIds, salt, expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateOperatorSetRegistrationDigestHash is a free data retrieval call binding the contract method 0x955e6696.
//
// Solidity: function calculateOperatorSetRegistrationDigestHash(address avs, uint32[] operatorSetIds, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_AVSDirectory *AVSDirectorySession) CalculateOperatorSetRegistrationDigestHash(avs common.Address, operatorSetIds []uint32, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _AVSDirectory.Contract.CalculateOperatorSetRegistrationDigestHash(&_AVSDirectory.CallOpts, avs, operatorSetIds, salt, expiry)
}

// CalculateOperatorSetRegistrationDigestHash is a free data retrieval call binding the contract method 0x955e6696.
//
// Solidity: function calculateOperatorSetRegistrationDigestHash(address avs, uint32[] operatorSetIds, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_AVSDirectory *AVSDirectoryCallerSession) CalculateOperatorSetRegistrationDigestHash(avs common.Address, operatorSetIds []uint32, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _AVSDirectory.Contract.CalculateOperatorSetRegistrationDigestHash(&_AVSDirectory.CallOpts, avs, operatorSetIds, salt, expiry)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_AVSDirectory *AVSDirectoryCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AVSDirectory.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_AVSDirectory *AVSDirectorySession) Delegation() (common.Address, error) {
	return _AVSDirectory.Contract.Delegation(&_AVSDirectory.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_AVSDirectory *AVSDirectoryCallerSession) Delegation() (common.Address, error) {
	return _AVSDirectory.Contract.Delegation(&_AVSDirectory.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_AVSDirectory *AVSDirectoryCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AVSDirectory.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_AVSDirectory *AVSDirectorySession) DomainSeparator() ([32]byte, error) {
	return _AVSDirectory.Contract.DomainSeparator(&_AVSDirectory.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_AVSDirectory *AVSDirectoryCallerSession) DomainSeparator() ([32]byte, error) {
	return _AVSDirectory.Contract.DomainSeparator(&_AVSDirectory.CallOpts)
}

// FreeMagnitude is a free data retrieval call binding the contract method 0x5fd6abfd.
//
// Solidity: function freeMagnitude(address , address ) view returns(uint64)
func (_AVSDirectory *AVSDirectoryCaller) FreeMagnitude(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (uint64, error) {
	var out []interface{}
	err := _AVSDirectory.contract.Call(opts, &out, "freeMagnitude", arg0, arg1)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// FreeMagnitude is a free data retrieval call binding the contract method 0x5fd6abfd.
//
// Solidity: function freeMagnitude(address , address ) view returns(uint64)
func (_AVSDirectory *AVSDirectorySession) FreeMagnitude(arg0 common.Address, arg1 common.Address) (uint64, error) {
	return _AVSDirectory.Contract.FreeMagnitude(&_AVSDirectory.CallOpts, arg0, arg1)
}

// FreeMagnitude is a free data retrieval call binding the contract method 0x5fd6abfd.
//
// Solidity: function freeMagnitude(address , address ) view returns(uint64)
func (_AVSDirectory *AVSDirectoryCallerSession) FreeMagnitude(arg0 common.Address, arg1 common.Address) (uint64, error) {
	return _AVSDirectory.Contract.FreeMagnitude(&_AVSDirectory.CallOpts, arg0, arg1)
}

// GetSlashableBips is a free data retrieval call binding the contract method 0x33429a6a.
//
// Solidity: function getSlashableBips(address operator, (address,uint32) operatorSet, address strategy, uint32 timestamp) view returns(uint16)
func (_AVSDirectory *AVSDirectoryCaller) GetSlashableBips(opts *bind.CallOpts, operator common.Address, operatorSet IAVSDirectoryOperatorSet, strategy common.Address, timestamp uint32) (uint16, error) {
	var out []interface{}
	err := _AVSDirectory.contract.Call(opts, &out, "getSlashableBips", operator, operatorSet, strategy, timestamp)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetSlashableBips is a free data retrieval call binding the contract method 0x33429a6a.
//
// Solidity: function getSlashableBips(address operator, (address,uint32) operatorSet, address strategy, uint32 timestamp) view returns(uint16)
func (_AVSDirectory *AVSDirectorySession) GetSlashableBips(operator common.Address, operatorSet IAVSDirectoryOperatorSet, strategy common.Address, timestamp uint32) (uint16, error) {
	return _AVSDirectory.Contract.GetSlashableBips(&_AVSDirectory.CallOpts, operator, operatorSet, strategy, timestamp)
}

// GetSlashableBips is a free data retrieval call binding the contract method 0x33429a6a.
//
// Solidity: function getSlashableBips(address operator, (address,uint32) operatorSet, address strategy, uint32 timestamp) view returns(uint16)
func (_AVSDirectory *AVSDirectoryCallerSession) GetSlashableBips(operator common.Address, operatorSet IAVSDirectoryOperatorSet, strategy common.Address, timestamp uint32) (uint16, error) {
	return _AVSDirectory.Contract.GetSlashableBips(&_AVSDirectory.CallOpts, operator, operatorSet, strategy, timestamp)
}

// InTotalOperatorSets is a free data retrieval call binding the contract method 0xcbdf0e42.
//
// Solidity: function inTotalOperatorSets(address operator) view returns(uint256)
func (_AVSDirectory *AVSDirectoryCaller) InTotalOperatorSets(opts *bind.CallOpts, operator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AVSDirectory.contract.Call(opts, &out, "inTotalOperatorSets", operator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InTotalOperatorSets is a free data retrieval call binding the contract method 0xcbdf0e42.
//
// Solidity: function inTotalOperatorSets(address operator) view returns(uint256)
func (_AVSDirectory *AVSDirectorySession) InTotalOperatorSets(operator common.Address) (*big.Int, error) {
	return _AVSDirectory.Contract.InTotalOperatorSets(&_AVSDirectory.CallOpts, operator)
}

// InTotalOperatorSets is a free data retrieval call binding the contract method 0xcbdf0e42.
//
// Solidity: function inTotalOperatorSets(address operator) view returns(uint256)
func (_AVSDirectory *AVSDirectoryCallerSession) InTotalOperatorSets(operator common.Address) (*big.Int, error) {
	return _AVSDirectory.Contract.InTotalOperatorSets(&_AVSDirectory.CallOpts, operator)
}

// IsMember is a free data retrieval call binding the contract method 0xda2ff05d.
//
// Solidity: function isMember(address operator, (address,uint32) operatorSet) view returns(bool)
func (_AVSDirectory *AVSDirectoryCaller) IsMember(opts *bind.CallOpts, operator common.Address, operatorSet IAVSDirectoryOperatorSet) (bool, error) {
	var out []interface{}
	err := _AVSDirectory.contract.Call(opts, &out, "isMember", operator, operatorSet)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMember is a free data retrieval call binding the contract method 0xda2ff05d.
//
// Solidity: function isMember(address operator, (address,uint32) operatorSet) view returns(bool)
func (_AVSDirectory *AVSDirectorySession) IsMember(operator common.Address, operatorSet IAVSDirectoryOperatorSet) (bool, error) {
	return _AVSDirectory.Contract.IsMember(&_AVSDirectory.CallOpts, operator, operatorSet)
}

// IsMember is a free data retrieval call binding the contract method 0xda2ff05d.
//
// Solidity: function isMember(address operator, (address,uint32) operatorSet) view returns(bool)
func (_AVSDirectory *AVSDirectoryCallerSession) IsMember(operator common.Address, operatorSet IAVSDirectoryOperatorSet) (bool, error) {
	return _AVSDirectory.Contract.IsMember(&_AVSDirectory.CallOpts, operator, operatorSet)
}

// IsOperatorSet is a free data retrieval call binding the contract method 0x84d76f7b.
//
// Solidity: function isOperatorSet(address , uint32 ) view returns(bool)
func (_AVSDirectory *AVSDirectoryCaller) IsOperatorSet(opts *bind.CallOpts, arg0 common.Address, arg1 uint32) (bool, error) {
	var out []interface{}
	err := _AVSDirectory.contract.Call(opts, &out, "isOperatorSet", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperatorSet is a free data retrieval call binding the contract method 0x84d76f7b.
//
// Solidity: function isOperatorSet(address , uint32 ) view returns(bool)
func (_AVSDirectory *AVSDirectorySession) IsOperatorSet(arg0 common.Address, arg1 uint32) (bool, error) {
	return _AVSDirectory.Contract.IsOperatorSet(&_AVSDirectory.CallOpts, arg0, arg1)
}

// IsOperatorSet is a free data retrieval call binding the contract method 0x84d76f7b.
//
// Solidity: function isOperatorSet(address , uint32 ) view returns(bool)
func (_AVSDirectory *AVSDirectoryCallerSession) IsOperatorSet(arg0 common.Address, arg1 uint32) (bool, error) {
	return _AVSDirectory.Contract.IsOperatorSet(&_AVSDirectory.CallOpts, arg0, arg1)
}

// IsOperatorSetAVS is a free data retrieval call binding the contract method 0x7673e93a.
//
// Solidity: function isOperatorSetAVS(address ) view returns(bool)
func (_AVSDirectory *AVSDirectoryCaller) IsOperatorSetAVS(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _AVSDirectory.contract.Call(opts, &out, "isOperatorSetAVS", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperatorSetAVS is a free data retrieval call binding the contract method 0x7673e93a.
//
// Solidity: function isOperatorSetAVS(address ) view returns(bool)
func (_AVSDirectory *AVSDirectorySession) IsOperatorSetAVS(arg0 common.Address) (bool, error) {
	return _AVSDirectory.Contract.IsOperatorSetAVS(&_AVSDirectory.CallOpts, arg0)
}

// IsOperatorSetAVS is a free data retrieval call binding the contract method 0x7673e93a.
//
// Solidity: function isOperatorSetAVS(address ) view returns(bool)
func (_AVSDirectory *AVSDirectoryCallerSession) IsOperatorSetAVS(arg0 common.Address) (bool, error) {
	return _AVSDirectory.Contract.IsOperatorSetAVS(&_AVSDirectory.CallOpts, arg0)
}

// IsOperatorSlashable is a free data retrieval call binding the contract method 0x1352c3e6.
//
// Solidity: function isOperatorSlashable(address operator, (address,uint32) operatorSet) view returns(bool)
func (_AVSDirectory *AVSDirectoryCaller) IsOperatorSlashable(opts *bind.CallOpts, operator common.Address, operatorSet IAVSDirectoryOperatorSet) (bool, error) {
	var out []interface{}
	err := _AVSDirectory.contract.Call(opts, &out, "isOperatorSlashable", operator, operatorSet)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperatorSlashable is a free data retrieval call binding the contract method 0x1352c3e6.
//
// Solidity: function isOperatorSlashable(address operator, (address,uint32) operatorSet) view returns(bool)
func (_AVSDirectory *AVSDirectorySession) IsOperatorSlashable(operator common.Address, operatorSet IAVSDirectoryOperatorSet) (bool, error) {
	return _AVSDirectory.Contract.IsOperatorSlashable(&_AVSDirectory.CallOpts, operator, operatorSet)
}

// IsOperatorSlashable is a free data retrieval call binding the contract method 0x1352c3e6.
//
// Solidity: function isOperatorSlashable(address operator, (address,uint32) operatorSet) view returns(bool)
func (_AVSDirectory *AVSDirectoryCallerSession) IsOperatorSlashable(operator common.Address, operatorSet IAVSDirectoryOperatorSet) (bool, error) {
	return _AVSDirectory.Contract.IsOperatorSlashable(&_AVSDirectory.CallOpts, operator, operatorSet)
}

// OperatorSaltIsSpent is a free data retrieval call binding the contract method 0x374823b5.
//
// Solidity: function operatorSaltIsSpent(address , bytes32 ) view returns(bool)
func (_AVSDirectory *AVSDirectoryCaller) OperatorSaltIsSpent(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (bool, error) {
	var out []interface{}
	err := _AVSDirectory.contract.Call(opts, &out, "operatorSaltIsSpent", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OperatorSaltIsSpent is a free data retrieval call binding the contract method 0x374823b5.
//
// Solidity: function operatorSaltIsSpent(address , bytes32 ) view returns(bool)
func (_AVSDirectory *AVSDirectorySession) OperatorSaltIsSpent(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _AVSDirectory.Contract.OperatorSaltIsSpent(&_AVSDirectory.CallOpts, arg0, arg1)
}

// OperatorSaltIsSpent is a free data retrieval call binding the contract method 0x374823b5.
//
// Solidity: function operatorSaltIsSpent(address , bytes32 ) view returns(bool)
func (_AVSDirectory *AVSDirectoryCallerSession) OperatorSaltIsSpent(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _AVSDirectory.Contract.OperatorSaltIsSpent(&_AVSDirectory.CallOpts, arg0, arg1)
}

// OperatorSetMemberCount is a free data retrieval call binding the contract method 0xdae226b6.
//
// Solidity: function operatorSetMemberCount(address , uint32 ) view returns(uint256)
func (_AVSDirectory *AVSDirectoryCaller) OperatorSetMemberCount(opts *bind.CallOpts, arg0 common.Address, arg1 uint32) (*big.Int, error) {
	var out []interface{}
	err := _AVSDirectory.contract.Call(opts, &out, "operatorSetMemberCount", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatorSetMemberCount is a free data retrieval call binding the contract method 0xdae226b6.
//
// Solidity: function operatorSetMemberCount(address , uint32 ) view returns(uint256)
func (_AVSDirectory *AVSDirectorySession) OperatorSetMemberCount(arg0 common.Address, arg1 uint32) (*big.Int, error) {
	return _AVSDirectory.Contract.OperatorSetMemberCount(&_AVSDirectory.CallOpts, arg0, arg1)
}

// OperatorSetMemberCount is a free data retrieval call binding the contract method 0xdae226b6.
//
// Solidity: function operatorSetMemberCount(address , uint32 ) view returns(uint256)
func (_AVSDirectory *AVSDirectoryCallerSession) OperatorSetMemberCount(arg0 common.Address, arg1 uint32) (*big.Int, error) {
	return _AVSDirectory.Contract.OperatorSetMemberCount(&_AVSDirectory.CallOpts, arg0, arg1)
}

// OperatorSetStatus is a free data retrieval call binding the contract method 0x1e68134e.
//
// Solidity: function operatorSetStatus(address , address , uint32 ) view returns(bool registered, uint32 lastDeregisteredTimestamp)
func (_AVSDirectory *AVSDirectoryCaller) OperatorSetStatus(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 uint32) (struct {
	Registered                bool
	LastDeregisteredTimestamp uint32
}, error) {
	var out []interface{}
	err := _AVSDirectory.contract.Call(opts, &out, "operatorSetStatus", arg0, arg1, arg2)

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
func (_AVSDirectory *AVSDirectorySession) OperatorSetStatus(arg0 common.Address, arg1 common.Address, arg2 uint32) (struct {
	Registered                bool
	LastDeregisteredTimestamp uint32
}, error) {
	return _AVSDirectory.Contract.OperatorSetStatus(&_AVSDirectory.CallOpts, arg0, arg1, arg2)
}

// OperatorSetStatus is a free data retrieval call binding the contract method 0x1e68134e.
//
// Solidity: function operatorSetStatus(address , address , uint32 ) view returns(bool registered, uint32 lastDeregisteredTimestamp)
func (_AVSDirectory *AVSDirectoryCallerSession) OperatorSetStatus(arg0 common.Address, arg1 common.Address, arg2 uint32) (struct {
	Registered                bool
	LastDeregisteredTimestamp uint32
}, error) {
	return _AVSDirectory.Contract.OperatorSetStatus(&_AVSDirectory.CallOpts, arg0, arg1, arg2)
}

// OperatorSetsMemberOf is a free data retrieval call binding the contract method 0x0d5387c5.
//
// Solidity: function operatorSetsMemberOf(address operator, uint256 start, uint256 length) view returns((address,uint32)[] operatorSets)
func (_AVSDirectory *AVSDirectoryCaller) OperatorSetsMemberOf(opts *bind.CallOpts, operator common.Address, start *big.Int, length *big.Int) ([]IAVSDirectoryOperatorSet, error) {
	var out []interface{}
	err := _AVSDirectory.contract.Call(opts, &out, "operatorSetsMemberOf", operator, start, length)

	if err != nil {
		return *new([]IAVSDirectoryOperatorSet), err
	}

	out0 := *abi.ConvertType(out[0], new([]IAVSDirectoryOperatorSet)).(*[]IAVSDirectoryOperatorSet)

	return out0, err

}

// OperatorSetsMemberOf is a free data retrieval call binding the contract method 0x0d5387c5.
//
// Solidity: function operatorSetsMemberOf(address operator, uint256 start, uint256 length) view returns((address,uint32)[] operatorSets)
func (_AVSDirectory *AVSDirectorySession) OperatorSetsMemberOf(operator common.Address, start *big.Int, length *big.Int) ([]IAVSDirectoryOperatorSet, error) {
	return _AVSDirectory.Contract.OperatorSetsMemberOf(&_AVSDirectory.CallOpts, operator, start, length)
}

// OperatorSetsMemberOf is a free data retrieval call binding the contract method 0x0d5387c5.
//
// Solidity: function operatorSetsMemberOf(address operator, uint256 start, uint256 length) view returns((address,uint32)[] operatorSets)
func (_AVSDirectory *AVSDirectoryCallerSession) OperatorSetsMemberOf(operator common.Address, start *big.Int, length *big.Int) ([]IAVSDirectoryOperatorSet, error) {
	return _AVSDirectory.Contract.OperatorSetsMemberOf(&_AVSDirectory.CallOpts, operator, start, length)
}

// OperatorSetsMemberOf0 is a free data retrieval call binding the contract method 0x8de54944.
//
// Solidity: function operatorSetsMemberOf(address operator, uint256 index) view returns((address,uint32))
func (_AVSDirectory *AVSDirectoryCaller) OperatorSetsMemberOf0(opts *bind.CallOpts, operator common.Address, index *big.Int) (IAVSDirectoryOperatorSet, error) {
	var out []interface{}
	err := _AVSDirectory.contract.Call(opts, &out, "operatorSetsMemberOf0", operator, index)

	if err != nil {
		return *new(IAVSDirectoryOperatorSet), err
	}

	out0 := *abi.ConvertType(out[0], new(IAVSDirectoryOperatorSet)).(*IAVSDirectoryOperatorSet)

	return out0, err

}

// OperatorSetsMemberOf0 is a free data retrieval call binding the contract method 0x8de54944.
//
// Solidity: function operatorSetsMemberOf(address operator, uint256 index) view returns((address,uint32))
func (_AVSDirectory *AVSDirectorySession) OperatorSetsMemberOf0(operator common.Address, index *big.Int) (IAVSDirectoryOperatorSet, error) {
	return _AVSDirectory.Contract.OperatorSetsMemberOf0(&_AVSDirectory.CallOpts, operator, index)
}

// OperatorSetsMemberOf0 is a free data retrieval call binding the contract method 0x8de54944.
//
// Solidity: function operatorSetsMemberOf(address operator, uint256 index) view returns((address,uint32))
func (_AVSDirectory *AVSDirectoryCallerSession) OperatorSetsMemberOf0(operator common.Address, index *big.Int) (IAVSDirectoryOperatorSet, error) {
	return _AVSDirectory.Contract.OperatorSetsMemberOf0(&_AVSDirectory.CallOpts, operator, index)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AVSDirectory *AVSDirectoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AVSDirectory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AVSDirectory *AVSDirectorySession) Owner() (common.Address, error) {
	return _AVSDirectory.Contract.Owner(&_AVSDirectory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AVSDirectory *AVSDirectoryCallerSession) Owner() (common.Address, error) {
	return _AVSDirectory.Contract.Owner(&_AVSDirectory.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_AVSDirectory *AVSDirectoryCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _AVSDirectory.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_AVSDirectory *AVSDirectorySession) Paused(index uint8) (bool, error) {
	return _AVSDirectory.Contract.Paused(&_AVSDirectory.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_AVSDirectory *AVSDirectoryCallerSession) Paused(index uint8) (bool, error) {
	return _AVSDirectory.Contract.Paused(&_AVSDirectory.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_AVSDirectory *AVSDirectoryCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AVSDirectory.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_AVSDirectory *AVSDirectorySession) Paused0() (*big.Int, error) {
	return _AVSDirectory.Contract.Paused0(&_AVSDirectory.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_AVSDirectory *AVSDirectoryCallerSession) Paused0() (*big.Int, error) {
	return _AVSDirectory.Contract.Paused0(&_AVSDirectory.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_AVSDirectory *AVSDirectoryCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AVSDirectory.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_AVSDirectory *AVSDirectorySession) PauserRegistry() (common.Address, error) {
	return _AVSDirectory.Contract.PauserRegistry(&_AVSDirectory.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_AVSDirectory *AVSDirectoryCallerSession) PauserRegistry() (common.Address, error) {
	return _AVSDirectory.Contract.PauserRegistry(&_AVSDirectory.CallOpts)
}

// Allocate is a paid mutator transaction binding the contract method 0x70196708.
//
// Solidity: function allocate(address operator, (address,(address,uint32)[],uint64[])[] allocations, (bytes,bytes32,uint256) operatorSignature) returns()
func (_AVSDirectory *AVSDirectoryTransactor) Allocate(opts *bind.TransactOpts, operator common.Address, allocations []IAVSDirectoryMagnitudeAdjustment, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _AVSDirectory.contract.Transact(opts, "allocate", operator, allocations, operatorSignature)
}

// Allocate is a paid mutator transaction binding the contract method 0x70196708.
//
// Solidity: function allocate(address operator, (address,(address,uint32)[],uint64[])[] allocations, (bytes,bytes32,uint256) operatorSignature) returns()
func (_AVSDirectory *AVSDirectorySession) Allocate(operator common.Address, allocations []IAVSDirectoryMagnitudeAdjustment, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _AVSDirectory.Contract.Allocate(&_AVSDirectory.TransactOpts, operator, allocations, operatorSignature)
}

// Allocate is a paid mutator transaction binding the contract method 0x70196708.
//
// Solidity: function allocate(address operator, (address,(address,uint32)[],uint64[])[] allocations, (bytes,bytes32,uint256) operatorSignature) returns()
func (_AVSDirectory *AVSDirectoryTransactorSession) Allocate(operator common.Address, allocations []IAVSDirectoryMagnitudeAdjustment, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _AVSDirectory.Contract.Allocate(&_AVSDirectory.TransactOpts, operator, allocations, operatorSignature)
}

// BecomeOperatorSetAVS is a paid mutator transaction binding the contract method 0xaec205c5.
//
// Solidity: function becomeOperatorSetAVS() returns()
func (_AVSDirectory *AVSDirectoryTransactor) BecomeOperatorSetAVS(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AVSDirectory.contract.Transact(opts, "becomeOperatorSetAVS")
}

// BecomeOperatorSetAVS is a paid mutator transaction binding the contract method 0xaec205c5.
//
// Solidity: function becomeOperatorSetAVS() returns()
func (_AVSDirectory *AVSDirectorySession) BecomeOperatorSetAVS() (*types.Transaction, error) {
	return _AVSDirectory.Contract.BecomeOperatorSetAVS(&_AVSDirectory.TransactOpts)
}

// BecomeOperatorSetAVS is a paid mutator transaction binding the contract method 0xaec205c5.
//
// Solidity: function becomeOperatorSetAVS() returns()
func (_AVSDirectory *AVSDirectoryTransactorSession) BecomeOperatorSetAVS() (*types.Transaction, error) {
	return _AVSDirectory.Contract.BecomeOperatorSetAVS(&_AVSDirectory.TransactOpts)
}

// CancelSalt is a paid mutator transaction binding the contract method 0xec76f442.
//
// Solidity: function cancelSalt(bytes32 salt) returns()
func (_AVSDirectory *AVSDirectoryTransactor) CancelSalt(opts *bind.TransactOpts, salt [32]byte) (*types.Transaction, error) {
	return _AVSDirectory.contract.Transact(opts, "cancelSalt", salt)
}

// CancelSalt is a paid mutator transaction binding the contract method 0xec76f442.
//
// Solidity: function cancelSalt(bytes32 salt) returns()
func (_AVSDirectory *AVSDirectorySession) CancelSalt(salt [32]byte) (*types.Transaction, error) {
	return _AVSDirectory.Contract.CancelSalt(&_AVSDirectory.TransactOpts, salt)
}

// CancelSalt is a paid mutator transaction binding the contract method 0xec76f442.
//
// Solidity: function cancelSalt(bytes32 salt) returns()
func (_AVSDirectory *AVSDirectoryTransactorSession) CancelSalt(salt [32]byte) (*types.Transaction, error) {
	return _AVSDirectory.Contract.CancelSalt(&_AVSDirectory.TransactOpts, salt)
}

// CompleteDeallocations is a paid mutator transaction binding the contract method 0x5a922613.
//
// Solidity: function completeDeallocations(address operator, address[] strategies, (address,uint32)[][] operatorSets) returns()
func (_AVSDirectory *AVSDirectoryTransactor) CompleteDeallocations(opts *bind.TransactOpts, operator common.Address, strategies []common.Address, operatorSets [][]IAVSDirectoryOperatorSet) (*types.Transaction, error) {
	return _AVSDirectory.contract.Transact(opts, "completeDeallocations", operator, strategies, operatorSets)
}

// CompleteDeallocations is a paid mutator transaction binding the contract method 0x5a922613.
//
// Solidity: function completeDeallocations(address operator, address[] strategies, (address,uint32)[][] operatorSets) returns()
func (_AVSDirectory *AVSDirectorySession) CompleteDeallocations(operator common.Address, strategies []common.Address, operatorSets [][]IAVSDirectoryOperatorSet) (*types.Transaction, error) {
	return _AVSDirectory.Contract.CompleteDeallocations(&_AVSDirectory.TransactOpts, operator, strategies, operatorSets)
}

// CompleteDeallocations is a paid mutator transaction binding the contract method 0x5a922613.
//
// Solidity: function completeDeallocations(address operator, address[] strategies, (address,uint32)[][] operatorSets) returns()
func (_AVSDirectory *AVSDirectoryTransactorSession) CompleteDeallocations(operator common.Address, strategies []common.Address, operatorSets [][]IAVSDirectoryOperatorSet) (*types.Transaction, error) {
	return _AVSDirectory.Contract.CompleteDeallocations(&_AVSDirectory.TransactOpts, operator, strategies, operatorSets)
}

// CreateOperatorSets is a paid mutator transaction binding the contract method 0xafe02ed5.
//
// Solidity: function createOperatorSets(uint32[] operatorSetIds) returns()
func (_AVSDirectory *AVSDirectoryTransactor) CreateOperatorSets(opts *bind.TransactOpts, operatorSetIds []uint32) (*types.Transaction, error) {
	return _AVSDirectory.contract.Transact(opts, "createOperatorSets", operatorSetIds)
}

// CreateOperatorSets is a paid mutator transaction binding the contract method 0xafe02ed5.
//
// Solidity: function createOperatorSets(uint32[] operatorSetIds) returns()
func (_AVSDirectory *AVSDirectorySession) CreateOperatorSets(operatorSetIds []uint32) (*types.Transaction, error) {
	return _AVSDirectory.Contract.CreateOperatorSets(&_AVSDirectory.TransactOpts, operatorSetIds)
}

// CreateOperatorSets is a paid mutator transaction binding the contract method 0xafe02ed5.
//
// Solidity: function createOperatorSets(uint32[] operatorSetIds) returns()
func (_AVSDirectory *AVSDirectoryTransactorSession) CreateOperatorSets(operatorSetIds []uint32) (*types.Transaction, error) {
	return _AVSDirectory.Contract.CreateOperatorSets(&_AVSDirectory.TransactOpts, operatorSetIds)
}

// DeregisterOperatorFromOperatorSets is a paid mutator transaction binding the contract method 0xc1a8e2c5.
//
// Solidity: function deregisterOperatorFromOperatorSets(address operator, uint32[] operatorSetIds) returns()
func (_AVSDirectory *AVSDirectoryTransactor) DeregisterOperatorFromOperatorSets(opts *bind.TransactOpts, operator common.Address, operatorSetIds []uint32) (*types.Transaction, error) {
	return _AVSDirectory.contract.Transact(opts, "deregisterOperatorFromOperatorSets", operator, operatorSetIds)
}

// DeregisterOperatorFromOperatorSets is a paid mutator transaction binding the contract method 0xc1a8e2c5.
//
// Solidity: function deregisterOperatorFromOperatorSets(address operator, uint32[] operatorSetIds) returns()
func (_AVSDirectory *AVSDirectorySession) DeregisterOperatorFromOperatorSets(operator common.Address, operatorSetIds []uint32) (*types.Transaction, error) {
	return _AVSDirectory.Contract.DeregisterOperatorFromOperatorSets(&_AVSDirectory.TransactOpts, operator, operatorSetIds)
}

// DeregisterOperatorFromOperatorSets is a paid mutator transaction binding the contract method 0xc1a8e2c5.
//
// Solidity: function deregisterOperatorFromOperatorSets(address operator, uint32[] operatorSetIds) returns()
func (_AVSDirectory *AVSDirectoryTransactorSession) DeregisterOperatorFromOperatorSets(operator common.Address, operatorSetIds []uint32) (*types.Transaction, error) {
	return _AVSDirectory.Contract.DeregisterOperatorFromOperatorSets(&_AVSDirectory.TransactOpts, operator, operatorSetIds)
}

// ForceDeregisterFromOperatorSets is a paid mutator transaction binding the contract method 0x3fee332d.
//
// Solidity: function forceDeregisterFromOperatorSets(address operator, address avs, uint32[] operatorSetIds, (bytes,bytes32,uint256) operatorSignature) returns()
func (_AVSDirectory *AVSDirectoryTransactor) ForceDeregisterFromOperatorSets(opts *bind.TransactOpts, operator common.Address, avs common.Address, operatorSetIds []uint32, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _AVSDirectory.contract.Transact(opts, "forceDeregisterFromOperatorSets", operator, avs, operatorSetIds, operatorSignature)
}

// ForceDeregisterFromOperatorSets is a paid mutator transaction binding the contract method 0x3fee332d.
//
// Solidity: function forceDeregisterFromOperatorSets(address operator, address avs, uint32[] operatorSetIds, (bytes,bytes32,uint256) operatorSignature) returns()
func (_AVSDirectory *AVSDirectorySession) ForceDeregisterFromOperatorSets(operator common.Address, avs common.Address, operatorSetIds []uint32, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _AVSDirectory.Contract.ForceDeregisterFromOperatorSets(&_AVSDirectory.TransactOpts, operator, avs, operatorSetIds, operatorSignature)
}

// ForceDeregisterFromOperatorSets is a paid mutator transaction binding the contract method 0x3fee332d.
//
// Solidity: function forceDeregisterFromOperatorSets(address operator, address avs, uint32[] operatorSetIds, (bytes,bytes32,uint256) operatorSignature) returns()
func (_AVSDirectory *AVSDirectoryTransactorSession) ForceDeregisterFromOperatorSets(operator common.Address, avs common.Address, operatorSetIds []uint32, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _AVSDirectory.Contract.ForceDeregisterFromOperatorSets(&_AVSDirectory.TransactOpts, operator, avs, operatorSetIds, operatorSignature)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address initialOwner, address _pauserRegistry, uint256 initialPausedStatus) returns()
func (_AVSDirectory *AVSDirectoryTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, _pauserRegistry common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _AVSDirectory.contract.Transact(opts, "initialize", initialOwner, _pauserRegistry, initialPausedStatus)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address initialOwner, address _pauserRegistry, uint256 initialPausedStatus) returns()
func (_AVSDirectory *AVSDirectorySession) Initialize(initialOwner common.Address, _pauserRegistry common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _AVSDirectory.Contract.Initialize(&_AVSDirectory.TransactOpts, initialOwner, _pauserRegistry, initialPausedStatus)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address initialOwner, address _pauserRegistry, uint256 initialPausedStatus) returns()
func (_AVSDirectory *AVSDirectoryTransactorSession) Initialize(initialOwner common.Address, _pauserRegistry common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _AVSDirectory.Contract.Initialize(&_AVSDirectory.TransactOpts, initialOwner, _pauserRegistry, initialPausedStatus)
}

// MigrateOperatorsToOperatorSets is a paid mutator transaction binding the contract method 0xef2dfa8d.
//
// Solidity: function migrateOperatorsToOperatorSets(address[] operators, uint32[][] operatorSetIds) returns()
func (_AVSDirectory *AVSDirectoryTransactor) MigrateOperatorsToOperatorSets(opts *bind.TransactOpts, operators []common.Address, operatorSetIds [][]uint32) (*types.Transaction, error) {
	return _AVSDirectory.contract.Transact(opts, "migrateOperatorsToOperatorSets", operators, operatorSetIds)
}

// MigrateOperatorsToOperatorSets is a paid mutator transaction binding the contract method 0xef2dfa8d.
//
// Solidity: function migrateOperatorsToOperatorSets(address[] operators, uint32[][] operatorSetIds) returns()
func (_AVSDirectory *AVSDirectorySession) MigrateOperatorsToOperatorSets(operators []common.Address, operatorSetIds [][]uint32) (*types.Transaction, error) {
	return _AVSDirectory.Contract.MigrateOperatorsToOperatorSets(&_AVSDirectory.TransactOpts, operators, operatorSetIds)
}

// MigrateOperatorsToOperatorSets is a paid mutator transaction binding the contract method 0xef2dfa8d.
//
// Solidity: function migrateOperatorsToOperatorSets(address[] operators, uint32[][] operatorSetIds) returns()
func (_AVSDirectory *AVSDirectoryTransactorSession) MigrateOperatorsToOperatorSets(operators []common.Address, operatorSetIds [][]uint32) (*types.Transaction, error) {
	return _AVSDirectory.Contract.MigrateOperatorsToOperatorSets(&_AVSDirectory.TransactOpts, operators, operatorSetIds)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_AVSDirectory *AVSDirectoryTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _AVSDirectory.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_AVSDirectory *AVSDirectorySession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _AVSDirectory.Contract.Pause(&_AVSDirectory.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_AVSDirectory *AVSDirectoryTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _AVSDirectory.Contract.Pause(&_AVSDirectory.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_AVSDirectory *AVSDirectoryTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AVSDirectory.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_AVSDirectory *AVSDirectorySession) PauseAll() (*types.Transaction, error) {
	return _AVSDirectory.Contract.PauseAll(&_AVSDirectory.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_AVSDirectory *AVSDirectoryTransactorSession) PauseAll() (*types.Transaction, error) {
	return _AVSDirectory.Contract.PauseAll(&_AVSDirectory.TransactOpts)
}

// QueueDeallocate is a paid mutator transaction binding the contract method 0x4ef8c5d4.
//
// Solidity: function queueDeallocate(address operator, (address,(address,uint32)[],uint64[])[] deallocations, (bytes,bytes32,uint256) operatorSignature) returns()
func (_AVSDirectory *AVSDirectoryTransactor) QueueDeallocate(opts *bind.TransactOpts, operator common.Address, deallocations []IAVSDirectoryMagnitudeAdjustment, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _AVSDirectory.contract.Transact(opts, "queueDeallocate", operator, deallocations, operatorSignature)
}

// QueueDeallocate is a paid mutator transaction binding the contract method 0x4ef8c5d4.
//
// Solidity: function queueDeallocate(address operator, (address,(address,uint32)[],uint64[])[] deallocations, (bytes,bytes32,uint256) operatorSignature) returns()
func (_AVSDirectory *AVSDirectorySession) QueueDeallocate(operator common.Address, deallocations []IAVSDirectoryMagnitudeAdjustment, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _AVSDirectory.Contract.QueueDeallocate(&_AVSDirectory.TransactOpts, operator, deallocations, operatorSignature)
}

// QueueDeallocate is a paid mutator transaction binding the contract method 0x4ef8c5d4.
//
// Solidity: function queueDeallocate(address operator, (address,(address,uint32)[],uint64[])[] deallocations, (bytes,bytes32,uint256) operatorSignature) returns()
func (_AVSDirectory *AVSDirectoryTransactorSession) QueueDeallocate(operator common.Address, deallocations []IAVSDirectoryMagnitudeAdjustment, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _AVSDirectory.Contract.QueueDeallocate(&_AVSDirectory.TransactOpts, operator, deallocations, operatorSignature)
}

// RegisterOperatorToOperatorSets is a paid mutator transaction binding the contract method 0x1e2199e2.
//
// Solidity: function registerOperatorToOperatorSets(address operator, uint32[] operatorSetIds, (bytes,bytes32,uint256) operatorSignature) returns()
func (_AVSDirectory *AVSDirectoryTransactor) RegisterOperatorToOperatorSets(opts *bind.TransactOpts, operator common.Address, operatorSetIds []uint32, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _AVSDirectory.contract.Transact(opts, "registerOperatorToOperatorSets", operator, operatorSetIds, operatorSignature)
}

// RegisterOperatorToOperatorSets is a paid mutator transaction binding the contract method 0x1e2199e2.
//
// Solidity: function registerOperatorToOperatorSets(address operator, uint32[] operatorSetIds, (bytes,bytes32,uint256) operatorSignature) returns()
func (_AVSDirectory *AVSDirectorySession) RegisterOperatorToOperatorSets(operator common.Address, operatorSetIds []uint32, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _AVSDirectory.Contract.RegisterOperatorToOperatorSets(&_AVSDirectory.TransactOpts, operator, operatorSetIds, operatorSignature)
}

// RegisterOperatorToOperatorSets is a paid mutator transaction binding the contract method 0x1e2199e2.
//
// Solidity: function registerOperatorToOperatorSets(address operator, uint32[] operatorSetIds, (bytes,bytes32,uint256) operatorSignature) returns()
func (_AVSDirectory *AVSDirectoryTransactorSession) RegisterOperatorToOperatorSets(operator common.Address, operatorSetIds []uint32, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _AVSDirectory.Contract.RegisterOperatorToOperatorSets(&_AVSDirectory.TransactOpts, operator, operatorSetIds, operatorSignature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AVSDirectory *AVSDirectoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AVSDirectory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AVSDirectory *AVSDirectorySession) RenounceOwnership() (*types.Transaction, error) {
	return _AVSDirectory.Contract.RenounceOwnership(&_AVSDirectory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AVSDirectory *AVSDirectoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AVSDirectory.Contract.RenounceOwnership(&_AVSDirectory.TransactOpts)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_AVSDirectory *AVSDirectoryTransactor) SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error) {
	return _AVSDirectory.contract.Transact(opts, "setPauserRegistry", newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_AVSDirectory *AVSDirectorySession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _AVSDirectory.Contract.SetPauserRegistry(&_AVSDirectory.TransactOpts, newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_AVSDirectory *AVSDirectoryTransactorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _AVSDirectory.Contract.SetPauserRegistry(&_AVSDirectory.TransactOpts, newPauserRegistry)
}

// SlashOperator is a paid mutator transaction binding the contract method 0xbd74a06c.
//
// Solidity: function slashOperator(address operator, uint32 operatorSetId, address[] strategies, uint16 bipsToSlash) returns()
func (_AVSDirectory *AVSDirectoryTransactor) SlashOperator(opts *bind.TransactOpts, operator common.Address, operatorSetId uint32, strategies []common.Address, bipsToSlash uint16) (*types.Transaction, error) {
	return _AVSDirectory.contract.Transact(opts, "slashOperator", operator, operatorSetId, strategies, bipsToSlash)
}

// SlashOperator is a paid mutator transaction binding the contract method 0xbd74a06c.
//
// Solidity: function slashOperator(address operator, uint32 operatorSetId, address[] strategies, uint16 bipsToSlash) returns()
func (_AVSDirectory *AVSDirectorySession) SlashOperator(operator common.Address, operatorSetId uint32, strategies []common.Address, bipsToSlash uint16) (*types.Transaction, error) {
	return _AVSDirectory.Contract.SlashOperator(&_AVSDirectory.TransactOpts, operator, operatorSetId, strategies, bipsToSlash)
}

// SlashOperator is a paid mutator transaction binding the contract method 0xbd74a06c.
//
// Solidity: function slashOperator(address operator, uint32 operatorSetId, address[] strategies, uint16 bipsToSlash) returns()
func (_AVSDirectory *AVSDirectoryTransactorSession) SlashOperator(operator common.Address, operatorSetId uint32, strategies []common.Address, bipsToSlash uint16) (*types.Transaction, error) {
	return _AVSDirectory.Contract.SlashOperator(&_AVSDirectory.TransactOpts, operator, operatorSetId, strategies, bipsToSlash)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AVSDirectory *AVSDirectoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AVSDirectory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AVSDirectory *AVSDirectorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AVSDirectory.Contract.TransferOwnership(&_AVSDirectory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AVSDirectory *AVSDirectoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AVSDirectory.Contract.TransferOwnership(&_AVSDirectory.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_AVSDirectory *AVSDirectoryTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _AVSDirectory.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_AVSDirectory *AVSDirectorySession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _AVSDirectory.Contract.Unpause(&_AVSDirectory.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_AVSDirectory *AVSDirectoryTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _AVSDirectory.Contract.Unpause(&_AVSDirectory.TransactOpts, newPausedStatus)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string metadataURI) returns()
func (_AVSDirectory *AVSDirectoryTransactor) UpdateAVSMetadataURI(opts *bind.TransactOpts, metadataURI string) (*types.Transaction, error) {
	return _AVSDirectory.contract.Transact(opts, "updateAVSMetadataURI", metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string metadataURI) returns()
func (_AVSDirectory *AVSDirectorySession) UpdateAVSMetadataURI(metadataURI string) (*types.Transaction, error) {
	return _AVSDirectory.Contract.UpdateAVSMetadataURI(&_AVSDirectory.TransactOpts, metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string metadataURI) returns()
func (_AVSDirectory *AVSDirectoryTransactorSession) UpdateAVSMetadataURI(metadataURI string) (*types.Transaction, error) {
	return _AVSDirectory.Contract.UpdateAVSMetadataURI(&_AVSDirectory.TransactOpts, metadataURI)
}

// AVSDirectoryAVSMetadataURIUpdatedIterator is returned from FilterAVSMetadataURIUpdated and is used to iterate over the raw logs and unpacked data for AVSMetadataURIUpdated events raised by the AVSDirectory contract.
type AVSDirectoryAVSMetadataURIUpdatedIterator struct {
	Event *AVSDirectoryAVSMetadataURIUpdated // Event containing the contract specifics and raw log

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
func (it *AVSDirectoryAVSMetadataURIUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSDirectoryAVSMetadataURIUpdated)
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
		it.Event = new(AVSDirectoryAVSMetadataURIUpdated)
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
func (it *AVSDirectoryAVSMetadataURIUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSDirectoryAVSMetadataURIUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSDirectoryAVSMetadataURIUpdated represents a AVSMetadataURIUpdated event raised by the AVSDirectory contract.
type AVSDirectoryAVSMetadataURIUpdated struct {
	Avs         common.Address
	MetadataURI string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAVSMetadataURIUpdated is a free log retrieval operation binding the contract event 0xa89c1dc243d8908a96dd84944bcc97d6bc6ac00dd78e20621576be6a3c943713.
//
// Solidity: event AVSMetadataURIUpdated(address indexed avs, string metadataURI)
func (_AVSDirectory *AVSDirectoryFilterer) FilterAVSMetadataURIUpdated(opts *bind.FilterOpts, avs []common.Address) (*AVSDirectoryAVSMetadataURIUpdatedIterator, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _AVSDirectory.contract.FilterLogs(opts, "AVSMetadataURIUpdated", avsRule)
	if err != nil {
		return nil, err
	}
	return &AVSDirectoryAVSMetadataURIUpdatedIterator{contract: _AVSDirectory.contract, event: "AVSMetadataURIUpdated", logs: logs, sub: sub}, nil
}

// WatchAVSMetadataURIUpdated is a free log subscription operation binding the contract event 0xa89c1dc243d8908a96dd84944bcc97d6bc6ac00dd78e20621576be6a3c943713.
//
// Solidity: event AVSMetadataURIUpdated(address indexed avs, string metadataURI)
func (_AVSDirectory *AVSDirectoryFilterer) WatchAVSMetadataURIUpdated(opts *bind.WatchOpts, sink chan<- *AVSDirectoryAVSMetadataURIUpdated, avs []common.Address) (event.Subscription, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _AVSDirectory.contract.WatchLogs(opts, "AVSMetadataURIUpdated", avsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSDirectoryAVSMetadataURIUpdated)
				if err := _AVSDirectory.contract.UnpackLog(event, "AVSMetadataURIUpdated", log); err != nil {
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
func (_AVSDirectory *AVSDirectoryFilterer) ParseAVSMetadataURIUpdated(log types.Log) (*AVSDirectoryAVSMetadataURIUpdated, error) {
	event := new(AVSDirectoryAVSMetadataURIUpdated)
	if err := _AVSDirectory.contract.UnpackLog(event, "AVSMetadataURIUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSDirectoryAVSMigratedToOperatorSetsIterator is returned from FilterAVSMigratedToOperatorSets and is used to iterate over the raw logs and unpacked data for AVSMigratedToOperatorSets events raised by the AVSDirectory contract.
type AVSDirectoryAVSMigratedToOperatorSetsIterator struct {
	Event *AVSDirectoryAVSMigratedToOperatorSets // Event containing the contract specifics and raw log

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
func (it *AVSDirectoryAVSMigratedToOperatorSetsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSDirectoryAVSMigratedToOperatorSets)
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
		it.Event = new(AVSDirectoryAVSMigratedToOperatorSets)
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
func (it *AVSDirectoryAVSMigratedToOperatorSetsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSDirectoryAVSMigratedToOperatorSetsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSDirectoryAVSMigratedToOperatorSets represents a AVSMigratedToOperatorSets event raised by the AVSDirectory contract.
type AVSDirectoryAVSMigratedToOperatorSets struct {
	Avs common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAVSMigratedToOperatorSets is a free log retrieval operation binding the contract event 0x702b0c1f6cb1cf511aaa81f72bc05a215bb3497632d72c690c822b044ab494bf.
//
// Solidity: event AVSMigratedToOperatorSets(address indexed avs)
func (_AVSDirectory *AVSDirectoryFilterer) FilterAVSMigratedToOperatorSets(opts *bind.FilterOpts, avs []common.Address) (*AVSDirectoryAVSMigratedToOperatorSetsIterator, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _AVSDirectory.contract.FilterLogs(opts, "AVSMigratedToOperatorSets", avsRule)
	if err != nil {
		return nil, err
	}
	return &AVSDirectoryAVSMigratedToOperatorSetsIterator{contract: _AVSDirectory.contract, event: "AVSMigratedToOperatorSets", logs: logs, sub: sub}, nil
}

// WatchAVSMigratedToOperatorSets is a free log subscription operation binding the contract event 0x702b0c1f6cb1cf511aaa81f72bc05a215bb3497632d72c690c822b044ab494bf.
//
// Solidity: event AVSMigratedToOperatorSets(address indexed avs)
func (_AVSDirectory *AVSDirectoryFilterer) WatchAVSMigratedToOperatorSets(opts *bind.WatchOpts, sink chan<- *AVSDirectoryAVSMigratedToOperatorSets, avs []common.Address) (event.Subscription, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _AVSDirectory.contract.WatchLogs(opts, "AVSMigratedToOperatorSets", avsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSDirectoryAVSMigratedToOperatorSets)
				if err := _AVSDirectory.contract.UnpackLog(event, "AVSMigratedToOperatorSets", log); err != nil {
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
func (_AVSDirectory *AVSDirectoryFilterer) ParseAVSMigratedToOperatorSets(log types.Log) (*AVSDirectoryAVSMigratedToOperatorSets, error) {
	event := new(AVSDirectoryAVSMigratedToOperatorSets)
	if err := _AVSDirectory.contract.UnpackLog(event, "AVSMigratedToOperatorSets", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSDirectoryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the AVSDirectory contract.
type AVSDirectoryInitializedIterator struct {
	Event *AVSDirectoryInitialized // Event containing the contract specifics and raw log

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
func (it *AVSDirectoryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSDirectoryInitialized)
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
		it.Event = new(AVSDirectoryInitialized)
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
func (it *AVSDirectoryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSDirectoryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSDirectoryInitialized represents a Initialized event raised by the AVSDirectory contract.
type AVSDirectoryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AVSDirectory *AVSDirectoryFilterer) FilterInitialized(opts *bind.FilterOpts) (*AVSDirectoryInitializedIterator, error) {

	logs, sub, err := _AVSDirectory.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AVSDirectoryInitializedIterator{contract: _AVSDirectory.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AVSDirectory *AVSDirectoryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AVSDirectoryInitialized) (event.Subscription, error) {

	logs, sub, err := _AVSDirectory.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSDirectoryInitialized)
				if err := _AVSDirectory.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AVSDirectory *AVSDirectoryFilterer) ParseInitialized(log types.Log) (*AVSDirectoryInitialized, error) {
	event := new(AVSDirectoryInitialized)
	if err := _AVSDirectory.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSDirectoryMagnitudeAllocatedIterator is returned from FilterMagnitudeAllocated and is used to iterate over the raw logs and unpacked data for MagnitudeAllocated events raised by the AVSDirectory contract.
type AVSDirectoryMagnitudeAllocatedIterator struct {
	Event *AVSDirectoryMagnitudeAllocated // Event containing the contract specifics and raw log

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
func (it *AVSDirectoryMagnitudeAllocatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSDirectoryMagnitudeAllocated)
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
		it.Event = new(AVSDirectoryMagnitudeAllocated)
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
func (it *AVSDirectoryMagnitudeAllocatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSDirectoryMagnitudeAllocatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSDirectoryMagnitudeAllocated represents a MagnitudeAllocated event raised by the AVSDirectory contract.
type AVSDirectoryMagnitudeAllocated struct {
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
func (_AVSDirectory *AVSDirectoryFilterer) FilterMagnitudeAllocated(opts *bind.FilterOpts) (*AVSDirectoryMagnitudeAllocatedIterator, error) {

	logs, sub, err := _AVSDirectory.contract.FilterLogs(opts, "MagnitudeAllocated")
	if err != nil {
		return nil, err
	}
	return &AVSDirectoryMagnitudeAllocatedIterator{contract: _AVSDirectory.contract, event: "MagnitudeAllocated", logs: logs, sub: sub}, nil
}

// WatchMagnitudeAllocated is a free log subscription operation binding the contract event 0x6d7d0079582cb2c5e70d4135b37f36711415ee6c260778b716bd65e026eb4f1a.
//
// Solidity: event MagnitudeAllocated(address operator, address strategy, (address,uint32) operatorSet, uint64 magnitudeToAllocate, uint32 effectTimestamp)
func (_AVSDirectory *AVSDirectoryFilterer) WatchMagnitudeAllocated(opts *bind.WatchOpts, sink chan<- *AVSDirectoryMagnitudeAllocated) (event.Subscription, error) {

	logs, sub, err := _AVSDirectory.contract.WatchLogs(opts, "MagnitudeAllocated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSDirectoryMagnitudeAllocated)
				if err := _AVSDirectory.contract.UnpackLog(event, "MagnitudeAllocated", log); err != nil {
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
func (_AVSDirectory *AVSDirectoryFilterer) ParseMagnitudeAllocated(log types.Log) (*AVSDirectoryMagnitudeAllocated, error) {
	event := new(AVSDirectoryMagnitudeAllocated)
	if err := _AVSDirectory.contract.UnpackLog(event, "MagnitudeAllocated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSDirectoryMagnitudeDeallocationCompletedIterator is returned from FilterMagnitudeDeallocationCompleted and is used to iterate over the raw logs and unpacked data for MagnitudeDeallocationCompleted events raised by the AVSDirectory contract.
type AVSDirectoryMagnitudeDeallocationCompletedIterator struct {
	Event *AVSDirectoryMagnitudeDeallocationCompleted // Event containing the contract specifics and raw log

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
func (it *AVSDirectoryMagnitudeDeallocationCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSDirectoryMagnitudeDeallocationCompleted)
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
		it.Event = new(AVSDirectoryMagnitudeDeallocationCompleted)
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
func (it *AVSDirectoryMagnitudeDeallocationCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSDirectoryMagnitudeDeallocationCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSDirectoryMagnitudeDeallocationCompleted represents a MagnitudeDeallocationCompleted event raised by the AVSDirectory contract.
type AVSDirectoryMagnitudeDeallocationCompleted struct {
	Operator           common.Address
	Strategy           common.Address
	OperatorSet        IAVSDirectoryOperatorSet
	FreeMagnitudeAdded uint64
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMagnitudeDeallocationCompleted is a free log retrieval operation binding the contract event 0x1e5c8e13c62c31d6252ac205e592477d643c7e95831d5b46d99a3c60c2fad8db.
//
// Solidity: event MagnitudeDeallocationCompleted(address operator, address strategy, (address,uint32) operatorSet, uint64 freeMagnitudeAdded)
func (_AVSDirectory *AVSDirectoryFilterer) FilterMagnitudeDeallocationCompleted(opts *bind.FilterOpts) (*AVSDirectoryMagnitudeDeallocationCompletedIterator, error) {

	logs, sub, err := _AVSDirectory.contract.FilterLogs(opts, "MagnitudeDeallocationCompleted")
	if err != nil {
		return nil, err
	}
	return &AVSDirectoryMagnitudeDeallocationCompletedIterator{contract: _AVSDirectory.contract, event: "MagnitudeDeallocationCompleted", logs: logs, sub: sub}, nil
}

// WatchMagnitudeDeallocationCompleted is a free log subscription operation binding the contract event 0x1e5c8e13c62c31d6252ac205e592477d643c7e95831d5b46d99a3c60c2fad8db.
//
// Solidity: event MagnitudeDeallocationCompleted(address operator, address strategy, (address,uint32) operatorSet, uint64 freeMagnitudeAdded)
func (_AVSDirectory *AVSDirectoryFilterer) WatchMagnitudeDeallocationCompleted(opts *bind.WatchOpts, sink chan<- *AVSDirectoryMagnitudeDeallocationCompleted) (event.Subscription, error) {

	logs, sub, err := _AVSDirectory.contract.WatchLogs(opts, "MagnitudeDeallocationCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSDirectoryMagnitudeDeallocationCompleted)
				if err := _AVSDirectory.contract.UnpackLog(event, "MagnitudeDeallocationCompleted", log); err != nil {
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
func (_AVSDirectory *AVSDirectoryFilterer) ParseMagnitudeDeallocationCompleted(log types.Log) (*AVSDirectoryMagnitudeDeallocationCompleted, error) {
	event := new(AVSDirectoryMagnitudeDeallocationCompleted)
	if err := _AVSDirectory.contract.UnpackLog(event, "MagnitudeDeallocationCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSDirectoryMagnitudeQueueDeallocatedIterator is returned from FilterMagnitudeQueueDeallocated and is used to iterate over the raw logs and unpacked data for MagnitudeQueueDeallocated events raised by the AVSDirectory contract.
type AVSDirectoryMagnitudeQueueDeallocatedIterator struct {
	Event *AVSDirectoryMagnitudeQueueDeallocated // Event containing the contract specifics and raw log

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
func (it *AVSDirectoryMagnitudeQueueDeallocatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSDirectoryMagnitudeQueueDeallocated)
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
		it.Event = new(AVSDirectoryMagnitudeQueueDeallocated)
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
func (it *AVSDirectoryMagnitudeQueueDeallocatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSDirectoryMagnitudeQueueDeallocatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSDirectoryMagnitudeQueueDeallocated represents a MagnitudeQueueDeallocated event raised by the AVSDirectory contract.
type AVSDirectoryMagnitudeQueueDeallocated struct {
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
func (_AVSDirectory *AVSDirectoryFilterer) FilterMagnitudeQueueDeallocated(opts *bind.FilterOpts) (*AVSDirectoryMagnitudeQueueDeallocatedIterator, error) {

	logs, sub, err := _AVSDirectory.contract.FilterLogs(opts, "MagnitudeQueueDeallocated")
	if err != nil {
		return nil, err
	}
	return &AVSDirectoryMagnitudeQueueDeallocatedIterator{contract: _AVSDirectory.contract, event: "MagnitudeQueueDeallocated", logs: logs, sub: sub}, nil
}

// WatchMagnitudeQueueDeallocated is a free log subscription operation binding the contract event 0x2e68db1fe51107d7e451ae268d1631796989ab9d7925054e9b247854cb5be950.
//
// Solidity: event MagnitudeQueueDeallocated(address operator, address strategy, (address,uint32) operatorSet, uint64 magnitudeToDeallocate, uint32 completableTimestamp)
func (_AVSDirectory *AVSDirectoryFilterer) WatchMagnitudeQueueDeallocated(opts *bind.WatchOpts, sink chan<- *AVSDirectoryMagnitudeQueueDeallocated) (event.Subscription, error) {

	logs, sub, err := _AVSDirectory.contract.WatchLogs(opts, "MagnitudeQueueDeallocated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSDirectoryMagnitudeQueueDeallocated)
				if err := _AVSDirectory.contract.UnpackLog(event, "MagnitudeQueueDeallocated", log); err != nil {
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
func (_AVSDirectory *AVSDirectoryFilterer) ParseMagnitudeQueueDeallocated(log types.Log) (*AVSDirectoryMagnitudeQueueDeallocated, error) {
	event := new(AVSDirectoryMagnitudeQueueDeallocated)
	if err := _AVSDirectory.contract.UnpackLog(event, "MagnitudeQueueDeallocated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSDirectoryOperatorAVSRegistrationStatusUpdatedIterator is returned from FilterOperatorAVSRegistrationStatusUpdated and is used to iterate over the raw logs and unpacked data for OperatorAVSRegistrationStatusUpdated events raised by the AVSDirectory contract.
type AVSDirectoryOperatorAVSRegistrationStatusUpdatedIterator struct {
	Event *AVSDirectoryOperatorAVSRegistrationStatusUpdated // Event containing the contract specifics and raw log

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
func (it *AVSDirectoryOperatorAVSRegistrationStatusUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSDirectoryOperatorAVSRegistrationStatusUpdated)
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
		it.Event = new(AVSDirectoryOperatorAVSRegistrationStatusUpdated)
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
func (it *AVSDirectoryOperatorAVSRegistrationStatusUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSDirectoryOperatorAVSRegistrationStatusUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSDirectoryOperatorAVSRegistrationStatusUpdated represents a OperatorAVSRegistrationStatusUpdated event raised by the AVSDirectory contract.
type AVSDirectoryOperatorAVSRegistrationStatusUpdated struct {
	Operator common.Address
	Avs      common.Address
	Status   uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorAVSRegistrationStatusUpdated is a free log retrieval operation binding the contract event 0xf0952b1c65271d819d39983d2abb044b9cace59bcc4d4dd389f586ebdcb15b41.
//
// Solidity: event OperatorAVSRegistrationStatusUpdated(address indexed operator, address indexed avs, uint8 status)
func (_AVSDirectory *AVSDirectoryFilterer) FilterOperatorAVSRegistrationStatusUpdated(opts *bind.FilterOpts, operator []common.Address, avs []common.Address) (*AVSDirectoryOperatorAVSRegistrationStatusUpdatedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _AVSDirectory.contract.FilterLogs(opts, "OperatorAVSRegistrationStatusUpdated", operatorRule, avsRule)
	if err != nil {
		return nil, err
	}
	return &AVSDirectoryOperatorAVSRegistrationStatusUpdatedIterator{contract: _AVSDirectory.contract, event: "OperatorAVSRegistrationStatusUpdated", logs: logs, sub: sub}, nil
}

// WatchOperatorAVSRegistrationStatusUpdated is a free log subscription operation binding the contract event 0xf0952b1c65271d819d39983d2abb044b9cace59bcc4d4dd389f586ebdcb15b41.
//
// Solidity: event OperatorAVSRegistrationStatusUpdated(address indexed operator, address indexed avs, uint8 status)
func (_AVSDirectory *AVSDirectoryFilterer) WatchOperatorAVSRegistrationStatusUpdated(opts *bind.WatchOpts, sink chan<- *AVSDirectoryOperatorAVSRegistrationStatusUpdated, operator []common.Address, avs []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _AVSDirectory.contract.WatchLogs(opts, "OperatorAVSRegistrationStatusUpdated", operatorRule, avsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSDirectoryOperatorAVSRegistrationStatusUpdated)
				if err := _AVSDirectory.contract.UnpackLog(event, "OperatorAVSRegistrationStatusUpdated", log); err != nil {
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
func (_AVSDirectory *AVSDirectoryFilterer) ParseOperatorAVSRegistrationStatusUpdated(log types.Log) (*AVSDirectoryOperatorAVSRegistrationStatusUpdated, error) {
	event := new(AVSDirectoryOperatorAVSRegistrationStatusUpdated)
	if err := _AVSDirectory.contract.UnpackLog(event, "OperatorAVSRegistrationStatusUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSDirectoryOperatorAddedToOperatorSetIterator is returned from FilterOperatorAddedToOperatorSet and is used to iterate over the raw logs and unpacked data for OperatorAddedToOperatorSet events raised by the AVSDirectory contract.
type AVSDirectoryOperatorAddedToOperatorSetIterator struct {
	Event *AVSDirectoryOperatorAddedToOperatorSet // Event containing the contract specifics and raw log

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
func (it *AVSDirectoryOperatorAddedToOperatorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSDirectoryOperatorAddedToOperatorSet)
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
		it.Event = new(AVSDirectoryOperatorAddedToOperatorSet)
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
func (it *AVSDirectoryOperatorAddedToOperatorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSDirectoryOperatorAddedToOperatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSDirectoryOperatorAddedToOperatorSet represents a OperatorAddedToOperatorSet event raised by the AVSDirectory contract.
type AVSDirectoryOperatorAddedToOperatorSet struct {
	Operator    common.Address
	OperatorSet IAVSDirectoryOperatorSet
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorAddedToOperatorSet is a free log retrieval operation binding the contract event 0x43232edf9071753d2321e5fa7e018363ee248e5f2142e6c08edd3265bfb4895e.
//
// Solidity: event OperatorAddedToOperatorSet(address indexed operator, (address,uint32) operatorSet)
func (_AVSDirectory *AVSDirectoryFilterer) FilterOperatorAddedToOperatorSet(opts *bind.FilterOpts, operator []common.Address) (*AVSDirectoryOperatorAddedToOperatorSetIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _AVSDirectory.contract.FilterLogs(opts, "OperatorAddedToOperatorSet", operatorRule)
	if err != nil {
		return nil, err
	}
	return &AVSDirectoryOperatorAddedToOperatorSetIterator{contract: _AVSDirectory.contract, event: "OperatorAddedToOperatorSet", logs: logs, sub: sub}, nil
}

// WatchOperatorAddedToOperatorSet is a free log subscription operation binding the contract event 0x43232edf9071753d2321e5fa7e018363ee248e5f2142e6c08edd3265bfb4895e.
//
// Solidity: event OperatorAddedToOperatorSet(address indexed operator, (address,uint32) operatorSet)
func (_AVSDirectory *AVSDirectoryFilterer) WatchOperatorAddedToOperatorSet(opts *bind.WatchOpts, sink chan<- *AVSDirectoryOperatorAddedToOperatorSet, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _AVSDirectory.contract.WatchLogs(opts, "OperatorAddedToOperatorSet", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSDirectoryOperatorAddedToOperatorSet)
				if err := _AVSDirectory.contract.UnpackLog(event, "OperatorAddedToOperatorSet", log); err != nil {
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
func (_AVSDirectory *AVSDirectoryFilterer) ParseOperatorAddedToOperatorSet(log types.Log) (*AVSDirectoryOperatorAddedToOperatorSet, error) {
	event := new(AVSDirectoryOperatorAddedToOperatorSet)
	if err := _AVSDirectory.contract.UnpackLog(event, "OperatorAddedToOperatorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSDirectoryOperatorMigratedToOperatorSetsIterator is returned from FilterOperatorMigratedToOperatorSets and is used to iterate over the raw logs and unpacked data for OperatorMigratedToOperatorSets events raised by the AVSDirectory contract.
type AVSDirectoryOperatorMigratedToOperatorSetsIterator struct {
	Event *AVSDirectoryOperatorMigratedToOperatorSets // Event containing the contract specifics and raw log

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
func (it *AVSDirectoryOperatorMigratedToOperatorSetsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSDirectoryOperatorMigratedToOperatorSets)
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
		it.Event = new(AVSDirectoryOperatorMigratedToOperatorSets)
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
func (it *AVSDirectoryOperatorMigratedToOperatorSetsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSDirectoryOperatorMigratedToOperatorSetsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSDirectoryOperatorMigratedToOperatorSets represents a OperatorMigratedToOperatorSets event raised by the AVSDirectory contract.
type AVSDirectoryOperatorMigratedToOperatorSets struct {
	Operator       common.Address
	Avs            common.Address
	OperatorSetIds []uint32
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOperatorMigratedToOperatorSets is a free log retrieval operation binding the contract event 0x54f33cfdd1ca703d795986b986fd47d742eab1904ecd2a5fdb8d6595e5904a01.
//
// Solidity: event OperatorMigratedToOperatorSets(address indexed operator, address indexed avs, uint32[] operatorSetIds)
func (_AVSDirectory *AVSDirectoryFilterer) FilterOperatorMigratedToOperatorSets(opts *bind.FilterOpts, operator []common.Address, avs []common.Address) (*AVSDirectoryOperatorMigratedToOperatorSetsIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _AVSDirectory.contract.FilterLogs(opts, "OperatorMigratedToOperatorSets", operatorRule, avsRule)
	if err != nil {
		return nil, err
	}
	return &AVSDirectoryOperatorMigratedToOperatorSetsIterator{contract: _AVSDirectory.contract, event: "OperatorMigratedToOperatorSets", logs: logs, sub: sub}, nil
}

// WatchOperatorMigratedToOperatorSets is a free log subscription operation binding the contract event 0x54f33cfdd1ca703d795986b986fd47d742eab1904ecd2a5fdb8d6595e5904a01.
//
// Solidity: event OperatorMigratedToOperatorSets(address indexed operator, address indexed avs, uint32[] operatorSetIds)
func (_AVSDirectory *AVSDirectoryFilterer) WatchOperatorMigratedToOperatorSets(opts *bind.WatchOpts, sink chan<- *AVSDirectoryOperatorMigratedToOperatorSets, operator []common.Address, avs []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _AVSDirectory.contract.WatchLogs(opts, "OperatorMigratedToOperatorSets", operatorRule, avsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSDirectoryOperatorMigratedToOperatorSets)
				if err := _AVSDirectory.contract.UnpackLog(event, "OperatorMigratedToOperatorSets", log); err != nil {
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
func (_AVSDirectory *AVSDirectoryFilterer) ParseOperatorMigratedToOperatorSets(log types.Log) (*AVSDirectoryOperatorMigratedToOperatorSets, error) {
	event := new(AVSDirectoryOperatorMigratedToOperatorSets)
	if err := _AVSDirectory.contract.UnpackLog(event, "OperatorMigratedToOperatorSets", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSDirectoryOperatorRemovedFromOperatorSetIterator is returned from FilterOperatorRemovedFromOperatorSet and is used to iterate over the raw logs and unpacked data for OperatorRemovedFromOperatorSet events raised by the AVSDirectory contract.
type AVSDirectoryOperatorRemovedFromOperatorSetIterator struct {
	Event *AVSDirectoryOperatorRemovedFromOperatorSet // Event containing the contract specifics and raw log

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
func (it *AVSDirectoryOperatorRemovedFromOperatorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSDirectoryOperatorRemovedFromOperatorSet)
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
		it.Event = new(AVSDirectoryOperatorRemovedFromOperatorSet)
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
func (it *AVSDirectoryOperatorRemovedFromOperatorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSDirectoryOperatorRemovedFromOperatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSDirectoryOperatorRemovedFromOperatorSet represents a OperatorRemovedFromOperatorSet event raised by the AVSDirectory contract.
type AVSDirectoryOperatorRemovedFromOperatorSet struct {
	Operator    common.Address
	OperatorSet IAVSDirectoryOperatorSet
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorRemovedFromOperatorSet is a free log retrieval operation binding the contract event 0xad34c3070be1dffbcaa499d000ba2b8d9848aefcac3059df245dd95c4ece14fe.
//
// Solidity: event OperatorRemovedFromOperatorSet(address indexed operator, (address,uint32) operatorSet)
func (_AVSDirectory *AVSDirectoryFilterer) FilterOperatorRemovedFromOperatorSet(opts *bind.FilterOpts, operator []common.Address) (*AVSDirectoryOperatorRemovedFromOperatorSetIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _AVSDirectory.contract.FilterLogs(opts, "OperatorRemovedFromOperatorSet", operatorRule)
	if err != nil {
		return nil, err
	}
	return &AVSDirectoryOperatorRemovedFromOperatorSetIterator{contract: _AVSDirectory.contract, event: "OperatorRemovedFromOperatorSet", logs: logs, sub: sub}, nil
}

// WatchOperatorRemovedFromOperatorSet is a free log subscription operation binding the contract event 0xad34c3070be1dffbcaa499d000ba2b8d9848aefcac3059df245dd95c4ece14fe.
//
// Solidity: event OperatorRemovedFromOperatorSet(address indexed operator, (address,uint32) operatorSet)
func (_AVSDirectory *AVSDirectoryFilterer) WatchOperatorRemovedFromOperatorSet(opts *bind.WatchOpts, sink chan<- *AVSDirectoryOperatorRemovedFromOperatorSet, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _AVSDirectory.contract.WatchLogs(opts, "OperatorRemovedFromOperatorSet", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSDirectoryOperatorRemovedFromOperatorSet)
				if err := _AVSDirectory.contract.UnpackLog(event, "OperatorRemovedFromOperatorSet", log); err != nil {
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
func (_AVSDirectory *AVSDirectoryFilterer) ParseOperatorRemovedFromOperatorSet(log types.Log) (*AVSDirectoryOperatorRemovedFromOperatorSet, error) {
	event := new(AVSDirectoryOperatorRemovedFromOperatorSet)
	if err := _AVSDirectory.contract.UnpackLog(event, "OperatorRemovedFromOperatorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSDirectoryOperatorSetCreatedIterator is returned from FilterOperatorSetCreated and is used to iterate over the raw logs and unpacked data for OperatorSetCreated events raised by the AVSDirectory contract.
type AVSDirectoryOperatorSetCreatedIterator struct {
	Event *AVSDirectoryOperatorSetCreated // Event containing the contract specifics and raw log

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
func (it *AVSDirectoryOperatorSetCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSDirectoryOperatorSetCreated)
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
		it.Event = new(AVSDirectoryOperatorSetCreated)
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
func (it *AVSDirectoryOperatorSetCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSDirectoryOperatorSetCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSDirectoryOperatorSetCreated represents a OperatorSetCreated event raised by the AVSDirectory contract.
type AVSDirectoryOperatorSetCreated struct {
	OperatorSet IAVSDirectoryOperatorSet
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorSetCreated is a free log retrieval operation binding the contract event 0x31629285ead2335ae0933f86ed2ae63321f7af77b4e6eaabc42c057880977e6c.
//
// Solidity: event OperatorSetCreated((address,uint32) operatorSet)
func (_AVSDirectory *AVSDirectoryFilterer) FilterOperatorSetCreated(opts *bind.FilterOpts) (*AVSDirectoryOperatorSetCreatedIterator, error) {

	logs, sub, err := _AVSDirectory.contract.FilterLogs(opts, "OperatorSetCreated")
	if err != nil {
		return nil, err
	}
	return &AVSDirectoryOperatorSetCreatedIterator{contract: _AVSDirectory.contract, event: "OperatorSetCreated", logs: logs, sub: sub}, nil
}

// WatchOperatorSetCreated is a free log subscription operation binding the contract event 0x31629285ead2335ae0933f86ed2ae63321f7af77b4e6eaabc42c057880977e6c.
//
// Solidity: event OperatorSetCreated((address,uint32) operatorSet)
func (_AVSDirectory *AVSDirectoryFilterer) WatchOperatorSetCreated(opts *bind.WatchOpts, sink chan<- *AVSDirectoryOperatorSetCreated) (event.Subscription, error) {

	logs, sub, err := _AVSDirectory.contract.WatchLogs(opts, "OperatorSetCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSDirectoryOperatorSetCreated)
				if err := _AVSDirectory.contract.UnpackLog(event, "OperatorSetCreated", log); err != nil {
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
func (_AVSDirectory *AVSDirectoryFilterer) ParseOperatorSetCreated(log types.Log) (*AVSDirectoryOperatorSetCreated, error) {
	event := new(AVSDirectoryOperatorSetCreated)
	if err := _AVSDirectory.contract.UnpackLog(event, "OperatorSetCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSDirectoryOperatorSlashedIterator is returned from FilterOperatorSlashed and is used to iterate over the raw logs and unpacked data for OperatorSlashed events raised by the AVSDirectory contract.
type AVSDirectoryOperatorSlashedIterator struct {
	Event *AVSDirectoryOperatorSlashed // Event containing the contract specifics and raw log

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
func (it *AVSDirectoryOperatorSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSDirectoryOperatorSlashed)
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
		it.Event = new(AVSDirectoryOperatorSlashed)
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
func (it *AVSDirectoryOperatorSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSDirectoryOperatorSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSDirectoryOperatorSlashed represents a OperatorSlashed event raised by the AVSDirectory contract.
type AVSDirectoryOperatorSlashed struct {
	Operator      common.Address
	OperatorSetId uint32
	Strategy      common.Address
	BipsToSlash   uint16
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOperatorSlashed is a free log retrieval operation binding the contract event 0xe672839d3c371691acdb52de9fefc94b3dbf407dc0920ef566c7c059ad575b1c.
//
// Solidity: event OperatorSlashed(address operator, uint32 operatorSetId, address strategy, uint16 bipsToSlash)
func (_AVSDirectory *AVSDirectoryFilterer) FilterOperatorSlashed(opts *bind.FilterOpts) (*AVSDirectoryOperatorSlashedIterator, error) {

	logs, sub, err := _AVSDirectory.contract.FilterLogs(opts, "OperatorSlashed")
	if err != nil {
		return nil, err
	}
	return &AVSDirectoryOperatorSlashedIterator{contract: _AVSDirectory.contract, event: "OperatorSlashed", logs: logs, sub: sub}, nil
}

// WatchOperatorSlashed is a free log subscription operation binding the contract event 0xe672839d3c371691acdb52de9fefc94b3dbf407dc0920ef566c7c059ad575b1c.
//
// Solidity: event OperatorSlashed(address operator, uint32 operatorSetId, address strategy, uint16 bipsToSlash)
func (_AVSDirectory *AVSDirectoryFilterer) WatchOperatorSlashed(opts *bind.WatchOpts, sink chan<- *AVSDirectoryOperatorSlashed) (event.Subscription, error) {

	logs, sub, err := _AVSDirectory.contract.WatchLogs(opts, "OperatorSlashed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSDirectoryOperatorSlashed)
				if err := _AVSDirectory.contract.UnpackLog(event, "OperatorSlashed", log); err != nil {
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
func (_AVSDirectory *AVSDirectoryFilterer) ParseOperatorSlashed(log types.Log) (*AVSDirectoryOperatorSlashed, error) {
	event := new(AVSDirectoryOperatorSlashed)
	if err := _AVSDirectory.contract.UnpackLog(event, "OperatorSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSDirectoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AVSDirectory contract.
type AVSDirectoryOwnershipTransferredIterator struct {
	Event *AVSDirectoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AVSDirectoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSDirectoryOwnershipTransferred)
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
		it.Event = new(AVSDirectoryOwnershipTransferred)
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
func (it *AVSDirectoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSDirectoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSDirectoryOwnershipTransferred represents a OwnershipTransferred event raised by the AVSDirectory contract.
type AVSDirectoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AVSDirectory *AVSDirectoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AVSDirectoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AVSDirectory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AVSDirectoryOwnershipTransferredIterator{contract: _AVSDirectory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AVSDirectory *AVSDirectoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AVSDirectoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AVSDirectory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSDirectoryOwnershipTransferred)
				if err := _AVSDirectory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AVSDirectory *AVSDirectoryFilterer) ParseOwnershipTransferred(log types.Log) (*AVSDirectoryOwnershipTransferred, error) {
	event := new(AVSDirectoryOwnershipTransferred)
	if err := _AVSDirectory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSDirectoryPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the AVSDirectory contract.
type AVSDirectoryPausedIterator struct {
	Event *AVSDirectoryPaused // Event containing the contract specifics and raw log

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
func (it *AVSDirectoryPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSDirectoryPaused)
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
		it.Event = new(AVSDirectoryPaused)
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
func (it *AVSDirectoryPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSDirectoryPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSDirectoryPaused represents a Paused event raised by the AVSDirectory contract.
type AVSDirectoryPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_AVSDirectory *AVSDirectoryFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*AVSDirectoryPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AVSDirectory.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &AVSDirectoryPausedIterator{contract: _AVSDirectory.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_AVSDirectory *AVSDirectoryFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *AVSDirectoryPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AVSDirectory.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSDirectoryPaused)
				if err := _AVSDirectory.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_AVSDirectory *AVSDirectoryFilterer) ParsePaused(log types.Log) (*AVSDirectoryPaused, error) {
	event := new(AVSDirectoryPaused)
	if err := _AVSDirectory.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSDirectoryPauserRegistrySetIterator is returned from FilterPauserRegistrySet and is used to iterate over the raw logs and unpacked data for PauserRegistrySet events raised by the AVSDirectory contract.
type AVSDirectoryPauserRegistrySetIterator struct {
	Event *AVSDirectoryPauserRegistrySet // Event containing the contract specifics and raw log

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
func (it *AVSDirectoryPauserRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSDirectoryPauserRegistrySet)
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
		it.Event = new(AVSDirectoryPauserRegistrySet)
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
func (it *AVSDirectoryPauserRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSDirectoryPauserRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSDirectoryPauserRegistrySet represents a PauserRegistrySet event raised by the AVSDirectory contract.
type AVSDirectoryPauserRegistrySet struct {
	PauserRegistry    common.Address
	NewPauserRegistry common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPauserRegistrySet is a free log retrieval operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_AVSDirectory *AVSDirectoryFilterer) FilterPauserRegistrySet(opts *bind.FilterOpts) (*AVSDirectoryPauserRegistrySetIterator, error) {

	logs, sub, err := _AVSDirectory.contract.FilterLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return &AVSDirectoryPauserRegistrySetIterator{contract: _AVSDirectory.contract, event: "PauserRegistrySet", logs: logs, sub: sub}, nil
}

// WatchPauserRegistrySet is a free log subscription operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_AVSDirectory *AVSDirectoryFilterer) WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *AVSDirectoryPauserRegistrySet) (event.Subscription, error) {

	logs, sub, err := _AVSDirectory.contract.WatchLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSDirectoryPauserRegistrySet)
				if err := _AVSDirectory.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
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

// ParsePauserRegistrySet is a log parse operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_AVSDirectory *AVSDirectoryFilterer) ParsePauserRegistrySet(log types.Log) (*AVSDirectoryPauserRegistrySet, error) {
	event := new(AVSDirectoryPauserRegistrySet)
	if err := _AVSDirectory.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSDirectoryUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the AVSDirectory contract.
type AVSDirectoryUnpausedIterator struct {
	Event *AVSDirectoryUnpaused // Event containing the contract specifics and raw log

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
func (it *AVSDirectoryUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSDirectoryUnpaused)
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
		it.Event = new(AVSDirectoryUnpaused)
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
func (it *AVSDirectoryUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSDirectoryUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSDirectoryUnpaused represents a Unpaused event raised by the AVSDirectory contract.
type AVSDirectoryUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_AVSDirectory *AVSDirectoryFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*AVSDirectoryUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AVSDirectory.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &AVSDirectoryUnpausedIterator{contract: _AVSDirectory.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_AVSDirectory *AVSDirectoryFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *AVSDirectoryUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AVSDirectory.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSDirectoryUnpaused)
				if err := _AVSDirectory.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_AVSDirectory *AVSDirectoryFilterer) ParseUnpaused(log types.Log) (*AVSDirectoryUnpaused, error) {
	event := new(AVSDirectoryUnpaused)
	if err := _AVSDirectory.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
