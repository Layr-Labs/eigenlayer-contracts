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
<<<<<<< HEAD
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_bn254CertificateVerifier\",\"type\":\"address\",\"internalType\":\"contractIBN254CertificateVerifier\"},{\"name\":\"_ecdsaCertificateVerifier\",\"type\":\"address\",\"internalType\":\"contractIECDSACertificateVerifier\"},{\"name\":\"_version\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"GLOBAL_TABLE_ROOT_CERT_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"INITIAL_GLOBAL_TABLE_ROOT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_BPS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bn254CertificateVerifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBN254CertificateVerifier\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"confirmGlobalTableRoot\",\"inputs\":[{\"name\":\"globalTableRootCert\",\"type\":\"tuple\",\"internalType\":\"structIBN254CertificateVerifierTypes.BN254Certificate\",\"components\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"nonSignerWitnesses\",\"type\":\"tuple[]\",\"internalType\":\"structIBN254CertificateVerifierTypes.BN254OperatorInfoWitness[]\",\"components\":[{\"name\":\"operatorIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorInfoProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"operatorInfo\",\"type\":\"tuple\",\"internalType\":\"structIOperatorTableCalculatorTypes.BN254OperatorInfo\",\"components\":[{\"name\":\"pubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}]}]},{\"name\":\"globalTableRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"disableRoot\",\"inputs\":[{\"name\":\"globalTableRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ecdsaCertificateVerifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIECDSACertificateVerifier\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCertificateVerifier\",\"inputs\":[{\"name\":\"curveType\",\"type\":\"uint8\",\"internalType\":\"enumIKeyRegistrarTypes.CurveType\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentGlobalTableRoot\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getGenerator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getGeneratorReferenceTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getGlobalTableRootByTimestamp\",\"inputs\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getGlobalTableUpdateMessageHash\",\"inputs\":[{\"name\":\"globalTableRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getLatestReferenceBlockNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLatestReferenceTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getReferenceBlockNumberByTimestamp\",\"inputs\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getReferenceTimestampByBlockNumber\",\"inputs\":[{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"globalRootConfirmationThreshold\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_generator\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"_globalRootConfirmationThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"generatorInfo\",\"type\":\"tuple\",\"internalType\":\"structIOperatorTableCalculatorTypes.BN254OperatorSetInfo\",\"components\":[{\"name\":\"operatorInfoTreeRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"numOperators\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"aggregatePubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"totalWeights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]},{\"name\":\"generatorConfig\",\"type\":\"tuple\",\"internalType\":\"structICrossChainRegistryTypes.OperatorSetConfig\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxStalenessPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isRootValid\",\"inputs\":[{\"name\":\"globalTableRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRootValidByTimestamp\",\"inputs\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGenerator\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGlobalRootConfirmationThreshold\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateGenerator\",\"inputs\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"generatorInfo\",\"type\":\"tuple\",\"internalType\":\"structIOperatorTableCalculatorTypes.BN254OperatorSetInfo\",\"components\":[{\"name\":\"operatorInfoTreeRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"numOperators\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"aggregatePubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"totalWeights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]},{\"name\":\"generatorConfig\",\"type\":\"tuple\",\"internalType\":\"structICrossChainRegistryTypes.OperatorSetConfig\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxStalenessPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOperatorTable\",\"inputs\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"globalTableRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"operatorSetIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"operatorTableBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"GeneratorUpdated\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GlobalRootConfirmationThresholdUpdated\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GlobalRootDisabled\",\"inputs\":[{\"name\":\"globalTableRoot\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewGlobalTableRoot\",\"inputs\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"globalTableRoot\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CertificateInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"GlobalTableRootInFuture\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"GlobalTableRootStale\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidConfirmationThreshold\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCurveType\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidGlobalTableRoot\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidMessageHash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperatorSetProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProofLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRoot\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignatureLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"TableUpdateForPastTimestamp\",\"inputs\":[]}]",
	Bin: "0x60e060405234801561000f575f5ffd5b5060405161240238038061240283398101604081905261002e91610188565b6001600160a01b03808416608052821660a0528061004b8161005f565b60c052506100576100a5565b5050506102b9565b5f5f829050601f81511115610092578260405163305a27a960e01b8152600401610089919061025e565b60405180910390fd5b805161009d82610293565b179392505050565b5f54610100900460ff161561010c5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b6064820152608401610089565b5f5460ff9081161461015b575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b0381168114610171575f5ffd5b50565b634e487b7160e01b5f52604160045260245ffd5b5f5f5f6060848603121561019a575f5ffd5b83516101a58161015d565b60208501519093506101b68161015d565b60408501519092506001600160401b038111156101d1575f5ffd5b8401601f810186136101e1575f5ffd5b80516001600160401b038111156101fa576101fa610174565b604051601f8201601f19908116603f011681016001600160401b038111828210171561022857610228610174565b60405281815282820160200188101561023f575f5ffd5b8160208401602083015e5f602083830101528093505050509250925092565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b805160208083015191908110156102b3575f198160200360031b1b821691505b50919050565b60805160a05160c0516120ee6103145f395f61076601525f818161044a015281816107ec0152610aa901525f8181610471015281816107ac0152818161085a01528181610a0401528181610cf00152610fc401526120ee5ff3fe608060405234801561000f575f5ffd5b50600436106101c6575f3560e01c8063715018a6116100fe578063c252aa221161009e578063e944e0a81161006e578063e944e0a8146104ff578063eaaed9d514610512578063f2fde38b14610525578063fd967f4714610538575f5ffd5b8063c252aa2214610493578063c3621f0a146104b4578063c3be1e33146104c7578063c5916a39146104da575f5ffd5b80638da5cb5b116100d95780638da5cb5b146104215780639ea9477814610432578063ad0f958214610445578063b8c143061461046c575f5ffd5b8063715018a6146103fe5780637551ba341461040657806377d90e941461040e575f5ffd5b806331a599d2116101695780634af81d7a116101445780634af81d7a1461036257806354fd4d501461038957806364e1df841461039e5780636f728c50146103d3575f5ffd5b806331a599d2146103025780633ef6cd7a146103275780634624e6a31461034e575f5ffd5b80632370356c116101a45780632370356c1461026957806323b7b5b21461027c57806328522d79146102a457806330ef41b4146102d0575f5ffd5b8063021ab442146101ca578063193b79f3146101df5780631e2ca26014610221575b5f5ffd5b6101dd6101d8366004611350565b610541565b005b6102076101ed3660046113e0565b63ffffffff9081165f908152606960205260409020541690565b60405163ffffffff90911681526020015b60405180910390f35b6040805180820182525f80825260209182015281518083019092526066546001600160a01b0381168352600160a01b900463ffffffff16908201526040516102189190611420565b6101dd61027736600461142e565b61074b565b61020761028a3660046113e0565b63ffffffff9081165f908152606860205260409020541690565b60655462010000900463ffffffff165f908152606760205260409020545b604051908152602001610218565b6102f26102de366004611447565b5f908152606a602052604090205460ff1690565b6040519015158152602001610218565b60655462010000900463ffffffff9081165f9081526068602052604090205416610207565b6102c27f4491f5ee91595f938885ef73c9a1fa8a6d14ff9b9dab4aa24b8802bbb9bfc1cc81565b60655462010000900463ffffffff16610207565b6102c27f2eddfa6e51c2e0ba986436883fbc224e895ba21e8fc61421f6b10d11e25d008e81565b61039161075f565b604051610218919061145e565b6102f26103ac3660046113e0565b63ffffffff165f908152606760209081526040808320548352606a90915290205460ff1690565b6103e66103e13660046114a1565b61078f565b6040516001600160a01b039091168152602001610218565b6101dd61082e565b610207610841565b6101dd61041c3660046114ba565b6108cf565b6033546001600160a01b03166103e6565b6101dd610440366004611518565b6108e0565b6103e67f000000000000000000000000000000000000000000000000000000000000000081565b6103e67f000000000000000000000000000000000000000000000000000000000000000081565b6065546104a19061ffff1681565b60405161ffff9091168152602001610218565b6101dd6104c2366004611447565b610b0e565b6102c26104d53660046115b3565b610b83565b6102c26104e83660046113e0565b63ffffffff165f9081526067602052604090205490565b6101dd61050d3660046115f2565b610beb565b6101dd61052036600461164e565b610c03565b6101dd6105333660046116ba565b610e35565b6104a161271081565b5f54610100900460ff161580801561055f57505f54600160ff909116105b806105785750303b15801561057857505f5460ff166001145b6105e05760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b5f805460ff191660011790558015610601575f805461ff0019166101001790555b61060a87610ea7565b61061386610ef8565b61061c85610f42565b610627848484610fad565b63ffffffff8481165f8181526067602090815260408083207f2eddfa6e51c2e0ba986436883fbc224e895ba21e8fc61421f6b10d11e25d008e908190557f10b2e5f6b212b2c1d8f17fc240726b9af6c45b2a8fff776e3ef5e5aff1c35965805460ff191660011790556068835281842080544390971663ffffffff19978816811790915584526069909252808320805490951684179094556065805462010000850265ffffffff00001990911617905592517f010dcbe0d1e019c93357711f7bb6287d543b7ff7de74f29df3fb5ecceec8d3699190a38015610742575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050505050565b610753611029565b61075c81610f42565b50565b606061078a7f0000000000000000000000000000000000000000000000000000000000000000611083565b905090565b5f60028260028111156107a4576107a46116d5565b036107d057507f0000000000000000000000000000000000000000000000000000000000000000919050565b60018260028111156107e4576107e46116d5565b0361081057507f0000000000000000000000000000000000000000000000000000000000000000919050565b60405163fdea7c0960e01b815260040160405180910390fd5b919050565b610836611029565b61083f5f610ea7565b565b604051635ddb9b5b60e01b81525f906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690635ddb9b5b90610890906066906004016116e9565b602060405180830381865afa1580156108ab573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061078a9190611710565b6108d7611029565b61075c81610ef8565b5f5f5f5f6108ee86866110c0565b5f8e8152606a60205260409020549397509195509350915060ff166109265760405163504570e360e01b815260040160405180910390fd5b61092f8361078f565b6001600160a01b0316635ddb9b5b856040518263ffffffff1660e01b815260040161095a9190611420565b602060405180830381865afa158015610975573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109999190611710565b63ffffffff168b63ffffffff16116109c45760405163207617df60e01b815260040160405180910390fd5b6109e98b8b8b8b8b8b8b6040516109dc92919061172b565b6040518091039020611107565b60028360028111156109fd576109fd6116d5565b03610a8e577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316636738c40b858d610a3c856111a8565b866040518563ffffffff1660e01b8152600401610a5c9493929190611774565b5f604051808303815f87803b158015610a73575f5ffd5b505af1158015610a85573d5f5f3e3d5ffd5b50505050610b01565b6001836002811115610aa257610aa26116d5565b03610810577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166356d482f5858d610ae1856111ca565b866040518563ffffffff1660e01b8152600401610a5c94939291906117f0565b5050505050505050505050565b610b16611029565b5f818152606a602052604090205460ff16610b445760405163504570e360e01b815260040160405180910390fd5b5f818152606a6020526040808220805460ff191690555182917f8bd43de1250f58fe6ec9a78671a8b78dba70f0018656d157a3aeaabec389df3491a250565b604080517f4491f5ee91595f938885ef73c9a1fa8a6d14ff9b9dab4aa24b8802bbb9bfc1cc602082015290810184905263ffffffff8084166060830152821660808201525f9060a0016040516020818303038152906040528051906020012090509392505050565b610bf3611029565b610bfe838383610fad565b505050565b428263ffffffff161115610c2a57604051635a119db560e11b815260040160405180910390fd5b60655463ffffffff62010000909104811690831611610c5c5760405163037fa86b60e31b815260040160405180910390fd5b610c67838383610b83565b846020013514610c8a57604051638b56642d60e01b815260040160405180910390fd5b6040805160018082528183019092525f91602080830190803683375050606554825192935061ffff16918391505f90610cc557610cc56118a9565b61ffff90921660209283029190910190910152604051625f5e5d60e21b81525f906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063017d797490610d2a906066908a9087906004016119db565b6020604051808303815f875af1158015610d46573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610d6a9190611b68565b905080610d8a57604051633042041f60e21b815260040160405180910390fd5b6065805463ffffffff80871662010000810265ffffffff000019909316929092179092555f818152606860209081526040808320805495891663ffffffff1996871681179091558352606982528083208054909516841790945582825260678152838220899055888252606a9052828120805460ff19166001179055915187927f010dcbe0d1e019c93357711f7bb6287d543b7ff7de74f29df3fb5ecceec8d36991a3505050505050565b610e3d611029565b6001600160a01b038116610ea25760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016105d7565b61075c815b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b806066610f058282611b87565b9050507f3463431b09dfd43dec7349f8f24acfa753fe4cf40a26235402d213373df1585681604051610f379190611c18565b60405180910390a150565b61271061ffff82161115610f69576040516307336f0360e11b815260040160405180910390fd5b6065805461ffff191661ffff83169081179091556040519081527ff5d1836df8fcd7c1e54047e94ac8773d2855395603e2ef9ba5f5f16905f2259290602001610f37565b604051636738c40b60e01b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690636738c40b9061100090606690879087908790600401611c26565b5f604051808303815f87803b158015611017575f5ffd5b505af1158015610742573d5f5f3e3d5ffd5b6033546001600160a01b0316331461083f5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016105d7565b60605f61108f836111e0565b6040805160208082528183019092529192505f91906020820181803683375050509182525060208101929092525090565b604080518082019091525f8082526020820152604080518082019091525f808252602082018190529060606110f785870187611d79565b9299919850965090945092505050565b63ffffffff86165f90815260676020526040902054851461113b5760405163639d09b560e11b815260040160405180910390fd5b61118383838080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152508992508591505063ffffffff8816611207565b6111a05760405163afa42ca760e01b815260040160405180910390fd5b505050505050565b6111b06112b5565b818060200190518101906111c49190611eca565b92915050565b6060818060200190518101906111c49190611f79565b5f60ff8216601f8111156111c457604051632cd44ac360e21b815260040160405180910390fd5b5f8361121486858561121e565b1495945050505050565b5f6020845161122d919061207a565b1561124b576040516313717da960e21b815260040160405180910390fd5b8260205b855181116112ac5761126260028561207a565b5f0361128357815f528086015160205260405f20915060028404935061129a565b808601515f528160205260405f2091506002840493505b6112a5602082612099565b905061124f565b50949350505050565b60405180608001604052805f81526020015f81526020016112e760405180604001604052805f81526020015f81525090565b8152602001606081525090565b6001600160a01b038116811461075c575f5ffd5b5f60408284031215611318575f5ffd5b50919050565b803561ffff81168114610829575f5ffd5b63ffffffff8116811461075c575f5ffd5b5f60a08284031215611318575f5ffd5b5f5f5f5f5f5f6101008789031215611366575f5ffd5b8635611371816112f4565b95506113808860208901611308565b945061138e6060880161131e565b9350608087013561139e8161132f565b925060a08701356001600160401b038111156113b8575f5ffd5b6113c489828a01611340565b9250506113d48860c08901611308565b90509295509295509295565b5f602082840312156113f0575f5ffd5b81356113fb8161132f565b9392505050565b80516001600160a01b0316825260209081015163ffffffff16910152565b604081016111c48284611402565b5f6020828403121561143e575f5ffd5b6113fb8261131e565b5f60208284031215611457575f5ffd5b5035919050565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b803560038110610829575f5ffd5b5f602082840312156114b1575f5ffd5b6113fb82611493565b5f604082840312156114ca575f5ffd5b6113fb8383611308565b5f5f83601f8401126114e4575f5ffd5b5081356001600160401b038111156114fa575f5ffd5b602083019150836020828501011115611511575f5ffd5b9250929050565b5f5f5f5f5f5f5f60a0888a03121561152e575f5ffd5b87356115398161132f565b96506020880135955060408801356115508161132f565b945060608801356001600160401b0381111561156a575f5ffd5b6115768a828b016114d4565b90955093505060808801356001600160401b03811115611594575f5ffd5b6115a08a828b016114d4565b989b979a50959850939692959293505050565b5f5f5f606084860312156115c5575f5ffd5b8335925060208401356115d78161132f565b915060408401356115e78161132f565b809150509250925092565b5f5f5f60808486031215611604575f5ffd5b833561160f8161132f565b925060208401356001600160401b03811115611629575f5ffd5b61163586828701611340565b9250506116458560408601611308565b90509250925092565b5f5f5f5f60808587031215611661575f5ffd5b84356001600160401b03811115611676575f5ffd5b85016101208188031215611688575f5ffd5b935060208501359250604085013561169f8161132f565b915060608501356116af8161132f565b939692955090935050565b5f602082840312156116ca575f5ffd5b81356113fb816112f4565b634e487b7160e01b5f52602160045260245ffd5b604081016111c48284546001600160a01b038116825260a01c63ffffffff16602090910152565b5f60208284031215611720575f5ffd5b81516113fb8161132f565b818382375f9101908152919050565b5f8151808452602084019350602083015f5b8281101561176a57815186526020958601959091019060010161174c565b5093949350505050565b61177e8186611402565b63ffffffff8416604082015260c06060820152825160c0820152602083015160e08201525f60408401518051610100840152602081015161012084015250606084015160a06101408401526117d761016084018261173a565b9150506117e76080830184611402565b95945050505050565b5f60c082016117ff8388611402565b63ffffffff8616604084015260c0606084015280855180835260e08501915060e08160051b8601019250602087015f5b828110156118805786850360df19018452815180516001600160a01b0316865260209081015160409187018290529061186a9087018261173a565b955050602093840193919091019060010161182f565b50505050809150506117e76080830184611402565b634e487b7160e01b5f52604160045260245ffd5b634e487b7160e01b5f52603260045260245ffd5b5f5f8335601e198436030181126118d2575f5ffd5b83016020810192503590506001600160401b038111156118f0575f5ffd5b8060051b3603821315611511575f5ffd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b5f8235605e1983360301811261193d575f5ffd5b90910192915050565b8183525f6001600160fb1b0383111561195d575f5ffd5b8260051b80836020870137939093016020019392505050565b80358252602080820135908301525f61199260408301836118bd565b606060408601526117e7606086018284611946565b5f8151808452602084019350602083015f5b8281101561176a57815161ffff168652602095860195909101906001016119b9565b6119fe8185546001600160a01b038116825260a01c63ffffffff16602090910152565b608060408201525f6101a082018435611a168161132f565b63ffffffff166080840152602085013560a0840152604085013560c0840152606085013560e0840152604060808601610100850137604060c08601610140850137611a656101008601866118bd565b610120610180860152828184526101c0860190506101c08260051b8701019350825f5b83811015611b47578786036101bf19018352611aa48286611929565b8035611aaf8161132f565b63ffffffff168752602081013536829003601e19018112611ace575f5ffd5b81016020810190356001600160401b03811115611ae9575f5ffd5b803603821315611af7575f5ffd5b606060208a0152611b0c60608a018284611901565b915050611b1c6040830183611929565b91508781036040890152611b308183611976565b975050506020928301929190910190600101611a88565b50505050508281036060840152611b5e81856119a7565b9695505050505050565b5f60208284031215611b78575f5ffd5b815180151581146113fb575f5ffd5b8135611b92816112f4565b81546001600160a01b031981166001600160a01b039290921691821783556020840135611bbe8161132f565b6001600160c01b03199190911690911760a09190911b63ffffffff60a01b1617905550565b8035611bee816112f4565b6001600160a01b031682526020810135611c078161132f565b63ffffffff81166020840152505050565b604081016111c48284611be3565b611c498186546001600160a01b038116825260a01c63ffffffff16602090910152565b63ffffffff841660408281019190915260c06060808401829052853591840191909152602085013560e0840152908401356101008301528301356101208201525f611c9760808501856118bd565b60a0610140850152611cae61016085018284611946565b925050506117e76080830184611be3565b604080519081016001600160401b0381118282101715611ce157611ce1611895565b60405290565b604051608081016001600160401b0381118282101715611ce157611ce1611895565b604051601f8201601f191681016001600160401b0381118282101715611d3157611d31611895565b604052919050565b5f60408284031215611d49575f5ffd5b611d51611cbf565b90508135611d5e816112f4565b81526020820135611d6e8161132f565b602082015292915050565b5f5f5f5f60c08587031215611d8c575f5ffd5b611d968686611d39565b9350611da460408601611493565b9250611db38660608701611d39565b915060a08501356001600160401b03811115611dcd575f5ffd5b8501601f81018713611ddd575f5ffd5b80356001600160401b03811115611df657611df6611895565b611e09601f8201601f1916602001611d09565b818152886020838501011115611e1d575f5ffd5b816020840160208301375f6020838301015280935050505092959194509250565b5f6001600160401b03821115611e5657611e56611895565b5060051b60200190565b5f82601f830112611e6f575f5ffd5b8151611e82611e7d82611e3e565b611d09565b8082825260208201915060208360051b860101925085831115611ea3575f5ffd5b602085015b83811015611ec0578051835260209283019201611ea8565b5095945050505050565b5f60208284031215611eda575f5ffd5b81516001600160401b03811115611eef575f5ffd5b820180840360a0811215611f01575f5ffd5b611f09611ce7565b82518152602080840151908201526040603f1983011215611f28575f5ffd5b611f30611cbf565b604084810151825260608501516020830152820152608083015191506001600160401b03821115611f5f575f5ffd5b611f6b86838501611e60565b606082015295945050505050565b5f60208284031215611f89575f5ffd5b81516001600160401b03811115611f9e575f5ffd5b8201601f81018413611fae575f5ffd5b8051611fbc611e7d82611e3e565b8082825260208201915060208360051b850101925086831115611fdd575f5ffd5b602084015b8381101561206f5780516001600160401b03811115611fff575f5ffd5b85016040818a03601f19011215612014575f5ffd5b61201c611cbf565b602082015161202a816112f4565b815260408201516001600160401b03811115612044575f5ffd5b6120538b602083860101611e60565b6020830152508085525050602083019250602081019050611fe2565b509695505050505050565b5f8261209457634e487b7160e01b5f52601260045260245ffd5b500690565b808201808211156111c457634e487b7160e01b5f52601160045260245ffdfea2646970667358221220ea2c67407757dee4681cee5e8e6770618e0b9b7b41b36edddd03d470cabbb9e664736f6c634300081b0033",
=======
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_bn254CertificateVerifier\",\"type\":\"address\",\"internalType\":\"contractIBN254CertificateVerifier\"},{\"name\":\"_ecdsaCertificateVerifier\",\"type\":\"address\",\"internalType\":\"contractIECDSACertificateVerifier\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"_version\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"GLOBAL_TABLE_ROOT_CERT_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"INITIAL_GLOBAL_TABLE_ROOT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_BPS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bn254CertificateVerifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBN254CertificateVerifier\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"confirmGlobalTableRoot\",\"inputs\":[{\"name\":\"globalTableRootCert\",\"type\":\"tuple\",\"internalType\":\"structIBN254CertificateVerifierTypes.BN254Certificate\",\"components\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"nonSignerWitnesses\",\"type\":\"tuple[]\",\"internalType\":\"structIBN254CertificateVerifierTypes.BN254OperatorInfoWitness[]\",\"components\":[{\"name\":\"operatorIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorInfoProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"operatorInfo\",\"type\":\"tuple\",\"internalType\":\"structIOperatorTableCalculatorTypes.BN254OperatorInfo\",\"components\":[{\"name\":\"pubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}]}]},{\"name\":\"globalTableRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"disableRoot\",\"inputs\":[{\"name\":\"globalTableRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ecdsaCertificateVerifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIECDSACertificateVerifier\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCertificateVerifier\",\"inputs\":[{\"name\":\"curveType\",\"type\":\"uint8\",\"internalType\":\"enumIKeyRegistrarTypes.CurveType\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentGlobalTableRoot\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getGlobalConfirmerSetReferenceTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getGlobalRootConfirmerSet\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getGlobalTableRootByTimestamp\",\"inputs\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getGlobalTableUpdateMessageHash\",\"inputs\":[{\"name\":\"globalTableRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getLatestReferenceBlockNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLatestReferenceTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getReferenceBlockNumberByTimestamp\",\"inputs\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getReferenceTimestampByBlockNumber\",\"inputs\":[{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"globalRootConfirmationThreshold\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_globalRootConfirmerSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"_globalRootConfirmationThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"globalRootConfirmerSetInfo\",\"type\":\"tuple\",\"internalType\":\"structIOperatorTableCalculatorTypes.BN254OperatorSetInfo\",\"components\":[{\"name\":\"operatorInfoTreeRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"numOperators\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"aggregatePubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"totalWeights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]},{\"name\":\"globalRootConfirmerSetConfig\",\"type\":\"tuple\",\"internalType\":\"structICrossChainRegistryTypes.OperatorSetConfig\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxStalenessPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isRootValid\",\"inputs\":[{\"name\":\"globalTableRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRootValidByTimestamp\",\"inputs\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGlobalRootConfirmationThreshold\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGlobalRootConfirmerSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateGlobalRootConfirmerSet\",\"inputs\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"globalRootConfirmerSetInfo\",\"type\":\"tuple\",\"internalType\":\"structIOperatorTableCalculatorTypes.BN254OperatorSetInfo\",\"components\":[{\"name\":\"operatorInfoTreeRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"numOperators\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"aggregatePubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"totalWeights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]},{\"name\":\"globalRootConfirmerSetConfig\",\"type\":\"tuple\",\"internalType\":\"structICrossChainRegistryTypes.OperatorSetConfig\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxStalenessPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOperatorTable\",\"inputs\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"globalTableRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"operatorSetIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"operatorTableBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"GlobalRootConfirmationThresholdUpdated\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GlobalRootConfirmerSetUpdated\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GlobalRootDisabled\",\"inputs\":[{\"name\":\"globalTableRoot\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewGlobalTableRoot\",\"inputs\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"globalTableRoot\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CertificateInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CurrentlyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"GlobalTableRootInFuture\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"GlobalTableRootStale\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidConfirmationThreshold\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCurveType\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidGlobalTableRoot\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidMessageHash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNewPausedStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperatorSetProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProofLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRoot\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignatureLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyPauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnpauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"TableUpdateForPastTimestamp\",\"inputs\":[]}]",
	Bin: "0x610100604052348015610010575f5ffd5b506040516127c63803806127c683398101604081905261002f916101b9565b808484846001600160a01b03811661005a576040516339b190bb60e11b815260040160405180910390fd5b6001600160a01b0390811660805291821660a0521660c05261007b81610090565b60e052506100876100d6565b505050506102fe565b5f5f829050601f815111156100c3578260405163305a27a960e01b81526004016100ba91906102a3565b60405180910390fd5b80516100ce826102d8565b179392505050565b5f54610100900460ff161561013d5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b60648201526084016100ba565b5f5460ff9081161461018c575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146101a2575f5ffd5b50565b634e487b7160e01b5f52604160045260245ffd5b5f5f5f5f608085870312156101cc575f5ffd5b84516101d78161018e565b60208601519094506101e88161018e565b60408601519093506101f98161018e565b60608601519092506001600160401b03811115610214575f5ffd5b8501601f81018713610224575f5ffd5b80516001600160401b0381111561023d5761023d6101a5565b604051601f8201601f19908116603f011681016001600160401b038111828210171561026b5761026b6101a5565b604052818152828201602001891015610282575f5ffd5b8160208401602083015e5f6020838301015280935050505092959194509250565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b805160208083015191908110156102f8575f198160200360031b1b821691505b50919050565b60805160a05160c05160e0516124566103705f395f61071401525f81816104f9015281816107a901526109c501525f818161052001528181610630015281816107690152818161092001528181610bf401526111cb01525f81816104ae015281816110e9015261145401526124565ff3fe608060405234801561000f575f5ffd5b5060043610610208575f3560e01c80635c975abb1161011f578063c252aa22116100a9578063eaaed9d511610079578063eaaed9d5146105ae578063ed90dc60146105c1578063f2fde38b146105d4578063fabc1cbc146105e7578063fd967f47146105fa575f5ffd5b8063c252aa2214610542578063c3621f0a14610563578063c3be1e3314610576578063c5916a3914610589575f5ffd5b8063886f1195116100ef578063886f1195146104a95780638da5cb5b146104d05780639ea94778146104e1578063ad0f9582146104f4578063b8c143061461051b575f5ffd5b80635c975abb1461043957806364e1df84146104415780636f728c5014610476578063715018a6146104a1575f5ffd5b806330ef41b4116101a05780634628288911610170578063462828891461038a5780634af81d7a146103d257806354fd4d50146103f9578063595c6a671461040e5780635ac86ab714610416575f5ffd5b806330ef41b4146102f857806331a599d21461032a5780633ef6cd7a1461034f5780634624e6a314610376575f5ffd5b80631ab78d90116101db5780631ab78d901461027e5780632370356c1461029157806323b7b5b2146102a457806328522d79146102cc575f5ffd5b80630371406e1461020c5780630f3f8edd14610221578063136439dd14610243578063193b79f314610256575b5f5ffd5b61021f61021a36600461162d565b610603565b005b610229610617565b60405163ffffffff90911681526020015b60405180910390f35b61021f61025136600461164e565b6106aa565b610229610264366004611676565b63ffffffff9081165f908152609b60205260409020541690565b61021f61028c3660046116a1565b6106e4565b61021f61029f36600461170e565b6106fc565b6102296102b2366004611676565b63ffffffff9081165f908152609a60205260409020541690565b60975462010000900463ffffffff165f908152609960205260409020545b60405190815260200161023a565b61031a61030636600461164e565b5f908152609c602052604090205460ff1690565b604051901515815260200161023a565b60975462010000900463ffffffff9081165f908152609a602052604090205416610229565b6102ea7f4491f5ee91595f938885ef73c9a1fa8a6d14ff9b9dab4aa24b8802bbb9bfc1cc81565b60975462010000900463ffffffff16610229565b6040805180820182525f80825260209182015281518083019092526098546001600160a01b0381168352600160a01b900463ffffffff169082015260405161023a9190611745565b6102ea7f2eddfa6e51c2e0ba986436883fbc224e895ba21e8fc61421f6b10d11e25d008e81565b61040161070d565b60405161023a9190611753565b61021f610738565b61031a610424366004611788565b606654600160ff9092169190911b9081161490565b6066546102ea565b61031a61044f366004611676565b63ffffffff165f908152609960209081526040808320548352609c90915290205460ff1690565b6104896104843660046117b6565b61074c565b6040516001600160a01b03909116815260200161023a565b61021f6107eb565b6104897f000000000000000000000000000000000000000000000000000000000000000081565b6033546001600160a01b0316610489565b61021f6104ef366004611813565b6107fc565b6104897f000000000000000000000000000000000000000000000000000000000000000081565b6104897f000000000000000000000000000000000000000000000000000000000000000081565b6097546105509061ffff1681565b60405161ffff909116815260200161023a565b61021f61057136600461164e565b610a2a565b6102ea6105843660046118ae565b610a9f565b6102ea610597366004611676565b63ffffffff165f9081526099602052604090205490565b61021f6105bc3660046118ed565b610b07565b61021f6105cf36600461196d565b610d39565b61021f6105e2366004611a07565b610f4d565b61021f6105f536600461164e565b610fc3565b61055061271081565b61060b611030565b6106148161108a565b50565b604051635ddb9b5b60e01b81525f906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690635ddb9b5b9061066690609890600401611a22565b602060405180830381865afa158015610681573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106a59190611a49565b905090565b6106b26110d4565b60665481811681146106d75760405163c61dca5d60e01b815260040160405180910390fd5b6106e082611177565b5050565b6106ec611030565b6106f78383836111b4565b505050565b610704611030565b61061481611239565b60606106a57f00000000000000000000000000000000000000000000000000000000000000006112a4565b6107406110d4565b61074a5f19611177565b565b5f600282600281111561076157610761611a64565b0361078d57507f0000000000000000000000000000000000000000000000000000000000000000919050565b60018260028111156107a1576107a1611a64565b036107cd57507f0000000000000000000000000000000000000000000000000000000000000000919050565b60405163fdea7c0960e01b815260040160405180910390fd5b919050565b6107f3611030565b61074a5f6112e1565b5f5f5f5f61080a8686611332565b5f8e8152609c60205260409020549397509195509350915060ff166108425760405163504570e360e01b815260040160405180910390fd5b61084b8361074c565b6001600160a01b0316635ddb9b5b856040518263ffffffff1660e01b81526004016108769190611745565b602060405180830381865afa158015610891573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108b59190611a49565b63ffffffff168b63ffffffff16116108e05760405163207617df60e01b815260040160405180910390fd5b6109058b8b8b8b8b8b8b6040516108f8929190611a78565b6040518091039020611379565b600283600281111561091957610919611a64565b036109aa577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316636738c40b858d6109588561141a565b866040518563ffffffff1660e01b81526004016109789493929190611ac1565b5f604051808303815f87803b15801561098f575f5ffd5b505af11580156109a1573d5f5f3e3d5ffd5b50505050610a1d565b60018360028111156109be576109be611a64565b036107cd577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166356d482f5858d6109fd8561143c565b866040518563ffffffff1660e01b81526004016109789493929190611b3d565b5050505050505050505050565b610a326110d4565b5f818152609c602052604090205460ff16610a605760405163504570e360e01b815260040160405180910390fd5b5f818152609c6020526040808220805460ff191690555182917f8bd43de1250f58fe6ec9a78671a8b78dba70f0018656d157a3aeaabec389df3491a250565b604080517f4491f5ee91595f938885ef73c9a1fa8a6d14ff9b9dab4aa24b8802bbb9bfc1cc602082015290810184905263ffffffff8084166060830152821660808201525f9060a0016040516020818303038152906040528051906020012090509392505050565b428263ffffffff161115610b2e57604051635a119db560e11b815260040160405180910390fd5b60975463ffffffff62010000909104811690831611610b605760405163037fa86b60e31b815260040160405180910390fd5b610b6b838383610a9f565b846020013514610b8e57604051638b56642d60e01b815260040160405180910390fd5b6040805160018082528183019092525f91602080830190803683375050609754825192935061ffff16918391505f90610bc957610bc9611bf6565b61ffff90921660209283029190910190910152604051625f5e5d60e21b81525f906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063017d797490610c2e906098908a908790600401611d28565b6020604051808303815f875af1158015610c4a573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610c6e9190611eb5565b905080610c8e57604051633042041f60e21b815260040160405180910390fd5b6097805463ffffffff80871662010000810265ffffffff000019909316929092179092555f818152609a60209081526040808320805495891663ffffffff1996871681179091558352609b82528083208054909516841790945582825260998152838220899055888252609c9052828120805460ff19166001179055915187927f010dcbe0d1e019c93357711f7bb6287d543b7ff7de74f29df3fb5ecceec8d36991a3505050505050565b5f54610100900460ff1615808015610d5757505f54600160ff909116105b80610d705750303b158015610d7057505f5460ff166001145b610dd85760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b5f805460ff191660011790558015610df9575f805461ff0019166101001790555b610e02886112e1565b610e0b87611177565b610e148661108a565b610e1d85611239565b610e288484846111b4565b63ffffffff8481165f8181526099602090815260408083207f2eddfa6e51c2e0ba986436883fbc224e895ba21e8fc61421f6b10d11e25d008e908190557fd6e1a24cb7e68b47373042d0900dfd69bffcfc2c807e706164b25b408655b71f805460ff19166001179055609a835281842080544390971663ffffffff1997881681179091558452609b909252808320805490951684179094556097805462010000850265ffffffff00001990911617905592517f010dcbe0d1e019c93357711f7bb6287d543b7ff7de74f29df3fb5ecceec8d3699190a38015610f43575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050505050505050565b610f55611030565b6001600160a01b038116610fba5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610dcf565b610614816112e1565b610fcb611452565b60665480198219811614610ff25760405163c61dca5d60e01b815260040160405180910390fd5b606682905560405182815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c9060200160405180910390a25050565b6033546001600160a01b0316331461074a5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610dcf565b8060986110978282611ed4565b9050507f20100394950e66014c25009b45d12b675210a6e7a002044a0e3de6544e3c4b37816040516110c99190611f65565b60405180910390a150565b60405163237dfb4760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906346fbf68e90602401602060405180830381865afa158015611136573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061115a9190611eb5565b61074a57604051631d77d47760e21b815260040160405180910390fd5b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a250565b604051636738c40b60e01b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690636738c40b9061120790609890879087908790600401611f73565b5f604051808303815f87803b15801561121e575f5ffd5b505af1158015611230573d5f5f3e3d5ffd5b50505050505050565b61271061ffff82161115611260576040516307336f0360e11b815260040160405180910390fd5b6097805461ffff191661ffff83169081179091556040519081527ff5d1836df8fcd7c1e54047e94ac8773d2855395603e2ef9ba5f5f16905f22592906020016110c9565b60605f6112b083611503565b6040805160208082528183019092529192505f91906020820181803683375050509182525060208101929092525090565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b604080518082019091525f8082526020820152604080518082019091525f80825260208201819052906060611369858701876120c6565b9299919850965090945092505050565b63ffffffff86165f9081526099602052604090205485146113ad5760405163639d09b560e11b815260040160405180910390fd5b6113f583838080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152508992508591505063ffffffff881661152a565b6114125760405163afa42ca760e01b815260040160405180910390fd5b505050505050565b6114226115d8565b818060200190518101906114369190612217565b92915050565b60608180602001905181019061143691906122c6565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156114ae573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906114d291906123c7565b6001600160a01b0316336001600160a01b03161461074a5760405163794821ff60e01b815260040160405180910390fd5b5f60ff8216601f81111561143657604051632cd44ac360e21b815260040160405180910390fd5b5f83611537868585611541565b1495945050505050565b5f6020845161155091906123e2565b1561156e576040516313717da960e21b815260040160405180910390fd5b8260205b855181116115cf576115856002856123e2565b5f036115a657815f528086015160205260405f2091506002840493506115bd565b808601515f528160205260405f2091506002840493505b6115c8602082612401565b9050611572565b50949350505050565b60405180608001604052805f81526020015f815260200161160a60405180604001604052805f81526020015f81525090565b8152602001606081525090565b5f60408284031215611627575f5ffd5b50919050565b5f6040828403121561163d575f5ffd5b6116478383611617565b9392505050565b5f6020828403121561165e575f5ffd5b5035919050565b63ffffffff81168114610614575f5ffd5b5f60208284031215611686575f5ffd5b813561164781611665565b5f60a08284031215611627575f5ffd5b5f5f5f608084860312156116b3575f5ffd5b83356116be81611665565b925060208401356001600160401b038111156116d8575f5ffd5b6116e486828701611691565b9250506116f48560408601611617565b90509250925092565b803561ffff811681146107e6575f5ffd5b5f6020828403121561171e575f5ffd5b611647826116fd565b80516001600160a01b0316825260209081015163ffffffff16910152565b604081016114368284611727565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b5f60208284031215611798575f5ffd5b813560ff81168114611647575f5ffd5b8035600381106107e6575f5ffd5b5f602082840312156117c6575f5ffd5b611647826117a8565b5f5f83601f8401126117df575f5ffd5b5081356001600160401b038111156117f5575f5ffd5b60208301915083602082850101111561180c575f5ffd5b9250929050565b5f5f5f5f5f5f5f60a0888a031215611829575f5ffd5b873561183481611665565b965060208801359550604088013561184b81611665565b945060608801356001600160401b03811115611865575f5ffd5b6118718a828b016117cf565b90955093505060808801356001600160401b0381111561188f575f5ffd5b61189b8a828b016117cf565b989b979a50959850939692959293505050565b5f5f5f606084860312156118c0575f5ffd5b8335925060208401356118d281611665565b915060408401356118e281611665565b809150509250925092565b5f5f5f5f60808587031215611900575f5ffd5b84356001600160401b03811115611915575f5ffd5b85016101208188031215611927575f5ffd5b935060208501359250604085013561193e81611665565b9150606085013561194e81611665565b939692955090935050565b6001600160a01b0381168114610614575f5ffd5b5f5f5f5f5f5f5f610120888a031215611984575f5ffd5b873561198f81611959565b9650602088013595506119a58960408a01611617565b94506119b3608089016116fd565b935060a08801356119c381611665565b925060c08801356001600160401b038111156119dd575f5ffd5b6119e98a828b01611691565b9250506119f98960e08a01611617565b905092959891949750929550565b5f60208284031215611a17575f5ffd5b813561164781611959565b604081016114368284546001600160a01b038116825260a01c63ffffffff16602090910152565b5f60208284031215611a59575f5ffd5b815161164781611665565b634e487b7160e01b5f52602160045260245ffd5b818382375f9101908152919050565b5f8151808452602084019350602083015f5b82811015611ab7578151865260209586019590910190600101611a99565b5093949350505050565b611acb8186611727565b63ffffffff8416604082015260c06060820152825160c0820152602083015160e08201525f60408401518051610100840152602081015161012084015250606084015160a0610140840152611b24610160840182611a87565b915050611b346080830184611727565b95945050505050565b5f60c08201611b4c8388611727565b63ffffffff8616604084015260c0606084015280855180835260e08501915060e08160051b8601019250602087015f5b82811015611bcd5786850360df19018452815180516001600160a01b03168652602090810151604091870182905290611bb790870182611a87565b9550506020938401939190910190600101611b7c565b5050505080915050611b346080830184611727565b634e487b7160e01b5f52604160045260245ffd5b634e487b7160e01b5f52603260045260245ffd5b5f5f8335601e19843603018112611c1f575f5ffd5b83016020810192503590506001600160401b03811115611c3d575f5ffd5b8060051b360382131561180c575f5ffd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b5f8235605e19833603018112611c8a575f5ffd5b90910192915050565b8183525f6001600160fb1b03831115611caa575f5ffd5b8260051b80836020870137939093016020019392505050565b80358252602080820135908301525f611cdf6040830183611c0a565b60606040860152611b34606086018284611c93565b5f8151808452602084019350602083015f5b82811015611ab757815161ffff16865260209586019590910190600101611d06565b611d4b8185546001600160a01b038116825260a01c63ffffffff16602090910152565b608060408201525f6101a082018435611d6381611665565b63ffffffff166080840152602085013560a0840152604085013560c0840152606085013560e0840152604060808601610100850137604060c08601610140850137611db2610100860186611c0a565b610120610180860152828184526101c0860190506101c08260051b8701019350825f5b83811015611e94578786036101bf19018352611df18286611c76565b8035611dfc81611665565b63ffffffff168752602081013536829003601e19018112611e1b575f5ffd5b81016020810190356001600160401b03811115611e36575f5ffd5b803603821315611e44575f5ffd5b606060208a0152611e5960608a018284611c4e565b915050611e696040830183611c76565b91508781036040890152611e7d8183611cc3565b975050506020928301929190910190600101611dd5565b50505050508281036060840152611eab8185611cf4565b9695505050505050565b5f60208284031215611ec5575f5ffd5b81518015158114611647575f5ffd5b8135611edf81611959565b81546001600160a01b031981166001600160a01b039290921691821783556020840135611f0b81611665565b6001600160c01b03199190911690911760a09190911b63ffffffff60a01b1617905550565b8035611f3b81611959565b6001600160a01b031682526020810135611f5481611665565b63ffffffff81166020840152505050565b604081016114368284611f30565b611f968186546001600160a01b038116825260a01c63ffffffff16602090910152565b63ffffffff841660408281019190915260c06060808401829052853591840191909152602085013560e0840152908401356101008301528301356101208201525f611fe46080850185611c0a565b60a0610140850152611ffb61016085018284611c93565b92505050611b346080830184611f30565b604080519081016001600160401b038111828210171561202e5761202e611be2565b60405290565b604051608081016001600160401b038111828210171561202e5761202e611be2565b604051601f8201601f191681016001600160401b038111828210171561207e5761207e611be2565b604052919050565b5f60408284031215612096575f5ffd5b61209e61200c565b905081356120ab81611959565b815260208201356120bb81611665565b602082015292915050565b5f5f5f5f60c085870312156120d9575f5ffd5b6120e38686612086565b93506120f1604086016117a8565b92506121008660608701612086565b915060a08501356001600160401b0381111561211a575f5ffd5b8501601f8101871361212a575f5ffd5b80356001600160401b0381111561214357612143611be2565b612156601f8201601f1916602001612056565b81815288602083850101111561216a575f5ffd5b816020840160208301375f6020838301015280935050505092959194509250565b5f6001600160401b038211156121a3576121a3611be2565b5060051b60200190565b5f82601f8301126121bc575f5ffd5b81516121cf6121ca8261218b565b612056565b8082825260208201915060208360051b8601019250858311156121f0575f5ffd5b602085015b8381101561220d5780518352602092830192016121f5565b5095945050505050565b5f60208284031215612227575f5ffd5b81516001600160401b0381111561223c575f5ffd5b820180840360a081121561224e575f5ffd5b612256612034565b82518152602080840151908201526040603f1983011215612275575f5ffd5b61227d61200c565b604084810151825260608501516020830152820152608083015191506001600160401b038211156122ac575f5ffd5b6122b8868385016121ad565b606082015295945050505050565b5f602082840312156122d6575f5ffd5b81516001600160401b038111156122eb575f5ffd5b8201601f810184136122fb575f5ffd5b80516123096121ca8261218b565b8082825260208201915060208360051b85010192508683111561232a575f5ffd5b602084015b838110156123bc5780516001600160401b0381111561234c575f5ffd5b85016040818a03601f19011215612361575f5ffd5b61236961200c565b602082015161237781611959565b815260408201516001600160401b03811115612391575f5ffd5b6123a08b6020838601016121ad565b602083015250808552505060208301925060208101905061232f565b509695505050505050565b5f602082840312156123d7575f5ffd5b815161164781611959565b5f826123fc57634e487b7160e01b5f52601260045260245ffd5b500690565b8082018082111561143657634e487b7160e01b5f52601160045260245ffdfea2646970667358221220979609cbaca338c929eec7517264839afdb3b8231b809a7d5edd61bd9d970e9a64736f6c634300081b0033",
>>>>>>> 46842ded (feat: add pauser)
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

// INITIALGLOBALTABLEROOT is a free data retrieval call binding the contract method 0x4af81d7a.
//
// Solidity: function INITIAL_GLOBAL_TABLE_ROOT() view returns(bytes32)
func (_OperatorTableUpdater *OperatorTableUpdaterCaller) INITIALGLOBALTABLEROOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _OperatorTableUpdater.contract.Call(opts, &out, "INITIAL_GLOBAL_TABLE_ROOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// INITIALGLOBALTABLEROOT is a free data retrieval call binding the contract method 0x4af81d7a.
//
// Solidity: function INITIAL_GLOBAL_TABLE_ROOT() view returns(bytes32)
func (_OperatorTableUpdater *OperatorTableUpdaterSession) INITIALGLOBALTABLEROOT() ([32]byte, error) {
	return _OperatorTableUpdater.Contract.INITIALGLOBALTABLEROOT(&_OperatorTableUpdater.CallOpts)
}

// INITIALGLOBALTABLEROOT is a free data retrieval call binding the contract method 0x4af81d7a.
//
// Solidity: function INITIAL_GLOBAL_TABLE_ROOT() view returns(bytes32)
func (_OperatorTableUpdater *OperatorTableUpdaterCallerSession) INITIALGLOBALTABLEROOT() ([32]byte, error) {
	return _OperatorTableUpdater.Contract.INITIALGLOBALTABLEROOT(&_OperatorTableUpdater.CallOpts)
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

// Initialize is a paid mutator transaction binding the contract method 0xed90dc60.
//
<<<<<<< HEAD
// Solidity: function initialize(address owner, (address,uint32) _generator, uint16 _globalRootConfirmationThreshold, uint32 referenceTimestamp, (bytes32,uint256,(uint256,uint256),uint256[]) generatorInfo, (address,uint32) generatorConfig) returns()
func (_OperatorTableUpdater *OperatorTableUpdaterTransactor) Initialize(opts *bind.TransactOpts, owner common.Address, _generator OperatorSet, _globalRootConfirmationThreshold uint16, referenceTimestamp uint32, generatorInfo IOperatorTableCalculatorTypesBN254OperatorSetInfo, generatorConfig ICrossChainRegistryTypesOperatorSetConfig) (*types.Transaction, error) {
	return _OperatorTableUpdater.contract.Transact(opts, "initialize", owner, _generator, _globalRootConfirmationThreshold, referenceTimestamp, generatorInfo, generatorConfig)
=======
// Solidity: function initialize(address owner, uint256 initialPausedStatus, (address,uint32) _globalRootConfirmerSet, uint16 _globalRootConfirmationThreshold, uint32 referenceTimestamp, (bytes32,uint256,(uint256,uint256),uint256[]) globalRootConfirmerSetInfo, (address,uint32) globalRootConfirmerSetConfig) returns()
func (_OperatorTableUpdater *OperatorTableUpdaterTransactor) Initialize(opts *bind.TransactOpts, owner common.Address, initialPausedStatus *big.Int, _globalRootConfirmerSet OperatorSet, _globalRootConfirmationThreshold uint16, referenceTimestamp uint32, globalRootConfirmerSetInfo IOperatorTableCalculatorTypesBN254OperatorSetInfo, globalRootConfirmerSetConfig ICrossChainRegistryTypesOperatorSetConfig) (*types.Transaction, error) {
	return _OperatorTableUpdater.contract.Transact(opts, "initialize", owner, initialPausedStatus, _globalRootConfirmerSet, _globalRootConfirmationThreshold, referenceTimestamp, globalRootConfirmerSetInfo, globalRootConfirmerSetConfig)
>>>>>>> 46842ded (feat: add pauser)
}

// Initialize is a paid mutator transaction binding the contract method 0xed90dc60.
//
<<<<<<< HEAD
// Solidity: function initialize(address owner, (address,uint32) _generator, uint16 _globalRootConfirmationThreshold, uint32 referenceTimestamp, (bytes32,uint256,(uint256,uint256),uint256[]) generatorInfo, (address,uint32) generatorConfig) returns()
func (_OperatorTableUpdater *OperatorTableUpdaterSession) Initialize(owner common.Address, _generator OperatorSet, _globalRootConfirmationThreshold uint16, referenceTimestamp uint32, generatorInfo IOperatorTableCalculatorTypesBN254OperatorSetInfo, generatorConfig ICrossChainRegistryTypesOperatorSetConfig) (*types.Transaction, error) {
	return _OperatorTableUpdater.Contract.Initialize(&_OperatorTableUpdater.TransactOpts, owner, _generator, _globalRootConfirmationThreshold, referenceTimestamp, generatorInfo, generatorConfig)
=======
// Solidity: function initialize(address owner, uint256 initialPausedStatus, (address,uint32) _globalRootConfirmerSet, uint16 _globalRootConfirmationThreshold, uint32 referenceTimestamp, (bytes32,uint256,(uint256,uint256),uint256[]) globalRootConfirmerSetInfo, (address,uint32) globalRootConfirmerSetConfig) returns()
func (_OperatorTableUpdater *OperatorTableUpdaterSession) Initialize(owner common.Address, initialPausedStatus *big.Int, _globalRootConfirmerSet OperatorSet, _globalRootConfirmationThreshold uint16, referenceTimestamp uint32, globalRootConfirmerSetInfo IOperatorTableCalculatorTypesBN254OperatorSetInfo, globalRootConfirmerSetConfig ICrossChainRegistryTypesOperatorSetConfig) (*types.Transaction, error) {
	return _OperatorTableUpdater.Contract.Initialize(&_OperatorTableUpdater.TransactOpts, owner, initialPausedStatus, _globalRootConfirmerSet, _globalRootConfirmationThreshold, referenceTimestamp, globalRootConfirmerSetInfo, globalRootConfirmerSetConfig)
>>>>>>> 46842ded (feat: add pauser)
}

// Initialize is a paid mutator transaction binding the contract method 0xed90dc60.
//
<<<<<<< HEAD
// Solidity: function initialize(address owner, (address,uint32) _generator, uint16 _globalRootConfirmationThreshold, uint32 referenceTimestamp, (bytes32,uint256,(uint256,uint256),uint256[]) generatorInfo, (address,uint32) generatorConfig) returns()
func (_OperatorTableUpdater *OperatorTableUpdaterTransactorSession) Initialize(owner common.Address, _generator OperatorSet, _globalRootConfirmationThreshold uint16, referenceTimestamp uint32, generatorInfo IOperatorTableCalculatorTypesBN254OperatorSetInfo, generatorConfig ICrossChainRegistryTypesOperatorSetConfig) (*types.Transaction, error) {
	return _OperatorTableUpdater.Contract.Initialize(&_OperatorTableUpdater.TransactOpts, owner, _generator, _globalRootConfirmationThreshold, referenceTimestamp, generatorInfo, generatorConfig)
=======
// Solidity: function initialize(address owner, uint256 initialPausedStatus, (address,uint32) _globalRootConfirmerSet, uint16 _globalRootConfirmationThreshold, uint32 referenceTimestamp, (bytes32,uint256,(uint256,uint256),uint256[]) globalRootConfirmerSetInfo, (address,uint32) globalRootConfirmerSetConfig) returns()
func (_OperatorTableUpdater *OperatorTableUpdaterTransactorSession) Initialize(owner common.Address, initialPausedStatus *big.Int, _globalRootConfirmerSet OperatorSet, _globalRootConfirmationThreshold uint16, referenceTimestamp uint32, globalRootConfirmerSetInfo IOperatorTableCalculatorTypesBN254OperatorSetInfo, globalRootConfirmerSetConfig ICrossChainRegistryTypesOperatorSetConfig) (*types.Transaction, error) {
	return _OperatorTableUpdater.Contract.Initialize(&_OperatorTableUpdater.TransactOpts, owner, initialPausedStatus, _globalRootConfirmerSet, _globalRootConfirmationThreshold, referenceTimestamp, globalRootConfirmerSetInfo, globalRootConfirmerSetConfig)
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
>>>>>>> 46842ded (feat: add pauser)
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

// SetGenerator is a paid mutator transaction binding the contract method 0x77d90e94.
//
// Solidity: function setGenerator((address,uint32) operatorSet) returns()
func (_OperatorTableUpdater *OperatorTableUpdaterTransactor) SetGenerator(opts *bind.TransactOpts, operatorSet OperatorSet) (*types.Transaction, error) {
	return _OperatorTableUpdater.contract.Transact(opts, "setGenerator", operatorSet)
}

// SetGenerator is a paid mutator transaction binding the contract method 0x77d90e94.
//
// Solidity: function setGenerator((address,uint32) operatorSet) returns()
func (_OperatorTableUpdater *OperatorTableUpdaterSession) SetGenerator(operatorSet OperatorSet) (*types.Transaction, error) {
	return _OperatorTableUpdater.Contract.SetGenerator(&_OperatorTableUpdater.TransactOpts, operatorSet)
}

// SetGenerator is a paid mutator transaction binding the contract method 0x77d90e94.
//
// Solidity: function setGenerator((address,uint32) operatorSet) returns()
func (_OperatorTableUpdater *OperatorTableUpdaterTransactorSession) SetGenerator(operatorSet OperatorSet) (*types.Transaction, error) {
	return _OperatorTableUpdater.Contract.SetGenerator(&_OperatorTableUpdater.TransactOpts, operatorSet)
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

<<<<<<< HEAD
// UpdateGenerator is a paid mutator transaction binding the contract method 0xe944e0a8.
=======
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

// UpdateGlobalRootConfirmerSet is a paid mutator transaction binding the contract method 0x1ab78d90.
>>>>>>> 46842ded (feat: add pauser)
//
// Solidity: function updateGenerator(uint32 referenceTimestamp, (bytes32,uint256,(uint256,uint256),uint256[]) generatorInfo, (address,uint32) generatorConfig) returns()
func (_OperatorTableUpdater *OperatorTableUpdaterTransactor) UpdateGenerator(opts *bind.TransactOpts, referenceTimestamp uint32, generatorInfo IOperatorTableCalculatorTypesBN254OperatorSetInfo, generatorConfig ICrossChainRegistryTypesOperatorSetConfig) (*types.Transaction, error) {
	return _OperatorTableUpdater.contract.Transact(opts, "updateGenerator", referenceTimestamp, generatorInfo, generatorConfig)
}

// UpdateGenerator is a paid mutator transaction binding the contract method 0xe944e0a8.
//
// Solidity: function updateGenerator(uint32 referenceTimestamp, (bytes32,uint256,(uint256,uint256),uint256[]) generatorInfo, (address,uint32) generatorConfig) returns()
func (_OperatorTableUpdater *OperatorTableUpdaterSession) UpdateGenerator(referenceTimestamp uint32, generatorInfo IOperatorTableCalculatorTypesBN254OperatorSetInfo, generatorConfig ICrossChainRegistryTypesOperatorSetConfig) (*types.Transaction, error) {
	return _OperatorTableUpdater.Contract.UpdateGenerator(&_OperatorTableUpdater.TransactOpts, referenceTimestamp, generatorInfo, generatorConfig)
}

// UpdateGenerator is a paid mutator transaction binding the contract method 0xe944e0a8.
//
// Solidity: function updateGenerator(uint32 referenceTimestamp, (bytes32,uint256,(uint256,uint256),uint256[]) generatorInfo, (address,uint32) generatorConfig) returns()
func (_OperatorTableUpdater *OperatorTableUpdaterTransactorSession) UpdateGenerator(referenceTimestamp uint32, generatorInfo IOperatorTableCalculatorTypesBN254OperatorSetInfo, generatorConfig ICrossChainRegistryTypesOperatorSetConfig) (*types.Transaction, error) {
	return _OperatorTableUpdater.Contract.UpdateGenerator(&_OperatorTableUpdater.TransactOpts, referenceTimestamp, generatorInfo, generatorConfig)
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

// OperatorTableUpdaterGeneratorUpdatedIterator is returned from FilterGeneratorUpdated and is used to iterate over the raw logs and unpacked data for GeneratorUpdated events raised by the OperatorTableUpdater contract.
type OperatorTableUpdaterGeneratorUpdatedIterator struct {
	Event *OperatorTableUpdaterGeneratorUpdated // Event containing the contract specifics and raw log

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
func (it *OperatorTableUpdaterGeneratorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorTableUpdaterGeneratorUpdated)
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
		it.Event = new(OperatorTableUpdaterGeneratorUpdated)
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
func (it *OperatorTableUpdaterGeneratorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorTableUpdaterGeneratorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorTableUpdaterGeneratorUpdated represents a GeneratorUpdated event raised by the OperatorTableUpdater contract.
type OperatorTableUpdaterGeneratorUpdated struct {
	OperatorSet OperatorSet
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterGeneratorUpdated is a free log retrieval operation binding the contract event 0x3463431b09dfd43dec7349f8f24acfa753fe4cf40a26235402d213373df15856.
//
// Solidity: event GeneratorUpdated((address,uint32) operatorSet)
func (_OperatorTableUpdater *OperatorTableUpdaterFilterer) FilterGeneratorUpdated(opts *bind.FilterOpts) (*OperatorTableUpdaterGeneratorUpdatedIterator, error) {

	logs, sub, err := _OperatorTableUpdater.contract.FilterLogs(opts, "GeneratorUpdated")
	if err != nil {
		return nil, err
	}
	return &OperatorTableUpdaterGeneratorUpdatedIterator{contract: _OperatorTableUpdater.contract, event: "GeneratorUpdated", logs: logs, sub: sub}, nil
}

// WatchGeneratorUpdated is a free log subscription operation binding the contract event 0x3463431b09dfd43dec7349f8f24acfa753fe4cf40a26235402d213373df15856.
//
// Solidity: event GeneratorUpdated((address,uint32) operatorSet)
func (_OperatorTableUpdater *OperatorTableUpdaterFilterer) WatchGeneratorUpdated(opts *bind.WatchOpts, sink chan<- *OperatorTableUpdaterGeneratorUpdated) (event.Subscription, error) {

	logs, sub, err := _OperatorTableUpdater.contract.WatchLogs(opts, "GeneratorUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorTableUpdaterGeneratorUpdated)
				if err := _OperatorTableUpdater.contract.UnpackLog(event, "GeneratorUpdated", log); err != nil {
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
func (_OperatorTableUpdater *OperatorTableUpdaterFilterer) ParseGeneratorUpdated(log types.Log) (*OperatorTableUpdaterGeneratorUpdated, error) {
	event := new(OperatorTableUpdaterGeneratorUpdated)
	if err := _OperatorTableUpdater.contract.UnpackLog(event, "GeneratorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
