// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// IAllocationManagerTypesCreateSetParamsV2 is an auto generated low-level Go binding around an user-defined struct.
type IAllocationManagerTypesCreateSetParamsV2 struct {
	OperatorSetId uint32
	Strategies    []common.Address
	Slasher       common.Address
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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_delegation\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"},{\"name\":\"_eigenStrategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"_permissionController\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"},{\"name\":\"_DEALLOCATION_DELAY\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_ALLOCATION_CONFIGURATION_DELAY\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_version\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ALLOCATION_CONFIGURATION_DELAY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEALLOCATION_DELAY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addStrategiesToOperatorSet\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"clearDeallocationQueue\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"numToClear\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createOperatorSets\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"params\",\"type\":\"tuple[]\",\"internalType\":\"structIAllocationManagerTypes.CreateSetParams[]\",\"components\":[{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createOperatorSets\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"params\",\"type\":\"tuple[]\",\"internalType\":\"structIAllocationManagerTypes.CreateSetParamsV2[]\",\"components\":[{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"slasher\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createRedistributingOperatorSets\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"params\",\"type\":\"tuple[]\",\"internalType\":\"structIAllocationManagerTypes.CreateSetParams[]\",\"components\":[{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}]},{\"name\":\"redistributionRecipients\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createRedistributingOperatorSets\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"params\",\"type\":\"tuple[]\",\"internalType\":\"structIAllocationManagerTypes.CreateSetParamsV2[]\",\"components\":[{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"slasher\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"redistributionRecipients\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregisterFromOperatorSets\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIAllocationManagerTypes.DeregisterParams\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eigenStrategy\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAVSRegistrar\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAVSRegistrar\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocatableMagnitude\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocatedSets\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocatedStake\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[][]\",\"internalType\":\"uint256[][]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocatedStrategies\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocation\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIAllocationManagerTypes.Allocation\",\"components\":[{\"name\":\"currentMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"pendingDiff\",\"type\":\"int128\",\"internalType\":\"int128\"},{\"name\":\"effectBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocationDelay\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocations\",\"inputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIAllocationManagerTypes.Allocation[]\",\"components\":[{\"name\":\"currentMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"pendingDiff\",\"type\":\"int128\",\"internalType\":\"int128\"},{\"name\":\"effectBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getEncumberedMagnitude\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxMagnitude\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxMagnitudes\",\"inputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxMagnitudes\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxMagnitudesAtBlock\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMemberCount\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMembers\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMinimumSlashableStake\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"futureBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"slashableStake\",\"type\":\"uint256[][]\",\"internalType\":\"uint256[][]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorSetCount\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPendingSlasher\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRedistributionRecipient\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRegisteredSets\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSlashCount\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSlasher\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStrategiesInOperatorSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStrategyAllocations\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIAllocationManagerTypes.Allocation[]\",\"components\":[{\"name\":\"currentMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"pendingDiff\",\"type\":\"int128\",\"internalType\":\"int128\"},{\"name\":\"effectBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isMemberOfOperatorSet\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperatorRedistributable\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperatorSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperatorSlashable\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRedistributingOperatorSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"migrateSlashers\",\"inputs\":[{\"name\":\"operatorSets\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"modifyAllocations\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"params\",\"type\":\"tuple[]\",\"internalType\":\"structIAllocationManagerTypes.AllocateParams[]\",\"components\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"newMagnitudes\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"permissionController\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerForOperatorSets\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIAllocationManagerTypes.RegisterParams\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeStrategiesFromOperatorSet\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAVSRegistrar\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"registrar\",\"type\":\"address\",\"internalType\":\"contractIAVSRegistrar\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAllocationDelay\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delay\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slashOperator\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIAllocationManagerTypes.SlashingParams\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"wadsToSlash\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAVSMetadataURI\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateSlasher\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slasher\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AVSMetadataURIUpdated\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"metadataURI\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AVSRegistrarSet\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"registrar\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIAVSRegistrar\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllocationDelaySet\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"delay\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"effectBlock\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllocationUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"magnitude\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"effectBlock\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EncumberedMagnitudeUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"encumberedMagnitude\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MaxMagnitudeUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"maxMagnitude\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorAddedToOperatorSet\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRemovedFromOperatorSet\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSetCreated\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSlashed\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategies\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"contractIStrategy[]\"},{\"name\":\"wadSlashed\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"description\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RedistributionAddressSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"redistributionRecipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SlasherMigrated\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slasher\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SlasherUpdated\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slasher\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"effectBlock\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyAddedToOperatorSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyRemovedFromOperatorSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyMemberOfSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CurrentlyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Empty\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputArrayLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientMagnitude\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAVSRegistrar\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCaller\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNewPausedStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperatorSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPermissions\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRedistributionRecipient\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSnapshotOrdering\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidStrategy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidWadToSlash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ModificationAlreadyPending\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonexistentAVSMetadata\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotMemberOfSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyPauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnpauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorNotSlashable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorSetAlreadyMigrated\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OutOfBounds\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SameMagnitude\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SlasherNotSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategiesMustBeInAscendingOrder\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategyAlreadyInOperatorSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategyNotInOperatorSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"UninitializedAllocationDelay\",\"inputs\":[]}]",
	Bin: "0x610160604052348015610010575f5ffd5b50604051616dc3380380616dc383398101604081905261002f916101ef565b8084888886868a6001600160a01b03811661005d576040516339b190bb60e11b815260040160405180910390fd5b6001600160a01b0390811660805293841660a05291831660c05263ffffffff90811660e05216610100521661012052610095816100ae565b61014052506100a26100f4565b50505050505050610367565b5f5f829050601f815111156100e1578260405163305a27a960e01b81526004016100d8919061030c565b60405180910390fd5b80516100ec82610341565b179392505050565b5f54610100900460ff161561015b5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b60648201526084016100d8565b5f5460ff908116146101aa575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146101c0575f5ffd5b50565b805163ffffffff811681146101d6575f5ffd5b919050565b634e487b7160e01b5f52604160045260245ffd5b5f5f5f5f5f5f5f60e0888a031215610205575f5ffd5b8751610210816101ac565b6020890151909750610221816101ac565b6040890151909650610232816101ac565b6060890151909550610243816101ac565b9350610251608089016101c3565b925061025f60a089016101c3565b60c08901519092506001600160401b0381111561027a575f5ffd5b8801601f81018a1361028a575f5ffd5b80516001600160401b038111156102a3576102a36101db565b604051601f8201601f19908116603f011681016001600160401b03811182821017156102d1576102d16101db565b6040528181528282016020018c10156102e8575f5ffd5b8160208401602083015e5f6020838301015280935050505092959891949750929550565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b80516020808301519190811015610361575f198160200360031b1b821691505b50919050565b60805160a05160c05160e0516101005161012051610140516169a961041a5f395f6113c601525f818161055e015281816121ba015261326d01525f81816106e4015281816142ff015261450e01525f8181610454015281816118790152611f4401525f81816108af015261415401525f81816108e9015281816113fa01528181611465015281816124b70152818161349b0152613d3801525f818161070b015281816131390152614ae001526169a95ff3fe608060405234801561000f575f5ffd5b506004361061037d575f3560e01c80636cfb4481116101d4578063b66bd98911610109578063d7794857116100a9578063f231bd0811610079578063f231bd081461090b578063f605ce081461091e578063fabc1cbc14610931578063fe4b84df14610944575f5ffd5b8063d779485714610873578063db4df761146108aa578063dc2af692146108d1578063df5cf723146108e4575f5ffd5b8063c221d8ae116100e4578063c221d8ae14610827578063d1a83f541461083a578063d3d96ff41461084d578063d4a3fcce14610860575f5ffd5b8063b66bd989146107d2578063b9fbaed1146107e5578063ba1a84e514610814575f5ffd5b806394d7d00c11610174578063a9333ec81161014f578063a9333ec814610786578063a982182114610799578063adc2e3d9146107ac578063b2447af7146107bf575f5ffd5b806394d7d00c1461074d578063952899ee14610760578063957dc50b14610773575f5ffd5b806379ae50cd116101af57806379ae50cd146106cc5780637bc1ef61146106df578063886f1195146107065780638ce648541461072d575f5ffd5b80636cfb44811461067b5780636e3492b5146106a65780636e875dba146106b9575f5ffd5b806340120dab116102b5578063547afb87116102555780635ac86ab7116102255780635ac86ab71461062a5780635c975abb1461064d578063670d3ba2146106555780636c9d7c5814610668575f5ffd5b8063547afb87146105e757806354fd4d50146105fa57806356c483e61461060f578063595c6a6714610622575f5ffd5b80634a10ffe5116102905780634a10ffe5146105805780634b5046ef146105a05780634cfd2939146105b357806350feea20146105d4575f5ffd5b806340120dab146105185780634177a87c146105395780634657e26a14610559575f5ffd5b80632981eb7711610320578063304c10cd116102fb578063304c10cd146104be57806332a879e4146104d157806336352057146104e45780633dff8e7d14610505575f5ffd5b80632981eb771461044f5780632b453a9a1461048b5780632bab2c4a146104ab575f5ffd5b8063136439dd1161035b578063136439dd146103f457806315fe502814610409578063260dc75814610429578063261f84e01461043c575f5ffd5b80630f3df50e1461038157806310e1b9b8146103b15780631352c3e6146103d1575b5f5ffd5b61039461038f3660046154d3565b610957565b6040516001600160a01b0390911681526020015b60405180910390f35b6103c46103bf3660046154ed565b610998565b6040516103a89190615534565b6103e46103df366004615567565b6109d1565b60405190151581526020016103a8565b61040761040236600461559b565b610a4c565b005b61041c6104173660046155b2565b610a86565b6040516103a89190615630565b6103e46104373660046154d3565b610b9d565b61040761044a366004615682565b610bce565b6104767f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff90911681526020016103a8565b61049e610499366004615767565b610c0a565b6040516103a8919061580a565b61049e6104b936600461586d565b610c20565b6103946104cc3660046155b2565b610cbf565b6104076104df3660046158f1565b610cee565b6104f76104f2366004615971565b610d33565b6040516103a89291906159c3565b610407610513366004615ab9565b610e88565b61052b610526366004615afb565b610f21565b6040516103a8929190615b88565b61054c6105473660046154d3565b61109c565b6040516103a89190615be5565b6103947f000000000000000000000000000000000000000000000000000000000000000081565b61059361058e366004615bf7565b6110c0565b6040516103a89190615c3a565b6104076105ae3660046158f1565b611168565b6105c66105c13660046154d3565b611204565b6040519081526020016103a8565b6104076105e2366004615c85565b611226565b6105936105f5366004615ce3565b611317565b6106026113bf565b6040516103a89190615d25565b61040761061d366004615d5a565b6113ef565b6104076114fb565b6103e4610638366004615d84565b606654600160ff9092169190911b9081161490565b6066546105c6565b6103e4610663366004615567565b61150f565b610407610676366004615da4565b61153b565b61068e610689366004615afb565b6115ee565b6040516001600160401b0390911681526020016103a8565b6104076106b4366004615de5565b611603565b61054c6106c73660046154d3565b6119be565b61041c6106da3660046155b2565b6119cf565b6104767f000000000000000000000000000000000000000000000000000000000000000081565b6103947f000000000000000000000000000000000000000000000000000000000000000081565b61074061073b366004615e16565b611aa9565b6040516103a89190615e59565b61059361075b366004615e6b565b611b65565b61040761076e366004615ec6565b611c51565b61040761078136600461606f565b6120f9565b61068e610794366004615afb565b612385565b6104076107a73660046160ff565b6123b4565b6104076107ba36600461617d565b612466565b6105c66107cd3660046154d3565b6127af565b6104076107e0366004615c85565b6127d1565b6107f86107f33660046155b2565b61292b565b60408051921515835263ffffffff9091166020830152016103a8565b6105c66108223660046155b2565b6129c5565b61054c610835366004615567565b6129e5565b6104076108483660046161bf565b612a0e565b61040761085b366004615afb565b612b44565b61039461086e3660046154d3565b612c71565b6108866108813660046154d3565b612cf8565b604080516001600160a01b03909316835263ffffffff9091166020830152016103a8565b6103947f000000000000000000000000000000000000000000000000000000000000000081565b6103e46108df3660046155b2565b612d79565b6103947f000000000000000000000000000000000000000000000000000000000000000081565b6103e46109193660046154d3565b612dac565b61068e61092c366004615afb565b612dcb565b61040761093f36600461559b565b612dd7565b61040761095236600461559b565b612e44565b5f5f60a65f61096585612f55565b815260208101919091526040015f20546001600160a01b03169050801561098c5780610991565b620e16e45b9392505050565b604080516060810182525f80825260208201819052918101829052906109c7856109c186612f55565b85612fb8565b9695505050505050565b6001600160a01b0382165f908152609e602052604081208190816109f485612f55565b815260208082019290925260409081015f2081518083019092525460ff8116151580835261010090910463ffffffff1692820192909252915080610a425750806020015163ffffffff164311155b9150505b92915050565b610a54613124565b6066548181168114610a795760405163c61dca5d60e01b815260040160405180910390fd5b610a82826131c7565b5050565b6001600160a01b0381165f908152609d6020526040812060609190610aaa90613204565b90505f816001600160401b03811115610ac557610ac56153e3565b604051908082528060200260200182016040528015610b0957816020015b604080518082019091525f8082526020820152815260200190600190039081610ae35790505b5090505f5b82811015610b95576001600160a01b0385165f908152609d60205260409020610b7090610b3b908361320d565b604080518082019091525f80825260208201525060408051808201909152606082901c815263ffffffff909116602082015290565b828281518110610b8257610b8261621e565b6020908102919091010152600101610b0e565b509392505050565b60208082015182516001600160a01b03165f908152609890925260408220610a469163ffffffff9081169061321816565b82610bd88161322f565b610bf55760405163932d94f760e01b815260040160405180910390fd5b610c04846105138585886132d9565b50505050565b6060610c1884848443613447565b949350505050565b6060610c2e85858585613447565b90505f5b8451811015610cb657610c5e858281518110610c5057610c5061621e565b6020026020010151876109d1565b610cae575f5b8451811015610cac575f838381518110610c8057610c8061621e565b60200260200101518281518110610c9957610c9961621e565b6020908102919091010152600101610c64565b505b600101610c32565b50949350505050565b6001600160a01b038082165f908152609760205260408120549091168015610ce75780610991565b5090919050565b84610cf88161322f565b610d155760405163932d94f760e01b815260040160405180910390fd5b610d2b86610d2487878a6132d9565b8585612a0e565b505050505050565b5f60606001610d4181613734565b5f6040518060400160405280876001600160a01b03168152602001866020016020810190610d6f9190616232565b63ffffffff1690529050610d86606086018661624b565b9050610d95604087018761624b565b905014610db5576040516343714afd60e01b815260040160405180910390fd5b60208082015182516001600160a01b03165f90815260989092526040909120610de79163ffffffff9081169061321816565b610e0457604051631fb1705560e21b815260040160405180910390fd5b610e1a610e1460208701876155b2565b826109d1565b610e375760405163ebbff49760e01b815260040160405180910390fd5b610e4081612c71565b6001600160a01b0316336001600160a01b031614610e71576040516348f5c3ed60e01b815260040160405180910390fd5b610e7b8582613762565b9350935050509250929050565b81610e928161322f565b610eaf5760405163932d94f760e01b815260040160405180910390fd5b6001600160a01b0383165f90815260a4602052604090205460ff16610ee7576040516348f7dbb960e01b815260040160405180910390fd5b5f5b8251811015610c0457610f1984848381518110610f0857610f0861621e565b6020026020010151620e16e4613ea2565b600101610ee9565b6001600160a01b0382165f908152609d602052604081206060918291610f4690613204565b90505f816001600160401b03811115610f6157610f616153e3565b604051908082528060200260200182016040528015610fa557816020015b604080518082019091525f8082526020820152815260200190600190039081610f7f5790505b5090505f826001600160401b03811115610fc157610fc16153e3565b60405190808252806020026020018201604052801561100a57816020015b604080516060810182525f80825260208083018290529282015282525f19909201910181610fdf5790505b5090505f5b8381101561108d576001600160a01b0388165f908152609d6020526040812061103c90610b3b908461320d565b9050808483815181106110515761105161621e565b602002602001018190525061106789828a610998565b8383815181106110795761107961621e565b60209081029190910101525060010161100f565b509093509150505b9250929050565b60605f61099160995f6110ae86612f55565b81526020019081526020015f20614014565b60605f83516001600160401b038111156110dc576110dc6153e3565b604051908082528060200260200182016040528015611105578160200160208202803683370190505b5090505f5b8451811015610b95576111368582815181106111285761112861621e565b602002602001015185612385565b8282815181106111485761114861621e565b6001600160401b039092166020928302919091019091015260010161110a565b5f61117281613734565b838214611192576040516343714afd60e01b815260040160405180910390fd5b5f5b848110156111fb576111f3878787848181106111b2576111b261621e565b90506020020160208101906111c791906155b2565b8686858181106111d9576111d961621e565b90506020020160208101906111ee9190616290565b614020565b600101611194565b50505050505050565b5f60a55f61121184612f55565b81526020019081526020015f20549050919050565b836112308161322f565b61124d5760405163932d94f760e01b815260040160405180910390fd5b6040805180820182526001600160a01b03871680825263ffffffff80881660208085018290525f9384526098905293909120919261128c929161321816565b6112a957604051631fb1705560e21b815260040160405180910390fd5b5f5b838110156111fb5761130f828686848181106112c9576112c961621e565b90506020020160208101906112de91906155b2565b61130a60405180604001604052808c6001600160a01b031681526020018b63ffffffff16815250612dac565b614124565b6001016112ab565b60605f82516001600160401b03811115611333576113336153e3565b60405190808252806020026020018201604052801561135c578160200160208202803683370190505b5090505f5b8351811015610b955761138d858583815181106113805761138061621e565b6020026020010151612385565b82828151811061139f5761139f61621e565b6001600160401b0390921660209283029190910190910152600101611361565b60606113ea7f0000000000000000000000000000000000000000000000000000000000000000614204565b905090565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614806114eb576114298361322f565b611446576040516348f5c3ed60e01b815260040160405180910390fd5b6040516336b87bd760e11b81526001600160a01b0384811660048301527f00000000000000000000000000000000000000000000000000000000000000001690636d70f7ae90602401602060405180830381865afa1580156114aa573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906114ce91906162b1565b6114eb5760405163ccea9e6f60e01b815260040160405180910390fd5b6114f6838383614241565b505050565b611503613124565b61150d5f196131c7565b565b5f61099183609a5f61152086612f55565b81526020019081526020015f2061441190919063ffffffff16565b81516115468161322f565b6115635760405163932d94f760e01b815260040160405180910390fd5b60208084015184516001600160a01b03165f908152609890925260409091206115959163ffffffff9081169061321816565b6115b257604051631fb1705560e21b815260040160405180910390fd5b5f6115bc84612c71565b6001600160a01b0316036115e35760405163255b0f4160e01b815260040160405180910390fd5b6114f683835f614432565b5f5f6115fa84846145ee565b95945050505050565b600261160e81613734565b61162361161e60208401846155b2565b61322f565b8061163c575061163c61161e60408401602085016155b2565b611659576040516348f5c3ed60e01b815260040160405180910390fd5b5f5b611668604084018461624b565b905081101561192a575f604051806040016040528085602001602081019061169091906155b2565b6001600160a01b031681526020016116ab604087018761624b565b858181106116bb576116bb61621e565b90506020020160208101906116d09190616232565b63ffffffff16815250905061171d816020015163ffffffff1660985f8760200160208101906116ff91906155b2565b6001600160a01b0316815260208101919091526040015f2090613218565b61173a57604051631fb1705560e21b815260040160405180910390fd5b609e5f61174a60208701876155b2565b6001600160a01b03166001600160a01b031681526020019081526020015f205f61177383612f55565b815260208101919091526040015f205460ff166117a3576040516325131d4f60e01b815260040160405180910390fd5b6117dd6117af82612f55565b609c5f6117bf60208901896155b2565b6001600160a01b0316815260208101919091526040015f209061475d565b506118156117ee60208601866155b2565b609a5f6117fa85612f55565b81526020019081526020015f2061476890919063ffffffff16565b5061182360208501856155b2565b6001600160a01b03167fad34c3070be1dffbcaa499d000ba2b8d9848aefcac3059df245dd95c4ece14fe8260405161185b91906162d0565b60405180910390a2604080518082019091525f81526020810161189e7f0000000000000000000000000000000000000000000000000000000000000000436162f2565b63ffffffff169052609e5f6118b660208801886155b2565b6001600160a01b03166001600160a01b031681526020019081526020015f205f6118df84612f55565b81526020808201929092526040015f2082518154939092015163ffffffff166101000264ffffffff00199215159290921664ffffffffff19909316929092171790555060010161165b565b5061193e6104cc60408401602085016155b2565b6001600160a01b031663303ca95661195960208501856155b2565b61196960408601602087016155b2565b611976604087018761624b565b6040518563ffffffff1660e01b81526004016119959493929190616347565b5f604051808303815f87803b1580156119ac575f5ffd5b505af1158015610d2b573d5f5f3e3d5ffd5b6060610a46609a5f6110ae85612f55565b6001600160a01b0381165f908152609c60205260408120606091906119f390613204565b90505f816001600160401b03811115611a0e57611a0e6153e3565b604051908082528060200260200182016040528015611a5257816020015b604080518082019091525f8082526020820152815260200190600190039081611a2c5790505b5090505f5b82811015610b95576001600160a01b0385165f908152609c60205260409020611a8490610b3b908361320d565b828281518110611a9657611a9661621e565b6020908102919091010152600101611a57565b60605f84516001600160401b03811115611ac557611ac56153e3565b604051908082528060200260200182016040528015611b0e57816020015b604080516060810182525f80825260208083018290529282015282525f19909201910181611ae35790505b5090505f5b8551811015610cb657611b40868281518110611b3157611b3161621e565b60200260200101518686610998565b828281518110611b5257611b5261621e565b6020908102919091010152600101611b13565b60605f83516001600160401b03811115611b8157611b816153e3565b604051908082528060200260200182016040528015611baa578160200160208202803683370190505b5090505f5b8451811015610cb6576001600160a01b0386165f90815260a1602052604081208651611c1f92879291899086908110611bea57611bea61621e565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020015f2061477c90919063ffffffff16565b828281518110611c3157611c3161621e565b6001600160401b0390921660209283029190910190910152600101611baf565b5f611c5b81613734565b611c648361322f565b611c81576040516348f5c3ed60e01b815260040160405180910390fd5b5f5f5f611c8d8661292b565b9150915081611caf5760405163fa55fc8160e01b815260040160405180910390fd5b91505f90505b83518110156120f257838181518110611cd057611cd061621e565b60200260200101516040015151848281518110611cef57611cef61621e565b6020026020010151602001515114611d1a576040516343714afd60e01b815260040160405180910390fd5b5f848281518110611d2d57611d2d61621e565b602090810291909101810151518082015181516001600160a01b03165f90815260989093526040909220909250611d6d9163ffffffff9081169061321816565b611d8a57604051631fb1705560e21b815260040160405180910390fd5b5f611d9587836109d1565b90505f5b868481518110611dab57611dab61621e565b602002602001015160200151518110156120e7575f878581518110611dd257611dd261621e565b6020026020010151602001518281518110611def57611def61621e565b60200260200101519050611e06898261ffff614020565b5f5f611e158b6109c188612f55565b91509150806040015163ffffffff165f14611e4357604051630d8fcbe360e41b815260040160405180910390fd5b5f611e5087858489614790565b9050611e95825f01518c8a81518110611e6b57611e6b61621e565b6020026020010151604001518781518110611e8857611e8861621e565b60200260200101516147c6565b600f0b602083018190525f03611ebe57604051634606179360e11b815260040160405180910390fd5b5f8260200151600f0b1215612002578015611f8457611f3f611edf88612f55565b6001600160a01b03808f165f90815260a360209081526040808320938a16835292905220908154600160801b90819004600f0b5f818152600180860160205260409091209390935583546001600160801b03908116939091011602179055565b611f697f0000000000000000000000000000000000000000000000000000000000000000436162f2565b611f749060016162f2565b63ffffffff16604083015261206f565b611f96836020015183602001516147dd565b6001600160401b031660208401528a518b9089908110611fb857611fb861621e565b6020026020010151604001518581518110611fd557611fd561621e565b6020908102919091018101516001600160401b031683525f9083015263ffffffff4316604083015261206f565b5f8260200151600f0b131561206f57612023836020015183602001516147dd565b6001600160401b03908116602085018190528451909116101561205957604051636c9be0bf60e01b815260040160405180910390fd5b61206389436162f2565b63ffffffff1660408301525b6120848c61207c89612f55565b8686866147fc565b7f1487af5418c47ee5ea45ef4a93398668120890774a9e13487e61e9dc3baf76dd8c88866120b9865f015187602001516147dd565b86604001516040516120cf959493929190616373565b60405180910390a1505060019092019150611d999050565b505050600101611cb5565b5050505050565b5f5b8151811015610a82576121788282815181106121195761211961621e565b60200260200101516020015163ffffffff1660985f8585815181106121405761214061621e565b60200260200101515f01516001600160a01b03166001600160a01b031681526020019081526020015f2061321890919063ffffffff16565b1561237d575f6001600160a01b03166121a983838151811061219c5761219c61621e565b6020026020010151612c71565b6001600160a01b03160361237d575f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663fddbdefd8484815181106121f9576121f961621e565b6020908102919091010151516040516001600160e01b031960e084901b1681526001600160a01b039091166004820152306024820152633635205760e01b60448201526064015f604051808303815f875af115801561225a573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261228191908101906163c4565b90505f81515f14806122bd57505f6001600160a01b0316825f815181106122aa576122aa61621e565b60200260200101516001600160a01b0316145b156122e6578383815181106122d4576122d461621e565b60200260200101515f01519050612303565b815f815181106122f8576122f861621e565b602002602001015190505b6123288484815181106123185761231861621e565b6020026020010151826001614432565b7ff0c8fc7d71f647bd3a88ac369112517f6a4b8038e71913f2d20f71f877dfc72584848151811061235b5761235b61621e565b602002602001015182604051612372929190616453565b60405180910390a150505b6001016120fb565b6001600160a01b038083165f90815260a160209081526040808320938516835292905290812061099190614a34565b826123be8161322f565b6123db5760405163932d94f760e01b815260040160405180910390fd5b6001600160a01b0384165f90815260a4602052604090205460ff1661241d576001600160a01b0384165f90815260a460205260409020805460ff191660011790555b836001600160a01b03167fa89c1dc243d8908a96dd84944bcc97d6bc6ac00dd78e20621576be6a3c94371384846040516124589291906164a1565b60405180910390a250505050565b600261247181613734565b8261247b8161322f565b6124985760405163932d94f760e01b815260040160405180910390fd5b6040516336b87bd760e11b81526001600160a01b0385811660048301527f00000000000000000000000000000000000000000000000000000000000000001690636d70f7ae90602401602060405180830381865afa1580156124fc573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061252091906162b1565b61253d5760405163ccea9e6f60e01b815260040160405180910390fd5b5f5b61254c602085018561624b565b905081101561271457604080518082019091525f908061256f60208801886155b2565b6001600160a01b0316815260200186806020019061258d919061624b565b8581811061259d5761259d61621e565b90506020020160208101906125b29190616232565b63ffffffff90811690915260208083015183516001600160a01b03165f908152609890925260409091209293506125ee92919081169061321816565b61260b57604051631fb1705560e21b815260040160405180910390fd5b61261586826109d1565b1561263357604051636c6c6e2760e11b815260040160405180910390fd5b61265c61263f82612f55565b6001600160a01b0388165f908152609c6020526040902090614a47565b5061268886609a5f61266d85612f55565b81526020019081526020015f20614a5290919063ffffffff16565b50856001600160a01b03167f43232edf9071753d2321e5fa7e018363ee248e5f2142e6c08edd3265bfb4895e826040516126c291906162d0565b60405180910390a26001600160a01b0386165f908152609e602052604081206001916126ed84612f55565b815260208101919091526040015f20805460ff19169115159190911790555060010161253f565b506127256104cc60208501856155b2565b6001600160a01b031663c63fd5028561274160208701876155b2565b61274e602088018861624b565b61275b60408a018a6164b4565b6040518763ffffffff1660e01b815260040161277c969594939291906164f6565b5f604051808303815f87803b158015612793575f5ffd5b505af11580156127a5573d5f5f3e3d5ffd5b5050505050505050565b5f610a46609a5f6127bf85612f55565b81526020019081526020015f20613204565b836127db8161322f565b6127f85760405163932d94f760e01b815260040160405180910390fd5b6040805180820182526001600160a01b03871680825263ffffffff80881660208085018290525f93845260989052939091209192612837929161321816565b61285457604051631fb1705560e21b815260040160405180910390fd5b5f61285e82612f55565b90505f5b848110156127a5576128a786868381811061287f5761287f61621e565b905060200201602081019061289491906155b2565b5f84815260996020526040902090614768565b6128c4576040516331bc342760e11b815260040160405180910390fd5b7f7b4b073d80dcac55a11177d8459ad9f664ceeb91f71f27167bb14f8152a7eeee838787848181106128f8576128f861621e565b905060200201602081019061290d91906155b2565b60405161291b929190616453565b60405180910390a1600101612862565b6001600160a01b0381165f908152609b602090815260408083208151608081018352905463ffffffff80821680845260ff600160201b8404161515958401869052650100000000008304821694840194909452600160481b9091041660608201819052849391929190158015906129ac5750826060015163ffffffff164310155b156129bb575050604081015160015b9590945092505050565b6001600160a01b0381165f908152609860205260408120610a4690613204565b6001600160a01b0382165f908152609f6020526040812060609190610a4290826110ae86612f55565b83612a188161322f565b612a355760405163932d94f760e01b815260040160405180910390fd5b83518214612a56576040516343714afd60e01b815260040160405180910390fd5b6001600160a01b0385165f90815260a4602052604090205460ff16612a8e576040516348f7dbb960e01b815260040160405180910390fd5b5f5b8451811015610d2b575f848483818110612aac57612aac61621e565b9050602002016020810190612ac191906155b2565b90506001600160a01b038116612aea576040516339b190bb60e11b815260040160405180910390fd5b620e16e3196001600160a01b03821601612b17576040516364be1a3f60e11b815260040160405180910390fd5b612b3b87878481518110612b2d57612b2d61621e565b602002602001015183613ea2565b50600101612a90565b81612b4e8161322f565b612b6b5760405163932d94f760e01b815260040160405180910390fd5b60405163b526578760e01b81526001600160a01b03848116600483015283169063b526578790602401602060405180830381865afa158015612baf573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612bd391906162b1565b612bf057604051631d0b13c160e31b815260040160405180910390fd5b6001600160a01b038381165f90815260976020526040902080546001600160a01b0319169184169190911790557f2ae945c40c44dc0ec263f95609c3fdc6952e0aefa22d6374e44f2c997acedf8583612c4881610cbf565b604080516001600160a01b039384168152929091166020830152015b60405180910390a1505050565b5f5f60a75f612c7f85612f55565b815260208082019290925260409081015f20815160608101835281546001600160a01b0390811680835260019093015490811694820194909452600160a01b90930463ffffffff1691830182905291925015801590612ce85750816040015163ffffffff164310155b1561099157506020015192915050565b5f5f5f5f5f60a75f612d0988612f55565b815260208082019290925260409081015f20815160608101835281546001600160a01b03908116825260019092015491821693810193909352600160a01b900463ffffffff16908201819052909150431015612d6e5780602001519250806040015191505b509094909350915050565b5f5f612d84836119cf565b90505f612d9084610a86565b9050612d9c8483614a66565b80610c185750610c188482614a66565b5f620e16e4612dba83610957565b6001600160a01b0316141592915050565b5f5f610cb684846145ee565b612ddf614ade565b60665480198219811614612e065760405163c61dca5d60e01b815260040160405180910390fd5b606682905560405182815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c9060200160405180910390a25050565b5f54610100900460ff1615808015612e6257505f54600160ff909116105b80612e7b5750303b158015612e7b57505f5460ff166001145b612ee35760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b5f805460ff191660011790558015612f04575f805461ff0019166101001790555b612f0d826131c7565b8015610a82575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15050565b5f815f0151826020015163ffffffff16604051602001612fa092919060609290921b6bffffffffffffffffffffffff1916825260a01b6001600160a01b031916601482015260200190565b604051602081830303815290604052610a4690616542565b6040805180820182525f80825260208083018290528351606081018552828152808201839052808501839052845180860186526001600160a01b03898116855260a184528685209088168552909252938220929392819061301890614a34565b6001600160401b0390811682526001600160a01b038981165f81815260a260209081526040808320948c168084529482528083205486169682019690965291815260a082528481208b8252825284812092815291815290839020835160608101855290549283168152600160401b8304600f0b91810191909152600160c01b90910463ffffffff169181018290529192504310156130ba57909250905061311c565b6130cb815f015182602001516147dd565b6001600160401b0316815260208101515f600f9190910b1215613109576130fa826020015182602001516147dd565b6001600160401b031660208301525b5f60408201819052602082015290925090505b935093915050565b60405163237dfb4760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906346fbf68e90602401602060405180830381865afa158015613186573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906131aa91906162b1565b61150d57604051631d77d47760e21b815260040160405180910390fd5b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a250565b5f610a46825490565b5f6109918383614b8f565b5f8181526001830160205260408120541515610991565b604051631beb2b9760e31b81526001600160a01b0382811660048301523360248301523060448301525f80356001600160e01b0319166064840152917f00000000000000000000000000000000000000000000000000000000000000009091169063df595cb8906084016020604051808303815f875af11580156132b5573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a4691906162b1565b60605f836001600160401b038111156132f4576132f46153e3565b60405190808252806020026020018201604052801561335657816020015b61334360405180606001604052805f63ffffffff168152602001606081526020015f6001600160a01b031681525090565b8152602001906001900390816133125790505b5090505f5b84811015610cb65760405180606001604052808787848181106133805761338061621e565b90506020028101906133929190616565565b6133a0906020810190616232565b63ffffffff1681526020018787848181106133bd576133bd61621e565b90506020028101906133cf9190616565565b6133dd90602081019061624b565b808060200260200160405190810160405280939291908181526020018383602002808284375f920191909152505050908252506001600160a01b03861660209091015282518390839081106134345761343461621e565b602090810291909101015260010161335b565b606083516001600160401b03811115613462576134626153e3565b60405190808252806020026020018201604052801561349557816020015b60608152602001906001900390816134805790505b5090505f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f0e0e67686866040518363ffffffff1660e01b81526004016134e7929190616583565b5f60405180830381865afa158015613501573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261352891908101906165a7565b90505f5b855181101561372a575f8682815181106135485761354861621e565b6020026020010151905085516001600160401b0381111561356b5761356b6153e3565b604051908082528060200260200182016040528015613594578160200160208202803683370190505b508483815181106135a7576135a761621e565b60209081029190910101525f5b8651811015613720575f8782815181106135d0576135d061621e565b6020908102919091018101516001600160a01b038086165f90815260a184526040808220928416825291909352822090925061360b90614a34565b9050806001600160401b03165f03613624575050613718565b5f613630858d85610998565b90508863ffffffff16816040015163ffffffff161115801561365857505f8160200151600f0b125b1561367a5761366e815f015182602001516147dd565b6001600160401b031681525b80515f90613695906001600160401b03908116908516614bb5565b90506136dc818989815181106136ad576136ad61621e565b602002602001015187815181106136c6576136c661621e565b6020026020010151614bc990919063ffffffff16565b8988815181106136ee576136ee61621e565b602002602001015186815181106137075761370761621e565b602002602001018181525050505050505b6001016135b4565b505060010161352c565b5050949350505050565b606654600160ff83161b9081160361375f5760405163840a48d560e01b815260040160405180910390fd5b50565b5f606081613773604086018661624b565b90506001600160401b0381111561378c5761378c6153e3565b6040519080825280602002602001820160405280156137b5578160200160208202803683370190505b5090506137c5604086018661624b565b90506001600160401b038111156137de576137de6153e3565b604051908082528060200260200182016040528015613807578160200160208202803683370190505b50915060a55f61381686612f55565b81526020019081526020015f205f815461382f906166b3565b918290555092505f5b613845604087018761624b565b9050811015613e34578015806138d85750613863604087018761624b565b61386e6001846166cb565b81811061387d5761387d61621e565b905060200201602081019061389291906155b2565b6001600160a01b03166138a8604088018861624b565b838181106138b8576138b861621e565b90506020020160208101906138cd91906155b2565b6001600160a01b0316115b6138f557604051639f1c805360e01b815260040160405180910390fd5b613902606087018761624b565b828181106139125761391261621e565b905060200201355f1080156139525750670de0b6b3a7640000613938606088018861624b565b838181106139485761394861621e565b9050602002013511155b61396f57604051631353603160e01b815260040160405180910390fd5b6139b061397f604088018861624b565b8381811061398f5761398f61621e565b90506020020160208101906139a491906155b2565b60995f61152089612f55565b6139cd576040516331bc342760e11b815260040160405180910390fd5b5f80613a1f6139df60208a018a6155b2565b6139e889612f55565b6139f560408c018c61624b565b87818110613a0557613a0561621e565b9050602002016020810190613a1a91906155b2565b612fb8565b805191935091506001600160401b03165f03613a3c575050613e2c565b5f613a77613a4d60608b018b61624b565b86818110613a5d57613a5d61621e565b85516001600160401b031692602090910201359050614bdd565b8351909150613a926001600160401b03808416908316614bb5565b868681518110613aa457613aa461621e565b60200260200101818152505081835f01818151613ac191906166de565b6001600160401b0316905250835182908590613ade9083906166de565b6001600160401b0316905250602084018051839190613afe9083906166de565b6001600160401b031690525060208301515f600f9190910b1215613c16575f613b61613b2d60608d018d61624b565b88818110613b3d57613b3d61621e565b905060200201358560200151613b52906166fd565b6001600160801b031690614bdd565b9050806001600160401b031684602001818151613b7e9190616721565b600f0b9052507f1487af5418c47ee5ea45ef4a93398668120890774a9e13487e61e9dc3baf76dd613bb260208d018d6155b2565b8b613bc060408f018f61624b565b8a818110613bd057613bd061621e565b9050602002016020810190613be591906155b2565b613bf6885f015189602001516147dd565b8860400151604051613c0c959493929190616373565b60405180910390a1505b613c68613c2660208c018c6155b2565b613c2f8b612f55565b613c3c60408e018e61624b565b89818110613c4c57613c4c61621e565b9050602002016020810190613c6191906155b2565b87876147fc565b7f1487af5418c47ee5ea45ef4a93398668120890774a9e13487e61e9dc3baf76dd613c9660208c018c6155b2565b8a613ca460408e018e61624b565b89818110613cb457613cb461621e565b9050602002016020810190613cc991906155b2565b8651604051613cdd94939291904390616373565b60405180910390a1613d2e613cf560208c018c6155b2565b613d0260408d018d61624b565b88818110613d1257613d1261621e565b9050602002016020810190613d2791906155b2565b8651614bf3565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016635ae679a7613d6a60208d018d6155b2565b8b8b8e8060400190613d7c919061624b565b8b818110613d8c57613d8c61621e565b9050602002016020810190613da191906155b2565b89516040516001600160e01b031960e088901b168152613dc99594939291899160040161674e565b6020604051808303815f875af1158015613de5573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613e0991906167a1565b878681518110613e1b57613e1b61621e565b602002602001018181525050505050505b600101613838565b507f80969ad29428d6797ee7aad084f9e4a42a82fc506dcd2ca3b6fb431f85ccebe5613e6360208701876155b2565b85613e71604089018961624b565b85613e7f60808c018c6164b4565b604051613e9297969594939291906167b8565b60405180910390a1509250929050565b6040805180820182526001600160a01b038516808252845163ffffffff90811660208085018290525f93845260989052939091209192613ee39291614a4716565b613f0057604051631fb1705560e21b815260040160405180910390fd5b7f31629285ead2335ae0933f86ed2ae63321f7af77b4e6eaabc42c057880977e6c81604051613f2f91906162d0565b60405180910390a16001600160a01b038216620e16e414801590613fc4578260a65f613f5a85612f55565b81526020019081526020015f205f6101000a8154816001600160a01b0302191690836001600160a01b031602179055507f90a6fa2a9b79b910872ebca540cf3bd8be827f586e6420c30d8836e30012907e8284604051613fbb929190616453565b60405180910390a15b5f5b84602001515181101561400357613ffb8386602001518381518110613fed57613fed61621e565b602002602001015184614124565b600101613fc6565b506120f28285604001516001614432565b60605f61099183614c75565b6001600160a01b038381165f90815260a360209081526040808320938616835292905290812054600f81810b600160801b909204900b035b5f8111801561406a57508261ffff1682105b156120f2576001600160a01b038086165f90815260a360209081526040808320938816835292905290812061409e90614cce565b90505f5f6140ad888489612fb8565b91509150806040015163ffffffff164310156140cb575050506120f2565b6140d888848985856147fc565b6001600160a01b038089165f90815260a360209081526040808320938b1683529290522061410590614d20565b5061410f856166b3565b945061411a8461684e565b9350505050614058565b80156141a6576001600160a01b03821673beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac01480159061418957507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316826001600160a01b031614155b6141a657604051632711b74d60e11b815260040160405180910390fd5b6141b68260995f61266d87612f55565b6141d35760405163585cfb2f60e01b815260040160405180910390fd5b7f7ab260fe0af193db5f4986770d831bda4ea46099dc817e8b6716dcae8af8e88b8383604051612c64929190616453565b60605f61421083614d9d565b6040805160208082528183019092529192505f91906020820181803683375050509182525060208101929092525090565b6001600160a01b0383165f908152609b60209081526040918290208251608081018452905463ffffffff808216835260ff600160201b830416151593830193909352650100000000008104831693820193909352600160481b9092041660608201819052158015906142bd5750806060015163ffffffff164310155b156142d757604081015163ffffffff168152600160208201525b63ffffffff8316604082015281156142fa5763ffffffff4316606082015261433b565b6143247f0000000000000000000000000000000000000000000000000000000000000000436162f2565b61432f9060016162f2565b63ffffffff1660608201525b6001600160a01b0384165f818152609b60209081526040918290208451815486840151878601516060808a015163ffffffff95861664ffffffffff1990951694909417600160201b93151593909302929092176cffffffffffffffff0000000000191665010000000000918516919091026cffffffff000000000000000000191617600160481b92841692830217909355845195865290881692850192909252918301527f4e85751d6331506c6c62335f207eb31f12a61e570f34f5c17640308785c6d4db91015b60405180910390a150505050565b6001600160a01b0381165f9081526001830160205260408120541515610991565b6001600160a01b038216614459576040516339b190bb60e11b815260040160405180910390fd5b5f60a75f61446686612f55565b815260208082019290925260409081015f20815160608101835281546001600160a01b03908116825260019092015491821693810193909352600160a01b900463ffffffff16908201819052909150158015906144cd5750806040015163ffffffff164310155b156144e35760208101516001600160a01b031681525b6001600160a01b038316602082015281156145095763ffffffff4316604082015261454a565b6145337f0000000000000000000000000000000000000000000000000000000000000000436162f2565b61453e9060016162f2565b63ffffffff1660408201525b8060a75f61455787612f55565b815260208082019290925260409081015f20835181546001600160a01b039182166001600160a01b031990911617825592840151600190910180549483015163ffffffff16600160a01b026001600160c01b031990951691909316179290921790558181015190517f3873f29d7a65a4d75f5ba28909172f486216a1420e77c3c2720815951a6b4f57916144039187918791616863565b6001600160a01b038281165f81815260a2602090815260408083209486168084529482528083205493835260a38252808320948352939052918220546001600160401b039091169190600f81810b600160801b909204900b03815b81811015614719576001600160a01b038087165f90815260a36020908152604080832093891683529290529081206146819083614dc4565b6001600160a01b038881165f90815260a0602090815260408083208584528252808320938b16835292815290829020825160608101845290546001600160401b0381168252600160401b8104600f0b92820192909252600160c01b90910463ffffffff169181018290529192504310156146fc575050614719565b61470a8682602001516147dd565b95505050806001019050614649565b506001600160a01b038086165f90815260a160209081526040808320938816835292905220839061474990614a34565b61475391906166de565b9150509250929050565b5f6109918383614e33565b5f610991836001600160a01b038416614e33565b5f6109918383670de0b6b3a7640000614f16565b5f6147a18460995f61152089612f55565b80156147aa5750815b80156115fa57505090516001600160401b031615159392505050565b5f6109916001600160401b03808516908416616896565b5f6109916147f4836001600160401b038616616721565b600f0b614f6b565b6020808301516001600160a01b038088165f90815260a284526040808220928816825291909352909120546001600160401b039081169116146148c257602082810180516001600160a01b038881165f81815260a286526040808220938a1680835293875290819020805467ffffffffffffffff19166001600160401b0395861617905593518451918252948101919091529216908201527facf9095feb3a370c9cf692421c69ef320d4db5c66e6a7d29c7694eb02364fc559060600160405180910390a15b6001600160a01b038086165f90815260a060209081526040808320888452825280832093871683529281529082902083518154928501519385015163ffffffff16600160c01b0263ffffffff60c01b196001600160801b038616600160401b026001600160c01b03199095166001600160401b03909316929092179390931716919091179055600f0b156149a4576001600160a01b0385165f908152609f60209081526040808320878452909152902061497c9084614a52565b506001600160a01b0385165f908152609d6020526040902061499e9085614a47565b506120f2565b80516001600160401b03165f036120f2576001600160a01b0385165f908152609f6020908152604080832087845290915290206149e19084614768565b506001600160a01b0385165f908152609f602090815260408083208784529091529020614a0d90613204565b5f036120f2576001600160a01b0385165f908152609d60205260409020610d2b908561475d565b5f610a4682670de0b6b3a7640000614fd6565b5f610991838361500d565b5f610991836001600160a01b03841661500d565b5f805b8251811015614ad557614a9584848381518110614a8857614a8861621e565b60200260200101516109d1565b8015614abe5750614abe838281518110614ab157614ab161621e565b6020026020010151612dac565b15614acd576001915050610a46565b600101614a69565b505f9392505050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015614b3a573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190614b5e91906168c3565b6001600160a01b0316336001600160a01b03161461150d5760405163794821ff60e01b815260040160405180910390fd5b5f825f018281548110614ba457614ba461621e565b905f5260205f200154905092915050565b5f61099183670de0b6b3a764000084615059565b5f6109918383670de0b6b3a7640000615059565b5f6109918383670de0b6b3a7640000600161513e565b6001600160a01b038084165f90815260a160209081526040808320938616835292905220614c2290438361518d565b604080516001600160a01b038086168252841660208201526001600160401b038316918101919091527f1c6458079a41077d003c11faf9bf097e693bd67979e4e6500bac7b29db779b5c90606001612c64565b6060815f01805480602002602001604051908101604052809291908181526020018280548015614cc257602002820191905f5260205f20905b815481526020019060010190808311614cae575b50505050509050919050565b5f614ce88254600f81810b600160801b909204900b131590565b15614d0657604051631ed9509560e11b815260040160405180910390fd5b508054600f0b5f9081526001909101602052604090205490565b5f614d3a8254600f81810b600160801b909204900b131590565b15614d5857604051631ed9509560e11b815260040160405180910390fd5b508054600f0b5f818152600180840160205260408220805492905583546fffffffffffffffffffffffffffffffff191692016001600160801b03169190911790915590565b5f60ff8216601f811115610a4657604051632cd44ac360e21b815260040160405180910390fd5b5f5f614de6614dd2846151a1565b8554614de19190600f0b6168de565b61520a565b8454909150600160801b9004600f90810b9082900b12614e1957604051632d0483c560e21b815260040160405180910390fd5b600f0b5f9081526001939093016020525050604090205490565b5f8181526001830160205260408120548015614f0d575f614e556001836166cb565b85549091505f90614e68906001906166cb565b9050818114614ec7575f865f018281548110614e8657614e8661621e565b905f5260205f200154905080875f018481548110614ea657614ea661621e565b5f918252602080832090910192909255918252600188019052604090208390555b8554869080614ed857614ed8616905565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f905560019350505050610a46565b5f915050610a46565b82545f9081614f2786868385615273565b90508015614f6157614f4b86614f3e6001846166cb565b5f91825260209091200190565b54600160201b90046001600160e01b03166109c7565b5091949350505050565b5f6001600160401b03821115614fd25760405162461bcd60e51b815260206004820152602660248201527f53616665436173743a2076616c756520646f65736e27742066697420696e203660448201526534206269747360d01b6064820152608401612eda565b5090565b81545f90801561500557614fef84614f3e6001846166cb565b54600160201b90046001600160e01b0316610a42565b509092915050565b5f81815260018301602052604081205461505257508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155610a46565b505f610a46565b5f80805f19858709858702925082811083820303915050805f036150905783828161508657615086616919565b0492505050610991565b8084116150d75760405162461bcd60e51b81526020600482015260156024820152744d6174683a206d756c446976206f766572666c6f7760581b6044820152606401612eda565b5f8486880960026001871981018816978890046003810283188082028403028082028403028082028403028082028403028082028403029081029092039091025f889003889004909101858311909403939093029303949094049190911702949350505050565b5f5f61514b868686615059565b905060018360028111156151615761516161692d565b14801561517d57505f848061517857615178616919565b868809115b156115fa576109c7600182616941565b6114f683836001600160401b0384166152c6565b5f6001600160ff1b03821115614fd25760405162461bcd60e51b815260206004820152602860248201527f53616665436173743a2076616c756520646f65736e27742066697420696e2061604482015267371034b73a191a9b60c11b6064820152608401612eda565b80600f81900b811461526e5760405162461bcd60e51b815260206004820152602760248201527f53616665436173743a2076616c756520646f65736e27742066697420696e20316044820152663238206269747360c81b6064820152608401612eda565b919050565b5f5b81831015610b95575f61528884846153c9565b5f8781526020902090915063ffffffff86169082015463ffffffff1611156152b2578092506152c0565b6152bd816001616941565b93505b50615275565b8254801561537c575f6152de85614f3e6001856166cb565b60408051808201909152905463ffffffff808216808452600160201b9092046001600160e01b0316602084015291925090851610156153305760405163151b8e3f60e11b815260040160405180910390fd5b805163ffffffff80861691160361537a578261535186614f3e6001866166cb565b80546001600160e01b0392909216600160201b0263ffffffff9092169190911790555050505050565b505b506040805180820190915263ffffffff92831681526001600160e01b03918216602080830191825285546001810187555f968752952091519051909216600160201b029190921617910155565b5f6153d76002848418616954565b61099190848416616941565b634e487b7160e01b5f52604160045260245ffd5b604051606081016001600160401b0381118282101715615419576154196153e3565b60405290565b604051601f8201601f191681016001600160401b0381118282101715615447576154476153e3565b604052919050565b6001600160a01b038116811461375f575f5ffd5b803563ffffffff8116811461526e575f5ffd5b5f60408284031215615486575f5ffd5b604080519081016001600160401b03811182821017156154a8576154a86153e3565b60405290508082356154b98161544f565b81526154c760208401615463565b60208201525092915050565b5f604082840312156154e3575f5ffd5b6109918383615476565b5f5f5f608084860312156154ff575f5ffd5b833561550a8161544f565b92506155198560208601615476565b915060608401356155298161544f565b809150509250925092565b81516001600160401b03168152602080830151600f0b9082015260408083015163ffffffff169082015260608101610a46565b5f5f60608385031215615578575f5ffd5b82356155838161544f565b91506155928460208501615476565b90509250929050565b5f602082840312156155ab575f5ffd5b5035919050565b5f602082840312156155c2575f5ffd5b81356109918161544f565b80516001600160a01b0316825260209081015163ffffffff16910152565b5f8151808452602084019350602083015f5b82811015615626576156108683516155cd565b60409590950194602091909101906001016155fd565b5093949350505050565b602081525f61099160208301846155eb565b5f5f83601f840112615652575f5ffd5b5081356001600160401b03811115615668575f5ffd5b6020830191508360208260051b8501011115611095575f5ffd5b5f5f5f60408486031215615694575f5ffd5b833561569f8161544f565b925060208401356001600160401b038111156156b9575f5ffd5b6156c586828701615642565b9497909650939450505050565b5f6001600160401b038211156156ea576156ea6153e3565b5060051b60200190565b5f82601f830112615703575f5ffd5b8135615716615711826156d2565b61541f565b8082825260208201915060208360051b860101925085831115615737575f5ffd5b602085015b8381101561575d57803561574f8161544f565b83526020928301920161573c565b5095945050505050565b5f5f5f60808486031215615779575f5ffd5b6157838585615476565b925060408401356001600160401b0381111561579d575f5ffd5b6157a9868287016156f4565b92505060608401356001600160401b038111156157c4575f5ffd5b6157d0868287016156f4565b9150509250925092565b5f8151808452602084019350602083015f5b828110156156265781518652602095860195909101906001016157ec565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b8281101561586157603f1987860301845261584c8583516157da565b94506020938401939190910190600101615830565b50929695505050505050565b5f5f5f5f60a08587031215615880575f5ffd5b61588a8686615476565b935060408501356001600160401b038111156158a4575f5ffd5b6158b0878288016156f4565b93505060608501356001600160401b038111156158cb575f5ffd5b6158d7878288016156f4565b9250506158e660808601615463565b905092959194509250565b5f5f5f5f5f60608688031215615905575f5ffd5b85356159108161544f565b945060208601356001600160401b0381111561592a575f5ffd5b61593688828901615642565b90955093505060408601356001600160401b03811115615954575f5ffd5b61596088828901615642565b969995985093965092949392505050565b5f5f60408385031215615982575f5ffd5b823561598d8161544f565b915060208301356001600160401b038111156159a7575f5ffd5b830160a081860312156159b8575f5ffd5b809150509250929050565b828152604060208201525f610c1860408301846157da565b5f82601f8301126159ea575f5ffd5b81356159f8615711826156d2565b8082825260208201915060208360051b860101925085831115615a19575f5ffd5b602085015b8381101561575d5780356001600160401b03811115615a3b575f5ffd5b86016060818903601f19011215615a50575f5ffd5b615a586153f7565b615a6460208301615463565b815260408201356001600160401b03811115615a7e575f5ffd5b615a8d8a6020838601016156f4565b60208301525060608201359150615aa38261544f565b6040810191909152835260209283019201615a1e565b5f5f60408385031215615aca575f5ffd5b8235615ad58161544f565b915060208301356001600160401b03811115615aef575f5ffd5b614753858286016159db565b5f5f60408385031215615b0c575f5ffd5b8235615b178161544f565b915060208301356159b88161544f565b5f8151808452602084019350602083015f5b8281101561562657615b7286835180516001600160401b03168252602080820151600f0b9083015260409081015163ffffffff16910152565b6060959095019460209190910190600101615b39565b604081525f615b9a60408301856155eb565b82810360208401526115fa8185615b27565b5f8151808452602084019350602083015f5b828110156156265781516001600160a01b0316865260209586019590910190600101615bbe565b602081525f6109916020830184615bac565b5f5f60408385031215615c08575f5ffd5b82356001600160401b03811115615c1d575f5ffd5b615c29858286016156f4565b92505060208301356159b88161544f565b602080825282518282018190525f918401906040840190835b81811015615c7a5783516001600160401b0316835260209384019390920191600101615c53565b509095945050505050565b5f5f5f5f60608587031215615c98575f5ffd5b8435615ca38161544f565b9350615cb160208601615463565b925060408501356001600160401b03811115615ccb575f5ffd5b615cd787828801615642565b95989497509550505050565b5f5f60408385031215615cf4575f5ffd5b8235615cff8161544f565b915060208301356001600160401b03811115615d19575f5ffd5b614753858286016156f4565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b5f5f60408385031215615d6b575f5ffd5b8235615d768161544f565b915061559260208401615463565b5f60208284031215615d94575f5ffd5b813560ff81168114610991575f5ffd5b5f5f60608385031215615db5575f5ffd5b615dbf8484615476565b915060408301356159b88161544f565b5f60608284031215615ddf575f5ffd5b50919050565b5f60208284031215615df5575f5ffd5b81356001600160401b03811115615e0a575f5ffd5b610a4284828501615dcf565b5f5f5f60808486031215615e28575f5ffd5b83356001600160401b03811115615e3d575f5ffd5b615e49868287016156f4565b9350506155198560208601615476565b602081525f6109916020830184615b27565b5f5f5f60608486031215615e7d575f5ffd5b8335615e888161544f565b925060208401356001600160401b03811115615ea2575f5ffd5b615eae868287016156f4565b925050615ebd60408501615463565b90509250925092565b5f5f60408385031215615ed7575f5ffd5b8235615ee28161544f565b915060208301356001600160401b03811115615efc575f5ffd5b8301601f81018513615f0c575f5ffd5b8035615f1a615711826156d2565b8082825260208201915060208360051b850101925087831115615f3b575f5ffd5b602084015b838110156160605780356001600160401b03811115615f5d575f5ffd5b85016080818b03601f19011215615f72575f5ffd5b615f7a6153f7565b615f878b60208401615476565b815260608201356001600160401b03811115615fa1575f5ffd5b615fb08c6020838601016156f4565b60208301525060808201356001600160401b03811115615fce575f5ffd5b6020818401019250508a601f830112615fe5575f5ffd5b8135615ff3615711826156d2565b8082825260208201915060208360051b86010192508d831115616014575f5ffd5b6020850194505b8285101561604a5784356001600160401b0381168114616039575f5ffd5b82526020948501949091019061601b565b6040840152505084525060209283019201615f40565b50809450505050509250929050565b5f6020828403121561607f575f5ffd5b81356001600160401b03811115616094575f5ffd5b8201601f810184136160a4575f5ffd5b80356160b2615711826156d2565b8082825260208201915060208360061b8501019250868311156160d3575f5ffd5b6020840193505b828410156109c7576160ec8785615476565b82526020820191506040840193506160da565b5f5f5f60408486031215616111575f5ffd5b833561611c8161544f565b925060208401356001600160401b03811115616136575f5ffd5b8401601f81018613616146575f5ffd5b80356001600160401b0381111561615b575f5ffd5b86602082840101111561616c575f5ffd5b939660209190910195509293505050565b5f5f6040838503121561618e575f5ffd5b82356161998161544f565b915060208301356001600160401b038111156161b3575f5ffd5b61475385828601615dcf565b5f5f5f5f606085870312156161d2575f5ffd5b84356161dd8161544f565b935060208501356001600160401b038111156161f7575f5ffd5b616203878288016159db565b93505060408501356001600160401b03811115615ccb575f5ffd5b634e487b7160e01b5f52603260045260245ffd5b5f60208284031215616242575f5ffd5b61099182615463565b5f5f8335601e19843603018112616260575f5ffd5b8301803591506001600160401b03821115616279575f5ffd5b6020019150600581901b3603821315611095575f5ffd5b5f602082840312156162a0575f5ffd5b813561ffff81168114610991575f5ffd5b5f602082840312156162c1575f5ffd5b81518015158114610991575f5ffd5b60408101610a4682846155cd565b634e487b7160e01b5f52601160045260245ffd5b63ffffffff8181168382160190811115610a4657610a466162de565b8183526020830192505f815f5b848110156156265763ffffffff61633183615463565b168652602095860195919091019060010161631b565b6001600160a01b038581168252841660208201526060604082018190525f906109c7908301848661630e565b6001600160a01b038616815260c0810161639060208301876155cd565b6001600160a01b039490941660608201526001600160401b0392909216608083015263ffffffff1660a09091015292915050565b5f602082840312156163d4575f5ffd5b81516001600160401b038111156163e9575f5ffd5b8201601f810184136163f9575f5ffd5b8051616407615711826156d2565b8082825260208201915060208360051b850101925086831115616428575f5ffd5b6020840193505b828410156109c75783516164428161544f565b82526020938401939091019061642f565b6060810161646182856155cd565b6001600160a01b039290921660409190910152919050565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b602081525f610c18602083018486616479565b5f5f8335601e198436030181126164c9575f5ffd5b8301803591506001600160401b038211156164e2575f5ffd5b602001915036819003821315611095575f5ffd5b6001600160a01b038781168252861660208201526080604082018190525f90616522908301868861630e565b8281036060840152616535818587616479565b9998505050505050505050565b80516020808301519190811015615ddf575f1960209190910360031b1b16919050565b5f8235603e19833603018112616579575f5ffd5b9190910192915050565b604081525f6165956040830185615bac565b82810360208401526115fa8185615bac565b5f602082840312156165b7575f5ffd5b81516001600160401b038111156165cc575f5ffd5b8201601f810184136165dc575f5ffd5b80516165ea615711826156d2565b8082825260208201915060208360051b85010192508683111561660b575f5ffd5b602084015b838110156166a85780516001600160401b0381111561662d575f5ffd5b8501603f8101891361663d575f5ffd5b602081015161664e615711826156d2565b808282526020820191506020808460051b8601010192508b831115616671575f5ffd5b6040840193505b82841015616693578351825260209384019390910190616678565b86525050602093840193919091019050616610565b509695505050505050565b5f600182016166c4576166c46162de565b5060010190565b81810381811115610a4657610a466162de565b6001600160401b038281168282160390811115610a4657610a466162de565b5f81600f0b60016001607f1b03198103616719576167196162de565b5f0392915050565b600f81810b9083900b0160016001607f1b03811360016001607f1b031982121715610a4657610a466162de565b6001600160a01b038716815260e0810161676b60208301886155cd565b60608201959095526001600160a01b039390931660808401526001600160401b0391821660a08401521660c09091015292915050565b5f602082840312156167b1575f5ffd5b5051919050565b6001600160a01b03881681525f60c082016167d6602084018a6155cd565b60c060608401528690528660e083015f5b888110156168175782356167fa8161544f565b6001600160a01b03168252602092830192909101906001016167e7565b50838103608085015261682a81886157da565b91505082810360a0840152616840818587616479565b9a9950505050505050505050565b5f8161685c5761685c6162de565b505f190190565b6080810161687182866155cd565b6001600160a01b0393909316604082015263ffffffff91909116606090910152919050565b600f82810b9082900b0360016001607f1b0319811260016001607f1b0382131715610a4657610a466162de565b5f602082840312156168d3575f5ffd5b81516109918161544f565b8082018281125f8312801582168215821617156168fd576168fd6162de565b505092915050565b634e487b7160e01b5f52603160045260245ffd5b634e487b7160e01b5f52601260045260245ffd5b634e487b7160e01b5f52602160045260245ffd5b80820180821115610a4657610a466162de565b5f8261696e57634e487b7160e01b5f52601260045260245ffd5b50049056fea26469706673582212205c645b8edb93cacb51b54b9d2abca739a9a47c7255500417db11c96ef053da1164736f6c634300081b0033",
}

// AllocationManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use AllocationManagerMetaData.ABI instead.
var AllocationManagerABI = AllocationManagerMetaData.ABI

// AllocationManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AllocationManagerMetaData.Bin instead.
var AllocationManagerBin = AllocationManagerMetaData.Bin

// DeployAllocationManager deploys a new Ethereum contract, binding an instance of AllocationManager to it.
func DeployAllocationManager(auth *bind.TransactOpts, backend bind.ContractBackend, _delegation common.Address, _eigenStrategy common.Address, _pauserRegistry common.Address, _permissionController common.Address, _DEALLOCATION_DELAY uint32, _ALLOCATION_CONFIGURATION_DELAY uint32, _version string) (common.Address, *types.Transaction, *AllocationManager, error) {
	parsed, err := AllocationManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AllocationManagerBin), backend, _delegation, _eigenStrategy, _pauserRegistry, _permissionController, _DEALLOCATION_DELAY, _ALLOCATION_CONFIGURATION_DELAY, _version)
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

// EigenStrategy is a free data retrieval call binding the contract method 0xdb4df761.
//
// Solidity: function eigenStrategy() view returns(address)
func (_AllocationManager *AllocationManagerCaller) EigenStrategy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AllocationManager.contract.Call(opts, &out, "eigenStrategy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EigenStrategy is a free data retrieval call binding the contract method 0xdb4df761.
//
// Solidity: function eigenStrategy() view returns(address)
func (_AllocationManager *AllocationManagerSession) EigenStrategy() (common.Address, error) {
	return _AllocationManager.Contract.EigenStrategy(&_AllocationManager.CallOpts)
}

// EigenStrategy is a free data retrieval call binding the contract method 0xdb4df761.
//
// Solidity: function eigenStrategy() view returns(address)
func (_AllocationManager *AllocationManagerCallerSession) EigenStrategy() (common.Address, error) {
	return _AllocationManager.Contract.EigenStrategy(&_AllocationManager.CallOpts)
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

// GetPendingSlasher is a free data retrieval call binding the contract method 0xd7794857.
//
// Solidity: function getPendingSlasher((address,uint32) operatorSet) view returns(address, uint32)
func (_AllocationManager *AllocationManagerCaller) GetPendingSlasher(opts *bind.CallOpts, operatorSet OperatorSet) (common.Address, uint32, error) {
	var out []interface{}
	err := _AllocationManager.contract.Call(opts, &out, "getPendingSlasher", operatorSet)

	if err != nil {
		return *new(common.Address), *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return out0, out1, err

}

// GetPendingSlasher is a free data retrieval call binding the contract method 0xd7794857.
//
// Solidity: function getPendingSlasher((address,uint32) operatorSet) view returns(address, uint32)
func (_AllocationManager *AllocationManagerSession) GetPendingSlasher(operatorSet OperatorSet) (common.Address, uint32, error) {
	return _AllocationManager.Contract.GetPendingSlasher(&_AllocationManager.CallOpts, operatorSet)
}

// GetPendingSlasher is a free data retrieval call binding the contract method 0xd7794857.
//
// Solidity: function getPendingSlasher((address,uint32) operatorSet) view returns(address, uint32)
func (_AllocationManager *AllocationManagerCallerSession) GetPendingSlasher(operatorSet OperatorSet) (common.Address, uint32, error) {
	return _AllocationManager.Contract.GetPendingSlasher(&_AllocationManager.CallOpts, operatorSet)
}

// GetRedistributionRecipient is a free data retrieval call binding the contract method 0x0f3df50e.
//
// Solidity: function getRedistributionRecipient((address,uint32) operatorSet) view returns(address)
func (_AllocationManager *AllocationManagerCaller) GetRedistributionRecipient(opts *bind.CallOpts, operatorSet OperatorSet) (common.Address, error) {
	var out []interface{}
	err := _AllocationManager.contract.Call(opts, &out, "getRedistributionRecipient", operatorSet)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRedistributionRecipient is a free data retrieval call binding the contract method 0x0f3df50e.
//
// Solidity: function getRedistributionRecipient((address,uint32) operatorSet) view returns(address)
func (_AllocationManager *AllocationManagerSession) GetRedistributionRecipient(operatorSet OperatorSet) (common.Address, error) {
	return _AllocationManager.Contract.GetRedistributionRecipient(&_AllocationManager.CallOpts, operatorSet)
}

// GetRedistributionRecipient is a free data retrieval call binding the contract method 0x0f3df50e.
//
// Solidity: function getRedistributionRecipient((address,uint32) operatorSet) view returns(address)
func (_AllocationManager *AllocationManagerCallerSession) GetRedistributionRecipient(operatorSet OperatorSet) (common.Address, error) {
	return _AllocationManager.Contract.GetRedistributionRecipient(&_AllocationManager.CallOpts, operatorSet)
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

// GetSlashCount is a free data retrieval call binding the contract method 0x4cfd2939.
//
// Solidity: function getSlashCount((address,uint32) operatorSet) view returns(uint256)
func (_AllocationManager *AllocationManagerCaller) GetSlashCount(opts *bind.CallOpts, operatorSet OperatorSet) (*big.Int, error) {
	var out []interface{}
	err := _AllocationManager.contract.Call(opts, &out, "getSlashCount", operatorSet)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSlashCount is a free data retrieval call binding the contract method 0x4cfd2939.
//
// Solidity: function getSlashCount((address,uint32) operatorSet) view returns(uint256)
func (_AllocationManager *AllocationManagerSession) GetSlashCount(operatorSet OperatorSet) (*big.Int, error) {
	return _AllocationManager.Contract.GetSlashCount(&_AllocationManager.CallOpts, operatorSet)
}

// GetSlashCount is a free data retrieval call binding the contract method 0x4cfd2939.
//
// Solidity: function getSlashCount((address,uint32) operatorSet) view returns(uint256)
func (_AllocationManager *AllocationManagerCallerSession) GetSlashCount(operatorSet OperatorSet) (*big.Int, error) {
	return _AllocationManager.Contract.GetSlashCount(&_AllocationManager.CallOpts, operatorSet)
}

// GetSlasher is a free data retrieval call binding the contract method 0xd4a3fcce.
//
// Solidity: function getSlasher((address,uint32) operatorSet) view returns(address)
func (_AllocationManager *AllocationManagerCaller) GetSlasher(opts *bind.CallOpts, operatorSet OperatorSet) (common.Address, error) {
	var out []interface{}
	err := _AllocationManager.contract.Call(opts, &out, "getSlasher", operatorSet)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSlasher is a free data retrieval call binding the contract method 0xd4a3fcce.
//
// Solidity: function getSlasher((address,uint32) operatorSet) view returns(address)
func (_AllocationManager *AllocationManagerSession) GetSlasher(operatorSet OperatorSet) (common.Address, error) {
	return _AllocationManager.Contract.GetSlasher(&_AllocationManager.CallOpts, operatorSet)
}

// GetSlasher is a free data retrieval call binding the contract method 0xd4a3fcce.
//
// Solidity: function getSlasher((address,uint32) operatorSet) view returns(address)
func (_AllocationManager *AllocationManagerCallerSession) GetSlasher(operatorSet OperatorSet) (common.Address, error) {
	return _AllocationManager.Contract.GetSlasher(&_AllocationManager.CallOpts, operatorSet)
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

// IsOperatorRedistributable is a free data retrieval call binding the contract method 0xdc2af692.
//
// Solidity: function isOperatorRedistributable(address operator) view returns(bool)
func (_AllocationManager *AllocationManagerCaller) IsOperatorRedistributable(opts *bind.CallOpts, operator common.Address) (bool, error) {
	var out []interface{}
	err := _AllocationManager.contract.Call(opts, &out, "isOperatorRedistributable", operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperatorRedistributable is a free data retrieval call binding the contract method 0xdc2af692.
//
// Solidity: function isOperatorRedistributable(address operator) view returns(bool)
func (_AllocationManager *AllocationManagerSession) IsOperatorRedistributable(operator common.Address) (bool, error) {
	return _AllocationManager.Contract.IsOperatorRedistributable(&_AllocationManager.CallOpts, operator)
}

// IsOperatorRedistributable is a free data retrieval call binding the contract method 0xdc2af692.
//
// Solidity: function isOperatorRedistributable(address operator) view returns(bool)
func (_AllocationManager *AllocationManagerCallerSession) IsOperatorRedistributable(operator common.Address) (bool, error) {
	return _AllocationManager.Contract.IsOperatorRedistributable(&_AllocationManager.CallOpts, operator)
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

// IsRedistributingOperatorSet is a free data retrieval call binding the contract method 0xf231bd08.
//
// Solidity: function isRedistributingOperatorSet((address,uint32) operatorSet) view returns(bool)
func (_AllocationManager *AllocationManagerCaller) IsRedistributingOperatorSet(opts *bind.CallOpts, operatorSet OperatorSet) (bool, error) {
	var out []interface{}
	err := _AllocationManager.contract.Call(opts, &out, "isRedistributingOperatorSet", operatorSet)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRedistributingOperatorSet is a free data retrieval call binding the contract method 0xf231bd08.
//
// Solidity: function isRedistributingOperatorSet((address,uint32) operatorSet) view returns(bool)
func (_AllocationManager *AllocationManagerSession) IsRedistributingOperatorSet(operatorSet OperatorSet) (bool, error) {
	return _AllocationManager.Contract.IsRedistributingOperatorSet(&_AllocationManager.CallOpts, operatorSet)
}

// IsRedistributingOperatorSet is a free data retrieval call binding the contract method 0xf231bd08.
//
// Solidity: function isRedistributingOperatorSet((address,uint32) operatorSet) view returns(bool)
func (_AllocationManager *AllocationManagerCallerSession) IsRedistributingOperatorSet(operatorSet OperatorSet) (bool, error) {
	return _AllocationManager.Contract.IsRedistributingOperatorSet(&_AllocationManager.CallOpts, operatorSet)
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

// CreateOperatorSets0 is a paid mutator transaction binding the contract method 0x3dff8e7d.
//
// Solidity: function createOperatorSets(address avs, (uint32,address[],address)[] params) returns()
func (_AllocationManager *AllocationManagerTransactor) CreateOperatorSets0(opts *bind.TransactOpts, avs common.Address, params []IAllocationManagerTypesCreateSetParamsV2) (*types.Transaction, error) {
	return _AllocationManager.contract.Transact(opts, "createOperatorSets0", avs, params)
}

// CreateOperatorSets0 is a paid mutator transaction binding the contract method 0x3dff8e7d.
//
// Solidity: function createOperatorSets(address avs, (uint32,address[],address)[] params) returns()
func (_AllocationManager *AllocationManagerSession) CreateOperatorSets0(avs common.Address, params []IAllocationManagerTypesCreateSetParamsV2) (*types.Transaction, error) {
	return _AllocationManager.Contract.CreateOperatorSets0(&_AllocationManager.TransactOpts, avs, params)
}

// CreateOperatorSets0 is a paid mutator transaction binding the contract method 0x3dff8e7d.
//
// Solidity: function createOperatorSets(address avs, (uint32,address[],address)[] params) returns()
func (_AllocationManager *AllocationManagerTransactorSession) CreateOperatorSets0(avs common.Address, params []IAllocationManagerTypesCreateSetParamsV2) (*types.Transaction, error) {
	return _AllocationManager.Contract.CreateOperatorSets0(&_AllocationManager.TransactOpts, avs, params)
}

// CreateRedistributingOperatorSets is a paid mutator transaction binding the contract method 0x32a879e4.
//
// Solidity: function createRedistributingOperatorSets(address avs, (uint32,address[])[] params, address[] redistributionRecipients) returns()
func (_AllocationManager *AllocationManagerTransactor) CreateRedistributingOperatorSets(opts *bind.TransactOpts, avs common.Address, params []IAllocationManagerTypesCreateSetParams, redistributionRecipients []common.Address) (*types.Transaction, error) {
	return _AllocationManager.contract.Transact(opts, "createRedistributingOperatorSets", avs, params, redistributionRecipients)
}

// CreateRedistributingOperatorSets is a paid mutator transaction binding the contract method 0x32a879e4.
//
// Solidity: function createRedistributingOperatorSets(address avs, (uint32,address[])[] params, address[] redistributionRecipients) returns()
func (_AllocationManager *AllocationManagerSession) CreateRedistributingOperatorSets(avs common.Address, params []IAllocationManagerTypesCreateSetParams, redistributionRecipients []common.Address) (*types.Transaction, error) {
	return _AllocationManager.Contract.CreateRedistributingOperatorSets(&_AllocationManager.TransactOpts, avs, params, redistributionRecipients)
}

// CreateRedistributingOperatorSets is a paid mutator transaction binding the contract method 0x32a879e4.
//
// Solidity: function createRedistributingOperatorSets(address avs, (uint32,address[])[] params, address[] redistributionRecipients) returns()
func (_AllocationManager *AllocationManagerTransactorSession) CreateRedistributingOperatorSets(avs common.Address, params []IAllocationManagerTypesCreateSetParams, redistributionRecipients []common.Address) (*types.Transaction, error) {
	return _AllocationManager.Contract.CreateRedistributingOperatorSets(&_AllocationManager.TransactOpts, avs, params, redistributionRecipients)
}

// CreateRedistributingOperatorSets0 is a paid mutator transaction binding the contract method 0xd1a83f54.
//
// Solidity: function createRedistributingOperatorSets(address avs, (uint32,address[],address)[] params, address[] redistributionRecipients) returns()
func (_AllocationManager *AllocationManagerTransactor) CreateRedistributingOperatorSets0(opts *bind.TransactOpts, avs common.Address, params []IAllocationManagerTypesCreateSetParamsV2, redistributionRecipients []common.Address) (*types.Transaction, error) {
	return _AllocationManager.contract.Transact(opts, "createRedistributingOperatorSets0", avs, params, redistributionRecipients)
}

// CreateRedistributingOperatorSets0 is a paid mutator transaction binding the contract method 0xd1a83f54.
//
// Solidity: function createRedistributingOperatorSets(address avs, (uint32,address[],address)[] params, address[] redistributionRecipients) returns()
func (_AllocationManager *AllocationManagerSession) CreateRedistributingOperatorSets0(avs common.Address, params []IAllocationManagerTypesCreateSetParamsV2, redistributionRecipients []common.Address) (*types.Transaction, error) {
	return _AllocationManager.Contract.CreateRedistributingOperatorSets0(&_AllocationManager.TransactOpts, avs, params, redistributionRecipients)
}

// CreateRedistributingOperatorSets0 is a paid mutator transaction binding the contract method 0xd1a83f54.
//
// Solidity: function createRedistributingOperatorSets(address avs, (uint32,address[],address)[] params, address[] redistributionRecipients) returns()
func (_AllocationManager *AllocationManagerTransactorSession) CreateRedistributingOperatorSets0(avs common.Address, params []IAllocationManagerTypesCreateSetParamsV2, redistributionRecipients []common.Address) (*types.Transaction, error) {
	return _AllocationManager.Contract.CreateRedistributingOperatorSets0(&_AllocationManager.TransactOpts, avs, params, redistributionRecipients)
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

// Initialize is a paid mutator transaction binding the contract method 0xfe4b84df.
//
// Solidity: function initialize(uint256 initialPausedStatus) returns()
func (_AllocationManager *AllocationManagerTransactor) Initialize(opts *bind.TransactOpts, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _AllocationManager.contract.Transact(opts, "initialize", initialPausedStatus)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe4b84df.
//
// Solidity: function initialize(uint256 initialPausedStatus) returns()
func (_AllocationManager *AllocationManagerSession) Initialize(initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _AllocationManager.Contract.Initialize(&_AllocationManager.TransactOpts, initialPausedStatus)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe4b84df.
//
// Solidity: function initialize(uint256 initialPausedStatus) returns()
func (_AllocationManager *AllocationManagerTransactorSession) Initialize(initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _AllocationManager.Contract.Initialize(&_AllocationManager.TransactOpts, initialPausedStatus)
}

// MigrateSlashers is a paid mutator transaction binding the contract method 0x957dc50b.
//
// Solidity: function migrateSlashers((address,uint32)[] operatorSets) returns()
func (_AllocationManager *AllocationManagerTransactor) MigrateSlashers(opts *bind.TransactOpts, operatorSets []OperatorSet) (*types.Transaction, error) {
	return _AllocationManager.contract.Transact(opts, "migrateSlashers", operatorSets)
}

// MigrateSlashers is a paid mutator transaction binding the contract method 0x957dc50b.
//
// Solidity: function migrateSlashers((address,uint32)[] operatorSets) returns()
func (_AllocationManager *AllocationManagerSession) MigrateSlashers(operatorSets []OperatorSet) (*types.Transaction, error) {
	return _AllocationManager.Contract.MigrateSlashers(&_AllocationManager.TransactOpts, operatorSets)
}

// MigrateSlashers is a paid mutator transaction binding the contract method 0x957dc50b.
//
// Solidity: function migrateSlashers((address,uint32)[] operatorSets) returns()
func (_AllocationManager *AllocationManagerTransactorSession) MigrateSlashers(operatorSets []OperatorSet) (*types.Transaction, error) {
	return _AllocationManager.Contract.MigrateSlashers(&_AllocationManager.TransactOpts, operatorSets)
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
// Solidity: function slashOperator(address avs, (address,uint32,address[],uint256[],string) params) returns(uint256, uint256[])
func (_AllocationManager *AllocationManagerTransactor) SlashOperator(opts *bind.TransactOpts, avs common.Address, params IAllocationManagerTypesSlashingParams) (*types.Transaction, error) {
	return _AllocationManager.contract.Transact(opts, "slashOperator", avs, params)
}

// SlashOperator is a paid mutator transaction binding the contract method 0x36352057.
//
// Solidity: function slashOperator(address avs, (address,uint32,address[],uint256[],string) params) returns(uint256, uint256[])
func (_AllocationManager *AllocationManagerSession) SlashOperator(avs common.Address, params IAllocationManagerTypesSlashingParams) (*types.Transaction, error) {
	return _AllocationManager.Contract.SlashOperator(&_AllocationManager.TransactOpts, avs, params)
}

// SlashOperator is a paid mutator transaction binding the contract method 0x36352057.
//
// Solidity: function slashOperator(address avs, (address,uint32,address[],uint256[],string) params) returns(uint256, uint256[])
func (_AllocationManager *AllocationManagerTransactorSession) SlashOperator(avs common.Address, params IAllocationManagerTypesSlashingParams) (*types.Transaction, error) {
	return _AllocationManager.Contract.SlashOperator(&_AllocationManager.TransactOpts, avs, params)
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

// UpdateSlasher is a paid mutator transaction binding the contract method 0x6c9d7c58.
//
// Solidity: function updateSlasher((address,uint32) operatorSet, address slasher) returns()
func (_AllocationManager *AllocationManagerTransactor) UpdateSlasher(opts *bind.TransactOpts, operatorSet OperatorSet, slasher common.Address) (*types.Transaction, error) {
	return _AllocationManager.contract.Transact(opts, "updateSlasher", operatorSet, slasher)
}

// UpdateSlasher is a paid mutator transaction binding the contract method 0x6c9d7c58.
//
// Solidity: function updateSlasher((address,uint32) operatorSet, address slasher) returns()
func (_AllocationManager *AllocationManagerSession) UpdateSlasher(operatorSet OperatorSet, slasher common.Address) (*types.Transaction, error) {
	return _AllocationManager.Contract.UpdateSlasher(&_AllocationManager.TransactOpts, operatorSet, slasher)
}

// UpdateSlasher is a paid mutator transaction binding the contract method 0x6c9d7c58.
//
// Solidity: function updateSlasher((address,uint32) operatorSet, address slasher) returns()
func (_AllocationManager *AllocationManagerTransactorSession) UpdateSlasher(operatorSet OperatorSet, slasher common.Address) (*types.Transaction, error) {
	return _AllocationManager.Contract.UpdateSlasher(&_AllocationManager.TransactOpts, operatorSet, slasher)
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

// AllocationManagerRedistributionAddressSetIterator is returned from FilterRedistributionAddressSet and is used to iterate over the raw logs and unpacked data for RedistributionAddressSet events raised by the AllocationManager contract.
type AllocationManagerRedistributionAddressSetIterator struct {
	Event *AllocationManagerRedistributionAddressSet // Event containing the contract specifics and raw log

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
func (it *AllocationManagerRedistributionAddressSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AllocationManagerRedistributionAddressSet)
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
		it.Event = new(AllocationManagerRedistributionAddressSet)
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
func (it *AllocationManagerRedistributionAddressSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AllocationManagerRedistributionAddressSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AllocationManagerRedistributionAddressSet represents a RedistributionAddressSet event raised by the AllocationManager contract.
type AllocationManagerRedistributionAddressSet struct {
	OperatorSet             OperatorSet
	RedistributionRecipient common.Address
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterRedistributionAddressSet is a free log retrieval operation binding the contract event 0x90a6fa2a9b79b910872ebca540cf3bd8be827f586e6420c30d8836e30012907e.
//
// Solidity: event RedistributionAddressSet((address,uint32) operatorSet, address redistributionRecipient)
func (_AllocationManager *AllocationManagerFilterer) FilterRedistributionAddressSet(opts *bind.FilterOpts) (*AllocationManagerRedistributionAddressSetIterator, error) {

	logs, sub, err := _AllocationManager.contract.FilterLogs(opts, "RedistributionAddressSet")
	if err != nil {
		return nil, err
	}
	return &AllocationManagerRedistributionAddressSetIterator{contract: _AllocationManager.contract, event: "RedistributionAddressSet", logs: logs, sub: sub}, nil
}

// WatchRedistributionAddressSet is a free log subscription operation binding the contract event 0x90a6fa2a9b79b910872ebca540cf3bd8be827f586e6420c30d8836e30012907e.
//
// Solidity: event RedistributionAddressSet((address,uint32) operatorSet, address redistributionRecipient)
func (_AllocationManager *AllocationManagerFilterer) WatchRedistributionAddressSet(opts *bind.WatchOpts, sink chan<- *AllocationManagerRedistributionAddressSet) (event.Subscription, error) {

	logs, sub, err := _AllocationManager.contract.WatchLogs(opts, "RedistributionAddressSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AllocationManagerRedistributionAddressSet)
				if err := _AllocationManager.contract.UnpackLog(event, "RedistributionAddressSet", log); err != nil {
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

// ParseRedistributionAddressSet is a log parse operation binding the contract event 0x90a6fa2a9b79b910872ebca540cf3bd8be827f586e6420c30d8836e30012907e.
//
// Solidity: event RedistributionAddressSet((address,uint32) operatorSet, address redistributionRecipient)
func (_AllocationManager *AllocationManagerFilterer) ParseRedistributionAddressSet(log types.Log) (*AllocationManagerRedistributionAddressSet, error) {
	event := new(AllocationManagerRedistributionAddressSet)
	if err := _AllocationManager.contract.UnpackLog(event, "RedistributionAddressSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AllocationManagerSlasherMigratedIterator is returned from FilterSlasherMigrated and is used to iterate over the raw logs and unpacked data for SlasherMigrated events raised by the AllocationManager contract.
type AllocationManagerSlasherMigratedIterator struct {
	Event *AllocationManagerSlasherMigrated // Event containing the contract specifics and raw log

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
func (it *AllocationManagerSlasherMigratedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AllocationManagerSlasherMigrated)
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
		it.Event = new(AllocationManagerSlasherMigrated)
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
func (it *AllocationManagerSlasherMigratedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AllocationManagerSlasherMigratedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AllocationManagerSlasherMigrated represents a SlasherMigrated event raised by the AllocationManager contract.
type AllocationManagerSlasherMigrated struct {
	OperatorSet OperatorSet
	Slasher     common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSlasherMigrated is a free log retrieval operation binding the contract event 0xf0c8fc7d71f647bd3a88ac369112517f6a4b8038e71913f2d20f71f877dfc725.
//
// Solidity: event SlasherMigrated((address,uint32) operatorSet, address slasher)
func (_AllocationManager *AllocationManagerFilterer) FilterSlasherMigrated(opts *bind.FilterOpts) (*AllocationManagerSlasherMigratedIterator, error) {

	logs, sub, err := _AllocationManager.contract.FilterLogs(opts, "SlasherMigrated")
	if err != nil {
		return nil, err
	}
	return &AllocationManagerSlasherMigratedIterator{contract: _AllocationManager.contract, event: "SlasherMigrated", logs: logs, sub: sub}, nil
}

// WatchSlasherMigrated is a free log subscription operation binding the contract event 0xf0c8fc7d71f647bd3a88ac369112517f6a4b8038e71913f2d20f71f877dfc725.
//
// Solidity: event SlasherMigrated((address,uint32) operatorSet, address slasher)
func (_AllocationManager *AllocationManagerFilterer) WatchSlasherMigrated(opts *bind.WatchOpts, sink chan<- *AllocationManagerSlasherMigrated) (event.Subscription, error) {

	logs, sub, err := _AllocationManager.contract.WatchLogs(opts, "SlasherMigrated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AllocationManagerSlasherMigrated)
				if err := _AllocationManager.contract.UnpackLog(event, "SlasherMigrated", log); err != nil {
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

// ParseSlasherMigrated is a log parse operation binding the contract event 0xf0c8fc7d71f647bd3a88ac369112517f6a4b8038e71913f2d20f71f877dfc725.
//
// Solidity: event SlasherMigrated((address,uint32) operatorSet, address slasher)
func (_AllocationManager *AllocationManagerFilterer) ParseSlasherMigrated(log types.Log) (*AllocationManagerSlasherMigrated, error) {
	event := new(AllocationManagerSlasherMigrated)
	if err := _AllocationManager.contract.UnpackLog(event, "SlasherMigrated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AllocationManagerSlasherUpdatedIterator is returned from FilterSlasherUpdated and is used to iterate over the raw logs and unpacked data for SlasherUpdated events raised by the AllocationManager contract.
type AllocationManagerSlasherUpdatedIterator struct {
	Event *AllocationManagerSlasherUpdated // Event containing the contract specifics and raw log

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
func (it *AllocationManagerSlasherUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AllocationManagerSlasherUpdated)
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
		it.Event = new(AllocationManagerSlasherUpdated)
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
func (it *AllocationManagerSlasherUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AllocationManagerSlasherUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AllocationManagerSlasherUpdated represents a SlasherUpdated event raised by the AllocationManager contract.
type AllocationManagerSlasherUpdated struct {
	OperatorSet OperatorSet
	Slasher     common.Address
	EffectBlock uint32
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSlasherUpdated is a free log retrieval operation binding the contract event 0x3873f29d7a65a4d75f5ba28909172f486216a1420e77c3c2720815951a6b4f57.
//
// Solidity: event SlasherUpdated((address,uint32) operatorSet, address slasher, uint32 effectBlock)
func (_AllocationManager *AllocationManagerFilterer) FilterSlasherUpdated(opts *bind.FilterOpts) (*AllocationManagerSlasherUpdatedIterator, error) {

	logs, sub, err := _AllocationManager.contract.FilterLogs(opts, "SlasherUpdated")
	if err != nil {
		return nil, err
	}
	return &AllocationManagerSlasherUpdatedIterator{contract: _AllocationManager.contract, event: "SlasherUpdated", logs: logs, sub: sub}, nil
}

// WatchSlasherUpdated is a free log subscription operation binding the contract event 0x3873f29d7a65a4d75f5ba28909172f486216a1420e77c3c2720815951a6b4f57.
//
// Solidity: event SlasherUpdated((address,uint32) operatorSet, address slasher, uint32 effectBlock)
func (_AllocationManager *AllocationManagerFilterer) WatchSlasherUpdated(opts *bind.WatchOpts, sink chan<- *AllocationManagerSlasherUpdated) (event.Subscription, error) {

	logs, sub, err := _AllocationManager.contract.WatchLogs(opts, "SlasherUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AllocationManagerSlasherUpdated)
				if err := _AllocationManager.contract.UnpackLog(event, "SlasherUpdated", log); err != nil {
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

// ParseSlasherUpdated is a log parse operation binding the contract event 0x3873f29d7a65a4d75f5ba28909172f486216a1420e77c3c2720815951a6b4f57.
//
// Solidity: event SlasherUpdated((address,uint32) operatorSet, address slasher, uint32 effectBlock)
func (_AllocationManager *AllocationManagerFilterer) ParseSlasherUpdated(log types.Log) (*AllocationManagerSlasherUpdated, error) {
	event := new(AllocationManagerSlasherUpdated)
	if err := _AllocationManager.contract.UnpackLog(event, "SlasherUpdated", log); err != nil {
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
