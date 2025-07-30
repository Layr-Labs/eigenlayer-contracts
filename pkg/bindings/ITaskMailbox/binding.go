// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ITaskMailbox

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
	OperatorInfo      IOperatorTableCalculatorTypesBN254OperatorInfo
}

// IECDSACertificateVerifierTypesECDSACertificate is an auto generated low-level Go binding around an user-defined struct.
type IECDSACertificateVerifierTypesECDSACertificate struct {
	ReferenceTimestamp uint32
	MessageHash        [32]byte
	Sig                []byte
}

// IOperatorTableCalculatorTypesBN254OperatorInfo is an auto generated low-level Go binding around an user-defined struct.
type IOperatorTableCalculatorTypesBN254OperatorInfo struct {
	Pubkey  BN254G1Point
	Weights []*big.Int
}

// ITaskMailboxTypesConsensus is an auto generated low-level Go binding around an user-defined struct.
type ITaskMailboxTypesConsensus struct {
	ConsensusType uint8
	Value         []byte
}

// ITaskMailboxTypesExecutorOperatorSetTaskConfig is an auto generated low-level Go binding around an user-defined struct.
type ITaskMailboxTypesExecutorOperatorSetTaskConfig struct {
	TaskHook     common.Address
	TaskSLA      *big.Int
	FeeToken     common.Address
	CurveType    uint8
	FeeCollector common.Address
	Consensus    ITaskMailboxTypesConsensus
	TaskMetadata []byte
}

// ITaskMailboxTypesTask is an auto generated low-level Go binding around an user-defined struct.
type ITaskMailboxTypesTask struct {
	Creator                         common.Address
	CreationTime                    *big.Int
	Avs                             common.Address
	AvsFee                          *big.Int
	RefundCollector                 common.Address
	ExecutorOperatorSetId           uint32
	FeeSplit                        uint16
	Status                          uint8
	IsFeeRefunded                   bool
	OperatorTableReferenceTimestamp uint32
	ExecutorOperatorSetTaskConfig   ITaskMailboxTypesExecutorOperatorSetTaskConfig
	Payload                         []byte
	ExecutorCert                    []byte
	Result                          []byte
}

// ITaskMailboxTypesTaskParams is an auto generated low-level Go binding around an user-defined struct.
type ITaskMailboxTypesTaskParams struct {
	RefundCollector     common.Address
	ExecutorOperatorSet OperatorSet
	Payload             []byte
}

// OperatorSet is an auto generated low-level Go binding around an user-defined struct.
type OperatorSet struct {
	Avs common.Address
	Id  uint32
}

// ITaskMailboxMetaData contains all meta data concerning the ITaskMailbox contract.
var ITaskMailboxMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"BN254_CERTIFICATE_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ECDSA_CERTIFICATE_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createTask\",\"inputs\":[{\"name\":\"taskParams\",\"type\":\"tuple\",\"internalType\":\"structITaskMailboxTypes.TaskParams\",\"components\":[{\"name\":\"refundCollector\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"executorOperatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"taskHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"feeSplit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"feeSplitCollector\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBN254CertificateBytes\",\"inputs\":[{\"name\":\"cert\",\"type\":\"tuple\",\"internalType\":\"structIBN254CertificateVerifierTypes.BN254Certificate\",\"components\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"nonSignerWitnesses\",\"type\":\"tuple[]\",\"internalType\":\"structIBN254CertificateVerifierTypes.BN254OperatorInfoWitness[]\",\"components\":[{\"name\":\"operatorIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorInfoProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"operatorInfo\",\"type\":\"tuple\",\"internalType\":\"structIOperatorTableCalculatorTypes.BN254OperatorInfo\",\"components\":[{\"name\":\"pubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}]}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getECDSACertificateBytes\",\"inputs\":[{\"name\":\"cert\",\"type\":\"tuple\",\"internalType\":\"structIECDSACertificateVerifierTypes.ECDSACertificate\",\"components\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sig\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getExecutorOperatorSetTaskConfig\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structITaskMailboxTypes.ExecutorOperatorSetTaskConfig\",\"components\":[{\"name\":\"taskHook\",\"type\":\"address\",\"internalType\":\"contractIAVSTaskHook\"},{\"name\":\"taskSLA\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"curveType\",\"type\":\"uint8\",\"internalType\":\"enumIKeyRegistrarTypes.CurveType\"},{\"name\":\"feeCollector\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"consensus\",\"type\":\"tuple\",\"internalType\":\"structITaskMailboxTypes.Consensus\",\"components\":[{\"name\":\"consensusType\",\"type\":\"uint8\",\"internalType\":\"enumITaskMailboxTypes.ConsensusType\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"taskMetadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskInfo\",\"inputs\":[{\"name\":\"taskHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structITaskMailboxTypes.Task\",\"components\":[{\"name\":\"creator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"creationTime\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avsFee\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"refundCollector\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"executorOperatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeSplit\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumITaskMailboxTypes.TaskStatus\"},{\"name\":\"isFeeRefunded\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"operatorTableReferenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"executorOperatorSetTaskConfig\",\"type\":\"tuple\",\"internalType\":\"structITaskMailboxTypes.ExecutorOperatorSetTaskConfig\",\"components\":[{\"name\":\"taskHook\",\"type\":\"address\",\"internalType\":\"contractIAVSTaskHook\"},{\"name\":\"taskSLA\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"curveType\",\"type\":\"uint8\",\"internalType\":\"enumIKeyRegistrarTypes.CurveType\"},{\"name\":\"feeCollector\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"consensus\",\"type\":\"tuple\",\"internalType\":\"structITaskMailboxTypes.Consensus\",\"components\":[{\"name\":\"consensusType\",\"type\":\"uint8\",\"internalType\":\"enumITaskMailboxTypes.ConsensusType\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"taskMetadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"executorCert\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"result\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskResult\",\"inputs\":[{\"name\":\"taskHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskStatus\",\"inputs\":[{\"name\":\"taskHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumITaskMailboxTypes.TaskStatus\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeSplit\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"feeSplitCollector\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isExecutorOperatorSetRegistered\",\"inputs\":[{\"name\":\"operatorSetKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"refundFee\",\"inputs\":[{\"name\":\"taskHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerExecutorOperatorSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"isRegistered\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setExecutorOperatorSetTaskConfig\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structITaskMailboxTypes.ExecutorOperatorSetTaskConfig\",\"components\":[{\"name\":\"taskHook\",\"type\":\"address\",\"internalType\":\"contractIAVSTaskHook\"},{\"name\":\"taskSLA\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"curveType\",\"type\":\"uint8\",\"internalType\":\"enumIKeyRegistrarTypes.CurveType\"},{\"name\":\"feeCollector\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"consensus\",\"type\":\"tuple\",\"internalType\":\"structITaskMailboxTypes.Consensus\",\"components\":[{\"name\":\"consensusType\",\"type\":\"uint8\",\"internalType\":\"enumITaskMailboxTypes.ConsensusType\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"taskMetadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFeeSplit\",\"inputs\":[{\"name\":\"feeSplit\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFeeSplitCollector\",\"inputs\":[{\"name\":\"feeSplitCollector\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitResult\",\"inputs\":[{\"name\":\"taskHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"executorCert\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"result\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ExecutorOperatorSetRegistered\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"executorOperatorSetId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"isRegistered\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExecutorOperatorSetTaskConfigSet\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"executorOperatorSetId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structITaskMailboxTypes.ExecutorOperatorSetTaskConfig\",\"components\":[{\"name\":\"taskHook\",\"type\":\"address\",\"internalType\":\"contractIAVSTaskHook\"},{\"name\":\"taskSLA\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"curveType\",\"type\":\"uint8\",\"internalType\":\"enumIKeyRegistrarTypes.CurveType\"},{\"name\":\"feeCollector\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"consensus\",\"type\":\"tuple\",\"internalType\":\"structITaskMailboxTypes.Consensus\",\"components\":[{\"name\":\"consensusType\",\"type\":\"uint8\",\"internalType\":\"enumITaskMailboxTypes.ConsensusType\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"taskMetadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeRefunded\",\"inputs\":[{\"name\":\"refundCollector\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"taskHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"avsFee\",\"type\":\"uint96\",\"indexed\":false,\"internalType\":\"uint96\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeSplitCollectorSet\",\"inputs\":[{\"name\":\"feeSplitCollector\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeSplitSet\",\"inputs\":[{\"name\":\"feeSplit\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskCreated\",\"inputs\":[{\"name\":\"creator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"taskHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"executorOperatorSetId\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"operatorTableReferenceTimestamp\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"refundCollector\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"avsFee\",\"type\":\"uint96\",\"indexed\":false,\"internalType\":\"uint96\"},{\"name\":\"taskDeadline\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskVerified\",\"inputs\":[{\"name\":\"aggregator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"taskHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"executorOperatorSetId\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"executorCert\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"result\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"EmptyCertificateSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExecutorOperatorSetNotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExecutorOperatorSetTaskConfigNotSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FeeAlreadyRefunded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidConsensusType\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidConsensusValue\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCurveType\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidFeeReceiver\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidFeeSplit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidMessageHash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperatorSetOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidReferenceTimestamp\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTaskCreator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTaskStatus\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"enumITaskMailboxTypes.TaskStatus\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"enumITaskMailboxTypes.TaskStatus\"}]},{\"type\":\"error\",\"name\":\"OnlyRefundCollector\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PayloadIsEmpty\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ThresholdNotMet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TimestampAtCreation\",\"inputs\":[]}]",
}

// ITaskMailboxABI is the input ABI used to generate the binding from.
// Deprecated: Use ITaskMailboxMetaData.ABI instead.
var ITaskMailboxABI = ITaskMailboxMetaData.ABI

// ITaskMailbox is an auto generated Go binding around an Ethereum contract.
type ITaskMailbox struct {
	ITaskMailboxCaller     // Read-only binding to the contract
	ITaskMailboxTransactor // Write-only binding to the contract
	ITaskMailboxFilterer   // Log filterer for contract events
}

// ITaskMailboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type ITaskMailboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITaskMailboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ITaskMailboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITaskMailboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ITaskMailboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITaskMailboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ITaskMailboxSession struct {
	Contract     *ITaskMailbox     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ITaskMailboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ITaskMailboxCallerSession struct {
	Contract *ITaskMailboxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ITaskMailboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ITaskMailboxTransactorSession struct {
	Contract     *ITaskMailboxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ITaskMailboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type ITaskMailboxRaw struct {
	Contract *ITaskMailbox // Generic contract binding to access the raw methods on
}

// ITaskMailboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ITaskMailboxCallerRaw struct {
	Contract *ITaskMailboxCaller // Generic read-only contract binding to access the raw methods on
}

// ITaskMailboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ITaskMailboxTransactorRaw struct {
	Contract *ITaskMailboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewITaskMailbox creates a new instance of ITaskMailbox, bound to a specific deployed contract.
func NewITaskMailbox(address common.Address, backend bind.ContractBackend) (*ITaskMailbox, error) {
	contract, err := bindITaskMailbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ITaskMailbox{ITaskMailboxCaller: ITaskMailboxCaller{contract: contract}, ITaskMailboxTransactor: ITaskMailboxTransactor{contract: contract}, ITaskMailboxFilterer: ITaskMailboxFilterer{contract: contract}}, nil
}

// NewITaskMailboxCaller creates a new read-only instance of ITaskMailbox, bound to a specific deployed contract.
func NewITaskMailboxCaller(address common.Address, caller bind.ContractCaller) (*ITaskMailboxCaller, error) {
	contract, err := bindITaskMailbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ITaskMailboxCaller{contract: contract}, nil
}

// NewITaskMailboxTransactor creates a new write-only instance of ITaskMailbox, bound to a specific deployed contract.
func NewITaskMailboxTransactor(address common.Address, transactor bind.ContractTransactor) (*ITaskMailboxTransactor, error) {
	contract, err := bindITaskMailbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ITaskMailboxTransactor{contract: contract}, nil
}

// NewITaskMailboxFilterer creates a new log filterer instance of ITaskMailbox, bound to a specific deployed contract.
func NewITaskMailboxFilterer(address common.Address, filterer bind.ContractFilterer) (*ITaskMailboxFilterer, error) {
	contract, err := bindITaskMailbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ITaskMailboxFilterer{contract: contract}, nil
}

// bindITaskMailbox binds a generic wrapper to an already deployed contract.
func bindITaskMailbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ITaskMailboxMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITaskMailbox *ITaskMailboxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITaskMailbox.Contract.ITaskMailboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITaskMailbox *ITaskMailboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITaskMailbox.Contract.ITaskMailboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITaskMailbox *ITaskMailboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITaskMailbox.Contract.ITaskMailboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITaskMailbox *ITaskMailboxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITaskMailbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITaskMailbox *ITaskMailboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITaskMailbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITaskMailbox *ITaskMailboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITaskMailbox.Contract.contract.Transact(opts, method, params...)
}

// BN254CERTIFICATEVERIFIER is a free data retrieval call binding the contract method 0xf7424fc9.
//
// Solidity: function BN254_CERTIFICATE_VERIFIER() view returns(address)
func (_ITaskMailbox *ITaskMailboxCaller) BN254CERTIFICATEVERIFIER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ITaskMailbox.contract.Call(opts, &out, "BN254_CERTIFICATE_VERIFIER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BN254CERTIFICATEVERIFIER is a free data retrieval call binding the contract method 0xf7424fc9.
//
// Solidity: function BN254_CERTIFICATE_VERIFIER() view returns(address)
func (_ITaskMailbox *ITaskMailboxSession) BN254CERTIFICATEVERIFIER() (common.Address, error) {
	return _ITaskMailbox.Contract.BN254CERTIFICATEVERIFIER(&_ITaskMailbox.CallOpts)
}

// BN254CERTIFICATEVERIFIER is a free data retrieval call binding the contract method 0xf7424fc9.
//
// Solidity: function BN254_CERTIFICATE_VERIFIER() view returns(address)
func (_ITaskMailbox *ITaskMailboxCallerSession) BN254CERTIFICATEVERIFIER() (common.Address, error) {
	return _ITaskMailbox.Contract.BN254CERTIFICATEVERIFIER(&_ITaskMailbox.CallOpts)
}

// ECDSACERTIFICATEVERIFIER is a free data retrieval call binding the contract method 0x54743ad2.
//
// Solidity: function ECDSA_CERTIFICATE_VERIFIER() view returns(address)
func (_ITaskMailbox *ITaskMailboxCaller) ECDSACERTIFICATEVERIFIER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ITaskMailbox.contract.Call(opts, &out, "ECDSA_CERTIFICATE_VERIFIER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ECDSACERTIFICATEVERIFIER is a free data retrieval call binding the contract method 0x54743ad2.
//
// Solidity: function ECDSA_CERTIFICATE_VERIFIER() view returns(address)
func (_ITaskMailbox *ITaskMailboxSession) ECDSACERTIFICATEVERIFIER() (common.Address, error) {
	return _ITaskMailbox.Contract.ECDSACERTIFICATEVERIFIER(&_ITaskMailbox.CallOpts)
}

// ECDSACERTIFICATEVERIFIER is a free data retrieval call binding the contract method 0x54743ad2.
//
// Solidity: function ECDSA_CERTIFICATE_VERIFIER() view returns(address)
func (_ITaskMailbox *ITaskMailboxCallerSession) ECDSACERTIFICATEVERIFIER() (common.Address, error) {
	return _ITaskMailbox.Contract.ECDSACERTIFICATEVERIFIER(&_ITaskMailbox.CallOpts)
}

// FeeSplit is a free data retrieval call binding the contract method 0x6373ea69.
//
// Solidity: function feeSplit() view returns(uint16)
func (_ITaskMailbox *ITaskMailboxCaller) FeeSplit(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ITaskMailbox.contract.Call(opts, &out, "feeSplit")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// FeeSplit is a free data retrieval call binding the contract method 0x6373ea69.
//
// Solidity: function feeSplit() view returns(uint16)
func (_ITaskMailbox *ITaskMailboxSession) FeeSplit() (uint16, error) {
	return _ITaskMailbox.Contract.FeeSplit(&_ITaskMailbox.CallOpts)
}

// FeeSplit is a free data retrieval call binding the contract method 0x6373ea69.
//
// Solidity: function feeSplit() view returns(uint16)
func (_ITaskMailbox *ITaskMailboxCallerSession) FeeSplit() (uint16, error) {
	return _ITaskMailbox.Contract.FeeSplit(&_ITaskMailbox.CallOpts)
}

// FeeSplitCollector is a free data retrieval call binding the contract method 0x1a20c505.
//
// Solidity: function feeSplitCollector() view returns(address)
func (_ITaskMailbox *ITaskMailboxCaller) FeeSplitCollector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ITaskMailbox.contract.Call(opts, &out, "feeSplitCollector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeSplitCollector is a free data retrieval call binding the contract method 0x1a20c505.
//
// Solidity: function feeSplitCollector() view returns(address)
func (_ITaskMailbox *ITaskMailboxSession) FeeSplitCollector() (common.Address, error) {
	return _ITaskMailbox.Contract.FeeSplitCollector(&_ITaskMailbox.CallOpts)
}

// FeeSplitCollector is a free data retrieval call binding the contract method 0x1a20c505.
//
// Solidity: function feeSplitCollector() view returns(address)
func (_ITaskMailbox *ITaskMailboxCallerSession) FeeSplitCollector() (common.Address, error) {
	return _ITaskMailbox.Contract.FeeSplitCollector(&_ITaskMailbox.CallOpts)
}

// GetBN254CertificateBytes is a free data retrieval call binding the contract method 0x1ae370eb.
//
// Solidity: function getBN254CertificateBytes((uint32,bytes32,(uint256,uint256),(uint256[2],uint256[2]),(uint32,bytes,((uint256,uint256),uint256[]))[]) cert) pure returns(bytes)
func (_ITaskMailbox *ITaskMailboxCaller) GetBN254CertificateBytes(opts *bind.CallOpts, cert IBN254CertificateVerifierTypesBN254Certificate) ([]byte, error) {
	var out []interface{}
	err := _ITaskMailbox.contract.Call(opts, &out, "getBN254CertificateBytes", cert)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetBN254CertificateBytes is a free data retrieval call binding the contract method 0x1ae370eb.
//
// Solidity: function getBN254CertificateBytes((uint32,bytes32,(uint256,uint256),(uint256[2],uint256[2]),(uint32,bytes,((uint256,uint256),uint256[]))[]) cert) pure returns(bytes)
func (_ITaskMailbox *ITaskMailboxSession) GetBN254CertificateBytes(cert IBN254CertificateVerifierTypesBN254Certificate) ([]byte, error) {
	return _ITaskMailbox.Contract.GetBN254CertificateBytes(&_ITaskMailbox.CallOpts, cert)
}

// GetBN254CertificateBytes is a free data retrieval call binding the contract method 0x1ae370eb.
//
// Solidity: function getBN254CertificateBytes((uint32,bytes32,(uint256,uint256),(uint256[2],uint256[2]),(uint32,bytes,((uint256,uint256),uint256[]))[]) cert) pure returns(bytes)
func (_ITaskMailbox *ITaskMailboxCallerSession) GetBN254CertificateBytes(cert IBN254CertificateVerifierTypesBN254Certificate) ([]byte, error) {
	return _ITaskMailbox.Contract.GetBN254CertificateBytes(&_ITaskMailbox.CallOpts, cert)
}

// GetECDSACertificateBytes is a free data retrieval call binding the contract method 0x1270a892.
//
// Solidity: function getECDSACertificateBytes((uint32,bytes32,bytes) cert) pure returns(bytes)
func (_ITaskMailbox *ITaskMailboxCaller) GetECDSACertificateBytes(opts *bind.CallOpts, cert IECDSACertificateVerifierTypesECDSACertificate) ([]byte, error) {
	var out []interface{}
	err := _ITaskMailbox.contract.Call(opts, &out, "getECDSACertificateBytes", cert)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetECDSACertificateBytes is a free data retrieval call binding the contract method 0x1270a892.
//
// Solidity: function getECDSACertificateBytes((uint32,bytes32,bytes) cert) pure returns(bytes)
func (_ITaskMailbox *ITaskMailboxSession) GetECDSACertificateBytes(cert IECDSACertificateVerifierTypesECDSACertificate) ([]byte, error) {
	return _ITaskMailbox.Contract.GetECDSACertificateBytes(&_ITaskMailbox.CallOpts, cert)
}

// GetECDSACertificateBytes is a free data retrieval call binding the contract method 0x1270a892.
//
// Solidity: function getECDSACertificateBytes((uint32,bytes32,bytes) cert) pure returns(bytes)
func (_ITaskMailbox *ITaskMailboxCallerSession) GetECDSACertificateBytes(cert IECDSACertificateVerifierTypesECDSACertificate) ([]byte, error) {
	return _ITaskMailbox.Contract.GetECDSACertificateBytes(&_ITaskMailbox.CallOpts, cert)
}

// GetExecutorOperatorSetTaskConfig is a free data retrieval call binding the contract method 0x6bf6fad5.
//
// Solidity: function getExecutorOperatorSetTaskConfig((address,uint32) operatorSet) view returns((address,uint96,address,uint8,address,(uint8,bytes),bytes))
func (_ITaskMailbox *ITaskMailboxCaller) GetExecutorOperatorSetTaskConfig(opts *bind.CallOpts, operatorSet OperatorSet) (ITaskMailboxTypesExecutorOperatorSetTaskConfig, error) {
	var out []interface{}
	err := _ITaskMailbox.contract.Call(opts, &out, "getExecutorOperatorSetTaskConfig", operatorSet)

	if err != nil {
		return *new(ITaskMailboxTypesExecutorOperatorSetTaskConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(ITaskMailboxTypesExecutorOperatorSetTaskConfig)).(*ITaskMailboxTypesExecutorOperatorSetTaskConfig)

	return out0, err

}

// GetExecutorOperatorSetTaskConfig is a free data retrieval call binding the contract method 0x6bf6fad5.
//
// Solidity: function getExecutorOperatorSetTaskConfig((address,uint32) operatorSet) view returns((address,uint96,address,uint8,address,(uint8,bytes),bytes))
func (_ITaskMailbox *ITaskMailboxSession) GetExecutorOperatorSetTaskConfig(operatorSet OperatorSet) (ITaskMailboxTypesExecutorOperatorSetTaskConfig, error) {
	return _ITaskMailbox.Contract.GetExecutorOperatorSetTaskConfig(&_ITaskMailbox.CallOpts, operatorSet)
}

// GetExecutorOperatorSetTaskConfig is a free data retrieval call binding the contract method 0x6bf6fad5.
//
// Solidity: function getExecutorOperatorSetTaskConfig((address,uint32) operatorSet) view returns((address,uint96,address,uint8,address,(uint8,bytes),bytes))
func (_ITaskMailbox *ITaskMailboxCallerSession) GetExecutorOperatorSetTaskConfig(operatorSet OperatorSet) (ITaskMailboxTypesExecutorOperatorSetTaskConfig, error) {
	return _ITaskMailbox.Contract.GetExecutorOperatorSetTaskConfig(&_ITaskMailbox.CallOpts, operatorSet)
}

// GetTaskInfo is a free data retrieval call binding the contract method 0x4ad52e02.
//
// Solidity: function getTaskInfo(bytes32 taskHash) view returns((address,uint96,address,uint96,address,uint32,uint16,uint8,bool,uint32,(address,uint96,address,uint8,address,(uint8,bytes),bytes),bytes,bytes,bytes))
func (_ITaskMailbox *ITaskMailboxCaller) GetTaskInfo(opts *bind.CallOpts, taskHash [32]byte) (ITaskMailboxTypesTask, error) {
	var out []interface{}
	err := _ITaskMailbox.contract.Call(opts, &out, "getTaskInfo", taskHash)

	if err != nil {
		return *new(ITaskMailboxTypesTask), err
	}

	out0 := *abi.ConvertType(out[0], new(ITaskMailboxTypesTask)).(*ITaskMailboxTypesTask)

	return out0, err

}

// GetTaskInfo is a free data retrieval call binding the contract method 0x4ad52e02.
//
// Solidity: function getTaskInfo(bytes32 taskHash) view returns((address,uint96,address,uint96,address,uint32,uint16,uint8,bool,uint32,(address,uint96,address,uint8,address,(uint8,bytes),bytes),bytes,bytes,bytes))
func (_ITaskMailbox *ITaskMailboxSession) GetTaskInfo(taskHash [32]byte) (ITaskMailboxTypesTask, error) {
	return _ITaskMailbox.Contract.GetTaskInfo(&_ITaskMailbox.CallOpts, taskHash)
}

// GetTaskInfo is a free data retrieval call binding the contract method 0x4ad52e02.
//
// Solidity: function getTaskInfo(bytes32 taskHash) view returns((address,uint96,address,uint96,address,uint32,uint16,uint8,bool,uint32,(address,uint96,address,uint8,address,(uint8,bytes),bytes),bytes,bytes,bytes))
func (_ITaskMailbox *ITaskMailboxCallerSession) GetTaskInfo(taskHash [32]byte) (ITaskMailboxTypesTask, error) {
	return _ITaskMailbox.Contract.GetTaskInfo(&_ITaskMailbox.CallOpts, taskHash)
}

// GetTaskResult is a free data retrieval call binding the contract method 0x62fee037.
//
// Solidity: function getTaskResult(bytes32 taskHash) view returns(bytes)
func (_ITaskMailbox *ITaskMailboxCaller) GetTaskResult(opts *bind.CallOpts, taskHash [32]byte) ([]byte, error) {
	var out []interface{}
	err := _ITaskMailbox.contract.Call(opts, &out, "getTaskResult", taskHash)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetTaskResult is a free data retrieval call binding the contract method 0x62fee037.
//
// Solidity: function getTaskResult(bytes32 taskHash) view returns(bytes)
func (_ITaskMailbox *ITaskMailboxSession) GetTaskResult(taskHash [32]byte) ([]byte, error) {
	return _ITaskMailbox.Contract.GetTaskResult(&_ITaskMailbox.CallOpts, taskHash)
}

// GetTaskResult is a free data retrieval call binding the contract method 0x62fee037.
//
// Solidity: function getTaskResult(bytes32 taskHash) view returns(bytes)
func (_ITaskMailbox *ITaskMailboxCallerSession) GetTaskResult(taskHash [32]byte) ([]byte, error) {
	return _ITaskMailbox.Contract.GetTaskResult(&_ITaskMailbox.CallOpts, taskHash)
}

// GetTaskStatus is a free data retrieval call binding the contract method 0x2bf6cc79.
//
// Solidity: function getTaskStatus(bytes32 taskHash) view returns(uint8)
func (_ITaskMailbox *ITaskMailboxCaller) GetTaskStatus(opts *bind.CallOpts, taskHash [32]byte) (uint8, error) {
	var out []interface{}
	err := _ITaskMailbox.contract.Call(opts, &out, "getTaskStatus", taskHash)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetTaskStatus is a free data retrieval call binding the contract method 0x2bf6cc79.
//
// Solidity: function getTaskStatus(bytes32 taskHash) view returns(uint8)
func (_ITaskMailbox *ITaskMailboxSession) GetTaskStatus(taskHash [32]byte) (uint8, error) {
	return _ITaskMailbox.Contract.GetTaskStatus(&_ITaskMailbox.CallOpts, taskHash)
}

// GetTaskStatus is a free data retrieval call binding the contract method 0x2bf6cc79.
//
// Solidity: function getTaskStatus(bytes32 taskHash) view returns(uint8)
func (_ITaskMailbox *ITaskMailboxCallerSession) GetTaskStatus(taskHash [32]byte) (uint8, error) {
	return _ITaskMailbox.Contract.GetTaskStatus(&_ITaskMailbox.CallOpts, taskHash)
}

// IsExecutorOperatorSetRegistered is a free data retrieval call binding the contract method 0xfa2c0b37.
//
// Solidity: function isExecutorOperatorSetRegistered(bytes32 operatorSetKey) view returns(bool)
func (_ITaskMailbox *ITaskMailboxCaller) IsExecutorOperatorSetRegistered(opts *bind.CallOpts, operatorSetKey [32]byte) (bool, error) {
	var out []interface{}
	err := _ITaskMailbox.contract.Call(opts, &out, "isExecutorOperatorSetRegistered", operatorSetKey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExecutorOperatorSetRegistered is a free data retrieval call binding the contract method 0xfa2c0b37.
//
// Solidity: function isExecutorOperatorSetRegistered(bytes32 operatorSetKey) view returns(bool)
func (_ITaskMailbox *ITaskMailboxSession) IsExecutorOperatorSetRegistered(operatorSetKey [32]byte) (bool, error) {
	return _ITaskMailbox.Contract.IsExecutorOperatorSetRegistered(&_ITaskMailbox.CallOpts, operatorSetKey)
}

// IsExecutorOperatorSetRegistered is a free data retrieval call binding the contract method 0xfa2c0b37.
//
// Solidity: function isExecutorOperatorSetRegistered(bytes32 operatorSetKey) view returns(bool)
func (_ITaskMailbox *ITaskMailboxCallerSession) IsExecutorOperatorSetRegistered(operatorSetKey [32]byte) (bool, error) {
	return _ITaskMailbox.Contract.IsExecutorOperatorSetRegistered(&_ITaskMailbox.CallOpts, operatorSetKey)
}

// CreateTask is a paid mutator transaction binding the contract method 0x1fb66f5d.
//
// Solidity: function createTask((address,(address,uint32),bytes) taskParams) returns(bytes32 taskHash)
func (_ITaskMailbox *ITaskMailboxTransactor) CreateTask(opts *bind.TransactOpts, taskParams ITaskMailboxTypesTaskParams) (*types.Transaction, error) {
	return _ITaskMailbox.contract.Transact(opts, "createTask", taskParams)
}

// CreateTask is a paid mutator transaction binding the contract method 0x1fb66f5d.
//
// Solidity: function createTask((address,(address,uint32),bytes) taskParams) returns(bytes32 taskHash)
func (_ITaskMailbox *ITaskMailboxSession) CreateTask(taskParams ITaskMailboxTypesTaskParams) (*types.Transaction, error) {
	return _ITaskMailbox.Contract.CreateTask(&_ITaskMailbox.TransactOpts, taskParams)
}

// CreateTask is a paid mutator transaction binding the contract method 0x1fb66f5d.
//
// Solidity: function createTask((address,(address,uint32),bytes) taskParams) returns(bytes32 taskHash)
func (_ITaskMailbox *ITaskMailboxTransactorSession) CreateTask(taskParams ITaskMailboxTypesTaskParams) (*types.Transaction, error) {
	return _ITaskMailbox.Contract.CreateTask(&_ITaskMailbox.TransactOpts, taskParams)
}

// Initialize is a paid mutator transaction binding the contract method 0x708c0db9.
//
// Solidity: function initialize(address owner, uint16 feeSplit, address feeSplitCollector) returns()
func (_ITaskMailbox *ITaskMailboxTransactor) Initialize(opts *bind.TransactOpts, owner common.Address, feeSplit uint16, feeSplitCollector common.Address) (*types.Transaction, error) {
	return _ITaskMailbox.contract.Transact(opts, "initialize", owner, feeSplit, feeSplitCollector)
}

// Initialize is a paid mutator transaction binding the contract method 0x708c0db9.
//
// Solidity: function initialize(address owner, uint16 feeSplit, address feeSplitCollector) returns()
func (_ITaskMailbox *ITaskMailboxSession) Initialize(owner common.Address, feeSplit uint16, feeSplitCollector common.Address) (*types.Transaction, error) {
	return _ITaskMailbox.Contract.Initialize(&_ITaskMailbox.TransactOpts, owner, feeSplit, feeSplitCollector)
}

// Initialize is a paid mutator transaction binding the contract method 0x708c0db9.
//
// Solidity: function initialize(address owner, uint16 feeSplit, address feeSplitCollector) returns()
func (_ITaskMailbox *ITaskMailboxTransactorSession) Initialize(owner common.Address, feeSplit uint16, feeSplitCollector common.Address) (*types.Transaction, error) {
	return _ITaskMailbox.Contract.Initialize(&_ITaskMailbox.TransactOpts, owner, feeSplit, feeSplitCollector)
}

// RefundFee is a paid mutator transaction binding the contract method 0xb8694166.
//
// Solidity: function refundFee(bytes32 taskHash) returns()
func (_ITaskMailbox *ITaskMailboxTransactor) RefundFee(opts *bind.TransactOpts, taskHash [32]byte) (*types.Transaction, error) {
	return _ITaskMailbox.contract.Transact(opts, "refundFee", taskHash)
}

// RefundFee is a paid mutator transaction binding the contract method 0xb8694166.
//
// Solidity: function refundFee(bytes32 taskHash) returns()
func (_ITaskMailbox *ITaskMailboxSession) RefundFee(taskHash [32]byte) (*types.Transaction, error) {
	return _ITaskMailbox.Contract.RefundFee(&_ITaskMailbox.TransactOpts, taskHash)
}

// RefundFee is a paid mutator transaction binding the contract method 0xb8694166.
//
// Solidity: function refundFee(bytes32 taskHash) returns()
func (_ITaskMailbox *ITaskMailboxTransactorSession) RefundFee(taskHash [32]byte) (*types.Transaction, error) {
	return _ITaskMailbox.Contract.RefundFee(&_ITaskMailbox.TransactOpts, taskHash)
}

// RegisterExecutorOperatorSet is a paid mutator transaction binding the contract method 0x49acd884.
//
// Solidity: function registerExecutorOperatorSet((address,uint32) operatorSet, bool isRegistered) returns()
func (_ITaskMailbox *ITaskMailboxTransactor) RegisterExecutorOperatorSet(opts *bind.TransactOpts, operatorSet OperatorSet, isRegistered bool) (*types.Transaction, error) {
	return _ITaskMailbox.contract.Transact(opts, "registerExecutorOperatorSet", operatorSet, isRegistered)
}

// RegisterExecutorOperatorSet is a paid mutator transaction binding the contract method 0x49acd884.
//
// Solidity: function registerExecutorOperatorSet((address,uint32) operatorSet, bool isRegistered) returns()
func (_ITaskMailbox *ITaskMailboxSession) RegisterExecutorOperatorSet(operatorSet OperatorSet, isRegistered bool) (*types.Transaction, error) {
	return _ITaskMailbox.Contract.RegisterExecutorOperatorSet(&_ITaskMailbox.TransactOpts, operatorSet, isRegistered)
}

// RegisterExecutorOperatorSet is a paid mutator transaction binding the contract method 0x49acd884.
//
// Solidity: function registerExecutorOperatorSet((address,uint32) operatorSet, bool isRegistered) returns()
func (_ITaskMailbox *ITaskMailboxTransactorSession) RegisterExecutorOperatorSet(operatorSet OperatorSet, isRegistered bool) (*types.Transaction, error) {
	return _ITaskMailbox.Contract.RegisterExecutorOperatorSet(&_ITaskMailbox.TransactOpts, operatorSet, isRegistered)
}

// SetExecutorOperatorSetTaskConfig is a paid mutator transaction binding the contract method 0xf741e81a.
//
// Solidity: function setExecutorOperatorSetTaskConfig((address,uint32) operatorSet, (address,uint96,address,uint8,address,(uint8,bytes),bytes) config) returns()
func (_ITaskMailbox *ITaskMailboxTransactor) SetExecutorOperatorSetTaskConfig(opts *bind.TransactOpts, operatorSet OperatorSet, config ITaskMailboxTypesExecutorOperatorSetTaskConfig) (*types.Transaction, error) {
	return _ITaskMailbox.contract.Transact(opts, "setExecutorOperatorSetTaskConfig", operatorSet, config)
}

// SetExecutorOperatorSetTaskConfig is a paid mutator transaction binding the contract method 0xf741e81a.
//
// Solidity: function setExecutorOperatorSetTaskConfig((address,uint32) operatorSet, (address,uint96,address,uint8,address,(uint8,bytes),bytes) config) returns()
func (_ITaskMailbox *ITaskMailboxSession) SetExecutorOperatorSetTaskConfig(operatorSet OperatorSet, config ITaskMailboxTypesExecutorOperatorSetTaskConfig) (*types.Transaction, error) {
	return _ITaskMailbox.Contract.SetExecutorOperatorSetTaskConfig(&_ITaskMailbox.TransactOpts, operatorSet, config)
}

// SetExecutorOperatorSetTaskConfig is a paid mutator transaction binding the contract method 0xf741e81a.
//
// Solidity: function setExecutorOperatorSetTaskConfig((address,uint32) operatorSet, (address,uint96,address,uint8,address,(uint8,bytes),bytes) config) returns()
func (_ITaskMailbox *ITaskMailboxTransactorSession) SetExecutorOperatorSetTaskConfig(operatorSet OperatorSet, config ITaskMailboxTypesExecutorOperatorSetTaskConfig) (*types.Transaction, error) {
	return _ITaskMailbox.Contract.SetExecutorOperatorSetTaskConfig(&_ITaskMailbox.TransactOpts, operatorSet, config)
}

// SetFeeSplit is a paid mutator transaction binding the contract method 0x468c07a0.
//
// Solidity: function setFeeSplit(uint16 feeSplit) returns()
func (_ITaskMailbox *ITaskMailboxTransactor) SetFeeSplit(opts *bind.TransactOpts, feeSplit uint16) (*types.Transaction, error) {
	return _ITaskMailbox.contract.Transact(opts, "setFeeSplit", feeSplit)
}

// SetFeeSplit is a paid mutator transaction binding the contract method 0x468c07a0.
//
// Solidity: function setFeeSplit(uint16 feeSplit) returns()
func (_ITaskMailbox *ITaskMailboxSession) SetFeeSplit(feeSplit uint16) (*types.Transaction, error) {
	return _ITaskMailbox.Contract.SetFeeSplit(&_ITaskMailbox.TransactOpts, feeSplit)
}

// SetFeeSplit is a paid mutator transaction binding the contract method 0x468c07a0.
//
// Solidity: function setFeeSplit(uint16 feeSplit) returns()
func (_ITaskMailbox *ITaskMailboxTransactorSession) SetFeeSplit(feeSplit uint16) (*types.Transaction, error) {
	return _ITaskMailbox.Contract.SetFeeSplit(&_ITaskMailbox.TransactOpts, feeSplit)
}

// SetFeeSplitCollector is a paid mutator transaction binding the contract method 0x678fbdb3.
//
// Solidity: function setFeeSplitCollector(address feeSplitCollector) returns()
func (_ITaskMailbox *ITaskMailboxTransactor) SetFeeSplitCollector(opts *bind.TransactOpts, feeSplitCollector common.Address) (*types.Transaction, error) {
	return _ITaskMailbox.contract.Transact(opts, "setFeeSplitCollector", feeSplitCollector)
}

// SetFeeSplitCollector is a paid mutator transaction binding the contract method 0x678fbdb3.
//
// Solidity: function setFeeSplitCollector(address feeSplitCollector) returns()
func (_ITaskMailbox *ITaskMailboxSession) SetFeeSplitCollector(feeSplitCollector common.Address) (*types.Transaction, error) {
	return _ITaskMailbox.Contract.SetFeeSplitCollector(&_ITaskMailbox.TransactOpts, feeSplitCollector)
}

// SetFeeSplitCollector is a paid mutator transaction binding the contract method 0x678fbdb3.
//
// Solidity: function setFeeSplitCollector(address feeSplitCollector) returns()
func (_ITaskMailbox *ITaskMailboxTransactorSession) SetFeeSplitCollector(feeSplitCollector common.Address) (*types.Transaction, error) {
	return _ITaskMailbox.Contract.SetFeeSplitCollector(&_ITaskMailbox.TransactOpts, feeSplitCollector)
}

// SubmitResult is a paid mutator transaction binding the contract method 0xa5fabc81.
//
// Solidity: function submitResult(bytes32 taskHash, bytes executorCert, bytes result) returns()
func (_ITaskMailbox *ITaskMailboxTransactor) SubmitResult(opts *bind.TransactOpts, taskHash [32]byte, executorCert []byte, result []byte) (*types.Transaction, error) {
	return _ITaskMailbox.contract.Transact(opts, "submitResult", taskHash, executorCert, result)
}

// SubmitResult is a paid mutator transaction binding the contract method 0xa5fabc81.
//
// Solidity: function submitResult(bytes32 taskHash, bytes executorCert, bytes result) returns()
func (_ITaskMailbox *ITaskMailboxSession) SubmitResult(taskHash [32]byte, executorCert []byte, result []byte) (*types.Transaction, error) {
	return _ITaskMailbox.Contract.SubmitResult(&_ITaskMailbox.TransactOpts, taskHash, executorCert, result)
}

// SubmitResult is a paid mutator transaction binding the contract method 0xa5fabc81.
//
// Solidity: function submitResult(bytes32 taskHash, bytes executorCert, bytes result) returns()
func (_ITaskMailbox *ITaskMailboxTransactorSession) SubmitResult(taskHash [32]byte, executorCert []byte, result []byte) (*types.Transaction, error) {
	return _ITaskMailbox.Contract.SubmitResult(&_ITaskMailbox.TransactOpts, taskHash, executorCert, result)
}

// ITaskMailboxExecutorOperatorSetRegisteredIterator is returned from FilterExecutorOperatorSetRegistered and is used to iterate over the raw logs and unpacked data for ExecutorOperatorSetRegistered events raised by the ITaskMailbox contract.
type ITaskMailboxExecutorOperatorSetRegisteredIterator struct {
	Event *ITaskMailboxExecutorOperatorSetRegistered // Event containing the contract specifics and raw log

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
func (it *ITaskMailboxExecutorOperatorSetRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITaskMailboxExecutorOperatorSetRegistered)
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
		it.Event = new(ITaskMailboxExecutorOperatorSetRegistered)
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
func (it *ITaskMailboxExecutorOperatorSetRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITaskMailboxExecutorOperatorSetRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITaskMailboxExecutorOperatorSetRegistered represents a ExecutorOperatorSetRegistered event raised by the ITaskMailbox contract.
type ITaskMailboxExecutorOperatorSetRegistered struct {
	Caller                common.Address
	Avs                   common.Address
	ExecutorOperatorSetId uint32
	IsRegistered          bool
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterExecutorOperatorSetRegistered is a free log retrieval operation binding the contract event 0x48b63f21a1eb9dd6880e196de6d7db3fbd0c282b74f1298dcb4cf53472298f39.
//
// Solidity: event ExecutorOperatorSetRegistered(address indexed caller, address indexed avs, uint32 indexed executorOperatorSetId, bool isRegistered)
func (_ITaskMailbox *ITaskMailboxFilterer) FilterExecutorOperatorSetRegistered(opts *bind.FilterOpts, caller []common.Address, avs []common.Address, executorOperatorSetId []uint32) (*ITaskMailboxExecutorOperatorSetRegisteredIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}
	var executorOperatorSetIdRule []interface{}
	for _, executorOperatorSetIdItem := range executorOperatorSetId {
		executorOperatorSetIdRule = append(executorOperatorSetIdRule, executorOperatorSetIdItem)
	}

	logs, sub, err := _ITaskMailbox.contract.FilterLogs(opts, "ExecutorOperatorSetRegistered", callerRule, avsRule, executorOperatorSetIdRule)
	if err != nil {
		return nil, err
	}
	return &ITaskMailboxExecutorOperatorSetRegisteredIterator{contract: _ITaskMailbox.contract, event: "ExecutorOperatorSetRegistered", logs: logs, sub: sub}, nil
}

// WatchExecutorOperatorSetRegistered is a free log subscription operation binding the contract event 0x48b63f21a1eb9dd6880e196de6d7db3fbd0c282b74f1298dcb4cf53472298f39.
//
// Solidity: event ExecutorOperatorSetRegistered(address indexed caller, address indexed avs, uint32 indexed executorOperatorSetId, bool isRegistered)
func (_ITaskMailbox *ITaskMailboxFilterer) WatchExecutorOperatorSetRegistered(opts *bind.WatchOpts, sink chan<- *ITaskMailboxExecutorOperatorSetRegistered, caller []common.Address, avs []common.Address, executorOperatorSetId []uint32) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}
	var executorOperatorSetIdRule []interface{}
	for _, executorOperatorSetIdItem := range executorOperatorSetId {
		executorOperatorSetIdRule = append(executorOperatorSetIdRule, executorOperatorSetIdItem)
	}

	logs, sub, err := _ITaskMailbox.contract.WatchLogs(opts, "ExecutorOperatorSetRegistered", callerRule, avsRule, executorOperatorSetIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITaskMailboxExecutorOperatorSetRegistered)
				if err := _ITaskMailbox.contract.UnpackLog(event, "ExecutorOperatorSetRegistered", log); err != nil {
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

// ParseExecutorOperatorSetRegistered is a log parse operation binding the contract event 0x48b63f21a1eb9dd6880e196de6d7db3fbd0c282b74f1298dcb4cf53472298f39.
//
// Solidity: event ExecutorOperatorSetRegistered(address indexed caller, address indexed avs, uint32 indexed executorOperatorSetId, bool isRegistered)
func (_ITaskMailbox *ITaskMailboxFilterer) ParseExecutorOperatorSetRegistered(log types.Log) (*ITaskMailboxExecutorOperatorSetRegistered, error) {
	event := new(ITaskMailboxExecutorOperatorSetRegistered)
	if err := _ITaskMailbox.contract.UnpackLog(event, "ExecutorOperatorSetRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITaskMailboxExecutorOperatorSetTaskConfigSetIterator is returned from FilterExecutorOperatorSetTaskConfigSet and is used to iterate over the raw logs and unpacked data for ExecutorOperatorSetTaskConfigSet events raised by the ITaskMailbox contract.
type ITaskMailboxExecutorOperatorSetTaskConfigSetIterator struct {
	Event *ITaskMailboxExecutorOperatorSetTaskConfigSet // Event containing the contract specifics and raw log

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
func (it *ITaskMailboxExecutorOperatorSetTaskConfigSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITaskMailboxExecutorOperatorSetTaskConfigSet)
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
		it.Event = new(ITaskMailboxExecutorOperatorSetTaskConfigSet)
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
func (it *ITaskMailboxExecutorOperatorSetTaskConfigSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITaskMailboxExecutorOperatorSetTaskConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITaskMailboxExecutorOperatorSetTaskConfigSet represents a ExecutorOperatorSetTaskConfigSet event raised by the ITaskMailbox contract.
type ITaskMailboxExecutorOperatorSetTaskConfigSet struct {
	Caller                common.Address
	Avs                   common.Address
	ExecutorOperatorSetId uint32
	Config                ITaskMailboxTypesExecutorOperatorSetTaskConfig
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterExecutorOperatorSetTaskConfigSet is a free log retrieval operation binding the contract event 0x7cd76abd4025a20959a1b20f7c1536e3894a0735cd8de0215dde803ddea7f2d2.
//
// Solidity: event ExecutorOperatorSetTaskConfigSet(address indexed caller, address indexed avs, uint32 indexed executorOperatorSetId, (address,uint96,address,uint8,address,(uint8,bytes),bytes) config)
func (_ITaskMailbox *ITaskMailboxFilterer) FilterExecutorOperatorSetTaskConfigSet(opts *bind.FilterOpts, caller []common.Address, avs []common.Address, executorOperatorSetId []uint32) (*ITaskMailboxExecutorOperatorSetTaskConfigSetIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}
	var executorOperatorSetIdRule []interface{}
	for _, executorOperatorSetIdItem := range executorOperatorSetId {
		executorOperatorSetIdRule = append(executorOperatorSetIdRule, executorOperatorSetIdItem)
	}

	logs, sub, err := _ITaskMailbox.contract.FilterLogs(opts, "ExecutorOperatorSetTaskConfigSet", callerRule, avsRule, executorOperatorSetIdRule)
	if err != nil {
		return nil, err
	}
	return &ITaskMailboxExecutorOperatorSetTaskConfigSetIterator{contract: _ITaskMailbox.contract, event: "ExecutorOperatorSetTaskConfigSet", logs: logs, sub: sub}, nil
}

// WatchExecutorOperatorSetTaskConfigSet is a free log subscription operation binding the contract event 0x7cd76abd4025a20959a1b20f7c1536e3894a0735cd8de0215dde803ddea7f2d2.
//
// Solidity: event ExecutorOperatorSetTaskConfigSet(address indexed caller, address indexed avs, uint32 indexed executorOperatorSetId, (address,uint96,address,uint8,address,(uint8,bytes),bytes) config)
func (_ITaskMailbox *ITaskMailboxFilterer) WatchExecutorOperatorSetTaskConfigSet(opts *bind.WatchOpts, sink chan<- *ITaskMailboxExecutorOperatorSetTaskConfigSet, caller []common.Address, avs []common.Address, executorOperatorSetId []uint32) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}
	var executorOperatorSetIdRule []interface{}
	for _, executorOperatorSetIdItem := range executorOperatorSetId {
		executorOperatorSetIdRule = append(executorOperatorSetIdRule, executorOperatorSetIdItem)
	}

	logs, sub, err := _ITaskMailbox.contract.WatchLogs(opts, "ExecutorOperatorSetTaskConfigSet", callerRule, avsRule, executorOperatorSetIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITaskMailboxExecutorOperatorSetTaskConfigSet)
				if err := _ITaskMailbox.contract.UnpackLog(event, "ExecutorOperatorSetTaskConfigSet", log); err != nil {
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

// ParseExecutorOperatorSetTaskConfigSet is a log parse operation binding the contract event 0x7cd76abd4025a20959a1b20f7c1536e3894a0735cd8de0215dde803ddea7f2d2.
//
// Solidity: event ExecutorOperatorSetTaskConfigSet(address indexed caller, address indexed avs, uint32 indexed executorOperatorSetId, (address,uint96,address,uint8,address,(uint8,bytes),bytes) config)
func (_ITaskMailbox *ITaskMailboxFilterer) ParseExecutorOperatorSetTaskConfigSet(log types.Log) (*ITaskMailboxExecutorOperatorSetTaskConfigSet, error) {
	event := new(ITaskMailboxExecutorOperatorSetTaskConfigSet)
	if err := _ITaskMailbox.contract.UnpackLog(event, "ExecutorOperatorSetTaskConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITaskMailboxFeeRefundedIterator is returned from FilterFeeRefunded and is used to iterate over the raw logs and unpacked data for FeeRefunded events raised by the ITaskMailbox contract.
type ITaskMailboxFeeRefundedIterator struct {
	Event *ITaskMailboxFeeRefunded // Event containing the contract specifics and raw log

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
func (it *ITaskMailboxFeeRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITaskMailboxFeeRefunded)
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
		it.Event = new(ITaskMailboxFeeRefunded)
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
func (it *ITaskMailboxFeeRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITaskMailboxFeeRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITaskMailboxFeeRefunded represents a FeeRefunded event raised by the ITaskMailbox contract.
type ITaskMailboxFeeRefunded struct {
	RefundCollector common.Address
	TaskHash        [32]byte
	AvsFee          *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterFeeRefunded is a free log retrieval operation binding the contract event 0xe3ed40d31808582f7a92a30beacc0ec788d5091407ec6c10c1b999b3f317aea3.
//
// Solidity: event FeeRefunded(address indexed refundCollector, bytes32 indexed taskHash, uint96 avsFee)
func (_ITaskMailbox *ITaskMailboxFilterer) FilterFeeRefunded(opts *bind.FilterOpts, refundCollector []common.Address, taskHash [][32]byte) (*ITaskMailboxFeeRefundedIterator, error) {

	var refundCollectorRule []interface{}
	for _, refundCollectorItem := range refundCollector {
		refundCollectorRule = append(refundCollectorRule, refundCollectorItem)
	}
	var taskHashRule []interface{}
	for _, taskHashItem := range taskHash {
		taskHashRule = append(taskHashRule, taskHashItem)
	}

	logs, sub, err := _ITaskMailbox.contract.FilterLogs(opts, "FeeRefunded", refundCollectorRule, taskHashRule)
	if err != nil {
		return nil, err
	}
	return &ITaskMailboxFeeRefundedIterator{contract: _ITaskMailbox.contract, event: "FeeRefunded", logs: logs, sub: sub}, nil
}

// WatchFeeRefunded is a free log subscription operation binding the contract event 0xe3ed40d31808582f7a92a30beacc0ec788d5091407ec6c10c1b999b3f317aea3.
//
// Solidity: event FeeRefunded(address indexed refundCollector, bytes32 indexed taskHash, uint96 avsFee)
func (_ITaskMailbox *ITaskMailboxFilterer) WatchFeeRefunded(opts *bind.WatchOpts, sink chan<- *ITaskMailboxFeeRefunded, refundCollector []common.Address, taskHash [][32]byte) (event.Subscription, error) {

	var refundCollectorRule []interface{}
	for _, refundCollectorItem := range refundCollector {
		refundCollectorRule = append(refundCollectorRule, refundCollectorItem)
	}
	var taskHashRule []interface{}
	for _, taskHashItem := range taskHash {
		taskHashRule = append(taskHashRule, taskHashItem)
	}

	logs, sub, err := _ITaskMailbox.contract.WatchLogs(opts, "FeeRefunded", refundCollectorRule, taskHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITaskMailboxFeeRefunded)
				if err := _ITaskMailbox.contract.UnpackLog(event, "FeeRefunded", log); err != nil {
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

// ParseFeeRefunded is a log parse operation binding the contract event 0xe3ed40d31808582f7a92a30beacc0ec788d5091407ec6c10c1b999b3f317aea3.
//
// Solidity: event FeeRefunded(address indexed refundCollector, bytes32 indexed taskHash, uint96 avsFee)
func (_ITaskMailbox *ITaskMailboxFilterer) ParseFeeRefunded(log types.Log) (*ITaskMailboxFeeRefunded, error) {
	event := new(ITaskMailboxFeeRefunded)
	if err := _ITaskMailbox.contract.UnpackLog(event, "FeeRefunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITaskMailboxFeeSplitCollectorSetIterator is returned from FilterFeeSplitCollectorSet and is used to iterate over the raw logs and unpacked data for FeeSplitCollectorSet events raised by the ITaskMailbox contract.
type ITaskMailboxFeeSplitCollectorSetIterator struct {
	Event *ITaskMailboxFeeSplitCollectorSet // Event containing the contract specifics and raw log

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
func (it *ITaskMailboxFeeSplitCollectorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITaskMailboxFeeSplitCollectorSet)
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
		it.Event = new(ITaskMailboxFeeSplitCollectorSet)
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
func (it *ITaskMailboxFeeSplitCollectorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITaskMailboxFeeSplitCollectorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITaskMailboxFeeSplitCollectorSet represents a FeeSplitCollectorSet event raised by the ITaskMailbox contract.
type ITaskMailboxFeeSplitCollectorSet struct {
	FeeSplitCollector common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterFeeSplitCollectorSet is a free log retrieval operation binding the contract event 0x262aa27c244f6f0088cb3092548a0adcaddedf459070a9ccab2dc6a07abe701d.
//
// Solidity: event FeeSplitCollectorSet(address indexed feeSplitCollector)
func (_ITaskMailbox *ITaskMailboxFilterer) FilterFeeSplitCollectorSet(opts *bind.FilterOpts, feeSplitCollector []common.Address) (*ITaskMailboxFeeSplitCollectorSetIterator, error) {

	var feeSplitCollectorRule []interface{}
	for _, feeSplitCollectorItem := range feeSplitCollector {
		feeSplitCollectorRule = append(feeSplitCollectorRule, feeSplitCollectorItem)
	}

	logs, sub, err := _ITaskMailbox.contract.FilterLogs(opts, "FeeSplitCollectorSet", feeSplitCollectorRule)
	if err != nil {
		return nil, err
	}
	return &ITaskMailboxFeeSplitCollectorSetIterator{contract: _ITaskMailbox.contract, event: "FeeSplitCollectorSet", logs: logs, sub: sub}, nil
}

// WatchFeeSplitCollectorSet is a free log subscription operation binding the contract event 0x262aa27c244f6f0088cb3092548a0adcaddedf459070a9ccab2dc6a07abe701d.
//
// Solidity: event FeeSplitCollectorSet(address indexed feeSplitCollector)
func (_ITaskMailbox *ITaskMailboxFilterer) WatchFeeSplitCollectorSet(opts *bind.WatchOpts, sink chan<- *ITaskMailboxFeeSplitCollectorSet, feeSplitCollector []common.Address) (event.Subscription, error) {

	var feeSplitCollectorRule []interface{}
	for _, feeSplitCollectorItem := range feeSplitCollector {
		feeSplitCollectorRule = append(feeSplitCollectorRule, feeSplitCollectorItem)
	}

	logs, sub, err := _ITaskMailbox.contract.WatchLogs(opts, "FeeSplitCollectorSet", feeSplitCollectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITaskMailboxFeeSplitCollectorSet)
				if err := _ITaskMailbox.contract.UnpackLog(event, "FeeSplitCollectorSet", log); err != nil {
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

// ParseFeeSplitCollectorSet is a log parse operation binding the contract event 0x262aa27c244f6f0088cb3092548a0adcaddedf459070a9ccab2dc6a07abe701d.
//
// Solidity: event FeeSplitCollectorSet(address indexed feeSplitCollector)
func (_ITaskMailbox *ITaskMailboxFilterer) ParseFeeSplitCollectorSet(log types.Log) (*ITaskMailboxFeeSplitCollectorSet, error) {
	event := new(ITaskMailboxFeeSplitCollectorSet)
	if err := _ITaskMailbox.contract.UnpackLog(event, "FeeSplitCollectorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITaskMailboxFeeSplitSetIterator is returned from FilterFeeSplitSet and is used to iterate over the raw logs and unpacked data for FeeSplitSet events raised by the ITaskMailbox contract.
type ITaskMailboxFeeSplitSetIterator struct {
	Event *ITaskMailboxFeeSplitSet // Event containing the contract specifics and raw log

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
func (it *ITaskMailboxFeeSplitSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITaskMailboxFeeSplitSet)
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
		it.Event = new(ITaskMailboxFeeSplitSet)
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
func (it *ITaskMailboxFeeSplitSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITaskMailboxFeeSplitSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITaskMailboxFeeSplitSet represents a FeeSplitSet event raised by the ITaskMailbox contract.
type ITaskMailboxFeeSplitSet struct {
	FeeSplit uint16
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFeeSplitSet is a free log retrieval operation binding the contract event 0x886b2cfcb151fd8b19ed902cc88f4a06dd9fe351a4a9ab93f33fe84abc157edf.
//
// Solidity: event FeeSplitSet(uint16 feeSplit)
func (_ITaskMailbox *ITaskMailboxFilterer) FilterFeeSplitSet(opts *bind.FilterOpts) (*ITaskMailboxFeeSplitSetIterator, error) {

	logs, sub, err := _ITaskMailbox.contract.FilterLogs(opts, "FeeSplitSet")
	if err != nil {
		return nil, err
	}
	return &ITaskMailboxFeeSplitSetIterator{contract: _ITaskMailbox.contract, event: "FeeSplitSet", logs: logs, sub: sub}, nil
}

// WatchFeeSplitSet is a free log subscription operation binding the contract event 0x886b2cfcb151fd8b19ed902cc88f4a06dd9fe351a4a9ab93f33fe84abc157edf.
//
// Solidity: event FeeSplitSet(uint16 feeSplit)
func (_ITaskMailbox *ITaskMailboxFilterer) WatchFeeSplitSet(opts *bind.WatchOpts, sink chan<- *ITaskMailboxFeeSplitSet) (event.Subscription, error) {

	logs, sub, err := _ITaskMailbox.contract.WatchLogs(opts, "FeeSplitSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITaskMailboxFeeSplitSet)
				if err := _ITaskMailbox.contract.UnpackLog(event, "FeeSplitSet", log); err != nil {
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

// ParseFeeSplitSet is a log parse operation binding the contract event 0x886b2cfcb151fd8b19ed902cc88f4a06dd9fe351a4a9ab93f33fe84abc157edf.
//
// Solidity: event FeeSplitSet(uint16 feeSplit)
func (_ITaskMailbox *ITaskMailboxFilterer) ParseFeeSplitSet(log types.Log) (*ITaskMailboxFeeSplitSet, error) {
	event := new(ITaskMailboxFeeSplitSet)
	if err := _ITaskMailbox.contract.UnpackLog(event, "FeeSplitSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITaskMailboxTaskCreatedIterator is returned from FilterTaskCreated and is used to iterate over the raw logs and unpacked data for TaskCreated events raised by the ITaskMailbox contract.
type ITaskMailboxTaskCreatedIterator struct {
	Event *ITaskMailboxTaskCreated // Event containing the contract specifics and raw log

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
func (it *ITaskMailboxTaskCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITaskMailboxTaskCreated)
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
		it.Event = new(ITaskMailboxTaskCreated)
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
func (it *ITaskMailboxTaskCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITaskMailboxTaskCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITaskMailboxTaskCreated represents a TaskCreated event raised by the ITaskMailbox contract.
type ITaskMailboxTaskCreated struct {
	Creator                         common.Address
	TaskHash                        [32]byte
	Avs                             common.Address
	ExecutorOperatorSetId           uint32
	OperatorTableReferenceTimestamp uint32
	RefundCollector                 common.Address
	AvsFee                          *big.Int
	TaskDeadline                    *big.Int
	Payload                         []byte
	Raw                             types.Log // Blockchain specific contextual infos
}

// FilterTaskCreated is a free log retrieval operation binding the contract event 0x33add0b01e02278be5459fbfa3274aee699ec47f4ee7236b59e7a2c8b5000c26.
//
// Solidity: event TaskCreated(address indexed creator, bytes32 indexed taskHash, address indexed avs, uint32 executorOperatorSetId, uint32 operatorTableReferenceTimestamp, address refundCollector, uint96 avsFee, uint256 taskDeadline, bytes payload)
func (_ITaskMailbox *ITaskMailboxFilterer) FilterTaskCreated(opts *bind.FilterOpts, creator []common.Address, taskHash [][32]byte, avs []common.Address) (*ITaskMailboxTaskCreatedIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var taskHashRule []interface{}
	for _, taskHashItem := range taskHash {
		taskHashRule = append(taskHashRule, taskHashItem)
	}
	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _ITaskMailbox.contract.FilterLogs(opts, "TaskCreated", creatorRule, taskHashRule, avsRule)
	if err != nil {
		return nil, err
	}
	return &ITaskMailboxTaskCreatedIterator{contract: _ITaskMailbox.contract, event: "TaskCreated", logs: logs, sub: sub}, nil
}

// WatchTaskCreated is a free log subscription operation binding the contract event 0x33add0b01e02278be5459fbfa3274aee699ec47f4ee7236b59e7a2c8b5000c26.
//
// Solidity: event TaskCreated(address indexed creator, bytes32 indexed taskHash, address indexed avs, uint32 executorOperatorSetId, uint32 operatorTableReferenceTimestamp, address refundCollector, uint96 avsFee, uint256 taskDeadline, bytes payload)
func (_ITaskMailbox *ITaskMailboxFilterer) WatchTaskCreated(opts *bind.WatchOpts, sink chan<- *ITaskMailboxTaskCreated, creator []common.Address, taskHash [][32]byte, avs []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var taskHashRule []interface{}
	for _, taskHashItem := range taskHash {
		taskHashRule = append(taskHashRule, taskHashItem)
	}
	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _ITaskMailbox.contract.WatchLogs(opts, "TaskCreated", creatorRule, taskHashRule, avsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITaskMailboxTaskCreated)
				if err := _ITaskMailbox.contract.UnpackLog(event, "TaskCreated", log); err != nil {
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

// ParseTaskCreated is a log parse operation binding the contract event 0x33add0b01e02278be5459fbfa3274aee699ec47f4ee7236b59e7a2c8b5000c26.
//
// Solidity: event TaskCreated(address indexed creator, bytes32 indexed taskHash, address indexed avs, uint32 executorOperatorSetId, uint32 operatorTableReferenceTimestamp, address refundCollector, uint96 avsFee, uint256 taskDeadline, bytes payload)
func (_ITaskMailbox *ITaskMailboxFilterer) ParseTaskCreated(log types.Log) (*ITaskMailboxTaskCreated, error) {
	event := new(ITaskMailboxTaskCreated)
	if err := _ITaskMailbox.contract.UnpackLog(event, "TaskCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITaskMailboxTaskVerifiedIterator is returned from FilterTaskVerified and is used to iterate over the raw logs and unpacked data for TaskVerified events raised by the ITaskMailbox contract.
type ITaskMailboxTaskVerifiedIterator struct {
	Event *ITaskMailboxTaskVerified // Event containing the contract specifics and raw log

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
func (it *ITaskMailboxTaskVerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITaskMailboxTaskVerified)
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
		it.Event = new(ITaskMailboxTaskVerified)
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
func (it *ITaskMailboxTaskVerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITaskMailboxTaskVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITaskMailboxTaskVerified represents a TaskVerified event raised by the ITaskMailbox contract.
type ITaskMailboxTaskVerified struct {
	Aggregator            common.Address
	TaskHash              [32]byte
	Avs                   common.Address
	ExecutorOperatorSetId uint32
	ExecutorCert          []byte
	Result                []byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterTaskVerified is a free log retrieval operation binding the contract event 0x659f23b2e7edf490e5fd6561c5148691ed0375ed7ddd3ab1bcfcfdbec4f209a9.
//
// Solidity: event TaskVerified(address indexed aggregator, bytes32 indexed taskHash, address indexed avs, uint32 executorOperatorSetId, bytes executorCert, bytes result)
func (_ITaskMailbox *ITaskMailboxFilterer) FilterTaskVerified(opts *bind.FilterOpts, aggregator []common.Address, taskHash [][32]byte, avs []common.Address) (*ITaskMailboxTaskVerifiedIterator, error) {

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}
	var taskHashRule []interface{}
	for _, taskHashItem := range taskHash {
		taskHashRule = append(taskHashRule, taskHashItem)
	}
	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _ITaskMailbox.contract.FilterLogs(opts, "TaskVerified", aggregatorRule, taskHashRule, avsRule)
	if err != nil {
		return nil, err
	}
	return &ITaskMailboxTaskVerifiedIterator{contract: _ITaskMailbox.contract, event: "TaskVerified", logs: logs, sub: sub}, nil
}

// WatchTaskVerified is a free log subscription operation binding the contract event 0x659f23b2e7edf490e5fd6561c5148691ed0375ed7ddd3ab1bcfcfdbec4f209a9.
//
// Solidity: event TaskVerified(address indexed aggregator, bytes32 indexed taskHash, address indexed avs, uint32 executorOperatorSetId, bytes executorCert, bytes result)
func (_ITaskMailbox *ITaskMailboxFilterer) WatchTaskVerified(opts *bind.WatchOpts, sink chan<- *ITaskMailboxTaskVerified, aggregator []common.Address, taskHash [][32]byte, avs []common.Address) (event.Subscription, error) {

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}
	var taskHashRule []interface{}
	for _, taskHashItem := range taskHash {
		taskHashRule = append(taskHashRule, taskHashItem)
	}
	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _ITaskMailbox.contract.WatchLogs(opts, "TaskVerified", aggregatorRule, taskHashRule, avsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITaskMailboxTaskVerified)
				if err := _ITaskMailbox.contract.UnpackLog(event, "TaskVerified", log); err != nil {
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

// ParseTaskVerified is a log parse operation binding the contract event 0x659f23b2e7edf490e5fd6561c5148691ed0375ed7ddd3ab1bcfcfdbec4f209a9.
//
// Solidity: event TaskVerified(address indexed aggregator, bytes32 indexed taskHash, address indexed avs, uint32 executorOperatorSetId, bytes executorCert, bytes result)
func (_ITaskMailbox *ITaskMailboxFilterer) ParseTaskVerified(log types.Log) (*ITaskMailboxTaskVerified, error) {
	event := new(ITaskMailboxTaskVerified)
	if err := _ITaskMailbox.contract.UnpackLog(event, "TaskVerified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
