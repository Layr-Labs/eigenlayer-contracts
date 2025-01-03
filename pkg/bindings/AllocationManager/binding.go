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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_delegation\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"_permissionController\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"},{\"name\":\"_DEALLOCATION_DELAY\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_ALLOCATION_CONFIGURATION_DELAY\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ALLOCATION_CONFIGURATION_DELAY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEALLOCATION_DELAY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addStrategiesToOperatorSet\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"clearDeallocationQueue\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"numToClear\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createOperatorSets\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"params\",\"type\":\"tuple[]\",\"internalType\":\"structIAllocationManagerTypes.CreateSetParams[]\",\"components\":[{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregisterFromOperatorSets\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIAllocationManagerTypes.DeregisterParams\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"encumberedMagnitude\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAVSRegistrar\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAVSRegistrar\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocatableMagnitude\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocatedSets\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocatedStrategies\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocation\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIAllocationManagerTypes.Allocation\",\"components\":[{\"name\":\"currentMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"pendingDiff\",\"type\":\"int128\",\"internalType\":\"int128\"},{\"name\":\"effectBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocationDelay\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocations\",\"inputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIAllocationManagerTypes.Allocation[]\",\"components\":[{\"name\":\"currentMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"pendingDiff\",\"type\":\"int128\",\"internalType\":\"int128\"},{\"name\":\"effectBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxMagnitude\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxMagnitudes\",\"inputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxMagnitudes\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxMagnitudesAtBlock\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMemberCount\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMembers\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMinimumSlashableStake\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"futureBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"slashableStake\",\"type\":\"uint256[][]\",\"internalType\":\"uint256[][]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorSetCount\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRegisteredSets\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStrategiesInOperatorSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStrategyAllocations\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIAllocationManagerTypes.Allocation[]\",\"components\":[{\"name\":\"currentMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"pendingDiff\",\"type\":\"int128\",\"internalType\":\"int128\"},{\"name\":\"effectBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isMemberOfOperatorSet\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperatorSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"modifyAllocations\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"params\",\"type\":\"tuple[]\",\"internalType\":\"structIAllocationManagerTypes.AllocateParams[]\",\"components\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"newMagnitudes\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"permissionController\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerForOperatorSets\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIAllocationManagerTypes.RegisterParams\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeStrategiesFromOperatorSet\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAVSRegistrar\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"registrar\",\"type\":\"address\",\"internalType\":\"contractIAVSRegistrar\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAllocationDelay\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delay\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slashOperator\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIAllocationManagerTypes.SlashingParams\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"wadsToSlash\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAVSMetadataURI\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AVSMetadataURIUpdated\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"metadataURI\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AVSRegistrarSet\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"registrar\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIAVSRegistrar\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllocationDelaySet\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"delay\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"effectBlock\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllocationUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"magnitude\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"effectBlock\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EncumberedMagnitudeUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"encumberedMagnitude\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MaxMagnitudeUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"maxMagnitude\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorAddedToOperatorSet\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRemovedFromOperatorSet\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSetCreated\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSlashed\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategies\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"contractIStrategy[]\"},{\"name\":\"wadSlashed\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"description\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyAddedToOperatorSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyRemovedFromOperatorSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyMemberOfSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CurrentlyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Empty\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputArrayLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientMagnitude\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCaller\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNewPausedStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperatorSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPermissions\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSnapshotOrdering\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidWadToSlash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MaxStrategiesExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ModificationAlreadyPending\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotMemberOfSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyPauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnpauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorNotSlashable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OutOfBounds\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SameMagnitude\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignatureExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategiesMustBeInAscendingOrder\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategyAlreadyInOperatorSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategyNotInOperatorSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UninitializedAllocationDelay\",\"inputs\":[]}]",
	Bin: "0x610120604052348015610010575f5ffd5b50604051615abb380380615abb83398101604081905261002f91610180565b82858383876001600160a01b03811661005b576040516339b190bb60e11b815260040160405180910390fd5b6001600160a01b0390811660805292831660a05263ffffffff91821660c0521660e052166101005261008b610095565b50505050506101e9565b5f54610100900460ff16156101005760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff9081161461014f575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b0381168114610165575f5ffd5b50565b805163ffffffff8116811461017b575f5ffd5b919050565b5f5f5f5f5f60a08688031215610194575f5ffd5b855161019f81610151565b60208701519095506101b081610151565b60408701519094506101c181610151565b92506101cf60608701610168565b91506101dd60808701610168565b90509295509295909350565b60805160a05160c05160e0516101005161584161027a5f395f81816103f9015261374a01525f81816105480152613d4901525f818161031e01528181612290015261298701525f81816106fa01528181610cec0152818161163701528181611cd501528181611d3f0152612c4301525f818161056f0152818161079201528181611de401526133c101526158415ff3fe608060405234801561000f575f5ffd5b5060043610610281575f3560e01c80636e875dba11610156578063a984eb3a116100ca578063c221d8ae11610084578063c221d8ae146106bc578063cd6dc687146106cf578063d3d96ff4146106e2578063df5cf723146106f5578063f2fde38b1461071c578063fabc1cbc1461072f575f5ffd5b8063a984eb3a1461060e578063adc2e3d914610641578063b2447af714610654578063b66bd98914610667578063b9fbaed11461067a578063ba1a84e5146106a9575f5ffd5b80638ce648541161011b5780638ce64854146105915780638da5cb5b146105b157806394d7d00c146105c2578063952899ee146105d5578063a9333ec8146105e8578063a9821821146105fb575f5ffd5b80636e875dba14610515578063715018a61461052857806379ae50cd146105305780637bc1ef6114610543578063886f11951461056a575f5ffd5b80634657e26a116101f8578063595c6a67116101b2578063595c6a67146104875780635ac86ab71461048f5780635c975abb146104b2578063670d3ba2146104c45780636cfb4481146104d75780636e3492b514610502575f5ffd5b80634657e26a146103f45780634a10ffe51461041b5780634b5046ef1461043b57806350feea201461044e578063547afb871461046157806356c483e614610474575f5ffd5b80632981eb77116102495780632981eb77146103195780632bab2c4a14610355578063304c10cd1461037557806336352057146103a057806340120dab146103b35780634177a87c146103d4575f5ffd5b806310e1b9b814610285578063136439dd146102ae57806315fe5028146102c3578063260dc758146102e3578063261f84e014610306575b5f5ffd5b61029861029336600461476e565b610742565b6040516102a591906147b5565b60405180910390f35b6102c16102bc3660046147e8565b61077d565b005b6102d66102d13660046147ff565b610852565b6040516102a5919061487d565b6102f66102f136600461488f565b610969565b60405190151581526020016102a5565b6102c16103143660046148e9565b6109a0565b6103407f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff90911681526020016102a5565b6103686103633660046149ce565b610c98565b6040516102a59190614a82565b6103886103833660046147ff565b610f85565b6040516001600160a01b0390911681526020016102a5565b6102c16103ae366004614ae5565b610fb4565b6103c66103c1366004614b37565b61178c565b6040516102a5929190614bc4565b6103e76103e236600461488f565b611907565b6040516102a59190614c21565b6103887f000000000000000000000000000000000000000000000000000000000000000081565b61042e610429366004614c33565b61192b565b6040516102a59190614c76565b6102c1610449366004614cc1565b6119d3565b6102c161045c366004614d41565b611a8d565b61042e61046f366004614d9f565b611c22565b6102c1610482366004614deb565b611cca565b6102c1611dcf565b6102f661049d366004614e1e565b606654600160ff9092169190911b9081161490565b6066545b6040519081526020016102a5565b6102f66104d2366004614e3e565b611e7e565b6104ea6104e5366004614b37565b611e8f565b6040516001600160401b0390911681526020016102a5565b6102c1610510366004614e7f565b611ffc565b6103e761052336600461488f565b6123cc565b6102c16123dd565b6102d661053e3660046147ff565b6123ee565b6103407f000000000000000000000000000000000000000000000000000000000000000081565b6103887f000000000000000000000000000000000000000000000000000000000000000081565b6105a461059f366004614eb0565b6124c8565b6040516102a59190614ef3565b6033546001600160a01b0316610388565b61042e6105d0366004614f05565b61258d565b6102c16105e3366004614f60565b612679565b6104ea6105f6366004614b37565b612b35565b6102c1610609366004615109565b612b64565b6104ea61061c366004614b37565b60a260209081525f92835260408084209091529082529020546001600160401b031681565b6102c161064f366004615187565b612bd4565b6104b661066236600461488f565b612f23565b6102c1610675366004614d41565b612f45565b61068d6106883660046147ff565b61309f565b60408051921515835263ffffffff9091166020830152016102a5565b6104b66106b73660046147ff565b613139565b6103e76106ca366004614e3e565b613159565b6102c16106dd3660046151c9565b61318a565b6102c16106f0366004614b37565b6132a7565b6103887f000000000000000000000000000000000000000000000000000000000000000081565b6102c161072a3660046147ff565b613346565b6102c161073d3660046147e8565b6133bf565b604080516060810182525f80825260208201819052918101829052906107718561076b866134d5565b85613538565b925050505b9392505050565b60405163237dfb4760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906346fbf68e90602401602060405180830381865afa1580156107df573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061080391906151f3565b61082057604051631d77d47760e21b815260040160405180910390fd5b60665481811681146108455760405163c61dca5d60e01b815260040160405180910390fd5b61084e826136a4565b5050565b6001600160a01b0381165f908152609d6020526040812060609190610876906136e1565b90505f816001600160401b0381111561089157610891614692565b6040519080825280602002602001820160405280156108d557816020015b604080518082019091525f80825260208201528152602001906001900390816108af5790505b5090505f5b82811015610961576001600160a01b0385165f908152609d6020526040902061093c9061090790836136ea565b604080518082019091525f80825260208201525060408051808201909152606082901c815263ffffffff909116602082015290565b82828151811061094e5761094e615212565b60209081029190910101526001016108da565b509392505050565b60208082015182516001600160a01b03165f90815260989092526040822061099a9163ffffffff908116906136f516565b92915050565b826109aa8161370c565b6109c75760405163932d94f760e01b815260040160405180910390fd5b5f5b82811015610c915760218484838181106109e5576109e5615212565b90506020028101906109f79190615226565b610a05906020810190615244565b90501115610a26576040516301a1443960e31b815260040160405180910390fd5b5f6040518060400160405280876001600160a01b03168152602001868685818110610a5357610a53615212565b9050602002810190610a659190615226565b610a73906020810190615289565b63ffffffff168152509050610abd816020015163ffffffff1660985f896001600160a01b03166001600160a01b031681526020019081526020015f206137b690919063ffffffff16565b610ada57604051631fb1705560e21b815260040160405180910390fd5b7f31629285ead2335ae0933f86ed2ae63321f7af77b4e6eaabc42c057880977e6c6040518060400160405280886001600160a01b03168152602001836020015163ffffffff16815250604051610b3091906152a2565b60405180910390a15f610b42826134d5565b90505f5b868685818110610b5857610b58615212565b9050602002810190610b6a9190615226565b610b78906020810190615244565b9050811015610c8657610bee878786818110610b9657610b96615212565b9050602002810190610ba89190615226565b610bb6906020810190615244565b83818110610bc657610bc6615212565b9050602002016020810190610bdb91906147ff565b5f848152609960205260409020906137c1565b507f7ab260fe0af193db5f4986770d831bda4ea46099dc817e8b6716dcae8af8e88b83888887818110610c2357610c23615212565b9050602002810190610c359190615226565b610c43906020810190615244565b84818110610c5357610c53615212565b9050602002016020810190610c6891906147ff565b604051610c769291906152b0565b60405180910390a1600101610b46565b5050506001016109c9565b5050505050565b606083516001600160401b03811115610cb357610cb3614692565b604051908082528060200260200182016040528015610ce657816020015b6060815260200190600190039081610cd15790505b5090505f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f0e0e67686866040518363ffffffff1660e01b8152600401610d389291906152d6565b5f60405180830381865afa158015610d52573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610d7991908101906152fa565b90505f5b8551811015610f7b575f868281518110610d9957610d99615212565b6020026020010151905085516001600160401b03811115610dbc57610dbc614692565b604051908082528060200260200182016040528015610de5578160200160208202803683370190505b50848381518110610df857610df8615212565b60209081029190910101525f5b8651811015610f71575f878281518110610e2157610e21615212565b6020908102919091018101516001600160a01b038086165f90815260a1845260408082209284168252919093528220909250610e5c906137d5565b9050806001600160401b03165f03610e75575050610f69565b5f610e81858d85610742565b90508863ffffffff16816040015163ffffffff1611158015610ea957505f8160200151600f0b125b15610ecb57610ebf815f015182602001516137e8565b6001600160401b031681525b80515f90610ee6906001600160401b039081169085166137fc565b9050610f2d81898981518110610efe57610efe615212565b60200260200101518781518110610f1757610f17615212565b602002602001015161381090919063ffffffff16565b898881518110610f3f57610f3f615212565b60200260200101518681518110610f5857610f58615212565b602002602001018181525050505050505b600101610e05565b5050600101610d7d565b5050949350505050565b6001600160a01b038082165f908152609760205260408120549091168015610fad5780610776565b5090919050565b606654600190600290811603610fdd5760405163840a48d560e01b815260040160405180910390fd5b82610fe78161370c565b6110045760405163932d94f760e01b815260040160405180910390fd5b5f6040518060400160405280866001600160a01b031681526020018560200160208101906110329190615289565b63ffffffff16905290505f61105361104d60208701876147ff565b83613824565b60208084015184516001600160a01b03165f90815260989092526040909120919250611089919063ffffffff908116906136f516565b6110a657604051631fb1705560e21b815260040160405180910390fd5b806110c45760405163ebbff49760e01b815260040160405180910390fd5b5f6110d26040870187615244565b90506001600160401b038111156110eb576110eb614692565b604051908082528060200260200182016040528015611114578160200160208202803683370190505b5090505f5b6111266040880188615244565b905081101561171d578015806111b957506111446040880188615244565b61114f60018461541a565b81811061115e5761115e615212565b905060200201602081019061117391906147ff565b6001600160a01b03166111896040890189615244565b8381811061119957611199615212565b90506020020160208101906111ae91906147ff565b6001600160a01b0316115b6111d657604051639f1c805360e01b815260040160405180910390fd5b6111e36060880188615244565b828181106111f3576111f3615212565b905060200201355f1080156112335750670de0b6b3a76400006112196060890189615244565b8381811061122957611229615212565b9050602002013511155b61125057604051631353603160e01b815260040160405180910390fd5b6112ac6112606040890189615244565b8381811061127057611270615212565b905060200201602081019061128591906147ff565b60995f611291886134d5565b81526020019081526020015f2061389990919063ffffffff16565b6112c9576040516331bc342760e11b815260040160405180910390fd5b5f8061131b6112db60208b018b6147ff565b6112e4886134d5565b6112f160408d018d615244565b8781811061130157611301615212565b905060200201602081019061131691906147ff565b613538565b805191935091506001600160401b03165f03611338575050611715565b5f61137361134960608c018c615244565b8681811061135957611359615212565b85516001600160401b0316926020909102013590506138ba565b835190915061138e6001600160401b038084169083166137fc565b8686815181106113a0576113a0615212565b60200260200101818152505081835f018181516113bd919061542d565b6001600160401b03169052508351829085906113da90839061542d565b6001600160401b03169052506020840180518391906113fa90839061542d565b6001600160401b031690525060208301515f600f9190910b1215611515575f61145d61142960608e018e615244565b8881811061143957611439615212565b90506020020135856020015161144e9061544c565b6001600160801b0316906138ba565b9050806001600160401b03168460200181815161147a9190615470565b600f0b9052507f1487af5418c47ee5ea45ef4a93398668120890774a9e13487e61e9dc3baf76dd6114ae60208e018e6147ff565b8a8e80604001906114bf9190615244565b8a8181106114cf576114cf615212565b90506020020160208101906114e491906147ff565b6114f5885f015189602001516137e8565b886040015160405161150b95949392919061549d565b60405180910390a1505b61156761152560208d018d6147ff565b61152e8a6134d5565b61153b60408f018f615244565b8981811061154b5761154b615212565b905060200201602081019061156091906147ff565b87876138d0565b7f1487af5418c47ee5ea45ef4a93398668120890774a9e13487e61e9dc3baf76dd61159560208d018d6147ff565b896115a360408f018f615244565b898181106115b3576115b3615212565b90506020020160208101906115c891906147ff565b86516040516115dc9493929190439061549d565b60405180910390a161162d6115f460208d018d6147ff565b61160160408e018e615244565b8881811061161157611611615212565b905060200201602081019061162691906147ff565b8651613b10565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001663ee74937f61166960208e018e6147ff565b61167660408f018f615244565b8981811061168657611686615212565b905060200201602081019061169b91906147ff565b875160405160e085901b6001600160e01b03191681526001600160a01b0393841660048201529290911660248301526001600160401b0380861660448401521660648201526084015f604051808303815f87803b1580156116fa575f5ffd5b505af115801561170c573d5f5f3e3d5ffd5b50505050505050505b600101611119565b507f80969ad29428d6797ee7aad084f9e4a42a82fc506dcd2ca3b6fb431f85ccebe561174c60208801886147ff565b8461175a60408a018a615244565b8561176860808d018d6154ee565b60405161177b9796959493929190615558565b60405180910390a150505050505050565b6001600160a01b0382165f908152609d6020526040812060609182916117b1906136e1565b90505f816001600160401b038111156117cc576117cc614692565b60405190808252806020026020018201604052801561181057816020015b604080518082019091525f80825260208201528152602001906001900390816117ea5790505b5090505f826001600160401b0381111561182c5761182c614692565b60405190808252806020026020018201604052801561187557816020015b604080516060810182525f80825260208083018290529282015282525f1990920191018161184a5790505b5090505f5b838110156118f8576001600160a01b0388165f908152609d602052604081206118a79061090790846136ea565b9050808483815181106118bc576118bc615212565b60200260200101819052506118d289828a610742565b8383815181106118e4576118e4615212565b60209081029190910101525060010161187a565b509093509150505b9250929050565b60605f61077660995f611919866134d5565b81526020019081526020015f20613b92565b60605f83516001600160401b0381111561194757611947614692565b604051908082528060200260200182016040528015611970578160200160208202803683370190505b5090505f5b8451811015610961576119a185828151811061199357611993615212565b602002602001015185612b35565b8282815181106119b3576119b3615212565b6001600160401b0390921660209283029190910190910152600101611975565b6066545f906001908116036119fb5760405163840a48d560e01b815260040160405180910390fd5b838214611a1b576040516343714afd60e01b815260040160405180910390fd5b5f5b84811015611a8457611a7c87878784818110611a3b57611a3b615212565b9050602002016020810190611a5091906147ff565b868685818110611a6257611a62615212565b9050602002016020810190611a7791906155ee565b613b9e565b600101611a1d565b50505050505050565b83611a978161370c565b611ab45760405163932d94f760e01b815260040160405180910390fd5b604080518082019091526001600160a01b038616815263ffffffff851660208201525f611ae0826134d5565b5f8181526099602052604090209091506021908590611afe906136e1565b611b08919061560f565b1115611b27576040516301a1443960e31b815260040160405180910390fd5b6020808301516001600160a01b0389165f90815260989092526040909120611b589163ffffffff908116906136f516565b611b7557604051631fb1705560e21b815260040160405180910390fd5b5f5b84811015611c1857611b94868683818110610bc657610bc6615212565b611bb15760405163585cfb2f60e01b815260040160405180910390fd5b7f7ab260fe0af193db5f4986770d831bda4ea46099dc817e8b6716dcae8af8e88b83878784818110611be557611be5615212565b9050602002016020810190611bfa91906147ff565b604051611c089291906152b0565b60405180910390a1600101611b77565b5050505050505050565b60605f82516001600160401b03811115611c3e57611c3e614692565b604051908082528060200260200182016040528015611c67578160200160208202803683370190505b5090505f5b835181101561096157611c9885858381518110611c8b57611c8b615212565b6020026020010151612b35565b828281518110611caa57611caa615212565b6001600160401b0390921660209283029190910190910152600101611c6c565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614611dc557611d038261370c565b611d20576040516348f5c3ed60e01b815260040160405180910390fd5b6040516336b87bd760e11b81526001600160a01b0383811660048301527f00000000000000000000000000000000000000000000000000000000000000001690636d70f7ae90602401602060405180830381865afa158015611d84573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611da891906151f3565b611dc55760405163ccea9e6f60e01b815260040160405180910390fd5b61084e8282613ca2565b60405163237dfb4760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906346fbf68e90602401602060405180830381865afa158015611e31573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611e5591906151f3565b611e7257604051631d77d47760e21b815260040160405180910390fd5b611e7c5f196136a4565b565b5f61077683609a5f611291866134d5565b6001600160a01b038281165f81815260a2602090815260408083209486168084529482528083205493835260a38252808320948352939052918220546001600160401b0390911690600f81810b600160801b909204900b03825b81811015611fb9576001600160a01b038087165f90815260a3602090815260408083209389168352929052908120611f219083613e43565b6001600160a01b038881165f90815260a0602090815260408083208584528252808320938b16835292815290829020825160608101845290546001600160401b0381168252600160401b8104600f0b92820192909252600160c01b90910463ffffffff16918101829052919250431015611f9c575050611fb9565b611faa8582602001516137e8565b94505050806001019050611ee9565b506001600160a01b038086165f90815260a1602090815260408083209388168352929052208290611fe9906137d5565b611ff3919061542d565b95945050505050565b6066546002906004908116036120255760405163840a48d560e01b815260040160405180910390fd5b61203a61203560208401846147ff565b61370c565b80612053575061205361203560408401602085016147ff565b612070576040516348f5c3ed60e01b815260040160405180910390fd5b5f5b61207f6040840184615244565b9050811015612341575f60405180604001604052808560200160208101906120a791906147ff565b6001600160a01b031681526020016120c26040870187615244565b858181106120d2576120d2615212565b90506020020160208101906120e79190615289565b63ffffffff168152509050612134816020015163ffffffff1660985f87602001602081019061211691906147ff565b6001600160a01b0316815260208101919091526040015f20906136f5565b61215157604051631fb1705560e21b815260040160405180910390fd5b609e5f61216160208701876147ff565b6001600160a01b03166001600160a01b031681526020019081526020015f205f61218a836134d5565b815260208101919091526040015f205460ff166121ba576040516325131d4f60e01b815260040160405180910390fd5b6121f46121c6826134d5565b609c5f6121d660208901896147ff565b6001600160a01b0316815260208101919091526040015f2090613eb2565b5061222c61220560208601866147ff565b609a5f612211856134d5565b81526020019081526020015f20613ebd90919063ffffffff16565b5061223a60208501856147ff565b6001600160a01b03167fad34c3070be1dffbcaa499d000ba2b8d9848aefcac3059df245dd95c4ece14fe8260405161227291906152a2565b60405180910390a2604080518082019091525f8152602081016122b57f000000000000000000000000000000000000000000000000000000000000000043615622565b63ffffffff169052609e5f6122cd60208801886147ff565b6001600160a01b03166001600160a01b031681526020019081526020015f205f6122f6846134d5565b81526020808201929092526040015f2082518154939092015163ffffffff166101000264ffffffff00199215159290921664ffffffffff199093169290921717905550600101612072565b5061235561038360408401602085016147ff565b6001600160a01b0316639d8e0c2361237060208501856147ff565b61237d6040860186615244565b6040518463ffffffff1660e01b815260040161239b93929190615677565b5f604051808303815f87803b1580156123b2575f5ffd5b505af19250505080156123c3575060015b1561084e575050565b606061099a609a5f611919856134d5565b6123e5613ed1565b611e7c5f613f2b565b6001600160a01b0381165f908152609c6020526040812060609190612412906136e1565b90505f816001600160401b0381111561242d5761242d614692565b60405190808252806020026020018201604052801561247157816020015b604080518082019091525f808252602082015281526020019060019003908161244b5790505b5090505f5b82811015610961576001600160a01b0385165f908152609c602052604090206124a39061090790836136ea565b8282815181106124b5576124b5615212565b6020908102919091010152600101612476565b60605f84516001600160401b038111156124e4576124e4614692565b60405190808252806020026020018201604052801561252d57816020015b604080516060810182525f80825260208083018290529282015282525f199092019101816125025790505b5090505f5b85518110156125845761255f86828151811061255057612550615212565b60200260200101518686610742565b82828151811061257157612571615212565b6020908102919091010152600101612532565b50949350505050565b60605f83516001600160401b038111156125a9576125a9614692565b6040519080825280602002602001820160405280156125d2578160200160208202803683370190505b5090505f5b8451811015612584576001600160a01b0386165f90815260a16020526040812086516126479287929189908690811061261257612612615212565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020015f20613f7c90919063ffffffff16565b82828151811061265957612659615212565b6001600160401b03909216602092830291909101909101526001016125d7565b6066545f906001908116036126a15760405163840a48d560e01b815260040160405180910390fd5b6126aa8361370c565b6126c7576040516348f5c3ed60e01b815260040160405180910390fd5b5f5f5f6126d38661309f565b91509150816126f55760405163fa55fc8160e01b815260040160405180910390fd5b91505f90505b8351811015610c915783818151811061271657612716615212565b6020026020010151604001515184828151811061273557612735615212565b6020026020010151602001515114612760576040516343714afd60e01b815260040160405180910390fd5b5f84828151811061277357612773615212565b602090810291909101810151518082015181516001600160a01b03165f908152609890935260409092209092506127b39163ffffffff908116906136f516565b6127d057604051631fb1705560e21b815260040160405180910390fd5b5f6127db8783613824565b90505f5b8684815181106127f1576127f1615212565b60200260200101516020015151811015612b2a575f87858151811061281857612818615212565b602002602001015160200151828151811061283557612835615212565b6020026020010151905061284c898261ffff613b9e565b5f5f61285b8b61076b886134d5565b915091508060200151600f0b5f1461288657604051630d8fcbe360e41b815260040160405180910390fd5b5f61289387858489613f90565b90506128d8825f01518c8a815181106128ae576128ae615212565b60200260200101516040015187815181106128cb576128cb615212565b6020026020010151613fc6565b600f0b602083018190525f0361290157604051634606179360e11b815260040160405180910390fd5b5f8260200151600f0b1215612a3a5780156129bc57612982612922886134d5565b6001600160a01b03808f165f90815260a360209081526040808320938a16835292905220908154600160801b90819004600f0b5f818152600180860160205260409091209390935583546001600160801b03908116939091011602179055565b6129ac7f000000000000000000000000000000000000000000000000000000000000000043615622565b63ffffffff166040830152612aa7565b6129ce836020015183602001516137e8565b6001600160401b031660208401528a518b90899081106129f0576129f0615212565b6020026020010151604001518581518110612a0d57612a0d615212565b6020908102919091018101516001600160401b031683525f9083015263ffffffff43166040830152612aa7565b5f8260200151600f0b1315612aa757612a5b836020015183602001516137e8565b6001600160401b039081166020850181905284519091161015612a9157604051636c9be0bf60e01b815260040160405180910390fd5b612a9b8943615622565b63ffffffff1660408301525b612abc8c612ab4896134d5565b8686866138d0565b7f1487af5418c47ee5ea45ef4a93398668120890774a9e13487e61e9dc3baf76dd8c612aea6109078a6134d5565b86612afc865f015187602001516137e8565b8660400151604051612b1295949392919061549d565b60405180910390a15050600190920191506127df9050565b5050506001016126fb565b6001600160a01b038083165f90815260a1602090815260408083209385168352929052908120610776906137d5565b82612b6e8161370c565b612b8b5760405163932d94f760e01b815260040160405180910390fd5b836001600160a01b03167fa89c1dc243d8908a96dd84944bcc97d6bc6ac00dd78e20621576be6a3c9437138484604051612bc692919061569b565b60405180910390a250505050565b606654600290600490811603612bfd5760405163840a48d560e01b815260040160405180910390fd5b82612c078161370c565b612c245760405163932d94f760e01b815260040160405180910390fd5b6040516336b87bd760e11b81526001600160a01b0385811660048301527f00000000000000000000000000000000000000000000000000000000000000001690636d70f7ae90602401602060405180830381865afa158015612c88573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612cac91906151f3565b612cc95760405163ccea9e6f60e01b815260040160405180910390fd5b5f5b612cd86020850185615244565b9050811015612ea057604080518082019091525f9080612cfb60208801886147ff565b6001600160a01b03168152602001868060200190612d199190615244565b85818110612d2957612d29615212565b9050602002016020810190612d3e9190615289565b63ffffffff90811690915260208083015183516001600160a01b03165f90815260989092526040909120929350612d7a9291908116906136f516565b612d9757604051631fb1705560e21b815260040160405180910390fd5b612da18682613824565b15612dbf57604051636c6c6e2760e11b815260040160405180910390fd5b612de8612dcb826134d5565b6001600160a01b0388165f908152609c60205260409020906137b6565b50612e1486609a5f612df9856134d5565b81526020019081526020015f206137c190919063ffffffff16565b50856001600160a01b03167f43232edf9071753d2321e5fa7e018363ee248e5f2142e6c08edd3265bfb4895e82604051612e4e91906152a2565b60405180910390a26001600160a01b0386165f908152609e60205260408120600191612e79846134d5565b815260208101919091526040015f20805460ff191691151591909117905550600101612ccb565b50612eb161038360208501856147ff565b6001600160a01b031663adcf73f785612ecd6020870187615244565b612eda60408901896154ee565b6040518663ffffffff1660e01b8152600401612efa9594939291906156ae565b5f604051808303815f87803b158015612f11575f5ffd5b505af1158015611c18573d5f5f3e3d5ffd5b5f61099a609a5f612f33856134d5565b81526020019081526020015f206136e1565b83612f4f8161370c565b612f6c5760405163932d94f760e01b815260040160405180910390fd5b6040805180820182526001600160a01b03871680825263ffffffff80881660208085018290525f93845260989052939091209192612fab92916136f516565b612fc857604051631fb1705560e21b815260040160405180910390fd5b5f612fd2826134d5565b90505f5b84811015611c185761301b868683818110612ff357612ff3615212565b905060200201602081019061300891906147ff565b5f84815260996020526040902090613ebd565b613038576040516331bc342760e11b815260040160405180910390fd5b7f7b4b073d80dcac55a11177d8459ad9f664ceeb91f71f27167bb14f8152a7eeee8387878481811061306c5761306c615212565b905060200201602081019061308191906147ff565b60405161308f9291906152b0565b60405180910390a1600101612fd6565b6001600160a01b0381165f908152609b602090815260408083208151608081018352905463ffffffff80821680845260ff600160201b8404161515958401869052650100000000008304821694840194909452600160481b9091041660608201819052849391929190158015906131205750826060015163ffffffff164310155b1561312f575050604081015160015b9590945092505050565b6001600160a01b0381165f90815260986020526040812061099a906136e1565b6001600160a01b0382165f908152609f60205260408120606091906131829082611919866134d5565b949350505050565b5f54610100900460ff16158080156131a857505f54600160ff909116105b806131c15750303b1580156131c157505f5460ff166001145b6132295760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b5f805460ff19166001179055801561324a575f805461ff0019166101001790555b613253826136a4565b61325c83613f2b565b80156132a2575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498906020015b60405180910390a15b505050565b816132b18161370c565b6132ce5760405163932d94f760e01b815260040160405180910390fd5b6001600160a01b038381165f90815260976020526040902080546001600160a01b0319169184169190911790557f2ae945c40c44dc0ec263f95609c3fdc6952e0aefa22d6374e44f2c997acedf858361332681610f85565b604080516001600160a01b03938416815292909116602083015201613299565b61334e613ed1565b6001600160a01b0381166133b35760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401613220565b6133bc81613f2b565b50565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561341b573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061343f91906156f1565b6001600160a01b0316336001600160a01b0316146134705760405163794821ff60e01b815260040160405180910390fd5b606654801982198116146134975760405163c61dca5d60e01b815260040160405180910390fd5b606682905560405182815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c9060200160405180910390a25050565b5f815f0151826020015163ffffffff1660405160200161352092919060609290921b6bffffffffffffffffffffffff1916825260a01b6001600160a01b031916601482015260200190565b60405160208183030381529060405261099a9061570c565b6040805180820182525f80825260208083018290528351606081018552828152808201839052808501839052845180860186526001600160a01b03898116855260a1845286852090881685529092529382209293928190613598906137d5565b6001600160401b0390811682526001600160a01b038981165f81815260a260209081526040808320948c168084529482528083205486169682019690965291815260a082528481208b8252825284812092815291815290839020835160608101855290549283168152600160401b8304600f0b91810191909152600160c01b90910463ffffffff1691810182905291925043101561363a57909250905061369c565b61364b815f015182602001516137e8565b6001600160401b0316815260208101515f600f9190910b12156136895761367a826020015182602001516137e8565b6001600160401b031660208301525b5f60408201819052602082015290925090505b935093915050565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a250565b5f61099a825490565b5f6107768383613fdd565b5f8181526001830160205260408120541515610776565b604051631beb2b9760e31b81526001600160a01b0382811660048301523360248301523060448301525f80356001600160e01b0319166064840152917f00000000000000000000000000000000000000000000000000000000000000009091169063df595cb8906084016020604051808303815f875af1158015613792573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061099a91906151f3565b5f6107768383614003565b5f610776836001600160a01b038416614003565b5f61099a82670de0b6b3a764000061404f565b5f610776826001600160401b038516615470565b5f61077683670de0b6b3a764000084614093565b5f6107768383670de0b6b3a7640000614093565b6001600160a01b0382165f908152609e60205260408120819081613847856134d5565b815260208082019290925260409081015f2081518083019092525460ff8116151580835261010090910463ffffffff169282019290925291508061318257506020015163ffffffff1643109392505050565b6001600160a01b0381165f9081526001830160205260408120541515610776565b5f6107768383670de0b6b3a76400006001614178565b6020808301516001600160a01b038088165f90815260a284526040808220928816825291909352909120546001600160401b0390811691161461399657602082810180516001600160a01b038881165f81815260a286526040808220938a1680835293875290819020805467ffffffffffffffff19166001600160401b0395861617905593518451918252948101919091529216908201527facf9095feb3a370c9cf692421c69ef320d4db5c66e6a7d29c7694eb02364fc559060600160405180910390a15b6001600160a01b038086165f90815260a060209081526040808320888452825280832093871683529281529082902083518154928501519385015163ffffffff16600160c01b0263ffffffff60c01b196001600160801b038616600160401b026001600160c01b03199095166001600160401b03909316929092179390931716919091179055600f0b15613a78576001600160a01b0385165f908152609f602090815260408083208784529091529020613a5090846137c1565b506001600160a01b0385165f908152609d60205260409020613a7290856137b6565b50610c91565b80516001600160401b03165f03610c91576001600160a01b0385165f908152609f602090815260408083208784529091529020613ab59084613ebd565b506001600160a01b0385165f908152609f602090815260408083208784529091529020613ae1906136e1565b5f03610c91576001600160a01b0385165f908152609d60205260409020613b089085613eb2565b505050505050565b6001600160a01b038084165f90815260a160209081526040808320938616835292905220613b3f9043836141d1565b604080516001600160a01b038086168252841660208201526001600160401b038316918101919091527f1c6458079a41077d003c11faf9bf097e693bd67979e4e6500bac7b29db779b5c90606001613299565b60605f610776836141e5565b6001600160a01b038381165f90815260a360209081526040808320938616835292905290812054600f81810b600160801b909204900b035b5f81118015613be857508261ffff1682105b15610c91576001600160a01b038086165f90815260a3602090815260408083209388168352929052908120613c1c9061423e565b90505f5f613c2b888489613538565b91509150806040015163ffffffff16431015613c4957505050610c91565b613c5688848985856138d0565b6001600160a01b038089165f90815260a360209081526040808320938b16835292905220613c8390614290565b50613c8d8561572f565b9450613c9884615747565b9350505050613bd6565b6001600160a01b0382165f908152609b60209081526040918290208251608081018452905463ffffffff808216835260ff600160201b830416151593830193909352650100000000008104831693820193909352600160481b909204166060820181905215801590613d1e5750806060015163ffffffff164310155b15613d3857604081015163ffffffff168152600160208201525b63ffffffff82166040820152613d6e7f000000000000000000000000000000000000000000000000000000000000000043615622565b63ffffffff90811660608381019182526001600160a01b0386165f818152609b602090815260409182902087518154838a0151858b01519851928a1664ffffffffff1990921691909117600160201b91151591909102176cffffffffffffffff0000000000191665010000000000978916979097026cffffffff000000000000000000191696909617600160481b968816968702179055815192835294871694820194909452928301919091527f4e85751d6331506c6c62335f207eb31f12a61e570f34f5c17640308785c6d4db9101613299565b5f5f613e65613e518461430d565b8554613e609190600f0b61575c565b61437a565b8454909150600160801b9004600f90810b9082900b12613e9857604051632d0483c560e21b815260040160405180910390fd5b600f0b5f9081526001939093016020525050604090205490565b5f61077683836143e3565b5f610776836001600160a01b0384166143e3565b6033546001600160a01b03163314611e7c5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401613220565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b5f6107768383670de0b6b3a76400006144c6565b5f613fa18460995f611291896134d5565b8015613faa5750815b8015611ff357505090516001600160401b031615159392505050565b5f6107766001600160401b03808516908416615783565b5f825f018281548110613ff257613ff2615212565b905f5260205f200154905092915050565b5f81815260018301602052604081205461404857508154600181810184555f84815260208082209093018490558454848252828601909352604090209190915561099a565b505f61099a565b81545f90801561408b576140758461406860018461541a565b5f91825260209091200190565b54600160201b90046001600160e01b0316613182565b509092915050565b5f80805f19858709858702925082811083820303915050805f036140ca578382816140c0576140c06157b0565b0492505050610776565b8084116141115760405162461bcd60e51b81526020600482015260156024820152744d6174683a206d756c446976206f766572666c6f7760581b6044820152606401613220565b5f8486880960026001871981018816978890046003810283188082028403028082028403028082028403028082028403028082028403029081029092039091025f889003889004909101858311909403939093029303949094049190911702949350505050565b5f5f614185868686614093565b9050600183600281111561419b5761419b6157c4565b1480156141b757505f84806141b2576141b26157b0565b868809115b15611ff3576141c760018261560f565b9695505050505050565b6132a283836001600160401b03841661450e565b6060815f0180548060200260200160405190810160405280929190818152602001828054801561423257602002820191905f5260205f20905b81548152602001906001019080831161421e575b50505050509050919050565b5f6142588254600f81810b600160801b909204900b131590565b1561427657604051631ed9509560e11b815260040160405180910390fd5b508054600f0b5f9081526001909101602052604090205490565b5f6142aa8254600f81810b600160801b909204900b131590565b156142c857604051631ed9509560e11b815260040160405180910390fd5b508054600f0b5f818152600180840160205260408220805492905583546fffffffffffffffffffffffffffffffff191692016001600160801b03169190911790915590565b5f6001600160ff1b038211156143765760405162461bcd60e51b815260206004820152602860248201527f53616665436173743a2076616c756520646f65736e27742066697420696e2061604482015267371034b73a191a9b60c11b6064820152608401613220565b5090565b80600f81900b81146143de5760405162461bcd60e51b815260206004820152602760248201527f53616665436173743a2076616c756520646f65736e27742066697420696e20316044820152663238206269747360c81b6064820152608401613220565b919050565b5f81815260018301602052604081205480156144bd575f61440560018361541a565b85549091505f906144189060019061541a565b9050818114614477575f865f01828154811061443657614436615212565b905f5260205f200154905080875f01848154811061445657614456615212565b5f918252602080832090910192909255918252600188019052604090208390555b8554869080614488576144886157d8565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f90556001935050505061099a565b5f91505061099a565b82545f90816144d786868385614611565b90508015614504576144ee8661406860018461541a565b54600160201b90046001600160e01b0316610771565b5091949350505050565b825480156145c4575f6145268561406860018561541a565b60408051808201909152905463ffffffff808216808452600160201b9092046001600160e01b0316602084015291925090851610156145785760405163151b8e3f60e11b815260040160405180910390fd5b805163ffffffff8086169116036145c257826145998661406860018661541a565b80546001600160e01b0392909216600160201b0263ffffffff9092169190911790555050505050565b505b506040805180820190915263ffffffff92831681526001600160e01b03918216602080830191825285546001810187555f968752952091519051909216600160201b029190921617910155565b5f5b81831015610961575f6146268484614664565b5f8781526020902090915063ffffffff86169082015463ffffffff1611156146505780925061465e565b61465b81600161560f565b93505b50614613565b5f61467260028484186157ec565b6107769084841661560f565b6001600160a01b03811681146133bc575f5ffd5b634e487b7160e01b5f52604160045260245ffd5b604051606081016001600160401b03811182821017156146c8576146c8614692565b60405290565b604051601f8201601f191681016001600160401b03811182821017156146f6576146f6614692565b604052919050565b803563ffffffff811681146143de575f5ffd5b5f60408284031215614721575f5ffd5b604080519081016001600160401b038111828210171561474357614743614692565b60405290508082356147548161467e565b8152614762602084016146fe565b60208201525092915050565b5f5f5f60808486031215614780575f5ffd5b833561478b8161467e565b925061479a8560208601614711565b915060608401356147aa8161467e565b809150509250925092565b81516001600160401b03168152602080830151600f0b9082015260408083015163ffffffff16908201526060810161099a565b5f602082840312156147f8575f5ffd5b5035919050565b5f6020828403121561480f575f5ffd5b81356107768161467e565b80516001600160a01b0316825260209081015163ffffffff16910152565b5f8151808452602084019350602083015f5b828110156148735761485d86835161481a565b604095909501946020919091019060010161484a565b5093949350505050565b602081525f6107766020830184614838565b5f6040828403121561489f575f5ffd5b6107768383614711565b5f5f83601f8401126148b9575f5ffd5b5081356001600160401b038111156148cf575f5ffd5b6020830191508360208260051b8501011115611900575f5ffd5b5f5f5f604084860312156148fb575f5ffd5b83356149068161467e565b925060208401356001600160401b03811115614920575f5ffd5b61492c868287016148a9565b9497909650939450505050565b5f6001600160401b0382111561495157614951614692565b5060051b60200190565b5f82601f83011261496a575f5ffd5b813561497d61497882614939565b6146ce565b8082825260208201915060208360051b86010192508583111561499e575f5ffd5b602085015b838110156149c45780356149b68161467e565b8352602092830192016149a3565b5095945050505050565b5f5f5f5f60a085870312156149e1575f5ffd5b6149eb8686614711565b935060408501356001600160401b03811115614a05575f5ffd5b614a118782880161495b565b93505060608501356001600160401b03811115614a2c575f5ffd5b614a388782880161495b565b925050614a47608086016146fe565b905092959194509250565b5f8151808452602084019350602083015f5b82811015614873578151865260209586019590910190600101614a64565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b82811015614ad957603f19878603018452614ac4858351614a52565b94506020938401939190910190600101614aa8565b50929695505050505050565b5f5f60408385031215614af6575f5ffd5b8235614b018161467e565b915060208301356001600160401b03811115614b1b575f5ffd5b830160a08186031215614b2c575f5ffd5b809150509250929050565b5f5f60408385031215614b48575f5ffd5b8235614b538161467e565b91506020830135614b2c8161467e565b5f8151808452602084019350602083015f5b8281101561487357614bae86835180516001600160401b03168252602080820151600f0b9083015260409081015163ffffffff16910152565b6060959095019460209190910190600101614b75565b604081525f614bd66040830185614838565b8281036020840152611ff38185614b63565b5f8151808452602084019350602083015f5b828110156148735781516001600160a01b0316865260209586019590910190600101614bfa565b602081525f6107766020830184614be8565b5f5f60408385031215614c44575f5ffd5b82356001600160401b03811115614c59575f5ffd5b614c658582860161495b565b9250506020830135614b2c8161467e565b602080825282518282018190525f918401906040840190835b81811015614cb65783516001600160401b0316835260209384019390920191600101614c8f565b509095945050505050565b5f5f5f5f5f60608688031215614cd5575f5ffd5b8535614ce08161467e565b945060208601356001600160401b03811115614cfa575f5ffd5b614d06888289016148a9565b90955093505060408601356001600160401b03811115614d24575f5ffd5b614d30888289016148a9565b969995985093965092949392505050565b5f5f5f5f60608587031215614d54575f5ffd5b8435614d5f8161467e565b9350614d6d602086016146fe565b925060408501356001600160401b03811115614d87575f5ffd5b614d93878288016148a9565b95989497509550505050565b5f5f60408385031215614db0575f5ffd5b8235614dbb8161467e565b915060208301356001600160401b03811115614dd5575f5ffd5b614de18582860161495b565b9150509250929050565b5f5f60408385031215614dfc575f5ffd5b8235614e078161467e565b9150614e15602084016146fe565b90509250929050565b5f60208284031215614e2e575f5ffd5b813560ff81168114610776575f5ffd5b5f5f60608385031215614e4f575f5ffd5b8235614e5a8161467e565b9150614e158460208501614711565b5f60608284031215614e79575f5ffd5b50919050565b5f60208284031215614e8f575f5ffd5b81356001600160401b03811115614ea4575f5ffd5b61318284828501614e69565b5f5f5f60808486031215614ec2575f5ffd5b83356001600160401b03811115614ed7575f5ffd5b614ee38682870161495b565b93505061479a8560208601614711565b602081525f6107766020830184614b63565b5f5f5f60608486031215614f17575f5ffd5b8335614f228161467e565b925060208401356001600160401b03811115614f3c575f5ffd5b614f488682870161495b565b925050614f57604085016146fe565b90509250925092565b5f5f60408385031215614f71575f5ffd5b8235614f7c8161467e565b915060208301356001600160401b03811115614f96575f5ffd5b8301601f81018513614fa6575f5ffd5b8035614fb461497882614939565b8082825260208201915060208360051b850101925087831115614fd5575f5ffd5b602084015b838110156150fa5780356001600160401b03811115614ff7575f5ffd5b85016080818b03601f1901121561500c575f5ffd5b6150146146a6565b6150218b60208401614711565b815260608201356001600160401b0381111561503b575f5ffd5b61504a8c60208386010161495b565b60208301525060808201356001600160401b03811115615068575f5ffd5b6020818401019250508a601f83011261507f575f5ffd5b813561508d61497882614939565b8082825260208201915060208360051b86010192508d8311156150ae575f5ffd5b6020850194505b828510156150e45784356001600160401b03811681146150d3575f5ffd5b8252602094850194909101906150b5565b6040840152505084525060209283019201614fda565b50809450505050509250929050565b5f5f5f6040848603121561511b575f5ffd5b83356151268161467e565b925060208401356001600160401b03811115615140575f5ffd5b8401601f81018613615150575f5ffd5b80356001600160401b03811115615165575f5ffd5b866020828401011115615176575f5ffd5b939660209190910195509293505050565b5f5f60408385031215615198575f5ffd5b82356151a38161467e565b915060208301356001600160401b038111156151bd575f5ffd5b614de185828601614e69565b5f5f604083850312156151da575f5ffd5b82356151e58161467e565b946020939093013593505050565b5f60208284031215615203575f5ffd5b81518015158114610776575f5ffd5b634e487b7160e01b5f52603260045260245ffd5b5f8235603e1983360301811261523a575f5ffd5b9190910192915050565b5f5f8335601e19843603018112615259575f5ffd5b8301803591506001600160401b03821115615272575f5ffd5b6020019150600581901b3603821315611900575f5ffd5b5f60208284031215615299575f5ffd5b610776826146fe565b6040810161099a828461481a565b606081016152be828561481a565b6001600160a01b039290921660409190910152919050565b604081525f6152e86040830185614be8565b8281036020840152611ff38185614be8565b5f6020828403121561530a575f5ffd5b81516001600160401b0381111561531f575f5ffd5b8201601f8101841361532f575f5ffd5b805161533d61497882614939565b8082825260208201915060208360051b85010192508683111561535e575f5ffd5b602084015b838110156153fb5780516001600160401b03811115615380575f5ffd5b8501603f81018913615390575f5ffd5b60208101516153a161497882614939565b808282526020820191506020808460051b8601010192508b8311156153c4575f5ffd5b6040840193505b828410156153e65783518252602093840193909101906153cb565b86525050602093840193919091019050615363565b509695505050505050565b634e487b7160e01b5f52601160045260245ffd5b8181038181111561099a5761099a615406565b6001600160401b03828116828216039081111561099a5761099a615406565b5f81600f0b60016001607f1b0319810361546857615468615406565b5f0392915050565b600f81810b9083900b0160016001607f1b03811360016001607f1b03198212171561099a5761099a615406565b6001600160a01b038616815260c081016154ba602083018761481a565b6001600160a01b039490941660608201526001600160401b0392909216608083015263ffffffff1660a09091015292915050565b5f5f8335601e19843603018112615503575f5ffd5b8301803591506001600160401b0382111561551c575f5ffd5b602001915036819003821315611900575f5ffd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b6001600160a01b03881681525f60c08201615576602084018a61481a565b60c060608401528690528660e083015f5b888110156155b757823561559a8161467e565b6001600160a01b0316825260209283019290910190600101615587565b5083810360808501526155ca8188614a52565b91505082810360a08401526155e0818587615530565b9a9950505050505050505050565b5f602082840312156155fe575f5ffd5b813561ffff81168114610776575f5ffd5b8082018082111561099a5761099a615406565b63ffffffff818116838216019081111561099a5761099a615406565b8183526020830192505f815f5b848110156148735763ffffffff615661836146fe565b168652602095860195919091019060010161564b565b6001600160a01b03841681526040602082018190525f90611ff3908301848661563e565b602081525f613182602083018486615530565b6001600160a01b03861681526060602082018190525f906156d2908301868861563e565b82810360408401526156e5818587615530565b98975050505050505050565b5f60208284031215615701575f5ffd5b81516107768161467e565b80516020808301519190811015614e79575f1960209190910360031b1b16919050565b5f6001820161574057615740615406565b5060010190565b5f8161575557615755615406565b505f190190565b8082018281125f83128015821682158216171561577b5761577b615406565b505092915050565b600f82810b9082900b0360016001607f1b0319811260016001607f1b038213171561099a5761099a615406565b634e487b7160e01b5f52601260045260245ffd5b634e487b7160e01b5f52602160045260245ffd5b634e487b7160e01b5f52603160045260245ffd5b5f8261580657634e487b7160e01b5f52601260045260245ffd5b50049056fea26469706673582212202c0f483a478d14f76ed057d28cec9d6ef94a45015cd283d99238b804b9741b8264736f6c634300081b0033",
}

// AllocationManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use AllocationManagerMetaData.ABI instead.
var AllocationManagerABI = AllocationManagerMetaData.ABI

// AllocationManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AllocationManagerMetaData.Bin instead.
var AllocationManagerBin = AllocationManagerMetaData.Bin

// DeployAllocationManager deploys a new Ethereum contract, binding an instance of AllocationManager to it.
func DeployAllocationManager(auth *bind.TransactOpts, backend bind.ContractBackend, _delegation common.Address, _pauserRegistry common.Address, _permissionController common.Address, _DEALLOCATION_DELAY uint32, _ALLOCATION_CONFIGURATION_DELAY uint32) (common.Address, *types.Transaction, *AllocationManager, error) {
	parsed, err := AllocationManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AllocationManagerBin), backend, _delegation, _pauserRegistry, _permissionController, _DEALLOCATION_DELAY, _ALLOCATION_CONFIGURATION_DELAY)
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

// EncumberedMagnitude is a free data retrieval call binding the contract method 0xa984eb3a.
//
// Solidity: function encumberedMagnitude(address operator, address strategy) view returns(uint64)
func (_AllocationManager *AllocationManagerCaller) EncumberedMagnitude(opts *bind.CallOpts, operator common.Address, strategy common.Address) (uint64, error) {
	var out []interface{}
	err := _AllocationManager.contract.Call(opts, &out, "encumberedMagnitude", operator, strategy)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// EncumberedMagnitude is a free data retrieval call binding the contract method 0xa984eb3a.
//
// Solidity: function encumberedMagnitude(address operator, address strategy) view returns(uint64)
func (_AllocationManager *AllocationManagerSession) EncumberedMagnitude(operator common.Address, strategy common.Address) (uint64, error) {
	return _AllocationManager.Contract.EncumberedMagnitude(&_AllocationManager.CallOpts, operator, strategy)
}

// EncumberedMagnitude is a free data retrieval call binding the contract method 0xa984eb3a.
//
// Solidity: function encumberedMagnitude(address operator, address strategy) view returns(uint64)
func (_AllocationManager *AllocationManagerCallerSession) EncumberedMagnitude(operator common.Address, strategy common.Address) (uint64, error) {
	return _AllocationManager.Contract.EncumberedMagnitude(&_AllocationManager.CallOpts, operator, strategy)
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
