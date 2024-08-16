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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_delegation\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ALLOCATION_DELAY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DOMAIN_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_PENDING_UPDATES\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"OPERATOR_AVS_REGISTRATION_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"OPERATOR_SET_FORCE_DEREGISTRATION_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"OPERATOR_SET_REGISTRATION_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allocate\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allocations\",\"type\":\"tuple[]\",\"internalType\":\"structIAVSDirectory.MagnitudeAdjustment[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"operatorSets\",\"type\":\"tuple[]\",\"internalType\":\"structIAVSDirectory.OperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"magnitudeDiffs\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}]},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"avsOperatorStatus\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIAVSDirectory.OperatorAVSRegistrationStatus\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"becomeOperatorSetAVS\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"calculateOperatorAVSRegistrationDigestHash\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateOperatorSetForceDeregistrationTypehash\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateOperatorSetRegistrationDigestHash\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cancelSalt\",\"inputs\":[{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createOperatorSets\",\"inputs\":[{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deallocate\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"deallocations\",\"type\":\"tuple[]\",\"internalType\":\"structIAVSDirectory.MagnitudeAdjustment[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"operatorSets\",\"type\":\"tuple[]\",\"internalType\":\"structIAVSDirectory.OperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"magnitudeDiffs\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}]},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregisterOperatorFromOperatorSets\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"domainSeparator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"forceDeregisterFromOperatorSets\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"freeDeallocatedMagnitude\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"freeMagnitude\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSlashableBips\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structIAVSDirectory.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"timestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"inTotalOperatorSets\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"initialPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isMember\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structIAVSDirectory.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperatorSet\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperatorSetAVS\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperatorSlashable\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structIAVSDirectory.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"migrateOperatorsToOperatorSets\",\"inputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"operatorSaltIsSpent\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorSetMemberCount\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorSetStatus\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"registered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"lastDeregisteredTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorSetsMemberOf\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"operatorSets\",\"type\":\"tuple[]\",\"internalType\":\"structIAVSDirectory.OperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorSetsMemberOf\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIAVSDirectory.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerOperatorToOperatorSets\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slashOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"bipsToSlash\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAVSMetadataURI\",\"inputs\":[{\"name\":\"metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AVSMetadataURIUpdated\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"metadataURI\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AVSMigratedToOperatorSets\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MagnitudeAllocated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIAVSDirectory.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"magnitudeToAllocate\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"effectTimestamp\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MagnitudeDeallocationCompleted\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIAVSDirectory.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"freeMagnitudeAdded\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MagnitudeQueueDeallocated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIAVSDirectory.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"magnitudeToDeallocate\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"completableTimestamp\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorAVSRegistrationStatusUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"status\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumIAVSDirectory.OperatorAVSRegistrationStatus\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorAddedToOperatorSet\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIAVSDirectory.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorMigratedToOperatorSets\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"indexed\":false,\"internalType\":\"uint32[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRemovedFromOperatorSet\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIAVSDirectory.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSetCreated\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIAVSDirectory.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSlashed\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"bipsToSlash\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x60c06040523480156200001157600080fd5b5060405162005b9e38038062005b9e833981016040819052620000349162000118565b6001600160a01b0381166080526200004b62000056565b504660a0526200014a565b600054610100900460ff1615620000c35760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff908116101562000116576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6000602082840312156200012b57600080fd5b81516001600160a01b03811681146200014357600080fd5b9392505050565b60805160a051615a276200017760003960006126760152600081816107540152610da50152615a276000f3fe608060405234801561001057600080fd5b50600436106102955760003560e01c80638da5cb5b11610167578063cb5b37be116100ce578063df5cf72311610087578063df5cf7231461074f578063ec76f44214610776578063ef2dfa8d146107aa578063f2fde38b146107bd578063f698da25146107d0578063fabc1cbc146107d857600080fd5b8063cb5b37be1461069d578063cbdf0e42146106b0578063d79aceab146106c3578063da2ff05d146106ea578063dae226b6146106fd578063dce974b91461072857600080fd5b8063afe02ed511610120578063afe02ed514610622578063b2841d4814610635578063b3e52eaa14610648578063bd74a06c14610650578063c1a8e2c514610663578063c825fe681461067657600080fd5b80638da5cb5b146105b05780638de54944146105c1578063955e6696146105e1578063a1060c88146105f4578063a98fb35514610607578063aec205c51461061a57600080fd5b80633fee332d1161020b5780636fbb031c116101c45780636fbb031c146104fa5780637019670814610519578063715018a61461052c5780637673e93a1461053457806384d76f7b14610557578063886f11951461058557600080fd5b80633fee332d1461042d57806349075da314610440578063595c6a671461047b5780635ac86ab7146104835780635c975abb146104a65780635fd6abfd146104ae57600080fd5b80631e2199e21161025d5780631e2199e2146103215780631e68134e1461033457806320606b701461039157806333429a6a146103c65780633367a33c146103ec578063374823b5146103ff57600080fd5b80630d5387c51461029a57806310d67a2f146102c35780631352c3e6146102d8578063136439dd146102fb5780631794bb3c1461030e575b600080fd5b6102ad6102a8366004614ba1565b6107eb565b6040516102ba9190614bd6565b60405180910390f35b6102d66102d1366004614c3c565b610930565b005b6102eb6102e6366004614ce0565b6109ec565b60405190151581526020016102ba565b6102d6610309366004614d65565b610a7c565b6102d661031c366004614d7e565b610bbb565b6102d661032f366004614eb5565b610ce5565b610375610342366004614f30565b609e60209081526000938452604080852082529284528284209052825290205460ff811690610100900463ffffffff1682565b60408051921515835263ffffffff9091166020830152016102ba565b6103b87f8cad95687ba82c2ce50e74f7b754645e5117c3a5bec8151c0726d5857980a86681565b6040519081526020016102ba565b6103d96103d4366004614f77565b611005565b60405161ffff90911681526020016102ba565b6102d66103fa366004614fdb565b61110d565b6102eb61040d36600461505b565b609960209081526000928352604080842090915290825290205460ff1681565b6102d661043b366004615087565b6111aa565b61046e61044e366004615115565b609860209081526000928352604080842090915290825290205460ff1681565b6040516102ba9190615164565b6102d66113ec565b6102eb61049136600461518c565b606654600160ff9092169190911b9081161490565b6066546103b8565b6104e26104bc366004615115565b60a06020908152600092835260408084209091529082529020546001600160401b031681565b6040516001600160401b0390911681526020016102ba565b610504621baf8081565b60405163ffffffff90911681526020016102ba565b6102d6610527366004614fdb565b6114b3565b6102d6611523565b6102eb610542366004614c3c565b609a6020526000908152604090205460ff1681565b6102eb6105653660046151af565b609b60209081526000928352604080842090915290825290205460ff1681565b606554610598906001600160a01b031681565b6040516001600160a01b0390911681526020016102ba565b6033546001600160a01b0316610598565b6105d46105cf36600461505b565b611537565b6040516102ba91906151e4565b6103b86105ef36600461520a565b611578565b6103b8610602366004615270565b6115dd565b6102d66106153660046152b6565b611647565b6102d661168e565b6102d6610630366004615327565b611756565b6103b861064336600461520a565b611923565b6103b8600181565b6102d661065e366004615368565b611963565b6102d66106713660046153e9565b611e29565b6103b87f809c5ac049c45b7a7f050a20f00c16cf63797efbf8b1eb8d749fdfa39ff8f92981565b6102d66106ab3660046153e9565b611e5e565b6103b86106be366004614c3c565b611ea4565b6103b87fda2c89bafdd34776a2b8bb9c83c82f419e20cc8c67207f70edd58249b92661bd81565b6102eb6106f8366004614ce0565b611ec5565b6103b861070b3660046151af565b609c60209081526000928352604080842090915290825290205481565b6103b87f4ee65f64218c67b68da66fd0db16560040a6b973290b9e71912d661ee53fe49581565b6105987f000000000000000000000000000000000000000000000000000000000000000081565b6102d6610784366004614d65565b33600090815260996020908152604080832093835292905220805460ff19166001179055565b6102d66107b836600461543d565b611ef1565b6102d66107cb366004614c3c565b61229a565b6103b8612310565b6102d66107e6366004614d65565b61231f565b6001600160a01b0383166000908152609d602052604081206060919084906108129061247b565b61081c91906154be565b90508083111561082a578092505b826001600160401b0381111561084257610842614c59565b60405190808252806020026020018201604052801561088757816020015b60408051808201909152600080825260208201528152602001906001900390816108605790505b50915060005b83811015610927576108f96108c36108a583886154d5565b6001600160a01b0389166000908152609d6020526040902090612485565b60408051808201909152600080825260208201525060408051808201909152606082901c815263ffffffff909116602082015290565b83828151811061090b5761090b6154ed565b60200260200101819052508061092090615503565b905061088d565b50509392505050565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610983573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109a7919061551e565b6001600160a01b0316336001600160a01b0316146109e05760405162461bcd60e51b81526004016109d79061553b565b60405180910390fd5b6109e981612491565b50565b80516001600160a01b039081166000908152609e6020908152604080832093861683529281528282208185015163ffffffff908116845290825283832084518086019095525460ff8116151585526101009004169083015290610a4f8484611ec5565b80610a72575042621baf808260200151610a699190615585565b63ffffffff1610155b9150505b92915050565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015610ac4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ae891906155ad565b610b045760405162461bcd60e51b81526004016109d7906155cf565b60665481811614610b7d5760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c697479000000000000000060648201526084016109d7565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b600054610100900460ff1615808015610bdb5750600054600160ff909116105b80610bf55750303b158015610bf5575060005460ff166001145b610c585760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016109d7565b6000805460ff191660011790558015610c7b576000805461ff0019166101001790555b610c858383612588565b610c8d612672565b609755610c998461273b565b8015610cdf576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050565b60665460019060029081161415610d0e5760405162461bcd60e51b81526004016109d790615617565b4282604001511015610d865760405162461bcd60e51b815260206004820152604760248201526000805160206159d283398151915260448201527f70657261746f72536574733a206f70657261746f72207369676e617475726520606482015266195e1c1a5c995960ca1b608482015260a4016109d7565b6040516336b87bd760e11b81526001600160a01b0386811660048301527f00000000000000000000000000000000000000000000000000000000000000001690636d70f7ae90602401602060405180830381865afa158015610dec573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e1091906155ad565b610e8f5760405162461bcd60e51b815260206004820152605660248201526000805160206159d283398151915260448201527f70657261746f72536574733a206f70657261746f72206e6f7420726567697374606482015275195c9959081d1bc8115a59d95b93185e595c881e595d60521b608482015260a4016109d7565b336000908152609a602052604090205460ff16610f165760405162461bcd60e51b815260206004820152604b60248201526000805160206159d283398151915260448201527f70657261746f72536574733a20415653206973206e6f7420616e206f7065726160648201526a746f72207365742041565360a81b608482015260a4016109d7565b6001600160a01b038516600090815260996020908152604080832085830151845290915290205460ff1615610fa15760405162461bcd60e51b815260206004820152603f60248201526000805160206159d283398151915260448201527f70657261746f72536574733a2073616c7420616c7265616479207370656e740060648201526084016109d7565b610fc185610fba33878787602001518860400151611578565b845161278d565b6001600160a01b03851660009081526099602090815260408083208583015184529091529020805460ff19166001179055610ffe85338686612947565b5050505050565b6001600160a01b038085166000908152609f60209081526040808320938616835292905290812081906110389084612d16565b6001600160a01b03808816600090815260a1602090815260408083209389168352928152918120929350916110d5918691908490611078908b018b614c3c565b6001600160a01b03166001600160a01b0316815260200190815260200160002060008960200160208101906110ad919061564e565b63ffffffff1663ffffffff168152602001908152602001600020612d1690919063ffffffff16565b9050816001600160401b0316612710826001600160401b03166110f89190615669565b611102919061569e565b979650505050505050565b600061111c621baf8042615585565b905060005b838110156111a2576111648686868481811061113f5761113f6154ed565b905060200281019061115191906156c0565b61115f906020810190614c3c565b612d6e565b6111928686868481811061117a5761117a6154ed565b905060200281019061118c91906156c0565b84612e40565b61119b81615503565b9050611121565b505050505050565b606654600190600290811614156111d35760405162461bcd60e51b81526004016109d790615617565b81515161126b57336001600160a01b038716146112665760405162461bcd60e51b815260206004820152604560248201527f4156534469726563746f72792e666f7263654465726567697374657246726f6d60448201527f4f70657261746f72536574733a2063616c6c6572206d757374206265206f70656064820152643930ba37b960d91b608482015260a4016109d7565b6113e0565b42826040015110156112f65760405162461bcd60e51b815260206004820152604860248201527f4156534469726563746f72792e666f7263654465726567697374657246726f6d60448201527f4f70657261746f72536574733a206f70657261746f72207369676e617475726560648201526708195e1c1a5c995960c21b608482015260a4016109d7565b6001600160a01b038616600090815260996020908152604080832085830151845290915290205460ff1615611395576040805162461bcd60e51b81526020600482015260248101919091527f4156534469726563746f72792e666f7263654465726567697374657246726f6d60448201527f4f70657261746f72536574733a2073616c7420616c7265616479207370656e7460648201526084016109d7565b6113ae86610fba87878787602001518860400151611923565b6001600160a01b03861660009081526099602090815260408083208583015184529091529020805460ff191660011790555b6111a2858786866133eb565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015611434573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061145891906155ad565b6114745760405162461bcd60e51b81526004016109d7906155cf565b600019606681905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b60006114c2621baf8042615585565b905060005b838110156111a2576114e58686868481811061113f5761113f6154ed565b611513868686848181106114fb576114fb6154ed565b905060200281019061150d91906156c0565b846135cd565b61151c81615503565b90506114c7565b61152b613bc5565b611535600061273b565b565b60408051808201909152600080825260208201526001600160a01b0383166000908152609d60205260409020611571906108c39084612485565b9392505050565b60006115d37f809c5ac049c45b7a7f050a20f00c16cf63797efbf8b1eb8d749fdfa39ff8f92987878787876040516020016115b896959493929190615724565b60405160208183030381529060405280519060200120613c1f565b9695505050505050565b604080517fda2c89bafdd34776a2b8bb9c83c82f419e20cc8c67207f70edd58249b92661bd60208201526001600160a01b038087169282019290925290841660608201526080810183905260a0810182905260009061163e9060c0016115b8565b95945050505050565b336001600160a01b03167fa89c1dc243d8908a96dd84944bcc97d6bc6ac00dd78e20621576be6a3c9437138383604051611682929190615764565b60405180910390a25050565b336000908152609a602052604090205460ff16156117145760405162461bcd60e51b815260206004820152603e60248201527f4156534469726563746f72792e6265636f6d654f70657261746f72536574415660448201527f533a20616c726561647920616e206f70657261746f722073657420415653000060648201526084016109d7565b336000818152609a6020526040808220805460ff19166001179055517f702b0c1f6cb1cf511aaa81f72bc05a215bb3497632d72c690c822b044ab494bf9190a2565b60005b8181101561191e57336000908152609b6020526040812090848484818110611783576117836154ed565b9050602002016020810190611798919061564e565b63ffffffff16815260208101919091526040016000205460ff16156118255760405162461bcd60e51b815260206004820152603b60248201527f4156534469726563746f72792e6372656174654f70657261746f725365743a2060448201527f6f70657261746f722073657420616c726561647920657869737473000000000060648201526084016109d7565b336000908152609b60205260408120600191858585818110611849576118496154ed565b905060200201602081019061185e919061564e565b63ffffffff1663ffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055507f31629285ead2335ae0933f86ed2ae63321f7af77b4e6eaabc42c057880977e6c6040518060400160405280336001600160a01b031681526020018585858181106118dc576118dc6154ed565b90506020020160208101906118f1919061564e565b63ffffffff16905260405161190691906151e4565b60405180910390a161191781615503565b9050611759565b505050565b60006115d37f4ee65f64218c67b68da66fd0db16560040a6b973290b9e71912d661ee53fe49587878787876040516020016115b896959493929190615724565b611990856040518060400160405280336001600160a01b031681526020018763ffffffff168152506109ec565b611a0d5760405162461bcd60e51b815260206004820152604260248201527f4156534469726563746f72792e736c6173684f70657261746f723a206f70657260448201527f61746f72206e6f7420736c61736861626c6520666f72206f70657261746f7253606482015261195d60f21b608482015260a4016109d7565b60005b828110156111a2576001600160a01b038616600090815260a1602052604081208190611a9b90429083898988818110611a4b57611a4b6154ed565b9050602002016020810190611a609190614c3c565b6001600160a01b0316815260208082019290925260409081016000908120338252835281812063ffffffff808e1683529352209190613c6616565b9050612710611ab76001600160401b03831661ffff8716615669565b611ac1919061569e565b6001600160a01b038916600090815260a160205260408120919350611b519142916001600160401b03861691908a8a89818110611b0057611b006154ed565b9050602002016020810190611b159190614c3c565b6001600160a01b0316815260208082019290925260409081016000908120338252835281812063ffffffff808f168352935220929190613d0b16565b506001600160a01b038716600090815260a26020526040812081878786818110611b7d57611b7d6154ed565b9050602002016020810190611b929190614c3c565b6001600160a01b0316815260208082019290925260409081016000908120338252835281812063ffffffff8b1682529092529020549050805b8015611d2f576001600160a01b038916600090815260a26020526040812081898988818110611bfc57611bfc6154ed565b9050602002016020810190611c119190614c3c565b6001600160a01b0316815260208082019290925260409081016000908120338252835281812063ffffffff8d1682529092529020611c506001846154be565b81548110611c6057611c606154ed565b60009182526020909120018054909150600160401b900463ffffffff16611c8a621baf8042615585565b63ffffffff161115611d1857805460009061271090611cb6906001600160401b031661ffff8a16615669565b611cc0919061569e565b825490915081908390600090611ce09084906001600160401b0316615793565b92506101000a8154816001600160401b0302191690836001600160401b031602179055508085611d1091906157bb565b945050611d1e565b50611d2f565b50611d28816157dd565b9050611bcb565b50506001600160a01b0387166000908152609f60205260408120611e159142916001600160401b03851691611dab91908a8a89818110611d7157611d716154ed565b9050602002016020810190611d869190614c3c565b6001600160a01b03166001600160a01b03168152602001908152602001600020613da7565b611db591906157f4565b6001600160a01b038a166000908152609f6020526040812090898988818110611de057611de06154ed565b9050602002016020810190611df59190614c3c565b6001600160a01b0316815260208101919091526040016000209190613de0565b50505080611e2290615503565b9050611a10565b60665460019060029081161415611e525760405162461bcd60e51b81526004016109d790615617565b610cdf338585856133eb565b60005b81811015610cdf57611e9484848484818110611e7f57611e7f6154ed565b905060200201602081019061115f9190614c3c565b611e9d81615503565b9050611e61565b6001600160a01b0381166000908152609d60205260408120610a769061247b565b6000611571611ed383613dfb565b6001600160a01b0385166000908152609d6020526040902090613e60565b60665460019060029081161415611f1a5760405162461bcd60e51b81526004016109d790615617565b336000908152609a602052604090205460ff16611fb35760405162461bcd60e51b815260206004820152604b60248201527f4156534469726563746f72792e6d6967726174654f70657261746f7273546f4f60448201527f70657261746f72536574733a20415653206973206e6f7420616e206f7065726160648201526a746f72207365742041565360a81b608482015260a4016109d7565b60005b848110156111a257600133600090815260986020526040812090888885818110611fe257611fe26154ed565b9050602002016020810190611ff79190614c3c565b6001600160a01b0316815260208101919091526040016000205460ff1660018111156120255761202561514e565b146120d15760405162461bcd60e51b815260206004820152606a60248201527f4156534469726563746f72792e6d6967726174654f70657261746f7273546f4f60448201527f70657261746f72536574733a206f70657261746f7220616c7265616479206d6960648201527f677261746564206f72206e6f742061206c656761637920726567697374657265608482015269321037b832b930ba37b960b11b60a482015260c4016109d7565b6121258686838181106120e6576120e66154ed565b90506020020160208101906120fb9190614c3c565b3386868581811061210e5761210e6154ed565b90506020028101906121209190615814565b612947565b33600090815260986020526040812081888885818110612147576121476154ed565b905060200201602081019061215c9190614c3c565b6001600160a01b031681526020810191909152604001600020805460ff19166001838181111561218e5761218e61514e565b0217905550338686838181106121a6576121a66154ed565b90506020020160208101906121bb9190614c3c565b6001600160a01b03167ff0952b1c65271d819d39983d2abb044b9cace59bcc4d4dd389f586ebdcb15b4160006040516121f49190615164565b60405180910390a33386868381811061220f5761220f6154ed565b90506020020160208101906122249190614c3c565b6001600160a01b03167f54f33cfdd1ca703d795986b986fd47d742eab1904ecd2a5fdb8d6595e5904a01868685818110612260576122606154ed565b90506020028101906122729190615814565b60405161228092919061585d565b60405180910390a38061229281615503565b915050611fb6565b6122a2613bc5565b6001600160a01b0381166123075760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016109d7565b6109e98161273b565b600061231a612672565b905090565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612372573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612396919061551e565b6001600160a01b0316336001600160a01b0316146123c65760405162461bcd60e51b81526004016109d79061553b565b6066541981196066541916146124445760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c697479000000000000000060648201526084016109d7565b606681905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c90602001610bb0565b6000610a76825490565b60006115718383613e78565b6001600160a01b03811661251f5760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a4016109d7565b606554604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1606580546001600160a01b0319166001600160a01b0392909216919091179055565b6065546001600160a01b03161580156125a957506001600160a01b03821615155b61262b5760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a4016109d7565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a261266e82612491565b5050565b60007f00000000000000000000000000000000000000000000000000000000000000004614156126a3575060975490565b50604080518082018252600a81526922b4b3b2b72630bcb2b960b11b60209182015281517f8cad95687ba82c2ce50e74f7b754645e5117c3a5bec8151c0726d5857980a866818301527f71b625cfad44bac63b13dba07f2e1d6084ee04b6f8752101ece6126d584ee6ea81840152466060820152306080808301919091528351808303909101815260a0909101909252815191012090565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6001600160a01b0383163b156128a757604051630b135d3f60e11b808252906001600160a01b03851690631626ba7e906127cd9086908690600401615879565b602060405180830381865afa1580156127ea573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061280e91906158d6565b6001600160e01b0319161461191e5760405162461bcd60e51b815260206004820152605360248201527f454950313237315369676e61747572655574696c732e636865636b5369676e6160448201527f747572655f454950313237313a2045524331323731207369676e6174757265206064820152721d995c9a599a58d85d1a5bdb8819985a5b1959606a1b608482015260a4016109d7565b826001600160a01b03166128bb8383613ea2565b6001600160a01b03161461191e5760405162461bcd60e51b815260206004820152604760248201527f454950313237315369676e61747572655574696c732e636865636b5369676e6160448201527f747572655f454950313237313a207369676e6174757265206e6f742066726f6d6064820152661039b4b3b732b960c91b608482015260a4016109d7565b60005b81811015610ffe5760006040518060400160405280866001600160a01b03168152602001858585818110612980576129806154ed565b9050602002016020810190612995919061564e565b63ffffffff1690526001600160a01b0386166000908152609b602052604081209192508585858181106129ca576129ca6154ed565b90506020020160208101906129df919061564e565b63ffffffff16815260208101919091526040016000205460ff16612a765760405162461bcd60e51b815260206004820152604260248201527f4156534469726563746f72792e5f72656769737465724f70657261746f72546f60448201527f4f70657261746f72536574733a20696e76616c6964206f70657261746f722073606482015261195d60f21b608482015260a4016109d7565b612a808682611ec5565b15612b195760405162461bcd60e51b815260206004820152605960248201527f4156534469726563746f72792e5f72656769737465724f70657261746f72546f60448201527f4f70657261746f72536574733a206f70657261746f7220616c7265616479207260648201527f65676973746572656420746f206f70657261746f722073657400000000000000608482015260a4016109d7565b6001600160a01b0385166000908152609c6020526040812090858585818110612b4457612b446154ed565b9050602002016020810190612b59919061564e565b63ffffffff1663ffffffff16815260200190815260200160002060008154612b8090615503565b90915550612bae612b9082613dfb565b6001600160a01b0388166000908152609d6020526040902090613ec6565b506001600160a01b038086166000908152609e60209081526040808320938a16835292905290812081868686818110612be957612be96154ed565b9050602002016020810190612bfe919061564e565b63ffffffff1681526020810191909152604001600020805490915060ff1615612cb55760405162461bcd60e51b815260206004820152605a60248201527f4156534469726563746f72792e5f72656769737465724f70657261746f72546f60448201527f4f70657261746f72536574733a206f70657261746f7220616c7265616479207260648201527f65676973746572656420666f72206f70657261746f7220736574000000000000608482015260a4016109d7565b805460ff191660011781556040516001600160a01b038816907f43232edf9071753d2321e5fa7e018363ee248e5f2142e6c08edd3265bfb4895e90612cfb9085906151e4565b60405180910390a2505080612d0f90615503565b905061294a565b815460009081612d2885858385613ed2565b90508015612d6357612d4d85612d3f6001846154be565b600091825260209091200190565b54600160201b90046001600160e01b031661163e565b600095945050505050565b6001600160a01b0382166000908152609d602052604081208190612d919061247b565b905060005b81811015612dd5576000612daa8683611537565b9050612db7868683613f28565b612dc190856157bb565b93505080612dce90615503565b9050612d96565b506001600160a01b03808516600090815260a06020908152604080832093871683529290529081208054849290612e169084906001600160401b03166157bb565b92506101000a8154816001600160401b0302191690836001600160401b0316021790555050505050565b6000612e4f6020840184614c3c565b9050612e5e6040840184615814565b9050612e6d6020850185615900565b905014612ef45760405162461bcd60e51b815260206004820152604960248201527f4156534469726563746f72792e5f6465616c6c6f636174653a206f706572617460448201527f6f725365747320616e64206d61676e69747564654469666673206c656e677468606482015268040dad2e6dac2e8c6d60bb1b608482015260a4016109d7565b366000612f046020860186615900565b9150915060005b818110156133e2576001600160a01b03808816600090815260a1602090815260408083209388168352929052908120612fdd90429083878787818110612f5357612f536154ed565b612f699260206040909202019081019150614c3c565b6001600160a01b03166001600160a01b031681526020019081526020016000206000878787818110612f9d57612f9d6154ed565b9050604002016020016020810190612fb5919061564e565b63ffffffff1663ffffffff168152602001908152602001600020613c6690919063ffffffff16565b90506000612fee6040890189615814565b84818110612ffe57612ffe6154ed565b90506020020160208101906130139190615949565b6001600160401b03161161308f5760405162461bcd60e51b815260206004820152603e60248201527f4156534469726563746f72792e5f6465616c6c6f636174653a206d61676e697460448201527f75646544696666206d7573742062652067726561746572207468616e2030000060648201526084016109d7565b6001600160401b0381166130a66040890189615814565b848181106130b6576130b66154ed565b90506020020160208101906130cb9190615949565b6001600160401b031611156131585760405162461bcd60e51b815260206004820152604760248201527f4156534469726563746f72792e5f6465616c6c6f636174653a2063616e6e6f7460448201527f206465616c6c6f63617465206d6f7265207468616e207768617420697320616c6064820152661b1bd8d85d195960ca1b608482015260a4016109d7565b6132604261316960408a018a615814565b85818110613179576131796154ed565b905060200201602081019061318e9190615949565b6001600160a01b03808c16600090815260a160209081526040808320938c1683529290529081206001600160401b039290921691908888888181106131d5576131d56154ed565b6131eb9260206040909202019081019150614c3c565b6001600160a01b03166001600160a01b03168152602001908152602001600020600088888881811061321f5761321f6154ed565b9050604002016020016020810190613237919061564e565b63ffffffff1663ffffffff168152602001908152602001600020613d0b9092919063ffffffff16565b6132828886868686818110613277576132776154ed565b9050604002016140e4565b506001600160a01b03808916600090815260a2602090815260408083209389168352929052908120908585858181106132bd576132bd6154ed565b6132d39260206040909202019081019150614c3c565b6001600160a01b03166001600160a01b031681526020019081526020016000206000858585818110613307576133076154ed565b905060400201602001602081019061331f919061564e565b63ffffffff1663ffffffff16815260200190815260200160002060405180604001604052808980604001906133549190615814565b86818110613364576133646154ed565b90506020020160208101906133799190615949565b6001600160401b03908116825263ffffffff808b166020938401528454600181018655600095865294839020845195018054949093015116600160401b026bffffffffffffffffffffffff19909316931692909217179055506133db81615503565b9050612f0b565b50505050505050565b60005b81811015610ffe5760006040518060400160405280876001600160a01b03168152602001858585818110613424576134246154ed565b9050602002016020810190613439919061564e565b63ffffffff169052905061344d8582611ec5565b6134e55760405162461bcd60e51b815260206004820152605960248201527f4156534469726563746f72792e5f646572656769737465724f70657261746f7260448201527f46726f6d4f70657261746f725365743a206f70657261746f72206e6f7420726560648201527f676973746572656420666f72206f70657261746f722073657400000000000000608482015260a4016109d7565b6001600160a01b0386166000908152609c6020526040812090858585818110613510576135106154ed565b9050602002016020810190613525919061564e565b63ffffffff1663ffffffff1681526020019081526020016000206000815461354c906157dd565b9091555061357a61355c82613dfb565b6001600160a01b0387166000908152609d60205260409020906142e7565b50846001600160a01b03167fad34c3070be1dffbcaa499d000ba2b8d9848aefcac3059df245dd95c4ece14fe826040516135b491906151e4565b60405180910390a2506135c681615503565b90506133ee565b60006135dc6020840184614c3c565b90503660006135ee6020860186615900565b6001600160a01b03808916600090815260a0602090815260408083209389168352929052205491935091506001600160401b031661362c8785612d6e565b60005b82811015613b755760006136466040890189615814565b83818110613656576136566154ed565b905060200201602081019061366b9190615949565b6001600160401b0316116136f35760405162461bcd60e51b815260206004820152604360248201527f4156534469726563746f72792e7175657565416c6c6f636174696f6e733a206d60448201527f61676e697475646544696666206d75737420626520677265617465722074686160648201526206e20360ec1b608482015260a4016109d7565b6137006040880188615814565b82818110613710576137106154ed565b90506020020160208101906137259190615949565b6001600160401b0316826001600160401b031610156137c55760405162461bcd60e51b815260206004820152605060248201527f4156534469726563746f72792e7175657565416c6c6f636174696f6e733a206960448201527f6e73756666696369656e7420617661696c61626c652066726565206d61676e6960648201526f7475646520746f20616c6c6f6361746560801b608482015260a4016109d7565b6001600160a01b03808916600090815260a1602090815260408083209389168352929052908120819061389190429083898988818110613807576138076154ed565b61381d9260206040909202019081019150614c3c565b6001600160a01b03166001600160a01b031681526020019081526020016000206000898988818110613851576138516154ed565b9050604002016020016020810190613869919061564e565b63ffffffff1663ffffffff1681526020019081526020016000206142f390919063ffffffff16565b91509150816001600160e01b031660001415806138ad57508015155b15613a13576001600160a01b03808b16600090815260a160209081526040808320938b16835292905290812061396f918888878181106138ef576138ef6154ed565b6139059260206040909202019081019150614c3c565b6001600160a01b03166001600160a01b031681526020019081526020016000206000888887818110613939576139396154ed565b9050604002016020016020810190613951919061564e565b63ffffffff1663ffffffff1681526020019081526020016000205490565b61397a6001836154d5565b14613a135760405162461bcd60e51b815260206004820152605d60248201527f4156534469726563746f72792e7175657565416c6c6f636174696f6e733a206560448201527f7863656564206d61782070656e64696e6720616c6c6f636174696f6e7320616c60648201527f6c6f77656420666f72206f702c206f705365742c207374726174656779000000608482015260a4016109d7565b613b2188613a2460408c018c615814565b86818110613a3457613a346154ed565b9050602002016020810190613a499190615949565b613a5c906001600160401b031685615972565b6001600160a01b03808e16600090815260a160209081526040808320938e168352929052908120908a8a89818110613a9657613a966154ed565b613aac9260206040909202019081019150614c3c565b6001600160a01b03166001600160a01b0316815260200190815260200160002060008a8a89818110613ae057613ae06154ed565b9050604002016020016020810190613af8919061564e565b63ffffffff1663ffffffff168152602001908152602001600020613de09092919063ffffffff16565b50613b31905060408a018a615814565b84818110613b4157613b416154ed565b9050602002016020810190613b569190615949565b613b609085615793565b9350505080613b6e90615503565b905061362f565b506001600160a01b03968716600090815260a0602090815260408083209690991682529490945295909220805467ffffffffffffffff19166001600160401b039096169590951790945550505050565b6033546001600160a01b031633146115355760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016109d7565b6000613c29612672565b60405161190160f01b6020820152602281019190915260428101839052606201604051602081830303815290604052805190602001209050919050565b815460009081816005811115613cc3576000613c81846143ac565b613c8b90856154be565b60008881526020902090915081015463ffffffff9081169087161015613cb357809150613cc1565b613cbe8160016154d5565b92505b505b6000613cd187878585613ed2565b90508015613cfe57613ce887612d3f6001846154be565b54600160201b90046001600160e01b0316611102565b6000979650505050505050565b600080613d1885856142f3565b91509150816001600160e01b03166000148015613d33575080155b15613d3d57506000195b84545b808210156111a2576000868152602090208201805485908290600490613d77908490600160201b90046001600160e01b03166157f4565b92506101000a8154816001600160e01b0302191690836001600160e01b0316021790555082600101925050613d40565b80546000908015613dd757613dc183612d3f6001846154be565b54600160201b90046001600160e01b0316611571565b60009392505050565b600080613dee858585614491565b915091505b935093915050565b60008160000151826020015163ffffffff16604051602001613e4892919060609290921b6bffffffffffffffffffffffff1916825260a01b6001600160a01b031916601482015260200190565b604051602081830303815290604052610a7690615994565b60008181526001830160205260408120541515611571565b6000826000018281548110613e8f57613e8f6154ed565b9060005260206000200154905092915050565b6000806000613eb18585614637565b91509150613ebe816146a4565b509392505050565b6000611571838361485f565b60005b81831015613ebe576000613ee984846148ae565b60008781526020902090915063ffffffff86169082015463ffffffff161115613f1457809250613f22565b613f1f8160016154d5565b93505b50613ed5565b6001600160a01b03808416600090815260a26020908152604080832086851684528252808320855190941683529281528282208482015163ffffffff16835281528282208054845181840281018401909552808552929384939092849084015b82821015613fd457600084815260209081902060408051808201909152908401546001600160401b0381168252600160401b900463ffffffff1681830152825260019092019101613f88565b5050506001600160a01b03808816600090815260a3602090815260408083208a851684528252808320895190941683529281528282208882015163ffffffff16835290522054919250505b8151811015614099574263ffffffff16828281518110614041576140416154ed565b60200260200101516020015163ffffffff16111561405e57614099565b818181518110614070576140706154ed565b6020026020010151600001518361408791906157bb565b925061409281615503565b905061401f565b6001600160a01b03808716600090815260a36020908152604080832089851684528252808320885190941683529281528282208782015163ffffffff16835290522055509392505050565b6001600160a01b03808416600090815260a26020908152604080832093861683529281529181209091829190829061411e90860186614c3c565b6001600160a01b03166001600160a01b031681526020019081526020016000206000846020016020810190614153919061564e565b63ffffffff1663ffffffff168152602001908152602001600020805480602002602001604051908101604052809291908181526020016000905b828210156141d957600084815260209081902060408051808201909152908401546001600160401b0381168252600160401b900463ffffffff168183015282526001909201910161418d565b50508251929350829150505b80156142dd5763ffffffff4216836141fe6001846154be565b8151811061420e5761420e6154ed565b60200260200101516020015163ffffffff1610156142c857600161423282846154be565b61423d9060016154d5565b106142c35760405162461bcd60e51b815260206004820152604a60248201527f4156534469726563746f72792e5f636865636b50656e64696e674465616c6c6f60448201527f636174696f6e733a2065786365656473206d61782070656e64696e67206465616064820152696c6c6f636174696f6e7360b01b608482015260a4016109d7565b6142cd565b6142dd565b6142d6816157dd565b90506141e5565b5050509392505050565b600061157183836148c9565b8154600090819081816005811115614352576000614310846143ac565b61431a90856154be565b60008981526020902090915081015463ffffffff908116908816101561434257809150614350565b61434d8160016154d5565b92505b505b600061436088888585613ed2565b905080156143985761437788612d3f6001846154be565b54600160201b90046001600160e01b03166143936001836154be565b61439c565b6000805b95509550505050505b9250929050565b6000816143bb57506000919050565b600060016143c8846149bc565b901c6001901b905060018184816143e1576143e1615688565b048201901c905060018184816143f9576143f9615688565b048201901c9050600181848161441157614411615688565b048201901c9050600181848161442957614429615688565b048201901c9050600181848161444157614441615688565b048201901c9050600181848161445957614459615688565b048201901c9050600181848161447157614471615688565b048201901c90506115718182858161448b5761448b615688565b04614a50565b8254600090819080156145de5760006144af87612d3f6001856154be565b60408051808201909152905463ffffffff808216808452600160201b9092046001600160e01b0316602084015291925090871610156145305760405162461bcd60e51b815260206004820152601b60248201527f436865636b706f696e743a2064656372656173696e67206b657973000000000060448201526064016109d7565b8563ffffffff16816000015163ffffffff16141561457f578461455888612d3f6001866154be565b80546001600160e01b0392909216600160201b0263ffffffff9092169190911790556145ce565b6040805180820190915263ffffffff80881682526001600160e01b0380881660208085019182528b54600181018d5560008d81529190912094519151909216600160201b029216919091179101555b602001519250839150613df39050565b50506040805180820190915263ffffffff80851682526001600160e01b0380851660208085019182528854600181018a5560008a815291822095519251909316600160201b029190931617920191909155905081613df3565b60008082516041141561466e5760208301516040840151606085015160001a61466287828585614a66565b945094505050506143a5565b825160401415614698576020830151604084015161468d868383614b53565b9350935050506143a5565b506000905060026143a5565b60008160048111156146b8576146b861514e565b14156146c15750565b60018160048111156146d5576146d561514e565b14156147235760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e6174757265000000000000000060448201526064016109d7565b60028160048111156147375761473761514e565b14156147855760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e6774680060448201526064016109d7565b60038160048111156147995761479961514e565b14156147f25760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b60648201526084016109d7565b60048160048111156148065761480661514e565b14156109e95760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b60648201526084016109d7565b60008181526001830160205260408120546148a657508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610a76565b506000610a76565b60006148bd600284841861569e565b611571908484166154d5565b600081815260018301602052604081205480156149b25760006148ed6001836154be565b8554909150600090614901906001906154be565b9050818114614966576000866000018281548110614921576149216154ed565b9060005260206000200154905080876000018481548110614944576149446154ed565b6000918252602080832090910192909255918252600188019052604090208390555b8554869080614977576149776159bb565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050610a76565b6000915050610a76565b600080608083901c156149d157608092831c92015b604083901c156149e357604092831c92015b602083901c156149f557602092831c92015b601083901c15614a0757601092831c92015b600883901c15614a1957600892831c92015b600483901c15614a2b57600492831c92015b600283901c15614a3d57600292831c92015b600183901c15610a765760010192915050565b6000818310614a5f5781611571565b5090919050565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115614a9d5750600090506003614b4a565b8460ff16601b14158015614ab557508460ff16601c14155b15614ac65750600090506004614b4a565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015614b1a573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116614b4357600060019250925050614b4a565b9150600090505b94509492505050565b6000806001600160ff1b03831681614b7060ff86901c601b6154d5565b9050614b7e87828885614a66565b935093505050935093915050565b6001600160a01b03811681146109e957600080fd5b600080600060608486031215614bb657600080fd5b8335614bc181614b8c565b95602085013595506040909401359392505050565b602080825282518282018190526000919060409081850190868401855b82811015614c2f57614c1f84835180516001600160a01b0316825260209081015163ffffffff16910152565b9284019290850190600101614bf3565b5091979650505050505050565b600060208284031215614c4e57600080fd5b813561157181614b8c565b634e487b7160e01b600052604160045260246000fd5b604051606081016001600160401b0381118282101715614c9157614c91614c59565b60405290565b604051601f8201601f191681016001600160401b0381118282101715614cbf57614cbf614c59565b604052919050565b803563ffffffff81168114614cdb57600080fd5b919050565b6000808284036060811215614cf457600080fd5b8335614cff81614b8c565b92506040601f1982011215614d1357600080fd5b50604051604081018181106001600160401b0382111715614d3657614d36614c59565b6040526020840135614d4781614b8c565b8152614d5560408501614cc7565b6020820152809150509250929050565b600060208284031215614d7757600080fd5b5035919050565b600080600060608486031215614d9357600080fd5b8335614d9e81614b8c565b92506020840135614dae81614b8c565b929592945050506040919091013590565b60008083601f840112614dd157600080fd5b5081356001600160401b03811115614de857600080fd5b6020830191508360208260051b85010111156143a557600080fd5b600060608284031215614e1557600080fd5b614e1d614c6f565b905081356001600160401b0380821115614e3657600080fd5b818401915084601f830112614e4a57600080fd5b8135602082821115614e5e57614e5e614c59565b614e70601f8301601f19168201614c97565b92508183528681838601011115614e8657600080fd5b818185018285013760008183850101528285528086013581860152505050506040820135604082015292915050565b60008060008060608587031215614ecb57600080fd5b8435614ed681614b8c565b935060208501356001600160401b0380821115614ef257600080fd5b614efe88838901614dbf565b90955093506040870135915080821115614f1757600080fd5b50614f2487828801614e03565b91505092959194509250565b600080600060608486031215614f4557600080fd5b8335614f5081614b8c565b92506020840135614f6081614b8c565b9150614f6e60408501614cc7565b90509250925092565b60008060008084860360a0811215614f8e57600080fd5b8535614f9981614b8c565b94506040601f1982011215614fad57600080fd5b506020850192506060850135614fc281614b8c565b9150614fd060808601614cc7565b905092959194509250565b60008060008060608587031215614ff157600080fd5b8435614ffc81614b8c565b935060208501356001600160401b038082111561501857600080fd5b61502488838901614dbf565b9095509350604087013591508082111561503d57600080fd5b5085016060818803121561505057600080fd5b939692955090935050565b6000806040838503121561506e57600080fd5b823561507981614b8c565b946020939093013593505050565b60008060008060006080868803121561509f57600080fd5b85356150aa81614b8c565b945060208601356150ba81614b8c565b935060408601356001600160401b03808211156150d657600080fd5b6150e289838a01614dbf565b909550935060608801359150808211156150fb57600080fd5b5061510888828901614e03565b9150509295509295909350565b6000806040838503121561512857600080fd5b823561513381614b8c565b9150602083013561514381614b8c565b809150509250929050565b634e487b7160e01b600052602160045260246000fd5b602081016002831061518657634e487b7160e01b600052602160045260246000fd5b91905290565b60006020828403121561519e57600080fd5b813560ff8116811461157157600080fd5b600080604083850312156151c257600080fd5b82356151cd81614b8c565b91506151db60208401614cc7565b90509250929050565b81516001600160a01b0316815260208083015163ffffffff169082015260408101610a76565b60008060008060006080868803121561522257600080fd5b853561522d81614b8c565b945060208601356001600160401b0381111561524857600080fd5b61525488828901614dbf565b9699909850959660408101359660609091013595509350505050565b6000806000806080858703121561528657600080fd5b843561529181614b8c565b935060208501356152a181614b8c565b93969395505050506040820135916060013590565b600080602083850312156152c957600080fd5b82356001600160401b03808211156152e057600080fd5b818501915085601f8301126152f457600080fd5b81358181111561530357600080fd5b86602082850101111561531557600080fd5b60209290920196919550909350505050565b6000806020838503121561533a57600080fd5b82356001600160401b0381111561535057600080fd5b61535c85828601614dbf565b90969095509350505050565b60008060008060006080868803121561538057600080fd5b853561538b81614b8c565b945061539960208701614cc7565b935060408601356001600160401b038111156153b457600080fd5b6153c088828901614dbf565b909450925050606086013561ffff811681146153db57600080fd5b809150509295509295909350565b6000806000604084860312156153fe57600080fd5b833561540981614b8c565b925060208401356001600160401b0381111561542457600080fd5b61543086828701614dbf565b9497909650939450505050565b6000806000806040858703121561545357600080fd5b84356001600160401b038082111561546a57600080fd5b61547688838901614dbf565b9096509450602087013591508082111561548f57600080fd5b5061549c87828801614dbf565b95989497509550505050565b634e487b7160e01b600052601160045260246000fd5b6000828210156154d0576154d06154a8565b500390565b600082198211156154e8576154e86154a8565b500190565b634e487b7160e01b600052603260045260246000fd5b6000600019821415615517576155176154a8565b5060010190565b60006020828403121561553057600080fd5b815161157181614b8c565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b600063ffffffff8083168185168083038211156155a4576155a46154a8565b01949350505050565b6000602082840312156155bf57600080fd5b8151801515811461157157600080fd5b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b60208082526019908201527f5061757361626c653a20696e6465782069732070617573656400000000000000604082015260600190565b60006020828403121561566057600080fd5b61157182614cc7565b6000816000190483118215151615615683576156836154a8565b500290565b634e487b7160e01b600052601260045260246000fd5b6000826156bb57634e487b7160e01b600052601260045260246000fd5b500490565b60008235605e198336030181126156d657600080fd5b9190910192915050565b8183526000602080850194508260005b858110156157195763ffffffff61570683614cc7565b16875295820195908201906001016156f0565b509495945050505050565b8681526001600160a01b038616602082015260a06040820181905260009061574f90830186886156e0565b60608301949094525060800152949350505050565b60208152816020820152818360408301376000818301604090810191909152601f909201601f19160101919050565b60006001600160401b03838116908316818110156157b3576157b36154a8565b039392505050565b60006001600160401b038083168185168083038211156155a4576155a46154a8565b6000816157ec576157ec6154a8565b506000190190565b60006001600160e01b03838116908316818110156157b3576157b36154a8565b6000808335601e1984360301811261582b57600080fd5b8301803591506001600160401b0382111561584557600080fd5b6020019150600581901b36038213156143a557600080fd5b6020815260006158716020830184866156e0565b949350505050565b82815260006020604081840152835180604085015260005b818110156158ad57858101830151858201606001528201615891565b818111156158bf576000606083870101525b50601f01601f191692909201606001949350505050565b6000602082840312156158e857600080fd5b81516001600160e01b03198116811461157157600080fd5b6000808335601e1984360301811261591757600080fd5b8301803591506001600160401b0382111561593157600080fd5b6020019150600681901b36038213156143a557600080fd5b60006020828403121561595b57600080fd5b81356001600160401b038116811461157157600080fd5b60006001600160e01b038281168482168083038211156155a4576155a46154a8565b805160208083015191908110156159b5576000198160200360031b1b821691505b50919050565b634e487b7160e01b600052603160045260246000fdfe4156534469726563746f72792e72656769737465724f70657261746f72546f4fa264697066735822122088c48cb837a3e7abf13048ac5ede39b27b6384be7644eb089d9768d8d5f6369764736f6c634300080c0033",
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

// MAXPENDINGUPDATES is a free data retrieval call binding the contract method 0xb3e52eaa.
//
// Solidity: function MAX_PENDING_UPDATES() view returns(uint256)
func (_AVSDirectory *AVSDirectoryCaller) MAXPENDINGUPDATES(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AVSDirectory.contract.Call(opts, &out, "MAX_PENDING_UPDATES")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXPENDINGUPDATES is a free data retrieval call binding the contract method 0xb3e52eaa.
//
// Solidity: function MAX_PENDING_UPDATES() view returns(uint256)
func (_AVSDirectory *AVSDirectorySession) MAXPENDINGUPDATES() (*big.Int, error) {
	return _AVSDirectory.Contract.MAXPENDINGUPDATES(&_AVSDirectory.CallOpts)
}

// MAXPENDINGUPDATES is a free data retrieval call binding the contract method 0xb3e52eaa.
//
// Solidity: function MAX_PENDING_UPDATES() view returns(uint256)
func (_AVSDirectory *AVSDirectoryCallerSession) MAXPENDINGUPDATES() (*big.Int, error) {
	return _AVSDirectory.Contract.MAXPENDINGUPDATES(&_AVSDirectory.CallOpts)
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

<<<<<<< HEAD
<<<<<<< HEAD
// GetNumOperatorsInOperatorSet is a free data retrieval call binding the contract method 0x1023aa35.
//
// Solidity: function getNumOperatorsInOperatorSet((address,uint32) operatorSet) view returns(uint256)
func (_AVSDirectory *AVSDirectoryCaller) GetNumOperatorsInOperatorSet(opts *bind.CallOpts, operatorSet IAVSDirectoryOperatorSet) (*big.Int, error) {
	var out []interface{}
	err := _AVSDirectory.contract.Call(opts, &out, "getNumOperatorsInOperatorSet", operatorSet)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumOperatorsInOperatorSet is a free data retrieval call binding the contract method 0x1023aa35.
//
// Solidity: function getNumOperatorsInOperatorSet((address,uint32) operatorSet) view returns(uint256)
func (_AVSDirectory *AVSDirectorySession) GetNumOperatorsInOperatorSet(operatorSet IAVSDirectoryOperatorSet) (*big.Int, error) {
	return _AVSDirectory.Contract.GetNumOperatorsInOperatorSet(&_AVSDirectory.CallOpts, operatorSet)
}

// GetNumOperatorsInOperatorSet is a free data retrieval call binding the contract method 0x1023aa35.
//
// Solidity: function getNumOperatorsInOperatorSet((address,uint32) operatorSet) view returns(uint256)
func (_AVSDirectory *AVSDirectoryCallerSession) GetNumOperatorsInOperatorSet(operatorSet IAVSDirectoryOperatorSet) (*big.Int, error) {
	return _AVSDirectory.Contract.GetNumOperatorsInOperatorSet(&_AVSDirectory.CallOpts, operatorSet)
}

// GetOperatorSetsOfOperator is a free data retrieval call binding the contract method 0x16ae76cb.
//
// Solidity: function getOperatorSetsOfOperator(address operator, uint256 start, uint256 length) view returns((address,uint32)[] operatorSets)
func (_AVSDirectory *AVSDirectoryCaller) GetOperatorSetsOfOperator(opts *bind.CallOpts, operator common.Address, start *big.Int, length *big.Int) ([]IAVSDirectoryOperatorSet, error) {
	var out []interface{}
	err := _AVSDirectory.contract.Call(opts, &out, "getOperatorSetsOfOperator", operator, start, length)

	if err != nil {
		return *new([]IAVSDirectoryOperatorSet), err
	}

	out0 := *abi.ConvertType(out[0], new([]IAVSDirectoryOperatorSet)).(*[]IAVSDirectoryOperatorSet)

	return out0, err

}

// GetOperatorSetsOfOperator is a free data retrieval call binding the contract method 0x16ae76cb.
//
// Solidity: function getOperatorSetsOfOperator(address operator, uint256 start, uint256 length) view returns((address,uint32)[] operatorSets)
func (_AVSDirectory *AVSDirectorySession) GetOperatorSetsOfOperator(operator common.Address, start *big.Int, length *big.Int) ([]IAVSDirectoryOperatorSet, error) {
	return _AVSDirectory.Contract.GetOperatorSetsOfOperator(&_AVSDirectory.CallOpts, operator, start, length)
}

// GetOperatorSetsOfOperator is a free data retrieval call binding the contract method 0x16ae76cb.
//
// Solidity: function getOperatorSetsOfOperator(address operator, uint256 start, uint256 length) view returns((address,uint32)[] operatorSets)
func (_AVSDirectory *AVSDirectoryCallerSession) GetOperatorSetsOfOperator(operator common.Address, start *big.Int, length *big.Int) ([]IAVSDirectoryOperatorSet, error) {
	return _AVSDirectory.Contract.GetOperatorSetsOfOperator(&_AVSDirectory.CallOpts, operator, start, length)
}

// GetOperatorsInOperatorSet is a free data retrieval call binding the contract method 0x7357723b.
//
// Solidity: function getOperatorsInOperatorSet((address,uint32) operatorSet, uint256 start, uint256 length) view returns(address[] operators)
func (_AVSDirectory *AVSDirectoryCaller) GetOperatorsInOperatorSet(opts *bind.CallOpts, operatorSet IAVSDirectoryOperatorSet, start *big.Int, length *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _AVSDirectory.contract.Call(opts, &out, "getOperatorsInOperatorSet", operatorSet, start, length)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOperatorsInOperatorSet is a free data retrieval call binding the contract method 0x7357723b.
//
// Solidity: function getOperatorsInOperatorSet((address,uint32) operatorSet, uint256 start, uint256 length) view returns(address[] operators)
func (_AVSDirectory *AVSDirectorySession) GetOperatorsInOperatorSet(operatorSet IAVSDirectoryOperatorSet, start *big.Int, length *big.Int) ([]common.Address, error) {
	return _AVSDirectory.Contract.GetOperatorsInOperatorSet(&_AVSDirectory.CallOpts, operatorSet, start, length)
}

// GetOperatorsInOperatorSet is a free data retrieval call binding the contract method 0x7357723b.
//
// Solidity: function getOperatorsInOperatorSet((address,uint32) operatorSet, uint256 start, uint256 length) view returns(address[] operators)
func (_AVSDirectory *AVSDirectoryCallerSession) GetOperatorsInOperatorSet(operatorSet IAVSDirectoryOperatorSet, start *big.Int, length *big.Int) ([]common.Address, error) {
	return _AVSDirectory.Contract.GetOperatorsInOperatorSet(&_AVSDirectory.CallOpts, operatorSet, start, length)
}

// InTotalOperatorSets is a free data retrieval call binding the contract method 0xcbdf0e42.
=======
// IsMember is a free data retrieval call binding the contract method 0x3c4385d0.
>>>>>>> ce94473c (feat: operator commission bips (#627))
=======
// FreeMagnitude is a free data retrieval call binding the contract method 0x5fd6abfd.
>>>>>>> 6b2ab7ed (build: bindings)
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

<<<<<<< HEAD
<<<<<<< HEAD
// OperatorSetMemberAtIndex is a free data retrieval call binding the contract method 0x411d415b.
//
// Solidity: function operatorSetMemberAtIndex((address,uint32) operatorSet, uint256 index) view returns(address)
func (_AVSDirectory *AVSDirectoryCaller) OperatorSetMemberAtIndex(opts *bind.CallOpts, operatorSet IAVSDirectoryOperatorSet, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AVSDirectory.contract.Call(opts, &out, "operatorSetMemberAtIndex", operatorSet, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
=======
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
>>>>>>> 6b2ab7ed (build: bindings)

	return out0, err

}

<<<<<<< HEAD
// OperatorSetMemberAtIndex is a free data retrieval call binding the contract method 0x411d415b.
//
// Solidity: function operatorSetMemberAtIndex((address,uint32) operatorSet, uint256 index) view returns(address)
func (_AVSDirectory *AVSDirectorySession) OperatorSetMemberAtIndex(operatorSet IAVSDirectoryOperatorSet, index *big.Int) (common.Address, error) {
	return _AVSDirectory.Contract.OperatorSetMemberAtIndex(&_AVSDirectory.CallOpts, operatorSet, index)
}

// OperatorSetMemberAtIndex is a free data retrieval call binding the contract method 0x411d415b.
//
// Solidity: function operatorSetMemberAtIndex((address,uint32) operatorSet, uint256 index) view returns(address)
func (_AVSDirectory *AVSDirectoryCallerSession) OperatorSetMemberAtIndex(operatorSet IAVSDirectoryOperatorSet, index *big.Int) (common.Address, error) {
	return _AVSDirectory.Contract.OperatorSetMemberAtIndex(&_AVSDirectory.CallOpts, operatorSet, index)
}

// OperatorSetsMemberOfAtIndex is a free data retrieval call binding the contract method 0xb5a768ca.
//
// Solidity: function operatorSetsMemberOfAtIndex(address operator, uint256 index) view returns((address,uint32))
func (_AVSDirectory *AVSDirectoryCaller) OperatorSetsMemberOfAtIndex(opts *bind.CallOpts, operator common.Address, index *big.Int) (IAVSDirectoryOperatorSet, error) {
	var out []interface{}
	err := _AVSDirectory.contract.Call(opts, &out, "operatorSetsMemberOfAtIndex", operator, index)
=======
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
>>>>>>> 6b2ab7ed (build: bindings)

	if err != nil {
		return *new(IAVSDirectoryOperatorSet), err
	}

	out0 := *abi.ConvertType(out[0], new(IAVSDirectoryOperatorSet)).(*IAVSDirectoryOperatorSet)

	return out0, err

}

<<<<<<< HEAD
// OperatorSetsMemberOfAtIndex is a free data retrieval call binding the contract method 0xb5a768ca.
//
// Solidity: function operatorSetsMemberOfAtIndex(address operator, uint256 index) view returns((address,uint32))
func (_AVSDirectory *AVSDirectorySession) OperatorSetsMemberOfAtIndex(operator common.Address, index *big.Int) (IAVSDirectoryOperatorSet, error) {
	return _AVSDirectory.Contract.OperatorSetsMemberOfAtIndex(&_AVSDirectory.CallOpts, operator, index)
}

// OperatorSetsMemberOfAtIndex is a free data retrieval call binding the contract method 0xb5a768ca.
//
// Solidity: function operatorSetsMemberOfAtIndex(address operator, uint256 index) view returns((address,uint32))
func (_AVSDirectory *AVSDirectoryCallerSession) OperatorSetsMemberOfAtIndex(operator common.Address, index *big.Int) (IAVSDirectoryOperatorSet, error) {
	return _AVSDirectory.Contract.OperatorSetsMemberOfAtIndex(&_AVSDirectory.CallOpts, operator, index)
}

=======
>>>>>>> ce94473c (feat: operator commission bips (#627))
=======
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

>>>>>>> 6b2ab7ed (build: bindings)
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

// Deallocate is a paid mutator transaction binding the contract method 0x3367a33c.
//
// Solidity: function deallocate(address operator, (address,(address,uint32)[],uint64[])[] deallocations, (bytes,bytes32,uint256) operatorSignature) returns()
func (_AVSDirectory *AVSDirectoryTransactor) Deallocate(opts *bind.TransactOpts, operator common.Address, deallocations []IAVSDirectoryMagnitudeAdjustment, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _AVSDirectory.contract.Transact(opts, "deallocate", operator, deallocations, operatorSignature)
}

// Deallocate is a paid mutator transaction binding the contract method 0x3367a33c.
//
// Solidity: function deallocate(address operator, (address,(address,uint32)[],uint64[])[] deallocations, (bytes,bytes32,uint256) operatorSignature) returns()
func (_AVSDirectory *AVSDirectorySession) Deallocate(operator common.Address, deallocations []IAVSDirectoryMagnitudeAdjustment, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _AVSDirectory.Contract.Deallocate(&_AVSDirectory.TransactOpts, operator, deallocations, operatorSignature)
}

// Deallocate is a paid mutator transaction binding the contract method 0x3367a33c.
//
// Solidity: function deallocate(address operator, (address,(address,uint32)[],uint64[])[] deallocations, (bytes,bytes32,uint256) operatorSignature) returns()
func (_AVSDirectory *AVSDirectoryTransactorSession) Deallocate(operator common.Address, deallocations []IAVSDirectoryMagnitudeAdjustment, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _AVSDirectory.Contract.Deallocate(&_AVSDirectory.TransactOpts, operator, deallocations, operatorSignature)
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

// FreeDeallocatedMagnitude is a paid mutator transaction binding the contract method 0xcb5b37be.
//
// Solidity: function freeDeallocatedMagnitude(address operator, address[] strategies) returns()
func (_AVSDirectory *AVSDirectoryTransactor) FreeDeallocatedMagnitude(opts *bind.TransactOpts, operator common.Address, strategies []common.Address) (*types.Transaction, error) {
	return _AVSDirectory.contract.Transact(opts, "freeDeallocatedMagnitude", operator, strategies)
}

// FreeDeallocatedMagnitude is a paid mutator transaction binding the contract method 0xcb5b37be.
//
// Solidity: function freeDeallocatedMagnitude(address operator, address[] strategies) returns()
func (_AVSDirectory *AVSDirectorySession) FreeDeallocatedMagnitude(operator common.Address, strategies []common.Address) (*types.Transaction, error) {
	return _AVSDirectory.Contract.FreeDeallocatedMagnitude(&_AVSDirectory.TransactOpts, operator, strategies)
}

// FreeDeallocatedMagnitude is a paid mutator transaction binding the contract method 0xcb5b37be.
//
// Solidity: function freeDeallocatedMagnitude(address operator, address[] strategies) returns()
func (_AVSDirectory *AVSDirectoryTransactorSession) FreeDeallocatedMagnitude(operator common.Address, strategies []common.Address) (*types.Transaction, error) {
	return _AVSDirectory.Contract.FreeDeallocatedMagnitude(&_AVSDirectory.TransactOpts, operator, strategies)
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
