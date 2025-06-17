// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package OperatorTableUpdaterStorage

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

// BN254G1Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G1Point struct {
	X *big.Int
	Y *big.Int
}

// BN254G2Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G2Point struct {
	X [2]*big.Int
	Y [2]*big.Int
}

// IBN254CertificateVerifierTypesBN254Certificate is an auto generated low-level Go binding around an user-defined struct.
type IBN254CertificateVerifierTypesBN254Certificate struct {
	ReferenceTimestamp uint32
	MessageHash        [32]byte
	Signature          BN254G1Point
	Apk                BN254G2Point
	NonSignerWitnesses []IBN254CertificateVerifierTypesBN254OperatorInfoWitness
}

// IBN254CertificateVerifierTypesBN254OperatorInfoWitness is an auto generated low-level Go binding around an user-defined struct.
type IBN254CertificateVerifierTypesBN254OperatorInfoWitness struct {
	OperatorIndex     uint32
	OperatorInfoProof []byte
	OperatorInfo      IBN254TableCalculatorTypesBN254OperatorInfo
}

// IBN254TableCalculatorTypesBN254OperatorInfo is an auto generated low-level Go binding around an user-defined struct.
type IBN254TableCalculatorTypesBN254OperatorInfo struct {
	Pubkey  BN254G1Point
	Weights []*big.Int
}

// OperatorSet is an auto generated low-level Go binding around an user-defined struct.
type OperatorSet struct {
	Avs common.Address
	Id  uint32
}

// OperatorTableUpdaterStorageMetaData contains all meta data concerning the OperatorTableUpdaterStorage contract.
var OperatorTableUpdaterStorageMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"GLOBAL_TABLE_ROOT_CERT_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_BPS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bn254CertificateVerifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBN254CertificateVerifier\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"confirmGlobalTableRoot\",\"inputs\":[{\"name\":\"globalTableRootCert\",\"type\":\"tuple\",\"internalType\":\"structIBN254CertificateVerifierTypes.BN254Certificate\",\"components\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"nonSignerWitnesses\",\"type\":\"tuple[]\",\"internalType\":\"structIBN254CertificateVerifierTypes.BN254OperatorInfoWitness[]\",\"components\":[{\"name\":\"operatorIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorInfoProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"operatorInfo\",\"type\":\"tuple\",\"internalType\":\"structIBN254TableCalculatorTypes.BN254OperatorInfo\",\"components\":[{\"name\":\"pubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}]}]},{\"name\":\"globalTableRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ecdsaCertificateVerifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIECDSACertificateVerifier\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCertificateVerifier\",\"inputs\":[{\"name\":\"curveType\",\"type\":\"uint8\",\"internalType\":\"enumIKeyRegistrarTypes.CurveType\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentGlobalTableRoot\",\"inputs\":[],\"outputs\":[{\"name\":\"globalTableRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getGlobalConfirmerSetReferenceTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getGlobalRootConfirmerSet\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getGlobalTableRootByTimestamp\",\"inputs\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"tableRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getGlobalTableUpdateMessageHash\",\"inputs\":[{\"name\":\"globalTableRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLatestReferenceBlockNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLatestReferenceTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getReferenceBlockNumber\",\"inputs\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"globalRootConfirmationThreshold\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setGlobalRootConfirmationThreshold\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGlobalRootConfirmerSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOperatorTable\",\"inputs\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"globalTableRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"operatorSetIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"operatorTableBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"GlobalRootConfirmationThresholdUpdated\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GlobalRootConfirmerSetUpdated\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewGlobalTableRoot\",\"inputs\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"globalTableRoot\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CertificateInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"GlobalTableRootInFuture\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"GlobalTableRootStale\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidConfirmationThreshold\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCurveType\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidGlobalTableRoot\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidMessageHash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperatorSetProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignatureLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TableUpdateForPastTimestamp\",\"inputs\":[]}]",
}

// OperatorTableUpdaterStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use OperatorTableUpdaterStorageMetaData.ABI instead.
var OperatorTableUpdaterStorageABI = OperatorTableUpdaterStorageMetaData.ABI

// OperatorTableUpdaterStorage is an auto generated Go binding around an Ethereum contract.
type OperatorTableUpdaterStorage struct {
	OperatorTableUpdaterStorageCaller     // Read-only binding to the contract
	OperatorTableUpdaterStorageTransactor // Write-only binding to the contract
	OperatorTableUpdaterStorageFilterer   // Log filterer for contract events
}

// OperatorTableUpdaterStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type OperatorTableUpdaterStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OperatorTableUpdaterStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OperatorTableUpdaterStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OperatorTableUpdaterStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OperatorTableUpdaterStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OperatorTableUpdaterStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OperatorTableUpdaterStorageSession struct {
	Contract     *OperatorTableUpdaterStorage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// OperatorTableUpdaterStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OperatorTableUpdaterStorageCallerSession struct {
	Contract *OperatorTableUpdaterStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// OperatorTableUpdaterStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OperatorTableUpdaterStorageTransactorSession struct {
	Contract     *OperatorTableUpdaterStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// OperatorTableUpdaterStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type OperatorTableUpdaterStorageRaw struct {
	Contract *OperatorTableUpdaterStorage // Generic contract binding to access the raw methods on
}

// OperatorTableUpdaterStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OperatorTableUpdaterStorageCallerRaw struct {
	Contract *OperatorTableUpdaterStorageCaller // Generic read-only contract binding to access the raw methods on
}

// OperatorTableUpdaterStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OperatorTableUpdaterStorageTransactorRaw struct {
	Contract *OperatorTableUpdaterStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOperatorTableUpdaterStorage creates a new instance of OperatorTableUpdaterStorage, bound to a specific deployed contract.
func NewOperatorTableUpdaterStorage(address common.Address, backend bind.ContractBackend) (*OperatorTableUpdaterStorage, error) {
	contract, err := bindOperatorTableUpdaterStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OperatorTableUpdaterStorage{OperatorTableUpdaterStorageCaller: OperatorTableUpdaterStorageCaller{contract: contract}, OperatorTableUpdaterStorageTransactor: OperatorTableUpdaterStorageTransactor{contract: contract}, OperatorTableUpdaterStorageFilterer: OperatorTableUpdaterStorageFilterer{contract: contract}}, nil
}

// NewOperatorTableUpdaterStorageCaller creates a new read-only instance of OperatorTableUpdaterStorage, bound to a specific deployed contract.
func NewOperatorTableUpdaterStorageCaller(address common.Address, caller bind.ContractCaller) (*OperatorTableUpdaterStorageCaller, error) {
	contract, err := bindOperatorTableUpdaterStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OperatorTableUpdaterStorageCaller{contract: contract}, nil
}

// NewOperatorTableUpdaterStorageTransactor creates a new write-only instance of OperatorTableUpdaterStorage, bound to a specific deployed contract.
func NewOperatorTableUpdaterStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*OperatorTableUpdaterStorageTransactor, error) {
	contract, err := bindOperatorTableUpdaterStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OperatorTableUpdaterStorageTransactor{contract: contract}, nil
}

// NewOperatorTableUpdaterStorageFilterer creates a new log filterer instance of OperatorTableUpdaterStorage, bound to a specific deployed contract.
func NewOperatorTableUpdaterStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*OperatorTableUpdaterStorageFilterer, error) {
	contract, err := bindOperatorTableUpdaterStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OperatorTableUpdaterStorageFilterer{contract: contract}, nil
}

// bindOperatorTableUpdaterStorage binds a generic wrapper to an already deployed contract.
func bindOperatorTableUpdaterStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OperatorTableUpdaterStorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OperatorTableUpdaterStorage.Contract.OperatorTableUpdaterStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OperatorTableUpdaterStorage.Contract.OperatorTableUpdaterStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OperatorTableUpdaterStorage.Contract.OperatorTableUpdaterStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OperatorTableUpdaterStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OperatorTableUpdaterStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OperatorTableUpdaterStorage.Contract.contract.Transact(opts, method, params...)
}

// GLOBALTABLEROOTCERTTYPEHASH is a free data retrieval call binding the contract method 0x3ef6cd7a.
//
// Solidity: function GLOBAL_TABLE_ROOT_CERT_TYPEHASH() view returns(bytes32)
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageCaller) GLOBALTABLEROOTCERTTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _OperatorTableUpdaterStorage.contract.Call(opts, &out, "GLOBAL_TABLE_ROOT_CERT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GLOBALTABLEROOTCERTTYPEHASH is a free data retrieval call binding the contract method 0x3ef6cd7a.
//
// Solidity: function GLOBAL_TABLE_ROOT_CERT_TYPEHASH() view returns(bytes32)
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageSession) GLOBALTABLEROOTCERTTYPEHASH() ([32]byte, error) {
	return _OperatorTableUpdaterStorage.Contract.GLOBALTABLEROOTCERTTYPEHASH(&_OperatorTableUpdaterStorage.CallOpts)
}

// GLOBALTABLEROOTCERTTYPEHASH is a free data retrieval call binding the contract method 0x3ef6cd7a.
//
// Solidity: function GLOBAL_TABLE_ROOT_CERT_TYPEHASH() view returns(bytes32)
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageCallerSession) GLOBALTABLEROOTCERTTYPEHASH() ([32]byte, error) {
	return _OperatorTableUpdaterStorage.Contract.GLOBALTABLEROOTCERTTYPEHASH(&_OperatorTableUpdaterStorage.CallOpts)
}

// MAXBPS is a free data retrieval call binding the contract method 0xfd967f47.
//
// Solidity: function MAX_BPS() view returns(uint16)
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageCaller) MAXBPS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _OperatorTableUpdaterStorage.contract.Call(opts, &out, "MAX_BPS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MAXBPS is a free data retrieval call binding the contract method 0xfd967f47.
//
// Solidity: function MAX_BPS() view returns(uint16)
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageSession) MAXBPS() (uint16, error) {
	return _OperatorTableUpdaterStorage.Contract.MAXBPS(&_OperatorTableUpdaterStorage.CallOpts)
}

// MAXBPS is a free data retrieval call binding the contract method 0xfd967f47.
//
// Solidity: function MAX_BPS() view returns(uint16)
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageCallerSession) MAXBPS() (uint16, error) {
	return _OperatorTableUpdaterStorage.Contract.MAXBPS(&_OperatorTableUpdaterStorage.CallOpts)
}

// Bn254CertificateVerifier is a free data retrieval call binding the contract method 0xb8c14306.
//
// Solidity: function bn254CertificateVerifier() view returns(address)
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageCaller) Bn254CertificateVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OperatorTableUpdaterStorage.contract.Call(opts, &out, "bn254CertificateVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bn254CertificateVerifier is a free data retrieval call binding the contract method 0xb8c14306.
//
// Solidity: function bn254CertificateVerifier() view returns(address)
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageSession) Bn254CertificateVerifier() (common.Address, error) {
	return _OperatorTableUpdaterStorage.Contract.Bn254CertificateVerifier(&_OperatorTableUpdaterStorage.CallOpts)
}

// Bn254CertificateVerifier is a free data retrieval call binding the contract method 0xb8c14306.
//
// Solidity: function bn254CertificateVerifier() view returns(address)
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageCallerSession) Bn254CertificateVerifier() (common.Address, error) {
	return _OperatorTableUpdaterStorage.Contract.Bn254CertificateVerifier(&_OperatorTableUpdaterStorage.CallOpts)
}

// EcdsaCertificateVerifier is a free data retrieval call binding the contract method 0xad0f9582.
//
// Solidity: function ecdsaCertificateVerifier() view returns(address)
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageCaller) EcdsaCertificateVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OperatorTableUpdaterStorage.contract.Call(opts, &out, "ecdsaCertificateVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EcdsaCertificateVerifier is a free data retrieval call binding the contract method 0xad0f9582.
//
// Solidity: function ecdsaCertificateVerifier() view returns(address)
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageSession) EcdsaCertificateVerifier() (common.Address, error) {
	return _OperatorTableUpdaterStorage.Contract.EcdsaCertificateVerifier(&_OperatorTableUpdaterStorage.CallOpts)
}

// EcdsaCertificateVerifier is a free data retrieval call binding the contract method 0xad0f9582.
//
// Solidity: function ecdsaCertificateVerifier() view returns(address)
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageCallerSession) EcdsaCertificateVerifier() (common.Address, error) {
	return _OperatorTableUpdaterStorage.Contract.EcdsaCertificateVerifier(&_OperatorTableUpdaterStorage.CallOpts)
}

// GetCertificateVerifier is a free data retrieval call binding the contract method 0x6f728c50.
//
// Solidity: function getCertificateVerifier(uint8 curveType) view returns(address)
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageCaller) GetCertificateVerifier(opts *bind.CallOpts, curveType uint8) (common.Address, error) {
	var out []interface{}
	err := _OperatorTableUpdaterStorage.contract.Call(opts, &out, "getCertificateVerifier", curveType)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCertificateVerifier is a free data retrieval call binding the contract method 0x6f728c50.
//
// Solidity: function getCertificateVerifier(uint8 curveType) view returns(address)
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageSession) GetCertificateVerifier(curveType uint8) (common.Address, error) {
	return _OperatorTableUpdaterStorage.Contract.GetCertificateVerifier(&_OperatorTableUpdaterStorage.CallOpts, curveType)
}

// GetCertificateVerifier is a free data retrieval call binding the contract method 0x6f728c50.
//
// Solidity: function getCertificateVerifier(uint8 curveType) view returns(address)
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageCallerSession) GetCertificateVerifier(curveType uint8) (common.Address, error) {
	return _OperatorTableUpdaterStorage.Contract.GetCertificateVerifier(&_OperatorTableUpdaterStorage.CallOpts, curveType)
}

// GetCurrentGlobalTableRoot is a free data retrieval call binding the contract method 0x28522d79.
//
// Solidity: function getCurrentGlobalTableRoot() view returns(bytes32 globalTableRoot)
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageCaller) GetCurrentGlobalTableRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _OperatorTableUpdaterStorage.contract.Call(opts, &out, "getCurrentGlobalTableRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetCurrentGlobalTableRoot is a free data retrieval call binding the contract method 0x28522d79.
//
// Solidity: function getCurrentGlobalTableRoot() view returns(bytes32 globalTableRoot)
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageSession) GetCurrentGlobalTableRoot() ([32]byte, error) {
	return _OperatorTableUpdaterStorage.Contract.GetCurrentGlobalTableRoot(&_OperatorTableUpdaterStorage.CallOpts)
}

// GetCurrentGlobalTableRoot is a free data retrieval call binding the contract method 0x28522d79.
//
// Solidity: function getCurrentGlobalTableRoot() view returns(bytes32 globalTableRoot)
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageCallerSession) GetCurrentGlobalTableRoot() ([32]byte, error) {
	return _OperatorTableUpdaterStorage.Contract.GetCurrentGlobalTableRoot(&_OperatorTableUpdaterStorage.CallOpts)
}

// GetGlobalConfirmerSetReferenceTimestamp is a free data retrieval call binding the contract method 0x0f3f8edd.
//
// Solidity: function getGlobalConfirmerSetReferenceTimestamp() view returns(uint32)
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageCaller) GetGlobalConfirmerSetReferenceTimestamp(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _OperatorTableUpdaterStorage.contract.Call(opts, &out, "getGlobalConfirmerSetReferenceTimestamp")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetGlobalConfirmerSetReferenceTimestamp is a free data retrieval call binding the contract method 0x0f3f8edd.
//
// Solidity: function getGlobalConfirmerSetReferenceTimestamp() view returns(uint32)
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageSession) GetGlobalConfirmerSetReferenceTimestamp() (uint32, error) {
	return _OperatorTableUpdaterStorage.Contract.GetGlobalConfirmerSetReferenceTimestamp(&_OperatorTableUpdaterStorage.CallOpts)
}

// GetGlobalConfirmerSetReferenceTimestamp is a free data retrieval call binding the contract method 0x0f3f8edd.
//
// Solidity: function getGlobalConfirmerSetReferenceTimestamp() view returns(uint32)
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageCallerSession) GetGlobalConfirmerSetReferenceTimestamp() (uint32, error) {
	return _OperatorTableUpdaterStorage.Contract.GetGlobalConfirmerSetReferenceTimestamp(&_OperatorTableUpdaterStorage.CallOpts)
}

// GetGlobalRootConfirmerSet is a free data retrieval call binding the contract method 0x46282889.
//
// Solidity: function getGlobalRootConfirmerSet() view returns((address,uint32))
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageCaller) GetGlobalRootConfirmerSet(opts *bind.CallOpts) (OperatorSet, error) {
	var out []interface{}
	err := _OperatorTableUpdaterStorage.contract.Call(opts, &out, "getGlobalRootConfirmerSet")

	if err != nil {
		return *new(OperatorSet), err
	}

	out0 := *abi.ConvertType(out[0], new(OperatorSet)).(*OperatorSet)

	return out0, err

}

// GetGlobalRootConfirmerSet is a free data retrieval call binding the contract method 0x46282889.
//
// Solidity: function getGlobalRootConfirmerSet() view returns((address,uint32))
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageSession) GetGlobalRootConfirmerSet() (OperatorSet, error) {
	return _OperatorTableUpdaterStorage.Contract.GetGlobalRootConfirmerSet(&_OperatorTableUpdaterStorage.CallOpts)
}

// GetGlobalRootConfirmerSet is a free data retrieval call binding the contract method 0x46282889.
//
// Solidity: function getGlobalRootConfirmerSet() view returns((address,uint32))
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageCallerSession) GetGlobalRootConfirmerSet() (OperatorSet, error) {
	return _OperatorTableUpdaterStorage.Contract.GetGlobalRootConfirmerSet(&_OperatorTableUpdaterStorage.CallOpts)
}

// GetGlobalTableRootByTimestamp is a free data retrieval call binding the contract method 0xc5916a39.
//
// Solidity: function getGlobalTableRootByTimestamp(uint32 referenceTimestamp) view returns(bytes32 tableRoot)
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageCaller) GetGlobalTableRootByTimestamp(opts *bind.CallOpts, referenceTimestamp uint32) ([32]byte, error) {
	var out []interface{}
	err := _OperatorTableUpdaterStorage.contract.Call(opts, &out, "getGlobalTableRootByTimestamp", referenceTimestamp)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetGlobalTableRootByTimestamp is a free data retrieval call binding the contract method 0xc5916a39.
//
// Solidity: function getGlobalTableRootByTimestamp(uint32 referenceTimestamp) view returns(bytes32 tableRoot)
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageSession) GetGlobalTableRootByTimestamp(referenceTimestamp uint32) ([32]byte, error) {
	return _OperatorTableUpdaterStorage.Contract.GetGlobalTableRootByTimestamp(&_OperatorTableUpdaterStorage.CallOpts, referenceTimestamp)
}

// GetGlobalTableRootByTimestamp is a free data retrieval call binding the contract method 0xc5916a39.
//
// Solidity: function getGlobalTableRootByTimestamp(uint32 referenceTimestamp) view returns(bytes32 tableRoot)
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageCallerSession) GetGlobalTableRootByTimestamp(referenceTimestamp uint32) ([32]byte, error) {
	return _OperatorTableUpdaterStorage.Contract.GetGlobalTableRootByTimestamp(&_OperatorTableUpdaterStorage.CallOpts, referenceTimestamp)
}

// GetGlobalTableUpdateMessageHash is a free data retrieval call binding the contract method 0xc3be1e33.
//
// Solidity: function getGlobalTableUpdateMessageHash(bytes32 globalTableRoot, uint32 referenceTimestamp, uint32 referenceBlockNumber) view returns(bytes32)
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageCaller) GetGlobalTableUpdateMessageHash(opts *bind.CallOpts, globalTableRoot [32]byte, referenceTimestamp uint32, referenceBlockNumber uint32) ([32]byte, error) {
	var out []interface{}
	err := _OperatorTableUpdaterStorage.contract.Call(opts, &out, "getGlobalTableUpdateMessageHash", globalTableRoot, referenceTimestamp, referenceBlockNumber)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetGlobalTableUpdateMessageHash is a free data retrieval call binding the contract method 0xc3be1e33.
//
// Solidity: function getGlobalTableUpdateMessageHash(bytes32 globalTableRoot, uint32 referenceTimestamp, uint32 referenceBlockNumber) view returns(bytes32)
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageSession) GetGlobalTableUpdateMessageHash(globalTableRoot [32]byte, referenceTimestamp uint32, referenceBlockNumber uint32) ([32]byte, error) {
	return _OperatorTableUpdaterStorage.Contract.GetGlobalTableUpdateMessageHash(&_OperatorTableUpdaterStorage.CallOpts, globalTableRoot, referenceTimestamp, referenceBlockNumber)
}

// GetGlobalTableUpdateMessageHash is a free data retrieval call binding the contract method 0xc3be1e33.
//
// Solidity: function getGlobalTableUpdateMessageHash(bytes32 globalTableRoot, uint32 referenceTimestamp, uint32 referenceBlockNumber) view returns(bytes32)
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageCallerSession) GetGlobalTableUpdateMessageHash(globalTableRoot [32]byte, referenceTimestamp uint32, referenceBlockNumber uint32) ([32]byte, error) {
	return _OperatorTableUpdaterStorage.Contract.GetGlobalTableUpdateMessageHash(&_OperatorTableUpdaterStorage.CallOpts, globalTableRoot, referenceTimestamp, referenceBlockNumber)
}

// GetLatestReferenceBlockNumber is a free data retrieval call binding the contract method 0x31a599d2.
//
// Solidity: function getLatestReferenceBlockNumber() view returns(uint32)
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageCaller) GetLatestReferenceBlockNumber(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _OperatorTableUpdaterStorage.contract.Call(opts, &out, "getLatestReferenceBlockNumber")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetLatestReferenceBlockNumber is a free data retrieval call binding the contract method 0x31a599d2.
//
// Solidity: function getLatestReferenceBlockNumber() view returns(uint32)
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageSession) GetLatestReferenceBlockNumber() (uint32, error) {
	return _OperatorTableUpdaterStorage.Contract.GetLatestReferenceBlockNumber(&_OperatorTableUpdaterStorage.CallOpts)
}

// GetLatestReferenceBlockNumber is a free data retrieval call binding the contract method 0x31a599d2.
//
// Solidity: function getLatestReferenceBlockNumber() view returns(uint32)
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageCallerSession) GetLatestReferenceBlockNumber() (uint32, error) {
	return _OperatorTableUpdaterStorage.Contract.GetLatestReferenceBlockNumber(&_OperatorTableUpdaterStorage.CallOpts)
}

// GetLatestReferenceTimestamp is a free data retrieval call binding the contract method 0x4624e6a3.
//
// Solidity: function getLatestReferenceTimestamp() view returns(uint32)
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageCaller) GetLatestReferenceTimestamp(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _OperatorTableUpdaterStorage.contract.Call(opts, &out, "getLatestReferenceTimestamp")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetLatestReferenceTimestamp is a free data retrieval call binding the contract method 0x4624e6a3.
//
// Solidity: function getLatestReferenceTimestamp() view returns(uint32)
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageSession) GetLatestReferenceTimestamp() (uint32, error) {
	return _OperatorTableUpdaterStorage.Contract.GetLatestReferenceTimestamp(&_OperatorTableUpdaterStorage.CallOpts)
}

// GetLatestReferenceTimestamp is a free data retrieval call binding the contract method 0x4624e6a3.
//
// Solidity: function getLatestReferenceTimestamp() view returns(uint32)
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageCallerSession) GetLatestReferenceTimestamp() (uint32, error) {
	return _OperatorTableUpdaterStorage.Contract.GetLatestReferenceTimestamp(&_OperatorTableUpdaterStorage.CallOpts)
}

// GetReferenceBlockNumber is a free data retrieval call binding the contract method 0x95b6bccb.
//
// Solidity: function getReferenceBlockNumber(uint32 referenceTimestamp) view returns(uint32)
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageCaller) GetReferenceBlockNumber(opts *bind.CallOpts, referenceTimestamp uint32) (uint32, error) {
	var out []interface{}
	err := _OperatorTableUpdaterStorage.contract.Call(opts, &out, "getReferenceBlockNumber", referenceTimestamp)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetReferenceBlockNumber is a free data retrieval call binding the contract method 0x95b6bccb.
//
// Solidity: function getReferenceBlockNumber(uint32 referenceTimestamp) view returns(uint32)
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageSession) GetReferenceBlockNumber(referenceTimestamp uint32) (uint32, error) {
	return _OperatorTableUpdaterStorage.Contract.GetReferenceBlockNumber(&_OperatorTableUpdaterStorage.CallOpts, referenceTimestamp)
}

// GetReferenceBlockNumber is a free data retrieval call binding the contract method 0x95b6bccb.
//
// Solidity: function getReferenceBlockNumber(uint32 referenceTimestamp) view returns(uint32)
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageCallerSession) GetReferenceBlockNumber(referenceTimestamp uint32) (uint32, error) {
	return _OperatorTableUpdaterStorage.Contract.GetReferenceBlockNumber(&_OperatorTableUpdaterStorage.CallOpts, referenceTimestamp)
}

// GlobalRootConfirmationThreshold is a free data retrieval call binding the contract method 0xc252aa22.
//
// Solidity: function globalRootConfirmationThreshold() view returns(uint16)
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageCaller) GlobalRootConfirmationThreshold(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _OperatorTableUpdaterStorage.contract.Call(opts, &out, "globalRootConfirmationThreshold")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GlobalRootConfirmationThreshold is a free data retrieval call binding the contract method 0xc252aa22.
//
// Solidity: function globalRootConfirmationThreshold() view returns(uint16)
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageSession) GlobalRootConfirmationThreshold() (uint16, error) {
	return _OperatorTableUpdaterStorage.Contract.GlobalRootConfirmationThreshold(&_OperatorTableUpdaterStorage.CallOpts)
}

// GlobalRootConfirmationThreshold is a free data retrieval call binding the contract method 0xc252aa22.
//
// Solidity: function globalRootConfirmationThreshold() view returns(uint16)
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageCallerSession) GlobalRootConfirmationThreshold() (uint16, error) {
	return _OperatorTableUpdaterStorage.Contract.GlobalRootConfirmationThreshold(&_OperatorTableUpdaterStorage.CallOpts)
}

// ConfirmGlobalTableRoot is a paid mutator transaction binding the contract method 0xeaaed9d5.
//
// Solidity: function confirmGlobalTableRoot((uint32,bytes32,(uint256,uint256),(uint256[2],uint256[2]),(uint32,bytes,((uint256,uint256),uint256[]))[]) globalTableRootCert, bytes32 globalTableRoot, uint32 referenceTimestamp, uint32 referenceBlockNumber) returns()
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageTransactor) ConfirmGlobalTableRoot(opts *bind.TransactOpts, globalTableRootCert IBN254CertificateVerifierTypesBN254Certificate, globalTableRoot [32]byte, referenceTimestamp uint32, referenceBlockNumber uint32) (*types.Transaction, error) {
	return _OperatorTableUpdaterStorage.contract.Transact(opts, "confirmGlobalTableRoot", globalTableRootCert, globalTableRoot, referenceTimestamp, referenceBlockNumber)
}

// ConfirmGlobalTableRoot is a paid mutator transaction binding the contract method 0xeaaed9d5.
//
// Solidity: function confirmGlobalTableRoot((uint32,bytes32,(uint256,uint256),(uint256[2],uint256[2]),(uint32,bytes,((uint256,uint256),uint256[]))[]) globalTableRootCert, bytes32 globalTableRoot, uint32 referenceTimestamp, uint32 referenceBlockNumber) returns()
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageSession) ConfirmGlobalTableRoot(globalTableRootCert IBN254CertificateVerifierTypesBN254Certificate, globalTableRoot [32]byte, referenceTimestamp uint32, referenceBlockNumber uint32) (*types.Transaction, error) {
	return _OperatorTableUpdaterStorage.Contract.ConfirmGlobalTableRoot(&_OperatorTableUpdaterStorage.TransactOpts, globalTableRootCert, globalTableRoot, referenceTimestamp, referenceBlockNumber)
}

// ConfirmGlobalTableRoot is a paid mutator transaction binding the contract method 0xeaaed9d5.
//
// Solidity: function confirmGlobalTableRoot((uint32,bytes32,(uint256,uint256),(uint256[2],uint256[2]),(uint32,bytes,((uint256,uint256),uint256[]))[]) globalTableRootCert, bytes32 globalTableRoot, uint32 referenceTimestamp, uint32 referenceBlockNumber) returns()
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageTransactorSession) ConfirmGlobalTableRoot(globalTableRootCert IBN254CertificateVerifierTypesBN254Certificate, globalTableRoot [32]byte, referenceTimestamp uint32, referenceBlockNumber uint32) (*types.Transaction, error) {
	return _OperatorTableUpdaterStorage.Contract.ConfirmGlobalTableRoot(&_OperatorTableUpdaterStorage.TransactOpts, globalTableRootCert, globalTableRoot, referenceTimestamp, referenceBlockNumber)
}

// SetGlobalRootConfirmationThreshold is a paid mutator transaction binding the contract method 0x2370356c.
//
// Solidity: function setGlobalRootConfirmationThreshold(uint16 bps) returns()
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageTransactor) SetGlobalRootConfirmationThreshold(opts *bind.TransactOpts, bps uint16) (*types.Transaction, error) {
	return _OperatorTableUpdaterStorage.contract.Transact(opts, "setGlobalRootConfirmationThreshold", bps)
}

// SetGlobalRootConfirmationThreshold is a paid mutator transaction binding the contract method 0x2370356c.
//
// Solidity: function setGlobalRootConfirmationThreshold(uint16 bps) returns()
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageSession) SetGlobalRootConfirmationThreshold(bps uint16) (*types.Transaction, error) {
	return _OperatorTableUpdaterStorage.Contract.SetGlobalRootConfirmationThreshold(&_OperatorTableUpdaterStorage.TransactOpts, bps)
}

// SetGlobalRootConfirmationThreshold is a paid mutator transaction binding the contract method 0x2370356c.
//
// Solidity: function setGlobalRootConfirmationThreshold(uint16 bps) returns()
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageTransactorSession) SetGlobalRootConfirmationThreshold(bps uint16) (*types.Transaction, error) {
	return _OperatorTableUpdaterStorage.Contract.SetGlobalRootConfirmationThreshold(&_OperatorTableUpdaterStorage.TransactOpts, bps)
}

// SetGlobalRootConfirmerSet is a paid mutator transaction binding the contract method 0x0371406e.
//
// Solidity: function setGlobalRootConfirmerSet((address,uint32) operatorSet) returns()
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageTransactor) SetGlobalRootConfirmerSet(opts *bind.TransactOpts, operatorSet OperatorSet) (*types.Transaction, error) {
	return _OperatorTableUpdaterStorage.contract.Transact(opts, "setGlobalRootConfirmerSet", operatorSet)
}

// SetGlobalRootConfirmerSet is a paid mutator transaction binding the contract method 0x0371406e.
//
// Solidity: function setGlobalRootConfirmerSet((address,uint32) operatorSet) returns()
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageSession) SetGlobalRootConfirmerSet(operatorSet OperatorSet) (*types.Transaction, error) {
	return _OperatorTableUpdaterStorage.Contract.SetGlobalRootConfirmerSet(&_OperatorTableUpdaterStorage.TransactOpts, operatorSet)
}

// SetGlobalRootConfirmerSet is a paid mutator transaction binding the contract method 0x0371406e.
//
// Solidity: function setGlobalRootConfirmerSet((address,uint32) operatorSet) returns()
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageTransactorSession) SetGlobalRootConfirmerSet(operatorSet OperatorSet) (*types.Transaction, error) {
	return _OperatorTableUpdaterStorage.Contract.SetGlobalRootConfirmerSet(&_OperatorTableUpdaterStorage.TransactOpts, operatorSet)
}

// UpdateOperatorTable is a paid mutator transaction binding the contract method 0x9ea94778.
//
// Solidity: function updateOperatorTable(uint32 referenceTimestamp, bytes32 globalTableRoot, uint32 operatorSetIndex, bytes proof, bytes operatorTableBytes) returns()
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageTransactor) UpdateOperatorTable(opts *bind.TransactOpts, referenceTimestamp uint32, globalTableRoot [32]byte, operatorSetIndex uint32, proof []byte, operatorTableBytes []byte) (*types.Transaction, error) {
	return _OperatorTableUpdaterStorage.contract.Transact(opts, "updateOperatorTable", referenceTimestamp, globalTableRoot, operatorSetIndex, proof, operatorTableBytes)
}

// UpdateOperatorTable is a paid mutator transaction binding the contract method 0x9ea94778.
//
// Solidity: function updateOperatorTable(uint32 referenceTimestamp, bytes32 globalTableRoot, uint32 operatorSetIndex, bytes proof, bytes operatorTableBytes) returns()
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageSession) UpdateOperatorTable(referenceTimestamp uint32, globalTableRoot [32]byte, operatorSetIndex uint32, proof []byte, operatorTableBytes []byte) (*types.Transaction, error) {
	return _OperatorTableUpdaterStorage.Contract.UpdateOperatorTable(&_OperatorTableUpdaterStorage.TransactOpts, referenceTimestamp, globalTableRoot, operatorSetIndex, proof, operatorTableBytes)
}

// UpdateOperatorTable is a paid mutator transaction binding the contract method 0x9ea94778.
//
// Solidity: function updateOperatorTable(uint32 referenceTimestamp, bytes32 globalTableRoot, uint32 operatorSetIndex, bytes proof, bytes operatorTableBytes) returns()
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageTransactorSession) UpdateOperatorTable(referenceTimestamp uint32, globalTableRoot [32]byte, operatorSetIndex uint32, proof []byte, operatorTableBytes []byte) (*types.Transaction, error) {
	return _OperatorTableUpdaterStorage.Contract.UpdateOperatorTable(&_OperatorTableUpdaterStorage.TransactOpts, referenceTimestamp, globalTableRoot, operatorSetIndex, proof, operatorTableBytes)
}

// OperatorTableUpdaterStorageGlobalRootConfirmationThresholdUpdatedIterator is returned from FilterGlobalRootConfirmationThresholdUpdated and is used to iterate over the raw logs and unpacked data for GlobalRootConfirmationThresholdUpdated events raised by the OperatorTableUpdaterStorage contract.
type OperatorTableUpdaterStorageGlobalRootConfirmationThresholdUpdatedIterator struct {
	Event *OperatorTableUpdaterStorageGlobalRootConfirmationThresholdUpdated // Event containing the contract specifics and raw log

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
func (it *OperatorTableUpdaterStorageGlobalRootConfirmationThresholdUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorTableUpdaterStorageGlobalRootConfirmationThresholdUpdated)
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
		it.Event = new(OperatorTableUpdaterStorageGlobalRootConfirmationThresholdUpdated)
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
func (it *OperatorTableUpdaterStorageGlobalRootConfirmationThresholdUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorTableUpdaterStorageGlobalRootConfirmationThresholdUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorTableUpdaterStorageGlobalRootConfirmationThresholdUpdated represents a GlobalRootConfirmationThresholdUpdated event raised by the OperatorTableUpdaterStorage contract.
type OperatorTableUpdaterStorageGlobalRootConfirmationThresholdUpdated struct {
	Bps uint16
	Raw types.Log // Blockchain specific contextual infos
}

// FilterGlobalRootConfirmationThresholdUpdated is a free log retrieval operation binding the contract event 0xf5d1836df8fcd7c1e54047e94ac8773d2855395603e2ef9ba5f5f16905f22592.
//
// Solidity: event GlobalRootConfirmationThresholdUpdated(uint16 bps)
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageFilterer) FilterGlobalRootConfirmationThresholdUpdated(opts *bind.FilterOpts) (*OperatorTableUpdaterStorageGlobalRootConfirmationThresholdUpdatedIterator, error) {

	logs, sub, err := _OperatorTableUpdaterStorage.contract.FilterLogs(opts, "GlobalRootConfirmationThresholdUpdated")
	if err != nil {
		return nil, err
	}
	return &OperatorTableUpdaterStorageGlobalRootConfirmationThresholdUpdatedIterator{contract: _OperatorTableUpdaterStorage.contract, event: "GlobalRootConfirmationThresholdUpdated", logs: logs, sub: sub}, nil
}

// WatchGlobalRootConfirmationThresholdUpdated is a free log subscription operation binding the contract event 0xf5d1836df8fcd7c1e54047e94ac8773d2855395603e2ef9ba5f5f16905f22592.
//
// Solidity: event GlobalRootConfirmationThresholdUpdated(uint16 bps)
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageFilterer) WatchGlobalRootConfirmationThresholdUpdated(opts *bind.WatchOpts, sink chan<- *OperatorTableUpdaterStorageGlobalRootConfirmationThresholdUpdated) (event.Subscription, error) {

	logs, sub, err := _OperatorTableUpdaterStorage.contract.WatchLogs(opts, "GlobalRootConfirmationThresholdUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorTableUpdaterStorageGlobalRootConfirmationThresholdUpdated)
				if err := _OperatorTableUpdaterStorage.contract.UnpackLog(event, "GlobalRootConfirmationThresholdUpdated", log); err != nil {
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

// ParseGlobalRootConfirmationThresholdUpdated is a log parse operation binding the contract event 0xf5d1836df8fcd7c1e54047e94ac8773d2855395603e2ef9ba5f5f16905f22592.
//
// Solidity: event GlobalRootConfirmationThresholdUpdated(uint16 bps)
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageFilterer) ParseGlobalRootConfirmationThresholdUpdated(log types.Log) (*OperatorTableUpdaterStorageGlobalRootConfirmationThresholdUpdated, error) {
	event := new(OperatorTableUpdaterStorageGlobalRootConfirmationThresholdUpdated)
	if err := _OperatorTableUpdaterStorage.contract.UnpackLog(event, "GlobalRootConfirmationThresholdUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorTableUpdaterStorageGlobalRootConfirmerSetUpdatedIterator is returned from FilterGlobalRootConfirmerSetUpdated and is used to iterate over the raw logs and unpacked data for GlobalRootConfirmerSetUpdated events raised by the OperatorTableUpdaterStorage contract.
type OperatorTableUpdaterStorageGlobalRootConfirmerSetUpdatedIterator struct {
	Event *OperatorTableUpdaterStorageGlobalRootConfirmerSetUpdated // Event containing the contract specifics and raw log

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
func (it *OperatorTableUpdaterStorageGlobalRootConfirmerSetUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorTableUpdaterStorageGlobalRootConfirmerSetUpdated)
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
		it.Event = new(OperatorTableUpdaterStorageGlobalRootConfirmerSetUpdated)
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
func (it *OperatorTableUpdaterStorageGlobalRootConfirmerSetUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorTableUpdaterStorageGlobalRootConfirmerSetUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorTableUpdaterStorageGlobalRootConfirmerSetUpdated represents a GlobalRootConfirmerSetUpdated event raised by the OperatorTableUpdaterStorage contract.
type OperatorTableUpdaterStorageGlobalRootConfirmerSetUpdated struct {
	OperatorSet OperatorSet
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterGlobalRootConfirmerSetUpdated is a free log retrieval operation binding the contract event 0x20100394950e66014c25009b45d12b675210a6e7a002044a0e3de6544e3c4b37.
//
// Solidity: event GlobalRootConfirmerSetUpdated((address,uint32) operatorSet)
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageFilterer) FilterGlobalRootConfirmerSetUpdated(opts *bind.FilterOpts) (*OperatorTableUpdaterStorageGlobalRootConfirmerSetUpdatedIterator, error) {

	logs, sub, err := _OperatorTableUpdaterStorage.contract.FilterLogs(opts, "GlobalRootConfirmerSetUpdated")
	if err != nil {
		return nil, err
	}
	return &OperatorTableUpdaterStorageGlobalRootConfirmerSetUpdatedIterator{contract: _OperatorTableUpdaterStorage.contract, event: "GlobalRootConfirmerSetUpdated", logs: logs, sub: sub}, nil
}

// WatchGlobalRootConfirmerSetUpdated is a free log subscription operation binding the contract event 0x20100394950e66014c25009b45d12b675210a6e7a002044a0e3de6544e3c4b37.
//
// Solidity: event GlobalRootConfirmerSetUpdated((address,uint32) operatorSet)
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageFilterer) WatchGlobalRootConfirmerSetUpdated(opts *bind.WatchOpts, sink chan<- *OperatorTableUpdaterStorageGlobalRootConfirmerSetUpdated) (event.Subscription, error) {

	logs, sub, err := _OperatorTableUpdaterStorage.contract.WatchLogs(opts, "GlobalRootConfirmerSetUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorTableUpdaterStorageGlobalRootConfirmerSetUpdated)
				if err := _OperatorTableUpdaterStorage.contract.UnpackLog(event, "GlobalRootConfirmerSetUpdated", log); err != nil {
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

// ParseGlobalRootConfirmerSetUpdated is a log parse operation binding the contract event 0x20100394950e66014c25009b45d12b675210a6e7a002044a0e3de6544e3c4b37.
//
// Solidity: event GlobalRootConfirmerSetUpdated((address,uint32) operatorSet)
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageFilterer) ParseGlobalRootConfirmerSetUpdated(log types.Log) (*OperatorTableUpdaterStorageGlobalRootConfirmerSetUpdated, error) {
	event := new(OperatorTableUpdaterStorageGlobalRootConfirmerSetUpdated)
	if err := _OperatorTableUpdaterStorage.contract.UnpackLog(event, "GlobalRootConfirmerSetUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorTableUpdaterStorageNewGlobalTableRootIterator is returned from FilterNewGlobalTableRoot and is used to iterate over the raw logs and unpacked data for NewGlobalTableRoot events raised by the OperatorTableUpdaterStorage contract.
type OperatorTableUpdaterStorageNewGlobalTableRootIterator struct {
	Event *OperatorTableUpdaterStorageNewGlobalTableRoot // Event containing the contract specifics and raw log

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
func (it *OperatorTableUpdaterStorageNewGlobalTableRootIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorTableUpdaterStorageNewGlobalTableRoot)
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
		it.Event = new(OperatorTableUpdaterStorageNewGlobalTableRoot)
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
func (it *OperatorTableUpdaterStorageNewGlobalTableRootIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorTableUpdaterStorageNewGlobalTableRootIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorTableUpdaterStorageNewGlobalTableRoot represents a NewGlobalTableRoot event raised by the OperatorTableUpdaterStorage contract.
type OperatorTableUpdaterStorageNewGlobalTableRoot struct {
	ReferenceTimestamp uint32
	GlobalTableRoot    [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterNewGlobalTableRoot is a free log retrieval operation binding the contract event 0x010dcbe0d1e019c93357711f7bb6287d543b7ff7de74f29df3fb5ecceec8d369.
//
// Solidity: event NewGlobalTableRoot(uint32 indexed referenceTimestamp, bytes32 indexed globalTableRoot)
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageFilterer) FilterNewGlobalTableRoot(opts *bind.FilterOpts, referenceTimestamp []uint32, globalTableRoot [][32]byte) (*OperatorTableUpdaterStorageNewGlobalTableRootIterator, error) {

	var referenceTimestampRule []interface{}
	for _, referenceTimestampItem := range referenceTimestamp {
		referenceTimestampRule = append(referenceTimestampRule, referenceTimestampItem)
	}
	var globalTableRootRule []interface{}
	for _, globalTableRootItem := range globalTableRoot {
		globalTableRootRule = append(globalTableRootRule, globalTableRootItem)
	}

	logs, sub, err := _OperatorTableUpdaterStorage.contract.FilterLogs(opts, "NewGlobalTableRoot", referenceTimestampRule, globalTableRootRule)
	if err != nil {
		return nil, err
	}
	return &OperatorTableUpdaterStorageNewGlobalTableRootIterator{contract: _OperatorTableUpdaterStorage.contract, event: "NewGlobalTableRoot", logs: logs, sub: sub}, nil
}

// WatchNewGlobalTableRoot is a free log subscription operation binding the contract event 0x010dcbe0d1e019c93357711f7bb6287d543b7ff7de74f29df3fb5ecceec8d369.
//
// Solidity: event NewGlobalTableRoot(uint32 indexed referenceTimestamp, bytes32 indexed globalTableRoot)
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageFilterer) WatchNewGlobalTableRoot(opts *bind.WatchOpts, sink chan<- *OperatorTableUpdaterStorageNewGlobalTableRoot, referenceTimestamp []uint32, globalTableRoot [][32]byte) (event.Subscription, error) {

	var referenceTimestampRule []interface{}
	for _, referenceTimestampItem := range referenceTimestamp {
		referenceTimestampRule = append(referenceTimestampRule, referenceTimestampItem)
	}
	var globalTableRootRule []interface{}
	for _, globalTableRootItem := range globalTableRoot {
		globalTableRootRule = append(globalTableRootRule, globalTableRootItem)
	}

	logs, sub, err := _OperatorTableUpdaterStorage.contract.WatchLogs(opts, "NewGlobalTableRoot", referenceTimestampRule, globalTableRootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorTableUpdaterStorageNewGlobalTableRoot)
				if err := _OperatorTableUpdaterStorage.contract.UnpackLog(event, "NewGlobalTableRoot", log); err != nil {
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

// ParseNewGlobalTableRoot is a log parse operation binding the contract event 0x010dcbe0d1e019c93357711f7bb6287d543b7ff7de74f29df3fb5ecceec8d369.
//
// Solidity: event NewGlobalTableRoot(uint32 indexed referenceTimestamp, bytes32 indexed globalTableRoot)
func (_OperatorTableUpdaterStorage *OperatorTableUpdaterStorageFilterer) ParseNewGlobalTableRoot(log types.Log) (*OperatorTableUpdaterStorageNewGlobalTableRoot, error) {
	event := new(OperatorTableUpdaterStorageNewGlobalTableRoot)
	if err := _OperatorTableUpdaterStorage.contract.UnpackLog(event, "NewGlobalTableRoot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
