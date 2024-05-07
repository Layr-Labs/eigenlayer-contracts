// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IPaymentCoordinator

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

// IPaymentCoordinatorEarnerTreeMerkleLeaf is an auto generated low-level Go binding around an user-defined struct.
type IPaymentCoordinatorEarnerTreeMerkleLeaf struct {
	Earner          common.Address
	EarnerTokenRoot [32]byte
}

// IPaymentCoordinatorPaymentMerkleClaim is an auto generated low-level Go binding around an user-defined struct.
type IPaymentCoordinatorPaymentMerkleClaim struct {
	RootIndex       uint32
	EarnerIndex     uint32
	EarnerTreeProof []byte
	EarnerLeaf      IPaymentCoordinatorEarnerTreeMerkleLeaf
	TokenIndices    []uint32
	TokenTreeProofs [][]byte
	TokenLeaves     []IPaymentCoordinatorTokenTreeMerkleLeaf
}

// IPaymentCoordinatorRangePayment is an auto generated low-level Go binding around an user-defined struct.
type IPaymentCoordinatorRangePayment struct {
	StrategiesAndMultipliers []IPaymentCoordinatorStrategyAndMultiplier
	Token                    common.Address
	Amount                   *big.Int
	StartTimestamp           uint32
	Duration                 uint32
}

// IPaymentCoordinatorStrategyAndMultiplier is an auto generated low-level Go binding around an user-defined struct.
type IPaymentCoordinatorStrategyAndMultiplier struct {
	Strategy   common.Address
	Multiplier *big.Int
}

// IPaymentCoordinatorTokenTreeMerkleLeaf is an auto generated low-level Go binding around an user-defined struct.
type IPaymentCoordinatorTokenTreeMerkleLeaf struct {
	Token              common.Address
	CumulativeEarnings *big.Int
}

// IPaymentCoordinatorMetaData contains all meta data concerning the IPaymentCoordinator contract.
var IPaymentCoordinatorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"CALCULATION_INTERVAL_SECONDS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GENESIS_PAYMENT_TIMESTAMP\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_FUTURE_LENGTH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_PAYMENT_DURATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_RETROACTIVE_LENGTH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"activationDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateEarnerLeafHash\",\"inputs\":[{\"name\":\"leaf\",\"type\":\"tuple\",\"internalType\":\"structIPaymentCoordinator.EarnerTreeMerkleLeaf\",\"components\":[{\"name\":\"earner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"earnerTokenRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"calculateTokenLeafHash\",\"inputs\":[{\"name\":\"leaf\",\"type\":\"tuple\",\"internalType\":\"structIPaymentCoordinator.TokenTreeMerkleLeaf\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"cumulativeEarnings\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"checkClaim\",\"inputs\":[{\"name\":\"claim\",\"type\":\"tuple\",\"internalType\":\"structIPaymentCoordinator.PaymentMerkleClaim\",\"components\":[{\"name\":\"rootIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"earnerIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"earnerTreeProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"earnerLeaf\",\"type\":\"tuple\",\"internalType\":\"structIPaymentCoordinator.EarnerTreeMerkleLeaf\",\"components\":[{\"name\":\"earner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"earnerTokenRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"tokenIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"tokenTreeProofs\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"tokenLeaves\",\"type\":\"tuple[]\",\"internalType\":\"structIPaymentCoordinator.TokenTreeMerkleLeaf[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"cumulativeEarnings\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimerFor\",\"inputs\":[{\"name\":\"earner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cumulativeClaimed\",\"inputs\":[{\"name\":\"claimer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"currPaymentCalculationEndTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRootIndexFromHash\",\"inputs\":[{\"name\":\"rootHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"globalOperatorCommissionBips\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorCommissionBips\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"payAllForRange\",\"inputs\":[{\"name\":\"rangePayment\",\"type\":\"tuple[]\",\"internalType\":\"structIPaymentCoordinator.RangePayment[]\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIPaymentCoordinator.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"payForRange\",\"inputs\":[{\"name\":\"rangePayments\",\"type\":\"tuple[]\",\"internalType\":\"structIPaymentCoordinator.RangePayment[]\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIPaymentCoordinator.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paymentUpdater\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"processClaim\",\"inputs\":[{\"name\":\"claim\",\"type\":\"tuple\",\"internalType\":\"structIPaymentCoordinator.PaymentMerkleClaim\",\"components\":[{\"name\":\"rootIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"earnerIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"earnerTreeProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"earnerLeaf\",\"type\":\"tuple\",\"internalType\":\"structIPaymentCoordinator.EarnerTreeMerkleLeaf\",\"components\":[{\"name\":\"earner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"earnerTokenRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"tokenIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"tokenTreeProofs\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"tokenLeaves\",\"type\":\"tuple[]\",\"internalType\":\"structIPaymentCoordinator.TokenTreeMerkleLeaf[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"cumulativeEarnings\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setActivationDelay\",\"inputs\":[{\"name\":\"_activationDelay\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setClaimerFor\",\"inputs\":[{\"name\":\"claimer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGlobalOperatorCommission\",\"inputs\":[{\"name\":\"_globalCommissionBips\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPayAllForRangeSubmitter\",\"inputs\":[{\"name\":\"_submitter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_newValue\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPaymentUpdater\",\"inputs\":[{\"name\":\"_paymentUpdater\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitRoot\",\"inputs\":[{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"paymentCalculationEndTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ActivationDelaySet\",\"inputs\":[{\"name\":\"oldActivationDelay\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"newActivationDelay\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ClaimerForSet\",\"inputs\":[{\"name\":\"earner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"oldClaimer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"claimer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DistributionRootSubmitted\",\"inputs\":[{\"name\":\"rootIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"root\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"paymentCalculationEndTimestamp\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"activatedAt\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GlobalCommissionBipsSet\",\"inputs\":[{\"name\":\"oldGlobalCommissionBips\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"newGlobalCommissionBips\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PayAllForRangeSubmitterSet\",\"inputs\":[{\"name\":\"payAllForRangeSubmitter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"oldValue\",\"type\":\"bool\",\"indexed\":true,\"internalType\":\"bool\"},{\"name\":\"newValue\",\"type\":\"bool\",\"indexed\":true,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PaymentClaimed\",\"inputs\":[{\"name\":\"root\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"earner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"claimer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIERC20\"},{\"name\":\"claimedAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PaymentUpdaterSet\",\"inputs\":[{\"name\":\"oldPaymentUpdater\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPaymentUpdater\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RangePaymentCreated\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"paymentNonce\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"rangePaymentHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"rangePayment\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIPaymentCoordinator.RangePayment\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIPaymentCoordinator.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RangePaymentForAllCreated\",\"inputs\":[{\"name\":\"submitter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"paymentNonce\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"rangePaymentHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"rangePayment\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIPaymentCoordinator.RangePayment\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIPaymentCoordinator.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false}]",
}

// IPaymentCoordinatorABI is the input ABI used to generate the binding from.
// Deprecated: Use IPaymentCoordinatorMetaData.ABI instead.
var IPaymentCoordinatorABI = IPaymentCoordinatorMetaData.ABI

// IPaymentCoordinator is an auto generated Go binding around an Ethereum contract.
type IPaymentCoordinator struct {
	IPaymentCoordinatorCaller     // Read-only binding to the contract
	IPaymentCoordinatorTransactor // Write-only binding to the contract
	IPaymentCoordinatorFilterer   // Log filterer for contract events
}

// IPaymentCoordinatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPaymentCoordinatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPaymentCoordinatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPaymentCoordinatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPaymentCoordinatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPaymentCoordinatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPaymentCoordinatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPaymentCoordinatorSession struct {
	Contract     *IPaymentCoordinator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IPaymentCoordinatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPaymentCoordinatorCallerSession struct {
	Contract *IPaymentCoordinatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// IPaymentCoordinatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPaymentCoordinatorTransactorSession struct {
	Contract     *IPaymentCoordinatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IPaymentCoordinatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPaymentCoordinatorRaw struct {
	Contract *IPaymentCoordinator // Generic contract binding to access the raw methods on
}

// IPaymentCoordinatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPaymentCoordinatorCallerRaw struct {
	Contract *IPaymentCoordinatorCaller // Generic read-only contract binding to access the raw methods on
}

// IPaymentCoordinatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPaymentCoordinatorTransactorRaw struct {
	Contract *IPaymentCoordinatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPaymentCoordinator creates a new instance of IPaymentCoordinator, bound to a specific deployed contract.
func NewIPaymentCoordinator(address common.Address, backend bind.ContractBackend) (*IPaymentCoordinator, error) {
	contract, err := bindIPaymentCoordinator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPaymentCoordinator{IPaymentCoordinatorCaller: IPaymentCoordinatorCaller{contract: contract}, IPaymentCoordinatorTransactor: IPaymentCoordinatorTransactor{contract: contract}, IPaymentCoordinatorFilterer: IPaymentCoordinatorFilterer{contract: contract}}, nil
}

// NewIPaymentCoordinatorCaller creates a new read-only instance of IPaymentCoordinator, bound to a specific deployed contract.
func NewIPaymentCoordinatorCaller(address common.Address, caller bind.ContractCaller) (*IPaymentCoordinatorCaller, error) {
	contract, err := bindIPaymentCoordinator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPaymentCoordinatorCaller{contract: contract}, nil
}

// NewIPaymentCoordinatorTransactor creates a new write-only instance of IPaymentCoordinator, bound to a specific deployed contract.
func NewIPaymentCoordinatorTransactor(address common.Address, transactor bind.ContractTransactor) (*IPaymentCoordinatorTransactor, error) {
	contract, err := bindIPaymentCoordinator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPaymentCoordinatorTransactor{contract: contract}, nil
}

// NewIPaymentCoordinatorFilterer creates a new log filterer instance of IPaymentCoordinator, bound to a specific deployed contract.
func NewIPaymentCoordinatorFilterer(address common.Address, filterer bind.ContractFilterer) (*IPaymentCoordinatorFilterer, error) {
	contract, err := bindIPaymentCoordinator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPaymentCoordinatorFilterer{contract: contract}, nil
}

// bindIPaymentCoordinator binds a generic wrapper to an already deployed contract.
func bindIPaymentCoordinator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IPaymentCoordinatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPaymentCoordinator *IPaymentCoordinatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPaymentCoordinator.Contract.IPaymentCoordinatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPaymentCoordinator *IPaymentCoordinatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPaymentCoordinator.Contract.IPaymentCoordinatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPaymentCoordinator *IPaymentCoordinatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPaymentCoordinator.Contract.IPaymentCoordinatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPaymentCoordinator *IPaymentCoordinatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPaymentCoordinator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPaymentCoordinator *IPaymentCoordinatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPaymentCoordinator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPaymentCoordinator *IPaymentCoordinatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPaymentCoordinator.Contract.contract.Transact(opts, method, params...)
}

// CALCULATIONINTERVALSECONDS is a free data retrieval call binding the contract method 0x9d45c281.
//
// Solidity: function CALCULATION_INTERVAL_SECONDS() view returns(uint32)
func (_IPaymentCoordinator *IPaymentCoordinatorCaller) CALCULATIONINTERVALSECONDS(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _IPaymentCoordinator.contract.Call(opts, &out, "CALCULATION_INTERVAL_SECONDS")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// CALCULATIONINTERVALSECONDS is a free data retrieval call binding the contract method 0x9d45c281.
//
// Solidity: function CALCULATION_INTERVAL_SECONDS() view returns(uint32)
func (_IPaymentCoordinator *IPaymentCoordinatorSession) CALCULATIONINTERVALSECONDS() (uint32, error) {
	return _IPaymentCoordinator.Contract.CALCULATIONINTERVALSECONDS(&_IPaymentCoordinator.CallOpts)
}

// CALCULATIONINTERVALSECONDS is a free data retrieval call binding the contract method 0x9d45c281.
//
// Solidity: function CALCULATION_INTERVAL_SECONDS() view returns(uint32)
func (_IPaymentCoordinator *IPaymentCoordinatorCallerSession) CALCULATIONINTERVALSECONDS() (uint32, error) {
	return _IPaymentCoordinator.Contract.CALCULATIONINTERVALSECONDS(&_IPaymentCoordinator.CallOpts)
}

// GENESISPAYMENTTIMESTAMP is a free data retrieval call binding the contract method 0x2cfd45eb.
//
// Solidity: function GENESIS_PAYMENT_TIMESTAMP() view returns(uint32)
func (_IPaymentCoordinator *IPaymentCoordinatorCaller) GENESISPAYMENTTIMESTAMP(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _IPaymentCoordinator.contract.Call(opts, &out, "GENESIS_PAYMENT_TIMESTAMP")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GENESISPAYMENTTIMESTAMP is a free data retrieval call binding the contract method 0x2cfd45eb.
//
// Solidity: function GENESIS_PAYMENT_TIMESTAMP() view returns(uint32)
func (_IPaymentCoordinator *IPaymentCoordinatorSession) GENESISPAYMENTTIMESTAMP() (uint32, error) {
	return _IPaymentCoordinator.Contract.GENESISPAYMENTTIMESTAMP(&_IPaymentCoordinator.CallOpts)
}

// GENESISPAYMENTTIMESTAMP is a free data retrieval call binding the contract method 0x2cfd45eb.
//
// Solidity: function GENESIS_PAYMENT_TIMESTAMP() view returns(uint32)
func (_IPaymentCoordinator *IPaymentCoordinatorCallerSession) GENESISPAYMENTTIMESTAMP() (uint32, error) {
	return _IPaymentCoordinator.Contract.GENESISPAYMENTTIMESTAMP(&_IPaymentCoordinator.CallOpts)
}

// MAXFUTURELENGTH is a free data retrieval call binding the contract method 0x04a0c502.
//
// Solidity: function MAX_FUTURE_LENGTH() view returns(uint32)
func (_IPaymentCoordinator *IPaymentCoordinatorCaller) MAXFUTURELENGTH(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _IPaymentCoordinator.contract.Call(opts, &out, "MAX_FUTURE_LENGTH")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MAXFUTURELENGTH is a free data retrieval call binding the contract method 0x04a0c502.
//
// Solidity: function MAX_FUTURE_LENGTH() view returns(uint32)
func (_IPaymentCoordinator *IPaymentCoordinatorSession) MAXFUTURELENGTH() (uint32, error) {
	return _IPaymentCoordinator.Contract.MAXFUTURELENGTH(&_IPaymentCoordinator.CallOpts)
}

// MAXFUTURELENGTH is a free data retrieval call binding the contract method 0x04a0c502.
//
// Solidity: function MAX_FUTURE_LENGTH() view returns(uint32)
func (_IPaymentCoordinator *IPaymentCoordinatorCallerSession) MAXFUTURELENGTH() (uint32, error) {
	return _IPaymentCoordinator.Contract.MAXFUTURELENGTH(&_IPaymentCoordinator.CallOpts)
}

// MAXPAYMENTDURATION is a free data retrieval call binding the contract method 0xee619597.
//
// Solidity: function MAX_PAYMENT_DURATION() view returns(uint32)
func (_IPaymentCoordinator *IPaymentCoordinatorCaller) MAXPAYMENTDURATION(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _IPaymentCoordinator.contract.Call(opts, &out, "MAX_PAYMENT_DURATION")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MAXPAYMENTDURATION is a free data retrieval call binding the contract method 0xee619597.
//
// Solidity: function MAX_PAYMENT_DURATION() view returns(uint32)
func (_IPaymentCoordinator *IPaymentCoordinatorSession) MAXPAYMENTDURATION() (uint32, error) {
	return _IPaymentCoordinator.Contract.MAXPAYMENTDURATION(&_IPaymentCoordinator.CallOpts)
}

// MAXPAYMENTDURATION is a free data retrieval call binding the contract method 0xee619597.
//
// Solidity: function MAX_PAYMENT_DURATION() view returns(uint32)
func (_IPaymentCoordinator *IPaymentCoordinatorCallerSession) MAXPAYMENTDURATION() (uint32, error) {
	return _IPaymentCoordinator.Contract.MAXPAYMENTDURATION(&_IPaymentCoordinator.CallOpts)
}

// MAXRETROACTIVELENGTH is a free data retrieval call binding the contract method 0x37838ed0.
//
// Solidity: function MAX_RETROACTIVE_LENGTH() view returns(uint32)
func (_IPaymentCoordinator *IPaymentCoordinatorCaller) MAXRETROACTIVELENGTH(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _IPaymentCoordinator.contract.Call(opts, &out, "MAX_RETROACTIVE_LENGTH")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MAXRETROACTIVELENGTH is a free data retrieval call binding the contract method 0x37838ed0.
//
// Solidity: function MAX_RETROACTIVE_LENGTH() view returns(uint32)
func (_IPaymentCoordinator *IPaymentCoordinatorSession) MAXRETROACTIVELENGTH() (uint32, error) {
	return _IPaymentCoordinator.Contract.MAXRETROACTIVELENGTH(&_IPaymentCoordinator.CallOpts)
}

// MAXRETROACTIVELENGTH is a free data retrieval call binding the contract method 0x37838ed0.
//
// Solidity: function MAX_RETROACTIVE_LENGTH() view returns(uint32)
func (_IPaymentCoordinator *IPaymentCoordinatorCallerSession) MAXRETROACTIVELENGTH() (uint32, error) {
	return _IPaymentCoordinator.Contract.MAXRETROACTIVELENGTH(&_IPaymentCoordinator.CallOpts)
}

// ActivationDelay is a free data retrieval call binding the contract method 0x3a8c0786.
//
// Solidity: function activationDelay() view returns(uint32)
func (_IPaymentCoordinator *IPaymentCoordinatorCaller) ActivationDelay(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _IPaymentCoordinator.contract.Call(opts, &out, "activationDelay")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ActivationDelay is a free data retrieval call binding the contract method 0x3a8c0786.
//
// Solidity: function activationDelay() view returns(uint32)
func (_IPaymentCoordinator *IPaymentCoordinatorSession) ActivationDelay() (uint32, error) {
	return _IPaymentCoordinator.Contract.ActivationDelay(&_IPaymentCoordinator.CallOpts)
}

// ActivationDelay is a free data retrieval call binding the contract method 0x3a8c0786.
//
// Solidity: function activationDelay() view returns(uint32)
func (_IPaymentCoordinator *IPaymentCoordinatorCallerSession) ActivationDelay() (uint32, error) {
	return _IPaymentCoordinator.Contract.ActivationDelay(&_IPaymentCoordinator.CallOpts)
}

// CalculateEarnerLeafHash is a free data retrieval call binding the contract method 0x149bc872.
//
// Solidity: function calculateEarnerLeafHash((address,bytes32) leaf) pure returns(bytes32)
func (_IPaymentCoordinator *IPaymentCoordinatorCaller) CalculateEarnerLeafHash(opts *bind.CallOpts, leaf IPaymentCoordinatorEarnerTreeMerkleLeaf) ([32]byte, error) {
	var out []interface{}
	err := _IPaymentCoordinator.contract.Call(opts, &out, "calculateEarnerLeafHash", leaf)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateEarnerLeafHash is a free data retrieval call binding the contract method 0x149bc872.
//
// Solidity: function calculateEarnerLeafHash((address,bytes32) leaf) pure returns(bytes32)
func (_IPaymentCoordinator *IPaymentCoordinatorSession) CalculateEarnerLeafHash(leaf IPaymentCoordinatorEarnerTreeMerkleLeaf) ([32]byte, error) {
	return _IPaymentCoordinator.Contract.CalculateEarnerLeafHash(&_IPaymentCoordinator.CallOpts, leaf)
}

// CalculateEarnerLeafHash is a free data retrieval call binding the contract method 0x149bc872.
//
// Solidity: function calculateEarnerLeafHash((address,bytes32) leaf) pure returns(bytes32)
func (_IPaymentCoordinator *IPaymentCoordinatorCallerSession) CalculateEarnerLeafHash(leaf IPaymentCoordinatorEarnerTreeMerkleLeaf) ([32]byte, error) {
	return _IPaymentCoordinator.Contract.CalculateEarnerLeafHash(&_IPaymentCoordinator.CallOpts, leaf)
}

// CalculateTokenLeafHash is a free data retrieval call binding the contract method 0xf8cd8448.
//
// Solidity: function calculateTokenLeafHash((address,uint256) leaf) pure returns(bytes32)
func (_IPaymentCoordinator *IPaymentCoordinatorCaller) CalculateTokenLeafHash(opts *bind.CallOpts, leaf IPaymentCoordinatorTokenTreeMerkleLeaf) ([32]byte, error) {
	var out []interface{}
	err := _IPaymentCoordinator.contract.Call(opts, &out, "calculateTokenLeafHash", leaf)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateTokenLeafHash is a free data retrieval call binding the contract method 0xf8cd8448.
//
// Solidity: function calculateTokenLeafHash((address,uint256) leaf) pure returns(bytes32)
func (_IPaymentCoordinator *IPaymentCoordinatorSession) CalculateTokenLeafHash(leaf IPaymentCoordinatorTokenTreeMerkleLeaf) ([32]byte, error) {
	return _IPaymentCoordinator.Contract.CalculateTokenLeafHash(&_IPaymentCoordinator.CallOpts, leaf)
}

// CalculateTokenLeafHash is a free data retrieval call binding the contract method 0xf8cd8448.
//
// Solidity: function calculateTokenLeafHash((address,uint256) leaf) pure returns(bytes32)
func (_IPaymentCoordinator *IPaymentCoordinatorCallerSession) CalculateTokenLeafHash(leaf IPaymentCoordinatorTokenTreeMerkleLeaf) ([32]byte, error) {
	return _IPaymentCoordinator.Contract.CalculateTokenLeafHash(&_IPaymentCoordinator.CallOpts, leaf)
}

// CheckClaim is a free data retrieval call binding the contract method 0x5e9d8348.
//
// Solidity: function checkClaim((uint32,uint32,bytes,(address,bytes32),uint32[],bytes[],(address,uint256)[]) claim) view returns(bool)
func (_IPaymentCoordinator *IPaymentCoordinatorCaller) CheckClaim(opts *bind.CallOpts, claim IPaymentCoordinatorPaymentMerkleClaim) (bool, error) {
	var out []interface{}
	err := _IPaymentCoordinator.contract.Call(opts, &out, "checkClaim", claim)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckClaim is a free data retrieval call binding the contract method 0x5e9d8348.
//
// Solidity: function checkClaim((uint32,uint32,bytes,(address,bytes32),uint32[],bytes[],(address,uint256)[]) claim) view returns(bool)
func (_IPaymentCoordinator *IPaymentCoordinatorSession) CheckClaim(claim IPaymentCoordinatorPaymentMerkleClaim) (bool, error) {
	return _IPaymentCoordinator.Contract.CheckClaim(&_IPaymentCoordinator.CallOpts, claim)
}

// CheckClaim is a free data retrieval call binding the contract method 0x5e9d8348.
//
// Solidity: function checkClaim((uint32,uint32,bytes,(address,bytes32),uint32[],bytes[],(address,uint256)[]) claim) view returns(bool)
func (_IPaymentCoordinator *IPaymentCoordinatorCallerSession) CheckClaim(claim IPaymentCoordinatorPaymentMerkleClaim) (bool, error) {
	return _IPaymentCoordinator.Contract.CheckClaim(&_IPaymentCoordinator.CallOpts, claim)
}

// ClaimerFor is a free data retrieval call binding the contract method 0x2b9f64a4.
//
// Solidity: function claimerFor(address earner) view returns(address)
func (_IPaymentCoordinator *IPaymentCoordinatorCaller) ClaimerFor(opts *bind.CallOpts, earner common.Address) (common.Address, error) {
	var out []interface{}
	err := _IPaymentCoordinator.contract.Call(opts, &out, "claimerFor", earner)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ClaimerFor is a free data retrieval call binding the contract method 0x2b9f64a4.
//
// Solidity: function claimerFor(address earner) view returns(address)
func (_IPaymentCoordinator *IPaymentCoordinatorSession) ClaimerFor(earner common.Address) (common.Address, error) {
	return _IPaymentCoordinator.Contract.ClaimerFor(&_IPaymentCoordinator.CallOpts, earner)
}

// ClaimerFor is a free data retrieval call binding the contract method 0x2b9f64a4.
//
// Solidity: function claimerFor(address earner) view returns(address)
func (_IPaymentCoordinator *IPaymentCoordinatorCallerSession) ClaimerFor(earner common.Address) (common.Address, error) {
	return _IPaymentCoordinator.Contract.ClaimerFor(&_IPaymentCoordinator.CallOpts, earner)
}

// CumulativeClaimed is a free data retrieval call binding the contract method 0x865c6953.
//
// Solidity: function cumulativeClaimed(address claimer, address token) view returns(uint256)
func (_IPaymentCoordinator *IPaymentCoordinatorCaller) CumulativeClaimed(opts *bind.CallOpts, claimer common.Address, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IPaymentCoordinator.contract.Call(opts, &out, "cumulativeClaimed", claimer, token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CumulativeClaimed is a free data retrieval call binding the contract method 0x865c6953.
//
// Solidity: function cumulativeClaimed(address claimer, address token) view returns(uint256)
func (_IPaymentCoordinator *IPaymentCoordinatorSession) CumulativeClaimed(claimer common.Address, token common.Address) (*big.Int, error) {
	return _IPaymentCoordinator.Contract.CumulativeClaimed(&_IPaymentCoordinator.CallOpts, claimer, token)
}

// CumulativeClaimed is a free data retrieval call binding the contract method 0x865c6953.
//
// Solidity: function cumulativeClaimed(address claimer, address token) view returns(uint256)
func (_IPaymentCoordinator *IPaymentCoordinatorCallerSession) CumulativeClaimed(claimer common.Address, token common.Address) (*big.Int, error) {
	return _IPaymentCoordinator.Contract.CumulativeClaimed(&_IPaymentCoordinator.CallOpts, claimer, token)
}

// CurrPaymentCalculationEndTimestamp is a free data retrieval call binding the contract method 0x67ef8585.
//
// Solidity: function currPaymentCalculationEndTimestamp() view returns(uint32)
func (_IPaymentCoordinator *IPaymentCoordinatorCaller) CurrPaymentCalculationEndTimestamp(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _IPaymentCoordinator.contract.Call(opts, &out, "currPaymentCalculationEndTimestamp")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// CurrPaymentCalculationEndTimestamp is a free data retrieval call binding the contract method 0x67ef8585.
//
// Solidity: function currPaymentCalculationEndTimestamp() view returns(uint32)
func (_IPaymentCoordinator *IPaymentCoordinatorSession) CurrPaymentCalculationEndTimestamp() (uint32, error) {
	return _IPaymentCoordinator.Contract.CurrPaymentCalculationEndTimestamp(&_IPaymentCoordinator.CallOpts)
}

// CurrPaymentCalculationEndTimestamp is a free data retrieval call binding the contract method 0x67ef8585.
//
// Solidity: function currPaymentCalculationEndTimestamp() view returns(uint32)
func (_IPaymentCoordinator *IPaymentCoordinatorCallerSession) CurrPaymentCalculationEndTimestamp() (uint32, error) {
	return _IPaymentCoordinator.Contract.CurrPaymentCalculationEndTimestamp(&_IPaymentCoordinator.CallOpts)
}

// GetRootIndexFromHash is a free data retrieval call binding the contract method 0xe810ce21.
//
// Solidity: function getRootIndexFromHash(bytes32 rootHash) view returns(uint32)
func (_IPaymentCoordinator *IPaymentCoordinatorCaller) GetRootIndexFromHash(opts *bind.CallOpts, rootHash [32]byte) (uint32, error) {
	var out []interface{}
	err := _IPaymentCoordinator.contract.Call(opts, &out, "getRootIndexFromHash", rootHash)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetRootIndexFromHash is a free data retrieval call binding the contract method 0xe810ce21.
//
// Solidity: function getRootIndexFromHash(bytes32 rootHash) view returns(uint32)
func (_IPaymentCoordinator *IPaymentCoordinatorSession) GetRootIndexFromHash(rootHash [32]byte) (uint32, error) {
	return _IPaymentCoordinator.Contract.GetRootIndexFromHash(&_IPaymentCoordinator.CallOpts, rootHash)
}

// GetRootIndexFromHash is a free data retrieval call binding the contract method 0xe810ce21.
//
// Solidity: function getRootIndexFromHash(bytes32 rootHash) view returns(uint32)
func (_IPaymentCoordinator *IPaymentCoordinatorCallerSession) GetRootIndexFromHash(rootHash [32]byte) (uint32, error) {
	return _IPaymentCoordinator.Contract.GetRootIndexFromHash(&_IPaymentCoordinator.CallOpts, rootHash)
}

// GlobalOperatorCommissionBips is a free data retrieval call binding the contract method 0x092db007.
//
// Solidity: function globalOperatorCommissionBips() view returns(uint16)
func (_IPaymentCoordinator *IPaymentCoordinatorCaller) GlobalOperatorCommissionBips(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _IPaymentCoordinator.contract.Call(opts, &out, "globalOperatorCommissionBips")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GlobalOperatorCommissionBips is a free data retrieval call binding the contract method 0x092db007.
//
// Solidity: function globalOperatorCommissionBips() view returns(uint16)
func (_IPaymentCoordinator *IPaymentCoordinatorSession) GlobalOperatorCommissionBips() (uint16, error) {
	return _IPaymentCoordinator.Contract.GlobalOperatorCommissionBips(&_IPaymentCoordinator.CallOpts)
}

// GlobalOperatorCommissionBips is a free data retrieval call binding the contract method 0x092db007.
//
// Solidity: function globalOperatorCommissionBips() view returns(uint16)
func (_IPaymentCoordinator *IPaymentCoordinatorCallerSession) GlobalOperatorCommissionBips() (uint16, error) {
	return _IPaymentCoordinator.Contract.GlobalOperatorCommissionBips(&_IPaymentCoordinator.CallOpts)
}

// OperatorCommissionBips is a free data retrieval call binding the contract method 0x22f19a64.
//
// Solidity: function operatorCommissionBips(address operator, address avs) view returns(uint16)
func (_IPaymentCoordinator *IPaymentCoordinatorCaller) OperatorCommissionBips(opts *bind.CallOpts, operator common.Address, avs common.Address) (uint16, error) {
	var out []interface{}
	err := _IPaymentCoordinator.contract.Call(opts, &out, "operatorCommissionBips", operator, avs)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// OperatorCommissionBips is a free data retrieval call binding the contract method 0x22f19a64.
//
// Solidity: function operatorCommissionBips(address operator, address avs) view returns(uint16)
func (_IPaymentCoordinator *IPaymentCoordinatorSession) OperatorCommissionBips(operator common.Address, avs common.Address) (uint16, error) {
	return _IPaymentCoordinator.Contract.OperatorCommissionBips(&_IPaymentCoordinator.CallOpts, operator, avs)
}

// OperatorCommissionBips is a free data retrieval call binding the contract method 0x22f19a64.
//
// Solidity: function operatorCommissionBips(address operator, address avs) view returns(uint16)
func (_IPaymentCoordinator *IPaymentCoordinatorCallerSession) OperatorCommissionBips(operator common.Address, avs common.Address) (uint16, error) {
	return _IPaymentCoordinator.Contract.OperatorCommissionBips(&_IPaymentCoordinator.CallOpts, operator, avs)
}

// PaymentUpdater is a free data retrieval call binding the contract method 0x66d3b16b.
//
// Solidity: function paymentUpdater() view returns(address)
func (_IPaymentCoordinator *IPaymentCoordinatorCaller) PaymentUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IPaymentCoordinator.contract.Call(opts, &out, "paymentUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PaymentUpdater is a free data retrieval call binding the contract method 0x66d3b16b.
//
// Solidity: function paymentUpdater() view returns(address)
func (_IPaymentCoordinator *IPaymentCoordinatorSession) PaymentUpdater() (common.Address, error) {
	return _IPaymentCoordinator.Contract.PaymentUpdater(&_IPaymentCoordinator.CallOpts)
}

// PaymentUpdater is a free data retrieval call binding the contract method 0x66d3b16b.
//
// Solidity: function paymentUpdater() view returns(address)
func (_IPaymentCoordinator *IPaymentCoordinatorCallerSession) PaymentUpdater() (common.Address, error) {
	return _IPaymentCoordinator.Contract.PaymentUpdater(&_IPaymentCoordinator.CallOpts)
}

// PayAllForRange is a paid mutator transaction binding the contract method 0xb5715284.
//
// Solidity: function payAllForRange(((address,uint96)[],address,uint256,uint32,uint32)[] rangePayment) returns()
func (_IPaymentCoordinator *IPaymentCoordinatorTransactor) PayAllForRange(opts *bind.TransactOpts, rangePayment []IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _IPaymentCoordinator.contract.Transact(opts, "payAllForRange", rangePayment)
}

// PayAllForRange is a paid mutator transaction binding the contract method 0xb5715284.
//
// Solidity: function payAllForRange(((address,uint96)[],address,uint256,uint32,uint32)[] rangePayment) returns()
func (_IPaymentCoordinator *IPaymentCoordinatorSession) PayAllForRange(rangePayment []IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _IPaymentCoordinator.Contract.PayAllForRange(&_IPaymentCoordinator.TransactOpts, rangePayment)
}

// PayAllForRange is a paid mutator transaction binding the contract method 0xb5715284.
//
// Solidity: function payAllForRange(((address,uint96)[],address,uint256,uint32,uint32)[] rangePayment) returns()
func (_IPaymentCoordinator *IPaymentCoordinatorTransactorSession) PayAllForRange(rangePayment []IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _IPaymentCoordinator.Contract.PayAllForRange(&_IPaymentCoordinator.TransactOpts, rangePayment)
}

// PayForRange is a paid mutator transaction binding the contract method 0x1b445516.
//
// Solidity: function payForRange(((address,uint96)[],address,uint256,uint32,uint32)[] rangePayments) returns()
func (_IPaymentCoordinator *IPaymentCoordinatorTransactor) PayForRange(opts *bind.TransactOpts, rangePayments []IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _IPaymentCoordinator.contract.Transact(opts, "payForRange", rangePayments)
}

// PayForRange is a paid mutator transaction binding the contract method 0x1b445516.
//
// Solidity: function payForRange(((address,uint96)[],address,uint256,uint32,uint32)[] rangePayments) returns()
func (_IPaymentCoordinator *IPaymentCoordinatorSession) PayForRange(rangePayments []IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _IPaymentCoordinator.Contract.PayForRange(&_IPaymentCoordinator.TransactOpts, rangePayments)
}

// PayForRange is a paid mutator transaction binding the contract method 0x1b445516.
//
// Solidity: function payForRange(((address,uint96)[],address,uint256,uint32,uint32)[] rangePayments) returns()
func (_IPaymentCoordinator *IPaymentCoordinatorTransactorSession) PayForRange(rangePayments []IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _IPaymentCoordinator.Contract.PayForRange(&_IPaymentCoordinator.TransactOpts, rangePayments)
}

// ProcessClaim is a paid mutator transaction binding the contract method 0x3ccc861d.
//
// Solidity: function processClaim((uint32,uint32,bytes,(address,bytes32),uint32[],bytes[],(address,uint256)[]) claim, address recipient) returns()
func (_IPaymentCoordinator *IPaymentCoordinatorTransactor) ProcessClaim(opts *bind.TransactOpts, claim IPaymentCoordinatorPaymentMerkleClaim, recipient common.Address) (*types.Transaction, error) {
	return _IPaymentCoordinator.contract.Transact(opts, "processClaim", claim, recipient)
}

// ProcessClaim is a paid mutator transaction binding the contract method 0x3ccc861d.
//
// Solidity: function processClaim((uint32,uint32,bytes,(address,bytes32),uint32[],bytes[],(address,uint256)[]) claim, address recipient) returns()
func (_IPaymentCoordinator *IPaymentCoordinatorSession) ProcessClaim(claim IPaymentCoordinatorPaymentMerkleClaim, recipient common.Address) (*types.Transaction, error) {
	return _IPaymentCoordinator.Contract.ProcessClaim(&_IPaymentCoordinator.TransactOpts, claim, recipient)
}

// ProcessClaim is a paid mutator transaction binding the contract method 0x3ccc861d.
//
// Solidity: function processClaim((uint32,uint32,bytes,(address,bytes32),uint32[],bytes[],(address,uint256)[]) claim, address recipient) returns()
func (_IPaymentCoordinator *IPaymentCoordinatorTransactorSession) ProcessClaim(claim IPaymentCoordinatorPaymentMerkleClaim, recipient common.Address) (*types.Transaction, error) {
	return _IPaymentCoordinator.Contract.ProcessClaim(&_IPaymentCoordinator.TransactOpts, claim, recipient)
}

// SetActivationDelay is a paid mutator transaction binding the contract method 0x58baaa3e.
//
// Solidity: function setActivationDelay(uint32 _activationDelay) returns()
func (_IPaymentCoordinator *IPaymentCoordinatorTransactor) SetActivationDelay(opts *bind.TransactOpts, _activationDelay uint32) (*types.Transaction, error) {
	return _IPaymentCoordinator.contract.Transact(opts, "setActivationDelay", _activationDelay)
}

// SetActivationDelay is a paid mutator transaction binding the contract method 0x58baaa3e.
//
// Solidity: function setActivationDelay(uint32 _activationDelay) returns()
func (_IPaymentCoordinator *IPaymentCoordinatorSession) SetActivationDelay(_activationDelay uint32) (*types.Transaction, error) {
	return _IPaymentCoordinator.Contract.SetActivationDelay(&_IPaymentCoordinator.TransactOpts, _activationDelay)
}

// SetActivationDelay is a paid mutator transaction binding the contract method 0x58baaa3e.
//
// Solidity: function setActivationDelay(uint32 _activationDelay) returns()
func (_IPaymentCoordinator *IPaymentCoordinatorTransactorSession) SetActivationDelay(_activationDelay uint32) (*types.Transaction, error) {
	return _IPaymentCoordinator.Contract.SetActivationDelay(&_IPaymentCoordinator.TransactOpts, _activationDelay)
}

// SetClaimerFor is a paid mutator transaction binding the contract method 0xa0169ddd.
//
// Solidity: function setClaimerFor(address claimer) returns()
func (_IPaymentCoordinator *IPaymentCoordinatorTransactor) SetClaimerFor(opts *bind.TransactOpts, claimer common.Address) (*types.Transaction, error) {
	return _IPaymentCoordinator.contract.Transact(opts, "setClaimerFor", claimer)
}

// SetClaimerFor is a paid mutator transaction binding the contract method 0xa0169ddd.
//
// Solidity: function setClaimerFor(address claimer) returns()
func (_IPaymentCoordinator *IPaymentCoordinatorSession) SetClaimerFor(claimer common.Address) (*types.Transaction, error) {
	return _IPaymentCoordinator.Contract.SetClaimerFor(&_IPaymentCoordinator.TransactOpts, claimer)
}

// SetClaimerFor is a paid mutator transaction binding the contract method 0xa0169ddd.
//
// Solidity: function setClaimerFor(address claimer) returns()
func (_IPaymentCoordinator *IPaymentCoordinatorTransactorSession) SetClaimerFor(claimer common.Address) (*types.Transaction, error) {
	return _IPaymentCoordinator.Contract.SetClaimerFor(&_IPaymentCoordinator.TransactOpts, claimer)
}

// SetGlobalOperatorCommission is a paid mutator transaction binding the contract method 0xe221b245.
//
// Solidity: function setGlobalOperatorCommission(uint16 _globalCommissionBips) returns()
func (_IPaymentCoordinator *IPaymentCoordinatorTransactor) SetGlobalOperatorCommission(opts *bind.TransactOpts, _globalCommissionBips uint16) (*types.Transaction, error) {
	return _IPaymentCoordinator.contract.Transact(opts, "setGlobalOperatorCommission", _globalCommissionBips)
}

// SetGlobalOperatorCommission is a paid mutator transaction binding the contract method 0xe221b245.
//
// Solidity: function setGlobalOperatorCommission(uint16 _globalCommissionBips) returns()
func (_IPaymentCoordinator *IPaymentCoordinatorSession) SetGlobalOperatorCommission(_globalCommissionBips uint16) (*types.Transaction, error) {
	return _IPaymentCoordinator.Contract.SetGlobalOperatorCommission(&_IPaymentCoordinator.TransactOpts, _globalCommissionBips)
}

// SetGlobalOperatorCommission is a paid mutator transaction binding the contract method 0xe221b245.
//
// Solidity: function setGlobalOperatorCommission(uint16 _globalCommissionBips) returns()
func (_IPaymentCoordinator *IPaymentCoordinatorTransactorSession) SetGlobalOperatorCommission(_globalCommissionBips uint16) (*types.Transaction, error) {
	return _IPaymentCoordinator.Contract.SetGlobalOperatorCommission(&_IPaymentCoordinator.TransactOpts, _globalCommissionBips)
}

// SetPayAllForRangeSubmitter is a paid mutator transaction binding the contract method 0xec1680de.
//
// Solidity: function setPayAllForRangeSubmitter(address _submitter, bool _newValue) returns()
func (_IPaymentCoordinator *IPaymentCoordinatorTransactor) SetPayAllForRangeSubmitter(opts *bind.TransactOpts, _submitter common.Address, _newValue bool) (*types.Transaction, error) {
	return _IPaymentCoordinator.contract.Transact(opts, "setPayAllForRangeSubmitter", _submitter, _newValue)
}

// SetPayAllForRangeSubmitter is a paid mutator transaction binding the contract method 0xec1680de.
//
// Solidity: function setPayAllForRangeSubmitter(address _submitter, bool _newValue) returns()
func (_IPaymentCoordinator *IPaymentCoordinatorSession) SetPayAllForRangeSubmitter(_submitter common.Address, _newValue bool) (*types.Transaction, error) {
	return _IPaymentCoordinator.Contract.SetPayAllForRangeSubmitter(&_IPaymentCoordinator.TransactOpts, _submitter, _newValue)
}

// SetPayAllForRangeSubmitter is a paid mutator transaction binding the contract method 0xec1680de.
//
// Solidity: function setPayAllForRangeSubmitter(address _submitter, bool _newValue) returns()
func (_IPaymentCoordinator *IPaymentCoordinatorTransactorSession) SetPayAllForRangeSubmitter(_submitter common.Address, _newValue bool) (*types.Transaction, error) {
	return _IPaymentCoordinator.Contract.SetPayAllForRangeSubmitter(&_IPaymentCoordinator.TransactOpts, _submitter, _newValue)
}

// SetPaymentUpdater is a paid mutator transaction binding the contract method 0x18190f53.
//
// Solidity: function setPaymentUpdater(address _paymentUpdater) returns()
func (_IPaymentCoordinator *IPaymentCoordinatorTransactor) SetPaymentUpdater(opts *bind.TransactOpts, _paymentUpdater common.Address) (*types.Transaction, error) {
	return _IPaymentCoordinator.contract.Transact(opts, "setPaymentUpdater", _paymentUpdater)
}

// SetPaymentUpdater is a paid mutator transaction binding the contract method 0x18190f53.
//
// Solidity: function setPaymentUpdater(address _paymentUpdater) returns()
func (_IPaymentCoordinator *IPaymentCoordinatorSession) SetPaymentUpdater(_paymentUpdater common.Address) (*types.Transaction, error) {
	return _IPaymentCoordinator.Contract.SetPaymentUpdater(&_IPaymentCoordinator.TransactOpts, _paymentUpdater)
}

// SetPaymentUpdater is a paid mutator transaction binding the contract method 0x18190f53.
//
// Solidity: function setPaymentUpdater(address _paymentUpdater) returns()
func (_IPaymentCoordinator *IPaymentCoordinatorTransactorSession) SetPaymentUpdater(_paymentUpdater common.Address) (*types.Transaction, error) {
	return _IPaymentCoordinator.Contract.SetPaymentUpdater(&_IPaymentCoordinator.TransactOpts, _paymentUpdater)
}

// SubmitRoot is a paid mutator transaction binding the contract method 0x3efe1db6.
//
// Solidity: function submitRoot(bytes32 root, uint32 paymentCalculationEndTimestamp) returns()
func (_IPaymentCoordinator *IPaymentCoordinatorTransactor) SubmitRoot(opts *bind.TransactOpts, root [32]byte, paymentCalculationEndTimestamp uint32) (*types.Transaction, error) {
	return _IPaymentCoordinator.contract.Transact(opts, "submitRoot", root, paymentCalculationEndTimestamp)
}

// SubmitRoot is a paid mutator transaction binding the contract method 0x3efe1db6.
//
// Solidity: function submitRoot(bytes32 root, uint32 paymentCalculationEndTimestamp) returns()
func (_IPaymentCoordinator *IPaymentCoordinatorSession) SubmitRoot(root [32]byte, paymentCalculationEndTimestamp uint32) (*types.Transaction, error) {
	return _IPaymentCoordinator.Contract.SubmitRoot(&_IPaymentCoordinator.TransactOpts, root, paymentCalculationEndTimestamp)
}

// SubmitRoot is a paid mutator transaction binding the contract method 0x3efe1db6.
//
// Solidity: function submitRoot(bytes32 root, uint32 paymentCalculationEndTimestamp) returns()
func (_IPaymentCoordinator *IPaymentCoordinatorTransactorSession) SubmitRoot(root [32]byte, paymentCalculationEndTimestamp uint32) (*types.Transaction, error) {
	return _IPaymentCoordinator.Contract.SubmitRoot(&_IPaymentCoordinator.TransactOpts, root, paymentCalculationEndTimestamp)
}

// IPaymentCoordinatorActivationDelaySetIterator is returned from FilterActivationDelaySet and is used to iterate over the raw logs and unpacked data for ActivationDelaySet events raised by the IPaymentCoordinator contract.
type IPaymentCoordinatorActivationDelaySetIterator struct {
	Event *IPaymentCoordinatorActivationDelaySet // Event containing the contract specifics and raw log

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
func (it *IPaymentCoordinatorActivationDelaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPaymentCoordinatorActivationDelaySet)
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
		it.Event = new(IPaymentCoordinatorActivationDelaySet)
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
func (it *IPaymentCoordinatorActivationDelaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPaymentCoordinatorActivationDelaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPaymentCoordinatorActivationDelaySet represents a ActivationDelaySet event raised by the IPaymentCoordinator contract.
type IPaymentCoordinatorActivationDelaySet struct {
	OldActivationDelay uint32
	NewActivationDelay uint32
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterActivationDelaySet is a free log retrieval operation binding the contract event 0xaf557c6c02c208794817a705609cfa935f827312a1adfdd26494b6b95dd2b4b3.
//
// Solidity: event ActivationDelaySet(uint32 oldActivationDelay, uint32 newActivationDelay)
func (_IPaymentCoordinator *IPaymentCoordinatorFilterer) FilterActivationDelaySet(opts *bind.FilterOpts) (*IPaymentCoordinatorActivationDelaySetIterator, error) {

	logs, sub, err := _IPaymentCoordinator.contract.FilterLogs(opts, "ActivationDelaySet")
	if err != nil {
		return nil, err
	}
	return &IPaymentCoordinatorActivationDelaySetIterator{contract: _IPaymentCoordinator.contract, event: "ActivationDelaySet", logs: logs, sub: sub}, nil
}

// WatchActivationDelaySet is a free log subscription operation binding the contract event 0xaf557c6c02c208794817a705609cfa935f827312a1adfdd26494b6b95dd2b4b3.
//
// Solidity: event ActivationDelaySet(uint32 oldActivationDelay, uint32 newActivationDelay)
func (_IPaymentCoordinator *IPaymentCoordinatorFilterer) WatchActivationDelaySet(opts *bind.WatchOpts, sink chan<- *IPaymentCoordinatorActivationDelaySet) (event.Subscription, error) {

	logs, sub, err := _IPaymentCoordinator.contract.WatchLogs(opts, "ActivationDelaySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPaymentCoordinatorActivationDelaySet)
				if err := _IPaymentCoordinator.contract.UnpackLog(event, "ActivationDelaySet", log); err != nil {
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

// ParseActivationDelaySet is a log parse operation binding the contract event 0xaf557c6c02c208794817a705609cfa935f827312a1adfdd26494b6b95dd2b4b3.
//
// Solidity: event ActivationDelaySet(uint32 oldActivationDelay, uint32 newActivationDelay)
func (_IPaymentCoordinator *IPaymentCoordinatorFilterer) ParseActivationDelaySet(log types.Log) (*IPaymentCoordinatorActivationDelaySet, error) {
	event := new(IPaymentCoordinatorActivationDelaySet)
	if err := _IPaymentCoordinator.contract.UnpackLog(event, "ActivationDelaySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPaymentCoordinatorClaimerForSetIterator is returned from FilterClaimerForSet and is used to iterate over the raw logs and unpacked data for ClaimerForSet events raised by the IPaymentCoordinator contract.
type IPaymentCoordinatorClaimerForSetIterator struct {
	Event *IPaymentCoordinatorClaimerForSet // Event containing the contract specifics and raw log

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
func (it *IPaymentCoordinatorClaimerForSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPaymentCoordinatorClaimerForSet)
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
		it.Event = new(IPaymentCoordinatorClaimerForSet)
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
func (it *IPaymentCoordinatorClaimerForSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPaymentCoordinatorClaimerForSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPaymentCoordinatorClaimerForSet represents a ClaimerForSet event raised by the IPaymentCoordinator contract.
type IPaymentCoordinatorClaimerForSet struct {
	Earner     common.Address
	OldClaimer common.Address
	Claimer    common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterClaimerForSet is a free log retrieval operation binding the contract event 0xbab947934d42e0ad206f25c9cab18b5bb6ae144acfb00f40b4e3aa59590ca312.
//
// Solidity: event ClaimerForSet(address indexed earner, address indexed oldClaimer, address indexed claimer)
func (_IPaymentCoordinator *IPaymentCoordinatorFilterer) FilterClaimerForSet(opts *bind.FilterOpts, earner []common.Address, oldClaimer []common.Address, claimer []common.Address) (*IPaymentCoordinatorClaimerForSetIterator, error) {

	var earnerRule []interface{}
	for _, earnerItem := range earner {
		earnerRule = append(earnerRule, earnerItem)
	}
	var oldClaimerRule []interface{}
	for _, oldClaimerItem := range oldClaimer {
		oldClaimerRule = append(oldClaimerRule, oldClaimerItem)
	}
	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}

	logs, sub, err := _IPaymentCoordinator.contract.FilterLogs(opts, "ClaimerForSet", earnerRule, oldClaimerRule, claimerRule)
	if err != nil {
		return nil, err
	}
	return &IPaymentCoordinatorClaimerForSetIterator{contract: _IPaymentCoordinator.contract, event: "ClaimerForSet", logs: logs, sub: sub}, nil
}

// WatchClaimerForSet is a free log subscription operation binding the contract event 0xbab947934d42e0ad206f25c9cab18b5bb6ae144acfb00f40b4e3aa59590ca312.
//
// Solidity: event ClaimerForSet(address indexed earner, address indexed oldClaimer, address indexed claimer)
func (_IPaymentCoordinator *IPaymentCoordinatorFilterer) WatchClaimerForSet(opts *bind.WatchOpts, sink chan<- *IPaymentCoordinatorClaimerForSet, earner []common.Address, oldClaimer []common.Address, claimer []common.Address) (event.Subscription, error) {

	var earnerRule []interface{}
	for _, earnerItem := range earner {
		earnerRule = append(earnerRule, earnerItem)
	}
	var oldClaimerRule []interface{}
	for _, oldClaimerItem := range oldClaimer {
		oldClaimerRule = append(oldClaimerRule, oldClaimerItem)
	}
	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}

	logs, sub, err := _IPaymentCoordinator.contract.WatchLogs(opts, "ClaimerForSet", earnerRule, oldClaimerRule, claimerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPaymentCoordinatorClaimerForSet)
				if err := _IPaymentCoordinator.contract.UnpackLog(event, "ClaimerForSet", log); err != nil {
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

// ParseClaimerForSet is a log parse operation binding the contract event 0xbab947934d42e0ad206f25c9cab18b5bb6ae144acfb00f40b4e3aa59590ca312.
//
// Solidity: event ClaimerForSet(address indexed earner, address indexed oldClaimer, address indexed claimer)
func (_IPaymentCoordinator *IPaymentCoordinatorFilterer) ParseClaimerForSet(log types.Log) (*IPaymentCoordinatorClaimerForSet, error) {
	event := new(IPaymentCoordinatorClaimerForSet)
	if err := _IPaymentCoordinator.contract.UnpackLog(event, "ClaimerForSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPaymentCoordinatorDistributionRootSubmittedIterator is returned from FilterDistributionRootSubmitted and is used to iterate over the raw logs and unpacked data for DistributionRootSubmitted events raised by the IPaymentCoordinator contract.
type IPaymentCoordinatorDistributionRootSubmittedIterator struct {
	Event *IPaymentCoordinatorDistributionRootSubmitted // Event containing the contract specifics and raw log

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
func (it *IPaymentCoordinatorDistributionRootSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPaymentCoordinatorDistributionRootSubmitted)
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
		it.Event = new(IPaymentCoordinatorDistributionRootSubmitted)
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
func (it *IPaymentCoordinatorDistributionRootSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPaymentCoordinatorDistributionRootSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPaymentCoordinatorDistributionRootSubmitted represents a DistributionRootSubmitted event raised by the IPaymentCoordinator contract.
type IPaymentCoordinatorDistributionRootSubmitted struct {
	RootIndex                      uint32
	Root                           [32]byte
	PaymentCalculationEndTimestamp uint32
	ActivatedAt                    uint32
	Raw                            types.Log // Blockchain specific contextual infos
}

// FilterDistributionRootSubmitted is a free log retrieval operation binding the contract event 0xecd866c3c158fa00bf34d803d5f6023000b57080bcb48af004c2b4b46b3afd08.
//
// Solidity: event DistributionRootSubmitted(uint32 indexed rootIndex, bytes32 indexed root, uint32 indexed paymentCalculationEndTimestamp, uint32 activatedAt)
func (_IPaymentCoordinator *IPaymentCoordinatorFilterer) FilterDistributionRootSubmitted(opts *bind.FilterOpts, rootIndex []uint32, root [][32]byte, paymentCalculationEndTimestamp []uint32) (*IPaymentCoordinatorDistributionRootSubmittedIterator, error) {

	var rootIndexRule []interface{}
	for _, rootIndexItem := range rootIndex {
		rootIndexRule = append(rootIndexRule, rootIndexItem)
	}
	var rootRule []interface{}
	for _, rootItem := range root {
		rootRule = append(rootRule, rootItem)
	}
	var paymentCalculationEndTimestampRule []interface{}
	for _, paymentCalculationEndTimestampItem := range paymentCalculationEndTimestamp {
		paymentCalculationEndTimestampRule = append(paymentCalculationEndTimestampRule, paymentCalculationEndTimestampItem)
	}

	logs, sub, err := _IPaymentCoordinator.contract.FilterLogs(opts, "DistributionRootSubmitted", rootIndexRule, rootRule, paymentCalculationEndTimestampRule)
	if err != nil {
		return nil, err
	}
	return &IPaymentCoordinatorDistributionRootSubmittedIterator{contract: _IPaymentCoordinator.contract, event: "DistributionRootSubmitted", logs: logs, sub: sub}, nil
}

// WatchDistributionRootSubmitted is a free log subscription operation binding the contract event 0xecd866c3c158fa00bf34d803d5f6023000b57080bcb48af004c2b4b46b3afd08.
//
// Solidity: event DistributionRootSubmitted(uint32 indexed rootIndex, bytes32 indexed root, uint32 indexed paymentCalculationEndTimestamp, uint32 activatedAt)
func (_IPaymentCoordinator *IPaymentCoordinatorFilterer) WatchDistributionRootSubmitted(opts *bind.WatchOpts, sink chan<- *IPaymentCoordinatorDistributionRootSubmitted, rootIndex []uint32, root [][32]byte, paymentCalculationEndTimestamp []uint32) (event.Subscription, error) {

	var rootIndexRule []interface{}
	for _, rootIndexItem := range rootIndex {
		rootIndexRule = append(rootIndexRule, rootIndexItem)
	}
	var rootRule []interface{}
	for _, rootItem := range root {
		rootRule = append(rootRule, rootItem)
	}
	var paymentCalculationEndTimestampRule []interface{}
	for _, paymentCalculationEndTimestampItem := range paymentCalculationEndTimestamp {
		paymentCalculationEndTimestampRule = append(paymentCalculationEndTimestampRule, paymentCalculationEndTimestampItem)
	}

	logs, sub, err := _IPaymentCoordinator.contract.WatchLogs(opts, "DistributionRootSubmitted", rootIndexRule, rootRule, paymentCalculationEndTimestampRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPaymentCoordinatorDistributionRootSubmitted)
				if err := _IPaymentCoordinator.contract.UnpackLog(event, "DistributionRootSubmitted", log); err != nil {
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

// ParseDistributionRootSubmitted is a log parse operation binding the contract event 0xecd866c3c158fa00bf34d803d5f6023000b57080bcb48af004c2b4b46b3afd08.
//
// Solidity: event DistributionRootSubmitted(uint32 indexed rootIndex, bytes32 indexed root, uint32 indexed paymentCalculationEndTimestamp, uint32 activatedAt)
func (_IPaymentCoordinator *IPaymentCoordinatorFilterer) ParseDistributionRootSubmitted(log types.Log) (*IPaymentCoordinatorDistributionRootSubmitted, error) {
	event := new(IPaymentCoordinatorDistributionRootSubmitted)
	if err := _IPaymentCoordinator.contract.UnpackLog(event, "DistributionRootSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPaymentCoordinatorGlobalCommissionBipsSetIterator is returned from FilterGlobalCommissionBipsSet and is used to iterate over the raw logs and unpacked data for GlobalCommissionBipsSet events raised by the IPaymentCoordinator contract.
type IPaymentCoordinatorGlobalCommissionBipsSetIterator struct {
	Event *IPaymentCoordinatorGlobalCommissionBipsSet // Event containing the contract specifics and raw log

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
func (it *IPaymentCoordinatorGlobalCommissionBipsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPaymentCoordinatorGlobalCommissionBipsSet)
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
		it.Event = new(IPaymentCoordinatorGlobalCommissionBipsSet)
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
func (it *IPaymentCoordinatorGlobalCommissionBipsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPaymentCoordinatorGlobalCommissionBipsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPaymentCoordinatorGlobalCommissionBipsSet represents a GlobalCommissionBipsSet event raised by the IPaymentCoordinator contract.
type IPaymentCoordinatorGlobalCommissionBipsSet struct {
	OldGlobalCommissionBips uint16
	NewGlobalCommissionBips uint16
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterGlobalCommissionBipsSet is a free log retrieval operation binding the contract event 0x8cdc428b0431b82d1619763f443a48197db344ba96905f3949643acd1c863a06.
//
// Solidity: event GlobalCommissionBipsSet(uint16 oldGlobalCommissionBips, uint16 newGlobalCommissionBips)
func (_IPaymentCoordinator *IPaymentCoordinatorFilterer) FilterGlobalCommissionBipsSet(opts *bind.FilterOpts) (*IPaymentCoordinatorGlobalCommissionBipsSetIterator, error) {

	logs, sub, err := _IPaymentCoordinator.contract.FilterLogs(opts, "GlobalCommissionBipsSet")
	if err != nil {
		return nil, err
	}
	return &IPaymentCoordinatorGlobalCommissionBipsSetIterator{contract: _IPaymentCoordinator.contract, event: "GlobalCommissionBipsSet", logs: logs, sub: sub}, nil
}

// WatchGlobalCommissionBipsSet is a free log subscription operation binding the contract event 0x8cdc428b0431b82d1619763f443a48197db344ba96905f3949643acd1c863a06.
//
// Solidity: event GlobalCommissionBipsSet(uint16 oldGlobalCommissionBips, uint16 newGlobalCommissionBips)
func (_IPaymentCoordinator *IPaymentCoordinatorFilterer) WatchGlobalCommissionBipsSet(opts *bind.WatchOpts, sink chan<- *IPaymentCoordinatorGlobalCommissionBipsSet) (event.Subscription, error) {

	logs, sub, err := _IPaymentCoordinator.contract.WatchLogs(opts, "GlobalCommissionBipsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPaymentCoordinatorGlobalCommissionBipsSet)
				if err := _IPaymentCoordinator.contract.UnpackLog(event, "GlobalCommissionBipsSet", log); err != nil {
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

// ParseGlobalCommissionBipsSet is a log parse operation binding the contract event 0x8cdc428b0431b82d1619763f443a48197db344ba96905f3949643acd1c863a06.
//
// Solidity: event GlobalCommissionBipsSet(uint16 oldGlobalCommissionBips, uint16 newGlobalCommissionBips)
func (_IPaymentCoordinator *IPaymentCoordinatorFilterer) ParseGlobalCommissionBipsSet(log types.Log) (*IPaymentCoordinatorGlobalCommissionBipsSet, error) {
	event := new(IPaymentCoordinatorGlobalCommissionBipsSet)
	if err := _IPaymentCoordinator.contract.UnpackLog(event, "GlobalCommissionBipsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPaymentCoordinatorPayAllForRangeSubmitterSetIterator is returned from FilterPayAllForRangeSubmitterSet and is used to iterate over the raw logs and unpacked data for PayAllForRangeSubmitterSet events raised by the IPaymentCoordinator contract.
type IPaymentCoordinatorPayAllForRangeSubmitterSetIterator struct {
	Event *IPaymentCoordinatorPayAllForRangeSubmitterSet // Event containing the contract specifics and raw log

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
func (it *IPaymentCoordinatorPayAllForRangeSubmitterSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPaymentCoordinatorPayAllForRangeSubmitterSet)
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
		it.Event = new(IPaymentCoordinatorPayAllForRangeSubmitterSet)
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
func (it *IPaymentCoordinatorPayAllForRangeSubmitterSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPaymentCoordinatorPayAllForRangeSubmitterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPaymentCoordinatorPayAllForRangeSubmitterSet represents a PayAllForRangeSubmitterSet event raised by the IPaymentCoordinator contract.
type IPaymentCoordinatorPayAllForRangeSubmitterSet struct {
	PayAllForRangeSubmitter common.Address
	OldValue                bool
	NewValue                bool
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterPayAllForRangeSubmitterSet is a free log retrieval operation binding the contract event 0x17b06be02af2803593116eb96121b9c6e8bee1cc1b145e7c31c19c180e86189b.
//
// Solidity: event PayAllForRangeSubmitterSet(address indexed payAllForRangeSubmitter, bool indexed oldValue, bool indexed newValue)
func (_IPaymentCoordinator *IPaymentCoordinatorFilterer) FilterPayAllForRangeSubmitterSet(opts *bind.FilterOpts, payAllForRangeSubmitter []common.Address, oldValue []bool, newValue []bool) (*IPaymentCoordinatorPayAllForRangeSubmitterSetIterator, error) {

	var payAllForRangeSubmitterRule []interface{}
	for _, payAllForRangeSubmitterItem := range payAllForRangeSubmitter {
		payAllForRangeSubmitterRule = append(payAllForRangeSubmitterRule, payAllForRangeSubmitterItem)
	}
	var oldValueRule []interface{}
	for _, oldValueItem := range oldValue {
		oldValueRule = append(oldValueRule, oldValueItem)
	}
	var newValueRule []interface{}
	for _, newValueItem := range newValue {
		newValueRule = append(newValueRule, newValueItem)
	}

	logs, sub, err := _IPaymentCoordinator.contract.FilterLogs(opts, "PayAllForRangeSubmitterSet", payAllForRangeSubmitterRule, oldValueRule, newValueRule)
	if err != nil {
		return nil, err
	}
	return &IPaymentCoordinatorPayAllForRangeSubmitterSetIterator{contract: _IPaymentCoordinator.contract, event: "PayAllForRangeSubmitterSet", logs: logs, sub: sub}, nil
}

// WatchPayAllForRangeSubmitterSet is a free log subscription operation binding the contract event 0x17b06be02af2803593116eb96121b9c6e8bee1cc1b145e7c31c19c180e86189b.
//
// Solidity: event PayAllForRangeSubmitterSet(address indexed payAllForRangeSubmitter, bool indexed oldValue, bool indexed newValue)
func (_IPaymentCoordinator *IPaymentCoordinatorFilterer) WatchPayAllForRangeSubmitterSet(opts *bind.WatchOpts, sink chan<- *IPaymentCoordinatorPayAllForRangeSubmitterSet, payAllForRangeSubmitter []common.Address, oldValue []bool, newValue []bool) (event.Subscription, error) {

	var payAllForRangeSubmitterRule []interface{}
	for _, payAllForRangeSubmitterItem := range payAllForRangeSubmitter {
		payAllForRangeSubmitterRule = append(payAllForRangeSubmitterRule, payAllForRangeSubmitterItem)
	}
	var oldValueRule []interface{}
	for _, oldValueItem := range oldValue {
		oldValueRule = append(oldValueRule, oldValueItem)
	}
	var newValueRule []interface{}
	for _, newValueItem := range newValue {
		newValueRule = append(newValueRule, newValueItem)
	}

	logs, sub, err := _IPaymentCoordinator.contract.WatchLogs(opts, "PayAllForRangeSubmitterSet", payAllForRangeSubmitterRule, oldValueRule, newValueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPaymentCoordinatorPayAllForRangeSubmitterSet)
				if err := _IPaymentCoordinator.contract.UnpackLog(event, "PayAllForRangeSubmitterSet", log); err != nil {
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

// ParsePayAllForRangeSubmitterSet is a log parse operation binding the contract event 0x17b06be02af2803593116eb96121b9c6e8bee1cc1b145e7c31c19c180e86189b.
//
// Solidity: event PayAllForRangeSubmitterSet(address indexed payAllForRangeSubmitter, bool indexed oldValue, bool indexed newValue)
func (_IPaymentCoordinator *IPaymentCoordinatorFilterer) ParsePayAllForRangeSubmitterSet(log types.Log) (*IPaymentCoordinatorPayAllForRangeSubmitterSet, error) {
	event := new(IPaymentCoordinatorPayAllForRangeSubmitterSet)
	if err := _IPaymentCoordinator.contract.UnpackLog(event, "PayAllForRangeSubmitterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPaymentCoordinatorPaymentClaimedIterator is returned from FilterPaymentClaimed and is used to iterate over the raw logs and unpacked data for PaymentClaimed events raised by the IPaymentCoordinator contract.
type IPaymentCoordinatorPaymentClaimedIterator struct {
	Event *IPaymentCoordinatorPaymentClaimed // Event containing the contract specifics and raw log

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
func (it *IPaymentCoordinatorPaymentClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPaymentCoordinatorPaymentClaimed)
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
		it.Event = new(IPaymentCoordinatorPaymentClaimed)
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
func (it *IPaymentCoordinatorPaymentClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPaymentCoordinatorPaymentClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPaymentCoordinatorPaymentClaimed represents a PaymentClaimed event raised by the IPaymentCoordinator contract.
type IPaymentCoordinatorPaymentClaimed struct {
	Root          [32]byte
	Earner        common.Address
	Claimer       common.Address
	Recipient     common.Address
	Token         common.Address
	ClaimedAmount *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPaymentClaimed is a free log retrieval operation binding the contract event 0xbff1e5a32b3f6d3b3c0a7e675ead2091fea820852f35a77abdd6d2420bec4778.
//
// Solidity: event PaymentClaimed(bytes32 root, address indexed earner, address indexed claimer, address indexed recipient, address token, uint256 claimedAmount)
func (_IPaymentCoordinator *IPaymentCoordinatorFilterer) FilterPaymentClaimed(opts *bind.FilterOpts, earner []common.Address, claimer []common.Address, recipient []common.Address) (*IPaymentCoordinatorPaymentClaimedIterator, error) {

	var earnerRule []interface{}
	for _, earnerItem := range earner {
		earnerRule = append(earnerRule, earnerItem)
	}
	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _IPaymentCoordinator.contract.FilterLogs(opts, "PaymentClaimed", earnerRule, claimerRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &IPaymentCoordinatorPaymentClaimedIterator{contract: _IPaymentCoordinator.contract, event: "PaymentClaimed", logs: logs, sub: sub}, nil
}

// WatchPaymentClaimed is a free log subscription operation binding the contract event 0xbff1e5a32b3f6d3b3c0a7e675ead2091fea820852f35a77abdd6d2420bec4778.
//
// Solidity: event PaymentClaimed(bytes32 root, address indexed earner, address indexed claimer, address indexed recipient, address token, uint256 claimedAmount)
func (_IPaymentCoordinator *IPaymentCoordinatorFilterer) WatchPaymentClaimed(opts *bind.WatchOpts, sink chan<- *IPaymentCoordinatorPaymentClaimed, earner []common.Address, claimer []common.Address, recipient []common.Address) (event.Subscription, error) {

	var earnerRule []interface{}
	for _, earnerItem := range earner {
		earnerRule = append(earnerRule, earnerItem)
	}
	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _IPaymentCoordinator.contract.WatchLogs(opts, "PaymentClaimed", earnerRule, claimerRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPaymentCoordinatorPaymentClaimed)
				if err := _IPaymentCoordinator.contract.UnpackLog(event, "PaymentClaimed", log); err != nil {
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

// ParsePaymentClaimed is a log parse operation binding the contract event 0xbff1e5a32b3f6d3b3c0a7e675ead2091fea820852f35a77abdd6d2420bec4778.
//
// Solidity: event PaymentClaimed(bytes32 root, address indexed earner, address indexed claimer, address indexed recipient, address token, uint256 claimedAmount)
func (_IPaymentCoordinator *IPaymentCoordinatorFilterer) ParsePaymentClaimed(log types.Log) (*IPaymentCoordinatorPaymentClaimed, error) {
	event := new(IPaymentCoordinatorPaymentClaimed)
	if err := _IPaymentCoordinator.contract.UnpackLog(event, "PaymentClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPaymentCoordinatorPaymentUpdaterSetIterator is returned from FilterPaymentUpdaterSet and is used to iterate over the raw logs and unpacked data for PaymentUpdaterSet events raised by the IPaymentCoordinator contract.
type IPaymentCoordinatorPaymentUpdaterSetIterator struct {
	Event *IPaymentCoordinatorPaymentUpdaterSet // Event containing the contract specifics and raw log

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
func (it *IPaymentCoordinatorPaymentUpdaterSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPaymentCoordinatorPaymentUpdaterSet)
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
		it.Event = new(IPaymentCoordinatorPaymentUpdaterSet)
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
func (it *IPaymentCoordinatorPaymentUpdaterSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPaymentCoordinatorPaymentUpdaterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPaymentCoordinatorPaymentUpdaterSet represents a PaymentUpdaterSet event raised by the IPaymentCoordinator contract.
type IPaymentCoordinatorPaymentUpdaterSet struct {
	OldPaymentUpdater common.Address
	NewPaymentUpdater common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPaymentUpdaterSet is a free log retrieval operation binding the contract event 0x07d2890b3eb1206e7c3cb6bf8d46da31385ace3ce99abf85e5b690c83aa49678.
//
// Solidity: event PaymentUpdaterSet(address indexed oldPaymentUpdater, address indexed newPaymentUpdater)
func (_IPaymentCoordinator *IPaymentCoordinatorFilterer) FilterPaymentUpdaterSet(opts *bind.FilterOpts, oldPaymentUpdater []common.Address, newPaymentUpdater []common.Address) (*IPaymentCoordinatorPaymentUpdaterSetIterator, error) {

	var oldPaymentUpdaterRule []interface{}
	for _, oldPaymentUpdaterItem := range oldPaymentUpdater {
		oldPaymentUpdaterRule = append(oldPaymentUpdaterRule, oldPaymentUpdaterItem)
	}
	var newPaymentUpdaterRule []interface{}
	for _, newPaymentUpdaterItem := range newPaymentUpdater {
		newPaymentUpdaterRule = append(newPaymentUpdaterRule, newPaymentUpdaterItem)
	}

	logs, sub, err := _IPaymentCoordinator.contract.FilterLogs(opts, "PaymentUpdaterSet", oldPaymentUpdaterRule, newPaymentUpdaterRule)
	if err != nil {
		return nil, err
	}
	return &IPaymentCoordinatorPaymentUpdaterSetIterator{contract: _IPaymentCoordinator.contract, event: "PaymentUpdaterSet", logs: logs, sub: sub}, nil
}

// WatchPaymentUpdaterSet is a free log subscription operation binding the contract event 0x07d2890b3eb1206e7c3cb6bf8d46da31385ace3ce99abf85e5b690c83aa49678.
//
// Solidity: event PaymentUpdaterSet(address indexed oldPaymentUpdater, address indexed newPaymentUpdater)
func (_IPaymentCoordinator *IPaymentCoordinatorFilterer) WatchPaymentUpdaterSet(opts *bind.WatchOpts, sink chan<- *IPaymentCoordinatorPaymentUpdaterSet, oldPaymentUpdater []common.Address, newPaymentUpdater []common.Address) (event.Subscription, error) {

	var oldPaymentUpdaterRule []interface{}
	for _, oldPaymentUpdaterItem := range oldPaymentUpdater {
		oldPaymentUpdaterRule = append(oldPaymentUpdaterRule, oldPaymentUpdaterItem)
	}
	var newPaymentUpdaterRule []interface{}
	for _, newPaymentUpdaterItem := range newPaymentUpdater {
		newPaymentUpdaterRule = append(newPaymentUpdaterRule, newPaymentUpdaterItem)
	}

	logs, sub, err := _IPaymentCoordinator.contract.WatchLogs(opts, "PaymentUpdaterSet", oldPaymentUpdaterRule, newPaymentUpdaterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPaymentCoordinatorPaymentUpdaterSet)
				if err := _IPaymentCoordinator.contract.UnpackLog(event, "PaymentUpdaterSet", log); err != nil {
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

// ParsePaymentUpdaterSet is a log parse operation binding the contract event 0x07d2890b3eb1206e7c3cb6bf8d46da31385ace3ce99abf85e5b690c83aa49678.
//
// Solidity: event PaymentUpdaterSet(address indexed oldPaymentUpdater, address indexed newPaymentUpdater)
func (_IPaymentCoordinator *IPaymentCoordinatorFilterer) ParsePaymentUpdaterSet(log types.Log) (*IPaymentCoordinatorPaymentUpdaterSet, error) {
	event := new(IPaymentCoordinatorPaymentUpdaterSet)
	if err := _IPaymentCoordinator.contract.UnpackLog(event, "PaymentUpdaterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPaymentCoordinatorRangePaymentCreatedIterator is returned from FilterRangePaymentCreated and is used to iterate over the raw logs and unpacked data for RangePaymentCreated events raised by the IPaymentCoordinator contract.
type IPaymentCoordinatorRangePaymentCreatedIterator struct {
	Event *IPaymentCoordinatorRangePaymentCreated // Event containing the contract specifics and raw log

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
func (it *IPaymentCoordinatorRangePaymentCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPaymentCoordinatorRangePaymentCreated)
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
		it.Event = new(IPaymentCoordinatorRangePaymentCreated)
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
func (it *IPaymentCoordinatorRangePaymentCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPaymentCoordinatorRangePaymentCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPaymentCoordinatorRangePaymentCreated represents a RangePaymentCreated event raised by the IPaymentCoordinator contract.
type IPaymentCoordinatorRangePaymentCreated struct {
	Avs              common.Address
	PaymentNonce     *big.Int
	RangePaymentHash [32]byte
	RangePayment     IPaymentCoordinatorRangePayment
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRangePaymentCreated is a free log retrieval operation binding the contract event 0x2a0659fa4c38e0075469a0e0dd737e045dc316ffd6cb6e68755c119ee0882aea.
//
// Solidity: event RangePaymentCreated(address indexed avs, uint256 indexed paymentNonce, bytes32 indexed rangePaymentHash, ((address,uint96)[],address,uint256,uint32,uint32) rangePayment)
func (_IPaymentCoordinator *IPaymentCoordinatorFilterer) FilterRangePaymentCreated(opts *bind.FilterOpts, avs []common.Address, paymentNonce []*big.Int, rangePaymentHash [][32]byte) (*IPaymentCoordinatorRangePaymentCreatedIterator, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}
	var paymentNonceRule []interface{}
	for _, paymentNonceItem := range paymentNonce {
		paymentNonceRule = append(paymentNonceRule, paymentNonceItem)
	}
	var rangePaymentHashRule []interface{}
	for _, rangePaymentHashItem := range rangePaymentHash {
		rangePaymentHashRule = append(rangePaymentHashRule, rangePaymentHashItem)
	}

	logs, sub, err := _IPaymentCoordinator.contract.FilterLogs(opts, "RangePaymentCreated", avsRule, paymentNonceRule, rangePaymentHashRule)
	if err != nil {
		return nil, err
	}
	return &IPaymentCoordinatorRangePaymentCreatedIterator{contract: _IPaymentCoordinator.contract, event: "RangePaymentCreated", logs: logs, sub: sub}, nil
}

// WatchRangePaymentCreated is a free log subscription operation binding the contract event 0x2a0659fa4c38e0075469a0e0dd737e045dc316ffd6cb6e68755c119ee0882aea.
//
// Solidity: event RangePaymentCreated(address indexed avs, uint256 indexed paymentNonce, bytes32 indexed rangePaymentHash, ((address,uint96)[],address,uint256,uint32,uint32) rangePayment)
func (_IPaymentCoordinator *IPaymentCoordinatorFilterer) WatchRangePaymentCreated(opts *bind.WatchOpts, sink chan<- *IPaymentCoordinatorRangePaymentCreated, avs []common.Address, paymentNonce []*big.Int, rangePaymentHash [][32]byte) (event.Subscription, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}
	var paymentNonceRule []interface{}
	for _, paymentNonceItem := range paymentNonce {
		paymentNonceRule = append(paymentNonceRule, paymentNonceItem)
	}
	var rangePaymentHashRule []interface{}
	for _, rangePaymentHashItem := range rangePaymentHash {
		rangePaymentHashRule = append(rangePaymentHashRule, rangePaymentHashItem)
	}

	logs, sub, err := _IPaymentCoordinator.contract.WatchLogs(opts, "RangePaymentCreated", avsRule, paymentNonceRule, rangePaymentHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPaymentCoordinatorRangePaymentCreated)
				if err := _IPaymentCoordinator.contract.UnpackLog(event, "RangePaymentCreated", log); err != nil {
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

// ParseRangePaymentCreated is a log parse operation binding the contract event 0x2a0659fa4c38e0075469a0e0dd737e045dc316ffd6cb6e68755c119ee0882aea.
//
// Solidity: event RangePaymentCreated(address indexed avs, uint256 indexed paymentNonce, bytes32 indexed rangePaymentHash, ((address,uint96)[],address,uint256,uint32,uint32) rangePayment)
func (_IPaymentCoordinator *IPaymentCoordinatorFilterer) ParseRangePaymentCreated(log types.Log) (*IPaymentCoordinatorRangePaymentCreated, error) {
	event := new(IPaymentCoordinatorRangePaymentCreated)
	if err := _IPaymentCoordinator.contract.UnpackLog(event, "RangePaymentCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPaymentCoordinatorRangePaymentForAllCreatedIterator is returned from FilterRangePaymentForAllCreated and is used to iterate over the raw logs and unpacked data for RangePaymentForAllCreated events raised by the IPaymentCoordinator contract.
type IPaymentCoordinatorRangePaymentForAllCreatedIterator struct {
	Event *IPaymentCoordinatorRangePaymentForAllCreated // Event containing the contract specifics and raw log

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
func (it *IPaymentCoordinatorRangePaymentForAllCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPaymentCoordinatorRangePaymentForAllCreated)
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
		it.Event = new(IPaymentCoordinatorRangePaymentForAllCreated)
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
func (it *IPaymentCoordinatorRangePaymentForAllCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPaymentCoordinatorRangePaymentForAllCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPaymentCoordinatorRangePaymentForAllCreated represents a RangePaymentForAllCreated event raised by the IPaymentCoordinator contract.
type IPaymentCoordinatorRangePaymentForAllCreated struct {
	Submitter        common.Address
	PaymentNonce     *big.Int
	RangePaymentHash [32]byte
	RangePayment     IPaymentCoordinatorRangePayment
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRangePaymentForAllCreated is a free log retrieval operation binding the contract event 0xbc0782940ec5871f66c5490ef957f44d19a9adac1dac18b946ad0dd6579c30d6.
//
// Solidity: event RangePaymentForAllCreated(address indexed submitter, uint256 indexed paymentNonce, bytes32 indexed rangePaymentHash, ((address,uint96)[],address,uint256,uint32,uint32) rangePayment)
func (_IPaymentCoordinator *IPaymentCoordinatorFilterer) FilterRangePaymentForAllCreated(opts *bind.FilterOpts, submitter []common.Address, paymentNonce []*big.Int, rangePaymentHash [][32]byte) (*IPaymentCoordinatorRangePaymentForAllCreatedIterator, error) {

	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}
	var paymentNonceRule []interface{}
	for _, paymentNonceItem := range paymentNonce {
		paymentNonceRule = append(paymentNonceRule, paymentNonceItem)
	}
	var rangePaymentHashRule []interface{}
	for _, rangePaymentHashItem := range rangePaymentHash {
		rangePaymentHashRule = append(rangePaymentHashRule, rangePaymentHashItem)
	}

	logs, sub, err := _IPaymentCoordinator.contract.FilterLogs(opts, "RangePaymentForAllCreated", submitterRule, paymentNonceRule, rangePaymentHashRule)
	if err != nil {
		return nil, err
	}
	return &IPaymentCoordinatorRangePaymentForAllCreatedIterator{contract: _IPaymentCoordinator.contract, event: "RangePaymentForAllCreated", logs: logs, sub: sub}, nil
}

// WatchRangePaymentForAllCreated is a free log subscription operation binding the contract event 0xbc0782940ec5871f66c5490ef957f44d19a9adac1dac18b946ad0dd6579c30d6.
//
// Solidity: event RangePaymentForAllCreated(address indexed submitter, uint256 indexed paymentNonce, bytes32 indexed rangePaymentHash, ((address,uint96)[],address,uint256,uint32,uint32) rangePayment)
func (_IPaymentCoordinator *IPaymentCoordinatorFilterer) WatchRangePaymentForAllCreated(opts *bind.WatchOpts, sink chan<- *IPaymentCoordinatorRangePaymentForAllCreated, submitter []common.Address, paymentNonce []*big.Int, rangePaymentHash [][32]byte) (event.Subscription, error) {

	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}
	var paymentNonceRule []interface{}
	for _, paymentNonceItem := range paymentNonce {
		paymentNonceRule = append(paymentNonceRule, paymentNonceItem)
	}
	var rangePaymentHashRule []interface{}
	for _, rangePaymentHashItem := range rangePaymentHash {
		rangePaymentHashRule = append(rangePaymentHashRule, rangePaymentHashItem)
	}

	logs, sub, err := _IPaymentCoordinator.contract.WatchLogs(opts, "RangePaymentForAllCreated", submitterRule, paymentNonceRule, rangePaymentHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPaymentCoordinatorRangePaymentForAllCreated)
				if err := _IPaymentCoordinator.contract.UnpackLog(event, "RangePaymentForAllCreated", log); err != nil {
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

// ParseRangePaymentForAllCreated is a log parse operation binding the contract event 0xbc0782940ec5871f66c5490ef957f44d19a9adac1dac18b946ad0dd6579c30d6.
//
// Solidity: event RangePaymentForAllCreated(address indexed submitter, uint256 indexed paymentNonce, bytes32 indexed rangePaymentHash, ((address,uint96)[],address,uint256,uint32,uint32) rangePayment)
func (_IPaymentCoordinator *IPaymentCoordinatorFilterer) ParseRangePaymentForAllCreated(log types.Log) (*IPaymentCoordinatorRangePaymentForAllCreated, error) {
	event := new(IPaymentCoordinatorRangePaymentForAllCreated)
	if err := _IPaymentCoordinator.contract.UnpackLog(event, "RangePaymentForAllCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
