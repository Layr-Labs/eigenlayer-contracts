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
	Digest      [32]byte
	RegistryUrl string
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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_permissionController\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"},{\"name\":\"_version\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getLatestRelease\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIReleaseManagerTypes.Release\",\"components\":[{\"name\":\"artifacts\",\"type\":\"tuple[]\",\"internalType\":\"structIReleaseManagerTypes.Artifact[]\",\"components\":[{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"registryUrl\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"upgradeByTime\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLatestUpgradeByTime\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRelease\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"releaseId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIReleaseManagerTypes.Release\",\"components\":[{\"name\":\"artifacts\",\"type\":\"tuple[]\",\"internalType\":\"structIReleaseManagerTypes.Artifact[]\",\"components\":[{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"registryUrl\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"upgradeByTime\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalReleases\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isValidRelease\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"releaseId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"permissionController\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"publishRelease\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"release\",\"type\":\"tuple\",\"internalType\":\"structIReleaseManagerTypes.Release\",\"components\":[{\"name\":\"artifacts\",\"type\":\"tuple[]\",\"internalType\":\"structIReleaseManagerTypes.Artifact[]\",\"components\":[{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"registryUrl\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"upgradeByTime\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"releaseId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasePublished\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":true,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"releaseId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"release\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIReleaseManagerTypes.Release\",\"components\":[{\"name\":\"artifacts\",\"type\":\"tuple[]\",\"internalType\":\"structIReleaseManagerTypes.Artifact[]\",\"components\":[{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"registryUrl\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"upgradeByTime\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidPermissions\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidUpgradeByTime\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]}]",
	Bin: "0x60c060405234801561000f575f5ffd5b506040516111df3803806111df83398101604081905261002e9161016a565b6001600160a01b0382166080528061004581610058565b60a0525061005161009e565b5050610294565b5f5f829050601f8151111561008b578260405163305a27a960e01b81526004016100829190610239565b60405180910390fd5b80516100968261026e565b179392505050565b5f54610100900460ff16156101055760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b6064820152608401610082565b5f5460ff90811614610154575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b634e487b7160e01b5f52604160045260245ffd5b5f5f6040838503121561017b575f5ffd5b82516001600160a01b0381168114610191575f5ffd5b60208401519092506001600160401b038111156101ac575f5ffd5b8301601f810185136101bc575f5ffd5b80516001600160401b038111156101d5576101d5610156565b604051601f8201601f19908116603f011681016001600160401b038111828210171561020357610203610156565b60405281815282820160200187101561021a575f5ffd5b8160208401602083015e5f602083830101528093505050509250929050565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b8051602080830151919081101561028e575f198160200360031b1b821691505b50919050565b60805160a051610f246102bb5f395f61033001525f818160b7015261081b0152610f245ff3fe608060405234801561000f575f5ffd5b5060043610610085575f3560e01c806366f409f71161005857806366f409f7146101295780637c09ea821461014a578063a9e0ed681461015d578063d30eeb8814610185575f5ffd5b80633acab5fc146100895780634657e26a146100b2578063517e4068146100f157806354fd4d5014610114575b5f5ffd5b61009c610097366004610944565b6101a6565b6040516100a99190610a30565b60405180910390f35b6100d97f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016100a9565b6101046100ff366004610944565b61030a565b60405190151581526020016100a9565b61011c610329565b6040516100a99190610a49565b61013c610137366004610a5b565b610359565b6040519081526020016100a9565b61013c610158366004610a8b565b61037c565b61017061016b366004610a5b565b610553565b60405163ffffffff90911681526020016100a9565b610198610193366004610a5b565b6105b7565b6040516100a9929190610ad7565b60408051808201909152606081525f602082015260015f6101c68561073d565b81526020019081526020015f2082815481106101e4576101e4610af7565b905f5260205f2090600202016040518060400160405290815f8201805480602002602001604051908101604052809291908181526020015f905b828210156102e7578382905f5260205f2090600202016040518060400160405290815f820154815260200160018201805461025890610b0b565b80601f016020809104026020016040519081016040528092919081815260200182805461028490610b0b565b80156102cf5780601f106102a6576101008083540402835291602001916102cf565b820191905f5260205f20905b8154815290600101906020018083116102b257829003601f168201915b5050505050815250508152602001906001019061021e565b505050908252506001919091015463ffffffff1660209091015290505b92915050565b5f600161031684610359565b6103209190610b3d565b90911492915050565b60606103547f00000000000000000000000000000000000000000000000000000000000000006107a0565b905090565b5f60015f6103668461073d565b815260208101919091526040015f205492915050565b5f61038a6020840184610b5c565b610393816107dd565b6103b05760405163932d94f760e01b815260040160405180910390fd5b5f6001816103cb6103c636899003890189610a5b565b61073d565b81526020019081526020015f209050428460200160208101906103ee9190610b75565b63ffffffff1610156104135760405163325ec75f60e01b815260040160405180910390fd5b80546001810182555f8281529093505b61042d8580610b8e565b90508110156104ae5781848154811061044857610448610af7565b5f91825260209091206002909102016104618680610b8e565b8381811061047157610471610af7565b90506020028101906104839190610bdb565b81546001810183555f92835260209092209091600202016104a48282610c45565b5050600101610423565b506104bf6040850160208601610b75565b8184815481106104d1576104d1610af7565b905f5260205f2090600202016001015f6101000a81548163ffffffff021916908363ffffffff160217905550828560405161050c9190610d3e565b60405180910390207f2decd15222f7c4a8c3d4d2e14dcfdc5a0b52eb2d4b81796bfd010ee5cd972fd3866040516105439190610d71565b60405180910390a3505092915050565b5f5f60015f6105618561073d565b81526020019081526020015f2090505f600182805490506105829190610b3d565b905081818154811061059657610596610af7565b5f91825260209091206001600290920201015463ffffffff16949350505050565b60408051808201909152606081525f60208201819052905f60015f6105db8661073d565b81526020019081526020015f2090505f600182805490506105fc9190610b3d565b90508082828154811061061157610611610af7565b905f5260205f209060020201806040518060400160405290815f8201805480602002602001604051908101604052809291908181526020015f905b82821015610715578382905f5260205f2090600202016040518060400160405290815f820154815260200160018201805461068690610b0b565b80601f01602080910402602001604051908101604052809291908181526020018280546106b290610b0b565b80156106fd5780601f106106d4576101008083540402835291602001916106fd565b820191905f5260205f20905b8154815290600101906020018083116106e057829003601f168201915b5050505050815250508152602001906001019061064c565b505050908252506001919091015463ffffffff16602090910152919791965090945050505050565b5f815f0151826020015163ffffffff1660405160200161078892919060609290921b6bffffffffffffffffffffffff1916825260a01b6001600160a01b031916601482015260200190565b60405160208183030381529060405261030490610eac565b60605f6107ac83610887565b6040805160208082528183019092529192505f91906020820181803683375050509182525060208101929092525090565b604051631beb2b9760e31b81526001600160a01b0382811660048301523360248301523060448301525f80356001600160e01b0319166064840152917f00000000000000000000000000000000000000000000000000000000000000009091169063df595cb8906084016020604051808303815f875af1158015610863573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103049190610ecf565b5f60ff8216601f81111561030457604051632cd44ac360e21b815260040160405180910390fd5b80356001600160a01b03811681146108c4575f5ffd5b919050565b803563ffffffff811681146108c4575f5ffd5b5f604082840312156108ec575f5ffd5b6040516040810181811067ffffffffffffffff8211171561091b57634e487b7160e01b5f52604160045260245ffd5b60405290508061092a836108ae565b8152610938602084016108c9565b60208201525092915050565b5f5f60608385031215610955575f5ffd5b61095f84846108dc565b946040939093013593505050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b5f6040830182516040855281815180845260608701915060608160051b88010193506020830192505f5b81811015610a0d57605f1988860301835283518051865260208101519050604060208701526109f7604087018261096d565b95505060209384019392909201916001016109c5565b505050506020830151610a28602086018263ffffffff169052565b509392505050565b602081525f610a42602083018461099b565b9392505050565b602081525f610a42602083018461096d565b5f60408284031215610a6b575f5ffd5b610a4283836108dc565b5f60408284031215610a85575f5ffd5b50919050565b5f5f60608385031215610a9c575f5ffd5b610aa68484610a75565b9150604083013567ffffffffffffffff811115610ac1575f5ffd5b610acd85828601610a75565b9150509250929050565b828152604060208201525f610aef604083018461099b565b949350505050565b634e487b7160e01b5f52603260045260245ffd5b600181811c90821680610b1f57607f821691505b602082108103610a8557634e487b7160e01b5f52602260045260245ffd5b8181038181111561030457634e487b7160e01b5f52601160045260245ffd5b5f60208284031215610b6c575f5ffd5b610a42826108ae565b5f60208284031215610b85575f5ffd5b610a42826108c9565b5f5f8335601e19843603018112610ba3575f5ffd5b83018035915067ffffffffffffffff821115610bbd575f5ffd5b6020019150600581901b3603821315610bd4575f5ffd5b9250929050565b5f8235603e19833603018112610bef575f5ffd5b9190910192915050565b601f821115610c4057805f5260205f20601f840160051c81016020851015610c1e5750805b601f840160051c820191505b81811015610c3d575f8155600101610c2a565b50505b505050565b81358155600181016020830135601e19843603018112610c63575f5ffd5b8301803567ffffffffffffffff81118015610c7c575f5ffd5b813603602084011315610c8d575f5ffd5b5f905050610ca581610c9f8554610b0b565b85610bf9565b5f601f821160018114610cd9575f8315610cc25750838201602001355b5f19600385901b1c1916600184901b178555610d35565b5f85815260208120601f198516915b82811015610d0a57602085880181013583559485019460019092019101610ce8565b5084821015610d29575f1960f88660031b161c19602085880101351681555b505060018360011b0185555b50505050505050565b6001600160a01b03610d4f836108ae565b16815263ffffffff610d63602084016108c9565b166020820152604001919050565b602081525f606082018335601e19853603018112610d8d575f5ffd5b840180356020820167ffffffffffffffff821115610da9575f5ffd5b8160051b803603821315610dbb575f5ffd5b604060208801529382905260809386018401935f908701605e1936869003015b84831015610e8957888703607f190182528335818112610df9575f5ffd5b860160208101358852604081013536829003603e19018112610e19575f5ffd5b0160408101906020013567ffffffffffffffff811115610e37575f5ffd5b803603821315610e45575f5ffd5b604060208a01528060408a0152808260608b01375f6060828b0101526060601f19601f8301168a010198505050602084019350602082019150600183019250610ddb565b505050505050610e9b602085016108c9565b63ffffffff81166040850152610a28565b80516020808301519190811015610a85575f1960209190910360031b1b16919050565b5f60208284031215610edf575f5ffd5b81518015158114610a42575f5ffdfea26469706673582212205b948e2bb7c86c91dca9f27a088a2a0a65e46b42929ba7eace1c73f7c5ee499a64736f6c634300081b0033",
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
