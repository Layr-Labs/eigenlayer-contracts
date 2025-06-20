// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package BN254CertificateVerifierStorage

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

// OperatorSet is an auto generated low-level Go binding around an user-defined struct.
type OperatorSet struct {
	Avs common.Address
	Id  uint32
}

// BN254CertificateVerifierStorageMetaData contains all meta data concerning the BN254CertificateVerifierStorage contract.
var BN254CertificateVerifierStorageMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"getNonsignerOperatorInfo\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBN254TableCalculatorTypes.BN254OperatorInfo\",\"components\":[{\"name\":\"pubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorSetInfo\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBN254TableCalculatorTypes.BN254OperatorSetInfo\",\"components\":[{\"name\":\"operatorInfoTreeRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"numOperators\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"aggregatePubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"totalWeights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorSetOwner\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isNonsignerCached\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestReferenceTimestamp\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxOperatorTableStaleness\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorTableUpdater\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIOperatorTableUpdater\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"trySignatureVerification\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"aggPubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"signature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"pairingSuccessful\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"signatureValid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateOperatorTable\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorSetInfo\",\"type\":\"tuple\",\"internalType\":\"structIBN254TableCalculatorTypes.BN254OperatorSetInfo\",\"components\":[{\"name\":\"operatorInfoTreeRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"numOperators\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"aggregatePubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"totalWeights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]},{\"name\":\"operatorSetConfig\",\"type\":\"tuple\",\"internalType\":\"structICrossChainRegistryTypes.OperatorSetConfig\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxStalenessPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyCertificate\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"cert\",\"type\":\"tuple\",\"internalType\":\"structIBN254CertificateVerifierTypes.BN254Certificate\",\"components\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"nonSignerWitnesses\",\"type\":\"tuple[]\",\"internalType\":\"structIBN254CertificateVerifierTypes.BN254OperatorInfoWitness[]\",\"components\":[{\"name\":\"operatorIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorInfoProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"operatorInfo\",\"type\":\"tuple\",\"internalType\":\"structIBN254TableCalculatorTypes.BN254OperatorInfo\",\"components\":[{\"name\":\"pubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}]}]}],\"outputs\":[{\"name\":\"signedStakes\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyCertificateNominal\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"cert\",\"type\":\"tuple\",\"internalType\":\"structIBN254CertificateVerifierTypes.BN254Certificate\",\"components\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"nonSignerWitnesses\",\"type\":\"tuple[]\",\"internalType\":\"structIBN254CertificateVerifierTypes.BN254OperatorInfoWitness[]\",\"components\":[{\"name\":\"operatorIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorInfoProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"operatorInfo\",\"type\":\"tuple\",\"internalType\":\"structIBN254TableCalculatorTypes.BN254OperatorInfo\",\"components\":[{\"name\":\"pubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}]}]},{\"name\":\"totalStakeNominalThresholds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyCertificateProportion\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"cert\",\"type\":\"tuple\",\"internalType\":\"structIBN254CertificateVerifierTypes.BN254Certificate\",\"components\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"nonSignerWitnesses\",\"type\":\"tuple[]\",\"internalType\":\"structIBN254CertificateVerifierTypes.BN254OperatorInfoWitness[]\",\"components\":[{\"name\":\"operatorIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorInfoProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"operatorInfo\",\"type\":\"tuple\",\"internalType\":\"structIBN254TableCalculatorTypes.BN254OperatorInfo\",\"components\":[{\"name\":\"pubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}]}]},{\"name\":\"totalStakeProportionThresholds\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"MaxStalenessPeriodUpdated\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"maxStalenessPeriod\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSetOwnerUpdated\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TableUpdated\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"operatorSetInfo\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIBN254TableCalculatorTypes.BN254OperatorSetInfo\",\"components\":[{\"name\":\"operatorInfoTreeRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"numOperators\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"aggregatePubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"totalWeights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ArrayLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CertificateStale\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperatorIndex\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyTableUpdater\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReferenceTimestampDoesNotExist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TableUpdateStale\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"VerificationFailed\",\"inputs\":[]}]",
}

// BN254CertificateVerifierStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use BN254CertificateVerifierStorageMetaData.ABI instead.
var BN254CertificateVerifierStorageABI = BN254CertificateVerifierStorageMetaData.ABI

// BN254CertificateVerifierStorage is an auto generated Go binding around an Ethereum contract.
type BN254CertificateVerifierStorage struct {
	BN254CertificateVerifierStorageCaller     // Read-only binding to the contract
	BN254CertificateVerifierStorageTransactor // Write-only binding to the contract
	BN254CertificateVerifierStorageFilterer   // Log filterer for contract events
}

// BN254CertificateVerifierStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type BN254CertificateVerifierStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BN254CertificateVerifierStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BN254CertificateVerifierStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BN254CertificateVerifierStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BN254CertificateVerifierStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BN254CertificateVerifierStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BN254CertificateVerifierStorageSession struct {
	Contract     *BN254CertificateVerifierStorage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                    // Call options to use throughout this session
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// BN254CertificateVerifierStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BN254CertificateVerifierStorageCallerSession struct {
	Contract *BN254CertificateVerifierStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                          // Call options to use throughout this session
}

// BN254CertificateVerifierStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BN254CertificateVerifierStorageTransactorSession struct {
	Contract     *BN254CertificateVerifierStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                          // Transaction auth options to use throughout this session
}

// BN254CertificateVerifierStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type BN254CertificateVerifierStorageRaw struct {
	Contract *BN254CertificateVerifierStorage // Generic contract binding to access the raw methods on
}

// BN254CertificateVerifierStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BN254CertificateVerifierStorageCallerRaw struct {
	Contract *BN254CertificateVerifierStorageCaller // Generic read-only contract binding to access the raw methods on
}

// BN254CertificateVerifierStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BN254CertificateVerifierStorageTransactorRaw struct {
	Contract *BN254CertificateVerifierStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBN254CertificateVerifierStorage creates a new instance of BN254CertificateVerifierStorage, bound to a specific deployed contract.
func NewBN254CertificateVerifierStorage(address common.Address, backend bind.ContractBackend) (*BN254CertificateVerifierStorage, error) {
	contract, err := bindBN254CertificateVerifierStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BN254CertificateVerifierStorage{BN254CertificateVerifierStorageCaller: BN254CertificateVerifierStorageCaller{contract: contract}, BN254CertificateVerifierStorageTransactor: BN254CertificateVerifierStorageTransactor{contract: contract}, BN254CertificateVerifierStorageFilterer: BN254CertificateVerifierStorageFilterer{contract: contract}}, nil
}

// NewBN254CertificateVerifierStorageCaller creates a new read-only instance of BN254CertificateVerifierStorage, bound to a specific deployed contract.
func NewBN254CertificateVerifierStorageCaller(address common.Address, caller bind.ContractCaller) (*BN254CertificateVerifierStorageCaller, error) {
	contract, err := bindBN254CertificateVerifierStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BN254CertificateVerifierStorageCaller{contract: contract}, nil
}

// NewBN254CertificateVerifierStorageTransactor creates a new write-only instance of BN254CertificateVerifierStorage, bound to a specific deployed contract.
func NewBN254CertificateVerifierStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*BN254CertificateVerifierStorageTransactor, error) {
	contract, err := bindBN254CertificateVerifierStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BN254CertificateVerifierStorageTransactor{contract: contract}, nil
}

// NewBN254CertificateVerifierStorageFilterer creates a new log filterer instance of BN254CertificateVerifierStorage, bound to a specific deployed contract.
func NewBN254CertificateVerifierStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*BN254CertificateVerifierStorageFilterer, error) {
	contract, err := bindBN254CertificateVerifierStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BN254CertificateVerifierStorageFilterer{contract: contract}, nil
}

// bindBN254CertificateVerifierStorage binds a generic wrapper to an already deployed contract.
func bindBN254CertificateVerifierStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BN254CertificateVerifierStorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BN254CertificateVerifierStorage *BN254CertificateVerifierStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BN254CertificateVerifierStorage.Contract.BN254CertificateVerifierStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BN254CertificateVerifierStorage *BN254CertificateVerifierStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BN254CertificateVerifierStorage.Contract.BN254CertificateVerifierStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BN254CertificateVerifierStorage *BN254CertificateVerifierStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BN254CertificateVerifierStorage.Contract.BN254CertificateVerifierStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BN254CertificateVerifierStorage *BN254CertificateVerifierStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BN254CertificateVerifierStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BN254CertificateVerifierStorage *BN254CertificateVerifierStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BN254CertificateVerifierStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BN254CertificateVerifierStorage *BN254CertificateVerifierStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BN254CertificateVerifierStorage.Contract.contract.Transact(opts, method, params...)
}

// GetNonsignerOperatorInfo is a free data retrieval call binding the contract method 0x26af6a3c.
//
// Solidity: function getNonsignerOperatorInfo((address,uint32) operatorSet, uint32 referenceTimestamp, uint256 operatorIndex) view returns(((uint256,uint256),uint256[]))
func (_BN254CertificateVerifierStorage *BN254CertificateVerifierStorageCaller) GetNonsignerOperatorInfo(opts *bind.CallOpts, operatorSet OperatorSet, referenceTimestamp uint32, operatorIndex *big.Int) (IBN254TableCalculatorTypesBN254OperatorInfo, error) {
	var out []interface{}
	err := _BN254CertificateVerifierStorage.contract.Call(opts, &out, "getNonsignerOperatorInfo", operatorSet, referenceTimestamp, operatorIndex)

	if err != nil {
		return *new(IBN254TableCalculatorTypesBN254OperatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IBN254TableCalculatorTypesBN254OperatorInfo)).(*IBN254TableCalculatorTypesBN254OperatorInfo)

	return out0, err

}

// GetNonsignerOperatorInfo is a free data retrieval call binding the contract method 0x26af6a3c.
//
// Solidity: function getNonsignerOperatorInfo((address,uint32) operatorSet, uint32 referenceTimestamp, uint256 operatorIndex) view returns(((uint256,uint256),uint256[]))
func (_BN254CertificateVerifierStorage *BN254CertificateVerifierStorageSession) GetNonsignerOperatorInfo(operatorSet OperatorSet, referenceTimestamp uint32, operatorIndex *big.Int) (IBN254TableCalculatorTypesBN254OperatorInfo, error) {
	return _BN254CertificateVerifierStorage.Contract.GetNonsignerOperatorInfo(&_BN254CertificateVerifierStorage.CallOpts, operatorSet, referenceTimestamp, operatorIndex)
}

// GetNonsignerOperatorInfo is a free data retrieval call binding the contract method 0x26af6a3c.
//
// Solidity: function getNonsignerOperatorInfo((address,uint32) operatorSet, uint32 referenceTimestamp, uint256 operatorIndex) view returns(((uint256,uint256),uint256[]))
func (_BN254CertificateVerifierStorage *BN254CertificateVerifierStorageCallerSession) GetNonsignerOperatorInfo(operatorSet OperatorSet, referenceTimestamp uint32, operatorIndex *big.Int) (IBN254TableCalculatorTypesBN254OperatorInfo, error) {
	return _BN254CertificateVerifierStorage.Contract.GetNonsignerOperatorInfo(&_BN254CertificateVerifierStorage.CallOpts, operatorSet, referenceTimestamp, operatorIndex)
}

// GetOperatorSetInfo is a free data retrieval call binding the contract method 0xeb39e68f.
//
// Solidity: function getOperatorSetInfo((address,uint32) operatorSet, uint32 referenceTimestamp) view returns((bytes32,uint256,(uint256,uint256),uint256[]))
func (_BN254CertificateVerifierStorage *BN254CertificateVerifierStorageCaller) GetOperatorSetInfo(opts *bind.CallOpts, operatorSet OperatorSet, referenceTimestamp uint32) (IBN254TableCalculatorTypesBN254OperatorSetInfo, error) {
	var out []interface{}
	err := _BN254CertificateVerifierStorage.contract.Call(opts, &out, "getOperatorSetInfo", operatorSet, referenceTimestamp)

	if err != nil {
		return *new(IBN254TableCalculatorTypesBN254OperatorSetInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IBN254TableCalculatorTypesBN254OperatorSetInfo)).(*IBN254TableCalculatorTypesBN254OperatorSetInfo)

	return out0, err

}

// GetOperatorSetInfo is a free data retrieval call binding the contract method 0xeb39e68f.
//
// Solidity: function getOperatorSetInfo((address,uint32) operatorSet, uint32 referenceTimestamp) view returns((bytes32,uint256,(uint256,uint256),uint256[]))
func (_BN254CertificateVerifierStorage *BN254CertificateVerifierStorageSession) GetOperatorSetInfo(operatorSet OperatorSet, referenceTimestamp uint32) (IBN254TableCalculatorTypesBN254OperatorSetInfo, error) {
	return _BN254CertificateVerifierStorage.Contract.GetOperatorSetInfo(&_BN254CertificateVerifierStorage.CallOpts, operatorSet, referenceTimestamp)
}

// GetOperatorSetInfo is a free data retrieval call binding the contract method 0xeb39e68f.
//
// Solidity: function getOperatorSetInfo((address,uint32) operatorSet, uint32 referenceTimestamp) view returns((bytes32,uint256,(uint256,uint256),uint256[]))
func (_BN254CertificateVerifierStorage *BN254CertificateVerifierStorageCallerSession) GetOperatorSetInfo(operatorSet OperatorSet, referenceTimestamp uint32) (IBN254TableCalculatorTypesBN254OperatorSetInfo, error) {
	return _BN254CertificateVerifierStorage.Contract.GetOperatorSetInfo(&_BN254CertificateVerifierStorage.CallOpts, operatorSet, referenceTimestamp)
}

// GetOperatorSetOwner is a free data retrieval call binding the contract method 0x84818920.
//
// Solidity: function getOperatorSetOwner((address,uint32) operatorSet) view returns(address)
func (_BN254CertificateVerifierStorage *BN254CertificateVerifierStorageCaller) GetOperatorSetOwner(opts *bind.CallOpts, operatorSet OperatorSet) (common.Address, error) {
	var out []interface{}
	err := _BN254CertificateVerifierStorage.contract.Call(opts, &out, "getOperatorSetOwner", operatorSet)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOperatorSetOwner is a free data retrieval call binding the contract method 0x84818920.
//
// Solidity: function getOperatorSetOwner((address,uint32) operatorSet) view returns(address)
func (_BN254CertificateVerifierStorage *BN254CertificateVerifierStorageSession) GetOperatorSetOwner(operatorSet OperatorSet) (common.Address, error) {
	return _BN254CertificateVerifierStorage.Contract.GetOperatorSetOwner(&_BN254CertificateVerifierStorage.CallOpts, operatorSet)
}

// GetOperatorSetOwner is a free data retrieval call binding the contract method 0x84818920.
//
// Solidity: function getOperatorSetOwner((address,uint32) operatorSet) view returns(address)
func (_BN254CertificateVerifierStorage *BN254CertificateVerifierStorageCallerSession) GetOperatorSetOwner(operatorSet OperatorSet) (common.Address, error) {
	return _BN254CertificateVerifierStorage.Contract.GetOperatorSetOwner(&_BN254CertificateVerifierStorage.CallOpts, operatorSet)
}

// IsNonsignerCached is a free data retrieval call binding the contract method 0x5be87274.
//
// Solidity: function isNonsignerCached((address,uint32) operatorSet, uint32 referenceTimestamp, uint256 operatorIndex) view returns(bool)
func (_BN254CertificateVerifierStorage *BN254CertificateVerifierStorageCaller) IsNonsignerCached(opts *bind.CallOpts, operatorSet OperatorSet, referenceTimestamp uint32, operatorIndex *big.Int) (bool, error) {
	var out []interface{}
	err := _BN254CertificateVerifierStorage.contract.Call(opts, &out, "isNonsignerCached", operatorSet, referenceTimestamp, operatorIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsNonsignerCached is a free data retrieval call binding the contract method 0x5be87274.
//
// Solidity: function isNonsignerCached((address,uint32) operatorSet, uint32 referenceTimestamp, uint256 operatorIndex) view returns(bool)
func (_BN254CertificateVerifierStorage *BN254CertificateVerifierStorageSession) IsNonsignerCached(operatorSet OperatorSet, referenceTimestamp uint32, operatorIndex *big.Int) (bool, error) {
	return _BN254CertificateVerifierStorage.Contract.IsNonsignerCached(&_BN254CertificateVerifierStorage.CallOpts, operatorSet, referenceTimestamp, operatorIndex)
}

// IsNonsignerCached is a free data retrieval call binding the contract method 0x5be87274.
//
// Solidity: function isNonsignerCached((address,uint32) operatorSet, uint32 referenceTimestamp, uint256 operatorIndex) view returns(bool)
func (_BN254CertificateVerifierStorage *BN254CertificateVerifierStorageCallerSession) IsNonsignerCached(operatorSet OperatorSet, referenceTimestamp uint32, operatorIndex *big.Int) (bool, error) {
	return _BN254CertificateVerifierStorage.Contract.IsNonsignerCached(&_BN254CertificateVerifierStorage.CallOpts, operatorSet, referenceTimestamp, operatorIndex)
}

// LatestReferenceTimestamp is a free data retrieval call binding the contract method 0x5ddb9b5b.
//
// Solidity: function latestReferenceTimestamp((address,uint32) operatorSet) view returns(uint32)
func (_BN254CertificateVerifierStorage *BN254CertificateVerifierStorageCaller) LatestReferenceTimestamp(opts *bind.CallOpts, operatorSet OperatorSet) (uint32, error) {
	var out []interface{}
	err := _BN254CertificateVerifierStorage.contract.Call(opts, &out, "latestReferenceTimestamp", operatorSet)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LatestReferenceTimestamp is a free data retrieval call binding the contract method 0x5ddb9b5b.
//
// Solidity: function latestReferenceTimestamp((address,uint32) operatorSet) view returns(uint32)
func (_BN254CertificateVerifierStorage *BN254CertificateVerifierStorageSession) LatestReferenceTimestamp(operatorSet OperatorSet) (uint32, error) {
	return _BN254CertificateVerifierStorage.Contract.LatestReferenceTimestamp(&_BN254CertificateVerifierStorage.CallOpts, operatorSet)
}

// LatestReferenceTimestamp is a free data retrieval call binding the contract method 0x5ddb9b5b.
//
// Solidity: function latestReferenceTimestamp((address,uint32) operatorSet) view returns(uint32)
func (_BN254CertificateVerifierStorage *BN254CertificateVerifierStorageCallerSession) LatestReferenceTimestamp(operatorSet OperatorSet) (uint32, error) {
	return _BN254CertificateVerifierStorage.Contract.LatestReferenceTimestamp(&_BN254CertificateVerifierStorage.CallOpts, operatorSet)
}

// MaxOperatorTableStaleness is a free data retrieval call binding the contract method 0x6141879e.
//
// Solidity: function maxOperatorTableStaleness((address,uint32) operatorSet) view returns(uint32)
func (_BN254CertificateVerifierStorage *BN254CertificateVerifierStorageCaller) MaxOperatorTableStaleness(opts *bind.CallOpts, operatorSet OperatorSet) (uint32, error) {
	var out []interface{}
	err := _BN254CertificateVerifierStorage.contract.Call(opts, &out, "maxOperatorTableStaleness", operatorSet)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MaxOperatorTableStaleness is a free data retrieval call binding the contract method 0x6141879e.
//
// Solidity: function maxOperatorTableStaleness((address,uint32) operatorSet) view returns(uint32)
func (_BN254CertificateVerifierStorage *BN254CertificateVerifierStorageSession) MaxOperatorTableStaleness(operatorSet OperatorSet) (uint32, error) {
	return _BN254CertificateVerifierStorage.Contract.MaxOperatorTableStaleness(&_BN254CertificateVerifierStorage.CallOpts, operatorSet)
}

// MaxOperatorTableStaleness is a free data retrieval call binding the contract method 0x6141879e.
//
// Solidity: function maxOperatorTableStaleness((address,uint32) operatorSet) view returns(uint32)
func (_BN254CertificateVerifierStorage *BN254CertificateVerifierStorageCallerSession) MaxOperatorTableStaleness(operatorSet OperatorSet) (uint32, error) {
	return _BN254CertificateVerifierStorage.Contract.MaxOperatorTableStaleness(&_BN254CertificateVerifierStorage.CallOpts, operatorSet)
}

// OperatorTableUpdater is a free data retrieval call binding the contract method 0x68d6e081.
//
// Solidity: function operatorTableUpdater() view returns(address)
func (_BN254CertificateVerifierStorage *BN254CertificateVerifierStorageCaller) OperatorTableUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BN254CertificateVerifierStorage.contract.Call(opts, &out, "operatorTableUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OperatorTableUpdater is a free data retrieval call binding the contract method 0x68d6e081.
//
// Solidity: function operatorTableUpdater() view returns(address)
func (_BN254CertificateVerifierStorage *BN254CertificateVerifierStorageSession) OperatorTableUpdater() (common.Address, error) {
	return _BN254CertificateVerifierStorage.Contract.OperatorTableUpdater(&_BN254CertificateVerifierStorage.CallOpts)
}

// OperatorTableUpdater is a free data retrieval call binding the contract method 0x68d6e081.
//
// Solidity: function operatorTableUpdater() view returns(address)
func (_BN254CertificateVerifierStorage *BN254CertificateVerifierStorageCallerSession) OperatorTableUpdater() (common.Address, error) {
	return _BN254CertificateVerifierStorage.Contract.OperatorTableUpdater(&_BN254CertificateVerifierStorage.CallOpts)
}

// TrySignatureVerification is a free data retrieval call binding the contract method 0x1a18746c.
//
// Solidity: function trySignatureVerification(bytes32 msgHash, (uint256,uint256) aggPubkey, (uint256[2],uint256[2]) apkG2, (uint256,uint256) signature) view returns(bool pairingSuccessful, bool signatureValid)
func (_BN254CertificateVerifierStorage *BN254CertificateVerifierStorageCaller) TrySignatureVerification(opts *bind.CallOpts, msgHash [32]byte, aggPubkey BN254G1Point, apkG2 BN254G2Point, signature BN254G1Point) (struct {
	PairingSuccessful bool
	SignatureValid    bool
}, error) {
	var out []interface{}
	err := _BN254CertificateVerifierStorage.contract.Call(opts, &out, "trySignatureVerification", msgHash, aggPubkey, apkG2, signature)

	outstruct := new(struct {
		PairingSuccessful bool
		SignatureValid    bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PairingSuccessful = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.SignatureValid = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// TrySignatureVerification is a free data retrieval call binding the contract method 0x1a18746c.
//
// Solidity: function trySignatureVerification(bytes32 msgHash, (uint256,uint256) aggPubkey, (uint256[2],uint256[2]) apkG2, (uint256,uint256) signature) view returns(bool pairingSuccessful, bool signatureValid)
func (_BN254CertificateVerifierStorage *BN254CertificateVerifierStorageSession) TrySignatureVerification(msgHash [32]byte, aggPubkey BN254G1Point, apkG2 BN254G2Point, signature BN254G1Point) (struct {
	PairingSuccessful bool
	SignatureValid    bool
}, error) {
	return _BN254CertificateVerifierStorage.Contract.TrySignatureVerification(&_BN254CertificateVerifierStorage.CallOpts, msgHash, aggPubkey, apkG2, signature)
}

// TrySignatureVerification is a free data retrieval call binding the contract method 0x1a18746c.
//
// Solidity: function trySignatureVerification(bytes32 msgHash, (uint256,uint256) aggPubkey, (uint256[2],uint256[2]) apkG2, (uint256,uint256) signature) view returns(bool pairingSuccessful, bool signatureValid)
func (_BN254CertificateVerifierStorage *BN254CertificateVerifierStorageCallerSession) TrySignatureVerification(msgHash [32]byte, aggPubkey BN254G1Point, apkG2 BN254G2Point, signature BN254G1Point) (struct {
	PairingSuccessful bool
	SignatureValid    bool
}, error) {
	return _BN254CertificateVerifierStorage.Contract.TrySignatureVerification(&_BN254CertificateVerifierStorage.CallOpts, msgHash, aggPubkey, apkG2, signature)
}

// UpdateOperatorTable is a paid mutator transaction binding the contract method 0x6738c40b.
//
// Solidity: function updateOperatorTable((address,uint32) operatorSet, uint32 referenceTimestamp, (bytes32,uint256,(uint256,uint256),uint256[]) operatorSetInfo, (address,uint32) operatorSetConfig) returns()
func (_BN254CertificateVerifierStorage *BN254CertificateVerifierStorageTransactor) UpdateOperatorTable(opts *bind.TransactOpts, operatorSet OperatorSet, referenceTimestamp uint32, operatorSetInfo IBN254TableCalculatorTypesBN254OperatorSetInfo, operatorSetConfig ICrossChainRegistryTypesOperatorSetConfig) (*types.Transaction, error) {
	return _BN254CertificateVerifierStorage.contract.Transact(opts, "updateOperatorTable", operatorSet, referenceTimestamp, operatorSetInfo, operatorSetConfig)
}

// UpdateOperatorTable is a paid mutator transaction binding the contract method 0x6738c40b.
//
// Solidity: function updateOperatorTable((address,uint32) operatorSet, uint32 referenceTimestamp, (bytes32,uint256,(uint256,uint256),uint256[]) operatorSetInfo, (address,uint32) operatorSetConfig) returns()
func (_BN254CertificateVerifierStorage *BN254CertificateVerifierStorageSession) UpdateOperatorTable(operatorSet OperatorSet, referenceTimestamp uint32, operatorSetInfo IBN254TableCalculatorTypesBN254OperatorSetInfo, operatorSetConfig ICrossChainRegistryTypesOperatorSetConfig) (*types.Transaction, error) {
	return _BN254CertificateVerifierStorage.Contract.UpdateOperatorTable(&_BN254CertificateVerifierStorage.TransactOpts, operatorSet, referenceTimestamp, operatorSetInfo, operatorSetConfig)
}

// UpdateOperatorTable is a paid mutator transaction binding the contract method 0x6738c40b.
//
// Solidity: function updateOperatorTable((address,uint32) operatorSet, uint32 referenceTimestamp, (bytes32,uint256,(uint256,uint256),uint256[]) operatorSetInfo, (address,uint32) operatorSetConfig) returns()
func (_BN254CertificateVerifierStorage *BN254CertificateVerifierStorageTransactorSession) UpdateOperatorTable(operatorSet OperatorSet, referenceTimestamp uint32, operatorSetInfo IBN254TableCalculatorTypesBN254OperatorSetInfo, operatorSetConfig ICrossChainRegistryTypesOperatorSetConfig) (*types.Transaction, error) {
	return _BN254CertificateVerifierStorage.Contract.UpdateOperatorTable(&_BN254CertificateVerifierStorage.TransactOpts, operatorSet, referenceTimestamp, operatorSetInfo, operatorSetConfig)
}

// VerifyCertificate is a paid mutator transaction binding the contract method 0x080b7150.
//
// Solidity: function verifyCertificate((address,uint32) operatorSet, (uint32,bytes32,(uint256,uint256),(uint256[2],uint256[2]),(uint32,bytes,((uint256,uint256),uint256[]))[]) cert) returns(uint256[] signedStakes)
func (_BN254CertificateVerifierStorage *BN254CertificateVerifierStorageTransactor) VerifyCertificate(opts *bind.TransactOpts, operatorSet OperatorSet, cert IBN254CertificateVerifierTypesBN254Certificate) (*types.Transaction, error) {
	return _BN254CertificateVerifierStorage.contract.Transact(opts, "verifyCertificate", operatorSet, cert)
}

// VerifyCertificate is a paid mutator transaction binding the contract method 0x080b7150.
//
// Solidity: function verifyCertificate((address,uint32) operatorSet, (uint32,bytes32,(uint256,uint256),(uint256[2],uint256[2]),(uint32,bytes,((uint256,uint256),uint256[]))[]) cert) returns(uint256[] signedStakes)
func (_BN254CertificateVerifierStorage *BN254CertificateVerifierStorageSession) VerifyCertificate(operatorSet OperatorSet, cert IBN254CertificateVerifierTypesBN254Certificate) (*types.Transaction, error) {
	return _BN254CertificateVerifierStorage.Contract.VerifyCertificate(&_BN254CertificateVerifierStorage.TransactOpts, operatorSet, cert)
}

// VerifyCertificate is a paid mutator transaction binding the contract method 0x080b7150.
//
// Solidity: function verifyCertificate((address,uint32) operatorSet, (uint32,bytes32,(uint256,uint256),(uint256[2],uint256[2]),(uint32,bytes,((uint256,uint256),uint256[]))[]) cert) returns(uint256[] signedStakes)
func (_BN254CertificateVerifierStorage *BN254CertificateVerifierStorageTransactorSession) VerifyCertificate(operatorSet OperatorSet, cert IBN254CertificateVerifierTypesBN254Certificate) (*types.Transaction, error) {
	return _BN254CertificateVerifierStorage.Contract.VerifyCertificate(&_BN254CertificateVerifierStorage.TransactOpts, operatorSet, cert)
}

// VerifyCertificateNominal is a paid mutator transaction binding the contract method 0xdd2ae1b9.
//
// Solidity: function verifyCertificateNominal((address,uint32) operatorSet, (uint32,bytes32,(uint256,uint256),(uint256[2],uint256[2]),(uint32,bytes,((uint256,uint256),uint256[]))[]) cert, uint256[] totalStakeNominalThresholds) returns(bool)
func (_BN254CertificateVerifierStorage *BN254CertificateVerifierStorageTransactor) VerifyCertificateNominal(opts *bind.TransactOpts, operatorSet OperatorSet, cert IBN254CertificateVerifierTypesBN254Certificate, totalStakeNominalThresholds []*big.Int) (*types.Transaction, error) {
	return _BN254CertificateVerifierStorage.contract.Transact(opts, "verifyCertificateNominal", operatorSet, cert, totalStakeNominalThresholds)
}

// VerifyCertificateNominal is a paid mutator transaction binding the contract method 0xdd2ae1b9.
//
// Solidity: function verifyCertificateNominal((address,uint32) operatorSet, (uint32,bytes32,(uint256,uint256),(uint256[2],uint256[2]),(uint32,bytes,((uint256,uint256),uint256[]))[]) cert, uint256[] totalStakeNominalThresholds) returns(bool)
func (_BN254CertificateVerifierStorage *BN254CertificateVerifierStorageSession) VerifyCertificateNominal(operatorSet OperatorSet, cert IBN254CertificateVerifierTypesBN254Certificate, totalStakeNominalThresholds []*big.Int) (*types.Transaction, error) {
	return _BN254CertificateVerifierStorage.Contract.VerifyCertificateNominal(&_BN254CertificateVerifierStorage.TransactOpts, operatorSet, cert, totalStakeNominalThresholds)
}

// VerifyCertificateNominal is a paid mutator transaction binding the contract method 0xdd2ae1b9.
//
// Solidity: function verifyCertificateNominal((address,uint32) operatorSet, (uint32,bytes32,(uint256,uint256),(uint256[2],uint256[2]),(uint32,bytes,((uint256,uint256),uint256[]))[]) cert, uint256[] totalStakeNominalThresholds) returns(bool)
func (_BN254CertificateVerifierStorage *BN254CertificateVerifierStorageTransactorSession) VerifyCertificateNominal(operatorSet OperatorSet, cert IBN254CertificateVerifierTypesBN254Certificate, totalStakeNominalThresholds []*big.Int) (*types.Transaction, error) {
	return _BN254CertificateVerifierStorage.Contract.VerifyCertificateNominal(&_BN254CertificateVerifierStorage.TransactOpts, operatorSet, cert, totalStakeNominalThresholds)
}

// VerifyCertificateProportion is a paid mutator transaction binding the contract method 0x017d7974.
//
// Solidity: function verifyCertificateProportion((address,uint32) operatorSet, (uint32,bytes32,(uint256,uint256),(uint256[2],uint256[2]),(uint32,bytes,((uint256,uint256),uint256[]))[]) cert, uint16[] totalStakeProportionThresholds) returns(bool)
func (_BN254CertificateVerifierStorage *BN254CertificateVerifierStorageTransactor) VerifyCertificateProportion(opts *bind.TransactOpts, operatorSet OperatorSet, cert IBN254CertificateVerifierTypesBN254Certificate, totalStakeProportionThresholds []uint16) (*types.Transaction, error) {
	return _BN254CertificateVerifierStorage.contract.Transact(opts, "verifyCertificateProportion", operatorSet, cert, totalStakeProportionThresholds)
}

// VerifyCertificateProportion is a paid mutator transaction binding the contract method 0x017d7974.
//
// Solidity: function verifyCertificateProportion((address,uint32) operatorSet, (uint32,bytes32,(uint256,uint256),(uint256[2],uint256[2]),(uint32,bytes,((uint256,uint256),uint256[]))[]) cert, uint16[] totalStakeProportionThresholds) returns(bool)
func (_BN254CertificateVerifierStorage *BN254CertificateVerifierStorageSession) VerifyCertificateProportion(operatorSet OperatorSet, cert IBN254CertificateVerifierTypesBN254Certificate, totalStakeProportionThresholds []uint16) (*types.Transaction, error) {
	return _BN254CertificateVerifierStorage.Contract.VerifyCertificateProportion(&_BN254CertificateVerifierStorage.TransactOpts, operatorSet, cert, totalStakeProportionThresholds)
}

// VerifyCertificateProportion is a paid mutator transaction binding the contract method 0x017d7974.
//
// Solidity: function verifyCertificateProportion((address,uint32) operatorSet, (uint32,bytes32,(uint256,uint256),(uint256[2],uint256[2]),(uint32,bytes,((uint256,uint256),uint256[]))[]) cert, uint16[] totalStakeProportionThresholds) returns(bool)
func (_BN254CertificateVerifierStorage *BN254CertificateVerifierStorageTransactorSession) VerifyCertificateProportion(operatorSet OperatorSet, cert IBN254CertificateVerifierTypesBN254Certificate, totalStakeProportionThresholds []uint16) (*types.Transaction, error) {
	return _BN254CertificateVerifierStorage.Contract.VerifyCertificateProportion(&_BN254CertificateVerifierStorage.TransactOpts, operatorSet, cert, totalStakeProportionThresholds)
}

// BN254CertificateVerifierStorageMaxStalenessPeriodUpdatedIterator is returned from FilterMaxStalenessPeriodUpdated and is used to iterate over the raw logs and unpacked data for MaxStalenessPeriodUpdated events raised by the BN254CertificateVerifierStorage contract.
type BN254CertificateVerifierStorageMaxStalenessPeriodUpdatedIterator struct {
	Event *BN254CertificateVerifierStorageMaxStalenessPeriodUpdated // Event containing the contract specifics and raw log

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
func (it *BN254CertificateVerifierStorageMaxStalenessPeriodUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BN254CertificateVerifierStorageMaxStalenessPeriodUpdated)
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
		it.Event = new(BN254CertificateVerifierStorageMaxStalenessPeriodUpdated)
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
func (it *BN254CertificateVerifierStorageMaxStalenessPeriodUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BN254CertificateVerifierStorageMaxStalenessPeriodUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BN254CertificateVerifierStorageMaxStalenessPeriodUpdated represents a MaxStalenessPeriodUpdated event raised by the BN254CertificateVerifierStorage contract.
type BN254CertificateVerifierStorageMaxStalenessPeriodUpdated struct {
	OperatorSet        OperatorSet
	MaxStalenessPeriod uint32
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMaxStalenessPeriodUpdated is a free log retrieval operation binding the contract event 0x28539469fbbc8a5482e60966bf9376f7b9d25b2f0a65a9976f6baa3f0e3788da.
//
// Solidity: event MaxStalenessPeriodUpdated((address,uint32) operatorSet, uint32 maxStalenessPeriod)
func (_BN254CertificateVerifierStorage *BN254CertificateVerifierStorageFilterer) FilterMaxStalenessPeriodUpdated(opts *bind.FilterOpts) (*BN254CertificateVerifierStorageMaxStalenessPeriodUpdatedIterator, error) {

	logs, sub, err := _BN254CertificateVerifierStorage.contract.FilterLogs(opts, "MaxStalenessPeriodUpdated")
	if err != nil {
		return nil, err
	}
	return &BN254CertificateVerifierStorageMaxStalenessPeriodUpdatedIterator{contract: _BN254CertificateVerifierStorage.contract, event: "MaxStalenessPeriodUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxStalenessPeriodUpdated is a free log subscription operation binding the contract event 0x28539469fbbc8a5482e60966bf9376f7b9d25b2f0a65a9976f6baa3f0e3788da.
//
// Solidity: event MaxStalenessPeriodUpdated((address,uint32) operatorSet, uint32 maxStalenessPeriod)
func (_BN254CertificateVerifierStorage *BN254CertificateVerifierStorageFilterer) WatchMaxStalenessPeriodUpdated(opts *bind.WatchOpts, sink chan<- *BN254CertificateVerifierStorageMaxStalenessPeriodUpdated) (event.Subscription, error) {

	logs, sub, err := _BN254CertificateVerifierStorage.contract.WatchLogs(opts, "MaxStalenessPeriodUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BN254CertificateVerifierStorageMaxStalenessPeriodUpdated)
				if err := _BN254CertificateVerifierStorage.contract.UnpackLog(event, "MaxStalenessPeriodUpdated", log); err != nil {
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
func (_BN254CertificateVerifierStorage *BN254CertificateVerifierStorageFilterer) ParseMaxStalenessPeriodUpdated(log types.Log) (*BN254CertificateVerifierStorageMaxStalenessPeriodUpdated, error) {
	event := new(BN254CertificateVerifierStorageMaxStalenessPeriodUpdated)
	if err := _BN254CertificateVerifierStorage.contract.UnpackLog(event, "MaxStalenessPeriodUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BN254CertificateVerifierStorageOperatorSetOwnerUpdatedIterator is returned from FilterOperatorSetOwnerUpdated and is used to iterate over the raw logs and unpacked data for OperatorSetOwnerUpdated events raised by the BN254CertificateVerifierStorage contract.
type BN254CertificateVerifierStorageOperatorSetOwnerUpdatedIterator struct {
	Event *BN254CertificateVerifierStorageOperatorSetOwnerUpdated // Event containing the contract specifics and raw log

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
func (it *BN254CertificateVerifierStorageOperatorSetOwnerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BN254CertificateVerifierStorageOperatorSetOwnerUpdated)
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
		it.Event = new(BN254CertificateVerifierStorageOperatorSetOwnerUpdated)
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
func (it *BN254CertificateVerifierStorageOperatorSetOwnerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BN254CertificateVerifierStorageOperatorSetOwnerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BN254CertificateVerifierStorageOperatorSetOwnerUpdated represents a OperatorSetOwnerUpdated event raised by the BN254CertificateVerifierStorage contract.
type BN254CertificateVerifierStorageOperatorSetOwnerUpdated struct {
	OperatorSet OperatorSet
	Owner       common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorSetOwnerUpdated is a free log retrieval operation binding the contract event 0x806dc367095c0baf953d7144b7c4376261675ee0b4e0da2761e43673051c7375.
//
// Solidity: event OperatorSetOwnerUpdated((address,uint32) operatorSet, address owner)
func (_BN254CertificateVerifierStorage *BN254CertificateVerifierStorageFilterer) FilterOperatorSetOwnerUpdated(opts *bind.FilterOpts) (*BN254CertificateVerifierStorageOperatorSetOwnerUpdatedIterator, error) {

	logs, sub, err := _BN254CertificateVerifierStorage.contract.FilterLogs(opts, "OperatorSetOwnerUpdated")
	if err != nil {
		return nil, err
	}
	return &BN254CertificateVerifierStorageOperatorSetOwnerUpdatedIterator{contract: _BN254CertificateVerifierStorage.contract, event: "OperatorSetOwnerUpdated", logs: logs, sub: sub}, nil
}

// WatchOperatorSetOwnerUpdated is a free log subscription operation binding the contract event 0x806dc367095c0baf953d7144b7c4376261675ee0b4e0da2761e43673051c7375.
//
// Solidity: event OperatorSetOwnerUpdated((address,uint32) operatorSet, address owner)
func (_BN254CertificateVerifierStorage *BN254CertificateVerifierStorageFilterer) WatchOperatorSetOwnerUpdated(opts *bind.WatchOpts, sink chan<- *BN254CertificateVerifierStorageOperatorSetOwnerUpdated) (event.Subscription, error) {

	logs, sub, err := _BN254CertificateVerifierStorage.contract.WatchLogs(opts, "OperatorSetOwnerUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BN254CertificateVerifierStorageOperatorSetOwnerUpdated)
				if err := _BN254CertificateVerifierStorage.contract.UnpackLog(event, "OperatorSetOwnerUpdated", log); err != nil {
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
func (_BN254CertificateVerifierStorage *BN254CertificateVerifierStorageFilterer) ParseOperatorSetOwnerUpdated(log types.Log) (*BN254CertificateVerifierStorageOperatorSetOwnerUpdated, error) {
	event := new(BN254CertificateVerifierStorageOperatorSetOwnerUpdated)
	if err := _BN254CertificateVerifierStorage.contract.UnpackLog(event, "OperatorSetOwnerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BN254CertificateVerifierStorageTableUpdatedIterator is returned from FilterTableUpdated and is used to iterate over the raw logs and unpacked data for TableUpdated events raised by the BN254CertificateVerifierStorage contract.
type BN254CertificateVerifierStorageTableUpdatedIterator struct {
	Event *BN254CertificateVerifierStorageTableUpdated // Event containing the contract specifics and raw log

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
func (it *BN254CertificateVerifierStorageTableUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BN254CertificateVerifierStorageTableUpdated)
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
		it.Event = new(BN254CertificateVerifierStorageTableUpdated)
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
func (it *BN254CertificateVerifierStorageTableUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BN254CertificateVerifierStorageTableUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BN254CertificateVerifierStorageTableUpdated represents a TableUpdated event raised by the BN254CertificateVerifierStorage contract.
type BN254CertificateVerifierStorageTableUpdated struct {
	OperatorSet        OperatorSet
	ReferenceTimestamp uint32
	OperatorSetInfo    IBN254TableCalculatorTypesBN254OperatorSetInfo
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterTableUpdated is a free log retrieval operation binding the contract event 0x93e6bea1c9b5dce4a5c07b00261e956df2a4a253d9ab6ca070ca2037d72ada9e.
//
// Solidity: event TableUpdated((address,uint32) operatorSet, uint32 referenceTimestamp, (bytes32,uint256,(uint256,uint256),uint256[]) operatorSetInfo)
func (_BN254CertificateVerifierStorage *BN254CertificateVerifierStorageFilterer) FilterTableUpdated(opts *bind.FilterOpts) (*BN254CertificateVerifierStorageTableUpdatedIterator, error) {

	logs, sub, err := _BN254CertificateVerifierStorage.contract.FilterLogs(opts, "TableUpdated")
	if err != nil {
		return nil, err
	}
	return &BN254CertificateVerifierStorageTableUpdatedIterator{contract: _BN254CertificateVerifierStorage.contract, event: "TableUpdated", logs: logs, sub: sub}, nil
}

// WatchTableUpdated is a free log subscription operation binding the contract event 0x93e6bea1c9b5dce4a5c07b00261e956df2a4a253d9ab6ca070ca2037d72ada9e.
//
// Solidity: event TableUpdated((address,uint32) operatorSet, uint32 referenceTimestamp, (bytes32,uint256,(uint256,uint256),uint256[]) operatorSetInfo)
func (_BN254CertificateVerifierStorage *BN254CertificateVerifierStorageFilterer) WatchTableUpdated(opts *bind.WatchOpts, sink chan<- *BN254CertificateVerifierStorageTableUpdated) (event.Subscription, error) {

	logs, sub, err := _BN254CertificateVerifierStorage.contract.WatchLogs(opts, "TableUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BN254CertificateVerifierStorageTableUpdated)
				if err := _BN254CertificateVerifierStorage.contract.UnpackLog(event, "TableUpdated", log); err != nil {
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

// ParseTableUpdated is a log parse operation binding the contract event 0x93e6bea1c9b5dce4a5c07b00261e956df2a4a253d9ab6ca070ca2037d72ada9e.
//
// Solidity: event TableUpdated((address,uint32) operatorSet, uint32 referenceTimestamp, (bytes32,uint256,(uint256,uint256),uint256[]) operatorSetInfo)
func (_BN254CertificateVerifierStorage *BN254CertificateVerifierStorageFilterer) ParseTableUpdated(log types.Log) (*BN254CertificateVerifierStorageTableUpdated, error) {
	event := new(BN254CertificateVerifierStorageTableUpdated)
	if err := _BN254CertificateVerifierStorage.contract.UnpackLog(event, "TableUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
