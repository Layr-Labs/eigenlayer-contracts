// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package CrossChainRegistry

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

// ICrossChainRegistryTypesOperatorSetConfig is an auto generated low-level Go binding around an user-defined struct.
type ICrossChainRegistryTypesOperatorSetConfig struct {
	Owner              common.Address
	MaxStalenessPeriod uint32
}

// OperatorSet is an auto generated low-level Go binding around an user-defined struct.
type OperatorSet struct {
	Avs common.Address
	Id  uint32
}

// CrossChainRegistryMetaData contains all meta data concerning the CrossChainRegistry contract.
var CrossChainRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_allocationManager\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"},{\"name\":\"_keyRegistrar\",\"type\":\"address\",\"internalType\":\"contractIKeyRegistrar\"},{\"name\":\"_permissionController\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"_version\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addChainIDsToWhitelist\",\"inputs\":[{\"name\":\"chainIDs\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"operatorTableUpdaters\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allocationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateOperatorTableBytes\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createGenerationReservation\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operatorTableCalculator\",\"type\":\"address\",\"internalType\":\"contractIOperatorTableCalculator\"},{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structICrossChainRegistryTypes.OperatorSetConfig\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxStalenessPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getActiveGenerationReservationCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getActiveGenerationReservations\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getActiveGenerationReservationsByRange\",\"inputs\":[{\"name\":\"startIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"endIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorSetConfig\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structICrossChainRegistryTypes.OperatorSetConfig\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxStalenessPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorTableCalculator\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIOperatorTableCalculator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTableUpdateCadence\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasActiveGenerationReservation\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialTableUpdateCadence\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"initialPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"keyRegistrar\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIKeyRegistrar\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"permissionController\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeChainIDsFromWhitelist\",\"inputs\":[{\"name\":\"chainIDs\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeGenerationReservation\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOperatorSetConfig\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structICrossChainRegistryTypes.OperatorSetConfig\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxStalenessPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOperatorTableCalculator\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operatorTableCalculator\",\"type\":\"address\",\"internalType\":\"contractIOperatorTableCalculator\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTableUpdateCadence\",\"inputs\":[{\"name\":\"tableUpdateCadence\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"ChainIDAddedToWhitelist\",\"inputs\":[{\"name\":\"chainID\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"operatorTableUpdater\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainIDRemovedFromWhitelist\",\"inputs\":[{\"name\":\"chainID\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GenerationReservationCreated\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GenerationReservationRemoved\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSetConfigRemoved\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSetConfigSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structICrossChainRegistryTypes.OperatorSetConfig\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxStalenessPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorTableCalculatorRemoved\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorTableCalculatorSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operatorTableCalculator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIOperatorTableCalculator\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TableUpdateCadenceSet\",\"inputs\":[{\"name\":\"tableUpdateCadence\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ArrayLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainIDAlreadyWhitelisted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainIDNotWhitelisted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CurrentlyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EmptyChainIDsArray\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"GenerationReservationAlreadyExists\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"GenerationReservationDoesNotExist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidChainId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidEndIndex\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNewPausedStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperatorSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPermissions\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRange\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidStalenessPeriod\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTableUpdateCadence\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"KeyTypeNotSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyPauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnpauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]}]",
	Bin: "0x610120604052348015610010575f5ffd5b5060405161281438038061281483398101604081905261002f916101c1565b80838686856001600160a01b03811661005b576040516339b190bb60e11b815260040160405180910390fd5b6001600160a01b0390811660805291821660a052811660c0521660e05261008181610098565b610100525061008e6100de565b5050505050610319565b5f5f829050601f815111156100cb578260405163305a27a960e01b81526004016100c291906102be565b60405180910390fd5b80516100d6826102f3565b179392505050565b5f54610100900460ff16156101455760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b60648201526084016100c2565b5f5460ff90811614610194575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146101aa575f5ffd5b50565b634e487b7160e01b5f52604160045260245ffd5b5f5f5f5f5f60a086880312156101d5575f5ffd5b85516101e081610196565b60208701519095506101f181610196565b604087015190945061020281610196565b606087015190935061021381610196565b60808701519092506001600160401b0381111561022e575f5ffd5b8601601f8101881361023e575f5ffd5b80516001600160401b03811115610257576102576101ad565b604051601f8201601f19908116603f011681016001600160401b0381118282101715610285576102856101ad565b6040528181528282016020018a101561029c575f5ffd5b8160208401602083015e5f602083830101528093505050509295509295909350565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b80516020808301519190811015610313575f198160200360031b1b821691505b50919050565b60805160a05160c05160e0516101005161247f6103955f395f610b7201525f81816102e601526117d701525f818161028701528181610a29015261112201525f81816103ee015281816107920152818161092201528181610c05015261107901525f81816103800152818161170b0152611a9b015261247f5ff3fe608060405234801561000f575f5ffd5b50600436106101d1575f3560e01c8063715018a6116100fe578063ca8aa7c71161009e578063d9a6729e1161006e578063d9a6729e1461044b578063dfbd9dfd1461045e578063f2fde38b14610471578063fabc1cbc14610484575f5ffd5b8063ca8aa7c7146103e9578063d09b978b14610410578063d504491114610425578063d6db9e2514610438575f5ffd5b80638da5cb5b116100d95780638da5cb5b146103a2578063ac505f4b146103b3578063b186a60e146103cb578063c4bffe2b146103d3575f5ffd5b8063715018a61461036057806375e4b53914610368578063886f11951461037b575f5ffd5b80633ec45c7e11610174578063595c6a6711610144578063595c6a67146103105780635ac86ab7146103185780635c975abb1461033b5780636c55a37f1461034d575f5ffd5b80633ec45c7e1461028257806341ee6d0e146102c15780634657e26a146102e157806354fd4d5014610308575f5ffd5b80631ca9142a116101af5780631ca9142a1461021057806321fa7fdc14610223578063277e1e621461024c57806336b200de1461025f575f5ffd5b806304e98be3146101d55780630f19aaef146101ea578063136439dd146101fd575b5f5ffd5b6101e86101e3366004611d72565b610497565b005b6101e86101f8366004611e0a565b6105db565b6101e861020b366004611e46565b610701565b6101e861021e366004611e73565b61073b565b610236610231366004611f4c565b61086b565b6040516102439190611f84565b60405180910390f35b6101e861025a366004611f92565b6108cb565b61027261026d366004611f4c565b610a09565b6040519015158152602001610243565b6102a97f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b039091168152602001610243565b6102d46102cf366004611fc5565b610a24565b604051610243919061200d565b6102a97f000000000000000000000000000000000000000000000000000000000000000081565b6102d4610b6b565b6101e8610b9b565b61027261032636600461201f565b606654600160ff9092169190911b9081161490565b6066545b604051908152602001610243565b6101e861035b366004611fc5565b610baf565b6101e8610dd2565b6102a9610376366004611f4c565b610de3565b6102a97f000000000000000000000000000000000000000000000000000000000000000081565b6033546001600160a01b03166102a9565b609e5460405163ffffffff9091168152602001610243565b61033f610e0f565b6103db610e1a565b60405161024392919061203f565b6102a97f000000000000000000000000000000000000000000000000000000000000000081565b610418610f30565b60405161024391906120c8565b6101e8610433366004612115565b611023565b6101e8610446366004612159565b611269565b610418610459366004612172565b61127d565b6101e861046c366004612192565b6113c7565b6101e861047f3660046121d1565b61146d565b6101e8610492366004611e46565b6114e3565b61049f611550565b60036104aa816115aa565b8382146104ca5760405163512509d360e11b815260040160405180910390fd5b5f5b848110156105d3575f8686838181106104e7576104e76121ec565b905060200201359050805f0361051057604051633d23e4d160e11b815260040160405180910390fd5b61054481868685818110610526576105266121ec565b905060200201602081019061053b91906121d1565b609b91906115d5565b610561576040516324bf631b60e11b815260040160405180910390fd5b7f7a0a76d85b582b17996dd7371a407aa7a79b870db8539247fba315c7b6beff6281868685818110610595576105956121ec565b90506020020160208101906105aa91906121d1565b604080519283526001600160a01b0390911660208301520160405180910390a1506001016104cc565b505050505050565b5f54610100900460ff16158080156105f957505f54600160ff909116105b806106125750303b15801561061257505f5460ff166001145b61067a5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b5f805460ff19166001179055801561069b575f805461ff0019166101001790555b6106a4846115f4565b6106ad83611645565b6106b6826116b9565b80156106fb575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050565b6107096116f6565b606654818116811461072e5760405163c61dca5d60e01b815260040160405180910390fd5b610737826116b9565b5050565b6001610746816115aa565b61075360208401846121d1565b61075c81611799565b6107795760405163932d94f760e01b815260040160405180910390fd5b6040516304c1b8eb60e31b815284906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063260dc758906107c7908490600401612231565b602060405180830381865afa1580156107e2573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610806919061223f565b61082357604051631fb1705560e21b815260040160405180910390fd5b8461083661026d36839003830183611f4c565b61085357604051634d2baea960e11b815260040160405180910390fd5b6105d361086536889003880188611f4c565b86611843565b604080518082019091525f8082526020820152609a5f61088a846118bd565b815260208082019290925260409081015f208151808301909252546001600160a01b0381168252600160a01b900463ffffffff169181019190915292915050565b60026108d6816115aa565b6108e360208401846121d1565b6108ec81611799565b6109095760405163932d94f760e01b815260040160405180910390fd5b6040516304c1b8eb60e31b815284906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063260dc75890610957908490600401612231565b602060405180830381865afa158015610972573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610996919061223f565b6109b357604051631fb1705560e21b815260040160405180910390fd5b846109c661026d36839003830183611f4c565b6109e357604051634d2baea960e11b815260040160405180910390fd5b6105d36109f536889003880188611f4c565b610a0436889003880188611f4c565b611920565b5f610a1e610a16836118bd565b6097906119e9565b92915050565b6060817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316637cffe48c846040518263ffffffff1660e01b8152600401610a739190612231565b602060405180830381865afa158015610a8e573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610ab2919061225e565b610ac461023136869003860186611f4c565b610ad661037636879003870187611f4c565b6001600160a01b03166341ee6d0e866040518263ffffffff1660e01b8152600401610b019190612231565b5f60405180830381865afa158015610b1b573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610b42919081019061227c565b604051602001610b559493929190612324565b6040516020818303038152906040529050919050565b6060610b967f0000000000000000000000000000000000000000000000000000000000000000611a00565b905090565b610ba36116f6565b610bad5f196116b9565b565b5f610bb9816115aa565b610bc660208301836121d1565b610bcf81611799565b610bec5760405163932d94f760e01b815260040160405180910390fd5b6040516304c1b8eb60e31b815283906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063260dc75890610c3a908490600401612231565b602060405180830381865afa158015610c55573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610c79919061223f565b610c9657604051631fb1705560e21b815260040160405180910390fd5b83610ca961026d36839003830183611f4c565b610cc657604051634d2baea960e11b815260040160405180910390fd5b5f610cde610cd936889003880188611f4c565b6118bd565b5f818152609960205260409081902080546001600160a01b0319169055519091507fd7811913efd5d98fc7ea0d1fdd022b3d31987815360842d05b1d1cf55578d16a90610d2c908890612231565b60405180910390a15f818152609a60205260409081902080546001600160c01b0319169055517f210a1118a869246162804e2a7f21ef808ebd93f4be7ed512014fe29a7a8be02e90610d7f908890612231565b60405180910390a1610d92609782611a3d565b507f4ffdfdd59e9e1e3c301608788f78dd458e61cb8c045ca92b62a7b484c80824fb86604051610dc29190612231565b60405180910390a1505050505050565b610dda611550565b610bad5f6115f4565b5f60995f610df0846118bd565b815260208101919091526040015f20546001600160a01b031692915050565b5f610b966097611a48565b6060805f610e28609b611a51565b90505f8167ffffffffffffffff811115610e4457610e44611ea9565b604051908082528060200260200182016040528015610e6d578160200160208202803683370190505b5090505f8267ffffffffffffffff811115610e8a57610e8a611ea9565b604051908082528060200260200182016040528015610eb3578160200160208202803683370190505b5090505f5b83811015610f25575f80610ecd609b84611a5b565b9150915081858481518110610ee457610ee46121ec565b60200260200101818152505080848481518110610f0357610f036121ec565b6001600160a01b03909216602092830291909101909101525050600101610eb8565b509094909350915050565b60605f610f3d6097611a48565b90505f8167ffffffffffffffff811115610f5957610f59611ea9565b604051908082528060200260200182016040528015610f9d57816020015b604080518082019091525f8082526020820152815260200190600190039081610f775790505b5090505f5b8281101561101c575f610fb6609783611a78565b90505f610ff282604080518082019091525f80825260208201525060408051808201909152606082901c815263ffffffff909116602082015290565b905080848481518110611007576110076121ec565b60209081029190910101525050600101610fa2565b5092915050565b5f61102d816115aa565b61103a60208501856121d1565b61104381611799565b6110605760405163932d94f760e01b815260040160405180910390fd5b6040516304c1b8eb60e31b815285906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063260dc758906110ae908490600401612231565b602060405180830381865afa1580156110c9573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906110ed919061223f565b61110a57604051631fb1705560e21b815260040160405180910390fd5b5f604051631f3ff92360e21b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690637cffe48c90611157908a90600401612231565b602060405180830381865afa158015611172573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611196919061225e565b60028111156111a7576111a7612310565b036111c55760405163e57cacbd60e01b815260040160405180910390fd5b6111e26111da610cd936899003890189611f4c565b609790611a83565b6111ff57604051631883461560e01b815260040160405180910390fd5b7f4fb6efec7dd60036ce3a7af8d5c48425019daa0fb61eb471a966a7ac2c6fa6a68660405161122e9190612231565b60405180910390a161124861086536889003880188611f4c565b6105d361125a36889003880188611f4c565b610a0436879003870187611f4c565b611271611550565b61127a81611645565b50565b6060818311156112a05760405163561ce9bb60e01b815260040160405180910390fd5b6112aa6097611a48565b8211156112ca576040516302da361360e61b815260040160405180910390fd5b5f6112d58484612390565b90505f8167ffffffffffffffff8111156112f1576112f1611ea9565b60405190808252806020026020018201604052801561133557816020015b604080518082019091525f808252602082015281526020019060019003908161130f5790505b5090505f5b828110156113be575f61135861135083896123a3565b609790611a78565b90505f61139482604080518082019091525f80825260208201525060408051808201909152606082901c815263ffffffff909116602082015290565b9050808484815181106113a9576113a96121ec565b6020908102919091010152505060010161133a565b50949350505050565b6113cf611550565b60036113da816115aa565b5f5b828110156106fb575f8484838181106113f7576113f76121ec565b90506020020135905061141481609b611a8e90919063ffffffff16565b6114315760405163b3f92ba160e01b815260040160405180910390fd5b6040518181527f6824d36084ecf2cd819b137cb5d837cc6e73afce1e0e348c9fdecaa81d0341e59060200160405180910390a1506001016113dc565b611475611550565b6001600160a01b0381166114da5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610671565b61127a816115f4565b6114eb611a99565b606654801982198116146115125760405163c61dca5d60e01b815260040160405180910390fd5b606682905560405182815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c9060200160405180910390a25050565b6033546001600160a01b03163314610bad5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610671565b606654600160ff83161b9081160361127a5760405163840a48d560e01b815260040160405180910390fd5b5f6115ea84846001600160a01b038516611b4a565b90505b9392505050565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b5f8163ffffffff161161166b576040516316d98e1b60e31b815260040160405180910390fd5b609e805463ffffffff191663ffffffff83169081179091556040519081527f4fbcd0cca70015b33db8af4aa4f2bd6fd6c1efa9460b8e2333f252c1467a63279060200160405180910390a150565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a250565b60405163237dfb4760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906346fbf68e90602401602060405180830381865afa158015611758573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061177c919061223f565b610bad57604051631d77d47760e21b815260040160405180910390fd5b604051631beb2b9760e31b81526001600160a01b0382811660048301523360248301523060448301525f80356001600160e01b0319166064840152917f00000000000000000000000000000000000000000000000000000000000000009091169063df595cb8906084016020604051808303815f875af115801561181f573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a1e919061223f565b8060995f611850856118bd565b81526020019081526020015f205f6101000a8154816001600160a01b0302191690836001600160a01b031602179055507f7f7ccafd92d20fdb39dee184a0dce002a9da420ed0def461f2a027abc9b3f6df82826040516118b19291906123b6565b60405180910390a15050565b5f815f0151826020015163ffffffff1660405160200161190892919060609290921b6bffffffffffffffffffffffff1916825260a01b6001600160a01b031916601482015260200190565b604051602081830303815290604052610a1e906123dc565b602081015163ffffffff1615806119475750609e54602082015163ffffffff918216911610155b61196457604051632e46483160e11b815260040160405180910390fd5b80609a5f611971856118bd565b815260208082019290925260409081015f2083518154949093015163ffffffff16600160a01b026001600160c01b03199094166001600160a01b0390931692909217929092179055517f3147846ee526009000671c20380b856a633345691300f82585f90034715cf0e2906118b190849084906123ff565b5f81815260018301602052604081205415156115ed565b60605f611a0c83611b66565b6040805160208082528183019092529192505f91906020820181803683375050509182525060208101929092525090565b5f6115ed8383611b8d565b5f610a1e825490565b5f610a1e82611c70565b5f808080611a698686611c7a565b909450925050505b9250929050565b5f6115ed8383611ca3565b5f6115ed8383611cc9565b5f6115ed8383611d15565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611af5573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611b19919061241a565b6001600160a01b0316336001600160a01b031614610bad5760405163794821ff60e01b815260040160405180910390fd5b5f82815260028401602052604081208290556115ea8484611a83565b5f60ff8216601f811115610a1e57604051632cd44ac360e21b815260040160405180910390fd5b5f8181526001830160205260408120548015611c67575f611baf600183612390565b85549091505f90611bc290600190612390565b9050818114611c21575f865f018281548110611be057611be06121ec565b905f5260205f200154905080875f018481548110611c0057611c006121ec565b5f918252602080832090910192909255918252600188019052604090208390555b8554869080611c3257611c32612435565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f905560019350505050610a1e565b5f915050610a1e565b5f610a1e82611a48565b5f8080611c878585611a78565b5f81815260029690960160205260409095205494959350505050565b5f825f018281548110611cb857611cb86121ec565b905f5260205f200154905092915050565b5f818152600183016020526040812054611d0e57508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155610a1e565b505f610a1e565b5f81815260028301602052604081208190556115ed8383611a3d565b5f5f83601f840112611d41575f5ffd5b50813567ffffffffffffffff811115611d58575f5ffd5b6020830191508360208260051b8501011115611a71575f5ffd5b5f5f5f5f60408587031215611d85575f5ffd5b843567ffffffffffffffff811115611d9b575f5ffd5b611da787828801611d31565b909550935050602085013567ffffffffffffffff811115611dc6575f5ffd5b611dd287828801611d31565b95989497509550505050565b6001600160a01b038116811461127a575f5ffd5b803563ffffffff81168114611e05575f5ffd5b919050565b5f5f5f60608486031215611e1c575f5ffd5b8335611e2781611dde565b9250611e3560208501611df2565b929592945050506040919091013590565b5f60208284031215611e56575f5ffd5b5035919050565b5f60408284031215611e6d575f5ffd5b50919050565b5f5f60608385031215611e84575f5ffd5b611e8e8484611e5d565b91506040830135611e9e81611dde565b809150509250929050565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f1916810167ffffffffffffffff81118282101715611ee657611ee6611ea9565b604052919050565b5f60408284031215611efe575f5ffd5b6040805190810167ffffffffffffffff81118282101715611f2157611f21611ea9565b6040529050808235611f3281611dde565b8152611f4060208401611df2565b60208201525092915050565b5f60408284031215611f5c575f5ffd5b6115ed8383611eee565b80516001600160a01b0316825260209081015163ffffffff16910152565b60408101610a1e8284611f66565b5f5f60808385031215611fa3575f5ffd5b611fad8484611e5d565b9150611fbc8460408501611e5d565b90509250929050565b5f60408284031215611fd5575f5ffd5b6115ed8383611e5d565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f6115ed6020830184611fdf565b5f6020828403121561202f575f5ffd5b813560ff811681146115ed575f5ffd5b604080825283519082018190525f9060208501906060840190835b8181101561207857835183526020938401939092019160010161205a565b5050838103602080860191909152855180835291810192508501905f5b818110156120bc5782516001600160a01b0316845260209384019390920191600101612095565b50919695505050505050565b602080825282518282018190525f918401906040840190835b8181101561210a576120f4838551611f66565b60209390930192604092909201916001016120e1565b509095945050505050565b5f5f5f60a08486031215612127575f5ffd5b6121318585611e5d565b9250604084013561214181611dde565b91506121508560608601611e5d565b90509250925092565b5f60208284031215612169575f5ffd5b6115ed82611df2565b5f5f60408385031215612183575f5ffd5b50508035926020909101359150565b5f5f602083850312156121a3575f5ffd5b823567ffffffffffffffff8111156121b9575f5ffd5b6121c585828601611d31565b90969095509350505050565b5f602082840312156121e1575f5ffd5b81356115ed81611dde565b634e487b7160e01b5f52603260045260245ffd5b803561220b81611dde565b6001600160a01b0316825263ffffffff61222760208301611df2565b1660208301525050565b60408101610a1e8284612200565b5f6020828403121561224f575f5ffd5b815180151581146115ed575f5ffd5b5f6020828403121561226e575f5ffd5b8151600381106115ed575f5ffd5b5f6020828403121561228c575f5ffd5b815167ffffffffffffffff8111156122a2575f5ffd5b8201601f810184136122b2575f5ffd5b805167ffffffffffffffff8111156122cc576122cc611ea9565b6122df601f8201601f1916602001611ebd565b8181528560208385010111156122f3575f5ffd5b8160208401602083015e5f91810160200191909152949350505050565b634e487b7160e01b5f52602160045260245ffd5b61232e8186612200565b5f6003851061234b57634e487b7160e01b5f52602160045260245ffd5b84604083015261235e6060830185611f66565b60c060a083015261237260c0830184611fdf565b9695505050505050565b634e487b7160e01b5f52601160045260245ffd5b81810381811115610a1e57610a1e61237c565b80820180821115610a1e57610a1e61237c565b606081016123c48285611f66565b6001600160a01b039290921660409190910152919050565b80516020808301519190811015611e6d575f1960209190910360031b1b16919050565b6080810161240d8285611f66565b6115ed6040830184611f66565b5f6020828403121561242a575f5ffd5b81516115ed81611dde565b634e487b7160e01b5f52603160045260245ffdfea26469706673582212209eee65670a0acd546a18377ada41e4f6518f10b5c3d618724117493c69a324be64736f6c634300081b0033",
}

// CrossChainRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use CrossChainRegistryMetaData.ABI instead.
var CrossChainRegistryABI = CrossChainRegistryMetaData.ABI

// CrossChainRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CrossChainRegistryMetaData.Bin instead.
var CrossChainRegistryBin = CrossChainRegistryMetaData.Bin

// DeployCrossChainRegistry deploys a new Ethereum contract, binding an instance of CrossChainRegistry to it.
func DeployCrossChainRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, _allocationManager common.Address, _keyRegistrar common.Address, _permissionController common.Address, _pauserRegistry common.Address, _version string) (common.Address, *types.Transaction, *CrossChainRegistry, error) {
	parsed, err := CrossChainRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CrossChainRegistryBin), backend, _allocationManager, _keyRegistrar, _permissionController, _pauserRegistry, _version)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CrossChainRegistry{CrossChainRegistryCaller: CrossChainRegistryCaller{contract: contract}, CrossChainRegistryTransactor: CrossChainRegistryTransactor{contract: contract}, CrossChainRegistryFilterer: CrossChainRegistryFilterer{contract: contract}}, nil
}

// CrossChainRegistry is an auto generated Go binding around an Ethereum contract.
type CrossChainRegistry struct {
	CrossChainRegistryCaller     // Read-only binding to the contract
	CrossChainRegistryTransactor // Write-only binding to the contract
	CrossChainRegistryFilterer   // Log filterer for contract events
}

// CrossChainRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type CrossChainRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossChainRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CrossChainRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossChainRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrossChainRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossChainRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrossChainRegistrySession struct {
	Contract     *CrossChainRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// CrossChainRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrossChainRegistryCallerSession struct {
	Contract *CrossChainRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// CrossChainRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrossChainRegistryTransactorSession struct {
	Contract     *CrossChainRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// CrossChainRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type CrossChainRegistryRaw struct {
	Contract *CrossChainRegistry // Generic contract binding to access the raw methods on
}

// CrossChainRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrossChainRegistryCallerRaw struct {
	Contract *CrossChainRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// CrossChainRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrossChainRegistryTransactorRaw struct {
	Contract *CrossChainRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCrossChainRegistry creates a new instance of CrossChainRegistry, bound to a specific deployed contract.
func NewCrossChainRegistry(address common.Address, backend bind.ContractBackend) (*CrossChainRegistry, error) {
	contract, err := bindCrossChainRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CrossChainRegistry{CrossChainRegistryCaller: CrossChainRegistryCaller{contract: contract}, CrossChainRegistryTransactor: CrossChainRegistryTransactor{contract: contract}, CrossChainRegistryFilterer: CrossChainRegistryFilterer{contract: contract}}, nil
}

// NewCrossChainRegistryCaller creates a new read-only instance of CrossChainRegistry, bound to a specific deployed contract.
func NewCrossChainRegistryCaller(address common.Address, caller bind.ContractCaller) (*CrossChainRegistryCaller, error) {
	contract, err := bindCrossChainRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrossChainRegistryCaller{contract: contract}, nil
}

// NewCrossChainRegistryTransactor creates a new write-only instance of CrossChainRegistry, bound to a specific deployed contract.
func NewCrossChainRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*CrossChainRegistryTransactor, error) {
	contract, err := bindCrossChainRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrossChainRegistryTransactor{contract: contract}, nil
}

// NewCrossChainRegistryFilterer creates a new log filterer instance of CrossChainRegistry, bound to a specific deployed contract.
func NewCrossChainRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*CrossChainRegistryFilterer, error) {
	contract, err := bindCrossChainRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrossChainRegistryFilterer{contract: contract}, nil
}

// bindCrossChainRegistry binds a generic wrapper to an already deployed contract.
func bindCrossChainRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CrossChainRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossChainRegistry *CrossChainRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossChainRegistry.Contract.CrossChainRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossChainRegistry *CrossChainRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossChainRegistry.Contract.CrossChainRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossChainRegistry *CrossChainRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossChainRegistry.Contract.CrossChainRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossChainRegistry *CrossChainRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossChainRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossChainRegistry *CrossChainRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossChainRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossChainRegistry *CrossChainRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossChainRegistry.Contract.contract.Transact(opts, method, params...)
}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_CrossChainRegistry *CrossChainRegistryCaller) AllocationManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossChainRegistry.contract.Call(opts, &out, "allocationManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_CrossChainRegistry *CrossChainRegistrySession) AllocationManager() (common.Address, error) {
	return _CrossChainRegistry.Contract.AllocationManager(&_CrossChainRegistry.CallOpts)
}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_CrossChainRegistry *CrossChainRegistryCallerSession) AllocationManager() (common.Address, error) {
	return _CrossChainRegistry.Contract.AllocationManager(&_CrossChainRegistry.CallOpts)
}

// CalculateOperatorTableBytes is a free data retrieval call binding the contract method 0x41ee6d0e.
//
// Solidity: function calculateOperatorTableBytes((address,uint32) operatorSet) view returns(bytes)
func (_CrossChainRegistry *CrossChainRegistryCaller) CalculateOperatorTableBytes(opts *bind.CallOpts, operatorSet OperatorSet) ([]byte, error) {
	var out []interface{}
	err := _CrossChainRegistry.contract.Call(opts, &out, "calculateOperatorTableBytes", operatorSet)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// CalculateOperatorTableBytes is a free data retrieval call binding the contract method 0x41ee6d0e.
//
// Solidity: function calculateOperatorTableBytes((address,uint32) operatorSet) view returns(bytes)
func (_CrossChainRegistry *CrossChainRegistrySession) CalculateOperatorTableBytes(operatorSet OperatorSet) ([]byte, error) {
	return _CrossChainRegistry.Contract.CalculateOperatorTableBytes(&_CrossChainRegistry.CallOpts, operatorSet)
}

// CalculateOperatorTableBytes is a free data retrieval call binding the contract method 0x41ee6d0e.
//
// Solidity: function calculateOperatorTableBytes((address,uint32) operatorSet) view returns(bytes)
func (_CrossChainRegistry *CrossChainRegistryCallerSession) CalculateOperatorTableBytes(operatorSet OperatorSet) ([]byte, error) {
	return _CrossChainRegistry.Contract.CalculateOperatorTableBytes(&_CrossChainRegistry.CallOpts, operatorSet)
}

// GetActiveGenerationReservationCount is a free data retrieval call binding the contract method 0xb186a60e.
//
// Solidity: function getActiveGenerationReservationCount() view returns(uint256)
func (_CrossChainRegistry *CrossChainRegistryCaller) GetActiveGenerationReservationCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossChainRegistry.contract.Call(opts, &out, "getActiveGenerationReservationCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetActiveGenerationReservationCount is a free data retrieval call binding the contract method 0xb186a60e.
//
// Solidity: function getActiveGenerationReservationCount() view returns(uint256)
func (_CrossChainRegistry *CrossChainRegistrySession) GetActiveGenerationReservationCount() (*big.Int, error) {
	return _CrossChainRegistry.Contract.GetActiveGenerationReservationCount(&_CrossChainRegistry.CallOpts)
}

// GetActiveGenerationReservationCount is a free data retrieval call binding the contract method 0xb186a60e.
//
// Solidity: function getActiveGenerationReservationCount() view returns(uint256)
func (_CrossChainRegistry *CrossChainRegistryCallerSession) GetActiveGenerationReservationCount() (*big.Int, error) {
	return _CrossChainRegistry.Contract.GetActiveGenerationReservationCount(&_CrossChainRegistry.CallOpts)
}

// GetActiveGenerationReservations is a free data retrieval call binding the contract method 0xd09b978b.
//
// Solidity: function getActiveGenerationReservations() view returns((address,uint32)[])
func (_CrossChainRegistry *CrossChainRegistryCaller) GetActiveGenerationReservations(opts *bind.CallOpts) ([]OperatorSet, error) {
	var out []interface{}
	err := _CrossChainRegistry.contract.Call(opts, &out, "getActiveGenerationReservations")

	if err != nil {
		return *new([]OperatorSet), err
	}

	out0 := *abi.ConvertType(out[0], new([]OperatorSet)).(*[]OperatorSet)

	return out0, err

}

// GetActiveGenerationReservations is a free data retrieval call binding the contract method 0xd09b978b.
//
// Solidity: function getActiveGenerationReservations() view returns((address,uint32)[])
func (_CrossChainRegistry *CrossChainRegistrySession) GetActiveGenerationReservations() ([]OperatorSet, error) {
	return _CrossChainRegistry.Contract.GetActiveGenerationReservations(&_CrossChainRegistry.CallOpts)
}

// GetActiveGenerationReservations is a free data retrieval call binding the contract method 0xd09b978b.
//
// Solidity: function getActiveGenerationReservations() view returns((address,uint32)[])
func (_CrossChainRegistry *CrossChainRegistryCallerSession) GetActiveGenerationReservations() ([]OperatorSet, error) {
	return _CrossChainRegistry.Contract.GetActiveGenerationReservations(&_CrossChainRegistry.CallOpts)
}

// GetActiveGenerationReservationsByRange is a free data retrieval call binding the contract method 0xd9a6729e.
//
// Solidity: function getActiveGenerationReservationsByRange(uint256 startIndex, uint256 endIndex) view returns((address,uint32)[])
func (_CrossChainRegistry *CrossChainRegistryCaller) GetActiveGenerationReservationsByRange(opts *bind.CallOpts, startIndex *big.Int, endIndex *big.Int) ([]OperatorSet, error) {
	var out []interface{}
	err := _CrossChainRegistry.contract.Call(opts, &out, "getActiveGenerationReservationsByRange", startIndex, endIndex)

	if err != nil {
		return *new([]OperatorSet), err
	}

	out0 := *abi.ConvertType(out[0], new([]OperatorSet)).(*[]OperatorSet)

	return out0, err

}

// GetActiveGenerationReservationsByRange is a free data retrieval call binding the contract method 0xd9a6729e.
//
// Solidity: function getActiveGenerationReservationsByRange(uint256 startIndex, uint256 endIndex) view returns((address,uint32)[])
func (_CrossChainRegistry *CrossChainRegistrySession) GetActiveGenerationReservationsByRange(startIndex *big.Int, endIndex *big.Int) ([]OperatorSet, error) {
	return _CrossChainRegistry.Contract.GetActiveGenerationReservationsByRange(&_CrossChainRegistry.CallOpts, startIndex, endIndex)
}

// GetActiveGenerationReservationsByRange is a free data retrieval call binding the contract method 0xd9a6729e.
//
// Solidity: function getActiveGenerationReservationsByRange(uint256 startIndex, uint256 endIndex) view returns((address,uint32)[])
func (_CrossChainRegistry *CrossChainRegistryCallerSession) GetActiveGenerationReservationsByRange(startIndex *big.Int, endIndex *big.Int) ([]OperatorSet, error) {
	return _CrossChainRegistry.Contract.GetActiveGenerationReservationsByRange(&_CrossChainRegistry.CallOpts, startIndex, endIndex)
}

// GetOperatorSetConfig is a free data retrieval call binding the contract method 0x21fa7fdc.
//
// Solidity: function getOperatorSetConfig((address,uint32) operatorSet) view returns((address,uint32))
func (_CrossChainRegistry *CrossChainRegistryCaller) GetOperatorSetConfig(opts *bind.CallOpts, operatorSet OperatorSet) (ICrossChainRegistryTypesOperatorSetConfig, error) {
	var out []interface{}
	err := _CrossChainRegistry.contract.Call(opts, &out, "getOperatorSetConfig", operatorSet)

	if err != nil {
		return *new(ICrossChainRegistryTypesOperatorSetConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(ICrossChainRegistryTypesOperatorSetConfig)).(*ICrossChainRegistryTypesOperatorSetConfig)

	return out0, err

}

// GetOperatorSetConfig is a free data retrieval call binding the contract method 0x21fa7fdc.
//
// Solidity: function getOperatorSetConfig((address,uint32) operatorSet) view returns((address,uint32))
func (_CrossChainRegistry *CrossChainRegistrySession) GetOperatorSetConfig(operatorSet OperatorSet) (ICrossChainRegistryTypesOperatorSetConfig, error) {
	return _CrossChainRegistry.Contract.GetOperatorSetConfig(&_CrossChainRegistry.CallOpts, operatorSet)
}

// GetOperatorSetConfig is a free data retrieval call binding the contract method 0x21fa7fdc.
//
// Solidity: function getOperatorSetConfig((address,uint32) operatorSet) view returns((address,uint32))
func (_CrossChainRegistry *CrossChainRegistryCallerSession) GetOperatorSetConfig(operatorSet OperatorSet) (ICrossChainRegistryTypesOperatorSetConfig, error) {
	return _CrossChainRegistry.Contract.GetOperatorSetConfig(&_CrossChainRegistry.CallOpts, operatorSet)
}

// GetOperatorTableCalculator is a free data retrieval call binding the contract method 0x75e4b539.
//
// Solidity: function getOperatorTableCalculator((address,uint32) operatorSet) view returns(address)
func (_CrossChainRegistry *CrossChainRegistryCaller) GetOperatorTableCalculator(opts *bind.CallOpts, operatorSet OperatorSet) (common.Address, error) {
	var out []interface{}
	err := _CrossChainRegistry.contract.Call(opts, &out, "getOperatorTableCalculator", operatorSet)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOperatorTableCalculator is a free data retrieval call binding the contract method 0x75e4b539.
//
// Solidity: function getOperatorTableCalculator((address,uint32) operatorSet) view returns(address)
func (_CrossChainRegistry *CrossChainRegistrySession) GetOperatorTableCalculator(operatorSet OperatorSet) (common.Address, error) {
	return _CrossChainRegistry.Contract.GetOperatorTableCalculator(&_CrossChainRegistry.CallOpts, operatorSet)
}

// GetOperatorTableCalculator is a free data retrieval call binding the contract method 0x75e4b539.
//
// Solidity: function getOperatorTableCalculator((address,uint32) operatorSet) view returns(address)
func (_CrossChainRegistry *CrossChainRegistryCallerSession) GetOperatorTableCalculator(operatorSet OperatorSet) (common.Address, error) {
	return _CrossChainRegistry.Contract.GetOperatorTableCalculator(&_CrossChainRegistry.CallOpts, operatorSet)
}

// GetSupportedChains is a free data retrieval call binding the contract method 0xc4bffe2b.
//
// Solidity: function getSupportedChains() view returns(uint256[], address[])
func (_CrossChainRegistry *CrossChainRegistryCaller) GetSupportedChains(opts *bind.CallOpts) ([]*big.Int, []common.Address, error) {
	var out []interface{}
	err := _CrossChainRegistry.contract.Call(opts, &out, "getSupportedChains")

	if err != nil {
		return *new([]*big.Int), *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	out1 := *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)

	return out0, out1, err

}

// GetSupportedChains is a free data retrieval call binding the contract method 0xc4bffe2b.
//
// Solidity: function getSupportedChains() view returns(uint256[], address[])
func (_CrossChainRegistry *CrossChainRegistrySession) GetSupportedChains() ([]*big.Int, []common.Address, error) {
	return _CrossChainRegistry.Contract.GetSupportedChains(&_CrossChainRegistry.CallOpts)
}

// GetSupportedChains is a free data retrieval call binding the contract method 0xc4bffe2b.
//
// Solidity: function getSupportedChains() view returns(uint256[], address[])
func (_CrossChainRegistry *CrossChainRegistryCallerSession) GetSupportedChains() ([]*big.Int, []common.Address, error) {
	return _CrossChainRegistry.Contract.GetSupportedChains(&_CrossChainRegistry.CallOpts)
}

// GetTableUpdateCadence is a free data retrieval call binding the contract method 0xac505f4b.
//
// Solidity: function getTableUpdateCadence() view returns(uint32)
func (_CrossChainRegistry *CrossChainRegistryCaller) GetTableUpdateCadence(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _CrossChainRegistry.contract.Call(opts, &out, "getTableUpdateCadence")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetTableUpdateCadence is a free data retrieval call binding the contract method 0xac505f4b.
//
// Solidity: function getTableUpdateCadence() view returns(uint32)
func (_CrossChainRegistry *CrossChainRegistrySession) GetTableUpdateCadence() (uint32, error) {
	return _CrossChainRegistry.Contract.GetTableUpdateCadence(&_CrossChainRegistry.CallOpts)
}

// GetTableUpdateCadence is a free data retrieval call binding the contract method 0xac505f4b.
//
// Solidity: function getTableUpdateCadence() view returns(uint32)
func (_CrossChainRegistry *CrossChainRegistryCallerSession) GetTableUpdateCadence() (uint32, error) {
	return _CrossChainRegistry.Contract.GetTableUpdateCadence(&_CrossChainRegistry.CallOpts)
}

// HasActiveGenerationReservation is a free data retrieval call binding the contract method 0x36b200de.
//
// Solidity: function hasActiveGenerationReservation((address,uint32) operatorSet) view returns(bool)
func (_CrossChainRegistry *CrossChainRegistryCaller) HasActiveGenerationReservation(opts *bind.CallOpts, operatorSet OperatorSet) (bool, error) {
	var out []interface{}
	err := _CrossChainRegistry.contract.Call(opts, &out, "hasActiveGenerationReservation", operatorSet)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasActiveGenerationReservation is a free data retrieval call binding the contract method 0x36b200de.
//
// Solidity: function hasActiveGenerationReservation((address,uint32) operatorSet) view returns(bool)
func (_CrossChainRegistry *CrossChainRegistrySession) HasActiveGenerationReservation(operatorSet OperatorSet) (bool, error) {
	return _CrossChainRegistry.Contract.HasActiveGenerationReservation(&_CrossChainRegistry.CallOpts, operatorSet)
}

// HasActiveGenerationReservation is a free data retrieval call binding the contract method 0x36b200de.
//
// Solidity: function hasActiveGenerationReservation((address,uint32) operatorSet) view returns(bool)
func (_CrossChainRegistry *CrossChainRegistryCallerSession) HasActiveGenerationReservation(operatorSet OperatorSet) (bool, error) {
	return _CrossChainRegistry.Contract.HasActiveGenerationReservation(&_CrossChainRegistry.CallOpts, operatorSet)
}

// KeyRegistrar is a free data retrieval call binding the contract method 0x3ec45c7e.
//
// Solidity: function keyRegistrar() view returns(address)
func (_CrossChainRegistry *CrossChainRegistryCaller) KeyRegistrar(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossChainRegistry.contract.Call(opts, &out, "keyRegistrar")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// KeyRegistrar is a free data retrieval call binding the contract method 0x3ec45c7e.
//
// Solidity: function keyRegistrar() view returns(address)
func (_CrossChainRegistry *CrossChainRegistrySession) KeyRegistrar() (common.Address, error) {
	return _CrossChainRegistry.Contract.KeyRegistrar(&_CrossChainRegistry.CallOpts)
}

// KeyRegistrar is a free data retrieval call binding the contract method 0x3ec45c7e.
//
// Solidity: function keyRegistrar() view returns(address)
func (_CrossChainRegistry *CrossChainRegistryCallerSession) KeyRegistrar() (common.Address, error) {
	return _CrossChainRegistry.Contract.KeyRegistrar(&_CrossChainRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CrossChainRegistry *CrossChainRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossChainRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CrossChainRegistry *CrossChainRegistrySession) Owner() (common.Address, error) {
	return _CrossChainRegistry.Contract.Owner(&_CrossChainRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CrossChainRegistry *CrossChainRegistryCallerSession) Owner() (common.Address, error) {
	return _CrossChainRegistry.Contract.Owner(&_CrossChainRegistry.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_CrossChainRegistry *CrossChainRegistryCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _CrossChainRegistry.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_CrossChainRegistry *CrossChainRegistrySession) Paused(index uint8) (bool, error) {
	return _CrossChainRegistry.Contract.Paused(&_CrossChainRegistry.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_CrossChainRegistry *CrossChainRegistryCallerSession) Paused(index uint8) (bool, error) {
	return _CrossChainRegistry.Contract.Paused(&_CrossChainRegistry.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_CrossChainRegistry *CrossChainRegistryCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossChainRegistry.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_CrossChainRegistry *CrossChainRegistrySession) Paused0() (*big.Int, error) {
	return _CrossChainRegistry.Contract.Paused0(&_CrossChainRegistry.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_CrossChainRegistry *CrossChainRegistryCallerSession) Paused0() (*big.Int, error) {
	return _CrossChainRegistry.Contract.Paused0(&_CrossChainRegistry.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_CrossChainRegistry *CrossChainRegistryCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossChainRegistry.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_CrossChainRegistry *CrossChainRegistrySession) PauserRegistry() (common.Address, error) {
	return _CrossChainRegistry.Contract.PauserRegistry(&_CrossChainRegistry.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_CrossChainRegistry *CrossChainRegistryCallerSession) PauserRegistry() (common.Address, error) {
	return _CrossChainRegistry.Contract.PauserRegistry(&_CrossChainRegistry.CallOpts)
}

// PermissionController is a free data retrieval call binding the contract method 0x4657e26a.
//
// Solidity: function permissionController() view returns(address)
func (_CrossChainRegistry *CrossChainRegistryCaller) PermissionController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossChainRegistry.contract.Call(opts, &out, "permissionController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PermissionController is a free data retrieval call binding the contract method 0x4657e26a.
//
// Solidity: function permissionController() view returns(address)
func (_CrossChainRegistry *CrossChainRegistrySession) PermissionController() (common.Address, error) {
	return _CrossChainRegistry.Contract.PermissionController(&_CrossChainRegistry.CallOpts)
}

// PermissionController is a free data retrieval call binding the contract method 0x4657e26a.
//
// Solidity: function permissionController() view returns(address)
func (_CrossChainRegistry *CrossChainRegistryCallerSession) PermissionController() (common.Address, error) {
	return _CrossChainRegistry.Contract.PermissionController(&_CrossChainRegistry.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_CrossChainRegistry *CrossChainRegistryCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CrossChainRegistry.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_CrossChainRegistry *CrossChainRegistrySession) Version() (string, error) {
	return _CrossChainRegistry.Contract.Version(&_CrossChainRegistry.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_CrossChainRegistry *CrossChainRegistryCallerSession) Version() (string, error) {
	return _CrossChainRegistry.Contract.Version(&_CrossChainRegistry.CallOpts)
}

// AddChainIDsToWhitelist is a paid mutator transaction binding the contract method 0x04e98be3.
//
// Solidity: function addChainIDsToWhitelist(uint256[] chainIDs, address[] operatorTableUpdaters) returns()
func (_CrossChainRegistry *CrossChainRegistryTransactor) AddChainIDsToWhitelist(opts *bind.TransactOpts, chainIDs []*big.Int, operatorTableUpdaters []common.Address) (*types.Transaction, error) {
	return _CrossChainRegistry.contract.Transact(opts, "addChainIDsToWhitelist", chainIDs, operatorTableUpdaters)
}

// AddChainIDsToWhitelist is a paid mutator transaction binding the contract method 0x04e98be3.
//
// Solidity: function addChainIDsToWhitelist(uint256[] chainIDs, address[] operatorTableUpdaters) returns()
func (_CrossChainRegistry *CrossChainRegistrySession) AddChainIDsToWhitelist(chainIDs []*big.Int, operatorTableUpdaters []common.Address) (*types.Transaction, error) {
	return _CrossChainRegistry.Contract.AddChainIDsToWhitelist(&_CrossChainRegistry.TransactOpts, chainIDs, operatorTableUpdaters)
}

// AddChainIDsToWhitelist is a paid mutator transaction binding the contract method 0x04e98be3.
//
// Solidity: function addChainIDsToWhitelist(uint256[] chainIDs, address[] operatorTableUpdaters) returns()
func (_CrossChainRegistry *CrossChainRegistryTransactorSession) AddChainIDsToWhitelist(chainIDs []*big.Int, operatorTableUpdaters []common.Address) (*types.Transaction, error) {
	return _CrossChainRegistry.Contract.AddChainIDsToWhitelist(&_CrossChainRegistry.TransactOpts, chainIDs, operatorTableUpdaters)
}

// CreateGenerationReservation is a paid mutator transaction binding the contract method 0xd5044911.
//
// Solidity: function createGenerationReservation((address,uint32) operatorSet, address operatorTableCalculator, (address,uint32) config) returns()
func (_CrossChainRegistry *CrossChainRegistryTransactor) CreateGenerationReservation(opts *bind.TransactOpts, operatorSet OperatorSet, operatorTableCalculator common.Address, config ICrossChainRegistryTypesOperatorSetConfig) (*types.Transaction, error) {
	return _CrossChainRegistry.contract.Transact(opts, "createGenerationReservation", operatorSet, operatorTableCalculator, config)
}

// CreateGenerationReservation is a paid mutator transaction binding the contract method 0xd5044911.
//
// Solidity: function createGenerationReservation((address,uint32) operatorSet, address operatorTableCalculator, (address,uint32) config) returns()
func (_CrossChainRegistry *CrossChainRegistrySession) CreateGenerationReservation(operatorSet OperatorSet, operatorTableCalculator common.Address, config ICrossChainRegistryTypesOperatorSetConfig) (*types.Transaction, error) {
	return _CrossChainRegistry.Contract.CreateGenerationReservation(&_CrossChainRegistry.TransactOpts, operatorSet, operatorTableCalculator, config)
}

// CreateGenerationReservation is a paid mutator transaction binding the contract method 0xd5044911.
//
// Solidity: function createGenerationReservation((address,uint32) operatorSet, address operatorTableCalculator, (address,uint32) config) returns()
func (_CrossChainRegistry *CrossChainRegistryTransactorSession) CreateGenerationReservation(operatorSet OperatorSet, operatorTableCalculator common.Address, config ICrossChainRegistryTypesOperatorSetConfig) (*types.Transaction, error) {
	return _CrossChainRegistry.Contract.CreateGenerationReservation(&_CrossChainRegistry.TransactOpts, operatorSet, operatorTableCalculator, config)
}

// Initialize is a paid mutator transaction binding the contract method 0x0f19aaef.
//
// Solidity: function initialize(address initialOwner, uint32 initialTableUpdateCadence, uint256 initialPausedStatus) returns()
func (_CrossChainRegistry *CrossChainRegistryTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, initialTableUpdateCadence uint32, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _CrossChainRegistry.contract.Transact(opts, "initialize", initialOwner, initialTableUpdateCadence, initialPausedStatus)
}

// Initialize is a paid mutator transaction binding the contract method 0x0f19aaef.
//
// Solidity: function initialize(address initialOwner, uint32 initialTableUpdateCadence, uint256 initialPausedStatus) returns()
func (_CrossChainRegistry *CrossChainRegistrySession) Initialize(initialOwner common.Address, initialTableUpdateCadence uint32, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _CrossChainRegistry.Contract.Initialize(&_CrossChainRegistry.TransactOpts, initialOwner, initialTableUpdateCadence, initialPausedStatus)
}

// Initialize is a paid mutator transaction binding the contract method 0x0f19aaef.
//
// Solidity: function initialize(address initialOwner, uint32 initialTableUpdateCadence, uint256 initialPausedStatus) returns()
func (_CrossChainRegistry *CrossChainRegistryTransactorSession) Initialize(initialOwner common.Address, initialTableUpdateCadence uint32, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _CrossChainRegistry.Contract.Initialize(&_CrossChainRegistry.TransactOpts, initialOwner, initialTableUpdateCadence, initialPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_CrossChainRegistry *CrossChainRegistryTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _CrossChainRegistry.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_CrossChainRegistry *CrossChainRegistrySession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _CrossChainRegistry.Contract.Pause(&_CrossChainRegistry.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_CrossChainRegistry *CrossChainRegistryTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _CrossChainRegistry.Contract.Pause(&_CrossChainRegistry.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_CrossChainRegistry *CrossChainRegistryTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossChainRegistry.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_CrossChainRegistry *CrossChainRegistrySession) PauseAll() (*types.Transaction, error) {
	return _CrossChainRegistry.Contract.PauseAll(&_CrossChainRegistry.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_CrossChainRegistry *CrossChainRegistryTransactorSession) PauseAll() (*types.Transaction, error) {
	return _CrossChainRegistry.Contract.PauseAll(&_CrossChainRegistry.TransactOpts)
}

// RemoveChainIDsFromWhitelist is a paid mutator transaction binding the contract method 0xdfbd9dfd.
//
// Solidity: function removeChainIDsFromWhitelist(uint256[] chainIDs) returns()
func (_CrossChainRegistry *CrossChainRegistryTransactor) RemoveChainIDsFromWhitelist(opts *bind.TransactOpts, chainIDs []*big.Int) (*types.Transaction, error) {
	return _CrossChainRegistry.contract.Transact(opts, "removeChainIDsFromWhitelist", chainIDs)
}

// RemoveChainIDsFromWhitelist is a paid mutator transaction binding the contract method 0xdfbd9dfd.
//
// Solidity: function removeChainIDsFromWhitelist(uint256[] chainIDs) returns()
func (_CrossChainRegistry *CrossChainRegistrySession) RemoveChainIDsFromWhitelist(chainIDs []*big.Int) (*types.Transaction, error) {
	return _CrossChainRegistry.Contract.RemoveChainIDsFromWhitelist(&_CrossChainRegistry.TransactOpts, chainIDs)
}

// RemoveChainIDsFromWhitelist is a paid mutator transaction binding the contract method 0xdfbd9dfd.
//
// Solidity: function removeChainIDsFromWhitelist(uint256[] chainIDs) returns()
func (_CrossChainRegistry *CrossChainRegistryTransactorSession) RemoveChainIDsFromWhitelist(chainIDs []*big.Int) (*types.Transaction, error) {
	return _CrossChainRegistry.Contract.RemoveChainIDsFromWhitelist(&_CrossChainRegistry.TransactOpts, chainIDs)
}

// RemoveGenerationReservation is a paid mutator transaction binding the contract method 0x6c55a37f.
//
// Solidity: function removeGenerationReservation((address,uint32) operatorSet) returns()
func (_CrossChainRegistry *CrossChainRegistryTransactor) RemoveGenerationReservation(opts *bind.TransactOpts, operatorSet OperatorSet) (*types.Transaction, error) {
	return _CrossChainRegistry.contract.Transact(opts, "removeGenerationReservation", operatorSet)
}

// RemoveGenerationReservation is a paid mutator transaction binding the contract method 0x6c55a37f.
//
// Solidity: function removeGenerationReservation((address,uint32) operatorSet) returns()
func (_CrossChainRegistry *CrossChainRegistrySession) RemoveGenerationReservation(operatorSet OperatorSet) (*types.Transaction, error) {
	return _CrossChainRegistry.Contract.RemoveGenerationReservation(&_CrossChainRegistry.TransactOpts, operatorSet)
}

// RemoveGenerationReservation is a paid mutator transaction binding the contract method 0x6c55a37f.
//
// Solidity: function removeGenerationReservation((address,uint32) operatorSet) returns()
func (_CrossChainRegistry *CrossChainRegistryTransactorSession) RemoveGenerationReservation(operatorSet OperatorSet) (*types.Transaction, error) {
	return _CrossChainRegistry.Contract.RemoveGenerationReservation(&_CrossChainRegistry.TransactOpts, operatorSet)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CrossChainRegistry *CrossChainRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossChainRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CrossChainRegistry *CrossChainRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _CrossChainRegistry.Contract.RenounceOwnership(&_CrossChainRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CrossChainRegistry *CrossChainRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CrossChainRegistry.Contract.RenounceOwnership(&_CrossChainRegistry.TransactOpts)
}

// SetOperatorSetConfig is a paid mutator transaction binding the contract method 0x277e1e62.
//
// Solidity: function setOperatorSetConfig((address,uint32) operatorSet, (address,uint32) config) returns()
func (_CrossChainRegistry *CrossChainRegistryTransactor) SetOperatorSetConfig(opts *bind.TransactOpts, operatorSet OperatorSet, config ICrossChainRegistryTypesOperatorSetConfig) (*types.Transaction, error) {
	return _CrossChainRegistry.contract.Transact(opts, "setOperatorSetConfig", operatorSet, config)
}

// SetOperatorSetConfig is a paid mutator transaction binding the contract method 0x277e1e62.
//
// Solidity: function setOperatorSetConfig((address,uint32) operatorSet, (address,uint32) config) returns()
func (_CrossChainRegistry *CrossChainRegistrySession) SetOperatorSetConfig(operatorSet OperatorSet, config ICrossChainRegistryTypesOperatorSetConfig) (*types.Transaction, error) {
	return _CrossChainRegistry.Contract.SetOperatorSetConfig(&_CrossChainRegistry.TransactOpts, operatorSet, config)
}

// SetOperatorSetConfig is a paid mutator transaction binding the contract method 0x277e1e62.
//
// Solidity: function setOperatorSetConfig((address,uint32) operatorSet, (address,uint32) config) returns()
func (_CrossChainRegistry *CrossChainRegistryTransactorSession) SetOperatorSetConfig(operatorSet OperatorSet, config ICrossChainRegistryTypesOperatorSetConfig) (*types.Transaction, error) {
	return _CrossChainRegistry.Contract.SetOperatorSetConfig(&_CrossChainRegistry.TransactOpts, operatorSet, config)
}

// SetOperatorTableCalculator is a paid mutator transaction binding the contract method 0x1ca9142a.
//
// Solidity: function setOperatorTableCalculator((address,uint32) operatorSet, address operatorTableCalculator) returns()
func (_CrossChainRegistry *CrossChainRegistryTransactor) SetOperatorTableCalculator(opts *bind.TransactOpts, operatorSet OperatorSet, operatorTableCalculator common.Address) (*types.Transaction, error) {
	return _CrossChainRegistry.contract.Transact(opts, "setOperatorTableCalculator", operatorSet, operatorTableCalculator)
}

// SetOperatorTableCalculator is a paid mutator transaction binding the contract method 0x1ca9142a.
//
// Solidity: function setOperatorTableCalculator((address,uint32) operatorSet, address operatorTableCalculator) returns()
func (_CrossChainRegistry *CrossChainRegistrySession) SetOperatorTableCalculator(operatorSet OperatorSet, operatorTableCalculator common.Address) (*types.Transaction, error) {
	return _CrossChainRegistry.Contract.SetOperatorTableCalculator(&_CrossChainRegistry.TransactOpts, operatorSet, operatorTableCalculator)
}

// SetOperatorTableCalculator is a paid mutator transaction binding the contract method 0x1ca9142a.
//
// Solidity: function setOperatorTableCalculator((address,uint32) operatorSet, address operatorTableCalculator) returns()
func (_CrossChainRegistry *CrossChainRegistryTransactorSession) SetOperatorTableCalculator(operatorSet OperatorSet, operatorTableCalculator common.Address) (*types.Transaction, error) {
	return _CrossChainRegistry.Contract.SetOperatorTableCalculator(&_CrossChainRegistry.TransactOpts, operatorSet, operatorTableCalculator)
}

// SetTableUpdateCadence is a paid mutator transaction binding the contract method 0xd6db9e25.
//
// Solidity: function setTableUpdateCadence(uint32 tableUpdateCadence) returns()
func (_CrossChainRegistry *CrossChainRegistryTransactor) SetTableUpdateCadence(opts *bind.TransactOpts, tableUpdateCadence uint32) (*types.Transaction, error) {
	return _CrossChainRegistry.contract.Transact(opts, "setTableUpdateCadence", tableUpdateCadence)
}

// SetTableUpdateCadence is a paid mutator transaction binding the contract method 0xd6db9e25.
//
// Solidity: function setTableUpdateCadence(uint32 tableUpdateCadence) returns()
func (_CrossChainRegistry *CrossChainRegistrySession) SetTableUpdateCadence(tableUpdateCadence uint32) (*types.Transaction, error) {
	return _CrossChainRegistry.Contract.SetTableUpdateCadence(&_CrossChainRegistry.TransactOpts, tableUpdateCadence)
}

// SetTableUpdateCadence is a paid mutator transaction binding the contract method 0xd6db9e25.
//
// Solidity: function setTableUpdateCadence(uint32 tableUpdateCadence) returns()
func (_CrossChainRegistry *CrossChainRegistryTransactorSession) SetTableUpdateCadence(tableUpdateCadence uint32) (*types.Transaction, error) {
	return _CrossChainRegistry.Contract.SetTableUpdateCadence(&_CrossChainRegistry.TransactOpts, tableUpdateCadence)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CrossChainRegistry *CrossChainRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CrossChainRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CrossChainRegistry *CrossChainRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CrossChainRegistry.Contract.TransferOwnership(&_CrossChainRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CrossChainRegistry *CrossChainRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CrossChainRegistry.Contract.TransferOwnership(&_CrossChainRegistry.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_CrossChainRegistry *CrossChainRegistryTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _CrossChainRegistry.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_CrossChainRegistry *CrossChainRegistrySession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _CrossChainRegistry.Contract.Unpause(&_CrossChainRegistry.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_CrossChainRegistry *CrossChainRegistryTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _CrossChainRegistry.Contract.Unpause(&_CrossChainRegistry.TransactOpts, newPausedStatus)
}

// CrossChainRegistryChainIDAddedToWhitelistIterator is returned from FilterChainIDAddedToWhitelist and is used to iterate over the raw logs and unpacked data for ChainIDAddedToWhitelist events raised by the CrossChainRegistry contract.
type CrossChainRegistryChainIDAddedToWhitelistIterator struct {
	Event *CrossChainRegistryChainIDAddedToWhitelist // Event containing the contract specifics and raw log

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
func (it *CrossChainRegistryChainIDAddedToWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainRegistryChainIDAddedToWhitelist)
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
		it.Event = new(CrossChainRegistryChainIDAddedToWhitelist)
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
func (it *CrossChainRegistryChainIDAddedToWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossChainRegistryChainIDAddedToWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossChainRegistryChainIDAddedToWhitelist represents a ChainIDAddedToWhitelist event raised by the CrossChainRegistry contract.
type CrossChainRegistryChainIDAddedToWhitelist struct {
	ChainID              *big.Int
	OperatorTableUpdater common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterChainIDAddedToWhitelist is a free log retrieval operation binding the contract event 0x7a0a76d85b582b17996dd7371a407aa7a79b870db8539247fba315c7b6beff62.
//
// Solidity: event ChainIDAddedToWhitelist(uint256 chainID, address operatorTableUpdater)
func (_CrossChainRegistry *CrossChainRegistryFilterer) FilterChainIDAddedToWhitelist(opts *bind.FilterOpts) (*CrossChainRegistryChainIDAddedToWhitelistIterator, error) {

	logs, sub, err := _CrossChainRegistry.contract.FilterLogs(opts, "ChainIDAddedToWhitelist")
	if err != nil {
		return nil, err
	}
	return &CrossChainRegistryChainIDAddedToWhitelistIterator{contract: _CrossChainRegistry.contract, event: "ChainIDAddedToWhitelist", logs: logs, sub: sub}, nil
}

// WatchChainIDAddedToWhitelist is a free log subscription operation binding the contract event 0x7a0a76d85b582b17996dd7371a407aa7a79b870db8539247fba315c7b6beff62.
//
// Solidity: event ChainIDAddedToWhitelist(uint256 chainID, address operatorTableUpdater)
func (_CrossChainRegistry *CrossChainRegistryFilterer) WatchChainIDAddedToWhitelist(opts *bind.WatchOpts, sink chan<- *CrossChainRegistryChainIDAddedToWhitelist) (event.Subscription, error) {

	logs, sub, err := _CrossChainRegistry.contract.WatchLogs(opts, "ChainIDAddedToWhitelist")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossChainRegistryChainIDAddedToWhitelist)
				if err := _CrossChainRegistry.contract.UnpackLog(event, "ChainIDAddedToWhitelist", log); err != nil {
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

// ParseChainIDAddedToWhitelist is a log parse operation binding the contract event 0x7a0a76d85b582b17996dd7371a407aa7a79b870db8539247fba315c7b6beff62.
//
// Solidity: event ChainIDAddedToWhitelist(uint256 chainID, address operatorTableUpdater)
func (_CrossChainRegistry *CrossChainRegistryFilterer) ParseChainIDAddedToWhitelist(log types.Log) (*CrossChainRegistryChainIDAddedToWhitelist, error) {
	event := new(CrossChainRegistryChainIDAddedToWhitelist)
	if err := _CrossChainRegistry.contract.UnpackLog(event, "ChainIDAddedToWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossChainRegistryChainIDRemovedFromWhitelistIterator is returned from FilterChainIDRemovedFromWhitelist and is used to iterate over the raw logs and unpacked data for ChainIDRemovedFromWhitelist events raised by the CrossChainRegistry contract.
type CrossChainRegistryChainIDRemovedFromWhitelistIterator struct {
	Event *CrossChainRegistryChainIDRemovedFromWhitelist // Event containing the contract specifics and raw log

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
func (it *CrossChainRegistryChainIDRemovedFromWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainRegistryChainIDRemovedFromWhitelist)
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
		it.Event = new(CrossChainRegistryChainIDRemovedFromWhitelist)
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
func (it *CrossChainRegistryChainIDRemovedFromWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossChainRegistryChainIDRemovedFromWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossChainRegistryChainIDRemovedFromWhitelist represents a ChainIDRemovedFromWhitelist event raised by the CrossChainRegistry contract.
type CrossChainRegistryChainIDRemovedFromWhitelist struct {
	ChainID *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterChainIDRemovedFromWhitelist is a free log retrieval operation binding the contract event 0x6824d36084ecf2cd819b137cb5d837cc6e73afce1e0e348c9fdecaa81d0341e5.
//
// Solidity: event ChainIDRemovedFromWhitelist(uint256 chainID)
func (_CrossChainRegistry *CrossChainRegistryFilterer) FilterChainIDRemovedFromWhitelist(opts *bind.FilterOpts) (*CrossChainRegistryChainIDRemovedFromWhitelistIterator, error) {

	logs, sub, err := _CrossChainRegistry.contract.FilterLogs(opts, "ChainIDRemovedFromWhitelist")
	if err != nil {
		return nil, err
	}
	return &CrossChainRegistryChainIDRemovedFromWhitelistIterator{contract: _CrossChainRegistry.contract, event: "ChainIDRemovedFromWhitelist", logs: logs, sub: sub}, nil
}

// WatchChainIDRemovedFromWhitelist is a free log subscription operation binding the contract event 0x6824d36084ecf2cd819b137cb5d837cc6e73afce1e0e348c9fdecaa81d0341e5.
//
// Solidity: event ChainIDRemovedFromWhitelist(uint256 chainID)
func (_CrossChainRegistry *CrossChainRegistryFilterer) WatchChainIDRemovedFromWhitelist(opts *bind.WatchOpts, sink chan<- *CrossChainRegistryChainIDRemovedFromWhitelist) (event.Subscription, error) {

	logs, sub, err := _CrossChainRegistry.contract.WatchLogs(opts, "ChainIDRemovedFromWhitelist")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossChainRegistryChainIDRemovedFromWhitelist)
				if err := _CrossChainRegistry.contract.UnpackLog(event, "ChainIDRemovedFromWhitelist", log); err != nil {
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

// ParseChainIDRemovedFromWhitelist is a log parse operation binding the contract event 0x6824d36084ecf2cd819b137cb5d837cc6e73afce1e0e348c9fdecaa81d0341e5.
//
// Solidity: event ChainIDRemovedFromWhitelist(uint256 chainID)
func (_CrossChainRegistry *CrossChainRegistryFilterer) ParseChainIDRemovedFromWhitelist(log types.Log) (*CrossChainRegistryChainIDRemovedFromWhitelist, error) {
	event := new(CrossChainRegistryChainIDRemovedFromWhitelist)
	if err := _CrossChainRegistry.contract.UnpackLog(event, "ChainIDRemovedFromWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossChainRegistryGenerationReservationCreatedIterator is returned from FilterGenerationReservationCreated and is used to iterate over the raw logs and unpacked data for GenerationReservationCreated events raised by the CrossChainRegistry contract.
type CrossChainRegistryGenerationReservationCreatedIterator struct {
	Event *CrossChainRegistryGenerationReservationCreated // Event containing the contract specifics and raw log

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
func (it *CrossChainRegistryGenerationReservationCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainRegistryGenerationReservationCreated)
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
		it.Event = new(CrossChainRegistryGenerationReservationCreated)
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
func (it *CrossChainRegistryGenerationReservationCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossChainRegistryGenerationReservationCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossChainRegistryGenerationReservationCreated represents a GenerationReservationCreated event raised by the CrossChainRegistry contract.
type CrossChainRegistryGenerationReservationCreated struct {
	OperatorSet OperatorSet
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterGenerationReservationCreated is a free log retrieval operation binding the contract event 0x4fb6efec7dd60036ce3a7af8d5c48425019daa0fb61eb471a966a7ac2c6fa6a6.
//
// Solidity: event GenerationReservationCreated((address,uint32) operatorSet)
func (_CrossChainRegistry *CrossChainRegistryFilterer) FilterGenerationReservationCreated(opts *bind.FilterOpts) (*CrossChainRegistryGenerationReservationCreatedIterator, error) {

	logs, sub, err := _CrossChainRegistry.contract.FilterLogs(opts, "GenerationReservationCreated")
	if err != nil {
		return nil, err
	}
	return &CrossChainRegistryGenerationReservationCreatedIterator{contract: _CrossChainRegistry.contract, event: "GenerationReservationCreated", logs: logs, sub: sub}, nil
}

// WatchGenerationReservationCreated is a free log subscription operation binding the contract event 0x4fb6efec7dd60036ce3a7af8d5c48425019daa0fb61eb471a966a7ac2c6fa6a6.
//
// Solidity: event GenerationReservationCreated((address,uint32) operatorSet)
func (_CrossChainRegistry *CrossChainRegistryFilterer) WatchGenerationReservationCreated(opts *bind.WatchOpts, sink chan<- *CrossChainRegistryGenerationReservationCreated) (event.Subscription, error) {

	logs, sub, err := _CrossChainRegistry.contract.WatchLogs(opts, "GenerationReservationCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossChainRegistryGenerationReservationCreated)
				if err := _CrossChainRegistry.contract.UnpackLog(event, "GenerationReservationCreated", log); err != nil {
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

// ParseGenerationReservationCreated is a log parse operation binding the contract event 0x4fb6efec7dd60036ce3a7af8d5c48425019daa0fb61eb471a966a7ac2c6fa6a6.
//
// Solidity: event GenerationReservationCreated((address,uint32) operatorSet)
func (_CrossChainRegistry *CrossChainRegistryFilterer) ParseGenerationReservationCreated(log types.Log) (*CrossChainRegistryGenerationReservationCreated, error) {
	event := new(CrossChainRegistryGenerationReservationCreated)
	if err := _CrossChainRegistry.contract.UnpackLog(event, "GenerationReservationCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossChainRegistryGenerationReservationRemovedIterator is returned from FilterGenerationReservationRemoved and is used to iterate over the raw logs and unpacked data for GenerationReservationRemoved events raised by the CrossChainRegistry contract.
type CrossChainRegistryGenerationReservationRemovedIterator struct {
	Event *CrossChainRegistryGenerationReservationRemoved // Event containing the contract specifics and raw log

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
func (it *CrossChainRegistryGenerationReservationRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainRegistryGenerationReservationRemoved)
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
		it.Event = new(CrossChainRegistryGenerationReservationRemoved)
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
func (it *CrossChainRegistryGenerationReservationRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossChainRegistryGenerationReservationRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossChainRegistryGenerationReservationRemoved represents a GenerationReservationRemoved event raised by the CrossChainRegistry contract.
type CrossChainRegistryGenerationReservationRemoved struct {
	OperatorSet OperatorSet
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterGenerationReservationRemoved is a free log retrieval operation binding the contract event 0x4ffdfdd59e9e1e3c301608788f78dd458e61cb8c045ca92b62a7b484c80824fb.
//
// Solidity: event GenerationReservationRemoved((address,uint32) operatorSet)
func (_CrossChainRegistry *CrossChainRegistryFilterer) FilterGenerationReservationRemoved(opts *bind.FilterOpts) (*CrossChainRegistryGenerationReservationRemovedIterator, error) {

	logs, sub, err := _CrossChainRegistry.contract.FilterLogs(opts, "GenerationReservationRemoved")
	if err != nil {
		return nil, err
	}
	return &CrossChainRegistryGenerationReservationRemovedIterator{contract: _CrossChainRegistry.contract, event: "GenerationReservationRemoved", logs: logs, sub: sub}, nil
}

// WatchGenerationReservationRemoved is a free log subscription operation binding the contract event 0x4ffdfdd59e9e1e3c301608788f78dd458e61cb8c045ca92b62a7b484c80824fb.
//
// Solidity: event GenerationReservationRemoved((address,uint32) operatorSet)
func (_CrossChainRegistry *CrossChainRegistryFilterer) WatchGenerationReservationRemoved(opts *bind.WatchOpts, sink chan<- *CrossChainRegistryGenerationReservationRemoved) (event.Subscription, error) {

	logs, sub, err := _CrossChainRegistry.contract.WatchLogs(opts, "GenerationReservationRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossChainRegistryGenerationReservationRemoved)
				if err := _CrossChainRegistry.contract.UnpackLog(event, "GenerationReservationRemoved", log); err != nil {
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

// ParseGenerationReservationRemoved is a log parse operation binding the contract event 0x4ffdfdd59e9e1e3c301608788f78dd458e61cb8c045ca92b62a7b484c80824fb.
//
// Solidity: event GenerationReservationRemoved((address,uint32) operatorSet)
func (_CrossChainRegistry *CrossChainRegistryFilterer) ParseGenerationReservationRemoved(log types.Log) (*CrossChainRegistryGenerationReservationRemoved, error) {
	event := new(CrossChainRegistryGenerationReservationRemoved)
	if err := _CrossChainRegistry.contract.UnpackLog(event, "GenerationReservationRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossChainRegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the CrossChainRegistry contract.
type CrossChainRegistryInitializedIterator struct {
	Event *CrossChainRegistryInitialized // Event containing the contract specifics and raw log

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
func (it *CrossChainRegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainRegistryInitialized)
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
		it.Event = new(CrossChainRegistryInitialized)
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
func (it *CrossChainRegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossChainRegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossChainRegistryInitialized represents a Initialized event raised by the CrossChainRegistry contract.
type CrossChainRegistryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_CrossChainRegistry *CrossChainRegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*CrossChainRegistryInitializedIterator, error) {

	logs, sub, err := _CrossChainRegistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CrossChainRegistryInitializedIterator{contract: _CrossChainRegistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_CrossChainRegistry *CrossChainRegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CrossChainRegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _CrossChainRegistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossChainRegistryInitialized)
				if err := _CrossChainRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_CrossChainRegistry *CrossChainRegistryFilterer) ParseInitialized(log types.Log) (*CrossChainRegistryInitialized, error) {
	event := new(CrossChainRegistryInitialized)
	if err := _CrossChainRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossChainRegistryOperatorSetConfigRemovedIterator is returned from FilterOperatorSetConfigRemoved and is used to iterate over the raw logs and unpacked data for OperatorSetConfigRemoved events raised by the CrossChainRegistry contract.
type CrossChainRegistryOperatorSetConfigRemovedIterator struct {
	Event *CrossChainRegistryOperatorSetConfigRemoved // Event containing the contract specifics and raw log

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
func (it *CrossChainRegistryOperatorSetConfigRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainRegistryOperatorSetConfigRemoved)
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
		it.Event = new(CrossChainRegistryOperatorSetConfigRemoved)
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
func (it *CrossChainRegistryOperatorSetConfigRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossChainRegistryOperatorSetConfigRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossChainRegistryOperatorSetConfigRemoved represents a OperatorSetConfigRemoved event raised by the CrossChainRegistry contract.
type CrossChainRegistryOperatorSetConfigRemoved struct {
	OperatorSet OperatorSet
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorSetConfigRemoved is a free log retrieval operation binding the contract event 0x210a1118a869246162804e2a7f21ef808ebd93f4be7ed512014fe29a7a8be02e.
//
// Solidity: event OperatorSetConfigRemoved((address,uint32) operatorSet)
func (_CrossChainRegistry *CrossChainRegistryFilterer) FilterOperatorSetConfigRemoved(opts *bind.FilterOpts) (*CrossChainRegistryOperatorSetConfigRemovedIterator, error) {

	logs, sub, err := _CrossChainRegistry.contract.FilterLogs(opts, "OperatorSetConfigRemoved")
	if err != nil {
		return nil, err
	}
	return &CrossChainRegistryOperatorSetConfigRemovedIterator{contract: _CrossChainRegistry.contract, event: "OperatorSetConfigRemoved", logs: logs, sub: sub}, nil
}

// WatchOperatorSetConfigRemoved is a free log subscription operation binding the contract event 0x210a1118a869246162804e2a7f21ef808ebd93f4be7ed512014fe29a7a8be02e.
//
// Solidity: event OperatorSetConfigRemoved((address,uint32) operatorSet)
func (_CrossChainRegistry *CrossChainRegistryFilterer) WatchOperatorSetConfigRemoved(opts *bind.WatchOpts, sink chan<- *CrossChainRegistryOperatorSetConfigRemoved) (event.Subscription, error) {

	logs, sub, err := _CrossChainRegistry.contract.WatchLogs(opts, "OperatorSetConfigRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossChainRegistryOperatorSetConfigRemoved)
				if err := _CrossChainRegistry.contract.UnpackLog(event, "OperatorSetConfigRemoved", log); err != nil {
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

// ParseOperatorSetConfigRemoved is a log parse operation binding the contract event 0x210a1118a869246162804e2a7f21ef808ebd93f4be7ed512014fe29a7a8be02e.
//
// Solidity: event OperatorSetConfigRemoved((address,uint32) operatorSet)
func (_CrossChainRegistry *CrossChainRegistryFilterer) ParseOperatorSetConfigRemoved(log types.Log) (*CrossChainRegistryOperatorSetConfigRemoved, error) {
	event := new(CrossChainRegistryOperatorSetConfigRemoved)
	if err := _CrossChainRegistry.contract.UnpackLog(event, "OperatorSetConfigRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossChainRegistryOperatorSetConfigSetIterator is returned from FilterOperatorSetConfigSet and is used to iterate over the raw logs and unpacked data for OperatorSetConfigSet events raised by the CrossChainRegistry contract.
type CrossChainRegistryOperatorSetConfigSetIterator struct {
	Event *CrossChainRegistryOperatorSetConfigSet // Event containing the contract specifics and raw log

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
func (it *CrossChainRegistryOperatorSetConfigSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainRegistryOperatorSetConfigSet)
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
		it.Event = new(CrossChainRegistryOperatorSetConfigSet)
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
func (it *CrossChainRegistryOperatorSetConfigSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossChainRegistryOperatorSetConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossChainRegistryOperatorSetConfigSet represents a OperatorSetConfigSet event raised by the CrossChainRegistry contract.
type CrossChainRegistryOperatorSetConfigSet struct {
	OperatorSet OperatorSet
	Config      ICrossChainRegistryTypesOperatorSetConfig
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorSetConfigSet is a free log retrieval operation binding the contract event 0x3147846ee526009000671c20380b856a633345691300f82585f90034715cf0e2.
//
// Solidity: event OperatorSetConfigSet((address,uint32) operatorSet, (address,uint32) config)
func (_CrossChainRegistry *CrossChainRegistryFilterer) FilterOperatorSetConfigSet(opts *bind.FilterOpts) (*CrossChainRegistryOperatorSetConfigSetIterator, error) {

	logs, sub, err := _CrossChainRegistry.contract.FilterLogs(opts, "OperatorSetConfigSet")
	if err != nil {
		return nil, err
	}
	return &CrossChainRegistryOperatorSetConfigSetIterator{contract: _CrossChainRegistry.contract, event: "OperatorSetConfigSet", logs: logs, sub: sub}, nil
}

// WatchOperatorSetConfigSet is a free log subscription operation binding the contract event 0x3147846ee526009000671c20380b856a633345691300f82585f90034715cf0e2.
//
// Solidity: event OperatorSetConfigSet((address,uint32) operatorSet, (address,uint32) config)
func (_CrossChainRegistry *CrossChainRegistryFilterer) WatchOperatorSetConfigSet(opts *bind.WatchOpts, sink chan<- *CrossChainRegistryOperatorSetConfigSet) (event.Subscription, error) {

	logs, sub, err := _CrossChainRegistry.contract.WatchLogs(opts, "OperatorSetConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossChainRegistryOperatorSetConfigSet)
				if err := _CrossChainRegistry.contract.UnpackLog(event, "OperatorSetConfigSet", log); err != nil {
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

// ParseOperatorSetConfigSet is a log parse operation binding the contract event 0x3147846ee526009000671c20380b856a633345691300f82585f90034715cf0e2.
//
// Solidity: event OperatorSetConfigSet((address,uint32) operatorSet, (address,uint32) config)
func (_CrossChainRegistry *CrossChainRegistryFilterer) ParseOperatorSetConfigSet(log types.Log) (*CrossChainRegistryOperatorSetConfigSet, error) {
	event := new(CrossChainRegistryOperatorSetConfigSet)
	if err := _CrossChainRegistry.contract.UnpackLog(event, "OperatorSetConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossChainRegistryOperatorTableCalculatorRemovedIterator is returned from FilterOperatorTableCalculatorRemoved and is used to iterate over the raw logs and unpacked data for OperatorTableCalculatorRemoved events raised by the CrossChainRegistry contract.
type CrossChainRegistryOperatorTableCalculatorRemovedIterator struct {
	Event *CrossChainRegistryOperatorTableCalculatorRemoved // Event containing the contract specifics and raw log

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
func (it *CrossChainRegistryOperatorTableCalculatorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainRegistryOperatorTableCalculatorRemoved)
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
		it.Event = new(CrossChainRegistryOperatorTableCalculatorRemoved)
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
func (it *CrossChainRegistryOperatorTableCalculatorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossChainRegistryOperatorTableCalculatorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossChainRegistryOperatorTableCalculatorRemoved represents a OperatorTableCalculatorRemoved event raised by the CrossChainRegistry contract.
type CrossChainRegistryOperatorTableCalculatorRemoved struct {
	OperatorSet OperatorSet
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorTableCalculatorRemoved is a free log retrieval operation binding the contract event 0xd7811913efd5d98fc7ea0d1fdd022b3d31987815360842d05b1d1cf55578d16a.
//
// Solidity: event OperatorTableCalculatorRemoved((address,uint32) operatorSet)
func (_CrossChainRegistry *CrossChainRegistryFilterer) FilterOperatorTableCalculatorRemoved(opts *bind.FilterOpts) (*CrossChainRegistryOperatorTableCalculatorRemovedIterator, error) {

	logs, sub, err := _CrossChainRegistry.contract.FilterLogs(opts, "OperatorTableCalculatorRemoved")
	if err != nil {
		return nil, err
	}
	return &CrossChainRegistryOperatorTableCalculatorRemovedIterator{contract: _CrossChainRegistry.contract, event: "OperatorTableCalculatorRemoved", logs: logs, sub: sub}, nil
}

// WatchOperatorTableCalculatorRemoved is a free log subscription operation binding the contract event 0xd7811913efd5d98fc7ea0d1fdd022b3d31987815360842d05b1d1cf55578d16a.
//
// Solidity: event OperatorTableCalculatorRemoved((address,uint32) operatorSet)
func (_CrossChainRegistry *CrossChainRegistryFilterer) WatchOperatorTableCalculatorRemoved(opts *bind.WatchOpts, sink chan<- *CrossChainRegistryOperatorTableCalculatorRemoved) (event.Subscription, error) {

	logs, sub, err := _CrossChainRegistry.contract.WatchLogs(opts, "OperatorTableCalculatorRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossChainRegistryOperatorTableCalculatorRemoved)
				if err := _CrossChainRegistry.contract.UnpackLog(event, "OperatorTableCalculatorRemoved", log); err != nil {
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

// ParseOperatorTableCalculatorRemoved is a log parse operation binding the contract event 0xd7811913efd5d98fc7ea0d1fdd022b3d31987815360842d05b1d1cf55578d16a.
//
// Solidity: event OperatorTableCalculatorRemoved((address,uint32) operatorSet)
func (_CrossChainRegistry *CrossChainRegistryFilterer) ParseOperatorTableCalculatorRemoved(log types.Log) (*CrossChainRegistryOperatorTableCalculatorRemoved, error) {
	event := new(CrossChainRegistryOperatorTableCalculatorRemoved)
	if err := _CrossChainRegistry.contract.UnpackLog(event, "OperatorTableCalculatorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossChainRegistryOperatorTableCalculatorSetIterator is returned from FilterOperatorTableCalculatorSet and is used to iterate over the raw logs and unpacked data for OperatorTableCalculatorSet events raised by the CrossChainRegistry contract.
type CrossChainRegistryOperatorTableCalculatorSetIterator struct {
	Event *CrossChainRegistryOperatorTableCalculatorSet // Event containing the contract specifics and raw log

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
func (it *CrossChainRegistryOperatorTableCalculatorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainRegistryOperatorTableCalculatorSet)
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
		it.Event = new(CrossChainRegistryOperatorTableCalculatorSet)
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
func (it *CrossChainRegistryOperatorTableCalculatorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossChainRegistryOperatorTableCalculatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossChainRegistryOperatorTableCalculatorSet represents a OperatorTableCalculatorSet event raised by the CrossChainRegistry contract.
type CrossChainRegistryOperatorTableCalculatorSet struct {
	OperatorSet             OperatorSet
	OperatorTableCalculator common.Address
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterOperatorTableCalculatorSet is a free log retrieval operation binding the contract event 0x7f7ccafd92d20fdb39dee184a0dce002a9da420ed0def461f2a027abc9b3f6df.
//
// Solidity: event OperatorTableCalculatorSet((address,uint32) operatorSet, address operatorTableCalculator)
func (_CrossChainRegistry *CrossChainRegistryFilterer) FilterOperatorTableCalculatorSet(opts *bind.FilterOpts) (*CrossChainRegistryOperatorTableCalculatorSetIterator, error) {

	logs, sub, err := _CrossChainRegistry.contract.FilterLogs(opts, "OperatorTableCalculatorSet")
	if err != nil {
		return nil, err
	}
	return &CrossChainRegistryOperatorTableCalculatorSetIterator{contract: _CrossChainRegistry.contract, event: "OperatorTableCalculatorSet", logs: logs, sub: sub}, nil
}

// WatchOperatorTableCalculatorSet is a free log subscription operation binding the contract event 0x7f7ccafd92d20fdb39dee184a0dce002a9da420ed0def461f2a027abc9b3f6df.
//
// Solidity: event OperatorTableCalculatorSet((address,uint32) operatorSet, address operatorTableCalculator)
func (_CrossChainRegistry *CrossChainRegistryFilterer) WatchOperatorTableCalculatorSet(opts *bind.WatchOpts, sink chan<- *CrossChainRegistryOperatorTableCalculatorSet) (event.Subscription, error) {

	logs, sub, err := _CrossChainRegistry.contract.WatchLogs(opts, "OperatorTableCalculatorSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossChainRegistryOperatorTableCalculatorSet)
				if err := _CrossChainRegistry.contract.UnpackLog(event, "OperatorTableCalculatorSet", log); err != nil {
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

// ParseOperatorTableCalculatorSet is a log parse operation binding the contract event 0x7f7ccafd92d20fdb39dee184a0dce002a9da420ed0def461f2a027abc9b3f6df.
//
// Solidity: event OperatorTableCalculatorSet((address,uint32) operatorSet, address operatorTableCalculator)
func (_CrossChainRegistry *CrossChainRegistryFilterer) ParseOperatorTableCalculatorSet(log types.Log) (*CrossChainRegistryOperatorTableCalculatorSet, error) {
	event := new(CrossChainRegistryOperatorTableCalculatorSet)
	if err := _CrossChainRegistry.contract.UnpackLog(event, "OperatorTableCalculatorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossChainRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CrossChainRegistry contract.
type CrossChainRegistryOwnershipTransferredIterator struct {
	Event *CrossChainRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CrossChainRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainRegistryOwnershipTransferred)
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
		it.Event = new(CrossChainRegistryOwnershipTransferred)
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
func (it *CrossChainRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossChainRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossChainRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the CrossChainRegistry contract.
type CrossChainRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CrossChainRegistry *CrossChainRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CrossChainRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CrossChainRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CrossChainRegistryOwnershipTransferredIterator{contract: _CrossChainRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CrossChainRegistry *CrossChainRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CrossChainRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CrossChainRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossChainRegistryOwnershipTransferred)
				if err := _CrossChainRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_CrossChainRegistry *CrossChainRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*CrossChainRegistryOwnershipTransferred, error) {
	event := new(CrossChainRegistryOwnershipTransferred)
	if err := _CrossChainRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossChainRegistryPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the CrossChainRegistry contract.
type CrossChainRegistryPausedIterator struct {
	Event *CrossChainRegistryPaused // Event containing the contract specifics and raw log

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
func (it *CrossChainRegistryPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainRegistryPaused)
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
		it.Event = new(CrossChainRegistryPaused)
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
func (it *CrossChainRegistryPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossChainRegistryPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossChainRegistryPaused represents a Paused event raised by the CrossChainRegistry contract.
type CrossChainRegistryPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_CrossChainRegistry *CrossChainRegistryFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*CrossChainRegistryPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _CrossChainRegistry.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &CrossChainRegistryPausedIterator{contract: _CrossChainRegistry.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_CrossChainRegistry *CrossChainRegistryFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *CrossChainRegistryPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _CrossChainRegistry.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossChainRegistryPaused)
				if err := _CrossChainRegistry.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_CrossChainRegistry *CrossChainRegistryFilterer) ParsePaused(log types.Log) (*CrossChainRegistryPaused, error) {
	event := new(CrossChainRegistryPaused)
	if err := _CrossChainRegistry.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossChainRegistryTableUpdateCadenceSetIterator is returned from FilterTableUpdateCadenceSet and is used to iterate over the raw logs and unpacked data for TableUpdateCadenceSet events raised by the CrossChainRegistry contract.
type CrossChainRegistryTableUpdateCadenceSetIterator struct {
	Event *CrossChainRegistryTableUpdateCadenceSet // Event containing the contract specifics and raw log

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
func (it *CrossChainRegistryTableUpdateCadenceSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainRegistryTableUpdateCadenceSet)
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
		it.Event = new(CrossChainRegistryTableUpdateCadenceSet)
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
func (it *CrossChainRegistryTableUpdateCadenceSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossChainRegistryTableUpdateCadenceSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossChainRegistryTableUpdateCadenceSet represents a TableUpdateCadenceSet event raised by the CrossChainRegistry contract.
type CrossChainRegistryTableUpdateCadenceSet struct {
	TableUpdateCadence uint32
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterTableUpdateCadenceSet is a free log retrieval operation binding the contract event 0x4fbcd0cca70015b33db8af4aa4f2bd6fd6c1efa9460b8e2333f252c1467a6327.
//
// Solidity: event TableUpdateCadenceSet(uint32 tableUpdateCadence)
func (_CrossChainRegistry *CrossChainRegistryFilterer) FilterTableUpdateCadenceSet(opts *bind.FilterOpts) (*CrossChainRegistryTableUpdateCadenceSetIterator, error) {

	logs, sub, err := _CrossChainRegistry.contract.FilterLogs(opts, "TableUpdateCadenceSet")
	if err != nil {
		return nil, err
	}
	return &CrossChainRegistryTableUpdateCadenceSetIterator{contract: _CrossChainRegistry.contract, event: "TableUpdateCadenceSet", logs: logs, sub: sub}, nil
}

// WatchTableUpdateCadenceSet is a free log subscription operation binding the contract event 0x4fbcd0cca70015b33db8af4aa4f2bd6fd6c1efa9460b8e2333f252c1467a6327.
//
// Solidity: event TableUpdateCadenceSet(uint32 tableUpdateCadence)
func (_CrossChainRegistry *CrossChainRegistryFilterer) WatchTableUpdateCadenceSet(opts *bind.WatchOpts, sink chan<- *CrossChainRegistryTableUpdateCadenceSet) (event.Subscription, error) {

	logs, sub, err := _CrossChainRegistry.contract.WatchLogs(opts, "TableUpdateCadenceSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossChainRegistryTableUpdateCadenceSet)
				if err := _CrossChainRegistry.contract.UnpackLog(event, "TableUpdateCadenceSet", log); err != nil {
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

// ParseTableUpdateCadenceSet is a log parse operation binding the contract event 0x4fbcd0cca70015b33db8af4aa4f2bd6fd6c1efa9460b8e2333f252c1467a6327.
//
// Solidity: event TableUpdateCadenceSet(uint32 tableUpdateCadence)
func (_CrossChainRegistry *CrossChainRegistryFilterer) ParseTableUpdateCadenceSet(log types.Log) (*CrossChainRegistryTableUpdateCadenceSet, error) {
	event := new(CrossChainRegistryTableUpdateCadenceSet)
	if err := _CrossChainRegistry.contract.UnpackLog(event, "TableUpdateCadenceSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossChainRegistryUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the CrossChainRegistry contract.
type CrossChainRegistryUnpausedIterator struct {
	Event *CrossChainRegistryUnpaused // Event containing the contract specifics and raw log

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
func (it *CrossChainRegistryUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainRegistryUnpaused)
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
		it.Event = new(CrossChainRegistryUnpaused)
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
func (it *CrossChainRegistryUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossChainRegistryUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossChainRegistryUnpaused represents a Unpaused event raised by the CrossChainRegistry contract.
type CrossChainRegistryUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_CrossChainRegistry *CrossChainRegistryFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*CrossChainRegistryUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _CrossChainRegistry.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &CrossChainRegistryUnpausedIterator{contract: _CrossChainRegistry.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_CrossChainRegistry *CrossChainRegistryFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *CrossChainRegistryUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _CrossChainRegistry.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossChainRegistryUnpaused)
				if err := _CrossChainRegistry.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_CrossChainRegistry *CrossChainRegistryFilterer) ParseUnpaused(log types.Log) (*CrossChainRegistryUnpaused, error) {
	event := new(CrossChainRegistryUnpaused)
	if err := _CrossChainRegistry.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
