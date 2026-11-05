// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package StrategyManager

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

// OperatorSet is an auto generated low-level Go binding around an user-defined struct.
type OperatorSet struct {
	Avs common.Address
	Id  uint32
}

// StrategyManagerMetaData contains all meta data concerning the StrategyManager contract.
var StrategyManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_allocationManager\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"},{\"name\":\"_delegation\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"_version\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_BURN_ADDRESS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEPOSIT_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SLASH_RESOLUTION_DELAY_BLOCKS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addStrategiesToDepositWhitelist\",\"inputs\":[{\"name\":\"strategiesToWhitelist\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allocationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burnShares\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"calculateStrategyDepositDigestHash\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"clearBurnOrRedistributableShares\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"clearBurnOrRedistributableSharesByStrategy\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"depositIntoStrategy\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"depositShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositIntoStrategyWithSignature\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"depositShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"domainSeparator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBurnOrRedistributableCount\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBurnOrRedistributableShares\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBurnOrRedistributableShares\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBurnableShares\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDeposits\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPendingOperatorSets\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPendingSlashIds\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSlashResolutionBlock\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStakerStrategyList\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStrategiesWithBurnableShares\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"increaseBurnOrRedistributableShares\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"sharesToBurn\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialStrategyWhitelister\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"nonces\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeDepositShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"depositSharesToRemove\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeStrategiesFromDepositWhitelist\",\"inputs\":[{\"name\":\"strategiesToRemoveFromWhitelist\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSlashResolutionDelayBlocks\",\"inputs\":[{\"name\":\"newSlashResolutionDelayBlocks\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSlashResolutionDelaySetter\",\"inputs\":[{\"name\":\"newSlashResolutionDelaySetter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStrategyWhitelister\",\"inputs\":[{\"name\":\"newStrategyWhitelister\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slashResolutionDelaySetter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakerDepositShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakerStrategyList\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"strategies\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakerStrategyListLength\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"strategyIsWhitelistedForDeposit\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"whitelisted\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"strategyWhitelister\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawSharesAsTokens\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"BurnOrRedistributableSharesDecreased\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BurnOrRedistributableSharesIncreased\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BurnableSharesDecreased\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposit\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SlashResolutionBlockSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"resolutionBlock\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SlashResolutionDelayBlocksChanged\",\"inputs\":[{\"name\":\"previousDelayBlocks\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"newDelayBlocks\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SlashResolutionDelaySetterChanged\",\"inputs\":[{\"name\":\"previousSetter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newSetter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyAddedToDepositWhitelist\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyRemovedFromDepositWhitelist\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyWhitelisterChanged\",\"inputs\":[{\"name\":\"previousAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CurrentlyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNewPausedStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MaxStrategiesExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyDelegationManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyPauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlySlashResolutionDelaySetter\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyStrategyWhitelister\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnpauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SharesAmountTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SharesAmountZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignatureExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SlashResolutionDelayNotElapsed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StakerAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategyAlreadyInSlash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategyNotFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategyNotWhitelisted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]}]",
	Bin: "0x610100604052348015610010575f5ffd5b50604051613eca380380613eca83398101604081905261002f916101bc565b80808585856001600160a01b03811661005b576040516339b190bb60e11b815260040160405180910390fd5b6001600160a01b0390811660805291821660a0521660c05261007c81610093565b60e0525061008a90506100d9565b50505050610301565b5f5f829050601f815111156100c6578260405163305a27a960e01b81526004016100bd91906102a6565b60405180910390fd5b80516100d1826102db565b179392505050565b5f54610100900460ff16156101405760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b60648201526084016100bd565b5f5460ff9081161461018f575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146101a5575f5ffd5b50565b634e487b7160e01b5f52604160045260245ffd5b5f5f5f5f608085870312156101cf575f5ffd5b84516101da81610191565b60208601519094506101eb81610191565b60408601519093506101fc81610191565b60608601519092506001600160401b03811115610217575f5ffd5b8501601f81018713610227575f5ffd5b80516001600160401b03811115610240576102406101a8565b604051601f8201601f19908116603f011681016001600160401b038111828210171561026e5761026e6101a8565b604052818152828201602001891015610285575f5ffd5b8160208401602083015e5f6020838301015280935050505092959194509250565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b805160208083015191908110156102fb575f198160200360031b1b821691505b50919050565b60805160a05160c05160e051613b5061037a5f395f8181611359015261290a01525f81816106b801528181610a7f01528181610f5c015281816112eb0152818161159601526124de01525f818161065e015281816109d201526117ca01525f818161056101528181611e4901526129790152613b505ff3fe608060405234801561000f575f5ffd5b50600436106102b1575f3560e01c806376fb162b1161017b578063b5d8b5b8116100e4578063e7a050aa1161009e578063f698da2511610079578063f698da251461070a578063fabc1cbc14610712578063fd98042314610725578063fe243a1714610738575f5ffd5b8063e7a050aa146106da578063f2fde38b146106ed578063f3b4a00014610700575f5ffd5b8063b5d8b5b814610633578063c665670214610646578063ca8aa7c714610659578063cbc2bd6214610680578063de44acb614610693578063df5cf723146106b3575f5ffd5b80638b8aac3c116101355780638b8aac3c146105ae5780638da5cb5b146105d657806394f649dd146105e7578063967fc0d2146105fa57806396f022091461060d5780639ac01d6114610620575f5ffd5b806376fb162b146104f75780637def15641461050a5780637ecebe001461052a578063829fca7314610549578063886f11951461055c57806388c102991461059b575f5ffd5b80634b3262f71161021d5780635c975abb116101d75780635c975abb1461048c5780635de08ff214610494578063663c1de4146104a757806367f1b47e146104c9578063715018a6146104dc578063724af423146104e4575f5ffd5b80634b3262f7146103ee5780634b6d5d6e1461040157806350ff72251461041457806354fd4d501461043c578063595c6a67146104515780635ac86ab714610459575f5ffd5b806331f8fb4c1161026e57806331f8fb4c1461035557806332e89ace1461037657806336a8c500146103895780633f292b081461039f5780633fb99ca5146103b457806348825e94146103c7575f5ffd5b806307a79355146102b5578063136439dd146102e65780631794bb3c146102fb57806329acc8421461030e5780632d44def6146103215780632eae418c14610342575b5f5ffd5b60dc546102cc90600160a01b900463ffffffff1681565b60405163ffffffff90911681526020015b60405180910390f35b6102f96102f43660046132c9565b610762565b005b6102f96103093660046132f4565b61079c565b6102f961031c366004613332565b6108c2565b61033461032f366004613363565b61095a565b6040519081526020016102dd565b6102f96103503660046133a1565b610a74565b6103686103633660046133ef565b610b40565b6040516102dd92919061348b565b6103346103843660046134fc565b610cc9565b610391610d4e565b6040516102dd9291906135d6565b6103a7610e69565b6040516102dd919061362c565b6102f96103c2366004613689565b610f51565b6103347f4337f82d142e41f2a8c10547cd8c859bddb92262a61058e77842e24d9dea922481565b6102cc6103fc3660046133ef565b611147565b6102f961040f366004613332565b611189565b6104276104223660046132f4565b6112de565b604080519283526020830191909152016102dd565b610444611352565b6040516102dd91906136fb565b6102f9611382565b61047c61046736600461370d565b609854600160ff9092169190911b9081161490565b60405190151581526020016102dd565b609854610334565b6102f96104a236600461372d565b611396565b61047c6104b5366004613332565b60d16020525f908152604090205460ff1681565b6102f96104d73660046137b4565b6114dd565b6102f9611579565b6103346104f23660046132f4565b61158a565b610334610505366004613363565b6115fa565b61051d6105183660046137cd565b611649565b6040516102dd91906137e7565b610334610538366004613332565b60ca6020525f908152604090205481565b6103346105573660046133ef565b611675565b6105837f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016102dd565b61051d6105a93660046133ef565b6116af565b6103346105bc366004613332565b6001600160a01b03165f90815260ce602052604090205490565b6033546001600160a01b0316610583565b6103686105f5366004613332565b611848565b60cb54610583906001600160a01b031681565b60dc54610583906001600160a01b031681565b61033461062e3660046137f9565b6119bf565b6102f961064136600461372d565b611a50565b6102f9610654366004613332565b611b97565b6105837f000000000000000000000000000000000000000000000000000000000000000081565b61058361068e36600461385a565b611bbd565b6106a66106a1366004613332565b611bf1565b6040516102dd9190613884565b6105837f000000000000000000000000000000000000000000000000000000000000000081565b6103346106e83660046132f4565b611c64565b6102f96106fb366004613332565b611c83565b610583620e16e481565b610334611cf9565b6102f96107203660046132c9565b611db2565b610334610733366004613332565b611e1f565b610334610746366004613896565b60cd60209081525f928352604080842090915290825290205481565b61076a611e34565b609854818116811461078f5760405163c61dca5d60e01b815260040160405180910390fd5b61079882611ed7565b5050565b5f54610100900460ff16158080156107ba57505f54600160ff909116105b806107d35750303b1580156107d357505f5460ff166001145b61083b5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b5f805460ff19166001179055801561085c575f805461ff0019166101001790555b61086582611ed7565b61086e84611f14565b61087783611f65565b80156108bc575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050565b6108ca611fce565b6001600160a01b0381166108f1576040516339b190bb60e11b815260040160405180910390fd5b60dc54604080516001600160a01b03928316815291831660208301527f98355cb0f7088f4584521f1fbb1f7ffd33d427eb8740bc64997df3a6a54078a7910160405180910390a160dc80546001600160a01b0319166001600160a01b0392909216919091179055565b5f600161096681612028565b61096e612053565b60db5f610988610983368990038901896138cd565b6120ac565b815260208082019290925260409081015f90812087825290925290205463ffffffff164310156109ca57604051626f6f0760e41b815260040160405180910390fd5b610a608585857f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316630f3df50e8a6040518263ffffffff1660e01b8152600401610a1c919061395a565b602060405180830381865afa158015610a37573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a5b9190613968565b61210f565b9150610a6c6001606555565b509392505050565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610abd5760405163f739589b60e01b815260040160405180910390fd5b610ac5612053565b604051636ce5768960e11b81526001600160a01b0384169063d9caed1290610af590879086908690600401613983565b6020604051808303815f875af1158015610b11573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b3591906139a7565b506108bc6001606555565b6060805f60d781610b59610983368990038901896138cd565b81526020019081526020015f205f8581526020019081526020015f2090505f610b8182612353565b90505f81516001600160401b03811115610b9d57610b9d6134b8565b604051908082528060200260200182016040528015610bc6578160200160208202803683370190505b5090505f82516001600160401b03811115610be357610be36134b8565b604051908082528060200260200182016040528015610c0c578160200160208202803683370190505b5090505f5b8351811015610cb957838181518110610c2c57610c2c6139be565b6020026020010151838281518110610c4657610c466139be565b60200260200101906001600160a01b031690816001600160a01b031681525050610c92848281518110610c7b57610c7b6139be565b60200260200101518661235f90919063ffffffff16565b9050828281518110610ca657610ca66139be565b6020908102919091010152600101610c11565b50909450925050505b9250929050565b5f5f610cd481612028565b610cdc612053565b6001600160a01b0385165f90815260ca6020526040902054610d0d86610d06818c8c8c878c6119bf565b8688612383565b6001600160a01b0386165f90815260ca60205260409020600182019055610d36868a8a8a6123d5565b925050610d436001606555565b509695505050505050565b6060805f610d5c60d4612542565b90505f816001600160401b03811115610d7757610d776134b8565b604051908082528060200260200182016040528015610da0578160200160208202803683370190505b5090505f826001600160401b03811115610dbc57610dbc6134b8565b604051908082528060200260200182016040528015610de5578160200160208202803683370190505b5090505f5b83811015610e5e575f5f610dff60d48461254c565b9150915081858481518110610e1657610e166139be565b60200260200101906001600160a01b031690816001600160a01b03168152505080848481518110610e4957610e496139be565b60209081029190910101525050600101610dea565b509094909350915050565b60605f610e7660d861255a565b90505f816001600160401b03811115610e9157610e916134b8565b604051908082528060200260200182016040528015610ed557816020015b604080518082019091525f8082526020820152815260200190600190039081610eaf5790505b5090505f5b82811015610f4a57610f25610ef060d883612563565b604080518082019091525f80825260208201525060408051808201909152606082901c815263ffffffff909116602082015290565b828281518110610f3757610f376139be565b6020908102919091010152600101610eda565b5092915050565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610f9a5760405163f739589b60e01b815260040160405180910390fd5b610fa2612053565b5f60d781610fb8610983368990038901896138cd565b815260208082019290925260409081015f90812087825290925290209050610fe181848461256e565b610ffe5760405163ca354fa360e01b815260040160405180910390fd5b61101b611013610983368890038801886138cd565b60d890612583565b506110508460da5f611035610983368b90038b018b6138cd565b81526020019081526020015f2061258390919063ffffffff16565b156110ff5760dc545f9061107190600160a01b900463ffffffff16436139e6565b90508060db5f611089610983368b90038b018b6138cd565b81526020019081526020015f205f8781526020019081526020015f205f6101000a81548163ffffffff021916908363ffffffff1602179055507fa8fd1473d1b9b7b4c9033ecb92a0f6e43f107aec7720ede3709521d122f89fa58686836040516110f593929190613a02565b60405180910390a1505b7f5f5209798bbac45a16d2dc3bc67319fab26ee00153916d6f07b69f8a134a1e8b858585856040516111349493929190613a2a565b60405180910390a1506108bc6001606555565b5f60db8161115d610983368790038701876138cd565b815260208082019290925260409081015f90812085825290925290205463ffffffff1690505b92915050565b600161119481612028565b61119c612053565b5f6111a860d48461235f565b9150506111b660d48461258e565b50604080516001600160a01b0385168152602081018390527fd9d082c3ec4f3a3ffa55c324939a06407f5fbcb87d5e0ce3b9508c92c84ed839910160405180910390a180156112d357826001600160a01b031663d9caed12620e16e4856001600160a01b0316632495a5996040518163ffffffff1660e01b8152600401602060405180830381865afa15801561124e573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906112729190613968565b846040518463ffffffff1660e01b815260040161129193929190613983565b6020604051808303815f875af11580156112ad573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906112d191906139a7565b505b506107986001606555565b5f80336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146113295760405163f739589b60e01b815260040160405180910390fd5b611331612053565b61133c8585856125a2565b9150915061134a6001606555565b935093915050565b606061137d7f0000000000000000000000000000000000000000000000000000000000000000612766565b905090565b61138a611e34565b6113945f19611ed7565b565b60cb546001600160a01b031633146113c1576040516320ba3ff960e21b815260040160405180910390fd5b6113c9612053565b805f5b818110156112d15760d15f8585848181106113e9576113e96139be565b90506020020160208101906113fe9190613332565b6001600160a01b0316815260208101919091526040015f205460ff166114d557600160d15f868685818110611435576114356139be565b905060200201602081019061144a9190613332565b6001600160a01b0316815260208101919091526040015f20805460ff19169115159190911790557f0c35b17d91c96eb2751cd456e1252f42a386e524ef9ff26ecc9950859fdc04fe8484838181106114a4576114a46139be565b90506020020160208101906114b99190613332565b6040516001600160a01b03909116815260200160405180910390a15b6001016113cc565b60dc546001600160a01b03163314611508576040516328a01d4b60e01b815260040160405180910390fd5b60dc546040805163ffffffff600160a01b9093048316815291831660208301527f086f17854c753dd15c19bb0e7683f699b07a44ea9673872d2b5a07079c7ab29a910160405180910390a160dc805463ffffffff909216600160a01b0263ffffffff60a01b19909216919091179055565b611581611fce565b6113945f611f14565b5f336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146115d45760405163f739589b60e01b815260040160405180910390fd5b6115dc612053565b6115e78484846127a3565b90506115f36001606555565b9392505050565b5f8061163f8360d783611615610983368b90038b018b6138cd565b81526020019081526020015f205f8781526020019081526020015f2061235f90919063ffffffff16565b9695505050505050565b606061118360da5f611663610983368790038701876138cd565b81526020019081526020015f206128b0565b5f6115f360d78261168e610983368890038801886138cd565b81526020019081526020015f205f8481526020019081526020015f20612542565b606060016116bc81612028565b6116c4612053565b60db5f6116d9610983368890038801886138cd565b815260208082019290925260409081015f90812086825290925290205463ffffffff1643101561171b57604051626f6f0760e41b815260040160405180910390fd5b5f61175560d782611734610983368a90038a018a6138cd565b81526020019081526020015f205f8681526020019081526020015f20612353565b80519091505f816001600160401b03811115611773576117736134b8565b60405190808252806020026020018201604052801561179c578160200160208202803683370190505b5090505f5b828110156118395761181488888684815181106117c0576117c06139be565b60200260200101517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316630f3df50e8d6040518263ffffffff1660e01b8152600401610a1c919061395a565b828281518110611826576118266139be565b60209081029190910101526001016117a1565b5093505050610f4a6001606555565b6001600160a01b0381165f90815260ce6020526040812054606091829190816001600160401b0381111561187e5761187e6134b8565b6040519080825280602002602001820160405280156118a7578160200160208202803683370190505b5090505f5b82811015611935576001600160a01b0386165f90815260cd6020908152604080832060ce90925282208054919291849081106118ea576118ea6139be565b5f9182526020808320909101546001600160a01b031683528201929092526040019020548251839083908110611922576119226139be565b60209081029190910101526001016118ac565b5060ce5f866001600160a01b03166001600160a01b031681526020019081526020015f2081818054806020026020016040519081016040528092919081815260200182805480156119ad57602002820191905f5260205f20905b81546001600160a01b0316815260019091019060200180831161198f575b50505050509150935093505050915091565b604080517f4337f82d142e41f2a8c10547cd8c859bddb92262a61058e77842e24d9dea922460208201526001600160a01b03808916928201929092528187166060820152908516608082015260a0810184905260c0810183905260e081018290525f90611a459061010001604051602081830303815290604052805190602001206128bc565b979650505050505050565b60cb546001600160a01b03163314611a7b576040516320ba3ff960e21b815260040160405180910390fd5b611a83612053565b805f5b818110156112d15760d15f858584818110611aa357611aa36139be565b9050602002016020810190611ab89190613332565b6001600160a01b0316815260208101919091526040015f205460ff1615611b8f575f60d15f868685818110611aef57611aef6139be565b9050602002016020810190611b049190613332565b6001600160a01b0316815260208101919091526040015f20805460ff19169115159190911790557f4074413b4b443e4e58019f2855a8765113358c7c72e39509c6af45fc0f5ba030848483818110611b5e57611b5e6139be565b9050602002016020810190611b739190613332565b6040516001600160a01b03909116815260200160405180910390a15b600101611a86565b611b9f611fce565b611ba7612053565b611bb081611f65565b611bba6001606555565b50565b60ce602052815f5260405f208181548110611bd6575f80fd5b5f918252602090912001546001600160a01b03169150829050565b6001600160a01b0381165f90815260ce6020908152604091829020805483518184028101840190945280845260609392830182828015611c5857602002820191905f5260205f20905b81546001600160a01b03168152600190910190602001808311611c3a575b50505050509050919050565b5f5f611c6f81612028565b611c77612053565b610a60338686866123d5565b611c8b611fce565b6001600160a01b038116611cf05760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610832565b611bba81611f14565b60408051808201909152600a81526922b4b3b2b72630bcb2b960b11b6020909101525f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f7f71b625cfad44bac63b13dba07f2e1d6084ee04b6f8752101ece6126d584ee6ea611d66612902565b805160209182012060408051928301949094529281019190915260608101919091524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b611dba612977565b60985480198219811614611de15760405163c61dca5d60e01b815260040160405180910390fd5b609882905560405182815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c9060200160405180910390a25050565b5f5f611e2c60d48461235f565b949350505050565b60405163237dfb4760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906346fbf68e90602401602060405180830381865afa158015611e96573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611eba9190613a5c565b61139457604051631d77d47760e21b815260040160405180910390fd5b609881905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a250565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b60cb54604080516001600160a01b03928316815291831660208301527f4264275e593955ff9d6146a51a4525f6ddace2e81db9391abcc9d1ca48047d29910160405180910390a160cb80546001600160a01b0319166001600160a01b0392909216919091179055565b6033546001600160a01b031633146113945760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610832565b609854600160ff83161b90811603611bba5760405163840a48d560e01b815260040160405180910390fd5b6002606554036120a55760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610832565b6002606555565b5f815f0151826020015163ffffffff166040516020016120f792919060609290921b6bffffffffffffffffffffffff1916825260a01b6001600160a01b031916601482015260200190565b60405160208183030381529060405261118390613a7b565b5f8060d781612126610983368a90038a018a6138cd565b815260208082019290925260409081015f9081208882529092528120915061214e828661235f565b915061215c9050828661258e565b505f811561227357856001600160a01b031663d9caed1286886001600160a01b0316632495a5996040518163ffffffff1660e01b8152600401602060405180830381865afa1580156121b0573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906121d49190613968565b856040518463ffffffff1660e01b81526004016121f393929190613983565b6020604051808303815f875af115801561220f573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061223391906139a7565b90507fe6413aa0c789e437b0a06bf64b20926584f066c79a2d8b80a759c85472f7b0af8888888560405161226a9493929190613a2a565b60405180910390a15b5f61227d84612542565b9050805f03612347576122ba8860da5f61229f610983368f90038f018f6138cd565b81526020019081526020015f20612a2890919063ffffffff16565b5060db5f6122d0610983368d90038d018d6138cd565b815260208082019290925260409081015f9081208b82529092528120805463ffffffff191690556123229060da90612310610983368e90038e018e6138cd565b81526020019081526020015f2061255a565b5f036123475761234561233d610983368c90038c018c6138cd565b60d890612a28565b505b50979650505050505050565b60605f6115f383612a33565b5f808080612376866001600160a01b038716612a3e565b9097909650945050505050565b428110156123a457604051630819bdcd60e01b815260040160405180910390fd5b6123b86001600160a01b0385168484612a76565b6108bc57604051638baa579f60e01b815260040160405180910390fd5b6001600160a01b0383165f90815260d16020526040812054849060ff1661240f57604051632efd965160e11b815260040160405180910390fd5b6124246001600160a01b038516338786612aca565b6040516311f9fbc960e21b81526001600160a01b038581166004830152602482018590528616906347e7ef24906044016020604051808303815f875af1158015612470573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061249491906139a7565b91505f5f6124a38888866125a2565b604051631e328e7960e11b81526001600160a01b038b811660048301528a8116602483015260448201849052606482018390529294509092507f000000000000000000000000000000000000000000000000000000000000000090911690633c651cf2906084015f604051808303815f87803b158015612521575f5ffd5b505af1158015612533573d5f5f3e3d5ffd5b50505050505050949350505050565b5f61118382612b22565b5f8080806123768686612b2c565b5f611183825490565b5f6115f38383612b55565b5f611e2c846001600160a01b03851684612b7b565b5f6115f38383612b97565b5f6115f3836001600160a01b038416612be3565b5f806001600160a01b0385166125cb576040516316f2ccc960e01b815260040160405180910390fd5b825f036125eb576040516342061b2560e11b815260040160405180910390fd5b60405162e7c78560e71b81526001600160a01b038681166004830152602482018590528516906373e3c280906044015f604051808303815f87803b158015612631575f5ffd5b505af1158015612643573d5f5f3e3d5ffd5b5050506001600160a01b038087165f90815260cd6020908152604080832093891683529290529081205491508190036126ed576001600160a01b0386165f90815260ce6020908152604090912054106126af576040516301a1443960e31b815260040160405180910390fd5b6001600160a01b038681165f90815260ce602090815260408220805460018101825590835291200180546001600160a01b0319169187169190911790555b6126f78482613a9e565b6001600160a01b038088165f90815260cd60209081526040808320938a16835292905281902091909155517f5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f629061275390889088908890613983565b60405180910390a1959294509192505050565b60605f61277283612bff565b6040805160208082528183019092529192505f91906020820181803683375050509182525060208101929092525090565b5f815f036127c4576040516342061b2560e11b815260040160405180910390fd5b6001600160a01b038085165f90815260cd60209081526040808320938716835292905220548083111561280a57604051634b18b19360e01b815260040160405180910390fd5b6040516303e3e6eb60e01b81526001600160a01b038681166004830152602482018590528516906303e3e6eb906044015f604051808303815f87803b158015612851575f5ffd5b505af1158015612863573d5f5f3e3d5ffd5b5050505082816128739190613ab1565b6001600160a01b038087165f90815260cd602090815260408083209389168352929052908120829055909150819003611e2c57611e2c8585612c26565b60605f6115f383612da4565b5f6128c5611cf9565b60405161190160f01b6020820152602281019190915260428101839052606201604051602081830303815290604052805190602001209050919050565b60605f61292e7f0000000000000000000000000000000000000000000000000000000000000000612766565b9050805f81518110612942576129426139be565b016020908101516040516001600160f81b03199091169181019190915260210160405160208183030381529060405291505090565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156129d3573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906129f79190613968565b6001600160a01b0316336001600160a01b0316146113945760405163794821ff60e01b815260040160405180910390fd5b5f6115f38383612dfc565b6060611183826128b0565b5f818152600283016020526040812054819080612a6b57612a5f8585612edf565b92505f9150610cc29050565b600192509050610cc2565b5f5f5f612a838585612eea565b90925090505f816004811115612a9b57612a9b613ac4565b148015612ab95750856001600160a01b0316826001600160a01b0316145b8061163f575061163f868686612f29565b6108bc846323b872dd60e01b858585604051602401612aeb93929190613983565b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b031990931692909217909152613010565b5f6111838261255a565b5f8080612b398585612563565b5f81815260029690960160205260409095205494959350505050565b5f825f018281548110612b6a57612b6a6139be565b905f5260205f200154905092915050565b5f8281526002840160205260408120829055611e2c8484612583565b5f818152600183016020526040812054612bdc57508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155611183565b505f611183565b5f81815260028301602052604081208190556115f38383612a28565b5f60ff8216601f81111561118357604051632cd44ac360e21b815260040160405180910390fd5b6001600160a01b0382165f90815260ce6020526040812054905b81811015612d38576001600160a01b038481165f90815260ce6020526040902080549185169183908110612c7657612c766139be565b5f918252602090912001546001600160a01b031603612d30576001600160a01b0384165f90815260ce602052604090208054612cb490600190613ab1565b81548110612cc457612cc46139be565b5f9182526020808320909101546001600160a01b03878116845260ce9092526040909220805491909216919083908110612d0057612d006139be565b905f5260205f20015f6101000a8154816001600160a01b0302191690836001600160a01b03160217905550612d38565b600101612c40565b818103612d5857604051632df15a4160e11b815260040160405180910390fd5b6001600160a01b0384165f90815260ce60205260409020805480612d7e57612d7e613ad8565b5f8281526020902081015f1990810180546001600160a01b031916905501905550505050565b6060815f01805480602002602001604051908101604052809291908181526020018280548015611c5857602002820191905f5260205f20905b815481526020019060010190808311612ddd5750505050509050919050565b5f8181526001830160205260408120548015612ed6575f612e1e600183613ab1565b85549091505f90612e3190600190613ab1565b9050818114612e90575f865f018281548110612e4f57612e4f6139be565b905f5260205f200154905080875f018481548110612e6f57612e6f6139be565b5f918252602080832090910192909255918252600188019052604090208390555b8554869080612ea157612ea1613ad8565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f905560019350505050611183565b5f915050611183565b5f6115f383836130e8565b5f5f8251604103612f1e576020830151604084015160608501515f1a612f12878285856130ff565b94509450505050610cc2565b505f90506002610cc2565b5f5f5f856001600160a01b0316631626ba7e60e01b8686604051602401612f51929190613aec565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909416939093179092529051612f8f9190613b04565b5f60405180830381855afa9150503d805f8114612fc7576040519150601f19603f3d011682016040523d82523d5f602084013e612fcc565b606091505b5091509150818015612fe057506020815110155b801561163f57508051630b135d3f60e11b9061300590830160209081019084016139a7565b149695505050505050565b5f613064826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166131bc9092919063ffffffff16565b905080515f14806130845750808060200190518101906130849190613a5c565b6130e35760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608401610832565b505050565b5f81815260018301602052604081205415156115f3565b5f807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a083111561313457505f905060036131b3565b604080515f8082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015613185573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b0381166131ad575f600192509250506131b3565b91505f90505b94509492505050565b6060611e2c84845f85855f5f866001600160a01b031685876040516131e19190613b04565b5f6040518083038185875af1925050503d805f811461321b576040519150601f19603f3d011682016040523d82523d5f602084013e613220565b606091505b5091509150611a45878383876060831561329a5782515f03613293576001600160a01b0385163b6132935760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610832565b5081611e2c565b611e2c83838151156132af5781518083602001fd5b8060405162461bcd60e51b815260040161083291906136fb565b5f602082840312156132d9575f5ffd5b5035919050565b6001600160a01b0381168114611bba575f5ffd5b5f5f5f60608486031215613306575f5ffd5b8335613311816132e0565b92506020840135613321816132e0565b929592945050506040919091013590565b5f60208284031215613342575f5ffd5b81356115f3816132e0565b5f6040828403121561335d575f5ffd5b50919050565b5f5f5f60808486031215613375575f5ffd5b61337f858561334d565b9250604084013591506060840135613396816132e0565b809150509250925092565b5f5f5f5f608085870312156133b4575f5ffd5b84356133bf816132e0565b935060208501356133cf816132e0565b925060408501356133df816132e0565b9396929550929360600135925050565b5f5f60608385031215613400575f5ffd5b61340a848461334d565b946040939093013593505050565b5f8151808452602084019350602083015f5b828110156134515781516001600160a01b031686526020958601959091019060010161342a565b5093949350505050565b5f8151808452602084019350602083015f5b8281101561345157815186526020958601959091019060010161346d565b604081525f61349d6040830185613418565b82810360208401526134af818561345b565b95945050505050565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f191681016001600160401b03811182821017156134f4576134f46134b8565b604052919050565b5f5f5f5f5f5f60c08789031215613511575f5ffd5b863561351c816132e0565b9550602087013561352c816132e0565b9450604087013593506060870135613543816132e0565b92506080870135915060a08701356001600160401b03811115613564575f5ffd5b8701601f81018913613574575f5ffd5b80356001600160401b0381111561358d5761358d6134b8565b6135a0601f8201601f19166020016134cc565b8181528a60208385010111156135b4575f5ffd5b816020840160208301375f602083830101528093505050509295509295509295565b604080825283519082018190525f9060208501906060840190835b818110156136185783516001600160a01b03168352602093840193909201916001016135f1565b5050838103602085015261163f818661345b565b602080825282518282018190525f918401906040840190835b8181101561367e57835180516001600160a01b0316845260209081015163ffffffff168185015290930192604090920191600101613645565b509095945050505050565b5f5f5f5f60a0858703121561369c575f5ffd5b6136a6868661334d565b93506040850135925060608501356136bd816132e0565b9396929550929360800135925050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f6115f360208301846136cd565b5f6020828403121561371d575f5ffd5b813560ff811681146115f3575f5ffd5b5f5f6020838503121561373e575f5ffd5b82356001600160401b03811115613753575f5ffd5b8301601f81018513613763575f5ffd5b80356001600160401b03811115613778575f5ffd5b8560208260051b840101111561378c575f5ffd5b6020919091019590945092505050565b803563ffffffff811681146137af575f5ffd5b919050565b5f602082840312156137c4575f5ffd5b6115f38261379c565b5f604082840312156137dd575f5ffd5b6115f3838361334d565b602081525f6115f3602083018461345b565b5f5f5f5f5f5f60c0878903121561380e575f5ffd5b8635613819816132e0565b95506020870135613829816132e0565b94506040870135613839816132e0565b959894975094956060810135955060808101359460a0909101359350915050565b5f5f6040838503121561386b575f5ffd5b8235613876816132e0565b946020939093013593505050565b602081525f6115f36020830184613418565b5f5f604083850312156138a7575f5ffd5b82356138b2816132e0565b915060208301356138c2816132e0565b809150509250929050565b5f60408284031280156138de575f5ffd5b50604080519081016001600160401b0381118282101715613901576139016134b8565b604052823561390f816132e0565b815261391d6020840161379c565b60208201529392505050565b8035613934816132e0565b6001600160a01b0316825263ffffffff6139506020830161379c565b1660208301525050565b604081016111838284613929565b5f60208284031215613978575f5ffd5b81516115f3816132e0565b6001600160a01b039384168152919092166020820152604081019190915260600190565b5f602082840312156139b7575f5ffd5b5051919050565b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52601160045260245ffd5b63ffffffff8181168382160190811115611183576111836139d2565b60808101613a108286613929565b83604083015263ffffffff83166060830152949350505050565b60a08101613a388287613929565b60408201949094526001600160a01b03929092166060830152608090910152919050565b5f60208284031215613a6c575f5ffd5b815180151581146115f3575f5ffd5b8051602080830151919081101561335d575f1960209190910360031b1b16919050565b80820180821115611183576111836139d2565b81810381811115611183576111836139d2565b634e487b7160e01b5f52602160045260245ffd5b634e487b7160e01b5f52603160045260245ffd5b828152604060208201525f611e2c60408301846136cd565b5f82518060208501845e5f92019182525091905056fea26469706673582212206333f31cd8356e4acc87930920cb90f88a71a7a61b326503a63b331fe85d1ec664736f6c634300081e0033",
}

// StrategyManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use StrategyManagerMetaData.ABI instead.
var StrategyManagerABI = StrategyManagerMetaData.ABI

// StrategyManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StrategyManagerMetaData.Bin instead.
var StrategyManagerBin = StrategyManagerMetaData.Bin

// DeployStrategyManager deploys a new Ethereum contract, binding an instance of StrategyManager to it.
func DeployStrategyManager(auth *bind.TransactOpts, backend bind.ContractBackend, _allocationManager common.Address, _delegation common.Address, _pauserRegistry common.Address, _version string) (common.Address, *types.Transaction, *StrategyManager, error) {
	parsed, err := StrategyManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StrategyManagerBin), backend, _allocationManager, _delegation, _pauserRegistry, _version)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StrategyManager{StrategyManagerCaller: StrategyManagerCaller{contract: contract}, StrategyManagerTransactor: StrategyManagerTransactor{contract: contract}, StrategyManagerFilterer: StrategyManagerFilterer{contract: contract}}, nil
}

// StrategyManager is an auto generated Go binding around an Ethereum contract.
type StrategyManager struct {
	StrategyManagerCaller     // Read-only binding to the contract
	StrategyManagerTransactor // Write-only binding to the contract
	StrategyManagerFilterer   // Log filterer for contract events
}

// StrategyManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type StrategyManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrategyManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StrategyManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrategyManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StrategyManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrategyManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StrategyManagerSession struct {
	Contract     *StrategyManager  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StrategyManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StrategyManagerCallerSession struct {
	Contract *StrategyManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// StrategyManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StrategyManagerTransactorSession struct {
	Contract     *StrategyManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// StrategyManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type StrategyManagerRaw struct {
	Contract *StrategyManager // Generic contract binding to access the raw methods on
}

// StrategyManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StrategyManagerCallerRaw struct {
	Contract *StrategyManagerCaller // Generic read-only contract binding to access the raw methods on
}

// StrategyManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StrategyManagerTransactorRaw struct {
	Contract *StrategyManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStrategyManager creates a new instance of StrategyManager, bound to a specific deployed contract.
func NewStrategyManager(address common.Address, backend bind.ContractBackend) (*StrategyManager, error) {
	contract, err := bindStrategyManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StrategyManager{StrategyManagerCaller: StrategyManagerCaller{contract: contract}, StrategyManagerTransactor: StrategyManagerTransactor{contract: contract}, StrategyManagerFilterer: StrategyManagerFilterer{contract: contract}}, nil
}

// NewStrategyManagerCaller creates a new read-only instance of StrategyManager, bound to a specific deployed contract.
func NewStrategyManagerCaller(address common.Address, caller bind.ContractCaller) (*StrategyManagerCaller, error) {
	contract, err := bindStrategyManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StrategyManagerCaller{contract: contract}, nil
}

// NewStrategyManagerTransactor creates a new write-only instance of StrategyManager, bound to a specific deployed contract.
func NewStrategyManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*StrategyManagerTransactor, error) {
	contract, err := bindStrategyManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StrategyManagerTransactor{contract: contract}, nil
}

// NewStrategyManagerFilterer creates a new log filterer instance of StrategyManager, bound to a specific deployed contract.
func NewStrategyManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*StrategyManagerFilterer, error) {
	contract, err := bindStrategyManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StrategyManagerFilterer{contract: contract}, nil
}

// bindStrategyManager binds a generic wrapper to an already deployed contract.
func bindStrategyManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StrategyManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StrategyManager *StrategyManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StrategyManager.Contract.StrategyManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StrategyManager *StrategyManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrategyManager.Contract.StrategyManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StrategyManager *StrategyManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StrategyManager.Contract.StrategyManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StrategyManager *StrategyManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StrategyManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StrategyManager *StrategyManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrategyManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StrategyManager *StrategyManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StrategyManager.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTBURNADDRESS is a free data retrieval call binding the contract method 0xf3b4a000.
//
// Solidity: function DEFAULT_BURN_ADDRESS() view returns(address)
func (_StrategyManager *StrategyManagerCaller) DEFAULTBURNADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyManager.contract.Call(opts, &out, "DEFAULT_BURN_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DEFAULTBURNADDRESS is a free data retrieval call binding the contract method 0xf3b4a000.
//
// Solidity: function DEFAULT_BURN_ADDRESS() view returns(address)
func (_StrategyManager *StrategyManagerSession) DEFAULTBURNADDRESS() (common.Address, error) {
	return _StrategyManager.Contract.DEFAULTBURNADDRESS(&_StrategyManager.CallOpts)
}

// DEFAULTBURNADDRESS is a free data retrieval call binding the contract method 0xf3b4a000.
//
// Solidity: function DEFAULT_BURN_ADDRESS() view returns(address)
func (_StrategyManager *StrategyManagerCallerSession) DEFAULTBURNADDRESS() (common.Address, error) {
	return _StrategyManager.Contract.DEFAULTBURNADDRESS(&_StrategyManager.CallOpts)
}

// DEPOSITTYPEHASH is a free data retrieval call binding the contract method 0x48825e94.
//
// Solidity: function DEPOSIT_TYPEHASH() view returns(bytes32)
func (_StrategyManager *StrategyManagerCaller) DEPOSITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StrategyManager.contract.Call(opts, &out, "DEPOSIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEPOSITTYPEHASH is a free data retrieval call binding the contract method 0x48825e94.
//
// Solidity: function DEPOSIT_TYPEHASH() view returns(bytes32)
func (_StrategyManager *StrategyManagerSession) DEPOSITTYPEHASH() ([32]byte, error) {
	return _StrategyManager.Contract.DEPOSITTYPEHASH(&_StrategyManager.CallOpts)
}

// DEPOSITTYPEHASH is a free data retrieval call binding the contract method 0x48825e94.
//
// Solidity: function DEPOSIT_TYPEHASH() view returns(bytes32)
func (_StrategyManager *StrategyManagerCallerSession) DEPOSITTYPEHASH() ([32]byte, error) {
	return _StrategyManager.Contract.DEPOSITTYPEHASH(&_StrategyManager.CallOpts)
}

// SLASHRESOLUTIONDELAYBLOCKS is a free data retrieval call binding the contract method 0x07a79355.
//
// Solidity: function SLASH_RESOLUTION_DELAY_BLOCKS() view returns(uint32)
func (_StrategyManager *StrategyManagerCaller) SLASHRESOLUTIONDELAYBLOCKS(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _StrategyManager.contract.Call(opts, &out, "SLASH_RESOLUTION_DELAY_BLOCKS")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// SLASHRESOLUTIONDELAYBLOCKS is a free data retrieval call binding the contract method 0x07a79355.
//
// Solidity: function SLASH_RESOLUTION_DELAY_BLOCKS() view returns(uint32)
func (_StrategyManager *StrategyManagerSession) SLASHRESOLUTIONDELAYBLOCKS() (uint32, error) {
	return _StrategyManager.Contract.SLASHRESOLUTIONDELAYBLOCKS(&_StrategyManager.CallOpts)
}

// SLASHRESOLUTIONDELAYBLOCKS is a free data retrieval call binding the contract method 0x07a79355.
//
// Solidity: function SLASH_RESOLUTION_DELAY_BLOCKS() view returns(uint32)
func (_StrategyManager *StrategyManagerCallerSession) SLASHRESOLUTIONDELAYBLOCKS() (uint32, error) {
	return _StrategyManager.Contract.SLASHRESOLUTIONDELAYBLOCKS(&_StrategyManager.CallOpts)
}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_StrategyManager *StrategyManagerCaller) AllocationManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyManager.contract.Call(opts, &out, "allocationManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_StrategyManager *StrategyManagerSession) AllocationManager() (common.Address, error) {
	return _StrategyManager.Contract.AllocationManager(&_StrategyManager.CallOpts)
}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_StrategyManager *StrategyManagerCallerSession) AllocationManager() (common.Address, error) {
	return _StrategyManager.Contract.AllocationManager(&_StrategyManager.CallOpts)
}

// CalculateStrategyDepositDigestHash is a free data retrieval call binding the contract method 0x9ac01d61.
//
// Solidity: function calculateStrategyDepositDigestHash(address staker, address strategy, address token, uint256 amount, uint256 nonce, uint256 expiry) view returns(bytes32)
func (_StrategyManager *StrategyManagerCaller) CalculateStrategyDepositDigestHash(opts *bind.CallOpts, staker common.Address, strategy common.Address, token common.Address, amount *big.Int, nonce *big.Int, expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _StrategyManager.contract.Call(opts, &out, "calculateStrategyDepositDigestHash", staker, strategy, token, amount, nonce, expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateStrategyDepositDigestHash is a free data retrieval call binding the contract method 0x9ac01d61.
//
// Solidity: function calculateStrategyDepositDigestHash(address staker, address strategy, address token, uint256 amount, uint256 nonce, uint256 expiry) view returns(bytes32)
func (_StrategyManager *StrategyManagerSession) CalculateStrategyDepositDigestHash(staker common.Address, strategy common.Address, token common.Address, amount *big.Int, nonce *big.Int, expiry *big.Int) ([32]byte, error) {
	return _StrategyManager.Contract.CalculateStrategyDepositDigestHash(&_StrategyManager.CallOpts, staker, strategy, token, amount, nonce, expiry)
}

// CalculateStrategyDepositDigestHash is a free data retrieval call binding the contract method 0x9ac01d61.
//
// Solidity: function calculateStrategyDepositDigestHash(address staker, address strategy, address token, uint256 amount, uint256 nonce, uint256 expiry) view returns(bytes32)
func (_StrategyManager *StrategyManagerCallerSession) CalculateStrategyDepositDigestHash(staker common.Address, strategy common.Address, token common.Address, amount *big.Int, nonce *big.Int, expiry *big.Int) ([32]byte, error) {
	return _StrategyManager.Contract.CalculateStrategyDepositDigestHash(&_StrategyManager.CallOpts, staker, strategy, token, amount, nonce, expiry)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_StrategyManager *StrategyManagerCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyManager.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_StrategyManager *StrategyManagerSession) Delegation() (common.Address, error) {
	return _StrategyManager.Contract.Delegation(&_StrategyManager.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_StrategyManager *StrategyManagerCallerSession) Delegation() (common.Address, error) {
	return _StrategyManager.Contract.Delegation(&_StrategyManager.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_StrategyManager *StrategyManagerCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StrategyManager.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_StrategyManager *StrategyManagerSession) DomainSeparator() ([32]byte, error) {
	return _StrategyManager.Contract.DomainSeparator(&_StrategyManager.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_StrategyManager *StrategyManagerCallerSession) DomainSeparator() ([32]byte, error) {
	return _StrategyManager.Contract.DomainSeparator(&_StrategyManager.CallOpts)
}

// GetBurnOrRedistributableCount is a free data retrieval call binding the contract method 0x829fca73.
//
// Solidity: function getBurnOrRedistributableCount((address,uint32) operatorSet, uint256 slashId) view returns(uint256)
func (_StrategyManager *StrategyManagerCaller) GetBurnOrRedistributableCount(opts *bind.CallOpts, operatorSet OperatorSet, slashId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StrategyManager.contract.Call(opts, &out, "getBurnOrRedistributableCount", operatorSet, slashId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBurnOrRedistributableCount is a free data retrieval call binding the contract method 0x829fca73.
//
// Solidity: function getBurnOrRedistributableCount((address,uint32) operatorSet, uint256 slashId) view returns(uint256)
func (_StrategyManager *StrategyManagerSession) GetBurnOrRedistributableCount(operatorSet OperatorSet, slashId *big.Int) (*big.Int, error) {
	return _StrategyManager.Contract.GetBurnOrRedistributableCount(&_StrategyManager.CallOpts, operatorSet, slashId)
}

// GetBurnOrRedistributableCount is a free data retrieval call binding the contract method 0x829fca73.
//
// Solidity: function getBurnOrRedistributableCount((address,uint32) operatorSet, uint256 slashId) view returns(uint256)
func (_StrategyManager *StrategyManagerCallerSession) GetBurnOrRedistributableCount(operatorSet OperatorSet, slashId *big.Int) (*big.Int, error) {
	return _StrategyManager.Contract.GetBurnOrRedistributableCount(&_StrategyManager.CallOpts, operatorSet, slashId)
}

// GetBurnOrRedistributableShares is a free data retrieval call binding the contract method 0x31f8fb4c.
//
// Solidity: function getBurnOrRedistributableShares((address,uint32) operatorSet, uint256 slashId) view returns(address[], uint256[])
func (_StrategyManager *StrategyManagerCaller) GetBurnOrRedistributableShares(opts *bind.CallOpts, operatorSet OperatorSet, slashId *big.Int) ([]common.Address, []*big.Int, error) {
	var out []interface{}
	err := _StrategyManager.contract.Call(opts, &out, "getBurnOrRedistributableShares", operatorSet, slashId)

	if err != nil {
		return *new([]common.Address), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// GetBurnOrRedistributableShares is a free data retrieval call binding the contract method 0x31f8fb4c.
//
// Solidity: function getBurnOrRedistributableShares((address,uint32) operatorSet, uint256 slashId) view returns(address[], uint256[])
func (_StrategyManager *StrategyManagerSession) GetBurnOrRedistributableShares(operatorSet OperatorSet, slashId *big.Int) ([]common.Address, []*big.Int, error) {
	return _StrategyManager.Contract.GetBurnOrRedistributableShares(&_StrategyManager.CallOpts, operatorSet, slashId)
}

// GetBurnOrRedistributableShares is a free data retrieval call binding the contract method 0x31f8fb4c.
//
// Solidity: function getBurnOrRedistributableShares((address,uint32) operatorSet, uint256 slashId) view returns(address[], uint256[])
func (_StrategyManager *StrategyManagerCallerSession) GetBurnOrRedistributableShares(operatorSet OperatorSet, slashId *big.Int) ([]common.Address, []*big.Int, error) {
	return _StrategyManager.Contract.GetBurnOrRedistributableShares(&_StrategyManager.CallOpts, operatorSet, slashId)
}

// GetBurnOrRedistributableShares0 is a free data retrieval call binding the contract method 0x76fb162b.
//
// Solidity: function getBurnOrRedistributableShares((address,uint32) operatorSet, uint256 slashId, address strategy) view returns(uint256)
func (_StrategyManager *StrategyManagerCaller) GetBurnOrRedistributableShares0(opts *bind.CallOpts, operatorSet OperatorSet, slashId *big.Int, strategy common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StrategyManager.contract.Call(opts, &out, "getBurnOrRedistributableShares0", operatorSet, slashId, strategy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBurnOrRedistributableShares0 is a free data retrieval call binding the contract method 0x76fb162b.
//
// Solidity: function getBurnOrRedistributableShares((address,uint32) operatorSet, uint256 slashId, address strategy) view returns(uint256)
func (_StrategyManager *StrategyManagerSession) GetBurnOrRedistributableShares0(operatorSet OperatorSet, slashId *big.Int, strategy common.Address) (*big.Int, error) {
	return _StrategyManager.Contract.GetBurnOrRedistributableShares0(&_StrategyManager.CallOpts, operatorSet, slashId, strategy)
}

// GetBurnOrRedistributableShares0 is a free data retrieval call binding the contract method 0x76fb162b.
//
// Solidity: function getBurnOrRedistributableShares((address,uint32) operatorSet, uint256 slashId, address strategy) view returns(uint256)
func (_StrategyManager *StrategyManagerCallerSession) GetBurnOrRedistributableShares0(operatorSet OperatorSet, slashId *big.Int, strategy common.Address) (*big.Int, error) {
	return _StrategyManager.Contract.GetBurnOrRedistributableShares0(&_StrategyManager.CallOpts, operatorSet, slashId, strategy)
}

// GetBurnableShares is a free data retrieval call binding the contract method 0xfd980423.
//
// Solidity: function getBurnableShares(address strategy) view returns(uint256)
func (_StrategyManager *StrategyManagerCaller) GetBurnableShares(opts *bind.CallOpts, strategy common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StrategyManager.contract.Call(opts, &out, "getBurnableShares", strategy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBurnableShares is a free data retrieval call binding the contract method 0xfd980423.
//
// Solidity: function getBurnableShares(address strategy) view returns(uint256)
func (_StrategyManager *StrategyManagerSession) GetBurnableShares(strategy common.Address) (*big.Int, error) {
	return _StrategyManager.Contract.GetBurnableShares(&_StrategyManager.CallOpts, strategy)
}

// GetBurnableShares is a free data retrieval call binding the contract method 0xfd980423.
//
// Solidity: function getBurnableShares(address strategy) view returns(uint256)
func (_StrategyManager *StrategyManagerCallerSession) GetBurnableShares(strategy common.Address) (*big.Int, error) {
	return _StrategyManager.Contract.GetBurnableShares(&_StrategyManager.CallOpts, strategy)
}

// GetDeposits is a free data retrieval call binding the contract method 0x94f649dd.
//
// Solidity: function getDeposits(address staker) view returns(address[], uint256[])
func (_StrategyManager *StrategyManagerCaller) GetDeposits(opts *bind.CallOpts, staker common.Address) ([]common.Address, []*big.Int, error) {
	var out []interface{}
	err := _StrategyManager.contract.Call(opts, &out, "getDeposits", staker)

	if err != nil {
		return *new([]common.Address), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// GetDeposits is a free data retrieval call binding the contract method 0x94f649dd.
//
// Solidity: function getDeposits(address staker) view returns(address[], uint256[])
func (_StrategyManager *StrategyManagerSession) GetDeposits(staker common.Address) ([]common.Address, []*big.Int, error) {
	return _StrategyManager.Contract.GetDeposits(&_StrategyManager.CallOpts, staker)
}

// GetDeposits is a free data retrieval call binding the contract method 0x94f649dd.
//
// Solidity: function getDeposits(address staker) view returns(address[], uint256[])
func (_StrategyManager *StrategyManagerCallerSession) GetDeposits(staker common.Address) ([]common.Address, []*big.Int, error) {
	return _StrategyManager.Contract.GetDeposits(&_StrategyManager.CallOpts, staker)
}

// GetPendingOperatorSets is a free data retrieval call binding the contract method 0x3f292b08.
//
// Solidity: function getPendingOperatorSets() view returns((address,uint32)[])
func (_StrategyManager *StrategyManagerCaller) GetPendingOperatorSets(opts *bind.CallOpts) ([]OperatorSet, error) {
	var out []interface{}
	err := _StrategyManager.contract.Call(opts, &out, "getPendingOperatorSets")

	if err != nil {
		return *new([]OperatorSet), err
	}

	out0 := *abi.ConvertType(out[0], new([]OperatorSet)).(*[]OperatorSet)

	return out0, err

}

// GetPendingOperatorSets is a free data retrieval call binding the contract method 0x3f292b08.
//
// Solidity: function getPendingOperatorSets() view returns((address,uint32)[])
func (_StrategyManager *StrategyManagerSession) GetPendingOperatorSets() ([]OperatorSet, error) {
	return _StrategyManager.Contract.GetPendingOperatorSets(&_StrategyManager.CallOpts)
}

// GetPendingOperatorSets is a free data retrieval call binding the contract method 0x3f292b08.
//
// Solidity: function getPendingOperatorSets() view returns((address,uint32)[])
func (_StrategyManager *StrategyManagerCallerSession) GetPendingOperatorSets() ([]OperatorSet, error) {
	return _StrategyManager.Contract.GetPendingOperatorSets(&_StrategyManager.CallOpts)
}

// GetPendingSlashIds is a free data retrieval call binding the contract method 0x7def1564.
//
// Solidity: function getPendingSlashIds((address,uint32) operatorSet) view returns(uint256[])
func (_StrategyManager *StrategyManagerCaller) GetPendingSlashIds(opts *bind.CallOpts, operatorSet OperatorSet) ([]*big.Int, error) {
	var out []interface{}
	err := _StrategyManager.contract.Call(opts, &out, "getPendingSlashIds", operatorSet)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetPendingSlashIds is a free data retrieval call binding the contract method 0x7def1564.
//
// Solidity: function getPendingSlashIds((address,uint32) operatorSet) view returns(uint256[])
func (_StrategyManager *StrategyManagerSession) GetPendingSlashIds(operatorSet OperatorSet) ([]*big.Int, error) {
	return _StrategyManager.Contract.GetPendingSlashIds(&_StrategyManager.CallOpts, operatorSet)
}

// GetPendingSlashIds is a free data retrieval call binding the contract method 0x7def1564.
//
// Solidity: function getPendingSlashIds((address,uint32) operatorSet) view returns(uint256[])
func (_StrategyManager *StrategyManagerCallerSession) GetPendingSlashIds(operatorSet OperatorSet) ([]*big.Int, error) {
	return _StrategyManager.Contract.GetPendingSlashIds(&_StrategyManager.CallOpts, operatorSet)
}

// GetSlashResolutionBlock is a free data retrieval call binding the contract method 0x4b3262f7.
//
// Solidity: function getSlashResolutionBlock((address,uint32) operatorSet, uint256 slashId) view returns(uint32)
func (_StrategyManager *StrategyManagerCaller) GetSlashResolutionBlock(opts *bind.CallOpts, operatorSet OperatorSet, slashId *big.Int) (uint32, error) {
	var out []interface{}
	err := _StrategyManager.contract.Call(opts, &out, "getSlashResolutionBlock", operatorSet, slashId)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetSlashResolutionBlock is a free data retrieval call binding the contract method 0x4b3262f7.
//
// Solidity: function getSlashResolutionBlock((address,uint32) operatorSet, uint256 slashId) view returns(uint32)
func (_StrategyManager *StrategyManagerSession) GetSlashResolutionBlock(operatorSet OperatorSet, slashId *big.Int) (uint32, error) {
	return _StrategyManager.Contract.GetSlashResolutionBlock(&_StrategyManager.CallOpts, operatorSet, slashId)
}

// GetSlashResolutionBlock is a free data retrieval call binding the contract method 0x4b3262f7.
//
// Solidity: function getSlashResolutionBlock((address,uint32) operatorSet, uint256 slashId) view returns(uint32)
func (_StrategyManager *StrategyManagerCallerSession) GetSlashResolutionBlock(operatorSet OperatorSet, slashId *big.Int) (uint32, error) {
	return _StrategyManager.Contract.GetSlashResolutionBlock(&_StrategyManager.CallOpts, operatorSet, slashId)
}

// GetStakerStrategyList is a free data retrieval call binding the contract method 0xde44acb6.
//
// Solidity: function getStakerStrategyList(address staker) view returns(address[])
func (_StrategyManager *StrategyManagerCaller) GetStakerStrategyList(opts *bind.CallOpts, staker common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _StrategyManager.contract.Call(opts, &out, "getStakerStrategyList", staker)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetStakerStrategyList is a free data retrieval call binding the contract method 0xde44acb6.
//
// Solidity: function getStakerStrategyList(address staker) view returns(address[])
func (_StrategyManager *StrategyManagerSession) GetStakerStrategyList(staker common.Address) ([]common.Address, error) {
	return _StrategyManager.Contract.GetStakerStrategyList(&_StrategyManager.CallOpts, staker)
}

// GetStakerStrategyList is a free data retrieval call binding the contract method 0xde44acb6.
//
// Solidity: function getStakerStrategyList(address staker) view returns(address[])
func (_StrategyManager *StrategyManagerCallerSession) GetStakerStrategyList(staker common.Address) ([]common.Address, error) {
	return _StrategyManager.Contract.GetStakerStrategyList(&_StrategyManager.CallOpts, staker)
}

// GetStrategiesWithBurnableShares is a free data retrieval call binding the contract method 0x36a8c500.
//
// Solidity: function getStrategiesWithBurnableShares() view returns(address[], uint256[])
func (_StrategyManager *StrategyManagerCaller) GetStrategiesWithBurnableShares(opts *bind.CallOpts) ([]common.Address, []*big.Int, error) {
	var out []interface{}
	err := _StrategyManager.contract.Call(opts, &out, "getStrategiesWithBurnableShares")

	if err != nil {
		return *new([]common.Address), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// GetStrategiesWithBurnableShares is a free data retrieval call binding the contract method 0x36a8c500.
//
// Solidity: function getStrategiesWithBurnableShares() view returns(address[], uint256[])
func (_StrategyManager *StrategyManagerSession) GetStrategiesWithBurnableShares() ([]common.Address, []*big.Int, error) {
	return _StrategyManager.Contract.GetStrategiesWithBurnableShares(&_StrategyManager.CallOpts)
}

// GetStrategiesWithBurnableShares is a free data retrieval call binding the contract method 0x36a8c500.
//
// Solidity: function getStrategiesWithBurnableShares() view returns(address[], uint256[])
func (_StrategyManager *StrategyManagerCallerSession) GetStrategiesWithBurnableShares() ([]common.Address, []*big.Int, error) {
	return _StrategyManager.Contract.GetStrategiesWithBurnableShares(&_StrategyManager.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address signer) view returns(uint256 nonce)
func (_StrategyManager *StrategyManagerCaller) Nonces(opts *bind.CallOpts, signer common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StrategyManager.contract.Call(opts, &out, "nonces", signer)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address signer) view returns(uint256 nonce)
func (_StrategyManager *StrategyManagerSession) Nonces(signer common.Address) (*big.Int, error) {
	return _StrategyManager.Contract.Nonces(&_StrategyManager.CallOpts, signer)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address signer) view returns(uint256 nonce)
func (_StrategyManager *StrategyManagerCallerSession) Nonces(signer common.Address) (*big.Int, error) {
	return _StrategyManager.Contract.Nonces(&_StrategyManager.CallOpts, signer)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StrategyManager *StrategyManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StrategyManager *StrategyManagerSession) Owner() (common.Address, error) {
	return _StrategyManager.Contract.Owner(&_StrategyManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StrategyManager *StrategyManagerCallerSession) Owner() (common.Address, error) {
	return _StrategyManager.Contract.Owner(&_StrategyManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_StrategyManager *StrategyManagerCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _StrategyManager.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_StrategyManager *StrategyManagerSession) Paused(index uint8) (bool, error) {
	return _StrategyManager.Contract.Paused(&_StrategyManager.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_StrategyManager *StrategyManagerCallerSession) Paused(index uint8) (bool, error) {
	return _StrategyManager.Contract.Paused(&_StrategyManager.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_StrategyManager *StrategyManagerCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrategyManager.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_StrategyManager *StrategyManagerSession) Paused0() (*big.Int, error) {
	return _StrategyManager.Contract.Paused0(&_StrategyManager.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_StrategyManager *StrategyManagerCallerSession) Paused0() (*big.Int, error) {
	return _StrategyManager.Contract.Paused0(&_StrategyManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_StrategyManager *StrategyManagerCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyManager.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_StrategyManager *StrategyManagerSession) PauserRegistry() (common.Address, error) {
	return _StrategyManager.Contract.PauserRegistry(&_StrategyManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_StrategyManager *StrategyManagerCallerSession) PauserRegistry() (common.Address, error) {
	return _StrategyManager.Contract.PauserRegistry(&_StrategyManager.CallOpts)
}

// SlashResolutionDelaySetter is a free data retrieval call binding the contract method 0x96f02209.
//
// Solidity: function slashResolutionDelaySetter() view returns(address)
func (_StrategyManager *StrategyManagerCaller) SlashResolutionDelaySetter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyManager.contract.Call(opts, &out, "slashResolutionDelaySetter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SlashResolutionDelaySetter is a free data retrieval call binding the contract method 0x96f02209.
//
// Solidity: function slashResolutionDelaySetter() view returns(address)
func (_StrategyManager *StrategyManagerSession) SlashResolutionDelaySetter() (common.Address, error) {
	return _StrategyManager.Contract.SlashResolutionDelaySetter(&_StrategyManager.CallOpts)
}

// SlashResolutionDelaySetter is a free data retrieval call binding the contract method 0x96f02209.
//
// Solidity: function slashResolutionDelaySetter() view returns(address)
func (_StrategyManager *StrategyManagerCallerSession) SlashResolutionDelaySetter() (common.Address, error) {
	return _StrategyManager.Contract.SlashResolutionDelaySetter(&_StrategyManager.CallOpts)
}

// StakerDepositShares is a free data retrieval call binding the contract method 0xfe243a17.
//
// Solidity: function stakerDepositShares(address staker, address strategy) view returns(uint256 shares)
func (_StrategyManager *StrategyManagerCaller) StakerDepositShares(opts *bind.CallOpts, staker common.Address, strategy common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StrategyManager.contract.Call(opts, &out, "stakerDepositShares", staker, strategy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerDepositShares is a free data retrieval call binding the contract method 0xfe243a17.
//
// Solidity: function stakerDepositShares(address staker, address strategy) view returns(uint256 shares)
func (_StrategyManager *StrategyManagerSession) StakerDepositShares(staker common.Address, strategy common.Address) (*big.Int, error) {
	return _StrategyManager.Contract.StakerDepositShares(&_StrategyManager.CallOpts, staker, strategy)
}

// StakerDepositShares is a free data retrieval call binding the contract method 0xfe243a17.
//
// Solidity: function stakerDepositShares(address staker, address strategy) view returns(uint256 shares)
func (_StrategyManager *StrategyManagerCallerSession) StakerDepositShares(staker common.Address, strategy common.Address) (*big.Int, error) {
	return _StrategyManager.Contract.StakerDepositShares(&_StrategyManager.CallOpts, staker, strategy)
}

// StakerStrategyList is a free data retrieval call binding the contract method 0xcbc2bd62.
//
// Solidity: function stakerStrategyList(address staker, uint256 ) view returns(address strategies)
func (_StrategyManager *StrategyManagerCaller) StakerStrategyList(opts *bind.CallOpts, staker common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _StrategyManager.contract.Call(opts, &out, "stakerStrategyList", staker, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakerStrategyList is a free data retrieval call binding the contract method 0xcbc2bd62.
//
// Solidity: function stakerStrategyList(address staker, uint256 ) view returns(address strategies)
func (_StrategyManager *StrategyManagerSession) StakerStrategyList(staker common.Address, arg1 *big.Int) (common.Address, error) {
	return _StrategyManager.Contract.StakerStrategyList(&_StrategyManager.CallOpts, staker, arg1)
}

// StakerStrategyList is a free data retrieval call binding the contract method 0xcbc2bd62.
//
// Solidity: function stakerStrategyList(address staker, uint256 ) view returns(address strategies)
func (_StrategyManager *StrategyManagerCallerSession) StakerStrategyList(staker common.Address, arg1 *big.Int) (common.Address, error) {
	return _StrategyManager.Contract.StakerStrategyList(&_StrategyManager.CallOpts, staker, arg1)
}

// StakerStrategyListLength is a free data retrieval call binding the contract method 0x8b8aac3c.
//
// Solidity: function stakerStrategyListLength(address staker) view returns(uint256)
func (_StrategyManager *StrategyManagerCaller) StakerStrategyListLength(opts *bind.CallOpts, staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StrategyManager.contract.Call(opts, &out, "stakerStrategyListLength", staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerStrategyListLength is a free data retrieval call binding the contract method 0x8b8aac3c.
//
// Solidity: function stakerStrategyListLength(address staker) view returns(uint256)
func (_StrategyManager *StrategyManagerSession) StakerStrategyListLength(staker common.Address) (*big.Int, error) {
	return _StrategyManager.Contract.StakerStrategyListLength(&_StrategyManager.CallOpts, staker)
}

// StakerStrategyListLength is a free data retrieval call binding the contract method 0x8b8aac3c.
//
// Solidity: function stakerStrategyListLength(address staker) view returns(uint256)
func (_StrategyManager *StrategyManagerCallerSession) StakerStrategyListLength(staker common.Address) (*big.Int, error) {
	return _StrategyManager.Contract.StakerStrategyListLength(&_StrategyManager.CallOpts, staker)
}

// StrategyIsWhitelistedForDeposit is a free data retrieval call binding the contract method 0x663c1de4.
//
// Solidity: function strategyIsWhitelistedForDeposit(address strategy) view returns(bool whitelisted)
func (_StrategyManager *StrategyManagerCaller) StrategyIsWhitelistedForDeposit(opts *bind.CallOpts, strategy common.Address) (bool, error) {
	var out []interface{}
	err := _StrategyManager.contract.Call(opts, &out, "strategyIsWhitelistedForDeposit", strategy)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StrategyIsWhitelistedForDeposit is a free data retrieval call binding the contract method 0x663c1de4.
//
// Solidity: function strategyIsWhitelistedForDeposit(address strategy) view returns(bool whitelisted)
func (_StrategyManager *StrategyManagerSession) StrategyIsWhitelistedForDeposit(strategy common.Address) (bool, error) {
	return _StrategyManager.Contract.StrategyIsWhitelistedForDeposit(&_StrategyManager.CallOpts, strategy)
}

// StrategyIsWhitelistedForDeposit is a free data retrieval call binding the contract method 0x663c1de4.
//
// Solidity: function strategyIsWhitelistedForDeposit(address strategy) view returns(bool whitelisted)
func (_StrategyManager *StrategyManagerCallerSession) StrategyIsWhitelistedForDeposit(strategy common.Address) (bool, error) {
	return _StrategyManager.Contract.StrategyIsWhitelistedForDeposit(&_StrategyManager.CallOpts, strategy)
}

// StrategyWhitelister is a free data retrieval call binding the contract method 0x967fc0d2.
//
// Solidity: function strategyWhitelister() view returns(address)
func (_StrategyManager *StrategyManagerCaller) StrategyWhitelister(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyManager.contract.Call(opts, &out, "strategyWhitelister")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StrategyWhitelister is a free data retrieval call binding the contract method 0x967fc0d2.
//
// Solidity: function strategyWhitelister() view returns(address)
func (_StrategyManager *StrategyManagerSession) StrategyWhitelister() (common.Address, error) {
	return _StrategyManager.Contract.StrategyWhitelister(&_StrategyManager.CallOpts)
}

// StrategyWhitelister is a free data retrieval call binding the contract method 0x967fc0d2.
//
// Solidity: function strategyWhitelister() view returns(address)
func (_StrategyManager *StrategyManagerCallerSession) StrategyWhitelister() (common.Address, error) {
	return _StrategyManager.Contract.StrategyWhitelister(&_StrategyManager.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_StrategyManager *StrategyManagerCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StrategyManager.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_StrategyManager *StrategyManagerSession) Version() (string, error) {
	return _StrategyManager.Contract.Version(&_StrategyManager.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_StrategyManager *StrategyManagerCallerSession) Version() (string, error) {
	return _StrategyManager.Contract.Version(&_StrategyManager.CallOpts)
}

// AddShares is a paid mutator transaction binding the contract method 0x50ff7225.
//
// Solidity: function addShares(address staker, address strategy, uint256 shares) returns(uint256, uint256)
func (_StrategyManager *StrategyManagerTransactor) AddShares(opts *bind.TransactOpts, staker common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error) {
	return _StrategyManager.contract.Transact(opts, "addShares", staker, strategy, shares)
}

// AddShares is a paid mutator transaction binding the contract method 0x50ff7225.
//
// Solidity: function addShares(address staker, address strategy, uint256 shares) returns(uint256, uint256)
func (_StrategyManager *StrategyManagerSession) AddShares(staker common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error) {
	return _StrategyManager.Contract.AddShares(&_StrategyManager.TransactOpts, staker, strategy, shares)
}

// AddShares is a paid mutator transaction binding the contract method 0x50ff7225.
//
// Solidity: function addShares(address staker, address strategy, uint256 shares) returns(uint256, uint256)
func (_StrategyManager *StrategyManagerTransactorSession) AddShares(staker common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error) {
	return _StrategyManager.Contract.AddShares(&_StrategyManager.TransactOpts, staker, strategy, shares)
}

// AddStrategiesToDepositWhitelist is a paid mutator transaction binding the contract method 0x5de08ff2.
//
// Solidity: function addStrategiesToDepositWhitelist(address[] strategiesToWhitelist) returns()
func (_StrategyManager *StrategyManagerTransactor) AddStrategiesToDepositWhitelist(opts *bind.TransactOpts, strategiesToWhitelist []common.Address) (*types.Transaction, error) {
	return _StrategyManager.contract.Transact(opts, "addStrategiesToDepositWhitelist", strategiesToWhitelist)
}

// AddStrategiesToDepositWhitelist is a paid mutator transaction binding the contract method 0x5de08ff2.
//
// Solidity: function addStrategiesToDepositWhitelist(address[] strategiesToWhitelist) returns()
func (_StrategyManager *StrategyManagerSession) AddStrategiesToDepositWhitelist(strategiesToWhitelist []common.Address) (*types.Transaction, error) {
	return _StrategyManager.Contract.AddStrategiesToDepositWhitelist(&_StrategyManager.TransactOpts, strategiesToWhitelist)
}

// AddStrategiesToDepositWhitelist is a paid mutator transaction binding the contract method 0x5de08ff2.
//
// Solidity: function addStrategiesToDepositWhitelist(address[] strategiesToWhitelist) returns()
func (_StrategyManager *StrategyManagerTransactorSession) AddStrategiesToDepositWhitelist(strategiesToWhitelist []common.Address) (*types.Transaction, error) {
	return _StrategyManager.Contract.AddStrategiesToDepositWhitelist(&_StrategyManager.TransactOpts, strategiesToWhitelist)
}

// BurnShares is a paid mutator transaction binding the contract method 0x4b6d5d6e.
//
// Solidity: function burnShares(address strategy) returns()
func (_StrategyManager *StrategyManagerTransactor) BurnShares(opts *bind.TransactOpts, strategy common.Address) (*types.Transaction, error) {
	return _StrategyManager.contract.Transact(opts, "burnShares", strategy)
}

// BurnShares is a paid mutator transaction binding the contract method 0x4b6d5d6e.
//
// Solidity: function burnShares(address strategy) returns()
func (_StrategyManager *StrategyManagerSession) BurnShares(strategy common.Address) (*types.Transaction, error) {
	return _StrategyManager.Contract.BurnShares(&_StrategyManager.TransactOpts, strategy)
}

// BurnShares is a paid mutator transaction binding the contract method 0x4b6d5d6e.
//
// Solidity: function burnShares(address strategy) returns()
func (_StrategyManager *StrategyManagerTransactorSession) BurnShares(strategy common.Address) (*types.Transaction, error) {
	return _StrategyManager.Contract.BurnShares(&_StrategyManager.TransactOpts, strategy)
}

// ClearBurnOrRedistributableShares is a paid mutator transaction binding the contract method 0x88c10299.
//
// Solidity: function clearBurnOrRedistributableShares((address,uint32) operatorSet, uint256 slashId) returns(uint256[])
func (_StrategyManager *StrategyManagerTransactor) ClearBurnOrRedistributableShares(opts *bind.TransactOpts, operatorSet OperatorSet, slashId *big.Int) (*types.Transaction, error) {
	return _StrategyManager.contract.Transact(opts, "clearBurnOrRedistributableShares", operatorSet, slashId)
}

// ClearBurnOrRedistributableShares is a paid mutator transaction binding the contract method 0x88c10299.
//
// Solidity: function clearBurnOrRedistributableShares((address,uint32) operatorSet, uint256 slashId) returns(uint256[])
func (_StrategyManager *StrategyManagerSession) ClearBurnOrRedistributableShares(operatorSet OperatorSet, slashId *big.Int) (*types.Transaction, error) {
	return _StrategyManager.Contract.ClearBurnOrRedistributableShares(&_StrategyManager.TransactOpts, operatorSet, slashId)
}

// ClearBurnOrRedistributableShares is a paid mutator transaction binding the contract method 0x88c10299.
//
// Solidity: function clearBurnOrRedistributableShares((address,uint32) operatorSet, uint256 slashId) returns(uint256[])
func (_StrategyManager *StrategyManagerTransactorSession) ClearBurnOrRedistributableShares(operatorSet OperatorSet, slashId *big.Int) (*types.Transaction, error) {
	return _StrategyManager.Contract.ClearBurnOrRedistributableShares(&_StrategyManager.TransactOpts, operatorSet, slashId)
}

// ClearBurnOrRedistributableSharesByStrategy is a paid mutator transaction binding the contract method 0x2d44def6.
//
// Solidity: function clearBurnOrRedistributableSharesByStrategy((address,uint32) operatorSet, uint256 slashId, address strategy) returns(uint256)
func (_StrategyManager *StrategyManagerTransactor) ClearBurnOrRedistributableSharesByStrategy(opts *bind.TransactOpts, operatorSet OperatorSet, slashId *big.Int, strategy common.Address) (*types.Transaction, error) {
	return _StrategyManager.contract.Transact(opts, "clearBurnOrRedistributableSharesByStrategy", operatorSet, slashId, strategy)
}

// ClearBurnOrRedistributableSharesByStrategy is a paid mutator transaction binding the contract method 0x2d44def6.
//
// Solidity: function clearBurnOrRedistributableSharesByStrategy((address,uint32) operatorSet, uint256 slashId, address strategy) returns(uint256)
func (_StrategyManager *StrategyManagerSession) ClearBurnOrRedistributableSharesByStrategy(operatorSet OperatorSet, slashId *big.Int, strategy common.Address) (*types.Transaction, error) {
	return _StrategyManager.Contract.ClearBurnOrRedistributableSharesByStrategy(&_StrategyManager.TransactOpts, operatorSet, slashId, strategy)
}

// ClearBurnOrRedistributableSharesByStrategy is a paid mutator transaction binding the contract method 0x2d44def6.
//
// Solidity: function clearBurnOrRedistributableSharesByStrategy((address,uint32) operatorSet, uint256 slashId, address strategy) returns(uint256)
func (_StrategyManager *StrategyManagerTransactorSession) ClearBurnOrRedistributableSharesByStrategy(operatorSet OperatorSet, slashId *big.Int, strategy common.Address) (*types.Transaction, error) {
	return _StrategyManager.Contract.ClearBurnOrRedistributableSharesByStrategy(&_StrategyManager.TransactOpts, operatorSet, slashId, strategy)
}

// DepositIntoStrategy is a paid mutator transaction binding the contract method 0xe7a050aa.
//
// Solidity: function depositIntoStrategy(address strategy, address token, uint256 amount) returns(uint256 depositShares)
func (_StrategyManager *StrategyManagerTransactor) DepositIntoStrategy(opts *bind.TransactOpts, strategy common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StrategyManager.contract.Transact(opts, "depositIntoStrategy", strategy, token, amount)
}

// DepositIntoStrategy is a paid mutator transaction binding the contract method 0xe7a050aa.
//
// Solidity: function depositIntoStrategy(address strategy, address token, uint256 amount) returns(uint256 depositShares)
func (_StrategyManager *StrategyManagerSession) DepositIntoStrategy(strategy common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StrategyManager.Contract.DepositIntoStrategy(&_StrategyManager.TransactOpts, strategy, token, amount)
}

// DepositIntoStrategy is a paid mutator transaction binding the contract method 0xe7a050aa.
//
// Solidity: function depositIntoStrategy(address strategy, address token, uint256 amount) returns(uint256 depositShares)
func (_StrategyManager *StrategyManagerTransactorSession) DepositIntoStrategy(strategy common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StrategyManager.Contract.DepositIntoStrategy(&_StrategyManager.TransactOpts, strategy, token, amount)
}

// DepositIntoStrategyWithSignature is a paid mutator transaction binding the contract method 0x32e89ace.
//
// Solidity: function depositIntoStrategyWithSignature(address strategy, address token, uint256 amount, address staker, uint256 expiry, bytes signature) returns(uint256 depositShares)
func (_StrategyManager *StrategyManagerTransactor) DepositIntoStrategyWithSignature(opts *bind.TransactOpts, strategy common.Address, token common.Address, amount *big.Int, staker common.Address, expiry *big.Int, signature []byte) (*types.Transaction, error) {
	return _StrategyManager.contract.Transact(opts, "depositIntoStrategyWithSignature", strategy, token, amount, staker, expiry, signature)
}

// DepositIntoStrategyWithSignature is a paid mutator transaction binding the contract method 0x32e89ace.
//
// Solidity: function depositIntoStrategyWithSignature(address strategy, address token, uint256 amount, address staker, uint256 expiry, bytes signature) returns(uint256 depositShares)
func (_StrategyManager *StrategyManagerSession) DepositIntoStrategyWithSignature(strategy common.Address, token common.Address, amount *big.Int, staker common.Address, expiry *big.Int, signature []byte) (*types.Transaction, error) {
	return _StrategyManager.Contract.DepositIntoStrategyWithSignature(&_StrategyManager.TransactOpts, strategy, token, amount, staker, expiry, signature)
}

// DepositIntoStrategyWithSignature is a paid mutator transaction binding the contract method 0x32e89ace.
//
// Solidity: function depositIntoStrategyWithSignature(address strategy, address token, uint256 amount, address staker, uint256 expiry, bytes signature) returns(uint256 depositShares)
func (_StrategyManager *StrategyManagerTransactorSession) DepositIntoStrategyWithSignature(strategy common.Address, token common.Address, amount *big.Int, staker common.Address, expiry *big.Int, signature []byte) (*types.Transaction, error) {
	return _StrategyManager.Contract.DepositIntoStrategyWithSignature(&_StrategyManager.TransactOpts, strategy, token, amount, staker, expiry, signature)
}

// IncreaseBurnOrRedistributableShares is a paid mutator transaction binding the contract method 0x3fb99ca5.
//
// Solidity: function increaseBurnOrRedistributableShares((address,uint32) operatorSet, uint256 slashId, address strategy, uint256 sharesToBurn) returns()
func (_StrategyManager *StrategyManagerTransactor) IncreaseBurnOrRedistributableShares(opts *bind.TransactOpts, operatorSet OperatorSet, slashId *big.Int, strategy common.Address, sharesToBurn *big.Int) (*types.Transaction, error) {
	return _StrategyManager.contract.Transact(opts, "increaseBurnOrRedistributableShares", operatorSet, slashId, strategy, sharesToBurn)
}

// IncreaseBurnOrRedistributableShares is a paid mutator transaction binding the contract method 0x3fb99ca5.
//
// Solidity: function increaseBurnOrRedistributableShares((address,uint32) operatorSet, uint256 slashId, address strategy, uint256 sharesToBurn) returns()
func (_StrategyManager *StrategyManagerSession) IncreaseBurnOrRedistributableShares(operatorSet OperatorSet, slashId *big.Int, strategy common.Address, sharesToBurn *big.Int) (*types.Transaction, error) {
	return _StrategyManager.Contract.IncreaseBurnOrRedistributableShares(&_StrategyManager.TransactOpts, operatorSet, slashId, strategy, sharesToBurn)
}

// IncreaseBurnOrRedistributableShares is a paid mutator transaction binding the contract method 0x3fb99ca5.
//
// Solidity: function increaseBurnOrRedistributableShares((address,uint32) operatorSet, uint256 slashId, address strategy, uint256 sharesToBurn) returns()
func (_StrategyManager *StrategyManagerTransactorSession) IncreaseBurnOrRedistributableShares(operatorSet OperatorSet, slashId *big.Int, strategy common.Address, sharesToBurn *big.Int) (*types.Transaction, error) {
	return _StrategyManager.Contract.IncreaseBurnOrRedistributableShares(&_StrategyManager.TransactOpts, operatorSet, slashId, strategy, sharesToBurn)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address initialOwner, address initialStrategyWhitelister, uint256 initialPausedStatus) returns()
func (_StrategyManager *StrategyManagerTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, initialStrategyWhitelister common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _StrategyManager.contract.Transact(opts, "initialize", initialOwner, initialStrategyWhitelister, initialPausedStatus)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address initialOwner, address initialStrategyWhitelister, uint256 initialPausedStatus) returns()
func (_StrategyManager *StrategyManagerSession) Initialize(initialOwner common.Address, initialStrategyWhitelister common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _StrategyManager.Contract.Initialize(&_StrategyManager.TransactOpts, initialOwner, initialStrategyWhitelister, initialPausedStatus)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address initialOwner, address initialStrategyWhitelister, uint256 initialPausedStatus) returns()
func (_StrategyManager *StrategyManagerTransactorSession) Initialize(initialOwner common.Address, initialStrategyWhitelister common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _StrategyManager.Contract.Initialize(&_StrategyManager.TransactOpts, initialOwner, initialStrategyWhitelister, initialPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_StrategyManager *StrategyManagerTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _StrategyManager.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_StrategyManager *StrategyManagerSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _StrategyManager.Contract.Pause(&_StrategyManager.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_StrategyManager *StrategyManagerTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _StrategyManager.Contract.Pause(&_StrategyManager.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_StrategyManager *StrategyManagerTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrategyManager.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_StrategyManager *StrategyManagerSession) PauseAll() (*types.Transaction, error) {
	return _StrategyManager.Contract.PauseAll(&_StrategyManager.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_StrategyManager *StrategyManagerTransactorSession) PauseAll() (*types.Transaction, error) {
	return _StrategyManager.Contract.PauseAll(&_StrategyManager.TransactOpts)
}

// RemoveDepositShares is a paid mutator transaction binding the contract method 0x724af423.
//
// Solidity: function removeDepositShares(address staker, address strategy, uint256 depositSharesToRemove) returns(uint256)
func (_StrategyManager *StrategyManagerTransactor) RemoveDepositShares(opts *bind.TransactOpts, staker common.Address, strategy common.Address, depositSharesToRemove *big.Int) (*types.Transaction, error) {
	return _StrategyManager.contract.Transact(opts, "removeDepositShares", staker, strategy, depositSharesToRemove)
}

// RemoveDepositShares is a paid mutator transaction binding the contract method 0x724af423.
//
// Solidity: function removeDepositShares(address staker, address strategy, uint256 depositSharesToRemove) returns(uint256)
func (_StrategyManager *StrategyManagerSession) RemoveDepositShares(staker common.Address, strategy common.Address, depositSharesToRemove *big.Int) (*types.Transaction, error) {
	return _StrategyManager.Contract.RemoveDepositShares(&_StrategyManager.TransactOpts, staker, strategy, depositSharesToRemove)
}

// RemoveDepositShares is a paid mutator transaction binding the contract method 0x724af423.
//
// Solidity: function removeDepositShares(address staker, address strategy, uint256 depositSharesToRemove) returns(uint256)
func (_StrategyManager *StrategyManagerTransactorSession) RemoveDepositShares(staker common.Address, strategy common.Address, depositSharesToRemove *big.Int) (*types.Transaction, error) {
	return _StrategyManager.Contract.RemoveDepositShares(&_StrategyManager.TransactOpts, staker, strategy, depositSharesToRemove)
}

// RemoveStrategiesFromDepositWhitelist is a paid mutator transaction binding the contract method 0xb5d8b5b8.
//
// Solidity: function removeStrategiesFromDepositWhitelist(address[] strategiesToRemoveFromWhitelist) returns()
func (_StrategyManager *StrategyManagerTransactor) RemoveStrategiesFromDepositWhitelist(opts *bind.TransactOpts, strategiesToRemoveFromWhitelist []common.Address) (*types.Transaction, error) {
	return _StrategyManager.contract.Transact(opts, "removeStrategiesFromDepositWhitelist", strategiesToRemoveFromWhitelist)
}

// RemoveStrategiesFromDepositWhitelist is a paid mutator transaction binding the contract method 0xb5d8b5b8.
//
// Solidity: function removeStrategiesFromDepositWhitelist(address[] strategiesToRemoveFromWhitelist) returns()
func (_StrategyManager *StrategyManagerSession) RemoveStrategiesFromDepositWhitelist(strategiesToRemoveFromWhitelist []common.Address) (*types.Transaction, error) {
	return _StrategyManager.Contract.RemoveStrategiesFromDepositWhitelist(&_StrategyManager.TransactOpts, strategiesToRemoveFromWhitelist)
}

// RemoveStrategiesFromDepositWhitelist is a paid mutator transaction binding the contract method 0xb5d8b5b8.
//
// Solidity: function removeStrategiesFromDepositWhitelist(address[] strategiesToRemoveFromWhitelist) returns()
func (_StrategyManager *StrategyManagerTransactorSession) RemoveStrategiesFromDepositWhitelist(strategiesToRemoveFromWhitelist []common.Address) (*types.Transaction, error) {
	return _StrategyManager.Contract.RemoveStrategiesFromDepositWhitelist(&_StrategyManager.TransactOpts, strategiesToRemoveFromWhitelist)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StrategyManager *StrategyManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrategyManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StrategyManager *StrategyManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _StrategyManager.Contract.RenounceOwnership(&_StrategyManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StrategyManager *StrategyManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _StrategyManager.Contract.RenounceOwnership(&_StrategyManager.TransactOpts)
}

// SetSlashResolutionDelayBlocks is a paid mutator transaction binding the contract method 0x67f1b47e.
//
// Solidity: function setSlashResolutionDelayBlocks(uint32 newSlashResolutionDelayBlocks) returns()
func (_StrategyManager *StrategyManagerTransactor) SetSlashResolutionDelayBlocks(opts *bind.TransactOpts, newSlashResolutionDelayBlocks uint32) (*types.Transaction, error) {
	return _StrategyManager.contract.Transact(opts, "setSlashResolutionDelayBlocks", newSlashResolutionDelayBlocks)
}

// SetSlashResolutionDelayBlocks is a paid mutator transaction binding the contract method 0x67f1b47e.
//
// Solidity: function setSlashResolutionDelayBlocks(uint32 newSlashResolutionDelayBlocks) returns()
func (_StrategyManager *StrategyManagerSession) SetSlashResolutionDelayBlocks(newSlashResolutionDelayBlocks uint32) (*types.Transaction, error) {
	return _StrategyManager.Contract.SetSlashResolutionDelayBlocks(&_StrategyManager.TransactOpts, newSlashResolutionDelayBlocks)
}

// SetSlashResolutionDelayBlocks is a paid mutator transaction binding the contract method 0x67f1b47e.
//
// Solidity: function setSlashResolutionDelayBlocks(uint32 newSlashResolutionDelayBlocks) returns()
func (_StrategyManager *StrategyManagerTransactorSession) SetSlashResolutionDelayBlocks(newSlashResolutionDelayBlocks uint32) (*types.Transaction, error) {
	return _StrategyManager.Contract.SetSlashResolutionDelayBlocks(&_StrategyManager.TransactOpts, newSlashResolutionDelayBlocks)
}

// SetSlashResolutionDelaySetter is a paid mutator transaction binding the contract method 0x29acc842.
//
// Solidity: function setSlashResolutionDelaySetter(address newSlashResolutionDelaySetter) returns()
func (_StrategyManager *StrategyManagerTransactor) SetSlashResolutionDelaySetter(opts *bind.TransactOpts, newSlashResolutionDelaySetter common.Address) (*types.Transaction, error) {
	return _StrategyManager.contract.Transact(opts, "setSlashResolutionDelaySetter", newSlashResolutionDelaySetter)
}

// SetSlashResolutionDelaySetter is a paid mutator transaction binding the contract method 0x29acc842.
//
// Solidity: function setSlashResolutionDelaySetter(address newSlashResolutionDelaySetter) returns()
func (_StrategyManager *StrategyManagerSession) SetSlashResolutionDelaySetter(newSlashResolutionDelaySetter common.Address) (*types.Transaction, error) {
	return _StrategyManager.Contract.SetSlashResolutionDelaySetter(&_StrategyManager.TransactOpts, newSlashResolutionDelaySetter)
}

// SetSlashResolutionDelaySetter is a paid mutator transaction binding the contract method 0x29acc842.
//
// Solidity: function setSlashResolutionDelaySetter(address newSlashResolutionDelaySetter) returns()
func (_StrategyManager *StrategyManagerTransactorSession) SetSlashResolutionDelaySetter(newSlashResolutionDelaySetter common.Address) (*types.Transaction, error) {
	return _StrategyManager.Contract.SetSlashResolutionDelaySetter(&_StrategyManager.TransactOpts, newSlashResolutionDelaySetter)
}

// SetStrategyWhitelister is a paid mutator transaction binding the contract method 0xc6656702.
//
// Solidity: function setStrategyWhitelister(address newStrategyWhitelister) returns()
func (_StrategyManager *StrategyManagerTransactor) SetStrategyWhitelister(opts *bind.TransactOpts, newStrategyWhitelister common.Address) (*types.Transaction, error) {
	return _StrategyManager.contract.Transact(opts, "setStrategyWhitelister", newStrategyWhitelister)
}

// SetStrategyWhitelister is a paid mutator transaction binding the contract method 0xc6656702.
//
// Solidity: function setStrategyWhitelister(address newStrategyWhitelister) returns()
func (_StrategyManager *StrategyManagerSession) SetStrategyWhitelister(newStrategyWhitelister common.Address) (*types.Transaction, error) {
	return _StrategyManager.Contract.SetStrategyWhitelister(&_StrategyManager.TransactOpts, newStrategyWhitelister)
}

// SetStrategyWhitelister is a paid mutator transaction binding the contract method 0xc6656702.
//
// Solidity: function setStrategyWhitelister(address newStrategyWhitelister) returns()
func (_StrategyManager *StrategyManagerTransactorSession) SetStrategyWhitelister(newStrategyWhitelister common.Address) (*types.Transaction, error) {
	return _StrategyManager.Contract.SetStrategyWhitelister(&_StrategyManager.TransactOpts, newStrategyWhitelister)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StrategyManager *StrategyManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _StrategyManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StrategyManager *StrategyManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StrategyManager.Contract.TransferOwnership(&_StrategyManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StrategyManager *StrategyManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StrategyManager.Contract.TransferOwnership(&_StrategyManager.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_StrategyManager *StrategyManagerTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _StrategyManager.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_StrategyManager *StrategyManagerSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _StrategyManager.Contract.Unpause(&_StrategyManager.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_StrategyManager *StrategyManagerTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _StrategyManager.Contract.Unpause(&_StrategyManager.TransactOpts, newPausedStatus)
}

// WithdrawSharesAsTokens is a paid mutator transaction binding the contract method 0x2eae418c.
//
// Solidity: function withdrawSharesAsTokens(address staker, address strategy, address token, uint256 shares) returns()
func (_StrategyManager *StrategyManagerTransactor) WithdrawSharesAsTokens(opts *bind.TransactOpts, staker common.Address, strategy common.Address, token common.Address, shares *big.Int) (*types.Transaction, error) {
	return _StrategyManager.contract.Transact(opts, "withdrawSharesAsTokens", staker, strategy, token, shares)
}

// WithdrawSharesAsTokens is a paid mutator transaction binding the contract method 0x2eae418c.
//
// Solidity: function withdrawSharesAsTokens(address staker, address strategy, address token, uint256 shares) returns()
func (_StrategyManager *StrategyManagerSession) WithdrawSharesAsTokens(staker common.Address, strategy common.Address, token common.Address, shares *big.Int) (*types.Transaction, error) {
	return _StrategyManager.Contract.WithdrawSharesAsTokens(&_StrategyManager.TransactOpts, staker, strategy, token, shares)
}

// WithdrawSharesAsTokens is a paid mutator transaction binding the contract method 0x2eae418c.
//
// Solidity: function withdrawSharesAsTokens(address staker, address strategy, address token, uint256 shares) returns()
func (_StrategyManager *StrategyManagerTransactorSession) WithdrawSharesAsTokens(staker common.Address, strategy common.Address, token common.Address, shares *big.Int) (*types.Transaction, error) {
	return _StrategyManager.Contract.WithdrawSharesAsTokens(&_StrategyManager.TransactOpts, staker, strategy, token, shares)
}

// StrategyManagerBurnOrRedistributableSharesDecreasedIterator is returned from FilterBurnOrRedistributableSharesDecreased and is used to iterate over the raw logs and unpacked data for BurnOrRedistributableSharesDecreased events raised by the StrategyManager contract.
type StrategyManagerBurnOrRedistributableSharesDecreasedIterator struct {
	Event *StrategyManagerBurnOrRedistributableSharesDecreased // Event containing the contract specifics and raw log

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
func (it *StrategyManagerBurnOrRedistributableSharesDecreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyManagerBurnOrRedistributableSharesDecreased)
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
		it.Event = new(StrategyManagerBurnOrRedistributableSharesDecreased)
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
func (it *StrategyManagerBurnOrRedistributableSharesDecreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyManagerBurnOrRedistributableSharesDecreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyManagerBurnOrRedistributableSharesDecreased represents a BurnOrRedistributableSharesDecreased event raised by the StrategyManager contract.
type StrategyManagerBurnOrRedistributableSharesDecreased struct {
	OperatorSet OperatorSet
	SlashId     *big.Int
	Strategy    common.Address
	Shares      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBurnOrRedistributableSharesDecreased is a free log retrieval operation binding the contract event 0xe6413aa0c789e437b0a06bf64b20926584f066c79a2d8b80a759c85472f7b0af.
//
// Solidity: event BurnOrRedistributableSharesDecreased((address,uint32) operatorSet, uint256 slashId, address strategy, uint256 shares)
func (_StrategyManager *StrategyManagerFilterer) FilterBurnOrRedistributableSharesDecreased(opts *bind.FilterOpts) (*StrategyManagerBurnOrRedistributableSharesDecreasedIterator, error) {

	logs, sub, err := _StrategyManager.contract.FilterLogs(opts, "BurnOrRedistributableSharesDecreased")
	if err != nil {
		return nil, err
	}
	return &StrategyManagerBurnOrRedistributableSharesDecreasedIterator{contract: _StrategyManager.contract, event: "BurnOrRedistributableSharesDecreased", logs: logs, sub: sub}, nil
}

// WatchBurnOrRedistributableSharesDecreased is a free log subscription operation binding the contract event 0xe6413aa0c789e437b0a06bf64b20926584f066c79a2d8b80a759c85472f7b0af.
//
// Solidity: event BurnOrRedistributableSharesDecreased((address,uint32) operatorSet, uint256 slashId, address strategy, uint256 shares)
func (_StrategyManager *StrategyManagerFilterer) WatchBurnOrRedistributableSharesDecreased(opts *bind.WatchOpts, sink chan<- *StrategyManagerBurnOrRedistributableSharesDecreased) (event.Subscription, error) {

	logs, sub, err := _StrategyManager.contract.WatchLogs(opts, "BurnOrRedistributableSharesDecreased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyManagerBurnOrRedistributableSharesDecreased)
				if err := _StrategyManager.contract.UnpackLog(event, "BurnOrRedistributableSharesDecreased", log); err != nil {
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

// ParseBurnOrRedistributableSharesDecreased is a log parse operation binding the contract event 0xe6413aa0c789e437b0a06bf64b20926584f066c79a2d8b80a759c85472f7b0af.
//
// Solidity: event BurnOrRedistributableSharesDecreased((address,uint32) operatorSet, uint256 slashId, address strategy, uint256 shares)
func (_StrategyManager *StrategyManagerFilterer) ParseBurnOrRedistributableSharesDecreased(log types.Log) (*StrategyManagerBurnOrRedistributableSharesDecreased, error) {
	event := new(StrategyManagerBurnOrRedistributableSharesDecreased)
	if err := _StrategyManager.contract.UnpackLog(event, "BurnOrRedistributableSharesDecreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyManagerBurnOrRedistributableSharesIncreasedIterator is returned from FilterBurnOrRedistributableSharesIncreased and is used to iterate over the raw logs and unpacked data for BurnOrRedistributableSharesIncreased events raised by the StrategyManager contract.
type StrategyManagerBurnOrRedistributableSharesIncreasedIterator struct {
	Event *StrategyManagerBurnOrRedistributableSharesIncreased // Event containing the contract specifics and raw log

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
func (it *StrategyManagerBurnOrRedistributableSharesIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyManagerBurnOrRedistributableSharesIncreased)
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
		it.Event = new(StrategyManagerBurnOrRedistributableSharesIncreased)
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
func (it *StrategyManagerBurnOrRedistributableSharesIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyManagerBurnOrRedistributableSharesIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyManagerBurnOrRedistributableSharesIncreased represents a BurnOrRedistributableSharesIncreased event raised by the StrategyManager contract.
type StrategyManagerBurnOrRedistributableSharesIncreased struct {
	OperatorSet OperatorSet
	SlashId     *big.Int
	Strategy    common.Address
	Shares      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBurnOrRedistributableSharesIncreased is a free log retrieval operation binding the contract event 0x5f5209798bbac45a16d2dc3bc67319fab26ee00153916d6f07b69f8a134a1e8b.
//
// Solidity: event BurnOrRedistributableSharesIncreased((address,uint32) operatorSet, uint256 slashId, address strategy, uint256 shares)
func (_StrategyManager *StrategyManagerFilterer) FilterBurnOrRedistributableSharesIncreased(opts *bind.FilterOpts) (*StrategyManagerBurnOrRedistributableSharesIncreasedIterator, error) {

	logs, sub, err := _StrategyManager.contract.FilterLogs(opts, "BurnOrRedistributableSharesIncreased")
	if err != nil {
		return nil, err
	}
	return &StrategyManagerBurnOrRedistributableSharesIncreasedIterator{contract: _StrategyManager.contract, event: "BurnOrRedistributableSharesIncreased", logs: logs, sub: sub}, nil
}

// WatchBurnOrRedistributableSharesIncreased is a free log subscription operation binding the contract event 0x5f5209798bbac45a16d2dc3bc67319fab26ee00153916d6f07b69f8a134a1e8b.
//
// Solidity: event BurnOrRedistributableSharesIncreased((address,uint32) operatorSet, uint256 slashId, address strategy, uint256 shares)
func (_StrategyManager *StrategyManagerFilterer) WatchBurnOrRedistributableSharesIncreased(opts *bind.WatchOpts, sink chan<- *StrategyManagerBurnOrRedistributableSharesIncreased) (event.Subscription, error) {

	logs, sub, err := _StrategyManager.contract.WatchLogs(opts, "BurnOrRedistributableSharesIncreased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyManagerBurnOrRedistributableSharesIncreased)
				if err := _StrategyManager.contract.UnpackLog(event, "BurnOrRedistributableSharesIncreased", log); err != nil {
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

// ParseBurnOrRedistributableSharesIncreased is a log parse operation binding the contract event 0x5f5209798bbac45a16d2dc3bc67319fab26ee00153916d6f07b69f8a134a1e8b.
//
// Solidity: event BurnOrRedistributableSharesIncreased((address,uint32) operatorSet, uint256 slashId, address strategy, uint256 shares)
func (_StrategyManager *StrategyManagerFilterer) ParseBurnOrRedistributableSharesIncreased(log types.Log) (*StrategyManagerBurnOrRedistributableSharesIncreased, error) {
	event := new(StrategyManagerBurnOrRedistributableSharesIncreased)
	if err := _StrategyManager.contract.UnpackLog(event, "BurnOrRedistributableSharesIncreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyManagerBurnableSharesDecreasedIterator is returned from FilterBurnableSharesDecreased and is used to iterate over the raw logs and unpacked data for BurnableSharesDecreased events raised by the StrategyManager contract.
type StrategyManagerBurnableSharesDecreasedIterator struct {
	Event *StrategyManagerBurnableSharesDecreased // Event containing the contract specifics and raw log

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
func (it *StrategyManagerBurnableSharesDecreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyManagerBurnableSharesDecreased)
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
		it.Event = new(StrategyManagerBurnableSharesDecreased)
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
func (it *StrategyManagerBurnableSharesDecreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyManagerBurnableSharesDecreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyManagerBurnableSharesDecreased represents a BurnableSharesDecreased event raised by the StrategyManager contract.
type StrategyManagerBurnableSharesDecreased struct {
	Strategy common.Address
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBurnableSharesDecreased is a free log retrieval operation binding the contract event 0xd9d082c3ec4f3a3ffa55c324939a06407f5fbcb87d5e0ce3b9508c92c84ed839.
//
// Solidity: event BurnableSharesDecreased(address strategy, uint256 shares)
func (_StrategyManager *StrategyManagerFilterer) FilterBurnableSharesDecreased(opts *bind.FilterOpts) (*StrategyManagerBurnableSharesDecreasedIterator, error) {

	logs, sub, err := _StrategyManager.contract.FilterLogs(opts, "BurnableSharesDecreased")
	if err != nil {
		return nil, err
	}
	return &StrategyManagerBurnableSharesDecreasedIterator{contract: _StrategyManager.contract, event: "BurnableSharesDecreased", logs: logs, sub: sub}, nil
}

// WatchBurnableSharesDecreased is a free log subscription operation binding the contract event 0xd9d082c3ec4f3a3ffa55c324939a06407f5fbcb87d5e0ce3b9508c92c84ed839.
//
// Solidity: event BurnableSharesDecreased(address strategy, uint256 shares)
func (_StrategyManager *StrategyManagerFilterer) WatchBurnableSharesDecreased(opts *bind.WatchOpts, sink chan<- *StrategyManagerBurnableSharesDecreased) (event.Subscription, error) {

	logs, sub, err := _StrategyManager.contract.WatchLogs(opts, "BurnableSharesDecreased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyManagerBurnableSharesDecreased)
				if err := _StrategyManager.contract.UnpackLog(event, "BurnableSharesDecreased", log); err != nil {
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

// ParseBurnableSharesDecreased is a log parse operation binding the contract event 0xd9d082c3ec4f3a3ffa55c324939a06407f5fbcb87d5e0ce3b9508c92c84ed839.
//
// Solidity: event BurnableSharesDecreased(address strategy, uint256 shares)
func (_StrategyManager *StrategyManagerFilterer) ParseBurnableSharesDecreased(log types.Log) (*StrategyManagerBurnableSharesDecreased, error) {
	event := new(StrategyManagerBurnableSharesDecreased)
	if err := _StrategyManager.contract.UnpackLog(event, "BurnableSharesDecreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyManagerDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the StrategyManager contract.
type StrategyManagerDepositIterator struct {
	Event *StrategyManagerDeposit // Event containing the contract specifics and raw log

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
func (it *StrategyManagerDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyManagerDeposit)
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
		it.Event = new(StrategyManagerDeposit)
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
func (it *StrategyManagerDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyManagerDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyManagerDeposit represents a Deposit event raised by the StrategyManager contract.
type StrategyManagerDeposit struct {
	Staker   common.Address
	Strategy common.Address
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: event Deposit(address staker, address strategy, uint256 shares)
func (_StrategyManager *StrategyManagerFilterer) FilterDeposit(opts *bind.FilterOpts) (*StrategyManagerDepositIterator, error) {

	logs, sub, err := _StrategyManager.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &StrategyManagerDepositIterator{contract: _StrategyManager.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: event Deposit(address staker, address strategy, uint256 shares)
func (_StrategyManager *StrategyManagerFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *StrategyManagerDeposit) (event.Subscription, error) {

	logs, sub, err := _StrategyManager.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyManagerDeposit)
				if err := _StrategyManager.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: event Deposit(address staker, address strategy, uint256 shares)
func (_StrategyManager *StrategyManagerFilterer) ParseDeposit(log types.Log) (*StrategyManagerDeposit, error) {
	event := new(StrategyManagerDeposit)
	if err := _StrategyManager.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the StrategyManager contract.
type StrategyManagerInitializedIterator struct {
	Event *StrategyManagerInitialized // Event containing the contract specifics and raw log

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
func (it *StrategyManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyManagerInitialized)
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
		it.Event = new(StrategyManagerInitialized)
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
func (it *StrategyManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyManagerInitialized represents a Initialized event raised by the StrategyManager contract.
type StrategyManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StrategyManager *StrategyManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*StrategyManagerInitializedIterator, error) {

	logs, sub, err := _StrategyManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StrategyManagerInitializedIterator{contract: _StrategyManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StrategyManager *StrategyManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StrategyManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _StrategyManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyManagerInitialized)
				if err := _StrategyManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_StrategyManager *StrategyManagerFilterer) ParseInitialized(log types.Log) (*StrategyManagerInitialized, error) {
	event := new(StrategyManagerInitialized)
	if err := _StrategyManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the StrategyManager contract.
type StrategyManagerOwnershipTransferredIterator struct {
	Event *StrategyManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StrategyManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyManagerOwnershipTransferred)
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
		it.Event = new(StrategyManagerOwnershipTransferred)
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
func (it *StrategyManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyManagerOwnershipTransferred represents a OwnershipTransferred event raised by the StrategyManager contract.
type StrategyManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StrategyManager *StrategyManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StrategyManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StrategyManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StrategyManagerOwnershipTransferredIterator{contract: _StrategyManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StrategyManager *StrategyManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StrategyManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StrategyManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyManagerOwnershipTransferred)
				if err := _StrategyManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_StrategyManager *StrategyManagerFilterer) ParseOwnershipTransferred(log types.Log) (*StrategyManagerOwnershipTransferred, error) {
	event := new(StrategyManagerOwnershipTransferred)
	if err := _StrategyManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the StrategyManager contract.
type StrategyManagerPausedIterator struct {
	Event *StrategyManagerPaused // Event containing the contract specifics and raw log

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
func (it *StrategyManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyManagerPaused)
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
		it.Event = new(StrategyManagerPaused)
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
func (it *StrategyManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyManagerPaused represents a Paused event raised by the StrategyManager contract.
type StrategyManagerPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_StrategyManager *StrategyManagerFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*StrategyManagerPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StrategyManager.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &StrategyManagerPausedIterator{contract: _StrategyManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_StrategyManager *StrategyManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *StrategyManagerPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StrategyManager.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyManagerPaused)
				if err := _StrategyManager.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_StrategyManager *StrategyManagerFilterer) ParsePaused(log types.Log) (*StrategyManagerPaused, error) {
	event := new(StrategyManagerPaused)
	if err := _StrategyManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyManagerSlashResolutionBlockSetIterator is returned from FilterSlashResolutionBlockSet and is used to iterate over the raw logs and unpacked data for SlashResolutionBlockSet events raised by the StrategyManager contract.
type StrategyManagerSlashResolutionBlockSetIterator struct {
	Event *StrategyManagerSlashResolutionBlockSet // Event containing the contract specifics and raw log

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
func (it *StrategyManagerSlashResolutionBlockSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyManagerSlashResolutionBlockSet)
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
		it.Event = new(StrategyManagerSlashResolutionBlockSet)
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
func (it *StrategyManagerSlashResolutionBlockSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyManagerSlashResolutionBlockSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyManagerSlashResolutionBlockSet represents a SlashResolutionBlockSet event raised by the StrategyManager contract.
type StrategyManagerSlashResolutionBlockSet struct {
	OperatorSet     OperatorSet
	SlashId         *big.Int
	ResolutionBlock uint32
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSlashResolutionBlockSet is a free log retrieval operation binding the contract event 0xa8fd1473d1b9b7b4c9033ecb92a0f6e43f107aec7720ede3709521d122f89fa5.
//
// Solidity: event SlashResolutionBlockSet((address,uint32) operatorSet, uint256 slashId, uint32 resolutionBlock)
func (_StrategyManager *StrategyManagerFilterer) FilterSlashResolutionBlockSet(opts *bind.FilterOpts) (*StrategyManagerSlashResolutionBlockSetIterator, error) {

	logs, sub, err := _StrategyManager.contract.FilterLogs(opts, "SlashResolutionBlockSet")
	if err != nil {
		return nil, err
	}
	return &StrategyManagerSlashResolutionBlockSetIterator{contract: _StrategyManager.contract, event: "SlashResolutionBlockSet", logs: logs, sub: sub}, nil
}

// WatchSlashResolutionBlockSet is a free log subscription operation binding the contract event 0xa8fd1473d1b9b7b4c9033ecb92a0f6e43f107aec7720ede3709521d122f89fa5.
//
// Solidity: event SlashResolutionBlockSet((address,uint32) operatorSet, uint256 slashId, uint32 resolutionBlock)
func (_StrategyManager *StrategyManagerFilterer) WatchSlashResolutionBlockSet(opts *bind.WatchOpts, sink chan<- *StrategyManagerSlashResolutionBlockSet) (event.Subscription, error) {

	logs, sub, err := _StrategyManager.contract.WatchLogs(opts, "SlashResolutionBlockSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyManagerSlashResolutionBlockSet)
				if err := _StrategyManager.contract.UnpackLog(event, "SlashResolutionBlockSet", log); err != nil {
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

// ParseSlashResolutionBlockSet is a log parse operation binding the contract event 0xa8fd1473d1b9b7b4c9033ecb92a0f6e43f107aec7720ede3709521d122f89fa5.
//
// Solidity: event SlashResolutionBlockSet((address,uint32) operatorSet, uint256 slashId, uint32 resolutionBlock)
func (_StrategyManager *StrategyManagerFilterer) ParseSlashResolutionBlockSet(log types.Log) (*StrategyManagerSlashResolutionBlockSet, error) {
	event := new(StrategyManagerSlashResolutionBlockSet)
	if err := _StrategyManager.contract.UnpackLog(event, "SlashResolutionBlockSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyManagerSlashResolutionDelayBlocksChangedIterator is returned from FilterSlashResolutionDelayBlocksChanged and is used to iterate over the raw logs and unpacked data for SlashResolutionDelayBlocksChanged events raised by the StrategyManager contract.
type StrategyManagerSlashResolutionDelayBlocksChangedIterator struct {
	Event *StrategyManagerSlashResolutionDelayBlocksChanged // Event containing the contract specifics and raw log

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
func (it *StrategyManagerSlashResolutionDelayBlocksChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyManagerSlashResolutionDelayBlocksChanged)
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
		it.Event = new(StrategyManagerSlashResolutionDelayBlocksChanged)
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
func (it *StrategyManagerSlashResolutionDelayBlocksChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyManagerSlashResolutionDelayBlocksChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyManagerSlashResolutionDelayBlocksChanged represents a SlashResolutionDelayBlocksChanged event raised by the StrategyManager contract.
type StrategyManagerSlashResolutionDelayBlocksChanged struct {
	PreviousDelayBlocks uint32
	NewDelayBlocks      uint32
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSlashResolutionDelayBlocksChanged is a free log retrieval operation binding the contract event 0x086f17854c753dd15c19bb0e7683f699b07a44ea9673872d2b5a07079c7ab29a.
//
// Solidity: event SlashResolutionDelayBlocksChanged(uint32 previousDelayBlocks, uint32 newDelayBlocks)
func (_StrategyManager *StrategyManagerFilterer) FilterSlashResolutionDelayBlocksChanged(opts *bind.FilterOpts) (*StrategyManagerSlashResolutionDelayBlocksChangedIterator, error) {

	logs, sub, err := _StrategyManager.contract.FilterLogs(opts, "SlashResolutionDelayBlocksChanged")
	if err != nil {
		return nil, err
	}
	return &StrategyManagerSlashResolutionDelayBlocksChangedIterator{contract: _StrategyManager.contract, event: "SlashResolutionDelayBlocksChanged", logs: logs, sub: sub}, nil
}

// WatchSlashResolutionDelayBlocksChanged is a free log subscription operation binding the contract event 0x086f17854c753dd15c19bb0e7683f699b07a44ea9673872d2b5a07079c7ab29a.
//
// Solidity: event SlashResolutionDelayBlocksChanged(uint32 previousDelayBlocks, uint32 newDelayBlocks)
func (_StrategyManager *StrategyManagerFilterer) WatchSlashResolutionDelayBlocksChanged(opts *bind.WatchOpts, sink chan<- *StrategyManagerSlashResolutionDelayBlocksChanged) (event.Subscription, error) {

	logs, sub, err := _StrategyManager.contract.WatchLogs(opts, "SlashResolutionDelayBlocksChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyManagerSlashResolutionDelayBlocksChanged)
				if err := _StrategyManager.contract.UnpackLog(event, "SlashResolutionDelayBlocksChanged", log); err != nil {
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

// ParseSlashResolutionDelayBlocksChanged is a log parse operation binding the contract event 0x086f17854c753dd15c19bb0e7683f699b07a44ea9673872d2b5a07079c7ab29a.
//
// Solidity: event SlashResolutionDelayBlocksChanged(uint32 previousDelayBlocks, uint32 newDelayBlocks)
func (_StrategyManager *StrategyManagerFilterer) ParseSlashResolutionDelayBlocksChanged(log types.Log) (*StrategyManagerSlashResolutionDelayBlocksChanged, error) {
	event := new(StrategyManagerSlashResolutionDelayBlocksChanged)
	if err := _StrategyManager.contract.UnpackLog(event, "SlashResolutionDelayBlocksChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyManagerSlashResolutionDelaySetterChangedIterator is returned from FilterSlashResolutionDelaySetterChanged and is used to iterate over the raw logs and unpacked data for SlashResolutionDelaySetterChanged events raised by the StrategyManager contract.
type StrategyManagerSlashResolutionDelaySetterChangedIterator struct {
	Event *StrategyManagerSlashResolutionDelaySetterChanged // Event containing the contract specifics and raw log

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
func (it *StrategyManagerSlashResolutionDelaySetterChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyManagerSlashResolutionDelaySetterChanged)
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
		it.Event = new(StrategyManagerSlashResolutionDelaySetterChanged)
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
func (it *StrategyManagerSlashResolutionDelaySetterChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyManagerSlashResolutionDelaySetterChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyManagerSlashResolutionDelaySetterChanged represents a SlashResolutionDelaySetterChanged event raised by the StrategyManager contract.
type StrategyManagerSlashResolutionDelaySetterChanged struct {
	PreviousSetter common.Address
	NewSetter      common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSlashResolutionDelaySetterChanged is a free log retrieval operation binding the contract event 0x98355cb0f7088f4584521f1fbb1f7ffd33d427eb8740bc64997df3a6a54078a7.
//
// Solidity: event SlashResolutionDelaySetterChanged(address previousSetter, address newSetter)
func (_StrategyManager *StrategyManagerFilterer) FilterSlashResolutionDelaySetterChanged(opts *bind.FilterOpts) (*StrategyManagerSlashResolutionDelaySetterChangedIterator, error) {

	logs, sub, err := _StrategyManager.contract.FilterLogs(opts, "SlashResolutionDelaySetterChanged")
	if err != nil {
		return nil, err
	}
	return &StrategyManagerSlashResolutionDelaySetterChangedIterator{contract: _StrategyManager.contract, event: "SlashResolutionDelaySetterChanged", logs: logs, sub: sub}, nil
}

// WatchSlashResolutionDelaySetterChanged is a free log subscription operation binding the contract event 0x98355cb0f7088f4584521f1fbb1f7ffd33d427eb8740bc64997df3a6a54078a7.
//
// Solidity: event SlashResolutionDelaySetterChanged(address previousSetter, address newSetter)
func (_StrategyManager *StrategyManagerFilterer) WatchSlashResolutionDelaySetterChanged(opts *bind.WatchOpts, sink chan<- *StrategyManagerSlashResolutionDelaySetterChanged) (event.Subscription, error) {

	logs, sub, err := _StrategyManager.contract.WatchLogs(opts, "SlashResolutionDelaySetterChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyManagerSlashResolutionDelaySetterChanged)
				if err := _StrategyManager.contract.UnpackLog(event, "SlashResolutionDelaySetterChanged", log); err != nil {
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

// ParseSlashResolutionDelaySetterChanged is a log parse operation binding the contract event 0x98355cb0f7088f4584521f1fbb1f7ffd33d427eb8740bc64997df3a6a54078a7.
//
// Solidity: event SlashResolutionDelaySetterChanged(address previousSetter, address newSetter)
func (_StrategyManager *StrategyManagerFilterer) ParseSlashResolutionDelaySetterChanged(log types.Log) (*StrategyManagerSlashResolutionDelaySetterChanged, error) {
	event := new(StrategyManagerSlashResolutionDelaySetterChanged)
	if err := _StrategyManager.contract.UnpackLog(event, "SlashResolutionDelaySetterChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyManagerStrategyAddedToDepositWhitelistIterator is returned from FilterStrategyAddedToDepositWhitelist and is used to iterate over the raw logs and unpacked data for StrategyAddedToDepositWhitelist events raised by the StrategyManager contract.
type StrategyManagerStrategyAddedToDepositWhitelistIterator struct {
	Event *StrategyManagerStrategyAddedToDepositWhitelist // Event containing the contract specifics and raw log

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
func (it *StrategyManagerStrategyAddedToDepositWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyManagerStrategyAddedToDepositWhitelist)
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
		it.Event = new(StrategyManagerStrategyAddedToDepositWhitelist)
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
func (it *StrategyManagerStrategyAddedToDepositWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyManagerStrategyAddedToDepositWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyManagerStrategyAddedToDepositWhitelist represents a StrategyAddedToDepositWhitelist event raised by the StrategyManager contract.
type StrategyManagerStrategyAddedToDepositWhitelist struct {
	Strategy common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStrategyAddedToDepositWhitelist is a free log retrieval operation binding the contract event 0x0c35b17d91c96eb2751cd456e1252f42a386e524ef9ff26ecc9950859fdc04fe.
//
// Solidity: event StrategyAddedToDepositWhitelist(address strategy)
func (_StrategyManager *StrategyManagerFilterer) FilterStrategyAddedToDepositWhitelist(opts *bind.FilterOpts) (*StrategyManagerStrategyAddedToDepositWhitelistIterator, error) {

	logs, sub, err := _StrategyManager.contract.FilterLogs(opts, "StrategyAddedToDepositWhitelist")
	if err != nil {
		return nil, err
	}
	return &StrategyManagerStrategyAddedToDepositWhitelistIterator{contract: _StrategyManager.contract, event: "StrategyAddedToDepositWhitelist", logs: logs, sub: sub}, nil
}

// WatchStrategyAddedToDepositWhitelist is a free log subscription operation binding the contract event 0x0c35b17d91c96eb2751cd456e1252f42a386e524ef9ff26ecc9950859fdc04fe.
//
// Solidity: event StrategyAddedToDepositWhitelist(address strategy)
func (_StrategyManager *StrategyManagerFilterer) WatchStrategyAddedToDepositWhitelist(opts *bind.WatchOpts, sink chan<- *StrategyManagerStrategyAddedToDepositWhitelist) (event.Subscription, error) {

	logs, sub, err := _StrategyManager.contract.WatchLogs(opts, "StrategyAddedToDepositWhitelist")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyManagerStrategyAddedToDepositWhitelist)
				if err := _StrategyManager.contract.UnpackLog(event, "StrategyAddedToDepositWhitelist", log); err != nil {
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

// ParseStrategyAddedToDepositWhitelist is a log parse operation binding the contract event 0x0c35b17d91c96eb2751cd456e1252f42a386e524ef9ff26ecc9950859fdc04fe.
//
// Solidity: event StrategyAddedToDepositWhitelist(address strategy)
func (_StrategyManager *StrategyManagerFilterer) ParseStrategyAddedToDepositWhitelist(log types.Log) (*StrategyManagerStrategyAddedToDepositWhitelist, error) {
	event := new(StrategyManagerStrategyAddedToDepositWhitelist)
	if err := _StrategyManager.contract.UnpackLog(event, "StrategyAddedToDepositWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyManagerStrategyRemovedFromDepositWhitelistIterator is returned from FilterStrategyRemovedFromDepositWhitelist and is used to iterate over the raw logs and unpacked data for StrategyRemovedFromDepositWhitelist events raised by the StrategyManager contract.
type StrategyManagerStrategyRemovedFromDepositWhitelistIterator struct {
	Event *StrategyManagerStrategyRemovedFromDepositWhitelist // Event containing the contract specifics and raw log

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
func (it *StrategyManagerStrategyRemovedFromDepositWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyManagerStrategyRemovedFromDepositWhitelist)
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
		it.Event = new(StrategyManagerStrategyRemovedFromDepositWhitelist)
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
func (it *StrategyManagerStrategyRemovedFromDepositWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyManagerStrategyRemovedFromDepositWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyManagerStrategyRemovedFromDepositWhitelist represents a StrategyRemovedFromDepositWhitelist event raised by the StrategyManager contract.
type StrategyManagerStrategyRemovedFromDepositWhitelist struct {
	Strategy common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStrategyRemovedFromDepositWhitelist is a free log retrieval operation binding the contract event 0x4074413b4b443e4e58019f2855a8765113358c7c72e39509c6af45fc0f5ba030.
//
// Solidity: event StrategyRemovedFromDepositWhitelist(address strategy)
func (_StrategyManager *StrategyManagerFilterer) FilterStrategyRemovedFromDepositWhitelist(opts *bind.FilterOpts) (*StrategyManagerStrategyRemovedFromDepositWhitelistIterator, error) {

	logs, sub, err := _StrategyManager.contract.FilterLogs(opts, "StrategyRemovedFromDepositWhitelist")
	if err != nil {
		return nil, err
	}
	return &StrategyManagerStrategyRemovedFromDepositWhitelistIterator{contract: _StrategyManager.contract, event: "StrategyRemovedFromDepositWhitelist", logs: logs, sub: sub}, nil
}

// WatchStrategyRemovedFromDepositWhitelist is a free log subscription operation binding the contract event 0x4074413b4b443e4e58019f2855a8765113358c7c72e39509c6af45fc0f5ba030.
//
// Solidity: event StrategyRemovedFromDepositWhitelist(address strategy)
func (_StrategyManager *StrategyManagerFilterer) WatchStrategyRemovedFromDepositWhitelist(opts *bind.WatchOpts, sink chan<- *StrategyManagerStrategyRemovedFromDepositWhitelist) (event.Subscription, error) {

	logs, sub, err := _StrategyManager.contract.WatchLogs(opts, "StrategyRemovedFromDepositWhitelist")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyManagerStrategyRemovedFromDepositWhitelist)
				if err := _StrategyManager.contract.UnpackLog(event, "StrategyRemovedFromDepositWhitelist", log); err != nil {
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

// ParseStrategyRemovedFromDepositWhitelist is a log parse operation binding the contract event 0x4074413b4b443e4e58019f2855a8765113358c7c72e39509c6af45fc0f5ba030.
//
// Solidity: event StrategyRemovedFromDepositWhitelist(address strategy)
func (_StrategyManager *StrategyManagerFilterer) ParseStrategyRemovedFromDepositWhitelist(log types.Log) (*StrategyManagerStrategyRemovedFromDepositWhitelist, error) {
	event := new(StrategyManagerStrategyRemovedFromDepositWhitelist)
	if err := _StrategyManager.contract.UnpackLog(event, "StrategyRemovedFromDepositWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyManagerStrategyWhitelisterChangedIterator is returned from FilterStrategyWhitelisterChanged and is used to iterate over the raw logs and unpacked data for StrategyWhitelisterChanged events raised by the StrategyManager contract.
type StrategyManagerStrategyWhitelisterChangedIterator struct {
	Event *StrategyManagerStrategyWhitelisterChanged // Event containing the contract specifics and raw log

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
func (it *StrategyManagerStrategyWhitelisterChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyManagerStrategyWhitelisterChanged)
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
		it.Event = new(StrategyManagerStrategyWhitelisterChanged)
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
func (it *StrategyManagerStrategyWhitelisterChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyManagerStrategyWhitelisterChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyManagerStrategyWhitelisterChanged represents a StrategyWhitelisterChanged event raised by the StrategyManager contract.
type StrategyManagerStrategyWhitelisterChanged struct {
	PreviousAddress common.Address
	NewAddress      common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterStrategyWhitelisterChanged is a free log retrieval operation binding the contract event 0x4264275e593955ff9d6146a51a4525f6ddace2e81db9391abcc9d1ca48047d29.
//
// Solidity: event StrategyWhitelisterChanged(address previousAddress, address newAddress)
func (_StrategyManager *StrategyManagerFilterer) FilterStrategyWhitelisterChanged(opts *bind.FilterOpts) (*StrategyManagerStrategyWhitelisterChangedIterator, error) {

	logs, sub, err := _StrategyManager.contract.FilterLogs(opts, "StrategyWhitelisterChanged")
	if err != nil {
		return nil, err
	}
	return &StrategyManagerStrategyWhitelisterChangedIterator{contract: _StrategyManager.contract, event: "StrategyWhitelisterChanged", logs: logs, sub: sub}, nil
}

// WatchStrategyWhitelisterChanged is a free log subscription operation binding the contract event 0x4264275e593955ff9d6146a51a4525f6ddace2e81db9391abcc9d1ca48047d29.
//
// Solidity: event StrategyWhitelisterChanged(address previousAddress, address newAddress)
func (_StrategyManager *StrategyManagerFilterer) WatchStrategyWhitelisterChanged(opts *bind.WatchOpts, sink chan<- *StrategyManagerStrategyWhitelisterChanged) (event.Subscription, error) {

	logs, sub, err := _StrategyManager.contract.WatchLogs(opts, "StrategyWhitelisterChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyManagerStrategyWhitelisterChanged)
				if err := _StrategyManager.contract.UnpackLog(event, "StrategyWhitelisterChanged", log); err != nil {
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

// ParseStrategyWhitelisterChanged is a log parse operation binding the contract event 0x4264275e593955ff9d6146a51a4525f6ddace2e81db9391abcc9d1ca48047d29.
//
// Solidity: event StrategyWhitelisterChanged(address previousAddress, address newAddress)
func (_StrategyManager *StrategyManagerFilterer) ParseStrategyWhitelisterChanged(log types.Log) (*StrategyManagerStrategyWhitelisterChanged, error) {
	event := new(StrategyManagerStrategyWhitelisterChanged)
	if err := _StrategyManager.contract.UnpackLog(event, "StrategyWhitelisterChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the StrategyManager contract.
type StrategyManagerUnpausedIterator struct {
	Event *StrategyManagerUnpaused // Event containing the contract specifics and raw log

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
func (it *StrategyManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyManagerUnpaused)
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
		it.Event = new(StrategyManagerUnpaused)
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
func (it *StrategyManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyManagerUnpaused represents a Unpaused event raised by the StrategyManager contract.
type StrategyManagerUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_StrategyManager *StrategyManagerFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*StrategyManagerUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StrategyManager.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &StrategyManagerUnpausedIterator{contract: _StrategyManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_StrategyManager *StrategyManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *StrategyManagerUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StrategyManager.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyManagerUnpaused)
				if err := _StrategyManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_StrategyManager *StrategyManagerFilterer) ParseUnpaused(log types.Log) (*StrategyManagerUnpaused, error) {
	event := new(StrategyManagerUnpaused)
	if err := _StrategyManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
