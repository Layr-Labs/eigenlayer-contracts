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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_bn254CertificateVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_ecdsaCertificateVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_maxTaskSLA\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"_version\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"BN254_CERTIFICATE_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ECDSA_CERTIFICATE_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_TASK_SLA\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createTask\",\"inputs\":[{\"name\":\"taskParams\",\"type\":\"tuple\",\"internalType\":\"structITaskMailboxTypes.TaskParams\",\"components\":[{\"name\":\"refundCollector\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"executorOperatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"feeSplit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"feeSplitCollector\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBN254CertificateBytes\",\"inputs\":[{\"name\":\"cert\",\"type\":\"tuple\",\"internalType\":\"structIBN254CertificateVerifierTypes.BN254Certificate\",\"components\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"nonSignerWitnesses\",\"type\":\"tuple[]\",\"internalType\":\"structIBN254CertificateVerifierTypes.BN254OperatorInfoWitness[]\",\"components\":[{\"name\":\"operatorIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorInfoProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"operatorInfo\",\"type\":\"tuple\",\"internalType\":\"structIOperatorTableCalculatorTypes.BN254OperatorInfo\",\"components\":[{\"name\":\"pubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}]}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getECDSACertificateBytes\",\"inputs\":[{\"name\":\"cert\",\"type\":\"tuple\",\"internalType\":\"structIECDSACertificateVerifierTypes.ECDSACertificate\",\"components\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sig\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getExecutorOperatorSetTaskConfig\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structITaskMailboxTypes.ExecutorOperatorSetTaskConfig\",\"components\":[{\"name\":\"taskHook\",\"type\":\"address\",\"internalType\":\"contractIAVSTaskHook\"},{\"name\":\"taskSLA\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"curveType\",\"type\":\"uint8\",\"internalType\":\"enumIKeyRegistrarTypes.CurveType\"},{\"name\":\"feeCollector\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"consensus\",\"type\":\"tuple\",\"internalType\":\"structITaskMailboxTypes.Consensus\",\"components\":[{\"name\":\"consensusType\",\"type\":\"uint8\",\"internalType\":\"enumITaskMailboxTypes.ConsensusType\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"taskMetadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskInfo\",\"inputs\":[{\"name\":\"taskHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structITaskMailboxTypes.Task\",\"components\":[{\"name\":\"creator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"creationTime\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avsFee\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"refundCollector\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"executorOperatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeSplit\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumITaskMailboxTypes.TaskStatus\"},{\"name\":\"isFeeRefunded\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"operatorTableReferenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"executorOperatorSetTaskConfig\",\"type\":\"tuple\",\"internalType\":\"structITaskMailboxTypes.ExecutorOperatorSetTaskConfig\",\"components\":[{\"name\":\"taskHook\",\"type\":\"address\",\"internalType\":\"contractIAVSTaskHook\"},{\"name\":\"taskSLA\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"curveType\",\"type\":\"uint8\",\"internalType\":\"enumIKeyRegistrarTypes.CurveType\"},{\"name\":\"feeCollector\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"consensus\",\"type\":\"tuple\",\"internalType\":\"structITaskMailboxTypes.Consensus\",\"components\":[{\"name\":\"consensusType\",\"type\":\"uint8\",\"internalType\":\"enumITaskMailboxTypes.ConsensusType\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"taskMetadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"executorCert\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"result\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskResult\",\"inputs\":[{\"name\":\"taskHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskStatus\",\"inputs\":[{\"name\":\"taskHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumITaskMailboxTypes.TaskStatus\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_feeSplit\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"_feeSplitCollector\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isExecutorOperatorSetRegistered\",\"inputs\":[{\"name\":\"operatorSetKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"isRegistered\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"refundFee\",\"inputs\":[{\"name\":\"taskHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerExecutorOperatorSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"isRegistered\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setExecutorOperatorSetTaskConfig\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structITaskMailboxTypes.ExecutorOperatorSetTaskConfig\",\"components\":[{\"name\":\"taskHook\",\"type\":\"address\",\"internalType\":\"contractIAVSTaskHook\"},{\"name\":\"taskSLA\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"curveType\",\"type\":\"uint8\",\"internalType\":\"enumIKeyRegistrarTypes.CurveType\"},{\"name\":\"feeCollector\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"consensus\",\"type\":\"tuple\",\"internalType\":\"structITaskMailboxTypes.Consensus\",\"components\":[{\"name\":\"consensusType\",\"type\":\"uint8\",\"internalType\":\"enumITaskMailboxTypes.ConsensusType\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"taskMetadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFeeSplit\",\"inputs\":[{\"name\":\"_feeSplit\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFeeSplitCollector\",\"inputs\":[{\"name\":\"_feeSplitCollector\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitResult\",\"inputs\":[{\"name\":\"taskHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"executorCert\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"result\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"ExecutorOperatorSetRegistered\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"executorOperatorSetId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"isRegistered\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExecutorOperatorSetTaskConfigSet\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"executorOperatorSetId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structITaskMailboxTypes.ExecutorOperatorSetTaskConfig\",\"components\":[{\"name\":\"taskHook\",\"type\":\"address\",\"internalType\":\"contractIAVSTaskHook\"},{\"name\":\"taskSLA\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"curveType\",\"type\":\"uint8\",\"internalType\":\"enumIKeyRegistrarTypes.CurveType\"},{\"name\":\"feeCollector\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"consensus\",\"type\":\"tuple\",\"internalType\":\"structITaskMailboxTypes.Consensus\",\"components\":[{\"name\":\"consensusType\",\"type\":\"uint8\",\"internalType\":\"enumITaskMailboxTypes.ConsensusType\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"taskMetadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeRefunded\",\"inputs\":[{\"name\":\"refundCollector\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"taskHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"avsFee\",\"type\":\"uint96\",\"indexed\":false,\"internalType\":\"uint96\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeSplitCollectorSet\",\"inputs\":[{\"name\":\"feeSplitCollector\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeSplitSet\",\"inputs\":[{\"name\":\"feeSplit\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskCreated\",\"inputs\":[{\"name\":\"creator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"taskHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"executorOperatorSetId\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"operatorTableReferenceTimestamp\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"refundCollector\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"avsFee\",\"type\":\"uint96\",\"indexed\":false,\"internalType\":\"uint96\"},{\"name\":\"taskDeadline\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskVerified\",\"inputs\":[{\"name\":\"aggregator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"taskHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"executorOperatorSetId\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"executorCert\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"result\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CertificateStale\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EmptyCertificateSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExecutorOperatorSetNotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExecutorOperatorSetTaskConfigNotSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FeeAlreadyRefunded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidConsensusType\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidConsensusValue\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCurveType\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidFeeReceiver\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidFeeSplit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidMessageHash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperatorSetOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidReferenceTimestamp\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTaskCreator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTaskStatus\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"enumITaskMailboxTypes.TaskStatus\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"enumITaskMailboxTypes.TaskStatus\"}]},{\"type\":\"error\",\"name\":\"OnlyRefundCollector\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PayloadIsEmpty\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"TaskSLAExceedsMaximum\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ThresholdNotMet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TimestampAtCreation\",\"inputs\":[]}]",
	Bin: "0x610100604052348015610010575f5ffd5b50604051615c6e380380615c6e83398101604081905261002f9161019b565b6001600160a01b03808516608052831660a0526001600160601b03821660c052806100598161006e565b60e052506100656100b4565b505050506102e6565b5f5f829050601f815111156100a1578260405163305a27a960e01b8152600401610098919061028b565b60405180910390fd5b80516100ac826102c0565b179392505050565b5f54610100900460ff161561011b5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b6064820152608401610098565b5f5460ff9081161461016a575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b80516001600160a01b0381168114610182575f5ffd5b919050565b634e487b7160e01b5f52604160045260245ffd5b5f5f5f5f608085870312156101ae575f5ffd5b6101b78561016c565b93506101c56020860161016c565b60408601519093506001600160601b03811681146101e1575f5ffd5b60608601519092506001600160401b038111156101fc575f5ffd5b8501601f8101871361020c575f5ffd5b80516001600160401b0381111561022557610225610187565b604051601f8201601f19908116603f011681016001600160401b038111828210171561025357610253610187565b60405281815282820160200189101561026a575f5ffd5b8160208401602083015e5f6020838301015280935050505092959194509250565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b805160208083015191908110156102e0575f198160200360031b1b821691505b50919050565b60805160a05160c05160e05161592461034a5f395f611a8301525f818161033a0152612ebe01525f818161025201528181613233015281816138450152613a6801525f818161039f015281816131f301528181613773015261399501526159245ff3fe608060405234801561000f575f5ffd5b5060043610610153575f3560e01c8063678fbdb3116100bf578063b869416611610079578063b869416614610322578063d3e043aa14610335578063f2fde38b14610374578063f741e81a14610387578063f7424fc91461039a578063fa2c0b37146103c1575f5ffd5b8063678fbdb3146102b05780636bf6fad5146102c3578063708c0db9146102e3578063715018a6146102f65780638da5cb5b146102fe578063a5fabc811461030f575f5ffd5b806349acd8841161011057806349acd8841461021a5780634ad52e021461022d57806354743ad21461024d57806354fd4d501461027457806362fee0371461027c5780636373ea691461028f575f5ffd5b80631270a892146101575780631a20c505146101805780631ae370eb146101b15780631fb66f5d146101c45780632bf6cc79146101e5578063468c07a014610205575b5f5ffd5b61016a610165366004614215565b6103f3565b60405161017791906142ca565b60405180910390f35b609b54610199906201000090046001600160a01b031681565b6040516001600160a01b039091168152602001610177565b61016a6101bf366004614527565b61041c565b6101d76101d2366004614666565b61042f565b604051908152602001610177565b6101f86101f33660046146db565b610dd8565b604051610177919061471a565b610218610213366004614737565b611284565b005b61021861022836600461475f565b611298565b61024061023b3660046146db565b6114e3565b604051610177919061487b565b6101997f000000000000000000000000000000000000000000000000000000000000000081565b61016a611a7c565b61016a61028a3660046146db565b611aac565b609b5461029d9061ffff1681565b60405161ffff9091168152602001610177565b6102186102be3660046149c6565b611fa9565b6102d66102d13660046149e1565b611fba565b60405161017791906149fb565b6102186102f1366004614a0d565b6121ce565b6102186122ff565b6033546001600160a01b0316610199565b61021861031d366004614a55565b612312565b6102186103303660046146db565b612bee565b61035c7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160601b039091168152602001610177565b6102186103823660046149c6565b612e20565b610218610395366004614b47565b612e96565b6101997f000000000000000000000000000000000000000000000000000000000000000081565b6103e36103cf3660046146db565b60996020525f908152604090205460ff1681565b6040519015158152602001610177565b6060816040516020016104069190614c72565b6040516020818303038152906040529050919050565b6060816040516020016104069190614df6565b5f6104386130cb565b5f8260400151511161045d57604051636b1a1b6960e11b815260040160405180910390fd5b60995f61046d8460200151613124565b815260208101919091526040015f205460ff1661049d5760405163c292b29760e01b815260040160405180910390fd5b5f609a5f6104ae8560200151613124565b815260208082019290925260409081015f20815160e08101835281546001600160a01b038082168352600160a01b918290046001600160601b03169583019590955260018301549485169382019390935292909160608401910460ff16600281111561051c5761051c6146f2565b600281111561052d5761052d6146f2565b815260028201546001600160a01b0316602082015260408051808201825260038401805492909301929091829060ff16600181111561056e5761056e6146f2565b600181111561057f5761057f6146f2565b815260200160018201805461059390614e08565b80601f01602080910402602001604051908101604052809291908181526020018280546105bf90614e08565b801561060a5780601f106105e15761010080835404028352916020019161060a565b820191905f5260205f20905b8154815290600101906020018083116105ed57829003601f168201915b505050505081525050815260200160058201805461062790614e08565b80601f016020809104026020016040519081016040528092919081815260200182805461065390614e08565b801561069e5780601f106106755761010080835404028352916020019161069e565b820191905f5260205f20905b81548152906001019060200180831161068157829003601f168201915b50505050508152505090506106b28161318d565b6106cf576040516314b0a41d60e11b815260040160405180910390fd5b5f6106dd82606001516131d6565b90505f816001600160a01b0316635ddb9b5b86602001516040518263ffffffff1660e01b81526004016107109190614e5e565b602060405180830381865afa15801561072b573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061074f9190614e6c565b90505f826001600160a01b0316636141879e87602001516040518263ffffffff1660e01b81526004016107829190614e5e565b602060405180830381865afa15801561079d573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107c19190614e6c565b905063ffffffff811615806107fd57506107db8183614e9b565b63ffffffff1684602001516001600160601b0316426107fa9190614eb7565b11155b61081a5760405163640fcd6b60e11b815260040160405180910390fd5b508251604051630a3fc61360e31b81526001600160a01b03909116906351fe30989061084c9033908990600401614f05565b5f6040518083038186803b158015610862575f5ffd5b505afa158015610874573d5f5f3e3d5ffd5b50508451604051637036693f60e11b81525f93506001600160a01b03909116915063e06cd27e906108a9908990600401614f28565b602060405180830381865afa1580156108c4573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108e89190614f3a565b90505f6097543046896040516020016109049493929190614f55565b604051602081830303815290604052805190602001209050609754600161092b9190614eb7565b609755604080516101c081019091523381526020810161094a42613270565b6001600160601b03908116825260208a81018051516001600160a01b039081168386015292871660408501528b5190921660608401529051015163ffffffff166080820152609b5461ffff1660a082015260c001600181525f602080830182905263ffffffff80881660408086019190915260608086018c90528d8201516080808801919091528251808601845286815260a0808901919091528351808701855287815260c09889015289875260988652958390208851958901516001600160601b03908116600160a01b9081026001600160a01b03988916178355948a0151938a0151168402928616929092176001830155870151600282018054968901519789015161ffff16600160c01b0261ffff60c01b19989095169093026001600160c01b0319909616941693909317939093179384168117835560e0850151919391929160ff60d01b1990911662ffffff60c01b1990911617600160d01b836003811115610ab957610ab96146f2565b02179055506101008201516002808301805461012086015163ffffffff16600160e01b026001600160e01b03941515600160d81b02949094166001600160d81b039091161792909217909155610140830151805160208201516001600160601b0316600160a01b9081026001600160a01b0392831617600386019081556040840151600487018054919094166001600160a01b031982168117855560608601519596929594936001600160a81b031990921617918490811115610b7e57610b7e6146f2565b021790555060808201516002820180546001600160a01b0319166001600160a01b0390921691909117905560a08201518051600383018054909190829060ff191660018381811115610bd257610bd26146f2565b021790555060208201516001820190610beb9082614fd6565b50505060c08201516005820190610c029082614fd6565b5050506101608201516009820190610c1a9082614fd6565b50610180820151600a820190610c309082614fd6565b506101a0820151600b820190610c469082614fd6565b50505060408501516001600160a01b031615801590610c6d57505f826001600160601b0316115b15610ce85760808501516001600160a01b0316610c9d57604051633480121760e21b815260040160405180910390fd5b86516001600160a01b0316610cc557604051633480121760e21b815260040160405180910390fd5b6040850151610ce8906001600160a01b031633306001600160601b0386166132db565b8451604051629c5c4560e41b8152600481018390526001600160a01b03909116906309c5c450906024015f604051808303815f87803b158015610d29575f5ffd5b505af1158015610d3b573d5f5f3e3d5ffd5b5050505086602001515f01516001600160a01b031681336001600160a01b03167f33add0b01e02278be5459fbfa3274aee699ec47f4ee7236b59e7a2c8b5000c268a6020015160200151878c5f0151888c602001516001600160601b031642610da49190614eb7565b8f60400151604051610dbb96959493929190615090565b60405180910390a4945050505050610dd36001606555565b919050565b5f81815260986020908152604080832081516101c08101835281546001600160a01b038082168352600160a01b918290046001600160601b03908116968401969096526001840154808216958401959095529381900490941660608201526002820154928316608082015292820463ffffffff1660a0840152600160c01b820461ffff1660c084015283929160e0830190600160d01b900460ff166003811115610e8457610e846146f2565b6003811115610e9557610e956146f2565b8152600282810154600160d81b810460ff9081161515602080860191909152600160e01b90920463ffffffff16604080860191909152805160e0810182526003870180546001600160a01b038082168452600160a01b918290046001600160601b03169684019690965260048901549586169383019390935260609687019691959094918601939290910490911690811115610f3357610f336146f2565b6002811115610f4457610f446146f2565b815260028201546001600160a01b0316602082015260408051808201825260038401805492909301929091829060ff166001811115610f8557610f856146f2565b6001811115610f9657610f966146f2565b8152602001600182018054610faa90614e08565b80601f0160208091040260200160405190810160405280929190818152602001828054610fd690614e08565b80156110215780601f10610ff857610100808354040283529160200191611021565b820191905f5260205f20905b81548152906001019060200180831161100457829003601f168201915b505050505081525050815260200160058201805461103e90614e08565b80601f016020809104026020016040519081016040528092919081815260200182805461106a90614e08565b80156110b55780601f1061108c576101008083540402835291602001916110b5565b820191905f5260205f20905b81548152906001019060200180831161109857829003601f168201915b50505050508152505081526020016009820180546110d290614e08565b80601f01602080910402602001604051908101604052809291908181526020018280546110fe90614e08565b80156111495780601f1061112057610100808354040283529160200191611149565b820191905f5260205f20905b81548152906001019060200180831161112c57829003601f168201915b50505050508152602001600a8201805461116290614e08565b80601f016020809104026020016040519081016040528092919081815260200182805461118e90614e08565b80156111d95780601f106111b0576101008083540402835291602001916111d9565b820191905f5260205f20905b8154815290600101906020018083116111bc57829003601f168201915b50505050508152602001600b820180546111f290614e08565b80601f016020809104026020016040519081016040528092919081815260200182805461121e90614e08565b80156112695780601f1061124057610100808354040283529160200191611269565b820191905f5260205f20905b81548152906001019060200180831161124c57829003601f168201915b505050505081525050905061127d8161334d565b9392505050565b61128c6133a8565b61129581613402565b50565b5f609a5f6112a585613124565b815260208082019290925260409081015f20815160e08101835281546001600160a01b038082168352600160a01b918290046001600160601b03169583019590955260018301549485169382019390935292909160608401910460ff166002811115611313576113136146f2565b6002811115611324576113246146f2565b815260028201546001600160a01b0316602082015260408051808201825260038401805492909301929091829060ff166001811115611365576113656146f2565b6001811115611376576113766146f2565b815260200160018201805461138a90614e08565b80601f01602080910402602001604051908101604052809291908181526020018280546113b690614e08565b80156114015780601f106113d857610100808354040283529160200191611401565b820191905f5260205f20905b8154815290600101906020018083116113e457829003601f168201915b505050505081525050815260200160058201805461141e90614e08565b80601f016020809104026020016040519081016040528092919081815260200182805461144a90614e08565b80156114955780601f1061146c57610100808354040283529160200191611495565b820191905f5260205f20905b81548152906001019060200180831161147857829003601f168201915b50505050508152505090506114a98161318d565b6114c6576040516314b0a41d60e11b815260040160405180910390fd5b6114d4838260600151613473565b6114de8383613514565b505050565b6114eb613ff5565b5f82815260986020908152604080832081516101c08101835281546001600160a01b038082168352600160a01b918290046001600160601b03908116968401969096526001840154808216958401959095529381900490941660608201526002820154928316608082015292820463ffffffff1660a0840152600160c01b820461ffff1660c08401529060e0830190600160d01b900460ff166003811115611595576115956146f2565b60038111156115a6576115a66146f2565b8152600282810154600160d81b810460ff9081161515602080860191909152600160e01b90920463ffffffff16604080860191909152805160e0810182526003870180546001600160a01b038082168452600160a01b918290046001600160601b03169684019690965260048901549586169383019390935260609687019691959094918601939290910490911690811115611644576116446146f2565b6002811115611655576116556146f2565b815260028201546001600160a01b0316602082015260408051808201825260038401805492909301929091829060ff166001811115611696576116966146f2565b60018111156116a7576116a76146f2565b81526020016001820180546116bb90614e08565b80601f01602080910402602001604051908101604052809291908181526020018280546116e790614e08565b80156117325780601f1061170957610100808354040283529160200191611732565b820191905f5260205f20905b81548152906001019060200180831161171557829003601f168201915b505050505081525050815260200160058201805461174f90614e08565b80601f016020809104026020016040519081016040528092919081815260200182805461177b90614e08565b80156117c65780601f1061179d576101008083540402835291602001916117c6565b820191905f5260205f20905b8154815290600101906020018083116117a957829003601f168201915b50505050508152505081526020016009820180546117e390614e08565b80601f016020809104026020016040519081016040528092919081815260200182805461180f90614e08565b801561185a5780601f106118315761010080835404028352916020019161185a565b820191905f5260205f20905b81548152906001019060200180831161183d57829003601f168201915b50505050508152602001600a8201805461187390614e08565b80601f016020809104026020016040519081016040528092919081815260200182805461189f90614e08565b80156118ea5780601f106118c1576101008083540402835291602001916118ea565b820191905f5260205f20905b8154815290600101906020018083116118cd57829003601f168201915b50505050508152602001600b8201805461190390614e08565b80601f016020809104026020016040519081016040528092919081815260200182805461192f90614e08565b801561197a5780601f106119515761010080835404028352916020019161197a565b820191905f5260205f20905b81548152906001019060200180831161195d57829003601f168201915b5050505050815250509050604051806101c00160405280825f01516001600160a01b0316815260200182602001516001600160601b0316815260200182604001516001600160a01b0316815260200182606001516001600160601b0316815260200182608001516001600160a01b031681526020018260a0015163ffffffff1681526020018260c0015161ffff168152602001611a168361334d565b6003811115611a2757611a276146f2565b81526020018261010001511515815260200182610120015163ffffffff168152602001826101400151815260200182610160015181526020018261018001518152602001826101a00151815250915050919050565b6060611aa77f00000000000000000000000000000000000000000000000000000000000000006135a5565b905090565b5f81815260986020908152604080832081516101c08101835281546001600160a01b038082168352600160a01b918290046001600160601b0390811696840196909652600184015480821695840195909552938190049094166060808301919091526002830154938416608083015293830463ffffffff1660a0820152600160c01b830461ffff1660c08201529293929160e0830190600160d01b900460ff166003811115611b5d57611b5d6146f2565b6003811115611b6e57611b6e6146f2565b8152600282810154600160d81b810460ff9081161515602080860191909152600160e01b90920463ffffffff16604080860191909152805160e0810182526003870180546001600160a01b038082168452600160a01b918290046001600160601b03169684019690965260048901549586169383019390935260609687019691959094918601939290910490911690811115611c0c57611c0c6146f2565b6002811115611c1d57611c1d6146f2565b815260028201546001600160a01b0316602082015260408051808201825260038401805492909301929091829060ff166001811115611c5e57611c5e6146f2565b6001811115611c6f57611c6f6146f2565b8152602001600182018054611c8390614e08565b80601f0160208091040260200160405190810160405280929190818152602001828054611caf90614e08565b8015611cfa5780601f10611cd157610100808354040283529160200191611cfa565b820191905f5260205f20905b815481529060010190602001808311611cdd57829003601f168201915b5050505050815250508152602001600582018054611d1790614e08565b80601f0160208091040260200160405190810160405280929190818152602001828054611d4390614e08565b8015611d8e5780601f10611d6557610100808354040283529160200191611d8e565b820191905f5260205f20905b815481529060010190602001808311611d7157829003601f168201915b5050505050815250508152602001600982018054611dab90614e08565b80601f0160208091040260200160405190810160405280929190818152602001828054611dd790614e08565b8015611e225780601f10611df957610100808354040283529160200191611e22565b820191905f5260205f20905b815481529060010190602001808311611e0557829003601f168201915b50505050508152602001600a82018054611e3b90614e08565b80601f0160208091040260200160405190810160405280929190818152602001828054611e6790614e08565b8015611eb25780601f10611e8957610100808354040283529160200191611eb2565b820191905f5260205f20905b815481529060010190602001808311611e9557829003601f168201915b50505050508152602001600b82018054611ecb90614e08565b80601f0160208091040260200160405190810160405280929190818152602001828054611ef790614e08565b8015611f425780601f10611f1957610100808354040283529160200191611f42565b820191905f5260205f20905b815481529060010190602001808311611f2557829003601f168201915b50505050508152505090505f611f578261334d565b90506002816003811115611f6d57611f6d6146f2565b146002829091611f9b57604051634091b18960e11b8152600401611f929291906150e9565b60405180910390fd5b5050506101a0015192915050565b611fb16133a8565b611295816135e2565b611fc2614068565b609a5f611fce84613124565b815260208082019290925260409081015f20815160e08101835281546001600160a01b038082168352600160a01b918290046001600160601b03169583019590955260018301549485169382019390935292909160608401910460ff16600281111561203c5761203c6146f2565b600281111561204d5761204d6146f2565b815260028201546001600160a01b0316602082015260408051808201825260038401805492909301929091829060ff16600181111561208e5761208e6146f2565b600181111561209f5761209f6146f2565b81526020016001820180546120b390614e08565b80601f01602080910402602001604051908101604052809291908181526020018280546120df90614e08565b801561212a5780601f106121015761010080835404028352916020019161212a565b820191905f5260205f20905b81548152906001019060200180831161210d57829003601f168201915b505050505081525050815260200160058201805461214790614e08565b80601f016020809104026020016040519081016040528092919081815260200182805461217390614e08565b80156121be5780601f10612195576101008083540402835291602001916121be565b820191905f5260205f20905b8154815290600101906020018083116121a157829003601f168201915b5050505050815250509050919050565b5f54610100900460ff16158080156121ec57505f54600160ff909116105b806122055750303b15801561220557505f5460ff166001145b6122685760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401611f92565b5f805460ff191660011790558015612289575f805461ff0019166101001790555b61229161365c565b61229961368a565b6122a2846136b8565b6122ab83613402565b6122b4826135e2565b80156122f9575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050565b6123076133a8565b6123105f6136b8565b565b61231a6130cb565b5f83815260986020908152604080832081516101c08101835281546001600160a01b038082168352600160a01b918290046001600160601b03908116968401969096526001840154808216958401959095529381900490941660608201526002820154928316608082015292820463ffffffff1660a0840152600160c01b820461ffff1660c084015292916127be91849060e0830190600160d01b900460ff1660038111156123cb576123cb6146f2565b60038111156123dc576123dc6146f2565b8152600282810154600160d81b810460ff9081161515602080860191909152600160e01b90920463ffffffff16604080860191909152805160e0810182526003870180546001600160a01b038082168452600160a01b918290046001600160601b0316968401969096526004890154958616938301939093526060968701969195909491860193929091049091169081111561247a5761247a6146f2565b600281111561248b5761248b6146f2565b815260028201546001600160a01b0316602082015260408051808201825260038401805492909301929091829060ff1660018111156124cc576124cc6146f2565b60018111156124dd576124dd6146f2565b81526020016001820180546124f190614e08565b80601f016020809104026020016040519081016040528092919081815260200182805461251d90614e08565b80156125685780601f1061253f57610100808354040283529160200191612568565b820191905f5260205f20905b81548152906001019060200180831161254b57829003601f168201915b505050505081525050815260200160058201805461258590614e08565b80601f01602080910402602001604051908101604052809291908181526020018280546125b190614e08565b80156125fc5780601f106125d3576101008083540402835291602001916125fc565b820191905f5260205f20905b8154815290600101906020018083116125df57829003601f168201915b505050505081525050815260200160098201805461261990614e08565b80601f016020809104026020016040519081016040528092919081815260200182805461264590614e08565b80156126905780601f1061266757610100808354040283529160200191612690565b820191905f5260205f20905b81548152906001019060200180831161267357829003601f168201915b50505050508152602001600a820180546126a990614e08565b80601f01602080910402602001604051908101604052809291908181526020018280546126d590614e08565b80156127205780601f106126f757610100808354040283529160200191612720565b820191905f5260205f20905b81548152906001019060200180831161270357829003601f168201915b50505050508152602001600b8201805461273990614e08565b80601f016020809104026020016040519081016040528092919081815260200182805461276590614e08565b80156127b05780601f10612787576101008083540402835291602001916127b0565b820191905f5260205f20905b81548152906001019060200180831161279357829003601f168201915b50505050508152505061334d565b905060018160038111156127d4576127d46146f2565b1460018290916127f957604051634091b18960e11b8152600401611f929291906150e9565b50508154600160a01b90046001600160601b0316421161282c5760405163015a4b7560e51b815260040160405180910390fd5b600382015460405163ba33565d60e01b81526001600160a01b039091169063ba33565d90612864903390899089908990600401615104565b5f6040518083038186803b15801561287a575f5ffd5b505afa15801561288c573d5f5f3e3d5ffd5b50506040805180820182526001808701546001600160a01b03168252600287015463ffffffff600160a01b91829004166020840152600488015484518086019095526006890180549497506129c6965060ff92909104821694939092849291909116908111156128fe576128fe6146f2565b600181111561290f5761290f6146f2565b815260200160018201805461292390614e08565b80601f016020809104026020016040519081016040528092919081815260200182805461294f90614e08565b801561299a5780601f106129715761010080835404028352916020019161299a565b820191905f5260205f20905b81548152906001019060200180831161297d57829003601f168201915b505050919092525050506002860154875160208901208591600160e01b900463ffffffff16908a613709565b60028301805460ff60d01b1916600160d11b179055600a83016129e98682614fd6565b50600b83016129f88582614fd6565b5060048301546001600160a01b031615801590612a2857506001830154600160a01b90046001600160601b031615155b15612b1757600283015460018401545f91612a779161271091612a6891600160c01b90910461ffff16906001600160601b03600160a01b9091041661513c565b612a729190615153565b613270565b90506001600160601b03811615612ab757609b546004850154612ab7916001600160a01b039182169162010000909104166001600160601b038416613b2b565b60018401545f90612ad9908390600160a01b90046001600160601b0316615172565b90506001600160601b03811615612b145760058501546004860154612b14916001600160a01b0391821691166001600160601b038416613b2b565b50505b6003830154604051637041233f60e11b8152336004820152602481018890526001600160a01b039091169063e082467e906044015f604051808303815f87803b158015612b62575f5ffd5b505af1158015612b74573d5f5f3e3d5ffd5b505050600184015460028501546040516001600160a01b039092169250889133917f659f23b2e7edf490e5fd6561c5148691ed0375ed7ddd3ab1bcfcfdbec4f209a991612bd99163ffffffff600160a01b9091041690600a8a0190600b8b0190615210565b60405180910390a45050506114de6001606555565b612bf66130cb565b5f81815260986020526040902060028101546001600160a01b03163314612c30576040516370f43cb760e01b815260040160405180910390fd5b6002810154600160d81b900460ff1615612c5d57604051633e3d786960e01b815260040160405180910390fd5b604080516101c08101825282546001600160a01b038082168352600160a01b918290046001600160601b03908116602085015260018601548083169585019590955293829004909316606083015260028401549283166080830152820463ffffffff1660a0820152600160c01b820461ffff1660c08201525f91612cfe9190849060e0830190600160d01b900460ff1660038111156123cb576123cb6146f2565b90506003816003811115612d1457612d146146f2565b146003829091612d3957604051634091b18960e11b8152600401611f929291906150e9565b505060028201805460ff60d81b1916600160d81b17905560048201546001600160a01b031615801590612d7f57506001820154600160a01b90046001600160601b031615155b15612dbb57600282015460018301546004840154612dbb926001600160a01b0391821692911690600160a01b90046001600160601b0316613b2b565b60028201546001830154604051600160a01b9091046001600160601b0316815284916001600160a01b0316907fe3ed40d31808582f7a92a30beacc0ec788d5091407ec6c10c1b999b3f317aea39060200160405180910390a350506112956001606555565b612e286133a8565b6001600160a01b038116612e8d5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401611f92565b611295816136b8565b612e9f8161318d565b612ebc576040516314b0a41d60e11b815260040160405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160601b031681602001516001600160601b03161115612f135760405163a7cf966560e01b815260040160405180910390fd5b612f208160a00151613b5b565b612f2e828260600151613473565b80609a5f612f3b85613124565b815260208082019290925260409081015f208351928401516001600160601b0316600160a01b9081026001600160a01b0394851617825591840151600182018054919094166001600160a01b03198216811785556060860151929492936001600160a81b03199092161790836002811115612fb857612fb86146f2565b021790555060808201516002820180546001600160a01b0319166001600160a01b0390921691909117905560a08201518051600383018054909190829060ff19166001838181111561300c5761300c6146f2565b0217905550602082015160018201906130259082614fd6565b50505060c0820151600582019061303c9082614fd6565b50905050816020015163ffffffff16825f01516001600160a01b0316336001600160a01b03167f7cd76abd4025a20959a1b20f7c1536e3894a0735cd8de0215dde803ddea7f2d28460405161309191906149fb565b60405180910390a460995f6130a584613124565b815260208101919091526040015f205460ff166130c7576130c7826001613514565b5050565b60026065540361311d5760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401611f92565b6002606555565b5f815f0151826020015163ffffffff1660405160200161316f92919060609290921b6bffffffffffffffffffffffff1916825260a01b6001600160a01b031916601482015260200190565b60405160208183030381529060405261318790615240565b92915050565b5f80826060015160028111156131a5576131a56146f2565b141580156131bc575081516001600160a01b031615155b8015613187575050602001516001600160601b0316151590565b5f60028260028111156131eb576131eb6146f2565b0361321757507f0000000000000000000000000000000000000000000000000000000000000000919050565b600182600281111561322b5761322b6146f2565b0361325757507f0000000000000000000000000000000000000000000000000000000000000000919050565b60405163fdea7c0960e01b815260040160405180910390fd5b5f6001600160601b038211156132d75760405162461bcd60e51b815260206004820152602660248201527f53616665436173743a2076616c756520646f65736e27742066697420696e203960448201526536206269747360d01b6064820152608401611f92565b5090565b6040516001600160a01b03808516602483015283166044820152606481018290526122f99085906323b872dd60e01b906084015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b031990931692909217909152613c1a565b6001606555565b5f60018260e001516003811115613366576133666146f2565b14801561339357508161014001516020015182602001516133879190615263565b6001600160601b031642115b156133a057506003919050565b5060e0015190565b6033546001600160a01b031633146123105760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401611f92565b61271061ffff8216111561342957604051630601f69760e01b815260040160405180910390fd5b609b805461ffff191661ffff83169081179091556040519081527f886b2cfcb151fd8b19ed902cc88f4a06dd9fe351a4a9ab93f33fe84abc157edf9060200160405180910390a150565b5f61347d826131d6565b6040516304240c4960e51b815290915033906001600160a01b038316906384818920906134ae908790600401614e5e565b602060405180830381865afa1580156134c9573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906134ed9190615282565b6001600160a01b0316146114de576040516342ecfee960e11b815260040160405180910390fd5b8060995f61352185613124565b81526020019081526020015f205f6101000a81548160ff021916908315150217905550816020015163ffffffff16825f01516001600160a01b0316336001600160a01b03167f48b63f21a1eb9dd6880e196de6d7db3fbd0c282b74f1298dcb4cf53472298f3984604051613599911515815260200190565b60405180910390a45050565b60605f6135b183613ced565b6040805160208082528183019092529192505f91906020820181803683375050509182525060208101929092525090565b6001600160a01b03811661360957604051630863a45360e11b815260040160405180910390fd5b609b805462010000600160b01b031916620100006001600160a01b038416908102919091179091556040517f262aa27c244f6f0088cb3092548a0adcaddedf459070a9ccab2dc6a07abe701d905f90a250565b5f54610100900460ff166136825760405162461bcd60e51b8152600401611f929061529d565b612310613d14565b5f54610100900460ff166136b05760405162461bcd60e51b8152600401611f929061529d565b612310613d43565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b5f8551600181111561371d5761371d6146f2565b036138c5576002866002811115613736576137366146f2565b036137f3575f818060200190518101906137509190615548565b905061375d818585613d69565b6040516280b71560e41b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063080b7150906137aa908890859060040161561a565b5f604051808303815f875af11580156137c5573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526137ec9190810190615639565b5050613b23565b6001866002811115613807576138076146f2565b03613257575f81806020019051810190613821919061566a565b905061382e818585613df3565b6040516380c7d3f360e01b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906380c7d3f39061387c90889085906004016156e3565b5f60405180830381865afa158015613896573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526138bd9190810190615766565b505050613b23565b6001855160018111156138da576138da6146f2565b03613b0a575f85602001518060200190518101906138f891906157c9565b6040805160018082528183019092529192505f91906020808301908036833701905050905081815f81518110613930576139306157e4565b61ffff909216602092830291909101909101525f6002896002811115613958576139586146f2565b03613a16575f848060200190518101906139729190615548565b905061397f818888613d69565b604051625f5e5d60e21b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063017d7974906139ce908b9085908890600401615836565b6020604051808303815f875af11580156139ea573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613a0e9190615867565b915050613aec565b6001896002811115613a2a57613a2a6146f2565b03613257575f84806020019051810190613a44919061566a565b9050613a51818888613df3565b604051630606d12160e51b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063c0da242090613aa1908b9085908890600401615882565b5f60405180830381865afa158015613abb573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052613ae291908101906158a1565b509150613aec9050565b806138bd576040516359fa4a9360e01b815260040160405180910390fd5b6040516347d3772160e11b815260040160405180910390fd5b505050505050565b6040516001600160a01b0383166024820152604481018290526114de90849063a9059cbb60e01b9060640161330f565b5f81516001811115613b6f57613b6f6146f2565b03613b98576020810151511561129557604051631501e04760e21b815260040160405180910390fd5b600181516001811115613bad57613bad6146f2565b03613b0a57806020015151602014613bd857604051631501e04760e21b815260040160405180910390fd5b5f8160200151806020019051810190613bf191906157c9565b905061271061ffff821611156130c757604051631501e04760e21b815260040160405180910390fd5b5f613c6e826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316613e6b9092919063ffffffff16565b905080515f1480613c8e575080806020019051810190613c8e9190615867565b6114de5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608401611f92565b5f60ff8216601f81111561318757604051632cd44ac360e21b815260040160405180910390fd5b5f54610100900460ff16613d3a5760405162461bcd60e51b8152600401611f929061529d565b612310336136b8565b5f54610100900460ff166133465760405162461bcd60e51b8152600401611f929061529d565b8163ffffffff16835f015163ffffffff1614613d9857604051634534032960e01b815260040160405180910390fd5b80836020015114613dbc57604051638b56642d60e01b815260040160405180910390fd5b604083015151158015613dd55750604083015160200151155b156114de57604051637a8a1dbd60e11b815260040160405180910390fd5b8163ffffffff16835f015163ffffffff1614613e2257604051634534032960e01b815260040160405180910390fd5b80836020015114613e4657604051638b56642d60e01b815260040160405180910390fd5b5f836040015151116114de57604051637a8a1dbd60e11b815260040160405180910390fd5b6060613e7984845f85613e81565b949350505050565b606082471015613ee25760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b6064820152608401611f92565b5f5f866001600160a01b03168587604051613efd91906158d8565b5f6040518083038185875af1925050503d805f8114613f37576040519150601f19603f3d011682016040523d82523d5f602084013e613f3c565b606091505b5091509150613f4d87838387613f58565b979650505050505050565b60608315613fc65782515f03613fbf576001600160a01b0385163b613fbf5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401611f92565b5081613e79565b613e798383815115613fdb5781518083602001fd5b8060405162461bcd60e51b8152600401611f9291906142ca565b604080516101c0810182525f80825260208201819052918101829052606081018290526080810182905260a0810182905260c081018290529060e082019081525f60208201819052604082015260600161404d614068565b81526020016060815260200160608152602001606081525090565b6040805160e0810182525f8082526020820181905291810182905290606082019081525f602082015260400161409c6140a9565b8152602001606081525090565b60408051808201909152805f61409c565b634e487b7160e01b5f52604160045260245ffd5b604051606081016001600160401b03811182821017156140f0576140f06140ba565b60405290565b604080519081016001600160401b03811182821017156140f0576140f06140ba565b60405160a081016001600160401b03811182821017156140f0576140f06140ba565b60405160e081016001600160401b03811182821017156140f0576140f06140ba565b604051601f8201601f191681016001600160401b0381118282101715614184576141846140ba565b604052919050565b63ffffffff81168114611295575f5ffd5b5f6001600160401b038211156141b5576141b56140ba565b50601f01601f191660200190565b5f82601f8301126141d2575f5ffd5b81356141e56141e08261419d565b61415c565b8181528460208386010111156141f9575f5ffd5b816020850160208301375f918101602001919091529392505050565b5f60208284031215614225575f5ffd5b81356001600160401b0381111561423a575f5ffd5b82016060818503121561424b575f5ffd5b6142536140ce565b813561425e8161418c565b81526020828101359082015260408201356001600160401b03811115614282575f5ffd5b61428e868285016141c3565b604083015250949350505050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f61127d602083018461429c565b5f604082840312156142ec575f5ffd5b6142f46140f6565b823581526020928301359281019290925250919050565b5f82601f83011261431a575f5ffd5b614324604061415c565b806040840185811115614335575f5ffd5b845b8181101561434f578035845260209384019301614337565b509095945050505050565b5f6001600160401b03821115614372576143726140ba565b5060051b60200190565b5f82601f83011261438b575f5ffd5b81356143996141e08261435a565b8082825260208201915060208360051b8601019250858311156143ba575f5ffd5b602085015b8381101561451d5780356001600160401b038111156143dc575f5ffd5b86016060818903601f190112156143f1575f5ffd5b6143f96140ce565b60208201356144078161418c565b815260408201356001600160401b03811115614421575f5ffd5b6144308a6020838601016141c3565b60208301525060608201356001600160401b0381111561444e575f5ffd5b6020818401019250506060828a031215614466575f5ffd5b61446e6140f6565b6144788a846142dc565b815260408301356001600160401b03811115614492575f5ffd5b80840193505089601f8401126144a6575f5ffd5b82356144b46141e08261435a565b8082825260208201915060208360051b87010192508c8311156144d5575f5ffd5b6020860195505b828610156144f75785358252602095860195909101906144dc565b8060208501525050508060408301525080855250506020830192506020810190506143bf565b5095945050505050565b5f60208284031215614537575f5ffd5b81356001600160401b0381111561454c575f5ffd5b820180840361012081121561455f575f5ffd5b614567614118565b82356145728161418c565b81526020838101359082015261458b86604085016142dc565b60408201526080607f19830112156145a1575f5ffd5b6145a96140f6565b91506145b8866080850161430b565b82526145c78660c0850161430b565b602083015281606082015261010083013591506001600160401b038211156145ed575f5ffd5b6145f98683850161437c565b608082015295945050505050565b6001600160a01b0381168114611295575f5ffd5b8035610dd381614607565b5f60408284031215614636575f5ffd5b61463e6140f6565b9050813561464b81614607565b8152602082013561465b8161418c565b602082015292915050565b5f60208284031215614676575f5ffd5b81356001600160401b0381111561468b575f5ffd5b82016080818503121561469c575f5ffd5b6146a46140ce565b81356146af81614607565b81526146be8560208401614626565b602082015260608201356001600160401b03811115614282575f5ffd5b5f602082840312156146eb575f5ffd5b5035919050565b634e487b7160e01b5f52602160045260245ffd5b60048110614716576147166146f2565b9052565b602081016131878284614706565b61ffff81168114611295575f5ffd5b5f60208284031215614747575f5ffd5b813561127d81614728565b8015158114611295575f5ffd5b5f5f60608385031215614770575f5ffd5b61477a8484614626565b9150604083013561478a81614752565b809150509250929050565b60038110614716576147166146f2565b5f8151600281106147b8576147b86146f2565b80845250602082015160406020850152613e79604085018261429c565b80516001600160a01b031682526020808201516001600160601b0316908301526040808201515f91614811908501826001600160a01b03169052565b5060608201516148246060850182614795565b50608082015161483f60808501826001600160a01b03169052565b5060a082015160e060a085015261485960e08501826147a5565b905060c083015184820360c0860152614872828261429c565b95945050505050565b602081526148956020820183516001600160a01b03169052565b5f60208301516148b060408401826001600160601b03169052565b5060408301516001600160a01b03811660608401525060608301516001600160601b03811660808401525060808301516001600160a01b03811660a08401525060a083015163ffffffff811660c08401525060c083015161ffff811660e08401525060e0830151614925610100840182614706565b506101008301518015156101208401525061012083015163ffffffff8116610140840152506101408301516101c06101608401526149676101e08401826147d5565b9050610160840151601f1984830301610180850152614986828261429c565b915050610180840151601f19848303016101a08501526149a6828261429c565b9150506101a0840151601f19848303016101c0850152614872828261429c565b5f602082840312156149d6575f5ffd5b813561127d81614607565b5f604082840312156149f1575f5ffd5b61127d8383614626565b602081525f61127d60208301846147d5565b5f5f5f60608486031215614a1f575f5ffd5b8335614a2a81614607565b92506020840135614a3a81614728565b91506040840135614a4a81614607565b809150509250925092565b5f5f5f60608486031215614a67575f5ffd5b8335925060208401356001600160401b03811115614a83575f5ffd5b614a8f868287016141c3565b92505060408401356001600160401b03811115614aaa575f5ffd5b614ab6868287016141c3565b9150509250925092565b6001600160601b0381168114611295575f5ffd5b8035610dd381614ac0565b803560038110610dd3575f5ffd5b5f60408284031215614afd575f5ffd5b614b056140f6565b9050813560028110614b15575f5ffd5b815260208201356001600160401b03811115614b2f575f5ffd5b614b3b848285016141c3565b60208301525092915050565b5f5f60608385031215614b58575f5ffd5b614b628484614626565b915060408301356001600160401b03811115614b7c575f5ffd5b830160e08186031215614b8d575f5ffd5b614b9561413a565b614b9e8261461b565b8152614bac60208301614ad4565b6020820152614bbd6040830161461b565b6040820152614bce60608301614adf565b6060820152614bdf6080830161461b565b608082015260a08201356001600160401b03811115614bfc575f5ffd5b614c0887828501614aed565b60a08301525060c08201356001600160401b03811115614c26575f5ffd5b614c32878285016141c3565b60c08301525080925050509250929050565b63ffffffff8151168252602081015160208301525f604082015160606040850152613e79606085018261429c565b602081525f61127d6020830184614c44565b805f5b60028110156122f9578151845260209384019390910190600101614c87565b5f610120830163ffffffff8351168452602083015160208501526040830151614cdc604086018280518252602090810151910152565b506060830151614cf0608086018251614c84565b60200151614d0160c0860182614c84565b506080830151610120610100860152818151808452610140870191506101408160051b88010193506020830192505f5b81811015614dea5761013f19888603018352835163ffffffff8151168652602081015160606020880152614d68606088018261429c565b905060408201519150868103604088015260608101614d9282845180518252602090810151910152565b6020928301516060604084015280518083529301925f92608001905b80841015614dd15784518252602082019150602085019450600184019350614dae565b5097505050602094850194939093019250600101614d31565b50929695505050505050565b602081525f61127d6020830184614ca6565b600181811c90821680614e1c57607f821691505b602082108103614e3a57634e487b7160e01b5f52602260045260245ffd5b50919050565b80516001600160a01b0316825260209081015163ffffffff16910152565b604081016131878284614e40565b5f60208284031215614e7c575f5ffd5b815161127d8161418c565b634e487b7160e01b5f52601160045260245ffd5b63ffffffff818116838216019081111561318757613187614e87565b8082018082111561318757613187614e87565b80516001600160a01b031682526020808201515f91614eeb90850182614e40565b50604082015160806060850152613e79608085018261429c565b6001600160a01b03831681526040602082018190525f90613e7990830184614eca565b602081525f61127d6020830184614eca565b5f60208284031215614f4a575f5ffd5b815161127d81614ac0565b84815260018060a01b0384166020820152826040820152608060608201525f614f816080830184614eca565b9695505050505050565b601f8211156114de57805f5260205f20601f840160051c81016020851015614fb05750805b601f840160051c820191505b81811015614fcf575f8155600101614fbc565b5050505050565b81516001600160401b03811115614fef57614fef6140ba565b61500381614ffd8454614e08565b84614f8b565b6020601f821160018114615035575f831561501e5750848201515b5f19600385901b1c1916600184901b178455614fcf565b5f84815260208120601f198516915b828110156150645787850151825560209485019460019092019101615044565b508482101561508157868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b63ffffffff8716815263ffffffff8616602082015260018060a01b03851660408201526001600160601b038416606082015282608082015260c060a08201525f6150dd60c083018461429c565b98975050505050505050565b604081016150f78285614706565b61127d6020830184614706565b60018060a01b0385168152836020820152608060408201525f61512a608083018561429c565b8281036060840152613f4d818561429c565b808202811582820484141761318757613187614e87565b5f8261516d57634e487b7160e01b5f52601260045260245ffd5b500490565b6001600160601b03828116828216039081111561318757613187614e87565b5f815461519d81614e08565b8085526001821680156151b757600181146151d357615207565b60ff1983166020870152602082151560051b8701019350615207565b845f5260205f205f5b838110156151fe5781546020828a0101526001820191506020810190506151dc565b87016020019450505b50505092915050565b63ffffffff84168152606060208201525f61522e6060830185615191565b8281036040840152614f818185615191565b80516020808301519190811015614e3a575f1960209190910360031b1b16919050565b6001600160601b03818116838216019081111561318757613187614e87565b5f60208284031215615292575f5ffd5b815161127d81614607565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b5f604082840312156152f8575f5ffd5b6153006140f6565b825181526020928301519281019290925250919050565b5f82601f830112615326575f5ffd5b615330604061415c565b806040840185811115615341575f5ffd5b845b8181101561434f578051845260209384019301615343565b5f82601f83011261536a575f5ffd5b81516153786141e08261419d565b81815284602083860101111561538c575f5ffd5b8160208501602083015e5f918101602001919091529392505050565b5f82601f8301126153b7575f5ffd5b81516153c56141e08261435a565b8082825260208201915060208360051b8601019250858311156153e6575f5ffd5b602085015b8381101561451d5780518352602092830192016153eb565b5f82601f830112615412575f5ffd5b81516154206141e08261435a565b8082825260208201915060208360051b860101925085831115615441575f5ffd5b602085015b8381101561451d5780516001600160401b03811115615463575f5ffd5b86016060818903601f19011215615478575f5ffd5b6154806140ce565b602082015161548e8161418c565b815260408201516001600160401b038111156154a8575f5ffd5b6154b78a60208386010161535b565b60208301525060608201516001600160401b038111156154d5575f5ffd5b6020818401019250506060828a0312156154ed575f5ffd5b6154f56140f6565b6154ff8a846152e8565b815260408301516001600160401b03811115615519575f5ffd5b6155258b8286016153a8565b602083015250806040830152508085525050602083019250602081019050615446565b5f60208284031215615558575f5ffd5b81516001600160401b0381111561556d575f5ffd5b8201808403610120811215615580575f5ffd5b615588614118565b82516155938161418c565b8152602083810151908201526155ac86604085016152e8565b60408201526080607f19830112156155c2575f5ffd5b6155ca6140f6565b91506155d98660808501615317565b82526155e88660c08501615317565b602083015281606082015261010083015191506001600160401b0382111561560e575f5ffd5b6145f986838501615403565b6156248184614e40565b606060408201525f613e796060830184614ca6565b5f60208284031215615649575f5ffd5b81516001600160401b0381111561565e575f5ffd5b613e79848285016153a8565b5f6020828403121561567a575f5ffd5b81516001600160401b0381111561568f575f5ffd5b8201606081850312156156a0575f5ffd5b6156a86140ce565b81516156b38161418c565b81526020828101519082015260408201516001600160401b038111156156d7575f5ffd5b61428e8682850161535b565b6156ed8184614e40565b606060408201525f613e796060830184614c44565b5f82601f830112615711575f5ffd5b815161571f6141e08261435a565b8082825260208201915060208360051b860101925085831115615740575f5ffd5b602085015b8381101561451d57805161575881614607565b835260209283019201615745565b5f5f60408385031215615777575f5ffd5b82516001600160401b0381111561578c575f5ffd5b615798858286016153a8565b92505060208301516001600160401b038111156157b3575f5ffd5b6157bf85828601615702565b9150509250929050565b5f602082840312156157d9575f5ffd5b815161127d81614728565b634e487b7160e01b5f52603260045260245ffd5b5f8151808452602084019350602083015f5b8281101561582c57815161ffff1686526020958601959091019060010161580a565b5093949350505050565b6158408185614e40565b608060408201525f6158556080830185614ca6565b8281036060840152614f8181856157f8565b5f60208284031215615877575f5ffd5b815161127d81614752565b61588c8185614e40565b608060408201525f6158556080830185614c44565b5f5f604083850312156158b2575f5ffd5b82516158bd81614752565b60208401519092506001600160401b038111156157b3575f5ffd5b5f82518060208501845e5f92019182525091905056fea2646970667358221220006101586bc87eb74a0ad014af2d99348c9a5054e4845ebb8ab131bf40b5fb7164736f6c634300081b0033",
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
