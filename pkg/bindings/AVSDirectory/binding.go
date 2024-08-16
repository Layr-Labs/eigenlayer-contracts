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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_delegation\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ALLOCATION_DELAY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DOMAIN_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAGNITUDE_ADJUSTMENT_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_PENDING_UPDATES\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"OPERATOR_AVS_REGISTRATION_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"OPERATOR_SET_FORCE_DEREGISTRATION_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"OPERATOR_SET_REGISTRATION_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allocate\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allocations\",\"type\":\"tuple[]\",\"internalType\":\"structIAVSDirectory.MagnitudeAdjustment[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"operatorSets\",\"type\":\"tuple[]\",\"internalType\":\"structIAVSDirectory.OperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"magnitudeDiffs\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}]},{\"name\":\"allocatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allocatorSaltIsSpent\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"avsOperatorStatus\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIAVSDirectory.OperatorAVSRegistrationStatus\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"becomeOperatorSetAVS\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"calculateMagnitudeAdjustmentDigestHash\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"adjustments\",\"type\":\"tuple[]\",\"internalType\":\"structIAVSDirectory.MagnitudeAdjustment[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"operatorSets\",\"type\":\"tuple[]\",\"internalType\":\"structIAVSDirectory.OperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"magnitudeDiffs\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}]},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateOperatorAVSRegistrationDigestHash\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateOperatorSetForceDeregistrationTypehash\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateOperatorSetRegistrationDigestHash\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cancelSalt\",\"inputs\":[{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createOperatorSets\",\"inputs\":[{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deallocate\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"deallocations\",\"type\":\"tuple[]\",\"internalType\":\"structIAVSDirectory.MagnitudeAdjustment[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"operatorSets\",\"type\":\"tuple[]\",\"internalType\":\"structIAVSDirectory.OperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"magnitudeDiffs\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}]},{\"name\":\"allocatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregisterOperatorFromOperatorSets\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"domainSeparator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"forceDeregisterFromOperatorSets\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"freeMagnitude\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSlashableBips\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structIAVSDirectory.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"timestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"inTotalOperatorSets\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"initialPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isMember\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structIAVSDirectory.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperatorSet\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperatorSetAVS\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperatorSlashable\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structIAVSDirectory.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"migrateOperatorsToOperatorSets\",\"inputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"operatorSaltIsSpent\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorSetMemberCount\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorSetStatus\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"registered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"lastDeregisteredTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorSetsMemberOf\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"operatorSets\",\"type\":\"tuple[]\",\"internalType\":\"structIAVSDirectory.OperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorSetsMemberOf\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIAVSDirectory.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerOperatorToOperatorSets\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slashOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"bipsToSlash\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAVSMetadataURI\",\"inputs\":[{\"name\":\"metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateFreeMagnitude\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"numToComplete\",\"type\":\"uint8[]\",\"internalType\":\"uint8[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AVSMetadataURIUpdated\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"metadataURI\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AVSMigratedToOperatorSets\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MagnitudeAllocated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIAVSDirectory.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"magnitudeToAllocate\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"effectTimestamp\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MagnitudeDeallocationCompleted\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIAVSDirectory.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"freeMagnitudeAdded\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MagnitudeQueueDeallocated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIAVSDirectory.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"magnitudeToDeallocate\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"completableTimestamp\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorAVSRegistrationStatusUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"status\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumIAVSDirectory.OperatorAVSRegistrationStatus\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorAddedToOperatorSet\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIAVSDirectory.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorMigratedToOperatorSets\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"indexed\":false,\"internalType\":\"uint32[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRemovedFromOperatorSet\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIAVSDirectory.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSetCreated\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIAVSDirectory.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSlashed\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"bipsToSlash\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x60c06040523480156200001157600080fd5b50604051620062d8380380620062d8833981016040819052620000349162000118565b6001600160a01b0381166080526200004b62000056565b504660a0526200014a565b600054610100900460ff1615620000c35760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff908116101562000116576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6000602082840312156200012b57600080fd5b81516001600160a01b03811681146200014357600080fd5b9392505050565b60805160a05161615362000185600039600061295f0152600081816107fd01528181610e4e015281816111d8015261163701526161536000f3fe608060405234801561001057600080fd5b50600436106102d65760003560e01c80638522fcb111610182578063c1a8e2c5116100e9578063dce974b9116100a2578063ef2dfa8d1161007c578063ef2dfa8d14610853578063f2fde38b14610866578063f698da2514610879578063fabc1cbc1461088157600080fd5b8063dce974b9146107d1578063df5cf723146107f8578063ec76f4421461081f57600080fd5b8063c1a8e2c51461071f578063c825fe6814610732578063cbdf0e4214610759578063d79aceab1461076c578063da2ff05d14610793578063dae226b6146107a657600080fd5b8063a98fb3551161013b578063a98fb355146106c3578063aec205c5146106d6578063afe02ed5146106de578063b2841d48146106f1578063b3e52eaa14610704578063bd74a06c1461070c57600080fd5b80638522fcb11461062e578063886f1195146106415780638da5cb5b1461066c5780638de549441461067d578063955e66961461069d578063a1060c88146106b057600080fd5b806346fb3a00116102415780636fbb031c116101fa5780637673e93a116101d45780637673e93a146105a35780637b205de3146105c657806384b78fc2146105ed57806384d76f7b1461060057600080fd5b80636fbb031c146105695780637019670814610588578063715018a61461059b57600080fd5b806346fb3a001461048157806349075da3146104af578063595c6a67146104ea5780635ac86ab7146104f25780635c975abb146105155780635fd6abfd1461051d57600080fd5b80631e68134e116102935780631e68134e1461037557806320606b70146103d257806333429a6a146104075780633367a33c1461042d578063374823b5146104405780633fee332d1461046e57600080fd5b80630d5387c5146102db57806310d67a2f146103045780631352c3e614610319578063136439dd1461033c5780631794bb3c1461034f5780631e2199e214610362575b600080fd5b6102ee6102e9366004614fe1565b610894565b6040516102fb9190615016565b60405180910390f35b61031761031236600461507c565b6109d9565b005b61032c610327366004615120565b610a95565b60405190151581526020016102fb565b61031761034a3660046151a5565b610b25565b61031761035d3660046151be565b610c64565b6103176103703660046152f5565b610d8e565b6103b6610383366004615370565b609e60209081526000938452604080852082529284528284209052825290205460ff811690610100900463ffffffff1682565b60408051921515835263ffffffff9091166020830152016102fb565b6103f97f8cad95687ba82c2ce50e74f7b754645e5117c3a5bec8151c0726d5857980a86681565b6040519081526020016102fb565b61041a6104153660046153b7565b6110ae565b60405161ffff90911681526020016102fb565b61031761043b36600461541b565b6111b6565b61032c61044e36600461549b565b609960209081526000928352604080842090915290825290205460ff1681565b61031761047c3660046154c7565b611304565b61032c61048f36600461549b565b60a560209081526000928352604080842090915290825290205460ff1681565b6104dd6104bd366004615555565b609860209081526000928352604080842090915290825290205460ff1681565b6040516102fb91906155a4565b61031761154e565b61032c6105003660046155cc565b606654600160ff9092169190911b9081161490565b6066546103f9565b61055161052b366004615555565b60a16020908152600092835260408084209091529082529020546001600160401b031681565b6040516001600160401b0390911681526020016102fb565b610573621baf8081565b60405163ffffffff90911681526020016102fb565b61031761059636600461541b565b611615565b610317611733565b61032c6105b136600461507c565b609a6020526000908152604090205460ff1681565b6103f97fb3cc2868468e0461dcbecef0a4fcb7a8ef26c2e1f3e59c9d0de6afcc74448a2481565b6103176105fb3660046155ef565b611747565b61032c61060e366004615671565b609b60209081526000928352604080842090915290825290205460ff1681565b6103f961063c3660046156a6565b6117b9565b606554610654906001600160a01b031681565b6040516001600160a01b0390911681526020016102fb565b6033546001600160a01b0316610654565b61069061068b36600461549b565b61181e565b6040516102fb919061570c565b6103f96106ab3660046156a6565b61185f565b6103f96106be366004615732565b61189f565b6103176106d1366004615778565b611909565b610317611950565b6103176106ec3660046157e9565b611a18565b6103f96106ff3660046156a6565b611be5565b6103f9600181565b61031761071a36600461582a565b611c25565b61031761072d3660046158ab565b612158565b6103f97f809c5ac049c45b7a7f050a20f00c16cf63797efbf8b1eb8d749fdfa39ff8f92981565b6103f961076736600461507c565b61218d565b6103f97fda2c89bafdd34776a2b8bb9c83c82f419e20cc8c67207f70edd58249b92661bd81565b61032c6107a1366004615120565b6121ae565b6103f96107b4366004615671565b609c60209081526000928352604080842090915290825290205481565b6103f97f4ee65f64218c67b68da66fd0db16560040a6b973290b9e71912d661ee53fe49581565b6106547f000000000000000000000000000000000000000000000000000000000000000081565b61031761082d3660046151a5565b33600090815260996020908152604080832093835292905220805460ff19166001179055565b6103176108613660046158ff565b6121da565b61031761087436600461507c565b612583565b6103f96125f9565b61031761088f3660046151a5565b612608565b6001600160a01b0383166000908152609d602052604081206060919084906108bb90612764565b6108c59190615980565b9050808311156108d3578092505b826001600160401b038111156108eb576108eb615099565b60405190808252806020026020018201604052801561093057816020015b60408051808201909152600080825260208201528152602001906001900390816109095790505b50915060005b838110156109d0576109a261096c61094e8388615997565b6001600160a01b0389166000908152609d602052604090209061276e565b60408051808201909152600080825260208201525060408051808201909152606082901c815263ffffffff909116602082015290565b8382815181106109b4576109b46159af565b6020026020010181905250806109c9906159c5565b9050610936565b50509392505050565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610a2c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a5091906159e0565b6001600160a01b0316336001600160a01b031614610a895760405162461bcd60e51b8152600401610a80906159fd565b60405180910390fd5b610a928161277a565b50565b80516001600160a01b039081166000908152609e6020908152604080832093861683529281528282208185015163ffffffff908116845290825283832084518086019095525460ff8116151585526101009004169083015290610af884846121ae565b80610b1b575042621baf808260200151610b129190615a47565b63ffffffff1610155b9150505b92915050565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015610b6d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b919190615a6f565b610bad5760405162461bcd60e51b8152600401610a8090615a91565b60665481811614610c265760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c69747900000000000000006064820152608401610a80565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b600054610100900460ff1615808015610c845750600054600160ff909116105b80610c9e5750303b158015610c9e575060005460ff166001145b610d015760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610a80565b6000805460ff191660011790558015610d24576000805461ff0019166101001790555b610d2e8383612871565b610d3661295b565b609755610d4284612a24565b8015610d88576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050565b60665460019060029081161415610db75760405162461bcd60e51b8152600401610a8090615ad9565b4282604001511015610e2f5760405162461bcd60e51b815260206004820152604760248201526000805160206160fe83398151915260448201527f70657261746f72536574733a206f70657261746f72207369676e617475726520606482015266195e1c1a5c995960ca1b608482015260a401610a80565b6040516336b87bd760e11b81526001600160a01b0386811660048301527f00000000000000000000000000000000000000000000000000000000000000001690636d70f7ae90602401602060405180830381865afa158015610e95573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610eb99190615a6f565b610f385760405162461bcd60e51b815260206004820152605660248201526000805160206160fe83398151915260448201527f70657261746f72536574733a206f70657261746f72206e6f7420726567697374606482015275195c9959081d1bc8115a59d95b93185e595c881e595d60521b608482015260a401610a80565b336000908152609a602052604090205460ff16610fbf5760405162461bcd60e51b815260206004820152604b60248201526000805160206160fe83398151915260448201527f70657261746f72536574733a20415653206973206e6f7420616e206f7065726160648201526a746f72207365742041565360a81b608482015260a401610a80565b6001600160a01b038516600090815260996020908152604080832085830151845290915290205460ff161561104a5760405162461bcd60e51b815260206004820152603f60248201526000805160206160fe83398151915260448201527f70657261746f72536574733a2073616c7420616c7265616479207370656e74006064820152608401610a80565b61106a856110633387878760200151886040015161185f565b8451612a76565b6001600160a01b03851660009081526099602090815260408083208583015184529091529020805460ff191660011790556110a785338686612c30565b5050505050565b6001600160a01b038085166000908152609f60209081526040808320938616835292905290812081906110e19084612fff565b6001600160a01b03808816600090815260a06020908152604080832093891683529281529181209293509161117e918691908490611121908b018b61507c565b6001600160a01b03166001600160a01b0316815260200190815260200160002060008960200160208101906111569190615b10565b63ffffffff1663ffffffff168152602001908152602001600020612fff90919063ffffffff16565b9050816001600160401b0316612710826001600160401b03166111a19190615b2b565b6111ab9190615b60565b979650505050505050565b6040516369066d9160e01b81526001600160a01b0385811660048301526000917f0000000000000000000000000000000000000000000000000000000000000000909116906369066d9190602401602060405180830381865afa158015611221573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061124591906159e0565b9050336001600160a01b03821614611264576112648186868686613057565b6000611273621baf8042615a47565b905060005b848110156112fb576112bd87878784818110611296576112966159af565b90506020028101906112a89190615b82565b6112b690602081019061507c565b600061320b565b6112eb878787848181106112d3576112d36159af565b90506020028101906112e59190615b82565b846133b2565b6112f4816159c5565b9050611278565b50505050505050565b6066546001906002908116141561132d5760405162461bcd60e51b8152600401610a8090615ad9565b8151516113c557336001600160a01b038716146113c05760405162461bcd60e51b815260206004820152604560248201527f4156534469726563746f72792e666f7263654465726567697374657246726f6d60448201527f4f70657261746f72536574733a2063616c6c6572206d757374206265206f70656064820152643930ba37b960d91b608482015260a401610a80565b61153a565b42826040015110156114505760405162461bcd60e51b815260206004820152604860248201527f4156534469726563746f72792e666f7263654465726567697374657246726f6d60448201527f4f70657261746f72536574733a206f70657261746f72207369676e617475726560648201526708195e1c1a5c995960c21b608482015260a401610a80565b6001600160a01b038616600090815260996020908152604080832085830151845290915290205460ff16156114ef576040805162461bcd60e51b81526020600482015260248101919091527f4156534469726563746f72792e666f7263654465726567697374657246726f6d60448201527f4f70657261746f72536574733a2073616c7420616c7265616479207370656e746064820152608401610a80565b6115088661106387878787602001518860400151611be5565b6001600160a01b03861660009081526099602090815260408083208583015184529091529020805460ff191660011790555b61154685878686613983565b505050505050565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015611596573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115ba9190615a6f565b6115d65760405162461bcd60e51b8152600401610a8090615a91565b600019606681905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b6040516369066d9160e01b81526001600160a01b0385811660048301526000917f0000000000000000000000000000000000000000000000000000000000000000909116906369066d9190602401602060405180830381865afa158015611680573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116a491906159e0565b9050336001600160a01b038216146116c3576116c38186868686613057565b60006116d2621baf8042615a47565b905060005b848110156112fb576116f587878784818110611296576112966159af565b6117238787878481811061170b5761170b6159af565b905060200281019061171d9190615b82565b84613b65565b61172c816159c5565b90506116d7565b61173b614157565b6117456000612a24565b565b60005b83811015611546576117a986868684818110611768576117686159af565b905060200201602081019061177d919061507c565b85858581811061178f5761178f6159af565b90506020020160208101906117a491906155cc565b61320b565b6117b2816159c5565b905061174a565b60006118147fb3cc2868468e0461dcbecef0a4fcb7a8ef26c2e1f3e59c9d0de6afcc74448a2487878787876040516020016117f996959493929190615c9e565b604051602081830303815290604052805190602001206141b1565b9695505050505050565b60408051808201909152600080825260208201526001600160a01b0383166000908152609d602052604090206118589061096c908461276e565b9392505050565b60006118147f809c5ac049c45b7a7f050a20f00c16cf63797efbf8b1eb8d749fdfa39ff8f92987878787876040516020016117f996959493929190615df8565b604080517fda2c89bafdd34776a2b8bb9c83c82f419e20cc8c67207f70edd58249b92661bd60208201526001600160a01b038087169282019290925290841660608201526080810183905260a081018290526000906119009060c0016117f9565b95945050505050565b336001600160a01b03167fa89c1dc243d8908a96dd84944bcc97d6bc6ac00dd78e20621576be6a3c9437138383604051611944929190615e38565b60405180910390a25050565b336000908152609a602052604090205460ff16156119d65760405162461bcd60e51b815260206004820152603e60248201527f4156534469726563746f72792e6265636f6d654f70657261746f72536574415660448201527f533a20616c726561647920616e206f70657261746f72207365742041565300006064820152608401610a80565b336000818152609a6020526040808220805460ff19166001179055517f702b0c1f6cb1cf511aaa81f72bc05a215bb3497632d72c690c822b044ab494bf9190a2565b60005b81811015611be057336000908152609b6020526040812090848484818110611a4557611a456159af565b9050602002016020810190611a5a9190615b10565b63ffffffff16815260208101919091526040016000205460ff1615611ae75760405162461bcd60e51b815260206004820152603b60248201527f4156534469726563746f72792e6372656174654f70657261746f725365743a2060448201527f6f70657261746f722073657420616c72656164792065786973747300000000006064820152608401610a80565b336000908152609b60205260408120600191858585818110611b0b57611b0b6159af565b9050602002016020810190611b209190615b10565b63ffffffff1663ffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055507f31629285ead2335ae0933f86ed2ae63321f7af77b4e6eaabc42c057880977e6c6040518060400160405280336001600160a01b03168152602001858585818110611b9e57611b9e6159af565b9050602002016020810190611bb39190615b10565b63ffffffff169052604051611bc8919061570c565b60405180910390a1611bd9816159c5565b9050611a1b565b505050565b60006118147f4ee65f64218c67b68da66fd0db16560040a6b973290b9e71912d661ee53fe49587878787876040516020016117f996959493929190615df8565b611c52856040518060400160405280336001600160a01b031681526020018763ffffffff16815250610a95565b611ccf5760405162461bcd60e51b815260206004820152604260248201527f4156534469726563746f72792e736c6173684f70657261746f723a206f70657260448201527f61746f72206e6f7420736c61736861626c6520666f72206f70657261746f7253606482015261195d60f21b608482015260a401610a80565b60005b82811015611546576001600160a01b038616600090815260a0602052604081208190611d5d90429083898988818110611d0d57611d0d6159af565b9050602002016020810190611d22919061507c565b6001600160a01b0316815260208082019290925260409081016000908120338252835281812063ffffffff808e16835293522091906141f816565b9050612710611d796001600160401b03831661ffff8716615b2b565b611d839190615b60565b6001600160a01b038916600090815260a060205260408120919350611e139142916001600160401b03861691908a8a89818110611dc257611dc26159af565b9050602002016020810190611dd7919061507c565b6001600160a01b0316815260208082019290925260409081016000908120338252835281812063ffffffff808f16835293522092919061429d16565b506001600160a01b038716600090815260a46020526040812081878786818110611e3f57611e3f6159af565b9050602002016020810190611e54919061507c565b6001600160a01b0316815260208082019290925260409081016000908120338252835281812063ffffffff8b1682529092529020549050805b801561205e576001600160a01b038916600090815260a46020526040812081898988818110611ebe57611ebe6159af565b9050602002016020810190611ed3919061507c565b6001600160a01b0316815260208082019290925260409081016000908120338252835281812063ffffffff8d1682529092529020611f12600184615980565b81548110611f2257611f226159af565b60009182526020808320909101546001600160a01b038d16835260a290915260408220909250818a8a89818110611f5b57611f5b6159af565b9050602002016020810190611f70919061507c565b6001600160a01b03166001600160a01b031681526020019081526020016000208281548110611fa157611fa16159af565b6000918252602090912001805490915063ffffffff428116600160401b9092041610611fce57505061205e565b805460009061271090611fee906001600160401b031661ffff8b16615b2b565b611ff89190615b60565b8254909150819083906000906120189084906001600160401b0316615e67565b92506101000a8154816001600160401b0302191690836001600160401b0316021790555080866120489190615e8f565b95505050508061205790615eb1565b9050611e8d565b50506001600160a01b0387166000908152609f602052604081206121449142916001600160401b038516916120da91908a8a898181106120a0576120a06159af565b90506020020160208101906120b5919061507c565b6001600160a01b03166001600160a01b03168152602001908152602001600020614339565b6120e49190615ec8565b6001600160a01b038a166000908152609f602052604081209089898881811061210f5761210f6159af565b9050602002016020810190612124919061507c565b6001600160a01b0316815260208101919091526040016000209190614372565b50505080612151906159c5565b9050611cd2565b606654600190600290811614156121815760405162461bcd60e51b8152600401610a8090615ad9565b610d8833858585613983565b6001600160a01b0381166000908152609d60205260408120610b1f90612764565b60006118586121bc8361438d565b6001600160a01b0385166000908152609d60205260409020906143f2565b606654600190600290811614156122035760405162461bcd60e51b8152600401610a8090615ad9565b336000908152609a602052604090205460ff1661229c5760405162461bcd60e51b815260206004820152604b60248201527f4156534469726563746f72792e6d6967726174654f70657261746f7273546f4f60448201527f70657261746f72536574733a20415653206973206e6f7420616e206f7065726160648201526a746f72207365742041565360a81b608482015260a401610a80565b60005b84811015611546576001336000908152609860205260408120908888858181106122cb576122cb6159af565b90506020020160208101906122e0919061507c565b6001600160a01b0316815260208101919091526040016000205460ff16600181111561230e5761230e61558e565b146123ba5760405162461bcd60e51b815260206004820152606a60248201527f4156534469726563746f72792e6d6967726174654f70657261746f7273546f4f60448201527f70657261746f72536574733a206f70657261746f7220616c7265616479206d6960648201527f677261746564206f72206e6f742061206c656761637920726567697374657265608482015269321037b832b930ba37b960b11b60a482015260c401610a80565b61240e8686838181106123cf576123cf6159af565b90506020020160208101906123e4919061507c565b338686858181106123f7576123f76159af565b90506020028101906124099190615ee8565b612c30565b33600090815260986020526040812081888885818110612430576124306159af565b9050602002016020810190612445919061507c565b6001600160a01b031681526020810191909152604001600020805460ff1916600183818111156124775761247761558e565b02179055503386868381811061248f5761248f6159af565b90506020020160208101906124a4919061507c565b6001600160a01b03167ff0952b1c65271d819d39983d2abb044b9cace59bcc4d4dd389f586ebdcb15b4160006040516124dd91906155a4565b60405180910390a3338686838181106124f8576124f86159af565b905060200201602081019061250d919061507c565b6001600160a01b03167f54f33cfdd1ca703d795986b986fd47d742eab1904ecd2a5fdb8d6595e5904a01868685818110612549576125496159af565b905060200281019061255b9190615ee8565b604051612569929190615f31565b60405180910390a38061257b816159c5565b91505061229f565b61258b614157565b6001600160a01b0381166125f05760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610a80565b610a9281612a24565b600061260361295b565b905090565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561265b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061267f91906159e0565b6001600160a01b0316336001600160a01b0316146126af5760405162461bcd60e51b8152600401610a80906159fd565b60665419811960665419161461272d5760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c69747900000000000000006064820152608401610a80565b606681905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c90602001610c59565b6000610b1f825490565b6000611858838361440a565b6001600160a01b0381166128085760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a401610a80565b606554604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1606580546001600160a01b0319166001600160a01b0392909216919091179055565b6065546001600160a01b031615801561289257506001600160a01b03821615155b6129145760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a401610a80565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a26129578261277a565b5050565b60007f000000000000000000000000000000000000000000000000000000000000000046141561298c575060975490565b50604080518082018252600a81526922b4b3b2b72630bcb2b960b11b60209182015281517f8cad95687ba82c2ce50e74f7b754645e5117c3a5bec8151c0726d5857980a866818301527f71b625cfad44bac63b13dba07f2e1d6084ee04b6f8752101ece6126d584ee6ea81840152466060820152306080808301919091528351808303909101815260a0909101909252815191012090565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6001600160a01b0383163b15612b9057604051630b135d3f60e11b808252906001600160a01b03851690631626ba7e90612ab69086908690600401615f4d565b602060405180830381865afa158015612ad3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612af79190615faa565b6001600160e01b03191614611be05760405162461bcd60e51b815260206004820152605360248201527f454950313237315369676e61747572655574696c732e636865636b5369676e6160448201527f747572655f454950313237313a2045524331323731207369676e6174757265206064820152721d995c9a599a58d85d1a5bdb8819985a5b1959606a1b608482015260a401610a80565b826001600160a01b0316612ba48383614434565b6001600160a01b031614611be05760405162461bcd60e51b815260206004820152604760248201527f454950313237315369676e61747572655574696c732e636865636b5369676e6160448201527f747572655f454950313237313a207369676e6174757265206e6f742066726f6d6064820152661039b4b3b732b960c91b608482015260a401610a80565b60005b818110156110a75760006040518060400160405280866001600160a01b03168152602001858585818110612c6957612c696159af565b9050602002016020810190612c7e9190615b10565b63ffffffff1690526001600160a01b0386166000908152609b60205260408120919250858585818110612cb357612cb36159af565b9050602002016020810190612cc89190615b10565b63ffffffff16815260208101919091526040016000205460ff16612d5f5760405162461bcd60e51b815260206004820152604260248201527f4156534469726563746f72792e5f72656769737465724f70657261746f72546f60448201527f4f70657261746f72536574733a20696e76616c6964206f70657261746f722073606482015261195d60f21b608482015260a401610a80565b612d6986826121ae565b15612e025760405162461bcd60e51b815260206004820152605960248201527f4156534469726563746f72792e5f72656769737465724f70657261746f72546f60448201527f4f70657261746f72536574733a206f70657261746f7220616c7265616479207260648201527f65676973746572656420746f206f70657261746f722073657400000000000000608482015260a401610a80565b6001600160a01b0385166000908152609c6020526040812090858585818110612e2d57612e2d6159af565b9050602002016020810190612e429190615b10565b63ffffffff1663ffffffff16815260200190815260200160002060008154612e69906159c5565b90915550612e97612e798261438d565b6001600160a01b0388166000908152609d6020526040902090614458565b506001600160a01b038086166000908152609e60209081526040808320938a16835292905290812081868686818110612ed257612ed26159af565b9050602002016020810190612ee79190615b10565b63ffffffff1681526020810191909152604001600020805490915060ff1615612f9e5760405162461bcd60e51b815260206004820152605a60248201527f4156534469726563746f72792e5f72656769737465724f70657261746f72546f60448201527f4f70657261746f72536574733a206f70657261746f7220616c7265616479207260648201527f65676973746572656420666f72206f70657261746f7220736574000000000000608482015260a401610a80565b805460ff191660011781556040516001600160a01b038816907f43232edf9071753d2321e5fa7e018363ee248e5f2142e6c08edd3265bfb4895e90612fe490859061570c565b60405180910390a2505080612ff8906159c5565b9050612c33565b81546000908161301185858385614464565b9050801561304c5761303685613028600184615980565b600091825260209091200190565b54600160201b90046001600160e01b0316611900565b600095945050505050565b42816040013510156130dd5760405162461bcd60e51b815260206004820152604360248201527f4156534469726563746f72792e5f766572696679416c6c6f6361746f7253696760448201527f6e61747572653a20616c6c6f6361746f72207369676e617475726520657870696064820152621c995960ea1b608482015260a401610a80565b6001600160a01b038516600090815260a56020908152604080832084830135845290915290205460ff161561316f5760405162461bcd60e51b815260206004820152603260248201527f4156534469726563746f72792e5f766572696679416c6c6f6361746f725369676044820152711b985d1d5c994e881cd85b1d081cdc195b9d60721b6064820152608401610a80565b6000613186858585856020013586604001356117b9565b90506131d186826131978580615fd4565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250612a7692505050565b506001600160a01b03909416600090815260a56020908152604080832096820135835295905293909320805460ff19166001179055505050565b6001600160a01b03808416600081815260a3602090815260408083209487168084529482528083205493835260a2825280832094835293905291822054909160ff84161561325a57600061325d565b60ff5b90505b818310801561327457508360ff168160ff16105b1561337f576001600160a01b03808716600090815260a26020908152604080832093891683529290529081208054859081106132b2576132b26159af565b6000918252602091829020604080518082019091529101546001600160401b038116825263ffffffff600160401b9091048116928201839052909250421611156132fc575061337f565b80516001600160a01b03808916600090815260a160209081526040808320938b168352929052908120805490919061333e9084906001600160401b0316615e8f565b92506101000a8154816001600160401b0302191690836001600160401b031602179055508361336c906159c5565b93506133778261601a565b915050613260565b50506001600160a01b03938416600090815260a36020908152604080832095909616825293909352929091209190915550565b60006133c1602084018461507c565b90506133d06040840184615ee8565b90506133df602085018561603a565b9050146134665760405162461bcd60e51b815260206004820152604960248201527f4156534469726563746f72792e5f6465616c6c6f636174653a206f706572617460448201527f6f725365747320616e64206d61676e69747564654469666673206c656e677468606482015268040dad2e6dac2e8c6d60bb1b608482015260a401610a80565b366000613476602086018661603a565b9150915060005b818110156112fb576001600160a01b03808816600090815260a060209081526040808320938816835292905290812061354f904290838787878181106134c5576134c56159af565b6134db926020604090920201908101915061507c565b6001600160a01b03166001600160a01b03168152602001908152602001600020600087878781811061350f5761350f6159af565b90506040020160200160208101906135279190615b10565b63ffffffff1663ffffffff1681526020019081526020016000206141f890919063ffffffff16565b905060006135606040890189615ee8565b84818110613570576135706159af565b90506020020160208101906135859190616083565b6001600160401b0316116136015760405162461bcd60e51b815260206004820152603e60248201527f4156534469726563746f72792e5f6465616c6c6f636174653a206d61676e697460448201527f75646544696666206d7573742062652067726561746572207468616e203000006064820152608401610a80565b6001600160401b0381166136186040890189615ee8565b84818110613628576136286159af565b905060200201602081019061363d9190616083565b6001600160401b031611156136ca5760405162461bcd60e51b815260206004820152604760248201527f4156534469726563746f72792e5f6465616c6c6f636174653a2063616e6e6f7460448201527f206465616c6c6f63617465206d6f7265207468616e207768617420697320616c6064820152661b1bd8d85d195960ca1b608482015260a401610a80565b6137d2426136db60408a018a615ee8565b858181106136eb576136eb6159af565b90506020020160208101906137009190616083565b6001600160a01b03808c16600090815260a060209081526040808320938c1683529290529081206001600160401b03929092169190888888818110613747576137476159af565b61375d926020604090920201908101915061507c565b6001600160a01b03166001600160a01b031681526020019081526020016000206000888888818110613791576137916159af565b90506040020160200160208101906137a99190615b10565b63ffffffff1663ffffffff16815260200190815260200160002061429d9092919063ffffffff16565b6137f488868686868181106137e9576137e96159af565b9050604002016144ba565b6001600160a01b03888116600090815260a260209081526040808320938916835292905281902080548251808401845290928190613834908c018c615ee8565b87818110613844576138446159af565b90506020020160208101906138599190616083565b6001600160401b03908116825263ffffffff808c166020938401528454600181018655600095865283862085519101805495850151909216600160401b026bffffffffffffffffffffffff19909516921691909117929092179091556001600160a01b03808c16835260a482526040808420918a16845291528120908686868181106138e7576138e76159af565b6138fd926020604090920201908101915061507c565b6001600160a01b03166001600160a01b031681526020019081526020016000206000868686818110613931576139316159af565b90506040020160200160208101906139499190615b10565b63ffffffff1681526020808201929092526040016000908120805460018101825590825291902001555061397c816159c5565b905061347d565b60005b818110156110a75760006040518060400160405280876001600160a01b031681526020018585858181106139bc576139bc6159af565b90506020020160208101906139d19190615b10565b63ffffffff16905290506139e585826121ae565b613a7d5760405162461bcd60e51b815260206004820152605960248201527f4156534469726563746f72792e5f646572656769737465724f70657261746f7260448201527f46726f6d4f70657261746f725365743a206f70657261746f72206e6f7420726560648201527f676973746572656420666f72206f70657261746f722073657400000000000000608482015260a401610a80565b6001600160a01b0386166000908152609c6020526040812090858585818110613aa857613aa86159af565b9050602002016020810190613abd9190615b10565b63ffffffff1663ffffffff16815260200190815260200160002060008154613ae490615eb1565b90915550613b12613af48261438d565b6001600160a01b0387166000908152609d6020526040902090614727565b50846001600160a01b03167fad34c3070be1dffbcaa499d000ba2b8d9848aefcac3059df245dd95c4ece14fe82604051613b4c919061570c565b60405180910390a250613b5e816159c5565b9050613986565b6000613b74602084018461507c565b9050366000613b86602086018661603a565b6001600160a01b03808916600090815260a1602090815260408083209389168352929052908120549294509092506001600160401b03909116905b82811015614107576000613bd86040890189615ee8565b83818110613be857613be86159af565b9050602002016020810190613bfd9190616083565b6001600160401b031611613c855760405162461bcd60e51b815260206004820152604360248201527f4156534469726563746f72792e7175657565416c6c6f636174696f6e733a206d60448201527f61676e697475646544696666206d75737420626520677265617465722074686160648201526206e20360ec1b608482015260a401610a80565b613c926040880188615ee8565b82818110613ca257613ca26159af565b9050602002016020810190613cb79190616083565b6001600160401b0316826001600160401b03161015613d575760405162461bcd60e51b815260206004820152605060248201527f4156534469726563746f72792e7175657565416c6c6f636174696f6e733a206960448201527f6e73756666696369656e7420617661696c61626c652066726565206d61676e6960648201526f7475646520746f20616c6c6f6361746560801b608482015260a401610a80565b6001600160a01b03808916600090815260a06020908152604080832093891683529290529081208190613e2390429083898988818110613d9957613d996159af565b613daf926020604090920201908101915061507c565b6001600160a01b03166001600160a01b031681526020019081526020016000206000898988818110613de357613de36159af565b9050604002016020016020810190613dfb9190615b10565b63ffffffff1663ffffffff16815260200190815260200160002061473390919063ffffffff16565b91509150816001600160e01b03166000141580613e3f57508015155b15613fa5576001600160a01b03808b16600090815260a060209081526040808320938b168352929052908120613f0191888887818110613e8157613e816159af565b613e97926020604090920201908101915061507c565b6001600160a01b03166001600160a01b031681526020019081526020016000206000888887818110613ecb57613ecb6159af565b9050604002016020016020810190613ee39190615b10565b63ffffffff1663ffffffff1681526020019081526020016000205490565b613f0c600183615997565b14613fa55760405162461bcd60e51b815260206004820152605d60248201527f4156534469726563746f72792e7175657565416c6c6f636174696f6e733a206560448201527f7863656564206d61782070656e64696e6720616c6c6f636174696f6e7320616c60648201527f6c6f77656420666f72206f702c206f705365742c207374726174656779000000608482015260a401610a80565b6140b388613fb660408c018c615ee8565b86818110613fc657613fc66159af565b9050602002016020810190613fdb9190616083565b613fee906001600160401b03168561609e565b6001600160a01b03808e16600090815260a060209081526040808320938e168352929052908120908a8a89818110614028576140286159af565b61403e926020604090920201908101915061507c565b6001600160a01b03166001600160a01b0316815260200190815260200160002060008a8a89818110614072576140726159af565b905060400201602001602081019061408a9190615b10565b63ffffffff1663ffffffff1681526020019081526020016000206143729092919063ffffffff16565b506140c3905060408a018a615ee8565b848181106140d3576140d36159af565b90506020020160208101906140e89190616083565b6140f29085615e67565b9350505080614100906159c5565b9050613bc1565b506001600160a01b03968716600090815260a1602090815260408083209690991682529490945295909220805467ffffffffffffffff19166001600160401b039096169590951790945550505050565b6033546001600160a01b031633146117455760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610a80565b60006141bb61295b565b60405161190160f01b6020820152602281019190915260428101839052606201604051602081830303815290604052805190602001209050919050565b815460009081816005811115614255576000614213846147ec565b61421d9085615980565b60008881526020902090915081015463ffffffff908116908716101561424557809150614253565b614250816001615997565b92505b505b600061426387878585614464565b905080156142905761427a87613028600184615980565b54600160201b90046001600160e01b03166111ab565b6000979650505050505050565b6000806142aa8585614733565b91509150816001600160e01b031660001480156142c5575080155b156142cf57506000195b84545b80821015611546576000868152602090208201805485908290600490614309908490600160201b90046001600160e01b0316615ec8565b92506101000a8154816001600160e01b0302191690836001600160e01b03160217905550826001019250506142d2565b805460009080156143695761435383613028600184615980565b54600160201b90046001600160e01b0316611858565b60009392505050565b6000806143808585856148d1565b915091505b935093915050565b60008160000151826020015163ffffffff166040516020016143da92919060609290921b6bffffffffffffffffffffffff1916825260a01b6001600160a01b031916601482015260200190565b604051602081830303815290604052610b1f906160c0565b60008181526001830160205260408120541515611858565b6000826000018281548110614421576144216159af565b9060005260206000200154905092915050565b60008060006144438585614a77565b9150915061445081614ae4565b509392505050565b60006118588383614c9f565b60005b8183101561445057600061447b8484614cee565b60008781526020902090915063ffffffff86169082015463ffffffff1611156144a6578092506144b4565b6144b1816001615997565b93505b50614467565b6001600160a01b03808416600090815260a4602090815260408083209386168352928152918120909182906144f19085018561507c565b6001600160a01b03166001600160a01b0316815260200190815260200160002060008360200160208101906145269190615b10565b63ffffffff1681526020810191909152604001600020549050805b80156110a7576001600160a01b03808616600090815260a46020908152604080832093881683529281529181209091829061457e9087018761507c565b6001600160a01b03166001600160a01b0316815260200190815260200160002060008560200160208101906145b39190615b10565b63ffffffff16815260208101919091526040016000206145d4600184615980565b815481106145e4576145e46159af565b60009182526020808320909101546001600160a01b03808a16845260a283526040808520918a168552925290822080549193509083908110614628576146286159af565b6000918252602091829020604080518082019091529101546001600160401b038116825263ffffffff600160401b90910481169282018390529092504216111561470d5760016146788486615980565b614683906001615997565b106147085760405162461bcd60e51b815260206004820152604960248201527f4156534469726563746f72792e5f636865636b5175657565644465616c6c6f6360448201527f6174696f6e733a2065786365656473206d61782070656e64696e67206465616c6064820152686c6f636174696f6e7360b81b608482015260a401610a80565b614714565b50506110a7565b50508061472090615eb1565b9050614541565b60006118588383614d09565b8154600090819081816005811115614792576000614750846147ec565b61475a9085615980565b60008981526020902090915081015463ffffffff908116908816101561478257809150614790565b61478d816001615997565b92505b505b60006147a088888585614464565b905080156147d8576147b788613028600184615980565b54600160201b90046001600160e01b03166147d3600183615980565b6147dc565b6000805b95509550505050505b9250929050565b6000816147fb57506000919050565b6000600161480884614dfc565b901c6001901b9050600181848161482157614821615b4a565b048201901c9050600181848161483957614839615b4a565b048201901c9050600181848161485157614851615b4a565b048201901c9050600181848161486957614869615b4a565b048201901c9050600181848161488157614881615b4a565b048201901c9050600181848161489957614899615b4a565b048201901c905060018184816148b1576148b1615b4a565b048201901c9050611858818285816148cb576148cb615b4a565b04614e90565b825460009081908015614a1e5760006148ef87613028600185615980565b60408051808201909152905463ffffffff808216808452600160201b9092046001600160e01b0316602084015291925090871610156149705760405162461bcd60e51b815260206004820152601b60248201527f436865636b706f696e743a2064656372656173696e67206b65797300000000006044820152606401610a80565b8563ffffffff16816000015163ffffffff1614156149bf578461499888613028600186615980565b80546001600160e01b0392909216600160201b0263ffffffff909216919091179055614a0e565b6040805180820190915263ffffffff80881682526001600160e01b0380881660208085019182528b54600181018d5560008d81529190912094519151909216600160201b029216919091179101555b6020015192508391506143859050565b50506040805180820190915263ffffffff80851682526001600160e01b0380851660208085019182528854600181018a5560008a815291822095519251909316600160201b029190931617920191909155905081614385565b600080825160411415614aae5760208301516040840151606085015160001a614aa287828585614ea6565b945094505050506147e5565b825160401415614ad85760208301516040840151614acd868383614f93565b9350935050506147e5565b506000905060026147e5565b6000816004811115614af857614af861558e565b1415614b015750565b6001816004811115614b1557614b1561558e565b1415614b635760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610a80565b6002816004811115614b7757614b7761558e565b1415614bc55760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610a80565b6003816004811115614bd957614bd961558e565b1415614c325760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b6064820152608401610a80565b6004816004811115614c4657614c4661558e565b1415610a925760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b6064820152608401610a80565b6000818152600183016020526040812054614ce657508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610b1f565b506000610b1f565b6000614cfd6002848418615b60565b61185890848416615997565b60008181526001830160205260408120548015614df2576000614d2d600183615980565b8554909150600090614d4190600190615980565b9050818114614da6576000866000018281548110614d6157614d616159af565b9060005260206000200154905080876000018481548110614d8457614d846159af565b6000918252602080832090910192909255918252600188019052604090208390555b8554869080614db757614db76160e7565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050610b1f565b6000915050610b1f565b600080608083901c15614e1157608092831c92015b604083901c15614e2357604092831c92015b602083901c15614e3557602092831c92015b601083901c15614e4757601092831c92015b600883901c15614e5957600892831c92015b600483901c15614e6b57600492831c92015b600283901c15614e7d57600292831c92015b600183901c15610b1f5760010192915050565b6000818310614e9f5781611858565b5090919050565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115614edd5750600090506003614f8a565b8460ff16601b14158015614ef557508460ff16601c14155b15614f065750600090506004614f8a565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015614f5a573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116614f8357600060019250925050614f8a565b9150600090505b94509492505050565b6000806001600160ff1b03831681614fb060ff86901c601b615997565b9050614fbe87828885614ea6565b935093505050935093915050565b6001600160a01b0381168114610a9257600080fd5b600080600060608486031215614ff657600080fd5b833561500181614fcc565b95602085013595506040909401359392505050565b602080825282518282018190526000919060409081850190868401855b8281101561506f5761505f84835180516001600160a01b0316825260209081015163ffffffff16910152565b9284019290850190600101615033565b5091979650505050505050565b60006020828403121561508e57600080fd5b813561185881614fcc565b634e487b7160e01b600052604160045260246000fd5b604051606081016001600160401b03811182821017156150d1576150d1615099565b60405290565b604051601f8201601f191681016001600160401b03811182821017156150ff576150ff615099565b604052919050565b803563ffffffff8116811461511b57600080fd5b919050565b600080828403606081121561513457600080fd5b833561513f81614fcc565b92506040601f198201121561515357600080fd5b50604051604081018181106001600160401b038211171561517657615176615099565b604052602084013561518781614fcc565b815261519560408501615107565b6020820152809150509250929050565b6000602082840312156151b757600080fd5b5035919050565b6000806000606084860312156151d357600080fd5b83356151de81614fcc565b925060208401356151ee81614fcc565b929592945050506040919091013590565b60008083601f84011261521157600080fd5b5081356001600160401b0381111561522857600080fd5b6020830191508360208260051b85010111156147e557600080fd5b60006060828403121561525557600080fd5b61525d6150af565b905081356001600160401b038082111561527657600080fd5b818401915084601f83011261528a57600080fd5b813560208282111561529e5761529e615099565b6152b0601f8301601f191682016150d7565b925081835286818386010111156152c657600080fd5b818185018285013760008183850101528285528086013581860152505050506040820135604082015292915050565b6000806000806060858703121561530b57600080fd5b843561531681614fcc565b935060208501356001600160401b038082111561533257600080fd5b61533e888389016151ff565b9095509350604087013591508082111561535757600080fd5b5061536487828801615243565b91505092959194509250565b60008060006060848603121561538557600080fd5b833561539081614fcc565b925060208401356153a081614fcc565b91506153ae60408501615107565b90509250925092565b60008060008084860360a08112156153ce57600080fd5b85356153d981614fcc565b94506040601f19820112156153ed57600080fd5b50602085019250606085013561540281614fcc565b915061541060808601615107565b905092959194509250565b6000806000806060858703121561543157600080fd5b843561543c81614fcc565b935060208501356001600160401b038082111561545857600080fd5b615464888389016151ff565b9095509350604087013591508082111561547d57600080fd5b5085016060818803121561549057600080fd5b939692955090935050565b600080604083850312156154ae57600080fd5b82356154b981614fcc565b946020939093013593505050565b6000806000806000608086880312156154df57600080fd5b85356154ea81614fcc565b945060208601356154fa81614fcc565b935060408601356001600160401b038082111561551657600080fd5b61552289838a016151ff565b9095509350606088013591508082111561553b57600080fd5b5061554888828901615243565b9150509295509295909350565b6000806040838503121561556857600080fd5b823561557381614fcc565b9150602083013561558381614fcc565b809150509250929050565b634e487b7160e01b600052602160045260246000fd5b60208101600283106155c657634e487b7160e01b600052602160045260246000fd5b91905290565b6000602082840312156155de57600080fd5b813560ff8116811461185857600080fd5b60008060008060006060868803121561560757600080fd5b853561561281614fcc565b945060208601356001600160401b038082111561562e57600080fd5b61563a89838a016151ff565b9096509450604088013591508082111561565357600080fd5b50615660888289016151ff565b969995985093965092949392505050565b6000806040838503121561568457600080fd5b823561568f81614fcc565b915061569d60208401615107565b90509250929050565b6000806000806000608086880312156156be57600080fd5b85356156c981614fcc565b945060208601356001600160401b038111156156e457600080fd5b6156f0888289016151ff565b9699909850959660408101359660609091013595509350505050565b81516001600160a01b0316815260208083015163ffffffff169082015260408101610b1f565b6000806000806080858703121561574857600080fd5b843561575381614fcc565b9350602085013561576381614fcc565b93969395505050506040820135916060013590565b6000806020838503121561578b57600080fd5b82356001600160401b03808211156157a257600080fd5b818501915085601f8301126157b657600080fd5b8135818111156157c557600080fd5b8660208285010111156157d757600080fd5b60209290920196919550909350505050565b600080602083850312156157fc57600080fd5b82356001600160401b0381111561581257600080fd5b61581e858286016151ff565b90969095509350505050565b60008060008060006080868803121561584257600080fd5b853561584d81614fcc565b945061585b60208701615107565b935060408601356001600160401b0381111561587657600080fd5b615882888289016151ff565b909450925050606086013561ffff8116811461589d57600080fd5b809150509295509295909350565b6000806000604084860312156158c057600080fd5b83356158cb81614fcc565b925060208401356001600160401b038111156158e657600080fd5b6158f2868287016151ff565b9497909650939450505050565b6000806000806040858703121561591557600080fd5b84356001600160401b038082111561592c57600080fd5b615938888389016151ff565b9096509450602087013591508082111561595157600080fd5b5061595e878288016151ff565b95989497509550505050565b634e487b7160e01b600052601160045260246000fd5b6000828210156159925761599261596a565b500390565b600082198211156159aa576159aa61596a565b500190565b634e487b7160e01b600052603260045260246000fd5b60006000198214156159d9576159d961596a565b5060010190565b6000602082840312156159f257600080fd5b815161185881614fcc565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b600063ffffffff808316818516808303821115615a6657615a6661596a565b01949350505050565b600060208284031215615a8157600080fd5b8151801515811461185857600080fd5b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b60208082526019908201527f5061757361626c653a20696e6465782069732070617573656400000000000000604082015260600190565b600060208284031215615b2257600080fd5b61185882615107565b6000816000190483118215151615615b4557615b4561596a565b500290565b634e487b7160e01b600052601260045260246000fd5b600082615b7d57634e487b7160e01b600052601260045260246000fd5b500490565b60008235605e19833603018112615b9857600080fd5b9190910192915050565b8183526000602080850194508260005b85811015615bf8578135615bc581614fcc565b6001600160a01b0316875263ffffffff615be0838501615107565b16878401526040968701969190910190600101615bb2565b509495945050505050565b6000808335601e19843603018112615c1a57600080fd5b83016020810192503590506001600160401b03811115615c3957600080fd5b8060051b36038313156147e557600080fd5b80356001600160401b038116811461511b57600080fd5b8183526000602080850194508260005b85811015615bf8576001600160401b03615c8b83615c4b565b1687529582019590820190600101615c72565b600060a08201888352602060018060a01b03808a1682860152604060a0818701528389855260c08701905060c08a60051b88010194508a60005b8b811015615d9f5788870360bf190183528135368e9003605e19018112615cfe57600080fd5b8d0160608135615d0d81614fcc565b871689528188013536839003601e19018112615d2857600080fd5b820180356001600160401b03811115615d4057600080fd5b8060061b3603841315615d5257600080fd5b828a8c0152615d66838c01828c8501615ba2565b92505050615d7686830183615c03565b9250898203878b0152615d8a828483615c62565b99505050928601925090850190600101615cd8565b505050505060608401959095526080909201929092529695505050505050565b8183526000602080850194508260005b85811015615bf85763ffffffff615de583615107565b1687529582019590820190600101615dcf565b8681526001600160a01b038616602082015260a060408201819052600090615e239083018688615dbf565b60608301949094525060800152949350505050565b60208152816020820152818360408301376000818301604090810191909152601f909201601f19160101919050565b60006001600160401b0383811690831681811015615e8757615e8761596a565b039392505050565b60006001600160401b03808316818516808303821115615a6657615a6661596a565b600081615ec057615ec061596a565b506000190190565b60006001600160e01b0383811690831681811015615e8757615e8761596a565b6000808335601e19843603018112615eff57600080fd5b8301803591506001600160401b03821115615f1957600080fd5b6020019150600581901b36038213156147e557600080fd5b602081526000615f45602083018486615dbf565b949350505050565b82815260006020604081840152835180604085015260005b81811015615f8157858101830151858201606001528201615f65565b81811115615f93576000606083870101525b50601f01601f191692909201606001949350505050565b600060208284031215615fbc57600080fd5b81516001600160e01b03198116811461185857600080fd5b6000808335601e19843603018112615feb57600080fd5b8301803591506001600160401b0382111561600557600080fd5b6020019150368190038213156147e557600080fd5b600060ff821660ff8114156160315761603161596a565b60010192915050565b6000808335601e1984360301811261605157600080fd5b8301803591506001600160401b0382111561606b57600080fd5b6020019150600681901b36038213156147e557600080fd5b60006020828403121561609557600080fd5b61185882615c4b565b60006001600160e01b03828116848216808303821115615a6657615a6661596a565b805160208083015191908110156160e1576000198160200360031b1b821691505b50919050565b634e487b7160e01b600052603160045260246000fdfe4156534469726563746f72792e72656769737465724f70657261746f72546f4fa264697066735822122057615600b1b8ed2c594c0ca166b1ad1650f361bd0a718ba6d96322a3c846a50d64736f6c634300080c0033",
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

// MAGNITUDEADJUSTMENTTYPEHASH is a free data retrieval call binding the contract method 0x7b205de3.
//
// Solidity: function MAGNITUDE_ADJUSTMENT_TYPEHASH() view returns(bytes32)
func (_AVSDirectory *AVSDirectoryCaller) MAGNITUDEADJUSTMENTTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AVSDirectory.contract.Call(opts, &out, "MAGNITUDE_ADJUSTMENT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MAGNITUDEADJUSTMENTTYPEHASH is a free data retrieval call binding the contract method 0x7b205de3.
//
// Solidity: function MAGNITUDE_ADJUSTMENT_TYPEHASH() view returns(bytes32)
func (_AVSDirectory *AVSDirectorySession) MAGNITUDEADJUSTMENTTYPEHASH() ([32]byte, error) {
	return _AVSDirectory.Contract.MAGNITUDEADJUSTMENTTYPEHASH(&_AVSDirectory.CallOpts)
}

// MAGNITUDEADJUSTMENTTYPEHASH is a free data retrieval call binding the contract method 0x7b205de3.
//
// Solidity: function MAGNITUDE_ADJUSTMENT_TYPEHASH() view returns(bytes32)
func (_AVSDirectory *AVSDirectoryCallerSession) MAGNITUDEADJUSTMENTTYPEHASH() ([32]byte, error) {
	return _AVSDirectory.Contract.MAGNITUDEADJUSTMENTTYPEHASH(&_AVSDirectory.CallOpts)
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

// AllocatorSaltIsSpent is a free data retrieval call binding the contract method 0x46fb3a00.
//
// Solidity: function allocatorSaltIsSpent(address , bytes32 ) view returns(bool)
func (_AVSDirectory *AVSDirectoryCaller) AllocatorSaltIsSpent(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (bool, error) {
	var out []interface{}
	err := _AVSDirectory.contract.Call(opts, &out, "allocatorSaltIsSpent", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllocatorSaltIsSpent is a free data retrieval call binding the contract method 0x46fb3a00.
//
// Solidity: function allocatorSaltIsSpent(address , bytes32 ) view returns(bool)
func (_AVSDirectory *AVSDirectorySession) AllocatorSaltIsSpent(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _AVSDirectory.Contract.AllocatorSaltIsSpent(&_AVSDirectory.CallOpts, arg0, arg1)
}

// AllocatorSaltIsSpent is a free data retrieval call binding the contract method 0x46fb3a00.
//
// Solidity: function allocatorSaltIsSpent(address , bytes32 ) view returns(bool)
func (_AVSDirectory *AVSDirectoryCallerSession) AllocatorSaltIsSpent(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _AVSDirectory.Contract.AllocatorSaltIsSpent(&_AVSDirectory.CallOpts, arg0, arg1)
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

// CalculateMagnitudeAdjustmentDigestHash is a free data retrieval call binding the contract method 0x8522fcb1.
//
// Solidity: function calculateMagnitudeAdjustmentDigestHash(address operator, (address,(address,uint32)[],uint64[])[] adjustments, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_AVSDirectory *AVSDirectoryCaller) CalculateMagnitudeAdjustmentDigestHash(opts *bind.CallOpts, operator common.Address, adjustments []IAVSDirectoryMagnitudeAdjustment, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _AVSDirectory.contract.Call(opts, &out, "calculateMagnitudeAdjustmentDigestHash", operator, adjustments, salt, expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateMagnitudeAdjustmentDigestHash is a free data retrieval call binding the contract method 0x8522fcb1.
//
// Solidity: function calculateMagnitudeAdjustmentDigestHash(address operator, (address,(address,uint32)[],uint64[])[] adjustments, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_AVSDirectory *AVSDirectorySession) CalculateMagnitudeAdjustmentDigestHash(operator common.Address, adjustments []IAVSDirectoryMagnitudeAdjustment, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _AVSDirectory.Contract.CalculateMagnitudeAdjustmentDigestHash(&_AVSDirectory.CallOpts, operator, adjustments, salt, expiry)
}

// CalculateMagnitudeAdjustmentDigestHash is a free data retrieval call binding the contract method 0x8522fcb1.
//
// Solidity: function calculateMagnitudeAdjustmentDigestHash(address operator, (address,(address,uint32)[],uint64[])[] adjustments, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_AVSDirectory *AVSDirectoryCallerSession) CalculateMagnitudeAdjustmentDigestHash(operator common.Address, adjustments []IAVSDirectoryMagnitudeAdjustment, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _AVSDirectory.Contract.CalculateMagnitudeAdjustmentDigestHash(&_AVSDirectory.CallOpts, operator, adjustments, salt, expiry)
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
// Solidity: function allocate(address operator, (address,(address,uint32)[],uint64[])[] allocations, (bytes,bytes32,uint256) allocatorSignature) returns()
func (_AVSDirectory *AVSDirectoryTransactor) Allocate(opts *bind.TransactOpts, operator common.Address, allocations []IAVSDirectoryMagnitudeAdjustment, allocatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _AVSDirectory.contract.Transact(opts, "allocate", operator, allocations, allocatorSignature)
}

// Allocate is a paid mutator transaction binding the contract method 0x70196708.
//
// Solidity: function allocate(address operator, (address,(address,uint32)[],uint64[])[] allocations, (bytes,bytes32,uint256) allocatorSignature) returns()
func (_AVSDirectory *AVSDirectorySession) Allocate(operator common.Address, allocations []IAVSDirectoryMagnitudeAdjustment, allocatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _AVSDirectory.Contract.Allocate(&_AVSDirectory.TransactOpts, operator, allocations, allocatorSignature)
}

// Allocate is a paid mutator transaction binding the contract method 0x70196708.
//
// Solidity: function allocate(address operator, (address,(address,uint32)[],uint64[])[] allocations, (bytes,bytes32,uint256) allocatorSignature) returns()
func (_AVSDirectory *AVSDirectoryTransactorSession) Allocate(operator common.Address, allocations []IAVSDirectoryMagnitudeAdjustment, allocatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _AVSDirectory.Contract.Allocate(&_AVSDirectory.TransactOpts, operator, allocations, allocatorSignature)
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
// Solidity: function deallocate(address operator, (address,(address,uint32)[],uint64[])[] deallocations, (bytes,bytes32,uint256) allocatorSignature) returns()
func (_AVSDirectory *AVSDirectoryTransactor) Deallocate(opts *bind.TransactOpts, operator common.Address, deallocations []IAVSDirectoryMagnitudeAdjustment, allocatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _AVSDirectory.contract.Transact(opts, "deallocate", operator, deallocations, allocatorSignature)
}

// Deallocate is a paid mutator transaction binding the contract method 0x3367a33c.
//
// Solidity: function deallocate(address operator, (address,(address,uint32)[],uint64[])[] deallocations, (bytes,bytes32,uint256) allocatorSignature) returns()
func (_AVSDirectory *AVSDirectorySession) Deallocate(operator common.Address, deallocations []IAVSDirectoryMagnitudeAdjustment, allocatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _AVSDirectory.Contract.Deallocate(&_AVSDirectory.TransactOpts, operator, deallocations, allocatorSignature)
}

// Deallocate is a paid mutator transaction binding the contract method 0x3367a33c.
//
// Solidity: function deallocate(address operator, (address,(address,uint32)[],uint64[])[] deallocations, (bytes,bytes32,uint256) allocatorSignature) returns()
func (_AVSDirectory *AVSDirectoryTransactorSession) Deallocate(operator common.Address, deallocations []IAVSDirectoryMagnitudeAdjustment, allocatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _AVSDirectory.Contract.Deallocate(&_AVSDirectory.TransactOpts, operator, deallocations, allocatorSignature)
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

// UpdateFreeMagnitude is a paid mutator transaction binding the contract method 0x84b78fc2.
//
// Solidity: function updateFreeMagnitude(address operator, address[] strategies, uint8[] numToComplete) returns()
func (_AVSDirectory *AVSDirectoryTransactor) UpdateFreeMagnitude(opts *bind.TransactOpts, operator common.Address, strategies []common.Address, numToComplete []uint8) (*types.Transaction, error) {
	return _AVSDirectory.contract.Transact(opts, "updateFreeMagnitude", operator, strategies, numToComplete)
}

// UpdateFreeMagnitude is a paid mutator transaction binding the contract method 0x84b78fc2.
//
// Solidity: function updateFreeMagnitude(address operator, address[] strategies, uint8[] numToComplete) returns()
func (_AVSDirectory *AVSDirectorySession) UpdateFreeMagnitude(operator common.Address, strategies []common.Address, numToComplete []uint8) (*types.Transaction, error) {
	return _AVSDirectory.Contract.UpdateFreeMagnitude(&_AVSDirectory.TransactOpts, operator, strategies, numToComplete)
}

// UpdateFreeMagnitude is a paid mutator transaction binding the contract method 0x84b78fc2.
//
// Solidity: function updateFreeMagnitude(address operator, address[] strategies, uint8[] numToComplete) returns()
func (_AVSDirectory *AVSDirectoryTransactorSession) UpdateFreeMagnitude(operator common.Address, strategies []common.Address, numToComplete []uint8) (*types.Transaction, error) {
	return _AVSDirectory.Contract.UpdateFreeMagnitude(&_AVSDirectory.TransactOpts, operator, strategies, numToComplete)
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
