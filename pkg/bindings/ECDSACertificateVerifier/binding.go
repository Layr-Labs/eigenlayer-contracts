// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ECDSACertificateVerifier

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

// ICrossChainRegistryTypesOperatorSetConfig is an auto generated low-level Go binding around an user-defined struct.
type ICrossChainRegistryTypesOperatorSetConfig struct {
	Owner              common.Address
	MaxStalenessPeriod uint32
}

// IECDSACertificateVerifierTypesECDSACertificate is an auto generated low-level Go binding around an user-defined struct.
type IECDSACertificateVerifierTypesECDSACertificate struct {
	ReferenceTimestamp uint32
	MessageHash        [32]byte
	Sig                []byte
}

// IOperatorTableCalculatorTypesECDSAOperatorInfo is an auto generated low-level Go binding around an user-defined struct.
type IOperatorTableCalculatorTypesECDSAOperatorInfo struct {
	Pubkey  common.Address
	Weights []*big.Int
}

// OperatorSet is an auto generated low-level Go binding around an user-defined struct.
type OperatorSet struct {
	Avs common.Address
	Id  uint32
}

// ECDSACertificateVerifierMetaData contains all meta data concerning the ECDSACertificateVerifier contract.
var ECDSACertificateVerifierMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_operatorTableUpdater\",\"type\":\"address\",\"internalType\":\"contractIOperatorTableUpdater\"},{\"name\":\"_version\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"calculateCertificateDigest\",\"inputs\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateCertificateDigestBytes\",\"inputs\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"domainSeparator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorCount\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorInfo\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIOperatorTableCalculatorTypes.ECDSAOperatorInfo\",\"components\":[{\"name\":\"pubkey\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"weights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorInfos\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIOperatorTableCalculatorTypes.ECDSAOperatorInfo[]\",\"components\":[{\"name\":\"pubkey\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"weights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorSetOwner\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalStakeWeights\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isReferenceTimestampSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestReferenceTimestamp\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxOperatorTableStaleness\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorTableUpdater\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIOperatorTableUpdater\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateOperatorTable\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorInfos\",\"type\":\"tuple[]\",\"internalType\":\"structIOperatorTableCalculatorTypes.ECDSAOperatorInfo[]\",\"components\":[{\"name\":\"pubkey\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"weights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]},{\"name\":\"operatorSetConfig\",\"type\":\"tuple\",\"internalType\":\"structICrossChainRegistryTypes.OperatorSetConfig\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxStalenessPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyCertificate\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"cert\",\"type\":\"tuple\",\"internalType\":\"structIECDSACertificateVerifierTypes.ECDSACertificate\",\"components\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sig\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyCertificateNominal\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"cert\",\"type\":\"tuple\",\"internalType\":\"structIECDSACertificateVerifierTypes.ECDSACertificate\",\"components\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sig\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"totalStakeNominalThresholds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyCertificateProportion\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"cert\",\"type\":\"tuple\",\"internalType\":\"structIECDSACertificateVerifierTypes.ECDSACertificate\",\"components\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sig\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"totalStakeProportionThresholds\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MaxStalenessPeriodUpdated\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"maxStalenessPeriod\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSetOwnerUpdated\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TableUpdated\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"operatorInfos\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"structIOperatorTableCalculatorTypes.ECDSAOperatorInfo[]\",\"components\":[{\"name\":\"pubkey\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"weights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ArrayLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CertificateStale\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignatureLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyTableUpdater\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorCountZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReferenceTimestampDoesNotExist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RootDisabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignatureExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignersNotOrdered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"TableUpdateStale\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"VerificationFailed\",\"inputs\":[]}]",
	Bin: "0x60c060405234801561000f575f5ffd5b5060405161238238038061238283398101604081905261002e9161016d565b6001600160a01b03821660805280806100468161005b565b60a0525061005490506100a1565b5050610297565b5f5f829050601f8151111561008e578260405163305a27a960e01b8152600401610085919061023c565b60405180910390fd5b805161009982610271565b179392505050565b5f54610100900460ff16156101085760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b6064820152608401610085565b5f5460ff90811614610157575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b634e487b7160e01b5f52604160045260245ffd5b5f5f6040838503121561017e575f5ffd5b82516001600160a01b0381168114610194575f5ffd5b60208401519092506001600160401b038111156101af575f5ffd5b8301601f810185136101bf575f5ffd5b80516001600160401b038111156101d8576101d8610159565b604051601f8201601f19908116603f011681016001600160401b038111828210171561020657610206610159565b60405281815282820160200187101561021d575f5ffd5b8160208401602083015e5f602083830101528093505050509250929050565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b80516020808301519190811015610291575f198160200360031b1b821691505b50919050565b60805160a0516120b56102cd5f395f818161034201526110af01525f81816101ad015281816103760152610f4d01526120b55ff3fe608060405234801561000f575f5ffd5b5060043610610106575f3560e01c80637c85ac4c1161009e578063be86e0b21161006e578063be86e0b21461026e578063c0da24201461028f578063cd83a72b146102a2578063e49613fc146102c5578063f698da25146102e5575f5ffd5b80637c85ac4c146101fa5780637d1d1f5b1461021a57806380c7d3f31461023a578063848189201461025b575f5ffd5b80635ddb9b5b116100d95780635ddb9b5b146101825780636141879e1461019557806368d6e081146101a8578063702ca531146101e7575f5ffd5b8063184674341461010a57806323c2a3cb1461013057806354fd4d501461015857806356d482f51461016d575b5f5ffd5b61011d610118366004611685565b6102ed565b6040519081526020015b60405180910390f35b61014361013e366004611765565b610308565b60405163ffffffff9091168152602001610127565b61016061033b565b60405161012791906117c5565b61018061017b366004611834565b61036b565b005b6101436101903660046118a6565b610582565b6101436101a33660046118a6565b6105a8565b6101cf7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b039091168152602001610127565b6101606101f5366004611685565b6105ce565b61020d610208366004611765565b610665565b6040516101279190611915565b61022d610228366004611978565b6107c3565b60405161012791906119cd565b61024d6102483660046119ef565b6109bc565b604051610127929190611a73565b6101cf6102693660046118a6565b6109da565b61028161027c366004611aa0565b610a03565b604051610127929190611b86565b61028161029d366004611ba8565b610aa3565b6102b56102b0366004611765565b610ba5565b6040519015158152602001610127565b6102d86102d3366004611c1e565b610bdb565b6040516101279190611c59565b61011d610d07565b5f6102f883836105ce565b8051906020012090505b92915050565b5f5f61031384610dc7565b5f90815260046020908152604080832063ffffffff8716845290915290205491505092915050565b60606103667f0000000000000000000000000000000000000000000000000000000000000000610e2a565b905090565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146103b45760405163030c1b6b60e11b815260040160405180910390fd5b5f6103cc6103c7368890038801886118a6565b610dc7565b5f8181526003602052604090205490915063ffffffff9081169086161161040657604051632f20889f60e01b815260040160405180910390fd5b5f81815260046020908152604080832063ffffffff8916845290915281208490555b8381101561048e5784848281811061044257610442611c6b565b90506020028101906104549190611c7f565b5f83815260056020908152604080832063ffffffff8b168452825280832085845290915290206104848282611cc8565b5050600101610428565b505f818152600360209081526040909120805463ffffffff191663ffffffff88161790556104be90830183611dcb565b5f8281526001602090815260409182902080546001600160a01b0319166001600160a01b0394909416939093179092556104fd91908401908401611de6565b5f828152600260209081526040808320805463ffffffff191663ffffffff958616179055600682528083209389168352929052819020805460ff19166001179055517f4f588da9ec57976194a79b5594f8f8782923d93013df2b9ed12fe125805011ef90610572908890889088908890611dff565b60405180910390a1505050505050565b5f5f61058d83610dc7565b5f9081526003602052604090205463ffffffff169392505050565b5f5f6105b383610dc7565b5f9081526002602052604090205463ffffffff169392505050565b604080517fda346acb3ce99e7c5132bf8cafb159ad8085970ebfdba78007ef0fe163063d1460208083019190915263ffffffff85168284015260608083018590528351808403820181526080909301909352815191012061062d610d07565b60405161190160f01b602082015260228101919091526042810182905260620160405160208183030381529060405291505092915050565b60605f61067184610dc7565b5f81815260046020908152604080832063ffffffff8089168552925282205492935082166001600160401b038111156106ac576106ac6116ad565b6040519080825280602002602001820160405280156106f157816020015b604080518082019091525f8152606060208201528152602001906001900390816106ca5790505b5090505f5b8263ffffffff168110156107b9575f84815260056020908152604080832063ffffffff8a16845282528083208484528252918290208251808401845281546001600160a01b031681526001820180548551818602810186019096528086529194929385810193929083018282801561078b57602002820191905f5260205f20905b815481526020019060010190808311610777575b5050505050815250508282815181106107a6576107a6611c6b565b60209081029190910101526001016106f6565b5095945050505050565b60605f6107d86103c7368690038601866118a6565b5f8181526003602052604090205490915063ffffffff84811691161461081157604051630cad17b760e31b815260040160405180910390fd5b5f81815260046020908152604080832063ffffffff871684529091529020548061084e57604051631029081560e21b815260040160405180910390fd5b5f82815260056020908152604080832063ffffffff88168452825280832083805290915281206001015490816001600160401b03811115610891576108916116ad565b6040519080825280602002602001820160405280156108ba578160200160208202803683370190505b5090505f5b838110156109b1575f85815260056020908152604080832063ffffffff8b168452825280832084845282528083206001018054825181850281018501909352808352919290919083018282801561093357602002820191905f5260205f20905b81548152602001906001019080831161091f575b509394505f93505050505b81518110801561094d57508481105b156109a75781818151811061096457610964611c6b565b602002602001015184828151811061097e5761097e611c6b565b602002602001018181516109929190611f32565b9052508061099f81611f45565b91505061093e565b50506001016108bf565b509695505050505050565b6060805f5f6109cb8686610e67565b909450925050505b9250929050565b5f5f6109e583610dc7565b5f908152600160205260409020546001600160a01b03169392505050565b5f60605f5f610a128787610e67565b915091508451825114610a385760405163512509d360e11b815260040160405180910390fd5b5f5b8251811015610a9257858181518110610a5557610a55611c6b565b6020026020010151838281518110610a6f57610a6f611c6b565b60200260200101511015610a8a57505f93509150610a9b9050565b600101610a3a565b50600193509150505b935093915050565b5f60605f5f610ab28888610e67565b90925090505f610ac98961022860208b018b611de6565b83519091508614610aed5760405163512509d360e11b815260040160405180910390fd5b5f5b8351811015610b91575f612710898984818110610b0e57610b0e611c6b565b9050602002016020810190610b239190611f5d565b61ffff16848481518110610b3957610b39611c6b565b6020026020010151610b4b9190611cb1565b610b559190611f92565b905080858381518110610b6a57610b6a611c6b565b60200260200101511015610b88575f84965096505050505050610b9c565b50600101610aef565b506001945090925050505b94509492505050565b5f5f610bb084610dc7565b5f90815260066020908152604080832063ffffffff8716845290915290205460ff1691505092915050565b604080518082019091525f8152606060208201525f610bf985610dc7565b5f81815260046020908152604080832063ffffffff891684529091529020549091508310610c6d5760405162461bcd60e51b815260206004820152601c60248201527f4f70657261746f7220696e646578206f7574206f6620626f756e647300000000604482015260640160405180910390fd5b5f81815260056020908152604080832063ffffffff8816845282528083208684528252918290208251808401845281546001600160a01b0316815260018201805485518186028101860190965280865291949293858101939290830182828015610cf457602002820191905f5260205f20905b815481526020019060010190808311610ce0575b5050505050815250509150509392505050565b60408051808201909152600a81526922b4b3b2b72630bcb2b960b11b6020909101525f7f91ab3d17e3a50a9d89e63fd30b92be7f5336b03b287bb946787a83a9d62a27667f71b625cfad44bac63b13dba07f2e1d6084ee04b6f8752101ece6126d584ee6ea610d746110a7565b8051602091820120604051610dac949392309101938452602084019290925260408301526001600160a01b0316606082015260800190565b60405160208183030381529060405280519060200120905090565b5f815f0151826020015163ffffffff16604051602001610e1292919060609290921b6bffffffffffffffffffffffff1916825260a01b6001600160a01b031916601482015260200190565b60405160208183030381529060405261030290611fa5565b60605f610e368361111c565b6040805160208082528183019092529192505f91906020820181803683375050509182525060208101929092525090565b6060805f610e7d6103c7368790038701876118a6565b5f8181526002602052604090205490915063ffffffff161580610ed357505f8181526002602090815260409091205463ffffffff1690610ebf90860186611de6565b610ec99190611fc8565b63ffffffff164211155b610ef05760405163640fcd6b60e11b815260040160405180910390fd5b5f81815260066020908152604082209190610f0d90870187611de6565b63ffffffff16815260208101919091526040015f205460ff16610f4357604051630cad17b760e31b815260040160405180910390fd5b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000166364e1df84610f7f6020870187611de6565b6040516001600160e01b031960e084901b16815263ffffffff919091166004820152602401602060405180830381865afa158015610fbf573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610fe39190611fe4565b61100057604051631b14174b60e01b815260040160405180910390fd5b5f61101b6110116020870187611de6565b86602001356102ed565b90505f6110688261102f6040890189612003565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525061114392505050565b90505f61107c8861022860208a018a611de6565b5190505f6110988561109160208b018b611de6565b858561134a565b99929850919650505050505050565b60605f6110d37f0000000000000000000000000000000000000000000000000000000000000000610e2a565b9050805f815181106110e7576110e7611c6b565b016020908101516040516001600160f81b03199091169181019190915260210160405160208183030381529060405291505090565b5f60ff8216601f81111561030257604051632cd44ac360e21b815260040160405180910390fd5b60605f825111801561116057506041825161115e9190612045565b155b61117d57604051634be6321b60e01b815260040160405180910390fd5b5f6041835161118c9190611f92565b9050806001600160401b038111156111a6576111a66116ad565b6040519080825280602002602001820160405280156111cf578160200160208202803683370190505b5091505f5b8181101561134257604080516041808252608082019092525f916020820181803683370190505090505f5b604181101561126a578581611215856041611cb1565b61121f9190611f32565b8151811061122f5761122f611c6b565b602001015160f81c60f81b82828151811061124c5761124c611c6b565b60200101906001600160f81b03191690815f1a9053506001016111ff565b505f5f6112778884611574565b90925090505f81600481111561128f5761128f612058565b146112ad57604051638baa579f60e01b815260040160405180910390fd5b8315806112ee5750856112c160018661206c565b815181106112d1576112d1611c6b565b60200260200101516001600160a01b0316826001600160a01b0316115b61130b57604051630b550c5760e41b815260040160405180910390fd5b8186858151811061131e5761131e611c6b565b6001600160a01b0392909216602092830291909101909101525050506001016111d4565b505092915050565b5f84815260046020908152604080832063ffffffff87168452909152902054606090826001600160401b03811115611384576113846116ad565b6040519080825280602002602001820160405280156113ad578160200160208202803683370190505b5091505f5b845181101561156a575f8582815181106113ce576113ce611c6b565b602002602001015190505f5f905061140160405180604001604052805f6001600160a01b03168152602001606081525090565b5f5b858110156114ca575f8b815260056020908152604080832063ffffffff8e16845282528083208484528252918290208251808401845281546001600160a01b031681526001820180548551818602810186019096528086529194929385810193929083018282801561149257602002820191905f5260205f20905b81548152602001906001019080831161147e575b5050505050815250509150836001600160a01b0316825f01516001600160a01b0316036114c257600192506114ca565b600101611403565b50816114e95760405163439cc0cd60e01b815260040160405180910390fd5b60208101515f5b8151811080156114ff57508881105b156115595781818151811061151657611516611c6b565b602002602001015188828151811061153057611530611c6b565b602002602001018181516115449190611f32565b9052508061155181611f45565b9150506114f0565b5050600190930192506113b2915050565b5050949350505050565b5f5f82516041036115a8576020830151604084015160608501515f1a61159c878285856115b3565b945094505050506109d3565b505f905060026109d3565b5f807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08311156115e857505f90506003610b9c565b604080515f8082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015611639573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b038116611661575f60019250925050610b9c565b965f9650945050505050565b803563ffffffff81168114611680575f5ffd5b919050565b5f5f60408385031215611696575f5ffd5b61169f8361166d565b946020939093013593505050565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f191681016001600160401b03811182821017156116e9576116e96116ad565b604052919050565b6001600160a01b0381168114611705575f5ffd5b50565b5f60408284031215611718575f5ffd5b604080519081016001600160401b038111828210171561173a5761173a6116ad565b604052905080823561174b816116f1565b81526117596020840161166d565b60208201525092915050565b5f5f60608385031215611776575f5ffd5b6117808484611708565b915061178e6040840161166d565b90509250929050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f6117d76020830184611797565b9392505050565b5f604082840312156117ee575f5ffd5b50919050565b5f5f83601f840112611804575f5ffd5b5081356001600160401b0381111561181a575f5ffd5b6020830191508360208260051b85010111156109d3575f5ffd5b5f5f5f5f5f60c08688031215611848575f5ffd5b61185287876117de565b94506118606040870161166d565b935060608601356001600160401b0381111561187a575f5ffd5b611886888289016117f4565b909450925061189a905087608088016117de565b90509295509295909350565b5f604082840312156118b6575f5ffd5b6117d78383611708565b80516001600160a01b03168252602080820151604082850181905281519085018190525f929190910190829060608601905b808310156107b957835182526020820191506020840193506001830192506118f2565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b8281101561196c57603f198786030184526119578583516118c0565b9450602093840193919091019060010161193b565b50929695505050505050565b5f5f60608385031215611989575f5ffd5b61178084846117de565b5f8151808452602084019350602083015f5b828110156119c35781518652602095860195909101906001016119a5565b5093949350505050565b602081525f6117d76020830184611993565b5f606082840312156117ee575f5ffd5b5f5f60608385031215611a00575f5ffd5b611a0a84846117de565b915060408301356001600160401b03811115611a24575f5ffd5b611a30858286016119df565b9150509250929050565b5f8151808452602084019350602083015f5b828110156119c35781516001600160a01b0316865260209586019590910190600101611a4c565b604081525f611a856040830185611993565b8281036020840152611a978185611a3a565b95945050505050565b5f5f5f60808486031215611ab2575f5ffd5b611abc85856117de565b925060408401356001600160401b03811115611ad6575f5ffd5b611ae2868287016119df565b92505060608401356001600160401b03811115611afd575f5ffd5b8401601f81018613611b0d575f5ffd5b80356001600160401b03811115611b2657611b266116ad565b8060051b611b36602082016116c1565b91825260208184018101929081019089841115611b51575f5ffd5b6020850194505b83851015611b7757843580835260209586019590935090910190611b58565b80955050505050509250925092565b8215158152604060208201525f611ba06040830184611a3a565b949350505050565b5f5f5f5f60808587031215611bbb575f5ffd5b611bc586866117de565b935060408501356001600160401b03811115611bdf575f5ffd5b611beb878288016119df565b93505060608501356001600160401b03811115611c06575f5ffd5b611c12878288016117f4565b95989497509550505050565b5f5f5f60808486031215611c30575f5ffd5b611c3a8585611708565b9250611c486040850161166d565b929592945050506060919091013590565b602081525f6117d760208301846118c0565b634e487b7160e01b5f52603260045260245ffd5b5f8235603e19833603018112611c93575f5ffd5b9190910192915050565b634e487b7160e01b5f52601160045260245ffd5b808202811582820484141761030257610302611c9d565b8135611cd3816116f1565b81546001600160a01b0319166001600160a01b0391909116178155602082013536839003601e19018112611d05575f5ffd5b820180356001600160401b03811115611d1c575f5ffd5b6020820191508060051b3603821315611d33575f5ffd5b600183016001600160401b03821115611d4e57611d4e6116ad565b68010000000000000000821115611d6757611d676116ad565b805482825580831015611d9c575f828152602090208381019082015b80821015611d99575f8255600182019150611d83565b50505b505f90815260208120905b82811015611dc357833582820155602090930192600101611da7565b505050505050565b5f60208284031215611ddb575f5ffd5b81356117d7816116f1565b5f60208284031215611df6575f5ffd5b6117d78261166d565b5f608082018635611e0f816116f1565b6001600160a01b0316835263ffffffff611e2b6020890161166d565b16602084015263ffffffff861660408401526080606084015283905260a0600584901b83018101908301855f603e1936839003015b87821015611f2357868503609f190184528235818112611e7e575f5ffd5b89018035611e8b816116f1565b6001600160a01b03168652602081013536829003601e19018112611ead575f5ffd5b016020810190356001600160401b03811115611ec7575f5ffd5b8060051b803603831315611ed9575f5ffd5b60406020890181905288018290526001600160fb1b03821115611efa575f5ffd5b808360608a01376060818901019750505050602083019250602084019350600182019150611e60565b50929998505050505050505050565b8082018082111561030257610302611c9d565b5f60018201611f5657611f56611c9d565b5060010190565b5f60208284031215611f6d575f5ffd5b813561ffff811681146117d7575f5ffd5b634e487b7160e01b5f52601260045260245ffd5b5f82611fa057611fa0611f7e565b500490565b805160208083015191908110156117ee575f1960209190910360031b1b16919050565b63ffffffff818116838216019081111561030257610302611c9d565b5f60208284031215611ff4575f5ffd5b815180151581146117d7575f5ffd5b5f5f8335601e19843603018112612018575f5ffd5b8301803591506001600160401b03821115612031575f5ffd5b6020019150368190038213156109d3575f5ffd5b5f8261205357612053611f7e565b500690565b634e487b7160e01b5f52602160045260245ffd5b8181038181111561030257610302611c9d56fea2646970667358221220c9a677b0fee0279a98a52462cdca09f3177fef3dd8d4c2443f12f3a112afbce664736f6c634300081b0033",
}

// ECDSACertificateVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use ECDSACertificateVerifierMetaData.ABI instead.
var ECDSACertificateVerifierABI = ECDSACertificateVerifierMetaData.ABI

// ECDSACertificateVerifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ECDSACertificateVerifierMetaData.Bin instead.
var ECDSACertificateVerifierBin = ECDSACertificateVerifierMetaData.Bin

// DeployECDSACertificateVerifier deploys a new Ethereum contract, binding an instance of ECDSACertificateVerifier to it.
func DeployECDSACertificateVerifier(auth *bind.TransactOpts, backend bind.ContractBackend, _operatorTableUpdater common.Address, _version string) (common.Address, *types.Transaction, *ECDSACertificateVerifier, error) {
	parsed, err := ECDSACertificateVerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ECDSACertificateVerifierBin), backend, _operatorTableUpdater, _version)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ECDSACertificateVerifier{ECDSACertificateVerifierCaller: ECDSACertificateVerifierCaller{contract: contract}, ECDSACertificateVerifierTransactor: ECDSACertificateVerifierTransactor{contract: contract}, ECDSACertificateVerifierFilterer: ECDSACertificateVerifierFilterer{contract: contract}}, nil
}

// ECDSACertificateVerifier is an auto generated Go binding around an Ethereum contract.
type ECDSACertificateVerifier struct {
	ECDSACertificateVerifierCaller     // Read-only binding to the contract
	ECDSACertificateVerifierTransactor // Write-only binding to the contract
	ECDSACertificateVerifierFilterer   // Log filterer for contract events
}

// ECDSACertificateVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type ECDSACertificateVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSACertificateVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ECDSACertificateVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSACertificateVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ECDSACertificateVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSACertificateVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ECDSACertificateVerifierSession struct {
	Contract     *ECDSACertificateVerifier // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ECDSACertificateVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ECDSACertificateVerifierCallerSession struct {
	Contract *ECDSACertificateVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// ECDSACertificateVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ECDSACertificateVerifierTransactorSession struct {
	Contract     *ECDSACertificateVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// ECDSACertificateVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type ECDSACertificateVerifierRaw struct {
	Contract *ECDSACertificateVerifier // Generic contract binding to access the raw methods on
}

// ECDSACertificateVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ECDSACertificateVerifierCallerRaw struct {
	Contract *ECDSACertificateVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// ECDSACertificateVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ECDSACertificateVerifierTransactorRaw struct {
	Contract *ECDSACertificateVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewECDSACertificateVerifier creates a new instance of ECDSACertificateVerifier, bound to a specific deployed contract.
func NewECDSACertificateVerifier(address common.Address, backend bind.ContractBackend) (*ECDSACertificateVerifier, error) {
	contract, err := bindECDSACertificateVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ECDSACertificateVerifier{ECDSACertificateVerifierCaller: ECDSACertificateVerifierCaller{contract: contract}, ECDSACertificateVerifierTransactor: ECDSACertificateVerifierTransactor{contract: contract}, ECDSACertificateVerifierFilterer: ECDSACertificateVerifierFilterer{contract: contract}}, nil
}

// NewECDSACertificateVerifierCaller creates a new read-only instance of ECDSACertificateVerifier, bound to a specific deployed contract.
func NewECDSACertificateVerifierCaller(address common.Address, caller bind.ContractCaller) (*ECDSACertificateVerifierCaller, error) {
	contract, err := bindECDSACertificateVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ECDSACertificateVerifierCaller{contract: contract}, nil
}

// NewECDSACertificateVerifierTransactor creates a new write-only instance of ECDSACertificateVerifier, bound to a specific deployed contract.
func NewECDSACertificateVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*ECDSACertificateVerifierTransactor, error) {
	contract, err := bindECDSACertificateVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ECDSACertificateVerifierTransactor{contract: contract}, nil
}

// NewECDSACertificateVerifierFilterer creates a new log filterer instance of ECDSACertificateVerifier, bound to a specific deployed contract.
func NewECDSACertificateVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*ECDSACertificateVerifierFilterer, error) {
	contract, err := bindECDSACertificateVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ECDSACertificateVerifierFilterer{contract: contract}, nil
}

// bindECDSACertificateVerifier binds a generic wrapper to an already deployed contract.
func bindECDSACertificateVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ECDSACertificateVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECDSACertificateVerifier *ECDSACertificateVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ECDSACertificateVerifier.Contract.ECDSACertificateVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECDSACertificateVerifier *ECDSACertificateVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECDSACertificateVerifier.Contract.ECDSACertificateVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECDSACertificateVerifier *ECDSACertificateVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECDSACertificateVerifier.Contract.ECDSACertificateVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECDSACertificateVerifier *ECDSACertificateVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ECDSACertificateVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECDSACertificateVerifier *ECDSACertificateVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECDSACertificateVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECDSACertificateVerifier *ECDSACertificateVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECDSACertificateVerifier.Contract.contract.Transact(opts, method, params...)
}

// CalculateCertificateDigest is a free data retrieval call binding the contract method 0x18467434.
//
// Solidity: function calculateCertificateDigest(uint32 referenceTimestamp, bytes32 messageHash) view returns(bytes32)
func (_ECDSACertificateVerifier *ECDSACertificateVerifierCaller) CalculateCertificateDigest(opts *bind.CallOpts, referenceTimestamp uint32, messageHash [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ECDSACertificateVerifier.contract.Call(opts, &out, "calculateCertificateDigest", referenceTimestamp, messageHash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateCertificateDigest is a free data retrieval call binding the contract method 0x18467434.
//
// Solidity: function calculateCertificateDigest(uint32 referenceTimestamp, bytes32 messageHash) view returns(bytes32)
func (_ECDSACertificateVerifier *ECDSACertificateVerifierSession) CalculateCertificateDigest(referenceTimestamp uint32, messageHash [32]byte) ([32]byte, error) {
	return _ECDSACertificateVerifier.Contract.CalculateCertificateDigest(&_ECDSACertificateVerifier.CallOpts, referenceTimestamp, messageHash)
}

// CalculateCertificateDigest is a free data retrieval call binding the contract method 0x18467434.
//
// Solidity: function calculateCertificateDigest(uint32 referenceTimestamp, bytes32 messageHash) view returns(bytes32)
func (_ECDSACertificateVerifier *ECDSACertificateVerifierCallerSession) CalculateCertificateDigest(referenceTimestamp uint32, messageHash [32]byte) ([32]byte, error) {
	return _ECDSACertificateVerifier.Contract.CalculateCertificateDigest(&_ECDSACertificateVerifier.CallOpts, referenceTimestamp, messageHash)
}

// CalculateCertificateDigestBytes is a free data retrieval call binding the contract method 0x702ca531.
//
// Solidity: function calculateCertificateDigestBytes(uint32 referenceTimestamp, bytes32 messageHash) view returns(bytes)
func (_ECDSACertificateVerifier *ECDSACertificateVerifierCaller) CalculateCertificateDigestBytes(opts *bind.CallOpts, referenceTimestamp uint32, messageHash [32]byte) ([]byte, error) {
	var out []interface{}
	err := _ECDSACertificateVerifier.contract.Call(opts, &out, "calculateCertificateDigestBytes", referenceTimestamp, messageHash)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// CalculateCertificateDigestBytes is a free data retrieval call binding the contract method 0x702ca531.
//
// Solidity: function calculateCertificateDigestBytes(uint32 referenceTimestamp, bytes32 messageHash) view returns(bytes)
func (_ECDSACertificateVerifier *ECDSACertificateVerifierSession) CalculateCertificateDigestBytes(referenceTimestamp uint32, messageHash [32]byte) ([]byte, error) {
	return _ECDSACertificateVerifier.Contract.CalculateCertificateDigestBytes(&_ECDSACertificateVerifier.CallOpts, referenceTimestamp, messageHash)
}

// CalculateCertificateDigestBytes is a free data retrieval call binding the contract method 0x702ca531.
//
// Solidity: function calculateCertificateDigestBytes(uint32 referenceTimestamp, bytes32 messageHash) view returns(bytes)
func (_ECDSACertificateVerifier *ECDSACertificateVerifierCallerSession) CalculateCertificateDigestBytes(referenceTimestamp uint32, messageHash [32]byte) ([]byte, error) {
	return _ECDSACertificateVerifier.Contract.CalculateCertificateDigestBytes(&_ECDSACertificateVerifier.CallOpts, referenceTimestamp, messageHash)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_ECDSACertificateVerifier *ECDSACertificateVerifierCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ECDSACertificateVerifier.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_ECDSACertificateVerifier *ECDSACertificateVerifierSession) DomainSeparator() ([32]byte, error) {
	return _ECDSACertificateVerifier.Contract.DomainSeparator(&_ECDSACertificateVerifier.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_ECDSACertificateVerifier *ECDSACertificateVerifierCallerSession) DomainSeparator() ([32]byte, error) {
	return _ECDSACertificateVerifier.Contract.DomainSeparator(&_ECDSACertificateVerifier.CallOpts)
}

// GetOperatorCount is a free data retrieval call binding the contract method 0x23c2a3cb.
//
// Solidity: function getOperatorCount((address,uint32) operatorSet, uint32 referenceTimestamp) view returns(uint32)
func (_ECDSACertificateVerifier *ECDSACertificateVerifierCaller) GetOperatorCount(opts *bind.CallOpts, operatorSet OperatorSet, referenceTimestamp uint32) (uint32, error) {
	var out []interface{}
	err := _ECDSACertificateVerifier.contract.Call(opts, &out, "getOperatorCount", operatorSet, referenceTimestamp)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetOperatorCount is a free data retrieval call binding the contract method 0x23c2a3cb.
//
// Solidity: function getOperatorCount((address,uint32) operatorSet, uint32 referenceTimestamp) view returns(uint32)
func (_ECDSACertificateVerifier *ECDSACertificateVerifierSession) GetOperatorCount(operatorSet OperatorSet, referenceTimestamp uint32) (uint32, error) {
	return _ECDSACertificateVerifier.Contract.GetOperatorCount(&_ECDSACertificateVerifier.CallOpts, operatorSet, referenceTimestamp)
}

// GetOperatorCount is a free data retrieval call binding the contract method 0x23c2a3cb.
//
// Solidity: function getOperatorCount((address,uint32) operatorSet, uint32 referenceTimestamp) view returns(uint32)
func (_ECDSACertificateVerifier *ECDSACertificateVerifierCallerSession) GetOperatorCount(operatorSet OperatorSet, referenceTimestamp uint32) (uint32, error) {
	return _ECDSACertificateVerifier.Contract.GetOperatorCount(&_ECDSACertificateVerifier.CallOpts, operatorSet, referenceTimestamp)
}

// GetOperatorInfo is a free data retrieval call binding the contract method 0xe49613fc.
//
// Solidity: function getOperatorInfo((address,uint32) operatorSet, uint32 referenceTimestamp, uint256 operatorIndex) view returns((address,uint256[]))
func (_ECDSACertificateVerifier *ECDSACertificateVerifierCaller) GetOperatorInfo(opts *bind.CallOpts, operatorSet OperatorSet, referenceTimestamp uint32, operatorIndex *big.Int) (IOperatorTableCalculatorTypesECDSAOperatorInfo, error) {
	var out []interface{}
	err := _ECDSACertificateVerifier.contract.Call(opts, &out, "getOperatorInfo", operatorSet, referenceTimestamp, operatorIndex)

	if err != nil {
		return *new(IOperatorTableCalculatorTypesECDSAOperatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IOperatorTableCalculatorTypesECDSAOperatorInfo)).(*IOperatorTableCalculatorTypesECDSAOperatorInfo)

	return out0, err

}

// GetOperatorInfo is a free data retrieval call binding the contract method 0xe49613fc.
//
// Solidity: function getOperatorInfo((address,uint32) operatorSet, uint32 referenceTimestamp, uint256 operatorIndex) view returns((address,uint256[]))
func (_ECDSACertificateVerifier *ECDSACertificateVerifierSession) GetOperatorInfo(operatorSet OperatorSet, referenceTimestamp uint32, operatorIndex *big.Int) (IOperatorTableCalculatorTypesECDSAOperatorInfo, error) {
	return _ECDSACertificateVerifier.Contract.GetOperatorInfo(&_ECDSACertificateVerifier.CallOpts, operatorSet, referenceTimestamp, operatorIndex)
}

// GetOperatorInfo is a free data retrieval call binding the contract method 0xe49613fc.
//
// Solidity: function getOperatorInfo((address,uint32) operatorSet, uint32 referenceTimestamp, uint256 operatorIndex) view returns((address,uint256[]))
func (_ECDSACertificateVerifier *ECDSACertificateVerifierCallerSession) GetOperatorInfo(operatorSet OperatorSet, referenceTimestamp uint32, operatorIndex *big.Int) (IOperatorTableCalculatorTypesECDSAOperatorInfo, error) {
	return _ECDSACertificateVerifier.Contract.GetOperatorInfo(&_ECDSACertificateVerifier.CallOpts, operatorSet, referenceTimestamp, operatorIndex)
}

// GetOperatorInfos is a free data retrieval call binding the contract method 0x7c85ac4c.
//
// Solidity: function getOperatorInfos((address,uint32) operatorSet, uint32 referenceTimestamp) view returns((address,uint256[])[])
func (_ECDSACertificateVerifier *ECDSACertificateVerifierCaller) GetOperatorInfos(opts *bind.CallOpts, operatorSet OperatorSet, referenceTimestamp uint32) ([]IOperatorTableCalculatorTypesECDSAOperatorInfo, error) {
	var out []interface{}
	err := _ECDSACertificateVerifier.contract.Call(opts, &out, "getOperatorInfos", operatorSet, referenceTimestamp)

	if err != nil {
		return *new([]IOperatorTableCalculatorTypesECDSAOperatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IOperatorTableCalculatorTypesECDSAOperatorInfo)).(*[]IOperatorTableCalculatorTypesECDSAOperatorInfo)

	return out0, err

}

// GetOperatorInfos is a free data retrieval call binding the contract method 0x7c85ac4c.
//
// Solidity: function getOperatorInfos((address,uint32) operatorSet, uint32 referenceTimestamp) view returns((address,uint256[])[])
func (_ECDSACertificateVerifier *ECDSACertificateVerifierSession) GetOperatorInfos(operatorSet OperatorSet, referenceTimestamp uint32) ([]IOperatorTableCalculatorTypesECDSAOperatorInfo, error) {
	return _ECDSACertificateVerifier.Contract.GetOperatorInfos(&_ECDSACertificateVerifier.CallOpts, operatorSet, referenceTimestamp)
}

// GetOperatorInfos is a free data retrieval call binding the contract method 0x7c85ac4c.
//
// Solidity: function getOperatorInfos((address,uint32) operatorSet, uint32 referenceTimestamp) view returns((address,uint256[])[])
func (_ECDSACertificateVerifier *ECDSACertificateVerifierCallerSession) GetOperatorInfos(operatorSet OperatorSet, referenceTimestamp uint32) ([]IOperatorTableCalculatorTypesECDSAOperatorInfo, error) {
	return _ECDSACertificateVerifier.Contract.GetOperatorInfos(&_ECDSACertificateVerifier.CallOpts, operatorSet, referenceTimestamp)
}

// GetOperatorSetOwner is a free data retrieval call binding the contract method 0x84818920.
//
// Solidity: function getOperatorSetOwner((address,uint32) operatorSet) view returns(address)
func (_ECDSACertificateVerifier *ECDSACertificateVerifierCaller) GetOperatorSetOwner(opts *bind.CallOpts, operatorSet OperatorSet) (common.Address, error) {
	var out []interface{}
	err := _ECDSACertificateVerifier.contract.Call(opts, &out, "getOperatorSetOwner", operatorSet)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOperatorSetOwner is a free data retrieval call binding the contract method 0x84818920.
//
// Solidity: function getOperatorSetOwner((address,uint32) operatorSet) view returns(address)
func (_ECDSACertificateVerifier *ECDSACertificateVerifierSession) GetOperatorSetOwner(operatorSet OperatorSet) (common.Address, error) {
	return _ECDSACertificateVerifier.Contract.GetOperatorSetOwner(&_ECDSACertificateVerifier.CallOpts, operatorSet)
}

// GetOperatorSetOwner is a free data retrieval call binding the contract method 0x84818920.
//
// Solidity: function getOperatorSetOwner((address,uint32) operatorSet) view returns(address)
func (_ECDSACertificateVerifier *ECDSACertificateVerifierCallerSession) GetOperatorSetOwner(operatorSet OperatorSet) (common.Address, error) {
	return _ECDSACertificateVerifier.Contract.GetOperatorSetOwner(&_ECDSACertificateVerifier.CallOpts, operatorSet)
}

// GetTotalStakeWeights is a free data retrieval call binding the contract method 0x7d1d1f5b.
//
// Solidity: function getTotalStakeWeights((address,uint32) operatorSet, uint32 referenceTimestamp) view returns(uint256[])
func (_ECDSACertificateVerifier *ECDSACertificateVerifierCaller) GetTotalStakeWeights(opts *bind.CallOpts, operatorSet OperatorSet, referenceTimestamp uint32) ([]*big.Int, error) {
	var out []interface{}
	err := _ECDSACertificateVerifier.contract.Call(opts, &out, "getTotalStakeWeights", operatorSet, referenceTimestamp)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetTotalStakeWeights is a free data retrieval call binding the contract method 0x7d1d1f5b.
//
// Solidity: function getTotalStakeWeights((address,uint32) operatorSet, uint32 referenceTimestamp) view returns(uint256[])
func (_ECDSACertificateVerifier *ECDSACertificateVerifierSession) GetTotalStakeWeights(operatorSet OperatorSet, referenceTimestamp uint32) ([]*big.Int, error) {
	return _ECDSACertificateVerifier.Contract.GetTotalStakeWeights(&_ECDSACertificateVerifier.CallOpts, operatorSet, referenceTimestamp)
}

// GetTotalStakeWeights is a free data retrieval call binding the contract method 0x7d1d1f5b.
//
// Solidity: function getTotalStakeWeights((address,uint32) operatorSet, uint32 referenceTimestamp) view returns(uint256[])
func (_ECDSACertificateVerifier *ECDSACertificateVerifierCallerSession) GetTotalStakeWeights(operatorSet OperatorSet, referenceTimestamp uint32) ([]*big.Int, error) {
	return _ECDSACertificateVerifier.Contract.GetTotalStakeWeights(&_ECDSACertificateVerifier.CallOpts, operatorSet, referenceTimestamp)
}

// IsReferenceTimestampSet is a free data retrieval call binding the contract method 0xcd83a72b.
//
// Solidity: function isReferenceTimestampSet((address,uint32) operatorSet, uint32 referenceTimestamp) view returns(bool)
func (_ECDSACertificateVerifier *ECDSACertificateVerifierCaller) IsReferenceTimestampSet(opts *bind.CallOpts, operatorSet OperatorSet, referenceTimestamp uint32) (bool, error) {
	var out []interface{}
	err := _ECDSACertificateVerifier.contract.Call(opts, &out, "isReferenceTimestampSet", operatorSet, referenceTimestamp)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsReferenceTimestampSet is a free data retrieval call binding the contract method 0xcd83a72b.
//
// Solidity: function isReferenceTimestampSet((address,uint32) operatorSet, uint32 referenceTimestamp) view returns(bool)
func (_ECDSACertificateVerifier *ECDSACertificateVerifierSession) IsReferenceTimestampSet(operatorSet OperatorSet, referenceTimestamp uint32) (bool, error) {
	return _ECDSACertificateVerifier.Contract.IsReferenceTimestampSet(&_ECDSACertificateVerifier.CallOpts, operatorSet, referenceTimestamp)
}

// IsReferenceTimestampSet is a free data retrieval call binding the contract method 0xcd83a72b.
//
// Solidity: function isReferenceTimestampSet((address,uint32) operatorSet, uint32 referenceTimestamp) view returns(bool)
func (_ECDSACertificateVerifier *ECDSACertificateVerifierCallerSession) IsReferenceTimestampSet(operatorSet OperatorSet, referenceTimestamp uint32) (bool, error) {
	return _ECDSACertificateVerifier.Contract.IsReferenceTimestampSet(&_ECDSACertificateVerifier.CallOpts, operatorSet, referenceTimestamp)
}

// LatestReferenceTimestamp is a free data retrieval call binding the contract method 0x5ddb9b5b.
//
// Solidity: function latestReferenceTimestamp((address,uint32) operatorSet) view returns(uint32)
func (_ECDSACertificateVerifier *ECDSACertificateVerifierCaller) LatestReferenceTimestamp(opts *bind.CallOpts, operatorSet OperatorSet) (uint32, error) {
	var out []interface{}
	err := _ECDSACertificateVerifier.contract.Call(opts, &out, "latestReferenceTimestamp", operatorSet)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LatestReferenceTimestamp is a free data retrieval call binding the contract method 0x5ddb9b5b.
//
// Solidity: function latestReferenceTimestamp((address,uint32) operatorSet) view returns(uint32)
func (_ECDSACertificateVerifier *ECDSACertificateVerifierSession) LatestReferenceTimestamp(operatorSet OperatorSet) (uint32, error) {
	return _ECDSACertificateVerifier.Contract.LatestReferenceTimestamp(&_ECDSACertificateVerifier.CallOpts, operatorSet)
}

// LatestReferenceTimestamp is a free data retrieval call binding the contract method 0x5ddb9b5b.
//
// Solidity: function latestReferenceTimestamp((address,uint32) operatorSet) view returns(uint32)
func (_ECDSACertificateVerifier *ECDSACertificateVerifierCallerSession) LatestReferenceTimestamp(operatorSet OperatorSet) (uint32, error) {
	return _ECDSACertificateVerifier.Contract.LatestReferenceTimestamp(&_ECDSACertificateVerifier.CallOpts, operatorSet)
}

// MaxOperatorTableStaleness is a free data retrieval call binding the contract method 0x6141879e.
//
// Solidity: function maxOperatorTableStaleness((address,uint32) operatorSet) view returns(uint32)
func (_ECDSACertificateVerifier *ECDSACertificateVerifierCaller) MaxOperatorTableStaleness(opts *bind.CallOpts, operatorSet OperatorSet) (uint32, error) {
	var out []interface{}
	err := _ECDSACertificateVerifier.contract.Call(opts, &out, "maxOperatorTableStaleness", operatorSet)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MaxOperatorTableStaleness is a free data retrieval call binding the contract method 0x6141879e.
//
// Solidity: function maxOperatorTableStaleness((address,uint32) operatorSet) view returns(uint32)
func (_ECDSACertificateVerifier *ECDSACertificateVerifierSession) MaxOperatorTableStaleness(operatorSet OperatorSet) (uint32, error) {
	return _ECDSACertificateVerifier.Contract.MaxOperatorTableStaleness(&_ECDSACertificateVerifier.CallOpts, operatorSet)
}

// MaxOperatorTableStaleness is a free data retrieval call binding the contract method 0x6141879e.
//
// Solidity: function maxOperatorTableStaleness((address,uint32) operatorSet) view returns(uint32)
func (_ECDSACertificateVerifier *ECDSACertificateVerifierCallerSession) MaxOperatorTableStaleness(operatorSet OperatorSet) (uint32, error) {
	return _ECDSACertificateVerifier.Contract.MaxOperatorTableStaleness(&_ECDSACertificateVerifier.CallOpts, operatorSet)
}

// OperatorTableUpdater is a free data retrieval call binding the contract method 0x68d6e081.
//
// Solidity: function operatorTableUpdater() view returns(address)
func (_ECDSACertificateVerifier *ECDSACertificateVerifierCaller) OperatorTableUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ECDSACertificateVerifier.contract.Call(opts, &out, "operatorTableUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OperatorTableUpdater is a free data retrieval call binding the contract method 0x68d6e081.
//
// Solidity: function operatorTableUpdater() view returns(address)
func (_ECDSACertificateVerifier *ECDSACertificateVerifierSession) OperatorTableUpdater() (common.Address, error) {
	return _ECDSACertificateVerifier.Contract.OperatorTableUpdater(&_ECDSACertificateVerifier.CallOpts)
}

// OperatorTableUpdater is a free data retrieval call binding the contract method 0x68d6e081.
//
// Solidity: function operatorTableUpdater() view returns(address)
func (_ECDSACertificateVerifier *ECDSACertificateVerifierCallerSession) OperatorTableUpdater() (common.Address, error) {
	return _ECDSACertificateVerifier.Contract.OperatorTableUpdater(&_ECDSACertificateVerifier.CallOpts)
}

// VerifyCertificate is a free data retrieval call binding the contract method 0x80c7d3f3.
//
// Solidity: function verifyCertificate((address,uint32) operatorSet, (uint32,bytes32,bytes) cert) view returns(uint256[], address[])
func (_ECDSACertificateVerifier *ECDSACertificateVerifierCaller) VerifyCertificate(opts *bind.CallOpts, operatorSet OperatorSet, cert IECDSACertificateVerifierTypesECDSACertificate) ([]*big.Int, []common.Address, error) {
	var out []interface{}
	err := _ECDSACertificateVerifier.contract.Call(opts, &out, "verifyCertificate", operatorSet, cert)

	if err != nil {
		return *new([]*big.Int), *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	out1 := *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)

	return out0, out1, err

}

// VerifyCertificate is a free data retrieval call binding the contract method 0x80c7d3f3.
//
// Solidity: function verifyCertificate((address,uint32) operatorSet, (uint32,bytes32,bytes) cert) view returns(uint256[], address[])
func (_ECDSACertificateVerifier *ECDSACertificateVerifierSession) VerifyCertificate(operatorSet OperatorSet, cert IECDSACertificateVerifierTypesECDSACertificate) ([]*big.Int, []common.Address, error) {
	return _ECDSACertificateVerifier.Contract.VerifyCertificate(&_ECDSACertificateVerifier.CallOpts, operatorSet, cert)
}

// VerifyCertificate is a free data retrieval call binding the contract method 0x80c7d3f3.
//
// Solidity: function verifyCertificate((address,uint32) operatorSet, (uint32,bytes32,bytes) cert) view returns(uint256[], address[])
func (_ECDSACertificateVerifier *ECDSACertificateVerifierCallerSession) VerifyCertificate(operatorSet OperatorSet, cert IECDSACertificateVerifierTypesECDSACertificate) ([]*big.Int, []common.Address, error) {
	return _ECDSACertificateVerifier.Contract.VerifyCertificate(&_ECDSACertificateVerifier.CallOpts, operatorSet, cert)
}

// VerifyCertificateNominal is a free data retrieval call binding the contract method 0xbe86e0b2.
//
// Solidity: function verifyCertificateNominal((address,uint32) operatorSet, (uint32,bytes32,bytes) cert, uint256[] totalStakeNominalThresholds) view returns(bool, address[])
func (_ECDSACertificateVerifier *ECDSACertificateVerifierCaller) VerifyCertificateNominal(opts *bind.CallOpts, operatorSet OperatorSet, cert IECDSACertificateVerifierTypesECDSACertificate, totalStakeNominalThresholds []*big.Int) (bool, []common.Address, error) {
	var out []interface{}
	err := _ECDSACertificateVerifier.contract.Call(opts, &out, "verifyCertificateNominal", operatorSet, cert, totalStakeNominalThresholds)

	if err != nil {
		return *new(bool), *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)

	return out0, out1, err

}

// VerifyCertificateNominal is a free data retrieval call binding the contract method 0xbe86e0b2.
//
// Solidity: function verifyCertificateNominal((address,uint32) operatorSet, (uint32,bytes32,bytes) cert, uint256[] totalStakeNominalThresholds) view returns(bool, address[])
func (_ECDSACertificateVerifier *ECDSACertificateVerifierSession) VerifyCertificateNominal(operatorSet OperatorSet, cert IECDSACertificateVerifierTypesECDSACertificate, totalStakeNominalThresholds []*big.Int) (bool, []common.Address, error) {
	return _ECDSACertificateVerifier.Contract.VerifyCertificateNominal(&_ECDSACertificateVerifier.CallOpts, operatorSet, cert, totalStakeNominalThresholds)
}

// VerifyCertificateNominal is a free data retrieval call binding the contract method 0xbe86e0b2.
//
// Solidity: function verifyCertificateNominal((address,uint32) operatorSet, (uint32,bytes32,bytes) cert, uint256[] totalStakeNominalThresholds) view returns(bool, address[])
func (_ECDSACertificateVerifier *ECDSACertificateVerifierCallerSession) VerifyCertificateNominal(operatorSet OperatorSet, cert IECDSACertificateVerifierTypesECDSACertificate, totalStakeNominalThresholds []*big.Int) (bool, []common.Address, error) {
	return _ECDSACertificateVerifier.Contract.VerifyCertificateNominal(&_ECDSACertificateVerifier.CallOpts, operatorSet, cert, totalStakeNominalThresholds)
}

// VerifyCertificateProportion is a free data retrieval call binding the contract method 0xc0da2420.
//
// Solidity: function verifyCertificateProportion((address,uint32) operatorSet, (uint32,bytes32,bytes) cert, uint16[] totalStakeProportionThresholds) view returns(bool, address[])
func (_ECDSACertificateVerifier *ECDSACertificateVerifierCaller) VerifyCertificateProportion(opts *bind.CallOpts, operatorSet OperatorSet, cert IECDSACertificateVerifierTypesECDSACertificate, totalStakeProportionThresholds []uint16) (bool, []common.Address, error) {
	var out []interface{}
	err := _ECDSACertificateVerifier.contract.Call(opts, &out, "verifyCertificateProportion", operatorSet, cert, totalStakeProportionThresholds)

	if err != nil {
		return *new(bool), *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)

	return out0, out1, err

}

// VerifyCertificateProportion is a free data retrieval call binding the contract method 0xc0da2420.
//
// Solidity: function verifyCertificateProportion((address,uint32) operatorSet, (uint32,bytes32,bytes) cert, uint16[] totalStakeProportionThresholds) view returns(bool, address[])
func (_ECDSACertificateVerifier *ECDSACertificateVerifierSession) VerifyCertificateProportion(operatorSet OperatorSet, cert IECDSACertificateVerifierTypesECDSACertificate, totalStakeProportionThresholds []uint16) (bool, []common.Address, error) {
	return _ECDSACertificateVerifier.Contract.VerifyCertificateProportion(&_ECDSACertificateVerifier.CallOpts, operatorSet, cert, totalStakeProportionThresholds)
}

// VerifyCertificateProportion is a free data retrieval call binding the contract method 0xc0da2420.
//
// Solidity: function verifyCertificateProportion((address,uint32) operatorSet, (uint32,bytes32,bytes) cert, uint16[] totalStakeProportionThresholds) view returns(bool, address[])
func (_ECDSACertificateVerifier *ECDSACertificateVerifierCallerSession) VerifyCertificateProportion(operatorSet OperatorSet, cert IECDSACertificateVerifierTypesECDSACertificate, totalStakeProportionThresholds []uint16) (bool, []common.Address, error) {
	return _ECDSACertificateVerifier.Contract.VerifyCertificateProportion(&_ECDSACertificateVerifier.CallOpts, operatorSet, cert, totalStakeProportionThresholds)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ECDSACertificateVerifier *ECDSACertificateVerifierCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ECDSACertificateVerifier.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ECDSACertificateVerifier *ECDSACertificateVerifierSession) Version() (string, error) {
	return _ECDSACertificateVerifier.Contract.Version(&_ECDSACertificateVerifier.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ECDSACertificateVerifier *ECDSACertificateVerifierCallerSession) Version() (string, error) {
	return _ECDSACertificateVerifier.Contract.Version(&_ECDSACertificateVerifier.CallOpts)
}

// UpdateOperatorTable is a paid mutator transaction binding the contract method 0x56d482f5.
//
// Solidity: function updateOperatorTable((address,uint32) operatorSet, uint32 referenceTimestamp, (address,uint256[])[] operatorInfos, (address,uint32) operatorSetConfig) returns()
func (_ECDSACertificateVerifier *ECDSACertificateVerifierTransactor) UpdateOperatorTable(opts *bind.TransactOpts, operatorSet OperatorSet, referenceTimestamp uint32, operatorInfos []IOperatorTableCalculatorTypesECDSAOperatorInfo, operatorSetConfig ICrossChainRegistryTypesOperatorSetConfig) (*types.Transaction, error) {
	return _ECDSACertificateVerifier.contract.Transact(opts, "updateOperatorTable", operatorSet, referenceTimestamp, operatorInfos, operatorSetConfig)
}

// UpdateOperatorTable is a paid mutator transaction binding the contract method 0x56d482f5.
//
// Solidity: function updateOperatorTable((address,uint32) operatorSet, uint32 referenceTimestamp, (address,uint256[])[] operatorInfos, (address,uint32) operatorSetConfig) returns()
func (_ECDSACertificateVerifier *ECDSACertificateVerifierSession) UpdateOperatorTable(operatorSet OperatorSet, referenceTimestamp uint32, operatorInfos []IOperatorTableCalculatorTypesECDSAOperatorInfo, operatorSetConfig ICrossChainRegistryTypesOperatorSetConfig) (*types.Transaction, error) {
	return _ECDSACertificateVerifier.Contract.UpdateOperatorTable(&_ECDSACertificateVerifier.TransactOpts, operatorSet, referenceTimestamp, operatorInfos, operatorSetConfig)
}

// UpdateOperatorTable is a paid mutator transaction binding the contract method 0x56d482f5.
//
// Solidity: function updateOperatorTable((address,uint32) operatorSet, uint32 referenceTimestamp, (address,uint256[])[] operatorInfos, (address,uint32) operatorSetConfig) returns()
func (_ECDSACertificateVerifier *ECDSACertificateVerifierTransactorSession) UpdateOperatorTable(operatorSet OperatorSet, referenceTimestamp uint32, operatorInfos []IOperatorTableCalculatorTypesECDSAOperatorInfo, operatorSetConfig ICrossChainRegistryTypesOperatorSetConfig) (*types.Transaction, error) {
	return _ECDSACertificateVerifier.Contract.UpdateOperatorTable(&_ECDSACertificateVerifier.TransactOpts, operatorSet, referenceTimestamp, operatorInfos, operatorSetConfig)
}

// ECDSACertificateVerifierInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ECDSACertificateVerifier contract.
type ECDSACertificateVerifierInitializedIterator struct {
	Event *ECDSACertificateVerifierInitialized // Event containing the contract specifics and raw log

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
func (it *ECDSACertificateVerifierInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ECDSACertificateVerifierInitialized)
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
		it.Event = new(ECDSACertificateVerifierInitialized)
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
func (it *ECDSACertificateVerifierInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ECDSACertificateVerifierInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ECDSACertificateVerifierInitialized represents a Initialized event raised by the ECDSACertificateVerifier contract.
type ECDSACertificateVerifierInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ECDSACertificateVerifier *ECDSACertificateVerifierFilterer) FilterInitialized(opts *bind.FilterOpts) (*ECDSACertificateVerifierInitializedIterator, error) {

	logs, sub, err := _ECDSACertificateVerifier.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ECDSACertificateVerifierInitializedIterator{contract: _ECDSACertificateVerifier.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ECDSACertificateVerifier *ECDSACertificateVerifierFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ECDSACertificateVerifierInitialized) (event.Subscription, error) {

	logs, sub, err := _ECDSACertificateVerifier.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ECDSACertificateVerifierInitialized)
				if err := _ECDSACertificateVerifier.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ECDSACertificateVerifier *ECDSACertificateVerifierFilterer) ParseInitialized(log types.Log) (*ECDSACertificateVerifierInitialized, error) {
	event := new(ECDSACertificateVerifierInitialized)
	if err := _ECDSACertificateVerifier.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ECDSACertificateVerifierMaxStalenessPeriodUpdatedIterator is returned from FilterMaxStalenessPeriodUpdated and is used to iterate over the raw logs and unpacked data for MaxStalenessPeriodUpdated events raised by the ECDSACertificateVerifier contract.
type ECDSACertificateVerifierMaxStalenessPeriodUpdatedIterator struct {
	Event *ECDSACertificateVerifierMaxStalenessPeriodUpdated // Event containing the contract specifics and raw log

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
func (it *ECDSACertificateVerifierMaxStalenessPeriodUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ECDSACertificateVerifierMaxStalenessPeriodUpdated)
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
		it.Event = new(ECDSACertificateVerifierMaxStalenessPeriodUpdated)
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
func (it *ECDSACertificateVerifierMaxStalenessPeriodUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ECDSACertificateVerifierMaxStalenessPeriodUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ECDSACertificateVerifierMaxStalenessPeriodUpdated represents a MaxStalenessPeriodUpdated event raised by the ECDSACertificateVerifier contract.
type ECDSACertificateVerifierMaxStalenessPeriodUpdated struct {
	OperatorSet        OperatorSet
	MaxStalenessPeriod uint32
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMaxStalenessPeriodUpdated is a free log retrieval operation binding the contract event 0x28539469fbbc8a5482e60966bf9376f7b9d25b2f0a65a9976f6baa3f0e3788da.
//
// Solidity: event MaxStalenessPeriodUpdated((address,uint32) operatorSet, uint32 maxStalenessPeriod)
func (_ECDSACertificateVerifier *ECDSACertificateVerifierFilterer) FilterMaxStalenessPeriodUpdated(opts *bind.FilterOpts) (*ECDSACertificateVerifierMaxStalenessPeriodUpdatedIterator, error) {

	logs, sub, err := _ECDSACertificateVerifier.contract.FilterLogs(opts, "MaxStalenessPeriodUpdated")
	if err != nil {
		return nil, err
	}
	return &ECDSACertificateVerifierMaxStalenessPeriodUpdatedIterator{contract: _ECDSACertificateVerifier.contract, event: "MaxStalenessPeriodUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxStalenessPeriodUpdated is a free log subscription operation binding the contract event 0x28539469fbbc8a5482e60966bf9376f7b9d25b2f0a65a9976f6baa3f0e3788da.
//
// Solidity: event MaxStalenessPeriodUpdated((address,uint32) operatorSet, uint32 maxStalenessPeriod)
func (_ECDSACertificateVerifier *ECDSACertificateVerifierFilterer) WatchMaxStalenessPeriodUpdated(opts *bind.WatchOpts, sink chan<- *ECDSACertificateVerifierMaxStalenessPeriodUpdated) (event.Subscription, error) {

	logs, sub, err := _ECDSACertificateVerifier.contract.WatchLogs(opts, "MaxStalenessPeriodUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ECDSACertificateVerifierMaxStalenessPeriodUpdated)
				if err := _ECDSACertificateVerifier.contract.UnpackLog(event, "MaxStalenessPeriodUpdated", log); err != nil {
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
func (_ECDSACertificateVerifier *ECDSACertificateVerifierFilterer) ParseMaxStalenessPeriodUpdated(log types.Log) (*ECDSACertificateVerifierMaxStalenessPeriodUpdated, error) {
	event := new(ECDSACertificateVerifierMaxStalenessPeriodUpdated)
	if err := _ECDSACertificateVerifier.contract.UnpackLog(event, "MaxStalenessPeriodUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ECDSACertificateVerifierOperatorSetOwnerUpdatedIterator is returned from FilterOperatorSetOwnerUpdated and is used to iterate over the raw logs and unpacked data for OperatorSetOwnerUpdated events raised by the ECDSACertificateVerifier contract.
type ECDSACertificateVerifierOperatorSetOwnerUpdatedIterator struct {
	Event *ECDSACertificateVerifierOperatorSetOwnerUpdated // Event containing the contract specifics and raw log

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
func (it *ECDSACertificateVerifierOperatorSetOwnerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ECDSACertificateVerifierOperatorSetOwnerUpdated)
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
		it.Event = new(ECDSACertificateVerifierOperatorSetOwnerUpdated)
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
func (it *ECDSACertificateVerifierOperatorSetOwnerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ECDSACertificateVerifierOperatorSetOwnerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ECDSACertificateVerifierOperatorSetOwnerUpdated represents a OperatorSetOwnerUpdated event raised by the ECDSACertificateVerifier contract.
type ECDSACertificateVerifierOperatorSetOwnerUpdated struct {
	OperatorSet OperatorSet
	Owner       common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorSetOwnerUpdated is a free log retrieval operation binding the contract event 0x806dc367095c0baf953d7144b7c4376261675ee0b4e0da2761e43673051c7375.
//
// Solidity: event OperatorSetOwnerUpdated((address,uint32) operatorSet, address owner)
func (_ECDSACertificateVerifier *ECDSACertificateVerifierFilterer) FilterOperatorSetOwnerUpdated(opts *bind.FilterOpts) (*ECDSACertificateVerifierOperatorSetOwnerUpdatedIterator, error) {

	logs, sub, err := _ECDSACertificateVerifier.contract.FilterLogs(opts, "OperatorSetOwnerUpdated")
	if err != nil {
		return nil, err
	}
	return &ECDSACertificateVerifierOperatorSetOwnerUpdatedIterator{contract: _ECDSACertificateVerifier.contract, event: "OperatorSetOwnerUpdated", logs: logs, sub: sub}, nil
}

// WatchOperatorSetOwnerUpdated is a free log subscription operation binding the contract event 0x806dc367095c0baf953d7144b7c4376261675ee0b4e0da2761e43673051c7375.
//
// Solidity: event OperatorSetOwnerUpdated((address,uint32) operatorSet, address owner)
func (_ECDSACertificateVerifier *ECDSACertificateVerifierFilterer) WatchOperatorSetOwnerUpdated(opts *bind.WatchOpts, sink chan<- *ECDSACertificateVerifierOperatorSetOwnerUpdated) (event.Subscription, error) {

	logs, sub, err := _ECDSACertificateVerifier.contract.WatchLogs(opts, "OperatorSetOwnerUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ECDSACertificateVerifierOperatorSetOwnerUpdated)
				if err := _ECDSACertificateVerifier.contract.UnpackLog(event, "OperatorSetOwnerUpdated", log); err != nil {
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
func (_ECDSACertificateVerifier *ECDSACertificateVerifierFilterer) ParseOperatorSetOwnerUpdated(log types.Log) (*ECDSACertificateVerifierOperatorSetOwnerUpdated, error) {
	event := new(ECDSACertificateVerifierOperatorSetOwnerUpdated)
	if err := _ECDSACertificateVerifier.contract.UnpackLog(event, "OperatorSetOwnerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ECDSACertificateVerifierTableUpdatedIterator is returned from FilterTableUpdated and is used to iterate over the raw logs and unpacked data for TableUpdated events raised by the ECDSACertificateVerifier contract.
type ECDSACertificateVerifierTableUpdatedIterator struct {
	Event *ECDSACertificateVerifierTableUpdated // Event containing the contract specifics and raw log

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
func (it *ECDSACertificateVerifierTableUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ECDSACertificateVerifierTableUpdated)
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
		it.Event = new(ECDSACertificateVerifierTableUpdated)
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
func (it *ECDSACertificateVerifierTableUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ECDSACertificateVerifierTableUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ECDSACertificateVerifierTableUpdated represents a TableUpdated event raised by the ECDSACertificateVerifier contract.
type ECDSACertificateVerifierTableUpdated struct {
	OperatorSet        OperatorSet
	ReferenceTimestamp uint32
	OperatorInfos      []IOperatorTableCalculatorTypesECDSAOperatorInfo
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterTableUpdated is a free log retrieval operation binding the contract event 0x4f588da9ec57976194a79b5594f8f8782923d93013df2b9ed12fe125805011ef.
//
// Solidity: event TableUpdated((address,uint32) operatorSet, uint32 referenceTimestamp, (address,uint256[])[] operatorInfos)
func (_ECDSACertificateVerifier *ECDSACertificateVerifierFilterer) FilterTableUpdated(opts *bind.FilterOpts) (*ECDSACertificateVerifierTableUpdatedIterator, error) {

	logs, sub, err := _ECDSACertificateVerifier.contract.FilterLogs(opts, "TableUpdated")
	if err != nil {
		return nil, err
	}
	return &ECDSACertificateVerifierTableUpdatedIterator{contract: _ECDSACertificateVerifier.contract, event: "TableUpdated", logs: logs, sub: sub}, nil
}

// WatchTableUpdated is a free log subscription operation binding the contract event 0x4f588da9ec57976194a79b5594f8f8782923d93013df2b9ed12fe125805011ef.
//
// Solidity: event TableUpdated((address,uint32) operatorSet, uint32 referenceTimestamp, (address,uint256[])[] operatorInfos)
func (_ECDSACertificateVerifier *ECDSACertificateVerifierFilterer) WatchTableUpdated(opts *bind.WatchOpts, sink chan<- *ECDSACertificateVerifierTableUpdated) (event.Subscription, error) {

	logs, sub, err := _ECDSACertificateVerifier.contract.WatchLogs(opts, "TableUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ECDSACertificateVerifierTableUpdated)
				if err := _ECDSACertificateVerifier.contract.UnpackLog(event, "TableUpdated", log); err != nil {
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

// ParseTableUpdated is a log parse operation binding the contract event 0x4f588da9ec57976194a79b5594f8f8782923d93013df2b9ed12fe125805011ef.
//
// Solidity: event TableUpdated((address,uint32) operatorSet, uint32 referenceTimestamp, (address,uint256[])[] operatorInfos)
func (_ECDSACertificateVerifier *ECDSACertificateVerifierFilterer) ParseTableUpdated(log types.Log) (*ECDSACertificateVerifierTableUpdated, error) {
	event := new(ECDSACertificateVerifierTableUpdated)
	if err := _ECDSACertificateVerifier.contract.UnpackLog(event, "TableUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
