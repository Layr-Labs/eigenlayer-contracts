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
	Creator                       common.Address
	CreationTime                  *big.Int
	Avs                           common.Address
	AvsFee                        *big.Int
	RefundCollector               common.Address
	ExecutorOperatorSetId         uint32
	FeeSplit                      uint16
	Status                        uint8
	IsFeeRefunded                 bool
	ExecutorOperatorSetTaskConfig ITaskMailboxTypesExecutorOperatorSetTaskConfig
	Payload                       []byte
	ExecutorCert                  []byte
	Result                        []byte
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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_bn254CertificateVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_ecdsaCertificateVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_version\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"BN254_CERTIFICATE_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ECDSA_CERTIFICATE_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createTask\",\"inputs\":[{\"name\":\"taskParams\",\"type\":\"tuple\",\"internalType\":\"structITaskMailboxTypes.TaskParams\",\"components\":[{\"name\":\"refundCollector\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"executorOperatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executorOperatorSetTaskConfigs\",\"inputs\":[{\"name\":\"operatorSetKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"taskHook\",\"type\":\"address\",\"internalType\":\"contractIAVSTaskHook\"},{\"name\":\"taskSLA\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"curveType\",\"type\":\"uint8\",\"internalType\":\"enumIKeyRegistrarTypes.CurveType\"},{\"name\":\"feeCollector\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"consensus\",\"type\":\"tuple\",\"internalType\":\"structITaskMailboxTypes.Consensus\",\"components\":[{\"name\":\"consensusType\",\"type\":\"uint8\",\"internalType\":\"enumITaskMailboxTypes.ConsensusType\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"taskMetadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"feeSplit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"feeSplitCollector\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBN254CertificateBytes\",\"inputs\":[{\"name\":\"cert\",\"type\":\"tuple\",\"internalType\":\"structIBN254CertificateVerifierTypes.BN254Certificate\",\"components\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"nonSignerWitnesses\",\"type\":\"tuple[]\",\"internalType\":\"structIBN254CertificateVerifierTypes.BN254OperatorInfoWitness[]\",\"components\":[{\"name\":\"operatorIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorInfoProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"operatorInfo\",\"type\":\"tuple\",\"internalType\":\"structIOperatorTableCalculatorTypes.BN254OperatorInfo\",\"components\":[{\"name\":\"pubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}]}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getECDSACertificateBytes\",\"inputs\":[{\"name\":\"cert\",\"type\":\"tuple\",\"internalType\":\"structIECDSACertificateVerifierTypes.ECDSACertificate\",\"components\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sig\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getExecutorOperatorSetTaskConfig\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structITaskMailboxTypes.ExecutorOperatorSetTaskConfig\",\"components\":[{\"name\":\"taskHook\",\"type\":\"address\",\"internalType\":\"contractIAVSTaskHook\"},{\"name\":\"taskSLA\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"curveType\",\"type\":\"uint8\",\"internalType\":\"enumIKeyRegistrarTypes.CurveType\"},{\"name\":\"feeCollector\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"consensus\",\"type\":\"tuple\",\"internalType\":\"structITaskMailboxTypes.Consensus\",\"components\":[{\"name\":\"consensusType\",\"type\":\"uint8\",\"internalType\":\"enumITaskMailboxTypes.ConsensusType\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"taskMetadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFeeSplit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFeeSplitCollector\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskInfo\",\"inputs\":[{\"name\":\"taskHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structITaskMailboxTypes.Task\",\"components\":[{\"name\":\"creator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"creationTime\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avsFee\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"refundCollector\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"executorOperatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeSplit\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumITaskMailboxTypes.TaskStatus\"},{\"name\":\"isFeeRefunded\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"executorOperatorSetTaskConfig\",\"type\":\"tuple\",\"internalType\":\"structITaskMailboxTypes.ExecutorOperatorSetTaskConfig\",\"components\":[{\"name\":\"taskHook\",\"type\":\"address\",\"internalType\":\"contractIAVSTaskHook\"},{\"name\":\"taskSLA\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"curveType\",\"type\":\"uint8\",\"internalType\":\"enumIKeyRegistrarTypes.CurveType\"},{\"name\":\"feeCollector\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"consensus\",\"type\":\"tuple\",\"internalType\":\"structITaskMailboxTypes.Consensus\",\"components\":[{\"name\":\"consensusType\",\"type\":\"uint8\",\"internalType\":\"enumITaskMailboxTypes.ConsensusType\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"taskMetadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"executorCert\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"result\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskResult\",\"inputs\":[{\"name\":\"taskHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskStatus\",\"inputs\":[{\"name\":\"taskHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumITaskMailboxTypes.TaskStatus\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_feeSplit\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"_feeSplitCollector\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isExecutorOperatorSetRegistered\",\"inputs\":[{\"name\":\"operatorSetKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"isRegistered\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"refundFee\",\"inputs\":[{\"name\":\"taskHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerExecutorOperatorSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"isRegistered\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setExecutorOperatorSetTaskConfig\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structITaskMailboxTypes.ExecutorOperatorSetTaskConfig\",\"components\":[{\"name\":\"taskHook\",\"type\":\"address\",\"internalType\":\"contractIAVSTaskHook\"},{\"name\":\"taskSLA\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"curveType\",\"type\":\"uint8\",\"internalType\":\"enumIKeyRegistrarTypes.CurveType\"},{\"name\":\"feeCollector\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"consensus\",\"type\":\"tuple\",\"internalType\":\"structITaskMailboxTypes.Consensus\",\"components\":[{\"name\":\"consensusType\",\"type\":\"uint8\",\"internalType\":\"enumITaskMailboxTypes.ConsensusType\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"taskMetadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFeeSplit\",\"inputs\":[{\"name\":\"_feeSplit\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFeeSplitCollector\",\"inputs\":[{\"name\":\"_feeSplitCollector\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitResult\",\"inputs\":[{\"name\":\"taskHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"executorCert\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"result\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"ExecutorOperatorSetRegistered\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"executorOperatorSetId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"isRegistered\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExecutorOperatorSetTaskConfigSet\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"executorOperatorSetId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structITaskMailboxTypes.ExecutorOperatorSetTaskConfig\",\"components\":[{\"name\":\"taskHook\",\"type\":\"address\",\"internalType\":\"contractIAVSTaskHook\"},{\"name\":\"taskSLA\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"curveType\",\"type\":\"uint8\",\"internalType\":\"enumIKeyRegistrarTypes.CurveType\"},{\"name\":\"feeCollector\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"consensus\",\"type\":\"tuple\",\"internalType\":\"structITaskMailboxTypes.Consensus\",\"components\":[{\"name\":\"consensusType\",\"type\":\"uint8\",\"internalType\":\"enumITaskMailboxTypes.ConsensusType\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"taskMetadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeRefunded\",\"inputs\":[{\"name\":\"refundCollector\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"taskHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"avsFee\",\"type\":\"uint96\",\"indexed\":false,\"internalType\":\"uint96\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeSplitCollectorSet\",\"inputs\":[{\"name\":\"feeSplitCollector\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeSplitSet\",\"inputs\":[{\"name\":\"feeSplit\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskCreated\",\"inputs\":[{\"name\":\"creator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"taskHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"executorOperatorSetId\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"refundCollector\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"avsFee\",\"type\":\"uint96\",\"indexed\":false,\"internalType\":\"uint96\"},{\"name\":\"taskDeadline\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskVerified\",\"inputs\":[{\"name\":\"aggregator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"taskHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"executorOperatorSetId\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"executorCert\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"result\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CertificateVerificationFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EmptyCertificateSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExecutorOperatorSetNotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExecutorOperatorSetTaskConfigNotSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FeeAlreadyRefunded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidConsensusType\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidConsensusValue\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCurveType\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidFeeReceiver\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidFeeSplit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperatorSetOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTaskCreator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTaskStatus\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"enumITaskMailboxTypes.TaskStatus\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"enumITaskMailboxTypes.TaskStatus\"}]},{\"type\":\"error\",\"name\":\"OnlyRefundCollector\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PayloadIsEmpty\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"TaskSLAIsZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TimestampAtCreation\",\"inputs\":[]}]",
	Bin: "0x60e060405234801561000f575f5ffd5b506040516158f83803806158f883398101604081905261002e9161018c565b6001600160a01b03808416608052821660a0528061004b8161005f565b60c052506100576100a5565b5050506102b8565b5f5f829050601f81511115610092578260405163305a27a960e01b8152600401610089919061025d565b60405180910390fd5b805161009d82610292565b179392505050565b5f54610100900460ff161561010c5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b6064820152608401610089565b5f5460ff9081161461015b575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b80516001600160a01b0381168114610173575f5ffd5b919050565b634e487b7160e01b5f52604160045260245ffd5b5f5f5f6060848603121561019e575f5ffd5b6101a78461015d565b92506101b56020850161015d565b60408501519092506001600160401b038111156101d0575f5ffd5b8401601f810186136101e0575f5ffd5b80516001600160401b038111156101f9576101f9610178565b604051601f8201601f19908116603f011681016001600160401b038111828210171561022757610227610178565b60405281815282820160200188101561023e575f5ffd5b8160208401602083015e5f602083830101528093505050509250925092565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b805160208083015191908110156102b2575f198160200360031b1b821691505b50919050565b60805160a05160c0516155fa6102fe5f395f611b8301525f81816102bb015281816134c8015261397901525f81816103d501528181613488015261388c01526155fa5ff3fe608060405234801561000f575f5ffd5b5060043610610187575f3560e01c80636373ea69116100d9578063a5fabc8111610093578063f2fde38b1161006e578063f2fde38b146103aa578063f741e81a146103bd578063f7424fc9146103d0578063fa2c0b37146103f7575f5ffd5b8063a5fabc8114610378578063b86941661461038b578063eda0be691461039e575f5ffd5b80636373ea69146102f8578063678fbdb3146103195780636bf6fad51461032c578063708c0db91461034c578063715018a61461035f5780638da5cb5b14610367575f5ffd5b80632bf6cc79116101445780634ad52e021161011f5780634ad52e021461029657806354743ad2146102b657806354fd4d50146102dd57806362fee037146102e5575f5ffd5b80632bf6cc791461024e578063468c07a01461026e57806349acd88414610283575f5ffd5b806302a744801461018b5780631270a892146101bb5780631a20c505146101db5780631ae370eb146101f45780631c7edb17146102075780631fb66f5d1461022d575b5f5ffd5b609b546201000090046001600160a01b03165b6040516001600160a01b0390911681526020015b60405180910390f35b6101ce6101c9366004613f5c565b610429565b6040516101b29190614011565b609b5461019e906201000090046001600160a01b031681565b6101ce61020236600461426e565b610452565b61021a61021536600461434e565b610465565b6040516101b297969594939291906143bd565b61024061023b366004614492565b61060c565b6040519081526020016101b2565b61026161025c36600461434e565b610ea9565b6040516101b29190614517565b61028161027c366004614534565b61133d565b005b61028161029136600461455c565b611351565b6102a96102a436600461434e565b61160c565b6040516101b29190614638565b61019e7f000000000000000000000000000000000000000000000000000000000000000081565b6101ce611b7c565b6101ce6102f336600461434e565b611bac565b609b546103069061ffff1681565b60405161ffff90911681526020016101b2565b61028161032736600461476f565b612091565b61033f61033a36600461478a565b6120a2565b6040516101b291906147a4565b61028161035a3660046147b6565b6122b6565b6102816123e7565b6033546001600160a01b031661019e565b6102816103863660046147fe565b6123fa565b61028161039936600461434e565b612cbc565b609b5461ffff16610306565b6102816103b836600461476f565b612eee565b6102816103cb3660046148f0565b612f64565b61019e7f000000000000000000000000000000000000000000000000000000000000000081565b61041961040536600461434e565b60996020525f908152604090205460ff1681565b60405190151581526020016101b2565b60608160405160200161043c9190614a1b565b6040516020818303038152906040529050919050565b60608160405160200161043c9190614b9f565b609a6020525f9081526040908190208054600180830154600284015485518087019096526003850180546001600160a01b03808716986001600160601b03600160a01b9889900416988683169860ff970487169795909216959194909392849216908111156104d6576104d6614365565b60018111156104e7576104e7614365565b81526020016001820180546104fb90614bb1565b80601f016020809104026020016040519081016040528092919081815260200182805461052790614bb1565b80156105725780601f1061054957610100808354040283529160200191610572565b820191905f5260205f20905b81548152906001019060200180831161055557829003601f168201915b5050505050815250509080600501805461058b90614bb1565b80601f01602080910402602001604051908101604052809291908181526020018280546105b790614bb1565b80156106025780601f106105d957610100808354040283529160200191610602565b820191905f5260205f20905b8154815290600101906020018083116105e557829003601f168201915b5050505050905087565b5f6106156131a6565b5f8260400151511161063a57604051636b1a1b6960e11b815260040160405180910390fd5b60995f61064a84602001516131ff565b815260208101919091526040015f205460ff1661067a5760405163c292b29760e01b815260040160405180910390fd5b5f609a5f61068b85602001516131ff565b815260208082019290925260409081015f20815160e08101835281546001600160a01b038082168352600160a01b918290046001600160601b03169583019590955260018301549485169382019390935292909160608401910460ff1660028111156106f9576106f9614365565b600281111561070a5761070a614365565b815260028201546001600160a01b0316602082015260408051808201825260038401805492909301929091829060ff16600181111561074b5761074b614365565b600181111561075c5761075c614365565b815260200160018201805461077090614bb1565b80601f016020809104026020016040519081016040528092919081815260200182805461079c90614bb1565b80156107e75780601f106107be576101008083540402835291602001916107e7565b820191905f5260205f20905b8154815290600101906020018083116107ca57829003601f168201915b505050505081525050815260200160058201805461080490614bb1565b80601f016020809104026020016040519081016040528092919081815260200182805461083090614bb1565b801561087b5780601f106108525761010080835404028352916020019161087b565b820191905f5260205f20905b81548152906001019060200180831161085e57829003601f168201915b50505050508152505090505f600281111561089857610898614365565b816060015160028111156108ae576108ae614365565b141580156108c5575080516001600160a01b031615155b80156108dd57505f81602001516001600160601b0316115b80156108ff57505f60a08201515160018111156108fc576108fc614365565b14155b61091c576040516314b0a41d60e11b815260040160405180910390fd5b8051604051630a3fc61360e31b81526001600160a01b03909116906351fe30989061094d9033908790600401614c3d565b5f6040518083038186803b158015610963575f5ffd5b505afa158015610975573d5f5f3e3d5ffd5b50508251604051637036693f60e11b81525f93506001600160a01b03909116915063e06cd27e906109aa908790600401614c60565b602060405180830381865afa1580156109c5573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109e99190614c72565b90505f609754304687604051602001610a059493929190614c8d565b6040516020818303038152906040528051906020012090506097546001610a2c9190614cd7565b609755604080516101a0810190915233815260208101610a4b42613268565b6001600160601b03908116825260208881018051516001600160a01b03908116838601529287166040850152895190921660608401529051015163ffffffff166080820152609b5461ffff1660a082015260c001600181525f6020808301829052604080840188905289810151606080860191909152815180840183528481526080808701919091528251808501845285815260a09687015287855260988452938290208651938701516001600160601b03908116600160a01b9081026001600160a01b039687161783559388015192880151168302918416919091176001820155928501516002840180549587015160c088015161ffff16600160c01b0261ffff60c01b1963ffffffff9092169094026001600160c01b03199097169290941691909117949094179182168117845560e085015192939160ff60d01b1990911662ffffff60c01b1990911617600160d01b836003811115610baf57610baf614365565b021790555061010082015160028083018054921515600160d81b0260ff60d81b1990931692909217909155610120830151805160208201516001600160601b0316600160a01b9081026001600160a01b0392831617600386019081556040840151600487018054919094166001600160a01b031982168117855560608601519596929594936001600160a81b031990921617918490811115610c5357610c53614365565b021790555060808201516002820180546001600160a01b0319166001600160a01b0390921691909117905560a08201518051600383018054909190829060ff191660018381811115610ca757610ca7614365565b021790555060208201516001820190610cc09082614d35565b50505060c08201516005820190610cd79082614d35565b5050506101408201516009820190610cef9082614d35565b50610160820151600a820190610d059082614d35565b50610180820151600b820190610d1b9082614d35565b50505060408301516001600160a01b031615801590610d4257505f826001600160601b0316115b15610dbd5760808301516001600160a01b0316610d7257604051633480121760e21b815260040160405180910390fd5b84516001600160a01b0316610d9a57604051633480121760e21b815260040160405180910390fd5b6040830151610dbd906001600160a01b031633306001600160601b0386166132d3565b8251604051629c5c4560e41b8152600481018390526001600160a01b03909116906309c5c450906024015f604051808303815f87803b158015610dfe575f5ffd5b505af1158015610e10573d5f5f3e3d5ffd5b5050505084602001515f01516001600160a01b031681336001600160a01b03167f4a09af06a0e08fd1c053a8b400de7833019c88066be8a2d3b3b17174a74fe317886020015160200151895f01518789602001516001600160601b031642610e789190614cd7565b8c60400151604051610e8e959493929190614def565b60405180910390a492505050610ea46001606555565b919050565b5f81815260986020908152604080832081516101a08101835281546001600160a01b038082168352600160a01b918290046001600160601b03908116968401969096526001840154808216958401959095529381900490941660608201526002820154928316608082015292820463ffffffff1660a0840152600160c01b820461ffff1660c084015283929160e0830190600160d01b900460ff166003811115610f5557610f55614365565b6003811115610f6657610f66614365565b815260028281015460ff600160d81b909104811615156020808501919091526040805160e0810182526003870180546001600160a01b0380821684526001600160601b03600160a01b92839004169584019590955260048901549485168385015292909601959094909360608601939290920490911690811115610fec57610fec614365565b6002811115610ffd57610ffd614365565b815260028201546001600160a01b0316602082015260408051808201825260038401805492909301929091829060ff16600181111561103e5761103e614365565b600181111561104f5761104f614365565b815260200160018201805461106390614bb1565b80601f016020809104026020016040519081016040528092919081815260200182805461108f90614bb1565b80156110da5780601f106110b1576101008083540402835291602001916110da565b820191905f5260205f20905b8154815290600101906020018083116110bd57829003601f168201915b50505050508152505081526020016005820180546110f790614bb1565b80601f016020809104026020016040519081016040528092919081815260200182805461112390614bb1565b801561116e5780601f106111455761010080835404028352916020019161116e565b820191905f5260205f20905b81548152906001019060200180831161115157829003601f168201915b505050505081525050815260200160098201805461118b90614bb1565b80601f01602080910402602001604051908101604052809291908181526020018280546111b790614bb1565b80156112025780601f106111d957610100808354040283529160200191611202565b820191905f5260205f20905b8154815290600101906020018083116111e557829003601f168201915b50505050508152602001600a8201805461121b90614bb1565b80601f016020809104026020016040519081016040528092919081815260200182805461124790614bb1565b80156112925780601f1061126957610100808354040283529160200191611292565b820191905f5260205f20905b81548152906001019060200180831161127557829003601f168201915b50505050508152602001600b820180546112ab90614bb1565b80601f01602080910402602001604051908101604052809291908181526020018280546112d790614bb1565b80156113225780601f106112f957610100808354040283529160200191611322565b820191905f5260205f20905b81548152906001019060200180831161130557829003601f168201915b505050505081525050905061133681613345565b9392505050565b6113456133a0565b61134e816133fa565b50565b5f609a5f61135e856131ff565b815260208082019290925260409081015f20815160e08101835281546001600160a01b038082168352600160a01b918290046001600160601b03169583019590955260018301549485169382019390935292909160608401910460ff1660028111156113cc576113cc614365565b60028111156113dd576113dd614365565b815260028201546001600160a01b0316602082015260408051808201825260038401805492909301929091829060ff16600181111561141e5761141e614365565b600181111561142f5761142f614365565b815260200160018201805461144390614bb1565b80601f016020809104026020016040519081016040528092919081815260200182805461146f90614bb1565b80156114ba5780601f10611491576101008083540402835291602001916114ba565b820191905f5260205f20905b81548152906001019060200180831161149d57829003601f168201915b50505050508152505081526020016005820180546114d790614bb1565b80601f016020809104026020016040519081016040528092919081815260200182805461150390614bb1565b801561154e5780601f106115255761010080835404028352916020019161154e565b820191905f5260205f20905b81548152906001019060200180831161153157829003601f168201915b50505050508152505090505f600281111561156b5761156b614365565b8160600151600281111561158157611581614365565b14158015611598575080516001600160a01b031615155b80156115b057505f81602001516001600160601b0316115b80156115d257505f60a08201515160018111156115cf576115cf614365565b14155b6115ef576040516314b0a41d60e11b815260040160405180910390fd5b6115fd83826060015161346b565b6116078383613599565b505050565b611614613d43565b5f82815260986020908152604080832081516101a08101835281546001600160a01b038082168352600160a01b918290046001600160601b03908116968401969096526001840154808216958401959095529381900490941660608201526002820154928316608082015292820463ffffffff1660a0840152600160c01b820461ffff1660c08401529060e0830190600160d01b900460ff1660038111156116be576116be614365565b60038111156116cf576116cf614365565b815260028281015460ff600160d81b909104811615156020808501919091526040805160e0810182526003870180546001600160a01b0380821684526001600160601b03600160a01b9283900416958401959095526004890154948516838501529290960195909490936060860193929092049091169081111561175557611755614365565b600281111561176657611766614365565b815260028201546001600160a01b0316602082015260408051808201825260038401805492909301929091829060ff1660018111156117a7576117a7614365565b60018111156117b8576117b8614365565b81526020016001820180546117cc90614bb1565b80601f01602080910402602001604051908101604052809291908181526020018280546117f890614bb1565b80156118435780601f1061181a57610100808354040283529160200191611843565b820191905f5260205f20905b81548152906001019060200180831161182657829003601f168201915b505050505081525050815260200160058201805461186090614bb1565b80601f016020809104026020016040519081016040528092919081815260200182805461188c90614bb1565b80156118d75780601f106118ae576101008083540402835291602001916118d7565b820191905f5260205f20905b8154815290600101906020018083116118ba57829003601f168201915b50505050508152505081526020016009820180546118f490614bb1565b80601f016020809104026020016040519081016040528092919081815260200182805461192090614bb1565b801561196b5780601f106119425761010080835404028352916020019161196b565b820191905f5260205f20905b81548152906001019060200180831161194e57829003601f168201915b50505050508152602001600a8201805461198490614bb1565b80601f01602080910402602001604051908101604052809291908181526020018280546119b090614bb1565b80156119fb5780601f106119d2576101008083540402835291602001916119fb565b820191905f5260205f20905b8154815290600101906020018083116119de57829003601f168201915b50505050508152602001600b82018054611a1490614bb1565b80601f0160208091040260200160405190810160405280929190818152602001828054611a4090614bb1565b8015611a8b5780601f10611a6257610100808354040283529160200191611a8b565b820191905f5260205f20905b815481529060010190602001808311611a6e57829003601f168201915b5050505050815250509050604051806101a00160405280825f01516001600160a01b0316815260200182602001516001600160601b0316815260200182604001516001600160a01b0316815260200182606001516001600160601b0316815260200182608001516001600160a01b031681526020018260a0015163ffffffff1681526020018260c0015161ffff168152602001611b2783613345565b6003811115611b3857611b38614365565b815260200182610100015115158152602001826101200151815260200182610140015181526020018261016001518152602001826101800151815250915050919050565b6060611ba77f000000000000000000000000000000000000000000000000000000000000000061362a565b905090565b5f81815260986020908152604080832081516101a08101835281546001600160a01b038082168352600160a01b918290046001600160601b0390811696840196909652600184015480821695840195909552938190049094166060808301919091526002830154938416608083015293830463ffffffff1660a0820152600160c01b830461ffff1660c08201529293929160e0830190600160d01b900460ff166003811115611c5d57611c5d614365565b6003811115611c6e57611c6e614365565b815260028281015460ff600160d81b909104811615156020808501919091526040805160e0810182526003870180546001600160a01b0380821684526001600160601b03600160a01b92839004169584019590955260048901549485168385015292909601959094909360608601939290920490911690811115611cf457611cf4614365565b6002811115611d0557611d05614365565b815260028201546001600160a01b0316602082015260408051808201825260038401805492909301929091829060ff166001811115611d4657611d46614365565b6001811115611d5757611d57614365565b8152602001600182018054611d6b90614bb1565b80601f0160208091040260200160405190810160405280929190818152602001828054611d9790614bb1565b8015611de25780601f10611db957610100808354040283529160200191611de2565b820191905f5260205f20905b815481529060010190602001808311611dc557829003601f168201915b5050505050815250508152602001600582018054611dff90614bb1565b80601f0160208091040260200160405190810160405280929190818152602001828054611e2b90614bb1565b8015611e765780601f10611e4d57610100808354040283529160200191611e76565b820191905f5260205f20905b815481529060010190602001808311611e5957829003601f168201915b5050505050815250508152602001600982018054611e9390614bb1565b80601f0160208091040260200160405190810160405280929190818152602001828054611ebf90614bb1565b8015611f0a5780601f10611ee157610100808354040283529160200191611f0a565b820191905f5260205f20905b815481529060010190602001808311611eed57829003601f168201915b50505050508152602001600a82018054611f2390614bb1565b80601f0160208091040260200160405190810160405280929190818152602001828054611f4f90614bb1565b8015611f9a5780601f10611f7157610100808354040283529160200191611f9a565b820191905f5260205f20905b815481529060010190602001808311611f7d57829003601f168201915b50505050508152602001600b82018054611fb390614bb1565b80601f0160208091040260200160405190810160405280929190818152602001828054611fdf90614bb1565b801561202a5780601f106120015761010080835404028352916020019161202a565b820191905f5260205f20905b81548152906001019060200180831161200d57829003601f168201915b50505050508152505090505f61203f82613345565b9050600281600381111561205557612055614365565b14600282909161208357604051634091b18960e11b815260040161207a929190614e30565b60405180910390fd5b505050610180015192915050565b6120996133a0565b61134e81613667565b6120aa613daf565b609a5f6120b6846131ff565b815260208082019290925260409081015f20815160e08101835281546001600160a01b038082168352600160a01b918290046001600160601b03169583019590955260018301549485169382019390935292909160608401910460ff16600281111561212457612124614365565b600281111561213557612135614365565b815260028201546001600160a01b0316602082015260408051808201825260038401805492909301929091829060ff16600181111561217657612176614365565b600181111561218757612187614365565b815260200160018201805461219b90614bb1565b80601f01602080910402602001604051908101604052809291908181526020018280546121c790614bb1565b80156122125780601f106121e957610100808354040283529160200191612212565b820191905f5260205f20905b8154815290600101906020018083116121f557829003601f168201915b505050505081525050815260200160058201805461222f90614bb1565b80601f016020809104026020016040519081016040528092919081815260200182805461225b90614bb1565b80156122a65780601f1061227d576101008083540402835291602001916122a6565b820191905f5260205f20905b81548152906001019060200180831161228957829003601f168201915b5050505050815250509050919050565b5f54610100900460ff16158080156122d457505f54600160ff909116105b806122ed5750303b1580156122ed57505f5460ff166001145b6123505760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161207a565b5f805460ff191660011790558015612371575f805461ff0019166101001790555b6123796136e1565b61238161370f565b61238a8461373d565b612393836133fa565b61239c82613667565b80156123e1575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050565b6123ef6133a0565b6123f85f61373d565b565b6124026131a6565b5f83815260986020908152604080832081516101a08101835281546001600160a01b038082168352600160a01b918290046001600160601b03908116968401969096526001840154808216958401959095529381900490941660608201526002820154928316608082015292820463ffffffff1660a0840152600160c01b820461ffff1660c0840152929161288e91849060e0830190600160d01b900460ff1660038111156124b3576124b3614365565b60038111156124c4576124c4614365565b815260028281015460ff600160d81b909104811615156020808501919091526040805160e0810182526003870180546001600160a01b0380821684526001600160601b03600160a01b9283900416958401959095526004890154948516838501529290960195909490936060860193929092049091169081111561254a5761254a614365565b600281111561255b5761255b614365565b815260028201546001600160a01b0316602082015260408051808201825260038401805492909301929091829060ff16600181111561259c5761259c614365565b60018111156125ad576125ad614365565b81526020016001820180546125c190614bb1565b80601f01602080910402602001604051908101604052809291908181526020018280546125ed90614bb1565b80156126385780601f1061260f57610100808354040283529160200191612638565b820191905f5260205f20905b81548152906001019060200180831161261b57829003601f168201915b505050505081525050815260200160058201805461265590614bb1565b80601f016020809104026020016040519081016040528092919081815260200182805461268190614bb1565b80156126cc5780601f106126a3576101008083540402835291602001916126cc565b820191905f5260205f20905b8154815290600101906020018083116126af57829003601f168201915b50505050508152505081526020016009820180546126e990614bb1565b80601f016020809104026020016040519081016040528092919081815260200182805461271590614bb1565b80156127605780601f1061273757610100808354040283529160200191612760565b820191905f5260205f20905b81548152906001019060200180831161274357829003601f168201915b50505050508152602001600a8201805461277990614bb1565b80601f01602080910402602001604051908101604052809291908181526020018280546127a590614bb1565b80156127f05780601f106127c7576101008083540402835291602001916127f0565b820191905f5260205f20905b8154815290600101906020018083116127d357829003601f168201915b50505050508152602001600b8201805461280990614bb1565b80601f016020809104026020016040519081016040528092919081815260200182805461283590614bb1565b80156128805780601f1061285757610100808354040283529160200191612880565b820191905f5260205f20905b81548152906001019060200180831161286357829003601f168201915b505050505081525050613345565b905060018160038111156128a4576128a4614365565b1460018290916128c957604051634091b18960e11b815260040161207a929190614e30565b50508154600160a01b90046001600160601b031642116128fc5760405163015a4b7560e51b815260040160405180910390fd5b600382015460405163ba33565d60e01b81526001600160a01b039091169063ba33565d90612934903390899089908990600401614e4b565b5f6040518083038186803b15801561294a575f5ffd5b505afa15801561295c573d5f5f3e3d5ffd5b50506040805180820182526001808701546001600160a01b03168252600287015463ffffffff600160a01b91829004166020840152600488015484518086019095526006890180549497505f9650612a7a9560ff9390920483169491939092849216908111156129ce576129ce614365565b60018111156129df576129df614365565b81526020016001820180546129f390614bb1565b80601f0160208091040260200160405190810160405280929190818152602001828054612a1f90614bb1565b8015612a6a5780601f10612a4157610100808354040283529160200191612a6a565b820191905f5260205f20905b815481529060010190602001808311612a4d57829003601f168201915b505050505081525050848961378e565b905080612a9a5760405163efcb98e760e01b815260040160405180910390fd5b60028401805460ff60d01b1916600160d11b179055600a8401612abd8782614d35565b50600b8401612acc8682614d35565b5060048401546001600160a01b031615801590612afc57506001840154600160a01b90046001600160601b031615155b15612bea57600284015460018501545f91612b4a9161271091612b3b91600160c01b90910461ffff1690600160a01b90046001600160601b0316614e83565b612b459190614e9a565b613268565b90506001600160601b03811615612b8a57609b546004860154612b8a916001600160a01b039182169162010000909104166001600160601b038416613a25565b60018501545f90612bac908390600160a01b90046001600160601b0316614eb9565b90506001600160601b03811615612be75760058601546004870154612be7916001600160a01b0391821691166001600160601b038416613a25565b50505b600384015460405163db6ecf6760e01b8152600481018990526001600160a01b039091169063db6ecf67906024015f604051808303815f87803b158015612c2f575f5ffd5b505af1158015612c41573d5f5f3e3d5ffd5b505050600185015460028601546040516001600160a01b039092169250899133917f659f23b2e7edf490e5fd6561c5148691ed0375ed7ddd3ab1bcfcfdbec4f209a991612ca69163ffffffff600160a01b9091041690600a8b0190600b8c0190614f57565b60405180910390a4505050506116076001606555565b612cc46131a6565b5f81815260986020526040902060028101546001600160a01b03163314612cfe576040516370f43cb760e01b815260040160405180910390fd5b6002810154600160d81b900460ff1615612d2b57604051633e3d786960e01b815260040160405180910390fd5b604080516101a08101825282546001600160a01b038082168352600160a01b918290046001600160601b03908116602085015260018601548083169585019590955293829004909316606083015260028401549283166080830152820463ffffffff1660a0820152600160c01b820461ffff1660c08201525f91612dcc9190849060e0830190600160d01b900460ff1660038111156124b3576124b3614365565b90506003816003811115612de257612de2614365565b146003829091612e0757604051634091b18960e11b815260040161207a929190614e30565b505060028201805460ff60d81b1916600160d81b17905560048201546001600160a01b031615801590612e4d57506001820154600160a01b90046001600160601b031615155b15612e8957600282015460018301546004840154612e89926001600160a01b0391821692911690600160a01b90046001600160601b0316613a25565b60028201546001830154604051600160a01b9091046001600160601b0316815284916001600160a01b0316907fe3ed40d31808582f7a92a30beacc0ec788d5091407ec6c10c1b999b3f317aea39060200160405180910390a3505061134e6001606555565b612ef66133a0565b6001600160a01b038116612f5b5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161207a565b61134e8161373d565b5f81606001516002811115612f7b57612f7b614365565b03612f995760405163fdea7c0960e01b815260040160405180910390fd5b80516001600160a01b0316612fc157604051630863a45360e11b815260040160405180910390fd5b5f81602001516001600160601b031611612fee5760405163568added60e11b815260040160405180910390fd5b612ffb8160a00151613a55565b61300982826060015161346b565b80609a5f613016856131ff565b815260208082019290925260409081015f208351928401516001600160601b0316600160a01b9081026001600160a01b0394851617825591840151600182018054919094166001600160a01b03198216811785556060860151929492936001600160a81b0319909216179083600281111561309357613093614365565b021790555060808201516002820180546001600160a01b0319166001600160a01b0390921691909117905560a08201518051600383018054909190829060ff1916600183818111156130e7576130e7614365565b0217905550602082015160018201906131009082614d35565b50505060c082015160058201906131179082614d35565b50905050816020015163ffffffff16825f01516001600160a01b0316336001600160a01b03167f7cd76abd4025a20959a1b20f7c1536e3894a0735cd8de0215dde803ddea7f2d28460405161316c91906147a4565b60405180910390a460995f613180846131ff565b815260208101919091526040015f205460ff166131a2576131a2826001613599565b5050565b6002606554036131f85760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015260640161207a565b6002606555565b5f815f0151826020015163ffffffff1660405160200161324a92919060609290921b6bffffffffffffffffffffffff1916825260a01b6001600160a01b031916601482015260200190565b60405160208183030381529060405261326290614f87565b92915050565b5f6001600160601b038211156132cf5760405162461bcd60e51b815260206004820152602660248201527f53616665436173743a2076616c756520646f65736e27742066697420696e203960448201526536206269747360d01b606482015260840161207a565b5090565b6040516001600160a01b03808516602483015283166044820152606481018290526123e19085906323b872dd60e01b906084015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b031990931692909217909152613ad7565b6001606555565b5f60018260e00151600381111561335e5761335e614365565b14801561338b575081610120015160200151826020015161337f9190614faa565b6001600160601b031642115b1561339857506003919050565b5060e0015190565b6033546001600160a01b031633146123f85760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161207a565b6127108161ffff16111561342157604051630601f69760e01b815260040160405180910390fd5b609b805461ffff191661ffff83169081179091556040519081527f886b2cfcb151fd8b19ed902cc88f4a06dd9fe351a4a9ab93f33fe84abc157edf9060200160405180910390a150565b5f600282600281111561348057613480614365565b036134ac57507f0000000000000000000000000000000000000000000000000000000000000000613505565b60018260028111156134c0576134c0614365565b036134ec57507f0000000000000000000000000000000000000000000000000000000000000000613505565b60405163fdea7c0960e01b815260040160405180910390fd5b6040516304240c4960e51b815233906001600160a01b03831690638481892090613533908790600401614fc9565b602060405180830381865afa15801561354e573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906135729190614fef565b6001600160a01b031614611607576040516342ecfee960e11b815260040160405180910390fd5b8060995f6135a6856131ff565b81526020019081526020015f205f6101000a81548160ff021916908315150217905550816020015163ffffffff16825f01516001600160a01b0316336001600160a01b03167f48b63f21a1eb9dd6880e196de6d7db3fbd0c282b74f1298dcb4cf53472298f398460405161361e911515815260200190565b60405180910390a45050565b60605f61363683613baa565b6040805160208082528183019092529192505f91906020820181803683375050509182525060208101929092525090565b6001600160a01b03811661368e57604051630863a45360e11b815260040160405180910390fd5b609b805462010000600160b01b031916620100006001600160a01b038416908102919091179091556040517f262aa27c244f6f0088cb3092548a0adcaddedf459070a9ccab2dc6a07abe701d905f90a250565b5f54610100900460ff166137075760405162461bcd60e51b815260040161207a9061500a565b6123f8613bd1565b5f54610100900460ff166137355760405162461bcd60e51b815260040161207a9061500a565b6123f8613c00565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b5f6001845160018111156137a4576137a4614365565b03613a04575f84602001518060200190518101906137c29190615055565b6040805160018082528183019092529192505f91906020808301908036833701905050905081815f815181106137fa576137fa615070565b61ffff90921660209283029190910190910152600287600281111561382157613821614365565b0361390d575f8480602001905181019061383b91906152e5565b60408101515190915015801590613859575060408101516020015115155b61387657604051637a8a1dbd60e11b815260040160405180910390fd5b604051625f5e5d60e21b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063017d7974906138c5908990859087906004016153f5565b6020604051808303815f875af11580156138e1573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906139059190615439565b9350506139fd565b600187600281111561392157613921614365565b036134ec575f8480602001905181019061393b9190615454565b90505f8160400151511161396257604051637a8a1dbd60e11b815260040160405180910390fd5b604051630606d12160e51b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063c0da2420906139b2908990859087906004016154cd565b5f60405180830381865afa1580156139cc573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526139f391908101906154ff565b5093506139fd9050565b5050613a1d565b6040516347d3772160e11b815260040160405180910390fd5b949350505050565b6040516001600160a01b03831660248201526044810182905261160790849063a9059cbb60e01b90606401613307565b600181516001811115613a6a57613a6a614365565b03613a0457806020015151602014613a9557604051631501e04760e21b815260040160405180910390fd5b5f8160200151806020019051810190613aae9190615055565b90506127108161ffff1611156131a257604051631501e04760e21b815260040160405180910390fd5b5f613b2b826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316613c269092919063ffffffff16565b905080515f1480613b4b575080806020019051810190613b4b9190615439565b6116075760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b606482015260840161207a565b5f60ff8216601f81111561326257604051632cd44ac360e21b815260040160405180910390fd5b5f54610100900460ff16613bf75760405162461bcd60e51b815260040161207a9061500a565b6123f83361373d565b5f54610100900460ff1661333e5760405162461bcd60e51b815260040161207a9061500a565b6060613a1d84845f85855f5f866001600160a01b03168587604051613c4b91906155ae565b5f6040518083038185875af1925050503d805f8114613c85576040519150601f19603f3d011682016040523d82523d5f602084013e613c8a565b606091505b5091509150613c9b87838387613ca6565b979650505050505050565b60608315613d145782515f03613d0d576001600160a01b0385163b613d0d5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161207a565b5081613a1d565b613a1d8383815115613d295781518083602001fd5b8060405162461bcd60e51b815260040161207a9190614011565b604080516101a0810182525f80825260208201819052918101829052606081018290526080810182905260a0810182905260c081018290529060e082019081525f6020820152604001613d94613daf565b81526020016060815260200160608152602001606081525090565b6040805160e0810182525f8082526020820181905291810182905290606082019081525f6020820152604001613de3613df0565b8152602001606081525090565b60408051808201909152805f613de3565b634e487b7160e01b5f52604160045260245ffd5b604051606081016001600160401b0381118282101715613e3757613e37613e01565b60405290565b604080519081016001600160401b0381118282101715613e3757613e37613e01565b60405160a081016001600160401b0381118282101715613e3757613e37613e01565b60405160e081016001600160401b0381118282101715613e3757613e37613e01565b604051601f8201601f191681016001600160401b0381118282101715613ecb57613ecb613e01565b604052919050565b63ffffffff8116811461134e575f5ffd5b5f6001600160401b03821115613efc57613efc613e01565b50601f01601f191660200190565b5f82601f830112613f19575f5ffd5b8135613f2c613f2782613ee4565b613ea3565b818152846020838601011115613f40575f5ffd5b816020850160208301375f918101602001919091529392505050565b5f60208284031215613f6c575f5ffd5b81356001600160401b03811115613f81575f5ffd5b820160608185031215613f92575f5ffd5b613f9a613e15565b8135613fa581613ed3565b81526020828101359082015260408201356001600160401b03811115613fc9575f5ffd5b613fd586828501613f0a565b604083015250949350505050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f6113366020830184613fe3565b5f60408284031215614033575f5ffd5b61403b613e3d565b823581526020928301359281019290925250919050565b5f82601f830112614061575f5ffd5b61406b6040613ea3565b80604084018581111561407c575f5ffd5b845b8181101561409657803584526020938401930161407e565b509095945050505050565b5f6001600160401b038211156140b9576140b9613e01565b5060051b60200190565b5f82601f8301126140d2575f5ffd5b81356140e0613f27826140a1565b8082825260208201915060208360051b860101925085831115614101575f5ffd5b602085015b838110156142645780356001600160401b03811115614123575f5ffd5b86016060818903601f19011215614138575f5ffd5b614140613e15565b602082013561414e81613ed3565b815260408201356001600160401b03811115614168575f5ffd5b6141778a602083860101613f0a565b60208301525060608201356001600160401b03811115614195575f5ffd5b6020818401019250506060828a0312156141ad575f5ffd5b6141b5613e3d565b6141bf8a84614023565b815260408301356001600160401b038111156141d9575f5ffd5b80840193505089601f8401126141ed575f5ffd5b82356141fb613f27826140a1565b8082825260208201915060208360051b87010192508c83111561421c575f5ffd5b6020860195505b8286101561423e578535825260209586019590910190614223565b806020850152505050806040830152508085525050602083019250602081019050614106565b5095945050505050565b5f6020828403121561427e575f5ffd5b81356001600160401b03811115614293575f5ffd5b82018084036101208112156142a6575f5ffd5b6142ae613e5f565b82356142b981613ed3565b8152602083810135908201526142d28660408501614023565b60408201526080607f19830112156142e8575f5ffd5b6142f0613e3d565b91506142ff8660808501614052565b825261430e8660c08501614052565b602083015281606082015261010083013591506001600160401b03821115614334575f5ffd5b614340868385016140c3565b608082015295945050505050565b5f6020828403121561435e575f5ffd5b5035919050565b634e487b7160e01b5f52602160045260245ffd5b6003811061438957614389614365565b9052565b5f8151600281106143a0576143a0614365565b80845250602082015160406020850152613a1d6040850182613fe3565b6001600160a01b0388811682526001600160601b0388166020830152861660408201526143ed6060820186614379565b6001600160a01b038416608082015260e060a082018190525f906144139083018561438d565b82810360c08401526144258185613fe3565b9a9950505050505050505050565b6001600160a01b038116811461134e575f5ffd5b8035610ea481614433565b5f60408284031215614462575f5ffd5b61446a613e3d565b9050813561447781614433565b8152602082013561448781613ed3565b602082015292915050565b5f602082840312156144a2575f5ffd5b81356001600160401b038111156144b7575f5ffd5b8201608081850312156144c8575f5ffd5b6144d0613e15565b81356144db81614433565b81526144ea8560208401614452565b602082015260608201356001600160401b03811115613fc9575f5ffd5b6004811061438957614389614365565b602081016132628284614507565b61ffff8116811461134e575f5ffd5b5f60208284031215614544575f5ffd5b813561133681614525565b801515811461134e575f5ffd5b5f5f6060838503121561456d575f5ffd5b6145778484614452565b915060408301356145878161454f565b809150509250929050565b80516001600160a01b031682526020808201516001600160601b0316908301526040808201515f916145ce908501826001600160a01b03169052565b5060608201516145e16060850182614379565b5060808201516145fc60808501826001600160a01b03169052565b5060a082015160e060a085015261461660e085018261438d565b905060c083015184820360c086015261462f8282613fe3565b95945050505050565b602081526146526020820183516001600160a01b03169052565b5f602083015161466d60408401826001600160601b03169052565b5060408301516001600160a01b03811660608401525060608301516001600160601b03811660808401525060808301516001600160a01b03811660a08401525060a083015163ffffffff811660c08401525060c083015161ffff811660e08401525060e08301516146e2610100840182614507565b50610100830151801515610120840152506101208301516101a06101408401526147106101c0840182614592565b9050610140840151601f198483030161016085015261472f8282613fe3565b915050610160840151601f198483030161018085015261474f8282613fe3565b915050610180840151601f19848303016101a085015261462f8282613fe3565b5f6020828403121561477f575f5ffd5b813561133681614433565b5f6040828403121561479a575f5ffd5b6113368383614452565b602081525f6113366020830184614592565b5f5f5f606084860312156147c8575f5ffd5b83356147d381614433565b925060208401356147e381614525565b915060408401356147f381614433565b809150509250925092565b5f5f5f60608486031215614810575f5ffd5b8335925060208401356001600160401b0381111561482c575f5ffd5b61483886828701613f0a565b92505060408401356001600160401b03811115614853575f5ffd5b61485f86828701613f0a565b9150509250925092565b6001600160601b038116811461134e575f5ffd5b8035610ea481614869565b803560038110610ea4575f5ffd5b5f604082840312156148a6575f5ffd5b6148ae613e3d565b90508135600281106148be575f5ffd5b815260208201356001600160401b038111156148d8575f5ffd5b6148e484828501613f0a565b60208301525092915050565b5f5f60608385031215614901575f5ffd5b61490b8484614452565b915060408301356001600160401b03811115614925575f5ffd5b830160e08186031215614936575f5ffd5b61493e613e81565b61494782614447565b81526149556020830161487d565b602082015261496660408301614447565b604082015261497760608301614888565b606082015261498860808301614447565b608082015260a08201356001600160401b038111156149a5575f5ffd5b6149b187828501614896565b60a08301525060c08201356001600160401b038111156149cf575f5ffd5b6149db87828501613f0a565b60c08301525080925050509250929050565b63ffffffff8151168252602081015160208301525f604082015160606040850152613a1d6060850182613fe3565b602081525f61133660208301846149ed565b805f5b60028110156123e1578151845260209384019390910190600101614a30565b5f610120830163ffffffff8351168452602083015160208501526040830151614a85604086018280518252602090810151910152565b506060830151614a99608086018251614a2d565b60200151614aaa60c0860182614a2d565b506080830151610120610100860152818151808452610140870191506101408160051b88010193506020830192505f5b81811015614b935761013f19888603018352835163ffffffff8151168652602081015160606020880152614b116060880182613fe3565b905060408201519150868103604088015260608101614b3b82845180518252602090810151910152565b6020928301516060604084015280518083529301925f92608001905b80841015614b7a5784518252602082019150602085019450600184019350614b57565b5097505050602094850194939093019250600101614ada565b50929695505050505050565b602081525f6113366020830184614a4f565b600181811c90821680614bc557607f821691505b602082108103614be357634e487b7160e01b5f52602260045260245ffd5b50919050565b80516001600160a01b031682526020808201515f91614c239085018280516001600160a01b0316825260209081015163ffffffff16910152565b50604082015160806060850152613a1d6080850182613fe3565b6001600160a01b03831681526040602082018190525f90613a1d90830184614be9565b602081525f6113366020830184614be9565b5f60208284031215614c82575f5ffd5b815161133681614869565b84815260018060a01b0384166020820152826040820152608060608201525f614cb96080830184614be9565b9695505050505050565b634e487b7160e01b5f52601160045260245ffd5b8082018082111561326257613262614cc3565b601f82111561160757805f5260205f20601f840160051c81016020851015614d0f5750805b601f840160051c820191505b81811015614d2e575f8155600101614d1b565b5050505050565b81516001600160401b03811115614d4e57614d4e613e01565b614d6281614d5c8454614bb1565b84614cea565b6020601f821160018114614d94575f8315614d7d5750848201515b5f19600385901b1c1916600184901b178455614d2e565b5f84815260208120601f198516915b82811015614dc35787850151825560209485019460019092019101614da3565b5084821015614de057868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b63ffffffff8616815260018060a01b03851660208201526001600160601b038416604082015282606082015260a060808201525f613c9b60a0830184613fe3565b60408101614e3e8285614507565b6113366020830184614507565b60018060a01b0385168152836020820152608060408201525f614e716080830185613fe3565b8281036060840152613c9b8185613fe3565b808202811582820484141761326257613262614cc3565b5f82614eb457634e487b7160e01b5f52601260045260245ffd5b500490565b6001600160601b03828116828216039081111561326257613262614cc3565b5f8154614ee481614bb1565b808552600182168015614efe5760018114614f1a57614f4e565b60ff1983166020870152602082151560051b8701019350614f4e565b845f5260205f205f5b83811015614f455781546020828a010152600182019150602081019050614f23565b87016020019450505b50505092915050565b63ffffffff84168152606060208201525f614f756060830185614ed8565b8281036040840152614cb98185614ed8565b80516020808301519190811015614be3575f1960209190910360031b1b16919050565b6001600160601b03818116838216019081111561326257613262614cc3565b81516001600160a01b0316815260208083015163ffffffff169082015260408101613262565b5f60208284031215614fff575f5ffd5b815161133681614433565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b5f60208284031215615065575f5ffd5b815161133681614525565b634e487b7160e01b5f52603260045260245ffd5b5f60408284031215615094575f5ffd5b61509c613e3d565b825181526020928301519281019290925250919050565b5f82601f8301126150c2575f5ffd5b6150cc6040613ea3565b8060408401858111156150dd575f5ffd5b845b818110156140965780518452602093840193016150df565b5f82601f830112615106575f5ffd5b8151615114613f2782613ee4565b818152846020838601011115615128575f5ffd5b8160208501602083015e5f918101602001919091529392505050565b5f82601f830112615153575f5ffd5b8151615161613f27826140a1565b8082825260208201915060208360051b860101925085831115615182575f5ffd5b602085015b838110156142645780516001600160401b038111156151a4575f5ffd5b86016060818903601f190112156151b9575f5ffd5b6151c1613e15565b60208201516151cf81613ed3565b815260408201516001600160401b038111156151e9575f5ffd5b6151f88a6020838601016150f7565b60208301525060608201516001600160401b03811115615216575f5ffd5b6020818401019250506060828a03121561522e575f5ffd5b615236613e3d565b6152408a84615084565b815260408301516001600160401b0381111561525a575f5ffd5b80840193505089601f84011261526e575f5ffd5b825161527c613f27826140a1565b8082825260208201915060208360051b87010192508c83111561529d575f5ffd5b6020860195505b828610156152bf5785518252602095860195909101906152a4565b806020850152505050806040830152508085525050602083019250602081019050615187565b5f602082840312156152f5575f5ffd5b81516001600160401b0381111561530a575f5ffd5b820180840361012081121561531d575f5ffd5b615325613e5f565b825161533081613ed3565b8152602083810151908201526153498660408501615084565b60408201526080607f198301121561535f575f5ffd5b615367613e3d565b915061537686608085016150b3565b82526153858660c085016150b3565b602083015281606082015261010083015191506001600160401b038211156153ab575f5ffd5b61434086838501615144565b5f8151808452602084019350602083015f5b828110156153eb57815161ffff168652602095860195909101906001016153c9565b5093949350505050565b83516001600160a01b0316815260208085015163ffffffff1690820152608060408201525f6154276080830185614a4f565b8281036060840152614cb981856153b7565b5f60208284031215615449575f5ffd5b81516113368161454f565b5f60208284031215615464575f5ffd5b81516001600160401b03811115615479575f5ffd5b82016060818503121561548a575f5ffd5b615492613e15565b815161549d81613ed3565b81526020828101519082015260408201516001600160401b038111156154c1575f5ffd5b613fd5868285016150f7565b83516001600160a01b0316815260208085015163ffffffff1690820152608060408201525f61542760808301856149ed565b5f5f60408385031215615510575f5ffd5b825161551b8161454f565b60208401519092506001600160401b03811115615536575f5ffd5b8301601f81018513615546575f5ffd5b8051615554613f27826140a1565b8082825260208201915060208360051b850101925087831115615575575f5ffd5b6020840193505b828410156155a057835161558f81614433565b82526020938401939091019061557c565b809450505050509250929050565b5f82518060208501845e5f92019182525091905056fea2646970667358221220084d2c54e2947819ff9c81f7c2d549499d4abc859fb04936c2d83d18118a51f064736f6c634300081b0033",
}

// TaskMailboxABI is the input ABI used to generate the binding from.
// Deprecated: Use TaskMailboxMetaData.ABI instead.
var TaskMailboxABI = TaskMailboxMetaData.ABI

// TaskMailboxBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TaskMailboxMetaData.Bin instead.
var TaskMailboxBin = TaskMailboxMetaData.Bin

// DeployTaskMailbox deploys a new Ethereum contract, binding an instance of TaskMailbox to it.
func DeployTaskMailbox(auth *bind.TransactOpts, backend bind.ContractBackend, _bn254CertificateVerifier common.Address, _ecdsaCertificateVerifier common.Address, _version string) (common.Address, *types.Transaction, *TaskMailbox, error) {
	parsed, err := TaskMailboxMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TaskMailboxBin), backend, _bn254CertificateVerifier, _ecdsaCertificateVerifier, _version)
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

// ExecutorOperatorSetTaskConfigs is a free data retrieval call binding the contract method 0x1c7edb17.
//
// Solidity: function executorOperatorSetTaskConfigs(bytes32 operatorSetKey) view returns(address taskHook, uint96 taskSLA, address feeToken, uint8 curveType, address feeCollector, (uint8,bytes) consensus, bytes taskMetadata)
func (_TaskMailbox *TaskMailboxCaller) ExecutorOperatorSetTaskConfigs(opts *bind.CallOpts, operatorSetKey [32]byte) (struct {
	TaskHook     common.Address
	TaskSLA      *big.Int
	FeeToken     common.Address
	CurveType    uint8
	FeeCollector common.Address
	Consensus    ITaskMailboxTypesConsensus
	TaskMetadata []byte
}, error) {
	var out []interface{}
	err := _TaskMailbox.contract.Call(opts, &out, "executorOperatorSetTaskConfigs", operatorSetKey)

	outstruct := new(struct {
		TaskHook     common.Address
		TaskSLA      *big.Int
		FeeToken     common.Address
		CurveType    uint8
		FeeCollector common.Address
		Consensus    ITaskMailboxTypesConsensus
		TaskMetadata []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TaskHook = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.TaskSLA = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.FeeToken = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.CurveType = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.FeeCollector = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Consensus = *abi.ConvertType(out[5], new(ITaskMailboxTypesConsensus)).(*ITaskMailboxTypesConsensus)
	outstruct.TaskMetadata = *abi.ConvertType(out[6], new([]byte)).(*[]byte)

	return *outstruct, err

}

// ExecutorOperatorSetTaskConfigs is a free data retrieval call binding the contract method 0x1c7edb17.
//
// Solidity: function executorOperatorSetTaskConfigs(bytes32 operatorSetKey) view returns(address taskHook, uint96 taskSLA, address feeToken, uint8 curveType, address feeCollector, (uint8,bytes) consensus, bytes taskMetadata)
func (_TaskMailbox *TaskMailboxSession) ExecutorOperatorSetTaskConfigs(operatorSetKey [32]byte) (struct {
	TaskHook     common.Address
	TaskSLA      *big.Int
	FeeToken     common.Address
	CurveType    uint8
	FeeCollector common.Address
	Consensus    ITaskMailboxTypesConsensus
	TaskMetadata []byte
}, error) {
	return _TaskMailbox.Contract.ExecutorOperatorSetTaskConfigs(&_TaskMailbox.CallOpts, operatorSetKey)
}

// ExecutorOperatorSetTaskConfigs is a free data retrieval call binding the contract method 0x1c7edb17.
//
// Solidity: function executorOperatorSetTaskConfigs(bytes32 operatorSetKey) view returns(address taskHook, uint96 taskSLA, address feeToken, uint8 curveType, address feeCollector, (uint8,bytes) consensus, bytes taskMetadata)
func (_TaskMailbox *TaskMailboxCallerSession) ExecutorOperatorSetTaskConfigs(operatorSetKey [32]byte) (struct {
	TaskHook     common.Address
	TaskSLA      *big.Int
	FeeToken     common.Address
	CurveType    uint8
	FeeCollector common.Address
	Consensus    ITaskMailboxTypesConsensus
	TaskMetadata []byte
}, error) {
	return _TaskMailbox.Contract.ExecutorOperatorSetTaskConfigs(&_TaskMailbox.CallOpts, operatorSetKey)
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

// GetFeeSplit is a free data retrieval call binding the contract method 0xeda0be69.
//
// Solidity: function getFeeSplit() view returns(uint16)
func (_TaskMailbox *TaskMailboxCaller) GetFeeSplit(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _TaskMailbox.contract.Call(opts, &out, "getFeeSplit")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetFeeSplit is a free data retrieval call binding the contract method 0xeda0be69.
//
// Solidity: function getFeeSplit() view returns(uint16)
func (_TaskMailbox *TaskMailboxSession) GetFeeSplit() (uint16, error) {
	return _TaskMailbox.Contract.GetFeeSplit(&_TaskMailbox.CallOpts)
}

// GetFeeSplit is a free data retrieval call binding the contract method 0xeda0be69.
//
// Solidity: function getFeeSplit() view returns(uint16)
func (_TaskMailbox *TaskMailboxCallerSession) GetFeeSplit() (uint16, error) {
	return _TaskMailbox.Contract.GetFeeSplit(&_TaskMailbox.CallOpts)
}

// GetFeeSplitCollector is a free data retrieval call binding the contract method 0x02a74480.
//
// Solidity: function getFeeSplitCollector() view returns(address)
func (_TaskMailbox *TaskMailboxCaller) GetFeeSplitCollector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TaskMailbox.contract.Call(opts, &out, "getFeeSplitCollector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFeeSplitCollector is a free data retrieval call binding the contract method 0x02a74480.
//
// Solidity: function getFeeSplitCollector() view returns(address)
func (_TaskMailbox *TaskMailboxSession) GetFeeSplitCollector() (common.Address, error) {
	return _TaskMailbox.Contract.GetFeeSplitCollector(&_TaskMailbox.CallOpts)
}

// GetFeeSplitCollector is a free data retrieval call binding the contract method 0x02a74480.
//
// Solidity: function getFeeSplitCollector() view returns(address)
func (_TaskMailbox *TaskMailboxCallerSession) GetFeeSplitCollector() (common.Address, error) {
	return _TaskMailbox.Contract.GetFeeSplitCollector(&_TaskMailbox.CallOpts)
}

// GetTaskInfo is a free data retrieval call binding the contract method 0x4ad52e02.
//
// Solidity: function getTaskInfo(bytes32 taskHash) view returns((address,uint96,address,uint96,address,uint32,uint16,uint8,bool,(address,uint96,address,uint8,address,(uint8,bytes),bytes),bytes,bytes,bytes))
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
// Solidity: function getTaskInfo(bytes32 taskHash) view returns((address,uint96,address,uint96,address,uint32,uint16,uint8,bool,(address,uint96,address,uint8,address,(uint8,bytes),bytes),bytes,bytes,bytes))
func (_TaskMailbox *TaskMailboxSession) GetTaskInfo(taskHash [32]byte) (ITaskMailboxTypesTask, error) {
	return _TaskMailbox.Contract.GetTaskInfo(&_TaskMailbox.CallOpts, taskHash)
}

// GetTaskInfo is a free data retrieval call binding the contract method 0x4ad52e02.
//
// Solidity: function getTaskInfo(bytes32 taskHash) view returns((address,uint96,address,uint96,address,uint32,uint16,uint8,bool,(address,uint96,address,uint8,address,(uint8,bytes),bytes),bytes,bytes,bytes))
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
	Creator               common.Address
	TaskHash              [32]byte
	Avs                   common.Address
	ExecutorOperatorSetId uint32
	RefundCollector       common.Address
	AvsFee                *big.Int
	TaskDeadline          *big.Int
	Payload               []byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterTaskCreated is a free log retrieval operation binding the contract event 0x4a09af06a0e08fd1c053a8b400de7833019c88066be8a2d3b3b17174a74fe317.
//
// Solidity: event TaskCreated(address indexed creator, bytes32 indexed taskHash, address indexed avs, uint32 executorOperatorSetId, address refundCollector, uint96 avsFee, uint256 taskDeadline, bytes payload)
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

// WatchTaskCreated is a free log subscription operation binding the contract event 0x4a09af06a0e08fd1c053a8b400de7833019c88066be8a2d3b3b17174a74fe317.
//
// Solidity: event TaskCreated(address indexed creator, bytes32 indexed taskHash, address indexed avs, uint32 executorOperatorSetId, address refundCollector, uint96 avsFee, uint256 taskDeadline, bytes payload)
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

// ParseTaskCreated is a log parse operation binding the contract event 0x4a09af06a0e08fd1c053a8b400de7833019c88066be8a2d3b3b17174a74fe317.
//
// Solidity: event TaskCreated(address indexed creator, bytes32 indexed taskHash, address indexed avs, uint32 executorOperatorSetId, address refundCollector, uint96 avsFee, uint256 taskDeadline, bytes payload)
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
