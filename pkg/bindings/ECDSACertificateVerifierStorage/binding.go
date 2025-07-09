// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ECDSACertificateVerifierStorage

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

// ICrossChainRegistryTypesOperatorSetConfig is an auto generated low-level Go binding around an user-defined struct.
type ICrossChainRegistryTypesOperatorSetConfig struct {
	Owner              common.Address
	MaxStalenessPeriod uint32
}

// IECDSACertificateVerifierTypesECDSACertificate is an auto generated low-level Go binding around an user-defined struct.
type IECDSACertificateVerifierTypesECDSACertificate struct {
	ReferenceTimestamp uint32
	MessageHash        [32]byte
	Sig                []byte
}

// IOperatorTableCalculatorTypesECDSAOperatorInfo is an auto generated low-level Go binding around an user-defined struct.
type IOperatorTableCalculatorTypesECDSAOperatorInfo struct {
	Pubkey  common.Address
	Weights []*big.Int
}

// OperatorSet is an auto generated low-level Go binding around an user-defined struct.
type OperatorSet struct {
	Avs common.Address
	Id  uint32
}

// ECDSACertificateVerifierStorageMetaData contains all meta data concerning the ECDSACertificateVerifierStorage contract.
var ECDSACertificateVerifierStorageMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"calculateCertificateDigest\",\"inputs\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"domainSeparator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorCount\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorInfo\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIOperatorTableCalculatorTypes.ECDSAOperatorInfo\",\"components\":[{\"name\":\"pubkey\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"weights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorInfos\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIOperatorTableCalculatorTypes.ECDSAOperatorInfo[]\",\"components\":[{\"name\":\"pubkey\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"weights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorSetOwner\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalStakes\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestReferenceTimestamp\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxOperatorTableStaleness\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorTableUpdater\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIOperatorTableUpdater\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateOperatorTable\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorInfos\",\"type\":\"tuple[]\",\"internalType\":\"structIOperatorTableCalculatorTypes.ECDSAOperatorInfo[]\",\"components\":[{\"name\":\"pubkey\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"weights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]},{\"name\":\"operatorSetConfig\",\"type\":\"tuple\",\"internalType\":\"structICrossChainRegistryTypes.OperatorSetConfig\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxStalenessPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyCertificate\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"cert\",\"type\":\"tuple\",\"internalType\":\"structIECDSACertificateVerifierTypes.ECDSACertificate\",\"components\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sig\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"signedStakes\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"signers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyCertificateNominal\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"cert\",\"type\":\"tuple\",\"internalType\":\"structIECDSACertificateVerifierTypes.ECDSACertificate\",\"components\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sig\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"totalStakeNominalThresholds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"signers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyCertificateProportion\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"cert\",\"type\":\"tuple\",\"internalType\":\"structIECDSACertificateVerifierTypes.ECDSACertificate\",\"components\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sig\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"totalStakeProportionThresholds\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"signers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"MaxStalenessPeriodUpdated\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"maxStalenessPeriod\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSetOwnerUpdated\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TableUpdated\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"operatorInfos\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"structIOperatorTableCalculatorTypes.ECDSAOperatorInfo[]\",\"components\":[{\"name\":\"pubkey\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"weights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ArrayLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CertificateStale\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignatureLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyTableUpdater\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReferenceTimestampDoesNotExist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RootDisabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignersNotOrdered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TableUpdateStale\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"VerificationFailed\",\"inputs\":[]}]",
}

// ECDSACertificateVerifierStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use ECDSACertificateVerifierStorageMetaData.ABI instead.
var ECDSACertificateVerifierStorageABI = ECDSACertificateVerifierStorageMetaData.ABI

// ECDSACertificateVerifierStorage is an auto generated Go binding around an Ethereum contract.
type ECDSACertificateVerifierStorage struct {
	ECDSACertificateVerifierStorageCaller     // Read-only binding to the contract
	ECDSACertificateVerifierStorageTransactor // Write-only binding to the contract
	ECDSACertificateVerifierStorageFilterer   // Log filterer for contract events
}

// ECDSACertificateVerifierStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type ECDSACertificateVerifierStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSACertificateVerifierStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ECDSACertificateVerifierStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSACertificateVerifierStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ECDSACertificateVerifierStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSACertificateVerifierStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ECDSACertificateVerifierStorageSession struct {
	Contract     *ECDSACertificateVerifierStorage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                    // Call options to use throughout this session
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// ECDSACertificateVerifierStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ECDSACertificateVerifierStorageCallerSession struct {
	Contract *ECDSACertificateVerifierStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                          // Call options to use throughout this session
}

// ECDSACertificateVerifierStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ECDSACertificateVerifierStorageTransactorSession struct {
	Contract     *ECDSACertificateVerifierStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                          // Transaction auth options to use throughout this session
}

// ECDSACertificateVerifierStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type ECDSACertificateVerifierStorageRaw struct {
	Contract *ECDSACertificateVerifierStorage // Generic contract binding to access the raw methods on
}

// ECDSACertificateVerifierStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ECDSACertificateVerifierStorageCallerRaw struct {
	Contract *ECDSACertificateVerifierStorageCaller // Generic read-only contract binding to access the raw methods on
}

// ECDSACertificateVerifierStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ECDSACertificateVerifierStorageTransactorRaw struct {
	Contract *ECDSACertificateVerifierStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewECDSACertificateVerifierStorage creates a new instance of ECDSACertificateVerifierStorage, bound to a specific deployed contract.
func NewECDSACertificateVerifierStorage(address common.Address, backend bind.ContractBackend) (*ECDSACertificateVerifierStorage, error) {
	contract, err := bindECDSACertificateVerifierStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ECDSACertificateVerifierStorage{ECDSACertificateVerifierStorageCaller: ECDSACertificateVerifierStorageCaller{contract: contract}, ECDSACertificateVerifierStorageTransactor: ECDSACertificateVerifierStorageTransactor{contract: contract}, ECDSACertificateVerifierStorageFilterer: ECDSACertificateVerifierStorageFilterer{contract: contract}}, nil
}

// NewECDSACertificateVerifierStorageCaller creates a new read-only instance of ECDSACertificateVerifierStorage, bound to a specific deployed contract.
func NewECDSACertificateVerifierStorageCaller(address common.Address, caller bind.ContractCaller) (*ECDSACertificateVerifierStorageCaller, error) {
	contract, err := bindECDSACertificateVerifierStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ECDSACertificateVerifierStorageCaller{contract: contract}, nil
}

// NewECDSACertificateVerifierStorageTransactor creates a new write-only instance of ECDSACertificateVerifierStorage, bound to a specific deployed contract.
func NewECDSACertificateVerifierStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*ECDSACertificateVerifierStorageTransactor, error) {
	contract, err := bindECDSACertificateVerifierStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ECDSACertificateVerifierStorageTransactor{contract: contract}, nil
}

// NewECDSACertificateVerifierStorageFilterer creates a new log filterer instance of ECDSACertificateVerifierStorage, bound to a specific deployed contract.
func NewECDSACertificateVerifierStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*ECDSACertificateVerifierStorageFilterer, error) {
	contract, err := bindECDSACertificateVerifierStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ECDSACertificateVerifierStorageFilterer{contract: contract}, nil
}

// bindECDSACertificateVerifierStorage binds a generic wrapper to an already deployed contract.
func bindECDSACertificateVerifierStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ECDSACertificateVerifierStorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECDSACertificateVerifierStorage *ECDSACertificateVerifierStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ECDSACertificateVerifierStorage.Contract.ECDSACertificateVerifierStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECDSACertificateVerifierStorage *ECDSACertificateVerifierStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECDSACertificateVerifierStorage.Contract.ECDSACertificateVerifierStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECDSACertificateVerifierStorage *ECDSACertificateVerifierStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECDSACertificateVerifierStorage.Contract.ECDSACertificateVerifierStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECDSACertificateVerifierStorage *ECDSACertificateVerifierStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ECDSACertificateVerifierStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECDSACertificateVerifierStorage *ECDSACertificateVerifierStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECDSACertificateVerifierStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECDSACertificateVerifierStorage *ECDSACertificateVerifierStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECDSACertificateVerifierStorage.Contract.contract.Transact(opts, method, params...)
}

// CalculateCertificateDigest is a free data retrieval call binding the contract method 0x18467434.
//
// Solidity: function calculateCertificateDigest(uint32 referenceTimestamp, bytes32 messageHash) view returns(bytes32)
func (_ECDSACertificateVerifierStorage *ECDSACertificateVerifierStorageCaller) CalculateCertificateDigest(opts *bind.CallOpts, referenceTimestamp uint32, messageHash [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ECDSACertificateVerifierStorage.contract.Call(opts, &out, "calculateCertificateDigest", referenceTimestamp, messageHash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateCertificateDigest is a free data retrieval call binding the contract method 0x18467434.
//
// Solidity: function calculateCertificateDigest(uint32 referenceTimestamp, bytes32 messageHash) view returns(bytes32)
func (_ECDSACertificateVerifierStorage *ECDSACertificateVerifierStorageSession) CalculateCertificateDigest(referenceTimestamp uint32, messageHash [32]byte) ([32]byte, error) {
	return _ECDSACertificateVerifierStorage.Contract.CalculateCertificateDigest(&_ECDSACertificateVerifierStorage.CallOpts, referenceTimestamp, messageHash)
}

// CalculateCertificateDigest is a free data retrieval call binding the contract method 0x18467434.
//
// Solidity: function calculateCertificateDigest(uint32 referenceTimestamp, bytes32 messageHash) view returns(bytes32)
func (_ECDSACertificateVerifierStorage *ECDSACertificateVerifierStorageCallerSession) CalculateCertificateDigest(referenceTimestamp uint32, messageHash [32]byte) ([32]byte, error) {
	return _ECDSACertificateVerifierStorage.Contract.CalculateCertificateDigest(&_ECDSACertificateVerifierStorage.CallOpts, referenceTimestamp, messageHash)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_ECDSACertificateVerifierStorage *ECDSACertificateVerifierStorageCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ECDSACertificateVerifierStorage.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_ECDSACertificateVerifierStorage *ECDSACertificateVerifierStorageSession) DomainSeparator() ([32]byte, error) {
	return _ECDSACertificateVerifierStorage.Contract.DomainSeparator(&_ECDSACertificateVerifierStorage.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_ECDSACertificateVerifierStorage *ECDSACertificateVerifierStorageCallerSession) DomainSeparator() ([32]byte, error) {
	return _ECDSACertificateVerifierStorage.Contract.DomainSeparator(&_ECDSACertificateVerifierStorage.CallOpts)
}

// GetOperatorCount is a free data retrieval call binding the contract method 0x23c2a3cb.
//
// Solidity: function getOperatorCount((address,uint32) operatorSet, uint32 referenceTimestamp) view returns(uint32)
func (_ECDSACertificateVerifierStorage *ECDSACertificateVerifierStorageCaller) GetOperatorCount(opts *bind.CallOpts, operatorSet OperatorSet, referenceTimestamp uint32) (uint32, error) {
	var out []interface{}
	err := _ECDSACertificateVerifierStorage.contract.Call(opts, &out, "getOperatorCount", operatorSet, referenceTimestamp)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetOperatorCount is a free data retrieval call binding the contract method 0x23c2a3cb.
//
// Solidity: function getOperatorCount((address,uint32) operatorSet, uint32 referenceTimestamp) view returns(uint32)
func (_ECDSACertificateVerifierStorage *ECDSACertificateVerifierStorageSession) GetOperatorCount(operatorSet OperatorSet, referenceTimestamp uint32) (uint32, error) {
	return _ECDSACertificateVerifierStorage.Contract.GetOperatorCount(&_ECDSACertificateVerifierStorage.CallOpts, operatorSet, referenceTimestamp)
}

// GetOperatorCount is a free data retrieval call binding the contract method 0x23c2a3cb.
//
// Solidity: function getOperatorCount((address,uint32) operatorSet, uint32 referenceTimestamp) view returns(uint32)
func (_ECDSACertificateVerifierStorage *ECDSACertificateVerifierStorageCallerSession) GetOperatorCount(operatorSet OperatorSet, referenceTimestamp uint32) (uint32, error) {
	return _ECDSACertificateVerifierStorage.Contract.GetOperatorCount(&_ECDSACertificateVerifierStorage.CallOpts, operatorSet, referenceTimestamp)
}

// GetOperatorInfo is a free data retrieval call binding the contract method 0x082ef73d.
//
// Solidity: function getOperatorInfo((address,uint32) operatorSet, uint32 referenceTimestamp, uint32 operatorIndex) view returns((address,uint256[]))
func (_ECDSACertificateVerifierStorage *ECDSACertificateVerifierStorageCaller) GetOperatorInfo(opts *bind.CallOpts, operatorSet OperatorSet, referenceTimestamp uint32, operatorIndex uint32) (IOperatorTableCalculatorTypesECDSAOperatorInfo, error) {
	var out []interface{}
	err := _ECDSACertificateVerifierStorage.contract.Call(opts, &out, "getOperatorInfo", operatorSet, referenceTimestamp, operatorIndex)

	if err != nil {
		return *new(IOperatorTableCalculatorTypesECDSAOperatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IOperatorTableCalculatorTypesECDSAOperatorInfo)).(*IOperatorTableCalculatorTypesECDSAOperatorInfo)

	return out0, err

}

// GetOperatorInfo is a free data retrieval call binding the contract method 0x082ef73d.
//
// Solidity: function getOperatorInfo((address,uint32) operatorSet, uint32 referenceTimestamp, uint32 operatorIndex) view returns((address,uint256[]))
func (_ECDSACertificateVerifierStorage *ECDSACertificateVerifierStorageSession) GetOperatorInfo(operatorSet OperatorSet, referenceTimestamp uint32, operatorIndex uint32) (IOperatorTableCalculatorTypesECDSAOperatorInfo, error) {
	return _ECDSACertificateVerifierStorage.Contract.GetOperatorInfo(&_ECDSACertificateVerifierStorage.CallOpts, operatorSet, referenceTimestamp, operatorIndex)
}

// GetOperatorInfo is a free data retrieval call binding the contract method 0x082ef73d.
//
// Solidity: function getOperatorInfo((address,uint32) operatorSet, uint32 referenceTimestamp, uint32 operatorIndex) view returns((address,uint256[]))
func (_ECDSACertificateVerifierStorage *ECDSACertificateVerifierStorageCallerSession) GetOperatorInfo(operatorSet OperatorSet, referenceTimestamp uint32, operatorIndex uint32) (IOperatorTableCalculatorTypesECDSAOperatorInfo, error) {
	return _ECDSACertificateVerifierStorage.Contract.GetOperatorInfo(&_ECDSACertificateVerifierStorage.CallOpts, operatorSet, referenceTimestamp, operatorIndex)
}

// GetOperatorInfos is a free data retrieval call binding the contract method 0x7c85ac4c.
//
// Solidity: function getOperatorInfos((address,uint32) operatorSet, uint32 referenceTimestamp) view returns((address,uint256[])[])
func (_ECDSACertificateVerifierStorage *ECDSACertificateVerifierStorageCaller) GetOperatorInfos(opts *bind.CallOpts, operatorSet OperatorSet, referenceTimestamp uint32) ([]IOperatorTableCalculatorTypesECDSAOperatorInfo, error) {
	var out []interface{}
	err := _ECDSACertificateVerifierStorage.contract.Call(opts, &out, "getOperatorInfos", operatorSet, referenceTimestamp)

	if err != nil {
		return *new([]IOperatorTableCalculatorTypesECDSAOperatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IOperatorTableCalculatorTypesECDSAOperatorInfo)).(*[]IOperatorTableCalculatorTypesECDSAOperatorInfo)

	return out0, err

}

// GetOperatorInfos is a free data retrieval call binding the contract method 0x7c85ac4c.
//
// Solidity: function getOperatorInfos((address,uint32) operatorSet, uint32 referenceTimestamp) view returns((address,uint256[])[])
func (_ECDSACertificateVerifierStorage *ECDSACertificateVerifierStorageSession) GetOperatorInfos(operatorSet OperatorSet, referenceTimestamp uint32) ([]IOperatorTableCalculatorTypesECDSAOperatorInfo, error) {
	return _ECDSACertificateVerifierStorage.Contract.GetOperatorInfos(&_ECDSACertificateVerifierStorage.CallOpts, operatorSet, referenceTimestamp)
}

// GetOperatorInfos is a free data retrieval call binding the contract method 0x7c85ac4c.
//
// Solidity: function getOperatorInfos((address,uint32) operatorSet, uint32 referenceTimestamp) view returns((address,uint256[])[])
func (_ECDSACertificateVerifierStorage *ECDSACertificateVerifierStorageCallerSession) GetOperatorInfos(operatorSet OperatorSet, referenceTimestamp uint32) ([]IOperatorTableCalculatorTypesECDSAOperatorInfo, error) {
	return _ECDSACertificateVerifierStorage.Contract.GetOperatorInfos(&_ECDSACertificateVerifierStorage.CallOpts, operatorSet, referenceTimestamp)
}

// GetOperatorSetOwner is a free data retrieval call binding the contract method 0x84818920.
//
// Solidity: function getOperatorSetOwner((address,uint32) operatorSet) view returns(address)
func (_ECDSACertificateVerifierStorage *ECDSACertificateVerifierStorageCaller) GetOperatorSetOwner(opts *bind.CallOpts, operatorSet OperatorSet) (common.Address, error) {
	var out []interface{}
	err := _ECDSACertificateVerifierStorage.contract.Call(opts, &out, "getOperatorSetOwner", operatorSet)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOperatorSetOwner is a free data retrieval call binding the contract method 0x84818920.
//
// Solidity: function getOperatorSetOwner((address,uint32) operatorSet) view returns(address)
func (_ECDSACertificateVerifierStorage *ECDSACertificateVerifierStorageSession) GetOperatorSetOwner(operatorSet OperatorSet) (common.Address, error) {
	return _ECDSACertificateVerifierStorage.Contract.GetOperatorSetOwner(&_ECDSACertificateVerifierStorage.CallOpts, operatorSet)
}

// GetOperatorSetOwner is a free data retrieval call binding the contract method 0x84818920.
//
// Solidity: function getOperatorSetOwner((address,uint32) operatorSet) view returns(address)
func (_ECDSACertificateVerifierStorage *ECDSACertificateVerifierStorageCallerSession) GetOperatorSetOwner(operatorSet OperatorSet) (common.Address, error) {
	return _ECDSACertificateVerifierStorage.Contract.GetOperatorSetOwner(&_ECDSACertificateVerifierStorage.CallOpts, operatorSet)
}

// GetTotalStakes is a free data retrieval call binding the contract method 0x04cdbae4.
//
// Solidity: function getTotalStakes((address,uint32) operatorSet, uint32 referenceTimestamp) view returns(uint256[])
func (_ECDSACertificateVerifierStorage *ECDSACertificateVerifierStorageCaller) GetTotalStakes(opts *bind.CallOpts, operatorSet OperatorSet, referenceTimestamp uint32) ([]*big.Int, error) {
	var out []interface{}
	err := _ECDSACertificateVerifierStorage.contract.Call(opts, &out, "getTotalStakes", operatorSet, referenceTimestamp)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetTotalStakes is a free data retrieval call binding the contract method 0x04cdbae4.
//
// Solidity: function getTotalStakes((address,uint32) operatorSet, uint32 referenceTimestamp) view returns(uint256[])
func (_ECDSACertificateVerifierStorage *ECDSACertificateVerifierStorageSession) GetTotalStakes(operatorSet OperatorSet, referenceTimestamp uint32) ([]*big.Int, error) {
	return _ECDSACertificateVerifierStorage.Contract.GetTotalStakes(&_ECDSACertificateVerifierStorage.CallOpts, operatorSet, referenceTimestamp)
}

// GetTotalStakes is a free data retrieval call binding the contract method 0x04cdbae4.
//
// Solidity: function getTotalStakes((address,uint32) operatorSet, uint32 referenceTimestamp) view returns(uint256[])
func (_ECDSACertificateVerifierStorage *ECDSACertificateVerifierStorageCallerSession) GetTotalStakes(operatorSet OperatorSet, referenceTimestamp uint32) ([]*big.Int, error) {
	return _ECDSACertificateVerifierStorage.Contract.GetTotalStakes(&_ECDSACertificateVerifierStorage.CallOpts, operatorSet, referenceTimestamp)
}

// LatestReferenceTimestamp is a free data retrieval call binding the contract method 0x5ddb9b5b.
//
// Solidity: function latestReferenceTimestamp((address,uint32) operatorSet) view returns(uint32)
func (_ECDSACertificateVerifierStorage *ECDSACertificateVerifierStorageCaller) LatestReferenceTimestamp(opts *bind.CallOpts, operatorSet OperatorSet) (uint32, error) {
	var out []interface{}
	err := _ECDSACertificateVerifierStorage.contract.Call(opts, &out, "latestReferenceTimestamp", operatorSet)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LatestReferenceTimestamp is a free data retrieval call binding the contract method 0x5ddb9b5b.
//
// Solidity: function latestReferenceTimestamp((address,uint32) operatorSet) view returns(uint32)
func (_ECDSACertificateVerifierStorage *ECDSACertificateVerifierStorageSession) LatestReferenceTimestamp(operatorSet OperatorSet) (uint32, error) {
	return _ECDSACertificateVerifierStorage.Contract.LatestReferenceTimestamp(&_ECDSACertificateVerifierStorage.CallOpts, operatorSet)
}

// LatestReferenceTimestamp is a free data retrieval call binding the contract method 0x5ddb9b5b.
//
// Solidity: function latestReferenceTimestamp((address,uint32) operatorSet) view returns(uint32)
func (_ECDSACertificateVerifierStorage *ECDSACertificateVerifierStorageCallerSession) LatestReferenceTimestamp(operatorSet OperatorSet) (uint32, error) {
	return _ECDSACertificateVerifierStorage.Contract.LatestReferenceTimestamp(&_ECDSACertificateVerifierStorage.CallOpts, operatorSet)
}

// MaxOperatorTableStaleness is a free data retrieval call binding the contract method 0x6141879e.
//
// Solidity: function maxOperatorTableStaleness((address,uint32) operatorSet) view returns(uint32)
func (_ECDSACertificateVerifierStorage *ECDSACertificateVerifierStorageCaller) MaxOperatorTableStaleness(opts *bind.CallOpts, operatorSet OperatorSet) (uint32, error) {
	var out []interface{}
	err := _ECDSACertificateVerifierStorage.contract.Call(opts, &out, "maxOperatorTableStaleness", operatorSet)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MaxOperatorTableStaleness is a free data retrieval call binding the contract method 0x6141879e.
//
// Solidity: function maxOperatorTableStaleness((address,uint32) operatorSet) view returns(uint32)
func (_ECDSACertificateVerifierStorage *ECDSACertificateVerifierStorageSession) MaxOperatorTableStaleness(operatorSet OperatorSet) (uint32, error) {
	return _ECDSACertificateVerifierStorage.Contract.MaxOperatorTableStaleness(&_ECDSACertificateVerifierStorage.CallOpts, operatorSet)
}

// MaxOperatorTableStaleness is a free data retrieval call binding the contract method 0x6141879e.
//
// Solidity: function maxOperatorTableStaleness((address,uint32) operatorSet) view returns(uint32)
func (_ECDSACertificateVerifierStorage *ECDSACertificateVerifierStorageCallerSession) MaxOperatorTableStaleness(operatorSet OperatorSet) (uint32, error) {
	return _ECDSACertificateVerifierStorage.Contract.MaxOperatorTableStaleness(&_ECDSACertificateVerifierStorage.CallOpts, operatorSet)
}

// OperatorTableUpdater is a free data retrieval call binding the contract method 0x68d6e081.
//
// Solidity: function operatorTableUpdater() view returns(address)
func (_ECDSACertificateVerifierStorage *ECDSACertificateVerifierStorageCaller) OperatorTableUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ECDSACertificateVerifierStorage.contract.Call(opts, &out, "operatorTableUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OperatorTableUpdater is a free data retrieval call binding the contract method 0x68d6e081.
//
// Solidity: function operatorTableUpdater() view returns(address)
func (_ECDSACertificateVerifierStorage *ECDSACertificateVerifierStorageSession) OperatorTableUpdater() (common.Address, error) {
	return _ECDSACertificateVerifierStorage.Contract.OperatorTableUpdater(&_ECDSACertificateVerifierStorage.CallOpts)
}

// OperatorTableUpdater is a free data retrieval call binding the contract method 0x68d6e081.
//
// Solidity: function operatorTableUpdater() view returns(address)
func (_ECDSACertificateVerifierStorage *ECDSACertificateVerifierStorageCallerSession) OperatorTableUpdater() (common.Address, error) {
	return _ECDSACertificateVerifierStorage.Contract.OperatorTableUpdater(&_ECDSACertificateVerifierStorage.CallOpts)
}

// UpdateOperatorTable is a paid mutator transaction binding the contract method 0x56d482f5.
//
// Solidity: function updateOperatorTable((address,uint32) operatorSet, uint32 referenceTimestamp, (address,uint256[])[] operatorInfos, (address,uint32) operatorSetConfig) returns()
func (_ECDSACertificateVerifierStorage *ECDSACertificateVerifierStorageTransactor) UpdateOperatorTable(opts *bind.TransactOpts, operatorSet OperatorSet, referenceTimestamp uint32, operatorInfos []IOperatorTableCalculatorTypesECDSAOperatorInfo, operatorSetConfig ICrossChainRegistryTypesOperatorSetConfig) (*types.Transaction, error) {
	return _ECDSACertificateVerifierStorage.contract.Transact(opts, "updateOperatorTable", operatorSet, referenceTimestamp, operatorInfos, operatorSetConfig)
}

// UpdateOperatorTable is a paid mutator transaction binding the contract method 0x56d482f5.
//
// Solidity: function updateOperatorTable((address,uint32) operatorSet, uint32 referenceTimestamp, (address,uint256[])[] operatorInfos, (address,uint32) operatorSetConfig) returns()
func (_ECDSACertificateVerifierStorage *ECDSACertificateVerifierStorageSession) UpdateOperatorTable(operatorSet OperatorSet, referenceTimestamp uint32, operatorInfos []IOperatorTableCalculatorTypesECDSAOperatorInfo, operatorSetConfig ICrossChainRegistryTypesOperatorSetConfig) (*types.Transaction, error) {
	return _ECDSACertificateVerifierStorage.Contract.UpdateOperatorTable(&_ECDSACertificateVerifierStorage.TransactOpts, operatorSet, referenceTimestamp, operatorInfos, operatorSetConfig)
}

// UpdateOperatorTable is a paid mutator transaction binding the contract method 0x56d482f5.
//
// Solidity: function updateOperatorTable((address,uint32) operatorSet, uint32 referenceTimestamp, (address,uint256[])[] operatorInfos, (address,uint32) operatorSetConfig) returns()
func (_ECDSACertificateVerifierStorage *ECDSACertificateVerifierStorageTransactorSession) UpdateOperatorTable(operatorSet OperatorSet, referenceTimestamp uint32, operatorInfos []IOperatorTableCalculatorTypesECDSAOperatorInfo, operatorSetConfig ICrossChainRegistryTypesOperatorSetConfig) (*types.Transaction, error) {
	return _ECDSACertificateVerifierStorage.Contract.UpdateOperatorTable(&_ECDSACertificateVerifierStorage.TransactOpts, operatorSet, referenceTimestamp, operatorInfos, operatorSetConfig)
}

// VerifyCertificate is a paid mutator transaction binding the contract method 0x80c7d3f3.
//
// Solidity: function verifyCertificate((address,uint32) operatorSet, (uint32,bytes32,bytes) cert) returns(uint256[] signedStakes, address[] signers)
func (_ECDSACertificateVerifierStorage *ECDSACertificateVerifierStorageTransactor) VerifyCertificate(opts *bind.TransactOpts, operatorSet OperatorSet, cert IECDSACertificateVerifierTypesECDSACertificate) (*types.Transaction, error) {
	return _ECDSACertificateVerifierStorage.contract.Transact(opts, "verifyCertificate", operatorSet, cert)
}

// VerifyCertificate is a paid mutator transaction binding the contract method 0x80c7d3f3.
//
// Solidity: function verifyCertificate((address,uint32) operatorSet, (uint32,bytes32,bytes) cert) returns(uint256[] signedStakes, address[] signers)
func (_ECDSACertificateVerifierStorage *ECDSACertificateVerifierStorageSession) VerifyCertificate(operatorSet OperatorSet, cert IECDSACertificateVerifierTypesECDSACertificate) (*types.Transaction, error) {
	return _ECDSACertificateVerifierStorage.Contract.VerifyCertificate(&_ECDSACertificateVerifierStorage.TransactOpts, operatorSet, cert)
}

// VerifyCertificate is a paid mutator transaction binding the contract method 0x80c7d3f3.
//
// Solidity: function verifyCertificate((address,uint32) operatorSet, (uint32,bytes32,bytes) cert) returns(uint256[] signedStakes, address[] signers)
func (_ECDSACertificateVerifierStorage *ECDSACertificateVerifierStorageTransactorSession) VerifyCertificate(operatorSet OperatorSet, cert IECDSACertificateVerifierTypesECDSACertificate) (*types.Transaction, error) {
	return _ECDSACertificateVerifierStorage.Contract.VerifyCertificate(&_ECDSACertificateVerifierStorage.TransactOpts, operatorSet, cert)
}

// VerifyCertificateNominal is a paid mutator transaction binding the contract method 0xbe86e0b2.
//
// Solidity: function verifyCertificateNominal((address,uint32) operatorSet, (uint32,bytes32,bytes) cert, uint256[] totalStakeNominalThresholds) returns(bool, address[] signers)
func (_ECDSACertificateVerifierStorage *ECDSACertificateVerifierStorageTransactor) VerifyCertificateNominal(opts *bind.TransactOpts, operatorSet OperatorSet, cert IECDSACertificateVerifierTypesECDSACertificate, totalStakeNominalThresholds []*big.Int) (*types.Transaction, error) {
	return _ECDSACertificateVerifierStorage.contract.Transact(opts, "verifyCertificateNominal", operatorSet, cert, totalStakeNominalThresholds)
}

// VerifyCertificateNominal is a paid mutator transaction binding the contract method 0xbe86e0b2.
//
// Solidity: function verifyCertificateNominal((address,uint32) operatorSet, (uint32,bytes32,bytes) cert, uint256[] totalStakeNominalThresholds) returns(bool, address[] signers)
func (_ECDSACertificateVerifierStorage *ECDSACertificateVerifierStorageSession) VerifyCertificateNominal(operatorSet OperatorSet, cert IECDSACertificateVerifierTypesECDSACertificate, totalStakeNominalThresholds []*big.Int) (*types.Transaction, error) {
	return _ECDSACertificateVerifierStorage.Contract.VerifyCertificateNominal(&_ECDSACertificateVerifierStorage.TransactOpts, operatorSet, cert, totalStakeNominalThresholds)
}

// VerifyCertificateNominal is a paid mutator transaction binding the contract method 0xbe86e0b2.
//
// Solidity: function verifyCertificateNominal((address,uint32) operatorSet, (uint32,bytes32,bytes) cert, uint256[] totalStakeNominalThresholds) returns(bool, address[] signers)
func (_ECDSACertificateVerifierStorage *ECDSACertificateVerifierStorageTransactorSession) VerifyCertificateNominal(operatorSet OperatorSet, cert IECDSACertificateVerifierTypesECDSACertificate, totalStakeNominalThresholds []*big.Int) (*types.Transaction, error) {
	return _ECDSACertificateVerifierStorage.Contract.VerifyCertificateNominal(&_ECDSACertificateVerifierStorage.TransactOpts, operatorSet, cert, totalStakeNominalThresholds)
}

// VerifyCertificateProportion is a paid mutator transaction binding the contract method 0xc0da2420.
//
// Solidity: function verifyCertificateProportion((address,uint32) operatorSet, (uint32,bytes32,bytes) cert, uint16[] totalStakeProportionThresholds) returns(bool, address[] signers)
func (_ECDSACertificateVerifierStorage *ECDSACertificateVerifierStorageTransactor) VerifyCertificateProportion(opts *bind.TransactOpts, operatorSet OperatorSet, cert IECDSACertificateVerifierTypesECDSACertificate, totalStakeProportionThresholds []uint16) (*types.Transaction, error) {
	return _ECDSACertificateVerifierStorage.contract.Transact(opts, "verifyCertificateProportion", operatorSet, cert, totalStakeProportionThresholds)
}

// VerifyCertificateProportion is a paid mutator transaction binding the contract method 0xc0da2420.
//
// Solidity: function verifyCertificateProportion((address,uint32) operatorSet, (uint32,bytes32,bytes) cert, uint16[] totalStakeProportionThresholds) returns(bool, address[] signers)
func (_ECDSACertificateVerifierStorage *ECDSACertificateVerifierStorageSession) VerifyCertificateProportion(operatorSet OperatorSet, cert IECDSACertificateVerifierTypesECDSACertificate, totalStakeProportionThresholds []uint16) (*types.Transaction, error) {
	return _ECDSACertificateVerifierStorage.Contract.VerifyCertificateProportion(&_ECDSACertificateVerifierStorage.TransactOpts, operatorSet, cert, totalStakeProportionThresholds)
}

// VerifyCertificateProportion is a paid mutator transaction binding the contract method 0xc0da2420.
//
// Solidity: function verifyCertificateProportion((address,uint32) operatorSet, (uint32,bytes32,bytes) cert, uint16[] totalStakeProportionThresholds) returns(bool, address[] signers)
func (_ECDSACertificateVerifierStorage *ECDSACertificateVerifierStorageTransactorSession) VerifyCertificateProportion(operatorSet OperatorSet, cert IECDSACertificateVerifierTypesECDSACertificate, totalStakeProportionThresholds []uint16) (*types.Transaction, error) {
	return _ECDSACertificateVerifierStorage.Contract.VerifyCertificateProportion(&_ECDSACertificateVerifierStorage.TransactOpts, operatorSet, cert, totalStakeProportionThresholds)
}

// ECDSACertificateVerifierStorageMaxStalenessPeriodUpdatedIterator is returned from FilterMaxStalenessPeriodUpdated and is used to iterate over the raw logs and unpacked data for MaxStalenessPeriodUpdated events raised by the ECDSACertificateVerifierStorage contract.
type ECDSACertificateVerifierStorageMaxStalenessPeriodUpdatedIterator struct {
	Event *ECDSACertificateVerifierStorageMaxStalenessPeriodUpdated // Event containing the contract specifics and raw log

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
func (it *ECDSACertificateVerifierStorageMaxStalenessPeriodUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ECDSACertificateVerifierStorageMaxStalenessPeriodUpdated)
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
		it.Event = new(ECDSACertificateVerifierStorageMaxStalenessPeriodUpdated)
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
func (it *ECDSACertificateVerifierStorageMaxStalenessPeriodUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ECDSACertificateVerifierStorageMaxStalenessPeriodUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ECDSACertificateVerifierStorageMaxStalenessPeriodUpdated represents a MaxStalenessPeriodUpdated event raised by the ECDSACertificateVerifierStorage contract.
type ECDSACertificateVerifierStorageMaxStalenessPeriodUpdated struct {
	OperatorSet        OperatorSet
	MaxStalenessPeriod uint32
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMaxStalenessPeriodUpdated is a free log retrieval operation binding the contract event 0x28539469fbbc8a5482e60966bf9376f7b9d25b2f0a65a9976f6baa3f0e3788da.
//
// Solidity: event MaxStalenessPeriodUpdated((address,uint32) operatorSet, uint32 maxStalenessPeriod)
func (_ECDSACertificateVerifierStorage *ECDSACertificateVerifierStorageFilterer) FilterMaxStalenessPeriodUpdated(opts *bind.FilterOpts) (*ECDSACertificateVerifierStorageMaxStalenessPeriodUpdatedIterator, error) {

	logs, sub, err := _ECDSACertificateVerifierStorage.contract.FilterLogs(opts, "MaxStalenessPeriodUpdated")
	if err != nil {
		return nil, err
	}
	return &ECDSACertificateVerifierStorageMaxStalenessPeriodUpdatedIterator{contract: _ECDSACertificateVerifierStorage.contract, event: "MaxStalenessPeriodUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxStalenessPeriodUpdated is a free log subscription operation binding the contract event 0x28539469fbbc8a5482e60966bf9376f7b9d25b2f0a65a9976f6baa3f0e3788da.
//
// Solidity: event MaxStalenessPeriodUpdated((address,uint32) operatorSet, uint32 maxStalenessPeriod)
func (_ECDSACertificateVerifierStorage *ECDSACertificateVerifierStorageFilterer) WatchMaxStalenessPeriodUpdated(opts *bind.WatchOpts, sink chan<- *ECDSACertificateVerifierStorageMaxStalenessPeriodUpdated) (event.Subscription, error) {

	logs, sub, err := _ECDSACertificateVerifierStorage.contract.WatchLogs(opts, "MaxStalenessPeriodUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ECDSACertificateVerifierStorageMaxStalenessPeriodUpdated)
				if err := _ECDSACertificateVerifierStorage.contract.UnpackLog(event, "MaxStalenessPeriodUpdated", log); err != nil {
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

// ParseMaxStalenessPeriodUpdated is a log parse operation binding the contract event 0x28539469fbbc8a5482e60966bf9376f7b9d25b2f0a65a9976f6baa3f0e3788da.
//
// Solidity: event MaxStalenessPeriodUpdated((address,uint32) operatorSet, uint32 maxStalenessPeriod)
func (_ECDSACertificateVerifierStorage *ECDSACertificateVerifierStorageFilterer) ParseMaxStalenessPeriodUpdated(log types.Log) (*ECDSACertificateVerifierStorageMaxStalenessPeriodUpdated, error) {
	event := new(ECDSACertificateVerifierStorageMaxStalenessPeriodUpdated)
	if err := _ECDSACertificateVerifierStorage.contract.UnpackLog(event, "MaxStalenessPeriodUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ECDSACertificateVerifierStorageOperatorSetOwnerUpdatedIterator is returned from FilterOperatorSetOwnerUpdated and is used to iterate over the raw logs and unpacked data for OperatorSetOwnerUpdated events raised by the ECDSACertificateVerifierStorage contract.
type ECDSACertificateVerifierStorageOperatorSetOwnerUpdatedIterator struct {
	Event *ECDSACertificateVerifierStorageOperatorSetOwnerUpdated // Event containing the contract specifics and raw log

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
func (it *ECDSACertificateVerifierStorageOperatorSetOwnerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ECDSACertificateVerifierStorageOperatorSetOwnerUpdated)
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
		it.Event = new(ECDSACertificateVerifierStorageOperatorSetOwnerUpdated)
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
func (it *ECDSACertificateVerifierStorageOperatorSetOwnerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ECDSACertificateVerifierStorageOperatorSetOwnerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ECDSACertificateVerifierStorageOperatorSetOwnerUpdated represents a OperatorSetOwnerUpdated event raised by the ECDSACertificateVerifierStorage contract.
type ECDSACertificateVerifierStorageOperatorSetOwnerUpdated struct {
	OperatorSet OperatorSet
	Owner       common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorSetOwnerUpdated is a free log retrieval operation binding the contract event 0x806dc367095c0baf953d7144b7c4376261675ee0b4e0da2761e43673051c7375.
//
// Solidity: event OperatorSetOwnerUpdated((address,uint32) operatorSet, address owner)
func (_ECDSACertificateVerifierStorage *ECDSACertificateVerifierStorageFilterer) FilterOperatorSetOwnerUpdated(opts *bind.FilterOpts) (*ECDSACertificateVerifierStorageOperatorSetOwnerUpdatedIterator, error) {

	logs, sub, err := _ECDSACertificateVerifierStorage.contract.FilterLogs(opts, "OperatorSetOwnerUpdated")
	if err != nil {
		return nil, err
	}
	return &ECDSACertificateVerifierStorageOperatorSetOwnerUpdatedIterator{contract: _ECDSACertificateVerifierStorage.contract, event: "OperatorSetOwnerUpdated", logs: logs, sub: sub}, nil
}

// WatchOperatorSetOwnerUpdated is a free log subscription operation binding the contract event 0x806dc367095c0baf953d7144b7c4376261675ee0b4e0da2761e43673051c7375.
//
// Solidity: event OperatorSetOwnerUpdated((address,uint32) operatorSet, address owner)
func (_ECDSACertificateVerifierStorage *ECDSACertificateVerifierStorageFilterer) WatchOperatorSetOwnerUpdated(opts *bind.WatchOpts, sink chan<- *ECDSACertificateVerifierStorageOperatorSetOwnerUpdated) (event.Subscription, error) {

	logs, sub, err := _ECDSACertificateVerifierStorage.contract.WatchLogs(opts, "OperatorSetOwnerUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ECDSACertificateVerifierStorageOperatorSetOwnerUpdated)
				if err := _ECDSACertificateVerifierStorage.contract.UnpackLog(event, "OperatorSetOwnerUpdated", log); err != nil {
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

// ParseOperatorSetOwnerUpdated is a log parse operation binding the contract event 0x806dc367095c0baf953d7144b7c4376261675ee0b4e0da2761e43673051c7375.
//
// Solidity: event OperatorSetOwnerUpdated((address,uint32) operatorSet, address owner)
func (_ECDSACertificateVerifierStorage *ECDSACertificateVerifierStorageFilterer) ParseOperatorSetOwnerUpdated(log types.Log) (*ECDSACertificateVerifierStorageOperatorSetOwnerUpdated, error) {
	event := new(ECDSACertificateVerifierStorageOperatorSetOwnerUpdated)
	if err := _ECDSACertificateVerifierStorage.contract.UnpackLog(event, "OperatorSetOwnerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ECDSACertificateVerifierStorageTableUpdatedIterator is returned from FilterTableUpdated and is used to iterate over the raw logs and unpacked data for TableUpdated events raised by the ECDSACertificateVerifierStorage contract.
type ECDSACertificateVerifierStorageTableUpdatedIterator struct {
	Event *ECDSACertificateVerifierStorageTableUpdated // Event containing the contract specifics and raw log

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
func (it *ECDSACertificateVerifierStorageTableUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ECDSACertificateVerifierStorageTableUpdated)
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
		it.Event = new(ECDSACertificateVerifierStorageTableUpdated)
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
func (it *ECDSACertificateVerifierStorageTableUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ECDSACertificateVerifierStorageTableUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ECDSACertificateVerifierStorageTableUpdated represents a TableUpdated event raised by the ECDSACertificateVerifierStorage contract.
type ECDSACertificateVerifierStorageTableUpdated struct {
	OperatorSet        OperatorSet
	ReferenceTimestamp uint32
	OperatorInfos      []IOperatorTableCalculatorTypesECDSAOperatorInfo
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterTableUpdated is a free log retrieval operation binding the contract event 0x4f588da9ec57976194a79b5594f8f8782923d93013df2b9ed12fe125805011ef.
//
// Solidity: event TableUpdated((address,uint32) operatorSet, uint32 referenceTimestamp, (address,uint256[])[] operatorInfos)
func (_ECDSACertificateVerifierStorage *ECDSACertificateVerifierStorageFilterer) FilterTableUpdated(opts *bind.FilterOpts) (*ECDSACertificateVerifierStorageTableUpdatedIterator, error) {

	logs, sub, err := _ECDSACertificateVerifierStorage.contract.FilterLogs(opts, "TableUpdated")
	if err != nil {
		return nil, err
	}
	return &ECDSACertificateVerifierStorageTableUpdatedIterator{contract: _ECDSACertificateVerifierStorage.contract, event: "TableUpdated", logs: logs, sub: sub}, nil
}

// WatchTableUpdated is a free log subscription operation binding the contract event 0x4f588da9ec57976194a79b5594f8f8782923d93013df2b9ed12fe125805011ef.
//
// Solidity: event TableUpdated((address,uint32) operatorSet, uint32 referenceTimestamp, (address,uint256[])[] operatorInfos)
func (_ECDSACertificateVerifierStorage *ECDSACertificateVerifierStorageFilterer) WatchTableUpdated(opts *bind.WatchOpts, sink chan<- *ECDSACertificateVerifierStorageTableUpdated) (event.Subscription, error) {

	logs, sub, err := _ECDSACertificateVerifierStorage.contract.WatchLogs(opts, "TableUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ECDSACertificateVerifierStorageTableUpdated)
				if err := _ECDSACertificateVerifierStorage.contract.UnpackLog(event, "TableUpdated", log); err != nil {
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

// ParseTableUpdated is a log parse operation binding the contract event 0x4f588da9ec57976194a79b5594f8f8782923d93013df2b9ed12fe125805011ef.
//
// Solidity: event TableUpdated((address,uint32) operatorSet, uint32 referenceTimestamp, (address,uint256[])[] operatorInfos)
func (_ECDSACertificateVerifierStorage *ECDSACertificateVerifierStorageFilterer) ParseTableUpdated(log types.Log) (*ECDSACertificateVerifierStorageTableUpdated, error) {
	event := new(ECDSACertificateVerifierStorageTableUpdated)
	if err := _ECDSACertificateVerifierStorage.contract.UnpackLog(event, "TableUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
