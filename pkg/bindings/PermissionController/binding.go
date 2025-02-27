// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package PermissionController

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

// PermissionControllerMetaData contains all meta data concerning the PermissionController contract.
var PermissionControllerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_version\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptAdmin\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addPendingAdmin\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"canCall\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAdmins\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAppointeePermissions\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"appointee\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAppointees\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPendingAdmins\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isAdmin\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isPendingAdmin\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pendingAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeAdmin\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeAppointee\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"appointee\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removePendingAdmin\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAppointee\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"appointee\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AdminRemoved\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"admin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AdminSet\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"admin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AppointeeRemoved\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"appointee\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"indexed\":false,\"internalType\":\"bytes4\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AppointeeSet\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"appointee\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"indexed\":false,\"internalType\":\"bytes4\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PendingAdminAdded\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"admin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PendingAdminRemoved\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"admin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AdminAlreadyPending\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AdminAlreadySet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AdminNotPending\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AdminNotSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AppointeeAlreadySet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AppointeeNotSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotHaveZeroAdmins\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAdmin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]}]",
	Bin: "0x60a060405234801561000f575f5ffd5b506040516112b73803806112b783398101604081905261002e9161015c565b806100388161004a565b60805250610044610090565b50610267565b5f5f829050601f8151111561007d578260405163305a27a960e01b8152600401610074919061020c565b60405180910390fd5b805161008882610241565b179392505050565b5f54610100900460ff16156100f75760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b6064820152608401610074565b5f5460ff90811614610146575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b634e487b7160e01b5f52604160045260245ffd5b5f6020828403121561016c575f5ffd5b81516001600160401b03811115610181575f5ffd5b8201601f81018413610191575f5ffd5b80516001600160401b038111156101aa576101aa610148565b604051601f8201601f19908116603f011681016001600160401b03811182821017156101d8576101d8610148565b6040528181528282016020018610156101ef575f5ffd5b8160208401602083015e5f91810160200191909152949350505050565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b80516020808301519190811015610261575f198160200360031b1b821691505b50919050565b60805161103861027f5f395f6104d101526110385ff3fe608060405234801561000f575f5ffd5b50600436106100e5575f3560e01c80639100674511610088578063ad8aca7711610063578063ad8aca77146101df578063df595cb8146101f2578063eb5a4e8714610205578063fddbdefd14610218575f5ffd5b80639100674514610196578063950d806e146101b9578063ad5f2210146101cc575f5ffd5b806354fd4d50116100c357806354fd4d5014610124578063628806ef146101425780636bddfa1f14610155578063882a3b3814610175575f5ffd5b806306641201146100e9578063268959e5146100fe5780634f906cf914610111575b5f5ffd5b6100fc6100f7366004610dbd565b61022b565b005b6100fc61010c366004610e0e565b61034c565b6100fc61011f366004610e0e565b610427565b61012c6104ca565b6040516101399190610e3f565b60405180910390f35b6100fc610150366004610e74565b6104fa565b610168610163366004610e74565b610588565b6040516101399190610ed0565b610188610183366004610e0e565b6105b1565b604051610139929190610ee2565b6101a96101a4366004610e0e565b610712565b6040519015158152602001610139565b6100fc6101c7366004610dbd565b610782565b6101686101da366004610e74565b610893565b6101a96101ed366004610e0e565b610939565b6101a9610200366004610dbd565b61095a565b6100fc610213366004610e0e565b6109af565b610168610226366004610f44565b610a7d565b836102368133610712565b61025357604051637bfa4b9f60e01b815260040160405180910390fd5b6001600160a01b0385165f908152600160205260408120906102758585610abb565b6001600160a01b0387165f908152600484016020526040902090915061029b9082610ae8565b6102b85760405163262118cd60e01b815260040160405180910390fd5b6001600160a01b0386165f90815260048301602052604090206102db9082610aff565b505f81815260058301602052604090206102f59087610b0a565b50856001600160a01b0316876001600160a01b03167f18242326b6b862126970679759169f01f646bd55ec5bfcab85ba9f337a74e0c6878760405161033b929190610f84565b60405180910390a350505050505050565b816103578133610712565b61037457604051637bfa4b9f60e01b815260040160405180910390fd5b6001600160a01b0383165f9081526001602081905260409091206002019061039b82610b1e565b116103b9576040516310ce892b60e31b815260040160405180910390fd5b6103c38184610b0a565b6103e057604051630716d81b60e51b815260040160405180910390fd5b6040516001600160a01b0384811682528516907fdb9d5d31320daf5bc7181d565b6da4d12e30f0f4d5aa324a992426c14a1d19ce906020015b60405180910390a250505050565b816104328133610712565b61044f57604051637bfa4b9f60e01b815260040160405180910390fd5b6001600160a01b0383165f9081526001602052604090206104708184610b0a565b61048d5760405163bed8295f60e01b815260040160405180910390fd5b6040516001600160a01b0384811682528516907fd706ed7ae044d795b49e54c9f519f663053951011985f663a862cd9ee72a9ac790602001610419565b60606104f57f0000000000000000000000000000000000000000000000000000000000000000610b27565b905090565b6001600160a01b0381165f90815260016020526040902061051b8133610b0a565b6105385760405163bed8295f60e01b815260040160405180910390fd5b6105456002820133610b64565b506040513381526001600160a01b038316907fbf265e8326285a2747e33e54d5945f7111f2b5edb826eb8c08d4677779b3ff979060200160405180910390a25050565b6001600160a01b0381165f9081526001602052604090206060906105ab90610b78565b92915050565b6001600160a01b038083165f9081526001602090815260408083209385168352600490930190529081206060918291906105ea82610b1e565b90505f8167ffffffffffffffff81111561060657610606610fa7565b60405190808252806020026020018201604052801561062f578160200160208202803683370190505b5090505f8267ffffffffffffffff81111561064c5761064c610fa7565b604051908082528060200260200182016040528015610675578160200160208202803683370190505b5090505f5b83811015610704576106a861068f8683610b84565b606081901c9160a09190911b6001600160e01b03191690565b8483815181106106ba576106ba610fbb565b602002602001018484815181106106d3576106d3610fbb565b6001600160e01b0319909316602093840291909101909201919091526001600160a01b03909116905260010161067a565b509097909650945050505050565b6001600160a01b0382165f90815260016020526040812061073590600201610b1e565b5f0361075757816001600160a01b0316836001600160a01b03161490506105ab565b6001600160a01b0383165f90815260016020526040902061077b9060020183610b8f565b9392505050565b8361078d8133610712565b6107aa57604051637bfa4b9f60e01b815260040160405180910390fd5b6001600160a01b0385165f908152600160205260408120906107cc8585610abb565b6001600160a01b0387165f90815260048401602052604090209091506107f29082610ae8565b156108105760405163ad8efeb760e01b815260040160405180910390fd5b6001600160a01b0386165f90815260048301602052604090206108339082610bb0565b505f818152600583016020526040902061084d9087610b64565b50856001600160a01b0316876001600160a01b03167f037f03a2ad6b967df4a01779b6d2b4c85950df83925d9e31362b519422fc0169878760405161033b929190610f84565b6001600160a01b0381165f9081526001602052604090206060906108b990600201610b1e565b5f03610911576040805160018082528183019092525f916020808301908036833701905050905082815f815181106108f3576108f3610fbb565b6001600160a01b039092166020928302919091019091015292915050565b6001600160a01b0382165f9081526001602052604090206105ab90600201610b78565b919050565b6001600160a01b0382165f90815260016020526040812061077b9083610b8f565b5f6109658585610712565b806109a657506109a66109788484610abb565b6001600160a01b038088165f908152600160209081526040808320938a168352600490930190522090610ae8565b95945050505050565b816109ba8133610712565b6109d757604051637bfa4b9f60e01b815260040160405180910390fd5b6001600160a01b0383165f9081526001602052604090206109fb6002820184610b8f565b15610a195760405163130160e560e31b815260040160405180910390fd5b610a238184610b64565b610a40576040516319abede360e11b815260040160405180910390fd5b6040516001600160a01b0384811682528516907fb14b9a3d448c5b04f0e5b087b6f5193390db7955482a6ffb841e7b3ba61a460c90602001610419565b60605f610a8a8484610abb565b6001600160a01b0386165f90815260016020908152604080832084845260050190915290209091506109a690610b78565b60609190911b6bffffffffffffffffffffffff191660a09190911c6bffffffff0000000000000000161790565b5f818152600183016020526040812054151561077b565b5f61077b8383610bbb565b5f61077b836001600160a01b038416610bbb565b5f6105ab825490565b60605f610b3383610c9e565b6040805160208082528183019092529192505f91906020820181803683375050509182525060208101929092525090565b5f61077b836001600160a01b038416610cc5565b60605f61077b83610d11565b5f61077b8383610d6a565b6001600160a01b0381165f908152600183016020526040812054151561077b565b5f61077b8383610cc5565b5f8181526001830160205260408120548015610c95575f610bdd600183610fcf565b85549091505f90610bf090600190610fcf565b9050818114610c4f575f865f018281548110610c0e57610c0e610fbb565b905f5260205f200154905080875f018481548110610c2e57610c2e610fbb565b5f918252602080832090910192909255918252600188019052604090208390555b8554869080610c6057610c60610fee565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f9055600193505050506105ab565b5f9150506105ab565b5f60ff8216601f8111156105ab57604051632cd44ac360e21b815260040160405180910390fd5b5f818152600183016020526040812054610d0a57508154600181810184555f8481526020808220909301849055845484825282860190935260409020919091556105ab565b505f6105ab565b6060815f01805480602002602001604051908101604052809291908181526020018280548015610d5e57602002820191905f5260205f20905b815481526020019060010190808311610d4a575b50505050509050919050565b5f825f018281548110610d7f57610d7f610fbb565b905f5260205f200154905092915050565b80356001600160a01b0381168114610934575f5ffd5b80356001600160e01b031981168114610934575f5ffd5b5f5f5f5f60808587031215610dd0575f5ffd5b610dd985610d90565b9350610de760208601610d90565b9250610df560408601610d90565b9150610e0360608601610da6565b905092959194509250565b5f5f60408385031215610e1f575f5ffd5b610e2883610d90565b9150610e3660208401610d90565b90509250929050565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b5f60208284031215610e84575f5ffd5b61077b82610d90565b5f8151808452602084019350602083015f5b82811015610ec65781516001600160a01b0316865260209586019590910190600101610e9f565b5093949350505050565b602081525f61077b6020830184610e8d565b604081525f610ef46040830185610e8d565b82810360208401528084518083526020830191506020860192505f5b81811015610f385783516001600160e01b031916835260209384019390920191600101610f10565b50909695505050505050565b5f5f5f60608486031215610f56575f5ffd5b610f5f84610d90565b9250610f6d60208501610d90565b9150610f7b60408501610da6565b90509250925092565b6001600160a01b039290921682526001600160e01b031916602082015260400190565b634e487b7160e01b5f52604160045260245ffd5b634e487b7160e01b5f52603260045260245ffd5b818103818111156105ab57634e487b7160e01b5f52601160045260245ffd5b634e487b7160e01b5f52603160045260245ffdfea2646970667358221220ab9bde44aeb13051e26dbeabefa3bed0231ce62858e181557792c52c88a0538164736f6c634300081b0033",
}

// PermissionControllerABI is the input ABI used to generate the binding from.
// Deprecated: Use PermissionControllerMetaData.ABI instead.
var PermissionControllerABI = PermissionControllerMetaData.ABI

// PermissionControllerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PermissionControllerMetaData.Bin instead.
var PermissionControllerBin = PermissionControllerMetaData.Bin

// DeployPermissionController deploys a new Ethereum contract, binding an instance of PermissionController to it.
func DeployPermissionController(auth *bind.TransactOpts, backend bind.ContractBackend, _version string) (common.Address, *types.Transaction, *PermissionController, error) {
	parsed, err := PermissionControllerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PermissionControllerBin), backend, _version)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PermissionController{PermissionControllerCaller: PermissionControllerCaller{contract: contract}, PermissionControllerTransactor: PermissionControllerTransactor{contract: contract}, PermissionControllerFilterer: PermissionControllerFilterer{contract: contract}}, nil
}

// PermissionController is an auto generated Go binding around an Ethereum contract.
type PermissionController struct {
	PermissionControllerCaller     // Read-only binding to the contract
	PermissionControllerTransactor // Write-only binding to the contract
	PermissionControllerFilterer   // Log filterer for contract events
}

// PermissionControllerCaller is an auto generated read-only Go binding around an Ethereum contract.
type PermissionControllerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermissionControllerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PermissionControllerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermissionControllerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PermissionControllerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermissionControllerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PermissionControllerSession struct {
	Contract     *PermissionController // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PermissionControllerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PermissionControllerCallerSession struct {
	Contract *PermissionControllerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// PermissionControllerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PermissionControllerTransactorSession struct {
	Contract     *PermissionControllerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// PermissionControllerRaw is an auto generated low-level Go binding around an Ethereum contract.
type PermissionControllerRaw struct {
	Contract *PermissionController // Generic contract binding to access the raw methods on
}

// PermissionControllerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PermissionControllerCallerRaw struct {
	Contract *PermissionControllerCaller // Generic read-only contract binding to access the raw methods on
}

// PermissionControllerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PermissionControllerTransactorRaw struct {
	Contract *PermissionControllerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPermissionController creates a new instance of PermissionController, bound to a specific deployed contract.
func NewPermissionController(address common.Address, backend bind.ContractBackend) (*PermissionController, error) {
	contract, err := bindPermissionController(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PermissionController{PermissionControllerCaller: PermissionControllerCaller{contract: contract}, PermissionControllerTransactor: PermissionControllerTransactor{contract: contract}, PermissionControllerFilterer: PermissionControllerFilterer{contract: contract}}, nil
}

// NewPermissionControllerCaller creates a new read-only instance of PermissionController, bound to a specific deployed contract.
func NewPermissionControllerCaller(address common.Address, caller bind.ContractCaller) (*PermissionControllerCaller, error) {
	contract, err := bindPermissionController(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PermissionControllerCaller{contract: contract}, nil
}

// NewPermissionControllerTransactor creates a new write-only instance of PermissionController, bound to a specific deployed contract.
func NewPermissionControllerTransactor(address common.Address, transactor bind.ContractTransactor) (*PermissionControllerTransactor, error) {
	contract, err := bindPermissionController(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PermissionControllerTransactor{contract: contract}, nil
}

// NewPermissionControllerFilterer creates a new log filterer instance of PermissionController, bound to a specific deployed contract.
func NewPermissionControllerFilterer(address common.Address, filterer bind.ContractFilterer) (*PermissionControllerFilterer, error) {
	contract, err := bindPermissionController(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PermissionControllerFilterer{contract: contract}, nil
}

// bindPermissionController binds a generic wrapper to an already deployed contract.
func bindPermissionController(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PermissionControllerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PermissionController *PermissionControllerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PermissionController.Contract.PermissionControllerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PermissionController *PermissionControllerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PermissionController.Contract.PermissionControllerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PermissionController *PermissionControllerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PermissionController.Contract.PermissionControllerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PermissionController *PermissionControllerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PermissionController.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PermissionController *PermissionControllerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PermissionController.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PermissionController *PermissionControllerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PermissionController.Contract.contract.Transact(opts, method, params...)
}

// CanCall is a free data retrieval call binding the contract method 0xdf595cb8.
//
// Solidity: function canCall(address account, address caller, address target, bytes4 selector) view returns(bool)
func (_PermissionController *PermissionControllerCaller) CanCall(opts *bind.CallOpts, account common.Address, caller common.Address, target common.Address, selector [4]byte) (bool, error) {
	var out []interface{}
	err := _PermissionController.contract.Call(opts, &out, "canCall", account, caller, target, selector)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanCall is a free data retrieval call binding the contract method 0xdf595cb8.
//
// Solidity: function canCall(address account, address caller, address target, bytes4 selector) view returns(bool)
func (_PermissionController *PermissionControllerSession) CanCall(account common.Address, caller common.Address, target common.Address, selector [4]byte) (bool, error) {
	return _PermissionController.Contract.CanCall(&_PermissionController.CallOpts, account, caller, target, selector)
}

// CanCall is a free data retrieval call binding the contract method 0xdf595cb8.
//
// Solidity: function canCall(address account, address caller, address target, bytes4 selector) view returns(bool)
func (_PermissionController *PermissionControllerCallerSession) CanCall(account common.Address, caller common.Address, target common.Address, selector [4]byte) (bool, error) {
	return _PermissionController.Contract.CanCall(&_PermissionController.CallOpts, account, caller, target, selector)
}

// GetAdmins is a free data retrieval call binding the contract method 0xad5f2210.
//
// Solidity: function getAdmins(address account) view returns(address[])
func (_PermissionController *PermissionControllerCaller) GetAdmins(opts *bind.CallOpts, account common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _PermissionController.contract.Call(opts, &out, "getAdmins", account)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAdmins is a free data retrieval call binding the contract method 0xad5f2210.
//
// Solidity: function getAdmins(address account) view returns(address[])
func (_PermissionController *PermissionControllerSession) GetAdmins(account common.Address) ([]common.Address, error) {
	return _PermissionController.Contract.GetAdmins(&_PermissionController.CallOpts, account)
}

// GetAdmins is a free data retrieval call binding the contract method 0xad5f2210.
//
// Solidity: function getAdmins(address account) view returns(address[])
func (_PermissionController *PermissionControllerCallerSession) GetAdmins(account common.Address) ([]common.Address, error) {
	return _PermissionController.Contract.GetAdmins(&_PermissionController.CallOpts, account)
}

// GetAppointeePermissions is a free data retrieval call binding the contract method 0x882a3b38.
//
// Solidity: function getAppointeePermissions(address account, address appointee) view returns(address[], bytes4[])
func (_PermissionController *PermissionControllerCaller) GetAppointeePermissions(opts *bind.CallOpts, account common.Address, appointee common.Address) ([]common.Address, [][4]byte, error) {
	var out []interface{}
	err := _PermissionController.contract.Call(opts, &out, "getAppointeePermissions", account, appointee)

	if err != nil {
		return *new([]common.Address), *new([][4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([][4]byte)).(*[][4]byte)

	return out0, out1, err

}

// GetAppointeePermissions is a free data retrieval call binding the contract method 0x882a3b38.
//
// Solidity: function getAppointeePermissions(address account, address appointee) view returns(address[], bytes4[])
func (_PermissionController *PermissionControllerSession) GetAppointeePermissions(account common.Address, appointee common.Address) ([]common.Address, [][4]byte, error) {
	return _PermissionController.Contract.GetAppointeePermissions(&_PermissionController.CallOpts, account, appointee)
}

// GetAppointeePermissions is a free data retrieval call binding the contract method 0x882a3b38.
//
// Solidity: function getAppointeePermissions(address account, address appointee) view returns(address[], bytes4[])
func (_PermissionController *PermissionControllerCallerSession) GetAppointeePermissions(account common.Address, appointee common.Address) ([]common.Address, [][4]byte, error) {
	return _PermissionController.Contract.GetAppointeePermissions(&_PermissionController.CallOpts, account, appointee)
}

// GetAppointees is a free data retrieval call binding the contract method 0xfddbdefd.
//
// Solidity: function getAppointees(address account, address target, bytes4 selector) view returns(address[])
func (_PermissionController *PermissionControllerCaller) GetAppointees(opts *bind.CallOpts, account common.Address, target common.Address, selector [4]byte) ([]common.Address, error) {
	var out []interface{}
	err := _PermissionController.contract.Call(opts, &out, "getAppointees", account, target, selector)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAppointees is a free data retrieval call binding the contract method 0xfddbdefd.
//
// Solidity: function getAppointees(address account, address target, bytes4 selector) view returns(address[])
func (_PermissionController *PermissionControllerSession) GetAppointees(account common.Address, target common.Address, selector [4]byte) ([]common.Address, error) {
	return _PermissionController.Contract.GetAppointees(&_PermissionController.CallOpts, account, target, selector)
}

// GetAppointees is a free data retrieval call binding the contract method 0xfddbdefd.
//
// Solidity: function getAppointees(address account, address target, bytes4 selector) view returns(address[])
func (_PermissionController *PermissionControllerCallerSession) GetAppointees(account common.Address, target common.Address, selector [4]byte) ([]common.Address, error) {
	return _PermissionController.Contract.GetAppointees(&_PermissionController.CallOpts, account, target, selector)
}

// GetPendingAdmins is a free data retrieval call binding the contract method 0x6bddfa1f.
//
// Solidity: function getPendingAdmins(address account) view returns(address[])
func (_PermissionController *PermissionControllerCaller) GetPendingAdmins(opts *bind.CallOpts, account common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _PermissionController.contract.Call(opts, &out, "getPendingAdmins", account)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetPendingAdmins is a free data retrieval call binding the contract method 0x6bddfa1f.
//
// Solidity: function getPendingAdmins(address account) view returns(address[])
func (_PermissionController *PermissionControllerSession) GetPendingAdmins(account common.Address) ([]common.Address, error) {
	return _PermissionController.Contract.GetPendingAdmins(&_PermissionController.CallOpts, account)
}

// GetPendingAdmins is a free data retrieval call binding the contract method 0x6bddfa1f.
//
// Solidity: function getPendingAdmins(address account) view returns(address[])
func (_PermissionController *PermissionControllerCallerSession) GetPendingAdmins(account common.Address) ([]common.Address, error) {
	return _PermissionController.Contract.GetPendingAdmins(&_PermissionController.CallOpts, account)
}

// IsAdmin is a free data retrieval call binding the contract method 0x91006745.
//
// Solidity: function isAdmin(address account, address caller) view returns(bool)
func (_PermissionController *PermissionControllerCaller) IsAdmin(opts *bind.CallOpts, account common.Address, caller common.Address) (bool, error) {
	var out []interface{}
	err := _PermissionController.contract.Call(opts, &out, "isAdmin", account, caller)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x91006745.
//
// Solidity: function isAdmin(address account, address caller) view returns(bool)
func (_PermissionController *PermissionControllerSession) IsAdmin(account common.Address, caller common.Address) (bool, error) {
	return _PermissionController.Contract.IsAdmin(&_PermissionController.CallOpts, account, caller)
}

// IsAdmin is a free data retrieval call binding the contract method 0x91006745.
//
// Solidity: function isAdmin(address account, address caller) view returns(bool)
func (_PermissionController *PermissionControllerCallerSession) IsAdmin(account common.Address, caller common.Address) (bool, error) {
	return _PermissionController.Contract.IsAdmin(&_PermissionController.CallOpts, account, caller)
}

// IsPendingAdmin is a free data retrieval call binding the contract method 0xad8aca77.
//
// Solidity: function isPendingAdmin(address account, address pendingAdmin) view returns(bool)
func (_PermissionController *PermissionControllerCaller) IsPendingAdmin(opts *bind.CallOpts, account common.Address, pendingAdmin common.Address) (bool, error) {
	var out []interface{}
	err := _PermissionController.contract.Call(opts, &out, "isPendingAdmin", account, pendingAdmin)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPendingAdmin is a free data retrieval call binding the contract method 0xad8aca77.
//
// Solidity: function isPendingAdmin(address account, address pendingAdmin) view returns(bool)
func (_PermissionController *PermissionControllerSession) IsPendingAdmin(account common.Address, pendingAdmin common.Address) (bool, error) {
	return _PermissionController.Contract.IsPendingAdmin(&_PermissionController.CallOpts, account, pendingAdmin)
}

// IsPendingAdmin is a free data retrieval call binding the contract method 0xad8aca77.
//
// Solidity: function isPendingAdmin(address account, address pendingAdmin) view returns(bool)
func (_PermissionController *PermissionControllerCallerSession) IsPendingAdmin(account common.Address, pendingAdmin common.Address) (bool, error) {
	return _PermissionController.Contract.IsPendingAdmin(&_PermissionController.CallOpts, account, pendingAdmin)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_PermissionController *PermissionControllerCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PermissionController.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_PermissionController *PermissionControllerSession) Version() (string, error) {
	return _PermissionController.Contract.Version(&_PermissionController.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_PermissionController *PermissionControllerCallerSession) Version() (string, error) {
	return _PermissionController.Contract.Version(&_PermissionController.CallOpts)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0x628806ef.
//
// Solidity: function acceptAdmin(address account) returns()
func (_PermissionController *PermissionControllerTransactor) AcceptAdmin(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _PermissionController.contract.Transact(opts, "acceptAdmin", account)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0x628806ef.
//
// Solidity: function acceptAdmin(address account) returns()
func (_PermissionController *PermissionControllerSession) AcceptAdmin(account common.Address) (*types.Transaction, error) {
	return _PermissionController.Contract.AcceptAdmin(&_PermissionController.TransactOpts, account)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0x628806ef.
//
// Solidity: function acceptAdmin(address account) returns()
func (_PermissionController *PermissionControllerTransactorSession) AcceptAdmin(account common.Address) (*types.Transaction, error) {
	return _PermissionController.Contract.AcceptAdmin(&_PermissionController.TransactOpts, account)
}

// AddPendingAdmin is a paid mutator transaction binding the contract method 0xeb5a4e87.
//
// Solidity: function addPendingAdmin(address account, address admin) returns()
func (_PermissionController *PermissionControllerTransactor) AddPendingAdmin(opts *bind.TransactOpts, account common.Address, admin common.Address) (*types.Transaction, error) {
	return _PermissionController.contract.Transact(opts, "addPendingAdmin", account, admin)
}

// AddPendingAdmin is a paid mutator transaction binding the contract method 0xeb5a4e87.
//
// Solidity: function addPendingAdmin(address account, address admin) returns()
func (_PermissionController *PermissionControllerSession) AddPendingAdmin(account common.Address, admin common.Address) (*types.Transaction, error) {
	return _PermissionController.Contract.AddPendingAdmin(&_PermissionController.TransactOpts, account, admin)
}

// AddPendingAdmin is a paid mutator transaction binding the contract method 0xeb5a4e87.
//
// Solidity: function addPendingAdmin(address account, address admin) returns()
func (_PermissionController *PermissionControllerTransactorSession) AddPendingAdmin(account common.Address, admin common.Address) (*types.Transaction, error) {
	return _PermissionController.Contract.AddPendingAdmin(&_PermissionController.TransactOpts, account, admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x268959e5.
//
// Solidity: function removeAdmin(address account, address admin) returns()
func (_PermissionController *PermissionControllerTransactor) RemoveAdmin(opts *bind.TransactOpts, account common.Address, admin common.Address) (*types.Transaction, error) {
	return _PermissionController.contract.Transact(opts, "removeAdmin", account, admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x268959e5.
//
// Solidity: function removeAdmin(address account, address admin) returns()
func (_PermissionController *PermissionControllerSession) RemoveAdmin(account common.Address, admin common.Address) (*types.Transaction, error) {
	return _PermissionController.Contract.RemoveAdmin(&_PermissionController.TransactOpts, account, admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x268959e5.
//
// Solidity: function removeAdmin(address account, address admin) returns()
func (_PermissionController *PermissionControllerTransactorSession) RemoveAdmin(account common.Address, admin common.Address) (*types.Transaction, error) {
	return _PermissionController.Contract.RemoveAdmin(&_PermissionController.TransactOpts, account, admin)
}

// RemoveAppointee is a paid mutator transaction binding the contract method 0x06641201.
//
// Solidity: function removeAppointee(address account, address appointee, address target, bytes4 selector) returns()
func (_PermissionController *PermissionControllerTransactor) RemoveAppointee(opts *bind.TransactOpts, account common.Address, appointee common.Address, target common.Address, selector [4]byte) (*types.Transaction, error) {
	return _PermissionController.contract.Transact(opts, "removeAppointee", account, appointee, target, selector)
}

// RemoveAppointee is a paid mutator transaction binding the contract method 0x06641201.
//
// Solidity: function removeAppointee(address account, address appointee, address target, bytes4 selector) returns()
func (_PermissionController *PermissionControllerSession) RemoveAppointee(account common.Address, appointee common.Address, target common.Address, selector [4]byte) (*types.Transaction, error) {
	return _PermissionController.Contract.RemoveAppointee(&_PermissionController.TransactOpts, account, appointee, target, selector)
}

// RemoveAppointee is a paid mutator transaction binding the contract method 0x06641201.
//
// Solidity: function removeAppointee(address account, address appointee, address target, bytes4 selector) returns()
func (_PermissionController *PermissionControllerTransactorSession) RemoveAppointee(account common.Address, appointee common.Address, target common.Address, selector [4]byte) (*types.Transaction, error) {
	return _PermissionController.Contract.RemoveAppointee(&_PermissionController.TransactOpts, account, appointee, target, selector)
}

// RemovePendingAdmin is a paid mutator transaction binding the contract method 0x4f906cf9.
//
// Solidity: function removePendingAdmin(address account, address admin) returns()
func (_PermissionController *PermissionControllerTransactor) RemovePendingAdmin(opts *bind.TransactOpts, account common.Address, admin common.Address) (*types.Transaction, error) {
	return _PermissionController.contract.Transact(opts, "removePendingAdmin", account, admin)
}

// RemovePendingAdmin is a paid mutator transaction binding the contract method 0x4f906cf9.
//
// Solidity: function removePendingAdmin(address account, address admin) returns()
func (_PermissionController *PermissionControllerSession) RemovePendingAdmin(account common.Address, admin common.Address) (*types.Transaction, error) {
	return _PermissionController.Contract.RemovePendingAdmin(&_PermissionController.TransactOpts, account, admin)
}

// RemovePendingAdmin is a paid mutator transaction binding the contract method 0x4f906cf9.
//
// Solidity: function removePendingAdmin(address account, address admin) returns()
func (_PermissionController *PermissionControllerTransactorSession) RemovePendingAdmin(account common.Address, admin common.Address) (*types.Transaction, error) {
	return _PermissionController.Contract.RemovePendingAdmin(&_PermissionController.TransactOpts, account, admin)
}

// SetAppointee is a paid mutator transaction binding the contract method 0x950d806e.
//
// Solidity: function setAppointee(address account, address appointee, address target, bytes4 selector) returns()
func (_PermissionController *PermissionControllerTransactor) SetAppointee(opts *bind.TransactOpts, account common.Address, appointee common.Address, target common.Address, selector [4]byte) (*types.Transaction, error) {
	return _PermissionController.contract.Transact(opts, "setAppointee", account, appointee, target, selector)
}

// SetAppointee is a paid mutator transaction binding the contract method 0x950d806e.
//
// Solidity: function setAppointee(address account, address appointee, address target, bytes4 selector) returns()
func (_PermissionController *PermissionControllerSession) SetAppointee(account common.Address, appointee common.Address, target common.Address, selector [4]byte) (*types.Transaction, error) {
	return _PermissionController.Contract.SetAppointee(&_PermissionController.TransactOpts, account, appointee, target, selector)
}

// SetAppointee is a paid mutator transaction binding the contract method 0x950d806e.
//
// Solidity: function setAppointee(address account, address appointee, address target, bytes4 selector) returns()
func (_PermissionController *PermissionControllerTransactorSession) SetAppointee(account common.Address, appointee common.Address, target common.Address, selector [4]byte) (*types.Transaction, error) {
	return _PermissionController.Contract.SetAppointee(&_PermissionController.TransactOpts, account, appointee, target, selector)
}

// PermissionControllerAdminRemovedIterator is returned from FilterAdminRemoved and is used to iterate over the raw logs and unpacked data for AdminRemoved events raised by the PermissionController contract.
type PermissionControllerAdminRemovedIterator struct {
	Event *PermissionControllerAdminRemoved // Event containing the contract specifics and raw log

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
func (it *PermissionControllerAdminRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionControllerAdminRemoved)
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
		it.Event = new(PermissionControllerAdminRemoved)
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
func (it *PermissionControllerAdminRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionControllerAdminRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionControllerAdminRemoved represents a AdminRemoved event raised by the PermissionController contract.
type PermissionControllerAdminRemoved struct {
	Account common.Address
	Admin   common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAdminRemoved is a free log retrieval operation binding the contract event 0xdb9d5d31320daf5bc7181d565b6da4d12e30f0f4d5aa324a992426c14a1d19ce.
//
// Solidity: event AdminRemoved(address indexed account, address admin)
func (_PermissionController *PermissionControllerFilterer) FilterAdminRemoved(opts *bind.FilterOpts, account []common.Address) (*PermissionControllerAdminRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PermissionController.contract.FilterLogs(opts, "AdminRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &PermissionControllerAdminRemovedIterator{contract: _PermissionController.contract, event: "AdminRemoved", logs: logs, sub: sub}, nil
}

// WatchAdminRemoved is a free log subscription operation binding the contract event 0xdb9d5d31320daf5bc7181d565b6da4d12e30f0f4d5aa324a992426c14a1d19ce.
//
// Solidity: event AdminRemoved(address indexed account, address admin)
func (_PermissionController *PermissionControllerFilterer) WatchAdminRemoved(opts *bind.WatchOpts, sink chan<- *PermissionControllerAdminRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PermissionController.contract.WatchLogs(opts, "AdminRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionControllerAdminRemoved)
				if err := _PermissionController.contract.UnpackLog(event, "AdminRemoved", log); err != nil {
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

// ParseAdminRemoved is a log parse operation binding the contract event 0xdb9d5d31320daf5bc7181d565b6da4d12e30f0f4d5aa324a992426c14a1d19ce.
//
// Solidity: event AdminRemoved(address indexed account, address admin)
func (_PermissionController *PermissionControllerFilterer) ParseAdminRemoved(log types.Log) (*PermissionControllerAdminRemoved, error) {
	event := new(PermissionControllerAdminRemoved)
	if err := _PermissionController.contract.UnpackLog(event, "AdminRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionControllerAdminSetIterator is returned from FilterAdminSet and is used to iterate over the raw logs and unpacked data for AdminSet events raised by the PermissionController contract.
type PermissionControllerAdminSetIterator struct {
	Event *PermissionControllerAdminSet // Event containing the contract specifics and raw log

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
func (it *PermissionControllerAdminSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionControllerAdminSet)
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
		it.Event = new(PermissionControllerAdminSet)
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
func (it *PermissionControllerAdminSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionControllerAdminSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionControllerAdminSet represents a AdminSet event raised by the PermissionController contract.
type PermissionControllerAdminSet struct {
	Account common.Address
	Admin   common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAdminSet is a free log retrieval operation binding the contract event 0xbf265e8326285a2747e33e54d5945f7111f2b5edb826eb8c08d4677779b3ff97.
//
// Solidity: event AdminSet(address indexed account, address admin)
func (_PermissionController *PermissionControllerFilterer) FilterAdminSet(opts *bind.FilterOpts, account []common.Address) (*PermissionControllerAdminSetIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PermissionController.contract.FilterLogs(opts, "AdminSet", accountRule)
	if err != nil {
		return nil, err
	}
	return &PermissionControllerAdminSetIterator{contract: _PermissionController.contract, event: "AdminSet", logs: logs, sub: sub}, nil
}

// WatchAdminSet is a free log subscription operation binding the contract event 0xbf265e8326285a2747e33e54d5945f7111f2b5edb826eb8c08d4677779b3ff97.
//
// Solidity: event AdminSet(address indexed account, address admin)
func (_PermissionController *PermissionControllerFilterer) WatchAdminSet(opts *bind.WatchOpts, sink chan<- *PermissionControllerAdminSet, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PermissionController.contract.WatchLogs(opts, "AdminSet", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionControllerAdminSet)
				if err := _PermissionController.contract.UnpackLog(event, "AdminSet", log); err != nil {
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

// ParseAdminSet is a log parse operation binding the contract event 0xbf265e8326285a2747e33e54d5945f7111f2b5edb826eb8c08d4677779b3ff97.
//
// Solidity: event AdminSet(address indexed account, address admin)
func (_PermissionController *PermissionControllerFilterer) ParseAdminSet(log types.Log) (*PermissionControllerAdminSet, error) {
	event := new(PermissionControllerAdminSet)
	if err := _PermissionController.contract.UnpackLog(event, "AdminSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionControllerAppointeeRemovedIterator is returned from FilterAppointeeRemoved and is used to iterate over the raw logs and unpacked data for AppointeeRemoved events raised by the PermissionController contract.
type PermissionControllerAppointeeRemovedIterator struct {
	Event *PermissionControllerAppointeeRemoved // Event containing the contract specifics and raw log

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
func (it *PermissionControllerAppointeeRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionControllerAppointeeRemoved)
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
		it.Event = new(PermissionControllerAppointeeRemoved)
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
func (it *PermissionControllerAppointeeRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionControllerAppointeeRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionControllerAppointeeRemoved represents a AppointeeRemoved event raised by the PermissionController contract.
type PermissionControllerAppointeeRemoved struct {
	Account   common.Address
	Appointee common.Address
	Target    common.Address
	Selector  [4]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAppointeeRemoved is a free log retrieval operation binding the contract event 0x18242326b6b862126970679759169f01f646bd55ec5bfcab85ba9f337a74e0c6.
//
// Solidity: event AppointeeRemoved(address indexed account, address indexed appointee, address target, bytes4 selector)
func (_PermissionController *PermissionControllerFilterer) FilterAppointeeRemoved(opts *bind.FilterOpts, account []common.Address, appointee []common.Address) (*PermissionControllerAppointeeRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var appointeeRule []interface{}
	for _, appointeeItem := range appointee {
		appointeeRule = append(appointeeRule, appointeeItem)
	}

	logs, sub, err := _PermissionController.contract.FilterLogs(opts, "AppointeeRemoved", accountRule, appointeeRule)
	if err != nil {
		return nil, err
	}
	return &PermissionControllerAppointeeRemovedIterator{contract: _PermissionController.contract, event: "AppointeeRemoved", logs: logs, sub: sub}, nil
}

// WatchAppointeeRemoved is a free log subscription operation binding the contract event 0x18242326b6b862126970679759169f01f646bd55ec5bfcab85ba9f337a74e0c6.
//
// Solidity: event AppointeeRemoved(address indexed account, address indexed appointee, address target, bytes4 selector)
func (_PermissionController *PermissionControllerFilterer) WatchAppointeeRemoved(opts *bind.WatchOpts, sink chan<- *PermissionControllerAppointeeRemoved, account []common.Address, appointee []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var appointeeRule []interface{}
	for _, appointeeItem := range appointee {
		appointeeRule = append(appointeeRule, appointeeItem)
	}

	logs, sub, err := _PermissionController.contract.WatchLogs(opts, "AppointeeRemoved", accountRule, appointeeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionControllerAppointeeRemoved)
				if err := _PermissionController.contract.UnpackLog(event, "AppointeeRemoved", log); err != nil {
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

// ParseAppointeeRemoved is a log parse operation binding the contract event 0x18242326b6b862126970679759169f01f646bd55ec5bfcab85ba9f337a74e0c6.
//
// Solidity: event AppointeeRemoved(address indexed account, address indexed appointee, address target, bytes4 selector)
func (_PermissionController *PermissionControllerFilterer) ParseAppointeeRemoved(log types.Log) (*PermissionControllerAppointeeRemoved, error) {
	event := new(PermissionControllerAppointeeRemoved)
	if err := _PermissionController.contract.UnpackLog(event, "AppointeeRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionControllerAppointeeSetIterator is returned from FilterAppointeeSet and is used to iterate over the raw logs and unpacked data for AppointeeSet events raised by the PermissionController contract.
type PermissionControllerAppointeeSetIterator struct {
	Event *PermissionControllerAppointeeSet // Event containing the contract specifics and raw log

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
func (it *PermissionControllerAppointeeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionControllerAppointeeSet)
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
		it.Event = new(PermissionControllerAppointeeSet)
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
func (it *PermissionControllerAppointeeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionControllerAppointeeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionControllerAppointeeSet represents a AppointeeSet event raised by the PermissionController contract.
type PermissionControllerAppointeeSet struct {
	Account   common.Address
	Appointee common.Address
	Target    common.Address
	Selector  [4]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAppointeeSet is a free log retrieval operation binding the contract event 0x037f03a2ad6b967df4a01779b6d2b4c85950df83925d9e31362b519422fc0169.
//
// Solidity: event AppointeeSet(address indexed account, address indexed appointee, address target, bytes4 selector)
func (_PermissionController *PermissionControllerFilterer) FilterAppointeeSet(opts *bind.FilterOpts, account []common.Address, appointee []common.Address) (*PermissionControllerAppointeeSetIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var appointeeRule []interface{}
	for _, appointeeItem := range appointee {
		appointeeRule = append(appointeeRule, appointeeItem)
	}

	logs, sub, err := _PermissionController.contract.FilterLogs(opts, "AppointeeSet", accountRule, appointeeRule)
	if err != nil {
		return nil, err
	}
	return &PermissionControllerAppointeeSetIterator{contract: _PermissionController.contract, event: "AppointeeSet", logs: logs, sub: sub}, nil
}

// WatchAppointeeSet is a free log subscription operation binding the contract event 0x037f03a2ad6b967df4a01779b6d2b4c85950df83925d9e31362b519422fc0169.
//
// Solidity: event AppointeeSet(address indexed account, address indexed appointee, address target, bytes4 selector)
func (_PermissionController *PermissionControllerFilterer) WatchAppointeeSet(opts *bind.WatchOpts, sink chan<- *PermissionControllerAppointeeSet, account []common.Address, appointee []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var appointeeRule []interface{}
	for _, appointeeItem := range appointee {
		appointeeRule = append(appointeeRule, appointeeItem)
	}

	logs, sub, err := _PermissionController.contract.WatchLogs(opts, "AppointeeSet", accountRule, appointeeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionControllerAppointeeSet)
				if err := _PermissionController.contract.UnpackLog(event, "AppointeeSet", log); err != nil {
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

// ParseAppointeeSet is a log parse operation binding the contract event 0x037f03a2ad6b967df4a01779b6d2b4c85950df83925d9e31362b519422fc0169.
//
// Solidity: event AppointeeSet(address indexed account, address indexed appointee, address target, bytes4 selector)
func (_PermissionController *PermissionControllerFilterer) ParseAppointeeSet(log types.Log) (*PermissionControllerAppointeeSet, error) {
	event := new(PermissionControllerAppointeeSet)
	if err := _PermissionController.contract.UnpackLog(event, "AppointeeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionControllerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the PermissionController contract.
type PermissionControllerInitializedIterator struct {
	Event *PermissionControllerInitialized // Event containing the contract specifics and raw log

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
func (it *PermissionControllerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionControllerInitialized)
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
		it.Event = new(PermissionControllerInitialized)
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
func (it *PermissionControllerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionControllerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionControllerInitialized represents a Initialized event raised by the PermissionController contract.
type PermissionControllerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PermissionController *PermissionControllerFilterer) FilterInitialized(opts *bind.FilterOpts) (*PermissionControllerInitializedIterator, error) {

	logs, sub, err := _PermissionController.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PermissionControllerInitializedIterator{contract: _PermissionController.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PermissionController *PermissionControllerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PermissionControllerInitialized) (event.Subscription, error) {

	logs, sub, err := _PermissionController.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionControllerInitialized)
				if err := _PermissionController.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_PermissionController *PermissionControllerFilterer) ParseInitialized(log types.Log) (*PermissionControllerInitialized, error) {
	event := new(PermissionControllerInitialized)
	if err := _PermissionController.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionControllerPendingAdminAddedIterator is returned from FilterPendingAdminAdded and is used to iterate over the raw logs and unpacked data for PendingAdminAdded events raised by the PermissionController contract.
type PermissionControllerPendingAdminAddedIterator struct {
	Event *PermissionControllerPendingAdminAdded // Event containing the contract specifics and raw log

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
func (it *PermissionControllerPendingAdminAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionControllerPendingAdminAdded)
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
		it.Event = new(PermissionControllerPendingAdminAdded)
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
func (it *PermissionControllerPendingAdminAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionControllerPendingAdminAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionControllerPendingAdminAdded represents a PendingAdminAdded event raised by the PermissionController contract.
type PermissionControllerPendingAdminAdded struct {
	Account common.Address
	Admin   common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPendingAdminAdded is a free log retrieval operation binding the contract event 0xb14b9a3d448c5b04f0e5b087b6f5193390db7955482a6ffb841e7b3ba61a460c.
//
// Solidity: event PendingAdminAdded(address indexed account, address admin)
func (_PermissionController *PermissionControllerFilterer) FilterPendingAdminAdded(opts *bind.FilterOpts, account []common.Address) (*PermissionControllerPendingAdminAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PermissionController.contract.FilterLogs(opts, "PendingAdminAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &PermissionControllerPendingAdminAddedIterator{contract: _PermissionController.contract, event: "PendingAdminAdded", logs: logs, sub: sub}, nil
}

// WatchPendingAdminAdded is a free log subscription operation binding the contract event 0xb14b9a3d448c5b04f0e5b087b6f5193390db7955482a6ffb841e7b3ba61a460c.
//
// Solidity: event PendingAdminAdded(address indexed account, address admin)
func (_PermissionController *PermissionControllerFilterer) WatchPendingAdminAdded(opts *bind.WatchOpts, sink chan<- *PermissionControllerPendingAdminAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PermissionController.contract.WatchLogs(opts, "PendingAdminAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionControllerPendingAdminAdded)
				if err := _PermissionController.contract.UnpackLog(event, "PendingAdminAdded", log); err != nil {
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

// ParsePendingAdminAdded is a log parse operation binding the contract event 0xb14b9a3d448c5b04f0e5b087b6f5193390db7955482a6ffb841e7b3ba61a460c.
//
// Solidity: event PendingAdminAdded(address indexed account, address admin)
func (_PermissionController *PermissionControllerFilterer) ParsePendingAdminAdded(log types.Log) (*PermissionControllerPendingAdminAdded, error) {
	event := new(PermissionControllerPendingAdminAdded)
	if err := _PermissionController.contract.UnpackLog(event, "PendingAdminAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionControllerPendingAdminRemovedIterator is returned from FilterPendingAdminRemoved and is used to iterate over the raw logs and unpacked data for PendingAdminRemoved events raised by the PermissionController contract.
type PermissionControllerPendingAdminRemovedIterator struct {
	Event *PermissionControllerPendingAdminRemoved // Event containing the contract specifics and raw log

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
func (it *PermissionControllerPendingAdminRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionControllerPendingAdminRemoved)
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
		it.Event = new(PermissionControllerPendingAdminRemoved)
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
func (it *PermissionControllerPendingAdminRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionControllerPendingAdminRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionControllerPendingAdminRemoved represents a PendingAdminRemoved event raised by the PermissionController contract.
type PermissionControllerPendingAdminRemoved struct {
	Account common.Address
	Admin   common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPendingAdminRemoved is a free log retrieval operation binding the contract event 0xd706ed7ae044d795b49e54c9f519f663053951011985f663a862cd9ee72a9ac7.
//
// Solidity: event PendingAdminRemoved(address indexed account, address admin)
func (_PermissionController *PermissionControllerFilterer) FilterPendingAdminRemoved(opts *bind.FilterOpts, account []common.Address) (*PermissionControllerPendingAdminRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PermissionController.contract.FilterLogs(opts, "PendingAdminRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &PermissionControllerPendingAdminRemovedIterator{contract: _PermissionController.contract, event: "PendingAdminRemoved", logs: logs, sub: sub}, nil
}

// WatchPendingAdminRemoved is a free log subscription operation binding the contract event 0xd706ed7ae044d795b49e54c9f519f663053951011985f663a862cd9ee72a9ac7.
//
// Solidity: event PendingAdminRemoved(address indexed account, address admin)
func (_PermissionController *PermissionControllerFilterer) WatchPendingAdminRemoved(opts *bind.WatchOpts, sink chan<- *PermissionControllerPendingAdminRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PermissionController.contract.WatchLogs(opts, "PendingAdminRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionControllerPendingAdminRemoved)
				if err := _PermissionController.contract.UnpackLog(event, "PendingAdminRemoved", log); err != nil {
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

// ParsePendingAdminRemoved is a log parse operation binding the contract event 0xd706ed7ae044d795b49e54c9f519f663053951011985f663a862cd9ee72a9ac7.
//
// Solidity: event PendingAdminRemoved(address indexed account, address admin)
func (_PermissionController *PermissionControllerFilterer) ParsePendingAdminRemoved(log types.Log) (*PermissionControllerPendingAdminRemoved, error) {
	event := new(PermissionControllerPendingAdminRemoved)
	if err := _PermissionController.contract.UnpackLog(event, "PendingAdminRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
