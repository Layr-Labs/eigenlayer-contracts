// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ReleaseManager

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

// IReleaseManagerTypesArtifact is an auto generated low-level Go binding around an user-defined struct.
type IReleaseManagerTypesArtifact struct {
	Digest   [32]byte
	Registry string
}

// IReleaseManagerTypesRelease is an auto generated low-level Go binding around an user-defined struct.
type IReleaseManagerTypesRelease struct {
	Artifacts     []IReleaseManagerTypesArtifact
	UpgradeByTime uint32
}

// OperatorSet is an auto generated low-level Go binding around an user-defined struct.
type OperatorSet struct {
	Avs common.Address
	Id  uint32
}

// ReleaseManagerMetaData contains all meta data concerning the ReleaseManager contract.
var ReleaseManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_permissionController\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"},{\"name\":\"_version\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getLatestRelease\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIReleaseManagerTypes.Release\",\"components\":[{\"name\":\"artifacts\",\"type\":\"tuple[]\",\"internalType\":\"structIReleaseManagerTypes.Artifact[]\",\"components\":[{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"registry\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"upgradeByTime\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLatestUpgradeByTime\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMetadataURI\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRelease\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"releaseId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIReleaseManagerTypes.Release\",\"components\":[{\"name\":\"artifacts\",\"type\":\"tuple[]\",\"internalType\":\"structIReleaseManagerTypes.Artifact[]\",\"components\":[{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"registry\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"upgradeByTime\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalReleases\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isValidRelease\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"releaseId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"permissionController\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"publishMetadataURI\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"publishRelease\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"release\",\"type\":\"tuple\",\"internalType\":\"structIReleaseManagerTypes.Release\",\"components\":[{\"name\":\"artifacts\",\"type\":\"tuple[]\",\"internalType\":\"structIReleaseManagerTypes.Artifact[]\",\"components\":[{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"registry\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"upgradeByTime\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"releaseId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MetadataURIPublished\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":true,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"metadataURI\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasePublished\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":true,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"releaseId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"release\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIReleaseManagerTypes.Release\",\"components\":[{\"name\":\"artifacts\",\"type\":\"tuple[]\",\"internalType\":\"structIReleaseManagerTypes.Artifact[]\",\"components\":[{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"registry\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"upgradeByTime\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidMetadataURI\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPermissions\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidUpgradeByTime\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustPublishMetadataURI\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoReleases\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]}]",
	Bin: "0x60c060405234801561000f575f5ffd5b5060405161168938038061168983398101604081905261002e9161016a565b6001600160a01b0382166080528061004581610058565b60a0525061005161009e565b5050610294565b5f5f829050601f8151111561008b578260405163305a27a960e01b81526004016100829190610239565b60405180910390fd5b80516100968261026e565b179392505050565b5f54610100900460ff16156101055760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b6064820152608401610082565b5f5460ff90811614610154575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b634e487b7160e01b5f52604160045260245ffd5b5f5f6040838503121561017b575f5ffd5b82516001600160a01b0381168114610191575f5ffd5b60208401519092506001600160401b038111156101ac575f5ffd5b8301601f810185136101bc575f5ffd5b80516001600160401b038111156101d5576101d5610156565b604051601f8201601f19908116603f011681016001600160401b038111828210171561020357610203610156565b60405281815282820160200187101561021a575f5ffd5b8160208401602083015e5f602083830101528093505050509250929050565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b8051602080830151919081101561028e575f198160200360031b1b821691505b50919050565b60805160a0516113ce6102bb5f395f61047101525f818160cd01526109ae01526113ce5ff3fe608060405234801561000f575f5ffd5b506004361061009b575f3560e01c806366f409f71161006357806366f409f7146101545780637c09ea8214610175578063a9e0ed6814610188578063b053b56d146101b0578063d30eeb88146101c3575f5ffd5b80633acab5fc1461009f5780634657e26a146100c85780634840a67c14610107578063517e40681461011c57806354fd4d501461013f575b5f5ffd5b6100b26100ad366004610b2b565b6101e4565b6040516100bf9190610c17565b60405180910390f35b6100ef7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016100bf565b61011a610115366004610c46565b610348565b005b61012f61012a366004610b2b565b610428565b60405190151581526020016100bf565b61014761046a565b6040516100bf9190610cc5565b610167610162366004610cd7565b61049a565b6040519081526020016100bf565b610167610183366004610cf1565b6104bd565b61019b610196366004610cd7565b610636565b60405163ffffffff90911681526020016100bf565b6101476101be366004610cd7565b6106bc565b6101d66101d1366004610cd7565b610765565b6040516100bf929190610d3d565b60408051808201909152606081525f602082015260015f6102048561090d565b81526020019081526020015f20828154811061022257610222610d5d565b905f5260205f2090600202016040518060400160405290815f8201805480602002602001604051908101604052809291908181526020015f905b82821015610325578382905f5260205f2090600202016040518060400160405290815f820154815260200160018201805461029690610d71565b80601f01602080910402602001604051908101604052809291908181526020018280546102c290610d71565b801561030d5780601f106102e45761010080835404028352916020019161030d565b820191905f5260205f20905b8154815290600101906020018083116102f057829003601f168201915b5050505050815250508152602001906001019061025c565b505050908252506001919091015463ffffffff1660209091015290505b92915050565b6103556020840184610da3565b61035e81610970565b61037b5760405163932d94f760e01b815260040160405180910390fd5b5f82900361039c57604051630eec403f60e41b815260040160405180910390fd5b828260025f6103b86103b3368a90038a018a610cd7565b61090d565b81526020019081526020015f2091826103d2929190610e13565b50836040516103e19190610ecd565b60405180910390207f209e95fbe8dd14c5e1fbf791ee0a83234f45f20cb85504c7068d5ca0d6224588848460405161041a929190610f2c565b60405180910390a250505050565b5f5f6104338461049a565b90505f811161045557604051637a31a0a160e11b815260040160405180910390fd5b610460600182610f53565b9092149392505050565b60606104957f0000000000000000000000000000000000000000000000000000000000000000610a1a565b905090565b5f60015f6104a78461090d565b815260208101919091526040015f205492915050565b5f6104cb6020840184610da3565b6104d481610970565b6104f15760405163932d94f760e01b815260040160405180910390fd5b5f6001816105076103b336899003890189610cd7565b81526020019081526020015f20905060025f61052d878036038101906103b39190610cd7565b81526020019081526020015f20805461054590610d71565b90505f036105665760405163413e6e5760e11b815260040160405180910390fd5b6105766040850160208601610f66565b63ffffffff16158061059d5750426105946040860160208701610f66565b63ffffffff1610155b6105ba5760405163325ec75f60e01b815260040160405180910390fd5b80546001810182555f82815260209020909350849060028502016105de82826110a4565b505082856040516105ef9190610ecd565b60405180910390207f2decd15222f7c4a8c3d4d2e14dcfdc5a0b52eb2d4b81796bfd010ee5cd972fd386604051610626919061122f565b60405180910390a3505092915050565b5f5f60015f6106448561090d565b81526020019081526020015f2090505f81805490501161067757604051637a31a0a160e11b815260040160405180910390fd5b80545f9061068790600190610f53565b905081818154811061069b5761069b610d5d565b5f91825260209091206001600290920201015463ffffffff16949350505050565b606060025f6106ca8461090d565b81526020019081526020015f2080546106e290610d71565b80601f016020809104026020016040519081016040528092919081815260200182805461070e90610d71565b80156107595780601f1061073057610100808354040283529160200191610759565b820191905f5260205f20905b81548152906001019060200180831161073c57829003601f168201915b50505050509050919050565b60408051808201909152606081525f60208201819052905f60015f6107898661090d565b81526020019081526020015f2090505f8180549050116107bc57604051637a31a0a160e11b815260040160405180910390fd5b80545f906107cc90600190610f53565b9050808282815481106107e1576107e1610d5d565b905f5260205f209060020201806040518060400160405290815f8201805480602002602001604051908101604052809291908181526020015f905b828210156108e5578382905f5260205f2090600202016040518060400160405290815f820154815260200160018201805461085690610d71565b80601f016020809104026020016040519081016040528092919081815260200182805461088290610d71565b80156108cd5780601f106108a4576101008083540402835291602001916108cd565b820191905f5260205f20905b8154815290600101906020018083116108b057829003601f168201915b5050505050815250508152602001906001019061081c565b505050908252506001919091015463ffffffff16602090910152919791965090945050505050565b5f815f0151826020015163ffffffff1660405160200161095892919060609290921b6bffffffffffffffffffffffff1916825260a01b6001600160a01b031916601482015260200190565b60405160208183030381529060405261034290611356565b604051631beb2b9760e31b81526001600160a01b0382811660048301523360248301523060448301525f80356001600160e01b0319166064840152917f00000000000000000000000000000000000000000000000000000000000000009091169063df595cb8906084016020604051808303815f875af11580156109f6573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103429190611379565b60605f610a2683610a57565b6040805160208082528183019092529192505f91906020820181803683375050509182525060208101929092525090565b5f60ff8216601f81111561034257604051632cd44ac360e21b815260040160405180910390fd5b634e487b7160e01b5f52604160045260245ffd5b80356001600160a01b0381168114610aa8575f5ffd5b919050565b63ffffffff81168114610abe575f5ffd5b50565b8035610aa881610aad565b5f60408284031215610adc575f5ffd5b6040516040810181811067ffffffffffffffff82111715610aff57610aff610a7e565b604052905080610b0e83610a92565b81526020830135610b1e81610aad565b6020919091015292915050565b5f5f60608385031215610b3c575f5ffd5b610b468484610acc565b946040939093013593505050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b5f6040830182516040855281815180845260608701915060608160051b88010193506020830192505f5b81811015610bf457605f198886030183528351805186526020810151905060406020870152610bde6040870182610b54565b9550506020938401939290920191600101610bac565b505050506020830151610c0f602086018263ffffffff169052565b509392505050565b602081525f610c296020830184610b82565b9392505050565b5f60408284031215610c40575f5ffd5b50919050565b5f5f5f60608486031215610c58575f5ffd5b610c628585610c30565b9250604084013567ffffffffffffffff811115610c7d575f5ffd5b8401601f81018613610c8d575f5ffd5b803567ffffffffffffffff811115610ca3575f5ffd5b866020828401011115610cb4575f5ffd5b939660209190910195509293505050565b602081525f610c296020830184610b54565b5f60408284031215610ce7575f5ffd5b610c298383610acc565b5f5f60608385031215610d02575f5ffd5b610d0c8484610c30565b9150604083013567ffffffffffffffff811115610d27575f5ffd5b610d3385828601610c30565b9150509250929050565b828152604060208201525f610d556040830184610b82565b949350505050565b634e487b7160e01b5f52603260045260245ffd5b600181811c90821680610d8557607f821691505b602082108103610c4057634e487b7160e01b5f52602260045260245ffd5b5f60208284031215610db3575f5ffd5b610c2982610a92565b5b81811015610dd0575f8155600101610dbd565b5050565b601f821115610e0e57805f5260205f20601f840160051c81016020851015610df95750805b610e0b601f850160051c830182610dbc565b50505b505050565b67ffffffffffffffff831115610e2b57610e2b610a7e565b610e3f83610e398354610d71565b83610dd4565b5f601f841160018114610e70575f8515610e595750838201355b5f19600387901b1c1916600186901b178355610e0b565b5f83815260208120601f198716915b82811015610e9f5786850135825560209485019460019092019101610e7f565b5086821015610ebb575f1960f88860031b161c19848701351681555b505060018560011b0183555050505050565b6001600160a01b03610ede83610a92565b1681525f6020830135610ef081610aad565b63ffffffff16602083015250604001919050565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b602081525f610d55602083018486610f04565b634e487b7160e01b5f52601160045260245ffd5b8181038181111561034257610342610f3f565b5f60208284031215610f76575f5ffd5b8135610c2981610aad565b5f8235603e19833603018112610f95575f5ffd5b9190910192915050565b81358155600181016020830135601e19843603018112610fbd575f5ffd5b8301803567ffffffffffffffff81118015610fd6575f5ffd5b813603602084011315610fe7575f5ffd5b5f905050610fff81610ff98554610d71565b85610dd4565b5f601f821160018114611033575f831561101c5750838201602001355b5f19600385901b1c1916600184901b17855561108f565b5f85815260208120601f198516915b8281101561106457602085880181013583559485019460019092019101611042565b5084821015611083575f1960f88660031b161c19602085880101351681555b505060018360011b0185555b50505050505050565b5f813561034281610aad565b8135601e198336030181126110b7575f5ffd5b8201803567ffffffffffffffff8111156110cf575f5ffd5b6020820191508060051b36038213156110e6575f5ffd5b680100000000000000008111156110ff576110ff610a7e565b8254818455808210156111c3576001600160ff1b038116811461112457611124610f3f565b6001600160ff1b038216821461113c5761113c610f3f565b835f5260205f208160011b81018360011b820191505b808210156111c0575f825560018201805461116c90610d71565b80156111b357601f811160018114611186575f83556111b1565b5f838152602090206111a3601f840160051c820160018301610dbc565b505f83815260208120818555555b505b5050600282019150611152565b50505b505f8381526020812083915b83811015611200576111ea6111e48487610f81565b83610f9f565b60209290920191600291909101906001016111cf565b5050505050610dd061121460208401611098565b6001830163ffffffff821663ffffffff198254161781555050565b602081525f606082018335601e1985360301811261124b575f5ffd5b840180356020820167ffffffffffffffff821115611267575f5ffd5b8160051b803603821315611279575f5ffd5b604060208801529382905260809386018401935f908701605e1936869003015b8483101561133357888703607f1901825283358181126112b7575f5ffd5b860160208101358852604081013536829003603e190181126112d7575f5ffd5b0160408101906020013567ffffffffffffffff8111156112f5575f5ffd5b803603821315611303575f5ffd5b604060208a015261131860408a018284610f04565b98505050602084019350602082019150600183019250611299565b50505050505061134560208501610ac1565b63ffffffff81166040850152610c0f565b80516020808301519190811015610c40575f1960209190910360031b1b16919050565b5f60208284031215611389575f5ffd5b81518015158114610c29575f5ffdfea26469706673582212203379f2249603e9a59875b615be4754146f886916fc025c19fa13230b3b40e45864736f6c634300081b0033",
}

// ReleaseManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ReleaseManagerMetaData.ABI instead.
var ReleaseManagerABI = ReleaseManagerMetaData.ABI

// ReleaseManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ReleaseManagerMetaData.Bin instead.
var ReleaseManagerBin = ReleaseManagerMetaData.Bin

// DeployReleaseManager deploys a new Ethereum contract, binding an instance of ReleaseManager to it.
func DeployReleaseManager(auth *bind.TransactOpts, backend bind.ContractBackend, _permissionController common.Address, _version string) (common.Address, *types.Transaction, *ReleaseManager, error) {
	parsed, err := ReleaseManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ReleaseManagerBin), backend, _permissionController, _version)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ReleaseManager{ReleaseManagerCaller: ReleaseManagerCaller{contract: contract}, ReleaseManagerTransactor: ReleaseManagerTransactor{contract: contract}, ReleaseManagerFilterer: ReleaseManagerFilterer{contract: contract}}, nil
}

// ReleaseManager is an auto generated Go binding around an Ethereum contract.
type ReleaseManager struct {
	ReleaseManagerCaller     // Read-only binding to the contract
	ReleaseManagerTransactor // Write-only binding to the contract
	ReleaseManagerFilterer   // Log filterer for contract events
}

// ReleaseManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReleaseManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReleaseManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReleaseManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReleaseManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReleaseManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReleaseManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReleaseManagerSession struct {
	Contract     *ReleaseManager   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReleaseManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReleaseManagerCallerSession struct {
	Contract *ReleaseManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ReleaseManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReleaseManagerTransactorSession struct {
	Contract     *ReleaseManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ReleaseManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReleaseManagerRaw struct {
	Contract *ReleaseManager // Generic contract binding to access the raw methods on
}

// ReleaseManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReleaseManagerCallerRaw struct {
	Contract *ReleaseManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ReleaseManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReleaseManagerTransactorRaw struct {
	Contract *ReleaseManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReleaseManager creates a new instance of ReleaseManager, bound to a specific deployed contract.
func NewReleaseManager(address common.Address, backend bind.ContractBackend) (*ReleaseManager, error) {
	contract, err := bindReleaseManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReleaseManager{ReleaseManagerCaller: ReleaseManagerCaller{contract: contract}, ReleaseManagerTransactor: ReleaseManagerTransactor{contract: contract}, ReleaseManagerFilterer: ReleaseManagerFilterer{contract: contract}}, nil
}

// NewReleaseManagerCaller creates a new read-only instance of ReleaseManager, bound to a specific deployed contract.
func NewReleaseManagerCaller(address common.Address, caller bind.ContractCaller) (*ReleaseManagerCaller, error) {
	contract, err := bindReleaseManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReleaseManagerCaller{contract: contract}, nil
}

// NewReleaseManagerTransactor creates a new write-only instance of ReleaseManager, bound to a specific deployed contract.
func NewReleaseManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ReleaseManagerTransactor, error) {
	contract, err := bindReleaseManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReleaseManagerTransactor{contract: contract}, nil
}

// NewReleaseManagerFilterer creates a new log filterer instance of ReleaseManager, bound to a specific deployed contract.
func NewReleaseManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ReleaseManagerFilterer, error) {
	contract, err := bindReleaseManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReleaseManagerFilterer{contract: contract}, nil
}

// bindReleaseManager binds a generic wrapper to an already deployed contract.
func bindReleaseManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ReleaseManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReleaseManager *ReleaseManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReleaseManager.Contract.ReleaseManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReleaseManager *ReleaseManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReleaseManager.Contract.ReleaseManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReleaseManager *ReleaseManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReleaseManager.Contract.ReleaseManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReleaseManager *ReleaseManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReleaseManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReleaseManager *ReleaseManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReleaseManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReleaseManager *ReleaseManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReleaseManager.Contract.contract.Transact(opts, method, params...)
}

// GetLatestRelease is a free data retrieval call binding the contract method 0xd30eeb88.
//
// Solidity: function getLatestRelease((address,uint32) operatorSet) view returns(uint256, ((bytes32,string)[],uint32))
func (_ReleaseManager *ReleaseManagerCaller) GetLatestRelease(opts *bind.CallOpts, operatorSet OperatorSet) (*big.Int, IReleaseManagerTypesRelease, error) {
	var out []interface{}
	err := _ReleaseManager.contract.Call(opts, &out, "getLatestRelease", operatorSet)

	if err != nil {
		return *new(*big.Int), *new(IReleaseManagerTypesRelease), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(IReleaseManagerTypesRelease)).(*IReleaseManagerTypesRelease)

	return out0, out1, err

}

// GetLatestRelease is a free data retrieval call binding the contract method 0xd30eeb88.
//
// Solidity: function getLatestRelease((address,uint32) operatorSet) view returns(uint256, ((bytes32,string)[],uint32))
func (_ReleaseManager *ReleaseManagerSession) GetLatestRelease(operatorSet OperatorSet) (*big.Int, IReleaseManagerTypesRelease, error) {
	return _ReleaseManager.Contract.GetLatestRelease(&_ReleaseManager.CallOpts, operatorSet)
}

// GetLatestRelease is a free data retrieval call binding the contract method 0xd30eeb88.
//
// Solidity: function getLatestRelease((address,uint32) operatorSet) view returns(uint256, ((bytes32,string)[],uint32))
func (_ReleaseManager *ReleaseManagerCallerSession) GetLatestRelease(operatorSet OperatorSet) (*big.Int, IReleaseManagerTypesRelease, error) {
	return _ReleaseManager.Contract.GetLatestRelease(&_ReleaseManager.CallOpts, operatorSet)
}

// GetLatestUpgradeByTime is a free data retrieval call binding the contract method 0xa9e0ed68.
//
// Solidity: function getLatestUpgradeByTime((address,uint32) operatorSet) view returns(uint32)
func (_ReleaseManager *ReleaseManagerCaller) GetLatestUpgradeByTime(opts *bind.CallOpts, operatorSet OperatorSet) (uint32, error) {
	var out []interface{}
	err := _ReleaseManager.contract.Call(opts, &out, "getLatestUpgradeByTime", operatorSet)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetLatestUpgradeByTime is a free data retrieval call binding the contract method 0xa9e0ed68.
//
// Solidity: function getLatestUpgradeByTime((address,uint32) operatorSet) view returns(uint32)
func (_ReleaseManager *ReleaseManagerSession) GetLatestUpgradeByTime(operatorSet OperatorSet) (uint32, error) {
	return _ReleaseManager.Contract.GetLatestUpgradeByTime(&_ReleaseManager.CallOpts, operatorSet)
}

// GetLatestUpgradeByTime is a free data retrieval call binding the contract method 0xa9e0ed68.
//
// Solidity: function getLatestUpgradeByTime((address,uint32) operatorSet) view returns(uint32)
func (_ReleaseManager *ReleaseManagerCallerSession) GetLatestUpgradeByTime(operatorSet OperatorSet) (uint32, error) {
	return _ReleaseManager.Contract.GetLatestUpgradeByTime(&_ReleaseManager.CallOpts, operatorSet)
}

// GetMetadataURI is a free data retrieval call binding the contract method 0xb053b56d.
//
// Solidity: function getMetadataURI((address,uint32) operatorSet) view returns(string)
func (_ReleaseManager *ReleaseManagerCaller) GetMetadataURI(opts *bind.CallOpts, operatorSet OperatorSet) (string, error) {
	var out []interface{}
	err := _ReleaseManager.contract.Call(opts, &out, "getMetadataURI", operatorSet)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetMetadataURI is a free data retrieval call binding the contract method 0xb053b56d.
//
// Solidity: function getMetadataURI((address,uint32) operatorSet) view returns(string)
func (_ReleaseManager *ReleaseManagerSession) GetMetadataURI(operatorSet OperatorSet) (string, error) {
	return _ReleaseManager.Contract.GetMetadataURI(&_ReleaseManager.CallOpts, operatorSet)
}

// GetMetadataURI is a free data retrieval call binding the contract method 0xb053b56d.
//
// Solidity: function getMetadataURI((address,uint32) operatorSet) view returns(string)
func (_ReleaseManager *ReleaseManagerCallerSession) GetMetadataURI(operatorSet OperatorSet) (string, error) {
	return _ReleaseManager.Contract.GetMetadataURI(&_ReleaseManager.CallOpts, operatorSet)
}

// GetRelease is a free data retrieval call binding the contract method 0x3acab5fc.
//
// Solidity: function getRelease((address,uint32) operatorSet, uint256 releaseId) view returns(((bytes32,string)[],uint32))
func (_ReleaseManager *ReleaseManagerCaller) GetRelease(opts *bind.CallOpts, operatorSet OperatorSet, releaseId *big.Int) (IReleaseManagerTypesRelease, error) {
	var out []interface{}
	err := _ReleaseManager.contract.Call(opts, &out, "getRelease", operatorSet, releaseId)

	if err != nil {
		return *new(IReleaseManagerTypesRelease), err
	}

	out0 := *abi.ConvertType(out[0], new(IReleaseManagerTypesRelease)).(*IReleaseManagerTypesRelease)

	return out0, err

}

// GetRelease is a free data retrieval call binding the contract method 0x3acab5fc.
//
// Solidity: function getRelease((address,uint32) operatorSet, uint256 releaseId) view returns(((bytes32,string)[],uint32))
func (_ReleaseManager *ReleaseManagerSession) GetRelease(operatorSet OperatorSet, releaseId *big.Int) (IReleaseManagerTypesRelease, error) {
	return _ReleaseManager.Contract.GetRelease(&_ReleaseManager.CallOpts, operatorSet, releaseId)
}

// GetRelease is a free data retrieval call binding the contract method 0x3acab5fc.
//
// Solidity: function getRelease((address,uint32) operatorSet, uint256 releaseId) view returns(((bytes32,string)[],uint32))
func (_ReleaseManager *ReleaseManagerCallerSession) GetRelease(operatorSet OperatorSet, releaseId *big.Int) (IReleaseManagerTypesRelease, error) {
	return _ReleaseManager.Contract.GetRelease(&_ReleaseManager.CallOpts, operatorSet, releaseId)
}

// GetTotalReleases is a free data retrieval call binding the contract method 0x66f409f7.
//
// Solidity: function getTotalReleases((address,uint32) operatorSet) view returns(uint256)
func (_ReleaseManager *ReleaseManagerCaller) GetTotalReleases(opts *bind.CallOpts, operatorSet OperatorSet) (*big.Int, error) {
	var out []interface{}
	err := _ReleaseManager.contract.Call(opts, &out, "getTotalReleases", operatorSet)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalReleases is a free data retrieval call binding the contract method 0x66f409f7.
//
// Solidity: function getTotalReleases((address,uint32) operatorSet) view returns(uint256)
func (_ReleaseManager *ReleaseManagerSession) GetTotalReleases(operatorSet OperatorSet) (*big.Int, error) {
	return _ReleaseManager.Contract.GetTotalReleases(&_ReleaseManager.CallOpts, operatorSet)
}

// GetTotalReleases is a free data retrieval call binding the contract method 0x66f409f7.
//
// Solidity: function getTotalReleases((address,uint32) operatorSet) view returns(uint256)
func (_ReleaseManager *ReleaseManagerCallerSession) GetTotalReleases(operatorSet OperatorSet) (*big.Int, error) {
	return _ReleaseManager.Contract.GetTotalReleases(&_ReleaseManager.CallOpts, operatorSet)
}

// IsValidRelease is a free data retrieval call binding the contract method 0x517e4068.
//
// Solidity: function isValidRelease((address,uint32) operatorSet, uint256 releaseId) view returns(bool)
func (_ReleaseManager *ReleaseManagerCaller) IsValidRelease(opts *bind.CallOpts, operatorSet OperatorSet, releaseId *big.Int) (bool, error) {
	var out []interface{}
	err := _ReleaseManager.contract.Call(opts, &out, "isValidRelease", operatorSet, releaseId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidRelease is a free data retrieval call binding the contract method 0x517e4068.
//
// Solidity: function isValidRelease((address,uint32) operatorSet, uint256 releaseId) view returns(bool)
func (_ReleaseManager *ReleaseManagerSession) IsValidRelease(operatorSet OperatorSet, releaseId *big.Int) (bool, error) {
	return _ReleaseManager.Contract.IsValidRelease(&_ReleaseManager.CallOpts, operatorSet, releaseId)
}

// IsValidRelease is a free data retrieval call binding the contract method 0x517e4068.
//
// Solidity: function isValidRelease((address,uint32) operatorSet, uint256 releaseId) view returns(bool)
func (_ReleaseManager *ReleaseManagerCallerSession) IsValidRelease(operatorSet OperatorSet, releaseId *big.Int) (bool, error) {
	return _ReleaseManager.Contract.IsValidRelease(&_ReleaseManager.CallOpts, operatorSet, releaseId)
}

// PermissionController is a free data retrieval call binding the contract method 0x4657e26a.
//
// Solidity: function permissionController() view returns(address)
func (_ReleaseManager *ReleaseManagerCaller) PermissionController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ReleaseManager.contract.Call(opts, &out, "permissionController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PermissionController is a free data retrieval call binding the contract method 0x4657e26a.
//
// Solidity: function permissionController() view returns(address)
func (_ReleaseManager *ReleaseManagerSession) PermissionController() (common.Address, error) {
	return _ReleaseManager.Contract.PermissionController(&_ReleaseManager.CallOpts)
}

// PermissionController is a free data retrieval call binding the contract method 0x4657e26a.
//
// Solidity: function permissionController() view returns(address)
func (_ReleaseManager *ReleaseManagerCallerSession) PermissionController() (common.Address, error) {
	return _ReleaseManager.Contract.PermissionController(&_ReleaseManager.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ReleaseManager *ReleaseManagerCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ReleaseManager.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ReleaseManager *ReleaseManagerSession) Version() (string, error) {
	return _ReleaseManager.Contract.Version(&_ReleaseManager.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ReleaseManager *ReleaseManagerCallerSession) Version() (string, error) {
	return _ReleaseManager.Contract.Version(&_ReleaseManager.CallOpts)
}

// PublishMetadataURI is a paid mutator transaction binding the contract method 0x4840a67c.
//
// Solidity: function publishMetadataURI((address,uint32) operatorSet, string metadataURI) returns()
func (_ReleaseManager *ReleaseManagerTransactor) PublishMetadataURI(opts *bind.TransactOpts, operatorSet OperatorSet, metadataURI string) (*types.Transaction, error) {
	return _ReleaseManager.contract.Transact(opts, "publishMetadataURI", operatorSet, metadataURI)
}

// PublishMetadataURI is a paid mutator transaction binding the contract method 0x4840a67c.
//
// Solidity: function publishMetadataURI((address,uint32) operatorSet, string metadataURI) returns()
func (_ReleaseManager *ReleaseManagerSession) PublishMetadataURI(operatorSet OperatorSet, metadataURI string) (*types.Transaction, error) {
	return _ReleaseManager.Contract.PublishMetadataURI(&_ReleaseManager.TransactOpts, operatorSet, metadataURI)
}

// PublishMetadataURI is a paid mutator transaction binding the contract method 0x4840a67c.
//
// Solidity: function publishMetadataURI((address,uint32) operatorSet, string metadataURI) returns()
func (_ReleaseManager *ReleaseManagerTransactorSession) PublishMetadataURI(operatorSet OperatorSet, metadataURI string) (*types.Transaction, error) {
	return _ReleaseManager.Contract.PublishMetadataURI(&_ReleaseManager.TransactOpts, operatorSet, metadataURI)
}

// PublishRelease is a paid mutator transaction binding the contract method 0x7c09ea82.
//
// Solidity: function publishRelease((address,uint32) operatorSet, ((bytes32,string)[],uint32) release) returns(uint256 releaseId)
func (_ReleaseManager *ReleaseManagerTransactor) PublishRelease(opts *bind.TransactOpts, operatorSet OperatorSet, release IReleaseManagerTypesRelease) (*types.Transaction, error) {
	return _ReleaseManager.contract.Transact(opts, "publishRelease", operatorSet, release)
}

// PublishRelease is a paid mutator transaction binding the contract method 0x7c09ea82.
//
// Solidity: function publishRelease((address,uint32) operatorSet, ((bytes32,string)[],uint32) release) returns(uint256 releaseId)
func (_ReleaseManager *ReleaseManagerSession) PublishRelease(operatorSet OperatorSet, release IReleaseManagerTypesRelease) (*types.Transaction, error) {
	return _ReleaseManager.Contract.PublishRelease(&_ReleaseManager.TransactOpts, operatorSet, release)
}

// PublishRelease is a paid mutator transaction binding the contract method 0x7c09ea82.
//
// Solidity: function publishRelease((address,uint32) operatorSet, ((bytes32,string)[],uint32) release) returns(uint256 releaseId)
func (_ReleaseManager *ReleaseManagerTransactorSession) PublishRelease(operatorSet OperatorSet, release IReleaseManagerTypesRelease) (*types.Transaction, error) {
	return _ReleaseManager.Contract.PublishRelease(&_ReleaseManager.TransactOpts, operatorSet, release)
}

// ReleaseManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ReleaseManager contract.
type ReleaseManagerInitializedIterator struct {
	Event *ReleaseManagerInitialized // Event containing the contract specifics and raw log

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
func (it *ReleaseManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReleaseManagerInitialized)
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
		it.Event = new(ReleaseManagerInitialized)
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
func (it *ReleaseManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReleaseManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReleaseManagerInitialized represents a Initialized event raised by the ReleaseManager contract.
type ReleaseManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ReleaseManager *ReleaseManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ReleaseManagerInitializedIterator, error) {

	logs, sub, err := _ReleaseManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ReleaseManagerInitializedIterator{contract: _ReleaseManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ReleaseManager *ReleaseManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ReleaseManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ReleaseManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReleaseManagerInitialized)
				if err := _ReleaseManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ReleaseManager *ReleaseManagerFilterer) ParseInitialized(log types.Log) (*ReleaseManagerInitialized, error) {
	event := new(ReleaseManagerInitialized)
	if err := _ReleaseManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReleaseManagerMetadataURIPublishedIterator is returned from FilterMetadataURIPublished and is used to iterate over the raw logs and unpacked data for MetadataURIPublished events raised by the ReleaseManager contract.
type ReleaseManagerMetadataURIPublishedIterator struct {
	Event *ReleaseManagerMetadataURIPublished // Event containing the contract specifics and raw log

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
func (it *ReleaseManagerMetadataURIPublishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReleaseManagerMetadataURIPublished)
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
		it.Event = new(ReleaseManagerMetadataURIPublished)
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
func (it *ReleaseManagerMetadataURIPublishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReleaseManagerMetadataURIPublishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReleaseManagerMetadataURIPublished represents a MetadataURIPublished event raised by the ReleaseManager contract.
type ReleaseManagerMetadataURIPublished struct {
	OperatorSet OperatorSet
	MetadataURI string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMetadataURIPublished is a free log retrieval operation binding the contract event 0x209e95fbe8dd14c5e1fbf791ee0a83234f45f20cb85504c7068d5ca0d6224588.
//
// Solidity: event MetadataURIPublished((address,uint32) indexed operatorSet, string metadataURI)
func (_ReleaseManager *ReleaseManagerFilterer) FilterMetadataURIPublished(opts *bind.FilterOpts, operatorSet []OperatorSet) (*ReleaseManagerMetadataURIPublishedIterator, error) {

	var operatorSetRule []interface{}
	for _, operatorSetItem := range operatorSet {
		operatorSetRule = append(operatorSetRule, operatorSetItem)
	}

	logs, sub, err := _ReleaseManager.contract.FilterLogs(opts, "MetadataURIPublished", operatorSetRule)
	if err != nil {
		return nil, err
	}
	return &ReleaseManagerMetadataURIPublishedIterator{contract: _ReleaseManager.contract, event: "MetadataURIPublished", logs: logs, sub: sub}, nil
}

// WatchMetadataURIPublished is a free log subscription operation binding the contract event 0x209e95fbe8dd14c5e1fbf791ee0a83234f45f20cb85504c7068d5ca0d6224588.
//
// Solidity: event MetadataURIPublished((address,uint32) indexed operatorSet, string metadataURI)
func (_ReleaseManager *ReleaseManagerFilterer) WatchMetadataURIPublished(opts *bind.WatchOpts, sink chan<- *ReleaseManagerMetadataURIPublished, operatorSet []OperatorSet) (event.Subscription, error) {

	var operatorSetRule []interface{}
	for _, operatorSetItem := range operatorSet {
		operatorSetRule = append(operatorSetRule, operatorSetItem)
	}

	logs, sub, err := _ReleaseManager.contract.WatchLogs(opts, "MetadataURIPublished", operatorSetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReleaseManagerMetadataURIPublished)
				if err := _ReleaseManager.contract.UnpackLog(event, "MetadataURIPublished", log); err != nil {
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

// ParseMetadataURIPublished is a log parse operation binding the contract event 0x209e95fbe8dd14c5e1fbf791ee0a83234f45f20cb85504c7068d5ca0d6224588.
//
// Solidity: event MetadataURIPublished((address,uint32) indexed operatorSet, string metadataURI)
func (_ReleaseManager *ReleaseManagerFilterer) ParseMetadataURIPublished(log types.Log) (*ReleaseManagerMetadataURIPublished, error) {
	event := new(ReleaseManagerMetadataURIPublished)
	if err := _ReleaseManager.contract.UnpackLog(event, "MetadataURIPublished", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReleaseManagerReleasePublishedIterator is returned from FilterReleasePublished and is used to iterate over the raw logs and unpacked data for ReleasePublished events raised by the ReleaseManager contract.
type ReleaseManagerReleasePublishedIterator struct {
	Event *ReleaseManagerReleasePublished // Event containing the contract specifics and raw log

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
func (it *ReleaseManagerReleasePublishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReleaseManagerReleasePublished)
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
		it.Event = new(ReleaseManagerReleasePublished)
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
func (it *ReleaseManagerReleasePublishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReleaseManagerReleasePublishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReleaseManagerReleasePublished represents a ReleasePublished event raised by the ReleaseManager contract.
type ReleaseManagerReleasePublished struct {
	OperatorSet OperatorSet
	ReleaseId   *big.Int
	Release     IReleaseManagerTypesRelease
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterReleasePublished is a free log retrieval operation binding the contract event 0x2decd15222f7c4a8c3d4d2e14dcfdc5a0b52eb2d4b81796bfd010ee5cd972fd3.
//
// Solidity: event ReleasePublished((address,uint32) indexed operatorSet, uint256 indexed releaseId, ((bytes32,string)[],uint32) release)
func (_ReleaseManager *ReleaseManagerFilterer) FilterReleasePublished(opts *bind.FilterOpts, operatorSet []OperatorSet, releaseId []*big.Int) (*ReleaseManagerReleasePublishedIterator, error) {

	var operatorSetRule []interface{}
	for _, operatorSetItem := range operatorSet {
		operatorSetRule = append(operatorSetRule, operatorSetItem)
	}
	var releaseIdRule []interface{}
	for _, releaseIdItem := range releaseId {
		releaseIdRule = append(releaseIdRule, releaseIdItem)
	}

	logs, sub, err := _ReleaseManager.contract.FilterLogs(opts, "ReleasePublished", operatorSetRule, releaseIdRule)
	if err != nil {
		return nil, err
	}
	return &ReleaseManagerReleasePublishedIterator{contract: _ReleaseManager.contract, event: "ReleasePublished", logs: logs, sub: sub}, nil
}

// WatchReleasePublished is a free log subscription operation binding the contract event 0x2decd15222f7c4a8c3d4d2e14dcfdc5a0b52eb2d4b81796bfd010ee5cd972fd3.
//
// Solidity: event ReleasePublished((address,uint32) indexed operatorSet, uint256 indexed releaseId, ((bytes32,string)[],uint32) release)
func (_ReleaseManager *ReleaseManagerFilterer) WatchReleasePublished(opts *bind.WatchOpts, sink chan<- *ReleaseManagerReleasePublished, operatorSet []OperatorSet, releaseId []*big.Int) (event.Subscription, error) {

	var operatorSetRule []interface{}
	for _, operatorSetItem := range operatorSet {
		operatorSetRule = append(operatorSetRule, operatorSetItem)
	}
	var releaseIdRule []interface{}
	for _, releaseIdItem := range releaseId {
		releaseIdRule = append(releaseIdRule, releaseIdItem)
	}

	logs, sub, err := _ReleaseManager.contract.WatchLogs(opts, "ReleasePublished", operatorSetRule, releaseIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReleaseManagerReleasePublished)
				if err := _ReleaseManager.contract.UnpackLog(event, "ReleasePublished", log); err != nil {
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

// ParseReleasePublished is a log parse operation binding the contract event 0x2decd15222f7c4a8c3d4d2e14dcfdc5a0b52eb2d4b81796bfd010ee5cd972fd3.
//
// Solidity: event ReleasePublished((address,uint32) indexed operatorSet, uint256 indexed releaseId, ((bytes32,string)[],uint32) release)
func (_ReleaseManager *ReleaseManagerFilterer) ParseReleasePublished(log types.Log) (*ReleaseManagerReleasePublished, error) {
	event := new(ReleaseManagerReleasePublished)
	if err := _ReleaseManager.contract.UnpackLog(event, "ReleasePublished", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
