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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_strategyManager\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"blacklistTokens\",\"inputs\":[{\"name\":\"tokens\",\"type\":\"address[]\",\"internalType\":\"contractIERC20[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployNewStrategy\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"outputs\":[{\"name\":\"newStrategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployedStrategies\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_initialPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_strategyBeacon\",\"type\":\"address\",\"internalType\":\"contractIBeacon\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isBlacklisted\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeStrategiesFromWhitelist\",\"inputs\":[{\"name\":\"strategiesToRemoveFromWhitelist\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"strategyBeacon\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBeacon\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"strategyManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"whitelistStrategies\",\"inputs\":[{\"name\":\"strategiesToWhitelist\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyBeaconModified\",\"inputs\":[{\"name\":\"previousBeacon\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIBeacon\"},{\"name\":\"newBeacon\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIBeacon\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategySetForToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIERC20\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenBlacklisted\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIERC20\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyBlacklisted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BlacklistedToken\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CurrentlyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNewPausedStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyPauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnpauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategyAlreadyExists\",\"inputs\":[]}]",
	Bin: "0x60c060405234801561001057600080fd5b50604051611a2f380380611a2f83398101604081905261002f91610150565b806001600160a01b038116610057576040516339b190bb60e11b815260040160405180910390fd5b6001600160a01b03908116608052821660a052610072610079565b505061018a565b603354610100900460ff16156100e55760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60335460ff90811614610136576033805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b038116811461014d57600080fd5b50565b6000806040838503121561016357600080fd5b825161016e81610138565b602084015190925061017f81610138565b809150509250929050565b60805160a05161184f6101e060003960008181610148015281816105e90152818161088e015281816109310152610c7901526000818161021c015281816102f90152818161066c0152610b44015261184f6000f3fe608060405234801561001057600080fd5b50600436106101165760003560e01c8063886f1195116100a2578063f0062d9a11610071578063f0062d9a14610275578063f2fde38b14610288578063fabc1cbc1461029b578063fe38b32d146102ae578063fe575a87146102c157600080fd5b8063886f1195146102175780638da5cb5b1461023e578063b768ebc91461024f578063c350a1b51461026257600080fd5b8063595c6a67116100e9578063595c6a67146101b05780635ac86ab7146101b85780635c975abb146101eb5780636b9b6229146101fc578063715018a61461020f57600080fd5b8063136439dd1461011b57806323103c411461013057806339b70e3814610143578063581dfd6514610187575b600080fd5b61012e610129366004610e7a565b6102e4565b005b61012e61013e366004610edf565b6103bb565b61016a7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020015b60405180910390f35b61016a610195366004610f36565b6001602052600090815260409020546001600160a01b031681565b61012e610657565b6101db6101c6366004610f5a565b609954600160ff9092169190911b9081161490565b604051901515815260200161017e565b60995460405190815260200161017e565b61016a61020a366004610f36565b610709565b61012e610900565b61016a7f000000000000000000000000000000000000000000000000000000000000000081565b6066546001600160a01b031661016a565b61012e61025d366004610edf565b610912565b61012e610270366004610f7d565b61099e565b60005461016a906001600160a01b031681565b61012e610296366004610f36565b610ac9565b61012e6102a9366004610e7a565b610b42565b61012e6102bc366004610edf565b610c5a565b6101db6102cf366004610f36565b60026020526000908152604090205460ff1681565b60405163237dfb4760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906346fbf68e90602401602060405180830381865afa158015610348573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061036c9190610fbf565b61038957604051631d77d47760e21b815260040160405180910390fd5b60995481811681146103ae5760405163c61dca5d60e01b815260040160405180910390fd5b6103b782610cb0565b5050565b6103c3610ced565b60008167ffffffffffffffff8111156103de576103de610fe1565b604051908082528060200260200182016040528015610407578160200160208202803683370190505b5090506000805b838110156105c8576002600086868481811061042c5761042c610ff7565b90506020020160208101906104419190610f36565b6001600160a01b0316815260208101919091526040016000205460ff161561047c5760405163f53de75f60e01b815260040160405180910390fd5b60016002600087878581811061049457610494610ff7565b90506020020160208101906104a99190610f36565b6001600160a01b031681526020810191909152604001600020805460ff19169115159190911790557f75519c51f39873ec0e27dd3bbc09549e4865a113f505393fb9eab5898f6418b385858381811061050457610504610ff7565b90506020020160208101906105199190610f36565b6040516001600160a01b03909116815260200160405180910390a160006001600087878581811061054c5761054c610ff7565b90506020020160208101906105619190610f36565b6001600160a01b03908116825260208201929092526040016000205416905080156105bf578084848151811061059957610599610ff7565b6001600160a01b0390921660209283029190910190910152826105bb8161100d565b9350505b5060010161040e565b508082528015610651576040516316bb16b760e31b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063b5d8b5b89061061e908590600401611034565b600060405180830381600087803b15801561063857600080fd5b505af115801561064c573d6000803e3d6000fd5b505050505b50505050565b60405163237dfb4760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906346fbf68e90602401602060405180830381865afa1580156106bb573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106df9190610fbf565b6106fc57604051631d77d47760e21b815260040160405180910390fd5b610707600019610cb0565b565b60995460009081906001908116036107345760405163840a48d560e01b815260040160405180910390fd5b6001600160a01b03831660009081526002602052604090205460ff161561076e5760405163091867bd60e11b815260040160405180910390fd5b6001600160a01b0383811660009081526001602052604090205416156107a75760405163c45546f760e01b815260040160405180910390fd5b60008054604080516001600160a01b0387811660248084019190915283518084039091018152604490920183526020820180516001600160e01b031663189acdbd60e31b17905291519190921691906107ff90610e6d565b61080a929190611080565b604051809103906000f080158015610826573d6000803e3d6000fd5b5090506108338482610d47565b60408051600180825281830190925260009160208083019080368337019050509050818160008151811061086957610869610ff7565b6001600160a01b039283166020918202929092010152604051632ef047f960e11b81527f000000000000000000000000000000000000000000000000000000000000000090911690635de08ff2906108c5908490600401611034565b600060405180830381600087803b1580156108df57600080fd5b505af11580156108f3573d6000803e3d6000fd5b5093979650505050505050565b610908610ced565b6107076000610db2565b61091a610ced565b604051632ef047f960e11b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690635de08ff29061096890859085906004016110dd565b600060405180830381600087803b15801561098257600080fd5b505af1158015610996573d6000803e3d6000fd5b505050505050565b603354610100900460ff16158080156109be5750603354600160ff909116105b806109d85750303b1580156109d8575060335460ff166001145b610a405760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b6033805460ff191660011790558015610a63576033805461ff0019166101001790555b610a6c84610db2565b610a7583610cb0565b610a7e82610e04565b8015610651576033805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a150505050565b610ad1610ced565b6001600160a01b038116610b365760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610a37565b610b3f81610db2565b50565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610ba0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610bc4919061112a565b6001600160a01b0316336001600160a01b031614610bf55760405163794821ff60e01b815260040160405180910390fd5b60995480198219811614610c1c5760405163c61dca5d60e01b815260040160405180910390fd5b609982905560405182815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c9060200160405180910390a25050565b610c62610ced565b6040516316bb16b760e31b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063b5d8b5b89061096890859085906004016110dd565b609981905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a250565b6066546001600160a01b031633146107075760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610a37565b6001600160a01b0382811660008181526001602090815260409182902080546001600160a01b031916948616948517905581519283528201929092527f6852a55230ef089d785bce7ffbf757985de34026df90a87d7b4a6e56f95d251f910160405180910390a15050565b606680546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054604080516001600160a01b03928316815291831660208301527fe21755962a7d7e100b59b9c3e4d4b54085b146313719955efb6a7a25c5c7feee910160405180910390a1600080546001600160a01b0319166001600160a01b0392909216919091179055565b6106d28061114883390190565b600060208284031215610e8c57600080fd5b5035919050565b60008083601f840112610ea557600080fd5b50813567ffffffffffffffff811115610ebd57600080fd5b6020830191508360208260051b8501011115610ed857600080fd5b9250929050565b60008060208385031215610ef257600080fd5b823567ffffffffffffffff811115610f0957600080fd5b610f1585828601610e93565b90969095509350505050565b6001600160a01b0381168114610b3f57600080fd5b600060208284031215610f4857600080fd5b8135610f5381610f21565b9392505050565b600060208284031215610f6c57600080fd5b813560ff81168114610f5357600080fd5b600080600060608486031215610f9257600080fd5b8335610f9d81610f21565b9250602084013591506040840135610fb481610f21565b809150509250925092565b600060208284031215610fd157600080fd5b81518015158114610f5357600080fd5b634e487b7160e01b600052604160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b60006001820161102d57634e487b7160e01b600052601160045260246000fd5b5060010190565b602080825282518282018190526000918401906040840190835b818110156110755783516001600160a01b031683526020938401939092019160010161104e565b509095945050505050565b60018060a01b0383168152604060208201526000825180604084015260005b818110156110bc576020818601810151606086840101520161109f565b506000606082850101526060601f19601f8301168401019150509392505050565b6020808252810182905260008360408301825b8581101561112057823561110381610f21565b6001600160a01b03168252602092830192909101906001016110f0565b5095945050505050565b60006020828403121561113c57600080fd5b8151610f5381610f2156fe60806040526040516106d23803806106d283398101604081905261002291610420565b61002e82826000610035565b5050610550565b61003e836100f6565b6040516001600160a01b038416907f1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e90600090a260008251118061007f5750805b156100f1576100ef836001600160a01b0316635c60da1b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156100c5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906100e991906104e6565b8361027a565b505b505050565b6001600160a01b0381163b6101605760405162461bcd60e51b815260206004820152602560248201527f455243313936373a206e657720626561636f6e206973206e6f74206120636f6e6044820152641d1c9858dd60da1b60648201526084015b60405180910390fd5b6101d4816001600160a01b0316635c60da1b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156101a1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101c591906104e6565b6001600160a01b03163b151590565b6102395760405162461bcd60e51b815260206004820152603060248201527f455243313936373a20626561636f6e20696d706c656d656e746174696f6e206960448201526f1cc81b9bdd08184818dbdb9d1c9858dd60821b6064820152608401610157565b7fa3f0ad74e5423aebfd80d3ef4346578335a9a72aeaee59ff6cb3582b35133d5080546001600160a01b0319166001600160a01b0392909216919091179055565b606061029f83836040518060600160405280602781526020016106ab602791396102a6565b9392505050565b6060600080856001600160a01b0316856040516102c39190610501565b600060405180830381855af49150503d80600081146102fe576040519150601f19603f3d011682016040523d82523d6000602084013e610303565b606091505b5090925090506103158683838761031f565b9695505050505050565b6060831561038e578251600003610387576001600160a01b0385163b6103875760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610157565b5081610398565b61039883836103a0565b949350505050565b8151156103b05781518083602001fd5b8060405162461bcd60e51b8152600401610157919061051d565b80516001600160a01b03811681146103e157600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b60005b838110156104175781810151838201526020016103ff565b50506000910152565b6000806040838503121561043357600080fd5b61043c836103ca565b60208401519092506001600160401b0381111561045857600080fd5b8301601f8101851361046957600080fd5b80516001600160401b03811115610482576104826103e6565b604051601f8201601f19908116603f011681016001600160401b03811182821017156104b0576104b06103e6565b6040528181528282016020018710156104c857600080fd5b6104d98260208301602086016103fc565b8093505050509250929050565b6000602082840312156104f857600080fd5b61029f826103ca565b600082516105138184602087016103fc565b9190910192915050565b602081526000825180602084015261053c8160408501602087016103fc565b601f01601f19169190910160400192915050565b61014c8061055f6000396000f3fe60806040523661001357610011610017565b005b6100115b610027610022610029565b6100c2565b565b600061005c7fa3f0ad74e5423aebfd80d3ef4346578335a9a72aeaee59ff6cb3582b35133d50546001600160a01b031690565b6001600160a01b0316635c60da1b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610099573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906100bd91906100e6565b905090565b3660008037600080366000845af43d6000803e8080156100e1573d6000f35b3d6000fd5b6000602082840312156100f857600080fd5b81516001600160a01b038116811461010f57600080fd5b939250505056fea2646970667358221220c713350530d7743d63f059fac4a4079abefdb272c7fbd383fdcd392f715b496764736f6c634300081b0033416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a264697066735822122085185263c42b8cd0c639217192e04d20161f28d7441ff95a09e2a28ee823b6fa64736f6c634300081b0033",
}

// StrategyFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use StrategyFactoryMetaData.ABI instead.
var StrategyFactoryABI = StrategyFactoryMetaData.ABI

// StrategyFactoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StrategyFactoryMetaData.Bin instead.
var StrategyFactoryBin = StrategyFactoryMetaData.Bin

// DeployStrategyFactory deploys a new Ethereum contract, binding an instance of StrategyFactory to it.
func DeployStrategyFactory(auth *bind.TransactOpts, backend bind.ContractBackend, _strategyManager common.Address, _pauserRegistry common.Address) (common.Address, *types.Transaction, *StrategyFactory, error) {
	parsed, err := StrategyFactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StrategyFactoryBin), backend, _strategyManager, _pauserRegistry)
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

// Initialize is a paid mutator transaction binding the contract method 0xc350a1b5.
//
// Solidity: function initialize(address _initialOwner, uint256 _initialPausedStatus, address _strategyBeacon) returns()
func (_StrategyFactory *StrategyFactoryTransactor) Initialize(opts *bind.TransactOpts, _initialOwner common.Address, _initialPausedStatus *big.Int, _strategyBeacon common.Address) (*types.Transaction, error) {
	return _StrategyFactory.contract.Transact(opts, "initialize", _initialOwner, _initialPausedStatus, _strategyBeacon)
}

// Initialize is a paid mutator transaction binding the contract method 0xc350a1b5.
//
// Solidity: function initialize(address _initialOwner, uint256 _initialPausedStatus, address _strategyBeacon) returns()
func (_StrategyFactory *StrategyFactorySession) Initialize(_initialOwner common.Address, _initialPausedStatus *big.Int, _strategyBeacon common.Address) (*types.Transaction, error) {
	return _StrategyFactory.Contract.Initialize(&_StrategyFactory.TransactOpts, _initialOwner, _initialPausedStatus, _strategyBeacon)
}

// Initialize is a paid mutator transaction binding the contract method 0xc350a1b5.
//
// Solidity: function initialize(address _initialOwner, uint256 _initialPausedStatus, address _strategyBeacon) returns()
func (_StrategyFactory *StrategyFactoryTransactorSession) Initialize(_initialOwner common.Address, _initialPausedStatus *big.Int, _strategyBeacon common.Address) (*types.Transaction, error) {
	return _StrategyFactory.Contract.Initialize(&_StrategyFactory.TransactOpts, _initialOwner, _initialPausedStatus, _strategyBeacon)
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

// WhitelistStrategies is a paid mutator transaction binding the contract method 0xb768ebc9.
//
// Solidity: function whitelistStrategies(address[] strategiesToWhitelist) returns()
func (_StrategyFactory *StrategyFactoryTransactor) WhitelistStrategies(opts *bind.TransactOpts, strategiesToWhitelist []common.Address) (*types.Transaction, error) {
	return _StrategyFactory.contract.Transact(opts, "whitelistStrategies", strategiesToWhitelist)
}

// WhitelistStrategies is a paid mutator transaction binding the contract method 0xb768ebc9.
//
// Solidity: function whitelistStrategies(address[] strategiesToWhitelist) returns()
func (_StrategyFactory *StrategyFactorySession) WhitelistStrategies(strategiesToWhitelist []common.Address) (*types.Transaction, error) {
	return _StrategyFactory.Contract.WhitelistStrategies(&_StrategyFactory.TransactOpts, strategiesToWhitelist)
}

// WhitelistStrategies is a paid mutator transaction binding the contract method 0xb768ebc9.
//
// Solidity: function whitelistStrategies(address[] strategiesToWhitelist) returns()
func (_StrategyFactory *StrategyFactoryTransactorSession) WhitelistStrategies(strategiesToWhitelist []common.Address) (*types.Transaction, error) {
	return _StrategyFactory.Contract.WhitelistStrategies(&_StrategyFactory.TransactOpts, strategiesToWhitelist)
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
