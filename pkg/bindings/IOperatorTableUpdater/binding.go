// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IOperatorTableUpdater

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
	OperatorInfo      IOperatorTableCalculatorTypesBN254OperatorInfo
}

// ICrossChainRegistryTypesOperatorSetConfig is an auto generated low-level Go binding around an user-defined struct.
type ICrossChainRegistryTypesOperatorSetConfig struct {
	Owner              common.Address
	MaxStalenessPeriod uint32
}

// IOperatorTableCalculatorTypesBN254OperatorInfo is an auto generated low-level Go binding around an user-defined struct.
type IOperatorTableCalculatorTypesBN254OperatorInfo struct {
	Pubkey  BN254G1Point
	Weights []*big.Int
}

// IOperatorTableCalculatorTypesBN254OperatorSetInfo is an auto generated low-level Go binding around an user-defined struct.
type IOperatorTableCalculatorTypesBN254OperatorSetInfo struct {
	OperatorInfoTreeRoot [32]byte
	NumOperators         *big.Int
	AggregatePubkey      BN254G1Point
	TotalWeights         []*big.Int
}

// OperatorSet is an auto generated low-level Go binding around an user-defined struct.
type OperatorSet struct {
	Avs common.Address
	Id  uint32
}

// IOperatorTableUpdaterMetaData contains all meta data concerning the IOperatorTableUpdater contract.
var IOperatorTableUpdaterMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"confirmGlobalTableRoot\",\"inputs\":[{\"name\":\"globalTableRootCert\",\"type\":\"tuple\",\"internalType\":\"structIBN254CertificateVerifierTypes.BN254Certificate\",\"components\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"nonSignerWitnesses\",\"type\":\"tuple[]\",\"internalType\":\"structIBN254CertificateVerifierTypes.BN254OperatorInfoWitness[]\",\"components\":[{\"name\":\"operatorIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorInfoProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"operatorInfo\",\"type\":\"tuple\",\"internalType\":\"structIOperatorTableCalculatorTypes.BN254OperatorInfo\",\"components\":[{\"name\":\"pubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}]}]},{\"name\":\"globalTableRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"disableRoot\",\"inputs\":[{\"name\":\"globalTableRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getCertificateVerifier\",\"inputs\":[{\"name\":\"curveType\",\"type\":\"uint8\",\"internalType\":\"enumIKeyRegistrarTypes.CurveType\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentGlobalTableRoot\",\"inputs\":[],\"outputs\":[{\"name\":\"globalTableRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getGenerator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getGeneratorConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structICrossChainRegistryTypes.OperatorSetConfig\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxStalenessPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getGeneratorReferenceTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getGlobalTableRootByTimestamp\",\"inputs\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"tableRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getGlobalTableUpdateMessageHash\",\"inputs\":[{\"name\":\"globalTableRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getGlobalTableUpdateSignableDigest\",\"inputs\":[{\"name\":\"globalTableRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLatestReferenceBlockNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLatestReferenceTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getReferenceBlockNumberByTimestamp\",\"inputs\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getReferenceTimestampByBlockNumber\",\"inputs\":[{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRootValid\",\"inputs\":[{\"name\":\"globalTableRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRootValidByTimestamp\",\"inputs\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setGlobalRootConfirmationThreshold\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateGenerator\",\"inputs\":[{\"name\":\"generator\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"generatorInfo\",\"type\":\"tuple\",\"internalType\":\"structIOperatorTableCalculatorTypes.BN254OperatorSetInfo\",\"components\":[{\"name\":\"operatorInfoTreeRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"numOperators\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"aggregatePubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"totalWeights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOperatorTable\",\"inputs\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"globalTableRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"operatorSetIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"operatorTableBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"GeneratorUpdated\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GlobalRootConfirmationThresholdUpdated\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GlobalRootDisabled\",\"inputs\":[{\"name\":\"globalTableRoot\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewGlobalTableRoot\",\"inputs\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"globalTableRoot\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CannotDisableGeneratorRoot\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CertificateInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"GlobalTableRootInFuture\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"GlobalTableRootStale\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidConfirmationThreshold\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCurveType\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidGenerator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidGlobalTableRoot\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidMessageHash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperatorSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperatorSetProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRoot\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TableUpdateForPastTimestamp\",\"inputs\":[]}]",
}

// IOperatorTableUpdaterABI is the input ABI used to generate the binding from.
// Deprecated: Use IOperatorTableUpdaterMetaData.ABI instead.
var IOperatorTableUpdaterABI = IOperatorTableUpdaterMetaData.ABI

// IOperatorTableUpdater is an auto generated Go binding around an Ethereum contract.
type IOperatorTableUpdater struct {
	IOperatorTableUpdaterCaller     // Read-only binding to the contract
	IOperatorTableUpdaterTransactor // Write-only binding to the contract
	IOperatorTableUpdaterFilterer   // Log filterer for contract events
}

// IOperatorTableUpdaterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IOperatorTableUpdaterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOperatorTableUpdaterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IOperatorTableUpdaterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOperatorTableUpdaterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IOperatorTableUpdaterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOperatorTableUpdaterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IOperatorTableUpdaterSession struct {
	Contract     *IOperatorTableUpdater // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IOperatorTableUpdaterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IOperatorTableUpdaterCallerSession struct {
	Contract *IOperatorTableUpdaterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// IOperatorTableUpdaterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IOperatorTableUpdaterTransactorSession struct {
	Contract     *IOperatorTableUpdaterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// IOperatorTableUpdaterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IOperatorTableUpdaterRaw struct {
	Contract *IOperatorTableUpdater // Generic contract binding to access the raw methods on
}

// IOperatorTableUpdaterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IOperatorTableUpdaterCallerRaw struct {
	Contract *IOperatorTableUpdaterCaller // Generic read-only contract binding to access the raw methods on
}

// IOperatorTableUpdaterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IOperatorTableUpdaterTransactorRaw struct {
	Contract *IOperatorTableUpdaterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIOperatorTableUpdater creates a new instance of IOperatorTableUpdater, bound to a specific deployed contract.
func NewIOperatorTableUpdater(address common.Address, backend bind.ContractBackend) (*IOperatorTableUpdater, error) {
	contract, err := bindIOperatorTableUpdater(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IOperatorTableUpdater{IOperatorTableUpdaterCaller: IOperatorTableUpdaterCaller{contract: contract}, IOperatorTableUpdaterTransactor: IOperatorTableUpdaterTransactor{contract: contract}, IOperatorTableUpdaterFilterer: IOperatorTableUpdaterFilterer{contract: contract}}, nil
}

// NewIOperatorTableUpdaterCaller creates a new read-only instance of IOperatorTableUpdater, bound to a specific deployed contract.
func NewIOperatorTableUpdaterCaller(address common.Address, caller bind.ContractCaller) (*IOperatorTableUpdaterCaller, error) {
	contract, err := bindIOperatorTableUpdater(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IOperatorTableUpdaterCaller{contract: contract}, nil
}

// NewIOperatorTableUpdaterTransactor creates a new write-only instance of IOperatorTableUpdater, bound to a specific deployed contract.
func NewIOperatorTableUpdaterTransactor(address common.Address, transactor bind.ContractTransactor) (*IOperatorTableUpdaterTransactor, error) {
	contract, err := bindIOperatorTableUpdater(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IOperatorTableUpdaterTransactor{contract: contract}, nil
}

// NewIOperatorTableUpdaterFilterer creates a new log filterer instance of IOperatorTableUpdater, bound to a specific deployed contract.
func NewIOperatorTableUpdaterFilterer(address common.Address, filterer bind.ContractFilterer) (*IOperatorTableUpdaterFilterer, error) {
	contract, err := bindIOperatorTableUpdater(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IOperatorTableUpdaterFilterer{contract: contract}, nil
}

// bindIOperatorTableUpdater binds a generic wrapper to an already deployed contract.
func bindIOperatorTableUpdater(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IOperatorTableUpdaterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOperatorTableUpdater *IOperatorTableUpdaterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOperatorTableUpdater.Contract.IOperatorTableUpdaterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOperatorTableUpdater *IOperatorTableUpdaterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOperatorTableUpdater.Contract.IOperatorTableUpdaterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOperatorTableUpdater *IOperatorTableUpdaterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOperatorTableUpdater.Contract.IOperatorTableUpdaterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOperatorTableUpdater *IOperatorTableUpdaterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOperatorTableUpdater.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOperatorTableUpdater *IOperatorTableUpdaterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOperatorTableUpdater.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOperatorTableUpdater *IOperatorTableUpdaterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOperatorTableUpdater.Contract.contract.Transact(opts, method, params...)
}

// GetCertificateVerifier is a free data retrieval call binding the contract method 0x6f728c50.
//
// Solidity: function getCertificateVerifier(uint8 curveType) view returns(address)
func (_IOperatorTableUpdater *IOperatorTableUpdaterCaller) GetCertificateVerifier(opts *bind.CallOpts, curveType uint8) (common.Address, error) {
	var out []interface{}
	err := _IOperatorTableUpdater.contract.Call(opts, &out, "getCertificateVerifier", curveType)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCertificateVerifier is a free data retrieval call binding the contract method 0x6f728c50.
//
// Solidity: function getCertificateVerifier(uint8 curveType) view returns(address)
func (_IOperatorTableUpdater *IOperatorTableUpdaterSession) GetCertificateVerifier(curveType uint8) (common.Address, error) {
	return _IOperatorTableUpdater.Contract.GetCertificateVerifier(&_IOperatorTableUpdater.CallOpts, curveType)
}

// GetCertificateVerifier is a free data retrieval call binding the contract method 0x6f728c50.
//
// Solidity: function getCertificateVerifier(uint8 curveType) view returns(address)
func (_IOperatorTableUpdater *IOperatorTableUpdaterCallerSession) GetCertificateVerifier(curveType uint8) (common.Address, error) {
	return _IOperatorTableUpdater.Contract.GetCertificateVerifier(&_IOperatorTableUpdater.CallOpts, curveType)
}

// GetCurrentGlobalTableRoot is a free data retrieval call binding the contract method 0x28522d79.
//
// Solidity: function getCurrentGlobalTableRoot() view returns(bytes32 globalTableRoot)
func (_IOperatorTableUpdater *IOperatorTableUpdaterCaller) GetCurrentGlobalTableRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IOperatorTableUpdater.contract.Call(opts, &out, "getCurrentGlobalTableRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetCurrentGlobalTableRoot is a free data retrieval call binding the contract method 0x28522d79.
//
// Solidity: function getCurrentGlobalTableRoot() view returns(bytes32 globalTableRoot)
func (_IOperatorTableUpdater *IOperatorTableUpdaterSession) GetCurrentGlobalTableRoot() ([32]byte, error) {
	return _IOperatorTableUpdater.Contract.GetCurrentGlobalTableRoot(&_IOperatorTableUpdater.CallOpts)
}

// GetCurrentGlobalTableRoot is a free data retrieval call binding the contract method 0x28522d79.
//
// Solidity: function getCurrentGlobalTableRoot() view returns(bytes32 globalTableRoot)
func (_IOperatorTableUpdater *IOperatorTableUpdaterCallerSession) GetCurrentGlobalTableRoot() ([32]byte, error) {
	return _IOperatorTableUpdater.Contract.GetCurrentGlobalTableRoot(&_IOperatorTableUpdater.CallOpts)
}

// GetGenerator is a free data retrieval call binding the contract method 0x1e2ca260.
//
// Solidity: function getGenerator() view returns((address,uint32))
func (_IOperatorTableUpdater *IOperatorTableUpdaterCaller) GetGenerator(opts *bind.CallOpts) (OperatorSet, error) {
	var out []interface{}
	err := _IOperatorTableUpdater.contract.Call(opts, &out, "getGenerator")

	if err != nil {
		return *new(OperatorSet), err
	}

	out0 := *abi.ConvertType(out[0], new(OperatorSet)).(*OperatorSet)

	return out0, err

}

// GetGenerator is a free data retrieval call binding the contract method 0x1e2ca260.
//
// Solidity: function getGenerator() view returns((address,uint32))
func (_IOperatorTableUpdater *IOperatorTableUpdaterSession) GetGenerator() (OperatorSet, error) {
	return _IOperatorTableUpdater.Contract.GetGenerator(&_IOperatorTableUpdater.CallOpts)
}

// GetGenerator is a free data retrieval call binding the contract method 0x1e2ca260.
//
// Solidity: function getGenerator() view returns((address,uint32))
func (_IOperatorTableUpdater *IOperatorTableUpdaterCallerSession) GetGenerator() (OperatorSet, error) {
	return _IOperatorTableUpdater.Contract.GetGenerator(&_IOperatorTableUpdater.CallOpts)
}

// GetGeneratorConfig is a free data retrieval call binding the contract method 0xb0cb3a24.
//
// Solidity: function getGeneratorConfig() view returns((address,uint32))
func (_IOperatorTableUpdater *IOperatorTableUpdaterCaller) GetGeneratorConfig(opts *bind.CallOpts) (ICrossChainRegistryTypesOperatorSetConfig, error) {
	var out []interface{}
	err := _IOperatorTableUpdater.contract.Call(opts, &out, "getGeneratorConfig")

	if err != nil {
		return *new(ICrossChainRegistryTypesOperatorSetConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(ICrossChainRegistryTypesOperatorSetConfig)).(*ICrossChainRegistryTypesOperatorSetConfig)

	return out0, err

}

// GetGeneratorConfig is a free data retrieval call binding the contract method 0xb0cb3a24.
//
// Solidity: function getGeneratorConfig() view returns((address,uint32))
func (_IOperatorTableUpdater *IOperatorTableUpdaterSession) GetGeneratorConfig() (ICrossChainRegistryTypesOperatorSetConfig, error) {
	return _IOperatorTableUpdater.Contract.GetGeneratorConfig(&_IOperatorTableUpdater.CallOpts)
}

// GetGeneratorConfig is a free data retrieval call binding the contract method 0xb0cb3a24.
//
// Solidity: function getGeneratorConfig() view returns((address,uint32))
func (_IOperatorTableUpdater *IOperatorTableUpdaterCallerSession) GetGeneratorConfig() (ICrossChainRegistryTypesOperatorSetConfig, error) {
	return _IOperatorTableUpdater.Contract.GetGeneratorConfig(&_IOperatorTableUpdater.CallOpts)
}

// GetGeneratorReferenceTimestamp is a free data retrieval call binding the contract method 0x7551ba34.
//
// Solidity: function getGeneratorReferenceTimestamp() view returns(uint32)
func (_IOperatorTableUpdater *IOperatorTableUpdaterCaller) GetGeneratorReferenceTimestamp(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _IOperatorTableUpdater.contract.Call(opts, &out, "getGeneratorReferenceTimestamp")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetGeneratorReferenceTimestamp is a free data retrieval call binding the contract method 0x7551ba34.
//
// Solidity: function getGeneratorReferenceTimestamp() view returns(uint32)
func (_IOperatorTableUpdater *IOperatorTableUpdaterSession) GetGeneratorReferenceTimestamp() (uint32, error) {
	return _IOperatorTableUpdater.Contract.GetGeneratorReferenceTimestamp(&_IOperatorTableUpdater.CallOpts)
}

// GetGeneratorReferenceTimestamp is a free data retrieval call binding the contract method 0x7551ba34.
//
// Solidity: function getGeneratorReferenceTimestamp() view returns(uint32)
func (_IOperatorTableUpdater *IOperatorTableUpdaterCallerSession) GetGeneratorReferenceTimestamp() (uint32, error) {
	return _IOperatorTableUpdater.Contract.GetGeneratorReferenceTimestamp(&_IOperatorTableUpdater.CallOpts)
}

// GetGlobalTableRootByTimestamp is a free data retrieval call binding the contract method 0xc5916a39.
//
// Solidity: function getGlobalTableRootByTimestamp(uint32 referenceTimestamp) view returns(bytes32 tableRoot)
func (_IOperatorTableUpdater *IOperatorTableUpdaterCaller) GetGlobalTableRootByTimestamp(opts *bind.CallOpts, referenceTimestamp uint32) ([32]byte, error) {
	var out []interface{}
	err := _IOperatorTableUpdater.contract.Call(opts, &out, "getGlobalTableRootByTimestamp", referenceTimestamp)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetGlobalTableRootByTimestamp is a free data retrieval call binding the contract method 0xc5916a39.
//
// Solidity: function getGlobalTableRootByTimestamp(uint32 referenceTimestamp) view returns(bytes32 tableRoot)
func (_IOperatorTableUpdater *IOperatorTableUpdaterSession) GetGlobalTableRootByTimestamp(referenceTimestamp uint32) ([32]byte, error) {
	return _IOperatorTableUpdater.Contract.GetGlobalTableRootByTimestamp(&_IOperatorTableUpdater.CallOpts, referenceTimestamp)
}

// GetGlobalTableRootByTimestamp is a free data retrieval call binding the contract method 0xc5916a39.
//
// Solidity: function getGlobalTableRootByTimestamp(uint32 referenceTimestamp) view returns(bytes32 tableRoot)
func (_IOperatorTableUpdater *IOperatorTableUpdaterCallerSession) GetGlobalTableRootByTimestamp(referenceTimestamp uint32) ([32]byte, error) {
	return _IOperatorTableUpdater.Contract.GetGlobalTableRootByTimestamp(&_IOperatorTableUpdater.CallOpts, referenceTimestamp)
}

// GetGlobalTableUpdateMessageHash is a free data retrieval call binding the contract method 0xc3be1e33.
//
// Solidity: function getGlobalTableUpdateMessageHash(bytes32 globalTableRoot, uint32 referenceTimestamp, uint32 referenceBlockNumber) view returns(bytes32)
func (_IOperatorTableUpdater *IOperatorTableUpdaterCaller) GetGlobalTableUpdateMessageHash(opts *bind.CallOpts, globalTableRoot [32]byte, referenceTimestamp uint32, referenceBlockNumber uint32) ([32]byte, error) {
	var out []interface{}
	err := _IOperatorTableUpdater.contract.Call(opts, &out, "getGlobalTableUpdateMessageHash", globalTableRoot, referenceTimestamp, referenceBlockNumber)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetGlobalTableUpdateMessageHash is a free data retrieval call binding the contract method 0xc3be1e33.
//
// Solidity: function getGlobalTableUpdateMessageHash(bytes32 globalTableRoot, uint32 referenceTimestamp, uint32 referenceBlockNumber) view returns(bytes32)
func (_IOperatorTableUpdater *IOperatorTableUpdaterSession) GetGlobalTableUpdateMessageHash(globalTableRoot [32]byte, referenceTimestamp uint32, referenceBlockNumber uint32) ([32]byte, error) {
	return _IOperatorTableUpdater.Contract.GetGlobalTableUpdateMessageHash(&_IOperatorTableUpdater.CallOpts, globalTableRoot, referenceTimestamp, referenceBlockNumber)
}

// GetGlobalTableUpdateMessageHash is a free data retrieval call binding the contract method 0xc3be1e33.
//
// Solidity: function getGlobalTableUpdateMessageHash(bytes32 globalTableRoot, uint32 referenceTimestamp, uint32 referenceBlockNumber) view returns(bytes32)
func (_IOperatorTableUpdater *IOperatorTableUpdaterCallerSession) GetGlobalTableUpdateMessageHash(globalTableRoot [32]byte, referenceTimestamp uint32, referenceBlockNumber uint32) ([32]byte, error) {
	return _IOperatorTableUpdater.Contract.GetGlobalTableUpdateMessageHash(&_IOperatorTableUpdater.CallOpts, globalTableRoot, referenceTimestamp, referenceBlockNumber)
}

// GetGlobalTableUpdateSignableDigest is a free data retrieval call binding the contract method 0x401c370f.
//
// Solidity: function getGlobalTableUpdateSignableDigest(bytes32 globalTableRoot, uint32 referenceTimestamp, uint32 referenceBlockNumber) view returns(bytes32)
func (_IOperatorTableUpdater *IOperatorTableUpdaterCaller) GetGlobalTableUpdateSignableDigest(opts *bind.CallOpts, globalTableRoot [32]byte, referenceTimestamp uint32, referenceBlockNumber uint32) ([32]byte, error) {
	var out []interface{}
	err := _IOperatorTableUpdater.contract.Call(opts, &out, "getGlobalTableUpdateSignableDigest", globalTableRoot, referenceTimestamp, referenceBlockNumber)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetGlobalTableUpdateSignableDigest is a free data retrieval call binding the contract method 0x401c370f.
//
// Solidity: function getGlobalTableUpdateSignableDigest(bytes32 globalTableRoot, uint32 referenceTimestamp, uint32 referenceBlockNumber) view returns(bytes32)
func (_IOperatorTableUpdater *IOperatorTableUpdaterSession) GetGlobalTableUpdateSignableDigest(globalTableRoot [32]byte, referenceTimestamp uint32, referenceBlockNumber uint32) ([32]byte, error) {
	return _IOperatorTableUpdater.Contract.GetGlobalTableUpdateSignableDigest(&_IOperatorTableUpdater.CallOpts, globalTableRoot, referenceTimestamp, referenceBlockNumber)
}

// GetGlobalTableUpdateSignableDigest is a free data retrieval call binding the contract method 0x401c370f.
//
// Solidity: function getGlobalTableUpdateSignableDigest(bytes32 globalTableRoot, uint32 referenceTimestamp, uint32 referenceBlockNumber) view returns(bytes32)
func (_IOperatorTableUpdater *IOperatorTableUpdaterCallerSession) GetGlobalTableUpdateSignableDigest(globalTableRoot [32]byte, referenceTimestamp uint32, referenceBlockNumber uint32) ([32]byte, error) {
	return _IOperatorTableUpdater.Contract.GetGlobalTableUpdateSignableDigest(&_IOperatorTableUpdater.CallOpts, globalTableRoot, referenceTimestamp, referenceBlockNumber)
}

// GetLatestReferenceBlockNumber is a free data retrieval call binding the contract method 0x31a599d2.
//
// Solidity: function getLatestReferenceBlockNumber() view returns(uint32)
func (_IOperatorTableUpdater *IOperatorTableUpdaterCaller) GetLatestReferenceBlockNumber(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _IOperatorTableUpdater.contract.Call(opts, &out, "getLatestReferenceBlockNumber")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetLatestReferenceBlockNumber is a free data retrieval call binding the contract method 0x31a599d2.
//
// Solidity: function getLatestReferenceBlockNumber() view returns(uint32)
func (_IOperatorTableUpdater *IOperatorTableUpdaterSession) GetLatestReferenceBlockNumber() (uint32, error) {
	return _IOperatorTableUpdater.Contract.GetLatestReferenceBlockNumber(&_IOperatorTableUpdater.CallOpts)
}

// GetLatestReferenceBlockNumber is a free data retrieval call binding the contract method 0x31a599d2.
//
// Solidity: function getLatestReferenceBlockNumber() view returns(uint32)
func (_IOperatorTableUpdater *IOperatorTableUpdaterCallerSession) GetLatestReferenceBlockNumber() (uint32, error) {
	return _IOperatorTableUpdater.Contract.GetLatestReferenceBlockNumber(&_IOperatorTableUpdater.CallOpts)
}

// GetLatestReferenceTimestamp is a free data retrieval call binding the contract method 0x4624e6a3.
//
// Solidity: function getLatestReferenceTimestamp() view returns(uint32)
func (_IOperatorTableUpdater *IOperatorTableUpdaterCaller) GetLatestReferenceTimestamp(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _IOperatorTableUpdater.contract.Call(opts, &out, "getLatestReferenceTimestamp")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetLatestReferenceTimestamp is a free data retrieval call binding the contract method 0x4624e6a3.
//
// Solidity: function getLatestReferenceTimestamp() view returns(uint32)
func (_IOperatorTableUpdater *IOperatorTableUpdaterSession) GetLatestReferenceTimestamp() (uint32, error) {
	return _IOperatorTableUpdater.Contract.GetLatestReferenceTimestamp(&_IOperatorTableUpdater.CallOpts)
}

// GetLatestReferenceTimestamp is a free data retrieval call binding the contract method 0x4624e6a3.
//
// Solidity: function getLatestReferenceTimestamp() view returns(uint32)
func (_IOperatorTableUpdater *IOperatorTableUpdaterCallerSession) GetLatestReferenceTimestamp() (uint32, error) {
	return _IOperatorTableUpdater.Contract.GetLatestReferenceTimestamp(&_IOperatorTableUpdater.CallOpts)
}

// GetReferenceBlockNumberByTimestamp is a free data retrieval call binding the contract method 0x23b7b5b2.
//
// Solidity: function getReferenceBlockNumberByTimestamp(uint32 referenceTimestamp) view returns(uint32)
func (_IOperatorTableUpdater *IOperatorTableUpdaterCaller) GetReferenceBlockNumberByTimestamp(opts *bind.CallOpts, referenceTimestamp uint32) (uint32, error) {
	var out []interface{}
	err := _IOperatorTableUpdater.contract.Call(opts, &out, "getReferenceBlockNumberByTimestamp", referenceTimestamp)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetReferenceBlockNumberByTimestamp is a free data retrieval call binding the contract method 0x23b7b5b2.
//
// Solidity: function getReferenceBlockNumberByTimestamp(uint32 referenceTimestamp) view returns(uint32)
func (_IOperatorTableUpdater *IOperatorTableUpdaterSession) GetReferenceBlockNumberByTimestamp(referenceTimestamp uint32) (uint32, error) {
	return _IOperatorTableUpdater.Contract.GetReferenceBlockNumberByTimestamp(&_IOperatorTableUpdater.CallOpts, referenceTimestamp)
}

// GetReferenceBlockNumberByTimestamp is a free data retrieval call binding the contract method 0x23b7b5b2.
//
// Solidity: function getReferenceBlockNumberByTimestamp(uint32 referenceTimestamp) view returns(uint32)
func (_IOperatorTableUpdater *IOperatorTableUpdaterCallerSession) GetReferenceBlockNumberByTimestamp(referenceTimestamp uint32) (uint32, error) {
	return _IOperatorTableUpdater.Contract.GetReferenceBlockNumberByTimestamp(&_IOperatorTableUpdater.CallOpts, referenceTimestamp)
}

// GetReferenceTimestampByBlockNumber is a free data retrieval call binding the contract method 0x193b79f3.
//
// Solidity: function getReferenceTimestampByBlockNumber(uint32 referenceBlockNumber) view returns(uint32)
func (_IOperatorTableUpdater *IOperatorTableUpdaterCaller) GetReferenceTimestampByBlockNumber(opts *bind.CallOpts, referenceBlockNumber uint32) (uint32, error) {
	var out []interface{}
	err := _IOperatorTableUpdater.contract.Call(opts, &out, "getReferenceTimestampByBlockNumber", referenceBlockNumber)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetReferenceTimestampByBlockNumber is a free data retrieval call binding the contract method 0x193b79f3.
//
// Solidity: function getReferenceTimestampByBlockNumber(uint32 referenceBlockNumber) view returns(uint32)
func (_IOperatorTableUpdater *IOperatorTableUpdaterSession) GetReferenceTimestampByBlockNumber(referenceBlockNumber uint32) (uint32, error) {
	return _IOperatorTableUpdater.Contract.GetReferenceTimestampByBlockNumber(&_IOperatorTableUpdater.CallOpts, referenceBlockNumber)
}

// GetReferenceTimestampByBlockNumber is a free data retrieval call binding the contract method 0x193b79f3.
//
// Solidity: function getReferenceTimestampByBlockNumber(uint32 referenceBlockNumber) view returns(uint32)
func (_IOperatorTableUpdater *IOperatorTableUpdaterCallerSession) GetReferenceTimestampByBlockNumber(referenceBlockNumber uint32) (uint32, error) {
	return _IOperatorTableUpdater.Contract.GetReferenceTimestampByBlockNumber(&_IOperatorTableUpdater.CallOpts, referenceBlockNumber)
}

// IsRootValid is a free data retrieval call binding the contract method 0x30ef41b4.
//
// Solidity: function isRootValid(bytes32 globalTableRoot) view returns(bool)
func (_IOperatorTableUpdater *IOperatorTableUpdaterCaller) IsRootValid(opts *bind.CallOpts, globalTableRoot [32]byte) (bool, error) {
	var out []interface{}
	err := _IOperatorTableUpdater.contract.Call(opts, &out, "isRootValid", globalTableRoot)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRootValid is a free data retrieval call binding the contract method 0x30ef41b4.
//
// Solidity: function isRootValid(bytes32 globalTableRoot) view returns(bool)
func (_IOperatorTableUpdater *IOperatorTableUpdaterSession) IsRootValid(globalTableRoot [32]byte) (bool, error) {
	return _IOperatorTableUpdater.Contract.IsRootValid(&_IOperatorTableUpdater.CallOpts, globalTableRoot)
}

// IsRootValid is a free data retrieval call binding the contract method 0x30ef41b4.
//
// Solidity: function isRootValid(bytes32 globalTableRoot) view returns(bool)
func (_IOperatorTableUpdater *IOperatorTableUpdaterCallerSession) IsRootValid(globalTableRoot [32]byte) (bool, error) {
	return _IOperatorTableUpdater.Contract.IsRootValid(&_IOperatorTableUpdater.CallOpts, globalTableRoot)
}

// IsRootValidByTimestamp is a free data retrieval call binding the contract method 0x64e1df84.
//
// Solidity: function isRootValidByTimestamp(uint32 referenceTimestamp) view returns(bool)
func (_IOperatorTableUpdater *IOperatorTableUpdaterCaller) IsRootValidByTimestamp(opts *bind.CallOpts, referenceTimestamp uint32) (bool, error) {
	var out []interface{}
	err := _IOperatorTableUpdater.contract.Call(opts, &out, "isRootValidByTimestamp", referenceTimestamp)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRootValidByTimestamp is a free data retrieval call binding the contract method 0x64e1df84.
//
// Solidity: function isRootValidByTimestamp(uint32 referenceTimestamp) view returns(bool)
func (_IOperatorTableUpdater *IOperatorTableUpdaterSession) IsRootValidByTimestamp(referenceTimestamp uint32) (bool, error) {
	return _IOperatorTableUpdater.Contract.IsRootValidByTimestamp(&_IOperatorTableUpdater.CallOpts, referenceTimestamp)
}

// IsRootValidByTimestamp is a free data retrieval call binding the contract method 0x64e1df84.
//
// Solidity: function isRootValidByTimestamp(uint32 referenceTimestamp) view returns(bool)
func (_IOperatorTableUpdater *IOperatorTableUpdaterCallerSession) IsRootValidByTimestamp(referenceTimestamp uint32) (bool, error) {
	return _IOperatorTableUpdater.Contract.IsRootValidByTimestamp(&_IOperatorTableUpdater.CallOpts, referenceTimestamp)
}

// ConfirmGlobalTableRoot is a paid mutator transaction binding the contract method 0xeaaed9d5.
//
// Solidity: function confirmGlobalTableRoot((uint32,bytes32,(uint256,uint256),(uint256[2],uint256[2]),(uint32,bytes,((uint256,uint256),uint256[]))[]) globalTableRootCert, bytes32 globalTableRoot, uint32 referenceTimestamp, uint32 referenceBlockNumber) returns()
func (_IOperatorTableUpdater *IOperatorTableUpdaterTransactor) ConfirmGlobalTableRoot(opts *bind.TransactOpts, globalTableRootCert IBN254CertificateVerifierTypesBN254Certificate, globalTableRoot [32]byte, referenceTimestamp uint32, referenceBlockNumber uint32) (*types.Transaction, error) {
	return _IOperatorTableUpdater.contract.Transact(opts, "confirmGlobalTableRoot", globalTableRootCert, globalTableRoot, referenceTimestamp, referenceBlockNumber)
}

// ConfirmGlobalTableRoot is a paid mutator transaction binding the contract method 0xeaaed9d5.
//
// Solidity: function confirmGlobalTableRoot((uint32,bytes32,(uint256,uint256),(uint256[2],uint256[2]),(uint32,bytes,((uint256,uint256),uint256[]))[]) globalTableRootCert, bytes32 globalTableRoot, uint32 referenceTimestamp, uint32 referenceBlockNumber) returns()
func (_IOperatorTableUpdater *IOperatorTableUpdaterSession) ConfirmGlobalTableRoot(globalTableRootCert IBN254CertificateVerifierTypesBN254Certificate, globalTableRoot [32]byte, referenceTimestamp uint32, referenceBlockNumber uint32) (*types.Transaction, error) {
	return _IOperatorTableUpdater.Contract.ConfirmGlobalTableRoot(&_IOperatorTableUpdater.TransactOpts, globalTableRootCert, globalTableRoot, referenceTimestamp, referenceBlockNumber)
}

// ConfirmGlobalTableRoot is a paid mutator transaction binding the contract method 0xeaaed9d5.
//
// Solidity: function confirmGlobalTableRoot((uint32,bytes32,(uint256,uint256),(uint256[2],uint256[2]),(uint32,bytes,((uint256,uint256),uint256[]))[]) globalTableRootCert, bytes32 globalTableRoot, uint32 referenceTimestamp, uint32 referenceBlockNumber) returns()
func (_IOperatorTableUpdater *IOperatorTableUpdaterTransactorSession) ConfirmGlobalTableRoot(globalTableRootCert IBN254CertificateVerifierTypesBN254Certificate, globalTableRoot [32]byte, referenceTimestamp uint32, referenceBlockNumber uint32) (*types.Transaction, error) {
	return _IOperatorTableUpdater.Contract.ConfirmGlobalTableRoot(&_IOperatorTableUpdater.TransactOpts, globalTableRootCert, globalTableRoot, referenceTimestamp, referenceBlockNumber)
}

// DisableRoot is a paid mutator transaction binding the contract method 0xc3621f0a.
//
// Solidity: function disableRoot(bytes32 globalTableRoot) returns()
func (_IOperatorTableUpdater *IOperatorTableUpdaterTransactor) DisableRoot(opts *bind.TransactOpts, globalTableRoot [32]byte) (*types.Transaction, error) {
	return _IOperatorTableUpdater.contract.Transact(opts, "disableRoot", globalTableRoot)
}

// DisableRoot is a paid mutator transaction binding the contract method 0xc3621f0a.
//
// Solidity: function disableRoot(bytes32 globalTableRoot) returns()
func (_IOperatorTableUpdater *IOperatorTableUpdaterSession) DisableRoot(globalTableRoot [32]byte) (*types.Transaction, error) {
	return _IOperatorTableUpdater.Contract.DisableRoot(&_IOperatorTableUpdater.TransactOpts, globalTableRoot)
}

// DisableRoot is a paid mutator transaction binding the contract method 0xc3621f0a.
//
// Solidity: function disableRoot(bytes32 globalTableRoot) returns()
func (_IOperatorTableUpdater *IOperatorTableUpdaterTransactorSession) DisableRoot(globalTableRoot [32]byte) (*types.Transaction, error) {
	return _IOperatorTableUpdater.Contract.DisableRoot(&_IOperatorTableUpdater.TransactOpts, globalTableRoot)
}

// SetGlobalRootConfirmationThreshold is a paid mutator transaction binding the contract method 0x2370356c.
//
// Solidity: function setGlobalRootConfirmationThreshold(uint16 bps) returns()
func (_IOperatorTableUpdater *IOperatorTableUpdaterTransactor) SetGlobalRootConfirmationThreshold(opts *bind.TransactOpts, bps uint16) (*types.Transaction, error) {
	return _IOperatorTableUpdater.contract.Transact(opts, "setGlobalRootConfirmationThreshold", bps)
}

// SetGlobalRootConfirmationThreshold is a paid mutator transaction binding the contract method 0x2370356c.
//
// Solidity: function setGlobalRootConfirmationThreshold(uint16 bps) returns()
func (_IOperatorTableUpdater *IOperatorTableUpdaterSession) SetGlobalRootConfirmationThreshold(bps uint16) (*types.Transaction, error) {
	return _IOperatorTableUpdater.Contract.SetGlobalRootConfirmationThreshold(&_IOperatorTableUpdater.TransactOpts, bps)
}

// SetGlobalRootConfirmationThreshold is a paid mutator transaction binding the contract method 0x2370356c.
//
// Solidity: function setGlobalRootConfirmationThreshold(uint16 bps) returns()
func (_IOperatorTableUpdater *IOperatorTableUpdaterTransactorSession) SetGlobalRootConfirmationThreshold(bps uint16) (*types.Transaction, error) {
	return _IOperatorTableUpdater.Contract.SetGlobalRootConfirmationThreshold(&_IOperatorTableUpdater.TransactOpts, bps)
}

// UpdateGenerator is a paid mutator transaction binding the contract method 0x9f7e206f.
//
// Solidity: function updateGenerator((address,uint32) generator, (bytes32,uint256,(uint256,uint256),uint256[]) generatorInfo) returns()
func (_IOperatorTableUpdater *IOperatorTableUpdaterTransactor) UpdateGenerator(opts *bind.TransactOpts, generator OperatorSet, generatorInfo IOperatorTableCalculatorTypesBN254OperatorSetInfo) (*types.Transaction, error) {
	return _IOperatorTableUpdater.contract.Transact(opts, "updateGenerator", generator, generatorInfo)
}

// UpdateGenerator is a paid mutator transaction binding the contract method 0x9f7e206f.
//
// Solidity: function updateGenerator((address,uint32) generator, (bytes32,uint256,(uint256,uint256),uint256[]) generatorInfo) returns()
func (_IOperatorTableUpdater *IOperatorTableUpdaterSession) UpdateGenerator(generator OperatorSet, generatorInfo IOperatorTableCalculatorTypesBN254OperatorSetInfo) (*types.Transaction, error) {
	return _IOperatorTableUpdater.Contract.UpdateGenerator(&_IOperatorTableUpdater.TransactOpts, generator, generatorInfo)
}

// UpdateGenerator is a paid mutator transaction binding the contract method 0x9f7e206f.
//
// Solidity: function updateGenerator((address,uint32) generator, (bytes32,uint256,(uint256,uint256),uint256[]) generatorInfo) returns()
func (_IOperatorTableUpdater *IOperatorTableUpdaterTransactorSession) UpdateGenerator(generator OperatorSet, generatorInfo IOperatorTableCalculatorTypesBN254OperatorSetInfo) (*types.Transaction, error) {
	return _IOperatorTableUpdater.Contract.UpdateGenerator(&_IOperatorTableUpdater.TransactOpts, generator, generatorInfo)
}

// UpdateOperatorTable is a paid mutator transaction binding the contract method 0x9ea94778.
//
// Solidity: function updateOperatorTable(uint32 referenceTimestamp, bytes32 globalTableRoot, uint32 operatorSetIndex, bytes proof, bytes operatorTableBytes) returns()
func (_IOperatorTableUpdater *IOperatorTableUpdaterTransactor) UpdateOperatorTable(opts *bind.TransactOpts, referenceTimestamp uint32, globalTableRoot [32]byte, operatorSetIndex uint32, proof []byte, operatorTableBytes []byte) (*types.Transaction, error) {
	return _IOperatorTableUpdater.contract.Transact(opts, "updateOperatorTable", referenceTimestamp, globalTableRoot, operatorSetIndex, proof, operatorTableBytes)
}

// UpdateOperatorTable is a paid mutator transaction binding the contract method 0x9ea94778.
//
// Solidity: function updateOperatorTable(uint32 referenceTimestamp, bytes32 globalTableRoot, uint32 operatorSetIndex, bytes proof, bytes operatorTableBytes) returns()
func (_IOperatorTableUpdater *IOperatorTableUpdaterSession) UpdateOperatorTable(referenceTimestamp uint32, globalTableRoot [32]byte, operatorSetIndex uint32, proof []byte, operatorTableBytes []byte) (*types.Transaction, error) {
	return _IOperatorTableUpdater.Contract.UpdateOperatorTable(&_IOperatorTableUpdater.TransactOpts, referenceTimestamp, globalTableRoot, operatorSetIndex, proof, operatorTableBytes)
}

// UpdateOperatorTable is a paid mutator transaction binding the contract method 0x9ea94778.
//
// Solidity: function updateOperatorTable(uint32 referenceTimestamp, bytes32 globalTableRoot, uint32 operatorSetIndex, bytes proof, bytes operatorTableBytes) returns()
func (_IOperatorTableUpdater *IOperatorTableUpdaterTransactorSession) UpdateOperatorTable(referenceTimestamp uint32, globalTableRoot [32]byte, operatorSetIndex uint32, proof []byte, operatorTableBytes []byte) (*types.Transaction, error) {
	return _IOperatorTableUpdater.Contract.UpdateOperatorTable(&_IOperatorTableUpdater.TransactOpts, referenceTimestamp, globalTableRoot, operatorSetIndex, proof, operatorTableBytes)
}

// IOperatorTableUpdaterGeneratorUpdatedIterator is returned from FilterGeneratorUpdated and is used to iterate over the raw logs and unpacked data for GeneratorUpdated events raised by the IOperatorTableUpdater contract.
type IOperatorTableUpdaterGeneratorUpdatedIterator struct {
	Event *IOperatorTableUpdaterGeneratorUpdated // Event containing the contract specifics and raw log

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
func (it *IOperatorTableUpdaterGeneratorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOperatorTableUpdaterGeneratorUpdated)
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
		it.Event = new(IOperatorTableUpdaterGeneratorUpdated)
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
func (it *IOperatorTableUpdaterGeneratorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOperatorTableUpdaterGeneratorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOperatorTableUpdaterGeneratorUpdated represents a GeneratorUpdated event raised by the IOperatorTableUpdater contract.
type IOperatorTableUpdaterGeneratorUpdated struct {
	OperatorSet OperatorSet
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterGeneratorUpdated is a free log retrieval operation binding the contract event 0x3463431b09dfd43dec7349f8f24acfa753fe4cf40a26235402d213373df15856.
//
// Solidity: event GeneratorUpdated((address,uint32) operatorSet)
func (_IOperatorTableUpdater *IOperatorTableUpdaterFilterer) FilterGeneratorUpdated(opts *bind.FilterOpts) (*IOperatorTableUpdaterGeneratorUpdatedIterator, error) {

	logs, sub, err := _IOperatorTableUpdater.contract.FilterLogs(opts, "GeneratorUpdated")
	if err != nil {
		return nil, err
	}
	return &IOperatorTableUpdaterGeneratorUpdatedIterator{contract: _IOperatorTableUpdater.contract, event: "GeneratorUpdated", logs: logs, sub: sub}, nil
}

// WatchGeneratorUpdated is a free log subscription operation binding the contract event 0x3463431b09dfd43dec7349f8f24acfa753fe4cf40a26235402d213373df15856.
//
// Solidity: event GeneratorUpdated((address,uint32) operatorSet)
func (_IOperatorTableUpdater *IOperatorTableUpdaterFilterer) WatchGeneratorUpdated(opts *bind.WatchOpts, sink chan<- *IOperatorTableUpdaterGeneratorUpdated) (event.Subscription, error) {

	logs, sub, err := _IOperatorTableUpdater.contract.WatchLogs(opts, "GeneratorUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOperatorTableUpdaterGeneratorUpdated)
				if err := _IOperatorTableUpdater.contract.UnpackLog(event, "GeneratorUpdated", log); err != nil {
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

// ParseGeneratorUpdated is a log parse operation binding the contract event 0x3463431b09dfd43dec7349f8f24acfa753fe4cf40a26235402d213373df15856.
//
// Solidity: event GeneratorUpdated((address,uint32) operatorSet)
func (_IOperatorTableUpdater *IOperatorTableUpdaterFilterer) ParseGeneratorUpdated(log types.Log) (*IOperatorTableUpdaterGeneratorUpdated, error) {
	event := new(IOperatorTableUpdaterGeneratorUpdated)
	if err := _IOperatorTableUpdater.contract.UnpackLog(event, "GeneratorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOperatorTableUpdaterGlobalRootConfirmationThresholdUpdatedIterator is returned from FilterGlobalRootConfirmationThresholdUpdated and is used to iterate over the raw logs and unpacked data for GlobalRootConfirmationThresholdUpdated events raised by the IOperatorTableUpdater contract.
type IOperatorTableUpdaterGlobalRootConfirmationThresholdUpdatedIterator struct {
	Event *IOperatorTableUpdaterGlobalRootConfirmationThresholdUpdated // Event containing the contract specifics and raw log

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
func (it *IOperatorTableUpdaterGlobalRootConfirmationThresholdUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOperatorTableUpdaterGlobalRootConfirmationThresholdUpdated)
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
		it.Event = new(IOperatorTableUpdaterGlobalRootConfirmationThresholdUpdated)
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
func (it *IOperatorTableUpdaterGlobalRootConfirmationThresholdUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOperatorTableUpdaterGlobalRootConfirmationThresholdUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOperatorTableUpdaterGlobalRootConfirmationThresholdUpdated represents a GlobalRootConfirmationThresholdUpdated event raised by the IOperatorTableUpdater contract.
type IOperatorTableUpdaterGlobalRootConfirmationThresholdUpdated struct {
	Bps uint16
	Raw types.Log // Blockchain specific contextual infos
}

// FilterGlobalRootConfirmationThresholdUpdated is a free log retrieval operation binding the contract event 0xf5d1836df8fcd7c1e54047e94ac8773d2855395603e2ef9ba5f5f16905f22592.
//
// Solidity: event GlobalRootConfirmationThresholdUpdated(uint16 bps)
func (_IOperatorTableUpdater *IOperatorTableUpdaterFilterer) FilterGlobalRootConfirmationThresholdUpdated(opts *bind.FilterOpts) (*IOperatorTableUpdaterGlobalRootConfirmationThresholdUpdatedIterator, error) {

	logs, sub, err := _IOperatorTableUpdater.contract.FilterLogs(opts, "GlobalRootConfirmationThresholdUpdated")
	if err != nil {
		return nil, err
	}
	return &IOperatorTableUpdaterGlobalRootConfirmationThresholdUpdatedIterator{contract: _IOperatorTableUpdater.contract, event: "GlobalRootConfirmationThresholdUpdated", logs: logs, sub: sub}, nil
}

// WatchGlobalRootConfirmationThresholdUpdated is a free log subscription operation binding the contract event 0xf5d1836df8fcd7c1e54047e94ac8773d2855395603e2ef9ba5f5f16905f22592.
//
// Solidity: event GlobalRootConfirmationThresholdUpdated(uint16 bps)
func (_IOperatorTableUpdater *IOperatorTableUpdaterFilterer) WatchGlobalRootConfirmationThresholdUpdated(opts *bind.WatchOpts, sink chan<- *IOperatorTableUpdaterGlobalRootConfirmationThresholdUpdated) (event.Subscription, error) {

	logs, sub, err := _IOperatorTableUpdater.contract.WatchLogs(opts, "GlobalRootConfirmationThresholdUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOperatorTableUpdaterGlobalRootConfirmationThresholdUpdated)
				if err := _IOperatorTableUpdater.contract.UnpackLog(event, "GlobalRootConfirmationThresholdUpdated", log); err != nil {
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
func (_IOperatorTableUpdater *IOperatorTableUpdaterFilterer) ParseGlobalRootConfirmationThresholdUpdated(log types.Log) (*IOperatorTableUpdaterGlobalRootConfirmationThresholdUpdated, error) {
	event := new(IOperatorTableUpdaterGlobalRootConfirmationThresholdUpdated)
	if err := _IOperatorTableUpdater.contract.UnpackLog(event, "GlobalRootConfirmationThresholdUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOperatorTableUpdaterGlobalRootDisabledIterator is returned from FilterGlobalRootDisabled and is used to iterate over the raw logs and unpacked data for GlobalRootDisabled events raised by the IOperatorTableUpdater contract.
type IOperatorTableUpdaterGlobalRootDisabledIterator struct {
	Event *IOperatorTableUpdaterGlobalRootDisabled // Event containing the contract specifics and raw log

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
func (it *IOperatorTableUpdaterGlobalRootDisabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOperatorTableUpdaterGlobalRootDisabled)
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
		it.Event = new(IOperatorTableUpdaterGlobalRootDisabled)
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
func (it *IOperatorTableUpdaterGlobalRootDisabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOperatorTableUpdaterGlobalRootDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOperatorTableUpdaterGlobalRootDisabled represents a GlobalRootDisabled event raised by the IOperatorTableUpdater contract.
type IOperatorTableUpdaterGlobalRootDisabled struct {
	GlobalTableRoot [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterGlobalRootDisabled is a free log retrieval operation binding the contract event 0x8bd43de1250f58fe6ec9a78671a8b78dba70f0018656d157a3aeaabec389df34.
//
// Solidity: event GlobalRootDisabled(bytes32 indexed globalTableRoot)
func (_IOperatorTableUpdater *IOperatorTableUpdaterFilterer) FilterGlobalRootDisabled(opts *bind.FilterOpts, globalTableRoot [][32]byte) (*IOperatorTableUpdaterGlobalRootDisabledIterator, error) {

	var globalTableRootRule []interface{}
	for _, globalTableRootItem := range globalTableRoot {
		globalTableRootRule = append(globalTableRootRule, globalTableRootItem)
	}

	logs, sub, err := _IOperatorTableUpdater.contract.FilterLogs(opts, "GlobalRootDisabled", globalTableRootRule)
	if err != nil {
		return nil, err
	}
	return &IOperatorTableUpdaterGlobalRootDisabledIterator{contract: _IOperatorTableUpdater.contract, event: "GlobalRootDisabled", logs: logs, sub: sub}, nil
}

// WatchGlobalRootDisabled is a free log subscription operation binding the contract event 0x8bd43de1250f58fe6ec9a78671a8b78dba70f0018656d157a3aeaabec389df34.
//
// Solidity: event GlobalRootDisabled(bytes32 indexed globalTableRoot)
func (_IOperatorTableUpdater *IOperatorTableUpdaterFilterer) WatchGlobalRootDisabled(opts *bind.WatchOpts, sink chan<- *IOperatorTableUpdaterGlobalRootDisabled, globalTableRoot [][32]byte) (event.Subscription, error) {

	var globalTableRootRule []interface{}
	for _, globalTableRootItem := range globalTableRoot {
		globalTableRootRule = append(globalTableRootRule, globalTableRootItem)
	}

	logs, sub, err := _IOperatorTableUpdater.contract.WatchLogs(opts, "GlobalRootDisabled", globalTableRootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOperatorTableUpdaterGlobalRootDisabled)
				if err := _IOperatorTableUpdater.contract.UnpackLog(event, "GlobalRootDisabled", log); err != nil {
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

// ParseGlobalRootDisabled is a log parse operation binding the contract event 0x8bd43de1250f58fe6ec9a78671a8b78dba70f0018656d157a3aeaabec389df34.
//
// Solidity: event GlobalRootDisabled(bytes32 indexed globalTableRoot)
func (_IOperatorTableUpdater *IOperatorTableUpdaterFilterer) ParseGlobalRootDisabled(log types.Log) (*IOperatorTableUpdaterGlobalRootDisabled, error) {
	event := new(IOperatorTableUpdaterGlobalRootDisabled)
	if err := _IOperatorTableUpdater.contract.UnpackLog(event, "GlobalRootDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOperatorTableUpdaterNewGlobalTableRootIterator is returned from FilterNewGlobalTableRoot and is used to iterate over the raw logs and unpacked data for NewGlobalTableRoot events raised by the IOperatorTableUpdater contract.
type IOperatorTableUpdaterNewGlobalTableRootIterator struct {
	Event *IOperatorTableUpdaterNewGlobalTableRoot // Event containing the contract specifics and raw log

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
func (it *IOperatorTableUpdaterNewGlobalTableRootIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOperatorTableUpdaterNewGlobalTableRoot)
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
		it.Event = new(IOperatorTableUpdaterNewGlobalTableRoot)
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
func (it *IOperatorTableUpdaterNewGlobalTableRootIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOperatorTableUpdaterNewGlobalTableRootIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOperatorTableUpdaterNewGlobalTableRoot represents a NewGlobalTableRoot event raised by the IOperatorTableUpdater contract.
type IOperatorTableUpdaterNewGlobalTableRoot struct {
	ReferenceTimestamp uint32
	GlobalTableRoot    [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterNewGlobalTableRoot is a free log retrieval operation binding the contract event 0x010dcbe0d1e019c93357711f7bb6287d543b7ff7de74f29df3fb5ecceec8d369.
//
// Solidity: event NewGlobalTableRoot(uint32 indexed referenceTimestamp, bytes32 indexed globalTableRoot)
func (_IOperatorTableUpdater *IOperatorTableUpdaterFilterer) FilterNewGlobalTableRoot(opts *bind.FilterOpts, referenceTimestamp []uint32, globalTableRoot [][32]byte) (*IOperatorTableUpdaterNewGlobalTableRootIterator, error) {

	var referenceTimestampRule []interface{}
	for _, referenceTimestampItem := range referenceTimestamp {
		referenceTimestampRule = append(referenceTimestampRule, referenceTimestampItem)
	}
	var globalTableRootRule []interface{}
	for _, globalTableRootItem := range globalTableRoot {
		globalTableRootRule = append(globalTableRootRule, globalTableRootItem)
	}

	logs, sub, err := _IOperatorTableUpdater.contract.FilterLogs(opts, "NewGlobalTableRoot", referenceTimestampRule, globalTableRootRule)
	if err != nil {
		return nil, err
	}
	return &IOperatorTableUpdaterNewGlobalTableRootIterator{contract: _IOperatorTableUpdater.contract, event: "NewGlobalTableRoot", logs: logs, sub: sub}, nil
}

// WatchNewGlobalTableRoot is a free log subscription operation binding the contract event 0x010dcbe0d1e019c93357711f7bb6287d543b7ff7de74f29df3fb5ecceec8d369.
//
// Solidity: event NewGlobalTableRoot(uint32 indexed referenceTimestamp, bytes32 indexed globalTableRoot)
func (_IOperatorTableUpdater *IOperatorTableUpdaterFilterer) WatchNewGlobalTableRoot(opts *bind.WatchOpts, sink chan<- *IOperatorTableUpdaterNewGlobalTableRoot, referenceTimestamp []uint32, globalTableRoot [][32]byte) (event.Subscription, error) {

	var referenceTimestampRule []interface{}
	for _, referenceTimestampItem := range referenceTimestamp {
		referenceTimestampRule = append(referenceTimestampRule, referenceTimestampItem)
	}
	var globalTableRootRule []interface{}
	for _, globalTableRootItem := range globalTableRoot {
		globalTableRootRule = append(globalTableRootRule, globalTableRootItem)
	}

	logs, sub, err := _IOperatorTableUpdater.contract.WatchLogs(opts, "NewGlobalTableRoot", referenceTimestampRule, globalTableRootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOperatorTableUpdaterNewGlobalTableRoot)
				if err := _IOperatorTableUpdater.contract.UnpackLog(event, "NewGlobalTableRoot", log); err != nil {
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
func (_IOperatorTableUpdater *IOperatorTableUpdaterFilterer) ParseNewGlobalTableRoot(log types.Log) (*IOperatorTableUpdaterNewGlobalTableRoot, error) {
	event := new(IOperatorTableUpdaterNewGlobalTableRoot)
	if err := _IOperatorTableUpdater.contract.UnpackLog(event, "NewGlobalTableRoot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
