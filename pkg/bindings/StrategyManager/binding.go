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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_allocationManager\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"},{\"name\":\"_delegation\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"_version\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_BURN_ADDRESS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEPOSIT_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SLASH_RESOLUTION_DELAY_BLOCKS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addStrategiesToDepositWhitelist\",\"inputs\":[{\"name\":\"strategiesToWhitelist\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allocationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burnShares\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"calculateStrategyDepositDigestHash\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"clearBurnOrRedistributableShares\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"clearBurnOrRedistributableSharesByStrategy\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"depositIntoStrategy\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"depositShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositIntoStrategyWithSignature\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"depositShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"domainSeparator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBurnOrRedistributableCount\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBurnOrRedistributableShares\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBurnOrRedistributableShares\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBurnableShares\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDeposits\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPendingOperatorSets\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPendingSlashIds\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSlashResolutionBlock\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStakerStrategyList\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStrategiesWithBurnableShares\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"increaseBurnOrRedistributableShares\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"sharesToBurn\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialStrategyWhitelister\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"nonces\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeDepositShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"depositSharesToRemove\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeStrategiesFromDepositWhitelist\",\"inputs\":[{\"name\":\"strategiesToRemoveFromWhitelist\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStrategyWhitelister\",\"inputs\":[{\"name\":\"newStrategyWhitelister\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakerDepositShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakerStrategyList\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"strategies\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakerStrategyListLength\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"strategyIsWhitelistedForDeposit\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"whitelisted\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"strategyWhitelister\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawSharesAsTokens\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"BurnOrRedistributableSharesDecreased\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BurnOrRedistributableSharesIncreased\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BurnableSharesDecreased\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposit\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyAddedToDepositWhitelist\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyRemovedFromDepositWhitelist\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyWhitelisterChanged\",\"inputs\":[{\"name\":\"previousAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CurrentlyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNewPausedStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MaxStrategiesExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyDelegationManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyPauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyStrategyWhitelister\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnpauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SharesAmountTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SharesAmountZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignatureExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SlashResolutionDelayNotElapsed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StakerAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategyAlreadyInSlash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategyNotFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategyNotWhitelisted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]}]",
	Bin: "0x610100604052348015610010575f5ffd5b50604051613c8f380380613c8f83398101604081905261002f916101bc565b80808585856001600160a01b03811661005b576040516339b190bb60e11b815260040160405180910390fd5b6001600160a01b0390811660805291821660a0521660c05261007c81610093565b60e0525061008a90506100d9565b50505050610301565b5f5f829050601f815111156100c6578260405163305a27a960e01b81526004016100bd91906102a6565b60405180910390fd5b80516100d1826102db565b179392505050565b5f54610100900460ff16156101405760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b60648201526084016100bd565b5f5460ff9081161461018f575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146101a5575f5ffd5b50565b634e487b7160e01b5f52604160045260245ffd5b5f5f5f5f608085870312156101cf575f5ffd5b84516101da81610191565b60208601519094506101eb81610191565b60408601519093506101fc81610191565b60608601519092506001600160401b03811115610217575f5ffd5b8501601f81018713610227575f5ffd5b80516001600160401b03811115610240576102406101a8565b604051601f8201601f19908116603f011681016001600160401b038111828210171561026e5761026e6101a8565b604052818152828201602001891015610285575f5ffd5b8160208401602083015e5f6020838301015280935050505092959194509250565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b805160208083015191908110156102fb575f198160200360031b1b821691505b50919050565b60805160a05160c05160e05161391561037a5f395f81816111f2015261271001525f81816106410152818161096f01528181610e4c015281816111840152818161139f015261228a01525f81816105e7015281816108c201526115d201525f81816104fd01528181611c4e015261277f01526139155ff3fe608060405234801561000f575f5ffd5b5060043610610281575f3560e01c80637def156411610156578063c6656702116100ca578063f2fde38b11610084578063f2fde38b14610676578063f3b4a00014610689578063f698da2514610693578063fabc1cbc1461069b578063fd980423146106ae578063fe243a17146106c1575f5ffd5b8063c6656702146105cf578063ca8aa7c7146105e2578063cbc2bd6214610609578063de44acb61461061c578063df5cf7231461063c578063e7a050aa14610663575f5ffd5b80638b8aac3c1161011b5780638b8aac3c1461054a5780638da5cb5b1461057257806394f649dd14610583578063967fc0d2146105965780639ac01d61146105a9578063b5d8b5b8146105bc575f5ffd5b80637def1564146104a65780637ecebe00146104c6578063829fca73146104e5578063886f1195146104f857806388c1029914610537575f5ffd5b80634b3262f7116101f85780635c975abb116101b25780635c975abb1461043b5780635de08ff214610443578063663c1de414610456578063715018a614610478578063724af4231461048057806376fb162b14610493575f5ffd5b80634b3262f71461039d5780634b6d5d6e146103b057806350ff7225146103c357806354fd4d50146103eb578063595c6a67146104005780635ac86ab714610408575f5ffd5b806331f8fb4c1161024957806331f8fb4c1461030457806332e89ace1461032557806336a8c500146103385780633f292b081461034e5780633fb99ca51461036357806348825e9414610376575f5ffd5b806307a7935514610285578063136439dd146102a85780631794bb3c146102bd5780632d44def6146102d05780632eae418c146102f1575b5f5ffd5b61028e61c4e081565b60405163ffffffff90911681526020015b60405180910390f35b6102bb6102b63660046130cf565b6106eb565b005b6102bb6102cb3660046130fa565b610725565b6102e36102de36600461314e565b61084b565b60405190815260200161029f565b6102bb6102ff36600461318c565b610964565b6103176103123660046131da565b610a30565b60405161029f929190613276565b6102e36103333660046132e7565b610bb9565b610340610c3e565b60405161029f9291906133c1565b610356610d59565b60405161029f9190613417565b6102bb610371366004613474565b610e41565b6102e37f4337f82d142e41f2a8c10547cd8c859bddb92262a61058e77842e24d9dea922481565b61028e6103ab3660046131da565b610fe8565b6102bb6103be3660046134b8565b61102a565b6103d66103d13660046130fa565b611177565b6040805192835260208301919091520161029f565b6103f36111eb565b60405161029f9190613501565b6102bb61121b565b61042b610416366004613513565b609854600160ff9092169190911b9081161490565b604051901515815260200161029f565b6098546102e3565b6102bb610451366004613533565b61122f565b61042b6104643660046134b8565b60d16020525f908152604090205460ff1681565b6102bb611382565b6102e361048e3660046130fa565b611393565b6102e36104a136600461314e565b611403565b6104b96104b43660046135a2565b611452565b60405161029f91906135bc565b6102e36104d43660046134b8565b60ca6020525f908152604090205481565b6102e36104f33660046131da565b61147e565b61051f7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b03909116815260200161029f565b6104b96105453660046131da565b6114b8565b6102e36105583660046134b8565b6001600160a01b03165f90815260ce602052604090205490565b6033546001600160a01b031661051f565b6103176105913660046134b8565b611650565b60cb5461051f906001600160a01b031681565b6102e36105b73660046135ce565b6117c7565b6102bb6105ca366004613533565b611858565b6102bb6105dd3660046134b8565b61199f565b61051f7f000000000000000000000000000000000000000000000000000000000000000081565b61051f61061736600461362f565b6119c2565b61062f61062a3660046134b8565b6119f6565b60405161029f9190613659565b61051f7f000000000000000000000000000000000000000000000000000000000000000081565b6102e36106713660046130fa565b611a69565b6102bb6106843660046134b8565b611a88565b61051f620e16e481565b6102e3611afe565b6102bb6106a93660046130cf565b611bb7565b6102e36106bc3660046134b8565b611c24565b6102e36106cf36600461366b565b60cd60209081525f928352604080842090915290825290205481565b6106f3611c39565b60985481811681146107185760405163c61dca5d60e01b815260040160405180910390fd5b61072182611cdc565b5050565b5f54610100900460ff161580801561074357505f54600160ff909116105b8061075c5750303b15801561075c57505f5460ff166001145b6107c45760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b5f805460ff1916600117905580156107e5575f805461ff0019166101001790555b6107ee82611cdc565b6107f784611d19565b61080083611d6a565b8015610845575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050565b5f600161085781611dd3565b61085f611dfe565b60db5f610879610874368990038901896136ba565b611e57565b815260208082019290925260409081015f90812087825290925290205463ffffffff1643116108ba57604051626f6f0760e41b815260040160405180910390fd5b6109508585857f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316630f3df50e8a6040518263ffffffff1660e01b815260040161090c9190613747565b602060405180830381865afa158015610927573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061094b9190613755565b611eba565b915061095c6001606555565b509392505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146109ad5760405163f739589b60e01b815260040160405180910390fd5b6109b5611dfe565b604051636ce5768960e11b81526001600160a01b0384169063d9caed12906109e590879086908690600401613770565b6020604051808303815f875af1158015610a01573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a259190613794565b506108456001606555565b6060805f60d781610a49610874368990038901896136ba565b81526020019081526020015f205f8581526020019081526020015f2090505f610a71826120ff565b90505f81516001600160401b03811115610a8d57610a8d6132a3565b604051908082528060200260200182016040528015610ab6578160200160208202803683370190505b5090505f82516001600160401b03811115610ad357610ad36132a3565b604051908082528060200260200182016040528015610afc578160200160208202803683370190505b5090505f5b8351811015610ba957838181518110610b1c57610b1c6137ab565b6020026020010151838281518110610b3657610b366137ab565b60200260200101906001600160a01b031690816001600160a01b031681525050610b82848281518110610b6b57610b6b6137ab565b60200260200101518661210b90919063ffffffff16565b9050828281518110610b9657610b966137ab565b6020908102919091010152600101610b01565b50909450925050505b9250929050565b5f5f610bc481611dd3565b610bcc611dfe565b6001600160a01b0385165f90815260ca6020526040902054610bfd86610bf6818c8c8c878c6117c7565b868861212f565b6001600160a01b0386165f90815260ca60205260409020600182019055610c26868a8a8a612181565b925050610c336001606555565b509695505050505050565b6060805f610c4c60d46122ee565b90505f816001600160401b03811115610c6757610c676132a3565b604051908082528060200260200182016040528015610c90578160200160208202803683370190505b5090505f826001600160401b03811115610cac57610cac6132a3565b604051908082528060200260200182016040528015610cd5578160200160208202803683370190505b5090505f5b83811015610d4e575f5f610cef60d4846122f8565b9150915081858481518110610d0657610d066137ab565b60200260200101906001600160a01b031690816001600160a01b03168152505080848481518110610d3957610d396137ab565b60209081029190910101525050600101610cda565b509094909350915050565b60605f610d6660d8612306565b90505f816001600160401b03811115610d8157610d816132a3565b604051908082528060200260200182016040528015610dc557816020015b604080518082019091525f8082526020820152815260200190600190039081610d9f5790505b5090505f5b82811015610e3a57610e15610de060d88361230f565b604080518082019091525f80825260208201525060408051808201909152606082901c815263ffffffff909116602082015290565b828281518110610e2757610e276137ab565b6020908102919091010152600101610dca565b5092915050565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610e8a5760405163f739589b60e01b815260040160405180910390fd5b610e92611dfe565b5f60d781610ea8610874368990038901896136ba565b815260208082019290925260409081015f90812087825290925290209050610ed181848461231a565b610eee5760405163ca354fa360e01b815260040160405180910390fd5b610f0b610f03610874368890038801886136ba565b60d89061232f565b50610f408460da5f610f25610874368b90038b018b6136ba565b81526020019081526020015f2061232f90919063ffffffff16565b15610fa057610f5161c4e0436137d3565b60db5f610f66610874368a90038a018a6136ba565b81526020019081526020015f205f8681526020019081526020015f205f6101000a81548163ffffffff021916908363ffffffff1602179055505b7f5f5209798bbac45a16d2dc3bc67319fab26ee00153916d6f07b69f8a134a1e8b85858585604051610fd594939291906137ef565b60405180910390a1506108456001606555565b5f60db81610ffe610874368790038701876136ba565b815260208082019290925260409081015f90812085825290925290205463ffffffff1690505b92915050565b611032611dfe565b5f61103e60d48361210b565b91505061104c60d48361233a565b50604080516001600160a01b0384168152602081018390527fd9d082c3ec4f3a3ffa55c324939a06407f5fbcb87d5e0ce3b9508c92c84ed839910160405180910390a1801561116957816001600160a01b031663d9caed12620e16e4846001600160a01b0316632495a5996040518163ffffffff1660e01b8152600401602060405180830381865afa1580156110e4573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906111089190613755565b846040518463ffffffff1660e01b815260040161112793929190613770565b6020604051808303815f875af1158015611143573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906111679190613794565b505b506111746001606555565b50565b5f80336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146111c25760405163f739589b60e01b815260040160405180910390fd5b6111ca611dfe565b6111d585858561234e565b915091506111e36001606555565b935093915050565b60606112167f0000000000000000000000000000000000000000000000000000000000000000612512565b905090565b611223611c39565b61122d5f19611cdc565b565b60cb546001600160a01b0316331461125a576040516320ba3ff960e21b815260040160405180910390fd5b611262611dfe565b805f5b818110156113765760d15f858584818110611282576112826137ab565b905060200201602081019061129791906134b8565b6001600160a01b0316815260208101919091526040015f205460ff1661136e57600160d15f8686858181106112ce576112ce6137ab565b90506020020160208101906112e391906134b8565b6001600160a01b0316815260208101919091526040015f20805460ff19169115159190911790557f0c35b17d91c96eb2751cd456e1252f42a386e524ef9ff26ecc9950859fdc04fe84848381811061133d5761133d6137ab565b905060200201602081019061135291906134b8565b6040516001600160a01b03909116815260200160405180910390a15b600101611265565b50506107216001606555565b61138a61254f565b61122d5f611d19565b5f336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146113dd5760405163f739589b60e01b815260040160405180910390fd5b6113e5611dfe565b6113f08484846125a9565b90506113fc6001606555565b9392505050565b5f806114488360d78361141e610874368b90038b018b6136ba565b81526020019081526020015f205f8781526020019081526020015f2061210b90919063ffffffff16565b9695505050505050565b606061102460da5f61146c610874368790038701876136ba565b81526020019081526020015f206126b6565b5f6113fc60d782611497610874368890038801886136ba565b81526020019081526020015f205f8481526020019081526020015f206122ee565b606060016114c581611dd3565b6114cd611dfe565b60db5f6114e2610874368890038801886136ba565b815260208082019290925260409081015f90812086825290925290205463ffffffff16431161152357604051626f6f0760e41b815260040160405180910390fd5b5f61155d60d78261153c610874368a90038a018a6136ba565b81526020019081526020015f205f8681526020019081526020015f206120ff565b80519091505f816001600160401b0381111561157b5761157b6132a3565b6040519080825280602002602001820160405280156115a4578160200160208202803683370190505b5090505f5b828110156116415761161c88888684815181106115c8576115c86137ab565b60200260200101517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316630f3df50e8d6040518263ffffffff1660e01b815260040161090c9190613747565b82828151811061162e5761162e6137ab565b60209081029190910101526001016115a9565b5093505050610e3a6001606555565b6001600160a01b0381165f90815260ce6020526040812054606091829190816001600160401b03811115611686576116866132a3565b6040519080825280602002602001820160405280156116af578160200160208202803683370190505b5090505f5b8281101561173d576001600160a01b0386165f90815260cd6020908152604080832060ce90925282208054919291849081106116f2576116f26137ab565b5f9182526020808320909101546001600160a01b03168352820192909252604001902054825183908390811061172a5761172a6137ab565b60209081029190910101526001016116b4565b5060ce5f866001600160a01b03166001600160a01b031681526020019081526020015f2081818054806020026020016040519081016040528092919081815260200182805480156117b557602002820191905f5260205f20905b81546001600160a01b03168152600190910190602001808311611797575b50505050509150935093505050915091565b604080517f4337f82d142e41f2a8c10547cd8c859bddb92262a61058e77842e24d9dea922460208201526001600160a01b03808916928201929092528187166060820152908516608082015260a0810184905260c0810183905260e081018290525f9061184d9061010001604051602081830303815290604052805190602001206126c2565b979650505050505050565b60cb546001600160a01b03163314611883576040516320ba3ff960e21b815260040160405180910390fd5b61188b611dfe565b805f5b818110156113765760d15f8585848181106118ab576118ab6137ab565b90506020020160208101906118c091906134b8565b6001600160a01b0316815260208101919091526040015f205460ff1615611997575f60d15f8686858181106118f7576118f76137ab565b905060200201602081019061190c91906134b8565b6001600160a01b0316815260208101919091526040015f20805460ff19169115159190911790557f4074413b4b443e4e58019f2855a8765113358c7c72e39509c6af45fc0f5ba030848483818110611966576119666137ab565b905060200201602081019061197b91906134b8565b6040516001600160a01b03909116815260200160405180910390a15b60010161188e565b6119a761254f565b6119af611dfe565b6119b881611d6a565b6111746001606555565b60ce602052815f5260405f2081815481106119db575f80fd5b5f918252602090912001546001600160a01b03169150829050565b6001600160a01b0381165f90815260ce6020908152604091829020805483518184028101840190945280845260609392830182828015611a5d57602002820191905f5260205f20905b81546001600160a01b03168152600190910190602001808311611a3f575b50505050509050919050565b5f5f611a7481611dd3565b611a7c611dfe565b61095033868686612181565b611a9061254f565b6001600160a01b038116611af55760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016107bb565b61117481611d19565b60408051808201909152600a81526922b4b3b2b72630bcb2b960b11b6020909101525f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f7f71b625cfad44bac63b13dba07f2e1d6084ee04b6f8752101ece6126d584ee6ea611b6b612708565b805160209182012060408051928301949094529281019190915260608101919091524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b611bbf61277d565b60985480198219811614611be65760405163c61dca5d60e01b815260040160405180910390fd5b609882905560405182815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c9060200160405180910390a25050565b5f5f611c3160d48461210b565b949350505050565b60405163237dfb4760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906346fbf68e90602401602060405180830381865afa158015611c9b573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611cbf9190613821565b61122d57604051631d77d47760e21b815260040160405180910390fd5b609881905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a250565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b60cb54604080516001600160a01b03928316815291831660208301527f4264275e593955ff9d6146a51a4525f6ddace2e81db9391abcc9d1ca48047d29910160405180910390a160cb80546001600160a01b0319166001600160a01b0392909216919091179055565b609854600160ff83161b908116036111745760405163840a48d560e01b815260040160405180910390fd5b600260655403611e505760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064016107bb565b6002606555565b5f815f0151826020015163ffffffff16604051602001611ea292919060609290921b6bffffffffffffffffffffffff1916825260a01b6001600160a01b031916601482015260200190565b60405160208183030381529060405261102490613840565b5f8060d781611ed1610874368a90038a018a6136ba565b815260208082019290925260409081015f90812088825290925281209150611ef9828661210b565b9150611f079050828661233a565b505f811561201e57856001600160a01b031663d9caed1286886001600160a01b0316632495a5996040518163ffffffff1660e01b8152600401602060405180830381865afa158015611f5b573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611f7f9190613755565b856040518463ffffffff1660e01b8152600401611f9e93929190613770565b6020604051808303815f875af1158015611fba573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611fde9190613794565b90507fe6413aa0c789e437b0a06bf64b20926584f066c79a2d8b80a759c85472f7b0af8888888560405161201594939291906137ef565b60405180910390a15b5f612028846120ff565b519050805f036120f3576120668860da5f61204b610874368f90038f018f6136ba565b81526020019081526020015f2061282e90919063ffffffff16565b5060db5f61207c610874368d90038d018d6136ba565b815260208082019290925260409081015f9081208b82529092528120805463ffffffff191690556120ce9060da906120bc610874368e90038e018e6136ba565b81526020019081526020015f20612306565b5f036120f3576120f16120e9610874368c90038c018c6136ba565b60d89061282e565b505b50979650505050505050565b60605f6113fc83612839565b5f808080612122866001600160a01b038716612844565b9097909650945050505050565b4281101561215057604051630819bdcd60e01b815260040160405180910390fd5b6121646001600160a01b038516848461287c565b61084557604051638baa579f60e01b815260040160405180910390fd5b6001600160a01b0383165f90815260d16020526040812054849060ff166121bb57604051632efd965160e11b815260040160405180910390fd5b6121d06001600160a01b0385163387866128d0565b6040516311f9fbc960e21b81526001600160a01b038581166004830152602482018590528616906347e7ef24906044016020604051808303815f875af115801561221c573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906122409190613794565b91505f5f61224f88888661234e565b604051631e328e7960e11b81526001600160a01b038b811660048301528a8116602483015260448201849052606482018390529294509092507f000000000000000000000000000000000000000000000000000000000000000090911690633c651cf2906084015f604051808303815f87803b1580156122cd575f5ffd5b505af11580156122df573d5f5f3e3d5ffd5b50505050505050949350505050565b5f61102482612928565b5f8080806121228686612932565b5f611024825490565b5f6113fc838361295b565b5f611c31846001600160a01b03851684612981565b5f6113fc838361299d565b5f6113fc836001600160a01b0384166129e9565b5f806001600160a01b038516612377576040516316f2ccc960e01b815260040160405180910390fd5b825f03612397576040516342061b2560e11b815260040160405180910390fd5b60405162e7c78560e71b81526001600160a01b038681166004830152602482018590528516906373e3c280906044015f604051808303815f87803b1580156123dd575f5ffd5b505af11580156123ef573d5f5f3e3d5ffd5b5050506001600160a01b038087165f90815260cd602090815260408083209389168352929052908120549150819003612499576001600160a01b0386165f90815260ce60209081526040909120541061245b576040516301a1443960e31b815260040160405180910390fd5b6001600160a01b038681165f90815260ce602090815260408220805460018101825590835291200180546001600160a01b0319169187169190911790555b6124a38482613863565b6001600160a01b038088165f90815260cd60209081526040808320938a16835292905281902091909155517f5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62906124ff90889088908890613770565b60405180910390a1959294509192505050565b60605f61251e83612a05565b6040805160208082528183019092529192505f91906020820181803683375050509182525060208101929092525090565b6033546001600160a01b0316331461122d5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016107bb565b5f815f036125ca576040516342061b2560e11b815260040160405180910390fd5b6001600160a01b038085165f90815260cd60209081526040808320938716835292905220548083111561261057604051634b18b19360e01b815260040160405180910390fd5b6040516303e3e6eb60e01b81526001600160a01b038681166004830152602482018590528516906303e3e6eb906044015f604051808303815f87803b158015612657575f5ffd5b505af1158015612669573d5f5f3e3d5ffd5b5050505082816126799190613876565b6001600160a01b038087165f90815260cd602090815260408083209389168352929052908120829055909150819003611c3157611c318585612a2c565b60605f6113fc83612baa565b5f6126cb611afe565b60405161190160f01b6020820152602281019190915260428101839052606201604051602081830303815290604052805190602001209050919050565b60605f6127347f0000000000000000000000000000000000000000000000000000000000000000612512565b9050805f81518110612748576127486137ab565b016020908101516040516001600160f81b03199091169181019190915260210160405160208183030381529060405291505090565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156127d9573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906127fd9190613755565b6001600160a01b0316336001600160a01b03161461122d5760405163794821ff60e01b815260040160405180910390fd5b5f6113fc8383612c02565b6060611024826126b6565b5f818152600283016020526040812054819080612871576128658585612ce5565b92505f9150610bb29050565b600192509050610bb2565b5f5f5f6128898585612cf0565b90925090505f8160048111156128a1576128a1613889565b1480156128bf5750856001600160a01b0316826001600160a01b0316145b806114485750611448868686612d2f565b610845846323b872dd60e01b8585856040516024016128f193929190613770565b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b031990931692909217909152612e16565b5f61102482612306565b5f808061293f858561230f565b5f81815260029690960160205260409095205494959350505050565b5f825f018281548110612970576129706137ab565b905f5260205f200154905092915050565b5f8281526002840160205260408120829055611c31848461232f565b5f8181526001830160205260408120546129e257508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155611024565b505f611024565b5f81815260028301602052604081208190556113fc838361282e565b5f60ff8216601f81111561102457604051632cd44ac360e21b815260040160405180910390fd5b6001600160a01b0382165f90815260ce6020526040812054905b81811015612b3e576001600160a01b038481165f90815260ce6020526040902080549185169183908110612a7c57612a7c6137ab565b5f918252602090912001546001600160a01b031603612b36576001600160a01b0384165f90815260ce602052604090208054612aba90600190613876565b81548110612aca57612aca6137ab565b5f9182526020808320909101546001600160a01b03878116845260ce9092526040909220805491909216919083908110612b0657612b066137ab565b905f5260205f20015f6101000a8154816001600160a01b0302191690836001600160a01b03160217905550612b3e565b600101612a46565b818103612b5e57604051632df15a4160e11b815260040160405180910390fd5b6001600160a01b0384165f90815260ce60205260409020805480612b8457612b8461389d565b5f8281526020902081015f1990810180546001600160a01b031916905501905550505050565b6060815f01805480602002602001604051908101604052809291908181526020018280548015611a5d57602002820191905f5260205f20905b815481526020019060010190808311612be35750505050509050919050565b5f8181526001830160205260408120548015612cdc575f612c24600183613876565b85549091505f90612c3790600190613876565b9050818114612c96575f865f018281548110612c5557612c556137ab565b905f5260205f200154905080875f018481548110612c7557612c756137ab565b5f918252602080832090910192909255918252600188019052604090208390555b8554869080612ca757612ca761389d565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f905560019350505050611024565b5f915050611024565b5f6113fc8383612eee565b5f5f8251604103612d24576020830151604084015160608501515f1a612d1887828585612f05565b94509450505050610bb2565b505f90506002610bb2565b5f5f5f856001600160a01b0316631626ba7e60e01b8686604051602401612d579291906138b1565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909416939093179092529051612d9591906138c9565b5f60405180830381855afa9150503d805f8114612dcd576040519150601f19603f3d011682016040523d82523d5f602084013e612dd2565b606091505b5091509150818015612de657506020815110155b801561144857508051630b135d3f60e11b90612e0b9083016020908101908401613794565b149695505050505050565b5f612e6a826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316612fc29092919063ffffffff16565b905080515f1480612e8a575080806020019051810190612e8a9190613821565b612ee95760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b60648201526084016107bb565b505050565b5f81815260018301602052604081205415156113fc565b5f807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115612f3a57505f90506003612fb9565b604080515f8082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015612f8b573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b038116612fb3575f60019250925050612fb9565b91505f90505b94509492505050565b6060611c3184845f85855f5f866001600160a01b03168587604051612fe791906138c9565b5f6040518083038185875af1925050503d805f8114613021576040519150601f19603f3d011682016040523d82523d5f602084013e613026565b606091505b509150915061184d87838387606083156130a05782515f03613099576001600160a01b0385163b6130995760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016107bb565b5081611c31565b611c3183838151156130b55781518083602001fd5b8060405162461bcd60e51b81526004016107bb9190613501565b5f602082840312156130df575f5ffd5b5035919050565b6001600160a01b0381168114611174575f5ffd5b5f5f5f6060848603121561310c575f5ffd5b8335613117816130e6565b92506020840135613127816130e6565b929592945050506040919091013590565b5f60408284031215613148575f5ffd5b50919050565b5f5f5f60808486031215613160575f5ffd5b61316a8585613138565b9250604084013591506060840135613181816130e6565b809150509250925092565b5f5f5f5f6080858703121561319f575f5ffd5b84356131aa816130e6565b935060208501356131ba816130e6565b925060408501356131ca816130e6565b9396929550929360600135925050565b5f5f606083850312156131eb575f5ffd5b6131f58484613138565b946040939093013593505050565b5f8151808452602084019350602083015f5b8281101561323c5781516001600160a01b0316865260209586019590910190600101613215565b5093949350505050565b5f8151808452602084019350602083015f5b8281101561323c578151865260209586019590910190600101613258565b604081525f6132886040830185613203565b828103602084015261329a8185613246565b95945050505050565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f191681016001600160401b03811182821017156132df576132df6132a3565b604052919050565b5f5f5f5f5f5f60c087890312156132fc575f5ffd5b8635613307816130e6565b95506020870135613317816130e6565b945060408701359350606087013561332e816130e6565b92506080870135915060a08701356001600160401b0381111561334f575f5ffd5b8701601f8101891361335f575f5ffd5b80356001600160401b03811115613378576133786132a3565b61338b601f8201601f19166020016132b7565b8181528a602083850101111561339f575f5ffd5b816020840160208301375f602083830101528093505050509295509295509295565b604080825283519082018190525f9060208501906060840190835b818110156134035783516001600160a01b03168352602093840193909201916001016133dc565b505083810360208501526114488186613246565b602080825282518282018190525f918401906040840190835b8181101561346957835180516001600160a01b0316845260209081015163ffffffff168185015290930192604090920191600101613430565b509095945050505050565b5f5f5f5f60a08587031215613487575f5ffd5b6134918686613138565b93506040850135925060608501356134a8816130e6565b9396929550929360800135925050565b5f602082840312156134c8575f5ffd5b81356113fc816130e6565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f6113fc60208301846134d3565b5f60208284031215613523575f5ffd5b813560ff811681146113fc575f5ffd5b5f5f60208385031215613544575f5ffd5b82356001600160401b03811115613559575f5ffd5b8301601f81018513613569575f5ffd5b80356001600160401b0381111561357e575f5ffd5b8560208260051b8401011115613592575f5ffd5b6020919091019590945092505050565b5f604082840312156135b2575f5ffd5b6113fc8383613138565b602081525f6113fc6020830184613246565b5f5f5f5f5f5f60c087890312156135e3575f5ffd5b86356135ee816130e6565b955060208701356135fe816130e6565b9450604087013561360e816130e6565b959894975094956060810135955060808101359460a0909101359350915050565b5f5f60408385031215613640575f5ffd5b823561364b816130e6565b946020939093013593505050565b602081525f6113fc6020830184613203565b5f5f6040838503121561367c575f5ffd5b8235613687816130e6565b91506020830135613697816130e6565b809150509250929050565b803563ffffffff811681146136b5575f5ffd5b919050565b5f60408284031280156136cb575f5ffd5b50604080519081016001600160401b03811182821017156136ee576136ee6132a3565b60405282356136fc816130e6565b815261370a602084016136a2565b60208201529392505050565b8035613721816130e6565b6001600160a01b0316825263ffffffff61373d602083016136a2565b1660208301525050565b604081016110248284613716565b5f60208284031215613765575f5ffd5b81516113fc816130e6565b6001600160a01b039384168152919092166020820152604081019190915260600190565b5f602082840312156137a4575f5ffd5b5051919050565b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52601160045260245ffd5b63ffffffff8181168382160190811115611024576110246137bf565b60a081016137fd8287613716565b60408201949094526001600160a01b03929092166060830152608090910152919050565b5f60208284031215613831575f5ffd5b815180151581146113fc575f5ffd5b80516020808301519190811015613148575f1960209190910360031b1b16919050565b80820180821115611024576110246137bf565b81810381811115611024576110246137bf565b634e487b7160e01b5f52602160045260245ffd5b634e487b7160e01b5f52603160045260245ffd5b828152604060208201525f611c3160408301846134d3565b5f82518060208501845e5f92019182525091905056fea26469706673582212209349bdc8f3b89456e0f89a36df326a6ae97e0b0d600cbc6ebd03c74639d2c84f64736f6c634300081e0033",
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
