// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package StrategyFactory

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

// StrategyFactoryMetaData contains all meta data concerning the StrategyFactory contract.
var StrategyFactoryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_strategyManager\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"blacklistTokens\",\"inputs\":[{\"name\":\"tokens\",\"type\":\"address[]\",\"internalType\":\"contractIERC20[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployNewStrategy\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"outputs\":[{\"name\":\"newStrategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployedStrategies\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"_initialPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_strategyBeacon\",\"type\":\"address\",\"internalType\":\"contractIBeacon\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isBlacklisted\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeStrategiesFromWhitelist\",\"inputs\":[{\"name\":\"strategiesToRemoveFromWhitelist\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setThirdPartyTransfersForbidden\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"strategyBeacon\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBeacon\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"strategyManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"whitelistStrategies\",\"inputs\":[{\"name\":\"strategiesToWhitelist\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"thirdPartyTransfersForbiddenValues\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyBeaconModified\",\"inputs\":[{\"name\":\"previousBeacon\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIBeacon\"},{\"name\":\"newBeacon\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIBeacon\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategySetForToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIERC20\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenBlacklisted\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIERC20\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyBlacklisted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BlacklistedToken\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CurrentlyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNewPausedStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyPauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnpauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategyAlreadyExists\",\"inputs\":[]}]",
	Bin: "0x60a060405234801561001057600080fd5b50604051611dfe380380611dfe83398101604081905261002f91610109565b6001600160a01b03811660805261004461004a565b50610139565b603354610100900460ff16156100b65760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60335460ff90811614610107576033805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b60006020828403121561011b57600080fd5b81516001600160a01b038116811461013257600080fd5b9392505050565b608051611c8761017760003960008181610171015281816106d901528181610776015281816108be01528181610aee0152610e400152611c876000f3fe608060405234801561001057600080fd5b506004361061012c5760003560e01c80636b9b6229116100ad578063f0062d9a11610071578063f0062d9a1461029d578063f2fde38b146102b0578063fabc1cbc146102c3578063fe38b32d146102d6578063fe575a87146102e957600080fd5b80636b9b62291461024b578063715018a61461025e578063886f1195146102665780638da5cb5b14610279578063be2030941461028a57600080fd5b8063581dfd65116100f4578063581dfd65146101c3578063595c6a67146101ec5780635ac86ab7146101f45780635c975abb14610227578063697d54b41461023857600080fd5b806310d67a2f14610131578063136439dd1461014657806323103c411461015957806339b70e381461016c5780634e5a4263146101b0575b600080fd5b61014461013f36600461112e565b61030c565b005b610144610154366004611152565b6103c0565b6101446101673660046111b7565b6104ab565b6101937f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020015b60405180910390f35b6101446101be366004611207565b610747565b6101936101d136600461112e565b6001602052600090815260409020546001600160a01b031681565b6101446107d7565b610217610202366004611240565b609954600160ff9092169190911b9081161490565b60405190151581526020016101a7565b6099546040519081526020016101a7565b610144610246366004611263565b61089f565b61019361025936600461112e565b6108f9565b610144610b61565b609854610193906001600160a01b031681565b6066546001600160a01b0316610193565b6101446102983660046112d4565b610b75565b600054610193906001600160a01b031681565b6101446102be36600461112e565b610ca3565b6101446102d1366004611152565b610d19565b6101446102e43660046111b7565b610e21565b6102176102f736600461112e565b60026020526000908152604090205460ff1681565b609860009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561035f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103839190611327565b6001600160a01b0316336001600160a01b0316146103b45760405163794821ff60e01b815260040160405180910390fd5b6103bd81610e77565b50565b60985460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015610408573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061042c9190611344565b61044957604051631d77d47760e21b815260040160405180910390fd5b6099548181161461046d5760405163c61dca5d60e01b815260040160405180910390fd5b609981905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b6104b3610f07565b60008167ffffffffffffffff8111156104ce576104ce611361565b6040519080825280602002602001820160405280156104f7578160200160208202803683370190505b5090506000805b838110156106b8576002600086868481811061051c5761051c611377565b9050602002016020810190610531919061112e565b6001600160a01b0316815260208101919091526040016000205460ff161561056c5760405163f53de75f60e01b815260040160405180910390fd5b60016002600087878581811061058457610584611377565b9050602002016020810190610599919061112e565b6001600160a01b031681526020810191909152604001600020805460ff19169115159190911790557f75519c51f39873ec0e27dd3bbc09549e4865a113f505393fb9eab5898f6418b38585838181106105f4576105f4611377565b9050602002016020810190610609919061112e565b6040516001600160a01b03909116815260200160405180910390a160006001600087878581811061063c5761063c611377565b9050602002016020810190610651919061112e565b6001600160a01b03908116825260208201929092526040016000205416905080156106af578084848151811061068957610689611377565b6001600160a01b0390921660209283029190910190910152826106ab8161138d565b9350505b506001016104fe565b508082528015610741576040516316bb16b760e31b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063b5d8b5b89061070e9085906004016113f9565b600060405180830381600087803b15801561072857600080fd5b505af115801561073c573d6000803e3d6000fd5b505050505b50505050565b61074f610f07565b604051634e5a426360e01b81526001600160a01b03838116600483015282151560248301527f00000000000000000000000000000000000000000000000000000000000000001690634e5a4263906044015b600060405180830381600087803b1580156107bb57600080fd5b505af11580156107cf573d6000803e3d6000fd5b505050505050565b60985460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa15801561081f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108439190611344565b61086057604051631d77d47760e21b815260040160405180910390fd5b600019609981905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b6108a7610f07565b60405163df5b354760e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063df5b35479061070e90879087908790879060040161144c565b60995460009081906001908116036109245760405163840a48d560e01b815260040160405180910390fd5b6001600160a01b03831660009081526002602052604090205460ff161561095e5760405163091867bd60e11b815260040160405180910390fd5b6001600160a01b0383811660009081526001602052604090205416156109975760405163c45546f760e01b815260040160405180910390fd5b600080546098546040516001600160a01b038781166024830152918216604482015291169063485cc95560e01b9060640160408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909416939093179092529051610a059061110c565b610a109291906114aa565b604051809103906000f080158015610a2c573d6000803e3d6000fd5b509050610a398482610f61565b6040805160018082528183019092526000916020808301908036833750506040805160018082528183019092529293506000929150602080830190803683370190505090508282600081518110610a9257610a92611377565b60200260200101906001600160a01b031690816001600160a01b031681525050600081600081518110610ac757610ac7611377565b9115156020928302919091019091015260405163df5b354760e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063df5b354790610b259085908590600401611507565b600060405180830381600087803b158015610b3f57600080fd5b505af1158015610b53573d6000803e3d6000fd5b509498975050505050505050565b610b69610f07565b610b736000610fcc565b565b603354610100900460ff1615808015610b955750603354600160ff909116105b80610baf5750303b158015610baf575060335460ff166001145b610c175760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b6033805460ff191660011790558015610c3a576033805461ff0019166101001790555b610c4385610fcc565b610c4d848461101e565b610c56826110a3565b8015610c9c576033805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050505050565b610cab610f07565b6001600160a01b038116610d105760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610c0e565b6103bd81610fcc565b609860009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610d6c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d909190611327565b6001600160a01b0316336001600160a01b031614610dc15760405163794821ff60e01b815260040160405180910390fd5b609954198119609954191614610dea5760405163c61dca5d60e01b815260040160405180910390fd5b609981905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c906020016104a0565b610e29610f07565b6040516316bb16b760e31b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063b5d8b5b8906107a19085908590600401611563565b6001600160a01b038116610e9e576040516339b190bb60e11b815260040160405180910390fd5b609854604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1609880546001600160a01b0319166001600160a01b0392909216919091179055565b6066546001600160a01b03163314610b735760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610c0e565b6001600160a01b0382811660008181526001602090815260409182902080546001600160a01b031916948616948517905581519283528201929092527f6852a55230ef089d785bce7ffbf757985de34026df90a87d7b4a6e56f95d251f910160405180910390a15050565b606680546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6098546001600160a01b031615801561103f57506001600160a01b03821615155b61105c576040516339b190bb60e11b815260040160405180910390fd5b609981905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a261109f82610e77565b5050565b600054604080516001600160a01b03928316815291831660208301527fe21755962a7d7e100b59b9c3e4d4b54085b146313719955efb6a7a25c5c7feee910160405180910390a1600080546001600160a01b0319166001600160a01b0392909216919091179055565b6106d28061158083390190565b6001600160a01b03811681146103bd57600080fd5b60006020828403121561114057600080fd5b813561114b81611119565b9392505050565b60006020828403121561116457600080fd5b5035919050565b60008083601f84011261117d57600080fd5b50813567ffffffffffffffff81111561119557600080fd5b6020830191508360208260051b85010111156111b057600080fd5b9250929050565b600080602083850312156111ca57600080fd5b823567ffffffffffffffff8111156111e157600080fd5b6111ed8582860161116b565b90969095509350505050565b80151581146103bd57600080fd5b6000806040838503121561121a57600080fd5b823561122581611119565b91506020830135611235816111f9565b809150509250929050565b60006020828403121561125257600080fd5b813560ff8116811461114b57600080fd5b6000806000806040858703121561127957600080fd5b843567ffffffffffffffff81111561129057600080fd5b61129c8782880161116b565b909550935050602085013567ffffffffffffffff8111156112bc57600080fd5b6112c88782880161116b565b95989497509550505050565b600080600080608085870312156112ea57600080fd5b84356112f581611119565b9350602085013561130581611119565b925060408501359150606085013561131c81611119565b939692955090935050565b60006020828403121561133957600080fd5b815161114b81611119565b60006020828403121561135657600080fd5b815161114b816111f9565b634e487b7160e01b600052604160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b6000600182016113ad57634e487b7160e01b600052601160045260246000fd5b5060010190565b600081518084526020840193506020830160005b828110156113ef5781516001600160a01b03168652602095860195909101906001016113c8565b5093949350505050565b60208152600061114b60208301846113b4565b81835260208301925060008160005b848110156113ef57813561142e81611119565b6001600160a01b03168652602095860195919091019060010161141b565b60408152600061146060408301868861140c565b82810360208085019190915284825285910160005b8581101561149e578235611488816111f9565b1515825260209283019290910190600101611475565b50979650505050505050565b60018060a01b0383168152604060208201526000825180604084015260005b818110156114e657602081860181015160608684010152016114c9565b506000606082850101526060601f19601f8301168401019150509392505050565b60408152600061151a60408301856113b4565b828103602084015280845180835260208301915060208601925060005b818110156115575783511515835260209384019390920191600101611537565b50909695505050505050565b60208152600061157760208301848661140c565b94935050505056fe60806040526040516106d23803806106d283398101604081905261002291610420565b61002e82826000610035565b5050610550565b61003e836100f6565b6040516001600160a01b038416907f1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e90600090a260008251118061007f5750805b156100f1576100ef836001600160a01b0316635c60da1b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156100c5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906100e991906104e6565b8361027a565b505b505050565b6001600160a01b0381163b6101605760405162461bcd60e51b815260206004820152602560248201527f455243313936373a206e657720626561636f6e206973206e6f74206120636f6e6044820152641d1c9858dd60da1b60648201526084015b60405180910390fd5b6101d4816001600160a01b0316635c60da1b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156101a1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101c591906104e6565b6001600160a01b03163b151590565b6102395760405162461bcd60e51b815260206004820152603060248201527f455243313936373a20626561636f6e20696d706c656d656e746174696f6e206960448201526f1cc81b9bdd08184818dbdb9d1c9858dd60821b6064820152608401610157565b7fa3f0ad74e5423aebfd80d3ef4346578335a9a72aeaee59ff6cb3582b35133d5080546001600160a01b0319166001600160a01b0392909216919091179055565b606061029f83836040518060600160405280602781526020016106ab602791396102a6565b9392505050565b6060600080856001600160a01b0316856040516102c39190610501565b600060405180830381855af49150503d80600081146102fe576040519150601f19603f3d011682016040523d82523d6000602084013e610303565b606091505b5090925090506103158683838761031f565b9695505050505050565b6060831561038e578251600003610387576001600160a01b0385163b6103875760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610157565b5081610398565b61039883836103a0565b949350505050565b8151156103b05781518083602001fd5b8060405162461bcd60e51b8152600401610157919061051d565b80516001600160a01b03811681146103e157600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b60005b838110156104175781810151838201526020016103ff565b50506000910152565b6000806040838503121561043357600080fd5b61043c836103ca565b60208401519092506001600160401b0381111561045857600080fd5b8301601f8101851361046957600080fd5b80516001600160401b03811115610482576104826103e6565b604051601f8201601f19908116603f011681016001600160401b03811182821017156104b0576104b06103e6565b6040528181528282016020018710156104c857600080fd5b6104d98260208301602086016103fc565b8093505050509250929050565b6000602082840312156104f857600080fd5b61029f826103ca565b600082516105138184602087016103fc565b9190910192915050565b602081526000825180602084015261053c8160408501602087016103fc565b601f01601f19169190910160400192915050565b61014c8061055f6000396000f3fe60806040523661001357610011610017565b005b6100115b610027610022610029565b6100c2565b565b600061005c7fa3f0ad74e5423aebfd80d3ef4346578335a9a72aeaee59ff6cb3582b35133d50546001600160a01b031690565b6001600160a01b0316635c60da1b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610099573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906100bd91906100e6565b905090565b3660008037600080366000845af43d6000803e8080156100e1573d6000f35b3d6000fd5b6000602082840312156100f857600080fd5b81516001600160a01b038116811461010f57600080fd5b939250505056fea2646970667358221220cb944d3abd4df8c3f6ee344d91680dd11bc4b5cd56dbc90431015960711da23464736f6c634300081b0033416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220c9d8514d0938ade8407b3ea7be5e1cd523f80a5a594143de4e74eb1c5765188264736f6c634300081b0033",
}

// StrategyFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use StrategyFactoryMetaData.ABI instead.
var StrategyFactoryABI = StrategyFactoryMetaData.ABI

// StrategyFactoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StrategyFactoryMetaData.Bin instead.
var StrategyFactoryBin = StrategyFactoryMetaData.Bin

// DeployStrategyFactory deploys a new Ethereum contract, binding an instance of StrategyFactory to it.
func DeployStrategyFactory(auth *bind.TransactOpts, backend bind.ContractBackend, _strategyManager common.Address) (common.Address, *types.Transaction, *StrategyFactory, error) {
	parsed, err := StrategyFactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StrategyFactoryBin), backend, _strategyManager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StrategyFactory{StrategyFactoryCaller: StrategyFactoryCaller{contract: contract}, StrategyFactoryTransactor: StrategyFactoryTransactor{contract: contract}, StrategyFactoryFilterer: StrategyFactoryFilterer{contract: contract}}, nil
}

// StrategyFactory is an auto generated Go binding around an Ethereum contract.
type StrategyFactory struct {
	StrategyFactoryCaller     // Read-only binding to the contract
	StrategyFactoryTransactor // Write-only binding to the contract
	StrategyFactoryFilterer   // Log filterer for contract events
}

// StrategyFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type StrategyFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrategyFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StrategyFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrategyFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StrategyFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrategyFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StrategyFactorySession struct {
	Contract     *StrategyFactory  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StrategyFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StrategyFactoryCallerSession struct {
	Contract *StrategyFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// StrategyFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StrategyFactoryTransactorSession struct {
	Contract     *StrategyFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// StrategyFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type StrategyFactoryRaw struct {
	Contract *StrategyFactory // Generic contract binding to access the raw methods on
}

// StrategyFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StrategyFactoryCallerRaw struct {
	Contract *StrategyFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// StrategyFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StrategyFactoryTransactorRaw struct {
	Contract *StrategyFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStrategyFactory creates a new instance of StrategyFactory, bound to a specific deployed contract.
func NewStrategyFactory(address common.Address, backend bind.ContractBackend) (*StrategyFactory, error) {
	contract, err := bindStrategyFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StrategyFactory{StrategyFactoryCaller: StrategyFactoryCaller{contract: contract}, StrategyFactoryTransactor: StrategyFactoryTransactor{contract: contract}, StrategyFactoryFilterer: StrategyFactoryFilterer{contract: contract}}, nil
}

// NewStrategyFactoryCaller creates a new read-only instance of StrategyFactory, bound to a specific deployed contract.
func NewStrategyFactoryCaller(address common.Address, caller bind.ContractCaller) (*StrategyFactoryCaller, error) {
	contract, err := bindStrategyFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StrategyFactoryCaller{contract: contract}, nil
}

// NewStrategyFactoryTransactor creates a new write-only instance of StrategyFactory, bound to a specific deployed contract.
func NewStrategyFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*StrategyFactoryTransactor, error) {
	contract, err := bindStrategyFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StrategyFactoryTransactor{contract: contract}, nil
}

// NewStrategyFactoryFilterer creates a new log filterer instance of StrategyFactory, bound to a specific deployed contract.
func NewStrategyFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*StrategyFactoryFilterer, error) {
	contract, err := bindStrategyFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StrategyFactoryFilterer{contract: contract}, nil
}

// bindStrategyFactory binds a generic wrapper to an already deployed contract.
func bindStrategyFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StrategyFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StrategyFactory *StrategyFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StrategyFactory.Contract.StrategyFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StrategyFactory *StrategyFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrategyFactory.Contract.StrategyFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StrategyFactory *StrategyFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StrategyFactory.Contract.StrategyFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StrategyFactory *StrategyFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StrategyFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StrategyFactory *StrategyFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrategyFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StrategyFactory *StrategyFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StrategyFactory.Contract.contract.Transact(opts, method, params...)
}

// DeployedStrategies is a free data retrieval call binding the contract method 0x581dfd65.
//
// Solidity: function deployedStrategies(address ) view returns(address)
func (_StrategyFactory *StrategyFactoryCaller) DeployedStrategies(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _StrategyFactory.contract.Call(opts, &out, "deployedStrategies", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DeployedStrategies is a free data retrieval call binding the contract method 0x581dfd65.
//
// Solidity: function deployedStrategies(address ) view returns(address)
func (_StrategyFactory *StrategyFactorySession) DeployedStrategies(arg0 common.Address) (common.Address, error) {
	return _StrategyFactory.Contract.DeployedStrategies(&_StrategyFactory.CallOpts, arg0)
}

// DeployedStrategies is a free data retrieval call binding the contract method 0x581dfd65.
//
// Solidity: function deployedStrategies(address ) view returns(address)
func (_StrategyFactory *StrategyFactoryCallerSession) DeployedStrategies(arg0 common.Address) (common.Address, error) {
	return _StrategyFactory.Contract.DeployedStrategies(&_StrategyFactory.CallOpts, arg0)
}

// IsBlacklisted is a free data retrieval call binding the contract method 0xfe575a87.
//
// Solidity: function isBlacklisted(address ) view returns(bool)
func (_StrategyFactory *StrategyFactoryCaller) IsBlacklisted(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _StrategyFactory.contract.Call(opts, &out, "isBlacklisted", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBlacklisted is a free data retrieval call binding the contract method 0xfe575a87.
//
// Solidity: function isBlacklisted(address ) view returns(bool)
func (_StrategyFactory *StrategyFactorySession) IsBlacklisted(arg0 common.Address) (bool, error) {
	return _StrategyFactory.Contract.IsBlacklisted(&_StrategyFactory.CallOpts, arg0)
}

// IsBlacklisted is a free data retrieval call binding the contract method 0xfe575a87.
//
// Solidity: function isBlacklisted(address ) view returns(bool)
func (_StrategyFactory *StrategyFactoryCallerSession) IsBlacklisted(arg0 common.Address) (bool, error) {
	return _StrategyFactory.Contract.IsBlacklisted(&_StrategyFactory.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StrategyFactory *StrategyFactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyFactory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StrategyFactory *StrategyFactorySession) Owner() (common.Address, error) {
	return _StrategyFactory.Contract.Owner(&_StrategyFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StrategyFactory *StrategyFactoryCallerSession) Owner() (common.Address, error) {
	return _StrategyFactory.Contract.Owner(&_StrategyFactory.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_StrategyFactory *StrategyFactoryCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _StrategyFactory.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_StrategyFactory *StrategyFactorySession) Paused(index uint8) (bool, error) {
	return _StrategyFactory.Contract.Paused(&_StrategyFactory.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_StrategyFactory *StrategyFactoryCallerSession) Paused(index uint8) (bool, error) {
	return _StrategyFactory.Contract.Paused(&_StrategyFactory.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_StrategyFactory *StrategyFactoryCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrategyFactory.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_StrategyFactory *StrategyFactorySession) Paused0() (*big.Int, error) {
	return _StrategyFactory.Contract.Paused0(&_StrategyFactory.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_StrategyFactory *StrategyFactoryCallerSession) Paused0() (*big.Int, error) {
	return _StrategyFactory.Contract.Paused0(&_StrategyFactory.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_StrategyFactory *StrategyFactoryCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyFactory.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_StrategyFactory *StrategyFactorySession) PauserRegistry() (common.Address, error) {
	return _StrategyFactory.Contract.PauserRegistry(&_StrategyFactory.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_StrategyFactory *StrategyFactoryCallerSession) PauserRegistry() (common.Address, error) {
	return _StrategyFactory.Contract.PauserRegistry(&_StrategyFactory.CallOpts)
}

// StrategyBeacon is a free data retrieval call binding the contract method 0xf0062d9a.
//
// Solidity: function strategyBeacon() view returns(address)
func (_StrategyFactory *StrategyFactoryCaller) StrategyBeacon(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyFactory.contract.Call(opts, &out, "strategyBeacon")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StrategyBeacon is a free data retrieval call binding the contract method 0xf0062d9a.
//
// Solidity: function strategyBeacon() view returns(address)
func (_StrategyFactory *StrategyFactorySession) StrategyBeacon() (common.Address, error) {
	return _StrategyFactory.Contract.StrategyBeacon(&_StrategyFactory.CallOpts)
}

// StrategyBeacon is a free data retrieval call binding the contract method 0xf0062d9a.
//
// Solidity: function strategyBeacon() view returns(address)
func (_StrategyFactory *StrategyFactoryCallerSession) StrategyBeacon() (common.Address, error) {
	return _StrategyFactory.Contract.StrategyBeacon(&_StrategyFactory.CallOpts)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_StrategyFactory *StrategyFactoryCaller) StrategyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyFactory.contract.Call(opts, &out, "strategyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_StrategyFactory *StrategyFactorySession) StrategyManager() (common.Address, error) {
	return _StrategyFactory.Contract.StrategyManager(&_StrategyFactory.CallOpts)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_StrategyFactory *StrategyFactoryCallerSession) StrategyManager() (common.Address, error) {
	return _StrategyFactory.Contract.StrategyManager(&_StrategyFactory.CallOpts)
}

// BlacklistTokens is a paid mutator transaction binding the contract method 0x23103c41.
//
// Solidity: function blacklistTokens(address[] tokens) returns()
func (_StrategyFactory *StrategyFactoryTransactor) BlacklistTokens(opts *bind.TransactOpts, tokens []common.Address) (*types.Transaction, error) {
	return _StrategyFactory.contract.Transact(opts, "blacklistTokens", tokens)
}

// BlacklistTokens is a paid mutator transaction binding the contract method 0x23103c41.
//
// Solidity: function blacklistTokens(address[] tokens) returns()
func (_StrategyFactory *StrategyFactorySession) BlacklistTokens(tokens []common.Address) (*types.Transaction, error) {
	return _StrategyFactory.Contract.BlacklistTokens(&_StrategyFactory.TransactOpts, tokens)
}

// BlacklistTokens is a paid mutator transaction binding the contract method 0x23103c41.
//
// Solidity: function blacklistTokens(address[] tokens) returns()
func (_StrategyFactory *StrategyFactoryTransactorSession) BlacklistTokens(tokens []common.Address) (*types.Transaction, error) {
	return _StrategyFactory.Contract.BlacklistTokens(&_StrategyFactory.TransactOpts, tokens)
}

// DeployNewStrategy is a paid mutator transaction binding the contract method 0x6b9b6229.
//
// Solidity: function deployNewStrategy(address token) returns(address newStrategy)
func (_StrategyFactory *StrategyFactoryTransactor) DeployNewStrategy(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _StrategyFactory.contract.Transact(opts, "deployNewStrategy", token)
}

// DeployNewStrategy is a paid mutator transaction binding the contract method 0x6b9b6229.
//
// Solidity: function deployNewStrategy(address token) returns(address newStrategy)
func (_StrategyFactory *StrategyFactorySession) DeployNewStrategy(token common.Address) (*types.Transaction, error) {
	return _StrategyFactory.Contract.DeployNewStrategy(&_StrategyFactory.TransactOpts, token)
}

// DeployNewStrategy is a paid mutator transaction binding the contract method 0x6b9b6229.
//
// Solidity: function deployNewStrategy(address token) returns(address newStrategy)
func (_StrategyFactory *StrategyFactoryTransactorSession) DeployNewStrategy(token common.Address) (*types.Transaction, error) {
	return _StrategyFactory.Contract.DeployNewStrategy(&_StrategyFactory.TransactOpts, token)
}

// Initialize is a paid mutator transaction binding the contract method 0xbe203094.
//
// Solidity: function initialize(address _initialOwner, address _pauserRegistry, uint256 _initialPausedStatus, address _strategyBeacon) returns()
func (_StrategyFactory *StrategyFactoryTransactor) Initialize(opts *bind.TransactOpts, _initialOwner common.Address, _pauserRegistry common.Address, _initialPausedStatus *big.Int, _strategyBeacon common.Address) (*types.Transaction, error) {
	return _StrategyFactory.contract.Transact(opts, "initialize", _initialOwner, _pauserRegistry, _initialPausedStatus, _strategyBeacon)
}

// Initialize is a paid mutator transaction binding the contract method 0xbe203094.
//
// Solidity: function initialize(address _initialOwner, address _pauserRegistry, uint256 _initialPausedStatus, address _strategyBeacon) returns()
func (_StrategyFactory *StrategyFactorySession) Initialize(_initialOwner common.Address, _pauserRegistry common.Address, _initialPausedStatus *big.Int, _strategyBeacon common.Address) (*types.Transaction, error) {
	return _StrategyFactory.Contract.Initialize(&_StrategyFactory.TransactOpts, _initialOwner, _pauserRegistry, _initialPausedStatus, _strategyBeacon)
}

// Initialize is a paid mutator transaction binding the contract method 0xbe203094.
//
// Solidity: function initialize(address _initialOwner, address _pauserRegistry, uint256 _initialPausedStatus, address _strategyBeacon) returns()
func (_StrategyFactory *StrategyFactoryTransactorSession) Initialize(_initialOwner common.Address, _pauserRegistry common.Address, _initialPausedStatus *big.Int, _strategyBeacon common.Address) (*types.Transaction, error) {
	return _StrategyFactory.Contract.Initialize(&_StrategyFactory.TransactOpts, _initialOwner, _pauserRegistry, _initialPausedStatus, _strategyBeacon)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_StrategyFactory *StrategyFactoryTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _StrategyFactory.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_StrategyFactory *StrategyFactorySession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _StrategyFactory.Contract.Pause(&_StrategyFactory.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_StrategyFactory *StrategyFactoryTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _StrategyFactory.Contract.Pause(&_StrategyFactory.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_StrategyFactory *StrategyFactoryTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrategyFactory.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_StrategyFactory *StrategyFactorySession) PauseAll() (*types.Transaction, error) {
	return _StrategyFactory.Contract.PauseAll(&_StrategyFactory.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_StrategyFactory *StrategyFactoryTransactorSession) PauseAll() (*types.Transaction, error) {
	return _StrategyFactory.Contract.PauseAll(&_StrategyFactory.TransactOpts)
}

// RemoveStrategiesFromWhitelist is a paid mutator transaction binding the contract method 0xfe38b32d.
//
// Solidity: function removeStrategiesFromWhitelist(address[] strategiesToRemoveFromWhitelist) returns()
func (_StrategyFactory *StrategyFactoryTransactor) RemoveStrategiesFromWhitelist(opts *bind.TransactOpts, strategiesToRemoveFromWhitelist []common.Address) (*types.Transaction, error) {
	return _StrategyFactory.contract.Transact(opts, "removeStrategiesFromWhitelist", strategiesToRemoveFromWhitelist)
}

// RemoveStrategiesFromWhitelist is a paid mutator transaction binding the contract method 0xfe38b32d.
//
// Solidity: function removeStrategiesFromWhitelist(address[] strategiesToRemoveFromWhitelist) returns()
func (_StrategyFactory *StrategyFactorySession) RemoveStrategiesFromWhitelist(strategiesToRemoveFromWhitelist []common.Address) (*types.Transaction, error) {
	return _StrategyFactory.Contract.RemoveStrategiesFromWhitelist(&_StrategyFactory.TransactOpts, strategiesToRemoveFromWhitelist)
}

// RemoveStrategiesFromWhitelist is a paid mutator transaction binding the contract method 0xfe38b32d.
//
// Solidity: function removeStrategiesFromWhitelist(address[] strategiesToRemoveFromWhitelist) returns()
func (_StrategyFactory *StrategyFactoryTransactorSession) RemoveStrategiesFromWhitelist(strategiesToRemoveFromWhitelist []common.Address) (*types.Transaction, error) {
	return _StrategyFactory.Contract.RemoveStrategiesFromWhitelist(&_StrategyFactory.TransactOpts, strategiesToRemoveFromWhitelist)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StrategyFactory *StrategyFactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrategyFactory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StrategyFactory *StrategyFactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _StrategyFactory.Contract.RenounceOwnership(&_StrategyFactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StrategyFactory *StrategyFactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _StrategyFactory.Contract.RenounceOwnership(&_StrategyFactory.TransactOpts)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_StrategyFactory *StrategyFactoryTransactor) SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error) {
	return _StrategyFactory.contract.Transact(opts, "setPauserRegistry", newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_StrategyFactory *StrategyFactorySession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _StrategyFactory.Contract.SetPauserRegistry(&_StrategyFactory.TransactOpts, newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_StrategyFactory *StrategyFactoryTransactorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _StrategyFactory.Contract.SetPauserRegistry(&_StrategyFactory.TransactOpts, newPauserRegistry)
}

// SetThirdPartyTransfersForbidden is a paid mutator transaction binding the contract method 0x4e5a4263.
//
// Solidity: function setThirdPartyTransfersForbidden(address strategy, bool value) returns()
func (_StrategyFactory *StrategyFactoryTransactor) SetThirdPartyTransfersForbidden(opts *bind.TransactOpts, strategy common.Address, value bool) (*types.Transaction, error) {
	return _StrategyFactory.contract.Transact(opts, "setThirdPartyTransfersForbidden", strategy, value)
}

// SetThirdPartyTransfersForbidden is a paid mutator transaction binding the contract method 0x4e5a4263.
//
// Solidity: function setThirdPartyTransfersForbidden(address strategy, bool value) returns()
func (_StrategyFactory *StrategyFactorySession) SetThirdPartyTransfersForbidden(strategy common.Address, value bool) (*types.Transaction, error) {
	return _StrategyFactory.Contract.SetThirdPartyTransfersForbidden(&_StrategyFactory.TransactOpts, strategy, value)
}

// SetThirdPartyTransfersForbidden is a paid mutator transaction binding the contract method 0x4e5a4263.
//
// Solidity: function setThirdPartyTransfersForbidden(address strategy, bool value) returns()
func (_StrategyFactory *StrategyFactoryTransactorSession) SetThirdPartyTransfersForbidden(strategy common.Address, value bool) (*types.Transaction, error) {
	return _StrategyFactory.Contract.SetThirdPartyTransfersForbidden(&_StrategyFactory.TransactOpts, strategy, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StrategyFactory *StrategyFactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _StrategyFactory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StrategyFactory *StrategyFactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StrategyFactory.Contract.TransferOwnership(&_StrategyFactory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StrategyFactory *StrategyFactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StrategyFactory.Contract.TransferOwnership(&_StrategyFactory.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_StrategyFactory *StrategyFactoryTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _StrategyFactory.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_StrategyFactory *StrategyFactorySession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _StrategyFactory.Contract.Unpause(&_StrategyFactory.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_StrategyFactory *StrategyFactoryTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _StrategyFactory.Contract.Unpause(&_StrategyFactory.TransactOpts, newPausedStatus)
}

// WhitelistStrategies is a paid mutator transaction binding the contract method 0x697d54b4.
//
// Solidity: function whitelistStrategies(address[] strategiesToWhitelist, bool[] thirdPartyTransfersForbiddenValues) returns()
func (_StrategyFactory *StrategyFactoryTransactor) WhitelistStrategies(opts *bind.TransactOpts, strategiesToWhitelist []common.Address, thirdPartyTransfersForbiddenValues []bool) (*types.Transaction, error) {
	return _StrategyFactory.contract.Transact(opts, "whitelistStrategies", strategiesToWhitelist, thirdPartyTransfersForbiddenValues)
}

// WhitelistStrategies is a paid mutator transaction binding the contract method 0x697d54b4.
//
// Solidity: function whitelistStrategies(address[] strategiesToWhitelist, bool[] thirdPartyTransfersForbiddenValues) returns()
func (_StrategyFactory *StrategyFactorySession) WhitelistStrategies(strategiesToWhitelist []common.Address, thirdPartyTransfersForbiddenValues []bool) (*types.Transaction, error) {
	return _StrategyFactory.Contract.WhitelistStrategies(&_StrategyFactory.TransactOpts, strategiesToWhitelist, thirdPartyTransfersForbiddenValues)
}

// WhitelistStrategies is a paid mutator transaction binding the contract method 0x697d54b4.
//
// Solidity: function whitelistStrategies(address[] strategiesToWhitelist, bool[] thirdPartyTransfersForbiddenValues) returns()
func (_StrategyFactory *StrategyFactoryTransactorSession) WhitelistStrategies(strategiesToWhitelist []common.Address, thirdPartyTransfersForbiddenValues []bool) (*types.Transaction, error) {
	return _StrategyFactory.Contract.WhitelistStrategies(&_StrategyFactory.TransactOpts, strategiesToWhitelist, thirdPartyTransfersForbiddenValues)
}

// StrategyFactoryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the StrategyFactory contract.
type StrategyFactoryInitializedIterator struct {
	Event *StrategyFactoryInitialized // Event containing the contract specifics and raw log

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
func (it *StrategyFactoryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyFactoryInitialized)
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
		it.Event = new(StrategyFactoryInitialized)
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
func (it *StrategyFactoryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyFactoryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyFactoryInitialized represents a Initialized event raised by the StrategyFactory contract.
type StrategyFactoryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StrategyFactory *StrategyFactoryFilterer) FilterInitialized(opts *bind.FilterOpts) (*StrategyFactoryInitializedIterator, error) {

	logs, sub, err := _StrategyFactory.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StrategyFactoryInitializedIterator{contract: _StrategyFactory.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StrategyFactory *StrategyFactoryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StrategyFactoryInitialized) (event.Subscription, error) {

	logs, sub, err := _StrategyFactory.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyFactoryInitialized)
				if err := _StrategyFactory.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_StrategyFactory *StrategyFactoryFilterer) ParseInitialized(log types.Log) (*StrategyFactoryInitialized, error) {
	event := new(StrategyFactoryInitialized)
	if err := _StrategyFactory.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyFactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the StrategyFactory contract.
type StrategyFactoryOwnershipTransferredIterator struct {
	Event *StrategyFactoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StrategyFactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyFactoryOwnershipTransferred)
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
		it.Event = new(StrategyFactoryOwnershipTransferred)
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
func (it *StrategyFactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyFactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyFactoryOwnershipTransferred represents a OwnershipTransferred event raised by the StrategyFactory contract.
type StrategyFactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StrategyFactory *StrategyFactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StrategyFactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StrategyFactory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StrategyFactoryOwnershipTransferredIterator{contract: _StrategyFactory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StrategyFactory *StrategyFactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StrategyFactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StrategyFactory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyFactoryOwnershipTransferred)
				if err := _StrategyFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_StrategyFactory *StrategyFactoryFilterer) ParseOwnershipTransferred(log types.Log) (*StrategyFactoryOwnershipTransferred, error) {
	event := new(StrategyFactoryOwnershipTransferred)
	if err := _StrategyFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyFactoryPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the StrategyFactory contract.
type StrategyFactoryPausedIterator struct {
	Event *StrategyFactoryPaused // Event containing the contract specifics and raw log

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
func (it *StrategyFactoryPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyFactoryPaused)
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
		it.Event = new(StrategyFactoryPaused)
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
func (it *StrategyFactoryPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyFactoryPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyFactoryPaused represents a Paused event raised by the StrategyFactory contract.
type StrategyFactoryPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_StrategyFactory *StrategyFactoryFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*StrategyFactoryPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StrategyFactory.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &StrategyFactoryPausedIterator{contract: _StrategyFactory.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_StrategyFactory *StrategyFactoryFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *StrategyFactoryPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StrategyFactory.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyFactoryPaused)
				if err := _StrategyFactory.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_StrategyFactory *StrategyFactoryFilterer) ParsePaused(log types.Log) (*StrategyFactoryPaused, error) {
	event := new(StrategyFactoryPaused)
	if err := _StrategyFactory.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyFactoryPauserRegistrySetIterator is returned from FilterPauserRegistrySet and is used to iterate over the raw logs and unpacked data for PauserRegistrySet events raised by the StrategyFactory contract.
type StrategyFactoryPauserRegistrySetIterator struct {
	Event *StrategyFactoryPauserRegistrySet // Event containing the contract specifics and raw log

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
func (it *StrategyFactoryPauserRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyFactoryPauserRegistrySet)
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
		it.Event = new(StrategyFactoryPauserRegistrySet)
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
func (it *StrategyFactoryPauserRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyFactoryPauserRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyFactoryPauserRegistrySet represents a PauserRegistrySet event raised by the StrategyFactory contract.
type StrategyFactoryPauserRegistrySet struct {
	PauserRegistry    common.Address
	NewPauserRegistry common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPauserRegistrySet is a free log retrieval operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_StrategyFactory *StrategyFactoryFilterer) FilterPauserRegistrySet(opts *bind.FilterOpts) (*StrategyFactoryPauserRegistrySetIterator, error) {

	logs, sub, err := _StrategyFactory.contract.FilterLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return &StrategyFactoryPauserRegistrySetIterator{contract: _StrategyFactory.contract, event: "PauserRegistrySet", logs: logs, sub: sub}, nil
}

// WatchPauserRegistrySet is a free log subscription operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_StrategyFactory *StrategyFactoryFilterer) WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *StrategyFactoryPauserRegistrySet) (event.Subscription, error) {

	logs, sub, err := _StrategyFactory.contract.WatchLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyFactoryPauserRegistrySet)
				if err := _StrategyFactory.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
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
func (_StrategyFactory *StrategyFactoryFilterer) ParsePauserRegistrySet(log types.Log) (*StrategyFactoryPauserRegistrySet, error) {
	event := new(StrategyFactoryPauserRegistrySet)
	if err := _StrategyFactory.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyFactoryStrategyBeaconModifiedIterator is returned from FilterStrategyBeaconModified and is used to iterate over the raw logs and unpacked data for StrategyBeaconModified events raised by the StrategyFactory contract.
type StrategyFactoryStrategyBeaconModifiedIterator struct {
	Event *StrategyFactoryStrategyBeaconModified // Event containing the contract specifics and raw log

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
func (it *StrategyFactoryStrategyBeaconModifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyFactoryStrategyBeaconModified)
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
		it.Event = new(StrategyFactoryStrategyBeaconModified)
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
func (it *StrategyFactoryStrategyBeaconModifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyFactoryStrategyBeaconModifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyFactoryStrategyBeaconModified represents a StrategyBeaconModified event raised by the StrategyFactory contract.
type StrategyFactoryStrategyBeaconModified struct {
	PreviousBeacon common.Address
	NewBeacon      common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterStrategyBeaconModified is a free log retrieval operation binding the contract event 0xe21755962a7d7e100b59b9c3e4d4b54085b146313719955efb6a7a25c5c7feee.
//
// Solidity: event StrategyBeaconModified(address previousBeacon, address newBeacon)
func (_StrategyFactory *StrategyFactoryFilterer) FilterStrategyBeaconModified(opts *bind.FilterOpts) (*StrategyFactoryStrategyBeaconModifiedIterator, error) {

	logs, sub, err := _StrategyFactory.contract.FilterLogs(opts, "StrategyBeaconModified")
	if err != nil {
		return nil, err
	}
	return &StrategyFactoryStrategyBeaconModifiedIterator{contract: _StrategyFactory.contract, event: "StrategyBeaconModified", logs: logs, sub: sub}, nil
}

// WatchStrategyBeaconModified is a free log subscription operation binding the contract event 0xe21755962a7d7e100b59b9c3e4d4b54085b146313719955efb6a7a25c5c7feee.
//
// Solidity: event StrategyBeaconModified(address previousBeacon, address newBeacon)
func (_StrategyFactory *StrategyFactoryFilterer) WatchStrategyBeaconModified(opts *bind.WatchOpts, sink chan<- *StrategyFactoryStrategyBeaconModified) (event.Subscription, error) {

	logs, sub, err := _StrategyFactory.contract.WatchLogs(opts, "StrategyBeaconModified")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyFactoryStrategyBeaconModified)
				if err := _StrategyFactory.contract.UnpackLog(event, "StrategyBeaconModified", log); err != nil {
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

// ParseStrategyBeaconModified is a log parse operation binding the contract event 0xe21755962a7d7e100b59b9c3e4d4b54085b146313719955efb6a7a25c5c7feee.
//
// Solidity: event StrategyBeaconModified(address previousBeacon, address newBeacon)
func (_StrategyFactory *StrategyFactoryFilterer) ParseStrategyBeaconModified(log types.Log) (*StrategyFactoryStrategyBeaconModified, error) {
	event := new(StrategyFactoryStrategyBeaconModified)
	if err := _StrategyFactory.contract.UnpackLog(event, "StrategyBeaconModified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyFactoryStrategySetForTokenIterator is returned from FilterStrategySetForToken and is used to iterate over the raw logs and unpacked data for StrategySetForToken events raised by the StrategyFactory contract.
type StrategyFactoryStrategySetForTokenIterator struct {
	Event *StrategyFactoryStrategySetForToken // Event containing the contract specifics and raw log

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
func (it *StrategyFactoryStrategySetForTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyFactoryStrategySetForToken)
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
		it.Event = new(StrategyFactoryStrategySetForToken)
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
func (it *StrategyFactoryStrategySetForTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyFactoryStrategySetForTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyFactoryStrategySetForToken represents a StrategySetForToken event raised by the StrategyFactory contract.
type StrategyFactoryStrategySetForToken struct {
	Token    common.Address
	Strategy common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStrategySetForToken is a free log retrieval operation binding the contract event 0x6852a55230ef089d785bce7ffbf757985de34026df90a87d7b4a6e56f95d251f.
//
// Solidity: event StrategySetForToken(address token, address strategy)
func (_StrategyFactory *StrategyFactoryFilterer) FilterStrategySetForToken(opts *bind.FilterOpts) (*StrategyFactoryStrategySetForTokenIterator, error) {

	logs, sub, err := _StrategyFactory.contract.FilterLogs(opts, "StrategySetForToken")
	if err != nil {
		return nil, err
	}
	return &StrategyFactoryStrategySetForTokenIterator{contract: _StrategyFactory.contract, event: "StrategySetForToken", logs: logs, sub: sub}, nil
}

// WatchStrategySetForToken is a free log subscription operation binding the contract event 0x6852a55230ef089d785bce7ffbf757985de34026df90a87d7b4a6e56f95d251f.
//
// Solidity: event StrategySetForToken(address token, address strategy)
func (_StrategyFactory *StrategyFactoryFilterer) WatchStrategySetForToken(opts *bind.WatchOpts, sink chan<- *StrategyFactoryStrategySetForToken) (event.Subscription, error) {

	logs, sub, err := _StrategyFactory.contract.WatchLogs(opts, "StrategySetForToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyFactoryStrategySetForToken)
				if err := _StrategyFactory.contract.UnpackLog(event, "StrategySetForToken", log); err != nil {
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

// ParseStrategySetForToken is a log parse operation binding the contract event 0x6852a55230ef089d785bce7ffbf757985de34026df90a87d7b4a6e56f95d251f.
//
// Solidity: event StrategySetForToken(address token, address strategy)
func (_StrategyFactory *StrategyFactoryFilterer) ParseStrategySetForToken(log types.Log) (*StrategyFactoryStrategySetForToken, error) {
	event := new(StrategyFactoryStrategySetForToken)
	if err := _StrategyFactory.contract.UnpackLog(event, "StrategySetForToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyFactoryTokenBlacklistedIterator is returned from FilterTokenBlacklisted and is used to iterate over the raw logs and unpacked data for TokenBlacklisted events raised by the StrategyFactory contract.
type StrategyFactoryTokenBlacklistedIterator struct {
	Event *StrategyFactoryTokenBlacklisted // Event containing the contract specifics and raw log

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
func (it *StrategyFactoryTokenBlacklistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyFactoryTokenBlacklisted)
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
		it.Event = new(StrategyFactoryTokenBlacklisted)
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
func (it *StrategyFactoryTokenBlacklistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyFactoryTokenBlacklistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyFactoryTokenBlacklisted represents a TokenBlacklisted event raised by the StrategyFactory contract.
type StrategyFactoryTokenBlacklisted struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenBlacklisted is a free log retrieval operation binding the contract event 0x75519c51f39873ec0e27dd3bbc09549e4865a113f505393fb9eab5898f6418b3.
//
// Solidity: event TokenBlacklisted(address token)
func (_StrategyFactory *StrategyFactoryFilterer) FilterTokenBlacklisted(opts *bind.FilterOpts) (*StrategyFactoryTokenBlacklistedIterator, error) {

	logs, sub, err := _StrategyFactory.contract.FilterLogs(opts, "TokenBlacklisted")
	if err != nil {
		return nil, err
	}
	return &StrategyFactoryTokenBlacklistedIterator{contract: _StrategyFactory.contract, event: "TokenBlacklisted", logs: logs, sub: sub}, nil
}

// WatchTokenBlacklisted is a free log subscription operation binding the contract event 0x75519c51f39873ec0e27dd3bbc09549e4865a113f505393fb9eab5898f6418b3.
//
// Solidity: event TokenBlacklisted(address token)
func (_StrategyFactory *StrategyFactoryFilterer) WatchTokenBlacklisted(opts *bind.WatchOpts, sink chan<- *StrategyFactoryTokenBlacklisted) (event.Subscription, error) {

	logs, sub, err := _StrategyFactory.contract.WatchLogs(opts, "TokenBlacklisted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyFactoryTokenBlacklisted)
				if err := _StrategyFactory.contract.UnpackLog(event, "TokenBlacklisted", log); err != nil {
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

// ParseTokenBlacklisted is a log parse operation binding the contract event 0x75519c51f39873ec0e27dd3bbc09549e4865a113f505393fb9eab5898f6418b3.
//
// Solidity: event TokenBlacklisted(address token)
func (_StrategyFactory *StrategyFactoryFilterer) ParseTokenBlacklisted(log types.Log) (*StrategyFactoryTokenBlacklisted, error) {
	event := new(StrategyFactoryTokenBlacklisted)
	if err := _StrategyFactory.contract.UnpackLog(event, "TokenBlacklisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyFactoryUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the StrategyFactory contract.
type StrategyFactoryUnpausedIterator struct {
	Event *StrategyFactoryUnpaused // Event containing the contract specifics and raw log

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
func (it *StrategyFactoryUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyFactoryUnpaused)
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
		it.Event = new(StrategyFactoryUnpaused)
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
func (it *StrategyFactoryUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyFactoryUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyFactoryUnpaused represents a Unpaused event raised by the StrategyFactory contract.
type StrategyFactoryUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_StrategyFactory *StrategyFactoryFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*StrategyFactoryUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StrategyFactory.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &StrategyFactoryUnpausedIterator{contract: _StrategyFactory.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_StrategyFactory *StrategyFactoryFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *StrategyFactoryUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StrategyFactory.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyFactoryUnpaused)
				if err := _StrategyFactory.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_StrategyFactory *StrategyFactoryFilterer) ParseUnpaused(log types.Log) (*StrategyFactoryUnpaused, error) {
	event := new(StrategyFactoryUnpaused)
	if err := _StrategyFactory.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
