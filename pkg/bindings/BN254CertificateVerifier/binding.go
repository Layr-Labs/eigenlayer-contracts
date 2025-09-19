// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package BN254CertificateVerifier

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

// BN254CertificateVerifierMetaData contains all meta data concerning the BN254CertificateVerifier contract.
var BN254CertificateVerifierMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_operatorTableUpdater\",\"type\":\"address\",\"internalType\":\"contractIOperatorTableUpdater\"},{\"name\":\"_version\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"OPERATOR_INFO_LEAF_SALT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"OPERATOR_TABLE_LEAF_SALT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateCertificateDigest\",\"inputs\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"calculateOperatorInfoLeaf\",\"inputs\":[{\"name\":\"operatorInfo\",\"type\":\"tuple\",\"internalType\":\"structIOperatorTableCalculatorTypes.BN254OperatorInfo\",\"components\":[{\"name\":\"pubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"calculateOperatorTableLeaf\",\"inputs\":[{\"name\":\"operatorTableBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getNonsignerOperatorInfo\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIOperatorTableCalculatorTypes.BN254OperatorInfo\",\"components\":[{\"name\":\"pubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorCount\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorSetInfo\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIOperatorTableCalculatorTypes.BN254OperatorSetInfo\",\"components\":[{\"name\":\"operatorInfoTreeRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"numOperators\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"aggregatePubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"totalWeights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorSetOwner\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalStakeWeights\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isNonsignerCached\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isReferenceTimestampSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestReferenceTimestamp\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxOperatorTableStaleness\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorTableUpdater\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIOperatorTableUpdater\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"trySignatureVerification\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"aggPubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"signature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"pairingSuccessful\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"signatureValid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateOperatorTable\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorSetInfo\",\"type\":\"tuple\",\"internalType\":\"structIOperatorTableCalculatorTypes.BN254OperatorSetInfo\",\"components\":[{\"name\":\"operatorInfoTreeRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"numOperators\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"aggregatePubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"totalWeights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]},{\"name\":\"operatorSetConfig\",\"type\":\"tuple\",\"internalType\":\"structICrossChainRegistryTypes.OperatorSetConfig\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxStalenessPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyCertificate\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"cert\",\"type\":\"tuple\",\"internalType\":\"structIBN254CertificateVerifierTypes.BN254Certificate\",\"components\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"nonSignerWitnesses\",\"type\":\"tuple[]\",\"internalType\":\"structIBN254CertificateVerifierTypes.BN254OperatorInfoWitness[]\",\"components\":[{\"name\":\"operatorIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorInfoProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"operatorInfo\",\"type\":\"tuple\",\"internalType\":\"structIOperatorTableCalculatorTypes.BN254OperatorInfo\",\"components\":[{\"name\":\"pubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}]}]}],\"outputs\":[{\"name\":\"totalSignedStakeWeights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyCertificateNominal\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"cert\",\"type\":\"tuple\",\"internalType\":\"structIBN254CertificateVerifierTypes.BN254Certificate\",\"components\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"nonSignerWitnesses\",\"type\":\"tuple[]\",\"internalType\":\"structIBN254CertificateVerifierTypes.BN254OperatorInfoWitness[]\",\"components\":[{\"name\":\"operatorIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorInfoProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"operatorInfo\",\"type\":\"tuple\",\"internalType\":\"structIOperatorTableCalculatorTypes.BN254OperatorInfo\",\"components\":[{\"name\":\"pubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}]}]},{\"name\":\"totalStakeNominalThresholds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyCertificateProportion\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"cert\",\"type\":\"tuple\",\"internalType\":\"structIBN254CertificateVerifierTypes.BN254Certificate\",\"components\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"nonSignerWitnesses\",\"type\":\"tuple[]\",\"internalType\":\"structIBN254CertificateVerifierTypes.BN254OperatorInfoWitness[]\",\"components\":[{\"name\":\"operatorIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorInfoProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"operatorInfo\",\"type\":\"tuple\",\"internalType\":\"structIOperatorTableCalculatorTypes.BN254OperatorInfo\",\"components\":[{\"name\":\"pubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}]}]},{\"name\":\"totalStakeProportionThresholds\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MaxStalenessPeriodUpdated\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"maxStalenessPeriod\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSetOwnerUpdated\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TableUpdated\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"operatorSetInfo\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIOperatorTableCalculatorTypes.BN254OperatorSetInfo\",\"components\":[{\"name\":\"operatorInfoTreeRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"numOperators\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"aggregatePubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"totalWeights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ArrayLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CertificateStale\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECAddFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECMulFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECPairingFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EmptyRoot\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpModFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidIndex\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperatorIndex\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProofLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonSignerIndicesNotSorted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyTableUpdater\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReferenceTimestampDoesNotExist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RootDisabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"TableUpdateStale\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"VerificationFailed\",\"inputs\":[]}]",
	Bin: "0x60c060405234801561000f575f5ffd5b50604051612e71380380612e7183398101604081905261002e9161016a565b6001600160a01b0382166080528061004581610058565b60a0525061005161009e565b5050610294565b5f5f829050601f8151111561008b578260405163305a27a960e01b81526004016100829190610239565b60405180910390fd5b80516100968261026e565b179392505050565b5f54610100900460ff16156101055760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b6064820152608401610082565b5f5460ff90811614610154575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b634e487b7160e01b5f52604160045260245ffd5b5f5f6040838503121561017b575f5ffd5b82516001600160a01b0381168114610191575f5ffd5b60208401519092506001600160401b038111156101ac575f5ffd5b8301601f810185136101bc575f5ffd5b80516001600160401b038111156101d5576101d5610156565b604051601f8201601f19908116603f011681016001600160401b038111828210171561020357610203610156565b60405281815282820160200187101561021a575f5ffd5b8160208401602083015e5f602083830101528093505050509250929050565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b8051602080830151919081101561028e575f198160200360031b1b821691505b50919050565b60805160a051612bae6102c35f395f6106d501525f81816102a6015281816108200152610ff50152612bae5ff3fe608060405234801561000f575f5ffd5b5060043610610132575f3560e01c80635ddb9b5b116100b4578063848189201161007957806384818920146102f3578063a2c902f514610306578063a2f2e24d1461030e578063cd83a72b14610321578063dd2ae1b914610334578063eb39e68f14610347575f5ffd5b80635ddb9b5b146102515780636141879e146102795780636738c40b1461028c57806368d6e081146102a15780637d1d1f5b146102e0575f5ffd5b806323c2a3cb116100fa57806323c2a3cb146101e357806326af6a3c146101f6578063538a37901461021657806354fd4d50146102295780635be872741461023e575f5ffd5b8063017d797414610136578063080b71501461015e578063121409ea1461017e57806318467434146101985780631a18746c146101b9575b5f5ffd5b610149610144366004612461565b610367565b60405190151581526020015b60405180910390f35b61017161016c36600461253c565b6104f8565b6040516101559190612587565b610186608e81565b60405160ff9091168152602001610155565b6101ab6101a63660046125be565b61050d565b604051908152602001610155565b6101cc6101c73660046125e6565b610570565b604080519215158352901515602083015201610155565b6101ab6101f1366004612634565b610591565b610209610204366004612666565b6105c7565b60405161015591906126db565b6101ab610224366004612708565b61067f565b6102316106ce565b6040516101559190612739565b61014961024c366004612666565b6106fe565b61026461025f36600461276e565b6107c9565b60405163ffffffff9091168152602001610155565b61026461028736600461276e565b6107ef565b61029f61029a36600461279e565b610815565b005b6102c87f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b039091168152602001610155565b6101716102ee366004612634565b6109fb565b6102c861030136600461276e565b610a7b565b610186607581565b6101ab61031c366004612857565b610aa4565b61014961032f366004612634565b610abb565b6101496103423660046128c3565b610af1565b61035a610355366004612634565b610b84565b604051610155919061297b565b5f5f6103738585610c49565b90505f61037f86610de6565b5f8181526004602081815260408084208a5163ffffffff16855282528084208151608081018352815481526001820154818501528251808401845260028301548152600383015481860152818401529381018054835181860281018601909452808452969750949593949093606086019383018282801561041d57602002820191905f5260205f20905b815481526020019060010190808311610409575b50505050508152505090505f8160600151905085518451146104525760405163512509d360e11b815260040160405180910390fd5b5f5b84518110156104e7575f6127108883815181106104735761047361298d565b602002602001015161ffff168484815181106104915761049161298d565b60200260200101516104a391906129b5565b6104ad91906129e0565b9050808683815181106104c2576104c261298d565b602002602001015110156104de575f96505050505050506104f1565b50600101610454565b5060019450505050505b9392505050565b60606105048383610c49565b90505b92915050565b604080517fd9f77a423768f4b0526fa60a7c732334516a93f1d228dce50ad804ea74ced36e602082015263ffffffff841691810191909152606081018290525f906080015b60405160208183030381529060405280519060200120905092915050565b5f5f61058486848787600162061a80610e49565b9150915094509492505050565b5f5f61059c84610de6565b5f90815260046020908152604080832063ffffffff8716845290915290206001015491505092915050565b6105cf611dfa565b5f6105d985610de6565b5f81815260056020908152604080832063ffffffff89168452825280832087845282529182902082516080810184528154818501908152600183015460608301528152600282018054855181860281018601909652808652959650909491938584019390929083018282801561066c57602002820191905f5260205f20905b815481526020019060010190808311610658575b5050505050815250509150509392505050565b5f60758260405160200161069391906126db565b60408051601f19818403018152908290526106b192916020016129f3565b604051602081830303815290604052805190602001209050919050565b60606106f97f0000000000000000000000000000000000000000000000000000000000000000610f11565b905090565b5f5f61070985610de6565b5f81815260056020908152604080832063ffffffff8916845282528083208784528252808320815160808101835281548184019081526001830154606083015281526002820180548451818702810187019095528085529697509495909491938581019392919083018282801561079d57602002820191905f5260205f20905b815481526020019060010190808311610789575b50505091909252505081515191925050158015906107bf575080516020015115155b9695505050505050565b5f5f6107d483610de6565b5f9081526003602052604090205463ffffffff169392505050565b5f5f6107fa83610de6565b5f9081526002602052604090205463ffffffff169392505050565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461085e5760405163030c1b6b60e11b815260040160405180910390fd5b5f6108766108713687900387018761276e565b610de6565b5f8181526003602052604090205490915063ffffffff908116908516116108b057604051632f20889f60e01b815260040160405180910390fd5b5f81815260046020818152604080842063ffffffff8916855282529283902086518155818701516001820155928601518051600285015581015160038401556060860151805187949361090893908501920190611e24565b5050505f818152600360209081526040909120805463ffffffff191663ffffffff871617905561093a90830183612a1c565b5f8281526001602090815260409182902080546001600160a01b0319166001600160a01b03949094169390931790925561097991908401908401612a35565b5f828152600260209081526040808320805463ffffffff191663ffffffff958616179055600682528083209388168352929052819020805460ff19166001179055517f93e6bea1c9b5dce4a5c07b00261e956df2a4a253d9ab6ca070ca2037d72ada9e906109ec90879087908790612a4e565b60405180910390a15050505050565b60605f610a0784610de6565b5f81815260046020818152604080842063ffffffff891685528252928390209091018054835181840281018401909452808452939450919290830182828015610a6d57602002820191905f5260205f20905b815481526020019060010190808311610a59575b505050505091505092915050565b5f5f610a8683610de6565b5f908152600160205260409020546001600160a01b03169392505050565b5f608e838360405160200161055293929190612a9a565b5f5f610ac684610de6565b5f90815260066020908152604080832063ffffffff8716845290915290205460ff1691505092915050565b5f5f610afd8585610c49565b90508251815114610b215760405163512509d360e11b815260040160405180910390fd5b5f5b8151811015610b7857838181518110610b3e57610b3e61298d565b6020026020010151828281518110610b5857610b5861298d565b60200260200101511015610b70575f925050506104f1565b600101610b23565b50600195945050505050565b610b8c611e6d565b5f610b9684610de6565b5f81815260046020818152604080842063ffffffff8916855282529283902083516080810185528154815260018201548184015284518086018652600283015481526003830154818501528186015292810180548551818502810185019096528086529596509294909360608601939092909190830182828015610c3757602002820191905f5260205f20905b815481526020019060010190808311610c23575b50505050508152505091505092915050565b6060610c53611e9f565b610c5c84610de6565b8082528351610c6b9190610f4e565b80515f908152600460208181526040808420875163ffffffff1685528252928390208351608081018552815481526001820154818401528451808601865260028301548152600383015481850152818601529281018054855181850281018501909652808652939491936060860193830182828015610d0757602002820191905f5260205f20905b815481526020019060010190808311610cf3575b505050919092525050506020820181905260600151516001600160401b03811115610d3457610d34611fcd565b604051908082528060200260200182016040528015610d5d578160200160208202803683370190505b5060408201525f5b81602001516060015151811015610dc1578160200151606001518181518110610d9057610d9061298d565b602002602001015182604001518281518110610dae57610dae61298d565b6020908102919091010152600101610d65565b50610dcc8184611088565b6060820152610ddb81846111e5565b604001519392505050565b5f815f0151826020015163ffffffff16604051602001610e3192919060609290921b6bffffffffffffffffffffffff1916825260a01b6001600160a01b031916601482015260200190565b60405160208183030381529060405261050790612ac1565b5f5f5f610e5589611264565b90505f610e648a89898c6112ee565b90505f610e7b610e748a846113a2565b8b9061140a565b90505f610ebd610eb684610eb06040805180820182525f80825260209182015281518083019092526001825260029082015290565b906113a2565b859061140a565b90508715610ee257610ed982610ed161147e565b838c8b61153e565b96509450610f02565b610ef582610eee61147e565b838c611752565b95508515610f0257600194505b50505050965096945050505050565b60605f610f1d83611989565b6040805160208082528183019092529192505f91906020820181803683375050509182525060208101929092525090565b5f8281526002602052604090205463ffffffff16801580610f7e5750610f748183612ae4565b63ffffffff164211155b610f9b5760405163640fcd6b60e11b815260040160405180910390fd5b5f83815260066020908152604080832063ffffffff8616845290915290205460ff16610fda57604051630cad17b760e31b815260040160405180910390fd5b60405163193877e160e21b815263ffffffff831660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906364e1df8490602401602060405180830381865afa158015611042573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906110669190612b00565b61108357604051631b14174b60e01b815260040160405180910390fd5b505050565b6040805180820182525f80825260209182018190528251808401909352808352908201819052805b8360800151518110156111dd575f846080015182815181106110d4576110d461298d565b602002602001015190505f82111561111057805163ffffffff80851691161161111057604051631d8c4d1760e31b815260040160405180910390fd5b6020808701510151815163ffffffff161061113e576040516301fa53c760e11b815260040160405180910390fd5b855185515f9161114e91846119b0565b805190915061115e90869061140a565b94505f5b8160200151518110156111d0578760400151518110156111c857816020015181815181106111925761119261298d565b6020026020010151886040015182815181106111b0576111b061298d565b602002602001018181516111c49190612b1f565b9052505b600101611162565b50505191506001016110b0565b505092915050565b5f6112056111f68460600151611b27565b6020850151604001519061140a565b90505f611219835f0151846020015161050d565b90505f5f611231838587606001518860400151610570565b9150915081801561123f5750805b61125c5760405163439cc0cd60e01b815260040160405180910390fd5b505050505050565b604080518082019091525f80825260208201525f80806112915f516020612b595f395f51905f5286612b32565b90505b61129d81611bbd565b90935091505f516020612b595f395f51905f5282830983036112d5576040805180820190915290815260208101919091529392505050565b5f516020612b595f395f51905f52600182089050611294565b8251602080850151845180519083015186840151805190850151875188870151604080519889018e90528801989098526060870195909552608086019390935260a085019190915260c084015260e08301526101008201526101208101919091525f907f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000019061014001604051602081830303815290604052805190602001205f1c6113999190612b32565b95945050505050565b604080518082019091525f80825260208201526113bd611ee4565b835181526020808501519082015260408082018490525f908360608460076107d05a03fa905080806113eb57fe5b50806111dd57604051632319df1960e11b815260040160405180910390fd5b604080518082019091525f8082526020820152611425611f02565b835181526020808501518183015283516040808401919091529084015160608301525f908360808460066107d05a03fa9050808061145f57fe5b50806111dd5760405163d4b68fd760e01b815260040160405180910390fd5b611486611f20565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208381019190915281019190915290565b6040805180820182528681526020808201869052825180840190935286835282018490525f9182919061156f611f40565b5f5b6002811015611726575f6115868260066129b5565b905084826002811061159a5761159a61298d565b602002015151836115ab835f612b45565b600c81106115bb576115bb61298d565b60200201528482600281106115d2576115d261298d565b602002015160200151838260016115e99190612b45565b600c81106115f9576115f961298d565b60200201528382600281106116105761161061298d565b6020020151515183611623836002612b45565b600c81106116335761163361298d565b602002015283826002811061164a5761164a61298d565b6020020151516001602002015183611663836003612b45565b600c81106116735761167361298d565b602002015283826002811061168a5761168a61298d565b6020020151602001515f600281106116a4576116a461298d565b6020020151836116b5836004612b45565b600c81106116c5576116c561298d565b60200201528382600281106116dc576116dc61298d565b6020020151602001516001600281106116f7576116f761298d565b602002015183611708836005612b45565b600c81106117185761171861298d565b602002015250600101611571565b5061172f611f5f565b5f6020826101808560088cfa9151919c9115159b50909950505050505050505050565b6040805180820182528581526020808201859052825180840190935285835282018390525f91611780611f40565b5f5b6002811015611937575f6117978260066129b5565b90508482600281106117ab576117ab61298d565b602002015151836117bc835f612b45565b600c81106117cc576117cc61298d565b60200201528482600281106117e3576117e361298d565b602002015160200151838260016117fa9190612b45565b600c811061180a5761180a61298d565b60200201528382600281106118215761182161298d565b6020020151515183611834836002612b45565b600c81106118445761184461298d565b602002015283826002811061185b5761185b61298d565b6020020151516001602002015183611874836003612b45565b600c81106118845761188461298d565b602002015283826002811061189b5761189b61298d565b6020020151602001515f600281106118b5576118b561298d565b6020020151836118c6836004612b45565b600c81106118d6576118d661298d565b60200201528382600281106118ed576118ed61298d565b6020020151602001516001600281106119085761190861298d565b602002015183611919836005612b45565b600c81106119295761192961298d565b602002015250600101611782565b50611940611f5f565b5f6020826101808560086107d05a03fa9050808061195a57fe5b5080611979576040516324ccc79360e21b815260040160405180910390fd5b5051151598975050505050505050565b5f60ff8216601f81111561050757604051632cd44ac360e21b815260040160405180910390fd5b6119b8611dfa565b5f84815260056020908152604080832063ffffffff8088168552908352818420865190911684528252808320815160808101835281548184019081526001830154606083015281526002820180548451818702810187019095528085529194929385840193909290830182828015611a4d57602002820191905f5260205f20905b815481526020019060010190808311611a39575b5050509190925250508151519192505f911515905080611a71575081516020015115155b905080611b1a575f611a918787875f015188604001518960200151611c39565b905080611ab15760405163439cc0cd60e01b815260040160405180910390fd5b6040808601515f8981526005602090815283822063ffffffff808c1684529082528483208a5190911683528152929020815180518255830151600182015582820151805192939192611b099260028501920190611e24565b509050508460400151935050611b1e565b8192505b50509392505050565b604080518082019091525f80825260208201528151158015611b4b57506020820151155b15611b68575050604080518082019091525f808252602082015290565b6040518060400160405280835f015181526020015f516020612b595f395f51905f528460200151611b999190612b32565b611bb0905f516020612b595f395f51905f52612b1f565b905292915050565b919050565b5f80805f516020612b595f395f51905f5260035f516020612b595f395f51905f52865f516020612b595f395f51905f52888909090890505f611c2d827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f525f516020612b595f395f51905f52611c87565b91959194509092505050565b5f5f611c448461067f565b5f88815260046020908152604080832063ffffffff808c168552925290912054919250611c7b908590839085908a811690611d0016565b98975050505050505050565b5f5f611c91611f5f565b611c99611f7d565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa92508280611cd657fe5b5082611cf55760405163d51edae360e01b815260040160405180910390fd5b505195945050505050565b5f83611d1f576040516329e7276760e11b815260040160405180910390fd5b83611d2b868585611d35565b1495945050505050565b5f83515f03611d455750816104f1565b60208451611d539190612b32565b15611d71576040516313717da960e21b815260040160405180910390fd5b8260205b85518111611dd257611d88600285612b32565b5f03611da957815f528086015160205260405f209150600284049350611dc0565b808601515f528160205260405f2091506002840493505b611dcb602082612b45565b9050611d75565b508215611df2576040516363df817160e01b815260040160405180910390fd5b949350505050565b604080516080810182525f91810182815260608201929092529081905b8152602001606081525090565b828054828255905f5260205f20908101928215611e5d579160200282015b82811115611e5d578251825591602001919060010190611e42565b50611e69929150611f9b565b5090565b60405180608001604052805f81526020015f8152602001611e1760405180604001604052805f81526020015f81525090565b60405180608001604052805f8152602001611eb8611e6d565b815260200160608152602001611edf60405180604001604052805f81526020015f81525090565b905290565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b6040518060400160405280611f33611faf565b8152602001611edf611faf565b604051806101800160405280600c906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b5b80821115611e69575f8155600101611f9c565b60405180604001604052806002906020820280368337509192915050565b634e487b7160e01b5f52604160045260245ffd5b604080519081016001600160401b038111828210171561200357612003611fcd565b60405290565b60405160a081016001600160401b038111828210171561200357612003611fcd565b604051606081016001600160401b038111828210171561200357612003611fcd565b604051608081016001600160401b038111828210171561200357612003611fcd565b604051601f8201601f191681016001600160401b038111828210171561209757612097611fcd565b604052919050565b80356001600160a01b0381168114611bb8575f5ffd5b803563ffffffff81168114611bb8575f5ffd5b5f604082840312156120d8575f5ffd5b6120e0611fe1565b90506120eb8261209f565b81526120f9602083016120b5565b602082015292915050565b5f60408284031215612114575f5ffd5b61211c611fe1565b823581526020928301359281019290925250919050565b5f82601f830112612142575f5ffd5b61214a611fe1565b80604084018581111561215b575f5ffd5b845b8181101561217557803584526020938401930161215d565b509095945050505050565b5f60808284031215612190575f5ffd5b612198611fe1565b90506121a48383612133565b81526120f98360408401612133565b5f6001600160401b038211156121cb576121cb611fcd565b5060051b60200190565b5f82601f8301126121e4575f5ffd5b81356121f76121f2826121b3565b61206f565b8082825260208201915060208360051b860101925085831115612218575f5ffd5b602085015b8381101561223557803583526020928301920161221d565b5095945050505050565b5f6060828403121561224f575f5ffd5b612257611fe1565b90506122638383612104565b815260408201356001600160401b0381111561227d575f5ffd5b612289848285016121d5565b60208301525092915050565b5f61012082840312156122a6575f5ffd5b6122ae612009565b90506122b9826120b5565b8152602082810135908201526122d28360408401612104565b60408201526122e48360808401612180565b60608201526101008201356001600160401b03811115612302575f5ffd5b8201601f81018413612312575f5ffd5b80356123206121f2826121b3565b8082825260208201915060208360051b850101925086831115612341575f5ffd5b602084015b838110156124515780356001600160401b03811115612363575f5ffd5b85016060818a03601f19011215612378575f5ffd5b61238061202b565b61238c602083016120b5565b815260408201356001600160401b038111156123a6575f5ffd5b82016020810190603f018b136123ba575f5ffd5b80356001600160401b038111156123d3576123d3611fcd565b6123e6601f8201601f191660200161206f565b8181528c60208385010111156123fa575f5ffd5b816020840160208301375f6020838301015280602085015250505060608201356001600160401b0381111561242d575f5ffd5b61243c8b60208386010161223f565b60408301525084525060209283019201612346565b5060808501525091949350505050565b5f5f5f60808486031215612473575f5ffd5b61247d85856120c8565b925060408401356001600160401b03811115612497575f5ffd5b6124a386828701612295565b92505060608401356001600160401b038111156124be575f5ffd5b8401601f810186136124ce575f5ffd5b80356124dc6121f2826121b3565b8082825260208201915060208360051b8501019250888311156124fd575f5ffd5b6020840193505b8284101561252e57833561ffff8116811461251d575f5ffd5b825260209384019390910190612504565b809450505050509250925092565b5f5f6060838503121561254d575f5ffd5b61255784846120c8565b915060408301356001600160401b03811115612571575f5ffd5b61257d85828601612295565b9150509250929050565b602080825282518282018190525f918401906040840190835b818110156121755783518352602093840193909201916001016125a0565b5f5f604083850312156125cf575f5ffd5b6125d8836120b5565b946020939093013593505050565b5f5f5f5f61012085870312156125fa575f5ffd5b8435935061260b8660208701612104565b925061261a8660608701612180565b91506126298660e08701612104565b905092959194509250565b5f5f60608385031215612645575f5ffd5b61264f84846120c8565b915061265d604084016120b5565b90509250929050565b5f5f5f60808486031215612678575f5ffd5b61268285856120c8565b9250612690604085016120b5565b929592945050506060919091013590565b5f8151808452602084019350602083015f5b828110156126d15781518652602095860195909101906001016126b3565b5093949350505050565b60208082528251805183830152015160408201525f6020830151606080840152611df260808401826126a1565b5f60208284031215612718575f5ffd5b81356001600160401b0381111561272d575f5ffd5b611df28482850161223f565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b5f6040828403121561277e575f5ffd5b61050483836120c8565b5f60408284031215612798575f5ffd5b50919050565b5f5f5f5f60c085870312156127b1575f5ffd5b6127bb8686612788565b93506127c9604086016120b5565b925060608501356001600160401b038111156127e3575f5ffd5b850160a081880312156127f4575f5ffd5b6127fc61204d565b81358152602080830135908201526128178860408401612104565b604082015260808201356001600160401b03811115612834575f5ffd5b612840898285016121d5565b606083015250925061262990508660808701612788565b5f5f60208385031215612868575f5ffd5b82356001600160401b0381111561287d575f5ffd5b8301601f8101851361288d575f5ffd5b80356001600160401b038111156128a2575f5ffd5b8560208284010111156128b3575f5ffd5b6020919091019590945092505050565b5f5f5f608084860312156128d5575f5ffd5b6128df85856120c8565b925060408401356001600160401b038111156128f9575f5ffd5b61290586828701612295565b92505060608401356001600160401b03811115612920575f5ffd5b61292c868287016121d5565b9150509250925092565b80518252602081015160208301525f6040820151612961604085018280518252602090810151910152565b50606082015160a06080850152611df260a08501826126a1565b602081525f6105046020830184612936565b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52601160045260245ffd5b8082028115828204841417610507576105076129a1565b634e487b7160e01b5f52601260045260245ffd5b5f826129ee576129ee6129cc565b500490565b60ff60f81b8360f81b1681525f82518060208501600185015e5f92016001019182525092915050565b5f60208284031215612a2c575f5ffd5b6105048261209f565b5f60208284031215612a45575f5ffd5b610504826120b5565b6001600160a01b03612a5f8561209f565b16815263ffffffff612a73602086016120b5565b16602082015263ffffffff83166040820152608060608201525f6113996080830184612936565b60f884901b6001600160f81b0319168152818360018301375f910160010190815292915050565b80516020808301519190811015612798575f1960209190910360031b1b16919050565b63ffffffff8181168382160190811115610507576105076129a1565b5f60208284031215612b10575f5ffd5b815180151581146104f1575f5ffd5b81810381811115610507576105076129a1565b5f82612b4057612b406129cc565b500690565b80820180821115610507576105076129a156fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47a26469706673582212207db5218ea602fe4055ee92b6fa3a433a57d363164f08f843d254f0ebb2345ab464736f6c634300081b0033",
}

// BN254CertificateVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use BN254CertificateVerifierMetaData.ABI instead.
var BN254CertificateVerifierABI = BN254CertificateVerifierMetaData.ABI

// BN254CertificateVerifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BN254CertificateVerifierMetaData.Bin instead.
var BN254CertificateVerifierBin = BN254CertificateVerifierMetaData.Bin

// DeployBN254CertificateVerifier deploys a new Ethereum contract, binding an instance of BN254CertificateVerifier to it.
func DeployBN254CertificateVerifier(auth *bind.TransactOpts, backend bind.ContractBackend, _operatorTableUpdater common.Address, _version string) (common.Address, *types.Transaction, *BN254CertificateVerifier, error) {
	parsed, err := BN254CertificateVerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BN254CertificateVerifierBin), backend, _operatorTableUpdater, _version)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BN254CertificateVerifier{BN254CertificateVerifierCaller: BN254CertificateVerifierCaller{contract: contract}, BN254CertificateVerifierTransactor: BN254CertificateVerifierTransactor{contract: contract}, BN254CertificateVerifierFilterer: BN254CertificateVerifierFilterer{contract: contract}}, nil
}

// BN254CertificateVerifier is an auto generated Go binding around an Ethereum contract.
type BN254CertificateVerifier struct {
	BN254CertificateVerifierCaller     // Read-only binding to the contract
	BN254CertificateVerifierTransactor // Write-only binding to the contract
	BN254CertificateVerifierFilterer   // Log filterer for contract events
}

// BN254CertificateVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type BN254CertificateVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BN254CertificateVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BN254CertificateVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BN254CertificateVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BN254CertificateVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BN254CertificateVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BN254CertificateVerifierSession struct {
	Contract     *BN254CertificateVerifier // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// BN254CertificateVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BN254CertificateVerifierCallerSession struct {
	Contract *BN254CertificateVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// BN254CertificateVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BN254CertificateVerifierTransactorSession struct {
	Contract     *BN254CertificateVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// BN254CertificateVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type BN254CertificateVerifierRaw struct {
	Contract *BN254CertificateVerifier // Generic contract binding to access the raw methods on
}

// BN254CertificateVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BN254CertificateVerifierCallerRaw struct {
	Contract *BN254CertificateVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// BN254CertificateVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BN254CertificateVerifierTransactorRaw struct {
	Contract *BN254CertificateVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBN254CertificateVerifier creates a new instance of BN254CertificateVerifier, bound to a specific deployed contract.
func NewBN254CertificateVerifier(address common.Address, backend bind.ContractBackend) (*BN254CertificateVerifier, error) {
	contract, err := bindBN254CertificateVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BN254CertificateVerifier{BN254CertificateVerifierCaller: BN254CertificateVerifierCaller{contract: contract}, BN254CertificateVerifierTransactor: BN254CertificateVerifierTransactor{contract: contract}, BN254CertificateVerifierFilterer: BN254CertificateVerifierFilterer{contract: contract}}, nil
}

// NewBN254CertificateVerifierCaller creates a new read-only instance of BN254CertificateVerifier, bound to a specific deployed contract.
func NewBN254CertificateVerifierCaller(address common.Address, caller bind.ContractCaller) (*BN254CertificateVerifierCaller, error) {
	contract, err := bindBN254CertificateVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BN254CertificateVerifierCaller{contract: contract}, nil
}

// NewBN254CertificateVerifierTransactor creates a new write-only instance of BN254CertificateVerifier, bound to a specific deployed contract.
func NewBN254CertificateVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*BN254CertificateVerifierTransactor, error) {
	contract, err := bindBN254CertificateVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BN254CertificateVerifierTransactor{contract: contract}, nil
}

// NewBN254CertificateVerifierFilterer creates a new log filterer instance of BN254CertificateVerifier, bound to a specific deployed contract.
func NewBN254CertificateVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*BN254CertificateVerifierFilterer, error) {
	contract, err := bindBN254CertificateVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BN254CertificateVerifierFilterer{contract: contract}, nil
}

// bindBN254CertificateVerifier binds a generic wrapper to an already deployed contract.
func bindBN254CertificateVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BN254CertificateVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BN254CertificateVerifier *BN254CertificateVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BN254CertificateVerifier.Contract.BN254CertificateVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BN254CertificateVerifier *BN254CertificateVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BN254CertificateVerifier.Contract.BN254CertificateVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BN254CertificateVerifier *BN254CertificateVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BN254CertificateVerifier.Contract.BN254CertificateVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BN254CertificateVerifier *BN254CertificateVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BN254CertificateVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BN254CertificateVerifier *BN254CertificateVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BN254CertificateVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BN254CertificateVerifier *BN254CertificateVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BN254CertificateVerifier.Contract.contract.Transact(opts, method, params...)
}

// OPERATORINFOLEAFSALT is a free data retrieval call binding the contract method 0xa2c902f5.
//
// Solidity: function OPERATOR_INFO_LEAF_SALT() view returns(uint8)
func (_BN254CertificateVerifier *BN254CertificateVerifierCaller) OPERATORINFOLEAFSALT(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BN254CertificateVerifier.contract.Call(opts, &out, "OPERATOR_INFO_LEAF_SALT")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// OPERATORINFOLEAFSALT is a free data retrieval call binding the contract method 0xa2c902f5.
//
// Solidity: function OPERATOR_INFO_LEAF_SALT() view returns(uint8)
func (_BN254CertificateVerifier *BN254CertificateVerifierSession) OPERATORINFOLEAFSALT() (uint8, error) {
	return _BN254CertificateVerifier.Contract.OPERATORINFOLEAFSALT(&_BN254CertificateVerifier.CallOpts)
}

// OPERATORINFOLEAFSALT is a free data retrieval call binding the contract method 0xa2c902f5.
//
// Solidity: function OPERATOR_INFO_LEAF_SALT() view returns(uint8)
func (_BN254CertificateVerifier *BN254CertificateVerifierCallerSession) OPERATORINFOLEAFSALT() (uint8, error) {
	return _BN254CertificateVerifier.Contract.OPERATORINFOLEAFSALT(&_BN254CertificateVerifier.CallOpts)
}

// OPERATORTABLELEAFSALT is a free data retrieval call binding the contract method 0x121409ea.
//
// Solidity: function OPERATOR_TABLE_LEAF_SALT() view returns(uint8)
func (_BN254CertificateVerifier *BN254CertificateVerifierCaller) OPERATORTABLELEAFSALT(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BN254CertificateVerifier.contract.Call(opts, &out, "OPERATOR_TABLE_LEAF_SALT")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// OPERATORTABLELEAFSALT is a free data retrieval call binding the contract method 0x121409ea.
//
// Solidity: function OPERATOR_TABLE_LEAF_SALT() view returns(uint8)
func (_BN254CertificateVerifier *BN254CertificateVerifierSession) OPERATORTABLELEAFSALT() (uint8, error) {
	return _BN254CertificateVerifier.Contract.OPERATORTABLELEAFSALT(&_BN254CertificateVerifier.CallOpts)
}

// OPERATORTABLELEAFSALT is a free data retrieval call binding the contract method 0x121409ea.
//
// Solidity: function OPERATOR_TABLE_LEAF_SALT() view returns(uint8)
func (_BN254CertificateVerifier *BN254CertificateVerifierCallerSession) OPERATORTABLELEAFSALT() (uint8, error) {
	return _BN254CertificateVerifier.Contract.OPERATORTABLELEAFSALT(&_BN254CertificateVerifier.CallOpts)
}

// CalculateCertificateDigest is a free data retrieval call binding the contract method 0x18467434.
//
// Solidity: function calculateCertificateDigest(uint32 referenceTimestamp, bytes32 messageHash) pure returns(bytes32)
func (_BN254CertificateVerifier *BN254CertificateVerifierCaller) CalculateCertificateDigest(opts *bind.CallOpts, referenceTimestamp uint32, messageHash [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _BN254CertificateVerifier.contract.Call(opts, &out, "calculateCertificateDigest", referenceTimestamp, messageHash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateCertificateDigest is a free data retrieval call binding the contract method 0x18467434.
//
// Solidity: function calculateCertificateDigest(uint32 referenceTimestamp, bytes32 messageHash) pure returns(bytes32)
func (_BN254CertificateVerifier *BN254CertificateVerifierSession) CalculateCertificateDigest(referenceTimestamp uint32, messageHash [32]byte) ([32]byte, error) {
	return _BN254CertificateVerifier.Contract.CalculateCertificateDigest(&_BN254CertificateVerifier.CallOpts, referenceTimestamp, messageHash)
}

// CalculateCertificateDigest is a free data retrieval call binding the contract method 0x18467434.
//
// Solidity: function calculateCertificateDigest(uint32 referenceTimestamp, bytes32 messageHash) pure returns(bytes32)
func (_BN254CertificateVerifier *BN254CertificateVerifierCallerSession) CalculateCertificateDigest(referenceTimestamp uint32, messageHash [32]byte) ([32]byte, error) {
	return _BN254CertificateVerifier.Contract.CalculateCertificateDigest(&_BN254CertificateVerifier.CallOpts, referenceTimestamp, messageHash)
}

// CalculateOperatorInfoLeaf is a free data retrieval call binding the contract method 0x538a3790.
//
// Solidity: function calculateOperatorInfoLeaf(((uint256,uint256),uint256[]) operatorInfo) pure returns(bytes32)
func (_BN254CertificateVerifier *BN254CertificateVerifierCaller) CalculateOperatorInfoLeaf(opts *bind.CallOpts, operatorInfo IOperatorTableCalculatorTypesBN254OperatorInfo) ([32]byte, error) {
	var out []interface{}
	err := _BN254CertificateVerifier.contract.Call(opts, &out, "calculateOperatorInfoLeaf", operatorInfo)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateOperatorInfoLeaf is a free data retrieval call binding the contract method 0x538a3790.
//
// Solidity: function calculateOperatorInfoLeaf(((uint256,uint256),uint256[]) operatorInfo) pure returns(bytes32)
func (_BN254CertificateVerifier *BN254CertificateVerifierSession) CalculateOperatorInfoLeaf(operatorInfo IOperatorTableCalculatorTypesBN254OperatorInfo) ([32]byte, error) {
	return _BN254CertificateVerifier.Contract.CalculateOperatorInfoLeaf(&_BN254CertificateVerifier.CallOpts, operatorInfo)
}

// CalculateOperatorInfoLeaf is a free data retrieval call binding the contract method 0x538a3790.
//
// Solidity: function calculateOperatorInfoLeaf(((uint256,uint256),uint256[]) operatorInfo) pure returns(bytes32)
func (_BN254CertificateVerifier *BN254CertificateVerifierCallerSession) CalculateOperatorInfoLeaf(operatorInfo IOperatorTableCalculatorTypesBN254OperatorInfo) ([32]byte, error) {
	return _BN254CertificateVerifier.Contract.CalculateOperatorInfoLeaf(&_BN254CertificateVerifier.CallOpts, operatorInfo)
}

// CalculateOperatorTableLeaf is a free data retrieval call binding the contract method 0xa2f2e24d.
//
// Solidity: function calculateOperatorTableLeaf(bytes operatorTableBytes) pure returns(bytes32)
func (_BN254CertificateVerifier *BN254CertificateVerifierCaller) CalculateOperatorTableLeaf(opts *bind.CallOpts, operatorTableBytes []byte) ([32]byte, error) {
	var out []interface{}
	err := _BN254CertificateVerifier.contract.Call(opts, &out, "calculateOperatorTableLeaf", operatorTableBytes)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateOperatorTableLeaf is a free data retrieval call binding the contract method 0xa2f2e24d.
//
// Solidity: function calculateOperatorTableLeaf(bytes operatorTableBytes) pure returns(bytes32)
func (_BN254CertificateVerifier *BN254CertificateVerifierSession) CalculateOperatorTableLeaf(operatorTableBytes []byte) ([32]byte, error) {
	return _BN254CertificateVerifier.Contract.CalculateOperatorTableLeaf(&_BN254CertificateVerifier.CallOpts, operatorTableBytes)
}

// CalculateOperatorTableLeaf is a free data retrieval call binding the contract method 0xa2f2e24d.
//
// Solidity: function calculateOperatorTableLeaf(bytes operatorTableBytes) pure returns(bytes32)
func (_BN254CertificateVerifier *BN254CertificateVerifierCallerSession) CalculateOperatorTableLeaf(operatorTableBytes []byte) ([32]byte, error) {
	return _BN254CertificateVerifier.Contract.CalculateOperatorTableLeaf(&_BN254CertificateVerifier.CallOpts, operatorTableBytes)
}

// GetNonsignerOperatorInfo is a free data retrieval call binding the contract method 0x26af6a3c.
//
// Solidity: function getNonsignerOperatorInfo((address,uint32) operatorSet, uint32 referenceTimestamp, uint256 operatorIndex) view returns(((uint256,uint256),uint256[]))
func (_BN254CertificateVerifier *BN254CertificateVerifierCaller) GetNonsignerOperatorInfo(opts *bind.CallOpts, operatorSet OperatorSet, referenceTimestamp uint32, operatorIndex *big.Int) (IOperatorTableCalculatorTypesBN254OperatorInfo, error) {
	var out []interface{}
	err := _BN254CertificateVerifier.contract.Call(opts, &out, "getNonsignerOperatorInfo", operatorSet, referenceTimestamp, operatorIndex)

	if err != nil {
		return *new(IOperatorTableCalculatorTypesBN254OperatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IOperatorTableCalculatorTypesBN254OperatorInfo)).(*IOperatorTableCalculatorTypesBN254OperatorInfo)

	return out0, err

}

// GetNonsignerOperatorInfo is a free data retrieval call binding the contract method 0x26af6a3c.
//
// Solidity: function getNonsignerOperatorInfo((address,uint32) operatorSet, uint32 referenceTimestamp, uint256 operatorIndex) view returns(((uint256,uint256),uint256[]))
func (_BN254CertificateVerifier *BN254CertificateVerifierSession) GetNonsignerOperatorInfo(operatorSet OperatorSet, referenceTimestamp uint32, operatorIndex *big.Int) (IOperatorTableCalculatorTypesBN254OperatorInfo, error) {
	return _BN254CertificateVerifier.Contract.GetNonsignerOperatorInfo(&_BN254CertificateVerifier.CallOpts, operatorSet, referenceTimestamp, operatorIndex)
}

// GetNonsignerOperatorInfo is a free data retrieval call binding the contract method 0x26af6a3c.
//
// Solidity: function getNonsignerOperatorInfo((address,uint32) operatorSet, uint32 referenceTimestamp, uint256 operatorIndex) view returns(((uint256,uint256),uint256[]))
func (_BN254CertificateVerifier *BN254CertificateVerifierCallerSession) GetNonsignerOperatorInfo(operatorSet OperatorSet, referenceTimestamp uint32, operatorIndex *big.Int) (IOperatorTableCalculatorTypesBN254OperatorInfo, error) {
	return _BN254CertificateVerifier.Contract.GetNonsignerOperatorInfo(&_BN254CertificateVerifier.CallOpts, operatorSet, referenceTimestamp, operatorIndex)
}

// GetOperatorCount is a free data retrieval call binding the contract method 0x23c2a3cb.
//
// Solidity: function getOperatorCount((address,uint32) operatorSet, uint32 referenceTimestamp) view returns(uint256)
func (_BN254CertificateVerifier *BN254CertificateVerifierCaller) GetOperatorCount(opts *bind.CallOpts, operatorSet OperatorSet, referenceTimestamp uint32) (*big.Int, error) {
	var out []interface{}
	err := _BN254CertificateVerifier.contract.Call(opts, &out, "getOperatorCount", operatorSet, referenceTimestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOperatorCount is a free data retrieval call binding the contract method 0x23c2a3cb.
//
// Solidity: function getOperatorCount((address,uint32) operatorSet, uint32 referenceTimestamp) view returns(uint256)
func (_BN254CertificateVerifier *BN254CertificateVerifierSession) GetOperatorCount(operatorSet OperatorSet, referenceTimestamp uint32) (*big.Int, error) {
	return _BN254CertificateVerifier.Contract.GetOperatorCount(&_BN254CertificateVerifier.CallOpts, operatorSet, referenceTimestamp)
}

// GetOperatorCount is a free data retrieval call binding the contract method 0x23c2a3cb.
//
// Solidity: function getOperatorCount((address,uint32) operatorSet, uint32 referenceTimestamp) view returns(uint256)
func (_BN254CertificateVerifier *BN254CertificateVerifierCallerSession) GetOperatorCount(operatorSet OperatorSet, referenceTimestamp uint32) (*big.Int, error) {
	return _BN254CertificateVerifier.Contract.GetOperatorCount(&_BN254CertificateVerifier.CallOpts, operatorSet, referenceTimestamp)
}

// GetOperatorSetInfo is a free data retrieval call binding the contract method 0xeb39e68f.
//
// Solidity: function getOperatorSetInfo((address,uint32) operatorSet, uint32 referenceTimestamp) view returns((bytes32,uint256,(uint256,uint256),uint256[]))
func (_BN254CertificateVerifier *BN254CertificateVerifierCaller) GetOperatorSetInfo(opts *bind.CallOpts, operatorSet OperatorSet, referenceTimestamp uint32) (IOperatorTableCalculatorTypesBN254OperatorSetInfo, error) {
	var out []interface{}
	err := _BN254CertificateVerifier.contract.Call(opts, &out, "getOperatorSetInfo", operatorSet, referenceTimestamp)

	if err != nil {
		return *new(IOperatorTableCalculatorTypesBN254OperatorSetInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IOperatorTableCalculatorTypesBN254OperatorSetInfo)).(*IOperatorTableCalculatorTypesBN254OperatorSetInfo)

	return out0, err

}

// GetOperatorSetInfo is a free data retrieval call binding the contract method 0xeb39e68f.
//
// Solidity: function getOperatorSetInfo((address,uint32) operatorSet, uint32 referenceTimestamp) view returns((bytes32,uint256,(uint256,uint256),uint256[]))
func (_BN254CertificateVerifier *BN254CertificateVerifierSession) GetOperatorSetInfo(operatorSet OperatorSet, referenceTimestamp uint32) (IOperatorTableCalculatorTypesBN254OperatorSetInfo, error) {
	return _BN254CertificateVerifier.Contract.GetOperatorSetInfo(&_BN254CertificateVerifier.CallOpts, operatorSet, referenceTimestamp)
}

// GetOperatorSetInfo is a free data retrieval call binding the contract method 0xeb39e68f.
//
// Solidity: function getOperatorSetInfo((address,uint32) operatorSet, uint32 referenceTimestamp) view returns((bytes32,uint256,(uint256,uint256),uint256[]))
func (_BN254CertificateVerifier *BN254CertificateVerifierCallerSession) GetOperatorSetInfo(operatorSet OperatorSet, referenceTimestamp uint32) (IOperatorTableCalculatorTypesBN254OperatorSetInfo, error) {
	return _BN254CertificateVerifier.Contract.GetOperatorSetInfo(&_BN254CertificateVerifier.CallOpts, operatorSet, referenceTimestamp)
}

// GetOperatorSetOwner is a free data retrieval call binding the contract method 0x84818920.
//
// Solidity: function getOperatorSetOwner((address,uint32) operatorSet) view returns(address)
func (_BN254CertificateVerifier *BN254CertificateVerifierCaller) GetOperatorSetOwner(opts *bind.CallOpts, operatorSet OperatorSet) (common.Address, error) {
	var out []interface{}
	err := _BN254CertificateVerifier.contract.Call(opts, &out, "getOperatorSetOwner", operatorSet)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOperatorSetOwner is a free data retrieval call binding the contract method 0x84818920.
//
// Solidity: function getOperatorSetOwner((address,uint32) operatorSet) view returns(address)
func (_BN254CertificateVerifier *BN254CertificateVerifierSession) GetOperatorSetOwner(operatorSet OperatorSet) (common.Address, error) {
	return _BN254CertificateVerifier.Contract.GetOperatorSetOwner(&_BN254CertificateVerifier.CallOpts, operatorSet)
}

// GetOperatorSetOwner is a free data retrieval call binding the contract method 0x84818920.
//
// Solidity: function getOperatorSetOwner((address,uint32) operatorSet) view returns(address)
func (_BN254CertificateVerifier *BN254CertificateVerifierCallerSession) GetOperatorSetOwner(operatorSet OperatorSet) (common.Address, error) {
	return _BN254CertificateVerifier.Contract.GetOperatorSetOwner(&_BN254CertificateVerifier.CallOpts, operatorSet)
}

// GetTotalStakeWeights is a free data retrieval call binding the contract method 0x7d1d1f5b.
//
// Solidity: function getTotalStakeWeights((address,uint32) operatorSet, uint32 referenceTimestamp) view returns(uint256[])
func (_BN254CertificateVerifier *BN254CertificateVerifierCaller) GetTotalStakeWeights(opts *bind.CallOpts, operatorSet OperatorSet, referenceTimestamp uint32) ([]*big.Int, error) {
	var out []interface{}
	err := _BN254CertificateVerifier.contract.Call(opts, &out, "getTotalStakeWeights", operatorSet, referenceTimestamp)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetTotalStakeWeights is a free data retrieval call binding the contract method 0x7d1d1f5b.
//
// Solidity: function getTotalStakeWeights((address,uint32) operatorSet, uint32 referenceTimestamp) view returns(uint256[])
func (_BN254CertificateVerifier *BN254CertificateVerifierSession) GetTotalStakeWeights(operatorSet OperatorSet, referenceTimestamp uint32) ([]*big.Int, error) {
	return _BN254CertificateVerifier.Contract.GetTotalStakeWeights(&_BN254CertificateVerifier.CallOpts, operatorSet, referenceTimestamp)
}

// GetTotalStakeWeights is a free data retrieval call binding the contract method 0x7d1d1f5b.
//
// Solidity: function getTotalStakeWeights((address,uint32) operatorSet, uint32 referenceTimestamp) view returns(uint256[])
func (_BN254CertificateVerifier *BN254CertificateVerifierCallerSession) GetTotalStakeWeights(operatorSet OperatorSet, referenceTimestamp uint32) ([]*big.Int, error) {
	return _BN254CertificateVerifier.Contract.GetTotalStakeWeights(&_BN254CertificateVerifier.CallOpts, operatorSet, referenceTimestamp)
}

// IsNonsignerCached is a free data retrieval call binding the contract method 0x5be87274.
//
// Solidity: function isNonsignerCached((address,uint32) operatorSet, uint32 referenceTimestamp, uint256 operatorIndex) view returns(bool)
func (_BN254CertificateVerifier *BN254CertificateVerifierCaller) IsNonsignerCached(opts *bind.CallOpts, operatorSet OperatorSet, referenceTimestamp uint32, operatorIndex *big.Int) (bool, error) {
	var out []interface{}
	err := _BN254CertificateVerifier.contract.Call(opts, &out, "isNonsignerCached", operatorSet, referenceTimestamp, operatorIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsNonsignerCached is a free data retrieval call binding the contract method 0x5be87274.
//
// Solidity: function isNonsignerCached((address,uint32) operatorSet, uint32 referenceTimestamp, uint256 operatorIndex) view returns(bool)
func (_BN254CertificateVerifier *BN254CertificateVerifierSession) IsNonsignerCached(operatorSet OperatorSet, referenceTimestamp uint32, operatorIndex *big.Int) (bool, error) {
	return _BN254CertificateVerifier.Contract.IsNonsignerCached(&_BN254CertificateVerifier.CallOpts, operatorSet, referenceTimestamp, operatorIndex)
}

// IsNonsignerCached is a free data retrieval call binding the contract method 0x5be87274.
//
// Solidity: function isNonsignerCached((address,uint32) operatorSet, uint32 referenceTimestamp, uint256 operatorIndex) view returns(bool)
func (_BN254CertificateVerifier *BN254CertificateVerifierCallerSession) IsNonsignerCached(operatorSet OperatorSet, referenceTimestamp uint32, operatorIndex *big.Int) (bool, error) {
	return _BN254CertificateVerifier.Contract.IsNonsignerCached(&_BN254CertificateVerifier.CallOpts, operatorSet, referenceTimestamp, operatorIndex)
}

// IsReferenceTimestampSet is a free data retrieval call binding the contract method 0xcd83a72b.
//
// Solidity: function isReferenceTimestampSet((address,uint32) operatorSet, uint32 referenceTimestamp) view returns(bool)
func (_BN254CertificateVerifier *BN254CertificateVerifierCaller) IsReferenceTimestampSet(opts *bind.CallOpts, operatorSet OperatorSet, referenceTimestamp uint32) (bool, error) {
	var out []interface{}
	err := _BN254CertificateVerifier.contract.Call(opts, &out, "isReferenceTimestampSet", operatorSet, referenceTimestamp)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsReferenceTimestampSet is a free data retrieval call binding the contract method 0xcd83a72b.
//
// Solidity: function isReferenceTimestampSet((address,uint32) operatorSet, uint32 referenceTimestamp) view returns(bool)
func (_BN254CertificateVerifier *BN254CertificateVerifierSession) IsReferenceTimestampSet(operatorSet OperatorSet, referenceTimestamp uint32) (bool, error) {
	return _BN254CertificateVerifier.Contract.IsReferenceTimestampSet(&_BN254CertificateVerifier.CallOpts, operatorSet, referenceTimestamp)
}

// IsReferenceTimestampSet is a free data retrieval call binding the contract method 0xcd83a72b.
//
// Solidity: function isReferenceTimestampSet((address,uint32) operatorSet, uint32 referenceTimestamp) view returns(bool)
func (_BN254CertificateVerifier *BN254CertificateVerifierCallerSession) IsReferenceTimestampSet(operatorSet OperatorSet, referenceTimestamp uint32) (bool, error) {
	return _BN254CertificateVerifier.Contract.IsReferenceTimestampSet(&_BN254CertificateVerifier.CallOpts, operatorSet, referenceTimestamp)
}

// LatestReferenceTimestamp is a free data retrieval call binding the contract method 0x5ddb9b5b.
//
// Solidity: function latestReferenceTimestamp((address,uint32) operatorSet) view returns(uint32)
func (_BN254CertificateVerifier *BN254CertificateVerifierCaller) LatestReferenceTimestamp(opts *bind.CallOpts, operatorSet OperatorSet) (uint32, error) {
	var out []interface{}
	err := _BN254CertificateVerifier.contract.Call(opts, &out, "latestReferenceTimestamp", operatorSet)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LatestReferenceTimestamp is a free data retrieval call binding the contract method 0x5ddb9b5b.
//
// Solidity: function latestReferenceTimestamp((address,uint32) operatorSet) view returns(uint32)
func (_BN254CertificateVerifier *BN254CertificateVerifierSession) LatestReferenceTimestamp(operatorSet OperatorSet) (uint32, error) {
	return _BN254CertificateVerifier.Contract.LatestReferenceTimestamp(&_BN254CertificateVerifier.CallOpts, operatorSet)
}

// LatestReferenceTimestamp is a free data retrieval call binding the contract method 0x5ddb9b5b.
//
// Solidity: function latestReferenceTimestamp((address,uint32) operatorSet) view returns(uint32)
func (_BN254CertificateVerifier *BN254CertificateVerifierCallerSession) LatestReferenceTimestamp(operatorSet OperatorSet) (uint32, error) {
	return _BN254CertificateVerifier.Contract.LatestReferenceTimestamp(&_BN254CertificateVerifier.CallOpts, operatorSet)
}

// MaxOperatorTableStaleness is a free data retrieval call binding the contract method 0x6141879e.
//
// Solidity: function maxOperatorTableStaleness((address,uint32) operatorSet) view returns(uint32)
func (_BN254CertificateVerifier *BN254CertificateVerifierCaller) MaxOperatorTableStaleness(opts *bind.CallOpts, operatorSet OperatorSet) (uint32, error) {
	var out []interface{}
	err := _BN254CertificateVerifier.contract.Call(opts, &out, "maxOperatorTableStaleness", operatorSet)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MaxOperatorTableStaleness is a free data retrieval call binding the contract method 0x6141879e.
//
// Solidity: function maxOperatorTableStaleness((address,uint32) operatorSet) view returns(uint32)
func (_BN254CertificateVerifier *BN254CertificateVerifierSession) MaxOperatorTableStaleness(operatorSet OperatorSet) (uint32, error) {
	return _BN254CertificateVerifier.Contract.MaxOperatorTableStaleness(&_BN254CertificateVerifier.CallOpts, operatorSet)
}

// MaxOperatorTableStaleness is a free data retrieval call binding the contract method 0x6141879e.
//
// Solidity: function maxOperatorTableStaleness((address,uint32) operatorSet) view returns(uint32)
func (_BN254CertificateVerifier *BN254CertificateVerifierCallerSession) MaxOperatorTableStaleness(operatorSet OperatorSet) (uint32, error) {
	return _BN254CertificateVerifier.Contract.MaxOperatorTableStaleness(&_BN254CertificateVerifier.CallOpts, operatorSet)
}

// OperatorTableUpdater is a free data retrieval call binding the contract method 0x68d6e081.
//
// Solidity: function operatorTableUpdater() view returns(address)
func (_BN254CertificateVerifier *BN254CertificateVerifierCaller) OperatorTableUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BN254CertificateVerifier.contract.Call(opts, &out, "operatorTableUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OperatorTableUpdater is a free data retrieval call binding the contract method 0x68d6e081.
//
// Solidity: function operatorTableUpdater() view returns(address)
func (_BN254CertificateVerifier *BN254CertificateVerifierSession) OperatorTableUpdater() (common.Address, error) {
	return _BN254CertificateVerifier.Contract.OperatorTableUpdater(&_BN254CertificateVerifier.CallOpts)
}

// OperatorTableUpdater is a free data retrieval call binding the contract method 0x68d6e081.
//
// Solidity: function operatorTableUpdater() view returns(address)
func (_BN254CertificateVerifier *BN254CertificateVerifierCallerSession) OperatorTableUpdater() (common.Address, error) {
	return _BN254CertificateVerifier.Contract.OperatorTableUpdater(&_BN254CertificateVerifier.CallOpts)
}

// TrySignatureVerification is a free data retrieval call binding the contract method 0x1a18746c.
//
// Solidity: function trySignatureVerification(bytes32 msgHash, (uint256,uint256) aggPubkey, (uint256[2],uint256[2]) apkG2, (uint256,uint256) signature) view returns(bool pairingSuccessful, bool signatureValid)
func (_BN254CertificateVerifier *BN254CertificateVerifierCaller) TrySignatureVerification(opts *bind.CallOpts, msgHash [32]byte, aggPubkey BN254G1Point, apkG2 BN254G2Point, signature BN254G1Point) (struct {
	PairingSuccessful bool
	SignatureValid    bool
}, error) {
	var out []interface{}
	err := _BN254CertificateVerifier.contract.Call(opts, &out, "trySignatureVerification", msgHash, aggPubkey, apkG2, signature)

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
func (_BN254CertificateVerifier *BN254CertificateVerifierSession) TrySignatureVerification(msgHash [32]byte, aggPubkey BN254G1Point, apkG2 BN254G2Point, signature BN254G1Point) (struct {
	PairingSuccessful bool
	SignatureValid    bool
}, error) {
	return _BN254CertificateVerifier.Contract.TrySignatureVerification(&_BN254CertificateVerifier.CallOpts, msgHash, aggPubkey, apkG2, signature)
}

// TrySignatureVerification is a free data retrieval call binding the contract method 0x1a18746c.
//
// Solidity: function trySignatureVerification(bytes32 msgHash, (uint256,uint256) aggPubkey, (uint256[2],uint256[2]) apkG2, (uint256,uint256) signature) view returns(bool pairingSuccessful, bool signatureValid)
func (_BN254CertificateVerifier *BN254CertificateVerifierCallerSession) TrySignatureVerification(msgHash [32]byte, aggPubkey BN254G1Point, apkG2 BN254G2Point, signature BN254G1Point) (struct {
	PairingSuccessful bool
	SignatureValid    bool
}, error) {
	return _BN254CertificateVerifier.Contract.TrySignatureVerification(&_BN254CertificateVerifier.CallOpts, msgHash, aggPubkey, apkG2, signature)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_BN254CertificateVerifier *BN254CertificateVerifierCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BN254CertificateVerifier.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_BN254CertificateVerifier *BN254CertificateVerifierSession) Version() (string, error) {
	return _BN254CertificateVerifier.Contract.Version(&_BN254CertificateVerifier.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_BN254CertificateVerifier *BN254CertificateVerifierCallerSession) Version() (string, error) {
	return _BN254CertificateVerifier.Contract.Version(&_BN254CertificateVerifier.CallOpts)
}

// UpdateOperatorTable is a paid mutator transaction binding the contract method 0x6738c40b.
//
// Solidity: function updateOperatorTable((address,uint32) operatorSet, uint32 referenceTimestamp, (bytes32,uint256,(uint256,uint256),uint256[]) operatorSetInfo, (address,uint32) operatorSetConfig) returns()
func (_BN254CertificateVerifier *BN254CertificateVerifierTransactor) UpdateOperatorTable(opts *bind.TransactOpts, operatorSet OperatorSet, referenceTimestamp uint32, operatorSetInfo IOperatorTableCalculatorTypesBN254OperatorSetInfo, operatorSetConfig ICrossChainRegistryTypesOperatorSetConfig) (*types.Transaction, error) {
	return _BN254CertificateVerifier.contract.Transact(opts, "updateOperatorTable", operatorSet, referenceTimestamp, operatorSetInfo, operatorSetConfig)
}

// UpdateOperatorTable is a paid mutator transaction binding the contract method 0x6738c40b.
//
// Solidity: function updateOperatorTable((address,uint32) operatorSet, uint32 referenceTimestamp, (bytes32,uint256,(uint256,uint256),uint256[]) operatorSetInfo, (address,uint32) operatorSetConfig) returns()
func (_BN254CertificateVerifier *BN254CertificateVerifierSession) UpdateOperatorTable(operatorSet OperatorSet, referenceTimestamp uint32, operatorSetInfo IOperatorTableCalculatorTypesBN254OperatorSetInfo, operatorSetConfig ICrossChainRegistryTypesOperatorSetConfig) (*types.Transaction, error) {
	return _BN254CertificateVerifier.Contract.UpdateOperatorTable(&_BN254CertificateVerifier.TransactOpts, operatorSet, referenceTimestamp, operatorSetInfo, operatorSetConfig)
}

// UpdateOperatorTable is a paid mutator transaction binding the contract method 0x6738c40b.
//
// Solidity: function updateOperatorTable((address,uint32) operatorSet, uint32 referenceTimestamp, (bytes32,uint256,(uint256,uint256),uint256[]) operatorSetInfo, (address,uint32) operatorSetConfig) returns()
func (_BN254CertificateVerifier *BN254CertificateVerifierTransactorSession) UpdateOperatorTable(operatorSet OperatorSet, referenceTimestamp uint32, operatorSetInfo IOperatorTableCalculatorTypesBN254OperatorSetInfo, operatorSetConfig ICrossChainRegistryTypesOperatorSetConfig) (*types.Transaction, error) {
	return _BN254CertificateVerifier.Contract.UpdateOperatorTable(&_BN254CertificateVerifier.TransactOpts, operatorSet, referenceTimestamp, operatorSetInfo, operatorSetConfig)
}

// VerifyCertificate is a paid mutator transaction binding the contract method 0x080b7150.
//
// Solidity: function verifyCertificate((address,uint32) operatorSet, (uint32,bytes32,(uint256,uint256),(uint256[2],uint256[2]),(uint32,bytes,((uint256,uint256),uint256[]))[]) cert) returns(uint256[] totalSignedStakeWeights)
func (_BN254CertificateVerifier *BN254CertificateVerifierTransactor) VerifyCertificate(opts *bind.TransactOpts, operatorSet OperatorSet, cert IBN254CertificateVerifierTypesBN254Certificate) (*types.Transaction, error) {
	return _BN254CertificateVerifier.contract.Transact(opts, "verifyCertificate", operatorSet, cert)
}

// VerifyCertificate is a paid mutator transaction binding the contract method 0x080b7150.
//
// Solidity: function verifyCertificate((address,uint32) operatorSet, (uint32,bytes32,(uint256,uint256),(uint256[2],uint256[2]),(uint32,bytes,((uint256,uint256),uint256[]))[]) cert) returns(uint256[] totalSignedStakeWeights)
func (_BN254CertificateVerifier *BN254CertificateVerifierSession) VerifyCertificate(operatorSet OperatorSet, cert IBN254CertificateVerifierTypesBN254Certificate) (*types.Transaction, error) {
	return _BN254CertificateVerifier.Contract.VerifyCertificate(&_BN254CertificateVerifier.TransactOpts, operatorSet, cert)
}

// VerifyCertificate is a paid mutator transaction binding the contract method 0x080b7150.
//
// Solidity: function verifyCertificate((address,uint32) operatorSet, (uint32,bytes32,(uint256,uint256),(uint256[2],uint256[2]),(uint32,bytes,((uint256,uint256),uint256[]))[]) cert) returns(uint256[] totalSignedStakeWeights)
func (_BN254CertificateVerifier *BN254CertificateVerifierTransactorSession) VerifyCertificate(operatorSet OperatorSet, cert IBN254CertificateVerifierTypesBN254Certificate) (*types.Transaction, error) {
	return _BN254CertificateVerifier.Contract.VerifyCertificate(&_BN254CertificateVerifier.TransactOpts, operatorSet, cert)
}

// VerifyCertificateNominal is a paid mutator transaction binding the contract method 0xdd2ae1b9.
//
// Solidity: function verifyCertificateNominal((address,uint32) operatorSet, (uint32,bytes32,(uint256,uint256),(uint256[2],uint256[2]),(uint32,bytes,((uint256,uint256),uint256[]))[]) cert, uint256[] totalStakeNominalThresholds) returns(bool)
func (_BN254CertificateVerifier *BN254CertificateVerifierTransactor) VerifyCertificateNominal(opts *bind.TransactOpts, operatorSet OperatorSet, cert IBN254CertificateVerifierTypesBN254Certificate, totalStakeNominalThresholds []*big.Int) (*types.Transaction, error) {
	return _BN254CertificateVerifier.contract.Transact(opts, "verifyCertificateNominal", operatorSet, cert, totalStakeNominalThresholds)
}

// VerifyCertificateNominal is a paid mutator transaction binding the contract method 0xdd2ae1b9.
//
// Solidity: function verifyCertificateNominal((address,uint32) operatorSet, (uint32,bytes32,(uint256,uint256),(uint256[2],uint256[2]),(uint32,bytes,((uint256,uint256),uint256[]))[]) cert, uint256[] totalStakeNominalThresholds) returns(bool)
func (_BN254CertificateVerifier *BN254CertificateVerifierSession) VerifyCertificateNominal(operatorSet OperatorSet, cert IBN254CertificateVerifierTypesBN254Certificate, totalStakeNominalThresholds []*big.Int) (*types.Transaction, error) {
	return _BN254CertificateVerifier.Contract.VerifyCertificateNominal(&_BN254CertificateVerifier.TransactOpts, operatorSet, cert, totalStakeNominalThresholds)
}

// VerifyCertificateNominal is a paid mutator transaction binding the contract method 0xdd2ae1b9.
//
// Solidity: function verifyCertificateNominal((address,uint32) operatorSet, (uint32,bytes32,(uint256,uint256),(uint256[2],uint256[2]),(uint32,bytes,((uint256,uint256),uint256[]))[]) cert, uint256[] totalStakeNominalThresholds) returns(bool)
func (_BN254CertificateVerifier *BN254CertificateVerifierTransactorSession) VerifyCertificateNominal(operatorSet OperatorSet, cert IBN254CertificateVerifierTypesBN254Certificate, totalStakeNominalThresholds []*big.Int) (*types.Transaction, error) {
	return _BN254CertificateVerifier.Contract.VerifyCertificateNominal(&_BN254CertificateVerifier.TransactOpts, operatorSet, cert, totalStakeNominalThresholds)
}

// VerifyCertificateProportion is a paid mutator transaction binding the contract method 0x017d7974.
//
// Solidity: function verifyCertificateProportion((address,uint32) operatorSet, (uint32,bytes32,(uint256,uint256),(uint256[2],uint256[2]),(uint32,bytes,((uint256,uint256),uint256[]))[]) cert, uint16[] totalStakeProportionThresholds) returns(bool)
func (_BN254CertificateVerifier *BN254CertificateVerifierTransactor) VerifyCertificateProportion(opts *bind.TransactOpts, operatorSet OperatorSet, cert IBN254CertificateVerifierTypesBN254Certificate, totalStakeProportionThresholds []uint16) (*types.Transaction, error) {
	return _BN254CertificateVerifier.contract.Transact(opts, "verifyCertificateProportion", operatorSet, cert, totalStakeProportionThresholds)
}

// VerifyCertificateProportion is a paid mutator transaction binding the contract method 0x017d7974.
//
// Solidity: function verifyCertificateProportion((address,uint32) operatorSet, (uint32,bytes32,(uint256,uint256),(uint256[2],uint256[2]),(uint32,bytes,((uint256,uint256),uint256[]))[]) cert, uint16[] totalStakeProportionThresholds) returns(bool)
func (_BN254CertificateVerifier *BN254CertificateVerifierSession) VerifyCertificateProportion(operatorSet OperatorSet, cert IBN254CertificateVerifierTypesBN254Certificate, totalStakeProportionThresholds []uint16) (*types.Transaction, error) {
	return _BN254CertificateVerifier.Contract.VerifyCertificateProportion(&_BN254CertificateVerifier.TransactOpts, operatorSet, cert, totalStakeProportionThresholds)
}

// VerifyCertificateProportion is a paid mutator transaction binding the contract method 0x017d7974.
//
// Solidity: function verifyCertificateProportion((address,uint32) operatorSet, (uint32,bytes32,(uint256,uint256),(uint256[2],uint256[2]),(uint32,bytes,((uint256,uint256),uint256[]))[]) cert, uint16[] totalStakeProportionThresholds) returns(bool)
func (_BN254CertificateVerifier *BN254CertificateVerifierTransactorSession) VerifyCertificateProportion(operatorSet OperatorSet, cert IBN254CertificateVerifierTypesBN254Certificate, totalStakeProportionThresholds []uint16) (*types.Transaction, error) {
	return _BN254CertificateVerifier.Contract.VerifyCertificateProportion(&_BN254CertificateVerifier.TransactOpts, operatorSet, cert, totalStakeProportionThresholds)
}

// BN254CertificateVerifierInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the BN254CertificateVerifier contract.
type BN254CertificateVerifierInitializedIterator struct {
	Event *BN254CertificateVerifierInitialized // Event containing the contract specifics and raw log

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
func (it *BN254CertificateVerifierInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BN254CertificateVerifierInitialized)
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
		it.Event = new(BN254CertificateVerifierInitialized)
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
func (it *BN254CertificateVerifierInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BN254CertificateVerifierInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BN254CertificateVerifierInitialized represents a Initialized event raised by the BN254CertificateVerifier contract.
type BN254CertificateVerifierInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BN254CertificateVerifier *BN254CertificateVerifierFilterer) FilterInitialized(opts *bind.FilterOpts) (*BN254CertificateVerifierInitializedIterator, error) {

	logs, sub, err := _BN254CertificateVerifier.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BN254CertificateVerifierInitializedIterator{contract: _BN254CertificateVerifier.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BN254CertificateVerifier *BN254CertificateVerifierFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BN254CertificateVerifierInitialized) (event.Subscription, error) {

	logs, sub, err := _BN254CertificateVerifier.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BN254CertificateVerifierInitialized)
				if err := _BN254CertificateVerifier.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BN254CertificateVerifier *BN254CertificateVerifierFilterer) ParseInitialized(log types.Log) (*BN254CertificateVerifierInitialized, error) {
	event := new(BN254CertificateVerifierInitialized)
	if err := _BN254CertificateVerifier.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BN254CertificateVerifierMaxStalenessPeriodUpdatedIterator is returned from FilterMaxStalenessPeriodUpdated and is used to iterate over the raw logs and unpacked data for MaxStalenessPeriodUpdated events raised by the BN254CertificateVerifier contract.
type BN254CertificateVerifierMaxStalenessPeriodUpdatedIterator struct {
	Event *BN254CertificateVerifierMaxStalenessPeriodUpdated // Event containing the contract specifics and raw log

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
func (it *BN254CertificateVerifierMaxStalenessPeriodUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BN254CertificateVerifierMaxStalenessPeriodUpdated)
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
		it.Event = new(BN254CertificateVerifierMaxStalenessPeriodUpdated)
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
func (it *BN254CertificateVerifierMaxStalenessPeriodUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BN254CertificateVerifierMaxStalenessPeriodUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BN254CertificateVerifierMaxStalenessPeriodUpdated represents a MaxStalenessPeriodUpdated event raised by the BN254CertificateVerifier contract.
type BN254CertificateVerifierMaxStalenessPeriodUpdated struct {
	OperatorSet        OperatorSet
	MaxStalenessPeriod uint32
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMaxStalenessPeriodUpdated is a free log retrieval operation binding the contract event 0x28539469fbbc8a5482e60966bf9376f7b9d25b2f0a65a9976f6baa3f0e3788da.
//
// Solidity: event MaxStalenessPeriodUpdated((address,uint32) operatorSet, uint32 maxStalenessPeriod)
func (_BN254CertificateVerifier *BN254CertificateVerifierFilterer) FilterMaxStalenessPeriodUpdated(opts *bind.FilterOpts) (*BN254CertificateVerifierMaxStalenessPeriodUpdatedIterator, error) {

	logs, sub, err := _BN254CertificateVerifier.contract.FilterLogs(opts, "MaxStalenessPeriodUpdated")
	if err != nil {
		return nil, err
	}
	return &BN254CertificateVerifierMaxStalenessPeriodUpdatedIterator{contract: _BN254CertificateVerifier.contract, event: "MaxStalenessPeriodUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxStalenessPeriodUpdated is a free log subscription operation binding the contract event 0x28539469fbbc8a5482e60966bf9376f7b9d25b2f0a65a9976f6baa3f0e3788da.
//
// Solidity: event MaxStalenessPeriodUpdated((address,uint32) operatorSet, uint32 maxStalenessPeriod)
func (_BN254CertificateVerifier *BN254CertificateVerifierFilterer) WatchMaxStalenessPeriodUpdated(opts *bind.WatchOpts, sink chan<- *BN254CertificateVerifierMaxStalenessPeriodUpdated) (event.Subscription, error) {

	logs, sub, err := _BN254CertificateVerifier.contract.WatchLogs(opts, "MaxStalenessPeriodUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BN254CertificateVerifierMaxStalenessPeriodUpdated)
				if err := _BN254CertificateVerifier.contract.UnpackLog(event, "MaxStalenessPeriodUpdated", log); err != nil {
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
func (_BN254CertificateVerifier *BN254CertificateVerifierFilterer) ParseMaxStalenessPeriodUpdated(log types.Log) (*BN254CertificateVerifierMaxStalenessPeriodUpdated, error) {
	event := new(BN254CertificateVerifierMaxStalenessPeriodUpdated)
	if err := _BN254CertificateVerifier.contract.UnpackLog(event, "MaxStalenessPeriodUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BN254CertificateVerifierOperatorSetOwnerUpdatedIterator is returned from FilterOperatorSetOwnerUpdated and is used to iterate over the raw logs and unpacked data for OperatorSetOwnerUpdated events raised by the BN254CertificateVerifier contract.
type BN254CertificateVerifierOperatorSetOwnerUpdatedIterator struct {
	Event *BN254CertificateVerifierOperatorSetOwnerUpdated // Event containing the contract specifics and raw log

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
func (it *BN254CertificateVerifierOperatorSetOwnerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BN254CertificateVerifierOperatorSetOwnerUpdated)
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
		it.Event = new(BN254CertificateVerifierOperatorSetOwnerUpdated)
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
func (it *BN254CertificateVerifierOperatorSetOwnerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BN254CertificateVerifierOperatorSetOwnerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BN254CertificateVerifierOperatorSetOwnerUpdated represents a OperatorSetOwnerUpdated event raised by the BN254CertificateVerifier contract.
type BN254CertificateVerifierOperatorSetOwnerUpdated struct {
	OperatorSet OperatorSet
	Owner       common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorSetOwnerUpdated is a free log retrieval operation binding the contract event 0x806dc367095c0baf953d7144b7c4376261675ee0b4e0da2761e43673051c7375.
//
// Solidity: event OperatorSetOwnerUpdated((address,uint32) operatorSet, address owner)
func (_BN254CertificateVerifier *BN254CertificateVerifierFilterer) FilterOperatorSetOwnerUpdated(opts *bind.FilterOpts) (*BN254CertificateVerifierOperatorSetOwnerUpdatedIterator, error) {

	logs, sub, err := _BN254CertificateVerifier.contract.FilterLogs(opts, "OperatorSetOwnerUpdated")
	if err != nil {
		return nil, err
	}
	return &BN254CertificateVerifierOperatorSetOwnerUpdatedIterator{contract: _BN254CertificateVerifier.contract, event: "OperatorSetOwnerUpdated", logs: logs, sub: sub}, nil
}

// WatchOperatorSetOwnerUpdated is a free log subscription operation binding the contract event 0x806dc367095c0baf953d7144b7c4376261675ee0b4e0da2761e43673051c7375.
//
// Solidity: event OperatorSetOwnerUpdated((address,uint32) operatorSet, address owner)
func (_BN254CertificateVerifier *BN254CertificateVerifierFilterer) WatchOperatorSetOwnerUpdated(opts *bind.WatchOpts, sink chan<- *BN254CertificateVerifierOperatorSetOwnerUpdated) (event.Subscription, error) {

	logs, sub, err := _BN254CertificateVerifier.contract.WatchLogs(opts, "OperatorSetOwnerUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BN254CertificateVerifierOperatorSetOwnerUpdated)
				if err := _BN254CertificateVerifier.contract.UnpackLog(event, "OperatorSetOwnerUpdated", log); err != nil {
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
func (_BN254CertificateVerifier *BN254CertificateVerifierFilterer) ParseOperatorSetOwnerUpdated(log types.Log) (*BN254CertificateVerifierOperatorSetOwnerUpdated, error) {
	event := new(BN254CertificateVerifierOperatorSetOwnerUpdated)
	if err := _BN254CertificateVerifier.contract.UnpackLog(event, "OperatorSetOwnerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BN254CertificateVerifierTableUpdatedIterator is returned from FilterTableUpdated and is used to iterate over the raw logs and unpacked data for TableUpdated events raised by the BN254CertificateVerifier contract.
type BN254CertificateVerifierTableUpdatedIterator struct {
	Event *BN254CertificateVerifierTableUpdated // Event containing the contract specifics and raw log

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
func (it *BN254CertificateVerifierTableUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BN254CertificateVerifierTableUpdated)
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
		it.Event = new(BN254CertificateVerifierTableUpdated)
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
func (it *BN254CertificateVerifierTableUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BN254CertificateVerifierTableUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BN254CertificateVerifierTableUpdated represents a TableUpdated event raised by the BN254CertificateVerifier contract.
type BN254CertificateVerifierTableUpdated struct {
	OperatorSet        OperatorSet
	ReferenceTimestamp uint32
	OperatorSetInfo    IOperatorTableCalculatorTypesBN254OperatorSetInfo
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterTableUpdated is a free log retrieval operation binding the contract event 0x93e6bea1c9b5dce4a5c07b00261e956df2a4a253d9ab6ca070ca2037d72ada9e.
//
// Solidity: event TableUpdated((address,uint32) operatorSet, uint32 referenceTimestamp, (bytes32,uint256,(uint256,uint256),uint256[]) operatorSetInfo)
func (_BN254CertificateVerifier *BN254CertificateVerifierFilterer) FilterTableUpdated(opts *bind.FilterOpts) (*BN254CertificateVerifierTableUpdatedIterator, error) {

	logs, sub, err := _BN254CertificateVerifier.contract.FilterLogs(opts, "TableUpdated")
	if err != nil {
		return nil, err
	}
	return &BN254CertificateVerifierTableUpdatedIterator{contract: _BN254CertificateVerifier.contract, event: "TableUpdated", logs: logs, sub: sub}, nil
}

// WatchTableUpdated is a free log subscription operation binding the contract event 0x93e6bea1c9b5dce4a5c07b00261e956df2a4a253d9ab6ca070ca2037d72ada9e.
//
// Solidity: event TableUpdated((address,uint32) operatorSet, uint32 referenceTimestamp, (bytes32,uint256,(uint256,uint256),uint256[]) operatorSetInfo)
func (_BN254CertificateVerifier *BN254CertificateVerifierFilterer) WatchTableUpdated(opts *bind.WatchOpts, sink chan<- *BN254CertificateVerifierTableUpdated) (event.Subscription, error) {

	logs, sub, err := _BN254CertificateVerifier.contract.WatchLogs(opts, "TableUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BN254CertificateVerifierTableUpdated)
				if err := _BN254CertificateVerifier.contract.UnpackLog(event, "TableUpdated", log); err != nil {
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
func (_BN254CertificateVerifier *BN254CertificateVerifierFilterer) ParseTableUpdated(log types.Log) (*BN254CertificateVerifierTableUpdated, error) {
	event := new(BN254CertificateVerifierTableUpdated)
	if err := _BN254CertificateVerifier.contract.UnpackLog(event, "TableUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
