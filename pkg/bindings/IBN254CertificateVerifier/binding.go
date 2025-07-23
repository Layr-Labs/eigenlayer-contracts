// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IBN254CertificateVerifier

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

// IBN254CertificateVerifierMetaData contains all meta data concerning the IBN254CertificateVerifier contract.
var IBN254CertificateVerifierMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"getNonsignerOperatorInfo\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIOperatorTableCalculatorTypes.BN254OperatorInfo\",\"components\":[{\"name\":\"pubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorCount\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorSetInfo\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIOperatorTableCalculatorTypes.BN254OperatorSetInfo\",\"components\":[{\"name\":\"operatorInfoTreeRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"numOperators\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"aggregatePubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"totalWeights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorSetOwner\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalStakeWeights\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isNonsignerCached\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isReferenceTimestampSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestReferenceTimestamp\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxOperatorTableStaleness\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"trySignatureVerification\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"aggPubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"signature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"pairingSuccessful\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"signatureValid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateOperatorTable\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorSetInfo\",\"type\":\"tuple\",\"internalType\":\"structIOperatorTableCalculatorTypes.BN254OperatorSetInfo\",\"components\":[{\"name\":\"operatorInfoTreeRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"numOperators\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"aggregatePubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"totalWeights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]},{\"name\":\"operatorSetConfig\",\"type\":\"tuple\",\"internalType\":\"structICrossChainRegistryTypes.OperatorSetConfig\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxStalenessPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyCertificate\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"cert\",\"type\":\"tuple\",\"internalType\":\"structIBN254CertificateVerifierTypes.BN254Certificate\",\"components\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"nonSignerWitnesses\",\"type\":\"tuple[]\",\"internalType\":\"structIBN254CertificateVerifierTypes.BN254OperatorInfoWitness[]\",\"components\":[{\"name\":\"operatorIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorInfoProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"operatorInfo\",\"type\":\"tuple\",\"internalType\":\"structIOperatorTableCalculatorTypes.BN254OperatorInfo\",\"components\":[{\"name\":\"pubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}]}]}],\"outputs\":[{\"name\":\"signedStakes\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyCertificateNominal\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"cert\",\"type\":\"tuple\",\"internalType\":\"structIBN254CertificateVerifierTypes.BN254Certificate\",\"components\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"nonSignerWitnesses\",\"type\":\"tuple[]\",\"internalType\":\"structIBN254CertificateVerifierTypes.BN254OperatorInfoWitness[]\",\"components\":[{\"name\":\"operatorIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorInfoProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"operatorInfo\",\"type\":\"tuple\",\"internalType\":\"structIOperatorTableCalculatorTypes.BN254OperatorInfo\",\"components\":[{\"name\":\"pubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}]}]},{\"name\":\"totalStakeNominalThresholds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyCertificateProportion\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"cert\",\"type\":\"tuple\",\"internalType\":\"structIBN254CertificateVerifierTypes.BN254Certificate\",\"components\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"nonSignerWitnesses\",\"type\":\"tuple[]\",\"internalType\":\"structIBN254CertificateVerifierTypes.BN254OperatorInfoWitness[]\",\"components\":[{\"name\":\"operatorIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorInfoProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"operatorInfo\",\"type\":\"tuple\",\"internalType\":\"structIOperatorTableCalculatorTypes.BN254OperatorInfo\",\"components\":[{\"name\":\"pubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}]}]},{\"name\":\"totalStakeProportionThresholds\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"MaxStalenessPeriodUpdated\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"maxStalenessPeriod\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSetOwnerUpdated\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TableUpdated\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"operatorSetInfo\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIOperatorTableCalculatorTypes.BN254OperatorSetInfo\",\"components\":[{\"name\":\"operatorInfoTreeRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"numOperators\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"aggregatePubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"totalWeights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ArrayLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CertificateStale\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperatorIndex\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyTableUpdater\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReferenceTimestampDoesNotExist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RootDisabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TableUpdateStale\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"VerificationFailed\",\"inputs\":[]}]",
}

// IBN254CertificateVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use IBN254CertificateVerifierMetaData.ABI instead.
var IBN254CertificateVerifierABI = IBN254CertificateVerifierMetaData.ABI

// IBN254CertificateVerifier is an auto generated Go binding around an Ethereum contract.
type IBN254CertificateVerifier struct {
	IBN254CertificateVerifierCaller     // Read-only binding to the contract
	IBN254CertificateVerifierTransactor // Write-only binding to the contract
	IBN254CertificateVerifierFilterer   // Log filterer for contract events
}

// IBN254CertificateVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBN254CertificateVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBN254CertificateVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBN254CertificateVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBN254CertificateVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBN254CertificateVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBN254CertificateVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBN254CertificateVerifierSession struct {
	Contract     *IBN254CertificateVerifier // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IBN254CertificateVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBN254CertificateVerifierCallerSession struct {
	Contract *IBN254CertificateVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// IBN254CertificateVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBN254CertificateVerifierTransactorSession struct {
	Contract     *IBN254CertificateVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// IBN254CertificateVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBN254CertificateVerifierRaw struct {
	Contract *IBN254CertificateVerifier // Generic contract binding to access the raw methods on
}

// IBN254CertificateVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBN254CertificateVerifierCallerRaw struct {
	Contract *IBN254CertificateVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// IBN254CertificateVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBN254CertificateVerifierTransactorRaw struct {
	Contract *IBN254CertificateVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBN254CertificateVerifier creates a new instance of IBN254CertificateVerifier, bound to a specific deployed contract.
func NewIBN254CertificateVerifier(address common.Address, backend bind.ContractBackend) (*IBN254CertificateVerifier, error) {
	contract, err := bindIBN254CertificateVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBN254CertificateVerifier{IBN254CertificateVerifierCaller: IBN254CertificateVerifierCaller{contract: contract}, IBN254CertificateVerifierTransactor: IBN254CertificateVerifierTransactor{contract: contract}, IBN254CertificateVerifierFilterer: IBN254CertificateVerifierFilterer{contract: contract}}, nil
}

// NewIBN254CertificateVerifierCaller creates a new read-only instance of IBN254CertificateVerifier, bound to a specific deployed contract.
func NewIBN254CertificateVerifierCaller(address common.Address, caller bind.ContractCaller) (*IBN254CertificateVerifierCaller, error) {
	contract, err := bindIBN254CertificateVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBN254CertificateVerifierCaller{contract: contract}, nil
}

// NewIBN254CertificateVerifierTransactor creates a new write-only instance of IBN254CertificateVerifier, bound to a specific deployed contract.
func NewIBN254CertificateVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*IBN254CertificateVerifierTransactor, error) {
	contract, err := bindIBN254CertificateVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBN254CertificateVerifierTransactor{contract: contract}, nil
}

// NewIBN254CertificateVerifierFilterer creates a new log filterer instance of IBN254CertificateVerifier, bound to a specific deployed contract.
func NewIBN254CertificateVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*IBN254CertificateVerifierFilterer, error) {
	contract, err := bindIBN254CertificateVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBN254CertificateVerifierFilterer{contract: contract}, nil
}

// bindIBN254CertificateVerifier binds a generic wrapper to an already deployed contract.
func bindIBN254CertificateVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IBN254CertificateVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBN254CertificateVerifier *IBN254CertificateVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBN254CertificateVerifier.Contract.IBN254CertificateVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBN254CertificateVerifier *IBN254CertificateVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBN254CertificateVerifier.Contract.IBN254CertificateVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBN254CertificateVerifier *IBN254CertificateVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBN254CertificateVerifier.Contract.IBN254CertificateVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBN254CertificateVerifier *IBN254CertificateVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBN254CertificateVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBN254CertificateVerifier *IBN254CertificateVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBN254CertificateVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBN254CertificateVerifier *IBN254CertificateVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBN254CertificateVerifier.Contract.contract.Transact(opts, method, params...)
}

// GetNonsignerOperatorInfo is a free data retrieval call binding the contract method 0x26af6a3c.
//
// Solidity: function getNonsignerOperatorInfo((address,uint32) operatorSet, uint32 referenceTimestamp, uint256 operatorIndex) view returns(((uint256,uint256),uint256[]))
func (_IBN254CertificateVerifier *IBN254CertificateVerifierCaller) GetNonsignerOperatorInfo(opts *bind.CallOpts, operatorSet OperatorSet, referenceTimestamp uint32, operatorIndex *big.Int) (IOperatorTableCalculatorTypesBN254OperatorInfo, error) {
	var out []interface{}
	err := _IBN254CertificateVerifier.contract.Call(opts, &out, "getNonsignerOperatorInfo", operatorSet, referenceTimestamp, operatorIndex)

	if err != nil {
		return *new(IOperatorTableCalculatorTypesBN254OperatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IOperatorTableCalculatorTypesBN254OperatorInfo)).(*IOperatorTableCalculatorTypesBN254OperatorInfo)

	return out0, err

}

// GetNonsignerOperatorInfo is a free data retrieval call binding the contract method 0x26af6a3c.
//
// Solidity: function getNonsignerOperatorInfo((address,uint32) operatorSet, uint32 referenceTimestamp, uint256 operatorIndex) view returns(((uint256,uint256),uint256[]))
func (_IBN254CertificateVerifier *IBN254CertificateVerifierSession) GetNonsignerOperatorInfo(operatorSet OperatorSet, referenceTimestamp uint32, operatorIndex *big.Int) (IOperatorTableCalculatorTypesBN254OperatorInfo, error) {
	return _IBN254CertificateVerifier.Contract.GetNonsignerOperatorInfo(&_IBN254CertificateVerifier.CallOpts, operatorSet, referenceTimestamp, operatorIndex)
}

// GetNonsignerOperatorInfo is a free data retrieval call binding the contract method 0x26af6a3c.
//
// Solidity: function getNonsignerOperatorInfo((address,uint32) operatorSet, uint32 referenceTimestamp, uint256 operatorIndex) view returns(((uint256,uint256),uint256[]))
func (_IBN254CertificateVerifier *IBN254CertificateVerifierCallerSession) GetNonsignerOperatorInfo(operatorSet OperatorSet, referenceTimestamp uint32, operatorIndex *big.Int) (IOperatorTableCalculatorTypesBN254OperatorInfo, error) {
	return _IBN254CertificateVerifier.Contract.GetNonsignerOperatorInfo(&_IBN254CertificateVerifier.CallOpts, operatorSet, referenceTimestamp, operatorIndex)
}

// GetOperatorCount is a free data retrieval call binding the contract method 0x23c2a3cb.
//
// Solidity: function getOperatorCount((address,uint32) operatorSet, uint32 referenceTimestamp) view returns(uint256)
func (_IBN254CertificateVerifier *IBN254CertificateVerifierCaller) GetOperatorCount(opts *bind.CallOpts, operatorSet OperatorSet, referenceTimestamp uint32) (*big.Int, error) {
	var out []interface{}
	err := _IBN254CertificateVerifier.contract.Call(opts, &out, "getOperatorCount", operatorSet, referenceTimestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOperatorCount is a free data retrieval call binding the contract method 0x23c2a3cb.
//
// Solidity: function getOperatorCount((address,uint32) operatorSet, uint32 referenceTimestamp) view returns(uint256)
func (_IBN254CertificateVerifier *IBN254CertificateVerifierSession) GetOperatorCount(operatorSet OperatorSet, referenceTimestamp uint32) (*big.Int, error) {
	return _IBN254CertificateVerifier.Contract.GetOperatorCount(&_IBN254CertificateVerifier.CallOpts, operatorSet, referenceTimestamp)
}

// GetOperatorCount is a free data retrieval call binding the contract method 0x23c2a3cb.
//
// Solidity: function getOperatorCount((address,uint32) operatorSet, uint32 referenceTimestamp) view returns(uint256)
func (_IBN254CertificateVerifier *IBN254CertificateVerifierCallerSession) GetOperatorCount(operatorSet OperatorSet, referenceTimestamp uint32) (*big.Int, error) {
	return _IBN254CertificateVerifier.Contract.GetOperatorCount(&_IBN254CertificateVerifier.CallOpts, operatorSet, referenceTimestamp)
}

// GetOperatorSetInfo is a free data retrieval call binding the contract method 0xeb39e68f.
//
// Solidity: function getOperatorSetInfo((address,uint32) operatorSet, uint32 referenceTimestamp) view returns((bytes32,uint256,(uint256,uint256),uint256[]))
func (_IBN254CertificateVerifier *IBN254CertificateVerifierCaller) GetOperatorSetInfo(opts *bind.CallOpts, operatorSet OperatorSet, referenceTimestamp uint32) (IOperatorTableCalculatorTypesBN254OperatorSetInfo, error) {
	var out []interface{}
	err := _IBN254CertificateVerifier.contract.Call(opts, &out, "getOperatorSetInfo", operatorSet, referenceTimestamp)

	if err != nil {
		return *new(IOperatorTableCalculatorTypesBN254OperatorSetInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IOperatorTableCalculatorTypesBN254OperatorSetInfo)).(*IOperatorTableCalculatorTypesBN254OperatorSetInfo)

	return out0, err

}

// GetOperatorSetInfo is a free data retrieval call binding the contract method 0xeb39e68f.
//
// Solidity: function getOperatorSetInfo((address,uint32) operatorSet, uint32 referenceTimestamp) view returns((bytes32,uint256,(uint256,uint256),uint256[]))
func (_IBN254CertificateVerifier *IBN254CertificateVerifierSession) GetOperatorSetInfo(operatorSet OperatorSet, referenceTimestamp uint32) (IOperatorTableCalculatorTypesBN254OperatorSetInfo, error) {
	return _IBN254CertificateVerifier.Contract.GetOperatorSetInfo(&_IBN254CertificateVerifier.CallOpts, operatorSet, referenceTimestamp)
}

// GetOperatorSetInfo is a free data retrieval call binding the contract method 0xeb39e68f.
//
// Solidity: function getOperatorSetInfo((address,uint32) operatorSet, uint32 referenceTimestamp) view returns((bytes32,uint256,(uint256,uint256),uint256[]))
func (_IBN254CertificateVerifier *IBN254CertificateVerifierCallerSession) GetOperatorSetInfo(operatorSet OperatorSet, referenceTimestamp uint32) (IOperatorTableCalculatorTypesBN254OperatorSetInfo, error) {
	return _IBN254CertificateVerifier.Contract.GetOperatorSetInfo(&_IBN254CertificateVerifier.CallOpts, operatorSet, referenceTimestamp)
}

// GetOperatorSetOwner is a free data retrieval call binding the contract method 0x84818920.
//
// Solidity: function getOperatorSetOwner((address,uint32) operatorSet) view returns(address)
func (_IBN254CertificateVerifier *IBN254CertificateVerifierCaller) GetOperatorSetOwner(opts *bind.CallOpts, operatorSet OperatorSet) (common.Address, error) {
	var out []interface{}
	err := _IBN254CertificateVerifier.contract.Call(opts, &out, "getOperatorSetOwner", operatorSet)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOperatorSetOwner is a free data retrieval call binding the contract method 0x84818920.
//
// Solidity: function getOperatorSetOwner((address,uint32) operatorSet) view returns(address)
func (_IBN254CertificateVerifier *IBN254CertificateVerifierSession) GetOperatorSetOwner(operatorSet OperatorSet) (common.Address, error) {
	return _IBN254CertificateVerifier.Contract.GetOperatorSetOwner(&_IBN254CertificateVerifier.CallOpts, operatorSet)
}

// GetOperatorSetOwner is a free data retrieval call binding the contract method 0x84818920.
//
// Solidity: function getOperatorSetOwner((address,uint32) operatorSet) view returns(address)
func (_IBN254CertificateVerifier *IBN254CertificateVerifierCallerSession) GetOperatorSetOwner(operatorSet OperatorSet) (common.Address, error) {
	return _IBN254CertificateVerifier.Contract.GetOperatorSetOwner(&_IBN254CertificateVerifier.CallOpts, operatorSet)
}

// GetTotalStakeWeights is a free data retrieval call binding the contract method 0x7d1d1f5b.
//
// Solidity: function getTotalStakeWeights((address,uint32) operatorSet, uint32 referenceTimestamp) view returns(uint256[])
func (_IBN254CertificateVerifier *IBN254CertificateVerifierCaller) GetTotalStakeWeights(opts *bind.CallOpts, operatorSet OperatorSet, referenceTimestamp uint32) ([]*big.Int, error) {
	var out []interface{}
	err := _IBN254CertificateVerifier.contract.Call(opts, &out, "getTotalStakeWeights", operatorSet, referenceTimestamp)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetTotalStakeWeights is a free data retrieval call binding the contract method 0x7d1d1f5b.
//
// Solidity: function getTotalStakeWeights((address,uint32) operatorSet, uint32 referenceTimestamp) view returns(uint256[])
func (_IBN254CertificateVerifier *IBN254CertificateVerifierSession) GetTotalStakeWeights(operatorSet OperatorSet, referenceTimestamp uint32) ([]*big.Int, error) {
	return _IBN254CertificateVerifier.Contract.GetTotalStakeWeights(&_IBN254CertificateVerifier.CallOpts, operatorSet, referenceTimestamp)
}

// GetTotalStakeWeights is a free data retrieval call binding the contract method 0x7d1d1f5b.
//
// Solidity: function getTotalStakeWeights((address,uint32) operatorSet, uint32 referenceTimestamp) view returns(uint256[])
func (_IBN254CertificateVerifier *IBN254CertificateVerifierCallerSession) GetTotalStakeWeights(operatorSet OperatorSet, referenceTimestamp uint32) ([]*big.Int, error) {
	return _IBN254CertificateVerifier.Contract.GetTotalStakeWeights(&_IBN254CertificateVerifier.CallOpts, operatorSet, referenceTimestamp)
}

// IsNonsignerCached is a free data retrieval call binding the contract method 0x5be87274.
//
// Solidity: function isNonsignerCached((address,uint32) operatorSet, uint32 referenceTimestamp, uint256 operatorIndex) view returns(bool)
func (_IBN254CertificateVerifier *IBN254CertificateVerifierCaller) IsNonsignerCached(opts *bind.CallOpts, operatorSet OperatorSet, referenceTimestamp uint32, operatorIndex *big.Int) (bool, error) {
	var out []interface{}
	err := _IBN254CertificateVerifier.contract.Call(opts, &out, "isNonsignerCached", operatorSet, referenceTimestamp, operatorIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsNonsignerCached is a free data retrieval call binding the contract method 0x5be87274.
//
// Solidity: function isNonsignerCached((address,uint32) operatorSet, uint32 referenceTimestamp, uint256 operatorIndex) view returns(bool)
func (_IBN254CertificateVerifier *IBN254CertificateVerifierSession) IsNonsignerCached(operatorSet OperatorSet, referenceTimestamp uint32, operatorIndex *big.Int) (bool, error) {
	return _IBN254CertificateVerifier.Contract.IsNonsignerCached(&_IBN254CertificateVerifier.CallOpts, operatorSet, referenceTimestamp, operatorIndex)
}

// IsNonsignerCached is a free data retrieval call binding the contract method 0x5be87274.
//
// Solidity: function isNonsignerCached((address,uint32) operatorSet, uint32 referenceTimestamp, uint256 operatorIndex) view returns(bool)
func (_IBN254CertificateVerifier *IBN254CertificateVerifierCallerSession) IsNonsignerCached(operatorSet OperatorSet, referenceTimestamp uint32, operatorIndex *big.Int) (bool, error) {
	return _IBN254CertificateVerifier.Contract.IsNonsignerCached(&_IBN254CertificateVerifier.CallOpts, operatorSet, referenceTimestamp, operatorIndex)
}

// IsReferenceTimestampSet is a free data retrieval call binding the contract method 0xcd83a72b.
//
// Solidity: function isReferenceTimestampSet((address,uint32) operatorSet, uint32 referenceTimestamp) view returns(bool)
func (_IBN254CertificateVerifier *IBN254CertificateVerifierCaller) IsReferenceTimestampSet(opts *bind.CallOpts, operatorSet OperatorSet, referenceTimestamp uint32) (bool, error) {
	var out []interface{}
	err := _IBN254CertificateVerifier.contract.Call(opts, &out, "isReferenceTimestampSet", operatorSet, referenceTimestamp)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsReferenceTimestampSet is a free data retrieval call binding the contract method 0xcd83a72b.
//
// Solidity: function isReferenceTimestampSet((address,uint32) operatorSet, uint32 referenceTimestamp) view returns(bool)
func (_IBN254CertificateVerifier *IBN254CertificateVerifierSession) IsReferenceTimestampSet(operatorSet OperatorSet, referenceTimestamp uint32) (bool, error) {
	return _IBN254CertificateVerifier.Contract.IsReferenceTimestampSet(&_IBN254CertificateVerifier.CallOpts, operatorSet, referenceTimestamp)
}

// IsReferenceTimestampSet is a free data retrieval call binding the contract method 0xcd83a72b.
//
// Solidity: function isReferenceTimestampSet((address,uint32) operatorSet, uint32 referenceTimestamp) view returns(bool)
func (_IBN254CertificateVerifier *IBN254CertificateVerifierCallerSession) IsReferenceTimestampSet(operatorSet OperatorSet, referenceTimestamp uint32) (bool, error) {
	return _IBN254CertificateVerifier.Contract.IsReferenceTimestampSet(&_IBN254CertificateVerifier.CallOpts, operatorSet, referenceTimestamp)
}

// LatestReferenceTimestamp is a free data retrieval call binding the contract method 0x5ddb9b5b.
//
// Solidity: function latestReferenceTimestamp((address,uint32) operatorSet) view returns(uint32)
func (_IBN254CertificateVerifier *IBN254CertificateVerifierCaller) LatestReferenceTimestamp(opts *bind.CallOpts, operatorSet OperatorSet) (uint32, error) {
	var out []interface{}
	err := _IBN254CertificateVerifier.contract.Call(opts, &out, "latestReferenceTimestamp", operatorSet)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LatestReferenceTimestamp is a free data retrieval call binding the contract method 0x5ddb9b5b.
//
// Solidity: function latestReferenceTimestamp((address,uint32) operatorSet) view returns(uint32)
func (_IBN254CertificateVerifier *IBN254CertificateVerifierSession) LatestReferenceTimestamp(operatorSet OperatorSet) (uint32, error) {
	return _IBN254CertificateVerifier.Contract.LatestReferenceTimestamp(&_IBN254CertificateVerifier.CallOpts, operatorSet)
}

// LatestReferenceTimestamp is a free data retrieval call binding the contract method 0x5ddb9b5b.
//
// Solidity: function latestReferenceTimestamp((address,uint32) operatorSet) view returns(uint32)
func (_IBN254CertificateVerifier *IBN254CertificateVerifierCallerSession) LatestReferenceTimestamp(operatorSet OperatorSet) (uint32, error) {
	return _IBN254CertificateVerifier.Contract.LatestReferenceTimestamp(&_IBN254CertificateVerifier.CallOpts, operatorSet)
}

// MaxOperatorTableStaleness is a free data retrieval call binding the contract method 0x6141879e.
//
// Solidity: function maxOperatorTableStaleness((address,uint32) operatorSet) view returns(uint32)
func (_IBN254CertificateVerifier *IBN254CertificateVerifierCaller) MaxOperatorTableStaleness(opts *bind.CallOpts, operatorSet OperatorSet) (uint32, error) {
	var out []interface{}
	err := _IBN254CertificateVerifier.contract.Call(opts, &out, "maxOperatorTableStaleness", operatorSet)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MaxOperatorTableStaleness is a free data retrieval call binding the contract method 0x6141879e.
//
// Solidity: function maxOperatorTableStaleness((address,uint32) operatorSet) view returns(uint32)
func (_IBN254CertificateVerifier *IBN254CertificateVerifierSession) MaxOperatorTableStaleness(operatorSet OperatorSet) (uint32, error) {
	return _IBN254CertificateVerifier.Contract.MaxOperatorTableStaleness(&_IBN254CertificateVerifier.CallOpts, operatorSet)
}

// MaxOperatorTableStaleness is a free data retrieval call binding the contract method 0x6141879e.
//
// Solidity: function maxOperatorTableStaleness((address,uint32) operatorSet) view returns(uint32)
func (_IBN254CertificateVerifier *IBN254CertificateVerifierCallerSession) MaxOperatorTableStaleness(operatorSet OperatorSet) (uint32, error) {
	return _IBN254CertificateVerifier.Contract.MaxOperatorTableStaleness(&_IBN254CertificateVerifier.CallOpts, operatorSet)
}

// TrySignatureVerification is a free data retrieval call binding the contract method 0x1a18746c.
//
// Solidity: function trySignatureVerification(bytes32 msgHash, (uint256,uint256) aggPubkey, (uint256[2],uint256[2]) apkG2, (uint256,uint256) signature) view returns(bool pairingSuccessful, bool signatureValid)
func (_IBN254CertificateVerifier *IBN254CertificateVerifierCaller) TrySignatureVerification(opts *bind.CallOpts, msgHash [32]byte, aggPubkey BN254G1Point, apkG2 BN254G2Point, signature BN254G1Point) (struct {
	PairingSuccessful bool
	SignatureValid    bool
}, error) {
	var out []interface{}
	err := _IBN254CertificateVerifier.contract.Call(opts, &out, "trySignatureVerification", msgHash, aggPubkey, apkG2, signature)

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
func (_IBN254CertificateVerifier *IBN254CertificateVerifierSession) TrySignatureVerification(msgHash [32]byte, aggPubkey BN254G1Point, apkG2 BN254G2Point, signature BN254G1Point) (struct {
	PairingSuccessful bool
	SignatureValid    bool
}, error) {
	return _IBN254CertificateVerifier.Contract.TrySignatureVerification(&_IBN254CertificateVerifier.CallOpts, msgHash, aggPubkey, apkG2, signature)
}

// TrySignatureVerification is a free data retrieval call binding the contract method 0x1a18746c.
//
// Solidity: function trySignatureVerification(bytes32 msgHash, (uint256,uint256) aggPubkey, (uint256[2],uint256[2]) apkG2, (uint256,uint256) signature) view returns(bool pairingSuccessful, bool signatureValid)
func (_IBN254CertificateVerifier *IBN254CertificateVerifierCallerSession) TrySignatureVerification(msgHash [32]byte, aggPubkey BN254G1Point, apkG2 BN254G2Point, signature BN254G1Point) (struct {
	PairingSuccessful bool
	SignatureValid    bool
}, error) {
	return _IBN254CertificateVerifier.Contract.TrySignatureVerification(&_IBN254CertificateVerifier.CallOpts, msgHash, aggPubkey, apkG2, signature)
}

// UpdateOperatorTable is a paid mutator transaction binding the contract method 0x6738c40b.
//
// Solidity: function updateOperatorTable((address,uint32) operatorSet, uint32 referenceTimestamp, (bytes32,uint256,(uint256,uint256),uint256[]) operatorSetInfo, (address,uint32) operatorSetConfig) returns()
func (_IBN254CertificateVerifier *IBN254CertificateVerifierTransactor) UpdateOperatorTable(opts *bind.TransactOpts, operatorSet OperatorSet, referenceTimestamp uint32, operatorSetInfo IOperatorTableCalculatorTypesBN254OperatorSetInfo, operatorSetConfig ICrossChainRegistryTypesOperatorSetConfig) (*types.Transaction, error) {
	return _IBN254CertificateVerifier.contract.Transact(opts, "updateOperatorTable", operatorSet, referenceTimestamp, operatorSetInfo, operatorSetConfig)
}

// UpdateOperatorTable is a paid mutator transaction binding the contract method 0x6738c40b.
//
// Solidity: function updateOperatorTable((address,uint32) operatorSet, uint32 referenceTimestamp, (bytes32,uint256,(uint256,uint256),uint256[]) operatorSetInfo, (address,uint32) operatorSetConfig) returns()
func (_IBN254CertificateVerifier *IBN254CertificateVerifierSession) UpdateOperatorTable(operatorSet OperatorSet, referenceTimestamp uint32, operatorSetInfo IOperatorTableCalculatorTypesBN254OperatorSetInfo, operatorSetConfig ICrossChainRegistryTypesOperatorSetConfig) (*types.Transaction, error) {
	return _IBN254CertificateVerifier.Contract.UpdateOperatorTable(&_IBN254CertificateVerifier.TransactOpts, operatorSet, referenceTimestamp, operatorSetInfo, operatorSetConfig)
}

// UpdateOperatorTable is a paid mutator transaction binding the contract method 0x6738c40b.
//
// Solidity: function updateOperatorTable((address,uint32) operatorSet, uint32 referenceTimestamp, (bytes32,uint256,(uint256,uint256),uint256[]) operatorSetInfo, (address,uint32) operatorSetConfig) returns()
func (_IBN254CertificateVerifier *IBN254CertificateVerifierTransactorSession) UpdateOperatorTable(operatorSet OperatorSet, referenceTimestamp uint32, operatorSetInfo IOperatorTableCalculatorTypesBN254OperatorSetInfo, operatorSetConfig ICrossChainRegistryTypesOperatorSetConfig) (*types.Transaction, error) {
	return _IBN254CertificateVerifier.Contract.UpdateOperatorTable(&_IBN254CertificateVerifier.TransactOpts, operatorSet, referenceTimestamp, operatorSetInfo, operatorSetConfig)
}

// VerifyCertificate is a paid mutator transaction binding the contract method 0x080b7150.
//
// Solidity: function verifyCertificate((address,uint32) operatorSet, (uint32,bytes32,(uint256,uint256),(uint256[2],uint256[2]),(uint32,bytes,((uint256,uint256),uint256[]))[]) cert) returns(uint256[] signedStakes)
func (_IBN254CertificateVerifier *IBN254CertificateVerifierTransactor) VerifyCertificate(opts *bind.TransactOpts, operatorSet OperatorSet, cert IBN254CertificateVerifierTypesBN254Certificate) (*types.Transaction, error) {
	return _IBN254CertificateVerifier.contract.Transact(opts, "verifyCertificate", operatorSet, cert)
}

// VerifyCertificate is a paid mutator transaction binding the contract method 0x080b7150.
//
// Solidity: function verifyCertificate((address,uint32) operatorSet, (uint32,bytes32,(uint256,uint256),(uint256[2],uint256[2]),(uint32,bytes,((uint256,uint256),uint256[]))[]) cert) returns(uint256[] signedStakes)
func (_IBN254CertificateVerifier *IBN254CertificateVerifierSession) VerifyCertificate(operatorSet OperatorSet, cert IBN254CertificateVerifierTypesBN254Certificate) (*types.Transaction, error) {
	return _IBN254CertificateVerifier.Contract.VerifyCertificate(&_IBN254CertificateVerifier.TransactOpts, operatorSet, cert)
}

// VerifyCertificate is a paid mutator transaction binding the contract method 0x080b7150.
//
// Solidity: function verifyCertificate((address,uint32) operatorSet, (uint32,bytes32,(uint256,uint256),(uint256[2],uint256[2]),(uint32,bytes,((uint256,uint256),uint256[]))[]) cert) returns(uint256[] signedStakes)
func (_IBN254CertificateVerifier *IBN254CertificateVerifierTransactorSession) VerifyCertificate(operatorSet OperatorSet, cert IBN254CertificateVerifierTypesBN254Certificate) (*types.Transaction, error) {
	return _IBN254CertificateVerifier.Contract.VerifyCertificate(&_IBN254CertificateVerifier.TransactOpts, operatorSet, cert)
}

// VerifyCertificateNominal is a paid mutator transaction binding the contract method 0xdd2ae1b9.
//
// Solidity: function verifyCertificateNominal((address,uint32) operatorSet, (uint32,bytes32,(uint256,uint256),(uint256[2],uint256[2]),(uint32,bytes,((uint256,uint256),uint256[]))[]) cert, uint256[] totalStakeNominalThresholds) returns(bool)
func (_IBN254CertificateVerifier *IBN254CertificateVerifierTransactor) VerifyCertificateNominal(opts *bind.TransactOpts, operatorSet OperatorSet, cert IBN254CertificateVerifierTypesBN254Certificate, totalStakeNominalThresholds []*big.Int) (*types.Transaction, error) {
	return _IBN254CertificateVerifier.contract.Transact(opts, "verifyCertificateNominal", operatorSet, cert, totalStakeNominalThresholds)
}

// VerifyCertificateNominal is a paid mutator transaction binding the contract method 0xdd2ae1b9.
//
// Solidity: function verifyCertificateNominal((address,uint32) operatorSet, (uint32,bytes32,(uint256,uint256),(uint256[2],uint256[2]),(uint32,bytes,((uint256,uint256),uint256[]))[]) cert, uint256[] totalStakeNominalThresholds) returns(bool)
func (_IBN254CertificateVerifier *IBN254CertificateVerifierSession) VerifyCertificateNominal(operatorSet OperatorSet, cert IBN254CertificateVerifierTypesBN254Certificate, totalStakeNominalThresholds []*big.Int) (*types.Transaction, error) {
	return _IBN254CertificateVerifier.Contract.VerifyCertificateNominal(&_IBN254CertificateVerifier.TransactOpts, operatorSet, cert, totalStakeNominalThresholds)
}

// VerifyCertificateNominal is a paid mutator transaction binding the contract method 0xdd2ae1b9.
//
// Solidity: function verifyCertificateNominal((address,uint32) operatorSet, (uint32,bytes32,(uint256,uint256),(uint256[2],uint256[2]),(uint32,bytes,((uint256,uint256),uint256[]))[]) cert, uint256[] totalStakeNominalThresholds) returns(bool)
func (_IBN254CertificateVerifier *IBN254CertificateVerifierTransactorSession) VerifyCertificateNominal(operatorSet OperatorSet, cert IBN254CertificateVerifierTypesBN254Certificate, totalStakeNominalThresholds []*big.Int) (*types.Transaction, error) {
	return _IBN254CertificateVerifier.Contract.VerifyCertificateNominal(&_IBN254CertificateVerifier.TransactOpts, operatorSet, cert, totalStakeNominalThresholds)
}

// VerifyCertificateProportion is a paid mutator transaction binding the contract method 0x017d7974.
//
// Solidity: function verifyCertificateProportion((address,uint32) operatorSet, (uint32,bytes32,(uint256,uint256),(uint256[2],uint256[2]),(uint32,bytes,((uint256,uint256),uint256[]))[]) cert, uint16[] totalStakeProportionThresholds) returns(bool)
func (_IBN254CertificateVerifier *IBN254CertificateVerifierTransactor) VerifyCertificateProportion(opts *bind.TransactOpts, operatorSet OperatorSet, cert IBN254CertificateVerifierTypesBN254Certificate, totalStakeProportionThresholds []uint16) (*types.Transaction, error) {
	return _IBN254CertificateVerifier.contract.Transact(opts, "verifyCertificateProportion", operatorSet, cert, totalStakeProportionThresholds)
}

// VerifyCertificateProportion is a paid mutator transaction binding the contract method 0x017d7974.
//
// Solidity: function verifyCertificateProportion((address,uint32) operatorSet, (uint32,bytes32,(uint256,uint256),(uint256[2],uint256[2]),(uint32,bytes,((uint256,uint256),uint256[]))[]) cert, uint16[] totalStakeProportionThresholds) returns(bool)
func (_IBN254CertificateVerifier *IBN254CertificateVerifierSession) VerifyCertificateProportion(operatorSet OperatorSet, cert IBN254CertificateVerifierTypesBN254Certificate, totalStakeProportionThresholds []uint16) (*types.Transaction, error) {
	return _IBN254CertificateVerifier.Contract.VerifyCertificateProportion(&_IBN254CertificateVerifier.TransactOpts, operatorSet, cert, totalStakeProportionThresholds)
}

// VerifyCertificateProportion is a paid mutator transaction binding the contract method 0x017d7974.
//
// Solidity: function verifyCertificateProportion((address,uint32) operatorSet, (uint32,bytes32,(uint256,uint256),(uint256[2],uint256[2]),(uint32,bytes,((uint256,uint256),uint256[]))[]) cert, uint16[] totalStakeProportionThresholds) returns(bool)
func (_IBN254CertificateVerifier *IBN254CertificateVerifierTransactorSession) VerifyCertificateProportion(operatorSet OperatorSet, cert IBN254CertificateVerifierTypesBN254Certificate, totalStakeProportionThresholds []uint16) (*types.Transaction, error) {
	return _IBN254CertificateVerifier.Contract.VerifyCertificateProportion(&_IBN254CertificateVerifier.TransactOpts, operatorSet, cert, totalStakeProportionThresholds)
}

// IBN254CertificateVerifierMaxStalenessPeriodUpdatedIterator is returned from FilterMaxStalenessPeriodUpdated and is used to iterate over the raw logs and unpacked data for MaxStalenessPeriodUpdated events raised by the IBN254CertificateVerifier contract.
type IBN254CertificateVerifierMaxStalenessPeriodUpdatedIterator struct {
	Event *IBN254CertificateVerifierMaxStalenessPeriodUpdated // Event containing the contract specifics and raw log

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
func (it *IBN254CertificateVerifierMaxStalenessPeriodUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBN254CertificateVerifierMaxStalenessPeriodUpdated)
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
		it.Event = new(IBN254CertificateVerifierMaxStalenessPeriodUpdated)
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
func (it *IBN254CertificateVerifierMaxStalenessPeriodUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBN254CertificateVerifierMaxStalenessPeriodUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBN254CertificateVerifierMaxStalenessPeriodUpdated represents a MaxStalenessPeriodUpdated event raised by the IBN254CertificateVerifier contract.
type IBN254CertificateVerifierMaxStalenessPeriodUpdated struct {
	OperatorSet        OperatorSet
	MaxStalenessPeriod uint32
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMaxStalenessPeriodUpdated is a free log retrieval operation binding the contract event 0x28539469fbbc8a5482e60966bf9376f7b9d25b2f0a65a9976f6baa3f0e3788da.
//
// Solidity: event MaxStalenessPeriodUpdated((address,uint32) operatorSet, uint32 maxStalenessPeriod)
func (_IBN254CertificateVerifier *IBN254CertificateVerifierFilterer) FilterMaxStalenessPeriodUpdated(opts *bind.FilterOpts) (*IBN254CertificateVerifierMaxStalenessPeriodUpdatedIterator, error) {

	logs, sub, err := _IBN254CertificateVerifier.contract.FilterLogs(opts, "MaxStalenessPeriodUpdated")
	if err != nil {
		return nil, err
	}
	return &IBN254CertificateVerifierMaxStalenessPeriodUpdatedIterator{contract: _IBN254CertificateVerifier.contract, event: "MaxStalenessPeriodUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxStalenessPeriodUpdated is a free log subscription operation binding the contract event 0x28539469fbbc8a5482e60966bf9376f7b9d25b2f0a65a9976f6baa3f0e3788da.
//
// Solidity: event MaxStalenessPeriodUpdated((address,uint32) operatorSet, uint32 maxStalenessPeriod)
func (_IBN254CertificateVerifier *IBN254CertificateVerifierFilterer) WatchMaxStalenessPeriodUpdated(opts *bind.WatchOpts, sink chan<- *IBN254CertificateVerifierMaxStalenessPeriodUpdated) (event.Subscription, error) {

	logs, sub, err := _IBN254CertificateVerifier.contract.WatchLogs(opts, "MaxStalenessPeriodUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBN254CertificateVerifierMaxStalenessPeriodUpdated)
				if err := _IBN254CertificateVerifier.contract.UnpackLog(event, "MaxStalenessPeriodUpdated", log); err != nil {
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
func (_IBN254CertificateVerifier *IBN254CertificateVerifierFilterer) ParseMaxStalenessPeriodUpdated(log types.Log) (*IBN254CertificateVerifierMaxStalenessPeriodUpdated, error) {
	event := new(IBN254CertificateVerifierMaxStalenessPeriodUpdated)
	if err := _IBN254CertificateVerifier.contract.UnpackLog(event, "MaxStalenessPeriodUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBN254CertificateVerifierOperatorSetOwnerUpdatedIterator is returned from FilterOperatorSetOwnerUpdated and is used to iterate over the raw logs and unpacked data for OperatorSetOwnerUpdated events raised by the IBN254CertificateVerifier contract.
type IBN254CertificateVerifierOperatorSetOwnerUpdatedIterator struct {
	Event *IBN254CertificateVerifierOperatorSetOwnerUpdated // Event containing the contract specifics and raw log

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
func (it *IBN254CertificateVerifierOperatorSetOwnerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBN254CertificateVerifierOperatorSetOwnerUpdated)
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
		it.Event = new(IBN254CertificateVerifierOperatorSetOwnerUpdated)
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
func (it *IBN254CertificateVerifierOperatorSetOwnerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBN254CertificateVerifierOperatorSetOwnerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBN254CertificateVerifierOperatorSetOwnerUpdated represents a OperatorSetOwnerUpdated event raised by the IBN254CertificateVerifier contract.
type IBN254CertificateVerifierOperatorSetOwnerUpdated struct {
	OperatorSet OperatorSet
	Owner       common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorSetOwnerUpdated is a free log retrieval operation binding the contract event 0x806dc367095c0baf953d7144b7c4376261675ee0b4e0da2761e43673051c7375.
//
// Solidity: event OperatorSetOwnerUpdated((address,uint32) operatorSet, address owner)
func (_IBN254CertificateVerifier *IBN254CertificateVerifierFilterer) FilterOperatorSetOwnerUpdated(opts *bind.FilterOpts) (*IBN254CertificateVerifierOperatorSetOwnerUpdatedIterator, error) {

	logs, sub, err := _IBN254CertificateVerifier.contract.FilterLogs(opts, "OperatorSetOwnerUpdated")
	if err != nil {
		return nil, err
	}
	return &IBN254CertificateVerifierOperatorSetOwnerUpdatedIterator{contract: _IBN254CertificateVerifier.contract, event: "OperatorSetOwnerUpdated", logs: logs, sub: sub}, nil
}

// WatchOperatorSetOwnerUpdated is a free log subscription operation binding the contract event 0x806dc367095c0baf953d7144b7c4376261675ee0b4e0da2761e43673051c7375.
//
// Solidity: event OperatorSetOwnerUpdated((address,uint32) operatorSet, address owner)
func (_IBN254CertificateVerifier *IBN254CertificateVerifierFilterer) WatchOperatorSetOwnerUpdated(opts *bind.WatchOpts, sink chan<- *IBN254CertificateVerifierOperatorSetOwnerUpdated) (event.Subscription, error) {

	logs, sub, err := _IBN254CertificateVerifier.contract.WatchLogs(opts, "OperatorSetOwnerUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBN254CertificateVerifierOperatorSetOwnerUpdated)
				if err := _IBN254CertificateVerifier.contract.UnpackLog(event, "OperatorSetOwnerUpdated", log); err != nil {
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
func (_IBN254CertificateVerifier *IBN254CertificateVerifierFilterer) ParseOperatorSetOwnerUpdated(log types.Log) (*IBN254CertificateVerifierOperatorSetOwnerUpdated, error) {
	event := new(IBN254CertificateVerifierOperatorSetOwnerUpdated)
	if err := _IBN254CertificateVerifier.contract.UnpackLog(event, "OperatorSetOwnerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBN254CertificateVerifierTableUpdatedIterator is returned from FilterTableUpdated and is used to iterate over the raw logs and unpacked data for TableUpdated events raised by the IBN254CertificateVerifier contract.
type IBN254CertificateVerifierTableUpdatedIterator struct {
	Event *IBN254CertificateVerifierTableUpdated // Event containing the contract specifics and raw log

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
func (it *IBN254CertificateVerifierTableUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBN254CertificateVerifierTableUpdated)
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
		it.Event = new(IBN254CertificateVerifierTableUpdated)
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
func (it *IBN254CertificateVerifierTableUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBN254CertificateVerifierTableUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBN254CertificateVerifierTableUpdated represents a TableUpdated event raised by the IBN254CertificateVerifier contract.
type IBN254CertificateVerifierTableUpdated struct {
	OperatorSet        OperatorSet
	ReferenceTimestamp uint32
	OperatorSetInfo    IOperatorTableCalculatorTypesBN254OperatorSetInfo
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterTableUpdated is a free log retrieval operation binding the contract event 0x93e6bea1c9b5dce4a5c07b00261e956df2a4a253d9ab6ca070ca2037d72ada9e.
//
// Solidity: event TableUpdated((address,uint32) operatorSet, uint32 referenceTimestamp, (bytes32,uint256,(uint256,uint256),uint256[]) operatorSetInfo)
func (_IBN254CertificateVerifier *IBN254CertificateVerifierFilterer) FilterTableUpdated(opts *bind.FilterOpts) (*IBN254CertificateVerifierTableUpdatedIterator, error) {

	logs, sub, err := _IBN254CertificateVerifier.contract.FilterLogs(opts, "TableUpdated")
	if err != nil {
		return nil, err
	}
	return &IBN254CertificateVerifierTableUpdatedIterator{contract: _IBN254CertificateVerifier.contract, event: "TableUpdated", logs: logs, sub: sub}, nil
}

// WatchTableUpdated is a free log subscription operation binding the contract event 0x93e6bea1c9b5dce4a5c07b00261e956df2a4a253d9ab6ca070ca2037d72ada9e.
//
// Solidity: event TableUpdated((address,uint32) operatorSet, uint32 referenceTimestamp, (bytes32,uint256,(uint256,uint256),uint256[]) operatorSetInfo)
func (_IBN254CertificateVerifier *IBN254CertificateVerifierFilterer) WatchTableUpdated(opts *bind.WatchOpts, sink chan<- *IBN254CertificateVerifierTableUpdated) (event.Subscription, error) {

	logs, sub, err := _IBN254CertificateVerifier.contract.WatchLogs(opts, "TableUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBN254CertificateVerifierTableUpdated)
				if err := _IBN254CertificateVerifier.contract.UnpackLog(event, "TableUpdated", log); err != nil {
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
func (_IBN254CertificateVerifier *IBN254CertificateVerifierFilterer) ParseTableUpdated(log types.Log) (*IBN254CertificateVerifierTableUpdated, error) {
	event := new(IBN254CertificateVerifierTableUpdated)
	if err := _IBN254CertificateVerifier.contract.UnpackLog(event, "TableUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
