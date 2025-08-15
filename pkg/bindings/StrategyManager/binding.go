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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_allocationManager\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"},{\"name\":\"_delegation\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"_version\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_BURN_ADDRESS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEPOSIT_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addStrategiesToDepositWhitelist\",\"inputs\":[{\"name\":\"strategiesToWhitelist\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allocationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burnShares\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"calculateStrategyDepositDigestHash\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"clearBurnOrRedistributableShares\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"clearBurnOrRedistributableSharesByStrategy\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"depositIntoStrategy\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"depositShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositIntoStrategyWithSignature\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"depositShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"domainSeparator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBurnOrRedistributableCount\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBurnOrRedistributableShares\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBurnOrRedistributableShares\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBurnableShares\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDeposits\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPendingOperatorSets\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPendingSlashIds\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStakerStrategyList\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStrategiesWithBurnableShares\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"increaseBurnOrRedistributableShares\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"sharesToBurn\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialStrategyWhitelister\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"nonces\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeDepositShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"depositSharesToRemove\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeStrategiesFromDepositWhitelist\",\"inputs\":[{\"name\":\"strategiesToRemoveFromWhitelist\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStrategyWhitelister\",\"inputs\":[{\"name\":\"newStrategyWhitelister\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakerDepositShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakerStrategyList\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"strategies\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakerStrategyListLength\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"strategyIsWhitelistedForDeposit\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"whitelisted\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"strategyWhitelister\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawSharesAsTokens\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"BurnOrRedistributableSharesDecreased\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BurnOrRedistributableSharesIncreased\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BurnableSharesDecreased\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposit\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyAddedToDepositWhitelist\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyRemovedFromDepositWhitelist\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyWhitelisterChanged\",\"inputs\":[{\"name\":\"previousAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CurrentlyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNewPausedStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MaxStrategiesExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyDelegationManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyPauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyStrategyWhitelister\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnpauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SharesAmountTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SharesAmountZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignatureExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StakerAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategyAlreadyInSlash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategyNotFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategyNotWhitelisted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]}]",
	Bin: "0x610100604052348015610010575f5ffd5b506040516139d73803806139d783398101604081905261002f916101bc565b80808585856001600160a01b03811661005b576040516339b190bb60e11b815260040160405180910390fd5b6001600160a01b0390811660805291821660a0521660c05261007c81610093565b60e0525061008a90506100d9565b50505050610301565b5f5f829050601f815111156100c6578260405163305a27a960e01b81526004016100bd91906102a6565b60405180910390fd5b80516100d1826102db565b179392505050565b5f54610100900460ff16156101405760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b60648201526084016100bd565b5f5460ff9081161461018f575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146101a5575f5ffd5b50565b634e487b7160e01b5f52604160045260245ffd5b5f5f5f5f608085870312156101cf575f5ffd5b84516101da81610191565b60208601519094506101eb81610191565b60408601519093506101fc81610191565b60608601519092506001600160401b03811115610217575f5ffd5b8501601f81018713610227575f5ffd5b80516001600160401b03811115610240576102406101a8565b604051601f8201601f19908116603f011681016001600160401b038111828210171561026e5761026e6101a8565b604052818152828201602001891015610285575f5ffd5b8160208401602083015e5f6020838301015280935050505092959194509250565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b805160208083015191908110156102fb575f198160200360031b1b821691505b50919050565b60805160a05160c05160e05161365d61037a5f395f81816110a8015261247401525f81816105fa015281816108c101528181610da30152818161103a0152818161125501526120a901525f81816105a001528181610815015261141a01525f81816104b601528181611aaa01526124e3015261365d5ff3fe608060405234801561000f575f5ffd5b506004361061026b575f3560e01c80637ecebe001161014b578063ca8aa7c7116100bf578063f2fde38b11610084578063f2fde38b1461062f578063f3b4a00014610642578063f698da251461064c578063fabc1cbc14610654578063fd98042314610667578063fe243a171461067a575f5ffd5b8063ca8aa7c71461059b578063cbc2bd62146105c2578063de44acb6146105d5578063df5cf723146105f5578063e7a050aa1461061c575f5ffd5b80638da5cb5b116101105780638da5cb5b1461052b57806394f649dd1461053c578063967fc0d21461054f5780639ac01d6114610562578063b5d8b5b814610575578063c665670214610588575f5ffd5b80637ecebe001461047f578063829fca731461049e578063886f1195146104b157806388c10299146104f05780638b8aac3c14610503575f5ffd5b806350ff7225116101e25780635de08ff2116101a75780635de08ff2146103fc578063663c1de41461040f578063715018a614610431578063724af4231461043957806376fb162b1461044c5780637def15641461045f575f5ffd5b806350ff72251461037c57806354fd4d50146103a4578063595c6a67146103b95780635ac86ab7146103c15780635c975abb146103f4575f5ffd5b806332e89ace1161023357806332e89ace146102f157806336a8c500146103045780633f292b081461031a5780633fb99ca51461032f57806348825e94146103425780634b6d5d6e14610369575f5ffd5b8063136439dd1461026f5780631794bb3c146102845780632d44def6146102975780632eae418c146102bd57806331f8fb4c146102d0575b5f5ffd5b61028261027d366004612e33565b6106a4565b005b610282610292366004612e5e565b6106de565b6102aa6102a5366004612eb2565b610804565b6040519081526020015b60405180910390f35b6102826102cb366004612ef0565b6108b6565b6102e36102de366004612f3e565b610982565b6040516102b4929190612fda565b6102aa6102ff36600461304b565b610b10565b61030c610b95565b6040516102b4929190613125565b610322610cb0565b6040516102b4919061317b565b61028261033d3660046131d8565b610d98565b6102aa7f4337f82d142e41f2a8c10547cd8c859bddb92262a61058e77842e24d9dea922481565b61028261037736600461321c565b610ee0565b61038f61038a366004612e5e565b61102d565b604080519283526020830191909152016102b4565b6103ac6110a1565b6040516102b49190613265565b6102826110d1565b6103e46103cf366004613277565b609854600160ff9092169190911b9081161490565b60405190151581526020016102b4565b6098546102aa565b61028261040a366004613297565b6110e5565b6103e461041d36600461321c565b60d16020525f908152604090205460ff1681565b610282611238565b6102aa610447366004612e5e565b611249565b6102aa61045a366004612eb2565b6112a6565b61047261046d366004613306565b6112f5565b6040516102b49190613320565b6102aa61048d36600461321c565b60ca6020525f908152604090205481565b6102aa6104ac366004612f3e565b611327565b6104d87f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016102b4565b6104726104fe366004612f3e565b611361565b6102aa61051136600461321c565b6001600160a01b03165f90815260ce602052604090205490565b6033546001600160a01b03166104d8565b6102e361054a36600461321c565b611498565b60cb546104d8906001600160a01b031681565b6102aa610570366004613332565b61160f565b610282610583366004613297565b6116a0565b61028261059636600461321c565b6117e7565b6104d87f000000000000000000000000000000000000000000000000000000000000000081565b6104d86105d0366004613393565b61180a565b6105e86105e336600461321c565b61183e565b6040516102b491906133bd565b6104d87f000000000000000000000000000000000000000000000000000000000000000081565b6102aa61062a366004612e5e565b6118b1565b61028261063d36600461321c565b6118e4565b6104d8620e16e481565b6102aa61195a565b610282610662366004612e33565b611a13565b6102aa61067536600461321c565b611a80565b6102aa6106883660046133cf565b60cd60209081525f928352604080842090915290825290205481565b6106ac611a95565b60985481811681146106d15760405163c61dca5d60e01b815260040160405180910390fd5b6106da82611b38565b5050565b5f54610100900460ff16158080156106fc57505f54600160ff909116105b806107155750303b15801561071557505f5460ff166001145b61077d5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b5f805460ff19166001179055801561079e575f805461ff0019166101001790555b6107a782611b38565b6107b084611b75565b6107b983611bc6565b80156107fe575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050565b5f61080d611c2f565b6108a38484847f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316630f3df50e896040518263ffffffff1660e01b815260040161085f919061344f565b602060405180830381865afa15801561087a573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061089e919061345d565b611c88565b90506108af6001606555565b9392505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146108ff5760405163f739589b60e01b815260040160405180910390fd5b610907611c2f565b604051636ce5768960e11b81526001600160a01b0384169063d9caed129061093790879086908690600401613478565b6020604051808303815f875af1158015610953573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610977919061349c565b506107fe6001606555565b6060805f60d7816109a061099b368990038901896134b3565b611e90565b81526020019081526020015f205f8581526020019081526020015f2090505f6109c882611ef3565b90505f81516001600160401b038111156109e4576109e4613007565b604051908082528060200260200182016040528015610a0d578160200160208202803683370190505b5090505f82516001600160401b03811115610a2a57610a2a613007565b604051908082528060200260200182016040528015610a53578160200160208202803683370190505b5090505f5b8351811015610b0057838181518110610a7357610a7361350f565b6020026020010151838281518110610a8d57610a8d61350f565b60200260200101906001600160a01b031690816001600160a01b031681525050610ad9848281518110610ac257610ac261350f565b602002602001015186611eff90919063ffffffff16565b9050828281518110610aed57610aed61350f565b6020908102919091010152600101610a58565b50909450925050505b9250929050565b5f5f610b1b81611f23565b610b23611c2f565b6001600160a01b0385165f90815260ca6020526040902054610b5486610b4d818c8c8c878c61160f565b8688611f4e565b6001600160a01b0386165f90815260ca60205260409020600182019055610b7d868a8a8a611fa0565b925050610b8a6001606555565b509695505050505050565b6060805f610ba360d461210d565b90505f816001600160401b03811115610bbe57610bbe613007565b604051908082528060200260200182016040528015610be7578160200160208202803683370190505b5090505f826001600160401b03811115610c0357610c03613007565b604051908082528060200260200182016040528015610c2c578160200160208202803683370190505b5090505f5b83811015610ca5575f5f610c4660d484612117565b9150915081858481518110610c5d57610c5d61350f565b60200260200101906001600160a01b031690816001600160a01b03168152505080848481518110610c9057610c9061350f565b60209081029190910101525050600101610c31565b509094909350915050565b60605f610cbd60d8612125565b90505f816001600160401b03811115610cd857610cd8613007565b604051908082528060200260200182016040528015610d1c57816020015b604080518082019091525f8082526020820152815260200190600190039081610cf65790505b5090505f5b82811015610d9157610d6c610d3760d88361212e565b604080518082019091525f80825260208201525060408051808201909152606082901c815263ffffffff909116602082015290565b828281518110610d7e57610d7e61350f565b6020908102919091010152600101610d21565b5092915050565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610de15760405163f739589b60e01b815260040160405180910390fd5b610de9611c2f565b5f60d781610dff61099b368990038901896134b3565b815260208082019290925260409081015f90812087825290925290209050610e28818484612139565b610e455760405163ca354fa360e01b815260040160405180910390fd5b610e62610e5a61099b368890038801886134b3565b60d89061214e565b50610e978460da5f610e7c61099b368b90038b018b6134b3565b81526020019081526020015f2061214e90919063ffffffff16565b507f5f5209798bbac45a16d2dc3bc67319fab26ee00153916d6f07b69f8a134a1e8b85858585604051610ecd9493929190613523565b60405180910390a1506107fe6001606555565b610ee8611c2f565b5f610ef460d483611eff565b915050610f0260d483612159565b50604080516001600160a01b0384168152602081018390527fd9d082c3ec4f3a3ffa55c324939a06407f5fbcb87d5e0ce3b9508c92c84ed839910160405180910390a1801561101f57816001600160a01b031663d9caed12620e16e4846001600160a01b0316632495a5996040518163ffffffff1660e01b8152600401602060405180830381865afa158015610f9a573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610fbe919061345d565b846040518463ffffffff1660e01b8152600401610fdd93929190613478565b6020604051808303815f875af1158015610ff9573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061101d919061349c565b505b5061102a6001606555565b50565b5f80336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146110785760405163f739589b60e01b815260040160405180910390fd5b611080611c2f565b61108b85858561216d565b915091506110996001606555565b935093915050565b60606110cc7f00000000000000000000000000000000000000000000000000000000000000006122d5565b905090565b6110d9611a95565b6110e35f19611b38565b565b60cb546001600160a01b03163314611110576040516320ba3ff960e21b815260040160405180910390fd5b611118611c2f565b805f5b8181101561122c5760d15f8585848181106111385761113861350f565b905060200201602081019061114d919061321c565b6001600160a01b0316815260208101919091526040015f205460ff1661122457600160d15f8686858181106111845761118461350f565b9050602002016020810190611199919061321c565b6001600160a01b0316815260208101919091526040015f20805460ff19169115159190911790557f0c35b17d91c96eb2751cd456e1252f42a386e524ef9ff26ecc9950859fdc04fe8484838181106111f3576111f361350f565b9050602002016020810190611208919061321c565b6040516001600160a01b03909116815260200160405180910390a15b60010161111b565b50506106da6001606555565b611240612312565b6110e35f611b75565b5f336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146112935760405163f739589b60e01b815260040160405180910390fd5b61129b611c2f565b6108a384848461236c565b5f806112eb8360d7836112c161099b368b90038b018b6134b3565b81526020019081526020015f205f8781526020019081526020015f20611eff90919063ffffffff16565b9695505050505050565b606061132160da5f61130f61099b368790038701876134b3565b81526020019081526020015f2061241a565b92915050565b5f6108af60d78261134061099b368890038801886134b3565b81526020019081526020015f205f8481526020019081526020015f2061210d565b606061136b611c2f565b5f6113a560d78261138461099b368990038901896134b3565b81526020019081526020015f205f8581526020019081526020015f20611ef3565b80519091505f816001600160401b038111156113c3576113c3613007565b6040519080825280602002602001820160405280156113ec578160200160208202803683370190505b5090505f5b828110156114895761146487878684815181106114105761141061350f565b60200260200101517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316630f3df50e8c6040518263ffffffff1660e01b815260040161085f919061344f565b8282815181106114765761147661350f565b60209081029190910101526001016113f1565b50925050506113216001606555565b6001600160a01b0381165f90815260ce6020526040812054606091829190816001600160401b038111156114ce576114ce613007565b6040519080825280602002602001820160405280156114f7578160200160208202803683370190505b5090505f5b82811015611585576001600160a01b0386165f90815260cd6020908152604080832060ce909252822080549192918490811061153a5761153a61350f565b5f9182526020808320909101546001600160a01b0316835282019290925260400190205482518390839081106115725761157261350f565b60209081029190910101526001016114fc565b5060ce5f866001600160a01b03166001600160a01b031681526020019081526020015f2081818054806020026020016040519081016040528092919081815260200182805480156115fd57602002820191905f5260205f20905b81546001600160a01b031681526001909101906020018083116115df575b50505050509150935093505050915091565b604080517f4337f82d142e41f2a8c10547cd8c859bddb92262a61058e77842e24d9dea922460208201526001600160a01b03808916928201929092528187166060820152908516608082015260a0810184905260c0810183905260e081018290525f90611695906101000160405160208183030381529060405280519060200120612426565b979650505050505050565b60cb546001600160a01b031633146116cb576040516320ba3ff960e21b815260040160405180910390fd5b6116d3611c2f565b805f5b8181101561122c5760d15f8585848181106116f3576116f361350f565b9050602002016020810190611708919061321c565b6001600160a01b0316815260208101919091526040015f205460ff16156117df575f60d15f86868581811061173f5761173f61350f565b9050602002016020810190611754919061321c565b6001600160a01b0316815260208101919091526040015f20805460ff19169115159190911790557f4074413b4b443e4e58019f2855a8765113358c7c72e39509c6af45fc0f5ba0308484838181106117ae576117ae61350f565b90506020020160208101906117c3919061321c565b6040516001600160a01b03909116815260200160405180910390a15b6001016116d6565b6117ef612312565b6117f7611c2f565b61180081611bc6565b61102a6001606555565b60ce602052815f5260405f208181548110611823575f80fd5b5f918252602090912001546001600160a01b03169150829050565b6001600160a01b0381165f90815260ce60209081526040918290208054835181840281018401909452808452606093928301828280156118a557602002820191905f5260205f20905b81546001600160a01b03168152600190910190602001808311611887575b50505050509050919050565b5f5f6118bc81611f23565b6118c4611c2f565b6118d033868686611fa0565b91506118dc6001606555565b509392505050565b6118ec612312565b6001600160a01b0381166119515760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610774565b61102a81611b75565b60408051808201909152600a81526922b4b3b2b72630bcb2b960b11b6020909101525f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f7f71b625cfad44bac63b13dba07f2e1d6084ee04b6f8752101ece6126d584ee6ea6119c761246c565b805160209182012060408051928301949094529281019190915260608101919091524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b611a1b6124e1565b60985480198219811614611a425760405163c61dca5d60e01b815260040160405180910390fd5b609882905560405182815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c9060200160405180910390a25050565b5f5f611a8d60d484611eff565b949350505050565b60405163237dfb4760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906346fbf68e90602401602060405180830381865afa158015611af7573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611b1b9190613555565b6110e357604051631d77d47760e21b815260040160405180910390fd5b609881905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a250565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b60cb54604080516001600160a01b03928316815291831660208301527f4264275e593955ff9d6146a51a4525f6ddace2e81db9391abcc9d1ca48047d29910160405180910390a160cb80546001600160a01b0319166001600160a01b0392909216919091179055565b600260655403611c815760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610774565b6002606555565b5f8060d781611c9f61099b368a90038a018a6134b3565b815260208082019290925260409081015f90812088825290925281209150611cc78286611eff565b9150611cd590508286612159565b505f8115611dec57856001600160a01b031663d9caed1286886001600160a01b0316632495a5996040518163ffffffff1660e01b8152600401602060405180830381865afa158015611d29573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611d4d919061345d565b856040518463ffffffff1660e01b8152600401611d6c93929190613478565b6020604051808303815f875af1158015611d88573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611dac919061349c565b90507fe6413aa0c789e437b0a06bf64b20926584f066c79a2d8b80a759c85472f7b0af88888885604051611de39493929190613523565b60405180910390a15b5f611df684611ef3565b519050805f03611e8457611e348860da5f611e1961099b368f90038f018f6134b3565b81526020019081526020015f2061259290919063ffffffff16565b50611e5f60da5f611e4d61099b368e90038e018e6134b3565b81526020019081526020015f20612125565b5f03611e8457611e82611e7a61099b368c90038c018c6134b3565b60d890612592565b505b50979650505050505050565b5f815f0151826020015163ffffffff16604051602001611edb92919060609290921b6bffffffffffffffffffffffff1916825260a01b6001600160a01b031916601482015260200190565b60405160208183030381529060405261132190613574565b60605f6108af8361259d565b5f808080611f16866001600160a01b0387166125a8565b9097909650945050505050565b609854600160ff83161b9081160361102a5760405163840a48d560e01b815260040160405180910390fd5b42811015611f6f57604051630819bdcd60e01b815260040160405180910390fd5b611f836001600160a01b03851684846125e0565b6107fe57604051638baa579f60e01b815260040160405180910390fd5b6001600160a01b0383165f90815260d16020526040812054849060ff16611fda57604051632efd965160e11b815260040160405180910390fd5b611fef6001600160a01b038516338786612634565b6040516311f9fbc960e21b81526001600160a01b038581166004830152602482018590528616906347e7ef24906044016020604051808303815f875af115801561203b573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061205f919061349c565b91505f5f61206e88888661216d565b604051631e328e7960e11b81526001600160a01b038b811660048301528a8116602483015260448201849052606482018390529294509092507f000000000000000000000000000000000000000000000000000000000000000090911690633c651cf2906084015f604051808303815f87803b1580156120ec575f5ffd5b505af11580156120fe573d5f5f3e3d5ffd5b50505050505050949350505050565b5f6113218261268c565b5f808080611f168686612696565b5f611321825490565b5f6108af83836126bf565b5f611a8d846001600160a01b038516846126e5565b5f6108af8383612701565b5f6108af836001600160a01b03841661274d565b5f806001600160a01b038516612196576040516316f2ccc960e01b815260040160405180910390fd5b825f036121b6576040516342061b2560e11b815260040160405180910390fd5b6001600160a01b038086165f90815260cd602090815260408083209388168352929052908120549081900361225c576001600160a01b0386165f90815260ce60209081526040909120541061221e576040516301a1443960e31b815260040160405180910390fd5b6001600160a01b038681165f90815260ce602090815260408220805460018101825590835291200180546001600160a01b0319169187169190911790555b61226684826135ab565b6001600160a01b038088165f90815260cd60209081526040808320938a16835292905281902091909155517f5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62906122c290889088908890613478565b60405180910390a1959294509192505050565b60605f6122e183612769565b6040805160208082528183019092529192505f91906020820181803683375050509182525060208101929092525090565b6033546001600160a01b031633146110e35760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610774565b5f815f0361238d576040516342061b2560e11b815260040160405180910390fd5b6001600160a01b038085165f90815260cd6020908152604080832093871683529290522054808311156123d357604051634b18b19360e01b815260040160405180910390fd5b6123dd83826135be565b6001600160a01b038087165f90815260cd602090815260408083209389168352929052908120829055909150819003611a8d57611a8d8585612790565b60605f6108af8361290e565b5f61242f61195a565b60405161190160f01b6020820152602281019190915260428101839052606201604051602081830303815290604052805190602001209050919050565b60605f6124987f00000000000000000000000000000000000000000000000000000000000000006122d5565b9050805f815181106124ac576124ac61350f565b016020908101516040516001600160f81b03199091169181019190915260210160405160208183030381529060405291505090565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561253d573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612561919061345d565b6001600160a01b0316336001600160a01b0316146110e35760405163794821ff60e01b815260040160405180910390fd5b5f6108af8383612966565b60606113218261241a565b5f8181526002830160205260408120548190806125d5576125c98585612a49565b92505f9150610b099050565b600192509050610b09565b5f5f5f6125ed8585612a54565b90925090505f816004811115612605576126056135d1565b1480156126235750856001600160a01b0316826001600160a01b0316145b806112eb57506112eb868686612a93565b6107fe846323b872dd60e01b85858560405160240161265593929190613478565b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b031990931692909217909152612b7a565b5f61132182612125565b5f80806126a3858561212e565b5f81815260029690960160205260409095205494959350505050565b5f825f0182815481106126d4576126d461350f565b905f5260205f200154905092915050565b5f8281526002840160205260408120829055611a8d848461214e565b5f81815260018301602052604081205461274657508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155611321565b505f611321565b5f81815260028301602052604081208190556108af8383612592565b5f60ff8216601f81111561132157604051632cd44ac360e21b815260040160405180910390fd5b6001600160a01b0382165f90815260ce6020526040812054905b818110156128a2576001600160a01b038481165f90815260ce60205260409020805491851691839081106127e0576127e061350f565b5f918252602090912001546001600160a01b03160361289a576001600160a01b0384165f90815260ce60205260409020805461281e906001906135be565b8154811061282e5761282e61350f565b5f9182526020808320909101546001600160a01b03878116845260ce909252604090922080549190921691908390811061286a5761286a61350f565b905f5260205f20015f6101000a8154816001600160a01b0302191690836001600160a01b031602179055506128a2565b6001016127aa565b8181036128c257604051632df15a4160e11b815260040160405180910390fd5b6001600160a01b0384165f90815260ce602052604090208054806128e8576128e86135e5565b5f8281526020902081015f1990810180546001600160a01b031916905501905550505050565b6060815f018054806020026020016040519081016040528092919081815260200182805480156118a557602002820191905f5260205f20905b8154815260200190600101908083116129475750505050509050919050565b5f8181526001830160205260408120548015612a40575f6129886001836135be565b85549091505f9061299b906001906135be565b90508181146129fa575f865f0182815481106129b9576129b961350f565b905f5260205f200154905080875f0184815481106129d9576129d961350f565b5f918252602080832090910192909255918252600188019052604090208390555b8554869080612a0b57612a0b6135e5565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f905560019350505050611321565b5f915050611321565b5f6108af8383612c52565b5f5f8251604103612a88576020830151604084015160608501515f1a612a7c87828585612c69565b94509450505050610b09565b505f90506002610b09565b5f5f5f856001600160a01b0316631626ba7e60e01b8686604051602401612abb9291906135f9565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909416939093179092529051612af99190613611565b5f60405180830381855afa9150503d805f8114612b31576040519150601f19603f3d011682016040523d82523d5f602084013e612b36565b606091505b5091509150818015612b4a57506020815110155b80156112eb57508051630b135d3f60e11b90612b6f908301602090810190840161349c565b149695505050505050565b5f612bce826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316612d269092919063ffffffff16565b905080515f1480612bee575080806020019051810190612bee9190613555565b612c4d5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608401610774565b505050565b5f81815260018301602052604081205415156108af565b5f807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115612c9e57505f90506003612d1d565b604080515f8082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015612cef573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b038116612d17575f60019250925050612d1d565b91505f90505b94509492505050565b6060611a8d84845f85855f5f866001600160a01b03168587604051612d4b9190613611565b5f6040518083038185875af1925050503d805f8114612d85576040519150601f19603f3d011682016040523d82523d5f602084013e612d8a565b606091505b50915091506116958783838760608315612e045782515f03612dfd576001600160a01b0385163b612dfd5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610774565b5081611a8d565b611a8d8383815115612e195781518083602001fd5b8060405162461bcd60e51b81526004016107749190613265565b5f60208284031215612e43575f5ffd5b5035919050565b6001600160a01b038116811461102a575f5ffd5b5f5f5f60608486031215612e70575f5ffd5b8335612e7b81612e4a565b92506020840135612e8b81612e4a565b929592945050506040919091013590565b5f60408284031215612eac575f5ffd5b50919050565b5f5f5f60808486031215612ec4575f5ffd5b612ece8585612e9c565b9250604084013591506060840135612ee581612e4a565b809150509250925092565b5f5f5f5f60808587031215612f03575f5ffd5b8435612f0e81612e4a565b93506020850135612f1e81612e4a565b92506040850135612f2e81612e4a565b9396929550929360600135925050565b5f5f60608385031215612f4f575f5ffd5b612f598484612e9c565b946040939093013593505050565b5f8151808452602084019350602083015f5b82811015612fa05781516001600160a01b0316865260209586019590910190600101612f79565b5093949350505050565b5f8151808452602084019350602083015f5b82811015612fa0578151865260209586019590910190600101612fbc565b604081525f612fec6040830185612f67565b8281036020840152612ffe8185612faa565b95945050505050565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f191681016001600160401b038111828210171561304357613043613007565b604052919050565b5f5f5f5f5f5f60c08789031215613060575f5ffd5b863561306b81612e4a565b9550602087013561307b81612e4a565b945060408701359350606087013561309281612e4a565b92506080870135915060a08701356001600160401b038111156130b3575f5ffd5b8701601f810189136130c3575f5ffd5b80356001600160401b038111156130dc576130dc613007565b6130ef601f8201601f191660200161301b565b8181528a6020838501011115613103575f5ffd5b816020840160208301375f602083830101528093505050509295509295509295565b604080825283519082018190525f9060208501906060840190835b818110156131675783516001600160a01b0316835260209384019390920191600101613140565b505083810360208501526112eb8186612faa565b602080825282518282018190525f918401906040840190835b818110156131cd57835180516001600160a01b0316845260209081015163ffffffff168185015290930192604090920191600101613194565b509095945050505050565b5f5f5f5f60a085870312156131eb575f5ffd5b6131f58686612e9c565b935060408501359250606085013561320c81612e4a565b9396929550929360800135925050565b5f6020828403121561322c575f5ffd5b81356108af81612e4a565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f6108af6020830184613237565b5f60208284031215613287575f5ffd5b813560ff811681146108af575f5ffd5b5f5f602083850312156132a8575f5ffd5b82356001600160401b038111156132bd575f5ffd5b8301601f810185136132cd575f5ffd5b80356001600160401b038111156132e2575f5ffd5b8560208260051b84010111156132f6575f5ffd5b6020919091019590945092505050565b5f60408284031215613316575f5ffd5b6108af8383612e9c565b602081525f6108af6020830184612faa565b5f5f5f5f5f5f60c08789031215613347575f5ffd5b863561335281612e4a565b9550602087013561336281612e4a565b9450604087013561337281612e4a565b959894975094956060810135955060808101359460a0909101359350915050565b5f5f604083850312156133a4575f5ffd5b82356133af81612e4a565b946020939093013593505050565b602081525f6108af6020830184612f67565b5f5f604083850312156133e0575f5ffd5b82356133eb81612e4a565b915060208301356133fb81612e4a565b809150509250929050565b803563ffffffff81168114613419575f5ffd5b919050565b803561342981612e4a565b6001600160a01b0316825263ffffffff61344560208301613406565b1660208301525050565b60408101611321828461341e565b5f6020828403121561346d575f5ffd5b81516108af81612e4a565b6001600160a01b039384168152919092166020820152604081019190915260600190565b5f602082840312156134ac575f5ffd5b5051919050565b5f60408284031280156134c4575f5ffd5b50604080519081016001600160401b03811182821017156134e7576134e7613007565b60405282356134f581612e4a565b815261350360208401613406565b60208201529392505050565b634e487b7160e01b5f52603260045260245ffd5b60a08101613531828761341e565b60408201949094526001600160a01b03929092166060830152608090910152919050565b5f60208284031215613565575f5ffd5b815180151581146108af575f5ffd5b80516020808301519190811015612eac575f1960209190910360031b1b16919050565b634e487b7160e01b5f52601160045260245ffd5b8082018082111561132157611321613597565b8181038181111561132157611321613597565b634e487b7160e01b5f52602160045260245ffd5b634e487b7160e01b5f52603160045260245ffd5b828152604060208201525f611a8d6040830184613237565b5f82518060208501845e5f92019182525091905056fea264697066735822122051ad62e459ea065f2ef5780551bb5733a88f61bf29f219bce24fe29a3658744d64736f6c634300081b0033",
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
