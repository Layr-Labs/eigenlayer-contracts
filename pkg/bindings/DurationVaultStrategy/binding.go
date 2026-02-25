// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package DurationVaultStrategy

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

// IDurationVaultStrategyTypesVaultConfig is an auto generated low-level Go binding around an user-defined struct.
type IDurationVaultStrategyTypesVaultConfig struct {
	UnderlyingToken             common.Address
	VaultAdmin                  common.Address
	Arbitrator                  common.Address
	Duration                    uint32
	MaxPerDeposit               *big.Int
	StakeCap                    *big.Int
	MetadataURI                 string
	OperatorSet                 OperatorSet
	OperatorSetRegistrationData []byte
	DelegationApprover          common.Address
	OperatorMetadataURI         string
}

// OperatorSet is an auto generated low-level Go binding around an user-defined struct.
type OperatorSet struct {
	Avs common.Address
	Id  uint32
}

// DurationVaultStrategyMetaData contains all meta data concerning the DurationVaultStrategy contract.
var DurationVaultStrategyMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_strategyManager\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"_delegationManager\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"},{\"name\":\"_allocationManager\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"},{\"name\":\"_rewardsCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRewardsCoordinator\"},{\"name\":\"_strategyFactory\",\"type\":\"address\",\"internalType\":\"contractIStrategyFactory\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"advanceToWithdrawals\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allocationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allocationsActive\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"arbitrator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"beforeAddShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"beforeRemoveShares\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"newShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositsOpen\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"duration\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"explanation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getTVLLimits\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structIDurationVaultStrategyTypes.VaultConfig\",\"components\":[{\"name\":\"underlyingToken\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"vaultAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"arbitrator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxPerDeposit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"stakeCap\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"metadataURI\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operatorSetRegistrationData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"delegationApprover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorMetadataURI\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_underlyingToken\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isLocked\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isMatured\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lock\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockedAt\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"markMatured\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"maturedAt\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxPerDeposit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxTotalDeposits\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"metadataURI\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorIntegrationConfigured\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"operatorSetInfo\",\"inputs\":[],\"outputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorSetRegistered\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"rewardsCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRewardsCoordinator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setRewardsClaimer\",\"inputs\":[{\"name\":\"claimer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTVLLimits\",\"inputs\":[{\"name\":\"newMaxPerDeposit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"newMaxTotalDeposits\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"shares\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sharesToUnderlying\",\"inputs\":[{\"name\":\"amountShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sharesToUnderlyingView\",\"inputs\":[{\"name\":\"amountShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakeCap\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"state\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIDurationVaultStrategyTypes.VaultState\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"strategyFactory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategyFactory\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"strategyManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalShares\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"underlyingToShares\",\"inputs\":[{\"name\":\"amountUnderlying\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"underlyingToSharesView\",\"inputs\":[{\"name\":\"amountUnderlying\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"underlyingToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unlockAt\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unlockTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateDelegationApprover\",\"inputs\":[{\"name\":\"newDelegationApprover\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateMetadataURI\",\"inputs\":[{\"name\":\"newMetadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOperatorMetadataURI\",\"inputs\":[{\"name\":\"newOperatorMetadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateTVLLimits\",\"inputs\":[{\"name\":\"newMaxPerDeposit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"newStakeCap\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"userUnderlying\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"userUnderlyingView\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"vaultAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amountShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"amountOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawalsOpen\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"DeallocateAttempted\",\"inputs\":[{\"name\":\"success\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DeregisterAttempted\",\"inputs\":[{\"name\":\"success\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExchangeRateEmitted\",\"inputs\":[{\"name\":\"rate\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MaxPerDepositUpdated\",\"inputs\":[{\"name\":\"previousValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MaxTotalDepositsUpdated\",\"inputs\":[{\"name\":\"previousValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MetadataURIUpdated\",\"inputs\":[{\"name\":\"newMetadataURI\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyTokenSet\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIERC20\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VaultAdvancedToWithdrawals\",\"inputs\":[{\"name\":\"arbitrator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"maturedAt\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VaultInitialized\",\"inputs\":[{\"name\":\"vaultAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"arbitrator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"underlyingToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"contractIERC20\"},{\"name\":\"duration\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"maxPerDeposit\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"stakeCap\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"metadataURI\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VaultLocked\",\"inputs\":[{\"name\":\"lockedAt\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"unlockAt\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VaultMatured\",\"inputs\":[{\"name\":\"maturedAt\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BalanceExceedsMaxTotalDeposits\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CurrentlyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DepositExceedsMaxPerDeposit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DepositsLocked\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DurationAlreadyElapsed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DurationNotElapsed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidArbitrator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDuration\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNewPausedStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidVaultAdmin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MaxPerDepositExceedsMax\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeDelegatedToVaultOperator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NewSharesZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyArbitrator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyPauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyStrategyManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnderlyingToken\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnpauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyVaultAdmin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorIntegrationInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PendingAllocation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategyNotSupportedByOperatorSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TotalSharesExceedsMax\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnderlyingTokenBlacklisted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"VaultAlreadyLocked\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"VaultNotLocked\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawalAmountExceedsTotalDeposits\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawalsLockedDuringAllocations\",\"inputs\":[]}]",
	Bin: "0x610140604052348015610010575f5ffd5b5060405161420438038061420483398101604081905261002f916101e2565b8585806001600160a01b038116610059576040516339b190bb60e11b815260040160405180910390fd5b6001600160a01b03908116608052821660a05261007461010f565b50506001600160a01b0384161580159061009657506001600160a01b03831615155b80156100aa57506001600160a01b03821615155b80156100be57506001600160a01b03811615155b6100db57604051635881876760e01b815260040160405180910390fd5b6001600160a01b0380851660c05283811660e0528281166101005281166101205261010461010f565b505050505050610265565b5f54610100900460ff161561017a5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff908116146101c9575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146101df575f5ffd5b50565b5f5f5f5f5f5f60c087890312156101f7575f5ffd5b8651610202816101cb565b6020880151909650610213816101cb565b6040880151909550610224816101cb565b6060880151909450610235816101cb565b6080880151909350610246816101cb565b60a0880151909250610257816101cb565b809150509295509295509295565b60805160a05160c05160e0516101005161012051613e9b6103695f395f81816105f60152610ef701525f818161058e01528181611410015281816123b40152818161242101526124a601525f81816106f101528181610b94015281816119ac0152818161231f01528181612568015281816127e40152818161281f01528181612ab401528181612b970152612ca101525f81816107ab01528181610fa1015281816110bd01528181611260015281816113650152818161216801526121e501525f818161043b0152818161080e0152818161099d01528181610e7201528181611795015261181701525f818161056701528181611bd40152611d3c0152613e9b5ff3fe608060405234801561000f575f5ffd5b5060043610610372575f3560e01c80638c871019116101d4578063c19d93fb11610109578063df6fadc1116100a9578063f3e7387511610079578063f3e73875146107cd578063f83d08ba146107e0578063fabc1cbc146107e8578063fb4d86b4146107fb575f5ffd5b8063df6fadc114610765578063e3dae51c14610780578063e7f6f22514610793578063ea4d3c9b146107a6575f5ffd5b8063ca8aa7c7116100e4578063ca8aa7c7146106ec578063ce7c2ac214610713578063d4deae8114610726578063d9caed1214610752575f5ffd5b8063c19d93fb146106ac578063c2cca26d146106c6578063c4d66de8146106d9575f5ffd5b8063aa5dec6f11610174578063b21634821161014f578063b216348214610667578063b4e20f131461067e578063b501d66014610691578063ba28fd2e146106a4575f5ffd5b8063aa5dec6f14610635578063ab5921e11461064c578063af6eb2be14610654575f5ffd5b806399be81c8116101af57806399be81c8146105de5780639ef35710146105f1578063a4e2d63414610618578063aa082a9d14610620575f5ffd5b80638c871019146105b05780638f6a6240146105c357806394aad677146105d6575f5ffd5b8063553ca5f8116102aa5780636cc6cde11161024a5780637a8b2637116102255780637a8b2637146105475780637f2b6a0d1461055a578063886f1195146105625780638a2fc4e314610589575f5ffd5b80636cc6cde1146105195780636d8690a91461052c57806373e3c28014610534575f5ffd5b80635ac86ab7116102855780635ac86ab7146104e15780635c975abb1461050057806361b01b5d146105085780636325f65514610511575f5ffd5b8063553ca5f8146104be578063595c6a67146104d157806359d915ff146104d9575f5ffd5b806339b70e381161031557806347e7ef24116102f057806347e7ef241461047d57806353fd3e81146104905780635438a8c7146104a3578063549c4627146104b6575f5ffd5b806339b70e38146104365780633a98ef391461045d57806343fe08b014610474575f5ffd5b806311c70c9d1161035057806311c70c9d146103d5578063136439dd146103e85780631f1cc9a3146103fb5780632495a5991461040b575f5ffd5b806303e3e6eb1461037657806303ee438c1461038b5780630fb5a6b4146103a9575b5f5ffd5b610389610384366004613026565b610803565b005b61039361088e565b6040516103a0919061307e565b60405180910390f35b6033546103c090600160a01b900463ffffffff1681565b60405163ffffffff90911681526020016103a0565b6103896103e3366004613090565b61091a565b6103896103f63660046130b0565b610951565b6034546103c09063ffffffff1681565b60645461041e906001600160a01b031681565b6040516001600160a01b0390911681526020016103a0565b61041e7f000000000000000000000000000000000000000000000000000000000000000081565b61046660655481565b6040519081526020016103a0565b61046660375481565b61046661048b366004613026565b610987565b61038961049e3660046130c7565b610ab6565b60015b60405190151581526020016103a0565b6104a6610b2c565b6104666104cc366004613133565b610b54565b610389610b67565b6104a6610b7b565b6104a66104ef36600461315c565b6001805460ff9092161b9081161490565b600154610466565b61046660385481565b610389610c10565b60335461041e906001600160a01b031681565b610389610d70565b610389610542366004613026565b610e67565b6104666105553660046130b0565b61118d565b6104a66111d6565b61041e7f000000000000000000000000000000000000000000000000000000000000000081565b61041e7f000000000000000000000000000000000000000000000000000000000000000081565b6104666105be3660046130b0565b6111de565b6104666105d1366004613133565b6111e8565b6104a66111f5565b6103896105ec3660046130c7565b61121e565b61041e7f000000000000000000000000000000000000000000000000000000000000000081565b6104a66112c2565b603354600160e01b900463ffffffff166103c0565b6033546103c090600160e01b900463ffffffff1681565b6103936112ca565b610389610662366004613090565b6112ea565b6033546103c090600160c01b900463ffffffff1681565b61038961068c366004613133565b611315565b61038961069f366004613133565b6113c0565b603854610466565b603454600160201b900460ff166040516103a0919061318b565b6103896106d436600461330b565b61143f565b6103896106e7366004613133565b6116b0565b61041e7f000000000000000000000000000000000000000000000000000000000000000081565b610466610721366004613133565b61176e565b603654604080516001600160a01b0383168152600160a01b90920463ffffffff166020830152016103a0565b61046661076036600461344b565b611800565b603754603854604080519283526020830191909152016103a0565b61046661078e3660046130b0565b611902565b60325461041e906001600160a01b031681565b61041e7f000000000000000000000000000000000000000000000000000000000000000081565b6104666107db3660046130b0565b611939565b610389611943565b6103896107f63660046130b0565b611b5d565b6104a6611bca565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461084c576040516348da714f60e01b815260040160405180910390fd5b6002603454600160201b900460ff16600381111561086c5761086c613177565b0361088a576040516324d94d6960e11b815260040160405180910390fd5b5050565b6035805461089b90613489565b80601f01602080910402602001604051908101604052809291908181526020018280546108c790613489565b80156109125780601f106108e957610100808354040283529160200191610912565b820191905f5260205f20905b8154815290600101906020018083116108f557829003601f168201915b505050505081565b610922611bd2565b61092a610b2c565b6109475760405163faa4be9f60e01b815260040160405180910390fd5b61088a8282611c83565b610959611d27565b600154818116811461097e5760405163c61dca5d60e01b815260040160405180910390fd5b61088a82611dca565b5f5f61099281611e07565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146109db576040516348da714f60e01b815260040160405180910390fd5b6109e58484611e3d565b6065545f6109f56103e8836134d5565b90505f6103e8610a03611e6b565b610a0d91906134d5565b90505f610a1a87836134e8565b905080610a2784896134fb565b610a319190613512565b9550855f03610a5357604051630c392ed360e11b815260040160405180910390fd5b610a5d86856134d5565b60658190556f4b3b4ca85a86c47a098a223fffffffff1015610a9257604051632f14e8a360e11b815260040160405180910390fd5b610aab826103e8606554610aa691906134d5565b611ed5565b505050505092915050565b6032546001600160a01b03163314610ae157604051635b26872160e11b815260040160405180910390fd5b6035610aee828483613575565b507fefafb90526da1636e1335eac0151301742fb755d986954c613b90e891778ba398282604051610b20929190613656565b60405180910390a15050565b5f60015b603454600160201b900460ff166003811115610b4e57610b4e613177565b14905090565b5f610b616105558361176e565b92915050565b610b6f611d27565b610b795f19611dca565b565b6040516333869dd160e11b81525f906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063670d3ba290610bcc903090603690600401613669565b602060405180830381865afa158015610be7573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610c0b919061369f565b905090565b6033546001600160a01b03163314610c3b57604051631777988560e11b815260040160405180910390fd5b6003603454600160201b900460ff166003811115610c5b57610c5b613177565b03610c6857610b79611f21565b6002603454600160201b900460ff166003811115610c8857610c88613177565b14610ca6576040516329d0828960e01b815260040160405180910390fd5b603354600160e01b900463ffffffff164210610cd557604051632825c17f60e11b815260040160405180910390fd5b6034805463ffffffff421664ffffffffff199091168117640300000000179091556040519081527fff979382d3040b1602e0a02f0f2a454b2250aa36e891d2da0ceb95d70d11a8f29060200160405180910390a160345460405163ffffffff909116815233907f96c49d03ef64591194500229a104cd087b2d45c68234c96444c3a2a6abb0bb979060200160405180910390a2610b79611f21565b6003603454600160201b900460ff166003811115610d9057610d90613177565b03610d9d57610b79611f21565b6002603454600160201b900460ff166003811115610dbd57610dbd613177565b14610ddb576040516306560dcd60e41b815260040160405180910390fd5b603354600160e01b900463ffffffff16421015610e0b576040516306560dcd60e41b815260040160405180910390fd5b6034805463ffffffff421664ffffffffff199091168117640300000000179091556040519081527fff979382d3040b1602e0a02f0f2a454b2250aa36e891d2da0ceb95d70d11a8f29060200160405180910390a1610b79611f21565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610eb0576040516348da714f60e01b815260040160405180910390fd5b610eb8610b2c565b610ed55760405163faa4be9f60e01b815260040160405180910390fd5b60645460405163fe575a8760e01b81526001600160a01b0391821660048201527f00000000000000000000000000000000000000000000000000000000000000009091169063fe575a8790602401602060405180830381865afa158015610f3e573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610f62919061369f565b15610f80576040516337b068d160e21b815260040160405180910390fd5b604051631976849960e21b81526001600160a01b03838116600483015230917f0000000000000000000000000000000000000000000000000000000000000000909116906365da126490602401602060405180830381865afa158015610fe8573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061100c91906136be565b6001600160a01b031614611033576040516315192e1d60e11b815260040160405180910390fd5b5f61103d8261118d565b9050603754811115611062576040516334e2c93760e21b815260040160405180910390fd5b6040805160018082528183019092525f916020808301908036833701905050905030815f81518110611096576110966136d9565b6001600160a01b039283166020918202929092010152604051639004134760e01b81525f917f000000000000000000000000000000000000000000000000000000000000000016906390041347906110f49030908690600401613730565b5f60405180830381865afa15801561110e573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526111359190810190613775565b5f81518110611146576111466136d9565b602002602001015190505f611160858361055591906134d5565b90506038548111156111855760405163d86bae6760e01b815260040160405180910390fd5b505050505050565b5f5f6103e860655461119f91906134d5565b90505f6103e86111ad611e6b565b6111b791906134d5565b9050816111c485836134fb565b6111ce9190613512565b949350505050565b5f6003610b30565b5f610b6182611902565b5f610b616107db8361176e565b5f60025b603454600160201b900460ff16600381111561121757611217613177565b1415905090565b6032546001600160a01b0316331461124957604051635b26872160e11b815260040160405180910390fd5b6040516378296ec560e01b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906378296ec5906112999030908690869060040161380a565b5f604051808303815f87803b1580156112b0575f5ffd5b505af1158015611185573d5f5f3e3d5ffd5b5f60016111f9565b60606040518060800160405280604d8152602001613e19604d9139905090565b6032546001600160a01b0316331461092257604051635b26872160e11b815260040160405180910390fd5b6032546001600160a01b0316331461134057604051635b26872160e11b815260040160405180910390fd5b60405163152df25b60e21b81523060048201526001600160a01b0382811660248301527f000000000000000000000000000000000000000000000000000000000000000016906354b7c96c906044015b5f604051808303815f87803b1580156113a7575f5ffd5b505af11580156113b9573d5f5f3e3d5ffd5b5050505050565b6032546001600160a01b031633146113eb57604051635b26872160e11b815260040160405180910390fd5b60405163f22cef8560e01b81523060048201526001600160a01b0382811660248301527f0000000000000000000000000000000000000000000000000000000000000000169063f22cef8590604401611390565b5f54610100900460ff161580801561145d57505f54600160ff909116105b806114765750303b15801561147657505f5460ff166001145b61149b5760405162461bcd60e51b815260040161149290613837565b60405180910390fd5b5f805460ff1916600117905580156114bc575f805461ff0019166101001790555b60208201516001600160a01b03166114e757604051633b7dc71360e01b815260040160405180910390fd5b60408201516001600160a01b031661151157604051627528a560e41b815260040160405180910390fd5b606082015163ffffffff161580159061153e57506303c2670063ffffffff16826060015163ffffffff1611155b61155b57604051637616640160e01b815260040160405180910390fd5b61156d82608001518360a00151611c83565b815161157890611fa5565b6020820151603280546001600160a01b039283166001600160a01b0319909116179055604083015160338054606086015163ffffffff16600160a01b026001600160c01b0319909116929093169190911791909117905560c08201516035906115e19082613885565b506115eb826120f0565b6034805464ff000000001916600160201b1790558151603354603254608085015160a08601516040516001600160a01b0395861695808616959416937fbdbff63632f473bb2a7c6a4aafbc096b71fbda12e22c6b51643bfd64f13d2b9e9361166793600160a01b90920463ffffffff169290919060359061393f565b60405180910390a4801561088a575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249890602001610b20565b5f54610100900460ff16158080156116ce57505f54600160ff909116105b806116e75750303b1580156116e757505f5460ff166001145b6117035760405162461bcd60e51b815260040161149290613837565b5f805460ff191660011790558015611724575f805461ff0019166101001790555b61172d82611fa5565b801561088a575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249890602001610b20565b60405163fe243a1760e01b81526001600160a01b0382811660048301523060248301525f917f00000000000000000000000000000000000000000000000000000000000000009091169063fe243a1790604401602060405180830381865afa1580156117dc573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b6191906139e1565b5f600161180c81611e07565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614611855576040516348da714f60e01b815260040160405180910390fd5b61186085858561250d565b6065548084111561188457604051630b469df360e41b815260040160405180910390fd5b5f6118916103e8836134d5565b90505f6103e861189f611e6b565b6118a991906134d5565b9050816118b687836134fb565b6118c09190613512565b94506118cc86846134e8565b6065556118ec6118dc86836134e8565b6103e8606554610aa691906134d5565b6118f788888761253b565b505050509392505050565b5f5f6103e860655461191491906134d5565b90505f6103e8611922611e6b565b61192c91906134d5565b9050806111c483866134fb565b5f610b618261118d565b6032546001600160a01b0316331461196e57604051635b26872160e11b815260040160405180910390fd5b611976610b2c565b61199357604051630a6c62fd60e41b815260040160405180910390fd5b60405163105dea1f60e21b81525f906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690634177a87c906119e2906036906004016139f8565b5f60405180830381865afa1580156119fc573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052611a239190810190613a1f565b90505f805b8251811015611a7457306001600160a01b0316838281518110611a4d57611a4d6136d9565b60200260200101516001600160a01b031603611a6c5760019150611a74565b600101611a28565b5080611a93576040516329e0527160e11b815260040160405180910390fd5b6033805463ffffffff42818116600160c01b0263ffffffff60c01b1990931692909217928390559091611acf91600160a01b9091041682613aae565b603380546001600160e01b0316600160e01b63ffffffff938416810291909117918290556034805464020000000064ff000000001990911617905560408051600160c01b8404851681529190920490921660208301527f42cd6d7338516695d9c9ff8969dbdcf89ce22e3f2f76fda2fc11e973fe4860e4910160405180910390a1611b5861254f565b505050565b611b65611bd2565b60015480198219811614611b8c5760405163c61dca5d60e01b815260040160405180910390fd5b600182905560405182815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c9060200160405180910390a25050565b5f6002610b30565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611c2e573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611c5291906136be565b6001600160a01b0316336001600160a01b031614610b795760405163794821ff60e01b815260040160405180910390fd5b60375460408051918252602082018490527ff97ed4e083acac67830025ecbc756d8fe847cdbdca4cee3fe1e128e98b54ecb5910160405180910390a160385460408051918252602082018390527f6ab181e0440bfbf4bacdf2e99674735ce6638005490688c5f994f5399353e452910160405180910390a180821115611d1c5760405163052b07b760e21b815260040160405180910390fd5b603791909155603855565b60405163237dfb4760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906346fbf68e90602401602060405180830381865afa158015611d89573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611dad919061369f565b610b7957604051631d77d47760e21b815260040160405180910390fd5b600181905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a250565b611e1c816001805460ff9092161b9081161490565b15611e3a5760405163840a48d560e01b815260040160405180910390fd5b50565b6064546001600160a01b0383811691161461088a57604051630312abdd60e61b815260040160405180910390fd5b6064546040516370a0823160e01b81523060048201525f916001600160a01b0316906370a0823190602401602060405180830381865afa158015611eb1573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610c0b91906139e1565b7fd2494f3479e5da49d386657c292c610b5b01df313d07c62eb0cfa49924a31be881611f0984670de0b6b3a76400006134fb565b611f139190613512565b604051908152602001610b20565b5f611f2a61281b565b90507f72f957da7daaea6b52e4ff7820cb464206fd51e9f502f3027f45b5017caf4c8b81604051611f5f911515815260200190565b60405180910390a15f611f70612b7e565b90507fd0791dbc9180cb64588d7eb7658a1022dcf734b8825eb7eec68bd9516872d16881604051610b20911515815260200190565b5f54610100900460ff1661200f5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b6064820152608401611492565b606480546001600160a01b0319166001600160a01b0383161790556120335f611dca565b7f1c540707b00eb5427b6b774fc799d756516a54aee108b64b327acc55af55750760645f9054906101000a90046001600160a01b0316826001600160a01b031663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa1580156120a5573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906120c99190613aca565b604080516001600160a01b03909316835260ff90911660208301520160405180910390a150565b60e0810151516001600160a01b031661211c57604051635881876760e01b815260040160405180910390fd5b60e081015180516036805460209384015163ffffffff16600160a01b026001600160c01b03199091166001600160a01b0393841617179055604080516318891fd760e31b815290515f937f00000000000000000000000000000000000000000000000000000000000000009093169263c448feb892600480820193918290030181865afa1580156121af573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906121d39190613ae5565b90505f6121e1826001613aae565b90507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316632aa6d888846101200151838661014001516040518463ffffffff1660e01b815260040161223d93929190613b00565b5f604051808303815f87803b158015612254575f5ffd5b505af1158015612266573d5f5f3e3d5ffd5b5050505061229660405180606001604052805f6001600160a01b0316815260200160608152602001606081525090565b60e0840151516001600160a01b031681526040805160018082528183019092529060208083019080368337505050602082810182905260e0860151015181519091905f906122e6576122e66136d9565b63ffffffff909216602092830291909101909101526101008401516040808301919091525163adc2e3d960e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063adc2e3d9906123569030908590600401613b65565b5f604051808303815f87803b15801561236d575f5ffd5b505af115801561237f573d5f5f3e3d5ffd5b5050505060e08401515160405163dcbb03b360e01b81523060048201526001600160a01b0391821660248201525f60448201527f00000000000000000000000000000000000000000000000000000000000000009091169063dcbb03b3906064015f604051808303815f87803b1580156123f7575f5ffd5b505af1158015612409573d5f5f3e3d5ffd5b5050505060e0840151604051633dd3a3ab60e21b81527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169163f74e8eac916124609130915f90600401613bba565b5f604051808303815f87803b158015612477575f5ffd5b505af1158015612489573d5f5f3e3d5ffd5b505060405163059edd8760e51b81523060048201525f60248201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316925063b3dbb0e091506044015f604051808303815f87803b1580156124f1575f5ffd5b505af1158015612503573d5f5f3e3d5ffd5b5050505050505050565b6064546001600160a01b03838116911614611b5857604051630312abdd60e61b815260040160405180910390fd5b611b586001600160a01b0383168483612d68565b60405163021c373760e31b81525f906001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906310e1b9b8906125a29030906036908290600401613c02565b606060405180830381865afa1580156125bd573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906125e19190613c51565b9050806040015163ffffffff165f1461260d5760405163024b01ff60e11b815260040160405180910390fd5b6040805160018082528183019092525f91816020015b6126576040805160a081019091525f6060820181815260808301919091528190815260200160608152602001606081525090565b81526020019060019003908161262357905050604080518082019091526036546001600160a01b0381168252600160a01b900463ffffffff16602082015281519192509082905f906126ab576126ab6136d9565b6020908102919091010151526040805160018082528183019092529081602001602082028036833701905050815f815181106126e9576126e96136d9565b60200260200101516020018190525030815f8151811061270b5761270b6136d9565b6020026020010151602001515f81518110612728576127286136d9565b6001600160a01b039290921660209283029190910182015260408051600180825281830190925291828101908036833701905050815f8151811061276e5761276e6136d9565b602002602001015160400181905250670de0b6b3a7640000815f81518110612798576127986136d9565b6020026020010151604001515f815181106127b5576127b56136d9565b6001600160401b0390921660209283029190910190910152604051634a944cf760e11b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063952899ee906112999030908590600401613cd3565b5f5f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166310e1b9b8306036306040518463ffffffff1660e01b815260040161286e93929190613c02565b606060405180830381865afa158015612889573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906128ad9190613c51565b80519091506001600160401b03161580156128cd57506020810151600f0b155b156128da57600191505090565b604081015163ffffffff16156128f1575f91505090565b6040805160018082528183019092525f91816020015b61293b6040805160a081019091525f6060820181815260808301919091528190815260200160608152602001606081525090565b81526020019060019003908161290757905050604080518082019091526036546001600160a01b0381168252600160a01b900463ffffffff16602082015281519192509082905f9061298f5761298f6136d9565b6020908102919091010151526040805160018082528183019092529081602001602082028036833701905050815f815181106129cd576129cd6136d9565b60200260200101516020018190525030815f815181106129ef576129ef6136d9565b6020026020010151602001515f81518110612a0c57612a0c6136d9565b6001600160a01b039290921660209283029190910182015260408051600180825281830190925291828101908036833701905050815f81518110612a5257612a526136d9565b6020026020010151604001819052505f815f81518110612a7457612a746136d9565b6020026020010151604001515f81518110612a9157612a916136d9565b60200260200101906001600160401b031690816001600160401b0316815250505f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663952899ee60e01b3084604051602401612af7929190613cd3565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909416939093179092529051612b359190613dc6565b5f604051808303815f865af19150503d805f8114612b6e576040519150601f19603f3d011682016040523d82523d5f602084013e612b73565b606091505b509095945050505050565b6040516333869dd160e11b81525f906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063670d3ba290612bcf903090603690600401613669565b602060405180830381865afa158015612bea573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612c0e919061369f565b612c185750600190565b6040805160608082018352818301523081526036546001600160a01b0316602082015281516001808252818401909352909181602001602082028036833701905050604082018190526036548151600160a01b90910463ffffffff1691905f90612c8457612c846136d9565b602002602001019063ffffffff16908163ffffffff16815250505f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316636e3492b560e01b83604051602401612ce29190613ddc565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909416939093179092529051612d209190613dc6565b5f604051808303815f865af19150503d805f8114612d59576040519150601f19603f3d011682016040523d82523d5f602084013e612d5e565b606091505b5090949350505050565b604080516001600160a01b03848116602483015260448083018590528351808403909101815260649092018352602080830180516001600160e01b031663a9059cbb60e01b17905283518085019094528084527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c656490840152611b58928692915f91612df7918516908490612e76565b905080515f1480612e17575080806020019051810190612e17919061369f565b611b585760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608401611492565b6060612e8484845f85612e8e565b90505b9392505050565b606082471015612eef5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b6064820152608401611492565b5f5f866001600160a01b03168587604051612f0a9190613dc6565b5f6040518083038185875af1925050503d805f8114612f44576040519150601f19603f3d011682016040523d82523d5f602084013e612f49565b606091505b5091509150612f5a87838387612f65565b979650505050505050565b60608315612fd35782515f03612fcc576001600160a01b0385163b612fcc5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401611492565b50816111ce565b6111ce8383815115612fe85781518083602001fd5b8060405162461bcd60e51b8152600401611492919061307e565b6001600160a01b0381168114611e3a575f5ffd5b803561302181613002565b919050565b5f5f60408385031215613037575f5ffd5b823561304281613002565b946020939093013593505050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f612e876020830184613050565b5f5f604083850312156130a1575f5ffd5b50508035926020909101359150565b5f602082840312156130c0575f5ffd5b5035919050565b5f5f602083850312156130d8575f5ffd5b82356001600160401b038111156130ed575f5ffd5b8301601f810185136130fd575f5ffd5b80356001600160401b03811115613112575f5ffd5b856020828401011115613123575f5ffd5b6020919091019590945092505050565b5f60208284031215613143575f5ffd5b8135612e8781613002565b60ff81168114611e3a575f5ffd5b5f6020828403121561316c575f5ffd5b8135612e878161314e565b634e487b7160e01b5f52602160045260245ffd5b60208101600483106131ab57634e487b7160e01b5f52602160045260245ffd5b91905290565b634e487b7160e01b5f52604160045260245ffd5b60405161016081016001600160401b03811182821017156131e8576131e86131b1565b60405290565b604051601f8201601f191681016001600160401b0381118282101715613216576132166131b1565b604052919050565b63ffffffff81168114611e3a575f5ffd5b80356130218161321e565b5f82601f830112613249575f5ffd5b8135602083015f5f6001600160401b03841115613268576132686131b1565b50601f8301601f191660200161327d816131ee565b915050828152858383011115613291575f5ffd5b828260208301375f92810160200192909252509392505050565b5f604082840312156132bb575f5ffd5b604080519081016001600160401b03811182821017156132dd576132dd6131b1565b60405290508082356132ee81613002565b815260208301356132fe8161321e565b6020919091015292915050565b5f6020828403121561331b575f5ffd5b81356001600160401b03811115613330575f5ffd5b82016101808185031215613342575f5ffd5b61334a6131c5565b61335382613016565b815261336160208301613016565b602082015261337260408301613016565b60408201526133836060830161322f565b60608201526080828101359082015260a0808301359082015260c08201356001600160401b038111156133b4575f5ffd5b6133c08682850161323a565b60c0830152506133d38560e084016132ab565b60e08201526101208201356001600160401b038111156133f1575f5ffd5b6133fd8682850161323a565b610100830152506134116101408301613016565b6101208201526101608201356001600160401b03811115613430575f5ffd5b61343c8682850161323a565b61014083015250949350505050565b5f5f5f6060848603121561345d575f5ffd5b833561346881613002565b9250602084013561347881613002565b929592945050506040919091013590565b600181811c9082168061349d57607f821691505b6020821081036134bb57634e487b7160e01b5f52602260045260245ffd5b50919050565b634e487b7160e01b5f52601160045260245ffd5b80820180821115610b6157610b616134c1565b81810381811115610b6157610b616134c1565b8082028115828204841417610b6157610b616134c1565b5f8261352c57634e487b7160e01b5f52601260045260245ffd5b500490565b601f821115611b5857805f5260205f20601f840160051c810160208510156135565750805b601f840160051c820191505b818110156113b9575f8155600101613562565b6001600160401b0383111561358c5761358c6131b1565b6135a08361359a8354613489565b83613531565b5f601f8411600181146135d1575f85156135ba5750838201355b5f19600387901b1c1916600186901b1783556113b9565b5f83815260208120601f198716915b8281101561360057868501358255602094850194600190920191016135e0565b508682101561361c575f1960f88860031b161c19848701351681555b505060018560011b0183555050505050565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b602081525f612e8460208301848661362e565b6001600160a01b038316815260608101612e876020830184546001600160a01b038116825260a01c63ffffffff16602090910152565b5f602082840312156136af575f5ffd5b81518015158114612e87575f5ffd5b5f602082840312156136ce575f5ffd5b8151612e8781613002565b634e487b7160e01b5f52603260045260245ffd5b5f8151808452602084019350602083015f5b828110156137265781516001600160a01b03168652602095860195909101906001016136ff565b5093949350505050565b6001600160a01b03831681526040602082018190525f90612e84908301846136ed565b5f6001600160401b0382111561376b5761376b6131b1565b5060051b60200190565b5f60208284031215613785575f5ffd5b81516001600160401b0381111561379a575f5ffd5b8201601f810184136137aa575f5ffd5b80516137bd6137b882613753565b6131ee565b8082825260208201915060208360051b8501019250868311156137de575f5ffd5b6020840193505b828410156138005783518252602093840193909101906137e5565b9695505050505050565b6001600160a01b03841681526040602082018190525f9061382e908301848661362e565b95945050505050565b6020808252602e908201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160408201526d191e481a5b9a5d1a585b1a5e995960921b606082015260800190565b81516001600160401b0381111561389e5761389e6131b1565b6138b2816138ac8454613489565b84613531565b6020601f8211600181146138e4575f83156138cd5750848201515b5f19600385901b1c1916600184901b1784556113b9565b5f84815260208120601f198516915b8281101561391357878501518255602094850194600190920191016138f3565b508482101561393057868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b63ffffffff85168152836020820152826040820152608060608201525f5f835461396881613489565b806080860152600182165f811461398657600181146139a2576139d3565b60ff19831660a087015260a082151560051b87010193506139d3565b865f5260205f205f5b838110156139ca57815488820160a001526001909101906020016139ab565b870160a0019450505b509198975050505050505050565b5f602082840312156139f1575f5ffd5b5051919050565b60408101610b618284546001600160a01b038116825260a01c63ffffffff16602090910152565b5f60208284031215613a2f575f5ffd5b81516001600160401b03811115613a44575f5ffd5b8201601f81018413613a54575f5ffd5b8051613a626137b882613753565b8082825260208201915060208360051b850101925086831115613a83575f5ffd5b6020840193505b82841015613800578351613a9d81613002565b825260209384019390910190613a8a565b63ffffffff8181168382160190811115610b6157610b616134c1565b5f60208284031215613ada575f5ffd5b8151612e878161314e565b5f60208284031215613af5575f5ffd5b8151612e878161321e565b6001600160a01b038416815263ffffffff831660208201526060604082018190525f9061382e90830184613050565b5f8151808452602084019350602083015f5b8281101561372657815163ffffffff16865260209586019590910190600101613b41565b6001600160a01b038381168252604060208084018290528451909216908301528201516060808301525f90613b9d60a0840182613b2f565b90506040840151603f198483030160808501526138008282613050565b6001600160a01b038416815260808101613bf0602083018580516001600160a01b0316825260209081015163ffffffff16910152565b61ffff83166060830152949350505050565b6001600160a01b038416815260808101613c386020830185546001600160a01b038116825260a01c63ffffffff16602090910152565b6001600160a01b03929092166060919091015292915050565b5f6060828403128015613c62575f5ffd5b50604051606081016001600160401b0381118282101715613c8557613c856131b1565b60405282516001600160401b0381168114613c9e575f5ffd5b81526020830151600f81900b8114613cb4575f5ffd5b60208201526040830151613cc78161321e565b60408201529392505050565b5f6040820160018060a01b03851683526040602084015280845180835260608501915060608160051b8601019250602086015f5b82811015613db957868503605f190184528151805180516001600160a01b0316875260209081015163ffffffff1690870152602081015160806040880152613d5260808801826136ed565b60409290920151878303606089015280518084526020918201935f935091909101905b80831015613da1576001600160401b038451168252602082019150602084019350600183019250613d75565b50965050506020938401939190910190600101613d07565b5092979650505050505050565b5f82518060208501845e5f920191825250919050565b602080825282516001600160a01b039081168383015290830151166040808301919091528201516060808301525f906111ce6080840182613b2f56fe4475726174696f6e2d626f756e64207661756c74207374726174656779207769746820636f6e666967757261626c65206465706f736974206361707320616e64206c6f636b20706572696f6473a2646970667358221220f0b0de05bf99024d1cfd0e195e186b64d0d54aab320b8768384d02307521211b64736f6c634300081e0033",
}

// DurationVaultStrategyABI is the input ABI used to generate the binding from.
// Deprecated: Use DurationVaultStrategyMetaData.ABI instead.
var DurationVaultStrategyABI = DurationVaultStrategyMetaData.ABI

// DurationVaultStrategyBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DurationVaultStrategyMetaData.Bin instead.
var DurationVaultStrategyBin = DurationVaultStrategyMetaData.Bin

// DeployDurationVaultStrategy deploys a new Ethereum contract, binding an instance of DurationVaultStrategy to it.
func DeployDurationVaultStrategy(auth *bind.TransactOpts, backend bind.ContractBackend, _strategyManager common.Address, _pauserRegistry common.Address, _delegationManager common.Address, _allocationManager common.Address, _rewardsCoordinator common.Address, _strategyFactory common.Address) (common.Address, *types.Transaction, *DurationVaultStrategy, error) {
	parsed, err := DurationVaultStrategyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DurationVaultStrategyBin), backend, _strategyManager, _pauserRegistry, _delegationManager, _allocationManager, _rewardsCoordinator, _strategyFactory)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DurationVaultStrategy{DurationVaultStrategyCaller: DurationVaultStrategyCaller{contract: contract}, DurationVaultStrategyTransactor: DurationVaultStrategyTransactor{contract: contract}, DurationVaultStrategyFilterer: DurationVaultStrategyFilterer{contract: contract}}, nil
}

// DurationVaultStrategy is an auto generated Go binding around an Ethereum contract.
type DurationVaultStrategy struct {
	DurationVaultStrategyCaller     // Read-only binding to the contract
	DurationVaultStrategyTransactor // Write-only binding to the contract
	DurationVaultStrategyFilterer   // Log filterer for contract events
}

// DurationVaultStrategyCaller is an auto generated read-only Go binding around an Ethereum contract.
type DurationVaultStrategyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DurationVaultStrategyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DurationVaultStrategyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DurationVaultStrategyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DurationVaultStrategyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DurationVaultStrategySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DurationVaultStrategySession struct {
	Contract     *DurationVaultStrategy // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DurationVaultStrategyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DurationVaultStrategyCallerSession struct {
	Contract *DurationVaultStrategyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// DurationVaultStrategyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DurationVaultStrategyTransactorSession struct {
	Contract     *DurationVaultStrategyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// DurationVaultStrategyRaw is an auto generated low-level Go binding around an Ethereum contract.
type DurationVaultStrategyRaw struct {
	Contract *DurationVaultStrategy // Generic contract binding to access the raw methods on
}

// DurationVaultStrategyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DurationVaultStrategyCallerRaw struct {
	Contract *DurationVaultStrategyCaller // Generic read-only contract binding to access the raw methods on
}

// DurationVaultStrategyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DurationVaultStrategyTransactorRaw struct {
	Contract *DurationVaultStrategyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDurationVaultStrategy creates a new instance of DurationVaultStrategy, bound to a specific deployed contract.
func NewDurationVaultStrategy(address common.Address, backend bind.ContractBackend) (*DurationVaultStrategy, error) {
	contract, err := bindDurationVaultStrategy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DurationVaultStrategy{DurationVaultStrategyCaller: DurationVaultStrategyCaller{contract: contract}, DurationVaultStrategyTransactor: DurationVaultStrategyTransactor{contract: contract}, DurationVaultStrategyFilterer: DurationVaultStrategyFilterer{contract: contract}}, nil
}

// NewDurationVaultStrategyCaller creates a new read-only instance of DurationVaultStrategy, bound to a specific deployed contract.
func NewDurationVaultStrategyCaller(address common.Address, caller bind.ContractCaller) (*DurationVaultStrategyCaller, error) {
	contract, err := bindDurationVaultStrategy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DurationVaultStrategyCaller{contract: contract}, nil
}

// NewDurationVaultStrategyTransactor creates a new write-only instance of DurationVaultStrategy, bound to a specific deployed contract.
func NewDurationVaultStrategyTransactor(address common.Address, transactor bind.ContractTransactor) (*DurationVaultStrategyTransactor, error) {
	contract, err := bindDurationVaultStrategy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DurationVaultStrategyTransactor{contract: contract}, nil
}

// NewDurationVaultStrategyFilterer creates a new log filterer instance of DurationVaultStrategy, bound to a specific deployed contract.
func NewDurationVaultStrategyFilterer(address common.Address, filterer bind.ContractFilterer) (*DurationVaultStrategyFilterer, error) {
	contract, err := bindDurationVaultStrategy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DurationVaultStrategyFilterer{contract: contract}, nil
}

// bindDurationVaultStrategy binds a generic wrapper to an already deployed contract.
func bindDurationVaultStrategy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DurationVaultStrategyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DurationVaultStrategy *DurationVaultStrategyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DurationVaultStrategy.Contract.DurationVaultStrategyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DurationVaultStrategy *DurationVaultStrategyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DurationVaultStrategy.Contract.DurationVaultStrategyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DurationVaultStrategy *DurationVaultStrategyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DurationVaultStrategy.Contract.DurationVaultStrategyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DurationVaultStrategy *DurationVaultStrategyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DurationVaultStrategy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DurationVaultStrategy *DurationVaultStrategyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DurationVaultStrategy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DurationVaultStrategy *DurationVaultStrategyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DurationVaultStrategy.Contract.contract.Transact(opts, method, params...)
}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_DurationVaultStrategy *DurationVaultStrategyCaller) AllocationManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DurationVaultStrategy.contract.Call(opts, &out, "allocationManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_DurationVaultStrategy *DurationVaultStrategySession) AllocationManager() (common.Address, error) {
	return _DurationVaultStrategy.Contract.AllocationManager(&_DurationVaultStrategy.CallOpts)
}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_DurationVaultStrategy *DurationVaultStrategyCallerSession) AllocationManager() (common.Address, error) {
	return _DurationVaultStrategy.Contract.AllocationManager(&_DurationVaultStrategy.CallOpts)
}

// AllocationsActive is a free data retrieval call binding the contract method 0xfb4d86b4.
//
// Solidity: function allocationsActive() view returns(bool)
func (_DurationVaultStrategy *DurationVaultStrategyCaller) AllocationsActive(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DurationVaultStrategy.contract.Call(opts, &out, "allocationsActive")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllocationsActive is a free data retrieval call binding the contract method 0xfb4d86b4.
//
// Solidity: function allocationsActive() view returns(bool)
func (_DurationVaultStrategy *DurationVaultStrategySession) AllocationsActive() (bool, error) {
	return _DurationVaultStrategy.Contract.AllocationsActive(&_DurationVaultStrategy.CallOpts)
}

// AllocationsActive is a free data retrieval call binding the contract method 0xfb4d86b4.
//
// Solidity: function allocationsActive() view returns(bool)
func (_DurationVaultStrategy *DurationVaultStrategyCallerSession) AllocationsActive() (bool, error) {
	return _DurationVaultStrategy.Contract.AllocationsActive(&_DurationVaultStrategy.CallOpts)
}

// Arbitrator is a free data retrieval call binding the contract method 0x6cc6cde1.
//
// Solidity: function arbitrator() view returns(address)
func (_DurationVaultStrategy *DurationVaultStrategyCaller) Arbitrator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DurationVaultStrategy.contract.Call(opts, &out, "arbitrator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Arbitrator is a free data retrieval call binding the contract method 0x6cc6cde1.
//
// Solidity: function arbitrator() view returns(address)
func (_DurationVaultStrategy *DurationVaultStrategySession) Arbitrator() (common.Address, error) {
	return _DurationVaultStrategy.Contract.Arbitrator(&_DurationVaultStrategy.CallOpts)
}

// Arbitrator is a free data retrieval call binding the contract method 0x6cc6cde1.
//
// Solidity: function arbitrator() view returns(address)
func (_DurationVaultStrategy *DurationVaultStrategyCallerSession) Arbitrator() (common.Address, error) {
	return _DurationVaultStrategy.Contract.Arbitrator(&_DurationVaultStrategy.CallOpts)
}

// BeforeAddShares is a free data retrieval call binding the contract method 0x73e3c280.
//
// Solidity: function beforeAddShares(address staker, uint256 shares) view returns()
func (_DurationVaultStrategy *DurationVaultStrategyCaller) BeforeAddShares(opts *bind.CallOpts, staker common.Address, shares *big.Int) error {
	var out []interface{}
	err := _DurationVaultStrategy.contract.Call(opts, &out, "beforeAddShares", staker, shares)

	if err != nil {
		return err
	}

	return err

}

// BeforeAddShares is a free data retrieval call binding the contract method 0x73e3c280.
//
// Solidity: function beforeAddShares(address staker, uint256 shares) view returns()
func (_DurationVaultStrategy *DurationVaultStrategySession) BeforeAddShares(staker common.Address, shares *big.Int) error {
	return _DurationVaultStrategy.Contract.BeforeAddShares(&_DurationVaultStrategy.CallOpts, staker, shares)
}

// BeforeAddShares is a free data retrieval call binding the contract method 0x73e3c280.
//
// Solidity: function beforeAddShares(address staker, uint256 shares) view returns()
func (_DurationVaultStrategy *DurationVaultStrategyCallerSession) BeforeAddShares(staker common.Address, shares *big.Int) error {
	return _DurationVaultStrategy.Contract.BeforeAddShares(&_DurationVaultStrategy.CallOpts, staker, shares)
}

// BeforeRemoveShares is a free data retrieval call binding the contract method 0x03e3e6eb.
//
// Solidity: function beforeRemoveShares(address , uint256 ) view returns()
func (_DurationVaultStrategy *DurationVaultStrategyCaller) BeforeRemoveShares(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) error {
	var out []interface{}
	err := _DurationVaultStrategy.contract.Call(opts, &out, "beforeRemoveShares", arg0, arg1)

	if err != nil {
		return err
	}

	return err

}

// BeforeRemoveShares is a free data retrieval call binding the contract method 0x03e3e6eb.
//
// Solidity: function beforeRemoveShares(address , uint256 ) view returns()
func (_DurationVaultStrategy *DurationVaultStrategySession) BeforeRemoveShares(arg0 common.Address, arg1 *big.Int) error {
	return _DurationVaultStrategy.Contract.BeforeRemoveShares(&_DurationVaultStrategy.CallOpts, arg0, arg1)
}

// BeforeRemoveShares is a free data retrieval call binding the contract method 0x03e3e6eb.
//
// Solidity: function beforeRemoveShares(address , uint256 ) view returns()
func (_DurationVaultStrategy *DurationVaultStrategyCallerSession) BeforeRemoveShares(arg0 common.Address, arg1 *big.Int) error {
	return _DurationVaultStrategy.Contract.BeforeRemoveShares(&_DurationVaultStrategy.CallOpts, arg0, arg1)
}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_DurationVaultStrategy *DurationVaultStrategyCaller) DelegationManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DurationVaultStrategy.contract.Call(opts, &out, "delegationManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_DurationVaultStrategy *DurationVaultStrategySession) DelegationManager() (common.Address, error) {
	return _DurationVaultStrategy.Contract.DelegationManager(&_DurationVaultStrategy.CallOpts)
}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_DurationVaultStrategy *DurationVaultStrategyCallerSession) DelegationManager() (common.Address, error) {
	return _DurationVaultStrategy.Contract.DelegationManager(&_DurationVaultStrategy.CallOpts)
}

// DepositsOpen is a free data retrieval call binding the contract method 0x549c4627.
//
// Solidity: function depositsOpen() view returns(bool)
func (_DurationVaultStrategy *DurationVaultStrategyCaller) DepositsOpen(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DurationVaultStrategy.contract.Call(opts, &out, "depositsOpen")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DepositsOpen is a free data retrieval call binding the contract method 0x549c4627.
//
// Solidity: function depositsOpen() view returns(bool)
func (_DurationVaultStrategy *DurationVaultStrategySession) DepositsOpen() (bool, error) {
	return _DurationVaultStrategy.Contract.DepositsOpen(&_DurationVaultStrategy.CallOpts)
}

// DepositsOpen is a free data retrieval call binding the contract method 0x549c4627.
//
// Solidity: function depositsOpen() view returns(bool)
func (_DurationVaultStrategy *DurationVaultStrategyCallerSession) DepositsOpen() (bool, error) {
	return _DurationVaultStrategy.Contract.DepositsOpen(&_DurationVaultStrategy.CallOpts)
}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint32)
func (_DurationVaultStrategy *DurationVaultStrategyCaller) Duration(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _DurationVaultStrategy.contract.Call(opts, &out, "duration")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint32)
func (_DurationVaultStrategy *DurationVaultStrategySession) Duration() (uint32, error) {
	return _DurationVaultStrategy.Contract.Duration(&_DurationVaultStrategy.CallOpts)
}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint32)
func (_DurationVaultStrategy *DurationVaultStrategyCallerSession) Duration() (uint32, error) {
	return _DurationVaultStrategy.Contract.Duration(&_DurationVaultStrategy.CallOpts)
}

// Explanation is a free data retrieval call binding the contract method 0xab5921e1.
//
// Solidity: function explanation() pure returns(string)
func (_DurationVaultStrategy *DurationVaultStrategyCaller) Explanation(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DurationVaultStrategy.contract.Call(opts, &out, "explanation")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Explanation is a free data retrieval call binding the contract method 0xab5921e1.
//
// Solidity: function explanation() pure returns(string)
func (_DurationVaultStrategy *DurationVaultStrategySession) Explanation() (string, error) {
	return _DurationVaultStrategy.Contract.Explanation(&_DurationVaultStrategy.CallOpts)
}

// Explanation is a free data retrieval call binding the contract method 0xab5921e1.
//
// Solidity: function explanation() pure returns(string)
func (_DurationVaultStrategy *DurationVaultStrategyCallerSession) Explanation() (string, error) {
	return _DurationVaultStrategy.Contract.Explanation(&_DurationVaultStrategy.CallOpts)
}

// GetTVLLimits is a free data retrieval call binding the contract method 0xdf6fadc1.
//
// Solidity: function getTVLLimits() view returns(uint256, uint256)
func (_DurationVaultStrategy *DurationVaultStrategyCaller) GetTVLLimits(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _DurationVaultStrategy.contract.Call(opts, &out, "getTVLLimits")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetTVLLimits is a free data retrieval call binding the contract method 0xdf6fadc1.
//
// Solidity: function getTVLLimits() view returns(uint256, uint256)
func (_DurationVaultStrategy *DurationVaultStrategySession) GetTVLLimits() (*big.Int, *big.Int, error) {
	return _DurationVaultStrategy.Contract.GetTVLLimits(&_DurationVaultStrategy.CallOpts)
}

// GetTVLLimits is a free data retrieval call binding the contract method 0xdf6fadc1.
//
// Solidity: function getTVLLimits() view returns(uint256, uint256)
func (_DurationVaultStrategy *DurationVaultStrategyCallerSession) GetTVLLimits() (*big.Int, *big.Int, error) {
	return _DurationVaultStrategy.Contract.GetTVLLimits(&_DurationVaultStrategy.CallOpts)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() view returns(bool)
func (_DurationVaultStrategy *DurationVaultStrategyCaller) IsLocked(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DurationVaultStrategy.contract.Call(opts, &out, "isLocked")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() view returns(bool)
func (_DurationVaultStrategy *DurationVaultStrategySession) IsLocked() (bool, error) {
	return _DurationVaultStrategy.Contract.IsLocked(&_DurationVaultStrategy.CallOpts)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() view returns(bool)
func (_DurationVaultStrategy *DurationVaultStrategyCallerSession) IsLocked() (bool, error) {
	return _DurationVaultStrategy.Contract.IsLocked(&_DurationVaultStrategy.CallOpts)
}

// IsMatured is a free data retrieval call binding the contract method 0x7f2b6a0d.
//
// Solidity: function isMatured() view returns(bool)
func (_DurationVaultStrategy *DurationVaultStrategyCaller) IsMatured(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DurationVaultStrategy.contract.Call(opts, &out, "isMatured")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMatured is a free data retrieval call binding the contract method 0x7f2b6a0d.
//
// Solidity: function isMatured() view returns(bool)
func (_DurationVaultStrategy *DurationVaultStrategySession) IsMatured() (bool, error) {
	return _DurationVaultStrategy.Contract.IsMatured(&_DurationVaultStrategy.CallOpts)
}

// IsMatured is a free data retrieval call binding the contract method 0x7f2b6a0d.
//
// Solidity: function isMatured() view returns(bool)
func (_DurationVaultStrategy *DurationVaultStrategyCallerSession) IsMatured() (bool, error) {
	return _DurationVaultStrategy.Contract.IsMatured(&_DurationVaultStrategy.CallOpts)
}

// LockedAt is a free data retrieval call binding the contract method 0xb2163482.
//
// Solidity: function lockedAt() view returns(uint32)
func (_DurationVaultStrategy *DurationVaultStrategyCaller) LockedAt(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _DurationVaultStrategy.contract.Call(opts, &out, "lockedAt")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LockedAt is a free data retrieval call binding the contract method 0xb2163482.
//
// Solidity: function lockedAt() view returns(uint32)
func (_DurationVaultStrategy *DurationVaultStrategySession) LockedAt() (uint32, error) {
	return _DurationVaultStrategy.Contract.LockedAt(&_DurationVaultStrategy.CallOpts)
}

// LockedAt is a free data retrieval call binding the contract method 0xb2163482.
//
// Solidity: function lockedAt() view returns(uint32)
func (_DurationVaultStrategy *DurationVaultStrategyCallerSession) LockedAt() (uint32, error) {
	return _DurationVaultStrategy.Contract.LockedAt(&_DurationVaultStrategy.CallOpts)
}

// MaturedAt is a free data retrieval call binding the contract method 0x1f1cc9a3.
//
// Solidity: function maturedAt() view returns(uint32)
func (_DurationVaultStrategy *DurationVaultStrategyCaller) MaturedAt(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _DurationVaultStrategy.contract.Call(opts, &out, "maturedAt")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MaturedAt is a free data retrieval call binding the contract method 0x1f1cc9a3.
//
// Solidity: function maturedAt() view returns(uint32)
func (_DurationVaultStrategy *DurationVaultStrategySession) MaturedAt() (uint32, error) {
	return _DurationVaultStrategy.Contract.MaturedAt(&_DurationVaultStrategy.CallOpts)
}

// MaturedAt is a free data retrieval call binding the contract method 0x1f1cc9a3.
//
// Solidity: function maturedAt() view returns(uint32)
func (_DurationVaultStrategy *DurationVaultStrategyCallerSession) MaturedAt() (uint32, error) {
	return _DurationVaultStrategy.Contract.MaturedAt(&_DurationVaultStrategy.CallOpts)
}

// MaxPerDeposit is a free data retrieval call binding the contract method 0x43fe08b0.
//
// Solidity: function maxPerDeposit() view returns(uint256)
func (_DurationVaultStrategy *DurationVaultStrategyCaller) MaxPerDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DurationVaultStrategy.contract.Call(opts, &out, "maxPerDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxPerDeposit is a free data retrieval call binding the contract method 0x43fe08b0.
//
// Solidity: function maxPerDeposit() view returns(uint256)
func (_DurationVaultStrategy *DurationVaultStrategySession) MaxPerDeposit() (*big.Int, error) {
	return _DurationVaultStrategy.Contract.MaxPerDeposit(&_DurationVaultStrategy.CallOpts)
}

// MaxPerDeposit is a free data retrieval call binding the contract method 0x43fe08b0.
//
// Solidity: function maxPerDeposit() view returns(uint256)
func (_DurationVaultStrategy *DurationVaultStrategyCallerSession) MaxPerDeposit() (*big.Int, error) {
	return _DurationVaultStrategy.Contract.MaxPerDeposit(&_DurationVaultStrategy.CallOpts)
}

// MaxTotalDeposits is a free data retrieval call binding the contract method 0x61b01b5d.
//
// Solidity: function maxTotalDeposits() view returns(uint256)
func (_DurationVaultStrategy *DurationVaultStrategyCaller) MaxTotalDeposits(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DurationVaultStrategy.contract.Call(opts, &out, "maxTotalDeposits")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxTotalDeposits is a free data retrieval call binding the contract method 0x61b01b5d.
//
// Solidity: function maxTotalDeposits() view returns(uint256)
func (_DurationVaultStrategy *DurationVaultStrategySession) MaxTotalDeposits() (*big.Int, error) {
	return _DurationVaultStrategy.Contract.MaxTotalDeposits(&_DurationVaultStrategy.CallOpts)
}

// MaxTotalDeposits is a free data retrieval call binding the contract method 0x61b01b5d.
//
// Solidity: function maxTotalDeposits() view returns(uint256)
func (_DurationVaultStrategy *DurationVaultStrategyCallerSession) MaxTotalDeposits() (*big.Int, error) {
	return _DurationVaultStrategy.Contract.MaxTotalDeposits(&_DurationVaultStrategy.CallOpts)
}

// MetadataURI is a free data retrieval call binding the contract method 0x03ee438c.
//
// Solidity: function metadataURI() view returns(string)
func (_DurationVaultStrategy *DurationVaultStrategyCaller) MetadataURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DurationVaultStrategy.contract.Call(opts, &out, "metadataURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// MetadataURI is a free data retrieval call binding the contract method 0x03ee438c.
//
// Solidity: function metadataURI() view returns(string)
func (_DurationVaultStrategy *DurationVaultStrategySession) MetadataURI() (string, error) {
	return _DurationVaultStrategy.Contract.MetadataURI(&_DurationVaultStrategy.CallOpts)
}

// MetadataURI is a free data retrieval call binding the contract method 0x03ee438c.
//
// Solidity: function metadataURI() view returns(string)
func (_DurationVaultStrategy *DurationVaultStrategyCallerSession) MetadataURI() (string, error) {
	return _DurationVaultStrategy.Contract.MetadataURI(&_DurationVaultStrategy.CallOpts)
}

// OperatorIntegrationConfigured is a free data retrieval call binding the contract method 0x5438a8c7.
//
// Solidity: function operatorIntegrationConfigured() pure returns(bool)
func (_DurationVaultStrategy *DurationVaultStrategyCaller) OperatorIntegrationConfigured(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DurationVaultStrategy.contract.Call(opts, &out, "operatorIntegrationConfigured")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OperatorIntegrationConfigured is a free data retrieval call binding the contract method 0x5438a8c7.
//
// Solidity: function operatorIntegrationConfigured() pure returns(bool)
func (_DurationVaultStrategy *DurationVaultStrategySession) OperatorIntegrationConfigured() (bool, error) {
	return _DurationVaultStrategy.Contract.OperatorIntegrationConfigured(&_DurationVaultStrategy.CallOpts)
}

// OperatorIntegrationConfigured is a free data retrieval call binding the contract method 0x5438a8c7.
//
// Solidity: function operatorIntegrationConfigured() pure returns(bool)
func (_DurationVaultStrategy *DurationVaultStrategyCallerSession) OperatorIntegrationConfigured() (bool, error) {
	return _DurationVaultStrategy.Contract.OperatorIntegrationConfigured(&_DurationVaultStrategy.CallOpts)
}

// OperatorSetInfo is a free data retrieval call binding the contract method 0xd4deae81.
//
// Solidity: function operatorSetInfo() view returns(address avs, uint32 operatorSetId)
func (_DurationVaultStrategy *DurationVaultStrategyCaller) OperatorSetInfo(opts *bind.CallOpts) (struct {
	Avs           common.Address
	OperatorSetId uint32
}, error) {
	var out []interface{}
	err := _DurationVaultStrategy.contract.Call(opts, &out, "operatorSetInfo")

	outstruct := new(struct {
		Avs           common.Address
		OperatorSetId uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Avs = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.OperatorSetId = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

// OperatorSetInfo is a free data retrieval call binding the contract method 0xd4deae81.
//
// Solidity: function operatorSetInfo() view returns(address avs, uint32 operatorSetId)
func (_DurationVaultStrategy *DurationVaultStrategySession) OperatorSetInfo() (struct {
	Avs           common.Address
	OperatorSetId uint32
}, error) {
	return _DurationVaultStrategy.Contract.OperatorSetInfo(&_DurationVaultStrategy.CallOpts)
}

// OperatorSetInfo is a free data retrieval call binding the contract method 0xd4deae81.
//
// Solidity: function operatorSetInfo() view returns(address avs, uint32 operatorSetId)
func (_DurationVaultStrategy *DurationVaultStrategyCallerSession) OperatorSetInfo() (struct {
	Avs           common.Address
	OperatorSetId uint32
}, error) {
	return _DurationVaultStrategy.Contract.OperatorSetInfo(&_DurationVaultStrategy.CallOpts)
}

// OperatorSetRegistered is a free data retrieval call binding the contract method 0x59d915ff.
//
// Solidity: function operatorSetRegistered() view returns(bool)
func (_DurationVaultStrategy *DurationVaultStrategyCaller) OperatorSetRegistered(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DurationVaultStrategy.contract.Call(opts, &out, "operatorSetRegistered")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OperatorSetRegistered is a free data retrieval call binding the contract method 0x59d915ff.
//
// Solidity: function operatorSetRegistered() view returns(bool)
func (_DurationVaultStrategy *DurationVaultStrategySession) OperatorSetRegistered() (bool, error) {
	return _DurationVaultStrategy.Contract.OperatorSetRegistered(&_DurationVaultStrategy.CallOpts)
}

// OperatorSetRegistered is a free data retrieval call binding the contract method 0x59d915ff.
//
// Solidity: function operatorSetRegistered() view returns(bool)
func (_DurationVaultStrategy *DurationVaultStrategyCallerSession) OperatorSetRegistered() (bool, error) {
	return _DurationVaultStrategy.Contract.OperatorSetRegistered(&_DurationVaultStrategy.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_DurationVaultStrategy *DurationVaultStrategyCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _DurationVaultStrategy.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_DurationVaultStrategy *DurationVaultStrategySession) Paused(index uint8) (bool, error) {
	return _DurationVaultStrategy.Contract.Paused(&_DurationVaultStrategy.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_DurationVaultStrategy *DurationVaultStrategyCallerSession) Paused(index uint8) (bool, error) {
	return _DurationVaultStrategy.Contract.Paused(&_DurationVaultStrategy.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_DurationVaultStrategy *DurationVaultStrategyCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DurationVaultStrategy.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_DurationVaultStrategy *DurationVaultStrategySession) Paused0() (*big.Int, error) {
	return _DurationVaultStrategy.Contract.Paused0(&_DurationVaultStrategy.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_DurationVaultStrategy *DurationVaultStrategyCallerSession) Paused0() (*big.Int, error) {
	return _DurationVaultStrategy.Contract.Paused0(&_DurationVaultStrategy.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_DurationVaultStrategy *DurationVaultStrategyCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DurationVaultStrategy.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_DurationVaultStrategy *DurationVaultStrategySession) PauserRegistry() (common.Address, error) {
	return _DurationVaultStrategy.Contract.PauserRegistry(&_DurationVaultStrategy.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_DurationVaultStrategy *DurationVaultStrategyCallerSession) PauserRegistry() (common.Address, error) {
	return _DurationVaultStrategy.Contract.PauserRegistry(&_DurationVaultStrategy.CallOpts)
}

// RewardsCoordinator is a free data retrieval call binding the contract method 0x8a2fc4e3.
//
// Solidity: function rewardsCoordinator() view returns(address)
func (_DurationVaultStrategy *DurationVaultStrategyCaller) RewardsCoordinator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DurationVaultStrategy.contract.Call(opts, &out, "rewardsCoordinator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardsCoordinator is a free data retrieval call binding the contract method 0x8a2fc4e3.
//
// Solidity: function rewardsCoordinator() view returns(address)
func (_DurationVaultStrategy *DurationVaultStrategySession) RewardsCoordinator() (common.Address, error) {
	return _DurationVaultStrategy.Contract.RewardsCoordinator(&_DurationVaultStrategy.CallOpts)
}

// RewardsCoordinator is a free data retrieval call binding the contract method 0x8a2fc4e3.
//
// Solidity: function rewardsCoordinator() view returns(address)
func (_DurationVaultStrategy *DurationVaultStrategyCallerSession) RewardsCoordinator() (common.Address, error) {
	return _DurationVaultStrategy.Contract.RewardsCoordinator(&_DurationVaultStrategy.CallOpts)
}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares(address user) view returns(uint256)
func (_DurationVaultStrategy *DurationVaultStrategyCaller) Shares(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DurationVaultStrategy.contract.Call(opts, &out, "shares", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares(address user) view returns(uint256)
func (_DurationVaultStrategy *DurationVaultStrategySession) Shares(user common.Address) (*big.Int, error) {
	return _DurationVaultStrategy.Contract.Shares(&_DurationVaultStrategy.CallOpts, user)
}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares(address user) view returns(uint256)
func (_DurationVaultStrategy *DurationVaultStrategyCallerSession) Shares(user common.Address) (*big.Int, error) {
	return _DurationVaultStrategy.Contract.Shares(&_DurationVaultStrategy.CallOpts, user)
}

// SharesToUnderlying is a free data retrieval call binding the contract method 0xf3e73875.
//
// Solidity: function sharesToUnderlying(uint256 amountShares) view returns(uint256)
func (_DurationVaultStrategy *DurationVaultStrategyCaller) SharesToUnderlying(opts *bind.CallOpts, amountShares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DurationVaultStrategy.contract.Call(opts, &out, "sharesToUnderlying", amountShares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SharesToUnderlying is a free data retrieval call binding the contract method 0xf3e73875.
//
// Solidity: function sharesToUnderlying(uint256 amountShares) view returns(uint256)
func (_DurationVaultStrategy *DurationVaultStrategySession) SharesToUnderlying(amountShares *big.Int) (*big.Int, error) {
	return _DurationVaultStrategy.Contract.SharesToUnderlying(&_DurationVaultStrategy.CallOpts, amountShares)
}

// SharesToUnderlying is a free data retrieval call binding the contract method 0xf3e73875.
//
// Solidity: function sharesToUnderlying(uint256 amountShares) view returns(uint256)
func (_DurationVaultStrategy *DurationVaultStrategyCallerSession) SharesToUnderlying(amountShares *big.Int) (*big.Int, error) {
	return _DurationVaultStrategy.Contract.SharesToUnderlying(&_DurationVaultStrategy.CallOpts, amountShares)
}

// SharesToUnderlyingView is a free data retrieval call binding the contract method 0x7a8b2637.
//
// Solidity: function sharesToUnderlyingView(uint256 amountShares) view returns(uint256)
func (_DurationVaultStrategy *DurationVaultStrategyCaller) SharesToUnderlyingView(opts *bind.CallOpts, amountShares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DurationVaultStrategy.contract.Call(opts, &out, "sharesToUnderlyingView", amountShares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SharesToUnderlyingView is a free data retrieval call binding the contract method 0x7a8b2637.
//
// Solidity: function sharesToUnderlyingView(uint256 amountShares) view returns(uint256)
func (_DurationVaultStrategy *DurationVaultStrategySession) SharesToUnderlyingView(amountShares *big.Int) (*big.Int, error) {
	return _DurationVaultStrategy.Contract.SharesToUnderlyingView(&_DurationVaultStrategy.CallOpts, amountShares)
}

// SharesToUnderlyingView is a free data retrieval call binding the contract method 0x7a8b2637.
//
// Solidity: function sharesToUnderlyingView(uint256 amountShares) view returns(uint256)
func (_DurationVaultStrategy *DurationVaultStrategyCallerSession) SharesToUnderlyingView(amountShares *big.Int) (*big.Int, error) {
	return _DurationVaultStrategy.Contract.SharesToUnderlyingView(&_DurationVaultStrategy.CallOpts, amountShares)
}

// StakeCap is a free data retrieval call binding the contract method 0xba28fd2e.
//
// Solidity: function stakeCap() view returns(uint256)
func (_DurationVaultStrategy *DurationVaultStrategyCaller) StakeCap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DurationVaultStrategy.contract.Call(opts, &out, "stakeCap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeCap is a free data retrieval call binding the contract method 0xba28fd2e.
//
// Solidity: function stakeCap() view returns(uint256)
func (_DurationVaultStrategy *DurationVaultStrategySession) StakeCap() (*big.Int, error) {
	return _DurationVaultStrategy.Contract.StakeCap(&_DurationVaultStrategy.CallOpts)
}

// StakeCap is a free data retrieval call binding the contract method 0xba28fd2e.
//
// Solidity: function stakeCap() view returns(uint256)
func (_DurationVaultStrategy *DurationVaultStrategyCallerSession) StakeCap() (*big.Int, error) {
	return _DurationVaultStrategy.Contract.StakeCap(&_DurationVaultStrategy.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_DurationVaultStrategy *DurationVaultStrategyCaller) State(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _DurationVaultStrategy.contract.Call(opts, &out, "state")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_DurationVaultStrategy *DurationVaultStrategySession) State() (uint8, error) {
	return _DurationVaultStrategy.Contract.State(&_DurationVaultStrategy.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_DurationVaultStrategy *DurationVaultStrategyCallerSession) State() (uint8, error) {
	return _DurationVaultStrategy.Contract.State(&_DurationVaultStrategy.CallOpts)
}

// StrategyFactory is a free data retrieval call binding the contract method 0x9ef35710.
//
// Solidity: function strategyFactory() view returns(address)
func (_DurationVaultStrategy *DurationVaultStrategyCaller) StrategyFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DurationVaultStrategy.contract.Call(opts, &out, "strategyFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StrategyFactory is a free data retrieval call binding the contract method 0x9ef35710.
//
// Solidity: function strategyFactory() view returns(address)
func (_DurationVaultStrategy *DurationVaultStrategySession) StrategyFactory() (common.Address, error) {
	return _DurationVaultStrategy.Contract.StrategyFactory(&_DurationVaultStrategy.CallOpts)
}

// StrategyFactory is a free data retrieval call binding the contract method 0x9ef35710.
//
// Solidity: function strategyFactory() view returns(address)
func (_DurationVaultStrategy *DurationVaultStrategyCallerSession) StrategyFactory() (common.Address, error) {
	return _DurationVaultStrategy.Contract.StrategyFactory(&_DurationVaultStrategy.CallOpts)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_DurationVaultStrategy *DurationVaultStrategyCaller) StrategyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DurationVaultStrategy.contract.Call(opts, &out, "strategyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_DurationVaultStrategy *DurationVaultStrategySession) StrategyManager() (common.Address, error) {
	return _DurationVaultStrategy.Contract.StrategyManager(&_DurationVaultStrategy.CallOpts)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_DurationVaultStrategy *DurationVaultStrategyCallerSession) StrategyManager() (common.Address, error) {
	return _DurationVaultStrategy.Contract.StrategyManager(&_DurationVaultStrategy.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_DurationVaultStrategy *DurationVaultStrategyCaller) TotalShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DurationVaultStrategy.contract.Call(opts, &out, "totalShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_DurationVaultStrategy *DurationVaultStrategySession) TotalShares() (*big.Int, error) {
	return _DurationVaultStrategy.Contract.TotalShares(&_DurationVaultStrategy.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_DurationVaultStrategy *DurationVaultStrategyCallerSession) TotalShares() (*big.Int, error) {
	return _DurationVaultStrategy.Contract.TotalShares(&_DurationVaultStrategy.CallOpts)
}

// UnderlyingToShares is a free data retrieval call binding the contract method 0x8c871019.
//
// Solidity: function underlyingToShares(uint256 amountUnderlying) view returns(uint256)
func (_DurationVaultStrategy *DurationVaultStrategyCaller) UnderlyingToShares(opts *bind.CallOpts, amountUnderlying *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DurationVaultStrategy.contract.Call(opts, &out, "underlyingToShares", amountUnderlying)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnderlyingToShares is a free data retrieval call binding the contract method 0x8c871019.
//
// Solidity: function underlyingToShares(uint256 amountUnderlying) view returns(uint256)
func (_DurationVaultStrategy *DurationVaultStrategySession) UnderlyingToShares(amountUnderlying *big.Int) (*big.Int, error) {
	return _DurationVaultStrategy.Contract.UnderlyingToShares(&_DurationVaultStrategy.CallOpts, amountUnderlying)
}

// UnderlyingToShares is a free data retrieval call binding the contract method 0x8c871019.
//
// Solidity: function underlyingToShares(uint256 amountUnderlying) view returns(uint256)
func (_DurationVaultStrategy *DurationVaultStrategyCallerSession) UnderlyingToShares(amountUnderlying *big.Int) (*big.Int, error) {
	return _DurationVaultStrategy.Contract.UnderlyingToShares(&_DurationVaultStrategy.CallOpts, amountUnderlying)
}

// UnderlyingToSharesView is a free data retrieval call binding the contract method 0xe3dae51c.
//
// Solidity: function underlyingToSharesView(uint256 amountUnderlying) view returns(uint256)
func (_DurationVaultStrategy *DurationVaultStrategyCaller) UnderlyingToSharesView(opts *bind.CallOpts, amountUnderlying *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DurationVaultStrategy.contract.Call(opts, &out, "underlyingToSharesView", amountUnderlying)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnderlyingToSharesView is a free data retrieval call binding the contract method 0xe3dae51c.
//
// Solidity: function underlyingToSharesView(uint256 amountUnderlying) view returns(uint256)
func (_DurationVaultStrategy *DurationVaultStrategySession) UnderlyingToSharesView(amountUnderlying *big.Int) (*big.Int, error) {
	return _DurationVaultStrategy.Contract.UnderlyingToSharesView(&_DurationVaultStrategy.CallOpts, amountUnderlying)
}

// UnderlyingToSharesView is a free data retrieval call binding the contract method 0xe3dae51c.
//
// Solidity: function underlyingToSharesView(uint256 amountUnderlying) view returns(uint256)
func (_DurationVaultStrategy *DurationVaultStrategyCallerSession) UnderlyingToSharesView(amountUnderlying *big.Int) (*big.Int, error) {
	return _DurationVaultStrategy.Contract.UnderlyingToSharesView(&_DurationVaultStrategy.CallOpts, amountUnderlying)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_DurationVaultStrategy *DurationVaultStrategyCaller) UnderlyingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DurationVaultStrategy.contract.Call(opts, &out, "underlyingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_DurationVaultStrategy *DurationVaultStrategySession) UnderlyingToken() (common.Address, error) {
	return _DurationVaultStrategy.Contract.UnderlyingToken(&_DurationVaultStrategy.CallOpts)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_DurationVaultStrategy *DurationVaultStrategyCallerSession) UnderlyingToken() (common.Address, error) {
	return _DurationVaultStrategy.Contract.UnderlyingToken(&_DurationVaultStrategy.CallOpts)
}

// UnlockAt is a free data retrieval call binding the contract method 0xaa5dec6f.
//
// Solidity: function unlockAt() view returns(uint32)
func (_DurationVaultStrategy *DurationVaultStrategyCaller) UnlockAt(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _DurationVaultStrategy.contract.Call(opts, &out, "unlockAt")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// UnlockAt is a free data retrieval call binding the contract method 0xaa5dec6f.
//
// Solidity: function unlockAt() view returns(uint32)
func (_DurationVaultStrategy *DurationVaultStrategySession) UnlockAt() (uint32, error) {
	return _DurationVaultStrategy.Contract.UnlockAt(&_DurationVaultStrategy.CallOpts)
}

// UnlockAt is a free data retrieval call binding the contract method 0xaa5dec6f.
//
// Solidity: function unlockAt() view returns(uint32)
func (_DurationVaultStrategy *DurationVaultStrategyCallerSession) UnlockAt() (uint32, error) {
	return _DurationVaultStrategy.Contract.UnlockAt(&_DurationVaultStrategy.CallOpts)
}

// UnlockTimestamp is a free data retrieval call binding the contract method 0xaa082a9d.
//
// Solidity: function unlockTimestamp() view returns(uint32)
func (_DurationVaultStrategy *DurationVaultStrategyCaller) UnlockTimestamp(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _DurationVaultStrategy.contract.Call(opts, &out, "unlockTimestamp")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// UnlockTimestamp is a free data retrieval call binding the contract method 0xaa082a9d.
//
// Solidity: function unlockTimestamp() view returns(uint32)
func (_DurationVaultStrategy *DurationVaultStrategySession) UnlockTimestamp() (uint32, error) {
	return _DurationVaultStrategy.Contract.UnlockTimestamp(&_DurationVaultStrategy.CallOpts)
}

// UnlockTimestamp is a free data retrieval call binding the contract method 0xaa082a9d.
//
// Solidity: function unlockTimestamp() view returns(uint32)
func (_DurationVaultStrategy *DurationVaultStrategyCallerSession) UnlockTimestamp() (uint32, error) {
	return _DurationVaultStrategy.Contract.UnlockTimestamp(&_DurationVaultStrategy.CallOpts)
}

// UserUnderlyingView is a free data retrieval call binding the contract method 0x553ca5f8.
//
// Solidity: function userUnderlyingView(address user) view returns(uint256)
func (_DurationVaultStrategy *DurationVaultStrategyCaller) UserUnderlyingView(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DurationVaultStrategy.contract.Call(opts, &out, "userUnderlyingView", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserUnderlyingView is a free data retrieval call binding the contract method 0x553ca5f8.
//
// Solidity: function userUnderlyingView(address user) view returns(uint256)
func (_DurationVaultStrategy *DurationVaultStrategySession) UserUnderlyingView(user common.Address) (*big.Int, error) {
	return _DurationVaultStrategy.Contract.UserUnderlyingView(&_DurationVaultStrategy.CallOpts, user)
}

// UserUnderlyingView is a free data retrieval call binding the contract method 0x553ca5f8.
//
// Solidity: function userUnderlyingView(address user) view returns(uint256)
func (_DurationVaultStrategy *DurationVaultStrategyCallerSession) UserUnderlyingView(user common.Address) (*big.Int, error) {
	return _DurationVaultStrategy.Contract.UserUnderlyingView(&_DurationVaultStrategy.CallOpts, user)
}

// VaultAdmin is a free data retrieval call binding the contract method 0xe7f6f225.
//
// Solidity: function vaultAdmin() view returns(address)
func (_DurationVaultStrategy *DurationVaultStrategyCaller) VaultAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DurationVaultStrategy.contract.Call(opts, &out, "vaultAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VaultAdmin is a free data retrieval call binding the contract method 0xe7f6f225.
//
// Solidity: function vaultAdmin() view returns(address)
func (_DurationVaultStrategy *DurationVaultStrategySession) VaultAdmin() (common.Address, error) {
	return _DurationVaultStrategy.Contract.VaultAdmin(&_DurationVaultStrategy.CallOpts)
}

// VaultAdmin is a free data retrieval call binding the contract method 0xe7f6f225.
//
// Solidity: function vaultAdmin() view returns(address)
func (_DurationVaultStrategy *DurationVaultStrategyCallerSession) VaultAdmin() (common.Address, error) {
	return _DurationVaultStrategy.Contract.VaultAdmin(&_DurationVaultStrategy.CallOpts)
}

// WithdrawalsOpen is a free data retrieval call binding the contract method 0x94aad677.
//
// Solidity: function withdrawalsOpen() view returns(bool)
func (_DurationVaultStrategy *DurationVaultStrategyCaller) WithdrawalsOpen(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DurationVaultStrategy.contract.Call(opts, &out, "withdrawalsOpen")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WithdrawalsOpen is a free data retrieval call binding the contract method 0x94aad677.
//
// Solidity: function withdrawalsOpen() view returns(bool)
func (_DurationVaultStrategy *DurationVaultStrategySession) WithdrawalsOpen() (bool, error) {
	return _DurationVaultStrategy.Contract.WithdrawalsOpen(&_DurationVaultStrategy.CallOpts)
}

// WithdrawalsOpen is a free data retrieval call binding the contract method 0x94aad677.
//
// Solidity: function withdrawalsOpen() view returns(bool)
func (_DurationVaultStrategy *DurationVaultStrategyCallerSession) WithdrawalsOpen() (bool, error) {
	return _DurationVaultStrategy.Contract.WithdrawalsOpen(&_DurationVaultStrategy.CallOpts)
}

// AdvanceToWithdrawals is a paid mutator transaction binding the contract method 0x6325f655.
//
// Solidity: function advanceToWithdrawals() returns()
func (_DurationVaultStrategy *DurationVaultStrategyTransactor) AdvanceToWithdrawals(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DurationVaultStrategy.contract.Transact(opts, "advanceToWithdrawals")
}

// AdvanceToWithdrawals is a paid mutator transaction binding the contract method 0x6325f655.
//
// Solidity: function advanceToWithdrawals() returns()
func (_DurationVaultStrategy *DurationVaultStrategySession) AdvanceToWithdrawals() (*types.Transaction, error) {
	return _DurationVaultStrategy.Contract.AdvanceToWithdrawals(&_DurationVaultStrategy.TransactOpts)
}

// AdvanceToWithdrawals is a paid mutator transaction binding the contract method 0x6325f655.
//
// Solidity: function advanceToWithdrawals() returns()
func (_DurationVaultStrategy *DurationVaultStrategyTransactorSession) AdvanceToWithdrawals() (*types.Transaction, error) {
	return _DurationVaultStrategy.Contract.AdvanceToWithdrawals(&_DurationVaultStrategy.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns(uint256 newShares)
func (_DurationVaultStrategy *DurationVaultStrategyTransactor) Deposit(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DurationVaultStrategy.contract.Transact(opts, "deposit", token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns(uint256 newShares)
func (_DurationVaultStrategy *DurationVaultStrategySession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DurationVaultStrategy.Contract.Deposit(&_DurationVaultStrategy.TransactOpts, token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns(uint256 newShares)
func (_DurationVaultStrategy *DurationVaultStrategyTransactorSession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DurationVaultStrategy.Contract.Deposit(&_DurationVaultStrategy.TransactOpts, token, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0xc2cca26d.
//
// Solidity: function initialize((address,address,address,uint32,uint256,uint256,string,(address,uint32),bytes,address,string) config) returns()
func (_DurationVaultStrategy *DurationVaultStrategyTransactor) Initialize(opts *bind.TransactOpts, config IDurationVaultStrategyTypesVaultConfig) (*types.Transaction, error) {
	return _DurationVaultStrategy.contract.Transact(opts, "initialize", config)
}

// Initialize is a paid mutator transaction binding the contract method 0xc2cca26d.
//
// Solidity: function initialize((address,address,address,uint32,uint256,uint256,string,(address,uint32),bytes,address,string) config) returns()
func (_DurationVaultStrategy *DurationVaultStrategySession) Initialize(config IDurationVaultStrategyTypesVaultConfig) (*types.Transaction, error) {
	return _DurationVaultStrategy.Contract.Initialize(&_DurationVaultStrategy.TransactOpts, config)
}

// Initialize is a paid mutator transaction binding the contract method 0xc2cca26d.
//
// Solidity: function initialize((address,address,address,uint32,uint256,uint256,string,(address,uint32),bytes,address,string) config) returns()
func (_DurationVaultStrategy *DurationVaultStrategyTransactorSession) Initialize(config IDurationVaultStrategyTypesVaultConfig) (*types.Transaction, error) {
	return _DurationVaultStrategy.Contract.Initialize(&_DurationVaultStrategy.TransactOpts, config)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _underlyingToken) returns()
func (_DurationVaultStrategy *DurationVaultStrategyTransactor) Initialize0(opts *bind.TransactOpts, _underlyingToken common.Address) (*types.Transaction, error) {
	return _DurationVaultStrategy.contract.Transact(opts, "initialize0", _underlyingToken)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _underlyingToken) returns()
func (_DurationVaultStrategy *DurationVaultStrategySession) Initialize0(_underlyingToken common.Address) (*types.Transaction, error) {
	return _DurationVaultStrategy.Contract.Initialize0(&_DurationVaultStrategy.TransactOpts, _underlyingToken)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _underlyingToken) returns()
func (_DurationVaultStrategy *DurationVaultStrategyTransactorSession) Initialize0(_underlyingToken common.Address) (*types.Transaction, error) {
	return _DurationVaultStrategy.Contract.Initialize0(&_DurationVaultStrategy.TransactOpts, _underlyingToken)
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_DurationVaultStrategy *DurationVaultStrategyTransactor) Lock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DurationVaultStrategy.contract.Transact(opts, "lock")
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_DurationVaultStrategy *DurationVaultStrategySession) Lock() (*types.Transaction, error) {
	return _DurationVaultStrategy.Contract.Lock(&_DurationVaultStrategy.TransactOpts)
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_DurationVaultStrategy *DurationVaultStrategyTransactorSession) Lock() (*types.Transaction, error) {
	return _DurationVaultStrategy.Contract.Lock(&_DurationVaultStrategy.TransactOpts)
}

// MarkMatured is a paid mutator transaction binding the contract method 0x6d8690a9.
//
// Solidity: function markMatured() returns()
func (_DurationVaultStrategy *DurationVaultStrategyTransactor) MarkMatured(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DurationVaultStrategy.contract.Transact(opts, "markMatured")
}

// MarkMatured is a paid mutator transaction binding the contract method 0x6d8690a9.
//
// Solidity: function markMatured() returns()
func (_DurationVaultStrategy *DurationVaultStrategySession) MarkMatured() (*types.Transaction, error) {
	return _DurationVaultStrategy.Contract.MarkMatured(&_DurationVaultStrategy.TransactOpts)
}

// MarkMatured is a paid mutator transaction binding the contract method 0x6d8690a9.
//
// Solidity: function markMatured() returns()
func (_DurationVaultStrategy *DurationVaultStrategyTransactorSession) MarkMatured() (*types.Transaction, error) {
	return _DurationVaultStrategy.Contract.MarkMatured(&_DurationVaultStrategy.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_DurationVaultStrategy *DurationVaultStrategyTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _DurationVaultStrategy.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_DurationVaultStrategy *DurationVaultStrategySession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _DurationVaultStrategy.Contract.Pause(&_DurationVaultStrategy.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_DurationVaultStrategy *DurationVaultStrategyTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _DurationVaultStrategy.Contract.Pause(&_DurationVaultStrategy.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_DurationVaultStrategy *DurationVaultStrategyTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DurationVaultStrategy.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_DurationVaultStrategy *DurationVaultStrategySession) PauseAll() (*types.Transaction, error) {
	return _DurationVaultStrategy.Contract.PauseAll(&_DurationVaultStrategy.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_DurationVaultStrategy *DurationVaultStrategyTransactorSession) PauseAll() (*types.Transaction, error) {
	return _DurationVaultStrategy.Contract.PauseAll(&_DurationVaultStrategy.TransactOpts)
}

// SetRewardsClaimer is a paid mutator transaction binding the contract method 0xb501d660.
//
// Solidity: function setRewardsClaimer(address claimer) returns()
func (_DurationVaultStrategy *DurationVaultStrategyTransactor) SetRewardsClaimer(opts *bind.TransactOpts, claimer common.Address) (*types.Transaction, error) {
	return _DurationVaultStrategy.contract.Transact(opts, "setRewardsClaimer", claimer)
}

// SetRewardsClaimer is a paid mutator transaction binding the contract method 0xb501d660.
//
// Solidity: function setRewardsClaimer(address claimer) returns()
func (_DurationVaultStrategy *DurationVaultStrategySession) SetRewardsClaimer(claimer common.Address) (*types.Transaction, error) {
	return _DurationVaultStrategy.Contract.SetRewardsClaimer(&_DurationVaultStrategy.TransactOpts, claimer)
}

// SetRewardsClaimer is a paid mutator transaction binding the contract method 0xb501d660.
//
// Solidity: function setRewardsClaimer(address claimer) returns()
func (_DurationVaultStrategy *DurationVaultStrategyTransactorSession) SetRewardsClaimer(claimer common.Address) (*types.Transaction, error) {
	return _DurationVaultStrategy.Contract.SetRewardsClaimer(&_DurationVaultStrategy.TransactOpts, claimer)
}

// SetTVLLimits is a paid mutator transaction binding the contract method 0x11c70c9d.
//
// Solidity: function setTVLLimits(uint256 newMaxPerDeposit, uint256 newMaxTotalDeposits) returns()
func (_DurationVaultStrategy *DurationVaultStrategyTransactor) SetTVLLimits(opts *bind.TransactOpts, newMaxPerDeposit *big.Int, newMaxTotalDeposits *big.Int) (*types.Transaction, error) {
	return _DurationVaultStrategy.contract.Transact(opts, "setTVLLimits", newMaxPerDeposit, newMaxTotalDeposits)
}

// SetTVLLimits is a paid mutator transaction binding the contract method 0x11c70c9d.
//
// Solidity: function setTVLLimits(uint256 newMaxPerDeposit, uint256 newMaxTotalDeposits) returns()
func (_DurationVaultStrategy *DurationVaultStrategySession) SetTVLLimits(newMaxPerDeposit *big.Int, newMaxTotalDeposits *big.Int) (*types.Transaction, error) {
	return _DurationVaultStrategy.Contract.SetTVLLimits(&_DurationVaultStrategy.TransactOpts, newMaxPerDeposit, newMaxTotalDeposits)
}

// SetTVLLimits is a paid mutator transaction binding the contract method 0x11c70c9d.
//
// Solidity: function setTVLLimits(uint256 newMaxPerDeposit, uint256 newMaxTotalDeposits) returns()
func (_DurationVaultStrategy *DurationVaultStrategyTransactorSession) SetTVLLimits(newMaxPerDeposit *big.Int, newMaxTotalDeposits *big.Int) (*types.Transaction, error) {
	return _DurationVaultStrategy.Contract.SetTVLLimits(&_DurationVaultStrategy.TransactOpts, newMaxPerDeposit, newMaxTotalDeposits)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_DurationVaultStrategy *DurationVaultStrategyTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _DurationVaultStrategy.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_DurationVaultStrategy *DurationVaultStrategySession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _DurationVaultStrategy.Contract.Unpause(&_DurationVaultStrategy.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_DurationVaultStrategy *DurationVaultStrategyTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _DurationVaultStrategy.Contract.Unpause(&_DurationVaultStrategy.TransactOpts, newPausedStatus)
}

// UpdateDelegationApprover is a paid mutator transaction binding the contract method 0xb4e20f13.
//
// Solidity: function updateDelegationApprover(address newDelegationApprover) returns()
func (_DurationVaultStrategy *DurationVaultStrategyTransactor) UpdateDelegationApprover(opts *bind.TransactOpts, newDelegationApprover common.Address) (*types.Transaction, error) {
	return _DurationVaultStrategy.contract.Transact(opts, "updateDelegationApprover", newDelegationApprover)
}

// UpdateDelegationApprover is a paid mutator transaction binding the contract method 0xb4e20f13.
//
// Solidity: function updateDelegationApprover(address newDelegationApprover) returns()
func (_DurationVaultStrategy *DurationVaultStrategySession) UpdateDelegationApprover(newDelegationApprover common.Address) (*types.Transaction, error) {
	return _DurationVaultStrategy.Contract.UpdateDelegationApprover(&_DurationVaultStrategy.TransactOpts, newDelegationApprover)
}

// UpdateDelegationApprover is a paid mutator transaction binding the contract method 0xb4e20f13.
//
// Solidity: function updateDelegationApprover(address newDelegationApprover) returns()
func (_DurationVaultStrategy *DurationVaultStrategyTransactorSession) UpdateDelegationApprover(newDelegationApprover common.Address) (*types.Transaction, error) {
	return _DurationVaultStrategy.Contract.UpdateDelegationApprover(&_DurationVaultStrategy.TransactOpts, newDelegationApprover)
}

// UpdateMetadataURI is a paid mutator transaction binding the contract method 0x53fd3e81.
//
// Solidity: function updateMetadataURI(string newMetadataURI) returns()
func (_DurationVaultStrategy *DurationVaultStrategyTransactor) UpdateMetadataURI(opts *bind.TransactOpts, newMetadataURI string) (*types.Transaction, error) {
	return _DurationVaultStrategy.contract.Transact(opts, "updateMetadataURI", newMetadataURI)
}

// UpdateMetadataURI is a paid mutator transaction binding the contract method 0x53fd3e81.
//
// Solidity: function updateMetadataURI(string newMetadataURI) returns()
func (_DurationVaultStrategy *DurationVaultStrategySession) UpdateMetadataURI(newMetadataURI string) (*types.Transaction, error) {
	return _DurationVaultStrategy.Contract.UpdateMetadataURI(&_DurationVaultStrategy.TransactOpts, newMetadataURI)
}

// UpdateMetadataURI is a paid mutator transaction binding the contract method 0x53fd3e81.
//
// Solidity: function updateMetadataURI(string newMetadataURI) returns()
func (_DurationVaultStrategy *DurationVaultStrategyTransactorSession) UpdateMetadataURI(newMetadataURI string) (*types.Transaction, error) {
	return _DurationVaultStrategy.Contract.UpdateMetadataURI(&_DurationVaultStrategy.TransactOpts, newMetadataURI)
}

// UpdateOperatorMetadataURI is a paid mutator transaction binding the contract method 0x99be81c8.
//
// Solidity: function updateOperatorMetadataURI(string newOperatorMetadataURI) returns()
func (_DurationVaultStrategy *DurationVaultStrategyTransactor) UpdateOperatorMetadataURI(opts *bind.TransactOpts, newOperatorMetadataURI string) (*types.Transaction, error) {
	return _DurationVaultStrategy.contract.Transact(opts, "updateOperatorMetadataURI", newOperatorMetadataURI)
}

// UpdateOperatorMetadataURI is a paid mutator transaction binding the contract method 0x99be81c8.
//
// Solidity: function updateOperatorMetadataURI(string newOperatorMetadataURI) returns()
func (_DurationVaultStrategy *DurationVaultStrategySession) UpdateOperatorMetadataURI(newOperatorMetadataURI string) (*types.Transaction, error) {
	return _DurationVaultStrategy.Contract.UpdateOperatorMetadataURI(&_DurationVaultStrategy.TransactOpts, newOperatorMetadataURI)
}

// UpdateOperatorMetadataURI is a paid mutator transaction binding the contract method 0x99be81c8.
//
// Solidity: function updateOperatorMetadataURI(string newOperatorMetadataURI) returns()
func (_DurationVaultStrategy *DurationVaultStrategyTransactorSession) UpdateOperatorMetadataURI(newOperatorMetadataURI string) (*types.Transaction, error) {
	return _DurationVaultStrategy.Contract.UpdateOperatorMetadataURI(&_DurationVaultStrategy.TransactOpts, newOperatorMetadataURI)
}

// UpdateTVLLimits is a paid mutator transaction binding the contract method 0xaf6eb2be.
//
// Solidity: function updateTVLLimits(uint256 newMaxPerDeposit, uint256 newStakeCap) returns()
func (_DurationVaultStrategy *DurationVaultStrategyTransactor) UpdateTVLLimits(opts *bind.TransactOpts, newMaxPerDeposit *big.Int, newStakeCap *big.Int) (*types.Transaction, error) {
	return _DurationVaultStrategy.contract.Transact(opts, "updateTVLLimits", newMaxPerDeposit, newStakeCap)
}

// UpdateTVLLimits is a paid mutator transaction binding the contract method 0xaf6eb2be.
//
// Solidity: function updateTVLLimits(uint256 newMaxPerDeposit, uint256 newStakeCap) returns()
func (_DurationVaultStrategy *DurationVaultStrategySession) UpdateTVLLimits(newMaxPerDeposit *big.Int, newStakeCap *big.Int) (*types.Transaction, error) {
	return _DurationVaultStrategy.Contract.UpdateTVLLimits(&_DurationVaultStrategy.TransactOpts, newMaxPerDeposit, newStakeCap)
}

// UpdateTVLLimits is a paid mutator transaction binding the contract method 0xaf6eb2be.
//
// Solidity: function updateTVLLimits(uint256 newMaxPerDeposit, uint256 newStakeCap) returns()
func (_DurationVaultStrategy *DurationVaultStrategyTransactorSession) UpdateTVLLimits(newMaxPerDeposit *big.Int, newStakeCap *big.Int) (*types.Transaction, error) {
	return _DurationVaultStrategy.Contract.UpdateTVLLimits(&_DurationVaultStrategy.TransactOpts, newMaxPerDeposit, newStakeCap)
}

// UserUnderlying is a paid mutator transaction binding the contract method 0x8f6a6240.
//
// Solidity: function userUnderlying(address user) returns(uint256)
func (_DurationVaultStrategy *DurationVaultStrategyTransactor) UserUnderlying(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _DurationVaultStrategy.contract.Transact(opts, "userUnderlying", user)
}

// UserUnderlying is a paid mutator transaction binding the contract method 0x8f6a6240.
//
// Solidity: function userUnderlying(address user) returns(uint256)
func (_DurationVaultStrategy *DurationVaultStrategySession) UserUnderlying(user common.Address) (*types.Transaction, error) {
	return _DurationVaultStrategy.Contract.UserUnderlying(&_DurationVaultStrategy.TransactOpts, user)
}

// UserUnderlying is a paid mutator transaction binding the contract method 0x8f6a6240.
//
// Solidity: function userUnderlying(address user) returns(uint256)
func (_DurationVaultStrategy *DurationVaultStrategyTransactorSession) UserUnderlying(user common.Address) (*types.Transaction, error) {
	return _DurationVaultStrategy.Contract.UserUnderlying(&_DurationVaultStrategy.TransactOpts, user)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address recipient, address token, uint256 amountShares) returns(uint256 amountOut)
func (_DurationVaultStrategy *DurationVaultStrategyTransactor) Withdraw(opts *bind.TransactOpts, recipient common.Address, token common.Address, amountShares *big.Int) (*types.Transaction, error) {
	return _DurationVaultStrategy.contract.Transact(opts, "withdraw", recipient, token, amountShares)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address recipient, address token, uint256 amountShares) returns(uint256 amountOut)
func (_DurationVaultStrategy *DurationVaultStrategySession) Withdraw(recipient common.Address, token common.Address, amountShares *big.Int) (*types.Transaction, error) {
	return _DurationVaultStrategy.Contract.Withdraw(&_DurationVaultStrategy.TransactOpts, recipient, token, amountShares)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address recipient, address token, uint256 amountShares) returns(uint256 amountOut)
func (_DurationVaultStrategy *DurationVaultStrategyTransactorSession) Withdraw(recipient common.Address, token common.Address, amountShares *big.Int) (*types.Transaction, error) {
	return _DurationVaultStrategy.Contract.Withdraw(&_DurationVaultStrategy.TransactOpts, recipient, token, amountShares)
}

// DurationVaultStrategyDeallocateAttemptedIterator is returned from FilterDeallocateAttempted and is used to iterate over the raw logs and unpacked data for DeallocateAttempted events raised by the DurationVaultStrategy contract.
type DurationVaultStrategyDeallocateAttemptedIterator struct {
	Event *DurationVaultStrategyDeallocateAttempted // Event containing the contract specifics and raw log

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
func (it *DurationVaultStrategyDeallocateAttemptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DurationVaultStrategyDeallocateAttempted)
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
		it.Event = new(DurationVaultStrategyDeallocateAttempted)
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
func (it *DurationVaultStrategyDeallocateAttemptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DurationVaultStrategyDeallocateAttemptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DurationVaultStrategyDeallocateAttempted represents a DeallocateAttempted event raised by the DurationVaultStrategy contract.
type DurationVaultStrategyDeallocateAttempted struct {
	Success bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeallocateAttempted is a free log retrieval operation binding the contract event 0x72f957da7daaea6b52e4ff7820cb464206fd51e9f502f3027f45b5017caf4c8b.
//
// Solidity: event DeallocateAttempted(bool success)
func (_DurationVaultStrategy *DurationVaultStrategyFilterer) FilterDeallocateAttempted(opts *bind.FilterOpts) (*DurationVaultStrategyDeallocateAttemptedIterator, error) {

	logs, sub, err := _DurationVaultStrategy.contract.FilterLogs(opts, "DeallocateAttempted")
	if err != nil {
		return nil, err
	}
	return &DurationVaultStrategyDeallocateAttemptedIterator{contract: _DurationVaultStrategy.contract, event: "DeallocateAttempted", logs: logs, sub: sub}, nil
}

// WatchDeallocateAttempted is a free log subscription operation binding the contract event 0x72f957da7daaea6b52e4ff7820cb464206fd51e9f502f3027f45b5017caf4c8b.
//
// Solidity: event DeallocateAttempted(bool success)
func (_DurationVaultStrategy *DurationVaultStrategyFilterer) WatchDeallocateAttempted(opts *bind.WatchOpts, sink chan<- *DurationVaultStrategyDeallocateAttempted) (event.Subscription, error) {

	logs, sub, err := _DurationVaultStrategy.contract.WatchLogs(opts, "DeallocateAttempted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DurationVaultStrategyDeallocateAttempted)
				if err := _DurationVaultStrategy.contract.UnpackLog(event, "DeallocateAttempted", log); err != nil {
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

// ParseDeallocateAttempted is a log parse operation binding the contract event 0x72f957da7daaea6b52e4ff7820cb464206fd51e9f502f3027f45b5017caf4c8b.
//
// Solidity: event DeallocateAttempted(bool success)
func (_DurationVaultStrategy *DurationVaultStrategyFilterer) ParseDeallocateAttempted(log types.Log) (*DurationVaultStrategyDeallocateAttempted, error) {
	event := new(DurationVaultStrategyDeallocateAttempted)
	if err := _DurationVaultStrategy.contract.UnpackLog(event, "DeallocateAttempted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DurationVaultStrategyDeregisterAttemptedIterator is returned from FilterDeregisterAttempted and is used to iterate over the raw logs and unpacked data for DeregisterAttempted events raised by the DurationVaultStrategy contract.
type DurationVaultStrategyDeregisterAttemptedIterator struct {
	Event *DurationVaultStrategyDeregisterAttempted // Event containing the contract specifics and raw log

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
func (it *DurationVaultStrategyDeregisterAttemptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DurationVaultStrategyDeregisterAttempted)
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
		it.Event = new(DurationVaultStrategyDeregisterAttempted)
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
func (it *DurationVaultStrategyDeregisterAttemptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DurationVaultStrategyDeregisterAttemptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DurationVaultStrategyDeregisterAttempted represents a DeregisterAttempted event raised by the DurationVaultStrategy contract.
type DurationVaultStrategyDeregisterAttempted struct {
	Success bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeregisterAttempted is a free log retrieval operation binding the contract event 0xd0791dbc9180cb64588d7eb7658a1022dcf734b8825eb7eec68bd9516872d168.
//
// Solidity: event DeregisterAttempted(bool success)
func (_DurationVaultStrategy *DurationVaultStrategyFilterer) FilterDeregisterAttempted(opts *bind.FilterOpts) (*DurationVaultStrategyDeregisterAttemptedIterator, error) {

	logs, sub, err := _DurationVaultStrategy.contract.FilterLogs(opts, "DeregisterAttempted")
	if err != nil {
		return nil, err
	}
	return &DurationVaultStrategyDeregisterAttemptedIterator{contract: _DurationVaultStrategy.contract, event: "DeregisterAttempted", logs: logs, sub: sub}, nil
}

// WatchDeregisterAttempted is a free log subscription operation binding the contract event 0xd0791dbc9180cb64588d7eb7658a1022dcf734b8825eb7eec68bd9516872d168.
//
// Solidity: event DeregisterAttempted(bool success)
func (_DurationVaultStrategy *DurationVaultStrategyFilterer) WatchDeregisterAttempted(opts *bind.WatchOpts, sink chan<- *DurationVaultStrategyDeregisterAttempted) (event.Subscription, error) {

	logs, sub, err := _DurationVaultStrategy.contract.WatchLogs(opts, "DeregisterAttempted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DurationVaultStrategyDeregisterAttempted)
				if err := _DurationVaultStrategy.contract.UnpackLog(event, "DeregisterAttempted", log); err != nil {
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

// ParseDeregisterAttempted is a log parse operation binding the contract event 0xd0791dbc9180cb64588d7eb7658a1022dcf734b8825eb7eec68bd9516872d168.
//
// Solidity: event DeregisterAttempted(bool success)
func (_DurationVaultStrategy *DurationVaultStrategyFilterer) ParseDeregisterAttempted(log types.Log) (*DurationVaultStrategyDeregisterAttempted, error) {
	event := new(DurationVaultStrategyDeregisterAttempted)
	if err := _DurationVaultStrategy.contract.UnpackLog(event, "DeregisterAttempted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DurationVaultStrategyExchangeRateEmittedIterator is returned from FilterExchangeRateEmitted and is used to iterate over the raw logs and unpacked data for ExchangeRateEmitted events raised by the DurationVaultStrategy contract.
type DurationVaultStrategyExchangeRateEmittedIterator struct {
	Event *DurationVaultStrategyExchangeRateEmitted // Event containing the contract specifics and raw log

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
func (it *DurationVaultStrategyExchangeRateEmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DurationVaultStrategyExchangeRateEmitted)
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
		it.Event = new(DurationVaultStrategyExchangeRateEmitted)
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
func (it *DurationVaultStrategyExchangeRateEmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DurationVaultStrategyExchangeRateEmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DurationVaultStrategyExchangeRateEmitted represents a ExchangeRateEmitted event raised by the DurationVaultStrategy contract.
type DurationVaultStrategyExchangeRateEmitted struct {
	Rate *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterExchangeRateEmitted is a free log retrieval operation binding the contract event 0xd2494f3479e5da49d386657c292c610b5b01df313d07c62eb0cfa49924a31be8.
//
// Solidity: event ExchangeRateEmitted(uint256 rate)
func (_DurationVaultStrategy *DurationVaultStrategyFilterer) FilterExchangeRateEmitted(opts *bind.FilterOpts) (*DurationVaultStrategyExchangeRateEmittedIterator, error) {

	logs, sub, err := _DurationVaultStrategy.contract.FilterLogs(opts, "ExchangeRateEmitted")
	if err != nil {
		return nil, err
	}
	return &DurationVaultStrategyExchangeRateEmittedIterator{contract: _DurationVaultStrategy.contract, event: "ExchangeRateEmitted", logs: logs, sub: sub}, nil
}

// WatchExchangeRateEmitted is a free log subscription operation binding the contract event 0xd2494f3479e5da49d386657c292c610b5b01df313d07c62eb0cfa49924a31be8.
//
// Solidity: event ExchangeRateEmitted(uint256 rate)
func (_DurationVaultStrategy *DurationVaultStrategyFilterer) WatchExchangeRateEmitted(opts *bind.WatchOpts, sink chan<- *DurationVaultStrategyExchangeRateEmitted) (event.Subscription, error) {

	logs, sub, err := _DurationVaultStrategy.contract.WatchLogs(opts, "ExchangeRateEmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DurationVaultStrategyExchangeRateEmitted)
				if err := _DurationVaultStrategy.contract.UnpackLog(event, "ExchangeRateEmitted", log); err != nil {
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

// ParseExchangeRateEmitted is a log parse operation binding the contract event 0xd2494f3479e5da49d386657c292c610b5b01df313d07c62eb0cfa49924a31be8.
//
// Solidity: event ExchangeRateEmitted(uint256 rate)
func (_DurationVaultStrategy *DurationVaultStrategyFilterer) ParseExchangeRateEmitted(log types.Log) (*DurationVaultStrategyExchangeRateEmitted, error) {
	event := new(DurationVaultStrategyExchangeRateEmitted)
	if err := _DurationVaultStrategy.contract.UnpackLog(event, "ExchangeRateEmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DurationVaultStrategyInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the DurationVaultStrategy contract.
type DurationVaultStrategyInitializedIterator struct {
	Event *DurationVaultStrategyInitialized // Event containing the contract specifics and raw log

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
func (it *DurationVaultStrategyInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DurationVaultStrategyInitialized)
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
		it.Event = new(DurationVaultStrategyInitialized)
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
func (it *DurationVaultStrategyInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DurationVaultStrategyInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DurationVaultStrategyInitialized represents a Initialized event raised by the DurationVaultStrategy contract.
type DurationVaultStrategyInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_DurationVaultStrategy *DurationVaultStrategyFilterer) FilterInitialized(opts *bind.FilterOpts) (*DurationVaultStrategyInitializedIterator, error) {

	logs, sub, err := _DurationVaultStrategy.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &DurationVaultStrategyInitializedIterator{contract: _DurationVaultStrategy.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_DurationVaultStrategy *DurationVaultStrategyFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *DurationVaultStrategyInitialized) (event.Subscription, error) {

	logs, sub, err := _DurationVaultStrategy.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DurationVaultStrategyInitialized)
				if err := _DurationVaultStrategy.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_DurationVaultStrategy *DurationVaultStrategyFilterer) ParseInitialized(log types.Log) (*DurationVaultStrategyInitialized, error) {
	event := new(DurationVaultStrategyInitialized)
	if err := _DurationVaultStrategy.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DurationVaultStrategyMaxPerDepositUpdatedIterator is returned from FilterMaxPerDepositUpdated and is used to iterate over the raw logs and unpacked data for MaxPerDepositUpdated events raised by the DurationVaultStrategy contract.
type DurationVaultStrategyMaxPerDepositUpdatedIterator struct {
	Event *DurationVaultStrategyMaxPerDepositUpdated // Event containing the contract specifics and raw log

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
func (it *DurationVaultStrategyMaxPerDepositUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DurationVaultStrategyMaxPerDepositUpdated)
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
		it.Event = new(DurationVaultStrategyMaxPerDepositUpdated)
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
func (it *DurationVaultStrategyMaxPerDepositUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DurationVaultStrategyMaxPerDepositUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DurationVaultStrategyMaxPerDepositUpdated represents a MaxPerDepositUpdated event raised by the DurationVaultStrategy contract.
type DurationVaultStrategyMaxPerDepositUpdated struct {
	PreviousValue *big.Int
	NewValue      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMaxPerDepositUpdated is a free log retrieval operation binding the contract event 0xf97ed4e083acac67830025ecbc756d8fe847cdbdca4cee3fe1e128e98b54ecb5.
//
// Solidity: event MaxPerDepositUpdated(uint256 previousValue, uint256 newValue)
func (_DurationVaultStrategy *DurationVaultStrategyFilterer) FilterMaxPerDepositUpdated(opts *bind.FilterOpts) (*DurationVaultStrategyMaxPerDepositUpdatedIterator, error) {

	logs, sub, err := _DurationVaultStrategy.contract.FilterLogs(opts, "MaxPerDepositUpdated")
	if err != nil {
		return nil, err
	}
	return &DurationVaultStrategyMaxPerDepositUpdatedIterator{contract: _DurationVaultStrategy.contract, event: "MaxPerDepositUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxPerDepositUpdated is a free log subscription operation binding the contract event 0xf97ed4e083acac67830025ecbc756d8fe847cdbdca4cee3fe1e128e98b54ecb5.
//
// Solidity: event MaxPerDepositUpdated(uint256 previousValue, uint256 newValue)
func (_DurationVaultStrategy *DurationVaultStrategyFilterer) WatchMaxPerDepositUpdated(opts *bind.WatchOpts, sink chan<- *DurationVaultStrategyMaxPerDepositUpdated) (event.Subscription, error) {

	logs, sub, err := _DurationVaultStrategy.contract.WatchLogs(opts, "MaxPerDepositUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DurationVaultStrategyMaxPerDepositUpdated)
				if err := _DurationVaultStrategy.contract.UnpackLog(event, "MaxPerDepositUpdated", log); err != nil {
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

// ParseMaxPerDepositUpdated is a log parse operation binding the contract event 0xf97ed4e083acac67830025ecbc756d8fe847cdbdca4cee3fe1e128e98b54ecb5.
//
// Solidity: event MaxPerDepositUpdated(uint256 previousValue, uint256 newValue)
func (_DurationVaultStrategy *DurationVaultStrategyFilterer) ParseMaxPerDepositUpdated(log types.Log) (*DurationVaultStrategyMaxPerDepositUpdated, error) {
	event := new(DurationVaultStrategyMaxPerDepositUpdated)
	if err := _DurationVaultStrategy.contract.UnpackLog(event, "MaxPerDepositUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DurationVaultStrategyMaxTotalDepositsUpdatedIterator is returned from FilterMaxTotalDepositsUpdated and is used to iterate over the raw logs and unpacked data for MaxTotalDepositsUpdated events raised by the DurationVaultStrategy contract.
type DurationVaultStrategyMaxTotalDepositsUpdatedIterator struct {
	Event *DurationVaultStrategyMaxTotalDepositsUpdated // Event containing the contract specifics and raw log

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
func (it *DurationVaultStrategyMaxTotalDepositsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DurationVaultStrategyMaxTotalDepositsUpdated)
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
		it.Event = new(DurationVaultStrategyMaxTotalDepositsUpdated)
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
func (it *DurationVaultStrategyMaxTotalDepositsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DurationVaultStrategyMaxTotalDepositsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DurationVaultStrategyMaxTotalDepositsUpdated represents a MaxTotalDepositsUpdated event raised by the DurationVaultStrategy contract.
type DurationVaultStrategyMaxTotalDepositsUpdated struct {
	PreviousValue *big.Int
	NewValue      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMaxTotalDepositsUpdated is a free log retrieval operation binding the contract event 0x6ab181e0440bfbf4bacdf2e99674735ce6638005490688c5f994f5399353e452.
//
// Solidity: event MaxTotalDepositsUpdated(uint256 previousValue, uint256 newValue)
func (_DurationVaultStrategy *DurationVaultStrategyFilterer) FilterMaxTotalDepositsUpdated(opts *bind.FilterOpts) (*DurationVaultStrategyMaxTotalDepositsUpdatedIterator, error) {

	logs, sub, err := _DurationVaultStrategy.contract.FilterLogs(opts, "MaxTotalDepositsUpdated")
	if err != nil {
		return nil, err
	}
	return &DurationVaultStrategyMaxTotalDepositsUpdatedIterator{contract: _DurationVaultStrategy.contract, event: "MaxTotalDepositsUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxTotalDepositsUpdated is a free log subscription operation binding the contract event 0x6ab181e0440bfbf4bacdf2e99674735ce6638005490688c5f994f5399353e452.
//
// Solidity: event MaxTotalDepositsUpdated(uint256 previousValue, uint256 newValue)
func (_DurationVaultStrategy *DurationVaultStrategyFilterer) WatchMaxTotalDepositsUpdated(opts *bind.WatchOpts, sink chan<- *DurationVaultStrategyMaxTotalDepositsUpdated) (event.Subscription, error) {

	logs, sub, err := _DurationVaultStrategy.contract.WatchLogs(opts, "MaxTotalDepositsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DurationVaultStrategyMaxTotalDepositsUpdated)
				if err := _DurationVaultStrategy.contract.UnpackLog(event, "MaxTotalDepositsUpdated", log); err != nil {
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

// ParseMaxTotalDepositsUpdated is a log parse operation binding the contract event 0x6ab181e0440bfbf4bacdf2e99674735ce6638005490688c5f994f5399353e452.
//
// Solidity: event MaxTotalDepositsUpdated(uint256 previousValue, uint256 newValue)
func (_DurationVaultStrategy *DurationVaultStrategyFilterer) ParseMaxTotalDepositsUpdated(log types.Log) (*DurationVaultStrategyMaxTotalDepositsUpdated, error) {
	event := new(DurationVaultStrategyMaxTotalDepositsUpdated)
	if err := _DurationVaultStrategy.contract.UnpackLog(event, "MaxTotalDepositsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DurationVaultStrategyMetadataURIUpdatedIterator is returned from FilterMetadataURIUpdated and is used to iterate over the raw logs and unpacked data for MetadataURIUpdated events raised by the DurationVaultStrategy contract.
type DurationVaultStrategyMetadataURIUpdatedIterator struct {
	Event *DurationVaultStrategyMetadataURIUpdated // Event containing the contract specifics and raw log

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
func (it *DurationVaultStrategyMetadataURIUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DurationVaultStrategyMetadataURIUpdated)
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
		it.Event = new(DurationVaultStrategyMetadataURIUpdated)
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
func (it *DurationVaultStrategyMetadataURIUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DurationVaultStrategyMetadataURIUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DurationVaultStrategyMetadataURIUpdated represents a MetadataURIUpdated event raised by the DurationVaultStrategy contract.
type DurationVaultStrategyMetadataURIUpdated struct {
	NewMetadataURI string
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterMetadataURIUpdated is a free log retrieval operation binding the contract event 0xefafb90526da1636e1335eac0151301742fb755d986954c613b90e891778ba39.
//
// Solidity: event MetadataURIUpdated(string newMetadataURI)
func (_DurationVaultStrategy *DurationVaultStrategyFilterer) FilterMetadataURIUpdated(opts *bind.FilterOpts) (*DurationVaultStrategyMetadataURIUpdatedIterator, error) {

	logs, sub, err := _DurationVaultStrategy.contract.FilterLogs(opts, "MetadataURIUpdated")
	if err != nil {
		return nil, err
	}
	return &DurationVaultStrategyMetadataURIUpdatedIterator{contract: _DurationVaultStrategy.contract, event: "MetadataURIUpdated", logs: logs, sub: sub}, nil
}

// WatchMetadataURIUpdated is a free log subscription operation binding the contract event 0xefafb90526da1636e1335eac0151301742fb755d986954c613b90e891778ba39.
//
// Solidity: event MetadataURIUpdated(string newMetadataURI)
func (_DurationVaultStrategy *DurationVaultStrategyFilterer) WatchMetadataURIUpdated(opts *bind.WatchOpts, sink chan<- *DurationVaultStrategyMetadataURIUpdated) (event.Subscription, error) {

	logs, sub, err := _DurationVaultStrategy.contract.WatchLogs(opts, "MetadataURIUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DurationVaultStrategyMetadataURIUpdated)
				if err := _DurationVaultStrategy.contract.UnpackLog(event, "MetadataURIUpdated", log); err != nil {
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

// ParseMetadataURIUpdated is a log parse operation binding the contract event 0xefafb90526da1636e1335eac0151301742fb755d986954c613b90e891778ba39.
//
// Solidity: event MetadataURIUpdated(string newMetadataURI)
func (_DurationVaultStrategy *DurationVaultStrategyFilterer) ParseMetadataURIUpdated(log types.Log) (*DurationVaultStrategyMetadataURIUpdated, error) {
	event := new(DurationVaultStrategyMetadataURIUpdated)
	if err := _DurationVaultStrategy.contract.UnpackLog(event, "MetadataURIUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DurationVaultStrategyPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the DurationVaultStrategy contract.
type DurationVaultStrategyPausedIterator struct {
	Event *DurationVaultStrategyPaused // Event containing the contract specifics and raw log

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
func (it *DurationVaultStrategyPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DurationVaultStrategyPaused)
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
		it.Event = new(DurationVaultStrategyPaused)
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
func (it *DurationVaultStrategyPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DurationVaultStrategyPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DurationVaultStrategyPaused represents a Paused event raised by the DurationVaultStrategy contract.
type DurationVaultStrategyPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_DurationVaultStrategy *DurationVaultStrategyFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*DurationVaultStrategyPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _DurationVaultStrategy.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &DurationVaultStrategyPausedIterator{contract: _DurationVaultStrategy.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_DurationVaultStrategy *DurationVaultStrategyFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *DurationVaultStrategyPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _DurationVaultStrategy.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DurationVaultStrategyPaused)
				if err := _DurationVaultStrategy.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_DurationVaultStrategy *DurationVaultStrategyFilterer) ParsePaused(log types.Log) (*DurationVaultStrategyPaused, error) {
	event := new(DurationVaultStrategyPaused)
	if err := _DurationVaultStrategy.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DurationVaultStrategyStrategyTokenSetIterator is returned from FilterStrategyTokenSet and is used to iterate over the raw logs and unpacked data for StrategyTokenSet events raised by the DurationVaultStrategy contract.
type DurationVaultStrategyStrategyTokenSetIterator struct {
	Event *DurationVaultStrategyStrategyTokenSet // Event containing the contract specifics and raw log

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
func (it *DurationVaultStrategyStrategyTokenSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DurationVaultStrategyStrategyTokenSet)
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
		it.Event = new(DurationVaultStrategyStrategyTokenSet)
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
func (it *DurationVaultStrategyStrategyTokenSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DurationVaultStrategyStrategyTokenSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DurationVaultStrategyStrategyTokenSet represents a StrategyTokenSet event raised by the DurationVaultStrategy contract.
type DurationVaultStrategyStrategyTokenSet struct {
	Token    common.Address
	Decimals uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStrategyTokenSet is a free log retrieval operation binding the contract event 0x1c540707b00eb5427b6b774fc799d756516a54aee108b64b327acc55af557507.
//
// Solidity: event StrategyTokenSet(address token, uint8 decimals)
func (_DurationVaultStrategy *DurationVaultStrategyFilterer) FilterStrategyTokenSet(opts *bind.FilterOpts) (*DurationVaultStrategyStrategyTokenSetIterator, error) {

	logs, sub, err := _DurationVaultStrategy.contract.FilterLogs(opts, "StrategyTokenSet")
	if err != nil {
		return nil, err
	}
	return &DurationVaultStrategyStrategyTokenSetIterator{contract: _DurationVaultStrategy.contract, event: "StrategyTokenSet", logs: logs, sub: sub}, nil
}

// WatchStrategyTokenSet is a free log subscription operation binding the contract event 0x1c540707b00eb5427b6b774fc799d756516a54aee108b64b327acc55af557507.
//
// Solidity: event StrategyTokenSet(address token, uint8 decimals)
func (_DurationVaultStrategy *DurationVaultStrategyFilterer) WatchStrategyTokenSet(opts *bind.WatchOpts, sink chan<- *DurationVaultStrategyStrategyTokenSet) (event.Subscription, error) {

	logs, sub, err := _DurationVaultStrategy.contract.WatchLogs(opts, "StrategyTokenSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DurationVaultStrategyStrategyTokenSet)
				if err := _DurationVaultStrategy.contract.UnpackLog(event, "StrategyTokenSet", log); err != nil {
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

// ParseStrategyTokenSet is a log parse operation binding the contract event 0x1c540707b00eb5427b6b774fc799d756516a54aee108b64b327acc55af557507.
//
// Solidity: event StrategyTokenSet(address token, uint8 decimals)
func (_DurationVaultStrategy *DurationVaultStrategyFilterer) ParseStrategyTokenSet(log types.Log) (*DurationVaultStrategyStrategyTokenSet, error) {
	event := new(DurationVaultStrategyStrategyTokenSet)
	if err := _DurationVaultStrategy.contract.UnpackLog(event, "StrategyTokenSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DurationVaultStrategyUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the DurationVaultStrategy contract.
type DurationVaultStrategyUnpausedIterator struct {
	Event *DurationVaultStrategyUnpaused // Event containing the contract specifics and raw log

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
func (it *DurationVaultStrategyUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DurationVaultStrategyUnpaused)
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
		it.Event = new(DurationVaultStrategyUnpaused)
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
func (it *DurationVaultStrategyUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DurationVaultStrategyUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DurationVaultStrategyUnpaused represents a Unpaused event raised by the DurationVaultStrategy contract.
type DurationVaultStrategyUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_DurationVaultStrategy *DurationVaultStrategyFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*DurationVaultStrategyUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _DurationVaultStrategy.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &DurationVaultStrategyUnpausedIterator{contract: _DurationVaultStrategy.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_DurationVaultStrategy *DurationVaultStrategyFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *DurationVaultStrategyUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _DurationVaultStrategy.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DurationVaultStrategyUnpaused)
				if err := _DurationVaultStrategy.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_DurationVaultStrategy *DurationVaultStrategyFilterer) ParseUnpaused(log types.Log) (*DurationVaultStrategyUnpaused, error) {
	event := new(DurationVaultStrategyUnpaused)
	if err := _DurationVaultStrategy.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DurationVaultStrategyVaultAdvancedToWithdrawalsIterator is returned from FilterVaultAdvancedToWithdrawals and is used to iterate over the raw logs and unpacked data for VaultAdvancedToWithdrawals events raised by the DurationVaultStrategy contract.
type DurationVaultStrategyVaultAdvancedToWithdrawalsIterator struct {
	Event *DurationVaultStrategyVaultAdvancedToWithdrawals // Event containing the contract specifics and raw log

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
func (it *DurationVaultStrategyVaultAdvancedToWithdrawalsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DurationVaultStrategyVaultAdvancedToWithdrawals)
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
		it.Event = new(DurationVaultStrategyVaultAdvancedToWithdrawals)
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
func (it *DurationVaultStrategyVaultAdvancedToWithdrawalsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DurationVaultStrategyVaultAdvancedToWithdrawalsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DurationVaultStrategyVaultAdvancedToWithdrawals represents a VaultAdvancedToWithdrawals event raised by the DurationVaultStrategy contract.
type DurationVaultStrategyVaultAdvancedToWithdrawals struct {
	Arbitrator common.Address
	MaturedAt  uint32
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVaultAdvancedToWithdrawals is a free log retrieval operation binding the contract event 0x96c49d03ef64591194500229a104cd087b2d45c68234c96444c3a2a6abb0bb97.
//
// Solidity: event VaultAdvancedToWithdrawals(address indexed arbitrator, uint32 maturedAt)
func (_DurationVaultStrategy *DurationVaultStrategyFilterer) FilterVaultAdvancedToWithdrawals(opts *bind.FilterOpts, arbitrator []common.Address) (*DurationVaultStrategyVaultAdvancedToWithdrawalsIterator, error) {

	var arbitratorRule []interface{}
	for _, arbitratorItem := range arbitrator {
		arbitratorRule = append(arbitratorRule, arbitratorItem)
	}

	logs, sub, err := _DurationVaultStrategy.contract.FilterLogs(opts, "VaultAdvancedToWithdrawals", arbitratorRule)
	if err != nil {
		return nil, err
	}
	return &DurationVaultStrategyVaultAdvancedToWithdrawalsIterator{contract: _DurationVaultStrategy.contract, event: "VaultAdvancedToWithdrawals", logs: logs, sub: sub}, nil
}

// WatchVaultAdvancedToWithdrawals is a free log subscription operation binding the contract event 0x96c49d03ef64591194500229a104cd087b2d45c68234c96444c3a2a6abb0bb97.
//
// Solidity: event VaultAdvancedToWithdrawals(address indexed arbitrator, uint32 maturedAt)
func (_DurationVaultStrategy *DurationVaultStrategyFilterer) WatchVaultAdvancedToWithdrawals(opts *bind.WatchOpts, sink chan<- *DurationVaultStrategyVaultAdvancedToWithdrawals, arbitrator []common.Address) (event.Subscription, error) {

	var arbitratorRule []interface{}
	for _, arbitratorItem := range arbitrator {
		arbitratorRule = append(arbitratorRule, arbitratorItem)
	}

	logs, sub, err := _DurationVaultStrategy.contract.WatchLogs(opts, "VaultAdvancedToWithdrawals", arbitratorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DurationVaultStrategyVaultAdvancedToWithdrawals)
				if err := _DurationVaultStrategy.contract.UnpackLog(event, "VaultAdvancedToWithdrawals", log); err != nil {
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

// ParseVaultAdvancedToWithdrawals is a log parse operation binding the contract event 0x96c49d03ef64591194500229a104cd087b2d45c68234c96444c3a2a6abb0bb97.
//
// Solidity: event VaultAdvancedToWithdrawals(address indexed arbitrator, uint32 maturedAt)
func (_DurationVaultStrategy *DurationVaultStrategyFilterer) ParseVaultAdvancedToWithdrawals(log types.Log) (*DurationVaultStrategyVaultAdvancedToWithdrawals, error) {
	event := new(DurationVaultStrategyVaultAdvancedToWithdrawals)
	if err := _DurationVaultStrategy.contract.UnpackLog(event, "VaultAdvancedToWithdrawals", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DurationVaultStrategyVaultInitializedIterator is returned from FilterVaultInitialized and is used to iterate over the raw logs and unpacked data for VaultInitialized events raised by the DurationVaultStrategy contract.
type DurationVaultStrategyVaultInitializedIterator struct {
	Event *DurationVaultStrategyVaultInitialized // Event containing the contract specifics and raw log

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
func (it *DurationVaultStrategyVaultInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DurationVaultStrategyVaultInitialized)
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
		it.Event = new(DurationVaultStrategyVaultInitialized)
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
func (it *DurationVaultStrategyVaultInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DurationVaultStrategyVaultInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DurationVaultStrategyVaultInitialized represents a VaultInitialized event raised by the DurationVaultStrategy contract.
type DurationVaultStrategyVaultInitialized struct {
	VaultAdmin      common.Address
	Arbitrator      common.Address
	UnderlyingToken common.Address
	Duration        uint32
	MaxPerDeposit   *big.Int
	StakeCap        *big.Int
	MetadataURI     string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterVaultInitialized is a free log retrieval operation binding the contract event 0xbdbff63632f473bb2a7c6a4aafbc096b71fbda12e22c6b51643bfd64f13d2b9e.
//
// Solidity: event VaultInitialized(address indexed vaultAdmin, address indexed arbitrator, address indexed underlyingToken, uint32 duration, uint256 maxPerDeposit, uint256 stakeCap, string metadataURI)
func (_DurationVaultStrategy *DurationVaultStrategyFilterer) FilterVaultInitialized(opts *bind.FilterOpts, vaultAdmin []common.Address, arbitrator []common.Address, underlyingToken []common.Address) (*DurationVaultStrategyVaultInitializedIterator, error) {

	var vaultAdminRule []interface{}
	for _, vaultAdminItem := range vaultAdmin {
		vaultAdminRule = append(vaultAdminRule, vaultAdminItem)
	}
	var arbitratorRule []interface{}
	for _, arbitratorItem := range arbitrator {
		arbitratorRule = append(arbitratorRule, arbitratorItem)
	}
	var underlyingTokenRule []interface{}
	for _, underlyingTokenItem := range underlyingToken {
		underlyingTokenRule = append(underlyingTokenRule, underlyingTokenItem)
	}

	logs, sub, err := _DurationVaultStrategy.contract.FilterLogs(opts, "VaultInitialized", vaultAdminRule, arbitratorRule, underlyingTokenRule)
	if err != nil {
		return nil, err
	}
	return &DurationVaultStrategyVaultInitializedIterator{contract: _DurationVaultStrategy.contract, event: "VaultInitialized", logs: logs, sub: sub}, nil
}

// WatchVaultInitialized is a free log subscription operation binding the contract event 0xbdbff63632f473bb2a7c6a4aafbc096b71fbda12e22c6b51643bfd64f13d2b9e.
//
// Solidity: event VaultInitialized(address indexed vaultAdmin, address indexed arbitrator, address indexed underlyingToken, uint32 duration, uint256 maxPerDeposit, uint256 stakeCap, string metadataURI)
func (_DurationVaultStrategy *DurationVaultStrategyFilterer) WatchVaultInitialized(opts *bind.WatchOpts, sink chan<- *DurationVaultStrategyVaultInitialized, vaultAdmin []common.Address, arbitrator []common.Address, underlyingToken []common.Address) (event.Subscription, error) {

	var vaultAdminRule []interface{}
	for _, vaultAdminItem := range vaultAdmin {
		vaultAdminRule = append(vaultAdminRule, vaultAdminItem)
	}
	var arbitratorRule []interface{}
	for _, arbitratorItem := range arbitrator {
		arbitratorRule = append(arbitratorRule, arbitratorItem)
	}
	var underlyingTokenRule []interface{}
	for _, underlyingTokenItem := range underlyingToken {
		underlyingTokenRule = append(underlyingTokenRule, underlyingTokenItem)
	}

	logs, sub, err := _DurationVaultStrategy.contract.WatchLogs(opts, "VaultInitialized", vaultAdminRule, arbitratorRule, underlyingTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DurationVaultStrategyVaultInitialized)
				if err := _DurationVaultStrategy.contract.UnpackLog(event, "VaultInitialized", log); err != nil {
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

// ParseVaultInitialized is a log parse operation binding the contract event 0xbdbff63632f473bb2a7c6a4aafbc096b71fbda12e22c6b51643bfd64f13d2b9e.
//
// Solidity: event VaultInitialized(address indexed vaultAdmin, address indexed arbitrator, address indexed underlyingToken, uint32 duration, uint256 maxPerDeposit, uint256 stakeCap, string metadataURI)
func (_DurationVaultStrategy *DurationVaultStrategyFilterer) ParseVaultInitialized(log types.Log) (*DurationVaultStrategyVaultInitialized, error) {
	event := new(DurationVaultStrategyVaultInitialized)
	if err := _DurationVaultStrategy.contract.UnpackLog(event, "VaultInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DurationVaultStrategyVaultLockedIterator is returned from FilterVaultLocked and is used to iterate over the raw logs and unpacked data for VaultLocked events raised by the DurationVaultStrategy contract.
type DurationVaultStrategyVaultLockedIterator struct {
	Event *DurationVaultStrategyVaultLocked // Event containing the contract specifics and raw log

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
func (it *DurationVaultStrategyVaultLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DurationVaultStrategyVaultLocked)
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
		it.Event = new(DurationVaultStrategyVaultLocked)
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
func (it *DurationVaultStrategyVaultLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DurationVaultStrategyVaultLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DurationVaultStrategyVaultLocked represents a VaultLocked event raised by the DurationVaultStrategy contract.
type DurationVaultStrategyVaultLocked struct {
	LockedAt uint32
	UnlockAt uint32
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterVaultLocked is a free log retrieval operation binding the contract event 0x42cd6d7338516695d9c9ff8969dbdcf89ce22e3f2f76fda2fc11e973fe4860e4.
//
// Solidity: event VaultLocked(uint32 lockedAt, uint32 unlockAt)
func (_DurationVaultStrategy *DurationVaultStrategyFilterer) FilterVaultLocked(opts *bind.FilterOpts) (*DurationVaultStrategyVaultLockedIterator, error) {

	logs, sub, err := _DurationVaultStrategy.contract.FilterLogs(opts, "VaultLocked")
	if err != nil {
		return nil, err
	}
	return &DurationVaultStrategyVaultLockedIterator{contract: _DurationVaultStrategy.contract, event: "VaultLocked", logs: logs, sub: sub}, nil
}

// WatchVaultLocked is a free log subscription operation binding the contract event 0x42cd6d7338516695d9c9ff8969dbdcf89ce22e3f2f76fda2fc11e973fe4860e4.
//
// Solidity: event VaultLocked(uint32 lockedAt, uint32 unlockAt)
func (_DurationVaultStrategy *DurationVaultStrategyFilterer) WatchVaultLocked(opts *bind.WatchOpts, sink chan<- *DurationVaultStrategyVaultLocked) (event.Subscription, error) {

	logs, sub, err := _DurationVaultStrategy.contract.WatchLogs(opts, "VaultLocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DurationVaultStrategyVaultLocked)
				if err := _DurationVaultStrategy.contract.UnpackLog(event, "VaultLocked", log); err != nil {
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

// ParseVaultLocked is a log parse operation binding the contract event 0x42cd6d7338516695d9c9ff8969dbdcf89ce22e3f2f76fda2fc11e973fe4860e4.
//
// Solidity: event VaultLocked(uint32 lockedAt, uint32 unlockAt)
func (_DurationVaultStrategy *DurationVaultStrategyFilterer) ParseVaultLocked(log types.Log) (*DurationVaultStrategyVaultLocked, error) {
	event := new(DurationVaultStrategyVaultLocked)
	if err := _DurationVaultStrategy.contract.UnpackLog(event, "VaultLocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DurationVaultStrategyVaultMaturedIterator is returned from FilterVaultMatured and is used to iterate over the raw logs and unpacked data for VaultMatured events raised by the DurationVaultStrategy contract.
type DurationVaultStrategyVaultMaturedIterator struct {
	Event *DurationVaultStrategyVaultMatured // Event containing the contract specifics and raw log

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
func (it *DurationVaultStrategyVaultMaturedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DurationVaultStrategyVaultMatured)
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
		it.Event = new(DurationVaultStrategyVaultMatured)
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
func (it *DurationVaultStrategyVaultMaturedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DurationVaultStrategyVaultMaturedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DurationVaultStrategyVaultMatured represents a VaultMatured event raised by the DurationVaultStrategy contract.
type DurationVaultStrategyVaultMatured struct {
	MaturedAt uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterVaultMatured is a free log retrieval operation binding the contract event 0xff979382d3040b1602e0a02f0f2a454b2250aa36e891d2da0ceb95d70d11a8f2.
//
// Solidity: event VaultMatured(uint32 maturedAt)
func (_DurationVaultStrategy *DurationVaultStrategyFilterer) FilterVaultMatured(opts *bind.FilterOpts) (*DurationVaultStrategyVaultMaturedIterator, error) {

	logs, sub, err := _DurationVaultStrategy.contract.FilterLogs(opts, "VaultMatured")
	if err != nil {
		return nil, err
	}
	return &DurationVaultStrategyVaultMaturedIterator{contract: _DurationVaultStrategy.contract, event: "VaultMatured", logs: logs, sub: sub}, nil
}

// WatchVaultMatured is a free log subscription operation binding the contract event 0xff979382d3040b1602e0a02f0f2a454b2250aa36e891d2da0ceb95d70d11a8f2.
//
// Solidity: event VaultMatured(uint32 maturedAt)
func (_DurationVaultStrategy *DurationVaultStrategyFilterer) WatchVaultMatured(opts *bind.WatchOpts, sink chan<- *DurationVaultStrategyVaultMatured) (event.Subscription, error) {

	logs, sub, err := _DurationVaultStrategy.contract.WatchLogs(opts, "VaultMatured")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DurationVaultStrategyVaultMatured)
				if err := _DurationVaultStrategy.contract.UnpackLog(event, "VaultMatured", log); err != nil {
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

// ParseVaultMatured is a log parse operation binding the contract event 0xff979382d3040b1602e0a02f0f2a454b2250aa36e891d2da0ceb95d70d11a8f2.
//
// Solidity: event VaultMatured(uint32 maturedAt)
func (_DurationVaultStrategy *DurationVaultStrategyFilterer) ParseVaultMatured(log types.Log) (*DurationVaultStrategyVaultMatured, error) {
	event := new(DurationVaultStrategyVaultMatured)
	if err := _DurationVaultStrategy.contract.UnpackLog(event, "VaultMatured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
