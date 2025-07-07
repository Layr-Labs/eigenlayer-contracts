// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package OperatorTableUpdater

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

// OperatorTableUpdaterMetaData contains all meta data concerning the OperatorTableUpdater contract.
var OperatorTableUpdaterMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_bn254CertificateVerifier\",\"type\":\"address\",\"internalType\":\"contractIBN254CertificateVerifier\"},{\"name\":\"_ecdsaCertificateVerifier\",\"type\":\"address\",\"internalType\":\"contractIECDSACertificateVerifier\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"_version\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"GLOBAL_TABLE_ROOT_CERT_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_BPS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bn254CertificateVerifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBN254CertificateVerifier\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"confirmGlobalTableRoot\",\"inputs\":[{\"name\":\"globalTableRootCert\",\"type\":\"tuple\",\"internalType\":\"structIBN254CertificateVerifierTypes.BN254Certificate\",\"components\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"nonSignerWitnesses\",\"type\":\"tuple[]\",\"internalType\":\"structIBN254CertificateVerifierTypes.BN254OperatorInfoWitness[]\",\"components\":[{\"name\":\"operatorIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorInfoProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"operatorInfo\",\"type\":\"tuple\",\"internalType\":\"structIOperatorTableCalculatorTypes.BN254OperatorInfo\",\"components\":[{\"name\":\"pubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}]}]},{\"name\":\"globalTableRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"disableRoot\",\"inputs\":[{\"name\":\"globalTableRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ecdsaCertificateVerifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIECDSACertificateVerifier\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCertificateVerifier\",\"inputs\":[{\"name\":\"curveType\",\"type\":\"uint8\",\"internalType\":\"enumIKeyRegistrarTypes.CurveType\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentGlobalTableRoot\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getGenerator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getGeneratorReferenceTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getGlobalTableRootByTimestamp\",\"inputs\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getGlobalTableUpdateMessageHash\",\"inputs\":[{\"name\":\"globalTableRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getLatestReferenceBlockNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLatestReferenceTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getReferenceBlockNumberByTimestamp\",\"inputs\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getReferenceTimestampByBlockNumber\",\"inputs\":[{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"globalRootConfirmationThreshold\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"generator\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"_globalRootConfirmationThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"initialGlobalTableRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"generatorInfo\",\"type\":\"tuple\",\"internalType\":\"structIOperatorTableCalculatorTypes.BN254OperatorSetInfo\",\"components\":[{\"name\":\"operatorInfoTreeRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"numOperators\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"aggregatePubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"totalWeights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]},{\"name\":\"generatorConfig\",\"type\":\"tuple\",\"internalType\":\"structICrossChainRegistryTypes.OperatorSetConfig\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxStalenessPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isRootValid\",\"inputs\":[{\"name\":\"globalTableRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRootValidByTimestamp\",\"inputs\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGlobalRootConfirmationThreshold\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOperatorTable\",\"inputs\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"globalTableRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"operatorSetIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"operatorTableBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"GlobalRootConfirmationThresholdUpdated\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GlobalRootDisabled\",\"inputs\":[{\"name\":\"globalTableRoot\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewGlobalTableRoot\",\"inputs\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"globalTableRoot\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CertificateInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CurrentlyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"GlobalTableRootInFuture\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"GlobalTableRootStale\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidConfirmationThreshold\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCurveType\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidGlobalTableRoot\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidMessageHash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNewPausedStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperatorSetProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProofLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRoot\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignatureLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyPauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnpauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"TableUpdateForPastTimestamp\",\"inputs\":[]}]",
	Bin: "0x610100604052348015610010575f5ffd5b5060405161268338038061268383398101604081905261002f916101b9565b808484846001600160a01b03811661005a576040516339b190bb60e11b815260040160405180910390fd5b6001600160a01b0390811660805291821660a0521660c05261007b81610090565b60e052506100876100d6565b505050506102fe565b5f5f829050601f815111156100c3578260405163305a27a960e01b81526004016100ba91906102a3565b60405180910390fd5b80516100ce826102d8565b179392505050565b5f54610100900460ff161561013d5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b60648201526084016100ba565b5f5460ff9081161461018c575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146101a2575f5ffd5b50565b634e487b7160e01b5f52604160045260245ffd5b5f5f5f5f608085870312156101cc575f5ffd5b84516101d78161018e565b60208601519094506101e88161018e565b60408601519093506101f98161018e565b60608601519092506001600160401b03811115610214575f5ffd5b8501601f81018713610224575f5ffd5b80516001600160401b0381111561023d5761023d6101a5565b604051601f8201601f19908116603f011681016001600160401b038111828210171561026b5761026b6101a5565b604052818152828201602001891015610282575f5ffd5b8160208401602083015e5f6020838301015280935050505092959194509250565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b805160208083015191908110156102f8575f198160200360031b1b821691505b50919050565b60805160a05160c05160e0516123136103705f395f6105ea01525f818161049e015281816106840152610b2e01525f81816104c501528181610644015281816106f001528181610a8901528181610d6801526111e101525f818161044001528181610fa6015261139c01526123135ff3fe608060405234801561000f575f5ffd5b50600436106101e7575f3560e01c8063715018a611610109578063c252aa221161009e578063eaaed9d51161006e578063eaaed9d514610553578063f2fde38b14610566578063fabc1cbc14610579578063fd967f471461058c575f5ffd5b8063c252aa22146104e7578063c3621f0a14610508578063c3be1e331461051b578063c5916a391461052e575f5ffd5b80638f797718116100d95780638f797718146104735780639ea9477814610486578063ad0f958214610499578063b8c14306146104c0575f5ffd5b8063715018a61461042b5780637551ba3414610433578063886f11951461043b5780638da5cb5b14610462575f5ffd5b80633ef6cd7a1161017f5780635ac86ab71161014f5780635ac86ab7146103a05780635c975abb146103c357806364e1df84146103cb5780636f728c5014610400575f5ffd5b80633ef6cd7a146103485780634624e6a31461036f57806354fd4d5014610383578063595c6a6714610398575f5ffd5b806323b7b5b2116101ba57806323b7b5b21461029d57806328522d79146102c557806330ef41b4146102f157806331a599d214610323575f5ffd5b8063136439dd146101eb578063193b79f3146102005780631e2ca260146102425780632370356c1461028a575b5f5ffd5b6101fe6101f936600461155f565b610595565b005b61022861020e366004611587565b63ffffffff9081165f908152609b60205260409020541690565b60405163ffffffff90911681526020015b60405180910390f35b6040805180820182525f80825260209182015281518083019092526098546001600160a01b0381168352600160a01b900463ffffffff169082015260405161023991906115c7565b6101fe6102983660046115e6565b6105cf565b6102286102ab366004611587565b63ffffffff9081165f908152609a60205260409020541690565b60975462010000900463ffffffff165f908152609960205260409020545b604051908152602001610239565b6103136102ff36600461155f565b5f908152609c602052604090205460ff1690565b6040519015158152602001610239565b60975462010000900463ffffffff9081165f908152609a602052604090205416610228565b6102e37f4491f5ee91595f938885ef73c9a1fa8a6d14ff9b9dab4aa24b8802bbb9bfc1cc81565b60975462010000900463ffffffff16610228565b61038b6105e3565b60405161023991906115ff565b6101fe610613565b6103136103ae366004611634565b606654600160ff9092169190911b9081161490565b6066546102e3565b6103136103d9366004611587565b63ffffffff165f908152609960209081526040808320548352609c90915290205460ff1690565b61041361040e366004611662565b610627565b6040516001600160a01b039091168152602001610239565b6101fe6106c6565b6102286106d7565b6104137f000000000000000000000000000000000000000000000000000000000000000081565b6033546001600160a01b0316610413565b6101fe6104813660046116a5565b610765565b6101fe610494366004611791565b61095a565b6104137f000000000000000000000000000000000000000000000000000000000000000081565b6104137f000000000000000000000000000000000000000000000000000000000000000081565b6097546104f59061ffff1681565b60405161ffff9091168152602001610239565b6101fe61051636600461155f565b610b94565b6102e361052936600461182c565b610c09565b6102e361053c366004611587565b63ffffffff165f9081526099602052604090205490565b6101fe61056136600461186b565b610c71565b6101fe6105743660046118d7565b610eae565b6101fe61058736600461155f565b610f24565b6104f561271081565b61059d610f91565b60665481811681146105c25760405163c61dca5d60e01b815260040160405180910390fd5b6105cb82611034565b5050565b6105d7611071565b6105e0816110cb565b50565b606061060e7f000000000000000000000000000000000000000000000000000000000000000061113c565b905090565b61061b610f91565b6106255f19611034565b565b5f600282600281111561063c5761063c6118f2565b0361066857507f0000000000000000000000000000000000000000000000000000000000000000919050565b600182600281111561067c5761067c6118f2565b036106a857507f0000000000000000000000000000000000000000000000000000000000000000919050565b60405163fdea7c0960e01b815260040160405180910390fd5b919050565b6106ce611071565b6106255f611179565b604051635ddb9b5b60e01b81525f906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690635ddb9b5b9061072690609890600401611906565b602060405180830381865afa158015610741573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061060e919061192d565b5f54610100900460ff161580801561078357505f54600160ff909116105b8061079c5750303b15801561079c57505f5460ff166001145b6108045760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b5f805460ff191660011790558015610825575f805461ff0019166101001790555b61082e89611179565b61083788611034565b8660986108448282611948565b905050610850866110cb565b63ffffffff8085165f8181526099602090815260408083208a9055898352609c8252808320805460ff19166001179055838352609a825280832080544390961663ffffffff1996871681179091558352609b909152902080549092168117909155609780546201000090920265ffffffff0000199092169190911790556108d88484846111ca565b604051859063ffffffff8616907f010dcbe0d1e019c93357711f7bb6287d543b7ff7de74f29df3fb5ecceec8d369905f90a3801561094f575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050505050505050565b60016109658161124f565b5f5f5f5f610973878761127a565b5f8f8152609c60205260409020549397509195509350915060ff166109ab5760405163504570e360e01b815260040160405180910390fd5b6109b483610627565b6001600160a01b0316635ddb9b5b856040518263ffffffff1660e01b81526004016109df91906115c7565b602060405180830381865afa1580156109fa573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a1e919061192d565b63ffffffff168c63ffffffff1611610a495760405163207617df60e01b815260040160405180910390fd5b610a6e8c8c8c8c8c8c8c604051610a619291906119a4565b60405180910390206112c1565b6002836002811115610a8257610a826118f2565b03610b13577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316636738c40b858e610ac185611362565b866040518563ffffffff1660e01b8152600401610ae194939291906119ed565b5f604051808303815f87803b158015610af8575f5ffd5b505af1158015610b0a573d5f5f3e3d5ffd5b50505050610b86565b6001836002811115610b2757610b276118f2565b036106a8577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166356d482f5858e610b6685611384565b866040518563ffffffff1660e01b8152600401610ae19493929190611a69565b505050505050505050505050565b610b9c610f91565b5f818152609c602052604090205460ff16610bca5760405163504570e360e01b815260040160405180910390fd5b5f818152609c6020526040808220805460ff191690555182917f8bd43de1250f58fe6ec9a78671a8b78dba70f0018656d157a3aeaabec389df3491a250565b604080517f4491f5ee91595f938885ef73c9a1fa8a6d14ff9b9dab4aa24b8802bbb9bfc1cc602082015290810184905263ffffffff8084166060830152821660808201525f9060a0016040516020818303038152906040528051906020012090509392505050565b5f610c7b8161124f565b428363ffffffff161115610ca257604051635a119db560e11b815260040160405180910390fd5b60975463ffffffff62010000909104811690841611610cd45760405163037fa86b60e31b815260040160405180910390fd5b610cdf848484610c09565b856020013514610d0257604051638b56642d60e01b815260040160405180910390fd5b6040805160018082528183019092525f91602080830190803683375050609754825192935061ffff16918391505f90610d3d57610d3d611b22565b61ffff90921660209283029190910190910152604051625f5e5d60e21b81525f906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063017d797490610da2906098908b908790600401611c54565b6020604051808303815f875af1158015610dbe573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610de29190611de1565b905080610e0257604051633042041f60e21b815260040160405180910390fd5b6097805463ffffffff80881662010000810265ffffffff000019909316929092179092555f818152609a602090815260408083208054958a1663ffffffff1996871681179091558352609b825280832080549095168417909455828252609981528382208a9055898252609c9052828120805460ff19166001179055915188927f010dcbe0d1e019c93357711f7bb6287d543b7ff7de74f29df3fb5ecceec8d36991a350505050505050565b610eb6611071565b6001600160a01b038116610f1b5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016107fb565b6105e081611179565b610f2c61139a565b60665480198219811614610f535760405163c61dca5d60e01b815260040160405180910390fd5b606682905560405182815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c9060200160405180910390a25050565b60405163237dfb4760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906346fbf68e90602401602060405180830381865afa158015610ff3573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906110179190611de1565b61062557604051631d77d47760e21b815260040160405180910390fd5b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a250565b6033546001600160a01b031633146106255760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016107fb565b61271061ffff821611156110f2576040516307336f0360e11b815260040160405180910390fd5b6097805461ffff191661ffff83169081179091556040519081527ff5d1836df8fcd7c1e54047e94ac8773d2855395603e2ef9ba5f5f16905f225929060200160405180910390a150565b60605f6111488361144b565b6040805160208082528183019092529192505f91906020820181803683375050509182525060208101929092525090565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b604051636738c40b60e01b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690636738c40b9061121d90609890879087908790600401611e00565b5f604051808303815f87803b158015611234575f5ffd5b505af1158015611246573d5f5f3e3d5ffd5b50505050505050565b606654600160ff83161b908116036105e05760405163840a48d560e01b815260040160405180910390fd5b604080518082019091525f8082526020820152604080518082019091525f808252602082018190529060606112b185870187611f83565b9299919850965090945092505050565b63ffffffff86165f9081526099602052604090205485146112f55760405163639d09b560e11b815260040160405180910390fd5b61133d83838080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152508992508591505063ffffffff8816611472565b61135a5760405163afa42ca760e01b815260040160405180910390fd5b505050505050565b61136a611520565b8180602001905181019061137e91906120d4565b92915050565b60608180602001905181019061137e9190612183565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156113f6573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061141a9190612284565b6001600160a01b0316336001600160a01b0316146106255760405163794821ff60e01b815260040160405180910390fd5b5f60ff8216601f81111561137e57604051632cd44ac360e21b815260040160405180910390fd5b5f8361147f868585611489565b1495945050505050565b5f60208451611498919061229f565b156114b6576040516313717da960e21b815260040160405180910390fd5b8260205b85518111611517576114cd60028561229f565b5f036114ee57815f528086015160205260405f209150600284049350611505565b808601515f528160205260405f2091506002840493505b6115106020826122be565b90506114ba565b50949350505050565b60405180608001604052805f81526020015f815260200161155260405180604001604052805f81526020015f81525090565b8152602001606081525090565b5f6020828403121561156f575f5ffd5b5035919050565b63ffffffff811681146105e0575f5ffd5b5f60208284031215611597575f5ffd5b81356115a281611576565b9392505050565b80516001600160a01b0316825260209081015163ffffffff16910152565b6040810161137e82846115a9565b803561ffff811681146106c1575f5ffd5b5f602082840312156115f6575f5ffd5b6115a2826115d5565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b5f60208284031215611644575f5ffd5b813560ff811681146115a2575f5ffd5b8035600381106106c1575f5ffd5b5f60208284031215611672575f5ffd5b6115a282611654565b6001600160a01b03811681146105e0575f5ffd5b5f6040828403121561169f575f5ffd5b50919050565b5f5f5f5f5f5f5f5f610140898b0312156116bd575f5ffd5b88356116c88161167b565b9750602089013596506116de8a60408b0161168f565b95506116ec60808a016115d5565b945060a0890135935060c089013561170381611576565b925060e08901356001600160401b0381111561171d575f5ffd5b890160a0818c03121561172e575f5ffd5b915061173e8a6101008b0161168f565b90509295985092959890939650565b5f5f83601f84011261175d575f5ffd5b5081356001600160401b03811115611773575f5ffd5b60208301915083602082850101111561178a575f5ffd5b9250929050565b5f5f5f5f5f5f5f60a0888a0312156117a7575f5ffd5b87356117b281611576565b96506020880135955060408801356117c981611576565b945060608801356001600160401b038111156117e3575f5ffd5b6117ef8a828b0161174d565b90955093505060808801356001600160401b0381111561180d575f5ffd5b6118198a828b0161174d565b989b979a50959850939692959293505050565b5f5f5f6060848603121561183e575f5ffd5b83359250602084013561185081611576565b9150604084013561186081611576565b809150509250925092565b5f5f5f5f6080858703121561187e575f5ffd5b84356001600160401b03811115611893575f5ffd5b850161012081880312156118a5575f5ffd5b93506020850135925060408501356118bc81611576565b915060608501356118cc81611576565b939692955090935050565b5f602082840312156118e7575f5ffd5b81356115a28161167b565b634e487b7160e01b5f52602160045260245ffd5b6040810161137e8284546001600160a01b038116825260a01c63ffffffff16602090910152565b5f6020828403121561193d575f5ffd5b81516115a281611576565b81356119538161167b565b81546001600160a01b031981166001600160a01b03929092169182178355602084013561197f81611576565b6001600160c01b03199190911690911760a09190911b63ffffffff60a01b1617905550565b818382375f9101908152919050565b5f8151808452602084019350602083015f5b828110156119e35781518652602095860195909101906001016119c5565b5093949350505050565b6119f781866115a9565b63ffffffff8416604082015260c06060820152825160c0820152602083015160e08201525f60408401518051610100840152602081015161012084015250606084015160a0610140840152611a506101608401826119b3565b915050611a6060808301846115a9565b95945050505050565b5f60c08201611a7883886115a9565b63ffffffff8616604084015260c0606084015280855180835260e08501915060e08160051b8601019250602087015f5b82811015611af95786850360df19018452815180516001600160a01b03168652602090810151604091870182905290611ae3908701826119b3565b9550506020938401939190910190600101611aa8565b5050505080915050611a6060808301846115a9565b634e487b7160e01b5f52604160045260245ffd5b634e487b7160e01b5f52603260045260245ffd5b5f5f8335601e19843603018112611b4b575f5ffd5b83016020810192503590506001600160401b03811115611b69575f5ffd5b8060051b360382131561178a575f5ffd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b5f8235605e19833603018112611bb6575f5ffd5b90910192915050565b8183525f6001600160fb1b03831115611bd6575f5ffd5b8260051b80836020870137939093016020019392505050565b80358252602080820135908301525f611c0b6040830183611b36565b60606040860152611a60606086018284611bbf565b5f8151808452602084019350602083015f5b828110156119e357815161ffff16865260209586019590910190600101611c32565b611c778185546001600160a01b038116825260a01c63ffffffff16602090910152565b608060408201525f6101a082018435611c8f81611576565b63ffffffff166080840152602085013560a0840152604085013560c0840152606085013560e0840152604060808601610100850137604060c08601610140850137611cde610100860186611b36565b610120610180860152828184526101c0860190506101c08260051b8701019350825f5b83811015611dc0578786036101bf19018352611d1d8286611ba2565b8035611d2881611576565b63ffffffff168752602081013536829003601e19018112611d47575f5ffd5b81016020810190356001600160401b03811115611d62575f5ffd5b803603821315611d70575f5ffd5b606060208a0152611d8560608a018284611b7a565b915050611d956040830183611ba2565b91508781036040890152611da98183611bef565b975050506020928301929190910190600101611d01565b50505050508281036060840152611dd78185611c20565b9695505050505050565b5f60208284031215611df1575f5ffd5b815180151581146115a2575f5ffd5b611e238186546001600160a01b038116825260a01c63ffffffff16602090910152565b63ffffffff841660408281019190915260c06060808401829052853591840191909152602085013560e0840152908401356101008301528301356101208201525f611e716080850185611b36565b60a0610140850152611e8861016085018284611bbf565b925050508235611e978161167b565b6001600160a01b031660808301526020830135611eb381611576565b63ffffffff811660a08401525095945050505050565b604080519081016001600160401b0381118282101715611eeb57611eeb611b0e565b60405290565b604051608081016001600160401b0381118282101715611eeb57611eeb611b0e565b604051601f8201601f191681016001600160401b0381118282101715611f3b57611f3b611b0e565b604052919050565b5f60408284031215611f53575f5ffd5b611f5b611ec9565b90508135611f688161167b565b81526020820135611f7881611576565b602082015292915050565b5f5f5f5f60c08587031215611f96575f5ffd5b611fa08686611f43565b9350611fae60408601611654565b9250611fbd8660608701611f43565b915060a08501356001600160401b03811115611fd7575f5ffd5b8501601f81018713611fe7575f5ffd5b80356001600160401b0381111561200057612000611b0e565b612013601f8201601f1916602001611f13565b818152886020838501011115612027575f5ffd5b816020840160208301375f6020838301015280935050505092959194509250565b5f6001600160401b0382111561206057612060611b0e565b5060051b60200190565b5f82601f830112612079575f5ffd5b815161208c61208782612048565b611f13565b8082825260208201915060208360051b8601019250858311156120ad575f5ffd5b602085015b838110156120ca5780518352602092830192016120b2565b5095945050505050565b5f602082840312156120e4575f5ffd5b81516001600160401b038111156120f9575f5ffd5b820180840360a081121561210b575f5ffd5b612113611ef1565b82518152602080840151908201526040603f1983011215612132575f5ffd5b61213a611ec9565b604084810151825260608501516020830152820152608083015191506001600160401b03821115612169575f5ffd5b6121758683850161206a565b606082015295945050505050565b5f60208284031215612193575f5ffd5b81516001600160401b038111156121a8575f5ffd5b8201601f810184136121b8575f5ffd5b80516121c661208782612048565b8082825260208201915060208360051b8501019250868311156121e7575f5ffd5b602084015b838110156122795780516001600160401b03811115612209575f5ffd5b85016040818a03601f1901121561221e575f5ffd5b612226611ec9565b60208201516122348161167b565b815260408201516001600160401b0381111561224e575f5ffd5b61225d8b60208386010161206a565b60208301525080855250506020830192506020810190506121ec565b509695505050505050565b5f60208284031215612294575f5ffd5b81516115a28161167b565b5f826122b957634e487b7160e01b5f52601260045260245ffd5b500690565b8082018082111561137e57634e487b7160e01b5f52601160045260245ffdfea264697066735822122070e8c22e6e0af983e77b028aa66349b591500fd776d68dfc54a13e38f6cdee5364736f6c634300081b0033",
}

// OperatorTableUpdaterABI is the input ABI used to generate the binding from.
// Deprecated: Use OperatorTableUpdaterMetaData.ABI instead.
var OperatorTableUpdaterABI = OperatorTableUpdaterMetaData.ABI

// OperatorTableUpdaterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OperatorTableUpdaterMetaData.Bin instead.
var OperatorTableUpdaterBin = OperatorTableUpdaterMetaData.Bin

// DeployOperatorTableUpdater deploys a new Ethereum contract, binding an instance of OperatorTableUpdater to it.
func DeployOperatorTableUpdater(auth *bind.TransactOpts, backend bind.ContractBackend, _bn254CertificateVerifier common.Address, _ecdsaCertificateVerifier common.Address, _pauserRegistry common.Address, _version string) (common.Address, *types.Transaction, *OperatorTableUpdater, error) {
	parsed, err := OperatorTableUpdaterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OperatorTableUpdaterBin), backend, _bn254CertificateVerifier, _ecdsaCertificateVerifier, _pauserRegistry, _version)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OperatorTableUpdater{OperatorTableUpdaterCaller: OperatorTableUpdaterCaller{contract: contract}, OperatorTableUpdaterTransactor: OperatorTableUpdaterTransactor{contract: contract}, OperatorTableUpdaterFilterer: OperatorTableUpdaterFilterer{contract: contract}}, nil
}

// OperatorTableUpdater is an auto generated Go binding around an Ethereum contract.
type OperatorTableUpdater struct {
	OperatorTableUpdaterCaller     // Read-only binding to the contract
	OperatorTableUpdaterTransactor // Write-only binding to the contract
	OperatorTableUpdaterFilterer   // Log filterer for contract events
}

// OperatorTableUpdaterCaller is an auto generated read-only Go binding around an Ethereum contract.
type OperatorTableUpdaterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OperatorTableUpdaterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OperatorTableUpdaterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OperatorTableUpdaterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OperatorTableUpdaterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OperatorTableUpdaterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OperatorTableUpdaterSession struct {
	Contract     *OperatorTableUpdater // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// OperatorTableUpdaterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OperatorTableUpdaterCallerSession struct {
	Contract *OperatorTableUpdaterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// OperatorTableUpdaterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OperatorTableUpdaterTransactorSession struct {
	Contract     *OperatorTableUpdaterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// OperatorTableUpdaterRaw is an auto generated low-level Go binding around an Ethereum contract.
type OperatorTableUpdaterRaw struct {
	Contract *OperatorTableUpdater // Generic contract binding to access the raw methods on
}

// OperatorTableUpdaterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OperatorTableUpdaterCallerRaw struct {
	Contract *OperatorTableUpdaterCaller // Generic read-only contract binding to access the raw methods on
}

// OperatorTableUpdaterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OperatorTableUpdaterTransactorRaw struct {
	Contract *OperatorTableUpdaterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOperatorTableUpdater creates a new instance of OperatorTableUpdater, bound to a specific deployed contract.
func NewOperatorTableUpdater(address common.Address, backend bind.ContractBackend) (*OperatorTableUpdater, error) {
	contract, err := bindOperatorTableUpdater(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OperatorTableUpdater{OperatorTableUpdaterCaller: OperatorTableUpdaterCaller{contract: contract}, OperatorTableUpdaterTransactor: OperatorTableUpdaterTransactor{contract: contract}, OperatorTableUpdaterFilterer: OperatorTableUpdaterFilterer{contract: contract}}, nil
}

// NewOperatorTableUpdaterCaller creates a new read-only instance of OperatorTableUpdater, bound to a specific deployed contract.
func NewOperatorTableUpdaterCaller(address common.Address, caller bind.ContractCaller) (*OperatorTableUpdaterCaller, error) {
	contract, err := bindOperatorTableUpdater(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OperatorTableUpdaterCaller{contract: contract}, nil
}

// NewOperatorTableUpdaterTransactor creates a new write-only instance of OperatorTableUpdater, bound to a specific deployed contract.
func NewOperatorTableUpdaterTransactor(address common.Address, transactor bind.ContractTransactor) (*OperatorTableUpdaterTransactor, error) {
	contract, err := bindOperatorTableUpdater(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OperatorTableUpdaterTransactor{contract: contract}, nil
}

// NewOperatorTableUpdaterFilterer creates a new log filterer instance of OperatorTableUpdater, bound to a specific deployed contract.
func NewOperatorTableUpdaterFilterer(address common.Address, filterer bind.ContractFilterer) (*OperatorTableUpdaterFilterer, error) {
	contract, err := bindOperatorTableUpdater(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OperatorTableUpdaterFilterer{contract: contract}, nil
}

// bindOperatorTableUpdater binds a generic wrapper to an already deployed contract.
func bindOperatorTableUpdater(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OperatorTableUpdaterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OperatorTableUpdater *OperatorTableUpdaterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OperatorTableUpdater.Contract.OperatorTableUpdaterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OperatorTableUpdater *OperatorTableUpdaterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OperatorTableUpdater.Contract.OperatorTableUpdaterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OperatorTableUpdater *OperatorTableUpdaterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OperatorTableUpdater.Contract.OperatorTableUpdaterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OperatorTableUpdater *OperatorTableUpdaterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OperatorTableUpdater.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OperatorTableUpdater *OperatorTableUpdaterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OperatorTableUpdater.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OperatorTableUpdater *OperatorTableUpdaterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OperatorTableUpdater.Contract.contract.Transact(opts, method, params...)
}

// GLOBALTABLEROOTCERTTYPEHASH is a free data retrieval call binding the contract method 0x3ef6cd7a.
//
// Solidity: function GLOBAL_TABLE_ROOT_CERT_TYPEHASH() view returns(bytes32)
func (_OperatorTableUpdater *OperatorTableUpdaterCaller) GLOBALTABLEROOTCERTTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _OperatorTableUpdater.contract.Call(opts, &out, "GLOBAL_TABLE_ROOT_CERT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GLOBALTABLEROOTCERTTYPEHASH is a free data retrieval call binding the contract method 0x3ef6cd7a.
//
// Solidity: function GLOBAL_TABLE_ROOT_CERT_TYPEHASH() view returns(bytes32)
func (_OperatorTableUpdater *OperatorTableUpdaterSession) GLOBALTABLEROOTCERTTYPEHASH() ([32]byte, error) {
	return _OperatorTableUpdater.Contract.GLOBALTABLEROOTCERTTYPEHASH(&_OperatorTableUpdater.CallOpts)
}

// GLOBALTABLEROOTCERTTYPEHASH is a free data retrieval call binding the contract method 0x3ef6cd7a.
//
// Solidity: function GLOBAL_TABLE_ROOT_CERT_TYPEHASH() view returns(bytes32)
func (_OperatorTableUpdater *OperatorTableUpdaterCallerSession) GLOBALTABLEROOTCERTTYPEHASH() ([32]byte, error) {
	return _OperatorTableUpdater.Contract.GLOBALTABLEROOTCERTTYPEHASH(&_OperatorTableUpdater.CallOpts)
}

// MAXBPS is a free data retrieval call binding the contract method 0xfd967f47.
//
// Solidity: function MAX_BPS() view returns(uint16)
func (_OperatorTableUpdater *OperatorTableUpdaterCaller) MAXBPS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _OperatorTableUpdater.contract.Call(opts, &out, "MAX_BPS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MAXBPS is a free data retrieval call binding the contract method 0xfd967f47.
//
// Solidity: function MAX_BPS() view returns(uint16)
func (_OperatorTableUpdater *OperatorTableUpdaterSession) MAXBPS() (uint16, error) {
	return _OperatorTableUpdater.Contract.MAXBPS(&_OperatorTableUpdater.CallOpts)
}

// MAXBPS is a free data retrieval call binding the contract method 0xfd967f47.
//
// Solidity: function MAX_BPS() view returns(uint16)
func (_OperatorTableUpdater *OperatorTableUpdaterCallerSession) MAXBPS() (uint16, error) {
	return _OperatorTableUpdater.Contract.MAXBPS(&_OperatorTableUpdater.CallOpts)
}

// Bn254CertificateVerifier is a free data retrieval call binding the contract method 0xb8c14306.
//
// Solidity: function bn254CertificateVerifier() view returns(address)
func (_OperatorTableUpdater *OperatorTableUpdaterCaller) Bn254CertificateVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OperatorTableUpdater.contract.Call(opts, &out, "bn254CertificateVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bn254CertificateVerifier is a free data retrieval call binding the contract method 0xb8c14306.
//
// Solidity: function bn254CertificateVerifier() view returns(address)
func (_OperatorTableUpdater *OperatorTableUpdaterSession) Bn254CertificateVerifier() (common.Address, error) {
	return _OperatorTableUpdater.Contract.Bn254CertificateVerifier(&_OperatorTableUpdater.CallOpts)
}

// Bn254CertificateVerifier is a free data retrieval call binding the contract method 0xb8c14306.
//
// Solidity: function bn254CertificateVerifier() view returns(address)
func (_OperatorTableUpdater *OperatorTableUpdaterCallerSession) Bn254CertificateVerifier() (common.Address, error) {
	return _OperatorTableUpdater.Contract.Bn254CertificateVerifier(&_OperatorTableUpdater.CallOpts)
}

// EcdsaCertificateVerifier is a free data retrieval call binding the contract method 0xad0f9582.
//
// Solidity: function ecdsaCertificateVerifier() view returns(address)
func (_OperatorTableUpdater *OperatorTableUpdaterCaller) EcdsaCertificateVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OperatorTableUpdater.contract.Call(opts, &out, "ecdsaCertificateVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EcdsaCertificateVerifier is a free data retrieval call binding the contract method 0xad0f9582.
//
// Solidity: function ecdsaCertificateVerifier() view returns(address)
func (_OperatorTableUpdater *OperatorTableUpdaterSession) EcdsaCertificateVerifier() (common.Address, error) {
	return _OperatorTableUpdater.Contract.EcdsaCertificateVerifier(&_OperatorTableUpdater.CallOpts)
}

// EcdsaCertificateVerifier is a free data retrieval call binding the contract method 0xad0f9582.
//
// Solidity: function ecdsaCertificateVerifier() view returns(address)
func (_OperatorTableUpdater *OperatorTableUpdaterCallerSession) EcdsaCertificateVerifier() (common.Address, error) {
	return _OperatorTableUpdater.Contract.EcdsaCertificateVerifier(&_OperatorTableUpdater.CallOpts)
}

// GetCertificateVerifier is a free data retrieval call binding the contract method 0x6f728c50.
//
// Solidity: function getCertificateVerifier(uint8 curveType) view returns(address)
func (_OperatorTableUpdater *OperatorTableUpdaterCaller) GetCertificateVerifier(opts *bind.CallOpts, curveType uint8) (common.Address, error) {
	var out []interface{}
	err := _OperatorTableUpdater.contract.Call(opts, &out, "getCertificateVerifier", curveType)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCertificateVerifier is a free data retrieval call binding the contract method 0x6f728c50.
//
// Solidity: function getCertificateVerifier(uint8 curveType) view returns(address)
func (_OperatorTableUpdater *OperatorTableUpdaterSession) GetCertificateVerifier(curveType uint8) (common.Address, error) {
	return _OperatorTableUpdater.Contract.GetCertificateVerifier(&_OperatorTableUpdater.CallOpts, curveType)
}

// GetCertificateVerifier is a free data retrieval call binding the contract method 0x6f728c50.
//
// Solidity: function getCertificateVerifier(uint8 curveType) view returns(address)
func (_OperatorTableUpdater *OperatorTableUpdaterCallerSession) GetCertificateVerifier(curveType uint8) (common.Address, error) {
	return _OperatorTableUpdater.Contract.GetCertificateVerifier(&_OperatorTableUpdater.CallOpts, curveType)
}

// GetCurrentGlobalTableRoot is a free data retrieval call binding the contract method 0x28522d79.
//
// Solidity: function getCurrentGlobalTableRoot() view returns(bytes32)
func (_OperatorTableUpdater *OperatorTableUpdaterCaller) GetCurrentGlobalTableRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _OperatorTableUpdater.contract.Call(opts, &out, "getCurrentGlobalTableRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetCurrentGlobalTableRoot is a free data retrieval call binding the contract method 0x28522d79.
//
// Solidity: function getCurrentGlobalTableRoot() view returns(bytes32)
func (_OperatorTableUpdater *OperatorTableUpdaterSession) GetCurrentGlobalTableRoot() ([32]byte, error) {
	return _OperatorTableUpdater.Contract.GetCurrentGlobalTableRoot(&_OperatorTableUpdater.CallOpts)
}

// GetCurrentGlobalTableRoot is a free data retrieval call binding the contract method 0x28522d79.
//
// Solidity: function getCurrentGlobalTableRoot() view returns(bytes32)
func (_OperatorTableUpdater *OperatorTableUpdaterCallerSession) GetCurrentGlobalTableRoot() ([32]byte, error) {
	return _OperatorTableUpdater.Contract.GetCurrentGlobalTableRoot(&_OperatorTableUpdater.CallOpts)
}

// GetGenerator is a free data retrieval call binding the contract method 0x1e2ca260.
//
// Solidity: function getGenerator() view returns((address,uint32))
func (_OperatorTableUpdater *OperatorTableUpdaterCaller) GetGenerator(opts *bind.CallOpts) (OperatorSet, error) {
	var out []interface{}
	err := _OperatorTableUpdater.contract.Call(opts, &out, "getGenerator")

	if err != nil {
		return *new(OperatorSet), err
	}

	out0 := *abi.ConvertType(out[0], new(OperatorSet)).(*OperatorSet)

	return out0, err

}

// GetGenerator is a free data retrieval call binding the contract method 0x1e2ca260.
//
// Solidity: function getGenerator() view returns((address,uint32))
func (_OperatorTableUpdater *OperatorTableUpdaterSession) GetGenerator() (OperatorSet, error) {
	return _OperatorTableUpdater.Contract.GetGenerator(&_OperatorTableUpdater.CallOpts)
}

// GetGenerator is a free data retrieval call binding the contract method 0x1e2ca260.
//
// Solidity: function getGenerator() view returns((address,uint32))
func (_OperatorTableUpdater *OperatorTableUpdaterCallerSession) GetGenerator() (OperatorSet, error) {
	return _OperatorTableUpdater.Contract.GetGenerator(&_OperatorTableUpdater.CallOpts)
}

// GetGeneratorReferenceTimestamp is a free data retrieval call binding the contract method 0x7551ba34.
//
// Solidity: function getGeneratorReferenceTimestamp() view returns(uint32)
func (_OperatorTableUpdater *OperatorTableUpdaterCaller) GetGeneratorReferenceTimestamp(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _OperatorTableUpdater.contract.Call(opts, &out, "getGeneratorReferenceTimestamp")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetGeneratorReferenceTimestamp is a free data retrieval call binding the contract method 0x7551ba34.
//
// Solidity: function getGeneratorReferenceTimestamp() view returns(uint32)
func (_OperatorTableUpdater *OperatorTableUpdaterSession) GetGeneratorReferenceTimestamp() (uint32, error) {
	return _OperatorTableUpdater.Contract.GetGeneratorReferenceTimestamp(&_OperatorTableUpdater.CallOpts)
}

// GetGeneratorReferenceTimestamp is a free data retrieval call binding the contract method 0x7551ba34.
//
// Solidity: function getGeneratorReferenceTimestamp() view returns(uint32)
func (_OperatorTableUpdater *OperatorTableUpdaterCallerSession) GetGeneratorReferenceTimestamp() (uint32, error) {
	return _OperatorTableUpdater.Contract.GetGeneratorReferenceTimestamp(&_OperatorTableUpdater.CallOpts)
}

// GetGlobalTableRootByTimestamp is a free data retrieval call binding the contract method 0xc5916a39.
//
// Solidity: function getGlobalTableRootByTimestamp(uint32 referenceTimestamp) view returns(bytes32)
func (_OperatorTableUpdater *OperatorTableUpdaterCaller) GetGlobalTableRootByTimestamp(opts *bind.CallOpts, referenceTimestamp uint32) ([32]byte, error) {
	var out []interface{}
	err := _OperatorTableUpdater.contract.Call(opts, &out, "getGlobalTableRootByTimestamp", referenceTimestamp)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetGlobalTableRootByTimestamp is a free data retrieval call binding the contract method 0xc5916a39.
//
// Solidity: function getGlobalTableRootByTimestamp(uint32 referenceTimestamp) view returns(bytes32)
func (_OperatorTableUpdater *OperatorTableUpdaterSession) GetGlobalTableRootByTimestamp(referenceTimestamp uint32) ([32]byte, error) {
	return _OperatorTableUpdater.Contract.GetGlobalTableRootByTimestamp(&_OperatorTableUpdater.CallOpts, referenceTimestamp)
}

// GetGlobalTableRootByTimestamp is a free data retrieval call binding the contract method 0xc5916a39.
//
// Solidity: function getGlobalTableRootByTimestamp(uint32 referenceTimestamp) view returns(bytes32)
func (_OperatorTableUpdater *OperatorTableUpdaterCallerSession) GetGlobalTableRootByTimestamp(referenceTimestamp uint32) ([32]byte, error) {
	return _OperatorTableUpdater.Contract.GetGlobalTableRootByTimestamp(&_OperatorTableUpdater.CallOpts, referenceTimestamp)
}

// GetGlobalTableUpdateMessageHash is a free data retrieval call binding the contract method 0xc3be1e33.
//
// Solidity: function getGlobalTableUpdateMessageHash(bytes32 globalTableRoot, uint32 referenceTimestamp, uint32 referenceBlockNumber) pure returns(bytes32)
func (_OperatorTableUpdater *OperatorTableUpdaterCaller) GetGlobalTableUpdateMessageHash(opts *bind.CallOpts, globalTableRoot [32]byte, referenceTimestamp uint32, referenceBlockNumber uint32) ([32]byte, error) {
	var out []interface{}
	err := _OperatorTableUpdater.contract.Call(opts, &out, "getGlobalTableUpdateMessageHash", globalTableRoot, referenceTimestamp, referenceBlockNumber)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetGlobalTableUpdateMessageHash is a free data retrieval call binding the contract method 0xc3be1e33.
//
// Solidity: function getGlobalTableUpdateMessageHash(bytes32 globalTableRoot, uint32 referenceTimestamp, uint32 referenceBlockNumber) pure returns(bytes32)
func (_OperatorTableUpdater *OperatorTableUpdaterSession) GetGlobalTableUpdateMessageHash(globalTableRoot [32]byte, referenceTimestamp uint32, referenceBlockNumber uint32) ([32]byte, error) {
	return _OperatorTableUpdater.Contract.GetGlobalTableUpdateMessageHash(&_OperatorTableUpdater.CallOpts, globalTableRoot, referenceTimestamp, referenceBlockNumber)
}

// GetGlobalTableUpdateMessageHash is a free data retrieval call binding the contract method 0xc3be1e33.
//
// Solidity: function getGlobalTableUpdateMessageHash(bytes32 globalTableRoot, uint32 referenceTimestamp, uint32 referenceBlockNumber) pure returns(bytes32)
func (_OperatorTableUpdater *OperatorTableUpdaterCallerSession) GetGlobalTableUpdateMessageHash(globalTableRoot [32]byte, referenceTimestamp uint32, referenceBlockNumber uint32) ([32]byte, error) {
	return _OperatorTableUpdater.Contract.GetGlobalTableUpdateMessageHash(&_OperatorTableUpdater.CallOpts, globalTableRoot, referenceTimestamp, referenceBlockNumber)
}

// GetLatestReferenceBlockNumber is a free data retrieval call binding the contract method 0x31a599d2.
//
// Solidity: function getLatestReferenceBlockNumber() view returns(uint32)
func (_OperatorTableUpdater *OperatorTableUpdaterCaller) GetLatestReferenceBlockNumber(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _OperatorTableUpdater.contract.Call(opts, &out, "getLatestReferenceBlockNumber")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetLatestReferenceBlockNumber is a free data retrieval call binding the contract method 0x31a599d2.
//
// Solidity: function getLatestReferenceBlockNumber() view returns(uint32)
func (_OperatorTableUpdater *OperatorTableUpdaterSession) GetLatestReferenceBlockNumber() (uint32, error) {
	return _OperatorTableUpdater.Contract.GetLatestReferenceBlockNumber(&_OperatorTableUpdater.CallOpts)
}

// GetLatestReferenceBlockNumber is a free data retrieval call binding the contract method 0x31a599d2.
//
// Solidity: function getLatestReferenceBlockNumber() view returns(uint32)
func (_OperatorTableUpdater *OperatorTableUpdaterCallerSession) GetLatestReferenceBlockNumber() (uint32, error) {
	return _OperatorTableUpdater.Contract.GetLatestReferenceBlockNumber(&_OperatorTableUpdater.CallOpts)
}

// GetLatestReferenceTimestamp is a free data retrieval call binding the contract method 0x4624e6a3.
//
// Solidity: function getLatestReferenceTimestamp() view returns(uint32)
func (_OperatorTableUpdater *OperatorTableUpdaterCaller) GetLatestReferenceTimestamp(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _OperatorTableUpdater.contract.Call(opts, &out, "getLatestReferenceTimestamp")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetLatestReferenceTimestamp is a free data retrieval call binding the contract method 0x4624e6a3.
//
// Solidity: function getLatestReferenceTimestamp() view returns(uint32)
func (_OperatorTableUpdater *OperatorTableUpdaterSession) GetLatestReferenceTimestamp() (uint32, error) {
	return _OperatorTableUpdater.Contract.GetLatestReferenceTimestamp(&_OperatorTableUpdater.CallOpts)
}

// GetLatestReferenceTimestamp is a free data retrieval call binding the contract method 0x4624e6a3.
//
// Solidity: function getLatestReferenceTimestamp() view returns(uint32)
func (_OperatorTableUpdater *OperatorTableUpdaterCallerSession) GetLatestReferenceTimestamp() (uint32, error) {
	return _OperatorTableUpdater.Contract.GetLatestReferenceTimestamp(&_OperatorTableUpdater.CallOpts)
}

// GetReferenceBlockNumberByTimestamp is a free data retrieval call binding the contract method 0x23b7b5b2.
//
// Solidity: function getReferenceBlockNumberByTimestamp(uint32 referenceTimestamp) view returns(uint32)
func (_OperatorTableUpdater *OperatorTableUpdaterCaller) GetReferenceBlockNumberByTimestamp(opts *bind.CallOpts, referenceTimestamp uint32) (uint32, error) {
	var out []interface{}
	err := _OperatorTableUpdater.contract.Call(opts, &out, "getReferenceBlockNumberByTimestamp", referenceTimestamp)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetReferenceBlockNumberByTimestamp is a free data retrieval call binding the contract method 0x23b7b5b2.
//
// Solidity: function getReferenceBlockNumberByTimestamp(uint32 referenceTimestamp) view returns(uint32)
func (_OperatorTableUpdater *OperatorTableUpdaterSession) GetReferenceBlockNumberByTimestamp(referenceTimestamp uint32) (uint32, error) {
	return _OperatorTableUpdater.Contract.GetReferenceBlockNumberByTimestamp(&_OperatorTableUpdater.CallOpts, referenceTimestamp)
}

// GetReferenceBlockNumberByTimestamp is a free data retrieval call binding the contract method 0x23b7b5b2.
//
// Solidity: function getReferenceBlockNumberByTimestamp(uint32 referenceTimestamp) view returns(uint32)
func (_OperatorTableUpdater *OperatorTableUpdaterCallerSession) GetReferenceBlockNumberByTimestamp(referenceTimestamp uint32) (uint32, error) {
	return _OperatorTableUpdater.Contract.GetReferenceBlockNumberByTimestamp(&_OperatorTableUpdater.CallOpts, referenceTimestamp)
}

// GetReferenceTimestampByBlockNumber is a free data retrieval call binding the contract method 0x193b79f3.
//
// Solidity: function getReferenceTimestampByBlockNumber(uint32 referenceBlockNumber) view returns(uint32)
func (_OperatorTableUpdater *OperatorTableUpdaterCaller) GetReferenceTimestampByBlockNumber(opts *bind.CallOpts, referenceBlockNumber uint32) (uint32, error) {
	var out []interface{}
	err := _OperatorTableUpdater.contract.Call(opts, &out, "getReferenceTimestampByBlockNumber", referenceBlockNumber)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetReferenceTimestampByBlockNumber is a free data retrieval call binding the contract method 0x193b79f3.
//
// Solidity: function getReferenceTimestampByBlockNumber(uint32 referenceBlockNumber) view returns(uint32)
func (_OperatorTableUpdater *OperatorTableUpdaterSession) GetReferenceTimestampByBlockNumber(referenceBlockNumber uint32) (uint32, error) {
	return _OperatorTableUpdater.Contract.GetReferenceTimestampByBlockNumber(&_OperatorTableUpdater.CallOpts, referenceBlockNumber)
}

// GetReferenceTimestampByBlockNumber is a free data retrieval call binding the contract method 0x193b79f3.
//
// Solidity: function getReferenceTimestampByBlockNumber(uint32 referenceBlockNumber) view returns(uint32)
func (_OperatorTableUpdater *OperatorTableUpdaterCallerSession) GetReferenceTimestampByBlockNumber(referenceBlockNumber uint32) (uint32, error) {
	return _OperatorTableUpdater.Contract.GetReferenceTimestampByBlockNumber(&_OperatorTableUpdater.CallOpts, referenceBlockNumber)
}

// GlobalRootConfirmationThreshold is a free data retrieval call binding the contract method 0xc252aa22.
//
// Solidity: function globalRootConfirmationThreshold() view returns(uint16)
func (_OperatorTableUpdater *OperatorTableUpdaterCaller) GlobalRootConfirmationThreshold(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _OperatorTableUpdater.contract.Call(opts, &out, "globalRootConfirmationThreshold")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GlobalRootConfirmationThreshold is a free data retrieval call binding the contract method 0xc252aa22.
//
// Solidity: function globalRootConfirmationThreshold() view returns(uint16)
func (_OperatorTableUpdater *OperatorTableUpdaterSession) GlobalRootConfirmationThreshold() (uint16, error) {
	return _OperatorTableUpdater.Contract.GlobalRootConfirmationThreshold(&_OperatorTableUpdater.CallOpts)
}

// GlobalRootConfirmationThreshold is a free data retrieval call binding the contract method 0xc252aa22.
//
// Solidity: function globalRootConfirmationThreshold() view returns(uint16)
func (_OperatorTableUpdater *OperatorTableUpdaterCallerSession) GlobalRootConfirmationThreshold() (uint16, error) {
	return _OperatorTableUpdater.Contract.GlobalRootConfirmationThreshold(&_OperatorTableUpdater.CallOpts)
}

// IsRootValid is a free data retrieval call binding the contract method 0x30ef41b4.
//
// Solidity: function isRootValid(bytes32 globalTableRoot) view returns(bool)
func (_OperatorTableUpdater *OperatorTableUpdaterCaller) IsRootValid(opts *bind.CallOpts, globalTableRoot [32]byte) (bool, error) {
	var out []interface{}
	err := _OperatorTableUpdater.contract.Call(opts, &out, "isRootValid", globalTableRoot)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRootValid is a free data retrieval call binding the contract method 0x30ef41b4.
//
// Solidity: function isRootValid(bytes32 globalTableRoot) view returns(bool)
func (_OperatorTableUpdater *OperatorTableUpdaterSession) IsRootValid(globalTableRoot [32]byte) (bool, error) {
	return _OperatorTableUpdater.Contract.IsRootValid(&_OperatorTableUpdater.CallOpts, globalTableRoot)
}

// IsRootValid is a free data retrieval call binding the contract method 0x30ef41b4.
//
// Solidity: function isRootValid(bytes32 globalTableRoot) view returns(bool)
func (_OperatorTableUpdater *OperatorTableUpdaterCallerSession) IsRootValid(globalTableRoot [32]byte) (bool, error) {
	return _OperatorTableUpdater.Contract.IsRootValid(&_OperatorTableUpdater.CallOpts, globalTableRoot)
}

// IsRootValidByTimestamp is a free data retrieval call binding the contract method 0x64e1df84.
//
// Solidity: function isRootValidByTimestamp(uint32 referenceTimestamp) view returns(bool)
func (_OperatorTableUpdater *OperatorTableUpdaterCaller) IsRootValidByTimestamp(opts *bind.CallOpts, referenceTimestamp uint32) (bool, error) {
	var out []interface{}
	err := _OperatorTableUpdater.contract.Call(opts, &out, "isRootValidByTimestamp", referenceTimestamp)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRootValidByTimestamp is a free data retrieval call binding the contract method 0x64e1df84.
//
// Solidity: function isRootValidByTimestamp(uint32 referenceTimestamp) view returns(bool)
func (_OperatorTableUpdater *OperatorTableUpdaterSession) IsRootValidByTimestamp(referenceTimestamp uint32) (bool, error) {
	return _OperatorTableUpdater.Contract.IsRootValidByTimestamp(&_OperatorTableUpdater.CallOpts, referenceTimestamp)
}

// IsRootValidByTimestamp is a free data retrieval call binding the contract method 0x64e1df84.
//
// Solidity: function isRootValidByTimestamp(uint32 referenceTimestamp) view returns(bool)
func (_OperatorTableUpdater *OperatorTableUpdaterCallerSession) IsRootValidByTimestamp(referenceTimestamp uint32) (bool, error) {
	return _OperatorTableUpdater.Contract.IsRootValidByTimestamp(&_OperatorTableUpdater.CallOpts, referenceTimestamp)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OperatorTableUpdater *OperatorTableUpdaterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OperatorTableUpdater.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OperatorTableUpdater *OperatorTableUpdaterSession) Owner() (common.Address, error) {
	return _OperatorTableUpdater.Contract.Owner(&_OperatorTableUpdater.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OperatorTableUpdater *OperatorTableUpdaterCallerSession) Owner() (common.Address, error) {
	return _OperatorTableUpdater.Contract.Owner(&_OperatorTableUpdater.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_OperatorTableUpdater *OperatorTableUpdaterCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _OperatorTableUpdater.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_OperatorTableUpdater *OperatorTableUpdaterSession) Paused(index uint8) (bool, error) {
	return _OperatorTableUpdater.Contract.Paused(&_OperatorTableUpdater.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_OperatorTableUpdater *OperatorTableUpdaterCallerSession) Paused(index uint8) (bool, error) {
	return _OperatorTableUpdater.Contract.Paused(&_OperatorTableUpdater.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_OperatorTableUpdater *OperatorTableUpdaterCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OperatorTableUpdater.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_OperatorTableUpdater *OperatorTableUpdaterSession) Paused0() (*big.Int, error) {
	return _OperatorTableUpdater.Contract.Paused0(&_OperatorTableUpdater.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_OperatorTableUpdater *OperatorTableUpdaterCallerSession) Paused0() (*big.Int, error) {
	return _OperatorTableUpdater.Contract.Paused0(&_OperatorTableUpdater.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_OperatorTableUpdater *OperatorTableUpdaterCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OperatorTableUpdater.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_OperatorTableUpdater *OperatorTableUpdaterSession) PauserRegistry() (common.Address, error) {
	return _OperatorTableUpdater.Contract.PauserRegistry(&_OperatorTableUpdater.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_OperatorTableUpdater *OperatorTableUpdaterCallerSession) PauserRegistry() (common.Address, error) {
	return _OperatorTableUpdater.Contract.PauserRegistry(&_OperatorTableUpdater.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_OperatorTableUpdater *OperatorTableUpdaterCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OperatorTableUpdater.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_OperatorTableUpdater *OperatorTableUpdaterSession) Version() (string, error) {
	return _OperatorTableUpdater.Contract.Version(&_OperatorTableUpdater.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_OperatorTableUpdater *OperatorTableUpdaterCallerSession) Version() (string, error) {
	return _OperatorTableUpdater.Contract.Version(&_OperatorTableUpdater.CallOpts)
}

// ConfirmGlobalTableRoot is a paid mutator transaction binding the contract method 0xeaaed9d5.
//
// Solidity: function confirmGlobalTableRoot((uint32,bytes32,(uint256,uint256),(uint256[2],uint256[2]),(uint32,bytes,((uint256,uint256),uint256[]))[]) globalTableRootCert, bytes32 globalTableRoot, uint32 referenceTimestamp, uint32 referenceBlockNumber) returns()
func (_OperatorTableUpdater *OperatorTableUpdaterTransactor) ConfirmGlobalTableRoot(opts *bind.TransactOpts, globalTableRootCert IBN254CertificateVerifierTypesBN254Certificate, globalTableRoot [32]byte, referenceTimestamp uint32, referenceBlockNumber uint32) (*types.Transaction, error) {
	return _OperatorTableUpdater.contract.Transact(opts, "confirmGlobalTableRoot", globalTableRootCert, globalTableRoot, referenceTimestamp, referenceBlockNumber)
}

// ConfirmGlobalTableRoot is a paid mutator transaction binding the contract method 0xeaaed9d5.
//
// Solidity: function confirmGlobalTableRoot((uint32,bytes32,(uint256,uint256),(uint256[2],uint256[2]),(uint32,bytes,((uint256,uint256),uint256[]))[]) globalTableRootCert, bytes32 globalTableRoot, uint32 referenceTimestamp, uint32 referenceBlockNumber) returns()
func (_OperatorTableUpdater *OperatorTableUpdaterSession) ConfirmGlobalTableRoot(globalTableRootCert IBN254CertificateVerifierTypesBN254Certificate, globalTableRoot [32]byte, referenceTimestamp uint32, referenceBlockNumber uint32) (*types.Transaction, error) {
	return _OperatorTableUpdater.Contract.ConfirmGlobalTableRoot(&_OperatorTableUpdater.TransactOpts, globalTableRootCert, globalTableRoot, referenceTimestamp, referenceBlockNumber)
}

// ConfirmGlobalTableRoot is a paid mutator transaction binding the contract method 0xeaaed9d5.
//
// Solidity: function confirmGlobalTableRoot((uint32,bytes32,(uint256,uint256),(uint256[2],uint256[2]),(uint32,bytes,((uint256,uint256),uint256[]))[]) globalTableRootCert, bytes32 globalTableRoot, uint32 referenceTimestamp, uint32 referenceBlockNumber) returns()
func (_OperatorTableUpdater *OperatorTableUpdaterTransactorSession) ConfirmGlobalTableRoot(globalTableRootCert IBN254CertificateVerifierTypesBN254Certificate, globalTableRoot [32]byte, referenceTimestamp uint32, referenceBlockNumber uint32) (*types.Transaction, error) {
	return _OperatorTableUpdater.Contract.ConfirmGlobalTableRoot(&_OperatorTableUpdater.TransactOpts, globalTableRootCert, globalTableRoot, referenceTimestamp, referenceBlockNumber)
}

// DisableRoot is a paid mutator transaction binding the contract method 0xc3621f0a.
//
// Solidity: function disableRoot(bytes32 globalTableRoot) returns()
func (_OperatorTableUpdater *OperatorTableUpdaterTransactor) DisableRoot(opts *bind.TransactOpts, globalTableRoot [32]byte) (*types.Transaction, error) {
	return _OperatorTableUpdater.contract.Transact(opts, "disableRoot", globalTableRoot)
}

// DisableRoot is a paid mutator transaction binding the contract method 0xc3621f0a.
//
// Solidity: function disableRoot(bytes32 globalTableRoot) returns()
func (_OperatorTableUpdater *OperatorTableUpdaterSession) DisableRoot(globalTableRoot [32]byte) (*types.Transaction, error) {
	return _OperatorTableUpdater.Contract.DisableRoot(&_OperatorTableUpdater.TransactOpts, globalTableRoot)
}

// DisableRoot is a paid mutator transaction binding the contract method 0xc3621f0a.
//
// Solidity: function disableRoot(bytes32 globalTableRoot) returns()
func (_OperatorTableUpdater *OperatorTableUpdaterTransactorSession) DisableRoot(globalTableRoot [32]byte) (*types.Transaction, error) {
	return _OperatorTableUpdater.Contract.DisableRoot(&_OperatorTableUpdater.TransactOpts, globalTableRoot)
}

// Initialize is a paid mutator transaction binding the contract method 0x8f797718.
//
// Solidity: function initialize(address owner, uint256 initialPausedStatus, (address,uint32) generator, uint16 _globalRootConfirmationThreshold, bytes32 initialGlobalTableRoot, uint32 referenceTimestamp, (bytes32,uint256,(uint256,uint256),uint256[]) generatorInfo, (address,uint32) generatorConfig) returns()
func (_OperatorTableUpdater *OperatorTableUpdaterTransactor) Initialize(opts *bind.TransactOpts, owner common.Address, initialPausedStatus *big.Int, generator OperatorSet, _globalRootConfirmationThreshold uint16, initialGlobalTableRoot [32]byte, referenceTimestamp uint32, generatorInfo IOperatorTableCalculatorTypesBN254OperatorSetInfo, generatorConfig ICrossChainRegistryTypesOperatorSetConfig) (*types.Transaction, error) {
	return _OperatorTableUpdater.contract.Transact(opts, "initialize", owner, initialPausedStatus, generator, _globalRootConfirmationThreshold, initialGlobalTableRoot, referenceTimestamp, generatorInfo, generatorConfig)
}

// Initialize is a paid mutator transaction binding the contract method 0x8f797718.
//
// Solidity: function initialize(address owner, uint256 initialPausedStatus, (address,uint32) generator, uint16 _globalRootConfirmationThreshold, bytes32 initialGlobalTableRoot, uint32 referenceTimestamp, (bytes32,uint256,(uint256,uint256),uint256[]) generatorInfo, (address,uint32) generatorConfig) returns()
func (_OperatorTableUpdater *OperatorTableUpdaterSession) Initialize(owner common.Address, initialPausedStatus *big.Int, generator OperatorSet, _globalRootConfirmationThreshold uint16, initialGlobalTableRoot [32]byte, referenceTimestamp uint32, generatorInfo IOperatorTableCalculatorTypesBN254OperatorSetInfo, generatorConfig ICrossChainRegistryTypesOperatorSetConfig) (*types.Transaction, error) {
	return _OperatorTableUpdater.Contract.Initialize(&_OperatorTableUpdater.TransactOpts, owner, initialPausedStatus, generator, _globalRootConfirmationThreshold, initialGlobalTableRoot, referenceTimestamp, generatorInfo, generatorConfig)
}

// Initialize is a paid mutator transaction binding the contract method 0x8f797718.
//
// Solidity: function initialize(address owner, uint256 initialPausedStatus, (address,uint32) generator, uint16 _globalRootConfirmationThreshold, bytes32 initialGlobalTableRoot, uint32 referenceTimestamp, (bytes32,uint256,(uint256,uint256),uint256[]) generatorInfo, (address,uint32) generatorConfig) returns()
func (_OperatorTableUpdater *OperatorTableUpdaterTransactorSession) Initialize(owner common.Address, initialPausedStatus *big.Int, generator OperatorSet, _globalRootConfirmationThreshold uint16, initialGlobalTableRoot [32]byte, referenceTimestamp uint32, generatorInfo IOperatorTableCalculatorTypesBN254OperatorSetInfo, generatorConfig ICrossChainRegistryTypesOperatorSetConfig) (*types.Transaction, error) {
	return _OperatorTableUpdater.Contract.Initialize(&_OperatorTableUpdater.TransactOpts, owner, initialPausedStatus, generator, _globalRootConfirmationThreshold, initialGlobalTableRoot, referenceTimestamp, generatorInfo, generatorConfig)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_OperatorTableUpdater *OperatorTableUpdaterTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _OperatorTableUpdater.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_OperatorTableUpdater *OperatorTableUpdaterSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _OperatorTableUpdater.Contract.Pause(&_OperatorTableUpdater.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_OperatorTableUpdater *OperatorTableUpdaterTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _OperatorTableUpdater.Contract.Pause(&_OperatorTableUpdater.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_OperatorTableUpdater *OperatorTableUpdaterTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OperatorTableUpdater.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_OperatorTableUpdater *OperatorTableUpdaterSession) PauseAll() (*types.Transaction, error) {
	return _OperatorTableUpdater.Contract.PauseAll(&_OperatorTableUpdater.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_OperatorTableUpdater *OperatorTableUpdaterTransactorSession) PauseAll() (*types.Transaction, error) {
	return _OperatorTableUpdater.Contract.PauseAll(&_OperatorTableUpdater.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OperatorTableUpdater *OperatorTableUpdaterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OperatorTableUpdater.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OperatorTableUpdater *OperatorTableUpdaterSession) RenounceOwnership() (*types.Transaction, error) {
	return _OperatorTableUpdater.Contract.RenounceOwnership(&_OperatorTableUpdater.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OperatorTableUpdater *OperatorTableUpdaterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _OperatorTableUpdater.Contract.RenounceOwnership(&_OperatorTableUpdater.TransactOpts)
}

// SetGlobalRootConfirmationThreshold is a paid mutator transaction binding the contract method 0x2370356c.
//
// Solidity: function setGlobalRootConfirmationThreshold(uint16 bps) returns()
func (_OperatorTableUpdater *OperatorTableUpdaterTransactor) SetGlobalRootConfirmationThreshold(opts *bind.TransactOpts, bps uint16) (*types.Transaction, error) {
	return _OperatorTableUpdater.contract.Transact(opts, "setGlobalRootConfirmationThreshold", bps)
}

// SetGlobalRootConfirmationThreshold is a paid mutator transaction binding the contract method 0x2370356c.
//
// Solidity: function setGlobalRootConfirmationThreshold(uint16 bps) returns()
func (_OperatorTableUpdater *OperatorTableUpdaterSession) SetGlobalRootConfirmationThreshold(bps uint16) (*types.Transaction, error) {
	return _OperatorTableUpdater.Contract.SetGlobalRootConfirmationThreshold(&_OperatorTableUpdater.TransactOpts, bps)
}

// SetGlobalRootConfirmationThreshold is a paid mutator transaction binding the contract method 0x2370356c.
//
// Solidity: function setGlobalRootConfirmationThreshold(uint16 bps) returns()
func (_OperatorTableUpdater *OperatorTableUpdaterTransactorSession) SetGlobalRootConfirmationThreshold(bps uint16) (*types.Transaction, error) {
	return _OperatorTableUpdater.Contract.SetGlobalRootConfirmationThreshold(&_OperatorTableUpdater.TransactOpts, bps)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OperatorTableUpdater *OperatorTableUpdaterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _OperatorTableUpdater.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OperatorTableUpdater *OperatorTableUpdaterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OperatorTableUpdater.Contract.TransferOwnership(&_OperatorTableUpdater.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OperatorTableUpdater *OperatorTableUpdaterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OperatorTableUpdater.Contract.TransferOwnership(&_OperatorTableUpdater.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_OperatorTableUpdater *OperatorTableUpdaterTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _OperatorTableUpdater.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_OperatorTableUpdater *OperatorTableUpdaterSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _OperatorTableUpdater.Contract.Unpause(&_OperatorTableUpdater.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_OperatorTableUpdater *OperatorTableUpdaterTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _OperatorTableUpdater.Contract.Unpause(&_OperatorTableUpdater.TransactOpts, newPausedStatus)
}

// UpdateOperatorTable is a paid mutator transaction binding the contract method 0x9ea94778.
//
// Solidity: function updateOperatorTable(uint32 referenceTimestamp, bytes32 globalTableRoot, uint32 operatorSetIndex, bytes proof, bytes operatorTableBytes) returns()
func (_OperatorTableUpdater *OperatorTableUpdaterTransactor) UpdateOperatorTable(opts *bind.TransactOpts, referenceTimestamp uint32, globalTableRoot [32]byte, operatorSetIndex uint32, proof []byte, operatorTableBytes []byte) (*types.Transaction, error) {
	return _OperatorTableUpdater.contract.Transact(opts, "updateOperatorTable", referenceTimestamp, globalTableRoot, operatorSetIndex, proof, operatorTableBytes)
}

// UpdateOperatorTable is a paid mutator transaction binding the contract method 0x9ea94778.
//
// Solidity: function updateOperatorTable(uint32 referenceTimestamp, bytes32 globalTableRoot, uint32 operatorSetIndex, bytes proof, bytes operatorTableBytes) returns()
func (_OperatorTableUpdater *OperatorTableUpdaterSession) UpdateOperatorTable(referenceTimestamp uint32, globalTableRoot [32]byte, operatorSetIndex uint32, proof []byte, operatorTableBytes []byte) (*types.Transaction, error) {
	return _OperatorTableUpdater.Contract.UpdateOperatorTable(&_OperatorTableUpdater.TransactOpts, referenceTimestamp, globalTableRoot, operatorSetIndex, proof, operatorTableBytes)
}

// UpdateOperatorTable is a paid mutator transaction binding the contract method 0x9ea94778.
//
// Solidity: function updateOperatorTable(uint32 referenceTimestamp, bytes32 globalTableRoot, uint32 operatorSetIndex, bytes proof, bytes operatorTableBytes) returns()
func (_OperatorTableUpdater *OperatorTableUpdaterTransactorSession) UpdateOperatorTable(referenceTimestamp uint32, globalTableRoot [32]byte, operatorSetIndex uint32, proof []byte, operatorTableBytes []byte) (*types.Transaction, error) {
	return _OperatorTableUpdater.Contract.UpdateOperatorTable(&_OperatorTableUpdater.TransactOpts, referenceTimestamp, globalTableRoot, operatorSetIndex, proof, operatorTableBytes)
}

// OperatorTableUpdaterGlobalRootConfirmationThresholdUpdatedIterator is returned from FilterGlobalRootConfirmationThresholdUpdated and is used to iterate over the raw logs and unpacked data for GlobalRootConfirmationThresholdUpdated events raised by the OperatorTableUpdater contract.
type OperatorTableUpdaterGlobalRootConfirmationThresholdUpdatedIterator struct {
	Event *OperatorTableUpdaterGlobalRootConfirmationThresholdUpdated // Event containing the contract specifics and raw log

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
func (it *OperatorTableUpdaterGlobalRootConfirmationThresholdUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorTableUpdaterGlobalRootConfirmationThresholdUpdated)
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
		it.Event = new(OperatorTableUpdaterGlobalRootConfirmationThresholdUpdated)
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
func (it *OperatorTableUpdaterGlobalRootConfirmationThresholdUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorTableUpdaterGlobalRootConfirmationThresholdUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorTableUpdaterGlobalRootConfirmationThresholdUpdated represents a GlobalRootConfirmationThresholdUpdated event raised by the OperatorTableUpdater contract.
type OperatorTableUpdaterGlobalRootConfirmationThresholdUpdated struct {
	Bps uint16
	Raw types.Log // Blockchain specific contextual infos
}

// FilterGlobalRootConfirmationThresholdUpdated is a free log retrieval operation binding the contract event 0xf5d1836df8fcd7c1e54047e94ac8773d2855395603e2ef9ba5f5f16905f22592.
//
// Solidity: event GlobalRootConfirmationThresholdUpdated(uint16 bps)
func (_OperatorTableUpdater *OperatorTableUpdaterFilterer) FilterGlobalRootConfirmationThresholdUpdated(opts *bind.FilterOpts) (*OperatorTableUpdaterGlobalRootConfirmationThresholdUpdatedIterator, error) {

	logs, sub, err := _OperatorTableUpdater.contract.FilterLogs(opts, "GlobalRootConfirmationThresholdUpdated")
	if err != nil {
		return nil, err
	}
	return &OperatorTableUpdaterGlobalRootConfirmationThresholdUpdatedIterator{contract: _OperatorTableUpdater.contract, event: "GlobalRootConfirmationThresholdUpdated", logs: logs, sub: sub}, nil
}

// WatchGlobalRootConfirmationThresholdUpdated is a free log subscription operation binding the contract event 0xf5d1836df8fcd7c1e54047e94ac8773d2855395603e2ef9ba5f5f16905f22592.
//
// Solidity: event GlobalRootConfirmationThresholdUpdated(uint16 bps)
func (_OperatorTableUpdater *OperatorTableUpdaterFilterer) WatchGlobalRootConfirmationThresholdUpdated(opts *bind.WatchOpts, sink chan<- *OperatorTableUpdaterGlobalRootConfirmationThresholdUpdated) (event.Subscription, error) {

	logs, sub, err := _OperatorTableUpdater.contract.WatchLogs(opts, "GlobalRootConfirmationThresholdUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorTableUpdaterGlobalRootConfirmationThresholdUpdated)
				if err := _OperatorTableUpdater.contract.UnpackLog(event, "GlobalRootConfirmationThresholdUpdated", log); err != nil {
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
func (_OperatorTableUpdater *OperatorTableUpdaterFilterer) ParseGlobalRootConfirmationThresholdUpdated(log types.Log) (*OperatorTableUpdaterGlobalRootConfirmationThresholdUpdated, error) {
	event := new(OperatorTableUpdaterGlobalRootConfirmationThresholdUpdated)
	if err := _OperatorTableUpdater.contract.UnpackLog(event, "GlobalRootConfirmationThresholdUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorTableUpdaterGlobalRootDisabledIterator is returned from FilterGlobalRootDisabled and is used to iterate over the raw logs and unpacked data for GlobalRootDisabled events raised by the OperatorTableUpdater contract.
type OperatorTableUpdaterGlobalRootDisabledIterator struct {
	Event *OperatorTableUpdaterGlobalRootDisabled // Event containing the contract specifics and raw log

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
func (it *OperatorTableUpdaterGlobalRootDisabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorTableUpdaterGlobalRootDisabled)
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
		it.Event = new(OperatorTableUpdaterGlobalRootDisabled)
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
func (it *OperatorTableUpdaterGlobalRootDisabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorTableUpdaterGlobalRootDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorTableUpdaterGlobalRootDisabled represents a GlobalRootDisabled event raised by the OperatorTableUpdater contract.
type OperatorTableUpdaterGlobalRootDisabled struct {
	GlobalTableRoot [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterGlobalRootDisabled is a free log retrieval operation binding the contract event 0x8bd43de1250f58fe6ec9a78671a8b78dba70f0018656d157a3aeaabec389df34.
//
// Solidity: event GlobalRootDisabled(bytes32 indexed globalTableRoot)
func (_OperatorTableUpdater *OperatorTableUpdaterFilterer) FilterGlobalRootDisabled(opts *bind.FilterOpts, globalTableRoot [][32]byte) (*OperatorTableUpdaterGlobalRootDisabledIterator, error) {

	var globalTableRootRule []interface{}
	for _, globalTableRootItem := range globalTableRoot {
		globalTableRootRule = append(globalTableRootRule, globalTableRootItem)
	}

	logs, sub, err := _OperatorTableUpdater.contract.FilterLogs(opts, "GlobalRootDisabled", globalTableRootRule)
	if err != nil {
		return nil, err
	}
	return &OperatorTableUpdaterGlobalRootDisabledIterator{contract: _OperatorTableUpdater.contract, event: "GlobalRootDisabled", logs: logs, sub: sub}, nil
}

// WatchGlobalRootDisabled is a free log subscription operation binding the contract event 0x8bd43de1250f58fe6ec9a78671a8b78dba70f0018656d157a3aeaabec389df34.
//
// Solidity: event GlobalRootDisabled(bytes32 indexed globalTableRoot)
func (_OperatorTableUpdater *OperatorTableUpdaterFilterer) WatchGlobalRootDisabled(opts *bind.WatchOpts, sink chan<- *OperatorTableUpdaterGlobalRootDisabled, globalTableRoot [][32]byte) (event.Subscription, error) {

	var globalTableRootRule []interface{}
	for _, globalTableRootItem := range globalTableRoot {
		globalTableRootRule = append(globalTableRootRule, globalTableRootItem)
	}

	logs, sub, err := _OperatorTableUpdater.contract.WatchLogs(opts, "GlobalRootDisabled", globalTableRootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorTableUpdaterGlobalRootDisabled)
				if err := _OperatorTableUpdater.contract.UnpackLog(event, "GlobalRootDisabled", log); err != nil {
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
func (_OperatorTableUpdater *OperatorTableUpdaterFilterer) ParseGlobalRootDisabled(log types.Log) (*OperatorTableUpdaterGlobalRootDisabled, error) {
	event := new(OperatorTableUpdaterGlobalRootDisabled)
	if err := _OperatorTableUpdater.contract.UnpackLog(event, "GlobalRootDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorTableUpdaterInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the OperatorTableUpdater contract.
type OperatorTableUpdaterInitializedIterator struct {
	Event *OperatorTableUpdaterInitialized // Event containing the contract specifics and raw log

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
func (it *OperatorTableUpdaterInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorTableUpdaterInitialized)
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
		it.Event = new(OperatorTableUpdaterInitialized)
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
func (it *OperatorTableUpdaterInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorTableUpdaterInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorTableUpdaterInitialized represents a Initialized event raised by the OperatorTableUpdater contract.
type OperatorTableUpdaterInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_OperatorTableUpdater *OperatorTableUpdaterFilterer) FilterInitialized(opts *bind.FilterOpts) (*OperatorTableUpdaterInitializedIterator, error) {

	logs, sub, err := _OperatorTableUpdater.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &OperatorTableUpdaterInitializedIterator{contract: _OperatorTableUpdater.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_OperatorTableUpdater *OperatorTableUpdaterFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *OperatorTableUpdaterInitialized) (event.Subscription, error) {

	logs, sub, err := _OperatorTableUpdater.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorTableUpdaterInitialized)
				if err := _OperatorTableUpdater.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_OperatorTableUpdater *OperatorTableUpdaterFilterer) ParseInitialized(log types.Log) (*OperatorTableUpdaterInitialized, error) {
	event := new(OperatorTableUpdaterInitialized)
	if err := _OperatorTableUpdater.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorTableUpdaterNewGlobalTableRootIterator is returned from FilterNewGlobalTableRoot and is used to iterate over the raw logs and unpacked data for NewGlobalTableRoot events raised by the OperatorTableUpdater contract.
type OperatorTableUpdaterNewGlobalTableRootIterator struct {
	Event *OperatorTableUpdaterNewGlobalTableRoot // Event containing the contract specifics and raw log

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
func (it *OperatorTableUpdaterNewGlobalTableRootIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorTableUpdaterNewGlobalTableRoot)
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
		it.Event = new(OperatorTableUpdaterNewGlobalTableRoot)
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
func (it *OperatorTableUpdaterNewGlobalTableRootIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorTableUpdaterNewGlobalTableRootIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorTableUpdaterNewGlobalTableRoot represents a NewGlobalTableRoot event raised by the OperatorTableUpdater contract.
type OperatorTableUpdaterNewGlobalTableRoot struct {
	ReferenceTimestamp uint32
	GlobalTableRoot    [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterNewGlobalTableRoot is a free log retrieval operation binding the contract event 0x010dcbe0d1e019c93357711f7bb6287d543b7ff7de74f29df3fb5ecceec8d369.
//
// Solidity: event NewGlobalTableRoot(uint32 indexed referenceTimestamp, bytes32 indexed globalTableRoot)
func (_OperatorTableUpdater *OperatorTableUpdaterFilterer) FilterNewGlobalTableRoot(opts *bind.FilterOpts, referenceTimestamp []uint32, globalTableRoot [][32]byte) (*OperatorTableUpdaterNewGlobalTableRootIterator, error) {

	var referenceTimestampRule []interface{}
	for _, referenceTimestampItem := range referenceTimestamp {
		referenceTimestampRule = append(referenceTimestampRule, referenceTimestampItem)
	}
	var globalTableRootRule []interface{}
	for _, globalTableRootItem := range globalTableRoot {
		globalTableRootRule = append(globalTableRootRule, globalTableRootItem)
	}

	logs, sub, err := _OperatorTableUpdater.contract.FilterLogs(opts, "NewGlobalTableRoot", referenceTimestampRule, globalTableRootRule)
	if err != nil {
		return nil, err
	}
	return &OperatorTableUpdaterNewGlobalTableRootIterator{contract: _OperatorTableUpdater.contract, event: "NewGlobalTableRoot", logs: logs, sub: sub}, nil
}

// WatchNewGlobalTableRoot is a free log subscription operation binding the contract event 0x010dcbe0d1e019c93357711f7bb6287d543b7ff7de74f29df3fb5ecceec8d369.
//
// Solidity: event NewGlobalTableRoot(uint32 indexed referenceTimestamp, bytes32 indexed globalTableRoot)
func (_OperatorTableUpdater *OperatorTableUpdaterFilterer) WatchNewGlobalTableRoot(opts *bind.WatchOpts, sink chan<- *OperatorTableUpdaterNewGlobalTableRoot, referenceTimestamp []uint32, globalTableRoot [][32]byte) (event.Subscription, error) {

	var referenceTimestampRule []interface{}
	for _, referenceTimestampItem := range referenceTimestamp {
		referenceTimestampRule = append(referenceTimestampRule, referenceTimestampItem)
	}
	var globalTableRootRule []interface{}
	for _, globalTableRootItem := range globalTableRoot {
		globalTableRootRule = append(globalTableRootRule, globalTableRootItem)
	}

	logs, sub, err := _OperatorTableUpdater.contract.WatchLogs(opts, "NewGlobalTableRoot", referenceTimestampRule, globalTableRootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorTableUpdaterNewGlobalTableRoot)
				if err := _OperatorTableUpdater.contract.UnpackLog(event, "NewGlobalTableRoot", log); err != nil {
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
func (_OperatorTableUpdater *OperatorTableUpdaterFilterer) ParseNewGlobalTableRoot(log types.Log) (*OperatorTableUpdaterNewGlobalTableRoot, error) {
	event := new(OperatorTableUpdaterNewGlobalTableRoot)
	if err := _OperatorTableUpdater.contract.UnpackLog(event, "NewGlobalTableRoot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorTableUpdaterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the OperatorTableUpdater contract.
type OperatorTableUpdaterOwnershipTransferredIterator struct {
	Event *OperatorTableUpdaterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OperatorTableUpdaterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorTableUpdaterOwnershipTransferred)
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
		it.Event = new(OperatorTableUpdaterOwnershipTransferred)
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
func (it *OperatorTableUpdaterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorTableUpdaterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorTableUpdaterOwnershipTransferred represents a OwnershipTransferred event raised by the OperatorTableUpdater contract.
type OperatorTableUpdaterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OperatorTableUpdater *OperatorTableUpdaterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OperatorTableUpdaterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OperatorTableUpdater.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OperatorTableUpdaterOwnershipTransferredIterator{contract: _OperatorTableUpdater.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OperatorTableUpdater *OperatorTableUpdaterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OperatorTableUpdaterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OperatorTableUpdater.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorTableUpdaterOwnershipTransferred)
				if err := _OperatorTableUpdater.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OperatorTableUpdater *OperatorTableUpdaterFilterer) ParseOwnershipTransferred(log types.Log) (*OperatorTableUpdaterOwnershipTransferred, error) {
	event := new(OperatorTableUpdaterOwnershipTransferred)
	if err := _OperatorTableUpdater.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorTableUpdaterPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the OperatorTableUpdater contract.
type OperatorTableUpdaterPausedIterator struct {
	Event *OperatorTableUpdaterPaused // Event containing the contract specifics and raw log

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
func (it *OperatorTableUpdaterPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorTableUpdaterPaused)
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
		it.Event = new(OperatorTableUpdaterPaused)
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
func (it *OperatorTableUpdaterPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorTableUpdaterPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorTableUpdaterPaused represents a Paused event raised by the OperatorTableUpdater contract.
type OperatorTableUpdaterPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_OperatorTableUpdater *OperatorTableUpdaterFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*OperatorTableUpdaterPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _OperatorTableUpdater.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &OperatorTableUpdaterPausedIterator{contract: _OperatorTableUpdater.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_OperatorTableUpdater *OperatorTableUpdaterFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *OperatorTableUpdaterPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _OperatorTableUpdater.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorTableUpdaterPaused)
				if err := _OperatorTableUpdater.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_OperatorTableUpdater *OperatorTableUpdaterFilterer) ParsePaused(log types.Log) (*OperatorTableUpdaterPaused, error) {
	event := new(OperatorTableUpdaterPaused)
	if err := _OperatorTableUpdater.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorTableUpdaterUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the OperatorTableUpdater contract.
type OperatorTableUpdaterUnpausedIterator struct {
	Event *OperatorTableUpdaterUnpaused // Event containing the contract specifics and raw log

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
func (it *OperatorTableUpdaterUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorTableUpdaterUnpaused)
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
		it.Event = new(OperatorTableUpdaterUnpaused)
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
func (it *OperatorTableUpdaterUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorTableUpdaterUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorTableUpdaterUnpaused represents a Unpaused event raised by the OperatorTableUpdater contract.
type OperatorTableUpdaterUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_OperatorTableUpdater *OperatorTableUpdaterFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*OperatorTableUpdaterUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _OperatorTableUpdater.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &OperatorTableUpdaterUnpausedIterator{contract: _OperatorTableUpdater.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_OperatorTableUpdater *OperatorTableUpdaterFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *OperatorTableUpdaterUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _OperatorTableUpdater.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorTableUpdaterUnpaused)
				if err := _OperatorTableUpdater.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_OperatorTableUpdater *OperatorTableUpdaterFilterer) ParseUnpaused(log types.Log) (*OperatorTableUpdaterUnpaused, error) {
	event := new(OperatorTableUpdaterUnpaused)
	if err := _OperatorTableUpdater.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
