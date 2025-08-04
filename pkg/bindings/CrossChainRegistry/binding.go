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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_allocationManager\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"},{\"name\":\"_keyRegistrar\",\"type\":\"address\",\"internalType\":\"contractIKeyRegistrar\"},{\"name\":\"_permissionController\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"_version\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addChainIDsToWhitelist\",\"inputs\":[{\"name\":\"chainIDs\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"operatorTableUpdaters\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allocationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateOperatorTableBytes\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createGenerationReservation\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operatorTableCalculator\",\"type\":\"address\",\"internalType\":\"contractIOperatorTableCalculator\"},{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structICrossChainRegistryTypes.OperatorSetConfig\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxStalenessPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getActiveGenerationReservations\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorSetConfig\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structICrossChainRegistryTypes.OperatorSetConfig\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxStalenessPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorTableCalculator\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIOperatorTableCalculator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTableUpdateCadence\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialTableUpdateCadence\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"initialPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"keyRegistrar\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIKeyRegistrar\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"permissionController\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeChainIDsFromWhitelist\",\"inputs\":[{\"name\":\"chainIDs\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeGenerationReservation\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOperatorSetConfig\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structICrossChainRegistryTypes.OperatorSetConfig\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxStalenessPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOperatorTableCalculator\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operatorTableCalculator\",\"type\":\"address\",\"internalType\":\"contractIOperatorTableCalculator\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTableUpdateCadence\",\"inputs\":[{\"name\":\"tableUpdateCadence\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"ChainIDAddedToWhitelist\",\"inputs\":[{\"name\":\"chainID\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"operatorTableUpdater\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainIDRemovedFromWhitelist\",\"inputs\":[{\"name\":\"chainID\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GenerationReservationCreated\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GenerationReservationRemoved\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSetConfigRemoved\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSetConfigSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structICrossChainRegistryTypes.OperatorSetConfig\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxStalenessPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorTableCalculatorRemoved\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorTableCalculatorSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operatorTableCalculator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIOperatorTableCalculator\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TableUpdateCadenceSet\",\"inputs\":[{\"name\":\"tableUpdateCadence\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ArrayLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainIDAlreadyWhitelisted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainIDNotWhitelisted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CurrentlyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EmptyChainIDsArray\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"GenerationReservationAlreadyExists\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"GenerationReservationDoesNotExist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidChainId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNewPausedStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperatorSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPermissions\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidStalenessPeriod\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTableUpdateCadence\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"KeyTypeNotSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyPauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnpauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]}]",
	Bin: "0x610120604052348015610010575f5ffd5b5060405161261538038061261583398101604081905261002f916101c1565b80838686856001600160a01b03811661005b576040516339b190bb60e11b815260040160405180910390fd5b6001600160a01b0390811660805291821660a052811660c0521660e05261008181610098565b610100525061008e6100de565b5050505050610319565b5f5f829050601f815111156100cb578260405163305a27a960e01b81526004016100c291906102be565b60405180910390fd5b80516100d6826102f3565b179392505050565b5f54610100900460ff16156101455760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b60648201526084016100c2565b5f5460ff90811614610194575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146101aa575f5ffd5b50565b634e487b7160e01b5f52604160045260245ffd5b5f5f5f5f5f60a086880312156101d5575f5ffd5b85516101e081610196565b60208701519095506101f181610196565b604087015190945061020281610196565b606087015190935061021381610196565b60808701519092506001600160401b0381111561022e575f5ffd5b8601601f8101881361023e575f5ffd5b80516001600160401b03811115610257576102576101ad565b604051601f8201601f19908116603f011681016001600160401b0381118282101715610285576102856101ad565b6040528181528282016020018a101561029c575f5ffd5b8160208401602083015e5f602083830101528093505050509295509295909350565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b80516020808301519190811015610313575f198160200360031b1b821691505b50919050565b60805160a05160c05160e051610100516122806103955f395f610b1601525f8181610293015261162401525f8181610234015281816109cd01526110b901525f81816103a201528181610733015281816108de01528181610ba9015261101001525f818161033c0152818161155801526118d701526122805ff3fe608060405234801561000f575f5ffd5b50600436106101a1575f3560e01c80636c55a37f116100f3578063ca8aa7c711610093578063d6db9e251161006e578063d6db9e25146103ec578063dfbd9dfd146103ff578063f2fde38b14610412578063fabc1cbc14610425575f5ffd5b8063ca8aa7c71461039d578063d09b978b146103c4578063d5044911146103d9575f5ffd5b8063886f1195116100ce578063886f1195146103375780638da5cb5b1461035e578063ac505f4b1461036f578063c4bffe2b14610387575f5ffd5b80636c55a37f14610309578063715018a61461031c57806375e4b53914610324575f5ffd5b80633ec45c7e1161015e57806354fd4d501161013957806354fd4d50146102b5578063595c6a67146102bd5780635ac86ab7146102c55780635c975abb146102f8575f5ffd5b80633ec45c7e1461022f57806341ee6d0e1461026e5780634657e26a1461028e575f5ffd5b806304e98be3146101a55780630f19aaef146101ba578063136439dd146101cd5780631ca9142a146101e057806321fa7fdc146101f3578063277e1e621461021c575b5f5ffd5b6101b86101b3366004611bae565b610438565b005b6101b86101c8366004611c46565b61057c565b6101b86101db366004611c82565b6106a2565b6101b86101ee366004611caf565b6106dc565b610206610201366004611d88565b610827565b6040516102139190611dc0565b60405180910390f35b6101b861022a366004611dce565b610887565b6102567f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b039091168152602001610213565b61028161027c366004611e01565b6109c8565b6040516102139190611e49565b6102567f000000000000000000000000000000000000000000000000000000000000000081565b610281610b0f565b6101b8610b3f565b6102e86102d3366004611e5b565b606654600160ff9092169190911b9081161490565b6040519015158152602001610213565b606654604051908152602001610213565b6101b8610317366004611e01565b610b53565b6101b8610d74565b610256610332366004611d88565b610d85565b6102567f000000000000000000000000000000000000000000000000000000000000000081565b6033546001600160a01b0316610256565b609e5460405163ffffffff9091168152602001610213565b61038f610db1565b604051610213929190611e7b565b6102567f000000000000000000000000000000000000000000000000000000000000000081565b6103cc610ec7565b6040516102139190611f04565b6101b86103e7366004611f51565b610fba565b6101b86103fa366004611f95565b611200565b6101b861040d366004611fae565b611214565b6101b8610420366004611fed565b6112ba565b6101b8610433366004611c82565b611330565b61044061139d565b600361044b816113f7565b83821461046b5760405163512509d360e11b815260040160405180910390fd5b5f5b84811015610574575f86868381811061048857610488612008565b905060200201359050805f036104b157604051633d23e4d160e11b815260040160405180910390fd5b6104e5818686858181106104c7576104c7612008565b90506020020160208101906104dc9190611fed565b609b9190611422565b610502576040516324bf631b60e11b815260040160405180910390fd5b7f7a0a76d85b582b17996dd7371a407aa7a79b870db8539247fba315c7b6beff628186868581811061053657610536612008565b905060200201602081019061054b9190611fed565b604080519283526001600160a01b0390911660208301520160405180910390a15060010161046d565b505050505050565b5f54610100900460ff161580801561059a57505f54600160ff909116105b806105b35750303b1580156105b357505f5460ff166001145b61061b5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b5f805460ff19166001179055801561063c575f805461ff0019166101001790555b61064584611441565b61064e83611492565b61065782611506565b801561069c575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050565b6106aa611543565b60665481811681146106cf5760405163c61dca5d60e01b815260040160405180910390fd5b6106d882611506565b5050565b60016106e7816113f7565b6106f46020840184611fed565b6106fd816115e6565b61071a5760405163932d94f760e01b815260040160405180910390fd5b6040516304c1b8eb60e31b815284906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063260dc7589061076890849060040161204d565b602060405180830381865afa158015610783573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107a7919061205b565b6107c457604051631fb1705560e21b815260040160405180910390fd5b846107f26107df6107da36849003840184611d88565b611696565b5f90815260986020526040902054151590565b61080f57604051634d2baea960e11b815260040160405180910390fd5b61057461082136889003880188611d88565b866116f9565b604080518082019091525f8082526020820152609a5f61084684611696565b815260208082019290925260409081015f208151808301909252546001600160a01b0381168252600160a01b900463ffffffff169181019190915292915050565b6002610892816113f7565b61089f6020840184611fed565b6108a8816115e6565b6108c55760405163932d94f760e01b815260040160405180910390fd5b6040516304c1b8eb60e31b815284906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063260dc7589061091390849060040161204d565b602060405180830381865afa15801561092e573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610952919061205b565b61096f57604051631fb1705560e21b815260040160405180910390fd5b846109856107df6107da36849003840184611d88565b6109a257604051634d2baea960e11b815260040160405180910390fd5b6105746109b436889003880188611d88565b6109c336889003880188611d88565b611773565b6060817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316637cffe48c846040518263ffffffff1660e01b8152600401610a17919061204d565b602060405180830381865afa158015610a32573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a56919061207a565b610a6861020136869003860186611d88565b610a7a61033236879003870187611d88565b6001600160a01b03166341ee6d0e866040518263ffffffff1660e01b8152600401610aa5919061204d565b5f60405180830381865afa158015610abf573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610ae69190810190612098565b604051602001610af99493929190612140565b6040516020818303038152906040529050919050565b6060610b3a7f000000000000000000000000000000000000000000000000000000000000000061183c565b905090565b610b47611543565b610b515f19611506565b565b5f610b5d816113f7565b610b6a6020830183611fed565b610b73816115e6565b610b905760405163932d94f760e01b815260040160405180910390fd5b6040516304c1b8eb60e31b815283906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063260dc75890610bde90849060040161204d565b602060405180830381865afa158015610bf9573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610c1d919061205b565b610c3a57604051631fb1705560e21b815260040160405180910390fd5b83610c506107df6107da36849003840184611d88565b610c6d57604051634d2baea960e11b815260040160405180910390fd5b5f610c806107da36889003880188611d88565b5f818152609960205260409081902080546001600160a01b0319169055519091507fd7811913efd5d98fc7ea0d1fdd022b3d31987815360842d05b1d1cf55578d16a90610cce90889061204d565b60405180910390a15f818152609a60205260409081902080546001600160c01b0319169055517f210a1118a869246162804e2a7f21ef808ebd93f4be7ed512014fe29a7a8be02e90610d2190889061204d565b60405180910390a1610d34609782611879565b507f4ffdfdd59e9e1e3c301608788f78dd458e61cb8c045ca92b62a7b484c80824fb86604051610d64919061204d565b60405180910390a1505050505050565b610d7c61139d565b610b515f611441565b5f60995f610d9284611696565b815260208101919091526040015f20546001600160a01b031692915050565b6060805f610dbf609b611884565b90505f8167ffffffffffffffff811115610ddb57610ddb611ce5565b604051908082528060200260200182016040528015610e04578160200160208202803683370190505b5090505f8267ffffffffffffffff811115610e2157610e21611ce5565b604051908082528060200260200182016040528015610e4a578160200160208202803683370190505b5090505f5b83811015610ebc575f80610e64609b8461188e565b9150915081858481518110610e7b57610e7b612008565b60200260200101818152505080848481518110610e9a57610e9a612008565b6001600160a01b03909216602092830291909101909101525050600101610e4f565b509094909350915050565b60605f610ed460976118ab565b90505f8167ffffffffffffffff811115610ef057610ef0611ce5565b604051908082528060200260200182016040528015610f3457816020015b604080518082019091525f8082526020820152815260200190600190039081610f0e5790505b5090505f5b82811015610fb3575f610f4d6097836118b4565b90505f610f8982604080518082019091525f80825260208201525060408051808201909152606082901c815263ffffffff909116602082015290565b905080848481518110610f9e57610f9e612008565b60209081029190910101525050600101610f39565b5092915050565b5f610fc4816113f7565b610fd16020850185611fed565b610fda816115e6565b610ff75760405163932d94f760e01b815260040160405180910390fd5b6040516304c1b8eb60e31b815285906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063260dc7589061104590849060040161204d565b602060405180830381865afa158015611060573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611084919061205b565b6110a157604051631fb1705560e21b815260040160405180910390fd5b5f604051631f3ff92360e21b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690637cffe48c906110ee908a9060040161204d565b602060405180830381865afa158015611109573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061112d919061207a565b600281111561113e5761113e61212c565b0361115c5760405163e57cacbd60e01b815260040160405180910390fd5b6111796111716107da36899003890189611d88565b6097906118bf565b61119657604051631883461560e01b815260040160405180910390fd5b7f4fb6efec7dd60036ce3a7af8d5c48425019daa0fb61eb471a966a7ac2c6fa6a6866040516111c5919061204d565b60405180910390a16111df61082136889003880188611d88565b6105746111f136889003880188611d88565b6109c336879003870187611d88565b61120861139d565b61121181611492565b50565b61121c61139d565b6003611227816113f7565b5f5b8281101561069c575f84848381811061124457611244612008565b90506020020135905061126181609b6118ca90919063ffffffff16565b61127e5760405163b3f92ba160e01b815260040160405180910390fd5b6040518181527f6824d36084ecf2cd819b137cb5d837cc6e73afce1e0e348c9fdecaa81d0341e59060200160405180910390a150600101611229565b6112c261139d565b6001600160a01b0381166113275760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610612565b61121181611441565b6113386118d5565b6066548019821981161461135f5760405163c61dca5d60e01b815260040160405180910390fd5b606682905560405182815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c9060200160405180910390a25050565b6033546001600160a01b03163314610b515760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610612565b606654600160ff83161b908116036112115760405163840a48d560e01b815260040160405180910390fd5b5f61143784846001600160a01b038516611986565b90505b9392505050565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b5f8163ffffffff16116114b8576040516316d98e1b60e31b815260040160405180910390fd5b609e805463ffffffff191663ffffffff83169081179091556040519081527f4fbcd0cca70015b33db8af4aa4f2bd6fd6c1efa9460b8e2333f252c1467a63279060200160405180910390a150565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a250565b60405163237dfb4760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906346fbf68e90602401602060405180830381865afa1580156115a5573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906115c9919061205b565b610b5157604051631d77d47760e21b815260040160405180910390fd5b604051631beb2b9760e31b81526001600160a01b0382811660048301523360248301523060448301525f80356001600160e01b0319166064840152917f00000000000000000000000000000000000000000000000000000000000000009091169063df595cb8906084016020604051808303815f875af115801561166c573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611690919061205b565b92915050565b5f815f0151826020015163ffffffff166040516020016116e192919060609290921b6bffffffffffffffffffffffff1916825260a01b6001600160a01b031916601482015260200190565b60405160208183030381529060405261169090612198565b8060995f61170685611696565b81526020019081526020015f205f6101000a8154816001600160a01b0302191690836001600160a01b031602179055507f7f7ccafd92d20fdb39dee184a0dce002a9da420ed0def461f2a027abc9b3f6df82826040516117679291906121bb565b60405180910390a15050565b602081015163ffffffff16158061179a5750609e54602082015163ffffffff918216911610155b6117b757604051632e46483160e11b815260040160405180910390fd5b80609a5f6117c485611696565b815260208082019290925260409081015f2083518154949093015163ffffffff16600160a01b026001600160c01b03199094166001600160a01b0390931692909217929092179055517f3147846ee526009000671c20380b856a633345691300f82585f90034715cf0e29061176790849084906121e1565b60605f611848836119a2565b6040805160208082528183019092529192505f91906020820181803683375050509182525060208101929092525090565b5f61143a83836119c9565b5f61169082611aac565b5f80808061189c8686611ab6565b909450925050505b9250929050565b5f611690825490565b5f61143a8383611adf565b5f61143a8383611b05565b5f61143a8383611b51565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611931573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061195591906121fc565b6001600160a01b0316336001600160a01b031614610b515760405163794821ff60e01b815260040160405180910390fd5b5f828152600284016020526040812082905561143784846118bf565b5f60ff8216601f81111561169057604051632cd44ac360e21b815260040160405180910390fd5b5f8181526001830160205260408120548015611aa3575f6119eb600183612217565b85549091505f906119fe90600190612217565b9050818114611a5d575f865f018281548110611a1c57611a1c612008565b905f5260205f200154905080875f018481548110611a3c57611a3c612008565b5f918252602080832090910192909255918252600188019052604090208390555b8554869080611a6e57611a6e612236565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f905560019350505050611690565b5f915050611690565b5f611690826118ab565b5f8080611ac385856118b4565b5f81815260029690960160205260409095205494959350505050565b5f825f018281548110611af457611af4612008565b905f5260205f200154905092915050565b5f818152600183016020526040812054611b4a57508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155611690565b505f611690565b5f818152600283016020526040812081905561143a8383611879565b5f5f83601f840112611b7d575f5ffd5b50813567ffffffffffffffff811115611b94575f5ffd5b6020830191508360208260051b85010111156118a4575f5ffd5b5f5f5f5f60408587031215611bc1575f5ffd5b843567ffffffffffffffff811115611bd7575f5ffd5b611be387828801611b6d565b909550935050602085013567ffffffffffffffff811115611c02575f5ffd5b611c0e87828801611b6d565b95989497509550505050565b6001600160a01b0381168114611211575f5ffd5b803563ffffffff81168114611c41575f5ffd5b919050565b5f5f5f60608486031215611c58575f5ffd5b8335611c6381611c1a565b9250611c7160208501611c2e565b929592945050506040919091013590565b5f60208284031215611c92575f5ffd5b5035919050565b5f60408284031215611ca9575f5ffd5b50919050565b5f5f60608385031215611cc0575f5ffd5b611cca8484611c99565b91506040830135611cda81611c1a565b809150509250929050565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f1916810167ffffffffffffffff81118282101715611d2257611d22611ce5565b604052919050565b5f60408284031215611d3a575f5ffd5b6040805190810167ffffffffffffffff81118282101715611d5d57611d5d611ce5565b6040529050808235611d6e81611c1a565b8152611d7c60208401611c2e565b60208201525092915050565b5f60408284031215611d98575f5ffd5b61143a8383611d2a565b80516001600160a01b0316825260209081015163ffffffff16910152565b604081016116908284611da2565b5f5f60808385031215611ddf575f5ffd5b611de98484611c99565b9150611df88460408501611c99565b90509250929050565b5f60408284031215611e11575f5ffd5b61143a8383611c99565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f61143a6020830184611e1b565b5f60208284031215611e6b575f5ffd5b813560ff8116811461143a575f5ffd5b604080825283519082018190525f9060208501906060840190835b81811015611eb4578351835260209384019390920191600101611e96565b5050838103602080860191909152855180835291810192508501905f5b81811015611ef85782516001600160a01b0316845260209384019390920191600101611ed1565b50919695505050505050565b602080825282518282018190525f918401906040840190835b81811015611f4657611f30838551611da2565b6020939093019260409290920191600101611f1d565b509095945050505050565b5f5f5f60a08486031215611f63575f5ffd5b611f6d8585611c99565b92506040840135611f7d81611c1a565b9150611f8c8560608601611c99565b90509250925092565b5f60208284031215611fa5575f5ffd5b61143a82611c2e565b5f5f60208385031215611fbf575f5ffd5b823567ffffffffffffffff811115611fd5575f5ffd5b611fe185828601611b6d565b90969095509350505050565b5f60208284031215611ffd575f5ffd5b813561143a81611c1a565b634e487b7160e01b5f52603260045260245ffd5b803561202781611c1a565b6001600160a01b0316825263ffffffff61204360208301611c2e565b1660208301525050565b60408101611690828461201c565b5f6020828403121561206b575f5ffd5b8151801515811461143a575f5ffd5b5f6020828403121561208a575f5ffd5b81516003811061143a575f5ffd5b5f602082840312156120a8575f5ffd5b815167ffffffffffffffff8111156120be575f5ffd5b8201601f810184136120ce575f5ffd5b805167ffffffffffffffff8111156120e8576120e8611ce5565b6120fb601f8201601f1916602001611cf9565b81815285602083850101111561210f575f5ffd5b8160208401602083015e5f91810160200191909152949350505050565b634e487b7160e01b5f52602160045260245ffd5b61214a818661201c565b5f6003851061216757634e487b7160e01b5f52602160045260245ffd5b84604083015261217a6060830185611da2565b60c060a083015261218e60c0830184611e1b565b9695505050505050565b80516020808301519190811015611ca9575f1960209190910360031b1b16919050565b606081016121c98285611da2565b6001600160a01b039290921660409190910152919050565b608081016121ef8285611da2565b61143a6040830184611da2565b5f6020828403121561220c575f5ffd5b815161143a81611c1a565b8181038181111561169057634e487b7160e01b5f52601160045260245ffd5b634e487b7160e01b5f52603160045260245ffdfea2646970667358221220cb8c7b4ca2d61cdeb1901d00c02db4686587e36279b37adb63f638a846bb695264736f6c634300081b0033",
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
