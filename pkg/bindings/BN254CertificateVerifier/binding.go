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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_operatorTableUpdater\",\"type\":\"address\",\"internalType\":\"contractIOperatorTableUpdater\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getOperatorInfo\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBN254TableCalculatorTypes.BN254OperatorInfo\",\"components\":[{\"name\":\"pubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorSetInfo\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBN254TableCalculatorTypes.BN254OperatorSetInfo\",\"components\":[{\"name\":\"operatorInfoTreeRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"numOperators\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"aggregatePubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"totalWeights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorSetOwner\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestReferenceTimestamp\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxOperatorTableStaleness\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorTableUpdater\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIOperatorTableUpdater\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setMaxStalenessPeriod\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"maxStalenessPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOperatorTable\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorSetInfo\",\"type\":\"tuple\",\"internalType\":\"structIBN254TableCalculatorTypes.BN254OperatorSetInfo\",\"components\":[{\"name\":\"operatorInfoTreeRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"numOperators\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"aggregatePubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"totalWeights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]},{\"name\":\"operatorSetConfig\",\"type\":\"tuple\",\"internalType\":\"structICrossChainRegistryTypes.OperatorSetConfig\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxStalenessPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyCertificate\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"cert\",\"type\":\"tuple\",\"internalType\":\"structIBN254CertificateVerifierTypes.BN254Certificate\",\"components\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"nonSignerWitnesses\",\"type\":\"tuple[]\",\"internalType\":\"structIBN254CertificateVerifierTypes.BN254OperatorInfoWitness[]\",\"components\":[{\"name\":\"operatorIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorInfoProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"operatorInfo\",\"type\":\"tuple\",\"internalType\":\"structIBN254TableCalculatorTypes.BN254OperatorInfo\",\"components\":[{\"name\":\"pubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}]}]}],\"outputs\":[{\"name\":\"signedStakes\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyCertificateNominal\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"cert\",\"type\":\"tuple\",\"internalType\":\"structIBN254CertificateVerifierTypes.BN254Certificate\",\"components\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"nonSignerWitnesses\",\"type\":\"tuple[]\",\"internalType\":\"structIBN254CertificateVerifierTypes.BN254OperatorInfoWitness[]\",\"components\":[{\"name\":\"operatorIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorInfoProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"operatorInfo\",\"type\":\"tuple\",\"internalType\":\"structIBN254TableCalculatorTypes.BN254OperatorInfo\",\"components\":[{\"name\":\"pubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}]}]},{\"name\":\"totalStakeNominalThresholds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyCertificateProportion\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"cert\",\"type\":\"tuple\",\"internalType\":\"structIBN254CertificateVerifierTypes.BN254Certificate\",\"components\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"nonSignerWitnesses\",\"type\":\"tuple[]\",\"internalType\":\"structIBN254CertificateVerifierTypes.BN254OperatorInfoWitness[]\",\"components\":[{\"name\":\"operatorIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorInfoProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"operatorInfo\",\"type\":\"tuple\",\"internalType\":\"structIBN254TableCalculatorTypes.BN254OperatorInfo\",\"components\":[{\"name\":\"pubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}]}]},{\"name\":\"totalStakeProportionThresholds\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MaxStalenessPeriodUpdated\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"maxStalenessPeriod\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSetOwnerUpdated\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TableUpdated\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"operatorSetInfo\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIBN254TableCalculatorTypes.BN254OperatorSetInfo\",\"components\":[{\"name\":\"operatorInfoTreeRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"numOperators\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"aggregatePubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"totalWeights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ArrayLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CertificateStale\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECAddFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECMulFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpModFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperatorIndex\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProofLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyTableUpdater\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReferenceTimestampDoesNotExist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TableUpdateStale\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"VerificationFailed\",\"inputs\":[]}]",
	Bin: "0x60a060405234801561000f575f5ffd5b5060405161237d38038061237d83398101604081905261002e91610105565b6001600160a01b038116608052610043610049565b50610132565b5f54610100900460ff16156100b45760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff90811614610103575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b5f60208284031215610115575f5ffd5b81516001600160a01b038116811461012b575f5ffd5b9392505050565b60805161222c6101515f395f818161015a015261042b015261222c5ff3fe608060405234801561000f575f5ffd5b50600436106100a6575f3560e01c80636738c40b1161006e5780636738c40b1461014257806368d6e081146101555780638481892014610194578063dd2ae1b9146101a7578063e49613fc146101ba578063eb39e68f146101da575f5ffd5b8063017d7974146100aa578063080b7150146100d25780632a610b75146100f25780635ddb9b5b146101075780636141879e1461012f575b5f5ffd5b6100bd6100b8366004611c7a565b6101fa565b60405190151581526020015b60405180910390f35b6100e56100e0366004611d55565b61038b565b6040516100c99190611da0565b610105610100366004611dd7565b6103a0565b005b61011a610115366004611e09565b6103d4565b60405163ffffffff90911681526020016100c9565b61011a61013d366004611e09565b6103fa565b610105610150366004611e39565b610420565b61017c7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016100c9565b61017c6101a2366004611e09565b6105ed565b6100bd6101b5366004611efd565b610616565b6101cd6101c8366004611f70565b6106a9565b6040516100c99190611fe5565b6101ed6101e8366004611dd7565b610761565b6040516100c9919061205f565b5f5f6102068585610826565b90505f610212866109e6565b5f8181526004602081815260408084208a5163ffffffff1685528252808420815160808101835281548152600182015481850152825180840184526002830154815260038301548186015281840152938101805483518186028101860190945280845296975094959394909360608601938301828280156102b057602002820191905f5260205f20905b81548152602001906001019080831161029c575b50505050508152505090505f8160600151905085518451146102e55760405163512509d360e11b815260040160405180910390fd5b5f5b845181101561037a575f61271088838151811061030657610306612071565b602002602001015161ffff1684848151811061032457610324612071565b60200260200101516103369190612099565b61034091906120c4565b90508086838151811061035557610355612071565b60200260200101511015610371575f9650505050505050610384565b506001016102e7565b5060019450505050505b9392505050565b60606103978383610826565b90505b92915050565b5f6103aa836109e6565b5f908152600260205260409020805463ffffffff191663ffffffff93909316929092179091555050565b5f5f6103df836109e6565b5f9081526003602052604090205463ffffffff169392505050565b5f5f610405836109e6565b5f9081526002602052604090205463ffffffff169392505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146104695760405163030c1b6b60e11b815260040160405180910390fd5b5f61048161047c36879003870187611e09565b6109e6565b5f8181526003602052604090205490915063ffffffff908116908516116104bb57604051632f20889f60e01b815260040160405180910390fd5b5f81815260046020818152604080842063ffffffff8916855282529283902086518155818701516001820155928601518051600285015581015160038401556060860151805187949361051393908501920190611610565b5050505f818152600360209081526040909120805463ffffffff191663ffffffff8716179055610545908301836120d7565b5f8281526001602090815260409182902080546001600160a01b0319166001600160a01b039490941693909317909255610584919084019084016120f0565b5f8281526002602052604090819020805463ffffffff191663ffffffff9390931692909217909155517f93e6bea1c9b5dce4a5c07b00261e956df2a4a253d9ab6ca070ca2037d72ada9e906105de90879087908790612109565b60405180910390a15050505050565b5f5f6105f8836109e6565b5f908152600160205260409020546001600160a01b03169392505050565b5f5f6106228585610826565b905082518151146106465760405163512509d360e11b815260040160405180910390fd5b5f5b815181101561069d5783818151811061066357610663612071565b602002602001015182828151811061067d5761067d612071565b60200260200101511015610695575f92505050610384565b600101610648565b50600195945050505050565b6106b1611659565b5f6106bb856109e6565b5f81815260056020908152604080832063ffffffff89168452825280832087845282529182902082516080810184528154818501908152600183015460608301528152600282018054855181860281018601909652808652959650909491938584019390929083018282801561074e57602002820191905f5260205f20905b81548152602001906001019080831161073a575b5050505050815250509150509392505050565b610769611683565b5f610773846109e6565b5f81815260046020818152604080842063ffffffff891685528252928390208351608081018552815481526001820154818401528451808601865260028301548152600383015481850152818601529281018054855181850281018501909652808652959650929490936060860193909290919083018282801561081457602002820191905f5260205f20905b815481526020019060010190808311610800575b50505050508152505091505092915050565b60606108306116b5565b610839846109e6565b80825283516108489190610a49565b80515f908152600460208181526040808420875163ffffffff16855282529283902083516080810185528154815260018201548184015284518086018652600283015481526003830154818501528186015292810180548551818502810185019096528086529394919360608601938301828280156108e457602002820191905f5260205f20905b8154815260200190600101908083116108d0575b50505091909252505050602082018190525161091357604051630cad17b760e31b815260040160405180910390fd5b806020015160600151516001600160401b03811115610934576109346117e3565b60405190808252806020026020018201604052801561095d578160200160208202803683370190505b5060408201525f5b816020015160600151518110156109c157816020015160600151818151811061099057610990612071565b6020026020010151826040015182815181106109ae576109ae612071565b6020908102919091010152600101610965565b506109cc8184610a9b565b60608201526109db8184610bc5565b604001519392505050565b5f815f0151826020015163ffffffff16604051602001610a3192919060609290921b6bffffffffffffffffffffffff1916825260a01b6001600160a01b031916601482015260200190565b60405160208183030381529060405261039a9061215e565b5f8281526002602052604090205463ffffffff16801580610a795750610a6f8183612181565b63ffffffff164211155b610a965760405163640fcd6b60e11b815260040160405180910390fd5b505050565b6040805180820182525f808252602091820181905282518084019093528083529082018190525b826080015151811015610bbe575f83608001518281518110610ae657610ae6612071565b60200260200101519050846020015160200151815f015163ffffffff1610610b21576040516301fa53c760e11b815260040160405180910390fd5b845184515f91610b319184610c33565b8051909150610b41908590610da8565b93505f5b816020015151811015610bb357866040015151811015610bab5781602001518181518110610b7557610b75612071565b602002602001015187604001518281518110610b9357610b93612071565b60200260200101818151610ba7919061219d565b9052505b600101610b45565b505050600101610ac2565b5092915050565b5f610be5610bd68460600151610e24565b60208501516040015190610da8565b90505f5f610c0184602001518486606001518760400151610eba565b91509150818015610c0f5750805b610c2c5760405163439cc0cd60e01b815260040160405180910390fd5b5050505050565b610c3b611659565b5f84815260056020908152604080832063ffffffff80881685529083528184208651909116845290915281208054909190151580610c7c5750600182015415155b905080610d25575f610c9c8787875f01518860400151896020015161103c565b905080610cbc5760405163439cc0cd60e01b815260040160405180910390fd5b6040808601515f8981526005602090815283822063ffffffff808c1684529082528483208a5190911683528152929020815180518255830151600182015582820151805192939192610d149260028501920190611610565b509050508460400151935050610d9f565b6040805160808101825283548183019081526001850154606083015281526002840180548351602082810282018101909552818152929386938186019390929091830182828015610d9357602002820191905f5260205f20905b815481526020019060010190808311610d7f575b50505050508152505092505b50509392505050565b604080518082019091525f8082526020820152610dc36116fa565b835181526020808501518183015283516040808401919091529084015160608301525f908360808460066107d05a03fa90508080610dfd57fe5b5080610e1c5760405163d4b68fd760e01b815260040160405180910390fd5b505092915050565b604080518082019091525f80825260208201528151158015610e4857506020820151155b15610e65575050604080518082019091525f808252602082015290565b6040518060400160405280835f015181526020015f5160206121d75f395f51905f528460200151610e9691906121b0565b610ead905f5160206121d75f395f51905f5261219d565b905292915050565b919050565b5f5f5f7f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000187875f01518860200151885f01515f60028110610efd57610efd612071565b60200201518951600160200201518a602001515f60028110610f2157610f21612071565b60200201518b60200151600160028110610f3d57610f3d612071565b602090810291909101518c518d830151604051610f9a9a99989796959401988952602089019790975260408801959095526060870193909352608086019190915260a085015260c084015260e08301526101008201526101200190565b604051602081830303815290604052805190602001205f1c610fbc91906121b0565b905061102e610fd5610fce88846110a7565b8690610da8565b610fdd61110f565b6110246110158561100f6040805180820182525f80825260209182015281518083019092526001825260029082015290565b906110a7565b61101e8c6111cf565b90610da8565b8862061a80611259565b909890975095505050505050565b5f5f8360405160200161104f9190611fe5565b60408051601f1981840301815291815281516020928301205f8a81526004845282812063ffffffff808c168352945291909120549092509061109b908590839085908a81169061146d16565b98975050505050505050565b604080518082019091525f80825260208201526110c2611718565b835181526020808501519082015260408082018490525f908360608460076107d05a03fa905080806110f057fe5b5080610e1c57604051632319df1960e11b815260040160405180910390fd5b611117611736565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208381019190915281019190915290565b604080518082019091525f80825260208201525f80806111fc5f5160206121d75f395f51905f52866121b0565b90505b61120881611484565b90935091505f5160206121d75f395f51905f528283098303611240576040805180820190915290815260208101919091529392505050565b5f5160206121d75f395f51905f526001820890506111ff565b6040805180820182528681526020808201869052825180840190935286835282018490525f9182919061128a611756565b5f5b6002811015611441575f6112a1826006612099565b90508482600281106112b5576112b5612071565b602002015151836112c6835f6121c3565b600c81106112d6576112d6612071565b60200201528482600281106112ed576112ed612071565b6020020151602001518382600161130491906121c3565b600c811061131457611314612071565b602002015283826002811061132b5761132b612071565b602002015151518361133e8360026121c3565b600c811061134e5761134e612071565b602002015283826002811061136557611365612071565b602002015151600160200201518361137e8360036121c3565b600c811061138e5761138e612071565b60200201528382600281106113a5576113a5612071565b6020020151602001515f600281106113bf576113bf612071565b6020020151836113d08360046121c3565b600c81106113e0576113e0612071565b60200201528382600281106113f7576113f7612071565b60200201516020015160016002811061141257611412612071565b6020020151836114238360056121c3565b600c811061143357611433612071565b60200201525060010161128c565b5061144a611775565b5f6020826101808560088cfa9151919c9115159b50909950505050505050505050565b5f8361147a868585611500565b1495945050505050565b5f80805f5160206121d75f395f51905f5260035f5160206121d75f395f51905f52865f5160206121d75f395f51905f52888909090890505f6114f4827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f525f5160206121d75f395f51905f52611597565b91959194509092505050565b5f6020845161150f91906121b0565b1561152d576040516313717da960e21b815260040160405180910390fd5b8260205b8551811161158e576115446002856121b0565b5f0361156557815f528086015160205260405f20915060028404935061157c565b808601515f528160205260405f2091506002840493505b6115876020826121c3565b9050611531565b50949350505050565b5f5f6115a1611775565b6115a9611793565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa925082806115e657fe5b50826116055760405163d51edae360e01b815260040160405180910390fd5b505195945050505050565b828054828255905f5260205f20908101928215611649579160200282015b8281111561164957825182559160200191906001019061162e565b506116559291506117b1565b5090565b604080516080810182525f91810182815260608201929092529081905b8152602001606081525090565b60405180608001604052805f81526020015f815260200161167660405180604001604052805f81526020015f81525090565b60405180608001604052805f81526020016116ce611683565b8152602001606081526020016116f560405180604001604052805f81526020015f81525090565b905290565b60405180608001604052806004906020820280368337509192915050565b60405180606001604052806003906020820280368337509192915050565b60405180604001604052806117496117c5565b81526020016116f56117c5565b604051806101800160405280600c906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b5b80821115611655575f81556001016117b2565b60405180604001604052806002906020820280368337509192915050565b634e487b7160e01b5f52604160045260245ffd5b604080519081016001600160401b0381118282101715611819576118196117e3565b60405290565b604051606081016001600160401b0381118282101715611819576118196117e3565b60405160a081016001600160401b0381118282101715611819576118196117e3565b604051608081016001600160401b0381118282101715611819576118196117e3565b604051601f8201601f191681016001600160401b03811182821017156118ad576118ad6117e3565b604052919050565b80356001600160a01b0381168114610eb5575f5ffd5b803563ffffffff81168114610eb5575f5ffd5b5f604082840312156118ee575f5ffd5b6118f66117f7565b9050611901826118b5565b815261190f602083016118cb565b602082015292915050565b5f6040828403121561192a575f5ffd5b6119326117f7565b823581526020928301359281019290925250919050565b5f82601f830112611958575f5ffd5b6119606117f7565b806040840185811115611971575f5ffd5b845b8181101561198b578035845260209384019301611973565b509095945050505050565b5f6001600160401b038211156119ae576119ae6117e3565b5060051b60200190565b5f82601f8301126119c7575f5ffd5b81356119da6119d582611996565b611885565b8082825260208201915060208360051b8601019250858311156119fb575f5ffd5b602085015b83811015611a18578035835260209283019201611a00565b5095945050505050565b5f60608284031215611a32575f5ffd5b611a3a6117f7565b9050611a46838361191a565b815260408201356001600160401b03811115611a60575f5ffd5b611a6c848285016119b8565b60208301525092915050565b5f82601f830112611a87575f5ffd5b8135611a956119d582611996565b8082825260208201915060208360051b860101925085831115611ab6575f5ffd5b602085015b83811015611a185780356001600160401b03811115611ad8575f5ffd5b86016060818903601f19011215611aed575f5ffd5b611af561181f565b611b01602083016118cb565b815260408201356001600160401b03811115611b1b575f5ffd5b82016020810190603f018a13611b2f575f5ffd5b80356001600160401b03811115611b4857611b486117e3565b611b5b601f8201601f1916602001611885565b8181528b6020838501011115611b6f575f5ffd5b816020840160208301375f6020838301015280602085015250505060608201356001600160401b03811115611ba2575f5ffd5b611bb18a602083860101611a22565b60408301525084525060209283019201611abb565b5f818303610120811215611bd8575f5ffd5b611be0611841565b9150611beb836118cb565b825260208381013590830152611c04846040850161191a565b60408301526080607f1982011215611c1a575f5ffd5b50611c236117f7565b611c308460808501611949565b8152611c3f8460c08501611949565b602082015260608201526101008201356001600160401b03811115611c62575f5ffd5b611c6e84828501611a78565b60808301525092915050565b5f5f5f60808486031215611c8c575f5ffd5b611c9685856118de565b925060408401356001600160401b03811115611cb0575f5ffd5b611cbc86828701611bc6565b92505060608401356001600160401b03811115611cd7575f5ffd5b8401601f81018613611ce7575f5ffd5b8035611cf56119d582611996565b8082825260208201915060208360051b850101925088831115611d16575f5ffd5b6020840193505b82841015611d4757833561ffff81168114611d36575f5ffd5b825260209384019390910190611d1d565b809450505050509250925092565b5f5f60608385031215611d66575f5ffd5b611d7084846118de565b915060408301356001600160401b03811115611d8a575f5ffd5b611d9685828601611bc6565b9150509250929050565b602080825282518282018190525f918401906040840190835b8181101561198b578351835260209384019390920191600101611db9565b5f5f60608385031215611de8575f5ffd5b611df284846118de565b9150611e00604084016118cb565b90509250929050565b5f60408284031215611e19575f5ffd5b61039783836118de565b5f60408284031215611e33575f5ffd5b50919050565b5f5f5f5f60c08587031215611e4c575f5ffd5b611e568686611e23565b9350611e64604086016118cb565b925060608501356001600160401b03811115611e7e575f5ffd5b850160a08188031215611e8f575f5ffd5b611e97611863565b8135815260208083013590820152611eb2886040840161191a565b604082015260808201356001600160401b03811115611ecf575f5ffd5b611edb898285016119b8565b6060830152509250611ef290508660808701611e23565b905092959194509250565b5f5f5f60808486031215611f0f575f5ffd5b611f1985856118de565b925060408401356001600160401b03811115611f33575f5ffd5b611f3f86828701611bc6565b92505060608401356001600160401b03811115611f5a575f5ffd5b611f66868287016119b8565b9150509250925092565b5f5f5f60808486031215611f82575f5ffd5b611f8c85856118de565b9250611f9a604085016118cb565b929592945050506060919091013590565b5f8151808452602084019350602083015f5b82811015611fdb578151865260209586019590910190600101611fbd565b5093949350505050565b60208082528251805183830152015160408201525f60208301516060808401526120126080840182611fab565b949350505050565b80518252602081015160208301525f6040820151612045604085018280518252602090810151910152565b50606082015160a0608085015261201260a0850182611fab565b602081525f610397602083018461201a565b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52601160045260245ffd5b808202811582820484141761039a5761039a612085565b634e487b7160e01b5f52601260045260245ffd5b5f826120d2576120d26120b0565b500490565b5f602082840312156120e7575f5ffd5b610397826118b5565b5f60208284031215612100575f5ffd5b610397826118cb565b6001600160a01b0361211a856118b5565b16815263ffffffff61212e602086016118cb565b16602082015263ffffffff83166040820152608060608201525f612155608083018461201a565b95945050505050565b80516020808301519190811015611e33575f1960209190910360031b1b16919050565b63ffffffff818116838216019081111561039a5761039a612085565b8181038181111561039a5761039a612085565b5f826121be576121be6120b0565b500690565b8082018082111561039a5761039a61208556fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47a2646970667358221220e03cfd125b5348f01fb79ec4eab619c0c1f759b6ac0033c756bbb47cbc84f5ab64736f6c634300081b0033",
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

// SetMaxStalenessPeriod is a paid mutator transaction binding the contract method 0x2a610b75.
//
// Solidity: function setMaxStalenessPeriod((address,uint32) operatorSet, uint32 maxStalenessPeriod) returns()
func (_BN254CertificateVerifier *BN254CertificateVerifierTransactor) SetMaxStalenessPeriod(opts *bind.TransactOpts, operatorSet OperatorSet, maxStalenessPeriod uint32) (*types.Transaction, error) {
	return _BN254CertificateVerifier.contract.Transact(opts, "setMaxStalenessPeriod", operatorSet, maxStalenessPeriod)
}

// SetMaxStalenessPeriod is a paid mutator transaction binding the contract method 0x2a610b75.
//
// Solidity: function setMaxStalenessPeriod((address,uint32) operatorSet, uint32 maxStalenessPeriod) returns()
func (_BN254CertificateVerifier *BN254CertificateVerifierSession) SetMaxStalenessPeriod(operatorSet OperatorSet, maxStalenessPeriod uint32) (*types.Transaction, error) {
	return _BN254CertificateVerifier.Contract.SetMaxStalenessPeriod(&_BN254CertificateVerifier.TransactOpts, operatorSet, maxStalenessPeriod)
}

// SetMaxStalenessPeriod is a paid mutator transaction binding the contract method 0x2a610b75.
//
// Solidity: function setMaxStalenessPeriod((address,uint32) operatorSet, uint32 maxStalenessPeriod) returns()
func (_BN254CertificateVerifier *BN254CertificateVerifierTransactorSession) SetMaxStalenessPeriod(operatorSet OperatorSet, maxStalenessPeriod uint32) (*types.Transaction, error) {
	return _BN254CertificateVerifier.Contract.SetMaxStalenessPeriod(&_BN254CertificateVerifier.TransactOpts, operatorSet, maxStalenessPeriod)
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
