// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package EigenPod

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

// BeaconChainProofsBalanceContainerProof is an auto generated low-level Go binding around an user-defined struct.
type BeaconChainProofsBalanceContainerProof struct {
	BalanceContainerRoot [32]byte
	Proof                []byte
}

// BeaconChainProofsBalanceProof is an auto generated low-level Go binding around an user-defined struct.
type BeaconChainProofsBalanceProof struct {
	PubkeyHash  [32]byte
	BalanceRoot [32]byte
	Proof       []byte
}

// BeaconChainProofsStateRootProof is an auto generated low-level Go binding around an user-defined struct.
type BeaconChainProofsStateRootProof struct {
	BeaconStateRoot [32]byte
	Proof           []byte
}

// BeaconChainProofsValidatorProof is an auto generated low-level Go binding around an user-defined struct.
type BeaconChainProofsValidatorProof struct {
	ValidatorFields [][32]byte
	Proof           []byte
}

// IEigenPodCheckpoint is an auto generated low-level Go binding around an user-defined struct.
type IEigenPodCheckpoint struct {
	BeaconBlockRoot   [32]byte
	ProofsRemaining   *big.Int
	PodBalanceGwei    uint64
	BalanceDeltasGwei *big.Int
}

// IEigenPodValidatorInfo is an auto generated low-level Go binding around an user-defined struct.
type IEigenPodValidatorInfo struct {
	ValidatorIndex      uint64
	RestakedBalanceGwei uint64
	LastCheckpointedAt  uint64
	Status              uint8
}

// StdInvariantFuzzSelector is an auto generated low-level Go binding around an user-defined struct.
type StdInvariantFuzzSelector struct {
	Addr      common.Address
	Selectors [][4]byte
}

// EigenPodMetaData contains all meta data concerning the EigenPod contract.
var EigenPodMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_ethPOS\",\"type\":\"address\",\"internalType\":\"contractIETHPOSDeposit\"},{\"name\":\"_eigenPodManager\",\"type\":\"address\",\"internalType\":\"contractIEigenPodManager\"},{\"name\":\"_GENESIS_TIME\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"GENESIS_TIME\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"IS_TEST\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"activeValidatorCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkpointBalanceExitedGwei\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"currentCheckpoint\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIEigenPod.Checkpoint\",\"components\":[{\"name\":\"beaconBlockRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proofsRemaining\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"podBalanceGwei\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"balanceDeltasGwei\",\"type\":\"int128\",\"internalType\":\"int128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"currentCheckpointTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eigenPodManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEigenPodManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ethPOS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIETHPOSDeposit\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"failed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRestaked\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_podOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lastCheckpointTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"podOwner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"recoverTokens\",\"inputs\":[{\"name\":\"tokenList\",\"type\":\"address[]\",\"internalType\":\"contractIERC20[]\"},{\"name\":\"amountsToWithdraw\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stake\",\"inputs\":[{\"name\":\"pubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"depositDataRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"startCheckpoint\",\"inputs\":[{\"name\":\"revertIfNoBalance\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"targetArtifactSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifactSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validatorPubkeyHashToInfo\",\"inputs\":[{\"name\":\"validatorPubkeyHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIEigenPod.ValidatorInfo\",\"components\":[{\"name\":\"validatorIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"restakedBalanceGwei\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"lastCheckpointedAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumIEigenPod.VALIDATOR_STATUS\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validatorPubkeyToInfo\",\"inputs\":[{\"name\":\"validatorPubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIEigenPod.ValidatorInfo\",\"components\":[{\"name\":\"validatorIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"restakedBalanceGwei\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"lastCheckpointedAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumIEigenPod.VALIDATOR_STATUS\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validatorStatus\",\"inputs\":[{\"name\":\"validatorPubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIEigenPod.VALIDATOR_STATUS\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validatorStatus\",\"inputs\":[{\"name\":\"pubkeyHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIEigenPod.VALIDATOR_STATUS\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyCheckpointProofs\",\"inputs\":[{\"name\":\"balanceContainerProof\",\"type\":\"tuple\",\"internalType\":\"structBeaconChainProofs.BalanceContainerProof\",\"components\":[{\"name\":\"balanceContainerRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"proofs\",\"type\":\"tuple[]\",\"internalType\":\"structBeaconChainProofs.BalanceProof[]\",\"components\":[{\"name\":\"pubkeyHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"balanceRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyStaleBalance\",\"inputs\":[{\"name\":\"beaconTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"stateRootProof\",\"type\":\"tuple\",\"internalType\":\"structBeaconChainProofs.StateRootProof\",\"components\":[{\"name\":\"beaconStateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"proof\",\"type\":\"tuple\",\"internalType\":\"structBeaconChainProofs.ValidatorProof\",\"components\":[{\"name\":\"validatorFields\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyWithdrawalCredentials\",\"inputs\":[{\"name\":\"beaconTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"stateRootProof\",\"type\":\"tuple\",\"internalType\":\"structBeaconChainProofs.StateRootProof\",\"components\":[{\"name\":\"beaconStateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"validatorIndices\",\"type\":\"uint40[]\",\"internalType\":\"uint40[]\"},{\"name\":\"validatorFieldsProofs\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"validatorFields\",\"type\":\"bytes32[][]\",\"internalType\":\"bytes32[][]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawRestakedBeaconChainETH\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amountWei\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawableRestakedExecutionLayerGwei\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"CheckpointCreated\",\"inputs\":[{\"name\":\"checkpointTimestamp\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"beaconBlockRoot\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CheckpointFinalized\",\"inputs\":[{\"name\":\"checkpointTimestamp\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"totalShareDeltaWei\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EigenPodStaked\",\"inputs\":[{\"name\":\"pubkey\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NonBeaconChainETHReceived\",\"inputs\":[{\"name\":\"amountReceived\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RestakedBeaconChainETHWithdrawn\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RestakingActivated\",\"inputs\":[{\"name\":\"podOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorBalanceUpdated\",\"inputs\":[{\"name\":\"validatorIndex\",\"type\":\"uint40\",\"indexed\":false,\"internalType\":\"uint40\"},{\"name\":\"balanceTimestamp\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"newValidatorBalanceGwei\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorCheckpointed\",\"inputs\":[{\"name\":\"checkpointTimestamp\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"validatorIndex\",\"type\":\"uint40\",\"indexed\":true,\"internalType\":\"uint40\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorRestaked\",\"inputs\":[{\"name\":\"validatorIndex\",\"type\":\"uint40\",\"indexed\":false,\"internalType\":\"uint40\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorWithdrawn\",\"inputs\":[{\"name\":\"checkpointTimestamp\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"validatorIndex\",\"type\":\"uint40\",\"indexed\":true,\"internalType\":\"uint40\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_address\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes32\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_int\",\"inputs\":[{\"name\":\"\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_address\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes32\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_string\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_string\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_uint\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"logs\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false}]",
	Bin: "0x60e06040526001606360006101000a81548160ff0219169083151502179055506001606760006101000a81548160ff0219169083151502179055503480156200004757600080fd5b5060405162008ca938038062008ca983398181016040528101906200006d9190620002d9565b8273ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff16815250508173ffffffffffffffffffffffffffffffffffffffff1660a08173ffffffffffffffffffffffffffffffffffffffff16815250508067ffffffffffffffff1660c08167ffffffffffffffff1681525050620001016200010a60201b60201c565b50505062000419565b600060019054906101000a900460ff16156200015d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200015490620003bc565b60405180910390fd5b60ff801660008054906101000a900460ff1660ff161015620001cf5760ff6000806101000a81548160ff021916908360ff1602179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249860ff604051620001c69190620003fc565b60405180910390a15b565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006200020382620001d6565b9050919050565b60006200021782620001f6565b9050919050565b62000229816200020a565b81146200023557600080fd5b50565b60008151905062000249816200021e565b92915050565b60006200025c82620001f6565b9050919050565b6200026e816200024f565b81146200027a57600080fd5b50565b6000815190506200028e8162000263565b92915050565b600067ffffffffffffffff82169050919050565b620002b38162000294565b8114620002bf57600080fd5b50565b600081519050620002d381620002a8565b92915050565b600080600060608486031215620002f557620002f4620001d1565b5b6000620003058682870162000238565b935050602062000318868287016200027d565b92505060406200032b86828701620002c2565b9150509250925092565b600082825260208201905092915050565b7f496e697469616c697a61626c653a20636f6e747261637420697320696e69746960008201527f616c697a696e6700000000000000000000000000000000000000000000000000602082015250565b6000620003a460278362000335565b9150620003b18262000346565b604082019050919050565b60006020820190508181036000830152620003d78162000395565b9050919050565b600060ff82169050919050565b620003f681620003de565b82525050565b6000602082019050620004136000830184620003eb565b92915050565b60805160a05160c051618813620004966000396000612da7015260008181610801015281816108dd01528181610ece015281816111b6015281816113190152818161188201528181611b4d015281816120ce015281816125d10152818161281e015261401b0152600081816116f30152611c2601526188136000f3fe6080604052600436106101f25760003560e01c80637439841f1161010d578063ba414fa6116100a0578063e20c9f711161006f578063e20c9f7114610728578063ee94d67c14610753578063f074ba621461077e578063f2882461146107a7578063fa7626d4146107d257610230565b8063ba414fa614610682578063c4907442146106ad578063c4d66de8146106d6578063dda3346c146106ff57610230565b8063916a17c6116100dc578063916a17c6146105d35780639b4e4634146105fe578063b522538a1461061a578063b5508aa91461065757610230565b80637439841f1461051757806374cdd7981461055457806385226c811461057f57806388676cad146105aa57610230565b80633f7286f41161018557806352396a591161015457806352396a591461043557806358eaee791461047257806366d9a9a0146104af5780636fcd0e53146104da57610230565b80633f7286f41461038957806342ecff2a146103b45780634665bcda146103df57806347d283721461040a57610230565b80633106ab53116101c15780633106ab53146102df5780633474aa161461030a5780633e5e3c23146103355780633f65cf191461036057610230565b8063039157d2146102355780630b18ff661461025e5780631ed7831c146102895780632340e8d3146102b457610230565b36610230577f6fdd3dbdb173299608c0aa9f368735857c8842b581f8389238bf05bd04b3bf49346040516102269190614ca6565b60405180910390a1005b600080fd5b34801561024157600080fd5b5061025c60048036038101906102579190614d58565b6107fd565b005b34801561026a57600080fd5b50610273610cc5565b6040516102809190614e24565b60405180910390f35b34801561029557600080fd5b5061029e610ceb565b6040516102ab9190614efd565b60405180910390f35b3480156102c057600080fd5b506102c9610d79565b6040516102d69190614ca6565b60405180910390f35b3480156102eb57600080fd5b506102f4610d7f565b6040516103019190614f3a565b60405180910390f35b34801561031657600080fd5b5061031f610d92565b60405161032c9190614f64565b60405180910390f35b34801561034157600080fd5b5061034a610dac565b6040516103579190614efd565b60405180910390f35b34801561036c57600080fd5b5061038760048036038101906103829190615090565b610e3a565b005b34801561039557600080fd5b5061039e61126f565b6040516103ab9190614efd565b60405180910390f35b3480156103c057600080fd5b506103c96112fd565b6040516103d69190614f64565b60405180910390f35b3480156103eb57600080fd5b506103f4611317565b60405161040191906151e7565b60405180910390f35b34801561041657600080fd5b5061041f61133b565b60405161042c91906152b9565b60405180910390f35b34801561044157600080fd5b5061045c600480360381019061045791906152d4565b6113d1565b6040516104699190614f64565b60405180910390f35b34801561047e57600080fd5b5061049960048036038101906104949190615357565b6113f8565b6040516104a6919061541b565b60405180910390f35b3480156104bb57600080fd5b506104c4611477565b6040516104d1919061561f565b60405180910390f35b3480156104e657600080fd5b5061050160048036038101906104fc919061566d565b6115c6565b60405161050e91906156fe565b60405180910390f35b34801561052357600080fd5b5061053e6004803603810190610539919061566d565b6116c4565b60405161054b919061541b565b60405180910390f35b34801561056057600080fd5b506105696116f1565b604051610576919061573a565b60405180910390f35b34801561058b57600080fd5b50610594611715565b6040516105a191906158b0565b60405180910390f35b3480156105b657600080fd5b506105d160048036038101906105cc91906158fe565b6117ee565b005b3480156105df57600080fd5b506105e86119fc565b6040516105f5919061561f565b60405180910390f35b6106186004803603810190610613919061592b565b611b4b565b005b34801561062657600080fd5b50610641600480360381019061063c9190615357565b611d0a565b60405161064e91906156fe565b60405180910390f35b34801561066357600080fd5b5061066c611e55565b60405161067991906158b0565b60405180910390f35b34801561068e57600080fd5b50610697611f2e565b6040516106a49190614f3a565b60405180910390f35b3480156106b957600080fd5b506106d460048036038101906106cf9190615a18565b6120cc565b005b3480156106e257600080fd5b506106fd60048036038101906106f89190615a58565b6122d5565b005b34801561070b57600080fd5b5061072660048036038101906107219190615cc4565b61253d565b005b34801561073457600080fd5b5061073d612772565b60405161074a9190614efd565b60405180910390f35b34801561075f57600080fd5b50610768612800565b6040516107759190614f64565b60405180910390f35b34801561078a57600080fd5b506107a560048036038101906107a09190615dc4565b61281a565b005b3480156107b357600080fd5b506107bc612da5565b6040516107c99190614f64565b60405180910390f35b3480156107de57600080fd5b506107e7612dc9565b6040516107f49190614f3a565b60405180910390f35b60067f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16635ac86ab7826040518263ffffffff1660e01b81526004016108589190615e5c565b602060405180830381865afa158015610875573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108999190615e8c565b156108d9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108d090615f3c565b60405180910390fd5b60087f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16635ac86ab7826040518263ffffffff1660e01b81526004016109349190615e5c565b602060405180830381865afa158015610951573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109759190615e8c565b156109b5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109ac90615f3c565b60405180910390fd5b6000610a0f8480600001906109ca9190615f6b565b80806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f82011690508083019250505050505050612ddc565b90506000603660008381526020019081526020016000206040518060800160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900460ff166002811115610aec57610aeb6153a4565b5b6002811115610afe57610afd6153a4565b5b81525050905062127500816040015167ffffffffffffffff16610b219190615ffd565b8767ffffffffffffffff1611610b6c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b63906160c5565b60405180910390fd5b60016002811115610b8057610b7f6153a4565b5b81606001516002811115610b9757610b966153a4565b5b14610bd7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bce90616157565b60405180910390fd5b610c2f858060000190610bea9190615f6b565b80806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f82011690508083019250505050505050612e01565b610c6e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c659061620f565b60405180910390fd5b610c80610c7a88612e2c565b87612f9c565b610cb28660000135868060000190610c989190615f6b565b888060200190610ca8919061622f565b86600001516130a2565b610cbc6000613258565b50505050505050565b603360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60606070805480602002602001604051908101604052809291908181526020018280548015610d6f57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610d25575b5050505050905090565b60395481565b603460089054906101000a900460ff1681565b603460009054906101000a900467ffffffffffffffff1681565b60606072805480602002602001604051908101604052809291908181526020018280548015610e3057602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610de6575b5050505050905090565b603360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610eca576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ec190616304565b60405180910390fd5b60027f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16635ac86ab7826040518263ffffffff1660e01b8152600401610f259190615e5c565b602060405180830381865afa158015610f42573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f669190615e8c565b15610fa6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f9d90615f3c565b60405180910390fd5b603460089054906101000a900460ff16610ff5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fec90616396565b60405180910390fd5b848490508787905014801561100f57508282905085859050145b61104e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110459061644e565b60405180910390fd5b603a60009054906101000a900467ffffffffffffffff1667ffffffffffffffff168967ffffffffffffffff161180156110b05750603a60089054906101000a900467ffffffffffffffff1667ffffffffffffffff168967ffffffffffffffff16115b6110ef576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110e690616506565b60405180910390fd5b6111016110fb8a612e2c565b89612f9c565b600080600090505b888890508110156111b3576111938a600001358a8a8481811061112f5761112e616526565b5b90506020020160208101906111449190616592565b89898581811061115757611156616526565b5b9050602002810190611169919061622f565b89898781811061117c5761117b616526565b5b905060200281019061118e9190615f6b565b613469565b8261119e9190615ffd565b915080806111ab906165bf565b915050611109565b507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663c2c51c40603360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836040518363ffffffff1660e01b8152600401611231929190616621565b600060405180830381600087803b15801561124b57600080fd5b505af115801561125f573d6000803e3d6000fd5b5050505050505050505050505050565b606060718054806020026020016040519081016040528092919081815260200182805480156112f357602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116112a9575b5050505050905090565b603a60089054906101000a900467ffffffffffffffff1681565b7f000000000000000000000000000000000000000000000000000000000000000081565b611343614bf8565b603c604051806080016040529081600082015481526020016001820160009054906101000a900462ffffff1662ffffff1662ffffff1681526020016001820160039054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200160018201600b9054906101000a9004600f0b600f0b600f0b81525050905090565b603b6020528060005260406000206000915054906101000a900467ffffffffffffffff1681565b60008061144884848080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050613a5c565b90506036600082815260200190815260200160002060000160189054906101000a900460ff1691505092915050565b60606075805480602002602001604051908101604052809291908181526020016000905b828210156115bd57838290600052602060002090600202016040518060400160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182018054806020026020016040519081016040528092919081815260200182805480156115a557602002820191906000526020600020906000905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190600401906020826003010492830192600103820291508084116115525790505b5050505050815250508152602001906001019061149b565b50505050905090565b6115ce614c35565b603660008381526020019081526020016000206040518060800160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900460ff1660028111156116a7576116a66153a4565b5b60028111156116b9576116b86153a4565b5b815250509050919050565b60006036600083815260200190815260200160002060000160189054906101000a900460ff169050919050565b7f000000000000000000000000000000000000000000000000000000000000000081565b60606074805480602002602001604051908101604052809291908181526020016000905b828210156117e557838290600052602060002001805461175890616679565b80601f016020809104026020016040519081016040528092919081815260200182805461178490616679565b80156117d15780601f106117a6576101008083540402835291602001916117d1565b820191906000526020600020905b8154815290600101906020018083116117b457829003601f168201915b505050505081526020019060010190611739565b50505050905090565b603360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461187e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161187590616304565b60405180910390fd5b60067f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16635ac86ab7826040518263ffffffff1660e01b81526004016118d99190615e5c565b602060405180830381865afa1580156118f6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061191a9190615e8c565b1561195a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161195190615f3c565b60405180910390fd5b61196382613258565b603460089054906101000a900460ff166119f8576001603460086101000a81548160ff021916908315150217905550603360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167fca8dfc8c5e0a67a74501c072a3325f685259bebbae7cfd230ab85198a78b70cd60405160405180910390a25b5050565b60606076805480602002602001604051908101604052809291908181526020016000905b82821015611b4257838290600052602060002090600202016040518060400160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160018201805480602002602001604051908101604052809291908181526020018280548015611b2a57602002820191906000526020600020906000905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681526020019060040190602082600301049283019260010382029150808411611ad75790505b50505050508152505081526020019060010190611a20565b50505050905090565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611bd9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611bd09061671d565b60405180910390fd5b6801bc16d674ec8000003414611c24576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611c1b906167d5565b60405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663228951186801bc16d674ec8000008787611c74613b1e565b8888886040518863ffffffff1660e01b8152600401611c9896959493929190616895565b6000604051808303818588803b158015611cb157600080fd5b505af1158015611cc5573d6000803e3d6000fd5b50505050507f606865b7934a25d4aed43f6cdb426403353fa4b3009c4d228407474581b01e238585604051611cfb9291906168f3565b60405180910390a15050505050565b611d12614c35565b60366000611d6385858080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050613a5c565b81526020019081526020016000206040518060800160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900460ff166002811115611e3757611e366153a4565b5b6002811115611e4957611e486153a4565b5b81525050905092915050565b60606073805480602002602001604051908101604052809291908181526020016000905b82821015611f25578382906000526020600020018054611e9890616679565b80601f0160208091040260200160405190810160405280929190818152602001828054611ec490616679565b8015611f115780601f10611ee657610100808354040283529160200191611f11565b820191906000526020600020905b815481529060010190602001808311611ef457829003601f168201915b505050505081526020019060010190611e79565b50505050905090565b6000606360019054906101000a900460ff1615611f5c57606360019054906101000a900460ff1690506120c9565b6000611f66613b51565b156120c45760007f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c60601b60601c73ffffffffffffffffffffffffffffffffffffffff167f667f9d70ca411d70ead50d8d5c22070dafc36ad75f3dcf5e7237b22ade9aecc47f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c60601b60601c7f6661696c6564000000000000000000000000000000000000000000000000000060405160200161202a929190616917565b60405160208183030381529060405260405160200161204a92919061699d565b60405160208183030381529060405260405161206691906169c5565b6000604051808303816000865af19150503d80600081146120a3576040519150601f19603f3d011682016040523d82523d6000602084013e6120a8565b606091505b50915050808060200190518101906120c09190615e8c565b9150505b809150505b90565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461215a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016121519061671d565b60405180910390fd5b6000633b9aca008261216c9190616a0b565b146121ac576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016121a390616ad4565b60405180910390fd5b6000633b9aca00826121be9190616af4565b9050603460009054906101000a900467ffffffffffffffff1667ffffffffffffffff168167ffffffffffffffff16111561222d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161222490616be3565b60405180910390fd5b80603460008282829054906101000a900467ffffffffffffffff166122529190616c03565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055508273ffffffffffffffffffffffffffffffffffffffff167f8947fd2ce07ef9cc302c4e8f0461015615d91ce851564839e91cc804c2f49d8e836040516122be9190614ca6565b60405180910390a26122d08383613b7a565b505050565b60008060019054906101000a900460ff161590508080156123065750600160008054906101000a900460ff1660ff16105b80612333575061231530613b88565b1580156123325750600160008054906101000a900460ff1660ff16145b5b612372576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161236990616ca9565b60405180910390fd5b60016000806101000a81548160ff021916908360ff16021790555080156123af576001600060016101000a81548160ff0219169083151502179055505b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141561241f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161241690616d3b565b60405180910390fd5b81603360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506001603460086101000a81548160ff021916908315150217905550603360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167fca8dfc8c5e0a67a74501c072a3325f685259bebbae7cfd230ab85198a78b70cd60405160405180910390a280156125395760008060016101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249860016040516125309190616d96565b60405180910390a15b5050565b603360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146125cd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016125c490616304565b60405180910390fd5b60057f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16635ac86ab7826040518263ffffffff1660e01b81526004016126289190615e5c565b602060405180830381865afa158015612645573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126699190615e8c565b156126a9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016126a090615f3c565b60405180910390fd5b82518451146126ed576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016126e490616e49565b60405180910390fd5b60005b845181101561276b57612758838583815181106127105761270f616526565b5b602002602001015187848151811061272b5761272a616526565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff16613bab9092919063ffffffff16565b8080612763906165bf565b9150506126f0565b5050505050565b6060606f8054806020026020016040519081016040528092919081815260200182805480156127f657602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116127ac575b5050505050905090565b603a60009054906101000a900467ffffffffffffffff1681565b60077f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16635ac86ab7826040518263ffffffff1660e01b81526004016128759190615e5c565b602060405180830381865afa158015612892573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906128b69190615e8c565b156128f6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016128ed90615f3c565b60405180910390fd5b6000603a60089054906101000a900467ffffffffffffffff16905060008167ffffffffffffffff16141561295f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161295690616f01565b60405180910390fd5b6000603c604051806080016040529081600082015481526020016001820160009054906101000a900462ffffff1662ffffff1662ffffff1681526020016001820160039054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200160018201600b9054906101000a9004600f0b600f0b600f0b8152505090506129fa816000015187613c31565b600080600090505b86869050811015612d225736878783818110612a2157612a20616526565b5b9050602002810190612a339190616f21565b9050600060366000836000013581526020019081526020016000206040518060800160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900460ff166002811115612b1457612b136153a4565b5b6002811115612b2657612b256153a4565b5b81525050905060016002811115612b4057612b3f6153a4565b5b81606001516002811115612b5757612b566153a4565b5b14612b63575050612d0f565b8567ffffffffffffffff16816040015167ffffffffffffffff1610612b89575050612d0f565b600080612b9c83898e6000013587613d50565b91509150866020018051809190612bb290616f49565b62ffffff1662ffffff16815250508187606001818151612bd29190616f73565b915090600f0b9081600f0b815250508086612bed9190616ff7565b955082603660008660000135815260200190815260200160002060008201518160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060208201518160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060408201518160000160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060608201518160000160186101000a81548160ff02191690836002811115612cbf57612cbe6153a4565b5b0217905550905050826000015164ffffffffff168867ffffffffffffffff167fa91c59033c3423e18b54d0acecebb4972f9ea95aedf5f4cae3b677b02eaf3a3f60405160405180910390a3505050505b8080612d1a906165bf565b915050612a02565b5080603b60008567ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060008282829054906101000a900467ffffffffffffffff16612d6d9190616ff7565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550612d9c82613ecb565b50505050505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b606360009054906101000a900460ff1681565b600081600081518110612df257612df1616526565b5b60200260200101519050919050565b60008060001b82600381518110612e1b57612e1a616526565b5b602002602001015114159050919050565b6000600c611fff612e3d9190617035565b8267ffffffffffffffff1642612e53919061708f565b10612e93576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612e8a90617135565b60405180910390fd5b600080720f3df6d732807ef1319fb7b8bb8522d0beac0273ffffffffffffffffffffffffffffffffffffffff1684604051602001612ed19190614f64565b604051602081830303815290604052604051612eed91906169c5565b600060405180830381855afa9150503d8060008114612f28576040519150601f19603f3d011682016040523d82523d6000602084013e612f2d565b606091505b5091509150818015612f40575060008151115b612f7f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612f76906171c7565b60405180910390fd5b80806020019051810190612f9391906171fc565b92505050919050565b60036020612faa9190617035565b818060200190612fba919061622f565b905014612ffc576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612ff39061729b565b60405180910390fd5b61305f81806020019061300f919061622f565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505083836000013560036141d0565b61309e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161309590617353565b60405180910390fd5b5050565b600885859050146130e8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016130df9061740b565b60405180910390fd5b6005600160286130f89190615ffd565b6131029190615ffd565b602061310e9190617035565b8383905014613152576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613149906174c3565b60405180910390fd5b600061319e868680806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f820116905080830192505050505050506141e9565b905060008264ffffffffff16600160286131b89190615ffd565b600b901b17905061320f85858080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508984846141d0565b61324e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161324590617555565b60405180910390fd5b5050505050505050565b6000603a60089054906101000a900467ffffffffffffffff1667ffffffffffffffff16146132bb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016132b29061760d565b60405180910390fd5b4267ffffffffffffffff16603a60009054906101000a900467ffffffffffffffff1667ffffffffffffffff161415613328576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161331f9061769f565b60405180910390fd5b6000603460009054906101000a900467ffffffffffffffff16633b9aca00476133519190616af4565b61335b9190616c03565b9050818015613374575060008167ffffffffffffffff16145b156133b4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016133ab90617731565b60405180910390fd5b600060405180608001604052806133ca42612e2c565b815260200160395462ffffff1681526020018367ffffffffffffffff1681526020016000600f0b815250905042603a60086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555061342881613ecb565b80600001514267ffffffffffffffff167f17a875d60c621f9c2f4d68122fe3a5bd359233b6fd8fcb79a0f65a6e433c6bf660405160405180910390a3505050565b6000806134b6848480806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f82011690508083019250505050505050612ddc565b90506000603660008381526020019081526020016000206040518060800160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900460ff166002811115613593576135926153a4565b5b60028111156135a5576135a46153a4565b5b815250509050600060028111156135bf576135be6153a4565b5b816060015160028111156135d6576135d56153a4565b5b14613616576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161360d9061780f565b60405180910390fd5b67ffffffffffffffff801661366b868680806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f820116905080830192505050505050506144a8565b67ffffffffffffffff16146136b5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016136ac906178c7565b60405180910390fd5b6136bd613b1e565b6136c690617919565b613710868680806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f820116905080830192505050505050506144d5565b14613750576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161374790617a18565b60405180910390fd5b6000858580806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f8201169050808301925050505050505073__$4050e451cf207daf724f57201af415e6ef$__634534711b90916040518263ffffffff1660e01b81526004016137cc9190617af6565b602060405180830381865af41580156137e9573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061380d9190617b2d565b905061381d8a87878b8b8e6130a2565b60396000815480929190613830906165bf565b9190505550600080603a60089054906101000a900467ffffffffffffffff1667ffffffffffffffff161461387a57603a60089054906101000a900467ffffffffffffffff16613892565b603a60009054906101000a900467ffffffffffffffff165b905060405180608001604052808b64ffffffffff1667ffffffffffffffff1681526020018367ffffffffffffffff1681526020018267ffffffffffffffff168152602001600160028111156138ea576138e96153a4565b5b8152506036600086815260200190815260200160002060008201518160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060208201518160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060408201518160000160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060608201518160000160186101000a81548160ff021916908360028111156139b8576139b76153a4565b5b02179055509050507f2d0800bbc377ea54a08c5db6a87aafff5e3e9c8fead0eda110e40e0c104414498a6040516139ef9190617b69565b60405180910390a17f0e5fac175b83177cc047381e030d8fb3b42b37bd1c025e22c280facad62c32df8a8284604051613a2a93929190617b84565b60405180910390a1633b9aca008267ffffffffffffffff16613a4c9190617035565b9450505050509695505050505050565b60006030825114613aa2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613a9990617c53565b60405180910390fd5b600282600060801b604051602001613abb929190617cc0565b604051602081830303815290604052604051613ad791906169c5565b602060405180830381855afa158015613af4573d6000803e3d6000fd5b5050506040513d601f19601f82011682018060405250810190613b1791906171fc565b9050919050565b6060600160f81b600060a81b30604051602001613b3d93929190617dca565b604051602081830303815290604052905090565b60008060009050737109709ecfa91a80626ff3989d68f67f5b1dd12d3b90506000811191505090565b613b8482826144fa565b5050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b613c2c8363a9059cbb60e01b8484604051602401613bca929190617e07565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506145ee565b505050565b60056003613c3f9190615ffd565b6020613c4b9190617035565b818060200190613c5b919061622f565b905014613c9d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613c9490617ec8565b60405180910390fd5b6000600c60056003901b179050613d0c828060200190613cbd919061622f565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050848460000135846141d0565b613d4b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613d4290617f80565b60405180910390fd5b505050565b6000806000866000015190506000876020015190506000613d728784886146b5565b90508167ffffffffffffffff168167ffffffffffffffff1614613dd757613d9981836147f4565b94507f0e5fac175b83177cc047381e030d8fb3b42b37bd1c025e22c280facad62c32df838983604051613dce93929190617b84565b60405180910390a15b80896020019067ffffffffffffffff16908167ffffffffffffffff168152505087896040019067ffffffffffffffff16908167ffffffffffffffff168152505060008167ffffffffffffffff161415613ebf5760396000815480929190613e3d90617fa0565b9190505550600289606001906002811115613e5b57613e5a6153a4565b5b90816002811115613e6f57613e6e6153a4565b5b8152505084613e7d90617fca565b93508264ffffffffff168867ffffffffffffffff167f2a02361ffa66cf2c2da4682c2355a6adcaa9f6c227b6e6563e68480f9587626a60405160405180910390a35b50505094509492505050565b6000816020015162ffffff161415614126576000633b9aca008260600151836040015167ffffffffffffffff16613f029190616f73565b600f0b613f0f9190618013565b90508160400151603460008282829054906101000a900467ffffffffffffffff16613f3a9190616ff7565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550603a60089054906101000a900467ffffffffffffffff16603a60006101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550603a60086101000a81549067ffffffffffffffff0219169055603c6000808201600090556001820160006101000a81549062ffffff02191690556001820160036101000a81549067ffffffffffffffff021916905560018201600b6101000a8154906fffffffffffffffffffffffffffffffff021916905550507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663c2c51c40603360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836040518363ffffffff1660e01b8152600401614096929190616621565b600060405180830381600087803b1580156140b057600080fd5b505af11580156140c4573d6000803e3d6000fd5b50505050603a60009054906101000a900467ffffffffffffffff1667ffffffffffffffff167f525408c201bc1576eb44116f6478f1c2a54775b19a043bcfdc708364f74f8e4482604051614118919061812a565b60405180910390a2506141cd565b80603c6000820151816000015560208201518160010160006101000a81548162ffffff021916908362ffffff16021790555060408201518160010160036101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550606082015181600101600b6101000a8154816fffffffffffffffffffffffffffffffff0219169083600f0b6fffffffffffffffffffffffffffffffff1602179055509050505b50565b6000836141de86858561481e565b149050949350505050565b600080600283516141fa9190616af4565b905060008167ffffffffffffffff81111561421857614217615a85565b5b6040519080825280602002602001820160405280156142465781602001602082028036833780820191505090505b50905060005b8281101561434f576002858260026142649190617035565b8151811061427557614274616526565b5b602002602001015186600184600261428d9190617035565b6142979190615ffd565b815181106142a8576142a7616526565b5b60200260200101516040516020016142c1929190618166565b6040516020818303038152906040526040516142dd91906169c5565b602060405180830381855afa1580156142fa573d6000803e3d6000fd5b5050506040513d601f19601f8201168201806040525081019061431d91906171fc565b8282815181106143305761432f616526565b5b6020026020010181815250508080614347906165bf565b91505061424c565b5060028261435d9190616af4565b91505b600082146144835760005b8281101561446e576002828260026143839190617035565b8151811061439457614393616526565b5b60200260200101518360018460026143ac9190617035565b6143b69190615ffd565b815181106143c7576143c6616526565b5b60200260200101516040516020016143e0929190618166565b6040516020818303038152906040526040516143fc91906169c5565b602060405180830381855afa158015614419573d6000803e3d6000fd5b5050506040513d601f19601f8201168201806040525081019061443c91906171fc565b82828151811061444f5761444e616526565b5b6020026020010181815250508080614466906165bf565b91505061436b565b5060028261447c9190616af4565b9150614360565b8060008151811061449757614496616526565b5b602002602001015192505050919050565b60006144ce826006815181106144c1576144c0616526565b5b6020026020010151614946565b9050919050565b6000816001815181106144eb576144ea616526565b5b60200260200101519050919050565b8047101561453d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614534906181de565b60405180910390fd5b60008273ffffffffffffffffffffffffffffffffffffffff168260405161456390618224565b60006040518083038185875af1925050503d80600081146145a0576040519150601f19603f3d011682016040523d82523d6000602084013e6145a5565b606091505b50509050806145e9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016145e0906182ab565b60405180910390fd5b505050565b6000614650826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff16614a029092919063ffffffff16565b90506000815111156146b057808060200190518101906146709190615e8c565b6146af576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016146a69061833d565b60405180910390fd5b5b505050565b6000600160266146c59190615ffd565b60206146d19190617035565b8280604001906146e1919061622f565b905014614723576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161471a906183f5565b60405180910390fd5b60006004846147329190618415565b64ffffffffff16905061479d83806040019061474e919061622f565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050868560200135846141d0565b6147dc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016147d3906184b8565b60405180910390fd5b6147ea836020013585614a1a565b9150509392505050565b60008167ffffffffffffffff168367ffffffffffffffff1661481691906184d8565b905092915050565b60008084511415801561483e575060006020855161483c9190616a0b565b145b61487d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614874906185f4565b60405180910390fd5b600060405180602001604052808581525090506000602090505b855181116149225760006002856148ae9190616a0b565b14156148e3578151600052808601516020526020826040600060026107d05a03fa6148d857600080fd5b60028404935061490e565b8086015160005281516020526020826040600060026107d05a03fa61490757600080fd5b6002840493505b60208161491b9190615ffd565b9050614897565b508060006001811061493757614936616526565b5b60200201519150509392505050565b600060c082901c60001c905060388160ff1667ffffffffffffffff16901b60288261ff001667ffffffffffffffff16901b60188362ff00001667ffffffffffffffff16901b60088463ff0000001667ffffffffffffffff16901b60088564ff000000001667ffffffffffffffff16901c60188665ff00000000001667ffffffffffffffff16901c60288766ff0000000000001667ffffffffffffffff16901c60388867ffffffffffffffff16901c171717171717179050919050565b6060614a118484600085614a5a565b90509392505050565b6000806040600484614a2c9190618614565b614a369190618645565b64ffffffffff169050614a51818560001c901b60001b614946565b91505092915050565b606082471015614a9f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614a96906186f6565b60405180910390fd5b614aa885614b6e565b614ae7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614ade90618762565b60405180910390fd5b6000808673ffffffffffffffffffffffffffffffffffffffff168587604051614b1091906169c5565b60006040518083038185875af1925050503d8060008114614b4d576040519150601f19603f3d011682016040523d82523d6000602084013e614b52565b606091505b5091509150614b62828286614b91565b92505050949350505050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b60608315614ba157829050614bf1565b600083511115614bb45782518084602001fd5b816040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614be891906187bb565b60405180910390fd5b9392505050565b604051806080016040528060008019168152602001600062ffffff168152602001600067ffffffffffffffff1681526020016000600f0b81525090565b6040518060800160405280600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff16815260200160006002811115614c8757614c866153a4565b5b81525090565b6000819050919050565b614ca081614c8d565b82525050565b6000602082019050614cbb6000830184614c97565b92915050565b6000604051905090565b600080fd5b600080fd5b600067ffffffffffffffff82169050919050565b614cf281614cd5565b8114614cfd57600080fd5b50565b600081359050614d0f81614ce9565b92915050565b600080fd5b600060408284031215614d3057614d2f614d15565b5b81905092915050565b600060408284031215614d4f57614d4e614d15565b5b81905092915050565b600080600060608486031215614d7157614d70614ccb565b5b6000614d7f86828701614d00565b935050602084013567ffffffffffffffff811115614da057614d9f614cd0565b5b614dac86828701614d1a565b925050604084013567ffffffffffffffff811115614dcd57614dcc614cd0565b5b614dd986828701614d39565b9150509250925092565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000614e0e82614de3565b9050919050565b614e1e81614e03565b82525050565b6000602082019050614e396000830184614e15565b92915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b614e7481614e03565b82525050565b6000614e868383614e6b565b60208301905092915050565b6000602082019050919050565b6000614eaa82614e3f565b614eb48185614e4a565b9350614ebf83614e5b565b8060005b83811015614ef0578151614ed78882614e7a565b9750614ee283614e92565b925050600181019050614ec3565b5085935050505092915050565b60006020820190508181036000830152614f178184614e9f565b905092915050565b60008115159050919050565b614f3481614f1f565b82525050565b6000602082019050614f4f6000830184614f2b565b92915050565b614f5e81614cd5565b82525050565b6000602082019050614f796000830184614f55565b92915050565b600080fd5b600080fd5b600080fd5b60008083601f840112614fa457614fa3614f7f565b5b8235905067ffffffffffffffff811115614fc157614fc0614f84565b5b602083019150836020820283011115614fdd57614fdc614f89565b5b9250929050565b60008083601f840112614ffa57614ff9614f7f565b5b8235905067ffffffffffffffff81111561501757615016614f84565b5b60208301915083602082028301111561503357615032614f89565b5b9250929050565b60008083601f8401126150505761504f614f7f565b5b8235905067ffffffffffffffff81111561506d5761506c614f84565b5b60208301915083602082028301111561508957615088614f89565b5b9250929050565b60008060008060008060008060a0898b0312156150b0576150af614ccb565b5b60006150be8b828c01614d00565b985050602089013567ffffffffffffffff8111156150df576150de614cd0565b5b6150eb8b828c01614d1a565b975050604089013567ffffffffffffffff81111561510c5761510b614cd0565b5b6151188b828c01614f8e565b9650965050606089013567ffffffffffffffff81111561513b5761513a614cd0565b5b6151478b828c01614fe4565b9450945050608089013567ffffffffffffffff81111561516a57615169614cd0565b5b6151768b828c0161503a565b92509250509295985092959890939650565b6000819050919050565b60006151ad6151a86151a384614de3565b615188565b614de3565b9050919050565b60006151bf82615192565b9050919050565b60006151d1826151b4565b9050919050565b6151e1816151c6565b82525050565b60006020820190506151fc60008301846151d8565b92915050565b6000819050919050565b61521581615202565b82525050565b600062ffffff82169050919050565b6152338161521b565b82525050565b61524281614cd5565b82525050565b600081600f0b9050919050565b61525e81615248565b82525050565b60808201600082015161527a600085018261520c565b50602082015161528d602085018261522a565b5060408201516152a06040850182615239565b5060608201516152b36060850182615255565b50505050565b60006080820190506152ce6000830184615264565b92915050565b6000602082840312156152ea576152e9614ccb565b5b60006152f884828501614d00565b91505092915050565b60008083601f84011261531757615316614f7f565b5b8235905067ffffffffffffffff81111561533457615333614f84565b5b6020830191508360018202830111156153505761534f614f89565b5b9250929050565b6000806020838503121561536e5761536d614ccb565b5b600083013567ffffffffffffffff81111561538c5761538b614cd0565b5b61539885828601615301565b92509250509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600381106153e4576153e36153a4565b5b50565b60008190506153f5826153d3565b919050565b6000615405826153e7565b9050919050565b615415816153fa565b82525050565b6000602082019050615430600083018461540c565b92915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b6154c38161548e565b82525050565b60006154d583836154ba565b60208301905092915050565b6000602082019050919050565b60006154f982615462565b615503818561546d565b935061550e8361547e565b8060005b8381101561553f57815161552688826154c9565b9750615531836154e1565b925050600181019050615512565b5085935050505092915050565b60006040830160008301516155646000860182614e6b565b506020830151848203602086015261557c82826154ee565b9150508091505092915050565b6000615595838361554c565b905092915050565b6000602082019050919050565b60006155b582615436565b6155bf8185615441565b9350836020820285016155d185615452565b8060005b8581101561560d57848403895281516155ee8582615589565b94506155f98361559d565b925060208a019950506001810190506155d5565b50829750879550505050505092915050565b6000602082019050818103600083015261563981846155aa565b905092915050565b61564a81615202565b811461565557600080fd5b50565b60008135905061566781615641565b92915050565b60006020828403121561568357615682614ccb565b5b600061569184828501615658565b91505092915050565b6156a3816153fa565b82525050565b6080820160008201516156bf6000850182615239565b5060208201516156d26020850182615239565b5060408201516156e56040850182615239565b5060608201516156f8606085018261569a565b50505050565b600060808201905061571360008301846156a9565b92915050565b6000615724826151b4565b9050919050565b61573481615719565b82525050565b600060208201905061574f600083018461572b565b92915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600081519050919050565b600082825260208201905092915050565b60005b838110156157bb5780820151818401526020810190506157a0565b838111156157ca576000848401525b50505050565b6000601f19601f8301169050919050565b60006157ec82615781565b6157f6818561578c565b935061580681856020860161579d565b61580f816157d0565b840191505092915050565b600061582683836157e1565b905092915050565b6000602082019050919050565b600061584682615755565b6158508185615760565b93508360208202850161586285615771565b8060005b8581101561589e578484038952815161587f858261581a565b945061588a8361582e565b925060208a01995050600181019050615866565b50829750879550505050505092915050565b600060208201905081810360008301526158ca818461583b565b905092915050565b6158db81614f1f565b81146158e657600080fd5b50565b6000813590506158f8816158d2565b92915050565b60006020828403121561591457615913614ccb565b5b6000615922848285016158e9565b91505092915050565b60008060008060006060868803121561594757615946614ccb565b5b600086013567ffffffffffffffff81111561596557615964614cd0565b5b61597188828901615301565b9550955050602086013567ffffffffffffffff81111561599457615993614cd0565b5b6159a088828901615301565b935093505060406159b388828901615658565b9150509295509295909350565b6159c981614e03565b81146159d457600080fd5b50565b6000813590506159e6816159c0565b92915050565b6159f581614c8d565b8114615a0057600080fd5b50565b600081359050615a12816159ec565b92915050565b60008060408385031215615a2f57615a2e614ccb565b5b6000615a3d858286016159d7565b9250506020615a4e85828601615a03565b9150509250929050565b600060208284031215615a6e57615a6d614ccb565b5b6000615a7c848285016159d7565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b615abd826157d0565b810181811067ffffffffffffffff82111715615adc57615adb615a85565b5b80604052505050565b6000615aef614cc1565b9050615afb8282615ab4565b919050565b600067ffffffffffffffff821115615b1b57615b1a615a85565b5b602082029050602081019050919050565b6000615b3782614e03565b9050919050565b615b4781615b2c565b8114615b5257600080fd5b50565b600081359050615b6481615b3e565b92915050565b6000615b7d615b7884615b00565b615ae5565b90508083825260208201905060208402830185811115615ba057615b9f614f89565b5b835b81811015615bc95780615bb58882615b55565b845260208401935050602081019050615ba2565b5050509392505050565b600082601f830112615be857615be7614f7f565b5b8135615bf8848260208601615b6a565b91505092915050565b600067ffffffffffffffff821115615c1c57615c1b615a85565b5b602082029050602081019050919050565b6000615c40615c3b84615c01565b615ae5565b90508083825260208201905060208402830185811115615c6357615c62614f89565b5b835b81811015615c8c5780615c788882615a03565b845260208401935050602081019050615c65565b5050509392505050565b600082601f830112615cab57615caa614f7f565b5b8135615cbb848260208601615c2d565b91505092915050565b600080600060608486031215615cdd57615cdc614ccb565b5b600084013567ffffffffffffffff811115615cfb57615cfa614cd0565b5b615d0786828701615bd3565b935050602084013567ffffffffffffffff811115615d2857615d27614cd0565b5b615d3486828701615c96565b9250506040615d45868287016159d7565b9150509250925092565b600060408284031215615d6557615d64614d15565b5b81905092915050565b60008083601f840112615d8457615d83614f7f565b5b8235905067ffffffffffffffff811115615da157615da0614f84565b5b602083019150836020820283011115615dbd57615dbc614f89565b5b9250929050565b600080600060408486031215615ddd57615ddc614ccb565b5b600084013567ffffffffffffffff811115615dfb57615dfa614cd0565b5b615e0786828701615d4f565b935050602084013567ffffffffffffffff811115615e2857615e27614cd0565b5b615e3486828701615d6e565b92509250509250925092565b600060ff82169050919050565b615e5681615e40565b82525050565b6000602082019050615e716000830184615e4d565b92915050565b600081519050615e86816158d2565b92915050565b600060208284031215615ea257615ea1614ccb565b5b6000615eb084828501615e77565b91505092915050565b600082825260208201905092915050565b7f456967656e506f642e6f6e6c795768656e4e6f745061757365643a20696e646560008201527f782069732070617573656420696e20456967656e506f644d616e616765720000602082015250565b6000615f26603e83615eb9565b9150615f3182615eca565b604082019050919050565b60006020820190508181036000830152615f5581615f19565b9050919050565b600080fd5b600080fd5b600080fd5b60008083356001602003843603038112615f8857615f87615f5c565b5b80840192508235915067ffffffffffffffff821115615faa57615fa9615f61565b5b602083019250602082023603831315615fc657615fc5615f66565b5b509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061600882614c8d565b915061601383614c8d565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111561604857616047615fce565b5b828201905092915050565b7f456967656e506f642e7665726966795374616c6542616c616e63653a2076616c60008201527f696461746f722062616c616e6365206973206e6f74207374616c652079657400602082015250565b60006160af603f83615eb9565b91506160ba82616053565b604082019050919050565b600060208201905081810360008301526160de816160a2565b9050919050565b7f456967656e506f642e7665726966795374616c6542616c616e63653a2076616c60008201527f696461746f72206973206e6f7420616374697665000000000000000000000000602082015250565b6000616141603483615eb9565b915061614c826160e5565b604082019050919050565b6000602082019050818103600083015261617081616134565b9050919050565b7f456967656e506f642e7665726966795374616c6542616c616e63653a2076616c60008201527f696461746f72206d75737420626520736c617368656420746f206265206d617260208201527f6b6564207374616c650000000000000000000000000000000000000000000000604082015250565b60006161f9604983615eb9565b915061620482616177565b606082019050919050565b60006020820190508181036000830152616228816161ec565b9050919050565b6000808335600160200384360303811261624c5761624b615f5c565b5b80840192508235915067ffffffffffffffff82111561626e5761626d615f61565b5b60208301925060018202360383131561628a57616289615f66565b5b509250929050565b7f456967656e506f642e6f6e6c79456967656e506f644f776e65723a206e6f742060008201527f706f644f776e6572000000000000000000000000000000000000000000000000602082015250565b60006162ee602883615eb9565b91506162f982616292565b604082019050919050565b6000602082019050818103600083015261631d816162e1565b9050919050565b7f456967656e506f642e7665726966795769746864726177616c43726564656e7460008201527f69616c733a2072657374616b696e67206e6f7420616374697665000000000000602082015250565b6000616380603a83615eb9565b915061638b82616324565b604082019050919050565b600060208201905081810360008301526163af81616373565b9050919050565b7f456967656e506f642e7665726966795769746864726177616c43726564656e7460008201527f69616c733a2076616c696461746f72496e646963657320616e642070726f6f6660208201527f73206d7573742062652073616d65206c656e6774680000000000000000000000604082015250565b6000616438605583615eb9565b9150616443826163b6565b606082019050919050565b600060208201905081810360008301526164678161642b565b9050919050565b7f456967656e506f642e7665726966795769746864726177616c43726564656e7460008201527f69616c733a207370656369666965642074696d657374616d7020697320746f6f60208201527f2066617220696e20706173740000000000000000000000000000000000000000604082015250565b60006164f0604c83615eb9565b91506164fb8261646e565b606082019050919050565b6000602082019050818103600083015261651f816164e3565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600064ffffffffff82169050919050565b61656f81616555565b811461657a57600080fd5b50565b60008135905061658c81616566565b92915050565b6000602082840312156165a8576165a7614ccb565b5b60006165b68482850161657d565b91505092915050565b60006165ca82614c8d565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8214156165fd576165fc615fce565b5b600182019050919050565b6000819050919050565b61661b81616608565b82525050565b60006040820190506166366000830185614e15565b6166436020830184616612565b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061669157607f821691505b602082108114156166a5576166a461664a565b5b50919050565b7f456967656e506f642e6f6e6c79456967656e506f644d616e616765723a206e6f60008201527f7420656967656e506f644d616e61676572000000000000000000000000000000602082015250565b6000616707603183615eb9565b9150616712826166ab565b604082019050919050565b60006020820190508181036000830152616736816166fa565b9050919050565b7f456967656e506f642e7374616b653a206d75737420696e697469616c6c79207360008201527f74616b6520666f7220616e792076616c696461746f722077697468203332206560208201527f7468657200000000000000000000000000000000000000000000000000000000604082015250565b60006167bf604483615eb9565b91506167ca8261673d565b606082019050919050565b600060208201905081810360008301526167ee816167b2565b9050919050565b600082825260208201905092915050565b82818337600083830152505050565b600061682183856167f5565b935061682e838584616806565b616837836157d0565b840190509392505050565b600081519050919050565b600061685882616842565b61686281856167f5565b935061687281856020860161579d565b61687b816157d0565b840191505092915050565b61688f81615202565b82525050565b600060808201905081810360008301526168b081888a616815565b905081810360208301526168c4818761684d565b905081810360408301526168d9818587616815565b90506168e86060830184616886565b979650505050505050565b6000602082019050818103600083015261690e818486616815565b90509392505050565b600060408201905061692c6000830185614e15565b6169396020830184616886565b9392505050565b6000819050919050565b61695b6169568261548e565b616940565b82525050565b600081905092915050565b600061697782616842565b6169818185616961565b935061699181856020860161579d565b80840191505092915050565b60006169a9828561694a565b6004820191506169b9828461696c565b91508190509392505050565b60006169d1828461696c565b915081905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000616a1682614c8d565b9150616a2183614c8d565b925082616a3157616a306169dc565b5b828206905092915050565b7f456967656e506f642e776974686472617752657374616b6564426561636f6e4360008201527f6861696e4554483a20616d6f756e74576569206d75737420626520612077686f60208201527f6c65204777656920616d6f756e74000000000000000000000000000000000000604082015250565b6000616abe604e83615eb9565b9150616ac982616a3c565b606082019050919050565b60006020820190508181036000830152616aed81616ab1565b9050919050565b6000616aff82614c8d565b9150616b0a83614c8d565b925082616b1a57616b196169dc565b5b828204905092915050565b7f456967656e506f642e776974686472617752657374616b6564426561636f6e4360008201527f6861696e4554483a20616d6f756e74477765692065786365656473207769746860208201527f6472617761626c6552657374616b6564457865637574696f6e4c61796572477760408201527f6569000000000000000000000000000000000000000000000000000000000000606082015250565b6000616bcd606283615eb9565b9150616bd882616b25565b608082019050919050565b60006020820190508181036000830152616bfc81616bc0565b9050919050565b6000616c0e82614cd5565b9150616c1983614cd5565b925082821015616c2c57616c2b615fce565b5b828203905092915050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b6000616c93602e83615eb9565b9150616c9e82616c37565b604082019050919050565b60006020820190508181036000830152616cc281616c86565b9050919050565b7f456967656e506f642e696e697469616c697a653a20706f644f776e657220636160008201527f6e6e6f74206265207a65726f2061646472657373000000000000000000000000602082015250565b6000616d25603483615eb9565b9150616d3082616cc9565b604082019050919050565b60006020820190508181036000830152616d5481616d18565b9050919050565b6000819050919050565b6000616d80616d7b616d7684616d5b565b615188565b615e40565b9050919050565b616d9081616d65565b82525050565b6000602082019050616dab6000830184616d87565b92915050565b7f456967656e506f642e7265636f766572546f6b656e733a20746f6b656e4c697360008201527f7420616e6420616d6f756e7473546f5769746864726177206d7573742062652060208201527f73616d65206c656e677468000000000000000000000000000000000000000000604082015250565b6000616e33604b83615eb9565b9150616e3e82616db1565b606082019050919050565b60006020820190508181036000830152616e6281616e26565b9050919050565b7f456967656e506f642e766572696679436865636b706f696e7450726f6f66733a60008201527f206d75737420686176652061637469766520636865636b706f696e7420746f2060208201527f706572666f726d20636865636b706f696e742070726f6f660000000000000000604082015250565b6000616eeb605883615eb9565b9150616ef682616e69565b606082019050919050565b60006020820190508181036000830152616f1a81616ede565b9050919050565b600082356001606003833603038112616f3d57616f3c615f5c565b5b80830191505092915050565b6000616f548261521b565b91506000821415616f6857616f67615fce565b5b600182039050919050565b6000616f7e82615248565b9150616f8983615248565b9250816f7fffffffffffffffffffffffffffffff03831360008312151615616fb457616fb3615fce565b5b817fffffffffffffffffffffffffffffffff80000000000000000000000000000000038312600083121615616fec57616feb615fce565b5b828201905092915050565b600061700282614cd5565b915061700d83614cd5565b92508267ffffffffffffffff0382111561702a57617029615fce565b5b828201905092915050565b600061704082614c8d565b915061704b83614c8d565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561708457617083615fce565b5b828202905092915050565b600061709a82614c8d565b91506170a583614c8d565b9250828210156170b8576170b7615fce565b5b828203905092915050565b7f456967656e506f642e5f676574506172656e74426c6f636b526f6f743a20746960008201527f6d657374616d70206f7574206f662072616e6765000000000000000000000000602082015250565b600061711f603483615eb9565b915061712a826170c3565b604082019050919050565b6000602082019050818103600083015261714e81617112565b9050919050565b7f456967656e506f642e5f676574506172656e74426c6f636b526f6f743a20696e60008201527f76616c696420626c6f636b20726f6f742072657475726e656400000000000000602082015250565b60006171b1603983615eb9565b91506171bc82617155565b604082019050919050565b600060208201905081810360008301526171e0816171a4565b9050919050565b6000815190506171f681615641565b92915050565b60006020828403121561721257617211614ccb565b5b6000617220848285016171e7565b91505092915050565b7f426561636f6e436861696e50726f6f66732e7665726966795374617465526f6f60008201527f743a2050726f6f662068617320696e636f7272656374206c656e677468000000602082015250565b6000617285603d83615eb9565b915061729082617229565b604082019050919050565b600060208201905081810360008301526172b481617278565b9050919050565b7f426561636f6e436861696e50726f6f66732e7665726966795374617465526f6f60008201527f743a20496e76616c696420737461746520726f6f74206d65726b6c652070726f60208201527f6f66000000000000000000000000000000000000000000000000000000000000604082015250565b600061733d604283615eb9565b9150617348826172bb565b606082019050919050565b6000602082019050818103600083015261736c81617330565b9050919050565b7f426561636f6e436861696e50726f6f66732e76657269667956616c696461746f60008201527f724669656c64733a2056616c696461746f72206669656c64732068617320696e60208201527f636f7272656374206c656e677468000000000000000000000000000000000000604082015250565b60006173f5604e83615eb9565b915061740082617373565b606082019050919050565b60006020820190508181036000830152617424816173e8565b9050919050565b7f426561636f6e436861696e50726f6f66732e76657269667956616c696461746f60008201527f724669656c64733a2050726f6f662068617320696e636f7272656374206c656e60208201527f6774680000000000000000000000000000000000000000000000000000000000604082015250565b60006174ad604383615eb9565b91506174b88261742b565b606082019050919050565b600060208201905081810360008301526174dc816174a0565b9050919050565b7f426561636f6e436861696e50726f6f66732e76657269667956616c696461746f60008201527f724669656c64733a20496e76616c6964206d65726b6c652070726f6f66000000602082015250565b600061753f603d83615eb9565b915061754a826174e3565b604082019050919050565b6000602082019050818103600083015261756e81617532565b9050919050565b7f456967656e506f642e5f7374617274436865636b706f696e743a206d7573742060008201527f66696e6973682070726576696f757320636865636b706f696e74206265666f7260208201527f65207374617274696e6720616e6f746865720000000000000000000000000000604082015250565b60006175f7605283615eb9565b915061760282617575565b606082019050919050565b60006020820190508181036000830152617626816175ea565b9050919050565b7f456967656e506f642e5f7374617274436865636b706f696e743a2063616e6e6f60008201527f7420636865636b706f696e7420747769636520696e206f6e6520626c6f636b00602082015250565b6000617689603f83615eb9565b91506176948261762d565b604082019050919050565b600060208201905081810360008301526176b88161767c565b9050919050565b7f456967656e506f642e5f7374617274436865636b706f696e743a206e6f20626160008201527f6c616e636520617661696c61626c6520746f20636865636b706f696e74000000602082015250565b600061771b603d83615eb9565b9150617726826176bf565b604082019050919050565b6000602082019050818103600083015261774a8161770e565b9050919050565b7f456967656e506f642e5f7665726966795769746864726177616c43726564656e60008201527f7469616c733a2076616c696461746f72206d75737420626520696e616374697660208201527f6520746f2070726f7665207769746864726177616c2063726564656e7469616c60408201527f7300000000000000000000000000000000000000000000000000000000000000606082015250565b60006177f9606183615eb9565b915061780482617751565b608082019050919050565b60006020820190508181036000830152617828816177ec565b9050919050565b7f456967656e506f642e5f7665726966795769746864726177616c43726564656e60008201527f7469616c733a2076616c696461746f72206d757374206e6f742062652065786960208201527f74696e6700000000000000000000000000000000000000000000000000000000604082015250565b60006178b1604483615eb9565b91506178bc8261782f565b606082019050919050565b600060208201905081810360008301526178e0816178a4565b9050919050565b6000819050602082019050919050565b60006179038251615202565b80915050919050565b600082821b905092915050565b600061792482616842565b8261792e846178e7565b9050617939816178f7565b92506020821015617979576179747fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8360200360080261790c565b831692505b5050919050565b7f456967656e506f642e5f7665726966795769746864726177616c43726564656e60008201527f7469616c733a2070726f6f66206973206e6f7420666f7220746869732045696760208201527f656e506f64000000000000000000000000000000000000000000000000000000604082015250565b6000617a02604583615eb9565b9150617a0d82617980565b606082019050919050565b60006020820190508181036000830152617a31816179f5565b9050919050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b617a6d81615202565b82525050565b6000617a7f8383617a64565b60208301905092915050565b6000602082019050919050565b6000617aa382617a38565b617aad8185617a43565b9350617ab883617a54565b8060005b83811015617ae9578151617ad08882617a73565b9750617adb83617a8b565b925050600181019050617abc565b5085935050505092915050565b60006020820190508181036000830152617b108184617a98565b905092915050565b600081519050617b2781614ce9565b92915050565b600060208284031215617b4357617b42614ccb565b5b6000617b5184828501617b18565b91505092915050565b617b6381616555565b82525050565b6000602082019050617b7e6000830184617b5a565b92915050565b6000606082019050617b996000830186617b5a565b617ba66020830185614f55565b617bb36040830184614f55565b949350505050565b7f456967656e506f642e5f63616c63756c61746556616c696461746f725075626b60008201527f657948617368206d75737420626520612034382d6279746520424c532070756260208201527f6c6963206b657900000000000000000000000000000000000000000000000000604082015250565b6000617c3d604783615eb9565b9150617c4882617bbb565b606082019050919050565b60006020820190508181036000830152617c6c81617c30565b9050919050565b60007fffffffffffffffffffffffffffffffff0000000000000000000000000000000082169050919050565b6000819050919050565b617cba617cb582617c73565b617c9f565b82525050565b6000617ccc828561696c565b9150617cd88284617ca9565b6010820191508190509392505050565b60007fff0000000000000000000000000000000000000000000000000000000000000082169050919050565b6000819050919050565b617d2f617d2a82617ce8565b617d14565b82525050565b60007fffffffffffffffffffffff00000000000000000000000000000000000000000082169050919050565b6000819050919050565b617d7c617d7782617d35565b617d61565b82525050565b60008160601b9050919050565b6000617d9a82617d82565b9050919050565b6000617dac82617d8f565b9050919050565b617dc4617dbf82614e03565b617da1565b82525050565b6000617dd68286617d1e565b600182019150617de68285617d6b565b600b82019150617df68284617db3565b601482019150819050949350505050565b6000604082019050617e1c6000830185614e15565b617e296020830184614c97565b9392505050565b7f426561636f6e436861696e50726f6f66732e76657269667942616c616e63654360008201527f6f6e7461696e65723a2050726f6f662068617320696e636f7272656374206c6560208201527f6e67746800000000000000000000000000000000000000000000000000000000604082015250565b6000617eb2604483615eb9565b9150617ebd82617e30565b606082019050919050565b60006020820190508181036000830152617ee181617ea5565b9050919050565b7f426561636f6e436861696e50726f6f66732e76657269667942616c616e63654360008201527f6f6e7461696e65723a20696e76616c69642062616c616e636520636f6e74616960208201527f6e65722070726f6f660000000000000000000000000000000000000000000000604082015250565b6000617f6a604983615eb9565b9150617f7582617ee8565b606082019050919050565b60006020820190508181036000830152617f9981617f5d565b9050919050565b6000617fab82614c8d565b91506000821415617fbf57617fbe615fce565b5b600182039050919050565b6000617fd582615248565b91507fffffffffffffffffffffffffffffffff8000000000000000000000000000000082141561800857618007615fce565b5b816000039050919050565b600061801e82616608565b915061802983616608565b9250827f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048211600084136000841316161561806857618067615fce565b5b817f800000000000000000000000000000000000000000000000000000000000000005831260008412600084131616156180a5576180a4615fce565b5b827f800000000000000000000000000000000000000000000000000000000000000005821260008413600084121616156180e2576180e1615fce565b5b827f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff058212600084126000841216161561811f5761811e615fce565b5b828202905092915050565b600060208201905061813f6000830184616612565b92915050565b6000819050919050565b61816061815b82615202565b618145565b82525050565b6000618172828561814f565b602082019150618182828461814f565b6020820191508190509392505050565b7f416464726573733a20696e73756666696369656e742062616c616e6365000000600082015250565b60006181c8601d83615eb9565b91506181d382618192565b602082019050919050565b600060208201905081810360008301526181f7816181bb565b9050919050565b50565b600061820e600083616961565b9150618219826181fe565b600082019050919050565b600061822f82618201565b9150819050919050565b7f416464726573733a20756e61626c6520746f2073656e642076616c75652c207260008201527f6563697069656e74206d61792068617665207265766572746564000000000000602082015250565b6000618295603a83615eb9565b91506182a082618239565b604082019050919050565b600060208201905081810360008301526182c481618288565b9050919050565b7f5361666545524332303a204552433230206f7065726174696f6e20646964206e60008201527f6f74207375636365656400000000000000000000000000000000000000000000602082015250565b6000618327602a83615eb9565b9150618332826182cb565b604082019050919050565b600060208201905081810360008301526183568161831a565b9050919050565b7f426561636f6e436861696e50726f6f66732e76657269667956616c696461746f60008201527f7242616c616e63653a2050726f6f662068617320696e636f7272656374206c6560208201527f6e67746800000000000000000000000000000000000000000000000000000000604082015250565b60006183df604483615eb9565b91506183ea8261835d565b606082019050919050565b6000602082019050818103600083015261840e816183d2565b9050919050565b600061842082616555565b915061842b83616555565b92508261843b5761843a6169dc565b5b828204905092915050565b7f426561636f6e436861696e50726f6f66732e76657269667956616c696461746f60008201527f7242616c616e63653a20496e76616c6964206d65726b6c652070726f6f660000602082015250565b60006184a2603e83615eb9565b91506184ad82618446565b604082019050919050565b600060208201905081810360008301526184d181618495565b9050919050565b60006184e382615248565b91506184ee83615248565b9250827fffffffffffffffffffffffffffffffff800000000000000000000000000000000182126000841215161561852957618528615fce565b5b826f7fffffffffffffffffffffffffffffff01821360008412161561855157618550615fce565b5b828203905092915050565b7f4d65726b6c652e70726f63657373496e636c7573696f6e50726f6f665368613260008201527f35363a2070726f6f66206c656e6774682073686f756c642062652061206e6f6e60208201527f2d7a65726f206d756c7469706c65206f66203332000000000000000000000000604082015250565b60006185de605483615eb9565b91506185e98261855c565b606082019050919050565b6000602082019050818103600083015261860d816185d1565b9050919050565b600061861f82616555565b915061862a83616555565b92508261863a576186396169dc565b5b828206905092915050565b600061865082616555565b915061865b83616555565b92508164ffffffffff048311821515161561867957618678615fce565b5b828202905092915050565b7f416464726573733a20696e73756666696369656e742062616c616e636520666f60008201527f722063616c6c0000000000000000000000000000000000000000000000000000602082015250565b60006186e0602683615eb9565b91506186eb82618684565b604082019050919050565b6000602082019050818103600083015261870f816186d3565b9050919050565b7f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000600082015250565b600061874c601d83615eb9565b915061875782618716565b602082019050919050565b6000602082019050818103600083015261877b8161873f565b9050919050565b600061878d82615781565b6187978185615eb9565b93506187a781856020860161579d565b6187b0816157d0565b840191505092915050565b600060208201905081810360008301526187d58184618782565b90509291505056fea264697066735822122036ef7b9c17387f9ef55f0f047e0b0164277839d16e944961eff0048845f4177a64736f6c634300080c0033",
}

// EigenPodABI is the input ABI used to generate the binding from.
// Deprecated: Use EigenPodMetaData.ABI instead.
var EigenPodABI = EigenPodMetaData.ABI

// EigenPodBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EigenPodMetaData.Bin instead.
var EigenPodBin = EigenPodMetaData.Bin

// DeployEigenPod deploys a new Ethereum contract, binding an instance of EigenPod to it.
func DeployEigenPod(auth *bind.TransactOpts, backend bind.ContractBackend, _ethPOS common.Address, _eigenPodManager common.Address, _GENESIS_TIME uint64) (common.Address, *types.Transaction, *EigenPod, error) {
	parsed, err := EigenPodMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EigenPodBin), backend, _ethPOS, _eigenPodManager, _GENESIS_TIME)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EigenPod{EigenPodCaller: EigenPodCaller{contract: contract}, EigenPodTransactor: EigenPodTransactor{contract: contract}, EigenPodFilterer: EigenPodFilterer{contract: contract}}, nil
}

// EigenPod is an auto generated Go binding around an Ethereum contract.
type EigenPod struct {
	EigenPodCaller     // Read-only binding to the contract
	EigenPodTransactor // Write-only binding to the contract
	EigenPodFilterer   // Log filterer for contract events
}

// EigenPodCaller is an auto generated read-only Go binding around an Ethereum contract.
type EigenPodCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EigenPodTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EigenPodTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EigenPodFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EigenPodFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EigenPodSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EigenPodSession struct {
	Contract     *EigenPod         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EigenPodCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EigenPodCallerSession struct {
	Contract *EigenPodCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// EigenPodTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EigenPodTransactorSession struct {
	Contract     *EigenPodTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// EigenPodRaw is an auto generated low-level Go binding around an Ethereum contract.
type EigenPodRaw struct {
	Contract *EigenPod // Generic contract binding to access the raw methods on
}

// EigenPodCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EigenPodCallerRaw struct {
	Contract *EigenPodCaller // Generic read-only contract binding to access the raw methods on
}

// EigenPodTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EigenPodTransactorRaw struct {
	Contract *EigenPodTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEigenPod creates a new instance of EigenPod, bound to a specific deployed contract.
func NewEigenPod(address common.Address, backend bind.ContractBackend) (*EigenPod, error) {
	contract, err := bindEigenPod(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EigenPod{EigenPodCaller: EigenPodCaller{contract: contract}, EigenPodTransactor: EigenPodTransactor{contract: contract}, EigenPodFilterer: EigenPodFilterer{contract: contract}}, nil
}

// NewEigenPodCaller creates a new read-only instance of EigenPod, bound to a specific deployed contract.
func NewEigenPodCaller(address common.Address, caller bind.ContractCaller) (*EigenPodCaller, error) {
	contract, err := bindEigenPod(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EigenPodCaller{contract: contract}, nil
}

// NewEigenPodTransactor creates a new write-only instance of EigenPod, bound to a specific deployed contract.
func NewEigenPodTransactor(address common.Address, transactor bind.ContractTransactor) (*EigenPodTransactor, error) {
	contract, err := bindEigenPod(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EigenPodTransactor{contract: contract}, nil
}

// NewEigenPodFilterer creates a new log filterer instance of EigenPod, bound to a specific deployed contract.
func NewEigenPodFilterer(address common.Address, filterer bind.ContractFilterer) (*EigenPodFilterer, error) {
	contract, err := bindEigenPod(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EigenPodFilterer{contract: contract}, nil
}

// bindEigenPod binds a generic wrapper to an already deployed contract.
func bindEigenPod(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EigenPodMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EigenPod *EigenPodRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EigenPod.Contract.EigenPodCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EigenPod *EigenPodRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EigenPod.Contract.EigenPodTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EigenPod *EigenPodRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EigenPod.Contract.EigenPodTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EigenPod *EigenPodCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EigenPod.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EigenPod *EigenPodTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EigenPod.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EigenPod *EigenPodTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EigenPod.Contract.contract.Transact(opts, method, params...)
}

// GENESISTIME is a free data retrieval call binding the contract method 0xf2882461.
//
// Solidity: function GENESIS_TIME() view returns(uint64)
func (_EigenPod *EigenPodCaller) GENESISTIME(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _EigenPod.contract.Call(opts, &out, "GENESIS_TIME")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GENESISTIME is a free data retrieval call binding the contract method 0xf2882461.
//
// Solidity: function GENESIS_TIME() view returns(uint64)
func (_EigenPod *EigenPodSession) GENESISTIME() (uint64, error) {
	return _EigenPod.Contract.GENESISTIME(&_EigenPod.CallOpts)
}

// GENESISTIME is a free data retrieval call binding the contract method 0xf2882461.
//
// Solidity: function GENESIS_TIME() view returns(uint64)
func (_EigenPod *EigenPodCallerSession) GENESISTIME() (uint64, error) {
	return _EigenPod.Contract.GENESISTIME(&_EigenPod.CallOpts)
}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_EigenPod *EigenPodCaller) ISTEST(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EigenPod.contract.Call(opts, &out, "IS_TEST")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_EigenPod *EigenPodSession) ISTEST() (bool, error) {
	return _EigenPod.Contract.ISTEST(&_EigenPod.CallOpts)
}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_EigenPod *EigenPodCallerSession) ISTEST() (bool, error) {
	return _EigenPod.Contract.ISTEST(&_EigenPod.CallOpts)
}

// ActiveValidatorCount is a free data retrieval call binding the contract method 0x2340e8d3.
//
// Solidity: function activeValidatorCount() view returns(uint256)
func (_EigenPod *EigenPodCaller) ActiveValidatorCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EigenPod.contract.Call(opts, &out, "activeValidatorCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActiveValidatorCount is a free data retrieval call binding the contract method 0x2340e8d3.
//
// Solidity: function activeValidatorCount() view returns(uint256)
func (_EigenPod *EigenPodSession) ActiveValidatorCount() (*big.Int, error) {
	return _EigenPod.Contract.ActiveValidatorCount(&_EigenPod.CallOpts)
}

// ActiveValidatorCount is a free data retrieval call binding the contract method 0x2340e8d3.
//
// Solidity: function activeValidatorCount() view returns(uint256)
func (_EigenPod *EigenPodCallerSession) ActiveValidatorCount() (*big.Int, error) {
	return _EigenPod.Contract.ActiveValidatorCount(&_EigenPod.CallOpts)
}

// CheckpointBalanceExitedGwei is a free data retrieval call binding the contract method 0x52396a59.
//
// Solidity: function checkpointBalanceExitedGwei(uint64 ) view returns(uint64)
func (_EigenPod *EigenPodCaller) CheckpointBalanceExitedGwei(opts *bind.CallOpts, arg0 uint64) (uint64, error) {
	var out []interface{}
	err := _EigenPod.contract.Call(opts, &out, "checkpointBalanceExitedGwei", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// CheckpointBalanceExitedGwei is a free data retrieval call binding the contract method 0x52396a59.
//
// Solidity: function checkpointBalanceExitedGwei(uint64 ) view returns(uint64)
func (_EigenPod *EigenPodSession) CheckpointBalanceExitedGwei(arg0 uint64) (uint64, error) {
	return _EigenPod.Contract.CheckpointBalanceExitedGwei(&_EigenPod.CallOpts, arg0)
}

// CheckpointBalanceExitedGwei is a free data retrieval call binding the contract method 0x52396a59.
//
// Solidity: function checkpointBalanceExitedGwei(uint64 ) view returns(uint64)
func (_EigenPod *EigenPodCallerSession) CheckpointBalanceExitedGwei(arg0 uint64) (uint64, error) {
	return _EigenPod.Contract.CheckpointBalanceExitedGwei(&_EigenPod.CallOpts, arg0)
}

// CurrentCheckpoint is a free data retrieval call binding the contract method 0x47d28372.
//
// Solidity: function currentCheckpoint() view returns((bytes32,uint24,uint64,int128))
func (_EigenPod *EigenPodCaller) CurrentCheckpoint(opts *bind.CallOpts) (IEigenPodCheckpoint, error) {
	var out []interface{}
	err := _EigenPod.contract.Call(opts, &out, "currentCheckpoint")

	if err != nil {
		return *new(IEigenPodCheckpoint), err
	}

	out0 := *abi.ConvertType(out[0], new(IEigenPodCheckpoint)).(*IEigenPodCheckpoint)

	return out0, err

}

// CurrentCheckpoint is a free data retrieval call binding the contract method 0x47d28372.
//
// Solidity: function currentCheckpoint() view returns((bytes32,uint24,uint64,int128))
func (_EigenPod *EigenPodSession) CurrentCheckpoint() (IEigenPodCheckpoint, error) {
	return _EigenPod.Contract.CurrentCheckpoint(&_EigenPod.CallOpts)
}

// CurrentCheckpoint is a free data retrieval call binding the contract method 0x47d28372.
//
// Solidity: function currentCheckpoint() view returns((bytes32,uint24,uint64,int128))
func (_EigenPod *EigenPodCallerSession) CurrentCheckpoint() (IEigenPodCheckpoint, error) {
	return _EigenPod.Contract.CurrentCheckpoint(&_EigenPod.CallOpts)
}

// CurrentCheckpointTimestamp is a free data retrieval call binding the contract method 0x42ecff2a.
//
// Solidity: function currentCheckpointTimestamp() view returns(uint64)
func (_EigenPod *EigenPodCaller) CurrentCheckpointTimestamp(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _EigenPod.contract.Call(opts, &out, "currentCheckpointTimestamp")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// CurrentCheckpointTimestamp is a free data retrieval call binding the contract method 0x42ecff2a.
//
// Solidity: function currentCheckpointTimestamp() view returns(uint64)
func (_EigenPod *EigenPodSession) CurrentCheckpointTimestamp() (uint64, error) {
	return _EigenPod.Contract.CurrentCheckpointTimestamp(&_EigenPod.CallOpts)
}

// CurrentCheckpointTimestamp is a free data retrieval call binding the contract method 0x42ecff2a.
//
// Solidity: function currentCheckpointTimestamp() view returns(uint64)
func (_EigenPod *EigenPodCallerSession) CurrentCheckpointTimestamp() (uint64, error) {
	return _EigenPod.Contract.CurrentCheckpointTimestamp(&_EigenPod.CallOpts)
}

// EigenPodManager is a free data retrieval call binding the contract method 0x4665bcda.
//
// Solidity: function eigenPodManager() view returns(address)
func (_EigenPod *EigenPodCaller) EigenPodManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EigenPod.contract.Call(opts, &out, "eigenPodManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EigenPodManager is a free data retrieval call binding the contract method 0x4665bcda.
//
// Solidity: function eigenPodManager() view returns(address)
func (_EigenPod *EigenPodSession) EigenPodManager() (common.Address, error) {
	return _EigenPod.Contract.EigenPodManager(&_EigenPod.CallOpts)
}

// EigenPodManager is a free data retrieval call binding the contract method 0x4665bcda.
//
// Solidity: function eigenPodManager() view returns(address)
func (_EigenPod *EigenPodCallerSession) EigenPodManager() (common.Address, error) {
	return _EigenPod.Contract.EigenPodManager(&_EigenPod.CallOpts)
}

// EthPOS is a free data retrieval call binding the contract method 0x74cdd798.
//
// Solidity: function ethPOS() view returns(address)
func (_EigenPod *EigenPodCaller) EthPOS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EigenPod.contract.Call(opts, &out, "ethPOS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EthPOS is a free data retrieval call binding the contract method 0x74cdd798.
//
// Solidity: function ethPOS() view returns(address)
func (_EigenPod *EigenPodSession) EthPOS() (common.Address, error) {
	return _EigenPod.Contract.EthPOS(&_EigenPod.CallOpts)
}

// EthPOS is a free data retrieval call binding the contract method 0x74cdd798.
//
// Solidity: function ethPOS() view returns(address)
func (_EigenPod *EigenPodCallerSession) EthPOS() (common.Address, error) {
	return _EigenPod.Contract.EthPOS(&_EigenPod.CallOpts)
}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_EigenPod *EigenPodCaller) ExcludeArtifacts(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _EigenPod.contract.Call(opts, &out, "excludeArtifacts")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_EigenPod *EigenPodSession) ExcludeArtifacts() ([]string, error) {
	return _EigenPod.Contract.ExcludeArtifacts(&_EigenPod.CallOpts)
}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_EigenPod *EigenPodCallerSession) ExcludeArtifacts() ([]string, error) {
	return _EigenPod.Contract.ExcludeArtifacts(&_EigenPod.CallOpts)
}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_EigenPod *EigenPodCaller) ExcludeContracts(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EigenPod.contract.Call(opts, &out, "excludeContracts")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_EigenPod *EigenPodSession) ExcludeContracts() ([]common.Address, error) {
	return _EigenPod.Contract.ExcludeContracts(&_EigenPod.CallOpts)
}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_EigenPod *EigenPodCallerSession) ExcludeContracts() ([]common.Address, error) {
	return _EigenPod.Contract.ExcludeContracts(&_EigenPod.CallOpts)
}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_EigenPod *EigenPodCaller) ExcludeSenders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EigenPod.contract.Call(opts, &out, "excludeSenders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_EigenPod *EigenPodSession) ExcludeSenders() ([]common.Address, error) {
	return _EigenPod.Contract.ExcludeSenders(&_EigenPod.CallOpts)
}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_EigenPod *EigenPodCallerSession) ExcludeSenders() ([]common.Address, error) {
	return _EigenPod.Contract.ExcludeSenders(&_EigenPod.CallOpts)
}

// HasRestaked is a free data retrieval call binding the contract method 0x3106ab53.
//
// Solidity: function hasRestaked() view returns(bool)
func (_EigenPod *EigenPodCaller) HasRestaked(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EigenPod.contract.Call(opts, &out, "hasRestaked")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRestaked is a free data retrieval call binding the contract method 0x3106ab53.
//
// Solidity: function hasRestaked() view returns(bool)
func (_EigenPod *EigenPodSession) HasRestaked() (bool, error) {
	return _EigenPod.Contract.HasRestaked(&_EigenPod.CallOpts)
}

// HasRestaked is a free data retrieval call binding the contract method 0x3106ab53.
//
// Solidity: function hasRestaked() view returns(bool)
func (_EigenPod *EigenPodCallerSession) HasRestaked() (bool, error) {
	return _EigenPod.Contract.HasRestaked(&_EigenPod.CallOpts)
}

// LastCheckpointTimestamp is a free data retrieval call binding the contract method 0xee94d67c.
//
// Solidity: function lastCheckpointTimestamp() view returns(uint64)
func (_EigenPod *EigenPodCaller) LastCheckpointTimestamp(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _EigenPod.contract.Call(opts, &out, "lastCheckpointTimestamp")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastCheckpointTimestamp is a free data retrieval call binding the contract method 0xee94d67c.
//
// Solidity: function lastCheckpointTimestamp() view returns(uint64)
func (_EigenPod *EigenPodSession) LastCheckpointTimestamp() (uint64, error) {
	return _EigenPod.Contract.LastCheckpointTimestamp(&_EigenPod.CallOpts)
}

// LastCheckpointTimestamp is a free data retrieval call binding the contract method 0xee94d67c.
//
// Solidity: function lastCheckpointTimestamp() view returns(uint64)
func (_EigenPod *EigenPodCallerSession) LastCheckpointTimestamp() (uint64, error) {
	return _EigenPod.Contract.LastCheckpointTimestamp(&_EigenPod.CallOpts)
}

// PodOwner is a free data retrieval call binding the contract method 0x0b18ff66.
//
// Solidity: function podOwner() view returns(address)
func (_EigenPod *EigenPodCaller) PodOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EigenPod.contract.Call(opts, &out, "podOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PodOwner is a free data retrieval call binding the contract method 0x0b18ff66.
//
// Solidity: function podOwner() view returns(address)
func (_EigenPod *EigenPodSession) PodOwner() (common.Address, error) {
	return _EigenPod.Contract.PodOwner(&_EigenPod.CallOpts)
}

// PodOwner is a free data retrieval call binding the contract method 0x0b18ff66.
//
// Solidity: function podOwner() view returns(address)
func (_EigenPod *EigenPodCallerSession) PodOwner() (common.Address, error) {
	return _EigenPod.Contract.PodOwner(&_EigenPod.CallOpts)
}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((address,bytes4[])[] targetedArtifactSelectors_)
func (_EigenPod *EigenPodCaller) TargetArtifactSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzSelector, error) {
	var out []interface{}
	err := _EigenPod.contract.Call(opts, &out, "targetArtifactSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzSelector)).(*[]StdInvariantFuzzSelector)

	return out0, err

}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((address,bytes4[])[] targetedArtifactSelectors_)
func (_EigenPod *EigenPodSession) TargetArtifactSelectors() ([]StdInvariantFuzzSelector, error) {
	return _EigenPod.Contract.TargetArtifactSelectors(&_EigenPod.CallOpts)
}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((address,bytes4[])[] targetedArtifactSelectors_)
func (_EigenPod *EigenPodCallerSession) TargetArtifactSelectors() ([]StdInvariantFuzzSelector, error) {
	return _EigenPod.Contract.TargetArtifactSelectors(&_EigenPod.CallOpts)
}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_EigenPod *EigenPodCaller) TargetArtifacts(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _EigenPod.contract.Call(opts, &out, "targetArtifacts")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_EigenPod *EigenPodSession) TargetArtifacts() ([]string, error) {
	return _EigenPod.Contract.TargetArtifacts(&_EigenPod.CallOpts)
}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_EigenPod *EigenPodCallerSession) TargetArtifacts() ([]string, error) {
	return _EigenPod.Contract.TargetArtifacts(&_EigenPod.CallOpts)
}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_EigenPod *EigenPodCaller) TargetContracts(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EigenPod.contract.Call(opts, &out, "targetContracts")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_EigenPod *EigenPodSession) TargetContracts() ([]common.Address, error) {
	return _EigenPod.Contract.TargetContracts(&_EigenPod.CallOpts)
}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_EigenPod *EigenPodCallerSession) TargetContracts() ([]common.Address, error) {
	return _EigenPod.Contract.TargetContracts(&_EigenPod.CallOpts)
}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_EigenPod *EigenPodCaller) TargetSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzSelector, error) {
	var out []interface{}
	err := _EigenPod.contract.Call(opts, &out, "targetSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzSelector)).(*[]StdInvariantFuzzSelector)

	return out0, err

}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_EigenPod *EigenPodSession) TargetSelectors() ([]StdInvariantFuzzSelector, error) {
	return _EigenPod.Contract.TargetSelectors(&_EigenPod.CallOpts)
}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_EigenPod *EigenPodCallerSession) TargetSelectors() ([]StdInvariantFuzzSelector, error) {
	return _EigenPod.Contract.TargetSelectors(&_EigenPod.CallOpts)
}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_EigenPod *EigenPodCaller) TargetSenders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EigenPod.contract.Call(opts, &out, "targetSenders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_EigenPod *EigenPodSession) TargetSenders() ([]common.Address, error) {
	return _EigenPod.Contract.TargetSenders(&_EigenPod.CallOpts)
}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_EigenPod *EigenPodCallerSession) TargetSenders() ([]common.Address, error) {
	return _EigenPod.Contract.TargetSenders(&_EigenPod.CallOpts)
}

// ValidatorPubkeyHashToInfo is a free data retrieval call binding the contract method 0x6fcd0e53.
//
// Solidity: function validatorPubkeyHashToInfo(bytes32 validatorPubkeyHash) view returns((uint64,uint64,uint64,uint8))
func (_EigenPod *EigenPodCaller) ValidatorPubkeyHashToInfo(opts *bind.CallOpts, validatorPubkeyHash [32]byte) (IEigenPodValidatorInfo, error) {
	var out []interface{}
	err := _EigenPod.contract.Call(opts, &out, "validatorPubkeyHashToInfo", validatorPubkeyHash)

	if err != nil {
		return *new(IEigenPodValidatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IEigenPodValidatorInfo)).(*IEigenPodValidatorInfo)

	return out0, err

}

// ValidatorPubkeyHashToInfo is a free data retrieval call binding the contract method 0x6fcd0e53.
//
// Solidity: function validatorPubkeyHashToInfo(bytes32 validatorPubkeyHash) view returns((uint64,uint64,uint64,uint8))
func (_EigenPod *EigenPodSession) ValidatorPubkeyHashToInfo(validatorPubkeyHash [32]byte) (IEigenPodValidatorInfo, error) {
	return _EigenPod.Contract.ValidatorPubkeyHashToInfo(&_EigenPod.CallOpts, validatorPubkeyHash)
}

// ValidatorPubkeyHashToInfo is a free data retrieval call binding the contract method 0x6fcd0e53.
//
// Solidity: function validatorPubkeyHashToInfo(bytes32 validatorPubkeyHash) view returns((uint64,uint64,uint64,uint8))
func (_EigenPod *EigenPodCallerSession) ValidatorPubkeyHashToInfo(validatorPubkeyHash [32]byte) (IEigenPodValidatorInfo, error) {
	return _EigenPod.Contract.ValidatorPubkeyHashToInfo(&_EigenPod.CallOpts, validatorPubkeyHash)
}

// ValidatorPubkeyToInfo is a free data retrieval call binding the contract method 0xb522538a.
//
// Solidity: function validatorPubkeyToInfo(bytes validatorPubkey) view returns((uint64,uint64,uint64,uint8))
func (_EigenPod *EigenPodCaller) ValidatorPubkeyToInfo(opts *bind.CallOpts, validatorPubkey []byte) (IEigenPodValidatorInfo, error) {
	var out []interface{}
	err := _EigenPod.contract.Call(opts, &out, "validatorPubkeyToInfo", validatorPubkey)

	if err != nil {
		return *new(IEigenPodValidatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IEigenPodValidatorInfo)).(*IEigenPodValidatorInfo)

	return out0, err

}

// ValidatorPubkeyToInfo is a free data retrieval call binding the contract method 0xb522538a.
//
// Solidity: function validatorPubkeyToInfo(bytes validatorPubkey) view returns((uint64,uint64,uint64,uint8))
func (_EigenPod *EigenPodSession) ValidatorPubkeyToInfo(validatorPubkey []byte) (IEigenPodValidatorInfo, error) {
	return _EigenPod.Contract.ValidatorPubkeyToInfo(&_EigenPod.CallOpts, validatorPubkey)
}

// ValidatorPubkeyToInfo is a free data retrieval call binding the contract method 0xb522538a.
//
// Solidity: function validatorPubkeyToInfo(bytes validatorPubkey) view returns((uint64,uint64,uint64,uint8))
func (_EigenPod *EigenPodCallerSession) ValidatorPubkeyToInfo(validatorPubkey []byte) (IEigenPodValidatorInfo, error) {
	return _EigenPod.Contract.ValidatorPubkeyToInfo(&_EigenPod.CallOpts, validatorPubkey)
}

// ValidatorStatus is a free data retrieval call binding the contract method 0x58eaee79.
//
// Solidity: function validatorStatus(bytes validatorPubkey) view returns(uint8)
func (_EigenPod *EigenPodCaller) ValidatorStatus(opts *bind.CallOpts, validatorPubkey []byte) (uint8, error) {
	var out []interface{}
	err := _EigenPod.contract.Call(opts, &out, "validatorStatus", validatorPubkey)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ValidatorStatus is a free data retrieval call binding the contract method 0x58eaee79.
//
// Solidity: function validatorStatus(bytes validatorPubkey) view returns(uint8)
func (_EigenPod *EigenPodSession) ValidatorStatus(validatorPubkey []byte) (uint8, error) {
	return _EigenPod.Contract.ValidatorStatus(&_EigenPod.CallOpts, validatorPubkey)
}

// ValidatorStatus is a free data retrieval call binding the contract method 0x58eaee79.
//
// Solidity: function validatorStatus(bytes validatorPubkey) view returns(uint8)
func (_EigenPod *EigenPodCallerSession) ValidatorStatus(validatorPubkey []byte) (uint8, error) {
	return _EigenPod.Contract.ValidatorStatus(&_EigenPod.CallOpts, validatorPubkey)
}

// ValidatorStatus0 is a free data retrieval call binding the contract method 0x7439841f.
//
// Solidity: function validatorStatus(bytes32 pubkeyHash) view returns(uint8)
func (_EigenPod *EigenPodCaller) ValidatorStatus0(opts *bind.CallOpts, pubkeyHash [32]byte) (uint8, error) {
	var out []interface{}
	err := _EigenPod.contract.Call(opts, &out, "validatorStatus0", pubkeyHash)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ValidatorStatus0 is a free data retrieval call binding the contract method 0x7439841f.
//
// Solidity: function validatorStatus(bytes32 pubkeyHash) view returns(uint8)
func (_EigenPod *EigenPodSession) ValidatorStatus0(pubkeyHash [32]byte) (uint8, error) {
	return _EigenPod.Contract.ValidatorStatus0(&_EigenPod.CallOpts, pubkeyHash)
}

// ValidatorStatus0 is a free data retrieval call binding the contract method 0x7439841f.
//
// Solidity: function validatorStatus(bytes32 pubkeyHash) view returns(uint8)
func (_EigenPod *EigenPodCallerSession) ValidatorStatus0(pubkeyHash [32]byte) (uint8, error) {
	return _EigenPod.Contract.ValidatorStatus0(&_EigenPod.CallOpts, pubkeyHash)
}

// WithdrawableRestakedExecutionLayerGwei is a free data retrieval call binding the contract method 0x3474aa16.
//
// Solidity: function withdrawableRestakedExecutionLayerGwei() view returns(uint64)
func (_EigenPod *EigenPodCaller) WithdrawableRestakedExecutionLayerGwei(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _EigenPod.contract.Call(opts, &out, "withdrawableRestakedExecutionLayerGwei")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// WithdrawableRestakedExecutionLayerGwei is a free data retrieval call binding the contract method 0x3474aa16.
//
// Solidity: function withdrawableRestakedExecutionLayerGwei() view returns(uint64)
func (_EigenPod *EigenPodSession) WithdrawableRestakedExecutionLayerGwei() (uint64, error) {
	return _EigenPod.Contract.WithdrawableRestakedExecutionLayerGwei(&_EigenPod.CallOpts)
}

// WithdrawableRestakedExecutionLayerGwei is a free data retrieval call binding the contract method 0x3474aa16.
//
// Solidity: function withdrawableRestakedExecutionLayerGwei() view returns(uint64)
func (_EigenPod *EigenPodCallerSession) WithdrawableRestakedExecutionLayerGwei() (uint64, error) {
	return _EigenPod.Contract.WithdrawableRestakedExecutionLayerGwei(&_EigenPod.CallOpts)
}

// Failed is a paid mutator transaction binding the contract method 0xba414fa6.
//
// Solidity: function failed() returns(bool)
func (_EigenPod *EigenPodTransactor) Failed(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EigenPod.contract.Transact(opts, "failed")
}

// Failed is a paid mutator transaction binding the contract method 0xba414fa6.
//
// Solidity: function failed() returns(bool)
func (_EigenPod *EigenPodSession) Failed() (*types.Transaction, error) {
	return _EigenPod.Contract.Failed(&_EigenPod.TransactOpts)
}

// Failed is a paid mutator transaction binding the contract method 0xba414fa6.
//
// Solidity: function failed() returns(bool)
func (_EigenPod *EigenPodTransactorSession) Failed() (*types.Transaction, error) {
	return _EigenPod.Contract.Failed(&_EigenPod.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _podOwner) returns()
func (_EigenPod *EigenPodTransactor) Initialize(opts *bind.TransactOpts, _podOwner common.Address) (*types.Transaction, error) {
	return _EigenPod.contract.Transact(opts, "initialize", _podOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _podOwner) returns()
func (_EigenPod *EigenPodSession) Initialize(_podOwner common.Address) (*types.Transaction, error) {
	return _EigenPod.Contract.Initialize(&_EigenPod.TransactOpts, _podOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _podOwner) returns()
func (_EigenPod *EigenPodTransactorSession) Initialize(_podOwner common.Address) (*types.Transaction, error) {
	return _EigenPod.Contract.Initialize(&_EigenPod.TransactOpts, _podOwner)
}

// RecoverTokens is a paid mutator transaction binding the contract method 0xdda3346c.
//
// Solidity: function recoverTokens(address[] tokenList, uint256[] amountsToWithdraw, address recipient) returns()
func (_EigenPod *EigenPodTransactor) RecoverTokens(opts *bind.TransactOpts, tokenList []common.Address, amountsToWithdraw []*big.Int, recipient common.Address) (*types.Transaction, error) {
	return _EigenPod.contract.Transact(opts, "recoverTokens", tokenList, amountsToWithdraw, recipient)
}

// RecoverTokens is a paid mutator transaction binding the contract method 0xdda3346c.
//
// Solidity: function recoverTokens(address[] tokenList, uint256[] amountsToWithdraw, address recipient) returns()
func (_EigenPod *EigenPodSession) RecoverTokens(tokenList []common.Address, amountsToWithdraw []*big.Int, recipient common.Address) (*types.Transaction, error) {
	return _EigenPod.Contract.RecoverTokens(&_EigenPod.TransactOpts, tokenList, amountsToWithdraw, recipient)
}

// RecoverTokens is a paid mutator transaction binding the contract method 0xdda3346c.
//
// Solidity: function recoverTokens(address[] tokenList, uint256[] amountsToWithdraw, address recipient) returns()
func (_EigenPod *EigenPodTransactorSession) RecoverTokens(tokenList []common.Address, amountsToWithdraw []*big.Int, recipient common.Address) (*types.Transaction, error) {
	return _EigenPod.Contract.RecoverTokens(&_EigenPod.TransactOpts, tokenList, amountsToWithdraw, recipient)
}

// Stake is a paid mutator transaction binding the contract method 0x9b4e4634.
//
// Solidity: function stake(bytes pubkey, bytes signature, bytes32 depositDataRoot) payable returns()
func (_EigenPod *EigenPodTransactor) Stake(opts *bind.TransactOpts, pubkey []byte, signature []byte, depositDataRoot [32]byte) (*types.Transaction, error) {
	return _EigenPod.contract.Transact(opts, "stake", pubkey, signature, depositDataRoot)
}

// Stake is a paid mutator transaction binding the contract method 0x9b4e4634.
//
// Solidity: function stake(bytes pubkey, bytes signature, bytes32 depositDataRoot) payable returns()
func (_EigenPod *EigenPodSession) Stake(pubkey []byte, signature []byte, depositDataRoot [32]byte) (*types.Transaction, error) {
	return _EigenPod.Contract.Stake(&_EigenPod.TransactOpts, pubkey, signature, depositDataRoot)
}

// Stake is a paid mutator transaction binding the contract method 0x9b4e4634.
//
// Solidity: function stake(bytes pubkey, bytes signature, bytes32 depositDataRoot) payable returns()
func (_EigenPod *EigenPodTransactorSession) Stake(pubkey []byte, signature []byte, depositDataRoot [32]byte) (*types.Transaction, error) {
	return _EigenPod.Contract.Stake(&_EigenPod.TransactOpts, pubkey, signature, depositDataRoot)
}

// StartCheckpoint is a paid mutator transaction binding the contract method 0x88676cad.
//
// Solidity: function startCheckpoint(bool revertIfNoBalance) returns()
func (_EigenPod *EigenPodTransactor) StartCheckpoint(opts *bind.TransactOpts, revertIfNoBalance bool) (*types.Transaction, error) {
	return _EigenPod.contract.Transact(opts, "startCheckpoint", revertIfNoBalance)
}

// StartCheckpoint is a paid mutator transaction binding the contract method 0x88676cad.
//
// Solidity: function startCheckpoint(bool revertIfNoBalance) returns()
func (_EigenPod *EigenPodSession) StartCheckpoint(revertIfNoBalance bool) (*types.Transaction, error) {
	return _EigenPod.Contract.StartCheckpoint(&_EigenPod.TransactOpts, revertIfNoBalance)
}

// StartCheckpoint is a paid mutator transaction binding the contract method 0x88676cad.
//
// Solidity: function startCheckpoint(bool revertIfNoBalance) returns()
func (_EigenPod *EigenPodTransactorSession) StartCheckpoint(revertIfNoBalance bool) (*types.Transaction, error) {
	return _EigenPod.Contract.StartCheckpoint(&_EigenPod.TransactOpts, revertIfNoBalance)
}

// VerifyCheckpointProofs is a paid mutator transaction binding the contract method 0xf074ba62.
//
// Solidity: function verifyCheckpointProofs((bytes32,bytes) balanceContainerProof, (bytes32,bytes32,bytes)[] proofs) returns()
func (_EigenPod *EigenPodTransactor) VerifyCheckpointProofs(opts *bind.TransactOpts, balanceContainerProof BeaconChainProofsBalanceContainerProof, proofs []BeaconChainProofsBalanceProof) (*types.Transaction, error) {
	return _EigenPod.contract.Transact(opts, "verifyCheckpointProofs", balanceContainerProof, proofs)
}

// VerifyCheckpointProofs is a paid mutator transaction binding the contract method 0xf074ba62.
//
// Solidity: function verifyCheckpointProofs((bytes32,bytes) balanceContainerProof, (bytes32,bytes32,bytes)[] proofs) returns()
func (_EigenPod *EigenPodSession) VerifyCheckpointProofs(balanceContainerProof BeaconChainProofsBalanceContainerProof, proofs []BeaconChainProofsBalanceProof) (*types.Transaction, error) {
	return _EigenPod.Contract.VerifyCheckpointProofs(&_EigenPod.TransactOpts, balanceContainerProof, proofs)
}

// VerifyCheckpointProofs is a paid mutator transaction binding the contract method 0xf074ba62.
//
// Solidity: function verifyCheckpointProofs((bytes32,bytes) balanceContainerProof, (bytes32,bytes32,bytes)[] proofs) returns()
func (_EigenPod *EigenPodTransactorSession) VerifyCheckpointProofs(balanceContainerProof BeaconChainProofsBalanceContainerProof, proofs []BeaconChainProofsBalanceProof) (*types.Transaction, error) {
	return _EigenPod.Contract.VerifyCheckpointProofs(&_EigenPod.TransactOpts, balanceContainerProof, proofs)
}

// VerifyStaleBalance is a paid mutator transaction binding the contract method 0x039157d2.
//
// Solidity: function verifyStaleBalance(uint64 beaconTimestamp, (bytes32,bytes) stateRootProof, (bytes32[],bytes) proof) returns()
func (_EigenPod *EigenPodTransactor) VerifyStaleBalance(opts *bind.TransactOpts, beaconTimestamp uint64, stateRootProof BeaconChainProofsStateRootProof, proof BeaconChainProofsValidatorProof) (*types.Transaction, error) {
	return _EigenPod.contract.Transact(opts, "verifyStaleBalance", beaconTimestamp, stateRootProof, proof)
}

// VerifyStaleBalance is a paid mutator transaction binding the contract method 0x039157d2.
//
// Solidity: function verifyStaleBalance(uint64 beaconTimestamp, (bytes32,bytes) stateRootProof, (bytes32[],bytes) proof) returns()
func (_EigenPod *EigenPodSession) VerifyStaleBalance(beaconTimestamp uint64, stateRootProof BeaconChainProofsStateRootProof, proof BeaconChainProofsValidatorProof) (*types.Transaction, error) {
	return _EigenPod.Contract.VerifyStaleBalance(&_EigenPod.TransactOpts, beaconTimestamp, stateRootProof, proof)
}

// VerifyStaleBalance is a paid mutator transaction binding the contract method 0x039157d2.
//
// Solidity: function verifyStaleBalance(uint64 beaconTimestamp, (bytes32,bytes) stateRootProof, (bytes32[],bytes) proof) returns()
func (_EigenPod *EigenPodTransactorSession) VerifyStaleBalance(beaconTimestamp uint64, stateRootProof BeaconChainProofsStateRootProof, proof BeaconChainProofsValidatorProof) (*types.Transaction, error) {
	return _EigenPod.Contract.VerifyStaleBalance(&_EigenPod.TransactOpts, beaconTimestamp, stateRootProof, proof)
}

// VerifyWithdrawalCredentials is a paid mutator transaction binding the contract method 0x3f65cf19.
//
// Solidity: function verifyWithdrawalCredentials(uint64 beaconTimestamp, (bytes32,bytes) stateRootProof, uint40[] validatorIndices, bytes[] validatorFieldsProofs, bytes32[][] validatorFields) returns()
func (_EigenPod *EigenPodTransactor) VerifyWithdrawalCredentials(opts *bind.TransactOpts, beaconTimestamp uint64, stateRootProof BeaconChainProofsStateRootProof, validatorIndices []*big.Int, validatorFieldsProofs [][]byte, validatorFields [][][32]byte) (*types.Transaction, error) {
	return _EigenPod.contract.Transact(opts, "verifyWithdrawalCredentials", beaconTimestamp, stateRootProof, validatorIndices, validatorFieldsProofs, validatorFields)
}

// VerifyWithdrawalCredentials is a paid mutator transaction binding the contract method 0x3f65cf19.
//
// Solidity: function verifyWithdrawalCredentials(uint64 beaconTimestamp, (bytes32,bytes) stateRootProof, uint40[] validatorIndices, bytes[] validatorFieldsProofs, bytes32[][] validatorFields) returns()
func (_EigenPod *EigenPodSession) VerifyWithdrawalCredentials(beaconTimestamp uint64, stateRootProof BeaconChainProofsStateRootProof, validatorIndices []*big.Int, validatorFieldsProofs [][]byte, validatorFields [][][32]byte) (*types.Transaction, error) {
	return _EigenPod.Contract.VerifyWithdrawalCredentials(&_EigenPod.TransactOpts, beaconTimestamp, stateRootProof, validatorIndices, validatorFieldsProofs, validatorFields)
}

// VerifyWithdrawalCredentials is a paid mutator transaction binding the contract method 0x3f65cf19.
//
// Solidity: function verifyWithdrawalCredentials(uint64 beaconTimestamp, (bytes32,bytes) stateRootProof, uint40[] validatorIndices, bytes[] validatorFieldsProofs, bytes32[][] validatorFields) returns()
func (_EigenPod *EigenPodTransactorSession) VerifyWithdrawalCredentials(beaconTimestamp uint64, stateRootProof BeaconChainProofsStateRootProof, validatorIndices []*big.Int, validatorFieldsProofs [][]byte, validatorFields [][][32]byte) (*types.Transaction, error) {
	return _EigenPod.Contract.VerifyWithdrawalCredentials(&_EigenPod.TransactOpts, beaconTimestamp, stateRootProof, validatorIndices, validatorFieldsProofs, validatorFields)
}

// WithdrawRestakedBeaconChainETH is a paid mutator transaction binding the contract method 0xc4907442.
//
// Solidity: function withdrawRestakedBeaconChainETH(address recipient, uint256 amountWei) returns()
func (_EigenPod *EigenPodTransactor) WithdrawRestakedBeaconChainETH(opts *bind.TransactOpts, recipient common.Address, amountWei *big.Int) (*types.Transaction, error) {
	return _EigenPod.contract.Transact(opts, "withdrawRestakedBeaconChainETH", recipient, amountWei)
}

// WithdrawRestakedBeaconChainETH is a paid mutator transaction binding the contract method 0xc4907442.
//
// Solidity: function withdrawRestakedBeaconChainETH(address recipient, uint256 amountWei) returns()
func (_EigenPod *EigenPodSession) WithdrawRestakedBeaconChainETH(recipient common.Address, amountWei *big.Int) (*types.Transaction, error) {
	return _EigenPod.Contract.WithdrawRestakedBeaconChainETH(&_EigenPod.TransactOpts, recipient, amountWei)
}

// WithdrawRestakedBeaconChainETH is a paid mutator transaction binding the contract method 0xc4907442.
//
// Solidity: function withdrawRestakedBeaconChainETH(address recipient, uint256 amountWei) returns()
func (_EigenPod *EigenPodTransactorSession) WithdrawRestakedBeaconChainETH(recipient common.Address, amountWei *big.Int) (*types.Transaction, error) {
	return _EigenPod.Contract.WithdrawRestakedBeaconChainETH(&_EigenPod.TransactOpts, recipient, amountWei)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_EigenPod *EigenPodTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EigenPod.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_EigenPod *EigenPodSession) Receive() (*types.Transaction, error) {
	return _EigenPod.Contract.Receive(&_EigenPod.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_EigenPod *EigenPodTransactorSession) Receive() (*types.Transaction, error) {
	return _EigenPod.Contract.Receive(&_EigenPod.TransactOpts)
}

// EigenPodCheckpointCreatedIterator is returned from FilterCheckpointCreated and is used to iterate over the raw logs and unpacked data for CheckpointCreated events raised by the EigenPod contract.
type EigenPodCheckpointCreatedIterator struct {
	Event *EigenPodCheckpointCreated // Event containing the contract specifics and raw log

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
func (it *EigenPodCheckpointCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenPodCheckpointCreated)
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
		it.Event = new(EigenPodCheckpointCreated)
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
func (it *EigenPodCheckpointCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenPodCheckpointCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenPodCheckpointCreated represents a CheckpointCreated event raised by the EigenPod contract.
type EigenPodCheckpointCreated struct {
	CheckpointTimestamp uint64
	BeaconBlockRoot     [32]byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterCheckpointCreated is a free log retrieval operation binding the contract event 0x17a875d60c621f9c2f4d68122fe3a5bd359233b6fd8fcb79a0f65a6e433c6bf6.
//
// Solidity: event CheckpointCreated(uint64 indexed checkpointTimestamp, bytes32 indexed beaconBlockRoot)
func (_EigenPod *EigenPodFilterer) FilterCheckpointCreated(opts *bind.FilterOpts, checkpointTimestamp []uint64, beaconBlockRoot [][32]byte) (*EigenPodCheckpointCreatedIterator, error) {

	var checkpointTimestampRule []interface{}
	for _, checkpointTimestampItem := range checkpointTimestamp {
		checkpointTimestampRule = append(checkpointTimestampRule, checkpointTimestampItem)
	}
	var beaconBlockRootRule []interface{}
	for _, beaconBlockRootItem := range beaconBlockRoot {
		beaconBlockRootRule = append(beaconBlockRootRule, beaconBlockRootItem)
	}

	logs, sub, err := _EigenPod.contract.FilterLogs(opts, "CheckpointCreated", checkpointTimestampRule, beaconBlockRootRule)
	if err != nil {
		return nil, err
	}
	return &EigenPodCheckpointCreatedIterator{contract: _EigenPod.contract, event: "CheckpointCreated", logs: logs, sub: sub}, nil
}

// WatchCheckpointCreated is a free log subscription operation binding the contract event 0x17a875d60c621f9c2f4d68122fe3a5bd359233b6fd8fcb79a0f65a6e433c6bf6.
//
// Solidity: event CheckpointCreated(uint64 indexed checkpointTimestamp, bytes32 indexed beaconBlockRoot)
func (_EigenPod *EigenPodFilterer) WatchCheckpointCreated(opts *bind.WatchOpts, sink chan<- *EigenPodCheckpointCreated, checkpointTimestamp []uint64, beaconBlockRoot [][32]byte) (event.Subscription, error) {

	var checkpointTimestampRule []interface{}
	for _, checkpointTimestampItem := range checkpointTimestamp {
		checkpointTimestampRule = append(checkpointTimestampRule, checkpointTimestampItem)
	}
	var beaconBlockRootRule []interface{}
	for _, beaconBlockRootItem := range beaconBlockRoot {
		beaconBlockRootRule = append(beaconBlockRootRule, beaconBlockRootItem)
	}

	logs, sub, err := _EigenPod.contract.WatchLogs(opts, "CheckpointCreated", checkpointTimestampRule, beaconBlockRootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenPodCheckpointCreated)
				if err := _EigenPod.contract.UnpackLog(event, "CheckpointCreated", log); err != nil {
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

// ParseCheckpointCreated is a log parse operation binding the contract event 0x17a875d60c621f9c2f4d68122fe3a5bd359233b6fd8fcb79a0f65a6e433c6bf6.
//
// Solidity: event CheckpointCreated(uint64 indexed checkpointTimestamp, bytes32 indexed beaconBlockRoot)
func (_EigenPod *EigenPodFilterer) ParseCheckpointCreated(log types.Log) (*EigenPodCheckpointCreated, error) {
	event := new(EigenPodCheckpointCreated)
	if err := _EigenPod.contract.UnpackLog(event, "CheckpointCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenPodCheckpointFinalizedIterator is returned from FilterCheckpointFinalized and is used to iterate over the raw logs and unpacked data for CheckpointFinalized events raised by the EigenPod contract.
type EigenPodCheckpointFinalizedIterator struct {
	Event *EigenPodCheckpointFinalized // Event containing the contract specifics and raw log

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
func (it *EigenPodCheckpointFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenPodCheckpointFinalized)
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
		it.Event = new(EigenPodCheckpointFinalized)
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
func (it *EigenPodCheckpointFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenPodCheckpointFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenPodCheckpointFinalized represents a CheckpointFinalized event raised by the EigenPod contract.
type EigenPodCheckpointFinalized struct {
	CheckpointTimestamp uint64
	TotalShareDeltaWei  *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterCheckpointFinalized is a free log retrieval operation binding the contract event 0x525408c201bc1576eb44116f6478f1c2a54775b19a043bcfdc708364f74f8e44.
//
// Solidity: event CheckpointFinalized(uint64 indexed checkpointTimestamp, int256 totalShareDeltaWei)
func (_EigenPod *EigenPodFilterer) FilterCheckpointFinalized(opts *bind.FilterOpts, checkpointTimestamp []uint64) (*EigenPodCheckpointFinalizedIterator, error) {

	var checkpointTimestampRule []interface{}
	for _, checkpointTimestampItem := range checkpointTimestamp {
		checkpointTimestampRule = append(checkpointTimestampRule, checkpointTimestampItem)
	}

	logs, sub, err := _EigenPod.contract.FilterLogs(opts, "CheckpointFinalized", checkpointTimestampRule)
	if err != nil {
		return nil, err
	}
	return &EigenPodCheckpointFinalizedIterator{contract: _EigenPod.contract, event: "CheckpointFinalized", logs: logs, sub: sub}, nil
}

// WatchCheckpointFinalized is a free log subscription operation binding the contract event 0x525408c201bc1576eb44116f6478f1c2a54775b19a043bcfdc708364f74f8e44.
//
// Solidity: event CheckpointFinalized(uint64 indexed checkpointTimestamp, int256 totalShareDeltaWei)
func (_EigenPod *EigenPodFilterer) WatchCheckpointFinalized(opts *bind.WatchOpts, sink chan<- *EigenPodCheckpointFinalized, checkpointTimestamp []uint64) (event.Subscription, error) {

	var checkpointTimestampRule []interface{}
	for _, checkpointTimestampItem := range checkpointTimestamp {
		checkpointTimestampRule = append(checkpointTimestampRule, checkpointTimestampItem)
	}

	logs, sub, err := _EigenPod.contract.WatchLogs(opts, "CheckpointFinalized", checkpointTimestampRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenPodCheckpointFinalized)
				if err := _EigenPod.contract.UnpackLog(event, "CheckpointFinalized", log); err != nil {
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

// ParseCheckpointFinalized is a log parse operation binding the contract event 0x525408c201bc1576eb44116f6478f1c2a54775b19a043bcfdc708364f74f8e44.
//
// Solidity: event CheckpointFinalized(uint64 indexed checkpointTimestamp, int256 totalShareDeltaWei)
func (_EigenPod *EigenPodFilterer) ParseCheckpointFinalized(log types.Log) (*EigenPodCheckpointFinalized, error) {
	event := new(EigenPodCheckpointFinalized)
	if err := _EigenPod.contract.UnpackLog(event, "CheckpointFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenPodEigenPodStakedIterator is returned from FilterEigenPodStaked and is used to iterate over the raw logs and unpacked data for EigenPodStaked events raised by the EigenPod contract.
type EigenPodEigenPodStakedIterator struct {
	Event *EigenPodEigenPodStaked // Event containing the contract specifics and raw log

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
func (it *EigenPodEigenPodStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenPodEigenPodStaked)
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
		it.Event = new(EigenPodEigenPodStaked)
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
func (it *EigenPodEigenPodStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenPodEigenPodStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenPodEigenPodStaked represents a EigenPodStaked event raised by the EigenPod contract.
type EigenPodEigenPodStaked struct {
	Pubkey []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEigenPodStaked is a free log retrieval operation binding the contract event 0x606865b7934a25d4aed43f6cdb426403353fa4b3009c4d228407474581b01e23.
//
// Solidity: event EigenPodStaked(bytes pubkey)
func (_EigenPod *EigenPodFilterer) FilterEigenPodStaked(opts *bind.FilterOpts) (*EigenPodEigenPodStakedIterator, error) {

	logs, sub, err := _EigenPod.contract.FilterLogs(opts, "EigenPodStaked")
	if err != nil {
		return nil, err
	}
	return &EigenPodEigenPodStakedIterator{contract: _EigenPod.contract, event: "EigenPodStaked", logs: logs, sub: sub}, nil
}

// WatchEigenPodStaked is a free log subscription operation binding the contract event 0x606865b7934a25d4aed43f6cdb426403353fa4b3009c4d228407474581b01e23.
//
// Solidity: event EigenPodStaked(bytes pubkey)
func (_EigenPod *EigenPodFilterer) WatchEigenPodStaked(opts *bind.WatchOpts, sink chan<- *EigenPodEigenPodStaked) (event.Subscription, error) {

	logs, sub, err := _EigenPod.contract.WatchLogs(opts, "EigenPodStaked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenPodEigenPodStaked)
				if err := _EigenPod.contract.UnpackLog(event, "EigenPodStaked", log); err != nil {
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

// ParseEigenPodStaked is a log parse operation binding the contract event 0x606865b7934a25d4aed43f6cdb426403353fa4b3009c4d228407474581b01e23.
//
// Solidity: event EigenPodStaked(bytes pubkey)
func (_EigenPod *EigenPodFilterer) ParseEigenPodStaked(log types.Log) (*EigenPodEigenPodStaked, error) {
	event := new(EigenPodEigenPodStaked)
	if err := _EigenPod.contract.UnpackLog(event, "EigenPodStaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenPodInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the EigenPod contract.
type EigenPodInitializedIterator struct {
	Event *EigenPodInitialized // Event containing the contract specifics and raw log

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
func (it *EigenPodInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenPodInitialized)
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
		it.Event = new(EigenPodInitialized)
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
func (it *EigenPodInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenPodInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenPodInitialized represents a Initialized event raised by the EigenPod contract.
type EigenPodInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_EigenPod *EigenPodFilterer) FilterInitialized(opts *bind.FilterOpts) (*EigenPodInitializedIterator, error) {

	logs, sub, err := _EigenPod.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &EigenPodInitializedIterator{contract: _EigenPod.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_EigenPod *EigenPodFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *EigenPodInitialized) (event.Subscription, error) {

	logs, sub, err := _EigenPod.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenPodInitialized)
				if err := _EigenPod.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_EigenPod *EigenPodFilterer) ParseInitialized(log types.Log) (*EigenPodInitialized, error) {
	event := new(EigenPodInitialized)
	if err := _EigenPod.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenPodNonBeaconChainETHReceivedIterator is returned from FilterNonBeaconChainETHReceived and is used to iterate over the raw logs and unpacked data for NonBeaconChainETHReceived events raised by the EigenPod contract.
type EigenPodNonBeaconChainETHReceivedIterator struct {
	Event *EigenPodNonBeaconChainETHReceived // Event containing the contract specifics and raw log

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
func (it *EigenPodNonBeaconChainETHReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenPodNonBeaconChainETHReceived)
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
		it.Event = new(EigenPodNonBeaconChainETHReceived)
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
func (it *EigenPodNonBeaconChainETHReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenPodNonBeaconChainETHReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenPodNonBeaconChainETHReceived represents a NonBeaconChainETHReceived event raised by the EigenPod contract.
type EigenPodNonBeaconChainETHReceived struct {
	AmountReceived *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNonBeaconChainETHReceived is a free log retrieval operation binding the contract event 0x6fdd3dbdb173299608c0aa9f368735857c8842b581f8389238bf05bd04b3bf49.
//
// Solidity: event NonBeaconChainETHReceived(uint256 amountReceived)
func (_EigenPod *EigenPodFilterer) FilterNonBeaconChainETHReceived(opts *bind.FilterOpts) (*EigenPodNonBeaconChainETHReceivedIterator, error) {

	logs, sub, err := _EigenPod.contract.FilterLogs(opts, "NonBeaconChainETHReceived")
	if err != nil {
		return nil, err
	}
	return &EigenPodNonBeaconChainETHReceivedIterator{contract: _EigenPod.contract, event: "NonBeaconChainETHReceived", logs: logs, sub: sub}, nil
}

// WatchNonBeaconChainETHReceived is a free log subscription operation binding the contract event 0x6fdd3dbdb173299608c0aa9f368735857c8842b581f8389238bf05bd04b3bf49.
//
// Solidity: event NonBeaconChainETHReceived(uint256 amountReceived)
func (_EigenPod *EigenPodFilterer) WatchNonBeaconChainETHReceived(opts *bind.WatchOpts, sink chan<- *EigenPodNonBeaconChainETHReceived) (event.Subscription, error) {

	logs, sub, err := _EigenPod.contract.WatchLogs(opts, "NonBeaconChainETHReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenPodNonBeaconChainETHReceived)
				if err := _EigenPod.contract.UnpackLog(event, "NonBeaconChainETHReceived", log); err != nil {
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

// ParseNonBeaconChainETHReceived is a log parse operation binding the contract event 0x6fdd3dbdb173299608c0aa9f368735857c8842b581f8389238bf05bd04b3bf49.
//
// Solidity: event NonBeaconChainETHReceived(uint256 amountReceived)
func (_EigenPod *EigenPodFilterer) ParseNonBeaconChainETHReceived(log types.Log) (*EigenPodNonBeaconChainETHReceived, error) {
	event := new(EigenPodNonBeaconChainETHReceived)
	if err := _EigenPod.contract.UnpackLog(event, "NonBeaconChainETHReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenPodRestakedBeaconChainETHWithdrawnIterator is returned from FilterRestakedBeaconChainETHWithdrawn and is used to iterate over the raw logs and unpacked data for RestakedBeaconChainETHWithdrawn events raised by the EigenPod contract.
type EigenPodRestakedBeaconChainETHWithdrawnIterator struct {
	Event *EigenPodRestakedBeaconChainETHWithdrawn // Event containing the contract specifics and raw log

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
func (it *EigenPodRestakedBeaconChainETHWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenPodRestakedBeaconChainETHWithdrawn)
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
		it.Event = new(EigenPodRestakedBeaconChainETHWithdrawn)
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
func (it *EigenPodRestakedBeaconChainETHWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenPodRestakedBeaconChainETHWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenPodRestakedBeaconChainETHWithdrawn represents a RestakedBeaconChainETHWithdrawn event raised by the EigenPod contract.
type EigenPodRestakedBeaconChainETHWithdrawn struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRestakedBeaconChainETHWithdrawn is a free log retrieval operation binding the contract event 0x8947fd2ce07ef9cc302c4e8f0461015615d91ce851564839e91cc804c2f49d8e.
//
// Solidity: event RestakedBeaconChainETHWithdrawn(address indexed recipient, uint256 amount)
func (_EigenPod *EigenPodFilterer) FilterRestakedBeaconChainETHWithdrawn(opts *bind.FilterOpts, recipient []common.Address) (*EigenPodRestakedBeaconChainETHWithdrawnIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _EigenPod.contract.FilterLogs(opts, "RestakedBeaconChainETHWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return &EigenPodRestakedBeaconChainETHWithdrawnIterator{contract: _EigenPod.contract, event: "RestakedBeaconChainETHWithdrawn", logs: logs, sub: sub}, nil
}

// WatchRestakedBeaconChainETHWithdrawn is a free log subscription operation binding the contract event 0x8947fd2ce07ef9cc302c4e8f0461015615d91ce851564839e91cc804c2f49d8e.
//
// Solidity: event RestakedBeaconChainETHWithdrawn(address indexed recipient, uint256 amount)
func (_EigenPod *EigenPodFilterer) WatchRestakedBeaconChainETHWithdrawn(opts *bind.WatchOpts, sink chan<- *EigenPodRestakedBeaconChainETHWithdrawn, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _EigenPod.contract.WatchLogs(opts, "RestakedBeaconChainETHWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenPodRestakedBeaconChainETHWithdrawn)
				if err := _EigenPod.contract.UnpackLog(event, "RestakedBeaconChainETHWithdrawn", log); err != nil {
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

// ParseRestakedBeaconChainETHWithdrawn is a log parse operation binding the contract event 0x8947fd2ce07ef9cc302c4e8f0461015615d91ce851564839e91cc804c2f49d8e.
//
// Solidity: event RestakedBeaconChainETHWithdrawn(address indexed recipient, uint256 amount)
func (_EigenPod *EigenPodFilterer) ParseRestakedBeaconChainETHWithdrawn(log types.Log) (*EigenPodRestakedBeaconChainETHWithdrawn, error) {
	event := new(EigenPodRestakedBeaconChainETHWithdrawn)
	if err := _EigenPod.contract.UnpackLog(event, "RestakedBeaconChainETHWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenPodRestakingActivatedIterator is returned from FilterRestakingActivated and is used to iterate over the raw logs and unpacked data for RestakingActivated events raised by the EigenPod contract.
type EigenPodRestakingActivatedIterator struct {
	Event *EigenPodRestakingActivated // Event containing the contract specifics and raw log

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
func (it *EigenPodRestakingActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenPodRestakingActivated)
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
		it.Event = new(EigenPodRestakingActivated)
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
func (it *EigenPodRestakingActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenPodRestakingActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenPodRestakingActivated represents a RestakingActivated event raised by the EigenPod contract.
type EigenPodRestakingActivated struct {
	PodOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRestakingActivated is a free log retrieval operation binding the contract event 0xca8dfc8c5e0a67a74501c072a3325f685259bebbae7cfd230ab85198a78b70cd.
//
// Solidity: event RestakingActivated(address indexed podOwner)
func (_EigenPod *EigenPodFilterer) FilterRestakingActivated(opts *bind.FilterOpts, podOwner []common.Address) (*EigenPodRestakingActivatedIterator, error) {

	var podOwnerRule []interface{}
	for _, podOwnerItem := range podOwner {
		podOwnerRule = append(podOwnerRule, podOwnerItem)
	}

	logs, sub, err := _EigenPod.contract.FilterLogs(opts, "RestakingActivated", podOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EigenPodRestakingActivatedIterator{contract: _EigenPod.contract, event: "RestakingActivated", logs: logs, sub: sub}, nil
}

// WatchRestakingActivated is a free log subscription operation binding the contract event 0xca8dfc8c5e0a67a74501c072a3325f685259bebbae7cfd230ab85198a78b70cd.
//
// Solidity: event RestakingActivated(address indexed podOwner)
func (_EigenPod *EigenPodFilterer) WatchRestakingActivated(opts *bind.WatchOpts, sink chan<- *EigenPodRestakingActivated, podOwner []common.Address) (event.Subscription, error) {

	var podOwnerRule []interface{}
	for _, podOwnerItem := range podOwner {
		podOwnerRule = append(podOwnerRule, podOwnerItem)
	}

	logs, sub, err := _EigenPod.contract.WatchLogs(opts, "RestakingActivated", podOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenPodRestakingActivated)
				if err := _EigenPod.contract.UnpackLog(event, "RestakingActivated", log); err != nil {
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

// ParseRestakingActivated is a log parse operation binding the contract event 0xca8dfc8c5e0a67a74501c072a3325f685259bebbae7cfd230ab85198a78b70cd.
//
// Solidity: event RestakingActivated(address indexed podOwner)
func (_EigenPod *EigenPodFilterer) ParseRestakingActivated(log types.Log) (*EigenPodRestakingActivated, error) {
	event := new(EigenPodRestakingActivated)
	if err := _EigenPod.contract.UnpackLog(event, "RestakingActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenPodValidatorBalanceUpdatedIterator is returned from FilterValidatorBalanceUpdated and is used to iterate over the raw logs and unpacked data for ValidatorBalanceUpdated events raised by the EigenPod contract.
type EigenPodValidatorBalanceUpdatedIterator struct {
	Event *EigenPodValidatorBalanceUpdated // Event containing the contract specifics and raw log

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
func (it *EigenPodValidatorBalanceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenPodValidatorBalanceUpdated)
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
		it.Event = new(EigenPodValidatorBalanceUpdated)
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
func (it *EigenPodValidatorBalanceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenPodValidatorBalanceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenPodValidatorBalanceUpdated represents a ValidatorBalanceUpdated event raised by the EigenPod contract.
type EigenPodValidatorBalanceUpdated struct {
	ValidatorIndex          *big.Int
	BalanceTimestamp        uint64
	NewValidatorBalanceGwei uint64
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterValidatorBalanceUpdated is a free log retrieval operation binding the contract event 0x0e5fac175b83177cc047381e030d8fb3b42b37bd1c025e22c280facad62c32df.
//
// Solidity: event ValidatorBalanceUpdated(uint40 validatorIndex, uint64 balanceTimestamp, uint64 newValidatorBalanceGwei)
func (_EigenPod *EigenPodFilterer) FilterValidatorBalanceUpdated(opts *bind.FilterOpts) (*EigenPodValidatorBalanceUpdatedIterator, error) {

	logs, sub, err := _EigenPod.contract.FilterLogs(opts, "ValidatorBalanceUpdated")
	if err != nil {
		return nil, err
	}
	return &EigenPodValidatorBalanceUpdatedIterator{contract: _EigenPod.contract, event: "ValidatorBalanceUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorBalanceUpdated is a free log subscription operation binding the contract event 0x0e5fac175b83177cc047381e030d8fb3b42b37bd1c025e22c280facad62c32df.
//
// Solidity: event ValidatorBalanceUpdated(uint40 validatorIndex, uint64 balanceTimestamp, uint64 newValidatorBalanceGwei)
func (_EigenPod *EigenPodFilterer) WatchValidatorBalanceUpdated(opts *bind.WatchOpts, sink chan<- *EigenPodValidatorBalanceUpdated) (event.Subscription, error) {

	logs, sub, err := _EigenPod.contract.WatchLogs(opts, "ValidatorBalanceUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenPodValidatorBalanceUpdated)
				if err := _EigenPod.contract.UnpackLog(event, "ValidatorBalanceUpdated", log); err != nil {
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

// ParseValidatorBalanceUpdated is a log parse operation binding the contract event 0x0e5fac175b83177cc047381e030d8fb3b42b37bd1c025e22c280facad62c32df.
//
// Solidity: event ValidatorBalanceUpdated(uint40 validatorIndex, uint64 balanceTimestamp, uint64 newValidatorBalanceGwei)
func (_EigenPod *EigenPodFilterer) ParseValidatorBalanceUpdated(log types.Log) (*EigenPodValidatorBalanceUpdated, error) {
	event := new(EigenPodValidatorBalanceUpdated)
	if err := _EigenPod.contract.UnpackLog(event, "ValidatorBalanceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenPodValidatorCheckpointedIterator is returned from FilterValidatorCheckpointed and is used to iterate over the raw logs and unpacked data for ValidatorCheckpointed events raised by the EigenPod contract.
type EigenPodValidatorCheckpointedIterator struct {
	Event *EigenPodValidatorCheckpointed // Event containing the contract specifics and raw log

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
func (it *EigenPodValidatorCheckpointedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenPodValidatorCheckpointed)
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
		it.Event = new(EigenPodValidatorCheckpointed)
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
func (it *EigenPodValidatorCheckpointedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenPodValidatorCheckpointedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenPodValidatorCheckpointed represents a ValidatorCheckpointed event raised by the EigenPod contract.
type EigenPodValidatorCheckpointed struct {
	CheckpointTimestamp uint64
	ValidatorIndex      *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterValidatorCheckpointed is a free log retrieval operation binding the contract event 0xa91c59033c3423e18b54d0acecebb4972f9ea95aedf5f4cae3b677b02eaf3a3f.
//
// Solidity: event ValidatorCheckpointed(uint64 indexed checkpointTimestamp, uint40 indexed validatorIndex)
func (_EigenPod *EigenPodFilterer) FilterValidatorCheckpointed(opts *bind.FilterOpts, checkpointTimestamp []uint64, validatorIndex []*big.Int) (*EigenPodValidatorCheckpointedIterator, error) {

	var checkpointTimestampRule []interface{}
	for _, checkpointTimestampItem := range checkpointTimestamp {
		checkpointTimestampRule = append(checkpointTimestampRule, checkpointTimestampItem)
	}
	var validatorIndexRule []interface{}
	for _, validatorIndexItem := range validatorIndex {
		validatorIndexRule = append(validatorIndexRule, validatorIndexItem)
	}

	logs, sub, err := _EigenPod.contract.FilterLogs(opts, "ValidatorCheckpointed", checkpointTimestampRule, validatorIndexRule)
	if err != nil {
		return nil, err
	}
	return &EigenPodValidatorCheckpointedIterator{contract: _EigenPod.contract, event: "ValidatorCheckpointed", logs: logs, sub: sub}, nil
}

// WatchValidatorCheckpointed is a free log subscription operation binding the contract event 0xa91c59033c3423e18b54d0acecebb4972f9ea95aedf5f4cae3b677b02eaf3a3f.
//
// Solidity: event ValidatorCheckpointed(uint64 indexed checkpointTimestamp, uint40 indexed validatorIndex)
func (_EigenPod *EigenPodFilterer) WatchValidatorCheckpointed(opts *bind.WatchOpts, sink chan<- *EigenPodValidatorCheckpointed, checkpointTimestamp []uint64, validatorIndex []*big.Int) (event.Subscription, error) {

	var checkpointTimestampRule []interface{}
	for _, checkpointTimestampItem := range checkpointTimestamp {
		checkpointTimestampRule = append(checkpointTimestampRule, checkpointTimestampItem)
	}
	var validatorIndexRule []interface{}
	for _, validatorIndexItem := range validatorIndex {
		validatorIndexRule = append(validatorIndexRule, validatorIndexItem)
	}

	logs, sub, err := _EigenPod.contract.WatchLogs(opts, "ValidatorCheckpointed", checkpointTimestampRule, validatorIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenPodValidatorCheckpointed)
				if err := _EigenPod.contract.UnpackLog(event, "ValidatorCheckpointed", log); err != nil {
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

// ParseValidatorCheckpointed is a log parse operation binding the contract event 0xa91c59033c3423e18b54d0acecebb4972f9ea95aedf5f4cae3b677b02eaf3a3f.
//
// Solidity: event ValidatorCheckpointed(uint64 indexed checkpointTimestamp, uint40 indexed validatorIndex)
func (_EigenPod *EigenPodFilterer) ParseValidatorCheckpointed(log types.Log) (*EigenPodValidatorCheckpointed, error) {
	event := new(EigenPodValidatorCheckpointed)
	if err := _EigenPod.contract.UnpackLog(event, "ValidatorCheckpointed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenPodValidatorRestakedIterator is returned from FilterValidatorRestaked and is used to iterate over the raw logs and unpacked data for ValidatorRestaked events raised by the EigenPod contract.
type EigenPodValidatorRestakedIterator struct {
	Event *EigenPodValidatorRestaked // Event containing the contract specifics and raw log

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
func (it *EigenPodValidatorRestakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenPodValidatorRestaked)
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
		it.Event = new(EigenPodValidatorRestaked)
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
func (it *EigenPodValidatorRestakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenPodValidatorRestakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenPodValidatorRestaked represents a ValidatorRestaked event raised by the EigenPod contract.
type EigenPodValidatorRestaked struct {
	ValidatorIndex *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterValidatorRestaked is a free log retrieval operation binding the contract event 0x2d0800bbc377ea54a08c5db6a87aafff5e3e9c8fead0eda110e40e0c10441449.
//
// Solidity: event ValidatorRestaked(uint40 validatorIndex)
func (_EigenPod *EigenPodFilterer) FilterValidatorRestaked(opts *bind.FilterOpts) (*EigenPodValidatorRestakedIterator, error) {

	logs, sub, err := _EigenPod.contract.FilterLogs(opts, "ValidatorRestaked")
	if err != nil {
		return nil, err
	}
	return &EigenPodValidatorRestakedIterator{contract: _EigenPod.contract, event: "ValidatorRestaked", logs: logs, sub: sub}, nil
}

// WatchValidatorRestaked is a free log subscription operation binding the contract event 0x2d0800bbc377ea54a08c5db6a87aafff5e3e9c8fead0eda110e40e0c10441449.
//
// Solidity: event ValidatorRestaked(uint40 validatorIndex)
func (_EigenPod *EigenPodFilterer) WatchValidatorRestaked(opts *bind.WatchOpts, sink chan<- *EigenPodValidatorRestaked) (event.Subscription, error) {

	logs, sub, err := _EigenPod.contract.WatchLogs(opts, "ValidatorRestaked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenPodValidatorRestaked)
				if err := _EigenPod.contract.UnpackLog(event, "ValidatorRestaked", log); err != nil {
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

// ParseValidatorRestaked is a log parse operation binding the contract event 0x2d0800bbc377ea54a08c5db6a87aafff5e3e9c8fead0eda110e40e0c10441449.
//
// Solidity: event ValidatorRestaked(uint40 validatorIndex)
func (_EigenPod *EigenPodFilterer) ParseValidatorRestaked(log types.Log) (*EigenPodValidatorRestaked, error) {
	event := new(EigenPodValidatorRestaked)
	if err := _EigenPod.contract.UnpackLog(event, "ValidatorRestaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenPodValidatorWithdrawnIterator is returned from FilterValidatorWithdrawn and is used to iterate over the raw logs and unpacked data for ValidatorWithdrawn events raised by the EigenPod contract.
type EigenPodValidatorWithdrawnIterator struct {
	Event *EigenPodValidatorWithdrawn // Event containing the contract specifics and raw log

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
func (it *EigenPodValidatorWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenPodValidatorWithdrawn)
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
		it.Event = new(EigenPodValidatorWithdrawn)
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
func (it *EigenPodValidatorWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenPodValidatorWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenPodValidatorWithdrawn represents a ValidatorWithdrawn event raised by the EigenPod contract.
type EigenPodValidatorWithdrawn struct {
	CheckpointTimestamp uint64
	ValidatorIndex      *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterValidatorWithdrawn is a free log retrieval operation binding the contract event 0x2a02361ffa66cf2c2da4682c2355a6adcaa9f6c227b6e6563e68480f9587626a.
//
// Solidity: event ValidatorWithdrawn(uint64 indexed checkpointTimestamp, uint40 indexed validatorIndex)
func (_EigenPod *EigenPodFilterer) FilterValidatorWithdrawn(opts *bind.FilterOpts, checkpointTimestamp []uint64, validatorIndex []*big.Int) (*EigenPodValidatorWithdrawnIterator, error) {

	var checkpointTimestampRule []interface{}
	for _, checkpointTimestampItem := range checkpointTimestamp {
		checkpointTimestampRule = append(checkpointTimestampRule, checkpointTimestampItem)
	}
	var validatorIndexRule []interface{}
	for _, validatorIndexItem := range validatorIndex {
		validatorIndexRule = append(validatorIndexRule, validatorIndexItem)
	}

	logs, sub, err := _EigenPod.contract.FilterLogs(opts, "ValidatorWithdrawn", checkpointTimestampRule, validatorIndexRule)
	if err != nil {
		return nil, err
	}
	return &EigenPodValidatorWithdrawnIterator{contract: _EigenPod.contract, event: "ValidatorWithdrawn", logs: logs, sub: sub}, nil
}

// WatchValidatorWithdrawn is a free log subscription operation binding the contract event 0x2a02361ffa66cf2c2da4682c2355a6adcaa9f6c227b6e6563e68480f9587626a.
//
// Solidity: event ValidatorWithdrawn(uint64 indexed checkpointTimestamp, uint40 indexed validatorIndex)
func (_EigenPod *EigenPodFilterer) WatchValidatorWithdrawn(opts *bind.WatchOpts, sink chan<- *EigenPodValidatorWithdrawn, checkpointTimestamp []uint64, validatorIndex []*big.Int) (event.Subscription, error) {

	var checkpointTimestampRule []interface{}
	for _, checkpointTimestampItem := range checkpointTimestamp {
		checkpointTimestampRule = append(checkpointTimestampRule, checkpointTimestampItem)
	}
	var validatorIndexRule []interface{}
	for _, validatorIndexItem := range validatorIndex {
		validatorIndexRule = append(validatorIndexRule, validatorIndexItem)
	}

	logs, sub, err := _EigenPod.contract.WatchLogs(opts, "ValidatorWithdrawn", checkpointTimestampRule, validatorIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenPodValidatorWithdrawn)
				if err := _EigenPod.contract.UnpackLog(event, "ValidatorWithdrawn", log); err != nil {
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

// ParseValidatorWithdrawn is a log parse operation binding the contract event 0x2a02361ffa66cf2c2da4682c2355a6adcaa9f6c227b6e6563e68480f9587626a.
//
// Solidity: event ValidatorWithdrawn(uint64 indexed checkpointTimestamp, uint40 indexed validatorIndex)
func (_EigenPod *EigenPodFilterer) ParseValidatorWithdrawn(log types.Log) (*EigenPodValidatorWithdrawn, error) {
	event := new(EigenPodValidatorWithdrawn)
	if err := _EigenPod.contract.UnpackLog(event, "ValidatorWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenPodLogIterator is returned from FilterLog and is used to iterate over the raw logs and unpacked data for Log events raised by the EigenPod contract.
type EigenPodLogIterator struct {
	Event *EigenPodLog // Event containing the contract specifics and raw log

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
func (it *EigenPodLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenPodLog)
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
		it.Event = new(EigenPodLog)
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
func (it *EigenPodLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenPodLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenPodLog represents a Log event raised by the EigenPod contract.
type EigenPodLog struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLog is a free log retrieval operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_EigenPod *EigenPodFilterer) FilterLog(opts *bind.FilterOpts) (*EigenPodLogIterator, error) {

	logs, sub, err := _EigenPod.contract.FilterLogs(opts, "log")
	if err != nil {
		return nil, err
	}
	return &EigenPodLogIterator{contract: _EigenPod.contract, event: "log", logs: logs, sub: sub}, nil
}

// WatchLog is a free log subscription operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_EigenPod *EigenPodFilterer) WatchLog(opts *bind.WatchOpts, sink chan<- *EigenPodLog) (event.Subscription, error) {

	logs, sub, err := _EigenPod.contract.WatchLogs(opts, "log")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenPodLog)
				if err := _EigenPod.contract.UnpackLog(event, "log", log); err != nil {
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

// ParseLog is a log parse operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_EigenPod *EigenPodFilterer) ParseLog(log types.Log) (*EigenPodLog, error) {
	event := new(EigenPodLog)
	if err := _EigenPod.contract.UnpackLog(event, "log", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenPodLogAddressIterator is returned from FilterLogAddress and is used to iterate over the raw logs and unpacked data for LogAddress events raised by the EigenPod contract.
type EigenPodLogAddressIterator struct {
	Event *EigenPodLogAddress // Event containing the contract specifics and raw log

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
func (it *EigenPodLogAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenPodLogAddress)
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
		it.Event = new(EigenPodLogAddress)
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
func (it *EigenPodLogAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenPodLogAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenPodLogAddress represents a LogAddress event raised by the EigenPod contract.
type EigenPodLogAddress struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogAddress is a free log retrieval operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_EigenPod *EigenPodFilterer) FilterLogAddress(opts *bind.FilterOpts) (*EigenPodLogAddressIterator, error) {

	logs, sub, err := _EigenPod.contract.FilterLogs(opts, "log_address")
	if err != nil {
		return nil, err
	}
	return &EigenPodLogAddressIterator{contract: _EigenPod.contract, event: "log_address", logs: logs, sub: sub}, nil
}

// WatchLogAddress is a free log subscription operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_EigenPod *EigenPodFilterer) WatchLogAddress(opts *bind.WatchOpts, sink chan<- *EigenPodLogAddress) (event.Subscription, error) {

	logs, sub, err := _EigenPod.contract.WatchLogs(opts, "log_address")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenPodLogAddress)
				if err := _EigenPod.contract.UnpackLog(event, "log_address", log); err != nil {
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

// ParseLogAddress is a log parse operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_EigenPod *EigenPodFilterer) ParseLogAddress(log types.Log) (*EigenPodLogAddress, error) {
	event := new(EigenPodLogAddress)
	if err := _EigenPod.contract.UnpackLog(event, "log_address", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenPodLogArrayIterator is returned from FilterLogArray and is used to iterate over the raw logs and unpacked data for LogArray events raised by the EigenPod contract.
type EigenPodLogArrayIterator struct {
	Event *EigenPodLogArray // Event containing the contract specifics and raw log

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
func (it *EigenPodLogArrayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenPodLogArray)
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
		it.Event = new(EigenPodLogArray)
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
func (it *EigenPodLogArrayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenPodLogArrayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenPodLogArray represents a LogArray event raised by the EigenPod contract.
type EigenPodLogArray struct {
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray is a free log retrieval operation binding the contract event 0xfb102865d50addddf69da9b5aa1bced66c80cf869a5c8d0471a467e18ce9cab1.
//
// Solidity: event log_array(uint256[] val)
func (_EigenPod *EigenPodFilterer) FilterLogArray(opts *bind.FilterOpts) (*EigenPodLogArrayIterator, error) {

	logs, sub, err := _EigenPod.contract.FilterLogs(opts, "log_array")
	if err != nil {
		return nil, err
	}
	return &EigenPodLogArrayIterator{contract: _EigenPod.contract, event: "log_array", logs: logs, sub: sub}, nil
}

// WatchLogArray is a free log subscription operation binding the contract event 0xfb102865d50addddf69da9b5aa1bced66c80cf869a5c8d0471a467e18ce9cab1.
//
// Solidity: event log_array(uint256[] val)
func (_EigenPod *EigenPodFilterer) WatchLogArray(opts *bind.WatchOpts, sink chan<- *EigenPodLogArray) (event.Subscription, error) {

	logs, sub, err := _EigenPod.contract.WatchLogs(opts, "log_array")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenPodLogArray)
				if err := _EigenPod.contract.UnpackLog(event, "log_array", log); err != nil {
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

// ParseLogArray is a log parse operation binding the contract event 0xfb102865d50addddf69da9b5aa1bced66c80cf869a5c8d0471a467e18ce9cab1.
//
// Solidity: event log_array(uint256[] val)
func (_EigenPod *EigenPodFilterer) ParseLogArray(log types.Log) (*EigenPodLogArray, error) {
	event := new(EigenPodLogArray)
	if err := _EigenPod.contract.UnpackLog(event, "log_array", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenPodLogArray0Iterator is returned from FilterLogArray0 and is used to iterate over the raw logs and unpacked data for LogArray0 events raised by the EigenPod contract.
type EigenPodLogArray0Iterator struct {
	Event *EigenPodLogArray0 // Event containing the contract specifics and raw log

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
func (it *EigenPodLogArray0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenPodLogArray0)
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
		it.Event = new(EigenPodLogArray0)
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
func (it *EigenPodLogArray0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenPodLogArray0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenPodLogArray0 represents a LogArray0 event raised by the EigenPod contract.
type EigenPodLogArray0 struct {
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray0 is a free log retrieval operation binding the contract event 0x890a82679b470f2bd82816ed9b161f97d8b967f37fa3647c21d5bf39749e2dd5.
//
// Solidity: event log_array(int256[] val)
func (_EigenPod *EigenPodFilterer) FilterLogArray0(opts *bind.FilterOpts) (*EigenPodLogArray0Iterator, error) {

	logs, sub, err := _EigenPod.contract.FilterLogs(opts, "log_array0")
	if err != nil {
		return nil, err
	}
	return &EigenPodLogArray0Iterator{contract: _EigenPod.contract, event: "log_array0", logs: logs, sub: sub}, nil
}

// WatchLogArray0 is a free log subscription operation binding the contract event 0x890a82679b470f2bd82816ed9b161f97d8b967f37fa3647c21d5bf39749e2dd5.
//
// Solidity: event log_array(int256[] val)
func (_EigenPod *EigenPodFilterer) WatchLogArray0(opts *bind.WatchOpts, sink chan<- *EigenPodLogArray0) (event.Subscription, error) {

	logs, sub, err := _EigenPod.contract.WatchLogs(opts, "log_array0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenPodLogArray0)
				if err := _EigenPod.contract.UnpackLog(event, "log_array0", log); err != nil {
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

// ParseLogArray0 is a log parse operation binding the contract event 0x890a82679b470f2bd82816ed9b161f97d8b967f37fa3647c21d5bf39749e2dd5.
//
// Solidity: event log_array(int256[] val)
func (_EigenPod *EigenPodFilterer) ParseLogArray0(log types.Log) (*EigenPodLogArray0, error) {
	event := new(EigenPodLogArray0)
	if err := _EigenPod.contract.UnpackLog(event, "log_array0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenPodLogArray1Iterator is returned from FilterLogArray1 and is used to iterate over the raw logs and unpacked data for LogArray1 events raised by the EigenPod contract.
type EigenPodLogArray1Iterator struct {
	Event *EigenPodLogArray1 // Event containing the contract specifics and raw log

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
func (it *EigenPodLogArray1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenPodLogArray1)
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
		it.Event = new(EigenPodLogArray1)
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
func (it *EigenPodLogArray1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenPodLogArray1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenPodLogArray1 represents a LogArray1 event raised by the EigenPod contract.
type EigenPodLogArray1 struct {
	Val []common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray1 is a free log retrieval operation binding the contract event 0x40e1840f5769073d61bd01372d9b75baa9842d5629a0c99ff103be1178a8e9e2.
//
// Solidity: event log_array(address[] val)
func (_EigenPod *EigenPodFilterer) FilterLogArray1(opts *bind.FilterOpts) (*EigenPodLogArray1Iterator, error) {

	logs, sub, err := _EigenPod.contract.FilterLogs(opts, "log_array1")
	if err != nil {
		return nil, err
	}
	return &EigenPodLogArray1Iterator{contract: _EigenPod.contract, event: "log_array1", logs: logs, sub: sub}, nil
}

// WatchLogArray1 is a free log subscription operation binding the contract event 0x40e1840f5769073d61bd01372d9b75baa9842d5629a0c99ff103be1178a8e9e2.
//
// Solidity: event log_array(address[] val)
func (_EigenPod *EigenPodFilterer) WatchLogArray1(opts *bind.WatchOpts, sink chan<- *EigenPodLogArray1) (event.Subscription, error) {

	logs, sub, err := _EigenPod.contract.WatchLogs(opts, "log_array1")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenPodLogArray1)
				if err := _EigenPod.contract.UnpackLog(event, "log_array1", log); err != nil {
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

// ParseLogArray1 is a log parse operation binding the contract event 0x40e1840f5769073d61bd01372d9b75baa9842d5629a0c99ff103be1178a8e9e2.
//
// Solidity: event log_array(address[] val)
func (_EigenPod *EigenPodFilterer) ParseLogArray1(log types.Log) (*EigenPodLogArray1, error) {
	event := new(EigenPodLogArray1)
	if err := _EigenPod.contract.UnpackLog(event, "log_array1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenPodLogBytesIterator is returned from FilterLogBytes and is used to iterate over the raw logs and unpacked data for LogBytes events raised by the EigenPod contract.
type EigenPodLogBytesIterator struct {
	Event *EigenPodLogBytes // Event containing the contract specifics and raw log

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
func (it *EigenPodLogBytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenPodLogBytes)
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
		it.Event = new(EigenPodLogBytes)
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
func (it *EigenPodLogBytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenPodLogBytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenPodLogBytes represents a LogBytes event raised by the EigenPod contract.
type EigenPodLogBytes struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogBytes is a free log retrieval operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_EigenPod *EigenPodFilterer) FilterLogBytes(opts *bind.FilterOpts) (*EigenPodLogBytesIterator, error) {

	logs, sub, err := _EigenPod.contract.FilterLogs(opts, "log_bytes")
	if err != nil {
		return nil, err
	}
	return &EigenPodLogBytesIterator{contract: _EigenPod.contract, event: "log_bytes", logs: logs, sub: sub}, nil
}

// WatchLogBytes is a free log subscription operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_EigenPod *EigenPodFilterer) WatchLogBytes(opts *bind.WatchOpts, sink chan<- *EigenPodLogBytes) (event.Subscription, error) {

	logs, sub, err := _EigenPod.contract.WatchLogs(opts, "log_bytes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenPodLogBytes)
				if err := _EigenPod.contract.UnpackLog(event, "log_bytes", log); err != nil {
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

// ParseLogBytes is a log parse operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_EigenPod *EigenPodFilterer) ParseLogBytes(log types.Log) (*EigenPodLogBytes, error) {
	event := new(EigenPodLogBytes)
	if err := _EigenPod.contract.UnpackLog(event, "log_bytes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenPodLogBytes32Iterator is returned from FilterLogBytes32 and is used to iterate over the raw logs and unpacked data for LogBytes32 events raised by the EigenPod contract.
type EigenPodLogBytes32Iterator struct {
	Event *EigenPodLogBytes32 // Event containing the contract specifics and raw log

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
func (it *EigenPodLogBytes32Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenPodLogBytes32)
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
		it.Event = new(EigenPodLogBytes32)
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
func (it *EigenPodLogBytes32Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenPodLogBytes32Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenPodLogBytes32 represents a LogBytes32 event raised by the EigenPod contract.
type EigenPodLogBytes32 struct {
	Arg0 [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogBytes32 is a free log retrieval operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_EigenPod *EigenPodFilterer) FilterLogBytes32(opts *bind.FilterOpts) (*EigenPodLogBytes32Iterator, error) {

	logs, sub, err := _EigenPod.contract.FilterLogs(opts, "log_bytes32")
	if err != nil {
		return nil, err
	}
	return &EigenPodLogBytes32Iterator{contract: _EigenPod.contract, event: "log_bytes32", logs: logs, sub: sub}, nil
}

// WatchLogBytes32 is a free log subscription operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_EigenPod *EigenPodFilterer) WatchLogBytes32(opts *bind.WatchOpts, sink chan<- *EigenPodLogBytes32) (event.Subscription, error) {

	logs, sub, err := _EigenPod.contract.WatchLogs(opts, "log_bytes32")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenPodLogBytes32)
				if err := _EigenPod.contract.UnpackLog(event, "log_bytes32", log); err != nil {
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

// ParseLogBytes32 is a log parse operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_EigenPod *EigenPodFilterer) ParseLogBytes32(log types.Log) (*EigenPodLogBytes32, error) {
	event := new(EigenPodLogBytes32)
	if err := _EigenPod.contract.UnpackLog(event, "log_bytes32", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenPodLogIntIterator is returned from FilterLogInt and is used to iterate over the raw logs and unpacked data for LogInt events raised by the EigenPod contract.
type EigenPodLogIntIterator struct {
	Event *EigenPodLogInt // Event containing the contract specifics and raw log

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
func (it *EigenPodLogIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenPodLogInt)
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
		it.Event = new(EigenPodLogInt)
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
func (it *EigenPodLogIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenPodLogIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenPodLogInt represents a LogInt event raised by the EigenPod contract.
type EigenPodLogInt struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogInt is a free log retrieval operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_EigenPod *EigenPodFilterer) FilterLogInt(opts *bind.FilterOpts) (*EigenPodLogIntIterator, error) {

	logs, sub, err := _EigenPod.contract.FilterLogs(opts, "log_int")
	if err != nil {
		return nil, err
	}
	return &EigenPodLogIntIterator{contract: _EigenPod.contract, event: "log_int", logs: logs, sub: sub}, nil
}

// WatchLogInt is a free log subscription operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_EigenPod *EigenPodFilterer) WatchLogInt(opts *bind.WatchOpts, sink chan<- *EigenPodLogInt) (event.Subscription, error) {

	logs, sub, err := _EigenPod.contract.WatchLogs(opts, "log_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenPodLogInt)
				if err := _EigenPod.contract.UnpackLog(event, "log_int", log); err != nil {
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

// ParseLogInt is a log parse operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_EigenPod *EigenPodFilterer) ParseLogInt(log types.Log) (*EigenPodLogInt, error) {
	event := new(EigenPodLogInt)
	if err := _EigenPod.contract.UnpackLog(event, "log_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenPodLogNamedAddressIterator is returned from FilterLogNamedAddress and is used to iterate over the raw logs and unpacked data for LogNamedAddress events raised by the EigenPod contract.
type EigenPodLogNamedAddressIterator struct {
	Event *EigenPodLogNamedAddress // Event containing the contract specifics and raw log

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
func (it *EigenPodLogNamedAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenPodLogNamedAddress)
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
		it.Event = new(EigenPodLogNamedAddress)
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
func (it *EigenPodLogNamedAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenPodLogNamedAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenPodLogNamedAddress represents a LogNamedAddress event raised by the EigenPod contract.
type EigenPodLogNamedAddress struct {
	Key string
	Val common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedAddress is a free log retrieval operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_EigenPod *EigenPodFilterer) FilterLogNamedAddress(opts *bind.FilterOpts) (*EigenPodLogNamedAddressIterator, error) {

	logs, sub, err := _EigenPod.contract.FilterLogs(opts, "log_named_address")
	if err != nil {
		return nil, err
	}
	return &EigenPodLogNamedAddressIterator{contract: _EigenPod.contract, event: "log_named_address", logs: logs, sub: sub}, nil
}

// WatchLogNamedAddress is a free log subscription operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_EigenPod *EigenPodFilterer) WatchLogNamedAddress(opts *bind.WatchOpts, sink chan<- *EigenPodLogNamedAddress) (event.Subscription, error) {

	logs, sub, err := _EigenPod.contract.WatchLogs(opts, "log_named_address")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenPodLogNamedAddress)
				if err := _EigenPod.contract.UnpackLog(event, "log_named_address", log); err != nil {
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

// ParseLogNamedAddress is a log parse operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_EigenPod *EigenPodFilterer) ParseLogNamedAddress(log types.Log) (*EigenPodLogNamedAddress, error) {
	event := new(EigenPodLogNamedAddress)
	if err := _EigenPod.contract.UnpackLog(event, "log_named_address", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenPodLogNamedArrayIterator is returned from FilterLogNamedArray and is used to iterate over the raw logs and unpacked data for LogNamedArray events raised by the EigenPod contract.
type EigenPodLogNamedArrayIterator struct {
	Event *EigenPodLogNamedArray // Event containing the contract specifics and raw log

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
func (it *EigenPodLogNamedArrayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenPodLogNamedArray)
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
		it.Event = new(EigenPodLogNamedArray)
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
func (it *EigenPodLogNamedArrayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenPodLogNamedArrayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenPodLogNamedArray represents a LogNamedArray event raised by the EigenPod contract.
type EigenPodLogNamedArray struct {
	Key string
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray is a free log retrieval operation binding the contract event 0x00aaa39c9ffb5f567a4534380c737075702e1f7f14107fc95328e3b56c0325fb.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (_EigenPod *EigenPodFilterer) FilterLogNamedArray(opts *bind.FilterOpts) (*EigenPodLogNamedArrayIterator, error) {

	logs, sub, err := _EigenPod.contract.FilterLogs(opts, "log_named_array")
	if err != nil {
		return nil, err
	}
	return &EigenPodLogNamedArrayIterator{contract: _EigenPod.contract, event: "log_named_array", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray is a free log subscription operation binding the contract event 0x00aaa39c9ffb5f567a4534380c737075702e1f7f14107fc95328e3b56c0325fb.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (_EigenPod *EigenPodFilterer) WatchLogNamedArray(opts *bind.WatchOpts, sink chan<- *EigenPodLogNamedArray) (event.Subscription, error) {

	logs, sub, err := _EigenPod.contract.WatchLogs(opts, "log_named_array")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenPodLogNamedArray)
				if err := _EigenPod.contract.UnpackLog(event, "log_named_array", log); err != nil {
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

// ParseLogNamedArray is a log parse operation binding the contract event 0x00aaa39c9ffb5f567a4534380c737075702e1f7f14107fc95328e3b56c0325fb.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (_EigenPod *EigenPodFilterer) ParseLogNamedArray(log types.Log) (*EigenPodLogNamedArray, error) {
	event := new(EigenPodLogNamedArray)
	if err := _EigenPod.contract.UnpackLog(event, "log_named_array", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenPodLogNamedArray0Iterator is returned from FilterLogNamedArray0 and is used to iterate over the raw logs and unpacked data for LogNamedArray0 events raised by the EigenPod contract.
type EigenPodLogNamedArray0Iterator struct {
	Event *EigenPodLogNamedArray0 // Event containing the contract specifics and raw log

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
func (it *EigenPodLogNamedArray0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenPodLogNamedArray0)
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
		it.Event = new(EigenPodLogNamedArray0)
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
func (it *EigenPodLogNamedArray0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenPodLogNamedArray0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenPodLogNamedArray0 represents a LogNamedArray0 event raised by the EigenPod contract.
type EigenPodLogNamedArray0 struct {
	Key string
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray0 is a free log retrieval operation binding the contract event 0xa73eda09662f46dde729be4611385ff34fe6c44fbbc6f7e17b042b59a3445b57.
//
// Solidity: event log_named_array(string key, int256[] val)
func (_EigenPod *EigenPodFilterer) FilterLogNamedArray0(opts *bind.FilterOpts) (*EigenPodLogNamedArray0Iterator, error) {

	logs, sub, err := _EigenPod.contract.FilterLogs(opts, "log_named_array0")
	if err != nil {
		return nil, err
	}
	return &EigenPodLogNamedArray0Iterator{contract: _EigenPod.contract, event: "log_named_array0", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray0 is a free log subscription operation binding the contract event 0xa73eda09662f46dde729be4611385ff34fe6c44fbbc6f7e17b042b59a3445b57.
//
// Solidity: event log_named_array(string key, int256[] val)
func (_EigenPod *EigenPodFilterer) WatchLogNamedArray0(opts *bind.WatchOpts, sink chan<- *EigenPodLogNamedArray0) (event.Subscription, error) {

	logs, sub, err := _EigenPod.contract.WatchLogs(opts, "log_named_array0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenPodLogNamedArray0)
				if err := _EigenPod.contract.UnpackLog(event, "log_named_array0", log); err != nil {
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

// ParseLogNamedArray0 is a log parse operation binding the contract event 0xa73eda09662f46dde729be4611385ff34fe6c44fbbc6f7e17b042b59a3445b57.
//
// Solidity: event log_named_array(string key, int256[] val)
func (_EigenPod *EigenPodFilterer) ParseLogNamedArray0(log types.Log) (*EigenPodLogNamedArray0, error) {
	event := new(EigenPodLogNamedArray0)
	if err := _EigenPod.contract.UnpackLog(event, "log_named_array0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenPodLogNamedArray1Iterator is returned from FilterLogNamedArray1 and is used to iterate over the raw logs and unpacked data for LogNamedArray1 events raised by the EigenPod contract.
type EigenPodLogNamedArray1Iterator struct {
	Event *EigenPodLogNamedArray1 // Event containing the contract specifics and raw log

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
func (it *EigenPodLogNamedArray1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenPodLogNamedArray1)
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
		it.Event = new(EigenPodLogNamedArray1)
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
func (it *EigenPodLogNamedArray1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenPodLogNamedArray1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenPodLogNamedArray1 represents a LogNamedArray1 event raised by the EigenPod contract.
type EigenPodLogNamedArray1 struct {
	Key string
	Val []common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray1 is a free log retrieval operation binding the contract event 0x3bcfb2ae2e8d132dd1fce7cf278a9a19756a9fceabe470df3bdabb4bc577d1bd.
//
// Solidity: event log_named_array(string key, address[] val)
func (_EigenPod *EigenPodFilterer) FilterLogNamedArray1(opts *bind.FilterOpts) (*EigenPodLogNamedArray1Iterator, error) {

	logs, sub, err := _EigenPod.contract.FilterLogs(opts, "log_named_array1")
	if err != nil {
		return nil, err
	}
	return &EigenPodLogNamedArray1Iterator{contract: _EigenPod.contract, event: "log_named_array1", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray1 is a free log subscription operation binding the contract event 0x3bcfb2ae2e8d132dd1fce7cf278a9a19756a9fceabe470df3bdabb4bc577d1bd.
//
// Solidity: event log_named_array(string key, address[] val)
func (_EigenPod *EigenPodFilterer) WatchLogNamedArray1(opts *bind.WatchOpts, sink chan<- *EigenPodLogNamedArray1) (event.Subscription, error) {

	logs, sub, err := _EigenPod.contract.WatchLogs(opts, "log_named_array1")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenPodLogNamedArray1)
				if err := _EigenPod.contract.UnpackLog(event, "log_named_array1", log); err != nil {
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

// ParseLogNamedArray1 is a log parse operation binding the contract event 0x3bcfb2ae2e8d132dd1fce7cf278a9a19756a9fceabe470df3bdabb4bc577d1bd.
//
// Solidity: event log_named_array(string key, address[] val)
func (_EigenPod *EigenPodFilterer) ParseLogNamedArray1(log types.Log) (*EigenPodLogNamedArray1, error) {
	event := new(EigenPodLogNamedArray1)
	if err := _EigenPod.contract.UnpackLog(event, "log_named_array1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenPodLogNamedBytesIterator is returned from FilterLogNamedBytes and is used to iterate over the raw logs and unpacked data for LogNamedBytes events raised by the EigenPod contract.
type EigenPodLogNamedBytesIterator struct {
	Event *EigenPodLogNamedBytes // Event containing the contract specifics and raw log

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
func (it *EigenPodLogNamedBytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenPodLogNamedBytes)
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
		it.Event = new(EigenPodLogNamedBytes)
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
func (it *EigenPodLogNamedBytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenPodLogNamedBytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenPodLogNamedBytes represents a LogNamedBytes event raised by the EigenPod contract.
type EigenPodLogNamedBytes struct {
	Key string
	Val []byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedBytes is a free log retrieval operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_EigenPod *EigenPodFilterer) FilterLogNamedBytes(opts *bind.FilterOpts) (*EigenPodLogNamedBytesIterator, error) {

	logs, sub, err := _EigenPod.contract.FilterLogs(opts, "log_named_bytes")
	if err != nil {
		return nil, err
	}
	return &EigenPodLogNamedBytesIterator{contract: _EigenPod.contract, event: "log_named_bytes", logs: logs, sub: sub}, nil
}

// WatchLogNamedBytes is a free log subscription operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_EigenPod *EigenPodFilterer) WatchLogNamedBytes(opts *bind.WatchOpts, sink chan<- *EigenPodLogNamedBytes) (event.Subscription, error) {

	logs, sub, err := _EigenPod.contract.WatchLogs(opts, "log_named_bytes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenPodLogNamedBytes)
				if err := _EigenPod.contract.UnpackLog(event, "log_named_bytes", log); err != nil {
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

// ParseLogNamedBytes is a log parse operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_EigenPod *EigenPodFilterer) ParseLogNamedBytes(log types.Log) (*EigenPodLogNamedBytes, error) {
	event := new(EigenPodLogNamedBytes)
	if err := _EigenPod.contract.UnpackLog(event, "log_named_bytes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenPodLogNamedBytes32Iterator is returned from FilterLogNamedBytes32 and is used to iterate over the raw logs and unpacked data for LogNamedBytes32 events raised by the EigenPod contract.
type EigenPodLogNamedBytes32Iterator struct {
	Event *EigenPodLogNamedBytes32 // Event containing the contract specifics and raw log

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
func (it *EigenPodLogNamedBytes32Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenPodLogNamedBytes32)
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
		it.Event = new(EigenPodLogNamedBytes32)
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
func (it *EigenPodLogNamedBytes32Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenPodLogNamedBytes32Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenPodLogNamedBytes32 represents a LogNamedBytes32 event raised by the EigenPod contract.
type EigenPodLogNamedBytes32 struct {
	Key string
	Val [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedBytes32 is a free log retrieval operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_EigenPod *EigenPodFilterer) FilterLogNamedBytes32(opts *bind.FilterOpts) (*EigenPodLogNamedBytes32Iterator, error) {

	logs, sub, err := _EigenPod.contract.FilterLogs(opts, "log_named_bytes32")
	if err != nil {
		return nil, err
	}
	return &EigenPodLogNamedBytes32Iterator{contract: _EigenPod.contract, event: "log_named_bytes32", logs: logs, sub: sub}, nil
}

// WatchLogNamedBytes32 is a free log subscription operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_EigenPod *EigenPodFilterer) WatchLogNamedBytes32(opts *bind.WatchOpts, sink chan<- *EigenPodLogNamedBytes32) (event.Subscription, error) {

	logs, sub, err := _EigenPod.contract.WatchLogs(opts, "log_named_bytes32")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenPodLogNamedBytes32)
				if err := _EigenPod.contract.UnpackLog(event, "log_named_bytes32", log); err != nil {
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

// ParseLogNamedBytes32 is a log parse operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_EigenPod *EigenPodFilterer) ParseLogNamedBytes32(log types.Log) (*EigenPodLogNamedBytes32, error) {
	event := new(EigenPodLogNamedBytes32)
	if err := _EigenPod.contract.UnpackLog(event, "log_named_bytes32", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenPodLogNamedDecimalIntIterator is returned from FilterLogNamedDecimalInt and is used to iterate over the raw logs and unpacked data for LogNamedDecimalInt events raised by the EigenPod contract.
type EigenPodLogNamedDecimalIntIterator struct {
	Event *EigenPodLogNamedDecimalInt // Event containing the contract specifics and raw log

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
func (it *EigenPodLogNamedDecimalIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenPodLogNamedDecimalInt)
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
		it.Event = new(EigenPodLogNamedDecimalInt)
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
func (it *EigenPodLogNamedDecimalIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenPodLogNamedDecimalIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenPodLogNamedDecimalInt represents a LogNamedDecimalInt event raised by the EigenPod contract.
type EigenPodLogNamedDecimalInt struct {
	Key      string
	Val      *big.Int
	Decimals *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogNamedDecimalInt is a free log retrieval operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_EigenPod *EigenPodFilterer) FilterLogNamedDecimalInt(opts *bind.FilterOpts) (*EigenPodLogNamedDecimalIntIterator, error) {

	logs, sub, err := _EigenPod.contract.FilterLogs(opts, "log_named_decimal_int")
	if err != nil {
		return nil, err
	}
	return &EigenPodLogNamedDecimalIntIterator{contract: _EigenPod.contract, event: "log_named_decimal_int", logs: logs, sub: sub}, nil
}

// WatchLogNamedDecimalInt is a free log subscription operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_EigenPod *EigenPodFilterer) WatchLogNamedDecimalInt(opts *bind.WatchOpts, sink chan<- *EigenPodLogNamedDecimalInt) (event.Subscription, error) {

	logs, sub, err := _EigenPod.contract.WatchLogs(opts, "log_named_decimal_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenPodLogNamedDecimalInt)
				if err := _EigenPod.contract.UnpackLog(event, "log_named_decimal_int", log); err != nil {
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

// ParseLogNamedDecimalInt is a log parse operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_EigenPod *EigenPodFilterer) ParseLogNamedDecimalInt(log types.Log) (*EigenPodLogNamedDecimalInt, error) {
	event := new(EigenPodLogNamedDecimalInt)
	if err := _EigenPod.contract.UnpackLog(event, "log_named_decimal_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenPodLogNamedDecimalUintIterator is returned from FilterLogNamedDecimalUint and is used to iterate over the raw logs and unpacked data for LogNamedDecimalUint events raised by the EigenPod contract.
type EigenPodLogNamedDecimalUintIterator struct {
	Event *EigenPodLogNamedDecimalUint // Event containing the contract specifics and raw log

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
func (it *EigenPodLogNamedDecimalUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenPodLogNamedDecimalUint)
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
		it.Event = new(EigenPodLogNamedDecimalUint)
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
func (it *EigenPodLogNamedDecimalUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenPodLogNamedDecimalUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenPodLogNamedDecimalUint represents a LogNamedDecimalUint event raised by the EigenPod contract.
type EigenPodLogNamedDecimalUint struct {
	Key      string
	Val      *big.Int
	Decimals *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogNamedDecimalUint is a free log retrieval operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_EigenPod *EigenPodFilterer) FilterLogNamedDecimalUint(opts *bind.FilterOpts) (*EigenPodLogNamedDecimalUintIterator, error) {

	logs, sub, err := _EigenPod.contract.FilterLogs(opts, "log_named_decimal_uint")
	if err != nil {
		return nil, err
	}
	return &EigenPodLogNamedDecimalUintIterator{contract: _EigenPod.contract, event: "log_named_decimal_uint", logs: logs, sub: sub}, nil
}

// WatchLogNamedDecimalUint is a free log subscription operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_EigenPod *EigenPodFilterer) WatchLogNamedDecimalUint(opts *bind.WatchOpts, sink chan<- *EigenPodLogNamedDecimalUint) (event.Subscription, error) {

	logs, sub, err := _EigenPod.contract.WatchLogs(opts, "log_named_decimal_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenPodLogNamedDecimalUint)
				if err := _EigenPod.contract.UnpackLog(event, "log_named_decimal_uint", log); err != nil {
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

// ParseLogNamedDecimalUint is a log parse operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_EigenPod *EigenPodFilterer) ParseLogNamedDecimalUint(log types.Log) (*EigenPodLogNamedDecimalUint, error) {
	event := new(EigenPodLogNamedDecimalUint)
	if err := _EigenPod.contract.UnpackLog(event, "log_named_decimal_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenPodLogNamedIntIterator is returned from FilterLogNamedInt and is used to iterate over the raw logs and unpacked data for LogNamedInt events raised by the EigenPod contract.
type EigenPodLogNamedIntIterator struct {
	Event *EigenPodLogNamedInt // Event containing the contract specifics and raw log

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
func (it *EigenPodLogNamedIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenPodLogNamedInt)
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
		it.Event = new(EigenPodLogNamedInt)
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
func (it *EigenPodLogNamedIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenPodLogNamedIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenPodLogNamedInt represents a LogNamedInt event raised by the EigenPod contract.
type EigenPodLogNamedInt struct {
	Key string
	Val *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedInt is a free log retrieval operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_EigenPod *EigenPodFilterer) FilterLogNamedInt(opts *bind.FilterOpts) (*EigenPodLogNamedIntIterator, error) {

	logs, sub, err := _EigenPod.contract.FilterLogs(opts, "log_named_int")
	if err != nil {
		return nil, err
	}
	return &EigenPodLogNamedIntIterator{contract: _EigenPod.contract, event: "log_named_int", logs: logs, sub: sub}, nil
}

// WatchLogNamedInt is a free log subscription operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_EigenPod *EigenPodFilterer) WatchLogNamedInt(opts *bind.WatchOpts, sink chan<- *EigenPodLogNamedInt) (event.Subscription, error) {

	logs, sub, err := _EigenPod.contract.WatchLogs(opts, "log_named_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenPodLogNamedInt)
				if err := _EigenPod.contract.UnpackLog(event, "log_named_int", log); err != nil {
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

// ParseLogNamedInt is a log parse operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_EigenPod *EigenPodFilterer) ParseLogNamedInt(log types.Log) (*EigenPodLogNamedInt, error) {
	event := new(EigenPodLogNamedInt)
	if err := _EigenPod.contract.UnpackLog(event, "log_named_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenPodLogNamedStringIterator is returned from FilterLogNamedString and is used to iterate over the raw logs and unpacked data for LogNamedString events raised by the EigenPod contract.
type EigenPodLogNamedStringIterator struct {
	Event *EigenPodLogNamedString // Event containing the contract specifics and raw log

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
func (it *EigenPodLogNamedStringIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenPodLogNamedString)
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
		it.Event = new(EigenPodLogNamedString)
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
func (it *EigenPodLogNamedStringIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenPodLogNamedStringIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenPodLogNamedString represents a LogNamedString event raised by the EigenPod contract.
type EigenPodLogNamedString struct {
	Key string
	Val string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedString is a free log retrieval operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_EigenPod *EigenPodFilterer) FilterLogNamedString(opts *bind.FilterOpts) (*EigenPodLogNamedStringIterator, error) {

	logs, sub, err := _EigenPod.contract.FilterLogs(opts, "log_named_string")
	if err != nil {
		return nil, err
	}
	return &EigenPodLogNamedStringIterator{contract: _EigenPod.contract, event: "log_named_string", logs: logs, sub: sub}, nil
}

// WatchLogNamedString is a free log subscription operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_EigenPod *EigenPodFilterer) WatchLogNamedString(opts *bind.WatchOpts, sink chan<- *EigenPodLogNamedString) (event.Subscription, error) {

	logs, sub, err := _EigenPod.contract.WatchLogs(opts, "log_named_string")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenPodLogNamedString)
				if err := _EigenPod.contract.UnpackLog(event, "log_named_string", log); err != nil {
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

// ParseLogNamedString is a log parse operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_EigenPod *EigenPodFilterer) ParseLogNamedString(log types.Log) (*EigenPodLogNamedString, error) {
	event := new(EigenPodLogNamedString)
	if err := _EigenPod.contract.UnpackLog(event, "log_named_string", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenPodLogNamedUintIterator is returned from FilterLogNamedUint and is used to iterate over the raw logs and unpacked data for LogNamedUint events raised by the EigenPod contract.
type EigenPodLogNamedUintIterator struct {
	Event *EigenPodLogNamedUint // Event containing the contract specifics and raw log

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
func (it *EigenPodLogNamedUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenPodLogNamedUint)
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
		it.Event = new(EigenPodLogNamedUint)
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
func (it *EigenPodLogNamedUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenPodLogNamedUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenPodLogNamedUint represents a LogNamedUint event raised by the EigenPod contract.
type EigenPodLogNamedUint struct {
	Key string
	Val *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedUint is a free log retrieval operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_EigenPod *EigenPodFilterer) FilterLogNamedUint(opts *bind.FilterOpts) (*EigenPodLogNamedUintIterator, error) {

	logs, sub, err := _EigenPod.contract.FilterLogs(opts, "log_named_uint")
	if err != nil {
		return nil, err
	}
	return &EigenPodLogNamedUintIterator{contract: _EigenPod.contract, event: "log_named_uint", logs: logs, sub: sub}, nil
}

// WatchLogNamedUint is a free log subscription operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_EigenPod *EigenPodFilterer) WatchLogNamedUint(opts *bind.WatchOpts, sink chan<- *EigenPodLogNamedUint) (event.Subscription, error) {

	logs, sub, err := _EigenPod.contract.WatchLogs(opts, "log_named_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenPodLogNamedUint)
				if err := _EigenPod.contract.UnpackLog(event, "log_named_uint", log); err != nil {
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

// ParseLogNamedUint is a log parse operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_EigenPod *EigenPodFilterer) ParseLogNamedUint(log types.Log) (*EigenPodLogNamedUint, error) {
	event := new(EigenPodLogNamedUint)
	if err := _EigenPod.contract.UnpackLog(event, "log_named_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenPodLogStringIterator is returned from FilterLogString and is used to iterate over the raw logs and unpacked data for LogString events raised by the EigenPod contract.
type EigenPodLogStringIterator struct {
	Event *EigenPodLogString // Event containing the contract specifics and raw log

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
func (it *EigenPodLogStringIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenPodLogString)
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
		it.Event = new(EigenPodLogString)
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
func (it *EigenPodLogStringIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenPodLogStringIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenPodLogString represents a LogString event raised by the EigenPod contract.
type EigenPodLogString struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogString is a free log retrieval operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_EigenPod *EigenPodFilterer) FilterLogString(opts *bind.FilterOpts) (*EigenPodLogStringIterator, error) {

	logs, sub, err := _EigenPod.contract.FilterLogs(opts, "log_string")
	if err != nil {
		return nil, err
	}
	return &EigenPodLogStringIterator{contract: _EigenPod.contract, event: "log_string", logs: logs, sub: sub}, nil
}

// WatchLogString is a free log subscription operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_EigenPod *EigenPodFilterer) WatchLogString(opts *bind.WatchOpts, sink chan<- *EigenPodLogString) (event.Subscription, error) {

	logs, sub, err := _EigenPod.contract.WatchLogs(opts, "log_string")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenPodLogString)
				if err := _EigenPod.contract.UnpackLog(event, "log_string", log); err != nil {
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

// ParseLogString is a log parse operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_EigenPod *EigenPodFilterer) ParseLogString(log types.Log) (*EigenPodLogString, error) {
	event := new(EigenPodLogString)
	if err := _EigenPod.contract.UnpackLog(event, "log_string", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenPodLogUintIterator is returned from FilterLogUint and is used to iterate over the raw logs and unpacked data for LogUint events raised by the EigenPod contract.
type EigenPodLogUintIterator struct {
	Event *EigenPodLogUint // Event containing the contract specifics and raw log

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
func (it *EigenPodLogUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenPodLogUint)
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
		it.Event = new(EigenPodLogUint)
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
func (it *EigenPodLogUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenPodLogUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenPodLogUint represents a LogUint event raised by the EigenPod contract.
type EigenPodLogUint struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogUint is a free log retrieval operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_EigenPod *EigenPodFilterer) FilterLogUint(opts *bind.FilterOpts) (*EigenPodLogUintIterator, error) {

	logs, sub, err := _EigenPod.contract.FilterLogs(opts, "log_uint")
	if err != nil {
		return nil, err
	}
	return &EigenPodLogUintIterator{contract: _EigenPod.contract, event: "log_uint", logs: logs, sub: sub}, nil
}

// WatchLogUint is a free log subscription operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_EigenPod *EigenPodFilterer) WatchLogUint(opts *bind.WatchOpts, sink chan<- *EigenPodLogUint) (event.Subscription, error) {

	logs, sub, err := _EigenPod.contract.WatchLogs(opts, "log_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenPodLogUint)
				if err := _EigenPod.contract.UnpackLog(event, "log_uint", log); err != nil {
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

// ParseLogUint is a log parse operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_EigenPod *EigenPodFilterer) ParseLogUint(log types.Log) (*EigenPodLogUint, error) {
	event := new(EigenPodLogUint)
	if err := _EigenPod.contract.UnpackLog(event, "log_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenPodLogsIterator is returned from FilterLogs and is used to iterate over the raw logs and unpacked data for Logs events raised by the EigenPod contract.
type EigenPodLogsIterator struct {
	Event *EigenPodLogs // Event containing the contract specifics and raw log

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
func (it *EigenPodLogsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenPodLogs)
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
		it.Event = new(EigenPodLogs)
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
func (it *EigenPodLogsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenPodLogsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenPodLogs represents a Logs event raised by the EigenPod contract.
type EigenPodLogs struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogs is a free log retrieval operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_EigenPod *EigenPodFilterer) FilterLogs(opts *bind.FilterOpts) (*EigenPodLogsIterator, error) {

	logs, sub, err := _EigenPod.contract.FilterLogs(opts, "logs")
	if err != nil {
		return nil, err
	}
	return &EigenPodLogsIterator{contract: _EigenPod.contract, event: "logs", logs: logs, sub: sub}, nil
}

// WatchLogs is a free log subscription operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_EigenPod *EigenPodFilterer) WatchLogs(opts *bind.WatchOpts, sink chan<- *EigenPodLogs) (event.Subscription, error) {

	logs, sub, err := _EigenPod.contract.WatchLogs(opts, "logs")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenPodLogs)
				if err := _EigenPod.contract.UnpackLog(event, "logs", log); err != nil {
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

// ParseLogs is a log parse operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_EigenPod *EigenPodFilterer) ParseLogs(log types.Log) (*EigenPodLogs, error) {
	event := new(EigenPodLogs)
	if err := _EigenPod.contract.UnpackLog(event, "logs", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
