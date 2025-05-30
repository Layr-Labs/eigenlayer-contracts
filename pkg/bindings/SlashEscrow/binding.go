// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package SlashEscrow

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

// OperatorSet is an auto generated low-level Go binding around an user-defined struct.
type OperatorSet struct {
	Avs common.Address
	Id  uint32
}

// SlashEscrowMetaData contains all meta data concerning the SlashEscrow contract.
var SlashEscrowMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"releaseTokens\",\"inputs\":[{\"name\":\"slashEscrowFactory\",\"type\":\"address\",\"internalType\":\"contractISlashEscrowFactory\"},{\"name\":\"slashEscrowImplementation\",\"type\":\"address\",\"internalType\":\"contractISlashEscrow\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyDeploymentParameters\",\"inputs\":[{\"name\":\"slashEscrowFactory\",\"type\":\"address\",\"internalType\":\"contractISlashEscrowFactory\"},{\"name\":\"slashEscrowImplementation\",\"type\":\"address\",\"internalType\":\"contractISlashEscrow\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"InvalidDeploymentParameters\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlySlashEscrowFactory\",\"inputs\":[]}]",
	Bin: "0x6080604052348015600e575f5ffd5b5061077b8061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610034575f3560e01c80630d9e4ef114610038578063ff491e651461005f575b5f5ffd5b61004b61004636600461054b565b610074565b604051901515815260200160405180910390f35b61007261006d366004610598565b61012d565b005b5f306101188561009161008c3688900388018861060c565b610269565b6040805160208101929092528101869052606001604051602081830303815290604052805190602001208860405160388101919091526f5af43d82803e903d91602b57fd5bf3ff60248201526014810192909252733d602d80600a3d3981f3363d3d373d3d3d363d73825260588201526037600c8201206078820152605560439091012090565b6001600160a01b03161490505b949350505050565b61013986868686610074565b610156576040516308429e4760e31b815260040160405180910390fd5b336001600160a01b0387161461017f57604051637dbfb55b60e11b815260040160405180910390fd5b5f816001600160a01b0316632495a5996040518163ffffffff1660e01b8152600401602060405180830381865afa1580156101bc573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906101e0919061067f565b6040516370a0823160e01b81523060048201529091506102609084906001600160a01b038416906370a0823190602401602060405180830381865afa15801561022b573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061024f91906106a1565b6001600160a01b03841691906102d2565b50505050505050565b5f815f0151826020015163ffffffff166040516020016102b492919060609290921b6bffffffffffffffffffffffff1916825260a01b6001600160a01b031916601482015260200190565b6040516020818303038152906040526102cc906106b8565b92915050565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663a9059cbb60e01b179052610324908490610329565b505050565b5f61037d826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166104019092919063ffffffff16565b905080515f148061039d57508080602001905181019061039d91906106db565b6103245760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b60648201526084015b60405180910390fd5b606061012584845f85855f5f866001600160a01b0316858760405161042691906106fa565b5f6040518083038185875af1925050503d805f8114610460576040519150601f19603f3d011682016040523d82523d5f602084013e610465565b606091505b509150915061047687838387610481565b979650505050505050565b606083156104ef5782515f036104e8576001600160a01b0385163b6104e85760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016103f8565b5081610125565b61012583838151156105045781518083602001fd5b8060405162461bcd60e51b81526004016103f89190610710565b6001600160a01b0381168114610532575f5ffd5b50565b5f60408284031215610545575f5ffd5b50919050565b5f5f5f5f60a0858703121561055e575f5ffd5b84356105698161051e565b935060208501356105798161051e565b92506105888660408701610535565b9396929550929360800135925050565b5f5f5f5f5f5f60e087890312156105ad575f5ffd5b86356105b88161051e565b955060208701356105c88161051e565b94506105d78860408901610535565b93506080870135925060a08701356105ee8161051e565b915060c08701356105fe8161051e565b809150509295509295509295565b5f604082840312801561061d575f5ffd5b506040805190810167ffffffffffffffff8111828210171561064d57634e487b7160e01b5f52604160045260245ffd5b604052823561065b8161051e565b8152602083013563ffffffff81168114610673575f5ffd5b60208201529392505050565b5f6020828403121561068f575f5ffd5b815161069a8161051e565b9392505050565b5f602082840312156106b1575f5ffd5b5051919050565b80516020808301519190811015610545575f1960209190910360031b1b16919050565b5f602082840312156106eb575f5ffd5b8151801515811461069a575f5ffd5b5f82518060208501845e5f920191825250919050565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f8301168401019150509291505056fea26469706673582212205d2269668ed91fcb8ef86a7c07c9513e5dadc6b928f43dc2374c6b4ae1f09f9f64736f6c634300081b0033",
}

// SlashEscrowABI is the input ABI used to generate the binding from.
// Deprecated: Use SlashEscrowMetaData.ABI instead.
var SlashEscrowABI = SlashEscrowMetaData.ABI

// SlashEscrowBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SlashEscrowMetaData.Bin instead.
var SlashEscrowBin = SlashEscrowMetaData.Bin

// DeploySlashEscrow deploys a new Ethereum contract, binding an instance of SlashEscrow to it.
func DeploySlashEscrow(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SlashEscrow, error) {
	parsed, err := SlashEscrowMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SlashEscrowBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SlashEscrow{SlashEscrowCaller: SlashEscrowCaller{contract: contract}, SlashEscrowTransactor: SlashEscrowTransactor{contract: contract}, SlashEscrowFilterer: SlashEscrowFilterer{contract: contract}}, nil
}

// SlashEscrow is an auto generated Go binding around an Ethereum contract.
type SlashEscrow struct {
	SlashEscrowCaller     // Read-only binding to the contract
	SlashEscrowTransactor // Write-only binding to the contract
	SlashEscrowFilterer   // Log filterer for contract events
}

// SlashEscrowCaller is an auto generated read-only Go binding around an Ethereum contract.
type SlashEscrowCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SlashEscrowTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SlashEscrowTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SlashEscrowFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SlashEscrowFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SlashEscrowSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SlashEscrowSession struct {
	Contract     *SlashEscrow      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SlashEscrowCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SlashEscrowCallerSession struct {
	Contract *SlashEscrowCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// SlashEscrowTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SlashEscrowTransactorSession struct {
	Contract     *SlashEscrowTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SlashEscrowRaw is an auto generated low-level Go binding around an Ethereum contract.
type SlashEscrowRaw struct {
	Contract *SlashEscrow // Generic contract binding to access the raw methods on
}

// SlashEscrowCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SlashEscrowCallerRaw struct {
	Contract *SlashEscrowCaller // Generic read-only contract binding to access the raw methods on
}

// SlashEscrowTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SlashEscrowTransactorRaw struct {
	Contract *SlashEscrowTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSlashEscrow creates a new instance of SlashEscrow, bound to a specific deployed contract.
func NewSlashEscrow(address common.Address, backend bind.ContractBackend) (*SlashEscrow, error) {
	contract, err := bindSlashEscrow(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SlashEscrow{SlashEscrowCaller: SlashEscrowCaller{contract: contract}, SlashEscrowTransactor: SlashEscrowTransactor{contract: contract}, SlashEscrowFilterer: SlashEscrowFilterer{contract: contract}}, nil
}

// NewSlashEscrowCaller creates a new read-only instance of SlashEscrow, bound to a specific deployed contract.
func NewSlashEscrowCaller(address common.Address, caller bind.ContractCaller) (*SlashEscrowCaller, error) {
	contract, err := bindSlashEscrow(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SlashEscrowCaller{contract: contract}, nil
}

// NewSlashEscrowTransactor creates a new write-only instance of SlashEscrow, bound to a specific deployed contract.
func NewSlashEscrowTransactor(address common.Address, transactor bind.ContractTransactor) (*SlashEscrowTransactor, error) {
	contract, err := bindSlashEscrow(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SlashEscrowTransactor{contract: contract}, nil
}

// NewSlashEscrowFilterer creates a new log filterer instance of SlashEscrow, bound to a specific deployed contract.
func NewSlashEscrowFilterer(address common.Address, filterer bind.ContractFilterer) (*SlashEscrowFilterer, error) {
	contract, err := bindSlashEscrow(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SlashEscrowFilterer{contract: contract}, nil
}

// bindSlashEscrow binds a generic wrapper to an already deployed contract.
func bindSlashEscrow(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SlashEscrowMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SlashEscrow *SlashEscrowRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SlashEscrow.Contract.SlashEscrowCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SlashEscrow *SlashEscrowRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SlashEscrow.Contract.SlashEscrowTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SlashEscrow *SlashEscrowRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SlashEscrow.Contract.SlashEscrowTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SlashEscrow *SlashEscrowCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SlashEscrow.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SlashEscrow *SlashEscrowTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SlashEscrow.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SlashEscrow *SlashEscrowTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SlashEscrow.Contract.contract.Transact(opts, method, params...)
}

// VerifyDeploymentParameters is a free data retrieval call binding the contract method 0x0d9e4ef1.
//
// Solidity: function verifyDeploymentParameters(address slashEscrowFactory, address slashEscrowImplementation, (address,uint32) operatorSet, uint256 slashId) view returns(bool)
func (_SlashEscrow *SlashEscrowCaller) VerifyDeploymentParameters(opts *bind.CallOpts, slashEscrowFactory common.Address, slashEscrowImplementation common.Address, operatorSet OperatorSet, slashId *big.Int) (bool, error) {
	var out []interface{}
	err := _SlashEscrow.contract.Call(opts, &out, "verifyDeploymentParameters", slashEscrowFactory, slashEscrowImplementation, operatorSet, slashId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyDeploymentParameters is a free data retrieval call binding the contract method 0x0d9e4ef1.
//
// Solidity: function verifyDeploymentParameters(address slashEscrowFactory, address slashEscrowImplementation, (address,uint32) operatorSet, uint256 slashId) view returns(bool)
func (_SlashEscrow *SlashEscrowSession) VerifyDeploymentParameters(slashEscrowFactory common.Address, slashEscrowImplementation common.Address, operatorSet OperatorSet, slashId *big.Int) (bool, error) {
	return _SlashEscrow.Contract.VerifyDeploymentParameters(&_SlashEscrow.CallOpts, slashEscrowFactory, slashEscrowImplementation, operatorSet, slashId)
}

// VerifyDeploymentParameters is a free data retrieval call binding the contract method 0x0d9e4ef1.
//
// Solidity: function verifyDeploymentParameters(address slashEscrowFactory, address slashEscrowImplementation, (address,uint32) operatorSet, uint256 slashId) view returns(bool)
func (_SlashEscrow *SlashEscrowCallerSession) VerifyDeploymentParameters(slashEscrowFactory common.Address, slashEscrowImplementation common.Address, operatorSet OperatorSet, slashId *big.Int) (bool, error) {
	return _SlashEscrow.Contract.VerifyDeploymentParameters(&_SlashEscrow.CallOpts, slashEscrowFactory, slashEscrowImplementation, operatorSet, slashId)
}

// ReleaseTokens is a paid mutator transaction binding the contract method 0xff491e65.
//
// Solidity: function releaseTokens(address slashEscrowFactory, address slashEscrowImplementation, (address,uint32) operatorSet, uint256 slashId, address recipient, address strategy) returns()
func (_SlashEscrow *SlashEscrowTransactor) ReleaseTokens(opts *bind.TransactOpts, slashEscrowFactory common.Address, slashEscrowImplementation common.Address, operatorSet OperatorSet, slashId *big.Int, recipient common.Address, strategy common.Address) (*types.Transaction, error) {
	return _SlashEscrow.contract.Transact(opts, "releaseTokens", slashEscrowFactory, slashEscrowImplementation, operatorSet, slashId, recipient, strategy)
}

// ReleaseTokens is a paid mutator transaction binding the contract method 0xff491e65.
//
// Solidity: function releaseTokens(address slashEscrowFactory, address slashEscrowImplementation, (address,uint32) operatorSet, uint256 slashId, address recipient, address strategy) returns()
func (_SlashEscrow *SlashEscrowSession) ReleaseTokens(slashEscrowFactory common.Address, slashEscrowImplementation common.Address, operatorSet OperatorSet, slashId *big.Int, recipient common.Address, strategy common.Address) (*types.Transaction, error) {
	return _SlashEscrow.Contract.ReleaseTokens(&_SlashEscrow.TransactOpts, slashEscrowFactory, slashEscrowImplementation, operatorSet, slashId, recipient, strategy)
}

// ReleaseTokens is a paid mutator transaction binding the contract method 0xff491e65.
//
// Solidity: function releaseTokens(address slashEscrowFactory, address slashEscrowImplementation, (address,uint32) operatorSet, uint256 slashId, address recipient, address strategy) returns()
func (_SlashEscrow *SlashEscrowTransactorSession) ReleaseTokens(slashEscrowFactory common.Address, slashEscrowImplementation common.Address, operatorSet OperatorSet, slashId *big.Int, recipient common.Address, strategy common.Address) (*types.Transaction, error) {
	return _SlashEscrow.Contract.ReleaseTokens(&_SlashEscrow.TransactOpts, slashEscrowFactory, slashEscrowImplementation, operatorSet, slashId, recipient, strategy)
}
