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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_keyRegistrar\",\"type\":\"address\",\"internalType\":\"contractIKeyRegistrar\"},{\"name\":\"_allocationManager\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"},{\"name\":\"_LOOKAHEAD_BLOCKS\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"LOOKAHEAD_BLOCKS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allocationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateOperatorTable\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"operatorSetInfo\",\"type\":\"tuple\",\"internalType\":\"structIBN254TableCalculatorTypes.BN254OperatorSetInfo\",\"components\":[{\"name\":\"operatorInfoTreeRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"numOperators\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"aggregatePubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"totalWeights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateOperatorTableBytes\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"operatorTableBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorInfos\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIBN254TableCalculatorTypes.BN254OperatorInfo[]\",\"components\":[{\"name\":\"pubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorWeight\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorWeights\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"weights\",\"type\":\"uint256[][]\",\"internalType\":\"uint256[][]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"keyRegistrar\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIKeyRegistrar\"}],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"ECAddFailed\",\"inputs\":[]}]",
	Bin: "0x60e060405234801561000f575f5ffd5b5060405161174b38038061174b83398101604081905261002e91610060565b6001600160a01b03928316608052911660a05260c0526100a0565b6001600160a01b038116811461005d575f5ffd5b50565b5f5f5f60608486031215610072575f5ffd5b835161007d81610049565b602085015190935061008e81610049565b80925050604084015190509250925092565b60805160a05160c05161164a6101015f395f8181610133015261063301525f818161017b015281816104de0152818161056e015261060601525f818160d401528181610328015281816103d4015281816109570152610a72015261164a5ff3fe608060405234801561000f575f5ffd5b5060043610610085575f3560e01c80635e120ffc116100585780635e120ffc1461012e57806371ca71d914610155578063ca8aa7c714610176578063cf2d90ef1461019d575f5ffd5b80631088794a14610089578063124c87e0146100af5780633ec45c7e146100cf57806341ee6d0e1461010e575b5f5ffd5b61009c610097366004610ea2565b6101bd565b6040519081526020015b60405180910390f35b6100c26100bd366004610ed8565b61025d565b6040516100a69190610f33565b6100f67f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016100a6565b61012161011c366004610ed8565b61026e565b6040516100a69190610f86565b61009c7f000000000000000000000000000000000000000000000000000000000000000081565b610168610163366004610ed8565b61029f565b6040516100a6929190610ff4565b6100f67f000000000000000000000000000000000000000000000000000000000000000081565b6101b06101ab366004610ed8565b6102b4565b6040516100a69190611098565b5f5f5f6101c9856104d9565b90925090505f5b825181101561025057846001600160a01b03168382815181106101f5576101f56110fb565b60200260200101516001600160a01b0316036102485781818151811061021d5761021d6110fb565b60200260200101515f81518110610236576102366110fb565b60200260200101519350505050610257565b6001016101d0565b505f925050505b92915050565b610265610df7565b610257826107e8565b6060610279826107e8565b6040516020016102899190610f33565b6040516020818303038152906040529050919050565b6060806102ab836104d9565b91509150915091565b60605f5f6102c1846104d9565b915091505f82516001600160401b038111156102df576102df61110f565b60405190808252806020026020018201604052801561031857816020015b610305610e36565b8152602001906001900390816102fd5790505b5090505f5b83518110156104d0577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663bd30a0b987868481518110610368576103686110fb565b60200260200101516040518363ffffffff1660e01b815260040161038d92919061115c565b602060405180830381865afa1580156103a8573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103cc9190611182565b156104c8575f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639a43e3fb88878581518110610414576104146110fb565b60200260200101516040518363ffffffff1660e01b815260040161043992919061115c565b60c060405180830381865afa158015610454573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104789190611246565b509050604051806040016040528082815260200185848151811061049e5761049e6110fb565b60200260200101518152508383815181106104bb576104bb6110fb565b6020026020010181905250505b60010161031d565b50949350505050565b6060807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316636e875dba846040518263ffffffff1660e01b815260040161052891906112c3565b5f60405180830381865afa158015610542573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261056991908101906112f3565b91505f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316634177a87c856040518263ffffffff1660e01b81526004016105b891906112c3565b5f60405180830381865afa1580156105d2573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526105f99190810190611391565b90505f6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016632bab2c4a8686856106587f000000000000000000000000000000000000000000000000000000000000000043611434565b6040518563ffffffff1660e01b81526004016106779493929190611447565b5f60405180830381865afa158015610691573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526106b891908101906114c0565b905083516001600160401b038111156106d3576106d361110f565b60405190808252806020026020018201604052801561070657816020015b60608152602001906001900390816106f15790505b5092505f5b84518110156107e057604080516001808252818301909252906020808301908036833701905050848281518110610744576107446110fb565b60209081029190910101525f5b83518110156107d75782828151811061076c5761076c6110fb565b60200260200101518181518110610785576107856110fb565b602002602001015185838151811061079f5761079f6110fb565b60200260200101515f815181106107b8576107b86110fb565b602002602001018181516107cc9190611434565b905250600101610751565b5060010161070b565b505050915091565b6107f0610df7565b5f5f6107fb846104d9565b9150915081515f036108865760405180608001604052805f5f1b81526020015f815260200160405180604001604052805f81526020015f81525081526020015f6001600160401b038111156108525761085261110f565b60405190808252806020026020018201604052801561087b578160200160208202803683370190505b509052949350505050565b5f815f81518110610899576108996110fb565b60200260200101515190505f816001600160401b038111156108bd576108bd61110f565b6040519080825280602002602001820160405280156108e6578160200160208202803683370190505b5090505f84516001600160401b038111156109035761090361110f565b60405190808252806020026020018201604052801561092c578160200160208202803683370190505b50905061094a60405180604001604052805f81526020015f81525090565b5f5b8651811015610ba0577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663bd30a0b98a898481518110610997576109976110fb565b60200260200101516040518363ffffffff1660e01b81526004016109bc92919061115c565b602060405180830381865afa1580156109d7573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109fb9190611182565b15610b98575f5b85811015610a6e57868281518110610a1c57610a1c6110fb565b60200260200101518181518110610a3557610a356110fb565b6020026020010151858281518110610a4f57610a4f6110fb565b60200260200101818151610a639190611434565b905250600101610a02565b505f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639a43e3fb8b8a8581518110610ab257610ab26110fb565b60200260200101516040518363ffffffff1660e01b8152600401610ad792919061115c565b60c060405180830381865afa158015610af2573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b169190611246565b5090506040518060400160405280828152602001888481518110610b3c57610b3c6110fb565b6020026020010151815250604051602001610b5791906115cc565b60405160208183030381529060405280519060200120848381518110610b7f57610b7f6110fb565b6020908102919091010152610b948382610bd6565b9250505b60010161094c565b505f610bab83610c52565b6040805160808101825291825297516020820152968701919091525050606084015250909392505050565b604080518082019091525f8082526020820152610bf1610e57565b835181526020808501518183015283516040808401919091529084015160608301525f908360808460066107d05a03fa90508080610c2b57fe5b5080610c4a5760405163d4b68fd760e01b815260040160405180910390fd5b505092915050565b5f60015b8251811015610c7157610c6a6002826115de565b9050610c56565b5f816001600160401b03811115610c8a57610c8a61110f565b604051908082528060200260200182016040528015610cb3578160200160208202803683370190505b5090505f5b8451811015610d0057848181518110610cd357610cd36110fb565b6020026020010151828281518110610ced57610ced6110fb565b6020908102919091010152600101610cb8565b505b81600114610dd4575f610d166002846115f5565b90505f5b81811015610dcc5782610d2e8260026115de565b81518110610d3e57610d3e6110fb565b602002602001015183826002610d5491906115de565b610d5f906001611434565b81518110610d6f57610d6f6110fb565b6020026020010151604051602001610d91929190918252602082015260400190565b60405160208183030381529060405280519060200120838281518110610db957610db96110fb565b6020908102919091010152600101610d1a565b509150610d02565b805f81518110610de657610de66110fb565b602002602001015192505050919050565b60405180608001604052805f81526020015f8152602001610e2960405180604001604052805f81526020015f81525090565b8152602001606081525090565b604080516080810182525f9181018281526060820192909252908190610e29565b60405180608001604052806004906020820280368337509192915050565b5f60408284031215610e85575f5ffd5b50919050565b6001600160a01b0381168114610e9f575f5ffd5b50565b5f5f60608385031215610eb3575f5ffd5b610ebd8484610e75565b91506040830135610ecd81610e8b565b809150509250929050565b5f60408284031215610ee8575f5ffd5b610ef28383610e75565b9392505050565b5f8151808452602084019350602083015f5b82811015610f29578151865260209586019590910190600101610f0b565b5093949350505050565b6020815281516020820152602082015160408201525f6040830151610f65606084018280518252602090810151910152565b50606083015160a080840152610f7e60c0840182610ef9565b949350505050565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b5f8151808452602084019350602083015f5b82811015610f295781516001600160a01b0316865260209586019590910190600101610fcd565b604081525f6110066040830185610fbb565b828103602084015280845180835260208301915060208160051b840101602087015f5b8381101561105b57601f19868403018552611045838351610ef9565b6020958601959093509190910190600101611029565b509098975050505050505050565b61107e82825180518252602090810151910152565b5f602082015160606040850152610f7e6060850182610ef9565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b828110156110ef57603f198786030184526110da858351611069565b945060209384019391909101906001016110be565b50929695505050505050565b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52604160045260245ffd5b803561112e81610e8b565b6001600160a01b03168252602081013563ffffffff8116808214611150575f5ffd5b80602085015250505050565b6060810161116a8285611123565b6001600160a01b039290921660409190910152919050565b5f60208284031215611192575f5ffd5b81518015158114610ef2575f5ffd5b604080519081016001600160401b03811182821017156111c3576111c361110f565b60405290565b604051601f8201601f191681016001600160401b03811182821017156111f1576111f161110f565b604052919050565b5f82601f830112611208575f5ffd5b6112106111a1565b806040840185811115611221575f5ffd5b845b8181101561123b578051845260209384019301611223565b509095945050505050565b5f5f82840360c0811215611258575f5ffd5b6040811215611265575f5ffd5b61126d6111a1565b845181526020808601519082015292506080603f198201121561128e575f5ffd5b506112976111a1565b6112a485604086016111f9565b81526112b385608086016111f9565b6020820152809150509250929050565b604081016102578284611123565b5f6001600160401b038211156112e9576112e961110f565b5060051b60200190565b5f60208284031215611303575f5ffd5b81516001600160401b03811115611318575f5ffd5b8201601f81018413611328575f5ffd5b805161133b611336826112d1565b6111c9565b8082825260208201915060208360051b85010192508683111561135c575f5ffd5b6020840193505b8284101561138757835161137681610e8b565b825260209384019390910190611363565b9695505050505050565b5f602082840312156113a1575f5ffd5b81516001600160401b038111156113b6575f5ffd5b8201601f810184136113c6575f5ffd5b80516113d4611336826112d1565b8082825260208201915060208360051b8501019250868311156113f5575f5ffd5b6020840193505b8284101561138757835161140f81610e8b565b8252602093840193909101906113fc565b634e487b7160e01b5f52601160045260245ffd5b8082018082111561025757610257611420565b6114518186611123565b60a060408201525f61146660a0830186610fbb565b8281036060840152845180825260208087019201905f5b818110156114a45783516001600160a01b031683526020938401939092019160010161147d565b5050809250505063ffffffff8316608083015295945050505050565b5f602082840312156114d0575f5ffd5b81516001600160401b038111156114e5575f5ffd5b8201601f810184136114f5575f5ffd5b8051611503611336826112d1565b8082825260208201915060208360051b850101925086831115611524575f5ffd5b602084015b838110156115c15780516001600160401b03811115611546575f5ffd5b8501603f81018913611556575f5ffd5b6020810151611567611336826112d1565b808282526020820191506020808460051b8601010192508b83111561158a575f5ffd5b6040840193505b828410156115ac578351825260209384019390910190611591565b86525050602093840193919091019050611529565b509695505050505050565b602081525f610ef26020830184611069565b808202811582820484141761025757610257611420565b5f8261160f57634e487b7160e01b5f52601260045260245ffd5b50049056fea264697066735822122069947d39ac52bbd4772700d9d42b03f7fbe0444626961b3a020e8177cda39ca664736f6c634300081b0033",
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
