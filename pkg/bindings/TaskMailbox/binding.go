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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_bn254CertificateVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_ecdsaCertificateVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_version\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"BN254_CERTIFICATE_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ECDSA_CERTIFICATE_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createTask\",\"inputs\":[{\"name\":\"taskParams\",\"type\":\"tuple\",\"internalType\":\"structITaskMailboxTypes.TaskParams\",\"components\":[{\"name\":\"refundCollector\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"executorOperatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"feeSplit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"feeSplitCollector\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBN254CertificateBytes\",\"inputs\":[{\"name\":\"cert\",\"type\":\"tuple\",\"internalType\":\"structIBN254CertificateVerifierTypes.BN254Certificate\",\"components\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"nonSignerWitnesses\",\"type\":\"tuple[]\",\"internalType\":\"structIBN254CertificateVerifierTypes.BN254OperatorInfoWitness[]\",\"components\":[{\"name\":\"operatorIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorInfoProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"operatorInfo\",\"type\":\"tuple\",\"internalType\":\"structIOperatorTableCalculatorTypes.BN254OperatorInfo\",\"components\":[{\"name\":\"pubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}]}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getECDSACertificateBytes\",\"inputs\":[{\"name\":\"cert\",\"type\":\"tuple\",\"internalType\":\"structIECDSACertificateVerifierTypes.ECDSACertificate\",\"components\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sig\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getExecutorOperatorSetTaskConfig\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structITaskMailboxTypes.ExecutorOperatorSetTaskConfig\",\"components\":[{\"name\":\"taskHook\",\"type\":\"address\",\"internalType\":\"contractIAVSTaskHook\"},{\"name\":\"taskSLA\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"curveType\",\"type\":\"uint8\",\"internalType\":\"enumIKeyRegistrarTypes.CurveType\"},{\"name\":\"feeCollector\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"consensus\",\"type\":\"tuple\",\"internalType\":\"structITaskMailboxTypes.Consensus\",\"components\":[{\"name\":\"consensusType\",\"type\":\"uint8\",\"internalType\":\"enumITaskMailboxTypes.ConsensusType\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"taskMetadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskInfo\",\"inputs\":[{\"name\":\"taskHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structITaskMailboxTypes.Task\",\"components\":[{\"name\":\"creator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"creationTime\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avsFee\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"refundCollector\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"executorOperatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeSplit\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumITaskMailboxTypes.TaskStatus\"},{\"name\":\"isFeeRefunded\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"operatorTableReferenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"executorOperatorSetTaskConfig\",\"type\":\"tuple\",\"internalType\":\"structITaskMailboxTypes.ExecutorOperatorSetTaskConfig\",\"components\":[{\"name\":\"taskHook\",\"type\":\"address\",\"internalType\":\"contractIAVSTaskHook\"},{\"name\":\"taskSLA\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"curveType\",\"type\":\"uint8\",\"internalType\":\"enumIKeyRegistrarTypes.CurveType\"},{\"name\":\"feeCollector\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"consensus\",\"type\":\"tuple\",\"internalType\":\"structITaskMailboxTypes.Consensus\",\"components\":[{\"name\":\"consensusType\",\"type\":\"uint8\",\"internalType\":\"enumITaskMailboxTypes.ConsensusType\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"taskMetadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"executorCert\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"result\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskResult\",\"inputs\":[{\"name\":\"taskHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskStatus\",\"inputs\":[{\"name\":\"taskHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumITaskMailboxTypes.TaskStatus\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_feeSplit\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"_feeSplitCollector\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isExecutorOperatorSetRegistered\",\"inputs\":[{\"name\":\"operatorSetKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"isRegistered\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"refundFee\",\"inputs\":[{\"name\":\"taskHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerExecutorOperatorSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"isRegistered\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setExecutorOperatorSetTaskConfig\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structITaskMailboxTypes.ExecutorOperatorSetTaskConfig\",\"components\":[{\"name\":\"taskHook\",\"type\":\"address\",\"internalType\":\"contractIAVSTaskHook\"},{\"name\":\"taskSLA\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"curveType\",\"type\":\"uint8\",\"internalType\":\"enumIKeyRegistrarTypes.CurveType\"},{\"name\":\"feeCollector\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"consensus\",\"type\":\"tuple\",\"internalType\":\"structITaskMailboxTypes.Consensus\",\"components\":[{\"name\":\"consensusType\",\"type\":\"uint8\",\"internalType\":\"enumITaskMailboxTypes.ConsensusType\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"taskMetadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFeeSplit\",\"inputs\":[{\"name\":\"_feeSplit\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFeeSplitCollector\",\"inputs\":[{\"name\":\"_feeSplitCollector\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitResult\",\"inputs\":[{\"name\":\"taskHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"executorCert\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"result\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"ExecutorOperatorSetRegistered\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"executorOperatorSetId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"isRegistered\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExecutorOperatorSetTaskConfigSet\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"executorOperatorSetId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structITaskMailboxTypes.ExecutorOperatorSetTaskConfig\",\"components\":[{\"name\":\"taskHook\",\"type\":\"address\",\"internalType\":\"contractIAVSTaskHook\"},{\"name\":\"taskSLA\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"curveType\",\"type\":\"uint8\",\"internalType\":\"enumIKeyRegistrarTypes.CurveType\"},{\"name\":\"feeCollector\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"consensus\",\"type\":\"tuple\",\"internalType\":\"structITaskMailboxTypes.Consensus\",\"components\":[{\"name\":\"consensusType\",\"type\":\"uint8\",\"internalType\":\"enumITaskMailboxTypes.ConsensusType\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"taskMetadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeRefunded\",\"inputs\":[{\"name\":\"refundCollector\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"taskHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"avsFee\",\"type\":\"uint96\",\"indexed\":false,\"internalType\":\"uint96\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeSplitCollectorSet\",\"inputs\":[{\"name\":\"feeSplitCollector\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeSplitSet\",\"inputs\":[{\"name\":\"feeSplit\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskCreated\",\"inputs\":[{\"name\":\"creator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"taskHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"executorOperatorSetId\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"operatorTableReferenceTimestamp\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"refundCollector\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"avsFee\",\"type\":\"uint96\",\"indexed\":false,\"internalType\":\"uint96\"},{\"name\":\"taskDeadline\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskVerified\",\"inputs\":[{\"name\":\"aggregator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"taskHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"executorOperatorSetId\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"executorCert\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"result\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CertificateVerificationFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EmptyCertificateSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExecutorOperatorSetNotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExecutorOperatorSetTaskConfigNotSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FeeAlreadyRefunded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidConsensusType\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidConsensusValue\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCurveType\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidFeeReceiver\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidFeeSplit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidMessageHash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperatorSetOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidReferenceTimestamp\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTaskCreator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTaskStatus\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"enumITaskMailboxTypes.TaskStatus\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"enumITaskMailboxTypes.TaskStatus\"}]},{\"type\":\"error\",\"name\":\"OnlyRefundCollector\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PayloadIsEmpty\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"TaskSLAIsZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TimestampAtCreation\",\"inputs\":[]}]",
	Bin: "0x60e060405234801561000f575f5ffd5b506040516158fb3803806158fb83398101604081905261002e9161018c565b6001600160a01b03808416608052821660a0528061004b8161005f565b60c052506100576100a5565b5050506102b8565b5f5f829050601f81511115610092578260405163305a27a960e01b8152600401610089919061025d565b60405180910390fd5b805161009d82610292565b179392505050565b5f54610100900460ff161561010c5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b6064820152608401610089565b5f5460ff9081161461015b575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b80516001600160a01b0381168114610173575f5ffd5b919050565b634e487b7160e01b5f52604160045260245ffd5b5f5f5f6060848603121561019e575f5ffd5b6101a78461015d565b92506101b56020850161015d565b60408501519092506001600160401b038111156101d0575f5ffd5b8401601f810186136101e0575f5ffd5b80516001600160401b038111156101f9576101f9610178565b604051601f8201601f19908116603f011681016001600160401b038111828210171561022757610227610178565b60405281815282820160200188101561023e575f5ffd5b8160208401602083015e5f602083830101528093505050509250925092565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b805160208083015191908110156102b2575f198160200360031b1b821691505b50919050565b60805160a05160c0516155fd6102fe5f395f611a4a01525f8181610247015281816131d9015261393f01525f81816103550152818161319901526137ff01526155fd5ff3fe608060405234801561000f575f5ffd5b5060043610610148575f3560e01c80636373ea69116100bf578063a5fabc8111610079578063a5fabc8114610304578063b869416614610317578063f2fde38b1461032a578063f741e81a1461033d578063f7424fc914610350578063fa2c0b3714610377575f5ffd5b80636373ea6914610284578063678fbdb3146102a55780636bf6fad5146102b8578063708c0db9146102d8578063715018a6146102eb5780638da5cb5b146102f3575f5ffd5b8063468c07a011610110578063468c07a0146101fa57806349acd8841461020f5780634ad52e021461022257806354743ad21461024257806354fd4d501461026957806362fee03714610271575f5ffd5b80631270a8921461014c5780631a20c505146101755780631ae370eb146101a65780631fb66f5d146101b95780632bf6cc79146101da575b5f5ffd5b61015f61015a366004613f98565b6103a9565b60405161016c919061404d565b60405180910390f35b609b5461018e906201000090046001600160a01b031681565b6040516001600160a01b03909116815260200161016c565b61015f6101b43660046142aa565b6103d2565b6101cc6101c73660046143e9565b6103e5565b60405190815260200161016c565b6101ed6101e836600461445e565b610d2f565b60405161016c919061449d565b61020d6102083660046144ba565b6111db565b005b61020d61021d3660046144e2565b6111ef565b61023561023036600461445e565b6114aa565b60405161016c91906145fe565b61018e7f000000000000000000000000000000000000000000000000000000000000000081565b61015f611a43565b61015f61027f36600461445e565b611a73565b609b546102929061ffff1681565b60405161ffff909116815260200161016c565b61020d6102b3366004614749565b611f70565b6102cb6102c6366004614764565b611f81565b60405161016c919061477e565b61020d6102e6366004614790565b612195565b61020d6122c6565b6033546001600160a01b031661018e565b61020d6103123660046147d8565b6122d9565b61020d61032536600461445e565b612bd0565b61020d610338366004614749565b612e02565b61020d61034b3660046148ca565b612e78565b61018e7f000000000000000000000000000000000000000000000000000000000000000081565b61039961038536600461445e565b60996020525f908152604090205460ff1681565b604051901515815260200161016c565b6060816040516020016103bc91906149f5565b6040516020818303038152906040529050919050565b6060816040516020016103bc9190614b79565b5f6103ee6130ba565b5f8260400151511161041357604051636b1a1b6960e11b815260040160405180910390fd5b60995f6104238460200151613113565b815260208101919091526040015f205460ff166104535760405163c292b29760e01b815260040160405180910390fd5b5f609a5f6104648560200151613113565b815260208082019290925260409081015f20815160e08101835281546001600160a01b038082168352600160a01b918290046001600160601b03169583019590955260018301549485169382019390935292909160608401910460ff1660028111156104d2576104d2614475565b60028111156104e3576104e3614475565b815260028201546001600160a01b0316602082015260408051808201825260038401805492909301929091829060ff16600181111561052457610524614475565b600181111561053557610535614475565b815260200160018201805461054990614b8b565b80601f016020809104026020016040519081016040528092919081815260200182805461057590614b8b565b80156105c05780601f10610597576101008083540402835291602001916105c0565b820191905f5260205f20905b8154815290600101906020018083116105a357829003601f168201915b50505050508152505081526020016005820180546105dd90614b8b565b80601f016020809104026020016040519081016040528092919081815260200182805461060990614b8b565b80156106545780601f1061062b57610100808354040283529160200191610654565b820191905f5260205f20905b81548152906001019060200180831161063757829003601f168201915b50505050508152505090505f600281111561067157610671614475565b8160600151600281111561068757610687614475565b1415801561069e575080516001600160a01b031615155b80156106b657505f81602001516001600160601b0316115b80156106d857505f60a08201515160018111156106d5576106d5614475565b14155b6106f5576040516314b0a41d60e11b815260040160405180910390fd5b8051604051630a3fc61360e31b81526001600160a01b03909116906351fe3098906107269033908790600401614c17565b5f6040518083038186803b15801561073c575f5ffd5b505afa15801561074e573d5f5f3e3d5ffd5b50508251604051637036693f60e11b81525f93506001600160a01b03909116915063e06cd27e90610783908790600401614c3a565b602060405180830381865afa15801561079e573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107c29190614c4c565b90505f6107d2836060015161317c565b6001600160a01b0316635ddb9b5b86602001516040518263ffffffff1660e01b81526004016108019190614c67565b602060405180830381865afa15801561081c573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108409190614c8d565b90505f60975430468860405160200161085c9493929190614ca8565b60405160208183030381529060405280519060200120905060975460016108839190614ce8565b609755604080516101c08101909152338152602081016108a242613216565b6001600160601b03908116825260208981018051516001600160a01b039081168386015292881660408501528a5190921660608401529051015163ffffffff166080820152609b5461ffff1660a082015260c001600181525f602080830182905263ffffffff80871660408086019190915260608086018b90528c8201516080808801919091528251808601845286815260a0808901919091528351808701855287815260c09889015289875260988652958390208851958901516001600160601b03908116600160a01b9081026001600160a01b03988916178355948a0151938a0151168402928616929092176001830155870151600282018054968901519789015161ffff16600160c01b0261ffff60c01b19989095169093026001600160c01b0319909616941693909317939093179384168117835560e0850151919391929160ff60d01b1990911662ffffff60c01b1990911617600160d01b836003811115610a1157610a11614475565b02179055506101008201516002808301805461012086015163ffffffff16600160e01b026001600160e01b03941515600160d81b02949094166001600160d81b039091161792909217909155610140830151805160208201516001600160601b0316600160a01b9081026001600160a01b0392831617600386019081556040840151600487018054919094166001600160a01b031982168117855560608601519596929594936001600160a81b031990921617918490811115610ad657610ad6614475565b021790555060808201516002820180546001600160a01b0319166001600160a01b0390921691909117905560a08201518051600383018054909190829060ff191660018381811115610b2a57610b2a614475565b021790555060208201516001820190610b439082614d46565b50505060c08201516005820190610b5a9082614d46565b5050506101608201516009820190610b729082614d46565b50610180820151600a820190610b889082614d46565b506101a0820151600b820190610b9e9082614d46565b50505060408401516001600160a01b031615801590610bc557505f836001600160601b0316115b15610c405760808401516001600160a01b0316610bf557604051633480121760e21b815260040160405180910390fd5b85516001600160a01b0316610c1d57604051633480121760e21b815260040160405180910390fd5b6040840151610c40906001600160a01b031633306001600160601b038716613281565b8351604051629c5c4560e41b8152600481018390526001600160a01b03909116906309c5c450906024015f604051808303815f87803b158015610c81575f5ffd5b505af1158015610c93573d5f5f3e3d5ffd5b5050505085602001515f01516001600160a01b031681336001600160a01b03167f33add0b01e02278be5459fbfa3274aee699ec47f4ee7236b59e7a2c8b5000c26896020015160200151868b5f0151898b602001516001600160601b031642610cfc9190614ce8565b8e60400151604051610d1396959493929190614e00565b60405180910390a49350505050610d2a6001606555565b919050565b5f81815260986020908152604080832081516101c08101835281546001600160a01b038082168352600160a01b918290046001600160601b03908116968401969096526001840154808216958401959095529381900490941660608201526002820154928316608082015292820463ffffffff1660a0840152600160c01b820461ffff1660c084015283929160e0830190600160d01b900460ff166003811115610ddb57610ddb614475565b6003811115610dec57610dec614475565b8152600282810154600160d81b810460ff9081161515602080860191909152600160e01b90920463ffffffff16604080860191909152805160e0810182526003870180546001600160a01b038082168452600160a01b918290046001600160601b03169684019690965260048901549586169383019390935260609687019691959094918601939290910490911690811115610e8a57610e8a614475565b6002811115610e9b57610e9b614475565b815260028201546001600160a01b0316602082015260408051808201825260038401805492909301929091829060ff166001811115610edc57610edc614475565b6001811115610eed57610eed614475565b8152602001600182018054610f0190614b8b565b80601f0160208091040260200160405190810160405280929190818152602001828054610f2d90614b8b565b8015610f785780601f10610f4f57610100808354040283529160200191610f78565b820191905f5260205f20905b815481529060010190602001808311610f5b57829003601f168201915b5050505050815250508152602001600582018054610f9590614b8b565b80601f0160208091040260200160405190810160405280929190818152602001828054610fc190614b8b565b801561100c5780601f10610fe35761010080835404028352916020019161100c565b820191905f5260205f20905b815481529060010190602001808311610fef57829003601f168201915b505050505081525050815260200160098201805461102990614b8b565b80601f016020809104026020016040519081016040528092919081815260200182805461105590614b8b565b80156110a05780601f10611077576101008083540402835291602001916110a0565b820191905f5260205f20905b81548152906001019060200180831161108357829003601f168201915b50505050508152602001600a820180546110b990614b8b565b80601f01602080910402602001604051908101604052809291908181526020018280546110e590614b8b565b80156111305780601f1061110757610100808354040283529160200191611130565b820191905f5260205f20905b81548152906001019060200180831161111357829003601f168201915b50505050508152602001600b8201805461114990614b8b565b80601f016020809104026020016040519081016040528092919081815260200182805461117590614b8b565b80156111c05780601f10611197576101008083540402835291602001916111c0565b820191905f5260205f20905b8154815290600101906020018083116111a357829003601f168201915b50505050508152505090506111d4816132f3565b9392505050565b6111e361334e565b6111ec816133a8565b50565b5f609a5f6111fc85613113565b815260208082019290925260409081015f20815160e08101835281546001600160a01b038082168352600160a01b918290046001600160601b03169583019590955260018301549485169382019390935292909160608401910460ff16600281111561126a5761126a614475565b600281111561127b5761127b614475565b815260028201546001600160a01b0316602082015260408051808201825260038401805492909301929091829060ff1660018111156112bc576112bc614475565b60018111156112cd576112cd614475565b81526020016001820180546112e190614b8b565b80601f016020809104026020016040519081016040528092919081815260200182805461130d90614b8b565b80156113585780601f1061132f57610100808354040283529160200191611358565b820191905f5260205f20905b81548152906001019060200180831161133b57829003601f168201915b505050505081525050815260200160058201805461137590614b8b565b80601f01602080910402602001604051908101604052809291908181526020018280546113a190614b8b565b80156113ec5780601f106113c3576101008083540402835291602001916113ec565b820191905f5260205f20905b8154815290600101906020018083116113cf57829003601f168201915b50505050508152505090505f600281111561140957611409614475565b8160600151600281111561141f5761141f614475565b14158015611436575080516001600160a01b031615155b801561144e57505f81602001516001600160601b0316115b801561147057505f60a082015151600181111561146d5761146d614475565b14155b61148d576040516314b0a41d60e11b815260040160405180910390fd5b61149b838260600151613419565b6114a583836134ba565b505050565b6114b2613d78565b5f82815260986020908152604080832081516101c08101835281546001600160a01b038082168352600160a01b918290046001600160601b03908116968401969096526001840154808216958401959095529381900490941660608201526002820154928316608082015292820463ffffffff1660a0840152600160c01b820461ffff1660c08401529060e0830190600160d01b900460ff16600381111561155c5761155c614475565b600381111561156d5761156d614475565b8152600282810154600160d81b810460ff9081161515602080860191909152600160e01b90920463ffffffff16604080860191909152805160e0810182526003870180546001600160a01b038082168452600160a01b918290046001600160601b0316968401969096526004890154958616938301939093526060968701969195909491860193929091049091169081111561160b5761160b614475565b600281111561161c5761161c614475565b815260028201546001600160a01b0316602082015260408051808201825260038401805492909301929091829060ff16600181111561165d5761165d614475565b600181111561166e5761166e614475565b815260200160018201805461168290614b8b565b80601f01602080910402602001604051908101604052809291908181526020018280546116ae90614b8b565b80156116f95780601f106116d0576101008083540402835291602001916116f9565b820191905f5260205f20905b8154815290600101906020018083116116dc57829003601f168201915b505050505081525050815260200160058201805461171690614b8b565b80601f016020809104026020016040519081016040528092919081815260200182805461174290614b8b565b801561178d5780601f106117645761010080835404028352916020019161178d565b820191905f5260205f20905b81548152906001019060200180831161177057829003601f168201915b50505050508152505081526020016009820180546117aa90614b8b565b80601f01602080910402602001604051908101604052809291908181526020018280546117d690614b8b565b80156118215780601f106117f857610100808354040283529160200191611821565b820191905f5260205f20905b81548152906001019060200180831161180457829003601f168201915b50505050508152602001600a8201805461183a90614b8b565b80601f016020809104026020016040519081016040528092919081815260200182805461186690614b8b565b80156118b15780601f10611888576101008083540402835291602001916118b1565b820191905f5260205f20905b81548152906001019060200180831161189457829003601f168201915b50505050508152602001600b820180546118ca90614b8b565b80601f01602080910402602001604051908101604052809291908181526020018280546118f690614b8b565b80156119415780601f1061191857610100808354040283529160200191611941565b820191905f5260205f20905b81548152906001019060200180831161192457829003601f168201915b5050505050815250509050604051806101c00160405280825f01516001600160a01b0316815260200182602001516001600160601b0316815260200182604001516001600160a01b0316815260200182606001516001600160601b0316815260200182608001516001600160a01b031681526020018260a0015163ffffffff1681526020018260c0015161ffff1681526020016119dd836132f3565b60038111156119ee576119ee614475565b81526020018261010001511515815260200182610120015163ffffffff168152602001826101400151815260200182610160015181526020018261018001518152602001826101a00151815250915050919050565b6060611a6e7f000000000000000000000000000000000000000000000000000000000000000061354b565b905090565b5f81815260986020908152604080832081516101c08101835281546001600160a01b038082168352600160a01b918290046001600160601b0390811696840196909652600184015480821695840195909552938190049094166060808301919091526002830154938416608083015293830463ffffffff1660a0820152600160c01b830461ffff1660c08201529293929160e0830190600160d01b900460ff166003811115611b2457611b24614475565b6003811115611b3557611b35614475565b8152600282810154600160d81b810460ff9081161515602080860191909152600160e01b90920463ffffffff16604080860191909152805160e0810182526003870180546001600160a01b038082168452600160a01b918290046001600160601b03169684019690965260048901549586169383019390935260609687019691959094918601939290910490911690811115611bd357611bd3614475565b6002811115611be457611be4614475565b815260028201546001600160a01b0316602082015260408051808201825260038401805492909301929091829060ff166001811115611c2557611c25614475565b6001811115611c3657611c36614475565b8152602001600182018054611c4a90614b8b565b80601f0160208091040260200160405190810160405280929190818152602001828054611c7690614b8b565b8015611cc15780601f10611c9857610100808354040283529160200191611cc1565b820191905f5260205f20905b815481529060010190602001808311611ca457829003601f168201915b5050505050815250508152602001600582018054611cde90614b8b565b80601f0160208091040260200160405190810160405280929190818152602001828054611d0a90614b8b565b8015611d555780601f10611d2c57610100808354040283529160200191611d55565b820191905f5260205f20905b815481529060010190602001808311611d3857829003601f168201915b5050505050815250508152602001600982018054611d7290614b8b565b80601f0160208091040260200160405190810160405280929190818152602001828054611d9e90614b8b565b8015611de95780601f10611dc057610100808354040283529160200191611de9565b820191905f5260205f20905b815481529060010190602001808311611dcc57829003601f168201915b50505050508152602001600a82018054611e0290614b8b565b80601f0160208091040260200160405190810160405280929190818152602001828054611e2e90614b8b565b8015611e795780601f10611e5057610100808354040283529160200191611e79565b820191905f5260205f20905b815481529060010190602001808311611e5c57829003601f168201915b50505050508152602001600b82018054611e9290614b8b565b80601f0160208091040260200160405190810160405280929190818152602001828054611ebe90614b8b565b8015611f095780601f10611ee057610100808354040283529160200191611f09565b820191905f5260205f20905b815481529060010190602001808311611eec57829003601f168201915b50505050508152505090505f611f1e826132f3565b90506002816003811115611f3457611f34614475565b146002829091611f6257604051634091b18960e11b8152600401611f59929190614e59565b60405180910390fd5b5050506101a0015192915050565b611f7861334e565b6111ec81613588565b611f89613deb565b609a5f611f9584613113565b815260208082019290925260409081015f20815160e08101835281546001600160a01b038082168352600160a01b918290046001600160601b03169583019590955260018301549485169382019390935292909160608401910460ff16600281111561200357612003614475565b600281111561201457612014614475565b815260028201546001600160a01b0316602082015260408051808201825260038401805492909301929091829060ff16600181111561205557612055614475565b600181111561206657612066614475565b815260200160018201805461207a90614b8b565b80601f01602080910402602001604051908101604052809291908181526020018280546120a690614b8b565b80156120f15780601f106120c8576101008083540402835291602001916120f1565b820191905f5260205f20905b8154815290600101906020018083116120d457829003601f168201915b505050505081525050815260200160058201805461210e90614b8b565b80601f016020809104026020016040519081016040528092919081815260200182805461213a90614b8b565b80156121855780601f1061215c57610100808354040283529160200191612185565b820191905f5260205f20905b81548152906001019060200180831161216857829003601f168201915b5050505050815250509050919050565b5f54610100900460ff16158080156121b357505f54600160ff909116105b806121cc5750303b1580156121cc57505f5460ff166001145b61222f5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401611f59565b5f805460ff191660011790558015612250575f805461ff0019166101001790555b612258613602565b612260613630565b6122698461365e565b612272836133a8565b61227b82613588565b80156122c0575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050565b6122ce61334e565b6122d75f61365e565b565b6122e16130ba565b5f83815260986020908152604080832081516101c08101835281546001600160a01b038082168352600160a01b918290046001600160601b03908116968401969096526001840154808216958401959095529381900490941660608201526002820154928316608082015292820463ffffffff1660a0840152600160c01b820461ffff1660c0840152929161278591849060e0830190600160d01b900460ff16600381111561239257612392614475565b60038111156123a3576123a3614475565b8152600282810154600160d81b810460ff9081161515602080860191909152600160e01b90920463ffffffff16604080860191909152805160e0810182526003870180546001600160a01b038082168452600160a01b918290046001600160601b0316968401969096526004890154958616938301939093526060968701969195909491860193929091049091169081111561244157612441614475565b600281111561245257612452614475565b815260028201546001600160a01b0316602082015260408051808201825260038401805492909301929091829060ff16600181111561249357612493614475565b60018111156124a4576124a4614475565b81526020016001820180546124b890614b8b565b80601f01602080910402602001604051908101604052809291908181526020018280546124e490614b8b565b801561252f5780601f106125065761010080835404028352916020019161252f565b820191905f5260205f20905b81548152906001019060200180831161251257829003601f168201915b505050505081525050815260200160058201805461254c90614b8b565b80601f016020809104026020016040519081016040528092919081815260200182805461257890614b8b565b80156125c35780601f1061259a576101008083540402835291602001916125c3565b820191905f5260205f20905b8154815290600101906020018083116125a657829003601f168201915b50505050508152505081526020016009820180546125e090614b8b565b80601f016020809104026020016040519081016040528092919081815260200182805461260c90614b8b565b80156126575780601f1061262e57610100808354040283529160200191612657565b820191905f5260205f20905b81548152906001019060200180831161263a57829003601f168201915b50505050508152602001600a8201805461267090614b8b565b80601f016020809104026020016040519081016040528092919081815260200182805461269c90614b8b565b80156126e75780601f106126be576101008083540402835291602001916126e7565b820191905f5260205f20905b8154815290600101906020018083116126ca57829003601f168201915b50505050508152602001600b8201805461270090614b8b565b80601f016020809104026020016040519081016040528092919081815260200182805461272c90614b8b565b80156127775780601f1061274e57610100808354040283529160200191612777565b820191905f5260205f20905b81548152906001019060200180831161275a57829003601f168201915b5050505050815250506132f3565b9050600181600381111561279b5761279b614475565b1460018290916127c057604051634091b18960e11b8152600401611f59929190614e59565b50508154600160a01b90046001600160601b031642116127f35760405163015a4b7560e51b815260040160405180910390fd5b600382015460405163ba33565d60e01b81526001600160a01b039091169063ba33565d9061282b903390899089908990600401614e74565b5f6040518083038186803b158015612841575f5ffd5b505afa158015612853573d5f5f3e3d5ffd5b50506040805180820182526001808701546001600160a01b03168252600287015463ffffffff600160a01b91829004166020840152600488015484518086019095526006890180549497505f965061298d9560ff9390920483169491939092849216908111156128c5576128c5614475565b60018111156128d6576128d6614475565b81526020016001820180546128ea90614b8b565b80601f016020809104026020016040519081016040528092919081815260200182805461291690614b8b565b80156129615780601f1061293857610100808354040283529160200191612961565b820191905f5260205f20905b81548152906001019060200180831161294457829003601f168201915b505050919092525050506002870154885160208a01208691600160e01b900463ffffffff16908b6136af565b9050806129ad5760405163efcb98e760e01b815260040160405180910390fd5b60028401805460ff60d01b1916600160d11b179055600a84016129d08782614d46565b50600b84016129df8682614d46565b5060048401546001600160a01b031615801590612a0f57506001840154600160a01b90046001600160601b031615155b15612afe57600284015460018501545f91612a5e9161271091612a4f91600160c01b90910461ffff16906001600160601b03600160a01b90910416614eac565b612a599190614ec3565b613216565b90506001600160601b03811615612a9e57609b546004860154612a9e916001600160a01b039182169162010000909104166001600160601b0384166139ed565b60018501545f90612ac0908390600160a01b90046001600160601b0316614ee2565b90506001600160601b03811615612afb5760058601546004870154612afb916001600160a01b0391821691166001600160601b0384166139ed565b50505b600384015460405163db6ecf6760e01b8152600481018990526001600160a01b039091169063db6ecf67906024015f604051808303815f87803b158015612b43575f5ffd5b505af1158015612b55573d5f5f3e3d5ffd5b505050600185015460028601546040516001600160a01b039092169250899133917f659f23b2e7edf490e5fd6561c5148691ed0375ed7ddd3ab1bcfcfdbec4f209a991612bba9163ffffffff600160a01b9091041690600a8b0190600b8c0190614f80565b60405180910390a4505050506114a56001606555565b612bd86130ba565b5f81815260986020526040902060028101546001600160a01b03163314612c12576040516370f43cb760e01b815260040160405180910390fd5b6002810154600160d81b900460ff1615612c3f57604051633e3d786960e01b815260040160405180910390fd5b604080516101c08101825282546001600160a01b038082168352600160a01b918290046001600160601b03908116602085015260018601548083169585019590955293829004909316606083015260028401549283166080830152820463ffffffff1660a0820152600160c01b820461ffff1660c08201525f91612ce09190849060e0830190600160d01b900460ff16600381111561239257612392614475565b90506003816003811115612cf657612cf6614475565b146003829091612d1b57604051634091b18960e11b8152600401611f59929190614e59565b505060028201805460ff60d81b1916600160d81b17905560048201546001600160a01b031615801590612d6157506001820154600160a01b90046001600160601b031615155b15612d9d57600282015460018301546004840154612d9d926001600160a01b0391821692911690600160a01b90046001600160601b03166139ed565b60028201546001830154604051600160a01b9091046001600160601b0316815284916001600160a01b0316907fe3ed40d31808582f7a92a30beacc0ec788d5091407ec6c10c1b999b3f317aea39060200160405180910390a350506111ec6001606555565b612e0a61334e565b6001600160a01b038116612e6f5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401611f59565b6111ec8161365e565b5f81606001516002811115612e8f57612e8f614475565b03612ead5760405163fdea7c0960e01b815260040160405180910390fd5b80516001600160a01b0316612ed557604051630863a45360e11b815260040160405180910390fd5b5f81602001516001600160601b031611612f025760405163568added60e11b815260040160405180910390fd5b612f0f8160a00151613a1d565b612f1d828260600151613419565b80609a5f612f2a85613113565b815260208082019290925260409081015f208351928401516001600160601b0316600160a01b9081026001600160a01b0394851617825591840151600182018054919094166001600160a01b03198216811785556060860151929492936001600160a81b03199092161790836002811115612fa757612fa7614475565b021790555060808201516002820180546001600160a01b0319166001600160a01b0390921691909117905560a08201518051600383018054909190829060ff191660018381811115612ffb57612ffb614475565b0217905550602082015160018201906130149082614d46565b50505060c0820151600582019061302b9082614d46565b50905050816020015163ffffffff16825f01516001600160a01b0316336001600160a01b03167f7cd76abd4025a20959a1b20f7c1536e3894a0735cd8de0215dde803ddea7f2d284604051613080919061477e565b60405180910390a460995f61309484613113565b815260208101919091526040015f205460ff166130b6576130b68260016134ba565b5050565b60026065540361310c5760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401611f59565b6002606555565b5f815f0151826020015163ffffffff1660405160200161315e92919060609290921b6bffffffffffffffffffffffff1916825260a01b6001600160a01b031916601482015260200190565b60405160208183030381529060405261317690614fb0565b92915050565b5f600282600281111561319157613191614475565b036131bd57507f0000000000000000000000000000000000000000000000000000000000000000919050565b60018260028111156131d1576131d1614475565b036131fd57507f0000000000000000000000000000000000000000000000000000000000000000919050565b60405163fdea7c0960e01b815260040160405180910390fd5b5f6001600160601b0382111561327d5760405162461bcd60e51b815260206004820152602660248201527f53616665436173743a2076616c756520646f65736e27742066697420696e203960448201526536206269747360d01b6064820152608401611f59565b5090565b6040516001600160a01b03808516602483015283166044820152606481018290526122c09085906323b872dd60e01b906084015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b031990931692909217909152613a9f565b6001606555565b5f60018260e00151600381111561330c5761330c614475565b148015613339575081610140015160200151826020015161332d9190614fd3565b6001600160601b031642115b1561334657506003919050565b5060e0015190565b6033546001600160a01b031633146122d75760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401611f59565b61271061ffff821611156133cf57604051630601f69760e01b815260040160405180910390fd5b609b805461ffff191661ffff83169081179091556040519081527f886b2cfcb151fd8b19ed902cc88f4a06dd9fe351a4a9ab93f33fe84abc157edf9060200160405180910390a150565b5f6134238261317c565b6040516304240c4960e51b815290915033906001600160a01b03831690638481892090613454908790600401614c67565b602060405180830381865afa15801561346f573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906134939190614ff2565b6001600160a01b0316146114a5576040516342ecfee960e11b815260040160405180910390fd5b8060995f6134c785613113565b81526020019081526020015f205f6101000a81548160ff021916908315150217905550816020015163ffffffff16825f01516001600160a01b0316336001600160a01b03167f48b63f21a1eb9dd6880e196de6d7db3fbd0c282b74f1298dcb4cf53472298f398460405161353f911515815260200190565b60405180910390a45050565b60605f61355783613b72565b6040805160208082528183019092529192505f91906020820181803683375050509182525060208101929092525090565b6001600160a01b0381166135af57604051630863a45360e11b815260040160405180910390fd5b609b805462010000600160b01b031916620100006001600160a01b038416908102919091179091556040517f262aa27c244f6f0088cb3092548a0adcaddedf459070a9ccab2dc6a07abe701d905f90a250565b5f54610100900460ff166136285760405162461bcd60e51b8152600401611f599061500d565b6122d7613b99565b5f54610100900460ff166136565760405162461bcd60e51b8152600401611f599061500d565b6122d7613bc8565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b5f6001865160018111156136c5576136c5614475565b036139ca575f86602001518060200190518101906136e39190615058565b6040805160018082528183019092529192505f91906020808301908036833701905050905081815f8151811061371b5761371b615073565b61ffff90921660209283029190910190910152600289600281111561374257613742614475565b03613880575f8480602001905181019061375c91906152e8565b90508663ffffffff16815f015163ffffffff161461378d57604051634534032960e01b815260040160405180910390fd5b858160200151146137b157604051638b56642d60e01b815260040160405180910390fd5b604081015151158015906137cc575060408101516020015115155b6137e957604051637a8a1dbd60e11b815260040160405180910390fd5b604051625f5e5d60e21b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063017d797490613838908b90859087906004016153f8565b6020604051808303815f875af1158015613854573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613878919061543c565b9350506139c3565b600189600281111561389457613894614475565b036131fd575f848060200190518101906138ae9190615457565b90508663ffffffff16815f015163ffffffff16146138df57604051634534032960e01b815260040160405180910390fd5b8581602001511461390357604051638b56642d60e01b815260040160405180910390fd5b5f8160400151511161392857604051637a8a1dbd60e11b815260040160405180910390fd5b604051630606d12160e51b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063c0da242090613978908b90859087906004016154d0565b5f60405180830381865afa158015613992573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526139b99190810190615502565b5093506139c39050565b50506139e3565b6040516347d3772160e11b815260040160405180910390fd5b9695505050505050565b6040516001600160a01b0383166024820152604481018290526114a590849063a9059cbb60e01b906064016132b5565b600181516001811115613a3257613a32614475565b036139ca57806020015151602014613a5d57604051631501e04760e21b815260040160405180910390fd5b5f8160200151806020019051810190613a769190615058565b905061271061ffff821611156130b657604051631501e04760e21b815260040160405180910390fd5b5f613af3826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316613bee9092919063ffffffff16565b905080515f1480613b13575080806020019051810190613b13919061543c565b6114a55760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608401611f59565b5f60ff8216601f81111561317657604051632cd44ac360e21b815260040160405180910390fd5b5f54610100900460ff16613bbf5760405162461bcd60e51b8152600401611f599061500d565b6122d73361365e565b5f54610100900460ff166132ec5760405162461bcd60e51b8152600401611f599061500d565b6060613bfc84845f85613c04565b949350505050565b606082471015613c655760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b6064820152608401611f59565b5f5f866001600160a01b03168587604051613c8091906155b1565b5f6040518083038185875af1925050503d805f8114613cba576040519150601f19603f3d011682016040523d82523d5f602084013e613cbf565b606091505b5091509150613cd087838387613cdb565b979650505050505050565b60608315613d495782515f03613d42576001600160a01b0385163b613d425760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401611f59565b5081613bfc565b613bfc8383815115613d5e5781518083602001fd5b8060405162461bcd60e51b8152600401611f59919061404d565b604080516101c0810182525f80825260208201819052918101829052606081018290526080810182905260a0810182905260c081018290529060e082019081525f602082018190526040820152606001613dd0613deb565b81526020016060815260200160608152602001606081525090565b6040805160e0810182525f8082526020820181905291810182905290606082019081525f6020820152604001613e1f613e2c565b8152602001606081525090565b60408051808201909152805f613e1f565b634e487b7160e01b5f52604160045260245ffd5b604051606081016001600160401b0381118282101715613e7357613e73613e3d565b60405290565b604080519081016001600160401b0381118282101715613e7357613e73613e3d565b60405160a081016001600160401b0381118282101715613e7357613e73613e3d565b60405160e081016001600160401b0381118282101715613e7357613e73613e3d565b604051601f8201601f191681016001600160401b0381118282101715613f0757613f07613e3d565b604052919050565b63ffffffff811681146111ec575f5ffd5b5f6001600160401b03821115613f3857613f38613e3d565b50601f01601f191660200190565b5f82601f830112613f55575f5ffd5b8135613f68613f6382613f20565b613edf565b818152846020838601011115613f7c575f5ffd5b816020850160208301375f918101602001919091529392505050565b5f60208284031215613fa8575f5ffd5b81356001600160401b03811115613fbd575f5ffd5b820160608185031215613fce575f5ffd5b613fd6613e51565b8135613fe181613f0f565b81526020828101359082015260408201356001600160401b03811115614005575f5ffd5b61401186828501613f46565b604083015250949350505050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f6111d4602083018461401f565b5f6040828403121561406f575f5ffd5b614077613e79565b823581526020928301359281019290925250919050565b5f82601f83011261409d575f5ffd5b6140a76040613edf565b8060408401858111156140b8575f5ffd5b845b818110156140d25780358452602093840193016140ba565b509095945050505050565b5f6001600160401b038211156140f5576140f5613e3d565b5060051b60200190565b5f82601f83011261410e575f5ffd5b813561411c613f63826140dd565b8082825260208201915060208360051b86010192508583111561413d575f5ffd5b602085015b838110156142a05780356001600160401b0381111561415f575f5ffd5b86016060818903601f19011215614174575f5ffd5b61417c613e51565b602082013561418a81613f0f565b815260408201356001600160401b038111156141a4575f5ffd5b6141b38a602083860101613f46565b60208301525060608201356001600160401b038111156141d1575f5ffd5b6020818401019250506060828a0312156141e9575f5ffd5b6141f1613e79565b6141fb8a8461405f565b815260408301356001600160401b03811115614215575f5ffd5b80840193505089601f840112614229575f5ffd5b8235614237613f63826140dd565b8082825260208201915060208360051b87010192508c831115614258575f5ffd5b6020860195505b8286101561427a57853582526020958601959091019061425f565b806020850152505050806040830152508085525050602083019250602081019050614142565b5095945050505050565b5f602082840312156142ba575f5ffd5b81356001600160401b038111156142cf575f5ffd5b82018084036101208112156142e2575f5ffd5b6142ea613e9b565b82356142f581613f0f565b81526020838101359082015261430e866040850161405f565b60408201526080607f1983011215614324575f5ffd5b61432c613e79565b915061433b866080850161408e565b825261434a8660c0850161408e565b602083015281606082015261010083013591506001600160401b03821115614370575f5ffd5b61437c868385016140ff565b608082015295945050505050565b6001600160a01b03811681146111ec575f5ffd5b8035610d2a8161438a565b5f604082840312156143b9575f5ffd5b6143c1613e79565b905081356143ce8161438a565b815260208201356143de81613f0f565b602082015292915050565b5f602082840312156143f9575f5ffd5b81356001600160401b0381111561440e575f5ffd5b82016080818503121561441f575f5ffd5b614427613e51565b81356144328161438a565b815261444185602084016143a9565b602082015260608201356001600160401b03811115614005575f5ffd5b5f6020828403121561446e575f5ffd5b5035919050565b634e487b7160e01b5f52602160045260245ffd5b6004811061449957614499614475565b9052565b602081016131768284614489565b61ffff811681146111ec575f5ffd5b5f602082840312156144ca575f5ffd5b81356111d4816144ab565b80151581146111ec575f5ffd5b5f5f606083850312156144f3575f5ffd5b6144fd84846143a9565b9150604083013561450d816144d5565b809150509250929050565b6003811061449957614499614475565b5f81516002811061453b5761453b614475565b80845250602082015160406020850152613bfc604085018261401f565b80516001600160a01b031682526020808201516001600160601b0316908301526040808201515f91614594908501826001600160a01b03169052565b5060608201516145a76060850182614518565b5060808201516145c260808501826001600160a01b03169052565b5060a082015160e060a08501526145dc60e0850182614528565b905060c083015184820360c08601526145f5828261401f565b95945050505050565b602081526146186020820183516001600160a01b03169052565b5f602083015161463360408401826001600160601b03169052565b5060408301516001600160a01b03811660608401525060608301516001600160601b03811660808401525060808301516001600160a01b03811660a08401525060a083015163ffffffff811660c08401525060c083015161ffff811660e08401525060e08301516146a8610100840182614489565b506101008301518015156101208401525061012083015163ffffffff8116610140840152506101408301516101c06101608401526146ea6101e0840182614558565b9050610160840151601f1984830301610180850152614709828261401f565b915050610180840151601f19848303016101a0850152614729828261401f565b9150506101a0840151601f19848303016101c08501526145f5828261401f565b5f60208284031215614759575f5ffd5b81356111d48161438a565b5f60408284031215614774575f5ffd5b6111d483836143a9565b602081525f6111d46020830184614558565b5f5f5f606084860312156147a2575f5ffd5b83356147ad8161438a565b925060208401356147bd816144ab565b915060408401356147cd8161438a565b809150509250925092565b5f5f5f606084860312156147ea575f5ffd5b8335925060208401356001600160401b03811115614806575f5ffd5b61481286828701613f46565b92505060408401356001600160401b0381111561482d575f5ffd5b61483986828701613f46565b9150509250925092565b6001600160601b03811681146111ec575f5ffd5b8035610d2a81614843565b803560038110610d2a575f5ffd5b5f60408284031215614880575f5ffd5b614888613e79565b9050813560028110614898575f5ffd5b815260208201356001600160401b038111156148b2575f5ffd5b6148be84828501613f46565b60208301525092915050565b5f5f606083850312156148db575f5ffd5b6148e584846143a9565b915060408301356001600160401b038111156148ff575f5ffd5b830160e08186031215614910575f5ffd5b614918613ebd565b6149218261439e565b815261492f60208301614857565b60208201526149406040830161439e565b604082015261495160608301614862565b60608201526149626080830161439e565b608082015260a08201356001600160401b0381111561497f575f5ffd5b61498b87828501614870565b60a08301525060c08201356001600160401b038111156149a9575f5ffd5b6149b587828501613f46565b60c08301525080925050509250929050565b63ffffffff8151168252602081015160208301525f604082015160606040850152613bfc606085018261401f565b602081525f6111d460208301846149c7565b805f5b60028110156122c0578151845260209384019390910190600101614a0a565b5f610120830163ffffffff8351168452602083015160208501526040830151614a5f604086018280518252602090810151910152565b506060830151614a73608086018251614a07565b60200151614a8460c0860182614a07565b506080830151610120610100860152818151808452610140870191506101408160051b88010193506020830192505f5b81811015614b6d5761013f19888603018352835163ffffffff8151168652602081015160606020880152614aeb606088018261401f565b905060408201519150868103604088015260608101614b1582845180518252602090810151910152565b6020928301516060604084015280518083529301925f92608001905b80841015614b545784518252602082019150602085019450600184019350614b31565b5097505050602094850194939093019250600101614ab4565b50929695505050505050565b602081525f6111d46020830184614a29565b600181811c90821680614b9f57607f821691505b602082108103614bbd57634e487b7160e01b5f52602260045260245ffd5b50919050565b80516001600160a01b031682526020808201515f91614bfd9085018280516001600160a01b0316825260209081015163ffffffff16910152565b50604082015160806060850152613bfc608085018261401f565b6001600160a01b03831681526040602082018190525f90613bfc90830184614bc3565b602081525f6111d46020830184614bc3565b5f60208284031215614c5c575f5ffd5b81516111d481614843565b81516001600160a01b0316815260208083015163ffffffff169082015260408101613176565b5f60208284031215614c9d575f5ffd5b81516111d481613f0f565b84815260018060a01b0384166020820152826040820152608060608201525f6139e36080830184614bc3565b634e487b7160e01b5f52601160045260245ffd5b8082018082111561317657613176614cd4565b601f8211156114a557805f5260205f20601f840160051c81016020851015614d205750805b601f840160051c820191505b81811015614d3f575f8155600101614d2c565b5050505050565b81516001600160401b03811115614d5f57614d5f613e3d565b614d7381614d6d8454614b8b565b84614cfb565b6020601f821160018114614da5575f8315614d8e5750848201515b5f19600385901b1c1916600184901b178455614d3f565b5f84815260208120601f198516915b82811015614dd45787850151825560209485019460019092019101614db4565b5084821015614df157868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b63ffffffff8716815263ffffffff8616602082015260018060a01b03851660408201526001600160601b038416606082015282608082015260c060a08201525f614e4d60c083018461401f565b98975050505050505050565b60408101614e678285614489565b6111d46020830184614489565b60018060a01b0385168152836020820152608060408201525f614e9a608083018561401f565b8281036060840152613cd0818561401f565b808202811582820484141761317657613176614cd4565b5f82614edd57634e487b7160e01b5f52601260045260245ffd5b500490565b6001600160601b03828116828216039081111561317657613176614cd4565b5f8154614f0d81614b8b565b808552600182168015614f275760018114614f4357614f77565b60ff1983166020870152602082151560051b8701019350614f77565b845f5260205f205f5b83811015614f6e5781546020828a010152600182019150602081019050614f4c565b87016020019450505b50505092915050565b63ffffffff84168152606060208201525f614f9e6060830185614f01565b82810360408401526139e38185614f01565b80516020808301519190811015614bbd575f1960209190910360031b1b16919050565b6001600160601b03818116838216019081111561317657613176614cd4565b5f60208284031215615002575f5ffd5b81516111d48161438a565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b5f60208284031215615068575f5ffd5b81516111d4816144ab565b634e487b7160e01b5f52603260045260245ffd5b5f60408284031215615097575f5ffd5b61509f613e79565b825181526020928301519281019290925250919050565b5f82601f8301126150c5575f5ffd5b6150cf6040613edf565b8060408401858111156150e0575f5ffd5b845b818110156140d25780518452602093840193016150e2565b5f82601f830112615109575f5ffd5b8151615117613f6382613f20565b81815284602083860101111561512b575f5ffd5b8160208501602083015e5f918101602001919091529392505050565b5f82601f830112615156575f5ffd5b8151615164613f63826140dd565b8082825260208201915060208360051b860101925085831115615185575f5ffd5b602085015b838110156142a05780516001600160401b038111156151a7575f5ffd5b86016060818903601f190112156151bc575f5ffd5b6151c4613e51565b60208201516151d281613f0f565b815260408201516001600160401b038111156151ec575f5ffd5b6151fb8a6020838601016150fa565b60208301525060608201516001600160401b03811115615219575f5ffd5b6020818401019250506060828a031215615231575f5ffd5b615239613e79565b6152438a84615087565b815260408301516001600160401b0381111561525d575f5ffd5b80840193505089601f840112615271575f5ffd5b825161527f613f63826140dd565b8082825260208201915060208360051b87010192508c8311156152a0575f5ffd5b6020860195505b828610156152c25785518252602095860195909101906152a7565b80602085015250505080604083015250808552505060208301925060208101905061518a565b5f602082840312156152f8575f5ffd5b81516001600160401b0381111561530d575f5ffd5b8201808403610120811215615320575f5ffd5b615328613e9b565b825161533381613f0f565b81526020838101519082015261534c8660408501615087565b60408201526080607f1983011215615362575f5ffd5b61536a613e79565b915061537986608085016150b6565b82526153888660c085016150b6565b602083015281606082015261010083015191506001600160401b038211156153ae575f5ffd5b61437c86838501615147565b5f8151808452602084019350602083015f5b828110156153ee57815161ffff168652602095860195909101906001016153cc565b5093949350505050565b83516001600160a01b0316815260208085015163ffffffff1690820152608060408201525f61542a6080830185614a29565b82810360608401526139e381856153ba565b5f6020828403121561544c575f5ffd5b81516111d4816144d5565b5f60208284031215615467575f5ffd5b81516001600160401b0381111561547c575f5ffd5b82016060818503121561548d575f5ffd5b615495613e51565b81516154a081613f0f565b81526020828101519082015260408201516001600160401b038111156154c4575f5ffd5b614011868285016150fa565b83516001600160a01b0316815260208085015163ffffffff1690820152608060408201525f61542a60808301856149c7565b5f5f60408385031215615513575f5ffd5b825161551e816144d5565b60208401519092506001600160401b03811115615539575f5ffd5b8301601f81018513615549575f5ffd5b8051615557613f63826140dd565b8082825260208201915060208360051b850101925087831115615578575f5ffd5b6020840193505b828410156155a35783516155928161438a565b82526020938401939091019061557f565b809450505050509250929050565b5f82518060208501845e5f92019182525091905056fea2646970667358221220b4f338674d01e33301c995e4eeb8a9a730d3723c675189adc4527621e95ae28264736f6c634300081b0033",
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
