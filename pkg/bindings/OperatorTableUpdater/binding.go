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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_bn254CertificateVerifier\",\"type\":\"address\",\"internalType\":\"contractIBN254CertificateVerifier\"},{\"name\":\"_ecdsaCertificateVerifier\",\"type\":\"address\",\"internalType\":\"contractIECDSACertificateVerifier\"},{\"name\":\"_version\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"MAX_BPS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bn254CertificateVerifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBN254CertificateVerifier\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"confirmGlobalTableRoot\",\"inputs\":[{\"name\":\"globalTableRootCert\",\"type\":\"tuple\",\"internalType\":\"structIBN254CertificateVerifierTypes.BN254Certificate\",\"components\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"nonSignerWitnesses\",\"type\":\"tuple[]\",\"internalType\":\"structIBN254CertificateVerifierTypes.BN254OperatorInfoWitness[]\",\"components\":[{\"name\":\"operatorIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorInfoProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"operatorInfo\",\"type\":\"tuple\",\"internalType\":\"structIBN254TableCalculatorTypes.BN254OperatorInfo\",\"components\":[{\"name\":\"pubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}]}]},{\"name\":\"globalTableRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ecdsaCertificateVerifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIECDSACertificateVerifier\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCertificateVerifier\",\"inputs\":[{\"name\":\"curveType\",\"type\":\"uint8\",\"internalType\":\"enumIKeyRegistrarTypes.CurveType\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentGlobalTableRoot\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getGlobalRootConfirmerSet\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getGlobalTableRootByTimestamp\",\"inputs\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLatestReferenceTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"globalRootConfirmationThreshold\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_globalRootConfirmerSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"_globalRootConfirmationThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"globalRootConfirmerSetInfo\",\"type\":\"tuple\",\"internalType\":\"structIBN254TableCalculatorTypes.BN254OperatorSetInfo\",\"components\":[{\"name\":\"operatorInfoTreeRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"numOperators\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"aggregatePubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"totalWeights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]},{\"name\":\"globalRootConfirmerSetConfig\",\"type\":\"tuple\",\"internalType\":\"structICrossChainRegistryTypes.OperatorSetConfig\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxStalenessPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGlobalRootConfirmationThreshold\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGlobalRootConfirmerSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOperatorTable\",\"inputs\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"globalTableRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"operatorSetIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"operatorTableBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"GlobalRootConfirmationThresholdUpdated\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GlobalRootConfirmerSetUpdated\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewGlobalTableRoot\",\"inputs\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"globalTableRoot\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CertificateInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"GlobalTableRootInFuture\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"GlobalTableRootStale\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidConfirmationThreshold\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCurveType\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidGlobalTableRoot\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperatorSetProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProofLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"TableRootNotInCertificate\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TableUpdateForPastTimestamp\",\"inputs\":[]}]",
	Bin: "0x60e060405234801561000f575f5ffd5b50604051611dc5380380611dc583398101604081905261002e91610188565b6001600160a01b03808416608052821660a0528061004b8161005f565b60c052506100576100a5565b5050506102b9565b5f5f829050601f81511115610092578260405163305a27a960e01b8152600401610089919061025e565b60405180910390fd5b805161009d82610293565b179392505050565b5f54610100900460ff161561010c5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b6064820152608401610089565b5f5460ff9081161461015b575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b0381168114610171575f5ffd5b50565b634e487b7160e01b5f52604160045260245ffd5b5f5f5f6060848603121561019a575f5ffd5b83516101a58161015d565b60208501519093506101b68161015d565b60408501519092506001600160401b038111156101d1575f5ffd5b8401601f810186136101e1575f5ffd5b80516001600160401b038111156101fa576101fa610174565b604051601f8201601f19908116603f011681016001600160401b038111828210171561022857610228610174565b60405281815282820160200188101561023f575f5ffd5b8160208401602083015e5f602083830101528093505050509250925092565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b805160208083015191908110156102b3575f198160200360031b1b821691505b50919050565b60805160a05160c051611ab861030d5f395f61050c01525f818161027101528181610770015261095d01525f81816102980152818161040e015281816106180152818161073001526108b70152611ab85ff3fe608060405234801561000f575f5ffd5b5060043610610111575f3560e01c8063715018a61161009e578063b8c143061161006e578063b8c1430614610293578063c252aa22146102ba578063c5916a39146102db578063f2fde38b14610300578063fd967f4714610313575f5ffd5b8063715018a6146102405780638da5cb5b146102485780639ea9477814610259578063ad0f95821461026c575f5ffd5b80634624e6a3116100e45780634624e6a31461018157806346282889146101a557806354fd4d50146101ed5780636ab40904146102025780636f728c5014610215575f5ffd5b8063021ab442146101155780630371406e1461012a5780632370356c1461013d57806328522d7914610150575b5f5ffd5b610128610123366004610e47565b61031c565b005b610128610138366004610edb565b6104e0565b61012861014b366004610efc565b6104f4565b60655462010000900463ffffffff165f908152606760205260409020545b6040519081526020015b60405180910390f35b60655462010000900463ffffffff1660405163ffffffff9091168152602001610178565b6040805180820182525f80825260209182015281518083019092526066546001600160a01b0381168352600160a01b900463ffffffff16908201526040516101789190610f33565b6101f5610505565b6040516101789190610f41565b610128610210366004610f76565b610535565b610228610223366004610fdf565b610713565b6040516001600160a01b039091168152602001610178565b6101286107b2565b6033546001600160a01b0316610228565b61012861026736600461103c565b6107c5565b6102287f000000000000000000000000000000000000000000000000000000000000000081565b6102287f000000000000000000000000000000000000000000000000000000000000000081565b6065546102c89061ffff1681565b60405161ffff9091168152602001610178565b61016e6102e93660046110d7565b63ffffffff165f9081526067602052604090205490565b61012861030e3660046110f2565b6109c2565b6102c861271081565b5f54610100900460ff161580801561033a57505f54600160ff909116105b806103535750303b15801561035357505f5460ff166001145b6103bb5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b5f805460ff1916600117905580156103dc575f805461ff0019166101001790555b6103e587610a34565b6103ee86610a85565b6103f785610acf565b604051636738c40b60e01b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690636738c40b906104499089908890889088906004016111b6565b5f604051808303815f87803b158015610460575f5ffd5b505af1158015610472573d5f5f3e3d5ffd5b50506065805465ffffffff000019166201000063ffffffff891602179055505080156104d7575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050505050565b6104e8610b3a565b6104f181610a85565b50565b6104fc610b3a565b6104f181610acf565b60606105307f0000000000000000000000000000000000000000000000000000000000000000610b94565b905090565b428163ffffffff16111561055c57604051635a119db560e11b815260040160405180910390fd5b60655463ffffffff6201000090910481169082161161058e5760405163037fa86b60e31b815260040160405180910390fd5b826020013582146105b25760405163327f9d0760e21b815260040160405180910390fd5b6040805160018082528183019092525f91602080830190803683375050606554825192935061ffff16918391505f906105ed576105ed611253565b61ffff90921660209283029190910190910152604051625f5e5d60e21b81525f906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063017d797490610652906066908990879060040161131b565b6020604051808303815f875af115801561066e573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106929190611497565b9050806106b257604051633042041f60e21b815260040160405180910390fd5b6065805465ffffffff000019166201000063ffffffff8616908102919091179091555f81815260676020526040808220879055518692917f010dcbe0d1e019c93357711f7bb6287d543b7ff7de74f29df3fb5ecceec8d36991a35050505050565b5f6002826002811115610728576107286114b6565b0361075457507f0000000000000000000000000000000000000000000000000000000000000000919050565b6001826002811115610768576107686114b6565b0361079457507f0000000000000000000000000000000000000000000000000000000000000000919050565b60405163fdea7c0960e01b815260040160405180910390fd5b919050565b6107ba610b3a565b6107c35f610a34565b565b5f5f5f6107d28585610bd1565b9250925092506107e182610713565b6001600160a01b0316635ddb9b5b846040518263ffffffff1660e01b815260040161080c9190610f33565b6020604051808303815f875af1158015610828573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061084c91906114ca565b63ffffffff168a63ffffffff16116108775760405163207617df60e01b815260040160405180910390fd5b61089c8a8a8a8a8a8a8a60405161088f9291906114e5565b6040518091039020610c14565b60028260028111156108b0576108b06114b6565b03610942577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316636738c40b848c6108f08989610cb5565b856040518563ffffffff1660e01b81526004016109109493929190611524565b5f604051808303815f87803b158015610927575f5ffd5b505af1158015610939573d5f5f3e3d5ffd5b505050506109b6565b6001826002811115610956576109566114b6565b03610794577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166356d482f5848c6109968989610cd3565b856040518563ffffffff1660e01b81526004016109109493929190611597565b50505050505050505050565b6109ca610b3a565b6001600160a01b038116610a2f5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016103b2565b6104f1815b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b806066610a92828261163c565b9050507f20100394950e66014c25009b45d12b675210a6e7a002044a0e3de6544e3c4b3781604051610ac49190611698565b60405180910390a150565b61271061ffff82161115610af6576040516307336f0360e11b815260040160405180910390fd5b6065805461ffff191661ffff83169081179091556040519081527ff5d1836df8fcd7c1e54047e94ac8773d2855395603e2ef9ba5f5f16905f2259290602001610ac4565b6033546001600160a01b031633146107c35760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016103b2565b60605f610ba083610ce1565b6040805160208082528183019092529192505f91906020820181803683375050509182525060208101929092525090565b604080518082019091525f8082526020820152604080518082019091525f8082526020820181905290610c0684860186611760565b919790965090945092505050565b63ffffffff86165f908152606760205260409020548514610c485760405163639d09b560e11b815260040160405180910390fd5b610c9083838080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152508992508591505063ffffffff8816610d0e565b610cad5760405163afa42ca760e01b815260040160405180910390fd5b505050505050565b610cbd610dbc565b610cc98284018461182e565b9695505050505050565b6060610cc982840184611910565b5f60ff8216601f811115610d0857604051632cd44ac360e21b815260040160405180910390fd5b92915050565b5f83610d1b868585610d25565b1495945050505050565b5f60208451610d349190611a44565b15610d52576040516313717da960e21b815260040160405180910390fd5b8260205b85518111610db357610d69600285611a44565b5f03610d8a57815f528086015160205260405f209150600284049350610da1565b808601515f528160205260405f2091506002840493505b610dac602082611a63565b9050610d56565b50949350505050565b60405180608001604052805f81526020015f8152602001610dee60405180604001604052805f81526020015f81525090565b8152602001606081525090565b6001600160a01b03811681146104f1575f5ffd5b5f60408284031215610e1f575f5ffd5b50919050565b803561ffff811681146107ad575f5ffd5b63ffffffff811681146104f1575f5ffd5b5f5f5f5f5f5f6101008789031215610e5d575f5ffd5b8635610e6881610dfb565b9550610e778860208901610e0f565b9450610e8560608801610e25565b93506080870135610e9581610e36565b925060a08701356001600160401b03811115610eaf575f5ffd5b870160a0818a031215610ec0575f5ffd5b9150610ecf8860c08901610e0f565b90509295509295509295565b5f60408284031215610eeb575f5ffd5b610ef58383610e0f565b9392505050565b5f60208284031215610f0c575f5ffd5b610ef582610e25565b80516001600160a01b0316825260209081015163ffffffff16910152565b60408101610d088284610f15565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b5f5f5f60608486031215610f88575f5ffd5b83356001600160401b03811115610f9d575f5ffd5b84016101208187031215610faf575f5ffd5b9250602084013591506040840135610fc681610e36565b809150509250925092565b8035600381106107ad575f5ffd5b5f60208284031215610fef575f5ffd5b610ef582610fd1565b5f5f83601f840112611008575f5ffd5b5081356001600160401b0381111561101e575f5ffd5b602083019150836020828501011115611035575f5ffd5b9250929050565b5f5f5f5f5f5f5f60a0888a031215611052575f5ffd5b873561105d81610e36565b965060208801359550604088013561107481610e36565b945060608801356001600160401b0381111561108e575f5ffd5b61109a8a828b01610ff8565b90955093505060808801356001600160401b038111156110b8575f5ffd5b6110c48a828b01610ff8565b989b979a50959850939692959293505050565b5f602082840312156110e7575f5ffd5b8135610ef581610e36565b5f60208284031215611102575f5ffd5b8135610ef581610dfb565b803561111881610dfb565b6001600160a01b03168252602081013561113181610e36565b63ffffffff81166020840152505050565b5f5f8335601e19843603018112611157575f5ffd5b83016020810192503590506001600160401b03811115611175575f5ffd5b8060051b3603821315611035575f5ffd5b8183525f6001600160fb1b0383111561119d575f5ffd5b8260051b80836020870137939093016020019392505050565b6111c0818661110d565b63ffffffff841660408281019190915260c06060808401829052853591840191909152602085013560e0840152908401356101008301528301356101208201525f61120e6080850185611142565b60a061014085015261122561016085018284611186565b92505050611236608083018461110d565b95945050505050565b634e487b7160e01b5f52604160045260245ffd5b634e487b7160e01b5f52603260045260245ffd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b5f8235605e198336030181126112a3575f5ffd5b90910192915050565b80358252602080820135908301525f6112c86040830183611142565b60606040860152611236606086018284611186565b5f8151808452602084019350602083015f5b8281101561131157815161ffff168652602095860195909101906001016112ef565b5093949350505050565b83546001600160a01b038116825260a01c63ffffffff166020820152608060408201525f6101a08201843561134f81610e36565b63ffffffff166080840152602085013560a0840152604085013560c0840152606085013560e0840152604060808601610100850137604060c0860161014085013761139e610100860186611142565b610120610180860152828184526101c0860190506101c08260051b8701019350825f5b83811015611480578786036101bf190183526113dd828661128f565b80356113e881610e36565b63ffffffff168752602081013536829003601e19018112611407575f5ffd5b81016020810190356001600160401b03811115611422575f5ffd5b803603821315611430575f5ffd5b606060208a015261144560608a018284611267565b915050611455604083018361128f565b9150878103604089015261146981836112ac565b9750505060209283019291909101906001016113c1565b50505050508281036060840152610cc981856112dd565b5f602082840312156114a7575f5ffd5b81518015158114610ef5575f5ffd5b634e487b7160e01b5f52602160045260245ffd5b5f602082840312156114da575f5ffd5b8151610ef581610e36565b818382375f9101908152919050565b5f8151808452602084019350602083015f5b82811015611311578151865260209586019590910190600101611506565b61152e8186610f15565b63ffffffff8416604082015260c06060820152825160c0820152602083015160e08201525f60408401518051610100840152602081015161012084015250606084015160a06101408401526115876101608401826114f4565b9150506112366080830184610f15565b5f60c082016115a68388610f15565b63ffffffff8616604084015260c0606084015280855180835260e08501915060e08160051b8601019250602087015f5b828110156116275786850360df19018452815180516001600160a01b03168652602090810151604091870182905290611611908701826114f4565b95505060209384019391909101906001016115d6565b50505050809150506112366080830184610f15565b813561164781610dfb565b81546001600160a01b031981166001600160a01b03929092169182178355602084013561167381610e36565b6001600160c01b03199190911690911760a09190911b63ffffffff60a01b1617905550565b60408101610d08828461110d565b604080519081016001600160401b03811182821017156116c8576116c861123f565b60405290565b604051608081016001600160401b03811182821017156116c8576116c861123f565b604051601f8201601f191681016001600160401b03811182821017156117185761171861123f565b604052919050565b5f60408284031215611730575f5ffd5b6117386116a6565b9050813561174581610dfb565b8152602082013561175581610e36565b602082015292915050565b5f5f5f60a08486031215611772575f5ffd5b61177c8585611720565b925061178a60408501610fd1565b91506117998560608601611720565b90509250925092565b5f6001600160401b038211156117ba576117ba61123f565b5060051b60200190565b5f82601f8301126117d3575f5ffd5b81356117e66117e1826117a2565b6116f0565b8082825260208201915060208360051b860101925085831115611807575f5ffd5b602085015b8381101561182457803583526020928301920161180c565b5095945050505050565b5f5f5f5f60c08587031215611841575f5ffd5b61184b8686611720565b935061185960408601610fd1565b92506118688660608701611720565b915060a08501356001600160401b03811115611882575f5ffd5b850180870360a0811215611894575f5ffd5b61189c6116ce565b82358152602080840135908201526040603f19830112156118bb575f5ffd5b6118c36116a6565b604084810135825260608501356020830152820152608083013591506001600160401b038211156118f2575f5ffd5b6118fe898385016117c4565b60608201529598949750929550505050565b5f5f5f5f60c08587031215611923575f5ffd5b61192d8686611720565b935061193b60408601610fd1565b925061194a8660608701611720565b915060a08501356001600160401b03811115611964575f5ffd5b8501601f81018713611974575f5ffd5b80356119826117e1826117a2565b8082825260208201915060208360051b8501019250898311156119a3575f5ffd5b602084015b83811015611a355780356001600160401b038111156119c5575f5ffd5b85016040818d03601f190112156119da575f5ffd5b6119e26116a6565b60208201356119f081610dfb565b815260408201356001600160401b03811115611a0a575f5ffd5b611a198e6020838601016117c4565b60208301525080855250506020830192506020810190506119a8565b50969995985093965050505050565b5f82611a5e57634e487b7160e01b5f52601260045260245ffd5b500690565b80820180821115610d0857634e487b7160e01b5f52601160045260245ffdfea2646970667358221220574b7cf8d5fde8de74e9d16698bff854f7ee35f6e1951d56baa1445caf79c59c64736f6c634300081b0033",
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

// ConfirmGlobalTableRoot is a paid mutator transaction binding the contract method 0x6ab40904.
//
// Solidity: function confirmGlobalTableRoot((uint32,bytes32,(uint256,uint256),(uint256[2],uint256[2]),(uint32,bytes,((uint256,uint256),uint256[]))[]) globalTableRootCert, bytes32 globalTableRoot, uint32 referenceTimestamp) returns()
func (_OperatorTableUpdater *OperatorTableUpdaterTransactor) ConfirmGlobalTableRoot(opts *bind.TransactOpts, globalTableRootCert IBN254CertificateVerifierTypesBN254Certificate, globalTableRoot [32]byte, referenceTimestamp uint32) (*types.Transaction, error) {
	return _OperatorTableUpdater.contract.Transact(opts, "confirmGlobalTableRoot", globalTableRootCert, globalTableRoot, referenceTimestamp)
}

// ConfirmGlobalTableRoot is a paid mutator transaction binding the contract method 0x6ab40904.
//
// Solidity: function confirmGlobalTableRoot((uint32,bytes32,(uint256,uint256),(uint256[2],uint256[2]),(uint32,bytes,((uint256,uint256),uint256[]))[]) globalTableRootCert, bytes32 globalTableRoot, uint32 referenceTimestamp) returns()
func (_OperatorTableUpdater *OperatorTableUpdaterSession) ConfirmGlobalTableRoot(globalTableRootCert IBN254CertificateVerifierTypesBN254Certificate, globalTableRoot [32]byte, referenceTimestamp uint32) (*types.Transaction, error) {
	return _OperatorTableUpdater.Contract.ConfirmGlobalTableRoot(&_OperatorTableUpdater.TransactOpts, globalTableRootCert, globalTableRoot, referenceTimestamp)
}

// ConfirmGlobalTableRoot is a paid mutator transaction binding the contract method 0x6ab40904.
//
// Solidity: function confirmGlobalTableRoot((uint32,bytes32,(uint256,uint256),(uint256[2],uint256[2]),(uint32,bytes,((uint256,uint256),uint256[]))[]) globalTableRootCert, bytes32 globalTableRoot, uint32 referenceTimestamp) returns()
func (_OperatorTableUpdater *OperatorTableUpdaterTransactorSession) ConfirmGlobalTableRoot(globalTableRootCert IBN254CertificateVerifierTypesBN254Certificate, globalTableRoot [32]byte, referenceTimestamp uint32) (*types.Transaction, error) {
	return _OperatorTableUpdater.Contract.ConfirmGlobalTableRoot(&_OperatorTableUpdater.TransactOpts, globalTableRootCert, globalTableRoot, referenceTimestamp)
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
