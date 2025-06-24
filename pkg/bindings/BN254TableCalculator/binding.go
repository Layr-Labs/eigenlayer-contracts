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
	Bin: "0x60e060405234801561000f575f5ffd5b506040516118b23803806118b283398101604081905261002e91610060565b6001600160a01b03928316608052911660a05260c0526100a0565b6001600160a01b038116811461005d575f5ffd5b50565b5f5f5f60608486031215610072575f5ffd5b835161007d81610049565b602085015190935061008e81610049565b80925050604084015190509250925092565b60805160a05160c0516117b16101015f395f8181610133015261063401525f818161017b015281816104df0152818161056f015261060701525f818160d401528181610328015281816103d401528181610a040152610b1f01526117b15ff3fe608060405234801561000f575f5ffd5b5060043610610085575f3560e01c80635e120ffc116100585780635e120ffc1461012e57806371ca71d914610155578063ca8aa7c714610176578063cf2d90ef1461019d575f5ffd5b80631088794a14610089578063124c87e0146100af5780633ec45c7e146100cf57806341ee6d0e1461010e575b5f5ffd5b61009c610097366004610ff1565b6101bd565b6040519081526020015b60405180910390f35b6100c26100bd366004611027565b61025d565b6040516100a69190611082565b6100f67f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016100a6565b61012161011c366004611027565b61026e565b6040516100a691906110d5565b61009c7f000000000000000000000000000000000000000000000000000000000000000081565b610168610163366004611027565b61029f565b6040516100a6929190611143565b6100f67f000000000000000000000000000000000000000000000000000000000000000081565b6101b06101ab366004611027565b6102b4565b6040516100a691906111e7565b5f5f5f6101c9856104d9565b90925090505f5b825181101561025057846001600160a01b03168382815181106101f5576101f561124a565b60200260200101516001600160a01b0316036102485781818151811061021d5761021d61124a565b60200260200101515f815181106102365761023661124a565b60200260200101519350505050610257565b6001016101d0565b505f925050505b92915050565b610265610f46565b61025782610894565b606061027982610894565b6040516020016102899190611082565b6040516020818303038152906040529050919050565b6060806102ab836104d9565b91509150915091565b60605f5f6102c1846104d9565b915091505f82516001600160401b038111156102df576102df61125e565b60405190808252806020026020018201604052801561031857816020015b610305610f85565b8152602001906001900390816102fd5790505b5090505f5b83518110156104d0577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663bd30a0b9878684815181106103685761036861124a565b60200260200101516040518363ffffffff1660e01b815260040161038d9291906112ab565b602060405180830381865afa1580156103a8573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103cc91906112d1565b156104c8575f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639a43e3fb888785815181106104145761041461124a565b60200260200101516040518363ffffffff1660e01b81526004016104399291906112ab565b60c060405180830381865afa158015610454573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104789190611395565b509050604051806040016040528082815260200185848151811061049e5761049e61124a565b60200260200101518152508383815181106104bb576104bb61124a565b6020026020010181905250505b60010161031d565b50949350505050565b6060805f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316636e875dba856040518263ffffffff1660e01b81526004016105299190611412565b5f60405180830381865afa158015610543573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261056a9190810190611442565b90505f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316634177a87c866040518263ffffffff1660e01b81526004016105b99190611412565b5f60405180830381865afa1580156105d3573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526105fa91908101906114e0565b90505f6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016632bab2c4a8785856106597f000000000000000000000000000000000000000000000000000000000000000043611583565b6040518563ffffffff1660e01b81526004016106789493929190611596565b5f60405180830381865afa158015610692573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526106b9919081019061160f565b905082516001600160401b038111156106d4576106d461125e565b6040519080825280602002602001820160405280156106fd578160200160208202803683370190505b50945082516001600160401b038111156107195761071961125e565b60405190808252806020026020018201604052801561074c57816020015b60608152602001906001900390816107375790505b5093505f805b8451811015610884575f805b85518110156107af578483815181106107795761077961124a565b602002602001015181815181106107925761079261124a565b6020026020010151826107a59190611583565b915060010161075e565b50801561087b576040805160018082528183019092529060208083019080368337019050508784815181106107e6576107e661124a565b6020026020010181905250808784815181106108045761080461124a565b60200260200101515f8151811061081d5761081d61124a565b60200260200101818152505085828151811061083b5761083b61124a565b60200260200101518884815181106108555761085561124a565b6001600160a01b0390921660209283029190910190910152826108778161171b565b9350505b50600101610752565b5080865280855250505050915091565b61089c610f46565b5f5f6108a7846104d9565b9150915080515f036109325760405180608001604052805f5f1b81526020015f815260200160405180604001604052805f81526020015f81525081526020015f6001600160401b038111156108fe576108fe61125e565b604051908082528060200260200182016040528015610927578160200160208202803683370190505b509052949350505050565b5f815f815181106109455761094561124a565b60200260200101515190505f816001600160401b038111156109695761096961125e565b604051908082528060200260200182016040528015610992578160200160208202803683370190505b5090505f84516001600160401b038111156109af576109af61125e565b6040519080825280602002602001820160405280156109d8578160200160208202803683370190505b5090506109f660405180604001604052805f81526020015f81525090565b5f805b8751811015610c5a577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663bd30a0b98b8a8481518110610a4457610a4461124a565b60200260200101516040518363ffffffff1660e01b8152600401610a699291906112ab565b602060405180830381865afa158015610a84573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610aa891906112d1565b15610c52575f5b86811015610b1b57878281518110610ac957610ac961124a565b60200260200101518181518110610ae257610ae261124a565b6020026020010151868281518110610afc57610afc61124a565b60200260200101818151610b109190611583565b905250600101610aaf565b505f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639a43e3fb8c8b8581518110610b5f57610b5f61124a565b60200260200101516040518363ffffffff1660e01b8152600401610b849291906112ab565b60c060405180830381865afa158015610b9f573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610bc39190611395565b5090506040518060400160405280828152602001898481518110610be957610be961124a565b6020026020010151815250604051602001610c049190611733565b60405160208183030381529060405280519060200120858381518110610c2c57610c2c61124a565b6020908102919091010152610c418482610d25565b935082610c4d8161171b565b935050505b6001016109f9565b50805f03610ce65760405180608001604052805f5f1b81526020015f815260200160405180604001604052805f81526020015f81525081526020015f6001600160401b03811115610cad57610cad61125e565b604051908082528060200260200182016040528015610cd6578160200160208202803683370190505b5090529998505050505050505050565b8083525f610cf384610da1565b905060405180608001604052808281526020018381526020018481526020018681525098505050505050505050919050565b604080518082019091525f8082526020820152610d40610fa6565b835181526020808501518183015283516040808401919091529084015160608301525f908360808460066107d05a03fa90508080610d7a57fe5b5080610d995760405163d4b68fd760e01b815260040160405180910390fd5b505092915050565b5f60015b8251811015610dc057610db9600282611745565b9050610da5565b5f816001600160401b03811115610dd957610dd961125e565b604051908082528060200260200182016040528015610e02578160200160208202803683370190505b5090505f5b8451811015610e4f57848181518110610e2257610e2261124a565b6020026020010151828281518110610e3c57610e3c61124a565b6020908102919091010152600101610e07565b505b81600114610f23575f610e6560028461175c565b90505f5b81811015610f1b5782610e7d826002611745565b81518110610e8d57610e8d61124a565b602002602001015183826002610ea39190611745565b610eae906001611583565b81518110610ebe57610ebe61124a565b6020026020010151604051602001610ee0929190918252602082015260400190565b60405160208183030381529060405280519060200120838281518110610f0857610f0861124a565b6020908102919091010152600101610e69565b509150610e51565b805f81518110610f3557610f3561124a565b602002602001015192505050919050565b60405180608001604052805f81526020015f8152602001610f7860405180604001604052805f81526020015f81525090565b8152602001606081525090565b604080516080810182525f9181018281526060820192909252908190610f78565b60405180608001604052806004906020820280368337509192915050565b5f60408284031215610fd4575f5ffd5b50919050565b6001600160a01b0381168114610fee575f5ffd5b50565b5f5f60608385031215611002575f5ffd5b61100c8484610fc4565b9150604083013561101c81610fda565b809150509250929050565b5f60408284031215611037575f5ffd5b6110418383610fc4565b9392505050565b5f8151808452602084019350602083015f5b8281101561107857815186526020958601959091019060010161105a565b5093949350505050565b6020815281516020820152602082015160408201525f60408301516110b4606084018280518252602090810151910152565b50606083015160a0808401526110cd60c0840182611048565b949350505050565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b5f8151808452602084019350602083015f5b828110156110785781516001600160a01b031686526020958601959091019060010161111c565b604081525f611155604083018561110a565b828103602084015280845180835260208301915060208160051b840101602087015f5b838110156111aa57601f19868403018552611194838351611048565b6020958601959093509190910190600101611178565b509098975050505050505050565b6111cd82825180518252602090810151910152565b5f6020820151606060408501526110cd6060850182611048565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b8281101561123e57603f198786030184526112298583516111b8565b9450602093840193919091019060010161120d565b50929695505050505050565b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52604160045260245ffd5b803561127d81610fda565b6001600160a01b03168252602081013563ffffffff811680821461129f575f5ffd5b80602085015250505050565b606081016112b98285611272565b6001600160a01b039290921660409190910152919050565b5f602082840312156112e1575f5ffd5b81518015158114611041575f5ffd5b604080519081016001600160401b03811182821017156113125761131261125e565b60405290565b604051601f8201601f191681016001600160401b03811182821017156113405761134061125e565b604052919050565b5f82601f830112611357575f5ffd5b61135f6112f0565b806040840185811115611370575f5ffd5b845b8181101561138a578051845260209384019301611372565b509095945050505050565b5f5f82840360c08112156113a7575f5ffd5b60408112156113b4575f5ffd5b6113bc6112f0565b845181526020808601519082015292506080603f19820112156113dd575f5ffd5b506113e66112f0565b6113f38560408601611348565b81526114028560808601611348565b6020820152809150509250929050565b604081016102578284611272565b5f6001600160401b038211156114385761143861125e565b5060051b60200190565b5f60208284031215611452575f5ffd5b81516001600160401b03811115611467575f5ffd5b8201601f81018413611477575f5ffd5b805161148a61148582611420565b611318565b8082825260208201915060208360051b8501019250868311156114ab575f5ffd5b6020840193505b828410156114d65783516114c581610fda565b8252602093840193909101906114b2565b9695505050505050565b5f602082840312156114f0575f5ffd5b81516001600160401b03811115611505575f5ffd5b8201601f81018413611515575f5ffd5b805161152361148582611420565b8082825260208201915060208360051b850101925086831115611544575f5ffd5b6020840193505b828410156114d657835161155e81610fda565b82526020938401939091019061154b565b634e487b7160e01b5f52601160045260245ffd5b808201808211156102575761025761156f565b6115a08186611272565b60a060408201525f6115b560a083018661110a565b8281036060840152845180825260208087019201905f5b818110156115f35783516001600160a01b03168352602093840193909201916001016115cc565b5050809250505063ffffffff8316608083015295945050505050565b5f6020828403121561161f575f5ffd5b81516001600160401b03811115611634575f5ffd5b8201601f81018413611644575f5ffd5b805161165261148582611420565b8082825260208201915060208360051b850101925086831115611673575f5ffd5b602084015b838110156117105780516001600160401b03811115611695575f5ffd5b8501603f810189136116a5575f5ffd5b60208101516116b661148582611420565b808282526020820191506020808460051b8601010192508b8311156116d9575f5ffd5b6040840193505b828410156116fb5783518252602093840193909101906116e0565b86525050602093840193919091019050611678565b509695505050505050565b5f6001820161172c5761172c61156f565b5060010190565b602081525f61104160208301846111b8565b80820281158282048414176102575761025761156f565b5f8261177657634e487b7160e01b5f52601260045260245ffd5b50049056fea26469706673582212201b5a651309a6a5dca4bf019229739f96a3dfbba74509ef6e3025e4059ecb818164736f6c634300081b0033",
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
