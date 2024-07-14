// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Slasher

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

// ISignatureUtilsSignatureWithSaltAndExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithSaltAndExpiry struct {
	Signature []byte
	Salt      [32]byte
	Expiry    *big.Int
}

// ISlasherMagnitudeAdjustment is an auto generated low-level Go binding around an user-defined struct.
type ISlasherMagnitudeAdjustment struct {
	OperatorSet   ISlasherOperatorSet
	MagnitudeDiff uint64
}

// ISlasherMagnitudeAdjustmentsParam is an auto generated low-level Go binding around an user-defined struct.
type ISlasherMagnitudeAdjustmentsParam struct {
	Strategy                common.Address
	MagnitudeAdjustmentType uint8
	MagnitudeAdjustments    []ISlasherMagnitudeAdjustment
}

// ISlasherMagnitudeUpdate is an auto generated low-level Go binding around an user-defined struct.
type ISlasherMagnitudeUpdate struct {
	Timestamp uint32
	Magnitude uint64
}

// ISlasherOperatorSet is an auto generated low-level Go binding around an user-defined struct.
type ISlasherOperatorSet struct {
	Avs common.Address
	Id  uint32
}

// ISlasherTotalAndNonslashableUpdate is an auto generated low-level Go binding around an user-defined struct.
type ISlasherTotalAndNonslashableUpdate struct {
	Timestamp               uint32
	TotalMagnitude          uint64
	NonslashableMagnitude   uint64
	CumulativeAllocationSum uint64
}

// SlasherMetaData contains all meta data concerning the Slasher contract.
var SlasherMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_strategyManager\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"},{\"name\":\"_delegationManager\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DOMAIN_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAGNITUDE_CONCENTRATION_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAGNITUDE_DILUTION_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"REALLOCATION_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allocatorSaltIsSpent\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateMagnitudeConcentrationDigestHash\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"nonslashableToDecrement\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateMagnitudeDilutionDigestHash\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"nonslashableToAdd\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateReallocationDigestHash\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"adjustmentParams\",\"type\":\"tuple[]\",\"internalType\":\"structISlasher.MagnitudeAdjustmentsParam[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"magnitudeAdjustmentType\",\"type\":\"uint8\",\"internalType\":\"enumISlasher.MagnitudeAdjustmentType\"},{\"name\":\"magnitudeAdjustments\",\"type\":\"tuple[]\",\"internalType\":\"structISlasher.MagnitudeAdjustment[]\",\"components\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structISlasher.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"magnitudeDiff\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}]},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkPendingMagnitudeUpdates\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structISlasher.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkPendingTotalMagnitudeUpdates\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"domainSeparator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocatorFor\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMagnitudeUpdate\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structISlasher.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"timestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"magnitudeUpdate\",\"type\":\"tuple\",\"internalType\":\"structISlasher.MagnitudeUpdate\",\"components\":[{\"name\":\"timestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"magnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSlashableBips\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structISlasher.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"slashableBips\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSlashedRate\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structISlasher.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"epoch\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalAndNonslashableUpdate\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"timestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"totalAndNonslashableUpdate\",\"type\":\"tuple\",\"internalType\":\"structISlasher.TotalAndNonslashableUpdate\",\"components\":[{\"name\":\"timestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"totalMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"nonslashableMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"cumulativeAllocationSum\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"initialPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lastSlashed\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"queueMagnitudeConcentration\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"nonslashableToDecrement\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"allocatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"newNonslashableMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"newTotalMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"queueMagnitudeDilution\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"nonslashableToAdd\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"allocatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"newNonslashableMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"newTotalMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"queueReallocation\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"adjustmentParams\",\"type\":\"tuple[]\",\"internalType\":\"structISlasher.MagnitudeAdjustmentsParam[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"magnitudeAdjustmentType\",\"type\":\"uint8\",\"internalType\":\"enumISlasher.MagnitudeAdjustmentType\"},{\"name\":\"magnitudeAdjustments\",\"type\":\"tuple[]\",\"internalType\":\"structISlasher.MagnitudeAdjustment[]\",\"components\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structISlasher.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"magnitudeDiff\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}]},{\"name\":\"allocatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"effectTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"shareScalingFactor\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"shareScalingFactorAtEpoch\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"epoch\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"slashOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"bipsToSlash\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slashingEpochHistory\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"slashingUpdates\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"slashingRate\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"scalingFactor\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"strategyManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MagnitudeDecremented\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structISlasher.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"updatedSlashableMagnitude\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"effectTimestamp\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MagnitudeUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structISlasher.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"effectTimestamp\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"slashableMagnitude\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NonslashableMagnitudeDecremented\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"updatedNonslashableMagnitude\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"effectTimestamp\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSlashed\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structISlasher.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"bipsToSlash\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"contractIStrategy[]\"},{\"name\":\"slashingRates\",\"type\":\"uint64[]\",\"indexed\":false,\"internalType\":\"uint64[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TotalAndNonslashableMagnitudeUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"effectTimestamp\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"totalSlashableMagnitude\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"nonslashableMagnitude\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"cumulativeAllocationSum\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x60e06040523480156200001157600080fd5b50604051620046e7380380620046e7833981016040819052620000349162000069565b6001600160a01b039182166080521660a0524660c052620000a8565b6001600160a01b03811681146200006657600080fd5b50565b600080604083850312156200007d57600080fd5b82516200008a8162000050565b60208401519092506200009d8162000050565b809150509250929050565b60805160a05160c051614608620000df60003960006124c001526000818161064b01526116dc0152600061042701526146086000f3fe608060405234801561001057600080fd5b50600436106102325760003560e01c80635c120057116101305780639292a40d116100b8578063e49a1e841161007c578063e49a1e841461066d578063e565c62114610680578063f2fde38b14610693578063f698da25146106a6578063fabc1cbc146106ae57600080fd5b80639292a40d146105d2578063a952bff9146105e5578063ba7280621461060c578063d1920e3714610633578063df5cf7231461064657600080fd5b8063715018a6116100ff578063715018a61461056c5780638837866a14610574578063886f11951461059b5780638d60eab5146105ae5780638da5cb5b146105c157600080fd5b80635c1200571461052b5780635c975abb1461053e57806364565d3d1461054657806367e988de1461055957600080fd5b8063334f00d6116101be57806346fb3a001161018257806346fb3a001461049c5780634b0284de146104da578063595c6a67146104ed5780635ab112d6146104f55780635ac86ab71461050857600080fd5b8063334f00d6146103b2578063364bce6f146103c557806339b70e38146104225780633f5f49e9146104615780634279a7e61461047457600080fd5b8063136439dd11610205578063136439dd146102f457806316449dae146103075780631794bb3c146103445780631c32f5c61461035757806320606b701461037d57600080fd5b8063017ec6f9146102375780630b1b781e1461026f57806310c0a606146102b457806310d67a2f146102df575b600080fd5b61024a610245366004613743565b6106c1565b604080516001600160401b039384168152929091166020830152015b60405180910390f35b61024a61027d366004613809565b60986020908152600093845260408085208252928452828420905282529020546001600160401b0380821691600160401b90041682565b6102c76102c2366004613866565b61096e565b6040516001600160401b039091168152602001610266565b6102f26102ed3660046138c1565b6109d5565b005b6102f26103023660046138de565b610a88565b61031a61031536600461399f565b610bc7565b60408051825163ffffffff1681526020928301516001600160401b03169281019290925201610266565b6102f26103523660046139df565b610dc4565b61036a610365366004613a20565b610eee565b60405161ffff9091168152602001610266565b6103a47f8cad95687ba82c2ce50e74f7b754645e5117c3a5bec8151c0726d5857980a86681565b604051908152602001610266565b6102c76103c0366004613a5f565b610f4d565b6103d86103d3366004613809565b610f9d565b6040516102669190815163ffffffff1681526020808301516001600160401b0390811691830191909152604080840151821690830152606092830151169181019190915260800190565b6104497f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b039091168152602001610266565b6103a461046f366004613a98565b61119a565b6104876104823660046139df565b611205565b60405163ffffffff9091168152602001610266565b6104ca6104aa366004613b29565b609b60209081526000928352604080842090915290825290205460ff1681565b6040519015158152602001610266565b6103a46104e8366004613a98565b61125b565b6102f261129f565b610487610503366004613a5f565b611366565b6104ca610516366004613b55565b606654600160ff9092169190911b9081161490565b6102f2610539366004613b78565b61140b565b6066546103a4565b6104ca610554366004613a5f565b611622565b6104496105673660046138c1565b6116ba565b6102f2611750565b6103a47fea04577b50ab8413a5b69ebd4fc852645911c5d2220917f1717c99c7988f303781565b606554610449906001600160a01b031681565b6103a46105bc366004613c61565b611764565b6033546001600160a01b0316610449565b6104876105e0366004613cc7565b6117ae565b6103a47f33daef3b93f9ce098ed3ad97c293d3265bed7889b5b9c8dfe95c27cf7b8387ee81565b6103a47fd6b0dcfea99ac389c204d57700d0962238af3158458b4252d85998bfac29668e81565b6104ca610641366004613d42565b6118d5565b6104497f000000000000000000000000000000000000000000000000000000000000000081565b6102c761067b366004613809565b611a2a565b61024a61068e366004613743565b611a9e565b6102f26106a13660046138c1565b611cc6565b6103a4611d3c565b6102f26106bc3660046138de565b611d4b565b60008060006106cf896116ba565b9050336001600160a01b038216146107095760006106fa8a8a8a8a8a8a602001358b6040013561119a565b9050610707828287611ea7565b505b6000610718621baf8042613da0565b905060005b88811015610960576107508b8b8b8481811061073b5761073b613db8565b905060200201602081019061055491906138c1565b6107755760405162461bcd60e51b815260040161076c90613dce565b60405180910390fd5b60006107a78c8c8c8581811061078d5761078d613db8565b90506020020160208101906107a291906138c1565b61202d565b80549091506001600160401b03600160601b8204811691600160201b9004166b204fce5e3e250261100000008b8b868181106107e5576107e5613db8565b90506020020160208101906107fa9190613e58565b6001600160401b0316111561089d5760405162461bcd60e51b815260206004820152605a60248201527f536c61736865722e71756575654d61676e6974756465436f6e63656e7472617460448201527f696f6e3a206e6f6e736c61736861626c6541646465642065786365656473204660648201527f495845445f544f54414c5f4d41474e49545544455f4c494d4954000000000000608482015260a40161076c565b8a8a858181106108af576108af613db8565b90506020020160208101906108c49190613e58565b6108ce9082613e73565b96508a8a858181106108e2576108e2613db8565b90506020020160208101906108f79190613e58565b6109019083613e73565b975061094c8e8e8e8781811061091957610919613db8565b905060200201602081019061092e91906138c1565b855486908b908d90600160a01b90046001600160401b03168b612173565b5050508061095990613e9e565b905061071d565b505050965096945050505050565b60008060008061097f8888876122fa565b9150915080156109c9576001600160a01b038089166000908152609860209081526040808320938b16835292815282822063ffffffff86168352905220546001600160401b031692505b50909695505050505050565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610a28573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a4c9190613eb9565b6001600160a01b0316336001600160a01b031614610a7c5760405162461bcd60e51b815260040161076c90613ed6565b610a85816123c5565b50565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015610ad0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610af49190613f20565b610b105760405162461bcd60e51b815260040161076c90613f42565b60665481811614610b895760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c6974790000000000000000606482015260840161076c565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b604080518082018252600080825260208083018290526001600160a01b038881168352609c8252848320888216845282528483208751909116835281528382208682015163ffffffff16835290529190912054805b8015610d4a576001600160a01b038088166000908152609c602090815260408083208a851684528252808320895190941683529281528282208882015163ffffffff9081168452915291902090851690610c77600184613f8a565b81548110610c8757610c87613db8565b60009182526020909120015463ffffffff1611610d3a576001600160a01b038088166000908152609c602090815260408083208a851684528252808320895190941683529281528282208882015163ffffffff168352905220610ceb600183613f8a565b81548110610cfb57610cfb613db8565b60009182526020918290206040805180820190915291015463ffffffff81168252600160201b90046001600160401b0316918101919091529250610d4a565b610d4381613fa1565b9050610c1c565b5060405162461bcd60e51b815260206004820152604260248201527f536c61736865722e6765744d61676e69747564655570646174653a206e6f206d60448201527f61676e69747564652075706461746520666f756e642061742074696d6573746160648201526106d760f41b608482015260a40161076c565b600054610100900460ff1615808015610de45750600054600160ff909116105b80610dfe5750303b158015610dfe575060005460ff166001145b610e615760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161076c565b6000805460ff191660011790558015610e84576000805461ff0019166101001790555b610e8c6124bc565b609755610e998383612585565b610ea28461266f565b8015610ee8576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050565b600080610efd85848642610bc7565b90506000610f0c868542610f9d565b905080602001516001600160401b031661271083602001516001600160401b0316610f379190613fb8565b610f419190613fed565b925050505b9392505050565b6001600160a01b0380831660009081526099602090815260408083209385168352929052908120546001600160401b031680610f9457670de0b6b3a7640000915050610f97565b90505b92915050565b6040805160808101825260008082526020808301829052828401829052606083018290526001600160a01b038781168352609d8252848320908716835290529190912054805b80156110fa576001600160a01b038087166000908152609d6020908152604080832093891683529290522063ffffffff851690611021600184613f8a565b8154811061103157611031613db8565b60009182526020909120015463ffffffff16116110ea576001600160a01b038087166000908152609d60209081526040808320938916835292905220611078600183613f8a565b8154811061108857611088613db8565b600091825260209182902060408051608081018252919092015463ffffffff811682526001600160401b03600160201b8204811694830194909452600160601b8104841692820192909252600160a01b909104909116606082015292506110fa565b6110f381613fa1565b9050610fe3565b5060405162461bcd60e51b815260206004820152606260248201527f536c61736865722e676574546f74616c416e644e6f6e736c61736861626c655560448201527f70646174653a206e6f20746f74616c616e646e6f6e736c61736861626c65206d60648201527f61676e69747564652075706461746520666f756e642061742074696d6573746160848201526106d760f41b60a482015260c40161076c565b60006111f97fd6b0dcfea99ac389c204d57700d0962238af3158458b4252d85998bfac29668e898989898989896040516020016111de989796959493929190614001565b604051602081830303815290604052805190602001206126c1565b98975050505050505050565b609a602052826000526040600020602052816000526040600020818154811061122d57600080fd5b906000526020600020906008918282040191900660040292509250509054906101000a900463ffffffff1681565b60006111f97fea04577b50ab8413a5b69ebd4fc852645911c5d2220917f1717c99c7988f3037898989898989896040516020016111de989796959493929190614001565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa1580156112e7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061130b9190613f20565b6113275760405162461bcd60e51b815260040161076c90613f42565b600019606681905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b6001600160a01b038083166000908152609a602090815260408083209385168352929052908120548061139d576000915050610f97565b6001600160a01b038085166000908152609a602090815260408083209387168352929052206113cd600183613f8a565b815481106113dd576113dd613db8565b90600052602060002090600891828204019190066004029054906101000a900463ffffffff16915050610f97565b8063ffffffff16600010801561142857506127108163ffffffff16105b61148d5760405162461bcd60e51b815260206004820152603060248201527f536c61736865722e5f736c617368526571756573746564426970733a20696e7660448201526f0c2d8d2c840c4d2e0e6a8dea6d8c2e6d60831b606482015260840161076c565b6040805180820190915233815263ffffffff8416602082015260006114b0612708565b905060005b84518110156116195760008582815181106114d2576114d2613db8565b6020026020010151905060006114e9898684610eee565b6114fd9061ffff1663ffffffff88166140b5565b9050600061150b8a84610f4d565b905060006115198284612713565b6001600160a01b03808d166000818152609a602090815260408083209990941680835298815283822080546001810182559083528183206008820401805463ffffffff808f1660046007909516949094026101000a848102910219909116179055845180860186526001600160401b03998a168152958916868301818152858552609884528685208c865284528685209285529183528584209651875492518b16600160401b026fffffffffffffffffffffffffffffffff199093169a1699909917179094559081526099835281812096815295909152909320805467ffffffffffffffff191690921790915550611612905081613e9e565b90506114b5565b50505050505050565b6001600160a01b038083166000908152609d60209081526040808320938516835292905290812054600681101561165d576001915050610f97565b6001600160a01b038085166000908152609d60209081526040808320938716835292905220429061168f600684613f8a565b8154811061169f5761169f613db8565b60009182526020909120015463ffffffff1610949350505050565b60405163c5e480db60e01b81526001600160a01b0382811660048301526000917f00000000000000000000000000000000000000000000000000000000000000009091169063c5e480db90602401606060405180830381865afa158015611725573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061174991906140e4565b5192915050565b611758612868565b611762600061266f565b565b60006117a47f33daef3b93f9ce098ed3ad97c293d3265bed7889b5b9c8dfe95c27cf7b8387ee87878787876040516020016111de96959493929190614207565b9695505050505050565b6000806117ba866116ba565b9050336001600160a01b038216146117f25760006117e387878787602001358860400135611764565b90506117f0828286611ea7565b505b6117ff621baf8042613da0565b915060005b848110156118cb57600086868381811061182057611820613db8565b90506020028101906118329190614331565b611843906040810190602001614351565b600181111561185457611854614164565b141561188d576118888787878481811061187057611870613db8565b90506020028101906118829190614331565b856128c2565b6118bb565b6118bb878787848181106118a3576118a3613db8565b90506020028101906118b59190614331565b85612b0f565b6118c481613e9e565b9050611804565b5050949350505050565b6001600160a01b038084166000908152609c6020908152604080832093861683529281529181209091829190829061190f908601866138c1565b6001600160a01b03166001600160a01b031681526020019081526020016000206000846020016020810190611944919061436c565b63ffffffff16815260208101919091526040016000205490506003811015611970576001915050610f46565b6001600160a01b038086166000908152609c60209081526040808320938816835292815291812042929091906119a8908701876138c1565b6001600160a01b03166001600160a01b0316815260200190815260200160002060008560200160208101906119dd919061436c565b63ffffffff16815260208101919091526040016000206119fe600384613f8a565b81548110611a0e57611a0e613db8565b60009182526020909120015463ffffffff161095945050505050565b6000670de0b6b3a76400008180611a428787876122fa565b915091508015611a93576001600160a01b038781166000908152609860209081526040808320938a16835292815282822063ffffffff8616835290522054600160401b90046001600160401b031692505b509095945050505050565b6000806000611aac896116ba565b9050336001600160a01b03821614611ae6576000611ad78a8a8a8a8a8a602001358b6040013561125b565b9050611ae4828287611ea7565b505b6000611af5621baf8042613da0565b905060005b8881101561096057611b188b8b8b8481811061073b5761073b613db8565b611b345760405162461bcd60e51b815260040161076c90613dce565b6000611b4c8c8c8c8581811061078d5761078d613db8565b80549091506001600160401b03600160601b8204811691600160201b900416818b8b86818110611b7e57611b7e613db8565b9050602002016020810190611b939190613e58565b6001600160401b03161115611c365760405162461bcd60e51b815260206004820152605a60248201527f536c61736865722e71756575654d61676e6974756465436f6e63656e7472617460448201527f696f6e3a206e6f6e736c61736861626c6544656372656d656e7465642065786360648201527f65656473206e6f6e736c61736861626c654d61676e6974756465000000000000608482015260a40161076c565b8a8a85818110611c4857611c48613db8565b9050602002016020810190611c5d9190613e58565b611c679082614389565b96508a8a85818110611c7b57611c7b613db8565b9050602002016020810190611c909190613e58565b611c9a9083614389565b9750611cb28e8e8e8781811061091957610919613db8565b50505080611cbf90613e9e565b9050611afa565b611cce612868565b6001600160a01b038116611d335760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161076c565b610a858161266f565b6000611d466124bc565b905090565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611d9e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611dc29190613eb9565b6001600160a01b0316336001600160a01b031614611df25760405162461bcd60e51b815260040161076c90613ed6565b606654198119606654191614611e705760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c6974790000000000000000606482015260840161076c565b606681905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c90602001610bbc565b4281604001351015611f215760405162461bcd60e51b815260206004820152603e60248201527f536c61736865722e5f766572696679416c6c6f6361746f725369676e6174757260448201527f653a20616c6c6f6361746f72207369676e617475726520657870697265640000606482015260840161076c565b6001600160a01b0383166000908152609b6020908152604080832084830135845290915290205460ff1615611fad5760405162461bcd60e51b815260206004820152602c60248201527f4156534469726563746f72792e7570646174655374616e646279506172616d7360448201526b0e881cd85b1d081cdc195b9d60a21b606482015260840161076c565b611ff68383611fbc84806143b1565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250612da292505050565b6001600160a01b039092166000908152609b6020908152604080832094820135835293905291909120805460ff1916600117905550565b6001600160a01b038083166000908152609d6020908152604080832093851683529290529081205480612120576001600160a01b038481166000908152609d60209081526040808320938716835292815282822083516080810185528381528083018481529481018481526060820185815283546001810185559386529390942090519101805494519351925163ffffffff9092166001600160601b031990951694909417600160201b6001600160401b039485160217600160601b600160e01b031916600160601b9284169290920267ffffffffffffffff60a01b191691909117600160a01b92909116919091021790555b6001600160a01b038085166000908152609d60209081526040808320938716835292905220612150600183613f8a565b8154811061216057612160613db8565b9060005260206000200191505092915050565b845463ffffffff828116911614156121d75784546001600160401b03848116600160601b0267ffffffffffffffff60601b19918716600160201b029190911673ffffffffffffffffffffffffffffffff0000000019909216919091171785556122b0565b6001600160a01b038088166000908152609d60209081526040808320938a168352928152828220835160808101855263ffffffff80871682526001600160401b03808b168386019081528a82169784019788528982166060850190815285546001810187559588529590962092519290930180549551965194518416600160a01b0267ffffffffffffffff60a01b19958516600160601b0295909516600160601b600160e01b031997909416600160201b026001600160601b031990961692909116919091179390931793909316929092179190911790555b7f2bc1462fd4652ef49aae855563647b6564ff41eb9db10038d4b08e01d34f30e18787838787876040516122e9969594939291906143f7565b60405180910390a150505050505050565b6001600160a01b038084166000908152609a602090815260408083209386168352929052908120548190819081905b80156123b8576001600160a01b038089166000908152609a60209081526040808320938b1683529290522061235f600183613f8a565b8154811061236f5761236f613db8565b6000918252602090912060088204015460079091166004026101000a900463ffffffff9081169350861683116123a857600191506123b8565b6123b181613fa1565b9050612329565b5090969095509350505050565b6001600160a01b0381166124535760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a40161076c565b606554604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1606580546001600160a01b0319166001600160a01b0392909216919091179055565b60007f00000000000000000000000000000000000000000000000000000000000000004614156124ed575060975490565b50604080518082018252600a81526922b4b3b2b72630bcb2b960b11b60209182015281517f8cad95687ba82c2ce50e74f7b754645e5117c3a5bec8151c0726d5857980a866818301527f71b625cfad44bac63b13dba07f2e1d6084ee04b6f8752101ece6126d584ee6ea81840152466060820152306080808301919091528351808303909101815260a0909101909252815191012090565b6065546001600160a01b03161580156125a657506001600160a01b03821615155b6126285760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a40161076c565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a261266b826123c5565b5050565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b60006126cb6124bc565b60405161190160f01b6020820152602281019190915260428101839052606201604051602081830303815290604052805190602001209050919050565b6000611d4642612f61565b60006001600160401b0382166127615760405162461bcd60e51b815260206004820152601360248201527263616e6e6f7420736c61736820666f7220302560681b604482015260640161076c565b6305f5e1006001600160401b03831611156127ca5760405162461bcd60e51b815260206004820152602360248201527f63616e6e6f7420736c617368206d6f7265207468616e2031303025206174206f6044820152626e636560e81b606482015260840161076c565b60006001600160401b0383166305f5e100148061282757506001600160401b0380841690851661280e670de0b6b3a76400006bffffffffffffffffffffffff613fb8565b61281a90600019613fed565b6128249190613fed565b10155b1561283a57506001600160401b03610f94565b612848836305f5e100614389565b6128566305f5e100866140b5565b6128609190614441565b949350505050565b6033546001600160a01b031633146117625760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161076c565b60006128d160208401846138c1565b905060006128df858361202d565b905060006128ee868442610f9d565b606081015183549192506000916129159190600160a01b90046001600160401b0316614389565b82604001516129249190614389565b90506000805b6129376040890189614467565b9050811015612aca573661294e60408a018a614467565b8381811061295e5761295e613db8565b6060029190910191506000905061297860408b018b614467565b8481811061298857612988613db8565b90506060020160400160208101906129a09190613e58565b90506129ad8b89846118d5565b6129c95760405162461bcd60e51b815260040161076c906144af565b846001600160401b0316816001600160401b03161115612a5f5760405162461bcd60e51b815260206004820152604560248201527f536c61736865722e5f7175657565416c6c6f636174696f6e3a206d61676e697460448201527f756465446966662065786365656473206e6f6e736c61736861626c654d61676e606482015264697475646560d81b608482015260a40161076c565b6000612a6c8c8a85612ff9565b9050612a9e8c8a8584868660000160049054906101000a90046001600160401b0316612a989190613e73565b8f61320c565b612aa88287614389565b9550612ab48286613e73565b945050505080612ac390613e9e565b905061292a565b508354612b05908990879087906001600160401b03600160201b82048116918891612aff918991600160a01b90910416613e73565b8c612173565b5050505050505050565b6000612b1e60208401846138c1565b90506000612b2c858361202d565b8054909150600160601b90046001600160401b031660005b612b516040870187614467565b9050811015612d0d5736612b686040880188614467565b83818110612b7857612b78613db8565b60600291909101915060009050612b926040890189614467565b84818110612ba257612ba2613db8565b9050606002016040016020810190612bba9190613e58565b9050612bc78987846118d5565b612be35760405162461bcd60e51b815260040161076c906144af565b6000612bf08a8885612ff9565b80549091506001600160401b03600160201b9091048116908316811015612c8d5760405162461bcd60e51b8152602060048201526044602482018190527f536c61736865722e71756575654465616c6c6f636174696f6e3a206d61676e69908201527f7475646544696666206578636565647320616c6c6f6361746564206d61676e696064820152637475646560e01b608482015260a40161076c565b612ca48b898685612c9e8887614389565b8e61320c565b7f770e5720664253f46b99dcb3b922235f8ee78191bdb00bf7efb8c48ba3ab396d8b89868c612cd38887614389565b604051612ce4959493929190614501565b60405180910390a1612cf68387613e73565b95505050505080612d0690613e9e565b9050612b44565b508154612d3d908790859085906001600160401b03600160201b82048116918791600160a01b909104168a612173565b81546040517f2bc1462fd4652ef49aae855563647b6564ff41eb9db10038d4b08e01d34f30e191612d92918991879189916001600160401b03600160201b82048116928992600160a01b9004909116906143f7565b60405180910390a1505050505050565b6001600160a01b0383163b15612ec157604051630b135d3f60e11b808252906001600160a01b03851690631626ba7e90612de2908690869060040161454b565b602060405180830381865afa158015612dff573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612e2391906145a8565b6001600160e01b03191614612ebc5760405162461bcd60e51b815260206004820152605360248201527f454950313237315369676e61747572655574696c732e636865636b5369676e6160448201527f747572655f454950313237313a2045524331323731207369676e6174757265206064820152721d995c9a599a58d85d1a5bdb8819985a5b1959606a1b608482015260a40161076c565b505050565b826001600160a01b0316612ed5838361335d565b6001600160a01b031614612ebc5760405162461bcd60e51b815260206004820152604760248201527f454950313237315369676e61747572655574696c732e636865636b5369676e6160448201527f747572655f454950313237313a207369676e6174757265206e6f742066726f6d6064820152661039b4b3b732b960c91b608482015260a40161076c565b6000635fc63040821015612fdd5760405162461bcd60e51b815260206004820152603d60248201527f45706f63685574696c732e67657445706f636846726f6d54696d657374616d7060448201527f3a2074696d657374616d70206973206265666f72652067656e65736973000000606482015260840161076c565b62093a80612fef635fc6304084613f8a565b610f979190613fed565b6001600160a01b038084166000908152609c60209081526040808320938616835292815291812090918291908290613033908601866138c1565b6001600160a01b03166001600160a01b031681526020019081526020016000206000846020016020810190613068919061436c565b63ffffffff16815260208101919091526040016000205490508061315e576001600160a01b038086166000908152609c602090815260408083209388168352928152918120916130ba908601866138c1565b6001600160a01b03166001600160a01b0316815260200190815260200160002060008460200160208101906130ef919061436c565b63ffffffff908116825260208083019390935260409182016000908120835180850190945281845283850182815281546001810183559183529490912092519201805493516001600160401b0316600160201b026001600160601b031990941692909116919091179190911790555b6001600160a01b038086166000908152609c60209081526040808320938816835292815291812091613192908601866138c1565b6001600160a01b03166001600160a01b0316815260200190815260200160002060008460200160208101906131c7919061436c565b63ffffffff16815260208101919091526040016000206131e8600183613f8a565b815481106131f8576131f8613db8565b906000526020600020019150509392505050565b825463ffffffff828116911614156132475782546bffffffffffffffff000000001916600160201b6001600160401b03841602178355613326565b6001600160a01b038087166000908152609c6020908152604080832093891683529281529181209161327b908701876138c1565b6001600160a01b03166001600160a01b0316815260200190815260200160002060008560200160208101906132b0919061436c565b63ffffffff908116825260208083019390935260409182016000908120835180850190945285831684526001600160401b038088168587019081528254600181018455928452959092209351930180549451909116600160201b026001600160601b031990941692909116919091179190911790555b7f770e5720664253f46b99dcb3b922235f8ee78191bdb00bf7efb8c48ba3ab396d8686868486604051612d92959493929190614501565b600080600061336c8585613381565b91509150613379816133f1565b509392505050565b6000808251604114156133b85760208301516040840151606085015160001a6133ac878285856135ac565b945094505050506133ea565b8251604014156133e257602083015160408401516133d7868383613699565b9350935050506133ea565b506000905060025b9250929050565b600081600481111561340557613405614164565b141561340e5750565b600181600481111561342257613422614164565b14156134705760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e61747572650000000000000000604482015260640161076c565b600281600481111561348457613484614164565b14156134d25760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e67746800604482015260640161076c565b60038160048111156134e6576134e6614164565b141561353f5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b606482015260840161076c565b600481600481111561355357613553614164565b1415610a855760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b606482015260840161076c565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08311156135e35750600090506003613690565b8460ff16601b141580156135fb57508460ff16601c14155b1561360c5750600090506004613690565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015613660573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b03811661368957600060019250925050613690565b9150600090505b94509492505050565b6000806001600160ff1b038316816136b660ff86901c601b613da0565b90506136c4878288856135ac565b935093505050935093915050565b6001600160a01b0381168114610a8557600080fd5b60008083601f8401126136f957600080fd5b5081356001600160401b0381111561371057600080fd5b6020830191508360208260051b85010111156133ea57600080fd5b60006060828403121561373d57600080fd5b50919050565b6000806000806000806080878903121561375c57600080fd5b8635613767816136d2565b955060208701356001600160401b038082111561378357600080fd5b61378f8a838b016136e7565b909750955060408901359150808211156137a857600080fd5b6137b48a838b016136e7565b909550935060608901359150808211156137cd57600080fd5b506137da89828a0161372b565b9150509295509295509295565b63ffffffff81168114610a8557600080fd5b8035613804816137e7565b919050565b60008060006060848603121561381e57600080fd5b8335613829816136d2565b92506020840135613839816136d2565b91506040840135613849816137e7565b809150509250925092565b60006040828403121561373d57600080fd5b60008060008060a0858703121561387c57600080fd5b8435613887816136d2565b93506020850135613897816136d2565b92506138a68660408701613854565b915060808501356138b6816137e7565b939692955090935050565b6000602082840312156138d357600080fd5b8135610f94816136d2565b6000602082840312156138f057600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b0381118282101715613935576139356138f7565b604052919050565b60006040828403121561394f57600080fd5b604051604081018181106001600160401b0382111715613971576139716138f7565b6040529050808235613982816136d2565b81526020830135613992816137e7565b6020919091015292915050565b60008060008060a085870312156139b557600080fd5b84356139c0816136d2565b935060208501356139d0816136d2565b92506138a6866040870161393d565b6000806000606084860312156139f457600080fd5b83356139ff816136d2565b92506020840135613a0f816136d2565b929592945050506040919091013590565b600080600060808486031215613a3557600080fd5b8335613a40816136d2565b9250613a4f856020860161393d565b91506060840135613849816136d2565b60008060408385031215613a7257600080fd5b8235613a7d816136d2565b91506020830135613a8d816136d2565b809150509250929050565b600080600080600080600060a0888a031215613ab357600080fd5b8735613abe816136d2565b965060208801356001600160401b0380821115613ada57600080fd5b613ae68b838c016136e7565b909850965060408a0135915080821115613aff57600080fd5b50613b0c8a828b016136e7565b989b979a5095989597966060870135966080013595509350505050565b60008060408385031215613b3c57600080fd5b8235613b47816136d2565b946020939093013593505050565b600060208284031215613b6757600080fd5b813560ff81168114610f9457600080fd5b60008060008060808587031215613b8e57600080fd5b8435613b99816136d2565b9350602085810135613baa816137e7565b935060408601356001600160401b0380821115613bc657600080fd5b818801915088601f830112613bda57600080fd5b813581811115613bec57613bec6138f7565b8060051b9150613bfd84830161390d565b818152918301840191848101908b841115613c1757600080fd5b938501935b83851015613c415784359250613c31836136d2565b8282529385019390850190613c1c565b809750505050505050613c56606086016137f9565b905092959194509250565b600080600080600060808688031215613c7957600080fd5b8535613c84816136d2565b945060208601356001600160401b03811115613c9f57600080fd5b613cab888289016136e7565b9699909850959660408101359660609091013595509350505050565b60008060008060608587031215613cdd57600080fd5b8435613ce8816136d2565b935060208501356001600160401b0380821115613d0457600080fd5b613d10888389016136e7565b90955093506040870135915080821115613d2957600080fd5b50613d368782880161372b565b91505092959194509250565b600080600060808486031215613d5757600080fd5b8335613d62816136d2565b92506020840135613d72816136d2565b9150613d818560408601613854565b90509250925092565b634e487b7160e01b600052601160045260246000fd5b60008219821115613db357613db3613d8a565b500190565b634e487b7160e01b600052603260045260246000fd5b6020808252604d908201527f536c61736865722e71756575654d61676e6974756465436f6e63656e7472617460408201527f696f6e3a20746f6f206d616e792070656e64696e6720746f74616c206d61676e60608201526c6974756465207570646174657360981b608082015260a00190565b80356001600160401b038116811461380457600080fd5b600060208284031215613e6a57600080fd5b610f4682613e41565b60006001600160401b03808316818516808303821115613e9557613e95613d8a565b01949350505050565b6000600019821415613eb257613eb2613d8a565b5060010190565b600060208284031215613ecb57600080fd5b8151610f94816136d2565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b600060208284031215613f3257600080fd5b81518015158114610f9457600080fd5b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b600082821015613f9c57613f9c613d8a565b500390565b600081613fb057613fb0613d8a565b506000190190565b6000816000190483118215151615613fd257613fd2613d8a565b500290565b634e487b7160e01b600052601260045260246000fd5b600082613ffc57613ffc613fd7565b500490565b8881526001600160a01b0388811660208084019190915260c060408401819052830188905260009189919060e08501845b8b811015614059578435614045816136d2565b831682529383019390830190600101614032565b5085810360608701528881528201925088905060005b8881101561409b576001600160401b0361408883613e41565b168452928201929082019060010161406f565b50505060808301949094525060a001529695505050505050565b60006001600160401b03808316818516818304811182151516156140db576140db613d8a565b02949350505050565b6000606082840312156140f657600080fd5b604051606081018181106001600160401b0382111715614118576141186138f7565b6040528251614126816136d2565b81526020830151614136816136d2565b60208201526040830151614149816137e7565b60408201529392505050565b80356002811061380457600080fd5b634e487b7160e01b600052602160045260246000fd5b8035614185816136d2565b6001600160a01b03168252602081013561419e816137e7565b63ffffffff81166020840152505050565b81835260208301925060008160005b848110156141fd576141d0868361417a565b60406001600160401b036141e5828501613e41565b169087015260609586019591909101906001016141be565b5093949350505050565b600060a08201888352602060018060a01b03808a1682860152604060a0818701528389855260c08701905060c08a60051b88010194508a60005b8b8110156143115788870360bf190183528135368e9003605e1901811261426757600080fd5b8d0160608135614276816136d2565b87168952614285828901614155565b600281106142a357634e487b7160e01b600052602160045260246000fd5b898901528186013536839003601e190181126142be57600080fd5b820180356001600160401b038111156142d657600080fd5b82810236038413156142e757600080fd5b82888c01526142fb838c01828c85016141af565b9a50505093870193505090850190600101614241565b505050505060608401959095526080909201929092529695505050505050565b60008235605e1983360301811261434757600080fd5b9190910192915050565b60006020828403121561436357600080fd5b610f4682614155565b60006020828403121561437e57600080fd5b8135610f94816137e7565b60006001600160401b03838116908316818110156143a9576143a9613d8a565b039392505050565b6000808335601e198436030181126143c857600080fd5b8301803591506001600160401b038211156143e257600080fd5b6020019150368190038213156133ea57600080fd5b6001600160a01b03968716815294909516602085015263ffffffff9290921660408401526001600160401b039081166060840152908116608083015290911660a082015260c00190565b60006001600160401b038084168061445b5761445b613fd7565b92169190910492915050565b6000808335601e1984360301811261447e57600080fd5b8301803591506001600160401b0382111561449857600080fd5b60200191506060810236038213156133ea57600080fd5b60208082526032908201527f536c61736865722e5f7175657565416c6c6f636174696f6e3a20746f6f206d616040820152716e792070656e64696e67207570646174657360701b606082015260800190565b6001600160a01b0386811682528516602082015260c08101614526604083018661417a565b63ffffffff841660808301526001600160401b03831660a08301529695505050505050565b82815260006020604081840152835180604085015260005b8181101561457f57858101830151858201606001528201614563565b81811115614591576000606083870101525b50601f01601f191692909201606001949350505050565b6000602082840312156145ba57600080fd5b81516001600160e01b031981168114610f9457600080fdfea26469706673582212203704b7c5d9b59346371df9181e3035b103dc4722bdbe8aa9e0bb3461b41a524164736f6c634300080c0033",
}

// SlasherABI is the input ABI used to generate the binding from.
// Deprecated: Use SlasherMetaData.ABI instead.
var SlasherABI = SlasherMetaData.ABI

// SlasherBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SlasherMetaData.Bin instead.
var SlasherBin = SlasherMetaData.Bin

// DeploySlasher deploys a new Ethereum contract, binding an instance of Slasher to it.
func DeploySlasher(auth *bind.TransactOpts, backend bind.ContractBackend, _strategyManager common.Address, _delegationManager common.Address) (common.Address, *types.Transaction, *Slasher, error) {
	parsed, err := SlasherMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SlasherBin), backend, _strategyManager, _delegationManager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Slasher{SlasherCaller: SlasherCaller{contract: contract}, SlasherTransactor: SlasherTransactor{contract: contract}, SlasherFilterer: SlasherFilterer{contract: contract}}, nil
}

// Slasher is an auto generated Go binding around an Ethereum contract.
type Slasher struct {
	SlasherCaller     // Read-only binding to the contract
	SlasherTransactor // Write-only binding to the contract
	SlasherFilterer   // Log filterer for contract events
}

// SlasherCaller is an auto generated read-only Go binding around an Ethereum contract.
type SlasherCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SlasherTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SlasherTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SlasherFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SlasherFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SlasherSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SlasherSession struct {
	Contract     *Slasher          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SlasherCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SlasherCallerSession struct {
	Contract *SlasherCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// SlasherTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SlasherTransactorSession struct {
	Contract     *SlasherTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// SlasherRaw is an auto generated low-level Go binding around an Ethereum contract.
type SlasherRaw struct {
	Contract *Slasher // Generic contract binding to access the raw methods on
}

// SlasherCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SlasherCallerRaw struct {
	Contract *SlasherCaller // Generic read-only contract binding to access the raw methods on
}

// SlasherTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SlasherTransactorRaw struct {
	Contract *SlasherTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSlasher creates a new instance of Slasher, bound to a specific deployed contract.
func NewSlasher(address common.Address, backend bind.ContractBackend) (*Slasher, error) {
	contract, err := bindSlasher(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Slasher{SlasherCaller: SlasherCaller{contract: contract}, SlasherTransactor: SlasherTransactor{contract: contract}, SlasherFilterer: SlasherFilterer{contract: contract}}, nil
}

// NewSlasherCaller creates a new read-only instance of Slasher, bound to a specific deployed contract.
func NewSlasherCaller(address common.Address, caller bind.ContractCaller) (*SlasherCaller, error) {
	contract, err := bindSlasher(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SlasherCaller{contract: contract}, nil
}

// NewSlasherTransactor creates a new write-only instance of Slasher, bound to a specific deployed contract.
func NewSlasherTransactor(address common.Address, transactor bind.ContractTransactor) (*SlasherTransactor, error) {
	contract, err := bindSlasher(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SlasherTransactor{contract: contract}, nil
}

// NewSlasherFilterer creates a new log filterer instance of Slasher, bound to a specific deployed contract.
func NewSlasherFilterer(address common.Address, filterer bind.ContractFilterer) (*SlasherFilterer, error) {
	contract, err := bindSlasher(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SlasherFilterer{contract: contract}, nil
}

// bindSlasher binds a generic wrapper to an already deployed contract.
func bindSlasher(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SlasherMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Slasher *SlasherRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Slasher.Contract.SlasherCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Slasher *SlasherRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Slasher.Contract.SlasherTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Slasher *SlasherRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Slasher.Contract.SlasherTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Slasher *SlasherCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Slasher.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Slasher *SlasherTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Slasher.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Slasher *SlasherTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Slasher.Contract.contract.Transact(opts, method, params...)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_Slasher *SlasherCaller) DOMAINTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Slasher.contract.Call(opts, &out, "DOMAIN_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_Slasher *SlasherSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _Slasher.Contract.DOMAINTYPEHASH(&_Slasher.CallOpts)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_Slasher *SlasherCallerSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _Slasher.Contract.DOMAINTYPEHASH(&_Slasher.CallOpts)
}

// MAGNITUDECONCENTRATIONTYPEHASH is a free data retrieval call binding the contract method 0x8837866a.
//
// Solidity: function MAGNITUDE_CONCENTRATION_TYPEHASH() view returns(bytes32)
func (_Slasher *SlasherCaller) MAGNITUDECONCENTRATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Slasher.contract.Call(opts, &out, "MAGNITUDE_CONCENTRATION_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MAGNITUDECONCENTRATIONTYPEHASH is a free data retrieval call binding the contract method 0x8837866a.
//
// Solidity: function MAGNITUDE_CONCENTRATION_TYPEHASH() view returns(bytes32)
func (_Slasher *SlasherSession) MAGNITUDECONCENTRATIONTYPEHASH() ([32]byte, error) {
	return _Slasher.Contract.MAGNITUDECONCENTRATIONTYPEHASH(&_Slasher.CallOpts)
}

// MAGNITUDECONCENTRATIONTYPEHASH is a free data retrieval call binding the contract method 0x8837866a.
//
// Solidity: function MAGNITUDE_CONCENTRATION_TYPEHASH() view returns(bytes32)
func (_Slasher *SlasherCallerSession) MAGNITUDECONCENTRATIONTYPEHASH() ([32]byte, error) {
	return _Slasher.Contract.MAGNITUDECONCENTRATIONTYPEHASH(&_Slasher.CallOpts)
}

// MAGNITUDEDILUTIONTYPEHASH is a free data retrieval call binding the contract method 0xba728062.
//
// Solidity: function MAGNITUDE_DILUTION_TYPEHASH() view returns(bytes32)
func (_Slasher *SlasherCaller) MAGNITUDEDILUTIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Slasher.contract.Call(opts, &out, "MAGNITUDE_DILUTION_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MAGNITUDEDILUTIONTYPEHASH is a free data retrieval call binding the contract method 0xba728062.
//
// Solidity: function MAGNITUDE_DILUTION_TYPEHASH() view returns(bytes32)
func (_Slasher *SlasherSession) MAGNITUDEDILUTIONTYPEHASH() ([32]byte, error) {
	return _Slasher.Contract.MAGNITUDEDILUTIONTYPEHASH(&_Slasher.CallOpts)
}

// MAGNITUDEDILUTIONTYPEHASH is a free data retrieval call binding the contract method 0xba728062.
//
// Solidity: function MAGNITUDE_DILUTION_TYPEHASH() view returns(bytes32)
func (_Slasher *SlasherCallerSession) MAGNITUDEDILUTIONTYPEHASH() ([32]byte, error) {
	return _Slasher.Contract.MAGNITUDEDILUTIONTYPEHASH(&_Slasher.CallOpts)
}

// REALLOCATIONTYPEHASH is a free data retrieval call binding the contract method 0xa952bff9.
//
// Solidity: function REALLOCATION_TYPEHASH() view returns(bytes32)
func (_Slasher *SlasherCaller) REALLOCATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Slasher.contract.Call(opts, &out, "REALLOCATION_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REALLOCATIONTYPEHASH is a free data retrieval call binding the contract method 0xa952bff9.
//
// Solidity: function REALLOCATION_TYPEHASH() view returns(bytes32)
func (_Slasher *SlasherSession) REALLOCATIONTYPEHASH() ([32]byte, error) {
	return _Slasher.Contract.REALLOCATIONTYPEHASH(&_Slasher.CallOpts)
}

// REALLOCATIONTYPEHASH is a free data retrieval call binding the contract method 0xa952bff9.
//
// Solidity: function REALLOCATION_TYPEHASH() view returns(bytes32)
func (_Slasher *SlasherCallerSession) REALLOCATIONTYPEHASH() ([32]byte, error) {
	return _Slasher.Contract.REALLOCATIONTYPEHASH(&_Slasher.CallOpts)
}

// AllocatorSaltIsSpent is a free data retrieval call binding the contract method 0x46fb3a00.
//
// Solidity: function allocatorSaltIsSpent(address , bytes32 ) view returns(bool)
func (_Slasher *SlasherCaller) AllocatorSaltIsSpent(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (bool, error) {
	var out []interface{}
	err := _Slasher.contract.Call(opts, &out, "allocatorSaltIsSpent", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllocatorSaltIsSpent is a free data retrieval call binding the contract method 0x46fb3a00.
//
// Solidity: function allocatorSaltIsSpent(address , bytes32 ) view returns(bool)
func (_Slasher *SlasherSession) AllocatorSaltIsSpent(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _Slasher.Contract.AllocatorSaltIsSpent(&_Slasher.CallOpts, arg0, arg1)
}

// AllocatorSaltIsSpent is a free data retrieval call binding the contract method 0x46fb3a00.
//
// Solidity: function allocatorSaltIsSpent(address , bytes32 ) view returns(bool)
func (_Slasher *SlasherCallerSession) AllocatorSaltIsSpent(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _Slasher.Contract.AllocatorSaltIsSpent(&_Slasher.CallOpts, arg0, arg1)
}

// CalculateMagnitudeConcentrationDigestHash is a free data retrieval call binding the contract method 0x4b0284de.
//
// Solidity: function calculateMagnitudeConcentrationDigestHash(address operator, address[] strategies, uint64[] nonslashableToDecrement, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_Slasher *SlasherCaller) CalculateMagnitudeConcentrationDigestHash(opts *bind.CallOpts, operator common.Address, strategies []common.Address, nonslashableToDecrement []uint64, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Slasher.contract.Call(opts, &out, "calculateMagnitudeConcentrationDigestHash", operator, strategies, nonslashableToDecrement, salt, expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateMagnitudeConcentrationDigestHash is a free data retrieval call binding the contract method 0x4b0284de.
//
// Solidity: function calculateMagnitudeConcentrationDigestHash(address operator, address[] strategies, uint64[] nonslashableToDecrement, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_Slasher *SlasherSession) CalculateMagnitudeConcentrationDigestHash(operator common.Address, strategies []common.Address, nonslashableToDecrement []uint64, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _Slasher.Contract.CalculateMagnitudeConcentrationDigestHash(&_Slasher.CallOpts, operator, strategies, nonslashableToDecrement, salt, expiry)
}

// CalculateMagnitudeConcentrationDigestHash is a free data retrieval call binding the contract method 0x4b0284de.
//
// Solidity: function calculateMagnitudeConcentrationDigestHash(address operator, address[] strategies, uint64[] nonslashableToDecrement, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_Slasher *SlasherCallerSession) CalculateMagnitudeConcentrationDigestHash(operator common.Address, strategies []common.Address, nonslashableToDecrement []uint64, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _Slasher.Contract.CalculateMagnitudeConcentrationDigestHash(&_Slasher.CallOpts, operator, strategies, nonslashableToDecrement, salt, expiry)
}

// CalculateMagnitudeDilutionDigestHash is a free data retrieval call binding the contract method 0x3f5f49e9.
//
// Solidity: function calculateMagnitudeDilutionDigestHash(address operator, address[] strategies, uint64[] nonslashableToAdd, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_Slasher *SlasherCaller) CalculateMagnitudeDilutionDigestHash(opts *bind.CallOpts, operator common.Address, strategies []common.Address, nonslashableToAdd []uint64, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Slasher.contract.Call(opts, &out, "calculateMagnitudeDilutionDigestHash", operator, strategies, nonslashableToAdd, salt, expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateMagnitudeDilutionDigestHash is a free data retrieval call binding the contract method 0x3f5f49e9.
//
// Solidity: function calculateMagnitudeDilutionDigestHash(address operator, address[] strategies, uint64[] nonslashableToAdd, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_Slasher *SlasherSession) CalculateMagnitudeDilutionDigestHash(operator common.Address, strategies []common.Address, nonslashableToAdd []uint64, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _Slasher.Contract.CalculateMagnitudeDilutionDigestHash(&_Slasher.CallOpts, operator, strategies, nonslashableToAdd, salt, expiry)
}

// CalculateMagnitudeDilutionDigestHash is a free data retrieval call binding the contract method 0x3f5f49e9.
//
// Solidity: function calculateMagnitudeDilutionDigestHash(address operator, address[] strategies, uint64[] nonslashableToAdd, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_Slasher *SlasherCallerSession) CalculateMagnitudeDilutionDigestHash(operator common.Address, strategies []common.Address, nonslashableToAdd []uint64, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _Slasher.Contract.CalculateMagnitudeDilutionDigestHash(&_Slasher.CallOpts, operator, strategies, nonslashableToAdd, salt, expiry)
}

// CalculateReallocationDigestHash is a free data retrieval call binding the contract method 0x8d60eab5.
//
// Solidity: function calculateReallocationDigestHash(address operator, (address,uint8,((address,uint32),uint64)[])[] adjustmentParams, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_Slasher *SlasherCaller) CalculateReallocationDigestHash(opts *bind.CallOpts, operator common.Address, adjustmentParams []ISlasherMagnitudeAdjustmentsParam, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Slasher.contract.Call(opts, &out, "calculateReallocationDigestHash", operator, adjustmentParams, salt, expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateReallocationDigestHash is a free data retrieval call binding the contract method 0x8d60eab5.
//
// Solidity: function calculateReallocationDigestHash(address operator, (address,uint8,((address,uint32),uint64)[])[] adjustmentParams, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_Slasher *SlasherSession) CalculateReallocationDigestHash(operator common.Address, adjustmentParams []ISlasherMagnitudeAdjustmentsParam, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _Slasher.Contract.CalculateReallocationDigestHash(&_Slasher.CallOpts, operator, adjustmentParams, salt, expiry)
}

// CalculateReallocationDigestHash is a free data retrieval call binding the contract method 0x8d60eab5.
//
// Solidity: function calculateReallocationDigestHash(address operator, (address,uint8,((address,uint32),uint64)[])[] adjustmentParams, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_Slasher *SlasherCallerSession) CalculateReallocationDigestHash(operator common.Address, adjustmentParams []ISlasherMagnitudeAdjustmentsParam, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _Slasher.Contract.CalculateReallocationDigestHash(&_Slasher.CallOpts, operator, adjustmentParams, salt, expiry)
}

// CheckPendingMagnitudeUpdates is a free data retrieval call binding the contract method 0xd1920e37.
//
// Solidity: function checkPendingMagnitudeUpdates(address operator, address strategy, (address,uint32) operatorSet) view returns(bool)
func (_Slasher *SlasherCaller) CheckPendingMagnitudeUpdates(opts *bind.CallOpts, operator common.Address, strategy common.Address, operatorSet ISlasherOperatorSet) (bool, error) {
	var out []interface{}
	err := _Slasher.contract.Call(opts, &out, "checkPendingMagnitudeUpdates", operator, strategy, operatorSet)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckPendingMagnitudeUpdates is a free data retrieval call binding the contract method 0xd1920e37.
//
// Solidity: function checkPendingMagnitudeUpdates(address operator, address strategy, (address,uint32) operatorSet) view returns(bool)
func (_Slasher *SlasherSession) CheckPendingMagnitudeUpdates(operator common.Address, strategy common.Address, operatorSet ISlasherOperatorSet) (bool, error) {
	return _Slasher.Contract.CheckPendingMagnitudeUpdates(&_Slasher.CallOpts, operator, strategy, operatorSet)
}

// CheckPendingMagnitudeUpdates is a free data retrieval call binding the contract method 0xd1920e37.
//
// Solidity: function checkPendingMagnitudeUpdates(address operator, address strategy, (address,uint32) operatorSet) view returns(bool)
func (_Slasher *SlasherCallerSession) CheckPendingMagnitudeUpdates(operator common.Address, strategy common.Address, operatorSet ISlasherOperatorSet) (bool, error) {
	return _Slasher.Contract.CheckPendingMagnitudeUpdates(&_Slasher.CallOpts, operator, strategy, operatorSet)
}

// CheckPendingTotalMagnitudeUpdates is a free data retrieval call binding the contract method 0x64565d3d.
//
// Solidity: function checkPendingTotalMagnitudeUpdates(address operator, address strategy) view returns(bool)
func (_Slasher *SlasherCaller) CheckPendingTotalMagnitudeUpdates(opts *bind.CallOpts, operator common.Address, strategy common.Address) (bool, error) {
	var out []interface{}
	err := _Slasher.contract.Call(opts, &out, "checkPendingTotalMagnitudeUpdates", operator, strategy)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckPendingTotalMagnitudeUpdates is a free data retrieval call binding the contract method 0x64565d3d.
//
// Solidity: function checkPendingTotalMagnitudeUpdates(address operator, address strategy) view returns(bool)
func (_Slasher *SlasherSession) CheckPendingTotalMagnitudeUpdates(operator common.Address, strategy common.Address) (bool, error) {
	return _Slasher.Contract.CheckPendingTotalMagnitudeUpdates(&_Slasher.CallOpts, operator, strategy)
}

// CheckPendingTotalMagnitudeUpdates is a free data retrieval call binding the contract method 0x64565d3d.
//
// Solidity: function checkPendingTotalMagnitudeUpdates(address operator, address strategy) view returns(bool)
func (_Slasher *SlasherCallerSession) CheckPendingTotalMagnitudeUpdates(operator common.Address, strategy common.Address) (bool, error) {
	return _Slasher.Contract.CheckPendingTotalMagnitudeUpdates(&_Slasher.CallOpts, operator, strategy)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_Slasher *SlasherCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Slasher.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_Slasher *SlasherSession) Delegation() (common.Address, error) {
	return _Slasher.Contract.Delegation(&_Slasher.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_Slasher *SlasherCallerSession) Delegation() (common.Address, error) {
	return _Slasher.Contract.Delegation(&_Slasher.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_Slasher *SlasherCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Slasher.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_Slasher *SlasherSession) DomainSeparator() ([32]byte, error) {
	return _Slasher.Contract.DomainSeparator(&_Slasher.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_Slasher *SlasherCallerSession) DomainSeparator() ([32]byte, error) {
	return _Slasher.Contract.DomainSeparator(&_Slasher.CallOpts)
}

// GetAllocatorFor is a free data retrieval call binding the contract method 0x67e988de.
//
// Solidity: function getAllocatorFor(address operator) view returns(address)
func (_Slasher *SlasherCaller) GetAllocatorFor(opts *bind.CallOpts, operator common.Address) (common.Address, error) {
	var out []interface{}
	err := _Slasher.contract.Call(opts, &out, "getAllocatorFor", operator)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAllocatorFor is a free data retrieval call binding the contract method 0x67e988de.
//
// Solidity: function getAllocatorFor(address operator) view returns(address)
func (_Slasher *SlasherSession) GetAllocatorFor(operator common.Address) (common.Address, error) {
	return _Slasher.Contract.GetAllocatorFor(&_Slasher.CallOpts, operator)
}

// GetAllocatorFor is a free data retrieval call binding the contract method 0x67e988de.
//
// Solidity: function getAllocatorFor(address operator) view returns(address)
func (_Slasher *SlasherCallerSession) GetAllocatorFor(operator common.Address) (common.Address, error) {
	return _Slasher.Contract.GetAllocatorFor(&_Slasher.CallOpts, operator)
}

// GetMagnitudeUpdate is a free data retrieval call binding the contract method 0x16449dae.
//
// Solidity: function getMagnitudeUpdate(address operator, address strategy, (address,uint32) operatorSet, uint32 timestamp) view returns((uint32,uint64) magnitudeUpdate)
func (_Slasher *SlasherCaller) GetMagnitudeUpdate(opts *bind.CallOpts, operator common.Address, strategy common.Address, operatorSet ISlasherOperatorSet, timestamp uint32) (ISlasherMagnitudeUpdate, error) {
	var out []interface{}
	err := _Slasher.contract.Call(opts, &out, "getMagnitudeUpdate", operator, strategy, operatorSet, timestamp)

	if err != nil {
		return *new(ISlasherMagnitudeUpdate), err
	}

	out0 := *abi.ConvertType(out[0], new(ISlasherMagnitudeUpdate)).(*ISlasherMagnitudeUpdate)

	return out0, err

}

// GetMagnitudeUpdate is a free data retrieval call binding the contract method 0x16449dae.
//
// Solidity: function getMagnitudeUpdate(address operator, address strategy, (address,uint32) operatorSet, uint32 timestamp) view returns((uint32,uint64) magnitudeUpdate)
func (_Slasher *SlasherSession) GetMagnitudeUpdate(operator common.Address, strategy common.Address, operatorSet ISlasherOperatorSet, timestamp uint32) (ISlasherMagnitudeUpdate, error) {
	return _Slasher.Contract.GetMagnitudeUpdate(&_Slasher.CallOpts, operator, strategy, operatorSet, timestamp)
}

// GetMagnitudeUpdate is a free data retrieval call binding the contract method 0x16449dae.
//
// Solidity: function getMagnitudeUpdate(address operator, address strategy, (address,uint32) operatorSet, uint32 timestamp) view returns((uint32,uint64) magnitudeUpdate)
func (_Slasher *SlasherCallerSession) GetMagnitudeUpdate(operator common.Address, strategy common.Address, operatorSet ISlasherOperatorSet, timestamp uint32) (ISlasherMagnitudeUpdate, error) {
	return _Slasher.Contract.GetMagnitudeUpdate(&_Slasher.CallOpts, operator, strategy, operatorSet, timestamp)
}

// GetSlashableBips is a free data retrieval call binding the contract method 0x1c32f5c6.
//
// Solidity: function getSlashableBips(address operator, (address,uint32) operatorSet, address strategy) view returns(uint16 slashableBips)
func (_Slasher *SlasherCaller) GetSlashableBips(opts *bind.CallOpts, operator common.Address, operatorSet ISlasherOperatorSet, strategy common.Address) (uint16, error) {
	var out []interface{}
	err := _Slasher.contract.Call(opts, &out, "getSlashableBips", operator, operatorSet, strategy)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetSlashableBips is a free data retrieval call binding the contract method 0x1c32f5c6.
//
// Solidity: function getSlashableBips(address operator, (address,uint32) operatorSet, address strategy) view returns(uint16 slashableBips)
func (_Slasher *SlasherSession) GetSlashableBips(operator common.Address, operatorSet ISlasherOperatorSet, strategy common.Address) (uint16, error) {
	return _Slasher.Contract.GetSlashableBips(&_Slasher.CallOpts, operator, operatorSet, strategy)
}

// GetSlashableBips is a free data retrieval call binding the contract method 0x1c32f5c6.
//
// Solidity: function getSlashableBips(address operator, (address,uint32) operatorSet, address strategy) view returns(uint16 slashableBips)
func (_Slasher *SlasherCallerSession) GetSlashableBips(operator common.Address, operatorSet ISlasherOperatorSet, strategy common.Address) (uint16, error) {
	return _Slasher.Contract.GetSlashableBips(&_Slasher.CallOpts, operator, operatorSet, strategy)
}

// GetSlashedRate is a free data retrieval call binding the contract method 0x10c0a606.
//
// Solidity: function getSlashedRate(address operator, address strategy, (address,uint32) operatorSet, uint32 epoch) view returns(uint64)
func (_Slasher *SlasherCaller) GetSlashedRate(opts *bind.CallOpts, operator common.Address, strategy common.Address, operatorSet ISlasherOperatorSet, epoch uint32) (uint64, error) {
	var out []interface{}
	err := _Slasher.contract.Call(opts, &out, "getSlashedRate", operator, strategy, operatorSet, epoch)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetSlashedRate is a free data retrieval call binding the contract method 0x10c0a606.
//
// Solidity: function getSlashedRate(address operator, address strategy, (address,uint32) operatorSet, uint32 epoch) view returns(uint64)
func (_Slasher *SlasherSession) GetSlashedRate(operator common.Address, strategy common.Address, operatorSet ISlasherOperatorSet, epoch uint32) (uint64, error) {
	return _Slasher.Contract.GetSlashedRate(&_Slasher.CallOpts, operator, strategy, operatorSet, epoch)
}

// GetSlashedRate is a free data retrieval call binding the contract method 0x10c0a606.
//
// Solidity: function getSlashedRate(address operator, address strategy, (address,uint32) operatorSet, uint32 epoch) view returns(uint64)
func (_Slasher *SlasherCallerSession) GetSlashedRate(operator common.Address, strategy common.Address, operatorSet ISlasherOperatorSet, epoch uint32) (uint64, error) {
	return _Slasher.Contract.GetSlashedRate(&_Slasher.CallOpts, operator, strategy, operatorSet, epoch)
}

// GetTotalAndNonslashableUpdate is a free data retrieval call binding the contract method 0x364bce6f.
//
// Solidity: function getTotalAndNonslashableUpdate(address operator, address strategy, uint32 timestamp) view returns((uint32,uint64,uint64,uint64) totalAndNonslashableUpdate)
func (_Slasher *SlasherCaller) GetTotalAndNonslashableUpdate(opts *bind.CallOpts, operator common.Address, strategy common.Address, timestamp uint32) (ISlasherTotalAndNonslashableUpdate, error) {
	var out []interface{}
	err := _Slasher.contract.Call(opts, &out, "getTotalAndNonslashableUpdate", operator, strategy, timestamp)

	if err != nil {
		return *new(ISlasherTotalAndNonslashableUpdate), err
	}

	out0 := *abi.ConvertType(out[0], new(ISlasherTotalAndNonslashableUpdate)).(*ISlasherTotalAndNonslashableUpdate)

	return out0, err

}

// GetTotalAndNonslashableUpdate is a free data retrieval call binding the contract method 0x364bce6f.
//
// Solidity: function getTotalAndNonslashableUpdate(address operator, address strategy, uint32 timestamp) view returns((uint32,uint64,uint64,uint64) totalAndNonslashableUpdate)
func (_Slasher *SlasherSession) GetTotalAndNonslashableUpdate(operator common.Address, strategy common.Address, timestamp uint32) (ISlasherTotalAndNonslashableUpdate, error) {
	return _Slasher.Contract.GetTotalAndNonslashableUpdate(&_Slasher.CallOpts, operator, strategy, timestamp)
}

// GetTotalAndNonslashableUpdate is a free data retrieval call binding the contract method 0x364bce6f.
//
// Solidity: function getTotalAndNonslashableUpdate(address operator, address strategy, uint32 timestamp) view returns((uint32,uint64,uint64,uint64) totalAndNonslashableUpdate)
func (_Slasher *SlasherCallerSession) GetTotalAndNonslashableUpdate(operator common.Address, strategy common.Address, timestamp uint32) (ISlasherTotalAndNonslashableUpdate, error) {
	return _Slasher.Contract.GetTotalAndNonslashableUpdate(&_Slasher.CallOpts, operator, strategy, timestamp)
}

// LastSlashed is a free data retrieval call binding the contract method 0x5ab112d6.
//
// Solidity: function lastSlashed(address operator, address strategy) view returns(uint32)
func (_Slasher *SlasherCaller) LastSlashed(opts *bind.CallOpts, operator common.Address, strategy common.Address) (uint32, error) {
	var out []interface{}
	err := _Slasher.contract.Call(opts, &out, "lastSlashed", operator, strategy)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LastSlashed is a free data retrieval call binding the contract method 0x5ab112d6.
//
// Solidity: function lastSlashed(address operator, address strategy) view returns(uint32)
func (_Slasher *SlasherSession) LastSlashed(operator common.Address, strategy common.Address) (uint32, error) {
	return _Slasher.Contract.LastSlashed(&_Slasher.CallOpts, operator, strategy)
}

// LastSlashed is a free data retrieval call binding the contract method 0x5ab112d6.
//
// Solidity: function lastSlashed(address operator, address strategy) view returns(uint32)
func (_Slasher *SlasherCallerSession) LastSlashed(operator common.Address, strategy common.Address) (uint32, error) {
	return _Slasher.Contract.LastSlashed(&_Slasher.CallOpts, operator, strategy)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Slasher *SlasherCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Slasher.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Slasher *SlasherSession) Owner() (common.Address, error) {
	return _Slasher.Contract.Owner(&_Slasher.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Slasher *SlasherCallerSession) Owner() (common.Address, error) {
	return _Slasher.Contract.Owner(&_Slasher.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_Slasher *SlasherCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _Slasher.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_Slasher *SlasherSession) Paused(index uint8) (bool, error) {
	return _Slasher.Contract.Paused(&_Slasher.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_Slasher *SlasherCallerSession) Paused(index uint8) (bool, error) {
	return _Slasher.Contract.Paused(&_Slasher.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_Slasher *SlasherCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Slasher.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_Slasher *SlasherSession) Paused0() (*big.Int, error) {
	return _Slasher.Contract.Paused0(&_Slasher.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_Slasher *SlasherCallerSession) Paused0() (*big.Int, error) {
	return _Slasher.Contract.Paused0(&_Slasher.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_Slasher *SlasherCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Slasher.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_Slasher *SlasherSession) PauserRegistry() (common.Address, error) {
	return _Slasher.Contract.PauserRegistry(&_Slasher.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_Slasher *SlasherCallerSession) PauserRegistry() (common.Address, error) {
	return _Slasher.Contract.PauserRegistry(&_Slasher.CallOpts)
}

// ShareScalingFactor is a free data retrieval call binding the contract method 0x334f00d6.
//
// Solidity: function shareScalingFactor(address operator, address strategy) view returns(uint64)
func (_Slasher *SlasherCaller) ShareScalingFactor(opts *bind.CallOpts, operator common.Address, strategy common.Address) (uint64, error) {
	var out []interface{}
	err := _Slasher.contract.Call(opts, &out, "shareScalingFactor", operator, strategy)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ShareScalingFactor is a free data retrieval call binding the contract method 0x334f00d6.
//
// Solidity: function shareScalingFactor(address operator, address strategy) view returns(uint64)
func (_Slasher *SlasherSession) ShareScalingFactor(operator common.Address, strategy common.Address) (uint64, error) {
	return _Slasher.Contract.ShareScalingFactor(&_Slasher.CallOpts, operator, strategy)
}

// ShareScalingFactor is a free data retrieval call binding the contract method 0x334f00d6.
//
// Solidity: function shareScalingFactor(address operator, address strategy) view returns(uint64)
func (_Slasher *SlasherCallerSession) ShareScalingFactor(operator common.Address, strategy common.Address) (uint64, error) {
	return _Slasher.Contract.ShareScalingFactor(&_Slasher.CallOpts, operator, strategy)
}

// ShareScalingFactorAtEpoch is a free data retrieval call binding the contract method 0xe49a1e84.
//
// Solidity: function shareScalingFactorAtEpoch(address operator, address strategy, uint32 epoch) view returns(uint64)
func (_Slasher *SlasherCaller) ShareScalingFactorAtEpoch(opts *bind.CallOpts, operator common.Address, strategy common.Address, epoch uint32) (uint64, error) {
	var out []interface{}
	err := _Slasher.contract.Call(opts, &out, "shareScalingFactorAtEpoch", operator, strategy, epoch)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ShareScalingFactorAtEpoch is a free data retrieval call binding the contract method 0xe49a1e84.
//
// Solidity: function shareScalingFactorAtEpoch(address operator, address strategy, uint32 epoch) view returns(uint64)
func (_Slasher *SlasherSession) ShareScalingFactorAtEpoch(operator common.Address, strategy common.Address, epoch uint32) (uint64, error) {
	return _Slasher.Contract.ShareScalingFactorAtEpoch(&_Slasher.CallOpts, operator, strategy, epoch)
}

// ShareScalingFactorAtEpoch is a free data retrieval call binding the contract method 0xe49a1e84.
//
// Solidity: function shareScalingFactorAtEpoch(address operator, address strategy, uint32 epoch) view returns(uint64)
func (_Slasher *SlasherCallerSession) ShareScalingFactorAtEpoch(operator common.Address, strategy common.Address, epoch uint32) (uint64, error) {
	return _Slasher.Contract.ShareScalingFactorAtEpoch(&_Slasher.CallOpts, operator, strategy, epoch)
}

// SlashingEpochHistory is a free data retrieval call binding the contract method 0x4279a7e6.
//
// Solidity: function slashingEpochHistory(address , address , uint256 ) view returns(uint32)
func (_Slasher *SlasherCaller) SlashingEpochHistory(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) (uint32, error) {
	var out []interface{}
	err := _Slasher.contract.Call(opts, &out, "slashingEpochHistory", arg0, arg1, arg2)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// SlashingEpochHistory is a free data retrieval call binding the contract method 0x4279a7e6.
//
// Solidity: function slashingEpochHistory(address , address , uint256 ) view returns(uint32)
func (_Slasher *SlasherSession) SlashingEpochHistory(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (uint32, error) {
	return _Slasher.Contract.SlashingEpochHistory(&_Slasher.CallOpts, arg0, arg1, arg2)
}

// SlashingEpochHistory is a free data retrieval call binding the contract method 0x4279a7e6.
//
// Solidity: function slashingEpochHistory(address , address , uint256 ) view returns(uint32)
func (_Slasher *SlasherCallerSession) SlashingEpochHistory(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (uint32, error) {
	return _Slasher.Contract.SlashingEpochHistory(&_Slasher.CallOpts, arg0, arg1, arg2)
}

// SlashingUpdates is a free data retrieval call binding the contract method 0x0b1b781e.
//
// Solidity: function slashingUpdates(address , address , uint32 ) view returns(uint64 slashingRate, uint64 scalingFactor)
func (_Slasher *SlasherCaller) SlashingUpdates(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 uint32) (struct {
	SlashingRate  uint64
	ScalingFactor uint64
}, error) {
	var out []interface{}
	err := _Slasher.contract.Call(opts, &out, "slashingUpdates", arg0, arg1, arg2)

	outstruct := new(struct {
		SlashingRate  uint64
		ScalingFactor uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SlashingRate = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.ScalingFactor = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// SlashingUpdates is a free data retrieval call binding the contract method 0x0b1b781e.
//
// Solidity: function slashingUpdates(address , address , uint32 ) view returns(uint64 slashingRate, uint64 scalingFactor)
func (_Slasher *SlasherSession) SlashingUpdates(arg0 common.Address, arg1 common.Address, arg2 uint32) (struct {
	SlashingRate  uint64
	ScalingFactor uint64
}, error) {
	return _Slasher.Contract.SlashingUpdates(&_Slasher.CallOpts, arg0, arg1, arg2)
}

// SlashingUpdates is a free data retrieval call binding the contract method 0x0b1b781e.
//
// Solidity: function slashingUpdates(address , address , uint32 ) view returns(uint64 slashingRate, uint64 scalingFactor)
func (_Slasher *SlasherCallerSession) SlashingUpdates(arg0 common.Address, arg1 common.Address, arg2 uint32) (struct {
	SlashingRate  uint64
	ScalingFactor uint64
}, error) {
	return _Slasher.Contract.SlashingUpdates(&_Slasher.CallOpts, arg0, arg1, arg2)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_Slasher *SlasherCaller) StrategyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Slasher.contract.Call(opts, &out, "strategyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_Slasher *SlasherSession) StrategyManager() (common.Address, error) {
	return _Slasher.Contract.StrategyManager(&_Slasher.CallOpts)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_Slasher *SlasherCallerSession) StrategyManager() (common.Address, error) {
	return _Slasher.Contract.StrategyManager(&_Slasher.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address initialOwner, address _pauserRegistry, uint256 initialPausedStatus) returns()
func (_Slasher *SlasherTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, _pauserRegistry common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _Slasher.contract.Transact(opts, "initialize", initialOwner, _pauserRegistry, initialPausedStatus)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address initialOwner, address _pauserRegistry, uint256 initialPausedStatus) returns()
func (_Slasher *SlasherSession) Initialize(initialOwner common.Address, _pauserRegistry common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _Slasher.Contract.Initialize(&_Slasher.TransactOpts, initialOwner, _pauserRegistry, initialPausedStatus)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address initialOwner, address _pauserRegistry, uint256 initialPausedStatus) returns()
func (_Slasher *SlasherTransactorSession) Initialize(initialOwner common.Address, _pauserRegistry common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _Slasher.Contract.Initialize(&_Slasher.TransactOpts, initialOwner, _pauserRegistry, initialPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_Slasher *SlasherTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _Slasher.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_Slasher *SlasherSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _Slasher.Contract.Pause(&_Slasher.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_Slasher *SlasherTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _Slasher.Contract.Pause(&_Slasher.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_Slasher *SlasherTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Slasher.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_Slasher *SlasherSession) PauseAll() (*types.Transaction, error) {
	return _Slasher.Contract.PauseAll(&_Slasher.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_Slasher *SlasherTransactorSession) PauseAll() (*types.Transaction, error) {
	return _Slasher.Contract.PauseAll(&_Slasher.TransactOpts)
}

// QueueMagnitudeConcentration is a paid mutator transaction binding the contract method 0xe565c621.
//
// Solidity: function queueMagnitudeConcentration(address operator, address[] strategies, uint64[] nonslashableToDecrement, (bytes,bytes32,uint256) allocatorSignature) returns(uint64 newNonslashableMagnitude, uint64 newTotalMagnitude)
func (_Slasher *SlasherTransactor) QueueMagnitudeConcentration(opts *bind.TransactOpts, operator common.Address, strategies []common.Address, nonslashableToDecrement []uint64, allocatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _Slasher.contract.Transact(opts, "queueMagnitudeConcentration", operator, strategies, nonslashableToDecrement, allocatorSignature)
}

// QueueMagnitudeConcentration is a paid mutator transaction binding the contract method 0xe565c621.
//
// Solidity: function queueMagnitudeConcentration(address operator, address[] strategies, uint64[] nonslashableToDecrement, (bytes,bytes32,uint256) allocatorSignature) returns(uint64 newNonslashableMagnitude, uint64 newTotalMagnitude)
func (_Slasher *SlasherSession) QueueMagnitudeConcentration(operator common.Address, strategies []common.Address, nonslashableToDecrement []uint64, allocatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _Slasher.Contract.QueueMagnitudeConcentration(&_Slasher.TransactOpts, operator, strategies, nonslashableToDecrement, allocatorSignature)
}

// QueueMagnitudeConcentration is a paid mutator transaction binding the contract method 0xe565c621.
//
// Solidity: function queueMagnitudeConcentration(address operator, address[] strategies, uint64[] nonslashableToDecrement, (bytes,bytes32,uint256) allocatorSignature) returns(uint64 newNonslashableMagnitude, uint64 newTotalMagnitude)
func (_Slasher *SlasherTransactorSession) QueueMagnitudeConcentration(operator common.Address, strategies []common.Address, nonslashableToDecrement []uint64, allocatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _Slasher.Contract.QueueMagnitudeConcentration(&_Slasher.TransactOpts, operator, strategies, nonslashableToDecrement, allocatorSignature)
}

// QueueMagnitudeDilution is a paid mutator transaction binding the contract method 0x017ec6f9.
//
// Solidity: function queueMagnitudeDilution(address operator, address[] strategies, uint64[] nonslashableToAdd, (bytes,bytes32,uint256) allocatorSignature) returns(uint64 newNonslashableMagnitude, uint64 newTotalMagnitude)
func (_Slasher *SlasherTransactor) QueueMagnitudeDilution(opts *bind.TransactOpts, operator common.Address, strategies []common.Address, nonslashableToAdd []uint64, allocatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _Slasher.contract.Transact(opts, "queueMagnitudeDilution", operator, strategies, nonslashableToAdd, allocatorSignature)
}

// QueueMagnitudeDilution is a paid mutator transaction binding the contract method 0x017ec6f9.
//
// Solidity: function queueMagnitudeDilution(address operator, address[] strategies, uint64[] nonslashableToAdd, (bytes,bytes32,uint256) allocatorSignature) returns(uint64 newNonslashableMagnitude, uint64 newTotalMagnitude)
func (_Slasher *SlasherSession) QueueMagnitudeDilution(operator common.Address, strategies []common.Address, nonslashableToAdd []uint64, allocatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _Slasher.Contract.QueueMagnitudeDilution(&_Slasher.TransactOpts, operator, strategies, nonslashableToAdd, allocatorSignature)
}

// QueueMagnitudeDilution is a paid mutator transaction binding the contract method 0x017ec6f9.
//
// Solidity: function queueMagnitudeDilution(address operator, address[] strategies, uint64[] nonslashableToAdd, (bytes,bytes32,uint256) allocatorSignature) returns(uint64 newNonslashableMagnitude, uint64 newTotalMagnitude)
func (_Slasher *SlasherTransactorSession) QueueMagnitudeDilution(operator common.Address, strategies []common.Address, nonslashableToAdd []uint64, allocatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _Slasher.Contract.QueueMagnitudeDilution(&_Slasher.TransactOpts, operator, strategies, nonslashableToAdd, allocatorSignature)
}

// QueueReallocation is a paid mutator transaction binding the contract method 0x9292a40d.
//
// Solidity: function queueReallocation(address operator, (address,uint8,((address,uint32),uint64)[])[] adjustmentParams, (bytes,bytes32,uint256) allocatorSignature) returns(uint32 effectTimestamp)
func (_Slasher *SlasherTransactor) QueueReallocation(opts *bind.TransactOpts, operator common.Address, adjustmentParams []ISlasherMagnitudeAdjustmentsParam, allocatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _Slasher.contract.Transact(opts, "queueReallocation", operator, adjustmentParams, allocatorSignature)
}

// QueueReallocation is a paid mutator transaction binding the contract method 0x9292a40d.
//
// Solidity: function queueReallocation(address operator, (address,uint8,((address,uint32),uint64)[])[] adjustmentParams, (bytes,bytes32,uint256) allocatorSignature) returns(uint32 effectTimestamp)
func (_Slasher *SlasherSession) QueueReallocation(operator common.Address, adjustmentParams []ISlasherMagnitudeAdjustmentsParam, allocatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _Slasher.Contract.QueueReallocation(&_Slasher.TransactOpts, operator, adjustmentParams, allocatorSignature)
}

// QueueReallocation is a paid mutator transaction binding the contract method 0x9292a40d.
//
// Solidity: function queueReallocation(address operator, (address,uint8,((address,uint32),uint64)[])[] adjustmentParams, (bytes,bytes32,uint256) allocatorSignature) returns(uint32 effectTimestamp)
func (_Slasher *SlasherTransactorSession) QueueReallocation(operator common.Address, adjustmentParams []ISlasherMagnitudeAdjustmentsParam, allocatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _Slasher.Contract.QueueReallocation(&_Slasher.TransactOpts, operator, adjustmentParams, allocatorSignature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Slasher *SlasherTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Slasher.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Slasher *SlasherSession) RenounceOwnership() (*types.Transaction, error) {
	return _Slasher.Contract.RenounceOwnership(&_Slasher.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Slasher *SlasherTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Slasher.Contract.RenounceOwnership(&_Slasher.TransactOpts)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_Slasher *SlasherTransactor) SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error) {
	return _Slasher.contract.Transact(opts, "setPauserRegistry", newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_Slasher *SlasherSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _Slasher.Contract.SetPauserRegistry(&_Slasher.TransactOpts, newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_Slasher *SlasherTransactorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _Slasher.Contract.SetPauserRegistry(&_Slasher.TransactOpts, newPauserRegistry)
}

// SlashOperator is a paid mutator transaction binding the contract method 0x5c120057.
//
// Solidity: function slashOperator(address operator, uint32 operatorSetId, address[] strategies, uint32 bipsToSlash) returns()
func (_Slasher *SlasherTransactor) SlashOperator(opts *bind.TransactOpts, operator common.Address, operatorSetId uint32, strategies []common.Address, bipsToSlash uint32) (*types.Transaction, error) {
	return _Slasher.contract.Transact(opts, "slashOperator", operator, operatorSetId, strategies, bipsToSlash)
}

// SlashOperator is a paid mutator transaction binding the contract method 0x5c120057.
//
// Solidity: function slashOperator(address operator, uint32 operatorSetId, address[] strategies, uint32 bipsToSlash) returns()
func (_Slasher *SlasherSession) SlashOperator(operator common.Address, operatorSetId uint32, strategies []common.Address, bipsToSlash uint32) (*types.Transaction, error) {
	return _Slasher.Contract.SlashOperator(&_Slasher.TransactOpts, operator, operatorSetId, strategies, bipsToSlash)
}

// SlashOperator is a paid mutator transaction binding the contract method 0x5c120057.
//
// Solidity: function slashOperator(address operator, uint32 operatorSetId, address[] strategies, uint32 bipsToSlash) returns()
func (_Slasher *SlasherTransactorSession) SlashOperator(operator common.Address, operatorSetId uint32, strategies []common.Address, bipsToSlash uint32) (*types.Transaction, error) {
	return _Slasher.Contract.SlashOperator(&_Slasher.TransactOpts, operator, operatorSetId, strategies, bipsToSlash)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Slasher *SlasherTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Slasher.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Slasher *SlasherSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Slasher.Contract.TransferOwnership(&_Slasher.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Slasher *SlasherTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Slasher.Contract.TransferOwnership(&_Slasher.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_Slasher *SlasherTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _Slasher.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_Slasher *SlasherSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _Slasher.Contract.Unpause(&_Slasher.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_Slasher *SlasherTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _Slasher.Contract.Unpause(&_Slasher.TransactOpts, newPausedStatus)
}

// SlasherInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Slasher contract.
type SlasherInitializedIterator struct {
	Event *SlasherInitialized // Event containing the contract specifics and raw log

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
func (it *SlasherInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlasherInitialized)
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
		it.Event = new(SlasherInitialized)
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
func (it *SlasherInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlasherInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlasherInitialized represents a Initialized event raised by the Slasher contract.
type SlasherInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Slasher *SlasherFilterer) FilterInitialized(opts *bind.FilterOpts) (*SlasherInitializedIterator, error) {

	logs, sub, err := _Slasher.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SlasherInitializedIterator{contract: _Slasher.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Slasher *SlasherFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SlasherInitialized) (event.Subscription, error) {

	logs, sub, err := _Slasher.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlasherInitialized)
				if err := _Slasher.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Slasher *SlasherFilterer) ParseInitialized(log types.Log) (*SlasherInitialized, error) {
	event := new(SlasherInitialized)
	if err := _Slasher.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlasherMagnitudeDecrementedIterator is returned from FilterMagnitudeDecremented and is used to iterate over the raw logs and unpacked data for MagnitudeDecremented events raised by the Slasher contract.
type SlasherMagnitudeDecrementedIterator struct {
	Event *SlasherMagnitudeDecremented // Event containing the contract specifics and raw log

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
func (it *SlasherMagnitudeDecrementedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlasherMagnitudeDecremented)
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
		it.Event = new(SlasherMagnitudeDecremented)
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
func (it *SlasherMagnitudeDecrementedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlasherMagnitudeDecrementedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlasherMagnitudeDecremented represents a MagnitudeDecremented event raised by the Slasher contract.
type SlasherMagnitudeDecremented struct {
	Operator                  common.Address
	OperatorSet               ISlasherOperatorSet
	Strategy                  common.Address
	UpdatedSlashableMagnitude uint64
	EffectTimestamp           uint32
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterMagnitudeDecremented is a free log retrieval operation binding the contract event 0x543a5d2a6758e26af80ad687683e71136dbf9d41ba11bea98120b3a5be24652d.
//
// Solidity: event MagnitudeDecremented(address operator, (address,uint32) operatorSet, address strategy, uint64 updatedSlashableMagnitude, uint32 effectTimestamp)
func (_Slasher *SlasherFilterer) FilterMagnitudeDecremented(opts *bind.FilterOpts) (*SlasherMagnitudeDecrementedIterator, error) {

	logs, sub, err := _Slasher.contract.FilterLogs(opts, "MagnitudeDecremented")
	if err != nil {
		return nil, err
	}
	return &SlasherMagnitudeDecrementedIterator{contract: _Slasher.contract, event: "MagnitudeDecremented", logs: logs, sub: sub}, nil
}

// WatchMagnitudeDecremented is a free log subscription operation binding the contract event 0x543a5d2a6758e26af80ad687683e71136dbf9d41ba11bea98120b3a5be24652d.
//
// Solidity: event MagnitudeDecremented(address operator, (address,uint32) operatorSet, address strategy, uint64 updatedSlashableMagnitude, uint32 effectTimestamp)
func (_Slasher *SlasherFilterer) WatchMagnitudeDecremented(opts *bind.WatchOpts, sink chan<- *SlasherMagnitudeDecremented) (event.Subscription, error) {

	logs, sub, err := _Slasher.contract.WatchLogs(opts, "MagnitudeDecremented")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlasherMagnitudeDecremented)
				if err := _Slasher.contract.UnpackLog(event, "MagnitudeDecremented", log); err != nil {
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

// ParseMagnitudeDecremented is a log parse operation binding the contract event 0x543a5d2a6758e26af80ad687683e71136dbf9d41ba11bea98120b3a5be24652d.
//
// Solidity: event MagnitudeDecremented(address operator, (address,uint32) operatorSet, address strategy, uint64 updatedSlashableMagnitude, uint32 effectTimestamp)
func (_Slasher *SlasherFilterer) ParseMagnitudeDecremented(log types.Log) (*SlasherMagnitudeDecremented, error) {
	event := new(SlasherMagnitudeDecremented)
	if err := _Slasher.contract.UnpackLog(event, "MagnitudeDecremented", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlasherMagnitudeUpdatedIterator is returned from FilterMagnitudeUpdated and is used to iterate over the raw logs and unpacked data for MagnitudeUpdated events raised by the Slasher contract.
type SlasherMagnitudeUpdatedIterator struct {
	Event *SlasherMagnitudeUpdated // Event containing the contract specifics and raw log

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
func (it *SlasherMagnitudeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlasherMagnitudeUpdated)
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
		it.Event = new(SlasherMagnitudeUpdated)
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
func (it *SlasherMagnitudeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlasherMagnitudeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlasherMagnitudeUpdated represents a MagnitudeUpdated event raised by the Slasher contract.
type SlasherMagnitudeUpdated struct {
	Operator           common.Address
	Strategy           common.Address
	OperatorSet        ISlasherOperatorSet
	EffectTimestamp    uint32
	SlashableMagnitude uint64
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMagnitudeUpdated is a free log retrieval operation binding the contract event 0x770e5720664253f46b99dcb3b922235f8ee78191bdb00bf7efb8c48ba3ab396d.
//
// Solidity: event MagnitudeUpdated(address operator, address strategy, (address,uint32) operatorSet, uint32 effectTimestamp, uint64 slashableMagnitude)
func (_Slasher *SlasherFilterer) FilterMagnitudeUpdated(opts *bind.FilterOpts) (*SlasherMagnitudeUpdatedIterator, error) {

	logs, sub, err := _Slasher.contract.FilterLogs(opts, "MagnitudeUpdated")
	if err != nil {
		return nil, err
	}
	return &SlasherMagnitudeUpdatedIterator{contract: _Slasher.contract, event: "MagnitudeUpdated", logs: logs, sub: sub}, nil
}

// WatchMagnitudeUpdated is a free log subscription operation binding the contract event 0x770e5720664253f46b99dcb3b922235f8ee78191bdb00bf7efb8c48ba3ab396d.
//
// Solidity: event MagnitudeUpdated(address operator, address strategy, (address,uint32) operatorSet, uint32 effectTimestamp, uint64 slashableMagnitude)
func (_Slasher *SlasherFilterer) WatchMagnitudeUpdated(opts *bind.WatchOpts, sink chan<- *SlasherMagnitudeUpdated) (event.Subscription, error) {

	logs, sub, err := _Slasher.contract.WatchLogs(opts, "MagnitudeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlasherMagnitudeUpdated)
				if err := _Slasher.contract.UnpackLog(event, "MagnitudeUpdated", log); err != nil {
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

// ParseMagnitudeUpdated is a log parse operation binding the contract event 0x770e5720664253f46b99dcb3b922235f8ee78191bdb00bf7efb8c48ba3ab396d.
//
// Solidity: event MagnitudeUpdated(address operator, address strategy, (address,uint32) operatorSet, uint32 effectTimestamp, uint64 slashableMagnitude)
func (_Slasher *SlasherFilterer) ParseMagnitudeUpdated(log types.Log) (*SlasherMagnitudeUpdated, error) {
	event := new(SlasherMagnitudeUpdated)
	if err := _Slasher.contract.UnpackLog(event, "MagnitudeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlasherNonslashableMagnitudeDecrementedIterator is returned from FilterNonslashableMagnitudeDecremented and is used to iterate over the raw logs and unpacked data for NonslashableMagnitudeDecremented events raised by the Slasher contract.
type SlasherNonslashableMagnitudeDecrementedIterator struct {
	Event *SlasherNonslashableMagnitudeDecremented // Event containing the contract specifics and raw log

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
func (it *SlasherNonslashableMagnitudeDecrementedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlasherNonslashableMagnitudeDecremented)
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
		it.Event = new(SlasherNonslashableMagnitudeDecremented)
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
func (it *SlasherNonslashableMagnitudeDecrementedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlasherNonslashableMagnitudeDecrementedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlasherNonslashableMagnitudeDecremented represents a NonslashableMagnitudeDecremented event raised by the Slasher contract.
type SlasherNonslashableMagnitudeDecremented struct {
	Operator                     common.Address
	Strategy                     common.Address
	UpdatedNonslashableMagnitude uint64
	EffectTimestamp              uint32
	Raw                          types.Log // Blockchain specific contextual infos
}

// FilterNonslashableMagnitudeDecremented is a free log retrieval operation binding the contract event 0x96111829d0d3e9dd48c4e65cecac2a6ce99675f75740a2cbef2a99dca965a713.
//
// Solidity: event NonslashableMagnitudeDecremented(address operator, address strategy, uint64 updatedNonslashableMagnitude, uint32 effectTimestamp)
func (_Slasher *SlasherFilterer) FilterNonslashableMagnitudeDecremented(opts *bind.FilterOpts) (*SlasherNonslashableMagnitudeDecrementedIterator, error) {

	logs, sub, err := _Slasher.contract.FilterLogs(opts, "NonslashableMagnitudeDecremented")
	if err != nil {
		return nil, err
	}
	return &SlasherNonslashableMagnitudeDecrementedIterator{contract: _Slasher.contract, event: "NonslashableMagnitudeDecremented", logs: logs, sub: sub}, nil
}

// WatchNonslashableMagnitudeDecremented is a free log subscription operation binding the contract event 0x96111829d0d3e9dd48c4e65cecac2a6ce99675f75740a2cbef2a99dca965a713.
//
// Solidity: event NonslashableMagnitudeDecremented(address operator, address strategy, uint64 updatedNonslashableMagnitude, uint32 effectTimestamp)
func (_Slasher *SlasherFilterer) WatchNonslashableMagnitudeDecremented(opts *bind.WatchOpts, sink chan<- *SlasherNonslashableMagnitudeDecremented) (event.Subscription, error) {

	logs, sub, err := _Slasher.contract.WatchLogs(opts, "NonslashableMagnitudeDecremented")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlasherNonslashableMagnitudeDecremented)
				if err := _Slasher.contract.UnpackLog(event, "NonslashableMagnitudeDecremented", log); err != nil {
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

// ParseNonslashableMagnitudeDecremented is a log parse operation binding the contract event 0x96111829d0d3e9dd48c4e65cecac2a6ce99675f75740a2cbef2a99dca965a713.
//
// Solidity: event NonslashableMagnitudeDecremented(address operator, address strategy, uint64 updatedNonslashableMagnitude, uint32 effectTimestamp)
func (_Slasher *SlasherFilterer) ParseNonslashableMagnitudeDecremented(log types.Log) (*SlasherNonslashableMagnitudeDecremented, error) {
	event := new(SlasherNonslashableMagnitudeDecremented)
	if err := _Slasher.contract.UnpackLog(event, "NonslashableMagnitudeDecremented", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlasherOperatorSlashedIterator is returned from FilterOperatorSlashed and is used to iterate over the raw logs and unpacked data for OperatorSlashed events raised by the Slasher contract.
type SlasherOperatorSlashedIterator struct {
	Event *SlasherOperatorSlashed // Event containing the contract specifics and raw log

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
func (it *SlasherOperatorSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlasherOperatorSlashed)
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
		it.Event = new(SlasherOperatorSlashed)
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
func (it *SlasherOperatorSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlasherOperatorSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlasherOperatorSlashed represents a OperatorSlashed event raised by the Slasher contract.
type SlasherOperatorSlashed struct {
	Operator      common.Address
	OperatorSet   ISlasherOperatorSet
	BipsToSlash   uint32
	Strategies    []common.Address
	SlashingRates []uint64
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOperatorSlashed is a free log retrieval operation binding the contract event 0x715df5b5f92bd7692569b1129f8a601730c048c3daecf410b60b52d8a6015d40.
//
// Solidity: event OperatorSlashed(address operator, (address,uint32) operatorSet, uint32 bipsToSlash, address[] strategies, uint64[] slashingRates)
func (_Slasher *SlasherFilterer) FilterOperatorSlashed(opts *bind.FilterOpts) (*SlasherOperatorSlashedIterator, error) {

	logs, sub, err := _Slasher.contract.FilterLogs(opts, "OperatorSlashed")
	if err != nil {
		return nil, err
	}
	return &SlasherOperatorSlashedIterator{contract: _Slasher.contract, event: "OperatorSlashed", logs: logs, sub: sub}, nil
}

// WatchOperatorSlashed is a free log subscription operation binding the contract event 0x715df5b5f92bd7692569b1129f8a601730c048c3daecf410b60b52d8a6015d40.
//
// Solidity: event OperatorSlashed(address operator, (address,uint32) operatorSet, uint32 bipsToSlash, address[] strategies, uint64[] slashingRates)
func (_Slasher *SlasherFilterer) WatchOperatorSlashed(opts *bind.WatchOpts, sink chan<- *SlasherOperatorSlashed) (event.Subscription, error) {

	logs, sub, err := _Slasher.contract.WatchLogs(opts, "OperatorSlashed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlasherOperatorSlashed)
				if err := _Slasher.contract.UnpackLog(event, "OperatorSlashed", log); err != nil {
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

// ParseOperatorSlashed is a log parse operation binding the contract event 0x715df5b5f92bd7692569b1129f8a601730c048c3daecf410b60b52d8a6015d40.
//
// Solidity: event OperatorSlashed(address operator, (address,uint32) operatorSet, uint32 bipsToSlash, address[] strategies, uint64[] slashingRates)
func (_Slasher *SlasherFilterer) ParseOperatorSlashed(log types.Log) (*SlasherOperatorSlashed, error) {
	event := new(SlasherOperatorSlashed)
	if err := _Slasher.contract.UnpackLog(event, "OperatorSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlasherOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Slasher contract.
type SlasherOwnershipTransferredIterator struct {
	Event *SlasherOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SlasherOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlasherOwnershipTransferred)
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
		it.Event = new(SlasherOwnershipTransferred)
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
func (it *SlasherOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlasherOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlasherOwnershipTransferred represents a OwnershipTransferred event raised by the Slasher contract.
type SlasherOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Slasher *SlasherFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SlasherOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Slasher.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SlasherOwnershipTransferredIterator{contract: _Slasher.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Slasher *SlasherFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SlasherOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Slasher.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlasherOwnershipTransferred)
				if err := _Slasher.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Slasher *SlasherFilterer) ParseOwnershipTransferred(log types.Log) (*SlasherOwnershipTransferred, error) {
	event := new(SlasherOwnershipTransferred)
	if err := _Slasher.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlasherPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Slasher contract.
type SlasherPausedIterator struct {
	Event *SlasherPaused // Event containing the contract specifics and raw log

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
func (it *SlasherPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlasherPaused)
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
		it.Event = new(SlasherPaused)
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
func (it *SlasherPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlasherPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlasherPaused represents a Paused event raised by the Slasher contract.
type SlasherPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_Slasher *SlasherFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*SlasherPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Slasher.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &SlasherPausedIterator{contract: _Slasher.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_Slasher *SlasherFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *SlasherPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Slasher.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlasherPaused)
				if err := _Slasher.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Slasher *SlasherFilterer) ParsePaused(log types.Log) (*SlasherPaused, error) {
	event := new(SlasherPaused)
	if err := _Slasher.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlasherPauserRegistrySetIterator is returned from FilterPauserRegistrySet and is used to iterate over the raw logs and unpacked data for PauserRegistrySet events raised by the Slasher contract.
type SlasherPauserRegistrySetIterator struct {
	Event *SlasherPauserRegistrySet // Event containing the contract specifics and raw log

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
func (it *SlasherPauserRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlasherPauserRegistrySet)
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
		it.Event = new(SlasherPauserRegistrySet)
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
func (it *SlasherPauserRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlasherPauserRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlasherPauserRegistrySet represents a PauserRegistrySet event raised by the Slasher contract.
type SlasherPauserRegistrySet struct {
	PauserRegistry    common.Address
	NewPauserRegistry common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPauserRegistrySet is a free log retrieval operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_Slasher *SlasherFilterer) FilterPauserRegistrySet(opts *bind.FilterOpts) (*SlasherPauserRegistrySetIterator, error) {

	logs, sub, err := _Slasher.contract.FilterLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return &SlasherPauserRegistrySetIterator{contract: _Slasher.contract, event: "PauserRegistrySet", logs: logs, sub: sub}, nil
}

// WatchPauserRegistrySet is a free log subscription operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_Slasher *SlasherFilterer) WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *SlasherPauserRegistrySet) (event.Subscription, error) {

	logs, sub, err := _Slasher.contract.WatchLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlasherPauserRegistrySet)
				if err := _Slasher.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
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
func (_Slasher *SlasherFilterer) ParsePauserRegistrySet(log types.Log) (*SlasherPauserRegistrySet, error) {
	event := new(SlasherPauserRegistrySet)
	if err := _Slasher.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlasherTotalAndNonslashableMagnitudeUpdatedIterator is returned from FilterTotalAndNonslashableMagnitudeUpdated and is used to iterate over the raw logs and unpacked data for TotalAndNonslashableMagnitudeUpdated events raised by the Slasher contract.
type SlasherTotalAndNonslashableMagnitudeUpdatedIterator struct {
	Event *SlasherTotalAndNonslashableMagnitudeUpdated // Event containing the contract specifics and raw log

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
func (it *SlasherTotalAndNonslashableMagnitudeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlasherTotalAndNonslashableMagnitudeUpdated)
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
		it.Event = new(SlasherTotalAndNonslashableMagnitudeUpdated)
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
func (it *SlasherTotalAndNonslashableMagnitudeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlasherTotalAndNonslashableMagnitudeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlasherTotalAndNonslashableMagnitudeUpdated represents a TotalAndNonslashableMagnitudeUpdated event raised by the Slasher contract.
type SlasherTotalAndNonslashableMagnitudeUpdated struct {
	Operator                common.Address
	Strategy                common.Address
	EffectTimestamp         uint32
	TotalSlashableMagnitude uint64
	NonslashableMagnitude   uint64
	CumulativeAllocationSum uint64
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterTotalAndNonslashableMagnitudeUpdated is a free log retrieval operation binding the contract event 0x2bc1462fd4652ef49aae855563647b6564ff41eb9db10038d4b08e01d34f30e1.
//
// Solidity: event TotalAndNonslashableMagnitudeUpdated(address operator, address strategy, uint32 effectTimestamp, uint64 totalSlashableMagnitude, uint64 nonslashableMagnitude, uint64 cumulativeAllocationSum)
func (_Slasher *SlasherFilterer) FilterTotalAndNonslashableMagnitudeUpdated(opts *bind.FilterOpts) (*SlasherTotalAndNonslashableMagnitudeUpdatedIterator, error) {

	logs, sub, err := _Slasher.contract.FilterLogs(opts, "TotalAndNonslashableMagnitudeUpdated")
	if err != nil {
		return nil, err
	}
	return &SlasherTotalAndNonslashableMagnitudeUpdatedIterator{contract: _Slasher.contract, event: "TotalAndNonslashableMagnitudeUpdated", logs: logs, sub: sub}, nil
}

// WatchTotalAndNonslashableMagnitudeUpdated is a free log subscription operation binding the contract event 0x2bc1462fd4652ef49aae855563647b6564ff41eb9db10038d4b08e01d34f30e1.
//
// Solidity: event TotalAndNonslashableMagnitudeUpdated(address operator, address strategy, uint32 effectTimestamp, uint64 totalSlashableMagnitude, uint64 nonslashableMagnitude, uint64 cumulativeAllocationSum)
func (_Slasher *SlasherFilterer) WatchTotalAndNonslashableMagnitudeUpdated(opts *bind.WatchOpts, sink chan<- *SlasherTotalAndNonslashableMagnitudeUpdated) (event.Subscription, error) {

	logs, sub, err := _Slasher.contract.WatchLogs(opts, "TotalAndNonslashableMagnitudeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlasherTotalAndNonslashableMagnitudeUpdated)
				if err := _Slasher.contract.UnpackLog(event, "TotalAndNonslashableMagnitudeUpdated", log); err != nil {
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

// ParseTotalAndNonslashableMagnitudeUpdated is a log parse operation binding the contract event 0x2bc1462fd4652ef49aae855563647b6564ff41eb9db10038d4b08e01d34f30e1.
//
// Solidity: event TotalAndNonslashableMagnitudeUpdated(address operator, address strategy, uint32 effectTimestamp, uint64 totalSlashableMagnitude, uint64 nonslashableMagnitude, uint64 cumulativeAllocationSum)
func (_Slasher *SlasherFilterer) ParseTotalAndNonslashableMagnitudeUpdated(log types.Log) (*SlasherTotalAndNonslashableMagnitudeUpdated, error) {
	event := new(SlasherTotalAndNonslashableMagnitudeUpdated)
	if err := _Slasher.contract.UnpackLog(event, "TotalAndNonslashableMagnitudeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlasherUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Slasher contract.
type SlasherUnpausedIterator struct {
	Event *SlasherUnpaused // Event containing the contract specifics and raw log

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
func (it *SlasherUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlasherUnpaused)
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
		it.Event = new(SlasherUnpaused)
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
func (it *SlasherUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlasherUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlasherUnpaused represents a Unpaused event raised by the Slasher contract.
type SlasherUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_Slasher *SlasherFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*SlasherUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Slasher.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &SlasherUnpausedIterator{contract: _Slasher.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_Slasher *SlasherFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *SlasherUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Slasher.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlasherUnpaused)
				if err := _Slasher.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Slasher *SlasherFilterer) ParseUnpaused(log types.Log) (*SlasherUnpaused, error) {
	event := new(SlasherUnpaused)
	if err := _Slasher.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
