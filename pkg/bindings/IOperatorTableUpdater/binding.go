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
	OperatorInfo      IBN254TableCalculatorTypesBN254OperatorInfo
}

// IBN254TableCalculatorTypesBN254OperatorInfo is an auto generated low-level Go binding around an user-defined struct.
type IBN254TableCalculatorTypesBN254OperatorInfo struct {
	Pubkey  BN254G1Point
	Weights []*big.Int
}

// IBN254TableCalculatorTypesBN254OperatorSetInfo is an auto generated low-level Go binding around an user-defined struct.
type IBN254TableCalculatorTypesBN254OperatorSetInfo struct {
	OperatorInfoTreeRoot [32]byte
	NumOperators         *big.Int
	AggregatePubkey      BN254G1Point
	TotalWeights         []*big.Int
}

// ICrossChainRegistryTypesOperatorSetConfig is an auto generated low-level Go binding around an user-defined struct.
type ICrossChainRegistryTypesOperatorSetConfig struct {
	Owner              common.Address
	MaxStalenessPeriod uint32
}

// IECDSATableCalculatorTypesECDSAOperatorInfo is an auto generated low-level Go binding around an user-defined struct.
type IECDSATableCalculatorTypesECDSAOperatorInfo struct {
	Pubkey  common.Address
	Weights []*big.Int
}

// OperatorSet is an auto generated low-level Go binding around an user-defined struct.
type OperatorSet struct {
	Avs common.Address
	Id  uint32
}

// IOperatorTableUpdaterMetaData contains all meta data concerning the IOperatorTableUpdater contract.
var IOperatorTableUpdaterMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"confirmGlobalTableRoot\",\"inputs\":[{\"name\":\"globalTableRootCert\",\"type\":\"tuple\",\"internalType\":\"structIBN254CertificateVerifierTypes.BN254Certificate\",\"components\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"nonSignerWitnesses\",\"type\":\"tuple[]\",\"internalType\":\"structIBN254CertificateVerifierTypes.BN254OperatorInfoWitness[]\",\"components\":[{\"name\":\"operatorIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorInfoProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"operatorInfo\",\"type\":\"tuple\",\"internalType\":\"structIBN254TableCalculatorTypes.BN254OperatorInfo\",\"components\":[{\"name\":\"pubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}]}]},{\"name\":\"globalTableRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getCurrentGlobalTableRoot\",\"inputs\":[],\"outputs\":[{\"name\":\"globalTableRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTableRootByTimestamp\",\"inputs\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"tableRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setGlobalRootConfirmationThreshold\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGlobalRootConfirmerSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateBN254OperatorTable\",\"inputs\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"globalTableRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"operatorSetIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operatorSetInfo\",\"type\":\"tuple\",\"internalType\":\"structIBN254TableCalculatorTypes.BN254OperatorSetInfo\",\"components\":[{\"name\":\"operatorInfoTreeRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"numOperators\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"aggregatePubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"totalWeights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]},{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structICrossChainRegistryTypes.OperatorSetConfig\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxStalenessPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateECDSAOperatorTable\",\"inputs\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"globalTableRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"operatorSetIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operatorInfos\",\"type\":\"tuple[]\",\"internalType\":\"structIECDSATableCalculatorTypes.ECDSAOperatorInfo[]\",\"components\":[{\"name\":\"pubkey\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"weights\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}]},{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structICrossChainRegistryTypes.OperatorSetConfig\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxStalenessPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"NewglobalTableRoot\",\"inputs\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"globalTableRoot\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"GlobalTableRootUpdateFailed\",\"inputs\":[]}]",
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

// GetTableRootByTimestamp is a free data retrieval call binding the contract method 0x9ea8dbce.
//
// Solidity: function getTableRootByTimestamp(uint32 referenceTimestamp) view returns(bytes32 tableRoot)
func (_IOperatorTableUpdater *IOperatorTableUpdaterCaller) GetTableRootByTimestamp(opts *bind.CallOpts, referenceTimestamp uint32) ([32]byte, error) {
	var out []interface{}
	err := _IOperatorTableUpdater.contract.Call(opts, &out, "getTableRootByTimestamp", referenceTimestamp)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetTableRootByTimestamp is a free data retrieval call binding the contract method 0x9ea8dbce.
//
// Solidity: function getTableRootByTimestamp(uint32 referenceTimestamp) view returns(bytes32 tableRoot)
func (_IOperatorTableUpdater *IOperatorTableUpdaterSession) GetTableRootByTimestamp(referenceTimestamp uint32) ([32]byte, error) {
	return _IOperatorTableUpdater.Contract.GetTableRootByTimestamp(&_IOperatorTableUpdater.CallOpts, referenceTimestamp)
}

// GetTableRootByTimestamp is a free data retrieval call binding the contract method 0x9ea8dbce.
//
// Solidity: function getTableRootByTimestamp(uint32 referenceTimestamp) view returns(bytes32 tableRoot)
func (_IOperatorTableUpdater *IOperatorTableUpdaterCallerSession) GetTableRootByTimestamp(referenceTimestamp uint32) ([32]byte, error) {
	return _IOperatorTableUpdater.Contract.GetTableRootByTimestamp(&_IOperatorTableUpdater.CallOpts, referenceTimestamp)
}

// ConfirmGlobalTableRoot is a paid mutator transaction binding the contract method 0x6ab40904.
//
// Solidity: function confirmGlobalTableRoot((uint32,bytes32,(uint256,uint256),(uint256[2],uint256[2]),(uint32,bytes,((uint256,uint256),uint256[]))[]) globalTableRootCert, bytes32 globalTableRoot, uint32 referenceTimestamp) returns()
func (_IOperatorTableUpdater *IOperatorTableUpdaterTransactor) ConfirmGlobalTableRoot(opts *bind.TransactOpts, globalTableRootCert IBN254CertificateVerifierTypesBN254Certificate, globalTableRoot [32]byte, referenceTimestamp uint32) (*types.Transaction, error) {
	return _IOperatorTableUpdater.contract.Transact(opts, "confirmGlobalTableRoot", globalTableRootCert, globalTableRoot, referenceTimestamp)
}

// ConfirmGlobalTableRoot is a paid mutator transaction binding the contract method 0x6ab40904.
//
// Solidity: function confirmGlobalTableRoot((uint32,bytes32,(uint256,uint256),(uint256[2],uint256[2]),(uint32,bytes,((uint256,uint256),uint256[]))[]) globalTableRootCert, bytes32 globalTableRoot, uint32 referenceTimestamp) returns()
func (_IOperatorTableUpdater *IOperatorTableUpdaterSession) ConfirmGlobalTableRoot(globalTableRootCert IBN254CertificateVerifierTypesBN254Certificate, globalTableRoot [32]byte, referenceTimestamp uint32) (*types.Transaction, error) {
	return _IOperatorTableUpdater.Contract.ConfirmGlobalTableRoot(&_IOperatorTableUpdater.TransactOpts, globalTableRootCert, globalTableRoot, referenceTimestamp)
}

// ConfirmGlobalTableRoot is a paid mutator transaction binding the contract method 0x6ab40904.
//
// Solidity: function confirmGlobalTableRoot((uint32,bytes32,(uint256,uint256),(uint256[2],uint256[2]),(uint32,bytes,((uint256,uint256),uint256[]))[]) globalTableRootCert, bytes32 globalTableRoot, uint32 referenceTimestamp) returns()
func (_IOperatorTableUpdater *IOperatorTableUpdaterTransactorSession) ConfirmGlobalTableRoot(globalTableRootCert IBN254CertificateVerifierTypesBN254Certificate, globalTableRoot [32]byte, referenceTimestamp uint32) (*types.Transaction, error) {
	return _IOperatorTableUpdater.Contract.ConfirmGlobalTableRoot(&_IOperatorTableUpdater.TransactOpts, globalTableRootCert, globalTableRoot, referenceTimestamp)
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

// SetGlobalRootConfirmerSet is a paid mutator transaction binding the contract method 0x0371406e.
//
// Solidity: function setGlobalRootConfirmerSet((address,uint32) operatorSet) returns()
func (_IOperatorTableUpdater *IOperatorTableUpdaterTransactor) SetGlobalRootConfirmerSet(opts *bind.TransactOpts, operatorSet OperatorSet) (*types.Transaction, error) {
	return _IOperatorTableUpdater.contract.Transact(opts, "setGlobalRootConfirmerSet", operatorSet)
}

// SetGlobalRootConfirmerSet is a paid mutator transaction binding the contract method 0x0371406e.
//
// Solidity: function setGlobalRootConfirmerSet((address,uint32) operatorSet) returns()
func (_IOperatorTableUpdater *IOperatorTableUpdaterSession) SetGlobalRootConfirmerSet(operatorSet OperatorSet) (*types.Transaction, error) {
	return _IOperatorTableUpdater.Contract.SetGlobalRootConfirmerSet(&_IOperatorTableUpdater.TransactOpts, operatorSet)
}

// SetGlobalRootConfirmerSet is a paid mutator transaction binding the contract method 0x0371406e.
//
// Solidity: function setGlobalRootConfirmerSet((address,uint32) operatorSet) returns()
func (_IOperatorTableUpdater *IOperatorTableUpdaterTransactorSession) SetGlobalRootConfirmerSet(operatorSet OperatorSet) (*types.Transaction, error) {
	return _IOperatorTableUpdater.Contract.SetGlobalRootConfirmerSet(&_IOperatorTableUpdater.TransactOpts, operatorSet)
}

// UpdateBN254OperatorTable is a paid mutator transaction binding the contract method 0x71034fe1.
//
// Solidity: function updateBN254OperatorTable(uint32 referenceTimestamp, bytes32 globalTableRoot, uint32 operatorSetIndex, bytes proof, (address,uint32) operatorSet, (bytes32,uint256,(uint256,uint256),uint256[]) operatorSetInfo, (address,uint32) config) returns()
func (_IOperatorTableUpdater *IOperatorTableUpdaterTransactor) UpdateBN254OperatorTable(opts *bind.TransactOpts, referenceTimestamp uint32, globalTableRoot [32]byte, operatorSetIndex uint32, proof []byte, operatorSet OperatorSet, operatorSetInfo IBN254TableCalculatorTypesBN254OperatorSetInfo, config ICrossChainRegistryTypesOperatorSetConfig) (*types.Transaction, error) {
	return _IOperatorTableUpdater.contract.Transact(opts, "updateBN254OperatorTable", referenceTimestamp, globalTableRoot, operatorSetIndex, proof, operatorSet, operatorSetInfo, config)
}

// UpdateBN254OperatorTable is a paid mutator transaction binding the contract method 0x71034fe1.
//
// Solidity: function updateBN254OperatorTable(uint32 referenceTimestamp, bytes32 globalTableRoot, uint32 operatorSetIndex, bytes proof, (address,uint32) operatorSet, (bytes32,uint256,(uint256,uint256),uint256[]) operatorSetInfo, (address,uint32) config) returns()
func (_IOperatorTableUpdater *IOperatorTableUpdaterSession) UpdateBN254OperatorTable(referenceTimestamp uint32, globalTableRoot [32]byte, operatorSetIndex uint32, proof []byte, operatorSet OperatorSet, operatorSetInfo IBN254TableCalculatorTypesBN254OperatorSetInfo, config ICrossChainRegistryTypesOperatorSetConfig) (*types.Transaction, error) {
	return _IOperatorTableUpdater.Contract.UpdateBN254OperatorTable(&_IOperatorTableUpdater.TransactOpts, referenceTimestamp, globalTableRoot, operatorSetIndex, proof, operatorSet, operatorSetInfo, config)
}

// UpdateBN254OperatorTable is a paid mutator transaction binding the contract method 0x71034fe1.
//
// Solidity: function updateBN254OperatorTable(uint32 referenceTimestamp, bytes32 globalTableRoot, uint32 operatorSetIndex, bytes proof, (address,uint32) operatorSet, (bytes32,uint256,(uint256,uint256),uint256[]) operatorSetInfo, (address,uint32) config) returns()
func (_IOperatorTableUpdater *IOperatorTableUpdaterTransactorSession) UpdateBN254OperatorTable(referenceTimestamp uint32, globalTableRoot [32]byte, operatorSetIndex uint32, proof []byte, operatorSet OperatorSet, operatorSetInfo IBN254TableCalculatorTypesBN254OperatorSetInfo, config ICrossChainRegistryTypesOperatorSetConfig) (*types.Transaction, error) {
	return _IOperatorTableUpdater.Contract.UpdateBN254OperatorTable(&_IOperatorTableUpdater.TransactOpts, referenceTimestamp, globalTableRoot, operatorSetIndex, proof, operatorSet, operatorSetInfo, config)
}

// UpdateECDSAOperatorTable is a paid mutator transaction binding the contract method 0x3397579c.
//
// Solidity: function updateECDSAOperatorTable(uint32 referenceTimestamp, bytes32 globalTableRoot, uint32 operatorSetIndex, bytes proof, (address,uint32) operatorSet, (address,uint96[])[] operatorInfos, (address,uint32) config) returns()
func (_IOperatorTableUpdater *IOperatorTableUpdaterTransactor) UpdateECDSAOperatorTable(opts *bind.TransactOpts, referenceTimestamp uint32, globalTableRoot [32]byte, operatorSetIndex uint32, proof []byte, operatorSet OperatorSet, operatorInfos []IECDSATableCalculatorTypesECDSAOperatorInfo, config ICrossChainRegistryTypesOperatorSetConfig) (*types.Transaction, error) {
	return _IOperatorTableUpdater.contract.Transact(opts, "updateECDSAOperatorTable", referenceTimestamp, globalTableRoot, operatorSetIndex, proof, operatorSet, operatorInfos, config)
}

// UpdateECDSAOperatorTable is a paid mutator transaction binding the contract method 0x3397579c.
//
// Solidity: function updateECDSAOperatorTable(uint32 referenceTimestamp, bytes32 globalTableRoot, uint32 operatorSetIndex, bytes proof, (address,uint32) operatorSet, (address,uint96[])[] operatorInfos, (address,uint32) config) returns()
func (_IOperatorTableUpdater *IOperatorTableUpdaterSession) UpdateECDSAOperatorTable(referenceTimestamp uint32, globalTableRoot [32]byte, operatorSetIndex uint32, proof []byte, operatorSet OperatorSet, operatorInfos []IECDSATableCalculatorTypesECDSAOperatorInfo, config ICrossChainRegistryTypesOperatorSetConfig) (*types.Transaction, error) {
	return _IOperatorTableUpdater.Contract.UpdateECDSAOperatorTable(&_IOperatorTableUpdater.TransactOpts, referenceTimestamp, globalTableRoot, operatorSetIndex, proof, operatorSet, operatorInfos, config)
}

// UpdateECDSAOperatorTable is a paid mutator transaction binding the contract method 0x3397579c.
//
// Solidity: function updateECDSAOperatorTable(uint32 referenceTimestamp, bytes32 globalTableRoot, uint32 operatorSetIndex, bytes proof, (address,uint32) operatorSet, (address,uint96[])[] operatorInfos, (address,uint32) config) returns()
func (_IOperatorTableUpdater *IOperatorTableUpdaterTransactorSession) UpdateECDSAOperatorTable(referenceTimestamp uint32, globalTableRoot [32]byte, operatorSetIndex uint32, proof []byte, operatorSet OperatorSet, operatorInfos []IECDSATableCalculatorTypesECDSAOperatorInfo, config ICrossChainRegistryTypesOperatorSetConfig) (*types.Transaction, error) {
	return _IOperatorTableUpdater.Contract.UpdateECDSAOperatorTable(&_IOperatorTableUpdater.TransactOpts, referenceTimestamp, globalTableRoot, operatorSetIndex, proof, operatorSet, operatorInfos, config)
}

// IOperatorTableUpdaterNewglobalTableRootIterator is returned from FilterNewglobalTableRoot and is used to iterate over the raw logs and unpacked data for NewglobalTableRoot events raised by the IOperatorTableUpdater contract.
type IOperatorTableUpdaterNewglobalTableRootIterator struct {
	Event *IOperatorTableUpdaterNewglobalTableRoot // Event containing the contract specifics and raw log

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
func (it *IOperatorTableUpdaterNewglobalTableRootIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOperatorTableUpdaterNewglobalTableRoot)
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
		it.Event = new(IOperatorTableUpdaterNewglobalTableRoot)
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
func (it *IOperatorTableUpdaterNewglobalTableRootIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOperatorTableUpdaterNewglobalTableRootIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOperatorTableUpdaterNewglobalTableRoot represents a NewglobalTableRoot event raised by the IOperatorTableUpdater contract.
type IOperatorTableUpdaterNewglobalTableRoot struct {
	ReferenceTimestamp uint32
	GlobalTableRoot    [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterNewglobalTableRoot is a free log retrieval operation binding the contract event 0xbb96ecabc83b99776c3ba3d12d8131e99aeca109d673db36e42959d154359a65.
//
// Solidity: event NewglobalTableRoot(uint32 referenceTimestamp, bytes32 globalTableRoot)
func (_IOperatorTableUpdater *IOperatorTableUpdaterFilterer) FilterNewglobalTableRoot(opts *bind.FilterOpts) (*IOperatorTableUpdaterNewglobalTableRootIterator, error) {

	logs, sub, err := _IOperatorTableUpdater.contract.FilterLogs(opts, "NewglobalTableRoot")
	if err != nil {
		return nil, err
	}
	return &IOperatorTableUpdaterNewglobalTableRootIterator{contract: _IOperatorTableUpdater.contract, event: "NewglobalTableRoot", logs: logs, sub: sub}, nil
}

// WatchNewglobalTableRoot is a free log subscription operation binding the contract event 0xbb96ecabc83b99776c3ba3d12d8131e99aeca109d673db36e42959d154359a65.
//
// Solidity: event NewglobalTableRoot(uint32 referenceTimestamp, bytes32 globalTableRoot)
func (_IOperatorTableUpdater *IOperatorTableUpdaterFilterer) WatchNewglobalTableRoot(opts *bind.WatchOpts, sink chan<- *IOperatorTableUpdaterNewglobalTableRoot) (event.Subscription, error) {

	logs, sub, err := _IOperatorTableUpdater.contract.WatchLogs(opts, "NewglobalTableRoot")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOperatorTableUpdaterNewglobalTableRoot)
				if err := _IOperatorTableUpdater.contract.UnpackLog(event, "NewglobalTableRoot", log); err != nil {
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

// ParseNewglobalTableRoot is a log parse operation binding the contract event 0xbb96ecabc83b99776c3ba3d12d8131e99aeca109d673db36e42959d154359a65.
//
// Solidity: event NewglobalTableRoot(uint32 referenceTimestamp, bytes32 globalTableRoot)
func (_IOperatorTableUpdater *IOperatorTableUpdaterFilterer) ParseNewglobalTableRoot(log types.Log) (*IOperatorTableUpdaterNewglobalTableRoot, error) {
	event := new(IOperatorTableUpdaterNewglobalTableRoot)
	if err := _IOperatorTableUpdater.contract.UnpackLog(event, "NewglobalTableRoot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
