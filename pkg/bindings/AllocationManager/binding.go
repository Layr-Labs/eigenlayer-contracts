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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_allocationManagerView\",\"type\":\"address\",\"internalType\":\"contractIAllocationManagerView\"},{\"name\":\"_delegation\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"},{\"name\":\"_eigenStrategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"_permissionController\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"},{\"name\":\"_DEALLOCATION_DELAY\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_ALLOCATION_CONFIGURATION_DELAY\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_version\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ALLOCATION_CONFIGURATION_DELAY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEALLOCATION_DELAY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addStrategiesToOperatorSet\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"clearDeallocationQueue\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"numToClear\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createOperatorSets\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"params\",\"type\":\"tuple[]\",\"internalType\":\"structIAllocationManagerTypes.CreateSetParams[]\",\"components\":[{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createOperatorSets\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"params\",\"type\":\"tuple[]\",\"internalType\":\"structIAllocationManagerTypes.CreateSetParamsV2[]\",\"components\":[{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"slasher\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createRedistributingOperatorSets\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"params\",\"type\":\"tuple[]\",\"internalType\":\"structIAllocationManagerTypes.CreateSetParams[]\",\"components\":[{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}]},{\"name\":\"redistributionRecipients\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createRedistributingOperatorSets\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"params\",\"type\":\"tuple[]\",\"internalType\":\"structIAllocationManagerTypes.CreateSetParamsV2[]\",\"components\":[{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"slasher\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"redistributionRecipients\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregisterFromOperatorSets\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIAllocationManagerTypes.DeregisterParams\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eigenStrategy\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAVSRegistrar\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAVSRegistrar\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocatableMagnitude\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"allocatableMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocatedSets\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"operatorSets\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocatedStake\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[{\"name\":\"slashableStake\",\"type\":\"uint256[][]\",\"internalType\":\"uint256[][]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocatedStrategies\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocation\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"allocation\",\"type\":\"tuple\",\"internalType\":\"structIAllocationManagerTypes.Allocation\",\"components\":[{\"name\":\"currentMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"pendingDiff\",\"type\":\"int128\",\"internalType\":\"int128\"},{\"name\":\"effectBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocationDelay\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocations\",\"inputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"allocations\",\"type\":\"tuple[]\",\"internalType\":\"structIAllocationManagerTypes.Allocation[]\",\"components\":[{\"name\":\"currentMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"pendingDiff\",\"type\":\"int128\",\"internalType\":\"int128\"},{\"name\":\"effectBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getEncumberedMagnitude\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"encumberedMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxMagnitude\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"maxMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxMagnitudes\",\"inputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"maxMagnitudes\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxMagnitudes\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[{\"name\":\"maxMagnitudes\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxMagnitudesAtBlock\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"maxMagnitudes\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMemberCount\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"memberCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMembers\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMinimumSlashableStake\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"slashableStake\",\"type\":\"uint256[][]\",\"internalType\":\"uint256[][]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorSetCount\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPendingSlasher\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"pendingSlasher\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"effectBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRedistributionRecipient\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRegisteredSets\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"operatorSets\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSlashCount\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"slashCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSlasher\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStrategiesInOperatorSet\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStrategyAllocations\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"operatorSets\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"allocations\",\"type\":\"tuple[]\",\"internalType\":\"structIAllocationManagerTypes.Allocation[]\",\"components\":[{\"name\":\"currentMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"pendingDiff\",\"type\":\"int128\",\"internalType\":\"int128\"},{\"name\":\"effectBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isMemberOfOperatorSet\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"result\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperatorRedistributable\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperatorSet\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"result\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperatorSlashable\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRedistributingOperatorSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"migrateSlashers\",\"inputs\":[{\"name\":\"operatorSets\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"modifyAllocations\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"params\",\"type\":\"tuple[]\",\"internalType\":\"structIAllocationManagerTypes.AllocateParams[]\",\"components\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"newMagnitudes\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"permissionController\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerForOperatorSets\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIAllocationManagerTypes.RegisterParams\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeStrategiesFromOperatorSet\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAVSRegistrar\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"registrar\",\"type\":\"address\",\"internalType\":\"contractIAVSRegistrar\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAllocationDelay\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delay\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slashOperator\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIAllocationManagerTypes.SlashingParams\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"wadsToSlash\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAVSMetadataURI\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateSlasher\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slasher\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AVSMetadataURIUpdated\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"metadataURI\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AVSRegistrarSet\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"registrar\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIAVSRegistrar\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllocationDelaySet\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"delay\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"effectBlock\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllocationUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"magnitude\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"effectBlock\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EncumberedMagnitudeUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"encumberedMagnitude\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MaxMagnitudeUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"maxMagnitude\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorAddedToOperatorSet\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRemovedFromOperatorSet\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSetCreated\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSlashed\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategies\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"contractIStrategy[]\"},{\"name\":\"wadSlashed\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"description\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RedistributionAddressSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"redistributionRecipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SlasherMigrated\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slasher\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SlasherUpdated\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slasher\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"effectBlock\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyAddedToOperatorSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyRemovedFromOperatorSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyMemberOfSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CurrentlyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Empty\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputArrayLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientMagnitude\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAVSRegistrar\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCaller\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNewPausedStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperatorSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPermissions\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRedistributionRecipient\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSnapshotOrdering\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidStrategy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidWadToSlash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ModificationAlreadyPending\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonexistentAVSMetadata\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotMemberOfSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyPauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnpauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorNotSlashable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorSetAlreadyMigrated\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SameMagnitude\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SlasherNotSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategiesMustBeInAscendingOrder\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategyAlreadyInOperatorSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategyNotInOperatorSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"UninitializedAllocationDelay\",\"inputs\":[]}]",
	Bin: "0x610180604052348015610010575f5ffd5b50604051615cf9380380615cf983398101604081905261002f916101fe565b878185898987878b6001600160a01b03811661005e576040516339b190bb60e11b815260040160405180910390fd5b6001600160a01b0390811660805293841660a05291831660c05263ffffffff90811660e05216610100521661012052610096816100bd565b61014052506001600160a01b0316610160526100b0610103565b505050505050505061038a565b5f5f829050601f815111156100f0578260405163305a27a960e01b81526004016100e7919061032f565b60405180910390fd5b80516100fb82610364565b179392505050565b5f54610100900460ff161561016a5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b60648201526084016100e7565b5f5460ff908116146101b9575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146101cf575f5ffd5b50565b805163ffffffff811681146101e5575f5ffd5b919050565b634e487b7160e01b5f52604160045260245ffd5b5f5f5f5f5f5f5f5f610100898b031215610216575f5ffd5b8851610221816101bb565b60208a0151909850610232816101bb565b60408a0151909750610243816101bb565b60608a0151909650610254816101bb565b60808a0151909550610265816101bb565b935061027360a08a016101d2565b925061028160c08a016101d2565b60e08a01519092506001600160401b0381111561029c575f5ffd5b8901601f81018b136102ac575f5ffd5b80516001600160401b038111156102c5576102c56101ea565b604051601f8201601f19908116603f011681016001600160401b03811182821017156102f3576102f36101ea565b6040528181528282016020018d101561030a575f5ffd5b8160208401602083015e5f602083830101528093505050509295985092959890939650565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b80516020808301519190811015610384575f198160200360031b1b821691505b50919050565b60805160a05160c05160e051610100516101205161014051610160516158816104785f395f818161092d01528181610a0d01528181610a3c01528181610aa301528181610ace01528181610d640152818161107201528181612234015261254901525f610f2301525f818161055901528181611a7301526128a101525f81816106a501528181613675015261386301525f8181610454015281816113bf01526117fd01525f818161083b01526134ca01525f818161087001528181610f5701528181610fc201528181611d4101526130ba01525f81816106cc015281816127980152613d9501526158815ff3fe608060405234801561000f575f5ffd5b506004361061037d575f3560e01c80636cfb4481116101d4578063b66bd98911610109578063d7794857116100a9578063f231bd0811610079578063f231bd0814610892578063f605ce0814610667578063fabc1cbc146108a5578063fe4b84df146108b8575f5ffd5b8063d7794857146107ff578063db4df76114610836578063dc2af6921461085d578063df5cf7231461086b575f5ffd5b8063c221d8ae116100e4578063c221d8ae146107b3578063d1a83f54146107c6578063d3d96ff4146107d9578063d4a3fcce146107ec575f5ffd5b8063b66bd98914610763578063b9fbaed114610776578063ba1a84e5146107a5575f5ffd5b806394d7d00c11610174578063a9333ec81161014f578063a9333ec814610667578063a98218211461073d578063adc2e3d914610750578063b2447af7146105a9575f5ffd5b806394d7d00c14610709578063952899ee14610717578063957dc50b1461072a575f5ffd5b806379ae50cd116101af57806379ae50cd146104095780637bc1ef61146106a0578063886f1195146106c75780638ce64854146106ee575f5ffd5b80636cfb4481146106675780636e3492b51461068d5780636e875dba14610539575f5ffd5b806340120dab116102b5578063547afb87116102555780635ac86ab7116102255780635ac86ab7146106165780635c975abb14610639578063670d3ba2146106415780636c9d7c5814610654575f5ffd5b8063547afb87146105d857806354fd4d50146105e657806356c483e6146105fb578063595c6a671461060e575f5ffd5b80634a10ffe5116102905780634a10ffe51461057b5780634b5046ef146105965780634cfd2939146105a957806350feea20146105c5575f5ffd5b806340120dab146105185780634177a87c146105395780634657e26a14610554575f5ffd5b80632981eb7711610320578063304c10cd116102fb578063304c10cd146104be57806332a879e4146104d157806336352057146104e45780633dff8e7d14610505575f5ffd5b80632981eb771461044f5780632b453a9a1461048b5780632bab2c4a146104ab575f5ffd5b8063136439dd1161035b578063136439dd146103f457806315fe502814610409578063260dc75814610429578063261f84e01461043c575f5ffd5b80630f3df50e1461038157806310e1b9b8146103b15780631352c3e6146103d1575b5f5ffd5b61039461038f366004614564565b6108cb565b6040516001600160a01b0390911681526020015b60405180910390f35b6103c46103bf36600461457e565b61090c565b6040516103a891906145c5565b6103e46103df3660046145f8565b610951565b60405190151581526020016103a8565b61040761040236600461462c565b6109cc565b005b61041c610417366004614643565b610a06565b6040516103a891906146c1565b6103e4610437366004614564565b610a36565b61040761044a366004614713565b610a60565b6104767f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff90911681526020016103a8565b61049e6104993660046147f8565b610a9c565b6040516103a8919061489b565b61049e6104b93660046148fe565b610ac7565b6103946104cc366004614643565b610afa565b6104076104df366004614982565b610b29565b6104f76104f2366004614a02565b610b6e565b6040516103a8929190614a54565b610407610513366004614b4a565b610cc3565b61052b610526366004614b96565b610d5c565b6040516103a8929190614c23565b610547610417366004614564565b6040516103a89190614c47565b6103947f000000000000000000000000000000000000000000000000000000000000000081565b610589610499366004614c92565b6040516103a89190614cd9565b6104076105a4366004614982565b610d8f565b6105b7610437366004614564565b6040519081526020016103a8565b6104076105d3366004614d19565b610e2b565b610589610499366004614713565b6105ee610f1c565b6040516103a89190614d77565b610407610609366004614dac565b610f4c565b610407611058565b6103e4610624366004614dd6565b606654600160ff9092169190911b9081161490565b6066546105b7565b6103e461064f3660046145f8565b61106c565b610407610662366004614df6565b611096565b61067561064f366004614b96565b6040516001600160401b0390911681526020016103a8565b61040761069b366004614e37565b611149565b6104767f000000000000000000000000000000000000000000000000000000000000000081565b6103947f000000000000000000000000000000000000000000000000000000000000000081565b6106fc610499366004614e68565b6040516103a89190614eab565b6105896104b9366004614ebd565b610407610725366004614f14565b611504565b6104076107383660046150bd565b6119b2565b61040761074b36600461514d565b611c3e565b61040761075e3660046151cb565b611cf0565b610407610771366004614d19565b612039565b610789610784366004614643565b612193565b60408051921515835263ffffffff9091166020830152016103a8565b6105b7610437366004614643565b6105476107c13660046145f8565b61222d565b6104076107d436600461520d565b612258565b6104076107e7366004614b96565b61238e565b6103946107fa366004614564565b6124bb565b61081261080d366004614564565b612542565b604080516001600160a01b03909316835263ffffffff9091166020830152016103a8565b6103947f000000000000000000000000000000000000000000000000000000000000000081565b6103e4610437366004614643565b6103947f000000000000000000000000000000000000000000000000000000000000000081565b6103e46108a0366004614564565b612572565b6104076108b336600461462c565b612591565b6104076108c636600461462c565b6125fe565b5f5f60a65f6108d98561270f565b815260208101919091526040015f20546001600160a01b0316905080156109005780610905565b620e16e45b9392505050565b604080516060810182525f80825260208201819052918101919091526109057f0000000000000000000000000000000000000000000000000000000000000000612772565b6001600160a01b0382165f908152609e602052604081208190816109748561270f565b815260208082019290925260409081015f2081518083019092525460ff8116151580835261010090910463ffffffff16928201929092529150806109c25750806020015163ffffffff164311155b9150505b92915050565b6109d4612783565b60665481811681146109f95760405163c61dca5d60e01b815260040160405180910390fd5b610a0282612826565b5050565b6060610a317f0000000000000000000000000000000000000000000000000000000000000000612772565b919050565b5f610a317f0000000000000000000000000000000000000000000000000000000000000000612772565b82610a6a81612863565b610a875760405163932d94f760e01b815260040160405180910390fd5b610a968461051385858861290d565b50505050565b60606109057f0000000000000000000000000000000000000000000000000000000000000000612772565b6060610af27f0000000000000000000000000000000000000000000000000000000000000000612772565b949350505050565b6001600160a01b038082165f908152609760205260408120549091168015610b225780610905565b5090919050565b84610b3381612863565b610b505760405163932d94f760e01b815260040160405180910390fd5b610b6686610b5f87878a61290d565b8585612258565b505050505050565b5f60606001610b7c81612a84565b5f6040518060400160405280876001600160a01b03168152602001866020016020810190610baa919061526c565b63ffffffff1690529050610bc16060860186615285565b9050610bd06040870187615285565b905014610bf0576040516343714afd60e01b815260040160405180910390fd5b60208082015182516001600160a01b03165f90815260989092526040909120610c229163ffffffff90811690612ab216565b610c3f57604051631fb1705560e21b815260040160405180910390fd5b610c55610c4f6020870187614643565b82610951565b610c725760405163ebbff49760e01b815260040160405180910390fd5b610c7b816124bb565b6001600160a01b0316336001600160a01b031614610cac576040516348f5c3ed60e01b815260040160405180910390fd5b610cb68582612ac9565b9350935050509250929050565b81610ccd81612863565b610cea5760405163932d94f760e01b815260040160405180910390fd5b6001600160a01b0383165f90815260a4602052604090205460ff16610d22576040516348f7dbb960e01b815260040160405180910390fd5b5f5b8251811015610a9657610d5484848381518110610d4357610d436152ca565b6020026020010151620e16e4613224565b600101610d24565b606080610d887f0000000000000000000000000000000000000000000000000000000000000000612772565b9250929050565b5f610d9981612a84565b838214610db9576040516343714afd60e01b815260040160405180910390fd5b5f5b84811015610e2257610e1a87878784818110610dd957610dd96152ca565b9050602002016020810190610dee9190614643565b868685818110610e0057610e006152ca565b9050602002016020810190610e1591906152de565b613396565b600101610dbb565b50505050505050565b83610e3581612863565b610e525760405163932d94f760e01b815260040160405180910390fd5b6040805180820182526001600160a01b03871680825263ffffffff80881660208085018290525f93845260989052939091209192610e919291612ab216565b610eae57604051631fb1705560e21b815260040160405180910390fd5b5f5b83811015610e2257610f1482868684818110610ece57610ece6152ca565b9050602002016020810190610ee39190614643565b610f0f60405180604001604052808c6001600160a01b031681526020018b63ffffffff16815250612572565b61349a565b600101610eb0565b6060610f477f000000000000000000000000000000000000000000000000000000000000000061357a565b905090565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148061104857610f8683612863565b610fa3576040516348f5c3ed60e01b815260040160405180910390fd5b6040516336b87bd760e11b81526001600160a01b0384811660048301527f00000000000000000000000000000000000000000000000000000000000000001690636d70f7ae90602401602060405180830381865afa158015611007573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061102b91906152ff565b6110485760405163ccea9e6f60e01b815260040160405180910390fd5b6110538383836135b7565b505050565b611060612783565b61106a5f19612826565b565b5f6109c67f0000000000000000000000000000000000000000000000000000000000000000612772565b81516110a181612863565b6110be5760405163932d94f760e01b815260040160405180910390fd5b60208084015184516001600160a01b03165f908152609890925260409091206110f09163ffffffff90811690612ab216565b61110d57604051631fb1705560e21b815260040160405180910390fd5b5f611117846124bb565b6001600160a01b03160361113e5760405163255b0f4160e01b815260040160405180910390fd5b61105383835f613787565b600261115481612a84565b6111696111646020840184614643565b612863565b8061118257506111826111646040840160208501614643565b61119f576040516348f5c3ed60e01b815260040160405180910390fd5b5f5b6111ae6040840184615285565b9050811015611470575f60405180604001604052808560200160208101906111d69190614643565b6001600160a01b031681526020016111f16040870187615285565b85818110611201576112016152ca565b9050602002016020810190611216919061526c565b63ffffffff168152509050611263816020015163ffffffff1660985f8760200160208101906112459190614643565b6001600160a01b0316815260208101919091526040015f2090612ab2565b61128057604051631fb1705560e21b815260040160405180910390fd5b609e5f6112906020870187614643565b6001600160a01b03166001600160a01b031681526020019081526020015f205f6112b98361270f565b815260208101919091526040015f205460ff166112e9576040516325131d4f60e01b815260040160405180910390fd5b6113236112f58261270f565b609c5f6113056020890189614643565b6001600160a01b0316815260208101919091526040015f2090613943565b5061135b6113346020860186614643565b609a5f6113408561270f565b81526020019081526020015f2061394e90919063ffffffff16565b506113696020850185614643565b6001600160a01b03167fad34c3070be1dffbcaa499d000ba2b8d9848aefcac3059df245dd95c4ece14fe826040516113a1919061531e565b60405180910390a2604080518082019091525f8152602081016113e47f000000000000000000000000000000000000000000000000000000000000000043615340565b63ffffffff169052609e5f6113fc6020880188614643565b6001600160a01b03166001600160a01b031681526020019081526020015f205f6114258461270f565b81526020808201929092526040015f2082518154939092015163ffffffff166101000264ffffffff00199215159290921664ffffffffff1990931692909217179055506001016111a1565b506114846104cc6040840160208501614643565b6001600160a01b031663303ca95661149f6020850185614643565b6114af6040860160208701614643565b6114bc6040870187615285565b6040518563ffffffff1660e01b81526004016114db9493929190615395565b5f604051808303815f87803b1580156114f2575f5ffd5b505af1158015610b66573d5f5f3e3d5ffd5b5f61150e81612a84565b61151783612863565b611534576040516348f5c3ed60e01b815260040160405180910390fd5b5f5f5f61154086612193565b91509150816115625760405163fa55fc8160e01b815260040160405180910390fd5b91505f90505b83518110156119ab57838181518110611583576115836152ca565b602002602001015160400151518482815181106115a2576115a26152ca565b60200260200101516020015151146115cd576040516343714afd60e01b815260040160405180910390fd5b5f8482815181106115e0576115e06152ca565b602090810291909101810151518082015181516001600160a01b03165f908152609890935260409092209092506116209163ffffffff90811690612ab216565b61163d57604051631fb1705560e21b815260040160405180910390fd5b5f6116488783610951565b90505f5b86848151811061165e5761165e6152ca565b602002602001015160200151518110156119a0575f878581518110611685576116856152ca565b60200260200101516020015182815181106116a2576116a26152ca565b602002602001015190506116b9898261ffff613396565b5f5f6116ce8b6116c88861270f565b85613962565b91509150806040015163ffffffff165f146116fc57604051630d8fcbe360e41b815260040160405180910390fd5b5f61170987858489613ace565b905061174e825f01518c8a81518110611724576117246152ca565b6020026020010151604001518781518110611741576117416152ca565b6020026020010151613b06565b600f0b602083018190525f0361177757604051634606179360e11b815260040160405180910390fd5b5f8260200151600f0b12156118bb57801561183d576117f86117988861270f565b6001600160a01b03808f165f90815260a360209081526040808320938a16835292905220908154600160801b90819004600f0b5f818152600180860160205260409091209390935583546001600160801b03908116939091011602179055565b6118227f000000000000000000000000000000000000000000000000000000000000000043615340565b61182d906001615340565b63ffffffff166040830152611928565b61184f83602001518360200151613b1d565b6001600160401b031660208401528a518b9089908110611871576118716152ca565b602002602001015160400151858151811061188e5761188e6152ca565b6020908102919091018101516001600160401b031683525f9083015263ffffffff43166040830152611928565b5f8260200151600f0b1315611928576118dc83602001518360200151613b1d565b6001600160401b03908116602085018190528451909116101561191257604051636c9be0bf60e01b815260040160405180910390fd5b61191c8943615340565b63ffffffff1660408301525b61193d8c6119358961270f565b868686613b3c565b7f1487af5418c47ee5ea45ef4a93398668120890774a9e13487e61e9dc3baf76dd8c8886611972865f01518760200151613b1d565b86604001516040516119889594939291906153c1565b60405180910390a150506001909201915061164c9050565b505050600101611568565b5050505050565b5f5b8151811015610a0257611a318282815181106119d2576119d26152ca565b60200260200101516020015163ffffffff1660985f8585815181106119f9576119f96152ca565b60200260200101515f01516001600160a01b03166001600160a01b031681526020019081526020015f20612ab290919063ffffffff16565b15611c36575f6001600160a01b0316611a62838381518110611a5557611a556152ca565b60200260200101516124bb565b6001600160a01b031603611c36575f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663fddbdefd848481518110611ab257611ab26152ca565b6020908102919091010151516040516001600160e01b031960e084901b1681526001600160a01b039091166004820152306024820152633635205760e01b60448201526064015f604051808303815f875af1158015611b13573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052611b3a9190810190615412565b90505f81515f1480611b7657505f6001600160a01b0316825f81518110611b6357611b636152ca565b60200260200101516001600160a01b0316145b15611b9f57838381518110611b8d57611b8d6152ca565b60200260200101515f01519050611bbc565b815f81518110611bb157611bb16152ca565b602002602001015190505b611be1848481518110611bd157611bd16152ca565b6020026020010151826001613787565b7ff0c8fc7d71f647bd3a88ac369112517f6a4b8038e71913f2d20f71f877dfc725848481518110611c1457611c146152ca565b602002602001015182604051611c2b9291906154a1565b60405180910390a150505b6001016119b4565b82611c4881612863565b611c655760405163932d94f760e01b815260040160405180910390fd5b6001600160a01b0384165f90815260a4602052604090205460ff16611ca7576001600160a01b0384165f90815260a460205260409020805460ff191660011790555b836001600160a01b03167fa89c1dc243d8908a96dd84944bcc97d6bc6ac00dd78e20621576be6a3c9437138484604051611ce29291906154ef565b60405180910390a250505050565b6002611cfb81612a84565b82611d0581612863565b611d225760405163932d94f760e01b815260040160405180910390fd5b6040516336b87bd760e11b81526001600160a01b0385811660048301527f00000000000000000000000000000000000000000000000000000000000000001690636d70f7ae90602401602060405180830381865afa158015611d86573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611daa91906152ff565b611dc75760405163ccea9e6f60e01b815260040160405180910390fd5b5f5b611dd66020850185615285565b9050811015611f9e57604080518082019091525f9080611df96020880188614643565b6001600160a01b03168152602001868060200190611e179190615285565b85818110611e2757611e276152ca565b9050602002016020810190611e3c919061526c565b63ffffffff90811690915260208083015183516001600160a01b03165f90815260989092526040909120929350611e78929190811690612ab216565b611e9557604051631fb1705560e21b815260040160405180910390fd5b611e9f8682610951565b15611ebd57604051636c6c6e2760e11b815260040160405180910390fd5b611ee6611ec98261270f565b6001600160a01b0388165f908152609c6020526040902090613d74565b50611f1286609a5f611ef78561270f565b81526020019081526020015f20613d7f90919063ffffffff16565b50856001600160a01b03167f43232edf9071753d2321e5fa7e018363ee248e5f2142e6c08edd3265bfb4895e82604051611f4c919061531e565b60405180910390a26001600160a01b0386165f908152609e60205260408120600191611f778461270f565b815260208101919091526040015f20805460ff191691151591909117905550600101611dc9565b50611faf6104cc6020850185614643565b6001600160a01b031663c63fd50285611fcb6020870187614643565b611fd86020880188615285565b611fe560408a018a615502565b6040518763ffffffff1660e01b815260040161200696959493929190615544565b5f604051808303815f87803b15801561201d575f5ffd5b505af115801561202f573d5f5f3e3d5ffd5b5050505050505050565b8361204381612863565b6120605760405163932d94f760e01b815260040160405180910390fd5b6040805180820182526001600160a01b03871680825263ffffffff80881660208085018290525f9384526098905293909120919261209f9291612ab216565b6120bc57604051631fb1705560e21b815260040160405180910390fd5b5f6120c68261270f565b90505f5b8481101561202f5761210f8686838181106120e7576120e76152ca565b90506020020160208101906120fc9190614643565b5f8481526099602052604090209061394e565b61212c576040516331bc342760e11b815260040160405180910390fd5b7f7b4b073d80dcac55a11177d8459ad9f664ceeb91f71f27167bb14f8152a7eeee83878784818110612160576121606152ca565b90506020020160208101906121759190614643565b6040516121839291906154a1565b60405180910390a16001016120ca565b6001600160a01b0381165f908152609b602090815260408083208151608081018352905463ffffffff80821680845260ff600160201b8404161515958401869052650100000000008304821694840194909452600160481b9091041660608201819052849391929190158015906122145750826060015163ffffffff164310155b15612223575050604081015160015b9590945092505050565b60606109c67f0000000000000000000000000000000000000000000000000000000000000000612772565b8361226281612863565b61227f5760405163932d94f760e01b815260040160405180910390fd5b835182146122a0576040516343714afd60e01b815260040160405180910390fd5b6001600160a01b0385165f90815260a4602052604090205460ff166122d8576040516348f7dbb960e01b815260040160405180910390fd5b5f5b8451811015610b66575f8484838181106122f6576122f66152ca565b905060200201602081019061230b9190614643565b90506001600160a01b038116612334576040516339b190bb60e11b815260040160405180910390fd5b620e16e3196001600160a01b03821601612361576040516364be1a3f60e11b815260040160405180910390fd5b61238587878481518110612377576123776152ca565b602002602001015183613224565b506001016122da565b8161239881612863565b6123b55760405163932d94f760e01b815260040160405180910390fd5b60405163b526578760e01b81526001600160a01b03848116600483015283169063b526578790602401602060405180830381865afa1580156123f9573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061241d91906152ff565b61243a57604051631d0b13c160e31b815260040160405180910390fd5b6001600160a01b038381165f90815260976020526040902080546001600160a01b0319169184169190911790557f2ae945c40c44dc0ec263f95609c3fdc6952e0aefa22d6374e44f2c997acedf858361249281610afa565b604080516001600160a01b039384168152929091166020830152015b60405180910390a1505050565b5f5f60a75f6124c98561270f565b815260208082019290925260409081015f20815160608101835281546001600160a01b0390811680835260019093015490811694820194909452600160a01b90930463ffffffff16918301829052919250158015906125325750816040015163ffffffff164310155b1561090557506020015192915050565b5f5f61256d7f0000000000000000000000000000000000000000000000000000000000000000612772565b915091565b5f620e16e4612580836108cb565b6001600160a01b0316141592915050565b612599613d93565b606654801982198116146125c05760405163c61dca5d60e01b815260040160405180910390fd5b606682905560405182815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c9060200160405180910390a25050565b5f54610100900460ff161580801561261c57505f54600160ff909116105b806126355750303b15801561263557505f5460ff166001145b61269d5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b5f805460ff1916600117905580156126be575f805461ff0019166101001790555b6126c782612826565b8015610a02575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15050565b5f815f0151826020015163ffffffff1660405160200161275a92919060609290921b6bffffffffffffffffffffffff1916825260a01b6001600160a01b031916601482015260200190565b6040516020818303038152906040526109c690615590565b613e44806110538363ffffffff8316565b60405163237dfb4760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906346fbf68e90602401602060405180830381865afa1580156127e5573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061280991906152ff565b61106a57604051631d77d47760e21b815260040160405180910390fd5b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a250565b604051631beb2b9760e31b81526001600160a01b0382811660048301523360248301523060448301525f80356001600160e01b0319166064840152917f00000000000000000000000000000000000000000000000000000000000000009091169063df595cb8906084016020604051808303815f875af11580156128e9573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109c691906152ff565b60605f836001600160401b0381111561292857612928614474565b60405190808252806020026020018201604052801561298a57816020015b61297760405180606001604052805f63ffffffff168152602001606081526020015f6001600160a01b031681525090565b8152602001906001900390816129465790505b5090505f5b84811015612a7b5760405180606001604052808787848181106129b4576129b46152ca565b90506020028101906129c691906155b3565b6129d490602081019061526c565b63ffffffff1681526020018787848181106129f1576129f16152ca565b9050602002810190612a0391906155b3565b612a11906020810190615285565b808060200260200160405190810160405280939291908181526020018383602002808284375f920191909152505050908252506001600160a01b0386166020909101528251839083908110612a6857612a686152ca565b602090810291909101015260010161298f565b50949350505050565b606654600160ff83161b90811603612aaf5760405163840a48d560e01b815260040160405180910390fd5b50565b5f8181526001830160205260408120541515610905565b5f606081612ada6040860186615285565b90506001600160401b03811115612af357612af3614474565b604051908082528060200260200182016040528015612b1c578160200160208202803683370190505b509050612b2c6040860186615285565b90506001600160401b03811115612b4557612b45614474565b604051908082528060200260200182016040528015612b6e578160200160208202803683370190505b50915060a55f612b7d8661270f565b81526020019081526020015f205f8154612b96906155d1565b918290555092505f5b612bac6040870187615285565b90508110156131b657801580612c3f5750612bca6040870187615285565b612bd56001846155e9565b818110612be457612be46152ca565b9050602002016020810190612bf99190614643565b6001600160a01b0316612c0f6040880188615285565b83818110612c1f57612c1f6152ca565b9050602002016020810190612c349190614643565b6001600160a01b0316115b612c5c57604051639f1c805360e01b815260040160405180910390fd5b612c696060870187615285565b82818110612c7957612c796152ca565b905060200201355f108015612cb95750670de0b6b3a7640000612c9f6060880188615285565b83818110612caf57612caf6152ca565b9050602002013511155b612cd657604051631353603160e01b815260040160405180910390fd5b612d32612ce66040880188615285565b83818110612cf657612cf66152ca565b9050602002016020810190612d0b9190614643565b60995f612d178961270f565b81526020019081526020015f20613e6290919063ffffffff16565b612d4f576040516331bc342760e11b815260040160405180910390fd5b5f80612da1612d6160208a018a614643565b612d6a8961270f565b612d7760408c018c615285565b87818110612d8757612d876152ca565b9050602002016020810190612d9c9190614643565b613962565b805191935091506001600160401b03165f03612dbe5750506131ae565b5f612df9612dcf60608b018b615285565b86818110612ddf57612ddf6152ca565b85516001600160401b031692602090910201359050613e83565b8351909150612e146001600160401b03808416908316613e99565b868681518110612e2657612e266152ca565b60200260200101818152505081835f01818151612e4391906155fc565b6001600160401b0316905250835182908590612e609083906155fc565b6001600160401b0316905250602084018051839190612e809083906155fc565b6001600160401b031690525060208301515f600f9190910b1215612f98575f612ee3612eaf60608d018d615285565b88818110612ebf57612ebf6152ca565b905060200201358560200151612ed49061561b565b6001600160801b031690613e83565b9050806001600160401b031684602001818151612f00919061563f565b600f0b9052507f1487af5418c47ee5ea45ef4a93398668120890774a9e13487e61e9dc3baf76dd612f3460208d018d614643565b8b612f4260408f018f615285565b8a818110612f5257612f526152ca565b9050602002016020810190612f679190614643565b612f78885f01518960200151613b1d565b8860400151604051612f8e9594939291906153c1565b60405180910390a1505b612fea612fa860208c018c614643565b612fb18b61270f565b612fbe60408e018e615285565b89818110612fce57612fce6152ca565b9050602002016020810190612fe39190614643565b8787613b3c565b7f1487af5418c47ee5ea45ef4a93398668120890774a9e13487e61e9dc3baf76dd61301860208c018c614643565b8a61302660408e018e615285565b89818110613036576130366152ca565b905060200201602081019061304b9190614643565b865160405161305f949392919043906153c1565b60405180910390a16130b061307760208c018c614643565b61308460408d018d615285565b88818110613094576130946152ca565b90506020020160208101906130a99190614643565b8651613ead565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016635ae679a76130ec60208d018d614643565b8b8b8e80604001906130fe9190615285565b8b81811061310e5761310e6152ca565b90506020020160208101906131239190614643565b89516040516001600160e01b031960e088901b16815261314b9594939291899160040161566c565b6020604051808303815f875af1158015613167573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061318b91906156bf565b87868151811061319d5761319d6152ca565b602002602001018181525050505050505b600101612b9f565b507f80969ad29428d6797ee7aad084f9e4a42a82fc506dcd2ca3b6fb431f85ccebe56131e56020870187614643565b856131f36040890189615285565b8561320160808c018c615502565b60405161321497969594939291906156d6565b60405180910390a1509250929050565b6040805180820182526001600160a01b038516808252845163ffffffff90811660208085018290525f938452609890529390912091926132659291613d7416565b61328257604051631fb1705560e21b815260040160405180910390fd5b7f31629285ead2335ae0933f86ed2ae63321f7af77b4e6eaabc42c057880977e6c816040516132b1919061531e565b60405180910390a16001600160a01b038216620e16e414801590613346578260a65f6132dc8561270f565b81526020019081526020015f205f6101000a8154816001600160a01b0302191690836001600160a01b031602179055507f90a6fa2a9b79b910872ebca540cf3bd8be827f586e6420c30d8836e30012907e828460405161333d9291906154a1565b60405180910390a15b5f5b8460200151518110156133855761337d838660200151838151811061336f5761336f6152ca565b60200260200101518461349a565b600101613348565b506119ab8285604001516001613787565b6001600160a01b038381165f90815260a360209081526040808320938616835292905290812054600f81810b600160801b909204900b035b5f811180156133e057508261ffff1682105b156119ab576001600160a01b038086165f90815260a360209081526040808320938816835292905290812061341490613f2f565b90505f5f613423888489613962565b91509150806040015163ffffffff16431015613441575050506119ab565b61344e8884898585613b3c565b6001600160a01b038089165f90815260a360209081526040808320938b1683529290522061347b90613f81565b50613485856155d1565b94506134908461576c565b93505050506133ce565b801561351c576001600160a01b03821673beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac0148015906134ff57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316826001600160a01b031614155b61351c57604051632711b74d60e11b815260040160405180910390fd5b61352c8260995f611ef78761270f565b6135495760405163585cfb2f60e01b815260040160405180910390fd5b7f7ab260fe0af193db5f4986770d831bda4ea46099dc817e8b6716dcae8af8e88b83836040516124ae9291906154a1565b60605f61358683613ffe565b6040805160208082528183019092529192505f91906020820181803683375050509182525060208101929092525090565b6001600160a01b0383165f908152609b60209081526040918290208251608081018452905463ffffffff808216835260ff600160201b830416151593830193909352650100000000008104831693820193909352600160481b9092041660608201819052158015906136335750806060015163ffffffff164310155b1561364d57604081015163ffffffff168152600160208201525b63ffffffff8316604082015281156136705763ffffffff431660608201526136b1565b61369a7f000000000000000000000000000000000000000000000000000000000000000043615340565b6136a5906001615340565b63ffffffff1660608201525b6001600160a01b0384165f818152609b60209081526040918290208451815486840151878601516060808a015163ffffffff95861664ffffffffff1990951694909417600160201b93151593909302929092176cffffffffffffffff0000000000191665010000000000918516919091026cffffffff000000000000000000191617600160481b92841692830217909355845195865290881692850192909252918301527f4e85751d6331506c6c62335f207eb31f12a61e570f34f5c17640308785c6d4db91015b60405180910390a150505050565b6001600160a01b0382166137ae576040516339b190bb60e11b815260040160405180910390fd5b5f60a75f6137bb8661270f565b815260208082019290925260409081015f20815160608101835281546001600160a01b03908116825260019092015491821693810193909352600160a01b900463ffffffff16908201819052909150158015906138225750806040015163ffffffff164310155b156138385760208101516001600160a01b031681525b6001600160a01b0383166020820152811561385e5763ffffffff4316604082015261389f565b6138887f000000000000000000000000000000000000000000000000000000000000000043615340565b613893906001615340565b63ffffffff1660408201525b8060a75f6138ac8761270f565b815260208082019290925260409081015f20835181546001600160a01b039182166001600160a01b031990911617825592840151600190910180549483015163ffffffff16600160a01b026001600160c01b031990951691909316179290921790558181015190517f3873f29d7a65a4d75f5ba28909172f486216a1420e77c3c2720815951a6b4f57916137799187918791615781565b5f6109058383614025565b5f610905836001600160a01b038416614025565b6040805180820182525f80825260208083018290528351606081018552828152808201839052808501839052845180860186526001600160a01b03898116855260a18452868520908816855290925293822092939281906139c290614108565b6001600160401b0390811682526001600160a01b038981165f81815260a260209081526040808320948c168084529482528083205486169682019690965291815260a082528481208b8252825284812092815291815290839020835160608101855290549283168152600160401b8304600f0b91810191909152600160c01b90910463ffffffff16918101829052919250431015613a64579092509050613ac6565b613a75815f01518260200151613b1d565b6001600160401b0316815260208101515f600f9190910b1215613ab357613aa482602001518260200151613b1d565b6001600160401b031660208301525b5f60408201819052602082015290925090505b935093915050565b5f613adf8460995f612d178961270f565b8015613ae85750815b8015613afd575082516001600160401b031615155b95945050505050565b5f6109056001600160401b038085169084166157b4565b5f610905613b34836001600160401b03861661563f565b600f0b61411b565b6020808301516001600160a01b038088165f90815260a284526040808220928816825291909352909120546001600160401b03908116911614613c0257602082810180516001600160a01b038881165f81815260a286526040808220938a1680835293875290819020805467ffffffffffffffff19166001600160401b0395861617905593518451918252948101919091529216908201527facf9095feb3a370c9cf692421c69ef320d4db5c66e6a7d29c7694eb02364fc559060600160405180910390a15b6001600160a01b038086165f90815260a060209081526040808320888452825280832093871683529281529082902083518154928501519385015163ffffffff16600160c01b0263ffffffff60c01b196001600160801b038616600160401b026001600160c01b03199095166001600160401b03909316929092179390931716919091179055600f0b15613ce4576001600160a01b0385165f908152609f602090815260408083208784529091529020613cbc9084613d7f565b506001600160a01b0385165f908152609d60205260409020613cde9085613d74565b506119ab565b80516001600160401b03165f036119ab576001600160a01b0385165f908152609f602090815260408083208784529091529020613d21908461394e565b506001600160a01b0385165f908152609f602090815260408083208784529091529020613d4d90614186565b5f036119ab576001600160a01b0385165f908152609d60205260409020610b669085613943565b5f610905838361418f565b5f610905836001600160a01b03841661418f565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015613def573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613e1391906157e1565b6001600160a01b0316336001600160a01b03161461106a5760405163794821ff60e01b815260040160405180910390fd5b365f5f375f5f365f845af43d5f5f3e808015613e5e573d5ff35b3d5ffd5b6001600160a01b0381165f9081526001830160205260408120541515610905565b5f6109058383670de0b6b3a764000060016141db565b5f61090583670de0b6b3a764000084614234565b6001600160a01b038084165f90815260a160209081526040808320938616835292905220613edc904383614319565b604080516001600160a01b038086168252841660208201526001600160401b038316918101919091527f1c6458079a41077d003c11faf9bf097e693bd67979e4e6500bac7b29db779b5c906060016124ae565b5f613f498254600f81810b600160801b909204900b131590565b15613f6757604051631ed9509560e11b815260040160405180910390fd5b508054600f0b5f9081526001909101602052604090205490565b5f613f9b8254600f81810b600160801b909204900b131590565b15613fb957604051631ed9509560e11b815260040160405180910390fd5b508054600f0b5f818152600180840160205260408220805492905583546fffffffffffffffffffffffffffffffff191692016001600160801b03169190911790915590565b5f60ff8216601f8111156109c657604051632cd44ac360e21b815260040160405180910390fd5b5f81815260018301602052604081205480156140ff575f6140476001836155e9565b85549091505f9061405a906001906155e9565b90508181146140b9575f865f018281548110614078576140786152ca565b905f5260205f200154905080875f018481548110614098576140986152ca565b5f918252602080832090910192909255918252600188019052604090208390555b85548690806140ca576140ca6157fc565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f9055600193505050506109c6565b5f9150506109c6565b5f6109c682670de0b6b3a764000061432d565b5f6001600160401b038211156141825760405162461bcd60e51b815260206004820152602660248201527f53616665436173743a2076616c756520646f65736e27742066697420696e203660448201526534206269747360d01b6064820152608401612694565b5090565b5f6109c6825490565b5f8181526001830160205260408120546141d457508154600181810184555f8481526020808220909301849055845484825282860190935260409020919091556109c6565b505f6109c6565b5f5f6141e8868686614234565b905060018360028111156141fe576141fe615810565b14801561421a57505f848061421557614215615824565b868809115b15613afd5761422a600182615838565b9695505050505050565b5f80805f19858709858702925082811083820303915050805f0361426b5783828161426157614261615824565b0492505050610905565b8084116142b25760405162461bcd60e51b81526020600482015260156024820152744d6174683a206d756c446976206f766572666c6f7760581b6044820152606401612694565b5f8486880960026001871981018816978890046003810283188082028403028082028403028082028403028082028403028082028403029081029092039091025f889003889004909101858311909403939093029303949094049190911702949350505050565b61105383836001600160401b038416614371565b81545f90801561436957614353846143466001846155e9565b5f91825260209091200190565b54600160201b90046001600160e01b03166109c2565b509092915050565b82548015614427575f614389856143466001856155e9565b60408051808201909152905463ffffffff808216808452600160201b9092046001600160e01b0316602084015291925090851610156143db5760405163151b8e3f60e11b815260040160405180910390fd5b805163ffffffff80861691160361442557826143fc866143466001866155e9565b80546001600160e01b0392909216600160201b0263ffffffff9092169190911790555050505050565b505b506040805180820190915263ffffffff92831681526001600160e01b03918216602080830191825285546001810187555f968752952091519051909216600160201b029190921617910155565b634e487b7160e01b5f52604160045260245ffd5b604051606081016001600160401b03811182821017156144aa576144aa614474565b60405290565b604051601f8201601f191681016001600160401b03811182821017156144d8576144d8614474565b604052919050565b6001600160a01b0381168114612aaf575f5ffd5b803563ffffffff81168114610a31575f5ffd5b5f60408284031215614517575f5ffd5b604080519081016001600160401b038111828210171561453957614539614474565b604052905080823561454a816144e0565b8152614558602084016144f4565b60208201525092915050565b5f60408284031215614574575f5ffd5b6109058383614507565b5f5f5f60808486031215614590575f5ffd5b833561459b816144e0565b92506145aa8560208601614507565b915060608401356145ba816144e0565b809150509250925092565b81516001600160401b03168152602080830151600f0b9082015260408083015163ffffffff1690820152606081016109c6565b5f5f60608385031215614609575f5ffd5b8235614614816144e0565b91506146238460208501614507565b90509250929050565b5f6020828403121561463c575f5ffd5b5035919050565b5f60208284031215614653575f5ffd5b8135610905816144e0565b80516001600160a01b0316825260209081015163ffffffff16910152565b5f8151808452602084019350602083015f5b828110156146b7576146a186835161465e565b604095909501946020919091019060010161468e565b5093949350505050565b602081525f610905602083018461467c565b5f5f83601f8401126146e3575f5ffd5b5081356001600160401b038111156146f9575f5ffd5b6020830191508360208260051b8501011115610d88575f5ffd5b5f5f5f60408486031215614725575f5ffd5b8335614730816144e0565b925060208401356001600160401b0381111561474a575f5ffd5b614756868287016146d3565b9497909650939450505050565b5f6001600160401b0382111561477b5761477b614474565b5060051b60200190565b5f82601f830112614794575f5ffd5b81356147a76147a282614763565b6144b0565b8082825260208201915060208360051b8601019250858311156147c8575f5ffd5b602085015b838110156147ee5780356147e0816144e0565b8352602092830192016147cd565b5095945050505050565b5f5f5f6080848603121561480a575f5ffd5b6148148585614507565b925060408401356001600160401b0381111561482e575f5ffd5b61483a86828701614785565b92505060608401356001600160401b03811115614855575f5ffd5b61486186828701614785565b9150509250925092565b5f8151808452602084019350602083015f5b828110156146b757815186526020958601959091019060010161487d565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b828110156148f257603f198786030184526148dd85835161486b565b945060209384019391909101906001016148c1565b50929695505050505050565b5f5f5f5f60a08587031215614911575f5ffd5b61491b8686614507565b935060408501356001600160401b03811115614935575f5ffd5b61494187828801614785565b93505060608501356001600160401b0381111561495c575f5ffd5b61496887828801614785565b925050614977608086016144f4565b905092959194509250565b5f5f5f5f5f60608688031215614996575f5ffd5b85356149a1816144e0565b945060208601356001600160401b038111156149bb575f5ffd5b6149c7888289016146d3565b90955093505060408601356001600160401b038111156149e5575f5ffd5b6149f1888289016146d3565b969995985093965092949392505050565b5f5f60408385031215614a13575f5ffd5b8235614a1e816144e0565b915060208301356001600160401b03811115614a38575f5ffd5b830160a08186031215614a49575f5ffd5b809150509250929050565b828152604060208201525f610af2604083018461486b565b5f82601f830112614a7b575f5ffd5b8135614a896147a282614763565b8082825260208201915060208360051b860101925085831115614aaa575f5ffd5b602085015b838110156147ee5780356001600160401b03811115614acc575f5ffd5b86016060818903601f19011215614ae1575f5ffd5b614ae9614488565b614af5602083016144f4565b815260408201356001600160401b03811115614b0f575f5ffd5b614b1e8a602083860101614785565b60208301525060608201359150614b34826144e0565b6040810191909152835260209283019201614aaf565b5f5f60408385031215614b5b575f5ffd5b8235614b66816144e0565b915060208301356001600160401b03811115614b80575f5ffd5b614b8c85828601614a6c565b9150509250929050565b5f5f60408385031215614ba7575f5ffd5b8235614bb2816144e0565b91506020830135614a49816144e0565b5f8151808452602084019350602083015f5b828110156146b757614c0d86835180516001600160401b03168252602080820151600f0b9083015260409081015163ffffffff16910152565b6060959095019460209190910190600101614bd4565b604081525f614c35604083018561467c565b8281036020840152613afd8185614bc2565b602080825282518282018190525f918401906040840190835b81811015614c875783516001600160a01b0316835260209384019390920191600101614c60565b509095945050505050565b5f5f5f60408486031215614ca4575f5ffd5b83356001600160401b03811115614cb9575f5ffd5b614cc5868287016146d3565b90945092505060208401356145ba816144e0565b602080825282518282018190525f918401906040840190835b81811015614c875783516001600160401b0316835260209384019390920191600101614cf2565b5f5f5f5f60608587031215614d2c575f5ffd5b8435614d37816144e0565b9350614d45602086016144f4565b925060408501356001600160401b03811115614d5f575f5ffd5b614d6b878288016146d3565b95989497509550505050565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b5f5f60408385031215614dbd575f5ffd5b8235614dc8816144e0565b9150614623602084016144f4565b5f60208284031215614de6575f5ffd5b813560ff81168114610905575f5ffd5b5f5f60608385031215614e07575f5ffd5b614e118484614507565b91506040830135614a49816144e0565b5f60608284031215614e31575f5ffd5b50919050565b5f60208284031215614e47575f5ffd5b81356001600160401b03811115614e5c575f5ffd5b6109c284828501614e21565b5f5f5f60808486031215614e7a575f5ffd5b83356001600160401b03811115614e8f575f5ffd5b614e9b86828701614785565b9350506145aa8560208601614507565b602081525f6109056020830184614bc2565b5f5f5f5f60608587031215614ed0575f5ffd5b8435614edb816144e0565b935060208501356001600160401b03811115614ef5575f5ffd5b614f01878288016146d3565b90945092506149779050604086016144f4565b5f5f60408385031215614f25575f5ffd5b8235614f30816144e0565b915060208301356001600160401b03811115614f4a575f5ffd5b8301601f81018513614f5a575f5ffd5b8035614f686147a282614763565b8082825260208201915060208360051b850101925087831115614f89575f5ffd5b602084015b838110156150ae5780356001600160401b03811115614fab575f5ffd5b85016080818b03601f19011215614fc0575f5ffd5b614fc8614488565b614fd58b60208401614507565b815260608201356001600160401b03811115614fef575f5ffd5b614ffe8c602083860101614785565b60208301525060808201356001600160401b0381111561501c575f5ffd5b6020818401019250508a601f830112615033575f5ffd5b81356150416147a282614763565b8082825260208201915060208360051b86010192508d831115615062575f5ffd5b6020850194505b828510156150985784356001600160401b0381168114615087575f5ffd5b825260209485019490910190615069565b6040840152505084525060209283019201614f8e565b50809450505050509250929050565b5f602082840312156150cd575f5ffd5b81356001600160401b038111156150e2575f5ffd5b8201601f810184136150f2575f5ffd5b80356151006147a282614763565b8082825260208201915060208360061b850101925086831115615121575f5ffd5b6020840193505b8284101561422a5761513a8785614507565b8252602082019150604084019350615128565b5f5f5f6040848603121561515f575f5ffd5b833561516a816144e0565b925060208401356001600160401b03811115615184575f5ffd5b8401601f81018613615194575f5ffd5b80356001600160401b038111156151a9575f5ffd5b8660208284010111156151ba575f5ffd5b939660209190910195509293505050565b5f5f604083850312156151dc575f5ffd5b82356151e7816144e0565b915060208301356001600160401b03811115615201575f5ffd5b614b8c85828601614e21565b5f5f5f5f60608587031215615220575f5ffd5b843561522b816144e0565b935060208501356001600160401b03811115615245575f5ffd5b61525187828801614a6c565b93505060408501356001600160401b03811115614d5f575f5ffd5b5f6020828403121561527c575f5ffd5b610905826144f4565b5f5f8335601e1984360301811261529a575f5ffd5b8301803591506001600160401b038211156152b3575f5ffd5b6020019150600581901b3603821315610d88575f5ffd5b634e487b7160e01b5f52603260045260245ffd5b5f602082840312156152ee575f5ffd5b813561ffff81168114610905575f5ffd5b5f6020828403121561530f575f5ffd5b81518015158114610905575f5ffd5b604081016109c6828461465e565b634e487b7160e01b5f52601160045260245ffd5b63ffffffff81811683821601908111156109c6576109c661532c565b8183526020830192505f815f5b848110156146b75763ffffffff61537f836144f4565b1686526020958601959190910190600101615369565b6001600160a01b038581168252841660208201526060604082018190525f9061422a908301848661535c565b6001600160a01b038616815260c081016153de602083018761465e565b6001600160a01b039490941660608201526001600160401b0392909216608083015263ffffffff1660a09091015292915050565b5f60208284031215615422575f5ffd5b81516001600160401b03811115615437575f5ffd5b8201601f81018413615447575f5ffd5b80516154556147a282614763565b8082825260208201915060208360051b850101925086831115615476575f5ffd5b6020840193505b8284101561422a578351615490816144e0565b82526020938401939091019061547d565b606081016154af828561465e565b6001600160a01b039290921660409190910152919050565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b602081525f610af26020830184866154c7565b5f5f8335601e19843603018112615517575f5ffd5b8301803591506001600160401b03821115615530575f5ffd5b602001915036819003821315610d88575f5ffd5b6001600160a01b038781168252861660208201526080604082018190525f90615570908301868861535c565b82810360608401526155838185876154c7565b9998505050505050505050565b80516020808301519190811015614e31575f1960209190910360031b1b16919050565b5f8235603e198336030181126155c7575f5ffd5b9190910192915050565b5f600182016155e2576155e261532c565b5060010190565b818103818111156109c6576109c661532c565b6001600160401b0382811682821603908111156109c6576109c661532c565b5f81600f0b60016001607f1b031981036156375761563761532c565b5f0392915050565b600f81810b9083900b0160016001607f1b03811360016001607f1b0319821217156109c6576109c661532c565b6001600160a01b038716815260e08101615689602083018861465e565b60608201959095526001600160a01b039390931660808401526001600160401b0391821660a08401521660c09091015292915050565b5f602082840312156156cf575f5ffd5b5051919050565b6001600160a01b03881681525f60c082016156f4602084018a61465e565b60c060608401528690528660e083015f5b88811015615735578235615718816144e0565b6001600160a01b0316825260209283019290910190600101615705565b508381036080850152615748818861486b565b91505082810360a084015261575e8185876154c7565b9a9950505050505050505050565b5f8161577a5761577a61532c565b505f190190565b6080810161578f828661465e565b6001600160a01b0393909316604082015263ffffffff91909116606090910152919050565b600f82810b9082900b0360016001607f1b0319811260016001607f1b03821317156109c6576109c661532c565b5f602082840312156157f1575f5ffd5b8151610905816144e0565b634e487b7160e01b5f52603160045260245ffd5b634e487b7160e01b5f52602160045260245ffd5b634e487b7160e01b5f52601260045260245ffd5b808201808211156109c6576109c661532c56fea2646970667358221220aa63cfae26f1918f081c0d9ac5ba605fea04a6889a3fa02162c5e976e3f582d864736f6c634300081e0033",
}

// AllocationManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use AllocationManagerMetaData.ABI instead.
var AllocationManagerABI = AllocationManagerMetaData.ABI

// AllocationManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AllocationManagerMetaData.Bin instead.
var AllocationManagerBin = AllocationManagerMetaData.Bin

// DeployAllocationManager deploys a new Ethereum contract, binding an instance of AllocationManager to it.
func DeployAllocationManager(auth *bind.TransactOpts, backend bind.ContractBackend, _allocationManagerView common.Address, _delegation common.Address, _eigenStrategy common.Address, _pauserRegistry common.Address, _permissionController common.Address, _DEALLOCATION_DELAY uint32, _ALLOCATION_CONFIGURATION_DELAY uint32, _version string) (common.Address, *types.Transaction, *AllocationManager, error) {
	parsed, err := AllocationManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AllocationManagerBin), backend, _allocationManagerView, _delegation, _eigenStrategy, _pauserRegistry, _permissionController, _DEALLOCATION_DELAY, _ALLOCATION_CONFIGURATION_DELAY, _version)
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
// Solidity: function getAllocatableMagnitude(address , address ) view returns(uint64 allocatableMagnitude)
func (_AllocationManager *AllocationManagerCaller) GetAllocatableMagnitude(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (uint64, error) {
	var out []interface{}
	err := _AllocationManager.contract.Call(opts, &out, "getAllocatableMagnitude", arg0, arg1)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetAllocatableMagnitude is a free data retrieval call binding the contract method 0x6cfb4481.
//
// Solidity: function getAllocatableMagnitude(address , address ) view returns(uint64 allocatableMagnitude)
func (_AllocationManager *AllocationManagerSession) GetAllocatableMagnitude(arg0 common.Address, arg1 common.Address) (uint64, error) {
	return _AllocationManager.Contract.GetAllocatableMagnitude(&_AllocationManager.CallOpts, arg0, arg1)
}

// GetAllocatableMagnitude is a free data retrieval call binding the contract method 0x6cfb4481.
//
// Solidity: function getAllocatableMagnitude(address , address ) view returns(uint64 allocatableMagnitude)
func (_AllocationManager *AllocationManagerCallerSession) GetAllocatableMagnitude(arg0 common.Address, arg1 common.Address) (uint64, error) {
	return _AllocationManager.Contract.GetAllocatableMagnitude(&_AllocationManager.CallOpts, arg0, arg1)
}

// GetAllocatedSets is a free data retrieval call binding the contract method 0x15fe5028.
//
// Solidity: function getAllocatedSets(address ) view returns((address,uint32)[] operatorSets)
func (_AllocationManager *AllocationManagerCaller) GetAllocatedSets(opts *bind.CallOpts, arg0 common.Address) ([]OperatorSet, error) {
	var out []interface{}
	err := _AllocationManager.contract.Call(opts, &out, "getAllocatedSets", arg0)

	if err != nil {
		return *new([]OperatorSet), err
	}

	out0 := *abi.ConvertType(out[0], new([]OperatorSet)).(*[]OperatorSet)

	return out0, err

}

// GetAllocatedSets is a free data retrieval call binding the contract method 0x15fe5028.
//
// Solidity: function getAllocatedSets(address ) view returns((address,uint32)[] operatorSets)
func (_AllocationManager *AllocationManagerSession) GetAllocatedSets(arg0 common.Address) ([]OperatorSet, error) {
	return _AllocationManager.Contract.GetAllocatedSets(&_AllocationManager.CallOpts, arg0)
}

// GetAllocatedSets is a free data retrieval call binding the contract method 0x15fe5028.
//
// Solidity: function getAllocatedSets(address ) view returns((address,uint32)[] operatorSets)
func (_AllocationManager *AllocationManagerCallerSession) GetAllocatedSets(arg0 common.Address) ([]OperatorSet, error) {
	return _AllocationManager.Contract.GetAllocatedSets(&_AllocationManager.CallOpts, arg0)
}

// GetAllocatedStake is a free data retrieval call binding the contract method 0x2b453a9a.
//
// Solidity: function getAllocatedStake((address,uint32) , address[] , address[] ) view returns(uint256[][] slashableStake)
func (_AllocationManager *AllocationManagerCaller) GetAllocatedStake(opts *bind.CallOpts, arg0 OperatorSet, arg1 []common.Address, arg2 []common.Address) ([][]*big.Int, error) {
	var out []interface{}
	err := _AllocationManager.contract.Call(opts, &out, "getAllocatedStake", arg0, arg1, arg2)

	if err != nil {
		return *new([][]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([][]*big.Int)).(*[][]*big.Int)

	return out0, err

}

// GetAllocatedStake is a free data retrieval call binding the contract method 0x2b453a9a.
//
// Solidity: function getAllocatedStake((address,uint32) , address[] , address[] ) view returns(uint256[][] slashableStake)
func (_AllocationManager *AllocationManagerSession) GetAllocatedStake(arg0 OperatorSet, arg1 []common.Address, arg2 []common.Address) ([][]*big.Int, error) {
	return _AllocationManager.Contract.GetAllocatedStake(&_AllocationManager.CallOpts, arg0, arg1, arg2)
}

// GetAllocatedStake is a free data retrieval call binding the contract method 0x2b453a9a.
//
// Solidity: function getAllocatedStake((address,uint32) , address[] , address[] ) view returns(uint256[][] slashableStake)
func (_AllocationManager *AllocationManagerCallerSession) GetAllocatedStake(arg0 OperatorSet, arg1 []common.Address, arg2 []common.Address) ([][]*big.Int, error) {
	return _AllocationManager.Contract.GetAllocatedStake(&_AllocationManager.CallOpts, arg0, arg1, arg2)
}

// GetAllocatedStrategies is a free data retrieval call binding the contract method 0xc221d8ae.
//
// Solidity: function getAllocatedStrategies(address , (address,uint32) ) view returns(address[] strategies)
func (_AllocationManager *AllocationManagerCaller) GetAllocatedStrategies(opts *bind.CallOpts, arg0 common.Address, arg1 OperatorSet) ([]common.Address, error) {
	var out []interface{}
	err := _AllocationManager.contract.Call(opts, &out, "getAllocatedStrategies", arg0, arg1)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllocatedStrategies is a free data retrieval call binding the contract method 0xc221d8ae.
//
// Solidity: function getAllocatedStrategies(address , (address,uint32) ) view returns(address[] strategies)
func (_AllocationManager *AllocationManagerSession) GetAllocatedStrategies(arg0 common.Address, arg1 OperatorSet) ([]common.Address, error) {
	return _AllocationManager.Contract.GetAllocatedStrategies(&_AllocationManager.CallOpts, arg0, arg1)
}

// GetAllocatedStrategies is a free data retrieval call binding the contract method 0xc221d8ae.
//
// Solidity: function getAllocatedStrategies(address , (address,uint32) ) view returns(address[] strategies)
func (_AllocationManager *AllocationManagerCallerSession) GetAllocatedStrategies(arg0 common.Address, arg1 OperatorSet) ([]common.Address, error) {
	return _AllocationManager.Contract.GetAllocatedStrategies(&_AllocationManager.CallOpts, arg0, arg1)
}

// GetAllocation is a free data retrieval call binding the contract method 0x10e1b9b8.
//
// Solidity: function getAllocation(address , (address,uint32) , address ) view returns((uint64,int128,uint32) allocation)
func (_AllocationManager *AllocationManagerCaller) GetAllocation(opts *bind.CallOpts, arg0 common.Address, arg1 OperatorSet, arg2 common.Address) (IAllocationManagerTypesAllocation, error) {
	var out []interface{}
	err := _AllocationManager.contract.Call(opts, &out, "getAllocation", arg0, arg1, arg2)

	if err != nil {
		return *new(IAllocationManagerTypesAllocation), err
	}

	out0 := *abi.ConvertType(out[0], new(IAllocationManagerTypesAllocation)).(*IAllocationManagerTypesAllocation)

	return out0, err

}

// GetAllocation is a free data retrieval call binding the contract method 0x10e1b9b8.
//
// Solidity: function getAllocation(address , (address,uint32) , address ) view returns((uint64,int128,uint32) allocation)
func (_AllocationManager *AllocationManagerSession) GetAllocation(arg0 common.Address, arg1 OperatorSet, arg2 common.Address) (IAllocationManagerTypesAllocation, error) {
	return _AllocationManager.Contract.GetAllocation(&_AllocationManager.CallOpts, arg0, arg1, arg2)
}

// GetAllocation is a free data retrieval call binding the contract method 0x10e1b9b8.
//
// Solidity: function getAllocation(address , (address,uint32) , address ) view returns((uint64,int128,uint32) allocation)
func (_AllocationManager *AllocationManagerCallerSession) GetAllocation(arg0 common.Address, arg1 OperatorSet, arg2 common.Address) (IAllocationManagerTypesAllocation, error) {
	return _AllocationManager.Contract.GetAllocation(&_AllocationManager.CallOpts, arg0, arg1, arg2)
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
// Solidity: function getAllocations(address[] , (address,uint32) , address ) view returns((uint64,int128,uint32)[] allocations)
func (_AllocationManager *AllocationManagerCaller) GetAllocations(opts *bind.CallOpts, arg0 []common.Address, arg1 OperatorSet, arg2 common.Address) ([]IAllocationManagerTypesAllocation, error) {
	var out []interface{}
	err := _AllocationManager.contract.Call(opts, &out, "getAllocations", arg0, arg1, arg2)

	if err != nil {
		return *new([]IAllocationManagerTypesAllocation), err
	}

	out0 := *abi.ConvertType(out[0], new([]IAllocationManagerTypesAllocation)).(*[]IAllocationManagerTypesAllocation)

	return out0, err

}

// GetAllocations is a free data retrieval call binding the contract method 0x8ce64854.
//
// Solidity: function getAllocations(address[] , (address,uint32) , address ) view returns((uint64,int128,uint32)[] allocations)
func (_AllocationManager *AllocationManagerSession) GetAllocations(arg0 []common.Address, arg1 OperatorSet, arg2 common.Address) ([]IAllocationManagerTypesAllocation, error) {
	return _AllocationManager.Contract.GetAllocations(&_AllocationManager.CallOpts, arg0, arg1, arg2)
}

// GetAllocations is a free data retrieval call binding the contract method 0x8ce64854.
//
// Solidity: function getAllocations(address[] , (address,uint32) , address ) view returns((uint64,int128,uint32)[] allocations)
func (_AllocationManager *AllocationManagerCallerSession) GetAllocations(arg0 []common.Address, arg1 OperatorSet, arg2 common.Address) ([]IAllocationManagerTypesAllocation, error) {
	return _AllocationManager.Contract.GetAllocations(&_AllocationManager.CallOpts, arg0, arg1, arg2)
}

// GetEncumberedMagnitude is a free data retrieval call binding the contract method 0xf605ce08.
//
// Solidity: function getEncumberedMagnitude(address , address ) view returns(uint64 encumberedMagnitude)
func (_AllocationManager *AllocationManagerCaller) GetEncumberedMagnitude(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (uint64, error) {
	var out []interface{}
	err := _AllocationManager.contract.Call(opts, &out, "getEncumberedMagnitude", arg0, arg1)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetEncumberedMagnitude is a free data retrieval call binding the contract method 0xf605ce08.
//
// Solidity: function getEncumberedMagnitude(address , address ) view returns(uint64 encumberedMagnitude)
func (_AllocationManager *AllocationManagerSession) GetEncumberedMagnitude(arg0 common.Address, arg1 common.Address) (uint64, error) {
	return _AllocationManager.Contract.GetEncumberedMagnitude(&_AllocationManager.CallOpts, arg0, arg1)
}

// GetEncumberedMagnitude is a free data retrieval call binding the contract method 0xf605ce08.
//
// Solidity: function getEncumberedMagnitude(address , address ) view returns(uint64 encumberedMagnitude)
func (_AllocationManager *AllocationManagerCallerSession) GetEncumberedMagnitude(arg0 common.Address, arg1 common.Address) (uint64, error) {
	return _AllocationManager.Contract.GetEncumberedMagnitude(&_AllocationManager.CallOpts, arg0, arg1)
}

// GetMaxMagnitude is a free data retrieval call binding the contract method 0xa9333ec8.
//
// Solidity: function getMaxMagnitude(address , address ) view returns(uint64 maxMagnitude)
func (_AllocationManager *AllocationManagerCaller) GetMaxMagnitude(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (uint64, error) {
	var out []interface{}
	err := _AllocationManager.contract.Call(opts, &out, "getMaxMagnitude", arg0, arg1)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetMaxMagnitude is a free data retrieval call binding the contract method 0xa9333ec8.
//
// Solidity: function getMaxMagnitude(address , address ) view returns(uint64 maxMagnitude)
func (_AllocationManager *AllocationManagerSession) GetMaxMagnitude(arg0 common.Address, arg1 common.Address) (uint64, error) {
	return _AllocationManager.Contract.GetMaxMagnitude(&_AllocationManager.CallOpts, arg0, arg1)
}

// GetMaxMagnitude is a free data retrieval call binding the contract method 0xa9333ec8.
//
// Solidity: function getMaxMagnitude(address , address ) view returns(uint64 maxMagnitude)
func (_AllocationManager *AllocationManagerCallerSession) GetMaxMagnitude(arg0 common.Address, arg1 common.Address) (uint64, error) {
	return _AllocationManager.Contract.GetMaxMagnitude(&_AllocationManager.CallOpts, arg0, arg1)
}

// GetMaxMagnitudes is a free data retrieval call binding the contract method 0x4a10ffe5.
//
// Solidity: function getMaxMagnitudes(address[] , address ) view returns(uint64[] maxMagnitudes)
func (_AllocationManager *AllocationManagerCaller) GetMaxMagnitudes(opts *bind.CallOpts, arg0 []common.Address, arg1 common.Address) ([]uint64, error) {
	var out []interface{}
	err := _AllocationManager.contract.Call(opts, &out, "getMaxMagnitudes", arg0, arg1)

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

// GetMaxMagnitudes is a free data retrieval call binding the contract method 0x4a10ffe5.
//
// Solidity: function getMaxMagnitudes(address[] , address ) view returns(uint64[] maxMagnitudes)
func (_AllocationManager *AllocationManagerSession) GetMaxMagnitudes(arg0 []common.Address, arg1 common.Address) ([]uint64, error) {
	return _AllocationManager.Contract.GetMaxMagnitudes(&_AllocationManager.CallOpts, arg0, arg1)
}

// GetMaxMagnitudes is a free data retrieval call binding the contract method 0x4a10ffe5.
//
// Solidity: function getMaxMagnitudes(address[] , address ) view returns(uint64[] maxMagnitudes)
func (_AllocationManager *AllocationManagerCallerSession) GetMaxMagnitudes(arg0 []common.Address, arg1 common.Address) ([]uint64, error) {
	return _AllocationManager.Contract.GetMaxMagnitudes(&_AllocationManager.CallOpts, arg0, arg1)
}

// GetMaxMagnitudes0 is a free data retrieval call binding the contract method 0x547afb87.
//
// Solidity: function getMaxMagnitudes(address , address[] ) view returns(uint64[] maxMagnitudes)
func (_AllocationManager *AllocationManagerCaller) GetMaxMagnitudes0(opts *bind.CallOpts, arg0 common.Address, arg1 []common.Address) ([]uint64, error) {
	var out []interface{}
	err := _AllocationManager.contract.Call(opts, &out, "getMaxMagnitudes0", arg0, arg1)

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

// GetMaxMagnitudes0 is a free data retrieval call binding the contract method 0x547afb87.
//
// Solidity: function getMaxMagnitudes(address , address[] ) view returns(uint64[] maxMagnitudes)
func (_AllocationManager *AllocationManagerSession) GetMaxMagnitudes0(arg0 common.Address, arg1 []common.Address) ([]uint64, error) {
	return _AllocationManager.Contract.GetMaxMagnitudes0(&_AllocationManager.CallOpts, arg0, arg1)
}

// GetMaxMagnitudes0 is a free data retrieval call binding the contract method 0x547afb87.
//
// Solidity: function getMaxMagnitudes(address , address[] ) view returns(uint64[] maxMagnitudes)
func (_AllocationManager *AllocationManagerCallerSession) GetMaxMagnitudes0(arg0 common.Address, arg1 []common.Address) ([]uint64, error) {
	return _AllocationManager.Contract.GetMaxMagnitudes0(&_AllocationManager.CallOpts, arg0, arg1)
}

// GetMaxMagnitudesAtBlock is a free data retrieval call binding the contract method 0x94d7d00c.
//
// Solidity: function getMaxMagnitudesAtBlock(address , address[] , uint32 ) view returns(uint64[] maxMagnitudes)
func (_AllocationManager *AllocationManagerCaller) GetMaxMagnitudesAtBlock(opts *bind.CallOpts, arg0 common.Address, arg1 []common.Address, arg2 uint32) ([]uint64, error) {
	var out []interface{}
	err := _AllocationManager.contract.Call(opts, &out, "getMaxMagnitudesAtBlock", arg0, arg1, arg2)

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

// GetMaxMagnitudesAtBlock is a free data retrieval call binding the contract method 0x94d7d00c.
//
// Solidity: function getMaxMagnitudesAtBlock(address , address[] , uint32 ) view returns(uint64[] maxMagnitudes)
func (_AllocationManager *AllocationManagerSession) GetMaxMagnitudesAtBlock(arg0 common.Address, arg1 []common.Address, arg2 uint32) ([]uint64, error) {
	return _AllocationManager.Contract.GetMaxMagnitudesAtBlock(&_AllocationManager.CallOpts, arg0, arg1, arg2)
}

// GetMaxMagnitudesAtBlock is a free data retrieval call binding the contract method 0x94d7d00c.
//
// Solidity: function getMaxMagnitudesAtBlock(address , address[] , uint32 ) view returns(uint64[] maxMagnitudes)
func (_AllocationManager *AllocationManagerCallerSession) GetMaxMagnitudesAtBlock(arg0 common.Address, arg1 []common.Address, arg2 uint32) ([]uint64, error) {
	return _AllocationManager.Contract.GetMaxMagnitudesAtBlock(&_AllocationManager.CallOpts, arg0, arg1, arg2)
}

// GetMemberCount is a free data retrieval call binding the contract method 0xb2447af7.
//
// Solidity: function getMemberCount((address,uint32) ) view returns(uint256 memberCount)
func (_AllocationManager *AllocationManagerCaller) GetMemberCount(opts *bind.CallOpts, arg0 OperatorSet) (*big.Int, error) {
	var out []interface{}
	err := _AllocationManager.contract.Call(opts, &out, "getMemberCount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMemberCount is a free data retrieval call binding the contract method 0xb2447af7.
//
// Solidity: function getMemberCount((address,uint32) ) view returns(uint256 memberCount)
func (_AllocationManager *AllocationManagerSession) GetMemberCount(arg0 OperatorSet) (*big.Int, error) {
	return _AllocationManager.Contract.GetMemberCount(&_AllocationManager.CallOpts, arg0)
}

// GetMemberCount is a free data retrieval call binding the contract method 0xb2447af7.
//
// Solidity: function getMemberCount((address,uint32) ) view returns(uint256 memberCount)
func (_AllocationManager *AllocationManagerCallerSession) GetMemberCount(arg0 OperatorSet) (*big.Int, error) {
	return _AllocationManager.Contract.GetMemberCount(&_AllocationManager.CallOpts, arg0)
}

// GetMembers is a free data retrieval call binding the contract method 0x6e875dba.
//
// Solidity: function getMembers((address,uint32) ) view returns(address[] operators)
func (_AllocationManager *AllocationManagerCaller) GetMembers(opts *bind.CallOpts, arg0 OperatorSet) ([]common.Address, error) {
	var out []interface{}
	err := _AllocationManager.contract.Call(opts, &out, "getMembers", arg0)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetMembers is a free data retrieval call binding the contract method 0x6e875dba.
//
// Solidity: function getMembers((address,uint32) ) view returns(address[] operators)
func (_AllocationManager *AllocationManagerSession) GetMembers(arg0 OperatorSet) ([]common.Address, error) {
	return _AllocationManager.Contract.GetMembers(&_AllocationManager.CallOpts, arg0)
}

// GetMembers is a free data retrieval call binding the contract method 0x6e875dba.
//
// Solidity: function getMembers((address,uint32) ) view returns(address[] operators)
func (_AllocationManager *AllocationManagerCallerSession) GetMembers(arg0 OperatorSet) ([]common.Address, error) {
	return _AllocationManager.Contract.GetMembers(&_AllocationManager.CallOpts, arg0)
}

// GetMinimumSlashableStake is a free data retrieval call binding the contract method 0x2bab2c4a.
//
// Solidity: function getMinimumSlashableStake((address,uint32) , address[] , address[] , uint32 ) view returns(uint256[][] slashableStake)
func (_AllocationManager *AllocationManagerCaller) GetMinimumSlashableStake(opts *bind.CallOpts, arg0 OperatorSet, arg1 []common.Address, arg2 []common.Address, arg3 uint32) ([][]*big.Int, error) {
	var out []interface{}
	err := _AllocationManager.contract.Call(opts, &out, "getMinimumSlashableStake", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([][]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([][]*big.Int)).(*[][]*big.Int)

	return out0, err

}

// GetMinimumSlashableStake is a free data retrieval call binding the contract method 0x2bab2c4a.
//
// Solidity: function getMinimumSlashableStake((address,uint32) , address[] , address[] , uint32 ) view returns(uint256[][] slashableStake)
func (_AllocationManager *AllocationManagerSession) GetMinimumSlashableStake(arg0 OperatorSet, arg1 []common.Address, arg2 []common.Address, arg3 uint32) ([][]*big.Int, error) {
	return _AllocationManager.Contract.GetMinimumSlashableStake(&_AllocationManager.CallOpts, arg0, arg1, arg2, arg3)
}

// GetMinimumSlashableStake is a free data retrieval call binding the contract method 0x2bab2c4a.
//
// Solidity: function getMinimumSlashableStake((address,uint32) , address[] , address[] , uint32 ) view returns(uint256[][] slashableStake)
func (_AllocationManager *AllocationManagerCallerSession) GetMinimumSlashableStake(arg0 OperatorSet, arg1 []common.Address, arg2 []common.Address, arg3 uint32) ([][]*big.Int, error) {
	return _AllocationManager.Contract.GetMinimumSlashableStake(&_AllocationManager.CallOpts, arg0, arg1, arg2, arg3)
}

// GetOperatorSetCount is a free data retrieval call binding the contract method 0xba1a84e5.
//
// Solidity: function getOperatorSetCount(address ) view returns(uint256 count)
func (_AllocationManager *AllocationManagerCaller) GetOperatorSetCount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AllocationManager.contract.Call(opts, &out, "getOperatorSetCount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOperatorSetCount is a free data retrieval call binding the contract method 0xba1a84e5.
//
// Solidity: function getOperatorSetCount(address ) view returns(uint256 count)
func (_AllocationManager *AllocationManagerSession) GetOperatorSetCount(arg0 common.Address) (*big.Int, error) {
	return _AllocationManager.Contract.GetOperatorSetCount(&_AllocationManager.CallOpts, arg0)
}

// GetOperatorSetCount is a free data retrieval call binding the contract method 0xba1a84e5.
//
// Solidity: function getOperatorSetCount(address ) view returns(uint256 count)
func (_AllocationManager *AllocationManagerCallerSession) GetOperatorSetCount(arg0 common.Address) (*big.Int, error) {
	return _AllocationManager.Contract.GetOperatorSetCount(&_AllocationManager.CallOpts, arg0)
}

// GetPendingSlasher is a free data retrieval call binding the contract method 0xd7794857.
//
// Solidity: function getPendingSlasher((address,uint32) ) view returns(address pendingSlasher, uint32 effectBlock)
func (_AllocationManager *AllocationManagerCaller) GetPendingSlasher(opts *bind.CallOpts, arg0 OperatorSet) (struct {
	PendingSlasher common.Address
	EffectBlock    uint32
}, error) {
	var out []interface{}
	err := _AllocationManager.contract.Call(opts, &out, "getPendingSlasher", arg0)

	outstruct := new(struct {
		PendingSlasher common.Address
		EffectBlock    uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PendingSlasher = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.EffectBlock = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

// GetPendingSlasher is a free data retrieval call binding the contract method 0xd7794857.
//
// Solidity: function getPendingSlasher((address,uint32) ) view returns(address pendingSlasher, uint32 effectBlock)
func (_AllocationManager *AllocationManagerSession) GetPendingSlasher(arg0 OperatorSet) (struct {
	PendingSlasher common.Address
	EffectBlock    uint32
}, error) {
	return _AllocationManager.Contract.GetPendingSlasher(&_AllocationManager.CallOpts, arg0)
}

// GetPendingSlasher is a free data retrieval call binding the contract method 0xd7794857.
//
// Solidity: function getPendingSlasher((address,uint32) ) view returns(address pendingSlasher, uint32 effectBlock)
func (_AllocationManager *AllocationManagerCallerSession) GetPendingSlasher(arg0 OperatorSet) (struct {
	PendingSlasher common.Address
	EffectBlock    uint32
}, error) {
	return _AllocationManager.Contract.GetPendingSlasher(&_AllocationManager.CallOpts, arg0)
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
// Solidity: function getRegisteredSets(address ) view returns((address,uint32)[] operatorSets)
func (_AllocationManager *AllocationManagerCaller) GetRegisteredSets(opts *bind.CallOpts, arg0 common.Address) ([]OperatorSet, error) {
	var out []interface{}
	err := _AllocationManager.contract.Call(opts, &out, "getRegisteredSets", arg0)

	if err != nil {
		return *new([]OperatorSet), err
	}

	out0 := *abi.ConvertType(out[0], new([]OperatorSet)).(*[]OperatorSet)

	return out0, err

}

// GetRegisteredSets is a free data retrieval call binding the contract method 0x79ae50cd.
//
// Solidity: function getRegisteredSets(address ) view returns((address,uint32)[] operatorSets)
func (_AllocationManager *AllocationManagerSession) GetRegisteredSets(arg0 common.Address) ([]OperatorSet, error) {
	return _AllocationManager.Contract.GetRegisteredSets(&_AllocationManager.CallOpts, arg0)
}

// GetRegisteredSets is a free data retrieval call binding the contract method 0x79ae50cd.
//
// Solidity: function getRegisteredSets(address ) view returns((address,uint32)[] operatorSets)
func (_AllocationManager *AllocationManagerCallerSession) GetRegisteredSets(arg0 common.Address) ([]OperatorSet, error) {
	return _AllocationManager.Contract.GetRegisteredSets(&_AllocationManager.CallOpts, arg0)
}

// GetSlashCount is a free data retrieval call binding the contract method 0x4cfd2939.
//
// Solidity: function getSlashCount((address,uint32) ) view returns(uint256 slashCount)
func (_AllocationManager *AllocationManagerCaller) GetSlashCount(opts *bind.CallOpts, arg0 OperatorSet) (*big.Int, error) {
	var out []interface{}
	err := _AllocationManager.contract.Call(opts, &out, "getSlashCount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSlashCount is a free data retrieval call binding the contract method 0x4cfd2939.
//
// Solidity: function getSlashCount((address,uint32) ) view returns(uint256 slashCount)
func (_AllocationManager *AllocationManagerSession) GetSlashCount(arg0 OperatorSet) (*big.Int, error) {
	return _AllocationManager.Contract.GetSlashCount(&_AllocationManager.CallOpts, arg0)
}

// GetSlashCount is a free data retrieval call binding the contract method 0x4cfd2939.
//
// Solidity: function getSlashCount((address,uint32) ) view returns(uint256 slashCount)
func (_AllocationManager *AllocationManagerCallerSession) GetSlashCount(arg0 OperatorSet) (*big.Int, error) {
	return _AllocationManager.Contract.GetSlashCount(&_AllocationManager.CallOpts, arg0)
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
// Solidity: function getStrategiesInOperatorSet((address,uint32) ) view returns(address[] strategies)
func (_AllocationManager *AllocationManagerCaller) GetStrategiesInOperatorSet(opts *bind.CallOpts, arg0 OperatorSet) ([]common.Address, error) {
	var out []interface{}
	err := _AllocationManager.contract.Call(opts, &out, "getStrategiesInOperatorSet", arg0)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetStrategiesInOperatorSet is a free data retrieval call binding the contract method 0x4177a87c.
//
// Solidity: function getStrategiesInOperatorSet((address,uint32) ) view returns(address[] strategies)
func (_AllocationManager *AllocationManagerSession) GetStrategiesInOperatorSet(arg0 OperatorSet) ([]common.Address, error) {
	return _AllocationManager.Contract.GetStrategiesInOperatorSet(&_AllocationManager.CallOpts, arg0)
}

// GetStrategiesInOperatorSet is a free data retrieval call binding the contract method 0x4177a87c.
//
// Solidity: function getStrategiesInOperatorSet((address,uint32) ) view returns(address[] strategies)
func (_AllocationManager *AllocationManagerCallerSession) GetStrategiesInOperatorSet(arg0 OperatorSet) ([]common.Address, error) {
	return _AllocationManager.Contract.GetStrategiesInOperatorSet(&_AllocationManager.CallOpts, arg0)
}

// GetStrategyAllocations is a free data retrieval call binding the contract method 0x40120dab.
//
// Solidity: function getStrategyAllocations(address , address ) view returns((address,uint32)[] operatorSets, (uint64,int128,uint32)[] allocations)
func (_AllocationManager *AllocationManagerCaller) GetStrategyAllocations(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (struct {
	OperatorSets []OperatorSet
	Allocations  []IAllocationManagerTypesAllocation
}, error) {
	var out []interface{}
	err := _AllocationManager.contract.Call(opts, &out, "getStrategyAllocations", arg0, arg1)

	outstruct := new(struct {
		OperatorSets []OperatorSet
		Allocations  []IAllocationManagerTypesAllocation
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OperatorSets = *abi.ConvertType(out[0], new([]OperatorSet)).(*[]OperatorSet)
	outstruct.Allocations = *abi.ConvertType(out[1], new([]IAllocationManagerTypesAllocation)).(*[]IAllocationManagerTypesAllocation)

	return *outstruct, err

}

// GetStrategyAllocations is a free data retrieval call binding the contract method 0x40120dab.
//
// Solidity: function getStrategyAllocations(address , address ) view returns((address,uint32)[] operatorSets, (uint64,int128,uint32)[] allocations)
func (_AllocationManager *AllocationManagerSession) GetStrategyAllocations(arg0 common.Address, arg1 common.Address) (struct {
	OperatorSets []OperatorSet
	Allocations  []IAllocationManagerTypesAllocation
}, error) {
	return _AllocationManager.Contract.GetStrategyAllocations(&_AllocationManager.CallOpts, arg0, arg1)
}

// GetStrategyAllocations is a free data retrieval call binding the contract method 0x40120dab.
//
// Solidity: function getStrategyAllocations(address , address ) view returns((address,uint32)[] operatorSets, (uint64,int128,uint32)[] allocations)
func (_AllocationManager *AllocationManagerCallerSession) GetStrategyAllocations(arg0 common.Address, arg1 common.Address) (struct {
	OperatorSets []OperatorSet
	Allocations  []IAllocationManagerTypesAllocation
}, error) {
	return _AllocationManager.Contract.GetStrategyAllocations(&_AllocationManager.CallOpts, arg0, arg1)
}

// IsMemberOfOperatorSet is a free data retrieval call binding the contract method 0x670d3ba2.
//
// Solidity: function isMemberOfOperatorSet(address , (address,uint32) ) view returns(bool result)
func (_AllocationManager *AllocationManagerCaller) IsMemberOfOperatorSet(opts *bind.CallOpts, arg0 common.Address, arg1 OperatorSet) (bool, error) {
	var out []interface{}
	err := _AllocationManager.contract.Call(opts, &out, "isMemberOfOperatorSet", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMemberOfOperatorSet is a free data retrieval call binding the contract method 0x670d3ba2.
//
// Solidity: function isMemberOfOperatorSet(address , (address,uint32) ) view returns(bool result)
func (_AllocationManager *AllocationManagerSession) IsMemberOfOperatorSet(arg0 common.Address, arg1 OperatorSet) (bool, error) {
	return _AllocationManager.Contract.IsMemberOfOperatorSet(&_AllocationManager.CallOpts, arg0, arg1)
}

// IsMemberOfOperatorSet is a free data retrieval call binding the contract method 0x670d3ba2.
//
// Solidity: function isMemberOfOperatorSet(address , (address,uint32) ) view returns(bool result)
func (_AllocationManager *AllocationManagerCallerSession) IsMemberOfOperatorSet(arg0 common.Address, arg1 OperatorSet) (bool, error) {
	return _AllocationManager.Contract.IsMemberOfOperatorSet(&_AllocationManager.CallOpts, arg0, arg1)
}

// IsOperatorRedistributable is a free data retrieval call binding the contract method 0xdc2af692.
//
// Solidity: function isOperatorRedistributable(address ) view returns(bool result)
func (_AllocationManager *AllocationManagerCaller) IsOperatorRedistributable(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _AllocationManager.contract.Call(opts, &out, "isOperatorRedistributable", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperatorRedistributable is a free data retrieval call binding the contract method 0xdc2af692.
//
// Solidity: function isOperatorRedistributable(address ) view returns(bool result)
func (_AllocationManager *AllocationManagerSession) IsOperatorRedistributable(arg0 common.Address) (bool, error) {
	return _AllocationManager.Contract.IsOperatorRedistributable(&_AllocationManager.CallOpts, arg0)
}

// IsOperatorRedistributable is a free data retrieval call binding the contract method 0xdc2af692.
//
// Solidity: function isOperatorRedistributable(address ) view returns(bool result)
func (_AllocationManager *AllocationManagerCallerSession) IsOperatorRedistributable(arg0 common.Address) (bool, error) {
	return _AllocationManager.Contract.IsOperatorRedistributable(&_AllocationManager.CallOpts, arg0)
}

// IsOperatorSet is a free data retrieval call binding the contract method 0x260dc758.
//
// Solidity: function isOperatorSet((address,uint32) ) view returns(bool result)
func (_AllocationManager *AllocationManagerCaller) IsOperatorSet(opts *bind.CallOpts, arg0 OperatorSet) (bool, error) {
	var out []interface{}
	err := _AllocationManager.contract.Call(opts, &out, "isOperatorSet", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperatorSet is a free data retrieval call binding the contract method 0x260dc758.
//
// Solidity: function isOperatorSet((address,uint32) ) view returns(bool result)
func (_AllocationManager *AllocationManagerSession) IsOperatorSet(arg0 OperatorSet) (bool, error) {
	return _AllocationManager.Contract.IsOperatorSet(&_AllocationManager.CallOpts, arg0)
}

// IsOperatorSet is a free data retrieval call binding the contract method 0x260dc758.
//
// Solidity: function isOperatorSet((address,uint32) ) view returns(bool result)
func (_AllocationManager *AllocationManagerCallerSession) IsOperatorSet(arg0 OperatorSet) (bool, error) {
	return _AllocationManager.Contract.IsOperatorSet(&_AllocationManager.CallOpts, arg0)
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
