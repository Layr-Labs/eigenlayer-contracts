// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package AVSDirectoryStorage

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

// ISignatureUtilsSignatureWithSaltAndExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithSaltAndExpiry struct {
	Signature []byte
	Salt      [32]byte
	Expiry    *big.Int
}

// AVSDirectoryStorageMetaData contains all meta data concerning the AVSDirectoryStorage contract.
var AVSDirectoryStorageMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"DOMAIN_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"OPERATOR_AVS_REGISTRATION_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"OPERATOR_SET_REGISTRATION_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateOperatorAVSRegistrationDigestHash\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateOperatorSetRegistrationDigestHash\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregisterOperatorFromAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deregisterOperatorFromOperatorSets\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isMember\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"memberInfo\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"inTotalSets\",\"type\":\"uint248\",\"internalType\":\"uint248\"},{\"name\":\"isLegacyOperator\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorSaltIsSpent\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerOperatorToAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerOperatorToOperatorSets\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"signature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAVSMetadataURI\",\"inputs\":[{\"name\":\"metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AVSMetadataURIUpdated\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"metadataURI\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorAVSRegistrationStatusUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"isLegacyOperator\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorAddedToOperatorSet\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRemovedFromOperatorSet\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false}]",
}

// AVSDirectoryStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use AVSDirectoryStorageMetaData.ABI instead.
var AVSDirectoryStorageABI = AVSDirectoryStorageMetaData.ABI

// AVSDirectoryStorage is an auto generated Go binding around an Ethereum contract.
type AVSDirectoryStorage struct {
	AVSDirectoryStorageCaller     // Read-only binding to the contract
	AVSDirectoryStorageTransactor // Write-only binding to the contract
	AVSDirectoryStorageFilterer   // Log filterer for contract events
}

// AVSDirectoryStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type AVSDirectoryStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AVSDirectoryStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AVSDirectoryStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AVSDirectoryStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AVSDirectoryStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AVSDirectoryStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AVSDirectoryStorageSession struct {
	Contract     *AVSDirectoryStorage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// AVSDirectoryStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AVSDirectoryStorageCallerSession struct {
	Contract *AVSDirectoryStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// AVSDirectoryStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AVSDirectoryStorageTransactorSession struct {
	Contract     *AVSDirectoryStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// AVSDirectoryStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type AVSDirectoryStorageRaw struct {
	Contract *AVSDirectoryStorage // Generic contract binding to access the raw methods on
}

// AVSDirectoryStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AVSDirectoryStorageCallerRaw struct {
	Contract *AVSDirectoryStorageCaller // Generic read-only contract binding to access the raw methods on
}

// AVSDirectoryStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AVSDirectoryStorageTransactorRaw struct {
	Contract *AVSDirectoryStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAVSDirectoryStorage creates a new instance of AVSDirectoryStorage, bound to a specific deployed contract.
func NewAVSDirectoryStorage(address common.Address, backend bind.ContractBackend) (*AVSDirectoryStorage, error) {
	contract, err := bindAVSDirectoryStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AVSDirectoryStorage{AVSDirectoryStorageCaller: AVSDirectoryStorageCaller{contract: contract}, AVSDirectoryStorageTransactor: AVSDirectoryStorageTransactor{contract: contract}, AVSDirectoryStorageFilterer: AVSDirectoryStorageFilterer{contract: contract}}, nil
}

// NewAVSDirectoryStorageCaller creates a new read-only instance of AVSDirectoryStorage, bound to a specific deployed contract.
func NewAVSDirectoryStorageCaller(address common.Address, caller bind.ContractCaller) (*AVSDirectoryStorageCaller, error) {
	contract, err := bindAVSDirectoryStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AVSDirectoryStorageCaller{contract: contract}, nil
}

// NewAVSDirectoryStorageTransactor creates a new write-only instance of AVSDirectoryStorage, bound to a specific deployed contract.
func NewAVSDirectoryStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*AVSDirectoryStorageTransactor, error) {
	contract, err := bindAVSDirectoryStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AVSDirectoryStorageTransactor{contract: contract}, nil
}

// NewAVSDirectoryStorageFilterer creates a new log filterer instance of AVSDirectoryStorage, bound to a specific deployed contract.
func NewAVSDirectoryStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*AVSDirectoryStorageFilterer, error) {
	contract, err := bindAVSDirectoryStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AVSDirectoryStorageFilterer{contract: contract}, nil
}

// bindAVSDirectoryStorage binds a generic wrapper to an already deployed contract.
func bindAVSDirectoryStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AVSDirectoryStorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AVSDirectoryStorage *AVSDirectoryStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AVSDirectoryStorage.Contract.AVSDirectoryStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AVSDirectoryStorage *AVSDirectoryStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AVSDirectoryStorage.Contract.AVSDirectoryStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AVSDirectoryStorage *AVSDirectoryStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AVSDirectoryStorage.Contract.AVSDirectoryStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AVSDirectoryStorage *AVSDirectoryStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AVSDirectoryStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AVSDirectoryStorage *AVSDirectoryStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AVSDirectoryStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AVSDirectoryStorage *AVSDirectoryStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AVSDirectoryStorage.Contract.contract.Transact(opts, method, params...)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_AVSDirectoryStorage *AVSDirectoryStorageCaller) DOMAINTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AVSDirectoryStorage.contract.Call(opts, &out, "DOMAIN_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_AVSDirectoryStorage *AVSDirectoryStorageSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _AVSDirectoryStorage.Contract.DOMAINTYPEHASH(&_AVSDirectoryStorage.CallOpts)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_AVSDirectoryStorage *AVSDirectoryStorageCallerSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _AVSDirectoryStorage.Contract.DOMAINTYPEHASH(&_AVSDirectoryStorage.CallOpts)
}

// OPERATORAVSREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0xd79aceab.
//
// Solidity: function OPERATOR_AVS_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_AVSDirectoryStorage *AVSDirectoryStorageCaller) OPERATORAVSREGISTRATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AVSDirectoryStorage.contract.Call(opts, &out, "OPERATOR_AVS_REGISTRATION_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORAVSREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0xd79aceab.
//
// Solidity: function OPERATOR_AVS_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_AVSDirectoryStorage *AVSDirectoryStorageSession) OPERATORAVSREGISTRATIONTYPEHASH() ([32]byte, error) {
	return _AVSDirectoryStorage.Contract.OPERATORAVSREGISTRATIONTYPEHASH(&_AVSDirectoryStorage.CallOpts)
}

// OPERATORAVSREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0xd79aceab.
//
// Solidity: function OPERATOR_AVS_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_AVSDirectoryStorage *AVSDirectoryStorageCallerSession) OPERATORAVSREGISTRATIONTYPEHASH() ([32]byte, error) {
	return _AVSDirectoryStorage.Contract.OPERATORAVSREGISTRATIONTYPEHASH(&_AVSDirectoryStorage.CallOpts)
}

// OPERATORSETREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0xc825fe68.
//
// Solidity: function OPERATOR_SET_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_AVSDirectoryStorage *AVSDirectoryStorageCaller) OPERATORSETREGISTRATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AVSDirectoryStorage.contract.Call(opts, &out, "OPERATOR_SET_REGISTRATION_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORSETREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0xc825fe68.
//
// Solidity: function OPERATOR_SET_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_AVSDirectoryStorage *AVSDirectoryStorageSession) OPERATORSETREGISTRATIONTYPEHASH() ([32]byte, error) {
	return _AVSDirectoryStorage.Contract.OPERATORSETREGISTRATIONTYPEHASH(&_AVSDirectoryStorage.CallOpts)
}

// OPERATORSETREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0xc825fe68.
//
// Solidity: function OPERATOR_SET_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_AVSDirectoryStorage *AVSDirectoryStorageCallerSession) OPERATORSETREGISTRATIONTYPEHASH() ([32]byte, error) {
	return _AVSDirectoryStorage.Contract.OPERATORSETREGISTRATIONTYPEHASH(&_AVSDirectoryStorage.CallOpts)
}

// CalculateOperatorAVSRegistrationDigestHash is a free data retrieval call binding the contract method 0xa1060c88.
//
// Solidity: function calculateOperatorAVSRegistrationDigestHash(address operator, address avs, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_AVSDirectoryStorage *AVSDirectoryStorageCaller) CalculateOperatorAVSRegistrationDigestHash(opts *bind.CallOpts, operator common.Address, avs common.Address, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _AVSDirectoryStorage.contract.Call(opts, &out, "calculateOperatorAVSRegistrationDigestHash", operator, avs, salt, expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateOperatorAVSRegistrationDigestHash is a free data retrieval call binding the contract method 0xa1060c88.
//
// Solidity: function calculateOperatorAVSRegistrationDigestHash(address operator, address avs, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_AVSDirectoryStorage *AVSDirectoryStorageSession) CalculateOperatorAVSRegistrationDigestHash(operator common.Address, avs common.Address, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _AVSDirectoryStorage.Contract.CalculateOperatorAVSRegistrationDigestHash(&_AVSDirectoryStorage.CallOpts, operator, avs, salt, expiry)
}

// CalculateOperatorAVSRegistrationDigestHash is a free data retrieval call binding the contract method 0xa1060c88.
//
// Solidity: function calculateOperatorAVSRegistrationDigestHash(address operator, address avs, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_AVSDirectoryStorage *AVSDirectoryStorageCallerSession) CalculateOperatorAVSRegistrationDigestHash(operator common.Address, avs common.Address, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _AVSDirectoryStorage.Contract.CalculateOperatorAVSRegistrationDigestHash(&_AVSDirectoryStorage.CallOpts, operator, avs, salt, expiry)
}

// CalculateOperatorSetRegistrationDigestHash is a free data retrieval call binding the contract method 0x955e6696.
//
// Solidity: function calculateOperatorSetRegistrationDigestHash(address avs, uint32[] operatorSetIds, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_AVSDirectoryStorage *AVSDirectoryStorageCaller) CalculateOperatorSetRegistrationDigestHash(opts *bind.CallOpts, avs common.Address, operatorSetIds []uint32, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _AVSDirectoryStorage.contract.Call(opts, &out, "calculateOperatorSetRegistrationDigestHash", avs, operatorSetIds, salt, expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateOperatorSetRegistrationDigestHash is a free data retrieval call binding the contract method 0x955e6696.
//
// Solidity: function calculateOperatorSetRegistrationDigestHash(address avs, uint32[] operatorSetIds, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_AVSDirectoryStorage *AVSDirectoryStorageSession) CalculateOperatorSetRegistrationDigestHash(avs common.Address, operatorSetIds []uint32, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _AVSDirectoryStorage.Contract.CalculateOperatorSetRegistrationDigestHash(&_AVSDirectoryStorage.CallOpts, avs, operatorSetIds, salt, expiry)
}

// CalculateOperatorSetRegistrationDigestHash is a free data retrieval call binding the contract method 0x955e6696.
//
// Solidity: function calculateOperatorSetRegistrationDigestHash(address avs, uint32[] operatorSetIds, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_AVSDirectoryStorage *AVSDirectoryStorageCallerSession) CalculateOperatorSetRegistrationDigestHash(avs common.Address, operatorSetIds []uint32, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _AVSDirectoryStorage.Contract.CalculateOperatorSetRegistrationDigestHash(&_AVSDirectoryStorage.CallOpts, avs, operatorSetIds, salt, expiry)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_AVSDirectoryStorage *AVSDirectoryStorageCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AVSDirectoryStorage.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_AVSDirectoryStorage *AVSDirectoryStorageSession) Delegation() (common.Address, error) {
	return _AVSDirectoryStorage.Contract.Delegation(&_AVSDirectoryStorage.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_AVSDirectoryStorage *AVSDirectoryStorageCallerSession) Delegation() (common.Address, error) {
	return _AVSDirectoryStorage.Contract.Delegation(&_AVSDirectoryStorage.CallOpts)
}

// IsMember is a free data retrieval call binding the contract method 0x3c4385d0.
//
// Solidity: function isMember(address , address , uint32 ) view returns(bool)
func (_AVSDirectoryStorage *AVSDirectoryStorageCaller) IsMember(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 uint32) (bool, error) {
	var out []interface{}
	err := _AVSDirectoryStorage.contract.Call(opts, &out, "isMember", arg0, arg1, arg2)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMember is a free data retrieval call binding the contract method 0x3c4385d0.
//
// Solidity: function isMember(address , address , uint32 ) view returns(bool)
func (_AVSDirectoryStorage *AVSDirectoryStorageSession) IsMember(arg0 common.Address, arg1 common.Address, arg2 uint32) (bool, error) {
	return _AVSDirectoryStorage.Contract.IsMember(&_AVSDirectoryStorage.CallOpts, arg0, arg1, arg2)
}

// IsMember is a free data retrieval call binding the contract method 0x3c4385d0.
//
// Solidity: function isMember(address , address , uint32 ) view returns(bool)
func (_AVSDirectoryStorage *AVSDirectoryStorageCallerSession) IsMember(arg0 common.Address, arg1 common.Address, arg2 uint32) (bool, error) {
	return _AVSDirectoryStorage.Contract.IsMember(&_AVSDirectoryStorage.CallOpts, arg0, arg1, arg2)
}

// MemberInfo is a free data retrieval call binding the contract method 0xf37c521c.
//
// Solidity: function memberInfo(address , address ) view returns(uint248 inTotalSets, bool isLegacyOperator)
func (_AVSDirectoryStorage *AVSDirectoryStorageCaller) MemberInfo(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (struct {
	InTotalSets      *big.Int
	IsLegacyOperator bool
}, error) {
	var out []interface{}
	err := _AVSDirectoryStorage.contract.Call(opts, &out, "memberInfo", arg0, arg1)

	outstruct := new(struct {
		InTotalSets      *big.Int
		IsLegacyOperator bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.InTotalSets = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.IsLegacyOperator = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// MemberInfo is a free data retrieval call binding the contract method 0xf37c521c.
//
// Solidity: function memberInfo(address , address ) view returns(uint248 inTotalSets, bool isLegacyOperator)
func (_AVSDirectoryStorage *AVSDirectoryStorageSession) MemberInfo(arg0 common.Address, arg1 common.Address) (struct {
	InTotalSets      *big.Int
	IsLegacyOperator bool
}, error) {
	return _AVSDirectoryStorage.Contract.MemberInfo(&_AVSDirectoryStorage.CallOpts, arg0, arg1)
}

// MemberInfo is a free data retrieval call binding the contract method 0xf37c521c.
//
// Solidity: function memberInfo(address , address ) view returns(uint248 inTotalSets, bool isLegacyOperator)
func (_AVSDirectoryStorage *AVSDirectoryStorageCallerSession) MemberInfo(arg0 common.Address, arg1 common.Address) (struct {
	InTotalSets      *big.Int
	IsLegacyOperator bool
}, error) {
	return _AVSDirectoryStorage.Contract.MemberInfo(&_AVSDirectoryStorage.CallOpts, arg0, arg1)
}

// OperatorSaltIsSpent is a free data retrieval call binding the contract method 0x374823b5.
//
// Solidity: function operatorSaltIsSpent(address , bytes32 ) view returns(bool)
func (_AVSDirectoryStorage *AVSDirectoryStorageCaller) OperatorSaltIsSpent(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (bool, error) {
	var out []interface{}
	err := _AVSDirectoryStorage.contract.Call(opts, &out, "operatorSaltIsSpent", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OperatorSaltIsSpent is a free data retrieval call binding the contract method 0x374823b5.
//
// Solidity: function operatorSaltIsSpent(address , bytes32 ) view returns(bool)
func (_AVSDirectoryStorage *AVSDirectoryStorageSession) OperatorSaltIsSpent(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _AVSDirectoryStorage.Contract.OperatorSaltIsSpent(&_AVSDirectoryStorage.CallOpts, arg0, arg1)
}

// OperatorSaltIsSpent is a free data retrieval call binding the contract method 0x374823b5.
//
// Solidity: function operatorSaltIsSpent(address , bytes32 ) view returns(bool)
func (_AVSDirectoryStorage *AVSDirectoryStorageCallerSession) OperatorSaltIsSpent(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _AVSDirectoryStorage.Contract.OperatorSaltIsSpent(&_AVSDirectoryStorage.CallOpts, arg0, arg1)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_AVSDirectoryStorage *AVSDirectoryStorageTransactor) DeregisterOperatorFromAVS(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _AVSDirectoryStorage.contract.Transact(opts, "deregisterOperatorFromAVS", operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_AVSDirectoryStorage *AVSDirectoryStorageSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _AVSDirectoryStorage.Contract.DeregisterOperatorFromAVS(&_AVSDirectoryStorage.TransactOpts, operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_AVSDirectoryStorage *AVSDirectoryStorageTransactorSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _AVSDirectoryStorage.Contract.DeregisterOperatorFromAVS(&_AVSDirectoryStorage.TransactOpts, operator)
}

// DeregisterOperatorFromOperatorSets is a paid mutator transaction binding the contract method 0xc1a8e2c5.
//
// Solidity: function deregisterOperatorFromOperatorSets(address operator, uint32[] operatorSetIds) returns()
func (_AVSDirectoryStorage *AVSDirectoryStorageTransactor) DeregisterOperatorFromOperatorSets(opts *bind.TransactOpts, operator common.Address, operatorSetIds []uint32) (*types.Transaction, error) {
	return _AVSDirectoryStorage.contract.Transact(opts, "deregisterOperatorFromOperatorSets", operator, operatorSetIds)
}

// DeregisterOperatorFromOperatorSets is a paid mutator transaction binding the contract method 0xc1a8e2c5.
//
// Solidity: function deregisterOperatorFromOperatorSets(address operator, uint32[] operatorSetIds) returns()
func (_AVSDirectoryStorage *AVSDirectoryStorageSession) DeregisterOperatorFromOperatorSets(operator common.Address, operatorSetIds []uint32) (*types.Transaction, error) {
	return _AVSDirectoryStorage.Contract.DeregisterOperatorFromOperatorSets(&_AVSDirectoryStorage.TransactOpts, operator, operatorSetIds)
}

// DeregisterOperatorFromOperatorSets is a paid mutator transaction binding the contract method 0xc1a8e2c5.
//
// Solidity: function deregisterOperatorFromOperatorSets(address operator, uint32[] operatorSetIds) returns()
func (_AVSDirectoryStorage *AVSDirectoryStorageTransactorSession) DeregisterOperatorFromOperatorSets(operator common.Address, operatorSetIds []uint32) (*types.Transaction, error) {
	return _AVSDirectoryStorage.Contract.DeregisterOperatorFromOperatorSets(&_AVSDirectoryStorage.TransactOpts, operator, operatorSetIds)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_AVSDirectoryStorage *AVSDirectoryStorageTransactor) RegisterOperatorToAVS(opts *bind.TransactOpts, operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _AVSDirectoryStorage.contract.Transact(opts, "registerOperatorToAVS", operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_AVSDirectoryStorage *AVSDirectoryStorageSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _AVSDirectoryStorage.Contract.RegisterOperatorToAVS(&_AVSDirectoryStorage.TransactOpts, operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_AVSDirectoryStorage *AVSDirectoryStorageTransactorSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _AVSDirectoryStorage.Contract.RegisterOperatorToAVS(&_AVSDirectoryStorage.TransactOpts, operator, operatorSignature)
}

// RegisterOperatorToOperatorSets is a paid mutator transaction binding the contract method 0x1e2199e2.
//
// Solidity: function registerOperatorToOperatorSets(address operator, uint32[] operatorSetIds, (bytes,bytes32,uint256) signature) returns()
func (_AVSDirectoryStorage *AVSDirectoryStorageTransactor) RegisterOperatorToOperatorSets(opts *bind.TransactOpts, operator common.Address, operatorSetIds []uint32, signature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _AVSDirectoryStorage.contract.Transact(opts, "registerOperatorToOperatorSets", operator, operatorSetIds, signature)
}

// RegisterOperatorToOperatorSets is a paid mutator transaction binding the contract method 0x1e2199e2.
//
// Solidity: function registerOperatorToOperatorSets(address operator, uint32[] operatorSetIds, (bytes,bytes32,uint256) signature) returns()
func (_AVSDirectoryStorage *AVSDirectoryStorageSession) RegisterOperatorToOperatorSets(operator common.Address, operatorSetIds []uint32, signature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _AVSDirectoryStorage.Contract.RegisterOperatorToOperatorSets(&_AVSDirectoryStorage.TransactOpts, operator, operatorSetIds, signature)
}

// RegisterOperatorToOperatorSets is a paid mutator transaction binding the contract method 0x1e2199e2.
//
// Solidity: function registerOperatorToOperatorSets(address operator, uint32[] operatorSetIds, (bytes,bytes32,uint256) signature) returns()
func (_AVSDirectoryStorage *AVSDirectoryStorageTransactorSession) RegisterOperatorToOperatorSets(operator common.Address, operatorSetIds []uint32, signature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _AVSDirectoryStorage.Contract.RegisterOperatorToOperatorSets(&_AVSDirectoryStorage.TransactOpts, operator, operatorSetIds, signature)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string metadataURI) returns()
func (_AVSDirectoryStorage *AVSDirectoryStorageTransactor) UpdateAVSMetadataURI(opts *bind.TransactOpts, metadataURI string) (*types.Transaction, error) {
	return _AVSDirectoryStorage.contract.Transact(opts, "updateAVSMetadataURI", metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string metadataURI) returns()
func (_AVSDirectoryStorage *AVSDirectoryStorageSession) UpdateAVSMetadataURI(metadataURI string) (*types.Transaction, error) {
	return _AVSDirectoryStorage.Contract.UpdateAVSMetadataURI(&_AVSDirectoryStorage.TransactOpts, metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string metadataURI) returns()
func (_AVSDirectoryStorage *AVSDirectoryStorageTransactorSession) UpdateAVSMetadataURI(metadataURI string) (*types.Transaction, error) {
	return _AVSDirectoryStorage.Contract.UpdateAVSMetadataURI(&_AVSDirectoryStorage.TransactOpts, metadataURI)
}

// AVSDirectoryStorageAVSMetadataURIUpdatedIterator is returned from FilterAVSMetadataURIUpdated and is used to iterate over the raw logs and unpacked data for AVSMetadataURIUpdated events raised by the AVSDirectoryStorage contract.
type AVSDirectoryStorageAVSMetadataURIUpdatedIterator struct {
	Event *AVSDirectoryStorageAVSMetadataURIUpdated // Event containing the contract specifics and raw log

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
func (it *AVSDirectoryStorageAVSMetadataURIUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSDirectoryStorageAVSMetadataURIUpdated)
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
		it.Event = new(AVSDirectoryStorageAVSMetadataURIUpdated)
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
func (it *AVSDirectoryStorageAVSMetadataURIUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSDirectoryStorageAVSMetadataURIUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSDirectoryStorageAVSMetadataURIUpdated represents a AVSMetadataURIUpdated event raised by the AVSDirectoryStorage contract.
type AVSDirectoryStorageAVSMetadataURIUpdated struct {
	Avs         common.Address
	MetadataURI string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAVSMetadataURIUpdated is a free log retrieval operation binding the contract event 0xa89c1dc243d8908a96dd84944bcc97d6bc6ac00dd78e20621576be6a3c943713.
//
// Solidity: event AVSMetadataURIUpdated(address indexed avs, string metadataURI)
func (_AVSDirectoryStorage *AVSDirectoryStorageFilterer) FilterAVSMetadataURIUpdated(opts *bind.FilterOpts, avs []common.Address) (*AVSDirectoryStorageAVSMetadataURIUpdatedIterator, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _AVSDirectoryStorage.contract.FilterLogs(opts, "AVSMetadataURIUpdated", avsRule)
	if err != nil {
		return nil, err
	}
	return &AVSDirectoryStorageAVSMetadataURIUpdatedIterator{contract: _AVSDirectoryStorage.contract, event: "AVSMetadataURIUpdated", logs: logs, sub: sub}, nil
}

// WatchAVSMetadataURIUpdated is a free log subscription operation binding the contract event 0xa89c1dc243d8908a96dd84944bcc97d6bc6ac00dd78e20621576be6a3c943713.
//
// Solidity: event AVSMetadataURIUpdated(address indexed avs, string metadataURI)
func (_AVSDirectoryStorage *AVSDirectoryStorageFilterer) WatchAVSMetadataURIUpdated(opts *bind.WatchOpts, sink chan<- *AVSDirectoryStorageAVSMetadataURIUpdated, avs []common.Address) (event.Subscription, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _AVSDirectoryStorage.contract.WatchLogs(opts, "AVSMetadataURIUpdated", avsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSDirectoryStorageAVSMetadataURIUpdated)
				if err := _AVSDirectoryStorage.contract.UnpackLog(event, "AVSMetadataURIUpdated", log); err != nil {
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

// ParseAVSMetadataURIUpdated is a log parse operation binding the contract event 0xa89c1dc243d8908a96dd84944bcc97d6bc6ac00dd78e20621576be6a3c943713.
//
// Solidity: event AVSMetadataURIUpdated(address indexed avs, string metadataURI)
func (_AVSDirectoryStorage *AVSDirectoryStorageFilterer) ParseAVSMetadataURIUpdated(log types.Log) (*AVSDirectoryStorageAVSMetadataURIUpdated, error) {
	event := new(AVSDirectoryStorageAVSMetadataURIUpdated)
	if err := _AVSDirectoryStorage.contract.UnpackLog(event, "AVSMetadataURIUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSDirectoryStorageOperatorAVSRegistrationStatusUpdatedIterator is returned from FilterOperatorAVSRegistrationStatusUpdated and is used to iterate over the raw logs and unpacked data for OperatorAVSRegistrationStatusUpdated events raised by the AVSDirectoryStorage contract.
type AVSDirectoryStorageOperatorAVSRegistrationStatusUpdatedIterator struct {
	Event *AVSDirectoryStorageOperatorAVSRegistrationStatusUpdated // Event containing the contract specifics and raw log

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
func (it *AVSDirectoryStorageOperatorAVSRegistrationStatusUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSDirectoryStorageOperatorAVSRegistrationStatusUpdated)
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
		it.Event = new(AVSDirectoryStorageOperatorAVSRegistrationStatusUpdated)
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
func (it *AVSDirectoryStorageOperatorAVSRegistrationStatusUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSDirectoryStorageOperatorAVSRegistrationStatusUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSDirectoryStorageOperatorAVSRegistrationStatusUpdated represents a OperatorAVSRegistrationStatusUpdated event raised by the AVSDirectoryStorage contract.
type AVSDirectoryStorageOperatorAVSRegistrationStatusUpdated struct {
	Operator         common.Address
	Avs              common.Address
	IsLegacyOperator bool
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOperatorAVSRegistrationStatusUpdated is a free log retrieval operation binding the contract event 0x1864cf493d2c3314abbcef95e20acb37a637bccf3c51f4fe2f6f9492fbeabd7f.
//
// Solidity: event OperatorAVSRegistrationStatusUpdated(address indexed operator, address indexed avs, bool isLegacyOperator)
func (_AVSDirectoryStorage *AVSDirectoryStorageFilterer) FilterOperatorAVSRegistrationStatusUpdated(opts *bind.FilterOpts, operator []common.Address, avs []common.Address) (*AVSDirectoryStorageOperatorAVSRegistrationStatusUpdatedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _AVSDirectoryStorage.contract.FilterLogs(opts, "OperatorAVSRegistrationStatusUpdated", operatorRule, avsRule)
	if err != nil {
		return nil, err
	}
	return &AVSDirectoryStorageOperatorAVSRegistrationStatusUpdatedIterator{contract: _AVSDirectoryStorage.contract, event: "OperatorAVSRegistrationStatusUpdated", logs: logs, sub: sub}, nil
}

// WatchOperatorAVSRegistrationStatusUpdated is a free log subscription operation binding the contract event 0x1864cf493d2c3314abbcef95e20acb37a637bccf3c51f4fe2f6f9492fbeabd7f.
//
// Solidity: event OperatorAVSRegistrationStatusUpdated(address indexed operator, address indexed avs, bool isLegacyOperator)
func (_AVSDirectoryStorage *AVSDirectoryStorageFilterer) WatchOperatorAVSRegistrationStatusUpdated(opts *bind.WatchOpts, sink chan<- *AVSDirectoryStorageOperatorAVSRegistrationStatusUpdated, operator []common.Address, avs []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _AVSDirectoryStorage.contract.WatchLogs(opts, "OperatorAVSRegistrationStatusUpdated", operatorRule, avsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSDirectoryStorageOperatorAVSRegistrationStatusUpdated)
				if err := _AVSDirectoryStorage.contract.UnpackLog(event, "OperatorAVSRegistrationStatusUpdated", log); err != nil {
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

// ParseOperatorAVSRegistrationStatusUpdated is a log parse operation binding the contract event 0x1864cf493d2c3314abbcef95e20acb37a637bccf3c51f4fe2f6f9492fbeabd7f.
//
// Solidity: event OperatorAVSRegistrationStatusUpdated(address indexed operator, address indexed avs, bool isLegacyOperator)
func (_AVSDirectoryStorage *AVSDirectoryStorageFilterer) ParseOperatorAVSRegistrationStatusUpdated(log types.Log) (*AVSDirectoryStorageOperatorAVSRegistrationStatusUpdated, error) {
	event := new(AVSDirectoryStorageOperatorAVSRegistrationStatusUpdated)
	if err := _AVSDirectoryStorage.contract.UnpackLog(event, "OperatorAVSRegistrationStatusUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSDirectoryStorageOperatorAddedToOperatorSetIterator is returned from FilterOperatorAddedToOperatorSet and is used to iterate over the raw logs and unpacked data for OperatorAddedToOperatorSet events raised by the AVSDirectoryStorage contract.
type AVSDirectoryStorageOperatorAddedToOperatorSetIterator struct {
	Event *AVSDirectoryStorageOperatorAddedToOperatorSet // Event containing the contract specifics and raw log

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
func (it *AVSDirectoryStorageOperatorAddedToOperatorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSDirectoryStorageOperatorAddedToOperatorSet)
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
		it.Event = new(AVSDirectoryStorageOperatorAddedToOperatorSet)
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
func (it *AVSDirectoryStorageOperatorAddedToOperatorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSDirectoryStorageOperatorAddedToOperatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSDirectoryStorageOperatorAddedToOperatorSet represents a OperatorAddedToOperatorSet event raised by the AVSDirectoryStorage contract.
type AVSDirectoryStorageOperatorAddedToOperatorSet struct {
	Operator      common.Address
	Avs           common.Address
	OperatorSetId uint32
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOperatorAddedToOperatorSet is a free log retrieval operation binding the contract event 0x0e2d64f3a0f18391c95a0a816961e081b15096f8bb957d0b654897b96b7ba309.
//
// Solidity: event OperatorAddedToOperatorSet(address operator, address avs, uint32 operatorSetId)
func (_AVSDirectoryStorage *AVSDirectoryStorageFilterer) FilterOperatorAddedToOperatorSet(opts *bind.FilterOpts) (*AVSDirectoryStorageOperatorAddedToOperatorSetIterator, error) {

	logs, sub, err := _AVSDirectoryStorage.contract.FilterLogs(opts, "OperatorAddedToOperatorSet")
	if err != nil {
		return nil, err
	}
	return &AVSDirectoryStorageOperatorAddedToOperatorSetIterator{contract: _AVSDirectoryStorage.contract, event: "OperatorAddedToOperatorSet", logs: logs, sub: sub}, nil
}

// WatchOperatorAddedToOperatorSet is a free log subscription operation binding the contract event 0x0e2d64f3a0f18391c95a0a816961e081b15096f8bb957d0b654897b96b7ba309.
//
// Solidity: event OperatorAddedToOperatorSet(address operator, address avs, uint32 operatorSetId)
func (_AVSDirectoryStorage *AVSDirectoryStorageFilterer) WatchOperatorAddedToOperatorSet(opts *bind.WatchOpts, sink chan<- *AVSDirectoryStorageOperatorAddedToOperatorSet) (event.Subscription, error) {

	logs, sub, err := _AVSDirectoryStorage.contract.WatchLogs(opts, "OperatorAddedToOperatorSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSDirectoryStorageOperatorAddedToOperatorSet)
				if err := _AVSDirectoryStorage.contract.UnpackLog(event, "OperatorAddedToOperatorSet", log); err != nil {
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

// ParseOperatorAddedToOperatorSet is a log parse operation binding the contract event 0x0e2d64f3a0f18391c95a0a816961e081b15096f8bb957d0b654897b96b7ba309.
//
// Solidity: event OperatorAddedToOperatorSet(address operator, address avs, uint32 operatorSetId)
func (_AVSDirectoryStorage *AVSDirectoryStorageFilterer) ParseOperatorAddedToOperatorSet(log types.Log) (*AVSDirectoryStorageOperatorAddedToOperatorSet, error) {
	event := new(AVSDirectoryStorageOperatorAddedToOperatorSet)
	if err := _AVSDirectoryStorage.contract.UnpackLog(event, "OperatorAddedToOperatorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSDirectoryStorageOperatorRemovedFromOperatorSetIterator is returned from FilterOperatorRemovedFromOperatorSet and is used to iterate over the raw logs and unpacked data for OperatorRemovedFromOperatorSet events raised by the AVSDirectoryStorage contract.
type AVSDirectoryStorageOperatorRemovedFromOperatorSetIterator struct {
	Event *AVSDirectoryStorageOperatorRemovedFromOperatorSet // Event containing the contract specifics and raw log

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
func (it *AVSDirectoryStorageOperatorRemovedFromOperatorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSDirectoryStorageOperatorRemovedFromOperatorSet)
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
		it.Event = new(AVSDirectoryStorageOperatorRemovedFromOperatorSet)
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
func (it *AVSDirectoryStorageOperatorRemovedFromOperatorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSDirectoryStorageOperatorRemovedFromOperatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSDirectoryStorageOperatorRemovedFromOperatorSet represents a OperatorRemovedFromOperatorSet event raised by the AVSDirectoryStorage contract.
type AVSDirectoryStorageOperatorRemovedFromOperatorSet struct {
	Operator      common.Address
	Avs           common.Address
	OperatorSetId uint32
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOperatorRemovedFromOperatorSet is a free log retrieval operation binding the contract event 0x9a431e7219c8c180fcf08db82dbdfe7f5512a825a1febaee77e02b9b2db8df82.
//
// Solidity: event OperatorRemovedFromOperatorSet(address operator, address avs, uint32 operatorSetId)
func (_AVSDirectoryStorage *AVSDirectoryStorageFilterer) FilterOperatorRemovedFromOperatorSet(opts *bind.FilterOpts) (*AVSDirectoryStorageOperatorRemovedFromOperatorSetIterator, error) {

	logs, sub, err := _AVSDirectoryStorage.contract.FilterLogs(opts, "OperatorRemovedFromOperatorSet")
	if err != nil {
		return nil, err
	}
	return &AVSDirectoryStorageOperatorRemovedFromOperatorSetIterator{contract: _AVSDirectoryStorage.contract, event: "OperatorRemovedFromOperatorSet", logs: logs, sub: sub}, nil
}

// WatchOperatorRemovedFromOperatorSet is a free log subscription operation binding the contract event 0x9a431e7219c8c180fcf08db82dbdfe7f5512a825a1febaee77e02b9b2db8df82.
//
// Solidity: event OperatorRemovedFromOperatorSet(address operator, address avs, uint32 operatorSetId)
func (_AVSDirectoryStorage *AVSDirectoryStorageFilterer) WatchOperatorRemovedFromOperatorSet(opts *bind.WatchOpts, sink chan<- *AVSDirectoryStorageOperatorRemovedFromOperatorSet) (event.Subscription, error) {

	logs, sub, err := _AVSDirectoryStorage.contract.WatchLogs(opts, "OperatorRemovedFromOperatorSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSDirectoryStorageOperatorRemovedFromOperatorSet)
				if err := _AVSDirectoryStorage.contract.UnpackLog(event, "OperatorRemovedFromOperatorSet", log); err != nil {
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

// ParseOperatorRemovedFromOperatorSet is a log parse operation binding the contract event 0x9a431e7219c8c180fcf08db82dbdfe7f5512a825a1febaee77e02b9b2db8df82.
//
// Solidity: event OperatorRemovedFromOperatorSet(address operator, address avs, uint32 operatorSetId)
func (_AVSDirectoryStorage *AVSDirectoryStorageFilterer) ParseOperatorRemovedFromOperatorSet(log types.Log) (*AVSDirectoryStorageOperatorRemovedFromOperatorSet, error) {
	event := new(AVSDirectoryStorageOperatorRemovedFromOperatorSet)
	if err := _AVSDirectoryStorage.contract.UnpackLog(event, "OperatorRemovedFromOperatorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
