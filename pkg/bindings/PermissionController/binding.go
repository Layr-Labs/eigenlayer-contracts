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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptAdmin\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addPendingAdmin\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"canCall\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAdmins\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAppointeePermissions\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"appointee\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAppointees\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPendingAdmins\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isAdmin\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isPendingAdmin\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pendingAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeAdmin\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeAppointee\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"appointee\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removePendingAdmin\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAppointee\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"appointee\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AdminRemoved\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"admin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AdminSet\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"admin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AppointeeRemoved\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"appointee\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"indexed\":false,\"internalType\":\"bytes4\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AppointeeSet\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"appointee\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"indexed\":false,\"internalType\":\"bytes4\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PendingAdminAdded\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"admin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PendingAdminRemoved\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"admin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AdminAlreadyPending\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AdminAlreadySet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AdminNotPending\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AdminNotSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AppointeeAlreadySet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AppointeeNotSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotHaveZeroAdmins\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAdmin\",\"inputs\":[]}]",
	Bin: "0x6080604052348015600e575f5ffd5b5060156019565b60d3565b5f54610100900460ff161560835760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff9081161460d1575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b61107c806100e05f395ff3fe608060405234801561000f575f5ffd5b50600436106100e5575f3560e01c80639100674511610088578063ad8aca7711610063578063ad8aca77146101d2578063df595cb8146101e5578063eb5a4e87146101f8578063fddbdefd1461020b575f5ffd5b80639100674514610189578063950d806e146101ac578063ad5f2210146101bf575f5ffd5b8063628806ef116100c3578063628806ef146101245780636bddfa1f146101375780638129fc1c14610160578063882a3b3814610168575f5ffd5b806306641201146100e9578063268959e5146100fe5780634f906cf914610111575b5f5ffd5b6100fc6100f7366004610e36565b61021e565b005b6100fc61010c366004610e87565b61033f565b6100fc61011f366004610e87565b61041a565b6100fc610132366004610eb8565b6104bd565b61014a610145366004610eb8565b61054b565b6040516101579190610f14565b60405180910390f35b6100fc610574565b61017b610176366004610e87565b61067b565b604051610157929190610f26565b61019c610197366004610e87565b6107ef565b6040519015158152602001610157565b6100fc6101ba366004610e36565b61085f565b61014a6101cd366004610eb8565b610970565b61019c6101e0366004610e87565b610a16565b61019c6101f3366004610e36565b610a37565b6100fc610206366004610e87565b610a8c565b61014a610219366004610f88565b610b5a565b8361022981336107ef565b61024657604051637bfa4b9f60e01b815260040160405180910390fd5b6001600160a01b0385165f908152600160205260408120906102688585610b98565b6001600160a01b0387165f908152600484016020526040902090915061028e9082610bc5565b6102ab5760405163262118cd60e01b815260040160405180910390fd5b6001600160a01b0386165f90815260048301602052604090206102ce9082610bdc565b505f81815260058301602052604090206102e89087610be7565b50856001600160a01b0316876001600160a01b03167f18242326b6b862126970679759169f01f646bd55ec5bfcab85ba9f337a74e0c6878760405161032e929190610fc8565b60405180910390a350505050505050565b8161034a81336107ef565b61036757604051637bfa4b9f60e01b815260040160405180910390fd5b6001600160a01b0383165f9081526001602081905260409091206002019061038e82610bfb565b116103ac576040516310ce892b60e31b815260040160405180910390fd5b6103b68184610be7565b6103d357604051630716d81b60e51b815260040160405180910390fd5b6040516001600160a01b0384811682528516907fdb9d5d31320daf5bc7181d565b6da4d12e30f0f4d5aa324a992426c14a1d19ce906020015b60405180910390a250505050565b8161042581336107ef565b61044257604051637bfa4b9f60e01b815260040160405180910390fd5b6001600160a01b0383165f9081526001602052604090206104638184610be7565b6104805760405163bed8295f60e01b815260040160405180910390fd5b6040516001600160a01b0384811682528516907fd706ed7ae044d795b49e54c9f519f663053951011985f663a862cd9ee72a9ac79060200161040c565b6001600160a01b0381165f9081526001602052604090206104de8133610be7565b6104fb5760405163bed8295f60e01b815260040160405180910390fd5b6105086002820133610c04565b506040513381526001600160a01b038316907fbf265e8326285a2747e33e54d5945f7111f2b5edb826eb8c08d4677779b3ff979060200160405180910390a25050565b6001600160a01b0381165f90815260016020526040902060609061056e90610c18565b92915050565b5f54610100900460ff161580801561059257505f54600160ff909116105b806105ab5750303b1580156105ab57505f5460ff166001145b6106125760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840160405180910390fd5b5f805460ff191660011790558015610633575f805461ff0019166101001790555b8015610678575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50565b6001600160a01b038083165f9081526001602090815260408083209385168352600490930190529081206060918291906106b482610bfb565b90505f8167ffffffffffffffff8111156106d0576106d0610feb565b6040519080825280602002602001820160405280156106f9578160200160208202803683370190505b5090505f8267ffffffffffffffff81111561071657610716610feb565b60405190808252806020026020018201604052801561073f578160200160208202803683370190505b5090505f5b838110156107e1575f8061077461075b8885610c24565b606081901c9160a09190911b6001600160e01b03191690565b915091508185848151811061078b5761078b610fff565b60200260200101906001600160a01b031690816001600160a01b031681525050808484815181106107be576107be610fff565b6001600160e01b0319909216602092830291909101909101525050600101610744565b509097909650945050505050565b6001600160a01b0382165f90815260016020526040812061081290600201610bfb565b5f0361083457816001600160a01b0316836001600160a01b031614905061056e565b6001600160a01b0383165f9081526001602052604090206108589060020183610c2f565b9392505050565b8361086a81336107ef565b61088757604051637bfa4b9f60e01b815260040160405180910390fd5b6001600160a01b0385165f908152600160205260408120906108a98585610b98565b6001600160a01b0387165f90815260048401602052604090209091506108cf9082610bc5565b156108ed5760405163ad8efeb760e01b815260040160405180910390fd5b6001600160a01b0386165f90815260048301602052604090206109109082610c50565b505f818152600583016020526040902061092a9087610c04565b50856001600160a01b0316876001600160a01b03167f037f03a2ad6b967df4a01779b6d2b4c85950df83925d9e31362b519422fc0169878760405161032e929190610fc8565b6001600160a01b0381165f90815260016020526040902060609061099690600201610bfb565b5f036109ee576040805160018082528183019092525f916020808301908036833701905050905082815f815181106109d0576109d0610fff565b6001600160a01b039092166020928302919091019091015292915050565b6001600160a01b0382165f90815260016020526040902061056e90600201610c18565b919050565b6001600160a01b0382165f9081526001602052604081206108589083610c2f565b5f610a4285856107ef565b80610a835750610a83610a558484610b98565b6001600160a01b038088165f908152600160209081526040808320938a168352600490930190522090610bc5565b95945050505050565b81610a9781336107ef565b610ab457604051637bfa4b9f60e01b815260040160405180910390fd5b6001600160a01b0383165f908152600160205260409020610ad86002820184610c2f565b15610af65760405163130160e560e31b815260040160405180910390fd5b610b008184610c04565b610b1d576040516319abede360e11b815260040160405180910390fd5b6040516001600160a01b0384811682528516907fb14b9a3d448c5b04f0e5b087b6f5193390db7955482a6ffb841e7b3ba61a460c9060200161040c565b60605f610b678484610b98565b6001600160a01b0386165f9081526001602090815260408083208484526005019091529020909150610a8390610c18565b60609190911b6bffffffffffffffffffffffff191660a09190911c6bffffffff0000000000000000161790565b5f8181526001830160205260408120541515610858565b5f6108588383610c5b565b5f610858836001600160a01b038416610c5b565b5f61056e825490565b5f610858836001600160a01b038416610d3e565b60605f61085883610d8a565b5f6108588383610de3565b6001600160a01b0381165f9081526001830160205260408120541515610858565b5f6108588383610d3e565b5f8181526001830160205260408120548015610d35575f610c7d600183611013565b85549091505f90610c9090600190611013565b9050818114610cef575f865f018281548110610cae57610cae610fff565b905f5260205f200154905080875f018481548110610cce57610cce610fff565b5f918252602080832090910192909255918252600188019052604090208390555b8554869080610d0057610d00611032565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f90556001935050505061056e565b5f91505061056e565b5f818152600183016020526040812054610d8357508154600181810184555f84815260208082209093018490558454848252828601909352604090209190915561056e565b505f61056e565b6060815f01805480602002602001604051908101604052809291908181526020018280548015610dd757602002820191905f5260205f20905b815481526020019060010190808311610dc3575b50505050509050919050565b5f825f018281548110610df857610df8610fff565b905f5260205f200154905092915050565b80356001600160a01b0381168114610a11575f5ffd5b80356001600160e01b031981168114610a11575f5ffd5b5f5f5f5f60808587031215610e49575f5ffd5b610e5285610e09565b9350610e6060208601610e09565b9250610e6e60408601610e09565b9150610e7c60608601610e1f565b905092959194509250565b5f5f60408385031215610e98575f5ffd5b610ea183610e09565b9150610eaf60208401610e09565b90509250929050565b5f60208284031215610ec8575f5ffd5b61085882610e09565b5f8151808452602084019350602083015f5b82811015610f0a5781516001600160a01b0316865260209586019590910190600101610ee3565b5093949350505050565b602081525f6108586020830184610ed1565b604081525f610f386040830185610ed1565b82810360208401528084518083526020830191506020860192505f5b81811015610f7c5783516001600160e01b031916835260209384019390920191600101610f54565b50909695505050505050565b5f5f5f60608486031215610f9a575f5ffd5b610fa384610e09565b9250610fb160208501610e09565b9150610fbf60408501610e1f565b90509250925092565b6001600160a01b039290921682526001600160e01b031916602082015260400190565b634e487b7160e01b5f52604160045260245ffd5b634e487b7160e01b5f52603260045260245ffd5b8181038181111561056e57634e487b7160e01b5f52601160045260245ffd5b634e487b7160e01b5f52603160045260245ffdfea26469706673582212208ff3f5e26cf667d2994ce4894fb366fc8efa997d80b585e55d674cb4b0e20d9b64736f6c634300081b0033",
}

// PermissionControllerABI is the input ABI used to generate the binding from.
// Deprecated: Use PermissionControllerMetaData.ABI instead.
var PermissionControllerABI = PermissionControllerMetaData.ABI

// PermissionControllerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PermissionControllerMetaData.Bin instead.
var PermissionControllerBin = PermissionControllerMetaData.Bin

// DeployPermissionController deploys a new Ethereum contract, binding an instance of PermissionController to it.
func DeployPermissionController(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PermissionController, error) {
	parsed, err := PermissionControllerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PermissionControllerBin), backend)
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

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_PermissionController *PermissionControllerTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PermissionController.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_PermissionController *PermissionControllerSession) Initialize() (*types.Transaction, error) {
	return _PermissionController.Contract.Initialize(&_PermissionController.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_PermissionController *PermissionControllerTransactorSession) Initialize() (*types.Transaction, error) {
	return _PermissionController.Contract.Initialize(&_PermissionController.TransactOpts)
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
