// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package TaskMailbox

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

// TaskMailboxMetaData contains all meta data concerning the TaskMailbox contract.
var TaskMailboxMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_bn254CertificateVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_ecdsaCertificateVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_maxTaskSLA\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"_version\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"BN254_CERTIFICATE_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ECDSA_CERTIFICATE_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_TASK_SLA\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createTask\",\"inputs\":[{\"name\":\"taskParams\",\"type\":\"tuple\",\"internalType\":\"structITaskMailboxTypes.TaskParams\",\"components\":[{\"name\":\"refundCollector\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"executorOperatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"feeSplit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"feeSplitCollector\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBN254CertificateBytes\",\"inputs\":[{\"name\":\"cert\",\"type\":\"tuple\",\"internalType\":\"structIBN254CertificateVerifierTypes.BN254Certificate\",\"components\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"nonSignerWitnesses\",\"type\":\"tuple[]\",\"internalType\":\"structIBN254CertificateVerifierTypes.BN254OperatorInfoWitness[]\",\"components\":[{\"name\":\"operatorIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorInfoProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"operatorInfo\",\"type\":\"tuple\",\"internalType\":\"structIOperatorTableCalculatorTypes.BN254OperatorInfo\",\"components\":[{\"name\":\"pubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}]}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getECDSACertificateBytes\",\"inputs\":[{\"name\":\"cert\",\"type\":\"tuple\",\"internalType\":\"structIECDSACertificateVerifierTypes.ECDSACertificate\",\"components\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sig\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getExecutorOperatorSetTaskConfig\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structITaskMailboxTypes.ExecutorOperatorSetTaskConfig\",\"components\":[{\"name\":\"taskHook\",\"type\":\"address\",\"internalType\":\"contractIAVSTaskHook\"},{\"name\":\"taskSLA\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"curveType\",\"type\":\"uint8\",\"internalType\":\"enumIKeyRegistrarTypes.CurveType\"},{\"name\":\"feeCollector\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"consensus\",\"type\":\"tuple\",\"internalType\":\"structITaskMailboxTypes.Consensus\",\"components\":[{\"name\":\"consensusType\",\"type\":\"uint8\",\"internalType\":\"enumITaskMailboxTypes.ConsensusType\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"taskMetadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMessageHash\",\"inputs\":[{\"name\":\"taskHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"result\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getTaskInfo\",\"inputs\":[{\"name\":\"taskHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structITaskMailboxTypes.Task\",\"components\":[{\"name\":\"creator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"creationTime\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avsFee\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"refundCollector\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"executorOperatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeSplit\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumITaskMailboxTypes.TaskStatus\"},{\"name\":\"isFeeRefunded\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"operatorTableReferenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"executorOperatorSetTaskConfig\",\"type\":\"tuple\",\"internalType\":\"structITaskMailboxTypes.ExecutorOperatorSetTaskConfig\",\"components\":[{\"name\":\"taskHook\",\"type\":\"address\",\"internalType\":\"contractIAVSTaskHook\"},{\"name\":\"taskSLA\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"curveType\",\"type\":\"uint8\",\"internalType\":\"enumIKeyRegistrarTypes.CurveType\"},{\"name\":\"feeCollector\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"consensus\",\"type\":\"tuple\",\"internalType\":\"structITaskMailboxTypes.Consensus\",\"components\":[{\"name\":\"consensusType\",\"type\":\"uint8\",\"internalType\":\"enumITaskMailboxTypes.ConsensusType\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"taskMetadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"executorCert\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"result\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskResult\",\"inputs\":[{\"name\":\"taskHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskStatus\",\"inputs\":[{\"name\":\"taskHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumITaskMailboxTypes.TaskStatus\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_feeSplit\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"_feeSplitCollector\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isExecutorOperatorSetRegistered\",\"inputs\":[{\"name\":\"operatorSetKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"isRegistered\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"refundFee\",\"inputs\":[{\"name\":\"taskHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerExecutorOperatorSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"isRegistered\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setExecutorOperatorSetTaskConfig\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structITaskMailboxTypes.ExecutorOperatorSetTaskConfig\",\"components\":[{\"name\":\"taskHook\",\"type\":\"address\",\"internalType\":\"contractIAVSTaskHook\"},{\"name\":\"taskSLA\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"curveType\",\"type\":\"uint8\",\"internalType\":\"enumIKeyRegistrarTypes.CurveType\"},{\"name\":\"feeCollector\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"consensus\",\"type\":\"tuple\",\"internalType\":\"structITaskMailboxTypes.Consensus\",\"components\":[{\"name\":\"consensusType\",\"type\":\"uint8\",\"internalType\":\"enumITaskMailboxTypes.ConsensusType\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"taskMetadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFeeSplit\",\"inputs\":[{\"name\":\"_feeSplit\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFeeSplitCollector\",\"inputs\":[{\"name\":\"_feeSplitCollector\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitResult\",\"inputs\":[{\"name\":\"taskHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"executorCert\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"result\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"ExecutorOperatorSetRegistered\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"executorOperatorSetId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"isRegistered\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExecutorOperatorSetTaskConfigSet\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"executorOperatorSetId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structITaskMailboxTypes.ExecutorOperatorSetTaskConfig\",\"components\":[{\"name\":\"taskHook\",\"type\":\"address\",\"internalType\":\"contractIAVSTaskHook\"},{\"name\":\"taskSLA\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"curveType\",\"type\":\"uint8\",\"internalType\":\"enumIKeyRegistrarTypes.CurveType\"},{\"name\":\"feeCollector\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"consensus\",\"type\":\"tuple\",\"internalType\":\"structITaskMailboxTypes.Consensus\",\"components\":[{\"name\":\"consensusType\",\"type\":\"uint8\",\"internalType\":\"enumITaskMailboxTypes.ConsensusType\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"taskMetadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeRefunded\",\"inputs\":[{\"name\":\"refundCollector\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"taskHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"avsFee\",\"type\":\"uint96\",\"indexed\":false,\"internalType\":\"uint96\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeSplitCollectorSet\",\"inputs\":[{\"name\":\"feeSplitCollector\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeSplitSet\",\"inputs\":[{\"name\":\"feeSplit\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskCreated\",\"inputs\":[{\"name\":\"creator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"taskHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"executorOperatorSetId\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"operatorTableReferenceTimestamp\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"refundCollector\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"avsFee\",\"type\":\"uint96\",\"indexed\":false,\"internalType\":\"uint96\"},{\"name\":\"taskDeadline\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskVerified\",\"inputs\":[{\"name\":\"aggregator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"taskHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"executorOperatorSetId\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"executorCert\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"result\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CertificateStale\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EmptyCertificateSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExecutorOperatorSetNotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExecutorOperatorSetTaskConfigNotSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FeeAlreadyRefunded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidConsensusType\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidConsensusValue\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCurveType\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidFeeReceiver\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidFeeSplit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidMessageHash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperatorSetOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidReferenceTimestamp\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTaskCreator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTaskStatus\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"enumITaskMailboxTypes.TaskStatus\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"enumITaskMailboxTypes.TaskStatus\"}]},{\"type\":\"error\",\"name\":\"OnlyRefundCollector\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PayloadIsEmpty\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"TaskSLAExceedsMaximum\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ThresholdNotMet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TimestampAtCreation\",\"inputs\":[]}]",
	Bin: "0x610100604052348015610010575f5ffd5b50604051615d1b380380615d1b83398101604081905261002f9161019b565b6001600160a01b03808516608052831660a0526001600160601b03821660c052806100598161006e565b60e052506100656100b4565b505050506102e6565b5f5f829050601f815111156100a1578260405163305a27a960e01b8152600401610098919061028b565b60405180910390fd5b80516100ac826102c0565b179392505050565b5f54610100900460ff161561011b5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b6064820152608401610098565b5f5460ff9081161461016a575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b80516001600160a01b0381168114610182575f5ffd5b919050565b634e487b7160e01b5f52604160045260245ffd5b5f5f5f5f608085870312156101ae575f5ffd5b6101b78561016c565b93506101c56020860161016c565b60408601519093506001600160601b03811681146101e1575f5ffd5b60608601519092506001600160401b038111156101fc575f5ffd5b8501601f8101871361020c575f5ffd5b80516001600160401b0381111561022557610225610187565b604051601f8201601f19908116603f011681016001600160401b038111828210171561025357610253610187565b60405281815282820160200189101561026a575f5ffd5b8160208401602083015e5f6020838301015280935050505092959194509250565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b805160208083015191908110156102e0575f198160200360031b1b821691505b50919050565b60805160a05160c05160e0516159d161034a5f395f611ae301525f81816103670152612f2001525f818161027f0152818161328f015281816138a10152613ac401525f81816103cc0152818161324f015281816137cf01526139f101526159d15ff3fe608060405234801561000f575f5ffd5b506004361061016d575f3560e01c80636373ea69116100d9578063a5fabc8111610093578063f2fde38b1161006e578063f2fde38b146103a1578063f741e81a146103b4578063f7424fc9146103c7578063fa2c0b37146103ee575f5ffd5b8063a5fabc811461033c578063b86941661461034f578063d3e043aa14610362575f5ffd5b80636373ea69146102bc578063678fbdb3146102dd5780636bf6fad5146102f0578063708c0db914610310578063715018a6146103235780638da5cb5b1461032b575f5ffd5b8063468c07a01161012a578063468c07a01461023257806349acd884146102475780634ad52e021461025a57806354743ad21461027a57806354fd4d50146102a157806362fee037146102a9575f5ffd5b80631270a892146101715780631a20c5051461019a5780631ae370eb146101cb5780631fb66f5d146101de5780632bf6cc79146101ff57806337eaa1041461021f575b5f5ffd5b61018461017f366004614271565b610420565b6040516101919190614326565b60405180910390f35b609b546101b3906201000090046001600160a01b031681565b6040516001600160a01b039091168152602001610191565b6101846101d9366004614583565b610449565b6101f16101ec3660046146c2565b61045c565b604051908152602001610191565b61021261020d366004614737565b610e05565b6040516101919190614776565b6101f161022d366004614784565b6112b1565b6102456102403660046147d6565b6112e4565b005b6102456102553660046147fe565b6112f8565b61026d610268366004614737565b611543565b604051610191919061491a565b6101b37f000000000000000000000000000000000000000000000000000000000000000081565b610184611adc565b6101846102b7366004614737565b611b0c565b609b546102ca9061ffff1681565b60405161ffff9091168152602001610191565b6102456102eb366004614a65565b612009565b6103036102fe366004614a80565b61201a565b6040516101919190614a9a565b61024561031e366004614aac565b61222e565b61024561235f565b6033546001600160a01b03166101b3565b61024561034a366004614af4565b612372565b61024561035d366004614737565b612c50565b6103897f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160601b039091168152602001610191565b6102456103af366004614a65565b612e82565b6102456103c2366004614be6565b612ef8565b6101b37f000000000000000000000000000000000000000000000000000000000000000081565b6104106103fc366004614737565b60996020525f908152604090205460ff1681565b6040519015158152602001610191565b6060816040516020016104339190614d11565b6040516020818303038152906040529050919050565b6060816040516020016104339190614e95565b5f61046561312d565b5f8260400151511161048a57604051636b1a1b6960e11b815260040160405180910390fd5b60995f61049a8460200151613186565b815260208101919091526040015f205460ff166104ca5760405163c292b29760e01b815260040160405180910390fd5b5f609a5f6104db8560200151613186565b815260208082019290925260409081015f20815160e08101835281546001600160a01b038082168352600160a01b918290046001600160601b03169583019590955260018301549485169382019390935292909160608401910460ff1660028111156105495761054961474e565b600281111561055a5761055a61474e565b815260028201546001600160a01b0316602082015260408051808201825260038401805492909301929091829060ff16600181111561059b5761059b61474e565b60018111156105ac576105ac61474e565b81526020016001820180546105c090614ea7565b80601f01602080910402602001604051908101604052809291908181526020018280546105ec90614ea7565b80156106375780601f1061060e57610100808354040283529160200191610637565b820191905f5260205f20905b81548152906001019060200180831161061a57829003601f168201915b505050505081525050815260200160058201805461065490614ea7565b80601f016020809104026020016040519081016040528092919081815260200182805461068090614ea7565b80156106cb5780601f106106a2576101008083540402835291602001916106cb565b820191905f5260205f20905b8154815290600101906020018083116106ae57829003601f168201915b50505050508152505090506106df816131e9565b6106fc576040516314b0a41d60e11b815260040160405180910390fd5b5f61070a8260600151613232565b90505f816001600160a01b0316635ddb9b5b86602001516040518263ffffffff1660e01b815260040161073d9190614efd565b602060405180830381865afa158015610758573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061077c9190614f0b565b90505f826001600160a01b0316636141879e87602001516040518263ffffffff1660e01b81526004016107af9190614efd565b602060405180830381865afa1580156107ca573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107ee9190614f0b565b905063ffffffff8116158061082a57506108088183614f3a565b63ffffffff1684602001516001600160601b0316426108279190614f56565b11155b6108475760405163640fcd6b60e11b815260040160405180910390fd5b508251604051630a3fc61360e31b81526001600160a01b03909116906351fe3098906108799033908990600401614fa4565b5f6040518083038186803b15801561088f575f5ffd5b505afa1580156108a1573d5f5f3e3d5ffd5b50508451604051637036693f60e11b81525f93506001600160a01b03909116915063e06cd27e906108d6908990600401614fc7565b602060405180830381865afa1580156108f1573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109159190614fd9565b90505f6097543046896040516020016109319493929190614ff4565b60405160208183030381529060405280519060200120905060975460016109589190614f56565b609755604080516101c0810190915233815260208101610977426132cc565b6001600160601b03908116825260208a81018051516001600160a01b039081168386015292871660408501528b5190921660608401529051015163ffffffff166080820152609b5461ffff1660a082015260c001600181525f602080830182905263ffffffff80881660408086019190915260608086018c90528d8201516080808801919091528251808601845286815260a0808901919091528351808701855287815260c09889015289875260988652958390208851958901516001600160601b03908116600160a01b9081026001600160a01b03988916178355948a0151938a0151168402928616929092176001830155870151600282018054968901519789015161ffff16600160c01b0261ffff60c01b19989095169093026001600160c01b0319909616941693909317939093179384168117835560e0850151919391929160ff60d01b1990911662ffffff60c01b1990911617600160d01b836003811115610ae657610ae661474e565b02179055506101008201516002808301805461012086015163ffffffff16600160e01b026001600160e01b03941515600160d81b02949094166001600160d81b039091161792909217909155610140830151805160208201516001600160601b0316600160a01b9081026001600160a01b0392831617600386019081556040840151600487018054919094166001600160a01b031982168117855560608601519596929594936001600160a81b031990921617918490811115610bab57610bab61474e565b021790555060808201516002820180546001600160a01b0319166001600160a01b0390921691909117905560a08201518051600383018054909190829060ff191660018381811115610bff57610bff61474e565b021790555060208201516001820190610c189082615075565b50505060c08201516005820190610c2f9082615075565b5050506101608201516009820190610c479082615075565b50610180820151600a820190610c5d9082615075565b506101a0820151600b820190610c739082615075565b50505060408501516001600160a01b031615801590610c9a57505f826001600160601b0316115b15610d155760808501516001600160a01b0316610cca57604051633480121760e21b815260040160405180910390fd5b86516001600160a01b0316610cf257604051633480121760e21b815260040160405180910390fd5b6040850151610d15906001600160a01b031633306001600160601b038616613337565b8451604051629c5c4560e41b8152600481018390526001600160a01b03909116906309c5c450906024015f604051808303815f87803b158015610d56575f5ffd5b505af1158015610d68573d5f5f3e3d5ffd5b5050505086602001515f01516001600160a01b031681336001600160a01b03167f33add0b01e02278be5459fbfa3274aee699ec47f4ee7236b59e7a2c8b5000c268a6020015160200151878c5f0151888c602001516001600160601b031642610dd19190614f56565b8f60400151604051610de89695949392919061512f565b60405180910390a4945050505050610e006001606555565b919050565b5f81815260986020908152604080832081516101c08101835281546001600160a01b038082168352600160a01b918290046001600160601b03908116968401969096526001840154808216958401959095529381900490941660608201526002820154928316608082015292820463ffffffff1660a0840152600160c01b820461ffff1660c084015283929160e0830190600160d01b900460ff166003811115610eb157610eb161474e565b6003811115610ec257610ec261474e565b8152600282810154600160d81b810460ff9081161515602080860191909152600160e01b90920463ffffffff16604080860191909152805160e0810182526003870180546001600160a01b038082168452600160a01b918290046001600160601b03169684019690965260048901549586169383019390935260609687019691959094918601939290910490911690811115610f6057610f6061474e565b6002811115610f7157610f7161474e565b815260028201546001600160a01b0316602082015260408051808201825260038401805492909301929091829060ff166001811115610fb257610fb261474e565b6001811115610fc357610fc361474e565b8152602001600182018054610fd790614ea7565b80601f016020809104026020016040519081016040528092919081815260200182805461100390614ea7565b801561104e5780601f106110255761010080835404028352916020019161104e565b820191905f5260205f20905b81548152906001019060200180831161103157829003601f168201915b505050505081525050815260200160058201805461106b90614ea7565b80601f016020809104026020016040519081016040528092919081815260200182805461109790614ea7565b80156110e25780601f106110b9576101008083540402835291602001916110e2565b820191905f5260205f20905b8154815290600101906020018083116110c557829003601f168201915b50505050508152505081526020016009820180546110ff90614ea7565b80601f016020809104026020016040519081016040528092919081815260200182805461112b90614ea7565b80156111765780601f1061114d57610100808354040283529160200191611176565b820191905f5260205f20905b81548152906001019060200180831161115957829003601f168201915b50505050508152602001600a8201805461118f90614ea7565b80601f01602080910402602001604051908101604052809291908181526020018280546111bb90614ea7565b80156112065780601f106111dd57610100808354040283529160200191611206565b820191905f5260205f20905b8154815290600101906020018083116111e957829003601f168201915b50505050508152602001600b8201805461121f90614ea7565b80601f016020809104026020016040519081016040528092919081815260200182805461124b90614ea7565b80156112965780601f1061126d57610100808354040283529160200191611296565b820191905f5260205f20905b81548152906001019060200180831161127957829003601f168201915b50505050508152505090506112aa816133a9565b9392505050565b5f82826040516020016112c5929190615188565b6040516020818303038152906040528051906020012090505b92915050565b6112ec613404565b6112f58161345e565b50565b5f609a5f61130585613186565b815260208082019290925260409081015f20815160e08101835281546001600160a01b038082168352600160a01b918290046001600160601b03169583019590955260018301549485169382019390935292909160608401910460ff1660028111156113735761137361474e565b60028111156113845761138461474e565b815260028201546001600160a01b0316602082015260408051808201825260038401805492909301929091829060ff1660018111156113c5576113c561474e565b60018111156113d6576113d661474e565b81526020016001820180546113ea90614ea7565b80601f016020809104026020016040519081016040528092919081815260200182805461141690614ea7565b80156114615780601f1061143857610100808354040283529160200191611461565b820191905f5260205f20905b81548152906001019060200180831161144457829003601f168201915b505050505081525050815260200160058201805461147e90614ea7565b80601f01602080910402602001604051908101604052809291908181526020018280546114aa90614ea7565b80156114f55780601f106114cc576101008083540402835291602001916114f5565b820191905f5260205f20905b8154815290600101906020018083116114d857829003601f168201915b5050505050815250509050611509816131e9565b611526576040516314b0a41d60e11b815260040160405180910390fd5b6115348382606001516134cf565b61153e8383613570565b505050565b61154b614051565b5f82815260986020908152604080832081516101c08101835281546001600160a01b038082168352600160a01b918290046001600160601b03908116968401969096526001840154808216958401959095529381900490941660608201526002820154928316608082015292820463ffffffff1660a0840152600160c01b820461ffff1660c08401529060e0830190600160d01b900460ff1660038111156115f5576115f561474e565b60038111156116065761160661474e565b8152600282810154600160d81b810460ff9081161515602080860191909152600160e01b90920463ffffffff16604080860191909152805160e0810182526003870180546001600160a01b038082168452600160a01b918290046001600160601b031696840196909652600489015495861693830193909352606096870196919590949186019392909104909116908111156116a4576116a461474e565b60028111156116b5576116b561474e565b815260028201546001600160a01b0316602082015260408051808201825260038401805492909301929091829060ff1660018111156116f6576116f661474e565b60018111156117075761170761474e565b815260200160018201805461171b90614ea7565b80601f016020809104026020016040519081016040528092919081815260200182805461174790614ea7565b80156117925780601f1061176957610100808354040283529160200191611792565b820191905f5260205f20905b81548152906001019060200180831161177557829003601f168201915b50505050508152505081526020016005820180546117af90614ea7565b80601f01602080910402602001604051908101604052809291908181526020018280546117db90614ea7565b80156118265780601f106117fd57610100808354040283529160200191611826565b820191905f5260205f20905b81548152906001019060200180831161180957829003601f168201915b505050505081525050815260200160098201805461184390614ea7565b80601f016020809104026020016040519081016040528092919081815260200182805461186f90614ea7565b80156118ba5780601f10611891576101008083540402835291602001916118ba565b820191905f5260205f20905b81548152906001019060200180831161189d57829003601f168201915b50505050508152602001600a820180546118d390614ea7565b80601f01602080910402602001604051908101604052809291908181526020018280546118ff90614ea7565b801561194a5780601f106119215761010080835404028352916020019161194a565b820191905f5260205f20905b81548152906001019060200180831161192d57829003601f168201915b50505050508152602001600b8201805461196390614ea7565b80601f016020809104026020016040519081016040528092919081815260200182805461198f90614ea7565b80156119da5780601f106119b1576101008083540402835291602001916119da565b820191905f5260205f20905b8154815290600101906020018083116119bd57829003601f168201915b5050505050815250509050604051806101c00160405280825f01516001600160a01b0316815260200182602001516001600160601b0316815260200182604001516001600160a01b0316815260200182606001516001600160601b0316815260200182608001516001600160a01b031681526020018260a0015163ffffffff1681526020018260c0015161ffff168152602001611a76836133a9565b6003811115611a8757611a8761474e565b81526020018261010001511515815260200182610120015163ffffffff168152602001826101400151815260200182610160015181526020018261018001518152602001826101a00151815250915050919050565b6060611b077f0000000000000000000000000000000000000000000000000000000000000000613601565b905090565b5f81815260986020908152604080832081516101c08101835281546001600160a01b038082168352600160a01b918290046001600160601b0390811696840196909652600184015480821695840195909552938190049094166060808301919091526002830154938416608083015293830463ffffffff1660a0820152600160c01b830461ffff1660c08201529293929160e0830190600160d01b900460ff166003811115611bbd57611bbd61474e565b6003811115611bce57611bce61474e565b8152600282810154600160d81b810460ff9081161515602080860191909152600160e01b90920463ffffffff16604080860191909152805160e0810182526003870180546001600160a01b038082168452600160a01b918290046001600160601b03169684019690965260048901549586169383019390935260609687019691959094918601939290910490911690811115611c6c57611c6c61474e565b6002811115611c7d57611c7d61474e565b815260028201546001600160a01b0316602082015260408051808201825260038401805492909301929091829060ff166001811115611cbe57611cbe61474e565b6001811115611ccf57611ccf61474e565b8152602001600182018054611ce390614ea7565b80601f0160208091040260200160405190810160405280929190818152602001828054611d0f90614ea7565b8015611d5a5780601f10611d3157610100808354040283529160200191611d5a565b820191905f5260205f20905b815481529060010190602001808311611d3d57829003601f168201915b5050505050815250508152602001600582018054611d7790614ea7565b80601f0160208091040260200160405190810160405280929190818152602001828054611da390614ea7565b8015611dee5780601f10611dc557610100808354040283529160200191611dee565b820191905f5260205f20905b815481529060010190602001808311611dd157829003601f168201915b5050505050815250508152602001600982018054611e0b90614ea7565b80601f0160208091040260200160405190810160405280929190818152602001828054611e3790614ea7565b8015611e825780601f10611e5957610100808354040283529160200191611e82565b820191905f5260205f20905b815481529060010190602001808311611e6557829003601f168201915b50505050508152602001600a82018054611e9b90614ea7565b80601f0160208091040260200160405190810160405280929190818152602001828054611ec790614ea7565b8015611f125780601f10611ee957610100808354040283529160200191611f12565b820191905f5260205f20905b815481529060010190602001808311611ef557829003601f168201915b50505050508152602001600b82018054611f2b90614ea7565b80601f0160208091040260200160405190810160405280929190818152602001828054611f5790614ea7565b8015611fa25780601f10611f7957610100808354040283529160200191611fa2565b820191905f5260205f20905b815481529060010190602001808311611f8557829003601f168201915b50505050508152505090505f611fb7826133a9565b90506002816003811115611fcd57611fcd61474e565b146002829091611ffb57604051634091b18960e11b8152600401611ff29291906151a0565b60405180910390fd5b5050506101a0015192915050565b612011613404565b6112f58161363e565b6120226140c4565b609a5f61202e84613186565b815260208082019290925260409081015f20815160e08101835281546001600160a01b038082168352600160a01b918290046001600160601b03169583019590955260018301549485169382019390935292909160608401910460ff16600281111561209c5761209c61474e565b60028111156120ad576120ad61474e565b815260028201546001600160a01b0316602082015260408051808201825260038401805492909301929091829060ff1660018111156120ee576120ee61474e565b60018111156120ff576120ff61474e565b815260200160018201805461211390614ea7565b80601f016020809104026020016040519081016040528092919081815260200182805461213f90614ea7565b801561218a5780601f106121615761010080835404028352916020019161218a565b820191905f5260205f20905b81548152906001019060200180831161216d57829003601f168201915b50505050508152505081526020016005820180546121a790614ea7565b80601f01602080910402602001604051908101604052809291908181526020018280546121d390614ea7565b801561221e5780601f106121f55761010080835404028352916020019161221e565b820191905f5260205f20905b81548152906001019060200180831161220157829003601f168201915b5050505050815250509050919050565b5f54610100900460ff161580801561224c57505f54600160ff909116105b806122655750303b15801561226557505f5460ff166001145b6122c85760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401611ff2565b5f805460ff1916600117905580156122e9575f805461ff0019166101001790555b6122f16136b8565b6122f96136e6565b61230284613714565b61230b8361345e565b6123148261363e565b8015612359575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050565b612367613404565b6123705f613714565b565b61237a61312d565b5f83815260986020908152604080832081516101c08101835281546001600160a01b038082168352600160a01b918290046001600160601b03908116968401969096526001840154808216958401959095529381900490941660608201526002820154928316608082015292820463ffffffff1660a0840152600160c01b820461ffff1660c0840152929161281e91849060e0830190600160d01b900460ff16600381111561242b5761242b61474e565b600381111561243c5761243c61474e565b8152600282810154600160d81b810460ff9081161515602080860191909152600160e01b90920463ffffffff16604080860191909152805160e0810182526003870180546001600160a01b038082168452600160a01b918290046001600160601b031696840196909652600489015495861693830193909352606096870196919590949186019392909104909116908111156124da576124da61474e565b60028111156124eb576124eb61474e565b815260028201546001600160a01b0316602082015260408051808201825260038401805492909301929091829060ff16600181111561252c5761252c61474e565b600181111561253d5761253d61474e565b815260200160018201805461255190614ea7565b80601f016020809104026020016040519081016040528092919081815260200182805461257d90614ea7565b80156125c85780601f1061259f576101008083540402835291602001916125c8565b820191905f5260205f20905b8154815290600101906020018083116125ab57829003601f168201915b50505050508152505081526020016005820180546125e590614ea7565b80601f016020809104026020016040519081016040528092919081815260200182805461261190614ea7565b801561265c5780601f106126335761010080835404028352916020019161265c565b820191905f5260205f20905b81548152906001019060200180831161263f57829003601f168201915b505050505081525050815260200160098201805461267990614ea7565b80601f01602080910402602001604051908101604052809291908181526020018280546126a590614ea7565b80156126f05780601f106126c7576101008083540402835291602001916126f0565b820191905f5260205f20905b8154815290600101906020018083116126d357829003601f168201915b50505050508152602001600a8201805461270990614ea7565b80601f016020809104026020016040519081016040528092919081815260200182805461273590614ea7565b80156127805780601f1061275757610100808354040283529160200191612780565b820191905f5260205f20905b81548152906001019060200180831161276357829003601f168201915b50505050508152602001600b8201805461279990614ea7565b80601f01602080910402602001604051908101604052809291908181526020018280546127c590614ea7565b80156128105780601f106127e757610100808354040283529160200191612810565b820191905f5260205f20905b8154815290600101906020018083116127f357829003601f168201915b5050505050815250506133a9565b905060018160038111156128345761283461474e565b14600182909161285957604051634091b18960e11b8152600401611ff29291906151a0565b50508154600160a01b90046001600160601b0316421161288c5760405163015a4b7560e51b815260040160405180910390fd5b600382015460405163ba33565d60e01b81526001600160a01b039091169063ba33565d906128c49033908990899089906004016151bb565b5f6040518083038186803b1580156128da575f5ffd5b505afa1580156128ec573d5f5f3e3d5ffd5b50506040805180820182526001808701546001600160a01b03168252600287015463ffffffff600160a01b9182900416602084015260048801548451808601909552600689018054949750612a28965060ff929091048216949390928492919091169081111561295e5761295e61474e565b600181111561296f5761296f61474e565b815260200160018201805461298390614ea7565b80601f01602080910402602001604051908101604052809291908181526020018280546129af90614ea7565b80156129fa5780601f106129d1576101008083540402835291602001916129fa565b820191905f5260205f20905b8154815290600101906020018083116129dd57829003601f168201915b5050509190925250505060028601548490600160e01b900463ffffffff16612a228b8a6112b1565b8a613765565b60028301805460ff60d01b1916600160d11b179055600a8301612a4b8682615075565b50600b8301612a5a8582615075565b5060048301546001600160a01b031615801590612a8a57506001830154600160a01b90046001600160601b031615155b15612b7957600283015460018401545f91612ad99161271091612aca91600160c01b90910461ffff16906001600160601b03600160a01b909104166151f3565b612ad4919061520a565b6132cc565b90506001600160601b03811615612b1957609b546004850154612b19916001600160a01b039182169162010000909104166001600160601b038416613b87565b60018401545f90612b3b908390600160a01b90046001600160601b0316615229565b90506001600160601b03811615612b765760058501546004860154612b76916001600160a01b0391821691166001600160601b038416613b87565b50505b6003830154604051637041233f60e11b8152336004820152602481018890526001600160a01b039091169063e082467e906044015f604051808303815f87803b158015612bc4575f5ffd5b505af1158015612bd6573d5f5f3e3d5ffd5b505050600184015460028501546040516001600160a01b039092169250889133917f659f23b2e7edf490e5fd6561c5148691ed0375ed7ddd3ab1bcfcfdbec4f209a991612c3b9163ffffffff600160a01b9091041690600a8a0190600b8b01906152c7565b60405180910390a450505061153e6001606555565b612c5861312d565b5f81815260986020526040902060028101546001600160a01b03163314612c92576040516370f43cb760e01b815260040160405180910390fd5b6002810154600160d81b900460ff1615612cbf57604051633e3d786960e01b815260040160405180910390fd5b604080516101c08101825282546001600160a01b038082168352600160a01b918290046001600160601b03908116602085015260018601548083169585019590955293829004909316606083015260028401549283166080830152820463ffffffff1660a0820152600160c01b820461ffff1660c08201525f91612d609190849060e0830190600160d01b900460ff16600381111561242b5761242b61474e565b90506003816003811115612d7657612d7661474e565b146003829091612d9b57604051634091b18960e11b8152600401611ff29291906151a0565b505060028201805460ff60d81b1916600160d81b17905560048201546001600160a01b031615801590612de157506001820154600160a01b90046001600160601b031615155b15612e1d57600282015460018301546004840154612e1d926001600160a01b0391821692911690600160a01b90046001600160601b0316613b87565b60028201546001830154604051600160a01b9091046001600160601b0316815284916001600160a01b0316907fe3ed40d31808582f7a92a30beacc0ec788d5091407ec6c10c1b999b3f317aea39060200160405180910390a350506112f56001606555565b612e8a613404565b6001600160a01b038116612eef5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401611ff2565b6112f581613714565b612f01816131e9565b612f1e576040516314b0a41d60e11b815260040160405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160601b031681602001516001600160601b03161115612f755760405163a7cf966560e01b815260040160405180910390fd5b612f828160a00151613bb7565b612f908282606001516134cf565b80609a5f612f9d85613186565b815260208082019290925260409081015f208351928401516001600160601b0316600160a01b9081026001600160a01b0394851617825591840151600182018054919094166001600160a01b03198216811785556060860151929492936001600160a81b0319909216179083600281111561301a5761301a61474e565b021790555060808201516002820180546001600160a01b0319166001600160a01b0390921691909117905560a08201518051600383018054909190829060ff19166001838181111561306e5761306e61474e565b0217905550602082015160018201906130879082615075565b50505060c0820151600582019061309e9082615075565b50905050816020015163ffffffff16825f01516001600160a01b0316336001600160a01b03167f7cd76abd4025a20959a1b20f7c1536e3894a0735cd8de0215dde803ddea7f2d2846040516130f39190614a9a565b60405180910390a460995f61310784613186565b815260208101919091526040015f205460ff1661312957613129826001613570565b5050565b60026065540361317f5760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401611ff2565b6002606555565b5f815f0151826020015163ffffffff166040516020016131d192919060609290921b6bffffffffffffffffffffffff1916825260a01b6001600160a01b031916601482015260200190565b6040516020818303038152906040526112de906152f7565b5f80826060015160028111156132015761320161474e565b14158015613218575081516001600160a01b031615155b80156112de575050602001516001600160601b0316151590565b5f60028260028111156132475761324761474e565b0361327357507f0000000000000000000000000000000000000000000000000000000000000000919050565b60018260028111156132875761328761474e565b036132b357507f0000000000000000000000000000000000000000000000000000000000000000919050565b60405163fdea7c0960e01b815260040160405180910390fd5b5f6001600160601b038211156133335760405162461bcd60e51b815260206004820152602660248201527f53616665436173743a2076616c756520646f65736e27742066697420696e203960448201526536206269747360d01b6064820152608401611ff2565b5090565b6040516001600160a01b03808516602483015283166044820152606481018290526123599085906323b872dd60e01b906084015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b031990931692909217909152613c76565b6001606555565b5f60018260e0015160038111156133c2576133c261474e565b1480156133ef57508161014001516020015182602001516133e3919061531a565b6001600160601b031642115b156133fc57506003919050565b5060e0015190565b6033546001600160a01b031633146123705760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401611ff2565b61271061ffff8216111561348557604051630601f69760e01b815260040160405180910390fd5b609b805461ffff191661ffff83169081179091556040519081527f886b2cfcb151fd8b19ed902cc88f4a06dd9fe351a4a9ab93f33fe84abc157edf9060200160405180910390a150565b5f6134d982613232565b6040516304240c4960e51b815290915033906001600160a01b0383169063848189209061350a908790600401614efd565b602060405180830381865afa158015613525573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906135499190615339565b6001600160a01b03161461153e576040516342ecfee960e11b815260040160405180910390fd5b8060995f61357d85613186565b81526020019081526020015f205f6101000a81548160ff021916908315150217905550816020015163ffffffff16825f01516001600160a01b0316336001600160a01b03167f48b63f21a1eb9dd6880e196de6d7db3fbd0c282b74f1298dcb4cf53472298f39846040516135f5911515815260200190565b60405180910390a45050565b60605f61360d83613d49565b6040805160208082528183019092529192505f91906020820181803683375050509182525060208101929092525090565b6001600160a01b03811661366557604051630863a45360e11b815260040160405180910390fd5b609b805462010000600160b01b031916620100006001600160a01b038416908102919091179091556040517f262aa27c244f6f0088cb3092548a0adcaddedf459070a9ccab2dc6a07abe701d905f90a250565b5f54610100900460ff166136de5760405162461bcd60e51b8152600401611ff290615354565b612370613d70565b5f54610100900460ff1661370c5760405162461bcd60e51b8152600401611ff290615354565b612370613d9f565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b5f855160018111156137795761377961474e565b036139215760028660028111156137925761379261474e565b0361384f575f818060200190518101906137ac91906155ff565b90506137b9818585613dc5565b6040516280b71560e41b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063080b71509061380690889085906004016156d1565b5f604051808303815f875af1158015613821573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261384891908101906156f0565b5050613b7f565b60018660028111156138635761386361474e565b036132b3575f8180602001905181019061387d9190615721565b905061388a818585613e4f565b6040516380c7d3f360e01b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906380c7d3f3906138d8908890859060040161579a565b5f60405180830381865afa1580156138f2573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052613919919081019061581d565b505050613b7f565b6001855160018111156139365761393661474e565b03613b66575f85602001518060200190518101906139549190615876565b6040805160018082528183019092529192505f91906020808301908036833701905050905081815f8151811061398c5761398c615891565b61ffff909216602092830291909101909101525f60028960028111156139b4576139b461474e565b03613a72575f848060200190518101906139ce91906155ff565b90506139db818888613dc5565b604051625f5e5d60e21b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063017d797490613a2a908b90859088906004016158e3565b6020604051808303815f875af1158015613a46573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613a6a9190615914565b915050613b48565b6001896002811115613a8657613a8661474e565b036132b3575f84806020019051810190613aa09190615721565b9050613aad818888613e4f565b604051630606d12160e51b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063c0da242090613afd908b908590889060040161592f565b5f60405180830381865afa158015613b17573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052613b3e919081019061594e565b509150613b489050565b80613919576040516359fa4a9360e01b815260040160405180910390fd5b6040516347d3772160e11b815260040160405180910390fd5b505050505050565b6040516001600160a01b03831660248201526044810182905261153e90849063a9059cbb60e01b9060640161336b565b5f81516001811115613bcb57613bcb61474e565b03613bf457602081015151156112f557604051631501e04760e21b815260040160405180910390fd5b600181516001811115613c0957613c0961474e565b03613b6657806020015151602014613c3457604051631501e04760e21b815260040160405180910390fd5b5f8160200151806020019051810190613c4d9190615876565b905061271061ffff8216111561312957604051631501e04760e21b815260040160405180910390fd5b5f613cca826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316613ec79092919063ffffffff16565b905080515f1480613cea575080806020019051810190613cea9190615914565b61153e5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608401611ff2565b5f60ff8216601f8111156112de57604051632cd44ac360e21b815260040160405180910390fd5b5f54610100900460ff16613d965760405162461bcd60e51b8152600401611ff290615354565b61237033613714565b5f54610100900460ff166133a25760405162461bcd60e51b8152600401611ff290615354565b8163ffffffff16835f015163ffffffff1614613df457604051634534032960e01b815260040160405180910390fd5b80836020015114613e1857604051638b56642d60e01b815260040160405180910390fd5b604083015151158015613e315750604083015160200151155b1561153e57604051637a8a1dbd60e11b815260040160405180910390fd5b8163ffffffff16835f015163ffffffff1614613e7e57604051634534032960e01b815260040160405180910390fd5b80836020015114613ea257604051638b56642d60e01b815260040160405180910390fd5b5f8360400151511161153e57604051637a8a1dbd60e11b815260040160405180910390fd5b6060613ed584845f85613edd565b949350505050565b606082471015613f3e5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b6064820152608401611ff2565b5f5f866001600160a01b03168587604051613f599190615985565b5f6040518083038185875af1925050503d805f8114613f93576040519150601f19603f3d011682016040523d82523d5f602084013e613f98565b606091505b5091509150613fa987838387613fb4565b979650505050505050565b606083156140225782515f0361401b576001600160a01b0385163b61401b5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401611ff2565b5081613ed5565b613ed583838151156140375781518083602001fd5b8060405162461bcd60e51b8152600401611ff29190614326565b604080516101c0810182525f80825260208201819052918101829052606081018290526080810182905260a0810182905260c081018290529060e082019081525f6020820181905260408201526060016140a96140c4565b81526020016060815260200160608152602001606081525090565b6040805160e0810182525f8082526020820181905291810182905290606082019081525f60208201526040016140f8614105565b8152602001606081525090565b60408051808201909152805f6140f8565b634e487b7160e01b5f52604160045260245ffd5b604051606081016001600160401b038111828210171561414c5761414c614116565b60405290565b604080519081016001600160401b038111828210171561414c5761414c614116565b60405160a081016001600160401b038111828210171561414c5761414c614116565b60405160e081016001600160401b038111828210171561414c5761414c614116565b604051601f8201601f191681016001600160401b03811182821017156141e0576141e0614116565b604052919050565b63ffffffff811681146112f5575f5ffd5b5f6001600160401b0382111561421157614211614116565b50601f01601f191660200190565b5f82601f83011261422e575f5ffd5b813561424161423c826141f9565b6141b8565b818152846020838601011115614255575f5ffd5b816020850160208301375f918101602001919091529392505050565b5f60208284031215614281575f5ffd5b81356001600160401b03811115614296575f5ffd5b8201606081850312156142a7575f5ffd5b6142af61412a565b81356142ba816141e8565b81526020828101359082015260408201356001600160401b038111156142de575f5ffd5b6142ea8682850161421f565b604083015250949350505050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f6112aa60208301846142f8565b5f60408284031215614348575f5ffd5b614350614152565b823581526020928301359281019290925250919050565b5f82601f830112614376575f5ffd5b61438060406141b8565b806040840185811115614391575f5ffd5b845b818110156143ab578035845260209384019301614393565b509095945050505050565b5f6001600160401b038211156143ce576143ce614116565b5060051b60200190565b5f82601f8301126143e7575f5ffd5b81356143f561423c826143b6565b8082825260208201915060208360051b860101925085831115614416575f5ffd5b602085015b838110156145795780356001600160401b03811115614438575f5ffd5b86016060818903601f1901121561444d575f5ffd5b61445561412a565b6020820135614463816141e8565b815260408201356001600160401b0381111561447d575f5ffd5b61448c8a60208386010161421f565b60208301525060608201356001600160401b038111156144aa575f5ffd5b6020818401019250506060828a0312156144c2575f5ffd5b6144ca614152565b6144d48a84614338565b815260408301356001600160401b038111156144ee575f5ffd5b80840193505089601f840112614502575f5ffd5b823561451061423c826143b6565b8082825260208201915060208360051b87010192508c831115614531575f5ffd5b6020860195505b82861015614553578535825260209586019590910190614538565b80602085015250505080604083015250808552505060208301925060208101905061441b565b5095945050505050565b5f60208284031215614593575f5ffd5b81356001600160401b038111156145a8575f5ffd5b82018084036101208112156145bb575f5ffd5b6145c3614174565b82356145ce816141e8565b8152602083810135908201526145e78660408501614338565b60408201526080607f19830112156145fd575f5ffd5b614605614152565b91506146148660808501614367565b82526146238660c08501614367565b602083015281606082015261010083013591506001600160401b03821115614649575f5ffd5b614655868385016143d8565b608082015295945050505050565b6001600160a01b03811681146112f5575f5ffd5b8035610e0081614663565b5f60408284031215614692575f5ffd5b61469a614152565b905081356146a781614663565b815260208201356146b7816141e8565b602082015292915050565b5f602082840312156146d2575f5ffd5b81356001600160401b038111156146e7575f5ffd5b8201608081850312156146f8575f5ffd5b61470061412a565b813561470b81614663565b815261471a8560208401614682565b602082015260608201356001600160401b038111156142de575f5ffd5b5f60208284031215614747575f5ffd5b5035919050565b634e487b7160e01b5f52602160045260245ffd5b600481106147725761477261474e565b9052565b602081016112de8284614762565b5f5f60408385031215614795575f5ffd5b8235915060208301356001600160401b038111156147b1575f5ffd5b6147bd8582860161421f565b9150509250929050565b61ffff811681146112f5575f5ffd5b5f602082840312156147e6575f5ffd5b81356112aa816147c7565b80151581146112f5575f5ffd5b5f5f6060838503121561480f575f5ffd5b6148198484614682565b91506040830135614829816147f1565b809150509250929050565b600381106147725761477261474e565b5f8151600281106148575761485761474e565b80845250602082015160406020850152613ed560408501826142f8565b80516001600160a01b031682526020808201516001600160601b0316908301526040808201515f916148b0908501826001600160a01b03169052565b5060608201516148c36060850182614834565b5060808201516148de60808501826001600160a01b03169052565b5060a082015160e060a08501526148f860e0850182614844565b905060c083015184820360c086015261491182826142f8565b95945050505050565b602081526149346020820183516001600160a01b03169052565b5f602083015161494f60408401826001600160601b03169052565b5060408301516001600160a01b03811660608401525060608301516001600160601b03811660808401525060808301516001600160a01b03811660a08401525060a083015163ffffffff811660c08401525060c083015161ffff811660e08401525060e08301516149c4610100840182614762565b506101008301518015156101208401525061012083015163ffffffff8116610140840152506101408301516101c0610160840152614a066101e0840182614874565b9050610160840151601f1984830301610180850152614a2582826142f8565b915050610180840151601f19848303016101a0850152614a4582826142f8565b9150506101a0840151601f19848303016101c085015261491182826142f8565b5f60208284031215614a75575f5ffd5b81356112aa81614663565b5f60408284031215614a90575f5ffd5b6112aa8383614682565b602081525f6112aa6020830184614874565b5f5f5f60608486031215614abe575f5ffd5b8335614ac981614663565b92506020840135614ad9816147c7565b91506040840135614ae981614663565b809150509250925092565b5f5f5f60608486031215614b06575f5ffd5b8335925060208401356001600160401b03811115614b22575f5ffd5b614b2e8682870161421f565b92505060408401356001600160401b03811115614b49575f5ffd5b614b558682870161421f565b9150509250925092565b6001600160601b03811681146112f5575f5ffd5b8035610e0081614b5f565b803560038110610e00575f5ffd5b5f60408284031215614b9c575f5ffd5b614ba4614152565b9050813560028110614bb4575f5ffd5b815260208201356001600160401b03811115614bce575f5ffd5b614bda8482850161421f565b60208301525092915050565b5f5f60608385031215614bf7575f5ffd5b614c018484614682565b915060408301356001600160401b03811115614c1b575f5ffd5b830160e08186031215614c2c575f5ffd5b614c34614196565b614c3d82614677565b8152614c4b60208301614b73565b6020820152614c5c60408301614677565b6040820152614c6d60608301614b7e565b6060820152614c7e60808301614677565b608082015260a08201356001600160401b03811115614c9b575f5ffd5b614ca787828501614b8c565b60a08301525060c08201356001600160401b03811115614cc5575f5ffd5b614cd18782850161421f565b60c08301525080925050509250929050565b63ffffffff8151168252602081015160208301525f604082015160606040850152613ed560608501826142f8565b602081525f6112aa6020830184614ce3565b805f5b6002811015612359578151845260209384019390910190600101614d26565b5f610120830163ffffffff8351168452602083015160208501526040830151614d7b604086018280518252602090810151910152565b506060830151614d8f608086018251614d23565b60200151614da060c0860182614d23565b506080830151610120610100860152818151808452610140870191506101408160051b88010193506020830192505f5b81811015614e895761013f19888603018352835163ffffffff8151168652602081015160606020880152614e0760608801826142f8565b905060408201519150868103604088015260608101614e3182845180518252602090810151910152565b6020928301516060604084015280518083529301925f92608001905b80841015614e705784518252602082019150602085019450600184019350614e4d565b5097505050602094850194939093019250600101614dd0565b50929695505050505050565b602081525f6112aa6020830184614d45565b600181811c90821680614ebb57607f821691505b602082108103614ed957634e487b7160e01b5f52602260045260245ffd5b50919050565b80516001600160a01b0316825260209081015163ffffffff16910152565b604081016112de8284614edf565b5f60208284031215614f1b575f5ffd5b81516112aa816141e8565b634e487b7160e01b5f52601160045260245ffd5b63ffffffff81811683821601908111156112de576112de614f26565b808201808211156112de576112de614f26565b80516001600160a01b031682526020808201515f91614f8a90850182614edf565b50604082015160806060850152613ed560808501826142f8565b6001600160a01b03831681526040602082018190525f90613ed590830184614f69565b602081525f6112aa6020830184614f69565b5f60208284031215614fe9575f5ffd5b81516112aa81614b5f565b84815260018060a01b0384166020820152826040820152608060608201525f6150206080830184614f69565b9695505050505050565b601f82111561153e57805f5260205f20601f840160051c8101602085101561504f5750805b601f840160051c820191505b8181101561506e575f815560010161505b565b5050505050565b81516001600160401b0381111561508e5761508e614116565b6150a28161509c8454614ea7565b8461502a565b6020601f8211600181146150d4575f83156150bd5750848201515b5f19600385901b1c1916600184901b17845561506e565b5f84815260208120601f198516915b8281101561510357878501518255602094850194600190920191016150e3565b508482101561512057868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b63ffffffff8716815263ffffffff8616602082015260018060a01b03851660408201526001600160601b038416606082015282608082015260c060a08201525f61517c60c08301846142f8565b98975050505050505050565b828152604060208201525f613ed560408301846142f8565b604081016151ae8285614762565b6112aa6020830184614762565b60018060a01b0385168152836020820152608060408201525f6151e160808301856142f8565b8281036060840152613fa981856142f8565b80820281158282048414176112de576112de614f26565b5f8261522457634e487b7160e01b5f52601260045260245ffd5b500490565b6001600160601b0382811682821603908111156112de576112de614f26565b5f815461525481614ea7565b80855260018216801561526e576001811461528a576152be565b60ff1983166020870152602082151560051b87010193506152be565b845f5260205f205f5b838110156152b55781546020828a010152600182019150602081019050615293565b87016020019450505b50505092915050565b63ffffffff84168152606060208201525f6152e56060830185615248565b82810360408401526150208185615248565b80516020808301519190811015614ed9575f1960209190910360031b1b16919050565b6001600160601b0381811683821601908111156112de576112de614f26565b5f60208284031215615349575f5ffd5b81516112aa81614663565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b5f604082840312156153af575f5ffd5b6153b7614152565b825181526020928301519281019290925250919050565b5f82601f8301126153dd575f5ffd5b6153e760406141b8565b8060408401858111156153f8575f5ffd5b845b818110156143ab5780518452602093840193016153fa565b5f82601f830112615421575f5ffd5b815161542f61423c826141f9565b818152846020838601011115615443575f5ffd5b8160208501602083015e5f918101602001919091529392505050565b5f82601f83011261546e575f5ffd5b815161547c61423c826143b6565b8082825260208201915060208360051b86010192508583111561549d575f5ffd5b602085015b838110156145795780518352602092830192016154a2565b5f82601f8301126154c9575f5ffd5b81516154d761423c826143b6565b8082825260208201915060208360051b8601019250858311156154f8575f5ffd5b602085015b838110156145795780516001600160401b0381111561551a575f5ffd5b86016060818903601f1901121561552f575f5ffd5b61553761412a565b6020820151615545816141e8565b815260408201516001600160401b0381111561555f575f5ffd5b61556e8a602083860101615412565b60208301525060608201516001600160401b0381111561558c575f5ffd5b6020818401019250506060828a0312156155a4575f5ffd5b6155ac614152565b6155b68a8461539f565b815260408301516001600160401b038111156155d0575f5ffd5b6155dc8b82860161545f565b6020830152508060408301525080855250506020830192506020810190506154fd565b5f6020828403121561560f575f5ffd5b81516001600160401b03811115615624575f5ffd5b8201808403610120811215615637575f5ffd5b61563f614174565b825161564a816141e8565b815260208381015190820152615663866040850161539f565b60408201526080607f1983011215615679575f5ffd5b615681614152565b915061569086608085016153ce565b825261569f8660c085016153ce565b602083015281606082015261010083015191506001600160401b038211156156c5575f5ffd5b614655868385016154ba565b6156db8184614edf565b606060408201525f613ed56060830184614d45565b5f60208284031215615700575f5ffd5b81516001600160401b03811115615715575f5ffd5b613ed58482850161545f565b5f60208284031215615731575f5ffd5b81516001600160401b03811115615746575f5ffd5b820160608185031215615757575f5ffd5b61575f61412a565b815161576a816141e8565b81526020828101519082015260408201516001600160401b0381111561578e575f5ffd5b6142ea86828501615412565b6157a48184614edf565b606060408201525f613ed56060830184614ce3565b5f82601f8301126157c8575f5ffd5b81516157d661423c826143b6565b8082825260208201915060208360051b8601019250858311156157f7575f5ffd5b602085015b8381101561457957805161580f81614663565b8352602092830192016157fc565b5f5f6040838503121561582e575f5ffd5b82516001600160401b03811115615843575f5ffd5b61584f8582860161545f565b92505060208301516001600160401b0381111561586a575f5ffd5b6147bd858286016157b9565b5f60208284031215615886575f5ffd5b81516112aa816147c7565b634e487b7160e01b5f52603260045260245ffd5b5f8151808452602084019350602083015f5b828110156158d957815161ffff168652602095860195909101906001016158b7565b5093949350505050565b6158ed8185614edf565b608060408201525f6159026080830185614d45565b828103606084015261502081856158a5565b5f60208284031215615924575f5ffd5b81516112aa816147f1565b6159398185614edf565b608060408201525f6159026080830185614ce3565b5f5f6040838503121561595f575f5ffd5b825161596a816147f1565b60208401519092506001600160401b0381111561586a575f5ffd5b5f82518060208501845e5f92019182525091905056fea264697066735822122099f18466f6b06e40b9c04a196f23aa8ca171882301c7541ecb6407796eee978764736f6c634300081b0033",
}

// TaskMailboxABI is the input ABI used to generate the binding from.
// Deprecated: Use TaskMailboxMetaData.ABI instead.
var TaskMailboxABI = TaskMailboxMetaData.ABI

// TaskMailboxBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TaskMailboxMetaData.Bin instead.
var TaskMailboxBin = TaskMailboxMetaData.Bin

// DeployTaskMailbox deploys a new Ethereum contract, binding an instance of TaskMailbox to it.
func DeployTaskMailbox(auth *bind.TransactOpts, backend bind.ContractBackend, _bn254CertificateVerifier common.Address, _ecdsaCertificateVerifier common.Address, _maxTaskSLA *big.Int, _version string) (common.Address, *types.Transaction, *TaskMailbox, error) {
	parsed, err := TaskMailboxMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TaskMailboxBin), backend, _bn254CertificateVerifier, _ecdsaCertificateVerifier, _maxTaskSLA, _version)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TaskMailbox{TaskMailboxCaller: TaskMailboxCaller{contract: contract}, TaskMailboxTransactor: TaskMailboxTransactor{contract: contract}, TaskMailboxFilterer: TaskMailboxFilterer{contract: contract}}, nil
}

// TaskMailbox is an auto generated Go binding around an Ethereum contract.
type TaskMailbox struct {
	TaskMailboxCaller     // Read-only binding to the contract
	TaskMailboxTransactor // Write-only binding to the contract
	TaskMailboxFilterer   // Log filterer for contract events
}

// TaskMailboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type TaskMailboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskMailboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TaskMailboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskMailboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TaskMailboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskMailboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TaskMailboxSession struct {
	Contract     *TaskMailbox      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TaskMailboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TaskMailboxCallerSession struct {
	Contract *TaskMailboxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// TaskMailboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TaskMailboxTransactorSession struct {
	Contract     *TaskMailboxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TaskMailboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type TaskMailboxRaw struct {
	Contract *TaskMailbox // Generic contract binding to access the raw methods on
}

// TaskMailboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TaskMailboxCallerRaw struct {
	Contract *TaskMailboxCaller // Generic read-only contract binding to access the raw methods on
}

// TaskMailboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TaskMailboxTransactorRaw struct {
	Contract *TaskMailboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTaskMailbox creates a new instance of TaskMailbox, bound to a specific deployed contract.
func NewTaskMailbox(address common.Address, backend bind.ContractBackend) (*TaskMailbox, error) {
	contract, err := bindTaskMailbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TaskMailbox{TaskMailboxCaller: TaskMailboxCaller{contract: contract}, TaskMailboxTransactor: TaskMailboxTransactor{contract: contract}, TaskMailboxFilterer: TaskMailboxFilterer{contract: contract}}, nil
}

// NewTaskMailboxCaller creates a new read-only instance of TaskMailbox, bound to a specific deployed contract.
func NewTaskMailboxCaller(address common.Address, caller bind.ContractCaller) (*TaskMailboxCaller, error) {
	contract, err := bindTaskMailbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TaskMailboxCaller{contract: contract}, nil
}

// NewTaskMailboxTransactor creates a new write-only instance of TaskMailbox, bound to a specific deployed contract.
func NewTaskMailboxTransactor(address common.Address, transactor bind.ContractTransactor) (*TaskMailboxTransactor, error) {
	contract, err := bindTaskMailbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TaskMailboxTransactor{contract: contract}, nil
}

// NewTaskMailboxFilterer creates a new log filterer instance of TaskMailbox, bound to a specific deployed contract.
func NewTaskMailboxFilterer(address common.Address, filterer bind.ContractFilterer) (*TaskMailboxFilterer, error) {
	contract, err := bindTaskMailbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TaskMailboxFilterer{contract: contract}, nil
}

// bindTaskMailbox binds a generic wrapper to an already deployed contract.
func bindTaskMailbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TaskMailboxMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TaskMailbox *TaskMailboxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TaskMailbox.Contract.TaskMailboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TaskMailbox *TaskMailboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaskMailbox.Contract.TaskMailboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TaskMailbox *TaskMailboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TaskMailbox.Contract.TaskMailboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TaskMailbox *TaskMailboxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TaskMailbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TaskMailbox *TaskMailboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaskMailbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TaskMailbox *TaskMailboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TaskMailbox.Contract.contract.Transact(opts, method, params...)
}

// BN254CERTIFICATEVERIFIER is a free data retrieval call binding the contract method 0xf7424fc9.
//
// Solidity: function BN254_CERTIFICATE_VERIFIER() view returns(address)
func (_TaskMailbox *TaskMailboxCaller) BN254CERTIFICATEVERIFIER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TaskMailbox.contract.Call(opts, &out, "BN254_CERTIFICATE_VERIFIER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BN254CERTIFICATEVERIFIER is a free data retrieval call binding the contract method 0xf7424fc9.
//
// Solidity: function BN254_CERTIFICATE_VERIFIER() view returns(address)
func (_TaskMailbox *TaskMailboxSession) BN254CERTIFICATEVERIFIER() (common.Address, error) {
	return _TaskMailbox.Contract.BN254CERTIFICATEVERIFIER(&_TaskMailbox.CallOpts)
}

// BN254CERTIFICATEVERIFIER is a free data retrieval call binding the contract method 0xf7424fc9.
//
// Solidity: function BN254_CERTIFICATE_VERIFIER() view returns(address)
func (_TaskMailbox *TaskMailboxCallerSession) BN254CERTIFICATEVERIFIER() (common.Address, error) {
	return _TaskMailbox.Contract.BN254CERTIFICATEVERIFIER(&_TaskMailbox.CallOpts)
}

// ECDSACERTIFICATEVERIFIER is a free data retrieval call binding the contract method 0x54743ad2.
//
// Solidity: function ECDSA_CERTIFICATE_VERIFIER() view returns(address)
func (_TaskMailbox *TaskMailboxCaller) ECDSACERTIFICATEVERIFIER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TaskMailbox.contract.Call(opts, &out, "ECDSA_CERTIFICATE_VERIFIER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ECDSACERTIFICATEVERIFIER is a free data retrieval call binding the contract method 0x54743ad2.
//
// Solidity: function ECDSA_CERTIFICATE_VERIFIER() view returns(address)
func (_TaskMailbox *TaskMailboxSession) ECDSACERTIFICATEVERIFIER() (common.Address, error) {
	return _TaskMailbox.Contract.ECDSACERTIFICATEVERIFIER(&_TaskMailbox.CallOpts)
}

// ECDSACERTIFICATEVERIFIER is a free data retrieval call binding the contract method 0x54743ad2.
//
// Solidity: function ECDSA_CERTIFICATE_VERIFIER() view returns(address)
func (_TaskMailbox *TaskMailboxCallerSession) ECDSACERTIFICATEVERIFIER() (common.Address, error) {
	return _TaskMailbox.Contract.ECDSACERTIFICATEVERIFIER(&_TaskMailbox.CallOpts)
}

// MAXTASKSLA is a free data retrieval call binding the contract method 0xd3e043aa.
//
// Solidity: function MAX_TASK_SLA() view returns(uint96)
func (_TaskMailbox *TaskMailboxCaller) MAXTASKSLA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TaskMailbox.contract.Call(opts, &out, "MAX_TASK_SLA")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXTASKSLA is a free data retrieval call binding the contract method 0xd3e043aa.
//
// Solidity: function MAX_TASK_SLA() view returns(uint96)
func (_TaskMailbox *TaskMailboxSession) MAXTASKSLA() (*big.Int, error) {
	return _TaskMailbox.Contract.MAXTASKSLA(&_TaskMailbox.CallOpts)
}

// MAXTASKSLA is a free data retrieval call binding the contract method 0xd3e043aa.
//
// Solidity: function MAX_TASK_SLA() view returns(uint96)
func (_TaskMailbox *TaskMailboxCallerSession) MAXTASKSLA() (*big.Int, error) {
	return _TaskMailbox.Contract.MAXTASKSLA(&_TaskMailbox.CallOpts)
}

// FeeSplit is a free data retrieval call binding the contract method 0x6373ea69.
//
// Solidity: function feeSplit() view returns(uint16)
func (_TaskMailbox *TaskMailboxCaller) FeeSplit(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _TaskMailbox.contract.Call(opts, &out, "feeSplit")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// FeeSplit is a free data retrieval call binding the contract method 0x6373ea69.
//
// Solidity: function feeSplit() view returns(uint16)
func (_TaskMailbox *TaskMailboxSession) FeeSplit() (uint16, error) {
	return _TaskMailbox.Contract.FeeSplit(&_TaskMailbox.CallOpts)
}

// FeeSplit is a free data retrieval call binding the contract method 0x6373ea69.
//
// Solidity: function feeSplit() view returns(uint16)
func (_TaskMailbox *TaskMailboxCallerSession) FeeSplit() (uint16, error) {
	return _TaskMailbox.Contract.FeeSplit(&_TaskMailbox.CallOpts)
}

// FeeSplitCollector is a free data retrieval call binding the contract method 0x1a20c505.
//
// Solidity: function feeSplitCollector() view returns(address)
func (_TaskMailbox *TaskMailboxCaller) FeeSplitCollector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TaskMailbox.contract.Call(opts, &out, "feeSplitCollector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeSplitCollector is a free data retrieval call binding the contract method 0x1a20c505.
//
// Solidity: function feeSplitCollector() view returns(address)
func (_TaskMailbox *TaskMailboxSession) FeeSplitCollector() (common.Address, error) {
	return _TaskMailbox.Contract.FeeSplitCollector(&_TaskMailbox.CallOpts)
}

// FeeSplitCollector is a free data retrieval call binding the contract method 0x1a20c505.
//
// Solidity: function feeSplitCollector() view returns(address)
func (_TaskMailbox *TaskMailboxCallerSession) FeeSplitCollector() (common.Address, error) {
	return _TaskMailbox.Contract.FeeSplitCollector(&_TaskMailbox.CallOpts)
}

// GetBN254CertificateBytes is a free data retrieval call binding the contract method 0x1ae370eb.
//
// Solidity: function getBN254CertificateBytes((uint32,bytes32,(uint256,uint256),(uint256[2],uint256[2]),(uint32,bytes,((uint256,uint256),uint256[]))[]) cert) pure returns(bytes)
func (_TaskMailbox *TaskMailboxCaller) GetBN254CertificateBytes(opts *bind.CallOpts, cert IBN254CertificateVerifierTypesBN254Certificate) ([]byte, error) {
	var out []interface{}
	err := _TaskMailbox.contract.Call(opts, &out, "getBN254CertificateBytes", cert)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetBN254CertificateBytes is a free data retrieval call binding the contract method 0x1ae370eb.
//
// Solidity: function getBN254CertificateBytes((uint32,bytes32,(uint256,uint256),(uint256[2],uint256[2]),(uint32,bytes,((uint256,uint256),uint256[]))[]) cert) pure returns(bytes)
func (_TaskMailbox *TaskMailboxSession) GetBN254CertificateBytes(cert IBN254CertificateVerifierTypesBN254Certificate) ([]byte, error) {
	return _TaskMailbox.Contract.GetBN254CertificateBytes(&_TaskMailbox.CallOpts, cert)
}

// GetBN254CertificateBytes is a free data retrieval call binding the contract method 0x1ae370eb.
//
// Solidity: function getBN254CertificateBytes((uint32,bytes32,(uint256,uint256),(uint256[2],uint256[2]),(uint32,bytes,((uint256,uint256),uint256[]))[]) cert) pure returns(bytes)
func (_TaskMailbox *TaskMailboxCallerSession) GetBN254CertificateBytes(cert IBN254CertificateVerifierTypesBN254Certificate) ([]byte, error) {
	return _TaskMailbox.Contract.GetBN254CertificateBytes(&_TaskMailbox.CallOpts, cert)
}

// GetECDSACertificateBytes is a free data retrieval call binding the contract method 0x1270a892.
//
// Solidity: function getECDSACertificateBytes((uint32,bytes32,bytes) cert) pure returns(bytes)
func (_TaskMailbox *TaskMailboxCaller) GetECDSACertificateBytes(opts *bind.CallOpts, cert IECDSACertificateVerifierTypesECDSACertificate) ([]byte, error) {
	var out []interface{}
	err := _TaskMailbox.contract.Call(opts, &out, "getECDSACertificateBytes", cert)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetECDSACertificateBytes is a free data retrieval call binding the contract method 0x1270a892.
//
// Solidity: function getECDSACertificateBytes((uint32,bytes32,bytes) cert) pure returns(bytes)
func (_TaskMailbox *TaskMailboxSession) GetECDSACertificateBytes(cert IECDSACertificateVerifierTypesECDSACertificate) ([]byte, error) {
	return _TaskMailbox.Contract.GetECDSACertificateBytes(&_TaskMailbox.CallOpts, cert)
}

// GetECDSACertificateBytes is a free data retrieval call binding the contract method 0x1270a892.
//
// Solidity: function getECDSACertificateBytes((uint32,bytes32,bytes) cert) pure returns(bytes)
func (_TaskMailbox *TaskMailboxCallerSession) GetECDSACertificateBytes(cert IECDSACertificateVerifierTypesECDSACertificate) ([]byte, error) {
	return _TaskMailbox.Contract.GetECDSACertificateBytes(&_TaskMailbox.CallOpts, cert)
}

// GetExecutorOperatorSetTaskConfig is a free data retrieval call binding the contract method 0x6bf6fad5.
//
// Solidity: function getExecutorOperatorSetTaskConfig((address,uint32) operatorSet) view returns((address,uint96,address,uint8,address,(uint8,bytes),bytes))
func (_TaskMailbox *TaskMailboxCaller) GetExecutorOperatorSetTaskConfig(opts *bind.CallOpts, operatorSet OperatorSet) (ITaskMailboxTypesExecutorOperatorSetTaskConfig, error) {
	var out []interface{}
	err := _TaskMailbox.contract.Call(opts, &out, "getExecutorOperatorSetTaskConfig", operatorSet)

	if err != nil {
		return *new(ITaskMailboxTypesExecutorOperatorSetTaskConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(ITaskMailboxTypesExecutorOperatorSetTaskConfig)).(*ITaskMailboxTypesExecutorOperatorSetTaskConfig)

	return out0, err

}

// GetExecutorOperatorSetTaskConfig is a free data retrieval call binding the contract method 0x6bf6fad5.
//
// Solidity: function getExecutorOperatorSetTaskConfig((address,uint32) operatorSet) view returns((address,uint96,address,uint8,address,(uint8,bytes),bytes))
func (_TaskMailbox *TaskMailboxSession) GetExecutorOperatorSetTaskConfig(operatorSet OperatorSet) (ITaskMailboxTypesExecutorOperatorSetTaskConfig, error) {
	return _TaskMailbox.Contract.GetExecutorOperatorSetTaskConfig(&_TaskMailbox.CallOpts, operatorSet)
}

// GetExecutorOperatorSetTaskConfig is a free data retrieval call binding the contract method 0x6bf6fad5.
//
// Solidity: function getExecutorOperatorSetTaskConfig((address,uint32) operatorSet) view returns((address,uint96,address,uint8,address,(uint8,bytes),bytes))
func (_TaskMailbox *TaskMailboxCallerSession) GetExecutorOperatorSetTaskConfig(operatorSet OperatorSet) (ITaskMailboxTypesExecutorOperatorSetTaskConfig, error) {
	return _TaskMailbox.Contract.GetExecutorOperatorSetTaskConfig(&_TaskMailbox.CallOpts, operatorSet)
}

// GetMessageHash is a free data retrieval call binding the contract method 0x37eaa104.
//
// Solidity: function getMessageHash(bytes32 taskHash, bytes result) pure returns(bytes32)
func (_TaskMailbox *TaskMailboxCaller) GetMessageHash(opts *bind.CallOpts, taskHash [32]byte, result []byte) ([32]byte, error) {
	var out []interface{}
	err := _TaskMailbox.contract.Call(opts, &out, "getMessageHash", taskHash, result)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetMessageHash is a free data retrieval call binding the contract method 0x37eaa104.
//
// Solidity: function getMessageHash(bytes32 taskHash, bytes result) pure returns(bytes32)
func (_TaskMailbox *TaskMailboxSession) GetMessageHash(taskHash [32]byte, result []byte) ([32]byte, error) {
	return _TaskMailbox.Contract.GetMessageHash(&_TaskMailbox.CallOpts, taskHash, result)
}

// GetMessageHash is a free data retrieval call binding the contract method 0x37eaa104.
//
// Solidity: function getMessageHash(bytes32 taskHash, bytes result) pure returns(bytes32)
func (_TaskMailbox *TaskMailboxCallerSession) GetMessageHash(taskHash [32]byte, result []byte) ([32]byte, error) {
	return _TaskMailbox.Contract.GetMessageHash(&_TaskMailbox.CallOpts, taskHash, result)
}

// GetTaskInfo is a free data retrieval call binding the contract method 0x4ad52e02.
//
// Solidity: function getTaskInfo(bytes32 taskHash) view returns((address,uint96,address,uint96,address,uint32,uint16,uint8,bool,uint32,(address,uint96,address,uint8,address,(uint8,bytes),bytes),bytes,bytes,bytes))
func (_TaskMailbox *TaskMailboxCaller) GetTaskInfo(opts *bind.CallOpts, taskHash [32]byte) (ITaskMailboxTypesTask, error) {
	var out []interface{}
	err := _TaskMailbox.contract.Call(opts, &out, "getTaskInfo", taskHash)

	if err != nil {
		return *new(ITaskMailboxTypesTask), err
	}

	out0 := *abi.ConvertType(out[0], new(ITaskMailboxTypesTask)).(*ITaskMailboxTypesTask)

	return out0, err

}

// GetTaskInfo is a free data retrieval call binding the contract method 0x4ad52e02.
//
// Solidity: function getTaskInfo(bytes32 taskHash) view returns((address,uint96,address,uint96,address,uint32,uint16,uint8,bool,uint32,(address,uint96,address,uint8,address,(uint8,bytes),bytes),bytes,bytes,bytes))
func (_TaskMailbox *TaskMailboxSession) GetTaskInfo(taskHash [32]byte) (ITaskMailboxTypesTask, error) {
	return _TaskMailbox.Contract.GetTaskInfo(&_TaskMailbox.CallOpts, taskHash)
}

// GetTaskInfo is a free data retrieval call binding the contract method 0x4ad52e02.
//
// Solidity: function getTaskInfo(bytes32 taskHash) view returns((address,uint96,address,uint96,address,uint32,uint16,uint8,bool,uint32,(address,uint96,address,uint8,address,(uint8,bytes),bytes),bytes,bytes,bytes))
func (_TaskMailbox *TaskMailboxCallerSession) GetTaskInfo(taskHash [32]byte) (ITaskMailboxTypesTask, error) {
	return _TaskMailbox.Contract.GetTaskInfo(&_TaskMailbox.CallOpts, taskHash)
}

// GetTaskResult is a free data retrieval call binding the contract method 0x62fee037.
//
// Solidity: function getTaskResult(bytes32 taskHash) view returns(bytes)
func (_TaskMailbox *TaskMailboxCaller) GetTaskResult(opts *bind.CallOpts, taskHash [32]byte) ([]byte, error) {
	var out []interface{}
	err := _TaskMailbox.contract.Call(opts, &out, "getTaskResult", taskHash)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetTaskResult is a free data retrieval call binding the contract method 0x62fee037.
//
// Solidity: function getTaskResult(bytes32 taskHash) view returns(bytes)
func (_TaskMailbox *TaskMailboxSession) GetTaskResult(taskHash [32]byte) ([]byte, error) {
	return _TaskMailbox.Contract.GetTaskResult(&_TaskMailbox.CallOpts, taskHash)
}

// GetTaskResult is a free data retrieval call binding the contract method 0x62fee037.
//
// Solidity: function getTaskResult(bytes32 taskHash) view returns(bytes)
func (_TaskMailbox *TaskMailboxCallerSession) GetTaskResult(taskHash [32]byte) ([]byte, error) {
	return _TaskMailbox.Contract.GetTaskResult(&_TaskMailbox.CallOpts, taskHash)
}

// GetTaskStatus is a free data retrieval call binding the contract method 0x2bf6cc79.
//
// Solidity: function getTaskStatus(bytes32 taskHash) view returns(uint8)
func (_TaskMailbox *TaskMailboxCaller) GetTaskStatus(opts *bind.CallOpts, taskHash [32]byte) (uint8, error) {
	var out []interface{}
	err := _TaskMailbox.contract.Call(opts, &out, "getTaskStatus", taskHash)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetTaskStatus is a free data retrieval call binding the contract method 0x2bf6cc79.
//
// Solidity: function getTaskStatus(bytes32 taskHash) view returns(uint8)
func (_TaskMailbox *TaskMailboxSession) GetTaskStatus(taskHash [32]byte) (uint8, error) {
	return _TaskMailbox.Contract.GetTaskStatus(&_TaskMailbox.CallOpts, taskHash)
}

// GetTaskStatus is a free data retrieval call binding the contract method 0x2bf6cc79.
//
// Solidity: function getTaskStatus(bytes32 taskHash) view returns(uint8)
func (_TaskMailbox *TaskMailboxCallerSession) GetTaskStatus(taskHash [32]byte) (uint8, error) {
	return _TaskMailbox.Contract.GetTaskStatus(&_TaskMailbox.CallOpts, taskHash)
}

// IsExecutorOperatorSetRegistered is a free data retrieval call binding the contract method 0xfa2c0b37.
//
// Solidity: function isExecutorOperatorSetRegistered(bytes32 operatorSetKey) view returns(bool isRegistered)
func (_TaskMailbox *TaskMailboxCaller) IsExecutorOperatorSetRegistered(opts *bind.CallOpts, operatorSetKey [32]byte) (bool, error) {
	var out []interface{}
	err := _TaskMailbox.contract.Call(opts, &out, "isExecutorOperatorSetRegistered", operatorSetKey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExecutorOperatorSetRegistered is a free data retrieval call binding the contract method 0xfa2c0b37.
//
// Solidity: function isExecutorOperatorSetRegistered(bytes32 operatorSetKey) view returns(bool isRegistered)
func (_TaskMailbox *TaskMailboxSession) IsExecutorOperatorSetRegistered(operatorSetKey [32]byte) (bool, error) {
	return _TaskMailbox.Contract.IsExecutorOperatorSetRegistered(&_TaskMailbox.CallOpts, operatorSetKey)
}

// IsExecutorOperatorSetRegistered is a free data retrieval call binding the contract method 0xfa2c0b37.
//
// Solidity: function isExecutorOperatorSetRegistered(bytes32 operatorSetKey) view returns(bool isRegistered)
func (_TaskMailbox *TaskMailboxCallerSession) IsExecutorOperatorSetRegistered(operatorSetKey [32]byte) (bool, error) {
	return _TaskMailbox.Contract.IsExecutorOperatorSetRegistered(&_TaskMailbox.CallOpts, operatorSetKey)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TaskMailbox *TaskMailboxCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TaskMailbox.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TaskMailbox *TaskMailboxSession) Owner() (common.Address, error) {
	return _TaskMailbox.Contract.Owner(&_TaskMailbox.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TaskMailbox *TaskMailboxCallerSession) Owner() (common.Address, error) {
	return _TaskMailbox.Contract.Owner(&_TaskMailbox.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_TaskMailbox *TaskMailboxCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TaskMailbox.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_TaskMailbox *TaskMailboxSession) Version() (string, error) {
	return _TaskMailbox.Contract.Version(&_TaskMailbox.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_TaskMailbox *TaskMailboxCallerSession) Version() (string, error) {
	return _TaskMailbox.Contract.Version(&_TaskMailbox.CallOpts)
}

// CreateTask is a paid mutator transaction binding the contract method 0x1fb66f5d.
//
// Solidity: function createTask((address,(address,uint32),bytes) taskParams) returns(bytes32)
func (_TaskMailbox *TaskMailboxTransactor) CreateTask(opts *bind.TransactOpts, taskParams ITaskMailboxTypesTaskParams) (*types.Transaction, error) {
	return _TaskMailbox.contract.Transact(opts, "createTask", taskParams)
}

// CreateTask is a paid mutator transaction binding the contract method 0x1fb66f5d.
//
// Solidity: function createTask((address,(address,uint32),bytes) taskParams) returns(bytes32)
func (_TaskMailbox *TaskMailboxSession) CreateTask(taskParams ITaskMailboxTypesTaskParams) (*types.Transaction, error) {
	return _TaskMailbox.Contract.CreateTask(&_TaskMailbox.TransactOpts, taskParams)
}

// CreateTask is a paid mutator transaction binding the contract method 0x1fb66f5d.
//
// Solidity: function createTask((address,(address,uint32),bytes) taskParams) returns(bytes32)
func (_TaskMailbox *TaskMailboxTransactorSession) CreateTask(taskParams ITaskMailboxTypesTaskParams) (*types.Transaction, error) {
	return _TaskMailbox.Contract.CreateTask(&_TaskMailbox.TransactOpts, taskParams)
}

// Initialize is a paid mutator transaction binding the contract method 0x708c0db9.
//
// Solidity: function initialize(address _owner, uint16 _feeSplit, address _feeSplitCollector) returns()
func (_TaskMailbox *TaskMailboxTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address, _feeSplit uint16, _feeSplitCollector common.Address) (*types.Transaction, error) {
	return _TaskMailbox.contract.Transact(opts, "initialize", _owner, _feeSplit, _feeSplitCollector)
}

// Initialize is a paid mutator transaction binding the contract method 0x708c0db9.
//
// Solidity: function initialize(address _owner, uint16 _feeSplit, address _feeSplitCollector) returns()
func (_TaskMailbox *TaskMailboxSession) Initialize(_owner common.Address, _feeSplit uint16, _feeSplitCollector common.Address) (*types.Transaction, error) {
	return _TaskMailbox.Contract.Initialize(&_TaskMailbox.TransactOpts, _owner, _feeSplit, _feeSplitCollector)
}

// Initialize is a paid mutator transaction binding the contract method 0x708c0db9.
//
// Solidity: function initialize(address _owner, uint16 _feeSplit, address _feeSplitCollector) returns()
func (_TaskMailbox *TaskMailboxTransactorSession) Initialize(_owner common.Address, _feeSplit uint16, _feeSplitCollector common.Address) (*types.Transaction, error) {
	return _TaskMailbox.Contract.Initialize(&_TaskMailbox.TransactOpts, _owner, _feeSplit, _feeSplitCollector)
}

// RefundFee is a paid mutator transaction binding the contract method 0xb8694166.
//
// Solidity: function refundFee(bytes32 taskHash) returns()
func (_TaskMailbox *TaskMailboxTransactor) RefundFee(opts *bind.TransactOpts, taskHash [32]byte) (*types.Transaction, error) {
	return _TaskMailbox.contract.Transact(opts, "refundFee", taskHash)
}

// RefundFee is a paid mutator transaction binding the contract method 0xb8694166.
//
// Solidity: function refundFee(bytes32 taskHash) returns()
func (_TaskMailbox *TaskMailboxSession) RefundFee(taskHash [32]byte) (*types.Transaction, error) {
	return _TaskMailbox.Contract.RefundFee(&_TaskMailbox.TransactOpts, taskHash)
}

// RefundFee is a paid mutator transaction binding the contract method 0xb8694166.
//
// Solidity: function refundFee(bytes32 taskHash) returns()
func (_TaskMailbox *TaskMailboxTransactorSession) RefundFee(taskHash [32]byte) (*types.Transaction, error) {
	return _TaskMailbox.Contract.RefundFee(&_TaskMailbox.TransactOpts, taskHash)
}

// RegisterExecutorOperatorSet is a paid mutator transaction binding the contract method 0x49acd884.
//
// Solidity: function registerExecutorOperatorSet((address,uint32) operatorSet, bool isRegistered) returns()
func (_TaskMailbox *TaskMailboxTransactor) RegisterExecutorOperatorSet(opts *bind.TransactOpts, operatorSet OperatorSet, isRegistered bool) (*types.Transaction, error) {
	return _TaskMailbox.contract.Transact(opts, "registerExecutorOperatorSet", operatorSet, isRegistered)
}

// RegisterExecutorOperatorSet is a paid mutator transaction binding the contract method 0x49acd884.
//
// Solidity: function registerExecutorOperatorSet((address,uint32) operatorSet, bool isRegistered) returns()
func (_TaskMailbox *TaskMailboxSession) RegisterExecutorOperatorSet(operatorSet OperatorSet, isRegistered bool) (*types.Transaction, error) {
	return _TaskMailbox.Contract.RegisterExecutorOperatorSet(&_TaskMailbox.TransactOpts, operatorSet, isRegistered)
}

// RegisterExecutorOperatorSet is a paid mutator transaction binding the contract method 0x49acd884.
//
// Solidity: function registerExecutorOperatorSet((address,uint32) operatorSet, bool isRegistered) returns()
func (_TaskMailbox *TaskMailboxTransactorSession) RegisterExecutorOperatorSet(operatorSet OperatorSet, isRegistered bool) (*types.Transaction, error) {
	return _TaskMailbox.Contract.RegisterExecutorOperatorSet(&_TaskMailbox.TransactOpts, operatorSet, isRegistered)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TaskMailbox *TaskMailboxTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaskMailbox.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TaskMailbox *TaskMailboxSession) RenounceOwnership() (*types.Transaction, error) {
	return _TaskMailbox.Contract.RenounceOwnership(&_TaskMailbox.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TaskMailbox *TaskMailboxTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TaskMailbox.Contract.RenounceOwnership(&_TaskMailbox.TransactOpts)
}

// SetExecutorOperatorSetTaskConfig is a paid mutator transaction binding the contract method 0xf741e81a.
//
// Solidity: function setExecutorOperatorSetTaskConfig((address,uint32) operatorSet, (address,uint96,address,uint8,address,(uint8,bytes),bytes) config) returns()
func (_TaskMailbox *TaskMailboxTransactor) SetExecutorOperatorSetTaskConfig(opts *bind.TransactOpts, operatorSet OperatorSet, config ITaskMailboxTypesExecutorOperatorSetTaskConfig) (*types.Transaction, error) {
	return _TaskMailbox.contract.Transact(opts, "setExecutorOperatorSetTaskConfig", operatorSet, config)
}

// SetExecutorOperatorSetTaskConfig is a paid mutator transaction binding the contract method 0xf741e81a.
//
// Solidity: function setExecutorOperatorSetTaskConfig((address,uint32) operatorSet, (address,uint96,address,uint8,address,(uint8,bytes),bytes) config) returns()
func (_TaskMailbox *TaskMailboxSession) SetExecutorOperatorSetTaskConfig(operatorSet OperatorSet, config ITaskMailboxTypesExecutorOperatorSetTaskConfig) (*types.Transaction, error) {
	return _TaskMailbox.Contract.SetExecutorOperatorSetTaskConfig(&_TaskMailbox.TransactOpts, operatorSet, config)
}

// SetExecutorOperatorSetTaskConfig is a paid mutator transaction binding the contract method 0xf741e81a.
//
// Solidity: function setExecutorOperatorSetTaskConfig((address,uint32) operatorSet, (address,uint96,address,uint8,address,(uint8,bytes),bytes) config) returns()
func (_TaskMailbox *TaskMailboxTransactorSession) SetExecutorOperatorSetTaskConfig(operatorSet OperatorSet, config ITaskMailboxTypesExecutorOperatorSetTaskConfig) (*types.Transaction, error) {
	return _TaskMailbox.Contract.SetExecutorOperatorSetTaskConfig(&_TaskMailbox.TransactOpts, operatorSet, config)
}

// SetFeeSplit is a paid mutator transaction binding the contract method 0x468c07a0.
//
// Solidity: function setFeeSplit(uint16 _feeSplit) returns()
func (_TaskMailbox *TaskMailboxTransactor) SetFeeSplit(opts *bind.TransactOpts, _feeSplit uint16) (*types.Transaction, error) {
	return _TaskMailbox.contract.Transact(opts, "setFeeSplit", _feeSplit)
}

// SetFeeSplit is a paid mutator transaction binding the contract method 0x468c07a0.
//
// Solidity: function setFeeSplit(uint16 _feeSplit) returns()
func (_TaskMailbox *TaskMailboxSession) SetFeeSplit(_feeSplit uint16) (*types.Transaction, error) {
	return _TaskMailbox.Contract.SetFeeSplit(&_TaskMailbox.TransactOpts, _feeSplit)
}

// SetFeeSplit is a paid mutator transaction binding the contract method 0x468c07a0.
//
// Solidity: function setFeeSplit(uint16 _feeSplit) returns()
func (_TaskMailbox *TaskMailboxTransactorSession) SetFeeSplit(_feeSplit uint16) (*types.Transaction, error) {
	return _TaskMailbox.Contract.SetFeeSplit(&_TaskMailbox.TransactOpts, _feeSplit)
}

// SetFeeSplitCollector is a paid mutator transaction binding the contract method 0x678fbdb3.
//
// Solidity: function setFeeSplitCollector(address _feeSplitCollector) returns()
func (_TaskMailbox *TaskMailboxTransactor) SetFeeSplitCollector(opts *bind.TransactOpts, _feeSplitCollector common.Address) (*types.Transaction, error) {
	return _TaskMailbox.contract.Transact(opts, "setFeeSplitCollector", _feeSplitCollector)
}

// SetFeeSplitCollector is a paid mutator transaction binding the contract method 0x678fbdb3.
//
// Solidity: function setFeeSplitCollector(address _feeSplitCollector) returns()
func (_TaskMailbox *TaskMailboxSession) SetFeeSplitCollector(_feeSplitCollector common.Address) (*types.Transaction, error) {
	return _TaskMailbox.Contract.SetFeeSplitCollector(&_TaskMailbox.TransactOpts, _feeSplitCollector)
}

// SetFeeSplitCollector is a paid mutator transaction binding the contract method 0x678fbdb3.
//
// Solidity: function setFeeSplitCollector(address _feeSplitCollector) returns()
func (_TaskMailbox *TaskMailboxTransactorSession) SetFeeSplitCollector(_feeSplitCollector common.Address) (*types.Transaction, error) {
	return _TaskMailbox.Contract.SetFeeSplitCollector(&_TaskMailbox.TransactOpts, _feeSplitCollector)
}

// SubmitResult is a paid mutator transaction binding the contract method 0xa5fabc81.
//
// Solidity: function submitResult(bytes32 taskHash, bytes executorCert, bytes result) returns()
func (_TaskMailbox *TaskMailboxTransactor) SubmitResult(opts *bind.TransactOpts, taskHash [32]byte, executorCert []byte, result []byte) (*types.Transaction, error) {
	return _TaskMailbox.contract.Transact(opts, "submitResult", taskHash, executorCert, result)
}

// SubmitResult is a paid mutator transaction binding the contract method 0xa5fabc81.
//
// Solidity: function submitResult(bytes32 taskHash, bytes executorCert, bytes result) returns()
func (_TaskMailbox *TaskMailboxSession) SubmitResult(taskHash [32]byte, executorCert []byte, result []byte) (*types.Transaction, error) {
	return _TaskMailbox.Contract.SubmitResult(&_TaskMailbox.TransactOpts, taskHash, executorCert, result)
}

// SubmitResult is a paid mutator transaction binding the contract method 0xa5fabc81.
//
// Solidity: function submitResult(bytes32 taskHash, bytes executorCert, bytes result) returns()
func (_TaskMailbox *TaskMailboxTransactorSession) SubmitResult(taskHash [32]byte, executorCert []byte, result []byte) (*types.Transaction, error) {
	return _TaskMailbox.Contract.SubmitResult(&_TaskMailbox.TransactOpts, taskHash, executorCert, result)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TaskMailbox *TaskMailboxTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TaskMailbox.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TaskMailbox *TaskMailboxSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TaskMailbox.Contract.TransferOwnership(&_TaskMailbox.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TaskMailbox *TaskMailboxTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TaskMailbox.Contract.TransferOwnership(&_TaskMailbox.TransactOpts, newOwner)
}

// TaskMailboxExecutorOperatorSetRegisteredIterator is returned from FilterExecutorOperatorSetRegistered and is used to iterate over the raw logs and unpacked data for ExecutorOperatorSetRegistered events raised by the TaskMailbox contract.
type TaskMailboxExecutorOperatorSetRegisteredIterator struct {
	Event *TaskMailboxExecutorOperatorSetRegistered // Event containing the contract specifics and raw log

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
func (it *TaskMailboxExecutorOperatorSetRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskMailboxExecutorOperatorSetRegistered)
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
		it.Event = new(TaskMailboxExecutorOperatorSetRegistered)
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
func (it *TaskMailboxExecutorOperatorSetRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskMailboxExecutorOperatorSetRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskMailboxExecutorOperatorSetRegistered represents a ExecutorOperatorSetRegistered event raised by the TaskMailbox contract.
type TaskMailboxExecutorOperatorSetRegistered struct {
	Caller                common.Address
	Avs                   common.Address
	ExecutorOperatorSetId uint32
	IsRegistered          bool
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterExecutorOperatorSetRegistered is a free log retrieval operation binding the contract event 0x48b63f21a1eb9dd6880e196de6d7db3fbd0c282b74f1298dcb4cf53472298f39.
//
// Solidity: event ExecutorOperatorSetRegistered(address indexed caller, address indexed avs, uint32 indexed executorOperatorSetId, bool isRegistered)
func (_TaskMailbox *TaskMailboxFilterer) FilterExecutorOperatorSetRegistered(opts *bind.FilterOpts, caller []common.Address, avs []common.Address, executorOperatorSetId []uint32) (*TaskMailboxExecutorOperatorSetRegisteredIterator, error) {

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

	logs, sub, err := _TaskMailbox.contract.FilterLogs(opts, "ExecutorOperatorSetRegistered", callerRule, avsRule, executorOperatorSetIdRule)
	if err != nil {
		return nil, err
	}
	return &TaskMailboxExecutorOperatorSetRegisteredIterator{contract: _TaskMailbox.contract, event: "ExecutorOperatorSetRegistered", logs: logs, sub: sub}, nil
}

// WatchExecutorOperatorSetRegistered is a free log subscription operation binding the contract event 0x48b63f21a1eb9dd6880e196de6d7db3fbd0c282b74f1298dcb4cf53472298f39.
//
// Solidity: event ExecutorOperatorSetRegistered(address indexed caller, address indexed avs, uint32 indexed executorOperatorSetId, bool isRegistered)
func (_TaskMailbox *TaskMailboxFilterer) WatchExecutorOperatorSetRegistered(opts *bind.WatchOpts, sink chan<- *TaskMailboxExecutorOperatorSetRegistered, caller []common.Address, avs []common.Address, executorOperatorSetId []uint32) (event.Subscription, error) {

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

	logs, sub, err := _TaskMailbox.contract.WatchLogs(opts, "ExecutorOperatorSetRegistered", callerRule, avsRule, executorOperatorSetIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskMailboxExecutorOperatorSetRegistered)
				if err := _TaskMailbox.contract.UnpackLog(event, "ExecutorOperatorSetRegistered", log); err != nil {
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
func (_TaskMailbox *TaskMailboxFilterer) ParseExecutorOperatorSetRegistered(log types.Log) (*TaskMailboxExecutorOperatorSetRegistered, error) {
	event := new(TaskMailboxExecutorOperatorSetRegistered)
	if err := _TaskMailbox.contract.UnpackLog(event, "ExecutorOperatorSetRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskMailboxExecutorOperatorSetTaskConfigSetIterator is returned from FilterExecutorOperatorSetTaskConfigSet and is used to iterate over the raw logs and unpacked data for ExecutorOperatorSetTaskConfigSet events raised by the TaskMailbox contract.
type TaskMailboxExecutorOperatorSetTaskConfigSetIterator struct {
	Event *TaskMailboxExecutorOperatorSetTaskConfigSet // Event containing the contract specifics and raw log

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
func (it *TaskMailboxExecutorOperatorSetTaskConfigSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskMailboxExecutorOperatorSetTaskConfigSet)
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
		it.Event = new(TaskMailboxExecutorOperatorSetTaskConfigSet)
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
func (it *TaskMailboxExecutorOperatorSetTaskConfigSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskMailboxExecutorOperatorSetTaskConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskMailboxExecutorOperatorSetTaskConfigSet represents a ExecutorOperatorSetTaskConfigSet event raised by the TaskMailbox contract.
type TaskMailboxExecutorOperatorSetTaskConfigSet struct {
	Caller                common.Address
	Avs                   common.Address
	ExecutorOperatorSetId uint32
	Config                ITaskMailboxTypesExecutorOperatorSetTaskConfig
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterExecutorOperatorSetTaskConfigSet is a free log retrieval operation binding the contract event 0x7cd76abd4025a20959a1b20f7c1536e3894a0735cd8de0215dde803ddea7f2d2.
//
// Solidity: event ExecutorOperatorSetTaskConfigSet(address indexed caller, address indexed avs, uint32 indexed executorOperatorSetId, (address,uint96,address,uint8,address,(uint8,bytes),bytes) config)
func (_TaskMailbox *TaskMailboxFilterer) FilterExecutorOperatorSetTaskConfigSet(opts *bind.FilterOpts, caller []common.Address, avs []common.Address, executorOperatorSetId []uint32) (*TaskMailboxExecutorOperatorSetTaskConfigSetIterator, error) {

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

	logs, sub, err := _TaskMailbox.contract.FilterLogs(opts, "ExecutorOperatorSetTaskConfigSet", callerRule, avsRule, executorOperatorSetIdRule)
	if err != nil {
		return nil, err
	}
	return &TaskMailboxExecutorOperatorSetTaskConfigSetIterator{contract: _TaskMailbox.contract, event: "ExecutorOperatorSetTaskConfigSet", logs: logs, sub: sub}, nil
}

// WatchExecutorOperatorSetTaskConfigSet is a free log subscription operation binding the contract event 0x7cd76abd4025a20959a1b20f7c1536e3894a0735cd8de0215dde803ddea7f2d2.
//
// Solidity: event ExecutorOperatorSetTaskConfigSet(address indexed caller, address indexed avs, uint32 indexed executorOperatorSetId, (address,uint96,address,uint8,address,(uint8,bytes),bytes) config)
func (_TaskMailbox *TaskMailboxFilterer) WatchExecutorOperatorSetTaskConfigSet(opts *bind.WatchOpts, sink chan<- *TaskMailboxExecutorOperatorSetTaskConfigSet, caller []common.Address, avs []common.Address, executorOperatorSetId []uint32) (event.Subscription, error) {

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

	logs, sub, err := _TaskMailbox.contract.WatchLogs(opts, "ExecutorOperatorSetTaskConfigSet", callerRule, avsRule, executorOperatorSetIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskMailboxExecutorOperatorSetTaskConfigSet)
				if err := _TaskMailbox.contract.UnpackLog(event, "ExecutorOperatorSetTaskConfigSet", log); err != nil {
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
func (_TaskMailbox *TaskMailboxFilterer) ParseExecutorOperatorSetTaskConfigSet(log types.Log) (*TaskMailboxExecutorOperatorSetTaskConfigSet, error) {
	event := new(TaskMailboxExecutorOperatorSetTaskConfigSet)
	if err := _TaskMailbox.contract.UnpackLog(event, "ExecutorOperatorSetTaskConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskMailboxFeeRefundedIterator is returned from FilterFeeRefunded and is used to iterate over the raw logs and unpacked data for FeeRefunded events raised by the TaskMailbox contract.
type TaskMailboxFeeRefundedIterator struct {
	Event *TaskMailboxFeeRefunded // Event containing the contract specifics and raw log

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
func (it *TaskMailboxFeeRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskMailboxFeeRefunded)
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
		it.Event = new(TaskMailboxFeeRefunded)
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
func (it *TaskMailboxFeeRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskMailboxFeeRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskMailboxFeeRefunded represents a FeeRefunded event raised by the TaskMailbox contract.
type TaskMailboxFeeRefunded struct {
	RefundCollector common.Address
	TaskHash        [32]byte
	AvsFee          *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterFeeRefunded is a free log retrieval operation binding the contract event 0xe3ed40d31808582f7a92a30beacc0ec788d5091407ec6c10c1b999b3f317aea3.
//
// Solidity: event FeeRefunded(address indexed refundCollector, bytes32 indexed taskHash, uint96 avsFee)
func (_TaskMailbox *TaskMailboxFilterer) FilterFeeRefunded(opts *bind.FilterOpts, refundCollector []common.Address, taskHash [][32]byte) (*TaskMailboxFeeRefundedIterator, error) {

	var refundCollectorRule []interface{}
	for _, refundCollectorItem := range refundCollector {
		refundCollectorRule = append(refundCollectorRule, refundCollectorItem)
	}
	var taskHashRule []interface{}
	for _, taskHashItem := range taskHash {
		taskHashRule = append(taskHashRule, taskHashItem)
	}

	logs, sub, err := _TaskMailbox.contract.FilterLogs(opts, "FeeRefunded", refundCollectorRule, taskHashRule)
	if err != nil {
		return nil, err
	}
	return &TaskMailboxFeeRefundedIterator{contract: _TaskMailbox.contract, event: "FeeRefunded", logs: logs, sub: sub}, nil
}

// WatchFeeRefunded is a free log subscription operation binding the contract event 0xe3ed40d31808582f7a92a30beacc0ec788d5091407ec6c10c1b999b3f317aea3.
//
// Solidity: event FeeRefunded(address indexed refundCollector, bytes32 indexed taskHash, uint96 avsFee)
func (_TaskMailbox *TaskMailboxFilterer) WatchFeeRefunded(opts *bind.WatchOpts, sink chan<- *TaskMailboxFeeRefunded, refundCollector []common.Address, taskHash [][32]byte) (event.Subscription, error) {

	var refundCollectorRule []interface{}
	for _, refundCollectorItem := range refundCollector {
		refundCollectorRule = append(refundCollectorRule, refundCollectorItem)
	}
	var taskHashRule []interface{}
	for _, taskHashItem := range taskHash {
		taskHashRule = append(taskHashRule, taskHashItem)
	}

	logs, sub, err := _TaskMailbox.contract.WatchLogs(opts, "FeeRefunded", refundCollectorRule, taskHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskMailboxFeeRefunded)
				if err := _TaskMailbox.contract.UnpackLog(event, "FeeRefunded", log); err != nil {
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
func (_TaskMailbox *TaskMailboxFilterer) ParseFeeRefunded(log types.Log) (*TaskMailboxFeeRefunded, error) {
	event := new(TaskMailboxFeeRefunded)
	if err := _TaskMailbox.contract.UnpackLog(event, "FeeRefunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskMailboxFeeSplitCollectorSetIterator is returned from FilterFeeSplitCollectorSet and is used to iterate over the raw logs and unpacked data for FeeSplitCollectorSet events raised by the TaskMailbox contract.
type TaskMailboxFeeSplitCollectorSetIterator struct {
	Event *TaskMailboxFeeSplitCollectorSet // Event containing the contract specifics and raw log

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
func (it *TaskMailboxFeeSplitCollectorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskMailboxFeeSplitCollectorSet)
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
		it.Event = new(TaskMailboxFeeSplitCollectorSet)
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
func (it *TaskMailboxFeeSplitCollectorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskMailboxFeeSplitCollectorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskMailboxFeeSplitCollectorSet represents a FeeSplitCollectorSet event raised by the TaskMailbox contract.
type TaskMailboxFeeSplitCollectorSet struct {
	FeeSplitCollector common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterFeeSplitCollectorSet is a free log retrieval operation binding the contract event 0x262aa27c244f6f0088cb3092548a0adcaddedf459070a9ccab2dc6a07abe701d.
//
// Solidity: event FeeSplitCollectorSet(address indexed feeSplitCollector)
func (_TaskMailbox *TaskMailboxFilterer) FilterFeeSplitCollectorSet(opts *bind.FilterOpts, feeSplitCollector []common.Address) (*TaskMailboxFeeSplitCollectorSetIterator, error) {

	var feeSplitCollectorRule []interface{}
	for _, feeSplitCollectorItem := range feeSplitCollector {
		feeSplitCollectorRule = append(feeSplitCollectorRule, feeSplitCollectorItem)
	}

	logs, sub, err := _TaskMailbox.contract.FilterLogs(opts, "FeeSplitCollectorSet", feeSplitCollectorRule)
	if err != nil {
		return nil, err
	}
	return &TaskMailboxFeeSplitCollectorSetIterator{contract: _TaskMailbox.contract, event: "FeeSplitCollectorSet", logs: logs, sub: sub}, nil
}

// WatchFeeSplitCollectorSet is a free log subscription operation binding the contract event 0x262aa27c244f6f0088cb3092548a0adcaddedf459070a9ccab2dc6a07abe701d.
//
// Solidity: event FeeSplitCollectorSet(address indexed feeSplitCollector)
func (_TaskMailbox *TaskMailboxFilterer) WatchFeeSplitCollectorSet(opts *bind.WatchOpts, sink chan<- *TaskMailboxFeeSplitCollectorSet, feeSplitCollector []common.Address) (event.Subscription, error) {

	var feeSplitCollectorRule []interface{}
	for _, feeSplitCollectorItem := range feeSplitCollector {
		feeSplitCollectorRule = append(feeSplitCollectorRule, feeSplitCollectorItem)
	}

	logs, sub, err := _TaskMailbox.contract.WatchLogs(opts, "FeeSplitCollectorSet", feeSplitCollectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskMailboxFeeSplitCollectorSet)
				if err := _TaskMailbox.contract.UnpackLog(event, "FeeSplitCollectorSet", log); err != nil {
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
func (_TaskMailbox *TaskMailboxFilterer) ParseFeeSplitCollectorSet(log types.Log) (*TaskMailboxFeeSplitCollectorSet, error) {
	event := new(TaskMailboxFeeSplitCollectorSet)
	if err := _TaskMailbox.contract.UnpackLog(event, "FeeSplitCollectorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskMailboxFeeSplitSetIterator is returned from FilterFeeSplitSet and is used to iterate over the raw logs and unpacked data for FeeSplitSet events raised by the TaskMailbox contract.
type TaskMailboxFeeSplitSetIterator struct {
	Event *TaskMailboxFeeSplitSet // Event containing the contract specifics and raw log

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
func (it *TaskMailboxFeeSplitSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskMailboxFeeSplitSet)
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
		it.Event = new(TaskMailboxFeeSplitSet)
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
func (it *TaskMailboxFeeSplitSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskMailboxFeeSplitSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskMailboxFeeSplitSet represents a FeeSplitSet event raised by the TaskMailbox contract.
type TaskMailboxFeeSplitSet struct {
	FeeSplit uint16
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFeeSplitSet is a free log retrieval operation binding the contract event 0x886b2cfcb151fd8b19ed902cc88f4a06dd9fe351a4a9ab93f33fe84abc157edf.
//
// Solidity: event FeeSplitSet(uint16 feeSplit)
func (_TaskMailbox *TaskMailboxFilterer) FilterFeeSplitSet(opts *bind.FilterOpts) (*TaskMailboxFeeSplitSetIterator, error) {

	logs, sub, err := _TaskMailbox.contract.FilterLogs(opts, "FeeSplitSet")
	if err != nil {
		return nil, err
	}
	return &TaskMailboxFeeSplitSetIterator{contract: _TaskMailbox.contract, event: "FeeSplitSet", logs: logs, sub: sub}, nil
}

// WatchFeeSplitSet is a free log subscription operation binding the contract event 0x886b2cfcb151fd8b19ed902cc88f4a06dd9fe351a4a9ab93f33fe84abc157edf.
//
// Solidity: event FeeSplitSet(uint16 feeSplit)
func (_TaskMailbox *TaskMailboxFilterer) WatchFeeSplitSet(opts *bind.WatchOpts, sink chan<- *TaskMailboxFeeSplitSet) (event.Subscription, error) {

	logs, sub, err := _TaskMailbox.contract.WatchLogs(opts, "FeeSplitSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskMailboxFeeSplitSet)
				if err := _TaskMailbox.contract.UnpackLog(event, "FeeSplitSet", log); err != nil {
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
func (_TaskMailbox *TaskMailboxFilterer) ParseFeeSplitSet(log types.Log) (*TaskMailboxFeeSplitSet, error) {
	event := new(TaskMailboxFeeSplitSet)
	if err := _TaskMailbox.contract.UnpackLog(event, "FeeSplitSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskMailboxInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the TaskMailbox contract.
type TaskMailboxInitializedIterator struct {
	Event *TaskMailboxInitialized // Event containing the contract specifics and raw log

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
func (it *TaskMailboxInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskMailboxInitialized)
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
		it.Event = new(TaskMailboxInitialized)
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
func (it *TaskMailboxInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskMailboxInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskMailboxInitialized represents a Initialized event raised by the TaskMailbox contract.
type TaskMailboxInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TaskMailbox *TaskMailboxFilterer) FilterInitialized(opts *bind.FilterOpts) (*TaskMailboxInitializedIterator, error) {

	logs, sub, err := _TaskMailbox.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TaskMailboxInitializedIterator{contract: _TaskMailbox.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TaskMailbox *TaskMailboxFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TaskMailboxInitialized) (event.Subscription, error) {

	logs, sub, err := _TaskMailbox.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskMailboxInitialized)
				if err := _TaskMailbox.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_TaskMailbox *TaskMailboxFilterer) ParseInitialized(log types.Log) (*TaskMailboxInitialized, error) {
	event := new(TaskMailboxInitialized)
	if err := _TaskMailbox.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskMailboxOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TaskMailbox contract.
type TaskMailboxOwnershipTransferredIterator struct {
	Event *TaskMailboxOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TaskMailboxOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskMailboxOwnershipTransferred)
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
		it.Event = new(TaskMailboxOwnershipTransferred)
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
func (it *TaskMailboxOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskMailboxOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskMailboxOwnershipTransferred represents a OwnershipTransferred event raised by the TaskMailbox contract.
type TaskMailboxOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TaskMailbox *TaskMailboxFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TaskMailboxOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TaskMailbox.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TaskMailboxOwnershipTransferredIterator{contract: _TaskMailbox.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TaskMailbox *TaskMailboxFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TaskMailboxOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TaskMailbox.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskMailboxOwnershipTransferred)
				if err := _TaskMailbox.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TaskMailbox *TaskMailboxFilterer) ParseOwnershipTransferred(log types.Log) (*TaskMailboxOwnershipTransferred, error) {
	event := new(TaskMailboxOwnershipTransferred)
	if err := _TaskMailbox.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskMailboxTaskCreatedIterator is returned from FilterTaskCreated and is used to iterate over the raw logs and unpacked data for TaskCreated events raised by the TaskMailbox contract.
type TaskMailboxTaskCreatedIterator struct {
	Event *TaskMailboxTaskCreated // Event containing the contract specifics and raw log

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
func (it *TaskMailboxTaskCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskMailboxTaskCreated)
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
		it.Event = new(TaskMailboxTaskCreated)
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
func (it *TaskMailboxTaskCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskMailboxTaskCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskMailboxTaskCreated represents a TaskCreated event raised by the TaskMailbox contract.
type TaskMailboxTaskCreated struct {
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
func (_TaskMailbox *TaskMailboxFilterer) FilterTaskCreated(opts *bind.FilterOpts, creator []common.Address, taskHash [][32]byte, avs []common.Address) (*TaskMailboxTaskCreatedIterator, error) {

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

	logs, sub, err := _TaskMailbox.contract.FilterLogs(opts, "TaskCreated", creatorRule, taskHashRule, avsRule)
	if err != nil {
		return nil, err
	}
	return &TaskMailboxTaskCreatedIterator{contract: _TaskMailbox.contract, event: "TaskCreated", logs: logs, sub: sub}, nil
}

// WatchTaskCreated is a free log subscription operation binding the contract event 0x33add0b01e02278be5459fbfa3274aee699ec47f4ee7236b59e7a2c8b5000c26.
//
// Solidity: event TaskCreated(address indexed creator, bytes32 indexed taskHash, address indexed avs, uint32 executorOperatorSetId, uint32 operatorTableReferenceTimestamp, address refundCollector, uint96 avsFee, uint256 taskDeadline, bytes payload)
func (_TaskMailbox *TaskMailboxFilterer) WatchTaskCreated(opts *bind.WatchOpts, sink chan<- *TaskMailboxTaskCreated, creator []common.Address, taskHash [][32]byte, avs []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _TaskMailbox.contract.WatchLogs(opts, "TaskCreated", creatorRule, taskHashRule, avsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskMailboxTaskCreated)
				if err := _TaskMailbox.contract.UnpackLog(event, "TaskCreated", log); err != nil {
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
func (_TaskMailbox *TaskMailboxFilterer) ParseTaskCreated(log types.Log) (*TaskMailboxTaskCreated, error) {
	event := new(TaskMailboxTaskCreated)
	if err := _TaskMailbox.contract.UnpackLog(event, "TaskCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskMailboxTaskVerifiedIterator is returned from FilterTaskVerified and is used to iterate over the raw logs and unpacked data for TaskVerified events raised by the TaskMailbox contract.
type TaskMailboxTaskVerifiedIterator struct {
	Event *TaskMailboxTaskVerified // Event containing the contract specifics and raw log

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
func (it *TaskMailboxTaskVerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskMailboxTaskVerified)
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
		it.Event = new(TaskMailboxTaskVerified)
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
func (it *TaskMailboxTaskVerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskMailboxTaskVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskMailboxTaskVerified represents a TaskVerified event raised by the TaskMailbox contract.
type TaskMailboxTaskVerified struct {
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
func (_TaskMailbox *TaskMailboxFilterer) FilterTaskVerified(opts *bind.FilterOpts, aggregator []common.Address, taskHash [][32]byte, avs []common.Address) (*TaskMailboxTaskVerifiedIterator, error) {

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

	logs, sub, err := _TaskMailbox.contract.FilterLogs(opts, "TaskVerified", aggregatorRule, taskHashRule, avsRule)
	if err != nil {
		return nil, err
	}
	return &TaskMailboxTaskVerifiedIterator{contract: _TaskMailbox.contract, event: "TaskVerified", logs: logs, sub: sub}, nil
}

// WatchTaskVerified is a free log subscription operation binding the contract event 0x659f23b2e7edf490e5fd6561c5148691ed0375ed7ddd3ab1bcfcfdbec4f209a9.
//
// Solidity: event TaskVerified(address indexed aggregator, bytes32 indexed taskHash, address indexed avs, uint32 executorOperatorSetId, bytes executorCert, bytes result)
func (_TaskMailbox *TaskMailboxFilterer) WatchTaskVerified(opts *bind.WatchOpts, sink chan<- *TaskMailboxTaskVerified, aggregator []common.Address, taskHash [][32]byte, avs []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _TaskMailbox.contract.WatchLogs(opts, "TaskVerified", aggregatorRule, taskHashRule, avsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskMailboxTaskVerified)
				if err := _TaskMailbox.contract.UnpackLog(event, "TaskVerified", log); err != nil {
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
func (_TaskMailbox *TaskMailboxFilterer) ParseTaskVerified(log types.Log) (*TaskMailboxTaskVerified, error) {
	event := new(TaskMailboxTaskVerified)
	if err := _TaskMailbox.contract.UnpackLog(event, "TaskVerified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
