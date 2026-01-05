// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ProtocolRegistry

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

// IProtocolRegistryTypesDeploymentConfig is an auto generated low-level Go binding around an user-defined struct.
type IProtocolRegistryTypesDeploymentConfig struct {
	Pausable   bool
	Deprecated bool
}

// ProtocolRegistryMetaData contains all meta data concerning the ProtocolRegistry contract.
var ProtocolRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"configure\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structIProtocolRegistryTypes.DeploymentConfig\",\"components\":[{\"name\":\"pausable\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"deprecated\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAddress\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllDeployments\",\"inputs\":[],\"outputs\":[{\"name\":\"names\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"addresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"configs\",\"type\":\"tuple[]\",\"internalType\":\"structIProtocolRegistryTypes.DeploymentConfig[]\",\"components\":[{\"name\":\"pausable\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"deprecated\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDeployment\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structIProtocolRegistryTypes.DeploymentConfig\",\"components\":[{\"name\":\"pausable\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"deprecated\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleMember\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleMemberCount\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pauserMultisig\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"majorVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ship\",\"inputs\":[{\"name\":\"addresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"configs\",\"type\":\"tuple[]\",\"internalType\":\"structIProtocolRegistryTypes.DeploymentConfig[]\",\"components\":[{\"name\":\"pausable\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"deprecated\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"names\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"semanticVersion\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalDeployments\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"DeploymentConfigDeleted\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DeploymentConfigured\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIProtocolRegistryTypes.DeploymentConfig\",\"components\":[{\"name\":\"pausable\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"deprecated\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DeploymentShipped\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIProtocolRegistryTypes.DeploymentConfig\",\"components\":[{\"name\":\"pausable\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"deprecated\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SemanticVersionUpdated\",\"inputs\":[{\"name\":\"previousSemanticVersion\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"semanticVersion\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ArrayLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DeploymentNotShipped\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]}]",
	Bin: "0x6080604052348015600e575f5ffd5b5060156019565b60d3565b5f54610100900460ff161560835760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff9081161460d1575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b611b1a806100e05f395ff3fe608060405234801561000f575f5ffd5b5060043610610127575f3560e01c8063a217fddf116100a9578063d34663891161006e578063d346638914610286578063d547741f14610299578063e63ab1e9146102ac578063fb35b4e4146102b7578063fc434a8a146102bf575f5ffd5b8063a217fddf14610230578063a8091d9714610237578063af05a5c514610258578063bf40fac114610260578063ca15c87314610273575f5ffd5b806354fd4d50116100ef57806354fd4d50146101be578063595c6a67146101d35780638eec00b8146101db5780639010d07c146101f257806391d148541461021d575f5ffd5b806301ffc9a71461012b578063248a9ca3146101535780632f2ff15d1461018357806336568abe14610198578063485cc955146101ab575b5f5ffd5b61013e61013936600461144d565b6102d2565b60405190151581526020015b60405180910390f35b610175610161366004611474565b5f9081526065602052604090206001015490565b60405190815260200161014a565b6101966101913660046114a6565b6102fc565b005b6101966101a63660046114a6565b610325565b6101966101b93660046114d0565b6103a8565b6101c66104c5565b60405161014a9190611526565b6101966104d7565b6101e36105b7565b60405161014a9392919061158b565b610205610200366004611648565b61078a565b6040516001600160a01b03909116815260200161014a565b61013e61022b3660046114a6565b6107a8565b6101755f81565b61024a6102453660046116a5565b6107d2565b60405161014a9291906116e3565b6101c661087e565b61020561026e3660046116a5565b6108d5565b610175610281366004611474565b61091a565b61019661029436600461170e565b610930565b6101966102a73660046114a6565b6109cc565b610175600160f81b81565b6101756109f0565b6101966102cd36600461178a565b6109fb565b5f6001600160e01b03198216635a05180f60e01b14806102f657506102f682610b12565b92915050565b5f8281526065602052604090206001015461031681610b46565b6103208383610b53565b505050565b6001600160a01b038116331461039a5760405162461bcd60e51b815260206004820152602f60248201527f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560448201526e103937b632b9903337b91039b2b63360891b60648201526084015b60405180910390fd5b6103a48282610b74565b5050565b5f54610100900460ff16158080156103c657505f54600160ff909116105b806103df5750303b1580156103df57505f5460ff166001145b6104425760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610391565b5f805460ff191660011790558015610463575f805461ff0019166101001790555b61046d5f84610b53565b61047b600160f81b83610b53565b8015610320575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498906020015b60405180910390a1505050565b60606104d260c954610b95565b905090565b600160f81b6104e581610b46565b5f6104ee6109f0565b90505f5b81811015610320575f61050660ca83610bd2565b6001600160a01b0381165f90815260cd602090815260409182902082518084019093525460ff80821615801580865261010090930490911615159284019290925292945090925061055957508060200151155b156105ad57816001600160a01b031663595c6a676040518163ffffffff1660e01b81526004015f604051808303815f87803b158015610596575f5ffd5b505af11580156105a8573d5f5f3e3d5ffd5b505050505b50506001016104f2565b60608060605f6105c56109f0565b9050806001600160401b038111156105df576105df611885565b60405190808252806020026020018201604052801561061257816020015b60608152602001906001900390816105fd5790505b509350806001600160401b0381111561062d5761062d611885565b604051908082528060200260200182016040528015610656578160200160208202803683370190505b509250806001600160401b0381111561067157610671611885565b6040519080825280602002602001820160405280156106b557816020015b604080518082019091525f808252602082015281526020019060019003908161068f5790505b5091505f5b81811015610783575f806106cf60ca84610bd2565b90925090506106dd82610b95565b8784815181106106ef576106ef611899565b60200260200101819052508086848151811061070d5761070d611899565b6001600160a01b039283166020918202929092018101919091529082165f90815260cd825260409081902081518083019092525460ff8082161515835261010090910416151591810191909152855186908590811061076e5761076e611899565b602090810291909101015250506001016106ba565b5050909192565b5f8281526097602052604081206107a19083610bed565b9392505050565b5f9182526065602090815260408084206001600160a01b0393909316845291905290205460ff1690565b604080518082019091525f808252602082018190529061083761082f61082c86868080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250610bf892505050565b90565b60ca90610c35565b6001600160a01b0381165f90815260cd602090815260409182902082518084019093525460ff808216151584526101009091041615159082015290925090505b9250929050565b60605f61088c60c954610b95565b9050805f815181106108a0576108a0611899565b016020908101516040516001600160f81b03199091169181019190915260210160405160208183030381529060405291505090565b5f6107a161082f61082c85858080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250610bf892505050565b5f8181526097602052604081206102f690610c40565b5f61093a81610b46565b61094383610c49565b6109605760405163732e48b560e11b815260040160405180910390fd5b6001600160a01b0383165f90815260cd60205260409020829061098382826118ba565b905050826001600160a01b03167f6c5879ec82f910f6d12844857cfb8eb474dcecc9aa5b8257c7a77dcb42990e9d836040516109bf9190611900565b60405180910390a2505050565b5f828152606560205260409020600101546109e681610b46565b6103208383610b74565b5f6104d260ca610ca6565b5f610a0581610b46565b8786148015610a1357508584145b610a305760405163512509d360e11b815260040160405180910390fd5b610a3a8383610cb0565b5f5b88811015610b06575f8a8a83818110610a5757610a57611899565b9050602002016020810190610a6c9190611930565b6001600160a01b031603610a93576040516339b190bb60e11b815260040160405180910390fd5b610afe8a8a83818110610aa857610aa8611899565b9050602002016020810190610abd9190611930565b898984818110610acf57610acf611899565b905060400201888885818110610ae757610ae7611899565b9050602002810190610af99190611949565b610d33565b600101610a3c565b50505050505050505050565b5f6001600160e01b03198216637965db0b60e01b14806102f657506301ffc9a760e01b6001600160e01b03198316146102f6565b610b508133610e70565b50565b610b5d8282610ec9565b5f8281526097602052604090206103209082610f4e565b610b7e8282610f62565b5f8281526097602052604090206103209082610fc8565b60605f610ba183610fdc565b6040805160208082528183019092529192505f91906020820181803683375050509182525060208101929092525090565b5f808080610be08686611003565b9097909650945050505050565b5f6107a1838361102c565b5f5f829050601f81511115610c22578260405163305a27a960e01b81526004016103919190611526565b8051610c2d8261198b565b179392505050565b5f6107a18383611052565b5f6102f6825490565b5f5f610c5560ca610ca6565b90505f5b81811015610c9d575f610c6d60ca83610bd2565b915050846001600160a01b0316816001600160a01b031603610c9457506001949350505050565b50600101610c59565b505f9392505050565b5f6102f6826110c1565b5f610cbc60c954610b95565b9050610cfc83838080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250610bf892505050565b60c9556040517f769656e5cb3218f54758f85bd184d41e987639374d6cb9e244439cc9d1abe1e2906104b8908390869086906119b1565b5f610d7561082c84848080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250610bf892505050565b90505f80610d8460ca846110cb565b91509150818015610da75750866001600160a01b0316816001600160a01b031614155b15610df3576001600160a01b0381165f81815260cd6020526040808220805461ffff19169055517fa69cccaa8b056f2577aa7e06e1eb14ae0eb526356819b9403f5b31f41f3bc5099190a25b610dff60ca84896110d9565b506001600160a01b0387165f90815260cd602052604090208690610e2382826118ba565b905050866001600160a01b03167f4e81e5495dbc85919bf1b9c12a9a4bb6d546f75d3f92fa0a36867fc29a51467c87604051610e5f9190611900565b60405180910390a250505050505050565b610e7a82826107a8565b6103a457610e87816110f6565b610e92836020611108565b604051602001610ea3929190611a0c565b60408051601f198184030181529082905262461bcd60e51b825261039191600401611526565b610ed382826107a8565b6103a4575f8281526065602090815260408083206001600160a01b03851684529091529020805460ff19166001179055610f0a3390565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b5f6107a1836001600160a01b03841661129d565b610f6c82826107a8565b156103a4575f8281526065602090815260408083206001600160a01b0385168085529252808320805460ff1916905551339285917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45050565b5f6107a1836001600160a01b0384166112e9565b5f60ff8216601f8111156102f657604051632cd44ac360e21b815260040160405180910390fd5b5f80806110108585610bed565b5f81815260029690960160205260409095205494959350505050565b5f825f01828154811061104157611041611899565b905f5260205f200154905092915050565b5f81815260028301602052604081205480151580611075575061107584846113cc565b6107a15760405162461bcd60e51b815260206004820152601e60248201527f456e756d657261626c654d61703a206e6f6e6578697374656e74206b657900006044820152606401610391565b5f6102f682610c40565b5f808080610be086866113d7565b5f6110ee84846001600160a01b03851661140f565b949350505050565b60606102f66001600160a01b03831660145b60605f611116836002611a7e565b611121906002611a95565b6001600160401b0381111561113857611138611885565b6040519080825280601f01601f191660200182016040528015611162576020820181803683370190505b509050600360fc1b815f8151811061117c5761117c611899565b60200101906001600160f81b03191690815f1a905350600f60fb1b816001815181106111aa576111aa611899565b60200101906001600160f81b03191690815f1a9053505f6111cc846002611a7e565b6111d7906001611a95565b90505b600181111561124e576f181899199a1a9b1b9c1cb0b131b232b360811b85600f166010811061120b5761120b611899565b1a60f81b82828151811061122157611221611899565b60200101906001600160f81b03191690815f1a90535060049490941c9361124781611aa8565b90506111da565b5083156107a15760405162461bcd60e51b815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e746044820152606401610391565b5f8181526001830160205260408120546112e257508154600181810184555f8481526020808220909301849055845484825282860190935260409020919091556102f6565b505f6102f6565b5f81815260018301602052604081205480156113c3575f61130b600183611abd565b85549091505f9061131e90600190611abd565b905081811461137d575f865f01828154811061133c5761133c611899565b905f5260205f200154905080875f01848154811061135c5761135c611899565b5f918252602080832090910192909255918252600188019052604090208390555b855486908061138e5761138e611ad0565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f9055600193505050506102f6565b5f9150506102f6565b5f6107a1838361142b565b5f818152600283016020526040812054819080611404576113f885856113cc565b92505f91506108779050565b600192509050610877565b5f82815260028401602052604081208290556110ee8484611442565b5f81815260018301602052604081205415156107a1565b5f6107a1838361129d565b5f6020828403121561145d575f5ffd5b81356001600160e01b0319811681146107a1575f5ffd5b5f60208284031215611484575f5ffd5b5035919050565b80356001600160a01b03811681146114a1575f5ffd5b919050565b5f5f604083850312156114b7575f5ffd5b823591506114c76020840161148b565b90509250929050565b5f5f604083850312156114e1575f5ffd5b6114ea8361148b565b91506114c76020840161148b565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f6107a160208301846114f8565b5f8151808452602084019350602083015f5b828110156115815761156b8683518051151582526020908101511515910152565b604095909501946020919091019060010161154a565b5093949350505050565b5f606082016060835280865180835260808501915060808160051b8601019250602088015f5b828110156115e257607f198786030184526115cd8583516114f8565b945060209384019391909101906001016115b1565b5050505082810360208401528085518083526020830191506020870192505f5b818110156116295783516001600160a01b0316835260209384019390920191600101611602565b5050838103604085015261163d8186611538565b979650505050505050565b5f5f60408385031215611659575f5ffd5b50508035926020909101359150565b5f5f83601f840112611678575f5ffd5b5081356001600160401b0381111561168e575f5ffd5b602083019150836020828501011115610877575f5ffd5b5f5f602083850312156116b6575f5ffd5b82356001600160401b038111156116cb575f5ffd5b6116d785828601611668565b90969095509350505050565b6001600160a01b0383168152606081016107a160208301848051151582526020908101511515910152565b5f5f8284036060811215611720575f5ffd5b6117298461148b565b92506040601f198201121561173c575f5ffd5b506020830190509250929050565b5f5f83601f84011261175a575f5ffd5b5081356001600160401b03811115611770575f5ffd5b6020830191508360208260051b8501011115610877575f5ffd5b5f5f5f5f5f5f5f5f6080898b0312156117a1575f5ffd5b88356001600160401b038111156117b6575f5ffd5b6117c28b828c0161174a565b90995097505060208901356001600160401b038111156117e0575f5ffd5b8901601f81018b136117f0575f5ffd5b80356001600160401b03811115611805575f5ffd5b8b60208260061b8401011115611819575f5ffd5b6020919091019650945060408901356001600160401b0381111561183b575f5ffd5b6118478b828c0161174a565b90955093505060608901356001600160401b03811115611865575f5ffd5b6118718b828c01611668565b999c989b5096995094979396929594505050565b634e487b7160e01b5f52604160045260245ffd5b634e487b7160e01b5f52603260045260245ffd5b8015158114610b50575f5ffd5b81356118c5816118ad565b815460ff19811691151560ff16918217835560208401356118e5816118ad565b61ffff199190911690911790151560081b61ff001617905550565b60408101823561190f816118ad565b151582526020830135611921816118ad565b80151560208401525092915050565b5f60208284031215611940575f5ffd5b6107a18261148b565b5f5f8335601e1984360301811261195e575f5ffd5b8301803591506001600160401b03821115611977575f5ffd5b602001915036819003821315610877575f5ffd5b805160208083015191908110156119ab575f198160200360031b1b821691505b50919050565b604081525f6119c360408301866114f8565b8281036020840152838152838560208301375f602085830101526020601f19601f860116820101915050949350505050565b5f81518060208401855e5f93019283525090919050565b7f416363657373436f6e74726f6c3a206163636f756e742000000000000000000081525f611a3d60178301856119f5565b7001034b99036b4b9b9b4b733903937b6329607d1b8152611a6160118201856119f5565b95945050505050565b634e487b7160e01b5f52601160045260245ffd5b80820281158282048414176102f6576102f6611a6a565b808201808211156102f6576102f6611a6a565b5f81611ab657611ab6611a6a565b505f190190565b818103818111156102f6576102f6611a6a565b634e487b7160e01b5f52603160045260245ffdfea26469706673582212203b8b9273fddfc4566fa479d664c3bfe76b37d84c82f51fd071e3ad2eb9025b6964736f6c634300081e0033",
}

// ProtocolRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ProtocolRegistryMetaData.ABI instead.
var ProtocolRegistryABI = ProtocolRegistryMetaData.ABI

// ProtocolRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ProtocolRegistryMetaData.Bin instead.
var ProtocolRegistryBin = ProtocolRegistryMetaData.Bin

// DeployProtocolRegistry deploys a new Ethereum contract, binding an instance of ProtocolRegistry to it.
func DeployProtocolRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ProtocolRegistry, error) {
	parsed, err := ProtocolRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ProtocolRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ProtocolRegistry{ProtocolRegistryCaller: ProtocolRegistryCaller{contract: contract}, ProtocolRegistryTransactor: ProtocolRegistryTransactor{contract: contract}, ProtocolRegistryFilterer: ProtocolRegistryFilterer{contract: contract}}, nil
}

// ProtocolRegistry is an auto generated Go binding around an Ethereum contract.
type ProtocolRegistry struct {
	ProtocolRegistryCaller     // Read-only binding to the contract
	ProtocolRegistryTransactor // Write-only binding to the contract
	ProtocolRegistryFilterer   // Log filterer for contract events
}

// ProtocolRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProtocolRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProtocolRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProtocolRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProtocolRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProtocolRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProtocolRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProtocolRegistrySession struct {
	Contract     *ProtocolRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProtocolRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProtocolRegistryCallerSession struct {
	Contract *ProtocolRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ProtocolRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProtocolRegistryTransactorSession struct {
	Contract     *ProtocolRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ProtocolRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProtocolRegistryRaw struct {
	Contract *ProtocolRegistry // Generic contract binding to access the raw methods on
}

// ProtocolRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProtocolRegistryCallerRaw struct {
	Contract *ProtocolRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ProtocolRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProtocolRegistryTransactorRaw struct {
	Contract *ProtocolRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProtocolRegistry creates a new instance of ProtocolRegistry, bound to a specific deployed contract.
func NewProtocolRegistry(address common.Address, backend bind.ContractBackend) (*ProtocolRegistry, error) {
	contract, err := bindProtocolRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProtocolRegistry{ProtocolRegistryCaller: ProtocolRegistryCaller{contract: contract}, ProtocolRegistryTransactor: ProtocolRegistryTransactor{contract: contract}, ProtocolRegistryFilterer: ProtocolRegistryFilterer{contract: contract}}, nil
}

// NewProtocolRegistryCaller creates a new read-only instance of ProtocolRegistry, bound to a specific deployed contract.
func NewProtocolRegistryCaller(address common.Address, caller bind.ContractCaller) (*ProtocolRegistryCaller, error) {
	contract, err := bindProtocolRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProtocolRegistryCaller{contract: contract}, nil
}

// NewProtocolRegistryTransactor creates a new write-only instance of ProtocolRegistry, bound to a specific deployed contract.
func NewProtocolRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ProtocolRegistryTransactor, error) {
	contract, err := bindProtocolRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProtocolRegistryTransactor{contract: contract}, nil
}

// NewProtocolRegistryFilterer creates a new log filterer instance of ProtocolRegistry, bound to a specific deployed contract.
func NewProtocolRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ProtocolRegistryFilterer, error) {
	contract, err := bindProtocolRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProtocolRegistryFilterer{contract: contract}, nil
}

// bindProtocolRegistry binds a generic wrapper to an already deployed contract.
func bindProtocolRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ProtocolRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProtocolRegistry *ProtocolRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProtocolRegistry.Contract.ProtocolRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProtocolRegistry *ProtocolRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProtocolRegistry.Contract.ProtocolRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProtocolRegistry *ProtocolRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProtocolRegistry.Contract.ProtocolRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProtocolRegistry *ProtocolRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProtocolRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProtocolRegistry *ProtocolRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProtocolRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProtocolRegistry *ProtocolRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProtocolRegistry.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ProtocolRegistry *ProtocolRegistryCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ProtocolRegistry.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ProtocolRegistry *ProtocolRegistrySession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ProtocolRegistry.Contract.DEFAULTADMINROLE(&_ProtocolRegistry.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ProtocolRegistry *ProtocolRegistryCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ProtocolRegistry.Contract.DEFAULTADMINROLE(&_ProtocolRegistry.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_ProtocolRegistry *ProtocolRegistryCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ProtocolRegistry.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_ProtocolRegistry *ProtocolRegistrySession) PAUSERROLE() ([32]byte, error) {
	return _ProtocolRegistry.Contract.PAUSERROLE(&_ProtocolRegistry.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_ProtocolRegistry *ProtocolRegistryCallerSession) PAUSERROLE() ([32]byte, error) {
	return _ProtocolRegistry.Contract.PAUSERROLE(&_ProtocolRegistry.CallOpts)
}

// GetAddress is a free data retrieval call binding the contract method 0xbf40fac1.
//
// Solidity: function getAddress(string name) view returns(address)
func (_ProtocolRegistry *ProtocolRegistryCaller) GetAddress(opts *bind.CallOpts, name string) (common.Address, error) {
	var out []interface{}
	err := _ProtocolRegistry.contract.Call(opts, &out, "getAddress", name)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddress is a free data retrieval call binding the contract method 0xbf40fac1.
//
// Solidity: function getAddress(string name) view returns(address)
func (_ProtocolRegistry *ProtocolRegistrySession) GetAddress(name string) (common.Address, error) {
	return _ProtocolRegistry.Contract.GetAddress(&_ProtocolRegistry.CallOpts, name)
}

// GetAddress is a free data retrieval call binding the contract method 0xbf40fac1.
//
// Solidity: function getAddress(string name) view returns(address)
func (_ProtocolRegistry *ProtocolRegistryCallerSession) GetAddress(name string) (common.Address, error) {
	return _ProtocolRegistry.Contract.GetAddress(&_ProtocolRegistry.CallOpts, name)
}

// GetAllDeployments is a free data retrieval call binding the contract method 0x8eec00b8.
//
// Solidity: function getAllDeployments() view returns(string[] names, address[] addresses, (bool,bool)[] configs)
func (_ProtocolRegistry *ProtocolRegistryCaller) GetAllDeployments(opts *bind.CallOpts) (struct {
	Names     []string
	Addresses []common.Address
	Configs   []IProtocolRegistryTypesDeploymentConfig
}, error) {
	var out []interface{}
	err := _ProtocolRegistry.contract.Call(opts, &out, "getAllDeployments")

	outstruct := new(struct {
		Names     []string
		Addresses []common.Address
		Configs   []IProtocolRegistryTypesDeploymentConfig
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Names = *abi.ConvertType(out[0], new([]string)).(*[]string)
	outstruct.Addresses = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	outstruct.Configs = *abi.ConvertType(out[2], new([]IProtocolRegistryTypesDeploymentConfig)).(*[]IProtocolRegistryTypesDeploymentConfig)

	return *outstruct, err

}

// GetAllDeployments is a free data retrieval call binding the contract method 0x8eec00b8.
//
// Solidity: function getAllDeployments() view returns(string[] names, address[] addresses, (bool,bool)[] configs)
func (_ProtocolRegistry *ProtocolRegistrySession) GetAllDeployments() (struct {
	Names     []string
	Addresses []common.Address
	Configs   []IProtocolRegistryTypesDeploymentConfig
}, error) {
	return _ProtocolRegistry.Contract.GetAllDeployments(&_ProtocolRegistry.CallOpts)
}

// GetAllDeployments is a free data retrieval call binding the contract method 0x8eec00b8.
//
// Solidity: function getAllDeployments() view returns(string[] names, address[] addresses, (bool,bool)[] configs)
func (_ProtocolRegistry *ProtocolRegistryCallerSession) GetAllDeployments() (struct {
	Names     []string
	Addresses []common.Address
	Configs   []IProtocolRegistryTypesDeploymentConfig
}, error) {
	return _ProtocolRegistry.Contract.GetAllDeployments(&_ProtocolRegistry.CallOpts)
}

// GetDeployment is a free data retrieval call binding the contract method 0xa8091d97.
//
// Solidity: function getDeployment(string name) view returns(address addr, (bool,bool) config)
func (_ProtocolRegistry *ProtocolRegistryCaller) GetDeployment(opts *bind.CallOpts, name string) (struct {
	Addr   common.Address
	Config IProtocolRegistryTypesDeploymentConfig
}, error) {
	var out []interface{}
	err := _ProtocolRegistry.contract.Call(opts, &out, "getDeployment", name)

	outstruct := new(struct {
		Addr   common.Address
		Config IProtocolRegistryTypesDeploymentConfig
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Addr = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Config = *abi.ConvertType(out[1], new(IProtocolRegistryTypesDeploymentConfig)).(*IProtocolRegistryTypesDeploymentConfig)

	return *outstruct, err

}

// GetDeployment is a free data retrieval call binding the contract method 0xa8091d97.
//
// Solidity: function getDeployment(string name) view returns(address addr, (bool,bool) config)
func (_ProtocolRegistry *ProtocolRegistrySession) GetDeployment(name string) (struct {
	Addr   common.Address
	Config IProtocolRegistryTypesDeploymentConfig
}, error) {
	return _ProtocolRegistry.Contract.GetDeployment(&_ProtocolRegistry.CallOpts, name)
}

// GetDeployment is a free data retrieval call binding the contract method 0xa8091d97.
//
// Solidity: function getDeployment(string name) view returns(address addr, (bool,bool) config)
func (_ProtocolRegistry *ProtocolRegistryCallerSession) GetDeployment(name string) (struct {
	Addr   common.Address
	Config IProtocolRegistryTypesDeploymentConfig
}, error) {
	return _ProtocolRegistry.Contract.GetDeployment(&_ProtocolRegistry.CallOpts, name)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ProtocolRegistry *ProtocolRegistryCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ProtocolRegistry.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ProtocolRegistry *ProtocolRegistrySession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ProtocolRegistry.Contract.GetRoleAdmin(&_ProtocolRegistry.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ProtocolRegistry *ProtocolRegistryCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ProtocolRegistry.Contract.GetRoleAdmin(&_ProtocolRegistry.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_ProtocolRegistry *ProtocolRegistryCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ProtocolRegistry.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_ProtocolRegistry *ProtocolRegistrySession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _ProtocolRegistry.Contract.GetRoleMember(&_ProtocolRegistry.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_ProtocolRegistry *ProtocolRegistryCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _ProtocolRegistry.Contract.GetRoleMember(&_ProtocolRegistry.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_ProtocolRegistry *ProtocolRegistryCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ProtocolRegistry.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_ProtocolRegistry *ProtocolRegistrySession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _ProtocolRegistry.Contract.GetRoleMemberCount(&_ProtocolRegistry.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_ProtocolRegistry *ProtocolRegistryCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _ProtocolRegistry.Contract.GetRoleMemberCount(&_ProtocolRegistry.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ProtocolRegistry *ProtocolRegistryCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ProtocolRegistry.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ProtocolRegistry *ProtocolRegistrySession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ProtocolRegistry.Contract.HasRole(&_ProtocolRegistry.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ProtocolRegistry *ProtocolRegistryCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ProtocolRegistry.Contract.HasRole(&_ProtocolRegistry.CallOpts, role, account)
}

// MajorVersion is a free data retrieval call binding the contract method 0xaf05a5c5.
//
// Solidity: function majorVersion() view returns(string)
func (_ProtocolRegistry *ProtocolRegistryCaller) MajorVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ProtocolRegistry.contract.Call(opts, &out, "majorVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// MajorVersion is a free data retrieval call binding the contract method 0xaf05a5c5.
//
// Solidity: function majorVersion() view returns(string)
func (_ProtocolRegistry *ProtocolRegistrySession) MajorVersion() (string, error) {
	return _ProtocolRegistry.Contract.MajorVersion(&_ProtocolRegistry.CallOpts)
}

// MajorVersion is a free data retrieval call binding the contract method 0xaf05a5c5.
//
// Solidity: function majorVersion() view returns(string)
func (_ProtocolRegistry *ProtocolRegistryCallerSession) MajorVersion() (string, error) {
	return _ProtocolRegistry.Contract.MajorVersion(&_ProtocolRegistry.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ProtocolRegistry *ProtocolRegistryCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ProtocolRegistry.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ProtocolRegistry *ProtocolRegistrySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ProtocolRegistry.Contract.SupportsInterface(&_ProtocolRegistry.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ProtocolRegistry *ProtocolRegistryCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ProtocolRegistry.Contract.SupportsInterface(&_ProtocolRegistry.CallOpts, interfaceId)
}

// TotalDeployments is a free data retrieval call binding the contract method 0xfb35b4e4.
//
// Solidity: function totalDeployments() view returns(uint256)
func (_ProtocolRegistry *ProtocolRegistryCaller) TotalDeployments(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ProtocolRegistry.contract.Call(opts, &out, "totalDeployments")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDeployments is a free data retrieval call binding the contract method 0xfb35b4e4.
//
// Solidity: function totalDeployments() view returns(uint256)
func (_ProtocolRegistry *ProtocolRegistrySession) TotalDeployments() (*big.Int, error) {
	return _ProtocolRegistry.Contract.TotalDeployments(&_ProtocolRegistry.CallOpts)
}

// TotalDeployments is a free data retrieval call binding the contract method 0xfb35b4e4.
//
// Solidity: function totalDeployments() view returns(uint256)
func (_ProtocolRegistry *ProtocolRegistryCallerSession) TotalDeployments() (*big.Int, error) {
	return _ProtocolRegistry.Contract.TotalDeployments(&_ProtocolRegistry.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ProtocolRegistry *ProtocolRegistryCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ProtocolRegistry.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ProtocolRegistry *ProtocolRegistrySession) Version() (string, error) {
	return _ProtocolRegistry.Contract.Version(&_ProtocolRegistry.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ProtocolRegistry *ProtocolRegistryCallerSession) Version() (string, error) {
	return _ProtocolRegistry.Contract.Version(&_ProtocolRegistry.CallOpts)
}

// Configure is a paid mutator transaction binding the contract method 0xd3466389.
//
// Solidity: function configure(address addr, (bool,bool) config) returns()
func (_ProtocolRegistry *ProtocolRegistryTransactor) Configure(opts *bind.TransactOpts, addr common.Address, config IProtocolRegistryTypesDeploymentConfig) (*types.Transaction, error) {
	return _ProtocolRegistry.contract.Transact(opts, "configure", addr, config)
}

// Configure is a paid mutator transaction binding the contract method 0xd3466389.
//
// Solidity: function configure(address addr, (bool,bool) config) returns()
func (_ProtocolRegistry *ProtocolRegistrySession) Configure(addr common.Address, config IProtocolRegistryTypesDeploymentConfig) (*types.Transaction, error) {
	return _ProtocolRegistry.Contract.Configure(&_ProtocolRegistry.TransactOpts, addr, config)
}

// Configure is a paid mutator transaction binding the contract method 0xd3466389.
//
// Solidity: function configure(address addr, (bool,bool) config) returns()
func (_ProtocolRegistry *ProtocolRegistryTransactorSession) Configure(addr common.Address, config IProtocolRegistryTypesDeploymentConfig) (*types.Transaction, error) {
	return _ProtocolRegistry.Contract.Configure(&_ProtocolRegistry.TransactOpts, addr, config)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ProtocolRegistry *ProtocolRegistryTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ProtocolRegistry.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ProtocolRegistry *ProtocolRegistrySession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ProtocolRegistry.Contract.GrantRole(&_ProtocolRegistry.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ProtocolRegistry *ProtocolRegistryTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ProtocolRegistry.Contract.GrantRole(&_ProtocolRegistry.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address initialAdmin, address pauserMultisig) returns()
func (_ProtocolRegistry *ProtocolRegistryTransactor) Initialize(opts *bind.TransactOpts, initialAdmin common.Address, pauserMultisig common.Address) (*types.Transaction, error) {
	return _ProtocolRegistry.contract.Transact(opts, "initialize", initialAdmin, pauserMultisig)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address initialAdmin, address pauserMultisig) returns()
func (_ProtocolRegistry *ProtocolRegistrySession) Initialize(initialAdmin common.Address, pauserMultisig common.Address) (*types.Transaction, error) {
	return _ProtocolRegistry.Contract.Initialize(&_ProtocolRegistry.TransactOpts, initialAdmin, pauserMultisig)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address initialAdmin, address pauserMultisig) returns()
func (_ProtocolRegistry *ProtocolRegistryTransactorSession) Initialize(initialAdmin common.Address, pauserMultisig common.Address) (*types.Transaction, error) {
	return _ProtocolRegistry.Contract.Initialize(&_ProtocolRegistry.TransactOpts, initialAdmin, pauserMultisig)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ProtocolRegistry *ProtocolRegistryTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProtocolRegistry.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ProtocolRegistry *ProtocolRegistrySession) PauseAll() (*types.Transaction, error) {
	return _ProtocolRegistry.Contract.PauseAll(&_ProtocolRegistry.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ProtocolRegistry *ProtocolRegistryTransactorSession) PauseAll() (*types.Transaction, error) {
	return _ProtocolRegistry.Contract.PauseAll(&_ProtocolRegistry.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ProtocolRegistry *ProtocolRegistryTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ProtocolRegistry.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ProtocolRegistry *ProtocolRegistrySession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ProtocolRegistry.Contract.RenounceRole(&_ProtocolRegistry.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ProtocolRegistry *ProtocolRegistryTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ProtocolRegistry.Contract.RenounceRole(&_ProtocolRegistry.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ProtocolRegistry *ProtocolRegistryTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ProtocolRegistry.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ProtocolRegistry *ProtocolRegistrySession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ProtocolRegistry.Contract.RevokeRole(&_ProtocolRegistry.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ProtocolRegistry *ProtocolRegistryTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ProtocolRegistry.Contract.RevokeRole(&_ProtocolRegistry.TransactOpts, role, account)
}

// Ship is a paid mutator transaction binding the contract method 0xfc434a8a.
//
// Solidity: function ship(address[] addresses, (bool,bool)[] configs, string[] names, string semanticVersion) returns()
func (_ProtocolRegistry *ProtocolRegistryTransactor) Ship(opts *bind.TransactOpts, addresses []common.Address, configs []IProtocolRegistryTypesDeploymentConfig, names []string, semanticVersion string) (*types.Transaction, error) {
	return _ProtocolRegistry.contract.Transact(opts, "ship", addresses, configs, names, semanticVersion)
}

// Ship is a paid mutator transaction binding the contract method 0xfc434a8a.
//
// Solidity: function ship(address[] addresses, (bool,bool)[] configs, string[] names, string semanticVersion) returns()
func (_ProtocolRegistry *ProtocolRegistrySession) Ship(addresses []common.Address, configs []IProtocolRegistryTypesDeploymentConfig, names []string, semanticVersion string) (*types.Transaction, error) {
	return _ProtocolRegistry.Contract.Ship(&_ProtocolRegistry.TransactOpts, addresses, configs, names, semanticVersion)
}

// Ship is a paid mutator transaction binding the contract method 0xfc434a8a.
//
// Solidity: function ship(address[] addresses, (bool,bool)[] configs, string[] names, string semanticVersion) returns()
func (_ProtocolRegistry *ProtocolRegistryTransactorSession) Ship(addresses []common.Address, configs []IProtocolRegistryTypesDeploymentConfig, names []string, semanticVersion string) (*types.Transaction, error) {
	return _ProtocolRegistry.Contract.Ship(&_ProtocolRegistry.TransactOpts, addresses, configs, names, semanticVersion)
}

// ProtocolRegistryDeploymentConfigDeletedIterator is returned from FilterDeploymentConfigDeleted and is used to iterate over the raw logs and unpacked data for DeploymentConfigDeleted events raised by the ProtocolRegistry contract.
type ProtocolRegistryDeploymentConfigDeletedIterator struct {
	Event *ProtocolRegistryDeploymentConfigDeleted // Event containing the contract specifics and raw log

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
func (it *ProtocolRegistryDeploymentConfigDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProtocolRegistryDeploymentConfigDeleted)
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
		it.Event = new(ProtocolRegistryDeploymentConfigDeleted)
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
func (it *ProtocolRegistryDeploymentConfigDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProtocolRegistryDeploymentConfigDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProtocolRegistryDeploymentConfigDeleted represents a DeploymentConfigDeleted event raised by the ProtocolRegistry contract.
type ProtocolRegistryDeploymentConfigDeleted struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDeploymentConfigDeleted is a free log retrieval operation binding the contract event 0xa69cccaa8b056f2577aa7e06e1eb14ae0eb526356819b9403f5b31f41f3bc509.
//
// Solidity: event DeploymentConfigDeleted(address indexed addr)
func (_ProtocolRegistry *ProtocolRegistryFilterer) FilterDeploymentConfigDeleted(opts *bind.FilterOpts, addr []common.Address) (*ProtocolRegistryDeploymentConfigDeletedIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _ProtocolRegistry.contract.FilterLogs(opts, "DeploymentConfigDeleted", addrRule)
	if err != nil {
		return nil, err
	}
	return &ProtocolRegistryDeploymentConfigDeletedIterator{contract: _ProtocolRegistry.contract, event: "DeploymentConfigDeleted", logs: logs, sub: sub}, nil
}

// WatchDeploymentConfigDeleted is a free log subscription operation binding the contract event 0xa69cccaa8b056f2577aa7e06e1eb14ae0eb526356819b9403f5b31f41f3bc509.
//
// Solidity: event DeploymentConfigDeleted(address indexed addr)
func (_ProtocolRegistry *ProtocolRegistryFilterer) WatchDeploymentConfigDeleted(opts *bind.WatchOpts, sink chan<- *ProtocolRegistryDeploymentConfigDeleted, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _ProtocolRegistry.contract.WatchLogs(opts, "DeploymentConfigDeleted", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProtocolRegistryDeploymentConfigDeleted)
				if err := _ProtocolRegistry.contract.UnpackLog(event, "DeploymentConfigDeleted", log); err != nil {
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

// ParseDeploymentConfigDeleted is a log parse operation binding the contract event 0xa69cccaa8b056f2577aa7e06e1eb14ae0eb526356819b9403f5b31f41f3bc509.
//
// Solidity: event DeploymentConfigDeleted(address indexed addr)
func (_ProtocolRegistry *ProtocolRegistryFilterer) ParseDeploymentConfigDeleted(log types.Log) (*ProtocolRegistryDeploymentConfigDeleted, error) {
	event := new(ProtocolRegistryDeploymentConfigDeleted)
	if err := _ProtocolRegistry.contract.UnpackLog(event, "DeploymentConfigDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProtocolRegistryDeploymentConfiguredIterator is returned from FilterDeploymentConfigured and is used to iterate over the raw logs and unpacked data for DeploymentConfigured events raised by the ProtocolRegistry contract.
type ProtocolRegistryDeploymentConfiguredIterator struct {
	Event *ProtocolRegistryDeploymentConfigured // Event containing the contract specifics and raw log

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
func (it *ProtocolRegistryDeploymentConfiguredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProtocolRegistryDeploymentConfigured)
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
		it.Event = new(ProtocolRegistryDeploymentConfigured)
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
func (it *ProtocolRegistryDeploymentConfiguredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProtocolRegistryDeploymentConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProtocolRegistryDeploymentConfigured represents a DeploymentConfigured event raised by the ProtocolRegistry contract.
type ProtocolRegistryDeploymentConfigured struct {
	Addr   common.Address
	Config IProtocolRegistryTypesDeploymentConfig
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeploymentConfigured is a free log retrieval operation binding the contract event 0x6c5879ec82f910f6d12844857cfb8eb474dcecc9aa5b8257c7a77dcb42990e9d.
//
// Solidity: event DeploymentConfigured(address indexed addr, (bool,bool) config)
func (_ProtocolRegistry *ProtocolRegistryFilterer) FilterDeploymentConfigured(opts *bind.FilterOpts, addr []common.Address) (*ProtocolRegistryDeploymentConfiguredIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _ProtocolRegistry.contract.FilterLogs(opts, "DeploymentConfigured", addrRule)
	if err != nil {
		return nil, err
	}
	return &ProtocolRegistryDeploymentConfiguredIterator{contract: _ProtocolRegistry.contract, event: "DeploymentConfigured", logs: logs, sub: sub}, nil
}

// WatchDeploymentConfigured is a free log subscription operation binding the contract event 0x6c5879ec82f910f6d12844857cfb8eb474dcecc9aa5b8257c7a77dcb42990e9d.
//
// Solidity: event DeploymentConfigured(address indexed addr, (bool,bool) config)
func (_ProtocolRegistry *ProtocolRegistryFilterer) WatchDeploymentConfigured(opts *bind.WatchOpts, sink chan<- *ProtocolRegistryDeploymentConfigured, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _ProtocolRegistry.contract.WatchLogs(opts, "DeploymentConfigured", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProtocolRegistryDeploymentConfigured)
				if err := _ProtocolRegistry.contract.UnpackLog(event, "DeploymentConfigured", log); err != nil {
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

// ParseDeploymentConfigured is a log parse operation binding the contract event 0x6c5879ec82f910f6d12844857cfb8eb474dcecc9aa5b8257c7a77dcb42990e9d.
//
// Solidity: event DeploymentConfigured(address indexed addr, (bool,bool) config)
func (_ProtocolRegistry *ProtocolRegistryFilterer) ParseDeploymentConfigured(log types.Log) (*ProtocolRegistryDeploymentConfigured, error) {
	event := new(ProtocolRegistryDeploymentConfigured)
	if err := _ProtocolRegistry.contract.UnpackLog(event, "DeploymentConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProtocolRegistryDeploymentShippedIterator is returned from FilterDeploymentShipped and is used to iterate over the raw logs and unpacked data for DeploymentShipped events raised by the ProtocolRegistry contract.
type ProtocolRegistryDeploymentShippedIterator struct {
	Event *ProtocolRegistryDeploymentShipped // Event containing the contract specifics and raw log

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
func (it *ProtocolRegistryDeploymentShippedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProtocolRegistryDeploymentShipped)
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
		it.Event = new(ProtocolRegistryDeploymentShipped)
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
func (it *ProtocolRegistryDeploymentShippedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProtocolRegistryDeploymentShippedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProtocolRegistryDeploymentShipped represents a DeploymentShipped event raised by the ProtocolRegistry contract.
type ProtocolRegistryDeploymentShipped struct {
	Addr   common.Address
	Config IProtocolRegistryTypesDeploymentConfig
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeploymentShipped is a free log retrieval operation binding the contract event 0x4e81e5495dbc85919bf1b9c12a9a4bb6d546f75d3f92fa0a36867fc29a51467c.
//
// Solidity: event DeploymentShipped(address indexed addr, (bool,bool) config)
func (_ProtocolRegistry *ProtocolRegistryFilterer) FilterDeploymentShipped(opts *bind.FilterOpts, addr []common.Address) (*ProtocolRegistryDeploymentShippedIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _ProtocolRegistry.contract.FilterLogs(opts, "DeploymentShipped", addrRule)
	if err != nil {
		return nil, err
	}
	return &ProtocolRegistryDeploymentShippedIterator{contract: _ProtocolRegistry.contract, event: "DeploymentShipped", logs: logs, sub: sub}, nil
}

// WatchDeploymentShipped is a free log subscription operation binding the contract event 0x4e81e5495dbc85919bf1b9c12a9a4bb6d546f75d3f92fa0a36867fc29a51467c.
//
// Solidity: event DeploymentShipped(address indexed addr, (bool,bool) config)
func (_ProtocolRegistry *ProtocolRegistryFilterer) WatchDeploymentShipped(opts *bind.WatchOpts, sink chan<- *ProtocolRegistryDeploymentShipped, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _ProtocolRegistry.contract.WatchLogs(opts, "DeploymentShipped", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProtocolRegistryDeploymentShipped)
				if err := _ProtocolRegistry.contract.UnpackLog(event, "DeploymentShipped", log); err != nil {
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

// ParseDeploymentShipped is a log parse operation binding the contract event 0x4e81e5495dbc85919bf1b9c12a9a4bb6d546f75d3f92fa0a36867fc29a51467c.
//
// Solidity: event DeploymentShipped(address indexed addr, (bool,bool) config)
func (_ProtocolRegistry *ProtocolRegistryFilterer) ParseDeploymentShipped(log types.Log) (*ProtocolRegistryDeploymentShipped, error) {
	event := new(ProtocolRegistryDeploymentShipped)
	if err := _ProtocolRegistry.contract.UnpackLog(event, "DeploymentShipped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProtocolRegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ProtocolRegistry contract.
type ProtocolRegistryInitializedIterator struct {
	Event *ProtocolRegistryInitialized // Event containing the contract specifics and raw log

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
func (it *ProtocolRegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProtocolRegistryInitialized)
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
		it.Event = new(ProtocolRegistryInitialized)
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
func (it *ProtocolRegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProtocolRegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProtocolRegistryInitialized represents a Initialized event raised by the ProtocolRegistry contract.
type ProtocolRegistryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ProtocolRegistry *ProtocolRegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*ProtocolRegistryInitializedIterator, error) {

	logs, sub, err := _ProtocolRegistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ProtocolRegistryInitializedIterator{contract: _ProtocolRegistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ProtocolRegistry *ProtocolRegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ProtocolRegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _ProtocolRegistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProtocolRegistryInitialized)
				if err := _ProtocolRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ProtocolRegistry *ProtocolRegistryFilterer) ParseInitialized(log types.Log) (*ProtocolRegistryInitialized, error) {
	event := new(ProtocolRegistryInitialized)
	if err := _ProtocolRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProtocolRegistryRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the ProtocolRegistry contract.
type ProtocolRegistryRoleAdminChangedIterator struct {
	Event *ProtocolRegistryRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ProtocolRegistryRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProtocolRegistryRoleAdminChanged)
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
		it.Event = new(ProtocolRegistryRoleAdminChanged)
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
func (it *ProtocolRegistryRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProtocolRegistryRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProtocolRegistryRoleAdminChanged represents a RoleAdminChanged event raised by the ProtocolRegistry contract.
type ProtocolRegistryRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ProtocolRegistry *ProtocolRegistryFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ProtocolRegistryRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _ProtocolRegistry.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ProtocolRegistryRoleAdminChangedIterator{contract: _ProtocolRegistry.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ProtocolRegistry *ProtocolRegistryFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ProtocolRegistryRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _ProtocolRegistry.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProtocolRegistryRoleAdminChanged)
				if err := _ProtocolRegistry.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ProtocolRegistry *ProtocolRegistryFilterer) ParseRoleAdminChanged(log types.Log) (*ProtocolRegistryRoleAdminChanged, error) {
	event := new(ProtocolRegistryRoleAdminChanged)
	if err := _ProtocolRegistry.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProtocolRegistryRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ProtocolRegistry contract.
type ProtocolRegistryRoleGrantedIterator struct {
	Event *ProtocolRegistryRoleGranted // Event containing the contract specifics and raw log

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
func (it *ProtocolRegistryRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProtocolRegistryRoleGranted)
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
		it.Event = new(ProtocolRegistryRoleGranted)
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
func (it *ProtocolRegistryRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProtocolRegistryRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProtocolRegistryRoleGranted represents a RoleGranted event raised by the ProtocolRegistry contract.
type ProtocolRegistryRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ProtocolRegistry *ProtocolRegistryFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ProtocolRegistryRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ProtocolRegistry.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ProtocolRegistryRoleGrantedIterator{contract: _ProtocolRegistry.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ProtocolRegistry *ProtocolRegistryFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ProtocolRegistryRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ProtocolRegistry.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProtocolRegistryRoleGranted)
				if err := _ProtocolRegistry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ProtocolRegistry *ProtocolRegistryFilterer) ParseRoleGranted(log types.Log) (*ProtocolRegistryRoleGranted, error) {
	event := new(ProtocolRegistryRoleGranted)
	if err := _ProtocolRegistry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProtocolRegistryRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ProtocolRegistry contract.
type ProtocolRegistryRoleRevokedIterator struct {
	Event *ProtocolRegistryRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ProtocolRegistryRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProtocolRegistryRoleRevoked)
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
		it.Event = new(ProtocolRegistryRoleRevoked)
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
func (it *ProtocolRegistryRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProtocolRegistryRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProtocolRegistryRoleRevoked represents a RoleRevoked event raised by the ProtocolRegistry contract.
type ProtocolRegistryRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ProtocolRegistry *ProtocolRegistryFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ProtocolRegistryRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ProtocolRegistry.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ProtocolRegistryRoleRevokedIterator{contract: _ProtocolRegistry.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ProtocolRegistry *ProtocolRegistryFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ProtocolRegistryRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ProtocolRegistry.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProtocolRegistryRoleRevoked)
				if err := _ProtocolRegistry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ProtocolRegistry *ProtocolRegistryFilterer) ParseRoleRevoked(log types.Log) (*ProtocolRegistryRoleRevoked, error) {
	event := new(ProtocolRegistryRoleRevoked)
	if err := _ProtocolRegistry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProtocolRegistrySemanticVersionUpdatedIterator is returned from FilterSemanticVersionUpdated and is used to iterate over the raw logs and unpacked data for SemanticVersionUpdated events raised by the ProtocolRegistry contract.
type ProtocolRegistrySemanticVersionUpdatedIterator struct {
	Event *ProtocolRegistrySemanticVersionUpdated // Event containing the contract specifics and raw log

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
func (it *ProtocolRegistrySemanticVersionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProtocolRegistrySemanticVersionUpdated)
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
		it.Event = new(ProtocolRegistrySemanticVersionUpdated)
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
func (it *ProtocolRegistrySemanticVersionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProtocolRegistrySemanticVersionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProtocolRegistrySemanticVersionUpdated represents a SemanticVersionUpdated event raised by the ProtocolRegistry contract.
type ProtocolRegistrySemanticVersionUpdated struct {
	PreviousSemanticVersion string
	SemanticVersion         string
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterSemanticVersionUpdated is a free log retrieval operation binding the contract event 0x769656e5cb3218f54758f85bd184d41e987639374d6cb9e244439cc9d1abe1e2.
//
// Solidity: event SemanticVersionUpdated(string previousSemanticVersion, string semanticVersion)
func (_ProtocolRegistry *ProtocolRegistryFilterer) FilterSemanticVersionUpdated(opts *bind.FilterOpts) (*ProtocolRegistrySemanticVersionUpdatedIterator, error) {

	logs, sub, err := _ProtocolRegistry.contract.FilterLogs(opts, "SemanticVersionUpdated")
	if err != nil {
		return nil, err
	}
	return &ProtocolRegistrySemanticVersionUpdatedIterator{contract: _ProtocolRegistry.contract, event: "SemanticVersionUpdated", logs: logs, sub: sub}, nil
}

// WatchSemanticVersionUpdated is a free log subscription operation binding the contract event 0x769656e5cb3218f54758f85bd184d41e987639374d6cb9e244439cc9d1abe1e2.
//
// Solidity: event SemanticVersionUpdated(string previousSemanticVersion, string semanticVersion)
func (_ProtocolRegistry *ProtocolRegistryFilterer) WatchSemanticVersionUpdated(opts *bind.WatchOpts, sink chan<- *ProtocolRegistrySemanticVersionUpdated) (event.Subscription, error) {

	logs, sub, err := _ProtocolRegistry.contract.WatchLogs(opts, "SemanticVersionUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProtocolRegistrySemanticVersionUpdated)
				if err := _ProtocolRegistry.contract.UnpackLog(event, "SemanticVersionUpdated", log); err != nil {
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

// ParseSemanticVersionUpdated is a log parse operation binding the contract event 0x769656e5cb3218f54758f85bd184d41e987639374d6cb9e244439cc9d1abe1e2.
//
// Solidity: event SemanticVersionUpdated(string previousSemanticVersion, string semanticVersion)
func (_ProtocolRegistry *ProtocolRegistryFilterer) ParseSemanticVersionUpdated(log types.Log) (*ProtocolRegistrySemanticVersionUpdated, error) {
	event := new(ProtocolRegistrySemanticVersionUpdated)
	if err := _ProtocolRegistry.contract.UnpackLog(event, "SemanticVersionUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
