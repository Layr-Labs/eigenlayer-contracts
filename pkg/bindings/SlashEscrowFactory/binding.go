// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package SlashEscrowFactory

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

// SlashEscrowFactoryMetaData contains all meta data concerning the SlashEscrowFactory contract.
var SlashEscrowFactoryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_allocationManager\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"},{\"name\":\"_strategyManager\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"_slashEscrowImplementation\",\"type\":\"address\",\"internalType\":\"contractISlashEscrow\"},{\"name\":\"_version\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allocationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"computeSlashEscrowSalt\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getEscrowCompleteBlock\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getEscrowStartBlock\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getGlobalEscrowDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPendingEscrows\",\"inputs\":[],\"outputs\":[{\"name\":\"operatorSets\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"isRedistributing\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"},{\"name\":\"slashIds\",\"type\":\"uint256[][]\",\"internalType\":\"uint256[][]\"},{\"name\":\"completeBlocks\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPendingOperatorSets\",\"inputs\":[],\"outputs\":[{\"name\":\"operatorSets\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPendingSlashIds\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPendingStrategiesForSlashId\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPendingStrategiesForSlashIds\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"strategies\",\"type\":\"address[][]\",\"internalType\":\"contractIStrategy[][]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPendingUnderlyingAmountForStrategy\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSlashEscrow\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractISlashEscrow\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStrategyEscrowDelay\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalPendingOperatorSets\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalPendingSlashIds\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalPendingStrategiesForSlashId\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"initialGlobalDelayBlocks\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initiateSlashEscrow\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isDeployedSlashEscrow\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isDeployedSlashEscrow\",\"inputs\":[{\"name\":\"slashEscrow\",\"type\":\"address\",\"internalType\":\"contractISlashEscrow\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isEscrowPaused\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isPendingOperatorSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isPendingSlashId\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseEscrow\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releaseSlashEscrow\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseSlashEscrowByStrategy\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGlobalEscrowDelay\",\"inputs\":[{\"name\":\"delay\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStrategyEscrowDelay\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"delay\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slashEscrowImplementation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractISlashEscrow\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"strategyManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpauseEscrow\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"EscrowComplete\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EscrowPaused\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EscrowUnpaused\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GlobalEscrowDelaySet\",\"inputs\":[{\"name\":\"delay\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StartEscrow\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"startBlock\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyEscrowDelaySet\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"delay\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CurrentlyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EscrowDelayNotElapsed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EscrowNotMature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNewPausedStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyPauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyRedistributionRecipient\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyStrategyManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnpauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]}]",
	Bin: "0x610120604052348015610010575f5ffd5b50604051612e89380380612e8983398101604081905261002f916101c1565b6001600160a01b0380861660805280851660a05280831660c05281908490811661006c576040516339b190bb60e11b815260040160405180910390fd5b6001600160a01b031660e05261008181610098565b610100525061008e6100de565b5050505050610319565b5f5f829050601f815111156100cb578260405163305a27a960e01b81526004016100c291906102be565b60405180910390fd5b80516100d6826102f3565b179392505050565b5f54610100900460ff16156101455760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b60648201526084016100c2565b5f5460ff90811614610194575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146101aa575f5ffd5b50565b634e487b7160e01b5f52604160045260245ffd5b5f5f5f5f5f60a086880312156101d5575f5ffd5b85516101e081610196565b60208701519095506101f181610196565b604087015190945061020281610196565b606087015190935061021381610196565b60808701519092506001600160401b0381111561022e575f5ffd5b8601601f8101881361023e575f5ffd5b80516001600160401b03811115610257576102576101ad565b604051601f8201601f19908116603f011681016001600160401b0381118282101715610285576102856101ad565b6040528181528282016020018a101561029c575f5ffd5b8160208401602083015e5f602083830101528093505050509295509295909350565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b80516020808301519190811015610313575f198160200360031b1b821691505b50919050565b60805160a05160c05160e05161010051612ae66103a35f395f610a5601525f818161049601528181611c980152611f2e01525f8181610515015281816108f501528181611ad30152611ee701525f818161032a0152818161079701528181610b620152610f1701525f81816105b7015281816106ff01528181610aca01526115500152612ae65ff3fe608060405234801561000f575f5ffd5b506004361061024a575f3560e01c806378748459116101405780639b122356116100bf578063c50f4e4811610084578063c50f4e481461057d578063c8b5330c14610595578063ca8aa7c7146105b2578063e7ed076b146105d9578063f2fde38b146105ec578063fabc1cbc146105ff575f5ffd5b80639b122356146104fd578063a3c6564114610510578063a56b21e414610537578063b23ff83b14610557578063c2de70961461056a575f5ffd5b8063886f119511610105578063886f1195146104915780638a65d2d2146104b85780638d5d4036146104c65780638da5cb5b146104d95780638fc46be5146104ea575f5ffd5b8063787484591461043057806378cb9600146104435780637a9676111461044b5780637def15641461045e57806387420b071461047e575f5ffd5b8063595c6a67116101cc5780636729b5db116101915780636729b5db146103cf5780636c5ac81e146103e25780637130c423146103f5578063715018a61461041557806371e166e71461041d575f5ffd5b8063595c6a67146103765780635ac86ab71461037e5780635c975abb146103a15780635e0a64c5146103a95780635ffa5a81146103bc575f5ffd5b8063277a9f0e11610212578063277a9f0e146102e75780633453b234146102fa57806339b70e38146103255780633f292b081461034c57806354fd4d5014610361575f5ffd5b80630310f3e61461024e57806305a4dfbb1461027b5780630e475b171461029c578063136439dd146102b157806319f3db26146102c4575b5f5ffd5b61026161025c36600461233a565b610612565b60405163ffffffff90911681526020015b60405180910390f35b61028e61028936600461233a565b6106a4565b604051908152602001610272565b6102af6102aa366004612379565b6106dc565b005b6102af6102bf3660046123b7565b610838565b6102d76102d23660046123ce565b610872565b6040519015158152602001610272565b61028e6102f53660046123ce565b610887565b61030d6103083660046123ce565b6108c6565b6040516001600160a01b039091168152602001610272565b61030d7f000000000000000000000000000000000000000000000000000000000000000081565b610354610958565b6040516102729190612447565b610369610a4f565b6040516102729190612459565b6102af610a7f565b6102d761038c36600461248e565b609854600160ff9092169190911b9081161490565b60985461028e565b6102af6103b73660046124ae565b610a93565b6102af6103ca3660046123ce565b610aa7565b61028e6103dd3660046124c7565b610c65565b6102af6103f03660046124e1565b610c90565b61040861040336600461233a565b610db6565b6040516102729190612555565b6102af610e88565b6102d761042b3660046124c7565b610e99565b6102d761043e3660046123ce565b610eb7565b61028e610f01565b6102af610459366004612379565b610f0c565b61047161046c366004612567565b611090565b60405161027291906125b1565b6102d761048c3660046123ce565b6110b3565b61030d7f000000000000000000000000000000000000000000000000000000000000000081565b60075463ffffffff16610261565b6102616104d43660046125c3565b6110e8565b6065546001600160a01b031661030d565b6102af6104f83660046125de565b611126565b6102af61050b3660046123ce565b611198565b61030d7f000000000000000000000000000000000000000000000000000000000000000081565b61054a610545366004612567565b61126b565b6040516102729190612611565b61028e610565366004612379565b611329565b61028e6105783660046123ce565b611408565b61058561144d565b60405161027294939291906126fc565b6102d76105a33660046125c3565b6001600160a01b03163b151590565b61030d7f000000000000000000000000000000000000000000000000000000000000000081565b6102af6105e73660046123ce565b6117a0565b6102af6105fa3660046125c3565b611875565b6102af61060d3660046123b7565b6118eb565b5f5f61061e8484610db6565b90505f805b8251811015610672575f61064f848381518110610642576106426127ca565b60200260200101516110e8565b90508263ffffffff168163ffffffff161115610669578092505b50600101610623565b508063ffffffff1661068486866106a4565b61068e91906127f2565b6106999060016127f2565b925050505b92915050565b5f60055f6106b185611958565b815260208082019290925260409081015f90812085825290925290205463ffffffff16905092915050565b5f6106e6816119bb565b60405163079efa8760e11b81525f906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690630f3df50e90610734908890600401612836565b602060405180830381865afa15801561074f573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107739190612844565b90506107808585836119e6565b6040516316a26f7b60e11b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690632d44def6906107d09088908890889060040161285f565b6020604051808303815f875af11580156107ec573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610810919061288c565b50610827858561082088886108c6565b8487611a8d565b6108318585611b9d565b5050505050565b610840611c83565b60985481811681146108655760405163c61dca5d60e01b815260040160405180910390fd5b61086e82611d26565b5050565b5f6108806105a384846108c6565b9392505050565b5f6108806004826108a56108a036889003880188612567565b611958565b81526020019081526020015f205f8481526020019081526020015f20611d63565b5f6108806108d48484611408565b6040513060388201526f5af43d82803e903d91602b57fd5bf3ff60248201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166014820152733d602d80600a3d3981f3363d3d373d3d3d363d73815260588101919091526037600c8201206078820152605560439091012090565b60605f6109656001611d6c565b9050805167ffffffffffffffff8111156109815761098161226b565b6040519080825280602002602001820160405280156109c557816020015b604080518082019091525f808252602082015281526020019060019003908161099f5790505b5091505f5b8151811015610a4a57610a258282815181106109e8576109e86127ca565b6020026020010151604080518082019091525f80825260208201525060408051808201909152606082901c815263ffffffff909116602082015290565b838281518110610a3757610a376127ca565b60209081029190910101526001016109ca565b505090565b6060610a7a7f0000000000000000000000000000000000000000000000000000000000000000611d78565b905090565b610a87611c83565b610a915f19611d26565b565b610a9b611db5565b610aa481611e0f565b50565b5f610ab1816119bb565b60405163079efa8760e11b81525f906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690630f3df50e90610aff908790600401612836565b602060405180830381865afa158015610b1a573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b3e9190612844565b9050610b4b8484836119e6565b6040516388c1029960e01b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906388c1029990610b9990879087906004016128a3565b5f604051808303815f875af1158015610bb4573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610bdb91908101906128be565b505f610c16600482610bf56108a0368a90038a018a612567565b81526020019081526020015f205f8681526020019081526020015f20611d6c565b90505f5b8151811015610c5a57610c528686610c3289896108c6565b86868681518110610c4557610c456127ca565b6020026020010151611a8d565b600101610c1a565b506108318585611b9d565b5f61069e600382610c7e6108a036879003870187612567565b81526020019081526020015f20611d63565b5f54610100900460ff1615808015610cae57505f54600160ff909116105b80610cc75750303b158015610cc757505f5460ff166001145b610d2f5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b5f805460ff191660011790558015610d50575f805461ff0019166101001790555b610d5984611e5d565b610d6283611d26565b610d6b82611e0f565b8015610db0575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050565b60605f60045f610dc586611958565b81526020019081526020015f205f8481526020019081526020015f2090505f610ded82611d63565b90508067ffffffffffffffff811115610e0857610e0861226b565b604051908082528060200260200182016040528015610e31578160200160208202803683370190505b5092505f5b81811015610e7f575f610e498483611eae565b905080858381518110610e5e57610e5e6127ca565b6001600160a01b039092166020928302919091019091015250600101610e36565b50505092915050565b610e90611db5565b610a915f611e5d565b5f61069e610eaf6108a036859003850185612567565b600190611eb9565b5f600681610ecd6108a036879003870187612567565b815260208082019290925260409081015f90812085825290925290205460ff16806108805750609854600190811614610880565b5f610a7a6001611d63565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610f55576040516348da714f60e01b815260040160405180910390fd5b5f600381610f6b6108a036889003880188612567565b81526020019081526020015f2090505f60045f610f92878036038101906108a09190612567565b81526020019081526020015f205f8581526020019081526020015f209050610fba8585610872565b61104157610fc88585611ed0565b610fe5610fdd6108a036889003880188612567565b600190611f0d565b50610ff08285611f0d565b504360055f6110076108a0368a90038a018a612567565b81526020019081526020015f205f8681526020019081526020015f205f6101000a81548163ffffffff021916908363ffffffff1602179055505b61104b8184611f18565b507f3afae24c1d3dd2ce3649054ad82458a8c9967ebd9ce10a9a5a3d059f55bfaedb85858543604051611081949392919061296a565b60405180910390a15050505050565b606061069e60035f6110a185611958565b81526020019081526020015f20611d6c565b5f610880826003836110cd6108a036899003890189612567565b81526020019081526020015f20611eb990919063ffffffff16565b6007546001600160a01b0382165f90815260086020526040812054909163ffffffff908116911681811161111c578161111e565b805b949350505050565b61112e611db5565b6001600160a01b0382165f81815260086020908152604091829020805463ffffffff191663ffffffff86169081179091558251938452908301527f5d2b33f07ae22a809e79005f96ffac70c3715df85a3b011b025e0a86a23a007b91015b60405180910390a15050565b6111a0611f2c565b60065f6111b56108a036869003860186612567565b815260208082019290925260409081015f90812084825290925290205460ff166111f25760405163c61dca5d60e01b815260040160405180910390fd5b5f6006816112086108a036879003870187612567565b81526020019081526020015f205f8381526020019081526020015f205f6101000a81548160ff0219169083151502179055507fb8877c6bf02d5f6603188eb653c9269271836b75b2012a5d7f5f233e7cf2f241828260405161118c9291906128a3565b60605f60035f61127a85611958565b81526020019081526020015f2090505f61129382611d63565b90508067ffffffffffffffff8111156112ae576112ae61226b565b6040519080825280602002602001820160405280156112e157816020015b60608152602001906001900390816112cc5790505b5092505f5b81811015611321576112fc856104038584611eae565b84828151811061130e5761130e6127ca565b60209081029190910101526001016112e6565b505050919050565b5f816001600160a01b0316632495a5996040518163ffffffff1660e01b8152600401602060405180830381865afa158015611366573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061138a9190612844565b6001600160a01b03166370a082316113a286866108c6565b6040516001600160e01b031960e084901b1681526001600160a01b039091166004820152602401602060405180830381865afa1580156113e4573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061111e919061288c565b5f61141b6108a036859003850185612567565b604080516020810192909252810183905260600160405160208183030381529060405280519060200120905092915050565b60608060608061145b610958565b9350835167ffffffffffffffff8111156114775761147761226b565b6040519080825280602002602001820160405280156114a0578160200160208202803683370190505b509250835167ffffffffffffffff8111156114bd576114bd61226b565b6040519080825280602002602001820160405280156114f057816020015b60608152602001906001900390816114db5790505b509150835167ffffffffffffffff81111561150d5761150d61226b565b60405190808252806020026020018201604052801561154057816020015b606081526020019060019003908161152b5790505b5090505f5b8451811015611799577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f231bd0886838151811061158f5761158f6127ca565b60200260200101516040518263ffffffff1660e01b81526004016115b391906129a2565b602060405180830381865afa1580156115ce573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906115f291906129c8565b848281518110611604576116046127ca565b60200260200101901515908115158152505061163885828151811061162b5761162b6127ca565b6020026020010151611090565b83828151811061164a5761164a6127ca565b6020026020010181905250828181518110611667576116676127ca565b60200260200101515167ffffffffffffffff8111156116885761168861226b565b6040519080825280602002602001820160405280156116b1578160200160208202803683370190505b508282815181106116c4576116c46127ca565b60209081029190910101525f5b8382815181106116e3576116e36127ca565b60200260200101515181101561179057611748868381518110611708576117086127ca565b6020026020010151858481518110611722576117226127ca565b6020026020010151838151811061173b5761173b6127ca565b6020026020010151610612565b83838151811061175a5761175a6127ca565b60200260200101518281518110611773576117736127ca565b63ffffffff909216602092830291909101909101526001016116d1565b50600101611545565b5090919293565b6117a8611c83565b60065f6117bd6108a036869003860186612567565b815260208082019290925260409081015f90812084825290925290205460ff16156117fb5760405163c61dca5d60e01b815260040160405180910390fd5b600160065f6118126108a036879003870187612567565b81526020019081526020015f205f8381526020019081526020015f205f6101000a81548160ff0219169083151502179055507f050add19b1a78a4240cdebc8747899275f2dd070c88e83904a37ff7d1a539744828260405161118c9291906128a3565b61187d611db5565b6001600160a01b0381166118e25760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610d26565b610aa481611e5d565b6118f3611f2c565b6098548019821981161461191a5760405163c61dca5d60e01b815260040160405180910390fd5b609882905560405182815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c9060200160405180910390a25050565b5f815f0151826020015163ffffffff166040516020016119a392919060609290921b6bffffffffffffffffffffffff1916825260a01b6001600160a01b031916601482015260200190565b60405160208183030381529060405261069e906129e7565b609854600160ff83161b90811603610aa45760405163840a48d560e01b815260040160405180910390fd5b6001600160a01b038116620e16e414611a2257336001600160a01b03821614611a225760405163720116ad60e11b815260040160405180910390fd5b611a2c8383610eb7565b15611a4a5760405163840a48d560e01b815260040160405180910390fd5b611a62611a5c36859003850185612567565b83610612565b63ffffffff16431015611a88576040516331caa72160e01b815260040160405180910390fd5b505050565b5f600481611aa36108a0368a90038a018a612567565b81526020019081526020015f205f8681526020019081526020015f209050836001600160a01b031663ff491e65307f0000000000000000000000000000000000000000000000000000000000000000898988886040518763ffffffff1660e01b8152600401611b1796959493929190612a0a565b5f604051808303815f87803b158015611b2e575f5ffd5b505af1158015611b40573d5f5f3e3d5ffd5b50505050611b578282611fdd90919063ffffffff16565b507f32be306ad5a833e756b7cb9724d5312afe0feda6163bfc2dd98ee713346a9abc86868486604051611b8d9493929190612a56565b60405180910390a1505050505050565b60015f600381611bb56108a036889003880188612567565b81526020019081526020015f2090505f611c0060045f611bdf888036038101906108a09190612567565b81526020019081526020015f205f8681526020019081526020015f20611d63565b9050805f0361083157611c138285611ff1565b5060055f611c296108a036899003890189612567565b815260208082019290925260409081015f9081208782529092529020805463ffffffff19169055611c5982611d63565b5f0361083157611c7b611c746108a036889003880188612567565b8490611ff1565b505050505050565b60405163237dfb4760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906346fbf68e90602401602060405180830381865afa158015611ce5573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611d0991906129c8565b610a9157604051631d77d47760e21b815260040160405180910390fd5b609881905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a250565b5f61069e825490565b60605f61088083611ffc565b60605f611d8483612055565b6040805160208082528183019092529192505f91906020820181803683375050509182525060208101929092525090565b6065546001600160a01b03163314610a915760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610d26565b6007805463ffffffff191663ffffffff83169081179091556040519081527f67d0077d22e4e06f761dd87f6c9f2310ac879c9ce17de50d381e05b72f45fbf69060200160405180910390a150565b606580546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b5f610880838361207c565b5f8181526001830160205260408120541515610880565b611a88611edd8383611408565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906120a2565b5f610880838361213c565b5f610880836001600160a01b03841661213c565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611f88573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611fac9190612844565b6001600160a01b0316336001600160a01b031614610a915760405163794821ff60e01b815260040160405180910390fd5b5f610880836001600160a01b038416612188565b5f6108808383612188565b6060815f0180548060200260200160405190810160405280929190818152602001828054801561204957602002820191905f5260205f20905b815481526020019060010190808311612035575b50505050509050919050565b5f60ff8216601f81111561069e57604051632cd44ac360e21b815260040160405180910390fd5b5f825f018281548110612091576120916127ca565b905f5260205f200154905092915050565b5f763d602d80600a3d3981f3363d3d373d3d3d363d730000008360601b60e81c175f526e5af43d82803e903d91602b57fd5bf38360781b1760205281603760095ff590506001600160a01b03811661069e5760405162461bcd60e51b815260206004820152601760248201527f455243313136373a2063726561746532206661696c65640000000000000000006044820152606401610d26565b5f81815260018301602052604081205461218157508154600181810184555f84815260208082209093018490558454848252828601909352604090209190915561069e565b505f61069e565b5f8181526001830160205260408120548015612262575f6121aa600183612a89565b85549091505f906121bd90600190612a89565b905081811461221c575f865f0182815481106121db576121db6127ca565b905f5260205f200154905080875f0184815481106121fb576121fb6127ca565b5f918252602080832090910192909255918252600188019052604090208390555b855486908061222d5761222d612a9c565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f90556001935050505061069e565b5f91505061069e565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f1916810167ffffffffffffffff811182821017156122a8576122a861226b565b604052919050565b6001600160a01b0381168114610aa4575f5ffd5b803563ffffffff811681146122d7575f5ffd5b919050565b5f604082840312156122ec575f5ffd5b6040805190810167ffffffffffffffff8111828210171561230f5761230f61226b565b6040529050808235612320816122b0565b815261232e602084016122c4565b60208201525092915050565b5f5f6060838503121561234b575f5ffd5b61235584846122dc565b946040939093013593505050565b5f60408284031215612373575f5ffd5b50919050565b5f5f5f6080848603121561238b575f5ffd5b6123958585612363565b92506040840135915060608401356123ac816122b0565b809150509250925092565b5f602082840312156123c7575f5ffd5b5035919050565b5f5f606083850312156123df575f5ffd5b6123558484612363565b5f8151808452602084019350602083015f5b8281101561243d5761242786835180516001600160a01b0316825260209081015163ffffffff16910152565b60409590950194602091909101906001016123fb565b5093949350505050565b602081525f61088060208301846123e9565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b5f6020828403121561249e575f5ffd5b813560ff81168114610880575f5ffd5b5f602082840312156124be575f5ffd5b610880826122c4565b5f604082840312156124d7575f5ffd5b6108808383612363565b5f5f5f606084860312156124f3575f5ffd5b83356124fe816122b0565b925060208401359150612513604085016122c4565b90509250925092565b5f8151808452602084019350602083015f5b8281101561243d5781516001600160a01b031686526020958601959091019060010161252e565b602081525f610880602083018461251c565b5f60408284031215612577575f5ffd5b61088083836122dc565b5f8151808452602084019350602083015f5b8281101561243d578151865260209586019590910190600101612593565b602081525f6108806020830184612581565b5f602082840312156125d3575f5ffd5b8135610880816122b0565b5f5f604083850312156125ef575f5ffd5b82356125fa816122b0565b9150612608602084016122c4565b90509250929050565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b8281101561266857603f1987860301845261265385835161251c565b94506020938401939190910190600101612637565b50929695505050505050565b5f82825180855260208501945060208160051b830101602085015f5b838110156126f057848303601f19018852815180518085526020918201918501905f5b818110156126d757835163ffffffff168352602093840193909201916001016126b3565b50506020998a0199909450929092019150600101612690565b50909695505050505050565b608081525f61270e60808301876123e9565b82810360208401528086518083526020830191506020880192505f5b8181101561274a578351151583526020938401939092019160010161272a565b505083810360408501528091505f865180835260208301935060208160051b840101602089015f5b838110156127a457601f1986840301875261278e838351612581565b6020978801979093509190910190600101612772565b505080925086810360608801526127bb8189612674565b9b9a5050505050505050505050565b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52601160045260245ffd5b8082018082111561069e5761069e6127de565b8035612810816122b0565b6001600160a01b0316825263ffffffff61282c602083016122c4565b1660208301525050565b6040810161069e8284612805565b5f60208284031215612854575f5ffd5b8151610880816122b0565b6080810161286d8286612805565b60408201939093526001600160a01b0391909116606090910152919050565b5f6020828403121561289c575f5ffd5b5051919050565b606081016128b18285612805565b8260408301529392505050565b5f602082840312156128ce575f5ffd5b815167ffffffffffffffff8111156128e4575f5ffd5b8201601f810184136128f4575f5ffd5b805167ffffffffffffffff81111561290e5761290e61226b565b8060051b61291e6020820161227f565b91825260208184018101929081019087841115612939575f5ffd5b6020850194505b8385101561295f57845180835260209586019590935090910190612940565b979650505050505050565b60a081016129788287612805565b60408201949094526001600160a01b0392909216606083015263ffffffff16608090910152919050565b81516001600160a01b0316815260208083015163ffffffff16908201526040810161069e565b5f602082840312156129d8575f5ffd5b81518015158114610880575f5ffd5b80516020808301519190811015612373575f1960209190910360031b1b16919050565b6001600160a01b0387811682528616602082015260e08101612a2f6040830187612805565b60808201949094526001600160a01b0392831660a0820152911660c0909101529392505050565b60a08101612a648287612805565b60408201949094526001600160a01b0392831660608201529116608090910152919050565b8181038181111561069e5761069e6127de565b634e487b7160e01b5f52603160045260245ffdfea26469706673582212207217cb4087765640e0f75551498e65ab7942c6943b1e70b359e434aae5b214b064736f6c634300081b0033",
}

// SlashEscrowFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use SlashEscrowFactoryMetaData.ABI instead.
var SlashEscrowFactoryABI = SlashEscrowFactoryMetaData.ABI

// SlashEscrowFactoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SlashEscrowFactoryMetaData.Bin instead.
var SlashEscrowFactoryBin = SlashEscrowFactoryMetaData.Bin

// DeploySlashEscrowFactory deploys a new Ethereum contract, binding an instance of SlashEscrowFactory to it.
func DeploySlashEscrowFactory(auth *bind.TransactOpts, backend bind.ContractBackend, _allocationManager common.Address, _strategyManager common.Address, _pauserRegistry common.Address, _slashEscrowImplementation common.Address, _version string) (common.Address, *types.Transaction, *SlashEscrowFactory, error) {
	parsed, err := SlashEscrowFactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SlashEscrowFactoryBin), backend, _allocationManager, _strategyManager, _pauserRegistry, _slashEscrowImplementation, _version)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SlashEscrowFactory{SlashEscrowFactoryCaller: SlashEscrowFactoryCaller{contract: contract}, SlashEscrowFactoryTransactor: SlashEscrowFactoryTransactor{contract: contract}, SlashEscrowFactoryFilterer: SlashEscrowFactoryFilterer{contract: contract}}, nil
}

// SlashEscrowFactory is an auto generated Go binding around an Ethereum contract.
type SlashEscrowFactory struct {
	SlashEscrowFactoryCaller     // Read-only binding to the contract
	SlashEscrowFactoryTransactor // Write-only binding to the contract
	SlashEscrowFactoryFilterer   // Log filterer for contract events
}

// SlashEscrowFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type SlashEscrowFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SlashEscrowFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SlashEscrowFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SlashEscrowFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SlashEscrowFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SlashEscrowFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SlashEscrowFactorySession struct {
	Contract     *SlashEscrowFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SlashEscrowFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SlashEscrowFactoryCallerSession struct {
	Contract *SlashEscrowFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// SlashEscrowFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SlashEscrowFactoryTransactorSession struct {
	Contract     *SlashEscrowFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// SlashEscrowFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type SlashEscrowFactoryRaw struct {
	Contract *SlashEscrowFactory // Generic contract binding to access the raw methods on
}

// SlashEscrowFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SlashEscrowFactoryCallerRaw struct {
	Contract *SlashEscrowFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// SlashEscrowFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SlashEscrowFactoryTransactorRaw struct {
	Contract *SlashEscrowFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSlashEscrowFactory creates a new instance of SlashEscrowFactory, bound to a specific deployed contract.
func NewSlashEscrowFactory(address common.Address, backend bind.ContractBackend) (*SlashEscrowFactory, error) {
	contract, err := bindSlashEscrowFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SlashEscrowFactory{SlashEscrowFactoryCaller: SlashEscrowFactoryCaller{contract: contract}, SlashEscrowFactoryTransactor: SlashEscrowFactoryTransactor{contract: contract}, SlashEscrowFactoryFilterer: SlashEscrowFactoryFilterer{contract: contract}}, nil
}

// NewSlashEscrowFactoryCaller creates a new read-only instance of SlashEscrowFactory, bound to a specific deployed contract.
func NewSlashEscrowFactoryCaller(address common.Address, caller bind.ContractCaller) (*SlashEscrowFactoryCaller, error) {
	contract, err := bindSlashEscrowFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SlashEscrowFactoryCaller{contract: contract}, nil
}

// NewSlashEscrowFactoryTransactor creates a new write-only instance of SlashEscrowFactory, bound to a specific deployed contract.
func NewSlashEscrowFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*SlashEscrowFactoryTransactor, error) {
	contract, err := bindSlashEscrowFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SlashEscrowFactoryTransactor{contract: contract}, nil
}

// NewSlashEscrowFactoryFilterer creates a new log filterer instance of SlashEscrowFactory, bound to a specific deployed contract.
func NewSlashEscrowFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*SlashEscrowFactoryFilterer, error) {
	contract, err := bindSlashEscrowFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SlashEscrowFactoryFilterer{contract: contract}, nil
}

// bindSlashEscrowFactory binds a generic wrapper to an already deployed contract.
func bindSlashEscrowFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SlashEscrowFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SlashEscrowFactory *SlashEscrowFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SlashEscrowFactory.Contract.SlashEscrowFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SlashEscrowFactory *SlashEscrowFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SlashEscrowFactory.Contract.SlashEscrowFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SlashEscrowFactory *SlashEscrowFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SlashEscrowFactory.Contract.SlashEscrowFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SlashEscrowFactory *SlashEscrowFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SlashEscrowFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SlashEscrowFactory *SlashEscrowFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SlashEscrowFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SlashEscrowFactory *SlashEscrowFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SlashEscrowFactory.Contract.contract.Transact(opts, method, params...)
}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_SlashEscrowFactory *SlashEscrowFactoryCaller) AllocationManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SlashEscrowFactory.contract.Call(opts, &out, "allocationManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_SlashEscrowFactory *SlashEscrowFactorySession) AllocationManager() (common.Address, error) {
	return _SlashEscrowFactory.Contract.AllocationManager(&_SlashEscrowFactory.CallOpts)
}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_SlashEscrowFactory *SlashEscrowFactoryCallerSession) AllocationManager() (common.Address, error) {
	return _SlashEscrowFactory.Contract.AllocationManager(&_SlashEscrowFactory.CallOpts)
}

// ComputeSlashEscrowSalt is a free data retrieval call binding the contract method 0xc2de7096.
//
// Solidity: function computeSlashEscrowSalt((address,uint32) operatorSet, uint256 slashId) pure returns(bytes32)
func (_SlashEscrowFactory *SlashEscrowFactoryCaller) ComputeSlashEscrowSalt(opts *bind.CallOpts, operatorSet OperatorSet, slashId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _SlashEscrowFactory.contract.Call(opts, &out, "computeSlashEscrowSalt", operatorSet, slashId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ComputeSlashEscrowSalt is a free data retrieval call binding the contract method 0xc2de7096.
//
// Solidity: function computeSlashEscrowSalt((address,uint32) operatorSet, uint256 slashId) pure returns(bytes32)
func (_SlashEscrowFactory *SlashEscrowFactorySession) ComputeSlashEscrowSalt(operatorSet OperatorSet, slashId *big.Int) ([32]byte, error) {
	return _SlashEscrowFactory.Contract.ComputeSlashEscrowSalt(&_SlashEscrowFactory.CallOpts, operatorSet, slashId)
}

// ComputeSlashEscrowSalt is a free data retrieval call binding the contract method 0xc2de7096.
//
// Solidity: function computeSlashEscrowSalt((address,uint32) operatorSet, uint256 slashId) pure returns(bytes32)
func (_SlashEscrowFactory *SlashEscrowFactoryCallerSession) ComputeSlashEscrowSalt(operatorSet OperatorSet, slashId *big.Int) ([32]byte, error) {
	return _SlashEscrowFactory.Contract.ComputeSlashEscrowSalt(&_SlashEscrowFactory.CallOpts, operatorSet, slashId)
}

// GetEscrowCompleteBlock is a free data retrieval call binding the contract method 0x0310f3e6.
//
// Solidity: function getEscrowCompleteBlock((address,uint32) operatorSet, uint256 slashId) view returns(uint32)
func (_SlashEscrowFactory *SlashEscrowFactoryCaller) GetEscrowCompleteBlock(opts *bind.CallOpts, operatorSet OperatorSet, slashId *big.Int) (uint32, error) {
	var out []interface{}
	err := _SlashEscrowFactory.contract.Call(opts, &out, "getEscrowCompleteBlock", operatorSet, slashId)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetEscrowCompleteBlock is a free data retrieval call binding the contract method 0x0310f3e6.
//
// Solidity: function getEscrowCompleteBlock((address,uint32) operatorSet, uint256 slashId) view returns(uint32)
func (_SlashEscrowFactory *SlashEscrowFactorySession) GetEscrowCompleteBlock(operatorSet OperatorSet, slashId *big.Int) (uint32, error) {
	return _SlashEscrowFactory.Contract.GetEscrowCompleteBlock(&_SlashEscrowFactory.CallOpts, operatorSet, slashId)
}

// GetEscrowCompleteBlock is a free data retrieval call binding the contract method 0x0310f3e6.
//
// Solidity: function getEscrowCompleteBlock((address,uint32) operatorSet, uint256 slashId) view returns(uint32)
func (_SlashEscrowFactory *SlashEscrowFactoryCallerSession) GetEscrowCompleteBlock(operatorSet OperatorSet, slashId *big.Int) (uint32, error) {
	return _SlashEscrowFactory.Contract.GetEscrowCompleteBlock(&_SlashEscrowFactory.CallOpts, operatorSet, slashId)
}

// GetEscrowStartBlock is a free data retrieval call binding the contract method 0x05a4dfbb.
//
// Solidity: function getEscrowStartBlock((address,uint32) operatorSet, uint256 slashId) view returns(uint256)
func (_SlashEscrowFactory *SlashEscrowFactoryCaller) GetEscrowStartBlock(opts *bind.CallOpts, operatorSet OperatorSet, slashId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SlashEscrowFactory.contract.Call(opts, &out, "getEscrowStartBlock", operatorSet, slashId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEscrowStartBlock is a free data retrieval call binding the contract method 0x05a4dfbb.
//
// Solidity: function getEscrowStartBlock((address,uint32) operatorSet, uint256 slashId) view returns(uint256)
func (_SlashEscrowFactory *SlashEscrowFactorySession) GetEscrowStartBlock(operatorSet OperatorSet, slashId *big.Int) (*big.Int, error) {
	return _SlashEscrowFactory.Contract.GetEscrowStartBlock(&_SlashEscrowFactory.CallOpts, operatorSet, slashId)
}

// GetEscrowStartBlock is a free data retrieval call binding the contract method 0x05a4dfbb.
//
// Solidity: function getEscrowStartBlock((address,uint32) operatorSet, uint256 slashId) view returns(uint256)
func (_SlashEscrowFactory *SlashEscrowFactoryCallerSession) GetEscrowStartBlock(operatorSet OperatorSet, slashId *big.Int) (*big.Int, error) {
	return _SlashEscrowFactory.Contract.GetEscrowStartBlock(&_SlashEscrowFactory.CallOpts, operatorSet, slashId)
}

// GetGlobalEscrowDelay is a free data retrieval call binding the contract method 0x8a65d2d2.
//
// Solidity: function getGlobalEscrowDelay() view returns(uint32)
func (_SlashEscrowFactory *SlashEscrowFactoryCaller) GetGlobalEscrowDelay(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _SlashEscrowFactory.contract.Call(opts, &out, "getGlobalEscrowDelay")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetGlobalEscrowDelay is a free data retrieval call binding the contract method 0x8a65d2d2.
//
// Solidity: function getGlobalEscrowDelay() view returns(uint32)
func (_SlashEscrowFactory *SlashEscrowFactorySession) GetGlobalEscrowDelay() (uint32, error) {
	return _SlashEscrowFactory.Contract.GetGlobalEscrowDelay(&_SlashEscrowFactory.CallOpts)
}

// GetGlobalEscrowDelay is a free data retrieval call binding the contract method 0x8a65d2d2.
//
// Solidity: function getGlobalEscrowDelay() view returns(uint32)
func (_SlashEscrowFactory *SlashEscrowFactoryCallerSession) GetGlobalEscrowDelay() (uint32, error) {
	return _SlashEscrowFactory.Contract.GetGlobalEscrowDelay(&_SlashEscrowFactory.CallOpts)
}

// GetPendingEscrows is a free data retrieval call binding the contract method 0xc50f4e48.
//
// Solidity: function getPendingEscrows() view returns((address,uint32)[] operatorSets, bool[] isRedistributing, uint256[][] slashIds, uint32[][] completeBlocks)
func (_SlashEscrowFactory *SlashEscrowFactoryCaller) GetPendingEscrows(opts *bind.CallOpts) (struct {
	OperatorSets     []OperatorSet
	IsRedistributing []bool
	SlashIds         [][]*big.Int
	CompleteBlocks   [][]uint32
}, error) {
	var out []interface{}
	err := _SlashEscrowFactory.contract.Call(opts, &out, "getPendingEscrows")

	outstruct := new(struct {
		OperatorSets     []OperatorSet
		IsRedistributing []bool
		SlashIds         [][]*big.Int
		CompleteBlocks   [][]uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OperatorSets = *abi.ConvertType(out[0], new([]OperatorSet)).(*[]OperatorSet)
	outstruct.IsRedistributing = *abi.ConvertType(out[1], new([]bool)).(*[]bool)
	outstruct.SlashIds = *abi.ConvertType(out[2], new([][]*big.Int)).(*[][]*big.Int)
	outstruct.CompleteBlocks = *abi.ConvertType(out[3], new([][]uint32)).(*[][]uint32)

	return *outstruct, err

}

// GetPendingEscrows is a free data retrieval call binding the contract method 0xc50f4e48.
//
// Solidity: function getPendingEscrows() view returns((address,uint32)[] operatorSets, bool[] isRedistributing, uint256[][] slashIds, uint32[][] completeBlocks)
func (_SlashEscrowFactory *SlashEscrowFactorySession) GetPendingEscrows() (struct {
	OperatorSets     []OperatorSet
	IsRedistributing []bool
	SlashIds         [][]*big.Int
	CompleteBlocks   [][]uint32
}, error) {
	return _SlashEscrowFactory.Contract.GetPendingEscrows(&_SlashEscrowFactory.CallOpts)
}

// GetPendingEscrows is a free data retrieval call binding the contract method 0xc50f4e48.
//
// Solidity: function getPendingEscrows() view returns((address,uint32)[] operatorSets, bool[] isRedistributing, uint256[][] slashIds, uint32[][] completeBlocks)
func (_SlashEscrowFactory *SlashEscrowFactoryCallerSession) GetPendingEscrows() (struct {
	OperatorSets     []OperatorSet
	IsRedistributing []bool
	SlashIds         [][]*big.Int
	CompleteBlocks   [][]uint32
}, error) {
	return _SlashEscrowFactory.Contract.GetPendingEscrows(&_SlashEscrowFactory.CallOpts)
}

// GetPendingOperatorSets is a free data retrieval call binding the contract method 0x3f292b08.
//
// Solidity: function getPendingOperatorSets() view returns((address,uint32)[] operatorSets)
func (_SlashEscrowFactory *SlashEscrowFactoryCaller) GetPendingOperatorSets(opts *bind.CallOpts) ([]OperatorSet, error) {
	var out []interface{}
	err := _SlashEscrowFactory.contract.Call(opts, &out, "getPendingOperatorSets")

	if err != nil {
		return *new([]OperatorSet), err
	}

	out0 := *abi.ConvertType(out[0], new([]OperatorSet)).(*[]OperatorSet)

	return out0, err

}

// GetPendingOperatorSets is a free data retrieval call binding the contract method 0x3f292b08.
//
// Solidity: function getPendingOperatorSets() view returns((address,uint32)[] operatorSets)
func (_SlashEscrowFactory *SlashEscrowFactorySession) GetPendingOperatorSets() ([]OperatorSet, error) {
	return _SlashEscrowFactory.Contract.GetPendingOperatorSets(&_SlashEscrowFactory.CallOpts)
}

// GetPendingOperatorSets is a free data retrieval call binding the contract method 0x3f292b08.
//
// Solidity: function getPendingOperatorSets() view returns((address,uint32)[] operatorSets)
func (_SlashEscrowFactory *SlashEscrowFactoryCallerSession) GetPendingOperatorSets() ([]OperatorSet, error) {
	return _SlashEscrowFactory.Contract.GetPendingOperatorSets(&_SlashEscrowFactory.CallOpts)
}

// GetPendingSlashIds is a free data retrieval call binding the contract method 0x7def1564.
//
// Solidity: function getPendingSlashIds((address,uint32) operatorSet) view returns(uint256[])
func (_SlashEscrowFactory *SlashEscrowFactoryCaller) GetPendingSlashIds(opts *bind.CallOpts, operatorSet OperatorSet) ([]*big.Int, error) {
	var out []interface{}
	err := _SlashEscrowFactory.contract.Call(opts, &out, "getPendingSlashIds", operatorSet)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetPendingSlashIds is a free data retrieval call binding the contract method 0x7def1564.
//
// Solidity: function getPendingSlashIds((address,uint32) operatorSet) view returns(uint256[])
func (_SlashEscrowFactory *SlashEscrowFactorySession) GetPendingSlashIds(operatorSet OperatorSet) ([]*big.Int, error) {
	return _SlashEscrowFactory.Contract.GetPendingSlashIds(&_SlashEscrowFactory.CallOpts, operatorSet)
}

// GetPendingSlashIds is a free data retrieval call binding the contract method 0x7def1564.
//
// Solidity: function getPendingSlashIds((address,uint32) operatorSet) view returns(uint256[])
func (_SlashEscrowFactory *SlashEscrowFactoryCallerSession) GetPendingSlashIds(operatorSet OperatorSet) ([]*big.Int, error) {
	return _SlashEscrowFactory.Contract.GetPendingSlashIds(&_SlashEscrowFactory.CallOpts, operatorSet)
}

// GetPendingStrategiesForSlashId is a free data retrieval call binding the contract method 0x7130c423.
//
// Solidity: function getPendingStrategiesForSlashId((address,uint32) operatorSet, uint256 slashId) view returns(address[] strategies)
func (_SlashEscrowFactory *SlashEscrowFactoryCaller) GetPendingStrategiesForSlashId(opts *bind.CallOpts, operatorSet OperatorSet, slashId *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _SlashEscrowFactory.contract.Call(opts, &out, "getPendingStrategiesForSlashId", operatorSet, slashId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetPendingStrategiesForSlashId is a free data retrieval call binding the contract method 0x7130c423.
//
// Solidity: function getPendingStrategiesForSlashId((address,uint32) operatorSet, uint256 slashId) view returns(address[] strategies)
func (_SlashEscrowFactory *SlashEscrowFactorySession) GetPendingStrategiesForSlashId(operatorSet OperatorSet, slashId *big.Int) ([]common.Address, error) {
	return _SlashEscrowFactory.Contract.GetPendingStrategiesForSlashId(&_SlashEscrowFactory.CallOpts, operatorSet, slashId)
}

// GetPendingStrategiesForSlashId is a free data retrieval call binding the contract method 0x7130c423.
//
// Solidity: function getPendingStrategiesForSlashId((address,uint32) operatorSet, uint256 slashId) view returns(address[] strategies)
func (_SlashEscrowFactory *SlashEscrowFactoryCallerSession) GetPendingStrategiesForSlashId(operatorSet OperatorSet, slashId *big.Int) ([]common.Address, error) {
	return _SlashEscrowFactory.Contract.GetPendingStrategiesForSlashId(&_SlashEscrowFactory.CallOpts, operatorSet, slashId)
}

// GetPendingStrategiesForSlashIds is a free data retrieval call binding the contract method 0xa56b21e4.
//
// Solidity: function getPendingStrategiesForSlashIds((address,uint32) operatorSet) view returns(address[][] strategies)
func (_SlashEscrowFactory *SlashEscrowFactoryCaller) GetPendingStrategiesForSlashIds(opts *bind.CallOpts, operatorSet OperatorSet) ([][]common.Address, error) {
	var out []interface{}
	err := _SlashEscrowFactory.contract.Call(opts, &out, "getPendingStrategiesForSlashIds", operatorSet)

	if err != nil {
		return *new([][]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([][]common.Address)).(*[][]common.Address)

	return out0, err

}

// GetPendingStrategiesForSlashIds is a free data retrieval call binding the contract method 0xa56b21e4.
//
// Solidity: function getPendingStrategiesForSlashIds((address,uint32) operatorSet) view returns(address[][] strategies)
func (_SlashEscrowFactory *SlashEscrowFactorySession) GetPendingStrategiesForSlashIds(operatorSet OperatorSet) ([][]common.Address, error) {
	return _SlashEscrowFactory.Contract.GetPendingStrategiesForSlashIds(&_SlashEscrowFactory.CallOpts, operatorSet)
}

// GetPendingStrategiesForSlashIds is a free data retrieval call binding the contract method 0xa56b21e4.
//
// Solidity: function getPendingStrategiesForSlashIds((address,uint32) operatorSet) view returns(address[][] strategies)
func (_SlashEscrowFactory *SlashEscrowFactoryCallerSession) GetPendingStrategiesForSlashIds(operatorSet OperatorSet) ([][]common.Address, error) {
	return _SlashEscrowFactory.Contract.GetPendingStrategiesForSlashIds(&_SlashEscrowFactory.CallOpts, operatorSet)
}

// GetPendingUnderlyingAmountForStrategy is a free data retrieval call binding the contract method 0xb23ff83b.
//
// Solidity: function getPendingUnderlyingAmountForStrategy((address,uint32) operatorSet, uint256 slashId, address strategy) view returns(uint256)
func (_SlashEscrowFactory *SlashEscrowFactoryCaller) GetPendingUnderlyingAmountForStrategy(opts *bind.CallOpts, operatorSet OperatorSet, slashId *big.Int, strategy common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SlashEscrowFactory.contract.Call(opts, &out, "getPendingUnderlyingAmountForStrategy", operatorSet, slashId, strategy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPendingUnderlyingAmountForStrategy is a free data retrieval call binding the contract method 0xb23ff83b.
//
// Solidity: function getPendingUnderlyingAmountForStrategy((address,uint32) operatorSet, uint256 slashId, address strategy) view returns(uint256)
func (_SlashEscrowFactory *SlashEscrowFactorySession) GetPendingUnderlyingAmountForStrategy(operatorSet OperatorSet, slashId *big.Int, strategy common.Address) (*big.Int, error) {
	return _SlashEscrowFactory.Contract.GetPendingUnderlyingAmountForStrategy(&_SlashEscrowFactory.CallOpts, operatorSet, slashId, strategy)
}

// GetPendingUnderlyingAmountForStrategy is a free data retrieval call binding the contract method 0xb23ff83b.
//
// Solidity: function getPendingUnderlyingAmountForStrategy((address,uint32) operatorSet, uint256 slashId, address strategy) view returns(uint256)
func (_SlashEscrowFactory *SlashEscrowFactoryCallerSession) GetPendingUnderlyingAmountForStrategy(operatorSet OperatorSet, slashId *big.Int, strategy common.Address) (*big.Int, error) {
	return _SlashEscrowFactory.Contract.GetPendingUnderlyingAmountForStrategy(&_SlashEscrowFactory.CallOpts, operatorSet, slashId, strategy)
}

// GetSlashEscrow is a free data retrieval call binding the contract method 0x3453b234.
//
// Solidity: function getSlashEscrow((address,uint32) operatorSet, uint256 slashId) view returns(address)
func (_SlashEscrowFactory *SlashEscrowFactoryCaller) GetSlashEscrow(opts *bind.CallOpts, operatorSet OperatorSet, slashId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SlashEscrowFactory.contract.Call(opts, &out, "getSlashEscrow", operatorSet, slashId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSlashEscrow is a free data retrieval call binding the contract method 0x3453b234.
//
// Solidity: function getSlashEscrow((address,uint32) operatorSet, uint256 slashId) view returns(address)
func (_SlashEscrowFactory *SlashEscrowFactorySession) GetSlashEscrow(operatorSet OperatorSet, slashId *big.Int) (common.Address, error) {
	return _SlashEscrowFactory.Contract.GetSlashEscrow(&_SlashEscrowFactory.CallOpts, operatorSet, slashId)
}

// GetSlashEscrow is a free data retrieval call binding the contract method 0x3453b234.
//
// Solidity: function getSlashEscrow((address,uint32) operatorSet, uint256 slashId) view returns(address)
func (_SlashEscrowFactory *SlashEscrowFactoryCallerSession) GetSlashEscrow(operatorSet OperatorSet, slashId *big.Int) (common.Address, error) {
	return _SlashEscrowFactory.Contract.GetSlashEscrow(&_SlashEscrowFactory.CallOpts, operatorSet, slashId)
}

// GetStrategyEscrowDelay is a free data retrieval call binding the contract method 0x8d5d4036.
//
// Solidity: function getStrategyEscrowDelay(address strategy) view returns(uint32)
func (_SlashEscrowFactory *SlashEscrowFactoryCaller) GetStrategyEscrowDelay(opts *bind.CallOpts, strategy common.Address) (uint32, error) {
	var out []interface{}
	err := _SlashEscrowFactory.contract.Call(opts, &out, "getStrategyEscrowDelay", strategy)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetStrategyEscrowDelay is a free data retrieval call binding the contract method 0x8d5d4036.
//
// Solidity: function getStrategyEscrowDelay(address strategy) view returns(uint32)
func (_SlashEscrowFactory *SlashEscrowFactorySession) GetStrategyEscrowDelay(strategy common.Address) (uint32, error) {
	return _SlashEscrowFactory.Contract.GetStrategyEscrowDelay(&_SlashEscrowFactory.CallOpts, strategy)
}

// GetStrategyEscrowDelay is a free data retrieval call binding the contract method 0x8d5d4036.
//
// Solidity: function getStrategyEscrowDelay(address strategy) view returns(uint32)
func (_SlashEscrowFactory *SlashEscrowFactoryCallerSession) GetStrategyEscrowDelay(strategy common.Address) (uint32, error) {
	return _SlashEscrowFactory.Contract.GetStrategyEscrowDelay(&_SlashEscrowFactory.CallOpts, strategy)
}

// GetTotalPendingOperatorSets is a free data retrieval call binding the contract method 0x78cb9600.
//
// Solidity: function getTotalPendingOperatorSets() view returns(uint256)
func (_SlashEscrowFactory *SlashEscrowFactoryCaller) GetTotalPendingOperatorSets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SlashEscrowFactory.contract.Call(opts, &out, "getTotalPendingOperatorSets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalPendingOperatorSets is a free data retrieval call binding the contract method 0x78cb9600.
//
// Solidity: function getTotalPendingOperatorSets() view returns(uint256)
func (_SlashEscrowFactory *SlashEscrowFactorySession) GetTotalPendingOperatorSets() (*big.Int, error) {
	return _SlashEscrowFactory.Contract.GetTotalPendingOperatorSets(&_SlashEscrowFactory.CallOpts)
}

// GetTotalPendingOperatorSets is a free data retrieval call binding the contract method 0x78cb9600.
//
// Solidity: function getTotalPendingOperatorSets() view returns(uint256)
func (_SlashEscrowFactory *SlashEscrowFactoryCallerSession) GetTotalPendingOperatorSets() (*big.Int, error) {
	return _SlashEscrowFactory.Contract.GetTotalPendingOperatorSets(&_SlashEscrowFactory.CallOpts)
}

// GetTotalPendingSlashIds is a free data retrieval call binding the contract method 0x6729b5db.
//
// Solidity: function getTotalPendingSlashIds((address,uint32) operatorSet) view returns(uint256)
func (_SlashEscrowFactory *SlashEscrowFactoryCaller) GetTotalPendingSlashIds(opts *bind.CallOpts, operatorSet OperatorSet) (*big.Int, error) {
	var out []interface{}
	err := _SlashEscrowFactory.contract.Call(opts, &out, "getTotalPendingSlashIds", operatorSet)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalPendingSlashIds is a free data retrieval call binding the contract method 0x6729b5db.
//
// Solidity: function getTotalPendingSlashIds((address,uint32) operatorSet) view returns(uint256)
func (_SlashEscrowFactory *SlashEscrowFactorySession) GetTotalPendingSlashIds(operatorSet OperatorSet) (*big.Int, error) {
	return _SlashEscrowFactory.Contract.GetTotalPendingSlashIds(&_SlashEscrowFactory.CallOpts, operatorSet)
}

// GetTotalPendingSlashIds is a free data retrieval call binding the contract method 0x6729b5db.
//
// Solidity: function getTotalPendingSlashIds((address,uint32) operatorSet) view returns(uint256)
func (_SlashEscrowFactory *SlashEscrowFactoryCallerSession) GetTotalPendingSlashIds(operatorSet OperatorSet) (*big.Int, error) {
	return _SlashEscrowFactory.Contract.GetTotalPendingSlashIds(&_SlashEscrowFactory.CallOpts, operatorSet)
}

// GetTotalPendingStrategiesForSlashId is a free data retrieval call binding the contract method 0x277a9f0e.
//
// Solidity: function getTotalPendingStrategiesForSlashId((address,uint32) operatorSet, uint256 slashId) view returns(uint256)
func (_SlashEscrowFactory *SlashEscrowFactoryCaller) GetTotalPendingStrategiesForSlashId(opts *bind.CallOpts, operatorSet OperatorSet, slashId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SlashEscrowFactory.contract.Call(opts, &out, "getTotalPendingStrategiesForSlashId", operatorSet, slashId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalPendingStrategiesForSlashId is a free data retrieval call binding the contract method 0x277a9f0e.
//
// Solidity: function getTotalPendingStrategiesForSlashId((address,uint32) operatorSet, uint256 slashId) view returns(uint256)
func (_SlashEscrowFactory *SlashEscrowFactorySession) GetTotalPendingStrategiesForSlashId(operatorSet OperatorSet, slashId *big.Int) (*big.Int, error) {
	return _SlashEscrowFactory.Contract.GetTotalPendingStrategiesForSlashId(&_SlashEscrowFactory.CallOpts, operatorSet, slashId)
}

// GetTotalPendingStrategiesForSlashId is a free data retrieval call binding the contract method 0x277a9f0e.
//
// Solidity: function getTotalPendingStrategiesForSlashId((address,uint32) operatorSet, uint256 slashId) view returns(uint256)
func (_SlashEscrowFactory *SlashEscrowFactoryCallerSession) GetTotalPendingStrategiesForSlashId(operatorSet OperatorSet, slashId *big.Int) (*big.Int, error) {
	return _SlashEscrowFactory.Contract.GetTotalPendingStrategiesForSlashId(&_SlashEscrowFactory.CallOpts, operatorSet, slashId)
}

// IsDeployedSlashEscrow is a free data retrieval call binding the contract method 0x19f3db26.
//
// Solidity: function isDeployedSlashEscrow((address,uint32) operatorSet, uint256 slashId) view returns(bool)
func (_SlashEscrowFactory *SlashEscrowFactoryCaller) IsDeployedSlashEscrow(opts *bind.CallOpts, operatorSet OperatorSet, slashId *big.Int) (bool, error) {
	var out []interface{}
	err := _SlashEscrowFactory.contract.Call(opts, &out, "isDeployedSlashEscrow", operatorSet, slashId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDeployedSlashEscrow is a free data retrieval call binding the contract method 0x19f3db26.
//
// Solidity: function isDeployedSlashEscrow((address,uint32) operatorSet, uint256 slashId) view returns(bool)
func (_SlashEscrowFactory *SlashEscrowFactorySession) IsDeployedSlashEscrow(operatorSet OperatorSet, slashId *big.Int) (bool, error) {
	return _SlashEscrowFactory.Contract.IsDeployedSlashEscrow(&_SlashEscrowFactory.CallOpts, operatorSet, slashId)
}

// IsDeployedSlashEscrow is a free data retrieval call binding the contract method 0x19f3db26.
//
// Solidity: function isDeployedSlashEscrow((address,uint32) operatorSet, uint256 slashId) view returns(bool)
func (_SlashEscrowFactory *SlashEscrowFactoryCallerSession) IsDeployedSlashEscrow(operatorSet OperatorSet, slashId *big.Int) (bool, error) {
	return _SlashEscrowFactory.Contract.IsDeployedSlashEscrow(&_SlashEscrowFactory.CallOpts, operatorSet, slashId)
}

// IsDeployedSlashEscrow0 is a free data retrieval call binding the contract method 0xc8b5330c.
//
// Solidity: function isDeployedSlashEscrow(address slashEscrow) view returns(bool)
func (_SlashEscrowFactory *SlashEscrowFactoryCaller) IsDeployedSlashEscrow0(opts *bind.CallOpts, slashEscrow common.Address) (bool, error) {
	var out []interface{}
	err := _SlashEscrowFactory.contract.Call(opts, &out, "isDeployedSlashEscrow0", slashEscrow)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDeployedSlashEscrow0 is a free data retrieval call binding the contract method 0xc8b5330c.
//
// Solidity: function isDeployedSlashEscrow(address slashEscrow) view returns(bool)
func (_SlashEscrowFactory *SlashEscrowFactorySession) IsDeployedSlashEscrow0(slashEscrow common.Address) (bool, error) {
	return _SlashEscrowFactory.Contract.IsDeployedSlashEscrow0(&_SlashEscrowFactory.CallOpts, slashEscrow)
}

// IsDeployedSlashEscrow0 is a free data retrieval call binding the contract method 0xc8b5330c.
//
// Solidity: function isDeployedSlashEscrow(address slashEscrow) view returns(bool)
func (_SlashEscrowFactory *SlashEscrowFactoryCallerSession) IsDeployedSlashEscrow0(slashEscrow common.Address) (bool, error) {
	return _SlashEscrowFactory.Contract.IsDeployedSlashEscrow0(&_SlashEscrowFactory.CallOpts, slashEscrow)
}

// IsEscrowPaused is a free data retrieval call binding the contract method 0x78748459.
//
// Solidity: function isEscrowPaused((address,uint32) operatorSet, uint256 slashId) view returns(bool)
func (_SlashEscrowFactory *SlashEscrowFactoryCaller) IsEscrowPaused(opts *bind.CallOpts, operatorSet OperatorSet, slashId *big.Int) (bool, error) {
	var out []interface{}
	err := _SlashEscrowFactory.contract.Call(opts, &out, "isEscrowPaused", operatorSet, slashId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEscrowPaused is a free data retrieval call binding the contract method 0x78748459.
//
// Solidity: function isEscrowPaused((address,uint32) operatorSet, uint256 slashId) view returns(bool)
func (_SlashEscrowFactory *SlashEscrowFactorySession) IsEscrowPaused(operatorSet OperatorSet, slashId *big.Int) (bool, error) {
	return _SlashEscrowFactory.Contract.IsEscrowPaused(&_SlashEscrowFactory.CallOpts, operatorSet, slashId)
}

// IsEscrowPaused is a free data retrieval call binding the contract method 0x78748459.
//
// Solidity: function isEscrowPaused((address,uint32) operatorSet, uint256 slashId) view returns(bool)
func (_SlashEscrowFactory *SlashEscrowFactoryCallerSession) IsEscrowPaused(operatorSet OperatorSet, slashId *big.Int) (bool, error) {
	return _SlashEscrowFactory.Contract.IsEscrowPaused(&_SlashEscrowFactory.CallOpts, operatorSet, slashId)
}

// IsPendingOperatorSet is a free data retrieval call binding the contract method 0x71e166e7.
//
// Solidity: function isPendingOperatorSet((address,uint32) operatorSet) view returns(bool)
func (_SlashEscrowFactory *SlashEscrowFactoryCaller) IsPendingOperatorSet(opts *bind.CallOpts, operatorSet OperatorSet) (bool, error) {
	var out []interface{}
	err := _SlashEscrowFactory.contract.Call(opts, &out, "isPendingOperatorSet", operatorSet)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPendingOperatorSet is a free data retrieval call binding the contract method 0x71e166e7.
//
// Solidity: function isPendingOperatorSet((address,uint32) operatorSet) view returns(bool)
func (_SlashEscrowFactory *SlashEscrowFactorySession) IsPendingOperatorSet(operatorSet OperatorSet) (bool, error) {
	return _SlashEscrowFactory.Contract.IsPendingOperatorSet(&_SlashEscrowFactory.CallOpts, operatorSet)
}

// IsPendingOperatorSet is a free data retrieval call binding the contract method 0x71e166e7.
//
// Solidity: function isPendingOperatorSet((address,uint32) operatorSet) view returns(bool)
func (_SlashEscrowFactory *SlashEscrowFactoryCallerSession) IsPendingOperatorSet(operatorSet OperatorSet) (bool, error) {
	return _SlashEscrowFactory.Contract.IsPendingOperatorSet(&_SlashEscrowFactory.CallOpts, operatorSet)
}

// IsPendingSlashId is a free data retrieval call binding the contract method 0x87420b07.
//
// Solidity: function isPendingSlashId((address,uint32) operatorSet, uint256 slashId) view returns(bool)
func (_SlashEscrowFactory *SlashEscrowFactoryCaller) IsPendingSlashId(opts *bind.CallOpts, operatorSet OperatorSet, slashId *big.Int) (bool, error) {
	var out []interface{}
	err := _SlashEscrowFactory.contract.Call(opts, &out, "isPendingSlashId", operatorSet, slashId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPendingSlashId is a free data retrieval call binding the contract method 0x87420b07.
//
// Solidity: function isPendingSlashId((address,uint32) operatorSet, uint256 slashId) view returns(bool)
func (_SlashEscrowFactory *SlashEscrowFactorySession) IsPendingSlashId(operatorSet OperatorSet, slashId *big.Int) (bool, error) {
	return _SlashEscrowFactory.Contract.IsPendingSlashId(&_SlashEscrowFactory.CallOpts, operatorSet, slashId)
}

// IsPendingSlashId is a free data retrieval call binding the contract method 0x87420b07.
//
// Solidity: function isPendingSlashId((address,uint32) operatorSet, uint256 slashId) view returns(bool)
func (_SlashEscrowFactory *SlashEscrowFactoryCallerSession) IsPendingSlashId(operatorSet OperatorSet, slashId *big.Int) (bool, error) {
	return _SlashEscrowFactory.Contract.IsPendingSlashId(&_SlashEscrowFactory.CallOpts, operatorSet, slashId)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SlashEscrowFactory *SlashEscrowFactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SlashEscrowFactory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SlashEscrowFactory *SlashEscrowFactorySession) Owner() (common.Address, error) {
	return _SlashEscrowFactory.Contract.Owner(&_SlashEscrowFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SlashEscrowFactory *SlashEscrowFactoryCallerSession) Owner() (common.Address, error) {
	return _SlashEscrowFactory.Contract.Owner(&_SlashEscrowFactory.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_SlashEscrowFactory *SlashEscrowFactoryCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _SlashEscrowFactory.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_SlashEscrowFactory *SlashEscrowFactorySession) Paused(index uint8) (bool, error) {
	return _SlashEscrowFactory.Contract.Paused(&_SlashEscrowFactory.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_SlashEscrowFactory *SlashEscrowFactoryCallerSession) Paused(index uint8) (bool, error) {
	return _SlashEscrowFactory.Contract.Paused(&_SlashEscrowFactory.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_SlashEscrowFactory *SlashEscrowFactoryCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SlashEscrowFactory.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_SlashEscrowFactory *SlashEscrowFactorySession) Paused0() (*big.Int, error) {
	return _SlashEscrowFactory.Contract.Paused0(&_SlashEscrowFactory.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_SlashEscrowFactory *SlashEscrowFactoryCallerSession) Paused0() (*big.Int, error) {
	return _SlashEscrowFactory.Contract.Paused0(&_SlashEscrowFactory.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_SlashEscrowFactory *SlashEscrowFactoryCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SlashEscrowFactory.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_SlashEscrowFactory *SlashEscrowFactorySession) PauserRegistry() (common.Address, error) {
	return _SlashEscrowFactory.Contract.PauserRegistry(&_SlashEscrowFactory.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_SlashEscrowFactory *SlashEscrowFactoryCallerSession) PauserRegistry() (common.Address, error) {
	return _SlashEscrowFactory.Contract.PauserRegistry(&_SlashEscrowFactory.CallOpts)
}

// SlashEscrowImplementation is a free data retrieval call binding the contract method 0xa3c65641.
//
// Solidity: function slashEscrowImplementation() view returns(address)
func (_SlashEscrowFactory *SlashEscrowFactoryCaller) SlashEscrowImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SlashEscrowFactory.contract.Call(opts, &out, "slashEscrowImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SlashEscrowImplementation is a free data retrieval call binding the contract method 0xa3c65641.
//
// Solidity: function slashEscrowImplementation() view returns(address)
func (_SlashEscrowFactory *SlashEscrowFactorySession) SlashEscrowImplementation() (common.Address, error) {
	return _SlashEscrowFactory.Contract.SlashEscrowImplementation(&_SlashEscrowFactory.CallOpts)
}

// SlashEscrowImplementation is a free data retrieval call binding the contract method 0xa3c65641.
//
// Solidity: function slashEscrowImplementation() view returns(address)
func (_SlashEscrowFactory *SlashEscrowFactoryCallerSession) SlashEscrowImplementation() (common.Address, error) {
	return _SlashEscrowFactory.Contract.SlashEscrowImplementation(&_SlashEscrowFactory.CallOpts)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_SlashEscrowFactory *SlashEscrowFactoryCaller) StrategyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SlashEscrowFactory.contract.Call(opts, &out, "strategyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_SlashEscrowFactory *SlashEscrowFactorySession) StrategyManager() (common.Address, error) {
	return _SlashEscrowFactory.Contract.StrategyManager(&_SlashEscrowFactory.CallOpts)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_SlashEscrowFactory *SlashEscrowFactoryCallerSession) StrategyManager() (common.Address, error) {
	return _SlashEscrowFactory.Contract.StrategyManager(&_SlashEscrowFactory.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_SlashEscrowFactory *SlashEscrowFactoryCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SlashEscrowFactory.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_SlashEscrowFactory *SlashEscrowFactorySession) Version() (string, error) {
	return _SlashEscrowFactory.Contract.Version(&_SlashEscrowFactory.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_SlashEscrowFactory *SlashEscrowFactoryCallerSession) Version() (string, error) {
	return _SlashEscrowFactory.Contract.Version(&_SlashEscrowFactory.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x6c5ac81e.
//
// Solidity: function initialize(address initialOwner, uint256 initialPausedStatus, uint32 initialGlobalDelayBlocks) returns()
func (_SlashEscrowFactory *SlashEscrowFactoryTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, initialPausedStatus *big.Int, initialGlobalDelayBlocks uint32) (*types.Transaction, error) {
	return _SlashEscrowFactory.contract.Transact(opts, "initialize", initialOwner, initialPausedStatus, initialGlobalDelayBlocks)
}

// Initialize is a paid mutator transaction binding the contract method 0x6c5ac81e.
//
// Solidity: function initialize(address initialOwner, uint256 initialPausedStatus, uint32 initialGlobalDelayBlocks) returns()
func (_SlashEscrowFactory *SlashEscrowFactorySession) Initialize(initialOwner common.Address, initialPausedStatus *big.Int, initialGlobalDelayBlocks uint32) (*types.Transaction, error) {
	return _SlashEscrowFactory.Contract.Initialize(&_SlashEscrowFactory.TransactOpts, initialOwner, initialPausedStatus, initialGlobalDelayBlocks)
}

// Initialize is a paid mutator transaction binding the contract method 0x6c5ac81e.
//
// Solidity: function initialize(address initialOwner, uint256 initialPausedStatus, uint32 initialGlobalDelayBlocks) returns()
func (_SlashEscrowFactory *SlashEscrowFactoryTransactorSession) Initialize(initialOwner common.Address, initialPausedStatus *big.Int, initialGlobalDelayBlocks uint32) (*types.Transaction, error) {
	return _SlashEscrowFactory.Contract.Initialize(&_SlashEscrowFactory.TransactOpts, initialOwner, initialPausedStatus, initialGlobalDelayBlocks)
}

// InitiateSlashEscrow is a paid mutator transaction binding the contract method 0x7a967611.
//
// Solidity: function initiateSlashEscrow((address,uint32) operatorSet, uint256 slashId, address strategy) returns()
func (_SlashEscrowFactory *SlashEscrowFactoryTransactor) InitiateSlashEscrow(opts *bind.TransactOpts, operatorSet OperatorSet, slashId *big.Int, strategy common.Address) (*types.Transaction, error) {
	return _SlashEscrowFactory.contract.Transact(opts, "initiateSlashEscrow", operatorSet, slashId, strategy)
}

// InitiateSlashEscrow is a paid mutator transaction binding the contract method 0x7a967611.
//
// Solidity: function initiateSlashEscrow((address,uint32) operatorSet, uint256 slashId, address strategy) returns()
func (_SlashEscrowFactory *SlashEscrowFactorySession) InitiateSlashEscrow(operatorSet OperatorSet, slashId *big.Int, strategy common.Address) (*types.Transaction, error) {
	return _SlashEscrowFactory.Contract.InitiateSlashEscrow(&_SlashEscrowFactory.TransactOpts, operatorSet, slashId, strategy)
}

// InitiateSlashEscrow is a paid mutator transaction binding the contract method 0x7a967611.
//
// Solidity: function initiateSlashEscrow((address,uint32) operatorSet, uint256 slashId, address strategy) returns()
func (_SlashEscrowFactory *SlashEscrowFactoryTransactorSession) InitiateSlashEscrow(operatorSet OperatorSet, slashId *big.Int, strategy common.Address) (*types.Transaction, error) {
	return _SlashEscrowFactory.Contract.InitiateSlashEscrow(&_SlashEscrowFactory.TransactOpts, operatorSet, slashId, strategy)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_SlashEscrowFactory *SlashEscrowFactoryTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _SlashEscrowFactory.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_SlashEscrowFactory *SlashEscrowFactorySession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _SlashEscrowFactory.Contract.Pause(&_SlashEscrowFactory.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_SlashEscrowFactory *SlashEscrowFactoryTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _SlashEscrowFactory.Contract.Pause(&_SlashEscrowFactory.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_SlashEscrowFactory *SlashEscrowFactoryTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SlashEscrowFactory.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_SlashEscrowFactory *SlashEscrowFactorySession) PauseAll() (*types.Transaction, error) {
	return _SlashEscrowFactory.Contract.PauseAll(&_SlashEscrowFactory.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_SlashEscrowFactory *SlashEscrowFactoryTransactorSession) PauseAll() (*types.Transaction, error) {
	return _SlashEscrowFactory.Contract.PauseAll(&_SlashEscrowFactory.TransactOpts)
}

// PauseEscrow is a paid mutator transaction binding the contract method 0xe7ed076b.
//
// Solidity: function pauseEscrow((address,uint32) operatorSet, uint256 slashId) returns()
func (_SlashEscrowFactory *SlashEscrowFactoryTransactor) PauseEscrow(opts *bind.TransactOpts, operatorSet OperatorSet, slashId *big.Int) (*types.Transaction, error) {
	return _SlashEscrowFactory.contract.Transact(opts, "pauseEscrow", operatorSet, slashId)
}

// PauseEscrow is a paid mutator transaction binding the contract method 0xe7ed076b.
//
// Solidity: function pauseEscrow((address,uint32) operatorSet, uint256 slashId) returns()
func (_SlashEscrowFactory *SlashEscrowFactorySession) PauseEscrow(operatorSet OperatorSet, slashId *big.Int) (*types.Transaction, error) {
	return _SlashEscrowFactory.Contract.PauseEscrow(&_SlashEscrowFactory.TransactOpts, operatorSet, slashId)
}

// PauseEscrow is a paid mutator transaction binding the contract method 0xe7ed076b.
//
// Solidity: function pauseEscrow((address,uint32) operatorSet, uint256 slashId) returns()
func (_SlashEscrowFactory *SlashEscrowFactoryTransactorSession) PauseEscrow(operatorSet OperatorSet, slashId *big.Int) (*types.Transaction, error) {
	return _SlashEscrowFactory.Contract.PauseEscrow(&_SlashEscrowFactory.TransactOpts, operatorSet, slashId)
}

// ReleaseSlashEscrow is a paid mutator transaction binding the contract method 0x5ffa5a81.
//
// Solidity: function releaseSlashEscrow((address,uint32) operatorSet, uint256 slashId) returns()
func (_SlashEscrowFactory *SlashEscrowFactoryTransactor) ReleaseSlashEscrow(opts *bind.TransactOpts, operatorSet OperatorSet, slashId *big.Int) (*types.Transaction, error) {
	return _SlashEscrowFactory.contract.Transact(opts, "releaseSlashEscrow", operatorSet, slashId)
}

// ReleaseSlashEscrow is a paid mutator transaction binding the contract method 0x5ffa5a81.
//
// Solidity: function releaseSlashEscrow((address,uint32) operatorSet, uint256 slashId) returns()
func (_SlashEscrowFactory *SlashEscrowFactorySession) ReleaseSlashEscrow(operatorSet OperatorSet, slashId *big.Int) (*types.Transaction, error) {
	return _SlashEscrowFactory.Contract.ReleaseSlashEscrow(&_SlashEscrowFactory.TransactOpts, operatorSet, slashId)
}

// ReleaseSlashEscrow is a paid mutator transaction binding the contract method 0x5ffa5a81.
//
// Solidity: function releaseSlashEscrow((address,uint32) operatorSet, uint256 slashId) returns()
func (_SlashEscrowFactory *SlashEscrowFactoryTransactorSession) ReleaseSlashEscrow(operatorSet OperatorSet, slashId *big.Int) (*types.Transaction, error) {
	return _SlashEscrowFactory.Contract.ReleaseSlashEscrow(&_SlashEscrowFactory.TransactOpts, operatorSet, slashId)
}

// ReleaseSlashEscrowByStrategy is a paid mutator transaction binding the contract method 0x0e475b17.
//
// Solidity: function releaseSlashEscrowByStrategy((address,uint32) operatorSet, uint256 slashId, address strategy) returns()
func (_SlashEscrowFactory *SlashEscrowFactoryTransactor) ReleaseSlashEscrowByStrategy(opts *bind.TransactOpts, operatorSet OperatorSet, slashId *big.Int, strategy common.Address) (*types.Transaction, error) {
	return _SlashEscrowFactory.contract.Transact(opts, "releaseSlashEscrowByStrategy", operatorSet, slashId, strategy)
}

// ReleaseSlashEscrowByStrategy is a paid mutator transaction binding the contract method 0x0e475b17.
//
// Solidity: function releaseSlashEscrowByStrategy((address,uint32) operatorSet, uint256 slashId, address strategy) returns()
func (_SlashEscrowFactory *SlashEscrowFactorySession) ReleaseSlashEscrowByStrategy(operatorSet OperatorSet, slashId *big.Int, strategy common.Address) (*types.Transaction, error) {
	return _SlashEscrowFactory.Contract.ReleaseSlashEscrowByStrategy(&_SlashEscrowFactory.TransactOpts, operatorSet, slashId, strategy)
}

// ReleaseSlashEscrowByStrategy is a paid mutator transaction binding the contract method 0x0e475b17.
//
// Solidity: function releaseSlashEscrowByStrategy((address,uint32) operatorSet, uint256 slashId, address strategy) returns()
func (_SlashEscrowFactory *SlashEscrowFactoryTransactorSession) ReleaseSlashEscrowByStrategy(operatorSet OperatorSet, slashId *big.Int, strategy common.Address) (*types.Transaction, error) {
	return _SlashEscrowFactory.Contract.ReleaseSlashEscrowByStrategy(&_SlashEscrowFactory.TransactOpts, operatorSet, slashId, strategy)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SlashEscrowFactory *SlashEscrowFactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SlashEscrowFactory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SlashEscrowFactory *SlashEscrowFactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _SlashEscrowFactory.Contract.RenounceOwnership(&_SlashEscrowFactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SlashEscrowFactory *SlashEscrowFactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SlashEscrowFactory.Contract.RenounceOwnership(&_SlashEscrowFactory.TransactOpts)
}

// SetGlobalEscrowDelay is a paid mutator transaction binding the contract method 0x5e0a64c5.
//
// Solidity: function setGlobalEscrowDelay(uint32 delay) returns()
func (_SlashEscrowFactory *SlashEscrowFactoryTransactor) SetGlobalEscrowDelay(opts *bind.TransactOpts, delay uint32) (*types.Transaction, error) {
	return _SlashEscrowFactory.contract.Transact(opts, "setGlobalEscrowDelay", delay)
}

// SetGlobalEscrowDelay is a paid mutator transaction binding the contract method 0x5e0a64c5.
//
// Solidity: function setGlobalEscrowDelay(uint32 delay) returns()
func (_SlashEscrowFactory *SlashEscrowFactorySession) SetGlobalEscrowDelay(delay uint32) (*types.Transaction, error) {
	return _SlashEscrowFactory.Contract.SetGlobalEscrowDelay(&_SlashEscrowFactory.TransactOpts, delay)
}

// SetGlobalEscrowDelay is a paid mutator transaction binding the contract method 0x5e0a64c5.
//
// Solidity: function setGlobalEscrowDelay(uint32 delay) returns()
func (_SlashEscrowFactory *SlashEscrowFactoryTransactorSession) SetGlobalEscrowDelay(delay uint32) (*types.Transaction, error) {
	return _SlashEscrowFactory.Contract.SetGlobalEscrowDelay(&_SlashEscrowFactory.TransactOpts, delay)
}

// SetStrategyEscrowDelay is a paid mutator transaction binding the contract method 0x8fc46be5.
//
// Solidity: function setStrategyEscrowDelay(address strategy, uint32 delay) returns()
func (_SlashEscrowFactory *SlashEscrowFactoryTransactor) SetStrategyEscrowDelay(opts *bind.TransactOpts, strategy common.Address, delay uint32) (*types.Transaction, error) {
	return _SlashEscrowFactory.contract.Transact(opts, "setStrategyEscrowDelay", strategy, delay)
}

// SetStrategyEscrowDelay is a paid mutator transaction binding the contract method 0x8fc46be5.
//
// Solidity: function setStrategyEscrowDelay(address strategy, uint32 delay) returns()
func (_SlashEscrowFactory *SlashEscrowFactorySession) SetStrategyEscrowDelay(strategy common.Address, delay uint32) (*types.Transaction, error) {
	return _SlashEscrowFactory.Contract.SetStrategyEscrowDelay(&_SlashEscrowFactory.TransactOpts, strategy, delay)
}

// SetStrategyEscrowDelay is a paid mutator transaction binding the contract method 0x8fc46be5.
//
// Solidity: function setStrategyEscrowDelay(address strategy, uint32 delay) returns()
func (_SlashEscrowFactory *SlashEscrowFactoryTransactorSession) SetStrategyEscrowDelay(strategy common.Address, delay uint32) (*types.Transaction, error) {
	return _SlashEscrowFactory.Contract.SetStrategyEscrowDelay(&_SlashEscrowFactory.TransactOpts, strategy, delay)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SlashEscrowFactory *SlashEscrowFactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SlashEscrowFactory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SlashEscrowFactory *SlashEscrowFactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SlashEscrowFactory.Contract.TransferOwnership(&_SlashEscrowFactory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SlashEscrowFactory *SlashEscrowFactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SlashEscrowFactory.Contract.TransferOwnership(&_SlashEscrowFactory.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_SlashEscrowFactory *SlashEscrowFactoryTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _SlashEscrowFactory.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_SlashEscrowFactory *SlashEscrowFactorySession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _SlashEscrowFactory.Contract.Unpause(&_SlashEscrowFactory.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_SlashEscrowFactory *SlashEscrowFactoryTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _SlashEscrowFactory.Contract.Unpause(&_SlashEscrowFactory.TransactOpts, newPausedStatus)
}

// UnpauseEscrow is a paid mutator transaction binding the contract method 0x9b122356.
//
// Solidity: function unpauseEscrow((address,uint32) operatorSet, uint256 slashId) returns()
func (_SlashEscrowFactory *SlashEscrowFactoryTransactor) UnpauseEscrow(opts *bind.TransactOpts, operatorSet OperatorSet, slashId *big.Int) (*types.Transaction, error) {
	return _SlashEscrowFactory.contract.Transact(opts, "unpauseEscrow", operatorSet, slashId)
}

// UnpauseEscrow is a paid mutator transaction binding the contract method 0x9b122356.
//
// Solidity: function unpauseEscrow((address,uint32) operatorSet, uint256 slashId) returns()
func (_SlashEscrowFactory *SlashEscrowFactorySession) UnpauseEscrow(operatorSet OperatorSet, slashId *big.Int) (*types.Transaction, error) {
	return _SlashEscrowFactory.Contract.UnpauseEscrow(&_SlashEscrowFactory.TransactOpts, operatorSet, slashId)
}

// UnpauseEscrow is a paid mutator transaction binding the contract method 0x9b122356.
//
// Solidity: function unpauseEscrow((address,uint32) operatorSet, uint256 slashId) returns()
func (_SlashEscrowFactory *SlashEscrowFactoryTransactorSession) UnpauseEscrow(operatorSet OperatorSet, slashId *big.Int) (*types.Transaction, error) {
	return _SlashEscrowFactory.Contract.UnpauseEscrow(&_SlashEscrowFactory.TransactOpts, operatorSet, slashId)
}

// SlashEscrowFactoryEscrowCompleteIterator is returned from FilterEscrowComplete and is used to iterate over the raw logs and unpacked data for EscrowComplete events raised by the SlashEscrowFactory contract.
type SlashEscrowFactoryEscrowCompleteIterator struct {
	Event *SlashEscrowFactoryEscrowComplete // Event containing the contract specifics and raw log

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
func (it *SlashEscrowFactoryEscrowCompleteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlashEscrowFactoryEscrowComplete)
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
		it.Event = new(SlashEscrowFactoryEscrowComplete)
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
func (it *SlashEscrowFactoryEscrowCompleteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlashEscrowFactoryEscrowCompleteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlashEscrowFactoryEscrowComplete represents a EscrowComplete event raised by the SlashEscrowFactory contract.
type SlashEscrowFactoryEscrowComplete struct {
	OperatorSet OperatorSet
	SlashId     *big.Int
	Strategy    common.Address
	Recipient   common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterEscrowComplete is a free log retrieval operation binding the contract event 0x32be306ad5a833e756b7cb9724d5312afe0feda6163bfc2dd98ee713346a9abc.
//
// Solidity: event EscrowComplete((address,uint32) operatorSet, uint256 slashId, address strategy, address recipient)
func (_SlashEscrowFactory *SlashEscrowFactoryFilterer) FilterEscrowComplete(opts *bind.FilterOpts) (*SlashEscrowFactoryEscrowCompleteIterator, error) {

	logs, sub, err := _SlashEscrowFactory.contract.FilterLogs(opts, "EscrowComplete")
	if err != nil {
		return nil, err
	}
	return &SlashEscrowFactoryEscrowCompleteIterator{contract: _SlashEscrowFactory.contract, event: "EscrowComplete", logs: logs, sub: sub}, nil
}

// WatchEscrowComplete is a free log subscription operation binding the contract event 0x32be306ad5a833e756b7cb9724d5312afe0feda6163bfc2dd98ee713346a9abc.
//
// Solidity: event EscrowComplete((address,uint32) operatorSet, uint256 slashId, address strategy, address recipient)
func (_SlashEscrowFactory *SlashEscrowFactoryFilterer) WatchEscrowComplete(opts *bind.WatchOpts, sink chan<- *SlashEscrowFactoryEscrowComplete) (event.Subscription, error) {

	logs, sub, err := _SlashEscrowFactory.contract.WatchLogs(opts, "EscrowComplete")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlashEscrowFactoryEscrowComplete)
				if err := _SlashEscrowFactory.contract.UnpackLog(event, "EscrowComplete", log); err != nil {
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

// ParseEscrowComplete is a log parse operation binding the contract event 0x32be306ad5a833e756b7cb9724d5312afe0feda6163bfc2dd98ee713346a9abc.
//
// Solidity: event EscrowComplete((address,uint32) operatorSet, uint256 slashId, address strategy, address recipient)
func (_SlashEscrowFactory *SlashEscrowFactoryFilterer) ParseEscrowComplete(log types.Log) (*SlashEscrowFactoryEscrowComplete, error) {
	event := new(SlashEscrowFactoryEscrowComplete)
	if err := _SlashEscrowFactory.contract.UnpackLog(event, "EscrowComplete", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlashEscrowFactoryEscrowPausedIterator is returned from FilterEscrowPaused and is used to iterate over the raw logs and unpacked data for EscrowPaused events raised by the SlashEscrowFactory contract.
type SlashEscrowFactoryEscrowPausedIterator struct {
	Event *SlashEscrowFactoryEscrowPaused // Event containing the contract specifics and raw log

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
func (it *SlashEscrowFactoryEscrowPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlashEscrowFactoryEscrowPaused)
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
		it.Event = new(SlashEscrowFactoryEscrowPaused)
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
func (it *SlashEscrowFactoryEscrowPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlashEscrowFactoryEscrowPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlashEscrowFactoryEscrowPaused represents a EscrowPaused event raised by the SlashEscrowFactory contract.
type SlashEscrowFactoryEscrowPaused struct {
	OperatorSet OperatorSet
	SlashId     *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterEscrowPaused is a free log retrieval operation binding the contract event 0x050add19b1a78a4240cdebc8747899275f2dd070c88e83904a37ff7d1a539744.
//
// Solidity: event EscrowPaused((address,uint32) operatorSet, uint256 slashId)
func (_SlashEscrowFactory *SlashEscrowFactoryFilterer) FilterEscrowPaused(opts *bind.FilterOpts) (*SlashEscrowFactoryEscrowPausedIterator, error) {

	logs, sub, err := _SlashEscrowFactory.contract.FilterLogs(opts, "EscrowPaused")
	if err != nil {
		return nil, err
	}
	return &SlashEscrowFactoryEscrowPausedIterator{contract: _SlashEscrowFactory.contract, event: "EscrowPaused", logs: logs, sub: sub}, nil
}

// WatchEscrowPaused is a free log subscription operation binding the contract event 0x050add19b1a78a4240cdebc8747899275f2dd070c88e83904a37ff7d1a539744.
//
// Solidity: event EscrowPaused((address,uint32) operatorSet, uint256 slashId)
func (_SlashEscrowFactory *SlashEscrowFactoryFilterer) WatchEscrowPaused(opts *bind.WatchOpts, sink chan<- *SlashEscrowFactoryEscrowPaused) (event.Subscription, error) {

	logs, sub, err := _SlashEscrowFactory.contract.WatchLogs(opts, "EscrowPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlashEscrowFactoryEscrowPaused)
				if err := _SlashEscrowFactory.contract.UnpackLog(event, "EscrowPaused", log); err != nil {
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

// ParseEscrowPaused is a log parse operation binding the contract event 0x050add19b1a78a4240cdebc8747899275f2dd070c88e83904a37ff7d1a539744.
//
// Solidity: event EscrowPaused((address,uint32) operatorSet, uint256 slashId)
func (_SlashEscrowFactory *SlashEscrowFactoryFilterer) ParseEscrowPaused(log types.Log) (*SlashEscrowFactoryEscrowPaused, error) {
	event := new(SlashEscrowFactoryEscrowPaused)
	if err := _SlashEscrowFactory.contract.UnpackLog(event, "EscrowPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlashEscrowFactoryEscrowUnpausedIterator is returned from FilterEscrowUnpaused and is used to iterate over the raw logs and unpacked data for EscrowUnpaused events raised by the SlashEscrowFactory contract.
type SlashEscrowFactoryEscrowUnpausedIterator struct {
	Event *SlashEscrowFactoryEscrowUnpaused // Event containing the contract specifics and raw log

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
func (it *SlashEscrowFactoryEscrowUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlashEscrowFactoryEscrowUnpaused)
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
		it.Event = new(SlashEscrowFactoryEscrowUnpaused)
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
func (it *SlashEscrowFactoryEscrowUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlashEscrowFactoryEscrowUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlashEscrowFactoryEscrowUnpaused represents a EscrowUnpaused event raised by the SlashEscrowFactory contract.
type SlashEscrowFactoryEscrowUnpaused struct {
	OperatorSet OperatorSet
	SlashId     *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterEscrowUnpaused is a free log retrieval operation binding the contract event 0xb8877c6bf02d5f6603188eb653c9269271836b75b2012a5d7f5f233e7cf2f241.
//
// Solidity: event EscrowUnpaused((address,uint32) operatorSet, uint256 slashId)
func (_SlashEscrowFactory *SlashEscrowFactoryFilterer) FilterEscrowUnpaused(opts *bind.FilterOpts) (*SlashEscrowFactoryEscrowUnpausedIterator, error) {

	logs, sub, err := _SlashEscrowFactory.contract.FilterLogs(opts, "EscrowUnpaused")
	if err != nil {
		return nil, err
	}
	return &SlashEscrowFactoryEscrowUnpausedIterator{contract: _SlashEscrowFactory.contract, event: "EscrowUnpaused", logs: logs, sub: sub}, nil
}

// WatchEscrowUnpaused is a free log subscription operation binding the contract event 0xb8877c6bf02d5f6603188eb653c9269271836b75b2012a5d7f5f233e7cf2f241.
//
// Solidity: event EscrowUnpaused((address,uint32) operatorSet, uint256 slashId)
func (_SlashEscrowFactory *SlashEscrowFactoryFilterer) WatchEscrowUnpaused(opts *bind.WatchOpts, sink chan<- *SlashEscrowFactoryEscrowUnpaused) (event.Subscription, error) {

	logs, sub, err := _SlashEscrowFactory.contract.WatchLogs(opts, "EscrowUnpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlashEscrowFactoryEscrowUnpaused)
				if err := _SlashEscrowFactory.contract.UnpackLog(event, "EscrowUnpaused", log); err != nil {
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

// ParseEscrowUnpaused is a log parse operation binding the contract event 0xb8877c6bf02d5f6603188eb653c9269271836b75b2012a5d7f5f233e7cf2f241.
//
// Solidity: event EscrowUnpaused((address,uint32) operatorSet, uint256 slashId)
func (_SlashEscrowFactory *SlashEscrowFactoryFilterer) ParseEscrowUnpaused(log types.Log) (*SlashEscrowFactoryEscrowUnpaused, error) {
	event := new(SlashEscrowFactoryEscrowUnpaused)
	if err := _SlashEscrowFactory.contract.UnpackLog(event, "EscrowUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlashEscrowFactoryGlobalEscrowDelaySetIterator is returned from FilterGlobalEscrowDelaySet and is used to iterate over the raw logs and unpacked data for GlobalEscrowDelaySet events raised by the SlashEscrowFactory contract.
type SlashEscrowFactoryGlobalEscrowDelaySetIterator struct {
	Event *SlashEscrowFactoryGlobalEscrowDelaySet // Event containing the contract specifics and raw log

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
func (it *SlashEscrowFactoryGlobalEscrowDelaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlashEscrowFactoryGlobalEscrowDelaySet)
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
		it.Event = new(SlashEscrowFactoryGlobalEscrowDelaySet)
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
func (it *SlashEscrowFactoryGlobalEscrowDelaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlashEscrowFactoryGlobalEscrowDelaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlashEscrowFactoryGlobalEscrowDelaySet represents a GlobalEscrowDelaySet event raised by the SlashEscrowFactory contract.
type SlashEscrowFactoryGlobalEscrowDelaySet struct {
	Delay uint32
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterGlobalEscrowDelaySet is a free log retrieval operation binding the contract event 0x67d0077d22e4e06f761dd87f6c9f2310ac879c9ce17de50d381e05b72f45fbf6.
//
// Solidity: event GlobalEscrowDelaySet(uint32 delay)
func (_SlashEscrowFactory *SlashEscrowFactoryFilterer) FilterGlobalEscrowDelaySet(opts *bind.FilterOpts) (*SlashEscrowFactoryGlobalEscrowDelaySetIterator, error) {

	logs, sub, err := _SlashEscrowFactory.contract.FilterLogs(opts, "GlobalEscrowDelaySet")
	if err != nil {
		return nil, err
	}
	return &SlashEscrowFactoryGlobalEscrowDelaySetIterator{contract: _SlashEscrowFactory.contract, event: "GlobalEscrowDelaySet", logs: logs, sub: sub}, nil
}

// WatchGlobalEscrowDelaySet is a free log subscription operation binding the contract event 0x67d0077d22e4e06f761dd87f6c9f2310ac879c9ce17de50d381e05b72f45fbf6.
//
// Solidity: event GlobalEscrowDelaySet(uint32 delay)
func (_SlashEscrowFactory *SlashEscrowFactoryFilterer) WatchGlobalEscrowDelaySet(opts *bind.WatchOpts, sink chan<- *SlashEscrowFactoryGlobalEscrowDelaySet) (event.Subscription, error) {

	logs, sub, err := _SlashEscrowFactory.contract.WatchLogs(opts, "GlobalEscrowDelaySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlashEscrowFactoryGlobalEscrowDelaySet)
				if err := _SlashEscrowFactory.contract.UnpackLog(event, "GlobalEscrowDelaySet", log); err != nil {
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

// ParseGlobalEscrowDelaySet is a log parse operation binding the contract event 0x67d0077d22e4e06f761dd87f6c9f2310ac879c9ce17de50d381e05b72f45fbf6.
//
// Solidity: event GlobalEscrowDelaySet(uint32 delay)
func (_SlashEscrowFactory *SlashEscrowFactoryFilterer) ParseGlobalEscrowDelaySet(log types.Log) (*SlashEscrowFactoryGlobalEscrowDelaySet, error) {
	event := new(SlashEscrowFactoryGlobalEscrowDelaySet)
	if err := _SlashEscrowFactory.contract.UnpackLog(event, "GlobalEscrowDelaySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlashEscrowFactoryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SlashEscrowFactory contract.
type SlashEscrowFactoryInitializedIterator struct {
	Event *SlashEscrowFactoryInitialized // Event containing the contract specifics and raw log

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
func (it *SlashEscrowFactoryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlashEscrowFactoryInitialized)
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
		it.Event = new(SlashEscrowFactoryInitialized)
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
func (it *SlashEscrowFactoryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlashEscrowFactoryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlashEscrowFactoryInitialized represents a Initialized event raised by the SlashEscrowFactory contract.
type SlashEscrowFactoryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SlashEscrowFactory *SlashEscrowFactoryFilterer) FilterInitialized(opts *bind.FilterOpts) (*SlashEscrowFactoryInitializedIterator, error) {

	logs, sub, err := _SlashEscrowFactory.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SlashEscrowFactoryInitializedIterator{contract: _SlashEscrowFactory.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SlashEscrowFactory *SlashEscrowFactoryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SlashEscrowFactoryInitialized) (event.Subscription, error) {

	logs, sub, err := _SlashEscrowFactory.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlashEscrowFactoryInitialized)
				if err := _SlashEscrowFactory.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_SlashEscrowFactory *SlashEscrowFactoryFilterer) ParseInitialized(log types.Log) (*SlashEscrowFactoryInitialized, error) {
	event := new(SlashEscrowFactoryInitialized)
	if err := _SlashEscrowFactory.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlashEscrowFactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SlashEscrowFactory contract.
type SlashEscrowFactoryOwnershipTransferredIterator struct {
	Event *SlashEscrowFactoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SlashEscrowFactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlashEscrowFactoryOwnershipTransferred)
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
		it.Event = new(SlashEscrowFactoryOwnershipTransferred)
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
func (it *SlashEscrowFactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlashEscrowFactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlashEscrowFactoryOwnershipTransferred represents a OwnershipTransferred event raised by the SlashEscrowFactory contract.
type SlashEscrowFactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SlashEscrowFactory *SlashEscrowFactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SlashEscrowFactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SlashEscrowFactory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SlashEscrowFactoryOwnershipTransferredIterator{contract: _SlashEscrowFactory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SlashEscrowFactory *SlashEscrowFactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SlashEscrowFactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SlashEscrowFactory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlashEscrowFactoryOwnershipTransferred)
				if err := _SlashEscrowFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_SlashEscrowFactory *SlashEscrowFactoryFilterer) ParseOwnershipTransferred(log types.Log) (*SlashEscrowFactoryOwnershipTransferred, error) {
	event := new(SlashEscrowFactoryOwnershipTransferred)
	if err := _SlashEscrowFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlashEscrowFactoryPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the SlashEscrowFactory contract.
type SlashEscrowFactoryPausedIterator struct {
	Event *SlashEscrowFactoryPaused // Event containing the contract specifics and raw log

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
func (it *SlashEscrowFactoryPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlashEscrowFactoryPaused)
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
		it.Event = new(SlashEscrowFactoryPaused)
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
func (it *SlashEscrowFactoryPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlashEscrowFactoryPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlashEscrowFactoryPaused represents a Paused event raised by the SlashEscrowFactory contract.
type SlashEscrowFactoryPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_SlashEscrowFactory *SlashEscrowFactoryFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*SlashEscrowFactoryPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SlashEscrowFactory.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &SlashEscrowFactoryPausedIterator{contract: _SlashEscrowFactory.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_SlashEscrowFactory *SlashEscrowFactoryFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *SlashEscrowFactoryPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SlashEscrowFactory.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlashEscrowFactoryPaused)
				if err := _SlashEscrowFactory.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_SlashEscrowFactory *SlashEscrowFactoryFilterer) ParsePaused(log types.Log) (*SlashEscrowFactoryPaused, error) {
	event := new(SlashEscrowFactoryPaused)
	if err := _SlashEscrowFactory.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlashEscrowFactoryStartEscrowIterator is returned from FilterStartEscrow and is used to iterate over the raw logs and unpacked data for StartEscrow events raised by the SlashEscrowFactory contract.
type SlashEscrowFactoryStartEscrowIterator struct {
	Event *SlashEscrowFactoryStartEscrow // Event containing the contract specifics and raw log

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
func (it *SlashEscrowFactoryStartEscrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlashEscrowFactoryStartEscrow)
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
		it.Event = new(SlashEscrowFactoryStartEscrow)
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
func (it *SlashEscrowFactoryStartEscrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlashEscrowFactoryStartEscrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlashEscrowFactoryStartEscrow represents a StartEscrow event raised by the SlashEscrowFactory contract.
type SlashEscrowFactoryStartEscrow struct {
	OperatorSet OperatorSet
	SlashId     *big.Int
	Strategy    common.Address
	StartBlock  uint32
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStartEscrow is a free log retrieval operation binding the contract event 0x3afae24c1d3dd2ce3649054ad82458a8c9967ebd9ce10a9a5a3d059f55bfaedb.
//
// Solidity: event StartEscrow((address,uint32) operatorSet, uint256 slashId, address strategy, uint32 startBlock)
func (_SlashEscrowFactory *SlashEscrowFactoryFilterer) FilterStartEscrow(opts *bind.FilterOpts) (*SlashEscrowFactoryStartEscrowIterator, error) {

	logs, sub, err := _SlashEscrowFactory.contract.FilterLogs(opts, "StartEscrow")
	if err != nil {
		return nil, err
	}
	return &SlashEscrowFactoryStartEscrowIterator{contract: _SlashEscrowFactory.contract, event: "StartEscrow", logs: logs, sub: sub}, nil
}

// WatchStartEscrow is a free log subscription operation binding the contract event 0x3afae24c1d3dd2ce3649054ad82458a8c9967ebd9ce10a9a5a3d059f55bfaedb.
//
// Solidity: event StartEscrow((address,uint32) operatorSet, uint256 slashId, address strategy, uint32 startBlock)
func (_SlashEscrowFactory *SlashEscrowFactoryFilterer) WatchStartEscrow(opts *bind.WatchOpts, sink chan<- *SlashEscrowFactoryStartEscrow) (event.Subscription, error) {

	logs, sub, err := _SlashEscrowFactory.contract.WatchLogs(opts, "StartEscrow")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlashEscrowFactoryStartEscrow)
				if err := _SlashEscrowFactory.contract.UnpackLog(event, "StartEscrow", log); err != nil {
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

// ParseStartEscrow is a log parse operation binding the contract event 0x3afae24c1d3dd2ce3649054ad82458a8c9967ebd9ce10a9a5a3d059f55bfaedb.
//
// Solidity: event StartEscrow((address,uint32) operatorSet, uint256 slashId, address strategy, uint32 startBlock)
func (_SlashEscrowFactory *SlashEscrowFactoryFilterer) ParseStartEscrow(log types.Log) (*SlashEscrowFactoryStartEscrow, error) {
	event := new(SlashEscrowFactoryStartEscrow)
	if err := _SlashEscrowFactory.contract.UnpackLog(event, "StartEscrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlashEscrowFactoryStrategyEscrowDelaySetIterator is returned from FilterStrategyEscrowDelaySet and is used to iterate over the raw logs and unpacked data for StrategyEscrowDelaySet events raised by the SlashEscrowFactory contract.
type SlashEscrowFactoryStrategyEscrowDelaySetIterator struct {
	Event *SlashEscrowFactoryStrategyEscrowDelaySet // Event containing the contract specifics and raw log

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
func (it *SlashEscrowFactoryStrategyEscrowDelaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlashEscrowFactoryStrategyEscrowDelaySet)
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
		it.Event = new(SlashEscrowFactoryStrategyEscrowDelaySet)
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
func (it *SlashEscrowFactoryStrategyEscrowDelaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlashEscrowFactoryStrategyEscrowDelaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlashEscrowFactoryStrategyEscrowDelaySet represents a StrategyEscrowDelaySet event raised by the SlashEscrowFactory contract.
type SlashEscrowFactoryStrategyEscrowDelaySet struct {
	Strategy common.Address
	Delay    uint32
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStrategyEscrowDelaySet is a free log retrieval operation binding the contract event 0x5d2b33f07ae22a809e79005f96ffac70c3715df85a3b011b025e0a86a23a007b.
//
// Solidity: event StrategyEscrowDelaySet(address strategy, uint32 delay)
func (_SlashEscrowFactory *SlashEscrowFactoryFilterer) FilterStrategyEscrowDelaySet(opts *bind.FilterOpts) (*SlashEscrowFactoryStrategyEscrowDelaySetIterator, error) {

	logs, sub, err := _SlashEscrowFactory.contract.FilterLogs(opts, "StrategyEscrowDelaySet")
	if err != nil {
		return nil, err
	}
	return &SlashEscrowFactoryStrategyEscrowDelaySetIterator{contract: _SlashEscrowFactory.contract, event: "StrategyEscrowDelaySet", logs: logs, sub: sub}, nil
}

// WatchStrategyEscrowDelaySet is a free log subscription operation binding the contract event 0x5d2b33f07ae22a809e79005f96ffac70c3715df85a3b011b025e0a86a23a007b.
//
// Solidity: event StrategyEscrowDelaySet(address strategy, uint32 delay)
func (_SlashEscrowFactory *SlashEscrowFactoryFilterer) WatchStrategyEscrowDelaySet(opts *bind.WatchOpts, sink chan<- *SlashEscrowFactoryStrategyEscrowDelaySet) (event.Subscription, error) {

	logs, sub, err := _SlashEscrowFactory.contract.WatchLogs(opts, "StrategyEscrowDelaySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlashEscrowFactoryStrategyEscrowDelaySet)
				if err := _SlashEscrowFactory.contract.UnpackLog(event, "StrategyEscrowDelaySet", log); err != nil {
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

// ParseStrategyEscrowDelaySet is a log parse operation binding the contract event 0x5d2b33f07ae22a809e79005f96ffac70c3715df85a3b011b025e0a86a23a007b.
//
// Solidity: event StrategyEscrowDelaySet(address strategy, uint32 delay)
func (_SlashEscrowFactory *SlashEscrowFactoryFilterer) ParseStrategyEscrowDelaySet(log types.Log) (*SlashEscrowFactoryStrategyEscrowDelaySet, error) {
	event := new(SlashEscrowFactoryStrategyEscrowDelaySet)
	if err := _SlashEscrowFactory.contract.UnpackLog(event, "StrategyEscrowDelaySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlashEscrowFactoryUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the SlashEscrowFactory contract.
type SlashEscrowFactoryUnpausedIterator struct {
	Event *SlashEscrowFactoryUnpaused // Event containing the contract specifics and raw log

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
func (it *SlashEscrowFactoryUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlashEscrowFactoryUnpaused)
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
		it.Event = new(SlashEscrowFactoryUnpaused)
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
func (it *SlashEscrowFactoryUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlashEscrowFactoryUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlashEscrowFactoryUnpaused represents a Unpaused event raised by the SlashEscrowFactory contract.
type SlashEscrowFactoryUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_SlashEscrowFactory *SlashEscrowFactoryFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*SlashEscrowFactoryUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SlashEscrowFactory.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &SlashEscrowFactoryUnpausedIterator{contract: _SlashEscrowFactory.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_SlashEscrowFactory *SlashEscrowFactoryFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *SlashEscrowFactoryUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SlashEscrowFactory.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlashEscrowFactoryUnpaused)
				if err := _SlashEscrowFactory.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_SlashEscrowFactory *SlashEscrowFactoryFilterer) ParseUnpaused(log types.Log) (*SlashEscrowFactoryUnpaused, error) {
	event := new(SlashEscrowFactoryUnpaused)
	if err := _SlashEscrowFactory.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
