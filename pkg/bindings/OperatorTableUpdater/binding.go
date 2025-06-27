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

// OperatorTableUpdaterMetaData contains all meta data concerning the OperatorTableUpdater contract.
var OperatorTableUpdaterMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_bn254CertificateVerifier\",\"type\":\"address\",\"internalType\":\"contractIBN254CertificateVerifier\"},{\"name\":\"_ecdsaCertificateVerifier\",\"type\":\"address\",\"internalType\":\"contractIECDSACertificateVerifier\"},{\"name\":\"_version\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"GLOBAL_TABLE_ROOT_CERT_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_BPS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bn254CertificateVerifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBN254CertificateVerifier\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"confirmGlobalTableRoot\",\"inputs\":[{\"name\":\"globalTableRootCert\",\"type\":\"tuple\",\"internalType\":\"structIBN254CertificateVerifierTypes.BN254Certificate\",\"components\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"nonSignerWitnesses\",\"type\":\"tuple[]\",\"internalType\":\"structIBN254CertificateVerifierTypes.BN254OperatorInfoWitness[]\",\"components\":[{\"name\":\"operatorIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorInfoProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"operatorInfo\",\"type\":\"tuple\",\"internalType\":\"structIBN254TableCalculatorTypes.BN254OperatorInfo\",\"components\":[{\"name\":\"pubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}]}]},{\"name\":\"globalTableRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"disableRoot\",\"inputs\":[{\"name\":\"globalTableRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ecdsaCertificateVerifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIECDSACertificateVerifier\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCertificateVerifier\",\"inputs\":[{\"name\":\"curveType\",\"type\":\"uint8\",\"internalType\":\"enumIKeyRegistrarTypes.CurveType\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentGlobalTableRoot\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getGlobalConfirmerSetReferenceTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getGlobalRootConfirmerSet\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getGlobalTableRootByTimestamp\",\"inputs\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getGlobalTableUpdateMessageHash\",\"inputs\":[{\"name\":\"globalTableRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getLatestReferenceBlockNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLatestReferenceTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getReferenceBlockNumberByTimestamp\",\"inputs\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getReferenceTimestampByBlockNumber\",\"inputs\":[{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"globalRootConfirmationThreshold\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_globalRootConfirmerSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"_globalRootConfirmationThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"globalRootConfirmerSetInfo\",\"type\":\"tuple\",\"internalType\":\"structIBN254TableCalculatorTypes.BN254OperatorSetInfo\",\"components\":[{\"name\":\"operatorInfoTreeRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"numOperators\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"aggregatePubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"totalWeights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]},{\"name\":\"globalRootConfirmerSetConfig\",\"type\":\"tuple\",\"internalType\":\"structICrossChainRegistryTypes.OperatorSetConfig\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxStalenessPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isRootValid\",\"inputs\":[{\"name\":\"globalTableRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRootValidByTimestamp\",\"inputs\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGlobalRootConfirmationThreshold\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGlobalRootConfirmerSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateGlobalRootConfirmerSet\",\"inputs\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"globalRootConfirmerSetInfo\",\"type\":\"tuple\",\"internalType\":\"structIBN254TableCalculatorTypes.BN254OperatorSetInfo\",\"components\":[{\"name\":\"operatorInfoTreeRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"numOperators\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"aggregatePubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"totalWeights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]},{\"name\":\"globalRootConfirmerSetConfig\",\"type\":\"tuple\",\"internalType\":\"structICrossChainRegistryTypes.OperatorSetConfig\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxStalenessPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOperatorTable\",\"inputs\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"globalTableRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"operatorSetIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"operatorTableBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"GlobalRootConfirmationThresholdUpdated\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GlobalRootConfirmerSetUpdated\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GlobalRootDisabled\",\"inputs\":[{\"name\":\"globalTableRoot\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewGlobalTableRoot\",\"inputs\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"globalTableRoot\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CertificateInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"GlobalTableRootInFuture\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"GlobalTableRootStale\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidConfirmationThreshold\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCurveType\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidGlobalTableRoot\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidMessageHash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperatorSetProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProofLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRoot\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignatureLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"TableUpdateForPastTimestamp\",\"inputs\":[]}]",
	Bin: "0x60e060405234801561000f575f5ffd5b5060405161231638038061231683398101604081905261002e91610188565b6001600160a01b03808416608052821660a0528061004b8161005f565b60c052506100576100a5565b5050506102b9565b5f5f829050601f81511115610092578260405163305a27a960e01b8152600401610089919061025e565b60405180910390fd5b805161009d82610293565b179392505050565b5f54610100900460ff161561010c5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b6064820152608401610089565b5f5460ff9081161461015b575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b0381168114610171575f5ffd5b50565b634e487b7160e01b5f52604160045260245ffd5b5f5f5f6060848603121561019a575f5ffd5b83516101a58161015d565b60208501519093506101b68161015d565b60408501519092506001600160401b038111156101d1575f5ffd5b8401601f810186136101e1575f5ffd5b80516001600160401b038111156101fa576101fa610174565b604051601f8201601f19908116603f011681016001600160401b038111828210171561022857610228610174565b60405281815282820160200188101561023f575f5ffd5b8160208401602083015e5f602083830101528093505050509250925092565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b805160208083015191908110156102b3575f198160200360031b1b821691505b50919050565b60805160a05160c0516120026103145f395f61073601525f818161042b015281816107b701526109d501525f81816104520152818161068c015281816107770152818161093001528181610c040152610ed801526120025ff3fe608060405234801561000f575f5ffd5b50600436106101bb575f3560e01c806364e1df84116100f3578063c252aa2211610093578063c5916a391161006e578063c5916a39146104bb578063eaaed9d5146104e0578063f2fde38b146104f3578063fd967f4714610506575f5ffd5b8063c252aa2214610474578063c3621f0a14610495578063c3be1e33146104a8575f5ffd5b80638da5cb5b116100ce5780638da5cb5b146104025780639ea9477814610413578063ad0f958214610426578063b8c143061461044d575f5ffd5b806364e1df841461039a5780636f728c50146103cf578063715018a6146103fa575f5ffd5b806328522d791161015e5780633ef6cd7a116101395780633ef6cd7a146103025780634624e6a314610329578063462828891461033d57806354fd4d5014610385575f5ffd5b806328522d791461027f57806330ef41b4146102ab57806331a599d2146102dd575f5ffd5b8063193b79f311610199578063193b79f3146102095780631ab78d90146102315780632370356c1461024457806323b7b5b214610257575f5ffd5b8063021ab442146101bf5780630371406e146101d45780630f3f8edd146101e7575b5f5ffd5b6101d26101cd366004611264565b61050f565b005b6101d26101e23660046112f4565b61065f565b6101ef610673565b60405163ffffffff90911681526020015b60405180910390f35b6101ef610217366004611315565b63ffffffff9081165f908152606960205260409020541690565b6101d261023f366004611330565b610706565b6101d261025236600461138c565b61071e565b6101ef610265366004611315565b63ffffffff9081165f908152606860205260409020541690565b60655462010000900463ffffffff165f908152606760205260409020545b604051908152602001610200565b6102cd6102b93660046113a5565b5f908152606a602052604090205460ff1690565b6040519015158152602001610200565b60655462010000900463ffffffff9081165f90815260686020526040902054166101ef565b61029d7f4491f5ee91595f938885ef73c9a1fa8a6d14ff9b9dab4aa24b8802bbb9bfc1cc81565b60655462010000900463ffffffff166101ef565b6040805180820182525f80825260209182015281518083019092526066546001600160a01b0381168352600160a01b900463ffffffff169082015260405161020091906113da565b61038d61072f565b60405161020091906113e8565b6102cd6103a8366004611315565b63ffffffff165f908152606760209081526040808320548352606a90915290205460ff1690565b6103e26103dd36600461142b565b61075a565b6040516001600160a01b039091168152602001610200565b6101d26107f9565b6033546001600160a01b03166103e2565b6101d2610421366004611488565b61080c565b6103e27f000000000000000000000000000000000000000000000000000000000000000081565b6103e27f000000000000000000000000000000000000000000000000000000000000000081565b6065546104829061ffff1681565b60405161ffff9091168152602001610200565b6101d26104a33660046113a5565b610a3a565b61029d6104b6366004611523565b610aaf565b61029d6104c9366004611315565b63ffffffff165f9081526067602052604090205490565b6101d26104ee366004611562565b610b17565b6101d26105013660046115ce565b610d49565b61048261271081565b5f54610100900460ff161580801561052d57505f54600160ff909116105b806105465750303b15801561054657505f5460ff166001145b6105ae5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b5f805460ff1916600117905580156105cf575f805461ff0019166101001790555b6105d887610dbb565b6105e186610e0c565b6105ea85610e56565b6105f5848484610ec1565b6065805465ffffffff000019166201000063ffffffff8716021790558015610656575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050505050565b610667610f3d565b61067081610e0c565b50565b604051635ddb9b5b60e01b81525f906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690635ddb9b5b906106c2906066906004016115e9565b602060405180830381865afa1580156106dd573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107019190611610565b905090565b61070e610f3d565b610719838383610ec1565b505050565b610726610f3d565b61067081610e56565b60606107017f0000000000000000000000000000000000000000000000000000000000000000610f97565b5f600282600281111561076f5761076f61162b565b0361079b57507f0000000000000000000000000000000000000000000000000000000000000000919050565b60018260028111156107af576107af61162b565b036107db57507f0000000000000000000000000000000000000000000000000000000000000000919050565b60405163fdea7c0960e01b815260040160405180910390fd5b919050565b610801610f3d565b61080a5f610dbb565b565b5f5f5f5f61081a8686610fd4565b5f8e8152606a60205260409020549397509195509350915060ff166108525760405163504570e360e01b815260040160405180910390fd5b61085b8361075a565b6001600160a01b0316635ddb9b5b856040518263ffffffff1660e01b815260040161088691906113da565b602060405180830381865afa1580156108a1573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108c59190611610565b63ffffffff168b63ffffffff16116108f05760405163207617df60e01b815260040160405180910390fd5b6109158b8b8b8b8b8b8b60405161090892919061163f565b604051809103902061101b565b60028360028111156109295761092961162b565b036109ba577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316636738c40b858d610968856110bc565b866040518563ffffffff1660e01b81526004016109889493929190611688565b5f604051808303815f87803b15801561099f575f5ffd5b505af11580156109b1573d5f5f3e3d5ffd5b50505050610a2d565b60018360028111156109ce576109ce61162b565b036107db577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166356d482f5858d610a0d856110de565b866040518563ffffffff1660e01b81526004016109889493929190611704565b5050505050505050505050565b610a42610f3d565b5f818152606a602052604090205460ff16610a705760405163504570e360e01b815260040160405180910390fd5b5f818152606a6020526040808220805460ff191690555182917f8bd43de1250f58fe6ec9a78671a8b78dba70f0018656d157a3aeaabec389df3491a250565b604080517f4491f5ee91595f938885ef73c9a1fa8a6d14ff9b9dab4aa24b8802bbb9bfc1cc602082015290810184905263ffffffff8084166060830152821660808201525f9060a0016040516020818303038152906040528051906020012090509392505050565b428263ffffffff161115610b3e57604051635a119db560e11b815260040160405180910390fd5b60655463ffffffff62010000909104811690831611610b705760405163037fa86b60e31b815260040160405180910390fd5b610b7b838383610aaf565b846020013514610b9e57604051638b56642d60e01b815260040160405180910390fd5b6040805160018082528183019092525f91602080830190803683375050606554825192935061ffff16918391505f90610bd957610bd96117bd565b61ffff90921660209283029190910190910152604051625f5e5d60e21b81525f906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063017d797490610c3e906066908a9087906004016118ef565b6020604051808303815f875af1158015610c5a573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610c7e9190611a7c565b905080610c9e57604051633042041f60e21b815260040160405180910390fd5b6065805463ffffffff80871662010000810265ffffffff000019909316929092179092555f818152606860209081526040808320805495891663ffffffff1996871681179091558352606982528083208054909516841790945582825260678152838220899055888252606a9052828120805460ff19166001179055915187927f010dcbe0d1e019c93357711f7bb6287d543b7ff7de74f29df3fb5ecceec8d36991a3505050505050565b610d51610f3d565b6001600160a01b038116610db65760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016105a5565b610670815b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b806066610e198282611a9b565b9050507f20100394950e66014c25009b45d12b675210a6e7a002044a0e3de6544e3c4b3781604051610e4b9190611b2c565b60405180910390a150565b61271061ffff82161115610e7d576040516307336f0360e11b815260040160405180910390fd5b6065805461ffff191661ffff83169081179091556040519081527ff5d1836df8fcd7c1e54047e94ac8773d2855395603e2ef9ba5f5f16905f2259290602001610e4b565b604051636738c40b60e01b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690636738c40b90610f1490606690879087908790600401611b3a565b5f604051808303815f87803b158015610f2b575f5ffd5b505af1158015610656573d5f5f3e3d5ffd5b6033546001600160a01b0316331461080a5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016105a5565b60605f610fa3836110f4565b6040805160208082528183019092529192505f91906020820181803683375050509182525060208101929092525090565b604080518082019091525f8082526020820152604080518082019091525f8082526020820181905290606061100b85870187611c8d565b9299919850965090945092505050565b63ffffffff86165f90815260676020526040902054851461104f5760405163639d09b560e11b815260040160405180910390fd5b61109783838080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152508992508591505063ffffffff881661111b565b6110b45760405163afa42ca760e01b815260040160405180910390fd5b505050505050565b6110c46111c9565b818060200190518101906110d89190611dde565b92915050565b6060818060200190518101906110d89190611e8d565b5f60ff8216601f8111156110d857604051632cd44ac360e21b815260040160405180910390fd5b5f83611128868585611132565b1495945050505050565b5f602084516111419190611f8e565b1561115f576040516313717da960e21b815260040160405180910390fd5b8260205b855181116111c057611176600285611f8e565b5f0361119757815f528086015160205260405f2091506002840493506111ae565b808601515f528160205260405f2091506002840493505b6111b9602082611fad565b9050611163565b50949350505050565b60405180608001604052805f81526020015f81526020016111fb60405180604001604052805f81526020015f81525090565b8152602001606081525090565b6001600160a01b0381168114610670575f5ffd5b5f6040828403121561122c575f5ffd5b50919050565b803561ffff811681146107f4575f5ffd5b63ffffffff81168114610670575f5ffd5b5f60a0828403121561122c575f5ffd5b5f5f5f5f5f5f610100878903121561127a575f5ffd5b863561128581611208565b9550611294886020890161121c565b94506112a260608801611232565b935060808701356112b281611243565b925060a08701356001600160401b038111156112cc575f5ffd5b6112d889828a01611254565b9250506112e88860c0890161121c565b90509295509295509295565b5f60408284031215611304575f5ffd5b61130e838361121c565b9392505050565b5f60208284031215611325575f5ffd5b813561130e81611243565b5f5f5f60808486031215611342575f5ffd5b833561134d81611243565b925060208401356001600160401b03811115611367575f5ffd5b61137386828701611254565b925050611383856040860161121c565b90509250925092565b5f6020828403121561139c575f5ffd5b61130e82611232565b5f602082840312156113b5575f5ffd5b5035919050565b80516001600160a01b0316825260209081015163ffffffff16910152565b604081016110d882846113bc565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b8035600381106107f4575f5ffd5b5f6020828403121561143b575f5ffd5b61130e8261141d565b5f5f83601f840112611454575f5ffd5b5081356001600160401b0381111561146a575f5ffd5b602083019150836020828501011115611481575f5ffd5b9250929050565b5f5f5f5f5f5f5f60a0888a03121561149e575f5ffd5b87356114a981611243565b96506020880135955060408801356114c081611243565b945060608801356001600160401b038111156114da575f5ffd5b6114e68a828b01611444565b90955093505060808801356001600160401b03811115611504575f5ffd5b6115108a828b01611444565b989b979a50959850939692959293505050565b5f5f5f60608486031215611535575f5ffd5b83359250602084013561154781611243565b9150604084013561155781611243565b809150509250925092565b5f5f5f5f60808587031215611575575f5ffd5b84356001600160401b0381111561158a575f5ffd5b8501610120818803121561159c575f5ffd5b93506020850135925060408501356115b381611243565b915060608501356115c381611243565b939692955090935050565b5f602082840312156115de575f5ffd5b813561130e81611208565b604081016110d88284546001600160a01b038116825260a01c63ffffffff16602090910152565b5f60208284031215611620575f5ffd5b815161130e81611243565b634e487b7160e01b5f52602160045260245ffd5b818382375f9101908152919050565b5f8151808452602084019350602083015f5b8281101561167e578151865260209586019590910190600101611660565b5093949350505050565b61169281866113bc565b63ffffffff8416604082015260c06060820152825160c0820152602083015160e08201525f60408401518051610100840152602081015161012084015250606084015160a06101408401526116eb61016084018261164e565b9150506116fb60808301846113bc565b95945050505050565b5f60c0820161171383886113bc565b63ffffffff8616604084015260c0606084015280855180835260e08501915060e08160051b8601019250602087015f5b828110156117945786850360df19018452815180516001600160a01b0316865260209081015160409187018290529061177e9087018261164e565b9550506020938401939190910190600101611743565b50505050809150506116fb60808301846113bc565b634e487b7160e01b5f52604160045260245ffd5b634e487b7160e01b5f52603260045260245ffd5b5f5f8335601e198436030181126117e6575f5ffd5b83016020810192503590506001600160401b03811115611804575f5ffd5b8060051b3603821315611481575f5ffd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b5f8235605e19833603018112611851575f5ffd5b90910192915050565b8183525f6001600160fb1b03831115611871575f5ffd5b8260051b80836020870137939093016020019392505050565b80358252602080820135908301525f6118a660408301836117d1565b606060408601526116fb60608601828461185a565b5f8151808452602084019350602083015f5b8281101561167e57815161ffff168652602095860195909101906001016118cd565b6119128185546001600160a01b038116825260a01c63ffffffff16602090910152565b608060408201525f6101a08201843561192a81611243565b63ffffffff166080840152602085013560a0840152604085013560c0840152606085013560e0840152604060808601610100850137604060c086016101408501376119796101008601866117d1565b610120610180860152828184526101c0860190506101c08260051b8701019350825f5b83811015611a5b578786036101bf190183526119b8828661183d565b80356119c381611243565b63ffffffff168752602081013536829003601e190181126119e2575f5ffd5b81016020810190356001600160401b038111156119fd575f5ffd5b803603821315611a0b575f5ffd5b606060208a0152611a2060608a018284611815565b915050611a30604083018361183d565b91508781036040890152611a44818361188a565b97505050602092830192919091019060010161199c565b50505050508281036060840152611a7281856118bb565b9695505050505050565b5f60208284031215611a8c575f5ffd5b8151801515811461130e575f5ffd5b8135611aa681611208565b81546001600160a01b031981166001600160a01b039290921691821783556020840135611ad281611243565b6001600160c01b03199190911690911760a09190911b63ffffffff60a01b1617905550565b8035611b0281611208565b6001600160a01b031682526020810135611b1b81611243565b63ffffffff81166020840152505050565b604081016110d88284611af7565b611b5d8186546001600160a01b038116825260a01c63ffffffff16602090910152565b63ffffffff841660408281019190915260c06060808401829052853591840191909152602085013560e0840152908401356101008301528301356101208201525f611bab60808501856117d1565b60a0610140850152611bc26101608501828461185a565b925050506116fb6080830184611af7565b604080519081016001600160401b0381118282101715611bf557611bf56117a9565b60405290565b604051608081016001600160401b0381118282101715611bf557611bf56117a9565b604051601f8201601f191681016001600160401b0381118282101715611c4557611c456117a9565b604052919050565b5f60408284031215611c5d575f5ffd5b611c65611bd3565b90508135611c7281611208565b81526020820135611c8281611243565b602082015292915050565b5f5f5f5f60c08587031215611ca0575f5ffd5b611caa8686611c4d565b9350611cb86040860161141d565b9250611cc78660608701611c4d565b915060a08501356001600160401b03811115611ce1575f5ffd5b8501601f81018713611cf1575f5ffd5b80356001600160401b03811115611d0a57611d0a6117a9565b611d1d601f8201601f1916602001611c1d565b818152886020838501011115611d31575f5ffd5b816020840160208301375f6020838301015280935050505092959194509250565b5f6001600160401b03821115611d6a57611d6a6117a9565b5060051b60200190565b5f82601f830112611d83575f5ffd5b8151611d96611d9182611d52565b611c1d565b8082825260208201915060208360051b860101925085831115611db7575f5ffd5b602085015b83811015611dd4578051835260209283019201611dbc565b5095945050505050565b5f60208284031215611dee575f5ffd5b81516001600160401b03811115611e03575f5ffd5b820180840360a0811215611e15575f5ffd5b611e1d611bfb565b82518152602080840151908201526040603f1983011215611e3c575f5ffd5b611e44611bd3565b604084810151825260608501516020830152820152608083015191506001600160401b03821115611e73575f5ffd5b611e7f86838501611d74565b606082015295945050505050565b5f60208284031215611e9d575f5ffd5b81516001600160401b03811115611eb2575f5ffd5b8201601f81018413611ec2575f5ffd5b8051611ed0611d9182611d52565b8082825260208201915060208360051b850101925086831115611ef1575f5ffd5b602084015b83811015611f835780516001600160401b03811115611f13575f5ffd5b85016040818a03601f19011215611f28575f5ffd5b611f30611bd3565b6020820151611f3e81611208565b815260408201516001600160401b03811115611f58575f5ffd5b611f678b602083860101611d74565b6020830152508085525050602083019250602081019050611ef6565b509695505050505050565b5f82611fa857634e487b7160e01b5f52601260045260245ffd5b500690565b808201808211156110d857634e487b7160e01b5f52601160045260245ffdfea26469706673582212201633e49faf00f03a2a56870830f7f9b8bb992766b3faf552c112ca4a1c0d231964736f6c634300081b0033",
}

// OperatorTableUpdaterABI is the input ABI used to generate the binding from.
// Deprecated: Use OperatorTableUpdaterMetaData.ABI instead.
var OperatorTableUpdaterABI = OperatorTableUpdaterMetaData.ABI

// OperatorTableUpdaterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OperatorTableUpdaterMetaData.Bin instead.
var OperatorTableUpdaterBin = OperatorTableUpdaterMetaData.Bin

// DeployOperatorTableUpdater deploys a new Ethereum contract, binding an instance of OperatorTableUpdater to it.
func DeployOperatorTableUpdater(auth *bind.TransactOpts, backend bind.ContractBackend, _bn254CertificateVerifier common.Address, _ecdsaCertificateVerifier common.Address, _version string) (common.Address, *types.Transaction, *OperatorTableUpdater, error) {
	parsed, err := OperatorTableUpdaterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OperatorTableUpdaterBin), backend, _bn254CertificateVerifier, _ecdsaCertificateVerifier, _version)
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

// GetGlobalConfirmerSetReferenceTimestamp is a free data retrieval call binding the contract method 0x0f3f8edd.
//
// Solidity: function getGlobalConfirmerSetReferenceTimestamp() view returns(uint32)
func (_OperatorTableUpdater *OperatorTableUpdaterCaller) GetGlobalConfirmerSetReferenceTimestamp(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _OperatorTableUpdater.contract.Call(opts, &out, "getGlobalConfirmerSetReferenceTimestamp")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetGlobalConfirmerSetReferenceTimestamp is a free data retrieval call binding the contract method 0x0f3f8edd.
//
// Solidity: function getGlobalConfirmerSetReferenceTimestamp() view returns(uint32)
func (_OperatorTableUpdater *OperatorTableUpdaterSession) GetGlobalConfirmerSetReferenceTimestamp() (uint32, error) {
	return _OperatorTableUpdater.Contract.GetGlobalConfirmerSetReferenceTimestamp(&_OperatorTableUpdater.CallOpts)
}

// GetGlobalConfirmerSetReferenceTimestamp is a free data retrieval call binding the contract method 0x0f3f8edd.
//
// Solidity: function getGlobalConfirmerSetReferenceTimestamp() view returns(uint32)
func (_OperatorTableUpdater *OperatorTableUpdaterCallerSession) GetGlobalConfirmerSetReferenceTimestamp() (uint32, error) {
	return _OperatorTableUpdater.Contract.GetGlobalConfirmerSetReferenceTimestamp(&_OperatorTableUpdater.CallOpts)
}

// GetGlobalRootConfirmerSet is a free data retrieval call binding the contract method 0x46282889.
//
// Solidity: function getGlobalRootConfirmerSet() view returns((address,uint32))
func (_OperatorTableUpdater *OperatorTableUpdaterCaller) GetGlobalRootConfirmerSet(opts *bind.CallOpts) (OperatorSet, error) {
	var out []interface{}
	err := _OperatorTableUpdater.contract.Call(opts, &out, "getGlobalRootConfirmerSet")

	if err != nil {
		return *new(OperatorSet), err
	}

	out0 := *abi.ConvertType(out[0], new(OperatorSet)).(*OperatorSet)

	return out0, err

}

// GetGlobalRootConfirmerSet is a free data retrieval call binding the contract method 0x46282889.
//
// Solidity: function getGlobalRootConfirmerSet() view returns((address,uint32))
func (_OperatorTableUpdater *OperatorTableUpdaterSession) GetGlobalRootConfirmerSet() (OperatorSet, error) {
	return _OperatorTableUpdater.Contract.GetGlobalRootConfirmerSet(&_OperatorTableUpdater.CallOpts)
}

// GetGlobalRootConfirmerSet is a free data retrieval call binding the contract method 0x46282889.
//
// Solidity: function getGlobalRootConfirmerSet() view returns((address,uint32))
func (_OperatorTableUpdater *OperatorTableUpdaterCallerSession) GetGlobalRootConfirmerSet() (OperatorSet, error) {
	return _OperatorTableUpdater.Contract.GetGlobalRootConfirmerSet(&_OperatorTableUpdater.CallOpts)
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

// Initialize is a paid mutator transaction binding the contract method 0x021ab442.
//
// Solidity: function initialize(address owner, (address,uint32) _globalRootConfirmerSet, uint16 _globalRootConfirmationThreshold, uint32 referenceTimestamp, (bytes32,uint256,(uint256,uint256),uint256[]) globalRootConfirmerSetInfo, (address,uint32) globalRootConfirmerSetConfig) returns()
func (_OperatorTableUpdater *OperatorTableUpdaterTransactor) Initialize(opts *bind.TransactOpts, owner common.Address, _globalRootConfirmerSet OperatorSet, _globalRootConfirmationThreshold uint16, referenceTimestamp uint32, globalRootConfirmerSetInfo IBN254TableCalculatorTypesBN254OperatorSetInfo, globalRootConfirmerSetConfig ICrossChainRegistryTypesOperatorSetConfig) (*types.Transaction, error) {
	return _OperatorTableUpdater.contract.Transact(opts, "initialize", owner, _globalRootConfirmerSet, _globalRootConfirmationThreshold, referenceTimestamp, globalRootConfirmerSetInfo, globalRootConfirmerSetConfig)
}

// Initialize is a paid mutator transaction binding the contract method 0x021ab442.
//
// Solidity: function initialize(address owner, (address,uint32) _globalRootConfirmerSet, uint16 _globalRootConfirmationThreshold, uint32 referenceTimestamp, (bytes32,uint256,(uint256,uint256),uint256[]) globalRootConfirmerSetInfo, (address,uint32) globalRootConfirmerSetConfig) returns()
func (_OperatorTableUpdater *OperatorTableUpdaterSession) Initialize(owner common.Address, _globalRootConfirmerSet OperatorSet, _globalRootConfirmationThreshold uint16, referenceTimestamp uint32, globalRootConfirmerSetInfo IBN254TableCalculatorTypesBN254OperatorSetInfo, globalRootConfirmerSetConfig ICrossChainRegistryTypesOperatorSetConfig) (*types.Transaction, error) {
	return _OperatorTableUpdater.Contract.Initialize(&_OperatorTableUpdater.TransactOpts, owner, _globalRootConfirmerSet, _globalRootConfirmationThreshold, referenceTimestamp, globalRootConfirmerSetInfo, globalRootConfirmerSetConfig)
}

// Initialize is a paid mutator transaction binding the contract method 0x021ab442.
//
// Solidity: function initialize(address owner, (address,uint32) _globalRootConfirmerSet, uint16 _globalRootConfirmationThreshold, uint32 referenceTimestamp, (bytes32,uint256,(uint256,uint256),uint256[]) globalRootConfirmerSetInfo, (address,uint32) globalRootConfirmerSetConfig) returns()
func (_OperatorTableUpdater *OperatorTableUpdaterTransactorSession) Initialize(owner common.Address, _globalRootConfirmerSet OperatorSet, _globalRootConfirmationThreshold uint16, referenceTimestamp uint32, globalRootConfirmerSetInfo IBN254TableCalculatorTypesBN254OperatorSetInfo, globalRootConfirmerSetConfig ICrossChainRegistryTypesOperatorSetConfig) (*types.Transaction, error) {
	return _OperatorTableUpdater.Contract.Initialize(&_OperatorTableUpdater.TransactOpts, owner, _globalRootConfirmerSet, _globalRootConfirmationThreshold, referenceTimestamp, globalRootConfirmerSetInfo, globalRootConfirmerSetConfig)
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

// SetGlobalRootConfirmerSet is a paid mutator transaction binding the contract method 0x0371406e.
//
// Solidity: function setGlobalRootConfirmerSet((address,uint32) operatorSet) returns()
func (_OperatorTableUpdater *OperatorTableUpdaterTransactor) SetGlobalRootConfirmerSet(opts *bind.TransactOpts, operatorSet OperatorSet) (*types.Transaction, error) {
	return _OperatorTableUpdater.contract.Transact(opts, "setGlobalRootConfirmerSet", operatorSet)
}

// SetGlobalRootConfirmerSet is a paid mutator transaction binding the contract method 0x0371406e.
//
// Solidity: function setGlobalRootConfirmerSet((address,uint32) operatorSet) returns()
func (_OperatorTableUpdater *OperatorTableUpdaterSession) SetGlobalRootConfirmerSet(operatorSet OperatorSet) (*types.Transaction, error) {
	return _OperatorTableUpdater.Contract.SetGlobalRootConfirmerSet(&_OperatorTableUpdater.TransactOpts, operatorSet)
}

// SetGlobalRootConfirmerSet is a paid mutator transaction binding the contract method 0x0371406e.
//
// Solidity: function setGlobalRootConfirmerSet((address,uint32) operatorSet) returns()
func (_OperatorTableUpdater *OperatorTableUpdaterTransactorSession) SetGlobalRootConfirmerSet(operatorSet OperatorSet) (*types.Transaction, error) {
	return _OperatorTableUpdater.Contract.SetGlobalRootConfirmerSet(&_OperatorTableUpdater.TransactOpts, operatorSet)
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

// UpdateGlobalRootConfirmerSet is a paid mutator transaction binding the contract method 0x1ab78d90.
//
// Solidity: function updateGlobalRootConfirmerSet(uint32 referenceTimestamp, (bytes32,uint256,(uint256,uint256),uint256[]) globalRootConfirmerSetInfo, (address,uint32) globalRootConfirmerSetConfig) returns()
func (_OperatorTableUpdater *OperatorTableUpdaterTransactor) UpdateGlobalRootConfirmerSet(opts *bind.TransactOpts, referenceTimestamp uint32, globalRootConfirmerSetInfo IBN254TableCalculatorTypesBN254OperatorSetInfo, globalRootConfirmerSetConfig ICrossChainRegistryTypesOperatorSetConfig) (*types.Transaction, error) {
	return _OperatorTableUpdater.contract.Transact(opts, "updateGlobalRootConfirmerSet", referenceTimestamp, globalRootConfirmerSetInfo, globalRootConfirmerSetConfig)
}

// UpdateGlobalRootConfirmerSet is a paid mutator transaction binding the contract method 0x1ab78d90.
//
// Solidity: function updateGlobalRootConfirmerSet(uint32 referenceTimestamp, (bytes32,uint256,(uint256,uint256),uint256[]) globalRootConfirmerSetInfo, (address,uint32) globalRootConfirmerSetConfig) returns()
func (_OperatorTableUpdater *OperatorTableUpdaterSession) UpdateGlobalRootConfirmerSet(referenceTimestamp uint32, globalRootConfirmerSetInfo IBN254TableCalculatorTypesBN254OperatorSetInfo, globalRootConfirmerSetConfig ICrossChainRegistryTypesOperatorSetConfig) (*types.Transaction, error) {
	return _OperatorTableUpdater.Contract.UpdateGlobalRootConfirmerSet(&_OperatorTableUpdater.TransactOpts, referenceTimestamp, globalRootConfirmerSetInfo, globalRootConfirmerSetConfig)
}

// UpdateGlobalRootConfirmerSet is a paid mutator transaction binding the contract method 0x1ab78d90.
//
// Solidity: function updateGlobalRootConfirmerSet(uint32 referenceTimestamp, (bytes32,uint256,(uint256,uint256),uint256[]) globalRootConfirmerSetInfo, (address,uint32) globalRootConfirmerSetConfig) returns()
func (_OperatorTableUpdater *OperatorTableUpdaterTransactorSession) UpdateGlobalRootConfirmerSet(referenceTimestamp uint32, globalRootConfirmerSetInfo IBN254TableCalculatorTypesBN254OperatorSetInfo, globalRootConfirmerSetConfig ICrossChainRegistryTypesOperatorSetConfig) (*types.Transaction, error) {
	return _OperatorTableUpdater.Contract.UpdateGlobalRootConfirmerSet(&_OperatorTableUpdater.TransactOpts, referenceTimestamp, globalRootConfirmerSetInfo, globalRootConfirmerSetConfig)
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

// OperatorTableUpdaterGlobalRootConfirmerSetUpdatedIterator is returned from FilterGlobalRootConfirmerSetUpdated and is used to iterate over the raw logs and unpacked data for GlobalRootConfirmerSetUpdated events raised by the OperatorTableUpdater contract.
type OperatorTableUpdaterGlobalRootConfirmerSetUpdatedIterator struct {
	Event *OperatorTableUpdaterGlobalRootConfirmerSetUpdated // Event containing the contract specifics and raw log

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
func (it *OperatorTableUpdaterGlobalRootConfirmerSetUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorTableUpdaterGlobalRootConfirmerSetUpdated)
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
		it.Event = new(OperatorTableUpdaterGlobalRootConfirmerSetUpdated)
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
func (it *OperatorTableUpdaterGlobalRootConfirmerSetUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorTableUpdaterGlobalRootConfirmerSetUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorTableUpdaterGlobalRootConfirmerSetUpdated represents a GlobalRootConfirmerSetUpdated event raised by the OperatorTableUpdater contract.
type OperatorTableUpdaterGlobalRootConfirmerSetUpdated struct {
	OperatorSet OperatorSet
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterGlobalRootConfirmerSetUpdated is a free log retrieval operation binding the contract event 0x20100394950e66014c25009b45d12b675210a6e7a002044a0e3de6544e3c4b37.
//
// Solidity: event GlobalRootConfirmerSetUpdated((address,uint32) operatorSet)
func (_OperatorTableUpdater *OperatorTableUpdaterFilterer) FilterGlobalRootConfirmerSetUpdated(opts *bind.FilterOpts) (*OperatorTableUpdaterGlobalRootConfirmerSetUpdatedIterator, error) {

	logs, sub, err := _OperatorTableUpdater.contract.FilterLogs(opts, "GlobalRootConfirmerSetUpdated")
	if err != nil {
		return nil, err
	}
	return &OperatorTableUpdaterGlobalRootConfirmerSetUpdatedIterator{contract: _OperatorTableUpdater.contract, event: "GlobalRootConfirmerSetUpdated", logs: logs, sub: sub}, nil
}

// WatchGlobalRootConfirmerSetUpdated is a free log subscription operation binding the contract event 0x20100394950e66014c25009b45d12b675210a6e7a002044a0e3de6544e3c4b37.
//
// Solidity: event GlobalRootConfirmerSetUpdated((address,uint32) operatorSet)
func (_OperatorTableUpdater *OperatorTableUpdaterFilterer) WatchGlobalRootConfirmerSetUpdated(opts *bind.WatchOpts, sink chan<- *OperatorTableUpdaterGlobalRootConfirmerSetUpdated) (event.Subscription, error) {

	logs, sub, err := _OperatorTableUpdater.contract.WatchLogs(opts, "GlobalRootConfirmerSetUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorTableUpdaterGlobalRootConfirmerSetUpdated)
				if err := _OperatorTableUpdater.contract.UnpackLog(event, "GlobalRootConfirmerSetUpdated", log); err != nil {
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
func (_OperatorTableUpdater *OperatorTableUpdaterFilterer) ParseGlobalRootConfirmerSetUpdated(log types.Log) (*OperatorTableUpdaterGlobalRootConfirmerSetUpdated, error) {
	event := new(OperatorTableUpdaterGlobalRootConfirmerSetUpdated)
	if err := _OperatorTableUpdater.contract.UnpackLog(event, "GlobalRootConfirmerSetUpdated", log); err != nil {
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
