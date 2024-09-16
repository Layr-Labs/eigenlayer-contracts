// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package AllocationManager

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

// IAllocationManagerTypesAllocateParams is an auto generated low-level Go binding around an user-defined struct.
type IAllocationManagerTypesAllocateParams struct {
	OperatorSet   OperatorSet
	Strategies    []common.Address
	NewMagnitudes []uint64
}

// IAllocationManagerTypesAllocation is an auto generated low-level Go binding around an user-defined struct.
type IAllocationManagerTypesAllocation struct {
	CurrentMagnitude uint64
	PendingDiff      *big.Int
	EffectBlock      uint32
}

// IAllocationManagerTypesCreateSetParams is an auto generated low-level Go binding around an user-defined struct.
type IAllocationManagerTypesCreateSetParams struct {
	OperatorSetId uint32
	Strategies    []common.Address
}

// IAllocationManagerTypesDeregisterParams is an auto generated low-level Go binding around an user-defined struct.
type IAllocationManagerTypesDeregisterParams struct {
	Operator       common.Address
	Avs            common.Address
	OperatorSetIds []uint32
}

// IAllocationManagerTypesRegisterParams is an auto generated low-level Go binding around an user-defined struct.
type IAllocationManagerTypesRegisterParams struct {
	Avs            common.Address
	OperatorSetIds []uint32
	Data           []byte
}

// IAllocationManagerTypesSlashingParams is an auto generated low-level Go binding around an user-defined struct.
type IAllocationManagerTypesSlashingParams struct {
	Operator      common.Address
	OperatorSetId uint32
	Strategies    []common.Address
	WadsToSlash   []*big.Int
	Description   string
}

// OperatorSet is an auto generated low-level Go binding around an user-defined struct.
type OperatorSet struct {
	Avs common.Address
	Id  uint32
}

// AllocationManagerMetaData contains all meta data concerning the AllocationManager contract.
var AllocationManagerMetaData = &bind.MetaData{
<<<<<<< HEAD
<<<<<<< HEAD
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_delegation\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"_permissionController\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"},{\"name\":\"_DEALLOCATION_DELAY\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_ALLOCATION_CONFIGURATION_DELAY\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_version\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ALLOCATION_CONFIGURATION_DELAY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEALLOCATION_DELAY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addStrategiesToOperatorSet\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"clearDeallocationQueue\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"numToClear\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createOperatorSets\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"params\",\"type\":\"tuple[]\",\"internalType\":\"structIAllocationManagerTypes.CreateSetParams[]\",\"components\":[{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregisterFromOperatorSets\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIAllocationManagerTypes.DeregisterParams\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAVSRegistrar\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAVSRegistrar\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocatableMagnitude\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocatedSets\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocatedStake\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[][]\",\"internalType\":\"uint256[][]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocatedStrategies\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocation\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIAllocationManagerTypes.Allocation\",\"components\":[{\"name\":\"currentMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"pendingDiff\",\"type\":\"int128\",\"internalType\":\"int128\"},{\"name\":\"effectBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocationDelay\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocations\",\"inputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIAllocationManagerTypes.Allocation[]\",\"components\":[{\"name\":\"currentMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"pendingDiff\",\"type\":\"int128\",\"internalType\":\"int128\"},{\"name\":\"effectBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getEncumberedMagnitude\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxMagnitude\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxMagnitudes\",\"inputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxMagnitudes\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxMagnitudesAtBlock\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMemberCount\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMembers\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMinimumSlashableStake\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"futureBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"slashableStake\",\"type\":\"uint256[][]\",\"internalType\":\"uint256[][]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorSetCount\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRegisteredSets\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStrategiesInOperatorSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStrategyAllocations\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIAllocationManagerTypes.Allocation[]\",\"components\":[{\"name\":\"currentMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"pendingDiff\",\"type\":\"int128\",\"internalType\":\"int128\"},{\"name\":\"effectBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isMemberOfOperatorSet\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperatorSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperatorSlashable\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"modifyAllocations\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"params\",\"type\":\"tuple[]\",\"internalType\":\"structIAllocationManagerTypes.AllocateParams[]\",\"components\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"newMagnitudes\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"permissionController\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerForOperatorSets\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIAllocationManagerTypes.RegisterParams\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeStrategiesFromOperatorSet\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAVSRegistrar\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"registrar\",\"type\":\"address\",\"internalType\":\"contractIAVSRegistrar\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAllocationDelay\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delay\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slashOperator\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIAllocationManagerTypes.SlashingParams\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"wadsToSlash\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAVSMetadataURI\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AVSMetadataURIUpdated\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"metadataURI\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AVSRegistrarSet\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"registrar\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIAVSRegistrar\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllocationDelaySet\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"delay\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"effectBlock\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllocationUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"magnitude\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"effectBlock\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EncumberedMagnitudeUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"encumberedMagnitude\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MaxMagnitudeUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"maxMagnitude\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorAddedToOperatorSet\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRemovedFromOperatorSet\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSetCreated\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSlashed\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategies\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"contractIStrategy[]\"},{\"name\":\"wadSlashed\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"description\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyAddedToOperatorSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyRemovedFromOperatorSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyMemberOfSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CurrentlyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Empty\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputArrayLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientMagnitude\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAVSRegistrar\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCaller\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNewPausedStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperatorSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPermissions\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSnapshotOrdering\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidWadToSlash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ModificationAlreadyPending\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonexistentAVSMetadata\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotMemberOfSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyPauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnpauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorNotSlashable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OutOfBounds\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SameMagnitude\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategiesMustBeInAscendingOrder\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategyAlreadyInOperatorSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategyNotInOperatorSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"UninitializedAllocationDelay\",\"inputs\":[]}]",
	Bin: "0x610140604052348015610010575f5ffd5b50604051615fb7380380615fb783398101604081905261002f916101e6565b8084878585896001600160a01b03811661005c576040516339b190bb60e11b815260040160405180910390fd5b6001600160a01b0390811660805292831660a05263ffffffff91821660c0521660e052166101005261008d816100a5565b610120525061009a6100eb565b50505050505061034a565b5f5f829050601f815111156100d8578260405163305a27a960e01b81526004016100cf91906102ef565b60405180910390fd5b80516100e382610324565b179392505050565b5f54610100900460ff16156101525760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b60648201526084016100cf565b5f5460ff908116146101a1575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146101b7575f5ffd5b50565b805163ffffffff811681146101cd575f5ffd5b919050565b634e487b7160e01b5f52604160045260245ffd5b5f5f5f5f5f5f60c087890312156101fb575f5ffd5b8651610206816101a3565b6020880151909650610217816101a3565b6040880151909550610228816101a3565b9350610236606088016101ba565b9250610244608088016101ba565b60a08801519092506001600160401b0381111561025f575f5ffd5b8701601f8101891361026f575f5ffd5b80516001600160401b03811115610288576102886101d2565b604051601f8201601f19908116603f011681016001600160401b03811182821017156102b6576102b66101d2565b6040528181528282016020018b10156102cd575f5ffd5b8160208401602083015e5f602083830101528093505050509295509295509295565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b80516020808301519190811015610344575f198160200360031b1b821691505b50919050565b60805160a05160c05160e0516101005161012051615bd26103e55f395f611b3701525f818161044f015261356c01525f81816105b30152613dfc01525f818161036101528181611fce01526126d001525f8181610732015281816114d501528181611b6b01528181611bd5015281816129ce015261364b01525f81816105da0152818161085801528181611c7a01526131e30152615bd25ff3fe608060405234801561000f575f5ffd5b50600436106102b1575f3560e01c80636cfb44811161017b578063a9821821116100e4578063c221d8ae1161009e578063df5cf72311610079578063df5cf7231461072d578063f2fde38b14610754578063f605ce0814610767578063fabc1cbc1461077a575f5ffd5b8063c221d8ae146106f4578063cd6dc68714610707578063d3d96ff41461071a575f5ffd5b8063a982182114610666578063adc2e3d914610679578063b2447af71461068c578063b66bd9891461069f578063b9fbaed1146106b2578063ba1a84e5146106e1575f5ffd5b8063886f119511610135578063886f1195146105d55780638ce64854146105fc5780638da5cb5b1461061c57806394d7d00c1461062d578063952899ee14610640578063a9333ec814610653575f5ffd5b80636cfb4481146105425780636e3492b51461056d5780636e875dba14610580578063715018a61461059357806379ae50cd1461059b5780637bc1ef61146105ae575f5ffd5b80634177a87c1161021d57806354fd4d50116101d757806354fd4d50146104ca57806356c483e6146104df578063595c6a67146104f25780635ac86ab7146104fa5780635c975abb1461051d578063670d3ba21461052f575f5ffd5b80634177a87c1461042a5780634657e26a1461044a5780634a10ffe5146104715780634b5046ef1461049157806350feea20146104a4578063547afb87146104b7575f5ffd5b80632981eb771161026e5780632981eb771461035c5780632b453a9a146103985780632bab2c4a146103b8578063304c10cd146103cb57806336352057146103f657806340120dab14610409575f5ffd5b806310e1b9b8146102b55780631352c3e6146102de578063136439dd1461030157806315fe502814610316578063260dc75814610336578063261f84e014610349575b5f5ffd5b6102c86102c3366004614a50565b61078d565b6040516102d59190614a97565b60405180910390f35b6102f16102ec366004614aca565b6107c8565b60405190151581526020016102d5565b61031461030f366004614afe565b610843565b005b610329610324366004614b15565b610918565b6040516102d59190614b93565b6102f1610344366004614ba5565b610a2f565b610314610357366004614bff565b610a60565b6103837f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff90911681526020016102d5565b6103ab6103a6366004614ce4565b610d3b565b6040516102d59190614d87565b6103ab6103c6366004614dea565b610d51565b6103de6103d9366004614b15565b610df0565b6040516001600160a01b0390911681526020016102d5565b610314610404366004614e6e565b610e1f565b61041c610417366004614ec0565b611629565b6040516102d5929190614f4d565b61043d610438366004614ba5565b6117a4565b6040516102d59190614faa565b6103de7f000000000000000000000000000000000000000000000000000000000000000081565b61048461047f366004614fbc565b6117c8565b6040516102d59190614fff565b61031461049f36600461504a565b611870565b6103146104b23660046150ca565b61192a565b6104846104c5366004615128565b611a88565b6104d2611b30565b6040516102d5919061516a565b6103146104ed36600461519f565b611b60565b610314611c65565b6102f16105083660046151c9565b606654600160ff9092169190911b9081161490565b6066545b6040519081526020016102d5565b6102f161053d366004614aca565b611d14565b610555610550366004614ec0565b611d25565b6040516001600160401b0390911681526020016102d5565b61031461057b3660046151ff565b611d3a565b61043d61058e366004614ba5565b61211b565b61031461212c565b6103296105a9366004614b15565b61213d565b6103837f000000000000000000000000000000000000000000000000000000000000000081565b6103de7f000000000000000000000000000000000000000000000000000000000000000081565b61060f61060a366004615230565b612217565b6040516102d59190615273565b6033546001600160a01b03166103de565b61048461063b366004615285565b6122d3565b61031461064e3660046152e0565b6123bf565b610555610661366004614ec0565b61287e565b610314610674366004615489565b6128ad565b610314610687366004615507565b61295f565b61052161069a366004614ba5565b612cbc565b6103146106ad3660046150ca565b612cde565b6106c56106c0366004614b15565b612e38565b60408051921515835263ffffffff9091166020830152016102d5565b6105216106ef366004614b15565b612ed2565b61043d610702366004614aca565b612ef2565b610314610715366004615549565b612f1b565b610314610728366004614ec0565b613038565b6103de7f000000000000000000000000000000000000000000000000000000000000000081565b610314610762366004614b15565b61315c565b610555610775366004614ec0565b6131d5565b610314610788366004614afe565b6131e1565b604080516060810182525f80825260208201819052918101829052906107bc856107b6866132f7565b8561335a565b925050505b9392505050565b6001600160a01b0382165f908152609e602052604081208190816107eb856132f7565b815260208082019290925260409081015f2081518083019092525460ff8116151580835261010090910463ffffffff16928201929092529150806108395750806020015163ffffffff164311155b9150505b92915050565b60405163237dfb4760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906346fbf68e90602401602060405180830381865afa1580156108a5573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108c99190615573565b6108e657604051631d77d47760e21b815260040160405180910390fd5b606654818116811461090b5760405163c61dca5d60e01b815260040160405180910390fd5b610914826134c6565b5050565b6001600160a01b0381165f908152609d602052604081206060919061093c90613503565b90505f816001600160401b0381111561095757610957614974565b60405190808252806020026020018201604052801561099b57816020015b604080518082019091525f80825260208201528152602001906001900390816109755790505b5090505f5b82811015610a27576001600160a01b0385165f908152609d60205260409020610a02906109cd908361350c565b604080518082019091525f80825260208201525060408051808201909152606082901c815263ffffffff909116602082015290565b828281518110610a1457610a14615592565b60209081029190910101526001016109a0565b509392505050565b60208082015182516001600160a01b03165f90815260989092526040822061083d9163ffffffff9081169061351716565b82610a6a8161352e565b610a875760405163932d94f760e01b815260040160405180910390fd5b6001600160a01b0384165f90815260a4602052604090205460ff16610abf576040516348f7dbb960e01b815260040160405180910390fd5b5f5b82811015610d34575f6040518060400160405280876001600160a01b03168152602001868685818110610af657610af6615592565b9050602002810190610b0891906155a6565b610b169060208101906155c4565b63ffffffff168152509050610b60816020015163ffffffff1660985f896001600160a01b03166001600160a01b031681526020019081526020015f206135d890919063ffffffff16565b610b7d57604051631fb1705560e21b815260040160405180910390fd5b7f31629285ead2335ae0933f86ed2ae63321f7af77b4e6eaabc42c057880977e6c6040518060400160405280886001600160a01b03168152602001836020015163ffffffff16815250604051610bd391906155dd565b60405180910390a15f610be5826132f7565b90505f5b868685818110610bfb57610bfb615592565b9050602002810190610c0d91906155a6565b610c1b9060208101906155eb565b9050811015610d2957610c91878786818110610c3957610c39615592565b9050602002810190610c4b91906155a6565b610c599060208101906155eb565b83818110610c6957610c69615592565b9050602002016020810190610c7e9190614b15565b5f848152609960205260409020906135e3565b507f7ab260fe0af193db5f4986770d831bda4ea46099dc817e8b6716dcae8af8e88b83888887818110610cc657610cc6615592565b9050602002810190610cd891906155a6565b610ce69060208101906155eb565b84818110610cf657610cf6615592565b9050602002016020810190610d0b9190614b15565b604051610d19929190615630565b60405180910390a1600101610be9565b505050600101610ac1565b5050505050565b6060610d49848484436135f7565b949350505050565b6060610d5f858585856135f7565b90505f5b8451811015610de757610d8f858281518110610d8157610d81615592565b6020026020010151876107c8565b610ddf575f5b8451811015610ddd575f838381518110610db157610db1615592565b60200260200101518281518110610dca57610dca615592565b6020908102919091010152600101610d95565b505b600101610d63565b50949350505050565b6001600160a01b038082165f908152609760205260408120549091168015610e1857806107c1565b5090919050565b606654600190600290811603610e485760405163840a48d560e01b815260040160405180910390fd5b82610e528161352e565b610e6f5760405163932d94f760e01b815260040160405180910390fd5b5f6040518060400160405280866001600160a01b03168152602001856020016020810190610e9d91906155c4565b63ffffffff1690529050610eb460608501856155eb565b9050610ec360408601866155eb565b905014610ee3576040516343714afd60e01b815260040160405180910390fd5b60208082015182516001600160a01b03165f90815260989092526040909120610f159163ffffffff9081169061351716565b610f3257604051631fb1705560e21b815260040160405180910390fd5b610f48610f426020860186614b15565b826107c8565b610f655760405163ebbff49760e01b815260040160405180910390fd5b5f610f7360408601866155eb565b90506001600160401b03811115610f8c57610f8c614974565b604051908082528060200260200182016040528015610fb5578160200160208202803683370190505b5090505f5b610fc760408701876155eb565b90508110156115bb5780158061105a5750610fe560408701876155eb565b610ff060018461566a565b818110610fff57610fff615592565b90506020020160208101906110149190614b15565b6001600160a01b031661102a60408801886155eb565b8381811061103a5761103a615592565b905060200201602081019061104f9190614b15565b6001600160a01b0316115b61107757604051639f1c805360e01b815260040160405180910390fd5b61108460608701876155eb565b8281811061109457611094615592565b905060200201355f1080156110d45750670de0b6b3a76400006110ba60608801886155eb565b838181106110ca576110ca615592565b9050602002013511155b6110f157604051631353603160e01b815260040160405180910390fd5b61114d61110160408801886155eb565b8381811061111157611111615592565b90506020020160208101906111269190614b15565b60995f611132876132f7565b81526020019081526020015f206138e490919063ffffffff16565b61116a576040516331bc342760e11b815260040160405180910390fd5b5f806111bc61117c60208a018a614b15565b611185876132f7565b61119260408c018c6155eb565b878181106111a2576111a2615592565b90506020020160208101906111b79190614b15565b61335a565b805191935091506001600160401b03165f036111d95750506115b3565b5f6112146111ea60608b018b6155eb565b868181106111fa576111fa615592565b85516001600160401b031692602090910201359050613905565b835190915061122f6001600160401b0380841690831661391b565b86868151811061124157611241615592565b60200260200101818152505081835f0181815161125e919061567d565b6001600160401b031690525083518290859061127b90839061567d565b6001600160401b031690525060208401805183919061129b90839061567d565b6001600160401b031690525060208301515f600f9190910b12156113b3575f6112fe6112ca60608d018d6155eb565b888181106112da576112da615592565b9050602002013585602001516112ef9061569c565b6001600160801b031690613905565b9050806001600160401b03168460200181815161131b91906156c0565b600f0b9052507f1487af5418c47ee5ea45ef4a93398668120890774a9e13487e61e9dc3baf76dd61134f60208d018d614b15565b8961135d60408f018f6155eb565b8a81811061136d5761136d615592565b90506020020160208101906113829190614b15565b611393885f0151896020015161392f565b88604001516040516113a99594939291906156ed565b60405180910390a1505b6114056113c360208c018c614b15565b6113cc896132f7565b6113d960408e018e6155eb565b898181106113e9576113e9615592565b90506020020160208101906113fe9190614b15565b878761394e565b7f1487af5418c47ee5ea45ef4a93398668120890774a9e13487e61e9dc3baf76dd61143360208c018c614b15565b8861144160408e018e6155eb565b8981811061145157611451615592565b90506020020160208101906114669190614b15565b865160405161147a949392919043906156ed565b60405180910390a16114cb61149260208c018c614b15565b61149f60408d018d6155eb565b888181106114af576114af615592565b90506020020160208101906114c49190614b15565b8651613b86565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001663601bb36f61150760208d018d614b15565b61151460408e018e6155eb565b8981811061152457611524615592565b90506020020160208101906115399190614b15565b875160405160e085901b6001600160e01b03191681526001600160a01b0393841660048201529290911660248301526001600160401b0380861660448401521660648201526084015f604051808303815f87803b158015611598575f5ffd5b505af11580156115aa573d5f5f3e3d5ffd5b50505050505050505b600101610fba565b507f80969ad29428d6797ee7aad084f9e4a42a82fc506dcd2ca3b6fb431f85ccebe56115ea6020870187614b15565b836115f860408901896155eb565b8561160660808c018c61573e565b60405161161997969594939291906157a8565b60405180910390a1505050505050565b6001600160a01b0382165f908152609d60205260408120606091829161164e90613503565b90505f816001600160401b0381111561166957611669614974565b6040519080825280602002602001820160405280156116ad57816020015b604080518082019091525f80825260208201528152602001906001900390816116875790505b5090505f826001600160401b038111156116c9576116c9614974565b60405190808252806020026020018201604052801561171257816020015b604080516060810182525f80825260208083018290529282015282525f199092019101816116e75790505b5090505f5b83811015611795576001600160a01b0388165f908152609d60205260408120611744906109cd908461350c565b90508084838151811061175957611759615592565b602002602001018190525061176f89828a61078d565b83838151811061178157611781615592565b602090810291909101015250600101611717565b509093509150505b9250929050565b60605f6107c160995f6117b6866132f7565b81526020019081526020015f20613c08565b60605f83516001600160401b038111156117e4576117e4614974565b60405190808252806020026020018201604052801561180d578160200160208202803683370190505b5090505f5b8451811015610a275761183e85828151811061183057611830615592565b60200260200101518561287e565b82828151811061185057611850615592565b6001600160401b0390921660209283029190910190910152600101611812565b6066545f906001908116036118985760405163840a48d560e01b815260040160405180910390fd5b8382146118b8576040516343714afd60e01b815260040160405180910390fd5b5f5b8481101561192157611919878787848181106118d8576118d8615592565b90506020020160208101906118ed9190614b15565b8686858181106118ff576118ff615592565b9050602002016020810190611914919061583e565b613c14565b6001016118ba565b50505050505050565b836119348161352e565b6119515760405163932d94f760e01b815260040160405180910390fd5b604080518082019091526001600160a01b038616815263ffffffff851660208201525f61197d826132f7565b90506119be826020015163ffffffff1660985f8a6001600160a01b03166001600160a01b031681526020019081526020015f2061351790919063ffffffff16565b6119db57604051631fb1705560e21b815260040160405180910390fd5b5f5b84811015611a7e576119fa868683818110610c6957610c69615592565b611a175760405163585cfb2f60e01b815260040160405180910390fd5b7f7ab260fe0af193db5f4986770d831bda4ea46099dc817e8b6716dcae8af8e88b83878784818110611a4b57611a4b615592565b9050602002016020810190611a609190614b15565b604051611a6e929190615630565b60405180910390a16001016119dd565b5050505050505050565b60605f82516001600160401b03811115611aa457611aa4614974565b604051908082528060200260200182016040528015611acd578160200160208202803683370190505b5090505f5b8351811015610a2757611afe85858381518110611af157611af1615592565b602002602001015161287e565b828281518110611b1057611b10615592565b6001600160401b0390921660209283029190910190910152600101611ad2565b6060611b5b7f0000000000000000000000000000000000000000000000000000000000000000613d18565b905090565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614611c5b57611b998261352e565b611bb6576040516348f5c3ed60e01b815260040160405180910390fd5b6040516336b87bd760e11b81526001600160a01b0383811660048301527f00000000000000000000000000000000000000000000000000000000000000001690636d70f7ae90602401602060405180830381865afa158015611c1a573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611c3e9190615573565b611c5b5760405163ccea9e6f60e01b815260040160405180910390fd5b6109148282613d55565b60405163237dfb4760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906346fbf68e90602401602060405180830381865afa158015611cc7573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611ceb9190615573565b611d0857604051631d77d47760e21b815260040160405180910390fd5b611d125f196134c6565b565b5f6107c183609a5f611132866132f7565b5f5f611d318484613f01565b95945050505050565b606654600290600490811603611d635760405163840a48d560e01b815260040160405180910390fd5b611d78611d736020840184614b15565b61352e565b80611d915750611d91611d736040840160208501614b15565b611dae576040516348f5c3ed60e01b815260040160405180910390fd5b5f5b611dbd60408401846155eb565b905081101561207f575f6040518060400160405280856020016020810190611de59190614b15565b6001600160a01b03168152602001611e0060408701876155eb565b85818110611e1057611e10615592565b9050602002016020810190611e2591906155c4565b63ffffffff168152509050611e72816020015163ffffffff1660985f876020016020810190611e549190614b15565b6001600160a01b0316815260208101919091526040015f2090613517565b611e8f57604051631fb1705560e21b815260040160405180910390fd5b609e5f611e9f6020870187614b15565b6001600160a01b03166001600160a01b031681526020019081526020015f205f611ec8836132f7565b815260208101919091526040015f205460ff16611ef8576040516325131d4f60e01b815260040160405180910390fd5b611f32611f04826132f7565b609c5f611f146020890189614b15565b6001600160a01b0316815260208101919091526040015f2090614070565b50611f6a611f436020860186614b15565b609a5f611f4f856132f7565b81526020019081526020015f2061407b90919063ffffffff16565b50611f786020850185614b15565b6001600160a01b03167fad34c3070be1dffbcaa499d000ba2b8d9848aefcac3059df245dd95c4ece14fe82604051611fb091906155dd565b60405180910390a2604080518082019091525f815260208101611ff37f00000000000000000000000000000000000000000000000000000000000000004361585f565b63ffffffff169052609e5f61200b6020880188614b15565b6001600160a01b03166001600160a01b031681526020019081526020015f205f612034846132f7565b81526020808201929092526040015f2082518154939092015163ffffffff166101000264ffffffff00199215159290921664ffffffffff199093169290921717905550600101611db0565b506120936103d96040840160208501614b15565b6001600160a01b031663303ca9566120ae6020850185614b15565b6120be6040860160208701614b15565b6120cb60408701876155eb565b6040518563ffffffff1660e01b81526004016120ea94939291906158b4565b5f604051808303815f87803b158015612101575f5ffd5b505af1158015612113573d5f5f3e3d5ffd5b505050505050565b606061083d609a5f6117b6856132f7565b61213461408f565b611d125f6140e9565b6001600160a01b0381165f908152609c602052604081206060919061216190613503565b90505f816001600160401b0381111561217c5761217c614974565b6040519080825280602002602001820160405280156121c057816020015b604080518082019091525f808252602082015281526020019060019003908161219a5790505b5090505f5b82811015610a27576001600160a01b0385165f908152609c602052604090206121f2906109cd908361350c565b82828151811061220457612204615592565b60209081029190910101526001016121c5565b60605f84516001600160401b0381111561223357612233614974565b60405190808252806020026020018201604052801561227c57816020015b604080516060810182525f80825260208083018290529282015282525f199092019101816122515790505b5090505f5b8551811015610de7576122ae86828151811061229f5761229f615592565b6020026020010151868661078d565b8282815181106122c0576122c0615592565b6020908102919091010152600101612281565b60605f83516001600160401b038111156122ef576122ef614974565b604051908082528060200260200182016040528015612318578160200160208202803683370190505b5090505f5b8451811015610de7576001600160a01b0386165f90815260a160205260408120865161238d9287929189908690811061235857612358615592565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020015f2061413a90919063ffffffff16565b82828151811061239f5761239f615592565b6001600160401b039092166020928302919091019091015260010161231d565b6066545f906001908116036123e75760405163840a48d560e01b815260040160405180910390fd5b6123f08361352e565b61240d576040516348f5c3ed60e01b815260040160405180910390fd5b5f5f5f61241986612e38565b915091508161243b5760405163fa55fc8160e01b815260040160405180910390fd5b91505f90505b8351811015610d345783818151811061245c5761245c615592565b6020026020010151604001515184828151811061247b5761247b615592565b60200260200101516020015151146124a6576040516343714afd60e01b815260040160405180910390fd5b5f8482815181106124b9576124b9615592565b602090810291909101810151518082015181516001600160a01b03165f908152609890935260409092209092506124f99163ffffffff9081169061351716565b61251657604051631fb1705560e21b815260040160405180910390fd5b5f61252187836107c8565b90505f5b86848151811061253757612537615592565b60200260200101516020015151811015612873575f87858151811061255e5761255e615592565b602002602001015160200151828151811061257b5761257b615592565b60200260200101519050612592898261ffff613c14565b5f5f6125a18b6107b6886132f7565b91509150806040015163ffffffff165f146125cf57604051630d8fcbe360e41b815260040160405180910390fd5b5f6125dc8785848961414e565b9050612621825f01518c8a815181106125f7576125f7615592565b602002602001015160400151878151811061261457612614615592565b6020026020010151614184565b600f0b602083018190525f0361264a57604051634606179360e11b815260040160405180910390fd5b5f8260200151600f0b121561278e578015612710576126cb61266b886132f7565b6001600160a01b03808f165f90815260a360209081526040808320938a16835292905220908154600160801b90819004600f0b5f818152600180860160205260409091209390935583546001600160801b03908116939091011602179055565b6126f57f00000000000000000000000000000000000000000000000000000000000000004361585f565b61270090600161585f565b63ffffffff1660408301526127fb565b6127228360200151836020015161392f565b6001600160401b031660208401528a518b908990811061274457612744615592565b602002602001015160400151858151811061276157612761615592565b6020908102919091018101516001600160401b031683525f9083015263ffffffff431660408301526127fb565b5f8260200151600f0b13156127fb576127af8360200151836020015161392f565b6001600160401b0390811660208501819052845190911610156127e557604051636c9be0bf60e01b815260040160405180910390fd5b6127ef894361585f565b63ffffffff1660408301525b6128108c612808896132f7565b86868661394e565b7f1487af5418c47ee5ea45ef4a93398668120890774a9e13487e61e9dc3baf76dd8c8886612845865f0151876020015161392f565b866040015160405161285b9594939291906156ed565b60405180910390a15050600190920191506125259050565b505050600101612441565b6001600160a01b038083165f90815260a16020908152604080832093851683529290529081206107c19061419b565b826128b78161352e565b6128d45760405163932d94f760e01b815260040160405180910390fd5b6001600160a01b0384165f90815260a4602052604090205460ff16612916576001600160a01b0384165f90815260a460205260409020805460ff191660011790555b836001600160a01b03167fa89c1dc243d8908a96dd84944bcc97d6bc6ac00dd78e20621576be6a3c94371384846040516129519291906158e0565b60405180910390a250505050565b6066546002906004908116036129885760405163840a48d560e01b815260040160405180910390fd5b826129928161352e565b6129af5760405163932d94f760e01b815260040160405180910390fd5b6040516336b87bd760e11b81526001600160a01b0385811660048301527f00000000000000000000000000000000000000000000000000000000000000001690636d70f7ae90602401602060405180830381865afa158015612a13573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612a379190615573565b612a545760405163ccea9e6f60e01b815260040160405180910390fd5b5f5b612a6360208501856155eb565b9050811015612c2b57604080518082019091525f9080612a866020880188614b15565b6001600160a01b03168152602001868060200190612aa491906155eb565b85818110612ab457612ab4615592565b9050602002016020810190612ac991906155c4565b63ffffffff90811690915260208083015183516001600160a01b03165f90815260989092526040909120929350612b0592919081169061351716565b612b2257604051631fb1705560e21b815260040160405180910390fd5b612b2c86826107c8565b15612b4a57604051636c6c6e2760e11b815260040160405180910390fd5b612b73612b56826132f7565b6001600160a01b0388165f908152609c60205260409020906135d8565b50612b9f86609a5f612b84856132f7565b81526020019081526020015f206135e390919063ffffffff16565b50856001600160a01b03167f43232edf9071753d2321e5fa7e018363ee248e5f2142e6c08edd3265bfb4895e82604051612bd991906155dd565b60405180910390a26001600160a01b0386165f908152609e60205260408120600191612c04846132f7565b815260208101919091526040015f20805460ff191691151591909117905550600101612a56565b50612c3c6103d96020850185614b15565b6001600160a01b031663c63fd50285612c586020870187614b15565b612c6560208801886155eb565b612c7260408a018a61573e565b6040518763ffffffff1660e01b8152600401612c93969594939291906158f3565b5f604051808303815f87803b158015612caa575f5ffd5b505af1158015611a7e573d5f5f3e3d5ffd5b5f61083d609a5f612ccc856132f7565b81526020019081526020015f20613503565b83612ce88161352e565b612d055760405163932d94f760e01b815260040160405180910390fd5b6040805180820182526001600160a01b03871680825263ffffffff80881660208085018290525f93845260989052939091209192612d44929161351716565b612d6157604051631fb1705560e21b815260040160405180910390fd5b5f612d6b826132f7565b90505f5b84811015611a7e57612db4868683818110612d8c57612d8c615592565b9050602002016020810190612da19190614b15565b5f8481526099602052604090209061407b565b612dd1576040516331bc342760e11b815260040160405180910390fd5b7f7b4b073d80dcac55a11177d8459ad9f664ceeb91f71f27167bb14f8152a7eeee83878784818110612e0557612e05615592565b9050602002016020810190612e1a9190614b15565b604051612e28929190615630565b60405180910390a1600101612d6f565b6001600160a01b0381165f908152609b602090815260408083208151608081018352905463ffffffff80821680845260ff600160201b8404161515958401869052650100000000008304821694840194909452600160481b909104166060820181905284939192919015801590612eb95750826060015163ffffffff164310155b15612ec8575050604081015160015b9590945092505050565b6001600160a01b0381165f90815260986020526040812061083d90613503565b6001600160a01b0382165f908152609f602052604081206060919061083990826117b6866132f7565b5f54610100900460ff1615808015612f3957505f54600160ff909116105b80612f525750303b158015612f5257505f5460ff166001145b612fba5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b5f805460ff191660011790558015612fdb575f805461ff0019166101001790555b612fe4826134c6565b612fed836140e9565b8015613033575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498906020015b60405180910390a15b505050565b816130428161352e565b61305f5760405163932d94f760e01b815260040160405180910390fd5b60405163b526578760e01b81526001600160a01b03848116600483015283169063b526578790602401602060405180830381865afa1580156130a3573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906130c79190615573565b6130e457604051631d0b13c160e31b815260040160405180910390fd5b6001600160a01b038381165f90815260976020526040902080546001600160a01b0319169184169190911790557f2ae945c40c44dc0ec263f95609c3fdc6952e0aefa22d6374e44f2c997acedf858361313c81610df0565b604080516001600160a01b0393841681529290911660208301520161302a565b61316461408f565b6001600160a01b0381166131c95760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401612fb1565b6131d2816140e9565b50565b5f5f610de78484613f01565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561323d573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613261919061593f565b6001600160a01b0316336001600160a01b0316146132925760405163794821ff60e01b815260040160405180910390fd5b606654801982198116146132b95760405163c61dca5d60e01b815260040160405180910390fd5b606682905560405182815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c9060200160405180910390a25050565b5f815f0151826020015163ffffffff1660405160200161334292919060609290921b6bffffffffffffffffffffffff1916825260a01b6001600160a01b031916601482015260200190565b60405160208183030381529060405261083d9061595a565b6040805180820182525f80825260208083018290528351606081018552828152808201839052808501839052845180860186526001600160a01b03898116855260a18452868520908816855290925293822092939281906133ba9061419b565b6001600160401b0390811682526001600160a01b038981165f81815260a260209081526040808320948c168084529482528083205486169682019690965291815260a082528481208b8252825284812092815291815290839020835160608101855290549283168152600160401b8304600f0b91810191909152600160c01b90910463ffffffff1691810182905291925043101561345c5790925090506134be565b61346d815f0151826020015161392f565b6001600160401b0316815260208101515f600f9190910b12156134ab5761349c8260200151826020015161392f565b6001600160401b031660208301525b5f60408201819052602082015290925090505b935093915050565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a250565b5f61083d825490565b5f6107c183836141ae565b5f81815260018301602052604081205415156107c1565b604051631beb2b9760e31b81526001600160a01b0382811660048301523360248301523060448301525f80356001600160e01b0319166064840152917f00000000000000000000000000000000000000000000000000000000000000009091169063df595cb8906084016020604051808303815f875af11580156135b4573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061083d9190615573565b5f6107c183836141d4565b5f6107c1836001600160a01b0384166141d4565b606083516001600160401b0381111561361257613612614974565b60405190808252806020026020018201604052801561364557816020015b60608152602001906001900390816136305790505b5090505f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f0e0e67686866040518363ffffffff1660e01b815260040161369792919061597d565b5f60405180830381865afa1580156136b1573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526136d891908101906159a1565b90505f5b85518110156138da575f8682815181106136f8576136f8615592565b6020026020010151905085516001600160401b0381111561371b5761371b614974565b604051908082528060200260200182016040528015613744578160200160208202803683370190505b5084838151811061375757613757615592565b60209081029190910101525f5b86518110156138d0575f87828151811061378057613780615592565b6020908102919091018101516001600160a01b038086165f90815260a18452604080822092841682529190935282209092506137bb9061419b565b9050806001600160401b03165f036137d45750506138c8565b5f6137e0858d8561078d565b90508863ffffffff16816040015163ffffffff161115801561380857505f8160200151600f0b125b1561382a5761381e815f0151826020015161392f565b6001600160401b031681525b80515f90613845906001600160401b0390811690851661391b565b905061388c8189898151811061385d5761385d615592565b6020026020010151878151811061387657613876615592565b602002602001015161422090919063ffffffff16565b89888151811061389e5761389e615592565b602002602001015186815181106138b7576138b7615592565b602002602001018181525050505050505b600101613764565b50506001016136dc565b5050949350505050565b6001600160a01b0381165f90815260018301602052604081205415156107c1565b5f6107c18383670de0b6b3a76400006001614234565b5f6107c183670de0b6b3a76400008461428d565b5f6107c1613946836001600160401b0386166156c0565b600f0b614372565b6020808301516001600160a01b038088165f90815260a284526040808220928816825291909352909120546001600160401b03908116911614613a1457602082810180516001600160a01b038881165f81815260a286526040808220938a1680835293875290819020805467ffffffffffffffff19166001600160401b0395861617905593518451918252948101919091529216908201527facf9095feb3a370c9cf692421c69ef320d4db5c66e6a7d29c7694eb02364fc559060600160405180910390a15b6001600160a01b038086165f90815260a060209081526040808320888452825280832093871683529281529082902083518154928501519385015163ffffffff16600160c01b0263ffffffff60c01b196001600160801b038616600160401b026001600160c01b03199095166001600160401b03909316929092179390931716919091179055600f0b15613af6576001600160a01b0385165f908152609f602090815260408083208784529091529020613ace90846135e3565b506001600160a01b0385165f908152609d60205260409020613af090856135d8565b50610d34565b80516001600160401b03165f03610d34576001600160a01b0385165f908152609f602090815260408083208784529091529020613b33908461407b565b506001600160a01b0385165f908152609f602090815260408083208784529091529020613b5f90613503565b5f03610d34576001600160a01b0385165f908152609d602052604090206121139085614070565b6001600160a01b038084165f90815260a160209081526040808320938616835292905220613bb59043836143dd565b604080516001600160a01b038086168252841660208201526001600160401b038316918101919091527f1c6458079a41077d003c11faf9bf097e693bd67979e4e6500bac7b29db779b5c9060600161302a565b60605f6107c1836143f1565b6001600160a01b038381165f90815260a360209081526040808320938616835292905290812054600f81810b600160801b909204900b035b5f81118015613c5e57508261ffff1682105b15610d34576001600160a01b038086165f90815260a3602090815260408083209388168352929052908120613c929061444a565b90505f5f613ca188848961335a565b91509150806040015163ffffffff16431015613cbf57505050610d34565b613ccc888489858561394e565b6001600160a01b038089165f90815260a360209081526040808320938b16835292905220613cf99061449c565b50613d0385615aad565b9450613d0e84615ac5565b9350505050613c4c565b60605f613d2483614519565b6040805160208082528183019092529192505f91906020820181803683375050509182525060208101929092525090565b6001600160a01b0382165f908152609b60209081526040918290208251608081018452905463ffffffff808216835260ff600160201b830416151593830193909352650100000000008104831693820193909352600160481b909204166060820181905215801590613dd15750806060015163ffffffff164310155b15613deb57604081015163ffffffff168152600160208201525b63ffffffff82166040820152613e217f00000000000000000000000000000000000000000000000000000000000000004361585f565b613e2c90600161585f565b63ffffffff90811660608381019182526001600160a01b0386165f818152609b602090815260409182902087518154838a0151858b01519851928a1664ffffffffff1990921691909117600160201b91151591909102176cffffffffffffffff0000000000191665010000000000978916979097026cffffffff000000000000000000191696909617600160481b968816968702179055815192835294871694820194909452928301919091527f4e85751d6331506c6c62335f207eb31f12a61e570f34f5c17640308785c6d4db910161302a565b6001600160a01b038281165f81815260a2602090815260408083209486168084529482528083205493835260a38252808320948352939052918220546001600160401b039091169190600f81810b600160801b909204900b03815b8181101561402c576001600160a01b038087165f90815260a3602090815260408083209389168352929052908120613f949083614540565b6001600160a01b038881165f90815260a0602090815260408083208584528252808320938b16835292815290829020825160608101845290546001600160401b0381168252600160401b8104600f0b92820192909252600160c01b90910463ffffffff1691810182905291925043101561400f57505061402c565b61401d86826020015161392f565b95505050806001019050613f5c565b506001600160a01b038086165f90815260a160209081526040808320938816835292905220839061405c9061419b565b614066919061567d565b9150509250929050565b5f6107c183836145af565b5f6107c1836001600160a01b0384166145af565b6033546001600160a01b03163314611d125760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401612fb1565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b5f6107c18383670de0b6b3a7640000614692565b5f61415f8460995f611132896132f7565b80156141685750815b8015611d3157505090516001600160401b031615159392505050565b5f6107c16001600160401b03808516908416615ada565b5f61083d82670de0b6b3a76400006146e7565b5f825f0182815481106141c3576141c3615592565b905f5260205f200154905092915050565b5f81815260018301602052604081205461421957508154600181810184555f84815260208082209093018490558454848252828601909352604090209190915561083d565b505f61083d565b5f6107c18383670de0b6b3a764000061428d565b5f5f61424186868661428d565b9050600183600281111561425757614257615b07565b14801561427357505f848061426e5761426e615b1b565b868809115b15611d3157614283600182615b2f565b9695505050505050565b5f80805f19858709858702925082811083820303915050805f036142c4578382816142ba576142ba615b1b565b04925050506107c1565b80841161430b5760405162461bcd60e51b81526020600482015260156024820152744d6174683a206d756c446976206f766572666c6f7760581b6044820152606401612fb1565b5f8486880960026001871981018816978890046003810283188082028403028082028403028082028403028082028403028082028403029081029092039091025f889003889004909101858311909403939093029303949094049190911702949350505050565b5f6001600160401b038211156143d95760405162461bcd60e51b815260206004820152602660248201527f53616665436173743a2076616c756520646f65736e27742066697420696e203660448201526534206269747360d01b6064820152608401612fb1565b5090565b61303383836001600160401b03841661471e565b6060815f0180548060200260200160405190810160405280929190818152602001828054801561443e57602002820191905f5260205f20905b81548152602001906001019080831161442a575b50505050509050919050565b5f6144648254600f81810b600160801b909204900b131590565b1561448257604051631ed9509560e11b815260040160405180910390fd5b508054600f0b5f9081526001909101602052604090205490565b5f6144b68254600f81810b600160801b909204900b131590565b156144d457604051631ed9509560e11b815260040160405180910390fd5b508054600f0b5f818152600180840160205260408220805492905583546fffffffffffffffffffffffffffffffff191692016001600160801b03169190911790915590565b5f60ff8216601f81111561083d57604051632cd44ac360e21b815260040160405180910390fd5b5f5f61456261454e84614821565b855461455d9190600f0b615b42565b61488a565b8454909150600160801b9004600f90810b9082900b1261459557604051632d0483c560e21b815260040160405180910390fd5b600f0b5f9081526001939093016020525050604090205490565b5f8181526001830160205260408120548015614689575f6145d160018361566a565b85549091505f906145e49060019061566a565b9050818114614643575f865f01828154811061460257614602615592565b905f5260205f200154905080875f01848154811061462257614622615592565b5f918252602080832090910192909255918252600188019052604090208390555b855486908061465457614654615b69565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f90556001935050505061083d565b5f91505061083d565b82545f90816146a3868683856148f3565b905080156146dd576146c7866146ba60018461566a565b5f91825260209091200190565b54600160201b90046001600160e01b03166107bc565b5091949350505050565b81545f90801561471657614700846146ba60018461566a565b54600160201b90046001600160e01b0316610839565b509092915050565b825480156147d4575f614736856146ba60018561566a565b60408051808201909152905463ffffffff808216808452600160201b9092046001600160e01b0316602084015291925090851610156147885760405163151b8e3f60e11b815260040160405180910390fd5b805163ffffffff8086169116036147d257826147a9866146ba60018661566a565b80546001600160e01b0392909216600160201b0263ffffffff9092169190911790555050505050565b505b506040805180820190915263ffffffff92831681526001600160e01b03918216602080830191825285546001810187555f968752952091519051909216600160201b029190921617910155565b5f6001600160ff1b038211156143d95760405162461bcd60e51b815260206004820152602860248201527f53616665436173743a2076616c756520646f65736e27742066697420696e2061604482015267371034b73a191a9b60c11b6064820152608401612fb1565b80600f81900b81146148ee5760405162461bcd60e51b815260206004820152602760248201527f53616665436173743a2076616c756520646f65736e27742066697420696e20316044820152663238206269747360c81b6064820152608401612fb1565b919050565b5f5b81831015610a27575f6149088484614946565b5f8781526020902090915063ffffffff86169082015463ffffffff16111561493257809250614940565b61493d816001615b2f565b93505b506148f5565b5f6149546002848418615b7d565b6107c190848416615b2f565b6001600160a01b03811681146131d2575f5ffd5b634e487b7160e01b5f52604160045260245ffd5b604051606081016001600160401b03811182821017156149aa576149aa614974565b60405290565b604051601f8201601f191681016001600160401b03811182821017156149d8576149d8614974565b604052919050565b803563ffffffff811681146148ee575f5ffd5b5f60408284031215614a03575f5ffd5b604080519081016001600160401b0381118282101715614a2557614a25614974565b6040529050808235614a3681614960565b8152614a44602084016149e0565b60208201525092915050565b5f5f5f60808486031215614a62575f5ffd5b8335614a6d81614960565b9250614a7c85602086016149f3565b91506060840135614a8c81614960565b809150509250925092565b81516001600160401b03168152602080830151600f0b9082015260408083015163ffffffff16908201526060810161083d565b5f5f60608385031215614adb575f5ffd5b8235614ae681614960565b9150614af584602085016149f3565b90509250929050565b5f60208284031215614b0e575f5ffd5b5035919050565b5f60208284031215614b25575f5ffd5b81356107c181614960565b80516001600160a01b0316825260209081015163ffffffff16910152565b5f8151808452602084019350602083015f5b82811015614b8957614b73868351614b30565b6040959095019460209190910190600101614b60565b5093949350505050565b602081525f6107c16020830184614b4e565b5f60408284031215614bb5575f5ffd5b6107c183836149f3565b5f5f83601f840112614bcf575f5ffd5b5081356001600160401b03811115614be5575f5ffd5b6020830191508360208260051b850101111561179d575f5ffd5b5f5f5f60408486031215614c11575f5ffd5b8335614c1c81614960565b925060208401356001600160401b03811115614c36575f5ffd5b614c4286828701614bbf565b9497909650939450505050565b5f6001600160401b03821115614c6757614c67614974565b5060051b60200190565b5f82601f830112614c80575f5ffd5b8135614c93614c8e82614c4f565b6149b0565b8082825260208201915060208360051b860101925085831115614cb4575f5ffd5b602085015b83811015614cda578035614ccc81614960565b835260209283019201614cb9565b5095945050505050565b5f5f5f60808486031215614cf6575f5ffd5b614d0085856149f3565b925060408401356001600160401b03811115614d1a575f5ffd5b614d2686828701614c71565b92505060608401356001600160401b03811115614d41575f5ffd5b614d4d86828701614c71565b9150509250925092565b5f8151808452602084019350602083015f5b82811015614b89578151865260209586019590910190600101614d69565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b82811015614dde57603f19878603018452614dc9858351614d57565b94506020938401939190910190600101614dad565b50929695505050505050565b5f5f5f5f60a08587031215614dfd575f5ffd5b614e0786866149f3565b935060408501356001600160401b03811115614e21575f5ffd5b614e2d87828801614c71565b93505060608501356001600160401b03811115614e48575f5ffd5b614e5487828801614c71565b925050614e63608086016149e0565b905092959194509250565b5f5f60408385031215614e7f575f5ffd5b8235614e8a81614960565b915060208301356001600160401b03811115614ea4575f5ffd5b830160a08186031215614eb5575f5ffd5b809150509250929050565b5f5f60408385031215614ed1575f5ffd5b8235614edc81614960565b91506020830135614eb581614960565b5f8151808452602084019350602083015f5b82811015614b8957614f3786835180516001600160401b03168252602080820151600f0b9083015260409081015163ffffffff16910152565b6060959095019460209190910190600101614efe565b604081525f614f5f6040830185614b4e565b8281036020840152611d318185614eec565b5f8151808452602084019350602083015f5b82811015614b895781516001600160a01b0316865260209586019590910190600101614f83565b602081525f6107c16020830184614f71565b5f5f60408385031215614fcd575f5ffd5b82356001600160401b03811115614fe2575f5ffd5b614fee85828601614c71565b9250506020830135614eb581614960565b602080825282518282018190525f918401906040840190835b8181101561503f5783516001600160401b0316835260209384019390920191600101615018565b509095945050505050565b5f5f5f5f5f6060868803121561505e575f5ffd5b853561506981614960565b945060208601356001600160401b03811115615083575f5ffd5b61508f88828901614bbf565b90955093505060408601356001600160401b038111156150ad575f5ffd5b6150b988828901614bbf565b969995985093965092949392505050565b5f5f5f5f606085870312156150dd575f5ffd5b84356150e881614960565b93506150f6602086016149e0565b925060408501356001600160401b03811115615110575f5ffd5b61511c87828801614bbf565b95989497509550505050565b5f5f60408385031215615139575f5ffd5b823561514481614960565b915060208301356001600160401b0381111561515e575f5ffd5b61406685828601614c71565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b5f5f604083850312156151b0575f5ffd5b82356151bb81614960565b9150614af5602084016149e0565b5f602082840312156151d9575f5ffd5b813560ff811681146107c1575f5ffd5b5f606082840312156151f9575f5ffd5b50919050565b5f6020828403121561520f575f5ffd5b81356001600160401b03811115615224575f5ffd5b610839848285016151e9565b5f5f5f60808486031215615242575f5ffd5b83356001600160401b03811115615257575f5ffd5b61526386828701614c71565b935050614a7c85602086016149f3565b602081525f6107c16020830184614eec565b5f5f5f60608486031215615297575f5ffd5b83356152a281614960565b925060208401356001600160401b038111156152bc575f5ffd5b6152c886828701614c71565b9250506152d7604085016149e0565b90509250925092565b5f5f604083850312156152f1575f5ffd5b82356152fc81614960565b915060208301356001600160401b03811115615316575f5ffd5b8301601f81018513615326575f5ffd5b8035615334614c8e82614c4f565b8082825260208201915060208360051b850101925087831115615355575f5ffd5b602084015b8381101561547a5780356001600160401b03811115615377575f5ffd5b85016080818b03601f1901121561538c575f5ffd5b615394614988565b6153a18b602084016149f3565b815260608201356001600160401b038111156153bb575f5ffd5b6153ca8c602083860101614c71565b60208301525060808201356001600160401b038111156153e8575f5ffd5b6020818401019250508a601f8301126153ff575f5ffd5b813561540d614c8e82614c4f565b8082825260208201915060208360051b86010192508d83111561542e575f5ffd5b6020850194505b828510156154645784356001600160401b0381168114615453575f5ffd5b825260209485019490910190615435565b604084015250508452506020928301920161535a565b50809450505050509250929050565b5f5f5f6040848603121561549b575f5ffd5b83356154a681614960565b925060208401356001600160401b038111156154c0575f5ffd5b8401601f810186136154d0575f5ffd5b80356001600160401b038111156154e5575f5ffd5b8660208284010111156154f6575f5ffd5b939660209190910195509293505050565b5f5f60408385031215615518575f5ffd5b823561552381614960565b915060208301356001600160401b0381111561553d575f5ffd5b614066858286016151e9565b5f5f6040838503121561555a575f5ffd5b823561556581614960565b946020939093013593505050565b5f60208284031215615583575f5ffd5b815180151581146107c1575f5ffd5b634e487b7160e01b5f52603260045260245ffd5b5f8235603e198336030181126155ba575f5ffd5b9190910192915050565b5f602082840312156155d4575f5ffd5b6107c1826149e0565b6040810161083d8284614b30565b5f5f8335601e19843603018112615600575f5ffd5b8301803591506001600160401b03821115615619575f5ffd5b6020019150600581901b360382131561179d575f5ffd5b6060810161563e8285614b30565b6001600160a01b039290921660409190910152919050565b634e487b7160e01b5f52601160045260245ffd5b8181038181111561083d5761083d615656565b6001600160401b03828116828216039081111561083d5761083d615656565b5f81600f0b60016001607f1b031981036156b8576156b8615656565b5f0392915050565b600f81810b9083900b0160016001607f1b03811360016001607f1b03198212171561083d5761083d615656565b6001600160a01b038616815260c0810161570a6020830187614b30565b6001600160a01b039490941660608201526001600160401b0392909216608083015263ffffffff1660a09091015292915050565b5f5f8335601e19843603018112615753575f5ffd5b8301803591506001600160401b0382111561576c575f5ffd5b60200191503681900382131561179d575f5ffd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b6001600160a01b03881681525f60c082016157c6602084018a614b30565b60c060608401528690528660e083015f5b888110156158075782356157ea81614960565b6001600160a01b03168252602092830192909101906001016157d7565b50838103608085015261581a8188614d57565b91505082810360a0840152615830818587615780565b9a9950505050505050505050565b5f6020828403121561584e575f5ffd5b813561ffff811681146107c1575f5ffd5b63ffffffff818116838216019081111561083d5761083d615656565b8183526020830192505f815f5b84811015614b895763ffffffff61589e836149e0565b1686526020958601959190910190600101615888565b6001600160a01b038581168252841660208201526060604082018190525f90614283908301848661587b565b602081525f610d49602083018486615780565b6001600160a01b038781168252861660208201526080604082018190525f9061591f908301868861587b565b8281036060840152615932818587615780565b9998505050505050505050565b5f6020828403121561594f575f5ffd5b81516107c181614960565b805160208083015191908110156151f9575f1960209190910360031b1b16919050565b604081525f61598f6040830185614f71565b8281036020840152611d318185614f71565b5f602082840312156159b1575f5ffd5b81516001600160401b038111156159c6575f5ffd5b8201601f810184136159d6575f5ffd5b80516159e4614c8e82614c4f565b8082825260208201915060208360051b850101925086831115615a05575f5ffd5b602084015b83811015615aa25780516001600160401b03811115615a27575f5ffd5b8501603f81018913615a37575f5ffd5b6020810151615a48614c8e82614c4f565b808282526020820191506020808460051b8601010192508b831115615a6b575f5ffd5b6040840193505b82841015615a8d578351825260209384019390910190615a72565b86525050602093840193919091019050615a0a565b509695505050505050565b5f60018201615abe57615abe615656565b5060010190565b5f81615ad357615ad3615656565b505f190190565b600f82810b9082900b0360016001607f1b0319811260016001607f1b038213171561083d5761083d615656565b634e487b7160e01b5f52602160045260245ffd5b634e487b7160e01b5f52601260045260245ffd5b8082018082111561083d5761083d615656565b8082018281125f831280158216821582161715615b6157615b61615656565b505092915050565b634e487b7160e01b5f52603160045260245ffd5b5f82615b9757634e487b7160e01b5f52601260045260245ffd5b50049056fea2646970667358221220eb9c1f2290fb8dd6dd060f601e3b3d489683796b0ae15caffe25b8fdb623b55464736f6c634300081b0033",
=======
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_delegation\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"_permissionController\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"},{\"name\":\"_DEALLOCATION_DELAY\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_ALLOCATION_CONFIGURATION_DELAY\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ALLOCATION_CONFIGURATION_DELAY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEALLOCATION_DELAY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addStrategiesToOperatorSet\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"clearDeallocationQueue\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"numToClear\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createOperatorSets\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"params\",\"type\":\"tuple[]\",\"internalType\":\"structIAllocationManagerTypes.CreateSetParams[]\",\"components\":[{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregisterFromOperatorSets\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIAllocationManagerTypes.DeregisterParams\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAVSRegistrar\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAVSRegistrar\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocatableMagnitude\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocatedSets\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocatedStrategies\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocation\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIAllocationManagerTypes.Allocation\",\"components\":[{\"name\":\"currentMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"pendingDiff\",\"type\":\"int128\",\"internalType\":\"int128\"},{\"name\":\"effectBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocationDelay\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocations\",\"inputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIAllocationManagerTypes.Allocation[]\",\"components\":[{\"name\":\"currentMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"pendingDiff\",\"type\":\"int128\",\"internalType\":\"int128\"},{\"name\":\"effectBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getEncumberedMagnitude\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxMagnitude\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxMagnitudes\",\"inputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxMagnitudes\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxMagnitudesAtBlock\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMemberCount\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMembers\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMinimumSlashableStake\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"futureBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"slashableStake\",\"type\":\"uint256[][]\",\"internalType\":\"uint256[][]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorSetCount\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRegisteredSets\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStrategiesInOperatorSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStrategyAllocations\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIAllocationManagerTypes.Allocation[]\",\"components\":[{\"name\":\"currentMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"pendingDiff\",\"type\":\"int128\",\"internalType\":\"int128\"},{\"name\":\"effectBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isMemberOfOperatorSet\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperatorSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"modifyAllocations\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"params\",\"type\":\"tuple[]\",\"internalType\":\"structIAllocationManagerTypes.AllocateParams[]\",\"components\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"newMagnitudes\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"permissionController\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerForOperatorSets\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIAllocationManagerTypes.RegisterParams\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeStrategiesFromOperatorSet\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAVSRegistrar\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"registrar\",\"type\":\"address\",\"internalType\":\"contractIAVSRegistrar\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAllocationDelay\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delay\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slashOperator\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIAllocationManagerTypes.SlashingParams\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"wadsToSlash\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAVSMetadataURI\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AVSMetadataURIUpdated\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"metadataURI\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AVSRegistrarSet\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"registrar\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIAVSRegistrar\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllocationDelaySet\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"delay\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"effectBlock\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllocationUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"magnitude\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"effectBlock\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EncumberedMagnitudeUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"encumberedMagnitude\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MaxMagnitudeUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"maxMagnitude\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorAddedToOperatorSet\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRemovedFromOperatorSet\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSetCreated\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSlashed\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategies\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"contractIStrategy[]\"},{\"name\":\"wadSlashed\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"description\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyAddedToOperatorSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyRemovedFromOperatorSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyMemberOfSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CurrentlyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Empty\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputArrayLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientMagnitude\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCaller\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNewPausedStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperatorSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPermissions\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSnapshotOrdering\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidWadToSlash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ModificationAlreadyPending\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotMemberOfSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyPauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnpauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorNotSlashable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OutOfBounds\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SameMagnitude\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategiesMustBeInAscendingOrder\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategyAlreadyInOperatorSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategyNotInOperatorSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UninitializedAllocationDelay\",\"inputs\":[]}]",
	Bin: "0x610120604052348015610010575f5ffd5b50604051615a79380380615a7983398101604081905261002f91610180565b82858383876001600160a01b03811661005b576040516339b190bb60e11b815260040160405180910390fd5b6001600160a01b0390811660805292831660a05263ffffffff91821660c0521660e052166101005261008b610095565b50505050506101e9565b5f54610100900460ff16156101005760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff9081161461014f575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b0381168114610165575f5ffd5b50565b805163ffffffff8116811461017b575f5ffd5b919050565b5f5f5f5f5f60a08688031215610194575f5ffd5b855161019f81610151565b60208701519095506101b081610151565b60408701519094506101c181610151565b92506101cf60608701610168565b91506101dd60808701610168565b90509295509295909350565b60805160a05160c05160e051610100516157ff61027a5f395f81816103f9015261359701525f81816105480152613b9701525f818161031e015281816120c601526127bd01525f81816106c701528181610c77015281816115fc01528181611c6301528181611ccd0152612a8401525f818161056f0152818161077201528181611d72015261320e01526157ff5ff3fe608060405234801561000f575f5ffd5b5060043610610281575f3560e01c80636e875dba11610156578063adc2e3d9116100ca578063cd6dc68711610084578063cd6dc6871461069c578063d3d96ff4146106af578063df5cf723146106c2578063f2fde38b146106e9578063f605ce08146106fc578063fabc1cbc1461070f575f5ffd5b8063adc2e3d91461060e578063b2447af714610621578063b66bd98914610634578063b9fbaed114610647578063ba1a84e514610676578063c221d8ae14610689575f5ffd5b80638ce648541161011b5780638ce64854146105915780638da5cb5b146105b157806394d7d00c146105c2578063952899ee146105d5578063a9333ec8146105e8578063a9821821146105fb575f5ffd5b80636e875dba14610515578063715018a61461052857806379ae50cd146105305780637bc1ef6114610543578063886f11951461056a575f5ffd5b80634657e26a116101f8578063595c6a67116101b2578063595c6a67146104875780635ac86ab71461048f5780635c975abb146104b2578063670d3ba2146104c45780636cfb4481146104d75780636e3492b514610502575f5ffd5b80634657e26a146103f45780634a10ffe51461041b5780634b5046ef1461043b57806350feea201461044e578063547afb871461046157806356c483e614610474575f5ffd5b80632981eb77116102495780632981eb77146103195780632bab2c4a14610355578063304c10cd1461037557806336352057146103a057806340120dab146103b35780634177a87c146103d4575f5ffd5b806310e1b9b814610285578063136439dd146102ae57806315fe5028146102c3578063260dc758146102e3578063261f84e014610306575b5f5ffd5b610298610293366004614736565b610722565b6040516102a5919061477d565b60405180910390f35b6102c16102bc3660046147b0565b61075d565b005b6102d66102d13660046147c7565b610832565b6040516102a59190614845565b6102f66102f1366004614857565b610949565b60405190151581526020016102a5565b6102c16103143660046148b1565b610980565b6103407f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff90911681526020016102a5565b610368610363366004614996565b610c23565b6040516102a59190614a4a565b6103886103833660046147c7565b610f10565b6040516001600160a01b0390911681526020016102a5565b6102c16103ae366004614aad565b610f3f565b6103c66103c1366004614aff565b611751565b6040516102a5929190614b8c565b6103e76103e2366004614857565b6118cc565b6040516102a59190614be9565b6103887f000000000000000000000000000000000000000000000000000000000000000081565b61042e610429366004614bfb565b6118f0565b6040516102a59190614c3e565b6102c1610449366004614c89565b611998565b6102c161045c366004614d09565b611a52565b61042e61046f366004614d67565b611bb0565b6102c1610482366004614da9565b611c58565b6102c1611d5d565b6102f661049d366004614ddc565b606654600160ff9092169190911b9081161490565b6066545b6040519081526020016102a5565b6102f66104d2366004614dfc565b611e0c565b6104ea6104e5366004614aff565b611e1d565b6040516001600160401b0390911681526020016102a5565b6102c1610510366004614e3d565b611e32565b6103e7610523366004614857565b612202565b6102c1612213565b6102d661053e3660046147c7565b612224565b6103407f000000000000000000000000000000000000000000000000000000000000000081565b6103887f000000000000000000000000000000000000000000000000000000000000000081565b6105a461059f366004614e6e565b6122fe565b6040516102a59190614eb1565b6033546001600160a01b0316610388565b61042e6105d0366004614ec3565b6123c3565b6102c16105e3366004614f1e565b6124af565b6104ea6105f6366004614aff565b612976565b6102c16106093660046150c7565b6129a5565b6102c161061c366004615145565b612a15565b6104b661062f366004614857565b612d64565b6102c1610642366004614d09565b612d86565b61065a6106553660046147c7565b612ee0565b60408051921515835263ffffffff9091166020830152016102a5565b6104b66106843660046147c7565b612f7a565b6103e7610697366004614dfc565b612f9a565b6102c16106aa366004615187565b612fcb565b6102c16106bd366004614aff565b6130e8565b6103887f000000000000000000000000000000000000000000000000000000000000000081565b6102c16106f73660046147c7565b613187565b6104ea61070a366004614aff565b613200565b6102c161071d3660046147b0565b61320c565b604080516060810182525f80825260208201819052918101829052906107518561074b86613322565b85613385565b925050505b9392505050565b60405163237dfb4760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906346fbf68e90602401602060405180830381865afa1580156107bf573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107e391906151b1565b61080057604051631d77d47760e21b815260040160405180910390fd5b60665481811681146108255760405163c61dca5d60e01b815260040160405180910390fd5b61082e826134f1565b5050565b6001600160a01b0381165f908152609d60205260408120606091906108569061352e565b90505f816001600160401b038111156108715761087161465a565b6040519080825280602002602001820160405280156108b557816020015b604080518082019091525f808252602082015281526020019060019003908161088f5790505b5090505f5b82811015610941576001600160a01b0385165f908152609d6020526040902061091c906108e79083613537565b604080518082019091525f80825260208201525060408051808201909152606082901c815263ffffffff909116602082015290565b82828151811061092e5761092e6151d0565b60209081029190910101526001016108ba565b509392505050565b60208082015182516001600160a01b03165f90815260989092526040822061097a9163ffffffff9081169061354216565b92915050565b8261098a81613559565b6109a75760405163932d94f760e01b815260040160405180910390fd5b5f5b82811015610c1c575f6040518060400160405280876001600160a01b031681526020018686858181106109de576109de6151d0565b90506020028101906109f091906151e4565b6109fe906020810190615202565b63ffffffff168152509050610a48816020015163ffffffff1660985f896001600160a01b03166001600160a01b031681526020019081526020015f2061360390919063ffffffff16565b610a6557604051631fb1705560e21b815260040160405180910390fd5b7f31629285ead2335ae0933f86ed2ae63321f7af77b4e6eaabc42c057880977e6c6040518060400160405280886001600160a01b03168152602001836020015163ffffffff16815250604051610abb919061521b565b60405180910390a15f610acd82613322565b90505f5b868685818110610ae357610ae36151d0565b9050602002810190610af591906151e4565b610b03906020810190615229565b9050811015610c1157610b79878786818110610b2157610b216151d0565b9050602002810190610b3391906151e4565b610b41906020810190615229565b83818110610b5157610b516151d0565b9050602002016020810190610b6691906147c7565b5f8481526099602052604090209061360e565b507f7ab260fe0af193db5f4986770d831bda4ea46099dc817e8b6716dcae8af8e88b83888887818110610bae57610bae6151d0565b9050602002810190610bc091906151e4565b610bce906020810190615229565b84818110610bde57610bde6151d0565b9050602002016020810190610bf391906147c7565b604051610c0192919061526e565b60405180910390a1600101610ad1565b5050506001016109a9565b5050505050565b606083516001600160401b03811115610c3e57610c3e61465a565b604051908082528060200260200182016040528015610c7157816020015b6060815260200190600190039081610c5c5790505b5090505f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f0e0e67686866040518363ffffffff1660e01b8152600401610cc3929190615294565b5f60405180830381865afa158015610cdd573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610d0491908101906152b8565b90505f5b8551811015610f06575f868281518110610d2457610d246151d0565b6020026020010151905085516001600160401b03811115610d4757610d4761465a565b604051908082528060200260200182016040528015610d70578160200160208202803683370190505b50848381518110610d8357610d836151d0565b60209081029190910101525f5b8651811015610efc575f878281518110610dac57610dac6151d0565b6020908102919091018101516001600160a01b038086165f90815260a1845260408082209284168252919093528220909250610de790613622565b9050806001600160401b03165f03610e00575050610ef4565b5f610e0c858d85610722565b90508863ffffffff16816040015163ffffffff1611158015610e3457505f8160200151600f0b125b15610e5657610e4a815f01518260200151613635565b6001600160401b031681525b80515f90610e71906001600160401b03908116908516613649565b9050610eb881898981518110610e8957610e896151d0565b60200260200101518781518110610ea257610ea26151d0565b602002602001015161365d90919063ffffffff16565b898881518110610eca57610eca6151d0565b60200260200101518681518110610ee357610ee36151d0565b602002602001018181525050505050505b600101610d90565b5050600101610d08565b5050949350505050565b6001600160a01b038082165f908152609760205260408120549091168015610f385780610756565b5090919050565b606654600190600290811603610f685760405163840a48d560e01b815260040160405180910390fd5b82610f7281613559565b610f8f5760405163932d94f760e01b815260040160405180910390fd5b5f6040518060400160405280866001600160a01b03168152602001856020016020810190610fbd9190615202565b63ffffffff16905290505f610fde610fd860208701876147c7565b83613671565b9050610fed6060860186615229565b9050610ffc6040870187615229565b90501461101c576040516343714afd60e01b815260040160405180910390fd5b60208083015183516001600160a01b03165f9081526098909252604090912061104e9163ffffffff9081169061354216565b61106b57604051631fb1705560e21b815260040160405180910390fd5b806110895760405163ebbff49760e01b815260040160405180910390fd5b5f6110976040870187615229565b90506001600160401b038111156110b0576110b061465a565b6040519080825280602002602001820160405280156110d9578160200160208202803683370190505b5090505f5b6110eb6040880188615229565b90508110156116e25780158061117e57506111096040880188615229565b6111146001846153d8565b818110611123576111236151d0565b905060200201602081019061113891906147c7565b6001600160a01b031661114e6040890189615229565b8381811061115e5761115e6151d0565b905060200201602081019061117391906147c7565b6001600160a01b0316115b61119b57604051639f1c805360e01b815260040160405180910390fd5b6111a86060880188615229565b828181106111b8576111b86151d0565b905060200201355f1080156111f85750670de0b6b3a76400006111de6060890189615229565b838181106111ee576111ee6151d0565b9050602002013511155b61121557604051631353603160e01b815260040160405180910390fd5b6112716112256040890189615229565b83818110611235576112356151d0565b905060200201602081019061124a91906147c7565b60995f61125688613322565b81526020019081526020015f206136e790919063ffffffff16565b61128e576040516331bc342760e11b815260040160405180910390fd5b5f806112e06112a060208b018b6147c7565b6112a988613322565b6112b660408d018d615229565b878181106112c6576112c66151d0565b90506020020160208101906112db91906147c7565b613385565b805191935091506001600160401b03165f036112fd5750506116da565b5f61133861130e60608c018c615229565b8681811061131e5761131e6151d0565b85516001600160401b031692602090910201359050613708565b83519091506113536001600160401b03808416908316613649565b868681518110611365576113656151d0565b60200260200101818152505081835f0181815161138291906153eb565b6001600160401b031690525083518290859061139f9083906153eb565b6001600160401b03169052506020840180518391906113bf9083906153eb565b6001600160401b031690525060208301515f600f9190910b12156114da575f6114226113ee60608e018e615229565b888181106113fe576113fe6151d0565b9050602002013585602001516114139061540a565b6001600160801b031690613708565b9050806001600160401b03168460200181815161143f919061542e565b600f0b9052507f1487af5418c47ee5ea45ef4a93398668120890774a9e13487e61e9dc3baf76dd61147360208e018e6147c7565b8a8e80604001906114849190615229565b8a818110611494576114946151d0565b90506020020160208101906114a991906147c7565b6114ba885f01518960200151613635565b88604001516040516114d095949392919061545b565b60405180910390a1505b61152c6114ea60208d018d6147c7565b6114f38a613322565b61150060408f018f615229565b89818110611510576115106151d0565b905060200201602081019061152591906147c7565b878761371e565b7f1487af5418c47ee5ea45ef4a93398668120890774a9e13487e61e9dc3baf76dd61155a60208d018d6147c7565b8961156860408f018f615229565b89818110611578576115786151d0565b905060200201602081019061158d91906147c7565b86516040516115a19493929190439061545b565b60405180910390a16115f26115b960208d018d6147c7565b6115c660408e018e615229565b888181106115d6576115d66151d0565b90506020020160208101906115eb91906147c7565b865161395e565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001663601bb36f61162e60208e018e6147c7565b61163b60408f018f615229565b8981811061164b5761164b6151d0565b905060200201602081019061166091906147c7565b875160405160e085901b6001600160e01b03191681526001600160a01b0393841660048201529290911660248301526001600160401b0380861660448401521660648201526084015f604051808303815f87803b1580156116bf575f5ffd5b505af11580156116d1573d5f5f3e3d5ffd5b50505050505050505b6001016110de565b507f80969ad29428d6797ee7aad084f9e4a42a82fc506dcd2ca3b6fb431f85ccebe561171160208801886147c7565b8461171f60408a018a615229565b8561172d60808d018d6154ac565b6040516117409796959493929190615516565b60405180910390a150505050505050565b6001600160a01b0382165f908152609d6020526040812060609182916117769061352e565b90505f816001600160401b038111156117915761179161465a565b6040519080825280602002602001820160405280156117d557816020015b604080518082019091525f80825260208201528152602001906001900390816117af5790505b5090505f826001600160401b038111156117f1576117f161465a565b60405190808252806020026020018201604052801561183a57816020015b604080516060810182525f80825260208083018290529282015282525f1990920191018161180f5790505b5090505f5b838110156118bd576001600160a01b0388165f908152609d6020526040812061186c906108e79084613537565b905080848381518110611881576118816151d0565b602002602001018190525061189789828a610722565b8383815181106118a9576118a96151d0565b60209081029190910101525060010161183f565b509093509150505b9250929050565b60605f61075660995f6118de86613322565b81526020019081526020015f206139e0565b60605f83516001600160401b0381111561190c5761190c61465a565b604051908082528060200260200182016040528015611935578160200160208202803683370190505b5090505f5b845181101561094157611966858281518110611958576119586151d0565b602002602001015185612976565b828281518110611978576119786151d0565b6001600160401b039092166020928302919091019091015260010161193a565b6066545f906001908116036119c05760405163840a48d560e01b815260040160405180910390fd5b8382146119e0576040516343714afd60e01b815260040160405180910390fd5b5f5b84811015611a4957611a4187878784818110611a0057611a006151d0565b9050602002016020810190611a1591906147c7565b868685818110611a2757611a276151d0565b9050602002016020810190611a3c91906155ac565b6139ec565b6001016119e2565b50505050505050565b83611a5c81613559565b611a795760405163932d94f760e01b815260040160405180910390fd5b604080518082019091526001600160a01b038616815263ffffffff851660208201525f611aa582613322565b9050611ae6826020015163ffffffff1660985f8a6001600160a01b03166001600160a01b031681526020019081526020015f2061354290919063ffffffff16565b611b0357604051631fb1705560e21b815260040160405180910390fd5b5f5b84811015611ba657611b22868683818110610b5157610b516151d0565b611b3f5760405163585cfb2f60e01b815260040160405180910390fd5b7f7ab260fe0af193db5f4986770d831bda4ea46099dc817e8b6716dcae8af8e88b83878784818110611b7357611b736151d0565b9050602002016020810190611b8891906147c7565b604051611b9692919061526e565b60405180910390a1600101611b05565b5050505050505050565b60605f82516001600160401b03811115611bcc57611bcc61465a565b604051908082528060200260200182016040528015611bf5578160200160208202803683370190505b5090505f5b835181101561094157611c2685858381518110611c1957611c196151d0565b6020026020010151612976565b828281518110611c3857611c386151d0565b6001600160401b0390921660209283029190910190910152600101611bfa565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614611d5357611c9182613559565b611cae576040516348f5c3ed60e01b815260040160405180910390fd5b6040516336b87bd760e11b81526001600160a01b0383811660048301527f00000000000000000000000000000000000000000000000000000000000000001690636d70f7ae90602401602060405180830381865afa158015611d12573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611d3691906151b1565b611d535760405163ccea9e6f60e01b815260040160405180910390fd5b61082e8282613af0565b60405163237dfb4760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906346fbf68e90602401602060405180830381865afa158015611dbf573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611de391906151b1565b611e0057604051631d77d47760e21b815260040160405180910390fd5b611e0a5f196134f1565b565b5f61075683609a5f61125686613322565b5f5f611e298484613c9c565b95945050505050565b606654600290600490811603611e5b5760405163840a48d560e01b815260040160405180910390fd5b611e70611e6b60208401846147c7565b613559565b80611e895750611e89611e6b60408401602085016147c7565b611ea6576040516348f5c3ed60e01b815260040160405180910390fd5b5f5b611eb56040840184615229565b9050811015612177575f6040518060400160405280856020016020810190611edd91906147c7565b6001600160a01b03168152602001611ef86040870187615229565b85818110611f0857611f086151d0565b9050602002016020810190611f1d9190615202565b63ffffffff168152509050611f6a816020015163ffffffff1660985f876020016020810190611f4c91906147c7565b6001600160a01b0316815260208101919091526040015f2090613542565b611f8757604051631fb1705560e21b815260040160405180910390fd5b609e5f611f9760208701876147c7565b6001600160a01b03166001600160a01b031681526020019081526020015f205f611fc083613322565b815260208101919091526040015f205460ff16611ff0576040516325131d4f60e01b815260040160405180910390fd5b61202a611ffc82613322565b609c5f61200c60208901896147c7565b6001600160a01b0316815260208101919091526040015f2090613e0b565b5061206261203b60208601866147c7565b609a5f61204785613322565b81526020019081526020015f20613e1690919063ffffffff16565b5061207060208501856147c7565b6001600160a01b03167fad34c3070be1dffbcaa499d000ba2b8d9848aefcac3059df245dd95c4ece14fe826040516120a8919061521b565b60405180910390a2604080518082019091525f8152602081016120eb7f0000000000000000000000000000000000000000000000000000000000000000436155cd565b63ffffffff169052609e5f61210360208801886147c7565b6001600160a01b03166001600160a01b031681526020019081526020015f205f61212c84613322565b81526020808201929092526040015f2082518154939092015163ffffffff166101000264ffffffff00199215159290921664ffffffffff199093169290921717905550600101611ea8565b5061218b61038360408401602085016147c7565b6001600160a01b0316639d8e0c236121a660208501856147c7565b6121b36040860186615229565b6040518463ffffffff1660e01b81526004016121d193929190615622565b5f604051808303815f87803b1580156121e8575f5ffd5b505af19250505080156121f9575060015b1561082e575050565b606061097a609a5f6118de85613322565b61221b613e2a565b611e0a5f613e84565b6001600160a01b0381165f908152609c60205260408120606091906122489061352e565b90505f816001600160401b038111156122635761226361465a565b6040519080825280602002602001820160405280156122a757816020015b604080518082019091525f80825260208201528152602001906001900390816122815790505b5090505f5b82811015610941576001600160a01b0385165f908152609c602052604090206122d9906108e79083613537565b8282815181106122eb576122eb6151d0565b60209081029190910101526001016122ac565b60605f84516001600160401b0381111561231a5761231a61465a565b60405190808252806020026020018201604052801561236357816020015b604080516060810182525f80825260208083018290529282015282525f199092019101816123385790505b5090505f5b85518110156123ba57612395868281518110612386576123866151d0565b60200260200101518686610722565b8282815181106123a7576123a76151d0565b6020908102919091010152600101612368565b50949350505050565b60605f83516001600160401b038111156123df576123df61465a565b604051908082528060200260200182016040528015612408578160200160208202803683370190505b5090505f5b84518110156123ba576001600160a01b0386165f90815260a160205260408120865161247d92879291899086908110612448576124486151d0565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020015f20613ed590919063ffffffff16565b82828151811061248f5761248f6151d0565b6001600160401b039092166020928302919091019091015260010161240d565b6066545f906001908116036124d75760405163840a48d560e01b815260040160405180910390fd5b6124e083613559565b6124fd576040516348f5c3ed60e01b815260040160405180910390fd5b5f5f5f61250986612ee0565b915091508161252b5760405163fa55fc8160e01b815260040160405180910390fd5b91505f90505b8351811015610c1c5783818151811061254c5761254c6151d0565b6020026020010151604001515184828151811061256b5761256b6151d0565b6020026020010151602001515114612596576040516343714afd60e01b815260040160405180910390fd5b5f8482815181106125a9576125a96151d0565b602090810291909101810151518082015181516001600160a01b03165f908152609890935260409092209092506125e99163ffffffff9081169061354216565b61260657604051631fb1705560e21b815260040160405180910390fd5b5f6126118783613671565b90505f5b868481518110612627576126276151d0565b6020026020010151602001515181101561296b575f87858151811061264e5761264e6151d0565b602002602001015160200151828151811061266b5761266b6151d0565b60200260200101519050612682898261ffff6139ec565b5f5f6126918b61074b88613322565b915091508060200151600f0b5f146126bc57604051630d8fcbe360e41b815260040160405180910390fd5b5f6126c987858489613ee9565b905061270e825f01518c8a815181106126e4576126e46151d0565b6020026020010151604001518781518110612701576127016151d0565b6020026020010151613f1f565b600f0b602083018190525f0361273757604051634606179360e11b815260040160405180910390fd5b5f8260200151600f0b121561287b5780156127fd576127b861275888613322565b6001600160a01b03808f165f90815260a360209081526040808320938a16835292905220908154600160801b90819004600f0b5f818152600180860160205260409091209390935583546001600160801b03908116939091011602179055565b6127e27f0000000000000000000000000000000000000000000000000000000000000000436155cd565b6127ed9060016155cd565b63ffffffff1660408301526128e8565b61280f83602001518360200151613635565b6001600160401b031660208401528a518b9089908110612831576128316151d0565b602002602001015160400151858151811061284e5761284e6151d0565b6020908102919091018101516001600160401b031683525f9083015263ffffffff431660408301526128e8565b5f8260200151600f0b13156128e85761289c83602001518360200151613635565b6001600160401b0390811660208501819052845190911610156128d257604051636c9be0bf60e01b815260040160405180910390fd5b6128dc89436155cd565b63ffffffff1660408301525b6128fd8c6128f589613322565b86868661371e565b7f1487af5418c47ee5ea45ef4a93398668120890774a9e13487e61e9dc3baf76dd8c61292b6108e78a613322565b8661293d865f01518760200151613635565b866040015160405161295395949392919061545b565b60405180910390a15050600190920191506126159050565b505050600101612531565b6001600160a01b038083165f90815260a160209081526040808320938516835292905290812061075690613622565b826129af81613559565b6129cc5760405163932d94f760e01b815260040160405180910390fd5b836001600160a01b03167fa89c1dc243d8908a96dd84944bcc97d6bc6ac00dd78e20621576be6a3c9437138484604051612a07929190615646565b60405180910390a250505050565b606654600290600490811603612a3e5760405163840a48d560e01b815260040160405180910390fd5b82612a4881613559565b612a655760405163932d94f760e01b815260040160405180910390fd5b6040516336b87bd760e11b81526001600160a01b0385811660048301527f00000000000000000000000000000000000000000000000000000000000000001690636d70f7ae90602401602060405180830381865afa158015612ac9573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612aed91906151b1565b612b0a5760405163ccea9e6f60e01b815260040160405180910390fd5b5f5b612b196020850185615229565b9050811015612ce157604080518082019091525f9080612b3c60208801886147c7565b6001600160a01b03168152602001868060200190612b5a9190615229565b85818110612b6a57612b6a6151d0565b9050602002016020810190612b7f9190615202565b63ffffffff90811690915260208083015183516001600160a01b03165f90815260989092526040909120929350612bbb92919081169061354216565b612bd857604051631fb1705560e21b815260040160405180910390fd5b612be28682613671565b15612c0057604051636c6c6e2760e11b815260040160405180910390fd5b612c29612c0c82613322565b6001600160a01b0388165f908152609c6020526040902090613603565b50612c5586609a5f612c3a85613322565b81526020019081526020015f2061360e90919063ffffffff16565b50856001600160a01b03167f43232edf9071753d2321e5fa7e018363ee248e5f2142e6c08edd3265bfb4895e82604051612c8f919061521b565b60405180910390a26001600160a01b0386165f908152609e60205260408120600191612cba84613322565b815260208101919091526040015f20805460ff191691151591909117905550600101612b0c565b50612cf261038360208501856147c7565b6001600160a01b031663adcf73f785612d0e6020870187615229565b612d1b60408901896154ac565b6040518663ffffffff1660e01b8152600401612d3b959493929190615659565b5f604051808303815f87803b158015612d52575f5ffd5b505af1158015611ba6573d5f5f3e3d5ffd5b5f61097a609a5f612d7485613322565b81526020019081526020015f2061352e565b83612d9081613559565b612dad5760405163932d94f760e01b815260040160405180910390fd5b6040805180820182526001600160a01b03871680825263ffffffff80881660208085018290525f93845260989052939091209192612dec929161354216565b612e0957604051631fb1705560e21b815260040160405180910390fd5b5f612e1382613322565b90505f5b84811015611ba657612e5c868683818110612e3457612e346151d0565b9050602002016020810190612e4991906147c7565b5f84815260996020526040902090613e16565b612e79576040516331bc342760e11b815260040160405180910390fd5b7f7b4b073d80dcac55a11177d8459ad9f664ceeb91f71f27167bb14f8152a7eeee83878784818110612ead57612ead6151d0565b9050602002016020810190612ec291906147c7565b604051612ed092919061526e565b60405180910390a1600101612e17565b6001600160a01b0381165f908152609b602090815260408083208151608081018352905463ffffffff80821680845260ff600160201b8404161515958401869052650100000000008304821694840194909452600160481b909104166060820181905284939192919015801590612f615750826060015163ffffffff164310155b15612f70575050604081015160015b9590945092505050565b6001600160a01b0381165f90815260986020526040812061097a9061352e565b6001600160a01b0382165f908152609f6020526040812060609190612fc390826118de86613322565b949350505050565b5f54610100900460ff1615808015612fe957505f54600160ff909116105b806130025750303b15801561300257505f5460ff166001145b61306a5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b5f805460ff19166001179055801561308b575f805461ff0019166101001790555b613094826134f1565b61309d83613e84565b80156130e3575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498906020015b60405180910390a15b505050565b816130f281613559565b61310f5760405163932d94f760e01b815260040160405180910390fd5b6001600160a01b038381165f90815260976020526040902080546001600160a01b0319169184169190911790557f2ae945c40c44dc0ec263f95609c3fdc6952e0aefa22d6374e44f2c997acedf858361316781610f10565b604080516001600160a01b039384168152929091166020830152016130da565b61318f613e2a565b6001600160a01b0381166131f45760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401613061565b6131fd81613e84565b50565b5f5f6123ba8484613c9c565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015613268573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061328c919061569c565b6001600160a01b0316336001600160a01b0316146132bd5760405163794821ff60e01b815260040160405180910390fd5b606654801982198116146132e45760405163c61dca5d60e01b815260040160405180910390fd5b606682905560405182815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c9060200160405180910390a25050565b5f815f0151826020015163ffffffff1660405160200161336d92919060609290921b6bffffffffffffffffffffffff1916825260a01b6001600160a01b031916601482015260200190565b60405160208183030381529060405261097a906156b7565b6040805180820182525f80825260208083018290528351606081018552828152808201839052808501839052845180860186526001600160a01b03898116855260a18452868520908816855290925293822092939281906133e590613622565b6001600160401b0390811682526001600160a01b038981165f81815260a260209081526040808320948c168084529482528083205486169682019690965291815260a082528481208b8252825284812092815291815290839020835160608101855290549283168152600160401b8304600f0b91810191909152600160c01b90910463ffffffff169181018290529192504310156134875790925090506134e9565b613498815f01518260200151613635565b6001600160401b0316815260208101515f600f9190910b12156134d6576134c782602001518260200151613635565b6001600160401b031660208301525b5f60408201819052602082015290925090505b935093915050565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a250565b5f61097a825490565b5f6107568383613f36565b5f8181526001830160205260408120541515610756565b604051631beb2b9760e31b81526001600160a01b0382811660048301523360248301523060448301525f80356001600160e01b0319166064840152917f00000000000000000000000000000000000000000000000000000000000000009091169063df595cb8906084016020604051808303815f875af11580156135df573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061097a91906151b1565b5f6107568383613f5c565b5f610756836001600160a01b038416613f5c565b5f61097a82670de0b6b3a7640000613fa8565b5f610756826001600160401b03851661542e565b5f61075683670de0b6b3a764000084613fec565b5f6107568383670de0b6b3a7640000613fec565b6001600160a01b0382165f908152609e6020526040812081908161369485613322565b815260208082019290925260409081015f2081518083019092525460ff8116151580835261010090910463ffffffff1692820192909252915080612fc357506020015163ffffffff164311159392505050565b6001600160a01b0381165f9081526001830160205260408120541515610756565b5f6107568383670de0b6b3a764000060016140d1565b6020808301516001600160a01b038088165f90815260a284526040808220928816825291909352909120546001600160401b039081169116146137e457602082810180516001600160a01b038881165f81815260a286526040808220938a1680835293875290819020805467ffffffffffffffff19166001600160401b0395861617905593518451918252948101919091529216908201527facf9095feb3a370c9cf692421c69ef320d4db5c66e6a7d29c7694eb02364fc559060600160405180910390a15b6001600160a01b038086165f90815260a060209081526040808320888452825280832093871683529281529082902083518154928501519385015163ffffffff16600160c01b0263ffffffff60c01b196001600160801b038616600160401b026001600160c01b03199095166001600160401b03909316929092179390931716919091179055600f0b156138c6576001600160a01b0385165f908152609f60209081526040808320878452909152902061389e908461360e565b506001600160a01b0385165f908152609d602052604090206138c09085613603565b50610c1c565b80516001600160401b03165f03610c1c576001600160a01b0385165f908152609f6020908152604080832087845290915290206139039084613e16565b506001600160a01b0385165f908152609f60209081526040808320878452909152902061392f9061352e565b5f03610c1c576001600160a01b0385165f908152609d602052604090206139569085613e0b565b505050505050565b6001600160a01b038084165f90815260a16020908152604080832093861683529290522061398d90438361412a565b604080516001600160a01b038086168252841660208201526001600160401b038316918101919091527f1c6458079a41077d003c11faf9bf097e693bd67979e4e6500bac7b29db779b5c906060016130da565b60605f6107568361413e565b6001600160a01b038381165f90815260a360209081526040808320938616835292905290812054600f81810b600160801b909204900b035b5f81118015613a3657508261ffff1682105b15610c1c576001600160a01b038086165f90815260a3602090815260408083209388168352929052908120613a6a90614197565b90505f5f613a79888489613385565b91509150806040015163ffffffff16431015613a9757505050610c1c565b613aa4888489858561371e565b6001600160a01b038089165f90815260a360209081526040808320938b16835292905220613ad1906141e9565b50613adb856156da565b9450613ae6846156f2565b9350505050613a24565b6001600160a01b0382165f908152609b60209081526040918290208251608081018452905463ffffffff808216835260ff600160201b830416151593830193909352650100000000008104831693820193909352600160481b909204166060820181905215801590613b6c5750806060015163ffffffff164310155b15613b8657604081015163ffffffff168152600160208201525b63ffffffff82166040820152613bbc7f0000000000000000000000000000000000000000000000000000000000000000436155cd565b613bc79060016155cd565b63ffffffff90811660608381019182526001600160a01b0386165f818152609b602090815260409182902087518154838a0151858b01519851928a1664ffffffffff1990921691909117600160201b91151591909102176cffffffffffffffff0000000000191665010000000000978916979097026cffffffff000000000000000000191696909617600160481b968816968702179055815192835294871694820194909452928301919091527f4e85751d6331506c6c62335f207eb31f12a61e570f34f5c17640308785c6d4db91016130da565b6001600160a01b038281165f81815260a2602090815260408083209486168084529482528083205493835260a38252808320948352939052918220546001600160401b039091169190600f81810b600160801b909204900b03815b81811015613dc7576001600160a01b038087165f90815260a3602090815260408083209389168352929052908120613d2f9083614266565b6001600160a01b038881165f90815260a0602090815260408083208584528252808320938b16835292815290829020825160608101845290546001600160401b0381168252600160401b8104600f0b92820192909252600160c01b90910463ffffffff16918101829052919250431015613daa575050613dc7565b613db8868260200151613635565b95505050806001019050613cf7565b506001600160a01b038086165f90815260a1602090815260408083209388168352929052208390613df790613622565b613e0191906153eb565b9150509250929050565b5f61075683836142d5565b5f610756836001600160a01b0384166142d5565b6033546001600160a01b03163314611e0a5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401613061565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b5f6107568383670de0b6b3a76400006143b8565b5f613efa8460995f61125689613322565b8015613f035750815b8015611e2957505090516001600160401b031615159392505050565b5f6107566001600160401b03808516908416615707565b5f825f018281548110613f4b57613f4b6151d0565b905f5260205f200154905092915050565b5f818152600183016020526040812054613fa157508154600181810184555f84815260208082209093018490558454848252828601909352604090209190915561097a565b505f61097a565b81545f908015613fe457613fce84613fc16001846153d8565b5f91825260209091200190565b54600160201b90046001600160e01b0316612fc3565b509092915050565b5f80805f19858709858702925082811083820303915050805f036140235783828161401957614019615734565b0492505050610756565b80841161406a5760405162461bcd60e51b81526020600482015260156024820152744d6174683a206d756c446976206f766572666c6f7760581b6044820152606401613061565b5f8486880960026001871981018816978890046003810283188082028403028082028403028082028403028082028403028082028403029081029092039091025f889003889004909101858311909403939093029303949094049190911702949350505050565b5f5f6140de868686613fec565b905060018360028111156140f4576140f4615748565b14801561411057505f848061410b5761410b615734565b868809115b15611e295761412060018261575c565b9695505050505050565b6130e383836001600160401b038416614400565b6060815f0180548060200260200160405190810160405280929190818152602001828054801561418b57602002820191905f5260205f20905b815481526020019060010190808311614177575b50505050509050919050565b5f6141b18254600f81810b600160801b909204900b131590565b156141cf57604051631ed9509560e11b815260040160405180910390fd5b508054600f0b5f9081526001909101602052604090205490565b5f6142038254600f81810b600160801b909204900b131590565b1561422157604051631ed9509560e11b815260040160405180910390fd5b508054600f0b5f818152600180840160205260408220805492905583546fffffffffffffffffffffffffffffffff191692016001600160801b03169190911790915590565b5f5f61428861427484614503565b85546142839190600f0b61576f565b614570565b8454909150600160801b9004600f90810b9082900b126142bb57604051632d0483c560e21b815260040160405180910390fd5b600f0b5f9081526001939093016020525050604090205490565b5f81815260018301602052604081205480156143af575f6142f76001836153d8565b85549091505f9061430a906001906153d8565b9050818114614369575f865f018281548110614328576143286151d0565b905f5260205f200154905080875f018481548110614348576143486151d0565b5f918252602080832090910192909255918252600188019052604090208390555b855486908061437a5761437a615796565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f90556001935050505061097a565b5f91505061097a565b82545f90816143c9868683856145d9565b905080156143f6576143e086613fc16001846153d8565b54600160201b90046001600160e01b0316610751565b5091949350505050565b825480156144b6575f61441885613fc16001856153d8565b60408051808201909152905463ffffffff808216808452600160201b9092046001600160e01b03166020840152919250908516101561446a5760405163151b8e3f60e11b815260040160405180910390fd5b805163ffffffff8086169116036144b4578261448b86613fc16001866153d8565b80546001600160e01b0392909216600160201b0263ffffffff9092169190911790555050505050565b505b506040805180820190915263ffffffff92831681526001600160e01b03918216602080830191825285546001810187555f968752952091519051909216600160201b029190921617910155565b5f6001600160ff1b0382111561456c5760405162461bcd60e51b815260206004820152602860248201527f53616665436173743a2076616c756520646f65736e27742066697420696e2061604482015267371034b73a191a9b60c11b6064820152608401613061565b5090565b80600f81900b81146145d45760405162461bcd60e51b815260206004820152602760248201527f53616665436173743a2076616c756520646f65736e27742066697420696e20316044820152663238206269747360c81b6064820152608401613061565b919050565b5f5b81831015610941575f6145ee848461462c565b5f8781526020902090915063ffffffff86169082015463ffffffff16111561461857809250614626565b61462381600161575c565b93505b506145db565b5f61463a60028484186157aa565b6107569084841661575c565b6001600160a01b03811681146131fd575f5ffd5b634e487b7160e01b5f52604160045260245ffd5b604051606081016001600160401b03811182821017156146905761469061465a565b60405290565b604051601f8201601f191681016001600160401b03811182821017156146be576146be61465a565b604052919050565b803563ffffffff811681146145d4575f5ffd5b5f604082840312156146e9575f5ffd5b604080519081016001600160401b038111828210171561470b5761470b61465a565b604052905080823561471c81614646565b815261472a602084016146c6565b60208201525092915050565b5f5f5f60808486031215614748575f5ffd5b833561475381614646565b925061476285602086016146d9565b9150606084013561477281614646565b809150509250925092565b81516001600160401b03168152602080830151600f0b9082015260408083015163ffffffff16908201526060810161097a565b5f602082840312156147c0575f5ffd5b5035919050565b5f602082840312156147d7575f5ffd5b813561075681614646565b80516001600160a01b0316825260209081015163ffffffff16910152565b5f8151808452602084019350602083015f5b8281101561483b576148258683516147e2565b6040959095019460209190910190600101614812565b5093949350505050565b602081525f6107566020830184614800565b5f60408284031215614867575f5ffd5b61075683836146d9565b5f5f83601f840112614881575f5ffd5b5081356001600160401b03811115614897575f5ffd5b6020830191508360208260051b85010111156118c5575f5ffd5b5f5f5f604084860312156148c3575f5ffd5b83356148ce81614646565b925060208401356001600160401b038111156148e8575f5ffd5b6148f486828701614871565b9497909650939450505050565b5f6001600160401b038211156149195761491961465a565b5060051b60200190565b5f82601f830112614932575f5ffd5b813561494561494082614901565b614696565b8082825260208201915060208360051b860101925085831115614966575f5ffd5b602085015b8381101561498c57803561497e81614646565b83526020928301920161496b565b5095945050505050565b5f5f5f5f60a085870312156149a9575f5ffd5b6149b386866146d9565b935060408501356001600160401b038111156149cd575f5ffd5b6149d987828801614923565b93505060608501356001600160401b038111156149f4575f5ffd5b614a0087828801614923565b925050614a0f608086016146c6565b905092959194509250565b5f8151808452602084019350602083015f5b8281101561483b578151865260209586019590910190600101614a2c565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b82811015614aa157603f19878603018452614a8c858351614a1a565b94506020938401939190910190600101614a70565b50929695505050505050565b5f5f60408385031215614abe575f5ffd5b8235614ac981614646565b915060208301356001600160401b03811115614ae3575f5ffd5b830160a08186031215614af4575f5ffd5b809150509250929050565b5f5f60408385031215614b10575f5ffd5b8235614b1b81614646565b91506020830135614af481614646565b5f8151808452602084019350602083015f5b8281101561483b57614b7686835180516001600160401b03168252602080820151600f0b9083015260409081015163ffffffff16910152565b6060959095019460209190910190600101614b3d565b604081525f614b9e6040830185614800565b8281036020840152611e298185614b2b565b5f8151808452602084019350602083015f5b8281101561483b5781516001600160a01b0316865260209586019590910190600101614bc2565b602081525f6107566020830184614bb0565b5f5f60408385031215614c0c575f5ffd5b82356001600160401b03811115614c21575f5ffd5b614c2d85828601614923565b9250506020830135614af481614646565b602080825282518282018190525f918401906040840190835b81811015614c7e5783516001600160401b0316835260209384019390920191600101614c57565b509095945050505050565b5f5f5f5f5f60608688031215614c9d575f5ffd5b8535614ca881614646565b945060208601356001600160401b03811115614cc2575f5ffd5b614cce88828901614871565b90955093505060408601356001600160401b03811115614cec575f5ffd5b614cf888828901614871565b969995985093965092949392505050565b5f5f5f5f60608587031215614d1c575f5ffd5b8435614d2781614646565b9350614d35602086016146c6565b925060408501356001600160401b03811115614d4f575f5ffd5b614d5b87828801614871565b95989497509550505050565b5f5f60408385031215614d78575f5ffd5b8235614d8381614646565b915060208301356001600160401b03811115614d9d575f5ffd5b613e0185828601614923565b5f5f60408385031215614dba575f5ffd5b8235614dc581614646565b9150614dd3602084016146c6565b90509250929050565b5f60208284031215614dec575f5ffd5b813560ff81168114610756575f5ffd5b5f5f60608385031215614e0d575f5ffd5b8235614e1881614646565b9150614dd384602085016146d9565b5f60608284031215614e37575f5ffd5b50919050565b5f60208284031215614e4d575f5ffd5b81356001600160401b03811115614e62575f5ffd5b612fc384828501614e27565b5f5f5f60808486031215614e80575f5ffd5b83356001600160401b03811115614e95575f5ffd5b614ea186828701614923565b93505061476285602086016146d9565b602081525f6107566020830184614b2b565b5f5f5f60608486031215614ed5575f5ffd5b8335614ee081614646565b925060208401356001600160401b03811115614efa575f5ffd5b614f0686828701614923565b925050614f15604085016146c6565b90509250925092565b5f5f60408385031215614f2f575f5ffd5b8235614f3a81614646565b915060208301356001600160401b03811115614f54575f5ffd5b8301601f81018513614f64575f5ffd5b8035614f7261494082614901565b8082825260208201915060208360051b850101925087831115614f93575f5ffd5b602084015b838110156150b85780356001600160401b03811115614fb5575f5ffd5b85016080818b03601f19011215614fca575f5ffd5b614fd261466e565b614fdf8b602084016146d9565b815260608201356001600160401b03811115614ff9575f5ffd5b6150088c602083860101614923565b60208301525060808201356001600160401b03811115615026575f5ffd5b6020818401019250508a601f83011261503d575f5ffd5b813561504b61494082614901565b8082825260208201915060208360051b86010192508d83111561506c575f5ffd5b6020850194505b828510156150a25784356001600160401b0381168114615091575f5ffd5b825260209485019490910190615073565b6040840152505084525060209283019201614f98565b50809450505050509250929050565b5f5f5f604084860312156150d9575f5ffd5b83356150e481614646565b925060208401356001600160401b038111156150fe575f5ffd5b8401601f8101861361510e575f5ffd5b80356001600160401b03811115615123575f5ffd5b866020828401011115615134575f5ffd5b939660209190910195509293505050565b5f5f60408385031215615156575f5ffd5b823561516181614646565b915060208301356001600160401b0381111561517b575f5ffd5b613e0185828601614e27565b5f5f60408385031215615198575f5ffd5b82356151a381614646565b946020939093013593505050565b5f602082840312156151c1575f5ffd5b81518015158114610756575f5ffd5b634e487b7160e01b5f52603260045260245ffd5b5f8235603e198336030181126151f8575f5ffd5b9190910192915050565b5f60208284031215615212575f5ffd5b610756826146c6565b6040810161097a82846147e2565b5f5f8335601e1984360301811261523e575f5ffd5b8301803591506001600160401b03821115615257575f5ffd5b6020019150600581901b36038213156118c5575f5ffd5b6060810161527c82856147e2565b6001600160a01b039290921660409190910152919050565b604081525f6152a66040830185614bb0565b8281036020840152611e298185614bb0565b5f602082840312156152c8575f5ffd5b81516001600160401b038111156152dd575f5ffd5b8201601f810184136152ed575f5ffd5b80516152fb61494082614901565b8082825260208201915060208360051b85010192508683111561531c575f5ffd5b602084015b838110156153b95780516001600160401b0381111561533e575f5ffd5b8501603f8101891361534e575f5ffd5b602081015161535f61494082614901565b808282526020820191506020808460051b8601010192508b831115615382575f5ffd5b6040840193505b828410156153a4578351825260209384019390910190615389565b86525050602093840193919091019050615321565b509695505050505050565b634e487b7160e01b5f52601160045260245ffd5b8181038181111561097a5761097a6153c4565b6001600160401b03828116828216039081111561097a5761097a6153c4565b5f81600f0b60016001607f1b03198103615426576154266153c4565b5f0392915050565b600f81810b9083900b0160016001607f1b03811360016001607f1b03198212171561097a5761097a6153c4565b6001600160a01b038616815260c0810161547860208301876147e2565b6001600160a01b039490941660608201526001600160401b0392909216608083015263ffffffff1660a09091015292915050565b5f5f8335601e198436030181126154c1575f5ffd5b8301803591506001600160401b038211156154da575f5ffd5b6020019150368190038213156118c5575f5ffd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b6001600160a01b03881681525f60c08201615534602084018a6147e2565b60c060608401528690528660e083015f5b8881101561557557823561555881614646565b6001600160a01b0316825260209283019290910190600101615545565b5083810360808501526155888188614a1a565b91505082810360a084015261559e8185876154ee565b9a9950505050505050505050565b5f602082840312156155bc575f5ffd5b813561ffff81168114610756575f5ffd5b63ffffffff818116838216019081111561097a5761097a6153c4565b8183526020830192505f815f5b8481101561483b5763ffffffff61560c836146c6565b16865260209586019591909101906001016155f6565b6001600160a01b03841681526040602082018190525f90611e2990830184866155e9565b602081525f612fc36020830184866154ee565b6001600160a01b03861681526060602082018190525f9061567d90830186886155e9565b82810360408401526156908185876154ee565b98975050505050505050565b5f602082840312156156ac575f5ffd5b815161075681614646565b80516020808301519190811015614e37575f1960209190910360031b1b16919050565b5f600182016156eb576156eb6153c4565b5060010190565b5f81615700576157006153c4565b505f190190565b600f82810b9082900b0360016001607f1b0319811260016001607f1b038213171561097a5761097a6153c4565b634e487b7160e01b5f52601260045260245ffd5b634e487b7160e01b5f52602160045260245ffd5b8082018082111561097a5761097a6153c4565b8082018281125f83128015821682158216171561578e5761578e6153c4565b505092915050565b634e487b7160e01b5f52603160045260245ffd5b5f826157c457634e487b7160e01b5f52601260045260245ffd5b50049056fea2646970667358221220bf2c2a135f852605c736e114c992eb283845849273282d707e128f09f189664c64736f6c634300081b0033",
>>>>>>> 97b9d82b (feat: changing burnableShares to EnumerableMap (#1028))
=======
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_delegation\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"_permissionController\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"},{\"name\":\"_DEALLOCATION_DELAY\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_ALLOCATION_CONFIGURATION_DELAY\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ALLOCATION_CONFIGURATION_DELAY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEALLOCATION_DELAY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addStrategiesToOperatorSet\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"clearDeallocationQueue\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"numToClear\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createOperatorSets\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"params\",\"type\":\"tuple[]\",\"internalType\":\"structIAllocationManagerTypes.CreateSetParams[]\",\"components\":[{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregisterFromOperatorSets\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIAllocationManagerTypes.DeregisterParams\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAVSRegistrar\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAVSRegistrar\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocatableMagnitude\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocatedSets\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocatedStake\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[][]\",\"internalType\":\"uint256[][]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocatedStrategies\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocation\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIAllocationManagerTypes.Allocation\",\"components\":[{\"name\":\"currentMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"pendingDiff\",\"type\":\"int128\",\"internalType\":\"int128\"},{\"name\":\"effectBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocationDelay\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocations\",\"inputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIAllocationManagerTypes.Allocation[]\",\"components\":[{\"name\":\"currentMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"pendingDiff\",\"type\":\"int128\",\"internalType\":\"int128\"},{\"name\":\"effectBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getEncumberedMagnitude\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxMagnitude\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxMagnitudes\",\"inputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxMagnitudes\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxMagnitudesAtBlock\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMemberCount\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMembers\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMinimumSlashableStake\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"futureBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"slashableStake\",\"type\":\"uint256[][]\",\"internalType\":\"uint256[][]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorSetCount\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRegisteredSets\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStrategiesInOperatorSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStrategyAllocations\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIAllocationManagerTypes.Allocation[]\",\"components\":[{\"name\":\"currentMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"pendingDiff\",\"type\":\"int128\",\"internalType\":\"int128\"},{\"name\":\"effectBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isMemberOfOperatorSet\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperatorSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperatorSlashable\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"modifyAllocations\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"params\",\"type\":\"tuple[]\",\"internalType\":\"structIAllocationManagerTypes.AllocateParams[]\",\"components\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"newMagnitudes\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"permissionController\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerForOperatorSets\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIAllocationManagerTypes.RegisterParams\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeStrategiesFromOperatorSet\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAVSRegistrar\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"registrar\",\"type\":\"address\",\"internalType\":\"contractIAVSRegistrar\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAllocationDelay\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delay\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slashOperator\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIAllocationManagerTypes.SlashingParams\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"wadsToSlash\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAVSMetadataURI\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AVSMetadataURIUpdated\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"metadataURI\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AVSRegistrarSet\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"registrar\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIAVSRegistrar\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllocationDelaySet\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"delay\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"effectBlock\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllocationUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"magnitude\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"effectBlock\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EncumberedMagnitudeUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"encumberedMagnitude\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MaxMagnitudeUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"maxMagnitude\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorAddedToOperatorSet\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRemovedFromOperatorSet\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSetCreated\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSlashed\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategies\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"contractIStrategy[]\"},{\"name\":\"wadSlashed\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"description\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyAddedToOperatorSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyRemovedFromOperatorSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyMemberOfSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CurrentlyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Empty\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputArrayLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientMagnitude\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCaller\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNewPausedStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperatorSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPermissions\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSnapshotOrdering\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidWadToSlash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ModificationAlreadyPending\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotMemberOfSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyPauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnpauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorNotSlashable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OutOfBounds\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SameMagnitude\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategiesMustBeInAscendingOrder\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategyAlreadyInOperatorSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategyNotInOperatorSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UninitializedAllocationDelay\",\"inputs\":[]}]",
	Bin: "0x610120604052348015610010575f5ffd5b50604051615bc3380380615bc383398101604081905261002f91610180565b82858383876001600160a01b03811661005b576040516339b190bb60e11b815260040160405180910390fd5b6001600160a01b0390811660805292831660a05263ffffffff91821660c0521660e052166101005261008b610095565b50505050506101e9565b5f54610100900460ff16156101005760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff9081161461014f575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b0381168114610165575f5ffd5b50565b805163ffffffff8116811461017b575f5ffd5b919050565b5f5f5f5f5f60a08688031215610194575f5ffd5b855161019f81610151565b60208701519095506101b081610151565b60408701519094506101c181610151565b92506101cf60608701610168565b91506101dd60808701610168565b90509295509295909350565b60805160a05160c05160e0516101005161594961027a5f395f818161043501526133f701525f81816105840152613c4701525f818161034701528181611f37015261262501525f81816107030152818161146e01528181611ad401528181611b3e015281816128ec01526134d601525f81816105ab0152818161082901528181611be3015261306e01526159495ff3fe608060405234801561000f575f5ffd5b5060043610610297575f3560e01c80636e3492b511610161578063adc2e3d9116100ca578063cd6dc68711610084578063cd6dc687146106d8578063d3d96ff4146106eb578063df5cf723146106fe578063f2fde38b14610725578063f605ce0814610738578063fabc1cbc1461074b575f5ffd5b8063adc2e3d91461064a578063b2447af71461065d578063b66bd98914610670578063b9fbaed114610683578063ba1a84e5146106b2578063c221d8ae146106c5575f5ffd5b80638ce648541161011b5780638ce64854146105cd5780638da5cb5b146105ed57806394d7d00c146105fe578063952899ee14610611578063a9333ec814610624578063a982182114610637575f5ffd5b80636e3492b51461053e5780636e875dba14610551578063715018a61461056457806379ae50cd1461056c5780637bc1ef611461057f578063886f1195146105a6575f5ffd5b80634177a87c1161020357806356c483e6116101bd57806356c483e6146104b0578063595c6a67146104c35780635ac86ab7146104cb5780635c975abb146104ee578063670d3ba2146105005780636cfb448114610513575f5ffd5b80634177a87c146104105780634657e26a146104305780634a10ffe5146104575780634b5046ef1461047757806350feea201461048a578063547afb871461049d575f5ffd5b80632981eb77116102545780632981eb77146103425780632b453a9a1461037e5780632bab2c4a1461039e578063304c10cd146103b157806336352057146103dc57806340120dab146103ef575f5ffd5b806310e1b9b81461029b5780631352c3e6146102c4578063136439dd146102e757806315fe5028146102fc578063260dc7581461031c578063261f84e01461032f575b5f5ffd5b6102ae6102a936600461480d565b61075e565b6040516102bb9190614854565b60405180910390f35b6102d76102d2366004614887565b610799565b60405190151581526020016102bb565b6102fa6102f53660046148bb565b610814565b005b61030f61030a3660046148d2565b6108e9565b6040516102bb9190614950565b6102d761032a366004614962565b610a00565b6102fa61033d3660046149bc565b610a31565b6103697f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff90911681526020016102bb565b61039161038c366004614aa1565b610cd4565b6040516102bb9190614b44565b6103916103ac366004614ba7565b610cea565b6103c46103bf3660046148d2565b610d89565b6040516001600160a01b0390911681526020016102bb565b6102fa6103ea366004614c2b565b610db8565b6104026103fd366004614c7d565b6115c2565b6040516102bb929190614d0a565b61042361041e366004614962565b61173d565b6040516102bb9190614d67565b6103c47f000000000000000000000000000000000000000000000000000000000000000081565b61046a610465366004614d79565b611761565b6040516102bb9190614dbc565b6102fa610485366004614e07565b611809565b6102fa610498366004614e87565b6118c3565b61046a6104ab366004614ee5565b611a21565b6102fa6104be366004614f27565b611ac9565b6102fa611bce565b6102d76104d9366004614f51565b606654600160ff9092169190911b9081161490565b6066545b6040519081526020016102bb565b6102d761050e366004614887565b611c7d565b610526610521366004614c7d565b611c8e565b6040516001600160401b0390911681526020016102bb565b6102fa61054c366004614f87565b611ca3565b61042361055f366004614962565b612073565b6102fa612084565b61030f61057a3660046148d2565b612095565b6103697f000000000000000000000000000000000000000000000000000000000000000081565b6103c47f000000000000000000000000000000000000000000000000000000000000000081565b6105e06105db366004614fb8565b61216f565b6040516102bb9190614ffb565b6033546001600160a01b03166103c4565b61046a61060c36600461500d565b61222b565b6102fa61061f366004615068565b612317565b610526610632366004614c7d565b6127de565b6102fa610645366004615211565b61280d565b6102fa61065836600461528f565b61287d565b6104f261066b366004614962565b612bcc565b6102fa61067e366004614e87565b612bee565b6106966106913660046148d2565b612d48565b60408051921515835263ffffffff9091166020830152016102bb565b6104f26106c03660046148d2565b612de2565b6104236106d3366004614887565b612e02565b6102fa6106e63660046152d1565b612e2b565b6102fa6106f9366004614c7d565b612f48565b6103c47f000000000000000000000000000000000000000000000000000000000000000081565b6102fa6107333660046148d2565b612fe7565b610526610746366004614c7d565b613060565b6102fa6107593660046148bb565b61306c565b604080516060810182525f808252602082018190529181018290529061078d8561078786613182565b856131e5565b925050505b9392505050565b6001600160a01b0382165f908152609e602052604081208190816107bc85613182565b815260208082019290925260409081015f2081518083019092525460ff8116151580835261010090910463ffffffff169282019290925291508061080a5750806020015163ffffffff164311155b9150505b92915050565b60405163237dfb4760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906346fbf68e90602401602060405180830381865afa158015610876573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061089a91906152fb565b6108b757604051631d77d47760e21b815260040160405180910390fd5b60665481811681146108dc5760405163c61dca5d60e01b815260040160405180910390fd5b6108e582613351565b5050565b6001600160a01b0381165f908152609d602052604081206060919061090d9061338e565b90505f816001600160401b0381111561092857610928614731565b60405190808252806020026020018201604052801561096c57816020015b604080518082019091525f80825260208201528152602001906001900390816109465790505b5090505f5b828110156109f8576001600160a01b0385165f908152609d602052604090206109d39061099e9083613397565b604080518082019091525f80825260208201525060408051808201909152606082901c815263ffffffff909116602082015290565b8282815181106109e5576109e561531a565b6020908102919091010152600101610971565b509392505050565b60208082015182516001600160a01b03165f90815260989092526040822061080e9163ffffffff908116906133a216565b82610a3b816133b9565b610a585760405163932d94f760e01b815260040160405180910390fd5b5f5b82811015610ccd575f6040518060400160405280876001600160a01b03168152602001868685818110610a8f57610a8f61531a565b9050602002810190610aa1919061532e565b610aaf90602081019061534c565b63ffffffff168152509050610af9816020015163ffffffff1660985f896001600160a01b03166001600160a01b031681526020019081526020015f2061346390919063ffffffff16565b610b1657604051631fb1705560e21b815260040160405180910390fd5b7f31629285ead2335ae0933f86ed2ae63321f7af77b4e6eaabc42c057880977e6c6040518060400160405280886001600160a01b03168152602001836020015163ffffffff16815250604051610b6c9190615365565b60405180910390a15f610b7e82613182565b90505f5b868685818110610b9457610b9461531a565b9050602002810190610ba6919061532e565b610bb4906020810190615373565b9050811015610cc257610c2a878786818110610bd257610bd261531a565b9050602002810190610be4919061532e565b610bf2906020810190615373565b83818110610c0257610c0261531a565b9050602002016020810190610c1791906148d2565b5f8481526099602052604090209061346e565b507f7ab260fe0af193db5f4986770d831bda4ea46099dc817e8b6716dcae8af8e88b83888887818110610c5f57610c5f61531a565b9050602002810190610c71919061532e565b610c7f906020810190615373565b84818110610c8f57610c8f61531a565b9050602002016020810190610ca491906148d2565b604051610cb29291906153b8565b60405180910390a1600101610b82565b505050600101610a5a565b5050505050565b6060610ce284848443613482565b949350505050565b6060610cf885858585613482565b90505f5b8451811015610d8057610d28858281518110610d1a57610d1a61531a565b602002602001015187610799565b610d78575f5b8451811015610d76575f838381518110610d4a57610d4a61531a565b60200260200101518281518110610d6357610d6361531a565b6020908102919091010152600101610d2e565b505b600101610cfc565b50949350505050565b6001600160a01b038082165f908152609760205260408120549091168015610db15780610792565b5090919050565b606654600190600290811603610de15760405163840a48d560e01b815260040160405180910390fd5b82610deb816133b9565b610e085760405163932d94f760e01b815260040160405180910390fd5b5f6040518060400160405280866001600160a01b03168152602001856020016020810190610e36919061534c565b63ffffffff1690529050610e4d6060850185615373565b9050610e5c6040860186615373565b905014610e7c576040516343714afd60e01b815260040160405180910390fd5b60208082015182516001600160a01b03165f90815260989092526040909120610eae9163ffffffff908116906133a216565b610ecb57604051631fb1705560e21b815260040160405180910390fd5b610ee1610edb60208601866148d2565b82610799565b610efe5760405163ebbff49760e01b815260040160405180910390fd5b5f610f0c6040860186615373565b90506001600160401b03811115610f2557610f25614731565b604051908082528060200260200182016040528015610f4e578160200160208202803683370190505b5090505f5b610f606040870187615373565b905081101561155457801580610ff35750610f7e6040870187615373565b610f896001846153f2565b818110610f9857610f9861531a565b9050602002016020810190610fad91906148d2565b6001600160a01b0316610fc36040880188615373565b83818110610fd357610fd361531a565b9050602002016020810190610fe891906148d2565b6001600160a01b0316115b61101057604051639f1c805360e01b815260040160405180910390fd5b61101d6060870187615373565b8281811061102d5761102d61531a565b905060200201355f10801561106d5750670de0b6b3a76400006110536060880188615373565b838181106110635761106361531a565b9050602002013511155b61108a57604051631353603160e01b815260040160405180910390fd5b6110e661109a6040880188615373565b838181106110aa576110aa61531a565b90506020020160208101906110bf91906148d2565b60995f6110cb87613182565b81526020019081526020015f2061376f90919063ffffffff16565b611103576040516331bc342760e11b815260040160405180910390fd5b5f8061115561111560208a018a6148d2565b61111e87613182565b61112b60408c018c615373565b8781811061113b5761113b61531a565b905060200201602081019061115091906148d2565b6131e5565b805191935091506001600160401b03165f0361117257505061154c565b5f6111ad61118360608b018b615373565b868181106111935761119361531a565b85516001600160401b031692602090910201359050613790565b83519091506111c86001600160401b038084169083166137a6565b8686815181106111da576111da61531a565b60200260200101818152505081835f018181516111f79190615405565b6001600160401b0316905250835182908590611214908390615405565b6001600160401b0316905250602084018051839190611234908390615405565b6001600160401b031690525060208301515f600f9190910b121561134c575f61129761126360608d018d615373565b888181106112735761127361531a565b90506020020135856020015161128890615424565b6001600160801b031690613790565b9050806001600160401b0316846020018181516112b49190615448565b600f0b9052507f1487af5418c47ee5ea45ef4a93398668120890774a9e13487e61e9dc3baf76dd6112e860208d018d6148d2565b896112f660408f018f615373565b8a8181106113065761130661531a565b905060200201602081019061131b91906148d2565b61132c885f015189602001516137ba565b8860400151604051611342959493929190615475565b60405180910390a1505b61139e61135c60208c018c6148d2565b61136589613182565b61137260408e018e615373565b898181106113825761138261531a565b905060200201602081019061139791906148d2565b87876137ce565b7f1487af5418c47ee5ea45ef4a93398668120890774a9e13487e61e9dc3baf76dd6113cc60208c018c6148d2565b886113da60408e018e615373565b898181106113ea576113ea61531a565b90506020020160208101906113ff91906148d2565b865160405161141394939291904390615475565b60405180910390a161146461142b60208c018c6148d2565b61143860408d018d615373565b888181106114485761144861531a565b905060200201602081019061145d91906148d2565b8651613a0e565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001663601bb36f6114a060208d018d6148d2565b6114ad60408e018e615373565b898181106114bd576114bd61531a565b90506020020160208101906114d291906148d2565b875160405160e085901b6001600160e01b03191681526001600160a01b0393841660048201529290911660248301526001600160401b0380861660448401521660648201526084015f604051808303815f87803b158015611531575f5ffd5b505af1158015611543573d5f5f3e3d5ffd5b50505050505050505b600101610f53565b507f80969ad29428d6797ee7aad084f9e4a42a82fc506dcd2ca3b6fb431f85ccebe561158360208701876148d2565b836115916040890189615373565b8561159f60808c018c6154c6565b6040516115b29796959493929190615530565b60405180910390a1505050505050565b6001600160a01b0382165f908152609d6020526040812060609182916115e79061338e565b90505f816001600160401b0381111561160257611602614731565b60405190808252806020026020018201604052801561164657816020015b604080518082019091525f80825260208201528152602001906001900390816116205790505b5090505f826001600160401b0381111561166257611662614731565b6040519080825280602002602001820160405280156116ab57816020015b604080516060810182525f80825260208083018290529282015282525f199092019101816116805790505b5090505f5b8381101561172e576001600160a01b0388165f908152609d602052604081206116dd9061099e9084613397565b9050808483815181106116f2576116f261531a565b602002602001018190525061170889828a61075e565b83838151811061171a5761171a61531a565b6020908102919091010152506001016116b0565b509093509150505b9250929050565b60605f61079260995f61174f86613182565b81526020019081526020015f20613a90565b60605f83516001600160401b0381111561177d5761177d614731565b6040519080825280602002602001820160405280156117a6578160200160208202803683370190505b5090505f5b84518110156109f8576117d78582815181106117c9576117c961531a565b6020026020010151856127de565b8282815181106117e9576117e961531a565b6001600160401b03909216602092830291909101909101526001016117ab565b6066545f906001908116036118315760405163840a48d560e01b815260040160405180910390fd5b838214611851576040516343714afd60e01b815260040160405180910390fd5b5f5b848110156118ba576118b2878787848181106118715761187161531a565b905060200201602081019061188691906148d2565b8686858181106118985761189861531a565b90506020020160208101906118ad91906155c6565b613a9c565b600101611853565b50505050505050565b836118cd816133b9565b6118ea5760405163932d94f760e01b815260040160405180910390fd5b604080518082019091526001600160a01b038616815263ffffffff851660208201525f61191682613182565b9050611957826020015163ffffffff1660985f8a6001600160a01b03166001600160a01b031681526020019081526020015f206133a290919063ffffffff16565b61197457604051631fb1705560e21b815260040160405180910390fd5b5f5b84811015611a1757611993868683818110610c0257610c0261531a565b6119b05760405163585cfb2f60e01b815260040160405180910390fd5b7f7ab260fe0af193db5f4986770d831bda4ea46099dc817e8b6716dcae8af8e88b838787848181106119e4576119e461531a565b90506020020160208101906119f991906148d2565b604051611a079291906153b8565b60405180910390a1600101611976565b5050505050505050565b60605f82516001600160401b03811115611a3d57611a3d614731565b604051908082528060200260200182016040528015611a66578160200160208202803683370190505b5090505f5b83518110156109f857611a9785858381518110611a8a57611a8a61531a565b60200260200101516127de565b828281518110611aa957611aa961531a565b6001600160401b0390921660209283029190910190910152600101611a6b565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614611bc457611b02826133b9565b611b1f576040516348f5c3ed60e01b815260040160405180910390fd5b6040516336b87bd760e11b81526001600160a01b0383811660048301527f00000000000000000000000000000000000000000000000000000000000000001690636d70f7ae90602401602060405180830381865afa158015611b83573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611ba791906152fb565b611bc45760405163ccea9e6f60e01b815260040160405180910390fd5b6108e58282613ba0565b60405163237dfb4760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906346fbf68e90602401602060405180830381865afa158015611c30573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611c5491906152fb565b611c7157604051631d77d47760e21b815260040160405180910390fd5b611c7b5f19613351565b565b5f61079283609a5f6110cb86613182565b5f5f611c9a8484613d4c565b95945050505050565b606654600290600490811603611ccc5760405163840a48d560e01b815260040160405180910390fd5b611ce1611cdc60208401846148d2565b6133b9565b80611cfa5750611cfa611cdc60408401602085016148d2565b611d17576040516348f5c3ed60e01b815260040160405180910390fd5b5f5b611d266040840184615373565b9050811015611fe8575f6040518060400160405280856020016020810190611d4e91906148d2565b6001600160a01b03168152602001611d696040870187615373565b85818110611d7957611d7961531a565b9050602002016020810190611d8e919061534c565b63ffffffff168152509050611ddb816020015163ffffffff1660985f876020016020810190611dbd91906148d2565b6001600160a01b0316815260208101919091526040015f20906133a2565b611df857604051631fb1705560e21b815260040160405180910390fd5b609e5f611e0860208701876148d2565b6001600160a01b03166001600160a01b031681526020019081526020015f205f611e3183613182565b815260208101919091526040015f205460ff16611e61576040516325131d4f60e01b815260040160405180910390fd5b611e9b611e6d82613182565b609c5f611e7d60208901896148d2565b6001600160a01b0316815260208101919091526040015f2090613ebb565b50611ed3611eac60208601866148d2565b609a5f611eb885613182565b81526020019081526020015f20613ec690919063ffffffff16565b50611ee160208501856148d2565b6001600160a01b03167fad34c3070be1dffbcaa499d000ba2b8d9848aefcac3059df245dd95c4ece14fe82604051611f199190615365565b60405180910390a2604080518082019091525f815260208101611f5c7f0000000000000000000000000000000000000000000000000000000000000000436155e7565b63ffffffff169052609e5f611f7460208801886148d2565b6001600160a01b03166001600160a01b031681526020019081526020015f205f611f9d84613182565b81526020808201929092526040015f2082518154939092015163ffffffff166101000264ffffffff00199215159290921664ffffffffff199093169290921717905550600101611d19565b50611ffc6103bf60408401602085016148d2565b6001600160a01b0316639d8e0c2361201760208501856148d2565b6120246040860186615373565b6040518463ffffffff1660e01b81526004016120429392919061563c565b5f604051808303815f87803b158015612059575f5ffd5b505af192505050801561206a575060015b156108e5575050565b606061080e609a5f61174f85613182565b61208c613eda565b611c7b5f613f34565b6001600160a01b0381165f908152609c60205260408120606091906120b99061338e565b90505f816001600160401b038111156120d4576120d4614731565b60405190808252806020026020018201604052801561211857816020015b604080518082019091525f80825260208201528152602001906001900390816120f25790505b5090505f5b828110156109f8576001600160a01b0385165f908152609c6020526040902061214a9061099e9083613397565b82828151811061215c5761215c61531a565b602090810291909101015260010161211d565b60605f84516001600160401b0381111561218b5761218b614731565b6040519080825280602002602001820160405280156121d457816020015b604080516060810182525f80825260208083018290529282015282525f199092019101816121a95790505b5090505f5b8551811015610d80576122068682815181106121f7576121f761531a565b6020026020010151868661075e565b8282815181106122185761221861531a565b60209081029190910101526001016121d9565b60605f83516001600160401b0381111561224757612247614731565b604051908082528060200260200182016040528015612270578160200160208202803683370190505b5090505f5b8451811015610d80576001600160a01b0386165f90815260a16020526040812086516122e5928792918990869081106122b0576122b061531a565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020015f20613f8590919063ffffffff16565b8282815181106122f7576122f761531a565b6001600160401b0390921660209283029190910190910152600101612275565b6066545f9060019081160361233f5760405163840a48d560e01b815260040160405180910390fd5b612348836133b9565b612365576040516348f5c3ed60e01b815260040160405180910390fd5b5f5f5f61237186612d48565b91509150816123935760405163fa55fc8160e01b815260040160405180910390fd5b91505f90505b8351811015610ccd578381815181106123b4576123b461531a565b602002602001015160400151518482815181106123d3576123d361531a565b60200260200101516020015151146123fe576040516343714afd60e01b815260040160405180910390fd5b5f8482815181106124115761241161531a565b602090810291909101810151518082015181516001600160a01b03165f908152609890935260409092209092506124519163ffffffff908116906133a216565b61246e57604051631fb1705560e21b815260040160405180910390fd5b5f6124798783610799565b90505f5b86848151811061248f5761248f61531a565b602002602001015160200151518110156127d3575f8785815181106124b6576124b661531a565b60200260200101516020015182815181106124d3576124d361531a565b602002602001015190506124ea898261ffff613a9c565b5f5f6124f98b61078788613182565b915091508060200151600f0b5f1461252457604051630d8fcbe360e41b815260040160405180910390fd5b5f61253187858489613f99565b9050612576825f01518c8a8151811061254c5761254c61531a565b60200260200101516040015187815181106125695761256961531a565b6020026020010151613fcf565b600f0b602083018190525f0361259f57604051634606179360e11b815260040160405180910390fd5b5f8260200151600f0b12156126e3578015612665576126206125c088613182565b6001600160a01b03808f165f90815260a360209081526040808320938a16835292905220908154600160801b90819004600f0b5f818152600180860160205260409091209390935583546001600160801b03908116939091011602179055565b61264a7f0000000000000000000000000000000000000000000000000000000000000000436155e7565b6126559060016155e7565b63ffffffff166040830152612750565b612677836020015183602001516137ba565b6001600160401b031660208401528a518b90899081106126995761269961531a565b60200260200101516040015185815181106126b6576126b661531a565b6020908102919091018101516001600160401b031683525f9083015263ffffffff43166040830152612750565b5f8260200151600f0b131561275057612704836020015183602001516137ba565b6001600160401b03908116602085018190528451909116101561273a57604051636c9be0bf60e01b815260040160405180910390fd5b61274489436155e7565b63ffffffff1660408301525b6127658c61275d89613182565b8686866137ce565b7f1487af5418c47ee5ea45ef4a93398668120890774a9e13487e61e9dc3baf76dd8c61279361099e8a613182565b866127a5865f015187602001516137ba565b86604001516040516127bb959493929190615475565b60405180910390a150506001909201915061247d9050565b505050600101612399565b6001600160a01b038083165f90815260a160209081526040808320938516835292905290812061079290613fe6565b82612817816133b9565b6128345760405163932d94f760e01b815260040160405180910390fd5b836001600160a01b03167fa89c1dc243d8908a96dd84944bcc97d6bc6ac00dd78e20621576be6a3c943713848460405161286f929190615660565b60405180910390a250505050565b6066546002906004908116036128a65760405163840a48d560e01b815260040160405180910390fd5b826128b0816133b9565b6128cd5760405163932d94f760e01b815260040160405180910390fd5b6040516336b87bd760e11b81526001600160a01b0385811660048301527f00000000000000000000000000000000000000000000000000000000000000001690636d70f7ae90602401602060405180830381865afa158015612931573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061295591906152fb565b6129725760405163ccea9e6f60e01b815260040160405180910390fd5b5f5b6129816020850185615373565b9050811015612b4957604080518082019091525f90806129a460208801886148d2565b6001600160a01b031681526020018680602001906129c29190615373565b858181106129d2576129d261531a565b90506020020160208101906129e7919061534c565b63ffffffff90811690915260208083015183516001600160a01b03165f90815260989092526040909120929350612a239291908116906133a216565b612a4057604051631fb1705560e21b815260040160405180910390fd5b612a4a8682610799565b15612a6857604051636c6c6e2760e11b815260040160405180910390fd5b612a91612a7482613182565b6001600160a01b0388165f908152609c6020526040902090613463565b50612abd86609a5f612aa285613182565b81526020019081526020015f2061346e90919063ffffffff16565b50856001600160a01b03167f43232edf9071753d2321e5fa7e018363ee248e5f2142e6c08edd3265bfb4895e82604051612af79190615365565b60405180910390a26001600160a01b0386165f908152609e60205260408120600191612b2284613182565b815260208101919091526040015f20805460ff191691151591909117905550600101612974565b50612b5a6103bf60208501856148d2565b6001600160a01b031663adcf73f785612b766020870187615373565b612b8360408901896154c6565b6040518663ffffffff1660e01b8152600401612ba3959493929190615673565b5f604051808303815f87803b158015612bba575f5ffd5b505af1158015611a17573d5f5f3e3d5ffd5b5f61080e609a5f612bdc85613182565b81526020019081526020015f2061338e565b83612bf8816133b9565b612c155760405163932d94f760e01b815260040160405180910390fd5b6040805180820182526001600160a01b03871680825263ffffffff80881660208085018290525f93845260989052939091209192612c5492916133a216565b612c7157604051631fb1705560e21b815260040160405180910390fd5b5f612c7b82613182565b90505f5b84811015611a1757612cc4868683818110612c9c57612c9c61531a565b9050602002016020810190612cb191906148d2565b5f84815260996020526040902090613ec6565b612ce1576040516331bc342760e11b815260040160405180910390fd5b7f7b4b073d80dcac55a11177d8459ad9f664ceeb91f71f27167bb14f8152a7eeee83878784818110612d1557612d1561531a565b9050602002016020810190612d2a91906148d2565b604051612d389291906153b8565b60405180910390a1600101612c7f565b6001600160a01b0381165f908152609b602090815260408083208151608081018352905463ffffffff80821680845260ff600160201b8404161515958401869052650100000000008304821694840194909452600160481b909104166060820181905284939192919015801590612dc95750826060015163ffffffff164310155b15612dd8575050604081015160015b9590945092505050565b6001600160a01b0381165f90815260986020526040812061080e9061338e565b6001600160a01b0382165f908152609f602052604081206060919061080a908261174f86613182565b5f54610100900460ff1615808015612e4957505f54600160ff909116105b80612e625750303b158015612e6257505f5460ff166001145b612eca5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b5f805460ff191660011790558015612eeb575f805461ff0019166101001790555b612ef482613351565b612efd83613f34565b8015612f43575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498906020015b60405180910390a15b505050565b81612f52816133b9565b612f6f5760405163932d94f760e01b815260040160405180910390fd5b6001600160a01b038381165f90815260976020526040902080546001600160a01b0319169184169190911790557f2ae945c40c44dc0ec263f95609c3fdc6952e0aefa22d6374e44f2c997acedf8583612fc781610d89565b604080516001600160a01b03938416815292909116602083015201612f3a565b612fef613eda565b6001600160a01b0381166130545760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401612ec1565b61305d81613f34565b50565b5f5f610d808484613d4c565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156130c8573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906130ec91906156b6565b6001600160a01b0316336001600160a01b03161461311d5760405163794821ff60e01b815260040160405180910390fd5b606654801982198116146131445760405163c61dca5d60e01b815260040160405180910390fd5b606682905560405182815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c9060200160405180910390a25050565b5f815f0151826020015163ffffffff166040516020016131cd92919060609290921b6bffffffffffffffffffffffff1916825260a01b6001600160a01b031916601482015260200190565b60405160208183030381529060405261080e906156d1565b6040805180820182525f80825260208083018290528351606081018552828152808201839052808501839052845180860186526001600160a01b03898116855260a184528685209088168552909252938220929392819061324590613fe6565b6001600160401b0390811682526001600160a01b038981165f81815260a260209081526040808320948c168084529482528083205486169682019690965291815260a082528481208b8252825284812092815291815290839020835160608101855290549283168152600160401b8304600f0b91810191909152600160c01b90910463ffffffff169181018290529192504310156132e7579092509050613349565b6132f8815f015182602001516137ba565b6001600160401b0316815260208101515f600f9190910b121561333657613327826020015182602001516137ba565b6001600160401b031660208301525b5f60408201819052602082015290925090505b935093915050565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a250565b5f61080e825490565b5f6107928383613ff9565b5f8181526001830160205260408120541515610792565b604051631beb2b9760e31b81526001600160a01b0382811660048301523360248301523060448301525f80356001600160e01b0319166064840152917f00000000000000000000000000000000000000000000000000000000000000009091169063df595cb8906084016020604051808303815f875af115801561343f573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061080e91906152fb565b5f610792838361401f565b5f610792836001600160a01b03841661401f565b606083516001600160401b0381111561349d5761349d614731565b6040519080825280602002602001820160405280156134d057816020015b60608152602001906001900390816134bb5790505b5090505f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f0e0e67686866040518363ffffffff1660e01b81526004016135229291906156f4565b5f60405180830381865afa15801561353c573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526135639190810190615718565b90505f5b8551811015613765575f8682815181106135835761358361531a565b6020026020010151905085516001600160401b038111156135a6576135a6614731565b6040519080825280602002602001820160405280156135cf578160200160208202803683370190505b508483815181106135e2576135e261531a565b60209081029190910101525f5b865181101561375b575f87828151811061360b5761360b61531a565b6020908102919091018101516001600160a01b038086165f90815260a184526040808220928416825291909352822090925061364690613fe6565b9050806001600160401b03165f0361365f575050613753565b5f61366b858d8561075e565b90508863ffffffff16816040015163ffffffff161115801561369357505f8160200151600f0b125b156136b5576136a9815f015182602001516137ba565b6001600160401b031681525b80515f906136d0906001600160401b039081169085166137a6565b9050613717818989815181106136e8576136e861531a565b602002602001015187815181106137015761370161531a565b602002602001015161406b90919063ffffffff16565b8988815181106137295761372961531a565b602002602001015186815181106137425761374261531a565b602002602001018181525050505050505b6001016135ef565b5050600101613567565b5050949350505050565b6001600160a01b0381165f9081526001830160205260408120541515610792565b5f6107928383670de0b6b3a7640000600161407f565b5f61079283670de0b6b3a7640000846140d8565b5f610792826001600160401b038516615448565b6020808301516001600160a01b038088165f90815260a284526040808220928816825291909352909120546001600160401b0390811691161461389457602082810180516001600160a01b038881165f81815260a286526040808220938a1680835293875290819020805467ffffffffffffffff19166001600160401b0395861617905593518451918252948101919091529216908201527facf9095feb3a370c9cf692421c69ef320d4db5c66e6a7d29c7694eb02364fc559060600160405180910390a15b6001600160a01b038086165f90815260a060209081526040808320888452825280832093871683529281529082902083518154928501519385015163ffffffff16600160c01b0263ffffffff60c01b196001600160801b038616600160401b026001600160c01b03199095166001600160401b03909316929092179390931716919091179055600f0b15613976576001600160a01b0385165f908152609f60209081526040808320878452909152902061394e908461346e565b506001600160a01b0385165f908152609d602052604090206139709085613463565b50610ccd565b80516001600160401b03165f03610ccd576001600160a01b0385165f908152609f6020908152604080832087845290915290206139b39084613ec6565b506001600160a01b0385165f908152609f6020908152604080832087845290915290206139df9061338e565b5f03610ccd576001600160a01b0385165f908152609d60205260409020613a069085613ebb565b505050505050565b6001600160a01b038084165f90815260a160209081526040808320938616835292905220613a3d9043836141bd565b604080516001600160a01b038086168252841660208201526001600160401b038316918101919091527f1c6458079a41077d003c11faf9bf097e693bd67979e4e6500bac7b29db779b5c90606001612f3a565b60605f610792836141d1565b6001600160a01b038381165f90815260a360209081526040808320938616835292905290812054600f81810b600160801b909204900b035b5f81118015613ae657508261ffff1682105b15610ccd576001600160a01b038086165f90815260a3602090815260408083209388168352929052908120613b1a9061422a565b90505f5f613b298884896131e5565b91509150806040015163ffffffff16431015613b4757505050610ccd565b613b5488848985856137ce565b6001600160a01b038089165f90815260a360209081526040808320938b16835292905220613b819061427c565b50613b8b85615824565b9450613b968461583c565b9350505050613ad4565b6001600160a01b0382165f908152609b60209081526040918290208251608081018452905463ffffffff808216835260ff600160201b830416151593830193909352650100000000008104831693820193909352600160481b909204166060820181905215801590613c1c5750806060015163ffffffff164310155b15613c3657604081015163ffffffff168152600160208201525b63ffffffff82166040820152613c6c7f0000000000000000000000000000000000000000000000000000000000000000436155e7565b613c779060016155e7565b63ffffffff90811660608381019182526001600160a01b0386165f818152609b602090815260409182902087518154838a0151858b01519851928a1664ffffffffff1990921691909117600160201b91151591909102176cffffffffffffffff0000000000191665010000000000978916979097026cffffffff000000000000000000191696909617600160481b968816968702179055815192835294871694820194909452928301919091527f4e85751d6331506c6c62335f207eb31f12a61e570f34f5c17640308785c6d4db9101612f3a565b6001600160a01b038281165f81815260a2602090815260408083209486168084529482528083205493835260a38252808320948352939052918220546001600160401b039091169190600f81810b600160801b909204900b03815b81811015613e77576001600160a01b038087165f90815260a3602090815260408083209389168352929052908120613ddf90836142f9565b6001600160a01b038881165f90815260a0602090815260408083208584528252808320938b16835292815290829020825160608101845290546001600160401b0381168252600160401b8104600f0b92820192909252600160c01b90910463ffffffff16918101829052919250431015613e5a575050613e77565b613e688682602001516137ba565b95505050806001019050613da7565b506001600160a01b038086165f90815260a1602090815260408083209388168352929052208390613ea790613fe6565b613eb19190615405565b9150509250929050565b5f6107928383614368565b5f610792836001600160a01b038416614368565b6033546001600160a01b03163314611c7b5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401612ec1565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b5f6107928383670de0b6b3a764000061444b565b5f613faa8460995f6110cb89613182565b8015613fb35750815b8015611c9a57505090516001600160401b031615159392505050565b5f6107926001600160401b03808516908416615851565b5f61080e82670de0b6b3a76400006144a0565b5f825f01828154811061400e5761400e61531a565b905f5260205f200154905092915050565b5f81815260018301602052604081205461406457508154600181810184555f84815260208082209093018490558454848252828601909352604090209190915561080e565b505f61080e565b5f6107928383670de0b6b3a76400006140d8565b5f5f61408c8686866140d8565b905060018360028111156140a2576140a261587e565b1480156140be57505f84806140b9576140b9615892565b868809115b15611c9a576140ce6001826158a6565b9695505050505050565b5f80805f19858709858702925082811083820303915050805f0361410f5783828161410557614105615892565b0492505050610792565b8084116141565760405162461bcd60e51b81526020600482015260156024820152744d6174683a206d756c446976206f766572666c6f7760581b6044820152606401612ec1565b5f8486880960026001871981018816978890046003810283188082028403028082028403028082028403028082028403028082028403029081029092039091025f889003889004909101858311909403939093029303949094049190911702949350505050565b612f4383836001600160401b0384166144d7565b6060815f0180548060200260200160405190810160405280929190818152602001828054801561421e57602002820191905f5260205f20905b81548152602001906001019080831161420a575b50505050509050919050565b5f6142448254600f81810b600160801b909204900b131590565b1561426257604051631ed9509560e11b815260040160405180910390fd5b508054600f0b5f9081526001909101602052604090205490565b5f6142968254600f81810b600160801b909204900b131590565b156142b457604051631ed9509560e11b815260040160405180910390fd5b508054600f0b5f818152600180840160205260408220805492905583546fffffffffffffffffffffffffffffffff191692016001600160801b03169190911790915590565b5f5f61431b614307846145da565b85546143169190600f0b6158b9565b614647565b8454909150600160801b9004600f90810b9082900b1261434e57604051632d0483c560e21b815260040160405180910390fd5b600f0b5f9081526001939093016020525050604090205490565b5f8181526001830160205260408120548015614442575f61438a6001836153f2565b85549091505f9061439d906001906153f2565b90508181146143fc575f865f0182815481106143bb576143bb61531a565b905f5260205f200154905080875f0184815481106143db576143db61531a565b5f918252602080832090910192909255918252600188019052604090208390555b855486908061440d5761440d6158e0565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f90556001935050505061080e565b5f91505061080e565b82545f908161445c868683856146b0565b9050801561449657614480866144736001846153f2565b5f91825260209091200190565b54600160201b90046001600160e01b031661078d565b5091949350505050565b81545f9080156144cf576144b9846144736001846153f2565b54600160201b90046001600160e01b031661080a565b509092915050565b8254801561458d575f6144ef856144736001856153f2565b60408051808201909152905463ffffffff808216808452600160201b9092046001600160e01b0316602084015291925090851610156145415760405163151b8e3f60e11b815260040160405180910390fd5b805163ffffffff80861691160361458b5782614562866144736001866153f2565b80546001600160e01b0392909216600160201b0263ffffffff9092169190911790555050505050565b505b506040805180820190915263ffffffff92831681526001600160e01b03918216602080830191825285546001810187555f968752952091519051909216600160201b029190921617910155565b5f6001600160ff1b038211156146435760405162461bcd60e51b815260206004820152602860248201527f53616665436173743a2076616c756520646f65736e27742066697420696e2061604482015267371034b73a191a9b60c11b6064820152608401612ec1565b5090565b80600f81900b81146146ab5760405162461bcd60e51b815260206004820152602760248201527f53616665436173743a2076616c756520646f65736e27742066697420696e20316044820152663238206269747360c81b6064820152608401612ec1565b919050565b5f5b818310156109f8575f6146c58484614703565b5f8781526020902090915063ffffffff86169082015463ffffffff1611156146ef578092506146fd565b6146fa8160016158a6565b93505b506146b2565b5f61471160028484186158f4565b610792908484166158a6565b6001600160a01b038116811461305d575f5ffd5b634e487b7160e01b5f52604160045260245ffd5b604051606081016001600160401b038111828210171561476757614767614731565b60405290565b604051601f8201601f191681016001600160401b038111828210171561479557614795614731565b604052919050565b803563ffffffff811681146146ab575f5ffd5b5f604082840312156147c0575f5ffd5b604080519081016001600160401b03811182821017156147e2576147e2614731565b60405290508082356147f38161471d565b81526148016020840161479d565b60208201525092915050565b5f5f5f6080848603121561481f575f5ffd5b833561482a8161471d565b925061483985602086016147b0565b915060608401356148498161471d565b809150509250925092565b81516001600160401b03168152602080830151600f0b9082015260408083015163ffffffff16908201526060810161080e565b5f5f60608385031215614898575f5ffd5b82356148a38161471d565b91506148b284602085016147b0565b90509250929050565b5f602082840312156148cb575f5ffd5b5035919050565b5f602082840312156148e2575f5ffd5b81356107928161471d565b80516001600160a01b0316825260209081015163ffffffff16910152565b5f8151808452602084019350602083015f5b82811015614946576149308683516148ed565b604095909501946020919091019060010161491d565b5093949350505050565b602081525f610792602083018461490b565b5f60408284031215614972575f5ffd5b61079283836147b0565b5f5f83601f84011261498c575f5ffd5b5081356001600160401b038111156149a2575f5ffd5b6020830191508360208260051b8501011115611736575f5ffd5b5f5f5f604084860312156149ce575f5ffd5b83356149d98161471d565b925060208401356001600160401b038111156149f3575f5ffd5b6149ff8682870161497c565b9497909650939450505050565b5f6001600160401b03821115614a2457614a24614731565b5060051b60200190565b5f82601f830112614a3d575f5ffd5b8135614a50614a4b82614a0c565b61476d565b8082825260208201915060208360051b860101925085831115614a71575f5ffd5b602085015b83811015614a97578035614a898161471d565b835260209283019201614a76565b5095945050505050565b5f5f5f60808486031215614ab3575f5ffd5b614abd85856147b0565b925060408401356001600160401b03811115614ad7575f5ffd5b614ae386828701614a2e565b92505060608401356001600160401b03811115614afe575f5ffd5b614b0a86828701614a2e565b9150509250925092565b5f8151808452602084019350602083015f5b82811015614946578151865260209586019590910190600101614b26565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b82811015614b9b57603f19878603018452614b86858351614b14565b94506020938401939190910190600101614b6a565b50929695505050505050565b5f5f5f5f60a08587031215614bba575f5ffd5b614bc486866147b0565b935060408501356001600160401b03811115614bde575f5ffd5b614bea87828801614a2e565b93505060608501356001600160401b03811115614c05575f5ffd5b614c1187828801614a2e565b925050614c206080860161479d565b905092959194509250565b5f5f60408385031215614c3c575f5ffd5b8235614c478161471d565b915060208301356001600160401b03811115614c61575f5ffd5b830160a08186031215614c72575f5ffd5b809150509250929050565b5f5f60408385031215614c8e575f5ffd5b8235614c998161471d565b91506020830135614c728161471d565b5f8151808452602084019350602083015f5b8281101561494657614cf486835180516001600160401b03168252602080820151600f0b9083015260409081015163ffffffff16910152565b6060959095019460209190910190600101614cbb565b604081525f614d1c604083018561490b565b8281036020840152611c9a8185614ca9565b5f8151808452602084019350602083015f5b828110156149465781516001600160a01b0316865260209586019590910190600101614d40565b602081525f6107926020830184614d2e565b5f5f60408385031215614d8a575f5ffd5b82356001600160401b03811115614d9f575f5ffd5b614dab85828601614a2e565b9250506020830135614c728161471d565b602080825282518282018190525f918401906040840190835b81811015614dfc5783516001600160401b0316835260209384019390920191600101614dd5565b509095945050505050565b5f5f5f5f5f60608688031215614e1b575f5ffd5b8535614e268161471d565b945060208601356001600160401b03811115614e40575f5ffd5b614e4c8882890161497c565b90955093505060408601356001600160401b03811115614e6a575f5ffd5b614e768882890161497c565b969995985093965092949392505050565b5f5f5f5f60608587031215614e9a575f5ffd5b8435614ea58161471d565b9350614eb36020860161479d565b925060408501356001600160401b03811115614ecd575f5ffd5b614ed98782880161497c565b95989497509550505050565b5f5f60408385031215614ef6575f5ffd5b8235614f018161471d565b915060208301356001600160401b03811115614f1b575f5ffd5b613eb185828601614a2e565b5f5f60408385031215614f38575f5ffd5b8235614f438161471d565b91506148b26020840161479d565b5f60208284031215614f61575f5ffd5b813560ff81168114610792575f5ffd5b5f60608284031215614f81575f5ffd5b50919050565b5f60208284031215614f97575f5ffd5b81356001600160401b03811115614fac575f5ffd5b61080a84828501614f71565b5f5f5f60808486031215614fca575f5ffd5b83356001600160401b03811115614fdf575f5ffd5b614feb86828701614a2e565b93505061483985602086016147b0565b602081525f6107926020830184614ca9565b5f5f5f6060848603121561501f575f5ffd5b833561502a8161471d565b925060208401356001600160401b03811115615044575f5ffd5b61505086828701614a2e565b92505061505f6040850161479d565b90509250925092565b5f5f60408385031215615079575f5ffd5b82356150848161471d565b915060208301356001600160401b0381111561509e575f5ffd5b8301601f810185136150ae575f5ffd5b80356150bc614a4b82614a0c565b8082825260208201915060208360051b8501019250878311156150dd575f5ffd5b602084015b838110156152025780356001600160401b038111156150ff575f5ffd5b85016080818b03601f19011215615114575f5ffd5b61511c614745565b6151298b602084016147b0565b815260608201356001600160401b03811115615143575f5ffd5b6151528c602083860101614a2e565b60208301525060808201356001600160401b03811115615170575f5ffd5b6020818401019250508a601f830112615187575f5ffd5b8135615195614a4b82614a0c565b8082825260208201915060208360051b86010192508d8311156151b6575f5ffd5b6020850194505b828510156151ec5784356001600160401b03811681146151db575f5ffd5b8252602094850194909101906151bd565b60408401525050845250602092830192016150e2565b50809450505050509250929050565b5f5f5f60408486031215615223575f5ffd5b833561522e8161471d565b925060208401356001600160401b03811115615248575f5ffd5b8401601f81018613615258575f5ffd5b80356001600160401b0381111561526d575f5ffd5b86602082840101111561527e575f5ffd5b939660209190910195509293505050565b5f5f604083850312156152a0575f5ffd5b82356152ab8161471d565b915060208301356001600160401b038111156152c5575f5ffd5b613eb185828601614f71565b5f5f604083850312156152e2575f5ffd5b82356152ed8161471d565b946020939093013593505050565b5f6020828403121561530b575f5ffd5b81518015158114610792575f5ffd5b634e487b7160e01b5f52603260045260245ffd5b5f8235603e19833603018112615342575f5ffd5b9190910192915050565b5f6020828403121561535c575f5ffd5b6107928261479d565b6040810161080e82846148ed565b5f5f8335601e19843603018112615388575f5ffd5b8301803591506001600160401b038211156153a1575f5ffd5b6020019150600581901b3603821315611736575f5ffd5b606081016153c682856148ed565b6001600160a01b039290921660409190910152919050565b634e487b7160e01b5f52601160045260245ffd5b8181038181111561080e5761080e6153de565b6001600160401b03828116828216039081111561080e5761080e6153de565b5f81600f0b60016001607f1b03198103615440576154406153de565b5f0392915050565b600f81810b9083900b0160016001607f1b03811360016001607f1b03198212171561080e5761080e6153de565b6001600160a01b038616815260c0810161549260208301876148ed565b6001600160a01b039490941660608201526001600160401b0392909216608083015263ffffffff1660a09091015292915050565b5f5f8335601e198436030181126154db575f5ffd5b8301803591506001600160401b038211156154f4575f5ffd5b602001915036819003821315611736575f5ffd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b6001600160a01b03881681525f60c0820161554e602084018a6148ed565b60c060608401528690528660e083015f5b8881101561558f5782356155728161471d565b6001600160a01b031682526020928301929091019060010161555f565b5083810360808501526155a28188614b14565b91505082810360a08401526155b8818587615508565b9a9950505050505050505050565b5f602082840312156155d6575f5ffd5b813561ffff81168114610792575f5ffd5b63ffffffff818116838216019081111561080e5761080e6153de565b8183526020830192505f815f5b848110156149465763ffffffff6156268361479d565b1686526020958601959190910190600101615610565b6001600160a01b03841681526040602082018190525f90611c9a9083018486615603565b602081525f610ce2602083018486615508565b6001600160a01b03861681526060602082018190525f906156979083018688615603565b82810360408401526156aa818587615508565b98975050505050505050565b5f602082840312156156c6575f5ffd5b81516107928161471d565b80516020808301519190811015614f81575f1960209190910360031b1b16919050565b604081525f6157066040830185614d2e565b8281036020840152611c9a8185614d2e565b5f60208284031215615728575f5ffd5b81516001600160401b0381111561573d575f5ffd5b8201601f8101841361574d575f5ffd5b805161575b614a4b82614a0c565b8082825260208201915060208360051b85010192508683111561577c575f5ffd5b602084015b838110156158195780516001600160401b0381111561579e575f5ffd5b8501603f810189136157ae575f5ffd5b60208101516157bf614a4b82614a0c565b808282526020820191506020808460051b8601010192508b8311156157e2575f5ffd5b6040840193505b828410156158045783518252602093840193909101906157e9565b86525050602093840193919091019050615781565b509695505050505050565b5f60018201615835576158356153de565b5060010190565b5f8161584a5761584a6153de565b505f190190565b600f82810b9082900b0360016001607f1b0319811260016001607f1b038213171561080e5761080e6153de565b634e487b7160e01b5f52602160045260245ffd5b634e487b7160e01b5f52601260045260245ffd5b8082018082111561080e5761080e6153de565b8082018281125f8312801582168215821617156158d8576158d86153de565b505092915050565b634e487b7160e01b5f52603160045260245ffd5b5f8261590e57634e487b7160e01b5f52601260045260245ffd5b50049056fea2646970667358221220e8ca7d40cd0cf27e53eb080873f06137efd0ae4dfd7443d78e87c38cd40c571d64736f6c634300081b0033",
>>>>>>> bd2453c0 (feat: add `getAllocatedStake` and update `getMinimumSlashableStake` (#1037))
}

// AllocationManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use AllocationManagerMetaData.ABI instead.
var AllocationManagerABI = AllocationManagerMetaData.ABI

// AllocationManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AllocationManagerMetaData.Bin instead.
var AllocationManagerBin = AllocationManagerMetaData.Bin

// DeployAllocationManager deploys a new Ethereum contract, binding an instance of AllocationManager to it.
func DeployAllocationManager(auth *bind.TransactOpts, backend bind.ContractBackend, _delegation common.Address, _pauserRegistry common.Address, _permissionController common.Address, _DEALLOCATION_DELAY uint32, _ALLOCATION_CONFIGURATION_DELAY uint32, _version string) (common.Address, *types.Transaction, *AllocationManager, error) {
	parsed, err := AllocationManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AllocationManagerBin), backend, _delegation, _pauserRegistry, _permissionController, _DEALLOCATION_DELAY, _ALLOCATION_CONFIGURATION_DELAY, _version)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AllocationManager{AllocationManagerCaller: AllocationManagerCaller{contract: contract}, AllocationManagerTransactor: AllocationManagerTransactor{contract: contract}, AllocationManagerFilterer: AllocationManagerFilterer{contract: contract}}, nil
}

// AllocationManager is an auto generated Go binding around an Ethereum contract.
type AllocationManager struct {
	AllocationManagerCaller     // Read-only binding to the contract
	AllocationManagerTransactor // Write-only binding to the contract
	AllocationManagerFilterer   // Log filterer for contract events
}

// AllocationManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type AllocationManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AllocationManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AllocationManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AllocationManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AllocationManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AllocationManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AllocationManagerSession struct {
	Contract     *AllocationManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AllocationManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AllocationManagerCallerSession struct {
	Contract *AllocationManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// AllocationManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AllocationManagerTransactorSession struct {
	Contract     *AllocationManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// AllocationManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type AllocationManagerRaw struct {
	Contract *AllocationManager // Generic contract binding to access the raw methods on
}

// AllocationManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AllocationManagerCallerRaw struct {
	Contract *AllocationManagerCaller // Generic read-only contract binding to access the raw methods on
}

// AllocationManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AllocationManagerTransactorRaw struct {
	Contract *AllocationManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAllocationManager creates a new instance of AllocationManager, bound to a specific deployed contract.
func NewAllocationManager(address common.Address, backend bind.ContractBackend) (*AllocationManager, error) {
	contract, err := bindAllocationManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AllocationManager{AllocationManagerCaller: AllocationManagerCaller{contract: contract}, AllocationManagerTransactor: AllocationManagerTransactor{contract: contract}, AllocationManagerFilterer: AllocationManagerFilterer{contract: contract}}, nil
}

// NewAllocationManagerCaller creates a new read-only instance of AllocationManager, bound to a specific deployed contract.
func NewAllocationManagerCaller(address common.Address, caller bind.ContractCaller) (*AllocationManagerCaller, error) {
	contract, err := bindAllocationManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AllocationManagerCaller{contract: contract}, nil
}

// NewAllocationManagerTransactor creates a new write-only instance of AllocationManager, bound to a specific deployed contract.
func NewAllocationManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*AllocationManagerTransactor, error) {
	contract, err := bindAllocationManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AllocationManagerTransactor{contract: contract}, nil
}

// NewAllocationManagerFilterer creates a new log filterer instance of AllocationManager, bound to a specific deployed contract.
func NewAllocationManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*AllocationManagerFilterer, error) {
	contract, err := bindAllocationManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AllocationManagerFilterer{contract: contract}, nil
}

// bindAllocationManager binds a generic wrapper to an already deployed contract.
func bindAllocationManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AllocationManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AllocationManager *AllocationManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AllocationManager.Contract.AllocationManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AllocationManager *AllocationManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AllocationManager.Contract.AllocationManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AllocationManager *AllocationManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AllocationManager.Contract.AllocationManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AllocationManager *AllocationManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AllocationManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AllocationManager *AllocationManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AllocationManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AllocationManager *AllocationManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AllocationManager.Contract.contract.Transact(opts, method, params...)
}

// ALLOCATIONCONFIGURATIONDELAY is a free data retrieval call binding the contract method 0x7bc1ef61.
//
// Solidity: function ALLOCATION_CONFIGURATION_DELAY() view returns(uint32)
func (_AllocationManager *AllocationManagerCaller) ALLOCATIONCONFIGURATIONDELAY(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _AllocationManager.contract.Call(opts, &out, "ALLOCATION_CONFIGURATION_DELAY")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ALLOCATIONCONFIGURATIONDELAY is a free data retrieval call binding the contract method 0x7bc1ef61.
//
// Solidity: function ALLOCATION_CONFIGURATION_DELAY() view returns(uint32)
func (_AllocationManager *AllocationManagerSession) ALLOCATIONCONFIGURATIONDELAY() (uint32, error) {
	return _AllocationManager.Contract.ALLOCATIONCONFIGURATIONDELAY(&_AllocationManager.CallOpts)
}

// ALLOCATIONCONFIGURATIONDELAY is a free data retrieval call binding the contract method 0x7bc1ef61.
//
// Solidity: function ALLOCATION_CONFIGURATION_DELAY() view returns(uint32)
func (_AllocationManager *AllocationManagerCallerSession) ALLOCATIONCONFIGURATIONDELAY() (uint32, error) {
	return _AllocationManager.Contract.ALLOCATIONCONFIGURATIONDELAY(&_AllocationManager.CallOpts)
}

// DEALLOCATIONDELAY is a free data retrieval call binding the contract method 0x2981eb77.
//
// Solidity: function DEALLOCATION_DELAY() view returns(uint32)
func (_AllocationManager *AllocationManagerCaller) DEALLOCATIONDELAY(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _AllocationManager.contract.Call(opts, &out, "DEALLOCATION_DELAY")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// DEALLOCATIONDELAY is a free data retrieval call binding the contract method 0x2981eb77.
//
// Solidity: function DEALLOCATION_DELAY() view returns(uint32)
func (_AllocationManager *AllocationManagerSession) DEALLOCATIONDELAY() (uint32, error) {
	return _AllocationManager.Contract.DEALLOCATIONDELAY(&_AllocationManager.CallOpts)
}

// DEALLOCATIONDELAY is a free data retrieval call binding the contract method 0x2981eb77.
//
// Solidity: function DEALLOCATION_DELAY() view returns(uint32)
func (_AllocationManager *AllocationManagerCallerSession) DEALLOCATIONDELAY() (uint32, error) {
	return _AllocationManager.Contract.DEALLOCATIONDELAY(&_AllocationManager.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_AllocationManager *AllocationManagerCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AllocationManager.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_AllocationManager *AllocationManagerSession) Delegation() (common.Address, error) {
	return _AllocationManager.Contract.Delegation(&_AllocationManager.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_AllocationManager *AllocationManagerCallerSession) Delegation() (common.Address, error) {
	return _AllocationManager.Contract.Delegation(&_AllocationManager.CallOpts)
}

// GetAVSRegistrar is a free data retrieval call binding the contract method 0x304c10cd.
//
// Solidity: function getAVSRegistrar(address avs) view returns(address)
func (_AllocationManager *AllocationManagerCaller) GetAVSRegistrar(opts *bind.CallOpts, avs common.Address) (common.Address, error) {
	var out []interface{}
	err := _AllocationManager.contract.Call(opts, &out, "getAVSRegistrar", avs)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAVSRegistrar is a free data retrieval call binding the contract method 0x304c10cd.
//
// Solidity: function getAVSRegistrar(address avs) view returns(address)
func (_AllocationManager *AllocationManagerSession) GetAVSRegistrar(avs common.Address) (common.Address, error) {
	return _AllocationManager.Contract.GetAVSRegistrar(&_AllocationManager.CallOpts, avs)
}

// GetAVSRegistrar is a free data retrieval call binding the contract method 0x304c10cd.
//
// Solidity: function getAVSRegistrar(address avs) view returns(address)
func (_AllocationManager *AllocationManagerCallerSession) GetAVSRegistrar(avs common.Address) (common.Address, error) {
	return _AllocationManager.Contract.GetAVSRegistrar(&_AllocationManager.CallOpts, avs)
}

// GetAllocatableMagnitude is a free data retrieval call binding the contract method 0x6cfb4481.
//
// Solidity: function getAllocatableMagnitude(address operator, address strategy) view returns(uint64)
func (_AllocationManager *AllocationManagerCaller) GetAllocatableMagnitude(opts *bind.CallOpts, operator common.Address, strategy common.Address) (uint64, error) {
	var out []interface{}
	err := _AllocationManager.contract.Call(opts, &out, "getAllocatableMagnitude", operator, strategy)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetAllocatableMagnitude is a free data retrieval call binding the contract method 0x6cfb4481.
//
// Solidity: function getAllocatableMagnitude(address operator, address strategy) view returns(uint64)
func (_AllocationManager *AllocationManagerSession) GetAllocatableMagnitude(operator common.Address, strategy common.Address) (uint64, error) {
	return _AllocationManager.Contract.GetAllocatableMagnitude(&_AllocationManager.CallOpts, operator, strategy)
}

// GetAllocatableMagnitude is a free data retrieval call binding the contract method 0x6cfb4481.
//
// Solidity: function getAllocatableMagnitude(address operator, address strategy) view returns(uint64)
func (_AllocationManager *AllocationManagerCallerSession) GetAllocatableMagnitude(operator common.Address, strategy common.Address) (uint64, error) {
	return _AllocationManager.Contract.GetAllocatableMagnitude(&_AllocationManager.CallOpts, operator, strategy)
}

// GetAllocatedSets is a free data retrieval call binding the contract method 0x15fe5028.
//
// Solidity: function getAllocatedSets(address operator) view returns((address,uint32)[])
func (_AllocationManager *AllocationManagerCaller) GetAllocatedSets(opts *bind.CallOpts, operator common.Address) ([]OperatorSet, error) {
	var out []interface{}
	err := _AllocationManager.contract.Call(opts, &out, "getAllocatedSets", operator)

	if err != nil {
		return *new([]OperatorSet), err
	}

	out0 := *abi.ConvertType(out[0], new([]OperatorSet)).(*[]OperatorSet)

	return out0, err

}

// GetAllocatedSets is a free data retrieval call binding the contract method 0x15fe5028.
//
// Solidity: function getAllocatedSets(address operator) view returns((address,uint32)[])
func (_AllocationManager *AllocationManagerSession) GetAllocatedSets(operator common.Address) ([]OperatorSet, error) {
	return _AllocationManager.Contract.GetAllocatedSets(&_AllocationManager.CallOpts, operator)
}

// GetAllocatedSets is a free data retrieval call binding the contract method 0x15fe5028.
//
// Solidity: function getAllocatedSets(address operator) view returns((address,uint32)[])
func (_AllocationManager *AllocationManagerCallerSession) GetAllocatedSets(operator common.Address) ([]OperatorSet, error) {
	return _AllocationManager.Contract.GetAllocatedSets(&_AllocationManager.CallOpts, operator)
}

// GetAllocatedStake is a free data retrieval call binding the contract method 0x2b453a9a.
//
// Solidity: function getAllocatedStake((address,uint32) operatorSet, address[] operators, address[] strategies) view returns(uint256[][])
func (_AllocationManager *AllocationManagerCaller) GetAllocatedStake(opts *bind.CallOpts, operatorSet OperatorSet, operators []common.Address, strategies []common.Address) ([][]*big.Int, error) {
	var out []interface{}
	err := _AllocationManager.contract.Call(opts, &out, "getAllocatedStake", operatorSet, operators, strategies)

	if err != nil {
		return *new([][]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([][]*big.Int)).(*[][]*big.Int)

	return out0, err

}

// GetAllocatedStake is a free data retrieval call binding the contract method 0x2b453a9a.
//
// Solidity: function getAllocatedStake((address,uint32) operatorSet, address[] operators, address[] strategies) view returns(uint256[][])
func (_AllocationManager *AllocationManagerSession) GetAllocatedStake(operatorSet OperatorSet, operators []common.Address, strategies []common.Address) ([][]*big.Int, error) {
	return _AllocationManager.Contract.GetAllocatedStake(&_AllocationManager.CallOpts, operatorSet, operators, strategies)
}

// GetAllocatedStake is a free data retrieval call binding the contract method 0x2b453a9a.
//
// Solidity: function getAllocatedStake((address,uint32) operatorSet, address[] operators, address[] strategies) view returns(uint256[][])
func (_AllocationManager *AllocationManagerCallerSession) GetAllocatedStake(operatorSet OperatorSet, operators []common.Address, strategies []common.Address) ([][]*big.Int, error) {
	return _AllocationManager.Contract.GetAllocatedStake(&_AllocationManager.CallOpts, operatorSet, operators, strategies)
}

// GetAllocatedStrategies is a free data retrieval call binding the contract method 0xc221d8ae.
//
// Solidity: function getAllocatedStrategies(address operator, (address,uint32) operatorSet) view returns(address[])
func (_AllocationManager *AllocationManagerCaller) GetAllocatedStrategies(opts *bind.CallOpts, operator common.Address, operatorSet OperatorSet) ([]common.Address, error) {
	var out []interface{}
	err := _AllocationManager.contract.Call(opts, &out, "getAllocatedStrategies", operator, operatorSet)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllocatedStrategies is a free data retrieval call binding the contract method 0xc221d8ae.
//
// Solidity: function getAllocatedStrategies(address operator, (address,uint32) operatorSet) view returns(address[])
func (_AllocationManager *AllocationManagerSession) GetAllocatedStrategies(operator common.Address, operatorSet OperatorSet) ([]common.Address, error) {
	return _AllocationManager.Contract.GetAllocatedStrategies(&_AllocationManager.CallOpts, operator, operatorSet)
}

// GetAllocatedStrategies is a free data retrieval call binding the contract method 0xc221d8ae.
//
// Solidity: function getAllocatedStrategies(address operator, (address,uint32) operatorSet) view returns(address[])
func (_AllocationManager *AllocationManagerCallerSession) GetAllocatedStrategies(operator common.Address, operatorSet OperatorSet) ([]common.Address, error) {
	return _AllocationManager.Contract.GetAllocatedStrategies(&_AllocationManager.CallOpts, operator, operatorSet)
}

// GetAllocation is a free data retrieval call binding the contract method 0x10e1b9b8.
//
// Solidity: function getAllocation(address operator, (address,uint32) operatorSet, address strategy) view returns((uint64,int128,uint32))
func (_AllocationManager *AllocationManagerCaller) GetAllocation(opts *bind.CallOpts, operator common.Address, operatorSet OperatorSet, strategy common.Address) (IAllocationManagerTypesAllocation, error) {
	var out []interface{}
	err := _AllocationManager.contract.Call(opts, &out, "getAllocation", operator, operatorSet, strategy)

	if err != nil {
		return *new(IAllocationManagerTypesAllocation), err
	}

	out0 := *abi.ConvertType(out[0], new(IAllocationManagerTypesAllocation)).(*IAllocationManagerTypesAllocation)

	return out0, err

}

// GetAllocation is a free data retrieval call binding the contract method 0x10e1b9b8.
//
// Solidity: function getAllocation(address operator, (address,uint32) operatorSet, address strategy) view returns((uint64,int128,uint32))
func (_AllocationManager *AllocationManagerSession) GetAllocation(operator common.Address, operatorSet OperatorSet, strategy common.Address) (IAllocationManagerTypesAllocation, error) {
	return _AllocationManager.Contract.GetAllocation(&_AllocationManager.CallOpts, operator, operatorSet, strategy)
}

// GetAllocation is a free data retrieval call binding the contract method 0x10e1b9b8.
//
// Solidity: function getAllocation(address operator, (address,uint32) operatorSet, address strategy) view returns((uint64,int128,uint32))
func (_AllocationManager *AllocationManagerCallerSession) GetAllocation(operator common.Address, operatorSet OperatorSet, strategy common.Address) (IAllocationManagerTypesAllocation, error) {
	return _AllocationManager.Contract.GetAllocation(&_AllocationManager.CallOpts, operator, operatorSet, strategy)
}

// GetAllocationDelay is a free data retrieval call binding the contract method 0xb9fbaed1.
//
// Solidity: function getAllocationDelay(address operator) view returns(bool, uint32)
func (_AllocationManager *AllocationManagerCaller) GetAllocationDelay(opts *bind.CallOpts, operator common.Address) (bool, uint32, error) {
	var out []interface{}
	err := _AllocationManager.contract.Call(opts, &out, "getAllocationDelay", operator)

	if err != nil {
		return *new(bool), *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return out0, out1, err

}

// GetAllocationDelay is a free data retrieval call binding the contract method 0xb9fbaed1.
//
// Solidity: function getAllocationDelay(address operator) view returns(bool, uint32)
func (_AllocationManager *AllocationManagerSession) GetAllocationDelay(operator common.Address) (bool, uint32, error) {
	return _AllocationManager.Contract.GetAllocationDelay(&_AllocationManager.CallOpts, operator)
}

// GetAllocationDelay is a free data retrieval call binding the contract method 0xb9fbaed1.
//
// Solidity: function getAllocationDelay(address operator) view returns(bool, uint32)
func (_AllocationManager *AllocationManagerCallerSession) GetAllocationDelay(operator common.Address) (bool, uint32, error) {
	return _AllocationManager.Contract.GetAllocationDelay(&_AllocationManager.CallOpts, operator)
}

// GetAllocations is a free data retrieval call binding the contract method 0x8ce64854.
//
// Solidity: function getAllocations(address[] operators, (address,uint32) operatorSet, address strategy) view returns((uint64,int128,uint32)[])
func (_AllocationManager *AllocationManagerCaller) GetAllocations(opts *bind.CallOpts, operators []common.Address, operatorSet OperatorSet, strategy common.Address) ([]IAllocationManagerTypesAllocation, error) {
	var out []interface{}
	err := _AllocationManager.contract.Call(opts, &out, "getAllocations", operators, operatorSet, strategy)

	if err != nil {
		return *new([]IAllocationManagerTypesAllocation), err
	}

	out0 := *abi.ConvertType(out[0], new([]IAllocationManagerTypesAllocation)).(*[]IAllocationManagerTypesAllocation)

	return out0, err

}

// GetAllocations is a free data retrieval call binding the contract method 0x8ce64854.
//
// Solidity: function getAllocations(address[] operators, (address,uint32) operatorSet, address strategy) view returns((uint64,int128,uint32)[])
func (_AllocationManager *AllocationManagerSession) GetAllocations(operators []common.Address, operatorSet OperatorSet, strategy common.Address) ([]IAllocationManagerTypesAllocation, error) {
	return _AllocationManager.Contract.GetAllocations(&_AllocationManager.CallOpts, operators, operatorSet, strategy)
}

// GetAllocations is a free data retrieval call binding the contract method 0x8ce64854.
//
// Solidity: function getAllocations(address[] operators, (address,uint32) operatorSet, address strategy) view returns((uint64,int128,uint32)[])
func (_AllocationManager *AllocationManagerCallerSession) GetAllocations(operators []common.Address, operatorSet OperatorSet, strategy common.Address) ([]IAllocationManagerTypesAllocation, error) {
	return _AllocationManager.Contract.GetAllocations(&_AllocationManager.CallOpts, operators, operatorSet, strategy)
}

// GetEncumberedMagnitude is a free data retrieval call binding the contract method 0xf605ce08.
//
// Solidity: function getEncumberedMagnitude(address operator, address strategy) view returns(uint64)
func (_AllocationManager *AllocationManagerCaller) GetEncumberedMagnitude(opts *bind.CallOpts, operator common.Address, strategy common.Address) (uint64, error) {
	var out []interface{}
	err := _AllocationManager.contract.Call(opts, &out, "getEncumberedMagnitude", operator, strategy)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetEncumberedMagnitude is a free data retrieval call binding the contract method 0xf605ce08.
//
// Solidity: function getEncumberedMagnitude(address operator, address strategy) view returns(uint64)
func (_AllocationManager *AllocationManagerSession) GetEncumberedMagnitude(operator common.Address, strategy common.Address) (uint64, error) {
	return _AllocationManager.Contract.GetEncumberedMagnitude(&_AllocationManager.CallOpts, operator, strategy)
}

// GetEncumberedMagnitude is a free data retrieval call binding the contract method 0xf605ce08.
//
// Solidity: function getEncumberedMagnitude(address operator, address strategy) view returns(uint64)
func (_AllocationManager *AllocationManagerCallerSession) GetEncumberedMagnitude(operator common.Address, strategy common.Address) (uint64, error) {
	return _AllocationManager.Contract.GetEncumberedMagnitude(&_AllocationManager.CallOpts, operator, strategy)
}

// GetMaxMagnitude is a free data retrieval call binding the contract method 0xa9333ec8.
//
// Solidity: function getMaxMagnitude(address operator, address strategy) view returns(uint64)
func (_AllocationManager *AllocationManagerCaller) GetMaxMagnitude(opts *bind.CallOpts, operator common.Address, strategy common.Address) (uint64, error) {
	var out []interface{}
	err := _AllocationManager.contract.Call(opts, &out, "getMaxMagnitude", operator, strategy)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetMaxMagnitude is a free data retrieval call binding the contract method 0xa9333ec8.
//
// Solidity: function getMaxMagnitude(address operator, address strategy) view returns(uint64)
func (_AllocationManager *AllocationManagerSession) GetMaxMagnitude(operator common.Address, strategy common.Address) (uint64, error) {
	return _AllocationManager.Contract.GetMaxMagnitude(&_AllocationManager.CallOpts, operator, strategy)
}

// GetMaxMagnitude is a free data retrieval call binding the contract method 0xa9333ec8.
//
// Solidity: function getMaxMagnitude(address operator, address strategy) view returns(uint64)
func (_AllocationManager *AllocationManagerCallerSession) GetMaxMagnitude(operator common.Address, strategy common.Address) (uint64, error) {
	return _AllocationManager.Contract.GetMaxMagnitude(&_AllocationManager.CallOpts, operator, strategy)
}

// GetMaxMagnitudes is a free data retrieval call binding the contract method 0x4a10ffe5.
//
// Solidity: function getMaxMagnitudes(address[] operators, address strategy) view returns(uint64[])
func (_AllocationManager *AllocationManagerCaller) GetMaxMagnitudes(opts *bind.CallOpts, operators []common.Address, strategy common.Address) ([]uint64, error) {
	var out []interface{}
	err := _AllocationManager.contract.Call(opts, &out, "getMaxMagnitudes", operators, strategy)

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

// GetMaxMagnitudes is a free data retrieval call binding the contract method 0x4a10ffe5.
//
// Solidity: function getMaxMagnitudes(address[] operators, address strategy) view returns(uint64[])
func (_AllocationManager *AllocationManagerSession) GetMaxMagnitudes(operators []common.Address, strategy common.Address) ([]uint64, error) {
	return _AllocationManager.Contract.GetMaxMagnitudes(&_AllocationManager.CallOpts, operators, strategy)
}

// GetMaxMagnitudes is a free data retrieval call binding the contract method 0x4a10ffe5.
//
// Solidity: function getMaxMagnitudes(address[] operators, address strategy) view returns(uint64[])
func (_AllocationManager *AllocationManagerCallerSession) GetMaxMagnitudes(operators []common.Address, strategy common.Address) ([]uint64, error) {
	return _AllocationManager.Contract.GetMaxMagnitudes(&_AllocationManager.CallOpts, operators, strategy)
}

// GetMaxMagnitudes0 is a free data retrieval call binding the contract method 0x547afb87.
//
// Solidity: function getMaxMagnitudes(address operator, address[] strategies) view returns(uint64[])
func (_AllocationManager *AllocationManagerCaller) GetMaxMagnitudes0(opts *bind.CallOpts, operator common.Address, strategies []common.Address) ([]uint64, error) {
	var out []interface{}
	err := _AllocationManager.contract.Call(opts, &out, "getMaxMagnitudes0", operator, strategies)

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

// GetMaxMagnitudes0 is a free data retrieval call binding the contract method 0x547afb87.
//
// Solidity: function getMaxMagnitudes(address operator, address[] strategies) view returns(uint64[])
func (_AllocationManager *AllocationManagerSession) GetMaxMagnitudes0(operator common.Address, strategies []common.Address) ([]uint64, error) {
	return _AllocationManager.Contract.GetMaxMagnitudes0(&_AllocationManager.CallOpts, operator, strategies)
}

// GetMaxMagnitudes0 is a free data retrieval call binding the contract method 0x547afb87.
//
// Solidity: function getMaxMagnitudes(address operator, address[] strategies) view returns(uint64[])
func (_AllocationManager *AllocationManagerCallerSession) GetMaxMagnitudes0(operator common.Address, strategies []common.Address) ([]uint64, error) {
	return _AllocationManager.Contract.GetMaxMagnitudes0(&_AllocationManager.CallOpts, operator, strategies)
}

// GetMaxMagnitudesAtBlock is a free data retrieval call binding the contract method 0x94d7d00c.
//
// Solidity: function getMaxMagnitudesAtBlock(address operator, address[] strategies, uint32 blockNumber) view returns(uint64[])
func (_AllocationManager *AllocationManagerCaller) GetMaxMagnitudesAtBlock(opts *bind.CallOpts, operator common.Address, strategies []common.Address, blockNumber uint32) ([]uint64, error) {
	var out []interface{}
	err := _AllocationManager.contract.Call(opts, &out, "getMaxMagnitudesAtBlock", operator, strategies, blockNumber)

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

// GetMaxMagnitudesAtBlock is a free data retrieval call binding the contract method 0x94d7d00c.
//
// Solidity: function getMaxMagnitudesAtBlock(address operator, address[] strategies, uint32 blockNumber) view returns(uint64[])
func (_AllocationManager *AllocationManagerSession) GetMaxMagnitudesAtBlock(operator common.Address, strategies []common.Address, blockNumber uint32) ([]uint64, error) {
	return _AllocationManager.Contract.GetMaxMagnitudesAtBlock(&_AllocationManager.CallOpts, operator, strategies, blockNumber)
}

// GetMaxMagnitudesAtBlock is a free data retrieval call binding the contract method 0x94d7d00c.
//
// Solidity: function getMaxMagnitudesAtBlock(address operator, address[] strategies, uint32 blockNumber) view returns(uint64[])
func (_AllocationManager *AllocationManagerCallerSession) GetMaxMagnitudesAtBlock(operator common.Address, strategies []common.Address, blockNumber uint32) ([]uint64, error) {
	return _AllocationManager.Contract.GetMaxMagnitudesAtBlock(&_AllocationManager.CallOpts, operator, strategies, blockNumber)
}

// GetMemberCount is a free data retrieval call binding the contract method 0xb2447af7.
//
// Solidity: function getMemberCount((address,uint32) operatorSet) view returns(uint256)
func (_AllocationManager *AllocationManagerCaller) GetMemberCount(opts *bind.CallOpts, operatorSet OperatorSet) (*big.Int, error) {
	var out []interface{}
	err := _AllocationManager.contract.Call(opts, &out, "getMemberCount", operatorSet)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMemberCount is a free data retrieval call binding the contract method 0xb2447af7.
//
// Solidity: function getMemberCount((address,uint32) operatorSet) view returns(uint256)
func (_AllocationManager *AllocationManagerSession) GetMemberCount(operatorSet OperatorSet) (*big.Int, error) {
	return _AllocationManager.Contract.GetMemberCount(&_AllocationManager.CallOpts, operatorSet)
}

// GetMemberCount is a free data retrieval call binding the contract method 0xb2447af7.
//
// Solidity: function getMemberCount((address,uint32) operatorSet) view returns(uint256)
func (_AllocationManager *AllocationManagerCallerSession) GetMemberCount(operatorSet OperatorSet) (*big.Int, error) {
	return _AllocationManager.Contract.GetMemberCount(&_AllocationManager.CallOpts, operatorSet)
}

// GetMembers is a free data retrieval call binding the contract method 0x6e875dba.
//
// Solidity: function getMembers((address,uint32) operatorSet) view returns(address[])
func (_AllocationManager *AllocationManagerCaller) GetMembers(opts *bind.CallOpts, operatorSet OperatorSet) ([]common.Address, error) {
	var out []interface{}
	err := _AllocationManager.contract.Call(opts, &out, "getMembers", operatorSet)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetMembers is a free data retrieval call binding the contract method 0x6e875dba.
//
// Solidity: function getMembers((address,uint32) operatorSet) view returns(address[])
func (_AllocationManager *AllocationManagerSession) GetMembers(operatorSet OperatorSet) ([]common.Address, error) {
	return _AllocationManager.Contract.GetMembers(&_AllocationManager.CallOpts, operatorSet)
}

// GetMembers is a free data retrieval call binding the contract method 0x6e875dba.
//
// Solidity: function getMembers((address,uint32) operatorSet) view returns(address[])
func (_AllocationManager *AllocationManagerCallerSession) GetMembers(operatorSet OperatorSet) ([]common.Address, error) {
	return _AllocationManager.Contract.GetMembers(&_AllocationManager.CallOpts, operatorSet)
}

// GetMinimumSlashableStake is a free data retrieval call binding the contract method 0x2bab2c4a.
//
// Solidity: function getMinimumSlashableStake((address,uint32) operatorSet, address[] operators, address[] strategies, uint32 futureBlock) view returns(uint256[][] slashableStake)
func (_AllocationManager *AllocationManagerCaller) GetMinimumSlashableStake(opts *bind.CallOpts, operatorSet OperatorSet, operators []common.Address, strategies []common.Address, futureBlock uint32) ([][]*big.Int, error) {
	var out []interface{}
	err := _AllocationManager.contract.Call(opts, &out, "getMinimumSlashableStake", operatorSet, operators, strategies, futureBlock)

	if err != nil {
		return *new([][]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([][]*big.Int)).(*[][]*big.Int)

	return out0, err

}

// GetMinimumSlashableStake is a free data retrieval call binding the contract method 0x2bab2c4a.
//
// Solidity: function getMinimumSlashableStake((address,uint32) operatorSet, address[] operators, address[] strategies, uint32 futureBlock) view returns(uint256[][] slashableStake)
func (_AllocationManager *AllocationManagerSession) GetMinimumSlashableStake(operatorSet OperatorSet, operators []common.Address, strategies []common.Address, futureBlock uint32) ([][]*big.Int, error) {
	return _AllocationManager.Contract.GetMinimumSlashableStake(&_AllocationManager.CallOpts, operatorSet, operators, strategies, futureBlock)
}

// GetMinimumSlashableStake is a free data retrieval call binding the contract method 0x2bab2c4a.
//
// Solidity: function getMinimumSlashableStake((address,uint32) operatorSet, address[] operators, address[] strategies, uint32 futureBlock) view returns(uint256[][] slashableStake)
func (_AllocationManager *AllocationManagerCallerSession) GetMinimumSlashableStake(operatorSet OperatorSet, operators []common.Address, strategies []common.Address, futureBlock uint32) ([][]*big.Int, error) {
	return _AllocationManager.Contract.GetMinimumSlashableStake(&_AllocationManager.CallOpts, operatorSet, operators, strategies, futureBlock)
}

// GetOperatorSetCount is a free data retrieval call binding the contract method 0xba1a84e5.
//
// Solidity: function getOperatorSetCount(address avs) view returns(uint256)
func (_AllocationManager *AllocationManagerCaller) GetOperatorSetCount(opts *bind.CallOpts, avs common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AllocationManager.contract.Call(opts, &out, "getOperatorSetCount", avs)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOperatorSetCount is a free data retrieval call binding the contract method 0xba1a84e5.
//
// Solidity: function getOperatorSetCount(address avs) view returns(uint256)
func (_AllocationManager *AllocationManagerSession) GetOperatorSetCount(avs common.Address) (*big.Int, error) {
	return _AllocationManager.Contract.GetOperatorSetCount(&_AllocationManager.CallOpts, avs)
}

// GetOperatorSetCount is a free data retrieval call binding the contract method 0xba1a84e5.
//
// Solidity: function getOperatorSetCount(address avs) view returns(uint256)
func (_AllocationManager *AllocationManagerCallerSession) GetOperatorSetCount(avs common.Address) (*big.Int, error) {
	return _AllocationManager.Contract.GetOperatorSetCount(&_AllocationManager.CallOpts, avs)
}

// GetRegisteredSets is a free data retrieval call binding the contract method 0x79ae50cd.
//
// Solidity: function getRegisteredSets(address operator) view returns((address,uint32)[])
func (_AllocationManager *AllocationManagerCaller) GetRegisteredSets(opts *bind.CallOpts, operator common.Address) ([]OperatorSet, error) {
	var out []interface{}
	err := _AllocationManager.contract.Call(opts, &out, "getRegisteredSets", operator)

	if err != nil {
		return *new([]OperatorSet), err
	}

	out0 := *abi.ConvertType(out[0], new([]OperatorSet)).(*[]OperatorSet)

	return out0, err

}

// GetRegisteredSets is a free data retrieval call binding the contract method 0x79ae50cd.
//
// Solidity: function getRegisteredSets(address operator) view returns((address,uint32)[])
func (_AllocationManager *AllocationManagerSession) GetRegisteredSets(operator common.Address) ([]OperatorSet, error) {
	return _AllocationManager.Contract.GetRegisteredSets(&_AllocationManager.CallOpts, operator)
}

// GetRegisteredSets is a free data retrieval call binding the contract method 0x79ae50cd.
//
// Solidity: function getRegisteredSets(address operator) view returns((address,uint32)[])
func (_AllocationManager *AllocationManagerCallerSession) GetRegisteredSets(operator common.Address) ([]OperatorSet, error) {
	return _AllocationManager.Contract.GetRegisteredSets(&_AllocationManager.CallOpts, operator)
}

// GetStrategiesInOperatorSet is a free data retrieval call binding the contract method 0x4177a87c.
//
// Solidity: function getStrategiesInOperatorSet((address,uint32) operatorSet) view returns(address[])
func (_AllocationManager *AllocationManagerCaller) GetStrategiesInOperatorSet(opts *bind.CallOpts, operatorSet OperatorSet) ([]common.Address, error) {
	var out []interface{}
	err := _AllocationManager.contract.Call(opts, &out, "getStrategiesInOperatorSet", operatorSet)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetStrategiesInOperatorSet is a free data retrieval call binding the contract method 0x4177a87c.
//
// Solidity: function getStrategiesInOperatorSet((address,uint32) operatorSet) view returns(address[])
func (_AllocationManager *AllocationManagerSession) GetStrategiesInOperatorSet(operatorSet OperatorSet) ([]common.Address, error) {
	return _AllocationManager.Contract.GetStrategiesInOperatorSet(&_AllocationManager.CallOpts, operatorSet)
}

// GetStrategiesInOperatorSet is a free data retrieval call binding the contract method 0x4177a87c.
//
// Solidity: function getStrategiesInOperatorSet((address,uint32) operatorSet) view returns(address[])
func (_AllocationManager *AllocationManagerCallerSession) GetStrategiesInOperatorSet(operatorSet OperatorSet) ([]common.Address, error) {
	return _AllocationManager.Contract.GetStrategiesInOperatorSet(&_AllocationManager.CallOpts, operatorSet)
}

// GetStrategyAllocations is a free data retrieval call binding the contract method 0x40120dab.
//
// Solidity: function getStrategyAllocations(address operator, address strategy) view returns((address,uint32)[], (uint64,int128,uint32)[])
func (_AllocationManager *AllocationManagerCaller) GetStrategyAllocations(opts *bind.CallOpts, operator common.Address, strategy common.Address) ([]OperatorSet, []IAllocationManagerTypesAllocation, error) {
	var out []interface{}
	err := _AllocationManager.contract.Call(opts, &out, "getStrategyAllocations", operator, strategy)

	if err != nil {
		return *new([]OperatorSet), *new([]IAllocationManagerTypesAllocation), err
	}

	out0 := *abi.ConvertType(out[0], new([]OperatorSet)).(*[]OperatorSet)
	out1 := *abi.ConvertType(out[1], new([]IAllocationManagerTypesAllocation)).(*[]IAllocationManagerTypesAllocation)

	return out0, out1, err

}

// GetStrategyAllocations is a free data retrieval call binding the contract method 0x40120dab.
//
// Solidity: function getStrategyAllocations(address operator, address strategy) view returns((address,uint32)[], (uint64,int128,uint32)[])
func (_AllocationManager *AllocationManagerSession) GetStrategyAllocations(operator common.Address, strategy common.Address) ([]OperatorSet, []IAllocationManagerTypesAllocation, error) {
	return _AllocationManager.Contract.GetStrategyAllocations(&_AllocationManager.CallOpts, operator, strategy)
}

// GetStrategyAllocations is a free data retrieval call binding the contract method 0x40120dab.
//
// Solidity: function getStrategyAllocations(address operator, address strategy) view returns((address,uint32)[], (uint64,int128,uint32)[])
func (_AllocationManager *AllocationManagerCallerSession) GetStrategyAllocations(operator common.Address, strategy common.Address) ([]OperatorSet, []IAllocationManagerTypesAllocation, error) {
	return _AllocationManager.Contract.GetStrategyAllocations(&_AllocationManager.CallOpts, operator, strategy)
}

// IsMemberOfOperatorSet is a free data retrieval call binding the contract method 0x670d3ba2.
//
// Solidity: function isMemberOfOperatorSet(address operator, (address,uint32) operatorSet) view returns(bool)
func (_AllocationManager *AllocationManagerCaller) IsMemberOfOperatorSet(opts *bind.CallOpts, operator common.Address, operatorSet OperatorSet) (bool, error) {
	var out []interface{}
	err := _AllocationManager.contract.Call(opts, &out, "isMemberOfOperatorSet", operator, operatorSet)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMemberOfOperatorSet is a free data retrieval call binding the contract method 0x670d3ba2.
//
// Solidity: function isMemberOfOperatorSet(address operator, (address,uint32) operatorSet) view returns(bool)
func (_AllocationManager *AllocationManagerSession) IsMemberOfOperatorSet(operator common.Address, operatorSet OperatorSet) (bool, error) {
	return _AllocationManager.Contract.IsMemberOfOperatorSet(&_AllocationManager.CallOpts, operator, operatorSet)
}

// IsMemberOfOperatorSet is a free data retrieval call binding the contract method 0x670d3ba2.
//
// Solidity: function isMemberOfOperatorSet(address operator, (address,uint32) operatorSet) view returns(bool)
func (_AllocationManager *AllocationManagerCallerSession) IsMemberOfOperatorSet(operator common.Address, operatorSet OperatorSet) (bool, error) {
	return _AllocationManager.Contract.IsMemberOfOperatorSet(&_AllocationManager.CallOpts, operator, operatorSet)
}

// IsOperatorSet is a free data retrieval call binding the contract method 0x260dc758.
//
// Solidity: function isOperatorSet((address,uint32) operatorSet) view returns(bool)
func (_AllocationManager *AllocationManagerCaller) IsOperatorSet(opts *bind.CallOpts, operatorSet OperatorSet) (bool, error) {
	var out []interface{}
	err := _AllocationManager.contract.Call(opts, &out, "isOperatorSet", operatorSet)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperatorSet is a free data retrieval call binding the contract method 0x260dc758.
//
// Solidity: function isOperatorSet((address,uint32) operatorSet) view returns(bool)
func (_AllocationManager *AllocationManagerSession) IsOperatorSet(operatorSet OperatorSet) (bool, error) {
	return _AllocationManager.Contract.IsOperatorSet(&_AllocationManager.CallOpts, operatorSet)
}

// IsOperatorSet is a free data retrieval call binding the contract method 0x260dc758.
//
// Solidity: function isOperatorSet((address,uint32) operatorSet) view returns(bool)
func (_AllocationManager *AllocationManagerCallerSession) IsOperatorSet(operatorSet OperatorSet) (bool, error) {
	return _AllocationManager.Contract.IsOperatorSet(&_AllocationManager.CallOpts, operatorSet)
}

// IsOperatorSlashable is a free data retrieval call binding the contract method 0x1352c3e6.
//
// Solidity: function isOperatorSlashable(address operator, (address,uint32) operatorSet) view returns(bool)
func (_AllocationManager *AllocationManagerCaller) IsOperatorSlashable(opts *bind.CallOpts, operator common.Address, operatorSet OperatorSet) (bool, error) {
	var out []interface{}
	err := _AllocationManager.contract.Call(opts, &out, "isOperatorSlashable", operator, operatorSet)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperatorSlashable is a free data retrieval call binding the contract method 0x1352c3e6.
//
// Solidity: function isOperatorSlashable(address operator, (address,uint32) operatorSet) view returns(bool)
func (_AllocationManager *AllocationManagerSession) IsOperatorSlashable(operator common.Address, operatorSet OperatorSet) (bool, error) {
	return _AllocationManager.Contract.IsOperatorSlashable(&_AllocationManager.CallOpts, operator, operatorSet)
}

// IsOperatorSlashable is a free data retrieval call binding the contract method 0x1352c3e6.
//
// Solidity: function isOperatorSlashable(address operator, (address,uint32) operatorSet) view returns(bool)
func (_AllocationManager *AllocationManagerCallerSession) IsOperatorSlashable(operator common.Address, operatorSet OperatorSet) (bool, error) {
	return _AllocationManager.Contract.IsOperatorSlashable(&_AllocationManager.CallOpts, operator, operatorSet)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AllocationManager *AllocationManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AllocationManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AllocationManager *AllocationManagerSession) Owner() (common.Address, error) {
	return _AllocationManager.Contract.Owner(&_AllocationManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AllocationManager *AllocationManagerCallerSession) Owner() (common.Address, error) {
	return _AllocationManager.Contract.Owner(&_AllocationManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_AllocationManager *AllocationManagerCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _AllocationManager.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_AllocationManager *AllocationManagerSession) Paused(index uint8) (bool, error) {
	return _AllocationManager.Contract.Paused(&_AllocationManager.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_AllocationManager *AllocationManagerCallerSession) Paused(index uint8) (bool, error) {
	return _AllocationManager.Contract.Paused(&_AllocationManager.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_AllocationManager *AllocationManagerCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AllocationManager.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_AllocationManager *AllocationManagerSession) Paused0() (*big.Int, error) {
	return _AllocationManager.Contract.Paused0(&_AllocationManager.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_AllocationManager *AllocationManagerCallerSession) Paused0() (*big.Int, error) {
	return _AllocationManager.Contract.Paused0(&_AllocationManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_AllocationManager *AllocationManagerCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AllocationManager.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_AllocationManager *AllocationManagerSession) PauserRegistry() (common.Address, error) {
	return _AllocationManager.Contract.PauserRegistry(&_AllocationManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_AllocationManager *AllocationManagerCallerSession) PauserRegistry() (common.Address, error) {
	return _AllocationManager.Contract.PauserRegistry(&_AllocationManager.CallOpts)
}

// PermissionController is a free data retrieval call binding the contract method 0x4657e26a.
//
// Solidity: function permissionController() view returns(address)
func (_AllocationManager *AllocationManagerCaller) PermissionController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AllocationManager.contract.Call(opts, &out, "permissionController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PermissionController is a free data retrieval call binding the contract method 0x4657e26a.
//
// Solidity: function permissionController() view returns(address)
func (_AllocationManager *AllocationManagerSession) PermissionController() (common.Address, error) {
	return _AllocationManager.Contract.PermissionController(&_AllocationManager.CallOpts)
}

// PermissionController is a free data retrieval call binding the contract method 0x4657e26a.
//
// Solidity: function permissionController() view returns(address)
func (_AllocationManager *AllocationManagerCallerSession) PermissionController() (common.Address, error) {
	return _AllocationManager.Contract.PermissionController(&_AllocationManager.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_AllocationManager *AllocationManagerCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AllocationManager.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_AllocationManager *AllocationManagerSession) Version() (string, error) {
	return _AllocationManager.Contract.Version(&_AllocationManager.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_AllocationManager *AllocationManagerCallerSession) Version() (string, error) {
	return _AllocationManager.Contract.Version(&_AllocationManager.CallOpts)
}

// AddStrategiesToOperatorSet is a paid mutator transaction binding the contract method 0x50feea20.
//
// Solidity: function addStrategiesToOperatorSet(address avs, uint32 operatorSetId, address[] strategies) returns()
func (_AllocationManager *AllocationManagerTransactor) AddStrategiesToOperatorSet(opts *bind.TransactOpts, avs common.Address, operatorSetId uint32, strategies []common.Address) (*types.Transaction, error) {
	return _AllocationManager.contract.Transact(opts, "addStrategiesToOperatorSet", avs, operatorSetId, strategies)
}

// AddStrategiesToOperatorSet is a paid mutator transaction binding the contract method 0x50feea20.
//
// Solidity: function addStrategiesToOperatorSet(address avs, uint32 operatorSetId, address[] strategies) returns()
func (_AllocationManager *AllocationManagerSession) AddStrategiesToOperatorSet(avs common.Address, operatorSetId uint32, strategies []common.Address) (*types.Transaction, error) {
	return _AllocationManager.Contract.AddStrategiesToOperatorSet(&_AllocationManager.TransactOpts, avs, operatorSetId, strategies)
}

// AddStrategiesToOperatorSet is a paid mutator transaction binding the contract method 0x50feea20.
//
// Solidity: function addStrategiesToOperatorSet(address avs, uint32 operatorSetId, address[] strategies) returns()
func (_AllocationManager *AllocationManagerTransactorSession) AddStrategiesToOperatorSet(avs common.Address, operatorSetId uint32, strategies []common.Address) (*types.Transaction, error) {
	return _AllocationManager.Contract.AddStrategiesToOperatorSet(&_AllocationManager.TransactOpts, avs, operatorSetId, strategies)
}

// ClearDeallocationQueue is a paid mutator transaction binding the contract method 0x4b5046ef.
//
// Solidity: function clearDeallocationQueue(address operator, address[] strategies, uint16[] numToClear) returns()
func (_AllocationManager *AllocationManagerTransactor) ClearDeallocationQueue(opts *bind.TransactOpts, operator common.Address, strategies []common.Address, numToClear []uint16) (*types.Transaction, error) {
	return _AllocationManager.contract.Transact(opts, "clearDeallocationQueue", operator, strategies, numToClear)
}

// ClearDeallocationQueue is a paid mutator transaction binding the contract method 0x4b5046ef.
//
// Solidity: function clearDeallocationQueue(address operator, address[] strategies, uint16[] numToClear) returns()
func (_AllocationManager *AllocationManagerSession) ClearDeallocationQueue(operator common.Address, strategies []common.Address, numToClear []uint16) (*types.Transaction, error) {
	return _AllocationManager.Contract.ClearDeallocationQueue(&_AllocationManager.TransactOpts, operator, strategies, numToClear)
}

// ClearDeallocationQueue is a paid mutator transaction binding the contract method 0x4b5046ef.
//
// Solidity: function clearDeallocationQueue(address operator, address[] strategies, uint16[] numToClear) returns()
func (_AllocationManager *AllocationManagerTransactorSession) ClearDeallocationQueue(operator common.Address, strategies []common.Address, numToClear []uint16) (*types.Transaction, error) {
	return _AllocationManager.Contract.ClearDeallocationQueue(&_AllocationManager.TransactOpts, operator, strategies, numToClear)
}

// CreateOperatorSets is a paid mutator transaction binding the contract method 0x261f84e0.
//
// Solidity: function createOperatorSets(address avs, (uint32,address[])[] params) returns()
func (_AllocationManager *AllocationManagerTransactor) CreateOperatorSets(opts *bind.TransactOpts, avs common.Address, params []IAllocationManagerTypesCreateSetParams) (*types.Transaction, error) {
	return _AllocationManager.contract.Transact(opts, "createOperatorSets", avs, params)
}

// CreateOperatorSets is a paid mutator transaction binding the contract method 0x261f84e0.
//
// Solidity: function createOperatorSets(address avs, (uint32,address[])[] params) returns()
func (_AllocationManager *AllocationManagerSession) CreateOperatorSets(avs common.Address, params []IAllocationManagerTypesCreateSetParams) (*types.Transaction, error) {
	return _AllocationManager.Contract.CreateOperatorSets(&_AllocationManager.TransactOpts, avs, params)
}

// CreateOperatorSets is a paid mutator transaction binding the contract method 0x261f84e0.
//
// Solidity: function createOperatorSets(address avs, (uint32,address[])[] params) returns()
func (_AllocationManager *AllocationManagerTransactorSession) CreateOperatorSets(avs common.Address, params []IAllocationManagerTypesCreateSetParams) (*types.Transaction, error) {
	return _AllocationManager.Contract.CreateOperatorSets(&_AllocationManager.TransactOpts, avs, params)
}

// DeregisterFromOperatorSets is a paid mutator transaction binding the contract method 0x6e3492b5.
//
// Solidity: function deregisterFromOperatorSets((address,address,uint32[]) params) returns()
func (_AllocationManager *AllocationManagerTransactor) DeregisterFromOperatorSets(opts *bind.TransactOpts, params IAllocationManagerTypesDeregisterParams) (*types.Transaction, error) {
	return _AllocationManager.contract.Transact(opts, "deregisterFromOperatorSets", params)
}

// DeregisterFromOperatorSets is a paid mutator transaction binding the contract method 0x6e3492b5.
//
// Solidity: function deregisterFromOperatorSets((address,address,uint32[]) params) returns()
func (_AllocationManager *AllocationManagerSession) DeregisterFromOperatorSets(params IAllocationManagerTypesDeregisterParams) (*types.Transaction, error) {
	return _AllocationManager.Contract.DeregisterFromOperatorSets(&_AllocationManager.TransactOpts, params)
}

// DeregisterFromOperatorSets is a paid mutator transaction binding the contract method 0x6e3492b5.
//
// Solidity: function deregisterFromOperatorSets((address,address,uint32[]) params) returns()
func (_AllocationManager *AllocationManagerTransactorSession) DeregisterFromOperatorSets(params IAllocationManagerTypesDeregisterParams) (*types.Transaction, error) {
	return _AllocationManager.Contract.DeregisterFromOperatorSets(&_AllocationManager.TransactOpts, params)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address initialOwner, uint256 initialPausedStatus) returns()
func (_AllocationManager *AllocationManagerTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _AllocationManager.contract.Transact(opts, "initialize", initialOwner, initialPausedStatus)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address initialOwner, uint256 initialPausedStatus) returns()
func (_AllocationManager *AllocationManagerSession) Initialize(initialOwner common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _AllocationManager.Contract.Initialize(&_AllocationManager.TransactOpts, initialOwner, initialPausedStatus)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address initialOwner, uint256 initialPausedStatus) returns()
func (_AllocationManager *AllocationManagerTransactorSession) Initialize(initialOwner common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _AllocationManager.Contract.Initialize(&_AllocationManager.TransactOpts, initialOwner, initialPausedStatus)
}

// ModifyAllocations is a paid mutator transaction binding the contract method 0x952899ee.
//
// Solidity: function modifyAllocations(address operator, ((address,uint32),address[],uint64[])[] params) returns()
func (_AllocationManager *AllocationManagerTransactor) ModifyAllocations(opts *bind.TransactOpts, operator common.Address, params []IAllocationManagerTypesAllocateParams) (*types.Transaction, error) {
	return _AllocationManager.contract.Transact(opts, "modifyAllocations", operator, params)
}

// ModifyAllocations is a paid mutator transaction binding the contract method 0x952899ee.
//
// Solidity: function modifyAllocations(address operator, ((address,uint32),address[],uint64[])[] params) returns()
func (_AllocationManager *AllocationManagerSession) ModifyAllocations(operator common.Address, params []IAllocationManagerTypesAllocateParams) (*types.Transaction, error) {
	return _AllocationManager.Contract.ModifyAllocations(&_AllocationManager.TransactOpts, operator, params)
}

// ModifyAllocations is a paid mutator transaction binding the contract method 0x952899ee.
//
// Solidity: function modifyAllocations(address operator, ((address,uint32),address[],uint64[])[] params) returns()
func (_AllocationManager *AllocationManagerTransactorSession) ModifyAllocations(operator common.Address, params []IAllocationManagerTypesAllocateParams) (*types.Transaction, error) {
	return _AllocationManager.Contract.ModifyAllocations(&_AllocationManager.TransactOpts, operator, params)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_AllocationManager *AllocationManagerTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _AllocationManager.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_AllocationManager *AllocationManagerSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _AllocationManager.Contract.Pause(&_AllocationManager.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_AllocationManager *AllocationManagerTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _AllocationManager.Contract.Pause(&_AllocationManager.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_AllocationManager *AllocationManagerTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AllocationManager.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_AllocationManager *AllocationManagerSession) PauseAll() (*types.Transaction, error) {
	return _AllocationManager.Contract.PauseAll(&_AllocationManager.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_AllocationManager *AllocationManagerTransactorSession) PauseAll() (*types.Transaction, error) {
	return _AllocationManager.Contract.PauseAll(&_AllocationManager.TransactOpts)
}

// RegisterForOperatorSets is a paid mutator transaction binding the contract method 0xadc2e3d9.
//
// Solidity: function registerForOperatorSets(address operator, (address,uint32[],bytes) params) returns()
func (_AllocationManager *AllocationManagerTransactor) RegisterForOperatorSets(opts *bind.TransactOpts, operator common.Address, params IAllocationManagerTypesRegisterParams) (*types.Transaction, error) {
	return _AllocationManager.contract.Transact(opts, "registerForOperatorSets", operator, params)
}

// RegisterForOperatorSets is a paid mutator transaction binding the contract method 0xadc2e3d9.
//
// Solidity: function registerForOperatorSets(address operator, (address,uint32[],bytes) params) returns()
func (_AllocationManager *AllocationManagerSession) RegisterForOperatorSets(operator common.Address, params IAllocationManagerTypesRegisterParams) (*types.Transaction, error) {
	return _AllocationManager.Contract.RegisterForOperatorSets(&_AllocationManager.TransactOpts, operator, params)
}

// RegisterForOperatorSets is a paid mutator transaction binding the contract method 0xadc2e3d9.
//
// Solidity: function registerForOperatorSets(address operator, (address,uint32[],bytes) params) returns()
func (_AllocationManager *AllocationManagerTransactorSession) RegisterForOperatorSets(operator common.Address, params IAllocationManagerTypesRegisterParams) (*types.Transaction, error) {
	return _AllocationManager.Contract.RegisterForOperatorSets(&_AllocationManager.TransactOpts, operator, params)
}

// RemoveStrategiesFromOperatorSet is a paid mutator transaction binding the contract method 0xb66bd989.
//
// Solidity: function removeStrategiesFromOperatorSet(address avs, uint32 operatorSetId, address[] strategies) returns()
func (_AllocationManager *AllocationManagerTransactor) RemoveStrategiesFromOperatorSet(opts *bind.TransactOpts, avs common.Address, operatorSetId uint32, strategies []common.Address) (*types.Transaction, error) {
	return _AllocationManager.contract.Transact(opts, "removeStrategiesFromOperatorSet", avs, operatorSetId, strategies)
}

// RemoveStrategiesFromOperatorSet is a paid mutator transaction binding the contract method 0xb66bd989.
//
// Solidity: function removeStrategiesFromOperatorSet(address avs, uint32 operatorSetId, address[] strategies) returns()
func (_AllocationManager *AllocationManagerSession) RemoveStrategiesFromOperatorSet(avs common.Address, operatorSetId uint32, strategies []common.Address) (*types.Transaction, error) {
	return _AllocationManager.Contract.RemoveStrategiesFromOperatorSet(&_AllocationManager.TransactOpts, avs, operatorSetId, strategies)
}

// RemoveStrategiesFromOperatorSet is a paid mutator transaction binding the contract method 0xb66bd989.
//
// Solidity: function removeStrategiesFromOperatorSet(address avs, uint32 operatorSetId, address[] strategies) returns()
func (_AllocationManager *AllocationManagerTransactorSession) RemoveStrategiesFromOperatorSet(avs common.Address, operatorSetId uint32, strategies []common.Address) (*types.Transaction, error) {
	return _AllocationManager.Contract.RemoveStrategiesFromOperatorSet(&_AllocationManager.TransactOpts, avs, operatorSetId, strategies)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AllocationManager *AllocationManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AllocationManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AllocationManager *AllocationManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _AllocationManager.Contract.RenounceOwnership(&_AllocationManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AllocationManager *AllocationManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AllocationManager.Contract.RenounceOwnership(&_AllocationManager.TransactOpts)
}

// SetAVSRegistrar is a paid mutator transaction binding the contract method 0xd3d96ff4.
//
// Solidity: function setAVSRegistrar(address avs, address registrar) returns()
func (_AllocationManager *AllocationManagerTransactor) SetAVSRegistrar(opts *bind.TransactOpts, avs common.Address, registrar common.Address) (*types.Transaction, error) {
	return _AllocationManager.contract.Transact(opts, "setAVSRegistrar", avs, registrar)
}

// SetAVSRegistrar is a paid mutator transaction binding the contract method 0xd3d96ff4.
//
// Solidity: function setAVSRegistrar(address avs, address registrar) returns()
func (_AllocationManager *AllocationManagerSession) SetAVSRegistrar(avs common.Address, registrar common.Address) (*types.Transaction, error) {
	return _AllocationManager.Contract.SetAVSRegistrar(&_AllocationManager.TransactOpts, avs, registrar)
}

// SetAVSRegistrar is a paid mutator transaction binding the contract method 0xd3d96ff4.
//
// Solidity: function setAVSRegistrar(address avs, address registrar) returns()
func (_AllocationManager *AllocationManagerTransactorSession) SetAVSRegistrar(avs common.Address, registrar common.Address) (*types.Transaction, error) {
	return _AllocationManager.Contract.SetAVSRegistrar(&_AllocationManager.TransactOpts, avs, registrar)
}

// SetAllocationDelay is a paid mutator transaction binding the contract method 0x56c483e6.
//
// Solidity: function setAllocationDelay(address operator, uint32 delay) returns()
func (_AllocationManager *AllocationManagerTransactor) SetAllocationDelay(opts *bind.TransactOpts, operator common.Address, delay uint32) (*types.Transaction, error) {
	return _AllocationManager.contract.Transact(opts, "setAllocationDelay", operator, delay)
}

// SetAllocationDelay is a paid mutator transaction binding the contract method 0x56c483e6.
//
// Solidity: function setAllocationDelay(address operator, uint32 delay) returns()
func (_AllocationManager *AllocationManagerSession) SetAllocationDelay(operator common.Address, delay uint32) (*types.Transaction, error) {
	return _AllocationManager.Contract.SetAllocationDelay(&_AllocationManager.TransactOpts, operator, delay)
}

// SetAllocationDelay is a paid mutator transaction binding the contract method 0x56c483e6.
//
// Solidity: function setAllocationDelay(address operator, uint32 delay) returns()
func (_AllocationManager *AllocationManagerTransactorSession) SetAllocationDelay(operator common.Address, delay uint32) (*types.Transaction, error) {
	return _AllocationManager.Contract.SetAllocationDelay(&_AllocationManager.TransactOpts, operator, delay)
}

// SlashOperator is a paid mutator transaction binding the contract method 0x36352057.
//
// Solidity: function slashOperator(address avs, (address,uint32,address[],uint256[],string) params) returns()
func (_AllocationManager *AllocationManagerTransactor) SlashOperator(opts *bind.TransactOpts, avs common.Address, params IAllocationManagerTypesSlashingParams) (*types.Transaction, error) {
	return _AllocationManager.contract.Transact(opts, "slashOperator", avs, params)
}

// SlashOperator is a paid mutator transaction binding the contract method 0x36352057.
//
// Solidity: function slashOperator(address avs, (address,uint32,address[],uint256[],string) params) returns()
func (_AllocationManager *AllocationManagerSession) SlashOperator(avs common.Address, params IAllocationManagerTypesSlashingParams) (*types.Transaction, error) {
	return _AllocationManager.Contract.SlashOperator(&_AllocationManager.TransactOpts, avs, params)
}

// SlashOperator is a paid mutator transaction binding the contract method 0x36352057.
//
// Solidity: function slashOperator(address avs, (address,uint32,address[],uint256[],string) params) returns()
func (_AllocationManager *AllocationManagerTransactorSession) SlashOperator(avs common.Address, params IAllocationManagerTypesSlashingParams) (*types.Transaction, error) {
	return _AllocationManager.Contract.SlashOperator(&_AllocationManager.TransactOpts, avs, params)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AllocationManager *AllocationManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AllocationManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AllocationManager *AllocationManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AllocationManager.Contract.TransferOwnership(&_AllocationManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AllocationManager *AllocationManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AllocationManager.Contract.TransferOwnership(&_AllocationManager.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_AllocationManager *AllocationManagerTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _AllocationManager.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_AllocationManager *AllocationManagerSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _AllocationManager.Contract.Unpause(&_AllocationManager.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_AllocationManager *AllocationManagerTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _AllocationManager.Contract.Unpause(&_AllocationManager.TransactOpts, newPausedStatus)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa9821821.
//
// Solidity: function updateAVSMetadataURI(address avs, string metadataURI) returns()
func (_AllocationManager *AllocationManagerTransactor) UpdateAVSMetadataURI(opts *bind.TransactOpts, avs common.Address, metadataURI string) (*types.Transaction, error) {
	return _AllocationManager.contract.Transact(opts, "updateAVSMetadataURI", avs, metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa9821821.
//
// Solidity: function updateAVSMetadataURI(address avs, string metadataURI) returns()
func (_AllocationManager *AllocationManagerSession) UpdateAVSMetadataURI(avs common.Address, metadataURI string) (*types.Transaction, error) {
	return _AllocationManager.Contract.UpdateAVSMetadataURI(&_AllocationManager.TransactOpts, avs, metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa9821821.
//
// Solidity: function updateAVSMetadataURI(address avs, string metadataURI) returns()
func (_AllocationManager *AllocationManagerTransactorSession) UpdateAVSMetadataURI(avs common.Address, metadataURI string) (*types.Transaction, error) {
	return _AllocationManager.Contract.UpdateAVSMetadataURI(&_AllocationManager.TransactOpts, avs, metadataURI)
}

// AllocationManagerAVSMetadataURIUpdatedIterator is returned from FilterAVSMetadataURIUpdated and is used to iterate over the raw logs and unpacked data for AVSMetadataURIUpdated events raised by the AllocationManager contract.
type AllocationManagerAVSMetadataURIUpdatedIterator struct {
	Event *AllocationManagerAVSMetadataURIUpdated // Event containing the contract specifics and raw log

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
func (it *AllocationManagerAVSMetadataURIUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AllocationManagerAVSMetadataURIUpdated)
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
		it.Event = new(AllocationManagerAVSMetadataURIUpdated)
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
func (it *AllocationManagerAVSMetadataURIUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AllocationManagerAVSMetadataURIUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AllocationManagerAVSMetadataURIUpdated represents a AVSMetadataURIUpdated event raised by the AllocationManager contract.
type AllocationManagerAVSMetadataURIUpdated struct {
	Avs         common.Address
	MetadataURI string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAVSMetadataURIUpdated is a free log retrieval operation binding the contract event 0xa89c1dc243d8908a96dd84944bcc97d6bc6ac00dd78e20621576be6a3c943713.
//
// Solidity: event AVSMetadataURIUpdated(address indexed avs, string metadataURI)
func (_AllocationManager *AllocationManagerFilterer) FilterAVSMetadataURIUpdated(opts *bind.FilterOpts, avs []common.Address) (*AllocationManagerAVSMetadataURIUpdatedIterator, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _AllocationManager.contract.FilterLogs(opts, "AVSMetadataURIUpdated", avsRule)
	if err != nil {
		return nil, err
	}
	return &AllocationManagerAVSMetadataURIUpdatedIterator{contract: _AllocationManager.contract, event: "AVSMetadataURIUpdated", logs: logs, sub: sub}, nil
}

// WatchAVSMetadataURIUpdated is a free log subscription operation binding the contract event 0xa89c1dc243d8908a96dd84944bcc97d6bc6ac00dd78e20621576be6a3c943713.
//
// Solidity: event AVSMetadataURIUpdated(address indexed avs, string metadataURI)
func (_AllocationManager *AllocationManagerFilterer) WatchAVSMetadataURIUpdated(opts *bind.WatchOpts, sink chan<- *AllocationManagerAVSMetadataURIUpdated, avs []common.Address) (event.Subscription, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _AllocationManager.contract.WatchLogs(opts, "AVSMetadataURIUpdated", avsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AllocationManagerAVSMetadataURIUpdated)
				if err := _AllocationManager.contract.UnpackLog(event, "AVSMetadataURIUpdated", log); err != nil {
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
func (_AllocationManager *AllocationManagerFilterer) ParseAVSMetadataURIUpdated(log types.Log) (*AllocationManagerAVSMetadataURIUpdated, error) {
	event := new(AllocationManagerAVSMetadataURIUpdated)
	if err := _AllocationManager.contract.UnpackLog(event, "AVSMetadataURIUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AllocationManagerAVSRegistrarSetIterator is returned from FilterAVSRegistrarSet and is used to iterate over the raw logs and unpacked data for AVSRegistrarSet events raised by the AllocationManager contract.
type AllocationManagerAVSRegistrarSetIterator struct {
	Event *AllocationManagerAVSRegistrarSet // Event containing the contract specifics and raw log

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
func (it *AllocationManagerAVSRegistrarSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AllocationManagerAVSRegistrarSet)
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
		it.Event = new(AllocationManagerAVSRegistrarSet)
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
func (it *AllocationManagerAVSRegistrarSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AllocationManagerAVSRegistrarSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AllocationManagerAVSRegistrarSet represents a AVSRegistrarSet event raised by the AllocationManager contract.
type AllocationManagerAVSRegistrarSet struct {
	Avs       common.Address
	Registrar common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAVSRegistrarSet is a free log retrieval operation binding the contract event 0x2ae945c40c44dc0ec263f95609c3fdc6952e0aefa22d6374e44f2c997acedf85.
//
// Solidity: event AVSRegistrarSet(address avs, address registrar)
func (_AllocationManager *AllocationManagerFilterer) FilterAVSRegistrarSet(opts *bind.FilterOpts) (*AllocationManagerAVSRegistrarSetIterator, error) {

	logs, sub, err := _AllocationManager.contract.FilterLogs(opts, "AVSRegistrarSet")
	if err != nil {
		return nil, err
	}
	return &AllocationManagerAVSRegistrarSetIterator{contract: _AllocationManager.contract, event: "AVSRegistrarSet", logs: logs, sub: sub}, nil
}

// WatchAVSRegistrarSet is a free log subscription operation binding the contract event 0x2ae945c40c44dc0ec263f95609c3fdc6952e0aefa22d6374e44f2c997acedf85.
//
// Solidity: event AVSRegistrarSet(address avs, address registrar)
func (_AllocationManager *AllocationManagerFilterer) WatchAVSRegistrarSet(opts *bind.WatchOpts, sink chan<- *AllocationManagerAVSRegistrarSet) (event.Subscription, error) {

	logs, sub, err := _AllocationManager.contract.WatchLogs(opts, "AVSRegistrarSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AllocationManagerAVSRegistrarSet)
				if err := _AllocationManager.contract.UnpackLog(event, "AVSRegistrarSet", log); err != nil {
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

// ParseAVSRegistrarSet is a log parse operation binding the contract event 0x2ae945c40c44dc0ec263f95609c3fdc6952e0aefa22d6374e44f2c997acedf85.
//
// Solidity: event AVSRegistrarSet(address avs, address registrar)
func (_AllocationManager *AllocationManagerFilterer) ParseAVSRegistrarSet(log types.Log) (*AllocationManagerAVSRegistrarSet, error) {
	event := new(AllocationManagerAVSRegistrarSet)
	if err := _AllocationManager.contract.UnpackLog(event, "AVSRegistrarSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AllocationManagerAllocationDelaySetIterator is returned from FilterAllocationDelaySet and is used to iterate over the raw logs and unpacked data for AllocationDelaySet events raised by the AllocationManager contract.
type AllocationManagerAllocationDelaySetIterator struct {
	Event *AllocationManagerAllocationDelaySet // Event containing the contract specifics and raw log

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
func (it *AllocationManagerAllocationDelaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AllocationManagerAllocationDelaySet)
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
		it.Event = new(AllocationManagerAllocationDelaySet)
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
func (it *AllocationManagerAllocationDelaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AllocationManagerAllocationDelaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AllocationManagerAllocationDelaySet represents a AllocationDelaySet event raised by the AllocationManager contract.
type AllocationManagerAllocationDelaySet struct {
	Operator    common.Address
	Delay       uint32
	EffectBlock uint32
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAllocationDelaySet is a free log retrieval operation binding the contract event 0x4e85751d6331506c6c62335f207eb31f12a61e570f34f5c17640308785c6d4db.
//
// Solidity: event AllocationDelaySet(address operator, uint32 delay, uint32 effectBlock)
func (_AllocationManager *AllocationManagerFilterer) FilterAllocationDelaySet(opts *bind.FilterOpts) (*AllocationManagerAllocationDelaySetIterator, error) {

	logs, sub, err := _AllocationManager.contract.FilterLogs(opts, "AllocationDelaySet")
	if err != nil {
		return nil, err
	}
	return &AllocationManagerAllocationDelaySetIterator{contract: _AllocationManager.contract, event: "AllocationDelaySet", logs: logs, sub: sub}, nil
}

// WatchAllocationDelaySet is a free log subscription operation binding the contract event 0x4e85751d6331506c6c62335f207eb31f12a61e570f34f5c17640308785c6d4db.
//
// Solidity: event AllocationDelaySet(address operator, uint32 delay, uint32 effectBlock)
func (_AllocationManager *AllocationManagerFilterer) WatchAllocationDelaySet(opts *bind.WatchOpts, sink chan<- *AllocationManagerAllocationDelaySet) (event.Subscription, error) {

	logs, sub, err := _AllocationManager.contract.WatchLogs(opts, "AllocationDelaySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AllocationManagerAllocationDelaySet)
				if err := _AllocationManager.contract.UnpackLog(event, "AllocationDelaySet", log); err != nil {
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

// ParseAllocationDelaySet is a log parse operation binding the contract event 0x4e85751d6331506c6c62335f207eb31f12a61e570f34f5c17640308785c6d4db.
//
// Solidity: event AllocationDelaySet(address operator, uint32 delay, uint32 effectBlock)
func (_AllocationManager *AllocationManagerFilterer) ParseAllocationDelaySet(log types.Log) (*AllocationManagerAllocationDelaySet, error) {
	event := new(AllocationManagerAllocationDelaySet)
	if err := _AllocationManager.contract.UnpackLog(event, "AllocationDelaySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AllocationManagerAllocationUpdatedIterator is returned from FilterAllocationUpdated and is used to iterate over the raw logs and unpacked data for AllocationUpdated events raised by the AllocationManager contract.
type AllocationManagerAllocationUpdatedIterator struct {
	Event *AllocationManagerAllocationUpdated // Event containing the contract specifics and raw log

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
func (it *AllocationManagerAllocationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AllocationManagerAllocationUpdated)
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
		it.Event = new(AllocationManagerAllocationUpdated)
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
func (it *AllocationManagerAllocationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AllocationManagerAllocationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AllocationManagerAllocationUpdated represents a AllocationUpdated event raised by the AllocationManager contract.
type AllocationManagerAllocationUpdated struct {
	Operator    common.Address
	OperatorSet OperatorSet
	Strategy    common.Address
	Magnitude   uint64
	EffectBlock uint32
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAllocationUpdated is a free log retrieval operation binding the contract event 0x1487af5418c47ee5ea45ef4a93398668120890774a9e13487e61e9dc3baf76dd.
//
// Solidity: event AllocationUpdated(address operator, (address,uint32) operatorSet, address strategy, uint64 magnitude, uint32 effectBlock)
func (_AllocationManager *AllocationManagerFilterer) FilterAllocationUpdated(opts *bind.FilterOpts) (*AllocationManagerAllocationUpdatedIterator, error) {

	logs, sub, err := _AllocationManager.contract.FilterLogs(opts, "AllocationUpdated")
	if err != nil {
		return nil, err
	}
	return &AllocationManagerAllocationUpdatedIterator{contract: _AllocationManager.contract, event: "AllocationUpdated", logs: logs, sub: sub}, nil
}

// WatchAllocationUpdated is a free log subscription operation binding the contract event 0x1487af5418c47ee5ea45ef4a93398668120890774a9e13487e61e9dc3baf76dd.
//
// Solidity: event AllocationUpdated(address operator, (address,uint32) operatorSet, address strategy, uint64 magnitude, uint32 effectBlock)
func (_AllocationManager *AllocationManagerFilterer) WatchAllocationUpdated(opts *bind.WatchOpts, sink chan<- *AllocationManagerAllocationUpdated) (event.Subscription, error) {

	logs, sub, err := _AllocationManager.contract.WatchLogs(opts, "AllocationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AllocationManagerAllocationUpdated)
				if err := _AllocationManager.contract.UnpackLog(event, "AllocationUpdated", log); err != nil {
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

// ParseAllocationUpdated is a log parse operation binding the contract event 0x1487af5418c47ee5ea45ef4a93398668120890774a9e13487e61e9dc3baf76dd.
//
// Solidity: event AllocationUpdated(address operator, (address,uint32) operatorSet, address strategy, uint64 magnitude, uint32 effectBlock)
func (_AllocationManager *AllocationManagerFilterer) ParseAllocationUpdated(log types.Log) (*AllocationManagerAllocationUpdated, error) {
	event := new(AllocationManagerAllocationUpdated)
	if err := _AllocationManager.contract.UnpackLog(event, "AllocationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AllocationManagerEncumberedMagnitudeUpdatedIterator is returned from FilterEncumberedMagnitudeUpdated and is used to iterate over the raw logs and unpacked data for EncumberedMagnitudeUpdated events raised by the AllocationManager contract.
type AllocationManagerEncumberedMagnitudeUpdatedIterator struct {
	Event *AllocationManagerEncumberedMagnitudeUpdated // Event containing the contract specifics and raw log

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
func (it *AllocationManagerEncumberedMagnitudeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AllocationManagerEncumberedMagnitudeUpdated)
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
		it.Event = new(AllocationManagerEncumberedMagnitudeUpdated)
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
func (it *AllocationManagerEncumberedMagnitudeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AllocationManagerEncumberedMagnitudeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AllocationManagerEncumberedMagnitudeUpdated represents a EncumberedMagnitudeUpdated event raised by the AllocationManager contract.
type AllocationManagerEncumberedMagnitudeUpdated struct {
	Operator            common.Address
	Strategy            common.Address
	EncumberedMagnitude uint64
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterEncumberedMagnitudeUpdated is a free log retrieval operation binding the contract event 0xacf9095feb3a370c9cf692421c69ef320d4db5c66e6a7d29c7694eb02364fc55.
//
// Solidity: event EncumberedMagnitudeUpdated(address operator, address strategy, uint64 encumberedMagnitude)
func (_AllocationManager *AllocationManagerFilterer) FilterEncumberedMagnitudeUpdated(opts *bind.FilterOpts) (*AllocationManagerEncumberedMagnitudeUpdatedIterator, error) {

	logs, sub, err := _AllocationManager.contract.FilterLogs(opts, "EncumberedMagnitudeUpdated")
	if err != nil {
		return nil, err
	}
	return &AllocationManagerEncumberedMagnitudeUpdatedIterator{contract: _AllocationManager.contract, event: "EncumberedMagnitudeUpdated", logs: logs, sub: sub}, nil
}

// WatchEncumberedMagnitudeUpdated is a free log subscription operation binding the contract event 0xacf9095feb3a370c9cf692421c69ef320d4db5c66e6a7d29c7694eb02364fc55.
//
// Solidity: event EncumberedMagnitudeUpdated(address operator, address strategy, uint64 encumberedMagnitude)
func (_AllocationManager *AllocationManagerFilterer) WatchEncumberedMagnitudeUpdated(opts *bind.WatchOpts, sink chan<- *AllocationManagerEncumberedMagnitudeUpdated) (event.Subscription, error) {

	logs, sub, err := _AllocationManager.contract.WatchLogs(opts, "EncumberedMagnitudeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AllocationManagerEncumberedMagnitudeUpdated)
				if err := _AllocationManager.contract.UnpackLog(event, "EncumberedMagnitudeUpdated", log); err != nil {
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

// ParseEncumberedMagnitudeUpdated is a log parse operation binding the contract event 0xacf9095feb3a370c9cf692421c69ef320d4db5c66e6a7d29c7694eb02364fc55.
//
// Solidity: event EncumberedMagnitudeUpdated(address operator, address strategy, uint64 encumberedMagnitude)
func (_AllocationManager *AllocationManagerFilterer) ParseEncumberedMagnitudeUpdated(log types.Log) (*AllocationManagerEncumberedMagnitudeUpdated, error) {
	event := new(AllocationManagerEncumberedMagnitudeUpdated)
	if err := _AllocationManager.contract.UnpackLog(event, "EncumberedMagnitudeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AllocationManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the AllocationManager contract.
type AllocationManagerInitializedIterator struct {
	Event *AllocationManagerInitialized // Event containing the contract specifics and raw log

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
func (it *AllocationManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AllocationManagerInitialized)
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
		it.Event = new(AllocationManagerInitialized)
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
func (it *AllocationManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AllocationManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AllocationManagerInitialized represents a Initialized event raised by the AllocationManager contract.
type AllocationManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AllocationManager *AllocationManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*AllocationManagerInitializedIterator, error) {

	logs, sub, err := _AllocationManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AllocationManagerInitializedIterator{contract: _AllocationManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AllocationManager *AllocationManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AllocationManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _AllocationManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AllocationManagerInitialized)
				if err := _AllocationManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_AllocationManager *AllocationManagerFilterer) ParseInitialized(log types.Log) (*AllocationManagerInitialized, error) {
	event := new(AllocationManagerInitialized)
	if err := _AllocationManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AllocationManagerMaxMagnitudeUpdatedIterator is returned from FilterMaxMagnitudeUpdated and is used to iterate over the raw logs and unpacked data for MaxMagnitudeUpdated events raised by the AllocationManager contract.
type AllocationManagerMaxMagnitudeUpdatedIterator struct {
	Event *AllocationManagerMaxMagnitudeUpdated // Event containing the contract specifics and raw log

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
func (it *AllocationManagerMaxMagnitudeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AllocationManagerMaxMagnitudeUpdated)
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
		it.Event = new(AllocationManagerMaxMagnitudeUpdated)
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
func (it *AllocationManagerMaxMagnitudeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AllocationManagerMaxMagnitudeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AllocationManagerMaxMagnitudeUpdated represents a MaxMagnitudeUpdated event raised by the AllocationManager contract.
type AllocationManagerMaxMagnitudeUpdated struct {
	Operator     common.Address
	Strategy     common.Address
	MaxMagnitude uint64
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMaxMagnitudeUpdated is a free log retrieval operation binding the contract event 0x1c6458079a41077d003c11faf9bf097e693bd67979e4e6500bac7b29db779b5c.
//
// Solidity: event MaxMagnitudeUpdated(address operator, address strategy, uint64 maxMagnitude)
func (_AllocationManager *AllocationManagerFilterer) FilterMaxMagnitudeUpdated(opts *bind.FilterOpts) (*AllocationManagerMaxMagnitudeUpdatedIterator, error) {

	logs, sub, err := _AllocationManager.contract.FilterLogs(opts, "MaxMagnitudeUpdated")
	if err != nil {
		return nil, err
	}
	return &AllocationManagerMaxMagnitudeUpdatedIterator{contract: _AllocationManager.contract, event: "MaxMagnitudeUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxMagnitudeUpdated is a free log subscription operation binding the contract event 0x1c6458079a41077d003c11faf9bf097e693bd67979e4e6500bac7b29db779b5c.
//
// Solidity: event MaxMagnitudeUpdated(address operator, address strategy, uint64 maxMagnitude)
func (_AllocationManager *AllocationManagerFilterer) WatchMaxMagnitudeUpdated(opts *bind.WatchOpts, sink chan<- *AllocationManagerMaxMagnitudeUpdated) (event.Subscription, error) {

	logs, sub, err := _AllocationManager.contract.WatchLogs(opts, "MaxMagnitudeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AllocationManagerMaxMagnitudeUpdated)
				if err := _AllocationManager.contract.UnpackLog(event, "MaxMagnitudeUpdated", log); err != nil {
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

// ParseMaxMagnitudeUpdated is a log parse operation binding the contract event 0x1c6458079a41077d003c11faf9bf097e693bd67979e4e6500bac7b29db779b5c.
//
// Solidity: event MaxMagnitudeUpdated(address operator, address strategy, uint64 maxMagnitude)
func (_AllocationManager *AllocationManagerFilterer) ParseMaxMagnitudeUpdated(log types.Log) (*AllocationManagerMaxMagnitudeUpdated, error) {
	event := new(AllocationManagerMaxMagnitudeUpdated)
	if err := _AllocationManager.contract.UnpackLog(event, "MaxMagnitudeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AllocationManagerOperatorAddedToOperatorSetIterator is returned from FilterOperatorAddedToOperatorSet and is used to iterate over the raw logs and unpacked data for OperatorAddedToOperatorSet events raised by the AllocationManager contract.
type AllocationManagerOperatorAddedToOperatorSetIterator struct {
	Event *AllocationManagerOperatorAddedToOperatorSet // Event containing the contract specifics and raw log

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
func (it *AllocationManagerOperatorAddedToOperatorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AllocationManagerOperatorAddedToOperatorSet)
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
		it.Event = new(AllocationManagerOperatorAddedToOperatorSet)
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
func (it *AllocationManagerOperatorAddedToOperatorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AllocationManagerOperatorAddedToOperatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AllocationManagerOperatorAddedToOperatorSet represents a OperatorAddedToOperatorSet event raised by the AllocationManager contract.
type AllocationManagerOperatorAddedToOperatorSet struct {
	Operator    common.Address
	OperatorSet OperatorSet
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorAddedToOperatorSet is a free log retrieval operation binding the contract event 0x43232edf9071753d2321e5fa7e018363ee248e5f2142e6c08edd3265bfb4895e.
//
// Solidity: event OperatorAddedToOperatorSet(address indexed operator, (address,uint32) operatorSet)
func (_AllocationManager *AllocationManagerFilterer) FilterOperatorAddedToOperatorSet(opts *bind.FilterOpts, operator []common.Address) (*AllocationManagerOperatorAddedToOperatorSetIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _AllocationManager.contract.FilterLogs(opts, "OperatorAddedToOperatorSet", operatorRule)
	if err != nil {
		return nil, err
	}
	return &AllocationManagerOperatorAddedToOperatorSetIterator{contract: _AllocationManager.contract, event: "OperatorAddedToOperatorSet", logs: logs, sub: sub}, nil
}

// WatchOperatorAddedToOperatorSet is a free log subscription operation binding the contract event 0x43232edf9071753d2321e5fa7e018363ee248e5f2142e6c08edd3265bfb4895e.
//
// Solidity: event OperatorAddedToOperatorSet(address indexed operator, (address,uint32) operatorSet)
func (_AllocationManager *AllocationManagerFilterer) WatchOperatorAddedToOperatorSet(opts *bind.WatchOpts, sink chan<- *AllocationManagerOperatorAddedToOperatorSet, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _AllocationManager.contract.WatchLogs(opts, "OperatorAddedToOperatorSet", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AllocationManagerOperatorAddedToOperatorSet)
				if err := _AllocationManager.contract.UnpackLog(event, "OperatorAddedToOperatorSet", log); err != nil {
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
func (_AllocationManager *AllocationManagerFilterer) ParseOperatorAddedToOperatorSet(log types.Log) (*AllocationManagerOperatorAddedToOperatorSet, error) {
	event := new(AllocationManagerOperatorAddedToOperatorSet)
	if err := _AllocationManager.contract.UnpackLog(event, "OperatorAddedToOperatorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AllocationManagerOperatorRemovedFromOperatorSetIterator is returned from FilterOperatorRemovedFromOperatorSet and is used to iterate over the raw logs and unpacked data for OperatorRemovedFromOperatorSet events raised by the AllocationManager contract.
type AllocationManagerOperatorRemovedFromOperatorSetIterator struct {
	Event *AllocationManagerOperatorRemovedFromOperatorSet // Event containing the contract specifics and raw log

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
func (it *AllocationManagerOperatorRemovedFromOperatorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AllocationManagerOperatorRemovedFromOperatorSet)
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
		it.Event = new(AllocationManagerOperatorRemovedFromOperatorSet)
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
func (it *AllocationManagerOperatorRemovedFromOperatorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AllocationManagerOperatorRemovedFromOperatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AllocationManagerOperatorRemovedFromOperatorSet represents a OperatorRemovedFromOperatorSet event raised by the AllocationManager contract.
type AllocationManagerOperatorRemovedFromOperatorSet struct {
	Operator    common.Address
	OperatorSet OperatorSet
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorRemovedFromOperatorSet is a free log retrieval operation binding the contract event 0xad34c3070be1dffbcaa499d000ba2b8d9848aefcac3059df245dd95c4ece14fe.
//
// Solidity: event OperatorRemovedFromOperatorSet(address indexed operator, (address,uint32) operatorSet)
func (_AllocationManager *AllocationManagerFilterer) FilterOperatorRemovedFromOperatorSet(opts *bind.FilterOpts, operator []common.Address) (*AllocationManagerOperatorRemovedFromOperatorSetIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _AllocationManager.contract.FilterLogs(opts, "OperatorRemovedFromOperatorSet", operatorRule)
	if err != nil {
		return nil, err
	}
	return &AllocationManagerOperatorRemovedFromOperatorSetIterator{contract: _AllocationManager.contract, event: "OperatorRemovedFromOperatorSet", logs: logs, sub: sub}, nil
}

// WatchOperatorRemovedFromOperatorSet is a free log subscription operation binding the contract event 0xad34c3070be1dffbcaa499d000ba2b8d9848aefcac3059df245dd95c4ece14fe.
//
// Solidity: event OperatorRemovedFromOperatorSet(address indexed operator, (address,uint32) operatorSet)
func (_AllocationManager *AllocationManagerFilterer) WatchOperatorRemovedFromOperatorSet(opts *bind.WatchOpts, sink chan<- *AllocationManagerOperatorRemovedFromOperatorSet, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _AllocationManager.contract.WatchLogs(opts, "OperatorRemovedFromOperatorSet", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AllocationManagerOperatorRemovedFromOperatorSet)
				if err := _AllocationManager.contract.UnpackLog(event, "OperatorRemovedFromOperatorSet", log); err != nil {
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
func (_AllocationManager *AllocationManagerFilterer) ParseOperatorRemovedFromOperatorSet(log types.Log) (*AllocationManagerOperatorRemovedFromOperatorSet, error) {
	event := new(AllocationManagerOperatorRemovedFromOperatorSet)
	if err := _AllocationManager.contract.UnpackLog(event, "OperatorRemovedFromOperatorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AllocationManagerOperatorSetCreatedIterator is returned from FilterOperatorSetCreated and is used to iterate over the raw logs and unpacked data for OperatorSetCreated events raised by the AllocationManager contract.
type AllocationManagerOperatorSetCreatedIterator struct {
	Event *AllocationManagerOperatorSetCreated // Event containing the contract specifics and raw log

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
func (it *AllocationManagerOperatorSetCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AllocationManagerOperatorSetCreated)
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
		it.Event = new(AllocationManagerOperatorSetCreated)
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
func (it *AllocationManagerOperatorSetCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AllocationManagerOperatorSetCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AllocationManagerOperatorSetCreated represents a OperatorSetCreated event raised by the AllocationManager contract.
type AllocationManagerOperatorSetCreated struct {
	OperatorSet OperatorSet
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorSetCreated is a free log retrieval operation binding the contract event 0x31629285ead2335ae0933f86ed2ae63321f7af77b4e6eaabc42c057880977e6c.
//
// Solidity: event OperatorSetCreated((address,uint32) operatorSet)
func (_AllocationManager *AllocationManagerFilterer) FilterOperatorSetCreated(opts *bind.FilterOpts) (*AllocationManagerOperatorSetCreatedIterator, error) {

	logs, sub, err := _AllocationManager.contract.FilterLogs(opts, "OperatorSetCreated")
	if err != nil {
		return nil, err
	}
	return &AllocationManagerOperatorSetCreatedIterator{contract: _AllocationManager.contract, event: "OperatorSetCreated", logs: logs, sub: sub}, nil
}

// WatchOperatorSetCreated is a free log subscription operation binding the contract event 0x31629285ead2335ae0933f86ed2ae63321f7af77b4e6eaabc42c057880977e6c.
//
// Solidity: event OperatorSetCreated((address,uint32) operatorSet)
func (_AllocationManager *AllocationManagerFilterer) WatchOperatorSetCreated(opts *bind.WatchOpts, sink chan<- *AllocationManagerOperatorSetCreated) (event.Subscription, error) {

	logs, sub, err := _AllocationManager.contract.WatchLogs(opts, "OperatorSetCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AllocationManagerOperatorSetCreated)
				if err := _AllocationManager.contract.UnpackLog(event, "OperatorSetCreated", log); err != nil {
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
func (_AllocationManager *AllocationManagerFilterer) ParseOperatorSetCreated(log types.Log) (*AllocationManagerOperatorSetCreated, error) {
	event := new(AllocationManagerOperatorSetCreated)
	if err := _AllocationManager.contract.UnpackLog(event, "OperatorSetCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AllocationManagerOperatorSlashedIterator is returned from FilterOperatorSlashed and is used to iterate over the raw logs and unpacked data for OperatorSlashed events raised by the AllocationManager contract.
type AllocationManagerOperatorSlashedIterator struct {
	Event *AllocationManagerOperatorSlashed // Event containing the contract specifics and raw log

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
func (it *AllocationManagerOperatorSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AllocationManagerOperatorSlashed)
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
		it.Event = new(AllocationManagerOperatorSlashed)
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
func (it *AllocationManagerOperatorSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AllocationManagerOperatorSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AllocationManagerOperatorSlashed represents a OperatorSlashed event raised by the AllocationManager contract.
type AllocationManagerOperatorSlashed struct {
	Operator    common.Address
	OperatorSet OperatorSet
	Strategies  []common.Address
	WadSlashed  []*big.Int
	Description string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorSlashed is a free log retrieval operation binding the contract event 0x80969ad29428d6797ee7aad084f9e4a42a82fc506dcd2ca3b6fb431f85ccebe5.
//
// Solidity: event OperatorSlashed(address operator, (address,uint32) operatorSet, address[] strategies, uint256[] wadSlashed, string description)
func (_AllocationManager *AllocationManagerFilterer) FilterOperatorSlashed(opts *bind.FilterOpts) (*AllocationManagerOperatorSlashedIterator, error) {

	logs, sub, err := _AllocationManager.contract.FilterLogs(opts, "OperatorSlashed")
	if err != nil {
		return nil, err
	}
	return &AllocationManagerOperatorSlashedIterator{contract: _AllocationManager.contract, event: "OperatorSlashed", logs: logs, sub: sub}, nil
}

// WatchOperatorSlashed is a free log subscription operation binding the contract event 0x80969ad29428d6797ee7aad084f9e4a42a82fc506dcd2ca3b6fb431f85ccebe5.
//
// Solidity: event OperatorSlashed(address operator, (address,uint32) operatorSet, address[] strategies, uint256[] wadSlashed, string description)
func (_AllocationManager *AllocationManagerFilterer) WatchOperatorSlashed(opts *bind.WatchOpts, sink chan<- *AllocationManagerOperatorSlashed) (event.Subscription, error) {

	logs, sub, err := _AllocationManager.contract.WatchLogs(opts, "OperatorSlashed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AllocationManagerOperatorSlashed)
				if err := _AllocationManager.contract.UnpackLog(event, "OperatorSlashed", log); err != nil {
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

// ParseOperatorSlashed is a log parse operation binding the contract event 0x80969ad29428d6797ee7aad084f9e4a42a82fc506dcd2ca3b6fb431f85ccebe5.
//
// Solidity: event OperatorSlashed(address operator, (address,uint32) operatorSet, address[] strategies, uint256[] wadSlashed, string description)
func (_AllocationManager *AllocationManagerFilterer) ParseOperatorSlashed(log types.Log) (*AllocationManagerOperatorSlashed, error) {
	event := new(AllocationManagerOperatorSlashed)
	if err := _AllocationManager.contract.UnpackLog(event, "OperatorSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AllocationManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AllocationManager contract.
type AllocationManagerOwnershipTransferredIterator struct {
	Event *AllocationManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AllocationManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AllocationManagerOwnershipTransferred)
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
		it.Event = new(AllocationManagerOwnershipTransferred)
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
func (it *AllocationManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AllocationManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AllocationManagerOwnershipTransferred represents a OwnershipTransferred event raised by the AllocationManager contract.
type AllocationManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AllocationManager *AllocationManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AllocationManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AllocationManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AllocationManagerOwnershipTransferredIterator{contract: _AllocationManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AllocationManager *AllocationManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AllocationManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AllocationManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AllocationManagerOwnershipTransferred)
				if err := _AllocationManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_AllocationManager *AllocationManagerFilterer) ParseOwnershipTransferred(log types.Log) (*AllocationManagerOwnershipTransferred, error) {
	event := new(AllocationManagerOwnershipTransferred)
	if err := _AllocationManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AllocationManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the AllocationManager contract.
type AllocationManagerPausedIterator struct {
	Event *AllocationManagerPaused // Event containing the contract specifics and raw log

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
func (it *AllocationManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AllocationManagerPaused)
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
		it.Event = new(AllocationManagerPaused)
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
func (it *AllocationManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AllocationManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AllocationManagerPaused represents a Paused event raised by the AllocationManager contract.
type AllocationManagerPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_AllocationManager *AllocationManagerFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*AllocationManagerPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AllocationManager.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &AllocationManagerPausedIterator{contract: _AllocationManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_AllocationManager *AllocationManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *AllocationManagerPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AllocationManager.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AllocationManagerPaused)
				if err := _AllocationManager.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_AllocationManager *AllocationManagerFilterer) ParsePaused(log types.Log) (*AllocationManagerPaused, error) {
	event := new(AllocationManagerPaused)
	if err := _AllocationManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AllocationManagerStrategyAddedToOperatorSetIterator is returned from FilterStrategyAddedToOperatorSet and is used to iterate over the raw logs and unpacked data for StrategyAddedToOperatorSet events raised by the AllocationManager contract.
type AllocationManagerStrategyAddedToOperatorSetIterator struct {
	Event *AllocationManagerStrategyAddedToOperatorSet // Event containing the contract specifics and raw log

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
func (it *AllocationManagerStrategyAddedToOperatorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AllocationManagerStrategyAddedToOperatorSet)
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
		it.Event = new(AllocationManagerStrategyAddedToOperatorSet)
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
func (it *AllocationManagerStrategyAddedToOperatorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AllocationManagerStrategyAddedToOperatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AllocationManagerStrategyAddedToOperatorSet represents a StrategyAddedToOperatorSet event raised by the AllocationManager contract.
type AllocationManagerStrategyAddedToOperatorSet struct {
	OperatorSet OperatorSet
	Strategy    common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStrategyAddedToOperatorSet is a free log retrieval operation binding the contract event 0x7ab260fe0af193db5f4986770d831bda4ea46099dc817e8b6716dcae8af8e88b.
//
// Solidity: event StrategyAddedToOperatorSet((address,uint32) operatorSet, address strategy)
func (_AllocationManager *AllocationManagerFilterer) FilterStrategyAddedToOperatorSet(opts *bind.FilterOpts) (*AllocationManagerStrategyAddedToOperatorSetIterator, error) {

	logs, sub, err := _AllocationManager.contract.FilterLogs(opts, "StrategyAddedToOperatorSet")
	if err != nil {
		return nil, err
	}
	return &AllocationManagerStrategyAddedToOperatorSetIterator{contract: _AllocationManager.contract, event: "StrategyAddedToOperatorSet", logs: logs, sub: sub}, nil
}

// WatchStrategyAddedToOperatorSet is a free log subscription operation binding the contract event 0x7ab260fe0af193db5f4986770d831bda4ea46099dc817e8b6716dcae8af8e88b.
//
// Solidity: event StrategyAddedToOperatorSet((address,uint32) operatorSet, address strategy)
func (_AllocationManager *AllocationManagerFilterer) WatchStrategyAddedToOperatorSet(opts *bind.WatchOpts, sink chan<- *AllocationManagerStrategyAddedToOperatorSet) (event.Subscription, error) {

	logs, sub, err := _AllocationManager.contract.WatchLogs(opts, "StrategyAddedToOperatorSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AllocationManagerStrategyAddedToOperatorSet)
				if err := _AllocationManager.contract.UnpackLog(event, "StrategyAddedToOperatorSet", log); err != nil {
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

// ParseStrategyAddedToOperatorSet is a log parse operation binding the contract event 0x7ab260fe0af193db5f4986770d831bda4ea46099dc817e8b6716dcae8af8e88b.
//
// Solidity: event StrategyAddedToOperatorSet((address,uint32) operatorSet, address strategy)
func (_AllocationManager *AllocationManagerFilterer) ParseStrategyAddedToOperatorSet(log types.Log) (*AllocationManagerStrategyAddedToOperatorSet, error) {
	event := new(AllocationManagerStrategyAddedToOperatorSet)
	if err := _AllocationManager.contract.UnpackLog(event, "StrategyAddedToOperatorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AllocationManagerStrategyRemovedFromOperatorSetIterator is returned from FilterStrategyRemovedFromOperatorSet and is used to iterate over the raw logs and unpacked data for StrategyRemovedFromOperatorSet events raised by the AllocationManager contract.
type AllocationManagerStrategyRemovedFromOperatorSetIterator struct {
	Event *AllocationManagerStrategyRemovedFromOperatorSet // Event containing the contract specifics and raw log

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
func (it *AllocationManagerStrategyRemovedFromOperatorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AllocationManagerStrategyRemovedFromOperatorSet)
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
		it.Event = new(AllocationManagerStrategyRemovedFromOperatorSet)
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
func (it *AllocationManagerStrategyRemovedFromOperatorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AllocationManagerStrategyRemovedFromOperatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AllocationManagerStrategyRemovedFromOperatorSet represents a StrategyRemovedFromOperatorSet event raised by the AllocationManager contract.
type AllocationManagerStrategyRemovedFromOperatorSet struct {
	OperatorSet OperatorSet
	Strategy    common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStrategyRemovedFromOperatorSet is a free log retrieval operation binding the contract event 0x7b4b073d80dcac55a11177d8459ad9f664ceeb91f71f27167bb14f8152a7eeee.
//
// Solidity: event StrategyRemovedFromOperatorSet((address,uint32) operatorSet, address strategy)
func (_AllocationManager *AllocationManagerFilterer) FilterStrategyRemovedFromOperatorSet(opts *bind.FilterOpts) (*AllocationManagerStrategyRemovedFromOperatorSetIterator, error) {

	logs, sub, err := _AllocationManager.contract.FilterLogs(opts, "StrategyRemovedFromOperatorSet")
	if err != nil {
		return nil, err
	}
	return &AllocationManagerStrategyRemovedFromOperatorSetIterator{contract: _AllocationManager.contract, event: "StrategyRemovedFromOperatorSet", logs: logs, sub: sub}, nil
}

// WatchStrategyRemovedFromOperatorSet is a free log subscription operation binding the contract event 0x7b4b073d80dcac55a11177d8459ad9f664ceeb91f71f27167bb14f8152a7eeee.
//
// Solidity: event StrategyRemovedFromOperatorSet((address,uint32) operatorSet, address strategy)
func (_AllocationManager *AllocationManagerFilterer) WatchStrategyRemovedFromOperatorSet(opts *bind.WatchOpts, sink chan<- *AllocationManagerStrategyRemovedFromOperatorSet) (event.Subscription, error) {

	logs, sub, err := _AllocationManager.contract.WatchLogs(opts, "StrategyRemovedFromOperatorSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AllocationManagerStrategyRemovedFromOperatorSet)
				if err := _AllocationManager.contract.UnpackLog(event, "StrategyRemovedFromOperatorSet", log); err != nil {
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

// ParseStrategyRemovedFromOperatorSet is a log parse operation binding the contract event 0x7b4b073d80dcac55a11177d8459ad9f664ceeb91f71f27167bb14f8152a7eeee.
//
// Solidity: event StrategyRemovedFromOperatorSet((address,uint32) operatorSet, address strategy)
func (_AllocationManager *AllocationManagerFilterer) ParseStrategyRemovedFromOperatorSet(log types.Log) (*AllocationManagerStrategyRemovedFromOperatorSet, error) {
	event := new(AllocationManagerStrategyRemovedFromOperatorSet)
	if err := _AllocationManager.contract.UnpackLog(event, "StrategyRemovedFromOperatorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AllocationManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the AllocationManager contract.
type AllocationManagerUnpausedIterator struct {
	Event *AllocationManagerUnpaused // Event containing the contract specifics and raw log

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
func (it *AllocationManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AllocationManagerUnpaused)
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
		it.Event = new(AllocationManagerUnpaused)
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
func (it *AllocationManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AllocationManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AllocationManagerUnpaused represents a Unpaused event raised by the AllocationManager contract.
type AllocationManagerUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_AllocationManager *AllocationManagerFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*AllocationManagerUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AllocationManager.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &AllocationManagerUnpausedIterator{contract: _AllocationManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_AllocationManager *AllocationManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *AllocationManagerUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AllocationManager.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AllocationManagerUnpaused)
				if err := _AllocationManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_AllocationManager *AllocationManagerFilterer) ParseUnpaused(log types.Log) (*AllocationManagerUnpaused, error) {
	event := new(AllocationManagerUnpaused)
	if err := _AllocationManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
