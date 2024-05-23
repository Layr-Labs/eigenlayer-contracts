// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package PaymentCoordinatorStorage

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

// PaymentCoordinatorStorageMetaData contains all meta data concerning the PaymentCoordinatorStorage contract.
var PaymentCoordinatorStorageMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"CALCULATION_INTERVAL_SECONDS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GENESIS_PAYMENT_TIMESTAMP\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_FUTURE_LENGTH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_PAYMENT_DURATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_RETROACTIVE_LENGTH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"activationDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateEarnerLeafHash\",\"inputs\":[{\"name\":\"leaf\",\"type\":\"tuple\",\"internalType\":\"structIPaymentCoordinator.EarnerTreeMerkleLeaf\",\"components\":[{\"name\":\"earner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"earnerTokenRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"calculateTokenLeafHash\",\"inputs\":[{\"name\":\"leaf\",\"type\":\"tuple\",\"internalType\":\"structIPaymentCoordinator.TokenTreeMerkleLeaf\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"cumulativeEarnings\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"checkClaim\",\"inputs\":[{\"name\":\"claim\",\"type\":\"tuple\",\"internalType\":\"structIPaymentCoordinator.PaymentMerkleClaim\",\"components\":[{\"name\":\"rootIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"earnerIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"earnerTreeProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"earnerLeaf\",\"type\":\"tuple\",\"internalType\":\"structIPaymentCoordinator.EarnerTreeMerkleLeaf\",\"components\":[{\"name\":\"earner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"earnerTokenRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"tokenIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"tokenTreeProofs\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"tokenLeaves\",\"type\":\"tuple[]\",\"internalType\":\"structIPaymentCoordinator.TokenTreeMerkleLeaf[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"cumulativeEarnings\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimerFor\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cumulativeClaimed\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"currPaymentCalculationEndTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"distributionRoots\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"paymentCalculationEndTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"activatedAt\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDistributionRootsLength\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRootIndexFromHash\",\"inputs\":[{\"name\":\"rootHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"globalOperatorCommissionBips\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isPayAllForRangeSubmitter\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRangePaymentForAllHash\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRangePaymentHash\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorCommissionBips\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"payAllForRange\",\"inputs\":[{\"name\":\"rangePayment\",\"type\":\"tuple[]\",\"internalType\":\"structIPaymentCoordinator.RangePayment[]\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIPaymentCoordinator.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"payForRange\",\"inputs\":[{\"name\":\"rangePayments\",\"type\":\"tuple[]\",\"internalType\":\"structIPaymentCoordinator.RangePayment[]\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIPaymentCoordinator.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paymentNonce\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paymentUpdater\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"processClaim\",\"inputs\":[{\"name\":\"claim\",\"type\":\"tuple\",\"internalType\":\"structIPaymentCoordinator.PaymentMerkleClaim\",\"components\":[{\"name\":\"rootIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"earnerIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"earnerTreeProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"earnerLeaf\",\"type\":\"tuple\",\"internalType\":\"structIPaymentCoordinator.EarnerTreeMerkleLeaf\",\"components\":[{\"name\":\"earner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"earnerTokenRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"tokenIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"tokenTreeProofs\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"tokenLeaves\",\"type\":\"tuple[]\",\"internalType\":\"structIPaymentCoordinator.TokenTreeMerkleLeaf[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"cumulativeEarnings\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setActivationDelay\",\"inputs\":[{\"name\":\"_activationDelay\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setClaimerFor\",\"inputs\":[{\"name\":\"claimer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGlobalOperatorCommission\",\"inputs\":[{\"name\":\"_globalCommissionBips\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPayAllForRangeSubmitter\",\"inputs\":[{\"name\":\"_submitter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_newValue\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPaymentUpdater\",\"inputs\":[{\"name\":\"_paymentUpdater\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"strategyManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"submitRoot\",\"inputs\":[{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"paymentCalculationEndTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ActivationDelaySet\",\"inputs\":[{\"name\":\"oldActivationDelay\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"newActivationDelay\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ClaimerForSet\",\"inputs\":[{\"name\":\"earner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"oldClaimer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"claimer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DistributionRootSubmitted\",\"inputs\":[{\"name\":\"rootIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"root\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"paymentCalculationEndTimestamp\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"activatedAt\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GlobalCommissionBipsSet\",\"inputs\":[{\"name\":\"oldGlobalCommissionBips\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"newGlobalCommissionBips\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PayAllForRangeSubmitterSet\",\"inputs\":[{\"name\":\"payAllForRangeSubmitter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"oldValue\",\"type\":\"bool\",\"indexed\":true,\"internalType\":\"bool\"},{\"name\":\"newValue\",\"type\":\"bool\",\"indexed\":true,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PaymentClaimed\",\"inputs\":[{\"name\":\"root\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"earner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"claimer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIERC20\"},{\"name\":\"claimedAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PaymentUpdaterSet\",\"inputs\":[{\"name\":\"oldPaymentUpdater\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPaymentUpdater\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RangePaymentCreated\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"paymentNonce\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"rangePaymentHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"rangePayment\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIPaymentCoordinator.RangePayment\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIPaymentCoordinator.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RangePaymentForAllCreated\",\"inputs\":[{\"name\":\"submitter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"paymentNonce\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"rangePaymentHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"rangePayment\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIPaymentCoordinator.RangePayment\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIPaymentCoordinator.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false}]",
}

// PaymentCoordinatorStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use PaymentCoordinatorStorageMetaData.ABI instead.
var PaymentCoordinatorStorageABI = PaymentCoordinatorStorageMetaData.ABI

// PaymentCoordinatorStorage is an auto generated Go binding around an Ethereum contract.
type PaymentCoordinatorStorage struct {
	PaymentCoordinatorStorageCaller     // Read-only binding to the contract
	PaymentCoordinatorStorageTransactor // Write-only binding to the contract
	PaymentCoordinatorStorageFilterer   // Log filterer for contract events
}

// PaymentCoordinatorStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type PaymentCoordinatorStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentCoordinatorStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PaymentCoordinatorStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentCoordinatorStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PaymentCoordinatorStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentCoordinatorStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PaymentCoordinatorStorageSession struct {
	Contract     *PaymentCoordinatorStorage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// PaymentCoordinatorStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PaymentCoordinatorStorageCallerSession struct {
	Contract *PaymentCoordinatorStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// PaymentCoordinatorStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PaymentCoordinatorStorageTransactorSession struct {
	Contract     *PaymentCoordinatorStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// PaymentCoordinatorStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type PaymentCoordinatorStorageRaw struct {
	Contract *PaymentCoordinatorStorage // Generic contract binding to access the raw methods on
}

// PaymentCoordinatorStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PaymentCoordinatorStorageCallerRaw struct {
	Contract *PaymentCoordinatorStorageCaller // Generic read-only contract binding to access the raw methods on
}

// PaymentCoordinatorStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PaymentCoordinatorStorageTransactorRaw struct {
	Contract *PaymentCoordinatorStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPaymentCoordinatorStorage creates a new instance of PaymentCoordinatorStorage, bound to a specific deployed contract.
func NewPaymentCoordinatorStorage(address common.Address, backend bind.ContractBackend) (*PaymentCoordinatorStorage, error) {
	contract, err := bindPaymentCoordinatorStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PaymentCoordinatorStorage{PaymentCoordinatorStorageCaller: PaymentCoordinatorStorageCaller{contract: contract}, PaymentCoordinatorStorageTransactor: PaymentCoordinatorStorageTransactor{contract: contract}, PaymentCoordinatorStorageFilterer: PaymentCoordinatorStorageFilterer{contract: contract}}, nil
}

// NewPaymentCoordinatorStorageCaller creates a new read-only instance of PaymentCoordinatorStorage, bound to a specific deployed contract.
func NewPaymentCoordinatorStorageCaller(address common.Address, caller bind.ContractCaller) (*PaymentCoordinatorStorageCaller, error) {
	contract, err := bindPaymentCoordinatorStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PaymentCoordinatorStorageCaller{contract: contract}, nil
}

// NewPaymentCoordinatorStorageTransactor creates a new write-only instance of PaymentCoordinatorStorage, bound to a specific deployed contract.
func NewPaymentCoordinatorStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*PaymentCoordinatorStorageTransactor, error) {
	contract, err := bindPaymentCoordinatorStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PaymentCoordinatorStorageTransactor{contract: contract}, nil
}

// NewPaymentCoordinatorStorageFilterer creates a new log filterer instance of PaymentCoordinatorStorage, bound to a specific deployed contract.
func NewPaymentCoordinatorStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*PaymentCoordinatorStorageFilterer, error) {
	contract, err := bindPaymentCoordinatorStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PaymentCoordinatorStorageFilterer{contract: contract}, nil
}

// bindPaymentCoordinatorStorage binds a generic wrapper to an already deployed contract.
func bindPaymentCoordinatorStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PaymentCoordinatorStorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PaymentCoordinatorStorage.Contract.PaymentCoordinatorStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PaymentCoordinatorStorage.Contract.PaymentCoordinatorStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PaymentCoordinatorStorage.Contract.PaymentCoordinatorStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PaymentCoordinatorStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PaymentCoordinatorStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PaymentCoordinatorStorage.Contract.contract.Transact(opts, method, params...)
}

// CALCULATIONINTERVALSECONDS is a free data retrieval call binding the contract method 0x9d45c281.
//
// Solidity: function CALCULATION_INTERVAL_SECONDS() view returns(uint32)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageCaller) CALCULATIONINTERVALSECONDS(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _PaymentCoordinatorStorage.contract.Call(opts, &out, "CALCULATION_INTERVAL_SECONDS")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// CALCULATIONINTERVALSECONDS is a free data retrieval call binding the contract method 0x9d45c281.
//
// Solidity: function CALCULATION_INTERVAL_SECONDS() view returns(uint32)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageSession) CALCULATIONINTERVALSECONDS() (uint32, error) {
	return _PaymentCoordinatorStorage.Contract.CALCULATIONINTERVALSECONDS(&_PaymentCoordinatorStorage.CallOpts)
}

// CALCULATIONINTERVALSECONDS is a free data retrieval call binding the contract method 0x9d45c281.
//
// Solidity: function CALCULATION_INTERVAL_SECONDS() view returns(uint32)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageCallerSession) CALCULATIONINTERVALSECONDS() (uint32, error) {
	return _PaymentCoordinatorStorage.Contract.CALCULATIONINTERVALSECONDS(&_PaymentCoordinatorStorage.CallOpts)
}

// GENESISPAYMENTTIMESTAMP is a free data retrieval call binding the contract method 0x2cfd45eb.
//
// Solidity: function GENESIS_PAYMENT_TIMESTAMP() view returns(uint32)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageCaller) GENESISPAYMENTTIMESTAMP(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _PaymentCoordinatorStorage.contract.Call(opts, &out, "GENESIS_PAYMENT_TIMESTAMP")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GENESISPAYMENTTIMESTAMP is a free data retrieval call binding the contract method 0x2cfd45eb.
//
// Solidity: function GENESIS_PAYMENT_TIMESTAMP() view returns(uint32)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageSession) GENESISPAYMENTTIMESTAMP() (uint32, error) {
	return _PaymentCoordinatorStorage.Contract.GENESISPAYMENTTIMESTAMP(&_PaymentCoordinatorStorage.CallOpts)
}

// GENESISPAYMENTTIMESTAMP is a free data retrieval call binding the contract method 0x2cfd45eb.
//
// Solidity: function GENESIS_PAYMENT_TIMESTAMP() view returns(uint32)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageCallerSession) GENESISPAYMENTTIMESTAMP() (uint32, error) {
	return _PaymentCoordinatorStorage.Contract.GENESISPAYMENTTIMESTAMP(&_PaymentCoordinatorStorage.CallOpts)
}

// MAXFUTURELENGTH is a free data retrieval call binding the contract method 0x04a0c502.
//
// Solidity: function MAX_FUTURE_LENGTH() view returns(uint32)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageCaller) MAXFUTURELENGTH(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _PaymentCoordinatorStorage.contract.Call(opts, &out, "MAX_FUTURE_LENGTH")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MAXFUTURELENGTH is a free data retrieval call binding the contract method 0x04a0c502.
//
// Solidity: function MAX_FUTURE_LENGTH() view returns(uint32)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageSession) MAXFUTURELENGTH() (uint32, error) {
	return _PaymentCoordinatorStorage.Contract.MAXFUTURELENGTH(&_PaymentCoordinatorStorage.CallOpts)
}

// MAXFUTURELENGTH is a free data retrieval call binding the contract method 0x04a0c502.
//
// Solidity: function MAX_FUTURE_LENGTH() view returns(uint32)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageCallerSession) MAXFUTURELENGTH() (uint32, error) {
	return _PaymentCoordinatorStorage.Contract.MAXFUTURELENGTH(&_PaymentCoordinatorStorage.CallOpts)
}

// MAXPAYMENTDURATION is a free data retrieval call binding the contract method 0xee619597.
//
// Solidity: function MAX_PAYMENT_DURATION() view returns(uint32)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageCaller) MAXPAYMENTDURATION(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _PaymentCoordinatorStorage.contract.Call(opts, &out, "MAX_PAYMENT_DURATION")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MAXPAYMENTDURATION is a free data retrieval call binding the contract method 0xee619597.
//
// Solidity: function MAX_PAYMENT_DURATION() view returns(uint32)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageSession) MAXPAYMENTDURATION() (uint32, error) {
	return _PaymentCoordinatorStorage.Contract.MAXPAYMENTDURATION(&_PaymentCoordinatorStorage.CallOpts)
}

// MAXPAYMENTDURATION is a free data retrieval call binding the contract method 0xee619597.
//
// Solidity: function MAX_PAYMENT_DURATION() view returns(uint32)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageCallerSession) MAXPAYMENTDURATION() (uint32, error) {
	return _PaymentCoordinatorStorage.Contract.MAXPAYMENTDURATION(&_PaymentCoordinatorStorage.CallOpts)
}

// MAXRETROACTIVELENGTH is a free data retrieval call binding the contract method 0x37838ed0.
//
// Solidity: function MAX_RETROACTIVE_LENGTH() view returns(uint32)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageCaller) MAXRETROACTIVELENGTH(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _PaymentCoordinatorStorage.contract.Call(opts, &out, "MAX_RETROACTIVE_LENGTH")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MAXRETROACTIVELENGTH is a free data retrieval call binding the contract method 0x37838ed0.
//
// Solidity: function MAX_RETROACTIVE_LENGTH() view returns(uint32)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageSession) MAXRETROACTIVELENGTH() (uint32, error) {
	return _PaymentCoordinatorStorage.Contract.MAXRETROACTIVELENGTH(&_PaymentCoordinatorStorage.CallOpts)
}

// MAXRETROACTIVELENGTH is a free data retrieval call binding the contract method 0x37838ed0.
//
// Solidity: function MAX_RETROACTIVE_LENGTH() view returns(uint32)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageCallerSession) MAXRETROACTIVELENGTH() (uint32, error) {
	return _PaymentCoordinatorStorage.Contract.MAXRETROACTIVELENGTH(&_PaymentCoordinatorStorage.CallOpts)
}

// ActivationDelay is a free data retrieval call binding the contract method 0x3a8c0786.
//
// Solidity: function activationDelay() view returns(uint32)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageCaller) ActivationDelay(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _PaymentCoordinatorStorage.contract.Call(opts, &out, "activationDelay")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ActivationDelay is a free data retrieval call binding the contract method 0x3a8c0786.
//
// Solidity: function activationDelay() view returns(uint32)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageSession) ActivationDelay() (uint32, error) {
	return _PaymentCoordinatorStorage.Contract.ActivationDelay(&_PaymentCoordinatorStorage.CallOpts)
}

// ActivationDelay is a free data retrieval call binding the contract method 0x3a8c0786.
//
// Solidity: function activationDelay() view returns(uint32)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageCallerSession) ActivationDelay() (uint32, error) {
	return _PaymentCoordinatorStorage.Contract.ActivationDelay(&_PaymentCoordinatorStorage.CallOpts)
}

// CalculateEarnerLeafHash is a free data retrieval call binding the contract method 0x149bc872.
//
// Solidity: function calculateEarnerLeafHash((address,bytes32) leaf) pure returns(bytes32)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageCaller) CalculateEarnerLeafHash(opts *bind.CallOpts, leaf IPaymentCoordinatorEarnerTreeMerkleLeaf) ([32]byte, error) {
	var out []interface{}
	err := _PaymentCoordinatorStorage.contract.Call(opts, &out, "calculateEarnerLeafHash", leaf)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateEarnerLeafHash is a free data retrieval call binding the contract method 0x149bc872.
//
// Solidity: function calculateEarnerLeafHash((address,bytes32) leaf) pure returns(bytes32)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageSession) CalculateEarnerLeafHash(leaf IPaymentCoordinatorEarnerTreeMerkleLeaf) ([32]byte, error) {
	return _PaymentCoordinatorStorage.Contract.CalculateEarnerLeafHash(&_PaymentCoordinatorStorage.CallOpts, leaf)
}

// CalculateEarnerLeafHash is a free data retrieval call binding the contract method 0x149bc872.
//
// Solidity: function calculateEarnerLeafHash((address,bytes32) leaf) pure returns(bytes32)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageCallerSession) CalculateEarnerLeafHash(leaf IPaymentCoordinatorEarnerTreeMerkleLeaf) ([32]byte, error) {
	return _PaymentCoordinatorStorage.Contract.CalculateEarnerLeafHash(&_PaymentCoordinatorStorage.CallOpts, leaf)
}

// CalculateTokenLeafHash is a free data retrieval call binding the contract method 0xf8cd8448.
//
// Solidity: function calculateTokenLeafHash((address,uint256) leaf) pure returns(bytes32)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageCaller) CalculateTokenLeafHash(opts *bind.CallOpts, leaf IPaymentCoordinatorTokenTreeMerkleLeaf) ([32]byte, error) {
	var out []interface{}
	err := _PaymentCoordinatorStorage.contract.Call(opts, &out, "calculateTokenLeafHash", leaf)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateTokenLeafHash is a free data retrieval call binding the contract method 0xf8cd8448.
//
// Solidity: function calculateTokenLeafHash((address,uint256) leaf) pure returns(bytes32)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageSession) CalculateTokenLeafHash(leaf IPaymentCoordinatorTokenTreeMerkleLeaf) ([32]byte, error) {
	return _PaymentCoordinatorStorage.Contract.CalculateTokenLeafHash(&_PaymentCoordinatorStorage.CallOpts, leaf)
}

// CalculateTokenLeafHash is a free data retrieval call binding the contract method 0xf8cd8448.
//
// Solidity: function calculateTokenLeafHash((address,uint256) leaf) pure returns(bytes32)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageCallerSession) CalculateTokenLeafHash(leaf IPaymentCoordinatorTokenTreeMerkleLeaf) ([32]byte, error) {
	return _PaymentCoordinatorStorage.Contract.CalculateTokenLeafHash(&_PaymentCoordinatorStorage.CallOpts, leaf)
}

// CheckClaim is a free data retrieval call binding the contract method 0x5e9d8348.
//
// Solidity: function checkClaim((uint32,uint32,bytes,(address,bytes32),uint32[],bytes[],(address,uint256)[]) claim) view returns(bool)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageCaller) CheckClaim(opts *bind.CallOpts, claim IPaymentCoordinatorPaymentMerkleClaim) (bool, error) {
	var out []interface{}
	err := _PaymentCoordinatorStorage.contract.Call(opts, &out, "checkClaim", claim)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckClaim is a free data retrieval call binding the contract method 0x5e9d8348.
//
// Solidity: function checkClaim((uint32,uint32,bytes,(address,bytes32),uint32[],bytes[],(address,uint256)[]) claim) view returns(bool)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageSession) CheckClaim(claim IPaymentCoordinatorPaymentMerkleClaim) (bool, error) {
	return _PaymentCoordinatorStorage.Contract.CheckClaim(&_PaymentCoordinatorStorage.CallOpts, claim)
}

// CheckClaim is a free data retrieval call binding the contract method 0x5e9d8348.
//
// Solidity: function checkClaim((uint32,uint32,bytes,(address,bytes32),uint32[],bytes[],(address,uint256)[]) claim) view returns(bool)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageCallerSession) CheckClaim(claim IPaymentCoordinatorPaymentMerkleClaim) (bool, error) {
	return _PaymentCoordinatorStorage.Contract.CheckClaim(&_PaymentCoordinatorStorage.CallOpts, claim)
}

// ClaimerFor is a free data retrieval call binding the contract method 0x2b9f64a4.
//
// Solidity: function claimerFor(address ) view returns(address)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageCaller) ClaimerFor(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _PaymentCoordinatorStorage.contract.Call(opts, &out, "claimerFor", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ClaimerFor is a free data retrieval call binding the contract method 0x2b9f64a4.
//
// Solidity: function claimerFor(address ) view returns(address)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageSession) ClaimerFor(arg0 common.Address) (common.Address, error) {
	return _PaymentCoordinatorStorage.Contract.ClaimerFor(&_PaymentCoordinatorStorage.CallOpts, arg0)
}

// ClaimerFor is a free data retrieval call binding the contract method 0x2b9f64a4.
//
// Solidity: function claimerFor(address ) view returns(address)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageCallerSession) ClaimerFor(arg0 common.Address) (common.Address, error) {
	return _PaymentCoordinatorStorage.Contract.ClaimerFor(&_PaymentCoordinatorStorage.CallOpts, arg0)
}

// CumulativeClaimed is a free data retrieval call binding the contract method 0x865c6953.
//
// Solidity: function cumulativeClaimed(address , address ) view returns(uint256)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageCaller) CumulativeClaimed(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PaymentCoordinatorStorage.contract.Call(opts, &out, "cumulativeClaimed", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CumulativeClaimed is a free data retrieval call binding the contract method 0x865c6953.
//
// Solidity: function cumulativeClaimed(address , address ) view returns(uint256)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageSession) CumulativeClaimed(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _PaymentCoordinatorStorage.Contract.CumulativeClaimed(&_PaymentCoordinatorStorage.CallOpts, arg0, arg1)
}

// CumulativeClaimed is a free data retrieval call binding the contract method 0x865c6953.
//
// Solidity: function cumulativeClaimed(address , address ) view returns(uint256)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageCallerSession) CumulativeClaimed(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _PaymentCoordinatorStorage.Contract.CumulativeClaimed(&_PaymentCoordinatorStorage.CallOpts, arg0, arg1)
}

// CurrPaymentCalculationEndTimestamp is a free data retrieval call binding the contract method 0x67ef8585.
//
// Solidity: function currPaymentCalculationEndTimestamp() view returns(uint32)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageCaller) CurrPaymentCalculationEndTimestamp(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _PaymentCoordinatorStorage.contract.Call(opts, &out, "currPaymentCalculationEndTimestamp")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// CurrPaymentCalculationEndTimestamp is a free data retrieval call binding the contract method 0x67ef8585.
//
// Solidity: function currPaymentCalculationEndTimestamp() view returns(uint32)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageSession) CurrPaymentCalculationEndTimestamp() (uint32, error) {
	return _PaymentCoordinatorStorage.Contract.CurrPaymentCalculationEndTimestamp(&_PaymentCoordinatorStorage.CallOpts)
}

// CurrPaymentCalculationEndTimestamp is a free data retrieval call binding the contract method 0x67ef8585.
//
// Solidity: function currPaymentCalculationEndTimestamp() view returns(uint32)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageCallerSession) CurrPaymentCalculationEndTimestamp() (uint32, error) {
	return _PaymentCoordinatorStorage.Contract.CurrPaymentCalculationEndTimestamp(&_PaymentCoordinatorStorage.CallOpts)
}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageCaller) DelegationManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PaymentCoordinatorStorage.contract.Call(opts, &out, "delegationManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageSession) DelegationManager() (common.Address, error) {
	return _PaymentCoordinatorStorage.Contract.DelegationManager(&_PaymentCoordinatorStorage.CallOpts)
}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageCallerSession) DelegationManager() (common.Address, error) {
	return _PaymentCoordinatorStorage.Contract.DelegationManager(&_PaymentCoordinatorStorage.CallOpts)
}

// DistributionRoots is a free data retrieval call binding the contract method 0x5971b3f8.
//
// Solidity: function distributionRoots(uint256 ) view returns(bytes32 root, uint32 paymentCalculationEndTimestamp, uint32 activatedAt)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageCaller) DistributionRoots(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Root                           [32]byte
	PaymentCalculationEndTimestamp uint32
	ActivatedAt                    uint32
}, error) {
	var out []interface{}
	err := _PaymentCoordinatorStorage.contract.Call(opts, &out, "distributionRoots", arg0)

	outstruct := new(struct {
		Root                           [32]byte
		PaymentCalculationEndTimestamp uint32
		ActivatedAt                    uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Root = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.PaymentCalculationEndTimestamp = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.ActivatedAt = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

// DistributionRoots is a free data retrieval call binding the contract method 0x5971b3f8.
//
// Solidity: function distributionRoots(uint256 ) view returns(bytes32 root, uint32 paymentCalculationEndTimestamp, uint32 activatedAt)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageSession) DistributionRoots(arg0 *big.Int) (struct {
	Root                           [32]byte
	PaymentCalculationEndTimestamp uint32
	ActivatedAt                    uint32
}, error) {
	return _PaymentCoordinatorStorage.Contract.DistributionRoots(&_PaymentCoordinatorStorage.CallOpts, arg0)
}

// DistributionRoots is a free data retrieval call binding the contract method 0x5971b3f8.
//
// Solidity: function distributionRoots(uint256 ) view returns(bytes32 root, uint32 paymentCalculationEndTimestamp, uint32 activatedAt)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageCallerSession) DistributionRoots(arg0 *big.Int) (struct {
	Root                           [32]byte
	PaymentCalculationEndTimestamp uint32
	ActivatedAt                    uint32
}, error) {
	return _PaymentCoordinatorStorage.Contract.DistributionRoots(&_PaymentCoordinatorStorage.CallOpts, arg0)
}

// GetDistributionRootsLength is a free data retrieval call binding the contract method 0x7b8f8b05.
//
// Solidity: function getDistributionRootsLength() view returns(uint256)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageCaller) GetDistributionRootsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PaymentCoordinatorStorage.contract.Call(opts, &out, "getDistributionRootsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDistributionRootsLength is a free data retrieval call binding the contract method 0x7b8f8b05.
//
// Solidity: function getDistributionRootsLength() view returns(uint256)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageSession) GetDistributionRootsLength() (*big.Int, error) {
	return _PaymentCoordinatorStorage.Contract.GetDistributionRootsLength(&_PaymentCoordinatorStorage.CallOpts)
}

// GetDistributionRootsLength is a free data retrieval call binding the contract method 0x7b8f8b05.
//
// Solidity: function getDistributionRootsLength() view returns(uint256)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageCallerSession) GetDistributionRootsLength() (*big.Int, error) {
	return _PaymentCoordinatorStorage.Contract.GetDistributionRootsLength(&_PaymentCoordinatorStorage.CallOpts)
}

// GetRootIndexFromHash is a free data retrieval call binding the contract method 0xe810ce21.
//
// Solidity: function getRootIndexFromHash(bytes32 rootHash) view returns(uint32)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageCaller) GetRootIndexFromHash(opts *bind.CallOpts, rootHash [32]byte) (uint32, error) {
	var out []interface{}
	err := _PaymentCoordinatorStorage.contract.Call(opts, &out, "getRootIndexFromHash", rootHash)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetRootIndexFromHash is a free data retrieval call binding the contract method 0xe810ce21.
//
// Solidity: function getRootIndexFromHash(bytes32 rootHash) view returns(uint32)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageSession) GetRootIndexFromHash(rootHash [32]byte) (uint32, error) {
	return _PaymentCoordinatorStorage.Contract.GetRootIndexFromHash(&_PaymentCoordinatorStorage.CallOpts, rootHash)
}

// GetRootIndexFromHash is a free data retrieval call binding the contract method 0xe810ce21.
//
// Solidity: function getRootIndexFromHash(bytes32 rootHash) view returns(uint32)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageCallerSession) GetRootIndexFromHash(rootHash [32]byte) (uint32, error) {
	return _PaymentCoordinatorStorage.Contract.GetRootIndexFromHash(&_PaymentCoordinatorStorage.CallOpts, rootHash)
}

// GlobalOperatorCommissionBips is a free data retrieval call binding the contract method 0x092db007.
//
// Solidity: function globalOperatorCommissionBips() view returns(uint16)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageCaller) GlobalOperatorCommissionBips(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _PaymentCoordinatorStorage.contract.Call(opts, &out, "globalOperatorCommissionBips")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GlobalOperatorCommissionBips is a free data retrieval call binding the contract method 0x092db007.
//
// Solidity: function globalOperatorCommissionBips() view returns(uint16)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageSession) GlobalOperatorCommissionBips() (uint16, error) {
	return _PaymentCoordinatorStorage.Contract.GlobalOperatorCommissionBips(&_PaymentCoordinatorStorage.CallOpts)
}

// GlobalOperatorCommissionBips is a free data retrieval call binding the contract method 0x092db007.
//
// Solidity: function globalOperatorCommissionBips() view returns(uint16)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageCallerSession) GlobalOperatorCommissionBips() (uint16, error) {
	return _PaymentCoordinatorStorage.Contract.GlobalOperatorCommissionBips(&_PaymentCoordinatorStorage.CallOpts)
}

// IsPayAllForRangeSubmitter is a free data retrieval call binding the contract method 0x146cd61d.
//
// Solidity: function isPayAllForRangeSubmitter(address ) view returns(bool)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageCaller) IsPayAllForRangeSubmitter(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _PaymentCoordinatorStorage.contract.Call(opts, &out, "isPayAllForRangeSubmitter", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPayAllForRangeSubmitter is a free data retrieval call binding the contract method 0x146cd61d.
//
// Solidity: function isPayAllForRangeSubmitter(address ) view returns(bool)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageSession) IsPayAllForRangeSubmitter(arg0 common.Address) (bool, error) {
	return _PaymentCoordinatorStorage.Contract.IsPayAllForRangeSubmitter(&_PaymentCoordinatorStorage.CallOpts, arg0)
}

// IsPayAllForRangeSubmitter is a free data retrieval call binding the contract method 0x146cd61d.
//
// Solidity: function isPayAllForRangeSubmitter(address ) view returns(bool)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageCallerSession) IsPayAllForRangeSubmitter(arg0 common.Address) (bool, error) {
	return _PaymentCoordinatorStorage.Contract.IsPayAllForRangeSubmitter(&_PaymentCoordinatorStorage.CallOpts, arg0)
}

// IsRangePaymentForAllHash is a free data retrieval call binding the contract method 0x73f2fbea.
//
// Solidity: function isRangePaymentForAllHash(address , bytes32 ) view returns(bool)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageCaller) IsRangePaymentForAllHash(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (bool, error) {
	var out []interface{}
	err := _PaymentCoordinatorStorage.contract.Call(opts, &out, "isRangePaymentForAllHash", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRangePaymentForAllHash is a free data retrieval call binding the contract method 0x73f2fbea.
//
// Solidity: function isRangePaymentForAllHash(address , bytes32 ) view returns(bool)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageSession) IsRangePaymentForAllHash(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _PaymentCoordinatorStorage.Contract.IsRangePaymentForAllHash(&_PaymentCoordinatorStorage.CallOpts, arg0, arg1)
}

// IsRangePaymentForAllHash is a free data retrieval call binding the contract method 0x73f2fbea.
//
// Solidity: function isRangePaymentForAllHash(address , bytes32 ) view returns(bool)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageCallerSession) IsRangePaymentForAllHash(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _PaymentCoordinatorStorage.Contract.IsRangePaymentForAllHash(&_PaymentCoordinatorStorage.CallOpts, arg0, arg1)
}

// IsRangePaymentHash is a free data retrieval call binding the contract method 0xc8371b46.
//
// Solidity: function isRangePaymentHash(address , bytes32 ) view returns(bool)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageCaller) IsRangePaymentHash(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (bool, error) {
	var out []interface{}
	err := _PaymentCoordinatorStorage.contract.Call(opts, &out, "isRangePaymentHash", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRangePaymentHash is a free data retrieval call binding the contract method 0xc8371b46.
//
// Solidity: function isRangePaymentHash(address , bytes32 ) view returns(bool)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageSession) IsRangePaymentHash(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _PaymentCoordinatorStorage.Contract.IsRangePaymentHash(&_PaymentCoordinatorStorage.CallOpts, arg0, arg1)
}

// IsRangePaymentHash is a free data retrieval call binding the contract method 0xc8371b46.
//
// Solidity: function isRangePaymentHash(address , bytes32 ) view returns(bool)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageCallerSession) IsRangePaymentHash(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _PaymentCoordinatorStorage.Contract.IsRangePaymentHash(&_PaymentCoordinatorStorage.CallOpts, arg0, arg1)
}

// OperatorCommissionBips is a free data retrieval call binding the contract method 0x22f19a64.
//
// Solidity: function operatorCommissionBips(address operator, address avs) view returns(uint16)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageCaller) OperatorCommissionBips(opts *bind.CallOpts, operator common.Address, avs common.Address) (uint16, error) {
	var out []interface{}
	err := _PaymentCoordinatorStorage.contract.Call(opts, &out, "operatorCommissionBips", operator, avs)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// OperatorCommissionBips is a free data retrieval call binding the contract method 0x22f19a64.
//
// Solidity: function operatorCommissionBips(address operator, address avs) view returns(uint16)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageSession) OperatorCommissionBips(operator common.Address, avs common.Address) (uint16, error) {
	return _PaymentCoordinatorStorage.Contract.OperatorCommissionBips(&_PaymentCoordinatorStorage.CallOpts, operator, avs)
}

// OperatorCommissionBips is a free data retrieval call binding the contract method 0x22f19a64.
//
// Solidity: function operatorCommissionBips(address operator, address avs) view returns(uint16)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageCallerSession) OperatorCommissionBips(operator common.Address, avs common.Address) (uint16, error) {
	return _PaymentCoordinatorStorage.Contract.OperatorCommissionBips(&_PaymentCoordinatorStorage.CallOpts, operator, avs)
}

// PaymentNonce is a free data retrieval call binding the contract method 0xd588cefa.
//
// Solidity: function paymentNonce(address ) view returns(uint256)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageCaller) PaymentNonce(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PaymentCoordinatorStorage.contract.Call(opts, &out, "paymentNonce", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PaymentNonce is a free data retrieval call binding the contract method 0xd588cefa.
//
// Solidity: function paymentNonce(address ) view returns(uint256)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageSession) PaymentNonce(arg0 common.Address) (*big.Int, error) {
	return _PaymentCoordinatorStorage.Contract.PaymentNonce(&_PaymentCoordinatorStorage.CallOpts, arg0)
}

// PaymentNonce is a free data retrieval call binding the contract method 0xd588cefa.
//
// Solidity: function paymentNonce(address ) view returns(uint256)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageCallerSession) PaymentNonce(arg0 common.Address) (*big.Int, error) {
	return _PaymentCoordinatorStorage.Contract.PaymentNonce(&_PaymentCoordinatorStorage.CallOpts, arg0)
}

// PaymentUpdater is a free data retrieval call binding the contract method 0x66d3b16b.
//
// Solidity: function paymentUpdater() view returns(address)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageCaller) PaymentUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PaymentCoordinatorStorage.contract.Call(opts, &out, "paymentUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PaymentUpdater is a free data retrieval call binding the contract method 0x66d3b16b.
//
// Solidity: function paymentUpdater() view returns(address)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageSession) PaymentUpdater() (common.Address, error) {
	return _PaymentCoordinatorStorage.Contract.PaymentUpdater(&_PaymentCoordinatorStorage.CallOpts)
}

// PaymentUpdater is a free data retrieval call binding the contract method 0x66d3b16b.
//
// Solidity: function paymentUpdater() view returns(address)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageCallerSession) PaymentUpdater() (common.Address, error) {
	return _PaymentCoordinatorStorage.Contract.PaymentUpdater(&_PaymentCoordinatorStorage.CallOpts)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageCaller) StrategyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PaymentCoordinatorStorage.contract.Call(opts, &out, "strategyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageSession) StrategyManager() (common.Address, error) {
	return _PaymentCoordinatorStorage.Contract.StrategyManager(&_PaymentCoordinatorStorage.CallOpts)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageCallerSession) StrategyManager() (common.Address, error) {
	return _PaymentCoordinatorStorage.Contract.StrategyManager(&_PaymentCoordinatorStorage.CallOpts)
}

// PayAllForRange is a paid mutator transaction binding the contract method 0xb5715284.
//
// Solidity: function payAllForRange(((address,uint96)[],address,uint256,uint32,uint32)[] rangePayment) returns()
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageTransactor) PayAllForRange(opts *bind.TransactOpts, rangePayment []IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _PaymentCoordinatorStorage.contract.Transact(opts, "payAllForRange", rangePayment)
}

// PayAllForRange is a paid mutator transaction binding the contract method 0xb5715284.
//
// Solidity: function payAllForRange(((address,uint96)[],address,uint256,uint32,uint32)[] rangePayment) returns()
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageSession) PayAllForRange(rangePayment []IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _PaymentCoordinatorStorage.Contract.PayAllForRange(&_PaymentCoordinatorStorage.TransactOpts, rangePayment)
}

// PayAllForRange is a paid mutator transaction binding the contract method 0xb5715284.
//
// Solidity: function payAllForRange(((address,uint96)[],address,uint256,uint32,uint32)[] rangePayment) returns()
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageTransactorSession) PayAllForRange(rangePayment []IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _PaymentCoordinatorStorage.Contract.PayAllForRange(&_PaymentCoordinatorStorage.TransactOpts, rangePayment)
}

// PayForRange is a paid mutator transaction binding the contract method 0x1b445516.
//
// Solidity: function payForRange(((address,uint96)[],address,uint256,uint32,uint32)[] rangePayments) returns()
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageTransactor) PayForRange(opts *bind.TransactOpts, rangePayments []IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _PaymentCoordinatorStorage.contract.Transact(opts, "payForRange", rangePayments)
}

// PayForRange is a paid mutator transaction binding the contract method 0x1b445516.
//
// Solidity: function payForRange(((address,uint96)[],address,uint256,uint32,uint32)[] rangePayments) returns()
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageSession) PayForRange(rangePayments []IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _PaymentCoordinatorStorage.Contract.PayForRange(&_PaymentCoordinatorStorage.TransactOpts, rangePayments)
}

// PayForRange is a paid mutator transaction binding the contract method 0x1b445516.
//
// Solidity: function payForRange(((address,uint96)[],address,uint256,uint32,uint32)[] rangePayments) returns()
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageTransactorSession) PayForRange(rangePayments []IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _PaymentCoordinatorStorage.Contract.PayForRange(&_PaymentCoordinatorStorage.TransactOpts, rangePayments)
}

// ProcessClaim is a paid mutator transaction binding the contract method 0x3ccc861d.
//
// Solidity: function processClaim((uint32,uint32,bytes,(address,bytes32),uint32[],bytes[],(address,uint256)[]) claim, address recipient) returns()
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageTransactor) ProcessClaim(opts *bind.TransactOpts, claim IPaymentCoordinatorPaymentMerkleClaim, recipient common.Address) (*types.Transaction, error) {
	return _PaymentCoordinatorStorage.contract.Transact(opts, "processClaim", claim, recipient)
}

// ProcessClaim is a paid mutator transaction binding the contract method 0x3ccc861d.
//
// Solidity: function processClaim((uint32,uint32,bytes,(address,bytes32),uint32[],bytes[],(address,uint256)[]) claim, address recipient) returns()
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageSession) ProcessClaim(claim IPaymentCoordinatorPaymentMerkleClaim, recipient common.Address) (*types.Transaction, error) {
	return _PaymentCoordinatorStorage.Contract.ProcessClaim(&_PaymentCoordinatorStorage.TransactOpts, claim, recipient)
}

// ProcessClaim is a paid mutator transaction binding the contract method 0x3ccc861d.
//
// Solidity: function processClaim((uint32,uint32,bytes,(address,bytes32),uint32[],bytes[],(address,uint256)[]) claim, address recipient) returns()
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageTransactorSession) ProcessClaim(claim IPaymentCoordinatorPaymentMerkleClaim, recipient common.Address) (*types.Transaction, error) {
	return _PaymentCoordinatorStorage.Contract.ProcessClaim(&_PaymentCoordinatorStorage.TransactOpts, claim, recipient)
}

// SetActivationDelay is a paid mutator transaction binding the contract method 0x58baaa3e.
//
// Solidity: function setActivationDelay(uint32 _activationDelay) returns()
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageTransactor) SetActivationDelay(opts *bind.TransactOpts, _activationDelay uint32) (*types.Transaction, error) {
	return _PaymentCoordinatorStorage.contract.Transact(opts, "setActivationDelay", _activationDelay)
}

// SetActivationDelay is a paid mutator transaction binding the contract method 0x58baaa3e.
//
// Solidity: function setActivationDelay(uint32 _activationDelay) returns()
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageSession) SetActivationDelay(_activationDelay uint32) (*types.Transaction, error) {
	return _PaymentCoordinatorStorage.Contract.SetActivationDelay(&_PaymentCoordinatorStorage.TransactOpts, _activationDelay)
}

// SetActivationDelay is a paid mutator transaction binding the contract method 0x58baaa3e.
//
// Solidity: function setActivationDelay(uint32 _activationDelay) returns()
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageTransactorSession) SetActivationDelay(_activationDelay uint32) (*types.Transaction, error) {
	return _PaymentCoordinatorStorage.Contract.SetActivationDelay(&_PaymentCoordinatorStorage.TransactOpts, _activationDelay)
}

// SetClaimerFor is a paid mutator transaction binding the contract method 0xa0169ddd.
//
// Solidity: function setClaimerFor(address claimer) returns()
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageTransactor) SetClaimerFor(opts *bind.TransactOpts, claimer common.Address) (*types.Transaction, error) {
	return _PaymentCoordinatorStorage.contract.Transact(opts, "setClaimerFor", claimer)
}

// SetClaimerFor is a paid mutator transaction binding the contract method 0xa0169ddd.
//
// Solidity: function setClaimerFor(address claimer) returns()
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageSession) SetClaimerFor(claimer common.Address) (*types.Transaction, error) {
	return _PaymentCoordinatorStorage.Contract.SetClaimerFor(&_PaymentCoordinatorStorage.TransactOpts, claimer)
}

// SetClaimerFor is a paid mutator transaction binding the contract method 0xa0169ddd.
//
// Solidity: function setClaimerFor(address claimer) returns()
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageTransactorSession) SetClaimerFor(claimer common.Address) (*types.Transaction, error) {
	return _PaymentCoordinatorStorage.Contract.SetClaimerFor(&_PaymentCoordinatorStorage.TransactOpts, claimer)
}

// SetGlobalOperatorCommission is a paid mutator transaction binding the contract method 0xe221b245.
//
// Solidity: function setGlobalOperatorCommission(uint16 _globalCommissionBips) returns()
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageTransactor) SetGlobalOperatorCommission(opts *bind.TransactOpts, _globalCommissionBips uint16) (*types.Transaction, error) {
	return _PaymentCoordinatorStorage.contract.Transact(opts, "setGlobalOperatorCommission", _globalCommissionBips)
}

// SetGlobalOperatorCommission is a paid mutator transaction binding the contract method 0xe221b245.
//
// Solidity: function setGlobalOperatorCommission(uint16 _globalCommissionBips) returns()
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageSession) SetGlobalOperatorCommission(_globalCommissionBips uint16) (*types.Transaction, error) {
	return _PaymentCoordinatorStorage.Contract.SetGlobalOperatorCommission(&_PaymentCoordinatorStorage.TransactOpts, _globalCommissionBips)
}

// SetGlobalOperatorCommission is a paid mutator transaction binding the contract method 0xe221b245.
//
// Solidity: function setGlobalOperatorCommission(uint16 _globalCommissionBips) returns()
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageTransactorSession) SetGlobalOperatorCommission(_globalCommissionBips uint16) (*types.Transaction, error) {
	return _PaymentCoordinatorStorage.Contract.SetGlobalOperatorCommission(&_PaymentCoordinatorStorage.TransactOpts, _globalCommissionBips)
}

// SetPayAllForRangeSubmitter is a paid mutator transaction binding the contract method 0xec1680de.
//
// Solidity: function setPayAllForRangeSubmitter(address _submitter, bool _newValue) returns()
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageTransactor) SetPayAllForRangeSubmitter(opts *bind.TransactOpts, _submitter common.Address, _newValue bool) (*types.Transaction, error) {
	return _PaymentCoordinatorStorage.contract.Transact(opts, "setPayAllForRangeSubmitter", _submitter, _newValue)
}

// SetPayAllForRangeSubmitter is a paid mutator transaction binding the contract method 0xec1680de.
//
// Solidity: function setPayAllForRangeSubmitter(address _submitter, bool _newValue) returns()
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageSession) SetPayAllForRangeSubmitter(_submitter common.Address, _newValue bool) (*types.Transaction, error) {
	return _PaymentCoordinatorStorage.Contract.SetPayAllForRangeSubmitter(&_PaymentCoordinatorStorage.TransactOpts, _submitter, _newValue)
}

// SetPayAllForRangeSubmitter is a paid mutator transaction binding the contract method 0xec1680de.
//
// Solidity: function setPayAllForRangeSubmitter(address _submitter, bool _newValue) returns()
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageTransactorSession) SetPayAllForRangeSubmitter(_submitter common.Address, _newValue bool) (*types.Transaction, error) {
	return _PaymentCoordinatorStorage.Contract.SetPayAllForRangeSubmitter(&_PaymentCoordinatorStorage.TransactOpts, _submitter, _newValue)
}

// SetPaymentUpdater is a paid mutator transaction binding the contract method 0x18190f53.
//
// Solidity: function setPaymentUpdater(address _paymentUpdater) returns()
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageTransactor) SetPaymentUpdater(opts *bind.TransactOpts, _paymentUpdater common.Address) (*types.Transaction, error) {
	return _PaymentCoordinatorStorage.contract.Transact(opts, "setPaymentUpdater", _paymentUpdater)
}

// SetPaymentUpdater is a paid mutator transaction binding the contract method 0x18190f53.
//
// Solidity: function setPaymentUpdater(address _paymentUpdater) returns()
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageSession) SetPaymentUpdater(_paymentUpdater common.Address) (*types.Transaction, error) {
	return _PaymentCoordinatorStorage.Contract.SetPaymentUpdater(&_PaymentCoordinatorStorage.TransactOpts, _paymentUpdater)
}

// SetPaymentUpdater is a paid mutator transaction binding the contract method 0x18190f53.
//
// Solidity: function setPaymentUpdater(address _paymentUpdater) returns()
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageTransactorSession) SetPaymentUpdater(_paymentUpdater common.Address) (*types.Transaction, error) {
	return _PaymentCoordinatorStorage.Contract.SetPaymentUpdater(&_PaymentCoordinatorStorage.TransactOpts, _paymentUpdater)
}

// SubmitRoot is a paid mutator transaction binding the contract method 0x3efe1db6.
//
// Solidity: function submitRoot(bytes32 root, uint32 paymentCalculationEndTimestamp) returns()
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageTransactor) SubmitRoot(opts *bind.TransactOpts, root [32]byte, paymentCalculationEndTimestamp uint32) (*types.Transaction, error) {
	return _PaymentCoordinatorStorage.contract.Transact(opts, "submitRoot", root, paymentCalculationEndTimestamp)
}

// SubmitRoot is a paid mutator transaction binding the contract method 0x3efe1db6.
//
// Solidity: function submitRoot(bytes32 root, uint32 paymentCalculationEndTimestamp) returns()
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageSession) SubmitRoot(root [32]byte, paymentCalculationEndTimestamp uint32) (*types.Transaction, error) {
	return _PaymentCoordinatorStorage.Contract.SubmitRoot(&_PaymentCoordinatorStorage.TransactOpts, root, paymentCalculationEndTimestamp)
}

// SubmitRoot is a paid mutator transaction binding the contract method 0x3efe1db6.
//
// Solidity: function submitRoot(bytes32 root, uint32 paymentCalculationEndTimestamp) returns()
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageTransactorSession) SubmitRoot(root [32]byte, paymentCalculationEndTimestamp uint32) (*types.Transaction, error) {
	return _PaymentCoordinatorStorage.Contract.SubmitRoot(&_PaymentCoordinatorStorage.TransactOpts, root, paymentCalculationEndTimestamp)
}

// PaymentCoordinatorStorageActivationDelaySetIterator is returned from FilterActivationDelaySet and is used to iterate over the raw logs and unpacked data for ActivationDelaySet events raised by the PaymentCoordinatorStorage contract.
type PaymentCoordinatorStorageActivationDelaySetIterator struct {
	Event *PaymentCoordinatorStorageActivationDelaySet // Event containing the contract specifics and raw log

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
func (it *PaymentCoordinatorStorageActivationDelaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentCoordinatorStorageActivationDelaySet)
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
		it.Event = new(PaymentCoordinatorStorageActivationDelaySet)
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
func (it *PaymentCoordinatorStorageActivationDelaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentCoordinatorStorageActivationDelaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentCoordinatorStorageActivationDelaySet represents a ActivationDelaySet event raised by the PaymentCoordinatorStorage contract.
type PaymentCoordinatorStorageActivationDelaySet struct {
	OldActivationDelay uint32
	NewActivationDelay uint32
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterActivationDelaySet is a free log retrieval operation binding the contract event 0xaf557c6c02c208794817a705609cfa935f827312a1adfdd26494b6b95dd2b4b3.
//
// Solidity: event ActivationDelaySet(uint32 oldActivationDelay, uint32 newActivationDelay)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageFilterer) FilterActivationDelaySet(opts *bind.FilterOpts) (*PaymentCoordinatorStorageActivationDelaySetIterator, error) {

	logs, sub, err := _PaymentCoordinatorStorage.contract.FilterLogs(opts, "ActivationDelaySet")
	if err != nil {
		return nil, err
	}
	return &PaymentCoordinatorStorageActivationDelaySetIterator{contract: _PaymentCoordinatorStorage.contract, event: "ActivationDelaySet", logs: logs, sub: sub}, nil
}

// WatchActivationDelaySet is a free log subscription operation binding the contract event 0xaf557c6c02c208794817a705609cfa935f827312a1adfdd26494b6b95dd2b4b3.
//
// Solidity: event ActivationDelaySet(uint32 oldActivationDelay, uint32 newActivationDelay)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageFilterer) WatchActivationDelaySet(opts *bind.WatchOpts, sink chan<- *PaymentCoordinatorStorageActivationDelaySet) (event.Subscription, error) {

	logs, sub, err := _PaymentCoordinatorStorage.contract.WatchLogs(opts, "ActivationDelaySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentCoordinatorStorageActivationDelaySet)
				if err := _PaymentCoordinatorStorage.contract.UnpackLog(event, "ActivationDelaySet", log); err != nil {
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
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageFilterer) ParseActivationDelaySet(log types.Log) (*PaymentCoordinatorStorageActivationDelaySet, error) {
	event := new(PaymentCoordinatorStorageActivationDelaySet)
	if err := _PaymentCoordinatorStorage.contract.UnpackLog(event, "ActivationDelaySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentCoordinatorStorageClaimerForSetIterator is returned from FilterClaimerForSet and is used to iterate over the raw logs and unpacked data for ClaimerForSet events raised by the PaymentCoordinatorStorage contract.
type PaymentCoordinatorStorageClaimerForSetIterator struct {
	Event *PaymentCoordinatorStorageClaimerForSet // Event containing the contract specifics and raw log

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
func (it *PaymentCoordinatorStorageClaimerForSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentCoordinatorStorageClaimerForSet)
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
		it.Event = new(PaymentCoordinatorStorageClaimerForSet)
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
func (it *PaymentCoordinatorStorageClaimerForSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentCoordinatorStorageClaimerForSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentCoordinatorStorageClaimerForSet represents a ClaimerForSet event raised by the PaymentCoordinatorStorage contract.
type PaymentCoordinatorStorageClaimerForSet struct {
	Earner     common.Address
	OldClaimer common.Address
	Claimer    common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterClaimerForSet is a free log retrieval operation binding the contract event 0xbab947934d42e0ad206f25c9cab18b5bb6ae144acfb00f40b4e3aa59590ca312.
//
// Solidity: event ClaimerForSet(address indexed earner, address indexed oldClaimer, address indexed claimer)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageFilterer) FilterClaimerForSet(opts *bind.FilterOpts, earner []common.Address, oldClaimer []common.Address, claimer []common.Address) (*PaymentCoordinatorStorageClaimerForSetIterator, error) {

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

	logs, sub, err := _PaymentCoordinatorStorage.contract.FilterLogs(opts, "ClaimerForSet", earnerRule, oldClaimerRule, claimerRule)
	if err != nil {
		return nil, err
	}
	return &PaymentCoordinatorStorageClaimerForSetIterator{contract: _PaymentCoordinatorStorage.contract, event: "ClaimerForSet", logs: logs, sub: sub}, nil
}

// WatchClaimerForSet is a free log subscription operation binding the contract event 0xbab947934d42e0ad206f25c9cab18b5bb6ae144acfb00f40b4e3aa59590ca312.
//
// Solidity: event ClaimerForSet(address indexed earner, address indexed oldClaimer, address indexed claimer)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageFilterer) WatchClaimerForSet(opts *bind.WatchOpts, sink chan<- *PaymentCoordinatorStorageClaimerForSet, earner []common.Address, oldClaimer []common.Address, claimer []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PaymentCoordinatorStorage.contract.WatchLogs(opts, "ClaimerForSet", earnerRule, oldClaimerRule, claimerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentCoordinatorStorageClaimerForSet)
				if err := _PaymentCoordinatorStorage.contract.UnpackLog(event, "ClaimerForSet", log); err != nil {
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
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageFilterer) ParseClaimerForSet(log types.Log) (*PaymentCoordinatorStorageClaimerForSet, error) {
	event := new(PaymentCoordinatorStorageClaimerForSet)
	if err := _PaymentCoordinatorStorage.contract.UnpackLog(event, "ClaimerForSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentCoordinatorStorageDistributionRootSubmittedIterator is returned from FilterDistributionRootSubmitted and is used to iterate over the raw logs and unpacked data for DistributionRootSubmitted events raised by the PaymentCoordinatorStorage contract.
type PaymentCoordinatorStorageDistributionRootSubmittedIterator struct {
	Event *PaymentCoordinatorStorageDistributionRootSubmitted // Event containing the contract specifics and raw log

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
func (it *PaymentCoordinatorStorageDistributionRootSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentCoordinatorStorageDistributionRootSubmitted)
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
		it.Event = new(PaymentCoordinatorStorageDistributionRootSubmitted)
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
func (it *PaymentCoordinatorStorageDistributionRootSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentCoordinatorStorageDistributionRootSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentCoordinatorStorageDistributionRootSubmitted represents a DistributionRootSubmitted event raised by the PaymentCoordinatorStorage contract.
type PaymentCoordinatorStorageDistributionRootSubmitted struct {
	RootIndex                      uint32
	Root                           [32]byte
	PaymentCalculationEndTimestamp uint32
	ActivatedAt                    uint32
	Raw                            types.Log // Blockchain specific contextual infos
}

// FilterDistributionRootSubmitted is a free log retrieval operation binding the contract event 0xecd866c3c158fa00bf34d803d5f6023000b57080bcb48af004c2b4b46b3afd08.
//
// Solidity: event DistributionRootSubmitted(uint32 indexed rootIndex, bytes32 indexed root, uint32 indexed paymentCalculationEndTimestamp, uint32 activatedAt)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageFilterer) FilterDistributionRootSubmitted(opts *bind.FilterOpts, rootIndex []uint32, root [][32]byte, paymentCalculationEndTimestamp []uint32) (*PaymentCoordinatorStorageDistributionRootSubmittedIterator, error) {

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

	logs, sub, err := _PaymentCoordinatorStorage.contract.FilterLogs(opts, "DistributionRootSubmitted", rootIndexRule, rootRule, paymentCalculationEndTimestampRule)
	if err != nil {
		return nil, err
	}
	return &PaymentCoordinatorStorageDistributionRootSubmittedIterator{contract: _PaymentCoordinatorStorage.contract, event: "DistributionRootSubmitted", logs: logs, sub: sub}, nil
}

// WatchDistributionRootSubmitted is a free log subscription operation binding the contract event 0xecd866c3c158fa00bf34d803d5f6023000b57080bcb48af004c2b4b46b3afd08.
//
// Solidity: event DistributionRootSubmitted(uint32 indexed rootIndex, bytes32 indexed root, uint32 indexed paymentCalculationEndTimestamp, uint32 activatedAt)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageFilterer) WatchDistributionRootSubmitted(opts *bind.WatchOpts, sink chan<- *PaymentCoordinatorStorageDistributionRootSubmitted, rootIndex []uint32, root [][32]byte, paymentCalculationEndTimestamp []uint32) (event.Subscription, error) {

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

	logs, sub, err := _PaymentCoordinatorStorage.contract.WatchLogs(opts, "DistributionRootSubmitted", rootIndexRule, rootRule, paymentCalculationEndTimestampRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentCoordinatorStorageDistributionRootSubmitted)
				if err := _PaymentCoordinatorStorage.contract.UnpackLog(event, "DistributionRootSubmitted", log); err != nil {
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
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageFilterer) ParseDistributionRootSubmitted(log types.Log) (*PaymentCoordinatorStorageDistributionRootSubmitted, error) {
	event := new(PaymentCoordinatorStorageDistributionRootSubmitted)
	if err := _PaymentCoordinatorStorage.contract.UnpackLog(event, "DistributionRootSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentCoordinatorStorageGlobalCommissionBipsSetIterator is returned from FilterGlobalCommissionBipsSet and is used to iterate over the raw logs and unpacked data for GlobalCommissionBipsSet events raised by the PaymentCoordinatorStorage contract.
type PaymentCoordinatorStorageGlobalCommissionBipsSetIterator struct {
	Event *PaymentCoordinatorStorageGlobalCommissionBipsSet // Event containing the contract specifics and raw log

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
func (it *PaymentCoordinatorStorageGlobalCommissionBipsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentCoordinatorStorageGlobalCommissionBipsSet)
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
		it.Event = new(PaymentCoordinatorStorageGlobalCommissionBipsSet)
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
func (it *PaymentCoordinatorStorageGlobalCommissionBipsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentCoordinatorStorageGlobalCommissionBipsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentCoordinatorStorageGlobalCommissionBipsSet represents a GlobalCommissionBipsSet event raised by the PaymentCoordinatorStorage contract.
type PaymentCoordinatorStorageGlobalCommissionBipsSet struct {
	OldGlobalCommissionBips uint16
	NewGlobalCommissionBips uint16
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterGlobalCommissionBipsSet is a free log retrieval operation binding the contract event 0x8cdc428b0431b82d1619763f443a48197db344ba96905f3949643acd1c863a06.
//
// Solidity: event GlobalCommissionBipsSet(uint16 oldGlobalCommissionBips, uint16 newGlobalCommissionBips)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageFilterer) FilterGlobalCommissionBipsSet(opts *bind.FilterOpts) (*PaymentCoordinatorStorageGlobalCommissionBipsSetIterator, error) {

	logs, sub, err := _PaymentCoordinatorStorage.contract.FilterLogs(opts, "GlobalCommissionBipsSet")
	if err != nil {
		return nil, err
	}
	return &PaymentCoordinatorStorageGlobalCommissionBipsSetIterator{contract: _PaymentCoordinatorStorage.contract, event: "GlobalCommissionBipsSet", logs: logs, sub: sub}, nil
}

// WatchGlobalCommissionBipsSet is a free log subscription operation binding the contract event 0x8cdc428b0431b82d1619763f443a48197db344ba96905f3949643acd1c863a06.
//
// Solidity: event GlobalCommissionBipsSet(uint16 oldGlobalCommissionBips, uint16 newGlobalCommissionBips)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageFilterer) WatchGlobalCommissionBipsSet(opts *bind.WatchOpts, sink chan<- *PaymentCoordinatorStorageGlobalCommissionBipsSet) (event.Subscription, error) {

	logs, sub, err := _PaymentCoordinatorStorage.contract.WatchLogs(opts, "GlobalCommissionBipsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentCoordinatorStorageGlobalCommissionBipsSet)
				if err := _PaymentCoordinatorStorage.contract.UnpackLog(event, "GlobalCommissionBipsSet", log); err != nil {
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
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageFilterer) ParseGlobalCommissionBipsSet(log types.Log) (*PaymentCoordinatorStorageGlobalCommissionBipsSet, error) {
	event := new(PaymentCoordinatorStorageGlobalCommissionBipsSet)
	if err := _PaymentCoordinatorStorage.contract.UnpackLog(event, "GlobalCommissionBipsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentCoordinatorStoragePayAllForRangeSubmitterSetIterator is returned from FilterPayAllForRangeSubmitterSet and is used to iterate over the raw logs and unpacked data for PayAllForRangeSubmitterSet events raised by the PaymentCoordinatorStorage contract.
type PaymentCoordinatorStoragePayAllForRangeSubmitterSetIterator struct {
	Event *PaymentCoordinatorStoragePayAllForRangeSubmitterSet // Event containing the contract specifics and raw log

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
func (it *PaymentCoordinatorStoragePayAllForRangeSubmitterSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentCoordinatorStoragePayAllForRangeSubmitterSet)
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
		it.Event = new(PaymentCoordinatorStoragePayAllForRangeSubmitterSet)
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
func (it *PaymentCoordinatorStoragePayAllForRangeSubmitterSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentCoordinatorStoragePayAllForRangeSubmitterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentCoordinatorStoragePayAllForRangeSubmitterSet represents a PayAllForRangeSubmitterSet event raised by the PaymentCoordinatorStorage contract.
type PaymentCoordinatorStoragePayAllForRangeSubmitterSet struct {
	PayAllForRangeSubmitter common.Address
	OldValue                bool
	NewValue                bool
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterPayAllForRangeSubmitterSet is a free log retrieval operation binding the contract event 0x17b06be02af2803593116eb96121b9c6e8bee1cc1b145e7c31c19c180e86189b.
//
// Solidity: event PayAllForRangeSubmitterSet(address indexed payAllForRangeSubmitter, bool indexed oldValue, bool indexed newValue)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageFilterer) FilterPayAllForRangeSubmitterSet(opts *bind.FilterOpts, payAllForRangeSubmitter []common.Address, oldValue []bool, newValue []bool) (*PaymentCoordinatorStoragePayAllForRangeSubmitterSetIterator, error) {

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

	logs, sub, err := _PaymentCoordinatorStorage.contract.FilterLogs(opts, "PayAllForRangeSubmitterSet", payAllForRangeSubmitterRule, oldValueRule, newValueRule)
	if err != nil {
		return nil, err
	}
	return &PaymentCoordinatorStoragePayAllForRangeSubmitterSetIterator{contract: _PaymentCoordinatorStorage.contract, event: "PayAllForRangeSubmitterSet", logs: logs, sub: sub}, nil
}

// WatchPayAllForRangeSubmitterSet is a free log subscription operation binding the contract event 0x17b06be02af2803593116eb96121b9c6e8bee1cc1b145e7c31c19c180e86189b.
//
// Solidity: event PayAllForRangeSubmitterSet(address indexed payAllForRangeSubmitter, bool indexed oldValue, bool indexed newValue)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageFilterer) WatchPayAllForRangeSubmitterSet(opts *bind.WatchOpts, sink chan<- *PaymentCoordinatorStoragePayAllForRangeSubmitterSet, payAllForRangeSubmitter []common.Address, oldValue []bool, newValue []bool) (event.Subscription, error) {

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

	logs, sub, err := _PaymentCoordinatorStorage.contract.WatchLogs(opts, "PayAllForRangeSubmitterSet", payAllForRangeSubmitterRule, oldValueRule, newValueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentCoordinatorStoragePayAllForRangeSubmitterSet)
				if err := _PaymentCoordinatorStorage.contract.UnpackLog(event, "PayAllForRangeSubmitterSet", log); err != nil {
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
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageFilterer) ParsePayAllForRangeSubmitterSet(log types.Log) (*PaymentCoordinatorStoragePayAllForRangeSubmitterSet, error) {
	event := new(PaymentCoordinatorStoragePayAllForRangeSubmitterSet)
	if err := _PaymentCoordinatorStorage.contract.UnpackLog(event, "PayAllForRangeSubmitterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentCoordinatorStoragePaymentClaimedIterator is returned from FilterPaymentClaimed and is used to iterate over the raw logs and unpacked data for PaymentClaimed events raised by the PaymentCoordinatorStorage contract.
type PaymentCoordinatorStoragePaymentClaimedIterator struct {
	Event *PaymentCoordinatorStoragePaymentClaimed // Event containing the contract specifics and raw log

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
func (it *PaymentCoordinatorStoragePaymentClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentCoordinatorStoragePaymentClaimed)
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
		it.Event = new(PaymentCoordinatorStoragePaymentClaimed)
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
func (it *PaymentCoordinatorStoragePaymentClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentCoordinatorStoragePaymentClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentCoordinatorStoragePaymentClaimed represents a PaymentClaimed event raised by the PaymentCoordinatorStorage contract.
type PaymentCoordinatorStoragePaymentClaimed struct {
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
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageFilterer) FilterPaymentClaimed(opts *bind.FilterOpts, earner []common.Address, claimer []common.Address, recipient []common.Address) (*PaymentCoordinatorStoragePaymentClaimedIterator, error) {

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

	logs, sub, err := _PaymentCoordinatorStorage.contract.FilterLogs(opts, "PaymentClaimed", earnerRule, claimerRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &PaymentCoordinatorStoragePaymentClaimedIterator{contract: _PaymentCoordinatorStorage.contract, event: "PaymentClaimed", logs: logs, sub: sub}, nil
}

// WatchPaymentClaimed is a free log subscription operation binding the contract event 0xbff1e5a32b3f6d3b3c0a7e675ead2091fea820852f35a77abdd6d2420bec4778.
//
// Solidity: event PaymentClaimed(bytes32 root, address indexed earner, address indexed claimer, address indexed recipient, address token, uint256 claimedAmount)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageFilterer) WatchPaymentClaimed(opts *bind.WatchOpts, sink chan<- *PaymentCoordinatorStoragePaymentClaimed, earner []common.Address, claimer []common.Address, recipient []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PaymentCoordinatorStorage.contract.WatchLogs(opts, "PaymentClaimed", earnerRule, claimerRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentCoordinatorStoragePaymentClaimed)
				if err := _PaymentCoordinatorStorage.contract.UnpackLog(event, "PaymentClaimed", log); err != nil {
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
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageFilterer) ParsePaymentClaimed(log types.Log) (*PaymentCoordinatorStoragePaymentClaimed, error) {
	event := new(PaymentCoordinatorStoragePaymentClaimed)
	if err := _PaymentCoordinatorStorage.contract.UnpackLog(event, "PaymentClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentCoordinatorStoragePaymentUpdaterSetIterator is returned from FilterPaymentUpdaterSet and is used to iterate over the raw logs and unpacked data for PaymentUpdaterSet events raised by the PaymentCoordinatorStorage contract.
type PaymentCoordinatorStoragePaymentUpdaterSetIterator struct {
	Event *PaymentCoordinatorStoragePaymentUpdaterSet // Event containing the contract specifics and raw log

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
func (it *PaymentCoordinatorStoragePaymentUpdaterSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentCoordinatorStoragePaymentUpdaterSet)
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
		it.Event = new(PaymentCoordinatorStoragePaymentUpdaterSet)
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
func (it *PaymentCoordinatorStoragePaymentUpdaterSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentCoordinatorStoragePaymentUpdaterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentCoordinatorStoragePaymentUpdaterSet represents a PaymentUpdaterSet event raised by the PaymentCoordinatorStorage contract.
type PaymentCoordinatorStoragePaymentUpdaterSet struct {
	OldPaymentUpdater common.Address
	NewPaymentUpdater common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPaymentUpdaterSet is a free log retrieval operation binding the contract event 0x07d2890b3eb1206e7c3cb6bf8d46da31385ace3ce99abf85e5b690c83aa49678.
//
// Solidity: event PaymentUpdaterSet(address indexed oldPaymentUpdater, address indexed newPaymentUpdater)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageFilterer) FilterPaymentUpdaterSet(opts *bind.FilterOpts, oldPaymentUpdater []common.Address, newPaymentUpdater []common.Address) (*PaymentCoordinatorStoragePaymentUpdaterSetIterator, error) {

	var oldPaymentUpdaterRule []interface{}
	for _, oldPaymentUpdaterItem := range oldPaymentUpdater {
		oldPaymentUpdaterRule = append(oldPaymentUpdaterRule, oldPaymentUpdaterItem)
	}
	var newPaymentUpdaterRule []interface{}
	for _, newPaymentUpdaterItem := range newPaymentUpdater {
		newPaymentUpdaterRule = append(newPaymentUpdaterRule, newPaymentUpdaterItem)
	}

	logs, sub, err := _PaymentCoordinatorStorage.contract.FilterLogs(opts, "PaymentUpdaterSet", oldPaymentUpdaterRule, newPaymentUpdaterRule)
	if err != nil {
		return nil, err
	}
	return &PaymentCoordinatorStoragePaymentUpdaterSetIterator{contract: _PaymentCoordinatorStorage.contract, event: "PaymentUpdaterSet", logs: logs, sub: sub}, nil
}

// WatchPaymentUpdaterSet is a free log subscription operation binding the contract event 0x07d2890b3eb1206e7c3cb6bf8d46da31385ace3ce99abf85e5b690c83aa49678.
//
// Solidity: event PaymentUpdaterSet(address indexed oldPaymentUpdater, address indexed newPaymentUpdater)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageFilterer) WatchPaymentUpdaterSet(opts *bind.WatchOpts, sink chan<- *PaymentCoordinatorStoragePaymentUpdaterSet, oldPaymentUpdater []common.Address, newPaymentUpdater []common.Address) (event.Subscription, error) {

	var oldPaymentUpdaterRule []interface{}
	for _, oldPaymentUpdaterItem := range oldPaymentUpdater {
		oldPaymentUpdaterRule = append(oldPaymentUpdaterRule, oldPaymentUpdaterItem)
	}
	var newPaymentUpdaterRule []interface{}
	for _, newPaymentUpdaterItem := range newPaymentUpdater {
		newPaymentUpdaterRule = append(newPaymentUpdaterRule, newPaymentUpdaterItem)
	}

	logs, sub, err := _PaymentCoordinatorStorage.contract.WatchLogs(opts, "PaymentUpdaterSet", oldPaymentUpdaterRule, newPaymentUpdaterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentCoordinatorStoragePaymentUpdaterSet)
				if err := _PaymentCoordinatorStorage.contract.UnpackLog(event, "PaymentUpdaterSet", log); err != nil {
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
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageFilterer) ParsePaymentUpdaterSet(log types.Log) (*PaymentCoordinatorStoragePaymentUpdaterSet, error) {
	event := new(PaymentCoordinatorStoragePaymentUpdaterSet)
	if err := _PaymentCoordinatorStorage.contract.UnpackLog(event, "PaymentUpdaterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentCoordinatorStorageRangePaymentCreatedIterator is returned from FilterRangePaymentCreated and is used to iterate over the raw logs and unpacked data for RangePaymentCreated events raised by the PaymentCoordinatorStorage contract.
type PaymentCoordinatorStorageRangePaymentCreatedIterator struct {
	Event *PaymentCoordinatorStorageRangePaymentCreated // Event containing the contract specifics and raw log

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
func (it *PaymentCoordinatorStorageRangePaymentCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentCoordinatorStorageRangePaymentCreated)
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
		it.Event = new(PaymentCoordinatorStorageRangePaymentCreated)
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
func (it *PaymentCoordinatorStorageRangePaymentCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentCoordinatorStorageRangePaymentCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentCoordinatorStorageRangePaymentCreated represents a RangePaymentCreated event raised by the PaymentCoordinatorStorage contract.
type PaymentCoordinatorStorageRangePaymentCreated struct {
	Avs              common.Address
	PaymentNonce     *big.Int
	RangePaymentHash [32]byte
	RangePayment     IPaymentCoordinatorRangePayment
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRangePaymentCreated is a free log retrieval operation binding the contract event 0x2a0659fa4c38e0075469a0e0dd737e045dc316ffd6cb6e68755c119ee0882aea.
//
// Solidity: event RangePaymentCreated(address indexed avs, uint256 indexed paymentNonce, bytes32 indexed rangePaymentHash, ((address,uint96)[],address,uint256,uint32,uint32) rangePayment)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageFilterer) FilterRangePaymentCreated(opts *bind.FilterOpts, avs []common.Address, paymentNonce []*big.Int, rangePaymentHash [][32]byte) (*PaymentCoordinatorStorageRangePaymentCreatedIterator, error) {

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

	logs, sub, err := _PaymentCoordinatorStorage.contract.FilterLogs(opts, "RangePaymentCreated", avsRule, paymentNonceRule, rangePaymentHashRule)
	if err != nil {
		return nil, err
	}
	return &PaymentCoordinatorStorageRangePaymentCreatedIterator{contract: _PaymentCoordinatorStorage.contract, event: "RangePaymentCreated", logs: logs, sub: sub}, nil
}

// WatchRangePaymentCreated is a free log subscription operation binding the contract event 0x2a0659fa4c38e0075469a0e0dd737e045dc316ffd6cb6e68755c119ee0882aea.
//
// Solidity: event RangePaymentCreated(address indexed avs, uint256 indexed paymentNonce, bytes32 indexed rangePaymentHash, ((address,uint96)[],address,uint256,uint32,uint32) rangePayment)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageFilterer) WatchRangePaymentCreated(opts *bind.WatchOpts, sink chan<- *PaymentCoordinatorStorageRangePaymentCreated, avs []common.Address, paymentNonce []*big.Int, rangePaymentHash [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _PaymentCoordinatorStorage.contract.WatchLogs(opts, "RangePaymentCreated", avsRule, paymentNonceRule, rangePaymentHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentCoordinatorStorageRangePaymentCreated)
				if err := _PaymentCoordinatorStorage.contract.UnpackLog(event, "RangePaymentCreated", log); err != nil {
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
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageFilterer) ParseRangePaymentCreated(log types.Log) (*PaymentCoordinatorStorageRangePaymentCreated, error) {
	event := new(PaymentCoordinatorStorageRangePaymentCreated)
	if err := _PaymentCoordinatorStorage.contract.UnpackLog(event, "RangePaymentCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentCoordinatorStorageRangePaymentForAllCreatedIterator is returned from FilterRangePaymentForAllCreated and is used to iterate over the raw logs and unpacked data for RangePaymentForAllCreated events raised by the PaymentCoordinatorStorage contract.
type PaymentCoordinatorStorageRangePaymentForAllCreatedIterator struct {
	Event *PaymentCoordinatorStorageRangePaymentForAllCreated // Event containing the contract specifics and raw log

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
func (it *PaymentCoordinatorStorageRangePaymentForAllCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentCoordinatorStorageRangePaymentForAllCreated)
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
		it.Event = new(PaymentCoordinatorStorageRangePaymentForAllCreated)
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
func (it *PaymentCoordinatorStorageRangePaymentForAllCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentCoordinatorStorageRangePaymentForAllCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentCoordinatorStorageRangePaymentForAllCreated represents a RangePaymentForAllCreated event raised by the PaymentCoordinatorStorage contract.
type PaymentCoordinatorStorageRangePaymentForAllCreated struct {
	Submitter        common.Address
	PaymentNonce     *big.Int
	RangePaymentHash [32]byte
	RangePayment     IPaymentCoordinatorRangePayment
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRangePaymentForAllCreated is a free log retrieval operation binding the contract event 0xbc0782940ec5871f66c5490ef957f44d19a9adac1dac18b946ad0dd6579c30d6.
//
// Solidity: event RangePaymentForAllCreated(address indexed submitter, uint256 indexed paymentNonce, bytes32 indexed rangePaymentHash, ((address,uint96)[],address,uint256,uint32,uint32) rangePayment)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageFilterer) FilterRangePaymentForAllCreated(opts *bind.FilterOpts, submitter []common.Address, paymentNonce []*big.Int, rangePaymentHash [][32]byte) (*PaymentCoordinatorStorageRangePaymentForAllCreatedIterator, error) {

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

	logs, sub, err := _PaymentCoordinatorStorage.contract.FilterLogs(opts, "RangePaymentForAllCreated", submitterRule, paymentNonceRule, rangePaymentHashRule)
	if err != nil {
		return nil, err
	}
	return &PaymentCoordinatorStorageRangePaymentForAllCreatedIterator{contract: _PaymentCoordinatorStorage.contract, event: "RangePaymentForAllCreated", logs: logs, sub: sub}, nil
}

// WatchRangePaymentForAllCreated is a free log subscription operation binding the contract event 0xbc0782940ec5871f66c5490ef957f44d19a9adac1dac18b946ad0dd6579c30d6.
//
// Solidity: event RangePaymentForAllCreated(address indexed submitter, uint256 indexed paymentNonce, bytes32 indexed rangePaymentHash, ((address,uint96)[],address,uint256,uint32,uint32) rangePayment)
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageFilterer) WatchRangePaymentForAllCreated(opts *bind.WatchOpts, sink chan<- *PaymentCoordinatorStorageRangePaymentForAllCreated, submitter []common.Address, paymentNonce []*big.Int, rangePaymentHash [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _PaymentCoordinatorStorage.contract.WatchLogs(opts, "RangePaymentForAllCreated", submitterRule, paymentNonceRule, rangePaymentHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentCoordinatorStorageRangePaymentForAllCreated)
				if err := _PaymentCoordinatorStorage.contract.UnpackLog(event, "RangePaymentForAllCreated", log); err != nil {
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
func (_PaymentCoordinatorStorage *PaymentCoordinatorStorageFilterer) ParseRangePaymentForAllCreated(log types.Log) (*PaymentCoordinatorStorageRangePaymentForAllCreated, error) {
	event := new(PaymentCoordinatorStorageRangePaymentForAllCreated)
	if err := _PaymentCoordinatorStorage.contract.UnpackLog(event, "RangePaymentForAllCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
