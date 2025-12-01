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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_permissionController\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getLatestRelease\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIReleaseManagerTypes.Release\",\"components\":[{\"name\":\"artifacts\",\"type\":\"tuple[]\",\"internalType\":\"structIReleaseManagerTypes.Artifact[]\",\"components\":[{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"registry\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"upgradeByTime\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLatestUpgradeByTime\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMetadataURI\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRelease\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"releaseId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIReleaseManagerTypes.Release\",\"components\":[{\"name\":\"artifacts\",\"type\":\"tuple[]\",\"internalType\":\"structIReleaseManagerTypes.Artifact[]\",\"components\":[{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"registry\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"upgradeByTime\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalReleases\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isValidRelease\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"releaseId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"permissionController\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"publishMetadataURI\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"publishRelease\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"release\",\"type\":\"tuple\",\"internalType\":\"structIReleaseManagerTypes.Release\",\"components\":[{\"name\":\"artifacts\",\"type\":\"tuple[]\",\"internalType\":\"structIReleaseManagerTypes.Artifact[]\",\"components\":[{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"registry\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"upgradeByTime\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"releaseId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MetadataURIPublished\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":true,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"metadataURI\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasePublished\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":true,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"releaseId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"release\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIReleaseManagerTypes.Release\",\"components\":[{\"name\":\"artifacts\",\"type\":\"tuple[]\",\"internalType\":\"structIReleaseManagerTypes.Artifact[]\",\"components\":[{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"registry\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"upgradeByTime\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidMetadataURI\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPermissions\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidUpgradeByTime\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustPublishMetadataURI\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoReleases\",\"inputs\":[]}]",
	Bin: "0x60a060405234801561000f575f5ffd5b5060405161146238038061146283398101604081905261002e91610105565b6001600160a01b038116608052610043610049565b50610132565b5f54610100900460ff16156100b45760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff90811614610103575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b5f60208284031215610115575f5ffd5b81516001600160a01b038116811461012b575f5ffd5b9392505050565b6080516113126101505f395f818160c2015261095a01526113125ff3fe608060405234801561000f575f5ffd5b5060043610610090575f3560e01c806366f409f71161006357806366f409f7146101345780637c09ea8214610155578063a9e0ed6814610168578063b053b56d14610190578063d30eeb88146101b0575f5ffd5b80633acab5fc146100945780634657e26a146100bd5780634840a67c146100fc578063517e406814610111575b5f5ffd5b6100a76100a2366004610a6f565b6101d1565b6040516100b49190610b5b565b60405180910390f35b6100e47f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016100b4565b61010f61010a366004610b8a565b610335565b005b61012461011f366004610a6f565b6103f8565b60405190151581526020016100b4565b610147610142366004610c09565b61043a565b6040519081526020016100b4565b610147610163366004610c23565b61045d565b61017b610176366004610c09565b6105b9565b60405163ffffffff90911681526020016100b4565b6101a361019e366004610c09565b61063f565b6040516100b49190610c6f565b6101c36101be366004610c09565b6106e8565b6040516100b4929190610c81565b60408051808201909152606081525f602082015260015f6101f185610890565b81526020019081526020015f20828154811061020f5761020f610ca1565b905f5260205f2090600202016040518060400160405290815f8201805480602002602001604051908101604052809291908181526020015f905b82821015610312578382905f5260205f2090600202016040518060400160405290815f820154815260200160018201805461028390610cb5565b80601f01602080910402602001604051908101604052809291908181526020018280546102af90610cb5565b80156102fa5780601f106102d1576101008083540402835291602001916102fa565b820191905f5260205f20905b8154815290600101906020018083116102dd57829003601f168201915b50505050508152505081526020019060010190610249565b505050908252506001919091015463ffffffff1660209091015290505b92915050565b6103426020840184610ce7565b61034b816108f3565b5f82900361036c57604051630eec403f60e41b815260040160405180910390fd5b828260025f610388610383368a90038a018a610c09565b610890565b81526020019081526020015f2091826103a2929190610d57565b50836040516103b19190610e11565b60405180910390207f209e95fbe8dd14c5e1fbf791ee0a83234f45f20cb85504c7068d5ca0d622458884846040516103ea929190610e70565b60405180910390a250505050565b5f5f6104038461043a565b90505f811161042557604051637a31a0a160e11b815260040160405180910390fd5b610430600182610e97565b9092149392505050565b5f60015f61044784610890565b815260208101919091526040015f205492915050565b5f61046b6020840184610ce7565b610474816108f3565b5f60018161048a61038336899003890189610c09565b81526020019081526020015f20905060025f6104b0878036038101906103839190610c09565b81526020019081526020015f2080546104c890610cb5565b90505f036104e95760405163413e6e5760e11b815260040160405180910390fd5b6104f96040850160208601610eaa565b63ffffffff1615806105205750426105176040860160208701610eaa565b63ffffffff1610155b61053d5760405163325ec75f60e01b815260040160405180910390fd5b80546001810182555f82815260209020909350849060028502016105618282610fe8565b505082856040516105729190610e11565b60405180910390207f2decd15222f7c4a8c3d4d2e14dcfdc5a0b52eb2d4b81796bfd010ee5cd972fd3866040516105a99190611173565b60405180910390a3505092915050565b5f5f60015f6105c785610890565b81526020019081526020015f2090505f8180549050116105fa57604051637a31a0a160e11b815260040160405180910390fd5b80545f9061060a90600190610e97565b905081818154811061061e5761061e610ca1565b5f91825260209091206001600290920201015463ffffffff16949350505050565b606060025f61064d84610890565b81526020019081526020015f20805461066590610cb5565b80601f016020809104026020016040519081016040528092919081815260200182805461069190610cb5565b80156106dc5780601f106106b3576101008083540402835291602001916106dc565b820191905f5260205f20905b8154815290600101906020018083116106bf57829003601f168201915b50505050509050919050565b60408051808201909152606081525f60208201819052905f60015f61070c86610890565b81526020019081526020015f2090505f81805490501161073f57604051637a31a0a160e11b815260040160405180910390fd5b80545f9061074f90600190610e97565b90508082828154811061076457610764610ca1565b905f5260205f209060020201806040518060400160405290815f8201805480602002602001604051908101604052809291908181526020015f905b82821015610868578382905f5260205f2090600202016040518060400160405290815f82015481526020016001820180546107d990610cb5565b80601f016020809104026020016040519081016040528092919081815260200182805461080590610cb5565b80156108505780601f1061082757610100808354040283529160200191610850565b820191905f5260205f20905b81548152906001019060200180831161083357829003601f168201915b5050505050815250508152602001906001019061079f565b505050908252506001919091015463ffffffff16602090910152919791965090945050505050565b5f815f0151826020015163ffffffff166040516020016108db92919060609290921b6bffffffffffffffffffffffff1916825260a01b6001600160a01b031916601482015260200190565b60405160208183030381529060405261032f9061129a565b6108fc8161091c565b6109195760405163932d94f760e01b815260040160405180910390fd5b50565b604051631beb2b9760e31b81526001600160a01b0382811660048301523360248301523060448301525f80356001600160e01b0319166064840152917f00000000000000000000000000000000000000000000000000000000000000009091169063df595cb890608401602060405180830381865afa1580156109a1573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061032f91906112bd565b634e487b7160e01b5f52604160045260245ffd5b80356001600160a01b03811681146109ef575f5ffd5b919050565b63ffffffff81168114610919575f5ffd5b80356109ef816109f4565b5f60408284031215610a20575f5ffd5b6040516040810181811067ffffffffffffffff82111715610a4357610a436109c5565b604052905080610a52836109d9565b81526020830135610a62816109f4565b6020919091015292915050565b5f5f60608385031215610a80575f5ffd5b610a8a8484610a10565b946040939093013593505050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b5f6040830182516040855281815180845260608701915060608160051b88010193506020830192505f5b81811015610b3857605f198886030183528351805186526020810151905060406020870152610b226040870182610a98565b9550506020938401939290920191600101610af0565b505050506020830151610b53602086018263ffffffff169052565b509392505050565b602081525f610b6d6020830184610ac6565b9392505050565b5f60408284031215610b84575f5ffd5b50919050565b5f5f5f60608486031215610b9c575f5ffd5b610ba68585610b74565b9250604084013567ffffffffffffffff811115610bc1575f5ffd5b8401601f81018613610bd1575f5ffd5b803567ffffffffffffffff811115610be7575f5ffd5b866020828401011115610bf8575f5ffd5b939660209190910195509293505050565b5f60408284031215610c19575f5ffd5b610b6d8383610a10565b5f5f60608385031215610c34575f5ffd5b610c3e8484610b74565b9150604083013567ffffffffffffffff811115610c59575f5ffd5b610c6585828601610b74565b9150509250929050565b602081525f610b6d6020830184610a98565b828152604060208201525f610c996040830184610ac6565b949350505050565b634e487b7160e01b5f52603260045260245ffd5b600181811c90821680610cc957607f821691505b602082108103610b8457634e487b7160e01b5f52602260045260245ffd5b5f60208284031215610cf7575f5ffd5b610b6d826109d9565b5b81811015610d14575f8155600101610d01565b5050565b601f821115610d5257805f5260205f20601f840160051c81016020851015610d3d5750805b610d4f601f850160051c830182610d00565b50505b505050565b67ffffffffffffffff831115610d6f57610d6f6109c5565b610d8383610d7d8354610cb5565b83610d18565b5f601f841160018114610db4575f8515610d9d5750838201355b5f19600387901b1c1916600186901b178355610d4f565b5f83815260208120601f198716915b82811015610de35786850135825560209485019460019092019101610dc3565b5086821015610dff575f1960f88860031b161c19848701351681555b505060018560011b0183555050505050565b6001600160a01b03610e22836109d9565b1681525f6020830135610e34816109f4565b63ffffffff16602083015250604001919050565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b602081525f610c99602083018486610e48565b634e487b7160e01b5f52601160045260245ffd5b8181038181111561032f5761032f610e83565b5f60208284031215610eba575f5ffd5b8135610b6d816109f4565b5f8235603e19833603018112610ed9575f5ffd5b9190910192915050565b81358155600181016020830135601e19843603018112610f01575f5ffd5b8301803567ffffffffffffffff81118015610f1a575f5ffd5b813603602084011315610f2b575f5ffd5b5f905050610f4381610f3d8554610cb5565b85610d18565b5f601f821160018114610f77575f8315610f605750838201602001355b5f19600385901b1c1916600184901b178555610fd3565b5f85815260208120601f198516915b82811015610fa857602085880181013583559485019460019092019101610f86565b5084821015610fc7575f1960f88660031b161c19602085880101351681555b505060018360011b0185555b50505050505050565b5f813561032f816109f4565b8135601e19833603018112610ffb575f5ffd5b8201803567ffffffffffffffff811115611013575f5ffd5b6020820191508060051b360382131561102a575f5ffd5b68010000000000000000811115611043576110436109c5565b825481845580821015611107576001600160ff1b038116811461106857611068610e83565b6001600160ff1b038216821461108057611080610e83565b835f5260205f208160011b81018360011b820191505b80821015611104575f82556001820180546110b090610cb5565b80156110f757601f8111600181146110ca575f83556110f5565b5f838152602090206110e7601f840160051c820160018301610d00565b505f83815260208120818555555b505b5050600282019150611096565b50505b505f8381526020812083915b838110156111445761112e6111288487610ec5565b83610ee3565b6020929092019160029190910190600101611113565b5050505050610d1461115860208401610fdc565b6001830163ffffffff821663ffffffff198254161781555050565b602081525f606082018335601e1985360301811261118f575f5ffd5b840180356020820167ffffffffffffffff8211156111ab575f5ffd5b8160051b8036038213156111bd575f5ffd5b604060208801529382905260809386018401935f908701605e1936869003015b8483101561127757888703607f1901825283358181126111fb575f5ffd5b860160208101358852604081013536829003603e1901811261121b575f5ffd5b0160408101906020013567ffffffffffffffff811115611239575f5ffd5b803603821315611247575f5ffd5b604060208a015261125c60408a018284610e48565b985050506020840193506020820191506001830192506111dd565b50505050505061128960208501610a05565b63ffffffff81166040850152610b53565b80516020808301519190811015610b84575f1960209190910360031b1b16919050565b5f602082840312156112cd575f5ffd5b81518015158114610b6d575f5ffdfea264697066735822122005612e05e63b4f904e703ae970fa68ac728005f1ec5704420bf10d481e10d8da64736f6c634300081e0033",
}

// ReleaseManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ReleaseManagerMetaData.ABI instead.
var ReleaseManagerABI = ReleaseManagerMetaData.ABI

// ReleaseManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ReleaseManagerMetaData.Bin instead.
var ReleaseManagerBin = ReleaseManagerMetaData.Bin

// DeployReleaseManager deploys a new Ethereum contract, binding an instance of ReleaseManager to it.
func DeployReleaseManager(auth *bind.TransactOpts, backend bind.ContractBackend, _permissionController common.Address) (common.Address, *types.Transaction, *ReleaseManager, error) {
	parsed, err := ReleaseManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ReleaseManagerBin), backend, _permissionController)
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
