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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_permissionController\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"},{\"name\":\"_version\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getLatestRelease\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIReleaseManagerTypes.Release\",\"components\":[{\"name\":\"artifacts\",\"type\":\"tuple[]\",\"internalType\":\"structIReleaseManagerTypes.Artifact[]\",\"components\":[{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"registry\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"upgradeByTime\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLatestUpgradeByTime\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMetadataURI\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRelease\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"releaseId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIReleaseManagerTypes.Release\",\"components\":[{\"name\":\"artifacts\",\"type\":\"tuple[]\",\"internalType\":\"structIReleaseManagerTypes.Artifact[]\",\"components\":[{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"registry\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"upgradeByTime\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalReleases\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isValidRelease\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"releaseId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"permissionController\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"publishMetadataURI\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"publishRelease\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"release\",\"type\":\"tuple\",\"internalType\":\"structIReleaseManagerTypes.Release\",\"components\":[{\"name\":\"artifacts\",\"type\":\"tuple[]\",\"internalType\":\"structIReleaseManagerTypes.Artifact[]\",\"components\":[{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"registry\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"upgradeByTime\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"releaseId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MetadataURIPublished\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":true,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"metadataURI\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasePublished\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":true,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"releaseId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"release\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIReleaseManagerTypes.Release\",\"components\":[{\"name\":\"artifacts\",\"type\":\"tuple[]\",\"internalType\":\"structIReleaseManagerTypes.Artifact[]\",\"components\":[{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"registry\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"upgradeByTime\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidMetadataURI\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPermissions\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidUpgradeByTime\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustPublishMetadataURI\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]}]",
	Bin: "0x60c060405234801561000f575f5ffd5b5060405161155638038061155683398101604081905261002e9161016a565b6001600160a01b0382166080528061004581610058565b60a0525061005161009e565b5050610294565b5f5f829050601f8151111561008b578260405163305a27a960e01b81526004016100829190610239565b60405180910390fd5b80516100968261026e565b179392505050565b5f54610100900460ff16156101055760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b6064820152608401610082565b5f5460ff90811614610154575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b634e487b7160e01b5f52604160045260245ffd5b5f5f6040838503121561017b575f5ffd5b82516001600160a01b0381168114610191575f5ffd5b60208401519092506001600160401b038111156101ac575f5ffd5b8301601f810185136101bc575f5ffd5b80516001600160401b038111156101d5576101d5610156565b604051601f8201601f19908116603f011681016001600160401b038111828210171561020357610203610156565b60405281815282820160200187101561021a575f5ffd5b8160208401602083015e5f602083830101528093505050509250929050565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b8051602080830151919081101561028e575f198160200360031b1b821691505b50919050565b60805160a05161129b6102bb5f395f61044e01525f818160cd01526109ed015261129b5ff3fe608060405234801561000f575f5ffd5b506004361061009b575f3560e01c806366f409f71161006357806366f409f7146101545780637c09ea8214610175578063a9e0ed6814610188578063b053b56d146101b0578063d30eeb88146101c3575f5ffd5b80633acab5fc1461009f5780634657e26a146100c85780634840a67c14610107578063517e40681461011c57806354fd4d501461013f575b5f5ffd5b6100b26100ad366004610b5b565b6101e4565b6040516100bf9190610c47565b60405180910390f35b6100ef7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016100bf565b61011a610115366004610c76565b610348565b005b61012f61012a366004610b5b565b610428565b60405190151581526020016100bf565b610147610447565b6040516100bf9190610cf5565b610167610162366004610d07565b610477565b6040519081526020016100bf565b610167610183366004610d21565b61049a565b61019b610196366004610d07565b6106b9565b60405163ffffffff90911681526020016100bf565b6101476101be366004610d07565b61071d565b6101d66101d1366004610d07565b6107c6565b6040516100bf929190610d6d565b60408051808201909152606081525f602082015260015f6102048561094c565b81526020019081526020015f20828154811061022257610222610d8d565b905f5260205f2090600202016040518060400160405290815f8201805480602002602001604051908101604052809291908181526020015f905b82821015610325578382905f5260205f2090600202016040518060400160405290815f820154815260200160018201805461029690610da1565b80601f01602080910402602001604051908101604052809291908181526020018280546102c290610da1565b801561030d5780601f106102e45761010080835404028352916020019161030d565b820191905f5260205f20905b8154815290600101906020018083116102f057829003601f168201915b5050505050815250508152602001906001019061025c565b505050908252506001919091015463ffffffff1660209091015290505b92915050565b6103556020840184610dd3565b61035e816109af565b61037b5760405163932d94f760e01b815260040160405180910390fd5b5f82900361039c57604051630eec403f60e41b815260040160405180910390fd5b828260025f6103b86103b3368a90038a018a610d07565b61094c565b81526020019081526020015f2091826103d2929190610e38565b50836040516103e19190610ef2565b60405180910390207f209e95fbe8dd14c5e1fbf791ee0a83234f45f20cb85504c7068d5ca0d6224588848460405161041a929190610f4d565b60405180910390a250505050565b5f600161043484610477565b61043e9190610f60565b90911492915050565b60606104727f0000000000000000000000000000000000000000000000000000000000000000610a59565b905090565b5f60015f6104848461094c565b815260208101919091526040015f205492915050565b5f6104a86020840184610dd3565b6104b1816109af565b6104ce5760405163932d94f760e01b815260040160405180910390fd5b5f6001816104e46103b336899003890189610d07565b81526020019081526020015f20905060025f61050a878036038101906103b39190610d07565b81526020019081526020015f20805461052290610da1565b90505f036105435760405163413e6e5760e11b815260040160405180910390fd5b426105546040860160208701610f7f565b63ffffffff1610156105795760405163325ec75f60e01b815260040160405180910390fd5b80546001810182555f8281529093505b6105938580610f98565b9050811015610614578184815481106105ae576105ae610d8d565b5f91825260209091206002909102016105c78680610f98565b838181106105d7576105d7610d8d565b90506020028101906105e99190610fe5565b81546001810183555f928352602090922090916002020161060a8282611003565b5050600101610589565b506106256040850160208601610f7f565b81848154811061063757610637610d8d565b905f5260205f2090600202016001015f6101000a81548163ffffffff021916908363ffffffff16021790555082856040516106729190610ef2565b60405180910390207f2decd15222f7c4a8c3d4d2e14dcfdc5a0b52eb2d4b81796bfd010ee5cd972fd3866040516106a991906110fc565b60405180910390a3505092915050565b5f5f60015f6106c78561094c565b81526020019081526020015f2090505f600182805490506106e89190610f60565b90508181815481106106fc576106fc610d8d565b5f91825260209091206001600290920201015463ffffffff16949350505050565b606060025f61072b8461094c565b81526020019081526020015f20805461074390610da1565b80601f016020809104026020016040519081016040528092919081815260200182805461076f90610da1565b80156107ba5780601f10610791576101008083540402835291602001916107ba565b820191905f5260205f20905b81548152906001019060200180831161079d57829003601f168201915b50505050509050919050565b60408051808201909152606081525f60208201819052905f60015f6107ea8661094c565b81526020019081526020015f2090505f6001828054905061080b9190610f60565b90508082828154811061082057610820610d8d565b905f5260205f209060020201806040518060400160405290815f8201805480602002602001604051908101604052809291908181526020015f905b82821015610924578382905f5260205f2090600202016040518060400160405290815f820154815260200160018201805461089590610da1565b80601f01602080910402602001604051908101604052809291908181526020018280546108c190610da1565b801561090c5780601f106108e35761010080835404028352916020019161090c565b820191905f5260205f20905b8154815290600101906020018083116108ef57829003601f168201915b5050505050815250508152602001906001019061085b565b505050908252506001919091015463ffffffff16602090910152919791965090945050505050565b5f815f0151826020015163ffffffff1660405160200161099792919060609290921b6bffffffffffffffffffffffff1916825260a01b6001600160a01b031916601482015260200190565b60405160208183030381529060405261034290611223565b604051631beb2b9760e31b81526001600160a01b0382811660048301523360248301523060448301525f80356001600160e01b0319166064840152917f00000000000000000000000000000000000000000000000000000000000000009091169063df595cb8906084016020604051808303815f875af1158015610a35573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103429190611246565b60605f610a6583610a96565b6040805160208082528183019092529192505f91906020820181803683375050509182525060208101929092525090565b5f60ff8216601f81111561034257604051632cd44ac360e21b815260040160405180910390fd5b634e487b7160e01b5f52604160045260245ffd5b80356001600160a01b0381168114610ae7575f5ffd5b919050565b803563ffffffff81168114610ae7575f5ffd5b5f60408284031215610b0f575f5ffd5b6040516040810181811067ffffffffffffffff82111715610b3257610b32610abd565b604052905080610b4183610ad1565b8152610b4f60208401610aec565b60208201525092915050565b5f5f60608385031215610b6c575f5ffd5b610b768484610aff565b946040939093013593505050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b5f6040830182516040855281815180845260608701915060608160051b88010193506020830192505f5b81811015610c2457605f198886030183528351805186526020810151905060406020870152610c0e6040870182610b84565b9550506020938401939290920191600101610bdc565b505050506020830151610c3f602086018263ffffffff169052565b509392505050565b602081525f610c596020830184610bb2565b9392505050565b5f60408284031215610c70575f5ffd5b50919050565b5f5f5f60608486031215610c88575f5ffd5b610c928585610c60565b9250604084013567ffffffffffffffff811115610cad575f5ffd5b8401601f81018613610cbd575f5ffd5b803567ffffffffffffffff811115610cd3575f5ffd5b866020828401011115610ce4575f5ffd5b939660209190910195509293505050565b602081525f610c596020830184610b84565b5f60408284031215610d17575f5ffd5b610c598383610aff565b5f5f60608385031215610d32575f5ffd5b610d3c8484610c60565b9150604083013567ffffffffffffffff811115610d57575f5ffd5b610d6385828601610c60565b9150509250929050565b828152604060208201525f610d856040830184610bb2565b949350505050565b634e487b7160e01b5f52603260045260245ffd5b600181811c90821680610db557607f821691505b602082108103610c7057634e487b7160e01b5f52602260045260245ffd5b5f60208284031215610de3575f5ffd5b610c5982610ad1565b601f821115610e3357805f5260205f20601f840160051c81016020851015610e115750805b601f840160051c820191505b81811015610e30575f8155600101610e1d565b50505b505050565b67ffffffffffffffff831115610e5057610e50610abd565b610e6483610e5e8354610da1565b83610dec565b5f601f841160018114610e95575f8515610e7e5750838201355b5f19600387901b1c1916600186901b178355610e30565b5f83815260208120601f198716915b82811015610ec45786850135825560209485019460019092019101610ea4565b5086821015610ee0575f1960f88860031b161c19848701351681555b505060018560011b0183555050505050565b6001600160a01b03610f0383610ad1565b16815263ffffffff610f1760208401610aec565b166020820152604001919050565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b602081525f610d85602083018486610f25565b8181038181111561034257634e487b7160e01b5f52601160045260245ffd5b5f60208284031215610f8f575f5ffd5b610c5982610aec565b5f5f8335601e19843603018112610fad575f5ffd5b83018035915067ffffffffffffffff821115610fc7575f5ffd5b6020019150600581901b3603821315610fde575f5ffd5b9250929050565b5f8235603e19833603018112610ff9575f5ffd5b9190910192915050565b81358155600181016020830135601e19843603018112611021575f5ffd5b8301803567ffffffffffffffff8111801561103a575f5ffd5b81360360208401131561104b575f5ffd5b5f9050506110638161105d8554610da1565b85610dec565b5f601f821160018114611097575f83156110805750838201602001355b5f19600385901b1c1916600184901b1785556110f3565b5f85815260208120601f198516915b828110156110c8576020858801810135835594850194600190920191016110a6565b50848210156110e7575f1960f88660031b161c19602085880101351681555b505060018360011b0185555b50505050505050565b602081525f606082018335601e19853603018112611118575f5ffd5b840180356020820167ffffffffffffffff821115611134575f5ffd5b8160051b803603821315611146575f5ffd5b604060208801529382905260809386018401935f908701605e1936869003015b8483101561120057888703607f190182528335818112611184575f5ffd5b860160208101358852604081013536829003603e190181126111a4575f5ffd5b0160408101906020013567ffffffffffffffff8111156111c2575f5ffd5b8036038213156111d0575f5ffd5b604060208a01526111e560408a018284610f25565b98505050602084019350602082019150600183019250611166565b50505050505061121260208501610aec565b63ffffffff81166040850152610c3f565b80516020808301519190811015610c70575f1960209190910360031b1b16919050565b5f60208284031215611256575f5ffd5b81518015158114610c59575f5ffdfea264697066735822122041f0cff4b039206e205278a5bbae0ec194b6bf5c507e596ba497c1bbd73519f664736f6c634300081b0033",
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
