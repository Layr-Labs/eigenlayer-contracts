// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IComputeRegistry

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

// IComputeRegistryTypesTOSSignature is an auto generated low-level Go binding around an user-defined struct.
type IComputeRegistryTypesTOSSignature struct {
	Signer    common.Address
	TosHash   [32]byte
	Signature []byte
}

// OperatorSet is an auto generated low-level Go binding around an user-defined struct.
type OperatorSet struct {
	Avs common.Address
	Id  uint32
}

// IComputeRegistryMetaData contains all meta data concerning the IComputeRegistry contract.
var IComputeRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"ALLOCATION_MANAGER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"CROSS_CHAIN_REGISTRY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractICrossChainRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"KEY_REGISTRAR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIKeyRegistrar\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_EXPIRY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RELEASE_MANAGER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIReleaseManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TOS_AGREEMENT_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TOS_HASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateTOSAgreementDigest\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregisterFromCompute\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getOperatorSetTosSignature\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIComputeRegistryTypes.TOSSignature\",\"components\":[{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tosHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperatorSetRegistered\",\"inputs\":[{\"name\":\"operatorSetKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerForCompute\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"OperatorSetDeregistered\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":true,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSetRegistered\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":true,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"signer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tosHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CurveTypeNotSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperatorSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTOSSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoActiveGenerationReservation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorSetAlreadyRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorSetNotRegistered\",\"inputs\":[]}]",
}

// IComputeRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use IComputeRegistryMetaData.ABI instead.
var IComputeRegistryABI = IComputeRegistryMetaData.ABI

// IComputeRegistry is an auto generated Go binding around an Ethereum contract.
type IComputeRegistry struct {
	IComputeRegistryCaller     // Read-only binding to the contract
	IComputeRegistryTransactor // Write-only binding to the contract
	IComputeRegistryFilterer   // Log filterer for contract events
}

// IComputeRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IComputeRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IComputeRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IComputeRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IComputeRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IComputeRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IComputeRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IComputeRegistrySession struct {
	Contract     *IComputeRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IComputeRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IComputeRegistryCallerSession struct {
	Contract *IComputeRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IComputeRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IComputeRegistryTransactorSession struct {
	Contract     *IComputeRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IComputeRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IComputeRegistryRaw struct {
	Contract *IComputeRegistry // Generic contract binding to access the raw methods on
}

// IComputeRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IComputeRegistryCallerRaw struct {
	Contract *IComputeRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// IComputeRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IComputeRegistryTransactorRaw struct {
	Contract *IComputeRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIComputeRegistry creates a new instance of IComputeRegistry, bound to a specific deployed contract.
func NewIComputeRegistry(address common.Address, backend bind.ContractBackend) (*IComputeRegistry, error) {
	contract, err := bindIComputeRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IComputeRegistry{IComputeRegistryCaller: IComputeRegistryCaller{contract: contract}, IComputeRegistryTransactor: IComputeRegistryTransactor{contract: contract}, IComputeRegistryFilterer: IComputeRegistryFilterer{contract: contract}}, nil
}

// NewIComputeRegistryCaller creates a new read-only instance of IComputeRegistry, bound to a specific deployed contract.
func NewIComputeRegistryCaller(address common.Address, caller bind.ContractCaller) (*IComputeRegistryCaller, error) {
	contract, err := bindIComputeRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IComputeRegistryCaller{contract: contract}, nil
}

// NewIComputeRegistryTransactor creates a new write-only instance of IComputeRegistry, bound to a specific deployed contract.
func NewIComputeRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*IComputeRegistryTransactor, error) {
	contract, err := bindIComputeRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IComputeRegistryTransactor{contract: contract}, nil
}

// NewIComputeRegistryFilterer creates a new log filterer instance of IComputeRegistry, bound to a specific deployed contract.
func NewIComputeRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*IComputeRegistryFilterer, error) {
	contract, err := bindIComputeRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IComputeRegistryFilterer{contract: contract}, nil
}

// bindIComputeRegistry binds a generic wrapper to an already deployed contract.
func bindIComputeRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IComputeRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IComputeRegistry *IComputeRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IComputeRegistry.Contract.IComputeRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IComputeRegistry *IComputeRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IComputeRegistry.Contract.IComputeRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IComputeRegistry *IComputeRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IComputeRegistry.Contract.IComputeRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IComputeRegistry *IComputeRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IComputeRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IComputeRegistry *IComputeRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IComputeRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IComputeRegistry *IComputeRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IComputeRegistry.Contract.contract.Transact(opts, method, params...)
}

// ALLOCATIONMANAGER is a free data retrieval call binding the contract method 0x31232bc9.
//
// Solidity: function ALLOCATION_MANAGER() view returns(address)
func (_IComputeRegistry *IComputeRegistryCaller) ALLOCATIONMANAGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IComputeRegistry.contract.Call(opts, &out, "ALLOCATION_MANAGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ALLOCATIONMANAGER is a free data retrieval call binding the contract method 0x31232bc9.
//
// Solidity: function ALLOCATION_MANAGER() view returns(address)
func (_IComputeRegistry *IComputeRegistrySession) ALLOCATIONMANAGER() (common.Address, error) {
	return _IComputeRegistry.Contract.ALLOCATIONMANAGER(&_IComputeRegistry.CallOpts)
}

// ALLOCATIONMANAGER is a free data retrieval call binding the contract method 0x31232bc9.
//
// Solidity: function ALLOCATION_MANAGER() view returns(address)
func (_IComputeRegistry *IComputeRegistryCallerSession) ALLOCATIONMANAGER() (common.Address, error) {
	return _IComputeRegistry.Contract.ALLOCATIONMANAGER(&_IComputeRegistry.CallOpts)
}

// CROSSCHAINREGISTRY is a free data retrieval call binding the contract method 0x9b250844.
//
// Solidity: function CROSS_CHAIN_REGISTRY() view returns(address)
func (_IComputeRegistry *IComputeRegistryCaller) CROSSCHAINREGISTRY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IComputeRegistry.contract.Call(opts, &out, "CROSS_CHAIN_REGISTRY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CROSSCHAINREGISTRY is a free data retrieval call binding the contract method 0x9b250844.
//
// Solidity: function CROSS_CHAIN_REGISTRY() view returns(address)
func (_IComputeRegistry *IComputeRegistrySession) CROSSCHAINREGISTRY() (common.Address, error) {
	return _IComputeRegistry.Contract.CROSSCHAINREGISTRY(&_IComputeRegistry.CallOpts)
}

// CROSSCHAINREGISTRY is a free data retrieval call binding the contract method 0x9b250844.
//
// Solidity: function CROSS_CHAIN_REGISTRY() view returns(address)
func (_IComputeRegistry *IComputeRegistryCallerSession) CROSSCHAINREGISTRY() (common.Address, error) {
	return _IComputeRegistry.Contract.CROSSCHAINREGISTRY(&_IComputeRegistry.CallOpts)
}

// KEYREGISTRAR is a free data retrieval call binding the contract method 0xe6414b48.
//
// Solidity: function KEY_REGISTRAR() view returns(address)
func (_IComputeRegistry *IComputeRegistryCaller) KEYREGISTRAR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IComputeRegistry.contract.Call(opts, &out, "KEY_REGISTRAR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// KEYREGISTRAR is a free data retrieval call binding the contract method 0xe6414b48.
//
// Solidity: function KEY_REGISTRAR() view returns(address)
func (_IComputeRegistry *IComputeRegistrySession) KEYREGISTRAR() (common.Address, error) {
	return _IComputeRegistry.Contract.KEYREGISTRAR(&_IComputeRegistry.CallOpts)
}

// KEYREGISTRAR is a free data retrieval call binding the contract method 0xe6414b48.
//
// Solidity: function KEY_REGISTRAR() view returns(address)
func (_IComputeRegistry *IComputeRegistryCallerSession) KEYREGISTRAR() (common.Address, error) {
	return _IComputeRegistry.Contract.KEYREGISTRAR(&_IComputeRegistry.CallOpts)
}

// MAXEXPIRY is a free data retrieval call binding the contract method 0xb9671690.
//
// Solidity: function MAX_EXPIRY() view returns(uint256)
func (_IComputeRegistry *IComputeRegistryCaller) MAXEXPIRY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IComputeRegistry.contract.Call(opts, &out, "MAX_EXPIRY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXEXPIRY is a free data retrieval call binding the contract method 0xb9671690.
//
// Solidity: function MAX_EXPIRY() view returns(uint256)
func (_IComputeRegistry *IComputeRegistrySession) MAXEXPIRY() (*big.Int, error) {
	return _IComputeRegistry.Contract.MAXEXPIRY(&_IComputeRegistry.CallOpts)
}

// MAXEXPIRY is a free data retrieval call binding the contract method 0xb9671690.
//
// Solidity: function MAX_EXPIRY() view returns(uint256)
func (_IComputeRegistry *IComputeRegistryCallerSession) MAXEXPIRY() (*big.Int, error) {
	return _IComputeRegistry.Contract.MAXEXPIRY(&_IComputeRegistry.CallOpts)
}

// RELEASEMANAGER is a free data retrieval call binding the contract method 0xb39d254f.
//
// Solidity: function RELEASE_MANAGER() view returns(address)
func (_IComputeRegistry *IComputeRegistryCaller) RELEASEMANAGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IComputeRegistry.contract.Call(opts, &out, "RELEASE_MANAGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RELEASEMANAGER is a free data retrieval call binding the contract method 0xb39d254f.
//
// Solidity: function RELEASE_MANAGER() view returns(address)
func (_IComputeRegistry *IComputeRegistrySession) RELEASEMANAGER() (common.Address, error) {
	return _IComputeRegistry.Contract.RELEASEMANAGER(&_IComputeRegistry.CallOpts)
}

// RELEASEMANAGER is a free data retrieval call binding the contract method 0xb39d254f.
//
// Solidity: function RELEASE_MANAGER() view returns(address)
func (_IComputeRegistry *IComputeRegistryCallerSession) RELEASEMANAGER() (common.Address, error) {
	return _IComputeRegistry.Contract.RELEASEMANAGER(&_IComputeRegistry.CallOpts)
}

// TOSAGREEMENTTYPEHASH is a free data retrieval call binding the contract method 0x1de02dbb.
//
// Solidity: function TOS_AGREEMENT_TYPEHASH() view returns(bytes32)
func (_IComputeRegistry *IComputeRegistryCaller) TOSAGREEMENTTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IComputeRegistry.contract.Call(opts, &out, "TOS_AGREEMENT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TOSAGREEMENTTYPEHASH is a free data retrieval call binding the contract method 0x1de02dbb.
//
// Solidity: function TOS_AGREEMENT_TYPEHASH() view returns(bytes32)
func (_IComputeRegistry *IComputeRegistrySession) TOSAGREEMENTTYPEHASH() ([32]byte, error) {
	return _IComputeRegistry.Contract.TOSAGREEMENTTYPEHASH(&_IComputeRegistry.CallOpts)
}

// TOSAGREEMENTTYPEHASH is a free data retrieval call binding the contract method 0x1de02dbb.
//
// Solidity: function TOS_AGREEMENT_TYPEHASH() view returns(bytes32)
func (_IComputeRegistry *IComputeRegistryCallerSession) TOSAGREEMENTTYPEHASH() ([32]byte, error) {
	return _IComputeRegistry.Contract.TOSAGREEMENTTYPEHASH(&_IComputeRegistry.CallOpts)
}

// TOSHASH is a free data retrieval call binding the contract method 0x8df643c7.
//
// Solidity: function TOS_HASH() view returns(bytes32)
func (_IComputeRegistry *IComputeRegistryCaller) TOSHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IComputeRegistry.contract.Call(opts, &out, "TOS_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TOSHASH is a free data retrieval call binding the contract method 0x8df643c7.
//
// Solidity: function TOS_HASH() view returns(bytes32)
func (_IComputeRegistry *IComputeRegistrySession) TOSHASH() ([32]byte, error) {
	return _IComputeRegistry.Contract.TOSHASH(&_IComputeRegistry.CallOpts)
}

// TOSHASH is a free data retrieval call binding the contract method 0x8df643c7.
//
// Solidity: function TOS_HASH() view returns(bytes32)
func (_IComputeRegistry *IComputeRegistryCallerSession) TOSHASH() ([32]byte, error) {
	return _IComputeRegistry.Contract.TOSHASH(&_IComputeRegistry.CallOpts)
}

// CalculateTOSAgreementDigest is a free data retrieval call binding the contract method 0x7a1bb660.
//
// Solidity: function calculateTOSAgreementDigest((address,uint32) operatorSet, address signer) view returns(bytes32)
func (_IComputeRegistry *IComputeRegistryCaller) CalculateTOSAgreementDigest(opts *bind.CallOpts, operatorSet OperatorSet, signer common.Address) ([32]byte, error) {
	var out []interface{}
	err := _IComputeRegistry.contract.Call(opts, &out, "calculateTOSAgreementDigest", operatorSet, signer)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateTOSAgreementDigest is a free data retrieval call binding the contract method 0x7a1bb660.
//
// Solidity: function calculateTOSAgreementDigest((address,uint32) operatorSet, address signer) view returns(bytes32)
func (_IComputeRegistry *IComputeRegistrySession) CalculateTOSAgreementDigest(operatorSet OperatorSet, signer common.Address) ([32]byte, error) {
	return _IComputeRegistry.Contract.CalculateTOSAgreementDigest(&_IComputeRegistry.CallOpts, operatorSet, signer)
}

// CalculateTOSAgreementDigest is a free data retrieval call binding the contract method 0x7a1bb660.
//
// Solidity: function calculateTOSAgreementDigest((address,uint32) operatorSet, address signer) view returns(bytes32)
func (_IComputeRegistry *IComputeRegistryCallerSession) CalculateTOSAgreementDigest(operatorSet OperatorSet, signer common.Address) ([32]byte, error) {
	return _IComputeRegistry.Contract.CalculateTOSAgreementDigest(&_IComputeRegistry.CallOpts, operatorSet, signer)
}

// GetOperatorSetTosSignature is a free data retrieval call binding the contract method 0x130d6165.
//
// Solidity: function getOperatorSetTosSignature((address,uint32) operatorSet) view returns((address,bytes32,bytes))
func (_IComputeRegistry *IComputeRegistryCaller) GetOperatorSetTosSignature(opts *bind.CallOpts, operatorSet OperatorSet) (IComputeRegistryTypesTOSSignature, error) {
	var out []interface{}
	err := _IComputeRegistry.contract.Call(opts, &out, "getOperatorSetTosSignature", operatorSet)

	if err != nil {
		return *new(IComputeRegistryTypesTOSSignature), err
	}

	out0 := *abi.ConvertType(out[0], new(IComputeRegistryTypesTOSSignature)).(*IComputeRegistryTypesTOSSignature)

	return out0, err

}

// GetOperatorSetTosSignature is a free data retrieval call binding the contract method 0x130d6165.
//
// Solidity: function getOperatorSetTosSignature((address,uint32) operatorSet) view returns((address,bytes32,bytes))
func (_IComputeRegistry *IComputeRegistrySession) GetOperatorSetTosSignature(operatorSet OperatorSet) (IComputeRegistryTypesTOSSignature, error) {
	return _IComputeRegistry.Contract.GetOperatorSetTosSignature(&_IComputeRegistry.CallOpts, operatorSet)
}

// GetOperatorSetTosSignature is a free data retrieval call binding the contract method 0x130d6165.
//
// Solidity: function getOperatorSetTosSignature((address,uint32) operatorSet) view returns((address,bytes32,bytes))
func (_IComputeRegistry *IComputeRegistryCallerSession) GetOperatorSetTosSignature(operatorSet OperatorSet) (IComputeRegistryTypesTOSSignature, error) {
	return _IComputeRegistry.Contract.GetOperatorSetTosSignature(&_IComputeRegistry.CallOpts, operatorSet)
}

// IsOperatorSetRegistered is a free data retrieval call binding the contract method 0xc4a1ca05.
//
// Solidity: function isOperatorSetRegistered(bytes32 operatorSetKey) view returns(bool)
func (_IComputeRegistry *IComputeRegistryCaller) IsOperatorSetRegistered(opts *bind.CallOpts, operatorSetKey [32]byte) (bool, error) {
	var out []interface{}
	err := _IComputeRegistry.contract.Call(opts, &out, "isOperatorSetRegistered", operatorSetKey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperatorSetRegistered is a free data retrieval call binding the contract method 0xc4a1ca05.
//
// Solidity: function isOperatorSetRegistered(bytes32 operatorSetKey) view returns(bool)
func (_IComputeRegistry *IComputeRegistrySession) IsOperatorSetRegistered(operatorSetKey [32]byte) (bool, error) {
	return _IComputeRegistry.Contract.IsOperatorSetRegistered(&_IComputeRegistry.CallOpts, operatorSetKey)
}

// IsOperatorSetRegistered is a free data retrieval call binding the contract method 0xc4a1ca05.
//
// Solidity: function isOperatorSetRegistered(bytes32 operatorSetKey) view returns(bool)
func (_IComputeRegistry *IComputeRegistryCallerSession) IsOperatorSetRegistered(operatorSetKey [32]byte) (bool, error) {
	return _IComputeRegistry.Contract.IsOperatorSetRegistered(&_IComputeRegistry.CallOpts, operatorSetKey)
}

// DeregisterFromCompute is a paid mutator transaction binding the contract method 0x89fec15f.
//
// Solidity: function deregisterFromCompute((address,uint32) operatorSet) returns()
func (_IComputeRegistry *IComputeRegistryTransactor) DeregisterFromCompute(opts *bind.TransactOpts, operatorSet OperatorSet) (*types.Transaction, error) {
	return _IComputeRegistry.contract.Transact(opts, "deregisterFromCompute", operatorSet)
}

// DeregisterFromCompute is a paid mutator transaction binding the contract method 0x89fec15f.
//
// Solidity: function deregisterFromCompute((address,uint32) operatorSet) returns()
func (_IComputeRegistry *IComputeRegistrySession) DeregisterFromCompute(operatorSet OperatorSet) (*types.Transaction, error) {
	return _IComputeRegistry.Contract.DeregisterFromCompute(&_IComputeRegistry.TransactOpts, operatorSet)
}

// DeregisterFromCompute is a paid mutator transaction binding the contract method 0x89fec15f.
//
// Solidity: function deregisterFromCompute((address,uint32) operatorSet) returns()
func (_IComputeRegistry *IComputeRegistryTransactorSession) DeregisterFromCompute(operatorSet OperatorSet) (*types.Transaction, error) {
	return _IComputeRegistry.Contract.DeregisterFromCompute(&_IComputeRegistry.TransactOpts, operatorSet)
}

// RegisterForCompute is a paid mutator transaction binding the contract method 0x536b2353.
//
// Solidity: function registerForCompute((address,uint32) operatorSet, bytes signature) returns()
func (_IComputeRegistry *IComputeRegistryTransactor) RegisterForCompute(opts *bind.TransactOpts, operatorSet OperatorSet, signature []byte) (*types.Transaction, error) {
	return _IComputeRegistry.contract.Transact(opts, "registerForCompute", operatorSet, signature)
}

// RegisterForCompute is a paid mutator transaction binding the contract method 0x536b2353.
//
// Solidity: function registerForCompute((address,uint32) operatorSet, bytes signature) returns()
func (_IComputeRegistry *IComputeRegistrySession) RegisterForCompute(operatorSet OperatorSet, signature []byte) (*types.Transaction, error) {
	return _IComputeRegistry.Contract.RegisterForCompute(&_IComputeRegistry.TransactOpts, operatorSet, signature)
}

// RegisterForCompute is a paid mutator transaction binding the contract method 0x536b2353.
//
// Solidity: function registerForCompute((address,uint32) operatorSet, bytes signature) returns()
func (_IComputeRegistry *IComputeRegistryTransactorSession) RegisterForCompute(operatorSet OperatorSet, signature []byte) (*types.Transaction, error) {
	return _IComputeRegistry.Contract.RegisterForCompute(&_IComputeRegistry.TransactOpts, operatorSet, signature)
}

// IComputeRegistryOperatorSetDeregisteredIterator is returned from FilterOperatorSetDeregistered and is used to iterate over the raw logs and unpacked data for OperatorSetDeregistered events raised by the IComputeRegistry contract.
type IComputeRegistryOperatorSetDeregisteredIterator struct {
	Event *IComputeRegistryOperatorSetDeregistered // Event containing the contract specifics and raw log

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
func (it *IComputeRegistryOperatorSetDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IComputeRegistryOperatorSetDeregistered)
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
		it.Event = new(IComputeRegistryOperatorSetDeregistered)
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
func (it *IComputeRegistryOperatorSetDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IComputeRegistryOperatorSetDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IComputeRegistryOperatorSetDeregistered represents a OperatorSetDeregistered event raised by the IComputeRegistry contract.
type IComputeRegistryOperatorSetDeregistered struct {
	OperatorSet OperatorSet
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorSetDeregistered is a free log retrieval operation binding the contract event 0x1d9dd94cbbd6d9d2bc86468a197493d2b960bfc33eb09ec2c9f1731e60c45988.
//
// Solidity: event OperatorSetDeregistered((address,uint32) indexed operatorSet)
func (_IComputeRegistry *IComputeRegistryFilterer) FilterOperatorSetDeregistered(opts *bind.FilterOpts, operatorSet []OperatorSet) (*IComputeRegistryOperatorSetDeregisteredIterator, error) {

	var operatorSetRule []interface{}
	for _, operatorSetItem := range operatorSet {
		operatorSetRule = append(operatorSetRule, operatorSetItem)
	}

	logs, sub, err := _IComputeRegistry.contract.FilterLogs(opts, "OperatorSetDeregistered", operatorSetRule)
	if err != nil {
		return nil, err
	}
	return &IComputeRegistryOperatorSetDeregisteredIterator{contract: _IComputeRegistry.contract, event: "OperatorSetDeregistered", logs: logs, sub: sub}, nil
}

// WatchOperatorSetDeregistered is a free log subscription operation binding the contract event 0x1d9dd94cbbd6d9d2bc86468a197493d2b960bfc33eb09ec2c9f1731e60c45988.
//
// Solidity: event OperatorSetDeregistered((address,uint32) indexed operatorSet)
func (_IComputeRegistry *IComputeRegistryFilterer) WatchOperatorSetDeregistered(opts *bind.WatchOpts, sink chan<- *IComputeRegistryOperatorSetDeregistered, operatorSet []OperatorSet) (event.Subscription, error) {

	var operatorSetRule []interface{}
	for _, operatorSetItem := range operatorSet {
		operatorSetRule = append(operatorSetRule, operatorSetItem)
	}

	logs, sub, err := _IComputeRegistry.contract.WatchLogs(opts, "OperatorSetDeregistered", operatorSetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IComputeRegistryOperatorSetDeregistered)
				if err := _IComputeRegistry.contract.UnpackLog(event, "OperatorSetDeregistered", log); err != nil {
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

// ParseOperatorSetDeregistered is a log parse operation binding the contract event 0x1d9dd94cbbd6d9d2bc86468a197493d2b960bfc33eb09ec2c9f1731e60c45988.
//
// Solidity: event OperatorSetDeregistered((address,uint32) indexed operatorSet)
func (_IComputeRegistry *IComputeRegistryFilterer) ParseOperatorSetDeregistered(log types.Log) (*IComputeRegistryOperatorSetDeregistered, error) {
	event := new(IComputeRegistryOperatorSetDeregistered)
	if err := _IComputeRegistry.contract.UnpackLog(event, "OperatorSetDeregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IComputeRegistryOperatorSetRegisteredIterator is returned from FilterOperatorSetRegistered and is used to iterate over the raw logs and unpacked data for OperatorSetRegistered events raised by the IComputeRegistry contract.
type IComputeRegistryOperatorSetRegisteredIterator struct {
	Event *IComputeRegistryOperatorSetRegistered // Event containing the contract specifics and raw log

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
func (it *IComputeRegistryOperatorSetRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IComputeRegistryOperatorSetRegistered)
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
		it.Event = new(IComputeRegistryOperatorSetRegistered)
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
func (it *IComputeRegistryOperatorSetRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IComputeRegistryOperatorSetRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IComputeRegistryOperatorSetRegistered represents a OperatorSetRegistered event raised by the IComputeRegistry contract.
type IComputeRegistryOperatorSetRegistered struct {
	OperatorSet OperatorSet
	Signer      common.Address
	TosHash     [32]byte
	Signature   []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorSetRegistered is a free log retrieval operation binding the contract event 0xab9d9ee579ad2bc5e08fba9e2b3d59225f1400f1a5830745283697e35ce2dcc8.
//
// Solidity: event OperatorSetRegistered((address,uint32) indexed operatorSet, address indexed signer, bytes32 indexed tosHash, bytes signature)
func (_IComputeRegistry *IComputeRegistryFilterer) FilterOperatorSetRegistered(opts *bind.FilterOpts, operatorSet []OperatorSet, signer []common.Address, tosHash [][32]byte) (*IComputeRegistryOperatorSetRegisteredIterator, error) {

	var operatorSetRule []interface{}
	for _, operatorSetItem := range operatorSet {
		operatorSetRule = append(operatorSetRule, operatorSetItem)
	}
	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}
	var tosHashRule []interface{}
	for _, tosHashItem := range tosHash {
		tosHashRule = append(tosHashRule, tosHashItem)
	}

	logs, sub, err := _IComputeRegistry.contract.FilterLogs(opts, "OperatorSetRegistered", operatorSetRule, signerRule, tosHashRule)
	if err != nil {
		return nil, err
	}
	return &IComputeRegistryOperatorSetRegisteredIterator{contract: _IComputeRegistry.contract, event: "OperatorSetRegistered", logs: logs, sub: sub}, nil
}

// WatchOperatorSetRegistered is a free log subscription operation binding the contract event 0xab9d9ee579ad2bc5e08fba9e2b3d59225f1400f1a5830745283697e35ce2dcc8.
//
// Solidity: event OperatorSetRegistered((address,uint32) indexed operatorSet, address indexed signer, bytes32 indexed tosHash, bytes signature)
func (_IComputeRegistry *IComputeRegistryFilterer) WatchOperatorSetRegistered(opts *bind.WatchOpts, sink chan<- *IComputeRegistryOperatorSetRegistered, operatorSet []OperatorSet, signer []common.Address, tosHash [][32]byte) (event.Subscription, error) {

	var operatorSetRule []interface{}
	for _, operatorSetItem := range operatorSet {
		operatorSetRule = append(operatorSetRule, operatorSetItem)
	}
	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}
	var tosHashRule []interface{}
	for _, tosHashItem := range tosHash {
		tosHashRule = append(tosHashRule, tosHashItem)
	}

	logs, sub, err := _IComputeRegistry.contract.WatchLogs(opts, "OperatorSetRegistered", operatorSetRule, signerRule, tosHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IComputeRegistryOperatorSetRegistered)
				if err := _IComputeRegistry.contract.UnpackLog(event, "OperatorSetRegistered", log); err != nil {
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

// ParseOperatorSetRegistered is a log parse operation binding the contract event 0xab9d9ee579ad2bc5e08fba9e2b3d59225f1400f1a5830745283697e35ce2dcc8.
//
// Solidity: event OperatorSetRegistered((address,uint32) indexed operatorSet, address indexed signer, bytes32 indexed tosHash, bytes signature)
func (_IComputeRegistry *IComputeRegistryFilterer) ParseOperatorSetRegistered(log types.Log) (*IComputeRegistryOperatorSetRegistered, error) {
	event := new(IComputeRegistryOperatorSetRegistered)
	if err := _IComputeRegistry.contract.UnpackLog(event, "OperatorSetRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
