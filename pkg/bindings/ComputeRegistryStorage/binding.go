// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ComputeRegistryStorage

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

// ComputeRegistryStorageMetaData contains all meta data concerning the ComputeRegistryStorage contract.
var ComputeRegistryStorageMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"ALLOCATION_MANAGER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"CROSS_CHAIN_REGISTRY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractICrossChainRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"KEY_REGISTRAR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIKeyRegistrar\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_EXPIRY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RELEASE_MANAGER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIReleaseManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TOS_AGREEMENT_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TOS_HASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateTOSAgreementDigest\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregisterFromCompute\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getOperatorSetTosSignature\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIComputeRegistryTypes.TOSSignature\",\"components\":[{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tosHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperatorSetRegistered\",\"inputs\":[{\"name\":\"operatorSetKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"isRegistered\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerForCompute\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"OperatorSetDeregistered\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":true,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSetRegistered\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":true,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"signer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tosHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CurveTypeNotSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperatorSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTOSSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoActiveGenerationReservation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorSetAlreadyRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorSetNotRegistered\",\"inputs\":[]}]",
}

// ComputeRegistryStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use ComputeRegistryStorageMetaData.ABI instead.
var ComputeRegistryStorageABI = ComputeRegistryStorageMetaData.ABI

// ComputeRegistryStorage is an auto generated Go binding around an Ethereum contract.
type ComputeRegistryStorage struct {
	ComputeRegistryStorageCaller     // Read-only binding to the contract
	ComputeRegistryStorageTransactor // Write-only binding to the contract
	ComputeRegistryStorageFilterer   // Log filterer for contract events
}

// ComputeRegistryStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type ComputeRegistryStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ComputeRegistryStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ComputeRegistryStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ComputeRegistryStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ComputeRegistryStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ComputeRegistryStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ComputeRegistryStorageSession struct {
	Contract     *ComputeRegistryStorage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ComputeRegistryStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ComputeRegistryStorageCallerSession struct {
	Contract *ComputeRegistryStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// ComputeRegistryStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ComputeRegistryStorageTransactorSession struct {
	Contract     *ComputeRegistryStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// ComputeRegistryStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type ComputeRegistryStorageRaw struct {
	Contract *ComputeRegistryStorage // Generic contract binding to access the raw methods on
}

// ComputeRegistryStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ComputeRegistryStorageCallerRaw struct {
	Contract *ComputeRegistryStorageCaller // Generic read-only contract binding to access the raw methods on
}

// ComputeRegistryStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ComputeRegistryStorageTransactorRaw struct {
	Contract *ComputeRegistryStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewComputeRegistryStorage creates a new instance of ComputeRegistryStorage, bound to a specific deployed contract.
func NewComputeRegistryStorage(address common.Address, backend bind.ContractBackend) (*ComputeRegistryStorage, error) {
	contract, err := bindComputeRegistryStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ComputeRegistryStorage{ComputeRegistryStorageCaller: ComputeRegistryStorageCaller{contract: contract}, ComputeRegistryStorageTransactor: ComputeRegistryStorageTransactor{contract: contract}, ComputeRegistryStorageFilterer: ComputeRegistryStorageFilterer{contract: contract}}, nil
}

// NewComputeRegistryStorageCaller creates a new read-only instance of ComputeRegistryStorage, bound to a specific deployed contract.
func NewComputeRegistryStorageCaller(address common.Address, caller bind.ContractCaller) (*ComputeRegistryStorageCaller, error) {
	contract, err := bindComputeRegistryStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ComputeRegistryStorageCaller{contract: contract}, nil
}

// NewComputeRegistryStorageTransactor creates a new write-only instance of ComputeRegistryStorage, bound to a specific deployed contract.
func NewComputeRegistryStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*ComputeRegistryStorageTransactor, error) {
	contract, err := bindComputeRegistryStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ComputeRegistryStorageTransactor{contract: contract}, nil
}

// NewComputeRegistryStorageFilterer creates a new log filterer instance of ComputeRegistryStorage, bound to a specific deployed contract.
func NewComputeRegistryStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*ComputeRegistryStorageFilterer, error) {
	contract, err := bindComputeRegistryStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ComputeRegistryStorageFilterer{contract: contract}, nil
}

// bindComputeRegistryStorage binds a generic wrapper to an already deployed contract.
func bindComputeRegistryStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ComputeRegistryStorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ComputeRegistryStorage *ComputeRegistryStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ComputeRegistryStorage.Contract.ComputeRegistryStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ComputeRegistryStorage *ComputeRegistryStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ComputeRegistryStorage.Contract.ComputeRegistryStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ComputeRegistryStorage *ComputeRegistryStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ComputeRegistryStorage.Contract.ComputeRegistryStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ComputeRegistryStorage *ComputeRegistryStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ComputeRegistryStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ComputeRegistryStorage *ComputeRegistryStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ComputeRegistryStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ComputeRegistryStorage *ComputeRegistryStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ComputeRegistryStorage.Contract.contract.Transact(opts, method, params...)
}

// ALLOCATIONMANAGER is a free data retrieval call binding the contract method 0x31232bc9.
//
// Solidity: function ALLOCATION_MANAGER() view returns(address)
func (_ComputeRegistryStorage *ComputeRegistryStorageCaller) ALLOCATIONMANAGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ComputeRegistryStorage.contract.Call(opts, &out, "ALLOCATION_MANAGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ALLOCATIONMANAGER is a free data retrieval call binding the contract method 0x31232bc9.
//
// Solidity: function ALLOCATION_MANAGER() view returns(address)
func (_ComputeRegistryStorage *ComputeRegistryStorageSession) ALLOCATIONMANAGER() (common.Address, error) {
	return _ComputeRegistryStorage.Contract.ALLOCATIONMANAGER(&_ComputeRegistryStorage.CallOpts)
}

// ALLOCATIONMANAGER is a free data retrieval call binding the contract method 0x31232bc9.
//
// Solidity: function ALLOCATION_MANAGER() view returns(address)
func (_ComputeRegistryStorage *ComputeRegistryStorageCallerSession) ALLOCATIONMANAGER() (common.Address, error) {
	return _ComputeRegistryStorage.Contract.ALLOCATIONMANAGER(&_ComputeRegistryStorage.CallOpts)
}

// CROSSCHAINREGISTRY is a free data retrieval call binding the contract method 0x9b250844.
//
// Solidity: function CROSS_CHAIN_REGISTRY() view returns(address)
func (_ComputeRegistryStorage *ComputeRegistryStorageCaller) CROSSCHAINREGISTRY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ComputeRegistryStorage.contract.Call(opts, &out, "CROSS_CHAIN_REGISTRY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CROSSCHAINREGISTRY is a free data retrieval call binding the contract method 0x9b250844.
//
// Solidity: function CROSS_CHAIN_REGISTRY() view returns(address)
func (_ComputeRegistryStorage *ComputeRegistryStorageSession) CROSSCHAINREGISTRY() (common.Address, error) {
	return _ComputeRegistryStorage.Contract.CROSSCHAINREGISTRY(&_ComputeRegistryStorage.CallOpts)
}

// CROSSCHAINREGISTRY is a free data retrieval call binding the contract method 0x9b250844.
//
// Solidity: function CROSS_CHAIN_REGISTRY() view returns(address)
func (_ComputeRegistryStorage *ComputeRegistryStorageCallerSession) CROSSCHAINREGISTRY() (common.Address, error) {
	return _ComputeRegistryStorage.Contract.CROSSCHAINREGISTRY(&_ComputeRegistryStorage.CallOpts)
}

// KEYREGISTRAR is a free data retrieval call binding the contract method 0xe6414b48.
//
// Solidity: function KEY_REGISTRAR() view returns(address)
func (_ComputeRegistryStorage *ComputeRegistryStorageCaller) KEYREGISTRAR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ComputeRegistryStorage.contract.Call(opts, &out, "KEY_REGISTRAR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// KEYREGISTRAR is a free data retrieval call binding the contract method 0xe6414b48.
//
// Solidity: function KEY_REGISTRAR() view returns(address)
func (_ComputeRegistryStorage *ComputeRegistryStorageSession) KEYREGISTRAR() (common.Address, error) {
	return _ComputeRegistryStorage.Contract.KEYREGISTRAR(&_ComputeRegistryStorage.CallOpts)
}

// KEYREGISTRAR is a free data retrieval call binding the contract method 0xe6414b48.
//
// Solidity: function KEY_REGISTRAR() view returns(address)
func (_ComputeRegistryStorage *ComputeRegistryStorageCallerSession) KEYREGISTRAR() (common.Address, error) {
	return _ComputeRegistryStorage.Contract.KEYREGISTRAR(&_ComputeRegistryStorage.CallOpts)
}

// MAXEXPIRY is a free data retrieval call binding the contract method 0xb9671690.
//
// Solidity: function MAX_EXPIRY() view returns(uint256)
func (_ComputeRegistryStorage *ComputeRegistryStorageCaller) MAXEXPIRY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ComputeRegistryStorage.contract.Call(opts, &out, "MAX_EXPIRY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXEXPIRY is a free data retrieval call binding the contract method 0xb9671690.
//
// Solidity: function MAX_EXPIRY() view returns(uint256)
func (_ComputeRegistryStorage *ComputeRegistryStorageSession) MAXEXPIRY() (*big.Int, error) {
	return _ComputeRegistryStorage.Contract.MAXEXPIRY(&_ComputeRegistryStorage.CallOpts)
}

// MAXEXPIRY is a free data retrieval call binding the contract method 0xb9671690.
//
// Solidity: function MAX_EXPIRY() view returns(uint256)
func (_ComputeRegistryStorage *ComputeRegistryStorageCallerSession) MAXEXPIRY() (*big.Int, error) {
	return _ComputeRegistryStorage.Contract.MAXEXPIRY(&_ComputeRegistryStorage.CallOpts)
}

// RELEASEMANAGER is a free data retrieval call binding the contract method 0xb39d254f.
//
// Solidity: function RELEASE_MANAGER() view returns(address)
func (_ComputeRegistryStorage *ComputeRegistryStorageCaller) RELEASEMANAGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ComputeRegistryStorage.contract.Call(opts, &out, "RELEASE_MANAGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RELEASEMANAGER is a free data retrieval call binding the contract method 0xb39d254f.
//
// Solidity: function RELEASE_MANAGER() view returns(address)
func (_ComputeRegistryStorage *ComputeRegistryStorageSession) RELEASEMANAGER() (common.Address, error) {
	return _ComputeRegistryStorage.Contract.RELEASEMANAGER(&_ComputeRegistryStorage.CallOpts)
}

// RELEASEMANAGER is a free data retrieval call binding the contract method 0xb39d254f.
//
// Solidity: function RELEASE_MANAGER() view returns(address)
func (_ComputeRegistryStorage *ComputeRegistryStorageCallerSession) RELEASEMANAGER() (common.Address, error) {
	return _ComputeRegistryStorage.Contract.RELEASEMANAGER(&_ComputeRegistryStorage.CallOpts)
}

// TOSAGREEMENTTYPEHASH is a free data retrieval call binding the contract method 0x1de02dbb.
//
// Solidity: function TOS_AGREEMENT_TYPEHASH() view returns(bytes32)
func (_ComputeRegistryStorage *ComputeRegistryStorageCaller) TOSAGREEMENTTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ComputeRegistryStorage.contract.Call(opts, &out, "TOS_AGREEMENT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TOSAGREEMENTTYPEHASH is a free data retrieval call binding the contract method 0x1de02dbb.
//
// Solidity: function TOS_AGREEMENT_TYPEHASH() view returns(bytes32)
func (_ComputeRegistryStorage *ComputeRegistryStorageSession) TOSAGREEMENTTYPEHASH() ([32]byte, error) {
	return _ComputeRegistryStorage.Contract.TOSAGREEMENTTYPEHASH(&_ComputeRegistryStorage.CallOpts)
}

// TOSAGREEMENTTYPEHASH is a free data retrieval call binding the contract method 0x1de02dbb.
//
// Solidity: function TOS_AGREEMENT_TYPEHASH() view returns(bytes32)
func (_ComputeRegistryStorage *ComputeRegistryStorageCallerSession) TOSAGREEMENTTYPEHASH() ([32]byte, error) {
	return _ComputeRegistryStorage.Contract.TOSAGREEMENTTYPEHASH(&_ComputeRegistryStorage.CallOpts)
}

// TOSHASH is a free data retrieval call binding the contract method 0x8df643c7.
//
// Solidity: function TOS_HASH() view returns(bytes32)
func (_ComputeRegistryStorage *ComputeRegistryStorageCaller) TOSHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ComputeRegistryStorage.contract.Call(opts, &out, "TOS_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TOSHASH is a free data retrieval call binding the contract method 0x8df643c7.
//
// Solidity: function TOS_HASH() view returns(bytes32)
func (_ComputeRegistryStorage *ComputeRegistryStorageSession) TOSHASH() ([32]byte, error) {
	return _ComputeRegistryStorage.Contract.TOSHASH(&_ComputeRegistryStorage.CallOpts)
}

// TOSHASH is a free data retrieval call binding the contract method 0x8df643c7.
//
// Solidity: function TOS_HASH() view returns(bytes32)
func (_ComputeRegistryStorage *ComputeRegistryStorageCallerSession) TOSHASH() ([32]byte, error) {
	return _ComputeRegistryStorage.Contract.TOSHASH(&_ComputeRegistryStorage.CallOpts)
}

// CalculateTOSAgreementDigest is a free data retrieval call binding the contract method 0x7a1bb660.
//
// Solidity: function calculateTOSAgreementDigest((address,uint32) operatorSet, address signer) view returns(bytes32)
func (_ComputeRegistryStorage *ComputeRegistryStorageCaller) CalculateTOSAgreementDigest(opts *bind.CallOpts, operatorSet OperatorSet, signer common.Address) ([32]byte, error) {
	var out []interface{}
	err := _ComputeRegistryStorage.contract.Call(opts, &out, "calculateTOSAgreementDigest", operatorSet, signer)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateTOSAgreementDigest is a free data retrieval call binding the contract method 0x7a1bb660.
//
// Solidity: function calculateTOSAgreementDigest((address,uint32) operatorSet, address signer) view returns(bytes32)
func (_ComputeRegistryStorage *ComputeRegistryStorageSession) CalculateTOSAgreementDigest(operatorSet OperatorSet, signer common.Address) ([32]byte, error) {
	return _ComputeRegistryStorage.Contract.CalculateTOSAgreementDigest(&_ComputeRegistryStorage.CallOpts, operatorSet, signer)
}

// CalculateTOSAgreementDigest is a free data retrieval call binding the contract method 0x7a1bb660.
//
// Solidity: function calculateTOSAgreementDigest((address,uint32) operatorSet, address signer) view returns(bytes32)
func (_ComputeRegistryStorage *ComputeRegistryStorageCallerSession) CalculateTOSAgreementDigest(operatorSet OperatorSet, signer common.Address) ([32]byte, error) {
	return _ComputeRegistryStorage.Contract.CalculateTOSAgreementDigest(&_ComputeRegistryStorage.CallOpts, operatorSet, signer)
}

// GetOperatorSetTosSignature is a free data retrieval call binding the contract method 0x130d6165.
//
// Solidity: function getOperatorSetTosSignature((address,uint32) operatorSet) view returns((address,bytes32,bytes))
func (_ComputeRegistryStorage *ComputeRegistryStorageCaller) GetOperatorSetTosSignature(opts *bind.CallOpts, operatorSet OperatorSet) (IComputeRegistryTypesTOSSignature, error) {
	var out []interface{}
	err := _ComputeRegistryStorage.contract.Call(opts, &out, "getOperatorSetTosSignature", operatorSet)

	if err != nil {
		return *new(IComputeRegistryTypesTOSSignature), err
	}

	out0 := *abi.ConvertType(out[0], new(IComputeRegistryTypesTOSSignature)).(*IComputeRegistryTypesTOSSignature)

	return out0, err

}

// GetOperatorSetTosSignature is a free data retrieval call binding the contract method 0x130d6165.
//
// Solidity: function getOperatorSetTosSignature((address,uint32) operatorSet) view returns((address,bytes32,bytes))
func (_ComputeRegistryStorage *ComputeRegistryStorageSession) GetOperatorSetTosSignature(operatorSet OperatorSet) (IComputeRegistryTypesTOSSignature, error) {
	return _ComputeRegistryStorage.Contract.GetOperatorSetTosSignature(&_ComputeRegistryStorage.CallOpts, operatorSet)
}

// GetOperatorSetTosSignature is a free data retrieval call binding the contract method 0x130d6165.
//
// Solidity: function getOperatorSetTosSignature((address,uint32) operatorSet) view returns((address,bytes32,bytes))
func (_ComputeRegistryStorage *ComputeRegistryStorageCallerSession) GetOperatorSetTosSignature(operatorSet OperatorSet) (IComputeRegistryTypesTOSSignature, error) {
	return _ComputeRegistryStorage.Contract.GetOperatorSetTosSignature(&_ComputeRegistryStorage.CallOpts, operatorSet)
}

// IsOperatorSetRegistered is a free data retrieval call binding the contract method 0xc4a1ca05.
//
// Solidity: function isOperatorSetRegistered(bytes32 operatorSetKey) view returns(bool isRegistered)
func (_ComputeRegistryStorage *ComputeRegistryStorageCaller) IsOperatorSetRegistered(opts *bind.CallOpts, operatorSetKey [32]byte) (bool, error) {
	var out []interface{}
	err := _ComputeRegistryStorage.contract.Call(opts, &out, "isOperatorSetRegistered", operatorSetKey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperatorSetRegistered is a free data retrieval call binding the contract method 0xc4a1ca05.
//
// Solidity: function isOperatorSetRegistered(bytes32 operatorSetKey) view returns(bool isRegistered)
func (_ComputeRegistryStorage *ComputeRegistryStorageSession) IsOperatorSetRegistered(operatorSetKey [32]byte) (bool, error) {
	return _ComputeRegistryStorage.Contract.IsOperatorSetRegistered(&_ComputeRegistryStorage.CallOpts, operatorSetKey)
}

// IsOperatorSetRegistered is a free data retrieval call binding the contract method 0xc4a1ca05.
//
// Solidity: function isOperatorSetRegistered(bytes32 operatorSetKey) view returns(bool isRegistered)
func (_ComputeRegistryStorage *ComputeRegistryStorageCallerSession) IsOperatorSetRegistered(operatorSetKey [32]byte) (bool, error) {
	return _ComputeRegistryStorage.Contract.IsOperatorSetRegistered(&_ComputeRegistryStorage.CallOpts, operatorSetKey)
}

// DeregisterFromCompute is a paid mutator transaction binding the contract method 0x89fec15f.
//
// Solidity: function deregisterFromCompute((address,uint32) operatorSet) returns()
func (_ComputeRegistryStorage *ComputeRegistryStorageTransactor) DeregisterFromCompute(opts *bind.TransactOpts, operatorSet OperatorSet) (*types.Transaction, error) {
	return _ComputeRegistryStorage.contract.Transact(opts, "deregisterFromCompute", operatorSet)
}

// DeregisterFromCompute is a paid mutator transaction binding the contract method 0x89fec15f.
//
// Solidity: function deregisterFromCompute((address,uint32) operatorSet) returns()
func (_ComputeRegistryStorage *ComputeRegistryStorageSession) DeregisterFromCompute(operatorSet OperatorSet) (*types.Transaction, error) {
	return _ComputeRegistryStorage.Contract.DeregisterFromCompute(&_ComputeRegistryStorage.TransactOpts, operatorSet)
}

// DeregisterFromCompute is a paid mutator transaction binding the contract method 0x89fec15f.
//
// Solidity: function deregisterFromCompute((address,uint32) operatorSet) returns()
func (_ComputeRegistryStorage *ComputeRegistryStorageTransactorSession) DeregisterFromCompute(operatorSet OperatorSet) (*types.Transaction, error) {
	return _ComputeRegistryStorage.Contract.DeregisterFromCompute(&_ComputeRegistryStorage.TransactOpts, operatorSet)
}

// RegisterForCompute is a paid mutator transaction binding the contract method 0x536b2353.
//
// Solidity: function registerForCompute((address,uint32) operatorSet, bytes signature) returns()
func (_ComputeRegistryStorage *ComputeRegistryStorageTransactor) RegisterForCompute(opts *bind.TransactOpts, operatorSet OperatorSet, signature []byte) (*types.Transaction, error) {
	return _ComputeRegistryStorage.contract.Transact(opts, "registerForCompute", operatorSet, signature)
}

// RegisterForCompute is a paid mutator transaction binding the contract method 0x536b2353.
//
// Solidity: function registerForCompute((address,uint32) operatorSet, bytes signature) returns()
func (_ComputeRegistryStorage *ComputeRegistryStorageSession) RegisterForCompute(operatorSet OperatorSet, signature []byte) (*types.Transaction, error) {
	return _ComputeRegistryStorage.Contract.RegisterForCompute(&_ComputeRegistryStorage.TransactOpts, operatorSet, signature)
}

// RegisterForCompute is a paid mutator transaction binding the contract method 0x536b2353.
//
// Solidity: function registerForCompute((address,uint32) operatorSet, bytes signature) returns()
func (_ComputeRegistryStorage *ComputeRegistryStorageTransactorSession) RegisterForCompute(operatorSet OperatorSet, signature []byte) (*types.Transaction, error) {
	return _ComputeRegistryStorage.Contract.RegisterForCompute(&_ComputeRegistryStorage.TransactOpts, operatorSet, signature)
}

// ComputeRegistryStorageOperatorSetDeregisteredIterator is returned from FilterOperatorSetDeregistered and is used to iterate over the raw logs and unpacked data for OperatorSetDeregistered events raised by the ComputeRegistryStorage contract.
type ComputeRegistryStorageOperatorSetDeregisteredIterator struct {
	Event *ComputeRegistryStorageOperatorSetDeregistered // Event containing the contract specifics and raw log

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
func (it *ComputeRegistryStorageOperatorSetDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComputeRegistryStorageOperatorSetDeregistered)
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
		it.Event = new(ComputeRegistryStorageOperatorSetDeregistered)
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
func (it *ComputeRegistryStorageOperatorSetDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComputeRegistryStorageOperatorSetDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComputeRegistryStorageOperatorSetDeregistered represents a OperatorSetDeregistered event raised by the ComputeRegistryStorage contract.
type ComputeRegistryStorageOperatorSetDeregistered struct {
	OperatorSet OperatorSet
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorSetDeregistered is a free log retrieval operation binding the contract event 0x1d9dd94cbbd6d9d2bc86468a197493d2b960bfc33eb09ec2c9f1731e60c45988.
//
// Solidity: event OperatorSetDeregistered((address,uint32) indexed operatorSet)
func (_ComputeRegistryStorage *ComputeRegistryStorageFilterer) FilterOperatorSetDeregistered(opts *bind.FilterOpts, operatorSet []OperatorSet) (*ComputeRegistryStorageOperatorSetDeregisteredIterator, error) {

	var operatorSetRule []interface{}
	for _, operatorSetItem := range operatorSet {
		operatorSetRule = append(operatorSetRule, operatorSetItem)
	}

	logs, sub, err := _ComputeRegistryStorage.contract.FilterLogs(opts, "OperatorSetDeregistered", operatorSetRule)
	if err != nil {
		return nil, err
	}
	return &ComputeRegistryStorageOperatorSetDeregisteredIterator{contract: _ComputeRegistryStorage.contract, event: "OperatorSetDeregistered", logs: logs, sub: sub}, nil
}

// WatchOperatorSetDeregistered is a free log subscription operation binding the contract event 0x1d9dd94cbbd6d9d2bc86468a197493d2b960bfc33eb09ec2c9f1731e60c45988.
//
// Solidity: event OperatorSetDeregistered((address,uint32) indexed operatorSet)
func (_ComputeRegistryStorage *ComputeRegistryStorageFilterer) WatchOperatorSetDeregistered(opts *bind.WatchOpts, sink chan<- *ComputeRegistryStorageOperatorSetDeregistered, operatorSet []OperatorSet) (event.Subscription, error) {

	var operatorSetRule []interface{}
	for _, operatorSetItem := range operatorSet {
		operatorSetRule = append(operatorSetRule, operatorSetItem)
	}

	logs, sub, err := _ComputeRegistryStorage.contract.WatchLogs(opts, "OperatorSetDeregistered", operatorSetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComputeRegistryStorageOperatorSetDeregistered)
				if err := _ComputeRegistryStorage.contract.UnpackLog(event, "OperatorSetDeregistered", log); err != nil {
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
func (_ComputeRegistryStorage *ComputeRegistryStorageFilterer) ParseOperatorSetDeregistered(log types.Log) (*ComputeRegistryStorageOperatorSetDeregistered, error) {
	event := new(ComputeRegistryStorageOperatorSetDeregistered)
	if err := _ComputeRegistryStorage.contract.UnpackLog(event, "OperatorSetDeregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComputeRegistryStorageOperatorSetRegisteredIterator is returned from FilterOperatorSetRegistered and is used to iterate over the raw logs and unpacked data for OperatorSetRegistered events raised by the ComputeRegistryStorage contract.
type ComputeRegistryStorageOperatorSetRegisteredIterator struct {
	Event *ComputeRegistryStorageOperatorSetRegistered // Event containing the contract specifics and raw log

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
func (it *ComputeRegistryStorageOperatorSetRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComputeRegistryStorageOperatorSetRegistered)
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
		it.Event = new(ComputeRegistryStorageOperatorSetRegistered)
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
func (it *ComputeRegistryStorageOperatorSetRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComputeRegistryStorageOperatorSetRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComputeRegistryStorageOperatorSetRegistered represents a OperatorSetRegistered event raised by the ComputeRegistryStorage contract.
type ComputeRegistryStorageOperatorSetRegistered struct {
	OperatorSet OperatorSet
	Signer      common.Address
	TosHash     [32]byte
	Signature   []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorSetRegistered is a free log retrieval operation binding the contract event 0xab9d9ee579ad2bc5e08fba9e2b3d59225f1400f1a5830745283697e35ce2dcc8.
//
// Solidity: event OperatorSetRegistered((address,uint32) indexed operatorSet, address indexed signer, bytes32 indexed tosHash, bytes signature)
func (_ComputeRegistryStorage *ComputeRegistryStorageFilterer) FilterOperatorSetRegistered(opts *bind.FilterOpts, operatorSet []OperatorSet, signer []common.Address, tosHash [][32]byte) (*ComputeRegistryStorageOperatorSetRegisteredIterator, error) {

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

	logs, sub, err := _ComputeRegistryStorage.contract.FilterLogs(opts, "OperatorSetRegistered", operatorSetRule, signerRule, tosHashRule)
	if err != nil {
		return nil, err
	}
	return &ComputeRegistryStorageOperatorSetRegisteredIterator{contract: _ComputeRegistryStorage.contract, event: "OperatorSetRegistered", logs: logs, sub: sub}, nil
}

// WatchOperatorSetRegistered is a free log subscription operation binding the contract event 0xab9d9ee579ad2bc5e08fba9e2b3d59225f1400f1a5830745283697e35ce2dcc8.
//
// Solidity: event OperatorSetRegistered((address,uint32) indexed operatorSet, address indexed signer, bytes32 indexed tosHash, bytes signature)
func (_ComputeRegistryStorage *ComputeRegistryStorageFilterer) WatchOperatorSetRegistered(opts *bind.WatchOpts, sink chan<- *ComputeRegistryStorageOperatorSetRegistered, operatorSet []OperatorSet, signer []common.Address, tosHash [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _ComputeRegistryStorage.contract.WatchLogs(opts, "OperatorSetRegistered", operatorSetRule, signerRule, tosHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComputeRegistryStorageOperatorSetRegistered)
				if err := _ComputeRegistryStorage.contract.UnpackLog(event, "OperatorSetRegistered", log); err != nil {
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
func (_ComputeRegistryStorage *ComputeRegistryStorageFilterer) ParseOperatorSetRegistered(log types.Log) (*ComputeRegistryStorageOperatorSetRegistered, error) {
	event := new(ComputeRegistryStorageOperatorSetRegistered)
	if err := _ComputeRegistryStorage.contract.UnpackLog(event, "OperatorSetRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
