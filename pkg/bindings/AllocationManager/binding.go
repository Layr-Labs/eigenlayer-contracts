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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_delegation\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"_permissionController\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"},{\"name\":\"_DEALLOCATION_DELAY\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_ALLOCATION_CONFIGURATION_DELAY\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ALLOCATION_CONFIGURATION_DELAY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEALLOCATION_DELAY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addStrategiesToOperatorSet\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"clearDeallocationQueue\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"numToClear\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createOperatorSets\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"params\",\"type\":\"tuple[]\",\"internalType\":\"structIAllocationManagerTypes.CreateSetParams[]\",\"components\":[{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregisterFromOperatorSets\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIAllocationManagerTypes.DeregisterParams\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"encumberedMagnitude\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAVSRegistrar\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAVSRegistrar\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocatableMagnitude\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocatedSets\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocatedStrategies\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocation\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIAllocationManagerTypes.Allocation\",\"components\":[{\"name\":\"currentMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"pendingDiff\",\"type\":\"int128\",\"internalType\":\"int128\"},{\"name\":\"effectBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocationDelay\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocations\",\"inputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIAllocationManagerTypes.Allocation[]\",\"components\":[{\"name\":\"currentMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"pendingDiff\",\"type\":\"int128\",\"internalType\":\"int128\"},{\"name\":\"effectBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxMagnitude\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxMagnitudes\",\"inputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxMagnitudes\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxMagnitudesAtBlock\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMemberCount\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMembers\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMinimumSlashableStake\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"futureBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"slashableStake\",\"type\":\"uint256[][]\",\"internalType\":\"uint256[][]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorSetCount\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRegisteredSets\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStrategiesInOperatorSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStrategyAllocations\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIAllocationManagerTypes.Allocation[]\",\"components\":[{\"name\":\"currentMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"pendingDiff\",\"type\":\"int128\",\"internalType\":\"int128\"},{\"name\":\"effectBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isMemberOfOperatorSet\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperatorSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"modifyAllocations\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"params\",\"type\":\"tuple[]\",\"internalType\":\"structIAllocationManagerTypes.AllocateParams[]\",\"components\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"newMagnitudes\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"permissionController\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerForOperatorSets\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIAllocationManagerTypes.RegisterParams\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeStrategiesFromOperatorSet\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAVSRegistrar\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"registrar\",\"type\":\"address\",\"internalType\":\"contractIAVSRegistrar\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAllocationDelay\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delay\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slashOperator\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIAllocationManagerTypes.SlashingParams\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"wadsToSlash\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAVSMetadataURI\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AVSMetadataURIUpdated\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"metadataURI\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AVSRegistrarSet\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"registrar\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIAVSRegistrar\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllocationDelaySet\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"delay\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"effectBlock\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllocationUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"magnitude\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"effectBlock\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EncumberedMagnitudeUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"encumberedMagnitude\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MaxMagnitudeUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"maxMagnitude\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorAddedToOperatorSet\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRemovedFromOperatorSet\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSetCreated\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSlashed\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategies\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"contractIStrategy[]\"},{\"name\":\"wadSlashed\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"description\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyAddedToOperatorSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyRemovedFromOperatorSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyMemberOfSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CurrentlyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Empty\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputArrayLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientMagnitude\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAVSRegistrar\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAVSWithNoMetadataRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCaller\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNewPausedStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperatorSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPermissions\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSnapshotOrdering\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidWadToSlash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ModificationAlreadyPending\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotMemberOfSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyPauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnpauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorNotSlashable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OutOfBounds\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SameMagnitude\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategiesMustBeInAscendingOrder\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategyAlreadyInOperatorSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategyNotInOperatorSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UninitializedAllocationDelay\",\"inputs\":[]}]",
	Bin: "0x610120604052348015610010575f5ffd5b50604051615b9c380380615b9c83398101604081905261002f91610180565b82858383876001600160a01b03811661005b576040516339b190bb60e11b815260040160405180910390fd5b6001600160a01b0390811660805292831660a05263ffffffff91821660c0521660e052166101005261008b610095565b50505050506101e9565b5f54610100900460ff16156101005760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff9081161461014f575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b0381168114610165575f5ffd5b50565b805163ffffffff8116811461017b575f5ffd5b919050565b5f5f5f5f5f60a08688031215610194575f5ffd5b855161019f81610151565b60208701519095506101b081610151565b60408701519094506101c181610151565b92506101cf60608701610168565b91506101dd60808701610168565b90509295509295909350565b60805160a05160c05160e0516101005161592261027a5f395f81816103f9015261382101525f81816105480152613e1901525f818161031e01528181612276015261297e01525f81816106fa01528181610ccf0152818161165401528181611cbb01528181611d250152612c8701525f818161056f0152818161079201528181611dca015261349801526159225ff3fe608060405234801561000f575f5ffd5b5060043610610281575f3560e01c80636e875dba11610156578063a984eb3a116100ca578063c221d8ae11610084578063c221d8ae146106bc578063cd6dc687146106cf578063d3d96ff4146106e2578063df5cf723146106f5578063f2fde38b1461071c578063fabc1cbc1461072f575f5ffd5b8063a984eb3a1461060e578063adc2e3d914610641578063b2447af714610654578063b66bd98914610667578063b9fbaed11461067a578063ba1a84e5146106a9575f5ffd5b80638ce648541161011b5780638ce64854146105915780638da5cb5b146105b157806394d7d00c146105c2578063952899ee146105d5578063a9333ec8146105e8578063a9821821146105fb575f5ffd5b80636e875dba14610515578063715018a61461052857806379ae50cd146105305780637bc1ef6114610543578063886f11951461056a575f5ffd5b80634657e26a116101f8578063595c6a67116101b2578063595c6a67146104875780635ac86ab71461048f5780635c975abb146104b2578063670d3ba2146104c45780636cfb4481146104d75780636e3492b514610502575f5ffd5b80634657e26a146103f45780634a10ffe51461041b5780634b5046ef1461043b57806350feea201461044e578063547afb871461046157806356c483e614610474575f5ffd5b80632981eb77116102495780632981eb77146103195780632bab2c4a14610355578063304c10cd1461037557806336352057146103a057806340120dab146103b35780634177a87c146103d4575f5ffd5b806310e1b9b814610285578063136439dd146102ae57806315fe5028146102c3578063260dc758146102e3578063261f84e014610306575b5f5ffd5b61029861029336600461483e565b610742565b6040516102a59190614885565b60405180910390f35b6102c16102bc3660046148b8565b61077d565b005b6102d66102d13660046148cf565b610852565b6040516102a5919061494d565b6102f66102f136600461495f565b610969565b60405190151581526020016102a5565b6102c16103143660046149b9565b6109a0565b6103407f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff90911681526020016102a5565b610368610363366004614a9e565b610c7b565b6040516102a59190614b52565b6103886103833660046148cf565b610f68565b6040516001600160a01b0390911681526020016102a5565b6102c16103ae366004614bb5565b610f97565b6103c66103c1366004614c07565b6117a9565b6040516102a5929190614c94565b6103e76103e236600461495f565b611924565b6040516102a59190614cf1565b6103887f000000000000000000000000000000000000000000000000000000000000000081565b61042e610429366004614d03565b611948565b6040516102a59190614d46565b6102c1610449366004614d91565b6119f0565b6102c161045c366004614e11565b611aaa565b61042e61046f366004614e6f565b611c08565b6102c1610482366004614ebb565b611cb0565b6102c1611db5565b6102f661049d366004614eee565b606654600160ff9092169190911b9081161490565b6066545b6040519081526020016102a5565b6102f66104d2366004614f0e565b611e64565b6104ea6104e5366004614c07565b611e75565b6040516001600160401b0390911681526020016102a5565b6102c1610510366004614f4f565b611fe2565b6103e761052336600461495f565b6123c3565b6102c16123d4565b6102d661053e3660046148cf565b6123e5565b6103407f000000000000000000000000000000000000000000000000000000000000000081565b6103887f000000000000000000000000000000000000000000000000000000000000000081565b6105a461059f366004614f80565b6124bf565b6040516102a59190614fc3565b6033546001600160a01b0316610388565b61042e6105d0366004614fd5565b612584565b6102c16105e3366004615030565b612670565b6104ea6105f6366004614c07565b612b37565b6102c16106093660046151d9565b612b66565b6104ea61061c366004614c07565b60a260209081525f92835260408084209091529082529020546001600160401b031681565b6102c161064f366004615257565b612c18565b6104b661066236600461495f565b612f75565b6102c1610675366004614e11565b612f97565b61068d6106883660046148cf565b6130f1565b60408051921515835263ffffffff9091166020830152016102a5565b6104b66106b73660046148cf565b61318b565b6103e76106ca366004614f0e565b6131ab565b6102c16106dd366004615299565b6131dc565b6102c16106f0366004614c07565b6132f9565b6103887f000000000000000000000000000000000000000000000000000000000000000081565b6102c161072a3660046148cf565b61341d565b6102c161073d3660046148b8565b613496565b604080516060810182525f80825260208201819052918101829052906107718561076b866135ac565b8561360f565b925050505b9392505050565b60405163237dfb4760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906346fbf68e90602401602060405180830381865afa1580156107df573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061080391906152c3565b61082057604051631d77d47760e21b815260040160405180910390fd5b60665481811681146108455760405163c61dca5d60e01b815260040160405180910390fd5b61084e8261377b565b5050565b6001600160a01b0381165f908152609d6020526040812060609190610876906137b8565b90505f816001600160401b0381111561089157610891614762565b6040519080825280602002602001820160405280156108d557816020015b604080518082019091525f80825260208201528152602001906001900390816108af5790505b5090505f5b82811015610961576001600160a01b0385165f908152609d6020526040902061093c9061090790836137c1565b604080518082019091525f80825260208201525060408051808201909152606082901c815263ffffffff909116602082015290565b82828151811061094e5761094e6152e2565b60209081029190910101526001016108da565b509392505050565b60208082015182516001600160a01b03165f90815260989092526040822061099a9163ffffffff908116906137cc16565b92915050565b826109aa816137e3565b6109c75760405163932d94f760e01b815260040160405180910390fd5b6001600160a01b0384165f90815260a4602052604090205460ff166109ff5760405163caf9756160e01b815260040160405180910390fd5b5f5b82811015610c74575f6040518060400160405280876001600160a01b03168152602001868685818110610a3657610a366152e2565b9050602002810190610a4891906152f6565b610a56906020810190615314565b63ffffffff168152509050610aa0816020015163ffffffff1660985f896001600160a01b03166001600160a01b031681526020019081526020015f2061388d90919063ffffffff16565b610abd57604051631fb1705560e21b815260040160405180910390fd5b7f31629285ead2335ae0933f86ed2ae63321f7af77b4e6eaabc42c057880977e6c6040518060400160405280886001600160a01b03168152602001836020015163ffffffff16815250604051610b13919061532d565b60405180910390a15f610b25826135ac565b90505f5b868685818110610b3b57610b3b6152e2565b9050602002810190610b4d91906152f6565b610b5b90602081019061533b565b9050811015610c6957610bd1878786818110610b7957610b796152e2565b9050602002810190610b8b91906152f6565b610b9990602081019061533b565b83818110610ba957610ba96152e2565b9050602002016020810190610bbe91906148cf565b5f84815260996020526040902090613898565b507f7ab260fe0af193db5f4986770d831bda4ea46099dc817e8b6716dcae8af8e88b83888887818110610c0657610c066152e2565b9050602002810190610c1891906152f6565b610c2690602081019061533b565b84818110610c3657610c366152e2565b9050602002016020810190610c4b91906148cf565b604051610c59929190615380565b60405180910390a1600101610b29565b505050600101610a01565b5050505050565b606083516001600160401b03811115610c9657610c96614762565b604051908082528060200260200182016040528015610cc957816020015b6060815260200190600190039081610cb45790505b5090505f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f0e0e67686866040518363ffffffff1660e01b8152600401610d1b9291906153a6565b5f60405180830381865afa158015610d35573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610d5c91908101906153ca565b90505f5b8551811015610f5e575f868281518110610d7c57610d7c6152e2565b6020026020010151905085516001600160401b03811115610d9f57610d9f614762565b604051908082528060200260200182016040528015610dc8578160200160208202803683370190505b50848381518110610ddb57610ddb6152e2565b60209081029190910101525f5b8651811015610f54575f878281518110610e0457610e046152e2565b6020908102919091018101516001600160a01b038086165f90815260a1845260408082209284168252919093528220909250610e3f906138ac565b9050806001600160401b03165f03610e58575050610f4c565b5f610e64858d85610742565b90508863ffffffff16816040015163ffffffff1611158015610e8c57505f8160200151600f0b125b15610eae57610ea2815f015182602001516138bf565b6001600160401b031681525b80515f90610ec9906001600160401b039081169085166138d3565b9050610f1081898981518110610ee157610ee16152e2565b60200260200101518781518110610efa57610efa6152e2565b60200260200101516138e790919063ffffffff16565b898881518110610f2257610f226152e2565b60200260200101518681518110610f3b57610f3b6152e2565b602002602001018181525050505050505b600101610de8565b5050600101610d60565b5050949350505050565b6001600160a01b038082165f908152609760205260408120549091168015610f905780610776565b5090919050565b606654600190600290811603610fc05760405163840a48d560e01b815260040160405180910390fd5b82610fca816137e3565b610fe75760405163932d94f760e01b815260040160405180910390fd5b5f6040518060400160405280866001600160a01b031681526020018560200160208101906110159190615314565b63ffffffff16905290505f61103661103060208701876148cf565b836138fb565b9050611045606086018661533b565b9050611054604087018761533b565b905014611074576040516343714afd60e01b815260040160405180910390fd5b60208083015183516001600160a01b03165f908152609890925260409091206110a69163ffffffff908116906137cc16565b6110c357604051631fb1705560e21b815260040160405180910390fd5b806110e15760405163ebbff49760e01b815260040160405180910390fd5b5f6110ef604087018761533b565b90506001600160401b0381111561110857611108614762565b604051908082528060200260200182016040528015611131578160200160208202803683370190505b5090505f5b611143604088018861533b565b905081101561173a578015806111d65750611161604088018861533b565b61116c6001846154ea565b81811061117b5761117b6152e2565b905060200201602081019061119091906148cf565b6001600160a01b03166111a6604089018961533b565b838181106111b6576111b66152e2565b90506020020160208101906111cb91906148cf565b6001600160a01b0316115b6111f357604051639f1c805360e01b815260040160405180910390fd5b611200606088018861533b565b82818110611210576112106152e2565b905060200201355f1080156112505750670de0b6b3a7640000611236606089018961533b565b83818110611246576112466152e2565b9050602002013511155b61126d57604051631353603160e01b815260040160405180910390fd5b6112c961127d604089018961533b565b8381811061128d5761128d6152e2565b90506020020160208101906112a291906148cf565b60995f6112ae886135ac565b81526020019081526020015f2061397190919063ffffffff16565b6112e6576040516331bc342760e11b815260040160405180910390fd5b5f806113386112f860208b018b6148cf565b611301886135ac565b61130e60408d018d61533b565b8781811061131e5761131e6152e2565b905060200201602081019061133391906148cf565b61360f565b805191935091506001600160401b03165f03611355575050611732565b5f61139061136660608c018c61533b565b86818110611376576113766152e2565b85516001600160401b031692602090910201359050613992565b83519091506113ab6001600160401b038084169083166138d3565b8686815181106113bd576113bd6152e2565b60200260200101818152505081835f018181516113da91906154fd565b6001600160401b03169052508351829085906113f79083906154fd565b6001600160401b03169052506020840180518391906114179083906154fd565b6001600160401b031690525060208301515f600f9190910b1215611532575f61147a61144660608e018e61533b565b88818110611456576114566152e2565b90506020020135856020015161146b9061551c565b6001600160801b031690613992565b9050806001600160401b0316846020018181516114979190615540565b600f0b9052507f1487af5418c47ee5ea45ef4a93398668120890774a9e13487e61e9dc3baf76dd6114cb60208e018e6148cf565b8a8e80604001906114dc919061533b565b8a8181106114ec576114ec6152e2565b905060200201602081019061150191906148cf565b611512885f015189602001516138bf565b886040015160405161152895949392919061556d565b60405180910390a1505b61158461154260208d018d6148cf565b61154b8a6135ac565b61155860408f018f61533b565b89818110611568576115686152e2565b905060200201602081019061157d91906148cf565b87876139a8565b7f1487af5418c47ee5ea45ef4a93398668120890774a9e13487e61e9dc3baf76dd6115b260208d018d6148cf565b896115c060408f018f61533b565b898181106115d0576115d06152e2565b90506020020160208101906115e591906148cf565b86516040516115f99493929190439061556d565b60405180910390a161164a61161160208d018d6148cf565b61161e60408e018e61533b565b8881811061162e5761162e6152e2565b905060200201602081019061164391906148cf565b8651613be0565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001663ee74937f61168660208e018e6148cf565b61169360408f018f61533b565b898181106116a3576116a36152e2565b90506020020160208101906116b891906148cf565b875160405160e085901b6001600160e01b03191681526001600160a01b0393841660048201529290911660248301526001600160401b0380861660448401521660648201526084015f604051808303815f87803b158015611717575f5ffd5b505af1158015611729573d5f5f3e3d5ffd5b50505050505050505b600101611136565b507f80969ad29428d6797ee7aad084f9e4a42a82fc506dcd2ca3b6fb431f85ccebe561176960208801886148cf565b8461177760408a018a61533b565b8561178560808d018d6155be565b6040516117989796959493929190615628565b60405180910390a150505050505050565b6001600160a01b0382165f908152609d6020526040812060609182916117ce906137b8565b90505f816001600160401b038111156117e9576117e9614762565b60405190808252806020026020018201604052801561182d57816020015b604080518082019091525f80825260208201528152602001906001900390816118075790505b5090505f826001600160401b0381111561184957611849614762565b60405190808252806020026020018201604052801561189257816020015b604080516060810182525f80825260208083018290529282015282525f199092019101816118675790505b5090505f5b83811015611915576001600160a01b0388165f908152609d602052604081206118c49061090790846137c1565b9050808483815181106118d9576118d96152e2565b60200260200101819052506118ef89828a610742565b838381518110611901576119016152e2565b602090810291909101015250600101611897565b509093509150505b9250929050565b60605f61077660995f611936866135ac565b81526020019081526020015f20613c62565b60605f83516001600160401b0381111561196457611964614762565b60405190808252806020026020018201604052801561198d578160200160208202803683370190505b5090505f5b8451811015610961576119be8582815181106119b0576119b06152e2565b602002602001015185612b37565b8282815181106119d0576119d06152e2565b6001600160401b0390921660209283029190910190910152600101611992565b6066545f90600190811603611a185760405163840a48d560e01b815260040160405180910390fd5b838214611a38576040516343714afd60e01b815260040160405180910390fd5b5f5b84811015611aa157611a9987878784818110611a5857611a586152e2565b9050602002016020810190611a6d91906148cf565b868685818110611a7f57611a7f6152e2565b9050602002016020810190611a9491906156be565b613c6e565b600101611a3a565b50505050505050565b83611ab4816137e3565b611ad15760405163932d94f760e01b815260040160405180910390fd5b604080518082019091526001600160a01b038616815263ffffffff851660208201525f611afd826135ac565b9050611b3e826020015163ffffffff1660985f8a6001600160a01b03166001600160a01b031681526020019081526020015f206137cc90919063ffffffff16565b611b5b57604051631fb1705560e21b815260040160405180910390fd5b5f5b84811015611bfe57611b7a868683818110610ba957610ba96152e2565b611b975760405163585cfb2f60e01b815260040160405180910390fd5b7f7ab260fe0af193db5f4986770d831bda4ea46099dc817e8b6716dcae8af8e88b83878784818110611bcb57611bcb6152e2565b9050602002016020810190611be091906148cf565b604051611bee929190615380565b60405180910390a1600101611b5d565b5050505050505050565b60605f82516001600160401b03811115611c2457611c24614762565b604051908082528060200260200182016040528015611c4d578160200160208202803683370190505b5090505f5b835181101561096157611c7e85858381518110611c7157611c716152e2565b6020026020010151612b37565b828281518110611c9057611c906152e2565b6001600160401b0390921660209283029190910190910152600101611c52565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614611dab57611ce9826137e3565b611d06576040516348f5c3ed60e01b815260040160405180910390fd5b6040516336b87bd760e11b81526001600160a01b0383811660048301527f00000000000000000000000000000000000000000000000000000000000000001690636d70f7ae90602401602060405180830381865afa158015611d6a573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611d8e91906152c3565b611dab5760405163ccea9e6f60e01b815260040160405180910390fd5b61084e8282613d72565b60405163237dfb4760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906346fbf68e90602401602060405180830381865afa158015611e17573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611e3b91906152c3565b611e5857604051631d77d47760e21b815260040160405180910390fd5b611e625f1961377b565b565b5f61077683609a5f6112ae866135ac565b6001600160a01b038281165f81815260a2602090815260408083209486168084529482528083205493835260a38252808320948352939052918220546001600160401b0390911690600f81810b600160801b909204900b03825b81811015611f9f576001600160a01b038087165f90815260a3602090815260408083209389168352929052908120611f079083613f13565b6001600160a01b038881165f90815260a0602090815260408083208584528252808320938b16835292815290829020825160608101845290546001600160401b0381168252600160401b8104600f0b92820192909252600160c01b90910463ffffffff16918101829052919250431015611f82575050611f9f565b611f908582602001516138bf565b94505050806001019050611ecf565b506001600160a01b038086165f90815260a1602090815260408083209388168352929052208290611fcf906138ac565b611fd991906154fd565b95945050505050565b60665460029060049081160361200b5760405163840a48d560e01b815260040160405180910390fd5b61202061201b60208401846148cf565b6137e3565b80612039575061203961201b60408401602085016148cf565b612056576040516348f5c3ed60e01b815260040160405180910390fd5b5f5b612065604084018461533b565b9050811015612327575f604051806040016040528085602001602081019061208d91906148cf565b6001600160a01b031681526020016120a8604087018761533b565b858181106120b8576120b86152e2565b90506020020160208101906120cd9190615314565b63ffffffff16815250905061211a816020015163ffffffff1660985f8760200160208101906120fc91906148cf565b6001600160a01b0316815260208101919091526040015f20906137cc565b61213757604051631fb1705560e21b815260040160405180910390fd5b609e5f61214760208701876148cf565b6001600160a01b03166001600160a01b031681526020019081526020015f205f612170836135ac565b815260208101919091526040015f205460ff166121a0576040516325131d4f60e01b815260040160405180910390fd5b6121da6121ac826135ac565b609c5f6121bc60208901896148cf565b6001600160a01b0316815260208101919091526040015f2090613f82565b506122126121eb60208601866148cf565b609a5f6121f7856135ac565b81526020019081526020015f20613f8d90919063ffffffff16565b5061222060208501856148cf565b6001600160a01b03167fad34c3070be1dffbcaa499d000ba2b8d9848aefcac3059df245dd95c4ece14fe82604051612258919061532d565b60405180910390a2604080518082019091525f81526020810161229b7f0000000000000000000000000000000000000000000000000000000000000000436156df565b63ffffffff169052609e5f6122b360208801886148cf565b6001600160a01b03166001600160a01b031681526020019081526020015f205f6122dc846135ac565b81526020808201929092526040015f2082518154939092015163ffffffff166101000264ffffffff00199215159290921664ffffffffff199093169290921717905550600101612058565b5061233b61038360408401602085016148cf565b6001600160a01b031663303ca95661235660208501856148cf565b61236660408601602087016148cf565b612373604087018761533b565b6040518563ffffffff1660e01b81526004016123929493929190615734565b5f604051808303815f87803b1580156123a9575f5ffd5b505af11580156123bb573d5f5f3e3d5ffd5b505050505050565b606061099a609a5f611936856135ac565b6123dc613fa1565b611e625f613ffb565b6001600160a01b0381165f908152609c6020526040812060609190612409906137b8565b90505f816001600160401b0381111561242457612424614762565b60405190808252806020026020018201604052801561246857816020015b604080518082019091525f80825260208201528152602001906001900390816124425790505b5090505f5b82811015610961576001600160a01b0385165f908152609c6020526040902061249a9061090790836137c1565b8282815181106124ac576124ac6152e2565b602090810291909101015260010161246d565b60605f84516001600160401b038111156124db576124db614762565b60405190808252806020026020018201604052801561252457816020015b604080516060810182525f80825260208083018290529282015282525f199092019101816124f95790505b5090505f5b855181101561257b57612556868281518110612547576125476152e2565b60200260200101518686610742565b828281518110612568576125686152e2565b6020908102919091010152600101612529565b50949350505050565b60605f83516001600160401b038111156125a0576125a0614762565b6040519080825280602002602001820160405280156125c9578160200160208202803683370190505b5090505f5b845181101561257b576001600160a01b0386165f90815260a160205260408120865161263e92879291899086908110612609576126096152e2565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020015f2061404c90919063ffffffff16565b828281518110612650576126506152e2565b6001600160401b03909216602092830291909101909101526001016125ce565b6066545f906001908116036126985760405163840a48d560e01b815260040160405180910390fd5b6126a1836137e3565b6126be576040516348f5c3ed60e01b815260040160405180910390fd5b5f5f5f6126ca866130f1565b91509150816126ec5760405163fa55fc8160e01b815260040160405180910390fd5b91505f90505b8351811015610c745783818151811061270d5761270d6152e2565b6020026020010151604001515184828151811061272c5761272c6152e2565b6020026020010151602001515114612757576040516343714afd60e01b815260040160405180910390fd5b5f84828151811061276a5761276a6152e2565b602090810291909101810151518082015181516001600160a01b03165f908152609890935260409092209092506127aa9163ffffffff908116906137cc16565b6127c757604051631fb1705560e21b815260040160405180910390fd5b5f6127d287836138fb565b90505f5b8684815181106127e8576127e86152e2565b60200260200101516020015151811015612b2c575f87858151811061280f5761280f6152e2565b602002602001015160200151828151811061282c5761282c6152e2565b60200260200101519050612843898261ffff613c6e565b5f5f6128528b61076b886135ac565b915091508060200151600f0b5f1461287d57604051630d8fcbe360e41b815260040160405180910390fd5b5f61288a87858489614060565b90506128cf825f01518c8a815181106128a5576128a56152e2565b60200260200101516040015187815181106128c2576128c26152e2565b6020026020010151614096565b600f0b602083018190525f036128f857604051634606179360e11b815260040160405180910390fd5b5f8260200151600f0b1215612a3c5780156129be57612979612919886135ac565b6001600160a01b03808f165f90815260a360209081526040808320938a16835292905220908154600160801b90819004600f0b5f818152600180860160205260409091209390935583546001600160801b03908116939091011602179055565b6129a37f0000000000000000000000000000000000000000000000000000000000000000436156df565b6129ae9060016156df565b63ffffffff166040830152612aa9565b6129d0836020015183602001516138bf565b6001600160401b031660208401528a518b90899081106129f2576129f26152e2565b6020026020010151604001518581518110612a0f57612a0f6152e2565b6020908102919091018101516001600160401b031683525f9083015263ffffffff43166040830152612aa9565b5f8260200151600f0b1315612aa957612a5d836020015183602001516138bf565b6001600160401b039081166020850181905284519091161015612a9357604051636c9be0bf60e01b815260040160405180910390fd5b612a9d89436156df565b63ffffffff1660408301525b612abe8c612ab6896135ac565b8686866139a8565b7f1487af5418c47ee5ea45ef4a93398668120890774a9e13487e61e9dc3baf76dd8c612aec6109078a6135ac565b86612afe865f015187602001516138bf565b8660400151604051612b1495949392919061556d565b60405180910390a15050600190920191506127d69050565b5050506001016126f2565b6001600160a01b038083165f90815260a1602090815260408083209385168352929052908120610776906138ac565b82612b70816137e3565b612b8d5760405163932d94f760e01b815260040160405180910390fd5b6001600160a01b0384165f90815260a4602052604090205460ff16612bcf576001600160a01b0384165f90815260a460205260409020805460ff191660011790555b836001600160a01b03167fa89c1dc243d8908a96dd84944bcc97d6bc6ac00dd78e20621576be6a3c9437138484604051612c0a929190615760565b60405180910390a250505050565b606654600290600490811603612c415760405163840a48d560e01b815260040160405180910390fd5b82612c4b816137e3565b612c685760405163932d94f760e01b815260040160405180910390fd5b6040516336b87bd760e11b81526001600160a01b0385811660048301527f00000000000000000000000000000000000000000000000000000000000000001690636d70f7ae90602401602060405180830381865afa158015612ccc573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612cf091906152c3565b612d0d5760405163ccea9e6f60e01b815260040160405180910390fd5b5f5b612d1c602085018561533b565b9050811015612ee457604080518082019091525f9080612d3f60208801886148cf565b6001600160a01b03168152602001868060200190612d5d919061533b565b85818110612d6d57612d6d6152e2565b9050602002016020810190612d829190615314565b63ffffffff90811690915260208083015183516001600160a01b03165f90815260989092526040909120929350612dbe9291908116906137cc16565b612ddb57604051631fb1705560e21b815260040160405180910390fd5b612de586826138fb565b15612e0357604051636c6c6e2760e11b815260040160405180910390fd5b612e2c612e0f826135ac565b6001600160a01b0388165f908152609c602052604090209061388d565b50612e5886609a5f612e3d856135ac565b81526020019081526020015f2061389890919063ffffffff16565b50856001600160a01b03167f43232edf9071753d2321e5fa7e018363ee248e5f2142e6c08edd3265bfb4895e82604051612e92919061532d565b60405180910390a26001600160a01b0386165f908152609e60205260408120600191612ebd846135ac565b815260208101919091526040015f20805460ff191691151591909117905550600101612d0f565b50612ef561038360208501856148cf565b6001600160a01b031663c63fd50285612f1160208701876148cf565b612f1e602088018861533b565b612f2b60408a018a6155be565b6040518763ffffffff1660e01b8152600401612f4c96959493929190615773565b5f604051808303815f87803b158015612f63575f5ffd5b505af1158015611bfe573d5f5f3e3d5ffd5b5f61099a609a5f612f85856135ac565b81526020019081526020015f206137b8565b83612fa1816137e3565b612fbe5760405163932d94f760e01b815260040160405180910390fd5b6040805180820182526001600160a01b03871680825263ffffffff80881660208085018290525f93845260989052939091209192612ffd92916137cc16565b61301a57604051631fb1705560e21b815260040160405180910390fd5b5f613024826135ac565b90505f5b84811015611bfe5761306d868683818110613045576130456152e2565b905060200201602081019061305a91906148cf565b5f84815260996020526040902090613f8d565b61308a576040516331bc342760e11b815260040160405180910390fd5b7f7b4b073d80dcac55a11177d8459ad9f664ceeb91f71f27167bb14f8152a7eeee838787848181106130be576130be6152e2565b90506020020160208101906130d391906148cf565b6040516130e1929190615380565b60405180910390a1600101613028565b6001600160a01b0381165f908152609b602090815260408083208151608081018352905463ffffffff80821680845260ff600160201b8404161515958401869052650100000000008304821694840194909452600160481b9091041660608201819052849391929190158015906131725750826060015163ffffffff164310155b15613181575050604081015160015b9590945092505050565b6001600160a01b0381165f90815260986020526040812061099a906137b8565b6001600160a01b0382165f908152609f60205260408120606091906131d49082611936866135ac565b949350505050565b5f54610100900460ff16158080156131fa57505f54600160ff909116105b806132135750303b15801561321357505f5460ff166001145b61327b5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b5f805460ff19166001179055801561329c575f805461ff0019166101001790555b6132a58261377b565b6132ae83613ffb565b80156132f4575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498906020015b60405180910390a15b505050565b81613303816137e3565b6133205760405163932d94f760e01b815260040160405180910390fd5b60405163b526578760e01b81526001600160a01b03848116600483015283169063b526578790602401602060405180830381865afa158015613364573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061338891906152c3565b6133a557604051631d0b13c160e31b815260040160405180910390fd5b6001600160a01b038381165f90815260976020526040902080546001600160a01b0319169184169190911790557f2ae945c40c44dc0ec263f95609c3fdc6952e0aefa22d6374e44f2c997acedf85836133fd81610f68565b604080516001600160a01b039384168152929091166020830152016132eb565b613425613fa1565b6001600160a01b03811661348a5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401613272565b61349381613ffb565b50565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156134f2573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061351691906157bf565b6001600160a01b0316336001600160a01b0316146135475760405163794821ff60e01b815260040160405180910390fd5b6066548019821981161461356e5760405163c61dca5d60e01b815260040160405180910390fd5b606682905560405182815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c9060200160405180910390a25050565b5f815f0151826020015163ffffffff166040516020016135f792919060609290921b6bffffffffffffffffffffffff1916825260a01b6001600160a01b031916601482015260200190565b60405160208183030381529060405261099a906157da565b6040805180820182525f80825260208083018290528351606081018552828152808201839052808501839052845180860186526001600160a01b03898116855260a184528685209088168552909252938220929392819061366f906138ac565b6001600160401b0390811682526001600160a01b038981165f81815260a260209081526040808320948c168084529482528083205486169682019690965291815260a082528481208b8252825284812092815291815290839020835160608101855290549283168152600160401b8304600f0b91810191909152600160c01b90910463ffffffff16918101829052919250431015613711579092509050613773565b613722815f015182602001516138bf565b6001600160401b0316815260208101515f600f9190910b121561376057613751826020015182602001516138bf565b6001600160401b031660208301525b5f60408201819052602082015290925090505b935093915050565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a250565b5f61099a825490565b5f61077683836140ad565b5f8181526001830160205260408120541515610776565b604051631beb2b9760e31b81526001600160a01b0382811660048301523360248301523060448301525f80356001600160e01b0319166064840152917f00000000000000000000000000000000000000000000000000000000000000009091169063df595cb8906084016020604051808303815f875af1158015613869573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061099a91906152c3565b5f61077683836140d3565b5f610776836001600160a01b0384166140d3565b5f61099a82670de0b6b3a764000061411f565b5f610776826001600160401b038516615540565b5f61077683670de0b6b3a764000084614163565b5f6107768383670de0b6b3a7640000614163565b6001600160a01b0382165f908152609e6020526040812081908161391e856135ac565b815260208082019290925260409081015f2081518083019092525460ff8116151580835261010090910463ffffffff16928201929092529150806131d457506020015163ffffffff164311159392505050565b6001600160a01b0381165f9081526001830160205260408120541515610776565b5f6107768383670de0b6b3a76400006001614248565b6020808301516001600160a01b038088165f90815260a284526040808220928816825291909352909120546001600160401b03908116911614613a6e57602082810180516001600160a01b038881165f81815260a286526040808220938a1680835293875290819020805467ffffffffffffffff19166001600160401b0395861617905593518451918252948101919091529216908201527facf9095feb3a370c9cf692421c69ef320d4db5c66e6a7d29c7694eb02364fc559060600160405180910390a15b6001600160a01b038086165f90815260a060209081526040808320888452825280832093871683529281529082902083518154928501519385015163ffffffff16600160c01b0263ffffffff60c01b196001600160801b038616600160401b026001600160c01b03199095166001600160401b03909316929092179390931716919091179055600f0b15613b50576001600160a01b0385165f908152609f602090815260408083208784529091529020613b289084613898565b506001600160a01b0385165f908152609d60205260409020613b4a908561388d565b50610c74565b80516001600160401b03165f03610c74576001600160a01b0385165f908152609f602090815260408083208784529091529020613b8d9084613f8d565b506001600160a01b0385165f908152609f602090815260408083208784529091529020613bb9906137b8565b5f03610c74576001600160a01b0385165f908152609d602052604090206123bb9085613f82565b6001600160a01b038084165f90815260a160209081526040808320938616835292905220613c0f9043836142a1565b604080516001600160a01b038086168252841660208201526001600160401b038316918101919091527f1c6458079a41077d003c11faf9bf097e693bd67979e4e6500bac7b29db779b5c906060016132eb565b60605f610776836142b5565b6001600160a01b038381165f90815260a360209081526040808320938616835292905290812054600f81810b600160801b909204900b035b5f81118015613cb857508261ffff1682105b15610c74576001600160a01b038086165f90815260a3602090815260408083209388168352929052908120613cec9061430e565b90505f5f613cfb88848961360f565b91509150806040015163ffffffff16431015613d1957505050610c74565b613d2688848985856139a8565b6001600160a01b038089165f90815260a360209081526040808320938b16835292905220613d5390614360565b50613d5d856157fd565b9450613d6884615815565b9350505050613ca6565b6001600160a01b0382165f908152609b60209081526040918290208251608081018452905463ffffffff808216835260ff600160201b830416151593830193909352650100000000008104831693820193909352600160481b909204166060820181905215801590613dee5750806060015163ffffffff164310155b15613e0857604081015163ffffffff168152600160208201525b63ffffffff82166040820152613e3e7f0000000000000000000000000000000000000000000000000000000000000000436156df565b63ffffffff90811660608381019182526001600160a01b0386165f818152609b602090815260409182902087518154838a0151858b01519851928a1664ffffffffff1990921691909117600160201b91151591909102176cffffffffffffffff0000000000191665010000000000978916979097026cffffffff000000000000000000191696909617600160481b968816968702179055815192835294871694820194909452928301919091527f4e85751d6331506c6c62335f207eb31f12a61e570f34f5c17640308785c6d4db91016132eb565b5f5f613f35613f21846143dd565b8554613f309190600f0b61582a565b61444a565b8454909150600160801b9004600f90810b9082900b12613f6857604051632d0483c560e21b815260040160405180910390fd5b600f0b5f9081526001939093016020525050604090205490565b5f61077683836144b3565b5f610776836001600160a01b0384166144b3565b6033546001600160a01b03163314611e625760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401613272565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b5f6107768383670de0b6b3a7640000614596565b5f6140718460995f6112ae896135ac565b801561407a5750815b8015611fd957505090516001600160401b031615159392505050565b5f6107766001600160401b03808516908416615851565b5f825f0182815481106140c2576140c26152e2565b905f5260205f200154905092915050565b5f81815260018301602052604081205461411857508154600181810184555f84815260208082209093018490558454848252828601909352604090209190915561099a565b505f61099a565b81545f90801561415b57614145846141386001846154ea565b5f91825260209091200190565b54600160201b90046001600160e01b03166131d4565b509092915050565b5f80805f19858709858702925082811083820303915050805f0361419a578382816141905761419061587e565b0492505050610776565b8084116141e15760405162461bcd60e51b81526020600482015260156024820152744d6174683a206d756c446976206f766572666c6f7760581b6044820152606401613272565b5f8486880960026001871981018816978890046003810283188082028403028082028403028082028403028082028403028082028403029081029092039091025f889003889004909101858311909403939093029303949094049190911702949350505050565b5f5f614255868686614163565b9050600183600281111561426b5761426b615892565b14801561428757505f84806142825761428261587e565b868809115b15611fd9576142976001826158a6565b9695505050505050565b6132f483836001600160401b0384166145de565b6060815f0180548060200260200160405190810160405280929190818152602001828054801561430257602002820191905f5260205f20905b8154815260200190600101908083116142ee575b50505050509050919050565b5f6143288254600f81810b600160801b909204900b131590565b1561434657604051631ed9509560e11b815260040160405180910390fd5b508054600f0b5f9081526001909101602052604090205490565b5f61437a8254600f81810b600160801b909204900b131590565b1561439857604051631ed9509560e11b815260040160405180910390fd5b508054600f0b5f818152600180840160205260408220805492905583546fffffffffffffffffffffffffffffffff191692016001600160801b03169190911790915590565b5f6001600160ff1b038211156144465760405162461bcd60e51b815260206004820152602860248201527f53616665436173743a2076616c756520646f65736e27742066697420696e2061604482015267371034b73a191a9b60c11b6064820152608401613272565b5090565b80600f81900b81146144ae5760405162461bcd60e51b815260206004820152602760248201527f53616665436173743a2076616c756520646f65736e27742066697420696e20316044820152663238206269747360c81b6064820152608401613272565b919050565b5f818152600183016020526040812054801561458d575f6144d56001836154ea565b85549091505f906144e8906001906154ea565b9050818114614547575f865f018281548110614506576145066152e2565b905f5260205f200154905080875f018481548110614526576145266152e2565b5f918252602080832090910192909255918252600188019052604090208390555b8554869080614558576145586158b9565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f90556001935050505061099a565b5f91505061099a565b82545f90816145a7868683856146e1565b905080156145d4576145be866141386001846154ea565b54600160201b90046001600160e01b0316610771565b5091949350505050565b82548015614694575f6145f6856141386001856154ea565b60408051808201909152905463ffffffff808216808452600160201b9092046001600160e01b0316602084015291925090851610156146485760405163151b8e3f60e11b815260040160405180910390fd5b805163ffffffff8086169116036146925782614669866141386001866154ea565b80546001600160e01b0392909216600160201b0263ffffffff9092169190911790555050505050565b505b506040805180820190915263ffffffff92831681526001600160e01b03918216602080830191825285546001810187555f968752952091519051909216600160201b029190921617910155565b5f5b81831015610961575f6146f68484614734565b5f8781526020902090915063ffffffff86169082015463ffffffff1611156147205780925061472e565b61472b8160016158a6565b93505b506146e3565b5f61474260028484186158cd565b610776908484166158a6565b6001600160a01b0381168114613493575f5ffd5b634e487b7160e01b5f52604160045260245ffd5b604051606081016001600160401b038111828210171561479857614798614762565b60405290565b604051601f8201601f191681016001600160401b03811182821017156147c6576147c6614762565b604052919050565b803563ffffffff811681146144ae575f5ffd5b5f604082840312156147f1575f5ffd5b604080519081016001600160401b038111828210171561481357614813614762565b60405290508082356148248161474e565b8152614832602084016147ce565b60208201525092915050565b5f5f5f60808486031215614850575f5ffd5b833561485b8161474e565b925061486a85602086016147e1565b9150606084013561487a8161474e565b809150509250925092565b81516001600160401b03168152602080830151600f0b9082015260408083015163ffffffff16908201526060810161099a565b5f602082840312156148c8575f5ffd5b5035919050565b5f602082840312156148df575f5ffd5b81356107768161474e565b80516001600160a01b0316825260209081015163ffffffff16910152565b5f8151808452602084019350602083015f5b828110156149435761492d8683516148ea565b604095909501946020919091019060010161491a565b5093949350505050565b602081525f6107766020830184614908565b5f6040828403121561496f575f5ffd5b61077683836147e1565b5f5f83601f840112614989575f5ffd5b5081356001600160401b0381111561499f575f5ffd5b6020830191508360208260051b850101111561191d575f5ffd5b5f5f5f604084860312156149cb575f5ffd5b83356149d68161474e565b925060208401356001600160401b038111156149f0575f5ffd5b6149fc86828701614979565b9497909650939450505050565b5f6001600160401b03821115614a2157614a21614762565b5060051b60200190565b5f82601f830112614a3a575f5ffd5b8135614a4d614a4882614a09565b61479e565b8082825260208201915060208360051b860101925085831115614a6e575f5ffd5b602085015b83811015614a94578035614a868161474e565b835260209283019201614a73565b5095945050505050565b5f5f5f5f60a08587031215614ab1575f5ffd5b614abb86866147e1565b935060408501356001600160401b03811115614ad5575f5ffd5b614ae187828801614a2b565b93505060608501356001600160401b03811115614afc575f5ffd5b614b0887828801614a2b565b925050614b17608086016147ce565b905092959194509250565b5f8151808452602084019350602083015f5b82811015614943578151865260209586019590910190600101614b34565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b82811015614ba957603f19878603018452614b94858351614b22565b94506020938401939190910190600101614b78565b50929695505050505050565b5f5f60408385031215614bc6575f5ffd5b8235614bd18161474e565b915060208301356001600160401b03811115614beb575f5ffd5b830160a08186031215614bfc575f5ffd5b809150509250929050565b5f5f60408385031215614c18575f5ffd5b8235614c238161474e565b91506020830135614bfc8161474e565b5f8151808452602084019350602083015f5b8281101561494357614c7e86835180516001600160401b03168252602080820151600f0b9083015260409081015163ffffffff16910152565b6060959095019460209190910190600101614c45565b604081525f614ca66040830185614908565b8281036020840152611fd98185614c33565b5f8151808452602084019350602083015f5b828110156149435781516001600160a01b0316865260209586019590910190600101614cca565b602081525f6107766020830184614cb8565b5f5f60408385031215614d14575f5ffd5b82356001600160401b03811115614d29575f5ffd5b614d3585828601614a2b565b9250506020830135614bfc8161474e565b602080825282518282018190525f918401906040840190835b81811015614d865783516001600160401b0316835260209384019390920191600101614d5f565b509095945050505050565b5f5f5f5f5f60608688031215614da5575f5ffd5b8535614db08161474e565b945060208601356001600160401b03811115614dca575f5ffd5b614dd688828901614979565b90955093505060408601356001600160401b03811115614df4575f5ffd5b614e0088828901614979565b969995985093965092949392505050565b5f5f5f5f60608587031215614e24575f5ffd5b8435614e2f8161474e565b9350614e3d602086016147ce565b925060408501356001600160401b03811115614e57575f5ffd5b614e6387828801614979565b95989497509550505050565b5f5f60408385031215614e80575f5ffd5b8235614e8b8161474e565b915060208301356001600160401b03811115614ea5575f5ffd5b614eb185828601614a2b565b9150509250929050565b5f5f60408385031215614ecc575f5ffd5b8235614ed78161474e565b9150614ee5602084016147ce565b90509250929050565b5f60208284031215614efe575f5ffd5b813560ff81168114610776575f5ffd5b5f5f60608385031215614f1f575f5ffd5b8235614f2a8161474e565b9150614ee584602085016147e1565b5f60608284031215614f49575f5ffd5b50919050565b5f60208284031215614f5f575f5ffd5b81356001600160401b03811115614f74575f5ffd5b6131d484828501614f39565b5f5f5f60808486031215614f92575f5ffd5b83356001600160401b03811115614fa7575f5ffd5b614fb386828701614a2b565b93505061486a85602086016147e1565b602081525f6107766020830184614c33565b5f5f5f60608486031215614fe7575f5ffd5b8335614ff28161474e565b925060208401356001600160401b0381111561500c575f5ffd5b61501886828701614a2b565b925050615027604085016147ce565b90509250925092565b5f5f60408385031215615041575f5ffd5b823561504c8161474e565b915060208301356001600160401b03811115615066575f5ffd5b8301601f81018513615076575f5ffd5b8035615084614a4882614a09565b8082825260208201915060208360051b8501019250878311156150a5575f5ffd5b602084015b838110156151ca5780356001600160401b038111156150c7575f5ffd5b85016080818b03601f190112156150dc575f5ffd5b6150e4614776565b6150f18b602084016147e1565b815260608201356001600160401b0381111561510b575f5ffd5b61511a8c602083860101614a2b565b60208301525060808201356001600160401b03811115615138575f5ffd5b6020818401019250508a601f83011261514f575f5ffd5b813561515d614a4882614a09565b8082825260208201915060208360051b86010192508d83111561517e575f5ffd5b6020850194505b828510156151b45784356001600160401b03811681146151a3575f5ffd5b825260209485019490910190615185565b60408401525050845250602092830192016150aa565b50809450505050509250929050565b5f5f5f604084860312156151eb575f5ffd5b83356151f68161474e565b925060208401356001600160401b03811115615210575f5ffd5b8401601f81018613615220575f5ffd5b80356001600160401b03811115615235575f5ffd5b866020828401011115615246575f5ffd5b939660209190910195509293505050565b5f5f60408385031215615268575f5ffd5b82356152738161474e565b915060208301356001600160401b0381111561528d575f5ffd5b614eb185828601614f39565b5f5f604083850312156152aa575f5ffd5b82356152b58161474e565b946020939093013593505050565b5f602082840312156152d3575f5ffd5b81518015158114610776575f5ffd5b634e487b7160e01b5f52603260045260245ffd5b5f8235603e1983360301811261530a575f5ffd5b9190910192915050565b5f60208284031215615324575f5ffd5b610776826147ce565b6040810161099a82846148ea565b5f5f8335601e19843603018112615350575f5ffd5b8301803591506001600160401b03821115615369575f5ffd5b6020019150600581901b360382131561191d575f5ffd5b6060810161538e82856148ea565b6001600160a01b039290921660409190910152919050565b604081525f6153b86040830185614cb8565b8281036020840152611fd98185614cb8565b5f602082840312156153da575f5ffd5b81516001600160401b038111156153ef575f5ffd5b8201601f810184136153ff575f5ffd5b805161540d614a4882614a09565b8082825260208201915060208360051b85010192508683111561542e575f5ffd5b602084015b838110156154cb5780516001600160401b03811115615450575f5ffd5b8501603f81018913615460575f5ffd5b6020810151615471614a4882614a09565b808282526020820191506020808460051b8601010192508b831115615494575f5ffd5b6040840193505b828410156154b657835182526020938401939091019061549b565b86525050602093840193919091019050615433565b509695505050505050565b634e487b7160e01b5f52601160045260245ffd5b8181038181111561099a5761099a6154d6565b6001600160401b03828116828216039081111561099a5761099a6154d6565b5f81600f0b60016001607f1b03198103615538576155386154d6565b5f0392915050565b600f81810b9083900b0160016001607f1b03811360016001607f1b03198212171561099a5761099a6154d6565b6001600160a01b038616815260c0810161558a60208301876148ea565b6001600160a01b039490941660608201526001600160401b0392909216608083015263ffffffff1660a09091015292915050565b5f5f8335601e198436030181126155d3575f5ffd5b8301803591506001600160401b038211156155ec575f5ffd5b60200191503681900382131561191d575f5ffd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b6001600160a01b03881681525f60c08201615646602084018a6148ea565b60c060608401528690528660e083015f5b8881101561568757823561566a8161474e565b6001600160a01b0316825260209283019290910190600101615657565b50838103608085015261569a8188614b22565b91505082810360a08401526156b0818587615600565b9a9950505050505050505050565b5f602082840312156156ce575f5ffd5b813561ffff81168114610776575f5ffd5b63ffffffff818116838216019081111561099a5761099a6154d6565b8183526020830192505f815f5b848110156149435763ffffffff61571e836147ce565b1686526020958601959190910190600101615708565b6001600160a01b038581168252841660208201526060604082018190525f9061429790830184866156fb565b602081525f6131d4602083018486615600565b6001600160a01b038781168252861660208201526080604082018190525f9061579f90830186886156fb565b82810360608401526157b2818587615600565b9998505050505050505050565b5f602082840312156157cf575f5ffd5b81516107768161474e565b80516020808301519190811015614f49575f1960209190910360031b1b16919050565b5f6001820161580e5761580e6154d6565b5060010190565b5f81615823576158236154d6565b505f190190565b8082018281125f831280158216821582161715615849576158496154d6565b505092915050565b600f82810b9082900b0360016001607f1b0319811260016001607f1b038213171561099a5761099a6154d6565b634e487b7160e01b5f52601260045260245ffd5b634e487b7160e01b5f52602160045260245ffd5b8082018082111561099a5761099a6154d6565b634e487b7160e01b5f52603160045260245ffd5b5f826158e757634e487b7160e01b5f52601260045260245ffd5b50049056fea2646970667358221220271c916bfed2c4e6b796b3f2f4d779226127ece2a6bce7508308b4f4531fe0f164736f6c634300081b0033",
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
