// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package BN254TableCalculator

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

// OperatorSet is an auto generated low-level Go binding around an user-defined struct.
type OperatorSet struct {
	Avs common.Address
	Id  uint32
}

// BN254TableCalculatorMetaData contains all meta data concerning the BN254TableCalculator contract.
var BN254TableCalculatorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_keyRegistrar\",\"type\":\"address\",\"internalType\":\"contractIKeyRegistrar\"},{\"name\":\"_allocationManager\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"},{\"name\":\"_LOOKAHEAD_BLOCKS\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"LOOKAHEAD_BLOCKS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allocationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateOperatorTable\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"operatorSetInfo\",\"type\":\"tuple\",\"internalType\":\"structIBN254TableCalculatorTypes.BN254OperatorSetInfo\",\"components\":[{\"name\":\"operatorInfoTreeRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"numOperators\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"aggregatePubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"totalWeights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateOperatorTableBytes\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"operatorTableBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorInfos\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIBN254TableCalculatorTypes.BN254OperatorInfo[]\",\"components\":[{\"name\":\"pubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorWeight\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorWeights\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"weights\",\"type\":\"uint256[][]\",\"internalType\":\"uint256[][]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"keyRegistrar\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIKeyRegistrar\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"LookaheadBlocksSet\",\"inputs\":[{\"name\":\"lookaheadBlocks\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ECAddFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LookaheadBlocksTooHigh\",\"inputs\":[]}]",
	Bin: "0x60e060405234801561000f575f5ffd5b506040516116d03803806116d083398101604081905261002e91610060565b6001600160a01b03928316608052911660a05260c0526100a0565b6001600160a01b038116811461005d575f5ffd5b50565b5f5f5f60608486031215610072575f5ffd5b835161007d81610049565b602085015190935061008e81610049565b80925050604084015190509250925092565b60805160a05160c0516115cf6101015f395f8181610133015261063401525f818161017b015281816104df0152818161056f015261060701525f818160d401528181610329015281816103d5015281816108d401526109ef01526115cf5ff3fe608060405234801561000f575f5ffd5b5060043610610085575f3560e01c80635e120ffc116100585780635e120ffc1461012e57806371ca71d914610155578063ca8aa7c714610176578063cf2d90ef1461019d575f5ffd5b80631088794a14610089578063124c87e0146100af5780633ec45c7e146100cf57806341ee6d0e1461010e575b5f5ffd5b61009c610097366004610e20565b6101bd565b6040519081526020015b60405180910390f35b6100c26100bd366004610e56565b61025d565b6040516100a69190610eb1565b6100f67f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016100a6565b61012161011c366004610e56565b61026e565b6040516100a69190610f04565b61009c7f000000000000000000000000000000000000000000000000000000000000000081565b610168610163366004610e56565b61029f565b6040516100a6929190610f72565b6100f67f000000000000000000000000000000000000000000000000000000000000000081565b6101b06101ab366004610e56565b6102b4565b6040516100a69190611016565b5f5f5f6101c9856104da565b90925090505f5b825181101561025057846001600160a01b03168382815181106101f5576101f5611079565b60200260200101516001600160a01b0316036102485781818151811061021d5761021d611079565b60200260200101515f8151811061023657610236611079565b60200260200101519350505050610257565b6001016101d0565b505f925050505b92915050565b610265610d75565b610257826107ea565b6060610279826107ea565b6040516020016102899190610eb1565b6040516020818303038152906040529050919050565b6060806102ab836104da565b91509150915091565b60605f5f6102c1846104da565b915091505f825167ffffffffffffffff8111156102e0576102e061108d565b60405190808252806020026020018201604052801561031957816020015b610306610db4565b8152602001906001900390816102fe5790505b5090505f5b83518110156104d1577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663bd30a0b98786848151811061036957610369611079565b60200260200101516040518363ffffffff1660e01b815260040161038e9291906110da565b602060405180830381865afa1580156103a9573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103cd9190611100565b156104c9575f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639a43e3fb8887858151811061041557610415611079565b60200260200101516040518363ffffffff1660e01b815260040161043a9291906110da565b60c060405180830381865afa158015610455573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061047991906111c6565b509050604051806040016040528082815260200185848151811061049f5761049f611079565b60200260200101518152508383815181106104bc576104bc611079565b6020026020010181905250505b60010161031e565b50949350505050565b6060807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316636e875dba846040518263ffffffff1660e01b81526004016105299190611243565b5f60405180830381865afa158015610543573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261056a9190810190611274565b91505f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316634177a87c856040518263ffffffff1660e01b81526004016105b99190611243565b5f60405180830381865afa1580156105d3573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526105fa9190810190611313565b90505f6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016632bab2c4a8686856106597f0000000000000000000000000000000000000000000000000000000000000000436113b7565b6040518563ffffffff1660e01b815260040161067894939291906113ca565b5f60405180830381865afa158015610692573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526106b99190810190611443565b9050835167ffffffffffffffff8111156106d5576106d561108d565b60405190808252806020026020018201604052801561070857816020015b60608152602001906001900390816106f35790505b5092505f5b84518110156107e25760408051600180825281830190925290602080830190803683370190505084828151811061074657610746611079565b60209081029190910101525f5b83518110156107d95782828151811061076e5761076e611079565b6020026020010151818151811061078757610787611079565b60200260200101518583815181106107a1576107a1611079565b60200260200101515f815181106107ba576107ba611079565b602002602001018181516107ce91906113b7565b905250600101610753565b5060010161070d565b505050915091565b6107f2610d75565b5f5f6107fd846104da565b915091505f815f8151811061081457610814611079565b60200260200101515190505f8167ffffffffffffffff8111156108395761083961108d565b604051908082528060200260200182016040528015610862578160200160208202803683370190505b5090505f845167ffffffffffffffff8111156108805761088061108d565b6040519080825280602002602001820160405280156108a9578160200160208202803683370190505b5090506108c760405180604001604052805f81526020015f81525090565b5f5b8651811015610b1d577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663bd30a0b98a89848151811061091457610914611079565b60200260200101516040518363ffffffff1660e01b81526004016109399291906110da565b602060405180830381865afa158015610954573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109789190611100565b15610b15575f5b858110156109eb5786828151811061099957610999611079565b602002602001015181815181106109b2576109b2611079565b60200260200101518582815181106109cc576109cc611079565b602002602001018181516109e091906113b7565b90525060010161097f565b505f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639a43e3fb8b8a8581518110610a2f57610a2f611079565b60200260200101516040518363ffffffff1660e01b8152600401610a549291906110da565b60c060405180830381865afa158015610a6f573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a9391906111c6565b5090506040518060400160405280828152602001888481518110610ab957610ab9611079565b6020026020010151815250604051602001610ad49190611551565b60405160208183030381529060405280519060200120848381518110610afc57610afc611079565b6020908102919091010152610b118382610b53565b9250505b6001016108c9565b505f610b2883610bcf565b6040805160808101825291825297516020820152968701919091525050606084015250909392505050565b604080518082019091525f8082526020820152610b6e610dd5565b835181526020808501518183015283516040808401919091529084015160608301525f908360808460066107d05a03fa90508080610ba857fe5b5080610bc75760405163d4b68fd760e01b815260040160405180910390fd5b505092915050565b5f60015b8251811015610bee57610be7600282611563565b9050610bd3565b5f8167ffffffffffffffff811115610c0857610c0861108d565b604051908082528060200260200182016040528015610c31578160200160208202803683370190505b5090505f5b8451811015610c7e57848181518110610c5157610c51611079565b6020026020010151828281518110610c6b57610c6b611079565b6020908102919091010152600101610c36565b505b81600114610d52575f610c9460028461157a565b90505f5b81811015610d4a5782610cac826002611563565b81518110610cbc57610cbc611079565b602002602001015183826002610cd29190611563565b610cdd9060016113b7565b81518110610ced57610ced611079565b6020026020010151604051602001610d0f929190918252602082015260400190565b60405160208183030381529060405280519060200120838281518110610d3757610d37611079565b6020908102919091010152600101610c98565b509150610c80565b805f81518110610d6457610d64611079565b602002602001015192505050919050565b60405180608001604052805f81526020015f8152602001610da760405180604001604052805f81526020015f81525090565b8152602001606081525090565b604080516080810182525f9181018281526060820192909252908190610da7565b60405180608001604052806004906020820280368337509192915050565b5f60408284031215610e03575f5ffd5b50919050565b6001600160a01b0381168114610e1d575f5ffd5b50565b5f5f60608385031215610e31575f5ffd5b610e3b8484610df3565b91506040830135610e4b81610e09565b809150509250929050565b5f60408284031215610e66575f5ffd5b610e708383610df3565b9392505050565b5f8151808452602084019350602083015f5b82811015610ea7578151865260209586019590910190600101610e89565b5093949350505050565b6020815281516020820152602082015160408201525f6040830151610ee3606084018280518252602090810151910152565b50606083015160a080840152610efc60c0840182610e77565b949350505050565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b5f8151808452602084019350602083015f5b82811015610ea75781516001600160a01b0316865260209586019590910190600101610f4b565b604081525f610f846040830185610f39565b828103602084015280845180835260208301915060208160051b840101602087015f5b83811015610fd957601f19868403018552610fc3838351610e77565b6020958601959093509190910190600101610fa7565b509098975050505050505050565b610ffc82825180518252602090810151910152565b5f602082015160606040850152610efc6060850182610e77565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b8281101561106d57603f19878603018452611058858351610fe7565b9450602093840193919091019060010161103c565b50929695505050505050565b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52604160045260245ffd5b80356110ac81610e09565b6001600160a01b03168252602081013563ffffffff81168082146110ce575f5ffd5b80602085015250505050565b606081016110e882856110a1565b6001600160a01b039290921660409190910152919050565b5f60208284031215611110575f5ffd5b81518015158114610e70575f5ffd5b6040805190810167ffffffffffffffff811182821017156111425761114261108d565b60405290565b604051601f8201601f1916810167ffffffffffffffff811182821017156111715761117161108d565b604052919050565b5f82601f830112611188575f5ffd5b61119061111f565b8060408401858111156111a1575f5ffd5b845b818110156111bb5780518452602093840193016111a3565b509095945050505050565b5f5f82840360c08112156111d8575f5ffd5b60408112156111e5575f5ffd5b6111ed61111f565b845181526020808601519082015292506080603f198201121561120e575f5ffd5b5061121761111f565b6112248560408601611179565b81526112338560808601611179565b6020820152809150509250929050565b6040810161025782846110a1565b5f67ffffffffffffffff82111561126a5761126a61108d565b5060051b60200190565b5f60208284031215611284575f5ffd5b815167ffffffffffffffff81111561129a575f5ffd5b8201601f810184136112aa575f5ffd5b80516112bd6112b882611251565b611148565b8082825260208201915060208360051b8501019250868311156112de575f5ffd5b6020840193505b828410156113095783516112f881610e09565b8252602093840193909101906112e5565b9695505050505050565b5f60208284031215611323575f5ffd5b815167ffffffffffffffff811115611339575f5ffd5b8201601f81018413611349575f5ffd5b80516113576112b882611251565b8082825260208201915060208360051b850101925086831115611378575f5ffd5b6020840193505b8284101561130957835161139281610e09565b82526020938401939091019061137f565b634e487b7160e01b5f52601160045260245ffd5b80820180821115610257576102576113a3565b6113d481866110a1565b60a060408201525f6113e960a0830186610f39565b8281036060840152845180825260208087019201905f5b818110156114275783516001600160a01b0316835260209384019390920191600101611400565b5050809250505063ffffffff8316608083015295945050505050565b5f60208284031215611453575f5ffd5b815167ffffffffffffffff811115611469575f5ffd5b8201601f81018413611479575f5ffd5b80516114876112b882611251565b8082825260208201915060208360051b8501019250868311156114a8575f5ffd5b602084015b8381101561154657805167ffffffffffffffff8111156114cb575f5ffd5b8501603f810189136114db575f5ffd5b60208101516114ec6112b882611251565b808282526020820191506020808460051b8601010192508b83111561150f575f5ffd5b6040840193505b82841015611531578351825260209384019390910190611516565b865250506020938401939190910190506114ad565b509695505050505050565b602081525f610e706020830184610fe7565b8082028115828204841417610257576102576113a3565b5f8261159457634e487b7160e01b5f52601260045260245ffd5b50049056fea2646970667358221220b107f6a803276bf49886b13eaa27cb8636850339b95aae4eecdc1c6a46f9015564736f6c634300081b0033",
}

// BN254TableCalculatorABI is the input ABI used to generate the binding from.
// Deprecated: Use BN254TableCalculatorMetaData.ABI instead.
var BN254TableCalculatorABI = BN254TableCalculatorMetaData.ABI

// BN254TableCalculatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BN254TableCalculatorMetaData.Bin instead.
var BN254TableCalculatorBin = BN254TableCalculatorMetaData.Bin

// DeployBN254TableCalculator deploys a new Ethereum contract, binding an instance of BN254TableCalculator to it.
func DeployBN254TableCalculator(auth *bind.TransactOpts, backend bind.ContractBackend, _keyRegistrar common.Address, _allocationManager common.Address, _LOOKAHEAD_BLOCKS *big.Int) (common.Address, *types.Transaction, *BN254TableCalculator, error) {
	parsed, err := BN254TableCalculatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BN254TableCalculatorBin), backend, _keyRegistrar, _allocationManager, _LOOKAHEAD_BLOCKS)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BN254TableCalculator{BN254TableCalculatorCaller: BN254TableCalculatorCaller{contract: contract}, BN254TableCalculatorTransactor: BN254TableCalculatorTransactor{contract: contract}, BN254TableCalculatorFilterer: BN254TableCalculatorFilterer{contract: contract}}, nil
}

// BN254TableCalculator is an auto generated Go binding around an Ethereum contract.
type BN254TableCalculator struct {
	BN254TableCalculatorCaller     // Read-only binding to the contract
	BN254TableCalculatorTransactor // Write-only binding to the contract
	BN254TableCalculatorFilterer   // Log filterer for contract events
}

// BN254TableCalculatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type BN254TableCalculatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BN254TableCalculatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BN254TableCalculatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BN254TableCalculatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BN254TableCalculatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BN254TableCalculatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BN254TableCalculatorSession struct {
	Contract     *BN254TableCalculator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BN254TableCalculatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BN254TableCalculatorCallerSession struct {
	Contract *BN254TableCalculatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// BN254TableCalculatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BN254TableCalculatorTransactorSession struct {
	Contract     *BN254TableCalculatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// BN254TableCalculatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type BN254TableCalculatorRaw struct {
	Contract *BN254TableCalculator // Generic contract binding to access the raw methods on
}

// BN254TableCalculatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BN254TableCalculatorCallerRaw struct {
	Contract *BN254TableCalculatorCaller // Generic read-only contract binding to access the raw methods on
}

// BN254TableCalculatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BN254TableCalculatorTransactorRaw struct {
	Contract *BN254TableCalculatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBN254TableCalculator creates a new instance of BN254TableCalculator, bound to a specific deployed contract.
func NewBN254TableCalculator(address common.Address, backend bind.ContractBackend) (*BN254TableCalculator, error) {
	contract, err := bindBN254TableCalculator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BN254TableCalculator{BN254TableCalculatorCaller: BN254TableCalculatorCaller{contract: contract}, BN254TableCalculatorTransactor: BN254TableCalculatorTransactor{contract: contract}, BN254TableCalculatorFilterer: BN254TableCalculatorFilterer{contract: contract}}, nil
}

// NewBN254TableCalculatorCaller creates a new read-only instance of BN254TableCalculator, bound to a specific deployed contract.
func NewBN254TableCalculatorCaller(address common.Address, caller bind.ContractCaller) (*BN254TableCalculatorCaller, error) {
	contract, err := bindBN254TableCalculator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BN254TableCalculatorCaller{contract: contract}, nil
}

// NewBN254TableCalculatorTransactor creates a new write-only instance of BN254TableCalculator, bound to a specific deployed contract.
func NewBN254TableCalculatorTransactor(address common.Address, transactor bind.ContractTransactor) (*BN254TableCalculatorTransactor, error) {
	contract, err := bindBN254TableCalculator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BN254TableCalculatorTransactor{contract: contract}, nil
}

// NewBN254TableCalculatorFilterer creates a new log filterer instance of BN254TableCalculator, bound to a specific deployed contract.
func NewBN254TableCalculatorFilterer(address common.Address, filterer bind.ContractFilterer) (*BN254TableCalculatorFilterer, error) {
	contract, err := bindBN254TableCalculator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BN254TableCalculatorFilterer{contract: contract}, nil
}

// bindBN254TableCalculator binds a generic wrapper to an already deployed contract.
func bindBN254TableCalculator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BN254TableCalculatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BN254TableCalculator *BN254TableCalculatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BN254TableCalculator.Contract.BN254TableCalculatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BN254TableCalculator *BN254TableCalculatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BN254TableCalculator.Contract.BN254TableCalculatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BN254TableCalculator *BN254TableCalculatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BN254TableCalculator.Contract.BN254TableCalculatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BN254TableCalculator *BN254TableCalculatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BN254TableCalculator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BN254TableCalculator *BN254TableCalculatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BN254TableCalculator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BN254TableCalculator *BN254TableCalculatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BN254TableCalculator.Contract.contract.Transact(opts, method, params...)
}

// LOOKAHEADBLOCKS is a free data retrieval call binding the contract method 0x5e120ffc.
//
// Solidity: function LOOKAHEAD_BLOCKS() view returns(uint256)
func (_BN254TableCalculator *BN254TableCalculatorCaller) LOOKAHEADBLOCKS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BN254TableCalculator.contract.Call(opts, &out, "LOOKAHEAD_BLOCKS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LOOKAHEADBLOCKS is a free data retrieval call binding the contract method 0x5e120ffc.
//
// Solidity: function LOOKAHEAD_BLOCKS() view returns(uint256)
func (_BN254TableCalculator *BN254TableCalculatorSession) LOOKAHEADBLOCKS() (*big.Int, error) {
	return _BN254TableCalculator.Contract.LOOKAHEADBLOCKS(&_BN254TableCalculator.CallOpts)
}

// LOOKAHEADBLOCKS is a free data retrieval call binding the contract method 0x5e120ffc.
//
// Solidity: function LOOKAHEAD_BLOCKS() view returns(uint256)
func (_BN254TableCalculator *BN254TableCalculatorCallerSession) LOOKAHEADBLOCKS() (*big.Int, error) {
	return _BN254TableCalculator.Contract.LOOKAHEADBLOCKS(&_BN254TableCalculator.CallOpts)
}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_BN254TableCalculator *BN254TableCalculatorCaller) AllocationManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BN254TableCalculator.contract.Call(opts, &out, "allocationManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_BN254TableCalculator *BN254TableCalculatorSession) AllocationManager() (common.Address, error) {
	return _BN254TableCalculator.Contract.AllocationManager(&_BN254TableCalculator.CallOpts)
}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_BN254TableCalculator *BN254TableCalculatorCallerSession) AllocationManager() (common.Address, error) {
	return _BN254TableCalculator.Contract.AllocationManager(&_BN254TableCalculator.CallOpts)
}

// CalculateOperatorTable is a free data retrieval call binding the contract method 0x124c87e0.
//
// Solidity: function calculateOperatorTable((address,uint32) operatorSet) view returns((bytes32,uint256,(uint256,uint256),uint256[]) operatorSetInfo)
func (_BN254TableCalculator *BN254TableCalculatorCaller) CalculateOperatorTable(opts *bind.CallOpts, operatorSet OperatorSet) (IBN254TableCalculatorTypesBN254OperatorSetInfo, error) {
	var out []interface{}
	err := _BN254TableCalculator.contract.Call(opts, &out, "calculateOperatorTable", operatorSet)

	if err != nil {
		return *new(IBN254TableCalculatorTypesBN254OperatorSetInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IBN254TableCalculatorTypesBN254OperatorSetInfo)).(*IBN254TableCalculatorTypesBN254OperatorSetInfo)

	return out0, err

}

// CalculateOperatorTable is a free data retrieval call binding the contract method 0x124c87e0.
//
// Solidity: function calculateOperatorTable((address,uint32) operatorSet) view returns((bytes32,uint256,(uint256,uint256),uint256[]) operatorSetInfo)
func (_BN254TableCalculator *BN254TableCalculatorSession) CalculateOperatorTable(operatorSet OperatorSet) (IBN254TableCalculatorTypesBN254OperatorSetInfo, error) {
	return _BN254TableCalculator.Contract.CalculateOperatorTable(&_BN254TableCalculator.CallOpts, operatorSet)
}

// CalculateOperatorTable is a free data retrieval call binding the contract method 0x124c87e0.
//
// Solidity: function calculateOperatorTable((address,uint32) operatorSet) view returns((bytes32,uint256,(uint256,uint256),uint256[]) operatorSetInfo)
func (_BN254TableCalculator *BN254TableCalculatorCallerSession) CalculateOperatorTable(operatorSet OperatorSet) (IBN254TableCalculatorTypesBN254OperatorSetInfo, error) {
	return _BN254TableCalculator.Contract.CalculateOperatorTable(&_BN254TableCalculator.CallOpts, operatorSet)
}

// CalculateOperatorTableBytes is a free data retrieval call binding the contract method 0x41ee6d0e.
//
// Solidity: function calculateOperatorTableBytes((address,uint32) operatorSet) view returns(bytes operatorTableBytes)
func (_BN254TableCalculator *BN254TableCalculatorCaller) CalculateOperatorTableBytes(opts *bind.CallOpts, operatorSet OperatorSet) ([]byte, error) {
	var out []interface{}
	err := _BN254TableCalculator.contract.Call(opts, &out, "calculateOperatorTableBytes", operatorSet)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// CalculateOperatorTableBytes is a free data retrieval call binding the contract method 0x41ee6d0e.
//
// Solidity: function calculateOperatorTableBytes((address,uint32) operatorSet) view returns(bytes operatorTableBytes)
func (_BN254TableCalculator *BN254TableCalculatorSession) CalculateOperatorTableBytes(operatorSet OperatorSet) ([]byte, error) {
	return _BN254TableCalculator.Contract.CalculateOperatorTableBytes(&_BN254TableCalculator.CallOpts, operatorSet)
}

// CalculateOperatorTableBytes is a free data retrieval call binding the contract method 0x41ee6d0e.
//
// Solidity: function calculateOperatorTableBytes((address,uint32) operatorSet) view returns(bytes operatorTableBytes)
func (_BN254TableCalculator *BN254TableCalculatorCallerSession) CalculateOperatorTableBytes(operatorSet OperatorSet) ([]byte, error) {
	return _BN254TableCalculator.Contract.CalculateOperatorTableBytes(&_BN254TableCalculator.CallOpts, operatorSet)
}

// GetOperatorInfos is a free data retrieval call binding the contract method 0xcf2d90ef.
//
// Solidity: function getOperatorInfos((address,uint32) operatorSet) view returns(((uint256,uint256),uint256[])[])
func (_BN254TableCalculator *BN254TableCalculatorCaller) GetOperatorInfos(opts *bind.CallOpts, operatorSet OperatorSet) ([]IBN254TableCalculatorTypesBN254OperatorInfo, error) {
	var out []interface{}
	err := _BN254TableCalculator.contract.Call(opts, &out, "getOperatorInfos", operatorSet)

	if err != nil {
		return *new([]IBN254TableCalculatorTypesBN254OperatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IBN254TableCalculatorTypesBN254OperatorInfo)).(*[]IBN254TableCalculatorTypesBN254OperatorInfo)

	return out0, err

}

// GetOperatorInfos is a free data retrieval call binding the contract method 0xcf2d90ef.
//
// Solidity: function getOperatorInfos((address,uint32) operatorSet) view returns(((uint256,uint256),uint256[])[])
func (_BN254TableCalculator *BN254TableCalculatorSession) GetOperatorInfos(operatorSet OperatorSet) ([]IBN254TableCalculatorTypesBN254OperatorInfo, error) {
	return _BN254TableCalculator.Contract.GetOperatorInfos(&_BN254TableCalculator.CallOpts, operatorSet)
}

// GetOperatorInfos is a free data retrieval call binding the contract method 0xcf2d90ef.
//
// Solidity: function getOperatorInfos((address,uint32) operatorSet) view returns(((uint256,uint256),uint256[])[])
func (_BN254TableCalculator *BN254TableCalculatorCallerSession) GetOperatorInfos(operatorSet OperatorSet) ([]IBN254TableCalculatorTypesBN254OperatorInfo, error) {
	return _BN254TableCalculator.Contract.GetOperatorInfos(&_BN254TableCalculator.CallOpts, operatorSet)
}

// GetOperatorWeight is a free data retrieval call binding the contract method 0x1088794a.
//
// Solidity: function getOperatorWeight((address,uint32) operatorSet, address operator) view returns(uint256 weight)
func (_BN254TableCalculator *BN254TableCalculatorCaller) GetOperatorWeight(opts *bind.CallOpts, operatorSet OperatorSet, operator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BN254TableCalculator.contract.Call(opts, &out, "getOperatorWeight", operatorSet, operator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOperatorWeight is a free data retrieval call binding the contract method 0x1088794a.
//
// Solidity: function getOperatorWeight((address,uint32) operatorSet, address operator) view returns(uint256 weight)
func (_BN254TableCalculator *BN254TableCalculatorSession) GetOperatorWeight(operatorSet OperatorSet, operator common.Address) (*big.Int, error) {
	return _BN254TableCalculator.Contract.GetOperatorWeight(&_BN254TableCalculator.CallOpts, operatorSet, operator)
}

// GetOperatorWeight is a free data retrieval call binding the contract method 0x1088794a.
//
// Solidity: function getOperatorWeight((address,uint32) operatorSet, address operator) view returns(uint256 weight)
func (_BN254TableCalculator *BN254TableCalculatorCallerSession) GetOperatorWeight(operatorSet OperatorSet, operator common.Address) (*big.Int, error) {
	return _BN254TableCalculator.Contract.GetOperatorWeight(&_BN254TableCalculator.CallOpts, operatorSet, operator)
}

// GetOperatorWeights is a free data retrieval call binding the contract method 0x71ca71d9.
//
// Solidity: function getOperatorWeights((address,uint32) operatorSet) view returns(address[] operators, uint256[][] weights)
func (_BN254TableCalculator *BN254TableCalculatorCaller) GetOperatorWeights(opts *bind.CallOpts, operatorSet OperatorSet) (struct {
	Operators []common.Address
	Weights   [][]*big.Int
}, error) {
	var out []interface{}
	err := _BN254TableCalculator.contract.Call(opts, &out, "getOperatorWeights", operatorSet)

	outstruct := new(struct {
		Operators []common.Address
		Weights   [][]*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Operators = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Weights = *abi.ConvertType(out[1], new([][]*big.Int)).(*[][]*big.Int)

	return *outstruct, err

}

// GetOperatorWeights is a free data retrieval call binding the contract method 0x71ca71d9.
//
// Solidity: function getOperatorWeights((address,uint32) operatorSet) view returns(address[] operators, uint256[][] weights)
func (_BN254TableCalculator *BN254TableCalculatorSession) GetOperatorWeights(operatorSet OperatorSet) (struct {
	Operators []common.Address
	Weights   [][]*big.Int
}, error) {
	return _BN254TableCalculator.Contract.GetOperatorWeights(&_BN254TableCalculator.CallOpts, operatorSet)
}

// GetOperatorWeights is a free data retrieval call binding the contract method 0x71ca71d9.
//
// Solidity: function getOperatorWeights((address,uint32) operatorSet) view returns(address[] operators, uint256[][] weights)
func (_BN254TableCalculator *BN254TableCalculatorCallerSession) GetOperatorWeights(operatorSet OperatorSet) (struct {
	Operators []common.Address
	Weights   [][]*big.Int
}, error) {
	return _BN254TableCalculator.Contract.GetOperatorWeights(&_BN254TableCalculator.CallOpts, operatorSet)
}

// KeyRegistrar is a free data retrieval call binding the contract method 0x3ec45c7e.
//
// Solidity: function keyRegistrar() view returns(address)
func (_BN254TableCalculator *BN254TableCalculatorCaller) KeyRegistrar(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BN254TableCalculator.contract.Call(opts, &out, "keyRegistrar")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// KeyRegistrar is a free data retrieval call binding the contract method 0x3ec45c7e.
//
// Solidity: function keyRegistrar() view returns(address)
func (_BN254TableCalculator *BN254TableCalculatorSession) KeyRegistrar() (common.Address, error) {
	return _BN254TableCalculator.Contract.KeyRegistrar(&_BN254TableCalculator.CallOpts)
}

// KeyRegistrar is a free data retrieval call binding the contract method 0x3ec45c7e.
//
// Solidity: function keyRegistrar() view returns(address)
func (_BN254TableCalculator *BN254TableCalculatorCallerSession) KeyRegistrar() (common.Address, error) {
	return _BN254TableCalculator.Contract.KeyRegistrar(&_BN254TableCalculator.CallOpts)
}

// BN254TableCalculatorLookaheadBlocksSetIterator is returned from FilterLookaheadBlocksSet and is used to iterate over the raw logs and unpacked data for LookaheadBlocksSet events raised by the BN254TableCalculator contract.
type BN254TableCalculatorLookaheadBlocksSetIterator struct {
	Event *BN254TableCalculatorLookaheadBlocksSet // Event containing the contract specifics and raw log

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
func (it *BN254TableCalculatorLookaheadBlocksSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BN254TableCalculatorLookaheadBlocksSet)
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
		it.Event = new(BN254TableCalculatorLookaheadBlocksSet)
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
func (it *BN254TableCalculatorLookaheadBlocksSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BN254TableCalculatorLookaheadBlocksSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BN254TableCalculatorLookaheadBlocksSet represents a LookaheadBlocksSet event raised by the BN254TableCalculator contract.
type BN254TableCalculatorLookaheadBlocksSet struct {
	LookaheadBlocks *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLookaheadBlocksSet is a free log retrieval operation binding the contract event 0xa41e64dd47db91b61b43ccea8b57d75abfa496f23efc708c22753c4bc9d68842.
//
// Solidity: event LookaheadBlocksSet(uint256 lookaheadBlocks)
func (_BN254TableCalculator *BN254TableCalculatorFilterer) FilterLookaheadBlocksSet(opts *bind.FilterOpts) (*BN254TableCalculatorLookaheadBlocksSetIterator, error) {

	logs, sub, err := _BN254TableCalculator.contract.FilterLogs(opts, "LookaheadBlocksSet")
	if err != nil {
		return nil, err
	}
	return &BN254TableCalculatorLookaheadBlocksSetIterator{contract: _BN254TableCalculator.contract, event: "LookaheadBlocksSet", logs: logs, sub: sub}, nil
}

// WatchLookaheadBlocksSet is a free log subscription operation binding the contract event 0xa41e64dd47db91b61b43ccea8b57d75abfa496f23efc708c22753c4bc9d68842.
//
// Solidity: event LookaheadBlocksSet(uint256 lookaheadBlocks)
func (_BN254TableCalculator *BN254TableCalculatorFilterer) WatchLookaheadBlocksSet(opts *bind.WatchOpts, sink chan<- *BN254TableCalculatorLookaheadBlocksSet) (event.Subscription, error) {

	logs, sub, err := _BN254TableCalculator.contract.WatchLogs(opts, "LookaheadBlocksSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BN254TableCalculatorLookaheadBlocksSet)
				if err := _BN254TableCalculator.contract.UnpackLog(event, "LookaheadBlocksSet", log); err != nil {
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

// ParseLookaheadBlocksSet is a log parse operation binding the contract event 0xa41e64dd47db91b61b43ccea8b57d75abfa496f23efc708c22753c4bc9d68842.
//
// Solidity: event LookaheadBlocksSet(uint256 lookaheadBlocks)
func (_BN254TableCalculator *BN254TableCalculatorFilterer) ParseLookaheadBlocksSet(log types.Log) (*BN254TableCalculatorLookaheadBlocksSet, error) {
	event := new(BN254TableCalculatorLookaheadBlocksSet)
	if err := _BN254TableCalculator.contract.UnpackLog(event, "LookaheadBlocksSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
