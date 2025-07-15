// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IECDSACertificateVerifier

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

// IECDSACertificateVerifierMetaData contains all meta data concerning the IECDSACertificateVerifier contract.
var IECDSACertificateVerifierMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"calculateCertificateDigest\",\"inputs\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"domainSeparator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorCount\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorInfo\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIOperatorTableCalculatorTypes.ECDSAOperatorInfo\",\"components\":[{\"name\":\"pubkey\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"weights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorInfos\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIOperatorTableCalculatorTypes.ECDSAOperatorInfo[]\",\"components\":[{\"name\":\"pubkey\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"weights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorSetOwner\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalStakes\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isReferenceTimestampSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestReferenceTimestamp\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxOperatorTableStaleness\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateOperatorTable\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorInfos\",\"type\":\"tuple[]\",\"internalType\":\"structIOperatorTableCalculatorTypes.ECDSAOperatorInfo[]\",\"components\":[{\"name\":\"pubkey\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"weights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]},{\"name\":\"operatorSetConfig\",\"type\":\"tuple\",\"internalType\":\"structICrossChainRegistryTypes.OperatorSetConfig\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxStalenessPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyCertificate\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"cert\",\"type\":\"tuple\",\"internalType\":\"structIECDSACertificateVerifierTypes.ECDSACertificate\",\"components\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sig\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"signedStakes\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"signers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyCertificateNominal\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"cert\",\"type\":\"tuple\",\"internalType\":\"structIECDSACertificateVerifierTypes.ECDSACertificate\",\"components\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sig\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"totalStakeNominalThresholds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"signers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyCertificateProportion\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"cert\",\"type\":\"tuple\",\"internalType\":\"structIECDSACertificateVerifierTypes.ECDSACertificate\",\"components\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sig\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"totalStakeProportionThresholds\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"signers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"MaxStalenessPeriodUpdated\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"maxStalenessPeriod\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSetOwnerUpdated\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TableUpdated\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"operatorInfos\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"structIOperatorTableCalculatorTypes.ECDSAOperatorInfo[]\",\"components\":[{\"name\":\"pubkey\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"weights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ArrayLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CertificateStale\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignatureLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyTableUpdater\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReferenceTimestampDoesNotExist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RootDisabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignersNotOrdered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TableUpdateStale\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"VerificationFailed\",\"inputs\":[]}]",
}

// IECDSACertificateVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use IECDSACertificateVerifierMetaData.ABI instead.
var IECDSACertificateVerifierABI = IECDSACertificateVerifierMetaData.ABI

// IECDSACertificateVerifier is an auto generated Go binding around an Ethereum contract.
type IECDSACertificateVerifier struct {
	IECDSACertificateVerifierCaller     // Read-only binding to the contract
	IECDSACertificateVerifierTransactor // Write-only binding to the contract
	IECDSACertificateVerifierFilterer   // Log filterer for contract events
}

// IECDSACertificateVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type IECDSACertificateVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IECDSACertificateVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IECDSACertificateVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IECDSACertificateVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IECDSACertificateVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IECDSACertificateVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IECDSACertificateVerifierSession struct {
	Contract     *IECDSACertificateVerifier // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IECDSACertificateVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IECDSACertificateVerifierCallerSession struct {
	Contract *IECDSACertificateVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// IECDSACertificateVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IECDSACertificateVerifierTransactorSession struct {
	Contract     *IECDSACertificateVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// IECDSACertificateVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type IECDSACertificateVerifierRaw struct {
	Contract *IECDSACertificateVerifier // Generic contract binding to access the raw methods on
}

// IECDSACertificateVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IECDSACertificateVerifierCallerRaw struct {
	Contract *IECDSACertificateVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// IECDSACertificateVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IECDSACertificateVerifierTransactorRaw struct {
	Contract *IECDSACertificateVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIECDSACertificateVerifier creates a new instance of IECDSACertificateVerifier, bound to a specific deployed contract.
func NewIECDSACertificateVerifier(address common.Address, backend bind.ContractBackend) (*IECDSACertificateVerifier, error) {
	contract, err := bindIECDSACertificateVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IECDSACertificateVerifier{IECDSACertificateVerifierCaller: IECDSACertificateVerifierCaller{contract: contract}, IECDSACertificateVerifierTransactor: IECDSACertificateVerifierTransactor{contract: contract}, IECDSACertificateVerifierFilterer: IECDSACertificateVerifierFilterer{contract: contract}}, nil
}

// NewIECDSACertificateVerifierCaller creates a new read-only instance of IECDSACertificateVerifier, bound to a specific deployed contract.
func NewIECDSACertificateVerifierCaller(address common.Address, caller bind.ContractCaller) (*IECDSACertificateVerifierCaller, error) {
	contract, err := bindIECDSACertificateVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IECDSACertificateVerifierCaller{contract: contract}, nil
}

// NewIECDSACertificateVerifierTransactor creates a new write-only instance of IECDSACertificateVerifier, bound to a specific deployed contract.
func NewIECDSACertificateVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*IECDSACertificateVerifierTransactor, error) {
	contract, err := bindIECDSACertificateVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IECDSACertificateVerifierTransactor{contract: contract}, nil
}

// NewIECDSACertificateVerifierFilterer creates a new log filterer instance of IECDSACertificateVerifier, bound to a specific deployed contract.
func NewIECDSACertificateVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*IECDSACertificateVerifierFilterer, error) {
	contract, err := bindIECDSACertificateVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IECDSACertificateVerifierFilterer{contract: contract}, nil
}

// bindIECDSACertificateVerifier binds a generic wrapper to an already deployed contract.
func bindIECDSACertificateVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IECDSACertificateVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IECDSACertificateVerifier *IECDSACertificateVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IECDSACertificateVerifier.Contract.IECDSACertificateVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IECDSACertificateVerifier *IECDSACertificateVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IECDSACertificateVerifier.Contract.IECDSACertificateVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IECDSACertificateVerifier *IECDSACertificateVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IECDSACertificateVerifier.Contract.IECDSACertificateVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IECDSACertificateVerifier *IECDSACertificateVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IECDSACertificateVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IECDSACertificateVerifier *IECDSACertificateVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IECDSACertificateVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IECDSACertificateVerifier *IECDSACertificateVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IECDSACertificateVerifier.Contract.contract.Transact(opts, method, params...)
}

// CalculateCertificateDigest is a free data retrieval call binding the contract method 0x18467434.
//
// Solidity: function calculateCertificateDigest(uint32 referenceTimestamp, bytes32 messageHash) view returns(bytes32)
func (_IECDSACertificateVerifier *IECDSACertificateVerifierCaller) CalculateCertificateDigest(opts *bind.CallOpts, referenceTimestamp uint32, messageHash [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _IECDSACertificateVerifier.contract.Call(opts, &out, "calculateCertificateDigest", referenceTimestamp, messageHash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateCertificateDigest is a free data retrieval call binding the contract method 0x18467434.
//
// Solidity: function calculateCertificateDigest(uint32 referenceTimestamp, bytes32 messageHash) view returns(bytes32)
func (_IECDSACertificateVerifier *IECDSACertificateVerifierSession) CalculateCertificateDigest(referenceTimestamp uint32, messageHash [32]byte) ([32]byte, error) {
	return _IECDSACertificateVerifier.Contract.CalculateCertificateDigest(&_IECDSACertificateVerifier.CallOpts, referenceTimestamp, messageHash)
}

// CalculateCertificateDigest is a free data retrieval call binding the contract method 0x18467434.
//
// Solidity: function calculateCertificateDigest(uint32 referenceTimestamp, bytes32 messageHash) view returns(bytes32)
func (_IECDSACertificateVerifier *IECDSACertificateVerifierCallerSession) CalculateCertificateDigest(referenceTimestamp uint32, messageHash [32]byte) ([32]byte, error) {
	return _IECDSACertificateVerifier.Contract.CalculateCertificateDigest(&_IECDSACertificateVerifier.CallOpts, referenceTimestamp, messageHash)
}

// CalculateCertificateDigestBytes is a free data retrieval call binding the contract method 0x702ca531.
//
// Solidity: function calculateCertificateDigestBytes(uint32 referenceTimestamp, bytes32 messageHash) view returns(bytes)
func (_IECDSACertificateVerifier *IECDSACertificateVerifierCaller) CalculateCertificateDigestBytes(opts *bind.CallOpts, referenceTimestamp uint32, messageHash [32]byte) ([]byte, error) {
	var out []interface{}
	err := _IECDSACertificateVerifier.contract.Call(opts, &out, "calculateCertificateDigestBytes", referenceTimestamp, messageHash)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// CalculateCertificateDigestBytes is a free data retrieval call binding the contract method 0x702ca531.
//
// Solidity: function calculateCertificateDigestBytes(uint32 referenceTimestamp, bytes32 messageHash) view returns(bytes)
func (_IECDSACertificateVerifier *IECDSACertificateVerifierSession) CalculateCertificateDigestBytes(referenceTimestamp uint32, messageHash [32]byte) ([]byte, error) {
	return _IECDSACertificateVerifier.Contract.CalculateCertificateDigestBytes(&_IECDSACertificateVerifier.CallOpts, referenceTimestamp, messageHash)
}

// CalculateCertificateDigestBytes is a free data retrieval call binding the contract method 0x702ca531.
//
// Solidity: function calculateCertificateDigestBytes(uint32 referenceTimestamp, bytes32 messageHash) view returns(bytes)
func (_IECDSACertificateVerifier *IECDSACertificateVerifierCallerSession) CalculateCertificateDigestBytes(referenceTimestamp uint32, messageHash [32]byte) ([]byte, error) {
	return _IECDSACertificateVerifier.Contract.CalculateCertificateDigestBytes(&_IECDSACertificateVerifier.CallOpts, referenceTimestamp, messageHash)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_IECDSACertificateVerifier *IECDSACertificateVerifierCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IECDSACertificateVerifier.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_IECDSACertificateVerifier *IECDSACertificateVerifierSession) DomainSeparator() ([32]byte, error) {
	return _IECDSACertificateVerifier.Contract.DomainSeparator(&_IECDSACertificateVerifier.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_IECDSACertificateVerifier *IECDSACertificateVerifierCallerSession) DomainSeparator() ([32]byte, error) {
	return _IECDSACertificateVerifier.Contract.DomainSeparator(&_IECDSACertificateVerifier.CallOpts)
}

// GetOperatorCount is a free data retrieval call binding the contract method 0x23c2a3cb.
//
// Solidity: function getOperatorCount((address,uint32) operatorSet, uint32 referenceTimestamp) view returns(uint32)
func (_IECDSACertificateVerifier *IECDSACertificateVerifierCaller) GetOperatorCount(opts *bind.CallOpts, operatorSet OperatorSet, referenceTimestamp uint32) (uint32, error) {
	var out []interface{}
	err := _IECDSACertificateVerifier.contract.Call(opts, &out, "getOperatorCount", operatorSet, referenceTimestamp)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetOperatorCount is a free data retrieval call binding the contract method 0x23c2a3cb.
//
// Solidity: function getOperatorCount((address,uint32) operatorSet, uint32 referenceTimestamp) view returns(uint32)
func (_IECDSACertificateVerifier *IECDSACertificateVerifierSession) GetOperatorCount(operatorSet OperatorSet, referenceTimestamp uint32) (uint32, error) {
	return _IECDSACertificateVerifier.Contract.GetOperatorCount(&_IECDSACertificateVerifier.CallOpts, operatorSet, referenceTimestamp)
}

// GetOperatorCount is a free data retrieval call binding the contract method 0x23c2a3cb.
//
// Solidity: function getOperatorCount((address,uint32) operatorSet, uint32 referenceTimestamp) view returns(uint32)
func (_IECDSACertificateVerifier *IECDSACertificateVerifierCallerSession) GetOperatorCount(operatorSet OperatorSet, referenceTimestamp uint32) (uint32, error) {
	return _IECDSACertificateVerifier.Contract.GetOperatorCount(&_IECDSACertificateVerifier.CallOpts, operatorSet, referenceTimestamp)
}

// GetOperatorInfo is a free data retrieval call binding the contract method 0xe49613fc.
//
// Solidity: function getOperatorInfo((address,uint32) operatorSet, uint32 referenceTimestamp, uint256 operatorIndex) view returns((address,uint256[]))
func (_IECDSACertificateVerifier *IECDSACertificateVerifierCaller) GetOperatorInfo(opts *bind.CallOpts, operatorSet OperatorSet, referenceTimestamp uint32, operatorIndex *big.Int) (IOperatorTableCalculatorTypesECDSAOperatorInfo, error) {
	var out []interface{}
	err := _IECDSACertificateVerifier.contract.Call(opts, &out, "getOperatorInfo", operatorSet, referenceTimestamp, operatorIndex)

	if err != nil {
		return *new(IOperatorTableCalculatorTypesECDSAOperatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IOperatorTableCalculatorTypesECDSAOperatorInfo)).(*IOperatorTableCalculatorTypesECDSAOperatorInfo)

	return out0, err

}

// GetOperatorInfo is a free data retrieval call binding the contract method 0xe49613fc.
//
// Solidity: function getOperatorInfo((address,uint32) operatorSet, uint32 referenceTimestamp, uint256 operatorIndex) view returns((address,uint256[]))
func (_IECDSACertificateVerifier *IECDSACertificateVerifierSession) GetOperatorInfo(operatorSet OperatorSet, referenceTimestamp uint32, operatorIndex *big.Int) (IOperatorTableCalculatorTypesECDSAOperatorInfo, error) {
	return _IECDSACertificateVerifier.Contract.GetOperatorInfo(&_IECDSACertificateVerifier.CallOpts, operatorSet, referenceTimestamp, operatorIndex)
}

// GetOperatorInfo is a free data retrieval call binding the contract method 0xe49613fc.
//
// Solidity: function getOperatorInfo((address,uint32) operatorSet, uint32 referenceTimestamp, uint256 operatorIndex) view returns((address,uint256[]))
func (_IECDSACertificateVerifier *IECDSACertificateVerifierCallerSession) GetOperatorInfo(operatorSet OperatorSet, referenceTimestamp uint32, operatorIndex *big.Int) (IOperatorTableCalculatorTypesECDSAOperatorInfo, error) {
	return _IECDSACertificateVerifier.Contract.GetOperatorInfo(&_IECDSACertificateVerifier.CallOpts, operatorSet, referenceTimestamp, operatorIndex)
}

// GetOperatorInfos is a free data retrieval call binding the contract method 0x7c85ac4c.
//
// Solidity: function getOperatorInfos((address,uint32) operatorSet, uint32 referenceTimestamp) view returns((address,uint256[])[])
func (_IECDSACertificateVerifier *IECDSACertificateVerifierCaller) GetOperatorInfos(opts *bind.CallOpts, operatorSet OperatorSet, referenceTimestamp uint32) ([]IOperatorTableCalculatorTypesECDSAOperatorInfo, error) {
	var out []interface{}
	err := _IECDSACertificateVerifier.contract.Call(opts, &out, "getOperatorInfos", operatorSet, referenceTimestamp)

	if err != nil {
		return *new([]IOperatorTableCalculatorTypesECDSAOperatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IOperatorTableCalculatorTypesECDSAOperatorInfo)).(*[]IOperatorTableCalculatorTypesECDSAOperatorInfo)

	return out0, err

}

// GetOperatorInfos is a free data retrieval call binding the contract method 0x7c85ac4c.
//
// Solidity: function getOperatorInfos((address,uint32) operatorSet, uint32 referenceTimestamp) view returns((address,uint256[])[])
func (_IECDSACertificateVerifier *IECDSACertificateVerifierSession) GetOperatorInfos(operatorSet OperatorSet, referenceTimestamp uint32) ([]IOperatorTableCalculatorTypesECDSAOperatorInfo, error) {
	return _IECDSACertificateVerifier.Contract.GetOperatorInfos(&_IECDSACertificateVerifier.CallOpts, operatorSet, referenceTimestamp)
}

// GetOperatorInfos is a free data retrieval call binding the contract method 0x7c85ac4c.
//
// Solidity: function getOperatorInfos((address,uint32) operatorSet, uint32 referenceTimestamp) view returns((address,uint256[])[])
func (_IECDSACertificateVerifier *IECDSACertificateVerifierCallerSession) GetOperatorInfos(operatorSet OperatorSet, referenceTimestamp uint32) ([]IOperatorTableCalculatorTypesECDSAOperatorInfo, error) {
	return _IECDSACertificateVerifier.Contract.GetOperatorInfos(&_IECDSACertificateVerifier.CallOpts, operatorSet, referenceTimestamp)
}

// GetOperatorSetOwner is a free data retrieval call binding the contract method 0x84818920.
//
// Solidity: function getOperatorSetOwner((address,uint32) operatorSet) view returns(address)
func (_IECDSACertificateVerifier *IECDSACertificateVerifierCaller) GetOperatorSetOwner(opts *bind.CallOpts, operatorSet OperatorSet) (common.Address, error) {
	var out []interface{}
	err := _IECDSACertificateVerifier.contract.Call(opts, &out, "getOperatorSetOwner", operatorSet)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOperatorSetOwner is a free data retrieval call binding the contract method 0x84818920.
//
// Solidity: function getOperatorSetOwner((address,uint32) operatorSet) view returns(address)
func (_IECDSACertificateVerifier *IECDSACertificateVerifierSession) GetOperatorSetOwner(operatorSet OperatorSet) (common.Address, error) {
	return _IECDSACertificateVerifier.Contract.GetOperatorSetOwner(&_IECDSACertificateVerifier.CallOpts, operatorSet)
}

// GetOperatorSetOwner is a free data retrieval call binding the contract method 0x84818920.
//
// Solidity: function getOperatorSetOwner((address,uint32) operatorSet) view returns(address)
func (_IECDSACertificateVerifier *IECDSACertificateVerifierCallerSession) GetOperatorSetOwner(operatorSet OperatorSet) (common.Address, error) {
	return _IECDSACertificateVerifier.Contract.GetOperatorSetOwner(&_IECDSACertificateVerifier.CallOpts, operatorSet)
}

// GetTotalStakes is a free data retrieval call binding the contract method 0x04cdbae4.
//
// Solidity: function getTotalStakes((address,uint32) operatorSet, uint32 referenceTimestamp) view returns(uint256[])
func (_IECDSACertificateVerifier *IECDSACertificateVerifierCaller) GetTotalStakes(opts *bind.CallOpts, operatorSet OperatorSet, referenceTimestamp uint32) ([]*big.Int, error) {
	var out []interface{}
	err := _IECDSACertificateVerifier.contract.Call(opts, &out, "getTotalStakes", operatorSet, referenceTimestamp)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetTotalStakes is a free data retrieval call binding the contract method 0x04cdbae4.
//
// Solidity: function getTotalStakes((address,uint32) operatorSet, uint32 referenceTimestamp) view returns(uint256[])
func (_IECDSACertificateVerifier *IECDSACertificateVerifierSession) GetTotalStakes(operatorSet OperatorSet, referenceTimestamp uint32) ([]*big.Int, error) {
	return _IECDSACertificateVerifier.Contract.GetTotalStakes(&_IECDSACertificateVerifier.CallOpts, operatorSet, referenceTimestamp)
}

// GetTotalStakes is a free data retrieval call binding the contract method 0x04cdbae4.
//
// Solidity: function getTotalStakes((address,uint32) operatorSet, uint32 referenceTimestamp) view returns(uint256[])
func (_IECDSACertificateVerifier *IECDSACertificateVerifierCallerSession) GetTotalStakes(operatorSet OperatorSet, referenceTimestamp uint32) ([]*big.Int, error) {
	return _IECDSACertificateVerifier.Contract.GetTotalStakes(&_IECDSACertificateVerifier.CallOpts, operatorSet, referenceTimestamp)
}

// IsReferenceTimestampSet is a free data retrieval call binding the contract method 0xcd83a72b.
//
// Solidity: function isReferenceTimestampSet((address,uint32) operatorSet, uint32 referenceTimestamp) view returns(bool)
func (_IECDSACertificateVerifier *IECDSACertificateVerifierCaller) IsReferenceTimestampSet(opts *bind.CallOpts, operatorSet OperatorSet, referenceTimestamp uint32) (bool, error) {
	var out []interface{}
	err := _IECDSACertificateVerifier.contract.Call(opts, &out, "isReferenceTimestampSet", operatorSet, referenceTimestamp)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsReferenceTimestampSet is a free data retrieval call binding the contract method 0xcd83a72b.
//
// Solidity: function isReferenceTimestampSet((address,uint32) operatorSet, uint32 referenceTimestamp) view returns(bool)
func (_IECDSACertificateVerifier *IECDSACertificateVerifierSession) IsReferenceTimestampSet(operatorSet OperatorSet, referenceTimestamp uint32) (bool, error) {
	return _IECDSACertificateVerifier.Contract.IsReferenceTimestampSet(&_IECDSACertificateVerifier.CallOpts, operatorSet, referenceTimestamp)
}

// IsReferenceTimestampSet is a free data retrieval call binding the contract method 0xcd83a72b.
//
// Solidity: function isReferenceTimestampSet((address,uint32) operatorSet, uint32 referenceTimestamp) view returns(bool)
func (_IECDSACertificateVerifier *IECDSACertificateVerifierCallerSession) IsReferenceTimestampSet(operatorSet OperatorSet, referenceTimestamp uint32) (bool, error) {
	return _IECDSACertificateVerifier.Contract.IsReferenceTimestampSet(&_IECDSACertificateVerifier.CallOpts, operatorSet, referenceTimestamp)
}

// LatestReferenceTimestamp is a free data retrieval call binding the contract method 0x5ddb9b5b.
//
// Solidity: function latestReferenceTimestamp((address,uint32) operatorSet) view returns(uint32)
func (_IECDSACertificateVerifier *IECDSACertificateVerifierCaller) LatestReferenceTimestamp(opts *bind.CallOpts, operatorSet OperatorSet) (uint32, error) {
	var out []interface{}
	err := _IECDSACertificateVerifier.contract.Call(opts, &out, "latestReferenceTimestamp", operatorSet)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LatestReferenceTimestamp is a free data retrieval call binding the contract method 0x5ddb9b5b.
//
// Solidity: function latestReferenceTimestamp((address,uint32) operatorSet) view returns(uint32)
func (_IECDSACertificateVerifier *IECDSACertificateVerifierSession) LatestReferenceTimestamp(operatorSet OperatorSet) (uint32, error) {
	return _IECDSACertificateVerifier.Contract.LatestReferenceTimestamp(&_IECDSACertificateVerifier.CallOpts, operatorSet)
}

// LatestReferenceTimestamp is a free data retrieval call binding the contract method 0x5ddb9b5b.
//
// Solidity: function latestReferenceTimestamp((address,uint32) operatorSet) view returns(uint32)
func (_IECDSACertificateVerifier *IECDSACertificateVerifierCallerSession) LatestReferenceTimestamp(operatorSet OperatorSet) (uint32, error) {
	return _IECDSACertificateVerifier.Contract.LatestReferenceTimestamp(&_IECDSACertificateVerifier.CallOpts, operatorSet)
}

// MaxOperatorTableStaleness is a free data retrieval call binding the contract method 0x6141879e.
//
// Solidity: function maxOperatorTableStaleness((address,uint32) operatorSet) view returns(uint32)
func (_IECDSACertificateVerifier *IECDSACertificateVerifierCaller) MaxOperatorTableStaleness(opts *bind.CallOpts, operatorSet OperatorSet) (uint32, error) {
	var out []interface{}
	err := _IECDSACertificateVerifier.contract.Call(opts, &out, "maxOperatorTableStaleness", operatorSet)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MaxOperatorTableStaleness is a free data retrieval call binding the contract method 0x6141879e.
//
// Solidity: function maxOperatorTableStaleness((address,uint32) operatorSet) view returns(uint32)
func (_IECDSACertificateVerifier *IECDSACertificateVerifierSession) MaxOperatorTableStaleness(operatorSet OperatorSet) (uint32, error) {
	return _IECDSACertificateVerifier.Contract.MaxOperatorTableStaleness(&_IECDSACertificateVerifier.CallOpts, operatorSet)
}

// MaxOperatorTableStaleness is a free data retrieval call binding the contract method 0x6141879e.
//
// Solidity: function maxOperatorTableStaleness((address,uint32) operatorSet) view returns(uint32)
func (_IECDSACertificateVerifier *IECDSACertificateVerifierCallerSession) MaxOperatorTableStaleness(operatorSet OperatorSet) (uint32, error) {
	return _IECDSACertificateVerifier.Contract.MaxOperatorTableStaleness(&_IECDSACertificateVerifier.CallOpts, operatorSet)
}

// VerifyCertificate is a free data retrieval call binding the contract method 0x80c7d3f3.
//
// Solidity: function verifyCertificate((address,uint32) operatorSet, (uint32,bytes32,bytes) cert) view returns(uint256[] signedStakes, address[] signers)
func (_IECDSACertificateVerifier *IECDSACertificateVerifierCaller) VerifyCertificate(opts *bind.CallOpts, operatorSet OperatorSet, cert IECDSACertificateVerifierTypesECDSACertificate) (struct {
	SignedStakes []*big.Int
	Signers      []common.Address
}, error) {
	var out []interface{}
	err := _IECDSACertificateVerifier.contract.Call(opts, &out, "verifyCertificate", operatorSet, cert)

	outstruct := new(struct {
		SignedStakes []*big.Int
		Signers      []common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SignedStakes = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.Signers = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)

	return *outstruct, err

}

// VerifyCertificate is a free data retrieval call binding the contract method 0x80c7d3f3.
//
// Solidity: function verifyCertificate((address,uint32) operatorSet, (uint32,bytes32,bytes) cert) view returns(uint256[] signedStakes, address[] signers)
func (_IECDSACertificateVerifier *IECDSACertificateVerifierSession) VerifyCertificate(operatorSet OperatorSet, cert IECDSACertificateVerifierTypesECDSACertificate) (struct {
	SignedStakes []*big.Int
	Signers      []common.Address
}, error) {
	return _IECDSACertificateVerifier.Contract.VerifyCertificate(&_IECDSACertificateVerifier.CallOpts, operatorSet, cert)
}

// VerifyCertificate is a free data retrieval call binding the contract method 0x80c7d3f3.
//
// Solidity: function verifyCertificate((address,uint32) operatorSet, (uint32,bytes32,bytes) cert) view returns(uint256[] signedStakes, address[] signers)
func (_IECDSACertificateVerifier *IECDSACertificateVerifierCallerSession) VerifyCertificate(operatorSet OperatorSet, cert IECDSACertificateVerifierTypesECDSACertificate) (struct {
	SignedStakes []*big.Int
	Signers      []common.Address
}, error) {
	return _IECDSACertificateVerifier.Contract.VerifyCertificate(&_IECDSACertificateVerifier.CallOpts, operatorSet, cert)
}

// VerifyCertificateNominal is a free data retrieval call binding the contract method 0xbe86e0b2.
//
// Solidity: function verifyCertificateNominal((address,uint32) operatorSet, (uint32,bytes32,bytes) cert, uint256[] totalStakeNominalThresholds) view returns(bool, address[] signers)
func (_IECDSACertificateVerifier *IECDSACertificateVerifierCaller) VerifyCertificateNominal(opts *bind.CallOpts, operatorSet OperatorSet, cert IECDSACertificateVerifierTypesECDSACertificate, totalStakeNominalThresholds []*big.Int) (bool, []common.Address, error) {
	var out []interface{}
	err := _IECDSACertificateVerifier.contract.Call(opts, &out, "verifyCertificateNominal", operatorSet, cert, totalStakeNominalThresholds)

	if err != nil {
		return *new(bool), *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)

	return out0, out1, err

}

// VerifyCertificateNominal is a free data retrieval call binding the contract method 0xbe86e0b2.
//
// Solidity: function verifyCertificateNominal((address,uint32) operatorSet, (uint32,bytes32,bytes) cert, uint256[] totalStakeNominalThresholds) view returns(bool, address[] signers)
func (_IECDSACertificateVerifier *IECDSACertificateVerifierSession) VerifyCertificateNominal(operatorSet OperatorSet, cert IECDSACertificateVerifierTypesECDSACertificate, totalStakeNominalThresholds []*big.Int) (bool, []common.Address, error) {
	return _IECDSACertificateVerifier.Contract.VerifyCertificateNominal(&_IECDSACertificateVerifier.CallOpts, operatorSet, cert, totalStakeNominalThresholds)
}

// VerifyCertificateNominal is a free data retrieval call binding the contract method 0xbe86e0b2.
//
// Solidity: function verifyCertificateNominal((address,uint32) operatorSet, (uint32,bytes32,bytes) cert, uint256[] totalStakeNominalThresholds) view returns(bool, address[] signers)
func (_IECDSACertificateVerifier *IECDSACertificateVerifierCallerSession) VerifyCertificateNominal(operatorSet OperatorSet, cert IECDSACertificateVerifierTypesECDSACertificate, totalStakeNominalThresholds []*big.Int) (bool, []common.Address, error) {
	return _IECDSACertificateVerifier.Contract.VerifyCertificateNominal(&_IECDSACertificateVerifier.CallOpts, operatorSet, cert, totalStakeNominalThresholds)
}

// VerifyCertificateProportion is a free data retrieval call binding the contract method 0xc0da2420.
//
// Solidity: function verifyCertificateProportion((address,uint32) operatorSet, (uint32,bytes32,bytes) cert, uint16[] totalStakeProportionThresholds) view returns(bool, address[] signers)
func (_IECDSACertificateVerifier *IECDSACertificateVerifierCaller) VerifyCertificateProportion(opts *bind.CallOpts, operatorSet OperatorSet, cert IECDSACertificateVerifierTypesECDSACertificate, totalStakeProportionThresholds []uint16) (bool, []common.Address, error) {
	var out []interface{}
	err := _IECDSACertificateVerifier.contract.Call(opts, &out, "verifyCertificateProportion", operatorSet, cert, totalStakeProportionThresholds)

	if err != nil {
		return *new(bool), *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)

	return out0, out1, err

}

// VerifyCertificateProportion is a free data retrieval call binding the contract method 0xc0da2420.
//
// Solidity: function verifyCertificateProportion((address,uint32) operatorSet, (uint32,bytes32,bytes) cert, uint16[] totalStakeProportionThresholds) view returns(bool, address[] signers)
func (_IECDSACertificateVerifier *IECDSACertificateVerifierSession) VerifyCertificateProportion(operatorSet OperatorSet, cert IECDSACertificateVerifierTypesECDSACertificate, totalStakeProportionThresholds []uint16) (bool, []common.Address, error) {
	return _IECDSACertificateVerifier.Contract.VerifyCertificateProportion(&_IECDSACertificateVerifier.CallOpts, operatorSet, cert, totalStakeProportionThresholds)
}

// VerifyCertificateProportion is a free data retrieval call binding the contract method 0xc0da2420.
//
// Solidity: function verifyCertificateProportion((address,uint32) operatorSet, (uint32,bytes32,bytes) cert, uint16[] totalStakeProportionThresholds) view returns(bool, address[] signers)
func (_IECDSACertificateVerifier *IECDSACertificateVerifierCallerSession) VerifyCertificateProportion(operatorSet OperatorSet, cert IECDSACertificateVerifierTypesECDSACertificate, totalStakeProportionThresholds []uint16) (bool, []common.Address, error) {
	return _IECDSACertificateVerifier.Contract.VerifyCertificateProportion(&_IECDSACertificateVerifier.CallOpts, operatorSet, cert, totalStakeProportionThresholds)
}

// UpdateOperatorTable is a paid mutator transaction binding the contract method 0x56d482f5.
//
// Solidity: function updateOperatorTable((address,uint32) operatorSet, uint32 referenceTimestamp, (address,uint256[])[] operatorInfos, (address,uint32) operatorSetConfig) returns()
func (_IECDSACertificateVerifier *IECDSACertificateVerifierTransactor) UpdateOperatorTable(opts *bind.TransactOpts, operatorSet OperatorSet, referenceTimestamp uint32, operatorInfos []IOperatorTableCalculatorTypesECDSAOperatorInfo, operatorSetConfig ICrossChainRegistryTypesOperatorSetConfig) (*types.Transaction, error) {
	return _IECDSACertificateVerifier.contract.Transact(opts, "updateOperatorTable", operatorSet, referenceTimestamp, operatorInfos, operatorSetConfig)
}

// UpdateOperatorTable is a paid mutator transaction binding the contract method 0x56d482f5.
//
// Solidity: function updateOperatorTable((address,uint32) operatorSet, uint32 referenceTimestamp, (address,uint256[])[] operatorInfos, (address,uint32) operatorSetConfig) returns()
func (_IECDSACertificateVerifier *IECDSACertificateVerifierSession) UpdateOperatorTable(operatorSet OperatorSet, referenceTimestamp uint32, operatorInfos []IOperatorTableCalculatorTypesECDSAOperatorInfo, operatorSetConfig ICrossChainRegistryTypesOperatorSetConfig) (*types.Transaction, error) {
	return _IECDSACertificateVerifier.Contract.UpdateOperatorTable(&_IECDSACertificateVerifier.TransactOpts, operatorSet, referenceTimestamp, operatorInfos, operatorSetConfig)
}

// UpdateOperatorTable is a paid mutator transaction binding the contract method 0x56d482f5.
//
// Solidity: function updateOperatorTable((address,uint32) operatorSet, uint32 referenceTimestamp, (address,uint256[])[] operatorInfos, (address,uint32) operatorSetConfig) returns()
func (_IECDSACertificateVerifier *IECDSACertificateVerifierTransactorSession) UpdateOperatorTable(operatorSet OperatorSet, referenceTimestamp uint32, operatorInfos []IOperatorTableCalculatorTypesECDSAOperatorInfo, operatorSetConfig ICrossChainRegistryTypesOperatorSetConfig) (*types.Transaction, error) {
	return _IECDSACertificateVerifier.Contract.UpdateOperatorTable(&_IECDSACertificateVerifier.TransactOpts, operatorSet, referenceTimestamp, operatorInfos, operatorSetConfig)
}

// IECDSACertificateVerifierMaxStalenessPeriodUpdatedIterator is returned from FilterMaxStalenessPeriodUpdated and is used to iterate over the raw logs and unpacked data for MaxStalenessPeriodUpdated events raised by the IECDSACertificateVerifier contract.
type IECDSACertificateVerifierMaxStalenessPeriodUpdatedIterator struct {
	Event *IECDSACertificateVerifierMaxStalenessPeriodUpdated // Event containing the contract specifics and raw log

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
func (it *IECDSACertificateVerifierMaxStalenessPeriodUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IECDSACertificateVerifierMaxStalenessPeriodUpdated)
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
		it.Event = new(IECDSACertificateVerifierMaxStalenessPeriodUpdated)
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
func (it *IECDSACertificateVerifierMaxStalenessPeriodUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IECDSACertificateVerifierMaxStalenessPeriodUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IECDSACertificateVerifierMaxStalenessPeriodUpdated represents a MaxStalenessPeriodUpdated event raised by the IECDSACertificateVerifier contract.
type IECDSACertificateVerifierMaxStalenessPeriodUpdated struct {
	OperatorSet        OperatorSet
	MaxStalenessPeriod uint32
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMaxStalenessPeriodUpdated is a free log retrieval operation binding the contract event 0x28539469fbbc8a5482e60966bf9376f7b9d25b2f0a65a9976f6baa3f0e3788da.
//
// Solidity: event MaxStalenessPeriodUpdated((address,uint32) operatorSet, uint32 maxStalenessPeriod)
func (_IECDSACertificateVerifier *IECDSACertificateVerifierFilterer) FilterMaxStalenessPeriodUpdated(opts *bind.FilterOpts) (*IECDSACertificateVerifierMaxStalenessPeriodUpdatedIterator, error) {

	logs, sub, err := _IECDSACertificateVerifier.contract.FilterLogs(opts, "MaxStalenessPeriodUpdated")
	if err != nil {
		return nil, err
	}
	return &IECDSACertificateVerifierMaxStalenessPeriodUpdatedIterator{contract: _IECDSACertificateVerifier.contract, event: "MaxStalenessPeriodUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxStalenessPeriodUpdated is a free log subscription operation binding the contract event 0x28539469fbbc8a5482e60966bf9376f7b9d25b2f0a65a9976f6baa3f0e3788da.
//
// Solidity: event MaxStalenessPeriodUpdated((address,uint32) operatorSet, uint32 maxStalenessPeriod)
func (_IECDSACertificateVerifier *IECDSACertificateVerifierFilterer) WatchMaxStalenessPeriodUpdated(opts *bind.WatchOpts, sink chan<- *IECDSACertificateVerifierMaxStalenessPeriodUpdated) (event.Subscription, error) {

	logs, sub, err := _IECDSACertificateVerifier.contract.WatchLogs(opts, "MaxStalenessPeriodUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IECDSACertificateVerifierMaxStalenessPeriodUpdated)
				if err := _IECDSACertificateVerifier.contract.UnpackLog(event, "MaxStalenessPeriodUpdated", log); err != nil {
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
func (_IECDSACertificateVerifier *IECDSACertificateVerifierFilterer) ParseMaxStalenessPeriodUpdated(log types.Log) (*IECDSACertificateVerifierMaxStalenessPeriodUpdated, error) {
	event := new(IECDSACertificateVerifierMaxStalenessPeriodUpdated)
	if err := _IECDSACertificateVerifier.contract.UnpackLog(event, "MaxStalenessPeriodUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IECDSACertificateVerifierOperatorSetOwnerUpdatedIterator is returned from FilterOperatorSetOwnerUpdated and is used to iterate over the raw logs and unpacked data for OperatorSetOwnerUpdated events raised by the IECDSACertificateVerifier contract.
type IECDSACertificateVerifierOperatorSetOwnerUpdatedIterator struct {
	Event *IECDSACertificateVerifierOperatorSetOwnerUpdated // Event containing the contract specifics and raw log

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
func (it *IECDSACertificateVerifierOperatorSetOwnerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IECDSACertificateVerifierOperatorSetOwnerUpdated)
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
		it.Event = new(IECDSACertificateVerifierOperatorSetOwnerUpdated)
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
func (it *IECDSACertificateVerifierOperatorSetOwnerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IECDSACertificateVerifierOperatorSetOwnerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IECDSACertificateVerifierOperatorSetOwnerUpdated represents a OperatorSetOwnerUpdated event raised by the IECDSACertificateVerifier contract.
type IECDSACertificateVerifierOperatorSetOwnerUpdated struct {
	OperatorSet OperatorSet
	Owner       common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorSetOwnerUpdated is a free log retrieval operation binding the contract event 0x806dc367095c0baf953d7144b7c4376261675ee0b4e0da2761e43673051c7375.
//
// Solidity: event OperatorSetOwnerUpdated((address,uint32) operatorSet, address owner)
func (_IECDSACertificateVerifier *IECDSACertificateVerifierFilterer) FilterOperatorSetOwnerUpdated(opts *bind.FilterOpts) (*IECDSACertificateVerifierOperatorSetOwnerUpdatedIterator, error) {

	logs, sub, err := _IECDSACertificateVerifier.contract.FilterLogs(opts, "OperatorSetOwnerUpdated")
	if err != nil {
		return nil, err
	}
	return &IECDSACertificateVerifierOperatorSetOwnerUpdatedIterator{contract: _IECDSACertificateVerifier.contract, event: "OperatorSetOwnerUpdated", logs: logs, sub: sub}, nil
}

// WatchOperatorSetOwnerUpdated is a free log subscription operation binding the contract event 0x806dc367095c0baf953d7144b7c4376261675ee0b4e0da2761e43673051c7375.
//
// Solidity: event OperatorSetOwnerUpdated((address,uint32) operatorSet, address owner)
func (_IECDSACertificateVerifier *IECDSACertificateVerifierFilterer) WatchOperatorSetOwnerUpdated(opts *bind.WatchOpts, sink chan<- *IECDSACertificateVerifierOperatorSetOwnerUpdated) (event.Subscription, error) {

	logs, sub, err := _IECDSACertificateVerifier.contract.WatchLogs(opts, "OperatorSetOwnerUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IECDSACertificateVerifierOperatorSetOwnerUpdated)
				if err := _IECDSACertificateVerifier.contract.UnpackLog(event, "OperatorSetOwnerUpdated", log); err != nil {
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
func (_IECDSACertificateVerifier *IECDSACertificateVerifierFilterer) ParseOperatorSetOwnerUpdated(log types.Log) (*IECDSACertificateVerifierOperatorSetOwnerUpdated, error) {
	event := new(IECDSACertificateVerifierOperatorSetOwnerUpdated)
	if err := _IECDSACertificateVerifier.contract.UnpackLog(event, "OperatorSetOwnerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IECDSACertificateVerifierTableUpdatedIterator is returned from FilterTableUpdated and is used to iterate over the raw logs and unpacked data for TableUpdated events raised by the IECDSACertificateVerifier contract.
type IECDSACertificateVerifierTableUpdatedIterator struct {
	Event *IECDSACertificateVerifierTableUpdated // Event containing the contract specifics and raw log

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
func (it *IECDSACertificateVerifierTableUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IECDSACertificateVerifierTableUpdated)
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
		it.Event = new(IECDSACertificateVerifierTableUpdated)
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
func (it *IECDSACertificateVerifierTableUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IECDSACertificateVerifierTableUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IECDSACertificateVerifierTableUpdated represents a TableUpdated event raised by the IECDSACertificateVerifier contract.
type IECDSACertificateVerifierTableUpdated struct {
	OperatorSet        OperatorSet
	ReferenceTimestamp uint32
	OperatorInfos      []IOperatorTableCalculatorTypesECDSAOperatorInfo
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterTableUpdated is a free log retrieval operation binding the contract event 0x4f588da9ec57976194a79b5594f8f8782923d93013df2b9ed12fe125805011ef.
//
// Solidity: event TableUpdated((address,uint32) operatorSet, uint32 referenceTimestamp, (address,uint256[])[] operatorInfos)
func (_IECDSACertificateVerifier *IECDSACertificateVerifierFilterer) FilterTableUpdated(opts *bind.FilterOpts) (*IECDSACertificateVerifierTableUpdatedIterator, error) {

	logs, sub, err := _IECDSACertificateVerifier.contract.FilterLogs(opts, "TableUpdated")
	if err != nil {
		return nil, err
	}
	return &IECDSACertificateVerifierTableUpdatedIterator{contract: _IECDSACertificateVerifier.contract, event: "TableUpdated", logs: logs, sub: sub}, nil
}

// WatchTableUpdated is a free log subscription operation binding the contract event 0x4f588da9ec57976194a79b5594f8f8782923d93013df2b9ed12fe125805011ef.
//
// Solidity: event TableUpdated((address,uint32) operatorSet, uint32 referenceTimestamp, (address,uint256[])[] operatorInfos)
func (_IECDSACertificateVerifier *IECDSACertificateVerifierFilterer) WatchTableUpdated(opts *bind.WatchOpts, sink chan<- *IECDSACertificateVerifierTableUpdated) (event.Subscription, error) {

	logs, sub, err := _IECDSACertificateVerifier.contract.WatchLogs(opts, "TableUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IECDSACertificateVerifierTableUpdated)
				if err := _IECDSACertificateVerifier.contract.UnpackLog(event, "TableUpdated", log); err != nil {
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
func (_IECDSACertificateVerifier *IECDSACertificateVerifierFilterer) ParseTableUpdated(log types.Log) (*IECDSACertificateVerifierTableUpdated, error) {
	event := new(IECDSACertificateVerifierTableUpdated)
	if err := _IECDSACertificateVerifier.contract.UnpackLog(event, "TableUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
