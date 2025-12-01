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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_bn254CertificateVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_ecdsaCertificateVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_maxTaskSLA\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"BN254_CERTIFICATE_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ECDSA_CERTIFICATE_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_TASK_SLA\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createTask\",\"inputs\":[{\"name\":\"taskParams\",\"type\":\"tuple\",\"internalType\":\"structITaskMailboxTypes.TaskParams\",\"components\":[{\"name\":\"refundCollector\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"executorOperatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"feeSplit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"feeSplitCollector\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBN254CertificateBytes\",\"inputs\":[{\"name\":\"cert\",\"type\":\"tuple\",\"internalType\":\"structIBN254CertificateVerifierTypes.BN254Certificate\",\"components\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"nonSignerWitnesses\",\"type\":\"tuple[]\",\"internalType\":\"structIBN254CertificateVerifierTypes.BN254OperatorInfoWitness[]\",\"components\":[{\"name\":\"operatorIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorInfoProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"operatorInfo\",\"type\":\"tuple\",\"internalType\":\"structIOperatorTableCalculatorTypes.BN254OperatorInfo\",\"components\":[{\"name\":\"pubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}]}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getECDSACertificateBytes\",\"inputs\":[{\"name\":\"cert\",\"type\":\"tuple\",\"internalType\":\"structIECDSACertificateVerifierTypes.ECDSACertificate\",\"components\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sig\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getExecutorOperatorSetTaskConfig\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structITaskMailboxTypes.ExecutorOperatorSetTaskConfig\",\"components\":[{\"name\":\"taskHook\",\"type\":\"address\",\"internalType\":\"contractIAVSTaskHook\"},{\"name\":\"taskSLA\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"curveType\",\"type\":\"uint8\",\"internalType\":\"enumIKeyRegistrarTypes.CurveType\"},{\"name\":\"feeCollector\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"consensus\",\"type\":\"tuple\",\"internalType\":\"structITaskMailboxTypes.Consensus\",\"components\":[{\"name\":\"consensusType\",\"type\":\"uint8\",\"internalType\":\"enumITaskMailboxTypes.ConsensusType\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"taskMetadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMessageHash\",\"inputs\":[{\"name\":\"taskHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"result\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getTaskInfo\",\"inputs\":[{\"name\":\"taskHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structITaskMailboxTypes.Task\",\"components\":[{\"name\":\"creator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"creationTime\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avsFee\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"refundCollector\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"executorOperatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeSplit\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumITaskMailboxTypes.TaskStatus\"},{\"name\":\"isFeeRefunded\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"operatorTableReferenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"executorOperatorSetTaskConfig\",\"type\":\"tuple\",\"internalType\":\"structITaskMailboxTypes.ExecutorOperatorSetTaskConfig\",\"components\":[{\"name\":\"taskHook\",\"type\":\"address\",\"internalType\":\"contractIAVSTaskHook\"},{\"name\":\"taskSLA\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"curveType\",\"type\":\"uint8\",\"internalType\":\"enumIKeyRegistrarTypes.CurveType\"},{\"name\":\"feeCollector\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"consensus\",\"type\":\"tuple\",\"internalType\":\"structITaskMailboxTypes.Consensus\",\"components\":[{\"name\":\"consensusType\",\"type\":\"uint8\",\"internalType\":\"enumITaskMailboxTypes.ConsensusType\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"taskMetadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"executorCert\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"result\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskResult\",\"inputs\":[{\"name\":\"taskHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskStatus\",\"inputs\":[{\"name\":\"taskHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumITaskMailboxTypes.TaskStatus\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_feeSplit\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"_feeSplitCollector\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isExecutorOperatorSetRegistered\",\"inputs\":[{\"name\":\"operatorSetKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"isRegistered\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"refundFee\",\"inputs\":[{\"name\":\"taskHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerExecutorOperatorSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"isRegistered\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setExecutorOperatorSetTaskConfig\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structITaskMailboxTypes.ExecutorOperatorSetTaskConfig\",\"components\":[{\"name\":\"taskHook\",\"type\":\"address\",\"internalType\":\"contractIAVSTaskHook\"},{\"name\":\"taskSLA\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"curveType\",\"type\":\"uint8\",\"internalType\":\"enumIKeyRegistrarTypes.CurveType\"},{\"name\":\"feeCollector\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"consensus\",\"type\":\"tuple\",\"internalType\":\"structITaskMailboxTypes.Consensus\",\"components\":[{\"name\":\"consensusType\",\"type\":\"uint8\",\"internalType\":\"enumITaskMailboxTypes.ConsensusType\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"taskMetadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFeeSplit\",\"inputs\":[{\"name\":\"_feeSplit\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFeeSplitCollector\",\"inputs\":[{\"name\":\"_feeSplitCollector\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitResult\",\"inputs\":[{\"name\":\"taskHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"executorCert\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"result\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ExecutorOperatorSetRegistered\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"executorOperatorSetId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"isRegistered\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExecutorOperatorSetTaskConfigSet\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"executorOperatorSetId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structITaskMailboxTypes.ExecutorOperatorSetTaskConfig\",\"components\":[{\"name\":\"taskHook\",\"type\":\"address\",\"internalType\":\"contractIAVSTaskHook\"},{\"name\":\"taskSLA\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"curveType\",\"type\":\"uint8\",\"internalType\":\"enumIKeyRegistrarTypes.CurveType\"},{\"name\":\"feeCollector\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"consensus\",\"type\":\"tuple\",\"internalType\":\"structITaskMailboxTypes.Consensus\",\"components\":[{\"name\":\"consensusType\",\"type\":\"uint8\",\"internalType\":\"enumITaskMailboxTypes.ConsensusType\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"taskMetadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeRefunded\",\"inputs\":[{\"name\":\"refundCollector\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"taskHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"avsFee\",\"type\":\"uint96\",\"indexed\":false,\"internalType\":\"uint96\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeSplitCollectorSet\",\"inputs\":[{\"name\":\"feeSplitCollector\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeSplitSet\",\"inputs\":[{\"name\":\"feeSplit\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskCreated\",\"inputs\":[{\"name\":\"creator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"taskHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"executorOperatorSetId\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"operatorTableReferenceTimestamp\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"refundCollector\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"avsFee\",\"type\":\"uint96\",\"indexed\":false,\"internalType\":\"uint96\"},{\"name\":\"taskDeadline\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskVerified\",\"inputs\":[{\"name\":\"aggregator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"taskHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"executorOperatorSetId\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"executorCert\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"result\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CertificateStale\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EmptyCertificateSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExecutorOperatorSetNotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExecutorOperatorSetTaskConfigNotSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FeeAlreadyRefunded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidConsensusType\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidConsensusValue\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCurveType\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidFeeReceiver\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidFeeSplit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidMessageHash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperatorSetOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidReferenceTimestamp\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTaskCreator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTaskStatus\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"enumITaskMailboxTypes.TaskStatus\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"enumITaskMailboxTypes.TaskStatus\"}]},{\"type\":\"error\",\"name\":\"OnlyRefundCollector\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PayloadIsEmpty\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TaskSLAExceedsMaximum\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ThresholdNotMet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TimestampAtCreation\",\"inputs\":[]}]",
	Bin: "0x60e060405234801561000f575f5ffd5b50604051615afb380380615afb83398101604081905261002e91610135565b6001600160a01b03808416608052821660a0526001600160601b03811660c05261005661005e565b505050610185565b5f54610100900460ff16156100c95760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff90811614610118575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b80516001600160a01b0381168114610130575f5ffd5b919050565b5f5f5f60608486031215610147575f5ffd5b6101508461011a565b925061015e6020850161011a565b60408501519092506001600160601b038116811461017a575f5ffd5b809150509250925092565b60805160a05160c05161591b6101e05f395f81816103450152612ece01525f81816102650152818161323d015281816138120152613a3501525f81816103aa015281816131fd015281816137400152613962015261591b5ff3fe608060405234801561000f575f5ffd5b5060043610610153575f3560e01c8063678fbdb3116100bf578063b869416611610079578063b86941661461032d578063d3e043aa14610340578063f2fde38b1461037f578063f741e81a14610392578063f7424fc9146103a5578063fa2c0b37146103cc575f5ffd5b8063678fbdb3146102bb5780636bf6fad5146102ce578063708c0db9146102ee578063715018a6146103015780638da5cb5b14610309578063a5fabc811461031a575f5ffd5b8063468c07a011610110578063468c07a01461021857806349acd8841461022d5780634ad52e021461024057806354743ad21461026057806362fee037146102875780636373ea691461029a575f5ffd5b80631270a892146101575780631a20c505146101805780631ae370eb146101b15780631fb66f5d146101c45780632bf6cc79146101e557806337eaa10414610205575b5f5ffd5b61016a6101653660046141bb565b6103fe565b6040516101779190614270565b60405180910390f35b609b54610199906201000090046001600160a01b031681565b6040516001600160a01b039091168152602001610177565b61016a6101bf3660046144cd565b610427565b6101d76101d236600461460c565b61043a565b604051908152602001610177565b6101f86101f3366004614681565b610de3565b60405161017791906146c0565b6101d76102133660046146ce565b61128f565b61022b610226366004614720565b6112c2565b005b61022b61023b366004614748565b6112d6565b61025361024e366004614681565b611521565b6040516101779190614864565b6101997f000000000000000000000000000000000000000000000000000000000000000081565b61016a610295366004614681565b611aba565b609b546102a89061ffff1681565b60405161ffff9091168152602001610177565b61022b6102c93660046149af565b611fb7565b6102e16102dc3660046149ca565b611fc8565b60405161017791906149e4565b61022b6102fc3660046149f6565b6121dc565b61022b61230d565b6033546001600160a01b0316610199565b61022b610328366004614a3e565b612320565b61022b61033b366004614681565b612bfe565b6103677f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160601b039091168152602001610177565b61022b61038d3660046149af565b612e30565b61022b6103a0366004614b30565b612ea6565b6101997f000000000000000000000000000000000000000000000000000000000000000081565b6103ee6103da366004614681565b60996020525f908152604090205460ff1681565b6040519015158152602001610177565b6060816040516020016104119190614c5b565b6040516020818303038152906040529050919050565b6060816040516020016104119190614ddf565b5f6104436130db565b5f8260400151511161046857604051636b1a1b6960e11b815260040160405180910390fd5b60995f6104788460200151613134565b815260208101919091526040015f205460ff166104a85760405163c292b29760e01b815260040160405180910390fd5b5f609a5f6104b98560200151613134565b815260208082019290925260409081015f20815160e08101835281546001600160a01b038082168352600160a01b918290046001600160601b03169583019590955260018301549485169382019390935292909160608401910460ff16600281111561052757610527614698565b600281111561053857610538614698565b815260028201546001600160a01b0316602082015260408051808201825260038401805492909301929091829060ff16600181111561057957610579614698565b600181111561058a5761058a614698565b815260200160018201805461059e90614df1565b80601f01602080910402602001604051908101604052809291908181526020018280546105ca90614df1565b80156106155780601f106105ec57610100808354040283529160200191610615565b820191905f5260205f20905b8154815290600101906020018083116105f857829003601f168201915b505050505081525050815260200160058201805461063290614df1565b80601f016020809104026020016040519081016040528092919081815260200182805461065e90614df1565b80156106a95780601f10610680576101008083540402835291602001916106a9565b820191905f5260205f20905b81548152906001019060200180831161068c57829003601f168201915b50505050508152505090506106bd81613197565b6106da576040516314b0a41d60e11b815260040160405180910390fd5b5f6106e882606001516131e0565b90505f816001600160a01b0316635ddb9b5b86602001516040518263ffffffff1660e01b815260040161071b9190614e47565b602060405180830381865afa158015610736573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061075a9190614e55565b90505f826001600160a01b0316636141879e87602001516040518263ffffffff1660e01b815260040161078d9190614e47565b602060405180830381865afa1580156107a8573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107cc9190614e55565b905063ffffffff8116158061080857506107e68183614e84565b63ffffffff1684602001516001600160601b0316426108059190614ea0565b11155b6108255760405163640fcd6b60e11b815260040160405180910390fd5b508251604051630a3fc61360e31b81526001600160a01b03909116906351fe3098906108579033908990600401614eee565b5f6040518083038186803b15801561086d575f5ffd5b505afa15801561087f573d5f5f3e3d5ffd5b50508451604051637036693f60e11b81525f93506001600160a01b03909116915063e06cd27e906108b4908990600401614f11565b602060405180830381865afa1580156108cf573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108f39190614f23565b90505f60975430468960405160200161090f9493929190614f3e565b60405160208183030381529060405280519060200120905060975460016109369190614ea0565b609755604080516101c08101909152338152602081016109554261327a565b6001600160601b03908116825260208a81018051516001600160a01b039081168386015292871660408501528b5190921660608401529051015163ffffffff166080820152609b5461ffff1660a082015260c001600181525f602080830182905263ffffffff80881660408086019190915260608086018c90528d8201516080808801919091528251808601845286815260a0808901919091528351808701855287815260c09889015289875260988652958390208851958901516001600160601b03908116600160a01b9081026001600160a01b03988916178355948a0151938a0151168402928616929092176001830155870151600282018054968901519789015161ffff16600160c01b0261ffff60c01b19989095169093026001600160c01b0319909616941693909317939093179384168117835560e0850151919391929160ff60d01b1990911662ffffff60c01b1990911617600160d01b836003811115610ac457610ac4614698565b02179055506101008201516002808301805461012086015163ffffffff16600160e01b026001600160e01b03941515600160d81b02949094166001600160d81b039091161792909217909155610140830151805160208201516001600160601b0316600160a01b9081026001600160a01b0392831617600386019081556040840151600487018054919094166001600160a01b031982168117855560608601519596929594936001600160a81b031990921617918490811115610b8957610b89614698565b021790555060808201516002820180546001600160a01b0319166001600160a01b0390921691909117905560a08201518051600383018054909190829060ff191660018381811115610bdd57610bdd614698565b021790555060208201516001820190610bf69082614fbf565b50505060c08201516005820190610c0d9082614fbf565b5050506101608201516009820190610c259082614fbf565b50610180820151600a820190610c3b9082614fbf565b506101a0820151600b820190610c519082614fbf565b50505060408501516001600160a01b031615801590610c7857505f826001600160601b0316115b15610cf35760808501516001600160a01b0316610ca857604051633480121760e21b815260040160405180910390fd5b86516001600160a01b0316610cd057604051633480121760e21b815260040160405180910390fd5b6040850151610cf3906001600160a01b031633306001600160601b0386166132e5565b8451604051629c5c4560e41b8152600481018390526001600160a01b03909116906309c5c450906024015f604051808303815f87803b158015610d34575f5ffd5b505af1158015610d46573d5f5f3e3d5ffd5b5050505086602001515f01516001600160a01b031681336001600160a01b03167f33add0b01e02278be5459fbfa3274aee699ec47f4ee7236b59e7a2c8b5000c268a6020015160200151878c5f0151888c602001516001600160601b031642610daf9190614ea0565b8f60400151604051610dc696959493929190615079565b60405180910390a4945050505050610dde6001606555565b919050565b5f81815260986020908152604080832081516101c08101835281546001600160a01b038082168352600160a01b918290046001600160601b03908116968401969096526001840154808216958401959095529381900490941660608201526002820154928316608082015292820463ffffffff1660a0840152600160c01b820461ffff1660c084015283929160e0830190600160d01b900460ff166003811115610e8f57610e8f614698565b6003811115610ea057610ea0614698565b8152600282810154600160d81b810460ff9081161515602080860191909152600160e01b90920463ffffffff16604080860191909152805160e0810182526003870180546001600160a01b038082168452600160a01b918290046001600160601b03169684019690965260048901549586169383019390935260609687019691959094918601939290910490911690811115610f3e57610f3e614698565b6002811115610f4f57610f4f614698565b815260028201546001600160a01b0316602082015260408051808201825260038401805492909301929091829060ff166001811115610f9057610f90614698565b6001811115610fa157610fa1614698565b8152602001600182018054610fb590614df1565b80601f0160208091040260200160405190810160405280929190818152602001828054610fe190614df1565b801561102c5780601f106110035761010080835404028352916020019161102c565b820191905f5260205f20905b81548152906001019060200180831161100f57829003601f168201915b505050505081525050815260200160058201805461104990614df1565b80601f016020809104026020016040519081016040528092919081815260200182805461107590614df1565b80156110c05780601f10611097576101008083540402835291602001916110c0565b820191905f5260205f20905b8154815290600101906020018083116110a357829003601f168201915b50505050508152505081526020016009820180546110dd90614df1565b80601f016020809104026020016040519081016040528092919081815260200182805461110990614df1565b80156111545780601f1061112b57610100808354040283529160200191611154565b820191905f5260205f20905b81548152906001019060200180831161113757829003601f168201915b50505050508152602001600a8201805461116d90614df1565b80601f016020809104026020016040519081016040528092919081815260200182805461119990614df1565b80156111e45780601f106111bb576101008083540402835291602001916111e4565b820191905f5260205f20905b8154815290600101906020018083116111c757829003601f168201915b50505050508152602001600b820180546111fd90614df1565b80601f016020809104026020016040519081016040528092919081815260200182805461122990614df1565b80156112745780601f1061124b57610100808354040283529160200191611274565b820191905f5260205f20905b81548152906001019060200180831161125757829003601f168201915b505050505081525050905061128881613357565b9392505050565b5f82826040516020016112a39291906150d2565b6040516020818303038152906040528051906020012090505b92915050565b6112ca6133b2565b6112d38161340c565b50565b5f609a5f6112e385613134565b815260208082019290925260409081015f20815160e08101835281546001600160a01b038082168352600160a01b918290046001600160601b03169583019590955260018301549485169382019390935292909160608401910460ff16600281111561135157611351614698565b600281111561136257611362614698565b815260028201546001600160a01b0316602082015260408051808201825260038401805492909301929091829060ff1660018111156113a3576113a3614698565b60018111156113b4576113b4614698565b81526020016001820180546113c890614df1565b80601f01602080910402602001604051908101604052809291908181526020018280546113f490614df1565b801561143f5780601f106114165761010080835404028352916020019161143f565b820191905f5260205f20905b81548152906001019060200180831161142257829003601f168201915b505050505081525050815260200160058201805461145c90614df1565b80601f016020809104026020016040519081016040528092919081815260200182805461148890614df1565b80156114d35780601f106114aa576101008083540402835291602001916114d3565b820191905f5260205f20905b8154815290600101906020018083116114b657829003601f168201915b50505050508152505090506114e781613197565b611504576040516314b0a41d60e11b815260040160405180910390fd5b61151283826060015161347d565b61151c838361351e565b505050565b611529613f9b565b5f82815260986020908152604080832081516101c08101835281546001600160a01b038082168352600160a01b918290046001600160601b03908116968401969096526001840154808216958401959095529381900490941660608201526002820154928316608082015292820463ffffffff1660a0840152600160c01b820461ffff1660c08401529060e0830190600160d01b900460ff1660038111156115d3576115d3614698565b60038111156115e4576115e4614698565b8152600282810154600160d81b810460ff9081161515602080860191909152600160e01b90920463ffffffff16604080860191909152805160e0810182526003870180546001600160a01b038082168452600160a01b918290046001600160601b0316968401969096526004890154958616938301939093526060968701969195909491860193929091049091169081111561168257611682614698565b600281111561169357611693614698565b815260028201546001600160a01b0316602082015260408051808201825260038401805492909301929091829060ff1660018111156116d4576116d4614698565b60018111156116e5576116e5614698565b81526020016001820180546116f990614df1565b80601f016020809104026020016040519081016040528092919081815260200182805461172590614df1565b80156117705780601f1061174757610100808354040283529160200191611770565b820191905f5260205f20905b81548152906001019060200180831161175357829003601f168201915b505050505081525050815260200160058201805461178d90614df1565b80601f01602080910402602001604051908101604052809291908181526020018280546117b990614df1565b80156118045780601f106117db57610100808354040283529160200191611804565b820191905f5260205f20905b8154815290600101906020018083116117e757829003601f168201915b505050505081525050815260200160098201805461182190614df1565b80601f016020809104026020016040519081016040528092919081815260200182805461184d90614df1565b80156118985780601f1061186f57610100808354040283529160200191611898565b820191905f5260205f20905b81548152906001019060200180831161187b57829003601f168201915b50505050508152602001600a820180546118b190614df1565b80601f01602080910402602001604051908101604052809291908181526020018280546118dd90614df1565b80156119285780601f106118ff57610100808354040283529160200191611928565b820191905f5260205f20905b81548152906001019060200180831161190b57829003601f168201915b50505050508152602001600b8201805461194190614df1565b80601f016020809104026020016040519081016040528092919081815260200182805461196d90614df1565b80156119b85780601f1061198f576101008083540402835291602001916119b8565b820191905f5260205f20905b81548152906001019060200180831161199b57829003601f168201915b5050505050815250509050604051806101c00160405280825f01516001600160a01b0316815260200182602001516001600160601b0316815260200182604001516001600160a01b0316815260200182606001516001600160601b0316815260200182608001516001600160a01b031681526020018260a0015163ffffffff1681526020018260c0015161ffff168152602001611a5483613357565b6003811115611a6557611a65614698565b81526020018261010001511515815260200182610120015163ffffffff168152602001826101400151815260200182610160015181526020018261018001518152602001826101a00151815250915050919050565b5f81815260986020908152604080832081516101c08101835281546001600160a01b038082168352600160a01b918290046001600160601b0390811696840196909652600184015480821695840195909552938190049094166060808301919091526002830154938416608083015293830463ffffffff1660a0820152600160c01b830461ffff1660c08201529293929160e0830190600160d01b900460ff166003811115611b6b57611b6b614698565b6003811115611b7c57611b7c614698565b8152600282810154600160d81b810460ff9081161515602080860191909152600160e01b90920463ffffffff16604080860191909152805160e0810182526003870180546001600160a01b038082168452600160a01b918290046001600160601b03169684019690965260048901549586169383019390935260609687019691959094918601939290910490911690811115611c1a57611c1a614698565b6002811115611c2b57611c2b614698565b815260028201546001600160a01b0316602082015260408051808201825260038401805492909301929091829060ff166001811115611c6c57611c6c614698565b6001811115611c7d57611c7d614698565b8152602001600182018054611c9190614df1565b80601f0160208091040260200160405190810160405280929190818152602001828054611cbd90614df1565b8015611d085780601f10611cdf57610100808354040283529160200191611d08565b820191905f5260205f20905b815481529060010190602001808311611ceb57829003601f168201915b5050505050815250508152602001600582018054611d2590614df1565b80601f0160208091040260200160405190810160405280929190818152602001828054611d5190614df1565b8015611d9c5780601f10611d7357610100808354040283529160200191611d9c565b820191905f5260205f20905b815481529060010190602001808311611d7f57829003601f168201915b5050505050815250508152602001600982018054611db990614df1565b80601f0160208091040260200160405190810160405280929190818152602001828054611de590614df1565b8015611e305780601f10611e0757610100808354040283529160200191611e30565b820191905f5260205f20905b815481529060010190602001808311611e1357829003601f168201915b50505050508152602001600a82018054611e4990614df1565b80601f0160208091040260200160405190810160405280929190818152602001828054611e7590614df1565b8015611ec05780601f10611e9757610100808354040283529160200191611ec0565b820191905f5260205f20905b815481529060010190602001808311611ea357829003601f168201915b50505050508152602001600b82018054611ed990614df1565b80601f0160208091040260200160405190810160405280929190818152602001828054611f0590614df1565b8015611f505780601f10611f2757610100808354040283529160200191611f50565b820191905f5260205f20905b815481529060010190602001808311611f3357829003601f168201915b50505050508152505090505f611f6582613357565b90506002816003811115611f7b57611f7b614698565b146002829091611fa957604051634091b18960e11b8152600401611fa09291906150ea565b60405180910390fd5b5050506101a0015192915050565b611fbf6133b2565b6112d3816135af565b611fd061400e565b609a5f611fdc84613134565b815260208082019290925260409081015f20815160e08101835281546001600160a01b038082168352600160a01b918290046001600160601b03169583019590955260018301549485169382019390935292909160608401910460ff16600281111561204a5761204a614698565b600281111561205b5761205b614698565b815260028201546001600160a01b0316602082015260408051808201825260038401805492909301929091829060ff16600181111561209c5761209c614698565b60018111156120ad576120ad614698565b81526020016001820180546120c190614df1565b80601f01602080910402602001604051908101604052809291908181526020018280546120ed90614df1565b80156121385780601f1061210f57610100808354040283529160200191612138565b820191905f5260205f20905b81548152906001019060200180831161211b57829003601f168201915b505050505081525050815260200160058201805461215590614df1565b80601f016020809104026020016040519081016040528092919081815260200182805461218190614df1565b80156121cc5780601f106121a3576101008083540402835291602001916121cc565b820191905f5260205f20905b8154815290600101906020018083116121af57829003601f168201915b5050505050815250509050919050565b5f54610100900460ff16158080156121fa57505f54600160ff909116105b806122135750303b15801561221357505f5460ff166001145b6122765760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401611fa0565b5f805460ff191660011790558015612297575f805461ff0019166101001790555b61229f613629565b6122a7613657565b6122b084613685565b6122b98361340c565b6122c2826135af565b8015612307575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050565b6123156133b2565b61231e5f613685565b565b6123286130db565b5f83815260986020908152604080832081516101c08101835281546001600160a01b038082168352600160a01b918290046001600160601b03908116968401969096526001840154808216958401959095529381900490941660608201526002820154928316608082015292820463ffffffff1660a0840152600160c01b820461ffff1660c084015292916127cc91849060e0830190600160d01b900460ff1660038111156123d9576123d9614698565b60038111156123ea576123ea614698565b8152600282810154600160d81b810460ff9081161515602080860191909152600160e01b90920463ffffffff16604080860191909152805160e0810182526003870180546001600160a01b038082168452600160a01b918290046001600160601b0316968401969096526004890154958616938301939093526060968701969195909491860193929091049091169081111561248857612488614698565b600281111561249957612499614698565b815260028201546001600160a01b0316602082015260408051808201825260038401805492909301929091829060ff1660018111156124da576124da614698565b60018111156124eb576124eb614698565b81526020016001820180546124ff90614df1565b80601f016020809104026020016040519081016040528092919081815260200182805461252b90614df1565b80156125765780601f1061254d57610100808354040283529160200191612576565b820191905f5260205f20905b81548152906001019060200180831161255957829003601f168201915b505050505081525050815260200160058201805461259390614df1565b80601f01602080910402602001604051908101604052809291908181526020018280546125bf90614df1565b801561260a5780601f106125e15761010080835404028352916020019161260a565b820191905f5260205f20905b8154815290600101906020018083116125ed57829003601f168201915b505050505081525050815260200160098201805461262790614df1565b80601f016020809104026020016040519081016040528092919081815260200182805461265390614df1565b801561269e5780601f106126755761010080835404028352916020019161269e565b820191905f5260205f20905b81548152906001019060200180831161268157829003601f168201915b50505050508152602001600a820180546126b790614df1565b80601f01602080910402602001604051908101604052809291908181526020018280546126e390614df1565b801561272e5780601f106127055761010080835404028352916020019161272e565b820191905f5260205f20905b81548152906001019060200180831161271157829003601f168201915b50505050508152602001600b8201805461274790614df1565b80601f016020809104026020016040519081016040528092919081815260200182805461277390614df1565b80156127be5780601f10612795576101008083540402835291602001916127be565b820191905f5260205f20905b8154815290600101906020018083116127a157829003601f168201915b505050505081525050613357565b905060018160038111156127e2576127e2614698565b14600182909161280757604051634091b18960e11b8152600401611fa09291906150ea565b50508154600160a01b90046001600160601b0316421161283a5760405163015a4b7560e51b815260040160405180910390fd5b600382015460405163ba33565d60e01b81526001600160a01b039091169063ba33565d90612872903390899089908990600401615105565b5f6040518083038186803b158015612888575f5ffd5b505afa15801561289a573d5f5f3e3d5ffd5b50506040805180820182526001808701546001600160a01b03168252600287015463ffffffff600160a01b91829004166020840152600488015484518086019095526006890180549497506129d6965060ff929091048216949390928492919091169081111561290c5761290c614698565b600181111561291d5761291d614698565b815260200160018201805461293190614df1565b80601f016020809104026020016040519081016040528092919081815260200182805461295d90614df1565b80156129a85780601f1061297f576101008083540402835291602001916129a8565b820191905f5260205f20905b81548152906001019060200180831161298b57829003601f168201915b5050509190925250505060028601548490600160e01b900463ffffffff166129d08b8a61128f565b8a6136d6565b60028301805460ff60d01b1916600160d11b179055600a83016129f98682614fbf565b50600b8301612a088582614fbf565b5060048301546001600160a01b031615801590612a3857506001830154600160a01b90046001600160601b031615155b15612b2757600283015460018401545f91612a879161271091612a7891600160c01b90910461ffff16906001600160601b03600160a01b9091041661513d565b612a829190615154565b61327a565b90506001600160601b03811615612ac757609b546004850154612ac7916001600160a01b039182169162010000909104166001600160601b038416613af8565b60018401545f90612ae9908390600160a01b90046001600160601b0316615173565b90506001600160601b03811615612b245760058501546004860154612b24916001600160a01b0391821691166001600160601b038416613af8565b50505b6003830154604051637041233f60e11b8152336004820152602481018890526001600160a01b039091169063e082467e906044015f604051808303815f87803b158015612b72575f5ffd5b505af1158015612b84573d5f5f3e3d5ffd5b505050600184015460028501546040516001600160a01b039092169250889133917f659f23b2e7edf490e5fd6561c5148691ed0375ed7ddd3ab1bcfcfdbec4f209a991612be99163ffffffff600160a01b9091041690600a8a0190600b8b0190615211565b60405180910390a450505061151c6001606555565b612c066130db565b5f81815260986020526040902060028101546001600160a01b03163314612c40576040516370f43cb760e01b815260040160405180910390fd5b6002810154600160d81b900460ff1615612c6d57604051633e3d786960e01b815260040160405180910390fd5b604080516101c08101825282546001600160a01b038082168352600160a01b918290046001600160601b03908116602085015260018601548083169585019590955293829004909316606083015260028401549283166080830152820463ffffffff1660a0820152600160c01b820461ffff1660c08201525f91612d0e9190849060e0830190600160d01b900460ff1660038111156123d9576123d9614698565b90506003816003811115612d2457612d24614698565b146003829091612d4957604051634091b18960e11b8152600401611fa09291906150ea565b505060028201805460ff60d81b1916600160d81b17905560048201546001600160a01b031615801590612d8f57506001820154600160a01b90046001600160601b031615155b15612dcb57600282015460018301546004840154612dcb926001600160a01b0391821692911690600160a01b90046001600160601b0316613af8565b60028201546001830154604051600160a01b9091046001600160601b0316815284916001600160a01b0316907fe3ed40d31808582f7a92a30beacc0ec788d5091407ec6c10c1b999b3f317aea39060200160405180910390a350506112d36001606555565b612e386133b2565b6001600160a01b038116612e9d5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401611fa0565b6112d381613685565b612eaf81613197565b612ecc576040516314b0a41d60e11b815260040160405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160601b031681602001516001600160601b03161115612f235760405163a7cf966560e01b815260040160405180910390fd5b612f308160a00151613b28565b612f3e82826060015161347d565b80609a5f612f4b85613134565b815260208082019290925260409081015f208351928401516001600160601b0316600160a01b9081026001600160a01b0394851617825591840151600182018054919094166001600160a01b03198216811785556060860151929492936001600160a81b03199092161790836002811115612fc857612fc8614698565b021790555060808201516002820180546001600160a01b0319166001600160a01b0390921691909117905560a08201518051600383018054909190829060ff19166001838181111561301c5761301c614698565b0217905550602082015160018201906130359082614fbf565b50505060c0820151600582019061304c9082614fbf565b50905050816020015163ffffffff16825f01516001600160a01b0316336001600160a01b03167f7cd76abd4025a20959a1b20f7c1536e3894a0735cd8de0215dde803ddea7f2d2846040516130a191906149e4565b60405180910390a460995f6130b584613134565b815260208101919091526040015f205460ff166130d7576130d782600161351e565b5050565b60026065540361312d5760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401611fa0565b6002606555565b5f815f0151826020015163ffffffff1660405160200161317f92919060609290921b6bffffffffffffffffffffffff1916825260a01b6001600160a01b031916601482015260200190565b6040516020818303038152906040526112bc90615241565b5f80826060015160028111156131af576131af614698565b141580156131c6575081516001600160a01b031615155b80156112bc575050602001516001600160601b0316151590565b5f60028260028111156131f5576131f5614698565b0361322157507f0000000000000000000000000000000000000000000000000000000000000000919050565b600182600281111561323557613235614698565b0361326157507f0000000000000000000000000000000000000000000000000000000000000000919050565b60405163fdea7c0960e01b815260040160405180910390fd5b5f6001600160601b038211156132e15760405162461bcd60e51b815260206004820152602660248201527f53616665436173743a2076616c756520646f65736e27742066697420696e203960448201526536206269747360d01b6064820152608401611fa0565b5090565b6040516001600160a01b03808516602483015283166044820152606481018290526123079085906323b872dd60e01b906084015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b031990931692909217909152613be7565b6001606555565b5f60018260e00151600381111561337057613370614698565b14801561339d57508161014001516020015182602001516133919190615264565b6001600160601b031642115b156133aa57506003919050565b5060e0015190565b6033546001600160a01b0316331461231e5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401611fa0565b61271061ffff8216111561343357604051630601f69760e01b815260040160405180910390fd5b609b805461ffff191661ffff83169081179091556040519081527f886b2cfcb151fd8b19ed902cc88f4a06dd9fe351a4a9ab93f33fe84abc157edf9060200160405180910390a150565b5f613487826131e0565b6040516304240c4960e51b815290915033906001600160a01b038316906384818920906134b8908790600401614e47565b602060405180830381865afa1580156134d3573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906134f79190615283565b6001600160a01b03161461151c576040516342ecfee960e11b815260040160405180910390fd5b8060995f61352b85613134565b81526020019081526020015f205f6101000a81548160ff021916908315150217905550816020015163ffffffff16825f01516001600160a01b0316336001600160a01b03167f48b63f21a1eb9dd6880e196de6d7db3fbd0c282b74f1298dcb4cf53472298f39846040516135a3911515815260200190565b60405180910390a45050565b6001600160a01b0381166135d657604051630863a45360e11b815260040160405180910390fd5b609b805462010000600160b01b031916620100006001600160a01b038416908102919091179091556040517f262aa27c244f6f0088cb3092548a0adcaddedf459070a9ccab2dc6a07abe701d905f90a250565b5f54610100900460ff1661364f5760405162461bcd60e51b8152600401611fa09061529e565b61231e613cba565b5f54610100900460ff1661367d5760405162461bcd60e51b8152600401611fa09061529e565b61231e613ce9565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b5f855160018111156136ea576136ea614698565b0361389257600286600281111561370357613703614698565b036137c0575f8180602001905181019061371d9190615549565b905061372a818585613d0f565b6040516280b71560e41b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063080b715090613777908890859060040161561b565b5f604051808303815f875af1158015613792573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526137b9919081019061563a565b5050613af0565b60018660028111156137d4576137d4614698565b03613261575f818060200190518101906137ee919061566b565b90506137fb818585613d99565b6040516380c7d3f360e01b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906380c7d3f39061384990889085906004016156e4565b5f60405180830381865afa158015613863573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261388a9190810190615767565b505050613af0565b6001855160018111156138a7576138a7614698565b03613ad7575f85602001518060200190518101906138c591906157c0565b6040805160018082528183019092529192505f91906020808301908036833701905050905081815f815181106138fd576138fd6157db565b61ffff909216602092830291909101909101525f600289600281111561392557613925614698565b036139e3575f8480602001905181019061393f9190615549565b905061394c818888613d0f565b604051625f5e5d60e21b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063017d79749061399b908b908590889060040161582d565b6020604051808303815f875af11580156139b7573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906139db919061585e565b915050613ab9565b60018960028111156139f7576139f7614698565b03613261575f84806020019051810190613a11919061566b565b9050613a1e818888613d99565b604051630606d12160e51b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063c0da242090613a6e908b9085908890600401615879565b5f60405180830381865afa158015613a88573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052613aaf9190810190615898565b509150613ab99050565b8061388a576040516359fa4a9360e01b815260040160405180910390fd5b6040516347d3772160e11b815260040160405180910390fd5b505050505050565b6040516001600160a01b03831660248201526044810182905261151c90849063a9059cbb60e01b90606401613319565b5f81516001811115613b3c57613b3c614698565b03613b6557602081015151156112d357604051631501e04760e21b815260040160405180910390fd5b600181516001811115613b7a57613b7a614698565b03613ad757806020015151602014613ba557604051631501e04760e21b815260040160405180910390fd5b5f8160200151806020019051810190613bbe91906157c0565b905061271061ffff821611156130d757604051631501e04760e21b815260040160405180910390fd5b5f613c3b826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316613e119092919063ffffffff16565b905080515f1480613c5b575080806020019051810190613c5b919061585e565b61151c5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608401611fa0565b5f54610100900460ff16613ce05760405162461bcd60e51b8152600401611fa09061529e565b61231e33613685565b5f54610100900460ff166133505760405162461bcd60e51b8152600401611fa09061529e565b8163ffffffff16835f015163ffffffff1614613d3e57604051634534032960e01b815260040160405180910390fd5b80836020015114613d6257604051638b56642d60e01b815260040160405180910390fd5b604083015151158015613d7b5750604083015160200151155b1561151c57604051637a8a1dbd60e11b815260040160405180910390fd5b8163ffffffff16835f015163ffffffff1614613dc857604051634534032960e01b815260040160405180910390fd5b80836020015114613dec57604051638b56642d60e01b815260040160405180910390fd5b5f8360400151511161151c57604051637a8a1dbd60e11b815260040160405180910390fd5b6060613e1f84845f85613e27565b949350505050565b606082471015613e885760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b6064820152608401611fa0565b5f5f866001600160a01b03168587604051613ea391906158cf565b5f6040518083038185875af1925050503d805f8114613edd576040519150601f19603f3d011682016040523d82523d5f602084013e613ee2565b606091505b5091509150613ef387838387613efe565b979650505050505050565b60608315613f6c5782515f03613f65576001600160a01b0385163b613f655760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401611fa0565b5081613e1f565b613e1f8383815115613f815781518083602001fd5b8060405162461bcd60e51b8152600401611fa09190614270565b604080516101c0810182525f80825260208201819052918101829052606081018290526080810182905260a0810182905260c081018290529060e082019081525f602082018190526040820152606001613ff361400e565b81526020016060815260200160608152602001606081525090565b6040805160e0810182525f8082526020820181905291810182905290606082019081525f602082015260400161404261404f565b8152602001606081525090565b60408051808201909152805f614042565b634e487b7160e01b5f52604160045260245ffd5b604051606081016001600160401b038111828210171561409657614096614060565b60405290565b604080519081016001600160401b038111828210171561409657614096614060565b60405160a081016001600160401b038111828210171561409657614096614060565b60405160e081016001600160401b038111828210171561409657614096614060565b604051601f8201601f191681016001600160401b038111828210171561412a5761412a614060565b604052919050565b63ffffffff811681146112d3575f5ffd5b5f6001600160401b0382111561415b5761415b614060565b50601f01601f191660200190565b5f82601f830112614178575f5ffd5b813561418b61418682614143565b614102565b81815284602083860101111561419f575f5ffd5b816020850160208301375f918101602001919091529392505050565b5f602082840312156141cb575f5ffd5b81356001600160401b038111156141e0575f5ffd5b8201606081850312156141f1575f5ffd5b6141f9614074565b813561420481614132565b81526020828101359082015260408201356001600160401b03811115614228575f5ffd5b61423486828501614169565b604083015250949350505050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f6112886020830184614242565b5f60408284031215614292575f5ffd5b61429a61409c565b823581526020928301359281019290925250919050565b5f82601f8301126142c0575f5ffd5b6142ca6040614102565b8060408401858111156142db575f5ffd5b845b818110156142f55780358452602093840193016142dd565b509095945050505050565b5f6001600160401b0382111561431857614318614060565b5060051b60200190565b5f82601f830112614331575f5ffd5b813561433f61418682614300565b8082825260208201915060208360051b860101925085831115614360575f5ffd5b602085015b838110156144c35780356001600160401b03811115614382575f5ffd5b86016060818903601f19011215614397575f5ffd5b61439f614074565b60208201356143ad81614132565b815260408201356001600160401b038111156143c7575f5ffd5b6143d68a602083860101614169565b60208301525060608201356001600160401b038111156143f4575f5ffd5b6020818401019250506060828a03121561440c575f5ffd5b61441461409c565b61441e8a84614282565b815260408301356001600160401b03811115614438575f5ffd5b80840193505089601f84011261444c575f5ffd5b823561445a61418682614300565b8082825260208201915060208360051b87010192508c83111561447b575f5ffd5b6020860195505b8286101561449d578535825260209586019590910190614482565b806020850152505050806040830152508085525050602083019250602081019050614365565b5095945050505050565b5f602082840312156144dd575f5ffd5b81356001600160401b038111156144f2575f5ffd5b8201808403610120811215614505575f5ffd5b61450d6140be565b823561451881614132565b8152602083810135908201526145318660408501614282565b60408201526080607f1983011215614547575f5ffd5b61454f61409c565b915061455e86608085016142b1565b825261456d8660c085016142b1565b602083015281606082015261010083013591506001600160401b03821115614593575f5ffd5b61459f86838501614322565b608082015295945050505050565b6001600160a01b03811681146112d3575f5ffd5b8035610dde816145ad565b5f604082840312156145dc575f5ffd5b6145e461409c565b905081356145f1816145ad565b8152602082013561460181614132565b602082015292915050565b5f6020828403121561461c575f5ffd5b81356001600160401b03811115614631575f5ffd5b820160808185031215614642575f5ffd5b61464a614074565b8135614655816145ad565b815261466485602084016145cc565b602082015260608201356001600160401b03811115614228575f5ffd5b5f60208284031215614691575f5ffd5b5035919050565b634e487b7160e01b5f52602160045260245ffd5b600481106146bc576146bc614698565b9052565b602081016112bc82846146ac565b5f5f604083850312156146df575f5ffd5b8235915060208301356001600160401b038111156146fb575f5ffd5b61470785828601614169565b9150509250929050565b61ffff811681146112d3575f5ffd5b5f60208284031215614730575f5ffd5b813561128881614711565b80151581146112d3575f5ffd5b5f5f60608385031215614759575f5ffd5b61476384846145cc565b915060408301356147738161473b565b809150509250929050565b600381106146bc576146bc614698565b5f8151600281106147a1576147a1614698565b80845250602082015160406020850152613e1f6040850182614242565b80516001600160a01b031682526020808201516001600160601b0316908301526040808201515f916147fa908501826001600160a01b03169052565b50606082015161480d606085018261477e565b50608082015161482860808501826001600160a01b03169052565b5060a082015160e060a085015261484260e085018261478e565b905060c083015184820360c086015261485b8282614242565b95945050505050565b6020815261487e6020820183516001600160a01b03169052565b5f602083015161489960408401826001600160601b03169052565b5060408301516001600160a01b03811660608401525060608301516001600160601b03811660808401525060808301516001600160a01b03811660a08401525060a083015163ffffffff811660c08401525060c083015161ffff811660e08401525060e083015161490e6101008401826146ac565b506101008301518015156101208401525061012083015163ffffffff8116610140840152506101408301516101c06101608401526149506101e08401826147be565b9050610160840151601f198483030161018085015261496f8282614242565b915050610180840151601f19848303016101a085015261498f8282614242565b9150506101a0840151601f19848303016101c085015261485b8282614242565b5f602082840312156149bf575f5ffd5b8135611288816145ad565b5f604082840312156149da575f5ffd5b61128883836145cc565b602081525f61128860208301846147be565b5f5f5f60608486031215614a08575f5ffd5b8335614a13816145ad565b92506020840135614a2381614711565b91506040840135614a33816145ad565b809150509250925092565b5f5f5f60608486031215614a50575f5ffd5b8335925060208401356001600160401b03811115614a6c575f5ffd5b614a7886828701614169565b92505060408401356001600160401b03811115614a93575f5ffd5b614a9f86828701614169565b9150509250925092565b6001600160601b03811681146112d3575f5ffd5b8035610dde81614aa9565b803560038110610dde575f5ffd5b5f60408284031215614ae6575f5ffd5b614aee61409c565b9050813560028110614afe575f5ffd5b815260208201356001600160401b03811115614b18575f5ffd5b614b2484828501614169565b60208301525092915050565b5f5f60608385031215614b41575f5ffd5b614b4b84846145cc565b915060408301356001600160401b03811115614b65575f5ffd5b830160e08186031215614b76575f5ffd5b614b7e6140e0565b614b87826145c1565b8152614b9560208301614abd565b6020820152614ba6604083016145c1565b6040820152614bb760608301614ac8565b6060820152614bc8608083016145c1565b608082015260a08201356001600160401b03811115614be5575f5ffd5b614bf187828501614ad6565b60a08301525060c08201356001600160401b03811115614c0f575f5ffd5b614c1b87828501614169565b60c08301525080925050509250929050565b63ffffffff8151168252602081015160208301525f604082015160606040850152613e1f6060850182614242565b602081525f6112886020830184614c2d565b805f5b6002811015612307578151845260209384019390910190600101614c70565b5f610120830163ffffffff8351168452602083015160208501526040830151614cc5604086018280518252602090810151910152565b506060830151614cd9608086018251614c6d565b60200151614cea60c0860182614c6d565b506080830151610120610100860152818151808452610140870191506101408160051b88010193506020830192505f5b81811015614dd35761013f19888603018352835163ffffffff8151168652602081015160606020880152614d516060880182614242565b905060408201519150868103604088015260608101614d7b82845180518252602090810151910152565b6020928301516060604084015280518083529301925f92608001905b80841015614dba5784518252602082019150602085019450600184019350614d97565b5097505050602094850194939093019250600101614d1a565b50929695505050505050565b602081525f6112886020830184614c8f565b600181811c90821680614e0557607f821691505b602082108103614e2357634e487b7160e01b5f52602260045260245ffd5b50919050565b80516001600160a01b0316825260209081015163ffffffff16910152565b604081016112bc8284614e29565b5f60208284031215614e65575f5ffd5b815161128881614132565b634e487b7160e01b5f52601160045260245ffd5b63ffffffff81811683821601908111156112bc576112bc614e70565b808201808211156112bc576112bc614e70565b80516001600160a01b031682526020808201515f91614ed490850182614e29565b50604082015160806060850152613e1f6080850182614242565b6001600160a01b03831681526040602082018190525f90613e1f90830184614eb3565b602081525f6112886020830184614eb3565b5f60208284031215614f33575f5ffd5b815161128881614aa9565b84815260018060a01b0384166020820152826040820152608060608201525f614f6a6080830184614eb3565b9695505050505050565b601f82111561151c57805f5260205f20601f840160051c81016020851015614f995750805b601f840160051c820191505b81811015614fb8575f8155600101614fa5565b5050505050565b81516001600160401b03811115614fd857614fd8614060565b614fec81614fe68454614df1565b84614f74565b6020601f82116001811461501e575f83156150075750848201515b5f19600385901b1c1916600184901b178455614fb8565b5f84815260208120601f198516915b8281101561504d578785015182556020948501946001909201910161502d565b508482101561506a57868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b63ffffffff8716815263ffffffff8616602082015260018060a01b03851660408201526001600160601b038416606082015282608082015260c060a08201525f6150c660c0830184614242565b98975050505050505050565b828152604060208201525f613e1f6040830184614242565b604081016150f882856146ac565b61128860208301846146ac565b60018060a01b0385168152836020820152608060408201525f61512b6080830185614242565b8281036060840152613ef38185614242565b80820281158282048414176112bc576112bc614e70565b5f8261516e57634e487b7160e01b5f52601260045260245ffd5b500490565b6001600160601b0382811682821603908111156112bc576112bc614e70565b5f815461519e81614df1565b8085526001821680156151b857600181146151d457615208565b60ff1983166020870152602082151560051b8701019350615208565b845f5260205f205f5b838110156151ff5781546020828a0101526001820191506020810190506151dd565b87016020019450505b50505092915050565b63ffffffff84168152606060208201525f61522f6060830185615192565b8281036040840152614f6a8185615192565b80516020808301519190811015614e23575f1960209190910360031b1b16919050565b6001600160601b0381811683821601908111156112bc576112bc614e70565b5f60208284031215615293575f5ffd5b8151611288816145ad565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b5f604082840312156152f9575f5ffd5b61530161409c565b825181526020928301519281019290925250919050565b5f82601f830112615327575f5ffd5b6153316040614102565b806040840185811115615342575f5ffd5b845b818110156142f5578051845260209384019301615344565b5f82601f83011261536b575f5ffd5b815161537961418682614143565b81815284602083860101111561538d575f5ffd5b8160208501602083015e5f918101602001919091529392505050565b5f82601f8301126153b8575f5ffd5b81516153c661418682614300565b8082825260208201915060208360051b8601019250858311156153e7575f5ffd5b602085015b838110156144c35780518352602092830192016153ec565b5f82601f830112615413575f5ffd5b815161542161418682614300565b8082825260208201915060208360051b860101925085831115615442575f5ffd5b602085015b838110156144c35780516001600160401b03811115615464575f5ffd5b86016060818903601f19011215615479575f5ffd5b615481614074565b602082015161548f81614132565b815260408201516001600160401b038111156154a9575f5ffd5b6154b88a60208386010161535c565b60208301525060608201516001600160401b038111156154d6575f5ffd5b6020818401019250506060828a0312156154ee575f5ffd5b6154f661409c565b6155008a846152e9565b815260408301516001600160401b0381111561551a575f5ffd5b6155268b8286016153a9565b602083015250806040830152508085525050602083019250602081019050615447565b5f60208284031215615559575f5ffd5b81516001600160401b0381111561556e575f5ffd5b8201808403610120811215615581575f5ffd5b6155896140be565b825161559481614132565b8152602083810151908201526155ad86604085016152e9565b60408201526080607f19830112156155c3575f5ffd5b6155cb61409c565b91506155da8660808501615318565b82526155e98660c08501615318565b602083015281606082015261010083015191506001600160401b0382111561560f575f5ffd5b61459f86838501615404565b6156258184614e29565b606060408201525f613e1f6060830184614c8f565b5f6020828403121561564a575f5ffd5b81516001600160401b0381111561565f575f5ffd5b613e1f848285016153a9565b5f6020828403121561567b575f5ffd5b81516001600160401b03811115615690575f5ffd5b8201606081850312156156a1575f5ffd5b6156a9614074565b81516156b481614132565b81526020828101519082015260408201516001600160401b038111156156d8575f5ffd5b6142348682850161535c565b6156ee8184614e29565b606060408201525f613e1f6060830184614c2d565b5f82601f830112615712575f5ffd5b815161572061418682614300565b8082825260208201915060208360051b860101925085831115615741575f5ffd5b602085015b838110156144c3578051615759816145ad565b835260209283019201615746565b5f5f60408385031215615778575f5ffd5b82516001600160401b0381111561578d575f5ffd5b615799858286016153a9565b92505060208301516001600160401b038111156157b4575f5ffd5b61470785828601615703565b5f602082840312156157d0575f5ffd5b815161128881614711565b634e487b7160e01b5f52603260045260245ffd5b5f8151808452602084019350602083015f5b8281101561582357815161ffff16865260209586019590910190600101615801565b5093949350505050565b6158378185614e29565b608060408201525f61584c6080830185614c8f565b8281036060840152614f6a81856157ef565b5f6020828403121561586e575f5ffd5b81516112888161473b565b6158838185614e29565b608060408201525f61584c6080830185614c2d565b5f5f604083850312156158a9575f5ffd5b82516158b48161473b565b60208401519092506001600160401b038111156157b4575f5ffd5b5f82518060208501845e5f92019182525091905056fea2646970667358221220d08a27b23e2376124b17971ed8bc82fbf516441669f287e95f2ac7f6073ee2c864736f6c634300081e0033",
}

// TaskMailboxABI is the input ABI used to generate the binding from.
// Deprecated: Use TaskMailboxMetaData.ABI instead.
var TaskMailboxABI = TaskMailboxMetaData.ABI

// TaskMailboxBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TaskMailboxMetaData.Bin instead.
var TaskMailboxBin = TaskMailboxMetaData.Bin

// DeployTaskMailbox deploys a new Ethereum contract, binding an instance of TaskMailbox to it.
func DeployTaskMailbox(auth *bind.TransactOpts, backend bind.ContractBackend, _bn254CertificateVerifier common.Address, _ecdsaCertificateVerifier common.Address, _maxTaskSLA *big.Int) (common.Address, *types.Transaction, *TaskMailbox, error) {
	parsed, err := TaskMailboxMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TaskMailboxBin), backend, _bn254CertificateVerifier, _ecdsaCertificateVerifier, _maxTaskSLA)
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
