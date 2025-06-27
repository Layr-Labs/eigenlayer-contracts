// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ECDSATableCalculator

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

// IECDSATableCalculatorTypesECDSAOperatorInfo is an auto generated low-level Go binding around an user-defined struct.
type IECDSATableCalculatorTypesECDSAOperatorInfo struct {
	Pubkey  common.Address
	Weights []*big.Int
}

// OperatorSet is an auto generated low-level Go binding around an user-defined struct.
type OperatorSet struct {
	Avs common.Address
	Id  uint32
}

// ECDSATableCalculatorMetaData contains all meta data concerning the ECDSATableCalculator contract.
var ECDSATableCalculatorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_keyRegistrar\",\"type\":\"address\",\"internalType\":\"contractIKeyRegistrar\"},{\"name\":\"_allocationManager\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"},{\"name\":\"_LOOKAHEAD_BLOCKS\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"LOOKAHEAD_BLOCKS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allocationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateOperatorTable\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"operatorInfos\",\"type\":\"tuple[]\",\"internalType\":\"structIECDSATableCalculatorTypes.ECDSAOperatorInfo[]\",\"components\":[{\"name\":\"pubkey\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"weights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateOperatorTableBytes\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"operatorTableBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorWeight\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorWeights\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"weights\",\"type\":\"uint256[][]\",\"internalType\":\"uint256[][]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"keyRegistrar\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIKeyRegistrar\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"LookaheadBlocksSet\",\"inputs\":[{\"name\":\"lookaheadBlocks\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"LookaheadBlocksTooHigh\",\"inputs\":[]}]",
	Bin: "0x60e060405234801561000f575f5ffd5b5060405161108938038061108983398101604081905261002e91610060565b6001600160a01b03928316608052911660a05260c0526100a0565b6001600160a01b038116811461005d575f5ffd5b50565b5f5f5f60608486031215610072575f5ffd5b835161007d81610049565b602085015190935061008e81610049565b80925050604084015190509250925092565b60805160a05160c051610f966100f35f395f818161012801526103de01525f8181610170015281816102890152818161031901526103b101525f818160c90152818161070e01526107ba0152610f965ff3fe608060405234801561000f575f5ffd5b506004361061007a575f3560e01c806341ee6d0e1161005857806341ee6d0e146101035780635e120ffc1461012357806371ca71d91461014a578063ca8aa7c71461016b575f5ffd5b80631088794a1461007e578063124c87e0146100a45780633ec45c7e146100c4575b5f5ffd5b61009161008c366004610951565b610192565b6040519081526020015b60405180910390f35b6100b76100b2366004610987565b610232565b60405161009b91906109e2565b6100eb7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b03909116815260200161009b565b610116610111366004610987565b61023d565b60405161009b9190610a65565b6100917f000000000000000000000000000000000000000000000000000000000000000081565b61015d610158366004610987565b61026e565b60405161009b929190610ad3565b6100eb7f000000000000000000000000000000000000000000000000000000000000000081565b5f5f5f61019e85610283565b90925090505f5b825181101561022557846001600160a01b03168382815181106101ca576101ca610b48565b60200260200101516001600160a01b03160361021d578181815181106101f2576101f2610b48565b60200260200101515f8151811061020b5761020b610b48565b6020026020010151935050505061022c565b6001016101a5565b505f925050505b92915050565b606061022c82610640565b606061024882610640565b60405160200161025891906109e2565b6040516020818303038152906040529050919050565b60608061027a83610283565b91509150915091565b6060805f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316636e875dba856040518263ffffffff1660e01b81526004016102d39190610b95565b5f60405180830381865afa1580156102ed573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526103149190810190610c0b565b90505f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316634177a87c866040518263ffffffff1660e01b81526004016103639190610b95565b5f60405180830381865afa15801561037d573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526103a49190810190610caa565b90505f6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016632bab2c4a8785856104037f000000000000000000000000000000000000000000000000000000000000000043610d4e565b6040518563ffffffff1660e01b81526004016104229493929190610d61565b5f60405180830381865afa15801561043c573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526104639190810190610dda565b9050825167ffffffffffffffff81111561047f5761047f610ba3565b6040519080825280602002602001820160405280156104a8578160200160208202803683370190505b509450825167ffffffffffffffff8111156104c5576104c5610ba3565b6040519080825280602002602001820160405280156104f857816020015b60608152602001906001900390816104e35790505b5093505f805b8451811015610630575f805b855181101561055b5784838151811061052557610525610b48565b6020026020010151818151811061053e5761053e610b48565b6020026020010151826105519190610d4e565b915060010161050a565b5080156106275760408051600180825281830190925290602080830190803683370190505087848151811061059257610592610b48565b6020026020010181905250808784815181106105b0576105b0610b48565b60200260200101515f815181106105c9576105c9610b48565b6020026020010181815250508582815181106105e7576105e7610b48565b602002602001015188848151811061060157610601610b48565b6001600160a01b03909216602092830291909101909101528261062381610ee8565b9350505b506001016104fe565b5080865280855250505050915091565b60605f5f61064d84610283565b9150915080515f0361069e57604080515f8082526020820190925290610695565b604080518082019091525f81526060602082015281526020019060019003908161066e5790505b50949350505050565b815167ffffffffffffffff8111156106b8576106b8610ba3565b6040519080825280602002602001820160405280156106fd57816020015b604080518082019091525f8152606060208201528152602001906001900390816106d65790505b5092505f805b83518110156108cc577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663bd30a0b98786848151811061074e5761074e610b48565b60200260200101516040518363ffffffff1660e01b8152600401610773929190610f00565b602060405180830381865afa15801561078e573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107b29190610f26565b156108c4575f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316633b32a7bd888785815181106107fa576107fa610b48565b60200260200101516040518363ffffffff1660e01b815260040161081f929190610f00565b602060405180830381865afa15801561083a573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061085e9190610f45565b90506040518060400160405280826001600160a01b0316815260200185848151811061088c5761088c610b48565b60200260200101518152508684815181106108a9576108a9610b48565b602002602001018190525082806108bf90610ee8565b935050505b600101610703565b50805f0361091a57604080515f8082526020820190925290610910565b604080518082019091525f8152606060208201528152602001906001900390816108e95790505b5095945050505050565b8352509092915050565b5f60408284031215610934575f5ffd5b50919050565b6001600160a01b038116811461094e575f5ffd5b50565b5f5f60608385031215610962575f5ffd5b61096c8484610924565b9150604083013561097c8161093a565b809150509250929050565b5f60408284031215610997575f5ffd5b6109a18383610924565b9392505050565b5f8151808452602084019350602083015f5b828110156109d85781518652602095860195909101906001016109ba565b5093949350505050565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b82811015610a5957868503603f19018452815180516001600160a01b03168652602090810151604091870182905290610a43908701826109a8565b9550506020938401939190910190600101610a08565b50929695505050505050565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b5f8151808452602084019350602083015f5b828110156109d85781516001600160a01b0316865260209586019590910190600101610aac565b604081525f610ae56040830185610a9a565b828103602084015280845180835260208301915060208160051b840101602087015f5b83811015610b3a57601f19868403018552610b248383516109a8565b6020958601959093509190910190600101610b08565b509098975050505050505050565b634e487b7160e01b5f52603260045260245ffd5b8035610b678161093a565b6001600160a01b03168252602081013563ffffffff8116808214610b89575f5ffd5b80602085015250505050565b6040810161022c8284610b5c565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f1916810167ffffffffffffffff81118282101715610be057610be0610ba3565b604052919050565b5f67ffffffffffffffff821115610c0157610c01610ba3565b5060051b60200190565b5f60208284031215610c1b575f5ffd5b815167ffffffffffffffff811115610c31575f5ffd5b8201601f81018413610c41575f5ffd5b8051610c54610c4f82610be8565b610bb7565b8082825260208201915060208360051b850101925086831115610c75575f5ffd5b6020840193505b82841015610ca0578351610c8f8161093a565b825260209384019390910190610c7c565b9695505050505050565b5f60208284031215610cba575f5ffd5b815167ffffffffffffffff811115610cd0575f5ffd5b8201601f81018413610ce0575f5ffd5b8051610cee610c4f82610be8565b8082825260208201915060208360051b850101925086831115610d0f575f5ffd5b6020840193505b82841015610ca0578351610d298161093a565b825260209384019390910190610d16565b634e487b7160e01b5f52601160045260245ffd5b8082018082111561022c5761022c610d3a565b610d6b8186610b5c565b60a060408201525f610d8060a0830186610a9a565b8281036060840152845180825260208087019201905f5b81811015610dbe5783516001600160a01b0316835260209384019390920191600101610d97565b5050809250505063ffffffff8316608083015295945050505050565b5f60208284031215610dea575f5ffd5b815167ffffffffffffffff811115610e00575f5ffd5b8201601f81018413610e10575f5ffd5b8051610e1e610c4f82610be8565b8082825260208201915060208360051b850101925086831115610e3f575f5ffd5b602084015b83811015610edd57805167ffffffffffffffff811115610e62575f5ffd5b8501603f81018913610e72575f5ffd5b6020810151610e83610c4f82610be8565b808282526020820191506020808460051b8601010192508b831115610ea6575f5ffd5b6040840193505b82841015610ec8578351825260209384019390910190610ead565b86525050602093840193919091019050610e44565b509695505050505050565b5f60018201610ef957610ef9610d3a565b5060010190565b60608101610f0e8285610b5c565b6001600160a01b039290921660409190910152919050565b5f60208284031215610f36575f5ffd5b815180151581146109a1575f5ffd5b5f60208284031215610f55575f5ffd5b81516109a18161093a56fea2646970667358221220d84a68511f506942f3e0fa8d48df5dd9dda5a4c230f3307af1fd784b19b4d5c264736f6c634300081b0033",
}

// ECDSATableCalculatorABI is the input ABI used to generate the binding from.
// Deprecated: Use ECDSATableCalculatorMetaData.ABI instead.
var ECDSATableCalculatorABI = ECDSATableCalculatorMetaData.ABI

// ECDSATableCalculatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ECDSATableCalculatorMetaData.Bin instead.
var ECDSATableCalculatorBin = ECDSATableCalculatorMetaData.Bin

// DeployECDSATableCalculator deploys a new Ethereum contract, binding an instance of ECDSATableCalculator to it.
func DeployECDSATableCalculator(auth *bind.TransactOpts, backend bind.ContractBackend, _keyRegistrar common.Address, _allocationManager common.Address, _LOOKAHEAD_BLOCKS *big.Int) (common.Address, *types.Transaction, *ECDSATableCalculator, error) {
	parsed, err := ECDSATableCalculatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ECDSATableCalculatorBin), backend, _keyRegistrar, _allocationManager, _LOOKAHEAD_BLOCKS)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ECDSATableCalculator{ECDSATableCalculatorCaller: ECDSATableCalculatorCaller{contract: contract}, ECDSATableCalculatorTransactor: ECDSATableCalculatorTransactor{contract: contract}, ECDSATableCalculatorFilterer: ECDSATableCalculatorFilterer{contract: contract}}, nil
}

// ECDSATableCalculator is an auto generated Go binding around an Ethereum contract.
type ECDSATableCalculator struct {
	ECDSATableCalculatorCaller     // Read-only binding to the contract
	ECDSATableCalculatorTransactor // Write-only binding to the contract
	ECDSATableCalculatorFilterer   // Log filterer for contract events
}

// ECDSATableCalculatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type ECDSATableCalculatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSATableCalculatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ECDSATableCalculatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSATableCalculatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ECDSATableCalculatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSATableCalculatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ECDSATableCalculatorSession struct {
	Contract     *ECDSATableCalculator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ECDSATableCalculatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ECDSATableCalculatorCallerSession struct {
	Contract *ECDSATableCalculatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// ECDSATableCalculatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ECDSATableCalculatorTransactorSession struct {
	Contract     *ECDSATableCalculatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// ECDSATableCalculatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type ECDSATableCalculatorRaw struct {
	Contract *ECDSATableCalculator // Generic contract binding to access the raw methods on
}

// ECDSATableCalculatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ECDSATableCalculatorCallerRaw struct {
	Contract *ECDSATableCalculatorCaller // Generic read-only contract binding to access the raw methods on
}

// ECDSATableCalculatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ECDSATableCalculatorTransactorRaw struct {
	Contract *ECDSATableCalculatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewECDSATableCalculator creates a new instance of ECDSATableCalculator, bound to a specific deployed contract.
func NewECDSATableCalculator(address common.Address, backend bind.ContractBackend) (*ECDSATableCalculator, error) {
	contract, err := bindECDSATableCalculator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ECDSATableCalculator{ECDSATableCalculatorCaller: ECDSATableCalculatorCaller{contract: contract}, ECDSATableCalculatorTransactor: ECDSATableCalculatorTransactor{contract: contract}, ECDSATableCalculatorFilterer: ECDSATableCalculatorFilterer{contract: contract}}, nil
}

// NewECDSATableCalculatorCaller creates a new read-only instance of ECDSATableCalculator, bound to a specific deployed contract.
func NewECDSATableCalculatorCaller(address common.Address, caller bind.ContractCaller) (*ECDSATableCalculatorCaller, error) {
	contract, err := bindECDSATableCalculator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ECDSATableCalculatorCaller{contract: contract}, nil
}

// NewECDSATableCalculatorTransactor creates a new write-only instance of ECDSATableCalculator, bound to a specific deployed contract.
func NewECDSATableCalculatorTransactor(address common.Address, transactor bind.ContractTransactor) (*ECDSATableCalculatorTransactor, error) {
	contract, err := bindECDSATableCalculator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ECDSATableCalculatorTransactor{contract: contract}, nil
}

// NewECDSATableCalculatorFilterer creates a new log filterer instance of ECDSATableCalculator, bound to a specific deployed contract.
func NewECDSATableCalculatorFilterer(address common.Address, filterer bind.ContractFilterer) (*ECDSATableCalculatorFilterer, error) {
	contract, err := bindECDSATableCalculator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ECDSATableCalculatorFilterer{contract: contract}, nil
}

// bindECDSATableCalculator binds a generic wrapper to an already deployed contract.
func bindECDSATableCalculator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ECDSATableCalculatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECDSATableCalculator *ECDSATableCalculatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ECDSATableCalculator.Contract.ECDSATableCalculatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECDSATableCalculator *ECDSATableCalculatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECDSATableCalculator.Contract.ECDSATableCalculatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECDSATableCalculator *ECDSATableCalculatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECDSATableCalculator.Contract.ECDSATableCalculatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECDSATableCalculator *ECDSATableCalculatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ECDSATableCalculator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECDSATableCalculator *ECDSATableCalculatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECDSATableCalculator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECDSATableCalculator *ECDSATableCalculatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECDSATableCalculator.Contract.contract.Transact(opts, method, params...)
}

// LOOKAHEADBLOCKS is a free data retrieval call binding the contract method 0x5e120ffc.
//
// Solidity: function LOOKAHEAD_BLOCKS() view returns(uint256)
func (_ECDSATableCalculator *ECDSATableCalculatorCaller) LOOKAHEADBLOCKS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ECDSATableCalculator.contract.Call(opts, &out, "LOOKAHEAD_BLOCKS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LOOKAHEADBLOCKS is a free data retrieval call binding the contract method 0x5e120ffc.
//
// Solidity: function LOOKAHEAD_BLOCKS() view returns(uint256)
func (_ECDSATableCalculator *ECDSATableCalculatorSession) LOOKAHEADBLOCKS() (*big.Int, error) {
	return _ECDSATableCalculator.Contract.LOOKAHEADBLOCKS(&_ECDSATableCalculator.CallOpts)
}

// LOOKAHEADBLOCKS is a free data retrieval call binding the contract method 0x5e120ffc.
//
// Solidity: function LOOKAHEAD_BLOCKS() view returns(uint256)
func (_ECDSATableCalculator *ECDSATableCalculatorCallerSession) LOOKAHEADBLOCKS() (*big.Int, error) {
	return _ECDSATableCalculator.Contract.LOOKAHEADBLOCKS(&_ECDSATableCalculator.CallOpts)
}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_ECDSATableCalculator *ECDSATableCalculatorCaller) AllocationManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ECDSATableCalculator.contract.Call(opts, &out, "allocationManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_ECDSATableCalculator *ECDSATableCalculatorSession) AllocationManager() (common.Address, error) {
	return _ECDSATableCalculator.Contract.AllocationManager(&_ECDSATableCalculator.CallOpts)
}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_ECDSATableCalculator *ECDSATableCalculatorCallerSession) AllocationManager() (common.Address, error) {
	return _ECDSATableCalculator.Contract.AllocationManager(&_ECDSATableCalculator.CallOpts)
}

// CalculateOperatorTable is a free data retrieval call binding the contract method 0x124c87e0.
//
// Solidity: function calculateOperatorTable((address,uint32) operatorSet) view returns((address,uint256[])[] operatorInfos)
func (_ECDSATableCalculator *ECDSATableCalculatorCaller) CalculateOperatorTable(opts *bind.CallOpts, operatorSet OperatorSet) ([]IECDSATableCalculatorTypesECDSAOperatorInfo, error) {
	var out []interface{}
	err := _ECDSATableCalculator.contract.Call(opts, &out, "calculateOperatorTable", operatorSet)

	if err != nil {
		return *new([]IECDSATableCalculatorTypesECDSAOperatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IECDSATableCalculatorTypesECDSAOperatorInfo)).(*[]IECDSATableCalculatorTypesECDSAOperatorInfo)

	return out0, err

}

// CalculateOperatorTable is a free data retrieval call binding the contract method 0x124c87e0.
//
// Solidity: function calculateOperatorTable((address,uint32) operatorSet) view returns((address,uint256[])[] operatorInfos)
func (_ECDSATableCalculator *ECDSATableCalculatorSession) CalculateOperatorTable(operatorSet OperatorSet) ([]IECDSATableCalculatorTypesECDSAOperatorInfo, error) {
	return _ECDSATableCalculator.Contract.CalculateOperatorTable(&_ECDSATableCalculator.CallOpts, operatorSet)
}

// CalculateOperatorTable is a free data retrieval call binding the contract method 0x124c87e0.
//
// Solidity: function calculateOperatorTable((address,uint32) operatorSet) view returns((address,uint256[])[] operatorInfos)
func (_ECDSATableCalculator *ECDSATableCalculatorCallerSession) CalculateOperatorTable(operatorSet OperatorSet) ([]IECDSATableCalculatorTypesECDSAOperatorInfo, error) {
	return _ECDSATableCalculator.Contract.CalculateOperatorTable(&_ECDSATableCalculator.CallOpts, operatorSet)
}

// CalculateOperatorTableBytes is a free data retrieval call binding the contract method 0x41ee6d0e.
//
// Solidity: function calculateOperatorTableBytes((address,uint32) operatorSet) view returns(bytes operatorTableBytes)
func (_ECDSATableCalculator *ECDSATableCalculatorCaller) CalculateOperatorTableBytes(opts *bind.CallOpts, operatorSet OperatorSet) ([]byte, error) {
	var out []interface{}
	err := _ECDSATableCalculator.contract.Call(opts, &out, "calculateOperatorTableBytes", operatorSet)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// CalculateOperatorTableBytes is a free data retrieval call binding the contract method 0x41ee6d0e.
//
// Solidity: function calculateOperatorTableBytes((address,uint32) operatorSet) view returns(bytes operatorTableBytes)
func (_ECDSATableCalculator *ECDSATableCalculatorSession) CalculateOperatorTableBytes(operatorSet OperatorSet) ([]byte, error) {
	return _ECDSATableCalculator.Contract.CalculateOperatorTableBytes(&_ECDSATableCalculator.CallOpts, operatorSet)
}

// CalculateOperatorTableBytes is a free data retrieval call binding the contract method 0x41ee6d0e.
//
// Solidity: function calculateOperatorTableBytes((address,uint32) operatorSet) view returns(bytes operatorTableBytes)
func (_ECDSATableCalculator *ECDSATableCalculatorCallerSession) CalculateOperatorTableBytes(operatorSet OperatorSet) ([]byte, error) {
	return _ECDSATableCalculator.Contract.CalculateOperatorTableBytes(&_ECDSATableCalculator.CallOpts, operatorSet)
}

// GetOperatorWeight is a free data retrieval call binding the contract method 0x1088794a.
//
// Solidity: function getOperatorWeight((address,uint32) operatorSet, address operator) view returns(uint256 weight)
func (_ECDSATableCalculator *ECDSATableCalculatorCaller) GetOperatorWeight(opts *bind.CallOpts, operatorSet OperatorSet, operator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ECDSATableCalculator.contract.Call(opts, &out, "getOperatorWeight", operatorSet, operator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOperatorWeight is a free data retrieval call binding the contract method 0x1088794a.
//
// Solidity: function getOperatorWeight((address,uint32) operatorSet, address operator) view returns(uint256 weight)
func (_ECDSATableCalculator *ECDSATableCalculatorSession) GetOperatorWeight(operatorSet OperatorSet, operator common.Address) (*big.Int, error) {
	return _ECDSATableCalculator.Contract.GetOperatorWeight(&_ECDSATableCalculator.CallOpts, operatorSet, operator)
}

// GetOperatorWeight is a free data retrieval call binding the contract method 0x1088794a.
//
// Solidity: function getOperatorWeight((address,uint32) operatorSet, address operator) view returns(uint256 weight)
func (_ECDSATableCalculator *ECDSATableCalculatorCallerSession) GetOperatorWeight(operatorSet OperatorSet, operator common.Address) (*big.Int, error) {
	return _ECDSATableCalculator.Contract.GetOperatorWeight(&_ECDSATableCalculator.CallOpts, operatorSet, operator)
}

// GetOperatorWeights is a free data retrieval call binding the contract method 0x71ca71d9.
//
// Solidity: function getOperatorWeights((address,uint32) operatorSet) view returns(address[] operators, uint256[][] weights)
func (_ECDSATableCalculator *ECDSATableCalculatorCaller) GetOperatorWeights(opts *bind.CallOpts, operatorSet OperatorSet) (struct {
	Operators []common.Address
	Weights   [][]*big.Int
}, error) {
	var out []interface{}
	err := _ECDSATableCalculator.contract.Call(opts, &out, "getOperatorWeights", operatorSet)

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
func (_ECDSATableCalculator *ECDSATableCalculatorSession) GetOperatorWeights(operatorSet OperatorSet) (struct {
	Operators []common.Address
	Weights   [][]*big.Int
}, error) {
	return _ECDSATableCalculator.Contract.GetOperatorWeights(&_ECDSATableCalculator.CallOpts, operatorSet)
}

// GetOperatorWeights is a free data retrieval call binding the contract method 0x71ca71d9.
//
// Solidity: function getOperatorWeights((address,uint32) operatorSet) view returns(address[] operators, uint256[][] weights)
func (_ECDSATableCalculator *ECDSATableCalculatorCallerSession) GetOperatorWeights(operatorSet OperatorSet) (struct {
	Operators []common.Address
	Weights   [][]*big.Int
}, error) {
	return _ECDSATableCalculator.Contract.GetOperatorWeights(&_ECDSATableCalculator.CallOpts, operatorSet)
}

// KeyRegistrar is a free data retrieval call binding the contract method 0x3ec45c7e.
//
// Solidity: function keyRegistrar() view returns(address)
func (_ECDSATableCalculator *ECDSATableCalculatorCaller) KeyRegistrar(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ECDSATableCalculator.contract.Call(opts, &out, "keyRegistrar")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// KeyRegistrar is a free data retrieval call binding the contract method 0x3ec45c7e.
//
// Solidity: function keyRegistrar() view returns(address)
func (_ECDSATableCalculator *ECDSATableCalculatorSession) KeyRegistrar() (common.Address, error) {
	return _ECDSATableCalculator.Contract.KeyRegistrar(&_ECDSATableCalculator.CallOpts)
}

// KeyRegistrar is a free data retrieval call binding the contract method 0x3ec45c7e.
//
// Solidity: function keyRegistrar() view returns(address)
func (_ECDSATableCalculator *ECDSATableCalculatorCallerSession) KeyRegistrar() (common.Address, error) {
	return _ECDSATableCalculator.Contract.KeyRegistrar(&_ECDSATableCalculator.CallOpts)
}

// ECDSATableCalculatorLookaheadBlocksSetIterator is returned from FilterLookaheadBlocksSet and is used to iterate over the raw logs and unpacked data for LookaheadBlocksSet events raised by the ECDSATableCalculator contract.
type ECDSATableCalculatorLookaheadBlocksSetIterator struct {
	Event *ECDSATableCalculatorLookaheadBlocksSet // Event containing the contract specifics and raw log

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
func (it *ECDSATableCalculatorLookaheadBlocksSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ECDSATableCalculatorLookaheadBlocksSet)
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
		it.Event = new(ECDSATableCalculatorLookaheadBlocksSet)
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
func (it *ECDSATableCalculatorLookaheadBlocksSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ECDSATableCalculatorLookaheadBlocksSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ECDSATableCalculatorLookaheadBlocksSet represents a LookaheadBlocksSet event raised by the ECDSATableCalculator contract.
type ECDSATableCalculatorLookaheadBlocksSet struct {
	LookaheadBlocks *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLookaheadBlocksSet is a free log retrieval operation binding the contract event 0xa41e64dd47db91b61b43ccea8b57d75abfa496f23efc708c22753c4bc9d68842.
//
// Solidity: event LookaheadBlocksSet(uint256 lookaheadBlocks)
func (_ECDSATableCalculator *ECDSATableCalculatorFilterer) FilterLookaheadBlocksSet(opts *bind.FilterOpts) (*ECDSATableCalculatorLookaheadBlocksSetIterator, error) {

	logs, sub, err := _ECDSATableCalculator.contract.FilterLogs(opts, "LookaheadBlocksSet")
	if err != nil {
		return nil, err
	}
	return &ECDSATableCalculatorLookaheadBlocksSetIterator{contract: _ECDSATableCalculator.contract, event: "LookaheadBlocksSet", logs: logs, sub: sub}, nil
}

// WatchLookaheadBlocksSet is a free log subscription operation binding the contract event 0xa41e64dd47db91b61b43ccea8b57d75abfa496f23efc708c22753c4bc9d68842.
//
// Solidity: event LookaheadBlocksSet(uint256 lookaheadBlocks)
func (_ECDSATableCalculator *ECDSATableCalculatorFilterer) WatchLookaheadBlocksSet(opts *bind.WatchOpts, sink chan<- *ECDSATableCalculatorLookaheadBlocksSet) (event.Subscription, error) {

	logs, sub, err := _ECDSATableCalculator.contract.WatchLogs(opts, "LookaheadBlocksSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ECDSATableCalculatorLookaheadBlocksSet)
				if err := _ECDSATableCalculator.contract.UnpackLog(event, "LookaheadBlocksSet", log); err != nil {
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
func (_ECDSATableCalculator *ECDSATableCalculatorFilterer) ParseLookaheadBlocksSet(log types.Log) (*ECDSATableCalculatorLookaheadBlocksSet, error) {
	event := new(ECDSATableCalculatorLookaheadBlocksSet)
	if err := _ECDSATableCalculator.contract.UnpackLog(event, "LookaheadBlocksSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
