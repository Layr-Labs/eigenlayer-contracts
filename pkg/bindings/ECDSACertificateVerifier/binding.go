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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_operatorTableUpdater\",\"type\":\"address\",\"internalType\":\"contractIOperatorTableUpdater\"},{\"name\":\"_version\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"calculateCertificateDigest\",\"inputs\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"domainSeparator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorCount\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorInfo\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIOperatorTableCalculatorTypes.ECDSAOperatorInfo\",\"components\":[{\"name\":\"pubkey\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"weights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorInfos\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIOperatorTableCalculatorTypes.ECDSAOperatorInfo[]\",\"components\":[{\"name\":\"pubkey\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"weights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorSetOwner\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalStakes\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestReferenceTimestamp\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxOperatorTableStaleness\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorTableUpdater\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIOperatorTableUpdater\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateOperatorTable\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorInfos\",\"type\":\"tuple[]\",\"internalType\":\"structIOperatorTableCalculatorTypes.ECDSAOperatorInfo[]\",\"components\":[{\"name\":\"pubkey\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"weights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]},{\"name\":\"operatorSetConfig\",\"type\":\"tuple\",\"internalType\":\"structICrossChainRegistryTypes.OperatorSetConfig\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxStalenessPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyCertificate\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"cert\",\"type\":\"tuple\",\"internalType\":\"structIECDSACertificateVerifierTypes.ECDSACertificate\",\"components\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sig\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyCertificateNominal\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"cert\",\"type\":\"tuple\",\"internalType\":\"structIECDSACertificateVerifierTypes.ECDSACertificate\",\"components\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sig\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"totalStakeNominalThresholds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyCertificateProportion\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"cert\",\"type\":\"tuple\",\"internalType\":\"structIECDSACertificateVerifierTypes.ECDSACertificate\",\"components\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sig\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"totalStakeProportionThresholds\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MaxStalenessPeriodUpdated\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"maxStalenessPeriod\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSetOwnerUpdated\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TableUpdated\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"operatorInfos\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"structIOperatorTableCalculatorTypes.ECDSAOperatorInfo[]\",\"components\":[{\"name\":\"pubkey\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"weights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ArrayLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CertificateStale\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignatureLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyTableUpdater\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReferenceTimestampDoesNotExist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RootDisabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignatureExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignersNotOrdered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"TableUpdateStale\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"VerificationFailed\",\"inputs\":[]}]",
	Bin: "0x60c060405234801561000f575f5ffd5b506040516122ce3803806122ce83398101604081905261002e9161016d565b6001600160a01b03821660805280806100468161005b565b60a0525061005490506100a1565b5050610297565b5f5f829050601f8151111561008e578260405163305a27a960e01b8152600401610085919061023c565b60405180910390fd5b805161009982610271565b179392505050565b5f54610100900460ff16156101085760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b6064820152608401610085565b5f5460ff90811614610157575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b634e487b7160e01b5f52604160045260245ffd5b5f5f6040838503121561017e575f5ffd5b82516001600160a01b0381168114610194575f5ffd5b60208401519092506001600160401b038111156101af575f5ffd5b8301601f810185136101bf575f5ffd5b80516001600160401b038111156101d8576101d8610159565b604051601f8201601f19908116603f011681016001600160401b038111828210171561020657610206610159565b60405281815282820160200187101561021d575f5ffd5b8160208401602083015e5f602083830101528093505050509250929050565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b80516020808301519190811015610291575f198160200360031b1b821691505b50919050565b60805160a0516120016102cd5f395f818161054e015261123801525f81816101bb015281816105820152610e9101526120015ff3fe608060405234801561000f575f5ffd5b50600436106100f0575f3560e01c806368d6e08111610093578063be86e0b211610063578063be86e0b214610249578063c0da24201461026a578063e49613fc1461027d578063f698da251461029d575f5ffd5b806368d6e081146101b65780637c85ac4c146101f557806380c7d3f3146102155780638481892014610236575f5ffd5b806354fd4d50116100ce57806354fd4d501461016657806356d482f51461017b5780635ddb9b5b146101905780636141879e146101a3575f5ffd5b806304cdbae4146100f4578063184674341461011d57806323c2a3cb1461013e575b5f5ffd5b6101076101023660046115fa565b6102a5565b6040516101149190611666565b60405180910390f35b61013061012b36600461167f565b6104a5565b604051908152602001610114565b61015161014c36600461175f565b610514565b60405163ffffffff9091168152602001610114565b61016e610547565b604051610114919061177a565b61018e6101893660046117ef565b610577565b005b61015161019e366004611861565b610770565b6101516101b1366004611861565b610796565b6101dd7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b039091168152602001610114565b61020861020336600461175f565b6107bc565b60405161011491906118d0565b610228610223366004611943565b61091a565b6040516101149291906119c7565b6101dd610244366004611861565b610938565b61025c6102573660046119f4565b610961565b604051610114929190611ada565b61025c610278366004611af4565b610a01565b61029061028b366004611b6a565b610b03565b6040516101149190611ba5565b610130610c2f565b60605f6102bf6102ba36869003860186611861565b610cef565b5f8181526003602052604090205490915063ffffffff8481169116146102f857604051630cad17b760e31b815260040160405180910390fd5b5f81815260046020908152604080832063ffffffff871684529091529020548061033557604051630cad17b760e31b815260040160405180910390fd5b5f82815260056020908152604080832063ffffffff88168452825280832083805290915281206001015490816001600160401b03811115610378576103786116a7565b6040519080825280602002602001820160405280156103a1578160200160208202803683370190505b5090505f5b83811015610498575f85815260056020908152604080832063ffffffff8b168452825280832084845282528083206001018054825181850281018501909352808352919290919083018282801561041a57602002820191905f5260205f20905b815481526020019060010190808311610406575b509394505f93505050505b81518110801561043457508481105b1561048e5781818151811061044b5761044b611bb7565b602002602001015184828151811061046557610465611bb7565b602002602001018181516104799190611bdf565b9052508061048681611bf2565b915050610425565b50506001016103a6565b5093505050505b92915050565b604080517fda346acb3ce99e7c5132bf8cafb159ad8085970ebfdba78007ef0fe163063d14602082015263ffffffff841691810191909152606081018290525f90819060800160405160208183030381529060405280519060200120905061050c81610d52565b949350505050565b5f5f61051f84610cef565b5f90815260046020908152604080832063ffffffff8716845290915290205491505092915050565b60606105727f0000000000000000000000000000000000000000000000000000000000000000610d98565b905090565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146105c05760405163030c1b6b60e11b815260040160405180910390fd5b5f6105d36102ba36889003880188611861565b5f8181526003602052604090205490915063ffffffff9081169086161161060d57604051632f20889f60e01b815260040160405180910390fd5b5f81815260046020908152604080832063ffffffff8916845290915281208490555b838110156106955784848281811061064957610649611bb7565b905060200281019061065b9190611c0a565b5f83815260056020908152604080832063ffffffff8b1684528252808320858452909152902061068b8282611c3f565b505060010161062f565b505f818152600360209081526040909120805463ffffffff191663ffffffff88161790556106c590830183611d42565b5f8281526001602090815260409182902080546001600160a01b0319166001600160a01b03949094169390931790925561070491908401908401611d5d565b5f8281526002602052604090819020805463ffffffff191663ffffffff9390931692909217909155517f4f588da9ec57976194a79b5594f8f8782923d93013df2b9ed12fe125805011ef90610760908890889088908890611d76565b60405180910390a1505050505050565b5f5f61077b83610cef565b5f9081526003602052604090205463ffffffff169392505050565b5f5f6107a183610cef565b5f9081526002602052604090205463ffffffff169392505050565b60605f6107c884610cef565b5f81815260046020908152604080832063ffffffff8089168552925282205492935082166001600160401b03811115610803576108036116a7565b60405190808252806020026020018201604052801561084857816020015b604080518082019091525f8152606060208201528152602001906001900390816108215790505b5090505f5b8263ffffffff16811015610910575f84815260056020908152604080832063ffffffff8a16845282528083208484528252918290208251808401845281546001600160a01b03168152600182018054855181860281018601909652808652919492938581019392908301828280156108e257602002820191905f5260205f20905b8154815260200190600101908083116108ce575b5050505050815250508282815181106108fd576108fd611bb7565b602090810291909101015260010161084d565b5095945050505050565b6060805f5f6109298686610dd5565b909450925050505b9250929050565b5f5f61094383610cef565b5f908152600160205260409020546001600160a01b03169392505050565b5f60605f5f6109708787610dd5565b9150915084518251146109965760405163512509d360e11b815260040160405180910390fd5b5f5b82518110156109f0578581815181106109b3576109b3611bb7565b60200260200101518382815181106109cd576109cd611bb7565b602002602001015110156109e857505f935091506109f99050565b600101610998565b50600193509150505b935093915050565b5f60605f5f610a108888610dd5565b90925090505f610a278961010260208b018b611d5d565b83519091508614610a4b5760405163512509d360e11b815260040160405180910390fd5b5f5b8351811015610aef575f612710898984818110610a6c57610a6c611bb7565b9050602002016020810190610a819190611ea9565b61ffff16848481518110610a9757610a97611bb7565b6020026020010151610aa99190611c28565b610ab39190611ede565b905080858381518110610ac857610ac8611bb7565b60200260200101511015610ae6575f84965096505050505050610afa565b50600101610a4d565b506001945090925050505b94509492505050565b604080518082019091525f8152606060208201525f610b2185610cef565b5f81815260046020908152604080832063ffffffff891684529091529020549091508310610b955760405162461bcd60e51b815260206004820152601c60248201527f4f70657261746f7220696e646578206f7574206f6620626f756e647300000000604482015260640160405180910390fd5b5f81815260056020908152604080832063ffffffff8816845282528083208684528252918290208251808401845281546001600160a01b0316815260018201805485518186028101860190965280865291949293858101939290830182828015610c1c57602002820191905f5260205f20905b815481526020019060010190808311610c08575b5050505050815250509150509392505050565b60408051808201909152600a81526922b4b3b2b72630bcb2b960b11b6020909101525f7f91ab3d17e3a50a9d89e63fd30b92be7f5336b03b287bb946787a83a9d62a27667f71b625cfad44bac63b13dba07f2e1d6084ee04b6f8752101ece6126d584ee6ea610c9c611230565b8051602091820120604051610cd4949392309101938452602084019290925260408301526001600160a01b0316606082015260800190565b60405160208183030381529060405280519060200120905090565b5f815f0151826020015163ffffffff16604051602001610d3a92919060609290921b6bffffffffffffffffffffffff1916825260a01b6001600160a01b031916601482015260200190565b60405160208183030381529060405261049f90611ef1565b5f610d5b610c2f565b60405161190160f01b6020820152602281019190915260428101839052606201604051602081830303815290604052805190602001209050919050565b60605f610da4836112a5565b6040805160208082528183019092529192505f91906020820181803683375050509182525060208101929092525090565b6060805f610deb6102ba36879003870187611861565b5f8181526002602090815260409091205491925063ffffffff90911690610e1490860186611d5d565b610e1e9190611f14565b63ffffffff16421115610e445760405163640fcd6b60e11b815260040160405180910390fd5b610e516020850185611d5d565b5f8281526003602052604090205463ffffffff908116911614610e8757604051630cad17b760e31b815260040160405180910390fd5b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000166364e1df84610ec36020870187611d5d565b6040516001600160e01b031960e084901b16815263ffffffff919091166004820152602401602060405180830381865afa158015610f03573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610f279190611f30565b610f4457604051631b14174b60e01b815260040160405180910390fd5b5f610f56866101026020880188611d5d565b90505f81516001600160401b03811115610f7257610f726116a7565b604051908082528060200260200182016040528015610f9b578160200160208202803683370190505b5090505f610fb9610faf6020890189611d5d565b88602001356104a5565b90505f61100682610fcd60408b018b611f4f565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152506112cc92505050565b90505f5b8151811015611220575f82828151811061102657611026611bb7565b602002602001015190505f5f905061105960405180604001604052805f6001600160a01b03168152602001606081525090565b5f5b60045f8b81526020019081526020015f205f8e5f01602081019061107f9190611d5d565b63ffffffff1663ffffffff1681526020019081526020015f205481101561117f5760055f8b81526020019081526020015f205f8e5f0160208101906110c49190611d5d565b63ffffffff16815260208082019290925260409081015f90812084825283528190208151808301835281546001600160a01b03168152600182018054845181870281018701909552808552919492938584019390929083018282801561114757602002820191905f5260205f20905b815481526020019060010190808311611133575b5050505050815250509150836001600160a01b0316825f01516001600160a01b031603611177576001925061117f565b60010161105b565b508161119e5760405163439cc0cd60e01b815260040160405180910390fd5b60208101515f5b8151811080156111b55750885181105b1561120f578181815181106111cc576111cc611bb7565b60200260200101518982815181106111e6576111e6611bb7565b602002602001018181516111fa9190611bdf565b9052508061120781611bf2565b9150506111a5565b50506001909301925061100a915050565b5091989197509095505050505050565b60605f61125c7f0000000000000000000000000000000000000000000000000000000000000000610d98565b9050805f8151811061127057611270611bb7565b016020908101516040516001600160f81b03199091169181019190915260210160405160208183030381529060405291505090565b5f60ff8216601f81111561049f57604051632cd44ac360e21b815260040160405180910390fd5b60605f82511180156112e95750604182516112e79190611f91565b155b61130657604051634be6321b60e01b815260040160405180910390fd5b5f604183516113159190611ede565b9050806001600160401b0381111561132f5761132f6116a7565b604051908082528060200260200182016040528015611358578160200160208202803683370190505b5091505f5b818110156114cb57604080516041808252608082019092525f916020820181803683370190505090505f5b60418110156113f357858161139e856041611c28565b6113a89190611bdf565b815181106113b8576113b8611bb7565b602001015160f81c60f81b8282815181106113d5576113d5611bb7565b60200101906001600160f81b03191690815f1a905350600101611388565b505f5f61140088846114d3565b90925090505f81600481111561141857611418611fa4565b1461143657604051638baa579f60e01b815260040160405180910390fd5b83158061147757508561144a600186611fb8565b8151811061145a5761145a611bb7565b60200260200101516001600160a01b0316826001600160a01b0316115b61149457604051630b550c5760e41b815260040160405180910390fd5b818685815181106114a7576114a7611bb7565b6001600160a01b03929092166020928302919091019091015250505060010161135d565b505092915050565b5f5f8251604103611507576020830151604084015160608501515f1a6114fb87828585611512565b94509450505050610931565b505f90506002610931565b5f807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a083111561154757505f90506003610afa565b604080515f8082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015611598573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b0381166115c0575f60019250925050610afa565b965f9650945050505050565b5f604082840312156115dc575f5ffd5b50919050565b803563ffffffff811681146115f5575f5ffd5b919050565b5f5f6060838503121561160b575f5ffd5b61161584846115cc565b9150611623604084016115e2565b90509250929050565b5f8151808452602084019350602083015f5b8281101561165c57815186526020958601959091019060010161163e565b5093949350505050565b602081525f611678602083018461162c565b9392505050565b5f5f60408385031215611690575f5ffd5b611699836115e2565b946020939093013593505050565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f191681016001600160401b03811182821017156116e3576116e36116a7565b604052919050565b6001600160a01b03811681146116ff575f5ffd5b50565b5f60408284031215611712575f5ffd5b604080519081016001600160401b0381118282101715611734576117346116a7565b6040529050808235611745816116eb565b8152611753602084016115e2565b60208201525092915050565b5f5f60608385031215611770575f5ffd5b6116158484611702565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b5f5f83601f8401126117bf575f5ffd5b5081356001600160401b038111156117d5575f5ffd5b6020830191508360208260051b8501011115610931575f5ffd5b5f5f5f5f5f60c08688031215611803575f5ffd5b61180d87876115cc565b945061181b604087016115e2565b935060608601356001600160401b03811115611835575f5ffd5b611841888289016117af565b9094509250611855905087608088016115cc565b90509295509295909350565b5f60408284031215611871575f5ffd5b6116788383611702565b80516001600160a01b03168252602080820151604082850181905281519085018190525f929190910190829060608601905b8083101561091057835182526020820191506020840193506001830192506118ad565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b8281101561192757603f1987860301845261191285835161187b565b945060209384019391909101906001016118f6565b50929695505050505050565b5f606082840312156115dc575f5ffd5b5f5f60608385031215611954575f5ffd5b61195e84846115cc565b915060408301356001600160401b03811115611978575f5ffd5b61198485828601611933565b9150509250929050565b5f8151808452602084019350602083015f5b8281101561165c5781516001600160a01b03168652602095860195909101906001016119a0565b604081525f6119d9604083018561162c565b82810360208401526119eb818561198e565b95945050505050565b5f5f5f60808486031215611a06575f5ffd5b611a1085856115cc565b925060408401356001600160401b03811115611a2a575f5ffd5b611a3686828701611933565b92505060608401356001600160401b03811115611a51575f5ffd5b8401601f81018613611a61575f5ffd5b80356001600160401b03811115611a7a57611a7a6116a7565b8060051b611a8a602082016116bb565b91825260208184018101929081019089841115611aa5575f5ffd5b6020850194505b83851015611acb57843580835260209586019590935090910190611aac565b80955050505050509250925092565b8215158152604060208201525f61050c604083018461198e565b5f5f5f5f60808587031215611b07575f5ffd5b611b1186866115cc565b935060408501356001600160401b03811115611b2b575f5ffd5b611b3787828801611933565b93505060608501356001600160401b03811115611b52575f5ffd5b611b5e878288016117af565b95989497509550505050565b5f5f5f60808486031215611b7c575f5ffd5b611b868585611702565b9250611b94604085016115e2565b929592945050506060919091013590565b602081525f611678602083018461187b565b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52601160045260245ffd5b8082018082111561049f5761049f611bcb565b5f60018201611c0357611c03611bcb565b5060010190565b5f8235603e19833603018112611c1e575f5ffd5b9190910192915050565b808202811582820484141761049f5761049f611bcb565b8135611c4a816116eb565b81546001600160a01b0319166001600160a01b0391909116178155602082013536839003601e19018112611c7c575f5ffd5b820180356001600160401b03811115611c93575f5ffd5b6020820191508060051b3603821315611caa575f5ffd5b600183016001600160401b03821115611cc557611cc56116a7565b68010000000000000000821115611cde57611cde6116a7565b805482825580831015611d13575f828152602090208381019082015b80821015611d10575f8255600182019150611cfa565b50505b505f90815260208120905b82811015611d3a57833582820155602090930192600101611d1e565b505050505050565b5f60208284031215611d52575f5ffd5b8135611678816116eb565b5f60208284031215611d6d575f5ffd5b611678826115e2565b5f608082018635611d86816116eb565b6001600160a01b0316835263ffffffff611da2602089016115e2565b16602084015263ffffffff861660408401526080606084015283905260a0600584901b83018101908301855f603e1936839003015b87821015611e9a57868503609f190184528235818112611df5575f5ffd5b89018035611e02816116eb565b6001600160a01b03168652602081013536829003601e19018112611e24575f5ffd5b016020810190356001600160401b03811115611e3e575f5ffd5b8060051b803603831315611e50575f5ffd5b60406020890181905288018290526001600160fb1b03821115611e71575f5ffd5b808360608a01376060818901019750505050602083019250602084019350600182019150611dd7565b50929998505050505050505050565b5f60208284031215611eb9575f5ffd5b813561ffff81168114611678575f5ffd5b634e487b7160e01b5f52601260045260245ffd5b5f82611eec57611eec611eca565b500490565b805160208083015191908110156115dc575f1960209190910360031b1b16919050565b63ffffffff818116838216019081111561049f5761049f611bcb565b5f60208284031215611f40575f5ffd5b81518015158114611678575f5ffd5b5f5f8335601e19843603018112611f64575f5ffd5b8301803591506001600160401b03821115611f7d575f5ffd5b602001915036819003821315610931575f5ffd5b5f82611f9f57611f9f611eca565b500690565b634e487b7160e01b5f52602160045260245ffd5b8181038181111561049f5761049f611bcb56fea2646970667358221220f75c79887adc88462bd952997bf884e39246bab87cfa122a8d9db93fdc348b0364736f6c634300081b0033",
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

// GetTotalStakes is a free data retrieval call binding the contract method 0x04cdbae4.
//
// Solidity: function getTotalStakes((address,uint32) operatorSet, uint32 referenceTimestamp) view returns(uint256[])
func (_ECDSACertificateVerifier *ECDSACertificateVerifierCaller) GetTotalStakes(opts *bind.CallOpts, operatorSet OperatorSet, referenceTimestamp uint32) ([]*big.Int, error) {
	var out []interface{}
	err := _ECDSACertificateVerifier.contract.Call(opts, &out, "getTotalStakes", operatorSet, referenceTimestamp)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetTotalStakes is a free data retrieval call binding the contract method 0x04cdbae4.
//
// Solidity: function getTotalStakes((address,uint32) operatorSet, uint32 referenceTimestamp) view returns(uint256[])
func (_ECDSACertificateVerifier *ECDSACertificateVerifierSession) GetTotalStakes(operatorSet OperatorSet, referenceTimestamp uint32) ([]*big.Int, error) {
	return _ECDSACertificateVerifier.Contract.GetTotalStakes(&_ECDSACertificateVerifier.CallOpts, operatorSet, referenceTimestamp)
}

// GetTotalStakes is a free data retrieval call binding the contract method 0x04cdbae4.
//
// Solidity: function getTotalStakes((address,uint32) operatorSet, uint32 referenceTimestamp) view returns(uint256[])
func (_ECDSACertificateVerifier *ECDSACertificateVerifierCallerSession) GetTotalStakes(operatorSet OperatorSet, referenceTimestamp uint32) ([]*big.Int, error) {
	return _ECDSACertificateVerifier.Contract.GetTotalStakes(&_ECDSACertificateVerifier.CallOpts, operatorSet, referenceTimestamp)
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
