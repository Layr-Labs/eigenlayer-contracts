// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package AllocationManagerView

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

// IAllocationManagerTypesAllocation is an auto generated low-level Go binding around an user-defined struct.
type IAllocationManagerTypesAllocation struct {
	CurrentMagnitude uint64
	PendingDiff      *big.Int
	EffectBlock      uint32
}

// OperatorSet is an auto generated low-level Go binding around an user-defined struct.
type OperatorSet struct {
	Avs common.Address
	Id  uint32
}

// AllocationManagerViewMetaData contains all meta data concerning the AllocationManagerView contract.
var AllocationManagerViewMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_delegation\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"},{\"name\":\"_eigenStrategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"_DEALLOCATION_DELAY\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_ALLOCATION_CONFIGURATION_DELAY\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ALLOCATION_CONFIGURATION_DELAY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEALLOCATION_DELAY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SLASHER_CONFIGURATION_DELAY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eigenStrategy\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAVSRegistrar\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAVSRegistrar\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocatableMagnitude\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocatedSets\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocatedStake\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[][]\",\"internalType\":\"uint256[][]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocatedStrategies\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocation\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIAllocationManagerTypes.Allocation\",\"components\":[{\"name\":\"currentMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"pendingDiff\",\"type\":\"int128\",\"internalType\":\"int128\"},{\"name\":\"effectBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocationDelay\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocations\",\"inputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIAllocationManagerTypes.Allocation[]\",\"components\":[{\"name\":\"currentMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"pendingDiff\",\"type\":\"int128\",\"internalType\":\"int128\"},{\"name\":\"effectBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getEncumberedMagnitude\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxMagnitude\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxMagnitudes\",\"inputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxMagnitudes\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxMagnitudesAtBlock\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMemberCount\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMembers\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMinimumSlashableStake\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"futureBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"slashableStake\",\"type\":\"uint256[][]\",\"internalType\":\"uint256[][]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorSetCount\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPendingSlasher\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRedistributionRecipient\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRegisteredSets\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSlashCount\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSlasher\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStrategiesInOperatorSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStrategyAllocations\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIAllocationManagerTypes.Allocation[]\",\"components\":[{\"name\":\"currentMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"pendingDiff\",\"type\":\"int128\",\"internalType\":\"int128\"},{\"name\":\"effectBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isMemberOfOperatorSet\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperatorRedistributable\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperatorSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperatorSlashable\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRedistributingOperatorSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AVSMetadataURIUpdated\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"metadataURI\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AVSRegistrarSet\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"registrar\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIAVSRegistrar\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllocationDelaySet\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"delay\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"effectBlock\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllocationUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"magnitude\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"effectBlock\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EncumberedMagnitudeUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"encumberedMagnitude\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MaxMagnitudeUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"maxMagnitude\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorAddedToOperatorSet\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRemovedFromOperatorSet\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSetCreated\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSlashed\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategies\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"contractIStrategy[]\"},{\"name\":\"wadSlashed\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"description\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RedistributionAddressSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"redistributionRecipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SlasherMigrated\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slasher\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SlasherUpdated\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slasher\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"effectBlock\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyAddedToOperatorSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyRemovedFromOperatorSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyMemberOfSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputArrayLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientMagnitude\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAVSRegistrar\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCaller\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperatorSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRedistributionRecipient\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidStrategy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidWadToSlash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ModificationAlreadyPending\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonexistentAVSMetadata\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotMemberOfSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorNotSlashable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorSetAlreadyMigrated\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OutOfBounds\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SameMagnitude\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SlasherNotSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategiesMustBeInAscendingOrder\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategyAlreadyInOperatorSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategyNotInOperatorSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UninitializedAllocationDelay\",\"inputs\":[]}]",
	Bin: "0x610120604052348015610010575f5ffd5b5060405161277a38038061277a83398101604081905261002f9161008d565b6001600160a01b039384166080529290911660a05263ffffffff90811660c0521660e0819052610100526100e3565b6001600160a01b0381168114610072575f5ffd5b50565b805163ffffffff81168114610088575f5ffd5b919050565b5f5f5f5f608085870312156100a0575f5ffd5b84516100ab8161005e565b60208601519094506100bc8161005e565b92506100ca60408601610075565b91506100d860608601610075565b905092959194509250565b60805160a05160c05160e051610100516126536101275f395f6103d601525f61044e01525f6102ac01525f61056d01525f81816105a701526113e201526126535ff3fe608060405234801561000f575f5ffd5b50600436106101fd575f3560e01c80636e875dba11610114578063ba1a84e5116100a9578063db4df76111610079578063db4df76114610568578063dc2af6921461058f578063df5cf723146105a2578063f231bd08146105c9578063f605ce08146105dc575f5ffd5b8063ba1a84e5146104f8578063c221d8ae1461050b578063d4a3fcce1461051e578063d779485714610531575f5ffd5b806394d7d00c116100e457806394d7d00c14610490578063a9333ec8146104a3578063b2447af7146104b6578063b9fbaed1146104c9575f5ffd5b80636e875dba1461042357806379ae50cd146104365780637bc1ef61146104495780638ce6485414610470575f5ffd5b8063304c10cd116101955780634cfd2939116101655780634cfd29391461038a578063547afb87146103ab578063670d3ba2146103be57806367aeaa53146103d15780636cfb4481146103f8575f5ffd5b8063304c10cd1461031657806340120dab146103295780634177a87c1461034a5780634a10ffe51461036a575f5ffd5b8063260dc758116101d0578063260dc758146102945780632981eb77146102a75780632b453a9a146102e35780632bab2c4a14610303575f5ffd5b80630f3df50e1461020157806310e1b9b8146102315780631352c3e61461025157806315fe502814610274575b5f5ffd5b61021461020f366004611dd8565b6105ef565b6040516001600160a01b0390911681526020015b60405180910390f35b61024461023f366004611df2565b610630565b6040516102289190611e39565b61026461025f366004611e6c565b610669565b6040519015158152602001610228565b610287610282366004611ea0565b6106e4565b6040516102289190611f10565b6102646102a2366004611dd8565b6107fb565b6102ce7f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff9091168152602001610228565b6102f66102f1366004611fb7565b61082c565b604051610228919061202a565b6102f66103113660046120b6565b610842565b610214610324366004611ea0565b6108e1565b61033c61033736600461213a565b610910565b6040516102289291906121d2565b61035d610358366004611dd8565b610a89565b604051610228919061222f565b61037d610378366004612241565b610aad565b6040516102289190612284565b61039d610398366004611dd8565b610b55565b604051908152602001610228565b61037d6103b93660046122cf565b610b77565b6102646103cc366004611e6c565b610c1f565b6102ce7f000000000000000000000000000000000000000000000000000000000000000081565b61040b61040636600461213a565b610c4b565b6040516001600160401b039091168152602001610228565b61035d610431366004611dd8565b610c60565b610287610444366004611ea0565b610c71565b6102ce7f000000000000000000000000000000000000000000000000000000000000000081565b61048361047e366004612311565b610d4b565b6040516102289190612354565b61037d61049e366004612366565b610e07565b61040b6104b136600461213a565b610ef3565b61039d6104c4366004611dd8565b610f22565b6104dc6104d7366004611ea0565b610f44565b60408051921515835263ffffffff909116602083015201610228565b61039d610506366004611ea0565b610fe5565b61035d610519366004611e6c565b611005565b61021461052c366004611dd8565b61102e565b61054461053f366004611dd8565b6110b5565b604080516001600160a01b03909316835263ffffffff909116602083015201610228565b6102147f000000000000000000000000000000000000000000000000000000000000000081565b61026461059d366004611ea0565b611136565b6102147f000000000000000000000000000000000000000000000000000000000000000081565b6102646105d7366004611dd8565b611169565b61040b6105ea36600461213a565b611188565b5f5f60a65f6105fd85611194565b815260208101919091526040015f20546001600160a01b0316905080156106245780610629565b620e16e45b9392505050565b604080516060810182525f808252602082018190529181018290529061065f8561065986611194565b856111f7565b9695505050505050565b6001600160a01b0382165f908152609e6020526040812081908161068c85611194565b815260208082019290925260409081015f2081518083019092525460ff8116151580835261010090910463ffffffff16928201929092529150806106da5750806020015163ffffffff164311155b9150505b92915050565b6001600160a01b0381165f908152609d602052604081206060919061070890611363565b90505f816001600160401b0381111561072357610723611d0d565b60405190808252806020026020018201604052801561076757816020015b604080518082019091525f80825260208201528152602001906001900390816107415790505b5090505f5b828110156107f3576001600160a01b0385165f908152609d602052604090206107ce90610799908361136c565b604080518082019091525f80825260208201525060408051808201909152606082901c815263ffffffff909116602082015290565b8282815181106107e0576107e06123c1565b602090810291909101015260010161076c565b509392505050565b60208082015182516001600160a01b03165f9081526098909252604082206106de9163ffffffff9081169061137716565b606061083a8484844361138e565b949350505050565b60606108508585858561138e565b90505f5b84518110156108d857610880858281518110610872576108726123c1565b602002602001015187610669565b6108d0575f5b84518110156108ce575f8383815181106108a2576108a26123c1565b602002602001015182815181106108bb576108bb6123c1565b6020908102919091010152600101610886565b505b600101610854565b50949350505050565b6001600160a01b038082165f9081526097602052604081205490911680156109095780610629565b5090919050565b6001600160a01b0382165f908152609d60205260408120606091829161093590611363565b90505f816001600160401b0381111561095057610950611d0d565b60405190808252806020026020018201604052801561099457816020015b604080518082019091525f808252602082015281526020019060019003908161096e5790505b5090505f826001600160401b038111156109b0576109b0611d0d565b6040519080825280602002602001820160405280156109f957816020015b604080516060810182525f80825260208083018290529282015282525f199092019101816109ce5790505b5090505f5b83811015610a7c576001600160a01b0388165f908152609d60205260408120610a2b90610799908461136c565b905080848381518110610a4057610a406123c1565b6020026020010181905250610a5689828a610630565b838381518110610a6857610a686123c1565b6020908102919091010152506001016109fe565b5090969095509350505050565b60605f61062960995f610a9b86611194565b81526020019081526020015f2061167b565b60605f83516001600160401b03811115610ac957610ac9611d0d565b604051908082528060200260200182016040528015610af2578160200160208202803683370190505b5090505f5b84518110156107f357610b23858281518110610b1557610b156123c1565b602002602001015185610ef3565b828281518110610b3557610b356123c1565b6001600160401b0390921660209283029190910190910152600101610af7565b5f60a55f610b6284611194565b81526020019081526020015f20549050919050565b60605f82516001600160401b03811115610b9357610b93611d0d565b604051908082528060200260200182016040528015610bbc578160200160208202803683370190505b5090505f5b83518110156107f357610bed85858381518110610be057610be06123c1565b6020026020010151610ef3565b828281518110610bff57610bff6123c1565b6001600160401b0390921660209283029190910190910152600101610bc1565b5f61062983609a5f610c3086611194565b81526020019081526020015f2061168790919063ffffffff16565b5f5f610c5784846116a8565b95945050505050565b60606106de609a5f610a9b85611194565b6001600160a01b0381165f908152609c6020526040812060609190610c9590611363565b90505f816001600160401b03811115610cb057610cb0611d0d565b604051908082528060200260200182016040528015610cf457816020015b604080518082019091525f8082526020820152815260200190600190039081610cce5790505b5090505f5b828110156107f3576001600160a01b0385165f908152609c60205260409020610d2690610799908361136c565b828281518110610d3857610d386123c1565b6020908102919091010152600101610cf9565b60605f84516001600160401b03811115610d6757610d67611d0d565b604051908082528060200260200182016040528015610db057816020015b604080516060810182525f80825260208083018290529282015282525f19909201910181610d855790505b5090505f5b85518110156108d857610de2868281518110610dd357610dd36123c1565b60200260200101518686610630565b828281518110610df457610df46123c1565b6020908102919091010152600101610db5565b60605f83516001600160401b03811115610e2357610e23611d0d565b604051908082528060200260200182016040528015610e4c578160200160208202803683370190505b5090505f5b84518110156108d8576001600160a01b0386165f90815260a1602052604081208651610ec192879291899086908110610e8c57610e8c6123c1565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020015f2061181790919063ffffffff16565b828281518110610ed357610ed36123c1565b6001600160401b0390921660209283029190910190910152600101610e51565b6001600160a01b038083165f90815260a16020908152604080832093851683529290529081206106299061182b565b5f6106de609a5f610f3285611194565b81526020019081526020015f20611363565b6001600160a01b0381165f908152609b602090815260408083208151608081018352905463ffffffff80821680845260ff64010000000084041615159584018690526501000000000083048216948401949094526901000000000000000000909104166060820181905284939192919015801590610fcc5750826060015163ffffffff164310155b15610fdb575050604081015160015b9590945092505050565b6001600160a01b0381165f9081526098602052604081206106de90611363565b6001600160a01b0382165f908152609f60205260408120606091906106da9082610a9b86611194565b5f5f60a75f61103c85611194565b815260208082019290925260409081015f20815160608101835281546001600160a01b0390811680835260019093015490811694820194909452600160a01b90930463ffffffff16918301829052919250158015906110a55750816040015163ffffffff164310155b1561062957506020015192915050565b5f5f5f5f5f60a75f6110c688611194565b815260208082019290925260409081015f20815160608101835281546001600160a01b03908116825260019092015491821693810193909352600160a01b900463ffffffff1690820181905290915043101561112b5780602001519250806040015191505b509094909350915050565b5f5f61114183610c71565b90505f61114d846106e4565b9050611159848361183e565b8061083a575061083a848261183e565b5f620e16e4611177836105ef565b6001600160a01b0316141592915050565b5f5f6108d884846116a8565b5f815f0151826020015163ffffffff166040516020016111df92919060609290921b6bffffffffffffffffffffffff1916825260a01b6001600160a01b031916601482015260200190565b6040516020818303038152906040526106de906123d5565b6040805180820182525f80825260208083018290528351606081018552828152808201839052808501839052845180860186526001600160a01b03898116855260a18452868520908816855290925293822092939281906112579061182b565b6001600160401b0390811682526001600160a01b038981165f81815260a260209081526040808320948c168084529482528083205486169682019690965291815260a082528481208b8252825284812092815291815290839020835160608101855290549283168152600160401b8304600f0b91810191909152600160c01b90910463ffffffff169181018290529192504310156112f957909250905061135b565b61130a815f015182602001516118b6565b6001600160401b0316815260208101515f600f9190910b121561134857611339826020015182602001516118b6565b6001600160401b031660208301525b5f60408201819052602082015290925090505b935093915050565b5f6106de825490565b5f61062983836118d5565b5f8181526001830160205260408120541515610629565b606083516001600160401b038111156113a9576113a9611d0d565b6040519080825280602002602001820160405280156113dc57816020015b60608152602001906001900390816113c75790505b5090505f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f0e0e67686866040518363ffffffff1660e01b815260040161142e9291906123fb565b5f60405180830381865afa158015611448573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261146f919081019061241f565b90505f5b8551811015611671575f86828151811061148f5761148f6123c1565b6020026020010151905085516001600160401b038111156114b2576114b2611d0d565b6040519080825280602002602001820160405280156114db578160200160208202803683370190505b508483815181106114ee576114ee6123c1565b60209081029190910101525f5b8651811015611667575f878281518110611517576115176123c1565b6020908102919091018101516001600160a01b038086165f90815260a18452604080822092841682529190935282209092506115529061182b565b9050806001600160401b03165f0361156b57505061165f565b5f611577858d85610630565b90508863ffffffff16816040015163ffffffff161115801561159f57505f8160200151600f0b125b156115c1576115b5815f015182602001516118b6565b6001600160401b031681525b80515f906115dc906001600160401b039081169085166118fb565b9050611623818989815181106115f4576115f46123c1565b6020026020010151878151811061160d5761160d6123c1565b602002602001015161190f90919063ffffffff16565b898881518110611635576116356123c1565b6020026020010151868151811061164e5761164e6123c1565b602002602001018181525050505050505b6001016114fb565b5050600101611473565b5050949350505050565b60605f61062983611923565b6001600160a01b0381165f9081526001830160205260408120541515610629565b6001600160a01b038281165f81815260a2602090815260408083209486168084529482528083205493835260a38252808320948352939052918220546001600160401b039091169190600f81810b600160801b909204900b03815b818110156117d3576001600160a01b038087165f90815260a360209081526040808320938916835292905290812061173b908361197c565b6001600160a01b038881165f90815260a0602090815260408083208584528252808320938b16835292815290829020825160608101845290546001600160401b0381168252600160401b8104600f0b92820192909252600160c01b90910463ffffffff169181018290529192504310156117b65750506117d3565b6117c48682602001516118b6565b95505050806001019050611703565b506001600160a01b038086165f90815260a16020908152604080832093881683529290522083906118039061182b565b61180d919061253f565b9150509250929050565b5f6106298383670de0b6b3a76400006119eb565b5f6106de82670de0b6b3a7640000611a41565b5f805b82518110156118ad5761186d84848381518110611860576118606123c1565b6020026020010151610669565b80156118965750611896838281518110611889576118896123c1565b6020026020010151611169565b156118a55760019150506106de565b600101611841565b505f9392505050565b5f6106296118cd836001600160401b03861661255e565b600f0b611a79565b5f825f0182815481106118ea576118ea6123c1565b905f5260205f200154905092915050565b5f61062983670de0b6b3a764000084611ae9565b5f6106298383670de0b6b3a7640000611ae9565b6060815f0180548060200260200160405190810160405280929190818152602001828054801561197057602002820191905f5260205f20905b81548152602001906001019080831161195c575b50505050509050919050565b5f5f61199e61198a84611bce565b85546119999190600f0b61259d565b611c37565b8454909150600160801b9004600f90810b9082900b126119d157604051632d0483c560e21b815260040160405180910390fd5b600f0b5f9081526001939093016020525050604090205490565b82545f90816119fc86868385611ca0565b90508015611a3757611a2086611a136001846125c4565b5f91825260209091200190565b5464010000000090046001600160e01b031661065f565b5091949350505050565b81545f908015611a7157611a5a84611a136001846125c4565b5464010000000090046001600160e01b03166106da565b509092915050565b5f6001600160401b03821115611ae55760405162461bcd60e51b815260206004820152602660248201527f53616665436173743a2076616c756520646f65736e27742066697420696e203660448201526534206269747360d01b60648201526084015b60405180910390fd5b5090565b5f80805f19858709858702925082811083820303915050805f03611b2057838281611b1657611b166125d7565b0492505050610629565b808411611b675760405162461bcd60e51b81526020600482015260156024820152744d6174683a206d756c446976206f766572666c6f7760581b6044820152606401611adc565b5f8486880960026001871981018816978890046003810283188082028403028082028403028082028403028082028403028082028403029081029092039091025f889003889004909101858311909403939093029303949094049190911702949350505050565b5f6001600160ff1b03821115611ae55760405162461bcd60e51b815260206004820152602860248201527f53616665436173743a2076616c756520646f65736e27742066697420696e2061604482015267371034b73a191a9b60c11b6064820152608401611adc565b80600f81900b8114611c9b5760405162461bcd60e51b815260206004820152602760248201527f53616665436173743a2076616c756520646f65736e27742066697420696e20316044820152663238206269747360c81b6064820152608401611adc565b919050565b5f5b818310156107f3575f611cb58484611cf3565b5f8781526020902090915063ffffffff86169082015463ffffffff161115611cdf57809250611ced565b611cea8160016125eb565b93505b50611ca2565b5f611d0160028484186125fe565b610629908484166125eb565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f191681016001600160401b0381118282101715611d4957611d49611d0d565b604052919050565b6001600160a01b0381168114611d65575f5ffd5b50565b803563ffffffff81168114611c9b575f5ffd5b5f60408284031215611d8b575f5ffd5b604080519081016001600160401b0381118282101715611dad57611dad611d0d565b6040529050808235611dbe81611d51565b8152611dcc60208401611d68565b60208201525092915050565b5f60408284031215611de8575f5ffd5b6106298383611d7b565b5f5f5f60808486031215611e04575f5ffd5b8335611e0f81611d51565b9250611e1e8560208601611d7b565b91506060840135611e2e81611d51565b809150509250925092565b81516001600160401b03168152602080830151600f0b9082015260408083015163ffffffff1690820152606081016106de565b5f5f60608385031215611e7d575f5ffd5b8235611e8881611d51565b9150611e978460208501611d7b565b90509250929050565b5f60208284031215611eb0575f5ffd5b813561062981611d51565b5f8151808452602084019350602083015f5b82811015611f0657815180516001600160a01b0316875260209081015163ffffffff168188015260409096019590910190600101611ecd565b5093949350505050565b602081525f6106296020830184611ebb565b5f6001600160401b03821115611f3a57611f3a611d0d565b5060051b60200190565b5f82601f830112611f53575f5ffd5b8135611f66611f6182611f22565b611d21565b8082825260208201915060208360051b860101925085831115611f87575f5ffd5b602085015b83811015611fad578035611f9f81611d51565b835260209283019201611f8c565b5095945050505050565b5f5f5f60808486031215611fc9575f5ffd5b611fd38585611d7b565b925060408401356001600160401b03811115611fed575f5ffd5b611ff986828701611f44565b92505060608401356001600160401b03811115612014575f5ffd5b61202086828701611f44565b9150509250925092565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b828110156120aa57868503603f19018452815180518087526020918201918701905f5b81811015612091578351835260209384019390920191600101612073565b5090965050506020938401939190910190600101612050565b50929695505050505050565b5f5f5f5f60a085870312156120c9575f5ffd5b6120d38686611d7b565b935060408501356001600160401b038111156120ed575f5ffd5b6120f987828801611f44565b93505060608501356001600160401b03811115612114575f5ffd5b61212087828801611f44565b92505061212f60808601611d68565b905092959194509250565b5f5f6040838503121561214b575f5ffd5b823561215681611d51565b9150602083013561216681611d51565b809150509250929050565b5f8151808452602084019350602083015f5b82811015611f06576121bc86835180516001600160401b03168252602080820151600f0b9083015260409081015163ffffffff16910152565b6060959095019460209190910190600101612183565b604081525f6121e46040830185611ebb565b8281036020840152610c578185612171565b5f8151808452602084019350602083015f5b82811015611f065781516001600160a01b0316865260209586019590910190600101612208565b602081525f61062960208301846121f6565b5f5f60408385031215612252575f5ffd5b82356001600160401b03811115612267575f5ffd5b61227385828601611f44565b925050602083013561216681611d51565b602080825282518282018190525f918401906040840190835b818110156122c45783516001600160401b031683526020938401939092019160010161229d565b509095945050505050565b5f5f604083850312156122e0575f5ffd5b82356122eb81611d51565b915060208301356001600160401b03811115612305575f5ffd5b61180d85828601611f44565b5f5f5f60808486031215612323575f5ffd5b83356001600160401b03811115612338575f5ffd5b61234486828701611f44565b935050611e1e8560208601611d7b565b602081525f6106296020830184612171565b5f5f5f60608486031215612378575f5ffd5b833561238381611d51565b925060208401356001600160401b0381111561239d575f5ffd5b6123a986828701611f44565b9250506123b860408501611d68565b90509250925092565b634e487b7160e01b5f52603260045260245ffd5b805160208083015191908110156123f5575f198160200360031b1b821691505b50919050565b604081525f61240d60408301856121f6565b8281036020840152610c5781856121f6565b5f6020828403121561242f575f5ffd5b81516001600160401b03811115612444575f5ffd5b8201601f81018413612454575f5ffd5b8051612462611f6182611f22565b8082825260208201915060208360051b850101925086831115612483575f5ffd5b602084015b838110156125205780516001600160401b038111156124a5575f5ffd5b8501603f810189136124b5575f5ffd5b60208101516124c6611f6182611f22565b808282526020820191506020808460051b8601010192508b8311156124e9575f5ffd5b6040840193505b8284101561250b5783518252602093840193909101906124f0565b86525050602093840193919091019050612488565b509695505050505050565b634e487b7160e01b5f52601160045260245ffd5b6001600160401b0382811682821603908111156106de576106de61252b565b600f81810b9083900b016f7fffffffffffffffffffffffffffffff81136f7fffffffffffffffffffffffffffffff19821217156106de576106de61252b565b8082018281125f8312801582168215821617156125bc576125bc61252b565b505092915050565b818103818111156106de576106de61252b565b634e487b7160e01b5f52601260045260245ffd5b808201808211156106de576106de61252b565b5f8261261857634e487b7160e01b5f52601260045260245ffd5b50049056fea2646970667358221220800066159b94ad8279c77e9ea5168403873dfa0fd7176376207ccdac08ea7a9664736f6c634300081e0033",
}

// AllocationManagerViewABI is the input ABI used to generate the binding from.
// Deprecated: Use AllocationManagerViewMetaData.ABI instead.
var AllocationManagerViewABI = AllocationManagerViewMetaData.ABI

// AllocationManagerViewBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AllocationManagerViewMetaData.Bin instead.
var AllocationManagerViewBin = AllocationManagerViewMetaData.Bin

// DeployAllocationManagerView deploys a new Ethereum contract, binding an instance of AllocationManagerView to it.
func DeployAllocationManagerView(auth *bind.TransactOpts, backend bind.ContractBackend, _delegation common.Address, _eigenStrategy common.Address, _DEALLOCATION_DELAY uint32, _ALLOCATION_CONFIGURATION_DELAY uint32) (common.Address, *types.Transaction, *AllocationManagerView, error) {
	parsed, err := AllocationManagerViewMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AllocationManagerViewBin), backend, _delegation, _eigenStrategy, _DEALLOCATION_DELAY, _ALLOCATION_CONFIGURATION_DELAY)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AllocationManagerView{AllocationManagerViewCaller: AllocationManagerViewCaller{contract: contract}, AllocationManagerViewTransactor: AllocationManagerViewTransactor{contract: contract}, AllocationManagerViewFilterer: AllocationManagerViewFilterer{contract: contract}}, nil
}

// AllocationManagerView is an auto generated Go binding around an Ethereum contract.
type AllocationManagerView struct {
	AllocationManagerViewCaller     // Read-only binding to the contract
	AllocationManagerViewTransactor // Write-only binding to the contract
	AllocationManagerViewFilterer   // Log filterer for contract events
}

// AllocationManagerViewCaller is an auto generated read-only Go binding around an Ethereum contract.
type AllocationManagerViewCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AllocationManagerViewTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AllocationManagerViewTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AllocationManagerViewFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AllocationManagerViewFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AllocationManagerViewSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AllocationManagerViewSession struct {
	Contract     *AllocationManagerView // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// AllocationManagerViewCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AllocationManagerViewCallerSession struct {
	Contract *AllocationManagerViewCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// AllocationManagerViewTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AllocationManagerViewTransactorSession struct {
	Contract     *AllocationManagerViewTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// AllocationManagerViewRaw is an auto generated low-level Go binding around an Ethereum contract.
type AllocationManagerViewRaw struct {
	Contract *AllocationManagerView // Generic contract binding to access the raw methods on
}

// AllocationManagerViewCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AllocationManagerViewCallerRaw struct {
	Contract *AllocationManagerViewCaller // Generic read-only contract binding to access the raw methods on
}

// AllocationManagerViewTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AllocationManagerViewTransactorRaw struct {
	Contract *AllocationManagerViewTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAllocationManagerView creates a new instance of AllocationManagerView, bound to a specific deployed contract.
func NewAllocationManagerView(address common.Address, backend bind.ContractBackend) (*AllocationManagerView, error) {
	contract, err := bindAllocationManagerView(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AllocationManagerView{AllocationManagerViewCaller: AllocationManagerViewCaller{contract: contract}, AllocationManagerViewTransactor: AllocationManagerViewTransactor{contract: contract}, AllocationManagerViewFilterer: AllocationManagerViewFilterer{contract: contract}}, nil
}

// NewAllocationManagerViewCaller creates a new read-only instance of AllocationManagerView, bound to a specific deployed contract.
func NewAllocationManagerViewCaller(address common.Address, caller bind.ContractCaller) (*AllocationManagerViewCaller, error) {
	contract, err := bindAllocationManagerView(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AllocationManagerViewCaller{contract: contract}, nil
}

// NewAllocationManagerViewTransactor creates a new write-only instance of AllocationManagerView, bound to a specific deployed contract.
func NewAllocationManagerViewTransactor(address common.Address, transactor bind.ContractTransactor) (*AllocationManagerViewTransactor, error) {
	contract, err := bindAllocationManagerView(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AllocationManagerViewTransactor{contract: contract}, nil
}

// NewAllocationManagerViewFilterer creates a new log filterer instance of AllocationManagerView, bound to a specific deployed contract.
func NewAllocationManagerViewFilterer(address common.Address, filterer bind.ContractFilterer) (*AllocationManagerViewFilterer, error) {
	contract, err := bindAllocationManagerView(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AllocationManagerViewFilterer{contract: contract}, nil
}

// bindAllocationManagerView binds a generic wrapper to an already deployed contract.
func bindAllocationManagerView(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AllocationManagerViewMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AllocationManagerView *AllocationManagerViewRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AllocationManagerView.Contract.AllocationManagerViewCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AllocationManagerView *AllocationManagerViewRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AllocationManagerView.Contract.AllocationManagerViewTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AllocationManagerView *AllocationManagerViewRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AllocationManagerView.Contract.AllocationManagerViewTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AllocationManagerView *AllocationManagerViewCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AllocationManagerView.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AllocationManagerView *AllocationManagerViewTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AllocationManagerView.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AllocationManagerView *AllocationManagerViewTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AllocationManagerView.Contract.contract.Transact(opts, method, params...)
}

// ALLOCATIONCONFIGURATIONDELAY is a free data retrieval call binding the contract method 0x7bc1ef61.
//
// Solidity: function ALLOCATION_CONFIGURATION_DELAY() view returns(uint32)
func (_AllocationManagerView *AllocationManagerViewCaller) ALLOCATIONCONFIGURATIONDELAY(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _AllocationManagerView.contract.Call(opts, &out, "ALLOCATION_CONFIGURATION_DELAY")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ALLOCATIONCONFIGURATIONDELAY is a free data retrieval call binding the contract method 0x7bc1ef61.
//
// Solidity: function ALLOCATION_CONFIGURATION_DELAY() view returns(uint32)
func (_AllocationManagerView *AllocationManagerViewSession) ALLOCATIONCONFIGURATIONDELAY() (uint32, error) {
	return _AllocationManagerView.Contract.ALLOCATIONCONFIGURATIONDELAY(&_AllocationManagerView.CallOpts)
}

// ALLOCATIONCONFIGURATIONDELAY is a free data retrieval call binding the contract method 0x7bc1ef61.
//
// Solidity: function ALLOCATION_CONFIGURATION_DELAY() view returns(uint32)
func (_AllocationManagerView *AllocationManagerViewCallerSession) ALLOCATIONCONFIGURATIONDELAY() (uint32, error) {
	return _AllocationManagerView.Contract.ALLOCATIONCONFIGURATIONDELAY(&_AllocationManagerView.CallOpts)
}

// DEALLOCATIONDELAY is a free data retrieval call binding the contract method 0x2981eb77.
//
// Solidity: function DEALLOCATION_DELAY() view returns(uint32)
func (_AllocationManagerView *AllocationManagerViewCaller) DEALLOCATIONDELAY(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _AllocationManagerView.contract.Call(opts, &out, "DEALLOCATION_DELAY")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// DEALLOCATIONDELAY is a free data retrieval call binding the contract method 0x2981eb77.
//
// Solidity: function DEALLOCATION_DELAY() view returns(uint32)
func (_AllocationManagerView *AllocationManagerViewSession) DEALLOCATIONDELAY() (uint32, error) {
	return _AllocationManagerView.Contract.DEALLOCATIONDELAY(&_AllocationManagerView.CallOpts)
}

// DEALLOCATIONDELAY is a free data retrieval call binding the contract method 0x2981eb77.
//
// Solidity: function DEALLOCATION_DELAY() view returns(uint32)
func (_AllocationManagerView *AllocationManagerViewCallerSession) DEALLOCATIONDELAY() (uint32, error) {
	return _AllocationManagerView.Contract.DEALLOCATIONDELAY(&_AllocationManagerView.CallOpts)
}

// SLASHERCONFIGURATIONDELAY is a free data retrieval call binding the contract method 0x67aeaa53.
//
// Solidity: function SLASHER_CONFIGURATION_DELAY() view returns(uint32)
func (_AllocationManagerView *AllocationManagerViewCaller) SLASHERCONFIGURATIONDELAY(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _AllocationManagerView.contract.Call(opts, &out, "SLASHER_CONFIGURATION_DELAY")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// SLASHERCONFIGURATIONDELAY is a free data retrieval call binding the contract method 0x67aeaa53.
//
// Solidity: function SLASHER_CONFIGURATION_DELAY() view returns(uint32)
func (_AllocationManagerView *AllocationManagerViewSession) SLASHERCONFIGURATIONDELAY() (uint32, error) {
	return _AllocationManagerView.Contract.SLASHERCONFIGURATIONDELAY(&_AllocationManagerView.CallOpts)
}

// SLASHERCONFIGURATIONDELAY is a free data retrieval call binding the contract method 0x67aeaa53.
//
// Solidity: function SLASHER_CONFIGURATION_DELAY() view returns(uint32)
func (_AllocationManagerView *AllocationManagerViewCallerSession) SLASHERCONFIGURATIONDELAY() (uint32, error) {
	return _AllocationManagerView.Contract.SLASHERCONFIGURATIONDELAY(&_AllocationManagerView.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_AllocationManagerView *AllocationManagerViewCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AllocationManagerView.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_AllocationManagerView *AllocationManagerViewSession) Delegation() (common.Address, error) {
	return _AllocationManagerView.Contract.Delegation(&_AllocationManagerView.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_AllocationManagerView *AllocationManagerViewCallerSession) Delegation() (common.Address, error) {
	return _AllocationManagerView.Contract.Delegation(&_AllocationManagerView.CallOpts)
}

// EigenStrategy is a free data retrieval call binding the contract method 0xdb4df761.
//
// Solidity: function eigenStrategy() view returns(address)
func (_AllocationManagerView *AllocationManagerViewCaller) EigenStrategy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AllocationManagerView.contract.Call(opts, &out, "eigenStrategy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EigenStrategy is a free data retrieval call binding the contract method 0xdb4df761.
//
// Solidity: function eigenStrategy() view returns(address)
func (_AllocationManagerView *AllocationManagerViewSession) EigenStrategy() (common.Address, error) {
	return _AllocationManagerView.Contract.EigenStrategy(&_AllocationManagerView.CallOpts)
}

// EigenStrategy is a free data retrieval call binding the contract method 0xdb4df761.
//
// Solidity: function eigenStrategy() view returns(address)
func (_AllocationManagerView *AllocationManagerViewCallerSession) EigenStrategy() (common.Address, error) {
	return _AllocationManagerView.Contract.EigenStrategy(&_AllocationManagerView.CallOpts)
}

// GetAVSRegistrar is a free data retrieval call binding the contract method 0x304c10cd.
//
// Solidity: function getAVSRegistrar(address avs) view returns(address)
func (_AllocationManagerView *AllocationManagerViewCaller) GetAVSRegistrar(opts *bind.CallOpts, avs common.Address) (common.Address, error) {
	var out []interface{}
	err := _AllocationManagerView.contract.Call(opts, &out, "getAVSRegistrar", avs)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAVSRegistrar is a free data retrieval call binding the contract method 0x304c10cd.
//
// Solidity: function getAVSRegistrar(address avs) view returns(address)
func (_AllocationManagerView *AllocationManagerViewSession) GetAVSRegistrar(avs common.Address) (common.Address, error) {
	return _AllocationManagerView.Contract.GetAVSRegistrar(&_AllocationManagerView.CallOpts, avs)
}

// GetAVSRegistrar is a free data retrieval call binding the contract method 0x304c10cd.
//
// Solidity: function getAVSRegistrar(address avs) view returns(address)
func (_AllocationManagerView *AllocationManagerViewCallerSession) GetAVSRegistrar(avs common.Address) (common.Address, error) {
	return _AllocationManagerView.Contract.GetAVSRegistrar(&_AllocationManagerView.CallOpts, avs)
}

// GetAllocatableMagnitude is a free data retrieval call binding the contract method 0x6cfb4481.
//
// Solidity: function getAllocatableMagnitude(address operator, address strategy) view returns(uint64)
func (_AllocationManagerView *AllocationManagerViewCaller) GetAllocatableMagnitude(opts *bind.CallOpts, operator common.Address, strategy common.Address) (uint64, error) {
	var out []interface{}
	err := _AllocationManagerView.contract.Call(opts, &out, "getAllocatableMagnitude", operator, strategy)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetAllocatableMagnitude is a free data retrieval call binding the contract method 0x6cfb4481.
//
// Solidity: function getAllocatableMagnitude(address operator, address strategy) view returns(uint64)
func (_AllocationManagerView *AllocationManagerViewSession) GetAllocatableMagnitude(operator common.Address, strategy common.Address) (uint64, error) {
	return _AllocationManagerView.Contract.GetAllocatableMagnitude(&_AllocationManagerView.CallOpts, operator, strategy)
}

// GetAllocatableMagnitude is a free data retrieval call binding the contract method 0x6cfb4481.
//
// Solidity: function getAllocatableMagnitude(address operator, address strategy) view returns(uint64)
func (_AllocationManagerView *AllocationManagerViewCallerSession) GetAllocatableMagnitude(operator common.Address, strategy common.Address) (uint64, error) {
	return _AllocationManagerView.Contract.GetAllocatableMagnitude(&_AllocationManagerView.CallOpts, operator, strategy)
}

// GetAllocatedSets is a free data retrieval call binding the contract method 0x15fe5028.
//
// Solidity: function getAllocatedSets(address operator) view returns((address,uint32)[])
func (_AllocationManagerView *AllocationManagerViewCaller) GetAllocatedSets(opts *bind.CallOpts, operator common.Address) ([]OperatorSet, error) {
	var out []interface{}
	err := _AllocationManagerView.contract.Call(opts, &out, "getAllocatedSets", operator)

	if err != nil {
		return *new([]OperatorSet), err
	}

	out0 := *abi.ConvertType(out[0], new([]OperatorSet)).(*[]OperatorSet)

	return out0, err

}

// GetAllocatedSets is a free data retrieval call binding the contract method 0x15fe5028.
//
// Solidity: function getAllocatedSets(address operator) view returns((address,uint32)[])
func (_AllocationManagerView *AllocationManagerViewSession) GetAllocatedSets(operator common.Address) ([]OperatorSet, error) {
	return _AllocationManagerView.Contract.GetAllocatedSets(&_AllocationManagerView.CallOpts, operator)
}

// GetAllocatedSets is a free data retrieval call binding the contract method 0x15fe5028.
//
// Solidity: function getAllocatedSets(address operator) view returns((address,uint32)[])
func (_AllocationManagerView *AllocationManagerViewCallerSession) GetAllocatedSets(operator common.Address) ([]OperatorSet, error) {
	return _AllocationManagerView.Contract.GetAllocatedSets(&_AllocationManagerView.CallOpts, operator)
}

// GetAllocatedStake is a free data retrieval call binding the contract method 0x2b453a9a.
//
// Solidity: function getAllocatedStake((address,uint32) operatorSet, address[] operators, address[] strategies) view returns(uint256[][])
func (_AllocationManagerView *AllocationManagerViewCaller) GetAllocatedStake(opts *bind.CallOpts, operatorSet OperatorSet, operators []common.Address, strategies []common.Address) ([][]*big.Int, error) {
	var out []interface{}
	err := _AllocationManagerView.contract.Call(opts, &out, "getAllocatedStake", operatorSet, operators, strategies)

	if err != nil {
		return *new([][]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([][]*big.Int)).(*[][]*big.Int)

	return out0, err

}

// GetAllocatedStake is a free data retrieval call binding the contract method 0x2b453a9a.
//
// Solidity: function getAllocatedStake((address,uint32) operatorSet, address[] operators, address[] strategies) view returns(uint256[][])
func (_AllocationManagerView *AllocationManagerViewSession) GetAllocatedStake(operatorSet OperatorSet, operators []common.Address, strategies []common.Address) ([][]*big.Int, error) {
	return _AllocationManagerView.Contract.GetAllocatedStake(&_AllocationManagerView.CallOpts, operatorSet, operators, strategies)
}

// GetAllocatedStake is a free data retrieval call binding the contract method 0x2b453a9a.
//
// Solidity: function getAllocatedStake((address,uint32) operatorSet, address[] operators, address[] strategies) view returns(uint256[][])
func (_AllocationManagerView *AllocationManagerViewCallerSession) GetAllocatedStake(operatorSet OperatorSet, operators []common.Address, strategies []common.Address) ([][]*big.Int, error) {
	return _AllocationManagerView.Contract.GetAllocatedStake(&_AllocationManagerView.CallOpts, operatorSet, operators, strategies)
}

// GetAllocatedStrategies is a free data retrieval call binding the contract method 0xc221d8ae.
//
// Solidity: function getAllocatedStrategies(address operator, (address,uint32) operatorSet) view returns(address[])
func (_AllocationManagerView *AllocationManagerViewCaller) GetAllocatedStrategies(opts *bind.CallOpts, operator common.Address, operatorSet OperatorSet) ([]common.Address, error) {
	var out []interface{}
	err := _AllocationManagerView.contract.Call(opts, &out, "getAllocatedStrategies", operator, operatorSet)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllocatedStrategies is a free data retrieval call binding the contract method 0xc221d8ae.
//
// Solidity: function getAllocatedStrategies(address operator, (address,uint32) operatorSet) view returns(address[])
func (_AllocationManagerView *AllocationManagerViewSession) GetAllocatedStrategies(operator common.Address, operatorSet OperatorSet) ([]common.Address, error) {
	return _AllocationManagerView.Contract.GetAllocatedStrategies(&_AllocationManagerView.CallOpts, operator, operatorSet)
}

// GetAllocatedStrategies is a free data retrieval call binding the contract method 0xc221d8ae.
//
// Solidity: function getAllocatedStrategies(address operator, (address,uint32) operatorSet) view returns(address[])
func (_AllocationManagerView *AllocationManagerViewCallerSession) GetAllocatedStrategies(operator common.Address, operatorSet OperatorSet) ([]common.Address, error) {
	return _AllocationManagerView.Contract.GetAllocatedStrategies(&_AllocationManagerView.CallOpts, operator, operatorSet)
}

// GetAllocation is a free data retrieval call binding the contract method 0x10e1b9b8.
//
// Solidity: function getAllocation(address operator, (address,uint32) operatorSet, address strategy) view returns((uint64,int128,uint32))
func (_AllocationManagerView *AllocationManagerViewCaller) GetAllocation(opts *bind.CallOpts, operator common.Address, operatorSet OperatorSet, strategy common.Address) (IAllocationManagerTypesAllocation, error) {
	var out []interface{}
	err := _AllocationManagerView.contract.Call(opts, &out, "getAllocation", operator, operatorSet, strategy)

	if err != nil {
		return *new(IAllocationManagerTypesAllocation), err
	}

	out0 := *abi.ConvertType(out[0], new(IAllocationManagerTypesAllocation)).(*IAllocationManagerTypesAllocation)

	return out0, err

}

// GetAllocation is a free data retrieval call binding the contract method 0x10e1b9b8.
//
// Solidity: function getAllocation(address operator, (address,uint32) operatorSet, address strategy) view returns((uint64,int128,uint32))
func (_AllocationManagerView *AllocationManagerViewSession) GetAllocation(operator common.Address, operatorSet OperatorSet, strategy common.Address) (IAllocationManagerTypesAllocation, error) {
	return _AllocationManagerView.Contract.GetAllocation(&_AllocationManagerView.CallOpts, operator, operatorSet, strategy)
}

// GetAllocation is a free data retrieval call binding the contract method 0x10e1b9b8.
//
// Solidity: function getAllocation(address operator, (address,uint32) operatorSet, address strategy) view returns((uint64,int128,uint32))
func (_AllocationManagerView *AllocationManagerViewCallerSession) GetAllocation(operator common.Address, operatorSet OperatorSet, strategy common.Address) (IAllocationManagerTypesAllocation, error) {
	return _AllocationManagerView.Contract.GetAllocation(&_AllocationManagerView.CallOpts, operator, operatorSet, strategy)
}

// GetAllocationDelay is a free data retrieval call binding the contract method 0xb9fbaed1.
//
// Solidity: function getAllocationDelay(address operator) view returns(bool, uint32)
func (_AllocationManagerView *AllocationManagerViewCaller) GetAllocationDelay(opts *bind.CallOpts, operator common.Address) (bool, uint32, error) {
	var out []interface{}
	err := _AllocationManagerView.contract.Call(opts, &out, "getAllocationDelay", operator)

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
func (_AllocationManagerView *AllocationManagerViewSession) GetAllocationDelay(operator common.Address) (bool, uint32, error) {
	return _AllocationManagerView.Contract.GetAllocationDelay(&_AllocationManagerView.CallOpts, operator)
}

// GetAllocationDelay is a free data retrieval call binding the contract method 0xb9fbaed1.
//
// Solidity: function getAllocationDelay(address operator) view returns(bool, uint32)
func (_AllocationManagerView *AllocationManagerViewCallerSession) GetAllocationDelay(operator common.Address) (bool, uint32, error) {
	return _AllocationManagerView.Contract.GetAllocationDelay(&_AllocationManagerView.CallOpts, operator)
}

// GetAllocations is a free data retrieval call binding the contract method 0x8ce64854.
//
// Solidity: function getAllocations(address[] operators, (address,uint32) operatorSet, address strategy) view returns((uint64,int128,uint32)[])
func (_AllocationManagerView *AllocationManagerViewCaller) GetAllocations(opts *bind.CallOpts, operators []common.Address, operatorSet OperatorSet, strategy common.Address) ([]IAllocationManagerTypesAllocation, error) {
	var out []interface{}
	err := _AllocationManagerView.contract.Call(opts, &out, "getAllocations", operators, operatorSet, strategy)

	if err != nil {
		return *new([]IAllocationManagerTypesAllocation), err
	}

	out0 := *abi.ConvertType(out[0], new([]IAllocationManagerTypesAllocation)).(*[]IAllocationManagerTypesAllocation)

	return out0, err

}

// GetAllocations is a free data retrieval call binding the contract method 0x8ce64854.
//
// Solidity: function getAllocations(address[] operators, (address,uint32) operatorSet, address strategy) view returns((uint64,int128,uint32)[])
func (_AllocationManagerView *AllocationManagerViewSession) GetAllocations(operators []common.Address, operatorSet OperatorSet, strategy common.Address) ([]IAllocationManagerTypesAllocation, error) {
	return _AllocationManagerView.Contract.GetAllocations(&_AllocationManagerView.CallOpts, operators, operatorSet, strategy)
}

// GetAllocations is a free data retrieval call binding the contract method 0x8ce64854.
//
// Solidity: function getAllocations(address[] operators, (address,uint32) operatorSet, address strategy) view returns((uint64,int128,uint32)[])
func (_AllocationManagerView *AllocationManagerViewCallerSession) GetAllocations(operators []common.Address, operatorSet OperatorSet, strategy common.Address) ([]IAllocationManagerTypesAllocation, error) {
	return _AllocationManagerView.Contract.GetAllocations(&_AllocationManagerView.CallOpts, operators, operatorSet, strategy)
}

// GetEncumberedMagnitude is a free data retrieval call binding the contract method 0xf605ce08.
//
// Solidity: function getEncumberedMagnitude(address operator, address strategy) view returns(uint64)
func (_AllocationManagerView *AllocationManagerViewCaller) GetEncumberedMagnitude(opts *bind.CallOpts, operator common.Address, strategy common.Address) (uint64, error) {
	var out []interface{}
	err := _AllocationManagerView.contract.Call(opts, &out, "getEncumberedMagnitude", operator, strategy)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetEncumberedMagnitude is a free data retrieval call binding the contract method 0xf605ce08.
//
// Solidity: function getEncumberedMagnitude(address operator, address strategy) view returns(uint64)
func (_AllocationManagerView *AllocationManagerViewSession) GetEncumberedMagnitude(operator common.Address, strategy common.Address) (uint64, error) {
	return _AllocationManagerView.Contract.GetEncumberedMagnitude(&_AllocationManagerView.CallOpts, operator, strategy)
}

// GetEncumberedMagnitude is a free data retrieval call binding the contract method 0xf605ce08.
//
// Solidity: function getEncumberedMagnitude(address operator, address strategy) view returns(uint64)
func (_AllocationManagerView *AllocationManagerViewCallerSession) GetEncumberedMagnitude(operator common.Address, strategy common.Address) (uint64, error) {
	return _AllocationManagerView.Contract.GetEncumberedMagnitude(&_AllocationManagerView.CallOpts, operator, strategy)
}

// GetMaxMagnitude is a free data retrieval call binding the contract method 0xa9333ec8.
//
// Solidity: function getMaxMagnitude(address operator, address strategy) view returns(uint64)
func (_AllocationManagerView *AllocationManagerViewCaller) GetMaxMagnitude(opts *bind.CallOpts, operator common.Address, strategy common.Address) (uint64, error) {
	var out []interface{}
	err := _AllocationManagerView.contract.Call(opts, &out, "getMaxMagnitude", operator, strategy)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetMaxMagnitude is a free data retrieval call binding the contract method 0xa9333ec8.
//
// Solidity: function getMaxMagnitude(address operator, address strategy) view returns(uint64)
func (_AllocationManagerView *AllocationManagerViewSession) GetMaxMagnitude(operator common.Address, strategy common.Address) (uint64, error) {
	return _AllocationManagerView.Contract.GetMaxMagnitude(&_AllocationManagerView.CallOpts, operator, strategy)
}

// GetMaxMagnitude is a free data retrieval call binding the contract method 0xa9333ec8.
//
// Solidity: function getMaxMagnitude(address operator, address strategy) view returns(uint64)
func (_AllocationManagerView *AllocationManagerViewCallerSession) GetMaxMagnitude(operator common.Address, strategy common.Address) (uint64, error) {
	return _AllocationManagerView.Contract.GetMaxMagnitude(&_AllocationManagerView.CallOpts, operator, strategy)
}

// GetMaxMagnitudes is a free data retrieval call binding the contract method 0x4a10ffe5.
//
// Solidity: function getMaxMagnitudes(address[] operators, address strategy) view returns(uint64[])
func (_AllocationManagerView *AllocationManagerViewCaller) GetMaxMagnitudes(opts *bind.CallOpts, operators []common.Address, strategy common.Address) ([]uint64, error) {
	var out []interface{}
	err := _AllocationManagerView.contract.Call(opts, &out, "getMaxMagnitudes", operators, strategy)

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

// GetMaxMagnitudes is a free data retrieval call binding the contract method 0x4a10ffe5.
//
// Solidity: function getMaxMagnitudes(address[] operators, address strategy) view returns(uint64[])
func (_AllocationManagerView *AllocationManagerViewSession) GetMaxMagnitudes(operators []common.Address, strategy common.Address) ([]uint64, error) {
	return _AllocationManagerView.Contract.GetMaxMagnitudes(&_AllocationManagerView.CallOpts, operators, strategy)
}

// GetMaxMagnitudes is a free data retrieval call binding the contract method 0x4a10ffe5.
//
// Solidity: function getMaxMagnitudes(address[] operators, address strategy) view returns(uint64[])
func (_AllocationManagerView *AllocationManagerViewCallerSession) GetMaxMagnitudes(operators []common.Address, strategy common.Address) ([]uint64, error) {
	return _AllocationManagerView.Contract.GetMaxMagnitudes(&_AllocationManagerView.CallOpts, operators, strategy)
}

// GetMaxMagnitudes0 is a free data retrieval call binding the contract method 0x547afb87.
//
// Solidity: function getMaxMagnitudes(address operator, address[] strategies) view returns(uint64[])
func (_AllocationManagerView *AllocationManagerViewCaller) GetMaxMagnitudes0(opts *bind.CallOpts, operator common.Address, strategies []common.Address) ([]uint64, error) {
	var out []interface{}
	err := _AllocationManagerView.contract.Call(opts, &out, "getMaxMagnitudes0", operator, strategies)

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

// GetMaxMagnitudes0 is a free data retrieval call binding the contract method 0x547afb87.
//
// Solidity: function getMaxMagnitudes(address operator, address[] strategies) view returns(uint64[])
func (_AllocationManagerView *AllocationManagerViewSession) GetMaxMagnitudes0(operator common.Address, strategies []common.Address) ([]uint64, error) {
	return _AllocationManagerView.Contract.GetMaxMagnitudes0(&_AllocationManagerView.CallOpts, operator, strategies)
}

// GetMaxMagnitudes0 is a free data retrieval call binding the contract method 0x547afb87.
//
// Solidity: function getMaxMagnitudes(address operator, address[] strategies) view returns(uint64[])
func (_AllocationManagerView *AllocationManagerViewCallerSession) GetMaxMagnitudes0(operator common.Address, strategies []common.Address) ([]uint64, error) {
	return _AllocationManagerView.Contract.GetMaxMagnitudes0(&_AllocationManagerView.CallOpts, operator, strategies)
}

// GetMaxMagnitudesAtBlock is a free data retrieval call binding the contract method 0x94d7d00c.
//
// Solidity: function getMaxMagnitudesAtBlock(address operator, address[] strategies, uint32 blockNumber) view returns(uint64[])
func (_AllocationManagerView *AllocationManagerViewCaller) GetMaxMagnitudesAtBlock(opts *bind.CallOpts, operator common.Address, strategies []common.Address, blockNumber uint32) ([]uint64, error) {
	var out []interface{}
	err := _AllocationManagerView.contract.Call(opts, &out, "getMaxMagnitudesAtBlock", operator, strategies, blockNumber)

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

// GetMaxMagnitudesAtBlock is a free data retrieval call binding the contract method 0x94d7d00c.
//
// Solidity: function getMaxMagnitudesAtBlock(address operator, address[] strategies, uint32 blockNumber) view returns(uint64[])
func (_AllocationManagerView *AllocationManagerViewSession) GetMaxMagnitudesAtBlock(operator common.Address, strategies []common.Address, blockNumber uint32) ([]uint64, error) {
	return _AllocationManagerView.Contract.GetMaxMagnitudesAtBlock(&_AllocationManagerView.CallOpts, operator, strategies, blockNumber)
}

// GetMaxMagnitudesAtBlock is a free data retrieval call binding the contract method 0x94d7d00c.
//
// Solidity: function getMaxMagnitudesAtBlock(address operator, address[] strategies, uint32 blockNumber) view returns(uint64[])
func (_AllocationManagerView *AllocationManagerViewCallerSession) GetMaxMagnitudesAtBlock(operator common.Address, strategies []common.Address, blockNumber uint32) ([]uint64, error) {
	return _AllocationManagerView.Contract.GetMaxMagnitudesAtBlock(&_AllocationManagerView.CallOpts, operator, strategies, blockNumber)
}

// GetMemberCount is a free data retrieval call binding the contract method 0xb2447af7.
//
// Solidity: function getMemberCount((address,uint32) operatorSet) view returns(uint256)
func (_AllocationManagerView *AllocationManagerViewCaller) GetMemberCount(opts *bind.CallOpts, operatorSet OperatorSet) (*big.Int, error) {
	var out []interface{}
	err := _AllocationManagerView.contract.Call(opts, &out, "getMemberCount", operatorSet)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMemberCount is a free data retrieval call binding the contract method 0xb2447af7.
//
// Solidity: function getMemberCount((address,uint32) operatorSet) view returns(uint256)
func (_AllocationManagerView *AllocationManagerViewSession) GetMemberCount(operatorSet OperatorSet) (*big.Int, error) {
	return _AllocationManagerView.Contract.GetMemberCount(&_AllocationManagerView.CallOpts, operatorSet)
}

// GetMemberCount is a free data retrieval call binding the contract method 0xb2447af7.
//
// Solidity: function getMemberCount((address,uint32) operatorSet) view returns(uint256)
func (_AllocationManagerView *AllocationManagerViewCallerSession) GetMemberCount(operatorSet OperatorSet) (*big.Int, error) {
	return _AllocationManagerView.Contract.GetMemberCount(&_AllocationManagerView.CallOpts, operatorSet)
}

// GetMembers is a free data retrieval call binding the contract method 0x6e875dba.
//
// Solidity: function getMembers((address,uint32) operatorSet) view returns(address[])
func (_AllocationManagerView *AllocationManagerViewCaller) GetMembers(opts *bind.CallOpts, operatorSet OperatorSet) ([]common.Address, error) {
	var out []interface{}
	err := _AllocationManagerView.contract.Call(opts, &out, "getMembers", operatorSet)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetMembers is a free data retrieval call binding the contract method 0x6e875dba.
//
// Solidity: function getMembers((address,uint32) operatorSet) view returns(address[])
func (_AllocationManagerView *AllocationManagerViewSession) GetMembers(operatorSet OperatorSet) ([]common.Address, error) {
	return _AllocationManagerView.Contract.GetMembers(&_AllocationManagerView.CallOpts, operatorSet)
}

// GetMembers is a free data retrieval call binding the contract method 0x6e875dba.
//
// Solidity: function getMembers((address,uint32) operatorSet) view returns(address[])
func (_AllocationManagerView *AllocationManagerViewCallerSession) GetMembers(operatorSet OperatorSet) ([]common.Address, error) {
	return _AllocationManagerView.Contract.GetMembers(&_AllocationManagerView.CallOpts, operatorSet)
}

// GetMinimumSlashableStake is a free data retrieval call binding the contract method 0x2bab2c4a.
//
// Solidity: function getMinimumSlashableStake((address,uint32) operatorSet, address[] operators, address[] strategies, uint32 futureBlock) view returns(uint256[][] slashableStake)
func (_AllocationManagerView *AllocationManagerViewCaller) GetMinimumSlashableStake(opts *bind.CallOpts, operatorSet OperatorSet, operators []common.Address, strategies []common.Address, futureBlock uint32) ([][]*big.Int, error) {
	var out []interface{}
	err := _AllocationManagerView.contract.Call(opts, &out, "getMinimumSlashableStake", operatorSet, operators, strategies, futureBlock)

	if err != nil {
		return *new([][]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([][]*big.Int)).(*[][]*big.Int)

	return out0, err

}

// GetMinimumSlashableStake is a free data retrieval call binding the contract method 0x2bab2c4a.
//
// Solidity: function getMinimumSlashableStake((address,uint32) operatorSet, address[] operators, address[] strategies, uint32 futureBlock) view returns(uint256[][] slashableStake)
func (_AllocationManagerView *AllocationManagerViewSession) GetMinimumSlashableStake(operatorSet OperatorSet, operators []common.Address, strategies []common.Address, futureBlock uint32) ([][]*big.Int, error) {
	return _AllocationManagerView.Contract.GetMinimumSlashableStake(&_AllocationManagerView.CallOpts, operatorSet, operators, strategies, futureBlock)
}

// GetMinimumSlashableStake is a free data retrieval call binding the contract method 0x2bab2c4a.
//
// Solidity: function getMinimumSlashableStake((address,uint32) operatorSet, address[] operators, address[] strategies, uint32 futureBlock) view returns(uint256[][] slashableStake)
func (_AllocationManagerView *AllocationManagerViewCallerSession) GetMinimumSlashableStake(operatorSet OperatorSet, operators []common.Address, strategies []common.Address, futureBlock uint32) ([][]*big.Int, error) {
	return _AllocationManagerView.Contract.GetMinimumSlashableStake(&_AllocationManagerView.CallOpts, operatorSet, operators, strategies, futureBlock)
}

// GetOperatorSetCount is a free data retrieval call binding the contract method 0xba1a84e5.
//
// Solidity: function getOperatorSetCount(address avs) view returns(uint256)
func (_AllocationManagerView *AllocationManagerViewCaller) GetOperatorSetCount(opts *bind.CallOpts, avs common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AllocationManagerView.contract.Call(opts, &out, "getOperatorSetCount", avs)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOperatorSetCount is a free data retrieval call binding the contract method 0xba1a84e5.
//
// Solidity: function getOperatorSetCount(address avs) view returns(uint256)
func (_AllocationManagerView *AllocationManagerViewSession) GetOperatorSetCount(avs common.Address) (*big.Int, error) {
	return _AllocationManagerView.Contract.GetOperatorSetCount(&_AllocationManagerView.CallOpts, avs)
}

// GetOperatorSetCount is a free data retrieval call binding the contract method 0xba1a84e5.
//
// Solidity: function getOperatorSetCount(address avs) view returns(uint256)
func (_AllocationManagerView *AllocationManagerViewCallerSession) GetOperatorSetCount(avs common.Address) (*big.Int, error) {
	return _AllocationManagerView.Contract.GetOperatorSetCount(&_AllocationManagerView.CallOpts, avs)
}

// GetPendingSlasher is a free data retrieval call binding the contract method 0xd7794857.
//
// Solidity: function getPendingSlasher((address,uint32) operatorSet) view returns(address, uint32)
func (_AllocationManagerView *AllocationManagerViewCaller) GetPendingSlasher(opts *bind.CallOpts, operatorSet OperatorSet) (common.Address, uint32, error) {
	var out []interface{}
	err := _AllocationManagerView.contract.Call(opts, &out, "getPendingSlasher", operatorSet)

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
func (_AllocationManagerView *AllocationManagerViewSession) GetPendingSlasher(operatorSet OperatorSet) (common.Address, uint32, error) {
	return _AllocationManagerView.Contract.GetPendingSlasher(&_AllocationManagerView.CallOpts, operatorSet)
}

// GetPendingSlasher is a free data retrieval call binding the contract method 0xd7794857.
//
// Solidity: function getPendingSlasher((address,uint32) operatorSet) view returns(address, uint32)
func (_AllocationManagerView *AllocationManagerViewCallerSession) GetPendingSlasher(operatorSet OperatorSet) (common.Address, uint32, error) {
	return _AllocationManagerView.Contract.GetPendingSlasher(&_AllocationManagerView.CallOpts, operatorSet)
}

// GetRedistributionRecipient is a free data retrieval call binding the contract method 0x0f3df50e.
//
// Solidity: function getRedistributionRecipient((address,uint32) operatorSet) view returns(address)
func (_AllocationManagerView *AllocationManagerViewCaller) GetRedistributionRecipient(opts *bind.CallOpts, operatorSet OperatorSet) (common.Address, error) {
	var out []interface{}
	err := _AllocationManagerView.contract.Call(opts, &out, "getRedistributionRecipient", operatorSet)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRedistributionRecipient is a free data retrieval call binding the contract method 0x0f3df50e.
//
// Solidity: function getRedistributionRecipient((address,uint32) operatorSet) view returns(address)
func (_AllocationManagerView *AllocationManagerViewSession) GetRedistributionRecipient(operatorSet OperatorSet) (common.Address, error) {
	return _AllocationManagerView.Contract.GetRedistributionRecipient(&_AllocationManagerView.CallOpts, operatorSet)
}

// GetRedistributionRecipient is a free data retrieval call binding the contract method 0x0f3df50e.
//
// Solidity: function getRedistributionRecipient((address,uint32) operatorSet) view returns(address)
func (_AllocationManagerView *AllocationManagerViewCallerSession) GetRedistributionRecipient(operatorSet OperatorSet) (common.Address, error) {
	return _AllocationManagerView.Contract.GetRedistributionRecipient(&_AllocationManagerView.CallOpts, operatorSet)
}

// GetRegisteredSets is a free data retrieval call binding the contract method 0x79ae50cd.
//
// Solidity: function getRegisteredSets(address operator) view returns((address,uint32)[])
func (_AllocationManagerView *AllocationManagerViewCaller) GetRegisteredSets(opts *bind.CallOpts, operator common.Address) ([]OperatorSet, error) {
	var out []interface{}
	err := _AllocationManagerView.contract.Call(opts, &out, "getRegisteredSets", operator)

	if err != nil {
		return *new([]OperatorSet), err
	}

	out0 := *abi.ConvertType(out[0], new([]OperatorSet)).(*[]OperatorSet)

	return out0, err

}

// GetRegisteredSets is a free data retrieval call binding the contract method 0x79ae50cd.
//
// Solidity: function getRegisteredSets(address operator) view returns((address,uint32)[])
func (_AllocationManagerView *AllocationManagerViewSession) GetRegisteredSets(operator common.Address) ([]OperatorSet, error) {
	return _AllocationManagerView.Contract.GetRegisteredSets(&_AllocationManagerView.CallOpts, operator)
}

// GetRegisteredSets is a free data retrieval call binding the contract method 0x79ae50cd.
//
// Solidity: function getRegisteredSets(address operator) view returns((address,uint32)[])
func (_AllocationManagerView *AllocationManagerViewCallerSession) GetRegisteredSets(operator common.Address) ([]OperatorSet, error) {
	return _AllocationManagerView.Contract.GetRegisteredSets(&_AllocationManagerView.CallOpts, operator)
}

// GetSlashCount is a free data retrieval call binding the contract method 0x4cfd2939.
//
// Solidity: function getSlashCount((address,uint32) operatorSet) view returns(uint256)
func (_AllocationManagerView *AllocationManagerViewCaller) GetSlashCount(opts *bind.CallOpts, operatorSet OperatorSet) (*big.Int, error) {
	var out []interface{}
	err := _AllocationManagerView.contract.Call(opts, &out, "getSlashCount", operatorSet)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSlashCount is a free data retrieval call binding the contract method 0x4cfd2939.
//
// Solidity: function getSlashCount((address,uint32) operatorSet) view returns(uint256)
func (_AllocationManagerView *AllocationManagerViewSession) GetSlashCount(operatorSet OperatorSet) (*big.Int, error) {
	return _AllocationManagerView.Contract.GetSlashCount(&_AllocationManagerView.CallOpts, operatorSet)
}

// GetSlashCount is a free data retrieval call binding the contract method 0x4cfd2939.
//
// Solidity: function getSlashCount((address,uint32) operatorSet) view returns(uint256)
func (_AllocationManagerView *AllocationManagerViewCallerSession) GetSlashCount(operatorSet OperatorSet) (*big.Int, error) {
	return _AllocationManagerView.Contract.GetSlashCount(&_AllocationManagerView.CallOpts, operatorSet)
}

// GetSlasher is a free data retrieval call binding the contract method 0xd4a3fcce.
//
// Solidity: function getSlasher((address,uint32) operatorSet) view returns(address)
func (_AllocationManagerView *AllocationManagerViewCaller) GetSlasher(opts *bind.CallOpts, operatorSet OperatorSet) (common.Address, error) {
	var out []interface{}
	err := _AllocationManagerView.contract.Call(opts, &out, "getSlasher", operatorSet)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSlasher is a free data retrieval call binding the contract method 0xd4a3fcce.
//
// Solidity: function getSlasher((address,uint32) operatorSet) view returns(address)
func (_AllocationManagerView *AllocationManagerViewSession) GetSlasher(operatorSet OperatorSet) (common.Address, error) {
	return _AllocationManagerView.Contract.GetSlasher(&_AllocationManagerView.CallOpts, operatorSet)
}

// GetSlasher is a free data retrieval call binding the contract method 0xd4a3fcce.
//
// Solidity: function getSlasher((address,uint32) operatorSet) view returns(address)
func (_AllocationManagerView *AllocationManagerViewCallerSession) GetSlasher(operatorSet OperatorSet) (common.Address, error) {
	return _AllocationManagerView.Contract.GetSlasher(&_AllocationManagerView.CallOpts, operatorSet)
}

// GetStrategiesInOperatorSet is a free data retrieval call binding the contract method 0x4177a87c.
//
// Solidity: function getStrategiesInOperatorSet((address,uint32) operatorSet) view returns(address[])
func (_AllocationManagerView *AllocationManagerViewCaller) GetStrategiesInOperatorSet(opts *bind.CallOpts, operatorSet OperatorSet) ([]common.Address, error) {
	var out []interface{}
	err := _AllocationManagerView.contract.Call(opts, &out, "getStrategiesInOperatorSet", operatorSet)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetStrategiesInOperatorSet is a free data retrieval call binding the contract method 0x4177a87c.
//
// Solidity: function getStrategiesInOperatorSet((address,uint32) operatorSet) view returns(address[])
func (_AllocationManagerView *AllocationManagerViewSession) GetStrategiesInOperatorSet(operatorSet OperatorSet) ([]common.Address, error) {
	return _AllocationManagerView.Contract.GetStrategiesInOperatorSet(&_AllocationManagerView.CallOpts, operatorSet)
}

// GetStrategiesInOperatorSet is a free data retrieval call binding the contract method 0x4177a87c.
//
// Solidity: function getStrategiesInOperatorSet((address,uint32) operatorSet) view returns(address[])
func (_AllocationManagerView *AllocationManagerViewCallerSession) GetStrategiesInOperatorSet(operatorSet OperatorSet) ([]common.Address, error) {
	return _AllocationManagerView.Contract.GetStrategiesInOperatorSet(&_AllocationManagerView.CallOpts, operatorSet)
}

// GetStrategyAllocations is a free data retrieval call binding the contract method 0x40120dab.
//
// Solidity: function getStrategyAllocations(address operator, address strategy) view returns((address,uint32)[], (uint64,int128,uint32)[])
func (_AllocationManagerView *AllocationManagerViewCaller) GetStrategyAllocations(opts *bind.CallOpts, operator common.Address, strategy common.Address) ([]OperatorSet, []IAllocationManagerTypesAllocation, error) {
	var out []interface{}
	err := _AllocationManagerView.contract.Call(opts, &out, "getStrategyAllocations", operator, strategy)

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
func (_AllocationManagerView *AllocationManagerViewSession) GetStrategyAllocations(operator common.Address, strategy common.Address) ([]OperatorSet, []IAllocationManagerTypesAllocation, error) {
	return _AllocationManagerView.Contract.GetStrategyAllocations(&_AllocationManagerView.CallOpts, operator, strategy)
}

// GetStrategyAllocations is a free data retrieval call binding the contract method 0x40120dab.
//
// Solidity: function getStrategyAllocations(address operator, address strategy) view returns((address,uint32)[], (uint64,int128,uint32)[])
func (_AllocationManagerView *AllocationManagerViewCallerSession) GetStrategyAllocations(operator common.Address, strategy common.Address) ([]OperatorSet, []IAllocationManagerTypesAllocation, error) {
	return _AllocationManagerView.Contract.GetStrategyAllocations(&_AllocationManagerView.CallOpts, operator, strategy)
}

// IsMemberOfOperatorSet is a free data retrieval call binding the contract method 0x670d3ba2.
//
// Solidity: function isMemberOfOperatorSet(address operator, (address,uint32) operatorSet) view returns(bool)
func (_AllocationManagerView *AllocationManagerViewCaller) IsMemberOfOperatorSet(opts *bind.CallOpts, operator common.Address, operatorSet OperatorSet) (bool, error) {
	var out []interface{}
	err := _AllocationManagerView.contract.Call(opts, &out, "isMemberOfOperatorSet", operator, operatorSet)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMemberOfOperatorSet is a free data retrieval call binding the contract method 0x670d3ba2.
//
// Solidity: function isMemberOfOperatorSet(address operator, (address,uint32) operatorSet) view returns(bool)
func (_AllocationManagerView *AllocationManagerViewSession) IsMemberOfOperatorSet(operator common.Address, operatorSet OperatorSet) (bool, error) {
	return _AllocationManagerView.Contract.IsMemberOfOperatorSet(&_AllocationManagerView.CallOpts, operator, operatorSet)
}

// IsMemberOfOperatorSet is a free data retrieval call binding the contract method 0x670d3ba2.
//
// Solidity: function isMemberOfOperatorSet(address operator, (address,uint32) operatorSet) view returns(bool)
func (_AllocationManagerView *AllocationManagerViewCallerSession) IsMemberOfOperatorSet(operator common.Address, operatorSet OperatorSet) (bool, error) {
	return _AllocationManagerView.Contract.IsMemberOfOperatorSet(&_AllocationManagerView.CallOpts, operator, operatorSet)
}

// IsOperatorRedistributable is a free data retrieval call binding the contract method 0xdc2af692.
//
// Solidity: function isOperatorRedistributable(address operator) view returns(bool)
func (_AllocationManagerView *AllocationManagerViewCaller) IsOperatorRedistributable(opts *bind.CallOpts, operator common.Address) (bool, error) {
	var out []interface{}
	err := _AllocationManagerView.contract.Call(opts, &out, "isOperatorRedistributable", operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperatorRedistributable is a free data retrieval call binding the contract method 0xdc2af692.
//
// Solidity: function isOperatorRedistributable(address operator) view returns(bool)
func (_AllocationManagerView *AllocationManagerViewSession) IsOperatorRedistributable(operator common.Address) (bool, error) {
	return _AllocationManagerView.Contract.IsOperatorRedistributable(&_AllocationManagerView.CallOpts, operator)
}

// IsOperatorRedistributable is a free data retrieval call binding the contract method 0xdc2af692.
//
// Solidity: function isOperatorRedistributable(address operator) view returns(bool)
func (_AllocationManagerView *AllocationManagerViewCallerSession) IsOperatorRedistributable(operator common.Address) (bool, error) {
	return _AllocationManagerView.Contract.IsOperatorRedistributable(&_AllocationManagerView.CallOpts, operator)
}

// IsOperatorSet is a free data retrieval call binding the contract method 0x260dc758.
//
// Solidity: function isOperatorSet((address,uint32) operatorSet) view returns(bool)
func (_AllocationManagerView *AllocationManagerViewCaller) IsOperatorSet(opts *bind.CallOpts, operatorSet OperatorSet) (bool, error) {
	var out []interface{}
	err := _AllocationManagerView.contract.Call(opts, &out, "isOperatorSet", operatorSet)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperatorSet is a free data retrieval call binding the contract method 0x260dc758.
//
// Solidity: function isOperatorSet((address,uint32) operatorSet) view returns(bool)
func (_AllocationManagerView *AllocationManagerViewSession) IsOperatorSet(operatorSet OperatorSet) (bool, error) {
	return _AllocationManagerView.Contract.IsOperatorSet(&_AllocationManagerView.CallOpts, operatorSet)
}

// IsOperatorSet is a free data retrieval call binding the contract method 0x260dc758.
//
// Solidity: function isOperatorSet((address,uint32) operatorSet) view returns(bool)
func (_AllocationManagerView *AllocationManagerViewCallerSession) IsOperatorSet(operatorSet OperatorSet) (bool, error) {
	return _AllocationManagerView.Contract.IsOperatorSet(&_AllocationManagerView.CallOpts, operatorSet)
}

// IsOperatorSlashable is a free data retrieval call binding the contract method 0x1352c3e6.
//
// Solidity: function isOperatorSlashable(address operator, (address,uint32) operatorSet) view returns(bool)
func (_AllocationManagerView *AllocationManagerViewCaller) IsOperatorSlashable(opts *bind.CallOpts, operator common.Address, operatorSet OperatorSet) (bool, error) {
	var out []interface{}
	err := _AllocationManagerView.contract.Call(opts, &out, "isOperatorSlashable", operator, operatorSet)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperatorSlashable is a free data retrieval call binding the contract method 0x1352c3e6.
//
// Solidity: function isOperatorSlashable(address operator, (address,uint32) operatorSet) view returns(bool)
func (_AllocationManagerView *AllocationManagerViewSession) IsOperatorSlashable(operator common.Address, operatorSet OperatorSet) (bool, error) {
	return _AllocationManagerView.Contract.IsOperatorSlashable(&_AllocationManagerView.CallOpts, operator, operatorSet)
}

// IsOperatorSlashable is a free data retrieval call binding the contract method 0x1352c3e6.
//
// Solidity: function isOperatorSlashable(address operator, (address,uint32) operatorSet) view returns(bool)
func (_AllocationManagerView *AllocationManagerViewCallerSession) IsOperatorSlashable(operator common.Address, operatorSet OperatorSet) (bool, error) {
	return _AllocationManagerView.Contract.IsOperatorSlashable(&_AllocationManagerView.CallOpts, operator, operatorSet)
}

// IsRedistributingOperatorSet is a free data retrieval call binding the contract method 0xf231bd08.
//
// Solidity: function isRedistributingOperatorSet((address,uint32) operatorSet) view returns(bool)
func (_AllocationManagerView *AllocationManagerViewCaller) IsRedistributingOperatorSet(opts *bind.CallOpts, operatorSet OperatorSet) (bool, error) {
	var out []interface{}
	err := _AllocationManagerView.contract.Call(opts, &out, "isRedistributingOperatorSet", operatorSet)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRedistributingOperatorSet is a free data retrieval call binding the contract method 0xf231bd08.
//
// Solidity: function isRedistributingOperatorSet((address,uint32) operatorSet) view returns(bool)
func (_AllocationManagerView *AllocationManagerViewSession) IsRedistributingOperatorSet(operatorSet OperatorSet) (bool, error) {
	return _AllocationManagerView.Contract.IsRedistributingOperatorSet(&_AllocationManagerView.CallOpts, operatorSet)
}

// IsRedistributingOperatorSet is a free data retrieval call binding the contract method 0xf231bd08.
//
// Solidity: function isRedistributingOperatorSet((address,uint32) operatorSet) view returns(bool)
func (_AllocationManagerView *AllocationManagerViewCallerSession) IsRedistributingOperatorSet(operatorSet OperatorSet) (bool, error) {
	return _AllocationManagerView.Contract.IsRedistributingOperatorSet(&_AllocationManagerView.CallOpts, operatorSet)
}

// AllocationManagerViewAVSMetadataURIUpdatedIterator is returned from FilterAVSMetadataURIUpdated and is used to iterate over the raw logs and unpacked data for AVSMetadataURIUpdated events raised by the AllocationManagerView contract.
type AllocationManagerViewAVSMetadataURIUpdatedIterator struct {
	Event *AllocationManagerViewAVSMetadataURIUpdated // Event containing the contract specifics and raw log

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
func (it *AllocationManagerViewAVSMetadataURIUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AllocationManagerViewAVSMetadataURIUpdated)
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
		it.Event = new(AllocationManagerViewAVSMetadataURIUpdated)
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
func (it *AllocationManagerViewAVSMetadataURIUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AllocationManagerViewAVSMetadataURIUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AllocationManagerViewAVSMetadataURIUpdated represents a AVSMetadataURIUpdated event raised by the AllocationManagerView contract.
type AllocationManagerViewAVSMetadataURIUpdated struct {
	Avs         common.Address
	MetadataURI string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAVSMetadataURIUpdated is a free log retrieval operation binding the contract event 0xa89c1dc243d8908a96dd84944bcc97d6bc6ac00dd78e20621576be6a3c943713.
//
// Solidity: event AVSMetadataURIUpdated(address indexed avs, string metadataURI)
func (_AllocationManagerView *AllocationManagerViewFilterer) FilterAVSMetadataURIUpdated(opts *bind.FilterOpts, avs []common.Address) (*AllocationManagerViewAVSMetadataURIUpdatedIterator, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _AllocationManagerView.contract.FilterLogs(opts, "AVSMetadataURIUpdated", avsRule)
	if err != nil {
		return nil, err
	}
	return &AllocationManagerViewAVSMetadataURIUpdatedIterator{contract: _AllocationManagerView.contract, event: "AVSMetadataURIUpdated", logs: logs, sub: sub}, nil
}

// WatchAVSMetadataURIUpdated is a free log subscription operation binding the contract event 0xa89c1dc243d8908a96dd84944bcc97d6bc6ac00dd78e20621576be6a3c943713.
//
// Solidity: event AVSMetadataURIUpdated(address indexed avs, string metadataURI)
func (_AllocationManagerView *AllocationManagerViewFilterer) WatchAVSMetadataURIUpdated(opts *bind.WatchOpts, sink chan<- *AllocationManagerViewAVSMetadataURIUpdated, avs []common.Address) (event.Subscription, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _AllocationManagerView.contract.WatchLogs(opts, "AVSMetadataURIUpdated", avsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AllocationManagerViewAVSMetadataURIUpdated)
				if err := _AllocationManagerView.contract.UnpackLog(event, "AVSMetadataURIUpdated", log); err != nil {
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
func (_AllocationManagerView *AllocationManagerViewFilterer) ParseAVSMetadataURIUpdated(log types.Log) (*AllocationManagerViewAVSMetadataURIUpdated, error) {
	event := new(AllocationManagerViewAVSMetadataURIUpdated)
	if err := _AllocationManagerView.contract.UnpackLog(event, "AVSMetadataURIUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AllocationManagerViewAVSRegistrarSetIterator is returned from FilterAVSRegistrarSet and is used to iterate over the raw logs and unpacked data for AVSRegistrarSet events raised by the AllocationManagerView contract.
type AllocationManagerViewAVSRegistrarSetIterator struct {
	Event *AllocationManagerViewAVSRegistrarSet // Event containing the contract specifics and raw log

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
func (it *AllocationManagerViewAVSRegistrarSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AllocationManagerViewAVSRegistrarSet)
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
		it.Event = new(AllocationManagerViewAVSRegistrarSet)
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
func (it *AllocationManagerViewAVSRegistrarSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AllocationManagerViewAVSRegistrarSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AllocationManagerViewAVSRegistrarSet represents a AVSRegistrarSet event raised by the AllocationManagerView contract.
type AllocationManagerViewAVSRegistrarSet struct {
	Avs       common.Address
	Registrar common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAVSRegistrarSet is a free log retrieval operation binding the contract event 0x2ae945c40c44dc0ec263f95609c3fdc6952e0aefa22d6374e44f2c997acedf85.
//
// Solidity: event AVSRegistrarSet(address avs, address registrar)
func (_AllocationManagerView *AllocationManagerViewFilterer) FilterAVSRegistrarSet(opts *bind.FilterOpts) (*AllocationManagerViewAVSRegistrarSetIterator, error) {

	logs, sub, err := _AllocationManagerView.contract.FilterLogs(opts, "AVSRegistrarSet")
	if err != nil {
		return nil, err
	}
	return &AllocationManagerViewAVSRegistrarSetIterator{contract: _AllocationManagerView.contract, event: "AVSRegistrarSet", logs: logs, sub: sub}, nil
}

// WatchAVSRegistrarSet is a free log subscription operation binding the contract event 0x2ae945c40c44dc0ec263f95609c3fdc6952e0aefa22d6374e44f2c997acedf85.
//
// Solidity: event AVSRegistrarSet(address avs, address registrar)
func (_AllocationManagerView *AllocationManagerViewFilterer) WatchAVSRegistrarSet(opts *bind.WatchOpts, sink chan<- *AllocationManagerViewAVSRegistrarSet) (event.Subscription, error) {

	logs, sub, err := _AllocationManagerView.contract.WatchLogs(opts, "AVSRegistrarSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AllocationManagerViewAVSRegistrarSet)
				if err := _AllocationManagerView.contract.UnpackLog(event, "AVSRegistrarSet", log); err != nil {
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
func (_AllocationManagerView *AllocationManagerViewFilterer) ParseAVSRegistrarSet(log types.Log) (*AllocationManagerViewAVSRegistrarSet, error) {
	event := new(AllocationManagerViewAVSRegistrarSet)
	if err := _AllocationManagerView.contract.UnpackLog(event, "AVSRegistrarSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AllocationManagerViewAllocationDelaySetIterator is returned from FilterAllocationDelaySet and is used to iterate over the raw logs and unpacked data for AllocationDelaySet events raised by the AllocationManagerView contract.
type AllocationManagerViewAllocationDelaySetIterator struct {
	Event *AllocationManagerViewAllocationDelaySet // Event containing the contract specifics and raw log

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
func (it *AllocationManagerViewAllocationDelaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AllocationManagerViewAllocationDelaySet)
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
		it.Event = new(AllocationManagerViewAllocationDelaySet)
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
func (it *AllocationManagerViewAllocationDelaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AllocationManagerViewAllocationDelaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AllocationManagerViewAllocationDelaySet represents a AllocationDelaySet event raised by the AllocationManagerView contract.
type AllocationManagerViewAllocationDelaySet struct {
	Operator    common.Address
	Delay       uint32
	EffectBlock uint32
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAllocationDelaySet is a free log retrieval operation binding the contract event 0x4e85751d6331506c6c62335f207eb31f12a61e570f34f5c17640308785c6d4db.
//
// Solidity: event AllocationDelaySet(address operator, uint32 delay, uint32 effectBlock)
func (_AllocationManagerView *AllocationManagerViewFilterer) FilterAllocationDelaySet(opts *bind.FilterOpts) (*AllocationManagerViewAllocationDelaySetIterator, error) {

	logs, sub, err := _AllocationManagerView.contract.FilterLogs(opts, "AllocationDelaySet")
	if err != nil {
		return nil, err
	}
	return &AllocationManagerViewAllocationDelaySetIterator{contract: _AllocationManagerView.contract, event: "AllocationDelaySet", logs: logs, sub: sub}, nil
}

// WatchAllocationDelaySet is a free log subscription operation binding the contract event 0x4e85751d6331506c6c62335f207eb31f12a61e570f34f5c17640308785c6d4db.
//
// Solidity: event AllocationDelaySet(address operator, uint32 delay, uint32 effectBlock)
func (_AllocationManagerView *AllocationManagerViewFilterer) WatchAllocationDelaySet(opts *bind.WatchOpts, sink chan<- *AllocationManagerViewAllocationDelaySet) (event.Subscription, error) {

	logs, sub, err := _AllocationManagerView.contract.WatchLogs(opts, "AllocationDelaySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AllocationManagerViewAllocationDelaySet)
				if err := _AllocationManagerView.contract.UnpackLog(event, "AllocationDelaySet", log); err != nil {
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
func (_AllocationManagerView *AllocationManagerViewFilterer) ParseAllocationDelaySet(log types.Log) (*AllocationManagerViewAllocationDelaySet, error) {
	event := new(AllocationManagerViewAllocationDelaySet)
	if err := _AllocationManagerView.contract.UnpackLog(event, "AllocationDelaySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AllocationManagerViewAllocationUpdatedIterator is returned from FilterAllocationUpdated and is used to iterate over the raw logs and unpacked data for AllocationUpdated events raised by the AllocationManagerView contract.
type AllocationManagerViewAllocationUpdatedIterator struct {
	Event *AllocationManagerViewAllocationUpdated // Event containing the contract specifics and raw log

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
func (it *AllocationManagerViewAllocationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AllocationManagerViewAllocationUpdated)
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
		it.Event = new(AllocationManagerViewAllocationUpdated)
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
func (it *AllocationManagerViewAllocationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AllocationManagerViewAllocationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AllocationManagerViewAllocationUpdated represents a AllocationUpdated event raised by the AllocationManagerView contract.
type AllocationManagerViewAllocationUpdated struct {
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
func (_AllocationManagerView *AllocationManagerViewFilterer) FilterAllocationUpdated(opts *bind.FilterOpts) (*AllocationManagerViewAllocationUpdatedIterator, error) {

	logs, sub, err := _AllocationManagerView.contract.FilterLogs(opts, "AllocationUpdated")
	if err != nil {
		return nil, err
	}
	return &AllocationManagerViewAllocationUpdatedIterator{contract: _AllocationManagerView.contract, event: "AllocationUpdated", logs: logs, sub: sub}, nil
}

// WatchAllocationUpdated is a free log subscription operation binding the contract event 0x1487af5418c47ee5ea45ef4a93398668120890774a9e13487e61e9dc3baf76dd.
//
// Solidity: event AllocationUpdated(address operator, (address,uint32) operatorSet, address strategy, uint64 magnitude, uint32 effectBlock)
func (_AllocationManagerView *AllocationManagerViewFilterer) WatchAllocationUpdated(opts *bind.WatchOpts, sink chan<- *AllocationManagerViewAllocationUpdated) (event.Subscription, error) {

	logs, sub, err := _AllocationManagerView.contract.WatchLogs(opts, "AllocationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AllocationManagerViewAllocationUpdated)
				if err := _AllocationManagerView.contract.UnpackLog(event, "AllocationUpdated", log); err != nil {
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
func (_AllocationManagerView *AllocationManagerViewFilterer) ParseAllocationUpdated(log types.Log) (*AllocationManagerViewAllocationUpdated, error) {
	event := new(AllocationManagerViewAllocationUpdated)
	if err := _AllocationManagerView.contract.UnpackLog(event, "AllocationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AllocationManagerViewEncumberedMagnitudeUpdatedIterator is returned from FilterEncumberedMagnitudeUpdated and is used to iterate over the raw logs and unpacked data for EncumberedMagnitudeUpdated events raised by the AllocationManagerView contract.
type AllocationManagerViewEncumberedMagnitudeUpdatedIterator struct {
	Event *AllocationManagerViewEncumberedMagnitudeUpdated // Event containing the contract specifics and raw log

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
func (it *AllocationManagerViewEncumberedMagnitudeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AllocationManagerViewEncumberedMagnitudeUpdated)
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
		it.Event = new(AllocationManagerViewEncumberedMagnitudeUpdated)
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
func (it *AllocationManagerViewEncumberedMagnitudeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AllocationManagerViewEncumberedMagnitudeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AllocationManagerViewEncumberedMagnitudeUpdated represents a EncumberedMagnitudeUpdated event raised by the AllocationManagerView contract.
type AllocationManagerViewEncumberedMagnitudeUpdated struct {
	Operator            common.Address
	Strategy            common.Address
	EncumberedMagnitude uint64
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterEncumberedMagnitudeUpdated is a free log retrieval operation binding the contract event 0xacf9095feb3a370c9cf692421c69ef320d4db5c66e6a7d29c7694eb02364fc55.
//
// Solidity: event EncumberedMagnitudeUpdated(address operator, address strategy, uint64 encumberedMagnitude)
func (_AllocationManagerView *AllocationManagerViewFilterer) FilterEncumberedMagnitudeUpdated(opts *bind.FilterOpts) (*AllocationManagerViewEncumberedMagnitudeUpdatedIterator, error) {

	logs, sub, err := _AllocationManagerView.contract.FilterLogs(opts, "EncumberedMagnitudeUpdated")
	if err != nil {
		return nil, err
	}
	return &AllocationManagerViewEncumberedMagnitudeUpdatedIterator{contract: _AllocationManagerView.contract, event: "EncumberedMagnitudeUpdated", logs: logs, sub: sub}, nil
}

// WatchEncumberedMagnitudeUpdated is a free log subscription operation binding the contract event 0xacf9095feb3a370c9cf692421c69ef320d4db5c66e6a7d29c7694eb02364fc55.
//
// Solidity: event EncumberedMagnitudeUpdated(address operator, address strategy, uint64 encumberedMagnitude)
func (_AllocationManagerView *AllocationManagerViewFilterer) WatchEncumberedMagnitudeUpdated(opts *bind.WatchOpts, sink chan<- *AllocationManagerViewEncumberedMagnitudeUpdated) (event.Subscription, error) {

	logs, sub, err := _AllocationManagerView.contract.WatchLogs(opts, "EncumberedMagnitudeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AllocationManagerViewEncumberedMagnitudeUpdated)
				if err := _AllocationManagerView.contract.UnpackLog(event, "EncumberedMagnitudeUpdated", log); err != nil {
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
func (_AllocationManagerView *AllocationManagerViewFilterer) ParseEncumberedMagnitudeUpdated(log types.Log) (*AllocationManagerViewEncumberedMagnitudeUpdated, error) {
	event := new(AllocationManagerViewEncumberedMagnitudeUpdated)
	if err := _AllocationManagerView.contract.UnpackLog(event, "EncumberedMagnitudeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AllocationManagerViewMaxMagnitudeUpdatedIterator is returned from FilterMaxMagnitudeUpdated and is used to iterate over the raw logs and unpacked data for MaxMagnitudeUpdated events raised by the AllocationManagerView contract.
type AllocationManagerViewMaxMagnitudeUpdatedIterator struct {
	Event *AllocationManagerViewMaxMagnitudeUpdated // Event containing the contract specifics and raw log

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
func (it *AllocationManagerViewMaxMagnitudeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AllocationManagerViewMaxMagnitudeUpdated)
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
		it.Event = new(AllocationManagerViewMaxMagnitudeUpdated)
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
func (it *AllocationManagerViewMaxMagnitudeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AllocationManagerViewMaxMagnitudeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AllocationManagerViewMaxMagnitudeUpdated represents a MaxMagnitudeUpdated event raised by the AllocationManagerView contract.
type AllocationManagerViewMaxMagnitudeUpdated struct {
	Operator     common.Address
	Strategy     common.Address
	MaxMagnitude uint64
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMaxMagnitudeUpdated is a free log retrieval operation binding the contract event 0x1c6458079a41077d003c11faf9bf097e693bd67979e4e6500bac7b29db779b5c.
//
// Solidity: event MaxMagnitudeUpdated(address operator, address strategy, uint64 maxMagnitude)
func (_AllocationManagerView *AllocationManagerViewFilterer) FilterMaxMagnitudeUpdated(opts *bind.FilterOpts) (*AllocationManagerViewMaxMagnitudeUpdatedIterator, error) {

	logs, sub, err := _AllocationManagerView.contract.FilterLogs(opts, "MaxMagnitudeUpdated")
	if err != nil {
		return nil, err
	}
	return &AllocationManagerViewMaxMagnitudeUpdatedIterator{contract: _AllocationManagerView.contract, event: "MaxMagnitudeUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxMagnitudeUpdated is a free log subscription operation binding the contract event 0x1c6458079a41077d003c11faf9bf097e693bd67979e4e6500bac7b29db779b5c.
//
// Solidity: event MaxMagnitudeUpdated(address operator, address strategy, uint64 maxMagnitude)
func (_AllocationManagerView *AllocationManagerViewFilterer) WatchMaxMagnitudeUpdated(opts *bind.WatchOpts, sink chan<- *AllocationManagerViewMaxMagnitudeUpdated) (event.Subscription, error) {

	logs, sub, err := _AllocationManagerView.contract.WatchLogs(opts, "MaxMagnitudeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AllocationManagerViewMaxMagnitudeUpdated)
				if err := _AllocationManagerView.contract.UnpackLog(event, "MaxMagnitudeUpdated", log); err != nil {
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
func (_AllocationManagerView *AllocationManagerViewFilterer) ParseMaxMagnitudeUpdated(log types.Log) (*AllocationManagerViewMaxMagnitudeUpdated, error) {
	event := new(AllocationManagerViewMaxMagnitudeUpdated)
	if err := _AllocationManagerView.contract.UnpackLog(event, "MaxMagnitudeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AllocationManagerViewOperatorAddedToOperatorSetIterator is returned from FilterOperatorAddedToOperatorSet and is used to iterate over the raw logs and unpacked data for OperatorAddedToOperatorSet events raised by the AllocationManagerView contract.
type AllocationManagerViewOperatorAddedToOperatorSetIterator struct {
	Event *AllocationManagerViewOperatorAddedToOperatorSet // Event containing the contract specifics and raw log

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
func (it *AllocationManagerViewOperatorAddedToOperatorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AllocationManagerViewOperatorAddedToOperatorSet)
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
		it.Event = new(AllocationManagerViewOperatorAddedToOperatorSet)
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
func (it *AllocationManagerViewOperatorAddedToOperatorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AllocationManagerViewOperatorAddedToOperatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AllocationManagerViewOperatorAddedToOperatorSet represents a OperatorAddedToOperatorSet event raised by the AllocationManagerView contract.
type AllocationManagerViewOperatorAddedToOperatorSet struct {
	Operator    common.Address
	OperatorSet OperatorSet
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorAddedToOperatorSet is a free log retrieval operation binding the contract event 0x43232edf9071753d2321e5fa7e018363ee248e5f2142e6c08edd3265bfb4895e.
//
// Solidity: event OperatorAddedToOperatorSet(address indexed operator, (address,uint32) operatorSet)
func (_AllocationManagerView *AllocationManagerViewFilterer) FilterOperatorAddedToOperatorSet(opts *bind.FilterOpts, operator []common.Address) (*AllocationManagerViewOperatorAddedToOperatorSetIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _AllocationManagerView.contract.FilterLogs(opts, "OperatorAddedToOperatorSet", operatorRule)
	if err != nil {
		return nil, err
	}
	return &AllocationManagerViewOperatorAddedToOperatorSetIterator{contract: _AllocationManagerView.contract, event: "OperatorAddedToOperatorSet", logs: logs, sub: sub}, nil
}

// WatchOperatorAddedToOperatorSet is a free log subscription operation binding the contract event 0x43232edf9071753d2321e5fa7e018363ee248e5f2142e6c08edd3265bfb4895e.
//
// Solidity: event OperatorAddedToOperatorSet(address indexed operator, (address,uint32) operatorSet)
func (_AllocationManagerView *AllocationManagerViewFilterer) WatchOperatorAddedToOperatorSet(opts *bind.WatchOpts, sink chan<- *AllocationManagerViewOperatorAddedToOperatorSet, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _AllocationManagerView.contract.WatchLogs(opts, "OperatorAddedToOperatorSet", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AllocationManagerViewOperatorAddedToOperatorSet)
				if err := _AllocationManagerView.contract.UnpackLog(event, "OperatorAddedToOperatorSet", log); err != nil {
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
func (_AllocationManagerView *AllocationManagerViewFilterer) ParseOperatorAddedToOperatorSet(log types.Log) (*AllocationManagerViewOperatorAddedToOperatorSet, error) {
	event := new(AllocationManagerViewOperatorAddedToOperatorSet)
	if err := _AllocationManagerView.contract.UnpackLog(event, "OperatorAddedToOperatorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AllocationManagerViewOperatorRemovedFromOperatorSetIterator is returned from FilterOperatorRemovedFromOperatorSet and is used to iterate over the raw logs and unpacked data for OperatorRemovedFromOperatorSet events raised by the AllocationManagerView contract.
type AllocationManagerViewOperatorRemovedFromOperatorSetIterator struct {
	Event *AllocationManagerViewOperatorRemovedFromOperatorSet // Event containing the contract specifics and raw log

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
func (it *AllocationManagerViewOperatorRemovedFromOperatorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AllocationManagerViewOperatorRemovedFromOperatorSet)
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
		it.Event = new(AllocationManagerViewOperatorRemovedFromOperatorSet)
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
func (it *AllocationManagerViewOperatorRemovedFromOperatorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AllocationManagerViewOperatorRemovedFromOperatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AllocationManagerViewOperatorRemovedFromOperatorSet represents a OperatorRemovedFromOperatorSet event raised by the AllocationManagerView contract.
type AllocationManagerViewOperatorRemovedFromOperatorSet struct {
	Operator    common.Address
	OperatorSet OperatorSet
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorRemovedFromOperatorSet is a free log retrieval operation binding the contract event 0xad34c3070be1dffbcaa499d000ba2b8d9848aefcac3059df245dd95c4ece14fe.
//
// Solidity: event OperatorRemovedFromOperatorSet(address indexed operator, (address,uint32) operatorSet)
func (_AllocationManagerView *AllocationManagerViewFilterer) FilterOperatorRemovedFromOperatorSet(opts *bind.FilterOpts, operator []common.Address) (*AllocationManagerViewOperatorRemovedFromOperatorSetIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _AllocationManagerView.contract.FilterLogs(opts, "OperatorRemovedFromOperatorSet", operatorRule)
	if err != nil {
		return nil, err
	}
	return &AllocationManagerViewOperatorRemovedFromOperatorSetIterator{contract: _AllocationManagerView.contract, event: "OperatorRemovedFromOperatorSet", logs: logs, sub: sub}, nil
}

// WatchOperatorRemovedFromOperatorSet is a free log subscription operation binding the contract event 0xad34c3070be1dffbcaa499d000ba2b8d9848aefcac3059df245dd95c4ece14fe.
//
// Solidity: event OperatorRemovedFromOperatorSet(address indexed operator, (address,uint32) operatorSet)
func (_AllocationManagerView *AllocationManagerViewFilterer) WatchOperatorRemovedFromOperatorSet(opts *bind.WatchOpts, sink chan<- *AllocationManagerViewOperatorRemovedFromOperatorSet, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _AllocationManagerView.contract.WatchLogs(opts, "OperatorRemovedFromOperatorSet", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AllocationManagerViewOperatorRemovedFromOperatorSet)
				if err := _AllocationManagerView.contract.UnpackLog(event, "OperatorRemovedFromOperatorSet", log); err != nil {
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
func (_AllocationManagerView *AllocationManagerViewFilterer) ParseOperatorRemovedFromOperatorSet(log types.Log) (*AllocationManagerViewOperatorRemovedFromOperatorSet, error) {
	event := new(AllocationManagerViewOperatorRemovedFromOperatorSet)
	if err := _AllocationManagerView.contract.UnpackLog(event, "OperatorRemovedFromOperatorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AllocationManagerViewOperatorSetCreatedIterator is returned from FilterOperatorSetCreated and is used to iterate over the raw logs and unpacked data for OperatorSetCreated events raised by the AllocationManagerView contract.
type AllocationManagerViewOperatorSetCreatedIterator struct {
	Event *AllocationManagerViewOperatorSetCreated // Event containing the contract specifics and raw log

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
func (it *AllocationManagerViewOperatorSetCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AllocationManagerViewOperatorSetCreated)
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
		it.Event = new(AllocationManagerViewOperatorSetCreated)
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
func (it *AllocationManagerViewOperatorSetCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AllocationManagerViewOperatorSetCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AllocationManagerViewOperatorSetCreated represents a OperatorSetCreated event raised by the AllocationManagerView contract.
type AllocationManagerViewOperatorSetCreated struct {
	OperatorSet OperatorSet
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorSetCreated is a free log retrieval operation binding the contract event 0x31629285ead2335ae0933f86ed2ae63321f7af77b4e6eaabc42c057880977e6c.
//
// Solidity: event OperatorSetCreated((address,uint32) operatorSet)
func (_AllocationManagerView *AllocationManagerViewFilterer) FilterOperatorSetCreated(opts *bind.FilterOpts) (*AllocationManagerViewOperatorSetCreatedIterator, error) {

	logs, sub, err := _AllocationManagerView.contract.FilterLogs(opts, "OperatorSetCreated")
	if err != nil {
		return nil, err
	}
	return &AllocationManagerViewOperatorSetCreatedIterator{contract: _AllocationManagerView.contract, event: "OperatorSetCreated", logs: logs, sub: sub}, nil
}

// WatchOperatorSetCreated is a free log subscription operation binding the contract event 0x31629285ead2335ae0933f86ed2ae63321f7af77b4e6eaabc42c057880977e6c.
//
// Solidity: event OperatorSetCreated((address,uint32) operatorSet)
func (_AllocationManagerView *AllocationManagerViewFilterer) WatchOperatorSetCreated(opts *bind.WatchOpts, sink chan<- *AllocationManagerViewOperatorSetCreated) (event.Subscription, error) {

	logs, sub, err := _AllocationManagerView.contract.WatchLogs(opts, "OperatorSetCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AllocationManagerViewOperatorSetCreated)
				if err := _AllocationManagerView.contract.UnpackLog(event, "OperatorSetCreated", log); err != nil {
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
func (_AllocationManagerView *AllocationManagerViewFilterer) ParseOperatorSetCreated(log types.Log) (*AllocationManagerViewOperatorSetCreated, error) {
	event := new(AllocationManagerViewOperatorSetCreated)
	if err := _AllocationManagerView.contract.UnpackLog(event, "OperatorSetCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AllocationManagerViewOperatorSlashedIterator is returned from FilterOperatorSlashed and is used to iterate over the raw logs and unpacked data for OperatorSlashed events raised by the AllocationManagerView contract.
type AllocationManagerViewOperatorSlashedIterator struct {
	Event *AllocationManagerViewOperatorSlashed // Event containing the contract specifics and raw log

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
func (it *AllocationManagerViewOperatorSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AllocationManagerViewOperatorSlashed)
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
		it.Event = new(AllocationManagerViewOperatorSlashed)
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
func (it *AllocationManagerViewOperatorSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AllocationManagerViewOperatorSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AllocationManagerViewOperatorSlashed represents a OperatorSlashed event raised by the AllocationManagerView contract.
type AllocationManagerViewOperatorSlashed struct {
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
func (_AllocationManagerView *AllocationManagerViewFilterer) FilterOperatorSlashed(opts *bind.FilterOpts) (*AllocationManagerViewOperatorSlashedIterator, error) {

	logs, sub, err := _AllocationManagerView.contract.FilterLogs(opts, "OperatorSlashed")
	if err != nil {
		return nil, err
	}
	return &AllocationManagerViewOperatorSlashedIterator{contract: _AllocationManagerView.contract, event: "OperatorSlashed", logs: logs, sub: sub}, nil
}

// WatchOperatorSlashed is a free log subscription operation binding the contract event 0x80969ad29428d6797ee7aad084f9e4a42a82fc506dcd2ca3b6fb431f85ccebe5.
//
// Solidity: event OperatorSlashed(address operator, (address,uint32) operatorSet, address[] strategies, uint256[] wadSlashed, string description)
func (_AllocationManagerView *AllocationManagerViewFilterer) WatchOperatorSlashed(opts *bind.WatchOpts, sink chan<- *AllocationManagerViewOperatorSlashed) (event.Subscription, error) {

	logs, sub, err := _AllocationManagerView.contract.WatchLogs(opts, "OperatorSlashed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AllocationManagerViewOperatorSlashed)
				if err := _AllocationManagerView.contract.UnpackLog(event, "OperatorSlashed", log); err != nil {
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
func (_AllocationManagerView *AllocationManagerViewFilterer) ParseOperatorSlashed(log types.Log) (*AllocationManagerViewOperatorSlashed, error) {
	event := new(AllocationManagerViewOperatorSlashed)
	if err := _AllocationManagerView.contract.UnpackLog(event, "OperatorSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AllocationManagerViewRedistributionAddressSetIterator is returned from FilterRedistributionAddressSet and is used to iterate over the raw logs and unpacked data for RedistributionAddressSet events raised by the AllocationManagerView contract.
type AllocationManagerViewRedistributionAddressSetIterator struct {
	Event *AllocationManagerViewRedistributionAddressSet // Event containing the contract specifics and raw log

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
func (it *AllocationManagerViewRedistributionAddressSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AllocationManagerViewRedistributionAddressSet)
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
		it.Event = new(AllocationManagerViewRedistributionAddressSet)
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
func (it *AllocationManagerViewRedistributionAddressSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AllocationManagerViewRedistributionAddressSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AllocationManagerViewRedistributionAddressSet represents a RedistributionAddressSet event raised by the AllocationManagerView contract.
type AllocationManagerViewRedistributionAddressSet struct {
	OperatorSet             OperatorSet
	RedistributionRecipient common.Address
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterRedistributionAddressSet is a free log retrieval operation binding the contract event 0x90a6fa2a9b79b910872ebca540cf3bd8be827f586e6420c30d8836e30012907e.
//
// Solidity: event RedistributionAddressSet((address,uint32) operatorSet, address redistributionRecipient)
func (_AllocationManagerView *AllocationManagerViewFilterer) FilterRedistributionAddressSet(opts *bind.FilterOpts) (*AllocationManagerViewRedistributionAddressSetIterator, error) {

	logs, sub, err := _AllocationManagerView.contract.FilterLogs(opts, "RedistributionAddressSet")
	if err != nil {
		return nil, err
	}
	return &AllocationManagerViewRedistributionAddressSetIterator{contract: _AllocationManagerView.contract, event: "RedistributionAddressSet", logs: logs, sub: sub}, nil
}

// WatchRedistributionAddressSet is a free log subscription operation binding the contract event 0x90a6fa2a9b79b910872ebca540cf3bd8be827f586e6420c30d8836e30012907e.
//
// Solidity: event RedistributionAddressSet((address,uint32) operatorSet, address redistributionRecipient)
func (_AllocationManagerView *AllocationManagerViewFilterer) WatchRedistributionAddressSet(opts *bind.WatchOpts, sink chan<- *AllocationManagerViewRedistributionAddressSet) (event.Subscription, error) {

	logs, sub, err := _AllocationManagerView.contract.WatchLogs(opts, "RedistributionAddressSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AllocationManagerViewRedistributionAddressSet)
				if err := _AllocationManagerView.contract.UnpackLog(event, "RedistributionAddressSet", log); err != nil {
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
func (_AllocationManagerView *AllocationManagerViewFilterer) ParseRedistributionAddressSet(log types.Log) (*AllocationManagerViewRedistributionAddressSet, error) {
	event := new(AllocationManagerViewRedistributionAddressSet)
	if err := _AllocationManagerView.contract.UnpackLog(event, "RedistributionAddressSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AllocationManagerViewSlasherMigratedIterator is returned from FilterSlasherMigrated and is used to iterate over the raw logs and unpacked data for SlasherMigrated events raised by the AllocationManagerView contract.
type AllocationManagerViewSlasherMigratedIterator struct {
	Event *AllocationManagerViewSlasherMigrated // Event containing the contract specifics and raw log

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
func (it *AllocationManagerViewSlasherMigratedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AllocationManagerViewSlasherMigrated)
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
		it.Event = new(AllocationManagerViewSlasherMigrated)
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
func (it *AllocationManagerViewSlasherMigratedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AllocationManagerViewSlasherMigratedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AllocationManagerViewSlasherMigrated represents a SlasherMigrated event raised by the AllocationManagerView contract.
type AllocationManagerViewSlasherMigrated struct {
	OperatorSet OperatorSet
	Slasher     common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSlasherMigrated is a free log retrieval operation binding the contract event 0xf0c8fc7d71f647bd3a88ac369112517f6a4b8038e71913f2d20f71f877dfc725.
//
// Solidity: event SlasherMigrated((address,uint32) operatorSet, address slasher)
func (_AllocationManagerView *AllocationManagerViewFilterer) FilterSlasherMigrated(opts *bind.FilterOpts) (*AllocationManagerViewSlasherMigratedIterator, error) {

	logs, sub, err := _AllocationManagerView.contract.FilterLogs(opts, "SlasherMigrated")
	if err != nil {
		return nil, err
	}
	return &AllocationManagerViewSlasherMigratedIterator{contract: _AllocationManagerView.contract, event: "SlasherMigrated", logs: logs, sub: sub}, nil
}

// WatchSlasherMigrated is a free log subscription operation binding the contract event 0xf0c8fc7d71f647bd3a88ac369112517f6a4b8038e71913f2d20f71f877dfc725.
//
// Solidity: event SlasherMigrated((address,uint32) operatorSet, address slasher)
func (_AllocationManagerView *AllocationManagerViewFilterer) WatchSlasherMigrated(opts *bind.WatchOpts, sink chan<- *AllocationManagerViewSlasherMigrated) (event.Subscription, error) {

	logs, sub, err := _AllocationManagerView.contract.WatchLogs(opts, "SlasherMigrated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AllocationManagerViewSlasherMigrated)
				if err := _AllocationManagerView.contract.UnpackLog(event, "SlasherMigrated", log); err != nil {
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
func (_AllocationManagerView *AllocationManagerViewFilterer) ParseSlasherMigrated(log types.Log) (*AllocationManagerViewSlasherMigrated, error) {
	event := new(AllocationManagerViewSlasherMigrated)
	if err := _AllocationManagerView.contract.UnpackLog(event, "SlasherMigrated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AllocationManagerViewSlasherUpdatedIterator is returned from FilterSlasherUpdated and is used to iterate over the raw logs and unpacked data for SlasherUpdated events raised by the AllocationManagerView contract.
type AllocationManagerViewSlasherUpdatedIterator struct {
	Event *AllocationManagerViewSlasherUpdated // Event containing the contract specifics and raw log

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
func (it *AllocationManagerViewSlasherUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AllocationManagerViewSlasherUpdated)
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
		it.Event = new(AllocationManagerViewSlasherUpdated)
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
func (it *AllocationManagerViewSlasherUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AllocationManagerViewSlasherUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AllocationManagerViewSlasherUpdated represents a SlasherUpdated event raised by the AllocationManagerView contract.
type AllocationManagerViewSlasherUpdated struct {
	OperatorSet OperatorSet
	Slasher     common.Address
	EffectBlock uint32
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSlasherUpdated is a free log retrieval operation binding the contract event 0x3873f29d7a65a4d75f5ba28909172f486216a1420e77c3c2720815951a6b4f57.
//
// Solidity: event SlasherUpdated((address,uint32) operatorSet, address slasher, uint32 effectBlock)
func (_AllocationManagerView *AllocationManagerViewFilterer) FilterSlasherUpdated(opts *bind.FilterOpts) (*AllocationManagerViewSlasherUpdatedIterator, error) {

	logs, sub, err := _AllocationManagerView.contract.FilterLogs(opts, "SlasherUpdated")
	if err != nil {
		return nil, err
	}
	return &AllocationManagerViewSlasherUpdatedIterator{contract: _AllocationManagerView.contract, event: "SlasherUpdated", logs: logs, sub: sub}, nil
}

// WatchSlasherUpdated is a free log subscription operation binding the contract event 0x3873f29d7a65a4d75f5ba28909172f486216a1420e77c3c2720815951a6b4f57.
//
// Solidity: event SlasherUpdated((address,uint32) operatorSet, address slasher, uint32 effectBlock)
func (_AllocationManagerView *AllocationManagerViewFilterer) WatchSlasherUpdated(opts *bind.WatchOpts, sink chan<- *AllocationManagerViewSlasherUpdated) (event.Subscription, error) {

	logs, sub, err := _AllocationManagerView.contract.WatchLogs(opts, "SlasherUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AllocationManagerViewSlasherUpdated)
				if err := _AllocationManagerView.contract.UnpackLog(event, "SlasherUpdated", log); err != nil {
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
func (_AllocationManagerView *AllocationManagerViewFilterer) ParseSlasherUpdated(log types.Log) (*AllocationManagerViewSlasherUpdated, error) {
	event := new(AllocationManagerViewSlasherUpdated)
	if err := _AllocationManagerView.contract.UnpackLog(event, "SlasherUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AllocationManagerViewStrategyAddedToOperatorSetIterator is returned from FilterStrategyAddedToOperatorSet and is used to iterate over the raw logs and unpacked data for StrategyAddedToOperatorSet events raised by the AllocationManagerView contract.
type AllocationManagerViewStrategyAddedToOperatorSetIterator struct {
	Event *AllocationManagerViewStrategyAddedToOperatorSet // Event containing the contract specifics and raw log

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
func (it *AllocationManagerViewStrategyAddedToOperatorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AllocationManagerViewStrategyAddedToOperatorSet)
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
		it.Event = new(AllocationManagerViewStrategyAddedToOperatorSet)
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
func (it *AllocationManagerViewStrategyAddedToOperatorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AllocationManagerViewStrategyAddedToOperatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AllocationManagerViewStrategyAddedToOperatorSet represents a StrategyAddedToOperatorSet event raised by the AllocationManagerView contract.
type AllocationManagerViewStrategyAddedToOperatorSet struct {
	OperatorSet OperatorSet
	Strategy    common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStrategyAddedToOperatorSet is a free log retrieval operation binding the contract event 0x7ab260fe0af193db5f4986770d831bda4ea46099dc817e8b6716dcae8af8e88b.
//
// Solidity: event StrategyAddedToOperatorSet((address,uint32) operatorSet, address strategy)
func (_AllocationManagerView *AllocationManagerViewFilterer) FilterStrategyAddedToOperatorSet(opts *bind.FilterOpts) (*AllocationManagerViewStrategyAddedToOperatorSetIterator, error) {

	logs, sub, err := _AllocationManagerView.contract.FilterLogs(opts, "StrategyAddedToOperatorSet")
	if err != nil {
		return nil, err
	}
	return &AllocationManagerViewStrategyAddedToOperatorSetIterator{contract: _AllocationManagerView.contract, event: "StrategyAddedToOperatorSet", logs: logs, sub: sub}, nil
}

// WatchStrategyAddedToOperatorSet is a free log subscription operation binding the contract event 0x7ab260fe0af193db5f4986770d831bda4ea46099dc817e8b6716dcae8af8e88b.
//
// Solidity: event StrategyAddedToOperatorSet((address,uint32) operatorSet, address strategy)
func (_AllocationManagerView *AllocationManagerViewFilterer) WatchStrategyAddedToOperatorSet(opts *bind.WatchOpts, sink chan<- *AllocationManagerViewStrategyAddedToOperatorSet) (event.Subscription, error) {

	logs, sub, err := _AllocationManagerView.contract.WatchLogs(opts, "StrategyAddedToOperatorSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AllocationManagerViewStrategyAddedToOperatorSet)
				if err := _AllocationManagerView.contract.UnpackLog(event, "StrategyAddedToOperatorSet", log); err != nil {
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
func (_AllocationManagerView *AllocationManagerViewFilterer) ParseStrategyAddedToOperatorSet(log types.Log) (*AllocationManagerViewStrategyAddedToOperatorSet, error) {
	event := new(AllocationManagerViewStrategyAddedToOperatorSet)
	if err := _AllocationManagerView.contract.UnpackLog(event, "StrategyAddedToOperatorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AllocationManagerViewStrategyRemovedFromOperatorSetIterator is returned from FilterStrategyRemovedFromOperatorSet and is used to iterate over the raw logs and unpacked data for StrategyRemovedFromOperatorSet events raised by the AllocationManagerView contract.
type AllocationManagerViewStrategyRemovedFromOperatorSetIterator struct {
	Event *AllocationManagerViewStrategyRemovedFromOperatorSet // Event containing the contract specifics and raw log

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
func (it *AllocationManagerViewStrategyRemovedFromOperatorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AllocationManagerViewStrategyRemovedFromOperatorSet)
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
		it.Event = new(AllocationManagerViewStrategyRemovedFromOperatorSet)
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
func (it *AllocationManagerViewStrategyRemovedFromOperatorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AllocationManagerViewStrategyRemovedFromOperatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AllocationManagerViewStrategyRemovedFromOperatorSet represents a StrategyRemovedFromOperatorSet event raised by the AllocationManagerView contract.
type AllocationManagerViewStrategyRemovedFromOperatorSet struct {
	OperatorSet OperatorSet
	Strategy    common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStrategyRemovedFromOperatorSet is a free log retrieval operation binding the contract event 0x7b4b073d80dcac55a11177d8459ad9f664ceeb91f71f27167bb14f8152a7eeee.
//
// Solidity: event StrategyRemovedFromOperatorSet((address,uint32) operatorSet, address strategy)
func (_AllocationManagerView *AllocationManagerViewFilterer) FilterStrategyRemovedFromOperatorSet(opts *bind.FilterOpts) (*AllocationManagerViewStrategyRemovedFromOperatorSetIterator, error) {

	logs, sub, err := _AllocationManagerView.contract.FilterLogs(opts, "StrategyRemovedFromOperatorSet")
	if err != nil {
		return nil, err
	}
	return &AllocationManagerViewStrategyRemovedFromOperatorSetIterator{contract: _AllocationManagerView.contract, event: "StrategyRemovedFromOperatorSet", logs: logs, sub: sub}, nil
}

// WatchStrategyRemovedFromOperatorSet is a free log subscription operation binding the contract event 0x7b4b073d80dcac55a11177d8459ad9f664ceeb91f71f27167bb14f8152a7eeee.
//
// Solidity: event StrategyRemovedFromOperatorSet((address,uint32) operatorSet, address strategy)
func (_AllocationManagerView *AllocationManagerViewFilterer) WatchStrategyRemovedFromOperatorSet(opts *bind.WatchOpts, sink chan<- *AllocationManagerViewStrategyRemovedFromOperatorSet) (event.Subscription, error) {

	logs, sub, err := _AllocationManagerView.contract.WatchLogs(opts, "StrategyRemovedFromOperatorSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AllocationManagerViewStrategyRemovedFromOperatorSet)
				if err := _AllocationManagerView.contract.UnpackLog(event, "StrategyRemovedFromOperatorSet", log); err != nil {
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
func (_AllocationManagerView *AllocationManagerViewFilterer) ParseStrategyRemovedFromOperatorSet(log types.Log) (*AllocationManagerViewStrategyRemovedFromOperatorSet, error) {
	event := new(AllocationManagerViewStrategyRemovedFromOperatorSet)
	if err := _AllocationManagerView.contract.UnpackLog(event, "StrategyRemovedFromOperatorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
