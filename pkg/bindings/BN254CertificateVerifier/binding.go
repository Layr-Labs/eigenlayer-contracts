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

// BN254CertificateVerifierMetaData contains all meta data concerning the BN254CertificateVerifier contract.
var BN254CertificateVerifierMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_operatorTableUpdater\",\"type\":\"address\",\"internalType\":\"contractIOperatorTableUpdater\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getOperatorInfo\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBN254TableCalculatorTypes.BN254OperatorInfo\",\"components\":[{\"name\":\"pubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorSetInfo\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBN254TableCalculatorTypes.BN254OperatorSetInfo\",\"components\":[{\"name\":\"operatorInfoTreeRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"numOperators\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"aggregatePubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"totalWeights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorSetOwner\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestReferenceTimestamp\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxOperatorTableStaleness\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorTableUpdater\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIOperatorTableUpdater\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateOperatorTable\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorSetInfo\",\"type\":\"tuple\",\"internalType\":\"structIBN254TableCalculatorTypes.BN254OperatorSetInfo\",\"components\":[{\"name\":\"operatorInfoTreeRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"numOperators\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"aggregatePubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"totalWeights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]},{\"name\":\"operatorSetConfig\",\"type\":\"tuple\",\"internalType\":\"structICrossChainRegistryTypes.OperatorSetConfig\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxStalenessPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyCertificate\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"cert\",\"type\":\"tuple\",\"internalType\":\"structIBN254CertificateVerifierTypes.BN254Certificate\",\"components\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"nonSignerWitnesses\",\"type\":\"tuple[]\",\"internalType\":\"structIBN254CertificateVerifierTypes.BN254OperatorInfoWitness[]\",\"components\":[{\"name\":\"operatorIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorInfoProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"operatorInfo\",\"type\":\"tuple\",\"internalType\":\"structIBN254TableCalculatorTypes.BN254OperatorInfo\",\"components\":[{\"name\":\"pubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}]}]}],\"outputs\":[{\"name\":\"signedStakes\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyCertificateNominal\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"cert\",\"type\":\"tuple\",\"internalType\":\"structIBN254CertificateVerifierTypes.BN254Certificate\",\"components\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"nonSignerWitnesses\",\"type\":\"tuple[]\",\"internalType\":\"structIBN254CertificateVerifierTypes.BN254OperatorInfoWitness[]\",\"components\":[{\"name\":\"operatorIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorInfoProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"operatorInfo\",\"type\":\"tuple\",\"internalType\":\"structIBN254TableCalculatorTypes.BN254OperatorInfo\",\"components\":[{\"name\":\"pubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}]}]},{\"name\":\"totalStakeNominalThresholds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyCertificateProportion\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"cert\",\"type\":\"tuple\",\"internalType\":\"structIBN254CertificateVerifierTypes.BN254Certificate\",\"components\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"nonSignerWitnesses\",\"type\":\"tuple[]\",\"internalType\":\"structIBN254CertificateVerifierTypes.BN254OperatorInfoWitness[]\",\"components\":[{\"name\":\"operatorIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorInfoProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"operatorInfo\",\"type\":\"tuple\",\"internalType\":\"structIBN254TableCalculatorTypes.BN254OperatorInfo\",\"components\":[{\"name\":\"pubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}]}]},{\"name\":\"totalStakeProportionThresholds\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MaxStalenessPeriodUpdated\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"maxStalenessPeriod\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSetOwnerUpdated\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TableUpdated\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"operatorSetInfo\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIBN254TableCalculatorTypes.BN254OperatorSetInfo\",\"components\":[{\"name\":\"operatorInfoTreeRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"numOperators\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"aggregatePubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"totalWeights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ArrayLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CertificateStale\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECAddFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECMulFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpModFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperatorIndex\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProofLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyTableUpdater\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReferenceTimestampDoesNotExist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TableUpdateStale\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"VerificationFailed\",\"inputs\":[]}]",
	Bin: "0x60a060405234801561000f575f5ffd5b5060405161232b38038061232b83398101604081905261002e91610105565b6001600160a01b038116608052610043610049565b50610132565b5f54610100900460ff16156100b45760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff90811614610103575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b5f60208284031215610115575f5ffd5b81516001600160a01b038116811461012b575f5ffd5b9392505050565b6080516121da6101515f395f818161013c01526103d901526121da5ff3fe608060405234801561000f575f5ffd5b506004361061009b575f3560e01c806368d6e0811161006357806368d6e081146101375780638481892014610176578063dd2ae1b914610189578063e49613fc1461019c578063eb39e68f146101bc575f5ffd5b8063017d79741461009f578063080b7150146100c75780635ddb9b5b146100e75780636141879e1461010f5780636738c40b14610122575b5f5ffd5b6100b26100ad366004611c28565b6101dc565b60405190151581526020015b60405180910390f35b6100da6100d5366004611d03565b61036d565b6040516100be9190611d4e565b6100fa6100f5366004611d85565b610382565b60405163ffffffff90911681526020016100be565b6100fa61011d366004611d85565b6103a8565b610135610130366004611db5565b6103ce565b005b61015e7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016100be565b61015e610184366004611d85565b61059b565b6100b2610197366004611e79565b6105c4565b6101af6101aa366004611eec565b610657565b6040516100be9190611f61565b6101cf6101ca366004611f96565b61070f565b6040516100be919061200d565b5f5f6101e885856107d4565b90505f6101f486610994565b5f8181526004602081815260408084208a5163ffffffff16855282528084208151608081018352815481526001820154818501528251808401845260028301548152600383015481860152818401529381018054835181860281018601909452808452969750949593949093606086019383018282801561029257602002820191905f5260205f20905b81548152602001906001019080831161027e575b50505050508152505090505f8160600151905085518451146102c75760405163512509d360e11b815260040160405180910390fd5b5f5b845181101561035c575f6127108883815181106102e8576102e861201f565b602002602001015161ffff168484815181106103065761030661201f565b60200260200101516103189190612047565b6103229190612072565b9050808683815181106103375761033761201f565b60200260200101511015610353575f9650505050505050610366565b506001016102c9565b5060019450505050505b9392505050565b606061037983836107d4565b90505b92915050565b5f5f61038d83610994565b5f9081526003602052604090205463ffffffff169392505050565b5f5f6103b383610994565b5f9081526002602052604090205463ffffffff169392505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146104175760405163030c1b6b60e11b815260040160405180910390fd5b5f61042f61042a36879003870187611d85565b610994565b5f8181526003602052604090205490915063ffffffff9081169085161161046957604051632f20889f60e01b815260040160405180910390fd5b5f81815260046020818152604080842063ffffffff891685528252928390208651815581870151600182015592860151805160028501558101516003840155606086015180518794936104c1939085019201906115be565b5050505f818152600360209081526040909120805463ffffffff191663ffffffff87161790556104f390830183612085565b5f8281526001602090815260409182902080546001600160a01b0319166001600160a01b0394909416939093179092556105329190840190840161209e565b5f8281526002602052604090819020805463ffffffff191663ffffffff9390931692909217909155517f93e6bea1c9b5dce4a5c07b00261e956df2a4a253d9ab6ca070ca2037d72ada9e9061058c908790879087906120b7565b60405180910390a15050505050565b5f5f6105a683610994565b5f908152600160205260409020546001600160a01b03169392505050565b5f5f6105d085856107d4565b905082518151146105f45760405163512509d360e11b815260040160405180910390fd5b5f5b815181101561064b578381815181106106115761061161201f565b602002602001015182828151811061062b5761062b61201f565b60200260200101511015610643575f92505050610366565b6001016105f6565b50600195945050505050565b61065f611607565b5f61066985610994565b5f81815260056020908152604080832063ffffffff8916845282528083208784528252918290208251608081018452815481850190815260018301546060830152815260028201805485518186028101860190965280865295965090949193858401939092908301828280156106fc57602002820191905f5260205f20905b8154815260200190600101908083116106e8575b5050505050815250509150509392505050565b610717611631565b5f61072184610994565b5f81815260046020818152604080842063ffffffff89168552825292839020835160808101855281548152600182015481840152845180860186526002830154815260038301548185015281860152928101805485518185028101850190965280865295965092949093606086019390929091908301828280156107c257602002820191905f5260205f20905b8154815260200190600101908083116107ae575b50505050508152505091505092915050565b60606107de611663565b6107e784610994565b80825283516107f691906109f7565b80515f908152600460208181526040808420875163ffffffff168552825292839020835160808101855281548152600182015481840152845180860186526002830154815260038301548185015281860152928101805485518185028101850190965280865293949193606086019383018282801561089257602002820191905f5260205f20905b81548152602001906001019080831161087e575b5050509190925250505060208201819052516108c157604051630cad17b760e31b815260040160405180910390fd5b806020015160600151516001600160401b038111156108e2576108e2611791565b60405190808252806020026020018201604052801561090b578160200160208202803683370190505b5060408201525f5b8160200151606001515181101561096f57816020015160600151818151811061093e5761093e61201f565b60200260200101518260400151828151811061095c5761095c61201f565b6020908102919091010152600101610913565b5061097a8184610a49565b60608201526109898184610b73565b604001519392505050565b5f815f0151826020015163ffffffff166040516020016109df92919060609290921b6bffffffffffffffffffffffff1916825260a01b6001600160a01b031916601482015260200190565b60405160208183030381529060405261037c9061210c565b5f8281526002602052604090205463ffffffff16801580610a275750610a1d818361212f565b63ffffffff164211155b610a445760405163640fcd6b60e11b815260040160405180910390fd5b505050565b6040805180820182525f808252602091820181905282518084019093528083529082018190525b826080015151811015610b6c575f83608001518281518110610a9457610a9461201f565b60200260200101519050846020015160200151815f015163ffffffff1610610acf576040516301fa53c760e11b815260040160405180910390fd5b845184515f91610adf9184610be1565b8051909150610aef908590610d56565b93505f5b816020015151811015610b6157866040015151811015610b595781602001518181518110610b2357610b2361201f565b602002602001015187604001518281518110610b4157610b4161201f565b60200260200101818151610b55919061214b565b9052505b600101610af3565b505050600101610a70565b5092915050565b5f610b93610b848460600151610dd2565b60208501516040015190610d56565b90505f5f610baf84602001518486606001518760400151610e68565b91509150818015610bbd5750805b610bda5760405163439cc0cd60e01b815260040160405180910390fd5b5050505050565b610be9611607565b5f84815260056020908152604080832063ffffffff80881685529083528184208651909116845290915281208054909190151580610c2a5750600182015415155b905080610cd3575f610c4a8787875f015188604001518960200151610fea565b905080610c6a5760405163439cc0cd60e01b815260040160405180910390fd5b6040808601515f8981526005602090815283822063ffffffff808c1684529082528483208a5190911683528152929020815180518255830151600182015582820151805192939192610cc292600285019201906115be565b509050508460400151935050610d4d565b6040805160808101825283548183019081526001850154606083015281526002840180548351602082810282018101909552818152929386938186019390929091830182828015610d4157602002820191905f5260205f20905b815481526020019060010190808311610d2d575b50505050508152505092505b50509392505050565b604080518082019091525f8082526020820152610d716116a8565b835181526020808501518183015283516040808401919091529084015160608301525f908360808460066107d05a03fa90508080610dab57fe5b5080610dca5760405163d4b68fd760e01b815260040160405180910390fd5b505092915050565b604080518082019091525f80825260208201528151158015610df657506020820151155b15610e13575050604080518082019091525f808252602082015290565b6040518060400160405280835f015181526020015f5160206121855f395f51905f528460200151610e44919061215e565b610e5b905f5160206121855f395f51905f5261214b565b905292915050565b919050565b5f5f5f7f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000187875f01518860200151885f01515f60028110610eab57610eab61201f565b60200201518951600160200201518a602001515f60028110610ecf57610ecf61201f565b60200201518b60200151600160028110610eeb57610eeb61201f565b602090810291909101518c518d830151604051610f489a99989796959401988952602089019790975260408801959095526060870193909352608086019190915260a085015260c084015260e08301526101008201526101200190565b604051602081830303815290604052805190602001205f1c610f6a919061215e565b9050610fdc610f83610f7c8884611055565b8690610d56565b610f8b6110bd565b610fd2610fc385610fbd6040805180820182525f80825260209182015281518083019092526001825260029082015290565b90611055565b610fcc8c61117d565b90610d56565b8862061a80611207565b909890975095505050505050565b5f5f83604051602001610ffd9190611f61565b60408051601f1981840301815291815281516020928301205f8a81526004845282812063ffffffff808c1683529452919091205490925090611049908590839085908a81169061141b16565b98975050505050505050565b604080518082019091525f80825260208201526110706116c6565b835181526020808501519082015260408082018490525f908360608460076107d05a03fa9050808061109e57fe5b5080610dca57604051632319df1960e11b815260040160405180910390fd5b6110c56116e4565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208381019190915281019190915290565b604080518082019091525f80825260208201525f80806111aa5f5160206121855f395f51905f528661215e565b90505b6111b681611432565b90935091505f5160206121855f395f51905f5282830983036111ee576040805180820190915290815260208101919091529392505050565b5f5160206121855f395f51905f526001820890506111ad565b6040805180820182528681526020808201869052825180840190935286835282018490525f91829190611238611704565b5f5b60028110156113ef575f61124f826006612047565b90508482600281106112635761126361201f565b60200201515183611274835f612171565b600c81106112845761128461201f565b602002015284826002811061129b5761129b61201f565b602002015160200151838260016112b29190612171565b600c81106112c2576112c261201f565b60200201528382600281106112d9576112d961201f565b60200201515151836112ec836002612171565b600c81106112fc576112fc61201f565b60200201528382600281106113135761131361201f565b602002015151600160200201518361132c836003612171565b600c811061133c5761133c61201f565b60200201528382600281106113535761135361201f565b6020020151602001515f6002811061136d5761136d61201f565b60200201518361137e836004612171565b600c811061138e5761138e61201f565b60200201528382600281106113a5576113a561201f565b6020020151602001516001600281106113c0576113c061201f565b6020020151836113d1836005612171565b600c81106113e1576113e161201f565b60200201525060010161123a565b506113f8611723565b5f6020826101808560088cfa9151919c9115159b50909950505050505050505050565b5f836114288685856114ae565b1495945050505050565b5f80805f5160206121855f395f51905f5260035f5160206121855f395f51905f52865f5160206121855f395f51905f52888909090890505f6114a2827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f525f5160206121855f395f51905f52611545565b91959194509092505050565b5f602084516114bd919061215e565b156114db576040516313717da960e21b815260040160405180910390fd5b8260205b8551811161153c576114f260028561215e565b5f0361151357815f528086015160205260405f20915060028404935061152a565b808601515f528160205260405f2091506002840493505b611535602082612171565b90506114df565b50949350505050565b5f5f61154f611723565b611557611741565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa9250828061159457fe5b50826115b35760405163d51edae360e01b815260040160405180910390fd5b505195945050505050565b828054828255905f5260205f209081019282156115f7579160200282015b828111156115f75782518255916020019190600101906115dc565b5061160392915061175f565b5090565b604080516080810182525f91810182815260608201929092529081905b8152602001606081525090565b60405180608001604052805f81526020015f815260200161162460405180604001604052805f81526020015f81525090565b60405180608001604052805f815260200161167c611631565b8152602001606081526020016116a360405180604001604052805f81526020015f81525090565b905290565b60405180608001604052806004906020820280368337509192915050565b60405180606001604052806003906020820280368337509192915050565b60405180604001604052806116f7611773565b81526020016116a3611773565b604051806101800160405280600c906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b5b80821115611603575f8155600101611760565b60405180604001604052806002906020820280368337509192915050565b634e487b7160e01b5f52604160045260245ffd5b604080519081016001600160401b03811182821017156117c7576117c7611791565b60405290565b604051606081016001600160401b03811182821017156117c7576117c7611791565b60405160a081016001600160401b03811182821017156117c7576117c7611791565b604051608081016001600160401b03811182821017156117c7576117c7611791565b604051601f8201601f191681016001600160401b038111828210171561185b5761185b611791565b604052919050565b80356001600160a01b0381168114610e63575f5ffd5b803563ffffffff81168114610e63575f5ffd5b5f6040828403121561189c575f5ffd5b6118a46117a5565b90506118af82611863565b81526118bd60208301611879565b602082015292915050565b5f604082840312156118d8575f5ffd5b6118e06117a5565b823581526020928301359281019290925250919050565b5f82601f830112611906575f5ffd5b61190e6117a5565b80604084018581111561191f575f5ffd5b845b81811015611939578035845260209384019301611921565b509095945050505050565b5f6001600160401b0382111561195c5761195c611791565b5060051b60200190565b5f82601f830112611975575f5ffd5b813561198861198382611944565b611833565b8082825260208201915060208360051b8601019250858311156119a9575f5ffd5b602085015b838110156119c65780358352602092830192016119ae565b5095945050505050565b5f606082840312156119e0575f5ffd5b6119e86117a5565b90506119f483836118c8565b815260408201356001600160401b03811115611a0e575f5ffd5b611a1a84828501611966565b60208301525092915050565b5f82601f830112611a35575f5ffd5b8135611a4361198382611944565b8082825260208201915060208360051b860101925085831115611a64575f5ffd5b602085015b838110156119c65780356001600160401b03811115611a86575f5ffd5b86016060818903601f19011215611a9b575f5ffd5b611aa36117cd565b611aaf60208301611879565b815260408201356001600160401b03811115611ac9575f5ffd5b82016020810190603f018a13611add575f5ffd5b80356001600160401b03811115611af657611af6611791565b611b09601f8201601f1916602001611833565b8181528b6020838501011115611b1d575f5ffd5b816020840160208301375f6020838301015280602085015250505060608201356001600160401b03811115611b50575f5ffd5b611b5f8a6020838601016119d0565b60408301525084525060209283019201611a69565b5f818303610120811215611b86575f5ffd5b611b8e6117ef565b9150611b9983611879565b825260208381013590830152611bb284604085016118c8565b60408301526080607f1982011215611bc8575f5ffd5b50611bd16117a5565b611bde84608085016118f7565b8152611bed8460c085016118f7565b602082015260608201526101008201356001600160401b03811115611c10575f5ffd5b611c1c84828501611a26565b60808301525092915050565b5f5f5f60808486031215611c3a575f5ffd5b611c44858561188c565b925060408401356001600160401b03811115611c5e575f5ffd5b611c6a86828701611b74565b92505060608401356001600160401b03811115611c85575f5ffd5b8401601f81018613611c95575f5ffd5b8035611ca361198382611944565b8082825260208201915060208360051b850101925088831115611cc4575f5ffd5b6020840193505b82841015611cf557833561ffff81168114611ce4575f5ffd5b825260209384019390910190611ccb565b809450505050509250925092565b5f5f60608385031215611d14575f5ffd5b611d1e848461188c565b915060408301356001600160401b03811115611d38575f5ffd5b611d4485828601611b74565b9150509250929050565b602080825282518282018190525f918401906040840190835b81811015611939578351835260209384019390920191600101611d67565b5f60408284031215611d95575f5ffd5b610379838361188c565b5f60408284031215611daf575f5ffd5b50919050565b5f5f5f5f60c08587031215611dc8575f5ffd5b611dd28686611d9f565b9350611de060408601611879565b925060608501356001600160401b03811115611dfa575f5ffd5b850160a08188031215611e0b575f5ffd5b611e13611811565b8135815260208083013590820152611e2e88604084016118c8565b604082015260808201356001600160401b03811115611e4b575f5ffd5b611e5789828501611966565b6060830152509250611e6e90508660808701611d9f565b905092959194509250565b5f5f5f60808486031215611e8b575f5ffd5b611e95858561188c565b925060408401356001600160401b03811115611eaf575f5ffd5b611ebb86828701611b74565b92505060608401356001600160401b03811115611ed6575f5ffd5b611ee286828701611966565b9150509250925092565b5f5f5f60808486031215611efe575f5ffd5b611f08858561188c565b9250611f1660408501611879565b929592945050506060919091013590565b5f8151808452602084019350602083015f5b82811015611f57578151865260209586019590910190600101611f39565b5093949350505050565b60208082528251805183830152015160408201525f6020830151606080840152611f8e6080840182611f27565b949350505050565b5f5f60608385031215611fa7575f5ffd5b611fb1848461188c565b9150611fbf60408401611879565b90509250929050565b80518252602081015160208301525f6040820151611ff3604085018280518252602090810151910152565b50606082015160a06080850152611f8e60a0850182611f27565b602081525f6103796020830184611fc8565b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52601160045260245ffd5b808202811582820484141761037c5761037c612033565b634e487b7160e01b5f52601260045260245ffd5b5f826120805761208061205e565b500490565b5f60208284031215612095575f5ffd5b61037982611863565b5f602082840312156120ae575f5ffd5b61037982611879565b6001600160a01b036120c885611863565b16815263ffffffff6120dc60208601611879565b16602082015263ffffffff83166040820152608060608201525f6121036080830184611fc8565b95945050505050565b80516020808301519190811015611daf575f1960209190910360031b1b16919050565b63ffffffff818116838216019081111561037c5761037c612033565b8181038181111561037c5761037c612033565b5f8261216c5761216c61205e565b500690565b8082018082111561037c5761037c61203356fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47a2646970667358221220f6c03ec5f38d3ed4d9f0182a0319f0502da0f9e64a11006556a88af8f5d797f164736f6c634300081b0033",
}

// BN254CertificateVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use BN254CertificateVerifierMetaData.ABI instead.
var BN254CertificateVerifierABI = BN254CertificateVerifierMetaData.ABI

// BN254CertificateVerifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BN254CertificateVerifierMetaData.Bin instead.
var BN254CertificateVerifierBin = BN254CertificateVerifierMetaData.Bin

// DeployBN254CertificateVerifier deploys a new Ethereum contract, binding an instance of BN254CertificateVerifier to it.
func DeployBN254CertificateVerifier(auth *bind.TransactOpts, backend bind.ContractBackend, _operatorTableUpdater common.Address) (common.Address, *types.Transaction, *BN254CertificateVerifier, error) {
	parsed, err := BN254CertificateVerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BN254CertificateVerifierBin), backend, _operatorTableUpdater)
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

// GetOperatorInfo is a free data retrieval call binding the contract method 0xe49613fc.
//
// Solidity: function getOperatorInfo((address,uint32) operatorSet, uint32 referenceTimestamp, uint256 operatorIndex) view returns(((uint256,uint256),uint256[]))
func (_BN254CertificateVerifier *BN254CertificateVerifierCaller) GetOperatorInfo(opts *bind.CallOpts, operatorSet OperatorSet, referenceTimestamp uint32, operatorIndex *big.Int) (IBN254TableCalculatorTypesBN254OperatorInfo, error) {
	var out []interface{}
	err := _BN254CertificateVerifier.contract.Call(opts, &out, "getOperatorInfo", operatorSet, referenceTimestamp, operatorIndex)

	if err != nil {
		return *new(IBN254TableCalculatorTypesBN254OperatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IBN254TableCalculatorTypesBN254OperatorInfo)).(*IBN254TableCalculatorTypesBN254OperatorInfo)

	return out0, err

}

// GetOperatorInfo is a free data retrieval call binding the contract method 0xe49613fc.
//
// Solidity: function getOperatorInfo((address,uint32) operatorSet, uint32 referenceTimestamp, uint256 operatorIndex) view returns(((uint256,uint256),uint256[]))
func (_BN254CertificateVerifier *BN254CertificateVerifierSession) GetOperatorInfo(operatorSet OperatorSet, referenceTimestamp uint32, operatorIndex *big.Int) (IBN254TableCalculatorTypesBN254OperatorInfo, error) {
	return _BN254CertificateVerifier.Contract.GetOperatorInfo(&_BN254CertificateVerifier.CallOpts, operatorSet, referenceTimestamp, operatorIndex)
}

// GetOperatorInfo is a free data retrieval call binding the contract method 0xe49613fc.
//
// Solidity: function getOperatorInfo((address,uint32) operatorSet, uint32 referenceTimestamp, uint256 operatorIndex) view returns(((uint256,uint256),uint256[]))
func (_BN254CertificateVerifier *BN254CertificateVerifierCallerSession) GetOperatorInfo(operatorSet OperatorSet, referenceTimestamp uint32, operatorIndex *big.Int) (IBN254TableCalculatorTypesBN254OperatorInfo, error) {
	return _BN254CertificateVerifier.Contract.GetOperatorInfo(&_BN254CertificateVerifier.CallOpts, operatorSet, referenceTimestamp, operatorIndex)
}

// GetOperatorSetInfo is a free data retrieval call binding the contract method 0xeb39e68f.
//
// Solidity: function getOperatorSetInfo((address,uint32) operatorSet, uint32 referenceTimestamp) view returns((bytes32,uint256,(uint256,uint256),uint256[]))
func (_BN254CertificateVerifier *BN254CertificateVerifierCaller) GetOperatorSetInfo(opts *bind.CallOpts, operatorSet OperatorSet, referenceTimestamp uint32) (IBN254TableCalculatorTypesBN254OperatorSetInfo, error) {
	var out []interface{}
	err := _BN254CertificateVerifier.contract.Call(opts, &out, "getOperatorSetInfo", operatorSet, referenceTimestamp)

	if err != nil {
		return *new(IBN254TableCalculatorTypesBN254OperatorSetInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IBN254TableCalculatorTypesBN254OperatorSetInfo)).(*IBN254TableCalculatorTypesBN254OperatorSetInfo)

	return out0, err

}

// GetOperatorSetInfo is a free data retrieval call binding the contract method 0xeb39e68f.
//
// Solidity: function getOperatorSetInfo((address,uint32) operatorSet, uint32 referenceTimestamp) view returns((bytes32,uint256,(uint256,uint256),uint256[]))
func (_BN254CertificateVerifier *BN254CertificateVerifierSession) GetOperatorSetInfo(operatorSet OperatorSet, referenceTimestamp uint32) (IBN254TableCalculatorTypesBN254OperatorSetInfo, error) {
	return _BN254CertificateVerifier.Contract.GetOperatorSetInfo(&_BN254CertificateVerifier.CallOpts, operatorSet, referenceTimestamp)
}

// GetOperatorSetInfo is a free data retrieval call binding the contract method 0xeb39e68f.
//
// Solidity: function getOperatorSetInfo((address,uint32) operatorSet, uint32 referenceTimestamp) view returns((bytes32,uint256,(uint256,uint256),uint256[]))
func (_BN254CertificateVerifier *BN254CertificateVerifierCallerSession) GetOperatorSetInfo(operatorSet OperatorSet, referenceTimestamp uint32) (IBN254TableCalculatorTypesBN254OperatorSetInfo, error) {
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

// UpdateOperatorTable is a paid mutator transaction binding the contract method 0x6738c40b.
//
// Solidity: function updateOperatorTable((address,uint32) operatorSet, uint32 referenceTimestamp, (bytes32,uint256,(uint256,uint256),uint256[]) operatorSetInfo, (address,uint32) operatorSetConfig) returns()
func (_BN254CertificateVerifier *BN254CertificateVerifierTransactor) UpdateOperatorTable(opts *bind.TransactOpts, operatorSet OperatorSet, referenceTimestamp uint32, operatorSetInfo IBN254TableCalculatorTypesBN254OperatorSetInfo, operatorSetConfig ICrossChainRegistryTypesOperatorSetConfig) (*types.Transaction, error) {
	return _BN254CertificateVerifier.contract.Transact(opts, "updateOperatorTable", operatorSet, referenceTimestamp, operatorSetInfo, operatorSetConfig)
}

// UpdateOperatorTable is a paid mutator transaction binding the contract method 0x6738c40b.
//
// Solidity: function updateOperatorTable((address,uint32) operatorSet, uint32 referenceTimestamp, (bytes32,uint256,(uint256,uint256),uint256[]) operatorSetInfo, (address,uint32) operatorSetConfig) returns()
func (_BN254CertificateVerifier *BN254CertificateVerifierSession) UpdateOperatorTable(operatorSet OperatorSet, referenceTimestamp uint32, operatorSetInfo IBN254TableCalculatorTypesBN254OperatorSetInfo, operatorSetConfig ICrossChainRegistryTypesOperatorSetConfig) (*types.Transaction, error) {
	return _BN254CertificateVerifier.Contract.UpdateOperatorTable(&_BN254CertificateVerifier.TransactOpts, operatorSet, referenceTimestamp, operatorSetInfo, operatorSetConfig)
}

// UpdateOperatorTable is a paid mutator transaction binding the contract method 0x6738c40b.
//
// Solidity: function updateOperatorTable((address,uint32) operatorSet, uint32 referenceTimestamp, (bytes32,uint256,(uint256,uint256),uint256[]) operatorSetInfo, (address,uint32) operatorSetConfig) returns()
func (_BN254CertificateVerifier *BN254CertificateVerifierTransactorSession) UpdateOperatorTable(operatorSet OperatorSet, referenceTimestamp uint32, operatorSetInfo IBN254TableCalculatorTypesBN254OperatorSetInfo, operatorSetConfig ICrossChainRegistryTypesOperatorSetConfig) (*types.Transaction, error) {
	return _BN254CertificateVerifier.Contract.UpdateOperatorTable(&_BN254CertificateVerifier.TransactOpts, operatorSet, referenceTimestamp, operatorSetInfo, operatorSetConfig)
}

// VerifyCertificate is a paid mutator transaction binding the contract method 0x080b7150.
//
// Solidity: function verifyCertificate((address,uint32) operatorSet, (uint32,bytes32,(uint256,uint256),(uint256[2],uint256[2]),(uint32,bytes,((uint256,uint256),uint256[]))[]) cert) returns(uint256[] signedStakes)
func (_BN254CertificateVerifier *BN254CertificateVerifierTransactor) VerifyCertificate(opts *bind.TransactOpts, operatorSet OperatorSet, cert IBN254CertificateVerifierTypesBN254Certificate) (*types.Transaction, error) {
	return _BN254CertificateVerifier.contract.Transact(opts, "verifyCertificate", operatorSet, cert)
}

// VerifyCertificate is a paid mutator transaction binding the contract method 0x080b7150.
//
// Solidity: function verifyCertificate((address,uint32) operatorSet, (uint32,bytes32,(uint256,uint256),(uint256[2],uint256[2]),(uint32,bytes,((uint256,uint256),uint256[]))[]) cert) returns(uint256[] signedStakes)
func (_BN254CertificateVerifier *BN254CertificateVerifierSession) VerifyCertificate(operatorSet OperatorSet, cert IBN254CertificateVerifierTypesBN254Certificate) (*types.Transaction, error) {
	return _BN254CertificateVerifier.Contract.VerifyCertificate(&_BN254CertificateVerifier.TransactOpts, operatorSet, cert)
}

// VerifyCertificate is a paid mutator transaction binding the contract method 0x080b7150.
//
// Solidity: function verifyCertificate((address,uint32) operatorSet, (uint32,bytes32,(uint256,uint256),(uint256[2],uint256[2]),(uint32,bytes,((uint256,uint256),uint256[]))[]) cert) returns(uint256[] signedStakes)
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
	OperatorSetInfo    IBN254TableCalculatorTypesBN254OperatorSetInfo
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
